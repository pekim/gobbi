// This is a generated file - DO NOT EDIT

package gobject

import (
	"sync"
	"unsafe"
)

/*

	void object_notifyHandler(GObject *, GParamSpec *, gpointer);

	static gulong Object_signal_connect_notify(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "notify", G_CALLBACK(object_notifyHandler), data);
	}

*/
import "C"

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

// g_object_compat_control : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Unsupported : g_object_add_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer (gpointer*) for param weak_pointer_location

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// FreezeNotify is a wrapper around the C function g_object_freeze_notify.
func (recv *Object) FreezeNotify() {
	C.g_object_freeze_notify((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_get : unsupported parameter ... : varargs

// Unsupported : g_object_get_data : no return generator

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

// Unsupported : g_object_get_qdata : no return generator

// Unsupported : g_object_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Notify is a wrapper around the C function g_object_notify.
func (recv *Object) Notify(propertyName string) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	C.g_object_notify((*C.GObject)(recv.native), c_property_name)

	return
}

// Ref is a wrapper around the C function g_object_ref.
func (recv *Object) Ref() Object {
	retC := C.g_object_ref((C.gpointer)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_remove_weak_pointer : unsupported parameter weak_pointer_location : no type generator for gpointer (gpointer*) for param weak_pointer_location

// RunDispose is a wrapper around the C function g_object_run_dispose.
func (recv *Object) RunDispose() {
	C.g_object_run_dispose((*C.GObject)(recv.native))

	return
}

// Unsupported : g_object_set : unsupported parameter ... : varargs

// Unsupported : g_object_set_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

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

// Unsupported : g_object_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_qdata_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_object_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_object_steal_data : no return generator

// Unsupported : g_object_steal_qdata : no return generator

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

// Equals compares this ParamSpec with another ParamSpec, and returns true if they represent the same GObject.
func (recv *ParamSpec) Equals(other *ParamSpec) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecInternal is a wrapper around the C function g_param_spec_internal.
func ParamSpecInternal(paramType Type, name string, nick string, blurb string, flags ParamFlags) ParamSpec {
	c_param_type := (C.GType)(paramType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_internal(c_param_type, c_name, c_nick, c_blurb, c_flags)
	retGo := *ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBlurb is a wrapper around the C function g_param_spec_get_blurb.
func (recv *ParamSpec) GetBlurb() string {
	retC := C.g_param_spec_get_blurb((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_param_spec_get_name.
func (recv *ParamSpec) GetName() string {
	retC := C.g_param_spec_get_name((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNick is a wrapper around the C function g_param_spec_get_nick.
func (recv *ParamSpec) GetNick() string {
	retC := C.g_param_spec_get_nick((*C.GParamSpec)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_param_spec_get_qdata : no return generator

// Ref is a wrapper around the C function g_param_spec_ref.
func (recv *ParamSpec) Ref() *ParamSpec {
	retC := C.g_param_spec_ref((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_set_qdata : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Sink is a wrapper around the C function g_param_spec_sink.
func (recv *ParamSpec) Sink() {
	C.g_param_spec_sink((*C.GParamSpec)(recv.native))

	return
}

// Unsupported : g_param_spec_steal_qdata : no return generator

// Unref is a wrapper around the C function g_param_spec_unref.
func (recv *ParamSpec) Unref() {
	C.g_param_spec_unref((*C.GParamSpec)(recv.native))

	return
}

// Equals compares this ParamSpecBoolean with another ParamSpecBoolean, and returns true if they represent the same GObject.
func (recv *ParamSpecBoolean) Equals(other *ParamSpecBoolean) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoolean) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecBoxed with another ParamSpecBoxed, and returns true if they represent the same GObject.
func (recv *ParamSpecBoxed) Equals(other *ParamSpecBoxed) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoxed) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecChar with another ParamSpecChar, and returns true if they represent the same GObject.
func (recv *ParamSpecChar) Equals(other *ParamSpecChar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecDouble with another ParamSpecDouble, and returns true if they represent the same GObject.
func (recv *ParamSpecDouble) Equals(other *ParamSpecDouble) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecDouble) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecEnum with another ParamSpecEnum, and returns true if they represent the same GObject.
func (recv *ParamSpecEnum) Equals(other *ParamSpecEnum) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecEnum) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecFlags with another ParamSpecFlags, and returns true if they represent the same GObject.
func (recv *ParamSpecFlags) Equals(other *ParamSpecFlags) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFlags) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecFloat with another ParamSpecFloat, and returns true if they represent the same GObject.
func (recv *ParamSpecFloat) Equals(other *ParamSpecFloat) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFloat) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecInt with another ParamSpecInt, and returns true if they represent the same GObject.
func (recv *ParamSpecInt) Equals(other *ParamSpecInt) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecInt64 with another ParamSpecInt64, and returns true if they represent the same GObject.
func (recv *ParamSpecInt64) Equals(other *ParamSpecInt64) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecLong with another ParamSpecLong, and returns true if they represent the same GObject.
func (recv *ParamSpecLong) Equals(other *ParamSpecLong) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecLong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecObject with another ParamSpecObject, and returns true if they represent the same GObject.
func (recv *ParamSpecObject) Equals(other *ParamSpecObject) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecObject) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecParam with another ParamSpecParam, and returns true if they represent the same GObject.
func (recv *ParamSpecParam) Equals(other *ParamSpecParam) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecParam) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecPointer with another ParamSpecPointer, and returns true if they represent the same GObject.
func (recv *ParamSpecPointer) Equals(other *ParamSpecPointer) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecPointer) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecString with another ParamSpecString, and returns true if they represent the same GObject.
func (recv *ParamSpecString) Equals(other *ParamSpecString) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecString) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecUChar with another ParamSpecUChar, and returns true if they represent the same GObject.
func (recv *ParamSpecUChar) Equals(other *ParamSpecUChar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecUInt with another ParamSpecUInt, and returns true if they represent the same GObject.
func (recv *ParamSpecUInt) Equals(other *ParamSpecUInt) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecUInt64 with another ParamSpecUInt64, and returns true if they represent the same GObject.
func (recv *ParamSpecUInt64) Equals(other *ParamSpecUInt64) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecULong with another ParamSpecULong, and returns true if they represent the same GObject.
func (recv *ParamSpecULong) Equals(other *ParamSpecULong) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecULong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecUnichar with another ParamSpecUnichar, and returns true if they represent the same GObject.
func (recv *ParamSpecUnichar) Equals(other *ParamSpecUnichar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUnichar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Equals compares this ParamSpecValueArray with another ParamSpecValueArray, and returns true if they represent the same GObject.
func (recv *ParamSpecValueArray) Equals(other *ParamSpecValueArray) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecValueArray) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// Blacklisted : GTypeModule
