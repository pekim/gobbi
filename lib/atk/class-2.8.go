// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

package atk

import "C"

// GetObjectLocale is a wrapper around the C function atk_object_get_object_locale.
func (recv *Object) GetObjectLocale() string {
	retC := C.atk_object_get_object_locale((*C.AtkObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
