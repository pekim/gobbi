// This is a generated file - DO NOT EDIT

package fontconfig

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <stdlib.h>
import "C"

// Init is a wrapper around the C function FcInit.
func Init() {
	C.FcInit()

	return
}
