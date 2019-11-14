package atk

import (
	"github.com/pekim/gobbi/lib/gdkpixbuf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), gdkpixbuf.PIXBUF_MAJOR)
}
