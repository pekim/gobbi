package gtkx

import (
	"github.com/pekim/gobbi/lib/gtkx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	argsIn := []string{"one", "--g-fatal-warnings", "two", "--class", "cls", "three"}

	argsOut := gtkx.Init(argsIn)
	assert.Equal(t, []string{"one", "two", "three"}, argsOut)
}
