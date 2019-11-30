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

func StructNew(namespace string, structName string) (*Struct, error) {
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
		return nil, fmt.Errorf("Failed to find struct '%s' in namespace '%s'", structName, namespace)
	}

	return &Struct{
		namespace:  namespace,
		structName: structName,
		info:       struct_,
	}, nil
}

func (s *Struct) InvokerNew(funcName string) (*Function, error) {
	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_struct_info_find_method(
		s.info,
		cFuncName,
	)
	if invoker == nil {
		return nil, fmt.Errorf("Failed to find function '%s' in struct %s in namespace '%s'",
			funcName, s.structName, s.namespace)
	}

	return &Function{
		namespace:  s.namespace,
		structName: s.structName,
		funcName:   funcName,
		info:       invoker,
	}, nil
}

func (s *Struct) Alloc() uintptr {
	size := C.g_struct_info_get_size(s.info)
	struct_ := C.malloc(size)
	return uintptr(struct_)
}

func (s *Struct) Free(struct_ uintptr) {
	C.free(unsafe.Pointer(struct_))
}
