package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
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

// Arg hold an argument's value and type.
type Arg struct {
	value interface{}
	typ   ArgType

	// Qualifiying attributes of an argument's type.
	pointer bool
	array   bool
	// The index of the arg with array's length.
	// applicable if arrayNullTerminated is false
	arrayLengthArg int
	// array is null-terminated
	arrayNullTerminated bool
	in                  bool
	out                 bool
	transferOwnership   transferOwnership

	// A place for out arguments to leave a pointer.
	outPtr unsafe.Pointer
}
