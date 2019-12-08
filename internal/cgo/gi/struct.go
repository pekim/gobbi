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
		err := fmt.Errorf("Failed to find struct '%s' in namespace '%s'", structName, namespace)
		handleError(err)
		return nil, err
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
		err := fmt.Errorf("Failed to find function '%s' in struct %s in namespace '%s'",
			funcName, s.structName, s.namespace)
		handleError(err)
		return nil, err
	}

	return &Function{
		namespace: s.namespace,
		ownerName: s.structName,
		funcName:  funcName,
		info:      invoker,
	}, nil
}

func (s *Struct) Alloc() unsafe.Pointer {
	size := C.g_struct_info_get_size(s.info)
	struct_ := C.malloc(size)
	return unsafe.Pointer(struct_)
}

func (s *Struct) Free(struct_ unsafe.Pointer) {
	C.free(struct_)
}
