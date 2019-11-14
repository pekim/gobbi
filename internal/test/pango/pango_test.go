package atk

import (
	"github.com/pekim/gobbi/lib/pango"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), pango.VERSION_MIN_REQUIRED)
}
