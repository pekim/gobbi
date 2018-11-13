// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// All the fields in the GInitiallyUnowned structure
// are private to the #GInitiallyUnowned implementation and should never be
// accessed directly.
/*

C type

GInitiallyUnowned
*/
type InitiallyUnowned struct {
	native *C.GInitiallyUnowned
	// All fields are private
}

func InitiallyUnownedNewFromC(u unsafe.Pointer) *InitiallyUnowned {
	c := (*C.GInitiallyUnowned)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnowned{native: c}

	return g
}

func (recv *InitiallyUnowned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *InitiallyUnowned) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to InitiallyUnowned.
// Exercise care, as this is a potentially dangerous function if the Object is not a InitiallyUnowned.
func CastToInitiallyUnowned(object *Object) *InitiallyUnowned {
	return InitiallyUnownedNewFromC(object.ToC())
}

// All the fields in the GObject structure are private
// to the #GObject implementation and should never be accessed directly.
/*

C type

GObject
*/
type Object struct {
	native *C.GObject
	// All fields are private
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	c := (*C.GObject)(u)
	if c == nil {
		return nil
	}

	g := &Object{native: c}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CastToWidget down casts any arbitary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *Object) *Object {
	return ObjectNewFromC(object.ToC())
}

// Unsupported signal 'notify' for Object : unsupported parameter pspec : type ParamSpec : Blacklisted record : GParamSpec

// Unsupported : g_object_new : unsupported parameter ... : varargs

// Unsupported : g_object_new_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_object_newv : unsupported parameter parameters :

