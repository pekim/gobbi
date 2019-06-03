package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func SubclassCreate() {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GObjectClass
	typeInfo.instance_size = C.sizeof_GObject

	cTypeName := C.CString("my_type")
	defer C.free(unsafe.Pointer(cTypeName))

	//fmt.Println(typeInfo)
	gtype := C.g_type_register_static(C.G_TYPE_OBJECT, cTypeName, &typeInfo, 0)
	fmt.Println(gtype)
}
