package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Interface struct {
	namespace     string
	interfaceName string
	info          *C.GIInterfaceInfo
}

func InterfaceNew(namespace string, interfaceName string) (*Interface, error) {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cInterfaceName := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(cInterfaceName))

	interface_ := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cInterfaceName,
	)
	if interface_ == nil {
		err := fmt.Errorf("Failed to find interface '%s' in namespace '%s'", interfaceName, namespace)
		handleError(err)
		return nil, err
	}

	return &Interface{
		namespace:     namespace,
		interfaceName: interfaceName,
		info:          interface_,
	}, nil
}

func (s *Interface) InvokerNew(funcName string) (*Function, error) {
	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_interface_info_find_method(
		s.info,
		cFuncName,
	)
	if invoker == nil {
		err := fmt.Errorf("Failed to find function '%s' in interface %s in namespace '%s'",
			funcName, s.interfaceName, s.namespace)
		handleError(err)
		return nil, err
	}

	fi := &Function{
		namespace: s.namespace,
		ownerName: s.interfaceName,
		funcName:  funcName,
		info:      invoker,
	}
	fi.initTracing()

	return fi, nil
}
