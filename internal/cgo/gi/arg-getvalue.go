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
	var cArgPtr = unsafe.Pointer(&cArg)

	switch a.typ {
	case ArgType_boolean:
		var cValue C.gboolean = C.FALSE
		if a.value.(bool) {
			cValue = C.TRUE
		}
		(*(*C.gboolean)(cArgPtr)) = cValue
	case ArgType_int8:
		(*(*int8)(cArgPtr)) = a.value.(int8)
	case ArgType_uint8:
		(*(*uint8)(cArgPtr)) = a.value.(uint8)
	case ArgType_int16:
		(*(*int16)(cArgPtr)) = a.value.(int16)
	case ArgType_uint16:
		(*(*uint16)(cArgPtr)) = a.value.(uint16)
	case ArgType_int32:
		(*(*int32)(cArgPtr)) = a.value.(int32)
	case ArgType_uint32:
		(*(*uint32)(cArgPtr)) = a.value.(uint32)
	case ArgType_int64:
		(*(*int64)(cArgPtr)) = a.value.(int64)
	case ArgType_uint64:
		(*(*uint64)(cArgPtr)) = a.value.(uint64)
	case ArgType_float:
		(*(*float32)(cArgPtr)) = a.value.(float32)
	case ArgType_double:
		(*(*float64)(cArgPtr)) = a.value.(float64)
	case ArgType_short:
		(*(*int16)(cArgPtr)) = a.value.(int16)
	case ArgType_ushort:
		(*(*uint16)(cArgPtr)) = a.value.(uint16)
	case ArgType_int:
		(*(*int)(cArgPtr)) = a.value.(int)
	case ArgType_uint:
		(*(*uint)(cArgPtr)) = a.value.(uint)
	case ArgType_long:
		(*(*int64)(cArgPtr)) = a.value.(int64)
	case ArgType_ulong:
		(*(*uint64)(cArgPtr)) = a.value.(uint64)
	case ArgType_ssize:
		(*(*int)(cArgPtr)) = a.value.(int)
	case ArgType_size:
		(*(*uint)(cArgPtr)) = a.value.(uint)
	case ArgType_string:
		if a.array {
			a.getValueStringArray(cArgPtr)
		} else {
			a.getValueString(cArgPtr)
		}
	case ArgType_pointer:
		(*(*unsafe.Pointer)(cArgPtr)) = a.value.(unsafe.Pointer)
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}

	return cArg
}

func (a *Arg) getValueString(cArgPtr unsafe.Pointer) {
	if a.in && !a.out {
		cString := C.CString(a.value.(string))
		(*(*unsafe.Pointer)(cArgPtr)) = unsafe.Pointer(cString)
	}
	if a.out && !a.in {
		// where called function is to place output
		(*(*unsafe.Pointer)(cArgPtr)) = unsafe.Pointer(&a.outPtr)
	}
}

func (a *Arg) getValueStringArray(cArgPtr unsafe.Pointer) {
	if a.arrayNullTerminated {
		a.getValueStringArrayNullTerminated(cArgPtr)
	} else {
		panic("not supported : string array with length")
	}
}

func (a *Arg) getValueStringArrayNullTerminated(cArgPtr unsafe.Pointer) {
	strings := a.value.([]string)

	if strings == nil || len(strings) == 0 {
		(*(*unsafe.Pointer)(cArgPtr)) = nil
		return
	}

	count := len(strings) + 1
	arraySize := count * C.sizeof_gpointer
	cArray := (**C.gchar)(C.malloc(C.ulong(arraySize)))
	cStrings := (*[1 << 28]*C.char)(unsafe.Pointer(cArray))[:count:count]

	for i, str := range strings {
		cStrings[i] = C.CString(str)
	}

	// null terminator
	cStrings[count-1] = nil

	(*(*unsafe.Pointer)(cArgPtr)) = unsafe.Pointer(cArray)
}
