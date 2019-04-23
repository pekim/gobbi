// This is a generated file - DO NOT EDIT

package atk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void object_focusEventHandler(GObject *, gboolean, gpointer);

	static gulong Object_signal_connect_focus_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-event", G_CALLBACK(object_focusEventHandler), data);
	}

*/
/*

	void object_stateChangeHandler(GObject *, gchar*, gboolean, gpointer);

	static gulong Object_signal_connect_state_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-change", G_CALLBACK(object_stateChangeHandler), data);
	}

*/
/*

	void object_visibleDataChangedHandler(GObject *, gpointer);

	static gulong Object_signal_connect_visible_data_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "visible-data-changed", G_CALLBACK(object_visibleDataChangedHandler), data);
	}

*/
import "C"

// Blacklisted : AtkGObjectAccessible

// Blacklisted : AtkHyperlink

// Blacklisted : AtkMisc

// Blacklisted : AtkNoOpObject

// Blacklisted : AtkNoOpObjectFactory

// Object is a wrapper around the C record AtkObject.
type Object struct {
	native *C.AtkObject
	// parent : record
	Description string
	Name        string
	// accessible_parent : record
	Role Role
	// relation_set : record
	Layer Layer
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	c := (*C.AtkObject)(u)
	if c == nil {
		return nil
	}

	g := &Object{
		Description: C.GoString(c.description),
		Layer:       (Layer)(c.layer),
		Name:        C.GoString(c.name),
		Role:        (Role)(c.role),
		native:      c,
	}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {
	recv.native.description =
		C.CString(recv.Description)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.role =
		(C.AtkRole)(recv.Role)
	recv.native.layer =
		(C.AtkLayer)(recv.Layer)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Object with another Object, and returns true if they represent the same GObject.
func (recv *Object) Equals(other *Object) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Object) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *gobject.Object) *Object {
	return ObjectNewFromC(object.ToC())
}

// Unsupported signal 'active-descendant-changed' for Object : unsupported parameter arg1 : no type generator for gpointer, gpointer

// Unsupported signal 'children-changed' for Object : unsupported parameter arg2 : no type generator for gpointer, gpointer

type signalObjectFocusEventDetail struct {
	callback  ObjectSignalFocusEventCallback
	handlerID C.gulong
}

var signalObjectFocusEventId int
var signalObjectFocusEventMap = make(map[int]signalObjectFocusEventDetail)
var signalObjectFocusEventLock sync.RWMutex

// ObjectSignalFocusEventCallback is a callback function for a 'focus-event' signal emitted from a Object.
type ObjectSignalFocusEventCallback func(arg1 bool)

/*
ConnectFocusEvent connects the callback to the 'focus-event' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectFocusEvent to remove it.
*/
func (recv *Object) ConnectFocusEvent(callback ObjectSignalFocusEventCallback) int {
	signalObjectFocusEventLock.Lock()
	defer signalObjectFocusEventLock.Unlock()

	signalObjectFocusEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_focus_event(instance, C.gpointer(uintptr(signalObjectFocusEventId)))

	detail := signalObjectFocusEventDetail{callback, handlerID}
	signalObjectFocusEventMap[signalObjectFocusEventId] = detail

	return signalObjectFocusEventId
}

/*
DisconnectFocusEvent disconnects a callback from the 'focus-event' signal for the Object.

The connectionID should be a value returned from a call to ConnectFocusEvent.
*/
func (recv *Object) DisconnectFocusEvent(connectionID int) {
	signalObjectFocusEventLock.Lock()
	defer signalObjectFocusEventLock.Unlock()

	detail, exists := signalObjectFocusEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectFocusEventMap, connectionID)
}

//export object_focusEventHandler
func object_focusEventHandler(_ *C.GObject, c_arg1 C.gboolean, data C.gpointer) {
	signalObjectFocusEventLock.RLock()
	defer signalObjectFocusEventLock.RUnlock()

	arg1 := c_arg1 == C.TRUE

	index := int(uintptr(data))
	callback := signalObjectFocusEventMap[index].callback
	callback(arg1)
}

// Unsupported signal 'property-change' for Object : unsupported parameter arg1 : no type generator for gpointer, gpointer

type signalObjectStateChangeDetail struct {
	callback  ObjectSignalStateChangeCallback
	handlerID C.gulong
}

var signalObjectStateChangeId int
var signalObjectStateChangeMap = make(map[int]signalObjectStateChangeDetail)
var signalObjectStateChangeLock sync.RWMutex

// ObjectSignalStateChangeCallback is a callback function for a 'state-change' signal emitted from a Object.
type ObjectSignalStateChangeCallback func(arg1 string, arg2 bool)

