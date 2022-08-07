package cli

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_enc(t *testing.T) {
	key0 := "f79681852aad0428a5005a4dc0c25404bb3c3c2b387410a53cf6253e09e416db"
	key1 := "3aaf57cdd0fc902048388f7ebf9fe4f175ca3182c23109734e97323b5d719ab7"
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	filePath0 := path + "/../testdata/testdata0.jpg"
	filePath1 := path + "/../testdata/testdata1.png"
	err := enc([]string{filePath0, filePath1}, []string{key0, key1}, "test.enc")
	require.NoError(t, err)
}
