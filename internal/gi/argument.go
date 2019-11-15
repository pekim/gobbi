package gi

// #include <girepository.h>
import "C"
import "unsafe"

type Argument C.GIArgument

func (a Argument) Int8() int8 {
	return (int8)(*(*C.gint8)(unsafe.Pointer(&a)))
}

func (a Argument) Int16() int16 {
	return (int16)(*(*C.gint16)(unsafe.Pointer(&a)))
}

func (a Argument) Int32() int32 {
	return (int32)(*(*C.gint)(unsafe.Pointer(&a)))
}

func (a Argument) Uint32() uint32 {
	return (uint32)(*(*C.guint)(unsafe.Pointer(&a)))
}
