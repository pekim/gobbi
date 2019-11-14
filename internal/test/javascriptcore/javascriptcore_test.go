package javascriptcore

import (
	"github.com/pekim/gobbi/lib/javascriptcore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), javascriptcore.MAJOR_VERSION)
}
