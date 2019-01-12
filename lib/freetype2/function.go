// This is a generated file - DO NOT EDIT

package freetype2

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <stdlib.h>
import "C"

// LibraryVersion is a wrapper around the C function FT_Library_Version.
func LibraryVersion() {
	C.FT_Library_Version()

	return
}
