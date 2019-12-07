package pango

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), VERSION_MIN_REQUIRED)

	context := ContextNew()
	assert.NotNil(t, context)
}
