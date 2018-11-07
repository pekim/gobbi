package gdkpixbuf

import (
	"github.com/pekim/gobbi/lib/gdkpixbuf"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	gdkpixbuf.PixbufNewFromFile("badfilename")
}
