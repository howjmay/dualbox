package utils

func Uint32ToBytes(input uint32) []byte {
	buf := make([]byte, 4)
	buf[0] = byte((input >> (0 * 8)) & 0xff)
	buf[1] = byte((input >> (1 * 8)) & 0xff)
	buf[2] = byte((input >> (2 * 8)) & 0xff)
	buf[3] = byte((input >> (3 * 8)) & 0xff)
	return buf
}

func BytesToUint32(input []byte) uint32 {
	return (uint32(input[0]) << (0 * 8)) + (uint32(input[1]) << (1 * 8)) + (uint32(input[2]) << (2 * 8)) + (uint32(input[3]) << (3 * 8))
}
