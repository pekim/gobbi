// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "C"

// GetInstancePrivateOffset is a wrapper around the C function g_type_class_get_instance_private_offset.
func (recv *TypeClass) GetInstancePrivateOffset() int32 {
	retC := C.g_type_class_get_instance_private_offset((C.gpointer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
