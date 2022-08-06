package crypt

type Crypter interface {
	Encrypt(plaintext []byte) ([]byte, []byte, error)
	Decrypt(ciphertext, nonce []byte) ([]byte, error)
}

const (
	CRYPTO_TYPE_UNSPECIFIED = iota
	CRYPTO_TYPE_GCM
)

type CryptoType int

func NewCrypter(cryptoType CryptoType) Crypter {
	switch cryptoType {
	case CRYPTO_TYPE_GCM:
		return &crypterGCM{}
	default:
		return nil
	}
}
