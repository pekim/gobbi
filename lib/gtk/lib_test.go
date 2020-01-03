package gtk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionReturningInt(t *testing.T) {
	assert.Equal(t, uint(3), GetMajorVersion())
}
