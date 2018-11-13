// This is a generated file - DO NOT EDIT
// +build atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Get a list of properties applied to this object as a whole, as an #AtkAttributeSet consisting of
// name-value pairs. As such these attributes may be considered weakly-typed properties or annotations,
// as distinct from strongly-typed object data available via other get/set methods.
// Not all objects have explicit "name-value pair" #AtkAttributeSet properties.
/*

C function : atk_object_get_attributes
*/
func (recv *Object) GetAttributes() *AttributeSet {
	retC := C.atk_object_get_attributes((*C.AtkObject)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}
