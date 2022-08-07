package pack

import (
	"dualbox/pkg/crypt"
	"dualbox/pkg/utils"
	"fmt"
)

func Unpack(cpt crypt.Crypter, buf, key []byte) ([]byte, error) {
	header0 := buf[:HeaderSize]
	nonce0 := header0[:crypt.NonceSize]
	checksum0 := header0[crypt.NonceSize:]
	header1 := buf[HeaderSize : 2*HeaderSize]
	nonce1 := header1[:crypt.NonceSize]
	checksum1 := header1[crypt.NonceSize:]

	if dechecksum0, err := cpt.Decrypt(key, nonce0, checksum0); err == nil {
		from := utils.BytesToUint32(dechecksum0[:ChecksumFromToSize])
		to := utils.BytesToUint32(dechecksum0[+ChecksumFromToSize : +2*ChecksumFromToSize])
		plaintext, err := cpt.Decrypt(key, nonce0, buf[from:to])
		if err != nil {
			return nil, fmt.Errorf("can't decrypt cipher0: %w", err)
		}
		return plaintext, nil
	}

	if dechecksum1, err := cpt.Decrypt(key, nonce1, checksum1); err == nil {
		from := utils.BytesToUint32(dechecksum1[:ChecksumFromToSize])
		to := utils.BytesToUint32(dechecksum1[+ChecksumFromToSize : +2*ChecksumFromToSize])
		plaintext, err := cpt.Decrypt(key, nonce1, buf[from:to])
		if err != nil {
			return nil, fmt.Errorf("can't decrypt cipher1: %w", err)
		}
		return plaintext, nil
	}

	return nil, fmt.Errorf("none of the keys match")
}
