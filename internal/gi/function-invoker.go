package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type FunctionInvoker struct {
	info *C.GIFunctionInfo
}

func FunctionInvokerNew(namespace string, funcName string) *FunctionInvoker {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cFuncName,
	)

	return &FunctionInvoker{
		info: invoker,
	}
}
