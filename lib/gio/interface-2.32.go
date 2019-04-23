// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
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

// Blacklisted : g_action_group_query_action

// Blacklisted : g_action_map_add_action

// Unsupported : g_action_map_add_action_entries : unsupported parameter entries :

// Blacklisted : g_action_map_lookup_action

// Blacklisted : g_action_map_remove_action

// Blacklisted : g_dbus_interface_dup_object

// Blacklisted : g_drive_get_sort_key

// Blacklisted : g_file_new_tmp

// Blacklisted : g_mount_get_sort_key

// NetworkMonitor is a wrapper around the C record GNetworkMonitor.
type NetworkMonitor struct {
	native *C.GNetworkMonitor
}

func NetworkMonitorNewFromC(u unsafe.Pointer) *NetworkMonitor {
	c := (*C.GNetworkMonitor)(u)
	if c == nil {
		return nil
	}

	g := &NetworkMonitor{native: c}

	return g
}

func (recv *NetworkMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// Blacklisted : g_network_monitor_get_default

// Blacklisted : g_network_monitor_can_reach

// Unsupported : g_network_monitor_can_reach_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_network_monitor_can_reach_finish

// Blacklisted : g_network_monitor_get_network_available

// Blacklisted : g_remote_action_group_activate_action_full

// Blacklisted : g_remote_action_group_change_action_state_full

// Blacklisted : g_volume_get_sort_key
