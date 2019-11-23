package gobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(255), PARAM_MASK)
}

func TestBadFunction(t *testing.T) {
	err := Bad()
	assert.NotNil(t, err)
}
