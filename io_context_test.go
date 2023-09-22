package astiav_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/491467928/go-astiav"
	"github.com/stretchr/testify/require"
)

func TestIOContext(t *testing.T) {
	c := astiav.NewIOContext()
	path := filepath.Join(t.TempDir(), "iocontext.txt")
	err := c.Open(path, astiav.NewIOContextFlags(astiav.IOContextFlagWrite))
	require.NoError(t, err)
	c.Write(nil)
	c.Write([]byte("test"))
	err = c.Closep()
	require.NoError(t, err)
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "test", string(b))
	err = os.Remove(path)
	require.NoError(t, err)
}
