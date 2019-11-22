package atk

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), MAJOR_VERSION)

	version, err := GetVersion()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(version, "2."))
}
