// This is a generated file - DO NOT EDIT
// +build atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// MiscGetInstance is a wrapper around the C function atk_misc_get_instance.
func MiscGetInstance() *Misc {
	retC := C.atk_misc_get_instance()
	retGo := MiscNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ThreadsEnter is a wrapper around the C function atk_misc_threads_enter.
func (recv *Misc) ThreadsEnter() {
	C.atk_misc_threads_enter((*C.AtkMisc)(recv.native))

	return
}

// ThreadsLeave is a wrapper around the C function atk_misc_threads_leave.
func (recv *Misc) ThreadsLeave() {
	C.atk_misc_threads_leave((*C.AtkMisc)(recv.native))

	return
}
