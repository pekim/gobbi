// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionEntryStruct *gi.Struct
var actionEntryStructOnce sync.Once

func actionEntryStructSet() {
	actionEntryStructOnce.Do(func() {
		actionEntryStruct = gi.StructNew("Gio", "ActionEntry")
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

var actionGroupInterfaceStruct *gi.Struct
var actionGroupInterfaceStructOnce sync.Once

func actionGroupInterfaceStructSet() {
	actionGroupInterfaceStructOnce.Do(func() {
		actionGroupInterfaceStruct = gi.StructNew("Gio", "ActionGroupInterface")
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

var actionInterfaceStruct *gi.Struct
var actionInterfaceStructOnce sync.Once

func actionInterfaceStructSet() {
	actionInterfaceStructOnce.Do(func() {
		actionInterfaceStruct = gi.StructNew("Gio", "ActionInterface")
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

var actionMapInterfaceStruct *gi.Struct
var actionMapInterfaceStructOnce sync.Once

func actionMapInterfaceStructSet() {
	actionMapInterfaceStructOnce.Do(func() {
		actionMapInterfaceStruct = gi.StructNew("Gio", "ActionMapInterface")
	})
}

type ActionMapInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'lookup_action' : missing Type
	// UNSUPPORTED : C value 'add_action' : missing Type
	// UNSUPPORTED : C value 'remove_action' : missing Type
}

var appInfoIfaceStruct *gi.Struct
var appInfoIfaceStructOnce sync.Once

func appInfoIfaceStructSet() {
	appInfoIfaceStructOnce.Do(func() {
		appInfoIfaceStruct = gi.StructNew("Gio", "AppInfoIface")
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

var appLaunchContextClassStruct *gi.Struct
var appLaunchContextClassStructOnce sync.Once

func appLaunchContextClassStructSet() {
	appLaunchContextClassStructOnce.Do(func() {
		appLaunchContextClassStruct = gi.StructNew("Gio", "AppLaunchContextClass")
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

var appLaunchContextPrivateStruct *gi.Struct
var appLaunchContextPrivateStructOnce sync.Once

func appLaunchContextPrivateStructSet() {
	appLaunchContextPrivateStructOnce.Do(func() {
		appLaunchContextPrivateStruct = gi.StructNew("Gio", "AppLaunchContextPrivate")
	})
}

type AppLaunchContextPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStructOnce sync.Once

func applicationClassStructSet() {
	applicationClassStructOnce.Do(func() {
		applicationClassStruct = gi.StructNew("Gio", "ApplicationClass")
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

var applicationCommandLineClassStruct *gi.Struct
var applicationCommandLineClassStructOnce sync.Once

func applicationCommandLineClassStructSet() {
	applicationCommandLineClassStructOnce.Do(func() {
		applicationCommandLineClassStruct = gi.StructNew("Gio", "ApplicationCommandLineClass")
	})
}

type ApplicationCommandLineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'print_literal' : missing Type
	// UNSUPPORTED : C value 'printerr_literal' : missing Type
	// UNSUPPORTED : C value 'get_stdin' : missing Type
}

var applicationCommandLinePrivateStruct *gi.Struct
var applicationCommandLinePrivateStructOnce sync.Once

func applicationCommandLinePrivateStructSet() {
	applicationCommandLinePrivateStructOnce.Do(func() {
		applicationCommandLinePrivateStruct = gi.StructNew("Gio", "ApplicationCommandLinePrivate")
	})
}

type ApplicationCommandLinePrivate struct {
	native uintptr
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStructOnce sync.Once

func applicationPrivateStructSet() {
	applicationPrivateStructOnce.Do(func() {
		applicationPrivateStruct = gi.StructNew("Gio", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var asyncInitableIfaceStruct *gi.Struct
var asyncInitableIfaceStructOnce sync.Once

func asyncInitableIfaceStructSet() {
	asyncInitableIfaceStructOnce.Do(func() {
		asyncInitableIfaceStruct = gi.StructNew("Gio", "AsyncInitableIface")
	})
}

type AsyncInitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'init_async' : missing Type
	// UNSUPPORTED : C value 'init_finish' : missing Type
}

var asyncResultIfaceStruct *gi.Struct
var asyncResultIfaceStructOnce sync.Once

func asyncResultIfaceStructSet() {
	asyncResultIfaceStructOnce.Do(func() {
		asyncResultIfaceStruct = gi.StructNew("Gio", "AsyncResultIface")
	})
}

type AsyncResultIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_user_data' : missing Type
	// UNSUPPORTED : C value 'get_source_object' : missing Type
	// UNSUPPORTED : C value 'is_tagged' : missing Type
}

var bufferedInputStreamClassStruct *gi.Struct
var bufferedInputStreamClassStructOnce sync.Once

func bufferedInputStreamClassStructSet() {
	bufferedInputStreamClassStructOnce.Do(func() {
		bufferedInputStreamClassStruct = gi.StructNew("Gio", "BufferedInputStreamClass")
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

var bufferedInputStreamPrivateStruct *gi.Struct
var bufferedInputStreamPrivateStructOnce sync.Once

func bufferedInputStreamPrivateStructSet() {
	bufferedInputStreamPrivateStructOnce.Do(func() {
		bufferedInputStreamPrivateStruct = gi.StructNew("Gio", "BufferedInputStreamPrivate")
	})
}

type BufferedInputStreamPrivate struct {
	native uintptr
}

var bufferedOutputStreamClassStruct *gi.Struct
var bufferedOutputStreamClassStructOnce sync.Once

func bufferedOutputStreamClassStructSet() {
	bufferedOutputStreamClassStructOnce.Do(func() {
		bufferedOutputStreamClassStruct = gi.StructNew("Gio", "BufferedOutputStreamClass")
	})
}

type BufferedOutputStreamClass struct {
	native      uintptr
	ParentClass *FilterOutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var bufferedOutputStreamPrivateStruct *gi.Struct
var bufferedOutputStreamPrivateStructOnce sync.Once

func bufferedOutputStreamPrivateStructSet() {
	bufferedOutputStreamPrivateStructOnce.Do(func() {
		bufferedOutputStreamPrivateStruct = gi.StructNew("Gio", "BufferedOutputStreamPrivate")
	})
}

type BufferedOutputStreamPrivate struct {
	native uintptr
}

var cancellableClassStruct *gi.Struct
var cancellableClassStructOnce sync.Once

func cancellableClassStructSet() {
	cancellableClassStructOnce.Do(func() {
		cancellableClassStruct = gi.StructNew("Gio", "CancellableClass")
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

var cancellablePrivateStruct *gi.Struct
var cancellablePrivateStructOnce sync.Once

func cancellablePrivateStructSet() {
	cancellablePrivateStructOnce.Do(func() {
		cancellablePrivateStruct = gi.StructNew("Gio", "CancellablePrivate")
	})
}

type CancellablePrivate struct {
	native uintptr
}

var charsetConverterClassStruct *gi.Struct
var charsetConverterClassStructOnce sync.Once

func charsetConverterClassStructSet() {
	charsetConverterClassStructOnce.Do(func() {
		charsetConverterClassStruct = gi.StructNew("Gio", "CharsetConverterClass")
	})
}

type CharsetConverterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var converterIfaceStruct *gi.Struct
var converterIfaceStructOnce sync.Once

func converterIfaceStructSet() {
	converterIfaceStructOnce.Do(func() {
		converterIfaceStruct = gi.StructNew("Gio", "ConverterIface")
	})
}

type ConverterIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'convert' : missing Type
	// UNSUPPORTED : C value 'reset' : missing Type
}

var converterInputStreamClassStruct *gi.Struct
var converterInputStreamClassStructOnce sync.Once

func converterInputStreamClassStructSet() {
	converterInputStreamClassStructOnce.Do(func() {
		converterInputStreamClassStruct = gi.StructNew("Gio", "ConverterInputStreamClass")
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

var converterInputStreamPrivateStruct *gi.Struct
var converterInputStreamPrivateStructOnce sync.Once

func converterInputStreamPrivateStructSet() {
	converterInputStreamPrivateStructOnce.Do(func() {
		converterInputStreamPrivateStruct = gi.StructNew("Gio", "ConverterInputStreamPrivate")
	})
}

type ConverterInputStreamPrivate struct {
	native uintptr
}

var converterOutputStreamClassStruct *gi.Struct
var converterOutputStreamClassStructOnce sync.Once

func converterOutputStreamClassStructSet() {
	converterOutputStreamClassStructOnce.Do(func() {
		converterOutputStreamClassStruct = gi.StructNew("Gio", "ConverterOutputStreamClass")
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

var converterOutputStreamPrivateStruct *gi.Struct
var converterOutputStreamPrivateStructOnce sync.Once

func converterOutputStreamPrivateStructSet() {
	converterOutputStreamPrivateStructOnce.Do(func() {
		converterOutputStreamPrivateStruct = gi.StructNew("Gio", "ConverterOutputStreamPrivate")
	})
}

type ConverterOutputStreamPrivate struct {
	native uintptr
}

var credentialsClassStruct *gi.Struct
var credentialsClassStructOnce sync.Once

func credentialsClassStructSet() {
	credentialsClassStructOnce.Do(func() {
		credentialsClassStruct = gi.StructNew("Gio", "CredentialsClass")
	})
}

type CredentialsClass struct {
	native uintptr
}

var dBusAnnotationInfoStruct *gi.Struct
var dBusAnnotationInfoStructOnce sync.Once

func dBusAnnotationInfoStructSet() {
	dBusAnnotationInfoStructOnce.Do(func() {
		dBusAnnotationInfoStruct = gi.StructNew("Gio", "DBusAnnotationInfo")
	})
}

type DBusAnnotationInfo struct {
	native   uintptr
	RefCount int32
	Key      string
	Value    string
	// UNSUPPORTED : C value 'annotations' : missing Type
}

var dBusAnnotationInfoRefFunction *gi.Function
var dBusAnnotationInfoRefFunctionOnce sync.Once

func dBusAnnotationInfoRefFunctionSet() {
	dBusAnnotationInfoRefFunctionOnce.Do(func() {
		dBusAnnotationInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusAnnotationInfoUnrefFunction *gi.Function
var dBusAnnotationInfoUnrefFunctionOnce sync.Once

func dBusAnnotationInfoUnrefFunctionSet() {
	dBusAnnotationInfoUnrefFunctionOnce.Do(func() {
		dBusAnnotationInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusArgInfoStruct *gi.Struct
var dBusArgInfoStructOnce sync.Once

func dBusArgInfoStructSet() {
	dBusArgInfoStructOnce.Do(func() {
		dBusArgInfoStruct = gi.StructNew("Gio", "DBusArgInfo")
	})
}

type DBusArgInfo struct {
	native    uintptr
	RefCount  int32
	Name      string
	Signature string
	// UNSUPPORTED : C value 'annotations' : missing Type
}

var dBusArgInfoRefFunction *gi.Function
var dBusArgInfoRefFunctionOnce sync.Once

func dBusArgInfoRefFunctionSet() {
	dBusArgInfoRefFunctionOnce.Do(func() {
		dBusArgInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusArgInfoUnrefFunction *gi.Function
var dBusArgInfoUnrefFunctionOnce sync.Once

func dBusArgInfoUnrefFunctionSet() {
	dBusArgInfoUnrefFunctionOnce.Do(func() {
		dBusArgInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusErrorEntryStruct *gi.Struct
var dBusErrorEntryStructOnce sync.Once

func dBusErrorEntryStructSet() {
	dBusErrorEntryStructOnce.Do(func() {
		dBusErrorEntryStruct = gi.StructNew("Gio", "DBusErrorEntry")
	})
}

type DBusErrorEntry struct {
	native        uintptr
	ErrorCode     int32
	DbusErrorName string
}

var dBusInterfaceIfaceStruct *gi.Struct
var dBusInterfaceIfaceStructOnce sync.Once

func dBusInterfaceIfaceStructSet() {
	dBusInterfaceIfaceStructOnce.Do(func() {
		dBusInterfaceIfaceStruct = gi.StructNew("Gio", "DBusInterfaceIface")
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

var dBusInterfaceInfoStruct *gi.Struct
var dBusInterfaceInfoStructOnce sync.Once

func dBusInterfaceInfoStructSet() {
	dBusInterfaceInfoStructOnce.Do(func() {
		dBusInterfaceInfoStruct = gi.StructNew("Gio", "DBusInterfaceInfo")
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

var dBusInterfaceInfoCacheBuildFunction *gi.Function
var dBusInterfaceInfoCacheBuildFunctionOnce sync.Once

func dBusInterfaceInfoCacheBuildFunctionSet() {
	dBusInterfaceInfoCacheBuildFunctionOnce.Do(func() {
		dBusInterfaceInfoCacheBuildFunction = gi.FunctionInvokerNew("Gio", "cache_build")
	})
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

var dBusInterfaceInfoCacheReleaseFunction *gi.Function
var dBusInterfaceInfoCacheReleaseFunctionOnce sync.Once

func dBusInterfaceInfoCacheReleaseFunctionSet() {
	dBusInterfaceInfoCacheReleaseFunctionOnce.Do(func() {
		dBusInterfaceInfoCacheReleaseFunction = gi.FunctionInvokerNew("Gio", "cache_release")
	})
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

var dBusInterfaceInfoLookupMethodFunction *gi.Function
var dBusInterfaceInfoLookupMethodFunctionOnce sync.Once

func dBusInterfaceInfoLookupMethodFunctionSet() {
	dBusInterfaceInfoLookupMethodFunctionOnce.Do(func() {
		dBusInterfaceInfoLookupMethodFunction = gi.FunctionInvokerNew("Gio", "lookup_method")
	})
}

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

var dBusInterfaceInfoLookupPropertyFunction *gi.Function
var dBusInterfaceInfoLookupPropertyFunctionOnce sync.Once

func dBusInterfaceInfoLookupPropertyFunctionSet() {
	dBusInterfaceInfoLookupPropertyFunctionOnce.Do(func() {
		dBusInterfaceInfoLookupPropertyFunction = gi.FunctionInvokerNew("Gio", "lookup_property")
	})
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

var dBusInterfaceInfoLookupSignalFunction *gi.Function
var dBusInterfaceInfoLookupSignalFunctionOnce sync.Once

func dBusInterfaceInfoLookupSignalFunctionSet() {
	dBusInterfaceInfoLookupSignalFunctionOnce.Do(func() {
		dBusInterfaceInfoLookupSignalFunction = gi.FunctionInvokerNew("Gio", "lookup_signal")
	})
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

var dBusInterfaceInfoRefFunction *gi.Function
var dBusInterfaceInfoRefFunctionOnce sync.Once

func dBusInterfaceInfoRefFunctionSet() {
	dBusInterfaceInfoRefFunctionOnce.Do(func() {
		dBusInterfaceInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusInterfaceInfoUnrefFunction *gi.Function
var dBusInterfaceInfoUnrefFunctionOnce sync.Once

func dBusInterfaceInfoUnrefFunctionSet() {
	dBusInterfaceInfoUnrefFunctionOnce.Do(func() {
		dBusInterfaceInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusInterfaceSkeletonClassStruct *gi.Struct
var dBusInterfaceSkeletonClassStructOnce sync.Once

func dBusInterfaceSkeletonClassStructSet() {
	dBusInterfaceSkeletonClassStructOnce.Do(func() {
		dBusInterfaceSkeletonClassStruct = gi.StructNew("Gio", "DBusInterfaceSkeletonClass")
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

var dBusInterfaceSkeletonPrivateStruct *gi.Struct
var dBusInterfaceSkeletonPrivateStructOnce sync.Once

func dBusInterfaceSkeletonPrivateStructSet() {
	dBusInterfaceSkeletonPrivateStructOnce.Do(func() {
		dBusInterfaceSkeletonPrivateStruct = gi.StructNew("Gio", "DBusInterfaceSkeletonPrivate")
	})
}

type DBusInterfaceSkeletonPrivate struct {
	native uintptr
}

var dBusInterfaceVTableStruct *gi.Struct
var dBusInterfaceVTableStructOnce sync.Once

func dBusInterfaceVTableStructSet() {
	dBusInterfaceVTableStructOnce.Do(func() {
		dBusInterfaceVTableStruct = gi.StructNew("Gio", "DBusInterfaceVTable")
	})
}

type DBusInterfaceVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'method_call' : no Go type for 'DBusInterfaceMethodCallFunc'
	// UNSUPPORTED : C value 'get_property' : no Go type for 'DBusInterfaceGetPropertyFunc'
	// UNSUPPORTED : C value 'set_property' : no Go type for 'DBusInterfaceSetPropertyFunc'
}

var dBusMethodInfoStruct *gi.Struct
var dBusMethodInfoStructOnce sync.Once

func dBusMethodInfoStructSet() {
	dBusMethodInfoStructOnce.Do(func() {
		dBusMethodInfoStruct = gi.StructNew("Gio", "DBusMethodInfo")
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

var dBusMethodInfoRefFunction *gi.Function
var dBusMethodInfoRefFunctionOnce sync.Once

func dBusMethodInfoRefFunctionSet() {
	dBusMethodInfoRefFunctionOnce.Do(func() {
		dBusMethodInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusMethodInfoUnrefFunction *gi.Function
var dBusMethodInfoUnrefFunctionOnce sync.Once

func dBusMethodInfoUnrefFunctionSet() {
	dBusMethodInfoUnrefFunctionOnce.Do(func() {
		dBusMethodInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusNodeInfoStruct *gi.Struct
var dBusNodeInfoStructOnce sync.Once

func dBusNodeInfoStructSet() {
	dBusNodeInfoStructOnce.Do(func() {
		dBusNodeInfoStruct = gi.StructNew("Gio", "DBusNodeInfo")
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

var dBusNodeInfoNewForXmlFunction *gi.Function
var dBusNodeInfoNewForXmlFunctionOnce sync.Once

func dBusNodeInfoNewForXmlFunctionSet() {
	dBusNodeInfoNewForXmlFunctionOnce.Do(func() {
		dBusNodeInfoNewForXmlFunction = gi.FunctionInvokerNew("Gio", "new_for_xml")
	})
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

var dBusNodeInfoLookupInterfaceFunction *gi.Function
var dBusNodeInfoLookupInterfaceFunctionOnce sync.Once

func dBusNodeInfoLookupInterfaceFunctionSet() {
	dBusNodeInfoLookupInterfaceFunctionOnce.Do(func() {
		dBusNodeInfoLookupInterfaceFunction = gi.FunctionInvokerNew("Gio", "lookup_interface")
	})
}

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

var dBusNodeInfoRefFunction *gi.Function
var dBusNodeInfoRefFunctionOnce sync.Once

func dBusNodeInfoRefFunctionSet() {
	dBusNodeInfoRefFunctionOnce.Do(func() {
		dBusNodeInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusNodeInfoUnrefFunction *gi.Function
var dBusNodeInfoUnrefFunctionOnce sync.Once

func dBusNodeInfoUnrefFunctionSet() {
	dBusNodeInfoUnrefFunctionOnce.Do(func() {
		dBusNodeInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusObjectIfaceStruct *gi.Struct
var dBusObjectIfaceStructOnce sync.Once

func dBusObjectIfaceStructSet() {
	dBusObjectIfaceStructOnce.Do(func() {
		dBusObjectIfaceStruct = gi.StructNew("Gio", "DBusObjectIface")
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

var dBusObjectManagerClientClassStruct *gi.Struct
var dBusObjectManagerClientClassStructOnce sync.Once

func dBusObjectManagerClientClassStructSet() {
	dBusObjectManagerClientClassStructOnce.Do(func() {
		dBusObjectManagerClientClassStruct = gi.StructNew("Gio", "DBusObjectManagerClientClass")
	})
}

type DBusObjectManagerClientClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'interface_proxy_signal' : missing Type
	// UNSUPPORTED : C value 'interface_proxy_properties_changed' : missing Type
}

var dBusObjectManagerClientPrivateStruct *gi.Struct
var dBusObjectManagerClientPrivateStructOnce sync.Once

func dBusObjectManagerClientPrivateStructSet() {
	dBusObjectManagerClientPrivateStructOnce.Do(func() {
		dBusObjectManagerClientPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerClientPrivate")
	})
}

type DBusObjectManagerClientPrivate struct {
	native uintptr
}

var dBusObjectManagerIfaceStruct *gi.Struct
var dBusObjectManagerIfaceStructOnce sync.Once

func dBusObjectManagerIfaceStructSet() {
	dBusObjectManagerIfaceStructOnce.Do(func() {
		dBusObjectManagerIfaceStruct = gi.StructNew("Gio", "DBusObjectManagerIface")
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

var dBusObjectManagerServerClassStruct *gi.Struct
var dBusObjectManagerServerClassStructOnce sync.Once

func dBusObjectManagerServerClassStructSet() {
	dBusObjectManagerServerClassStructOnce.Do(func() {
		dBusObjectManagerServerClassStruct = gi.StructNew("Gio", "DBusObjectManagerServerClass")
	})
}

type DBusObjectManagerServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var dBusObjectManagerServerPrivateStruct *gi.Struct
var dBusObjectManagerServerPrivateStructOnce sync.Once

func dBusObjectManagerServerPrivateStructSet() {
	dBusObjectManagerServerPrivateStructOnce.Do(func() {
		dBusObjectManagerServerPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerServerPrivate")
	})
}

type DBusObjectManagerServerPrivate struct {
	native uintptr
}

var dBusObjectProxyClassStruct *gi.Struct
var dBusObjectProxyClassStructOnce sync.Once

func dBusObjectProxyClassStructSet() {
	dBusObjectProxyClassStructOnce.Do(func() {
		dBusObjectProxyClassStruct = gi.StructNew("Gio", "DBusObjectProxyClass")
	})
}

type DBusObjectProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var dBusObjectProxyPrivateStruct *gi.Struct
var dBusObjectProxyPrivateStructOnce sync.Once

func dBusObjectProxyPrivateStructSet() {
	dBusObjectProxyPrivateStructOnce.Do(func() {
		dBusObjectProxyPrivateStruct = gi.StructNew("Gio", "DBusObjectProxyPrivate")
	})
}

type DBusObjectProxyPrivate struct {
	native uintptr
}

var dBusObjectSkeletonClassStruct *gi.Struct
var dBusObjectSkeletonClassStructOnce sync.Once

func dBusObjectSkeletonClassStructSet() {
	dBusObjectSkeletonClassStructOnce.Do(func() {
		dBusObjectSkeletonClassStruct = gi.StructNew("Gio", "DBusObjectSkeletonClass")
	})
}

type DBusObjectSkeletonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authorize_method' : missing Type
}

var dBusObjectSkeletonPrivateStruct *gi.Struct
var dBusObjectSkeletonPrivateStructOnce sync.Once

func dBusObjectSkeletonPrivateStructSet() {
	dBusObjectSkeletonPrivateStructOnce.Do(func() {
		dBusObjectSkeletonPrivateStruct = gi.StructNew("Gio", "DBusObjectSkeletonPrivate")
	})
}

type DBusObjectSkeletonPrivate struct {
	native uintptr
}

var dBusPropertyInfoStruct *gi.Struct
var dBusPropertyInfoStructOnce sync.Once

func dBusPropertyInfoStructSet() {
	dBusPropertyInfoStructOnce.Do(func() {
		dBusPropertyInfoStruct = gi.StructNew("Gio", "DBusPropertyInfo")
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

var dBusPropertyInfoRefFunction *gi.Function
var dBusPropertyInfoRefFunctionOnce sync.Once

func dBusPropertyInfoRefFunctionSet() {
	dBusPropertyInfoRefFunctionOnce.Do(func() {
		dBusPropertyInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusPropertyInfoUnrefFunction *gi.Function
var dBusPropertyInfoUnrefFunctionOnce sync.Once

func dBusPropertyInfoUnrefFunctionSet() {
	dBusPropertyInfoUnrefFunctionOnce.Do(func() {
		dBusPropertyInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusProxyClassStruct *gi.Struct
var dBusProxyClassStructOnce sync.Once

func dBusProxyClassStructSet() {
	dBusProxyClassStructOnce.Do(func() {
		dBusProxyClassStruct = gi.StructNew("Gio", "DBusProxyClass")
	})
}

type DBusProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'g_properties_changed' : missing Type
	// UNSUPPORTED : C value 'g_signal' : missing Type
}

var dBusProxyPrivateStruct *gi.Struct
var dBusProxyPrivateStructOnce sync.Once

func dBusProxyPrivateStructSet() {
	dBusProxyPrivateStructOnce.Do(func() {
		dBusProxyPrivateStruct = gi.StructNew("Gio", "DBusProxyPrivate")
	})
}

type DBusProxyPrivate struct {
	native uintptr
}

var dBusSignalInfoStruct *gi.Struct
var dBusSignalInfoStructOnce sync.Once

func dBusSignalInfoStructSet() {
	dBusSignalInfoStructOnce.Do(func() {
		dBusSignalInfoStruct = gi.StructNew("Gio", "DBusSignalInfo")
	})
}

type DBusSignalInfo struct {
	native   uintptr
	RefCount int32
	Name     string
	// UNSUPPORTED : C value 'args' : missing Type
	// UNSUPPORTED : C value 'annotations' : missing Type
}

var dBusSignalInfoRefFunction *gi.Function
var dBusSignalInfoRefFunctionOnce sync.Once

func dBusSignalInfoRefFunctionSet() {
	dBusSignalInfoRefFunctionOnce.Do(func() {
		dBusSignalInfoRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var dBusSignalInfoUnrefFunction *gi.Function
var dBusSignalInfoUnrefFunctionOnce sync.Once

func dBusSignalInfoUnrefFunctionSet() {
	dBusSignalInfoUnrefFunctionOnce.Do(func() {
		dBusSignalInfoUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var dBusSubtreeVTableStruct *gi.Struct
var dBusSubtreeVTableStructOnce sync.Once

func dBusSubtreeVTableStructSet() {
	dBusSubtreeVTableStructOnce.Do(func() {
		dBusSubtreeVTableStruct = gi.StructNew("Gio", "DBusSubtreeVTable")
	})
}

type DBusSubtreeVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'enumerate' : no Go type for 'DBusSubtreeEnumerateFunc'
	// UNSUPPORTED : C value 'introspect' : no Go type for 'DBusSubtreeIntrospectFunc'
	// UNSUPPORTED : C value 'dispatch' : no Go type for 'DBusSubtreeDispatchFunc'
}

var dataInputStreamClassStruct *gi.Struct
var dataInputStreamClassStructOnce sync.Once

func dataInputStreamClassStructSet() {
	dataInputStreamClassStructOnce.Do(func() {
		dataInputStreamClassStruct = gi.StructNew("Gio", "DataInputStreamClass")
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

var dataInputStreamPrivateStruct *gi.Struct
var dataInputStreamPrivateStructOnce sync.Once

func dataInputStreamPrivateStructSet() {
	dataInputStreamPrivateStructOnce.Do(func() {
		dataInputStreamPrivateStruct = gi.StructNew("Gio", "DataInputStreamPrivate")
	})
}

type DataInputStreamPrivate struct {
	native uintptr
}

var dataOutputStreamClassStruct *gi.Struct
var dataOutputStreamClassStructOnce sync.Once

func dataOutputStreamClassStructSet() {
	dataOutputStreamClassStructOnce.Do(func() {
		dataOutputStreamClassStruct = gi.StructNew("Gio", "DataOutputStreamClass")
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

var dataOutputStreamPrivateStruct *gi.Struct
var dataOutputStreamPrivateStructOnce sync.Once

func dataOutputStreamPrivateStructSet() {
	dataOutputStreamPrivateStructOnce.Do(func() {
		dataOutputStreamPrivateStruct = gi.StructNew("Gio", "DataOutputStreamPrivate")
	})
}

type DataOutputStreamPrivate struct {
	native uintptr
}

var datagramBasedInterfaceStruct *gi.Struct
var datagramBasedInterfaceStructOnce sync.Once

func datagramBasedInterfaceStructSet() {
	datagramBasedInterfaceStructOnce.Do(func() {
		datagramBasedInterfaceStruct = gi.StructNew("Gio", "DatagramBasedInterface")
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

var desktopAppInfoClassStruct *gi.Struct
var desktopAppInfoClassStructOnce sync.Once

func desktopAppInfoClassStructSet() {
	desktopAppInfoClassStructOnce.Do(func() {
		desktopAppInfoClassStruct = gi.StructNew("Gio", "DesktopAppInfoClass")
	})
}

type DesktopAppInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var desktopAppInfoLookupIfaceStruct *gi.Struct
var desktopAppInfoLookupIfaceStructOnce sync.Once

func desktopAppInfoLookupIfaceStructSet() {
	desktopAppInfoLookupIfaceStructOnce.Do(func() {
		desktopAppInfoLookupIfaceStruct = gi.StructNew("Gio", "DesktopAppInfoLookupIface")
	})
}

type DesktopAppInfoLookupIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_default_for_uri_scheme' : missing Type
}

var driveIfaceStruct *gi.Struct
var driveIfaceStructOnce sync.Once

func driveIfaceStructSet() {
	driveIfaceStructOnce.Do(func() {
		driveIfaceStruct = gi.StructNew("Gio", "DriveIface")
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

var dtlsClientConnectionInterfaceStruct *gi.Struct
var dtlsClientConnectionInterfaceStructOnce sync.Once

func dtlsClientConnectionInterfaceStructSet() {
	dtlsClientConnectionInterfaceStructOnce.Do(func() {
		dtlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsClientConnectionInterface")
	})
}

type DtlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var dtlsConnectionInterfaceStruct *gi.Struct
var dtlsConnectionInterfaceStructOnce sync.Once

func dtlsConnectionInterfaceStructSet() {
	dtlsConnectionInterfaceStructOnce.Do(func() {
		dtlsConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsConnectionInterface")
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

var dtlsServerConnectionInterfaceStruct *gi.Struct
var dtlsServerConnectionInterfaceStructOnce sync.Once

func dtlsServerConnectionInterfaceStructSet() {
	dtlsServerConnectionInterfaceStructOnce.Do(func() {
		dtlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsServerConnectionInterface")
	})
}

type DtlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var emblemClassStruct *gi.Struct
var emblemClassStructOnce sync.Once

func emblemClassStructSet() {
	emblemClassStructOnce.Do(func() {
		emblemClassStruct = gi.StructNew("Gio", "EmblemClass")
	})
}

type EmblemClass struct {
	native uintptr
}

var emblemedIconClassStruct *gi.Struct
var emblemedIconClassStructOnce sync.Once

func emblemedIconClassStructSet() {
	emblemedIconClassStructOnce.Do(func() {
		emblemedIconClassStruct = gi.StructNew("Gio", "EmblemedIconClass")
	})
}

type EmblemedIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var emblemedIconPrivateStruct *gi.Struct
var emblemedIconPrivateStructOnce sync.Once

func emblemedIconPrivateStructSet() {
	emblemedIconPrivateStructOnce.Do(func() {
		emblemedIconPrivateStruct = gi.StructNew("Gio", "EmblemedIconPrivate")
	})
}

type EmblemedIconPrivate struct {
	native uintptr
}

var fileAttributeInfoStruct *gi.Struct
var fileAttributeInfoStructOnce sync.Once

func fileAttributeInfoStructSet() {
	fileAttributeInfoStructOnce.Do(func() {
		fileAttributeInfoStruct = gi.StructNew("Gio", "FileAttributeInfo")
	})
}

type FileAttributeInfo struct {
	native uintptr
	Name   string
	// UNSUPPORTED : C value 'type' : no Go type for 'FileAttributeType'
	// UNSUPPORTED : C value 'flags' : no Go type for 'FileAttributeInfoFlags'
}

var fileAttributeInfoListStruct *gi.Struct
var fileAttributeInfoListStructOnce sync.Once

func fileAttributeInfoListStructSet() {
	fileAttributeInfoListStructOnce.Do(func() {
		fileAttributeInfoListStruct = gi.StructNew("Gio", "FileAttributeInfoList")
	})
}

type FileAttributeInfoList struct {
	native uintptr
	Infos  *FileAttributeInfo
	NInfos int32
}

var fileAttributeInfoListNewFunction *gi.Function
var fileAttributeInfoListNewFunctionOnce sync.Once

func fileAttributeInfoListNewFunctionSet() {
	fileAttributeInfoListNewFunctionOnce.Do(func() {
		fileAttributeInfoListNewFunction = gi.FunctionInvokerNew("Gio", "new")
	})
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

var fileAttributeInfoListDupFunction *gi.Function
var fileAttributeInfoListDupFunctionOnce sync.Once

func fileAttributeInfoListDupFunctionSet() {
	fileAttributeInfoListDupFunctionOnce.Do(func() {
		fileAttributeInfoListDupFunction = gi.FunctionInvokerNew("Gio", "dup")
	})
}

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

var fileAttributeInfoListLookupFunction *gi.Function
var fileAttributeInfoListLookupFunctionOnce sync.Once

func fileAttributeInfoListLookupFunctionSet() {
	fileAttributeInfoListLookupFunctionOnce.Do(func() {
		fileAttributeInfoListLookupFunction = gi.FunctionInvokerNew("Gio", "lookup")
	})
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

var fileAttributeInfoListRefFunction *gi.Function
var fileAttributeInfoListRefFunctionOnce sync.Once

func fileAttributeInfoListRefFunctionSet() {
	fileAttributeInfoListRefFunctionOnce.Do(func() {
		fileAttributeInfoListRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var fileAttributeInfoListUnrefFunction *gi.Function
var fileAttributeInfoListUnrefFunctionOnce sync.Once

func fileAttributeInfoListUnrefFunctionSet() {
	fileAttributeInfoListUnrefFunctionOnce.Do(func() {
		fileAttributeInfoListUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var fileAttributeMatcherStruct *gi.Struct
var fileAttributeMatcherStructOnce sync.Once

func fileAttributeMatcherStructSet() {
	fileAttributeMatcherStructOnce.Do(func() {
		fileAttributeMatcherStruct = gi.StructNew("Gio", "FileAttributeMatcher")
	})
}

type FileAttributeMatcher struct {
	native uintptr
}

var fileAttributeMatcherNewFunction *gi.Function
var fileAttributeMatcherNewFunctionOnce sync.Once

func fileAttributeMatcherNewFunctionSet() {
	fileAttributeMatcherNewFunctionOnce.Do(func() {
		fileAttributeMatcherNewFunction = gi.FunctionInvokerNew("Gio", "new")
	})
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

var fileAttributeMatcherEnumerateNextFunction *gi.Function
var fileAttributeMatcherEnumerateNextFunctionOnce sync.Once

func fileAttributeMatcherEnumerateNextFunctionSet() {
	fileAttributeMatcherEnumerateNextFunctionOnce.Do(func() {
		fileAttributeMatcherEnumerateNextFunction = gi.FunctionInvokerNew("Gio", "enumerate_next")
	})
}

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

var fileAttributeMatcherRefFunction *gi.Function
var fileAttributeMatcherRefFunctionOnce sync.Once

func fileAttributeMatcherRefFunctionSet() {
	fileAttributeMatcherRefFunctionOnce.Do(func() {
		fileAttributeMatcherRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
}

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

var fileAttributeMatcherToStringFunction *gi.Function
var fileAttributeMatcherToStringFunctionOnce sync.Once

func fileAttributeMatcherToStringFunctionSet() {
	fileAttributeMatcherToStringFunctionOnce.Do(func() {
		fileAttributeMatcherToStringFunction = gi.FunctionInvokerNew("Gio", "to_string")
	})
}

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

var fileAttributeMatcherUnrefFunction *gi.Function
var fileAttributeMatcherUnrefFunctionOnce sync.Once

func fileAttributeMatcherUnrefFunctionSet() {
	fileAttributeMatcherUnrefFunctionOnce.Do(func() {
		fileAttributeMatcherUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var fileDescriptorBasedIfaceStruct *gi.Struct
var fileDescriptorBasedIfaceStructOnce sync.Once

func fileDescriptorBasedIfaceStructSet() {
	fileDescriptorBasedIfaceStructOnce.Do(func() {
		fileDescriptorBasedIfaceStruct = gi.StructNew("Gio", "FileDescriptorBasedIface")
	})
}

type FileDescriptorBasedIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_fd' : missing Type
}

var fileEnumeratorClassStruct *gi.Struct
var fileEnumeratorClassStructOnce sync.Once

func fileEnumeratorClassStructSet() {
	fileEnumeratorClassStructOnce.Do(func() {
		fileEnumeratorClassStruct = gi.StructNew("Gio", "FileEnumeratorClass")
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

var fileEnumeratorPrivateStruct *gi.Struct
var fileEnumeratorPrivateStructOnce sync.Once

func fileEnumeratorPrivateStructSet() {
	fileEnumeratorPrivateStructOnce.Do(func() {
		fileEnumeratorPrivateStruct = gi.StructNew("Gio", "FileEnumeratorPrivate")
	})
}

type FileEnumeratorPrivate struct {
	native uintptr
}

var fileIOStreamClassStruct *gi.Struct
var fileIOStreamClassStructOnce sync.Once

func fileIOStreamClassStructSet() {
	fileIOStreamClassStructOnce.Do(func() {
		fileIOStreamClassStruct = gi.StructNew("Gio", "FileIOStreamClass")
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

var fileIOStreamPrivateStruct *gi.Struct
var fileIOStreamPrivateStructOnce sync.Once

func fileIOStreamPrivateStructSet() {
	fileIOStreamPrivateStructOnce.Do(func() {
		fileIOStreamPrivateStruct = gi.StructNew("Gio", "FileIOStreamPrivate")
	})
}

type FileIOStreamPrivate struct {
	native uintptr
}

var fileIconClassStruct *gi.Struct
var fileIconClassStructOnce sync.Once

func fileIconClassStructSet() {
	fileIconClassStructOnce.Do(func() {
		fileIconClassStruct = gi.StructNew("Gio", "FileIconClass")
	})
}

type FileIconClass struct {
	native uintptr
}

var fileIfaceStruct *gi.Struct
var fileIfaceStructOnce sync.Once

func fileIfaceStructSet() {
	fileIfaceStructOnce.Do(func() {
		fileIfaceStruct = gi.StructNew("Gio", "FileIface")
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

var fileInfoClassStruct *gi.Struct
var fileInfoClassStructOnce sync.Once

func fileInfoClassStructSet() {
	fileInfoClassStructOnce.Do(func() {
		fileInfoClassStruct = gi.StructNew("Gio", "FileInfoClass")
	})
}

type FileInfoClass struct {
	native uintptr
}

var fileInputStreamClassStruct *gi.Struct
var fileInputStreamClassStructOnce sync.Once

func fileInputStreamClassStructSet() {
	fileInputStreamClassStructOnce.Do(func() {
		fileInputStreamClassStruct = gi.StructNew("Gio", "FileInputStreamClass")
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

var fileInputStreamPrivateStruct *gi.Struct
var fileInputStreamPrivateStructOnce sync.Once

func fileInputStreamPrivateStructSet() {
	fileInputStreamPrivateStructOnce.Do(func() {
		fileInputStreamPrivateStruct = gi.StructNew("Gio", "FileInputStreamPrivate")
	})
}

type FileInputStreamPrivate struct {
	native uintptr
}

var fileMonitorClassStruct *gi.Struct
var fileMonitorClassStructOnce sync.Once

func fileMonitorClassStructSet() {
	fileMonitorClassStructOnce.Do(func() {
		fileMonitorClassStruct = gi.StructNew("Gio", "FileMonitorClass")
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

var fileMonitorPrivateStruct *gi.Struct
var fileMonitorPrivateStructOnce sync.Once

func fileMonitorPrivateStructSet() {
	fileMonitorPrivateStructOnce.Do(func() {
		fileMonitorPrivateStruct = gi.StructNew("Gio", "FileMonitorPrivate")
	})
}

type FileMonitorPrivate struct {
	native uintptr
}

var fileOutputStreamClassStruct *gi.Struct
var fileOutputStreamClassStructOnce sync.Once

func fileOutputStreamClassStructSet() {
	fileOutputStreamClassStructOnce.Do(func() {
		fileOutputStreamClassStruct = gi.StructNew("Gio", "FileOutputStreamClass")
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

var fileOutputStreamPrivateStruct *gi.Struct
var fileOutputStreamPrivateStructOnce sync.Once

func fileOutputStreamPrivateStructSet() {
	fileOutputStreamPrivateStructOnce.Do(func() {
		fileOutputStreamPrivateStruct = gi.StructNew("Gio", "FileOutputStreamPrivate")
	})
}

type FileOutputStreamPrivate struct {
	native uintptr
}

var filenameCompleterClassStruct *gi.Struct
var filenameCompleterClassStructOnce sync.Once

func filenameCompleterClassStructSet() {
	filenameCompleterClassStructOnce.Do(func() {
		filenameCompleterClassStruct = gi.StructNew("Gio", "FilenameCompleterClass")
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

var filterInputStreamClassStruct *gi.Struct
var filterInputStreamClassStructOnce sync.Once

func filterInputStreamClassStructSet() {
	filterInputStreamClassStructOnce.Do(func() {
		filterInputStreamClassStruct = gi.StructNew("Gio", "FilterInputStreamClass")
	})
}

type FilterInputStreamClass struct {
	native      uintptr
	ParentClass *InputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
	// UNSUPPORTED : C value '_g_reserved3' : missing Type
}

var filterOutputStreamClassStruct *gi.Struct
var filterOutputStreamClassStructOnce sync.Once

func filterOutputStreamClassStructSet() {
	filterOutputStreamClassStructOnce.Do(func() {
		filterOutputStreamClassStruct = gi.StructNew("Gio", "FilterOutputStreamClass")
	})
}

type FilterOutputStreamClass struct {
	native      uintptr
	ParentClass *OutputStreamClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
	// UNSUPPORTED : C value '_g_reserved3' : missing Type
}

var iOExtensionStruct *gi.Struct
var iOExtensionStructOnce sync.Once

func iOExtensionStructSet() {
	iOExtensionStructOnce.Do(func() {
		iOExtensionStruct = gi.StructNew("Gio", "IOExtension")
	})
}

type IOExtension struct {
	native uintptr
}

var iOExtensionGetNameFunction *gi.Function
var iOExtensionGetNameFunctionOnce sync.Once

func iOExtensionGetNameFunctionSet() {
	iOExtensionGetNameFunctionOnce.Do(func() {
		iOExtensionGetNameFunction = gi.FunctionInvokerNew("Gio", "get_name")
	})
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

var iOExtensionGetPriorityFunction *gi.Function
var iOExtensionGetPriorityFunctionOnce sync.Once

func iOExtensionGetPriorityFunctionSet() {
	iOExtensionGetPriorityFunctionOnce.Do(func() {
		iOExtensionGetPriorityFunction = gi.FunctionInvokerNew("Gio", "get_priority")
	})
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

var iOExtensionPointStruct *gi.Struct
var iOExtensionPointStructOnce sync.Once

func iOExtensionPointStructSet() {
	iOExtensionPointStructOnce.Do(func() {
		iOExtensionPointStruct = gi.StructNew("Gio", "IOExtensionPoint")
	})
}

type IOExtensionPoint struct {
	native uintptr
}

var iOExtensionPointGetExtensionByNameFunction *gi.Function
var iOExtensionPointGetExtensionByNameFunctionOnce sync.Once

func iOExtensionPointGetExtensionByNameFunctionSet() {
	iOExtensionPointGetExtensionByNameFunctionOnce.Do(func() {
		iOExtensionPointGetExtensionByNameFunction = gi.FunctionInvokerNew("Gio", "get_extension_by_name")
	})
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

var iOModuleClassStruct *gi.Struct
var iOModuleClassStructOnce sync.Once

func iOModuleClassStructSet() {
	iOModuleClassStructOnce.Do(func() {
		iOModuleClassStruct = gi.StructNew("Gio", "IOModuleClass")
	})
}

type IOModuleClass struct {
	native uintptr
}

var iOModuleScopeStruct *gi.Struct
var iOModuleScopeStructOnce sync.Once

func iOModuleScopeStructSet() {
	iOModuleScopeStructOnce.Do(func() {
		iOModuleScopeStruct = gi.StructNew("Gio", "IOModuleScope")
	})
}

type IOModuleScope struct {
	native uintptr
}

var iOModuleScopeBlockFunction *gi.Function
var iOModuleScopeBlockFunctionOnce sync.Once

func iOModuleScopeBlockFunctionSet() {
	iOModuleScopeBlockFunctionOnce.Do(func() {
		iOModuleScopeBlockFunction = gi.FunctionInvokerNew("Gio", "block")
	})
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

var iOModuleScopeFreeFunction *gi.Function
var iOModuleScopeFreeFunctionOnce sync.Once

func iOModuleScopeFreeFunctionSet() {
	iOModuleScopeFreeFunctionOnce.Do(func() {
		iOModuleScopeFreeFunction = gi.FunctionInvokerNew("Gio", "free")
	})
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

var iOSchedulerJobStruct *gi.Struct
var iOSchedulerJobStructOnce sync.Once

func iOSchedulerJobStructSet() {
	iOSchedulerJobStructOnce.Do(func() {
		iOSchedulerJobStruct = gi.StructNew("Gio", "IOSchedulerJob")
	})
}

type IOSchedulerJob struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

var iOStreamAdapterStruct *gi.Struct
var iOStreamAdapterStructOnce sync.Once

func iOStreamAdapterStructSet() {
	iOStreamAdapterStructOnce.Do(func() {
		iOStreamAdapterStruct = gi.StructNew("Gio", "IOStreamAdapter")
	})
}

type IOStreamAdapter struct {
	native uintptr
}

var iOStreamClassStruct *gi.Struct
var iOStreamClassStructOnce sync.Once

func iOStreamClassStructSet() {
	iOStreamClassStructOnce.Do(func() {
		iOStreamClassStruct = gi.StructNew("Gio", "IOStreamClass")
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

var iOStreamPrivateStruct *gi.Struct
var iOStreamPrivateStructOnce sync.Once

func iOStreamPrivateStructSet() {
	iOStreamPrivateStructOnce.Do(func() {
		iOStreamPrivateStruct = gi.StructNew("Gio", "IOStreamPrivate")
	})
}

type IOStreamPrivate struct {
	native uintptr
}

var iconIfaceStruct *gi.Struct
var iconIfaceStructOnce sync.Once

func iconIfaceStructSet() {
	iconIfaceStructOnce.Do(func() {
		iconIfaceStruct = gi.StructNew("Gio", "IconIface")
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

var inetAddressClassStruct *gi.Struct
var inetAddressClassStructOnce sync.Once

func inetAddressClassStructSet() {
	inetAddressClassStructOnce.Do(func() {
		inetAddressClassStruct = gi.StructNew("Gio", "InetAddressClass")
	})
}

type InetAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'to_string' : missing Type
	// UNSUPPORTED : C value 'to_bytes' : missing Type
}

var inetAddressMaskClassStruct *gi.Struct
var inetAddressMaskClassStructOnce sync.Once

func inetAddressMaskClassStructSet() {
	inetAddressMaskClassStructOnce.Do(func() {
		inetAddressMaskClassStruct = gi.StructNew("Gio", "InetAddressMaskClass")
	})
}

type InetAddressMaskClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var inetAddressMaskPrivateStruct *gi.Struct
var inetAddressMaskPrivateStructOnce sync.Once

func inetAddressMaskPrivateStructSet() {
	inetAddressMaskPrivateStructOnce.Do(func() {
		inetAddressMaskPrivateStruct = gi.StructNew("Gio", "InetAddressMaskPrivate")
	})
}

type InetAddressMaskPrivate struct {
	native uintptr
}

var inetAddressPrivateStruct *gi.Struct
var inetAddressPrivateStructOnce sync.Once

func inetAddressPrivateStructSet() {
	inetAddressPrivateStructOnce.Do(func() {
		inetAddressPrivateStruct = gi.StructNew("Gio", "InetAddressPrivate")
	})
}

type InetAddressPrivate struct {
	native uintptr
}

var inetSocketAddressClassStruct *gi.Struct
var inetSocketAddressClassStructOnce sync.Once

func inetSocketAddressClassStructSet() {
	inetSocketAddressClassStructOnce.Do(func() {
		inetSocketAddressClassStruct = gi.StructNew("Gio", "InetSocketAddressClass")
	})
}

type InetSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var inetSocketAddressPrivateStruct *gi.Struct
var inetSocketAddressPrivateStructOnce sync.Once

func inetSocketAddressPrivateStructSet() {
	inetSocketAddressPrivateStructOnce.Do(func() {
		inetSocketAddressPrivateStruct = gi.StructNew("Gio", "InetSocketAddressPrivate")
	})
}

type InetSocketAddressPrivate struct {
	native uintptr
}

var initableIfaceStruct *gi.Struct
var initableIfaceStructOnce sync.Once

func initableIfaceStructSet() {
	initableIfaceStructOnce.Do(func() {
		initableIfaceStruct = gi.StructNew("Gio", "InitableIface")
	})
}

type InitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'init' : missing Type
}

var inputMessageStruct *gi.Struct
var inputMessageStructOnce sync.Once

func inputMessageStructSet() {
	inputMessageStructOnce.Do(func() {
		inputMessageStruct = gi.StructNew("Gio", "InputMessage")
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

var inputStreamClassStruct *gi.Struct
var inputStreamClassStructOnce sync.Once

func inputStreamClassStructSet() {
	inputStreamClassStructOnce.Do(func() {
		inputStreamClassStruct = gi.StructNew("Gio", "InputStreamClass")
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

var inputStreamPrivateStruct *gi.Struct
var inputStreamPrivateStructOnce sync.Once

func inputStreamPrivateStructSet() {
	inputStreamPrivateStructOnce.Do(func() {
		inputStreamPrivateStruct = gi.StructNew("Gio", "InputStreamPrivate")
	})
}

type InputStreamPrivate struct {
	native uintptr
}

var inputVectorStruct *gi.Struct
var inputVectorStructOnce sync.Once

func inputVectorStructSet() {
	inputVectorStructOnce.Do(func() {
		inputVectorStruct = gi.StructNew("Gio", "InputVector")
	})
}

type InputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var listModelInterfaceStruct *gi.Struct
var listModelInterfaceStructOnce sync.Once

func listModelInterfaceStructSet() {
	listModelInterfaceStructOnce.Do(func() {
		listModelInterfaceStruct = gi.StructNew("Gio", "ListModelInterface")
	})
}

type ListModelInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_item_type' : missing Type
	// UNSUPPORTED : C value 'get_n_items' : missing Type
	// UNSUPPORTED : C value 'get_item' : missing Type
}

var listStoreClassStruct *gi.Struct
var listStoreClassStructOnce sync.Once

func listStoreClassStructSet() {
	listStoreClassStructOnce.Do(func() {
		listStoreClassStruct = gi.StructNew("Gio", "ListStoreClass")
	})
}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var loadableIconIfaceStruct *gi.Struct
var loadableIconIfaceStructOnce sync.Once

func loadableIconIfaceStructSet() {
	loadableIconIfaceStructOnce.Do(func() {
		loadableIconIfaceStruct = gi.StructNew("Gio", "LoadableIconIface")
	})
}

type LoadableIconIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'load' : missing Type
	// UNSUPPORTED : C value 'load_async' : missing Type
	// UNSUPPORTED : C value 'load_finish' : missing Type
}

var memoryInputStreamClassStruct *gi.Struct
var memoryInputStreamClassStructOnce sync.Once

func memoryInputStreamClassStructSet() {
	memoryInputStreamClassStructOnce.Do(func() {
		memoryInputStreamClassStruct = gi.StructNew("Gio", "MemoryInputStreamClass")
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

var memoryInputStreamPrivateStruct *gi.Struct
var memoryInputStreamPrivateStructOnce sync.Once

func memoryInputStreamPrivateStructSet() {
	memoryInputStreamPrivateStructOnce.Do(func() {
		memoryInputStreamPrivateStruct = gi.StructNew("Gio", "MemoryInputStreamPrivate")
	})
}

type MemoryInputStreamPrivate struct {
	native uintptr
}

var memoryOutputStreamClassStruct *gi.Struct
var memoryOutputStreamClassStructOnce sync.Once

func memoryOutputStreamClassStructSet() {
	memoryOutputStreamClassStructOnce.Do(func() {
		memoryOutputStreamClassStruct = gi.StructNew("Gio", "MemoryOutputStreamClass")
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

var memoryOutputStreamPrivateStruct *gi.Struct
var memoryOutputStreamPrivateStructOnce sync.Once

func memoryOutputStreamPrivateStructSet() {
	memoryOutputStreamPrivateStructOnce.Do(func() {
		memoryOutputStreamPrivateStruct = gi.StructNew("Gio", "MemoryOutputStreamPrivate")
	})
}

type MemoryOutputStreamPrivate struct {
	native uintptr
}

var menuAttributeIterClassStruct *gi.Struct
var menuAttributeIterClassStructOnce sync.Once

func menuAttributeIterClassStructSet() {
	menuAttributeIterClassStructOnce.Do(func() {
		menuAttributeIterClassStruct = gi.StructNew("Gio", "MenuAttributeIterClass")
	})
}

type MenuAttributeIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var menuAttributeIterPrivateStruct *gi.Struct
var menuAttributeIterPrivateStructOnce sync.Once

func menuAttributeIterPrivateStructSet() {
	menuAttributeIterPrivateStructOnce.Do(func() {
		menuAttributeIterPrivateStruct = gi.StructNew("Gio", "MenuAttributeIterPrivate")
	})
}

type MenuAttributeIterPrivate struct {
	native uintptr
}

var menuLinkIterClassStruct *gi.Struct
var menuLinkIterClassStructOnce sync.Once

func menuLinkIterClassStructSet() {
	menuLinkIterClassStructOnce.Do(func() {
		menuLinkIterClassStruct = gi.StructNew("Gio", "MenuLinkIterClass")
	})
}

type MenuLinkIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var menuLinkIterPrivateStruct *gi.Struct
var menuLinkIterPrivateStructOnce sync.Once

func menuLinkIterPrivateStructSet() {
	menuLinkIterPrivateStructOnce.Do(func() {
		menuLinkIterPrivateStruct = gi.StructNew("Gio", "MenuLinkIterPrivate")
	})
}

type MenuLinkIterPrivate struct {
	native uintptr
}

var menuModelClassStruct *gi.Struct
var menuModelClassStructOnce sync.Once

func menuModelClassStructSet() {
	menuModelClassStructOnce.Do(func() {
		menuModelClassStruct = gi.StructNew("Gio", "MenuModelClass")
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

var menuModelPrivateStruct *gi.Struct
var menuModelPrivateStructOnce sync.Once

func menuModelPrivateStructSet() {
	menuModelPrivateStructOnce.Do(func() {
		menuModelPrivateStruct = gi.StructNew("Gio", "MenuModelPrivate")
	})
}

type MenuModelPrivate struct {
	native uintptr
}

var mountIfaceStruct *gi.Struct
var mountIfaceStructOnce sync.Once

func mountIfaceStructSet() {
	mountIfaceStructOnce.Do(func() {
		mountIfaceStruct = gi.StructNew("Gio", "MountIface")
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

var mountOperationClassStruct *gi.Struct
var mountOperationClassStructOnce sync.Once

func mountOperationClassStructSet() {
	mountOperationClassStructOnce.Do(func() {
		mountOperationClassStruct = gi.StructNew("Gio", "MountOperationClass")
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

var mountOperationPrivateStruct *gi.Struct
var mountOperationPrivateStructOnce sync.Once

func mountOperationPrivateStructSet() {
	mountOperationPrivateStructOnce.Do(func() {
		mountOperationPrivateStruct = gi.StructNew("Gio", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeSocketAddressClassStruct *gi.Struct
var nativeSocketAddressClassStructOnce sync.Once

func nativeSocketAddressClassStructSet() {
	nativeSocketAddressClassStructOnce.Do(func() {
		nativeSocketAddressClassStruct = gi.StructNew("Gio", "NativeSocketAddressClass")
	})
}

type NativeSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var nativeSocketAddressPrivateStruct *gi.Struct
var nativeSocketAddressPrivateStructOnce sync.Once

func nativeSocketAddressPrivateStructSet() {
	nativeSocketAddressPrivateStructOnce.Do(func() {
		nativeSocketAddressPrivateStruct = gi.StructNew("Gio", "NativeSocketAddressPrivate")
	})
}

type NativeSocketAddressPrivate struct {
	native uintptr
}

var nativeVolumeMonitorClassStruct *gi.Struct
var nativeVolumeMonitorClassStructOnce sync.Once

func nativeVolumeMonitorClassStructSet() {
	nativeVolumeMonitorClassStructOnce.Do(func() {
		nativeVolumeMonitorClassStruct = gi.StructNew("Gio", "NativeVolumeMonitorClass")
	})
}

type NativeVolumeMonitorClass struct {
	native      uintptr
	ParentClass *VolumeMonitorClass
	// UNSUPPORTED : C value 'get_mount_for_mount_path' : missing Type
}

var networkAddressClassStruct *gi.Struct
var networkAddressClassStructOnce sync.Once

func networkAddressClassStructSet() {
	networkAddressClassStructOnce.Do(func() {
		networkAddressClassStruct = gi.StructNew("Gio", "NetworkAddressClass")
	})
}

type NetworkAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var networkAddressPrivateStruct *gi.Struct
var networkAddressPrivateStructOnce sync.Once

func networkAddressPrivateStructSet() {
	networkAddressPrivateStructOnce.Do(func() {
		networkAddressPrivateStruct = gi.StructNew("Gio", "NetworkAddressPrivate")
	})
}

type NetworkAddressPrivate struct {
	native uintptr
}

var networkMonitorInterfaceStruct *gi.Struct
var networkMonitorInterfaceStructOnce sync.Once

func networkMonitorInterfaceStructSet() {
	networkMonitorInterfaceStructOnce.Do(func() {
		networkMonitorInterfaceStruct = gi.StructNew("Gio", "NetworkMonitorInterface")
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

var networkServiceClassStruct *gi.Struct
var networkServiceClassStructOnce sync.Once

func networkServiceClassStructSet() {
	networkServiceClassStructOnce.Do(func() {
		networkServiceClassStruct = gi.StructNew("Gio", "NetworkServiceClass")
	})
}

type NetworkServiceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var networkServicePrivateStruct *gi.Struct
var networkServicePrivateStructOnce sync.Once

func networkServicePrivateStructSet() {
	networkServicePrivateStructOnce.Do(func() {
		networkServicePrivateStruct = gi.StructNew("Gio", "NetworkServicePrivate")
	})
}

type NetworkServicePrivate struct {
	native uintptr
}

var outputMessageStruct *gi.Struct
var outputMessageStructOnce sync.Once

func outputMessageStructSet() {
	outputMessageStructOnce.Do(func() {
		outputMessageStruct = gi.StructNew("Gio", "OutputMessage")
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

var outputStreamClassStruct *gi.Struct
var outputStreamClassStructOnce sync.Once

func outputStreamClassStructSet() {
	outputStreamClassStructOnce.Do(func() {
		outputStreamClassStruct = gi.StructNew("Gio", "OutputStreamClass")
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

var outputStreamPrivateStruct *gi.Struct
var outputStreamPrivateStructOnce sync.Once

func outputStreamPrivateStructSet() {
	outputStreamPrivateStructOnce.Do(func() {
		outputStreamPrivateStruct = gi.StructNew("Gio", "OutputStreamPrivate")
	})
}

type OutputStreamPrivate struct {
	native uintptr
}

var outputVectorStruct *gi.Struct
var outputVectorStructOnce sync.Once

func outputVectorStructSet() {
	outputVectorStructOnce.Do(func() {
		outputVectorStruct = gi.StructNew("Gio", "OutputVector")
	})
}

type OutputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var permissionClassStruct *gi.Struct
var permissionClassStructOnce sync.Once

func permissionClassStructSet() {
	permissionClassStructOnce.Do(func() {
		permissionClassStruct = gi.StructNew("Gio", "PermissionClass")
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

var permissionPrivateStruct *gi.Struct
var permissionPrivateStructOnce sync.Once

func permissionPrivateStructSet() {
	permissionPrivateStructOnce.Do(func() {
		permissionPrivateStruct = gi.StructNew("Gio", "PermissionPrivate")
	})
}

type PermissionPrivate struct {
	native uintptr
}

var pollableInputStreamInterfaceStruct *gi.Struct
var pollableInputStreamInterfaceStructOnce sync.Once

func pollableInputStreamInterfaceStructSet() {
	pollableInputStreamInterfaceStructOnce.Do(func() {
		pollableInputStreamInterfaceStruct = gi.StructNew("Gio", "PollableInputStreamInterface")
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

var pollableOutputStreamInterfaceStruct *gi.Struct
var pollableOutputStreamInterfaceStructOnce sync.Once

func pollableOutputStreamInterfaceStructSet() {
	pollableOutputStreamInterfaceStructOnce.Do(func() {
		pollableOutputStreamInterfaceStruct = gi.StructNew("Gio", "PollableOutputStreamInterface")
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

var proxyAddressClassStruct *gi.Struct
var proxyAddressClassStructOnce sync.Once

func proxyAddressClassStructSet() {
	proxyAddressClassStructOnce.Do(func() {
		proxyAddressClassStruct = gi.StructNew("Gio", "ProxyAddressClass")
	})
}

type ProxyAddressClass struct {
	native      uintptr
	ParentClass *InetSocketAddressClass
}

var proxyAddressEnumeratorClassStruct *gi.Struct
var proxyAddressEnumeratorClassStructOnce sync.Once

func proxyAddressEnumeratorClassStructSet() {
	proxyAddressEnumeratorClassStructOnce.Do(func() {
		proxyAddressEnumeratorClassStruct = gi.StructNew("Gio", "ProxyAddressEnumeratorClass")
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

var proxyAddressEnumeratorPrivateStruct *gi.Struct
var proxyAddressEnumeratorPrivateStructOnce sync.Once

func proxyAddressEnumeratorPrivateStructSet() {
	proxyAddressEnumeratorPrivateStructOnce.Do(func() {
		proxyAddressEnumeratorPrivateStruct = gi.StructNew("Gio", "ProxyAddressEnumeratorPrivate")
	})
}

type ProxyAddressEnumeratorPrivate struct {
	native uintptr
}

var proxyAddressPrivateStruct *gi.Struct
var proxyAddressPrivateStructOnce sync.Once

func proxyAddressPrivateStructSet() {
	proxyAddressPrivateStructOnce.Do(func() {
		proxyAddressPrivateStruct = gi.StructNew("Gio", "ProxyAddressPrivate")
	})
}

type ProxyAddressPrivate struct {
	native uintptr
}

var proxyInterfaceStruct *gi.Struct
var proxyInterfaceStructOnce sync.Once

func proxyInterfaceStructSet() {
	proxyInterfaceStructOnce.Do(func() {
		proxyInterfaceStruct = gi.StructNew("Gio", "ProxyInterface")
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

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStructOnce sync.Once

func proxyResolverInterfaceStructSet() {
	proxyResolverInterfaceStructOnce.Do(func() {
		proxyResolverInterfaceStruct = gi.StructNew("Gio", "ProxyResolverInterface")
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

var remoteActionGroupInterfaceStruct *gi.Struct
var remoteActionGroupInterfaceStructOnce sync.Once

func remoteActionGroupInterfaceStructSet() {
	remoteActionGroupInterfaceStructOnce.Do(func() {
		remoteActionGroupInterfaceStruct = gi.StructNew("Gio", "RemoteActionGroupInterface")
	})
}

type RemoteActionGroupInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'activate_action_full' : missing Type
	// UNSUPPORTED : C value 'change_action_state_full' : missing Type
}

var resolverClassStruct *gi.Struct
var resolverClassStructOnce sync.Once

func resolverClassStructSet() {
	resolverClassStructOnce.Do(func() {
		resolverClassStruct = gi.StructNew("Gio", "ResolverClass")
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

var resolverPrivateStruct *gi.Struct
var resolverPrivateStructOnce sync.Once

func resolverPrivateStructSet() {
	resolverPrivateStructOnce.Do(func() {
		resolverPrivateStruct = gi.StructNew("Gio", "ResolverPrivate")
	})
}

type ResolverPrivate struct {
	native uintptr
}

var resourceStruct *gi.Struct
var resourceStructOnce sync.Once

func resourceStructSet() {
	resourceStructOnce.Do(func() {
		resourceStruct = gi.StructNew("Gio", "Resource")
	})
}

type Resource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_resource_new_from_data' : parameter 'data' of type 'GLib.Bytes' not supported

var resourceRegisterFunction *gi.Function
var resourceRegisterFunctionOnce sync.Once

func resourceRegisterFunctionSet() {
	resourceRegisterFunctionOnce.Do(func() {
		resourceRegisterFunction = gi.FunctionInvokerNew("Gio", "_register")
	})
}

var registerResourceInvoker *gi.Function

// Register is a representation of the C type g_resources_register.
func (recv *Resource) Register() {
	if registerResourceInvoker == nil {
		registerResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "_register")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	registerResourceInvoker.Invoke(inArgs[:], nil)

}

var resourceUnregisterFunction *gi.Function
var resourceUnregisterFunctionOnce sync.Once

func resourceUnregisterFunctionSet() {
	resourceUnregisterFunctionOnce.Do(func() {
		resourceUnregisterFunction = gi.FunctionInvokerNew("Gio", "_unregister")
	})
}

var unregisterResourceInvoker *gi.Function

// Unregister is a representation of the C type g_resources_unregister.
func (recv *Resource) Unregister() {
	if unregisterResourceInvoker == nil {
		unregisterResourceInvoker = gi.StructFunctionInvokerNew("Gio", "Resource", "_unregister")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unregisterResourceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_resource_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var resourceRefFunction *gi.Function
var resourceRefFunctionOnce sync.Once

func resourceRefFunctionSet() {
	resourceRefFunctionOnce.Do(func() {
		resourceRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
}

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

var resourceUnrefFunction *gi.Function
var resourceUnrefFunctionOnce sync.Once

func resourceUnrefFunctionSet() {
	resourceUnrefFunctionOnce.Do(func() {
		resourceUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var seekableIfaceStruct *gi.Struct
var seekableIfaceStructOnce sync.Once

func seekableIfaceStructSet() {
	seekableIfaceStructOnce.Do(func() {
		seekableIfaceStruct = gi.StructNew("Gio", "SeekableIface")
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

var settingsBackendClassStruct *gi.Struct
var settingsBackendClassStructOnce sync.Once

func settingsBackendClassStructSet() {
	settingsBackendClassStructOnce.Do(func() {
		settingsBackendClassStruct = gi.StructNew("Gio", "SettingsBackendClass")
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

var settingsBackendPrivateStruct *gi.Struct
var settingsBackendPrivateStructOnce sync.Once

func settingsBackendPrivateStructSet() {
	settingsBackendPrivateStructOnce.Do(func() {
		settingsBackendPrivateStruct = gi.StructNew("Gio", "SettingsBackendPrivate")
	})
}

type SettingsBackendPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStructOnce sync.Once

func settingsClassStructSet() {
	settingsClassStructOnce.Do(func() {
		settingsClassStruct = gi.StructNew("Gio", "SettingsClass")
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

var settingsPrivateStruct *gi.Struct
var settingsPrivateStructOnce sync.Once

func settingsPrivateStructSet() {
	settingsPrivateStructOnce.Do(func() {
		settingsPrivateStruct = gi.StructNew("Gio", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var settingsSchemaStruct *gi.Struct
var settingsSchemaStructOnce sync.Once

func settingsSchemaStructSet() {
	settingsSchemaStructOnce.Do(func() {
		settingsSchemaStruct = gi.StructNew("Gio", "SettingsSchema")
	})
}

type SettingsSchema struct {
	native uintptr
}

var settingsSchemaGetIdFunction *gi.Function
var settingsSchemaGetIdFunctionOnce sync.Once

func settingsSchemaGetIdFunctionSet() {
	settingsSchemaGetIdFunctionOnce.Do(func() {
		settingsSchemaGetIdFunction = gi.FunctionInvokerNew("Gio", "get_id")
	})
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

var settingsSchemaGetKeyFunction *gi.Function
var settingsSchemaGetKeyFunctionOnce sync.Once

func settingsSchemaGetKeyFunctionSet() {
	settingsSchemaGetKeyFunctionOnce.Do(func() {
		settingsSchemaGetKeyFunction = gi.FunctionInvokerNew("Gio", "get_key")
	})
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

var settingsSchemaGetPathFunction *gi.Function
var settingsSchemaGetPathFunctionOnce sync.Once

func settingsSchemaGetPathFunctionSet() {
	settingsSchemaGetPathFunctionOnce.Do(func() {
		settingsSchemaGetPathFunction = gi.FunctionInvokerNew("Gio", "get_path")
	})
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

var settingsSchemaListChildrenFunction *gi.Function
var settingsSchemaListChildrenFunctionOnce sync.Once

func settingsSchemaListChildrenFunctionSet() {
	settingsSchemaListChildrenFunctionOnce.Do(func() {
		settingsSchemaListChildrenFunction = gi.FunctionInvokerNew("Gio", "list_children")
	})
}

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

var settingsSchemaListKeysFunction *gi.Function
var settingsSchemaListKeysFunctionOnce sync.Once

func settingsSchemaListKeysFunctionSet() {
	settingsSchemaListKeysFunctionOnce.Do(func() {
		settingsSchemaListKeysFunction = gi.FunctionInvokerNew("Gio", "list_keys")
	})
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

var settingsSchemaRefFunction *gi.Function
var settingsSchemaRefFunctionOnce sync.Once

func settingsSchemaRefFunctionSet() {
	settingsSchemaRefFunctionOnce.Do(func() {
		settingsSchemaRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
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

var settingsSchemaUnrefFunction *gi.Function
var settingsSchemaUnrefFunctionOnce sync.Once

func settingsSchemaUnrefFunctionSet() {
	settingsSchemaUnrefFunctionOnce.Do(func() {
		settingsSchemaUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var settingsSchemaKeyStruct *gi.Struct
var settingsSchemaKeyStructOnce sync.Once

func settingsSchemaKeyStructSet() {
	settingsSchemaKeyStructOnce.Do(func() {
		settingsSchemaKeyStruct = gi.StructNew("Gio", "SettingsSchemaKey")
	})
}

type SettingsSchemaKey struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_default_value' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetDescriptionFunction *gi.Function
var settingsSchemaKeyGetDescriptionFunctionOnce sync.Once

func settingsSchemaKeyGetDescriptionFunctionSet() {
	settingsSchemaKeyGetDescriptionFunctionOnce.Do(func() {
		settingsSchemaKeyGetDescriptionFunction = gi.FunctionInvokerNew("Gio", "get_description")
	})
}

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

var settingsSchemaKeyGetNameFunction *gi.Function
var settingsSchemaKeyGetNameFunctionOnce sync.Once

func settingsSchemaKeyGetNameFunctionSet() {
	settingsSchemaKeyGetNameFunctionOnce.Do(func() {
		settingsSchemaKeyGetNameFunction = gi.FunctionInvokerNew("Gio", "get_name")
	})
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

var settingsSchemaKeyGetSummaryFunction *gi.Function
var settingsSchemaKeyGetSummaryFunctionOnce sync.Once

func settingsSchemaKeyGetSummaryFunctionSet() {
	settingsSchemaKeyGetSummaryFunctionOnce.Do(func() {
		settingsSchemaKeyGetSummaryFunction = gi.FunctionInvokerNew("Gio", "get_summary")
	})
}

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

var settingsSchemaKeyRefFunction *gi.Function
var settingsSchemaKeyRefFunctionOnce sync.Once

func settingsSchemaKeyRefFunctionSet() {
	settingsSchemaKeyRefFunctionOnce.Do(func() {
		settingsSchemaKeyRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
}

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

var settingsSchemaKeyUnrefFunction *gi.Function
var settingsSchemaKeyUnrefFunctionOnce sync.Once

func settingsSchemaKeyUnrefFunctionSet() {
	settingsSchemaKeyUnrefFunctionOnce.Do(func() {
		settingsSchemaKeyUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var settingsSchemaSourceStruct *gi.Struct
var settingsSchemaSourceStructOnce sync.Once

func settingsSchemaSourceStructSet() {
	settingsSchemaSourceStructOnce.Do(func() {
		settingsSchemaSourceStruct = gi.StructNew("Gio", "SettingsSchemaSource")
	})
}

type SettingsSchemaSource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_source_new_from_directory' : parameter 'directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_list_schemas' : parameter 'recursive' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_lookup' : parameter 'recursive' of type 'gboolean' not supported

var settingsSchemaSourceRefFunction *gi.Function
var settingsSchemaSourceRefFunctionOnce sync.Once

func settingsSchemaSourceRefFunctionSet() {
	settingsSchemaSourceRefFunctionOnce.Do(func() {
		settingsSchemaSourceRefFunction = gi.FunctionInvokerNew("Gio", "ref")
	})
}

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

var settingsSchemaSourceUnrefFunction *gi.Function
var settingsSchemaSourceUnrefFunctionOnce sync.Once

func settingsSchemaSourceUnrefFunctionSet() {
	settingsSchemaSourceUnrefFunctionOnce.Do(func() {
		settingsSchemaSourceUnrefFunction = gi.FunctionInvokerNew("Gio", "unref")
	})
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

var simpleActionGroupClassStruct *gi.Struct
var simpleActionGroupClassStructOnce sync.Once

func simpleActionGroupClassStructSet() {
	simpleActionGroupClassStructOnce.Do(func() {
		simpleActionGroupClassStruct = gi.StructNew("Gio", "SimpleActionGroupClass")
	})
}

type SimpleActionGroupClass struct {
	native uintptr
}

var simpleActionGroupPrivateStruct *gi.Struct
var simpleActionGroupPrivateStructOnce sync.Once

func simpleActionGroupPrivateStructSet() {
	simpleActionGroupPrivateStructOnce.Do(func() {
		simpleActionGroupPrivateStruct = gi.StructNew("Gio", "SimpleActionGroupPrivate")
	})
}

type SimpleActionGroupPrivate struct {
	native uintptr
}

var simpleAsyncResultClassStruct *gi.Struct
var simpleAsyncResultClassStructOnce sync.Once

func simpleAsyncResultClassStructSet() {
	simpleAsyncResultClassStructOnce.Do(func() {
		simpleAsyncResultClassStruct = gi.StructNew("Gio", "SimpleAsyncResultClass")
	})
}

type SimpleAsyncResultClass struct {
	native uintptr
}

var simpleProxyResolverClassStruct *gi.Struct
var simpleProxyResolverClassStructOnce sync.Once

func simpleProxyResolverClassStructSet() {
	simpleProxyResolverClassStructOnce.Do(func() {
		simpleProxyResolverClassStruct = gi.StructNew("Gio", "SimpleProxyResolverClass")
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

var simpleProxyResolverPrivateStruct *gi.Struct
var simpleProxyResolverPrivateStructOnce sync.Once

func simpleProxyResolverPrivateStructSet() {
	simpleProxyResolverPrivateStructOnce.Do(func() {
		simpleProxyResolverPrivateStruct = gi.StructNew("Gio", "SimpleProxyResolverPrivate")
	})
}

type SimpleProxyResolverPrivate struct {
	native uintptr
}

var socketAddressClassStruct *gi.Struct
var socketAddressClassStructOnce sync.Once

func socketAddressClassStructSet() {
	socketAddressClassStructOnce.Do(func() {
		socketAddressClassStruct = gi.StructNew("Gio", "SocketAddressClass")
	})
}

type SocketAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_family' : missing Type
	// UNSUPPORTED : C value 'get_native_size' : missing Type
	// UNSUPPORTED : C value 'to_native' : missing Type
}

var socketAddressEnumeratorClassStruct *gi.Struct
var socketAddressEnumeratorClassStructOnce sync.Once

func socketAddressEnumeratorClassStructSet() {
	socketAddressEnumeratorClassStructOnce.Do(func() {
		socketAddressEnumeratorClassStruct = gi.StructNew("Gio", "SocketAddressEnumeratorClass")
	})
}

type SocketAddressEnumeratorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'next' : missing Type
	// UNSUPPORTED : C value 'next_async' : missing Type
	// UNSUPPORTED : C value 'next_finish' : missing Type
}

var socketClassStruct *gi.Struct
var socketClassStructOnce sync.Once

func socketClassStructSet() {
	socketClassStructOnce.Do(func() {
		socketClassStruct = gi.StructNew("Gio", "SocketClass")
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

var socketClientClassStruct *gi.Struct
var socketClientClassStructOnce sync.Once

func socketClientClassStructSet() {
	socketClientClassStructOnce.Do(func() {
		socketClientClassStruct = gi.StructNew("Gio", "SocketClientClass")
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

var socketClientPrivateStruct *gi.Struct
var socketClientPrivateStructOnce sync.Once

func socketClientPrivateStructSet() {
	socketClientPrivateStructOnce.Do(func() {
		socketClientPrivateStruct = gi.StructNew("Gio", "SocketClientPrivate")
	})
}

type SocketClientPrivate struct {
	native uintptr
}

var socketConnectableIfaceStruct *gi.Struct
var socketConnectableIfaceStructOnce sync.Once

func socketConnectableIfaceStructSet() {
	socketConnectableIfaceStructOnce.Do(func() {
		socketConnectableIfaceStruct = gi.StructNew("Gio", "SocketConnectableIface")
	})
}

type SocketConnectableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'enumerate' : missing Type
	// UNSUPPORTED : C value 'proxy_enumerate' : missing Type
	// UNSUPPORTED : C value 'to_string' : missing Type
}

var socketConnectionClassStruct *gi.Struct
var socketConnectionClassStructOnce sync.Once

func socketConnectionClassStructSet() {
	socketConnectionClassStructOnce.Do(func() {
		socketConnectionClassStruct = gi.StructNew("Gio", "SocketConnectionClass")
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

var socketConnectionPrivateStruct *gi.Struct
var socketConnectionPrivateStructOnce sync.Once

func socketConnectionPrivateStructSet() {
	socketConnectionPrivateStructOnce.Do(func() {
		socketConnectionPrivateStruct = gi.StructNew("Gio", "SocketConnectionPrivate")
	})
}

type SocketConnectionPrivate struct {
	native uintptr
}

var socketControlMessageClassStruct *gi.Struct
var socketControlMessageClassStructOnce sync.Once

func socketControlMessageClassStructSet() {
	socketControlMessageClassStructOnce.Do(func() {
		socketControlMessageClassStruct = gi.StructNew("Gio", "SocketControlMessageClass")
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

var socketControlMessagePrivateStruct *gi.Struct
var socketControlMessagePrivateStructOnce sync.Once

func socketControlMessagePrivateStructSet() {
	socketControlMessagePrivateStructOnce.Do(func() {
		socketControlMessagePrivateStruct = gi.StructNew("Gio", "SocketControlMessagePrivate")
	})
}

type SocketControlMessagePrivate struct {
	native uintptr
}

var socketListenerClassStruct *gi.Struct
var socketListenerClassStructOnce sync.Once

func socketListenerClassStructSet() {
	socketListenerClassStructOnce.Do(func() {
		socketListenerClassStruct = gi.StructNew("Gio", "SocketListenerClass")
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

var socketListenerPrivateStruct *gi.Struct
var socketListenerPrivateStructOnce sync.Once

func socketListenerPrivateStructSet() {
	socketListenerPrivateStructOnce.Do(func() {
		socketListenerPrivateStruct = gi.StructNew("Gio", "SocketListenerPrivate")
	})
}

type SocketListenerPrivate struct {
	native uintptr
}

var socketPrivateStruct *gi.Struct
var socketPrivateStructOnce sync.Once

func socketPrivateStructSet() {
	socketPrivateStructOnce.Do(func() {
		socketPrivateStruct = gi.StructNew("Gio", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var socketServiceClassStruct *gi.Struct
var socketServiceClassStructOnce sync.Once

func socketServiceClassStructSet() {
	socketServiceClassStructOnce.Do(func() {
		socketServiceClassStruct = gi.StructNew("Gio", "SocketServiceClass")
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

var socketServicePrivateStruct *gi.Struct
var socketServicePrivateStructOnce sync.Once

func socketServicePrivateStructSet() {
	socketServicePrivateStructOnce.Do(func() {
		socketServicePrivateStruct = gi.StructNew("Gio", "SocketServicePrivate")
	})
}

type SocketServicePrivate struct {
	native uintptr
}

var srvTargetStruct *gi.Struct
var srvTargetStructOnce sync.Once

func srvTargetStructSet() {
	srvTargetStructOnce.Do(func() {
		srvTargetStruct = gi.StructNew("Gio", "SrvTarget")
	})
}

type SrvTarget struct {
	native uintptr
}

var srvTargetNewFunction *gi.Function
var srvTargetNewFunctionOnce sync.Once

func srvTargetNewFunctionSet() {
	srvTargetNewFunctionOnce.Do(func() {
		srvTargetNewFunction = gi.FunctionInvokerNew("Gio", "new")
	})
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

var srvTargetCopyFunction *gi.Function
var srvTargetCopyFunctionOnce sync.Once

func srvTargetCopyFunctionSet() {
	srvTargetCopyFunctionOnce.Do(func() {
		srvTargetCopyFunction = gi.FunctionInvokerNew("Gio", "copy")
	})
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

var srvTargetFreeFunction *gi.Function
var srvTargetFreeFunctionOnce sync.Once

func srvTargetFreeFunctionSet() {
	srvTargetFreeFunctionOnce.Do(func() {
		srvTargetFreeFunction = gi.FunctionInvokerNew("Gio", "free")
	})
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

var srvTargetGetHostnameFunction *gi.Function
var srvTargetGetHostnameFunctionOnce sync.Once

func srvTargetGetHostnameFunctionSet() {
	srvTargetGetHostnameFunctionOnce.Do(func() {
		srvTargetGetHostnameFunction = gi.FunctionInvokerNew("Gio", "get_hostname")
	})
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

var srvTargetGetPortFunction *gi.Function
var srvTargetGetPortFunctionOnce sync.Once

func srvTargetGetPortFunctionSet() {
	srvTargetGetPortFunctionOnce.Do(func() {
		srvTargetGetPortFunction = gi.FunctionInvokerNew("Gio", "get_port")
	})
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

var srvTargetGetPriorityFunction *gi.Function
var srvTargetGetPriorityFunctionOnce sync.Once

func srvTargetGetPriorityFunctionSet() {
	srvTargetGetPriorityFunctionOnce.Do(func() {
		srvTargetGetPriorityFunction = gi.FunctionInvokerNew("Gio", "get_priority")
	})
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

var srvTargetGetWeightFunction *gi.Function
var srvTargetGetWeightFunctionOnce sync.Once

func srvTargetGetWeightFunctionSet() {
	srvTargetGetWeightFunctionOnce.Do(func() {
		srvTargetGetWeightFunction = gi.FunctionInvokerNew("Gio", "get_weight")
	})
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

var staticResourceStruct *gi.Struct
var staticResourceStructOnce sync.Once

func staticResourceStructSet() {
	staticResourceStructOnce.Do(func() {
		staticResourceStruct = gi.StructNew("Gio", "StaticResource")
	})
}

type StaticResource struct {
	native uintptr
}

var staticResourceFiniFunction *gi.Function
var staticResourceFiniFunctionOnce sync.Once

func staticResourceFiniFunctionSet() {
	staticResourceFiniFunctionOnce.Do(func() {
		staticResourceFiniFunction = gi.FunctionInvokerNew("Gio", "fini")
	})
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

var staticResourceGetResourceFunction *gi.Function
var staticResourceGetResourceFunctionOnce sync.Once

func staticResourceGetResourceFunctionSet() {
	staticResourceGetResourceFunctionOnce.Do(func() {
		staticResourceGetResourceFunction = gi.FunctionInvokerNew("Gio", "get_resource")
	})
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

var staticResourceInitFunction *gi.Function
var staticResourceInitFunctionOnce sync.Once

func staticResourceInitFunctionSet() {
	staticResourceInitFunctionOnce.Do(func() {
		staticResourceInitFunction = gi.FunctionInvokerNew("Gio", "init")
	})
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

var taskClassStruct *gi.Struct
var taskClassStructOnce sync.Once

func taskClassStructSet() {
	taskClassStructOnce.Do(func() {
		taskClassStruct = gi.StructNew("Gio", "TaskClass")
	})
}

type TaskClass struct {
	native uintptr
}

var tcpConnectionClassStruct *gi.Struct
var tcpConnectionClassStructOnce sync.Once

func tcpConnectionClassStructSet() {
	tcpConnectionClassStructOnce.Do(func() {
		tcpConnectionClassStruct = gi.StructNew("Gio", "TcpConnectionClass")
	})
}

type TcpConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var tcpConnectionPrivateStruct *gi.Struct
var tcpConnectionPrivateStructOnce sync.Once

func tcpConnectionPrivateStructSet() {
	tcpConnectionPrivateStructOnce.Do(func() {
		tcpConnectionPrivateStruct = gi.StructNew("Gio", "TcpConnectionPrivate")
	})
}

type TcpConnectionPrivate struct {
	native uintptr
}

var tcpWrapperConnectionClassStruct *gi.Struct
var tcpWrapperConnectionClassStructOnce sync.Once

func tcpWrapperConnectionClassStructSet() {
	tcpWrapperConnectionClassStructOnce.Do(func() {
		tcpWrapperConnectionClassStruct = gi.StructNew("Gio", "TcpWrapperConnectionClass")
	})
}

type TcpWrapperConnectionClass struct {
	native      uintptr
	ParentClass *TcpConnectionClass
}

var tcpWrapperConnectionPrivateStruct *gi.Struct
var tcpWrapperConnectionPrivateStructOnce sync.Once

func tcpWrapperConnectionPrivateStructSet() {
	tcpWrapperConnectionPrivateStructOnce.Do(func() {
		tcpWrapperConnectionPrivateStruct = gi.StructNew("Gio", "TcpWrapperConnectionPrivate")
	})
}

type TcpWrapperConnectionPrivate struct {
	native uintptr
}

var themedIconClassStruct *gi.Struct
var themedIconClassStructOnce sync.Once

func themedIconClassStructSet() {
	themedIconClassStructOnce.Do(func() {
		themedIconClassStruct = gi.StructNew("Gio", "ThemedIconClass")
	})
}

type ThemedIconClass struct {
	native uintptr
}

var threadedSocketServiceClassStruct *gi.Struct
var threadedSocketServiceClassStructOnce sync.Once

func threadedSocketServiceClassStructSet() {
	threadedSocketServiceClassStructOnce.Do(func() {
		threadedSocketServiceClassStruct = gi.StructNew("Gio", "ThreadedSocketServiceClass")
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

var threadedSocketServicePrivateStruct *gi.Struct
var threadedSocketServicePrivateStructOnce sync.Once

func threadedSocketServicePrivateStructSet() {
	threadedSocketServicePrivateStructOnce.Do(func() {
		threadedSocketServicePrivateStruct = gi.StructNew("Gio", "ThreadedSocketServicePrivate")
	})
}

type ThreadedSocketServicePrivate struct {
	native uintptr
}

var tlsBackendInterfaceStruct *gi.Struct
var tlsBackendInterfaceStructOnce sync.Once

func tlsBackendInterfaceStructSet() {
	tlsBackendInterfaceStructOnce.Do(func() {
		tlsBackendInterfaceStruct = gi.StructNew("Gio", "TlsBackendInterface")
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

var tlsCertificateClassStruct *gi.Struct
var tlsCertificateClassStructOnce sync.Once

func tlsCertificateClassStructSet() {
	tlsCertificateClassStructOnce.Do(func() {
		tlsCertificateClassStruct = gi.StructNew("Gio", "TlsCertificateClass")
	})
}

type TlsCertificateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'verify' : missing Type
}

var tlsCertificatePrivateStruct *gi.Struct
var tlsCertificatePrivateStructOnce sync.Once

func tlsCertificatePrivateStructSet() {
	tlsCertificatePrivateStructOnce.Do(func() {
		tlsCertificatePrivateStruct = gi.StructNew("Gio", "TlsCertificatePrivate")
	})
}

type TlsCertificatePrivate struct {
	native uintptr
}

var tlsClientConnectionInterfaceStruct *gi.Struct
var tlsClientConnectionInterfaceStructOnce sync.Once

func tlsClientConnectionInterfaceStructSet() {
	tlsClientConnectionInterfaceStructOnce.Do(func() {
		tlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "TlsClientConnectionInterface")
	})
}

type TlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'copy_session_state' : missing Type
}

var tlsConnectionClassStruct *gi.Struct
var tlsConnectionClassStructOnce sync.Once

func tlsConnectionClassStructSet() {
	tlsConnectionClassStructOnce.Do(func() {
		tlsConnectionClassStruct = gi.StructNew("Gio", "TlsConnectionClass")
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

var tlsConnectionPrivateStruct *gi.Struct
var tlsConnectionPrivateStructOnce sync.Once

func tlsConnectionPrivateStructSet() {
	tlsConnectionPrivateStructOnce.Do(func() {
		tlsConnectionPrivateStruct = gi.StructNew("Gio", "TlsConnectionPrivate")
	})
}

type TlsConnectionPrivate struct {
	native uintptr
}

var tlsDatabaseClassStruct *gi.Struct
var tlsDatabaseClassStructOnce sync.Once

func tlsDatabaseClassStructSet() {
	tlsDatabaseClassStructOnce.Do(func() {
		tlsDatabaseClassStruct = gi.StructNew("Gio", "TlsDatabaseClass")
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

var tlsDatabasePrivateStruct *gi.Struct
var tlsDatabasePrivateStructOnce sync.Once

func tlsDatabasePrivateStructSet() {
	tlsDatabasePrivateStructOnce.Do(func() {
		tlsDatabasePrivateStruct = gi.StructNew("Gio", "TlsDatabasePrivate")
	})
}

type TlsDatabasePrivate struct {
	native uintptr
}

var tlsFileDatabaseInterfaceStruct *gi.Struct
var tlsFileDatabaseInterfaceStructOnce sync.Once

func tlsFileDatabaseInterfaceStructSet() {
	tlsFileDatabaseInterfaceStructOnce.Do(func() {
		tlsFileDatabaseInterfaceStruct = gi.StructNew("Gio", "TlsFileDatabaseInterface")
	})
}

type TlsFileDatabaseInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var tlsInteractionClassStruct *gi.Struct
var tlsInteractionClassStructOnce sync.Once

func tlsInteractionClassStructSet() {
	tlsInteractionClassStructOnce.Do(func() {
		tlsInteractionClassStruct = gi.StructNew("Gio", "TlsInteractionClass")
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

var tlsInteractionPrivateStruct *gi.Struct
var tlsInteractionPrivateStructOnce sync.Once

func tlsInteractionPrivateStructSet() {
	tlsInteractionPrivateStructOnce.Do(func() {
		tlsInteractionPrivateStruct = gi.StructNew("Gio", "TlsInteractionPrivate")
	})
}

type TlsInteractionPrivate struct {
	native uintptr
}

var tlsPasswordClassStruct *gi.Struct
var tlsPasswordClassStructOnce sync.Once

func tlsPasswordClassStructSet() {
	tlsPasswordClassStructOnce.Do(func() {
		tlsPasswordClassStruct = gi.StructNew("Gio", "TlsPasswordClass")
	})
}

type TlsPasswordClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_value' : missing Type
	// UNSUPPORTED : C value 'set_value' : missing Type
	// UNSUPPORTED : C value 'get_default_warning' : missing Type
}

var tlsPasswordPrivateStruct *gi.Struct
var tlsPasswordPrivateStructOnce sync.Once

func tlsPasswordPrivateStructSet() {
	tlsPasswordPrivateStructOnce.Do(func() {
		tlsPasswordPrivateStruct = gi.StructNew("Gio", "TlsPasswordPrivate")
	})
}

type TlsPasswordPrivate struct {
	native uintptr
}

var tlsServerConnectionInterfaceStruct *gi.Struct
var tlsServerConnectionInterfaceStructOnce sync.Once

func tlsServerConnectionInterfaceStructSet() {
	tlsServerConnectionInterfaceStructOnce.Do(func() {
		tlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "TlsServerConnectionInterface")
	})
}

type TlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var unixConnectionClassStruct *gi.Struct
var unixConnectionClassStructOnce sync.Once

func unixConnectionClassStructSet() {
	unixConnectionClassStructOnce.Do(func() {
		unixConnectionClassStruct = gi.StructNew("Gio", "UnixConnectionClass")
	})
}

type UnixConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var unixConnectionPrivateStruct *gi.Struct
var unixConnectionPrivateStructOnce sync.Once

func unixConnectionPrivateStructSet() {
	unixConnectionPrivateStructOnce.Do(func() {
		unixConnectionPrivateStruct = gi.StructNew("Gio", "UnixConnectionPrivate")
	})
}

type UnixConnectionPrivate struct {
	native uintptr
}

var unixCredentialsMessageClassStruct *gi.Struct
var unixCredentialsMessageClassStructOnce sync.Once

func unixCredentialsMessageClassStructSet() {
	unixCredentialsMessageClassStructOnce.Do(func() {
		unixCredentialsMessageClassStruct = gi.StructNew("Gio", "UnixCredentialsMessageClass")
	})
}

type UnixCredentialsMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var unixCredentialsMessagePrivateStruct *gi.Struct
var unixCredentialsMessagePrivateStructOnce sync.Once

func unixCredentialsMessagePrivateStructSet() {
	unixCredentialsMessagePrivateStructOnce.Do(func() {
		unixCredentialsMessagePrivateStruct = gi.StructNew("Gio", "UnixCredentialsMessagePrivate")
	})
}

type UnixCredentialsMessagePrivate struct {
	native uintptr
}

var unixFDListClassStruct *gi.Struct
var unixFDListClassStructOnce sync.Once

func unixFDListClassStructSet() {
	unixFDListClassStructOnce.Do(func() {
		unixFDListClassStruct = gi.StructNew("Gio", "UnixFDListClass")
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

var unixFDListPrivateStruct *gi.Struct
var unixFDListPrivateStructOnce sync.Once

func unixFDListPrivateStructSet() {
	unixFDListPrivateStructOnce.Do(func() {
		unixFDListPrivateStruct = gi.StructNew("Gio", "UnixFDListPrivate")
	})
}

type UnixFDListPrivate struct {
	native uintptr
}

var unixFDMessageClassStruct *gi.Struct
var unixFDMessageClassStructOnce sync.Once

func unixFDMessageClassStructSet() {
	unixFDMessageClassStructOnce.Do(func() {
		unixFDMessageClassStruct = gi.StructNew("Gio", "UnixFDMessageClass")
	})
}

type UnixFDMessageClass struct {
	native      uintptr
	ParentClass *SocketControlMessageClass
	// UNSUPPORTED : C value '_g_reserved1' : missing Type
	// UNSUPPORTED : C value '_g_reserved2' : missing Type
}

var unixFDMessagePrivateStruct *gi.Struct
var unixFDMessagePrivateStructOnce sync.Once

func unixFDMessagePrivateStructSet() {
	unixFDMessagePrivateStructOnce.Do(func() {
		unixFDMessagePrivateStruct = gi.StructNew("Gio", "UnixFDMessagePrivate")
	})
}

type UnixFDMessagePrivate struct {
	native uintptr
}

var unixInputStreamClassStruct *gi.Struct
var unixInputStreamClassStructOnce sync.Once

func unixInputStreamClassStructSet() {
	unixInputStreamClassStructOnce.Do(func() {
		unixInputStreamClassStruct = gi.StructNew("Gio", "UnixInputStreamClass")
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

var unixInputStreamPrivateStruct *gi.Struct
var unixInputStreamPrivateStructOnce sync.Once

func unixInputStreamPrivateStructSet() {
	unixInputStreamPrivateStructOnce.Do(func() {
		unixInputStreamPrivateStruct = gi.StructNew("Gio", "UnixInputStreamPrivate")
	})
}

type UnixInputStreamPrivate struct {
	native uintptr
}

var unixMountEntryStruct *gi.Struct
var unixMountEntryStructOnce sync.Once

func unixMountEntryStructSet() {
	unixMountEntryStructOnce.Do(func() {
		unixMountEntryStruct = gi.StructNew("Gio", "UnixMountEntry")
	})
}

type UnixMountEntry struct {
	native uintptr
}

var unixMountMonitorClassStruct *gi.Struct
var unixMountMonitorClassStructOnce sync.Once

func unixMountMonitorClassStructSet() {
	unixMountMonitorClassStructOnce.Do(func() {
		unixMountMonitorClassStruct = gi.StructNew("Gio", "UnixMountMonitorClass")
	})
}

type UnixMountMonitorClass struct {
	native uintptr
}

var unixMountPointStruct *gi.Struct
var unixMountPointStructOnce sync.Once

func unixMountPointStructSet() {
	unixMountPointStructOnce.Do(func() {
		unixMountPointStruct = gi.StructNew("Gio", "UnixMountPoint")
	})
}

type UnixMountPoint struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_unix_mount_point_compare' : parameter 'mount2' of type 'UnixMountPoint' not supported

var unixMountPointCopyFunction *gi.Function
var unixMountPointCopyFunctionOnce sync.Once

func unixMountPointCopyFunctionSet() {
	unixMountPointCopyFunctionOnce.Do(func() {
		unixMountPointCopyFunction = gi.FunctionInvokerNew("Gio", "copy")
	})
}

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

var unixMountPointFreeFunction *gi.Function
var unixMountPointFreeFunctionOnce sync.Once

func unixMountPointFreeFunctionSet() {
	unixMountPointFreeFunctionOnce.Do(func() {
		unixMountPointFreeFunction = gi.FunctionInvokerNew("Gio", "free")
	})
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

var unixMountPointGetFsTypeFunction *gi.Function
var unixMountPointGetFsTypeFunctionOnce sync.Once

func unixMountPointGetFsTypeFunctionSet() {
	unixMountPointGetFsTypeFunctionOnce.Do(func() {
		unixMountPointGetFsTypeFunction = gi.FunctionInvokerNew("Gio", "get_fs_type")
	})
}

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

var unixMountPointGetOptionsFunction *gi.Function
var unixMountPointGetOptionsFunctionOnce sync.Once

func unixMountPointGetOptionsFunctionSet() {
	unixMountPointGetOptionsFunctionOnce.Do(func() {
		unixMountPointGetOptionsFunction = gi.FunctionInvokerNew("Gio", "get_options")
	})
}

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

var unixMountPointGuessNameFunction *gi.Function
var unixMountPointGuessNameFunctionOnce sync.Once

func unixMountPointGuessNameFunctionSet() {
	unixMountPointGuessNameFunctionOnce.Do(func() {
		unixMountPointGuessNameFunction = gi.FunctionInvokerNew("Gio", "guess_name")
	})
}

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

var unixOutputStreamClassStruct *gi.Struct
var unixOutputStreamClassStructOnce sync.Once

func unixOutputStreamClassStructSet() {
	unixOutputStreamClassStructOnce.Do(func() {
		unixOutputStreamClassStruct = gi.StructNew("Gio", "UnixOutputStreamClass")
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

var unixOutputStreamPrivateStruct *gi.Struct
var unixOutputStreamPrivateStructOnce sync.Once

func unixOutputStreamPrivateStructSet() {
	unixOutputStreamPrivateStructOnce.Do(func() {
		unixOutputStreamPrivateStruct = gi.StructNew("Gio", "UnixOutputStreamPrivate")
	})
}

type UnixOutputStreamPrivate struct {
	native uintptr
}

var unixSocketAddressClassStruct *gi.Struct
var unixSocketAddressClassStructOnce sync.Once

func unixSocketAddressClassStructSet() {
	unixSocketAddressClassStructOnce.Do(func() {
		unixSocketAddressClassStruct = gi.StructNew("Gio", "UnixSocketAddressClass")
	})
}

type UnixSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var unixSocketAddressPrivateStruct *gi.Struct
var unixSocketAddressPrivateStructOnce sync.Once

func unixSocketAddressPrivateStructSet() {
	unixSocketAddressPrivateStructOnce.Do(func() {
		unixSocketAddressPrivateStruct = gi.StructNew("Gio", "UnixSocketAddressPrivate")
	})
}

type UnixSocketAddressPrivate struct {
	native uintptr
}

var vfsClassStruct *gi.Struct
var vfsClassStructOnce sync.Once

func vfsClassStructSet() {
	vfsClassStructOnce.Do(func() {
		vfsClassStruct = gi.StructNew("Gio", "VfsClass")
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

var volumeIfaceStruct *gi.Struct
var volumeIfaceStructOnce sync.Once

func volumeIfaceStructSet() {
	volumeIfaceStructOnce.Do(func() {
		volumeIfaceStruct = gi.StructNew("Gio", "VolumeIface")
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

var volumeMonitorClassStruct *gi.Struct
var volumeMonitorClassStructOnce sync.Once

func volumeMonitorClassStructSet() {
	volumeMonitorClassStructOnce.Do(func() {
		volumeMonitorClassStruct = gi.StructNew("Gio", "VolumeMonitorClass")
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

var zlibCompressorClassStruct *gi.Struct
var zlibCompressorClassStructOnce sync.Once

func zlibCompressorClassStructSet() {
	zlibCompressorClassStructOnce.Do(func() {
		zlibCompressorClassStruct = gi.StructNew("Gio", "ZlibCompressorClass")
	})
}

type ZlibCompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var zlibDecompressorClassStruct *gi.Struct
var zlibDecompressorClassStructOnce sync.Once

func zlibDecompressorClassStructSet() {
	zlibDecompressorClassStructOnce.Do(func() {
		zlibDecompressorClassStruct = gi.StructNew("Gio", "ZlibDecompressorClass")
	})
}

type ZlibDecompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}
