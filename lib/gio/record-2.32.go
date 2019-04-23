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

// Equals compares this ActionMapInterface with another ActionMapInterface, and returns true if they represent the same GObject.
func (recv *ActionMapInterface) Equals(other *ActionMapInterface) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_file_attribute_matcher_to_string

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

// Equals compares this NetworkMonitorInterface with another NetworkMonitorInterface, and returns true if they represent the same GObject.
func (recv *NetworkMonitorInterface) Equals(other *NetworkMonitorInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this RemoteActionGroupInterface with another RemoteActionGroupInterface, and returns true if they represent the same GObject.
func (recv *RemoteActionGroupInterface) Equals(other *RemoteActionGroupInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this Resource with another Resource, and returns true if they represent the same GObject.
func (recv *Resource) Equals(other *Resource) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_resource_new_from_data

// Blacklisted : g_resource_load

// Blacklisted : g_resources_register

// Blacklisted : g_resources_unregister

// Blacklisted : g_resource_enumerate_children

// Blacklisted : g_resource_get_info

// Blacklisted : g_resource_lookup_data

// Blacklisted : g_resource_open_stream

// Blacklisted : g_resource_ref

// Blacklisted : g_resource_unref

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

// Equals compares this SettingsSchema with another SettingsSchema, and returns true if they represent the same GObject.
func (recv *SettingsSchema) Equals(other *SettingsSchema) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_settings_schema_get_id

// Blacklisted : g_settings_schema_get_path

// Blacklisted : g_settings_schema_ref

// Blacklisted : g_settings_schema_unref

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

// Equals compares this SettingsSchemaSource with another SettingsSchemaSource, and returns true if they represent the same GObject.
func (recv *SettingsSchemaSource) Equals(other *SettingsSchemaSource) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_settings_schema_source_new_from_directory

// Blacklisted : g_settings_schema_source_get_default

// Blacklisted : g_settings_schema_source_lookup

// Blacklisted : g_settings_schema_source_ref

// Blacklisted : g_settings_schema_source_unref

// Blacklisted : g_static_resource_fini

// Blacklisted : g_static_resource_get_resource

// Init is a wrapper around the C function g_static_resource_init.
func (recv *StaticResource) Init() {
	C.g_static_resource_init((*C.GStaticResource)(recv.native))

	return
}

// Blacklisted : g_unix_mount_point_get_options
