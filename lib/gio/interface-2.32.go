// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void networkmonitor_networkChangedHandler(GObject *, gboolean, gpointer);

	static gulong NetworkMonitor_signal_connect_network_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "network-changed", G_CALLBACK(networkmonitor_networkChangedHandler), data);
	}

*/
import "C"

// QueryAction is a wrapper around the C function g_action_group_query_action.
func (recv *ActionGroup) QueryAction(actionName string) (bool, bool, *glib.VariantType, *glib.VariantType, *glib.Variant, *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	var c_enabled C.gboolean

	var c_parameter_type *C.GVariantType

	var c_state_type *C.GVariantType

	var c_state_hint *C.GVariant

	var c_state *C.GVariant

	retC := C.g_action_group_query_action((*C.GActionGroup)(recv.native), c_action_name, &c_enabled, &c_parameter_type, &c_state_type, &c_state_hint, &c_state)
	retGo := retC == C.TRUE

	enabled := c_enabled == C.TRUE

	parameterType := glib.VariantTypeNewFromC(unsafe.Pointer(c_parameter_type))

	stateType := glib.VariantTypeNewFromC(unsafe.Pointer(c_state_type))

	stateHint := glib.VariantNewFromC(unsafe.Pointer(c_state_hint))

	state := glib.VariantNewFromC(unsafe.Pointer(c_state))

	return retGo, enabled, parameterType, stateType, stateHint, state
}

// AddAction is a wrapper around the C function g_action_map_add_action.
func (recv *ActionMap) AddAction(action *Action) {
	c_action := (*C.GAction)(action.ToC())

	C.g_action_map_add_action((*C.GActionMap)(recv.native), c_action)

	return
}

// Unsupported : g_action_map_add_action_entries : unsupported parameter entries :

