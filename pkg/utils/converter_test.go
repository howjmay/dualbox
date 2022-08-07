package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Uint32AndBytes(t *testing.T) {
	testNum := uint32(123)
	b := Uint32ToBytes(testNum)
	res := BytesToUint32(b)
	require.Equal(t, testNum, res)

}
