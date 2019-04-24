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
	version := int(pango.Version())
	major := version / 10000
	minor := (version - major) / 100

	assert.Equal(t, 1, major)
	assert.Equal(t, 0, minor%2) // even -> stable, odd -> unstable
}
