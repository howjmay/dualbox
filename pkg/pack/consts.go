package pack

import "dualbox/pkg/crypt"

const HeaderSize = crypt.NonceSize + CheckSumEncryptedSize
const CheckSumEncryptedSize = 24
const CheckSumSize = 8
const ChecksumFromToSize = 4
