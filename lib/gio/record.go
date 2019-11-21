// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionEntryStruct *gi.Struct
var actionEntryStruct_Once sync.Once

func actionEntryStruct_Set() {
	actionEntryStruct_Once.Do(func() {
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
var actionGroupInterfaceStruct_Once sync.Once

func actionGroupInterfaceStruct_Set() {
	actionGroupInterfaceStruct_Once.Do(func() {
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
var actionInterfaceStruct_Once sync.Once

func actionInterfaceStruct_Set() {
	actionInterfaceStruct_Once.Do(func() {
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
var actionMapInterfaceStruct_Once sync.Once

func actionMapInterfaceStruct_Set() {
	actionMapInterfaceStruct_Once.Do(func() {
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
var appInfoIfaceStruct_Once sync.Once

func appInfoIfaceStruct_Set() {
	appInfoIfaceStruct_Once.Do(func() {
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
var appLaunchContextClassStruct_Once sync.Once

func appLaunchContextClassStruct_Set() {
	appLaunchContextClassStruct_Once.Do(func() {
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
var appLaunchContextPrivateStruct_Once sync.Once

func appLaunchContextPrivateStruct_Set() {
	appLaunchContextPrivateStruct_Once.Do(func() {
		appLaunchContextPrivateStruct = gi.StructNew("Gio", "AppLaunchContextPrivate")
	})
}

type AppLaunchContextPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStruct_Once sync.Once

func applicationClassStruct_Set() {
	applicationClassStruct_Once.Do(func() {
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
var applicationCommandLineClassStruct_Once sync.Once

func applicationCommandLineClassStruct_Set() {
	applicationCommandLineClassStruct_Once.Do(func() {
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
var applicationCommandLinePrivateStruct_Once sync.Once

func applicationCommandLinePrivateStruct_Set() {
	applicationCommandLinePrivateStruct_Once.Do(func() {
		applicationCommandLinePrivateStruct = gi.StructNew("Gio", "ApplicationCommandLinePrivate")
	})
}

type ApplicationCommandLinePrivate struct {
	native uintptr
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStruct_Once sync.Once

func applicationPrivateStruct_Set() {
	applicationPrivateStruct_Once.Do(func() {
		applicationPrivateStruct = gi.StructNew("Gio", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var asyncInitableIfaceStruct *gi.Struct
var asyncInitableIfaceStruct_Once sync.Once

func asyncInitableIfaceStruct_Set() {
	asyncInitableIfaceStruct_Once.Do(func() {
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
var asyncResultIfaceStruct_Once sync.Once

func asyncResultIfaceStruct_Set() {
	asyncResultIfaceStruct_Once.Do(func() {
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
var bufferedInputStreamClassStruct_Once sync.Once

func bufferedInputStreamClassStruct_Set() {
	bufferedInputStreamClassStruct_Once.Do(func() {
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
var bufferedInputStreamPrivateStruct_Once sync.Once

func bufferedInputStreamPrivateStruct_Set() {
	bufferedInputStreamPrivateStruct_Once.Do(func() {
		bufferedInputStreamPrivateStruct = gi.StructNew("Gio", "BufferedInputStreamPrivate")
	})
}

type BufferedInputStreamPrivate struct {
	native uintptr
}

var bufferedOutputStreamClassStruct *gi.Struct
var bufferedOutputStreamClassStruct_Once sync.Once

func bufferedOutputStreamClassStruct_Set() {
	bufferedOutputStreamClassStruct_Once.Do(func() {
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
var bufferedOutputStreamPrivateStruct_Once sync.Once

func bufferedOutputStreamPrivateStruct_Set() {
	bufferedOutputStreamPrivateStruct_Once.Do(func() {
		bufferedOutputStreamPrivateStruct = gi.StructNew("Gio", "BufferedOutputStreamPrivate")
	})
}

type BufferedOutputStreamPrivate struct {
	native uintptr
}

var cancellableClassStruct *gi.Struct
var cancellableClassStruct_Once sync.Once

func cancellableClassStruct_Set() {
	cancellableClassStruct_Once.Do(func() {
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
var cancellablePrivateStruct_Once sync.Once

func cancellablePrivateStruct_Set() {
	cancellablePrivateStruct_Once.Do(func() {
		cancellablePrivateStruct = gi.StructNew("Gio", "CancellablePrivate")
	})
}

type CancellablePrivate struct {
	native uintptr
}

var charsetConverterClassStruct *gi.Struct
var charsetConverterClassStruct_Once sync.Once

func charsetConverterClassStruct_Set() {
	charsetConverterClassStruct_Once.Do(func() {
		charsetConverterClassStruct = gi.StructNew("Gio", "CharsetConverterClass")
	})
}

type CharsetConverterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var converterIfaceStruct *gi.Struct
var converterIfaceStruct_Once sync.Once

func converterIfaceStruct_Set() {
	converterIfaceStruct_Once.Do(func() {
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
var converterInputStreamClassStruct_Once sync.Once

func converterInputStreamClassStruct_Set() {
	converterInputStreamClassStruct_Once.Do(func() {
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
var converterInputStreamPrivateStruct_Once sync.Once

func converterInputStreamPrivateStruct_Set() {
	converterInputStreamPrivateStruct_Once.Do(func() {
		converterInputStreamPrivateStruct = gi.StructNew("Gio", "ConverterInputStreamPrivate")
	})
}

type ConverterInputStreamPrivate struct {
	native uintptr
}

var converterOutputStreamClassStruct *gi.Struct
var converterOutputStreamClassStruct_Once sync.Once

func converterOutputStreamClassStruct_Set() {
	converterOutputStreamClassStruct_Once.Do(func() {
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
var converterOutputStreamPrivateStruct_Once sync.Once

func converterOutputStreamPrivateStruct_Set() {
	converterOutputStreamPrivateStruct_Once.Do(func() {
		converterOutputStreamPrivateStruct = gi.StructNew("Gio", "ConverterOutputStreamPrivate")
	})
}

type ConverterOutputStreamPrivate struct {
	native uintptr
}

var credentialsClassStruct *gi.Struct
var credentialsClassStruct_Once sync.Once

func credentialsClassStruct_Set() {
	credentialsClassStruct_Once.Do(func() {
		credentialsClassStruct = gi.StructNew("Gio", "CredentialsClass")
	})
}

type CredentialsClass struct {
	native uintptr
}

var dBusAnnotationInfoStruct *gi.Struct
var dBusAnnotationInfoStruct_Once sync.Once

func dBusAnnotationInfoStruct_Set() {
	dBusAnnotationInfoStruct_Once.Do(func() {
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
var dBusAnnotationInfoRefFunction_Once sync.Once

func dBusAnnotationInfoRefFunction_Set() {
	dBusAnnotationInfoRefFunction_Once.Do(func() {
		dBusAnnotationInfoStruct_Set()
		dBusAnnotationInfoRefFunction = dBusAnnotationInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_annotation_info_ref.
func (recv *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	dBusAnnotationInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusAnnotationInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusAnnotationInfo{native: ret.Pointer()}

	return retGo
}

var dBusAnnotationInfoUnrefFunction *gi.Function
var dBusAnnotationInfoUnrefFunction_Once sync.Once

func dBusAnnotationInfoUnrefFunction_Set() {
	dBusAnnotationInfoUnrefFunction_Once.Do(func() {
		dBusAnnotationInfoStruct_Set()
		dBusAnnotationInfoUnrefFunction = dBusAnnotationInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_annotation_info_unref.
func (recv *DBusAnnotationInfo) Unref() {
	dBusAnnotationInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusAnnotationInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusArgInfoStruct *gi.Struct
var dBusArgInfoStruct_Once sync.Once

func dBusArgInfoStruct_Set() {
	dBusArgInfoStruct_Once.Do(func() {
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
var dBusArgInfoRefFunction_Once sync.Once

func dBusArgInfoRefFunction_Set() {
	dBusArgInfoRefFunction_Once.Do(func() {
		dBusArgInfoStruct_Set()
		dBusArgInfoRefFunction = dBusArgInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_arg_info_ref.
func (recv *DBusArgInfo) Ref() *DBusArgInfo {
	dBusArgInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusArgInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusArgInfo{native: ret.Pointer()}

	return retGo
}

var dBusArgInfoUnrefFunction *gi.Function
var dBusArgInfoUnrefFunction_Once sync.Once

func dBusArgInfoUnrefFunction_Set() {
	dBusArgInfoUnrefFunction_Once.Do(func() {
		dBusArgInfoStruct_Set()
		dBusArgInfoUnrefFunction = dBusArgInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_arg_info_unref.
func (recv *DBusArgInfo) Unref() {
	dBusArgInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusArgInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusErrorEntryStruct *gi.Struct
var dBusErrorEntryStruct_Once sync.Once

func dBusErrorEntryStruct_Set() {
	dBusErrorEntryStruct_Once.Do(func() {
		dBusErrorEntryStruct = gi.StructNew("Gio", "DBusErrorEntry")
	})
}

type DBusErrorEntry struct {
	native        uintptr
	ErrorCode     int32
	DbusErrorName string
}

var dBusInterfaceIfaceStruct *gi.Struct
var dBusInterfaceIfaceStruct_Once sync.Once

func dBusInterfaceIfaceStruct_Set() {
	dBusInterfaceIfaceStruct_Once.Do(func() {
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
var dBusInterfaceInfoStruct_Once sync.Once

func dBusInterfaceInfoStruct_Set() {
	dBusInterfaceInfoStruct_Once.Do(func() {
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
var dBusInterfaceInfoCacheBuildFunction_Once sync.Once

func dBusInterfaceInfoCacheBuildFunction_Set() {
	dBusInterfaceInfoCacheBuildFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoCacheBuildFunction = dBusInterfaceInfoStruct.InvokerNew("cache_build")
	})
}

// CacheBuild is a representation of the C type g_dbus_interface_info_cache_build.
func (recv *DBusInterfaceInfo) CacheBuild() {
	dBusInterfaceInfoCacheBuildFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusInterfaceInfoCacheBuildFunction.Invoke(inArgs[:], nil)

}

var dBusInterfaceInfoCacheReleaseFunction *gi.Function
var dBusInterfaceInfoCacheReleaseFunction_Once sync.Once

func dBusInterfaceInfoCacheReleaseFunction_Set() {
	dBusInterfaceInfoCacheReleaseFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoCacheReleaseFunction = dBusInterfaceInfoStruct.InvokerNew("cache_release")
	})
}

// CacheRelease is a representation of the C type g_dbus_interface_info_cache_release.
func (recv *DBusInterfaceInfo) CacheRelease() {
	dBusInterfaceInfoCacheReleaseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusInterfaceInfoCacheReleaseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_dbus_interface_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusInterfaceInfoLookupMethodFunction *gi.Function
var dBusInterfaceInfoLookupMethodFunction_Once sync.Once

func dBusInterfaceInfoLookupMethodFunction_Set() {
	dBusInterfaceInfoLookupMethodFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoLookupMethodFunction = dBusInterfaceInfoStruct.InvokerNew("lookup_method")
	})
}

// LookupMethod is a representation of the C type g_dbus_interface_info_lookup_method.
func (recv *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	dBusInterfaceInfoLookupMethodFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := dBusInterfaceInfoLookupMethodFunction.Invoke(inArgs[:], nil)

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo
}

var dBusInterfaceInfoLookupPropertyFunction *gi.Function
var dBusInterfaceInfoLookupPropertyFunction_Once sync.Once

func dBusInterfaceInfoLookupPropertyFunction_Set() {
	dBusInterfaceInfoLookupPropertyFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoLookupPropertyFunction = dBusInterfaceInfoStruct.InvokerNew("lookup_property")
	})
}

// LookupProperty is a representation of the C type g_dbus_interface_info_lookup_property.
func (recv *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	dBusInterfaceInfoLookupPropertyFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := dBusInterfaceInfoLookupPropertyFunction.Invoke(inArgs[:], nil)

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo
}

var dBusInterfaceInfoLookupSignalFunction *gi.Function
var dBusInterfaceInfoLookupSignalFunction_Once sync.Once

func dBusInterfaceInfoLookupSignalFunction_Set() {
	dBusInterfaceInfoLookupSignalFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoLookupSignalFunction = dBusInterfaceInfoStruct.InvokerNew("lookup_signal")
	})
}

// LookupSignal is a representation of the C type g_dbus_interface_info_lookup_signal.
func (recv *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	dBusInterfaceInfoLookupSignalFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := dBusInterfaceInfoLookupSignalFunction.Invoke(inArgs[:], nil)

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo
}

var dBusInterfaceInfoRefFunction *gi.Function
var dBusInterfaceInfoRefFunction_Once sync.Once

func dBusInterfaceInfoRefFunction_Set() {
	dBusInterfaceInfoRefFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoRefFunction = dBusInterfaceInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_interface_info_ref.
func (recv *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	dBusInterfaceInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusInterfaceInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo
}

var dBusInterfaceInfoUnrefFunction *gi.Function
var dBusInterfaceInfoUnrefFunction_Once sync.Once

func dBusInterfaceInfoUnrefFunction_Set() {
	dBusInterfaceInfoUnrefFunction_Once.Do(func() {
		dBusInterfaceInfoStruct_Set()
		dBusInterfaceInfoUnrefFunction = dBusInterfaceInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_interface_info_unref.
func (recv *DBusInterfaceInfo) Unref() {
	dBusInterfaceInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusInterfaceInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusInterfaceSkeletonClassStruct *gi.Struct
var dBusInterfaceSkeletonClassStruct_Once sync.Once

func dBusInterfaceSkeletonClassStruct_Set() {
	dBusInterfaceSkeletonClassStruct_Once.Do(func() {
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
var dBusInterfaceSkeletonPrivateStruct_Once sync.Once

func dBusInterfaceSkeletonPrivateStruct_Set() {
	dBusInterfaceSkeletonPrivateStruct_Once.Do(func() {
		dBusInterfaceSkeletonPrivateStruct = gi.StructNew("Gio", "DBusInterfaceSkeletonPrivate")
	})
}

type DBusInterfaceSkeletonPrivate struct {
	native uintptr
}

var dBusInterfaceVTableStruct *gi.Struct
var dBusInterfaceVTableStruct_Once sync.Once

func dBusInterfaceVTableStruct_Set() {
	dBusInterfaceVTableStruct_Once.Do(func() {
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
var dBusMethodInfoStruct_Once sync.Once

func dBusMethodInfoStruct_Set() {
	dBusMethodInfoStruct_Once.Do(func() {
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
var dBusMethodInfoRefFunction_Once sync.Once

func dBusMethodInfoRefFunction_Set() {
	dBusMethodInfoRefFunction_Once.Do(func() {
		dBusMethodInfoStruct_Set()
		dBusMethodInfoRefFunction = dBusMethodInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_method_info_ref.
func (recv *DBusMethodInfo) Ref() *DBusMethodInfo {
	dBusMethodInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusMethodInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo
}

var dBusMethodInfoUnrefFunction *gi.Function
var dBusMethodInfoUnrefFunction_Once sync.Once

func dBusMethodInfoUnrefFunction_Set() {
	dBusMethodInfoUnrefFunction_Once.Do(func() {
		dBusMethodInfoStruct_Set()
		dBusMethodInfoUnrefFunction = dBusMethodInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_method_info_unref.
func (recv *DBusMethodInfo) Unref() {
	dBusMethodInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusMethodInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusNodeInfoStruct *gi.Struct
var dBusNodeInfoStruct_Once sync.Once

func dBusNodeInfoStruct_Set() {
	dBusNodeInfoStruct_Once.Do(func() {
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
var dBusNodeInfoNewForXmlFunction_Once sync.Once

func dBusNodeInfoNewForXmlFunction_Set() {
	dBusNodeInfoNewForXmlFunction_Once.Do(func() {
		dBusNodeInfoStruct_Set()
		dBusNodeInfoNewForXmlFunction = dBusNodeInfoStruct.InvokerNew("new_for_xml")
	})
}

// DBusNodeInfoNewForXml is a representation of the C type g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) *DBusNodeInfo {
	dBusNodeInfoNewForXmlFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(xmlData)

	ret := dBusNodeInfoNewForXmlFunction.Invoke(inArgs[:], nil)

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_node_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusNodeInfoLookupInterfaceFunction *gi.Function
var dBusNodeInfoLookupInterfaceFunction_Once sync.Once

func dBusNodeInfoLookupInterfaceFunction_Set() {
	dBusNodeInfoLookupInterfaceFunction_Once.Do(func() {
		dBusNodeInfoStruct_Set()
		dBusNodeInfoLookupInterfaceFunction = dBusNodeInfoStruct.InvokerNew("lookup_interface")
	})
}

// LookupInterface is a representation of the C type g_dbus_node_info_lookup_interface.
func (recv *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	dBusNodeInfoLookupInterfaceFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := dBusNodeInfoLookupInterfaceFunction.Invoke(inArgs[:], nil)

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo
}

var dBusNodeInfoRefFunction *gi.Function
var dBusNodeInfoRefFunction_Once sync.Once

func dBusNodeInfoRefFunction_Set() {
	dBusNodeInfoRefFunction_Once.Do(func() {
		dBusNodeInfoStruct_Set()
		dBusNodeInfoRefFunction = dBusNodeInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_node_info_ref.
func (recv *DBusNodeInfo) Ref() *DBusNodeInfo {
	dBusNodeInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusNodeInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo
}

var dBusNodeInfoUnrefFunction *gi.Function
var dBusNodeInfoUnrefFunction_Once sync.Once

func dBusNodeInfoUnrefFunction_Set() {
	dBusNodeInfoUnrefFunction_Once.Do(func() {
		dBusNodeInfoStruct_Set()
		dBusNodeInfoUnrefFunction = dBusNodeInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_node_info_unref.
func (recv *DBusNodeInfo) Unref() {
	dBusNodeInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusNodeInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusObjectIfaceStruct *gi.Struct
var dBusObjectIfaceStruct_Once sync.Once

func dBusObjectIfaceStruct_Set() {
	dBusObjectIfaceStruct_Once.Do(func() {
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
var dBusObjectManagerClientClassStruct_Once sync.Once

func dBusObjectManagerClientClassStruct_Set() {
	dBusObjectManagerClientClassStruct_Once.Do(func() {
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
var dBusObjectManagerClientPrivateStruct_Once sync.Once

func dBusObjectManagerClientPrivateStruct_Set() {
	dBusObjectManagerClientPrivateStruct_Once.Do(func() {
		dBusObjectManagerClientPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerClientPrivate")
	})
}

type DBusObjectManagerClientPrivate struct {
	native uintptr
}

var dBusObjectManagerIfaceStruct *gi.Struct
var dBusObjectManagerIfaceStruct_Once sync.Once

func dBusObjectManagerIfaceStruct_Set() {
	dBusObjectManagerIfaceStruct_Once.Do(func() {
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
var dBusObjectManagerServerClassStruct_Once sync.Once

func dBusObjectManagerServerClassStruct_Set() {
	dBusObjectManagerServerClassStruct_Once.Do(func() {
		dBusObjectManagerServerClassStruct = gi.StructNew("Gio", "DBusObjectManagerServerClass")
	})
}

type DBusObjectManagerServerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var dBusObjectManagerServerPrivateStruct *gi.Struct
var dBusObjectManagerServerPrivateStruct_Once sync.Once

func dBusObjectManagerServerPrivateStruct_Set() {
	dBusObjectManagerServerPrivateStruct_Once.Do(func() {
		dBusObjectManagerServerPrivateStruct = gi.StructNew("Gio", "DBusObjectManagerServerPrivate")
	})
}

type DBusObjectManagerServerPrivate struct {
	native uintptr
}

var dBusObjectProxyClassStruct *gi.Struct
var dBusObjectProxyClassStruct_Once sync.Once

func dBusObjectProxyClassStruct_Set() {
	dBusObjectProxyClassStruct_Once.Do(func() {
		dBusObjectProxyClassStruct = gi.StructNew("Gio", "DBusObjectProxyClass")
	})
}

type DBusObjectProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var dBusObjectProxyPrivateStruct *gi.Struct
var dBusObjectProxyPrivateStruct_Once sync.Once

func dBusObjectProxyPrivateStruct_Set() {
	dBusObjectProxyPrivateStruct_Once.Do(func() {
		dBusObjectProxyPrivateStruct = gi.StructNew("Gio", "DBusObjectProxyPrivate")
	})
}

type DBusObjectProxyPrivate struct {
	native uintptr
}

var dBusObjectSkeletonClassStruct *gi.Struct
var dBusObjectSkeletonClassStruct_Once sync.Once

func dBusObjectSkeletonClassStruct_Set() {
	dBusObjectSkeletonClassStruct_Once.Do(func() {
		dBusObjectSkeletonClassStruct = gi.StructNew("Gio", "DBusObjectSkeletonClass")
	})
}

type DBusObjectSkeletonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'authorize_method' : missing Type
}

var dBusObjectSkeletonPrivateStruct *gi.Struct
var dBusObjectSkeletonPrivateStruct_Once sync.Once

func dBusObjectSkeletonPrivateStruct_Set() {
	dBusObjectSkeletonPrivateStruct_Once.Do(func() {
		dBusObjectSkeletonPrivateStruct = gi.StructNew("Gio", "DBusObjectSkeletonPrivate")
	})
}

type DBusObjectSkeletonPrivate struct {
	native uintptr
}

var dBusPropertyInfoStruct *gi.Struct
var dBusPropertyInfoStruct_Once sync.Once

func dBusPropertyInfoStruct_Set() {
	dBusPropertyInfoStruct_Once.Do(func() {
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
var dBusPropertyInfoRefFunction_Once sync.Once

func dBusPropertyInfoRefFunction_Set() {
	dBusPropertyInfoRefFunction_Once.Do(func() {
		dBusPropertyInfoStruct_Set()
		dBusPropertyInfoRefFunction = dBusPropertyInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_property_info_ref.
func (recv *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	dBusPropertyInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusPropertyInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo
}

var dBusPropertyInfoUnrefFunction *gi.Function
var dBusPropertyInfoUnrefFunction_Once sync.Once

func dBusPropertyInfoUnrefFunction_Set() {
	dBusPropertyInfoUnrefFunction_Once.Do(func() {
		dBusPropertyInfoStruct_Set()
		dBusPropertyInfoUnrefFunction = dBusPropertyInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_property_info_unref.
func (recv *DBusPropertyInfo) Unref() {
	dBusPropertyInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusPropertyInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusProxyClassStruct *gi.Struct
var dBusProxyClassStruct_Once sync.Once

func dBusProxyClassStruct_Set() {
	dBusProxyClassStruct_Once.Do(func() {
		dBusProxyClassStruct = gi.StructNew("Gio", "DBusProxyClass")
	})
}

type DBusProxyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'g_properties_changed' : missing Type
	// UNSUPPORTED : C value 'g_signal' : missing Type
}

var dBusProxyPrivateStruct *gi.Struct
var dBusProxyPrivateStruct_Once sync.Once

func dBusProxyPrivateStruct_Set() {
	dBusProxyPrivateStruct_Once.Do(func() {
		dBusProxyPrivateStruct = gi.StructNew("Gio", "DBusProxyPrivate")
	})
}

type DBusProxyPrivate struct {
	native uintptr
}

var dBusSignalInfoStruct *gi.Struct
var dBusSignalInfoStruct_Once sync.Once

func dBusSignalInfoStruct_Set() {
	dBusSignalInfoStruct_Once.Do(func() {
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
var dBusSignalInfoRefFunction_Once sync.Once

func dBusSignalInfoRefFunction_Set() {
	dBusSignalInfoRefFunction_Once.Do(func() {
		dBusSignalInfoStruct_Set()
		dBusSignalInfoRefFunction = dBusSignalInfoStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_dbus_signal_info_ref.
func (recv *DBusSignalInfo) Ref() *DBusSignalInfo {
	dBusSignalInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dBusSignalInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo
}

var dBusSignalInfoUnrefFunction *gi.Function
var dBusSignalInfoUnrefFunction_Once sync.Once

func dBusSignalInfoUnrefFunction_Set() {
	dBusSignalInfoUnrefFunction_Once.Do(func() {
		dBusSignalInfoStruct_Set()
		dBusSignalInfoUnrefFunction = dBusSignalInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_dbus_signal_info_unref.
func (recv *DBusSignalInfo) Unref() {
	dBusSignalInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dBusSignalInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var dBusSubtreeVTableStruct *gi.Struct
var dBusSubtreeVTableStruct_Once sync.Once

func dBusSubtreeVTableStruct_Set() {
	dBusSubtreeVTableStruct_Once.Do(func() {
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
var dataInputStreamClassStruct_Once sync.Once

func dataInputStreamClassStruct_Set() {
	dataInputStreamClassStruct_Once.Do(func() {
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
var dataInputStreamPrivateStruct_Once sync.Once

func dataInputStreamPrivateStruct_Set() {
	dataInputStreamPrivateStruct_Once.Do(func() {
		dataInputStreamPrivateStruct = gi.StructNew("Gio", "DataInputStreamPrivate")
	})
}

type DataInputStreamPrivate struct {
	native uintptr
}

var dataOutputStreamClassStruct *gi.Struct
var dataOutputStreamClassStruct_Once sync.Once

func dataOutputStreamClassStruct_Set() {
	dataOutputStreamClassStruct_Once.Do(func() {
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
var dataOutputStreamPrivateStruct_Once sync.Once

func dataOutputStreamPrivateStruct_Set() {
	dataOutputStreamPrivateStruct_Once.Do(func() {
		dataOutputStreamPrivateStruct = gi.StructNew("Gio", "DataOutputStreamPrivate")
	})
}

type DataOutputStreamPrivate struct {
	native uintptr
}

var datagramBasedInterfaceStruct *gi.Struct
var datagramBasedInterfaceStruct_Once sync.Once

func datagramBasedInterfaceStruct_Set() {
	datagramBasedInterfaceStruct_Once.Do(func() {
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
var desktopAppInfoClassStruct_Once sync.Once

func desktopAppInfoClassStruct_Set() {
	desktopAppInfoClassStruct_Once.Do(func() {
		desktopAppInfoClassStruct = gi.StructNew("Gio", "DesktopAppInfoClass")
	})
}

type DesktopAppInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var desktopAppInfoLookupIfaceStruct *gi.Struct
var desktopAppInfoLookupIfaceStruct_Once sync.Once

func desktopAppInfoLookupIfaceStruct_Set() {
	desktopAppInfoLookupIfaceStruct_Once.Do(func() {
		desktopAppInfoLookupIfaceStruct = gi.StructNew("Gio", "DesktopAppInfoLookupIface")
	})
}

type DesktopAppInfoLookupIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_default_for_uri_scheme' : missing Type
}

var driveIfaceStruct *gi.Struct
var driveIfaceStruct_Once sync.Once

func driveIfaceStruct_Set() {
	driveIfaceStruct_Once.Do(func() {
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
var dtlsClientConnectionInterfaceStruct_Once sync.Once

func dtlsClientConnectionInterfaceStruct_Set() {
	dtlsClientConnectionInterfaceStruct_Once.Do(func() {
		dtlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsClientConnectionInterface")
	})
}

type DtlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var dtlsConnectionInterfaceStruct *gi.Struct
var dtlsConnectionInterfaceStruct_Once sync.Once

func dtlsConnectionInterfaceStruct_Set() {
	dtlsConnectionInterfaceStruct_Once.Do(func() {
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
var dtlsServerConnectionInterfaceStruct_Once sync.Once

func dtlsServerConnectionInterfaceStruct_Set() {
	dtlsServerConnectionInterfaceStruct_Once.Do(func() {
		dtlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "DtlsServerConnectionInterface")
	})
}

type DtlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var emblemClassStruct *gi.Struct
var emblemClassStruct_Once sync.Once

func emblemClassStruct_Set() {
	emblemClassStruct_Once.Do(func() {
		emblemClassStruct = gi.StructNew("Gio", "EmblemClass")
	})
}

type EmblemClass struct {
	native uintptr
}

var emblemedIconClassStruct *gi.Struct
var emblemedIconClassStruct_Once sync.Once

func emblemedIconClassStruct_Set() {
	emblemedIconClassStruct_Once.Do(func() {
		emblemedIconClassStruct = gi.StructNew("Gio", "EmblemedIconClass")
	})
}

type EmblemedIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var emblemedIconPrivateStruct *gi.Struct
var emblemedIconPrivateStruct_Once sync.Once

func emblemedIconPrivateStruct_Set() {
	emblemedIconPrivateStruct_Once.Do(func() {
		emblemedIconPrivateStruct = gi.StructNew("Gio", "EmblemedIconPrivate")
	})
}

type EmblemedIconPrivate struct {
	native uintptr
}

var fileAttributeInfoStruct *gi.Struct
var fileAttributeInfoStruct_Once sync.Once

func fileAttributeInfoStruct_Set() {
	fileAttributeInfoStruct_Once.Do(func() {
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
var fileAttributeInfoListStruct_Once sync.Once

func fileAttributeInfoListStruct_Set() {
	fileAttributeInfoListStruct_Once.Do(func() {
		fileAttributeInfoListStruct = gi.StructNew("Gio", "FileAttributeInfoList")
	})
}

type FileAttributeInfoList struct {
	native uintptr
	Infos  *FileAttributeInfo
	NInfos int32
}

var fileAttributeInfoListNewFunction *gi.Function
var fileAttributeInfoListNewFunction_Once sync.Once

func fileAttributeInfoListNewFunction_Set() {
	fileAttributeInfoListNewFunction_Once.Do(func() {
		fileAttributeInfoListStruct_Set()
		fileAttributeInfoListNewFunction = fileAttributeInfoListStruct.InvokerNew("new")
	})
}

// FileAttributeInfoListNew is a representation of the C type g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {
	fileAttributeInfoListNewFunction_Set()

	ret := fileAttributeInfoListNewFunction.Invoke(nil, nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_info_list_add' : parameter 'type' of type 'FileAttributeType' not supported

var fileAttributeInfoListDupFunction *gi.Function
var fileAttributeInfoListDupFunction_Once sync.Once

func fileAttributeInfoListDupFunction_Set() {
	fileAttributeInfoListDupFunction_Once.Do(func() {
		fileAttributeInfoListStruct_Set()
		fileAttributeInfoListDupFunction = fileAttributeInfoListStruct.InvokerNew("dup")
	})
}

// Dup is a representation of the C type g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	fileAttributeInfoListDupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fileAttributeInfoListDupFunction.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

var fileAttributeInfoListLookupFunction *gi.Function
var fileAttributeInfoListLookupFunction_Once sync.Once

func fileAttributeInfoListLookupFunction_Set() {
	fileAttributeInfoListLookupFunction_Once.Do(func() {
		fileAttributeInfoListStruct_Set()
		fileAttributeInfoListLookupFunction = fileAttributeInfoListStruct.InvokerNew("lookup")
	})
}

// Lookup is a representation of the C type g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	fileAttributeInfoListLookupFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := fileAttributeInfoListLookupFunction.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfo{native: ret.Pointer()}

	return retGo
}

var fileAttributeInfoListRefFunction *gi.Function
var fileAttributeInfoListRefFunction_Once sync.Once

func fileAttributeInfoListRefFunction_Set() {
	fileAttributeInfoListRefFunction_Once.Do(func() {
		fileAttributeInfoListStruct_Set()
		fileAttributeInfoListRefFunction = fileAttributeInfoListStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	fileAttributeInfoListRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fileAttributeInfoListRefFunction.Invoke(inArgs[:], nil)

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo
}

var fileAttributeInfoListUnrefFunction *gi.Function
var fileAttributeInfoListUnrefFunction_Once sync.Once

func fileAttributeInfoListUnrefFunction_Set() {
	fileAttributeInfoListUnrefFunction_Once.Do(func() {
		fileAttributeInfoListStruct_Set()
		fileAttributeInfoListUnrefFunction = fileAttributeInfoListStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_file_attribute_info_list_unref.
func (recv *FileAttributeInfoList) Unref() {
	fileAttributeInfoListUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	fileAttributeInfoListUnrefFunction.Invoke(inArgs[:], nil)

}

var fileAttributeMatcherStruct *gi.Struct
var fileAttributeMatcherStruct_Once sync.Once

func fileAttributeMatcherStruct_Set() {
	fileAttributeMatcherStruct_Once.Do(func() {
		fileAttributeMatcherStruct = gi.StructNew("Gio", "FileAttributeMatcher")
	})
}

type FileAttributeMatcher struct {
	native uintptr
}

var fileAttributeMatcherNewFunction *gi.Function
var fileAttributeMatcherNewFunction_Once sync.Once

func fileAttributeMatcherNewFunction_Set() {
	fileAttributeMatcherNewFunction_Once.Do(func() {
		fileAttributeMatcherStruct_Set()
		fileAttributeMatcherNewFunction = fileAttributeMatcherStruct.InvokerNew("new")
	})
}

// FileAttributeMatcherNew is a representation of the C type g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	fileAttributeMatcherNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(attributes)

	ret := fileAttributeMatcherNewFunction.Invoke(inArgs[:], nil)

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_enumerate_namespace' : return type 'gboolean' not supported

var fileAttributeMatcherEnumerateNextFunction *gi.Function
var fileAttributeMatcherEnumerateNextFunction_Once sync.Once

func fileAttributeMatcherEnumerateNextFunction_Set() {
	fileAttributeMatcherEnumerateNextFunction_Once.Do(func() {
		fileAttributeMatcherStruct_Set()
		fileAttributeMatcherEnumerateNextFunction = fileAttributeMatcherStruct.InvokerNew("enumerate_next")
	})
}

// EnumerateNext is a representation of the C type g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() string {
	fileAttributeMatcherEnumerateNextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fileAttributeMatcherEnumerateNextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_file_attribute_matcher_matches_only' : return type 'gboolean' not supported

var fileAttributeMatcherRefFunction *gi.Function
var fileAttributeMatcherRefFunction_Once sync.Once

func fileAttributeMatcherRefFunction_Set() {
	fileAttributeMatcherRefFunction_Once.Do(func() {
		fileAttributeMatcherStruct_Set()
		fileAttributeMatcherRefFunction = fileAttributeMatcherStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	fileAttributeMatcherRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fileAttributeMatcherRefFunction.Invoke(inArgs[:], nil)

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_file_attribute_matcher_subtract' : parameter 'subtract' of type 'FileAttributeMatcher' not supported

var fileAttributeMatcherToStringFunction *gi.Function
var fileAttributeMatcherToStringFunction_Once sync.Once

func fileAttributeMatcherToStringFunction_Set() {
	fileAttributeMatcherToStringFunction_Once.Do(func() {
		fileAttributeMatcherStruct_Set()
		fileAttributeMatcherToStringFunction = fileAttributeMatcherStruct.InvokerNew("to_string")
	})
}

// ToString is a representation of the C type g_file_attribute_matcher_to_string.
func (recv *FileAttributeMatcher) ToString() string {
	fileAttributeMatcherToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fileAttributeMatcherToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fileAttributeMatcherUnrefFunction *gi.Function
var fileAttributeMatcherUnrefFunction_Once sync.Once

func fileAttributeMatcherUnrefFunction_Set() {
	fileAttributeMatcherUnrefFunction_Once.Do(func() {
		fileAttributeMatcherStruct_Set()
		fileAttributeMatcherUnrefFunction = fileAttributeMatcherStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_file_attribute_matcher_unref.
func (recv *FileAttributeMatcher) Unref() {
	fileAttributeMatcherUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	fileAttributeMatcherUnrefFunction.Invoke(inArgs[:], nil)

}

var fileDescriptorBasedIfaceStruct *gi.Struct
var fileDescriptorBasedIfaceStruct_Once sync.Once

func fileDescriptorBasedIfaceStruct_Set() {
	fileDescriptorBasedIfaceStruct_Once.Do(func() {
		fileDescriptorBasedIfaceStruct = gi.StructNew("Gio", "FileDescriptorBasedIface")
	})
}

type FileDescriptorBasedIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_fd' : missing Type
}

var fileEnumeratorClassStruct *gi.Struct
var fileEnumeratorClassStruct_Once sync.Once

func fileEnumeratorClassStruct_Set() {
	fileEnumeratorClassStruct_Once.Do(func() {
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
var fileEnumeratorPrivateStruct_Once sync.Once

func fileEnumeratorPrivateStruct_Set() {
	fileEnumeratorPrivateStruct_Once.Do(func() {
		fileEnumeratorPrivateStruct = gi.StructNew("Gio", "FileEnumeratorPrivate")
	})
}

type FileEnumeratorPrivate struct {
	native uintptr
}

var fileIOStreamClassStruct *gi.Struct
var fileIOStreamClassStruct_Once sync.Once

func fileIOStreamClassStruct_Set() {
	fileIOStreamClassStruct_Once.Do(func() {
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
var fileIOStreamPrivateStruct_Once sync.Once

func fileIOStreamPrivateStruct_Set() {
	fileIOStreamPrivateStruct_Once.Do(func() {
		fileIOStreamPrivateStruct = gi.StructNew("Gio", "FileIOStreamPrivate")
	})
}

type FileIOStreamPrivate struct {
	native uintptr
}

var fileIconClassStruct *gi.Struct
var fileIconClassStruct_Once sync.Once

func fileIconClassStruct_Set() {
	fileIconClassStruct_Once.Do(func() {
		fileIconClassStruct = gi.StructNew("Gio", "FileIconClass")
	})
}

type FileIconClass struct {
	native uintptr
}

var fileIfaceStruct *gi.Struct
var fileIfaceStruct_Once sync.Once

func fileIfaceStruct_Set() {
	fileIfaceStruct_Once.Do(func() {
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
var fileInfoClassStruct_Once sync.Once

func fileInfoClassStruct_Set() {
	fileInfoClassStruct_Once.Do(func() {
		fileInfoClassStruct = gi.StructNew("Gio", "FileInfoClass")
	})
}

type FileInfoClass struct {
	native uintptr
}

var fileInputStreamClassStruct *gi.Struct
var fileInputStreamClassStruct_Once sync.Once

func fileInputStreamClassStruct_Set() {
	fileInputStreamClassStruct_Once.Do(func() {
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
var fileInputStreamPrivateStruct_Once sync.Once

func fileInputStreamPrivateStruct_Set() {
	fileInputStreamPrivateStruct_Once.Do(func() {
		fileInputStreamPrivateStruct = gi.StructNew("Gio", "FileInputStreamPrivate")
	})
}

type FileInputStreamPrivate struct {
	native uintptr
}

var fileMonitorClassStruct *gi.Struct
var fileMonitorClassStruct_Once sync.Once

func fileMonitorClassStruct_Set() {
	fileMonitorClassStruct_Once.Do(func() {
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
var fileMonitorPrivateStruct_Once sync.Once

func fileMonitorPrivateStruct_Set() {
	fileMonitorPrivateStruct_Once.Do(func() {
		fileMonitorPrivateStruct = gi.StructNew("Gio", "FileMonitorPrivate")
	})
}

type FileMonitorPrivate struct {
	native uintptr
}

var fileOutputStreamClassStruct *gi.Struct
var fileOutputStreamClassStruct_Once sync.Once

func fileOutputStreamClassStruct_Set() {
	fileOutputStreamClassStruct_Once.Do(func() {
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
var fileOutputStreamPrivateStruct_Once sync.Once

func fileOutputStreamPrivateStruct_Set() {
	fileOutputStreamPrivateStruct_Once.Do(func() {
		fileOutputStreamPrivateStruct = gi.StructNew("Gio", "FileOutputStreamPrivate")
	})
}

type FileOutputStreamPrivate struct {
	native uintptr
}

var filenameCompleterClassStruct *gi.Struct
var filenameCompleterClassStruct_Once sync.Once

func filenameCompleterClassStruct_Set() {
	filenameCompleterClassStruct_Once.Do(func() {
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
var filterInputStreamClassStruct_Once sync.Once

func filterInputStreamClassStruct_Set() {
	filterInputStreamClassStruct_Once.Do(func() {
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
var filterOutputStreamClassStruct_Once sync.Once

func filterOutputStreamClassStruct_Set() {
	filterOutputStreamClassStruct_Once.Do(func() {
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
var iOExtensionStruct_Once sync.Once

func iOExtensionStruct_Set() {
	iOExtensionStruct_Once.Do(func() {
		iOExtensionStruct = gi.StructNew("Gio", "IOExtension")
	})
}

type IOExtension struct {
	native uintptr
}

var iOExtensionGetNameFunction *gi.Function
var iOExtensionGetNameFunction_Once sync.Once

func iOExtensionGetNameFunction_Set() {
	iOExtensionGetNameFunction_Once.Do(func() {
		iOExtensionStruct_Set()
		iOExtensionGetNameFunction = iOExtensionStruct.InvokerNew("get_name")
	})
}

// GetName is a representation of the C type g_io_extension_get_name.
func (recv *IOExtension) GetName() string {
	iOExtensionGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iOExtensionGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var iOExtensionGetPriorityFunction *gi.Function
var iOExtensionGetPriorityFunction_Once sync.Once

func iOExtensionGetPriorityFunction_Set() {
	iOExtensionGetPriorityFunction_Once.Do(func() {
		iOExtensionStruct_Set()
		iOExtensionGetPriorityFunction = iOExtensionStruct.InvokerNew("get_priority")
	})
}

// GetPriority is a representation of the C type g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() int32 {
	iOExtensionGetPriorityFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iOExtensionGetPriorityFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_get_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_ref_class' : return type 'GObject.TypeClass' not supported

var iOExtensionPointStruct *gi.Struct
var iOExtensionPointStruct_Once sync.Once

func iOExtensionPointStruct_Set() {
	iOExtensionPointStruct_Once.Do(func() {
		iOExtensionPointStruct = gi.StructNew("Gio", "IOExtensionPoint")
	})
}

type IOExtensionPoint struct {
	native uintptr
}

var iOExtensionPointGetExtensionByNameFunction *gi.Function
var iOExtensionPointGetExtensionByNameFunction_Once sync.Once

func iOExtensionPointGetExtensionByNameFunction_Set() {
	iOExtensionPointGetExtensionByNameFunction_Once.Do(func() {
		iOExtensionPointStruct_Set()
		iOExtensionPointGetExtensionByNameFunction = iOExtensionPointStruct.InvokerNew("get_extension_by_name")
	})
}

// GetExtensionByName is a representation of the C type g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	iOExtensionPointGetExtensionByNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := iOExtensionPointGetExtensionByNameFunction.Invoke(inArgs[:], nil)

	retGo := &IOExtension{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_io_extension_point_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_io_extension_point_get_required_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_point_set_required_type' : parameter 'type' of type 'GType' not supported

var iOModuleClassStruct *gi.Struct
var iOModuleClassStruct_Once sync.Once

func iOModuleClassStruct_Set() {
	iOModuleClassStruct_Once.Do(func() {
		iOModuleClassStruct = gi.StructNew("Gio", "IOModuleClass")
	})
}

type IOModuleClass struct {
	native uintptr
}

var iOModuleScopeStruct *gi.Struct
var iOModuleScopeStruct_Once sync.Once

func iOModuleScopeStruct_Set() {
	iOModuleScopeStruct_Once.Do(func() {
		iOModuleScopeStruct = gi.StructNew("Gio", "IOModuleScope")
	})
}

type IOModuleScope struct {
	native uintptr
}

var iOModuleScopeBlockFunction *gi.Function
var iOModuleScopeBlockFunction_Once sync.Once

func iOModuleScopeBlockFunction_Set() {
	iOModuleScopeBlockFunction_Once.Do(func() {
		iOModuleScopeStruct_Set()
		iOModuleScopeBlockFunction = iOModuleScopeStruct.InvokerNew("block")
	})
}

// Block is a representation of the C type g_io_module_scope_block.
func (recv *IOModuleScope) Block(basename string) {
	iOModuleScopeBlockFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(basename)

	iOModuleScopeBlockFunction.Invoke(inArgs[:], nil)

}

var iOModuleScopeFreeFunction *gi.Function
var iOModuleScopeFreeFunction_Once sync.Once

func iOModuleScopeFreeFunction_Set() {
	iOModuleScopeFreeFunction_Once.Do(func() {
		iOModuleScopeStruct_Set()
		iOModuleScopeFreeFunction = iOModuleScopeStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type g_io_module_scope_free.
func (recv *IOModuleScope) Free() {
	iOModuleScopeFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iOModuleScopeFreeFunction.Invoke(inArgs[:], nil)

}

var iOSchedulerJobStruct *gi.Struct
var iOSchedulerJobStruct_Once sync.Once

func iOSchedulerJobStruct_Set() {
	iOSchedulerJobStruct_Once.Do(func() {
		iOSchedulerJobStruct = gi.StructNew("Gio", "IOSchedulerJob")
	})
}

type IOSchedulerJob struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

var iOStreamAdapterStruct *gi.Struct
var iOStreamAdapterStruct_Once sync.Once

func iOStreamAdapterStruct_Set() {
	iOStreamAdapterStruct_Once.Do(func() {
		iOStreamAdapterStruct = gi.StructNew("Gio", "IOStreamAdapter")
	})
}

type IOStreamAdapter struct {
	native uintptr
}

var iOStreamClassStruct *gi.Struct
var iOStreamClassStruct_Once sync.Once

func iOStreamClassStruct_Set() {
	iOStreamClassStruct_Once.Do(func() {
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
var iOStreamPrivateStruct_Once sync.Once

func iOStreamPrivateStruct_Set() {
	iOStreamPrivateStruct_Once.Do(func() {
		iOStreamPrivateStruct = gi.StructNew("Gio", "IOStreamPrivate")
	})
}

type IOStreamPrivate struct {
	native uintptr
}

var iconIfaceStruct *gi.Struct
var iconIfaceStruct_Once sync.Once

func iconIfaceStruct_Set() {
	iconIfaceStruct_Once.Do(func() {
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
var inetAddressClassStruct_Once sync.Once

func inetAddressClassStruct_Set() {
	inetAddressClassStruct_Once.Do(func() {
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
var inetAddressMaskClassStruct_Once sync.Once

func inetAddressMaskClassStruct_Set() {
	inetAddressMaskClassStruct_Once.Do(func() {
		inetAddressMaskClassStruct = gi.StructNew("Gio", "InetAddressMaskClass")
	})
}

type InetAddressMaskClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var inetAddressMaskPrivateStruct *gi.Struct
var inetAddressMaskPrivateStruct_Once sync.Once

func inetAddressMaskPrivateStruct_Set() {
	inetAddressMaskPrivateStruct_Once.Do(func() {
		inetAddressMaskPrivateStruct = gi.StructNew("Gio", "InetAddressMaskPrivate")
	})
}

type InetAddressMaskPrivate struct {
	native uintptr
}

var inetAddressPrivateStruct *gi.Struct
var inetAddressPrivateStruct_Once sync.Once

func inetAddressPrivateStruct_Set() {
	inetAddressPrivateStruct_Once.Do(func() {
		inetAddressPrivateStruct = gi.StructNew("Gio", "InetAddressPrivate")
	})
}

type InetAddressPrivate struct {
	native uintptr
}

var inetSocketAddressClassStruct *gi.Struct
var inetSocketAddressClassStruct_Once sync.Once

func inetSocketAddressClassStruct_Set() {
	inetSocketAddressClassStruct_Once.Do(func() {
		inetSocketAddressClassStruct = gi.StructNew("Gio", "InetSocketAddressClass")
	})
}

type InetSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var inetSocketAddressPrivateStruct *gi.Struct
var inetSocketAddressPrivateStruct_Once sync.Once

func inetSocketAddressPrivateStruct_Set() {
	inetSocketAddressPrivateStruct_Once.Do(func() {
		inetSocketAddressPrivateStruct = gi.StructNew("Gio", "InetSocketAddressPrivate")
	})
}

type InetSocketAddressPrivate struct {
	native uintptr
}

var initableIfaceStruct *gi.Struct
var initableIfaceStruct_Once sync.Once

func initableIfaceStruct_Set() {
	initableIfaceStruct_Once.Do(func() {
		initableIfaceStruct = gi.StructNew("Gio", "InitableIface")
	})
}

type InitableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'init' : missing Type
}

var inputMessageStruct *gi.Struct
var inputMessageStruct_Once sync.Once

func inputMessageStruct_Set() {
	inputMessageStruct_Once.Do(func() {
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
var inputStreamClassStruct_Once sync.Once

func inputStreamClassStruct_Set() {
	inputStreamClassStruct_Once.Do(func() {
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
var inputStreamPrivateStruct_Once sync.Once

func inputStreamPrivateStruct_Set() {
	inputStreamPrivateStruct_Once.Do(func() {
		inputStreamPrivateStruct = gi.StructNew("Gio", "InputStreamPrivate")
	})
}

type InputStreamPrivate struct {
	native uintptr
}

var inputVectorStruct *gi.Struct
var inputVectorStruct_Once sync.Once

func inputVectorStruct_Set() {
	inputVectorStruct_Once.Do(func() {
		inputVectorStruct = gi.StructNew("Gio", "InputVector")
	})
}

type InputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var listModelInterfaceStruct *gi.Struct
var listModelInterfaceStruct_Once sync.Once

func listModelInterfaceStruct_Set() {
	listModelInterfaceStruct_Once.Do(func() {
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
var listStoreClassStruct_Once sync.Once

func listStoreClassStruct_Set() {
	listStoreClassStruct_Once.Do(func() {
		listStoreClassStruct = gi.StructNew("Gio", "ListStoreClass")
	})
}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var loadableIconIfaceStruct *gi.Struct
var loadableIconIfaceStruct_Once sync.Once

func loadableIconIfaceStruct_Set() {
	loadableIconIfaceStruct_Once.Do(func() {
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
var memoryInputStreamClassStruct_Once sync.Once

func memoryInputStreamClassStruct_Set() {
	memoryInputStreamClassStruct_Once.Do(func() {
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
var memoryInputStreamPrivateStruct_Once sync.Once

func memoryInputStreamPrivateStruct_Set() {
	memoryInputStreamPrivateStruct_Once.Do(func() {
		memoryInputStreamPrivateStruct = gi.StructNew("Gio", "MemoryInputStreamPrivate")
	})
}

type MemoryInputStreamPrivate struct {
	native uintptr
}

var memoryOutputStreamClassStruct *gi.Struct
var memoryOutputStreamClassStruct_Once sync.Once

func memoryOutputStreamClassStruct_Set() {
	memoryOutputStreamClassStruct_Once.Do(func() {
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
var memoryOutputStreamPrivateStruct_Once sync.Once

func memoryOutputStreamPrivateStruct_Set() {
	memoryOutputStreamPrivateStruct_Once.Do(func() {
		memoryOutputStreamPrivateStruct = gi.StructNew("Gio", "MemoryOutputStreamPrivate")
	})
}

type MemoryOutputStreamPrivate struct {
	native uintptr
}

var menuAttributeIterClassStruct *gi.Struct
var menuAttributeIterClassStruct_Once sync.Once

func menuAttributeIterClassStruct_Set() {
	menuAttributeIterClassStruct_Once.Do(func() {
		menuAttributeIterClassStruct = gi.StructNew("Gio", "MenuAttributeIterClass")
	})
}

type MenuAttributeIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var menuAttributeIterPrivateStruct *gi.Struct
var menuAttributeIterPrivateStruct_Once sync.Once

func menuAttributeIterPrivateStruct_Set() {
	menuAttributeIterPrivateStruct_Once.Do(func() {
		menuAttributeIterPrivateStruct = gi.StructNew("Gio", "MenuAttributeIterPrivate")
	})
}

type MenuAttributeIterPrivate struct {
	native uintptr
}

var menuLinkIterClassStruct *gi.Struct
var menuLinkIterClassStruct_Once sync.Once

func menuLinkIterClassStruct_Set() {
	menuLinkIterClassStruct_Once.Do(func() {
		menuLinkIterClassStruct = gi.StructNew("Gio", "MenuLinkIterClass")
	})
}

type MenuLinkIterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_next' : missing Type
}

var menuLinkIterPrivateStruct *gi.Struct
var menuLinkIterPrivateStruct_Once sync.Once

func menuLinkIterPrivateStruct_Set() {
	menuLinkIterPrivateStruct_Once.Do(func() {
		menuLinkIterPrivateStruct = gi.StructNew("Gio", "MenuLinkIterPrivate")
	})
}

type MenuLinkIterPrivate struct {
	native uintptr
}

var menuModelClassStruct *gi.Struct
var menuModelClassStruct_Once sync.Once

func menuModelClassStruct_Set() {
	menuModelClassStruct_Once.Do(func() {
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
var menuModelPrivateStruct_Once sync.Once

func menuModelPrivateStruct_Set() {
	menuModelPrivateStruct_Once.Do(func() {
		menuModelPrivateStruct = gi.StructNew("Gio", "MenuModelPrivate")
	})
}

type MenuModelPrivate struct {
	native uintptr
}

var mountIfaceStruct *gi.Struct
var mountIfaceStruct_Once sync.Once

func mountIfaceStruct_Set() {
	mountIfaceStruct_Once.Do(func() {
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
var mountOperationClassStruct_Once sync.Once

func mountOperationClassStruct_Set() {
	mountOperationClassStruct_Once.Do(func() {
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
var mountOperationPrivateStruct_Once sync.Once

func mountOperationPrivateStruct_Set() {
	mountOperationPrivateStruct_Once.Do(func() {
		mountOperationPrivateStruct = gi.StructNew("Gio", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeSocketAddressClassStruct *gi.Struct
var nativeSocketAddressClassStruct_Once sync.Once

func nativeSocketAddressClassStruct_Set() {
	nativeSocketAddressClassStruct_Once.Do(func() {
		nativeSocketAddressClassStruct = gi.StructNew("Gio", "NativeSocketAddressClass")
	})
}

type NativeSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var nativeSocketAddressPrivateStruct *gi.Struct
var nativeSocketAddressPrivateStruct_Once sync.Once

func nativeSocketAddressPrivateStruct_Set() {
	nativeSocketAddressPrivateStruct_Once.Do(func() {
		nativeSocketAddressPrivateStruct = gi.StructNew("Gio", "NativeSocketAddressPrivate")
	})
}

type NativeSocketAddressPrivate struct {
	native uintptr
}

var nativeVolumeMonitorClassStruct *gi.Struct
var nativeVolumeMonitorClassStruct_Once sync.Once

func nativeVolumeMonitorClassStruct_Set() {
	nativeVolumeMonitorClassStruct_Once.Do(func() {
		nativeVolumeMonitorClassStruct = gi.StructNew("Gio", "NativeVolumeMonitorClass")
	})
}

type NativeVolumeMonitorClass struct {
	native      uintptr
	ParentClass *VolumeMonitorClass
	// UNSUPPORTED : C value 'get_mount_for_mount_path' : missing Type
}

var networkAddressClassStruct *gi.Struct
var networkAddressClassStruct_Once sync.Once

func networkAddressClassStruct_Set() {
	networkAddressClassStruct_Once.Do(func() {
		networkAddressClassStruct = gi.StructNew("Gio", "NetworkAddressClass")
	})
}

type NetworkAddressClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var networkAddressPrivateStruct *gi.Struct
var networkAddressPrivateStruct_Once sync.Once

func networkAddressPrivateStruct_Set() {
	networkAddressPrivateStruct_Once.Do(func() {
		networkAddressPrivateStruct = gi.StructNew("Gio", "NetworkAddressPrivate")
	})
}

type NetworkAddressPrivate struct {
	native uintptr
}

var networkMonitorInterfaceStruct *gi.Struct
var networkMonitorInterfaceStruct_Once sync.Once

func networkMonitorInterfaceStruct_Set() {
	networkMonitorInterfaceStruct_Once.Do(func() {
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
var networkServiceClassStruct_Once sync.Once

func networkServiceClassStruct_Set() {
	networkServiceClassStruct_Once.Do(func() {
		networkServiceClassStruct = gi.StructNew("Gio", "NetworkServiceClass")
	})
}

type NetworkServiceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var networkServicePrivateStruct *gi.Struct
var networkServicePrivateStruct_Once sync.Once

func networkServicePrivateStruct_Set() {
	networkServicePrivateStruct_Once.Do(func() {
		networkServicePrivateStruct = gi.StructNew("Gio", "NetworkServicePrivate")
	})
}

type NetworkServicePrivate struct {
	native uintptr
}

var outputMessageStruct *gi.Struct
var outputMessageStruct_Once sync.Once

func outputMessageStruct_Set() {
	outputMessageStruct_Once.Do(func() {
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
var outputStreamClassStruct_Once sync.Once

func outputStreamClassStruct_Set() {
	outputStreamClassStruct_Once.Do(func() {
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
var outputStreamPrivateStruct_Once sync.Once

func outputStreamPrivateStruct_Set() {
	outputStreamPrivateStruct_Once.Do(func() {
		outputStreamPrivateStruct = gi.StructNew("Gio", "OutputStreamPrivate")
	})
}

type OutputStreamPrivate struct {
	native uintptr
}

var outputVectorStruct *gi.Struct
var outputVectorStruct_Once sync.Once

func outputVectorStruct_Set() {
	outputVectorStruct_Once.Do(func() {
		outputVectorStruct = gi.StructNew("Gio", "OutputVector")
	})
}

type OutputVector struct {
	native uintptr
	// UNSUPPORTED : C value 'buffer' : no Go type for 'gpointer'
	Size uintptr
}

var permissionClassStruct *gi.Struct
var permissionClassStruct_Once sync.Once

func permissionClassStruct_Set() {
	permissionClassStruct_Once.Do(func() {
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
var permissionPrivateStruct_Once sync.Once

func permissionPrivateStruct_Set() {
	permissionPrivateStruct_Once.Do(func() {
		permissionPrivateStruct = gi.StructNew("Gio", "PermissionPrivate")
	})
}

type PermissionPrivate struct {
	native uintptr
}

var pollableInputStreamInterfaceStruct *gi.Struct
var pollableInputStreamInterfaceStruct_Once sync.Once

func pollableInputStreamInterfaceStruct_Set() {
	pollableInputStreamInterfaceStruct_Once.Do(func() {
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
var pollableOutputStreamInterfaceStruct_Once sync.Once

func pollableOutputStreamInterfaceStruct_Set() {
	pollableOutputStreamInterfaceStruct_Once.Do(func() {
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
var proxyAddressClassStruct_Once sync.Once

func proxyAddressClassStruct_Set() {
	proxyAddressClassStruct_Once.Do(func() {
		proxyAddressClassStruct = gi.StructNew("Gio", "ProxyAddressClass")
	})
}

type ProxyAddressClass struct {
	native      uintptr
	ParentClass *InetSocketAddressClass
}

var proxyAddressEnumeratorClassStruct *gi.Struct
var proxyAddressEnumeratorClassStruct_Once sync.Once

func proxyAddressEnumeratorClassStruct_Set() {
	proxyAddressEnumeratorClassStruct_Once.Do(func() {
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
var proxyAddressEnumeratorPrivateStruct_Once sync.Once

func proxyAddressEnumeratorPrivateStruct_Set() {
	proxyAddressEnumeratorPrivateStruct_Once.Do(func() {
		proxyAddressEnumeratorPrivateStruct = gi.StructNew("Gio", "ProxyAddressEnumeratorPrivate")
	})
}

type ProxyAddressEnumeratorPrivate struct {
	native uintptr
}

var proxyAddressPrivateStruct *gi.Struct
var proxyAddressPrivateStruct_Once sync.Once

func proxyAddressPrivateStruct_Set() {
	proxyAddressPrivateStruct_Once.Do(func() {
		proxyAddressPrivateStruct = gi.StructNew("Gio", "ProxyAddressPrivate")
	})
}

type ProxyAddressPrivate struct {
	native uintptr
}

var proxyInterfaceStruct *gi.Struct
var proxyInterfaceStruct_Once sync.Once

func proxyInterfaceStruct_Set() {
	proxyInterfaceStruct_Once.Do(func() {
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
var proxyResolverInterfaceStruct_Once sync.Once

func proxyResolverInterfaceStruct_Set() {
	proxyResolverInterfaceStruct_Once.Do(func() {
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
var remoteActionGroupInterfaceStruct_Once sync.Once

func remoteActionGroupInterfaceStruct_Set() {
	remoteActionGroupInterfaceStruct_Once.Do(func() {
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
var resolverClassStruct_Once sync.Once

func resolverClassStruct_Set() {
	resolverClassStruct_Once.Do(func() {
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
var resolverPrivateStruct_Once sync.Once

func resolverPrivateStruct_Set() {
	resolverPrivateStruct_Once.Do(func() {
		resolverPrivateStruct = gi.StructNew("Gio", "ResolverPrivate")
	})
}

type ResolverPrivate struct {
	native uintptr
}

var resourceStruct *gi.Struct
var resourceStruct_Once sync.Once

func resourceStruct_Set() {
	resourceStruct_Once.Do(func() {
		resourceStruct = gi.StructNew("Gio", "Resource")
	})
}

type Resource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_resource_new_from_data' : parameter 'data' of type 'GLib.Bytes' not supported

var resourceRegisterFunction *gi.Function
var resourceRegisterFunction_Once sync.Once

func resourceRegisterFunction_Set() {
	resourceRegisterFunction_Once.Do(func() {
		resourceStruct_Set()
		resourceRegisterFunction = resourceStruct.InvokerNew("_register")
	})
}

// Register is a representation of the C type g_resources_register.
func (recv *Resource) Register() {
	resourceRegisterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	resourceRegisterFunction.Invoke(inArgs[:], nil)

}

var resourceUnregisterFunction *gi.Function
var resourceUnregisterFunction_Once sync.Once

func resourceUnregisterFunction_Set() {
	resourceUnregisterFunction_Once.Do(func() {
		resourceStruct_Set()
		resourceUnregisterFunction = resourceStruct.InvokerNew("_unregister")
	})
}

// Unregister is a representation of the C type g_resources_unregister.
func (recv *Resource) Unregister() {
	resourceUnregisterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	resourceUnregisterFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_resource_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var resourceRefFunction *gi.Function
var resourceRefFunction_Once sync.Once

func resourceRefFunction_Set() {
	resourceRefFunction_Once.Do(func() {
		resourceStruct_Set()
		resourceRefFunction = resourceStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_resource_ref.
func (recv *Resource) Ref() *Resource {
	resourceRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := resourceRefFunction.Invoke(inArgs[:], nil)

	retGo := &Resource{native: ret.Pointer()}

	return retGo
}

var resourceUnrefFunction *gi.Function
var resourceUnrefFunction_Once sync.Once

func resourceUnrefFunction_Set() {
	resourceUnrefFunction_Once.Do(func() {
		resourceStruct_Set()
		resourceUnrefFunction = resourceStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_resource_unref.
func (recv *Resource) Unref() {
	resourceUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	resourceUnrefFunction.Invoke(inArgs[:], nil)

}

var seekableIfaceStruct *gi.Struct
var seekableIfaceStruct_Once sync.Once

func seekableIfaceStruct_Set() {
	seekableIfaceStruct_Once.Do(func() {
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
var settingsBackendClassStruct_Once sync.Once

func settingsBackendClassStruct_Set() {
	settingsBackendClassStruct_Once.Do(func() {
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
var settingsBackendPrivateStruct_Once sync.Once

func settingsBackendPrivateStruct_Set() {
	settingsBackendPrivateStruct_Once.Do(func() {
		settingsBackendPrivateStruct = gi.StructNew("Gio", "SettingsBackendPrivate")
	})
}

type SettingsBackendPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() {
	settingsClassStruct_Once.Do(func() {
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
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() {
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct = gi.StructNew("Gio", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var settingsSchemaStruct *gi.Struct
var settingsSchemaStruct_Once sync.Once

func settingsSchemaStruct_Set() {
	settingsSchemaStruct_Once.Do(func() {
		settingsSchemaStruct = gi.StructNew("Gio", "SettingsSchema")
	})
}

type SettingsSchema struct {
	native uintptr
}

var settingsSchemaGetIdFunction *gi.Function
var settingsSchemaGetIdFunction_Once sync.Once

func settingsSchemaGetIdFunction_Set() {
	settingsSchemaGetIdFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaGetIdFunction = settingsSchemaStruct.InvokerNew("get_id")
	})
}

// GetId is a representation of the C type g_settings_schema_get_id.
func (recv *SettingsSchema) GetId() string {
	settingsSchemaGetIdFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaGetIdFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var settingsSchemaGetKeyFunction *gi.Function
var settingsSchemaGetKeyFunction_Once sync.Once

func settingsSchemaGetKeyFunction_Set() {
	settingsSchemaGetKeyFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaGetKeyFunction = settingsSchemaStruct.InvokerNew("get_key")
	})
}

// GetKey is a representation of the C type g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	settingsSchemaGetKeyFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := settingsSchemaGetKeyFunction.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo
}

var settingsSchemaGetPathFunction *gi.Function
var settingsSchemaGetPathFunction_Once sync.Once

func settingsSchemaGetPathFunction_Set() {
	settingsSchemaGetPathFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaGetPathFunction = settingsSchemaStruct.InvokerNew("get_path")
	})
}

// GetPath is a representation of the C type g_settings_schema_get_path.
func (recv *SettingsSchema) GetPath() string {
	settingsSchemaGetPathFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaGetPathFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_has_key' : return type 'gboolean' not supported

var settingsSchemaListChildrenFunction *gi.Function
var settingsSchemaListChildrenFunction_Once sync.Once

func settingsSchemaListChildrenFunction_Set() {
	settingsSchemaListChildrenFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaListChildrenFunction = settingsSchemaStruct.InvokerNew("list_children")
	})
}

// ListChildren is a representation of the C type g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() {
	settingsSchemaListChildrenFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	settingsSchemaListChildrenFunction.Invoke(inArgs[:], nil)

}

var settingsSchemaListKeysFunction *gi.Function
var settingsSchemaListKeysFunction_Once sync.Once

func settingsSchemaListKeysFunction_Set() {
	settingsSchemaListKeysFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaListKeysFunction = settingsSchemaStruct.InvokerNew("list_keys")
	})
}

// ListKeys is a representation of the C type g_settings_schema_list_keys.
func (recv *SettingsSchema) ListKeys() {
	settingsSchemaListKeysFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	settingsSchemaListKeysFunction.Invoke(inArgs[:], nil)

}

var settingsSchemaRefFunction *gi.Function
var settingsSchemaRefFunction_Once sync.Once

func settingsSchemaRefFunction_Set() {
	settingsSchemaRefFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaRefFunction = settingsSchemaStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_settings_schema_ref.
func (recv *SettingsSchema) Ref() *SettingsSchema {
	settingsSchemaRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaRefFunction.Invoke(inArgs[:], nil)

	retGo := &SettingsSchema{native: ret.Pointer()}

	return retGo
}

var settingsSchemaUnrefFunction *gi.Function
var settingsSchemaUnrefFunction_Once sync.Once

func settingsSchemaUnrefFunction_Set() {
	settingsSchemaUnrefFunction_Once.Do(func() {
		settingsSchemaStruct_Set()
		settingsSchemaUnrefFunction = settingsSchemaStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_settings_schema_unref.
func (recv *SettingsSchema) Unref() {
	settingsSchemaUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	settingsSchemaUnrefFunction.Invoke(inArgs[:], nil)

}

var settingsSchemaKeyStruct *gi.Struct
var settingsSchemaKeyStruct_Once sync.Once

func settingsSchemaKeyStruct_Set() {
	settingsSchemaKeyStruct_Once.Do(func() {
		settingsSchemaKeyStruct = gi.StructNew("Gio", "SettingsSchemaKey")
	})
}

type SettingsSchemaKey struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_default_value' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetDescriptionFunction *gi.Function
var settingsSchemaKeyGetDescriptionFunction_Once sync.Once

func settingsSchemaKeyGetDescriptionFunction_Set() {
	settingsSchemaKeyGetDescriptionFunction_Once.Do(func() {
		settingsSchemaKeyStruct_Set()
		settingsSchemaKeyGetDescriptionFunction = settingsSchemaKeyStruct.InvokerNew("get_description")
	})
}

// GetDescription is a representation of the C type g_settings_schema_key_get_description.
func (recv *SettingsSchemaKey) GetDescription() string {
	settingsSchemaKeyGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaKeyGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var settingsSchemaKeyGetNameFunction *gi.Function
var settingsSchemaKeyGetNameFunction_Once sync.Once

func settingsSchemaKeyGetNameFunction_Set() {
	settingsSchemaKeyGetNameFunction_Once.Do(func() {
		settingsSchemaKeyStruct_Set()
		settingsSchemaKeyGetNameFunction = settingsSchemaKeyStruct.InvokerNew("get_name")
	})
}

// GetName is a representation of the C type g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() string {
	settingsSchemaKeyGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaKeyGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_range' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetSummaryFunction *gi.Function
var settingsSchemaKeyGetSummaryFunction_Once sync.Once

func settingsSchemaKeyGetSummaryFunction_Set() {
	settingsSchemaKeyGetSummaryFunction_Once.Do(func() {
		settingsSchemaKeyStruct_Set()
		settingsSchemaKeyGetSummaryFunction = settingsSchemaKeyStruct.InvokerNew("get_summary")
	})
}

// GetSummary is a representation of the C type g_settings_schema_key_get_summary.
func (recv *SettingsSchemaKey) GetSummary() string {
	settingsSchemaKeyGetSummaryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaKeyGetSummaryFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_value_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_settings_schema_key_range_check' : parameter 'value' of type 'GLib.Variant' not supported

var settingsSchemaKeyRefFunction *gi.Function
var settingsSchemaKeyRefFunction_Once sync.Once

func settingsSchemaKeyRefFunction_Set() {
	settingsSchemaKeyRefFunction_Once.Do(func() {
		settingsSchemaKeyStruct_Set()
		settingsSchemaKeyRefFunction = settingsSchemaKeyStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	settingsSchemaKeyRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaKeyRefFunction.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo
}

var settingsSchemaKeyUnrefFunction *gi.Function
var settingsSchemaKeyUnrefFunction_Once sync.Once

func settingsSchemaKeyUnrefFunction_Set() {
	settingsSchemaKeyUnrefFunction_Once.Do(func() {
		settingsSchemaKeyStruct_Set()
		settingsSchemaKeyUnrefFunction = settingsSchemaKeyStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	settingsSchemaKeyUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	settingsSchemaKeyUnrefFunction.Invoke(inArgs[:], nil)

}

var settingsSchemaSourceStruct *gi.Struct
var settingsSchemaSourceStruct_Once sync.Once

func settingsSchemaSourceStruct_Set() {
	settingsSchemaSourceStruct_Once.Do(func() {
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
var settingsSchemaSourceRefFunction_Once sync.Once

func settingsSchemaSourceRefFunction_Set() {
	settingsSchemaSourceRefFunction_Once.Do(func() {
		settingsSchemaSourceStruct_Set()
		settingsSchemaSourceRefFunction = settingsSchemaSourceStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type g_settings_schema_source_ref.
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	settingsSchemaSourceRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := settingsSchemaSourceRefFunction.Invoke(inArgs[:], nil)

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo
}

var settingsSchemaSourceUnrefFunction *gi.Function
var settingsSchemaSourceUnrefFunction_Once sync.Once

func settingsSchemaSourceUnrefFunction_Set() {
	settingsSchemaSourceUnrefFunction_Once.Do(func() {
		settingsSchemaSourceStruct_Set()
		settingsSchemaSourceUnrefFunction = settingsSchemaSourceStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type g_settings_schema_source_unref.
func (recv *SettingsSchemaSource) Unref() {
	settingsSchemaSourceUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	settingsSchemaSourceUnrefFunction.Invoke(inArgs[:], nil)

}

var simpleActionGroupClassStruct *gi.Struct
var simpleActionGroupClassStruct_Once sync.Once

func simpleActionGroupClassStruct_Set() {
	simpleActionGroupClassStruct_Once.Do(func() {
		simpleActionGroupClassStruct = gi.StructNew("Gio", "SimpleActionGroupClass")
	})
}

type SimpleActionGroupClass struct {
	native uintptr
}

var simpleActionGroupPrivateStruct *gi.Struct
var simpleActionGroupPrivateStruct_Once sync.Once

func simpleActionGroupPrivateStruct_Set() {
	simpleActionGroupPrivateStruct_Once.Do(func() {
		simpleActionGroupPrivateStruct = gi.StructNew("Gio", "SimpleActionGroupPrivate")
	})
}

type SimpleActionGroupPrivate struct {
	native uintptr
}

var simpleAsyncResultClassStruct *gi.Struct
var simpleAsyncResultClassStruct_Once sync.Once

func simpleAsyncResultClassStruct_Set() {
	simpleAsyncResultClassStruct_Once.Do(func() {
		simpleAsyncResultClassStruct = gi.StructNew("Gio", "SimpleAsyncResultClass")
	})
}

type SimpleAsyncResultClass struct {
	native uintptr
}

var simpleProxyResolverClassStruct *gi.Struct
var simpleProxyResolverClassStruct_Once sync.Once

func simpleProxyResolverClassStruct_Set() {
	simpleProxyResolverClassStruct_Once.Do(func() {
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
var simpleProxyResolverPrivateStruct_Once sync.Once

func simpleProxyResolverPrivateStruct_Set() {
	simpleProxyResolverPrivateStruct_Once.Do(func() {
		simpleProxyResolverPrivateStruct = gi.StructNew("Gio", "SimpleProxyResolverPrivate")
	})
}

type SimpleProxyResolverPrivate struct {
	native uintptr
}

var socketAddressClassStruct *gi.Struct
var socketAddressClassStruct_Once sync.Once

func socketAddressClassStruct_Set() {
	socketAddressClassStruct_Once.Do(func() {
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
var socketAddressEnumeratorClassStruct_Once sync.Once

func socketAddressEnumeratorClassStruct_Set() {
	socketAddressEnumeratorClassStruct_Once.Do(func() {
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
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() {
	socketClassStruct_Once.Do(func() {
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
var socketClientClassStruct_Once sync.Once

func socketClientClassStruct_Set() {
	socketClientClassStruct_Once.Do(func() {
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
var socketClientPrivateStruct_Once sync.Once

func socketClientPrivateStruct_Set() {
	socketClientPrivateStruct_Once.Do(func() {
		socketClientPrivateStruct = gi.StructNew("Gio", "SocketClientPrivate")
	})
}

type SocketClientPrivate struct {
	native uintptr
}

var socketConnectableIfaceStruct *gi.Struct
var socketConnectableIfaceStruct_Once sync.Once

func socketConnectableIfaceStruct_Set() {
	socketConnectableIfaceStruct_Once.Do(func() {
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
var socketConnectionClassStruct_Once sync.Once

func socketConnectionClassStruct_Set() {
	socketConnectionClassStruct_Once.Do(func() {
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
var socketConnectionPrivateStruct_Once sync.Once

func socketConnectionPrivateStruct_Set() {
	socketConnectionPrivateStruct_Once.Do(func() {
		socketConnectionPrivateStruct = gi.StructNew("Gio", "SocketConnectionPrivate")
	})
}

type SocketConnectionPrivate struct {
	native uintptr
}

var socketControlMessageClassStruct *gi.Struct
var socketControlMessageClassStruct_Once sync.Once

func socketControlMessageClassStruct_Set() {
	socketControlMessageClassStruct_Once.Do(func() {
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
var socketControlMessagePrivateStruct_Once sync.Once

func socketControlMessagePrivateStruct_Set() {
	socketControlMessagePrivateStruct_Once.Do(func() {
		socketControlMessagePrivateStruct = gi.StructNew("Gio", "SocketControlMessagePrivate")
	})
}

type SocketControlMessagePrivate struct {
	native uintptr
}

var socketListenerClassStruct *gi.Struct
var socketListenerClassStruct_Once sync.Once

func socketListenerClassStruct_Set() {
	socketListenerClassStruct_Once.Do(func() {
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
var socketListenerPrivateStruct_Once sync.Once

func socketListenerPrivateStruct_Set() {
	socketListenerPrivateStruct_Once.Do(func() {
		socketListenerPrivateStruct = gi.StructNew("Gio", "SocketListenerPrivate")
	})
}

type SocketListenerPrivate struct {
	native uintptr
}

var socketPrivateStruct *gi.Struct
var socketPrivateStruct_Once sync.Once

func socketPrivateStruct_Set() {
	socketPrivateStruct_Once.Do(func() {
		socketPrivateStruct = gi.StructNew("Gio", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var socketServiceClassStruct *gi.Struct
var socketServiceClassStruct_Once sync.Once

func socketServiceClassStruct_Set() {
	socketServiceClassStruct_Once.Do(func() {
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
var socketServicePrivateStruct_Once sync.Once

func socketServicePrivateStruct_Set() {
	socketServicePrivateStruct_Once.Do(func() {
		socketServicePrivateStruct = gi.StructNew("Gio", "SocketServicePrivate")
	})
}

type SocketServicePrivate struct {
	native uintptr
}

var srvTargetStruct *gi.Struct
var srvTargetStruct_Once sync.Once

func srvTargetStruct_Set() {
	srvTargetStruct_Once.Do(func() {
		srvTargetStruct = gi.StructNew("Gio", "SrvTarget")
	})
}

type SrvTarget struct {
	native uintptr
}

var srvTargetNewFunction *gi.Function
var srvTargetNewFunction_Once sync.Once

func srvTargetNewFunction_Set() {
	srvTargetNewFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetNewFunction = srvTargetStruct.InvokerNew("new")
	})
}

// SrvTargetNew is a representation of the C type g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	srvTargetNewFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)
	inArgs[2].SetUint16(priority)
	inArgs[3].SetUint16(weight)

	ret := srvTargetNewFunction.Invoke(inArgs[:], nil)

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo
}

var srvTargetCopyFunction *gi.Function
var srvTargetCopyFunction_Once sync.Once

func srvTargetCopyFunction_Set() {
	srvTargetCopyFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetCopyFunction = srvTargetStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type g_srv_target_copy.
func (recv *SrvTarget) Copy() *SrvTarget {
	srvTargetCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := srvTargetCopyFunction.Invoke(inArgs[:], nil)

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo
}

var srvTargetFreeFunction *gi.Function
var srvTargetFreeFunction_Once sync.Once

func srvTargetFreeFunction_Set() {
	srvTargetFreeFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetFreeFunction = srvTargetStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type g_srv_target_free.
func (recv *SrvTarget) Free() {
	srvTargetFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	srvTargetFreeFunction.Invoke(inArgs[:], nil)

}

var srvTargetGetHostnameFunction *gi.Function
var srvTargetGetHostnameFunction_Once sync.Once

func srvTargetGetHostnameFunction_Set() {
	srvTargetGetHostnameFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetGetHostnameFunction = srvTargetStruct.InvokerNew("get_hostname")
	})
}

// GetHostname is a representation of the C type g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() string {
	srvTargetGetHostnameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := srvTargetGetHostnameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var srvTargetGetPortFunction *gi.Function
var srvTargetGetPortFunction_Once sync.Once

func srvTargetGetPortFunction_Set() {
	srvTargetGetPortFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetGetPortFunction = srvTargetStruct.InvokerNew("get_port")
	})
}

// GetPort is a representation of the C type g_srv_target_get_port.
func (recv *SrvTarget) GetPort() uint16 {
	srvTargetGetPortFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := srvTargetGetPortFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var srvTargetGetPriorityFunction *gi.Function
var srvTargetGetPriorityFunction_Once sync.Once

func srvTargetGetPriorityFunction_Set() {
	srvTargetGetPriorityFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetGetPriorityFunction = srvTargetStruct.InvokerNew("get_priority")
	})
}

// GetPriority is a representation of the C type g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() uint16 {
	srvTargetGetPriorityFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := srvTargetGetPriorityFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var srvTargetGetWeightFunction *gi.Function
var srvTargetGetWeightFunction_Once sync.Once

func srvTargetGetWeightFunction_Set() {
	srvTargetGetWeightFunction_Once.Do(func() {
		srvTargetStruct_Set()
		srvTargetGetWeightFunction = srvTargetStruct.InvokerNew("get_weight")
	})
}

// GetWeight is a representation of the C type g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() uint16 {
	srvTargetGetWeightFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := srvTargetGetWeightFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var staticResourceStruct *gi.Struct
var staticResourceStruct_Once sync.Once

func staticResourceStruct_Set() {
	staticResourceStruct_Once.Do(func() {
		staticResourceStruct = gi.StructNew("Gio", "StaticResource")
	})
}

type StaticResource struct {
	native uintptr
}

var staticResourceFiniFunction *gi.Function
var staticResourceFiniFunction_Once sync.Once

func staticResourceFiniFunction_Set() {
	staticResourceFiniFunction_Once.Do(func() {
		staticResourceStruct_Set()
		staticResourceFiniFunction = staticResourceStruct.InvokerNew("fini")
	})
}

// Fini is a representation of the C type g_static_resource_fini.
func (recv *StaticResource) Fini() {
	staticResourceFiniFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	staticResourceFiniFunction.Invoke(inArgs[:], nil)

}

var staticResourceGetResourceFunction *gi.Function
var staticResourceGetResourceFunction_Once sync.Once

func staticResourceGetResourceFunction_Set() {
	staticResourceGetResourceFunction_Once.Do(func() {
		staticResourceStruct_Set()
		staticResourceGetResourceFunction = staticResourceStruct.InvokerNew("get_resource")
	})
}

// GetResource is a representation of the C type g_static_resource_get_resource.
func (recv *StaticResource) GetResource() *Resource {
	staticResourceGetResourceFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := staticResourceGetResourceFunction.Invoke(inArgs[:], nil)

	retGo := &Resource{native: ret.Pointer()}

	return retGo
}

var staticResourceInitFunction *gi.Function
var staticResourceInitFunction_Once sync.Once

func staticResourceInitFunction_Set() {
	staticResourceInitFunction_Once.Do(func() {
		staticResourceStruct_Set()
		staticResourceInitFunction = staticResourceStruct.InvokerNew("init")
	})
}

// Init is a representation of the C type g_static_resource_init.
func (recv *StaticResource) Init() {
	staticResourceInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	staticResourceInitFunction.Invoke(inArgs[:], nil)

}

var taskClassStruct *gi.Struct
var taskClassStruct_Once sync.Once

func taskClassStruct_Set() {
	taskClassStruct_Once.Do(func() {
		taskClassStruct = gi.StructNew("Gio", "TaskClass")
	})
}

type TaskClass struct {
	native uintptr
}

var tcpConnectionClassStruct *gi.Struct
var tcpConnectionClassStruct_Once sync.Once

func tcpConnectionClassStruct_Set() {
	tcpConnectionClassStruct_Once.Do(func() {
		tcpConnectionClassStruct = gi.StructNew("Gio", "TcpConnectionClass")
	})
}

type TcpConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var tcpConnectionPrivateStruct *gi.Struct
var tcpConnectionPrivateStruct_Once sync.Once

func tcpConnectionPrivateStruct_Set() {
	tcpConnectionPrivateStruct_Once.Do(func() {
		tcpConnectionPrivateStruct = gi.StructNew("Gio", "TcpConnectionPrivate")
	})
}

type TcpConnectionPrivate struct {
	native uintptr
}

var tcpWrapperConnectionClassStruct *gi.Struct
var tcpWrapperConnectionClassStruct_Once sync.Once

func tcpWrapperConnectionClassStruct_Set() {
	tcpWrapperConnectionClassStruct_Once.Do(func() {
		tcpWrapperConnectionClassStruct = gi.StructNew("Gio", "TcpWrapperConnectionClass")
	})
}

type TcpWrapperConnectionClass struct {
	native      uintptr
	ParentClass *TcpConnectionClass
}

var tcpWrapperConnectionPrivateStruct *gi.Struct
var tcpWrapperConnectionPrivateStruct_Once sync.Once

func tcpWrapperConnectionPrivateStruct_Set() {
	tcpWrapperConnectionPrivateStruct_Once.Do(func() {
		tcpWrapperConnectionPrivateStruct = gi.StructNew("Gio", "TcpWrapperConnectionPrivate")
	})
}

type TcpWrapperConnectionPrivate struct {
	native uintptr
}

var themedIconClassStruct *gi.Struct
var themedIconClassStruct_Once sync.Once

func themedIconClassStruct_Set() {
	themedIconClassStruct_Once.Do(func() {
		themedIconClassStruct = gi.StructNew("Gio", "ThemedIconClass")
	})
}

type ThemedIconClass struct {
	native uintptr
}

var threadedSocketServiceClassStruct *gi.Struct
var threadedSocketServiceClassStruct_Once sync.Once

func threadedSocketServiceClassStruct_Set() {
	threadedSocketServiceClassStruct_Once.Do(func() {
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
var threadedSocketServicePrivateStruct_Once sync.Once

func threadedSocketServicePrivateStruct_Set() {
	threadedSocketServicePrivateStruct_Once.Do(func() {
		threadedSocketServicePrivateStruct = gi.StructNew("Gio", "ThreadedSocketServicePrivate")
	})
}

type ThreadedSocketServicePrivate struct {
	native uintptr
}

var tlsBackendInterfaceStruct *gi.Struct
var tlsBackendInterfaceStruct_Once sync.Once

func tlsBackendInterfaceStruct_Set() {
	tlsBackendInterfaceStruct_Once.Do(func() {
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
var tlsCertificateClassStruct_Once sync.Once

func tlsCertificateClassStruct_Set() {
	tlsCertificateClassStruct_Once.Do(func() {
		tlsCertificateClassStruct = gi.StructNew("Gio", "TlsCertificateClass")
	})
}

type TlsCertificateClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'verify' : missing Type
}

var tlsCertificatePrivateStruct *gi.Struct
var tlsCertificatePrivateStruct_Once sync.Once

func tlsCertificatePrivateStruct_Set() {
	tlsCertificatePrivateStruct_Once.Do(func() {
		tlsCertificatePrivateStruct = gi.StructNew("Gio", "TlsCertificatePrivate")
	})
}

type TlsCertificatePrivate struct {
	native uintptr
}

var tlsClientConnectionInterfaceStruct *gi.Struct
var tlsClientConnectionInterfaceStruct_Once sync.Once

func tlsClientConnectionInterfaceStruct_Set() {
	tlsClientConnectionInterfaceStruct_Once.Do(func() {
		tlsClientConnectionInterfaceStruct = gi.StructNew("Gio", "TlsClientConnectionInterface")
	})
}

type TlsClientConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'copy_session_state' : missing Type
}

var tlsConnectionClassStruct *gi.Struct
var tlsConnectionClassStruct_Once sync.Once

func tlsConnectionClassStruct_Set() {
	tlsConnectionClassStruct_Once.Do(func() {
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
var tlsConnectionPrivateStruct_Once sync.Once

func tlsConnectionPrivateStruct_Set() {
	tlsConnectionPrivateStruct_Once.Do(func() {
		tlsConnectionPrivateStruct = gi.StructNew("Gio", "TlsConnectionPrivate")
	})
}

type TlsConnectionPrivate struct {
	native uintptr
}

var tlsDatabaseClassStruct *gi.Struct
var tlsDatabaseClassStruct_Once sync.Once

func tlsDatabaseClassStruct_Set() {
	tlsDatabaseClassStruct_Once.Do(func() {
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
var tlsDatabasePrivateStruct_Once sync.Once

func tlsDatabasePrivateStruct_Set() {
	tlsDatabasePrivateStruct_Once.Do(func() {
		tlsDatabasePrivateStruct = gi.StructNew("Gio", "TlsDatabasePrivate")
	})
}

type TlsDatabasePrivate struct {
	native uintptr
}

var tlsFileDatabaseInterfaceStruct *gi.Struct
var tlsFileDatabaseInterfaceStruct_Once sync.Once

func tlsFileDatabaseInterfaceStruct_Set() {
	tlsFileDatabaseInterfaceStruct_Once.Do(func() {
		tlsFileDatabaseInterfaceStruct = gi.StructNew("Gio", "TlsFileDatabaseInterface")
	})
}

type TlsFileDatabaseInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var tlsInteractionClassStruct *gi.Struct
var tlsInteractionClassStruct_Once sync.Once

func tlsInteractionClassStruct_Set() {
	tlsInteractionClassStruct_Once.Do(func() {
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
var tlsInteractionPrivateStruct_Once sync.Once

func tlsInteractionPrivateStruct_Set() {
	tlsInteractionPrivateStruct_Once.Do(func() {
		tlsInteractionPrivateStruct = gi.StructNew("Gio", "TlsInteractionPrivate")
	})
}

type TlsInteractionPrivate struct {
	native uintptr
}

var tlsPasswordClassStruct *gi.Struct
var tlsPasswordClassStruct_Once sync.Once

func tlsPasswordClassStruct_Set() {
	tlsPasswordClassStruct_Once.Do(func() {
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
var tlsPasswordPrivateStruct_Once sync.Once

func tlsPasswordPrivateStruct_Set() {
	tlsPasswordPrivateStruct_Once.Do(func() {
		tlsPasswordPrivateStruct = gi.StructNew("Gio", "TlsPasswordPrivate")
	})
}

type TlsPasswordPrivate struct {
	native uintptr
}

var tlsServerConnectionInterfaceStruct *gi.Struct
var tlsServerConnectionInterfaceStruct_Once sync.Once

func tlsServerConnectionInterfaceStruct_Set() {
	tlsServerConnectionInterfaceStruct_Once.Do(func() {
		tlsServerConnectionInterfaceStruct = gi.StructNew("Gio", "TlsServerConnectionInterface")
	})
}

type TlsServerConnectionInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
}

var unixConnectionClassStruct *gi.Struct
var unixConnectionClassStruct_Once sync.Once

func unixConnectionClassStruct_Set() {
	unixConnectionClassStruct_Once.Do(func() {
		unixConnectionClassStruct = gi.StructNew("Gio", "UnixConnectionClass")
	})
}

type UnixConnectionClass struct {
	native      uintptr
	ParentClass *SocketConnectionClass
}

var unixConnectionPrivateStruct *gi.Struct
var unixConnectionPrivateStruct_Once sync.Once

func unixConnectionPrivateStruct_Set() {
	unixConnectionPrivateStruct_Once.Do(func() {
		unixConnectionPrivateStruct = gi.StructNew("Gio", "UnixConnectionPrivate")
	})
}

type UnixConnectionPrivate struct {
	native uintptr
}

var unixCredentialsMessageClassStruct *gi.Struct
var unixCredentialsMessageClassStruct_Once sync.Once

func unixCredentialsMessageClassStruct_Set() {
	unixCredentialsMessageClassStruct_Once.Do(func() {
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
var unixCredentialsMessagePrivateStruct_Once sync.Once

func unixCredentialsMessagePrivateStruct_Set() {
	unixCredentialsMessagePrivateStruct_Once.Do(func() {
		unixCredentialsMessagePrivateStruct = gi.StructNew("Gio", "UnixCredentialsMessagePrivate")
	})
}

type UnixCredentialsMessagePrivate struct {
	native uintptr
}

var unixFDListClassStruct *gi.Struct
var unixFDListClassStruct_Once sync.Once

func unixFDListClassStruct_Set() {
	unixFDListClassStruct_Once.Do(func() {
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
var unixFDListPrivateStruct_Once sync.Once

func unixFDListPrivateStruct_Set() {
	unixFDListPrivateStruct_Once.Do(func() {
		unixFDListPrivateStruct = gi.StructNew("Gio", "UnixFDListPrivate")
	})
}

type UnixFDListPrivate struct {
	native uintptr
}

var unixFDMessageClassStruct *gi.Struct
var unixFDMessageClassStruct_Once sync.Once

func unixFDMessageClassStruct_Set() {
	unixFDMessageClassStruct_Once.Do(func() {
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
var unixFDMessagePrivateStruct_Once sync.Once

func unixFDMessagePrivateStruct_Set() {
	unixFDMessagePrivateStruct_Once.Do(func() {
		unixFDMessagePrivateStruct = gi.StructNew("Gio", "UnixFDMessagePrivate")
	})
}

type UnixFDMessagePrivate struct {
	native uintptr
}

var unixInputStreamClassStruct *gi.Struct
var unixInputStreamClassStruct_Once sync.Once

func unixInputStreamClassStruct_Set() {
	unixInputStreamClassStruct_Once.Do(func() {
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
var unixInputStreamPrivateStruct_Once sync.Once

func unixInputStreamPrivateStruct_Set() {
	unixInputStreamPrivateStruct_Once.Do(func() {
		unixInputStreamPrivateStruct = gi.StructNew("Gio", "UnixInputStreamPrivate")
	})
}

type UnixInputStreamPrivate struct {
	native uintptr
}

var unixMountEntryStruct *gi.Struct
var unixMountEntryStruct_Once sync.Once

func unixMountEntryStruct_Set() {
	unixMountEntryStruct_Once.Do(func() {
		unixMountEntryStruct = gi.StructNew("Gio", "UnixMountEntry")
	})
}

type UnixMountEntry struct {
	native uintptr
}

var unixMountMonitorClassStruct *gi.Struct
var unixMountMonitorClassStruct_Once sync.Once

func unixMountMonitorClassStruct_Set() {
	unixMountMonitorClassStruct_Once.Do(func() {
		unixMountMonitorClassStruct = gi.StructNew("Gio", "UnixMountMonitorClass")
	})
}

type UnixMountMonitorClass struct {
	native uintptr
}

var unixMountPointStruct *gi.Struct
var unixMountPointStruct_Once sync.Once

func unixMountPointStruct_Set() {
	unixMountPointStruct_Once.Do(func() {
		unixMountPointStruct = gi.StructNew("Gio", "UnixMountPoint")
	})
}

type UnixMountPoint struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_unix_mount_point_compare' : parameter 'mount2' of type 'UnixMountPoint' not supported

var unixMountPointCopyFunction *gi.Function
var unixMountPointCopyFunction_Once sync.Once

func unixMountPointCopyFunction_Set() {
	unixMountPointCopyFunction_Once.Do(func() {
		unixMountPointStruct_Set()
		unixMountPointCopyFunction = unixMountPointStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type g_unix_mount_point_copy.
func (recv *UnixMountPoint) Copy() *UnixMountPoint {
	unixMountPointCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unixMountPointCopyFunction.Invoke(inArgs[:], nil)

	retGo := &UnixMountPoint{native: ret.Pointer()}

	return retGo
}

var unixMountPointFreeFunction *gi.Function
var unixMountPointFreeFunction_Once sync.Once

func unixMountPointFreeFunction_Set() {
	unixMountPointFreeFunction_Once.Do(func() {
		unixMountPointStruct_Set()
		unixMountPointFreeFunction = unixMountPointStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type g_unix_mount_point_free.
func (recv *UnixMountPoint) Free() {
	unixMountPointFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unixMountPointFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_unix_mount_point_get_device_path' : return type 'filename' not supported

var unixMountPointGetFsTypeFunction *gi.Function
var unixMountPointGetFsTypeFunction_Once sync.Once

func unixMountPointGetFsTypeFunction_Set() {
	unixMountPointGetFsTypeFunction_Once.Do(func() {
		unixMountPointStruct_Set()
		unixMountPointGetFsTypeFunction = unixMountPointStruct.InvokerNew("get_fs_type")
	})
}

// GetFsType is a representation of the C type g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() string {
	unixMountPointGetFsTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unixMountPointGetFsTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_get_mount_path' : return type 'filename' not supported

var unixMountPointGetOptionsFunction *gi.Function
var unixMountPointGetOptionsFunction_Once sync.Once

func unixMountPointGetOptionsFunction_Set() {
	unixMountPointGetOptionsFunction_Once.Do(func() {
		unixMountPointStruct_Set()
		unixMountPointGetOptionsFunction = unixMountPointStruct.InvokerNew("get_options")
	})
}

// GetOptions is a representation of the C type g_unix_mount_point_get_options.
func (recv *UnixMountPoint) GetOptions() string {
	unixMountPointGetOptionsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unixMountPointGetOptionsFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_can_eject' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_guess_icon' : return type 'Icon' not supported

var unixMountPointGuessNameFunction *gi.Function
var unixMountPointGuessNameFunction_Once sync.Once

func unixMountPointGuessNameFunction_Set() {
	unixMountPointGuessNameFunction_Once.Do(func() {
		unixMountPointStruct_Set()
		unixMountPointGuessNameFunction = unixMountPointStruct.InvokerNew("guess_name")
	})
}

// GuessName is a representation of the C type g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() string {
	unixMountPointGuessNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unixMountPointGuessNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_loopback' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_readonly' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_point_is_user_mountable' : return type 'gboolean' not supported

var unixOutputStreamClassStruct *gi.Struct
var unixOutputStreamClassStruct_Once sync.Once

func unixOutputStreamClassStruct_Set() {
	unixOutputStreamClassStruct_Once.Do(func() {
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
var unixOutputStreamPrivateStruct_Once sync.Once

func unixOutputStreamPrivateStruct_Set() {
	unixOutputStreamPrivateStruct_Once.Do(func() {
		unixOutputStreamPrivateStruct = gi.StructNew("Gio", "UnixOutputStreamPrivate")
	})
}

type UnixOutputStreamPrivate struct {
	native uintptr
}

var unixSocketAddressClassStruct *gi.Struct
var unixSocketAddressClassStruct_Once sync.Once

func unixSocketAddressClassStruct_Set() {
	unixSocketAddressClassStruct_Once.Do(func() {
		unixSocketAddressClassStruct = gi.StructNew("Gio", "UnixSocketAddressClass")
	})
}

type UnixSocketAddressClass struct {
	native      uintptr
	ParentClass *SocketAddressClass
}

var unixSocketAddressPrivateStruct *gi.Struct
var unixSocketAddressPrivateStruct_Once sync.Once

func unixSocketAddressPrivateStruct_Set() {
	unixSocketAddressPrivateStruct_Once.Do(func() {
		unixSocketAddressPrivateStruct = gi.StructNew("Gio", "UnixSocketAddressPrivate")
	})
}

type UnixSocketAddressPrivate struct {
	native uintptr
}

var vfsClassStruct *gi.Struct
var vfsClassStruct_Once sync.Once

func vfsClassStruct_Set() {
	vfsClassStruct_Once.Do(func() {
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
var volumeIfaceStruct_Once sync.Once

func volumeIfaceStruct_Set() {
	volumeIfaceStruct_Once.Do(func() {
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
var volumeMonitorClassStruct_Once sync.Once

func volumeMonitorClassStruct_Set() {
	volumeMonitorClassStruct_Once.Do(func() {
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
var zlibCompressorClassStruct_Once sync.Once

func zlibCompressorClassStruct_Set() {
	zlibCompressorClassStruct_Once.Do(func() {
		zlibCompressorClassStruct = gi.StructNew("Gio", "ZlibCompressorClass")
	})
}

type ZlibCompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var zlibDecompressorClassStruct *gi.Struct
var zlibDecompressorClassStruct_Once sync.Once

func zlibDecompressorClassStruct_Set() {
	zlibDecompressorClassStruct_Once.Do(func() {
		zlibDecompressorClassStruct = gi.StructNew("Gio", "ZlibDecompressorClass")
	})
}

type ZlibDecompressorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}
