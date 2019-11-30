package gdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(2), BUTTON_MIDDLE)
}

func TestFieldGet(t *testing.T) {
	epsilon := 0.004

	color := RGBAStruct()
	ok := color.Parse("#123456")

	assert.True(t, ok)
	assert.InEpsilon(t, 0x12/float64(0x100), color.FieldRed(), epsilon)
	assert.InEpsilon(t, 0x34/float64(0x100), color.FieldGreen(), epsilon)
	assert.InEpsilon(t, 0x56/float64(0x100), color.FieldBlue(), epsilon)
}

func TestFieldSet(t *testing.T) {
	epsilon := 0.004
	color := RGBAStruct()

	ok := color.Parse("#123456")
	assert.True(t, ok)
	assert.InEpsilon(t, 0x12/float64(0x100), color.FieldRed(), epsilon)

	color.SetFieldRed(0xcc / float64(0x100))
	assert.InEpsilon(t, 0xcc/float64(0x100), color.FieldRed(), epsilon)
}
