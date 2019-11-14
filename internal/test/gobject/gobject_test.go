package gobject

import (
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(255), gobject.PARAM_MASK)
}
