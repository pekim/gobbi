// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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
import "C"

// ActionMapInterface is a wrapper around the C record GActionMapInterface.
type ActionMapInterface struct {
	native *C.GActionMapInterface
	// g_iface : record
	// no type for lookup_action
	// no type for add_action
	// no type for remove_action
}

func ActionMapInterfaceNewFromC(u unsafe.Pointer) *ActionMapInterface {
	c := (*C.GActionMapInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionMapInterface{native: c}

	return g
}

func (recv *ActionMapInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkMonitorInterface is a wrapper around the C record GNetworkMonitorInterface.
type NetworkMonitorInterface struct {
	native *C.GNetworkMonitorInterface
	// g_iface : record
	// no type for network_changed
	// no type for can_reach
	// no type for can_reach_async
	// no type for can_reach_finish
}

func NetworkMonitorInterfaceNewFromC(u unsafe.Pointer) *NetworkMonitorInterface {
	c := (*C.GNetworkMonitorInterface)(u)
	if c == nil {
		return nil
	}

	g := &NetworkMonitorInterface{native: c}

	return g
}

func (recv *NetworkMonitorInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RemoteActionGroupInterface is a wrapper around the C record GRemoteActionGroupInterface.
type RemoteActionGroupInterface struct {
	native *C.GRemoteActionGroupInterface
	// g_iface : record
	// no type for activate_action_full
	// no type for change_action_state_full
}

func RemoteActionGroupInterfaceNewFromC(u unsafe.Pointer) *RemoteActionGroupInterface {
	c := (*C.GRemoteActionGroupInterface)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroupInterface{native: c}

	return g
}

func (recv *RemoteActionGroupInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Resource is a wrapper around the C record GResource.
type Resource struct {
	native *C.GResource
}

func ResourceNewFromC(u unsafe.Pointer) *Resource {
	c := (*C.GResource)(u)
	if c == nil {
		return nil
	}

	g := &Resource{native: c}

	return g
}

func (recv *Resource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsSchema is a wrapper around the C record GSettingsSchema.
type SettingsSchema struct {
	native *C.GSettingsSchema
}

func SettingsSchemaNewFromC(u unsafe.Pointer) *SettingsSchema {
	c := (*C.GSettingsSchema)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchema{native: c}

	return g
}

func (recv *SettingsSchema) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsSchemaSource is a wrapper around the C record GSettingsSchemaSource.
type SettingsSchemaSource struct {
	native *C.GSettingsSchemaSource
}

func SettingsSchemaSourceNewFromC(u unsafe.Pointer) *SettingsSchemaSource {
	c := (*C.GSettingsSchemaSource)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchemaSource{native: c}

	return g
}

func (recv *SettingsSchemaSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
