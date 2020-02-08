package gi

// #include <girepository.h>
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
}

func (a *Arg) getValue() C.GIArgument {
	var cArg C.GIArgument

	switch a.typ {
	case ArgType_size:
		(*(*uint64)(unsafe.Pointer(&cArg))) = a.value.(uint64)
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}

	return cArg
}

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
		a.value = (int)(*(*C.gint)(ptrValue))
	case ArgType_uint:
		a.value = (uint)(*(*C.guint)(ptrValue))
	case ArgType_long:
		a.value = (int64)(*(*C.glong)(ptrValue))
	case ArgType_ulong:
		a.value = (uint64)(*(*C.gulong)(ptrValue))
	case ArgType_ssize:
		a.value = (int)(*(*C.gssize)(ptrValue))
	case ArgType_size:
		a.value = (uint)(*(*C.gsize)(ptrValue))
	case ArgType_pointer:
		a.value = unsafe.Pointer(*(*C.gpointer)(ptrValue))
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}
}
