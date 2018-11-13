// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Registers @property_id as referring to a property with the name
// @name in a parent class or in an interface implemented by @oclass.
// This allows this class to "override" a property implementation in
// a parent class or to provide the implementation of a property from
// an interface.
//
// Internally, overriding is implemented by creating a property of type
// #GParamSpecOverride; generally operations that query the properties of
// the object class, such as g_object_class_find_property() or
// g_object_class_list_properties() will return the overridden
// property. However, in one case, the @construct_properties argument of
// the @constructor virtual function, the #GParamSpecOverride is passed
// instead, so that the @param_id field of the #GParamSpec will be
// correct.  For virtually all uses, this makes no difference. If you
// need to get the overridden property, you can call
// g_param_spec_get_redirect_target().
/*

C function

g_object_class_override_property
*/
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	c_property_id := (C.guint)(propertyId)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_object_class_override_property((*C.GObjectClass)(recv.native), c_property_id, c_name)

	return
}

// Registers a private structure for an instantiatable type.
//
// When an object is allocated, the private structures for
// the type and all of its parent types are allocated
// sequentially in the same memory block as the public
// structures, and are zero-filled.
//
// Note that the accumulated size of the private structures of
// a type and all its parent types cannot exceed 64 KiB.
//
// This function should be called in the type's class_init() function.
// The private structure can be retrieved using the
// G_TYPE_INSTANCE_GET_PRIVATE() macro.
//
// The following example shows attaching a private structure
// MyObjectPrivate to an object MyObject defined in the standard
// GObject fashion in the type's class_init() function.
//
// Note the use of a structure member "priv" to avoid the overhead
// of repeatedly calling MY_OBJECT_GET_PRIVATE().
//
// |[<!-- language="C" -->
// typedef struct _MyObject        MyObject;
// typedef struct _MyObjectPrivate MyObjectPrivate;
//
// struct _MyObject {
// GObject parent;
//
// MyObjectPrivate *priv;
// };
//
// struct _MyObjectPrivate {
// int some_field;
// };
//
// static void
// my_object_class_init (MyObjectClass *klass)
// {
// g_type_class_add_private (klass, sizeof (MyObjectPrivate));
// }
//
// static void
// my_object_init (MyObject *my_object)
// {
// my_object->priv = G_TYPE_INSTANCE_GET_PRIVATE (my_object,
// MY_TYPE_OBJECT,
// MyObjectPrivate);
// my_object->priv->some_field will be automatically initialised to 0
// }
//
// static int
// my_object_get_some_field (MyObject *my_object)
// {
// MyObjectPrivate *priv;
//
// g_return_val_if_fail (MY_IS_OBJECT (my_object), 0);
//
// priv = my_object->priv;
//
// return priv->some_field;
// }
// ]|
/*

C function

g_type_class_add_private
*/
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	c_private_size := (C.gsize)(privateSize)

	C.g_type_class_add_private((C.gpointer)(recv.native), c_private_size)

	return
}

// Sets the contents of a %G_TYPE_BOXED derived #GValue to @v_boxed
// and takes over the ownership of the callers reference to @v_boxed;
// the caller doesn't have to unref it any more.
/*

C function

g_value_take_boxed
*/
func (recv *Value) TakeBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_take_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// Sets the contents of a %G_TYPE_OBJECT derived #GValue to @v_object
// and takes over the ownership of the callers reference to @v_object;
// the caller doesn't have to unref it any more (i.e. the reference
// count of the object is not increased).
//
// If you want the #GValue to hold its own reference to @v_object, use
// g_value_set_object() instead.
/*

C function

g_value_take_object
*/
func (recv *Value) TakeObject(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_take_object((*C.GValue)(recv.native), c_v_object)

	return
}

// Unsupported : g_value_take_param : unsupported parameter param : Blacklisted record : GParamSpec

// Sets the contents of a %G_TYPE_STRING #GValue to @v_string.
/*

C function

g_value_take_string
*/
func (recv *Value) TakeString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_take_string((*C.GValue)(recv.native), c_v_string)

	return
}
