package gi

import (
	"unsafe"
)

// #cgo pkg-config: gobject-introspection-1.0
// #include <girepository.h>
// #include <stdlib.h>
import "C"

const Qaz = 42

var defaultRepo *C.GIRepository

func init() {
	defaultRepo = C.g_irepository_get_default()
}

func Require(name string, version string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cVersion := C.CString(version)
	defer C.free(unsafe.Pointer(cVersion))

	var err *C.GError

	C.g_irepository_require(
		defaultRepo,
		cName,
		cVersion,
		0,
		&err,
	)

	if err != nil {
		message := C.GoString(err.message)
		panic(message)
	}
}
