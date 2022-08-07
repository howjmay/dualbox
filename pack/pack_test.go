package pack

import (
	"crypto/rand"
	"dualbox/crypt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PackUnpack(t *testing.T) {
	cpt := crypt.NewCrypter(crypt.CRYPTO_TYPE_GCM_AES256)

	b0 := make([]byte, 40)
	rand.Read(b0)
	b1 := make([]byte, 40)
	rand.Read(b1)

	key0 := cpt.GenRandKey(crypt.CRYPTO_TYPE_GCM_AES256)
	key1 := cpt.GenRandKey(crypt.CRYPTO_TYPE_GCM_AES256)
	c0, n0, err := cpt.Encrypt(key0, nil, b0)
	require.NoError(t, err)
	c1, n1, err := cpt.Encrypt(key1, nil, b1)
	require.NoError(t, err)

	b, err := Pack(cpt, key0, key1, n0, n1, c0, c1)
	require.NoError(t, err)

	decrypttext0, err := Unpack(cpt, b, key0)
	require.NoError(t, err)
	require.Equal(t, b0, decrypttext0)
	decrypttext1, err := Unpack(cpt, b, key1)
	require.NoError(t, err)
	require.Equal(t, b1, decrypttext1)
}
