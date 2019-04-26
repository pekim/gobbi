package call

// #cgo LDFLAGS: -ldl
// #cgo LDFLAGS: -lavcall
//
// #include "call.h"
import "C"

import (
	"errors"
	"log"
)

func init() {
	err := open()
	if err != nil {
		log.Fatal(err)
	}
}

func open() error {
	cString := C.open()
	if cString != nil {
		return errors.New(C.GoString(cString))
	}

	return nil
}

func Function(index int, data *Data) {
	cData := C.CallData{
		return_type: C.Type(data.Return.Type),
	}

	C.call_function(C.int(index), &cData)

	switch data.Return.Type {
	case TYPE_INT:
		data.Return.value = cData.value_int
	case TYPE_UINT:
		data.Return.value = cData.value_uint
	case TYPE_LONG:
		data.Return.value = cData.value_long
	case TYPE_DOUBLE:
		data.Return.value = cData.value_double
	case TYPE_VOID:
		// nothing
	}
}

func Close() {
	C.close()
}
