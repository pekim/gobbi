// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var libraryVersionFunction *gi.Function
var libraryVersionFunction_Once sync.Once

func libraryVersionFunction_Set() {
	libraryVersionFunction_Once.Do(func() {
		libraryVersionFunction = gi.FunctionInvokerNew("freetype2", "library_version")
	})
}

var libraryVersionInvoker *gi.Function

// LibraryVersion is a representation of the C type FT_Library_Version.
func LibraryVersion() {
	libraryVersionFunction_Set()

	libraryVersionFunction.Invoke(nil, nil)

}
