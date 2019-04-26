package call

// #include "call.h"
import "C"

type Type int

const (
	TYPE_VOID   Type = C.type_void
	TYPE_INT         = C.type_int
	TYPE_UINT        = C.type_uint
	TYPE_LONG        = C.type_long
	TYPE_DOUBLE      = C.type_double
)

type Value struct {
	Type  Type
	value interface{}
}

func (v *Value) Int32() int32 {
	return int32(v.value.(C.int))
}

func (v *Value) Uint32() uint32 {
	return uint32(v.value.(C.uint))
}

func (v *Value) Int64() int64 {
	return int64(v.value.(C.long))
}

func (v *Value) Float64() float64 {
	return float64(v.value.(C.double))
}

type Data struct {
	Return Value
	Params []Value
}
