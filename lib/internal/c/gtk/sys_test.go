package gtk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilds(t *testing.T) {
}

func TestStringArrayInout(t *testing.T) {
	args := []string{"qaz", "--g-fatal-warnings", "qwerty"}
	var count int = len(args)
	Fn_gtk_init(&count, &args)

	assert.Equal(t, 2, count)
	assert.Equal(t, []string{"qaz", "qwerty"}, args)
}
