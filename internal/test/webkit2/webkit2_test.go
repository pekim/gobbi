package webkit2

import (
	"github.com/pekim/gobbi/lib/webkit2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), webkit2.MAJOR_VERSION)
}
