package atk

import (
	"github.com/pekim/gobbi/lib/soup"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, "host", soup.AuthHost)
}
