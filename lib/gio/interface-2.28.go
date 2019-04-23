// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// Blacklisted : g_action_activate

// Blacklisted : g_action_get_enabled

// Blacklisted : g_action_get_name

// Blacklisted : g_action_get_parameter_type

// Blacklisted : g_action_get_state

// Blacklisted : g_action_get_state_hint

// Blacklisted : g_action_get_state_type

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

// Blacklisted : g_action_group_action_added

// Blacklisted : g_action_group_action_enabled_changed

// Blacklisted : g_action_group_action_removed

// Blacklisted : g_action_group_action_state_changed

// Blacklisted : g_action_group_activate_action

// Blacklisted : g_action_group_change_action_state

// Blacklisted : g_action_group_get_action_enabled

// Blacklisted : g_action_group_get_action_parameter_type

// Blacklisted : g_action_group_get_action_state

// Blacklisted : g_action_group_get_action_state_hint

// Blacklisted : g_action_group_get_action_state_type

// Blacklisted : g_action_group_has_action

// Blacklisted : g_action_group_list_actions

// Blacklisted : g_app_info_get_fallback_for_type

// Blacklisted : g_app_info_get_recommended_for_type

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

// Blacklisted : g_pollable_input_stream_can_poll

// Blacklisted : g_pollable_input_stream_create_source

// Blacklisted : g_pollable_input_stream_is_readable

// Blacklisted : g_pollable_input_stream_read_nonblocking

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

// Blacklisted : g_pollable_output_stream_can_poll

// Blacklisted : g_pollable_output_stream_create_source

// Blacklisted : g_pollable_output_stream_is_writable

// Blacklisted : g_pollable_output_stream_write_nonblocking

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

// Blacklisted : g_tls_backend_get_default

// Blacklisted : g_tls_backend_get_certificate_type

// Blacklisted : g_tls_backend_get_client_connection_type

// Blacklisted : g_tls_backend_get_server_connection_type

// Blacklisted : g_tls_backend_supports_tls

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

// Blacklisted : g_tls_client_connection_new

// Blacklisted : g_tls_client_connection_get_accepted_cas

// Blacklisted : g_tls_client_connection_get_server_identity

// Blacklisted : g_tls_client_connection_get_use_ssl3

// Blacklisted : g_tls_client_connection_get_validation_flags

// Blacklisted : g_tls_client_connection_set_server_identity

// Blacklisted : g_tls_client_connection_set_use_ssl3

// Blacklisted : g_tls_client_connection_set_validation_flags

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

// Blacklisted : g_tls_server_connection_new
