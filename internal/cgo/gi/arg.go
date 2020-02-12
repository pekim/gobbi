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

/*
Arg holds an argument's value and type.

An Arg's value may be extracted with getValue method, and placed in a C.GIArgument.
Its value may set from a C.GIArgument with the setValue method.

When the Arg's value is represented in a C.GIArgument it may do so in
several different forms depending on the type of the value and various flags.

A number type.
	Go : int
	C  : int

The number value is simply set in the GIArgument.

	GIArgument 	= value


A pointer to a number type.
Typically for an output parameter
	Go : *int
	C  : *int
Or can be an optional input parameter, where GIArgument could hold nil.
	Go : *int
	C  : int*

	GIArgument	= pointer to arg.value
	arg.ptr		= value


A string.
	Go : string
	C  : gchar*

	GIArgument	= pointer to a null terminated C string


An output string.
	Go : *string
	C  : gchar**
The ptr field is used to provide somewhere for the function
to place the pointer to the null terminated string.

	GIArgument	= pointer to arg.ptr


An array of numbers
	Go : []int
	C  : int*

	GIArgument 	= pointer to arg.value[0]

*/
type Arg struct {
	value interface{}
	typ   ArgType

	// Qualifiying attributes of an argument's type.
	pointer bool
	array   bool
	// The index of the arg with array's length.
	// applicable if arrayNullTerminated is false
	arrayLengthArg int
	arrayLength    int
	// array is null-terminated
	arrayNullTerminated bool
	in                  bool
	out                 bool
	transferOwnership   transferOwnership

	valuePtr unsafe.Pointer
	// A place for out arguments to leave a pointer.
	outPtr unsafe.Pointer
}
