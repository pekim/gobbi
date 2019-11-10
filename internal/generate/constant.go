package generate

// #include <gobject-introspection-1.0/girepository.h>
import "C"

import ()

func (r *repository) generateConstant(info *C.GIBaseInfo) {
	//cName := C.g_base_info_get_name(info)
	//name := C.GoString(cName)
	//
	//constant := (*C.GIConstantInfo)(info)
	//var value C.GIArgument
	//size := C.g_constant_info_get_value(constant, &value)
	//
	//fmt.Println(name, size)
}
