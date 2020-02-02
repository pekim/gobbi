package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func (a Argument) Int8() int8 {
	return (int8)(*(*C.gint8)(unsafe.Pointer(&a)))
}

func (a Argument) Int16() int16 {
	return (int16)(*(*C.gint16)(unsafe.Pointer(&a)))
}

func (a Argument) Int32() int32 {
	return (int32)(*(*C.gint32)(unsafe.Pointer(&a)))
}

func (a Argument) Int64() int64 {
	return (int64)(*(*C.gint64)(unsafe.Pointer(&a)))
}

func (a Argument) Uint8() uint8 {
	return (uint8)(*(*C.guint8)(unsafe.Pointer(&a)))
}

func (a Argument) Uint16() uint16 {
	return (uint16)(*(*C.guint16)(unsafe.Pointer(&a)))
}

func (a Argument) Uint32() uint32 {
	return (uint32)(*(*C.guint32)(unsafe.Pointer(&a)))
}

func (a Argument) Uint64() uint64 {
	return (uint64)(*(*C.guint64)(unsafe.Pointer(&a)))
}

func (a Argument) Float32() float32 {
	return (float32)(*(*C.gfloat)(unsafe.Pointer(&a)))
}

func (a Argument) Float64() float64 {
	return (float64)(*(*C.gdouble)(unsafe.Pointer(&a)))
}

func (a Argument) Boolean() bool {
	value := *(*C.gboolean)(unsafe.Pointer(&a))
	return value == C.TRUE
}

func (a Argument) String(transferOwnership bool) string {
	cString := *(**C.gchar)(unsafe.Pointer(&a))
	if transferOwnership {
		defer C.free(unsafe.Pointer(cString))
	}

	goString := C.GoString(cString)
	return goString
}

func (a Argument) StringArray(transferOwnership bool) []string {
	var strings []string

	cArray := *(**C.gchar)(unsafe.Pointer(&a))
	for cArray != nil {
		string := C.GoString(cArray)
		strings = append(strings, string)

		// advance to next string
		cArray = (*C.gchar)(incptr(unsafe.Pointer(cArray), C.sizeof_gpointer))
	}

	return strings
}

func (a Argument) Pointer() unsafe.Pointer {
	return unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&a)))
}
