// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetObjectLocale is a wrapper around the C function atk_object_get_object_locale.
func (recv *Object) GetObjectLocale() string {
	retC := C.atk_object_get_object_locale((*C.AtkObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
