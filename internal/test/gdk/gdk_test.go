package gdk

import (
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), gdk.BUTTON_MIDDLE)
}
