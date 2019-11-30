package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func fieldInfo(structInfo *Struct, name string) (*C.GIFieldInfo, bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	fieldInfo := C.g_struct_info_find_field(
		structInfo.info,
		cName,
	)
	if fieldInfo == nil {
		err := fmt.Errorf("Failed to find field '%s' in struct %s in namespace '%s'",
			name, structInfo.structName, structInfo.namespace)
		handleError(err)
		return nil, false
	}

	return fieldInfo, true
}

func FieldGet(structInfo *Struct, struct_ uintptr, name string) Argument {
	var value Argument

	info, gotInfo := fieldInfo(structInfo, name)
	if !gotInfo {
		return value
	}
	defer C.g_base_info_unref(info)

	ok := C.g_field_info_get_field(
		info,
		C.gpointer(struct_),
		(*C.GIArgument)(&value),
	)
	if ok != C.TRUE {
		err := fmt.Errorf("Failed to get value for field '%s' in struct %s in namespace '%s'",
			name, structInfo.structName, structInfo.namespace)
		handleError(err)
	}

	return value
}

func FieldSet(structInfo *Struct, struct_ uintptr, name string, value Argument) {
	info, gotInfo := fieldInfo(structInfo, name)
	if !gotInfo {
		return
	}
	defer C.g_base_info_unref(info)

	ok := C.g_field_info_set_field(
		info,
		C.gpointer(struct_),
		(*C.GIArgument)(&value),
	)
	if ok != C.TRUE {
		err := fmt.Errorf("Failed to set value for field '%s' in struct %s in namespace '%s'",
			name, structInfo.structName, structInfo.namespace)
		handleError(err)
	}
}
