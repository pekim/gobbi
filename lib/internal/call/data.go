package call

// #include "call.h"
import "C"

type ReturnType int

const (
	RT_VOID = C.rt_void
	RT_INT  = C.rt_int
)

type Return struct {
	Type  ReturnType
	value interface{}
}

func (r *Return) Int32() int32 {
	return int32(r.value.(C.int))
}

type Data struct {
	Return Return
}
