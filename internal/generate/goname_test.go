package generate

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGoName(t *testing.T) {
	for i, test := range []struct {
		cName              string
		expectedUnexported string
		expectedExported   string
	}{
		{"abc_def", "abcDef", "AbcDef"},
		{"abcDef", "abcDef", "AbcDef"},
		{"Abc_Def", "abcDef", "AbcDef"},
		{"Abc_def", "abcDef", "AbcDef"},
		{"AbcDef", "abcDef", "AbcDef"},
		{"_abc_def", "abcDef", "AbcDef"},
		{"__abc_def", "abcDef", "AbcDef"},
		{"aBCDef", "aBCDef", "ABCDef"},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, test.expectedUnexported, makeUnexportedGoName(test.cName))
			assert.Equal(t, test.expectedExported, makeExportedGoName(test.cName))
		})
	}
}
