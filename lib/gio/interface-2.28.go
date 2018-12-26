// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
/*

	void actiongroup_actionAddedHandler(GObject *, gchar*, gpointer);

	static gulong ActionGroup_signal_connect_action_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-added", G_CALLBACK(actiongroup_actionAddedHandler), data);
	}

*/
/*

	void actiongroup_actionEnabledChangedHandler(GObject *, gchar*, gboolean, gpointer);

	static gulong ActionGroup_signal_connect_action_enabled_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-enabled-changed", G_CALLBACK(actiongroup_actionEnabledChangedHandler), data);
	}

*/
/*

	void actiongroup_actionRemovedHandler(GObject *, gchar*, gpointer);

	static gulong ActionGroup_signal_connect_action_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-removed", G_CALLBACK(actiongroup_actionRemovedHandler), data);
	}

*/
/*

	void actiongroup_actionStateChangedHandler(GObject *, gchar*, GVariant *, gpointer);

	static gulong ActionGroup_signal_connect_action_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-state-changed", G_CALLBACK(actiongroup_actionStateChangedHandler), data);
	}

*/
import "C"

// Activate is a wrapper around the C function g_action_activate.
func (recv *Action) Activate(parameter *glib.Variant) {
	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.g_action_activate((*C.GAction)(recv.native), c_parameter)

	return
}

