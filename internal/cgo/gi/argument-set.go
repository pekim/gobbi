package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func (a *Argument) SetInt8(value int8) {
	(*(*int8)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetInt16(value int16) {
	(*(*int16)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetInt32(value int32) {
	(*(*int32)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetInt64(value int64) {
	(*(*int64)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetUint8(value uint8) {
	(*(*uint8)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetUint16(value uint16) {
	(*(*uint16)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetUint32(value uint32) {
	(*(*uint32)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetUint64(value uint64) {
	(*(*uint64)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetFloat32(value float32) {
	(*(*float32)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetFloat64(value float64) {
	(*(*float64)(unsafe.Pointer(a))) = value
}

func (a *Argument) SetBoolean(value bool) {
	var cValue C.gboolean = C.FALSE
	if value {
		cValue = C.TRUE
	}
	(*(*C.gboolean)(unsafe.Pointer(a))) = cValue
}

func (a *Argument) SetString(value string) {
	cString := C.CString(value)
	*(**C.gchar)(unsafe.Pointer(a)) = cString
}

func (a *Argument) SetStringArray(value []string) {
	if value == nil || len(value) == 0 {
		*(***C.gchar)(unsafe.Pointer(&a)) = nil
		return
	}

	//cArray := make([](*C.gchar), len(value), len(value))
	//for i, str := range value {
	//	cArray[i] = C.CString(str)
	//}
	//
	//*(***C.gchar)(unsafe.Pointer(a)) = &cArray[0]

	arraySize := len(value) * C.sizeof_gpointer
	cArray := (**C.gchar)(C.malloc(C.ulong(arraySize)))
	cStringPtr := cArray
	for _, str := range value {
		*cStringPtr = C.CString(str)

		// advance to next
		cStringPtr = (**C.gchar)(incrPointer(unsafe.Pointer(cStringPtr), C.sizeof_gpointer))
	}

	//*(***C.gchar)(unsafe.Pointer(a)) = cArray
	a.SetPointer(unsafe.Pointer(cArray))
}

func (a *Argument) SetPointer(value unsafe.Pointer) {
	(*(*uintptr)(unsafe.Pointer(a))) = uintptr(value)
}
