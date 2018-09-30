package gobject

import (
	"github.com/pekim/gobbi/lib/gobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, gobject.SignalFlags(32), gobject.SIGNAL_ACTION)
}
