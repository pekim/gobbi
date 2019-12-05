package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func objectFieldInfo(objectInfo *Object, name string) (*C.GIFieldInfo, bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	// TODO There is no g_object_info_find_field, and it's a bit silly to execute this look every time.
	// So maybe do it one them the Object is created, and keep the results in a map?
	var fieldInfo *C.GIFieldInfo
	for i := C.int(0); i < C.g_object_info_get_n_fields(objectInfo.info); i++ {
		info := C.g_object_info_get_field(objectInfo.info, i)
		cName := C.g_base_info_get_name(info)
		goName := C.GoString(cName)
		if name == goName {
			fieldInfo = info
			break
		}
	}

	if fieldInfo == nil {
		err := fmt.Errorf("Failed to find field '%s' in object %s in namespace '%s'",
			name, objectInfo.objectName, objectInfo.namespace)
		handleError(err)
		return nil, false
	}

	return fieldInfo, true
}

func ObjectFieldGet(objectInfo *Object, object unsafe.Pointer, name string) Argument {
	var value Argument

	info, gotInfo := objectFieldInfo(objectInfo, name)
	if !gotInfo {
		return value
	}
	defer C.g_base_info_unref(info)

	ok := C.g_field_info_get_field(
		info,
		C.gpointer(object),
		(*C.GIArgument)(&value),
	)
	if ok != C.TRUE {
		err := fmt.Errorf("Failed to get value for field '%s' in object %s in namespace '%s'",
			name, objectInfo.objectName, objectInfo.namespace)
		handleError(err)
	}

	return value
}

func ObjectFieldSet(objectInfo *Object, object unsafe.Pointer, name string, value Argument) {
	info, gotInfo := objectFieldInfo(objectInfo, name)
	if !gotInfo {
		return
	}
	defer C.g_base_info_unref(info)

	ok := C.g_field_info_set_field(
		info,
		C.gpointer(object),
		(*C.GIArgument)(&value),
	)
	if ok != C.TRUE {
		err := fmt.Errorf("Failed to set value for field '%s' in object %s in namespace '%s'",
			name, objectInfo.objectName, objectInfo.namespace)
		handleError(err)
	}
}
