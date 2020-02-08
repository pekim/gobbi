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
	switch a.typ {
	case ArgType_boolean:
		a.value = (*(*C.gboolean)(unsafe.Pointer(&value))) == C.TRUE
	case ArgType_uint:
		a.value = (uint)(*(*C.guint)(unsafe.Pointer(&value)))
	case ArgType_pointer:
		a.value = unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&value)))
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}
}
