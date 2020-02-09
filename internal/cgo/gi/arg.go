package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type ArgType int

const (
	ArgType_boolean ArgType = iota
	ArgType_int8
	ArgType_uint8
	ArgType_int16
	ArgType_uint16
	ArgType_int32
	ArgType_uint32
	ArgType_int64
	ArgType_uint64
	ArgType_float
	ArgType_double
	ArgType_short
	ArgType_ushort
	ArgType_int
	ArgType_uint
	ArgType_long
	ArgType_ulong
	ArgType_ssize
	ArgType_size
	ArgType_string
	ArgType_pointer
)

type transferOwnership int

const (
	TransferOwnershipNone = iota
	TransferOwnershipFull
	TransferOwnershipContainer
)

type Arg struct {
	value             interface{}
	typ               ArgType
	pointer           bool
	array             bool
	in                bool
	out               bool
	transferOwnership transferOwnership
	outPtr            unsafe.Pointer
}

// getValue extracts the value field in to a C.GIArgument
// suitable for passing to a gi function invoker.
func (a *Arg) getValue() C.GIArgument {
	var cArg C.GIArgument

	switch a.typ {
	case ArgType_boolean:
		var cValue C.gboolean = C.FALSE
		if a.value.(bool) {
			cValue = C.TRUE
		}
		(*(*C.gboolean)(unsafe.Pointer(&cArg))) = cValue
	case ArgType_int8:
		(*(*int8)(unsafe.Pointer(&cArg))) = a.value.(int8)
	case ArgType_uint8:
		(*(*uint8)(unsafe.Pointer(&cArg))) = a.value.(uint8)
	case ArgType_int16:
		(*(*int16)(unsafe.Pointer(&cArg))) = a.value.(int16)
	case ArgType_uint16:
		(*(*uint16)(unsafe.Pointer(&cArg))) = a.value.(uint16)
	case ArgType_int32:
		(*(*int32)(unsafe.Pointer(&cArg))) = a.value.(int32)
	case ArgType_uint32:
		(*(*uint32)(unsafe.Pointer(&cArg))) = a.value.(uint32)
	case ArgType_int64:
		(*(*int64)(unsafe.Pointer(&cArg))) = a.value.(int64)
	case ArgType_uint64:
		(*(*uint64)(unsafe.Pointer(&cArg))) = a.value.(uint64)
	case ArgType_float:
		(*(*float32)(unsafe.Pointer(&cArg))) = a.value.(float32)
	case ArgType_double:
		(*(*float64)(unsafe.Pointer(&cArg))) = a.value.(float64)
	case ArgType_short:
		(*(*int16)(unsafe.Pointer(&cArg))) = a.value.(int16)
	case ArgType_ushort:
		(*(*uint16)(unsafe.Pointer(&cArg))) = a.value.(uint16)
	case ArgType_int:
		(*(*int32)(unsafe.Pointer(&cArg))) = a.value.(int32)
	case ArgType_uint:
		(*(*uint32)(unsafe.Pointer(&cArg))) = a.value.(uint32)
	case ArgType_long:
		(*(*int64)(unsafe.Pointer(&cArg))) = a.value.(int64)
	case ArgType_ulong:
		(*(*uint64)(unsafe.Pointer(&cArg))) = a.value.(uint64)
	case ArgType_ssize:
		(*(*int)(unsafe.Pointer(&cArg))) = a.value.(int)
	case ArgType_size:
		(*(*uint)(unsafe.Pointer(&cArg))) = a.value.(uint)
	case ArgType_string:
		if a.in && !a.out {
			cString := C.CString(a.value.(string))
			(*(*unsafe.Pointer)(unsafe.Pointer(&cArg))) = unsafe.Pointer(cString)
		}
		if a.out && !a.in {
			// where called function is to place output
			(*(*unsafe.Pointer)(unsafe.Pointer(&cArg))) = unsafe.Pointer(&a.outPtr)
		}
	case ArgType_pointer:
		(*(*unsafe.Pointer)(unsafe.Pointer(&cArg))) = a.value.(unsafe.Pointer)
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}

	return cArg
}

// setValue takes a C.GIArgument with a value returned from a gi
// function invocation, and sets its value in the value field.
func (a *Arg) setValue(value C.GIArgument) {
	ptrValue := unsafe.Pointer(&value)

	switch a.typ {
	case ArgType_boolean:
		a.value = (*(*C.gboolean)(ptrValue)) == C.TRUE
	case ArgType_int8:
		a.value = (int8)(*(*C.gint8)(ptrValue))
	case ArgType_uint8:
		a.value = (uint8)(*(*C.guint8)(ptrValue))
	case ArgType_int16:
		a.value = (int16)(*(*C.gint16)(ptrValue))
	case ArgType_uint16:
		a.value = (uint16)(*(*C.guint16)(ptrValue))
	case ArgType_int32:
		a.value = (int32)(*(*C.gint32)(ptrValue))
	case ArgType_uint32:
		a.value = (uint32)(*(*C.guint32)(ptrValue))
	case ArgType_int64:
		a.value = (int64)(*(*C.gint64)(ptrValue))
	case ArgType_uint64:
		a.value = (uint64)(*(*C.guint64)(ptrValue))
	case ArgType_float:
		a.value = (float32)(*(*C.gfloat)(ptrValue))
	case ArgType_double:
		a.value = (float64)(*(*C.gdouble)(ptrValue))
	case ArgType_short:
		a.value = (int16)(*(*C.gshort)(ptrValue))
	case ArgType_ushort:
		a.value = (uint16)(*(*C.gushort)(ptrValue))
	case ArgType_int:
		a.value = (int32)(*(*C.gint)(ptrValue))
	case ArgType_uint:
		a.value = (uint32)(*(*C.guint)(ptrValue))
	case ArgType_long:
		a.value = (int64)(*(*C.glong)(ptrValue))
	case ArgType_ulong:
		a.value = (uint64)(*(*C.gulong)(ptrValue))
	case ArgType_ssize:
		a.value = (int)(*(*C.gssize)(ptrValue))
	case ArgType_size:
		a.value = (uint)(*(*C.gsize)(ptrValue))
	case ArgType_string:
		var cString *C.gchar
		if a.out {
			cString = *(**C.gchar)(unsafe.Pointer(&a.outPtr))
		} else {
			cString = *(**C.gchar)(ptrValue)
		}

		a.value = C.GoString(cString)

		if a.transferOwnership != TransferOwnershipNone {
			C.free(unsafe.Pointer(cString))
		}
	case ArgType_pointer:
		a.value = unsafe.Pointer(*(*C.gpointer)(ptrValue))
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}
}
