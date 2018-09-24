// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// ActionEntry is a wrapper around the C record GActionEntry.
type ActionEntry struct {
	native *C.GActionEntry
	Name   string
	// no type for activate
	ParameterType string
	State         string
	// no type for change_state
	// Private : padding
}

func ActionEntryNewFromC(u unsafe.Pointer) *ActionEntry {
	c := (*C.GActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &ActionEntry{
		Name:          C.GoString(c.name),
		ParameterType: C.GoString(c.parameter_type),
		State:         C.GoString(c.state),
		native:        c,
	}

	return g
}

func (recv *ActionEntry) toC() *C.GActionEntry {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.parameter_type =
		C.CString(recv.ParameterType)
	recv.native.state =
		C.CString(recv.State)

	return recv.native
}

// AppInfoIface is a wrapper around the C record GAppInfoIface.
type AppInfoIface struct {
	native *C.GAppInfoIface
	// g_iface : record
	// no type for dup
	// no type for equal
	// no type for get_id
	// no type for get_name
	// no type for get_description
	// no type for get_executable
	// no type for get_icon
	// no type for launch
	// no type for supports_uris
	// no type for supports_files
	// no type for launch_uris
	// no type for should_show
	// no type for set_as_default_for_type
	// no type for set_as_default_for_extension
	// no type for add_supports_type
	// no type for can_remove_supports_type
	// no type for remove_supports_type
	// no type for can_delete
	// no type for do_delete
	// no type for get_commandline
	// no type for get_display_name
	// no type for set_as_last_used_for_type
	// no type for get_supported_types
}

func AppInfoIfaceNewFromC(u unsafe.Pointer) *AppInfoIface {
	c := (*C.GAppInfoIface)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoIface{native: c}

	return g
}

func (recv *AppInfoIface) toC() *C.GAppInfoIface {

	return recv.native
}

// AppLaunchContextClass is a wrapper around the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native *C.GAppLaunchContextClass
	// parent_class : record
	// no type for get_display
	// no type for get_startup_notify_id
	// no type for launch_failed
	// no type for launched
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func AppLaunchContextClassNewFromC(u unsafe.Pointer) *AppLaunchContextClass {
	c := (*C.GAppLaunchContextClass)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextClass{native: c}

	return g
}

func (recv *AppLaunchContextClass) toC() *C.GAppLaunchContextClass {

	return recv.native
}

// AppLaunchContextPrivate is a wrapper around the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native *C.GAppLaunchContextPrivate
}

func AppLaunchContextPrivateNewFromC(u unsafe.Pointer) *AppLaunchContextPrivate {
	c := (*C.GAppLaunchContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextPrivate{native: c}

	return g
}

func (recv *AppLaunchContextPrivate) toC() *C.GAppLaunchContextPrivate {

	return recv.native
}

// ApplicationCommandLinePrivate is a wrapper around the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native *C.GApplicationCommandLinePrivate
}

func ApplicationCommandLinePrivateNewFromC(u unsafe.Pointer) *ApplicationCommandLinePrivate {
	c := (*C.GApplicationCommandLinePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLinePrivate{native: c}

	return g
}

func (recv *ApplicationCommandLinePrivate) toC() *C.GApplicationCommandLinePrivate {

	return recv.native
}

// ApplicationPrivate is a wrapper around the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native *C.GApplicationPrivate
}

