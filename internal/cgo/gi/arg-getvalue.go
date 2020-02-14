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
	case ArgType_boolean,
		ArgType_int8, ArgType_uint8, ArgType_int16, ArgType_uint16,
		ArgType_int32, ArgType_uint32, ArgType_int64, ArgType_uint64,
		ArgType_float, ArgType_double,
		ArgType_short, ArgType_ushort, ArgType_int, ArgType_uint, ArgType_long, ArgType_ulong,
		ArgType_ssize, ArgType_size,
		ArgType_pointer:
		if a.array {
			a.getValueSimpleArray(cArgPtr)
		} else {
			a.getValueSimple(cArgPtr)
		}
	case ArgType_string:
		if a.array {
			a.getValueStringArray(cArgPtr)
		} else {
			a.getValueString(cArgPtr)
		}
	default:
		panic(fmt.Sprintf("Unhandled arg type, %#v", a))
	}

	return cArg
}

func (a *Arg) getValueSimple(cArgPtr unsafe.Pointer) {
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
	case ArgType_pointer:
		(*(*unsafe.Pointer)(cArgPtr)) = a.value.(unsafe.Pointer)
	}
}

func (a *Arg) getValueSimpleArray(cArgPtr unsafe.Pointer) {
	if a.arrayNullTerminated {
		panic("not supported : simple null-terminated array")
	}

	var array unsafe.Pointer
	switch a.typ {
	case ArgType_boolean:
		array = a.getValueBoolArray()
	case ArgType_int8:
		array = unsafe.Pointer(&(a.value.([]int8))[0])
	case ArgType_uint8:
		array = unsafe.Pointer(&(a.value.([]uint8))[0])
	case ArgType_int16:
		array = unsafe.Pointer(&(a.value.([]int16))[0])
	case ArgType_uint16:
		array = unsafe.Pointer(&(a.value.([]uint16))[0])
	case ArgType_int32:
		array = unsafe.Pointer(&(a.value.([]int32))[0])
	case ArgType_uint32:
		array = unsafe.Pointer(&(a.value.([]uint32))[0])
	case ArgType_int64:
		array = unsafe.Pointer(&(a.value.([]int64))[0])
	case ArgType_uint64:
		array = unsafe.Pointer(&(a.value.([]uint64))[0])
	case ArgType_float:
		array = unsafe.Pointer(&(a.value.([]float32))[0])
	case ArgType_double:
		array = unsafe.Pointer(&(a.value.([]float64))[0])
	case ArgType_short:
		array = unsafe.Pointer(&(a.value.([]int16))[0])
	case ArgType_ushort:
		array = unsafe.Pointer(&(a.value.([]uint16))[0])
	case ArgType_int:
		array = unsafe.Pointer(&(a.value.([]int))[0])
	case ArgType_uint:
		array = unsafe.Pointer(&(a.value.([]uint))[0])
	case ArgType_long:
		array = unsafe.Pointer(&(a.value.([]int64))[0])
	case ArgType_ulong:
		array = unsafe.Pointer(&(a.value.([]uint64))[0])
	case ArgType_ssize:
		array = unsafe.Pointer(&(a.value.([]int))[0])
	case ArgType_size:
		array = unsafe.Pointer(&(a.value.([]uint))[0])
	case ArgType_pointer:
		array = unsafe.Pointer(&(a.value.([]unsafe.Pointer))[0])
	default:
		panic(fmt.Sprintf("Unsupported array of type %#v", a.value))
	}

	(*(*unsafe.Pointer)(cArgPtr)) = array
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

func (a *Arg) getValueBoolArray() unsafe.Pointer {
	bools := a.value.([]bool)

	if bools == nil || len(bools) == 0 {
		return unsafe.Pointer(nil)
	}

	count := len(bools)
	arraySize := count * C.sizeof_gboolean
	cArray := (*C.gboolean)(C.malloc(C.ulong(arraySize)))
	cBooleans := (*[1 << 28]C.gboolean)(unsafe.Pointer(cArray))[:count:count]

	for i, boolean := range bools {
		var cBoolean C.gboolean = C.FALSE
		if boolean {
			cBoolean = C.TRUE
		}

		cBooleans[i] = cBoolean
	}

	return unsafe.Pointer(cArray)
}
