package pack

import (
	"dualbox/pkg/crypt"
	"dualbox/pkg/utils"
	"fmt"

	"github.com/sirupsen/logrus"
)

func Unpack(cpt crypt.Crypter, buf, key []byte) ([]byte, error) {
	header0 := buf[:HeaderSize]
	nonce0 := header0[:crypt.NonceSize]
	checksum0 := header0[crypt.NonceSize:]
	header1 := buf[HeaderSize : 2*HeaderSize]
	nonce1 := header1[:crypt.NonceSize]
	checksum1 := header1[crypt.NonceSize:]

	if dechecksum0, err := cpt.Decrypt(key, nonce0, checksum0); err == nil {
		if string(dechecksum0[:2]) == ChecksumMsg {
			from := utils.BytesToUint32(dechecksum0[ChecksumMsgSize : ChecksumMsgSize+ChecksumFromToSize])
			to := utils.BytesToUint32(dechecksum0[ChecksumMsgSize+ChecksumFromToSize : ChecksumMsgSize+2*ChecksumFromToSize])
			plaintext, err := cpt.Decrypt(key, nonce0, buf[from:to])
			if err != nil {
				logrus.Fatal(err)
			}
			return plaintext, nil
		}
	}

	if dechecksum1, err := cpt.Decrypt(key, nonce1, checksum1); err == nil {
		if string(dechecksum1[:2]) == ChecksumMsg {
			from := utils.BytesToUint32(dechecksum1[ChecksumMsgSize : ChecksumMsgSize+ChecksumFromToSize])
			to := utils.BytesToUint32(dechecksum1[ChecksumMsgSize+ChecksumFromToSize : ChecksumMsgSize+2*ChecksumFromToSize])
			plaintext, err := cpt.Decrypt(key, nonce1, buf[from:to])
			if err != nil {
				logrus.Fatal(err)
			}
			return plaintext, nil
		}
	}

	return nil, fmt.Errorf("nothing")
}