// Adds a weak reference from weak_pointer to @object to indicate that
// the pointer located at @weak_pointer_location is only valid during
// the lifetime of @object. When the @object is finalized,
// @weak_pointer will be set to %NULL.
//
// Note that as with g_object_weak_ref(), the weak references created by
// this method are not thread-safe: they cannot safely be used in one
// thread if the object's last g_object_unref() might happen in another
// thread. Use #GWeakRef if thread-safety is required.
/*

C function

g_object_add_weak_pointer
*/
func (recv *Object) AddWeakPointer(weakPointerLocation uintptr) {
	c_weak_pointer_location := (C.gpointer)(weakPointerLocation)

	C.g_object_add_weak_pointer((*C.GObject)(recv.native), &c_weak_pointer_location)

	return
}

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// Increases the freeze count on @object. If the freeze count is
// non-zero, the emission of "notify" signals on @object is
// stopped. The signals are queued until the freeze count is decreased
// to zero. Duplicate notifications are squashed so that at most one
// #GObject::notify signal is emitted for each property modified while the
// object is frozen.
//
// This is necessary for accessors that modify multiple properties to prevent
// premature notification while the object is still being modified.
/*

C function

g_object_freeze_notify
*/
func (recv *Object) FreezeNotify() {
	C.g_object_freeze_notify((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_get : unsupported parameter ... : varargs

// Gets a named field from the objects table of associations (see g_object_set_data()).
/*

C function

g_object_get_data
*/
func (recv *Object) GetData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_get_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets a property of an object. @value must have been initialized to the
// expected type of the property (or a type to which the expected type can be
// transformed) using g_value_init().
//
// In general, a copy is made of the property contents and the caller is
// responsible for freeing the memory by calling g_value_unset().
//
// Note that g_object_get_property() is really intended for language
// bindings, g_object_get() is much more convenient for C programming.
/*

C function

g_object_get_property
*/
func (recv *Object) GetProperty(propertyName string, value *Value) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.g_object_get_property((*C.GObject)(recv.native), c_property_name, c_value)

	return
}

// This function gets back user data pointers stored via
// g_object_set_qdata().
/*

C function

g_object_get_qdata
*/
func (recv *Object) GetQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_get_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Emits a "notify" signal for the property @property_name on @object.
//
// When possible, eg. when signaling a property change from within the class
// that registered the property, you should use g_object_notify_by_pspec()
// instead.
//
// Note that emission of the notify signal may be blocked with
// g_object_freeze_notify(). In this case, the signal emissions are queued
// and will be emitted (in reverse order) when g_object_thaw_notify() is
// called.
/*

C function

g_object_notify
*/
func (recv *Object) Notify(propertyName string) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	C.g_object_notify((*C.GObject)(recv.native), c_property_name)

	return
}

// Increases the reference count of @object.
//
// Since GLib 2.56, if `GLIB_VERSION_MAX_ALLOWED` is 2.56 or greater, the type
// of @object will be propagated to the return type (using the GCC typeof()
// extension), so any casting the caller needs to do on the return type must be
// explicit.
/*

C function

g_object_ref
*/
func (recv *Object) Ref() uintptr {
	retC := C.g_object_ref((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Removes a weak reference from @object that was previously added
// using g_object_add_weak_pointer(). The @weak_pointer_location has
// to match the one used with g_object_add_weak_pointer().
/*

C function

g_object_remove_weak_pointer
*/
func (recv *Object) RemoveWeakPointer(weakPointerLocation uintptr) {
	c_weak_pointer_location := (C.gpointer)(weakPointerLocation)

	C.g_object_remove_weak_pointer((*C.GObject)(recv.native), &c_weak_pointer_location)

	return
}

// Releases all references to other objects. This can be used to break
// reference cycles.
//
// This function should only be called from object system implementations.
/*

C function

g_object_run_dispose
*/
func (recv *Object) RunDispose() {
	C.g_object_run_dispose((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Each object carries around a table of associations from
// strings to pointers.  This function lets you set an association.
//
// If the object already had an association with that name,
// the old association will be destroyed.
/*

C function

g_object_set_data
*/
func (recv *Object) SetData(key string, data uintptr) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_data := (C.gpointer)(data)

	C.g_object_set_data((*C.GObject)(recv.native), c_key, c_data)

	return
}

// Unsupported : g_object_set_data_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Sets a property on an object.
/*

C function

g_object_set_property
*/
func (recv *Object) SetProperty(propertyName string, value *Value) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.g_object_set_property((*C.GObject)(recv.native), c_property_name, c_value)

	return
}

// This sets an opaque, named pointer on an object.
// The name is specified through a #GQuark (retrived e.g. via
// g_quark_from_static_string()), and the pointer
// can be gotten back from the @object with g_object_get_qdata()
// until the @object is finalized.
// Setting a previously set user data pointer, overrides (frees)
// the old pointer set, using #NULL as pointer essentially
// removes the data stored.
/*

C function

g_object_set_qdata
*/
func (recv *Object) SetQdata(quark glib.Quark, data uintptr) {
	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_object_set_qdata((*C.GObject)(recv.native), c_quark, c_data)

	return
}

// Unsupported : g_object_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Remove a specified datum from the object's data associations,
// without invoking the association's destroy handler.
/*

C function

g_object_steal_data
*/
func (recv *Object) StealData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_steal_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function gets back user data pointers stored via
// g_object_set_qdata() and removes the @data from object
// without invoking its destroy() function (if any was
// set).
// Usually, calling this function is only required to update
// user data pointers with a destroy notifier, for example:
// |[<!-- language="C" -->
// void
// object_add_to_user_list (GObject     *object,
// const gchar *new_string)
// {
// the quark, naming the object data
// GQuark quark_string_list = g_quark_from_static_string ("my-string-list");
// retrive the old string list
// GList *list = g_object_steal_qdata (object, quark_string_list);
//
// prepend new string
// list = g_list_prepend (list, g_strdup (new_string));
// this changed 'list', so we need to set it again
// g_object_set_qdata_full (object, quark_string_list, list, free_string_list);
// }
// static void
// free_string_list (gpointer data)
// {
// GList *node, *list = data;
//
// for (node = list; node; node = node->next)
// g_free (node->data);
// g_list_free (list);
// }
// ]|
// Using g_object_get_qdata() in the above example, instead of
// g_object_steal_qdata() would have left the destroy function set,
// and thus the partial string list would have been freed upon
// g_object_set_qdata_full().
/*

C function

g_object_steal_qdata
*/
func (recv *Object) StealQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_steal_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Reverts the effect of a previous call to
// g_object_freeze_notify(). The freeze count is decreased on @object
// and when it reaches zero, queued "notify" signals are emitted.
//
// Duplicate notifications for each property are squashed so that at most one
// #GObject::notify signal is emitted for each property, in the reverse order
// in which they have been queued.
//
// It is an error to call this function when the freeze count is zero.
/*

C function

g_object_thaw_notify
*/
func (recv *Object) ThawNotify() {
	C.g_object_thaw_notify((*C.GObject)(recv.native))

	return
}

// Decreases the reference count of @object. When its reference count
// drops to 0, the object is finalized (i.e. its memory is freed).
//
// If the pointer to the #GObject may be reused in future (for example, if it is
// an instance variable of another object), it is recommended to clear the
// pointer to %NULL rather than retain a dangling pointer to a potentially
// invalid #GObject instance. Use g_clear_object() for this.
/*

C function

g_object_unref
*/
func (recv *Object) Unref() {
	C.g_object_unref((C.gpointer)(recv.native))

	return
}

// This function essentially limits the life time of the @closure to
// the life time of the object. That is, when the object is finalized,
// the @closure is invalidated by calling g_closure_invalidate() on
// it, in order to prevent invocations of the closure with a finalized
// (nonexisting) object. Also, g_object_ref() and g_object_unref() are
// added as marshal guards to the @closure, to ensure that an extra
// reference count is held on @object during invocation of the
// @closure.  Usually, this function will be called on closures that
// use this @object as closure data.
/*

C function

g_object_watch_closure
*/
func (recv *Object) WatchClosure(closure *Closure) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	C.g_object_watch_closure((*C.GObject)(recv.native), c_closure)

	return
}

// Unsupported : g_object_weak_ref : unsupported parameter notify : no type generator for WeakNotify (GWeakNotify) for param notify

// Unsupported : g_object_weak_unref : unsupported parameter notify : no type generator for WeakNotify (GWeakNotify) for param notify

// Blacklisted : GParamSpec

// Blacklisted : GParamSpecBoolean

// Blacklisted : GParamSpecBoxed

// Blacklisted : GParamSpecChar

// Blacklisted : GParamSpecDouble

// Blacklisted : GParamSpecEnum

// Blacklisted : GParamSpecFlags

// Blacklisted : GParamSpecFloat

// Blacklisted : GParamSpecInt

// Blacklisted : GParamSpecInt64

// Blacklisted : GParamSpecLong

// Blacklisted : GParamSpecObject

// Blacklisted : GParamSpecParam

// Blacklisted : GParamSpecPointer

// Blacklisted : GParamSpecString

// Blacklisted : GParamSpecUChar

// Blacklisted : GParamSpecUInt

// Blacklisted : GParamSpecUInt64

// Blacklisted : GParamSpecULong

// Blacklisted : GParamSpecUnichar

// Blacklisted : GParamSpecValueArray

// Blacklisted : GTypeModule
