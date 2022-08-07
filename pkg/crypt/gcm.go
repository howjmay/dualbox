package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type crypterGCM struct {
	key []byte
}

func (c *crypterGCM) GenRandKey(cryptoType CryptoType) []byte {
	switch cryptoType {
	case CRYPTO_TYPE_GCM_AES128:
		key := make([]byte, 16)
		rand.Read(key)
		return key
	case CRYPTO_TYPE_GCM_AES256:
		key := make([]byte, 32)
		rand.Read(key)
		return key
	default:
		return nil
	}
}

func (c *crypterGCM) Encrypt(key, nonce, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	if nonce == nil {
		nonce = make([]byte, NonceSize)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return nil, nil, err
		}
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nonce, err
}

func (c *crypterGCM) Decrypt(key, nonce, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
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
