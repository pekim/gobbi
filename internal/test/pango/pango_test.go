package pango

import (
	"github.com/pekim/gobbi/lib/pango"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	context := pango.ContextNew()
	assert.NotNil(t, context)
}
