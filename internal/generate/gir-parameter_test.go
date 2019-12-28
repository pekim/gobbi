package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
	"go/format"
	"strings"
	"testing"
)

func TestParameterSysParamGoType(t *testing.T) {
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
		param    Parameter
		expected string
	}{
		{
			name: "array of simple type",
			param: Parameter{
				Direction: "",
				Argument: Argument{
					Type: nil,
					Array: &Array{
						CType: "gint*",
						Type:  &Type{Name: "gint", CType: "gint"},
					},
				},
			},
			expected: "var array []int",
		},
		{
			name: "out array of simple type",
			param: Parameter{
				Direction: "out",
				Argument: Argument{
					Type: nil,
					Array: &Array{
						CType: "gint*",
						Type:  &Type{Name: "gint", CType: "gint"},
					},
				},
			},
			expected: "var array *[]int",
		},
		{
			name: "array of string",
			param: Parameter{
				Direction: "",
				Argument: Argument{
					Type: nil,
					Array: &Array{
						CType: "const gchar**",
						Type:  &Type{Name: "utf8", CType: "const gchar*"},
					},
				},
			},
			expected: "var array []string",
		},
		{
			name: "array of record",
			param: Parameter{
				Direction: "",
				Argument: Argument{
					Type: nil,
					Array: &Array{
						CType: "GtkIconSize**",
						Type:  &Type{Name: "IconSize", CType: "GtkIconSize*"},
					},
				},
			},
			expected: "var array []unsafe.Pointer",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.param.init(ns)

			// generate
			f := jen.NewFile("test")
			f.Var().Id("array").Add(test.param.sysParamGoType())
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
