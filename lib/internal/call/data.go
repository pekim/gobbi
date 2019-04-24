package call

// #include "call.h"
import "C"

type ReturnType int

const (
	RT_VOID = C.rt_void
	RT_INT  = C.rt_int
)

type Return struct {
	Type ReturnType
	Int  int
}

type Data struct {
	Return Return
}

//func (d*Data)
