package glib

import (
	"github.com/pekim/gobbi/lib/glib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(4321), glib.BIG_ENDIAN)
}

func TestFunctionCall(t *testing.T) {
	glib.ClearError()
}
