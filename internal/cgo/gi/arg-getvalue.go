package gi

// #include <girepository.h>
import "C"

import (
	"fmt"
	"unsafe"
)

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
