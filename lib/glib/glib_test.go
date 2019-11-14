package glib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(4321), BIG_ENDIAN)
}

func TestFunctionCall(t *testing.T) {
	ClearError()
}
