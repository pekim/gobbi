package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type FunctionInvoker struct {
	namespace string
	funcName  string
	info      *C.GIFunctionInfo
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
		namespace: namespace,
		funcName:  funcName,
		info:      invoker,
	}
}

func (fi *FunctionInvoker) Call() {
	var returnValue C.GIArgument
	var err *C.GError

	invoked := C.g_function_info_invoke(
		fi.info,
		nil, 0,
		nil, 0,
		&returnValue,
		&err,
	) == C.TRUE

	if err != nil {
		message := C.GoString(err.message)
		panic(message)
	}

	if !invoked {
		panic(fmt.Sprintf("%s.%s not called", fi.namespace, fi.funcName))
	}

	fmt.Println("called!!!!!!!!!!!", invoked)
}
