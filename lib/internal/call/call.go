package call

// #cgo LDFLAGS: -ldl -lavcall
// #include "call.h"
import "C"

import (
	"errors"
	"log"
)

type ReturnType int

const (
	RT_INT = C.rt_int
)

type Data struct {
	ReturnType ReturnType
}

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

func Function(index int, data Data) {
	cData := C.CallData{
		return_type: C.ReturnType(data.ReturnType),
	}

	C.call_function(C.int(index), &cData)
}

func Close() {
	C.close()
}