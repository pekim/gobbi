package glib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, AsciiType(16), ASCII_GRAPH)
}

func TestFunctionCallWithIntegers(t *testing.T) {
	assert.Equal(t, int32(4), AsciiDigitValue('4'))
}

func TestFunctionCallWithStringsNotFreeReturnedString(t *testing.T) {
	assert.Equal(t, "edcba", Strreverse("abcde"))
}

func TestFunctionCallWithReturnedStringToFree(t *testing.T) {
	assert.Equal(t, "ABCDE", AsciiStrup("aBcDe", -1))
}

func TestFunctionCallWithAliasParamAndReturn(t *testing.T) {
	s := "abc"
	q := QuarkFromString(s)
	assert.Equal(t, s, QuarkToString(q))
}
