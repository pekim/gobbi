package gobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, SignalFlags(32), SIGNAL_ACTION)
}
