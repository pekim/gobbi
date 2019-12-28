package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
	"go/format"
	"strings"
	"testing"
)

func TestTypeSysParamGoPlainType(t *testing.T) {
	ns := &Namespace{
		Name: "GLib",
		Records: Records{
			&Record{
				Name: "IconSize",
			},
		},
	}
	ns.init(nil, namespaces{"GLib": ns})

	tests := []struct {
		name     string
		typ      *Type
		expected string
	}{
		{
			name:     "simple type",
			typ:      &Type{Name: "gint", CType: "gint"},
			expected: "var v int",
		},
		{
			name:     "string",
			typ:      &Type{Name: "utf8", CType: "const gchar*"},
			expected: "var v string",
		},
		{
			name:     "record",
			typ:      &Type{Name: "IconSize", CType: "GtkIconSize"},
			expected: "var v IconSize",
		},
		{
			name:     "record pointer",
			typ:      &Type{Name: "IconSize", CType: "GtkIconSize*"},
			expected: "var v unsafe.Pointer",
		},
		//{
		//	name:     "void *",
		//	typ:      &Type{Name: "none", CType: "void"},
		//	expected: "var v ",
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.typ.init(ns)

			// generate
			f := jen.NewFile("test")
			f.Var().Id("v").Add(test.typ.sysParamGoPlainType())
			generated := fmt.Sprintf("%#v\n", f)

			// take the last non-empty line, with the var declaration
			lines := strings.Split(generated, "\n")
			actual := lines[len(lines)-3]

			// format the expected value, to ensure consistency
			formattedExpected, err := format.Source([]byte(test.expected))
			if err != nil {
				t.Fatalf("Expected value could not be formatted -- %v", err)
			}
			expected := string(formattedExpected)

			assert.Equal(t, expected, actual)
		})
	}
}
