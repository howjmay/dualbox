package pack

import (
	"dualbox/pkg/crypt"
	"dualbox/pkg/utils"

	"github.com/sirupsen/logrus"
)

func Pack(cpt crypt.Crypter, key0, key1, nonce0, nonce1, cipher0, cipher1 []byte) ([]byte, error) {
	from0 := HeaderSize * 2
	to0 := from0 + len(cipher0)
	buf0 := make([]byte, CheckSumSize)
	copy(buf0, utils.Uint32ToBytes(uint32(from0)))
	copy(buf0[ChecksumFromToSize:], utils.Uint32ToBytes(uint32(to0)))
	checksum0, _, err := cpt.Encrypt(key0, nonce0, buf0)
	if err != nil {
		logrus.Fatal(err)
	}

	from1 := to0
	to1 := from1 + len(cipher1)
	buf1 := make([]byte, CheckSumSize)
	copy(buf1, utils.Uint32ToBytes(uint32(from1)))
	copy(buf1[ChecksumFromToSize:], utils.Uint32ToBytes(uint32(to1)))
	checksum1, _, err := cpt.Encrypt(key1, nonce1, buf1)
	if err != nil {
		logrus.Fatal(err)
	}

	ret := make([]byte, to1)
	copy(ret, nonce0)
	copy(ret[crypt.NonceSize:], checksum0)
	copy(ret[HeaderSize:], nonce1)
	copy(ret[HeaderSize+crypt.NonceSize:], checksum1)
	copy(ret[HeaderSize*2:], cipher0)
	copy(ret[HeaderSize*2+len(cipher0):], cipher1)

	return ret, nil
}
