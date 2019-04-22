package pango

import (
	"github.com/pekim/gobbi/lib/pango"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestCleanBuild(t *testing.T) {
//	context := pango.ContextNew()
//	assert.NotNil(t, context)
//}

func TestVersion(t *testing.T) {
	assert.Equal(t, int32(13801), pango.Version())
}