/*
ConnectStateChange connects the callback to the 'state-change' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectStateChange to remove it.
*/
func (recv *Object) ConnectStateChange(callback ObjectSignalStateChangeCallback) int {
	signalObjectStateChangeLock.Lock()
	defer signalObjectStateChangeLock.Unlock()

	signalObjectStateChangeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_state_change(instance, C.gpointer(uintptr(signalObjectStateChangeId)))

	detail := signalObjectStateChangeDetail{callback, handlerID}
	signalObjectStateChangeMap[signalObjectStateChangeId] = detail

	return signalObjectStateChangeId
}

/*
DisconnectStateChange disconnects a callback from the 'state-change' signal for the Object.

The connectionID should be a value returned from a call to ConnectStateChange.
*/
func (recv *Object) DisconnectStateChange(connectionID int) {
	signalObjectStateChangeLock.Lock()
	defer signalObjectStateChangeLock.Unlock()

	detail, exists := signalObjectStateChangeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectStateChangeMap, connectionID)
}

//export object_stateChangeHandler
func object_stateChangeHandler(_ *C.GObject, c_arg1 *C.gchar, c_arg2 C.gboolean, data C.gpointer) {
	signalObjectStateChangeLock.RLock()
	defer signalObjectStateChangeLock.RUnlock()

	arg1 := C.GoString(c_arg1)

	arg2 := c_arg2 == C.TRUE

	index := int(uintptr(data))
	callback := signalObjectStateChangeMap[index].callback
	callback(arg1, arg2)
}

type signalObjectVisibleDataChangedDetail struct {
	callback  ObjectSignalVisibleDataChangedCallback
	handlerID C.gulong
}

var signalObjectVisibleDataChangedId int
var signalObjectVisibleDataChangedMap = make(map[int]signalObjectVisibleDataChangedDetail)
var signalObjectVisibleDataChangedLock sync.RWMutex

// ObjectSignalVisibleDataChangedCallback is a callback function for a 'visible-data-changed' signal emitted from a Object.
type ObjectSignalVisibleDataChangedCallback func()

/*
ConnectVisibleDataChanged connects the callback to the 'visible-data-changed' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectVisibleDataChanged to remove it.
*/
func (recv *Object) ConnectVisibleDataChanged(callback ObjectSignalVisibleDataChangedCallback) int {
	signalObjectVisibleDataChangedLock.Lock()
	defer signalObjectVisibleDataChangedLock.Unlock()

	signalObjectVisibleDataChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_visible_data_changed(instance, C.gpointer(uintptr(signalObjectVisibleDataChangedId)))

	detail := signalObjectVisibleDataChangedDetail{callback, handlerID}
	signalObjectVisibleDataChangedMap[signalObjectVisibleDataChangedId] = detail

	return signalObjectVisibleDataChangedId
}

/*
DisconnectVisibleDataChanged disconnects a callback from the 'visible-data-changed' signal for the Object.

The connectionID should be a value returned from a call to ConnectVisibleDataChanged.
*/
func (recv *Object) DisconnectVisibleDataChanged(connectionID int) {
	signalObjectVisibleDataChangedLock.Lock()
	defer signalObjectVisibleDataChangedLock.Unlock()

	detail, exists := signalObjectVisibleDataChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectVisibleDataChangedMap, connectionID)
}

//export object_visibleDataChangedHandler
func object_visibleDataChangedHandler(_ *C.GObject, data C.gpointer) {
	signalObjectVisibleDataChangedLock.RLock()
	defer signalObjectVisibleDataChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalObjectVisibleDataChangedMap[index].callback
	callback()
}

// Blacklisted : atk_object_add_relationship

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler (AtkPropertyChangeHandler*) for param handler

// Blacklisted : atk_object_get_description

// Blacklisted : atk_object_get_index_in_parent

// Blacklisted : atk_object_get_layer

// Blacklisted : atk_object_get_mdi_zorder

// Blacklisted : atk_object_get_n_accessible_children

// Blacklisted : atk_object_get_name

// Blacklisted : atk_object_get_parent

// Blacklisted : atk_object_get_role

// Unsupported : atk_object_initialize : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : atk_object_notify_state_change

// Blacklisted : atk_object_peek_parent

// Blacklisted : atk_object_ref_accessible_child

// Blacklisted : atk_object_ref_relation_set

// Blacklisted : atk_object_ref_state_set

// Blacklisted : atk_object_remove_property_change_handler

// Blacklisted : atk_object_remove_relationship

// Blacklisted : atk_object_set_description

// Blacklisted : atk_object_set_name

// Blacklisted : atk_object_set_parent

// Blacklisted : atk_object_set_role

// Blacklisted : AtkObjectFactory

// Blacklisted : AtkPlug

// Blacklisted : AtkRegistry

// Blacklisted : AtkRelation

// Blacklisted : AtkRelationSet

// Blacklisted : AtkSocket

// Blacklisted : AtkStateSet

// Blacklisted : AtkUtil