// LookupAction is a wrapper around the C function g_action_map_lookup_action.
func (recv *ActionMap) LookupAction(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_map_lookup_action((*C.GActionMap)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveAction is a wrapper around the C function g_action_map_remove_action.
func (recv *ActionMap) RemoveAction(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_map_remove_action((*C.GActionMap)(recv.native), c_action_name)

	return
}

// DupObject is a wrapper around the C function g_dbus_interface_dup_object.
func (recv *DBusInterface) DupObject() *DBusObject {
	retC := C.g_dbus_interface_dup_object((*C.GDBusInterface)(recv.native))
	retGo := DBusObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSortKey is a wrapper around the C function g_drive_get_sort_key.
func (recv *Drive) GetSortKey() string {
	retC := C.g_drive_get_sort_key((*C.GDrive)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// FileNewTmp is a wrapper around the C function g_file_new_tmp.
func FileNewTmp(tmpl string) (*File, *FileIOStream, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var c_iostream *C.GFileIOStream

	var cThrowableError *C.GError

	retC := C.g_file_new_tmp(c_tmpl, &c_iostream, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	iostream := FileIOStreamNewFromC(unsafe.Pointer(c_iostream))

	return retGo, iostream, goError
}

// GetSortKey is a wrapper around the C function g_mount_get_sort_key.
func (recv *Mount) GetSortKey() string {
	retC := C.g_mount_get_sort_key((*C.GMount)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Equals compares this NetworkMonitor with another NetworkMonitor, and returns true if they represent the same GObject.
func (recv *NetworkMonitor) Equals(other *NetworkMonitor) bool {
	return other.ToC() == recv.ToC()
}

type signalNetworkMonitorNetworkChangedDetail struct {
	callback  NetworkMonitorSignalNetworkChangedCallback
	handlerID C.gulong
}

var signalNetworkMonitorNetworkChangedId int
var signalNetworkMonitorNetworkChangedMap = make(map[int]signalNetworkMonitorNetworkChangedDetail)
var signalNetworkMonitorNetworkChangedLock sync.RWMutex

// NetworkMonitorSignalNetworkChangedCallback is a callback function for a 'network-changed' signal emitted from a NetworkMonitor.
type NetworkMonitorSignalNetworkChangedCallback func(networkAvailable bool)

/*
ConnectNetworkChanged connects the callback to the 'network-changed' signal for the NetworkMonitor.

The returned value represents the connection, and may be passed to DisconnectNetworkChanged to remove it.
*/
func (recv *NetworkMonitor) ConnectNetworkChanged(callback NetworkMonitorSignalNetworkChangedCallback) int {
	signalNetworkMonitorNetworkChangedLock.Lock()
	defer signalNetworkMonitorNetworkChangedLock.Unlock()

	signalNetworkMonitorNetworkChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.NetworkMonitor_signal_connect_network_changed(instance, C.gpointer(uintptr(signalNetworkMonitorNetworkChangedId)))

	detail := signalNetworkMonitorNetworkChangedDetail{callback, handlerID}
	signalNetworkMonitorNetworkChangedMap[signalNetworkMonitorNetworkChangedId] = detail

	return signalNetworkMonitorNetworkChangedId
}

/*
DisconnectNetworkChanged disconnects a callback from the 'network-changed' signal for the NetworkMonitor.

The connectionID should be a value returned from a call to ConnectNetworkChanged.
*/
func (recv *NetworkMonitor) DisconnectNetworkChanged(connectionID int) {
	signalNetworkMonitorNetworkChangedLock.Lock()
	defer signalNetworkMonitorNetworkChangedLock.Unlock()

	detail, exists := signalNetworkMonitorNetworkChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNetworkMonitorNetworkChangedMap, connectionID)
}

//export networkmonitor_networkChangedHandler
func networkmonitor_networkChangedHandler(_ *C.GObject, c_network_available C.gboolean, data C.gpointer) {
	signalNetworkMonitorNetworkChangedLock.RLock()
	defer signalNetworkMonitorNetworkChangedLock.RUnlock()

	networkAvailable := c_network_available == C.TRUE

	index := int(uintptr(data))
	callback := signalNetworkMonitorNetworkChangedMap[index].callback
	callback(networkAvailable)
}

// NetworkMonitorGetDefault is a wrapper around the C function g_network_monitor_get_default.
func NetworkMonitorGetDefault() *NetworkMonitor {
	retC := C.g_network_monitor_get_default()
	retGo := NetworkMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CanReach is a wrapper around the C function g_network_monitor_can_reach.
func (recv *NetworkMonitor) CanReach(connectable *SocketConnectable, cancellable *Cancellable) (bool, error) {
	c_connectable := (*C.GSocketConnectable)(connectable.ToC())

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_network_monitor_can_reach((*C.GNetworkMonitor)(recv.native), c_connectable, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_network_monitor_can_reach_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CanReachFinish is a wrapper around the C function g_network_monitor_can_reach_finish.
func (recv *NetworkMonitor) CanReachFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_network_monitor_can_reach_finish((*C.GNetworkMonitor)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetNetworkAvailable is a wrapper around the C function g_network_monitor_get_network_available.
func (recv *NetworkMonitor) GetNetworkAvailable() bool {
	retC := C.g_network_monitor_get_network_available((*C.GNetworkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ActivateActionFull is a wrapper around the C function g_remote_action_group_activate_action_full.
func (recv *RemoteActionGroup) ActivateActionFull(actionName string, parameter *glib.Variant, platformData *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	c_platform_data := (*C.GVariant)(C.NULL)
	if platformData != nil {
		c_platform_data = (*C.GVariant)(platformData.ToC())
	}

	C.g_remote_action_group_activate_action_full((*C.GRemoteActionGroup)(recv.native), c_action_name, c_parameter, c_platform_data)

	return
}

// ChangeActionStateFull is a wrapper around the C function g_remote_action_group_change_action_state_full.
func (recv *RemoteActionGroup) ChangeActionStateFull(actionName string, value *glib.Variant, platformData *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	c_platform_data := (*C.GVariant)(C.NULL)
	if platformData != nil {
		c_platform_data = (*C.GVariant)(platformData.ToC())
	}

	C.g_remote_action_group_change_action_state_full((*C.GRemoteActionGroup)(recv.native), c_action_name, c_value, c_platform_data)

	return
}

// GetSortKey is a wrapper around the C function g_volume_get_sort_key.
func (recv *Volume) GetSortKey() string {
	retC := C.g_volume_get_sort_key((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
