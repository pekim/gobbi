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
		return_type: C.ReturnType(data.Return.Type),
	}

	C.call_function(C.int(index), &cData)

	switch data.Return.Type {
	case RT_INT:
		data.Return.Int = int(cData.return_int)
	case RT_VOID:
		// nothing
	}
}

func Close() {
	C.close()
}
