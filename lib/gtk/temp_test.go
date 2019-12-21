package gtk

import (
	"github.com/pekim/gobbi/lib/internal/c/gtk"
	"testing"
)

func TestFuncCall(t *testing.T) {
	var count int
	gtk.Fn_init(&count, "")
}
