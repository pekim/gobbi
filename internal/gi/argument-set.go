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

func (a *Argument) SetString(value string) {
	cString := C.CString(value)
	*(**C.gchar)(unsafe.Pointer(a)) = cString
}

func (a *Argument) SetPointer(value uintptr) {
	(*(*uintptr)(unsafe.Pointer(a))) = value
}
