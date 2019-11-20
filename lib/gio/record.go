// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var ActionEntryStruct *gi.Struct
var ActionEntryStructOnce sync.Once

func ActionEntryStructSet() {
	ActionEntryStructOnce.Do(func() {
		ActionEntryStruct = gi.StructNew("Gio", "ActionEntry")
	})
}

type ActionEntry struct {
	native uintptr
	Name   string
	// UNSUPPORTED : C value 'activate' : missing Type
	ParameterType string
	State         string
	// UNSUPPORTED : C value 'change_state' : missing Type
}

var ActionGroupInterfaceStruct *gi.Struct
var ActionGroupInterfaceStructOnce sync.Once

func ActionGroupInterfaceStructSet() {
	ActionGroupInterfaceStructOnce.Do(func() {
		ActionGroupInterfaceStruct = gi.StructNew("Gio", "ActionGroupInterface")
	})
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

var ActionInterfaceStruct *gi.Struct
var ActionInterfaceStructOnce sync.Once

func ActionInterfaceStructSet() {
	ActionInterfaceStructOnce.Do(func() {
		ActionInterfaceStruct = gi.StructNew("Gio", "ActionInterface")
	})
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

var ActionMapInterfaceStruct *gi.Struct
var ActionMapInterfaceStructOnce sync.Once

func ActionMapInterfaceStructSet() {
	ActionMapInterfaceStructOnce.Do(func() {
		ActionMapInterfaceStruct = gi.StructNew("Gio", "ActionMapInterface")
	})
}

type ActionMapInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'lookup_action' : missing Type
	// UNSUPPORTED : C value 'add_action' : missing Type
	// UNSUPPORTED : C value 'remove_action' : missing Type
}

var AppInfoIfaceStruct *gi.Struct
var AppInfoIfaceStructOnce sync.Once

func AppInfoIfaceStructSet() {
	AppInfoIfaceStructOnce.Do(func() {
		AppInfoIfaceStruct = gi.StructNew("Gio", "AppInfoIface")
	})
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

var AppLaunchContextClassStruct *gi.Struct
var AppLaunchContextClassStructOnce sync.Once

func AppLaunchContextClassStructSet() {
	AppLaunchContextClassStructOnce.Do(func() {
		AppLaunchContextClassStruct = gi.StructNew("Gio", "AppLaunchContextClass")
	})
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

var AppLaunchContextPrivateStruct *gi.Struct
var AppLaunchContextPrivateStructOnce sync.Once

func AppLaunchContextPrivateStructSet() {
	AppLaunchContextPrivateStructOnce.Do(func() {
		AppLaunchContextPrivateStruct = gi.StructNew("Gio", "AppLaunchContextPrivate")
	})
}

type AppLaunchContextPrivate struct {
	native uintptr
}

var ApplicationClassStruct *gi.Struct
var ApplicationClassStructOnce sync.Once

func ApplicationClassStructSet() {
	ApplicationClassStructOnce.Do(func() {
		ApplicationClassStruct = gi.StructNew("Gio", "ApplicationClass")
	})
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

var ApplicationCommandLineClassStruct *gi.Struct
var ApplicationCommandLineClassStructOnce sync.Once

func ApplicationCommandLineClassStructSet() {
	ApplicationCommandLineClassStructOnce.Do(func() {
		ApplicationCommandLineClassStruct = gi.StructNew("Gio", "ApplicationCommandLineClass")
	})
}

type ApplicationCommandLineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'print_literal' : missing Type
	// UNSUPPORTED : C value 'printerr_literal' : missing Type
	// UNSUPPORTED : C value 'get_stdin' : missing Type
}

var ApplicationCommandLinePrivateStruct *gi.Struct
var ApplicationCommandLinePrivateStructOnce sync.Once

func ApplicationCommandLinePrivateStructSet() {
	ApplicationCommandLinePrivateStructOnce.Do(func() {
		ApplicationCommandLinePrivateStruct = gi.StructNew("Gio", "ApplicationCommandLinePrivate")
	})
}

type ApplicationCommandLinePrivate struct {
	native uintptr
}

var ApplicationPrivateStruct *gi.Struct
var ApplicationPrivateStructOnce sync.Once

func ApplicationPrivateStructSet() {
	ApplicationPrivateStructOnce.Do(func() {
		ApplicationPrivateStruct = gi.StructNew("Gio", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var AsyncInitableIfaceStruct *gi.Struct
var AsyncInitableIfaceStructOnce sync.Once

func AsyncInitableIfaceStructSet() {
	AsyncInitableIfaceStructOnce.Do(func() {
		AsyncInitableIfaceStruct = gi.StructNew("Gio", "AsyncInitableIface")
	})
}

type AsyncInitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'init_async' : missing Type
	// UNSUPPORTED : C value 'init_finish' : missing Type
}

var AsyncResultIfaceStruct *gi.Struct
var AsyncResultIfaceStructOnce sync.Once

func AsyncResultIfaceStructSet() {
	AsyncResultIfaceStructOnce.Do(func() {
		AsyncResultIfaceStruct = gi.StructNew("Gio", "AsyncResultIface")
	})
}

type AsyncResultIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_user_data' : missing Type
	// UNSUPPORTED : C value 'get_source_object' : missing Type
	// UNSUPPORTED : C value 'is_tagged' : missing Type
}

var BufferedInputStreamClassStruct *gi.Struct
var BufferedInputStreamClassStructOnce sync.Once

func BufferedInputStreamClassStructSet() {
	BufferedInputStreamClassStructOnce.Do(func() {
		BufferedInputStreamClassStruct = gi.StructNew("Gio", "BufferedInputStreamClass")
	})
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

var BufferedInputStreamPrivateStruct *gi.Struct
var BufferedInputStreamPrivateStructOnce sync.Once

func BufferedInputStreamPrivateStructSet() {
	BufferedInputStreamPrivateStructOnce.Do(func() {
		BufferedInputStreamPrivateStruct = gi.StructNew("Gio", "BufferedInputStreamPrivate")
	})
}

type BufferedInputStreamPrivate struct {
	native uintptr
}

var BufferedOutputStreamClassStruct *gi.Struct
var BufferedOutputStreamClassStructOnce sync.Once

func BufferedOutputStreamClassStructSet() {
	BufferedOutputStreamClassStructOnce.Do(func() {
		BufferedOutputStreamClassStruct = gi.StructNew("Gio", "BufferedOutputStreamClass")
	})
}

type BufferedOutputStreamClass struct {
	native      uintptr
	ParentClass *FilterOutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var BufferedOutputStreamPrivateStruct *gi.Struct
var BufferedOutputStreamPrivateStructOnce sync.Once

func BufferedOutputStreamPrivateStructSet() {
	BufferedOutputStreamPrivateStructOnce.Do(func() {
		BufferedOutputStreamPrivateStruct = gi.StructNew("Gio", "BufferedOutputStreamPrivate")
	})
}

type BufferedOutputStreamPrivate struct {
	native uintptr
}

var CancellableClassStruct *gi.Struct
var CancellableClassStructOnce sync.Once

func CancellableClassStructSet() {
	CancellableClassStructOnce.Do(func() {
		CancellableClassStruct = gi.StructNew("Gio", "CancellableClass")
	})
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

var CancellablePrivateStruct *gi.Struct
var CancellablePrivateStructOnce sync.Once

func CancellablePrivateStructSet() {
	CancellablePrivateStructOnce.Do(func() {
		CancellablePrivateStruct = gi.StructNew("Gio", "CancellablePrivate")
	})
}

type CancellablePrivate struct {
	native uintptr
}

var CharsetConverterClassStruct *gi.Struct
var CharsetConverterClassStructOnce sync.Once

func CharsetConverterClassStructSet() {
	CharsetConverterClassStructOnce.Do(func() {
		CharsetConverterClassStruct = gi.StructNew("Gio", "CharsetConverterClass")
	})
}

type CharsetConverterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var ConverterIfaceStruct *gi.Struct
var ConverterIfaceStructOnce sync.Once

func ConverterIfaceStructSet() {
	ConverterIfaceStructOnce.Do(func() {
		ConverterIfaceStruct = gi.StructNew("Gio", "ConverterIface")
	})
}

type ConverterIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'convert' : missing Type
	// UNSUPPORTED : C value 'reset' : missing Type
}

var ConverterInputStreamClassStruct *gi.Struct
var ConverterInputStreamClassStructOnce sync.Once

func ConverterInputStreamClassStructSet() {
	ConverterInputStreamClassStructOnce.Do(func() {
		ConverterInputStreamClassStruct = gi.StructNew("Gio", "ConverterInputStreamClass")
	})
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

var ConverterInputStreamPrivateStruct *gi.Struct
var ConverterInputStreamPrivateStructOnce sync.Once

func ConverterInputStreamPrivateStructSet() {
	ConverterInputStreamPrivateStructOnce.Do(func() {
		ConverterInputStreamPrivateStruct = gi.StructNew("Gio", "ConverterInputStreamPrivate")
	})
}

type ConverterInputStreamPrivate struct {
	native uintptr
}

var ConverterOutputStreamClassStruct *gi.Struct
var ConverterOutputStreamClassStructOnce sync.Once

func ConverterOutputStreamClassStructSet() {
	ConverterOutputStreamClassStructOnce.Do(func() {
		ConverterOutputStreamClassStruct = gi.StructNew("Gio", "ConverterOutputStreamClass")
	})
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

var ConverterOutputStreamPrivateStruct *gi.Struct
var ConverterOutputStreamPrivateStructOnce sync.Once

func ConverterOutputStreamPrivateStructSet() {
	ConverterOutputStreamPrivateStructOnce.Do(func() {
		ConverterOutputStreamPrivateStruct = gi.StructNew("Gio", "ConverterOutputStreamPrivate")
	})
}

type ConverterOutputStreamPrivate struct {
	native uintptr
}

var CredentialsClassStruct *gi.Struct
var CredentialsClassStructOnce sync.Once

func CredentialsClassStructSet() {
	CredentialsClassStructOnce.Do(func() {
		CredentialsClassStruct = gi.StructNew("Gio", "CredentialsClass")
	})
}

type CredentialsClass struct {
	native uintptr
}

var DBusAnnotationInfoStruct *gi.Struct
var DBusAnnotationInfoStructOnce sync.Once

func DBusAnnotationInfoStructSet() {
	DBusAnnotationInfoStructOnce.Do(func() {
		DBusAnnotationInfoStruct = gi.StructNew("Gio", "DBusAnnotationInfo")
	})
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

var DBusArgInfoStruct *gi.Struct
var DBusArgInfoStructOnce sync.Once

func DBusArgInfoStructSet() {
	DBusArgInfoStructOnce.Do(func() {
		DBusArgInfoStruct = gi.StructNew("Gio", "DBusArgInfo")
	})
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

var DBusErrorEntryStruct *gi.Struct
var DBusErrorEntryStructOnce sync.Once

func DBusErrorEntryStructSet() {
	DBusErrorEntryStructOnce.Do(func() {
		DBusErrorEntryStruct = gi.StructNew("Gio", "DBusErrorEntry")
	})
}

type DBusErrorEntry struct {
	native        uintptr
	ErrorCode     int32
	DbusErrorName string
}

var DBusInterfaceIfaceStruct *gi.Struct
var DBusInterfaceIfaceStructOnce sync.Once

func DBusInterfaceIfaceStructSet() {
	DBusInterfaceIfaceStructOnce.Do(func() {
		DBusInterfaceIfaceStruct = gi.StructNew("Gio", "DBusInterfaceIface")
	})
}

type DBusInterfaceIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_info' : missing Type
	// UNSUPPORTED : C value 'get_object' : missing Type
	// UNSUPPORTED : C value 'set_object' : missing Type
	// UNSUPPORTED : C value 'dup_object' : missing Type
}

var DBusInterfaceInfoStruct *gi.Struct
var DBusInterfaceInfoStructOnce sync.Once

func DBusInterfaceInfoStructSet() {
	DBusInterfaceInfoStructOnce.Do(func() {
		DBusInterfaceInfoStruct = gi.StructNew("Gio", "DBusInterfaceInfo")
	})
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

var DBusInterfaceSkeletonClassStruct *gi.Struct
var DBusInterfaceSkeletonClassStructOnce sync.Once

func DBusInterfaceSkeletonClassStructSet() {
	DBusInterfaceSkeletonClassStructOnce.Do(func() {
		DBusInterfaceSkeletonClassStruct = gi.StructNew("Gio", "DBusInterfaceSkeletonClass")
	})
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

var DBusInterfaceSkeletonPrivateStruct *gi.Struct
var DBusInterfaceSkeletonPrivateStructOnce sync.Once

func DBusInterfaceSkeletonPrivateStructSet() {
	DBusInterfaceSkeletonPrivateStructOnce.Do(func() {
		DBusInterfaceSkeletonPrivateStruct = gi.StructNew("Gio", "DBusInterfaceSkeletonPrivate")
	})
}

type DBusInterfaceSkeletonPrivate struct {
	native uintptr
}

var DBusInterfaceVTableStruct *gi.Struct
var DBusInterfaceVTableStructOnce sync.Once

func DBusInterfaceVTableStructSet() {
	DBusInterfaceVTableStructOnce.Do(func() {
		DBusInterfaceVTableStruct = gi.StructNew("Gio", "DBusInterfaceVTable")
	})
}

type DBusInterfaceVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'method_call' : no Go type for 'DBusInterfaceMethodCallFunc'
	// UNSUPPORTED : C value 'get_property' : no Go type for 'DBusInterfaceGetPropertyFunc'
	// UNSUPPORTED : C value 'set_property' : no Go type for 'DBusInterfaceSetPropertyFunc'
}

var DBusMethodInfoStruct *gi.Struct
var DBusMethodInfoStructOnce sync.Once

func DBusMethodInfoStructSet() {
	DBusMethodInfoStructOnce.Do(func() {
		DBusMethodInfoStruct = gi.StructNew("Gio", "DBusMethodInfo")
	})
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

var DBusNodeInfoStruct *gi.Struct
var DBusNodeInfoStructOnce sync.Once

func DBusNodeInfoStructSet() {
	DBusNodeInfoStructOnce.Do(func() {
		DBusNodeInfoStruct = gi.StructNew("Gio", "DBusNodeInfo")
	})
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

var DBusObjectIfaceStruct *gi.Struct
var DBusObjectIfaceStructOnce sync.Once

func DBusObjectIfaceStructSet() {
	DBusObjectIfaceStructOnce.Do(func() {
		DBusObjectIfaceStruct = gi.StructNew("Gio", "DBusObjectIface")
	})
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

var DBusObjectManagerClientClassStruct *gi.Struct
var DBusObjectManagerClientClassStructOnce sync.Once

func DBusObjectManagerClientClassStructSet() {
	DBusObjectManagerClientClassStructOnce.Do(func() {
		DBusObjectManagerClientClassStruct = gi.StructNew("Gio", "DBusObjectManagerClientClass")
	})
}

type DBusObjectManagerClientClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'interface_proxy_signal' : missing Type
	// UNSUPPORTED : C value 'interface_proxy_properties_changed' : missing Type
}

var DBusObjectManagerClientPrivateStruct *gi.Struct
var DBusObjectManagerClientPrivateStructOnce sync.Once

func DBusObjectManagerClientPrivateStructSet() {
	DBusObjectManagerClientPrivateStructOnce.Do(func() {
		DBusObjectManagerClientPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerClientPrivate")
	})
}

type DBusObjectManagerClientPrivate struct {
	native uintptr
}

var DBusObjectManagerIfaceStruct *gi.Struct
var DBusObjectManagerIfaceStructOnce sync.Once

func DBusObjectManagerIfaceStructSet() {
	DBusObjectManagerIfaceStructOnce.Do(func() {
		DBusObjectManagerIfaceStruct = gi.StructNew("Gio", "DBusObjectManagerIface")
	})
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

var DBusObjectManagerServerClassStruct *gi.Struct
var DBusObjectManagerServerClassStructOnce sync.Once

func DBusObjectManagerServerClassStructSet() {
	DBusObjectManagerServerClassStructOnce.Do(func() {
		DBusObjectManagerServerClassStruct = gi.StructNew("Gio", "DBusObjectManagerServerClass")
	})
}

type DBusObjectManagerServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var DBusObjectManagerServerPrivateStruct *gi.Struct
var DBusObjectManagerServerPrivateStructOnce sync.Once

func DBusObjectManagerServerPrivateStructSet() {
	DBusObjectManagerServerPrivateStructOnce.Do(func() {
		DBusObjectManagerServerPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerServerPrivate")
	})
}

type DBusObjectManagerServerPrivate struct {
	native uintptr
}

var DBusObjectProxyClassStruct *gi.Struct
var DBusObjectProxyClassStructOnce sync.Once

func DBusObjectProxyClassStructSet() {
	DBusObjectProxyClassStructOnce.Do(func() {
		DBusObjectProxyClassStruct = gi.StructNew("Gio", "DBusObjectProxyClass")
	})
}

type DBusObjectProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var DBusObjectProxyPrivateStruct *gi.Struct
var DBusObjectProxyPrivateStructOnce sync.Once

func DBusObjectProxyPrivateStructSet() {
	DBusObjectProxyPrivateStructOnce.Do(func() {
		DBusObjectProxyPrivateStruct = gi.StructNew("Gio", "DBusObjectProxyPrivate")
	})
}

type DBusObjectProxyPrivate struct {
	native uintptr
}

var DBusObjectSkeletonClassStruct *gi.Struct
var DBusObjectSkeletonClassStructOnce sync.Once

func DBusObjectSkeletonClassStructSet() {
	DBusObjectSkeletonClassStructOnce.Do(func() {
		DBusObjectSkeletonClassStruct = gi.StructNew("Gio", "DBusObjectSkeletonClass")
	})
}

type DBusObjectSkeletonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authorize_method' : missing Type
}

var DBusObjectSkeletonPrivateStruct *gi.Struct
var DBusObjectSkeletonPrivateStructOnce sync.Once

func DBusObjectSkeletonPrivateStructSet() {
	DBusObjectSkeletonPrivateStructOnce.Do(func() {
		DBusObjectSkeletonPrivateStruct = gi.StructNew("Gio", "DBusObjectSkeletonPrivate")
	})
}

type DBusObjectSkeletonPrivate struct {
	native uintptr
}

var DBusPropertyInfoStruct *gi.Struct
var DBusPropertyInfoStructOnce sync.Once

func DBusPropertyInfoStructSet() {
	DBusPropertyInfoStructOnce.Do(func() {
		DBusPropertyInfoStruct = gi.StructNew("Gio", "DBusPropertyInfo")
	})
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

var DBusProxyClassStruct *gi.Struct
var DBusProxyClassStructOnce sync.Once

func DBusProxyClassStructSet() {
	DBusProxyClassStructOnce.Do(func() {
		DBusProxyClassStruct = gi.StructNew("Gio", "DBusProxyClass")
	})
}

type DBusProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'g_properties_changed' : missing Type
	// UNSUPPORTED : C value 'g_signal' : missing Type
}

var DBusProxyPrivateStruct *gi.Struct
var DBusProxyPrivateStructOnce sync.Once

func DBusProxyPrivateStructSet() {
	DBusProxyPrivateStructOnce.Do(func() {
		DBusProxyPrivateStruct = gi.StructNew("Gio", "DBusProxyPrivate")
	})
}

type DBusProxyPrivate struct {
	native uintptr
}

var DBusSignalInfoStruct *gi.Struct
var DBusSignalInfoStructOnce sync.Once

func DBusSignalInfoStructSet() {
	DBusSignalInfoStructOnce.Do(func() {
		DBusSignalInfoStruct = gi.StructNew("Gio", "DBusSignalInfo")
	})
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

var DBusSubtreeVTableStruct *gi.Struct
var DBusSubtreeVTableStructOnce sync.Once

func DBusSubtreeVTableStructSet() {
	DBusSubtreeVTableStructOnce.Do(func() {
		DBusSubtreeVTableStruct = gi.StructNew("Gio", "DBusSubtreeVTable")
	})
}

type DBusSubtreeVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'enumerate' : no Go type for 'DBusSubtreeEnumerateFunc'
	// UNSUPPORTED : C value 'introspect' : no Go type for 'DBusSubtreeIntrospectFunc'
	// UNSUPPORTED : C value 'dispatch' : no Go type for 'DBusSubtreeDispatchFunc'
}

var DataInputStreamClassStruct *gi.Struct
var DataInputStreamClassStructOnce sync.Once

func DataInputStreamClassStructSet() {
	DataInputStreamClassStructOnce.Do(func() {
		DataInputStreamClassStruct = gi.StructNew("Gio", "DataInputStreamClass")
	})
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

var DataInputStreamPrivateStruct *gi.Struct
var DataInputStreamPrivateStructOnce sync.Once

func DataInputStreamPrivateStructSet() {
	DataInputStreamPrivateStructOnce.Do(func() {
		DataInputStreamPrivateStruct = gi.StructNew("Gio", "DataInputStreamPrivate")
	})
}

type DataInputStreamPrivate struct {
	native uintptr
}

var DataOutputStreamClassStruct *gi.Struct
var DataOutputStreamClassStructOnce sync.Once

func DataOutputStreamClassStructSet() {
	DataOutputStreamClassStructOnce.Do(func() {
		DataOutputStreamClassStruct = gi.StructNew("Gio", "DataOutputStreamClass")
	})
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

var DataOutputStreamPrivateStruct *gi.Struct
var DataOutputStreamPrivateStructOnce sync.Once

func DataOutputStreamPrivateStructSet() {
	DataOutputStreamPrivateStructOnce.Do(func() {
		DataOutputStreamPrivateStruct = gi.StructNew("Gio", "DataOutputStreamPrivate")
	})
}

type DataOutputStreamPrivate struct {
	native uintptr
}

var DatagramBasedInterfaceStruct *gi.Struct
var DatagramBasedInterfaceStructOnce sync.Once

func DatagramBasedInterfaceStructSet() {
	DatagramBasedInterfaceStructOnce.Do(func() {
		DatagramBasedInterfaceStruct = gi.StructNew("Gio", "DatagramBasedInterface")
	})
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

var DesktopAppInfoClassStruct *gi.Struct
var DesktopAppInfoClassStructOnce sync.Once

func DesktopAppInfoClassStructSet() {
	DesktopAppInfoClassStructOnce.Do(func() {
		DesktopAppInfoClassStruct = gi.StructNew("Gio", "DesktopAppInfoClass")
	})
}

type DesktopAppInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var DesktopAppInfoLookupIfaceStruct *gi.Struct
var DesktopAppInfoLookupIfaceStructOnce sync.Once

func DesktopAppInfoLookupIfaceStructSet() {
	DesktopAppInfoLookupIfaceStructOnce.Do(func() {
		DesktopAppInfoLookupIfaceStruct = gi.StructNew("Gio", "DesktopAppInfoLookupIface")
	})
}

type DesktopAppInfoLookupIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_default_for_uri_scheme' : missing Type
}

var DriveIfaceStruct *gi.Struct
var DriveIfaceStructOnce sync.Once

func DriveIfaceStructSet() {
	DriveIfaceStructOnce.Do(func() {
		DriveIfaceStruct = gi.StructNew("Gio", "DriveIface")
	})
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

var DtlsClientConnectionInterfaceStruct *gi.Struct
var DtlsClientConnectionInterfaceStructOnce sync.Once

func DtlsClientConnectionInterfaceStructSet() {
	DtlsClientConnectionInterfaceStructOnce.Do(func() {
		DtlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsClientConnectionInterface")
	})
}

type DtlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var DtlsConnectionInterfaceStruct *gi.Struct
var DtlsConnectionInterfaceStructOnce sync.Once

func DtlsConnectionInterfaceStructSet() {
	DtlsConnectionInterfaceStructOnce.Do(func() {
		DtlsConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsConnectionInterface")
	})
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

var DtlsServerConnectionInterfaceStruct *gi.Struct
var DtlsServerConnectionInterfaceStructOnce sync.Once

func DtlsServerConnectionInterfaceStructSet() {
	DtlsServerConnectionInterfaceStructOnce.Do(func() {
		DtlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsServerConnectionInterface")
	})
}

type DtlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var EmblemClassStruct *gi.Struct
var EmblemClassStructOnce sync.Once

func EmblemClassStructSet() {
	EmblemClassStructOnce.Do(func() {
		EmblemClassStruct = gi.StructNew("Gio", "EmblemClass")
	})
}

type EmblemClass struct {
	native uintptr
}

var EmblemedIconClassStruct *gi.Struct
var EmblemedIconClassStructOnce sync.Once

func EmblemedIconClassStructSet() {
	EmblemedIconClassStructOnce.Do(func() {
		EmblemedIconClassStruct = gi.StructNew("Gio", "EmblemedIconClass")
	})
}

type EmblemedIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var EmblemedIconPrivateStruct *gi.Struct
var EmblemedIconPrivateStructOnce sync.Once

func EmblemedIconPrivateStructSet() {
	EmblemedIconPrivateStructOnce.Do(func() {
		EmblemedIconPrivateStruct = gi.StructNew("Gio", "EmblemedIconPrivate")
	})
}

type EmblemedIconPrivate struct {
	native uintptr
}

var FileAttributeInfoStruct *gi.Struct
var FileAttributeInfoStructOnce sync.Once

func FileAttributeInfoStructSet() {
	FileAttributeInfoStructOnce.Do(func() {
		FileAttributeInfoStruct = gi.StructNew("Gio", "FileAttributeInfo")
	})
}

type FileAttributeInfo struct {
	native uintptr
	Name   string
	// UNSUPPORTED : C value 'type' : no Go type for 'FileAttributeType'
	// UNSUPPORTED : C value 'flags' : no Go type for 'FileAttributeInfoFlags'
}

var FileAttributeInfoListStruct *gi.Struct
var FileAttributeInfoListStructOnce sync.Once

func FileAttributeInfoListStructSet() {
	FileAttributeInfoListStructOnce.Do(func() {
		FileAttributeInfoListStruct = gi.StructNew("Gio", "FileAttributeInfoList")
	})
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

var FileAttributeMatcherStruct *gi.Struct
var FileAttributeMatcherStructOnce sync.Once

func FileAttributeMatcherStructSet() {
	FileAttributeMatcherStructOnce.Do(func() {
		FileAttributeMatcherStruct = gi.StructNew("Gio", "FileAttributeMatcher")
	})
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

var FileDescriptorBasedIfaceStruct *gi.Struct
var FileDescriptorBasedIfaceStructOnce sync.Once

func FileDescriptorBasedIfaceStructSet() {
	FileDescriptorBasedIfaceStructOnce.Do(func() {
		FileDescriptorBasedIfaceStruct = gi.StructNew("Gio", "FileDescriptorBasedIface")
	})
}

type FileDescriptorBasedIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_fd' : missing Type
}

var FileEnumeratorClassStruct *gi.Struct
var FileEnumeratorClassStructOnce sync.Once

func FileEnumeratorClassStructSet() {
	FileEnumeratorClassStructOnce.Do(func() {
		FileEnumeratorClassStruct = gi.StructNew("Gio", "FileEnumeratorClass")
	})
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

var FileEnumeratorPrivateStruct *gi.Struct
var FileEnumeratorPrivateStructOnce sync.Once

func FileEnumeratorPrivateStructSet() {
	FileEnumeratorPrivateStructOnce.Do(func() {
		FileEnumeratorPrivateStruct = gi.StructNew("Gio", "FileEnumeratorPrivate")
	})
}

type FileEnumeratorPrivate struct {
	native uintptr
}

var FileIOStreamClassStruct *gi.Struct
var FileIOStreamClassStructOnce sync.Once

func FileIOStreamClassStructSet() {
	FileIOStreamClassStructOnce.Do(func() {
		FileIOStreamClassStruct = gi.StructNew("Gio", "FileIOStreamClass")
	})
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

var FileIOStreamPrivateStruct *gi.Struct
var FileIOStreamPrivateStructOnce sync.Once

func FileIOStreamPrivateStructSet() {
	FileIOStreamPrivateStructOnce.Do(func() {
		FileIOStreamPrivateStruct = gi.StructNew("Gio", "FileIOStreamPrivate")
	})
}

type FileIOStreamPrivate struct {
	native uintptr
}

var FileIconClassStruct *gi.Struct
var FileIconClassStructOnce sync.Once

func FileIconClassStructSet() {
	FileIconClassStructOnce.Do(func() {
		FileIconClassStruct = gi.StructNew("Gio", "FileIconClass")
	})
}

type FileIconClass struct {
	native uintptr
}

var FileIfaceStruct *gi.Struct
var FileIfaceStructOnce sync.Once

func FileIfaceStructSet() {
	FileIfaceStructOnce.Do(func() {
		FileIfaceStruct = gi.StructNew("Gio", "FileIface")
	})
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

var FileInfoClassStruct *gi.Struct
var FileInfoClassStructOnce sync.Once

func FileInfoClassStructSet() {
	FileInfoClassStructOnce.Do(func() {
		FileInfoClassStruct = gi.StructNew("Gio", "FileInfoClass")
	})
}

type FileInfoClass struct {
	native uintptr
}

var FileInputStreamClassStruct *gi.Struct
var FileInputStreamClassStructOnce sync.Once

func FileInputStreamClassStructSet() {
	FileInputStreamClassStructOnce.Do(func() {
		FileInputStreamClassStruct = gi.StructNew("Gio", "FileInputStreamClass")
	})
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

var FileInputStreamPrivateStruct *gi.Struct
var FileInputStreamPrivateStructOnce sync.Once

func FileInputStreamPrivateStructSet() {
	FileInputStreamPrivateStructOnce.Do(func() {
		FileInputStreamPrivateStruct = gi.StructNew("Gio", "FileInputStreamPrivate")
	})
}

type FileInputStreamPrivate struct {
	native uintptr
}

var FileMonitorClassStruct *gi.Struct
var FileMonitorClassStructOnce sync.Once

func FileMonitorClassStructSet() {
	FileMonitorClassStructOnce.Do(func() {
		FileMonitorClassStruct = gi.StructNew("Gio", "FileMonitorClass")
	})
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

var FileMonitorPrivateStruct *gi.Struct
var FileMonitorPrivateStructOnce sync.Once

func FileMonitorPrivateStructSet() {
	FileMonitorPrivateStructOnce.Do(func() {
		FileMonitorPrivateStruct = gi.StructNew("Gio", "FileMonitorPrivate")
	})
}

type FileMonitorPrivate struct {
	native uintptr
}

var FileOutputStreamClassStruct *gi.Struct
var FileOutputStreamClassStructOnce sync.Once

func FileOutputStreamClassStructSet() {
	FileOutputStreamClassStructOnce.Do(func() {
		FileOutputStreamClassStruct = gi.StructNew("Gio", "FileOutputStreamClass")
	})
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

var FileOutputStreamPrivateStruct *gi.Struct
var FileOutputStreamPrivateStructOnce sync.Once

func FileOutputStreamPrivateStructSet() {
	FileOutputStreamPrivateStructOnce.Do(func() {
		FileOutputStreamPrivateStruct = gi.StructNew("Gio", "FileOutputStreamPrivate")
	})
}

type FileOutputStreamPrivate struct {
	native uintptr
}

var FilenameCompleterClassStruct *gi.Struct
var FilenameCompleterClassStructOnce sync.Once

func FilenameCompleterClassStructSet() {
	FilenameCompleterClassStructOnce.Do(func() {
		FilenameCompleterClassStruct = gi.StructNew("Gio", "FilenameCompleterClass")
	})
}

type FilenameCompleterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'got_completion_data' : missing Type
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
	// UNSUPPORTED : C value '_g_reserved3' : missing Type
}

var FilterInputStreamClassStruct *gi.Struct
var FilterInputStreamClassStructOnce sync.Once

func FilterInputStreamClassStructSet() {
	FilterInputStreamClassStructOnce.Do(func() {
		FilterInputStreamClassStruct = gi.StructNew("Gio", "FilterInputStreamClass")
	})
}

type FilterInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
	// UNSUPPORTED : C value '_g_reserved3' : missing Type
}

var FilterOutputStreamClassStruct *gi.Struct
var FilterOutputStreamClassStructOnce sync.Once

func FilterOutputStreamClassStructSet() {
	FilterOutputStreamClassStructOnce.Do(func() {
		FilterOutputStreamClassStruct = gi.StructNew("Gio", "FilterOutputStreamClass")
	})
}

type FilterOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
	// UNSUPPORTED : C value '_g_reserved3' : missing Type
}

var IOExtensionStruct *gi.Struct
var IOExtensionStructOnce sync.Once

func IOExtensionStructSet() {
	IOExtensionStructOnce.Do(func() {
		IOExtensionStruct = gi.StructNew("Gio", "IOExtension")
	})
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

var IOExtensionPointStruct *gi.Struct
var IOExtensionPointStructOnce sync.Once

func IOExtensionPointStructSet() {
	IOExtensionPointStructOnce.Do(func() {
		IOExtensionPointStruct = gi.StructNew("Gio", "IOExtensionPoint")
	})
}

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

var IOModuleClassStruct *gi.Struct
var IOModuleClassStructOnce sync.Once

func IOModuleClassStructSet() {
	IOModuleClassStructOnce.Do(func() {
		IOModuleClassStruct = gi.StructNew("Gio", "IOModuleClass")
	})
}

type IOModuleClass struct {
	native uintptr
}

var IOModuleScopeStruct *gi.Struct
var IOModuleScopeStructOnce sync.Once

func IOModuleScopeStructSet() {
	IOModuleScopeStructOnce.Do(func() {
		IOModuleScopeStruct = gi.StructNew("Gio", "IOModuleScope")
	})
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

var IOSchedulerJobStruct *gi.Struct
var IOSchedulerJobStructOnce sync.Once

func IOSchedulerJobStructSet() {
	IOSchedulerJobStructOnce.Do(func() {
		IOSchedulerJobStruct = gi.StructNew("Gio", "IOSchedulerJob")
	})
}

type IOSchedulerJob struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

var IOStreamAdapterStruct *gi.Struct
var IOStreamAdapterStructOnce sync.Once

func IOStreamAdapterStructSet() {
	IOStreamAdapterStructOnce.Do(func() {
		IOStreamAdapterStruct = gi.StructNew("Gio", "IOStreamAdapter")
	})
}

type IOStreamAdapter struct {
	native uintptr
}

var IOStreamClassStruct *gi.Struct
var IOStreamClassStructOnce sync.Once

func IOStreamClassStructSet() {
	IOStreamClassStructOnce.Do(func() {
		IOStreamClassStruct = gi.StructNew("Gio", "IOStreamClass")
	})
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

var IOStreamPrivateStruct *gi.Struct
var IOStreamPrivateStructOnce sync.Once

func IOStreamPrivateStructSet() {
	IOStreamPrivateStructOnce.Do(func() {
		IOStreamPrivateStruct = gi.StructNew("Gio", "IOStreamPrivate")
	})
}

type IOStreamPrivate struct {
	native uintptr
}

var IconIfaceStruct *gi.Struct
var IconIfaceStructOnce sync.Once

func IconIfaceStructSet() {
	IconIfaceStructOnce.Do(func() {
		IconIfaceStruct = gi.StructNew("Gio", "IconIface")
	})
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

var InetAddressClassStruct *gi.Struct
var InetAddressClassStructOnce sync.Once

func InetAddressClassStructSet() {
	InetAddressClassStructOnce.Do(func() {
		InetAddressClassStruct = gi.StructNew("Gio", "InetAddressClass")
	})
}

type InetAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'to_string' : missing Type
	// UNSUPPORTED : C value 'to_bytes' : missing Type
}

var InetAddressMaskClassStruct *gi.Struct
var InetAddressMaskClassStructOnce sync.Once

func InetAddressMaskClassStructSet() {
	InetAddressMaskClassStructOnce.Do(func() {
		InetAddressMaskClassStruct = gi.StructNew("Gio", "InetAddressMaskClass")
	})
}

type InetAddressMaskClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var InetAddressMaskPrivateStruct *gi.Struct
var InetAddressMaskPrivateStructOnce sync.Once

func InetAddressMaskPrivateStructSet() {
	InetAddressMaskPrivateStructOnce.Do(func() {
		InetAddressMaskPrivateStruct = gi.StructNew("Gio", "InetAddressMaskPrivate")
	})
}

type InetAddressMaskPrivate struct {
	native uintptr
}

var InetAddressPrivateStruct *gi.Struct
var InetAddressPrivateStructOnce sync.Once

func InetAddressPrivateStructSet() {
	InetAddressPrivateStructOnce.Do(func() {
		InetAddressPrivateStruct = gi.StructNew("Gio", "InetAddressPrivate")
	})
}

type InetAddressPrivate struct {
	native uintptr
}

var InetSocketAddressClassStruct *gi.Struct
var InetSocketAddressClassStructOnce sync.Once

func InetSocketAddressClassStructSet() {
	InetSocketAddressClassStructOnce.Do(func() {
		InetSocketAddressClassStruct = gi.StructNew("Gio", "InetSocketAddressClass")
	})
}

type InetSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var InetSocketAddressPrivateStruct *gi.Struct
var InetSocketAddressPrivateStructOnce sync.Once

func InetSocketAddressPrivateStructSet() {
	InetSocketAddressPrivateStructOnce.Do(func() {
		InetSocketAddressPrivateStruct = gi.StructNew("Gio", "InetSocketAddressPrivate")
	})
}

type InetSocketAddressPrivate struct {
	native uintptr
}

var InitableIfaceStruct *gi.Struct
var InitableIfaceStructOnce sync.Once

func InitableIfaceStructSet() {
	InitableIfaceStructOnce.Do(func() {
		InitableIfaceStruct = gi.StructNew("Gio", "InitableIface")
	})
}

type InitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'init' : missing Type
}

var InputMessageStruct *gi.Struct
var InputMessageStructOnce sync.Once

func InputMessageStructSet() {
	InputMessageStructOnce.Do(func() {
		InputMessageStruct = gi.StructNew("Gio", "InputMessage")
	})
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

var InputStreamClassStruct *gi.Struct
var InputStreamClassStructOnce sync.Once

func InputStreamClassStructSet() {
	InputStreamClassStructOnce.Do(func() {
		InputStreamClassStruct = gi.StructNew("Gio", "InputStreamClass")
	})
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

var InputStreamPrivateStruct *gi.Struct
var InputStreamPrivateStructOnce sync.Once

func InputStreamPrivateStructSet() {
	InputStreamPrivateStructOnce.Do(func() {
		InputStreamPrivateStruct = gi.StructNew("Gio", "InputStreamPrivate")
	})
}

type InputStreamPrivate struct {
	native uintptr
}

var InputVectorStruct *gi.Struct
var InputVectorStructOnce sync.Once

func InputVectorStructSet() {
	InputVectorStructOnce.Do(func() {
		InputVectorStruct = gi.StructNew("Gio", "InputVector")
	})
}

type InputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var ListModelInterfaceStruct *gi.Struct
var ListModelInterfaceStructOnce sync.Once

func ListModelInterfaceStructSet() {
	ListModelInterfaceStructOnce.Do(func() {
		ListModelInterfaceStruct = gi.StructNew("Gio", "ListModelInterface")
	})
}

type ListModelInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_item_type' : missing Type
	// UNSUPPORTED : C value 'get_n_items' : missing Type
	// UNSUPPORTED : C value 'get_item' : missing Type
}

var ListStoreClassStruct *gi.Struct
var ListStoreClassStructOnce sync.Once

func ListStoreClassStructSet() {
	ListStoreClassStructOnce.Do(func() {
		ListStoreClassStruct = gi.StructNew("Gio", "ListStoreClass")
	})
}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var LoadableIconIfaceStruct *gi.Struct
var LoadableIconIfaceStructOnce sync.Once

func LoadableIconIfaceStructSet() {
	LoadableIconIfaceStructOnce.Do(func() {
		LoadableIconIfaceStruct = gi.StructNew("Gio", "LoadableIconIface")
	})
}

type LoadableIconIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'load' : missing Type
	// UNSUPPORTED : C value 'load_async' : missing Type
	// UNSUPPORTED : C value 'load_finish' : missing Type
}

var MemoryInputStreamClassStruct *gi.Struct
var MemoryInputStreamClassStructOnce sync.Once

func MemoryInputStreamClassStructSet() {
	MemoryInputStreamClassStructOnce.Do(func() {
		MemoryInputStreamClassStruct = gi.StructNew("Gio", "MemoryInputStreamClass")
	})
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

var MemoryInputStreamPrivateStruct *gi.Struct
var MemoryInputStreamPrivateStructOnce sync.Once

func MemoryInputStreamPrivateStructSet() {
	MemoryInputStreamPrivateStructOnce.Do(func() {
		MemoryInputStreamPrivateStruct = gi.StructNew("Gio", "MemoryInputStreamPrivate")
	})
}

type MemoryInputStreamPrivate struct {
	native uintptr
}

var MemoryOutputStreamClassStruct *gi.Struct
var MemoryOutputStreamClassStructOnce sync.Once

func MemoryOutputStreamClassStructSet() {
	MemoryOutputStreamClassStructOnce.Do(func() {
		MemoryOutputStreamClassStruct = gi.StructNew("Gio", "MemoryOutputStreamClass")
	})
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

var MemoryOutputStreamPrivateStruct *gi.Struct
var MemoryOutputStreamPrivateStructOnce sync.Once

func MemoryOutputStreamPrivateStructSet() {
	MemoryOutputStreamPrivateStructOnce.Do(func() {
		MemoryOutputStreamPrivateStruct = gi.StructNew("Gio", "MemoryOutputStreamPrivate")
	})
}

type MemoryOutputStreamPrivate struct {
	native uintptr
}

var MenuAttributeIterClassStruct *gi.Struct
var MenuAttributeIterClassStructOnce sync.Once

func MenuAttributeIterClassStructSet() {
	MenuAttributeIterClassStructOnce.Do(func() {
		MenuAttributeIterClassStruct = gi.StructNew("Gio", "MenuAttributeIterClass")
	})
}

type MenuAttributeIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var MenuAttributeIterPrivateStruct *gi.Struct
var MenuAttributeIterPrivateStructOnce sync.Once

func MenuAttributeIterPrivateStructSet() {
	MenuAttributeIterPrivateStructOnce.Do(func() {
		MenuAttributeIterPrivateStruct = gi.StructNew("Gio", "MenuAttributeIterPrivate")
	})
}

type MenuAttributeIterPrivate struct {
	native uintptr
}

var MenuLinkIterClassStruct *gi.Struct
var MenuLinkIterClassStructOnce sync.Once

func MenuLinkIterClassStructSet() {
	MenuLinkIterClassStructOnce.Do(func() {
		MenuLinkIterClassStruct = gi.StructNew("Gio", "MenuLinkIterClass")
	})
}

type MenuLinkIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var MenuLinkIterPrivateStruct *gi.Struct
var MenuLinkIterPrivateStructOnce sync.Once

func MenuLinkIterPrivateStructSet() {
	MenuLinkIterPrivateStructOnce.Do(func() {
		MenuLinkIterPrivateStruct = gi.StructNew("Gio", "MenuLinkIterPrivate")
	})
}

type MenuLinkIterPrivate struct {
	native uintptr
}

var MenuModelClassStruct *gi.Struct
var MenuModelClassStructOnce sync.Once

func MenuModelClassStructSet() {
	MenuModelClassStructOnce.Do(func() {
		MenuModelClassStruct = gi.StructNew("Gio", "MenuModelClass")
	})
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

var MenuModelPrivateStruct *gi.Struct
var MenuModelPrivateStructOnce sync.Once

func MenuModelPrivateStructSet() {
	MenuModelPrivateStructOnce.Do(func() {
		MenuModelPrivateStruct = gi.StructNew("Gio", "MenuModelPrivate")
	})
}

type MenuModelPrivate struct {
	native uintptr
}

var MountIfaceStruct *gi.Struct
var MountIfaceStructOnce sync.Once

func MountIfaceStructSet() {
	MountIfaceStructOnce.Do(func() {
		MountIfaceStruct = gi.StructNew("Gio", "MountIface")
	})
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

var MountOperationClassStruct *gi.Struct
var MountOperationClassStructOnce sync.Once

func MountOperationClassStructSet() {
	MountOperationClassStructOnce.Do(func() {
		MountOperationClassStruct = gi.StructNew("Gio", "MountOperationClass")
	})
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

var MountOperationPrivateStruct *gi.Struct
var MountOperationPrivateStructOnce sync.Once

func MountOperationPrivateStructSet() {
	MountOperationPrivateStructOnce.Do(func() {
		MountOperationPrivateStruct = gi.StructNew("Gio", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var NativeSocketAddressClassStruct *gi.Struct
var NativeSocketAddressClassStructOnce sync.Once

func NativeSocketAddressClassStructSet() {
	NativeSocketAddressClassStructOnce.Do(func() {
		NativeSocketAddressClassStruct = gi.StructNew("Gio", "NativeSocketAddressClass")
	})
}

type NativeSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var NativeSocketAddressPrivateStruct *gi.Struct
var NativeSocketAddressPrivateStructOnce sync.Once

func NativeSocketAddressPrivateStructSet() {
	NativeSocketAddressPrivateStructOnce.Do(func() {
		NativeSocketAddressPrivateStruct = gi.StructNew("Gio", "NativeSocketAddressPrivate")
	})
}

type NativeSocketAddressPrivate struct {
	native uintptr
}

var NativeVolumeMonitorClassStruct *gi.Struct
var NativeVolumeMonitorClassStructOnce sync.Once

func NativeVolumeMonitorClassStructSet() {
	NativeVolumeMonitorClassStructOnce.Do(func() {
		NativeVolumeMonitorClassStruct = gi.StructNew("Gio", "NativeVolumeMonitorClass")
	})
}

type NativeVolumeMonitorClass struct {
	native      uintptr
	ParentClass *VolumeMonitorClass
	// UNSUPPORTED : C value 'get_mount_for_mount_path' : missing Type
}

var NetworkAddressClassStruct *gi.Struct
var NetworkAddressClassStructOnce sync.Once

func NetworkAddressClassStructSet() {
	NetworkAddressClassStructOnce.Do(func() {
		NetworkAddressClassStruct = gi.StructNew("Gio", "NetworkAddressClass")
	})
}

type NetworkAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var NetworkAddressPrivateStruct *gi.Struct
var NetworkAddressPrivateStructOnce sync.Once

func NetworkAddressPrivateStructSet() {
	NetworkAddressPrivateStructOnce.Do(func() {
		NetworkAddressPrivateStruct = gi.StructNew("Gio", "NetworkAddressPrivate")
	})
}

type NetworkAddressPrivate struct {
	native uintptr
}

var NetworkMonitorInterfaceStruct *gi.Struct
var NetworkMonitorInterfaceStructOnce sync.Once

func NetworkMonitorInterfaceStructSet() {
	NetworkMonitorInterfaceStructOnce.Do(func() {
		NetworkMonitorInterfaceStruct = gi.StructNew("Gio", "NetworkMonitorInterface")
	})
}

type NetworkMonitorInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'network_changed' : missing Type
	// UNSUPPORTED : C value 'can_reach' : missing Type
	// UNSUPPORTED : C value 'can_reach_async' : missing Type
	// UNSUPPORTED : C value 'can_reach_finish' : missing Type
}

var NetworkServiceClassStruct *gi.Struct
var NetworkServiceClassStructOnce sync.Once

func NetworkServiceClassStructSet() {
	NetworkServiceClassStructOnce.Do(func() {
		NetworkServiceClassStruct = gi.StructNew("Gio", "NetworkServiceClass")
	})
}

type NetworkServiceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var NetworkServicePrivateStruct *gi.Struct
var NetworkServicePrivateStructOnce sync.Once

func NetworkServicePrivateStructSet() {
	NetworkServicePrivateStructOnce.Do(func() {
		NetworkServicePrivateStruct = gi.StructNew("Gio", "NetworkServicePrivate")
	})
}

type NetworkServicePrivate struct {
	native uintptr
}

var OutputMessageStruct *gi.Struct
var OutputMessageStructOnce sync.Once

func OutputMessageStructSet() {
	OutputMessageStructOnce.Do(func() {
		OutputMessageStruct = gi.StructNew("Gio", "OutputMessage")
	})
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

var OutputStreamClassStruct *gi.Struct
var OutputStreamClassStructOnce sync.Once

func OutputStreamClassStructSet() {
	OutputStreamClassStructOnce.Do(func() {
		OutputStreamClassStruct = gi.StructNew("Gio", "OutputStreamClass")
	})
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

var OutputStreamPrivateStruct *gi.Struct
var OutputStreamPrivateStructOnce sync.Once

func OutputStreamPrivateStructSet() {
	OutputStreamPrivateStructOnce.Do(func() {
		OutputStreamPrivateStruct = gi.StructNew("Gio", "OutputStreamPrivate")
	})
}

type OutputStreamPrivate struct {
	native uintptr
}

var OutputVectorStruct *gi.Struct
var OutputVectorStructOnce sync.Once

func OutputVectorStructSet() {
	OutputVectorStructOnce.Do(func() {
		OutputVectorStruct = gi.StructNew("Gio", "OutputVector")
	})
}

type OutputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var PermissionClassStruct *gi.Struct
var PermissionClassStructOnce sync.Once

func PermissionClassStructSet() {
	PermissionClassStructOnce.Do(func() {
		PermissionClassStruct = gi.StructNew("Gio", "PermissionClass")
	})
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

var PermissionPrivateStruct *gi.Struct
var PermissionPrivateStructOnce sync.Once

func PermissionPrivateStructSet() {
	PermissionPrivateStructOnce.Do(func() {
		PermissionPrivateStruct = gi.StructNew("Gio", "PermissionPrivate")
	})
}

type PermissionPrivate struct {
	native uintptr
}

var PollableInputStreamInterfaceStruct *gi.Struct
var PollableInputStreamInterfaceStructOnce sync.Once

func PollableInputStreamInterfaceStructSet() {
	PollableInputStreamInterfaceStructOnce.Do(func() {
		PollableInputStreamInterfaceStruct = gi.StructNew("Gio", "PollableInputStreamInterface")
	})
}

type PollableInputStreamInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'can_poll' : missing Type
	// UNSUPPORTED : C value 'is_readable' : missing Type
	// UNSUPPORTED : C value 'create_source' : missing Type
	// UNSUPPORTED : C value 'read_nonblocking' : missing Type
}

var PollableOutputStreamInterfaceStruct *gi.Struct
var PollableOutputStreamInterfaceStructOnce sync.Once

func PollableOutputStreamInterfaceStructSet() {
	PollableOutputStreamInterfaceStructOnce.Do(func() {
		PollableOutputStreamInterfaceStruct = gi.StructNew("Gio", "PollableOutputStreamInterface")
	})
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

var ProxyAddressClassStruct *gi.Struct
var ProxyAddressClassStructOnce sync.Once

func ProxyAddressClassStructSet() {
	ProxyAddressClassStructOnce.Do(func() {
		ProxyAddressClassStruct = gi.StructNew("Gio", "ProxyAddressClass")
	})
}

type ProxyAddressClass struct {
	native      uintptr
	ParentClass *InetSocketAddressClass
}

var ProxyAddressEnumeratorClassStruct *gi.Struct
var ProxyAddressEnumeratorClassStructOnce sync.Once

func ProxyAddressEnumeratorClassStructSet() {
	ProxyAddressEnumeratorClassStructOnce.Do(func() {
		ProxyAddressEnumeratorClassStruct = gi.StructNew("Gio", "ProxyAddressEnumeratorClass")
	})
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

var ProxyAddressEnumeratorPrivateStruct *gi.Struct
var ProxyAddressEnumeratorPrivateStructOnce sync.Once

func ProxyAddressEnumeratorPrivateStructSet() {
	ProxyAddressEnumeratorPrivateStructOnce.Do(func() {
		ProxyAddressEnumeratorPrivateStruct = gi.StructNew("Gio", "ProxyAddressEnumeratorPrivate")
	})
}

type ProxyAddressEnumeratorPrivate struct {
	native uintptr
}

var ProxyAddressPrivateStruct *gi.Struct
var ProxyAddressPrivateStructOnce sync.Once

func ProxyAddressPrivateStructSet() {
	ProxyAddressPrivateStructOnce.Do(func() {
		ProxyAddressPrivateStruct = gi.StructNew("Gio", "ProxyAddressPrivate")
	})
}

type ProxyAddressPrivate struct {
	native uintptr
}

var ProxyInterfaceStruct *gi.Struct
var ProxyInterfaceStructOnce sync.Once

func ProxyInterfaceStructSet() {
	ProxyInterfaceStructOnce.Do(func() {
		ProxyInterfaceStruct = gi.StructNew("Gio", "ProxyInterface")
	})
}

type ProxyInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'connect' : missing Type
	// UNSUPPORTED : C value 'connect_async' : missing Type
	// UNSUPPORTED : C value 'connect_finish' : missing Type
	// UNSUPPORTED : C value 'supports_hostname' : missing Type
}

var ProxyResolverInterfaceStruct *gi.Struct
var ProxyResolverInterfaceStructOnce sync.Once

func ProxyResolverInterfaceStructSet() {
	ProxyResolverInterfaceStructOnce.Do(func() {
		ProxyResolverInterfaceStruct = gi.StructNew("Gio", "ProxyResolverInterface")
	})
}

type ProxyResolverInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'is_supported' : missing Type
	// UNSUPPORTED : C value 'lookup' : missing Type
	// UNSUPPORTED : C value 'lookup_async' : missing Type
	// UNSUPPORTED : C value 'lookup_finish' : missing Type
}

var RemoteActionGroupInterfaceStruct *gi.Struct
var RemoteActionGroupInterfaceStructOnce sync.Once

func RemoteActionGroupInterfaceStructSet() {
	RemoteActionGroupInterfaceStructOnce.Do(func() {
		RemoteActionGroupInterfaceStruct = gi.StructNew("Gio", "RemoteActionGroupInterface")
	})
}

type RemoteActionGroupInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'activate_action_full' : missing Type
	// UNSUPPORTED : C value 'change_action_state_full' : missing Type
}

var ResolverClassStruct *gi.Struct
var ResolverClassStructOnce sync.Once

func ResolverClassStructSet() {
	ResolverClassStructOnce.Do(func() {
		ResolverClassStruct = gi.StructNew("Gio", "ResolverClass")
	})
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

var ResolverPrivateStruct *gi.Struct
var ResolverPrivateStructOnce sync.Once

func ResolverPrivateStructSet() {
	ResolverPrivateStructOnce.Do(func() {
		ResolverPrivateStruct = gi.StructNew("Gio", "ResolverPrivate")
	})
}

type ResolverPrivate struct {
	native uintptr
}

var ResourceStruct *gi.Struct
var ResourceStructOnce sync.Once

func ResourceStructSet() {
	ResourceStructOnce.Do(func() {
		ResourceStruct = gi.StructNew("Gio", "Resource")
	})
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

var SeekableIfaceStruct *gi.Struct
var SeekableIfaceStructOnce sync.Once

func SeekableIfaceStructSet() {
	SeekableIfaceStructOnce.Do(func() {
		SeekableIfaceStruct = gi.StructNew("Gio", "SeekableIface")
	})
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

var SettingsBackendClassStruct *gi.Struct
var SettingsBackendClassStructOnce sync.Once

func SettingsBackendClassStructSet() {
	SettingsBackendClassStructOnce.Do(func() {
		SettingsBackendClassStruct = gi.StructNew("Gio", "SettingsBackendClass")
	})
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

var SettingsBackendPrivateStruct *gi.Struct
var SettingsBackendPrivateStructOnce sync.Once

func SettingsBackendPrivateStructSet() {
	SettingsBackendPrivateStructOnce.Do(func() {
		SettingsBackendPrivateStruct = gi.StructNew("Gio", "SettingsBackendPrivate")
	})
}

type SettingsBackendPrivate struct {
	native uintptr
}

var SettingsClassStruct *gi.Struct
var SettingsClassStructOnce sync.Once

func SettingsClassStructSet() {
	SettingsClassStructOnce.Do(func() {
		SettingsClassStruct = gi.StructNew("Gio", "SettingsClass")
	})
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

var SettingsPrivateStruct *gi.Struct
var SettingsPrivateStructOnce sync.Once

func SettingsPrivateStructSet() {
	SettingsPrivateStructOnce.Do(func() {
		SettingsPrivateStruct = gi.StructNew("Gio", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var SettingsSchemaStruct *gi.Struct
var SettingsSchemaStructOnce sync.Once

func SettingsSchemaStructSet() {
	SettingsSchemaStructOnce.Do(func() {
		SettingsSchemaStruct = gi.StructNew("Gio", "SettingsSchema")
	})
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

var SettingsSchemaKeyStruct *gi.Struct
var SettingsSchemaKeyStructOnce sync.Once

func SettingsSchemaKeyStructSet() {
	SettingsSchemaKeyStructOnce.Do(func() {
		SettingsSchemaKeyStruct = gi.StructNew("Gio", "SettingsSchemaKey")
	})
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

var SettingsSchemaSourceStruct *gi.Struct
var SettingsSchemaSourceStructOnce sync.Once

func SettingsSchemaSourceStructSet() {
	SettingsSchemaSourceStructOnce.Do(func() {
		SettingsSchemaSourceStruct = gi.StructNew("Gio", "SettingsSchemaSource")
	})
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

var SimpleActionGroupClassStruct *gi.Struct
var SimpleActionGroupClassStructOnce sync.Once

func SimpleActionGroupClassStructSet() {
	SimpleActionGroupClassStructOnce.Do(func() {
		SimpleActionGroupClassStruct = gi.StructNew("Gio", "SimpleActionGroupClass")
	})
}

type SimpleActionGroupClass struct {
	native uintptr
}

var SimpleActionGroupPrivateStruct *gi.Struct
var SimpleActionGroupPrivateStructOnce sync.Once

func SimpleActionGroupPrivateStructSet() {
	SimpleActionGroupPrivateStructOnce.Do(func() {
		SimpleActionGroupPrivateStruct = gi.StructNew("Gio", "SimpleActionGroupPrivate")
	})
}

type SimpleActionGroupPrivate struct {
	native uintptr
}

var SimpleAsyncResultClassStruct *gi.Struct
var SimpleAsyncResultClassStructOnce sync.Once

func SimpleAsyncResultClassStructSet() {
	SimpleAsyncResultClassStructOnce.Do(func() {
		SimpleAsyncResultClassStruct = gi.StructNew("Gio", "SimpleAsyncResultClass")
	})
}

type SimpleAsyncResultClass struct {
	native uintptr
}

var SimpleProxyResolverClassStruct *gi.Struct
var SimpleProxyResolverClassStructOnce sync.Once

func SimpleProxyResolverClassStructSet() {
	SimpleProxyResolverClassStructOnce.Do(func() {
		SimpleProxyResolverClassStruct = gi.StructNew("Gio", "SimpleProxyResolverClass")
	})
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

var SimpleProxyResolverPrivateStruct *gi.Struct
var SimpleProxyResolverPrivateStructOnce sync.Once

func SimpleProxyResolverPrivateStructSet() {
	SimpleProxyResolverPrivateStructOnce.Do(func() {
		SimpleProxyResolverPrivateStruct = gi.StructNew("Gio", "SimpleProxyResolverPrivate")
	})
}

type SimpleProxyResolverPrivate struct {
	native uintptr
}

var SocketAddressClassStruct *gi.Struct
var SocketAddressClassStructOnce sync.Once

func SocketAddressClassStructSet() {
	SocketAddressClassStructOnce.Do(func() {
		SocketAddressClassStruct = gi.StructNew("Gio", "SocketAddressClass")
	})
}

type SocketAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_family' : missing Type
	// UNSUPPORTED : C value 'get_native_size' : missing Type
	// UNSUPPORTED : C value 'to_native' : missing Type
}

var SocketAddressEnumeratorClassStruct *gi.Struct
var SocketAddressEnumeratorClassStructOnce sync.Once

func SocketAddressEnumeratorClassStructSet() {
	SocketAddressEnumeratorClassStructOnce.Do(func() {
		SocketAddressEnumeratorClassStruct = gi.StructNew("Gio", "SocketAddressEnumeratorClass")
	})
}

type SocketAddressEnumeratorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'next' : missing Type
	// UNSUPPORTED : C value 'next_async' : missing Type
	// UNSUPPORTED : C value 'next_finish' : missing Type
}

var SocketClassStruct *gi.Struct
var SocketClassStructOnce sync.Once

func SocketClassStructSet() {
	SocketClassStructOnce.Do(func() {
		SocketClassStruct = gi.StructNew("Gio", "SocketClass")
	})
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

var SocketClientClassStruct *gi.Struct
var SocketClientClassStructOnce sync.Once

func SocketClientClassStructSet() {
	SocketClientClassStructOnce.Do(func() {
		SocketClientClassStruct = gi.StructNew("Gio", "SocketClientClass")
	})
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

var SocketClientPrivateStruct *gi.Struct
var SocketClientPrivateStructOnce sync.Once

func SocketClientPrivateStructSet() {
	SocketClientPrivateStructOnce.Do(func() {
		SocketClientPrivateStruct = gi.StructNew("Gio", "SocketClientPrivate")
	})
}

type SocketClientPrivate struct {
	native uintptr
}

var SocketConnectableIfaceStruct *gi.Struct
var SocketConnectableIfaceStructOnce sync.Once

func SocketConnectableIfaceStructSet() {
	SocketConnectableIfaceStructOnce.Do(func() {
		SocketConnectableIfaceStruct = gi.StructNew("Gio", "SocketConnectableIface")
	})
}

type SocketConnectableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'enumerate' : missing Type
	// UNSUPPORTED : C value 'proxy_enumerate' : missing Type
	// UNSUPPORTED : C value 'to_string' : missing Type
}

var SocketConnectionClassStruct *gi.Struct
var SocketConnectionClassStructOnce sync.Once

func SocketConnectionClassStructSet() {
	SocketConnectionClassStructOnce.Do(func() {
		SocketConnectionClassStruct = gi.StructNew("Gio", "SocketConnectionClass")
	})
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

var SocketConnectionPrivateStruct *gi.Struct
var SocketConnectionPrivateStructOnce sync.Once

func SocketConnectionPrivateStructSet() {
	SocketConnectionPrivateStructOnce.Do(func() {
		SocketConnectionPrivateStruct = gi.StructNew("Gio", "SocketConnectionPrivate")
	})
}

type SocketConnectionPrivate struct {
	native uintptr
}

var SocketControlMessageClassStruct *gi.Struct
var SocketControlMessageClassStructOnce sync.Once

func SocketControlMessageClassStructSet() {
	SocketControlMessageClassStructOnce.Do(func() {
		SocketControlMessageClassStruct = gi.StructNew("Gio", "SocketControlMessageClass")
	})
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

var SocketControlMessagePrivateStruct *gi.Struct
var SocketControlMessagePrivateStructOnce sync.Once

func SocketControlMessagePrivateStructSet() {
	SocketControlMessagePrivateStructOnce.Do(func() {
		SocketControlMessagePrivateStruct = gi.StructNew("Gio", "SocketControlMessagePrivate")
	})
}

type SocketControlMessagePrivate struct {
	native uintptr
}

var SocketListenerClassStruct *gi.Struct
var SocketListenerClassStructOnce sync.Once

func SocketListenerClassStructSet() {
	SocketListenerClassStructOnce.Do(func() {
		SocketListenerClassStruct = gi.StructNew("Gio", "SocketListenerClass")
	})
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

var SocketListenerPrivateStruct *gi.Struct
var SocketListenerPrivateStructOnce sync.Once

func SocketListenerPrivateStructSet() {
	SocketListenerPrivateStructOnce.Do(func() {
		SocketListenerPrivateStruct = gi.StructNew("Gio", "SocketListenerPrivate")
	})
}

type SocketListenerPrivate struct {
	native uintptr
}

var SocketPrivateStruct *gi.Struct
var SocketPrivateStructOnce sync.Once

func SocketPrivateStructSet() {
	SocketPrivateStructOnce.Do(func() {
		SocketPrivateStruct = gi.StructNew("Gio", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var SocketServiceClassStruct *gi.Struct
var SocketServiceClassStructOnce sync.Once

func SocketServiceClassStructSet() {
	SocketServiceClassStructOnce.Do(func() {
		SocketServiceClassStruct = gi.StructNew("Gio", "SocketServiceClass")
	})
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

var SocketServicePrivateStruct *gi.Struct
var SocketServicePrivateStructOnce sync.Once

func SocketServicePrivateStructSet() {
	SocketServicePrivateStructOnce.Do(func() {
		SocketServicePrivateStruct = gi.StructNew("Gio", "SocketServicePrivate")
	})
}

type SocketServicePrivate struct {
	native uintptr
}

var SrvTargetStruct *gi.Struct
var SrvTargetStructOnce sync.Once

func SrvTargetStructSet() {
	SrvTargetStructOnce.Do(func() {
		SrvTargetStruct = gi.StructNew("Gio", "SrvTarget")
	})
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

var StaticResourceStruct *gi.Struct
var StaticResourceStructOnce sync.Once

func StaticResourceStructSet() {
	StaticResourceStructOnce.Do(func() {
		StaticResourceStruct = gi.StructNew("Gio", "StaticResource")
	})
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

var TaskClassStruct *gi.Struct
var TaskClassStructOnce sync.Once

func TaskClassStructSet() {
	TaskClassStructOnce.Do(func() {
		TaskClassStruct = gi.StructNew("Gio", "TaskClass")
	})
}

type TaskClass struct {
	native uintptr
}

var TcpConnectionClassStruct *gi.Struct
var TcpConnectionClassStructOnce sync.Once

func TcpConnectionClassStructSet() {
	TcpConnectionClassStructOnce.Do(func() {
		TcpConnectionClassStruct = gi.StructNew("Gio", "TcpConnectionClass")
	})
}

type TcpConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var TcpConnectionPrivateStruct *gi.Struct
var TcpConnectionPrivateStructOnce sync.Once

func TcpConnectionPrivateStructSet() {
	TcpConnectionPrivateStructOnce.Do(func() {
		TcpConnectionPrivateStruct = gi.StructNew("Gio", "TcpConnectionPrivate")
	})
}

type TcpConnectionPrivate struct {
	native uintptr
}

var TcpWrapperConnectionClassStruct *gi.Struct
var TcpWrapperConnectionClassStructOnce sync.Once

func TcpWrapperConnectionClassStructSet() {
	TcpWrapperConnectionClassStructOnce.Do(func() {
		TcpWrapperConnectionClassStruct = gi.StructNew("Gio", "TcpWrapperConnectionClass")
	})
}

type TcpWrapperConnectionClass struct {
	native      uintptr
	ParentClass *TcpConnectionClass
}

var TcpWrapperConnectionPrivateStruct *gi.Struct
var TcpWrapperConnectionPrivateStructOnce sync.Once

func TcpWrapperConnectionPrivateStructSet() {
	TcpWrapperConnectionPrivateStructOnce.Do(func() {
		TcpWrapperConnectionPrivateStruct = gi.StructNew("Gio", "TcpWrapperConnectionPrivate")
	})
}

type TcpWrapperConnectionPrivate struct {
	native uintptr
}

var ThemedIconClassStruct *gi.Struct
var ThemedIconClassStructOnce sync.Once

func ThemedIconClassStructSet() {
	ThemedIconClassStructOnce.Do(func() {
		ThemedIconClassStruct = gi.StructNew("Gio", "ThemedIconClass")
	})
}

type ThemedIconClass struct {
	native uintptr
}

var ThreadedSocketServiceClassStruct *gi.Struct
var ThreadedSocketServiceClassStructOnce sync.Once

func ThreadedSocketServiceClassStructSet() {
	ThreadedSocketServiceClassStructOnce.Do(func() {
		ThreadedSocketServiceClassStruct = gi.StructNew("Gio", "ThreadedSocketServiceClass")
	})
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

var ThreadedSocketServicePrivateStruct *gi.Struct
var ThreadedSocketServicePrivateStructOnce sync.Once

func ThreadedSocketServicePrivateStructSet() {
	ThreadedSocketServicePrivateStructOnce.Do(func() {
		ThreadedSocketServicePrivateStruct = gi.StructNew("Gio", "ThreadedSocketServicePrivate")
	})
}

type ThreadedSocketServicePrivate struct {
	native uintptr
}

var TlsBackendInterfaceStruct *gi.Struct
var TlsBackendInterfaceStructOnce sync.Once

func TlsBackendInterfaceStructSet() {
	TlsBackendInterfaceStructOnce.Do(func() {
		TlsBackendInterfaceStruct = gi.StructNew("Gio", "TlsBackendInterface")
	})
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

var TlsCertificateClassStruct *gi.Struct
var TlsCertificateClassStructOnce sync.Once

func TlsCertificateClassStructSet() {
	TlsCertificateClassStructOnce.Do(func() {
		TlsCertificateClassStruct = gi.StructNew("Gio", "TlsCertificateClass")
	})
}

type TlsCertificateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'verify' : missing Type
}

var TlsCertificatePrivateStruct *gi.Struct
var TlsCertificatePrivateStructOnce sync.Once

func TlsCertificatePrivateStructSet() {
	TlsCertificatePrivateStructOnce.Do(func() {
		TlsCertificatePrivateStruct = gi.StructNew("Gio", "TlsCertificatePrivate")
	})
}

type TlsCertificatePrivate struct {
	native uintptr
}

var TlsClientConnectionInterfaceStruct *gi.Struct
var TlsClientConnectionInterfaceStructOnce sync.Once

func TlsClientConnectionInterfaceStructSet() {
	TlsClientConnectionInterfaceStructOnce.Do(func() {
		TlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "TlsClientConnectionInterface")
	})
}

type TlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'copy_session_state' : missing Type
}

var TlsConnectionClassStruct *gi.Struct
var TlsConnectionClassStructOnce sync.Once

func TlsConnectionClassStructSet() {
	TlsConnectionClassStructOnce.Do(func() {
		TlsConnectionClassStruct = gi.StructNew("Gio", "TlsConnectionClass")
	})
}

type TlsConnectionClass struct {
	native      uintptr
	ParentClass *IOStreamClass
	// UNSUPPORTED : C value 'accept_certificate' : missing Type
	// UNSUPPORTED : C value 'handshake' : missing Type
	// UNSUPPORTED : C value 'handshake_async' : missing Type
	// UNSUPPORTED : C value 'handshake_finish' : missing Type
}

var TlsConnectionPrivateStruct *gi.Struct
var TlsConnectionPrivateStructOnce sync.Once

func TlsConnectionPrivateStructSet() {
	TlsConnectionPrivateStructOnce.Do(func() {
		TlsConnectionPrivateStruct = gi.StructNew("Gio", "TlsConnectionPrivate")
	})
}

type TlsConnectionPrivate struct {
	native uintptr
}

var TlsDatabaseClassStruct *gi.Struct
var TlsDatabaseClassStructOnce sync.Once

func TlsDatabaseClassStructSet() {
	TlsDatabaseClassStructOnce.Do(func() {
		TlsDatabaseClassStruct = gi.StructNew("Gio", "TlsDatabaseClass")
	})
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

var TlsDatabasePrivateStruct *gi.Struct
var TlsDatabasePrivateStructOnce sync.Once

func TlsDatabasePrivateStructSet() {
	TlsDatabasePrivateStructOnce.Do(func() {
		TlsDatabasePrivateStruct = gi.StructNew("Gio", "TlsDatabasePrivate")
	})
}

type TlsDatabasePrivate struct {
	native uintptr
}

var TlsFileDatabaseInterfaceStruct *gi.Struct
var TlsFileDatabaseInterfaceStructOnce sync.Once

func TlsFileDatabaseInterfaceStructSet() {
	TlsFileDatabaseInterfaceStructOnce.Do(func() {
		TlsFileDatabaseInterfaceStruct = gi.StructNew("Gio", "TlsFileDatabaseInterface")
	})
}

type TlsFileDatabaseInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var TlsInteractionClassStruct *gi.Struct
var TlsInteractionClassStructOnce sync.Once

func TlsInteractionClassStructSet() {
	TlsInteractionClassStructOnce.Do(func() {
		TlsInteractionClassStruct = gi.StructNew("Gio", "TlsInteractionClass")
	})
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

var TlsInteractionPrivateStruct *gi.Struct
var TlsInteractionPrivateStructOnce sync.Once

func TlsInteractionPrivateStructSet() {
	TlsInteractionPrivateStructOnce.Do(func() {
		TlsInteractionPrivateStruct = gi.StructNew("Gio", "TlsInteractionPrivate")
	})
}

type TlsInteractionPrivate struct {
	native uintptr
}

var TlsPasswordClassStruct *gi.Struct
var TlsPasswordClassStructOnce sync.Once

func TlsPasswordClassStructSet() {
	TlsPasswordClassStructOnce.Do(func() {
		TlsPasswordClassStruct = gi.StructNew("Gio", "TlsPasswordClass")
	})
}

type TlsPasswordClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_value' : missing Type
	// UNSUPPORTED : C value 'set_value' : missing Type
	// UNSUPPORTED : C value 'get_default_warning' : missing Type
}

var TlsPasswordPrivateStruct *gi.Struct
var TlsPasswordPrivateStructOnce sync.Once

func TlsPasswordPrivateStructSet() {
	TlsPasswordPrivateStructOnce.Do(func() {
		TlsPasswordPrivateStruct = gi.StructNew("Gio", "TlsPasswordPrivate")
	})
}

type TlsPasswordPrivate struct {
	native uintptr
}

var TlsServerConnectionInterfaceStruct *gi.Struct
var TlsServerConnectionInterfaceStructOnce sync.Once

func TlsServerConnectionInterfaceStructSet() {
	TlsServerConnectionInterfaceStructOnce.Do(func() {
		TlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "TlsServerConnectionInterface")
	})
}

type TlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var UnixConnectionClassStruct *gi.Struct
var UnixConnectionClassStructOnce sync.Once

func UnixConnectionClassStructSet() {
	UnixConnectionClassStructOnce.Do(func() {
		UnixConnectionClassStruct = gi.StructNew("Gio", "UnixConnectionClass")
	})
}

type UnixConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var UnixConnectionPrivateStruct *gi.Struct
var UnixConnectionPrivateStructOnce sync.Once

func UnixConnectionPrivateStructSet() {
	UnixConnectionPrivateStructOnce.Do(func() {
		UnixConnectionPrivateStruct = gi.StructNew("Gio", "UnixConnectionPrivate")
	})
}

type UnixConnectionPrivate struct {
	native uintptr
}

var UnixCredentialsMessageClassStruct *gi.Struct
var UnixCredentialsMessageClassStructOnce sync.Once

func UnixCredentialsMessageClassStructSet() {
	UnixCredentialsMessageClassStructOnce.Do(func() {
		UnixCredentialsMessageClassStruct = gi.StructNew("Gio", "UnixCredentialsMessageClass")
	})
}

type UnixCredentialsMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var UnixCredentialsMessagePrivateStruct *gi.Struct
var UnixCredentialsMessagePrivateStructOnce sync.Once

func UnixCredentialsMessagePrivateStructSet() {
	UnixCredentialsMessagePrivateStructOnce.Do(func() {
		UnixCredentialsMessagePrivateStruct = gi.StructNew("Gio", "UnixCredentialsMessagePrivate")
	})
}

type UnixCredentialsMessagePrivate struct {
	native uintptr
}

var UnixFDListClassStruct *gi.Struct
var UnixFDListClassStructOnce sync.Once

func UnixFDListClassStructSet() {
	UnixFDListClassStructOnce.Do(func() {
		UnixFDListClassStruct = gi.StructNew("Gio", "UnixFDListClass")
	})
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

var UnixFDListPrivateStruct *gi.Struct
var UnixFDListPrivateStructOnce sync.Once

func UnixFDListPrivateStructSet() {
	UnixFDListPrivateStructOnce.Do(func() {
		UnixFDListPrivateStruct = gi.StructNew("Gio", "UnixFDListPrivate")
	})
}

type UnixFDListPrivate struct {
	native uintptr
}

var UnixFDMessageClassStruct *gi.Struct
var UnixFDMessageClassStructOnce sync.Once

func UnixFDMessageClassStructSet() {
	UnixFDMessageClassStructOnce.Do(func() {
		UnixFDMessageClassStruct = gi.StructNew("Gio", "UnixFDMessageClass")
	})
}

type UnixFDMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var UnixFDMessagePrivateStruct *gi.Struct
var UnixFDMessagePrivateStructOnce sync.Once

func UnixFDMessagePrivateStructSet() {
	UnixFDMessagePrivateStructOnce.Do(func() {
		UnixFDMessagePrivateStruct = gi.StructNew("Gio", "UnixFDMessagePrivate")
	})
}

type UnixFDMessagePrivate struct {
	native uintptr
}

var UnixInputStreamClassStruct *gi.Struct
var UnixInputStreamClassStructOnce sync.Once

func UnixInputStreamClassStructSet() {
	UnixInputStreamClassStructOnce.Do(func() {
		UnixInputStreamClassStruct = gi.StructNew("Gio", "UnixInputStreamClass")
	})
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

var UnixInputStreamPrivateStruct *gi.Struct
var UnixInputStreamPrivateStructOnce sync.Once

func UnixInputStreamPrivateStructSet() {
	UnixInputStreamPrivateStructOnce.Do(func() {
		UnixInputStreamPrivateStruct = gi.StructNew("Gio", "UnixInputStreamPrivate")
	})
}

type UnixInputStreamPrivate struct {
	native uintptr
}

var UnixMountEntryStruct *gi.Struct
var UnixMountEntryStructOnce sync.Once

func UnixMountEntryStructSet() {
	UnixMountEntryStructOnce.Do(func() {
		UnixMountEntryStruct = gi.StructNew("Gio", "UnixMountEntry")
	})
}

type UnixMountEntry struct {
	native uintptr
}

var UnixMountMonitorClassStruct *gi.Struct
var UnixMountMonitorClassStructOnce sync.Once

func UnixMountMonitorClassStructSet() {
	UnixMountMonitorClassStructOnce.Do(func() {
		UnixMountMonitorClassStruct = gi.StructNew("Gio", "UnixMountMonitorClass")
	})
}

type UnixMountMonitorClass struct {
	native uintptr
}

var UnixMountPointStruct *gi.Struct
var UnixMountPointStructOnce sync.Once

func UnixMountPointStructSet() {
	UnixMountPointStructOnce.Do(func() {
		UnixMountPointStruct = gi.StructNew("Gio", "UnixMountPoint")
	})
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

var UnixOutputStreamClassStruct *gi.Struct
var UnixOutputStreamClassStructOnce sync.Once

func UnixOutputStreamClassStructSet() {
	UnixOutputStreamClassStructOnce.Do(func() {
		UnixOutputStreamClassStruct = gi.StructNew("Gio", "UnixOutputStreamClass")
	})
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

var UnixOutputStreamPrivateStruct *gi.Struct
var UnixOutputStreamPrivateStructOnce sync.Once

func UnixOutputStreamPrivateStructSet() {
	UnixOutputStreamPrivateStructOnce.Do(func() {
		UnixOutputStreamPrivateStruct = gi.StructNew("Gio", "UnixOutputStreamPrivate")
	})
}

type UnixOutputStreamPrivate struct {
	native uintptr
}

var UnixSocketAddressClassStruct *gi.Struct
var UnixSocketAddressClassStructOnce sync.Once

func UnixSocketAddressClassStructSet() {
	UnixSocketAddressClassStructOnce.Do(func() {
		UnixSocketAddressClassStruct = gi.StructNew("Gio", "UnixSocketAddressClass")
	})
}

type UnixSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var UnixSocketAddressPrivateStruct *gi.Struct
var UnixSocketAddressPrivateStructOnce sync.Once

func UnixSocketAddressPrivateStructSet() {
	UnixSocketAddressPrivateStructOnce.Do(func() {
		UnixSocketAddressPrivateStruct = gi.StructNew("Gio", "UnixSocketAddressPrivate")
	})
}

type UnixSocketAddressPrivate struct {
	native uintptr
}

var VfsClassStruct *gi.Struct
var VfsClassStructOnce sync.Once

func VfsClassStructSet() {
	VfsClassStructOnce.Do(func() {
		VfsClassStruct = gi.StructNew("Gio", "VfsClass")
	})
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

var VolumeIfaceStruct *gi.Struct
var VolumeIfaceStructOnce sync.Once

func VolumeIfaceStructSet() {
	VolumeIfaceStructOnce.Do(func() {
		VolumeIfaceStruct = gi.StructNew("Gio", "VolumeIface")
	})
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

var VolumeMonitorClassStruct *gi.Struct
var VolumeMonitorClassStructOnce sync.Once

func VolumeMonitorClassStructSet() {
	VolumeMonitorClassStructOnce.Do(func() {
		VolumeMonitorClassStruct = gi.StructNew("Gio", "VolumeMonitorClass")
	})
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

var ZlibCompressorClassStruct *gi.Struct
var ZlibCompressorClassStructOnce sync.Once

func ZlibCompressorClassStructSet() {
	ZlibCompressorClassStructOnce.Do(func() {
		ZlibCompressorClassStruct = gi.StructNew("Gio", "ZlibCompressorClass")
	})
}

type ZlibCompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var ZlibDecompressorClassStruct *gi.Struct
var ZlibDecompressorClassStructOnce sync.Once

func ZlibDecompressorClassStructSet() {
	ZlibDecompressorClassStructOnce.Do(func() {
		ZlibDecompressorClassStruct = gi.StructNew("Gio", "ZlibDecompressorClass")
	})
}

type ZlibDecompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}
