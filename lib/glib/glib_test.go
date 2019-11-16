package glib

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(4321), BIG_ENDIAN)
}

func TestFunctionCall(t *testing.T) {
	assert.True(t, GetNumProcessors() > 0)
}

func TestReturnedTransferOwnershipSring(t *testing.T) {
	assert.True(t, strings.HasPrefix(GetCodeset(), "ANSI_"))
}

func TestFunctionCallWithIntegers(t *testing.T) {
	assert.Equal(t, int32(4), AsciiDigitValue('4'))
}

func BenchmarkReturnedTransferOwnershipSring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetCodeset()
	}
}
