package astiav_test

import (
	"testing"

	"github.com/491467928/go-astiav"
	"github.com/stretchr/testify/require"
)

func TestCodecID(t *testing.T) {
	require.Equal(t, astiav.MediaTypeVideo, astiav.CodecIDH264.MediaType())
	require.Equal(t, "h264", astiav.CodecIDH264.Name())
	require.Equal(t, "h264", astiav.CodecIDH264.String())
}
