// This is a generated file - DO NOT EDIT
// +build atk_1.20 atk_2.8

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetVersion is a wrapper around the C function atk_get_version.
func GetVersion() string {
	retC := C.atk_get_version()
	retGo := C.GoString(retC)

	return retGo
}
