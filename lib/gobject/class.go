// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// InitiallyUnowned is a wrapper around the C record GInitiallyUnowned.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InitiallyUnowned) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InitiallyUnowned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InitiallyUnowned with another InitiallyUnowned, and returns true if they represent the same GObject.
func (recv *InitiallyUnowned) Equals(other *InitiallyUnowned) bool {
	return other.ToC() == recv.ToC()
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

// Object is a wrapper around the C record GObject.
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

// Equals compares this Object with another Object, and returns true if they represent the same GObject.
func (recv *Object) Equals(other *Object) bool {
	return other.ToC() == recv.ToC()
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

// ObjectCompatControl is a wrapper around the C function g_object_compat_control.
func ObjectCompatControl(what uint64, data uintptr) uint64 {
	c_what := (C.gsize)(what)

	c_data := (C.gpointer)(data)

	retC := C.g_object_compat_control(c_what, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// AddWeakPointer is a wrapper around the C function g_object_add_weak_pointer.
func (recv *Object) AddWeakPointer(weakPointerLocation uintptr) {
	c_weak_pointer_location := (C.gpointer)(weakPointerLocation)

	C.g_object_add_weak_pointer((*C.GObject)(recv.native), &c_weak_pointer_location)

	return
}

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// FreezeNotify is a wrapper around the C function g_object_freeze_notify.
func (recv *Object) FreezeNotify() {
	C.g_object_freeze_notify((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_get : unsupported parameter ... : varargs

// GetData is a wrapper around the C function g_object_get_data.
func (recv *Object) GetData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_get_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetProperty is a wrapper around the C function g_object_get_property.
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

// GetQdata is a wrapper around the C function g_object_get_qdata.
func (recv *Object) GetQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_get_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Notify is a wrapper around the C function g_object_notify.
func (recv *Object) Notify(propertyName string) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	C.g_object_notify((*C.GObject)(recv.native), c_property_name)

	return
}

// Ref is a wrapper around the C function g_object_ref.
func (recv *Object) Ref() uintptr {
	retC := C.g_object_ref((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// RemoveWeakPointer is a wrapper around the C function g_object_remove_weak_pointer.
func (recv *Object) RemoveWeakPointer(weakPointerLocation uintptr) {
	c_weak_pointer_location := (C.gpointer)(weakPointerLocation)

	C.g_object_remove_weak_pointer((*C.GObject)(recv.native), &c_weak_pointer_location)

	return
}

// RunDispose is a wrapper around the C function g_object_run_dispose.
func (recv *Object) RunDispose() {
	C.g_object_run_dispose((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_set : unsupported parameter ... : varargs

// SetData is a wrapper around the C function g_object_set_data.
func (recv *Object) SetData(key string, data uintptr) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_data := (C.gpointer)(data)

	C.g_object_set_data((*C.GObject)(recv.native), c_key, c_data)

	return
}

// Unsupported : g_object_set_data_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// SetProperty is a wrapper around the C function g_object_set_property.
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

// SetQdata is a wrapper around the C function g_object_set_qdata.
func (recv *Object) SetQdata(quark glib.Quark, data uintptr) {
	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_object_set_qdata((*C.GObject)(recv.native), c_quark, c_data)

	return
}

// Unsupported : g_object_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// StealData is a wrapper around the C function g_object_steal_data.
func (recv *Object) StealData(key string) uintptr {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_object_steal_data((*C.GObject)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// StealQdata is a wrapper around the C function g_object_steal_qdata.
func (recv *Object) StealQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_object_steal_qdata((*C.GObject)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// ThawNotify is a wrapper around the C function g_object_thaw_notify.
func (recv *Object) ThawNotify() {
	C.g_object_thaw_notify((*C.GObject)(recv.native))

	return
}

// Unref is a wrapper around the C function g_object_unref.
func (recv *Object) Unref() {
	C.g_object_unref((C.gpointer)(recv.native))

	return
}

// WatchClosure is a wrapper around the C function g_object_watch_closure.
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
