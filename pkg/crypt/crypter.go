package crypt

type Crypter interface {
	GenRandKey(cryptoType CryptoType) []byte
	Encrypt(key, nonce, plaintext []byte) ([]byte, []byte, error)
	Decrypt(key, nonce, ciphertext []byte) ([]byte, error)
}

const (
	CRYPTO_TYPE_UNSPECIFIED = iota
	CRYPTO_TYPE_GCM_AES128
	CRYPTO_TYPE_GCM_AES256
)

const (
	NonceSize = 12
)

type CryptoType int

func NewCrypter(cryptoType CryptoType) Crypter {
	switch cryptoType {
	case CRYPTO_TYPE_GCM_AES128, CRYPTO_TYPE_GCM_AES256:
		return &crypterGCM{}
	default:
		return nil
	}
}
