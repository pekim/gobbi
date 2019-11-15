// Code generated - DO NOT EDIT.

package freetype2

import gi "github.com/pekim/gobbi/internal/gi"

var libraryVersionInvoker *gi.Function

// LibraryVersion is a representation of the C type FT_Library_Version.
func LibraryVersion() {
	if libraryVersionInvoker == nil {
		libraryVersionInvoker = gi.FunctionInvokerNew("freetype2", "library_version")
	}

	libraryVersionInvoker.Invoke()
}
