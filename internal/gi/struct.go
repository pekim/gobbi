package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Struct struct {
	namespace  string
	structName string
	info       *C.GIStructInfo
}

func StructNew(namespace string, structName string) *Struct {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cStructName := C.CString(structName)
	defer C.free(unsafe.Pointer(cStructName))

	struct_ := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cStructName,
	)
	if struct_ == nil {
		panic(fmt.Sprintf("Failed to find struct '%s' in namespace '%s'", structName, namespace))
	}

	return &Struct{
		namespace:  namespace,
		structName: structName,
		info:       struct_,
	}
}

func (s *Struct) InvokerNew(funcName string) *Function {
	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_struct_info_find_method(
		s.info,
		cFuncName,
	)
	if invoker == nil {
		panic(fmt.Sprintf("Failed to find function '%s' in struct %s in namespace '%s'",
			funcName, s.structName, s.namespace))
	}

	return &Function{
		namespace:  s.namespace,
		structName: s.structName,
		funcName:   funcName,
		info:       invoker,
	}
}
