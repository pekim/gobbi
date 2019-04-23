// This is a generated file - DO NOT EDIT

package gobject

import (
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
/*

	void object_notifyHandler(GObject *, GParamSpec *, gpointer);

	static gulong Object_signal_connect_notify(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "notify", G_CALLBACK(object_notifyHandler), data);
	}

*/
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

// CastToWidget down casts any arbitrary Object to InitiallyUnowned.
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

// CastToWidget down casts any arbitrary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *Object) *Object {
	return ObjectNewFromC(object.ToC())
}

type signalObjectNotifyDetail struct {
	callback  ObjectSignalNotifyCallback
	handlerID C.gulong
}

var signalObjectNotifyId int
var signalObjectNotifyMap = make(map[int]signalObjectNotifyDetail)
var signalObjectNotifyLock sync.RWMutex

// ObjectSignalNotifyCallback is a callback function for a 'notify' signal emitted from a Object.
type ObjectSignalNotifyCallback func(pspec *ParamSpec)

/*
ConnectNotify connects the callback to the 'notify' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectNotify to remove it.
*/
func (recv *Object) ConnectNotify(callback ObjectSignalNotifyCallback) int {
	signalObjectNotifyLock.Lock()
	defer signalObjectNotifyLock.Unlock()

	signalObjectNotifyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_notify(instance, C.gpointer(uintptr(signalObjectNotifyId)))

	detail := signalObjectNotifyDetail{callback, handlerID}
	signalObjectNotifyMap[signalObjectNotifyId] = detail

	return signalObjectNotifyId
}

/*
DisconnectNotify disconnects a callback from the 'notify' signal for the Object.

The connectionID should be a value returned from a call to ConnectNotify.
*/
func (recv *Object) DisconnectNotify(connectionID int) {
	signalObjectNotifyLock.Lock()
	defer signalObjectNotifyLock.Unlock()

	detail, exists := signalObjectNotifyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectNotifyMap, connectionID)
}

//export object_notifyHandler
func object_notifyHandler(_ *C.GObject, c_pspec *C.GParamSpec, data C.gpointer) {
	signalObjectNotifyLock.RLock()
	defer signalObjectNotifyLock.RUnlock()

	pspec := ParamSpecNewFromC(unsafe.Pointer(c_pspec))

	index := int(uintptr(data))
	callback := signalObjectNotifyMap[index].callback
	callback(pspec)
}

// Unsupported : g_object_new : unsupported parameter ... : varargs

// Unsupported : g_object_new_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_object_newv : unsupported parameter parameters :

// g_object_compat_control : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Unsupported : g_object_add_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer (gpointer*) for param weak_pointer_location

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// Blacklisted : g_object_freeze_notify

// Unsupported : g_object_get : unsupported parameter ... : varargs

// Unsupported : g_object_get_data : no return generator

// Blacklisted : g_object_get_property

// Unsupported : g_object_get_qdata : no return generator

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : g_object_notify

// Blacklisted : g_object_ref

// Unsupported : g_object_remove_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer (gpointer*) for param weak_pointer_location

// Blacklisted : g_object_run_dispose

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Unsupported : g_object_set_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_object_set_property

// Unsupported : g_object_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_qdata_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_object_steal_data : no return generator

// Unsupported : g_object_steal_qdata : no return generator

// Blacklisted : g_object_thaw_notify

// Blacklisted : g_object_unref

// Blacklisted : g_object_watch_closure

// Unsupported : g_object_weak_ref : unsupported parameter notify : no type generator for WeakNotify (GWeakNotify) for param notify

// Unsupported : g_object_weak_unref : unsupported parameter notify : no type generator for WeakNotify (GWeakNotify) for param notify

// ParamSpec is a wrapper around the C record GParamSpec.
type ParamSpec struct {
	native *C.GParamSpec
	// g_type_instance : record
	Name      string
	Flags     ParamFlags
	ValueType Type
	OwnerType Type
	// Private : _nick
	// Private : _blurb
	// Private : qdata
	// Private : ref_count
	// Private : param_id
}

func ParamSpecNewFromC(u unsafe.Pointer) *ParamSpec {
	c := (*C.GParamSpec)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpec{
		Flags:     (ParamFlags)(c.flags),
		Name:      C.GoString(c.name),
		OwnerType: (Type)(c.owner_type),
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpec) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.flags =
		(C.GParamFlags)(recv.Flags)
	recv.native.value_type =
		(C.GType)(recv.ValueType)
	recv.native.owner_type =
		(C.GType)(recv.OwnerType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpec with another ParamSpec, and returns true if they represent the same GObject.
func (recv *ParamSpec) Equals(other *ParamSpec) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_param_spec_internal

// Blacklisted : g_param_spec_get_blurb

// Blacklisted : g_param_spec_get_name

// Blacklisted : g_param_spec_get_nick

// Unsupported : g_param_spec_get_qdata : no return generator

// Blacklisted : g_param_spec_ref

// Unsupported : g_param_spec_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_param_spec_sink

// Unsupported : g_param_spec_steal_qdata : no return generator

// Blacklisted : g_param_spec_unref

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
