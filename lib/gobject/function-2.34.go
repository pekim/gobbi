// This is a generated file - DO NOT EDIT
// +build gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Ensures that the indicated @type has been registered with the
// type system, and its _class_init() method has been run.
//
// In theory, simply calling the type's _get_type() method (or using
// the corresponding macro) is supposed take care of this. However,
// _get_type() methods are often marked %G_GNUC_CONST for performance
// reasons, even though this is technically incorrect (since
// %G_GNUC_CONST requires that the function not have side effects,
// which _get_type() methods do on the first call). As a result, if
// you write a bare call to a _get_type() macro, it may get optimized
// out by the compiler. Using g_type_ensure() guarantees that the
// type's _get_type() method is called.
/*

C function : g_type_ensure
*/
func TypeEnsure(type_ Type) {
	c_type := (C.GType)(type_)

	C.g_type_ensure(c_type)

	return
}
