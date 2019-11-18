// Code generated - DO NOT EDIT.

package gio

import gi "github.com/pekim/gobbi/internal/gi"

type ActionEntry struct {
	native uintptr
	Name   string
	// UNSUPPORTED : C value 'activate' : missing Type

	ParameterType string
	State         string
	// UNSUPPORTED : C value 'change_state' : missing Type

}

type ActionGroupInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'has_action' : missing Type

	// UNSUPPORTED : C value 'list_actions' : missing Type

	// UNSUPPORTED : C value 'get_action_enabled' : missing Type

	// UNSUPPORTED : C value 'get_action_parameter_type' : missing Type

	// UNSUPPORTED : C value 'get_action_state_type' : missing Type

	// UNSUPPORTED : C value 'get_action_state_hint' : missing Type

	// UNSUPPORTED : C value 'get_action_state' : missing Type

	// UNSUPPORTED : C value 'change_action_state' : missing Type

	// UNSUPPORTED : C value 'activate_action' : missing Type

	// UNSUPPORTED : C value 'action_added' : missing Type

	// UNSUPPORTED : C value 'action_removed' : missing Type

	// UNSUPPORTED : C value 'action_enabled_changed' : missing Type

	// UNSUPPORTED : C value 'action_state_changed' : missing Type

	// UNSUPPORTED : C value 'query_action' : missing Type

}

type ActionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_parameter_type' : missing Type

	// UNSUPPORTED : C value 'get_state_type' : missing Type

	// UNSUPPORTED : C value 'get_state_hint' : missing Type

	// UNSUPPORTED : C value 'get_enabled' : missing Type

	// UNSUPPORTED : C value 'get_state' : missing Type

	// UNSUPPORTED : C value 'change_state' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

}

type ActionMapInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'lookup_action' : missing Type

	// UNSUPPORTED : C value 'add_action' : missing Type

	// UNSUPPORTED : C value 'remove_action' : missing Type

}

type AppInfoIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'dup' : missing Type

	// UNSUPPORTED : C value 'equal' : missing Type

	// UNSUPPORTED : C value 'get_id' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_description' : missing Type

	// UNSUPPORTED : C value 'get_executable' : missing Type

	// UNSUPPORTED : C value 'get_icon' : missing Type

	// UNSUPPORTED : C value 'launch' : missing Type

	// UNSUPPORTED : C value 'supports_uris' : missing Type

	// UNSUPPORTED : C value 'supports_files' : missing Type

	// UNSUPPORTED : C value 'launch_uris' : missing Type

	// UNSUPPORTED : C value 'should_show' : missing Type

	// UNSUPPORTED : C value 'set_as_default_for_type' : missing Type

	// UNSUPPORTED : C value 'set_as_default_for_extension' : missing Type

	// UNSUPPORTED : C value 'add_supports_type' : missing Type

	// UNSUPPORTED : C value 'can_remove_supports_type' : missing Type

	// UNSUPPORTED : C value 'remove_supports_type' : missing Type

	// UNSUPPORTED : C value 'can_delete' : missing Type

	// UNSUPPORTED : C value 'do_delete' : missing Type

	// UNSUPPORTED : C value 'get_commandline' : missing Type

	// UNSUPPORTED : C value 'get_display_name' : missing Type

	// UNSUPPORTED : C value 'set_as_last_used_for_type' : missing Type

	// UNSUPPORTED : C value 'get_supported_types' : missing Type

	// UNSUPPORTED : C value 'launch_uris_async' : missing Type

	// UNSUPPORTED : C value 'launch_uris_finish' : missing Type

}

type AppLaunchContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_display' : missing Type

	// UNSUPPORTED : C value 'get_startup_notify_id' : missing Type

	// UNSUPPORTED : C value 'launch_failed' : missing Type

	// UNSUPPORTED : C value 'launched' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

}

type AppLaunchContextPrivate struct {
	native uintptr
}

type ApplicationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'startup' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'open' : missing Type

	// UNSUPPORTED : C value 'command_line' : missing Type

	// UNSUPPORTED : C value 'local_command_line' : missing Type

	// UNSUPPORTED : C value 'before_emit' : missing Type

	// UNSUPPORTED : C value 'after_emit' : missing Type

	// UNSUPPORTED : C value 'add_platform_data' : missing Type

	// UNSUPPORTED : C value 'quit_mainloop' : missing Type

	// UNSUPPORTED : C value 'run_mainloop' : missing Type

	// UNSUPPORTED : C value 'shutdown' : missing Type

	// UNSUPPORTED : C value 'dbus_register' : missing Type

	// UNSUPPORTED : C value 'dbus_unregister' : missing Type

	// UNSUPPORTED : C value 'handle_local_options' : missing Type

	// UNSUPPORTED : C value 'name_lost' : missing Type

}

type ApplicationCommandLineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'print_literal' : missing Type

	// UNSUPPORTED : C value 'printerr_literal' : missing Type

	// UNSUPPORTED : C value 'get_stdin' : missing Type

}

type ApplicationCommandLinePrivate struct {
	native uintptr
}

type ApplicationPrivate struct {
	native uintptr
}

type AsyncInitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'init_async' : missing Type

	// UNSUPPORTED : C value 'init_finish' : missing Type

}

type AsyncResultIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_user_data' : missing Type

	// UNSUPPORTED : C value 'get_source_object' : missing Type

	// UNSUPPORTED : C value 'is_tagged' : missing Type

}

type BufferedInputStreamClass struct {
	native      uintptr
	ParentClass *FilterInputStreamClass
	// UNSUPPORTED : C value 'fill' : missing Type

	// UNSUPPORTED : C value 'fill_async' : missing Type

	// UNSUPPORTED : C value 'fill_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type BufferedInputStreamPrivate struct {
	native uintptr
}

type BufferedOutputStreamClass struct {
	native      uintptr
	ParentClass *FilterOutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

}

type BufferedOutputStreamPrivate struct {
	native uintptr
}

type CancellableClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'cancelled' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type CancellablePrivate struct {
	native uintptr
}

type CharsetConverterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type ConverterIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'convert' : missing Type

	// UNSUPPORTED : C value 'reset' : missing Type

}