func ApplicationPrivateNewFromC(u unsafe.Pointer) *ApplicationPrivate {
	c := (*C.GApplicationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationPrivate{native: c}

	return g
}

func (recv *ApplicationPrivate) toC() *C.GApplicationPrivate {

	return recv.native
}

// AsyncResultIface is a wrapper around the C record GAsyncResultIface.
type AsyncResultIface struct {
	native *C.GAsyncResultIface
	// g_iface : record
	// no type for get_user_data
	// no type for get_source_object
	// no type for is_tagged
}

func AsyncResultIfaceNewFromC(u unsafe.Pointer) *AsyncResultIface {
	c := (*C.GAsyncResultIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResultIface{native: c}

	return g
}

func (recv *AsyncResultIface) toC() *C.GAsyncResultIface {

	return recv.native
}

// BufferedInputStreamClass is a wrapper around the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native *C.GBufferedInputStreamClass
	// parent_class : record
	// no type for fill
	// no type for fill_async
	// no type for fill_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func BufferedInputStreamClassNewFromC(u unsafe.Pointer) *BufferedInputStreamClass {
	c := (*C.GBufferedInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamClass{native: c}

	return g
}

func (recv *BufferedInputStreamClass) toC() *C.GBufferedInputStreamClass {

	return recv.native
}

// BufferedInputStreamPrivate is a wrapper around the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native *C.GBufferedInputStreamPrivate
}

func BufferedInputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedInputStreamPrivate {
	c := (*C.GBufferedInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamPrivate{native: c}

	return g
}

func (recv *BufferedInputStreamPrivate) toC() *C.GBufferedInputStreamPrivate {

	return recv.native
}

// BufferedOutputStreamClass is a wrapper around the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native *C.GBufferedOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func BufferedOutputStreamClassNewFromC(u unsafe.Pointer) *BufferedOutputStreamClass {
	c := (*C.GBufferedOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamClass{native: c}

	return g
}

func (recv *BufferedOutputStreamClass) toC() *C.GBufferedOutputStreamClass {

	return recv.native
}

// BufferedOutputStreamPrivate is a wrapper around the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native *C.GBufferedOutputStreamPrivate
}

func BufferedOutputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedOutputStreamPrivate {
	c := (*C.GBufferedOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamPrivate{native: c}

	return g
}

func (recv *BufferedOutputStreamPrivate) toC() *C.GBufferedOutputStreamPrivate {

	return recv.native
}

// CancellableClass is a wrapper around the C record GCancellableClass.
type CancellableClass struct {
	native *C.GCancellableClass
	// parent_class : record
	// no type for cancelled
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func CancellableClassNewFromC(u unsafe.Pointer) *CancellableClass {
	c := (*C.GCancellableClass)(u)
	if c == nil {
		return nil
	}

	g := &CancellableClass{native: c}

	return g
}

func (recv *CancellableClass) toC() *C.GCancellableClass {

	return recv.native
}

// CancellablePrivate is a wrapper around the C record GCancellablePrivate.
type CancellablePrivate struct {
	native *C.GCancellablePrivate
}

func CancellablePrivateNewFromC(u unsafe.Pointer) *CancellablePrivate {
	c := (*C.GCancellablePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CancellablePrivate{native: c}

	return g
}

func (recv *CancellablePrivate) toC() *C.GCancellablePrivate {

	return recv.native
}

// CharsetConverterClass is a wrapper around the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native *C.GCharsetConverterClass
	// parent_class : record
}

func CharsetConverterClassNewFromC(u unsafe.Pointer) *CharsetConverterClass {
	c := (*C.GCharsetConverterClass)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverterClass{native: c}

	return g
}

func (recv *CharsetConverterClass) toC() *C.GCharsetConverterClass {

	return recv.native
}

// ConverterInputStreamClass is a wrapper around the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native *C.GConverterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterInputStreamClassNewFromC(u unsafe.Pointer) *ConverterInputStreamClass {
	c := (*C.GConverterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamClass{native: c}

	return g
}

func (recv *ConverterInputStreamClass) toC() *C.GConverterInputStreamClass {

	return recv.native
}

// ConverterInputStreamPrivate is a wrapper around the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native *C.GConverterInputStreamPrivate
}

func ConverterInputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterInputStreamPrivate {
	c := (*C.GConverterInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamPrivate{native: c}

	return g
}

func (recv *ConverterInputStreamPrivate) toC() *C.GConverterInputStreamPrivate {

	return recv.native
}

// ConverterOutputStreamClass is a wrapper around the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native *C.GConverterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterOutputStreamClassNewFromC(u unsafe.Pointer) *ConverterOutputStreamClass {
	c := (*C.GConverterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamClass{native: c}

	return g
}

func (recv *ConverterOutputStreamClass) toC() *C.GConverterOutputStreamClass {

	return recv.native
}

// ConverterOutputStreamPrivate is a wrapper around the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native *C.GConverterOutputStreamPrivate
}

func ConverterOutputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterOutputStreamPrivate {
	c := (*C.GConverterOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamPrivate{native: c}

	return g
}

func (recv *ConverterOutputStreamPrivate) toC() *C.GConverterOutputStreamPrivate {

	return recv.native
}

// DBusInterfaceSkeletonPrivate is a wrapper around the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native *C.GDBusInterfaceSkeletonPrivate
}

func DBusInterfaceSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	c := (*C.GDBusInterfaceSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeletonPrivate{native: c}

	return g
}

func (recv *DBusInterfaceSkeletonPrivate) toC() *C.GDBusInterfaceSkeletonPrivate {

	return recv.native
}

// DBusObjectManagerClientPrivate is a wrapper around the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native *C.GDBusObjectManagerClientPrivate
}

func DBusObjectManagerClientPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerClientPrivate {
	c := (*C.GDBusObjectManagerClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClientPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerClientPrivate) toC() *C.GDBusObjectManagerClientPrivate {

	return recv.native
}

// DBusObjectManagerServerPrivate is a wrapper around the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native *C.GDBusObjectManagerServerPrivate
}

func DBusObjectManagerServerPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerServerPrivate {
	c := (*C.GDBusObjectManagerServerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServerPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerServerPrivate) toC() *C.GDBusObjectManagerServerPrivate {

	return recv.native
}

// DBusObjectProxyPrivate is a wrapper around the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native *C.GDBusObjectProxyPrivate
}

func DBusObjectProxyPrivateNewFromC(u unsafe.Pointer) *DBusObjectProxyPrivate {
	c := (*C.GDBusObjectProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxyPrivate{native: c}

	return g
}

func (recv *DBusObjectProxyPrivate) toC() *C.GDBusObjectProxyPrivate {

	return recv.native
}

// DBusObjectSkeletonPrivate is a wrapper around the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native *C.GDBusObjectSkeletonPrivate
}

func DBusObjectSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusObjectSkeletonPrivate {
	c := (*C.GDBusObjectSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeletonPrivate{native: c}

	return g
}

func (recv *DBusObjectSkeletonPrivate) toC() *C.GDBusObjectSkeletonPrivate {

	return recv.native
}

// DBusProxyPrivate is a wrapper around the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native *C.GDBusProxyPrivate
}

func DBusProxyPrivateNewFromC(u unsafe.Pointer) *DBusProxyPrivate {
	c := (*C.GDBusProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxyPrivate{native: c}

	return g
}

func (recv *DBusProxyPrivate) toC() *C.GDBusProxyPrivate {

	return recv.native
}

// DataInputStreamClass is a wrapper around the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native *C.GDataInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataInputStreamClassNewFromC(u unsafe.Pointer) *DataInputStreamClass {
	c := (*C.GDataInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamClass{native: c}

	return g
}

func (recv *DataInputStreamClass) toC() *C.GDataInputStreamClass {

	return recv.native
}

// DataInputStreamPrivate is a wrapper around the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native *C.GDataInputStreamPrivate
}

func DataInputStreamPrivateNewFromC(u unsafe.Pointer) *DataInputStreamPrivate {
	c := (*C.GDataInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamPrivate{native: c}

	return g
}

func (recv *DataInputStreamPrivate) toC() *C.GDataInputStreamPrivate {

	return recv.native
}

// DataOutputStreamClass is a wrapper around the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native *C.GDataOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataOutputStreamClassNewFromC(u unsafe.Pointer) *DataOutputStreamClass {
	c := (*C.GDataOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamClass{native: c}

	return g
}

func (recv *DataOutputStreamClass) toC() *C.GDataOutputStreamClass {

	return recv.native
}

// DataOutputStreamPrivate is a wrapper around the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native *C.GDataOutputStreamPrivate
}

func DataOutputStreamPrivateNewFromC(u unsafe.Pointer) *DataOutputStreamPrivate {
	c := (*C.GDataOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamPrivate{native: c}

	return g
}

func (recv *DataOutputStreamPrivate) toC() *C.GDataOutputStreamPrivate {

	return recv.native
}

// DesktopAppInfoClass is a wrapper around the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native *C.GDesktopAppInfoClass
	// parent_class : record
}

func DesktopAppInfoClassNewFromC(u unsafe.Pointer) *DesktopAppInfoClass {
	c := (*C.GDesktopAppInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoClass{native: c}

	return g
}

func (recv *DesktopAppInfoClass) toC() *C.GDesktopAppInfoClass {

	return recv.native
}

// DesktopAppInfoLookupIface is a wrapper around the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native *C.GDesktopAppInfoLookupIface
	// g_iface : record
	// no type for get_default_for_uri_scheme
}

func DesktopAppInfoLookupIfaceNewFromC(u unsafe.Pointer) *DesktopAppInfoLookupIface {
	c := (*C.GDesktopAppInfoLookupIface)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookupIface{native: c}

	return g
}

func (recv *DesktopAppInfoLookupIface) toC() *C.GDesktopAppInfoLookupIface {

	return recv.native
}

// DriveIface is a wrapper around the C record GDriveIface.
type DriveIface struct {
	native *C.GDriveIface
	// g_iface : record
	// no type for changed
	// no type for disconnected
	// no type for eject_button
	// no type for get_name
	// no type for get_icon
	// no type for has_volumes
	// no type for get_volumes
	// no type for is_media_removable
	// no type for has_media
	// no type for is_media_check_automatic
	// no type for can_eject
	// no type for can_poll_for_media
	// no type for eject
	// no type for eject_finish
	// no type for poll_for_media
	// no type for poll_for_media_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for get_start_stop_type
	// no type for can_start
	// no type for can_start_degraded
	// no type for start
	// no type for start_finish
	// no type for can_stop
	// no type for stop
	// no type for stop_finish
	// no type for stop_button
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
	// no type for is_removable
}

func DriveIfaceNewFromC(u unsafe.Pointer) *DriveIface {
	c := (*C.GDriveIface)(u)
	if c == nil {
		return nil
	}

	g := &DriveIface{native: c}

	return g
}

func (recv *DriveIface) toC() *C.GDriveIface {

	return recv.native
}

// EmblemClass is a wrapper around the C record GEmblemClass.
type EmblemClass struct {
	native *C.GEmblemClass
}

func EmblemClassNewFromC(u unsafe.Pointer) *EmblemClass {
	c := (*C.GEmblemClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemClass{native: c}

	return g
}

func (recv *EmblemClass) toC() *C.GEmblemClass {

	return recv.native
}

// EmblemedIconClass is a wrapper around the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native *C.GEmblemedIconClass
	// parent_class : record
}

func EmblemedIconClassNewFromC(u unsafe.Pointer) *EmblemedIconClass {
	c := (*C.GEmblemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconClass{native: c}

	return g
}

func (recv *EmblemedIconClass) toC() *C.GEmblemedIconClass {

	return recv.native
}

// EmblemedIconPrivate is a wrapper around the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native *C.GEmblemedIconPrivate
}

func EmblemedIconPrivateNewFromC(u unsafe.Pointer) *EmblemedIconPrivate {
	c := (*C.GEmblemedIconPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconPrivate{native: c}

	return g
}

func (recv *EmblemedIconPrivate) toC() *C.GEmblemedIconPrivate {

	return recv.native
}

// FileAttributeInfo is a wrapper around the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native *C.GFileAttributeInfo
	Name   string
	Type   FileAttributeType
	Flags  FileAttributeInfoFlags
}

func FileAttributeInfoNewFromC(u unsafe.Pointer) *FileAttributeInfo {
	c := (*C.GFileAttributeInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfo{
		Flags:  (FileAttributeInfoFlags)(c.flags),
		Name:   C.GoString(c.name),
		Type:   (FileAttributeType)(c._type),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfo) toC() *C.GFileAttributeInfo {
	recv.native.name =
		C.CString(recv.Name)
	recv.native._type =
		(C.GFileAttributeType)(recv.Type)
	recv.native.flags =
		(C.GFileAttributeInfoFlags)(recv.Flags)

	return recv.native
}

// FileAttributeInfoList is a wrapper around the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native *C.GFileAttributeInfoList
	// infos : record
	NInfos int32
}

func FileAttributeInfoListNewFromC(u unsafe.Pointer) *FileAttributeInfoList {
	c := (*C.GFileAttributeInfoList)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfoList{
		NInfos: (int32)(c.n_infos),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfoList) toC() *C.GFileAttributeInfoList {
	recv.native.n_infos =
		(C.int)(recv.NInfos)

	return recv.native
}

// FileAttributeInfoListNew is a wrapper around the C function g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_new()
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_attribute_info_list_add : no return generator

// Dup is a wrapper around the C function g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_dup((*C.GFileAttributeInfoList)(recv.native))
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lookup is a wrapper around the C function g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_file_attribute_info_list_lookup((*C.GFileAttributeInfoList)(recv.native), c_name)
	retGo := FileAttributeInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_ref((*C.GFileAttributeInfoList)(recv.native))
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_attribute_info_list_unref : no return generator

// FileAttributeMatcher is a wrapper around the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native *C.GFileAttributeMatcher
}

func FileAttributeMatcherNewFromC(u unsafe.Pointer) *FileAttributeMatcher {
	c := (*C.GFileAttributeMatcher)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeMatcher{native: c}

	return g
}

func (recv *FileAttributeMatcher) toC() *C.GFileAttributeMatcher {

	return recv.native
}

// FileAttributeMatcherNew is a wrapper around the C function g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	retC := C.g_file_attribute_matcher_new(c_attributes)
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumerateNamespace is a wrapper around the C function g_file_attribute_matcher_enumerate_namespace.
func (recv *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	c_ns := C.CString(ns)
	defer C.free(unsafe.Pointer(c_ns))

	retC := C.g_file_attribute_matcher_enumerate_namespace((*C.GFileAttributeMatcher)(recv.native), c_ns)
	retGo := retC == C.TRUE

	return retGo
}

// EnumerateNext is a wrapper around the C function g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() string {
	retC := C.g_file_attribute_matcher_enumerate_next((*C.GFileAttributeMatcher)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Matches is a wrapper around the C function g_file_attribute_matcher_matches.
func (recv *FileAttributeMatcher) Matches(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_attribute_matcher_matches((*C.GFileAttributeMatcher)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// MatchesOnly is a wrapper around the C function g_file_attribute_matcher_matches_only.
func (recv *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_attribute_matcher_matches_only((*C.GFileAttributeMatcher)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	retC := C.g_file_attribute_matcher_ref((*C.GFileAttributeMatcher)(recv.native))
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Subtract is a wrapper around the C function g_file_attribute_matcher_subtract.
func (recv *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) *FileAttributeMatcher {
	c_subtract := subtract.toC()

	retC := C.g_file_attribute_matcher_subtract((*C.GFileAttributeMatcher)(recv.native), c_subtract)
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_attribute_matcher_unref : no return generator

// FileDescriptorBasedIface is a wrapper around the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native *C.GFileDescriptorBasedIface
	// g_iface : record
	// no type for get_fd
}

func FileDescriptorBasedIfaceNewFromC(u unsafe.Pointer) *FileDescriptorBasedIface {
	c := (*C.GFileDescriptorBasedIface)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBasedIface{native: c}

	return g
}

func (recv *FileDescriptorBasedIface) toC() *C.GFileDescriptorBasedIface {

	return recv.native
}

// FileEnumeratorClass is a wrapper around the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native *C.GFileEnumeratorClass
	// parent_class : record
	// no type for next_file
	// no type for close_fn
	// no type for next_files_async
	// no type for next_files_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func FileEnumeratorClassNewFromC(u unsafe.Pointer) *FileEnumeratorClass {
	c := (*C.GFileEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorClass{native: c}

	return g
}

func (recv *FileEnumeratorClass) toC() *C.GFileEnumeratorClass {

	return recv.native
}

// FileEnumeratorPrivate is a wrapper around the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native *C.GFileEnumeratorPrivate
}

func FileEnumeratorPrivateNewFromC(u unsafe.Pointer) *FileEnumeratorPrivate {
	c := (*C.GFileEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorPrivate{native: c}

	return g
}

func (recv *FileEnumeratorPrivate) toC() *C.GFileEnumeratorPrivate {

	return recv.native
}

// FileIOStreamClass is a wrapper around the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native *C.GFileIOStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileIOStreamClassNewFromC(u unsafe.Pointer) *FileIOStreamClass {
	c := (*C.GFileIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamClass{native: c}

	return g
}

func (recv *FileIOStreamClass) toC() *C.GFileIOStreamClass {

	return recv.native
}

// FileIOStreamPrivate is a wrapper around the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native *C.GFileIOStreamPrivate
}

func FileIOStreamPrivateNewFromC(u unsafe.Pointer) *FileIOStreamPrivate {
	c := (*C.GFileIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamPrivate{native: c}

	return g
}

func (recv *FileIOStreamPrivate) toC() *C.GFileIOStreamPrivate {

	return recv.native
}

// FileIconClass is a wrapper around the C record GFileIconClass.
type FileIconClass struct {
	native *C.GFileIconClass
}

func FileIconClassNewFromC(u unsafe.Pointer) *FileIconClass {
	c := (*C.GFileIconClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIconClass{native: c}

	return g
}

func (recv *FileIconClass) toC() *C.GFileIconClass {

	return recv.native
}

// FileIface is a wrapper around the C record GFileIface.
type FileIface struct {
	native *C.GFileIface
	// g_iface : record
	// no type for dup
	// no type for hash
	// no type for equal
	// no type for is_native
	// no type for has_uri_scheme
	// no type for get_uri_scheme
	// no type for get_basename
	// no type for get_path
	// no type for get_uri
	// no type for get_parse_name
	// no type for get_parent
	// no type for prefix_matches
	// no type for get_relative_path
	// no type for resolve_relative_path
	// no type for get_child_for_display_name
	// no type for enumerate_children
	// no type for enumerate_children_async
	// no type for enumerate_children_finish
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for query_filesystem_info
	// no type for query_filesystem_info_async
	// no type for query_filesystem_info_finish
	// no type for find_enclosing_mount
	// no type for find_enclosing_mount_async
	// no type for find_enclosing_mount_finish
	// no type for set_display_name
	// no type for set_display_name_async
	// no type for set_display_name_finish
	// no type for query_settable_attributes
	// no type for _query_settable_attributes_async
	// no type for _query_settable_attributes_finish
	// no type for query_writable_namespaces
	// no type for _query_writable_namespaces_async
	// no type for _query_writable_namespaces_finish
	// no type for set_attribute
	// no type for set_attributes_from_info
	// no type for set_attributes_async
	// no type for set_attributes_finish
	// no type for read_fn
	// no type for read_async
	// no type for read_finish
	// no type for append_to
	// no type for append_to_async
	// no type for append_to_finish
	// no type for create
	// no type for create_async
	// no type for create_finish
	// no type for replace
	// no type for replace_async
	// no type for replace_finish
	// no type for delete_file
	// no type for delete_file_async
	// no type for delete_file_finish
	// no type for trash
	// no type for trash_async
	// no type for trash_finish
	// no type for make_directory
	// no type for make_directory_async
	// no type for make_directory_finish
	// no type for make_symbolic_link
	// no type for _make_symbolic_link_async
	// no type for _make_symbolic_link_finish
	// no type for copy
	// no type for copy_async
	// no type for copy_finish
	// no type for move
	// no type for _move_async
	// no type for _move_finish
	// no type for mount_mountable
	// no type for mount_mountable_finish
	// no type for unmount_mountable
	// no type for unmount_mountable_finish
	// no type for eject_mountable
	// no type for eject_mountable_finish
	// no type for mount_enclosing_volume
	// no type for mount_enclosing_volume_finish
	// no type for monitor_dir
	// no type for monitor_file
	// no type for open_readwrite
	// no type for open_readwrite_async
	// no type for open_readwrite_finish
	// no type for create_readwrite
	// no type for create_readwrite_async
	// no type for create_readwrite_finish
	// no type for replace_readwrite
	// no type for replace_readwrite_async
	// no type for replace_readwrite_finish
	// no type for start_mountable
	// no type for start_mountable_finish
	// no type for stop_mountable
	// no type for stop_mountable_finish
	SupportsThreadContexts bool
	// no type for unmount_mountable_with_operation
	// no type for unmount_mountable_with_operation_finish
	// no type for eject_mountable_with_operation
	// no type for eject_mountable_with_operation_finish
	// no type for poll_mountable
	// no type for poll_mountable_finish
	// no type for measure_disk_usage
	// no type for measure_disk_usage_async
	// no type for measure_disk_usage_finish
}

func FileIfaceNewFromC(u unsafe.Pointer) *FileIface {
	c := (*C.GFileIface)(u)
	if c == nil {
		return nil
	}

	g := &FileIface{
		SupportsThreadContexts: c.supports_thread_contexts == C.TRUE,
		native:                 c,
	}

	return g
}

func (recv *FileIface) toC() *C.GFileIface {
	recv.native.supports_thread_contexts =
		boolToGboolean(recv.SupportsThreadContexts)

	return recv.native
}

// FileInfoClass is a wrapper around the C record GFileInfoClass.
type FileInfoClass struct {
	native *C.GFileInfoClass
}

func FileInfoClassNewFromC(u unsafe.Pointer) *FileInfoClass {
	c := (*C.GFileInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInfoClass{native: c}

	return g
}

func (recv *FileInfoClass) toC() *C.GFileInfoClass {

	return recv.native
}

// FileInputStreamClass is a wrapper around the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native *C.GFileInputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileInputStreamClassNewFromC(u unsafe.Pointer) *FileInputStreamClass {
	c := (*C.GFileInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamClass{native: c}

	return g
}

func (recv *FileInputStreamClass) toC() *C.GFileInputStreamClass {

	return recv.native
}

// FileInputStreamPrivate is a wrapper around the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native *C.GFileInputStreamPrivate
}

func FileInputStreamPrivateNewFromC(u unsafe.Pointer) *FileInputStreamPrivate {
	c := (*C.GFileInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamPrivate{native: c}

	return g
}

func (recv *FileInputStreamPrivate) toC() *C.GFileInputStreamPrivate {

	return recv.native
}

// FileMonitorClass is a wrapper around the C record GFileMonitorClass.
type FileMonitorClass struct {
	native *C.GFileMonitorClass
	// parent_class : record
	// no type for changed
	// no type for cancel
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileMonitorClassNewFromC(u unsafe.Pointer) *FileMonitorClass {
	c := (*C.GFileMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorClass{native: c}

	return g
}

func (recv *FileMonitorClass) toC() *C.GFileMonitorClass {

	return recv.native
}

// FileMonitorPrivate is a wrapper around the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native *C.GFileMonitorPrivate
}

func FileMonitorPrivateNewFromC(u unsafe.Pointer) *FileMonitorPrivate {
	c := (*C.GFileMonitorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorPrivate{native: c}

	return g
}

func (recv *FileMonitorPrivate) toC() *C.GFileMonitorPrivate {

	return recv.native
}

// FileOutputStreamClass is a wrapper around the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native *C.GFileOutputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileOutputStreamClassNewFromC(u unsafe.Pointer) *FileOutputStreamClass {
	c := (*C.GFileOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamClass{native: c}

	return g
}

func (recv *FileOutputStreamClass) toC() *C.GFileOutputStreamClass {

	return recv.native
}

// FileOutputStreamPrivate is a wrapper around the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native *C.GFileOutputStreamPrivate
}

func FileOutputStreamPrivateNewFromC(u unsafe.Pointer) *FileOutputStreamPrivate {
	c := (*C.GFileOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamPrivate{native: c}

	return g
}

func (recv *FileOutputStreamPrivate) toC() *C.GFileOutputStreamPrivate {

	return recv.native
}

// FilenameCompleterClass is a wrapper around the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native *C.GFilenameCompleterClass
	// parent_class : record
	// no type for got_completion_data
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilenameCompleterClassNewFromC(u unsafe.Pointer) *FilenameCompleterClass {
	c := (*C.GFilenameCompleterClass)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleterClass{native: c}

	return g
}

func (recv *FilenameCompleterClass) toC() *C.GFilenameCompleterClass {

	return recv.native
}

// FilterInputStreamClass is a wrapper around the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native *C.GFilterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterInputStreamClassNewFromC(u unsafe.Pointer) *FilterInputStreamClass {
	c := (*C.GFilterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStreamClass{native: c}

	return g
}

func (recv *FilterInputStreamClass) toC() *C.GFilterInputStreamClass {

	return recv.native
}

// FilterOutputStreamClass is a wrapper around the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native *C.GFilterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterOutputStreamClassNewFromC(u unsafe.Pointer) *FilterOutputStreamClass {
	c := (*C.GFilterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStreamClass{native: c}

	return g
}

func (recv *FilterOutputStreamClass) toC() *C.GFilterOutputStreamClass {

	return recv.native
}

// IOExtension is a wrapper around the C record GIOExtension.
type IOExtension struct {
	native *C.GIOExtension
}

func IOExtensionNewFromC(u unsafe.Pointer) *IOExtension {
	c := (*C.GIOExtension)(u)
	if c == nil {
		return nil
	}

	g := &IOExtension{native: c}

	return g
}

func (recv *IOExtension) toC() *C.GIOExtension {

	return recv.native
}

// GetName is a wrapper around the C function g_io_extension_get_name.
func (recv *IOExtension) GetName() string {
	retC := C.g_io_extension_get_name((*C.GIOExtension)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() int32 {
	retC := C.g_io_extension_get_priority((*C.GIOExtension)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_io_extension_get_type : no return generator

// RefClass is a wrapper around the C function g_io_extension_ref_class.
func (recv *IOExtension) RefClass() *gobject.TypeClass {
	retC := C.g_io_extension_ref_class((*C.GIOExtension)(recv.native))
	retGo := gobject.TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOExtensionPoint is a wrapper around the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native *C.GIOExtensionPoint
}

func IOExtensionPointNewFromC(u unsafe.Pointer) *IOExtensionPoint {
	c := (*C.GIOExtensionPoint)(u)
	if c == nil {
		return nil
	}

	g := &IOExtensionPoint{native: c}

	return g
}

func (recv *IOExtensionPoint) toC() *C.GIOExtensionPoint {

	return recv.native
}

// GetExtensionByName is a wrapper around the C function g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_get_extension_by_name((*C.GIOExtensionPoint)(recv.native), c_name)
	retGo := IOExtensionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExtensions is a wrapper around the C function g_io_extension_point_get_extensions.
func (recv *IOExtensionPoint) GetExtensions() *glib.List {
	retC := C.g_io_extension_point_get_extensions((*C.GIOExtensionPoint)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_io_extension_point_get_required_type : no return generator

// Unsupported : g_io_extension_point_set_required_type : unsupported parameter type : no type generator for GType, GType

// IOModuleClass is a wrapper around the C record GIOModuleClass.
type IOModuleClass struct {
	native *C.GIOModuleClass
}

func IOModuleClassNewFromC(u unsafe.Pointer) *IOModuleClass {
	c := (*C.GIOModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &IOModuleClass{native: c}

	return g
}

func (recv *IOModuleClass) toC() *C.GIOModuleClass {

	return recv.native
}

// IOSchedulerJob is a wrapper around the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native *C.GIOSchedulerJob
}

func IOSchedulerJobNewFromC(u unsafe.Pointer) *IOSchedulerJob {
	c := (*C.GIOSchedulerJob)(u)
	if c == nil {
		return nil
	}

	g := &IOSchedulerJob{native: c}

	return g
}

func (recv *IOSchedulerJob) toC() *C.GIOSchedulerJob {

	return recv.native
}

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// IOStreamAdapter is a wrapper around the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native *C.GIOStreamAdapter
}

func IOStreamAdapterNewFromC(u unsafe.Pointer) *IOStreamAdapter {
	c := (*C.GIOStreamAdapter)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamAdapter{native: c}

	return g
}

func (recv *IOStreamAdapter) toC() *C.GIOStreamAdapter {

	return recv.native
}

// IOStreamClass is a wrapper around the C record GIOStreamClass.
type IOStreamClass struct {
	native *C.GIOStreamClass
	// parent_class : record
	// no type for get_input_stream
	// no type for get_output_stream
	// no type for close_fn
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func IOStreamClassNewFromC(u unsafe.Pointer) *IOStreamClass {
	c := (*C.GIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamClass{native: c}

	return g
}

func (recv *IOStreamClass) toC() *C.GIOStreamClass {

	return recv.native
}

// IOStreamPrivate is a wrapper around the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native *C.GIOStreamPrivate
}

func IOStreamPrivateNewFromC(u unsafe.Pointer) *IOStreamPrivate {
	c := (*C.GIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamPrivate{native: c}

	return g
}

func (recv *IOStreamPrivate) toC() *C.GIOStreamPrivate {

	return recv.native
}

// IconIface is a wrapper around the C record GIconIface.
type IconIface struct {
	native *C.GIconIface
	// g_iface : record
	// no type for hash
	// no type for equal
	// no type for to_tokens
	// no type for from_tokens
	// no type for serialize
}

func IconIfaceNewFromC(u unsafe.Pointer) *IconIface {
	c := (*C.GIconIface)(u)
	if c == nil {
		return nil
	}

	g := &IconIface{native: c}

	return g
}

func (recv *IconIface) toC() *C.GIconIface {

	return recv.native
}

// InetAddressClass is a wrapper around the C record GInetAddressClass.
type InetAddressClass struct {
	native *C.GInetAddressClass
	// parent_class : record
	// no type for to_string
	// no type for to_bytes
}

func InetAddressClassNewFromC(u unsafe.Pointer) *InetAddressClass {
	c := (*C.GInetAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressClass{native: c}

	return g
}

func (recv *InetAddressClass) toC() *C.GInetAddressClass {

	return recv.native
}

// InetAddressMaskClass is a wrapper around the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native *C.GInetAddressMaskClass
	// parent_class : record
}

func InetAddressMaskClassNewFromC(u unsafe.Pointer) *InetAddressMaskClass {
	c := (*C.GInetAddressMaskClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskClass{native: c}

	return g
}

func (recv *InetAddressMaskClass) toC() *C.GInetAddressMaskClass {

	return recv.native
}

// InetAddressMaskPrivate is a wrapper around the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native *C.GInetAddressMaskPrivate
}

func InetAddressMaskPrivateNewFromC(u unsafe.Pointer) *InetAddressMaskPrivate {
	c := (*C.GInetAddressMaskPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskPrivate{native: c}

	return g
}

func (recv *InetAddressMaskPrivate) toC() *C.GInetAddressMaskPrivate {

	return recv.native
}

// InetAddressPrivate is a wrapper around the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native *C.GInetAddressPrivate
}

func InetAddressPrivateNewFromC(u unsafe.Pointer) *InetAddressPrivate {
	c := (*C.GInetAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressPrivate{native: c}

	return g
}

func (recv *InetAddressPrivate) toC() *C.GInetAddressPrivate {

	return recv.native
}

// InetSocketAddressClass is a wrapper around the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native *C.GInetSocketAddressClass
	// parent_class : record
}

func InetSocketAddressClassNewFromC(u unsafe.Pointer) *InetSocketAddressClass {
	c := (*C.GInetSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressClass{native: c}

	return g
}

func (recv *InetSocketAddressClass) toC() *C.GInetSocketAddressClass {

	return recv.native
}

// InetSocketAddressPrivate is a wrapper around the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native *C.GInetSocketAddressPrivate
}

func InetSocketAddressPrivateNewFromC(u unsafe.Pointer) *InetSocketAddressPrivate {
	c := (*C.GInetSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressPrivate{native: c}

	return g
}

func (recv *InetSocketAddressPrivate) toC() *C.GInetSocketAddressPrivate {

	return recv.native
}

// InputStreamClass is a wrapper around the C record GInputStreamClass.
type InputStreamClass struct {
	native *C.GInputStreamClass
	// parent_class : record
	// no type for read_fn
	// no type for skip
	// no type for close_fn
	// no type for read_async
	// no type for read_finish
	// no type for skip_async
	// no type for skip_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func InputStreamClassNewFromC(u unsafe.Pointer) *InputStreamClass {
	c := (*C.GInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamClass{native: c}

	return g
}

func (recv *InputStreamClass) toC() *C.GInputStreamClass {

	return recv.native
}

// InputStreamPrivate is a wrapper around the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native *C.GInputStreamPrivate
}

func InputStreamPrivateNewFromC(u unsafe.Pointer) *InputStreamPrivate {
	c := (*C.GInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamPrivate{native: c}

	return g
}

func (recv *InputStreamPrivate) toC() *C.GInputStreamPrivate {

	return recv.native
}

// ListStoreClass is a wrapper around the C record GListStoreClass.
type ListStoreClass struct {
	native *C.GListStoreClass
	// parent_class : record
}

func ListStoreClassNewFromC(u unsafe.Pointer) *ListStoreClass {
	c := (*C.GListStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &ListStoreClass{native: c}

	return g
}

func (recv *ListStoreClass) toC() *C.GListStoreClass {

	return recv.native
}

// LoadableIconIface is a wrapper around the C record GLoadableIconIface.
type LoadableIconIface struct {
	native *C.GLoadableIconIface
	// g_iface : record
	// no type for load
	// no type for load_async
	// no type for load_finish
}

func LoadableIconIfaceNewFromC(u unsafe.Pointer) *LoadableIconIface {
	c := (*C.GLoadableIconIface)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIconIface{native: c}

	return g
}

func (recv *LoadableIconIface) toC() *C.GLoadableIconIface {

	return recv.native
}

// MemoryInputStreamClass is a wrapper around the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native *C.GMemoryInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryInputStreamClassNewFromC(u unsafe.Pointer) *MemoryInputStreamClass {
	c := (*C.GMemoryInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamClass{native: c}

	return g
}

func (recv *MemoryInputStreamClass) toC() *C.GMemoryInputStreamClass {

	return recv.native
}

// MemoryInputStreamPrivate is a wrapper around the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native *C.GMemoryInputStreamPrivate
}

func MemoryInputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryInputStreamPrivate {
	c := (*C.GMemoryInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamPrivate{native: c}

	return g
}

func (recv *MemoryInputStreamPrivate) toC() *C.GMemoryInputStreamPrivate {

	return recv.native
}

// MemoryOutputStreamClass is a wrapper around the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native *C.GMemoryOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryOutputStreamClassNewFromC(u unsafe.Pointer) *MemoryOutputStreamClass {
	c := (*C.GMemoryOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamClass{native: c}

	return g
}

func (recv *MemoryOutputStreamClass) toC() *C.GMemoryOutputStreamClass {

	return recv.native
}

// MemoryOutputStreamPrivate is a wrapper around the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native *C.GMemoryOutputStreamPrivate
}

func MemoryOutputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryOutputStreamPrivate {
	c := (*C.GMemoryOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamPrivate{native: c}

	return g
}

func (recv *MemoryOutputStreamPrivate) toC() *C.GMemoryOutputStreamPrivate {

	return recv.native
}

// MenuAttributeIterClass is a wrapper around the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native *C.GMenuAttributeIterClass
	// parent_class : record
	// no type for get_next
}

func MenuAttributeIterClassNewFromC(u unsafe.Pointer) *MenuAttributeIterClass {
	c := (*C.GMenuAttributeIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterClass{native: c}

	return g
}

func (recv *MenuAttributeIterClass) toC() *C.GMenuAttributeIterClass {

	return recv.native
}

// MenuAttributeIterPrivate is a wrapper around the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native *C.GMenuAttributeIterPrivate
}

func MenuAttributeIterPrivateNewFromC(u unsafe.Pointer) *MenuAttributeIterPrivate {
	c := (*C.GMenuAttributeIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterPrivate{native: c}

	return g
}

func (recv *MenuAttributeIterPrivate) toC() *C.GMenuAttributeIterPrivate {

	return recv.native
}

// MenuLinkIterClass is a wrapper around the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native *C.GMenuLinkIterClass
	// parent_class : record
	// no type for get_next
}

func MenuLinkIterClassNewFromC(u unsafe.Pointer) *MenuLinkIterClass {
	c := (*C.GMenuLinkIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterClass{native: c}

	return g
}

func (recv *MenuLinkIterClass) toC() *C.GMenuLinkIterClass {

	return recv.native
}

// MenuLinkIterPrivate is a wrapper around the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native *C.GMenuLinkIterPrivate
}

func MenuLinkIterPrivateNewFromC(u unsafe.Pointer) *MenuLinkIterPrivate {
	c := (*C.GMenuLinkIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterPrivate{native: c}

	return g
}

func (recv *MenuLinkIterPrivate) toC() *C.GMenuLinkIterPrivate {

	return recv.native
}

// MenuModelClass is a wrapper around the C record GMenuModelClass.
type MenuModelClass struct {
	native *C.GMenuModelClass
	// parent_class : record
	// no type for is_mutable
	// no type for get_n_items
	// no type for get_item_attributes
	// no type for iterate_item_attributes
	// no type for get_item_attribute_value
	// no type for get_item_links
	// no type for iterate_item_links
	// no type for get_item_link
}

func MenuModelClassNewFromC(u unsafe.Pointer) *MenuModelClass {
	c := (*C.GMenuModelClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelClass{native: c}

	return g
}

func (recv *MenuModelClass) toC() *C.GMenuModelClass {

	return recv.native
}

// MenuModelPrivate is a wrapper around the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native *C.GMenuModelPrivate
}

func MenuModelPrivateNewFromC(u unsafe.Pointer) *MenuModelPrivate {
	c := (*C.GMenuModelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelPrivate{native: c}

	return g
}

func (recv *MenuModelPrivate) toC() *C.GMenuModelPrivate {

	return recv.native
}

// MountIface is a wrapper around the C record GMountIface.
type MountIface struct {
	native *C.GMountIface
	// g_iface : record
	// no type for changed
	// no type for unmounted
	// no type for get_root
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_volume
	// no type for get_drive
	// no type for can_unmount
	// no type for can_eject
	// no type for unmount
	// no type for unmount_finish
	// no type for eject
	// no type for eject_finish
	// no type for remount
	// no type for remount_finish
	// no type for guess_content_type
	// no type for guess_content_type_finish
	// no type for guess_content_type_sync
	// no type for pre_unmount
	// no type for unmount_with_operation
	// no type for unmount_with_operation_finish
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_default_location
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func MountIfaceNewFromC(u unsafe.Pointer) *MountIface {
	c := (*C.GMountIface)(u)
	if c == nil {
		return nil
	}

	g := &MountIface{native: c}

	return g
}

func (recv *MountIface) toC() *C.GMountIface {

	return recv.native
}

// MountOperationClass is a wrapper around the C record GMountOperationClass.
type MountOperationClass struct {
	native *C.GMountOperationClass
	// parent_class : record
	// no type for ask_password
	// no type for ask_question
	// no type for reply
	// no type for aborted
	// no type for show_processes
	// no type for show_unmount_progress
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
}

func MountOperationClassNewFromC(u unsafe.Pointer) *MountOperationClass {
	c := (*C.GMountOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationClass{native: c}

	return g
}

func (recv *MountOperationClass) toC() *C.GMountOperationClass {

	return recv.native
}

// MountOperationPrivate is a wrapper around the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native *C.GMountOperationPrivate
}

func MountOperationPrivateNewFromC(u unsafe.Pointer) *MountOperationPrivate {
	c := (*C.GMountOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationPrivate{native: c}

	return g
}

func (recv *MountOperationPrivate) toC() *C.GMountOperationPrivate {

	return recv.native
}

// NativeSocketAddress is a wrapper around the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native *C.GNativeSocketAddress
}

func NativeSocketAddressNewFromC(u unsafe.Pointer) *NativeSocketAddress {
	c := (*C.GNativeSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &NativeSocketAddress{native: c}

	return g
}

func (recv *NativeSocketAddress) toC() *C.GNativeSocketAddress {

	return recv.native
}

// NativeVolumeMonitorClass is a wrapper around the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native *C.GNativeVolumeMonitorClass
	// parent_class : record
	// no type for get_mount_for_mount_path
}

func NativeVolumeMonitorClassNewFromC(u unsafe.Pointer) *NativeVolumeMonitorClass {
	c := (*C.GNativeVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitorClass{native: c}

	return g
}

func (recv *NativeVolumeMonitorClass) toC() *C.GNativeVolumeMonitorClass {

	return recv.native
}

// NetworkAddressClass is a wrapper around the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native *C.GNetworkAddressClass
	// parent_class : record
}

func NetworkAddressClassNewFromC(u unsafe.Pointer) *NetworkAddressClass {
	c := (*C.GNetworkAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressClass{native: c}

	return g
}

func (recv *NetworkAddressClass) toC() *C.GNetworkAddressClass {

	return recv.native
}

// NetworkAddressPrivate is a wrapper around the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native *C.GNetworkAddressPrivate
}

func NetworkAddressPrivateNewFromC(u unsafe.Pointer) *NetworkAddressPrivate {
	c := (*C.GNetworkAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressPrivate{native: c}

	return g
}

func (recv *NetworkAddressPrivate) toC() *C.GNetworkAddressPrivate {

	return recv.native
}

// NetworkServiceClass is a wrapper around the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native *C.GNetworkServiceClass
	// parent_class : record
}

func NetworkServiceClassNewFromC(u unsafe.Pointer) *NetworkServiceClass {
	c := (*C.GNetworkServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServiceClass{native: c}

	return g
}

func (recv *NetworkServiceClass) toC() *C.GNetworkServiceClass {

	return recv.native
}

// NetworkServicePrivate is a wrapper around the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native *C.GNetworkServicePrivate
}

func NetworkServicePrivateNewFromC(u unsafe.Pointer) *NetworkServicePrivate {
	c := (*C.GNetworkServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServicePrivate{native: c}

	return g
}

func (recv *NetworkServicePrivate) toC() *C.GNetworkServicePrivate {

	return recv.native
}

// OutputStreamClass is a wrapper around the C record GOutputStreamClass.
type OutputStreamClass struct {
	native *C.GOutputStreamClass
	// parent_class : record
	// no type for write_fn
	// no type for splice
	// no type for flush
	// no type for close_fn
	// no type for write_async
	// no type for write_finish
	// no type for splice_async
	// no type for splice_finish
	// no type for flush_async
	// no type for flush_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
}

func OutputStreamClassNewFromC(u unsafe.Pointer) *OutputStreamClass {
	c := (*C.GOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamClass{native: c}

	return g
}

func (recv *OutputStreamClass) toC() *C.GOutputStreamClass {

	return recv.native
}

// OutputStreamPrivate is a wrapper around the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native *C.GOutputStreamPrivate
}

func OutputStreamPrivateNewFromC(u unsafe.Pointer) *OutputStreamPrivate {
	c := (*C.GOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamPrivate{native: c}

	return g
}

func (recv *OutputStreamPrivate) toC() *C.GOutputStreamPrivate {

	return recv.native
}

// PermissionClass is a wrapper around the C record GPermissionClass.
type PermissionClass struct {
	native *C.GPermissionClass
	// parent_class : record
	// no type for acquire
	// no type for acquire_async
	// no type for acquire_finish
	// no type for release
	// no type for release_async
	// no type for release_finish
	// no type for reserved
}

func PermissionClassNewFromC(u unsafe.Pointer) *PermissionClass {
	c := (*C.GPermissionClass)(u)
	if c == nil {
		return nil
	}

	g := &PermissionClass{native: c}

	return g
}

func (recv *PermissionClass) toC() *C.GPermissionClass {

	return recv.native
}

// PermissionPrivate is a wrapper around the C record GPermissionPrivate.
type PermissionPrivate struct {
	native *C.GPermissionPrivate
}

func PermissionPrivateNewFromC(u unsafe.Pointer) *PermissionPrivate {
	c := (*C.GPermissionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PermissionPrivate{native: c}

	return g
}

func (recv *PermissionPrivate) toC() *C.GPermissionPrivate {

	return recv.native
}

// ProxyAddressEnumeratorClass is a wrapper around the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native *C.GProxyAddressEnumeratorClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func ProxyAddressEnumeratorClassNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorClass {
	c := (*C.GProxyAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorClass{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorClass) toC() *C.GProxyAddressEnumeratorClass {

	return recv.native
}

// ProxyAddressEnumeratorPrivate is a wrapper around the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native *C.GProxyAddressEnumeratorPrivate
}

func ProxyAddressEnumeratorPrivateNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	c := (*C.GProxyAddressEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorPrivate{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorPrivate) toC() *C.GProxyAddressEnumeratorPrivate {

	return recv.native
}

// ProxyAddressPrivate is a wrapper around the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native *C.GProxyAddressPrivate
}

func ProxyAddressPrivateNewFromC(u unsafe.Pointer) *ProxyAddressPrivate {
	c := (*C.GProxyAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressPrivate{native: c}

	return g
}

func (recv *ProxyAddressPrivate) toC() *C.GProxyAddressPrivate {

	return recv.native
}

// ProxyResolverInterface is a wrapper around the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native *C.GProxyResolverInterface
	// g_iface : record
	// no type for is_supported
	// no type for lookup
	// no type for lookup_async
	// no type for lookup_finish
}

func ProxyResolverInterfaceNewFromC(u unsafe.Pointer) *ProxyResolverInterface {
	c := (*C.GProxyResolverInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolverInterface{native: c}

	return g
}

func (recv *ProxyResolverInterface) toC() *C.GProxyResolverInterface {

	return recv.native
}

// ResolverClass is a wrapper around the C record GResolverClass.
type ResolverClass struct {
	native *C.GResolverClass
	// parent_class : record
	// no type for reload
	// no type for lookup_by_name
	// no type for lookup_by_name_async
	// no type for lookup_by_name_finish
	// no type for lookup_by_address
	// no type for lookup_by_address_async
	// no type for lookup_by_address_finish
	// no type for lookup_service
	// no type for lookup_service_async
	// no type for lookup_service_finish
	// no type for lookup_records
	// no type for lookup_records_async
	// no type for lookup_records_finish
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func ResolverClassNewFromC(u unsafe.Pointer) *ResolverClass {
	c := (*C.GResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &ResolverClass{native: c}

	return g
}

func (recv *ResolverClass) toC() *C.GResolverClass {

	return recv.native
}

// ResolverPrivate is a wrapper around the C record GResolverPrivate.
type ResolverPrivate struct {
	native *C.GResolverPrivate
}

func ResolverPrivateNewFromC(u unsafe.Pointer) *ResolverPrivate {
	c := (*C.GResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ResolverPrivate{native: c}

	return g
}

func (recv *ResolverPrivate) toC() *C.GResolverPrivate {

	return recv.native
}

// SeekableIface is a wrapper around the C record GSeekableIface.
type SeekableIface struct {
	native *C.GSeekableIface
	// g_iface : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
}

func SeekableIfaceNewFromC(u unsafe.Pointer) *SeekableIface {
	c := (*C.GSeekableIface)(u)
	if c == nil {
		return nil
	}

	g := &SeekableIface{native: c}

	return g
}

func (recv *SeekableIface) toC() *C.GSeekableIface {

	return recv.native
}

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// SettingsClass is a wrapper around the C record GSettingsClass.
type SettingsClass struct {
	native *C.GSettingsClass
	// parent_class : record
	// no type for writable_changed
	// no type for changed
	// no type for writable_change_event
	// no type for change_event
	// no type for padding
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	c := (*C.GSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SettingsClass{native: c}

	return g
}

func (recv *SettingsClass) toC() *C.GSettingsClass {

	return recv.native
}

// SettingsPrivate is a wrapper around the C record GSettingsPrivate.
type SettingsPrivate struct {
	native *C.GSettingsPrivate
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	c := (*C.GSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SettingsPrivate{native: c}

	return g
}

func (recv *SettingsPrivate) toC() *C.GSettingsPrivate {

	return recv.native
}

// SettingsSchemaKey is a wrapper around the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native *C.GSettingsSchemaKey
}

func SettingsSchemaKeyNewFromC(u unsafe.Pointer) *SettingsSchemaKey {
	c := (*C.GSettingsSchemaKey)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchemaKey{native: c}

	return g
}

func (recv *SettingsSchemaKey) toC() *C.GSettingsSchemaKey {

	return recv.native
}

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_unref : no return generator

// SimpleActionGroupClass is a wrapper around the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native *C.GSimpleActionGroupClass
	// Private : parent_class
	// Private : padding
}

func SimpleActionGroupClassNewFromC(u unsafe.Pointer) *SimpleActionGroupClass {
	c := (*C.GSimpleActionGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupClass{native: c}

	return g
}

func (recv *SimpleActionGroupClass) toC() *C.GSimpleActionGroupClass {

	return recv.native
}

// SimpleActionGroupPrivate is a wrapper around the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native *C.GSimpleActionGroupPrivate
}

func SimpleActionGroupPrivateNewFromC(u unsafe.Pointer) *SimpleActionGroupPrivate {
	c := (*C.GSimpleActionGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupPrivate{native: c}

	return g
}

func (recv *SimpleActionGroupPrivate) toC() *C.GSimpleActionGroupPrivate {

	return recv.native
}

// SimpleAsyncResultClass is a wrapper around the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native *C.GSimpleAsyncResultClass
}

func SimpleAsyncResultClassNewFromC(u unsafe.Pointer) *SimpleAsyncResultClass {
	c := (*C.GSimpleAsyncResultClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResultClass{native: c}

	return g
}

func (recv *SimpleAsyncResultClass) toC() *C.GSimpleAsyncResultClass {

	return recv.native
}

// SimpleProxyResolverClass is a wrapper around the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native *C.GSimpleProxyResolverClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SimpleProxyResolverClassNewFromC(u unsafe.Pointer) *SimpleProxyResolverClass {
	c := (*C.GSimpleProxyResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverClass{native: c}

	return g
}

func (recv *SimpleProxyResolverClass) toC() *C.GSimpleProxyResolverClass {

	return recv.native
}

// SimpleProxyResolverPrivate is a wrapper around the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native *C.GSimpleProxyResolverPrivate
}

func SimpleProxyResolverPrivateNewFromC(u unsafe.Pointer) *SimpleProxyResolverPrivate {
	c := (*C.GSimpleProxyResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverPrivate{native: c}

	return g
}

func (recv *SimpleProxyResolverPrivate) toC() *C.GSimpleProxyResolverPrivate {

	return recv.native
}

// SocketAddressClass is a wrapper around the C record GSocketAddressClass.
type SocketAddressClass struct {
	native *C.GSocketAddressClass
	// parent_class : record
	// no type for get_family
	// no type for get_native_size
	// no type for to_native
}

func SocketAddressClassNewFromC(u unsafe.Pointer) *SocketAddressClass {
	c := (*C.GSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressClass{native: c}

	return g
}

func (recv *SocketAddressClass) toC() *C.GSocketAddressClass {

	return recv.native
}

// SocketAddressEnumeratorClass is a wrapper around the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native *C.GSocketAddressEnumeratorClass
	// parent_class : record
	// no type for next
	// no type for next_async
	// no type for next_finish
}

func SocketAddressEnumeratorClassNewFromC(u unsafe.Pointer) *SocketAddressEnumeratorClass {
	c := (*C.GSocketAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumeratorClass{native: c}

	return g
}

func (recv *SocketAddressEnumeratorClass) toC() *C.GSocketAddressEnumeratorClass {

	return recv.native
}

// SocketClass is a wrapper around the C record GSocketClass.
type SocketClass struct {
	native *C.GSocketClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.GSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) toC() *C.GSocketClass {

	return recv.native
}

// SocketClientClass is a wrapper around the C record GSocketClientClass.
type SocketClientClass struct {
	native *C.GSocketClientClass
	// parent_class : record
	// no type for event
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func SocketClientClassNewFromC(u unsafe.Pointer) *SocketClientClass {
	c := (*C.GSocketClientClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientClass{native: c}

	return g
}

func (recv *SocketClientClass) toC() *C.GSocketClientClass {

	return recv.native
}

// SocketClientPrivate is a wrapper around the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native *C.GSocketClientPrivate
}

func SocketClientPrivateNewFromC(u unsafe.Pointer) *SocketClientPrivate {
	c := (*C.GSocketClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientPrivate{native: c}

	return g
}

func (recv *SocketClientPrivate) toC() *C.GSocketClientPrivate {

	return recv.native
}

// SocketConnectableIface is a wrapper around the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native *C.GSocketConnectableIface
	// g_iface : record
	// no type for enumerate
	// no type for proxy_enumerate
	// no type for to_string
}

func SocketConnectableIfaceNewFromC(u unsafe.Pointer) *SocketConnectableIface {
	c := (*C.GSocketConnectableIface)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectableIface{native: c}

	return g
}

func (recv *SocketConnectableIface) toC() *C.GSocketConnectableIface {

	return recv.native
}

// SocketConnectionClass is a wrapper around the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native *C.GSocketConnectionClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketConnectionClassNewFromC(u unsafe.Pointer) *SocketConnectionClass {
	c := (*C.GSocketConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionClass{native: c}

	return g
}

func (recv *SocketConnectionClass) toC() *C.GSocketConnectionClass {

	return recv.native
}

// SocketConnectionPrivate is a wrapper around the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native *C.GSocketConnectionPrivate
}

func SocketConnectionPrivateNewFromC(u unsafe.Pointer) *SocketConnectionPrivate {
	c := (*C.GSocketConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionPrivate{native: c}

	return g
}

func (recv *SocketConnectionPrivate) toC() *C.GSocketConnectionPrivate {

	return recv.native
}

// SocketControlMessageClass is a wrapper around the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native *C.GSocketControlMessageClass
	// parent_class : record
	// no type for get_size
	// no type for get_level
	// no type for get_type
	// no type for serialize
	// no type for deserialize
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SocketControlMessageClassNewFromC(u unsafe.Pointer) *SocketControlMessageClass {
	c := (*C.GSocketControlMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessageClass{native: c}

	return g
}

func (recv *SocketControlMessageClass) toC() *C.GSocketControlMessageClass {

	return recv.native
}

// SocketControlMessagePrivate is a wrapper around the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native *C.GSocketControlMessagePrivate
}

func SocketControlMessagePrivateNewFromC(u unsafe.Pointer) *SocketControlMessagePrivate {
	c := (*C.GSocketControlMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessagePrivate{native: c}

	return g
}

func (recv *SocketControlMessagePrivate) toC() *C.GSocketControlMessagePrivate {

	return recv.native
}

// SocketListenerClass is a wrapper around the C record GSocketListenerClass.
type SocketListenerClass struct {
	native *C.GSocketListenerClass
	// parent_class : record
	// no type for changed
	// no type for event
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketListenerClassNewFromC(u unsafe.Pointer) *SocketListenerClass {
	c := (*C.GSocketListenerClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerClass{native: c}

	return g
}

func (recv *SocketListenerClass) toC() *C.GSocketListenerClass {

	return recv.native
}

// SocketListenerPrivate is a wrapper around the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native *C.GSocketListenerPrivate
}

func SocketListenerPrivateNewFromC(u unsafe.Pointer) *SocketListenerPrivate {
	c := (*C.GSocketListenerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerPrivate{native: c}

	return g
}

func (recv *SocketListenerPrivate) toC() *C.GSocketListenerPrivate {

	return recv.native
}

// SocketPrivate is a wrapper around the C record GSocketPrivate.
type SocketPrivate struct {
	native *C.GSocketPrivate
}

func SocketPrivateNewFromC(u unsafe.Pointer) *SocketPrivate {
	c := (*C.GSocketPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketPrivate{native: c}

	return g
}

func (recv *SocketPrivate) toC() *C.GSocketPrivate {

	return recv.native
}

// SocketServiceClass is a wrapper around the C record GSocketServiceClass.
type SocketServiceClass struct {
	native *C.GSocketServiceClass
	// parent_class : record
	// no type for incoming
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketServiceClassNewFromC(u unsafe.Pointer) *SocketServiceClass {
	c := (*C.GSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketServiceClass{native: c}

	return g
}

func (recv *SocketServiceClass) toC() *C.GSocketServiceClass {

	return recv.native
}

// SocketServicePrivate is a wrapper around the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native *C.GSocketServicePrivate
}

func SocketServicePrivateNewFromC(u unsafe.Pointer) *SocketServicePrivate {
	c := (*C.GSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketServicePrivate{native: c}

	return g
}

func (recv *SocketServicePrivate) toC() *C.GSocketServicePrivate {

	return recv.native
}

// SrvTarget is a wrapper around the C record GSrvTarget.
type SrvTarget struct {
	native *C.GSrvTarget
}

func SrvTargetNewFromC(u unsafe.Pointer) *SrvTarget {
	c := (*C.GSrvTarget)(u)
	if c == nil {
		return nil
	}

	g := &SrvTarget{native: c}

	return g
}

func (recv *SrvTarget) toC() *C.GSrvTarget {

	return recv.native
}

// SrvTargetNew is a wrapper around the C function g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	c_priority := (C.guint16)(priority)

	c_weight := (C.guint16)(weight)

	retC := C.g_srv_target_new(c_hostname, c_port, c_priority, c_weight)
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_srv_target_free : no return generator

// StaticResource is a wrapper around the C record GStaticResource.
type StaticResource struct {
	native *C.GStaticResource
	// Private : data
	// Private : data_len
	// Private : resource
	// Private : next
	// Private : padding
}

func StaticResourceNewFromC(u unsafe.Pointer) *StaticResource {
	c := (*C.GStaticResource)(u)
	if c == nil {
		return nil
	}

	g := &StaticResource{native: c}

	return g
}

func (recv *StaticResource) toC() *C.GStaticResource {

	return recv.native
}

// Unsupported : g_static_resource_fini : no return generator

// Unsupported : g_static_resource_init : no return generator

// TaskClass is a wrapper around the C record GTaskClass.
type TaskClass struct {
	native *C.GTaskClass
}

func TaskClassNewFromC(u unsafe.Pointer) *TaskClass {
	c := (*C.GTaskClass)(u)
	if c == nil {
		return nil
	}

	g := &TaskClass{native: c}

	return g
}

func (recv *TaskClass) toC() *C.GTaskClass {

	return recv.native
}

// TcpConnectionClass is a wrapper around the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native *C.GTcpConnectionClass
	// parent_class : record
}

func TcpConnectionClassNewFromC(u unsafe.Pointer) *TcpConnectionClass {
	c := (*C.GTcpConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionClass{native: c}

	return g
}

func (recv *TcpConnectionClass) toC() *C.GTcpConnectionClass {

	return recv.native
}

// TcpConnectionPrivate is a wrapper around the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native *C.GTcpConnectionPrivate
}

func TcpConnectionPrivateNewFromC(u unsafe.Pointer) *TcpConnectionPrivate {
	c := (*C.GTcpConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionPrivate{native: c}

	return g
}

func (recv *TcpConnectionPrivate) toC() *C.GTcpConnectionPrivate {

	return recv.native
}

// TcpWrapperConnectionClass is a wrapper around the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native *C.GTcpWrapperConnectionClass
	// parent_class : record
}

func TcpWrapperConnectionClassNewFromC(u unsafe.Pointer) *TcpWrapperConnectionClass {
	c := (*C.GTcpWrapperConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionClass{native: c}

	return g
}

func (recv *TcpWrapperConnectionClass) toC() *C.GTcpWrapperConnectionClass {

	return recv.native
}

// TcpWrapperConnectionPrivate is a wrapper around the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native *C.GTcpWrapperConnectionPrivate
}

func TcpWrapperConnectionPrivateNewFromC(u unsafe.Pointer) *TcpWrapperConnectionPrivate {
	c := (*C.GTcpWrapperConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionPrivate{native: c}

	return g
}

func (recv *TcpWrapperConnectionPrivate) toC() *C.GTcpWrapperConnectionPrivate {

	return recv.native
}

// ThemedIconClass is a wrapper around the C record GThemedIconClass.
type ThemedIconClass struct {
	native *C.GThemedIconClass
}

func ThemedIconClassNewFromC(u unsafe.Pointer) *ThemedIconClass {
	c := (*C.GThemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIconClass{native: c}

	return g
}

func (recv *ThemedIconClass) toC() *C.GThemedIconClass {

	return recv.native
}

// ThreadedSocketServiceClass is a wrapper around the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native *C.GThreadedSocketServiceClass
	// parent_class : record
	// no type for run
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ThreadedSocketServiceClassNewFromC(u unsafe.Pointer) *ThreadedSocketServiceClass {
	c := (*C.GThreadedSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServiceClass{native: c}

	return g
}

func (recv *ThreadedSocketServiceClass) toC() *C.GThreadedSocketServiceClass {

	return recv.native
}

// ThreadedSocketServicePrivate is a wrapper around the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native *C.GThreadedSocketServicePrivate
}

func ThreadedSocketServicePrivateNewFromC(u unsafe.Pointer) *ThreadedSocketServicePrivate {
	c := (*C.GThreadedSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServicePrivate{native: c}

	return g
}

func (recv *ThreadedSocketServicePrivate) toC() *C.GThreadedSocketServicePrivate {

	return recv.native
}

// TlsCertificateClass is a wrapper around the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native *C.GTlsCertificateClass
	// parent_class : record
	// no type for verify
	// Private : padding
}

func TlsCertificateClassNewFromC(u unsafe.Pointer) *TlsCertificateClass {
	c := (*C.GTlsCertificateClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificateClass{native: c}

	return g
}

func (recv *TlsCertificateClass) toC() *C.GTlsCertificateClass {

	return recv.native
}

// TlsCertificatePrivate is a wrapper around the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native *C.GTlsCertificatePrivate
}

func TlsCertificatePrivateNewFromC(u unsafe.Pointer) *TlsCertificatePrivate {
	c := (*C.GTlsCertificatePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificatePrivate{native: c}

	return g
}

func (recv *TlsCertificatePrivate) toC() *C.GTlsCertificatePrivate {

	return recv.native
}

// TlsConnectionClass is a wrapper around the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native *C.GTlsConnectionClass
	// parent_class : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// Private : padding
}

func TlsConnectionClassNewFromC(u unsafe.Pointer) *TlsConnectionClass {
	c := (*C.GTlsConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionClass{native: c}

	return g
}

func (recv *TlsConnectionClass) toC() *C.GTlsConnectionClass {

	return recv.native
}

// TlsConnectionPrivate is a wrapper around the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native *C.GTlsConnectionPrivate
}

func TlsConnectionPrivateNewFromC(u unsafe.Pointer) *TlsConnectionPrivate {
	c := (*C.GTlsConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionPrivate{native: c}

	return g
}

func (recv *TlsConnectionPrivate) toC() *C.GTlsConnectionPrivate {

	return recv.native
}

// TlsDatabasePrivate is a wrapper around the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native *C.GTlsDatabasePrivate
}

func TlsDatabasePrivateNewFromC(u unsafe.Pointer) *TlsDatabasePrivate {
	c := (*C.GTlsDatabasePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabasePrivate{native: c}

	return g
}

func (recv *TlsDatabasePrivate) toC() *C.GTlsDatabasePrivate {

	return recv.native
}

// TlsFileDatabaseInterface is a wrapper around the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native *C.GTlsFileDatabaseInterface
	// g_iface : record
	// Private : padding
}

func TlsFileDatabaseInterfaceNewFromC(u unsafe.Pointer) *TlsFileDatabaseInterface {
	c := (*C.GTlsFileDatabaseInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsFileDatabaseInterface{native: c}

	return g
}

func (recv *TlsFileDatabaseInterface) toC() *C.GTlsFileDatabaseInterface {

	return recv.native
}

// TlsInteractionPrivate is a wrapper around the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native *C.GTlsInteractionPrivate
}

func TlsInteractionPrivateNewFromC(u unsafe.Pointer) *TlsInteractionPrivate {
	c := (*C.GTlsInteractionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteractionPrivate{native: c}

	return g
}

func (recv *TlsInteractionPrivate) toC() *C.GTlsInteractionPrivate {

	return recv.native
}

// TlsPasswordClass is a wrapper around the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native *C.GTlsPasswordClass
	// parent_class : record
	// no type for get_value
	// no type for set_value
	// no type for get_default_warning
	// Private : padding
}

func TlsPasswordClassNewFromC(u unsafe.Pointer) *TlsPasswordClass {
	c := (*C.GTlsPasswordClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordClass{native: c}

	return g
}

func (recv *TlsPasswordClass) toC() *C.GTlsPasswordClass {

	return recv.native
}

// TlsPasswordPrivate is a wrapper around the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native *C.GTlsPasswordPrivate
}

func TlsPasswordPrivateNewFromC(u unsafe.Pointer) *TlsPasswordPrivate {
	c := (*C.GTlsPasswordPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordPrivate{native: c}

	return g
}

func (recv *TlsPasswordPrivate) toC() *C.GTlsPasswordPrivate {

	return recv.native
}

// UnixConnectionClass is a wrapper around the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native *C.GUnixConnectionClass
	// parent_class : record
}

func UnixConnectionClassNewFromC(u unsafe.Pointer) *UnixConnectionClass {
	c := (*C.GUnixConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionClass{native: c}

	return g
}

func (recv *UnixConnectionClass) toC() *C.GUnixConnectionClass {

	return recv.native
}

// UnixConnectionPrivate is a wrapper around the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native *C.GUnixConnectionPrivate
}

func UnixConnectionPrivateNewFromC(u unsafe.Pointer) *UnixConnectionPrivate {
	c := (*C.GUnixConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionPrivate{native: c}

	return g
}

func (recv *UnixConnectionPrivate) toC() *C.GUnixConnectionPrivate {

	return recv.native
}

// UnixCredentialsMessagePrivate is a wrapper around the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native *C.GUnixCredentialsMessagePrivate
}

func UnixCredentialsMessagePrivateNewFromC(u unsafe.Pointer) *UnixCredentialsMessagePrivate {
	c := (*C.GUnixCredentialsMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessagePrivate{native: c}

	return g
}

func (recv *UnixCredentialsMessagePrivate) toC() *C.GUnixCredentialsMessagePrivate {

	return recv.native
}

// UnixFDListClass is a wrapper around the C record GUnixFDListClass.
type UnixFDListClass struct {
	native *C.GUnixFDListClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixFDListClassNewFromC(u unsafe.Pointer) *UnixFDListClass {
	c := (*C.GUnixFDListClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListClass{native: c}

	return g
}

func (recv *UnixFDListClass) toC() *C.GUnixFDListClass {

	return recv.native
}

// UnixFDListPrivate is a wrapper around the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native *C.GUnixFDListPrivate
}

func UnixFDListPrivateNewFromC(u unsafe.Pointer) *UnixFDListPrivate {
	c := (*C.GUnixFDListPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListPrivate{native: c}

	return g
}

func (recv *UnixFDListPrivate) toC() *C.GUnixFDListPrivate {

	return recv.native
}

// UnixFDMessageClass is a wrapper around the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native *C.GUnixFDMessageClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func UnixFDMessageClassNewFromC(u unsafe.Pointer) *UnixFDMessageClass {
	c := (*C.GUnixFDMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessageClass{native: c}

	return g
}

func (recv *UnixFDMessageClass) toC() *C.GUnixFDMessageClass {

	return recv.native
}

// UnixFDMessagePrivate is a wrapper around the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native *C.GUnixFDMessagePrivate
}

func UnixFDMessagePrivateNewFromC(u unsafe.Pointer) *UnixFDMessagePrivate {
	c := (*C.GUnixFDMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessagePrivate{native: c}

	return g
}

func (recv *UnixFDMessagePrivate) toC() *C.GUnixFDMessagePrivate {

	return recv.native
}

// UnixInputStreamClass is a wrapper around the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native *C.GUnixInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixInputStreamClassNewFromC(u unsafe.Pointer) *UnixInputStreamClass {
	c := (*C.GUnixInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamClass{native: c}

	return g
}

func (recv *UnixInputStreamClass) toC() *C.GUnixInputStreamClass {

	return recv.native
}

// UnixInputStreamPrivate is a wrapper around the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native *C.GUnixInputStreamPrivate
}

func UnixInputStreamPrivateNewFromC(u unsafe.Pointer) *UnixInputStreamPrivate {
	c := (*C.GUnixInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamPrivate{native: c}

	return g
}

func (recv *UnixInputStreamPrivate) toC() *C.GUnixInputStreamPrivate {

	return recv.native
}

// UnixMountEntry is a wrapper around the C record GUnixMountEntry.
type UnixMountEntry struct {
	native *C.GUnixMountEntry
}

func UnixMountEntryNewFromC(u unsafe.Pointer) *UnixMountEntry {
	c := (*C.GUnixMountEntry)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountEntry{native: c}

	return g
}

func (recv *UnixMountEntry) toC() *C.GUnixMountEntry {

	return recv.native
}

// UnixMountMonitorClass is a wrapper around the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native *C.GUnixMountMonitorClass
}

func UnixMountMonitorClassNewFromC(u unsafe.Pointer) *UnixMountMonitorClass {
	c := (*C.GUnixMountMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitorClass{native: c}

	return g
}

func (recv *UnixMountMonitorClass) toC() *C.GUnixMountMonitorClass {

	return recv.native
}

// UnixMountPoint is a wrapper around the C record GUnixMountPoint.
type UnixMountPoint struct {
	native *C.GUnixMountPoint
}

func UnixMountPointNewFromC(u unsafe.Pointer) *UnixMountPoint {
	c := (*C.GUnixMountPoint)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountPoint{native: c}

	return g
}

func (recv *UnixMountPoint) toC() *C.GUnixMountPoint {

	return recv.native
}

// Compare is a wrapper around the C function g_unix_mount_point_compare.
func (recv *UnixMountPoint) Compare(mount2 *UnixMountPoint) int32 {
	c_mount2 := mount2.toC()

	retC := C.g_unix_mount_point_compare((*C.GUnixMountPoint)(recv.native), c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_unix_mount_point_free : no return generator

// GetDevicePath is a wrapper around the C function g_unix_mount_point_get_device_path.
func (recv *UnixMountPoint) GetDevicePath() string {
	retC := C.g_unix_mount_point_get_device_path((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFsType is a wrapper around the C function g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() string {
	retC := C.g_unix_mount_point_get_fs_type((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMountPath is a wrapper around the C function g_unix_mount_point_get_mount_path.
func (recv *UnixMountPoint) GetMountPath() string {
	retC := C.g_unix_mount_point_get_mount_path((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GuessCanEject is a wrapper around the C function g_unix_mount_point_guess_can_eject.
func (recv *UnixMountPoint) GuessCanEject() bool {
	retC := C.g_unix_mount_point_guess_can_eject((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_unix_mount_point_guess_icon : no return generator

// GuessName is a wrapper around the C function g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() string {
	retC := C.g_unix_mount_point_guess_name((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_mount_point_guess_symbolic_icon : no return generator

// IsLoopback is a wrapper around the C function g_unix_mount_point_is_loopback.
func (recv *UnixMountPoint) IsLoopback() bool {
	retC := C.g_unix_mount_point_is_loopback((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsReadonly is a wrapper around the C function g_unix_mount_point_is_readonly.
func (recv *UnixMountPoint) IsReadonly() bool {
	retC := C.g_unix_mount_point_is_readonly((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsUserMountable is a wrapper around the C function g_unix_mount_point_is_user_mountable.
func (recv *UnixMountPoint) IsUserMountable() bool {
	retC := C.g_unix_mount_point_is_user_mountable((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnixOutputStreamClass is a wrapper around the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native *C.GUnixOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixOutputStreamClassNewFromC(u unsafe.Pointer) *UnixOutputStreamClass {
	c := (*C.GUnixOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamClass{native: c}

	return g
}

func (recv *UnixOutputStreamClass) toC() *C.GUnixOutputStreamClass {

	return recv.native
}

// UnixOutputStreamPrivate is a wrapper around the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native *C.GUnixOutputStreamPrivate
}

func UnixOutputStreamPrivateNewFromC(u unsafe.Pointer) *UnixOutputStreamPrivate {
	c := (*C.GUnixOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamPrivate{native: c}

	return g
}

func (recv *UnixOutputStreamPrivate) toC() *C.GUnixOutputStreamPrivate {

	return recv.native
}

// UnixSocketAddressClass is a wrapper around the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native *C.GUnixSocketAddressClass
	// parent_class : record
}

func UnixSocketAddressClassNewFromC(u unsafe.Pointer) *UnixSocketAddressClass {
	c := (*C.GUnixSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressClass{native: c}

	return g
}

func (recv *UnixSocketAddressClass) toC() *C.GUnixSocketAddressClass {

	return recv.native
}

// UnixSocketAddressPrivate is a wrapper around the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native *C.GUnixSocketAddressPrivate
}

func UnixSocketAddressPrivateNewFromC(u unsafe.Pointer) *UnixSocketAddressPrivate {
	c := (*C.GUnixSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressPrivate{native: c}

	return g
}

func (recv *UnixSocketAddressPrivate) toC() *C.GUnixSocketAddressPrivate {

	return recv.native
}

// VfsClass is a wrapper around the C record GVfsClass.
type VfsClass struct {
	native *C.GVfsClass
	// parent_class : record
	// no type for is_active
	// no type for get_file_for_path
	// no type for get_file_for_uri
	// no type for get_supported_uri_schemes
	// no type for parse_name
	// no type for local_file_add_info
	// no type for add_writable_namespaces
	// no type for local_file_set_attributes
	// no type for local_file_removed
	// no type for local_file_moved
	// no type for deserialize_icon
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VfsClassNewFromC(u unsafe.Pointer) *VfsClass {
	c := (*C.GVfsClass)(u)
	if c == nil {
		return nil
	}

	g := &VfsClass{native: c}

	return g
}

func (recv *VfsClass) toC() *C.GVfsClass {

	return recv.native
}

// VolumeIface is a wrapper around the C record GVolumeIface.
type VolumeIface struct {
	native *C.GVolumeIface
	// g_iface : record
	// no type for changed
	// no type for removed
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_drive
	// no type for get_mount
	// no type for can_mount
	// no type for can_eject
	// no type for mount_fn
	// no type for mount_finish
	// no type for eject
	// no type for eject_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for should_automount
	// no type for get_activation_root
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func VolumeIfaceNewFromC(u unsafe.Pointer) *VolumeIface {
	c := (*C.GVolumeIface)(u)
	if c == nil {
		return nil
	}

	g := &VolumeIface{native: c}

	return g
}

func (recv *VolumeIface) toC() *C.GVolumeIface {

	return recv.native
}

// VolumeMonitorClass is a wrapper around the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native *C.GVolumeMonitorClass
	// parent_class : record
	// no type for volume_added
	// no type for volume_removed
	// no type for volume_changed
	// no type for mount_added
	// no type for mount_removed
	// no type for mount_pre_unmount
	// no type for mount_changed
	// no type for drive_connected
	// no type for drive_disconnected
	// no type for drive_changed
	// no type for is_supported
	// no type for get_connected_drives
	// no type for get_volumes
	// no type for get_mounts
	// no type for get_volume_for_uuid
	// no type for get_mount_for_uuid
	// no type for adopt_orphan_mount
	// no type for drive_eject_button
	// no type for drive_stop_button
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VolumeMonitorClassNewFromC(u unsafe.Pointer) *VolumeMonitorClass {
	c := (*C.GVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitorClass{native: c}

	return g
}

func (recv *VolumeMonitorClass) toC() *C.GVolumeMonitorClass {

	return recv.native
}

// ZlibCompressorClass is a wrapper around the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native *C.GZlibCompressorClass
	// parent_class : record
}

func ZlibCompressorClassNewFromC(u unsafe.Pointer) *ZlibCompressorClass {
	c := (*C.GZlibCompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressorClass{native: c}

	return g
}

func (recv *ZlibCompressorClass) toC() *C.GZlibCompressorClass {

	return recv.native
}

// ZlibDecompressorClass is a wrapper around the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native *C.GZlibDecompressorClass
	// parent_class : record
}

func ZlibDecompressorClassNewFromC(u unsafe.Pointer) *ZlibDecompressorClass {
	c := (*C.GZlibDecompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressorClass{native: c}

	return g
}

func (recv *ZlibDecompressorClass) toC() *C.GZlibDecompressorClass {

	return recv.native
}
