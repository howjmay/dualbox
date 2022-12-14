package cli

import (
	"dualbox/crypt"
	"dualbox/loader"
	"dualbox/pack"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

func Dec() *cli.Command {
	var filePath string
	var key string
	var pwd string
	return &cli.Command{
		Name:  "dec",
		Usage: "decrypt file with given key",
		Action: func(ctx *cli.Context) error {
			return dec(filePath, key, pwd)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "target decrypt file path",
				Destination: &filePath,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k"},
				Usage:       "decryption key",
				Destination: &key,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p", "pwd"},
				Usage:       "decryption password",
				Destination: &pwd,
			},
		},
	}
}

func dec(filePath, keyRaw, pwd string) error {
	var key []byte
	cpt := crypt.NewCrypter(crypt.CRYPTO_TYPE_GCM_AES256)

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %s: %w", filePath, err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file: %s: %w", filePath, err)
	}

	if pwd != "" {
		key, err = cpt.PasswordToKey(pwd)
		if err != nil {
			return fmt.Errorf("failed to convert pwd to key: %w", err)
		}
	} else {
		key, err = hex.DecodeString(keyRaw)
		if err != nil {
			return fmt.Errorf("failed to decode key in hex: %w", err)
		}
	}

	plaintext, err := pack.Unpack(cpt, b, key)
	if err != nil {
		return fmt.Errorf("failed to unpack encrypted file: %w", err)
	}

	err = loader.WriteFile("", plaintext)
	if err != nil {
		return fmt.Errorf("failed to write decrypted file: %w", err)
	}

	return nil
}
