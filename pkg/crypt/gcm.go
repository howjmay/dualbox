package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type GCMType int

const (
	GCM_TYPE_UNSPECIFIED = iota
	GCM_TYPE_AES128
	GCM_TYPE_AES256
)

type crypterGCM struct {
	key []byte
}

func (c *crypterGCM) GenRandKey(gcmType GCMType) []byte {
	switch gcmType {
	case GCM_TYPE_AES128:
		key := make([]byte, 16)
		rand.Read(key)
		return key
	case GCM_TYPE_AES256:
		key := make([]byte, 32)
		rand.Read(key)
		return key
	default:
		return nil
	}
}

func (c *crypterGCM) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	c.key = c.GenRandKey(GCM_TYPE_AES256)

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)

	return ciphertext, nonce, err
}

func (c *crypterGCM) Decrypt(ciphertext, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