// GetEnabled is a wrapper around the C function g_action_get_enabled.
func (recv *Action) GetEnabled() bool {
	retC := C.g_action_get_enabled((*C.GAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function g_action_get_name.
func (recv *Action) GetName() string {
	retC := C.g_action_get_name((*C.GAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetParameterType is a wrapper around the C function g_action_get_parameter_type.
func (recv *Action) GetParameterType() *glib.VariantType {
	retC := C.g_action_get_parameter_type((*C.GAction)(recv.native))
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetState is a wrapper around the C function g_action_get_state.
func (recv *Action) GetState() *glib.Variant {
	retC := C.g_action_get_state((*C.GAction)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStateHint is a wrapper around the C function g_action_get_state_hint.
func (recv *Action) GetStateHint() *glib.Variant {
	retC := C.g_action_get_state_hint((*C.GAction)(recv.native))
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStateType is a wrapper around the C function g_action_get_state_type.
func (recv *Action) GetStateType() *glib.VariantType {
	retC := C.g_action_get_state_type((*C.GAction)(recv.native))
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type signalActionGroupActionAddedDetail struct {
	callback  ActionGroupSignalActionAddedCallback
	handlerID C.gulong
}

var signalActionGroupActionAddedId int
var signalActionGroupActionAddedMap = make(map[int]signalActionGroupActionAddedDetail)
var signalActionGroupActionAddedLock sync.RWMutex

// ActionGroupSignalActionAddedCallback is a callback function for a 'action-added' signal emitted from a ActionGroup.
type ActionGroupSignalActionAddedCallback func(actionName string)

/*
ConnectActionAdded connects the callback to the 'action-added' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionAdded to remove it.
*/
func (recv *ActionGroup) ConnectActionAdded(callback ActionGroupSignalActionAddedCallback) int {
	signalActionGroupActionAddedLock.Lock()
	defer signalActionGroupActionAddedLock.Unlock()

	signalActionGroupActionAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_added(instance, C.gpointer(uintptr(signalActionGroupActionAddedId)))

	detail := signalActionGroupActionAddedDetail{callback, handlerID}
	signalActionGroupActionAddedMap[signalActionGroupActionAddedId] = detail

	return signalActionGroupActionAddedId
}

/*
DisconnectActionAdded disconnects a callback from the 'action-added' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionAdded.
*/
func (recv *ActionGroup) DisconnectActionAdded(connectionID int) {
	signalActionGroupActionAddedLock.Lock()
	defer signalActionGroupActionAddedLock.Unlock()

	detail, exists := signalActionGroupActionAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionAddedMap, connectionID)
}

//export actiongroup_actionAddedHandler
func actiongroup_actionAddedHandler(_ *C.GObject, c_action_name *C.gchar, data C.gpointer) {
	signalActionGroupActionAddedLock.RLock()
	defer signalActionGroupActionAddedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	index := int(uintptr(data))
	callback := signalActionGroupActionAddedMap[index].callback
	callback(actionName)
}

type signalActionGroupActionEnabledChangedDetail struct {
	callback  ActionGroupSignalActionEnabledChangedCallback
	handlerID C.gulong
}

var signalActionGroupActionEnabledChangedId int
var signalActionGroupActionEnabledChangedMap = make(map[int]signalActionGroupActionEnabledChangedDetail)
var signalActionGroupActionEnabledChangedLock sync.RWMutex

// ActionGroupSignalActionEnabledChangedCallback is a callback function for a 'action-enabled-changed' signal emitted from a ActionGroup.
type ActionGroupSignalActionEnabledChangedCallback func(actionName string, enabled bool)

/*
ConnectActionEnabledChanged connects the callback to the 'action-enabled-changed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionEnabledChanged to remove it.
*/
func (recv *ActionGroup) ConnectActionEnabledChanged(callback ActionGroupSignalActionEnabledChangedCallback) int {
	signalActionGroupActionEnabledChangedLock.Lock()
	defer signalActionGroupActionEnabledChangedLock.Unlock()

	signalActionGroupActionEnabledChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_enabled_changed(instance, C.gpointer(uintptr(signalActionGroupActionEnabledChangedId)))

	detail := signalActionGroupActionEnabledChangedDetail{callback, handlerID}
	signalActionGroupActionEnabledChangedMap[signalActionGroupActionEnabledChangedId] = detail

	return signalActionGroupActionEnabledChangedId
}

/*
DisconnectActionEnabledChanged disconnects a callback from the 'action-enabled-changed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionEnabledChanged.
*/
func (recv *ActionGroup) DisconnectActionEnabledChanged(connectionID int) {
	signalActionGroupActionEnabledChangedLock.Lock()
	defer signalActionGroupActionEnabledChangedLock.Unlock()

	detail, exists := signalActionGroupActionEnabledChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionEnabledChangedMap, connectionID)
}

//export actiongroup_actionEnabledChangedHandler
func actiongroup_actionEnabledChangedHandler(_ *C.GObject, c_action_name *C.gchar, c_enabled C.gboolean, data C.gpointer) {
	signalActionGroupActionEnabledChangedLock.RLock()
	defer signalActionGroupActionEnabledChangedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	enabled := c_enabled == C.TRUE

	index := int(uintptr(data))
	callback := signalActionGroupActionEnabledChangedMap[index].callback
	callback(actionName, enabled)
}

type signalActionGroupActionRemovedDetail struct {
	callback  ActionGroupSignalActionRemovedCallback
	handlerID C.gulong
}

var signalActionGroupActionRemovedId int
var signalActionGroupActionRemovedMap = make(map[int]signalActionGroupActionRemovedDetail)
var signalActionGroupActionRemovedLock sync.RWMutex

// ActionGroupSignalActionRemovedCallback is a callback function for a 'action-removed' signal emitted from a ActionGroup.
type ActionGroupSignalActionRemovedCallback func(actionName string)

/*
ConnectActionRemoved connects the callback to the 'action-removed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionRemoved to remove it.
*/
func (recv *ActionGroup) ConnectActionRemoved(callback ActionGroupSignalActionRemovedCallback) int {
	signalActionGroupActionRemovedLock.Lock()
	defer signalActionGroupActionRemovedLock.Unlock()

	signalActionGroupActionRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_removed(instance, C.gpointer(uintptr(signalActionGroupActionRemovedId)))

	detail := signalActionGroupActionRemovedDetail{callback, handlerID}
	signalActionGroupActionRemovedMap[signalActionGroupActionRemovedId] = detail

	return signalActionGroupActionRemovedId
}

/*
DisconnectActionRemoved disconnects a callback from the 'action-removed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionRemoved.
*/
func (recv *ActionGroup) DisconnectActionRemoved(connectionID int) {
	signalActionGroupActionRemovedLock.Lock()
	defer signalActionGroupActionRemovedLock.Unlock()

	detail, exists := signalActionGroupActionRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionRemovedMap, connectionID)
}

//export actiongroup_actionRemovedHandler
func actiongroup_actionRemovedHandler(_ *C.GObject, c_action_name *C.gchar, data C.gpointer) {
	signalActionGroupActionRemovedLock.RLock()
	defer signalActionGroupActionRemovedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	index := int(uintptr(data))
	callback := signalActionGroupActionRemovedMap[index].callback
	callback(actionName)
}

type signalActionGroupActionStateChangedDetail struct {
	callback  ActionGroupSignalActionStateChangedCallback
	handlerID C.gulong
}

var signalActionGroupActionStateChangedId int
var signalActionGroupActionStateChangedMap = make(map[int]signalActionGroupActionStateChangedDetail)
var signalActionGroupActionStateChangedLock sync.RWMutex

// ActionGroupSignalActionStateChangedCallback is a callback function for a 'action-state-changed' signal emitted from a ActionGroup.
type ActionGroupSignalActionStateChangedCallback func(actionName string, value *glib.Variant)

/*
ConnectActionStateChanged connects the callback to the 'action-state-changed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionStateChanged to remove it.
*/
func (recv *ActionGroup) ConnectActionStateChanged(callback ActionGroupSignalActionStateChangedCallback) int {
	signalActionGroupActionStateChangedLock.Lock()
	defer signalActionGroupActionStateChangedLock.Unlock()

	signalActionGroupActionStateChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_state_changed(instance, C.gpointer(uintptr(signalActionGroupActionStateChangedId)))

	detail := signalActionGroupActionStateChangedDetail{callback, handlerID}
	signalActionGroupActionStateChangedMap[signalActionGroupActionStateChangedId] = detail

	return signalActionGroupActionStateChangedId
}

/*
DisconnectActionStateChanged disconnects a callback from the 'action-state-changed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionStateChanged.
*/
func (recv *ActionGroup) DisconnectActionStateChanged(connectionID int) {
	signalActionGroupActionStateChangedLock.Lock()
	defer signalActionGroupActionStateChangedLock.Unlock()

	detail, exists := signalActionGroupActionStateChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionStateChangedMap, connectionID)
}

//export actiongroup_actionStateChangedHandler
func actiongroup_actionStateChangedHandler(_ *C.GObject, c_action_name *C.gchar, c_value *C.GVariant, data C.gpointer) {
	signalActionGroupActionStateChangedLock.RLock()
	defer signalActionGroupActionStateChangedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	value := glib.VariantNewFromC(unsafe.Pointer(c_value))

	index := int(uintptr(data))
	callback := signalActionGroupActionStateChangedMap[index].callback
	callback(actionName, value)
}

// ActionAdded is a wrapper around the C function g_action_group_action_added.
func (recv *ActionGroup) ActionAdded(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_added((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// ActionEnabledChanged is a wrapper around the C function g_action_group_action_enabled_changed.
func (recv *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_enabled :=
		boolToGboolean(enabled)

	C.g_action_group_action_enabled_changed((*C.GActionGroup)(recv.native), c_action_name, c_enabled)

	return
}

// ActionRemoved is a wrapper around the C function g_action_group_action_removed.
func (recv *ActionGroup) ActionRemoved(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_removed((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// ActionStateChanged is a wrapper around the C function g_action_group_action_state_changed.
func (recv *ActionGroup) ActionStateChanged(actionName string, state *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_state := (*C.GVariant)(C.NULL)
	if state != nil {
		c_state = (*C.GVariant)(state.ToC())
	}

	C.g_action_group_action_state_changed((*C.GActionGroup)(recv.native), c_action_name, c_state)

	return
}

// ActivateAction is a wrapper around the C function g_action_group_activate_action.
func (recv *ActionGroup) ActivateAction(actionName string, parameter *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.g_action_group_activate_action((*C.GActionGroup)(recv.native), c_action_name, c_parameter)

	return
}

// ChangeActionState is a wrapper around the C function g_action_group_change_action_state.
func (recv *ActionGroup) ChangeActionState(actionName string, value *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_action_group_change_action_state((*C.GActionGroup)(recv.native), c_action_name, c_value)

	return
}

// GetActionEnabled is a wrapper around the C function g_action_group_get_action_enabled.
func (recv *ActionGroup) GetActionEnabled(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_enabled((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// GetActionParameterType is a wrapper around the C function g_action_group_get_action_parameter_type.
func (recv *ActionGroup) GetActionParameterType(actionName string) *glib.VariantType {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_parameter_type((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionState is a wrapper around the C function g_action_group_get_action_state.
func (recv *ActionGroup) GetActionState(actionName string) *glib.Variant {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionStateHint is a wrapper around the C function g_action_group_get_action_state_hint.
func (recv *ActionGroup) GetActionStateHint(actionName string) *glib.Variant {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state_hint((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionStateType is a wrapper around the C function g_action_group_get_action_state_type.
func (recv *ActionGroup) GetActionStateType(actionName string) *glib.VariantType {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state_type((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// HasAction is a wrapper around the C function g_action_group_has_action.
func (recv *ActionGroup) HasAction(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_has_action((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_action_group_list_actions : no return type

// AppInfoGetFallbackForType is a wrapper around the C function g_app_info_get_fallback_for_type.
func AppInfoGetFallbackForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_fallback_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetRecommendedForType is a wrapper around the C function g_app_info_get_recommended_for_type.
func AppInfoGetRecommendedForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_recommended_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PollableInputStream is a wrapper around the C record GPollableInputStream.
type PollableInputStream struct {
	native *C.GPollableInputStream
}

func PollableInputStreamNewFromC(u unsafe.Pointer) *PollableInputStream {
	c := (*C.GPollableInputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStream{native: c}

	return g
}

func (recv *PollableInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableInputStream with another PollableInputStream, and returns true if they represent the same GObject.
func (recv *PollableInputStream) Equals(other *PollableInputStream) bool {
	return other.ToC() == recv.ToC()
}

// CanPoll is a wrapper around the C function g_pollable_input_stream_can_poll.
func (recv *PollableInputStream) CanPoll() bool {
	retC := C.g_pollable_input_stream_can_poll((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_input_stream_create_source.
func (recv *PollableInputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_input_stream_create_source((*C.GPollableInputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsReadable is a wrapper around the C function g_pollable_input_stream_is_readable.
func (recv *PollableInputStream) IsReadable() bool {
	retC := C.g_pollable_input_stream_is_readable((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReadNonblocking is a wrapper around the C function g_pollable_input_stream_read_nonblocking.
func (recv *PollableInputStream) ReadNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_input_stream_read_nonblocking((*C.GPollableInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PollableOutputStream is a wrapper around the C record GPollableOutputStream.
type PollableOutputStream struct {
	native *C.GPollableOutputStream
}

func PollableOutputStreamNewFromC(u unsafe.Pointer) *PollableOutputStream {
	c := (*C.GPollableOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStream{native: c}

	return g
}

func (recv *PollableOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableOutputStream with another PollableOutputStream, and returns true if they represent the same GObject.
func (recv *PollableOutputStream) Equals(other *PollableOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// CanPoll is a wrapper around the C function g_pollable_output_stream_can_poll.
func (recv *PollableOutputStream) CanPoll() bool {
	retC := C.g_pollable_output_stream_can_poll((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_output_stream_create_source.
func (recv *PollableOutputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_output_stream_create_source((*C.GPollableOutputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsWritable is a wrapper around the C function g_pollable_output_stream_is_writable.
func (recv *PollableOutputStream) IsWritable() bool {
	retC := C.g_pollable_output_stream_is_writable((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WriteNonblocking is a wrapper around the C function g_pollable_output_stream_write_nonblocking.
func (recv *PollableOutputStream) WriteNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_output_stream_write_nonblocking((*C.GPollableOutputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsBackend is a wrapper around the C record GTlsBackend.
type TlsBackend struct {
	native *C.GTlsBackend
}

func TlsBackendNewFromC(u unsafe.Pointer) *TlsBackend {
	c := (*C.GTlsBackend)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackend{native: c}

	return g
}

func (recv *TlsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsBackend with another TlsBackend, and returns true if they represent the same GObject.
func (recv *TlsBackend) Equals(other *TlsBackend) bool {
	return other.ToC() == recv.ToC()
}

// TlsBackendGetDefault is a wrapper around the C function g_tls_backend_get_default.
func TlsBackendGetDefault() *TlsBackend {
	retC := C.g_tls_backend_get_default()
	retGo := TlsBackendNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCertificateType is a wrapper around the C function g_tls_backend_get_certificate_type.
func (recv *TlsBackend) GetCertificateType() gobject.Type {
	retC := C.g_tls_backend_get_certificate_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetClientConnectionType is a wrapper around the C function g_tls_backend_get_client_connection_type.
func (recv *TlsBackend) GetClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetServerConnectionType is a wrapper around the C function g_tls_backend_get_server_connection_type.
func (recv *TlsBackend) GetServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// SupportsTls is a wrapper around the C function g_tls_backend_supports_tls.
func (recv *TlsBackend) SupportsTls() bool {
	retC := C.g_tls_backend_supports_tls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TlsClientConnection is a wrapper around the C record GTlsClientConnection.
type TlsClientConnection struct {
	native *C.GTlsClientConnection
}

func TlsClientConnectionNewFromC(u unsafe.Pointer) *TlsClientConnection {
	c := (*C.GTlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnection{native: c}

	return g
}

func (recv *TlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsClientConnection with another TlsClientConnection, and returns true if they represent the same GObject.
func (recv *TlsClientConnection) Equals(other *TlsClientConnection) bool {
	return other.ToC() == recv.ToC()
}

// TlsClientConnectionNew is a wrapper around the C function g_tls_client_connection_new.
func TlsClientConnectionNew(baseIoStream *IOStream, serverIdentity *SocketConnectable) (*TlsClientConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_server_identity := (*C.GSocketConnectable)(serverIdentity.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_client_connection_new(c_base_io_stream, c_server_identity, &cThrowableError)
	retGo := TlsClientConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAcceptedCas is a wrapper around the C function g_tls_client_connection_get_accepted_cas.
func (recv *TlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_tls_client_connection_get_accepted_cas((*C.GTlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetServerIdentity is a wrapper around the C function g_tls_client_connection_get_server_identity.
func (recv *TlsClientConnection) GetServerIdentity() *SocketConnectable {
	retC := C.g_tls_client_connection_get_server_identity((*C.GTlsClientConnection)(recv.native))
	retGo := SocketConnectableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUseSsl3 is a wrapper around the C function g_tls_client_connection_get_use_ssl3.
func (recv *TlsClientConnection) GetUseSsl3() bool {
	retC := C.g_tls_client_connection_get_use_ssl3((*C.GTlsClientConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValidationFlags is a wrapper around the C function g_tls_client_connection_get_validation_flags.
func (recv *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_tls_client_connection_get_validation_flags((*C.GTlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// SetServerIdentity is a wrapper around the C function g_tls_client_connection_set_server_identity.
func (recv *TlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	C.g_tls_client_connection_set_server_identity((*C.GTlsClientConnection)(recv.native), c_identity)

	return
}

// SetUseSsl3 is a wrapper around the C function g_tls_client_connection_set_use_ssl3.
func (recv *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	c_use_ssl3 :=
		boolToGboolean(useSsl3)

	C.g_tls_client_connection_set_use_ssl3((*C.GTlsClientConnection)(recv.native), c_use_ssl3)

	return
}

// SetValidationFlags is a wrapper around the C function g_tls_client_connection_set_validation_flags.
func (recv *TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_tls_client_connection_set_validation_flags((*C.GTlsClientConnection)(recv.native), c_flags)

	return
}

// TlsServerConnection is a wrapper around the C record GTlsServerConnection.
type TlsServerConnection struct {
	native *C.GTlsServerConnection
}

func TlsServerConnectionNewFromC(u unsafe.Pointer) *TlsServerConnection {
	c := (*C.GTlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnection{native: c}

	return g
}

func (recv *TlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsServerConnection with another TlsServerConnection, and returns true if they represent the same GObject.
func (recv *TlsServerConnection) Equals(other *TlsServerConnection) bool {
	return other.ToC() == recv.ToC()
}

// TlsServerConnectionNew is a wrapper around the C function g_tls_server_connection_new.
func TlsServerConnectionNew(baseIoStream *IOStream, certificate *TlsCertificate) (*TlsServerConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_server_connection_new(c_base_io_stream, c_certificate, &cThrowableError)
	retGo := TlsServerConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
