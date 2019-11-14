// Code generated - DO NOT EDIT.

package freetype2

import (
	"fmt"
	gi "github.com/pekim/gobbi/internal/gi"
)

// UNSUPPORTED : C value 'FT_Library_Version' : non trivial function

// LibraryVersion is a representation of the C type FT_Library_Version.
func LibraryVersion() {
	invoker := gi.FunctionInvokerNew("freetype2", "library_version")
	fmt.Println(invoker)
}
