package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func FieldGet(structInfo *Struct, struct_ uintptr, name string) Argument {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var value Argument

	fieldInfo := C.g_struct_info_find_field(
		structInfo.info,
		cName,
	)
	if fieldInfo == nil {
		err := fmt.Errorf("Failed to find field '%s' in struct %s in namespace '%s'",
			name, structInfo.structName, structInfo.namespace)
		handleError(err)
		return value
	}
	defer C.g_base_info_unref(fieldInfo)

	ok := C.g_field_info_get_field(
		fieldInfo,
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
