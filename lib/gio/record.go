// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionEntryStruct *gi.Struct
var actionEntryStruct_Once sync.Once

func actionEntryStruct_Set() error {
	var err error
	actionEntryStruct_Once.Do(func() {
		actionEntryStruct, err = gi.StructNew("Gio", "ActionEntry")
	})
	return err
}

type ActionEntry struct {
	native uintptr
}

var actionGroupInterfaceStruct *gi.Struct
var actionGroupInterfaceStruct_Once sync.Once

func actionGroupInterfaceStruct_Set() error {
	var err error
	actionGroupInterfaceStruct_Once.Do(func() {
		actionGroupInterfaceStruct, err = gi.StructNew("Gio", "ActionGroupInterface")
	})
	return err
}

type ActionGroupInterface struct {
	native uintptr
}

var actionInterfaceStruct *gi.Struct
var actionInterfaceStruct_Once sync.Once

func actionInterfaceStruct_Set() error {
	var err error
	actionInterfaceStruct_Once.Do(func() {
		actionInterfaceStruct, err = gi.StructNew("Gio", "ActionInterface")
	})
	return err
}

type ActionInterface struct {
	native uintptr
}

var actionMapInterfaceStruct *gi.Struct
var actionMapInterfaceStruct_Once sync.Once

func actionMapInterfaceStruct_Set() error {
	var err error
	actionMapInterfaceStruct_Once.Do(func() {
		actionMapInterfaceStruct, err = gi.StructNew("Gio", "ActionMapInterface")
	})
	return err
}

type ActionMapInterface struct {
	native uintptr
}

var appInfoIfaceStruct *gi.Struct
var appInfoIfaceStruct_Once sync.Once

func appInfoIfaceStruct_Set() error {
	var err error
	appInfoIfaceStruct_Once.Do(func() {
		appInfoIfaceStruct, err = gi.StructNew("Gio", "AppInfoIface")
	})
	return err
}

type AppInfoIface struct {
	native uintptr
}

var appLaunchContextClassStruct *gi.Struct
var appLaunchContextClassStruct_Once sync.Once

func appLaunchContextClassStruct_Set() error {
	var err error
	appLaunchContextClassStruct_Once.Do(func() {
		appLaunchContextClassStruct, err = gi.StructNew("Gio", "AppLaunchContextClass")
	})
	return err
}

type AppLaunchContextClass struct {
	native uintptr
}

var appLaunchContextPrivateStruct *gi.Struct
var appLaunchContextPrivateStruct_Once sync.Once

func appLaunchContextPrivateStruct_Set() error {
	var err error
	appLaunchContextPrivateStruct_Once.Do(func() {
		appLaunchContextPrivateStruct, err = gi.StructNew("Gio", "AppLaunchContextPrivate")
	})
	return err
}

type AppLaunchContextPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStruct_Once sync.Once

func applicationClassStruct_Set() error {
	var err error
	applicationClassStruct_Once.Do(func() {
		applicationClassStruct, err = gi.StructNew("Gio", "ApplicationClass")
	})
	return err
}

type ApplicationClass struct {
	native uintptr
}

var applicationCommandLineClassStruct *gi.Struct
var applicationCommandLineClassStruct_Once sync.Once

func applicationCommandLineClassStruct_Set() error {
	var err error
	applicationCommandLineClassStruct_Once.Do(func() {
		applicationCommandLineClassStruct, err = gi.StructNew("Gio", "ApplicationCommandLineClass")
	})
	return err
}

type ApplicationCommandLineClass struct {
	native uintptr
}

var applicationCommandLinePrivateStruct *gi.Struct
var applicationCommandLinePrivateStruct_Once sync.Once

func applicationCommandLinePrivateStruct_Set() error {
	var err error
	applicationCommandLinePrivateStruct_Once.Do(func() {
		applicationCommandLinePrivateStruct, err = gi.StructNew("Gio", "ApplicationCommandLinePrivate")
	})
	return err
}

