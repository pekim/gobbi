package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/internal/cgo"
	"unsafe"
)

type Function2 struct {
	namespace      string
	ownerName      string
	funcName       string
	fullName       string
	hasReturnValue bool
	info           *C.GIFunctionInfo
}

func Function2InvokerNew(namespace string, funcName string) (*Function2, error) {
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

	fi := &Function2{
		namespace: namespace,
		funcName:  funcName,
		info:      invoker,
	}
	fi.initTracing()

	return fi, nil
}

func (fi *Function2) initTracing() {
	if fi.ownerName != "" {
		fi.fullName = fmt.Sprintf("%s : %s.%s", fi.namespace, fi.ownerName, fi.funcName)
	} else {
		fi.fullName = fmt.Sprintf("%s : %s", fi.namespace, fi.funcName)
	}

	returnTypeInfo := C.g_callable_info_get_return_type(fi.info)
	returnTypeTag := C.g_type_info_get_tag(returnTypeInfo)
	fi.hasReturnValue = returnTypeTag != C.GI_TYPE_TAG_VOID
}

func (fi *Function2) Invoke(args []Arg, inLen int, outLen int, returnArg Arg) Arg {
	var cReturnValue C.GIArgument
	var err *C.GError

	inArgs := make([]C.GIArgument, inLen, inLen)
	outArgs := make([]C.GIArgument, outLen, outLen)

	// populate inArgs and outArgs
	inIndex := 0
	outIndex := 0
	for _, arg := range args {
		cArg := arg.getValue()

		if arg.in {
			inArgs[inIndex] = cArg
			inIndex++
		}

		if arg.in {
			outArgs[outIndex] = cArg
			outIndex++
		}
	}

	cInLen := C.int(inLen)
	var cIn *C.GIArgument
	if inLen > 0 {
		cIn = (*C.GIArgument)(&inArgs[0])
	}

	cOutLen := C.int(outLen)
	var cOut *C.GIArgument
	if outLen > 0 {
		cOut = (*C.GIArgument)(&outArgs[0])
	}

	// invoke
	invoked := C.g_function_info_invoke(
		fi.info,
		cIn, cInLen,
		cOut, cOutLen,
		(*C.GIArgument)(&cReturnValue),
		&err,
	) == C.TRUE

	(&returnArg).setValue(cReturnValue)

	if cgo.Tracing() {
		//fi.trace(in, out, returnValue)
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
	return returnArg

	//// invoke
	//invoked := C.g_function_info_invoke(
	//	fi.info,
	//	cIn, cInLen,
	//	cOut, cOutLen,
	//	(*C.GIArgument)(&returnValue),
	//	&err,
	//) == C.TRUE
	//
	//if cgo.Tracing() {
	//	fi.trace(in, out, returnValue)
	//}
	//
	//// check error
	//if err != nil {
	//	message := C.GoString(err.message)
	//	panic(message)
	//}
	//
	//// verify invoke happened
	//if !invoked {
	//	panic(fmt.Sprintf("%s.%s not called", fi.namespace, fi.funcName))
	//}
	//return returnValue
}

func (fi *Function2) trace(in []Argument, out []Argument, returnValue Argument) {
	cgo.Trace(fmt.Sprintf("%s\n%s%s%s\n",
		fi.fullName,
		fi.formatTraceData(in, "in", len(in) > 0),
		fi.formatTraceData(out, "out", len(out) > 0),
		fi.formatTraceData(returnValue, "ret", fi.hasReturnValue),
	))
}

func (fi *Function2) formatTraceData(data interface{}, name string, include bool) string {
	if include {
		return fmt.Sprintf("  %3s %v\n", name, data)
	}

	return ""
}
