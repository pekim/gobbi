package call

// #include "call.h"
import "C"

type ReturnType int

const (
	RT_VOID   = C.rt_void
	RT_INT    = C.rt_int
	RT_UINT   = C.rt_uint
	RT_LONG   = C.rt_long
	RT_DOUBLE = C.rt_double
)

type Return struct {
	Type  ReturnType
	value interface{}
}

func (r *Return) Int32() int32 {
	return int32(r.value.(C.int))
}

func (r *Return) Uint32() uint32 {
	return uint32(r.value.(C.uint))
}

func (r *Return) Int64() int64 {
	return int64(r.value.(C.long))
}

func (r *Return) Float64() float64 {
	return float64(r.value.(C.double))
}

type Data struct {
	Return Return
}
