package cli

import (
	"dualbox/crypt"
	"dualbox/loader"
	"dualbox/pack"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

func Enc() *cli.Command {
	var filePaths [2]string
	var keys [2]string
	var output string
	return &cli.Command{
		Name:  "enc",
		Usage: "encrypt files with given keys",
		Action: func(ctx *cli.Context) error {
			return enc(filePaths[:], keys[:], output)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file0",
				Aliases:     []string{"f0"},
				Usage:       "target file path with index 0",
				Destination: &filePaths[0],
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "file1",
				Aliases:     []string{"f1"},
				Usage:       "target file path with index 1",
				Destination: &filePaths[1],
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "key0",
				Aliases:     []string{"k0"},
				Usage:       "encryption key with index 0",
				Destination: &keys[0],
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "key1",
				Aliases:     []string{"k1"},
				Usage:       "encryption key with index 1",
				Destination: &keys[1],
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output file name",
				Destination: &output,
			},
		},
	}
}

func enc(filePaths []string, keys []string, output string) error {
	cpt := crypt.NewCrypter(crypt.CRYPTO_TYPE_GCM_AES256)

	b0, err := loader.OpenPlainFile(filePaths[0])
	if err != nil {
		return fmt.Errorf("failed to open file0: %w", err)
	}

	b1, err := loader.OpenPlainFile(filePaths[1])
	if err != nil {
		return fmt.Errorf("failed to open file1: %w", err)
	}

	key0, err := hex.DecodeString(keys[0])
	if err != nil {
		return fmt.Errorf("failed to decode key0 in hex: %w", err)
	}
	key1, err := hex.DecodeString(keys[1])
	if err != nil {
		return fmt.Errorf("failed to decode key1 in hex: %w", err)
	}

	c0, n0, err := cpt.Encrypt(key0, nil, b0)
	if err != nil {
		return fmt.Errorf("failed to encrypt file0: %w", err)
	}
	c1, n1, err := cpt.Encrypt(key1, nil, b1)
	if err != nil {
		return fmt.Errorf("failed to encrypt file1: %w", err)
	}

	b, err := pack.Pack(cpt, key0, key1, n0, n1, c0, c1)
	if err != nil {
		return fmt.Errorf("failed to pack files: %w", err)
	}

	if output == "" {
		cipherFilename := time.Now().Format(time.RFC3339) + loader.FilenameExtension
		if err = loader.WriteFile(cipherFilename, b); err != nil {
			return fmt.Errorf("failed to cipher to file: %w", err)
		}
		return nil
	}

	if err = loader.WriteFile(output, b); err != nil {
		return fmt.Errorf("failed to cipher to file: %w", err)
	}
	return nil
}
