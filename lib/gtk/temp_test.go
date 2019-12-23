package gtk

import (
	"github.com/pekim/gobbi/lib/internal/c/gtk"
	"testing"
)

func TestFuncCall(t *testing.T) {
	var count int
	var args []string
	gtk.Fn_gtk_init(&count, &args)
}
