package ffi

// #include <ffi.h>
import "C"

type Type C.ffi_type

var (
	Type_void       = Type(C.ffi_type_void)
	Type_uint8      = Type(C.ffi_type_uint8)
	Type_sint8      = Type(C.ffi_type_sint8)
	Type_uint16     = Type(C.ffi_type_uint16)
	Type_sint16     = Type(C.ffi_type_sint16)
	Type_uint32     = Type(C.ffi_type_uint32)
	Type_sint32     = Type(C.ffi_type_sint32)
	Type_uint64     = Type(C.ffi_type_uint64)
	Type_sint64     = Type(C.ffi_type_sint64)
	Type_float      = Type(C.ffi_type_float)
	Type_double     = Type(C.ffi_type_double)
	Type_uchar      = Type(C.ffi_type_uchar)
	Type_schar      = Type(C.ffi_type_schar)
	Type_ushort     = Type(C.ffi_type_ushort)
	Type_sshort     = Type(C.ffi_type_sshort)
	Type_uint       = Type(C.ffi_type_uint)
	Type_sint       = Type(C.ffi_type_sint)
	Type_ulong      = Type(C.ffi_type_ulong)
	Type_slong      = Type(C.ffi_type_slong)
	Type_longdouble = Type(C.ffi_type_longdouble)
	Type_pointer    = Type(C.ffi_type_pointer)
)
