package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Object struct {
	namespace  string
	objectName string
	info       *C.GIObjectInfo
}

func ObjectNew(namespace string, objectName string) (*Object, error) {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cObjectName := C.CString(objectName)
	defer C.free(unsafe.Pointer(cObjectName))

	object_ := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cObjectName,
	)
	if object_ == nil {
		err := fmt.Errorf("Failed to find object '%s' in namespace '%s'", objectName, namespace)
		handleError(err)
		return nil, err
	}

	return &Object{
		namespace:  namespace,
		objectName: objectName,
		info:       object_,
	}, nil
}

func (s *Object) InvokerNew(funcName string) (*Function, error) {
	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_object_info_find_method(
		s.info,
		cFuncName,
	)
	if invoker == nil {
		err := fmt.Errorf("Failed to find function '%s' in object %s in namespace '%s'",
			funcName, s.objectName, s.namespace)
		handleError(err)
		return nil, err
	}

	fi := &Function{
		namespace: s.namespace,
		ownerName: s.objectName,
		funcName:  funcName,
		info:      invoker,
	}
	fi.initTracing()

	return fi, nil
}
