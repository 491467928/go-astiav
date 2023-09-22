package astiav_test

import (
	"testing"

	"github.com/491467928/go-astiav"
	"github.com/stretchr/testify/require"
)

func TestMediaType(t *testing.T) {
	require.Equal(t, "video", astiav.MediaTypeVideo.String())
}
