// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// ActionMapInterface is a wrapper around the C record GActionMapInterface.
type ActionMapInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for lookup_action
	// no type for add_action
	// no type for remove_action
}

// NetworkMonitorInterface is a wrapper around the C record GNetworkMonitorInterface.
type NetworkMonitorInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for network_changed
	// no type for can_reach
	// no type for can_reach_async
	// no type for can_reach_finish
}

// RemoteActionGroupInterface is a wrapper around the C record GRemoteActionGroupInterface.
type RemoteActionGroupInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for activate_action_full
	// no type for change_action_state_full
}

// Resource is a wrapper around the C record GResource.
type Resource struct {
	native unsafe.Pointer
}

// SettingsSchema is a wrapper around the C record GSettingsSchema.
type SettingsSchema struct {
	native unsafe.Pointer
}

// SettingsSchemaSource is a wrapper around the C record GSettingsSchemaSource.
type SettingsSchemaSource struct {
	native unsafe.Pointer
}
