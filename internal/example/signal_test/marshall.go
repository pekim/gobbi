package main

//#include <glib-object.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

//export goMarshal
func goMarshal(closure *C.GClosure, retValue *C.GValue,
	nParams C.guint, params *C.GValue,
	invocationHint C.gpointer, marshalData *C.GValue,
) {
	callValues := []reflect.Value{}

	closureLock.Lock()
	defer closureLock.Unlock()
	cb := closures[closure]
	cbFunc := reflect.ValueOf(cb)
	cbType := reflect.TypeOf(cb)

	param := params
	for p := 0; p < int(nParams); p++ {
		if C.g_type_check_value_holds(param, C.G_TYPE_OBJECT) == C.TRUE {
			if p > 0 {
				object := C.g_value_get_object(param)
				fmt.Println("object", object)
			}
		} else if C.g_type_check_value_holds(param, C.G_TYPE_POINTER) == C.TRUE {
			pointer := C.g_value_get_pointer(param)
			fmt.Println("pointer", pointer)
		} else if C.g_type_check_value_holds(param, C.G_TYPE_INVALID) == C.TRUE {
			fmt.Println("G_TYPE_INVALID")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_NONE) == C.TRUE {
			fmt.Println("G_TYPE_NONE")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_INTERFACE) == C.TRUE {
			fmt.Println("G_TYPE_INTERFACE")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_CHAR) == C.TRUE {
			fmt.Println("G_TYPE_CHAR")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_UCHAR) == C.TRUE {
			fmt.Println("G_TYPE_UCHAR")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_BOOLEAN) == C.TRUE {
			fmt.Println("G_TYPE_BOOLEAN")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_INT) == C.TRUE {
			fmt.Println("G_TYPE_INT")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_UINT) == C.TRUE {
			fmt.Println("G_TYPE_UINT")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_LONG) == C.TRUE {
			fmt.Println("G_TYPE_LONG")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_ULONG) == C.TRUE {
			fmt.Println("G_TYPE_ULONG")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_INT64) == C.TRUE {
			fmt.Println("G_TYPE_INT64")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_UINT64) == C.TRUE {
			fmt.Println("G_TYPE_UINT64")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_ENUM) == C.TRUE {
			fmt.Println("G_TYPE_ENUM")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_FLAGS) == C.TRUE {
			fmt.Println("G_TYPE_FLAGS")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_FLOAT) == C.TRUE {
			fmt.Println("G_TYPE_FLOAT")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_DOUBLE) == C.TRUE {
			fmt.Println("G_TYPE_DOUBLE")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_STRING) == C.TRUE {
			fmt.Println("G_TYPE_STRING")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_POINTER) == C.TRUE {
			fmt.Println("G_TYPE_POINTER")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_BOXED) == C.TRUE {
			if !assertMarshallParamType(cbType, p, "unsafe.Pointer") {
				return
			}

			boxed := C.g_value_get_boxed(param)
			paramValue := reflect.ValueOf(unsafe.Pointer(boxed))
			callValues = append(callValues, paramValue)
		} else if C.g_type_check_value_holds(param, C.G_TYPE_PARAM) == C.TRUE {
			fmt.Println("G_TYPE_PARAM")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_OBJECT) == C.TRUE {
			fmt.Println("G_TYPE_OBJECT")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_GTYPE) == C.TRUE {
			fmt.Println("G_TYPE_GTYPE")
		} else if C.g_type_check_value_holds(param, C.G_TYPE_VARIANT) == C.TRUE {
			fmt.Println("G_TYPE_VARIANT")
		} else {
			panic("unexpected param type")
		}

		// increment, to point to next param
		paramsAsUintptr := uintptr(unsafe.Pointer(param))
		paramsAsUintptr += C.sizeof_GValue
		param = (*C.GValue)(unsafe.Pointer(paramsAsUintptr))
	}

	cbFunc.Call(callValues)
}

func assertMarshallParamType(cbType reflect.Type, p int, expected string) bool {
	paramType := cbType.In(p - 1)
	if paramType.String() != expected {
		fmt.Printf("Expected param #%d of '%s' to be %s, but is %s; callback will not be called\n",
			p, cbType, expected, paramType)
		return false
	}

	return true
}
