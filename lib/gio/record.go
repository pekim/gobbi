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

var refDBusAnnotationInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_annotation_info_ref.
func (recv *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	if refDBusAnnotationInfoInvoker == nil {
		refDBusAnnotationInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusAnnotationInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusAnnotationInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusAnnotationInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusAnnotationInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_annotation_info_unref.
func (recv *DBusAnnotationInfo) Unref() {
	if unrefDBusAnnotationInfoInvoker == nil {
		unrefDBusAnnotationInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusAnnotationInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusAnnotationInfoInvoker.Invoke(inArgs[:], nil)

}

type DBusArgInfo struct {
	native    uintptr
	RefCount  int32
	Name      string
	Signature string
	// UNSUPPORTED : C value 'annotations' : missing Type
}

var refDBusArgInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_arg_info_ref.
func (recv *DBusArgInfo) Ref() *DBusArgInfo {
	if refDBusArgInfoInvoker == nil {
		refDBusArgInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusArgInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusArgInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusArgInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusArgInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_arg_info_unref.
func (recv *DBusArgInfo) Unref() {
	if unrefDBusArgInfoInvoker == nil {
		unrefDBusArgInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusArgInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusArgInfoInvoker.Invoke(inArgs[:], nil)

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

var cacheBuildDBusInterfaceInfoInvoker *gi.Function

// CacheBuild is a representation of the C type g_dbus_interface_info_cache_build.
func (recv *DBusInterfaceInfo) CacheBuild() {
	if cacheBuildDBusInterfaceInfoInvoker == nil {
		cacheBuildDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "cache_build")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cacheBuildDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

}

var cacheReleaseDBusInterfaceInfoInvoker *gi.Function

// CacheRelease is a representation of the C type g_dbus_interface_info_cache_release.
func (recv *DBusInterfaceInfo) CacheRelease() {
	if cacheReleaseDBusInterfaceInfoInvoker == nil {
		cacheReleaseDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "cache_release")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cacheReleaseDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_dbus_interface_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var lookupMethodDBusInterfaceInfoInvoker *gi.Function

// LookupMethod is a representation of the C type g_dbus_interface_info_lookup_method.
func (recv *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	if lookupMethodDBusInterfaceInfoInvoker == nil {
		lookupMethodDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "lookup_method")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := lookupMethodDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo
}

var lookupPropertyDBusInterfaceInfoInvoker *gi.Function

// LookupProperty is a representation of the C type g_dbus_interface_info_lookup_property.
func (recv *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	if lookupPropertyDBusInterfaceInfoInvoker == nil {
		lookupPropertyDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "lookup_property")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := lookupPropertyDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo
}

var lookupSignalDBusInterfaceInfoInvoker *gi.Function

// LookupSignal is a representation of the C type g_dbus_interface_info_lookup_signal.
func (recv *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	if lookupSignalDBusInterfaceInfoInvoker == nil {
		lookupSignalDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "lookup_signal")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := lookupSignalDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo
}

var refDBusInterfaceInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_interface_info_ref.
func (recv *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	if refDBusInterfaceInfoInvoker == nil {
		refDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusInterfaceInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_interface_info_unref.
func (recv *DBusInterfaceInfo) Unref() {
	if unrefDBusInterfaceInfoInvoker == nil {
		unrefDBusInterfaceInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusInterfaceInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusInterfaceInfoInvoker.Invoke(inArgs[:], nil)

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

var refDBusMethodInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_method_info_ref.
func (recv *DBusMethodInfo) Ref() *DBusMethodInfo {
	if refDBusMethodInfoInvoker == nil {
		refDBusMethodInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusMethodInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusMethodInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusMethodInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_method_info_unref.
func (recv *DBusMethodInfo) Unref() {
	if unrefDBusMethodInfoInvoker == nil {
		unrefDBusMethodInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusMethodInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusMethodInfoInvoker.Invoke(inArgs[:], nil)

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
		newForXmlDBusNodeInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusNodeInfo", "new_for_xml")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(xmlData)

	ret := newForXmlDBusNodeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_node_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var lookupInterfaceDBusNodeInfoInvoker *gi.Function

// LookupInterface is a representation of the C type g_dbus_node_info_lookup_interface.
func (recv *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	if lookupInterfaceDBusNodeInfoInvoker == nil {
		lookupInterfaceDBusNodeInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusNodeInfo", "lookup_interface")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := lookupInterfaceDBusNodeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo
}

var refDBusNodeInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_node_info_ref.
func (recv *DBusNodeInfo) Ref() *DBusNodeInfo {
	if refDBusNodeInfoInvoker == nil {
		refDBusNodeInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusNodeInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusNodeInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusNodeInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_node_info_unref.
func (recv *DBusNodeInfo) Unref() {
	if unrefDBusNodeInfoInvoker == nil {
		unrefDBusNodeInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusNodeInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusNodeInfoInvoker.Invoke(inArgs[:], nil)

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

var refDBusPropertyInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_property_info_ref.
func (recv *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	if refDBusPropertyInfoInvoker == nil {
		refDBusPropertyInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusPropertyInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusPropertyInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusPropertyInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_property_info_unref.
func (recv *DBusPropertyInfo) Unref() {
	if unrefDBusPropertyInfoInvoker == nil {
		unrefDBusPropertyInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusPropertyInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusPropertyInfoInvoker.Invoke(inArgs[:], nil)

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

var refDBusSignalInfoInvoker *gi.Function

// Ref is a representation of the C type g_dbus_signal_info_ref.
func (recv *DBusSignalInfo) Ref() *DBusSignalInfo {
	if refDBusSignalInfoInvoker == nil {
		refDBusSignalInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusSignalInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDBusSignalInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo
}

var unrefDBusSignalInfoInvoker *gi.Function

// Unref is a representation of the C type g_dbus_signal_info_unref.
func (recv *DBusSignalInfo) Unref() {
	if unrefDBusSignalInfoInvoker == nil {
		unrefDBusSignalInfoInvoker = gi.StructFunctionInvokerNew("Gio", "DBusSignalInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDBusSignalInfoInvoker.Invoke(inArgs[:], nil)

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
		newFileAttributeInfoListInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeInfoList", "new")
	}

	ret := newFileAttributeInfoListInvoker.Invoke(nil, nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_info_list_add' : parameter 'type' of type 'FileAttributeType' not supported

var dupFileAttributeInfoListInvoker *gi.Function

// Dup is a representation of the C type g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	if dupFileAttributeInfoListInvoker == nil {
		dupFileAttributeInfoListInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeInfoList", "dup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dupFileAttributeInfoListInvoker.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

var lookupFileAttributeInfoListInvoker *gi.Function

// Lookup is a representation of the C type g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	if lookupFileAttributeInfoListInvoker == nil {
		lookupFileAttributeInfoListInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeInfoList", "lookup")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := lookupFileAttributeInfoListInvoker.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfo{native: ret.Pointer()}

	return retGo
}

var refFileAttributeInfoListInvoker *gi.Function

// Ref is a representation of the C type g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	if refFileAttributeInfoListInvoker == nil {
		refFileAttributeInfoListInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeInfoList", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refFileAttributeInfoListInvoker.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

var unrefFileAttributeInfoListInvoker *gi.Function

// Unref is a representation of the C type g_file_attribute_info_list_unref.
func (recv *FileAttributeInfoList) Unref() {
	if unrefFileAttributeInfoListInvoker == nil {
		unrefFileAttributeInfoListInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeInfoList", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefFileAttributeInfoListInvoker.Invoke(inArgs[:], nil)

}

type FileAttributeMatcher struct {
	native uintptr
}

var newFileAttributeMatcherInvoker *gi.Function

// FileAttributeMatcherNew is a representation of the C type g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	if newFileAttributeMatcherInvoker == nil {
		newFileAttributeMatcherInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeMatcher", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(attributes)

	ret := newFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_enumerate_namespace' : return type 'gboolean' not supported

var enumerateNextFileAttributeMatcherInvoker *gi.Function

// EnumerateNext is a representation of the C type g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() string {
	if enumerateNextFileAttributeMatcherInvoker == nil {
		enumerateNextFileAttributeMatcherInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeMatcher", "enumerate_next")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := enumerateNextFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_file_attribute_matcher_matches_only' : return type 'gboolean' not supported

var refFileAttributeMatcherInvoker *gi.Function

// Ref is a representation of the C type g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	if refFileAttributeMatcherInvoker == nil {
		refFileAttributeMatcherInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeMatcher", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_subtract' : parameter 'subtract' of type 'FileAttributeMatcher' not supported

var toStringFileAttributeMatcherInvoker *gi.Function

// ToString is a representation of the C type g_file_attribute_matcher_to_string.
func (recv *FileAttributeMatcher) ToString() string {
	if toStringFileAttributeMatcherInvoker == nil {
		toStringFileAttributeMatcherInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeMatcher", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var unrefFileAttributeMatcherInvoker *gi.Function

// Unref is a representation of the C type g_file_attribute_matcher_unref.
func (recv *FileAttributeMatcher) Unref() {
	if unrefFileAttributeMatcherInvoker == nil {
		unrefFileAttributeMatcherInvoker = gi.StructFunctionInvokerNew("Gio", "FileAttributeMatcher", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefFileAttributeMatcherInvoker.Invoke(inArgs[:], nil)

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

var getNameIOExtensionInvoker *gi.Function

// GetName is a representation of the C type g_io_extension_get_name.
func (recv *IOExtension) GetName() string {
	if getNameIOExtensionInvoker == nil {
		getNameIOExtensionInvoker = gi.StructFunctionInvokerNew("Gio", "IOExtension", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameIOExtensionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPriorityIOExtensionInvoker *gi.Function

// GetPriority is a representation of the C type g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() int32 {
	if getPriorityIOExtensionInvoker == nil {
		getPriorityIOExtensionInvoker = gi.StructFunctionInvokerNew("Gio", "IOExtension", "get_priority")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPriorityIOExtensionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_get_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_ref_class' : return type 'GObject.TypeClass' not supported

type IOExtensionPoint struct {
	native uintptr
}

var getExtensionByNameIOExtensionPointInvoker *gi.Function

// GetExtensionByName is a representation of the C type g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	if getExtensionByNameIOExtensionPointInvoker == nil {
		getExtensionByNameIOExtensionPointInvoker = gi.StructFunctionInvokerNew("Gio", "IOExtensionPoint", "get_extension_by_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getExtensionByNameIOExtensionPointInvoker.Invoke(inArgs[:], nil)

	retGo := &IOExtension{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_point_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_io_extension_point_get_required_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_point_set_required_type' : parameter 'type' of type 'GType' not supported

type IOModuleClass struct {
	native uintptr
}

type IOModuleScope struct {
	native uintptr
}

var blockIOModuleScopeInvoker *gi.Function

// Block is a representation of the C type g_io_module_scope_block.
func (recv *IOModuleScope) Block(basename string) {
	if blockIOModuleScopeInvoker == nil {
		blockIOModuleScopeInvoker = gi.StructFunctionInvokerNew("Gio", "IOModuleScope", "block")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(basename)

	blockIOModuleScopeInvoker.Invoke(inArgs[:], nil)

}

var freeIOModuleScopeInvoker *gi.Function

// Free is a representation of the C type g_io_module_scope_free.
func (recv *IOModuleScope) Free() {
	if freeIOModuleScopeInvoker == nil {
		freeIOModuleScopeInvoker = gi.StructFunctionInvokerNew("Gio", "IOModuleScope", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeIOModuleScopeInvoker.Invoke(inArgs[:], nil)

}

type IOSchedulerJob struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

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

var RegisterResourceInvoker *gi.Function

// Register is a representation of the C type g_resources_register.
func (recv *Resource) Register() {
	if RegisterResourceInvoker == nil {
		RegisterResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "_register")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	RegisterResourceInvoker.Invoke(inArgs[:], nil)

}

var UnregisterResourceInvoker *gi.Function

// Unregister is a representation of the C type g_resources_unregister.
func (recv *Resource) Unregister() {
	if UnregisterResourceInvoker == nil {
		UnregisterResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "_unregister")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	UnregisterResourceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_resource_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var refResourceInvoker *gi.Function

// Ref is a representation of the C type g_resource_ref.
func (recv *Resource) Ref() *Resource {
	if refResourceInvoker == nil {
		refResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refResourceInvoker.Invoke(inArgs[:], nil)

	retGo := &Resource{native: ret.Pointer()}

	return retGo
}

var unrefResourceInvoker *gi.Function

// Unref is a representation of the C type g_resource_unref.
func (recv *Resource) Unref() {
	if unrefResourceInvoker == nil {
		unrefResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefResourceInvoker.Invoke(inArgs[:], nil)

}

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

var getIdSettingsSchemaInvoker *gi.Function

// GetId is a representation of the C type g_settings_schema_get_id.
func (recv *SettingsSchema) GetId() string {
	if getIdSettingsSchemaInvoker == nil {
		getIdSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "get_id")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIdSettingsSchemaInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getKeySettingsSchemaInvoker *gi.Function

// GetKey is a representation of the C type g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	if getKeySettingsSchemaInvoker == nil {
		getKeySettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "get_key")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getKeySettingsSchemaInvoker.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo
}

var getPathSettingsSchemaInvoker *gi.Function

// GetPath is a representation of the C type g_settings_schema_get_path.
func (recv *SettingsSchema) GetPath() string {
	if getPathSettingsSchemaInvoker == nil {
		getPathSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathSettingsSchemaInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_has_key' : return type 'gboolean' not supported

var listChildrenSettingsSchemaInvoker *gi.Function

// ListChildren is a representation of the C type g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() {
	if listChildrenSettingsSchemaInvoker == nil {
		listChildrenSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "list_children")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	listChildrenSettingsSchemaInvoker.Invoke(inArgs[:], nil)

}

var listKeysSettingsSchemaInvoker *gi.Function

// ListKeys is a representation of the C type g_settings_schema_list_keys.
func (recv *SettingsSchema) ListKeys() {
	if listKeysSettingsSchemaInvoker == nil {
		listKeysSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "list_keys")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	listKeysSettingsSchemaInvoker.Invoke(inArgs[:], nil)

}

var refSettingsSchemaInvoker *gi.Function

// Ref is a representation of the C type g_settings_schema_ref.
func (recv *SettingsSchema) Ref() *SettingsSchema {
	if refSettingsSchemaInvoker == nil {
		refSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSettingsSchemaInvoker.Invoke(inArgs[:], nil)

	retGo := &SettingsSchema{native: ret.Pointer()}

	return retGo
}

var unrefSettingsSchemaInvoker *gi.Function

// Unref is a representation of the C type g_settings_schema_unref.
func (recv *SettingsSchema) Unref() {
	if unrefSettingsSchemaInvoker == nil {
		unrefSettingsSchemaInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchema", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSettingsSchemaInvoker.Invoke(inArgs[:], nil)

}

type SettingsSchemaKey struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_default_value' : return type 'GLib.Variant' not supported

var getDescriptionSettingsSchemaKeyInvoker *gi.Function

// GetDescription is a representation of the C type g_settings_schema_key_get_description.
func (recv *SettingsSchemaKey) GetDescription() string {
	if getDescriptionSettingsSchemaKeyInvoker == nil {
		getDescriptionSettingsSchemaKeyInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaKey", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionSettingsSchemaKeyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getNameSettingsSchemaKeyInvoker *gi.Function

// GetName is a representation of the C type g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() string {
	if getNameSettingsSchemaKeyInvoker == nil {
		getNameSettingsSchemaKeyInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaKey", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameSettingsSchemaKeyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_range' : return type 'GLib.Variant' not supported

var getSummarySettingsSchemaKeyInvoker *gi.Function

// GetSummary is a representation of the C type g_settings_schema_key_get_summary.
func (recv *SettingsSchemaKey) GetSummary() string {
	if getSummarySettingsSchemaKeyInvoker == nil {
		getSummarySettingsSchemaKeyInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaKey", "get_summary")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSummarySettingsSchemaKeyInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_value_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_settings_schema_key_range_check' : parameter 'value' of type 'GLib.Variant' not supported

var refSettingsSchemaKeyInvoker *gi.Function

// Ref is a representation of the C type g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	if refSettingsSchemaKeyInvoker == nil {
		refSettingsSchemaKeyInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaKey", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSettingsSchemaKeyInvoker.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo
}

var unrefSettingsSchemaKeyInvoker *gi.Function

// Unref is a representation of the C type g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	if unrefSettingsSchemaKeyInvoker == nil {
		unrefSettingsSchemaKeyInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaKey", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSettingsSchemaKeyInvoker.Invoke(inArgs[:], nil)

}

type SettingsSchemaSource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_source_new_from_directory' : parameter 'directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_list_schemas' : parameter 'recursive' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_lookup' : parameter 'recursive' of type 'gboolean' not supported

var refSettingsSchemaSourceInvoker *gi.Function

// Ref is a representation of the C type g_settings_schema_source_ref.
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	if refSettingsSchemaSourceInvoker == nil {
		refSettingsSchemaSourceInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaSource", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSettingsSchemaSourceInvoker.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo
}

var unrefSettingsSchemaSourceInvoker *gi.Function

// Unref is a representation of the C type g_settings_schema_source_unref.
func (recv *SettingsSchemaSource) Unref() {
	if unrefSettingsSchemaSourceInvoker == nil {
		unrefSettingsSchemaSourceInvoker = gi.StructFunctionInvokerNew("Gio", "SettingsSchemaSource", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSettingsSchemaSourceInvoker.Invoke(inArgs[:], nil)

}

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
		newSrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "new")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)
	inArgs[2].SetUint16(priority)
	inArgs[3].SetUint16(weight)

	ret := newSrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo
}

var copySrvTargetInvoker *gi.Function

// Copy is a representation of the C type g_srv_target_copy.
func (recv *SrvTarget) Copy() *SrvTarget {
	if copySrvTargetInvoker == nil {
		copySrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copySrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo
}

var freeSrvTargetInvoker *gi.Function

// Free is a representation of the C type g_srv_target_free.
func (recv *SrvTarget) Free() {
	if freeSrvTargetInvoker == nil {
		freeSrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeSrvTargetInvoker.Invoke(inArgs[:], nil)

}

var getHostnameSrvTargetInvoker *gi.Function

// GetHostname is a representation of the C type g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() string {
	if getHostnameSrvTargetInvoker == nil {
		getHostnameSrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "get_hostname")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHostnameSrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPortSrvTargetInvoker *gi.Function

// GetPort is a representation of the C type g_srv_target_get_port.
func (recv *SrvTarget) GetPort() uint16 {
	if getPortSrvTargetInvoker == nil {
		getPortSrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "get_port")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPortSrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var getPrioritySrvTargetInvoker *gi.Function

// GetPriority is a representation of the C type g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() uint16 {
	if getPrioritySrvTargetInvoker == nil {
		getPrioritySrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "get_priority")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPrioritySrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var getWeightSrvTargetInvoker *gi.Function

// GetWeight is a representation of the C type g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() uint16 {
	if getWeightSrvTargetInvoker == nil {
		getWeightSrvTargetInvoker = gi.StructFunctionInvokerNew("Gio", "SrvTarget", "get_weight")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getWeightSrvTargetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

type StaticResource struct {
	native uintptr
}

var finiStaticResourceInvoker *gi.Function

// Fini is a representation of the C type g_static_resource_fini.
func (recv *StaticResource) Fini() {
	if finiStaticResourceInvoker == nil {
		finiStaticResourceInvoker = gi.StructFunctionInvokerNew("Gio", "StaticResource", "fini")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	finiStaticResourceInvoker.Invoke(inArgs[:], nil)

}

var getResourceStaticResourceInvoker *gi.Function

// GetResource is a representation of the C type g_static_resource_get_resource.
func (recv *StaticResource) GetResource() *Resource {
	if getResourceStaticResourceInvoker == nil {
		getResourceStaticResourceInvoker = gi.StructFunctionInvokerNew("Gio", "StaticResource", "get_resource")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getResourceStaticResourceInvoker.Invoke(inArgs[:], nil)

	retGo := &Resource{native: ret.Pointer()}

	return retGo
}

var initStaticResourceInvoker *gi.Function

// Init is a representation of the C type g_static_resource_init.
func (recv *StaticResource) Init() {
	if initStaticResourceInvoker == nil {
		initStaticResourceInvoker = gi.StructFunctionInvokerNew("Gio", "StaticResource", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initStaticResourceInvoker.Invoke(inArgs[:], nil)

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

// UNSUPPORTED : C value 'g_unix_mount_point_compare' : parameter 'mount2' of type 'UnixMountPoint' not supported

var copyUnixMountPointInvoker *gi.Function

// Copy is a representation of the C type g_unix_mount_point_copy.
func (recv *UnixMountPoint) Copy() *UnixMountPoint {
	if copyUnixMountPointInvoker == nil {
		copyUnixMountPointInvoker = gi.StructFunctionInvokerNew("Gio", "UnixMountPoint", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyUnixMountPointInvoker.Invoke(inArgs[:], nil)

	retGo := &UnixMountPoint{native: ret.Pointer()}

	return retGo
}

var freeUnixMountPointInvoker *gi.Function

// Free is a representation of the C type g_unix_mount_point_free.
func (recv *UnixMountPoint) Free() {
	if freeUnixMountPointInvoker == nil {
		freeUnixMountPointInvoker = gi.StructFunctionInvokerNew("Gio", "UnixMountPoint", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeUnixMountPointInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_unix_mount_point_get_device_path' : return type 'filename' not supported

var getFsTypeUnixMountPointInvoker *gi.Function

// GetFsType is a representation of the C type g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() string {
	if getFsTypeUnixMountPointInvoker == nil {
		getFsTypeUnixMountPointInvoker = gi.StructFunctionInvokerNew("Gio", "UnixMountPoint", "get_fs_type")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFsTypeUnixMountPointInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_get_mount_path' : return type 'filename' not supported

var getOptionsUnixMountPointInvoker *gi.Function

// GetOptions is a representation of the C type g_unix_mount_point_get_options.
func (recv *UnixMountPoint) GetOptions() string {
	if getOptionsUnixMountPointInvoker == nil {
		getOptionsUnixMountPointInvoker = gi.StructFunctionInvokerNew("Gio", "UnixMountPoint", "get_options")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getOptionsUnixMountPointInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_can_eject' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_guess_icon' : return type 'Icon' not supported

var guessNameUnixMountPointInvoker *gi.Function

// GuessName is a representation of the C type g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() string {
	if guessNameUnixMountPointInvoker == nil {
		guessNameUnixMountPointInvoker = gi.StructFunctionInvokerNew("Gio", "UnixMountPoint", "guess_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := guessNameUnixMountPointInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_loopback' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_readonly' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_user_mountable' : return type 'gboolean' not supported

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
