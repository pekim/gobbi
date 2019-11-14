package atk

import (
	"github.com/pekim/gobbi/lib/atk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), atk.MAJOR_VERSION)
}
