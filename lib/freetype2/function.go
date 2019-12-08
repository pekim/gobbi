// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	"sync"
)

var libraryVersionFunction *gi.Function
var libraryVersionFunction_Once sync.Once

func libraryVersionFunction_Set() error {
	var err error
	libraryVersionFunction_Once.Do(func() {
		libraryVersionFunction, err = gi.FunctionInvokerNew("freetype2", "library_version")
	})
	return err
}

// LibraryVersion is a representation of the C type FT_Library_Version.
func LibraryVersion() {

	err := libraryVersionFunction_Set()
	if err == nil {
		libraryVersionFunction.Invoke(nil, nil)
	}

	return
}
