// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// ActionGroupInterface is a wrapper around the C record GActionGroupInterface.
type ActionGroupInterface struct {
	native *C.GActionGroupInterface
	// g_iface : record
	// no type for has_action
	// no type for list_actions
	// no type for get_action_enabled
	// no type for get_action_parameter_type
	// no type for get_action_state_type
	// no type for get_action_state_hint
	// no type for get_action_state
	// no type for change_action_state
	// no type for activate_action
	// no type for action_added
	// no type for action_removed
	// no type for action_enabled_changed
	// no type for action_state_changed
	// no type for query_action
}

func ActionGroupInterfaceNewFromC(u unsafe.Pointer) *ActionGroupInterface {
	c := (*C.GActionGroupInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroupInterface{native: c}

	return g
}

func (recv *ActionGroupInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionInterface is a wrapper around the C record GActionInterface.
type ActionInterface struct {
	native *C.GActionInterface
	// g_iface : record
	// no type for get_name
	// no type for get_parameter_type
	// no type for get_state_type
	// no type for get_state_hint
	// no type for get_enabled
	// no type for get_state
	// no type for change_state
	// no type for activate
}

func ActionInterfaceNewFromC(u unsafe.Pointer) *ActionInterface {
	c := (*C.GActionInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionInterface{native: c}

	return g
}

func (recv *ActionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationClass is a wrapper around the C record GApplicationClass.
type ApplicationClass struct {
	native *C.GApplicationClass
	// Private : parent_class
	// no type for startup
	// no type for activate
	// no type for open
	// no type for command_line
	// no type for local_command_line
	// no type for before_emit
	// no type for after_emit
	// no type for add_platform_data
	// no type for quit_mainloop
	// no type for run_mainloop
	// no type for shutdown
	// no type for dbus_register
	// no type for dbus_unregister
	// no type for handle_local_options
	// Private : padding
}

func ApplicationClassNewFromC(u unsafe.Pointer) *ApplicationClass {
	c := (*C.GApplicationClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationClass{native: c}

	return g
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationCommandLineClass is a wrapper around the C record GApplicationCommandLineClass.
type ApplicationCommandLineClass struct {
	native *C.GApplicationCommandLineClass
	// Private : parent_class
	// no type for print_literal
	// no type for printerr_literal
	// no type for get_stdin
	// Private : padding
}

func ApplicationCommandLineClassNewFromC(u unsafe.Pointer) *ApplicationCommandLineClass {
	c := (*C.GApplicationCommandLineClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLineClass{native: c}

	return g
}

func (recv *ApplicationCommandLineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_io_extension_get_type : no return generator

// Unsupported : g_io_extension_point_get_required_type : no return generator

// Unsupported : g_io_extension_point_set_required_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// PollableInputStreamInterface is a wrapper around the C record GPollableInputStreamInterface.
type PollableInputStreamInterface struct {
	native *C.GPollableInputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_readable
	// no type for create_source
	// no type for read_nonblocking
}

func PollableInputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableInputStreamInterface {
	c := (*C.GPollableInputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStreamInterface{native: c}

	return g
}

func (recv *PollableInputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollableOutputStreamInterface is a wrapper around the C record GPollableOutputStreamInterface.
type PollableOutputStreamInterface struct {
	native *C.GPollableOutputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_writable
	// no type for create_source
	// no type for write_nonblocking
}

func PollableOutputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableOutputStreamInterface {
	c := (*C.GPollableOutputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStreamInterface{native: c}

	return g
}

func (recv *PollableOutputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_resource_enumerate_children : no return type

// Unsupported : g_resource_get_info : unsupported parameter size : no type generator for gsize, gsize*

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// Unsupported : g_settings_schema_list_children : no return type

// Unsupported : g_settings_schema_list_keys : no return type

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : no param type

// TlsBackendInterface is a wrapper around the C record GTlsBackendInterface.
type TlsBackendInterface struct {
	native *C.GTlsBackendInterface
	// g_iface : record
	// no type for supports_tls
	// no type for get_certificate_type
	// no type for get_client_connection_type
	// no type for get_server_connection_type
	// no type for get_file_database_type
	// no type for get_default_database
	// no type for supports_dtls
	// no type for get_dtls_client_connection_type
	// no type for get_dtls_server_connection_type
}

func TlsBackendInterfaceNewFromC(u unsafe.Pointer) *TlsBackendInterface {
	c := (*C.GTlsBackendInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackendInterface{native: c}

	return g
}

func (recv *TlsBackendInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_unix_mount_point_guess_icon : no return generator

// Unsupported : g_unix_mount_point_guess_symbolic_icon : no return generator
