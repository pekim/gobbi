package gobject

import (
	"github.com/pekim/gobbi"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(255), PARAM_MASK)
}

func TestBadFunction(t *testing.T) {
	handlerCalledCount := 0
	gobbi.SetErrorHandler(func(err error) {
		assert.True(t, strings.Contains(err.Error(), "function"))
		assert.True(t, strings.Contains(err.Error(), "bad"))

		handlerCalledCount++
	})

	// good
	SignalName(0)
	assert.Equal(t, 0, handlerCalledCount)

	// bad
	Bad()
	assert.Equal(t, 1, handlerCalledCount)
}