// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// DBusInterfaceIface is a wrapper around the C record GDBusInterfaceIface.
type DBusInterfaceIface struct {
	native unsafe.Pointer
	// parent_iface : record
	// no type for get_info
	// no type for get_object
	// no type for set_object
	// no type for dup_object
}

// DBusInterfaceSkeletonClass is a wrapper around the C record GDBusInterfaceSkeletonClass.
type DBusInterfaceSkeletonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_info
	// no type for get_vtable
	// no type for get_properties
	// no type for flush
	// Private : vfunc_padding
	// no type for g_authorize_method
	// Private : signal_padding
}

// DBusObjectIface is a wrapper around the C record GDBusObjectIface.
type DBusObjectIface struct {
	native unsafe.Pointer
	// parent_iface : record
	// no type for get_object_path
	// no type for get_interfaces
	// no type for get_interface
	// no type for interface_added
	// no type for interface_removed
}

// DBusObjectManagerClientClass is a wrapper around the C record GDBusObjectManagerClientClass.
type DBusObjectManagerClientClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for interface_proxy_signal
	// no type for interface_proxy_properties_changed
	// Private : padding
}

// DBusObjectManagerIface is a wrapper around the C record GDBusObjectManagerIface.
type DBusObjectManagerIface struct {
	native unsafe.Pointer
	// parent_iface : record
	// no type for get_object_path
	// no type for get_objects
	// no type for get_object
	// no type for get_interface
	// no type for object_added
	// no type for object_removed
	// no type for interface_added
	// no type for interface_removed
}

// DBusObjectManagerServerClass is a wrapper around the C record GDBusObjectManagerServerClass.
type DBusObjectManagerServerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

// DBusObjectProxyClass is a wrapper around the C record GDBusObjectProxyClass.
type DBusObjectProxyClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

// DBusObjectSkeletonClass is a wrapper around the C record GDBusObjectSkeletonClass.
type DBusObjectSkeletonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for authorize_method
	// Private : padding
}

// IOModuleScope is a wrapper around the C record GIOModuleScope.
type IOModuleScope struct {
	native unsafe.Pointer
}

// TlsDatabaseClass is a wrapper around the C record GTlsDatabaseClass.
type TlsDatabaseClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for verify_chain
	// no type for verify_chain_async
	// no type for verify_chain_finish
	// no type for create_certificate_handle
	// no type for lookup_certificate_for_handle
	// no type for lookup_certificate_for_handle_async
	// no type for lookup_certificate_for_handle_finish
	// no type for lookup_certificate_issuer
	// no type for lookup_certificate_issuer_async
	// no type for lookup_certificate_issuer_finish
	// no type for lookup_certificates_issued_by
	// no type for lookup_certificates_issued_by_async
	// no type for lookup_certificates_issued_by_finish
	// Private : padding
}

// TlsInteractionClass is a wrapper around the C record GTlsInteractionClass.
type TlsInteractionClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for ask_password
	// no type for ask_password_async
	// no type for ask_password_finish
	// no type for request_certificate
	// no type for request_certificate_async
	// no type for request_certificate_finish
	// Private : padding
}
