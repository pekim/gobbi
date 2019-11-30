package gdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), BUTTON_MIDDLE)
}

func TestFieldGet(t *testing.T) {
	for i := 0; i < 1000; i++ {
		color := RGBAStruct()
		ok := color.Parse("#123456")

		assert.True(t, ok)
		assert.InEpsilon(t, 0x12/float64(0x100), color.Red(), 0.005)
		assert.InEpsilon(t, 0x34/float64(0x100), color.Green(), 0.005)
		assert.InEpsilon(t, 0x56/float64(0x100), color.Blue(), 0.005)
	}
}
