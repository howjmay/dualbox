package pack

import (
	"dualbox/pkg/crypt"
	"dualbox/pkg/utils"

	"github.com/sirupsen/logrus"
)

func Pack(cpt crypt.Crypter, key0, key1, nonce0, nonce1, cipher0, cipher1 []byte) ([]byte, error) {
	buf0 := []byte(ChecksumMsg)
	from0 := HeaderSize * 2
	to0 := from0 + len(cipher0)
	buf0 = append(buf0, utils.Uint32ToBytes(uint32(from0))...)
	buf0 = append(buf0, utils.Uint32ToBytes(uint32(to0))...)
	checksum0, _, err := cpt.Encrypt(key0, nonce0, buf0)
	if err != nil {
		logrus.Fatal(err)
	}

	buf1 := []byte(ChecksumMsg)
	from1 := to0
	to1 := from1 + len(cipher1)
	buf1 = append(buf1, utils.Uint32ToBytes(uint32(from1))...)
	buf1 = append(buf1, utils.Uint32ToBytes(uint32(to1))...)
	checksum1, _, err := cpt.Encrypt(key1, nonce1, buf1)
	if err != nil {
		logrus.Fatal(err)
	}

	var ret []byte
	ret = append(ret, nonce0...)
	ret = append(ret, checksum0...)
	ret = append(ret, nonce1...)
	ret = append(ret, checksum1...)
	ret = append(ret, cipher0...)
	ret = append(ret, cipher1...)

	return ret, nil
}