type ConverterInputStreamClass struct {
	native      uintptr
	ParentClass *FilterInputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type ConverterInputStreamPrivate struct {
	native uintptr
}

type ConverterOutputStreamClass struct {
	native      uintptr
	ParentClass *FilterOutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type ConverterOutputStreamPrivate struct {
	native uintptr
}

type CredentialsClass struct {
	native uintptr
}

type DBusAnnotationInfo struct {
	native   uintptr
	RefCount int32
	Key      string
	Value    string
	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusArgInfo struct {
	native    uintptr
	RefCount  int32
	Name      string
	Signature string
	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusErrorEntry struct {
	native        uintptr
	ErrorCode     int32
	DbusErrorName string
}

type DBusInterfaceIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_info' : missing Type

	// UNSUPPORTED : C value 'get_object' : missing Type

	// UNSUPPORTED : C value 'set_object' : missing Type

	// UNSUPPORTED : C value 'dup_object' : missing Type

}

type DBusInterfaceInfo struct {
	native   uintptr
	RefCount int32
	Name     string
	// UNSUPPORTED : C value 'methods' : missing Type

	// UNSUPPORTED : C value 'signals' : missing Type

	// UNSUPPORTED : C value 'properties' : missing Type

	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusInterfaceSkeletonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_info' : missing Type

	// UNSUPPORTED : C value 'get_vtable' : missing Type

	// UNSUPPORTED : C value 'get_properties' : missing Type

	// UNSUPPORTED : C value 'flush' : missing Type

	// UNSUPPORTED : C value 'g_authorize_method' : missing Type

}

type DBusInterfaceSkeletonPrivate struct {
	native uintptr
}

type DBusInterfaceVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'method_call' : no Go type for 'DBusInterfaceMethodCallFunc'

	// UNSUPPORTED : C value 'get_property' : no Go type for 'DBusInterfaceGetPropertyFunc'

	// UNSUPPORTED : C value 'set_property' : no Go type for 'DBusInterfaceSetPropertyFunc'

}

type DBusMethodInfo struct {
	native   uintptr
	RefCount int32
	Name     string
	// UNSUPPORTED : C value 'in_args' : missing Type

	// UNSUPPORTED : C value 'out_args' : missing Type

	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusNodeInfo struct {
	native   uintptr
	RefCount int32
	Path     string
	// UNSUPPORTED : C value 'interfaces' : missing Type

	// UNSUPPORTED : C value 'nodes' : missing Type

	// UNSUPPORTED : C value 'annotations' : missing Type

}

var newForXmlDBusNodeInfoInvoker *gi.Function

// DBusNodeInfoNewForXml is a representation of the C type g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) *DBusNodeInfo {
	if newForXmlDBusNodeInfoInvoker == nil {
		newForXmlDBusNodeInfoInvoker = gi.FunctionInvokerNew("Gio", "new_for_xml")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(xmlData)

	ret := newForXmlDBusNodeInfoInvoker.Invoke(inArgs[:], nil)

	return &DBusNodeInfo{native: ret.Pointer()}
}

type DBusObjectIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_object_path' : missing Type

	// UNSUPPORTED : C value 'get_interfaces' : missing Type

	// UNSUPPORTED : C value 'get_interface' : missing Type

	// UNSUPPORTED : C value 'interface_added' : missing Type

	// UNSUPPORTED : C value 'interface_removed' : missing Type

}

type DBusObjectManagerClientClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'interface_proxy_signal' : missing Type

	// UNSUPPORTED : C value 'interface_proxy_properties_changed' : missing Type

}

type DBusObjectManagerClientPrivate struct {
	native uintptr
}

type DBusObjectManagerIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_object_path' : missing Type

	// UNSUPPORTED : C value 'get_objects' : missing Type

	// UNSUPPORTED : C value 'get_object' : missing Type

	// UNSUPPORTED : C value 'get_interface' : missing Type

	// UNSUPPORTED : C value 'object_added' : missing Type

	// UNSUPPORTED : C value 'object_removed' : missing Type

	// UNSUPPORTED : C value 'interface_added' : missing Type

	// UNSUPPORTED : C value 'interface_removed' : missing Type

}

type DBusObjectManagerServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type DBusObjectManagerServerPrivate struct {
	native uintptr
}

type DBusObjectProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type DBusObjectProxyPrivate struct {
	native uintptr
}

type DBusObjectSkeletonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'authorize_method' : missing Type

}

type DBusObjectSkeletonPrivate struct {
	native uintptr
}

type DBusPropertyInfo struct {
	native    uintptr
	RefCount  int32
	Name      string
	Signature string
	// UNSUPPORTED : C value 'flags' : no Go type for 'DBusPropertyInfoFlags'

	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'g_properties_changed' : missing Type

	// UNSUPPORTED : C value 'g_signal' : missing Type

}

type DBusProxyPrivate struct {
	native uintptr
}

type DBusSignalInfo struct {
	native   uintptr
	RefCount int32
	Name     string
	// UNSUPPORTED : C value 'args' : missing Type

	// UNSUPPORTED : C value 'annotations' : missing Type

}

type DBusSubtreeVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'enumerate' : no Go type for 'DBusSubtreeEnumerateFunc'

	// UNSUPPORTED : C value 'introspect' : no Go type for 'DBusSubtreeIntrospectFunc'

	// UNSUPPORTED : C value 'dispatch' : no Go type for 'DBusSubtreeDispatchFunc'

}

type DataInputStreamClass struct {
	native      uintptr
	ParentClass *BufferedInputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type DataInputStreamPrivate struct {
	native uintptr
}

type DataOutputStreamClass struct {
	native      uintptr
	ParentClass *FilterOutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type DataOutputStreamPrivate struct {
	native uintptr
}

type DatagramBasedInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'receive_messages' : missing Type

	// UNSUPPORTED : C value 'send_messages' : missing Type

	// UNSUPPORTED : C value 'create_source' : missing Type

	// UNSUPPORTED : C value 'condition_check' : missing Type

	// UNSUPPORTED : C value 'condition_wait' : missing Type

}

type DesktopAppInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type DesktopAppInfoLookupIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_default_for_uri_scheme' : missing Type

}

type DriveIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'disconnected' : missing Type

	// UNSUPPORTED : C value 'eject_button' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_icon' : missing Type

	// UNSUPPORTED : C value 'has_volumes' : missing Type

	// UNSUPPORTED : C value 'get_volumes' : missing Type

	// UNSUPPORTED : C value 'is_media_removable' : missing Type

	// UNSUPPORTED : C value 'has_media' : missing Type

	// UNSUPPORTED : C value 'is_media_check_automatic' : missing Type

	// UNSUPPORTED : C value 'can_eject' : missing Type

	// UNSUPPORTED : C value 'can_poll_for_media' : missing Type

	// UNSUPPORTED : C value 'eject' : missing Type

	// UNSUPPORTED : C value 'eject_finish' : missing Type

	// UNSUPPORTED : C value 'poll_for_media' : missing Type

	// UNSUPPORTED : C value 'poll_for_media_finish' : missing Type

	// UNSUPPORTED : C value 'get_identifier' : missing Type

	// UNSUPPORTED : C value 'enumerate_identifiers' : missing Type

	// UNSUPPORTED : C value 'get_start_stop_type' : missing Type

	// UNSUPPORTED : C value 'can_start' : missing Type

	// UNSUPPORTED : C value 'can_start_degraded' : missing Type

	// UNSUPPORTED : C value 'start' : missing Type

	// UNSUPPORTED : C value 'start_finish' : missing Type

	// UNSUPPORTED : C value 'can_stop' : missing Type

	// UNSUPPORTED : C value 'stop' : missing Type

	// UNSUPPORTED : C value 'stop_finish' : missing Type

	// UNSUPPORTED : C value 'stop_button' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'get_sort_key' : missing Type

	// UNSUPPORTED : C value 'get_symbolic_icon' : missing Type

	// UNSUPPORTED : C value 'is_removable' : missing Type

}

type DtlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

}

type DtlsConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'accept_certificate' : missing Type

	// UNSUPPORTED : C value 'handshake' : missing Type

	// UNSUPPORTED : C value 'handshake_async' : missing Type

	// UNSUPPORTED : C value 'handshake_finish' : missing Type

	// UNSUPPORTED : C value 'shutdown' : missing Type

	// UNSUPPORTED : C value 'shutdown_async' : missing Type

	// UNSUPPORTED : C value 'shutdown_finish' : missing Type

	// UNSUPPORTED : C value 'set_advertised_protocols' : missing Type

	// UNSUPPORTED : C value 'get_negotiated_protocol' : missing Type

}

type DtlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

}

type EmblemClass struct {
	native uintptr
}

type EmblemedIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type EmblemedIconPrivate struct {
	native uintptr
}

type FileAttributeInfo struct {
	native uintptr
	Name   string
	// UNSUPPORTED : C value 'type' : no Go type for 'FileAttributeType'

	// UNSUPPORTED : C value 'flags' : no Go type for 'FileAttributeInfoFlags'

}

type FileAttributeInfoList struct {
	native uintptr
	Infos  *FileAttributeInfo
	NInfos int32
}

var newFileAttributeInfoListInvoker *gi.Function

// FileAttributeInfoListNew is a representation of the C type g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {
	if newFileAttributeInfoListInvoker == nil {
		newFileAttributeInfoListInvoker = gi.FunctionInvokerNew("Gio", "new")
	}

	ret := newFileAttributeInfoListInvoker.Invoke(nil, nil)

	return &FileAttributeInfoList{native: ret.Pointer()}
}

type FileAttributeMatcher struct {
	native uintptr
}

var newFileAttributeMatcherInvoker *gi.Function

// FileAttributeMatcherNew is a representation of the C type g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	if newFileAttributeMatcherInvoker == nil {
		newFileAttributeMatcherInvoker = gi.FunctionInvokerNew("Gio", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(attributes)

	ret := newFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

	return &FileAttributeMatcher{native: ret.Pointer()}
}

type FileDescriptorBasedIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_fd' : missing Type

}

type FileEnumeratorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'next_file' : missing Type

	// UNSUPPORTED : C value 'close_fn' : missing Type

	// UNSUPPORTED : C value 'next_files_async' : missing Type

	// UNSUPPORTED : C value 'next_files_finish' : missing Type

	// UNSUPPORTED : C value 'close_async' : missing Type

	// UNSUPPORTED : C value 'close_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

}

type FileEnumeratorPrivate struct {
	native uintptr
}

type FileIOStreamClass struct {
	native      uintptr
	ParentClass *IOStreamClass
	// UNSUPPORTED : C value 'tell' : missing Type

	// UNSUPPORTED : C value 'can_seek' : missing Type

	// UNSUPPORTED : C value 'seek' : missing Type

	// UNSUPPORTED : C value 'can_truncate' : missing Type

	// UNSUPPORTED : C value 'truncate_fn' : missing Type

	// UNSUPPORTED : C value 'query_info' : missing Type

	// UNSUPPORTED : C value 'query_info_async' : missing Type

	// UNSUPPORTED : C value 'query_info_finish' : missing Type

	// UNSUPPORTED : C value 'get_etag' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type FileIOStreamPrivate struct {
	native uintptr
}

type FileIconClass struct {
	native uintptr
}

type FileIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'dup' : missing Type

	// UNSUPPORTED : C value 'hash' : missing Type

	// UNSUPPORTED : C value 'equal' : missing Type

	// UNSUPPORTED : C value 'is_native' : missing Type

	// UNSUPPORTED : C value 'has_uri_scheme' : missing Type

	// UNSUPPORTED : C value 'get_uri_scheme' : missing Type

	// UNSUPPORTED : C value 'get_basename' : missing Type

	// UNSUPPORTED : C value 'get_path' : missing Type

	// UNSUPPORTED : C value 'get_uri' : missing Type

	// UNSUPPORTED : C value 'get_parse_name' : missing Type

	// UNSUPPORTED : C value 'get_parent' : missing Type

	// UNSUPPORTED : C value 'prefix_matches' : missing Type

	// UNSUPPORTED : C value 'get_relative_path' : missing Type

	// UNSUPPORTED : C value 'resolve_relative_path' : missing Type

	// UNSUPPORTED : C value 'get_child_for_display_name' : missing Type

	// UNSUPPORTED : C value 'enumerate_children' : missing Type

	// UNSUPPORTED : C value 'enumerate_children_async' : missing Type

	// UNSUPPORTED : C value 'enumerate_children_finish' : missing Type

	// UNSUPPORTED : C value 'query_info' : missing Type

	// UNSUPPORTED : C value 'query_info_async' : missing Type

	// UNSUPPORTED : C value 'query_info_finish' : missing Type

	// UNSUPPORTED : C value 'query_filesystem_info' : missing Type

	// UNSUPPORTED : C value 'query_filesystem_info_async' : missing Type

	// UNSUPPORTED : C value 'query_filesystem_info_finish' : missing Type

	// UNSUPPORTED : C value 'find_enclosing_mount' : missing Type

	// UNSUPPORTED : C value 'find_enclosing_mount_async' : missing Type

	// UNSUPPORTED : C value 'find_enclosing_mount_finish' : missing Type

	// UNSUPPORTED : C value 'set_display_name' : missing Type

	// UNSUPPORTED : C value 'set_display_name_async' : missing Type

	// UNSUPPORTED : C value 'set_display_name_finish' : missing Type

	// UNSUPPORTED : C value 'query_settable_attributes' : missing Type

	// UNSUPPORTED : C value '_query_settable_attributes_async' : missing Type

	// UNSUPPORTED : C value '_query_settable_attributes_finish' : missing Type

	// UNSUPPORTED : C value 'query_writable_namespaces' : missing Type

	// UNSUPPORTED : C value '_query_writable_namespaces_async' : missing Type

	// UNSUPPORTED : C value '_query_writable_namespaces_finish' : missing Type

	// UNSUPPORTED : C value 'set_attribute' : missing Type

	// UNSUPPORTED : C value 'set_attributes_from_info' : missing Type

	// UNSUPPORTED : C value 'set_attributes_async' : missing Type

	// UNSUPPORTED : C value 'set_attributes_finish' : missing Type

	// UNSUPPORTED : C value 'read_fn' : missing Type

	// UNSUPPORTED : C value 'read_async' : missing Type

	// UNSUPPORTED : C value 'read_finish' : missing Type

	// UNSUPPORTED : C value 'append_to' : missing Type

	// UNSUPPORTED : C value 'append_to_async' : missing Type

	// UNSUPPORTED : C value 'append_to_finish' : missing Type

	// UNSUPPORTED : C value 'create' : missing Type

	// UNSUPPORTED : C value 'create_async' : missing Type

	// UNSUPPORTED : C value 'create_finish' : missing Type

	// UNSUPPORTED : C value 'replace' : missing Type

	// UNSUPPORTED : C value 'replace_async' : missing Type

	// UNSUPPORTED : C value 'replace_finish' : missing Type

	// UNSUPPORTED : C value 'delete_file' : missing Type

	// UNSUPPORTED : C value 'delete_file_async' : missing Type

	// UNSUPPORTED : C value 'delete_file_finish' : missing Type

	// UNSUPPORTED : C value 'trash' : missing Type

	// UNSUPPORTED : C value 'trash_async' : missing Type

	// UNSUPPORTED : C value 'trash_finish' : missing Type

	// UNSUPPORTED : C value 'make_directory' : missing Type

	// UNSUPPORTED : C value 'make_directory_async' : missing Type

	// UNSUPPORTED : C value 'make_directory_finish' : missing Type

	// UNSUPPORTED : C value 'make_symbolic_link' : missing Type

	// UNSUPPORTED : C value '_make_symbolic_link_async' : missing Type

	// UNSUPPORTED : C value '_make_symbolic_link_finish' : missing Type

	// UNSUPPORTED : C value 'copy' : missing Type

	// UNSUPPORTED : C value 'copy_async' : missing Type

	// UNSUPPORTED : C value 'copy_finish' : missing Type

	// UNSUPPORTED : C value 'move' : missing Type

	// UNSUPPORTED : C value '_move_async' : missing Type

	// UNSUPPORTED : C value '_move_finish' : missing Type

	// UNSUPPORTED : C value 'mount_mountable' : missing Type

	// UNSUPPORTED : C value 'mount_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'unmount_mountable' : missing Type

	// UNSUPPORTED : C value 'unmount_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'eject_mountable' : missing Type

	// UNSUPPORTED : C value 'eject_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'mount_enclosing_volume' : missing Type

	// UNSUPPORTED : C value 'mount_enclosing_volume_finish' : missing Type

	// UNSUPPORTED : C value 'monitor_dir' : missing Type

	// UNSUPPORTED : C value 'monitor_file' : missing Type

	// UNSUPPORTED : C value 'open_readwrite' : missing Type

	// UNSUPPORTED : C value 'open_readwrite_async' : missing Type

	// UNSUPPORTED : C value 'open_readwrite_finish' : missing Type

	// UNSUPPORTED : C value 'create_readwrite' : missing Type

	// UNSUPPORTED : C value 'create_readwrite_async' : missing Type

	// UNSUPPORTED : C value 'create_readwrite_finish' : missing Type

	// UNSUPPORTED : C value 'replace_readwrite' : missing Type

	// UNSUPPORTED : C value 'replace_readwrite_async' : missing Type

	// UNSUPPORTED : C value 'replace_readwrite_finish' : missing Type

	// UNSUPPORTED : C value 'start_mountable' : missing Type

	// UNSUPPORTED : C value 'start_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'stop_mountable' : missing Type

	// UNSUPPORTED : C value 'stop_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'supports_thread_contexts' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'unmount_mountable_with_operation' : missing Type

	// UNSUPPORTED : C value 'unmount_mountable_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'eject_mountable_with_operation' : missing Type

	// UNSUPPORTED : C value 'eject_mountable_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'poll_mountable' : missing Type

	// UNSUPPORTED : C value 'poll_mountable_finish' : missing Type

	// UNSUPPORTED : C value 'measure_disk_usage' : missing Type

	// UNSUPPORTED : C value 'measure_disk_usage_async' : missing Type

	// UNSUPPORTED : C value 'measure_disk_usage_finish' : missing Type

}

type FileInfoClass struct {
	native uintptr
}

type FileInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value 'tell' : missing Type

	// UNSUPPORTED : C value 'can_seek' : missing Type

	// UNSUPPORTED : C value 'seek' : missing Type

	// UNSUPPORTED : C value 'query_info' : missing Type

	// UNSUPPORTED : C value 'query_info_async' : missing Type

	// UNSUPPORTED : C value 'query_info_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type FileInputStreamPrivate struct {
	native uintptr
}

type FileMonitorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'cancel' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type FileMonitorPrivate struct {
	native uintptr
}

type FileOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value 'tell' : missing Type

	// UNSUPPORTED : C value 'can_seek' : missing Type

	// UNSUPPORTED : C value 'seek' : missing Type

	// UNSUPPORTED : C value 'can_truncate' : missing Type

	// UNSUPPORTED : C value 'truncate_fn' : missing Type

	// UNSUPPORTED : C value 'query_info' : missing Type

	// UNSUPPORTED : C value 'query_info_async' : missing Type

	// UNSUPPORTED : C value 'query_info_finish' : missing Type

	// UNSUPPORTED : C value 'get_etag' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type FileOutputStreamPrivate struct {
	native uintptr
}

type FilenameCompleterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'got_completion_data' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

}

type FilterInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

}

type FilterOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

}

type IOExtension struct {
	native uintptr
}

type IOExtensionPoint struct {
	native uintptr
}

type IOModuleClass struct {
	native uintptr
}

type IOModuleScope struct {
	native uintptr
}

type IOSchedulerJob struct {
	native uintptr
}

type IOStreamAdapter struct {
	native uintptr
}

type IOStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_input_stream' : missing Type

	// UNSUPPORTED : C value 'get_output_stream' : missing Type

	// UNSUPPORTED : C value 'close_fn' : missing Type

	// UNSUPPORTED : C value 'close_async' : missing Type

	// UNSUPPORTED : C value 'close_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

	// UNSUPPORTED : C value '_g_reserved8' : missing Type

	// UNSUPPORTED : C value '_g_reserved9' : missing Type

	// UNSUPPORTED : C value '_g_reserved10' : missing Type

}

type IOStreamPrivate struct {
	native uintptr
}

type IconIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'hash' : missing Type

	// UNSUPPORTED : C value 'equal' : missing Type

	// UNSUPPORTED : C value 'to_tokens' : missing Type

	// UNSUPPORTED : C value 'from_tokens' : missing Type

	// UNSUPPORTED : C value 'serialize' : missing Type

}

type InetAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'to_string' : missing Type

	// UNSUPPORTED : C value 'to_bytes' : missing Type

}

type InetAddressMaskClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type InetAddressMaskPrivate struct {
	native uintptr
}

type InetAddressPrivate struct {
	native uintptr
}

type InetSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

type InetSocketAddressPrivate struct {
	native uintptr
}

type InitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'init' : missing Type

}

type InputMessage struct {
	native uintptr
	// UNSUPPORTED : C value 'address' : no Go type for 'SocketAddress'

	// UNSUPPORTED : C value 'vectors' : missing Type

	NumVectors    uint32
	BytesReceived uintptr
	Flags         int32
	// UNSUPPORTED : C value 'control_messages' : missing Type

	NumControlMessages uint32
}

type InputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'read_fn' : missing Type

	// UNSUPPORTED : C value 'skip' : missing Type

	// UNSUPPORTED : C value 'close_fn' : missing Type

	// UNSUPPORTED : C value 'read_async' : missing Type

	// UNSUPPORTED : C value 'read_finish' : missing Type

	// UNSUPPORTED : C value 'skip_async' : missing Type

	// UNSUPPORTED : C value 'skip_finish' : missing Type

	// UNSUPPORTED : C value 'close_async' : missing Type

	// UNSUPPORTED : C value 'close_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type InputStreamPrivate struct {
	native uintptr
}

type InputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'

	Size uintptr
}

type ListModelInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_item_type' : missing Type

	// UNSUPPORTED : C value 'get_n_items' : missing Type

	// UNSUPPORTED : C value 'get_item' : missing Type

}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type LoadableIconIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'load' : missing Type

	// UNSUPPORTED : C value 'load_async' : missing Type

	// UNSUPPORTED : C value 'load_finish' : missing Type

}

type MemoryInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type MemoryInputStreamPrivate struct {
	native uintptr
}

type MemoryOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type MemoryOutputStreamPrivate struct {
	native uintptr
}

type MenuAttributeIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_next' : missing Type

}

type MenuAttributeIterPrivate struct {
	native uintptr
}

type MenuLinkIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_next' : missing Type

}

type MenuLinkIterPrivate struct {
	native uintptr
}

type MenuModelClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'is_mutable' : missing Type

	// UNSUPPORTED : C value 'get_n_items' : missing Type

	// UNSUPPORTED : C value 'get_item_attributes' : missing Type

	// UNSUPPORTED : C value 'iterate_item_attributes' : missing Type

	// UNSUPPORTED : C value 'get_item_attribute_value' : missing Type

	// UNSUPPORTED : C value 'get_item_links' : missing Type

	// UNSUPPORTED : C value 'iterate_item_links' : missing Type

	// UNSUPPORTED : C value 'get_item_link' : missing Type

}

type MenuModelPrivate struct {
	native uintptr
}

type MountIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'unmounted' : missing Type

	// UNSUPPORTED : C value 'get_root' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_icon' : missing Type

	// UNSUPPORTED : C value 'get_uuid' : missing Type

	// UNSUPPORTED : C value 'get_volume' : missing Type

	// UNSUPPORTED : C value 'get_drive' : missing Type

	// UNSUPPORTED : C value 'can_unmount' : missing Type

	// UNSUPPORTED : C value 'can_eject' : missing Type

	// UNSUPPORTED : C value 'unmount' : missing Type

	// UNSUPPORTED : C value 'unmount_finish' : missing Type

	// UNSUPPORTED : C value 'eject' : missing Type

	// UNSUPPORTED : C value 'eject_finish' : missing Type

	// UNSUPPORTED : C value 'remount' : missing Type

	// UNSUPPORTED : C value 'remount_finish' : missing Type

	// UNSUPPORTED : C value 'guess_content_type' : missing Type

	// UNSUPPORTED : C value 'guess_content_type_finish' : missing Type

	// UNSUPPORTED : C value 'guess_content_type_sync' : missing Type

	// UNSUPPORTED : C value 'pre_unmount' : missing Type

	// UNSUPPORTED : C value 'unmount_with_operation' : missing Type

	// UNSUPPORTED : C value 'unmount_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'get_default_location' : missing Type

	// UNSUPPORTED : C value 'get_sort_key' : missing Type

	// UNSUPPORTED : C value 'get_symbolic_icon' : missing Type

}

type MountOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'ask_password' : missing Type

	// UNSUPPORTED : C value 'ask_question' : missing Type

	// UNSUPPORTED : C value 'reply' : missing Type

	// UNSUPPORTED : C value 'aborted' : missing Type

	// UNSUPPORTED : C value 'show_processes' : missing Type

	// UNSUPPORTED : C value 'show_unmount_progress' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

	// UNSUPPORTED : C value '_g_reserved8' : missing Type

	// UNSUPPORTED : C value '_g_reserved9' : missing Type

}

type MountOperationPrivate struct {
	native uintptr
}

type NativeSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

type NativeSocketAddressPrivate struct {
	native uintptr
}

type NativeVolumeMonitorClass struct {
	native      uintptr
	ParentClass *VolumeMonitorClass
	// UNSUPPORTED : C value 'get_mount_for_mount_path' : missing Type

}

type NetworkAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type NetworkAddressPrivate struct {
	native uintptr
}

type NetworkMonitorInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'network_changed' : missing Type

	// UNSUPPORTED : C value 'can_reach' : missing Type

	// UNSUPPORTED : C value 'can_reach_async' : missing Type

	// UNSUPPORTED : C value 'can_reach_finish' : missing Type

}

type NetworkServiceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type NetworkServicePrivate struct {
	native uintptr
}

type OutputMessage struct {
	native uintptr
	// UNSUPPORTED : C value 'address' : no Go type for 'SocketAddress'

	Vectors    *OutputVector
	NumVectors uint32
	BytesSent  uint32
	// UNSUPPORTED : C value 'control_messages' : missing Type

	NumControlMessages uint32
}

type OutputStreamClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'write_fn' : missing Type

	// UNSUPPORTED : C value 'splice' : missing Type

	// UNSUPPORTED : C value 'flush' : missing Type

	// UNSUPPORTED : C value 'close_fn' : missing Type

	// UNSUPPORTED : C value 'write_async' : missing Type

	// UNSUPPORTED : C value 'write_finish' : missing Type

	// UNSUPPORTED : C value 'splice_async' : missing Type

	// UNSUPPORTED : C value 'splice_finish' : missing Type

	// UNSUPPORTED : C value 'flush_async' : missing Type

	// UNSUPPORTED : C value 'flush_finish' : missing Type

	// UNSUPPORTED : C value 'close_async' : missing Type

	// UNSUPPORTED : C value 'close_finish' : missing Type

	// UNSUPPORTED : C value 'writev_fn' : missing Type

	// UNSUPPORTED : C value 'writev_async' : missing Type

	// UNSUPPORTED : C value 'writev_finish' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

	// UNSUPPORTED : C value '_g_reserved8' : missing Type

}

type OutputStreamPrivate struct {
	native uintptr
}

type OutputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'

	Size uintptr
}

type PermissionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'acquire' : missing Type

	// UNSUPPORTED : C value 'acquire_async' : missing Type

	// UNSUPPORTED : C value 'acquire_finish' : missing Type

	// UNSUPPORTED : C value 'release' : missing Type

	// UNSUPPORTED : C value 'release_async' : missing Type

	// UNSUPPORTED : C value 'release_finish' : missing Type

	// UNSUPPORTED : C value 'reserved' : missing Type

}

type PermissionPrivate struct {
	native uintptr
}

type PollableInputStreamInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'can_poll' : missing Type

	// UNSUPPORTED : C value 'is_readable' : missing Type

	// UNSUPPORTED : C value 'create_source' : missing Type

	// UNSUPPORTED : C value 'read_nonblocking' : missing Type

}

type PollableOutputStreamInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'can_poll' : missing Type

	// UNSUPPORTED : C value 'is_writable' : missing Type

	// UNSUPPORTED : C value 'create_source' : missing Type

	// UNSUPPORTED : C value 'write_nonblocking' : missing Type

	// UNSUPPORTED : C value 'writev_nonblocking' : missing Type

}

type ProxyAddressClass struct {
	native      uintptr
	ParentClass *InetSocketAddressClass
}

type ProxyAddressEnumeratorClass struct {
	native uintptr
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

}

type ProxyAddressEnumeratorPrivate struct {
	native uintptr
}

type ProxyAddressPrivate struct {
	native uintptr
}

type ProxyInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'connect' : missing Type

	// UNSUPPORTED : C value 'connect_async' : missing Type

	// UNSUPPORTED : C value 'connect_finish' : missing Type

	// UNSUPPORTED : C value 'supports_hostname' : missing Type

}

type ProxyResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'is_supported' : missing Type

	// UNSUPPORTED : C value 'lookup' : missing Type

	// UNSUPPORTED : C value 'lookup_async' : missing Type

	// UNSUPPORTED : C value 'lookup_finish' : missing Type

}

type RemoteActionGroupInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'activate_action_full' : missing Type

	// UNSUPPORTED : C value 'change_action_state_full' : missing Type

}

type ResolverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'reload' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name_async' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_by_address' : missing Type

	// UNSUPPORTED : C value 'lookup_by_address_async' : missing Type

	// UNSUPPORTED : C value 'lookup_by_address_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_service' : missing Type

	// UNSUPPORTED : C value 'lookup_service_async' : missing Type

	// UNSUPPORTED : C value 'lookup_service_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_records' : missing Type

	// UNSUPPORTED : C value 'lookup_records_async' : missing Type

	// UNSUPPORTED : C value 'lookup_records_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name_with_flags_async' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name_with_flags_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_by_name_with_flags' : missing Type

}

type ResolverPrivate struct {
	native uintptr
}

type Resource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_resource_new_from_data' : parameter 'data' of type 'GLib.Bytes' not supported

type SeekableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'tell' : missing Type

	// UNSUPPORTED : C value 'can_seek' : missing Type

	// UNSUPPORTED : C value 'seek' : missing Type

	// UNSUPPORTED : C value 'can_truncate' : missing Type

	// UNSUPPORTED : C value 'truncate_fn' : missing Type

}

type SettingsBackendClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'read' : missing Type

	// UNSUPPORTED : C value 'get_writable' : missing Type

	// UNSUPPORTED : C value 'write' : missing Type

	// UNSUPPORTED : C value 'write_tree' : missing Type

	// UNSUPPORTED : C value 'reset' : missing Type

	// UNSUPPORTED : C value 'subscribe' : missing Type

	// UNSUPPORTED : C value 'unsubscribe' : missing Type

	// UNSUPPORTED : C value 'sync' : missing Type

	// UNSUPPORTED : C value 'get_permission' : missing Type

	// UNSUPPORTED : C value 'read_user_value' : missing Type

}

type SettingsBackendPrivate struct {
	native uintptr
}

type SettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'writable_changed' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'writable_change_event' : missing Type

	// UNSUPPORTED : C value 'change_event' : missing Type

	// UNSUPPORTED : C value 'padding' : missing Type

}

type SettingsPrivate struct {
	native uintptr
}

type SettingsSchema struct {
	native uintptr
}

type SettingsSchemaKey struct {
	native uintptr
}

type SettingsSchemaSource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_source_new_from_directory' : parameter 'directory' of type 'filename' not supported

type SimpleActionGroupClass struct {
	native uintptr
}

type SimpleActionGroupPrivate struct {
	native uintptr
}

type SimpleAsyncResultClass struct {
	native uintptr
}

type SimpleProxyResolverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type SimpleProxyResolverPrivate struct {
	native uintptr
}

type SocketAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_family' : missing Type

	// UNSUPPORTED : C value 'get_native_size' : missing Type

	// UNSUPPORTED : C value 'to_native' : missing Type

}

type SocketAddressEnumeratorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'next' : missing Type

	// UNSUPPORTED : C value 'next_async' : missing Type

	// UNSUPPORTED : C value 'next_finish' : missing Type

}

type SocketClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

	// UNSUPPORTED : C value '_g_reserved7' : missing Type

	// UNSUPPORTED : C value '_g_reserved8' : missing Type

	// UNSUPPORTED : C value '_g_reserved9' : missing Type

	// UNSUPPORTED : C value '_g_reserved10' : missing Type

}

type SocketClientClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'event' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

}

type SocketClientPrivate struct {
	native uintptr
}

type SocketConnectableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'enumerate' : missing Type

	// UNSUPPORTED : C value 'proxy_enumerate' : missing Type

	// UNSUPPORTED : C value 'to_string' : missing Type

}

type SocketConnectionClass struct {
	native      uintptr
	ParentClass *IOStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

}

type SocketConnectionPrivate struct {
	native uintptr
}

type SocketControlMessageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_size' : missing Type

	// UNSUPPORTED : C value 'get_level' : missing Type

	// UNSUPPORTED : C value 'get_type' : missing Type

	// UNSUPPORTED : C value 'serialize' : missing Type

	// UNSUPPORTED : C value 'deserialize' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type SocketControlMessagePrivate struct {
	native uintptr
}

type SocketListenerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'event' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

}

type SocketListenerPrivate struct {
	native uintptr
}

type SocketPrivate struct {
	native uintptr
}

type SocketServiceClass struct {
	native      uintptr
	ParentClass *SocketListenerClass
	// UNSUPPORTED : C value 'incoming' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

}

type SocketServicePrivate struct {
	native uintptr
}

type SrvTarget struct {
	native uintptr
}

var newSrvTargetInvoker *gi.Function

// SrvTargetNew is a representation of the C type g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	if newSrvTargetInvoker == nil {
		newSrvTargetInvoker = gi.FunctionInvokerNew("Gio", "new")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)
	inArgs[2].SetUint16(priority)
	inArgs[3].SetUint16(weight)

	ret := newSrvTargetInvoker.Invoke(inArgs[:], nil)

	return &SrvTarget{native: ret.Pointer()}
}

type StaticResource struct {
	native uintptr
}

type TaskClass struct {
	native uintptr
}

type TcpConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

type TcpConnectionPrivate struct {
	native uintptr
}

type TcpWrapperConnectionClass struct {
	native      uintptr
	ParentClass *TcpConnectionClass
}

type TcpWrapperConnectionPrivate struct {
	native uintptr
}

type ThemedIconClass struct {
	native uintptr
}

type ThreadedSocketServiceClass struct {
	native      uintptr
	ParentClass *SocketServiceClass
	// UNSUPPORTED : C value 'run' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type ThreadedSocketServicePrivate struct {
	native uintptr
}

type TlsBackendInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'supports_tls' : missing Type

	// UNSUPPORTED : C value 'get_certificate_type' : missing Type

	// UNSUPPORTED : C value 'get_client_connection_type' : missing Type

	// UNSUPPORTED : C value 'get_server_connection_type' : missing Type

	// UNSUPPORTED : C value 'get_file_database_type' : missing Type

	// UNSUPPORTED : C value 'get_default_database' : missing Type

	// UNSUPPORTED : C value 'supports_dtls' : missing Type

	// UNSUPPORTED : C value 'get_dtls_client_connection_type' : missing Type

	// UNSUPPORTED : C value 'get_dtls_server_connection_type' : missing Type

}

type TlsCertificateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'verify' : missing Type

}

type TlsCertificatePrivate struct {
	native uintptr
}

type TlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'copy_session_state' : missing Type

}

type TlsConnectionClass struct {
	native      uintptr
	ParentClass *IOStreamClass
	// UNSUPPORTED : C value 'accept_certificate' : missing Type

	// UNSUPPORTED : C value 'handshake' : missing Type

	// UNSUPPORTED : C value 'handshake_async' : missing Type

	// UNSUPPORTED : C value 'handshake_finish' : missing Type

}

type TlsConnectionPrivate struct {
	native uintptr
}

type TlsDatabaseClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'verify_chain' : missing Type

	// UNSUPPORTED : C value 'verify_chain_async' : missing Type

	// UNSUPPORTED : C value 'verify_chain_finish' : missing Type

	// UNSUPPORTED : C value 'create_certificate_handle' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_for_handle' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_for_handle_async' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_for_handle_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_issuer' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_issuer_async' : missing Type

	// UNSUPPORTED : C value 'lookup_certificate_issuer_finish' : missing Type

	// UNSUPPORTED : C value 'lookup_certificates_issued_by' : missing Type

	// UNSUPPORTED : C value 'lookup_certificates_issued_by_async' : missing Type

	// UNSUPPORTED : C value 'lookup_certificates_issued_by_finish' : missing Type

}

type TlsDatabasePrivate struct {
	native uintptr
}

type TlsFileDatabaseInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

}

type TlsInteractionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'ask_password' : missing Type

	// UNSUPPORTED : C value 'ask_password_async' : missing Type

	// UNSUPPORTED : C value 'ask_password_finish' : missing Type

	// UNSUPPORTED : C value 'request_certificate' : missing Type

	// UNSUPPORTED : C value 'request_certificate_async' : missing Type

	// UNSUPPORTED : C value 'request_certificate_finish' : missing Type

}

type TlsInteractionPrivate struct {
	native uintptr
}

type TlsPasswordClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_value' : missing Type

	// UNSUPPORTED : C value 'set_value' : missing Type

	// UNSUPPORTED : C value 'get_default_warning' : missing Type

}

type TlsPasswordPrivate struct {
	native uintptr
}

type TlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

}

type UnixConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

type UnixConnectionPrivate struct {
	native uintptr
}

type UnixCredentialsMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

}

type UnixCredentialsMessagePrivate struct {
	native uintptr
}

type UnixFDListClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type UnixFDListPrivate struct {
	native uintptr
}

type UnixFDMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

}

type UnixFDMessagePrivate struct {
	native uintptr
}

type UnixInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type UnixInputStreamPrivate struct {
	native uintptr
}

type UnixMountEntry struct {
	native uintptr
}

type UnixMountMonitorClass struct {
	native uintptr
}

type UnixMountPoint struct {
	native uintptr
}

type UnixOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

}

type UnixOutputStreamPrivate struct {
	native uintptr
}

type UnixSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

type UnixSocketAddressPrivate struct {
	native uintptr
}

type VfsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'is_active' : missing Type

	// UNSUPPORTED : C value 'get_file_for_path' : missing Type

	// UNSUPPORTED : C value 'get_file_for_uri' : missing Type

	// UNSUPPORTED : C value 'get_supported_uri_schemes' : missing Type

	// UNSUPPORTED : C value 'parse_name' : missing Type

	// UNSUPPORTED : C value 'local_file_add_info' : missing Type

	// UNSUPPORTED : C value 'add_writable_namespaces' : missing Type

	// UNSUPPORTED : C value 'local_file_set_attributes' : missing Type

	// UNSUPPORTED : C value 'local_file_removed' : missing Type

	// UNSUPPORTED : C value 'local_file_moved' : missing Type

	// UNSUPPORTED : C value 'deserialize_icon' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

}

type VolumeIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'removed' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_icon' : missing Type

	// UNSUPPORTED : C value 'get_uuid' : missing Type

	// UNSUPPORTED : C value 'get_drive' : missing Type

	// UNSUPPORTED : C value 'get_mount' : missing Type

	// UNSUPPORTED : C value 'can_mount' : missing Type

	// UNSUPPORTED : C value 'can_eject' : missing Type

	// UNSUPPORTED : C value 'mount_fn' : missing Type

	// UNSUPPORTED : C value 'mount_finish' : missing Type

	// UNSUPPORTED : C value 'eject' : missing Type

	// UNSUPPORTED : C value 'eject_finish' : missing Type

	// UNSUPPORTED : C value 'get_identifier' : missing Type

	// UNSUPPORTED : C value 'enumerate_identifiers' : missing Type

	// UNSUPPORTED : C value 'should_automount' : missing Type

	// UNSUPPORTED : C value 'get_activation_root' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation' : missing Type

	// UNSUPPORTED : C value 'eject_with_operation_finish' : missing Type

	// UNSUPPORTED : C value 'get_sort_key' : missing Type

	// UNSUPPORTED : C value 'get_symbolic_icon' : missing Type

}

type VolumeMonitorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'volume_added' : missing Type

	// UNSUPPORTED : C value 'volume_removed' : missing Type

	// UNSUPPORTED : C value 'volume_changed' : missing Type

	// UNSUPPORTED : C value 'mount_added' : missing Type

	// UNSUPPORTED : C value 'mount_removed' : missing Type

	// UNSUPPORTED : C value 'mount_pre_unmount' : missing Type

	// UNSUPPORTED : C value 'mount_changed' : missing Type

	// UNSUPPORTED : C value 'drive_connected' : missing Type

	// UNSUPPORTED : C value 'drive_disconnected' : missing Type

	// UNSUPPORTED : C value 'drive_changed' : missing Type

	// UNSUPPORTED : C value 'is_supported' : missing Type

	// UNSUPPORTED : C value 'get_connected_drives' : missing Type

	// UNSUPPORTED : C value 'get_volumes' : missing Type

	// UNSUPPORTED : C value 'get_mounts' : missing Type

	// UNSUPPORTED : C value 'get_volume_for_uuid' : missing Type

	// UNSUPPORTED : C value 'get_mount_for_uuid' : missing Type

	// UNSUPPORTED : C value 'adopt_orphan_mount' : missing Type

	// UNSUPPORTED : C value 'drive_eject_button' : missing Type

	// UNSUPPORTED : C value 'drive_stop_button' : missing Type

	// UNSUPPORTED : C value '_g_reserved1' : missing Type

	// UNSUPPORTED : C value '_g_reserved2' : missing Type

	// UNSUPPORTED : C value '_g_reserved3' : missing Type

	// UNSUPPORTED : C value '_g_reserved4' : missing Type

	// UNSUPPORTED : C value '_g_reserved5' : missing Type

	// UNSUPPORTED : C value '_g_reserved6' : missing Type

}

type ZlibCompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type ZlibDecompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}
