// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Explicitly releases the binding between the source and the target
// property expressed by @binding.
//
// This function will release the reference that is being held on
// the @binding instance; if you want to hold on to the #GBinding instance
// after calling g_binding_unbind(), you will need to hold a reference
// to it.
/*

C function : g_binding_unbind
*/
func (recv *Binding) Unbind() {
	C.g_binding_unbind((*C.GBinding)(recv.native))

	return
}
