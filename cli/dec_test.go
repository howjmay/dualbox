package cli

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dec(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	filePath := path + "/test.enc"
	err := dec(filePath, "f79681852aad0428a5005a4dc0c25404bb3c3c2b387410a53cf6253e09e416db")
	require.NoError(t, err)
}
