// This is a generated file - DO NOT EDIT
// +build atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetVersion is a wrapper around the C function atk_get_version.
func GetVersion() string {
	retC := C.atk_get_version()
	retGo := C.GoString(retC)

	return retGo
}
