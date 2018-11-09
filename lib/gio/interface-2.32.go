// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_network_monitor_can_reach_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CanReachFinish is a wrapper around the C function g_network_monitor_can_reach_finish.
func (recv *NetworkMonitor) CanReachFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_network_monitor_can_reach_finish((*C.GNetworkMonitor)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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
