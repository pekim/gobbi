package glib

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, int32(4321), BIG_ENDIAN)
}

func TestFunctionCall(t *testing.T) {
	numProc := GetNumProcessors()
	assert.True(t, numProc > 0)
}

func TestReturnedTransferOwnershipSring(t *testing.T) {
	codeset := GetCodeset()
	assert.True(t, strings.HasPrefix(codeset, "ANSI_"))
}

func TestFunctionCallWithIntegers(t *testing.T) {
	value := AsciiDigitValue('4')
	assert.Equal(t, int32(4), value)
}

func TestFunctionCallArgStringReturnString(t *testing.T) {
	result := Utf8Strdown("AaBbCc-Zz", -1)
	assert.Equal(t, "aabbcc-zz", result)
}

func TestFunctionInAndOutParams(t *testing.T) {
	value, end := AsciiStrtoll("1234xyx", 16)
	assert.Equal(t, int64(4660), value)
	assert.Equal(t, "xyz", end)
}

func TestRecordConstructor(t *testing.T) {
	date := DateNew()
	assert.NotNil(t, date.native)
}

func TestRecordMethod(t *testing.T) {
	date := DateNewJulian(100)
	date.AddDays(2)
	dayOfYear := date.GetDayOfYear()

	assert.Equal(t, uint32(102), dayOfYear)
}

func TestReturnArrayOfNullTerminatedStrings(t *testing.T) {
	assert.Equal(t, len(os.Environ()), len(GetEnviron()))
}

func BenchmarkReturnedTransferOwnershipSring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetCodeset()
	}
}
