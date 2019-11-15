package atk

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), MAJOR_VERSION)

	assert.True(t, strings.HasPrefix(GetVersion(), "2."))
}
