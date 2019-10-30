// Code generated - DO NOT EDIT.
// +build gobject_2.18

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"reflect"
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

var gobjectClassGoTypeMap = make(map[string]reflect.Type)

// Unsupported : alias has no type generator for ClosureMarshal, GClosureMarshal

// Unsupported : alias has no type generator for VaClosureMarshal, GVaClosureMarshal

// Type is a representation of the C alias GType.
type Type uint64

type ConnectFlags C.GConnectFlags

const (
	CONNECT_AFTER   ConnectFlags = 1
	CONNECT_SWAPPED ConnectFlags = 2
)

type ParamFlags C.GParamFlags

const (
	PARAM_READABLE        ParamFlags = 1
	PARAM_WRITABLE        ParamFlags = 2
	PARAM_READWRITE       ParamFlags = 3
	PARAM_CONSTRUCT       ParamFlags = 4
	PARAM_CONSTRUCT_ONLY  ParamFlags = 8
	PARAM_LAX_VALIDATION  ParamFlags = 16
	PARAM_STATIC_NAME     ParamFlags = 32
	PARAM_PRIVATE         ParamFlags = 32
	PARAM_STATIC_NICK     ParamFlags = 64
	PARAM_STATIC_BLURB    ParamFlags = 128
	PARAM_EXPLICIT_NOTIFY ParamFlags = 1073741824

// Blacklisted member : G_PARAM_DEPRECATED
)

type SignalFlags C.GSignalFlags

const (
	SIGNAL_RUN_FIRST    SignalFlags = 1
	SIGNAL_RUN_LAST     SignalFlags = 2
	SIGNAL_RUN_CLEANUP  SignalFlags = 4
	SIGNAL_NO_RECURSE   SignalFlags = 8
	SIGNAL_DETAILED     SignalFlags = 16
	SIGNAL_ACTION       SignalFlags = 32
	SIGNAL_NO_HOOKS     SignalFlags = 64
	SIGNAL_MUST_COLLECT SignalFlags = 128
	SIGNAL_DEPRECATED   SignalFlags = 256
)

type SignalMatchType C.GSignalMatchType

const (
	SIGNAL_MATCH_ID        SignalMatchType = 1
	SIGNAL_MATCH_DETAIL    SignalMatchType = 2
	SIGNAL_MATCH_CLOSURE   SignalMatchType = 4
	SIGNAL_MATCH_FUNC      SignalMatchType = 8
	SIGNAL_MATCH_DATA      SignalMatchType = 16
	SIGNAL_MATCH_UNBLOCKED SignalMatchType = 32
)

type TypeDebugFlags C.GTypeDebugFlags

const (
	TYPE_DEBUG_NONE           TypeDebugFlags = 0
	TYPE_DEBUG_OBJECTS        TypeDebugFlags = 1
	TYPE_DEBUG_SIGNALS        TypeDebugFlags = 2
	TYPE_DEBUG_INSTANCE_COUNT TypeDebugFlags = 4
	TYPE_DEBUG_MASK           TypeDebugFlags = 7
)

type TypeFlags C.GTypeFlags

const (
	TYPE_FLAG_ABSTRACT       TypeFlags = 16
	TYPE_FLAG_VALUE_ABSTRACT TypeFlags = 32
)

type TypeFundamentalFlags C.GTypeFundamentalFlags

const (
	TYPE_FLAG_CLASSED        TypeFundamentalFlags = 1
	TYPE_FLAG_INSTANTIATABLE TypeFundamentalFlags = 2
	TYPE_FLAG_DERIVABLE      TypeFundamentalFlags = 4
	TYPE_FLAG_DEEP_DERIVABLE TypeFundamentalFlags = 8
)

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

