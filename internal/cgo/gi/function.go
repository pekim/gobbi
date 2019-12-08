package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Function struct {
	namespace      string
	ownerName      string
	funcName       string
	fullName       string
	hasReturnValue bool
	info           *C.GIFunctionInfo
}

func FunctionInvokerNew(namespace string, funcName string) (*Function, error) {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cFuncName,
	)

	if invoker == nil {
		err := fmt.Errorf("Failed to find function '%s' in namespace '%s'", funcName, namespace)
		handleError(err)
		return nil, err
	}

	fi := &Function{
		namespace: namespace,
		funcName:  funcName,
		info:      invoker,
	}
	fi.initTracing()

	return fi, nil
}

func (fi *Function) initTracing() {
	if fi.ownerName != "" {
		fi.fullName = fmt.Sprintf("%s  %s  %s", fi.namespace, fi.ownerName, fi.funcName)
	} else {
		fi.fullName = fmt.Sprintf("%s  %s", fi.namespace, fi.funcName)
	}

	returnTypeInfo := C.g_callable_info_get_return_type(fi.info)
	returnTypeTag := C.g_type_info_get_tag(returnTypeInfo)
	fi.hasReturnValue = returnTypeTag != C.GI_TYPE_TAG_VOID
}

func (fi *Function) Invoke(in []Argument, out []Argument) Argument {
	var returnValue Argument
	var err *C.GError

	// prepare in args
	var cIn *C.GIArgument
	var cInLen C.int
	if in != nil {
		cIn = (*C.GIArgument)(&in[0])
		cInLen = C.int(len(in))
	}

	// prepare out args
	var cOut *C.GIArgument
	var cOutLen C.int
	if out != nil {
		cOut = (*C.GIArgument)(&out[0])
		cOutLen = C.int(len(out))
	}

	// invoke
	invoked := C.g_function_info_invoke(
		fi.info,
		cIn, cInLen,
		cOut, cOutLen,
		(*C.GIArgument)(&returnValue),
		&err,
	) == C.TRUE

	if tracing() {
		fi.trace(in, out, returnValue)
	}

	// check error
	if err != nil {
		message := C.GoString(err.message)
		panic(message)
	}

	// verify invoke happened
	if !invoked {
		panic(fmt.Sprintf("%s.%s not called", fi.namespace, fi.funcName))
	}

	return returnValue
}

func (fi *Function) trace(in []Argument, out []Argument, returnValue Argument) {
	trace(fmt.Sprintf("%s\n%s%s%s\n",
		fi.fullName,
		fi.formatTraceData(in, "in", len(in) > 0),
		fi.formatTraceData(out, "out", len(out) > 0),
		fi.formatTraceData(returnValue, "ret", fi.hasReturnValue),
	))
}

func (fi *Function) formatTraceData(data interface{}, name string, include bool) string {
	if include {
		return fmt.Sprintf("  %3s %v\n", name, data)
	}

	return ""
}
