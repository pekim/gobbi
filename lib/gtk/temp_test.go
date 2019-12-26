package gtk

import (
	"fmt"
	"github.com/pekim/gobbi/lib/internal/c/gtk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringArrayInoutEmpty(t *testing.T) {
	args := []string{}
	var count int = len(args)

	gtk.Fn_gtk_init(&count, &args)
	assert.Equal(t, 0, count)
}

func TestStringArrayInout(t *testing.T) {
	args := []string{"qaz", "--g-fatal-warnings", "qwerty"}
	var count int = len(args)

	gtk.Fn_gtk_init(&count, &args)
	fmt.Println(args)
	assert.Equal(t, 2, count)
	assert.Equal(t, []string{"qaz", "qwerty"}, args)
}
