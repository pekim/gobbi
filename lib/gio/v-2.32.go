// Code generated - DO NOT EDIT.
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import (
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// FILE_ATTRIBUTE_FILESYSTEM_USED is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_USED.
//
// since 2.32
const FILE_ATTRIBUTE_FILESYSTEM_USED = "filesystem::used"

// MENU_ATTRIBUTE_ACTION is a representation of the C constant G_MENU_ATTRIBUTE_ACTION.
//
// since 2.32
const MENU_ATTRIBUTE_ACTION = "action"

// MENU_ATTRIBUTE_LABEL is a representation of the C constant G_MENU_ATTRIBUTE_LABEL.
//
// since 2.32
const MENU_ATTRIBUTE_LABEL = "label"

// MENU_ATTRIBUTE_TARGET is a representation of the C constant G_MENU_ATTRIBUTE_TARGET.
//
// since 2.32
const MENU_ATTRIBUTE_TARGET = "target"

// MENU_LINK_SECTION is a representation of the C constant G_MENU_LINK_SECTION.
//
// since 2.32
const MENU_LINK_SECTION = "section"

// MENU_LINK_SUBMENU is a representation of the C constant G_MENU_LINK_SUBMENU.
//
// since 2.32
const MENU_LINK_SUBMENU = "submenu"

// ResourceFlags is a representation of the C bitfield GResourceFlags.
type ResourceFlags int

// ResourceFlags_none is a representation of the C bitfield member G_RESOURCE_FLAGS_NONE.
const ResourceFlags_none = ResourceFlags(0)

// ResourceFlags_compressed is a representation of the C bitfield member G_RESOURCE_FLAGS_COMPRESSED.
const ResourceFlags_compressed = ResourceFlags(1)

// ResourceLookupFlags is a representation of the C bitfield GResourceLookupFlags.
type ResourceLookupFlags int

// ResourceLookupFlags_none is a representation of the C bitfield member G_RESOURCE_LOOKUP_FLAGS_NONE.
const ResourceLookupFlags_none = ResourceLookupFlags(0)

// ResourceError is a representation of the C enumeration GResourceError.
type ResourceError int

// ResourceError_not_found is a representation of the C enumeration member G_RESOURCE_ERROR_NOT_FOUND.
const ResourceError_not_found = ResourceError(0)

// ResourceError_internal is a representation of the C enumeration member G_RESOURCE_ERROR_INTERNAL.
const ResourceError_internal = ResourceError(1)

// SocketClientEvent is a representation of the C enumeration GSocketClientEvent.
type SocketClientEvent int

// SocketClientEvent_resolving is a representation of the C enumeration member G_SOCKET_CLIENT_RESOLVING.
const SocketClientEvent_resolving = SocketClientEvent(0)

// SocketClientEvent_resolved is a representation of the C enumeration member G_SOCKET_CLIENT_RESOLVED.
const SocketClientEvent_resolved = SocketClientEvent(1)

// SocketClientEvent_connecting is a representation of the C enumeration member G_SOCKET_CLIENT_CONNECTING.
const SocketClientEvent_connecting = SocketClientEvent(2)

// SocketClientEvent_connected is a representation of the C enumeration member G_SOCKET_CLIENT_CONNECTED.
const SocketClientEvent_connected = SocketClientEvent(3)

// SocketClientEvent_proxy_negotiating is a representation of the C enumeration member G_SOCKET_CLIENT_PROXY_NEGOTIATING.
const SocketClientEvent_proxy_negotiating = SocketClientEvent(4)

// SocketClientEvent_proxy_negotiated is a representation of the C enumeration member G_SOCKET_CLIENT_PROXY_NEGOTIATED.
const SocketClientEvent_proxy_negotiated = SocketClientEvent(5)

// SocketClientEvent_tls_handshaking is a representation of the C enumeration member G_SOCKET_CLIENT_TLS_HANDSHAKING.
const SocketClientEvent_tls_handshaking = SocketClientEvent(6)

// SocketClientEvent_tls_handshaked is a representation of the C enumeration member G_SOCKET_CLIENT_TLS_HANDSHAKED.
const SocketClientEvent_tls_handshaked = SocketClientEvent(7)

// SocketClientEvent_complete is a representation of the C enumeration member G_SOCKET_CLIENT_COMPLETE.
const SocketClientEvent_complete = SocketClientEvent(8)

// UNSUPPORTED : g_action_parse_detailed_name : throws

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// ResourcesRegister wraps the C function g_resources_register.
//
// since 2.32
func ResourcesRegister(resource *Resource) {
	sys_resource := resource.ToC()
	gio.Fn_g_resources_register(sys_resource)
}

// ResourcesUnregister wraps the C function g_resources_unregister.
//
// since 2.32
func ResourcesUnregister(resource *Resource) {
	sys_resource := resource.ToC()
	gio.Fn_g_resources_unregister(sys_resource)
}

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UNSUPPORTED : g_unix_mount_at : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_for : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// ActionMapInterface is a representation of the C record GActionMapInterface.
//
// since 2.32
type ActionMapInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionMapInterface that represents the ActionMapInterface.
func (recv *ActionMapInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionMapInterfaceNewFromC creates a new ActionMapInterface from a pointer to the C GActionMapInterface that represents the ActionMapInterface.
func ActionMapInterfaceNewFromC(native unsafe.Pointer) *ActionMapInterface {
	return &ActionMapInterface{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// NetworkMonitorInterface is a representation of the C record GNetworkMonitorInterface.
//
// since 2.32
type NetworkMonitorInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkMonitorInterface that represents the NetworkMonitorInterface.
func (recv *NetworkMonitorInterface) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkMonitorInterfaceNewFromC creates a new NetworkMonitorInterface from a pointer to the C GNetworkMonitorInterface that represents the NetworkMonitorInterface.
func NetworkMonitorInterfaceNewFromC(native unsafe.Pointer) *NetworkMonitorInterface {
	return &NetworkMonitorInterface{native: native}
}

// RemoteActionGroupInterface is a representation of the C record GRemoteActionGroupInterface.
//
// since 2.32
type RemoteActionGroupInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GRemoteActionGroupInterface that represents the RemoteActionGroupInterface.
func (recv *RemoteActionGroupInterface) ToC() unsafe.Pointer {
	return recv.native
}

// RemoteActionGroupInterfaceNewFromC creates a new RemoteActionGroupInterface from a pointer to the C GRemoteActionGroupInterface that represents the RemoteActionGroupInterface.
func RemoteActionGroupInterfaceNewFromC(native unsafe.Pointer) *RemoteActionGroupInterface {
	return &RemoteActionGroupInterface{native: native}
}

// Resource is a representation of the C record GResource.
//
// since 2.32
type Resource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResource that represents the Resource.
func (recv *Resource) ToC() unsafe.Pointer {
	return recv.native
}

// ResourceNewFromC creates a new Resource from a pointer to the C GResource that represents the Resource.
func ResourceNewFromC(native unsafe.Pointer) *Resource {
	return &Resource{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// SettingsSchema is a representation of the C record GSettingsSchema.
//
// since 2.32
type SettingsSchema struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchema that represents the SettingsSchema.
func (recv *SettingsSchema) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaNewFromC creates a new SettingsSchema from a pointer to the C GSettingsSchema that represents the SettingsSchema.
func SettingsSchemaNewFromC(native unsafe.Pointer) *SettingsSchema {
	return &SettingsSchema{native: native}
}

// SettingsSchemaSource is a representation of the C record GSettingsSchemaSource.
//
// since 2.32
type SettingsSchemaSource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchemaSource that represents the SettingsSchemaSource.
func (recv *SettingsSchemaSource) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaSourceNewFromC creates a new SettingsSchemaSource from a pointer to the C GSettingsSchemaSource that represents the SettingsSchemaSource.
func SettingsSchemaSourceNewFromC(native unsafe.Pointer) *SettingsSchemaSource {
	return &SettingsSchemaSource{native: native}
}

// InetAddressMask is a representation of the C record GInetAddressMask.
//
// since 2.32
type InetAddressMask struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMask that represents the InetAddressMask.
func (recv *InetAddressMask) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskNewFromC creates a new InetAddressMask from a pointer to the C GInetAddressMask that represents the InetAddressMask.
func InetAddressMaskNewFromC(native unsafe.Pointer) *InetAddressMask {
	return &InetAddressMask{native: native}
}

// Menu is a representation of the C record GMenu.
//
// since 2.32
type Menu struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenu that represents the Menu.
func (recv *Menu) ToC() unsafe.Pointer {
	return recv.native
}

// MenuNewFromC creates a new Menu from a pointer to the C GMenu that represents the Menu.
func MenuNewFromC(native unsafe.Pointer) *Menu {
	return &Menu{native: native}
}

// MenuAttributeIter is a representation of the C record GMenuAttributeIter.
//
// since 2.32
type MenuAttributeIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIter that represents the MenuAttributeIter.
func (recv *MenuAttributeIter) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterNewFromC creates a new MenuAttributeIter from a pointer to the C GMenuAttributeIter that represents the MenuAttributeIter.
func MenuAttributeIterNewFromC(native unsafe.Pointer) *MenuAttributeIter {
	return &MenuAttributeIter{native: native}
}

// MenuItem is a representation of the C record GMenuItem.
//
// since 2.32
type MenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuItem that represents the MenuItem.
func (recv *MenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemNewFromC creates a new MenuItem from a pointer to the C GMenuItem that represents the MenuItem.
func MenuItemNewFromC(native unsafe.Pointer) *MenuItem {
	return &MenuItem{native: native}
}

// MenuLinkIter is a representation of the C record GMenuLinkIter.
//
// since 2.32
type MenuLinkIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIter that represents the MenuLinkIter.
func (recv *MenuLinkIter) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterNewFromC creates a new MenuLinkIter from a pointer to the C GMenuLinkIter that represents the MenuLinkIter.
func MenuLinkIterNewFromC(native unsafe.Pointer) *MenuLinkIter {
	return &MenuLinkIter{native: native}
}

// MenuModel is a representation of the C record GMenuModel.
//
// since 2.32
type MenuModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModel that represents the MenuModel.
func (recv *MenuModel) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelNewFromC creates a new MenuModel from a pointer to the C GMenuModel that represents the MenuModel.
func MenuModelNewFromC(native unsafe.Pointer) *MenuModel {
	return &MenuModel{native: native}
}

// ActionMap is a representation of the C interface GActionMap.
//
// since 2.32
type ActionMap struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionMap that represents the ActionMap.
func (recv *ActionMap) ToC() unsafe.Pointer {
	return recv.native
}

// ActionMapNewFromC creates a new ActionMap from a pointer to the C GActionMap that represents the ActionMap.
func ActionMapNewFromC(native unsafe.Pointer) *ActionMap {
	return &ActionMap{native: native}
}

// NetworkMonitor is a representation of the C interface GNetworkMonitor.
//
// since 2.32
type NetworkMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkMonitor that represents the NetworkMonitor.
func (recv *NetworkMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkMonitorNewFromC creates a new NetworkMonitor from a pointer to the C GNetworkMonitor that represents the NetworkMonitor.
func NetworkMonitorNewFromC(native unsafe.Pointer) *NetworkMonitor {
	return &NetworkMonitor{native: native}
}

// RemoteActionGroup is a representation of the C interface GRemoteActionGroup.
//
// since 2.32
type RemoteActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GRemoteActionGroup that represents the RemoteActionGroup.
func (recv *RemoteActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// RemoteActionGroupNewFromC creates a new RemoteActionGroup from a pointer to the C GRemoteActionGroup that represents the RemoteActionGroup.
func RemoteActionGroupNewFromC(native unsafe.Pointer) *RemoteActionGroup {
	return &RemoteActionGroup{native: native}
}