type ApplicationCommandLinePrivate struct {
	native uintptr
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStruct_Once sync.Once

func applicationPrivateStruct_Set() error {
	var err error
	applicationPrivateStruct_Once.Do(func() {
		applicationPrivateStruct, err = gi.StructNew("Gio", "ApplicationPrivate")
	})
	return err
}

type ApplicationPrivate struct {
	native uintptr
}

var asyncInitableIfaceStruct *gi.Struct
var asyncInitableIfaceStruct_Once sync.Once

func asyncInitableIfaceStruct_Set() error {
	var err error
	asyncInitableIfaceStruct_Once.Do(func() {
		asyncInitableIfaceStruct, err = gi.StructNew("Gio", "AsyncInitableIface")
	})
	return err
}

type AsyncInitableIface struct {
	native uintptr
}

var asyncResultIfaceStruct *gi.Struct
var asyncResultIfaceStruct_Once sync.Once

func asyncResultIfaceStruct_Set() error {
	var err error
	asyncResultIfaceStruct_Once.Do(func() {
		asyncResultIfaceStruct, err = gi.StructNew("Gio", "AsyncResultIface")
	})
	return err
}

type AsyncResultIface struct {
	native uintptr
}

var bufferedInputStreamClassStruct *gi.Struct
var bufferedInputStreamClassStruct_Once sync.Once

func bufferedInputStreamClassStruct_Set() error {
	var err error
	bufferedInputStreamClassStruct_Once.Do(func() {
		bufferedInputStreamClassStruct, err = gi.StructNew("Gio", "BufferedInputStreamClass")
	})
	return err
}

type BufferedInputStreamClass struct {
	native uintptr
}

var bufferedInputStreamPrivateStruct *gi.Struct
var bufferedInputStreamPrivateStruct_Once sync.Once

func bufferedInputStreamPrivateStruct_Set() error {
	var err error
	bufferedInputStreamPrivateStruct_Once.Do(func() {
		bufferedInputStreamPrivateStruct, err = gi.StructNew("Gio", "BufferedInputStreamPrivate")
	})
	return err
}

type BufferedInputStreamPrivate struct {
	native uintptr
}

var bufferedOutputStreamClassStruct *gi.Struct
var bufferedOutputStreamClassStruct_Once sync.Once

func bufferedOutputStreamClassStruct_Set() error {
	var err error
	bufferedOutputStreamClassStruct_Once.Do(func() {
		bufferedOutputStreamClassStruct, err = gi.StructNew("Gio", "BufferedOutputStreamClass")
	})
	return err
}

type BufferedOutputStreamClass struct {
	native uintptr
}

var bufferedOutputStreamPrivateStruct *gi.Struct
var bufferedOutputStreamPrivateStruct_Once sync.Once

func bufferedOutputStreamPrivateStruct_Set() error {
	var err error
	bufferedOutputStreamPrivateStruct_Once.Do(func() {
		bufferedOutputStreamPrivateStruct, err = gi.StructNew("Gio", "BufferedOutputStreamPrivate")
	})
	return err
}

type BufferedOutputStreamPrivate struct {
	native uintptr
}

var cancellableClassStruct *gi.Struct
var cancellableClassStruct_Once sync.Once

func cancellableClassStruct_Set() error {
	var err error
	cancellableClassStruct_Once.Do(func() {
		cancellableClassStruct, err = gi.StructNew("Gio", "CancellableClass")
	})
	return err
}

type CancellableClass struct {
	native uintptr
}

var cancellablePrivateStruct *gi.Struct
var cancellablePrivateStruct_Once sync.Once

func cancellablePrivateStruct_Set() error {
	var err error
	cancellablePrivateStruct_Once.Do(func() {
		cancellablePrivateStruct, err = gi.StructNew("Gio", "CancellablePrivate")
	})
	return err
}

type CancellablePrivate struct {
	native uintptr
}

var charsetConverterClassStruct *gi.Struct
var charsetConverterClassStruct_Once sync.Once

func charsetConverterClassStruct_Set() error {
	var err error
	charsetConverterClassStruct_Once.Do(func() {
		charsetConverterClassStruct, err = gi.StructNew("Gio", "CharsetConverterClass")
	})
	return err
}

type CharsetConverterClass struct {
	native uintptr
}

var converterIfaceStruct *gi.Struct
var converterIfaceStruct_Once sync.Once

func converterIfaceStruct_Set() error {
	var err error
	converterIfaceStruct_Once.Do(func() {
		converterIfaceStruct, err = gi.StructNew("Gio", "ConverterIface")
	})
	return err
}

type ConverterIface struct {
	native uintptr
}

var converterInputStreamClassStruct *gi.Struct
var converterInputStreamClassStruct_Once sync.Once

func converterInputStreamClassStruct_Set() error {
	var err error
	converterInputStreamClassStruct_Once.Do(func() {
		converterInputStreamClassStruct, err = gi.StructNew("Gio", "ConverterInputStreamClass")
	})
	return err
}

type ConverterInputStreamClass struct {
	native uintptr
}

var converterInputStreamPrivateStruct *gi.Struct
var converterInputStreamPrivateStruct_Once sync.Once

func converterInputStreamPrivateStruct_Set() error {
	var err error
	converterInputStreamPrivateStruct_Once.Do(func() {
		converterInputStreamPrivateStruct, err = gi.StructNew("Gio", "ConverterInputStreamPrivate")
	})
	return err
}

type ConverterInputStreamPrivate struct {
	native uintptr
}

var converterOutputStreamClassStruct *gi.Struct
var converterOutputStreamClassStruct_Once sync.Once

func converterOutputStreamClassStruct_Set() error {
	var err error
	converterOutputStreamClassStruct_Once.Do(func() {
		converterOutputStreamClassStruct, err = gi.StructNew("Gio", "ConverterOutputStreamClass")
	})
	return err
}

type ConverterOutputStreamClass struct {
	native uintptr
}

var converterOutputStreamPrivateStruct *gi.Struct
var converterOutputStreamPrivateStruct_Once sync.Once

func converterOutputStreamPrivateStruct_Set() error {
	var err error
	converterOutputStreamPrivateStruct_Once.Do(func() {
		converterOutputStreamPrivateStruct, err = gi.StructNew("Gio", "ConverterOutputStreamPrivate")
	})
	return err
}

type ConverterOutputStreamPrivate struct {
	native uintptr
}

var credentialsClassStruct *gi.Struct
var credentialsClassStruct_Once sync.Once

func credentialsClassStruct_Set() error {
	var err error
	credentialsClassStruct_Once.Do(func() {
		credentialsClassStruct, err = gi.StructNew("Gio", "CredentialsClass")
	})
	return err
}

type CredentialsClass struct {
	native uintptr
}

var dBusAnnotationInfoStruct *gi.Struct
var dBusAnnotationInfoStruct_Once sync.Once

func dBusAnnotationInfoStruct_Set() error {
	var err error
	dBusAnnotationInfoStruct_Once.Do(func() {
		dBusAnnotationInfoStruct, err = gi.StructNew("Gio", "DBusAnnotationInfo")
	})
	return err
}

type DBusAnnotationInfo struct {
	native uintptr
}

var dBusAnnotationInfoRefFunction *gi.Function
var dBusAnnotationInfoRefFunction_Once sync.Once

func dBusAnnotationInfoRefFunction_Set() error {
	var err error
	dBusAnnotationInfoRefFunction_Once.Do(func() {
		err = dBusAnnotationInfoStruct_Set()
		if err != nil {
			return
		}
		dBusAnnotationInfoRefFunction, err = dBusAnnotationInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_annotation_info_ref.
func (recv *DBusAnnotationInfo) Ref() (*DBusAnnotationInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusAnnotationInfoRefFunction_Set()
	if err == nil {
		ret = dBusAnnotationInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusAnnotationInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusAnnotationInfoUnrefFunction *gi.Function
var dBusAnnotationInfoUnrefFunction_Once sync.Once

func dBusAnnotationInfoUnrefFunction_Set() error {
	var err error
	dBusAnnotationInfoUnrefFunction_Once.Do(func() {
		err = dBusAnnotationInfoStruct_Set()
		if err != nil {
			return
		}
		dBusAnnotationInfoUnrefFunction, err = dBusAnnotationInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_annotation_info_unref.
func (recv *DBusAnnotationInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusAnnotationInfoUnrefFunction_Set()
	if err == nil {
		dBusAnnotationInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusArgInfoStruct *gi.Struct
var dBusArgInfoStruct_Once sync.Once

func dBusArgInfoStruct_Set() error {
	var err error
	dBusArgInfoStruct_Once.Do(func() {
		dBusArgInfoStruct, err = gi.StructNew("Gio", "DBusArgInfo")
	})
	return err
}

type DBusArgInfo struct {
	native uintptr
}

var dBusArgInfoRefFunction *gi.Function
var dBusArgInfoRefFunction_Once sync.Once

func dBusArgInfoRefFunction_Set() error {
	var err error
	dBusArgInfoRefFunction_Once.Do(func() {
		err = dBusArgInfoStruct_Set()
		if err != nil {
			return
		}
		dBusArgInfoRefFunction, err = dBusArgInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_arg_info_ref.
func (recv *DBusArgInfo) Ref() (*DBusArgInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusArgInfoRefFunction_Set()
	if err == nil {
		ret = dBusArgInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusArgInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusArgInfoUnrefFunction *gi.Function
var dBusArgInfoUnrefFunction_Once sync.Once

func dBusArgInfoUnrefFunction_Set() error {
	var err error
	dBusArgInfoUnrefFunction_Once.Do(func() {
		err = dBusArgInfoStruct_Set()
		if err != nil {
			return
		}
		dBusArgInfoUnrefFunction, err = dBusArgInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_arg_info_unref.
func (recv *DBusArgInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusArgInfoUnrefFunction_Set()
	if err == nil {
		dBusArgInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusErrorEntryStruct *gi.Struct
var dBusErrorEntryStruct_Once sync.Once

func dBusErrorEntryStruct_Set() error {
	var err error
	dBusErrorEntryStruct_Once.Do(func() {
		dBusErrorEntryStruct, err = gi.StructNew("Gio", "DBusErrorEntry")
	})
	return err
}

type DBusErrorEntry struct {
	native uintptr
}

var dBusInterfaceIfaceStruct *gi.Struct
var dBusInterfaceIfaceStruct_Once sync.Once

func dBusInterfaceIfaceStruct_Set() error {
	var err error
	dBusInterfaceIfaceStruct_Once.Do(func() {
		dBusInterfaceIfaceStruct, err = gi.StructNew("Gio", "DBusInterfaceIface")
	})
	return err
}

type DBusInterfaceIface struct {
	native uintptr
}

var dBusInterfaceInfoStruct *gi.Struct
var dBusInterfaceInfoStruct_Once sync.Once

func dBusInterfaceInfoStruct_Set() error {
	var err error
	dBusInterfaceInfoStruct_Once.Do(func() {
		dBusInterfaceInfoStruct, err = gi.StructNew("Gio", "DBusInterfaceInfo")
	})
	return err
}

type DBusInterfaceInfo struct {
	native uintptr
}

var dBusInterfaceInfoCacheBuildFunction *gi.Function
var dBusInterfaceInfoCacheBuildFunction_Once sync.Once

func dBusInterfaceInfoCacheBuildFunction_Set() error {
	var err error
	dBusInterfaceInfoCacheBuildFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoCacheBuildFunction, err = dBusInterfaceInfoStruct.InvokerNew("cache_build")
	})
	return err
}

// CacheBuild is a representation of the C type g_dbus_interface_info_cache_build.
func (recv *DBusInterfaceInfo) CacheBuild() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoCacheBuildFunction_Set()
	if err == nil {
		dBusInterfaceInfoCacheBuildFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusInterfaceInfoCacheReleaseFunction *gi.Function
var dBusInterfaceInfoCacheReleaseFunction_Once sync.Once

func dBusInterfaceInfoCacheReleaseFunction_Set() error {
	var err error
	dBusInterfaceInfoCacheReleaseFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoCacheReleaseFunction, err = dBusInterfaceInfoStruct.InvokerNew("cache_release")
	})
	return err
}

// CacheRelease is a representation of the C type g_dbus_interface_info_cache_release.
func (recv *DBusInterfaceInfo) CacheRelease() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoCacheReleaseFunction_Set()
	if err == nil {
		dBusInterfaceInfoCacheReleaseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_dbus_interface_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusInterfaceInfoLookupMethodFunction *gi.Function
var dBusInterfaceInfoLookupMethodFunction_Once sync.Once

func dBusInterfaceInfoLookupMethodFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupMethodFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupMethodFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_method")
	})
	return err
}

// LookupMethod is a representation of the C type g_dbus_interface_info_lookup_method.
func (recv *DBusInterfaceInfo) LookupMethod(name string) (*DBusMethodInfo, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupMethodFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupMethodFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusInterfaceInfoLookupPropertyFunction *gi.Function
var dBusInterfaceInfoLookupPropertyFunction_Once sync.Once

func dBusInterfaceInfoLookupPropertyFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupPropertyFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupPropertyFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_property")
	})
	return err
}

// LookupProperty is a representation of the C type g_dbus_interface_info_lookup_property.
func (recv *DBusInterfaceInfo) LookupProperty(name string) (*DBusPropertyInfo, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupPropertyFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupPropertyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusInterfaceInfoLookupSignalFunction *gi.Function
var dBusInterfaceInfoLookupSignalFunction_Once sync.Once

func dBusInterfaceInfoLookupSignalFunction_Set() error {
	var err error
	dBusInterfaceInfoLookupSignalFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoLookupSignalFunction, err = dBusInterfaceInfoStruct.InvokerNew("lookup_signal")
	})
	return err
}

// LookupSignal is a representation of the C type g_dbus_interface_info_lookup_signal.
func (recv *DBusInterfaceInfo) LookupSignal(name string) (*DBusSignalInfo, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusInterfaceInfoLookupSignalFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoLookupSignalFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusInterfaceInfoRefFunction *gi.Function
var dBusInterfaceInfoRefFunction_Once sync.Once

func dBusInterfaceInfoRefFunction_Set() error {
	var err error
	dBusInterfaceInfoRefFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoRefFunction, err = dBusInterfaceInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_interface_info_ref.
func (recv *DBusInterfaceInfo) Ref() (*DBusInterfaceInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusInterfaceInfoRefFunction_Set()
	if err == nil {
		ret = dBusInterfaceInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusInterfaceInfoUnrefFunction *gi.Function
var dBusInterfaceInfoUnrefFunction_Once sync.Once

func dBusInterfaceInfoUnrefFunction_Set() error {
	var err error
	dBusInterfaceInfoUnrefFunction_Once.Do(func() {
		err = dBusInterfaceInfoStruct_Set()
		if err != nil {
			return
		}
		dBusInterfaceInfoUnrefFunction, err = dBusInterfaceInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_interface_info_unref.
func (recv *DBusInterfaceInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusInterfaceInfoUnrefFunction_Set()
	if err == nil {
		dBusInterfaceInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusInterfaceSkeletonClassStruct *gi.Struct
var dBusInterfaceSkeletonClassStruct_Once sync.Once

func dBusInterfaceSkeletonClassStruct_Set() error {
	var err error
	dBusInterfaceSkeletonClassStruct_Once.Do(func() {
		dBusInterfaceSkeletonClassStruct, err = gi.StructNew("Gio", "DBusInterfaceSkeletonClass")
	})
	return err
}

type DBusInterfaceSkeletonClass struct {
	native uintptr
}

var dBusInterfaceSkeletonPrivateStruct *gi.Struct
var dBusInterfaceSkeletonPrivateStruct_Once sync.Once

func dBusInterfaceSkeletonPrivateStruct_Set() error {
	var err error
	dBusInterfaceSkeletonPrivateStruct_Once.Do(func() {
		dBusInterfaceSkeletonPrivateStruct, err = gi.StructNew("Gio", "DBusInterfaceSkeletonPrivate")
	})
	return err
}

type DBusInterfaceSkeletonPrivate struct {
	native uintptr
}

var dBusInterfaceVTableStruct *gi.Struct
var dBusInterfaceVTableStruct_Once sync.Once

func dBusInterfaceVTableStruct_Set() error {
	var err error
	dBusInterfaceVTableStruct_Once.Do(func() {
		dBusInterfaceVTableStruct, err = gi.StructNew("Gio", "DBusInterfaceVTable")
	})
	return err
}

type DBusInterfaceVTable struct {
	native uintptr
}

var dBusMethodInfoStruct *gi.Struct
var dBusMethodInfoStruct_Once sync.Once

func dBusMethodInfoStruct_Set() error {
	var err error
	dBusMethodInfoStruct_Once.Do(func() {
		dBusMethodInfoStruct, err = gi.StructNew("Gio", "DBusMethodInfo")
	})
	return err
}

type DBusMethodInfo struct {
	native uintptr
}

var dBusMethodInfoRefFunction *gi.Function
var dBusMethodInfoRefFunction_Once sync.Once

func dBusMethodInfoRefFunction_Set() error {
	var err error
	dBusMethodInfoRefFunction_Once.Do(func() {
		err = dBusMethodInfoStruct_Set()
		if err != nil {
			return
		}
		dBusMethodInfoRefFunction, err = dBusMethodInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_method_info_ref.
func (recv *DBusMethodInfo) Ref() (*DBusMethodInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusMethodInfoRefFunction_Set()
	if err == nil {
		ret = dBusMethodInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusMethodInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusMethodInfoUnrefFunction *gi.Function
var dBusMethodInfoUnrefFunction_Once sync.Once

func dBusMethodInfoUnrefFunction_Set() error {
	var err error
	dBusMethodInfoUnrefFunction_Once.Do(func() {
		err = dBusMethodInfoStruct_Set()
		if err != nil {
			return
		}
		dBusMethodInfoUnrefFunction, err = dBusMethodInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_method_info_unref.
func (recv *DBusMethodInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusMethodInfoUnrefFunction_Set()
	if err == nil {
		dBusMethodInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusNodeInfoStruct *gi.Struct
var dBusNodeInfoStruct_Once sync.Once

func dBusNodeInfoStruct_Set() error {
	var err error
	dBusNodeInfoStruct_Once.Do(func() {
		dBusNodeInfoStruct, err = gi.StructNew("Gio", "DBusNodeInfo")
	})
	return err
}

type DBusNodeInfo struct {
	native uintptr
}

var dBusNodeInfoNewForXmlFunction *gi.Function
var dBusNodeInfoNewForXmlFunction_Once sync.Once

func dBusNodeInfoNewForXmlFunction_Set() error {
	var err error
	dBusNodeInfoNewForXmlFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoNewForXmlFunction, err = dBusNodeInfoStruct.InvokerNew("new_for_xml")
	})
	return err
}

// DBusNodeInfoNewForXml is a representation of the C type g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) (*DBusNodeInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(xmlData)

	var ret gi.Argument

	err := dBusNodeInfoNewForXmlFunction_Set()
	if err == nil {
		ret = dBusNodeInfoNewForXmlFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_dbus_node_info_generate_xml' : parameter 'string_builder' of type 'GLib.String' not supported

var dBusNodeInfoLookupInterfaceFunction *gi.Function
var dBusNodeInfoLookupInterfaceFunction_Once sync.Once

func dBusNodeInfoLookupInterfaceFunction_Set() error {
	var err error
	dBusNodeInfoLookupInterfaceFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoLookupInterfaceFunction, err = dBusNodeInfoStruct.InvokerNew("lookup_interface")
	})
	return err
}

// LookupInterface is a representation of the C type g_dbus_node_info_lookup_interface.
func (recv *DBusNodeInfo) LookupInterface(name string) (*DBusInterfaceInfo, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := dBusNodeInfoLookupInterfaceFunction_Set()
	if err == nil {
		ret = dBusNodeInfoLookupInterfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusInterfaceInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusNodeInfoRefFunction *gi.Function
var dBusNodeInfoRefFunction_Once sync.Once

func dBusNodeInfoRefFunction_Set() error {
	var err error
	dBusNodeInfoRefFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoRefFunction, err = dBusNodeInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_node_info_ref.
func (recv *DBusNodeInfo) Ref() (*DBusNodeInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusNodeInfoRefFunction_Set()
	if err == nil {
		ret = dBusNodeInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusNodeInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusNodeInfoUnrefFunction *gi.Function
var dBusNodeInfoUnrefFunction_Once sync.Once

func dBusNodeInfoUnrefFunction_Set() error {
	var err error
	dBusNodeInfoUnrefFunction_Once.Do(func() {
		err = dBusNodeInfoStruct_Set()
		if err != nil {
			return
		}
		dBusNodeInfoUnrefFunction, err = dBusNodeInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_node_info_unref.
func (recv *DBusNodeInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusNodeInfoUnrefFunction_Set()
	if err == nil {
		dBusNodeInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusObjectIfaceStruct *gi.Struct
var dBusObjectIfaceStruct_Once sync.Once

func dBusObjectIfaceStruct_Set() error {
	var err error
	dBusObjectIfaceStruct_Once.Do(func() {
		dBusObjectIfaceStruct, err = gi.StructNew("Gio", "DBusObjectIface")
	})
	return err
}

type DBusObjectIface struct {
	native uintptr
}

var dBusObjectManagerClientClassStruct *gi.Struct
var dBusObjectManagerClientClassStruct_Once sync.Once

func dBusObjectManagerClientClassStruct_Set() error {
	var err error
	dBusObjectManagerClientClassStruct_Once.Do(func() {
		dBusObjectManagerClientClassStruct, err = gi.StructNew("Gio", "DBusObjectManagerClientClass")
	})
	return err
}

type DBusObjectManagerClientClass struct {
	native uintptr
}

var dBusObjectManagerClientPrivateStruct *gi.Struct
var dBusObjectManagerClientPrivateStruct_Once sync.Once

func dBusObjectManagerClientPrivateStruct_Set() error {
	var err error
	dBusObjectManagerClientPrivateStruct_Once.Do(func() {
		dBusObjectManagerClientPrivateStruct, err = gi.StructNew("Gio", "DBusObjectManagerClientPrivate")
	})
	return err
}

type DBusObjectManagerClientPrivate struct {
	native uintptr
}

var dBusObjectManagerIfaceStruct *gi.Struct
var dBusObjectManagerIfaceStruct_Once sync.Once

func dBusObjectManagerIfaceStruct_Set() error {
	var err error
	dBusObjectManagerIfaceStruct_Once.Do(func() {
		dBusObjectManagerIfaceStruct, err = gi.StructNew("Gio", "DBusObjectManagerIface")
	})
	return err
}

type DBusObjectManagerIface struct {
	native uintptr
}

var dBusObjectManagerServerClassStruct *gi.Struct
var dBusObjectManagerServerClassStruct_Once sync.Once

func dBusObjectManagerServerClassStruct_Set() error {
	var err error
	dBusObjectManagerServerClassStruct_Once.Do(func() {
		dBusObjectManagerServerClassStruct, err = gi.StructNew("Gio", "DBusObjectManagerServerClass")
	})
	return err
}

type DBusObjectManagerServerClass struct {
	native uintptr
}

var dBusObjectManagerServerPrivateStruct *gi.Struct
var dBusObjectManagerServerPrivateStruct_Once sync.Once

func dBusObjectManagerServerPrivateStruct_Set() error {
	var err error
	dBusObjectManagerServerPrivateStruct_Once.Do(func() {
		dBusObjectManagerServerPrivateStruct, err = gi.StructNew("Gio", "DBusObjectManagerServerPrivate")
	})
	return err
}

type DBusObjectManagerServerPrivate struct {
	native uintptr
}

var dBusObjectProxyClassStruct *gi.Struct
var dBusObjectProxyClassStruct_Once sync.Once

func dBusObjectProxyClassStruct_Set() error {
	var err error
	dBusObjectProxyClassStruct_Once.Do(func() {
		dBusObjectProxyClassStruct, err = gi.StructNew("Gio", "DBusObjectProxyClass")
	})
	return err
}

type DBusObjectProxyClass struct {
	native uintptr
}

var dBusObjectProxyPrivateStruct *gi.Struct
var dBusObjectProxyPrivateStruct_Once sync.Once

func dBusObjectProxyPrivateStruct_Set() error {
	var err error
	dBusObjectProxyPrivateStruct_Once.Do(func() {
		dBusObjectProxyPrivateStruct, err = gi.StructNew("Gio", "DBusObjectProxyPrivate")
	})
	return err
}

type DBusObjectProxyPrivate struct {
	native uintptr
}

var dBusObjectSkeletonClassStruct *gi.Struct
var dBusObjectSkeletonClassStruct_Once sync.Once

func dBusObjectSkeletonClassStruct_Set() error {
	var err error
	dBusObjectSkeletonClassStruct_Once.Do(func() {
		dBusObjectSkeletonClassStruct, err = gi.StructNew("Gio", "DBusObjectSkeletonClass")
	})
	return err
}

type DBusObjectSkeletonClass struct {
	native uintptr
}

var dBusObjectSkeletonPrivateStruct *gi.Struct
var dBusObjectSkeletonPrivateStruct_Once sync.Once

func dBusObjectSkeletonPrivateStruct_Set() error {
	var err error
	dBusObjectSkeletonPrivateStruct_Once.Do(func() {
		dBusObjectSkeletonPrivateStruct, err = gi.StructNew("Gio", "DBusObjectSkeletonPrivate")
	})
	return err
}

type DBusObjectSkeletonPrivate struct {
	native uintptr
}

var dBusPropertyInfoStruct *gi.Struct
var dBusPropertyInfoStruct_Once sync.Once

func dBusPropertyInfoStruct_Set() error {
	var err error
	dBusPropertyInfoStruct_Once.Do(func() {
		dBusPropertyInfoStruct, err = gi.StructNew("Gio", "DBusPropertyInfo")
	})
	return err
}

type DBusPropertyInfo struct {
	native uintptr
}

var dBusPropertyInfoRefFunction *gi.Function
var dBusPropertyInfoRefFunction_Once sync.Once

func dBusPropertyInfoRefFunction_Set() error {
	var err error
	dBusPropertyInfoRefFunction_Once.Do(func() {
		err = dBusPropertyInfoStruct_Set()
		if err != nil {
			return
		}
		dBusPropertyInfoRefFunction, err = dBusPropertyInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_property_info_ref.
func (recv *DBusPropertyInfo) Ref() (*DBusPropertyInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusPropertyInfoRefFunction_Set()
	if err == nil {
		ret = dBusPropertyInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusPropertyInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusPropertyInfoUnrefFunction *gi.Function
var dBusPropertyInfoUnrefFunction_Once sync.Once

func dBusPropertyInfoUnrefFunction_Set() error {
	var err error
	dBusPropertyInfoUnrefFunction_Once.Do(func() {
		err = dBusPropertyInfoStruct_Set()
		if err != nil {
			return
		}
		dBusPropertyInfoUnrefFunction, err = dBusPropertyInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_property_info_unref.
func (recv *DBusPropertyInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusPropertyInfoUnrefFunction_Set()
	if err == nil {
		dBusPropertyInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusProxyClassStruct *gi.Struct
var dBusProxyClassStruct_Once sync.Once

func dBusProxyClassStruct_Set() error {
	var err error
	dBusProxyClassStruct_Once.Do(func() {
		dBusProxyClassStruct, err = gi.StructNew("Gio", "DBusProxyClass")
	})
	return err
}

type DBusProxyClass struct {
	native uintptr
}

var dBusProxyPrivateStruct *gi.Struct
var dBusProxyPrivateStruct_Once sync.Once

func dBusProxyPrivateStruct_Set() error {
	var err error
	dBusProxyPrivateStruct_Once.Do(func() {
		dBusProxyPrivateStruct, err = gi.StructNew("Gio", "DBusProxyPrivate")
	})
	return err
}

type DBusProxyPrivate struct {
	native uintptr
}

var dBusSignalInfoStruct *gi.Struct
var dBusSignalInfoStruct_Once sync.Once

func dBusSignalInfoStruct_Set() error {
	var err error
	dBusSignalInfoStruct_Once.Do(func() {
		dBusSignalInfoStruct, err = gi.StructNew("Gio", "DBusSignalInfo")
	})
	return err
}

type DBusSignalInfo struct {
	native uintptr
}

var dBusSignalInfoRefFunction *gi.Function
var dBusSignalInfoRefFunction_Once sync.Once

func dBusSignalInfoRefFunction_Set() error {
	var err error
	dBusSignalInfoRefFunction_Once.Do(func() {
		err = dBusSignalInfoStruct_Set()
		if err != nil {
			return
		}
		dBusSignalInfoRefFunction, err = dBusSignalInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_dbus_signal_info_ref.
func (recv *DBusSignalInfo) Ref() (*DBusSignalInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dBusSignalInfoRefFunction_Set()
	if err == nil {
		ret = dBusSignalInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DBusSignalInfo{native: ret.Pointer()}

	return retGo, err
}

var dBusSignalInfoUnrefFunction *gi.Function
var dBusSignalInfoUnrefFunction_Once sync.Once

func dBusSignalInfoUnrefFunction_Set() error {
	var err error
	dBusSignalInfoUnrefFunction_Once.Do(func() {
		err = dBusSignalInfoStruct_Set()
		if err != nil {
			return
		}
		dBusSignalInfoUnrefFunction, err = dBusSignalInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_dbus_signal_info_unref.
func (recv *DBusSignalInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dBusSignalInfoUnrefFunction_Set()
	if err == nil {
		dBusSignalInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dBusSubtreeVTableStruct *gi.Struct
var dBusSubtreeVTableStruct_Once sync.Once

func dBusSubtreeVTableStruct_Set() error {
	var err error
	dBusSubtreeVTableStruct_Once.Do(func() {
		dBusSubtreeVTableStruct, err = gi.StructNew("Gio", "DBusSubtreeVTable")
	})
	return err
}

type DBusSubtreeVTable struct {
	native uintptr
}

var dataInputStreamClassStruct *gi.Struct
var dataInputStreamClassStruct_Once sync.Once

func dataInputStreamClassStruct_Set() error {
	var err error
	dataInputStreamClassStruct_Once.Do(func() {
		dataInputStreamClassStruct, err = gi.StructNew("Gio", "DataInputStreamClass")
	})
	return err
}

type DataInputStreamClass struct {
	native uintptr
}

var dataInputStreamPrivateStruct *gi.Struct
var dataInputStreamPrivateStruct_Once sync.Once

func dataInputStreamPrivateStruct_Set() error {
	var err error
	dataInputStreamPrivateStruct_Once.Do(func() {
		dataInputStreamPrivateStruct, err = gi.StructNew("Gio", "DataInputStreamPrivate")
	})
	return err
}

type DataInputStreamPrivate struct {
	native uintptr
}

var dataOutputStreamClassStruct *gi.Struct
var dataOutputStreamClassStruct_Once sync.Once

func dataOutputStreamClassStruct_Set() error {
	var err error
	dataOutputStreamClassStruct_Once.Do(func() {
		dataOutputStreamClassStruct, err = gi.StructNew("Gio", "DataOutputStreamClass")
	})
	return err
}

type DataOutputStreamClass struct {
	native uintptr
}

var dataOutputStreamPrivateStruct *gi.Struct
var dataOutputStreamPrivateStruct_Once sync.Once

func dataOutputStreamPrivateStruct_Set() error {
	var err error
	dataOutputStreamPrivateStruct_Once.Do(func() {
		dataOutputStreamPrivateStruct, err = gi.StructNew("Gio", "DataOutputStreamPrivate")
	})
	return err
}

type DataOutputStreamPrivate struct {
	native uintptr
}

var datagramBasedInterfaceStruct *gi.Struct
var datagramBasedInterfaceStruct_Once sync.Once

func datagramBasedInterfaceStruct_Set() error {
	var err error
	datagramBasedInterfaceStruct_Once.Do(func() {
		datagramBasedInterfaceStruct, err = gi.StructNew("Gio", "DatagramBasedInterface")
	})
	return err
}

type DatagramBasedInterface struct {
	native uintptr
}

var desktopAppInfoClassStruct *gi.Struct
var desktopAppInfoClassStruct_Once sync.Once

func desktopAppInfoClassStruct_Set() error {
	var err error
	desktopAppInfoClassStruct_Once.Do(func() {
		desktopAppInfoClassStruct, err = gi.StructNew("Gio", "DesktopAppInfoClass")
	})
	return err
}

type DesktopAppInfoClass struct {
	native uintptr
}

var desktopAppInfoLookupIfaceStruct *gi.Struct
var desktopAppInfoLookupIfaceStruct_Once sync.Once

func desktopAppInfoLookupIfaceStruct_Set() error {
	var err error
	desktopAppInfoLookupIfaceStruct_Once.Do(func() {
		desktopAppInfoLookupIfaceStruct, err = gi.StructNew("Gio", "DesktopAppInfoLookupIface")
	})
	return err
}

type DesktopAppInfoLookupIface struct {
	native uintptr
}

var driveIfaceStruct *gi.Struct
var driveIfaceStruct_Once sync.Once

func driveIfaceStruct_Set() error {
	var err error
	driveIfaceStruct_Once.Do(func() {
		driveIfaceStruct, err = gi.StructNew("Gio", "DriveIface")
	})
	return err
}

type DriveIface struct {
	native uintptr
}

var dtlsClientConnectionInterfaceStruct *gi.Struct
var dtlsClientConnectionInterfaceStruct_Once sync.Once

func dtlsClientConnectionInterfaceStruct_Set() error {
	var err error
	dtlsClientConnectionInterfaceStruct_Once.Do(func() {
		dtlsClientConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsClientConnectionInterface")
	})
	return err
}

type DtlsClientConnectionInterface struct {
	native uintptr
}

var dtlsConnectionInterfaceStruct *gi.Struct
var dtlsConnectionInterfaceStruct_Once sync.Once

func dtlsConnectionInterfaceStruct_Set() error {
	var err error
	dtlsConnectionInterfaceStruct_Once.Do(func() {
		dtlsConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsConnectionInterface")
	})
	return err
}

type DtlsConnectionInterface struct {
	native uintptr
}

var dtlsServerConnectionInterfaceStruct *gi.Struct
var dtlsServerConnectionInterfaceStruct_Once sync.Once

func dtlsServerConnectionInterfaceStruct_Set() error {
	var err error
	dtlsServerConnectionInterfaceStruct_Once.Do(func() {
		dtlsServerConnectionInterfaceStruct, err = gi.StructNew("Gio", "DtlsServerConnectionInterface")
	})
	return err
}

type DtlsServerConnectionInterface struct {
	native uintptr
}

var emblemClassStruct *gi.Struct
var emblemClassStruct_Once sync.Once

func emblemClassStruct_Set() error {
	var err error
	emblemClassStruct_Once.Do(func() {
		emblemClassStruct, err = gi.StructNew("Gio", "EmblemClass")
	})
	return err
}

type EmblemClass struct {
	native uintptr
}

var emblemedIconClassStruct *gi.Struct
var emblemedIconClassStruct_Once sync.Once

func emblemedIconClassStruct_Set() error {
	var err error
	emblemedIconClassStruct_Once.Do(func() {
		emblemedIconClassStruct, err = gi.StructNew("Gio", "EmblemedIconClass")
	})
	return err
}

type EmblemedIconClass struct {
	native uintptr
}

var emblemedIconPrivateStruct *gi.Struct
var emblemedIconPrivateStruct_Once sync.Once

func emblemedIconPrivateStruct_Set() error {
	var err error
	emblemedIconPrivateStruct_Once.Do(func() {
		emblemedIconPrivateStruct, err = gi.StructNew("Gio", "EmblemedIconPrivate")
	})
	return err
}

type EmblemedIconPrivate struct {
	native uintptr
}

var fileAttributeInfoStruct *gi.Struct
var fileAttributeInfoStruct_Once sync.Once

func fileAttributeInfoStruct_Set() error {
	var err error
	fileAttributeInfoStruct_Once.Do(func() {
		fileAttributeInfoStruct, err = gi.StructNew("Gio", "FileAttributeInfo")
	})
	return err
}

type FileAttributeInfo struct {
	native uintptr
}

var fileAttributeInfoListStruct *gi.Struct
var fileAttributeInfoListStruct_Once sync.Once

func fileAttributeInfoListStruct_Set() error {
	var err error
	fileAttributeInfoListStruct_Once.Do(func() {
		fileAttributeInfoListStruct, err = gi.StructNew("Gio", "FileAttributeInfoList")
	})
	return err
}

type FileAttributeInfoList struct {
	native uintptr
}

var fileAttributeInfoListNewFunction *gi.Function
var fileAttributeInfoListNewFunction_Once sync.Once

func fileAttributeInfoListNewFunction_Set() error {
	var err error
	fileAttributeInfoListNewFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListNewFunction, err = fileAttributeInfoListStruct.InvokerNew("new")
	})
	return err
}

// FileAttributeInfoListNew is a representation of the C type g_file_attribute_info_list_new.
func FileAttributeInfoListNew() (*FileAttributeInfoList, error) {

	var ret gi.Argument

	err := fileAttributeInfoListNewFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListNewFunction.Invoke(nil, nil)
	}

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_file_attribute_info_list_add' : parameter 'type' of type 'FileAttributeType' not supported

var fileAttributeInfoListDupFunction *gi.Function
var fileAttributeInfoListDupFunction_Once sync.Once

func fileAttributeInfoListDupFunction_Set() error {
	var err error
	fileAttributeInfoListDupFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListDupFunction, err = fileAttributeInfoListStruct.InvokerNew("dup")
	})
	return err
}

// Dup is a representation of the C type g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() (*FileAttributeInfoList, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeInfoListDupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListDupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeInfoListLookupFunction *gi.Function
var fileAttributeInfoListLookupFunction_Once sync.Once

func fileAttributeInfoListLookupFunction_Set() error {
	var err error
	fileAttributeInfoListLookupFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListLookupFunction, err = fileAttributeInfoListStruct.InvokerNew("lookup")
	})
	return err
}

// Lookup is a representation of the C type g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) (*FileAttributeInfo, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := fileAttributeInfoListLookupFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfo{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeInfoListRefFunction *gi.Function
var fileAttributeInfoListRefFunction_Once sync.Once

func fileAttributeInfoListRefFunction_Set() error {
	var err error
	fileAttributeInfoListRefFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListRefFunction, err = fileAttributeInfoListStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() (*FileAttributeInfoList, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeInfoListRefFunction_Set()
	if err == nil {
		ret = fileAttributeInfoListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeInfoList{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeInfoListUnrefFunction *gi.Function
var fileAttributeInfoListUnrefFunction_Once sync.Once

func fileAttributeInfoListUnrefFunction_Set() error {
	var err error
	fileAttributeInfoListUnrefFunction_Once.Do(func() {
		err = fileAttributeInfoListStruct_Set()
		if err != nil {
			return
		}
		fileAttributeInfoListUnrefFunction, err = fileAttributeInfoListStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_file_attribute_info_list_unref.
func (recv *FileAttributeInfoList) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileAttributeInfoListUnrefFunction_Set()
	if err == nil {
		fileAttributeInfoListUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var fileAttributeMatcherStruct *gi.Struct
var fileAttributeMatcherStruct_Once sync.Once

func fileAttributeMatcherStruct_Set() error {
	var err error
	fileAttributeMatcherStruct_Once.Do(func() {
		fileAttributeMatcherStruct, err = gi.StructNew("Gio", "FileAttributeMatcher")
	})
	return err
}

type FileAttributeMatcher struct {
	native uintptr
}

var fileAttributeMatcherNewFunction *gi.Function
var fileAttributeMatcherNewFunction_Once sync.Once

func fileAttributeMatcherNewFunction_Set() error {
	var err error
	fileAttributeMatcherNewFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherNewFunction, err = fileAttributeMatcherStruct.InvokerNew("new")
	})
	return err
}

// FileAttributeMatcherNew is a representation of the C type g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) (*FileAttributeMatcher, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(attributes)

	var ret gi.Argument

	err := fileAttributeMatcherNewFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeMatcherEnumerateNamespaceFunction *gi.Function
var fileAttributeMatcherEnumerateNamespaceFunction_Once sync.Once

func fileAttributeMatcherEnumerateNamespaceFunction_Set() error {
	var err error
	fileAttributeMatcherEnumerateNamespaceFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherEnumerateNamespaceFunction, err = fileAttributeMatcherStruct.InvokerNew("enumerate_namespace")
	})
	return err
}

// EnumerateNamespace is a representation of the C type g_file_attribute_matcher_enumerate_namespace.
func (recv *FileAttributeMatcher) EnumerateNamespace(ns string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(ns)

	var ret gi.Argument

	err := fileAttributeMatcherEnumerateNamespaceFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherEnumerateNamespaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var fileAttributeMatcherEnumerateNextFunction *gi.Function
var fileAttributeMatcherEnumerateNextFunction_Once sync.Once

func fileAttributeMatcherEnumerateNextFunction_Set() error {
	var err error
	fileAttributeMatcherEnumerateNextFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherEnumerateNextFunction, err = fileAttributeMatcherStruct.InvokerNew("enumerate_next")
	})
	return err
}

// EnumerateNext is a representation of the C type g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherEnumerateNextFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherEnumerateNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var fileAttributeMatcherMatchesFunction *gi.Function
var fileAttributeMatcherMatchesFunction_Once sync.Once

func fileAttributeMatcherMatchesFunction_Set() error {
	var err error
	fileAttributeMatcherMatchesFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherMatchesFunction, err = fileAttributeMatcherStruct.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type g_file_attribute_matcher_matches.
func (recv *FileAttributeMatcher) Matches(attribute string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileAttributeMatcherMatchesFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var fileAttributeMatcherMatchesOnlyFunction *gi.Function
var fileAttributeMatcherMatchesOnlyFunction_Once sync.Once

func fileAttributeMatcherMatchesOnlyFunction_Set() error {
	var err error
	fileAttributeMatcherMatchesOnlyFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherMatchesOnlyFunction, err = fileAttributeMatcherStruct.InvokerNew("matches_only")
	})
	return err
}

// MatchesOnly is a representation of the C type g_file_attribute_matcher_matches_only.
func (recv *FileAttributeMatcher) MatchesOnly(attribute string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(attribute)

	var ret gi.Argument

	err := fileAttributeMatcherMatchesOnlyFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherMatchesOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var fileAttributeMatcherRefFunction *gi.Function
var fileAttributeMatcherRefFunction_Once sync.Once

func fileAttributeMatcherRefFunction_Set() error {
	var err error
	fileAttributeMatcherRefFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherRefFunction, err = fileAttributeMatcherStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() (*FileAttributeMatcher, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherRefFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeMatcherSubtractFunction *gi.Function
var fileAttributeMatcherSubtractFunction_Once sync.Once

func fileAttributeMatcherSubtractFunction_Set() error {
	var err error
	fileAttributeMatcherSubtractFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherSubtractFunction, err = fileAttributeMatcherStruct.InvokerNew("subtract")
	})
	return err
}

// Subtract is a representation of the C type g_file_attribute_matcher_subtract.
func (recv *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) (*FileAttributeMatcher, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(subtract.native)

	var ret gi.Argument

	err := fileAttributeMatcherSubtractFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherSubtractFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileAttributeMatcher{native: ret.Pointer()}

	return retGo, err
}

var fileAttributeMatcherToStringFunction *gi.Function
var fileAttributeMatcherToStringFunction_Once sync.Once

func fileAttributeMatcherToStringFunction_Set() error {
	var err error
	fileAttributeMatcherToStringFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherToStringFunction, err = fileAttributeMatcherStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type g_file_attribute_matcher_to_string.
func (recv *FileAttributeMatcher) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fileAttributeMatcherToStringFunction_Set()
	if err == nil {
		ret = fileAttributeMatcherToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var fileAttributeMatcherUnrefFunction *gi.Function
var fileAttributeMatcherUnrefFunction_Once sync.Once

func fileAttributeMatcherUnrefFunction_Set() error {
	var err error
	fileAttributeMatcherUnrefFunction_Once.Do(func() {
		err = fileAttributeMatcherStruct_Set()
		if err != nil {
			return
		}
		fileAttributeMatcherUnrefFunction, err = fileAttributeMatcherStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_file_attribute_matcher_unref.
func (recv *FileAttributeMatcher) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fileAttributeMatcherUnrefFunction_Set()
	if err == nil {
		fileAttributeMatcherUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var fileDescriptorBasedIfaceStruct *gi.Struct
var fileDescriptorBasedIfaceStruct_Once sync.Once

func fileDescriptorBasedIfaceStruct_Set() error {
	var err error
	fileDescriptorBasedIfaceStruct_Once.Do(func() {
		fileDescriptorBasedIfaceStruct, err = gi.StructNew("Gio", "FileDescriptorBasedIface")
	})
	return err
}

type FileDescriptorBasedIface struct {
	native uintptr
}

var fileEnumeratorClassStruct *gi.Struct
var fileEnumeratorClassStruct_Once sync.Once

func fileEnumeratorClassStruct_Set() error {
	var err error
	fileEnumeratorClassStruct_Once.Do(func() {
		fileEnumeratorClassStruct, err = gi.StructNew("Gio", "FileEnumeratorClass")
	})
	return err
}

type FileEnumeratorClass struct {
	native uintptr
}

var fileEnumeratorPrivateStruct *gi.Struct
var fileEnumeratorPrivateStruct_Once sync.Once

func fileEnumeratorPrivateStruct_Set() error {
	var err error
	fileEnumeratorPrivateStruct_Once.Do(func() {
		fileEnumeratorPrivateStruct, err = gi.StructNew("Gio", "FileEnumeratorPrivate")
	})
	return err
}

type FileEnumeratorPrivate struct {
	native uintptr
}

var fileIOStreamClassStruct *gi.Struct
var fileIOStreamClassStruct_Once sync.Once

func fileIOStreamClassStruct_Set() error {
	var err error
	fileIOStreamClassStruct_Once.Do(func() {
		fileIOStreamClassStruct, err = gi.StructNew("Gio", "FileIOStreamClass")
	})
	return err
}

type FileIOStreamClass struct {
	native uintptr
}

var fileIOStreamPrivateStruct *gi.Struct
var fileIOStreamPrivateStruct_Once sync.Once

func fileIOStreamPrivateStruct_Set() error {
	var err error
	fileIOStreamPrivateStruct_Once.Do(func() {
		fileIOStreamPrivateStruct, err = gi.StructNew("Gio", "FileIOStreamPrivate")
	})
	return err
}

type FileIOStreamPrivate struct {
	native uintptr
}

var fileIconClassStruct *gi.Struct
var fileIconClassStruct_Once sync.Once

func fileIconClassStruct_Set() error {
	var err error
	fileIconClassStruct_Once.Do(func() {
		fileIconClassStruct, err = gi.StructNew("Gio", "FileIconClass")
	})
	return err
}

type FileIconClass struct {
	native uintptr
}

var fileIfaceStruct *gi.Struct
var fileIfaceStruct_Once sync.Once

func fileIfaceStruct_Set() error {
	var err error
	fileIfaceStruct_Once.Do(func() {
		fileIfaceStruct, err = gi.StructNew("Gio", "FileIface")
	})
	return err
}

type FileIface struct {
	native uintptr
}

var fileInfoClassStruct *gi.Struct
var fileInfoClassStruct_Once sync.Once

func fileInfoClassStruct_Set() error {
	var err error
	fileInfoClassStruct_Once.Do(func() {
		fileInfoClassStruct, err = gi.StructNew("Gio", "FileInfoClass")
	})
	return err
}

type FileInfoClass struct {
	native uintptr
}

var fileInputStreamClassStruct *gi.Struct
var fileInputStreamClassStruct_Once sync.Once

func fileInputStreamClassStruct_Set() error {
	var err error
	fileInputStreamClassStruct_Once.Do(func() {
		fileInputStreamClassStruct, err = gi.StructNew("Gio", "FileInputStreamClass")
	})
	return err
}

type FileInputStreamClass struct {
	native uintptr
}

var fileInputStreamPrivateStruct *gi.Struct
var fileInputStreamPrivateStruct_Once sync.Once

func fileInputStreamPrivateStruct_Set() error {
	var err error
	fileInputStreamPrivateStruct_Once.Do(func() {
		fileInputStreamPrivateStruct, err = gi.StructNew("Gio", "FileInputStreamPrivate")
	})
	return err
}

type FileInputStreamPrivate struct {
	native uintptr
}

var fileMonitorClassStruct *gi.Struct
var fileMonitorClassStruct_Once sync.Once

func fileMonitorClassStruct_Set() error {
	var err error
	fileMonitorClassStruct_Once.Do(func() {
		fileMonitorClassStruct, err = gi.StructNew("Gio", "FileMonitorClass")
	})
	return err
}

type FileMonitorClass struct {
	native uintptr
}

var fileMonitorPrivateStruct *gi.Struct
var fileMonitorPrivateStruct_Once sync.Once

func fileMonitorPrivateStruct_Set() error {
	var err error
	fileMonitorPrivateStruct_Once.Do(func() {
		fileMonitorPrivateStruct, err = gi.StructNew("Gio", "FileMonitorPrivate")
	})
	return err
}

type FileMonitorPrivate struct {
	native uintptr
}

var fileOutputStreamClassStruct *gi.Struct
var fileOutputStreamClassStruct_Once sync.Once

func fileOutputStreamClassStruct_Set() error {
	var err error
	fileOutputStreamClassStruct_Once.Do(func() {
		fileOutputStreamClassStruct, err = gi.StructNew("Gio", "FileOutputStreamClass")
	})
	return err
}

type FileOutputStreamClass struct {
	native uintptr
}

var fileOutputStreamPrivateStruct *gi.Struct
var fileOutputStreamPrivateStruct_Once sync.Once

func fileOutputStreamPrivateStruct_Set() error {
	var err error
	fileOutputStreamPrivateStruct_Once.Do(func() {
		fileOutputStreamPrivateStruct, err = gi.StructNew("Gio", "FileOutputStreamPrivate")
	})
	return err
}

type FileOutputStreamPrivate struct {
	native uintptr
}

var filenameCompleterClassStruct *gi.Struct
var filenameCompleterClassStruct_Once sync.Once

func filenameCompleterClassStruct_Set() error {
	var err error
	filenameCompleterClassStruct_Once.Do(func() {
		filenameCompleterClassStruct, err = gi.StructNew("Gio", "FilenameCompleterClass")
	})
	return err
}

type FilenameCompleterClass struct {
	native uintptr
}

var filterInputStreamClassStruct *gi.Struct
var filterInputStreamClassStruct_Once sync.Once

func filterInputStreamClassStruct_Set() error {
	var err error
	filterInputStreamClassStruct_Once.Do(func() {
		filterInputStreamClassStruct, err = gi.StructNew("Gio", "FilterInputStreamClass")
	})
	return err
}

type FilterInputStreamClass struct {
	native uintptr
}

var filterOutputStreamClassStruct *gi.Struct
var filterOutputStreamClassStruct_Once sync.Once

func filterOutputStreamClassStruct_Set() error {
	var err error
	filterOutputStreamClassStruct_Once.Do(func() {
		filterOutputStreamClassStruct, err = gi.StructNew("Gio", "FilterOutputStreamClass")
	})
	return err
}

type FilterOutputStreamClass struct {
	native uintptr
}

var iOExtensionStruct *gi.Struct
var iOExtensionStruct_Once sync.Once

func iOExtensionStruct_Set() error {
	var err error
	iOExtensionStruct_Once.Do(func() {
		iOExtensionStruct, err = gi.StructNew("Gio", "IOExtension")
	})
	return err
}

type IOExtension struct {
	native uintptr
}

var iOExtensionGetNameFunction *gi.Function
var iOExtensionGetNameFunction_Once sync.Once

func iOExtensionGetNameFunction_Set() error {
	var err error
	iOExtensionGetNameFunction_Once.Do(func() {
		err = iOExtensionStruct_Set()
		if err != nil {
			return
		}
		iOExtensionGetNameFunction, err = iOExtensionStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_io_extension_get_name.
func (recv *IOExtension) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOExtensionGetNameFunction_Set()
	if err == nil {
		ret = iOExtensionGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var iOExtensionGetPriorityFunction *gi.Function
var iOExtensionGetPriorityFunction_Once sync.Once

func iOExtensionGetPriorityFunction_Set() error {
	var err error
	iOExtensionGetPriorityFunction_Once.Do(func() {
		err = iOExtensionStruct_Set()
		if err != nil {
			return
		}
		iOExtensionGetPriorityFunction, err = iOExtensionStruct.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOExtensionGetPriorityFunction_Set()
	if err == nil {
		ret = iOExtensionGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_extension_get_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_ref_class' : return type 'GObject.TypeClass' not supported

var iOExtensionPointStruct *gi.Struct
var iOExtensionPointStruct_Once sync.Once

func iOExtensionPointStruct_Set() error {
	var err error
	iOExtensionPointStruct_Once.Do(func() {
		iOExtensionPointStruct, err = gi.StructNew("Gio", "IOExtensionPoint")
	})
	return err
}

type IOExtensionPoint struct {
	native uintptr
}

var iOExtensionPointGetExtensionByNameFunction *gi.Function
var iOExtensionPointGetExtensionByNameFunction_Once sync.Once

func iOExtensionPointGetExtensionByNameFunction_Set() error {
	var err error
	iOExtensionPointGetExtensionByNameFunction_Once.Do(func() {
		err = iOExtensionPointStruct_Set()
		if err != nil {
			return
		}
		iOExtensionPointGetExtensionByNameFunction, err = iOExtensionPointStruct.InvokerNew("get_extension_by_name")
	})
	return err
}

// GetExtensionByName is a representation of the C type g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) (*IOExtension, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := iOExtensionPointGetExtensionByNameFunction_Set()
	if err == nil {
		ret = iOExtensionPointGetExtensionByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOExtension{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_extension_point_get_extensions' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_io_extension_point_get_required_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_io_extension_point_set_required_type' : parameter 'type' of type 'GType' not supported

var iOModuleClassStruct *gi.Struct
var iOModuleClassStruct_Once sync.Once

func iOModuleClassStruct_Set() error {
	var err error
	iOModuleClassStruct_Once.Do(func() {
		iOModuleClassStruct, err = gi.StructNew("Gio", "IOModuleClass")
	})
	return err
}

type IOModuleClass struct {
	native uintptr
}

var iOModuleScopeStruct *gi.Struct
var iOModuleScopeStruct_Once sync.Once

func iOModuleScopeStruct_Set() error {
	var err error
	iOModuleScopeStruct_Once.Do(func() {
		iOModuleScopeStruct, err = gi.StructNew("Gio", "IOModuleScope")
	})
	return err
}

type IOModuleScope struct {
	native uintptr
}

var iOModuleScopeBlockFunction *gi.Function
var iOModuleScopeBlockFunction_Once sync.Once

func iOModuleScopeBlockFunction_Set() error {
	var err error
	iOModuleScopeBlockFunction_Once.Do(func() {
		err = iOModuleScopeStruct_Set()
		if err != nil {
			return
		}
		iOModuleScopeBlockFunction, err = iOModuleScopeStruct.InvokerNew("block")
	})
	return err
}

// Block is a representation of the C type g_io_module_scope_block.
func (recv *IOModuleScope) Block(basename string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(basename)

	err := iOModuleScopeBlockFunction_Set()
	if err == nil {
		iOModuleScopeBlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var iOModuleScopeFreeFunction *gi.Function
var iOModuleScopeFreeFunction_Once sync.Once

func iOModuleScopeFreeFunction_Set() error {
	var err error
	iOModuleScopeFreeFunction_Once.Do(func() {
		err = iOModuleScopeStruct_Set()
		if err != nil {
			return
		}
		iOModuleScopeFreeFunction, err = iOModuleScopeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_io_module_scope_free.
func (recv *IOModuleScope) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iOModuleScopeFreeFunction_Set()
	if err == nil {
		iOModuleScopeFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var iOSchedulerJobStruct *gi.Struct
var iOSchedulerJobStruct_Once sync.Once

func iOSchedulerJobStruct_Set() error {
	var err error
	iOSchedulerJobStruct_Once.Do(func() {
		iOSchedulerJobStruct, err = gi.StructNew("Gio", "IOSchedulerJob")
	})
	return err
}

type IOSchedulerJob struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop' : parameter 'func' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'g_io_scheduler_job_send_to_mainloop_async' : parameter 'func' of type 'GLib.SourceFunc' not supported

var iOStreamAdapterStruct *gi.Struct
var iOStreamAdapterStruct_Once sync.Once

func iOStreamAdapterStruct_Set() error {
	var err error
	iOStreamAdapterStruct_Once.Do(func() {
		iOStreamAdapterStruct, err = gi.StructNew("Gio", "IOStreamAdapter")
	})
	return err
}

type IOStreamAdapter struct {
	native uintptr
}

var iOStreamClassStruct *gi.Struct
var iOStreamClassStruct_Once sync.Once

func iOStreamClassStruct_Set() error {
	var err error
	iOStreamClassStruct_Once.Do(func() {
		iOStreamClassStruct, err = gi.StructNew("Gio", "IOStreamClass")
	})
	return err
}

type IOStreamClass struct {
	native uintptr
}

var iOStreamPrivateStruct *gi.Struct
var iOStreamPrivateStruct_Once sync.Once

func iOStreamPrivateStruct_Set() error {
	var err error
	iOStreamPrivateStruct_Once.Do(func() {
		iOStreamPrivateStruct, err = gi.StructNew("Gio", "IOStreamPrivate")
	})
	return err
}

type IOStreamPrivate struct {
	native uintptr
}

var iconIfaceStruct *gi.Struct
var iconIfaceStruct_Once sync.Once

func iconIfaceStruct_Set() error {
	var err error
	iconIfaceStruct_Once.Do(func() {
		iconIfaceStruct, err = gi.StructNew("Gio", "IconIface")
	})
	return err
}

type IconIface struct {
	native uintptr
}

var inetAddressClassStruct *gi.Struct
var inetAddressClassStruct_Once sync.Once

func inetAddressClassStruct_Set() error {
	var err error
	inetAddressClassStruct_Once.Do(func() {
		inetAddressClassStruct, err = gi.StructNew("Gio", "InetAddressClass")
	})
	return err
}

type InetAddressClass struct {
	native uintptr
}

var inetAddressMaskClassStruct *gi.Struct
var inetAddressMaskClassStruct_Once sync.Once

func inetAddressMaskClassStruct_Set() error {
	var err error
	inetAddressMaskClassStruct_Once.Do(func() {
		inetAddressMaskClassStruct, err = gi.StructNew("Gio", "InetAddressMaskClass")
	})
	return err
}

type InetAddressMaskClass struct {
	native uintptr
}

var inetAddressMaskPrivateStruct *gi.Struct
var inetAddressMaskPrivateStruct_Once sync.Once

func inetAddressMaskPrivateStruct_Set() error {
	var err error
	inetAddressMaskPrivateStruct_Once.Do(func() {
		inetAddressMaskPrivateStruct, err = gi.StructNew("Gio", "InetAddressMaskPrivate")
	})
	return err
}

type InetAddressMaskPrivate struct {
	native uintptr
}

var inetAddressPrivateStruct *gi.Struct
var inetAddressPrivateStruct_Once sync.Once

func inetAddressPrivateStruct_Set() error {
	var err error
	inetAddressPrivateStruct_Once.Do(func() {
		inetAddressPrivateStruct, err = gi.StructNew("Gio", "InetAddressPrivate")
	})
	return err
}

type InetAddressPrivate struct {
	native uintptr
}

var inetSocketAddressClassStruct *gi.Struct
var inetSocketAddressClassStruct_Once sync.Once

func inetSocketAddressClassStruct_Set() error {
	var err error
	inetSocketAddressClassStruct_Once.Do(func() {
		inetSocketAddressClassStruct, err = gi.StructNew("Gio", "InetSocketAddressClass")
	})
	return err
}

type InetSocketAddressClass struct {
	native uintptr
}

var inetSocketAddressPrivateStruct *gi.Struct
var inetSocketAddressPrivateStruct_Once sync.Once

func inetSocketAddressPrivateStruct_Set() error {
	var err error
	inetSocketAddressPrivateStruct_Once.Do(func() {
		inetSocketAddressPrivateStruct, err = gi.StructNew("Gio", "InetSocketAddressPrivate")
	})
	return err
}

type InetSocketAddressPrivate struct {
	native uintptr
}

var initableIfaceStruct *gi.Struct
var initableIfaceStruct_Once sync.Once

func initableIfaceStruct_Set() error {
	var err error
	initableIfaceStruct_Once.Do(func() {
		initableIfaceStruct, err = gi.StructNew("Gio", "InitableIface")
	})
	return err
}

type InitableIface struct {
	native uintptr
}

var inputMessageStruct *gi.Struct
var inputMessageStruct_Once sync.Once

func inputMessageStruct_Set() error {
	var err error
	inputMessageStruct_Once.Do(func() {
		inputMessageStruct, err = gi.StructNew("Gio", "InputMessage")
	})
	return err
}

type InputMessage struct {
	native uintptr
}

var inputStreamClassStruct *gi.Struct
var inputStreamClassStruct_Once sync.Once

func inputStreamClassStruct_Set() error {
	var err error
	inputStreamClassStruct_Once.Do(func() {
		inputStreamClassStruct, err = gi.StructNew("Gio", "InputStreamClass")
	})
	return err
}

type InputStreamClass struct {
	native uintptr
}

var inputStreamPrivateStruct *gi.Struct
var inputStreamPrivateStruct_Once sync.Once

func inputStreamPrivateStruct_Set() error {
	var err error
	inputStreamPrivateStruct_Once.Do(func() {
		inputStreamPrivateStruct, err = gi.StructNew("Gio", "InputStreamPrivate")
	})
	return err
}

type InputStreamPrivate struct {
	native uintptr
}

var inputVectorStruct *gi.Struct
var inputVectorStruct_Once sync.Once

func inputVectorStruct_Set() error {
	var err error
	inputVectorStruct_Once.Do(func() {
		inputVectorStruct, err = gi.StructNew("Gio", "InputVector")
	})
	return err
}

type InputVector struct {
	native uintptr
}

var listModelInterfaceStruct *gi.Struct
var listModelInterfaceStruct_Once sync.Once

func listModelInterfaceStruct_Set() error {
	var err error
	listModelInterfaceStruct_Once.Do(func() {
		listModelInterfaceStruct, err = gi.StructNew("Gio", "ListModelInterface")
	})
	return err
}

type ListModelInterface struct {
	native uintptr
}

var listStoreClassStruct *gi.Struct
var listStoreClassStruct_Once sync.Once

func listStoreClassStruct_Set() error {
	var err error
	listStoreClassStruct_Once.Do(func() {
		listStoreClassStruct, err = gi.StructNew("Gio", "ListStoreClass")
	})
	return err
}

type ListStoreClass struct {
	native uintptr
}

var loadableIconIfaceStruct *gi.Struct
var loadableIconIfaceStruct_Once sync.Once

func loadableIconIfaceStruct_Set() error {
	var err error
	loadableIconIfaceStruct_Once.Do(func() {
		loadableIconIfaceStruct, err = gi.StructNew("Gio", "LoadableIconIface")
	})
	return err
}

type LoadableIconIface struct {
	native uintptr
}

var memoryInputStreamClassStruct *gi.Struct
var memoryInputStreamClassStruct_Once sync.Once

func memoryInputStreamClassStruct_Set() error {
	var err error
	memoryInputStreamClassStruct_Once.Do(func() {
		memoryInputStreamClassStruct, err = gi.StructNew("Gio", "MemoryInputStreamClass")
	})
	return err
}

type MemoryInputStreamClass struct {
	native uintptr
}

var memoryInputStreamPrivateStruct *gi.Struct
var memoryInputStreamPrivateStruct_Once sync.Once

func memoryInputStreamPrivateStruct_Set() error {
	var err error
	memoryInputStreamPrivateStruct_Once.Do(func() {
		memoryInputStreamPrivateStruct, err = gi.StructNew("Gio", "MemoryInputStreamPrivate")
	})
	return err
}

type MemoryInputStreamPrivate struct {
	native uintptr
}

var memoryOutputStreamClassStruct *gi.Struct
var memoryOutputStreamClassStruct_Once sync.Once

func memoryOutputStreamClassStruct_Set() error {
	var err error
	memoryOutputStreamClassStruct_Once.Do(func() {
		memoryOutputStreamClassStruct, err = gi.StructNew("Gio", "MemoryOutputStreamClass")
	})
	return err
}

type MemoryOutputStreamClass struct {
	native uintptr
}

var memoryOutputStreamPrivateStruct *gi.Struct
var memoryOutputStreamPrivateStruct_Once sync.Once

func memoryOutputStreamPrivateStruct_Set() error {
	var err error
	memoryOutputStreamPrivateStruct_Once.Do(func() {
		memoryOutputStreamPrivateStruct, err = gi.StructNew("Gio", "MemoryOutputStreamPrivate")
	})
	return err
}

type MemoryOutputStreamPrivate struct {
	native uintptr
}

var menuAttributeIterClassStruct *gi.Struct
var menuAttributeIterClassStruct_Once sync.Once

func menuAttributeIterClassStruct_Set() error {
	var err error
	menuAttributeIterClassStruct_Once.Do(func() {
		menuAttributeIterClassStruct, err = gi.StructNew("Gio", "MenuAttributeIterClass")
	})
	return err
}

type MenuAttributeIterClass struct {
	native uintptr
}

var menuAttributeIterPrivateStruct *gi.Struct
var menuAttributeIterPrivateStruct_Once sync.Once

func menuAttributeIterPrivateStruct_Set() error {
	var err error
	menuAttributeIterPrivateStruct_Once.Do(func() {
		menuAttributeIterPrivateStruct, err = gi.StructNew("Gio", "MenuAttributeIterPrivate")
	})
	return err
}

type MenuAttributeIterPrivate struct {
	native uintptr
}

var menuLinkIterClassStruct *gi.Struct
var menuLinkIterClassStruct_Once sync.Once

func menuLinkIterClassStruct_Set() error {
	var err error
	menuLinkIterClassStruct_Once.Do(func() {
		menuLinkIterClassStruct, err = gi.StructNew("Gio", "MenuLinkIterClass")
	})
	return err
}

type MenuLinkIterClass struct {
	native uintptr
}

var menuLinkIterPrivateStruct *gi.Struct
var menuLinkIterPrivateStruct_Once sync.Once

func menuLinkIterPrivateStruct_Set() error {
	var err error
	menuLinkIterPrivateStruct_Once.Do(func() {
		menuLinkIterPrivateStruct, err = gi.StructNew("Gio", "MenuLinkIterPrivate")
	})
	return err
}

type MenuLinkIterPrivate struct {
	native uintptr
}

var menuModelClassStruct *gi.Struct
var menuModelClassStruct_Once sync.Once

func menuModelClassStruct_Set() error {
	var err error
	menuModelClassStruct_Once.Do(func() {
		menuModelClassStruct, err = gi.StructNew("Gio", "MenuModelClass")
	})
	return err
}

type MenuModelClass struct {
	native uintptr
}

var menuModelPrivateStruct *gi.Struct
var menuModelPrivateStruct_Once sync.Once

func menuModelPrivateStruct_Set() error {
	var err error
	menuModelPrivateStruct_Once.Do(func() {
		menuModelPrivateStruct, err = gi.StructNew("Gio", "MenuModelPrivate")
	})
	return err
}

type MenuModelPrivate struct {
	native uintptr
}

var mountIfaceStruct *gi.Struct
var mountIfaceStruct_Once sync.Once

func mountIfaceStruct_Set() error {
	var err error
	mountIfaceStruct_Once.Do(func() {
		mountIfaceStruct, err = gi.StructNew("Gio", "MountIface")
	})
	return err
}

type MountIface struct {
	native uintptr
}

var mountOperationClassStruct *gi.Struct
var mountOperationClassStruct_Once sync.Once

func mountOperationClassStruct_Set() error {
	var err error
	mountOperationClassStruct_Once.Do(func() {
		mountOperationClassStruct, err = gi.StructNew("Gio", "MountOperationClass")
	})
	return err
}

type MountOperationClass struct {
	native uintptr
}

var mountOperationPrivateStruct *gi.Struct
var mountOperationPrivateStruct_Once sync.Once

func mountOperationPrivateStruct_Set() error {
	var err error
	mountOperationPrivateStruct_Once.Do(func() {
		mountOperationPrivateStruct, err = gi.StructNew("Gio", "MountOperationPrivate")
	})
	return err
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeSocketAddressClassStruct *gi.Struct
var nativeSocketAddressClassStruct_Once sync.Once

func nativeSocketAddressClassStruct_Set() error {
	var err error
	nativeSocketAddressClassStruct_Once.Do(func() {
		nativeSocketAddressClassStruct, err = gi.StructNew("Gio", "NativeSocketAddressClass")
	})
	return err
}

type NativeSocketAddressClass struct {
	native uintptr
}

var nativeSocketAddressPrivateStruct *gi.Struct
var nativeSocketAddressPrivateStruct_Once sync.Once

func nativeSocketAddressPrivateStruct_Set() error {
	var err error
	nativeSocketAddressPrivateStruct_Once.Do(func() {
		nativeSocketAddressPrivateStruct, err = gi.StructNew("Gio", "NativeSocketAddressPrivate")
	})
	return err
}

type NativeSocketAddressPrivate struct {
	native uintptr
}

var nativeVolumeMonitorClassStruct *gi.Struct
var nativeVolumeMonitorClassStruct_Once sync.Once

func nativeVolumeMonitorClassStruct_Set() error {
	var err error
	nativeVolumeMonitorClassStruct_Once.Do(func() {
		nativeVolumeMonitorClassStruct, err = gi.StructNew("Gio", "NativeVolumeMonitorClass")
	})
	return err
}

type NativeVolumeMonitorClass struct {
	native uintptr
}

var networkAddressClassStruct *gi.Struct
var networkAddressClassStruct_Once sync.Once

func networkAddressClassStruct_Set() error {
	var err error
	networkAddressClassStruct_Once.Do(func() {
		networkAddressClassStruct, err = gi.StructNew("Gio", "NetworkAddressClass")
	})
	return err
}

type NetworkAddressClass struct {
	native uintptr
}

var networkAddressPrivateStruct *gi.Struct
var networkAddressPrivateStruct_Once sync.Once

func networkAddressPrivateStruct_Set() error {
	var err error
	networkAddressPrivateStruct_Once.Do(func() {
		networkAddressPrivateStruct, err = gi.StructNew("Gio", "NetworkAddressPrivate")
	})
	return err
}

type NetworkAddressPrivate struct {
	native uintptr
}

var networkMonitorInterfaceStruct *gi.Struct
var networkMonitorInterfaceStruct_Once sync.Once

func networkMonitorInterfaceStruct_Set() error {
	var err error
	networkMonitorInterfaceStruct_Once.Do(func() {
		networkMonitorInterfaceStruct, err = gi.StructNew("Gio", "NetworkMonitorInterface")
	})
	return err
}

type NetworkMonitorInterface struct {
	native uintptr
}

var networkServiceClassStruct *gi.Struct
var networkServiceClassStruct_Once sync.Once

func networkServiceClassStruct_Set() error {
	var err error
	networkServiceClassStruct_Once.Do(func() {
		networkServiceClassStruct, err = gi.StructNew("Gio", "NetworkServiceClass")
	})
	return err
}

type NetworkServiceClass struct {
	native uintptr
}

var networkServicePrivateStruct *gi.Struct
var networkServicePrivateStruct_Once sync.Once

func networkServicePrivateStruct_Set() error {
	var err error
	networkServicePrivateStruct_Once.Do(func() {
		networkServicePrivateStruct, err = gi.StructNew("Gio", "NetworkServicePrivate")
	})
	return err
}

type NetworkServicePrivate struct {
	native uintptr
}

var outputMessageStruct *gi.Struct
var outputMessageStruct_Once sync.Once

func outputMessageStruct_Set() error {
	var err error
	outputMessageStruct_Once.Do(func() {
		outputMessageStruct, err = gi.StructNew("Gio", "OutputMessage")
	})
	return err
}

type OutputMessage struct {
	native uintptr
}

var outputStreamClassStruct *gi.Struct
var outputStreamClassStruct_Once sync.Once

func outputStreamClassStruct_Set() error {
	var err error
	outputStreamClassStruct_Once.Do(func() {
		outputStreamClassStruct, err = gi.StructNew("Gio", "OutputStreamClass")
	})
	return err
}

type OutputStreamClass struct {
	native uintptr
}

var outputStreamPrivateStruct *gi.Struct
var outputStreamPrivateStruct_Once sync.Once

func outputStreamPrivateStruct_Set() error {
	var err error
	outputStreamPrivateStruct_Once.Do(func() {
		outputStreamPrivateStruct, err = gi.StructNew("Gio", "OutputStreamPrivate")
	})
	return err
}

type OutputStreamPrivate struct {
	native uintptr
}

var outputVectorStruct *gi.Struct
var outputVectorStruct_Once sync.Once

func outputVectorStruct_Set() error {
	var err error
	outputVectorStruct_Once.Do(func() {
		outputVectorStruct, err = gi.StructNew("Gio", "OutputVector")
	})
	return err
}

type OutputVector struct {
	native uintptr
}

var permissionClassStruct *gi.Struct
var permissionClassStruct_Once sync.Once

func permissionClassStruct_Set() error {
	var err error
	permissionClassStruct_Once.Do(func() {
		permissionClassStruct, err = gi.StructNew("Gio", "PermissionClass")
	})
	return err
}

type PermissionClass struct {
	native uintptr
}

var permissionPrivateStruct *gi.Struct
var permissionPrivateStruct_Once sync.Once

func permissionPrivateStruct_Set() error {
	var err error
	permissionPrivateStruct_Once.Do(func() {
		permissionPrivateStruct, err = gi.StructNew("Gio", "PermissionPrivate")
	})
	return err
}

type PermissionPrivate struct {
	native uintptr
}

var pollableInputStreamInterfaceStruct *gi.Struct
var pollableInputStreamInterfaceStruct_Once sync.Once

func pollableInputStreamInterfaceStruct_Set() error {
	var err error
	pollableInputStreamInterfaceStruct_Once.Do(func() {
		pollableInputStreamInterfaceStruct, err = gi.StructNew("Gio", "PollableInputStreamInterface")
	})
	return err
}

type PollableInputStreamInterface struct {
	native uintptr
}

var pollableOutputStreamInterfaceStruct *gi.Struct
var pollableOutputStreamInterfaceStruct_Once sync.Once

func pollableOutputStreamInterfaceStruct_Set() error {
	var err error
	pollableOutputStreamInterfaceStruct_Once.Do(func() {
		pollableOutputStreamInterfaceStruct, err = gi.StructNew("Gio", "PollableOutputStreamInterface")
	})
	return err
}

type PollableOutputStreamInterface struct {
	native uintptr
}

var proxyAddressClassStruct *gi.Struct
var proxyAddressClassStruct_Once sync.Once

func proxyAddressClassStruct_Set() error {
	var err error
	proxyAddressClassStruct_Once.Do(func() {
		proxyAddressClassStruct, err = gi.StructNew("Gio", "ProxyAddressClass")
	})
	return err
}

type ProxyAddressClass struct {
	native uintptr
}

var proxyAddressEnumeratorClassStruct *gi.Struct
var proxyAddressEnumeratorClassStruct_Once sync.Once

func proxyAddressEnumeratorClassStruct_Set() error {
	var err error
	proxyAddressEnumeratorClassStruct_Once.Do(func() {
		proxyAddressEnumeratorClassStruct, err = gi.StructNew("Gio", "ProxyAddressEnumeratorClass")
	})
	return err
}

type ProxyAddressEnumeratorClass struct {
	native uintptr
}

var proxyAddressEnumeratorPrivateStruct *gi.Struct
var proxyAddressEnumeratorPrivateStruct_Once sync.Once

func proxyAddressEnumeratorPrivateStruct_Set() error {
	var err error
	proxyAddressEnumeratorPrivateStruct_Once.Do(func() {
		proxyAddressEnumeratorPrivateStruct, err = gi.StructNew("Gio", "ProxyAddressEnumeratorPrivate")
	})
	return err
}

type ProxyAddressEnumeratorPrivate struct {
	native uintptr
}

var proxyAddressPrivateStruct *gi.Struct
var proxyAddressPrivateStruct_Once sync.Once

func proxyAddressPrivateStruct_Set() error {
	var err error
	proxyAddressPrivateStruct_Once.Do(func() {
		proxyAddressPrivateStruct, err = gi.StructNew("Gio", "ProxyAddressPrivate")
	})
	return err
}

type ProxyAddressPrivate struct {
	native uintptr
}

var proxyInterfaceStruct *gi.Struct
var proxyInterfaceStruct_Once sync.Once

func proxyInterfaceStruct_Set() error {
	var err error
	proxyInterfaceStruct_Once.Do(func() {
		proxyInterfaceStruct, err = gi.StructNew("Gio", "ProxyInterface")
	})
	return err
}

type ProxyInterface struct {
	native uintptr
}

var proxyResolverInterfaceStruct *gi.Struct
var proxyResolverInterfaceStruct_Once sync.Once

func proxyResolverInterfaceStruct_Set() error {
	var err error
	proxyResolverInterfaceStruct_Once.Do(func() {
		proxyResolverInterfaceStruct, err = gi.StructNew("Gio", "ProxyResolverInterface")
	})
	return err
}

type ProxyResolverInterface struct {
	native uintptr
}

var remoteActionGroupInterfaceStruct *gi.Struct
var remoteActionGroupInterfaceStruct_Once sync.Once

func remoteActionGroupInterfaceStruct_Set() error {
	var err error
	remoteActionGroupInterfaceStruct_Once.Do(func() {
		remoteActionGroupInterfaceStruct, err = gi.StructNew("Gio", "RemoteActionGroupInterface")
	})
	return err
}

type RemoteActionGroupInterface struct {
	native uintptr
}

var resolverClassStruct *gi.Struct
var resolverClassStruct_Once sync.Once

func resolverClassStruct_Set() error {
	var err error
	resolverClassStruct_Once.Do(func() {
		resolverClassStruct, err = gi.StructNew("Gio", "ResolverClass")
	})
	return err
}

type ResolverClass struct {
	native uintptr
}

var resolverPrivateStruct *gi.Struct
var resolverPrivateStruct_Once sync.Once

func resolverPrivateStruct_Set() error {
	var err error
	resolverPrivateStruct_Once.Do(func() {
		resolverPrivateStruct, err = gi.StructNew("Gio", "ResolverPrivate")
	})
	return err
}

type ResolverPrivate struct {
	native uintptr
}

var resourceStruct *gi.Struct
var resourceStruct_Once sync.Once

func resourceStruct_Set() error {
	var err error
	resourceStruct_Once.Do(func() {
		resourceStruct, err = gi.StructNew("Gio", "Resource")
	})
	return err
}

type Resource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_resource_new_from_data' : parameter 'data' of type 'GLib.Bytes' not supported

var resourceRegisterFunction *gi.Function
var resourceRegisterFunction_Once sync.Once

func resourceRegisterFunction_Set() error {
	var err error
	resourceRegisterFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceRegisterFunction, err = resourceStruct.InvokerNew("_register")
	})
	return err
}

// Register is a representation of the C type g_resources_register.
func (recv *Resource) Register() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceRegisterFunction_Set()
	if err == nil {
		resourceRegisterFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var resourceUnregisterFunction *gi.Function
var resourceUnregisterFunction_Once sync.Once

func resourceUnregisterFunction_Set() error {
	var err error
	resourceUnregisterFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceUnregisterFunction, err = resourceStruct.InvokerNew("_unregister")
	})
	return err
}

// Unregister is a representation of the C type g_resources_unregister.
func (recv *Resource) Unregister() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceUnregisterFunction_Set()
	if err == nil {
		resourceUnregisterFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_resource_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resource_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

var resourceRefFunction *gi.Function
var resourceRefFunction_Once sync.Once

func resourceRefFunction_Set() error {
	var err error
	resourceRefFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceRefFunction, err = resourceStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_resource_ref.
func (recv *Resource) Ref() (*Resource, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := resourceRefFunction_Set()
	if err == nil {
		ret = resourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Resource{native: ret.Pointer()}

	return retGo, err
}

var resourceUnrefFunction *gi.Function
var resourceUnrefFunction_Once sync.Once

func resourceUnrefFunction_Set() error {
	var err error
	resourceUnrefFunction_Once.Do(func() {
		err = resourceStruct_Set()
		if err != nil {
			return
		}
		resourceUnrefFunction, err = resourceStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_resource_unref.
func (recv *Resource) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := resourceUnrefFunction_Set()
	if err == nil {
		resourceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var seekableIfaceStruct *gi.Struct
var seekableIfaceStruct_Once sync.Once

func seekableIfaceStruct_Set() error {
	var err error
	seekableIfaceStruct_Once.Do(func() {
		seekableIfaceStruct, err = gi.StructNew("Gio", "SeekableIface")
	})
	return err
}

type SeekableIface struct {
	native uintptr
}

var settingsBackendClassStruct *gi.Struct
var settingsBackendClassStruct_Once sync.Once

func settingsBackendClassStruct_Set() error {
	var err error
	settingsBackendClassStruct_Once.Do(func() {
		settingsBackendClassStruct, err = gi.StructNew("Gio", "SettingsBackendClass")
	})
	return err
}

type SettingsBackendClass struct {
	native uintptr
}

var settingsBackendPrivateStruct *gi.Struct
var settingsBackendPrivateStruct_Once sync.Once

func settingsBackendPrivateStruct_Set() error {
	var err error
	settingsBackendPrivateStruct_Once.Do(func() {
		settingsBackendPrivateStruct, err = gi.StructNew("Gio", "SettingsBackendPrivate")
	})
	return err
}

type SettingsBackendPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() error {
	var err error
	settingsClassStruct_Once.Do(func() {
		settingsClassStruct, err = gi.StructNew("Gio", "SettingsClass")
	})
	return err
}

type SettingsClass struct {
	native uintptr
}

var settingsPrivateStruct *gi.Struct
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() error {
	var err error
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct, err = gi.StructNew("Gio", "SettingsPrivate")
	})
	return err
}

type SettingsPrivate struct {
	native uintptr
}

var settingsSchemaStruct *gi.Struct
var settingsSchemaStruct_Once sync.Once

func settingsSchemaStruct_Set() error {
	var err error
	settingsSchemaStruct_Once.Do(func() {
		settingsSchemaStruct, err = gi.StructNew("Gio", "SettingsSchema")
	})
	return err
}

type SettingsSchema struct {
	native uintptr
}

var settingsSchemaGetIdFunction *gi.Function
var settingsSchemaGetIdFunction_Once sync.Once

func settingsSchemaGetIdFunction_Set() error {
	var err error
	settingsSchemaGetIdFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetIdFunction, err = settingsSchemaStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type g_settings_schema_get_id.
func (recv *SettingsSchema) GetId() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaGetIdFunction_Set()
	if err == nil {
		ret = settingsSchemaGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var settingsSchemaGetKeyFunction *gi.Function
var settingsSchemaGetKeyFunction_Once sync.Once

func settingsSchemaGetKeyFunction_Set() error {
	var err error
	settingsSchemaGetKeyFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetKeyFunction, err = settingsSchemaStruct.InvokerNew("get_key")
	})
	return err
}

// GetKey is a representation of the C type g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) (*SettingsSchemaKey, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsSchemaGetKeyFunction_Set()
	if err == nil {
		ret = settingsSchemaGetKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo, err
}

var settingsSchemaGetPathFunction *gi.Function
var settingsSchemaGetPathFunction_Once sync.Once

func settingsSchemaGetPathFunction_Set() error {
	var err error
	settingsSchemaGetPathFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaGetPathFunction, err = settingsSchemaStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type g_settings_schema_get_path.
func (recv *SettingsSchema) GetPath() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaGetPathFunction_Set()
	if err == nil {
		ret = settingsSchemaGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var settingsSchemaHasKeyFunction *gi.Function
var settingsSchemaHasKeyFunction_Once sync.Once

func settingsSchemaHasKeyFunction_Set() error {
	var err error
	settingsSchemaHasKeyFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaHasKeyFunction, err = settingsSchemaStruct.InvokerNew("has_key")
	})
	return err
}

// HasKey is a representation of the C type g_settings_schema_has_key.
func (recv *SettingsSchema) HasKey(name string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := settingsSchemaHasKeyFunction_Set()
	if err == nil {
		ret = settingsSchemaHasKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var settingsSchemaListChildrenFunction *gi.Function
var settingsSchemaListChildrenFunction_Once sync.Once

func settingsSchemaListChildrenFunction_Set() error {
	var err error
	settingsSchemaListChildrenFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaListChildrenFunction, err = settingsSchemaStruct.InvokerNew("list_children")
	})
	return err
}

// ListChildren is a representation of the C type g_settings_schema_list_children.
func (recv *SettingsSchema) ListChildren() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaListChildrenFunction_Set()
	if err == nil {
		settingsSchemaListChildrenFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var settingsSchemaListKeysFunction *gi.Function
var settingsSchemaListKeysFunction_Once sync.Once

func settingsSchemaListKeysFunction_Set() error {
	var err error
	settingsSchemaListKeysFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaListKeysFunction, err = settingsSchemaStruct.InvokerNew("list_keys")
	})
	return err
}

// ListKeys is a representation of the C type g_settings_schema_list_keys.
func (recv *SettingsSchema) ListKeys() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaListKeysFunction_Set()
	if err == nil {
		settingsSchemaListKeysFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var settingsSchemaRefFunction *gi.Function
var settingsSchemaRefFunction_Once sync.Once

func settingsSchemaRefFunction_Set() error {
	var err error
	settingsSchemaRefFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaRefFunction, err = settingsSchemaStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_ref.
func (recv *SettingsSchema) Ref() (*SettingsSchema, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaRefFunction_Set()
	if err == nil {
		ret = settingsSchemaRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchema{native: ret.Pointer()}

	return retGo, err
}

var settingsSchemaUnrefFunction *gi.Function
var settingsSchemaUnrefFunction_Once sync.Once

func settingsSchemaUnrefFunction_Set() error {
	var err error
	settingsSchemaUnrefFunction_Once.Do(func() {
		err = settingsSchemaStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaUnrefFunction, err = settingsSchemaStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_unref.
func (recv *SettingsSchema) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaUnrefFunction_Set()
	if err == nil {
		settingsSchemaUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var settingsSchemaKeyStruct *gi.Struct
var settingsSchemaKeyStruct_Once sync.Once

func settingsSchemaKeyStruct_Set() error {
	var err error
	settingsSchemaKeyStruct_Once.Do(func() {
		settingsSchemaKeyStruct, err = gi.StructNew("Gio", "SettingsSchemaKey")
	})
	return err
}

type SettingsSchemaKey struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_default_value' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetDescriptionFunction *gi.Function
var settingsSchemaKeyGetDescriptionFunction_Once sync.Once

func settingsSchemaKeyGetDescriptionFunction_Set() error {
	var err error
	settingsSchemaKeyGetDescriptionFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetDescriptionFunction, err = settingsSchemaKeyStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_settings_schema_key_get_description.
func (recv *SettingsSchemaKey) GetDescription() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetDescriptionFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var settingsSchemaKeyGetNameFunction *gi.Function
var settingsSchemaKeyGetNameFunction_Once sync.Once

func settingsSchemaKeyGetNameFunction_Set() error {
	var err error
	settingsSchemaKeyGetNameFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetNameFunction, err = settingsSchemaKeyStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_settings_schema_key_get_name.
func (recv *SettingsSchemaKey) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetNameFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_range' : return type 'GLib.Variant' not supported

var settingsSchemaKeyGetSummaryFunction *gi.Function
var settingsSchemaKeyGetSummaryFunction_Once sync.Once

func settingsSchemaKeyGetSummaryFunction_Set() error {
	var err error
	settingsSchemaKeyGetSummaryFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyGetSummaryFunction, err = settingsSchemaKeyStruct.InvokerNew("get_summary")
	})
	return err
}

// GetSummary is a representation of the C type g_settings_schema_key_get_summary.
func (recv *SettingsSchemaKey) GetSummary() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyGetSummaryFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyGetSummaryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_settings_schema_key_get_value_type' : return type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_settings_schema_key_range_check' : parameter 'value' of type 'GLib.Variant' not supported

var settingsSchemaKeyRefFunction *gi.Function
var settingsSchemaKeyRefFunction_Once sync.Once

func settingsSchemaKeyRefFunction_Set() error {
	var err error
	settingsSchemaKeyRefFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyRefFunction, err = settingsSchemaKeyStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() (*SettingsSchemaKey, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaKeyRefFunction_Set()
	if err == nil {
		ret = settingsSchemaKeyRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaKey{native: ret.Pointer()}

	return retGo, err
}

var settingsSchemaKeyUnrefFunction *gi.Function
var settingsSchemaKeyUnrefFunction_Once sync.Once

func settingsSchemaKeyUnrefFunction_Set() error {
	var err error
	settingsSchemaKeyUnrefFunction_Once.Do(func() {
		err = settingsSchemaKeyStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaKeyUnrefFunction, err = settingsSchemaKeyStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaKeyUnrefFunction_Set()
	if err == nil {
		settingsSchemaKeyUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var settingsSchemaSourceStruct *gi.Struct
var settingsSchemaSourceStruct_Once sync.Once

func settingsSchemaSourceStruct_Set() error {
	var err error
	settingsSchemaSourceStruct_Once.Do(func() {
		settingsSchemaSourceStruct, err = gi.StructNew("Gio", "SettingsSchemaSource")
	})
	return err
}

type SettingsSchemaSource struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_settings_schema_source_new_from_directory' : parameter 'directory' of type 'filename' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_list_schemas' : parameter 'non_relocatable' has no type

var settingsSchemaSourceLookupFunction *gi.Function
var settingsSchemaSourceLookupFunction_Once sync.Once

func settingsSchemaSourceLookupFunction_Set() error {
	var err error
	settingsSchemaSourceLookupFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceLookupFunction, err = settingsSchemaSourceStruct.InvokerNew("lookup")
	})
	return err
}

// Lookup is a representation of the C type g_settings_schema_source_lookup.
func (recv *SettingsSchemaSource) Lookup(schemaId string, recursive bool) (*SettingsSchema, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(schemaId)
	inArgs[2].SetBoolean(recursive)

	var ret gi.Argument

	err := settingsSchemaSourceLookupFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchema{native: ret.Pointer()}

	return retGo, err
}

var settingsSchemaSourceRefFunction *gi.Function
var settingsSchemaSourceRefFunction_Once sync.Once

func settingsSchemaSourceRefFunction_Set() error {
	var err error
	settingsSchemaSourceRefFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceRefFunction, err = settingsSchemaSourceStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_settings_schema_source_ref.
func (recv *SettingsSchemaSource) Ref() (*SettingsSchemaSource, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := settingsSchemaSourceRefFunction_Set()
	if err == nil {
		ret = settingsSchemaSourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo, err
}

var settingsSchemaSourceUnrefFunction *gi.Function
var settingsSchemaSourceUnrefFunction_Once sync.Once

func settingsSchemaSourceUnrefFunction_Set() error {
	var err error
	settingsSchemaSourceUnrefFunction_Once.Do(func() {
		err = settingsSchemaSourceStruct_Set()
		if err != nil {
			return
		}
		settingsSchemaSourceUnrefFunction, err = settingsSchemaSourceStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_settings_schema_source_unref.
func (recv *SettingsSchemaSource) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := settingsSchemaSourceUnrefFunction_Set()
	if err == nil {
		settingsSchemaSourceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var simpleActionGroupClassStruct *gi.Struct
var simpleActionGroupClassStruct_Once sync.Once

func simpleActionGroupClassStruct_Set() error {
	var err error
	simpleActionGroupClassStruct_Once.Do(func() {
		simpleActionGroupClassStruct, err = gi.StructNew("Gio", "SimpleActionGroupClass")
	})
	return err
}

type SimpleActionGroupClass struct {
	native uintptr
}

var simpleActionGroupPrivateStruct *gi.Struct
var simpleActionGroupPrivateStruct_Once sync.Once

func simpleActionGroupPrivateStruct_Set() error {
	var err error
	simpleActionGroupPrivateStruct_Once.Do(func() {
		simpleActionGroupPrivateStruct, err = gi.StructNew("Gio", "SimpleActionGroupPrivate")
	})
	return err
}

type SimpleActionGroupPrivate struct {
	native uintptr
}

var simpleAsyncResultClassStruct *gi.Struct
var simpleAsyncResultClassStruct_Once sync.Once

func simpleAsyncResultClassStruct_Set() error {
	var err error
	simpleAsyncResultClassStruct_Once.Do(func() {
		simpleAsyncResultClassStruct, err = gi.StructNew("Gio", "SimpleAsyncResultClass")
	})
	return err
}

type SimpleAsyncResultClass struct {
	native uintptr
}

var simpleProxyResolverClassStruct *gi.Struct
var simpleProxyResolverClassStruct_Once sync.Once

func simpleProxyResolverClassStruct_Set() error {
	var err error
	simpleProxyResolverClassStruct_Once.Do(func() {
		simpleProxyResolverClassStruct, err = gi.StructNew("Gio", "SimpleProxyResolverClass")
	})
	return err
}

type SimpleProxyResolverClass struct {
	native uintptr
}

var simpleProxyResolverPrivateStruct *gi.Struct
var simpleProxyResolverPrivateStruct_Once sync.Once

func simpleProxyResolverPrivateStruct_Set() error {
	var err error
	simpleProxyResolverPrivateStruct_Once.Do(func() {
		simpleProxyResolverPrivateStruct, err = gi.StructNew("Gio", "SimpleProxyResolverPrivate")
	})
	return err
}

type SimpleProxyResolverPrivate struct {
	native uintptr
}

var socketAddressClassStruct *gi.Struct
var socketAddressClassStruct_Once sync.Once

func socketAddressClassStruct_Set() error {
	var err error
	socketAddressClassStruct_Once.Do(func() {
		socketAddressClassStruct, err = gi.StructNew("Gio", "SocketAddressClass")
	})
	return err
}

type SocketAddressClass struct {
	native uintptr
}

var socketAddressEnumeratorClassStruct *gi.Struct
var socketAddressEnumeratorClassStruct_Once sync.Once

func socketAddressEnumeratorClassStruct_Set() error {
	var err error
	socketAddressEnumeratorClassStruct_Once.Do(func() {
		socketAddressEnumeratorClassStruct, err = gi.StructNew("Gio", "SocketAddressEnumeratorClass")
	})
	return err
}

type SocketAddressEnumeratorClass struct {
	native uintptr
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Gio", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native uintptr
}

var socketClientClassStruct *gi.Struct
var socketClientClassStruct_Once sync.Once

func socketClientClassStruct_Set() error {
	var err error
	socketClientClassStruct_Once.Do(func() {
		socketClientClassStruct, err = gi.StructNew("Gio", "SocketClientClass")
	})
	return err
}

type SocketClientClass struct {
	native uintptr
}

var socketClientPrivateStruct *gi.Struct
var socketClientPrivateStruct_Once sync.Once

func socketClientPrivateStruct_Set() error {
	var err error
	socketClientPrivateStruct_Once.Do(func() {
		socketClientPrivateStruct, err = gi.StructNew("Gio", "SocketClientPrivate")
	})
	return err
}

type SocketClientPrivate struct {
	native uintptr
}

var socketConnectableIfaceStruct *gi.Struct
var socketConnectableIfaceStruct_Once sync.Once

func socketConnectableIfaceStruct_Set() error {
	var err error
	socketConnectableIfaceStruct_Once.Do(func() {
		socketConnectableIfaceStruct, err = gi.StructNew("Gio", "SocketConnectableIface")
	})
	return err
}

type SocketConnectableIface struct {
	native uintptr
}

var socketConnectionClassStruct *gi.Struct
var socketConnectionClassStruct_Once sync.Once

func socketConnectionClassStruct_Set() error {
	var err error
	socketConnectionClassStruct_Once.Do(func() {
		socketConnectionClassStruct, err = gi.StructNew("Gio", "SocketConnectionClass")
	})
	return err
}

type SocketConnectionClass struct {
	native uintptr
}

var socketConnectionPrivateStruct *gi.Struct
var socketConnectionPrivateStruct_Once sync.Once

func socketConnectionPrivateStruct_Set() error {
	var err error
	socketConnectionPrivateStruct_Once.Do(func() {
		socketConnectionPrivateStruct, err = gi.StructNew("Gio", "SocketConnectionPrivate")
	})
	return err
}

type SocketConnectionPrivate struct {
	native uintptr
}

var socketControlMessageClassStruct *gi.Struct
var socketControlMessageClassStruct_Once sync.Once

func socketControlMessageClassStruct_Set() error {
	var err error
	socketControlMessageClassStruct_Once.Do(func() {
		socketControlMessageClassStruct, err = gi.StructNew("Gio", "SocketControlMessageClass")
	})
	return err
}

type SocketControlMessageClass struct {
	native uintptr
}

var socketControlMessagePrivateStruct *gi.Struct
var socketControlMessagePrivateStruct_Once sync.Once

func socketControlMessagePrivateStruct_Set() error {
	var err error
	socketControlMessagePrivateStruct_Once.Do(func() {
		socketControlMessagePrivateStruct, err = gi.StructNew("Gio", "SocketControlMessagePrivate")
	})
	return err
}

type SocketControlMessagePrivate struct {
	native uintptr
}

var socketListenerClassStruct *gi.Struct
var socketListenerClassStruct_Once sync.Once

func socketListenerClassStruct_Set() error {
	var err error
	socketListenerClassStruct_Once.Do(func() {
		socketListenerClassStruct, err = gi.StructNew("Gio", "SocketListenerClass")
	})
	return err
}

type SocketListenerClass struct {
	native uintptr
}

var socketListenerPrivateStruct *gi.Struct
var socketListenerPrivateStruct_Once sync.Once

func socketListenerPrivateStruct_Set() error {
	var err error
	socketListenerPrivateStruct_Once.Do(func() {
		socketListenerPrivateStruct, err = gi.StructNew("Gio", "SocketListenerPrivate")
	})
	return err
}

type SocketListenerPrivate struct {
	native uintptr
}

var socketPrivateStruct *gi.Struct
var socketPrivateStruct_Once sync.Once

func socketPrivateStruct_Set() error {
	var err error
	socketPrivateStruct_Once.Do(func() {
		socketPrivateStruct, err = gi.StructNew("Gio", "SocketPrivate")
	})
	return err
}

type SocketPrivate struct {
	native uintptr
}

var socketServiceClassStruct *gi.Struct
var socketServiceClassStruct_Once sync.Once

func socketServiceClassStruct_Set() error {
	var err error
	socketServiceClassStruct_Once.Do(func() {
		socketServiceClassStruct, err = gi.StructNew("Gio", "SocketServiceClass")
	})
	return err
}

type SocketServiceClass struct {
	native uintptr
}

var socketServicePrivateStruct *gi.Struct
var socketServicePrivateStruct_Once sync.Once

func socketServicePrivateStruct_Set() error {
	var err error
	socketServicePrivateStruct_Once.Do(func() {
		socketServicePrivateStruct, err = gi.StructNew("Gio", "SocketServicePrivate")
	})
	return err
}

type SocketServicePrivate struct {
	native uintptr
}

var srvTargetStruct *gi.Struct
var srvTargetStruct_Once sync.Once

func srvTargetStruct_Set() error {
	var err error
	srvTargetStruct_Once.Do(func() {
		srvTargetStruct, err = gi.StructNew("Gio", "SrvTarget")
	})
	return err
}

type SrvTarget struct {
	native uintptr
}

var srvTargetNewFunction *gi.Function
var srvTargetNewFunction_Once sync.Once

func srvTargetNewFunction_Set() error {
	var err error
	srvTargetNewFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetNewFunction, err = srvTargetStruct.InvokerNew("new")
	})
	return err
}

// SrvTargetNew is a representation of the C type g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) (*SrvTarget, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(hostname)
	inArgs[1].SetUint16(port)
	inArgs[2].SetUint16(priority)
	inArgs[3].SetUint16(weight)

	var ret gi.Argument

	err := srvTargetNewFunction_Set()
	if err == nil {
		ret = srvTargetNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo, err
}

var srvTargetCopyFunction *gi.Function
var srvTargetCopyFunction_Once sync.Once

func srvTargetCopyFunction_Set() error {
	var err error
	srvTargetCopyFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetCopyFunction, err = srvTargetStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_srv_target_copy.
func (recv *SrvTarget) Copy() (*SrvTarget, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetCopyFunction_Set()
	if err == nil {
		ret = srvTargetCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SrvTarget{native: ret.Pointer()}

	return retGo, err
}

var srvTargetFreeFunction *gi.Function
var srvTargetFreeFunction_Once sync.Once

func srvTargetFreeFunction_Set() error {
	var err error
	srvTargetFreeFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetFreeFunction, err = srvTargetStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_srv_target_free.
func (recv *SrvTarget) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := srvTargetFreeFunction_Set()
	if err == nil {
		srvTargetFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var srvTargetGetHostnameFunction *gi.Function
var srvTargetGetHostnameFunction_Once sync.Once

func srvTargetGetHostnameFunction_Set() error {
	var err error
	srvTargetGetHostnameFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetHostnameFunction, err = srvTargetStruct.InvokerNew("get_hostname")
	})
	return err
}

// GetHostname is a representation of the C type g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetHostnameFunction_Set()
	if err == nil {
		ret = srvTargetGetHostnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var srvTargetGetPortFunction *gi.Function
var srvTargetGetPortFunction_Once sync.Once

func srvTargetGetPortFunction_Set() error {
	var err error
	srvTargetGetPortFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetPortFunction, err = srvTargetStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type g_srv_target_get_port.
func (recv *SrvTarget) GetPort() (uint16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetPortFunction_Set()
	if err == nil {
		ret = srvTargetGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo, err
}

var srvTargetGetPriorityFunction *gi.Function
var srvTargetGetPriorityFunction_Once sync.Once

func srvTargetGetPriorityFunction_Set() error {
	var err error
	srvTargetGetPriorityFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetPriorityFunction, err = srvTargetStruct.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() (uint16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetPriorityFunction_Set()
	if err == nil {
		ret = srvTargetGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo, err
}

var srvTargetGetWeightFunction *gi.Function
var srvTargetGetWeightFunction_Once sync.Once

func srvTargetGetWeightFunction_Set() error {
	var err error
	srvTargetGetWeightFunction_Once.Do(func() {
		err = srvTargetStruct_Set()
		if err != nil {
			return
		}
		srvTargetGetWeightFunction, err = srvTargetStruct.InvokerNew("get_weight")
	})
	return err
}

// GetWeight is a representation of the C type g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() (uint16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := srvTargetGetWeightFunction_Set()
	if err == nil {
		ret = srvTargetGetWeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo, err
}

var staticResourceStruct *gi.Struct
var staticResourceStruct_Once sync.Once

func staticResourceStruct_Set() error {
	var err error
	staticResourceStruct_Once.Do(func() {
		staticResourceStruct, err = gi.StructNew("Gio", "StaticResource")
	})
	return err
}

type StaticResource struct {
	native uintptr
}

var staticResourceFiniFunction *gi.Function
var staticResourceFiniFunction_Once sync.Once

func staticResourceFiniFunction_Set() error {
	var err error
	staticResourceFiniFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceFiniFunction, err = staticResourceStruct.InvokerNew("fini")
	})
	return err
}

// Fini is a representation of the C type g_static_resource_fini.
func (recv *StaticResource) Fini() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := staticResourceFiniFunction_Set()
	if err == nil {
		staticResourceFiniFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var staticResourceGetResourceFunction *gi.Function
var staticResourceGetResourceFunction_Once sync.Once

func staticResourceGetResourceFunction_Set() error {
	var err error
	staticResourceGetResourceFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceGetResourceFunction, err = staticResourceStruct.InvokerNew("get_resource")
	})
	return err
}

// GetResource is a representation of the C type g_static_resource_get_resource.
func (recv *StaticResource) GetResource() (*Resource, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := staticResourceGetResourceFunction_Set()
	if err == nil {
		ret = staticResourceGetResourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Resource{native: ret.Pointer()}

	return retGo, err
}

var staticResourceInitFunction *gi.Function
var staticResourceInitFunction_Once sync.Once

func staticResourceInitFunction_Set() error {
	var err error
	staticResourceInitFunction_Once.Do(func() {
		err = staticResourceStruct_Set()
		if err != nil {
			return
		}
		staticResourceInitFunction, err = staticResourceStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_static_resource_init.
func (recv *StaticResource) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := staticResourceInitFunction_Set()
	if err == nil {
		staticResourceInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var taskClassStruct *gi.Struct
var taskClassStruct_Once sync.Once

func taskClassStruct_Set() error {
	var err error
	taskClassStruct_Once.Do(func() {
		taskClassStruct, err = gi.StructNew("Gio", "TaskClass")
	})
	return err
}

type TaskClass struct {
	native uintptr
}

var tcpConnectionClassStruct *gi.Struct
var tcpConnectionClassStruct_Once sync.Once

func tcpConnectionClassStruct_Set() error {
	var err error
	tcpConnectionClassStruct_Once.Do(func() {
		tcpConnectionClassStruct, err = gi.StructNew("Gio", "TcpConnectionClass")
	})
	return err
}

type TcpConnectionClass struct {
	native uintptr
}

var tcpConnectionPrivateStruct *gi.Struct
var tcpConnectionPrivateStruct_Once sync.Once

func tcpConnectionPrivateStruct_Set() error {
	var err error
	tcpConnectionPrivateStruct_Once.Do(func() {
		tcpConnectionPrivateStruct, err = gi.StructNew("Gio", "TcpConnectionPrivate")
	})
	return err
}

type TcpConnectionPrivate struct {
	native uintptr
}

var tcpWrapperConnectionClassStruct *gi.Struct
var tcpWrapperConnectionClassStruct_Once sync.Once

func tcpWrapperConnectionClassStruct_Set() error {
	var err error
	tcpWrapperConnectionClassStruct_Once.Do(func() {
		tcpWrapperConnectionClassStruct, err = gi.StructNew("Gio", "TcpWrapperConnectionClass")
	})
	return err
}

type TcpWrapperConnectionClass struct {
	native uintptr
}

var tcpWrapperConnectionPrivateStruct *gi.Struct
var tcpWrapperConnectionPrivateStruct_Once sync.Once

func tcpWrapperConnectionPrivateStruct_Set() error {
	var err error
	tcpWrapperConnectionPrivateStruct_Once.Do(func() {
		tcpWrapperConnectionPrivateStruct, err = gi.StructNew("Gio", "TcpWrapperConnectionPrivate")
	})
	return err
}

type TcpWrapperConnectionPrivate struct {
	native uintptr
}

var themedIconClassStruct *gi.Struct
var themedIconClassStruct_Once sync.Once

func themedIconClassStruct_Set() error {
	var err error
	themedIconClassStruct_Once.Do(func() {
		themedIconClassStruct, err = gi.StructNew("Gio", "ThemedIconClass")
	})
	return err
}

type ThemedIconClass struct {
	native uintptr
}

var threadedSocketServiceClassStruct *gi.Struct
var threadedSocketServiceClassStruct_Once sync.Once

func threadedSocketServiceClassStruct_Set() error {
	var err error
	threadedSocketServiceClassStruct_Once.Do(func() {
		threadedSocketServiceClassStruct, err = gi.StructNew("Gio", "ThreadedSocketServiceClass")
	})
	return err
}

type ThreadedSocketServiceClass struct {
	native uintptr
}

var threadedSocketServicePrivateStruct *gi.Struct
var threadedSocketServicePrivateStruct_Once sync.Once

func threadedSocketServicePrivateStruct_Set() error {
	var err error
	threadedSocketServicePrivateStruct_Once.Do(func() {
		threadedSocketServicePrivateStruct, err = gi.StructNew("Gio", "ThreadedSocketServicePrivate")
	})
	return err
}

type ThreadedSocketServicePrivate struct {
	native uintptr
}

var tlsBackendInterfaceStruct *gi.Struct
var tlsBackendInterfaceStruct_Once sync.Once

func tlsBackendInterfaceStruct_Set() error {
	var err error
	tlsBackendInterfaceStruct_Once.Do(func() {
		tlsBackendInterfaceStruct, err = gi.StructNew("Gio", "TlsBackendInterface")
	})
	return err
}

type TlsBackendInterface struct {
	native uintptr
}

var tlsCertificateClassStruct *gi.Struct
var tlsCertificateClassStruct_Once sync.Once

func tlsCertificateClassStruct_Set() error {
	var err error
	tlsCertificateClassStruct_Once.Do(func() {
		tlsCertificateClassStruct, err = gi.StructNew("Gio", "TlsCertificateClass")
	})
	return err
}

type TlsCertificateClass struct {
	native uintptr
}

var tlsCertificatePrivateStruct *gi.Struct
var tlsCertificatePrivateStruct_Once sync.Once

func tlsCertificatePrivateStruct_Set() error {
	var err error
	tlsCertificatePrivateStruct_Once.Do(func() {
		tlsCertificatePrivateStruct, err = gi.StructNew("Gio", "TlsCertificatePrivate")
	})
	return err
}

type TlsCertificatePrivate struct {
	native uintptr
}

var tlsClientConnectionInterfaceStruct *gi.Struct
var tlsClientConnectionInterfaceStruct_Once sync.Once

func tlsClientConnectionInterfaceStruct_Set() error {
	var err error
	tlsClientConnectionInterfaceStruct_Once.Do(func() {
		tlsClientConnectionInterfaceStruct, err = gi.StructNew("Gio", "TlsClientConnectionInterface")
	})
	return err
}

type TlsClientConnectionInterface struct {
	native uintptr
}

var tlsConnectionClassStruct *gi.Struct
var tlsConnectionClassStruct_Once sync.Once

func tlsConnectionClassStruct_Set() error {
	var err error
	tlsConnectionClassStruct_Once.Do(func() {
		tlsConnectionClassStruct, err = gi.StructNew("Gio", "TlsConnectionClass")
	})
	return err
}

type TlsConnectionClass struct {
	native uintptr
}

var tlsConnectionPrivateStruct *gi.Struct
var tlsConnectionPrivateStruct_Once sync.Once

func tlsConnectionPrivateStruct_Set() error {
	var err error
	tlsConnectionPrivateStruct_Once.Do(func() {
		tlsConnectionPrivateStruct, err = gi.StructNew("Gio", "TlsConnectionPrivate")
	})
	return err
}

type TlsConnectionPrivate struct {
	native uintptr
}

var tlsDatabaseClassStruct *gi.Struct
var tlsDatabaseClassStruct_Once sync.Once

func tlsDatabaseClassStruct_Set() error {
	var err error
	tlsDatabaseClassStruct_Once.Do(func() {
		tlsDatabaseClassStruct, err = gi.StructNew("Gio", "TlsDatabaseClass")
	})
	return err
}

type TlsDatabaseClass struct {
	native uintptr
}

var tlsDatabasePrivateStruct *gi.Struct
var tlsDatabasePrivateStruct_Once sync.Once

func tlsDatabasePrivateStruct_Set() error {
	var err error
	tlsDatabasePrivateStruct_Once.Do(func() {
		tlsDatabasePrivateStruct, err = gi.StructNew("Gio", "TlsDatabasePrivate")
	})
	return err
}

type TlsDatabasePrivate struct {
	native uintptr
}

var tlsFileDatabaseInterfaceStruct *gi.Struct
var tlsFileDatabaseInterfaceStruct_Once sync.Once

func tlsFileDatabaseInterfaceStruct_Set() error {
	var err error
	tlsFileDatabaseInterfaceStruct_Once.Do(func() {
		tlsFileDatabaseInterfaceStruct, err = gi.StructNew("Gio", "TlsFileDatabaseInterface")
	})
	return err
}

type TlsFileDatabaseInterface struct {
	native uintptr
}

var tlsInteractionClassStruct *gi.Struct
var tlsInteractionClassStruct_Once sync.Once

func tlsInteractionClassStruct_Set() error {
	var err error
	tlsInteractionClassStruct_Once.Do(func() {
		tlsInteractionClassStruct, err = gi.StructNew("Gio", "TlsInteractionClass")
	})
	return err
}

type TlsInteractionClass struct {
	native uintptr
}

var tlsInteractionPrivateStruct *gi.Struct
var tlsInteractionPrivateStruct_Once sync.Once

func tlsInteractionPrivateStruct_Set() error {
	var err error
	tlsInteractionPrivateStruct_Once.Do(func() {
		tlsInteractionPrivateStruct, err = gi.StructNew("Gio", "TlsInteractionPrivate")
	})
	return err
}

type TlsInteractionPrivate struct {
	native uintptr
}

var tlsPasswordClassStruct *gi.Struct
var tlsPasswordClassStruct_Once sync.Once

func tlsPasswordClassStruct_Set() error {
	var err error
	tlsPasswordClassStruct_Once.Do(func() {
		tlsPasswordClassStruct, err = gi.StructNew("Gio", "TlsPasswordClass")
	})
	return err
}

type TlsPasswordClass struct {
	native uintptr
}

var tlsPasswordPrivateStruct *gi.Struct
var tlsPasswordPrivateStruct_Once sync.Once

func tlsPasswordPrivateStruct_Set() error {
	var err error
	tlsPasswordPrivateStruct_Once.Do(func() {
		tlsPasswordPrivateStruct, err = gi.StructNew("Gio", "TlsPasswordPrivate")
	})
	return err
}

type TlsPasswordPrivate struct {
	native uintptr
}

var tlsServerConnectionInterfaceStruct *gi.Struct
var tlsServerConnectionInterfaceStruct_Once sync.Once

func tlsServerConnectionInterfaceStruct_Set() error {
	var err error
	tlsServerConnectionInterfaceStruct_Once.Do(func() {
		tlsServerConnectionInterfaceStruct, err = gi.StructNew("Gio", "TlsServerConnectionInterface")
	})
	return err
}

type TlsServerConnectionInterface struct {
	native uintptr
}

var unixConnectionClassStruct *gi.Struct
var unixConnectionClassStruct_Once sync.Once

func unixConnectionClassStruct_Set() error {
	var err error
	unixConnectionClassStruct_Once.Do(func() {
		unixConnectionClassStruct, err = gi.StructNew("Gio", "UnixConnectionClass")
	})
	return err
}

type UnixConnectionClass struct {
	native uintptr
}

var unixConnectionPrivateStruct *gi.Struct
var unixConnectionPrivateStruct_Once sync.Once

func unixConnectionPrivateStruct_Set() error {
	var err error
	unixConnectionPrivateStruct_Once.Do(func() {
		unixConnectionPrivateStruct, err = gi.StructNew("Gio", "UnixConnectionPrivate")
	})
	return err
}

type UnixConnectionPrivate struct {
	native uintptr
}

var unixCredentialsMessageClassStruct *gi.Struct
var unixCredentialsMessageClassStruct_Once sync.Once

func unixCredentialsMessageClassStruct_Set() error {
	var err error
	unixCredentialsMessageClassStruct_Once.Do(func() {
		unixCredentialsMessageClassStruct, err = gi.StructNew("Gio", "UnixCredentialsMessageClass")
	})
	return err
}

type UnixCredentialsMessageClass struct {
	native uintptr
}

var unixCredentialsMessagePrivateStruct *gi.Struct
var unixCredentialsMessagePrivateStruct_Once sync.Once

func unixCredentialsMessagePrivateStruct_Set() error {
	var err error
	unixCredentialsMessagePrivateStruct_Once.Do(func() {
		unixCredentialsMessagePrivateStruct, err = gi.StructNew("Gio", "UnixCredentialsMessagePrivate")
	})
	return err
}

type UnixCredentialsMessagePrivate struct {
	native uintptr
}

var unixFDListClassStruct *gi.Struct
var unixFDListClassStruct_Once sync.Once

func unixFDListClassStruct_Set() error {
	var err error
	unixFDListClassStruct_Once.Do(func() {
		unixFDListClassStruct, err = gi.StructNew("Gio", "UnixFDListClass")
	})
	return err
}

type UnixFDListClass struct {
	native uintptr
}

var unixFDListPrivateStruct *gi.Struct
var unixFDListPrivateStruct_Once sync.Once

func unixFDListPrivateStruct_Set() error {
	var err error
	unixFDListPrivateStruct_Once.Do(func() {
		unixFDListPrivateStruct, err = gi.StructNew("Gio", "UnixFDListPrivate")
	})
	return err
}

type UnixFDListPrivate struct {
	native uintptr
}

var unixFDMessageClassStruct *gi.Struct
var unixFDMessageClassStruct_Once sync.Once

func unixFDMessageClassStruct_Set() error {
	var err error
	unixFDMessageClassStruct_Once.Do(func() {
		unixFDMessageClassStruct, err = gi.StructNew("Gio", "UnixFDMessageClass")
	})
	return err
}

type UnixFDMessageClass struct {
	native uintptr
}

var unixFDMessagePrivateStruct *gi.Struct
var unixFDMessagePrivateStruct_Once sync.Once

func unixFDMessagePrivateStruct_Set() error {
	var err error
	unixFDMessagePrivateStruct_Once.Do(func() {
		unixFDMessagePrivateStruct, err = gi.StructNew("Gio", "UnixFDMessagePrivate")
	})
	return err
}

type UnixFDMessagePrivate struct {
	native uintptr
}

var unixInputStreamClassStruct *gi.Struct
var unixInputStreamClassStruct_Once sync.Once

func unixInputStreamClassStruct_Set() error {
	var err error
	unixInputStreamClassStruct_Once.Do(func() {
		unixInputStreamClassStruct, err = gi.StructNew("Gio", "UnixInputStreamClass")
	})
	return err
}

type UnixInputStreamClass struct {
	native uintptr
}

var unixInputStreamPrivateStruct *gi.Struct
var unixInputStreamPrivateStruct_Once sync.Once

func unixInputStreamPrivateStruct_Set() error {
	var err error
	unixInputStreamPrivateStruct_Once.Do(func() {
		unixInputStreamPrivateStruct, err = gi.StructNew("Gio", "UnixInputStreamPrivate")
	})
	return err
}

type UnixInputStreamPrivate struct {
	native uintptr
}

var unixMountEntryStruct *gi.Struct
var unixMountEntryStruct_Once sync.Once

func unixMountEntryStruct_Set() error {
	var err error
	unixMountEntryStruct_Once.Do(func() {
		unixMountEntryStruct, err = gi.StructNew("Gio", "UnixMountEntry")
	})
	return err
}

type UnixMountEntry struct {
	native uintptr
}

var unixMountMonitorClassStruct *gi.Struct
var unixMountMonitorClassStruct_Once sync.Once

func unixMountMonitorClassStruct_Set() error {
	var err error
	unixMountMonitorClassStruct_Once.Do(func() {
		unixMountMonitorClassStruct, err = gi.StructNew("Gio", "UnixMountMonitorClass")
	})
	return err
}

type UnixMountMonitorClass struct {
	native uintptr
}

var unixMountPointStruct *gi.Struct
var unixMountPointStruct_Once sync.Once

func unixMountPointStruct_Set() error {
	var err error
	unixMountPointStruct_Once.Do(func() {
		unixMountPointStruct, err = gi.StructNew("Gio", "UnixMountPoint")
	})
	return err
}

type UnixMountPoint struct {
	native uintptr
}

var unixMountPointCompareFunction *gi.Function
var unixMountPointCompareFunction_Once sync.Once

func unixMountPointCompareFunction_Set() error {
	var err error
	unixMountPointCompareFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointCompareFunction, err = unixMountPointStruct.InvokerNew("compare")
	})
	return err
}

// Compare is a representation of the C type g_unix_mount_point_compare.
func (recv *UnixMountPoint) Compare(mount2 *UnixMountPoint) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(mount2.native)

	var ret gi.Argument

	err := unixMountPointCompareFunction_Set()
	if err == nil {
		ret = unixMountPointCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var unixMountPointCopyFunction *gi.Function
var unixMountPointCopyFunction_Once sync.Once

func unixMountPointCopyFunction_Set() error {
	var err error
	unixMountPointCopyFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointCopyFunction, err = unixMountPointStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_unix_mount_point_copy.
func (recv *UnixMountPoint) Copy() (*UnixMountPoint, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointCopyFunction_Set()
	if err == nil {
		ret = unixMountPointCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &UnixMountPoint{native: ret.Pointer()}

	return retGo, err
}

var unixMountPointFreeFunction *gi.Function
var unixMountPointFreeFunction_Once sync.Once

func unixMountPointFreeFunction_Set() error {
	var err error
	unixMountPointFreeFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointFreeFunction, err = unixMountPointStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_unix_mount_point_free.
func (recv *UnixMountPoint) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := unixMountPointFreeFunction_Set()
	if err == nil {
		unixMountPointFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_unix_mount_point_get_device_path' : return type 'filename' not supported

var unixMountPointGetFsTypeFunction *gi.Function
var unixMountPointGetFsTypeFunction_Once sync.Once

func unixMountPointGetFsTypeFunction_Set() error {
	var err error
	unixMountPointGetFsTypeFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetFsTypeFunction, err = unixMountPointStruct.InvokerNew("get_fs_type")
	})
	return err
}

// GetFsType is a representation of the C type g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetFsTypeFunction_Set()
	if err == nil {
		ret = unixMountPointGetFsTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_mount_point_get_mount_path' : return type 'filename' not supported

var unixMountPointGetOptionsFunction *gi.Function
var unixMountPointGetOptionsFunction_Once sync.Once

func unixMountPointGetOptionsFunction_Set() error {
	var err error
	unixMountPointGetOptionsFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGetOptionsFunction, err = unixMountPointStruct.InvokerNew("get_options")
	})
	return err
}

// GetOptions is a representation of the C type g_unix_mount_point_get_options.
func (recv *UnixMountPoint) GetOptions() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGetOptionsFunction_Set()
	if err == nil {
		ret = unixMountPointGetOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var unixMountPointGuessCanEjectFunction *gi.Function
var unixMountPointGuessCanEjectFunction_Once sync.Once

func unixMountPointGuessCanEjectFunction_Set() error {
	var err error
	unixMountPointGuessCanEjectFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGuessCanEjectFunction, err = unixMountPointStruct.InvokerNew("guess_can_eject")
	})
	return err
}

// GuessCanEject is a representation of the C type g_unix_mount_point_guess_can_eject.
func (recv *UnixMountPoint) GuessCanEject() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGuessCanEjectFunction_Set()
	if err == nil {
		ret = unixMountPointGuessCanEjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_icon' : return type 'Icon' not supported

var unixMountPointGuessNameFunction *gi.Function
var unixMountPointGuessNameFunction_Once sync.Once

func unixMountPointGuessNameFunction_Set() error {
	var err error
	unixMountPointGuessNameFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointGuessNameFunction, err = unixMountPointStruct.InvokerNew("guess_name")
	})
	return err
}

// GuessName is a representation of the C type g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointGuessNameFunction_Set()
	if err == nil {
		ret = unixMountPointGuessNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_unix_mount_point_guess_symbolic_icon' : return type 'Icon' not supported

var unixMountPointIsLoopbackFunction *gi.Function
var unixMountPointIsLoopbackFunction_Once sync.Once

func unixMountPointIsLoopbackFunction_Set() error {
	var err error
	unixMountPointIsLoopbackFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsLoopbackFunction, err = unixMountPointStruct.InvokerNew("is_loopback")
	})
	return err
}

// IsLoopback is a representation of the C type g_unix_mount_point_is_loopback.
func (recv *UnixMountPoint) IsLoopback() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsLoopbackFunction_Set()
	if err == nil {
		ret = unixMountPointIsLoopbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var unixMountPointIsReadonlyFunction *gi.Function
var unixMountPointIsReadonlyFunction_Once sync.Once

func unixMountPointIsReadonlyFunction_Set() error {
	var err error
	unixMountPointIsReadonlyFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsReadonlyFunction, err = unixMountPointStruct.InvokerNew("is_readonly")
	})
	return err
}

// IsReadonly is a representation of the C type g_unix_mount_point_is_readonly.
func (recv *UnixMountPoint) IsReadonly() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsReadonlyFunction_Set()
	if err == nil {
		ret = unixMountPointIsReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var unixMountPointIsUserMountableFunction *gi.Function
var unixMountPointIsUserMountableFunction_Once sync.Once

func unixMountPointIsUserMountableFunction_Set() error {
	var err error
	unixMountPointIsUserMountableFunction_Once.Do(func() {
		err = unixMountPointStruct_Set()
		if err != nil {
			return
		}
		unixMountPointIsUserMountableFunction, err = unixMountPointStruct.InvokerNew("is_user_mountable")
	})
	return err
}

// IsUserMountable is a representation of the C type g_unix_mount_point_is_user_mountable.
func (recv *UnixMountPoint) IsUserMountable() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := unixMountPointIsUserMountableFunction_Set()
	if err == nil {
		ret = unixMountPointIsUserMountableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var unixOutputStreamClassStruct *gi.Struct
var unixOutputStreamClassStruct_Once sync.Once

func unixOutputStreamClassStruct_Set() error {
	var err error
	unixOutputStreamClassStruct_Once.Do(func() {
		unixOutputStreamClassStruct, err = gi.StructNew("Gio", "UnixOutputStreamClass")
	})
	return err
}

type UnixOutputStreamClass struct {
	native uintptr
}

var unixOutputStreamPrivateStruct *gi.Struct
var unixOutputStreamPrivateStruct_Once sync.Once

func unixOutputStreamPrivateStruct_Set() error {
	var err error
	unixOutputStreamPrivateStruct_Once.Do(func() {
		unixOutputStreamPrivateStruct, err = gi.StructNew("Gio", "UnixOutputStreamPrivate")
	})
	return err
}

type UnixOutputStreamPrivate struct {
	native uintptr
}

var unixSocketAddressClassStruct *gi.Struct
var unixSocketAddressClassStruct_Once sync.Once

func unixSocketAddressClassStruct_Set() error {
	var err error
	unixSocketAddressClassStruct_Once.Do(func() {
		unixSocketAddressClassStruct, err = gi.StructNew("Gio", "UnixSocketAddressClass")
	})
	return err
}

type UnixSocketAddressClass struct {
	native uintptr
}

var unixSocketAddressPrivateStruct *gi.Struct
var unixSocketAddressPrivateStruct_Once sync.Once

func unixSocketAddressPrivateStruct_Set() error {
	var err error
	unixSocketAddressPrivateStruct_Once.Do(func() {
		unixSocketAddressPrivateStruct, err = gi.StructNew("Gio", "UnixSocketAddressPrivate")
	})
	return err
}

type UnixSocketAddressPrivate struct {
	native uintptr
}

var vfsClassStruct *gi.Struct
var vfsClassStruct_Once sync.Once

func vfsClassStruct_Set() error {
	var err error
	vfsClassStruct_Once.Do(func() {
		vfsClassStruct, err = gi.StructNew("Gio", "VfsClass")
	})
	return err
}

type VfsClass struct {
	native uintptr
}

var volumeIfaceStruct *gi.Struct
var volumeIfaceStruct_Once sync.Once

func volumeIfaceStruct_Set() error {
	var err error
	volumeIfaceStruct_Once.Do(func() {
		volumeIfaceStruct, err = gi.StructNew("Gio", "VolumeIface")
	})
	return err
}

type VolumeIface struct {
	native uintptr
}

var volumeMonitorClassStruct *gi.Struct
var volumeMonitorClassStruct_Once sync.Once

func volumeMonitorClassStruct_Set() error {
	var err error
	volumeMonitorClassStruct_Once.Do(func() {
		volumeMonitorClassStruct, err = gi.StructNew("Gio", "VolumeMonitorClass")
	})
	return err
}

type VolumeMonitorClass struct {
	native uintptr
}

var zlibCompressorClassStruct *gi.Struct
var zlibCompressorClassStruct_Once sync.Once

func zlibCompressorClassStruct_Set() error {
	var err error
	zlibCompressorClassStruct_Once.Do(func() {
		zlibCompressorClassStruct, err = gi.StructNew("Gio", "ZlibCompressorClass")
	})
	return err
}

type ZlibCompressorClass struct {
	native uintptr
}

var zlibDecompressorClassStruct *gi.Struct
var zlibDecompressorClassStruct_Once sync.Once

func zlibDecompressorClassStruct_Set() error {
	var err error
	zlibDecompressorClassStruct_Once.Do(func() {
		zlibDecompressorClassStruct, err = gi.StructNew("Gio", "ZlibDecompressorClass")
	})
	return err
}

type ZlibDecompressorClass struct {
	native uintptr
}
