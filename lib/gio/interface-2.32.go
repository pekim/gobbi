// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
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
import "C"

// Unsupported : g_action_group_query_action : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_action_map_add_action : unsupported parameter action : no type generator for Action (GAction*) for param action

// Unsupported : g_action_map_add_action_entries : unsupported parameter entries : no type generator for ActionEntry () for array param entries

// Unsupported : g_action_map_lookup_action : no return generator

// RemoveAction is a wrapper around the C function g_action_map_remove_action.
func (recv *ActionMap) RemoveAction(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_map_remove_action((*C.GActionMap)(recv.native), c_action_name)

	return
}

// Unsupported : g_dbus_interface_dup_object : no return generator

// GetSortKey is a wrapper around the C function g_drive_get_sort_key.
func (recv *Drive) GetSortKey() string {
	retC := C.g_drive_get_sort_key((*C.GDrive)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSortKey is a wrapper around the C function g_mount_get_sort_key.
func (recv *Mount) GetSortKey() string {
	retC := C.g_mount_get_sort_key((*C.GMount)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

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

// Unsupported : g_network_monitor_can_reach : unsupported parameter connectable : no type generator for SocketConnectable (GSocketConnectable*) for param connectable

// Unsupported : g_network_monitor_can_reach_async : unsupported parameter connectable : no type generator for SocketConnectable (GSocketConnectable*) for param connectable

// Unsupported : g_network_monitor_can_reach_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// GetNetworkAvailable is a wrapper around the C function g_network_monitor_get_network_available.
func (recv *NetworkMonitor) GetNetworkAvailable() bool {
	retC := C.g_network_monitor_get_network_available((*C.GNetworkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_remote_action_group_activate_action_full : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : g_remote_action_group_change_action_state_full : unsupported parameter value : Blacklisted record : GVariant

// GetSortKey is a wrapper around the C function g_volume_get_sort_key.
func (recv *Volume) GetSortKey() string {
	retC := C.g_volume_get_sort_key((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