// ObjectCompatControl is a wrapper around the C function g_object_compat_control.
func ObjectCompatControl(what uint64, data uintptr) uint64 {
	c_what := (C.gsize)(what)

	c_data := (C.gpointer)(data)

	retC := C.g_object_compat_control(c_what, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// ObjectInterfaceFindProperty is a wrapper around the C function g_object_interface_find_property.
func ObjectInterfaceFindProperty(gIface *TypeInterface, propertyName string) *ParamSpec {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_object_interface_find_property(c_g_iface, c_property_name)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInterfaceInstallProperty is a wrapper around the C function g_object_interface_install_property.
func ObjectInterfaceInstallProperty(gIface *TypeInterface, pspec *ParamSpec) {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_object_interface_install_property(c_g_iface, c_pspec)

	return
}

// g_object_interface_list_properties : array return type :
// Unsupported : g_object_add_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify (GToggleNotify) for param notify

// AddWeakPointer is a wrapper around the C function g_object_add_weak_pointer.
func (recv *Object) AddWeakPointer(weakPointerLocation uintptr) {
	c_weak_pointer_location := (C.gpointer)(weakPointerLocation)

	C.g_object_add_weak_pointer((*C.GObject)(recv.native), &c_weak_pointer_location)

	return
}

// Unsupported : g_object_connect : unsupported parameter ... : varargs

// Unsupported : g_object_disconnect : unsupported parameter ... : varargs

// ForceFloating is a wrapper around the C function g_object_force_floating.
func (recv *Object) ForceFloating() {
	C.g_object_force_floating((*C.GObject)(recv.native))

	return
}

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

// IsFloating is a wrapper around the C function g_object_is_floating.
func (recv *Object) IsFloating() bool {
	retC := C.g_object_is_floating((C.gpointer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// RefSink is a wrapper around the C function g_object_ref_sink.
func (recv *Object) RefSink() Object {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_object_remove_toggle_ref : unsupported parameter notify : no type generator for ToggleNotify (GToggleNotify) for param notify

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

// GetQdata is a wrapper around the C function g_param_spec_get_qdata.
func (recv *ParamSpec) GetQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_param_spec_get_qdata((*C.GParamSpec)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetRedirectTarget is a wrapper around the C function g_param_spec_get_redirect_target.
func (recv *ParamSpec) GetRedirectTarget() *ParamSpec {
	retC := C.g_param_spec_get_redirect_target((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_param_spec_ref.
func (recv *ParamSpec) Ref() *ParamSpec {
	retC := C.g_param_spec_ref((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefSink is a wrapper around the C function g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	retC := C.g_param_spec_ref_sink((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetQdata is a wrapper around the C function g_param_spec_set_qdata.
func (recv *ParamSpec) SetQdata(quark glib.Quark, data uintptr) {
	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_param_spec_set_qdata((*C.GParamSpec)(recv.native), c_quark, c_data)

	return
}

// Unsupported : g_param_spec_set_qdata_full : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Sink is a wrapper around the C function g_param_spec_sink.
func (recv *ParamSpec) Sink() {
	C.g_param_spec_sink((*C.GParamSpec)(recv.native))

	return
}

// StealQdata is a wrapper around the C function g_param_spec_steal_qdata.
func (recv *ParamSpec) StealQdata(quark glib.Quark) uintptr {
	c_quark := (C.GQuark)(quark)

	retC := C.g_param_spec_steal_qdata((*C.GParamSpec)(recv.native), c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_param_spec_unref.
func (recv *ParamSpec) Unref() {
	C.g_param_spec_unref((*C.GParamSpec)(recv.native))

	return
}

// ParamSpecBoolean is a wrapper around the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native *C.GParamSpecBoolean
	// parent_instance : record
	DefaultValue bool
}

func ParamSpecBooleanNewFromC(u unsafe.Pointer) *ParamSpecBoolean {
	c := (*C.GParamSpecBoolean)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecBoolean{
		DefaultValue: c.default_value == C.TRUE,
		native:       c,
	}

	return g
}

func (recv *ParamSpecBoolean) ToC() unsafe.Pointer {
	recv.native.default_value =
		boolToGboolean(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecBoolean with another ParamSpecBoolean, and returns true if they represent the same GObject.
func (recv *ParamSpecBoolean) Equals(other *ParamSpecBoolean) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoolean) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecBoxed is a wrapper around the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native *C.GParamSpecBoxed
	// parent_instance : record
}

func ParamSpecBoxedNewFromC(u unsafe.Pointer) *ParamSpecBoxed {
	c := (*C.GParamSpecBoxed)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecBoxed{native: c}

	return g
}

func (recv *ParamSpecBoxed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecBoxed with another ParamSpecBoxed, and returns true if they represent the same GObject.
func (recv *ParamSpecBoxed) Equals(other *ParamSpecBoxed) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecBoxed) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecChar is a wrapper around the C record GParamSpecChar.
type ParamSpecChar struct {
	native *C.GParamSpecChar
	// parent_instance : record
	Minimum      int8
	Maximum      int8
	DefaultValue int8
}

func ParamSpecCharNewFromC(u unsafe.Pointer) *ParamSpecChar {
	c := (*C.GParamSpecChar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecChar{
		DefaultValue: (int8)(c.default_value),
		Maximum:      (int8)(c.maximum),
		Minimum:      (int8)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecChar) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint8)(recv.Minimum)
	recv.native.maximum =
		(C.gint8)(recv.Maximum)
	recv.native.default_value =
		(C.gint8)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecChar with another ParamSpecChar, and returns true if they represent the same GObject.
func (recv *ParamSpecChar) Equals(other *ParamSpecChar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecDouble is a wrapper around the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native *C.GParamSpecDouble
	// parent_instance : record
	Minimum      float64
	Maximum      float64
	DefaultValue float64
	Epsilon      float64
}

func ParamSpecDoubleNewFromC(u unsafe.Pointer) *ParamSpecDouble {
	c := (*C.GParamSpecDouble)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecDouble{
		DefaultValue: (float64)(c.default_value),
		Epsilon:      (float64)(c.epsilon),
		Maximum:      (float64)(c.maximum),
		Minimum:      (float64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecDouble) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gdouble)(recv.Minimum)
	recv.native.maximum =
		(C.gdouble)(recv.Maximum)
	recv.native.default_value =
		(C.gdouble)(recv.DefaultValue)
	recv.native.epsilon =
		(C.gdouble)(recv.Epsilon)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecDouble with another ParamSpecDouble, and returns true if they represent the same GObject.
func (recv *ParamSpecDouble) Equals(other *ParamSpecDouble) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecDouble) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecEnum is a wrapper around the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native *C.GParamSpecEnum
	// parent_instance : record
	// enum_class : record
	DefaultValue int32
}

func ParamSpecEnumNewFromC(u unsafe.Pointer) *ParamSpecEnum {
	c := (*C.GParamSpecEnum)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecEnum{
		DefaultValue: (int32)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecEnum) ToC() unsafe.Pointer {
	recv.native.default_value =
		(C.gint)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecEnum with another ParamSpecEnum, and returns true if they represent the same GObject.
func (recv *ParamSpecEnum) Equals(other *ParamSpecEnum) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecEnum) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecFlags is a wrapper around the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native *C.GParamSpecFlags
	// parent_instance : record
	// flags_class : record
	DefaultValue uint32
}

func ParamSpecFlagsNewFromC(u unsafe.Pointer) *ParamSpecFlags {
	c := (*C.GParamSpecFlags)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecFlags{
		DefaultValue: (uint32)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecFlags) ToC() unsafe.Pointer {
	recv.native.default_value =
		(C.guint)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecFlags with another ParamSpecFlags, and returns true if they represent the same GObject.
func (recv *ParamSpecFlags) Equals(other *ParamSpecFlags) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFlags) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecFloat is a wrapper around the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native *C.GParamSpecFloat
	// parent_instance : record
	Minimum      float32
	Maximum      float32
	DefaultValue float32
	Epsilon      float32
}

func ParamSpecFloatNewFromC(u unsafe.Pointer) *ParamSpecFloat {
	c := (*C.GParamSpecFloat)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecFloat{
		DefaultValue: (float32)(c.default_value),
		Epsilon:      (float32)(c.epsilon),
		Maximum:      (float32)(c.maximum),
		Minimum:      (float32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecFloat) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gfloat)(recv.Minimum)
	recv.native.maximum =
		(C.gfloat)(recv.Maximum)
	recv.native.default_value =
		(C.gfloat)(recv.DefaultValue)
	recv.native.epsilon =
		(C.gfloat)(recv.Epsilon)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecFloat with another ParamSpecFloat, and returns true if they represent the same GObject.
func (recv *ParamSpecFloat) Equals(other *ParamSpecFloat) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecFloat) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecGType is a wrapper around the C record GParamSpecGType.
type ParamSpecGType struct {
	native *C.GParamSpecGType
	// parent_instance : record
	IsAType Type
}

func ParamSpecGTypeNewFromC(u unsafe.Pointer) *ParamSpecGType {
	c := (*C.GParamSpecGType)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecGType{
		IsAType: (Type)(c.is_a_type),
		native:  c,
	}

	return g
}

func (recv *ParamSpecGType) ToC() unsafe.Pointer {
	recv.native.is_a_type =
		(C.GType)(recv.IsAType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecGType with another ParamSpecGType, and returns true if they represent the same GObject.
func (recv *ParamSpecGType) Equals(other *ParamSpecGType) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecGType) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecInt is a wrapper around the C record GParamSpecInt.
type ParamSpecInt struct {
	native *C.GParamSpecInt
	// parent_instance : record
	Minimum      int32
	Maximum      int32
	DefaultValue int32
}

func ParamSpecIntNewFromC(u unsafe.Pointer) *ParamSpecInt {
	c := (*C.GParamSpecInt)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecInt{
		DefaultValue: (int32)(c.default_value),
		Maximum:      (int32)(c.maximum),
		Minimum:      (int32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecInt) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.default_value =
		(C.gint)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecInt with another ParamSpecInt, and returns true if they represent the same GObject.
func (recv *ParamSpecInt) Equals(other *ParamSpecInt) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecInt64 is a wrapper around the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native *C.GParamSpecInt64
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func ParamSpecInt64NewFromC(u unsafe.Pointer) *ParamSpecInt64 {
	c := (*C.GParamSpecInt64)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecInt64{
		DefaultValue: (int64)(c.default_value),
		Maximum:      (int64)(c.maximum),
		Minimum:      (int64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecInt64) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint64)(recv.Minimum)
	recv.native.maximum =
		(C.gint64)(recv.Maximum)
	recv.native.default_value =
		(C.gint64)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecInt64 with another ParamSpecInt64, and returns true if they represent the same GObject.
func (recv *ParamSpecInt64) Equals(other *ParamSpecInt64) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecLong is a wrapper around the C record GParamSpecLong.
type ParamSpecLong struct {
	native *C.GParamSpecLong
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

func ParamSpecLongNewFromC(u unsafe.Pointer) *ParamSpecLong {
	c := (*C.GParamSpecLong)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecLong{
		DefaultValue: (int64)(c.default_value),
		Maximum:      (int64)(c.maximum),
		Minimum:      (int64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecLong) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.glong)(recv.Minimum)
	recv.native.maximum =
		(C.glong)(recv.Maximum)
	recv.native.default_value =
		(C.glong)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecLong with another ParamSpecLong, and returns true if they represent the same GObject.
func (recv *ParamSpecLong) Equals(other *ParamSpecLong) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecLong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecObject is a wrapper around the C record GParamSpecObject.
type ParamSpecObject struct {
	native *C.GParamSpecObject
	// parent_instance : record
}

func ParamSpecObjectNewFromC(u unsafe.Pointer) *ParamSpecObject {
	c := (*C.GParamSpecObject)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecObject{native: c}

	return g
}

func (recv *ParamSpecObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecObject with another ParamSpecObject, and returns true if they represent the same GObject.
func (recv *ParamSpecObject) Equals(other *ParamSpecObject) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecObject) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecOverride is a wrapper around the C record GParamSpecOverride.
type ParamSpecOverride struct {
	native *C.GParamSpecOverride
	// Private : parent_instance
	// Private : overridden
}

func ParamSpecOverrideNewFromC(u unsafe.Pointer) *ParamSpecOverride {
	c := (*C.GParamSpecOverride)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecOverride{native: c}

	return g
}

func (recv *ParamSpecOverride) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecOverride with another ParamSpecOverride, and returns true if they represent the same GObject.
func (recv *ParamSpecOverride) Equals(other *ParamSpecOverride) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecOverride) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecParam is a wrapper around the C record GParamSpecParam.
type ParamSpecParam struct {
	native *C.GParamSpecParam
	// parent_instance : record
}

func ParamSpecParamNewFromC(u unsafe.Pointer) *ParamSpecParam {
	c := (*C.GParamSpecParam)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecParam{native: c}

	return g
}

func (recv *ParamSpecParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecParam with another ParamSpecParam, and returns true if they represent the same GObject.
func (recv *ParamSpecParam) Equals(other *ParamSpecParam) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecParam) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecPointer is a wrapper around the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native *C.GParamSpecPointer
	// parent_instance : record
}

func ParamSpecPointerNewFromC(u unsafe.Pointer) *ParamSpecPointer {
	c := (*C.GParamSpecPointer)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPointer{native: c}

	return g
}

func (recv *ParamSpecPointer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecPointer with another ParamSpecPointer, and returns true if they represent the same GObject.
func (recv *ParamSpecPointer) Equals(other *ParamSpecPointer) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecPointer) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecString is a wrapper around the C record GParamSpecString.
type ParamSpecString struct {
	native *C.GParamSpecString
	// parent_instance : record
	DefaultValue string
	CsetFirst    string
	CsetNth      string
	Substitutor  rune
	// Bitfield not supported :  1 null_fold_if_empty
	// Bitfield not supported :  1 ensure_non_null
}

func ParamSpecStringNewFromC(u unsafe.Pointer) *ParamSpecString {
	c := (*C.GParamSpecString)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecString{
		CsetFirst:    C.GoString(c.cset_first),
		CsetNth:      C.GoString(c.cset_nth),
		DefaultValue: C.GoString(c.default_value),
		Substitutor:  (rune)(c.substitutor),
		native:       c,
	}

	return g
}

func (recv *ParamSpecString) ToC() unsafe.Pointer {
	recv.native.default_value =
		C.CString(recv.DefaultValue)
	recv.native.cset_first =
		C.CString(recv.CsetFirst)
	recv.native.cset_nth =
		C.CString(recv.CsetNth)
	recv.native.substitutor =
		(C.gchar)(recv.Substitutor)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecString with another ParamSpecString, and returns true if they represent the same GObject.
func (recv *ParamSpecString) Equals(other *ParamSpecString) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecString) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecUChar is a wrapper around the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native *C.GParamSpecUChar
	// parent_instance : record
	Minimum      uint8
	Maximum      uint8
	DefaultValue uint8
}

func ParamSpecUCharNewFromC(u unsafe.Pointer) *ParamSpecUChar {
	c := (*C.GParamSpecUChar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUChar{
		DefaultValue: (uint8)(c.default_value),
		Maximum:      (uint8)(c.maximum),
		Minimum:      (uint8)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUChar) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.guint8)(recv.Minimum)
	recv.native.maximum =
		(C.guint8)(recv.Maximum)
	recv.native.default_value =
		(C.guint8)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecUChar with another ParamSpecUChar, and returns true if they represent the same GObject.
func (recv *ParamSpecUChar) Equals(other *ParamSpecUChar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUChar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecUInt is a wrapper around the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native *C.GParamSpecUInt
	// parent_instance : record
	Minimum      uint32
	Maximum      uint32
	DefaultValue uint32
}

func ParamSpecUIntNewFromC(u unsafe.Pointer) *ParamSpecUInt {
	c := (*C.GParamSpecUInt)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUInt{
		DefaultValue: (uint32)(c.default_value),
		Maximum:      (uint32)(c.maximum),
		Minimum:      (uint32)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUInt) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.guint)(recv.Minimum)
	recv.native.maximum =
		(C.guint)(recv.Maximum)
	recv.native.default_value =
		(C.guint)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecUInt with another ParamSpecUInt, and returns true if they represent the same GObject.
func (recv *ParamSpecUInt) Equals(other *ParamSpecUInt) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecUInt64 is a wrapper around the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native *C.GParamSpecUInt64
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func ParamSpecUInt64NewFromC(u unsafe.Pointer) *ParamSpecUInt64 {
	c := (*C.GParamSpecUInt64)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUInt64{
		DefaultValue: (uint64)(c.default_value),
		Maximum:      (uint64)(c.maximum),
		Minimum:      (uint64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUInt64) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.guint64)(recv.Minimum)
	recv.native.maximum =
		(C.guint64)(recv.Maximum)
	recv.native.default_value =
		(C.guint64)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecUInt64 with another ParamSpecUInt64, and returns true if they represent the same GObject.
func (recv *ParamSpecUInt64) Equals(other *ParamSpecUInt64) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUInt64) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecULong is a wrapper around the C record GParamSpecULong.
type ParamSpecULong struct {
	native *C.GParamSpecULong
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

func ParamSpecULongNewFromC(u unsafe.Pointer) *ParamSpecULong {
	c := (*C.GParamSpecULong)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecULong{
		DefaultValue: (uint64)(c.default_value),
		Maximum:      (uint64)(c.maximum),
		Minimum:      (uint64)(c.minimum),
		native:       c,
	}

	return g
}

func (recv *ParamSpecULong) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gulong)(recv.Minimum)
	recv.native.maximum =
		(C.gulong)(recv.Maximum)
	recv.native.default_value =
		(C.gulong)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecULong with another ParamSpecULong, and returns true if they represent the same GObject.
func (recv *ParamSpecULong) Equals(other *ParamSpecULong) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecULong) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecUnichar is a wrapper around the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native *C.GParamSpecUnichar
	// parent_instance : record
	DefaultValue rune
}

func ParamSpecUnicharNewFromC(u unsafe.Pointer) *ParamSpecUnichar {
	c := (*C.GParamSpecUnichar)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecUnichar{
		DefaultValue: (rune)(c.default_value),
		native:       c,
	}

	return g
}

func (recv *ParamSpecUnichar) ToC() unsafe.Pointer {
	recv.native.default_value =
		(C.gunichar)(recv.DefaultValue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecUnichar with another ParamSpecUnichar, and returns true if they represent the same GObject.
func (recv *ParamSpecUnichar) Equals(other *ParamSpecUnichar) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecUnichar) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}

// ParamSpecValueArray is a wrapper around the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native *C.GParamSpecValueArray
	// parent_instance : record
	// element_spec : record
	FixedNElements uint32
}

func ParamSpecValueArrayNewFromC(u unsafe.Pointer) *ParamSpecValueArray {
	c := (*C.GParamSpecValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecValueArray{
		FixedNElements: (uint32)(c.fixed_n_elements),
		native:         c,
	}

	return g
}

func (recv *ParamSpecValueArray) ToC() unsafe.Pointer {
	recv.native.fixed_n_elements =
		(C.guint)(recv.FixedNElements)

	return (unsafe.Pointer)(recv.native)
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

const PARAM_MASK int32 = C.G_PARAM_MASK
const PARAM_STATIC_STRINGS int32 = C.G_PARAM_STATIC_STRINGS
const PARAM_USER_SHIFT int32 = C.G_PARAM_USER_SHIFT
const SIGNAL_FLAGS_MASK int32 = C.G_SIGNAL_FLAGS_MASK
const SIGNAL_MATCH_MASK int32 = C.G_SIGNAL_MATCH_MASK

// Unsupported : type GLib.Type for TYPE_FLAG_RESERVED_ID_BIT

const TYPE_FUNDAMENTAL_MAX int32 = C.G_TYPE_FUNDAMENTAL_MAX
const TYPE_FUNDAMENTAL_SHIFT int32 = C.G_TYPE_FUNDAMENTAL_SHIFT
const TYPE_RESERVED_BSE_FIRST int32 = C.G_TYPE_RESERVED_BSE_FIRST
const TYPE_RESERVED_BSE_LAST int32 = C.G_TYPE_RESERVED_BSE_LAST
const TYPE_RESERVED_GLIB_FIRST int32 = C.G_TYPE_RESERVED_GLIB_FIRST
const TYPE_RESERVED_GLIB_LAST int32 = C.G_TYPE_RESERVED_GLIB_LAST
const TYPE_RESERVED_USER_FIRST int32 = C.G_TYPE_RESERVED_USER_FIRST

// Blacklisted : VALUE_COLLECT_FORMAT_MAX_LENGTH

const VALUE_NOCOPY_CONTENTS int32 = C.G_VALUE_NOCOPY_CONTENTS

// BoxedCopy is a wrapper around the C function g_boxed_copy.
func BoxedCopy(boxedType Type, srcBoxed uintptr) uintptr {
	c_boxed_type := (C.GType)(boxedType)

	c_src_boxed := (C.gconstpointer)(srcBoxed)

	retC := C.g_boxed_copy(c_boxed_type, c_src_boxed)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// BoxedFree is a wrapper around the C function g_boxed_free.
func BoxedFree(boxedType Type, boxed uintptr) {
	c_boxed_type := (C.GType)(boxedType)

	c_boxed := (C.gpointer)(boxed)

	C.g_boxed_free(c_boxed_type, c_boxed)

	return
}

// Unsupported : g_boxed_type_register_static : unsupported parameter boxed_copy : no type generator for BoxedCopyFunc (GBoxedCopyFunc) for param boxed_copy

// Unsupported : g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// Unsupported : g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func

// EnumCompleteTypeInfo is a wrapper around the C function g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType Type, constValues *EnumValue) *TypeInfo {
	c_g_enum_type := (C.GType)(gEnumType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GEnumValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GEnumValue)(constValues.ToC())
	}

	C.g_enum_complete_type_info(c_g_enum_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// EnumGetValue is a wrapper around the C function g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_value := (C.gint)(value)

	retC := C.g_enum_get_value(c_enum_class, c_value)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByName is a wrapper around the C function g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_enum_get_value_by_name(c_enum_class, c_name)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumGetValueByNick is a wrapper around the C function g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	c_enum_class := (*C.GEnumClass)(C.NULL)
	if enumClass != nil {
		c_enum_class = (*C.GEnumClass)(enumClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_enum_get_value_by_nick(c_enum_class, c_nick)
	retGo := EnumValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumRegisterStatic is a wrapper around the C function g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GEnumValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GEnumValue)(constStaticValues.ToC())
	}

	retC := C.g_enum_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// FlagsCompleteTypeInfo is a wrapper around the C function g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType Type, constValues *FlagsValue) *TypeInfo {
	c_g_flags_type := (C.GType)(gFlagsType)

	var c_info C.GTypeInfo

	c_const_values := (*C.GFlagsValue)(C.NULL)
	if constValues != nil {
		c_const_values = (*C.GFlagsValue)(constValues.ToC())
	}

	C.g_flags_complete_type_info(c_g_flags_type, &c_info, c_const_values)

	info := TypeInfoNewFromC(unsafe.Pointer(&c_info))

	return info
}

// FlagsGetFirstValue is a wrapper around the C function g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_value := (C.guint)(value)

	retC := C.g_flags_get_first_value(c_flags_class, c_value)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByName is a wrapper around the C function g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_flags_get_value_by_name(c_flags_class, c_name)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsGetValueByNick is a wrapper around the C function g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	c_flags_class := (*C.GFlagsClass)(C.NULL)
	if flagsClass != nil {
		c_flags_class = (*C.GFlagsClass)(flagsClass.ToC())
	}

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	retC := C.g_flags_get_value_by_nick(c_flags_class, c_nick)
	retGo := FlagsValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlagsRegisterStatic is a wrapper around the C function g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_const_static_values := (*C.GFlagsValue)(C.NULL)
	if constStaticValues != nil {
		c_const_static_values = (*C.GFlagsValue)(constStaticValues.ToC())
	}

	retC := C.g_flags_register_static(c_name, c_const_static_values)
	retGo := (Type)(retC)

	return retGo
}

// GtypeGetType is a wrapper around the C function g_gtype_get_type.
func GtypeGetType() Type {
	retC := C.g_gtype_get_type()
	retGo := (Type)(retC)

	return retGo
}

// ParamSpecBoolean_ is a wrapper around the C function g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value :=
		boolToGboolean(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_boolean(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecBoxed_ is a wrapper around the C function g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_boxed_type := (C.GType)(boxedType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_boxed(c_name, c_nick, c_blurb, c_boxed_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecChar_ is a wrapper around the C function g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint8)(minimum)

	c_maximum := (C.gint8)(maximum)

	c_default_value := (C.gint8)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_char(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecDouble_ is a wrapper around the C function g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gdouble)(minimum)

	c_maximum := (C.gdouble)(maximum)

	c_default_value := (C.gdouble)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_double(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecEnum_ is a wrapper around the C function g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType Type, defaultValue int32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_enum_type := (C.GType)(enumType)

	c_default_value := (C.gint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_enum(c_name, c_nick, c_blurb, c_enum_type, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecFlags_ is a wrapper around the C function g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType Type, defaultValue uint32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_flags_type := (C.GType)(flagsType)

	c_default_value := (C.guint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_flags(c_name, c_nick, c_blurb, c_flags_type, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecFloat_ is a wrapper around the C function g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gfloat)(minimum)

	c_maximum := (C.gfloat)(maximum)

	c_default_value := (C.gfloat)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_float(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecGtype is a wrapper around the C function g_param_spec_gtype.
func ParamSpecGtype(name string, nick string, blurb string, isAType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_is_a_type := (C.GType)(isAType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_gtype(c_name, c_nick, c_blurb, c_is_a_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecInt_ is a wrapper around the C function g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int32, maximum int32, defaultValue int32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint)(minimum)

	c_maximum := (C.gint)(maximum)

	c_default_value := (C.gint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_int(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecInt64_ is a wrapper around the C function g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gint64)(minimum)

	c_maximum := (C.gint64)(maximum)

	c_default_value := (C.gint64)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_int64(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamLong_ is a wrapper around the C function g_param_spec_long.
func ParamLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.glong)(minimum)

	c_maximum := (C.glong)(maximum)

	c_default_value := (C.glong)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_long(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecObject_ is a wrapper around the C function g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_object_type := (C.GType)(objectType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_object(c_name, c_nick, c_blurb, c_object_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecOverride_ is a wrapper around the C function g_param_spec_override.
func ParamSpecOverride_(name string, overridden *ParamSpec) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_overridden := (*C.GParamSpec)(C.NULL)
	if overridden != nil {
		c_overridden = (*C.GParamSpec)(overridden.ToC())
	}

	retC := C.g_param_spec_override(c_name, c_overridden)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecParam_ is a wrapper around the C function g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType Type, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_param_type := (C.GType)(paramType)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_param(c_name, c_nick, c_blurb, c_param_type, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecPointer_ is a wrapper around the C function g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_pointer(c_name, c_nick, c_blurb, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecString_ is a wrapper around the C function g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value := C.CString(defaultValue)
	defer C.free(unsafe.Pointer(c_default_value))

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_string(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUchar is a wrapper around the C function g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint8)(minimum)

	c_maximum := (C.guint8)(maximum)

	c_default_value := (C.guint8)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uchar(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUint is a wrapper around the C function g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint32, maximum uint32, defaultValue uint32, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint)(minimum)

	c_maximum := (C.guint)(maximum)

	c_default_value := (C.guint)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uint(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUint64 is a wrapper around the C function g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.guint64)(minimum)

	c_maximum := (C.guint64)(maximum)

	c_default_value := (C.guint64)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_uint64(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUlong is a wrapper around the C function g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_minimum := (C.gulong)(minimum)

	c_maximum := (C.gulong)(maximum)

	c_default_value := (C.gulong)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_ulong(c_name, c_nick, c_blurb, c_minimum, c_maximum, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecUnichar_ is a wrapper around the C function g_param_spec_unichar.
func ParamSpecUnichar_(name string, nick string, blurb string, defaultValue rune, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_default_value := (C.gunichar)(defaultValue)

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_unichar(c_name, c_nick, c_blurb, c_default_value, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamSpecValueArray_ is a wrapper around the C function g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags ParamFlags) *ParamSpec {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_nick := C.CString(nick)
	defer C.free(unsafe.Pointer(c_nick))

	c_blurb := C.CString(blurb)
	defer C.free(unsafe.Pointer(c_blurb))

	c_element_spec := (*C.GParamSpec)(C.NULL)
	if elementSpec != nil {
		c_element_spec = (*C.GParamSpec)(elementSpec.ToC())
	}

	c_flags := (C.GParamFlags)(flags)

	retC := C.g_param_spec_value_array(c_name, c_nick, c_blurb, c_element_spec, c_flags)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParamTypeRegisterStatic is a wrapper around the C function g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_pspec_info := (*C.GParamSpecTypeInfo)(C.NULL)
	if pspecInfo != nil {
		c_pspec_info = (*C.GParamSpecTypeInfo)(pspecInfo.ToC())
	}

	retC := C.g_param_type_register_static(c_name, c_pspec_info)
	retGo := (Type)(retC)

	return retGo
}

// ParamValueConvert is a wrapper around the C function g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_src_value := (*C.GValue)(C.NULL)
	if srcValue != nil {
		c_src_value = (*C.GValue)(srcValue.ToC())
	}

	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	c_strict_validation :=
		boolToGboolean(strictValidation)

	retC := C.g_param_value_convert(c_pspec, c_src_value, c_dest_value, c_strict_validation)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueDefaults is a wrapper around the C function g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_param_value_defaults(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValueSetDefault is a wrapper around the C function g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.g_param_value_set_default(c_pspec, c_value)

	return
}

// ParamValueValidate is a wrapper around the C function g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_param_value_validate(c_pspec, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// ParamValuesCmp is a wrapper around the C function g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int32 {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_value1 := (*C.GValue)(C.NULL)
	if value1 != nil {
		c_value1 = (*C.GValue)(value1.ToC())
	}

	c_value2 := (*C.GValue)(C.NULL)
	if value2 != nil {
		c_value2 = (*C.GValue)(value2.ToC())
	}

	retC := C.g_param_values_cmp(c_pspec, c_value1, c_value2)
	retGo := (int32)(retC)

	return retGo
}

// PointerTypeRegisterStatic is a wrapper around the C function g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_pointer_type_register_static(c_name)
	retGo := (Type)(retC)

	return retGo
}

// SignalAccumulatorTrueHandled is a wrapper around the C function g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy uintptr) bool {
	c_ihint := (*C.GSignalInvocationHint)(C.NULL)
	if ihint != nil {
		c_ihint = (*C.GSignalInvocationHint)(ihint.ToC())
	}

	c_return_accu := (*C.GValue)(C.NULL)
	if returnAccu != nil {
		c_return_accu = (*C.GValue)(returnAccu.ToC())
	}

	c_handler_return := (*C.GValue)(C.NULL)
	if handlerReturn != nil {
		c_handler_return = (*C.GValue)(handlerReturn.ToC())
	}

	c_dummy := (C.gpointer)(dummy)

	retC := C.g_signal_accumulator_true_handled(c_ihint, c_return_accu, c_handler_return, c_dummy)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_signal_add_emission_hook : unsupported parameter hook_func : no type generator for SignalEmissionHook (GSignalEmissionHook) for param hook_func

// Unsupported : g_signal_chain_from_overridden : unsupported parameter instance_and_params :

// Unsupported : g_signal_chain_from_overridden_handler : unsupported parameter ... : varargs

// SignalConnectClosure is a wrapper around the C function g_signal_connect_closure.
func SignalConnectClosure(instance *Object, detailedSignal string, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure(c_instance, c_detailed_signal, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// SignalConnectClosureById is a wrapper around the C function g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance *Object, signalId uint32, detail glib.Quark, closure *Closure, after bool) uint64 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_after :=
		boolToGboolean(after)

	retC := C.g_signal_connect_closure_by_id(c_instance, c_signal_id, c_detail, c_closure, c_after)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_signal_connect_data : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_connect_object : unsupported parameter c_handler : no type generator for Callback (GCallback) for param c_handler

// Unsupported : g_signal_emit : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_by_name : unsupported parameter ... : varargs

// Unsupported : g_signal_emit_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : g_signal_emitv : unsupported parameter instance_and_params :

// SignalGetInvocationHint is a wrapper around the C function g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance *Object) *SignalInvocationHint {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	retC := C.g_signal_get_invocation_hint(c_instance)
	retGo := SignalInvocationHintNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SignalHandlerBlock is a wrapper around the C function g_signal_handler_block.
func SignalHandlerBlock(instance *Object, handlerId uint64) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_block(c_instance, c_handler_id)

	return
}

// SignalHandlerDisconnect is a wrapper around the C function g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance *Object, handlerId uint64) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_disconnect(c_instance, c_handler_id)

	return
}

// SignalHandlerFind is a wrapper around the C function g_signal_handler_find.
func SignalHandlerFind(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint64 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handler_find(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint64)(retC)

	return retGo
}

// SignalHandlerIsConnected is a wrapper around the C function g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance *Object, handlerId uint64) bool {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_handler_id := (C.gulong)(handlerId)

	retC := C.g_signal_handler_is_connected(c_instance, c_handler_id)
	retGo := retC == C.TRUE

	return retGo
}

// SignalHandlerUnblock is a wrapper around the C function g_signal_handler_unblock.
func SignalHandlerUnblock(instance *Object, handlerId uint64) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_handler_id := (C.gulong)(handlerId)

	C.g_signal_handler_unblock(c_instance, c_handler_id)

	return
}

// SignalHandlersBlockMatched is a wrapper around the C function g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_block_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHandlersDestroy is a wrapper around the C function g_signal_handlers_destroy.
func SignalHandlersDestroy(instance *Object) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	C.g_signal_handlers_destroy(c_instance)

	return
}

// SignalHandlersDisconnectMatched is a wrapper around the C function g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_disconnect_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHandlersUnblockMatched is a wrapper around the C function g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ uintptr, data uintptr) uint32 {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_mask := (C.GSignalMatchType)(mask)

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_signal_handlers_unblock_matched(c_instance, c_mask, c_signal_id, c_detail, c_closure, c_func, c_data)
	retGo := (uint32)(retC)

	return retGo
}

// SignalHasHandlerPending is a wrapper around the C function g_signal_has_handler_pending.
func SignalHasHandlerPending(instance *Object, signalId uint32, detail glib.Quark, mayBeBlocked bool) bool {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	c_may_be_blocked :=
		boolToGboolean(mayBeBlocked)

	retC := C.g_signal_has_handler_pending(c_instance, c_signal_id, c_detail, c_may_be_blocked)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_signal_list_ids : array return type :

// SignalLookup is a wrapper around the C function g_signal_lookup.
func SignalLookup(name string, itype Type) uint32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_itype := (C.GType)(itype)

	retC := C.g_signal_lookup(c_name, c_itype)
	retGo := (uint32)(retC)

	return retGo
}

// SignalName is a wrapper around the C function g_signal_name.
func SignalName(signalId uint32) string {
	c_signal_id := (C.guint)(signalId)

	retC := C.g_signal_name(c_signal_id)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_signal_new : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_new_class_handler : unsupported parameter class_handler : no type generator for Callback (GCallback) for param class_handler

// Unsupported : g_signal_new_valist : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// Unsupported : g_signal_newv : unsupported parameter accumulator : no type generator for SignalAccumulator (GSignalAccumulator) for param accumulator

// SignalOverrideClassClosure is a wrapper around the C function g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId uint32, instanceType Type, classClosure *Closure) {
	c_signal_id := (C.guint)(signalId)

	c_instance_type := (C.GType)(instanceType)

	c_class_closure := (*C.GClosure)(C.NULL)
	if classClosure != nil {
		c_class_closure = (*C.GClosure)(classClosure.ToC())
	}

	C.g_signal_override_class_closure(c_signal_id, c_instance_type, c_class_closure)

	return
}

// Unsupported : g_signal_override_class_handler : unsupported parameter class_handler : no type generator for Callback (GCallback) for param class_handler

// SignalParseName is a wrapper around the C function g_signal_parse_name.
func SignalParseName(detailedSignal string, itype Type, forceDetailQuark bool) (bool, uint32, glib.Quark) {
	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	c_itype := (C.GType)(itype)

	var c_signal_id_p C.guint

	var c_detail_p C.GQuark

	c_force_detail_quark :=
		boolToGboolean(forceDetailQuark)

	retC := C.g_signal_parse_name(c_detailed_signal, c_itype, &c_signal_id_p, &c_detail_p, c_force_detail_quark)
	retGo := retC == C.TRUE

	signalIdP := (uint32)(c_signal_id_p)

	detailP := (glib.Quark)(c_detail_p)

	return retGo, signalIdP, detailP
}

// SignalQuery_ is a wrapper around the C function g_signal_query.
func SignalQuery_(signalId uint32) *SignalQuery {
	c_signal_id := (C.guint)(signalId)

	var c_query C.GSignalQuery

	C.g_signal_query(c_signal_id, &c_query)

	query := SignalQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// SignalRemoveEmissionHook is a wrapper around the C function g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	c_signal_id := (C.guint)(signalId)

	c_hook_id := (C.gulong)(hookId)

	C.g_signal_remove_emission_hook(c_signal_id, c_hook_id)

	return
}

// SignalStopEmission is a wrapper around the C function g_signal_stop_emission.
func SignalStopEmission(instance *Object, signalId uint32, detail glib.Quark) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_signal_id := (C.guint)(signalId)

	c_detail := (C.GQuark)(detail)

	C.g_signal_stop_emission(c_instance, c_signal_id, c_detail)

	return
}

// SignalStopEmissionByName is a wrapper around the C function g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance *Object, detailedSignal string) {
	c_instance := (C.gpointer)(C.NULL)
	if instance != nil {
		c_instance = (C.gpointer)(instance.ToC())
	}

	c_detailed_signal := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(c_detailed_signal))

	C.g_signal_stop_emission_by_name(c_instance, c_detailed_signal)

	return
}

// SignalTypeCclosureNew is a wrapper around the C function g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype Type, structOffset uint32) *Closure {
	c_itype := (C.GType)(itype)

	c_struct_offset := (C.guint)(structOffset)

	retC := C.g_signal_type_cclosure_new(c_itype, c_struct_offset)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SourceSetClosure is a wrapper around the C function g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	C.g_source_set_closure(c_source, c_closure)

	return
}

// SourceSetDummyCallback is a wrapper around the C function g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	c_source := (*C.GSource)(C.NULL)
	if source != nil {
		c_source = (*C.GSource)(source.ToC())
	}

	C.g_source_set_dummy_callback(c_source)

	return
}

// StrdupValueContents is a wrapper around the C function g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_strdup_value_contents(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_type_add_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// TypeAddInstancePrivate is a wrapper around the C function g_type_add_instance_private.
func TypeAddInstancePrivate(classType Type, privateSize uint64) int32 {
	c_class_type := (C.GType)(classType)

	c_private_size := (C.gsize)(privateSize)

	retC := C.g_type_add_instance_private(c_class_type, c_private_size)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_type_add_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func

// TypeAddInterfaceDynamic is a wrapper around the C function g_type_add_interface_dynamic.
func TypeAddInterfaceDynamic(instanceType Type, interfaceType Type, plugin *TypePlugin) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	C.g_type_add_interface_dynamic(c_instance_type, c_interface_type, c_plugin)

	return
}

// TypeAddInterfaceStatic is a wrapper around the C function g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_add_interface_static(c_instance_type, c_interface_type, c_info)

	return
}

// TypeCheckClassCast is a wrapper around the C function g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType Type) *TypeClass {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_cast(c_g_class, c_is_a_type)
	retGo := TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckClassIsA is a wrapper around the C function g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType Type) bool {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_check_class_is_a(c_g_class, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstance is a wrapper around the C function g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_check_instance(c_instance)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceCast is a wrapper around the C function g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType Type) *TypeInstance {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_cast(c_instance, c_iface_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeCheckInstanceIsA is a wrapper around the C function g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_check_instance_is_a(c_instance, c_iface_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckInstanceIsFundamentallyA is a wrapper around the C function g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType Type) bool {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	c_fundamental_type := (C.GType)(fundamentalType)

	retC := C.g_type_check_instance_is_fundamentally_a(c_instance, c_fundamental_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckIsValueType is a wrapper around the C function g_type_check_is_value_type.
func TypeCheckIsValueType(type_ Type) bool {
	c_type := (C.GType)(type_)

	retC := C.g_type_check_is_value_type(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckValue is a wrapper around the C function g_type_check_value.
func TypeCheckValue(value *Value) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_type_check_value(c_value)
	retGo := retC == C.TRUE

	return retGo
}

// TypeCheckValueHolds is a wrapper around the C function g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ Type) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	c_type := (C.GType)(type_)

	retC := C.g_type_check_value_holds(c_value, c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_type_children : array return type :

// TypeCreateInstance is a wrapper around the C function g_type_create_instance.
func TypeCreateInstance(type_ Type) *TypeInstance {
	c_type := (C.GType)(type_)

	retC := C.g_type_create_instance(c_type)
	retGo := TypeInstanceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDefaultInterfacePeek is a wrapper around the C function g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType Type) TypeInterface {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_peek(c_g_type)
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDefaultInterfaceRef is a wrapper around the C function g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType Type) TypeInterface {
	c_g_type := (C.GType)(gType)

	retC := C.g_type_default_interface_ref(c_g_type)
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeDefaultInterfaceUnref is a wrapper around the C function g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface *TypeInterface) {
	c_g_iface := (C.gpointer)(C.NULL)
	if gIface != nil {
		c_g_iface = (C.gpointer)(gIface.ToC())
	}

	C.g_type_default_interface_unref(c_g_iface)

	return
}

// TypeDepth is a wrapper around the C function g_type_depth.
func TypeDepth(type_ Type) uint32 {
	c_type := (C.GType)(type_)

	retC := C.g_type_depth(c_type)
	retGo := (uint32)(retC)

	return retGo
}

// TypeFreeInstance is a wrapper around the C function g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	C.g_type_free_instance(c_instance)

	return
}

// TypeFromName is a wrapper around the C function g_type_from_name.
func TypeFromName(name string) Type {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_type_from_name(c_name)
	retGo := (Type)(retC)

	return retGo
}

// TypeFundamental is a wrapper around the C function g_type_fundamental.
func TypeFundamental(typeId Type) Type {
	c_type_id := (C.GType)(typeId)

	retC := C.g_type_fundamental(c_type_id)
	retGo := (Type)(retC)

	return retGo
}

// TypeFundamentalNext is a wrapper around the C function g_type_fundamental_next.
func TypeFundamentalNext() Type {
	retC := C.g_type_fundamental_next()
	retGo := (Type)(retC)

	return retGo
}

// TypeGetPlugin is a wrapper around the C function g_type_get_plugin.
func TypeGetPlugin(type_ Type) *TypePlugin {
	c_type := (C.GType)(type_)

	retC := C.g_type_get_plugin(c_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeGetQdata is a wrapper around the C function g_type_get_qdata.
func TypeGetQdata(type_ Type, quark glib.Quark) uintptr {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	retC := C.g_type_get_qdata(c_type, c_quark)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TypeInit is a wrapper around the C function g_type_init.
func TypeInit() {
	C.g_type_init()

	return
}

// TypeInitWithDebugFlags is a wrapper around the C function g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	c_debug_flags := (C.GTypeDebugFlags)(debugFlags)

	C.g_type_init_with_debug_flags(c_debug_flags)

	return
}

// Unsupported : g_type_interface_prerequisites : array return type :

// Unsupported : g_type_interfaces : array return type :

// TypeIsA is a wrapper around the C function g_type_is_a.
func TypeIsA(type_ Type, isAType Type) bool {
	c_type := (C.GType)(type_)

	c_is_a_type := (C.GType)(isAType)

	retC := C.g_type_is_a(c_type, c_is_a_type)
	retGo := retC == C.TRUE

	return retGo
}

// TypeName is a wrapper around the C function g_type_name.
func TypeName(type_ Type) string {
	c_type := (C.GType)(type_)

	retC := C.g_type_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromClass is a wrapper around the C function g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	c_g_class := (*C.GTypeClass)(C.NULL)
	if gClass != nil {
		c_g_class = (*C.GTypeClass)(gClass.ToC())
	}

	retC := C.g_type_name_from_class(c_g_class)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNameFromInstance is a wrapper around the C function g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	c_instance := (*C.GTypeInstance)(C.NULL)
	if instance != nil {
		c_instance = (*C.GTypeInstance)(instance.ToC())
	}

	retC := C.g_type_name_from_instance(c_instance)
	retGo := C.GoString(retC)

	return retGo
}

// TypeNextBase is a wrapper around the C function g_type_next_base.
func TypeNextBase(leafType Type, rootType Type) Type {
	c_leaf_type := (C.GType)(leafType)

	c_root_type := (C.GType)(rootType)

	retC := C.g_type_next_base(c_leaf_type, c_root_type)
	retGo := (Type)(retC)

	return retGo
}

// TypeParent is a wrapper around the C function g_type_parent.
func TypeParent(type_ Type) Type {
	c_type := (C.GType)(type_)

	retC := C.g_type_parent(c_type)
	retGo := (Type)(retC)

	return retGo
}

// TypeQname is a wrapper around the C function g_type_qname.
func TypeQname(type_ Type) glib.Quark {
	c_type := (C.GType)(type_)

	retC := C.g_type_qname(c_type)
	retGo := (glib.Quark)(retC)

	return retGo
}

// TypeQuery_ is a wrapper around the C function g_type_query.
func TypeQuery_(type_ Type) *TypeQuery {
	c_type := (C.GType)(type_)

	var c_query C.GTypeQuery

	C.g_type_query(c_type, &c_query)

	query := TypeQueryNewFromC(unsafe.Pointer(&c_query))

	return query
}

// TypeRegisterDynamic is a wrapper around the C function g_type_register_dynamic.
func TypeRegisterDynamic(parentType Type, typeName string, plugin *TypePlugin, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_plugin := (*C.GTypePlugin)(plugin.ToC())

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_dynamic(c_parent_type, c_type_name, c_plugin, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// TypeRegisterFundamental is a wrapper around the C function g_type_register_fundamental.
func TypeRegisterFundamental(typeId Type, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) Type {
	c_type_id := (C.GType)(typeId)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_finfo := (*C.GTypeFundamentalInfo)(C.NULL)
	if finfo != nil {
		c_finfo = (*C.GTypeFundamentalInfo)(finfo.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_fundamental(c_type_id, c_type_name, c_info, c_finfo, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// TypeRegisterStatic is a wrapper around the C function g_type_register_static.
func TypeRegisterStatic(parentType Type, typeName string, info *TypeInfo, flags TypeFlags) Type {
	c_parent_type := (C.GType)(parentType)

	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_flags := (C.GTypeFlags)(flags)

	retC := C.g_type_register_static(c_parent_type, c_type_name, c_info, c_flags)
	retGo := (Type)(retC)

	return retGo
}

// Unsupported : g_type_register_static_simple : unsupported parameter class_init : no type generator for ClassInitFunc (GClassInitFunc) for param class_init

// Unsupported : g_type_remove_class_cache_func : unsupported parameter cache_func : no type generator for TypeClassCacheFunc (GTypeClassCacheFunc) for param cache_func

// Unsupported : g_type_remove_interface_check : unsupported parameter check_func : no type generator for TypeInterfaceCheckFunc (GTypeInterfaceCheckFunc) for param check_func

// TypeSetQdata is a wrapper around the C function g_type_set_qdata.
func TypeSetQdata(type_ Type, quark glib.Quark, data uintptr) {
	c_type := (C.GType)(type_)

	c_quark := (C.GQuark)(quark)

	c_data := (C.gpointer)(data)

	C.g_type_set_qdata(c_type, c_quark, c_data)

	return
}

// TypeTestFlags is a wrapper around the C function g_type_test_flags.
func TypeTestFlags(type_ Type, flags uint32) bool {
	c_type := (C.GType)(type_)

	c_flags := (C.guint)(flags)

	retC := C.g_type_test_flags(c_type, c_flags)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func

// TypePlugin is a wrapper around the C record GTypePlugin.
type TypePlugin struct {
	native *C.GTypePlugin
}

func TypePluginNewFromC(u unsafe.Pointer) *TypePlugin {
	c := (*C.GTypePlugin)(u)
	if c == nil {
		return nil
	}

	g := &TypePlugin{native: c}

	return g
}

func (recv *TypePlugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypePlugin with another TypePlugin, and returns true if they represent the same GObject.
func (recv *TypePlugin) Equals(other *TypePlugin) bool {
	return other.ToC() == recv.ToC()
}

// CompleteInterfaceInfo is a wrapper around the C function g_type_plugin_complete_interface_info.
func (recv *TypePlugin) CompleteInterfaceInfo(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_plugin_complete_interface_info((*C.GTypePlugin)(recv.native), c_instance_type, c_interface_type, c_info)

	return
}

// CompleteTypeInfo is a wrapper around the C function g_type_plugin_complete_type_info.
func (recv *TypePlugin) CompleteTypeInfo(gType Type, info *TypeInfo, valueTable *TypeValueTable) {
	c_g_type := (C.GType)(gType)

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_value_table := (*C.GTypeValueTable)(C.NULL)
	if valueTable != nil {
		c_value_table = (*C.GTypeValueTable)(valueTable.ToC())
	}

	C.g_type_plugin_complete_type_info((*C.GTypePlugin)(recv.native), c_g_type, c_info, c_value_table)

	return
}

// Unuse is a wrapper around the C function g_type_plugin_unuse.
func (recv *TypePlugin) Unuse() {
	C.g_type_plugin_unuse((*C.GTypePlugin)(recv.native))

	return
}

// Use is a wrapper around the C function g_type_plugin_use.
func (recv *TypePlugin) Use() {
	C.g_type_plugin_use((*C.GTypePlugin)(recv.native))

	return
}

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native *C.GCClosure
	// closure : record
	Callback uintptr
}

func CClosureNewFromC(u unsafe.Pointer) *CClosure {
	c := (*C.GCClosure)(u)
	if c == nil {
		return nil
	}

	g := &CClosure{
		Callback: (uintptr)(c.callback),
		native:   c,
	}

	return g
}

func (recv *CClosure) ToC() unsafe.Pointer {
	recv.native.callback =
		(C.gpointer)(recv.Callback)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CClosure with another CClosure, and returns true if they represent the same GObject.
func (recv *CClosure) Equals(other *CClosure) bool {
	return other.ToC() == recv.ToC()
}

// CClosureMarshalBooleanBoxedBoxed is a wrapper around the C function g_cclosure_marshal_BOOLEAN__BOXED_BOXED.
func CClosureMarshalBooleanBoxedBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_BOOLEAN__BOXED_BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalBooleanFlags is a wrapper around the C function g_cclosure_marshal_BOOLEAN__FLAGS.
func CClosureMarshalBooleanFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_BOOLEAN__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_BOOLEAN__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalStringObjectPointer is a wrapper around the C function g_cclosure_marshal_STRING__OBJECT_POINTER.
func CClosureMarshalStringObjectPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_STRING__OBJECT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_STRING__OBJECT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidBoolean is a wrapper around the C function g_cclosure_marshal_VOID__BOOLEAN.
func CClosureMarshalVoidBoolean(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOOLEAN(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__BOOLEANv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidBoxed is a wrapper around the C function g_cclosure_marshal_VOID__BOXED.
func CClosureMarshalVoidBoxed(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__BOXED(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__BOXEDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidChar is a wrapper around the C function g_cclosure_marshal_VOID__CHAR.
func CClosureMarshalVoidChar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__CHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__CHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidDouble is a wrapper around the C function g_cclosure_marshal_VOID__DOUBLE.
func CClosureMarshalVoidDouble(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__DOUBLE(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__DOUBLEv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidEnum is a wrapper around the C function g_cclosure_marshal_VOID__ENUM.
func CClosureMarshalVoidEnum(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ENUM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__ENUMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidFlags is a wrapper around the C function g_cclosure_marshal_VOID__FLAGS.
func CClosureMarshalVoidFlags(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLAGS(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__FLAGSv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidFloat is a wrapper around the C function g_cclosure_marshal_VOID__FLOAT.
func CClosureMarshalVoidFloat(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__FLOAT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__FLOATv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidInt is a wrapper around the C function g_cclosure_marshal_VOID__INT.
func CClosureMarshalVoidInt(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__INT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__INTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidLong is a wrapper around the C function g_cclosure_marshal_VOID__LONG.
func CClosureMarshalVoidLong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__LONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__LONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidObject is a wrapper around the C function g_cclosure_marshal_VOID__OBJECT.
func CClosureMarshalVoidObject(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__OBJECT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__OBJECTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidParam is a wrapper around the C function g_cclosure_marshal_VOID__PARAM.
func CClosureMarshalVoidParam(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__PARAM(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__PARAMv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidPointer is a wrapper around the C function g_cclosure_marshal_VOID__POINTER.
func CClosureMarshalVoidPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidString is a wrapper around the C function g_cclosure_marshal_VOID__STRING.
func CClosureMarshalVoidString(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__STRING(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__STRINGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUchar is a wrapper around the C function g_cclosure_marshal_VOID__UCHAR.
func CClosureMarshalVoidUchar(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UCHAR(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__UCHARv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUint is a wrapper around the C function g_cclosure_marshal_VOID__UINT.
func CClosureMarshalVoidUint(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// CClosureMarshalVoidUintPointer is a wrapper around the C function g_cclosure_marshal_VOID__UINT_POINTER.
func CClosureMarshalVoidUintPointer(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__UINT_POINTER(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__UINT_POINTERv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_marshal_VOID__UINTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidUlong is a wrapper around the C function g_cclosure_marshal_VOID__ULONG.
func CClosureMarshalVoidUlong(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__ULONG(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__ULONGv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidVariant is a wrapper around the C function g_cclosure_marshal_VOID__VARIANT.
func CClosureMarshalVoidVariant(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VARIANT(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__VARIANTv : unsupported parameter args : no type generator for va_list (va_list) for param args
// CClosureMarshalVoidVoid is a wrapper around the C function g_cclosure_marshal_VOID__VOID.
func CClosureMarshalVoidVoid(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint uintptr, marshalData uintptr) {
	c_closure := (*C.GClosure)(C.NULL)
	if closure != nil {
		c_closure = (*C.GClosure)(closure.ToC())
	}

	c_return_value := (*C.GValue)(C.NULL)
	if returnValue != nil {
		c_return_value = (*C.GValue)(returnValue.ToC())
	}

	c_n_param_values := (C.guint)(nParamValues)

	c_param_values := (*C.GValue)(C.NULL)
	if paramValues != nil {
		c_param_values = (*C.GValue)(paramValues.ToC())
	}

	c_invocation_hint := (C.gpointer)(invocationHint)

	c_marshal_data := (C.gpointer)(marshalData)

	C.g_cclosure_marshal_VOID__VOID(c_closure, c_return_value, c_n_param_values, c_param_values, c_invocation_hint, c_marshal_data)

	return
}

// g_cclosure_marshal_VOID__VOIDv : unsupported parameter args : no type generator for va_list (va_list) for param args
// g_cclosure_new : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_object_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// g_cclosure_new_swap : unsupported parameter callback_func : no type generator for Callback (GCallback) for param callback_func
// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native *C.GClosure
	// Private : ref_count
	// Private : meta_marshal_nouse
	// Private : n_guards
	// Private : n_fnotifiers
	// Private : n_inotifiers
	// Private : in_inotify
	// Private : floating
	// Private : derivative_flag
	// Bitfield not supported :  1 in_marshal
	// Bitfield not supported :  1 is_invalid
	// no type for marshal
	// Private : data
	// Private : notifiers
}

func ClosureNewFromC(u unsafe.Pointer) *Closure {
	c := (*C.GClosure)(u)
	if c == nil {
		return nil
	}

	g := &Closure{native: c}

	return g
}

func (recv *Closure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Closure with another Closure, and returns true if they represent the same GObject.
func (recv *Closure) Equals(other *Closure) bool {
	return other.ToC() == recv.ToC()
}

// ClosureNewObject is a wrapper around the C function g_closure_new_object.
func ClosureNewObject(sizeofClosure uint32, object *Object) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	retC := C.g_closure_new_object(c_sizeof_closure, c_object)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClosureNewSimple is a wrapper around the C function g_closure_new_simple.
func ClosureNewSimple(sizeofClosure uint32, data uintptr) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_data := (C.gpointer)(data)

	retC := C.g_closure_new_simple(c_sizeof_closure, c_data)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_notify : no type generator for ClosureNotify (GClosureNotify) for param pre_marshal_notify

// Invalidate is a wrapper around the C function g_closure_invalidate.
func (recv *Closure) Invalidate() {
	C.g_closure_invalidate((*C.GClosure)(recv.native))

	return
}

// Unsupported : g_closure_invoke : unsupported parameter param_values :

// Ref is a wrapper around the C function g_closure_ref.
func (recv *Closure) Ref() *Closure {
	retC := C.g_closure_ref((*C.GClosure)(recv.native))
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal (GClosureMarshal) for param marshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter meta_marshal : no type generator for ClosureMarshal (GClosureMarshal) for param meta_marshal

// Sink is a wrapper around the C function g_closure_sink.
func (recv *Closure) Sink() {
	C.g_closure_sink((*C.GClosure)(recv.native))

	return
}

// Unref is a wrapper around the C function g_closure_unref.
func (recv *Closure) Unref() {
	C.g_closure_unref((*C.GClosure)(recv.native))

	return
}

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native *C.GClosureNotifyData
	Data   uintptr
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func ClosureNotifyDataNewFromC(u unsafe.Pointer) *ClosureNotifyData {
	c := (*C.GClosureNotifyData)(u)
	if c == nil {
		return nil
	}

	g := &ClosureNotifyData{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ClosureNotifyData with another ClosureNotifyData, and returns true if they represent the same GObject.
func (recv *ClosureNotifyData) Equals(other *ClosureNotifyData) bool {
	return other.ToC() == recv.ToC()
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native *C.GEnumClass
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

func EnumClassNewFromC(u unsafe.Pointer) *EnumClass {
	c := (*C.GEnumClass)(u)
	if c == nil {
		return nil
	}

	g := &EnumClass{
		Maximum: (int32)(c.maximum),
		Minimum: (int32)(c.minimum),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *EnumClass) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumClass with another EnumClass, and returns true if they represent the same GObject.
func (recv *EnumClass) Equals(other *EnumClass) bool {
	return other.ToC() == recv.ToC()
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    *C.GEnumValue
	Value     int32
	ValueName string
	ValueNick string
}

func EnumValueNewFromC(u unsafe.Pointer) *EnumValue {
	c := (*C.GEnumValue)(u)
	if c == nil {
		return nil
	}

	g := &EnumValue{
		Value:     (int32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *EnumValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.gint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EnumValue with another EnumValue, and returns true if they represent the same GObject.
func (recv *EnumValue) Equals(other *EnumValue) bool {
	return other.ToC() == recv.ToC()
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native *C.GFlagsClass
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func FlagsClassNewFromC(u unsafe.Pointer) *FlagsClass {
	c := (*C.GFlagsClass)(u)
	if c == nil {
		return nil
	}

	g := &FlagsClass{
		Mask:    (uint32)(c.mask),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *FlagsClass) ToC() unsafe.Pointer {
	recv.native.mask =
		(C.guint)(recv.Mask)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsClass with another FlagsClass, and returns true if they represent the same GObject.
func (recv *FlagsClass) Equals(other *FlagsClass) bool {
	return other.ToC() == recv.ToC()
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    *C.GFlagsValue
	Value     uint32
	ValueName string
	ValueNick string
}

func FlagsValueNewFromC(u unsafe.Pointer) *FlagsValue {
	c := (*C.GFlagsValue)(u)
	if c == nil {
		return nil
	}

	g := &FlagsValue{
		Value:     (uint32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *FlagsValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.guint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FlagsValue with another FlagsValue, and returns true if they represent the same GObject.
func (recv *FlagsValue) Equals(other *FlagsValue) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native *C.GInitiallyUnownedClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func InitiallyUnownedClassNewFromC(u unsafe.Pointer) *InitiallyUnownedClass {
	c := (*C.GInitiallyUnownedClass)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: c}

	return g
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InitiallyUnownedClass with another InitiallyUnownedClass, and returns true if they represent the same GObject.
func (recv *InitiallyUnownedClass) Equals(other *InitiallyUnownedClass) bool {
	return other.ToC() == recv.ToC()
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native *C.GInterfaceInfo
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	InterfaceData uintptr
}

func InterfaceInfoNewFromC(u unsafe.Pointer) *InterfaceInfo {
	c := (*C.GInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &InterfaceInfo{
		InterfaceData: (uintptr)(c.interface_data),
		native:        c,
	}

	return g
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	recv.native.interface_data =
		(C.gpointer)(recv.InterfaceData)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InterfaceInfo with another InterfaceInfo, and returns true if they represent the same GObject.
func (recv *InterfaceInfo) Equals(other *InterfaceInfo) bool {
	return other.ToC() == recv.ToC()
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native *C.GObjectClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.GObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same GObject.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// FindProperty is a wrapper around the C function g_object_class_find_property.
func (recv *ObjectClass) FindProperty(propertyName string) *ParamSpec {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_object_class_find_property((*C.GObjectClass)(recv.native), c_property_name)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InstallProperty is a wrapper around the C function g_object_class_install_property.
func (recv *ObjectClass) InstallProperty(propertyId uint32, pspec *ParamSpec) {
	c_property_id := (C.guint)(propertyId)

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_object_class_install_property((*C.GObjectClass)(recv.native), c_property_id, c_pspec)

	return
}

// Unsupported : g_object_class_list_properties : array return type :

// OverrideProperty is a wrapper around the C function g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	c_property_id := (C.guint)(propertyId)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_object_class_override_property((*C.GObjectClass)(recv.native), c_property_id, c_name)

	return
}

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native *C.GObjectConstructParam
	// pspec : record
	// value : record
}

func ObjectConstructParamNewFromC(u unsafe.Pointer) *ObjectConstructParam {
	c := (*C.GObjectConstructParam)(u)
	if c == nil {
		return nil
	}

	g := &ObjectConstructParam{native: c}

	return g
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectConstructParam with another ObjectConstructParam, and returns true if they represent the same GObject.
func (recv *ObjectConstructParam) Equals(other *ObjectConstructParam) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native *C.GParamSpecClass
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func ParamSpecClassNewFromC(u unsafe.Pointer) *ParamSpecClass {
	c := (*C.GParamSpecClass)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecClass{
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecClass with another ParamSpecClass, and returns true if they represent the same GObject.
func (recv *ParamSpecClass) Equals(other *ParamSpecClass) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native *C.GParamSpecPool
}

func ParamSpecPoolNewFromC(u unsafe.Pointer) *ParamSpecPool {
	c := (*C.GParamSpecPool)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPool{native: c}

	return g
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecPool with another ParamSpecPool, and returns true if they represent the same GObject.
func (recv *ParamSpecPool) Equals(other *ParamSpecPool) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpecPoolNew is a wrapper around the C function g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	c_type_prefixing :=
		boolToGboolean(typePrefixing)

	retC := C.g_param_spec_pool_new(c_type_prefixing)
	retGo := ParamSpecPoolNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_param_spec_pool_insert.
func (recv *ParamSpecPool) Insert(pspec *ParamSpec, ownerType Type) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_owner_type := (C.GType)(ownerType)

	C.g_param_spec_pool_insert((*C.GParamSpecPool)(recv.native), c_pspec, c_owner_type)

	return
}

// Unsupported : g_param_spec_pool_list : array return type :

// ListOwned is a wrapper around the C function g_param_spec_pool_list_owned.
func (recv *ParamSpecPool) ListOwned(ownerType Type) *glib.List {
	c_owner_type := (C.GType)(ownerType)

	retC := C.g_param_spec_pool_list_owned((*C.GParamSpecPool)(recv.native), c_owner_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lookup is a wrapper around the C function g_param_spec_pool_lookup.
func (recv *ParamSpecPool) Lookup(paramName string, ownerType Type, walkAncestors bool) *ParamSpec {
	c_param_name := C.CString(paramName)
	defer C.free(unsafe.Pointer(c_param_name))

	c_owner_type := (C.GType)(ownerType)

	c_walk_ancestors :=
		boolToGboolean(walkAncestors)

	retC := C.g_param_spec_pool_lookup((*C.GParamSpecPool)(recv.native), c_param_name, c_owner_type, c_walk_ancestors)
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_param_spec_pool_remove.
func (recv *ParamSpecPool) Remove(pspec *ParamSpec) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.g_param_spec_pool_remove((*C.GParamSpecPool)(recv.native), c_pspec)

	return
}

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       *C.GParamSpecTypeInfo
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func ParamSpecTypeInfoNewFromC(u unsafe.Pointer) *ParamSpecTypeInfo {
	c := (*C.GParamSpecTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		ValueType:    (Type)(c.value_type),
		native:       c,
	}

	return g
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ParamSpecTypeInfo with another ParamSpecTypeInfo, and returns true if they represent the same GObject.
func (recv *ParamSpecTypeInfo) Equals(other *ParamSpecTypeInfo) bool {
	return other.ToC() == recv.ToC()
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native *C.GParameter
	Name   string
	// value : record
}

func ParameterNewFromC(u unsafe.Pointer) *Parameter {
	c := (*C.GParameter)(u)
	if c == nil {
		return nil
	}

	g := &Parameter{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *Parameter) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Parameter with another Parameter, and returns true if they represent the same GObject.
func (recv *Parameter) Equals(other *Parameter) bool {
	return other.ToC() == recv.ToC()
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   *C.GSignalInvocationHint
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

func SignalInvocationHintNewFromC(u unsafe.Pointer) *SignalInvocationHint {
	c := (*C.GSignalInvocationHint)(u)
	if c == nil {
		return nil
	}

	g := &SignalInvocationHint{
		Detail:   (glib.Quark)(c.detail),
		RunType:  (SignalFlags)(c.run_type),
		SignalId: (uint32)(c.signal_id),
		native:   c,
	}

	return g
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.detail =
		(C.GQuark)(recv.Detail)
	recv.native.run_type =
		(C.GSignalFlags)(recv.RunType)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalInvocationHint with another SignalInvocationHint, and returns true if they represent the same GObject.
func (recv *SignalInvocationHint) Equals(other *SignalInvocationHint) bool {
	return other.ToC() == recv.ToC()
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      *C.GSignalQuery
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

func SignalQueryNewFromC(u unsafe.Pointer) *SignalQuery {
	c := (*C.GSignalQuery)(u)
	if c == nil {
		return nil
	}

	g := &SignalQuery{
		Itype:       (Type)(c.itype),
		NParams:     (uint32)(c.n_params),
		ReturnType:  (Type)(c.return_type),
		SignalFlags: (SignalFlags)(c.signal_flags),
		SignalId:    (uint32)(c.signal_id),
		SignalName:  C.GoString(c.signal_name),
		native:      c,
	}

	return g
}

func (recv *SignalQuery) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.itype =
		(C.GType)(recv.Itype)
	recv.native.signal_flags =
		(C.GSignalFlags)(recv.SignalFlags)
	recv.native.return_type =
		(C.GType)(recv.ReturnType)
	recv.native.n_params =
		(C.guint)(recv.NParams)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SignalQuery with another SignalQuery, and returns true if they represent the same GObject.
func (recv *SignalQuery) Equals(other *SignalQuery) bool {
	return other.ToC() == recv.ToC()
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native *C.GTypeClass
	// Private : g_type
}

func TypeClassNewFromC(u unsafe.Pointer) *TypeClass {
	c := (*C.GTypeClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeClass{native: c}

	return g
}

func (recv *TypeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeClass with another TypeClass, and returns true if they represent the same GObject.
func (recv *TypeClass) Equals(other *TypeClass) bool {
	return other.ToC() == recv.ToC()
}

// TypeClassAdjustPrivateOffset is a wrapper around the C function g_type_class_adjust_private_offset.
func TypeClassAdjustPrivateOffset(gClass uintptr, privateSizeOrOffset int32) {
	c_g_class := (C.gpointer)(gClass)

	c_private_size_or_offset := (C.gint)(privateSizeOrOffset)

	C.g_type_class_adjust_private_offset(c_g_class, &c_private_size_or_offset)

	return
}

// TypeClassPeek is a wrapper around the C function g_type_class_peek.
func TypeClassPeek(type_ Type) TypeClass {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek(c_type)
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeClassPeekStatic is a wrapper around the C function g_type_class_peek_static.
func TypeClassPeekStatic(type_ Type) TypeClass {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek_static(c_type)
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeClassRef is a wrapper around the C function g_type_class_ref.
func TypeClassRef(type_ Type) TypeClass {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_ref(c_type)
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddPrivate is a wrapper around the C function g_type_class_add_private.
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	c_private_size := (C.gsize)(privateSize)

	C.g_type_class_add_private((C.gpointer)(recv.native), c_private_size)

	return
}

// GetPrivate is a wrapper around the C function g_type_class_get_private.
func (recv *TypeClass) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_class_get_private((*C.GTypeClass)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PeekParent is a wrapper around the C function g_type_class_peek_parent.
func (recv *TypeClass) PeekParent() TypeClass {
	retC := C.g_type_class_peek_parent((C.gpointer)(recv.native))
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_type_class_unref.
func (recv *TypeClass) Unref() {
	C.g_type_class_unref((C.gpointer)(recv.native))

	return
}

// UnrefUncached is a wrapper around the C function g_type_class_unref_uncached.
func (recv *TypeClass) UnrefUncached() {
	C.g_type_class_unref_uncached((C.gpointer)(recv.native))

	return
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    *C.GTypeFundamentalInfo
	TypeFlags TypeFundamentalFlags
}

func TypeFundamentalInfoNewFromC(u unsafe.Pointer) *TypeFundamentalInfo {
	c := (*C.GTypeFundamentalInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeFundamentalInfo{
		TypeFlags: (TypeFundamentalFlags)(c.type_flags),
		native:    c,
	}

	return g
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	recv.native.type_flags =
		(C.GTypeFundamentalFlags)(recv.TypeFlags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeFundamentalInfo with another TypeFundamentalInfo, and returns true if they represent the same GObject.
func (recv *TypeFundamentalInfo) Equals(other *TypeFundamentalInfo) bool {
	return other.ToC() == recv.ToC()
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    *C.GTypeInfo
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	ClassData    uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func TypeInfoNewFromC(u unsafe.Pointer) *TypeInfo {
	c := (*C.GTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeInfo{
		ClassData:    (uintptr)(c.class_data),
		ClassSize:    (uint16)(c.class_size),
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *TypeInfo) ToC() unsafe.Pointer {
	recv.native.class_size =
		(C.guint16)(recv.ClassSize)
	recv.native.class_data =
		(C.gconstpointer)(recv.ClassData)
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInfo with another TypeInfo, and returns true if they represent the same GObject.
func (recv *TypeInfo) Equals(other *TypeInfo) bool {
	return other.ToC() == recv.ToC()
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native *C.GTypeInstance
	// Private : g_class
}

func TypeInstanceNewFromC(u unsafe.Pointer) *TypeInstance {
	c := (*C.GTypeInstance)(u)
	if c == nil {
		return nil
	}

	g := &TypeInstance{native: c}

	return g
}

func (recv *TypeInstance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInstance with another TypeInstance, and returns true if they represent the same GObject.
func (recv *TypeInstance) Equals(other *TypeInstance) bool {
	return other.ToC() == recv.ToC()
}

// GetPrivate is a wrapper around the C function g_type_instance_get_private.
func (recv *TypeInstance) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_instance_get_private((*C.GTypeInstance)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native *C.GTypeInterface
	// Private : g_type
	// Private : g_instance_type
}

func TypeInterfaceNewFromC(u unsafe.Pointer) *TypeInterface {
	c := (*C.GTypeInterface)(u)
	if c == nil {
		return nil
	}

	g := &TypeInterface{native: c}

	return g
}

func (recv *TypeInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeInterface with another TypeInterface, and returns true if they represent the same GObject.
func (recv *TypeInterface) Equals(other *TypeInterface) bool {
	return other.ToC() == recv.ToC()
}

// TypeInterfaceAddPrerequisite is a wrapper around the C function g_type_interface_add_prerequisite.
func TypeInterfaceAddPrerequisite(interfaceType Type, prerequisiteType Type) {
	c_interface_type := (C.GType)(interfaceType)

	c_prerequisite_type := (C.GType)(prerequisiteType)

	C.g_type_interface_add_prerequisite(c_interface_type, c_prerequisite_type)

	return
}

// TypeInterfaceGetPlugin is a wrapper around the C function g_type_interface_get_plugin.
func TypeInterfaceGetPlugin(instanceType Type, interfaceType Type) *TypePlugin {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	retC := C.g_type_interface_get_plugin(c_instance_type, c_interface_type)
	retGo := TypePluginNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeInterfacePeek is a wrapper around the C function g_type_interface_peek.
func TypeInterfacePeek(instanceClass *TypeClass, ifaceType Type) TypeInterface {
	c_instance_class := (C.gpointer)(C.NULL)
	if instanceClass != nil {
		c_instance_class = (C.gpointer)(instanceClass.ToC())
	}

	c_iface_type := (C.GType)(ifaceType)

	retC := C.g_type_interface_peek(c_instance_class, c_iface_type)
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_type_interface_prerequisites : array return type :
// PeekParent is a wrapper around the C function g_type_interface_peek_parent.
func (recv *TypeInterface) PeekParent() TypeInterface {
	retC := C.g_type_interface_peek_parent((C.gpointer)(recv.native))
	retGo := *TypeInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native *C.GTypeModuleClass
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

func TypeModuleClassNewFromC(u unsafe.Pointer) *TypeModuleClass {
	c := (*C.GTypeModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeModuleClass{native: c}

	return g
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeModuleClass with another TypeModuleClass, and returns true if they represent the same GObject.
func (recv *TypeModuleClass) Equals(other *TypeModuleClass) bool {
	return other.ToC() == recv.ToC()
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native *C.GTypePluginClass
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

func TypePluginClassNewFromC(u unsafe.Pointer) *TypePluginClass {
	c := (*C.GTypePluginClass)(u)
	if c == nil {
		return nil
	}

	g := &TypePluginClass{native: c}

	return g
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypePluginClass with another TypePluginClass, and returns true if they represent the same GObject.
func (recv *TypePluginClass) Equals(other *TypePluginClass) bool {
	return other.ToC() == recv.ToC()
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       *C.GTypeQuery
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func TypeQueryNewFromC(u unsafe.Pointer) *TypeQuery {
	c := (*C.GTypeQuery)(u)
	if c == nil {
		return nil
	}

	g := &TypeQuery{
		ClassSize:    (uint32)(c.class_size),
		InstanceSize: (uint32)(c.instance_size),
		Type:         (Type)(c._type),
		TypeName:     C.GoString(c.type_name),
		native:       c,
	}

	return g
}

func (recv *TypeQuery) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GType)(recv.Type)
	recv.native.type_name =
		C.CString(recv.TypeName)
	recv.native.class_size =
		(C.guint)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint)(recv.InstanceSize)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeQuery with another TypeQuery, and returns true if they represent the same GObject.
func (recv *TypeQuery) Equals(other *TypeQuery) bool {
	return other.ToC() == recv.ToC()
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native *C.GTypeValueTable
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

func TypeValueTableNewFromC(u unsafe.Pointer) *TypeValueTable {
	c := (*C.GTypeValueTable)(u)
	if c == nil {
		return nil
	}

	g := &TypeValueTable{
		CollectFormat: C.GoString(c.collect_format),
		LcopyFormat:   C.GoString(c.lcopy_format),
		native:        c,
	}

	return g
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {
	recv.native.collect_format =
		C.CString(recv.CollectFormat)
	recv.native.lcopy_format =
		C.CString(recv.LcopyFormat)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TypeValueTable with another TypeValueTable, and returns true if they represent the same GObject.
func (recv *TypeValueTable) Equals(other *TypeValueTable) bool {
	return other.ToC() == recv.ToC()
}

// TypeValueTablePeek is a wrapper around the C function g_type_value_table_peek.
func TypeValueTablePeek(type_ Type) *TypeValueTable {
	c_type := (C.GType)(type_)

	retC := C.g_type_value_table_peek(c_type)
	retGo := TypeValueTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.GValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
// ValueTypeCompatible is a wrapper around the C function g_value_type_compatible.
func ValueTypeCompatible(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_compatible(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// ValueTypeTransformable is a wrapper around the C function g_value_type_transformable.
func ValueTypeTransformable(srcType Type, destType Type) bool {
	c_src_type := (C.GType)(srcType)

	c_dest_type := (C.GType)(destType)

	retC := C.g_value_type_transformable(c_src_type, c_dest_type)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function g_value_copy.
func (recv *Value) Copy(destValue *Value) {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	C.g_value_copy((*C.GValue)(recv.native), c_dest_value)

	return
}

// DupBoxed is a wrapper around the C function g_value_dup_boxed.
func (recv *Value) DupBoxed() uintptr {
	retC := C.g_value_dup_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DupObject is a wrapper around the C function g_value_dup_object.
func (recv *Value) DupObject() Object {
	retC := C.g_value_dup_object((*C.GValue)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupParam is a wrapper around the C function g_value_dup_param.
func (recv *Value) DupParam() *ParamSpec {
	retC := C.g_value_dup_param((*C.GValue)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupString is a wrapper around the C function g_value_dup_string.
func (recv *Value) DupString() string {
	retC := C.g_value_dup_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FitsPointer is a wrapper around the C function g_value_fits_pointer.
func (recv *Value) FitsPointer() bool {
	retC := C.g_value_fits_pointer((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBoolean is a wrapper around the C function g_value_get_boolean.
func (recv *Value) GetBoolean() bool {
	retC := C.g_value_get_boolean((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBoxed is a wrapper around the C function g_value_get_boxed.
func (recv *Value) GetBoxed() uintptr {
	retC := C.g_value_get_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetChar is a wrapper around the C function g_value_get_char.
func (recv *Value) GetChar() rune {
	retC := C.g_value_get_char((*C.GValue)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetDouble is a wrapper around the C function g_value_get_double.
func (recv *Value) GetDouble() float64 {
	retC := C.g_value_get_double((*C.GValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetEnum is a wrapper around the C function g_value_get_enum.
func (recv *Value) GetEnum() int32 {
	retC := C.g_value_get_enum((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_value_get_flags.
func (recv *Value) GetFlags() uint32 {
	retC := C.g_value_get_flags((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetFloat is a wrapper around the C function g_value_get_float.
func (recv *Value) GetFloat() float32 {
	retC := C.g_value_get_float((*C.GValue)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetGtype is a wrapper around the C function g_value_get_gtype.
func (recv *Value) GetGtype() Type {
	retC := C.g_value_get_gtype((*C.GValue)(recv.native))
	retGo := (Type)(retC)

	return retGo
}

// GetInt is a wrapper around the C function g_value_get_int.
func (recv *Value) GetInt() int32 {
	retC := C.g_value_get_int((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt64 is a wrapper around the C function g_value_get_int64.
func (recv *Value) GetInt64() int64 {
	retC := C.g_value_get_int64((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetLong is a wrapper around the C function g_value_get_long.
func (recv *Value) GetLong() int64 {
	retC := C.g_value_get_long((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetObject is a wrapper around the C function g_value_get_object.
func (recv *Value) GetObject() Object {
	retC := C.g_value_get_object((*C.GValue)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParam is a wrapper around the C function g_value_get_param.
func (recv *Value) GetParam() *ParamSpec {
	retC := C.g_value_get_param((*C.GValue)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPointer is a wrapper around the C function g_value_get_pointer.
func (recv *Value) GetPointer() uintptr {
	retC := C.g_value_get_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_value_get_string.
func (recv *Value) GetString() string {
	retC := C.g_value_get_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUchar is a wrapper around the C function g_value_get_uchar.
func (recv *Value) GetUchar() uint8 {
	retC := C.g_value_get_uchar((*C.GValue)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// GetUint is a wrapper around the C function g_value_get_uint.
func (recv *Value) GetUint() uint32 {
	retC := C.g_value_get_uint((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_value_get_uint64.
func (recv *Value) GetUint64() uint64 {
	retC := C.g_value_get_uint64((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetUlong is a wrapper around the C function g_value_get_ulong.
func (recv *Value) GetUlong() uint64 {
	retC := C.g_value_get_ulong((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Init is a wrapper around the C function g_value_init.
func (recv *Value) Init(gType Type) *Value {
	c_g_type := (C.GType)(gType)

	retC := C.g_value_init((*C.GValue)(recv.native), c_g_type)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekPointer is a wrapper around the C function g_value_peek_pointer.
func (recv *Value) PeekPointer() uintptr {
	retC := C.g_value_peek_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Reset is a wrapper around the C function g_value_reset.
func (recv *Value) Reset() *Value {
	retC := C.g_value_reset((*C.GValue)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetBoolean is a wrapper around the C function g_value_set_boolean.
func (recv *Value) SetBoolean(vBoolean bool) {
	c_v_boolean :=
		boolToGboolean(vBoolean)

	C.g_value_set_boolean((*C.GValue)(recv.native), c_v_boolean)

	return
}

// SetBoxed is a wrapper around the C function g_value_set_boxed.
func (recv *Value) SetBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetBoxedTakeOwnership is a wrapper around the C function g_value_set_boxed_take_ownership.
func (recv *Value) SetBoxedTakeOwnership(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed_take_ownership((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetChar is a wrapper around the C function g_value_set_char.
func (recv *Value) SetChar(vChar rune) {
	c_v_char := (C.gchar)(vChar)

	C.g_value_set_char((*C.GValue)(recv.native), c_v_char)

	return
}

// SetDouble is a wrapper around the C function g_value_set_double.
func (recv *Value) SetDouble(vDouble float64) {
	c_v_double := (C.gdouble)(vDouble)

	C.g_value_set_double((*C.GValue)(recv.native), c_v_double)

	return
}

// SetEnum is a wrapper around the C function g_value_set_enum.
func (recv *Value) SetEnum(vEnum int32) {
	c_v_enum := (C.gint)(vEnum)

	C.g_value_set_enum((*C.GValue)(recv.native), c_v_enum)

	return
}

// SetFlags is a wrapper around the C function g_value_set_flags.
func (recv *Value) SetFlags(vFlags uint32) {
	c_v_flags := (C.guint)(vFlags)

	C.g_value_set_flags((*C.GValue)(recv.native), c_v_flags)

	return
}

// SetFloat is a wrapper around the C function g_value_set_float.
func (recv *Value) SetFloat(vFloat float32) {
	c_v_float := (C.gfloat)(vFloat)

	C.g_value_set_float((*C.GValue)(recv.native), c_v_float)

	return
}

// SetGtype is a wrapper around the C function g_value_set_gtype.
func (recv *Value) SetGtype(vGtype Type) {
	c_v_gtype := (C.GType)(vGtype)

	C.g_value_set_gtype((*C.GValue)(recv.native), c_v_gtype)

	return
}

// SetInstance is a wrapper around the C function g_value_set_instance.
func (recv *Value) SetInstance(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_value_set_instance((*C.GValue)(recv.native), c_instance)

	return
}

// SetInt is a wrapper around the C function g_value_set_int.
func (recv *Value) SetInt(vInt int32) {
	c_v_int := (C.gint)(vInt)

	C.g_value_set_int((*C.GValue)(recv.native), c_v_int)

	return
}

// SetInt64 is a wrapper around the C function g_value_set_int64.
func (recv *Value) SetInt64(vInt64 int64) {
	c_v_int64 := (C.gint64)(vInt64)

	C.g_value_set_int64((*C.GValue)(recv.native), c_v_int64)

	return
}

// SetLong is a wrapper around the C function g_value_set_long.
func (recv *Value) SetLong(vLong int64) {
	c_v_long := (C.glong)(vLong)

	C.g_value_set_long((*C.GValue)(recv.native), c_v_long)

	return
}

// SetObject is a wrapper around the C function g_value_set_object.
func (recv *Value) SetObject(vObject *Object) {
	c_v_object := (C.gpointer)(C.NULL)
	if vObject != nil {
		c_v_object = (C.gpointer)(vObject.ToC())
	}

	C.g_value_set_object((*C.GValue)(recv.native), c_v_object)

	return
}

// SetObjectTakeOwnership is a wrapper around the C function g_value_set_object_take_ownership.
func (recv *Value) SetObjectTakeOwnership(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_set_object_take_ownership((*C.GValue)(recv.native), c_v_object)

	return
}

// SetParam is a wrapper around the C function g_value_set_param.
func (recv *Value) SetParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_set_param((*C.GValue)(recv.native), c_param)

	return
}

// SetParamTakeOwnership is a wrapper around the C function g_value_set_param_take_ownership.
func (recv *Value) SetParamTakeOwnership(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_set_param_take_ownership((*C.GValue)(recv.native), c_param)

	return
}

// SetPointer is a wrapper around the C function g_value_set_pointer.
func (recv *Value) SetPointer(vPointer uintptr) {
	c_v_pointer := (C.gpointer)(vPointer)

	C.g_value_set_pointer((*C.GValue)(recv.native), c_v_pointer)

	return
}

// SetStaticBoxed is a wrapper around the C function g_value_set_static_boxed.
func (recv *Value) SetStaticBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_static_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// SetStaticString is a wrapper around the C function g_value_set_static_string.
func (recv *Value) SetStaticString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_static_string((*C.GValue)(recv.native), c_v_string)

	return
}

// SetString is a wrapper around the C function g_value_set_string.
func (recv *Value) SetString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string((*C.GValue)(recv.native), c_v_string)

	return
}

// SetStringTakeOwnership is a wrapper around the C function g_value_set_string_take_ownership.
func (recv *Value) SetStringTakeOwnership(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string_take_ownership((*C.GValue)(recv.native), c_v_string)

	return
}

// SetUchar is a wrapper around the C function g_value_set_uchar.
func (recv *Value) SetUchar(vUchar uint8) {
	c_v_uchar := (C.guchar)(vUchar)

	C.g_value_set_uchar((*C.GValue)(recv.native), c_v_uchar)

	return
}

// SetUint is a wrapper around the C function g_value_set_uint.
func (recv *Value) SetUint(vUint uint32) {
	c_v_uint := (C.guint)(vUint)

	C.g_value_set_uint((*C.GValue)(recv.native), c_v_uint)

	return
}

// SetUint64 is a wrapper around the C function g_value_set_uint64.
func (recv *Value) SetUint64(vUint64 uint64) {
	c_v_uint64 := (C.guint64)(vUint64)

	C.g_value_set_uint64((*C.GValue)(recv.native), c_v_uint64)

	return
}

// SetUlong is a wrapper around the C function g_value_set_ulong.
func (recv *Value) SetUlong(vUlong uint64) {
	c_v_ulong := (C.gulong)(vUlong)

	C.g_value_set_ulong((*C.GValue)(recv.native), c_v_ulong)

	return
}

// TakeBoxed is a wrapper around the C function g_value_take_boxed.
func (recv *Value) TakeBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_take_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// TakeObject is a wrapper around the C function g_value_take_object.
func (recv *Value) TakeObject(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_take_object((*C.GValue)(recv.native), c_v_object)

	return
}

// TakeParam is a wrapper around the C function g_value_take_param.
func (recv *Value) TakeParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_take_param((*C.GValue)(recv.native), c_param)

	return
}

// TakeString is a wrapper around the C function g_value_take_string.
func (recv *Value) TakeString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_take_string((*C.GValue)(recv.native), c_v_string)

	return
}

// Transform is a wrapper around the C function g_value_transform.
func (recv *Value) Transform(destValue *Value) bool {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	retC := C.g_value_transform((*C.GValue)(recv.native), c_dest_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unset is a wrapper around the C function g_value_unset.
func (recv *Value) Unset() {
	C.g_value_unset((*C.GValue)(recv.native))

	return
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  *C.GValueArray
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func ValueArrayNewFromC(u unsafe.Pointer) *ValueArray {
	c := (*C.GValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ValueArray{
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *ValueArray) ToC() unsafe.Pointer {
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValueArray with another ValueArray, and returns true if they represent the same GObject.
func (recv *ValueArray) Equals(other *ValueArray) bool {
	return other.ToC() == recv.ToC()
}

// ValueArrayNew is a wrapper around the C function g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	c_n_prealloced := (C.guint)(nPrealloced)

	retC := C.g_value_array_new(c_n_prealloced)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function g_value_array_append.
func (recv *ValueArray) Append(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_append((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_value_array_copy.
func (recv *ValueArray) Copy() *ValueArray {
	retC := C.g_value_array_copy((*C.GValueArray)(recv.native))
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_value_array_free.
func (recv *ValueArray) Free() {
	C.g_value_array_free((*C.GValueArray)(recv.native))

	return
}

// GetNth is a wrapper around the C function g_value_array_get_nth.
func (recv *ValueArray) GetNth(index uint32) *Value {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_get_nth((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_value_array_insert.
func (recv *ValueArray) Insert(index uint32, value *Value) *ValueArray {
	c_index_ := (C.guint)(index)

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_insert((*C.GValueArray)(recv.native), c_index_, c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_value_array_prepend.
func (recv *ValueArray) Prepend(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_prepend((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_value_array_remove.
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_remove((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc (GCompareFunc) for param compare_func

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native *C.GWeakRef
}

func WeakRefNewFromC(u unsafe.Pointer) *WeakRef {
	c := (*C.GWeakRef)(u)
	if c == nil {
		return nil
	}

	g := &WeakRef{native: c}

	return g
}

func (recv *WeakRef) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WeakRef with another WeakRef, and returns true if they represent the same GObject.
func (recv *WeakRef) Equals(other *WeakRef) bool {
	return other.ToC() == recv.ToC()
}
