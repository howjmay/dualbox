package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/sha3"
)

type crypterGCM struct {
	cryptoType CryptoType
}

func (c *crypterGCM) PasswordToKey(pwd string) ([]byte, error) {
	switch c.cryptoType {
	case CRYPTO_TYPE_GCM_AES128:
		h := sha3.New256()
		_, err := h.Write([]byte(pwd))
		if err != nil {
			return nil, err
		}
		buf := h.Sum(nil)
		ret := make([]byte, 16)
		bufHalfLen := len(buf)
		for i := 0; i < bufHalfLen; i++ {
			ret[i] = buf[i] ^ buf[i+bufHalfLen]
		}
		return ret, nil

	case CRYPTO_TYPE_GCM_AES256:
		h := sha3.New256()
		_, err := h.Write([]byte(pwd))
		if err != nil {
			return nil, err
		}
		return h.Sum(nil), nil
	default:
		panic(fmt.Sprintf("unsupported crypto type: %d", c.cryptoType))
	}
}

func (c *crypterGCM) GenRandKey() []byte {
	switch c.cryptoType {
	case CRYPTO_TYPE_GCM_AES128:
		key := make([]byte, 16)
		rand.Read(key)
		return key
	case CRYPTO_TYPE_GCM_AES256:
		key := make([]byte, 32)
		rand.Read(key)
		return key
	default:
		panic(fmt.Sprintf("unsupported crypto type: %d", c.cryptoType))
	}
}

func (c *crypterGCM) Encrypt(key, nonce, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create new cipher in Encrypt(): %w", err)
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
		return nil, nil, fmt.Errorf("failed to call NewGCM in Encrypt(): %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nonce, nil
}

func (c *crypterGCM) Decrypt(key, nonce, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create new cipher in Decrypt(): %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to call NewGCM in Decrypt(): %w", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open cipher: %w", err)
	}

	return plaintext, nil
}
