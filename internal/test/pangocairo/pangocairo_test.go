package pangocairo

import (
	"github.com/pekim/gobbi/lib/pangocairo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	fm := pangocairo.FontMapNew()
	assert.NotNil(t, fm)
}
