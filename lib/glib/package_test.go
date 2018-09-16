package glib

import (
	"strings"
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
	assert.Equal(t, "abcde", InternString("abcde"))
}

func TestFunctionCallWithReturnedStringToFree(t *testing.T) {
	assert.Equal(t, "ABCDE", AsciiStrup("aBcDe", -1))
}

func TestFunctionCallWithAliasParamAndReturn(t *testing.T) {
	s := "abc"
	q := QuarkFromString(s)
	assert.Equal(t, s, QuarkToString(q))
}

func TestFunctionReturningError(t *testing.T) {
	_, err := FileOpenTmp("bad/should/not/contain/a/slash")
	assert.NotNil(t, err)

	glibError := err.(*Error)
	assert.True(t,
		strings.Contains(glibError.Message, "should not contain"),
		"error message not correct")
}

func TestFunctionReturningNoError(t *testing.T) {
	_, err := FilenameFromUri("file:///good")
	assert.Nil(t, err)
}
