// This is a generated file - DO NOT EDIT

package gio

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
/*

	void drive_changedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(drive_changedHandler), data);
	}

*/
/*

	void drive_disconnectedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_disconnected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnected", G_CALLBACK(drive_disconnectedHandler), data);
	}

*/
/*

	void drive_ejectButtonHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "eject-button", G_CALLBACK(drive_ejectButtonHandler), data);
	}

*/
/*

	void mount_changedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(mount_changedHandler), data);
	}

*/
/*

	void mount_unmountedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_unmounted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmounted", G_CALLBACK(mount_unmountedHandler), data);
	}

*/
/*

	void volume_changedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(volume_changedHandler), data);
	}

*/
/*

	void volume_removedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "removed", G_CALLBACK(volume_removedHandler), data);
	}

*/
import "C"

// Action is a wrapper around the C record GAction.
type Action struct {
	native *C.GAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// ActionGroup is a wrapper around the C record GActionGroup.
type ActionGroup struct {
	native *C.GActionGroup
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionGroup with another ActionGroup, and returns true if they represent the same GObject.
func (recv *ActionGroup) Equals(other *ActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// ActionMap is a wrapper around the C record GActionMap.
type ActionMap struct {
	native *C.GActionMap
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	c := (*C.GActionMap)(u)
	if c == nil {
		return nil
	}

	g := &ActionMap{native: c}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionMap with another ActionMap, and returns true if they represent the same GObject.
func (recv *ActionMap) Equals(other *ActionMap) bool {
	return other.ToC() == recv.ToC()
}

// AppInfo is a wrapper around the C record GAppInfo.
type AppInfo struct {
	native *C.GAppInfo
}

func AppInfoNewFromC(u unsafe.Pointer) *AppInfo {
	c := (*C.GAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &AppInfo{native: c}

	return g
}

func (recv *AppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppInfo with another AppInfo, and returns true if they represent the same GObject.
func (recv *AppInfo) Equals(other *AppInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_app_info_create_from_commandline

// Blacklisted : g_app_info_get_all

// Blacklisted : g_app_info_get_all_for_type

// Blacklisted : g_app_info_get_default_for_type

// Blacklisted : g_app_info_get_default_for_uri_scheme

// Blacklisted : g_app_info_launch_default_for_uri

// Blacklisted : g_app_info_add_supports_type

// Blacklisted : g_app_info_can_remove_supports_type

// Blacklisted : g_app_info_dup

// Blacklisted : g_app_info_equal

// Blacklisted : g_app_info_get_description

// Blacklisted : g_app_info_get_executable

// Blacklisted : g_app_info_get_icon

// Blacklisted : g_app_info_get_id

// Blacklisted : g_app_info_get_name

// Blacklisted : g_app_info_launch

// Blacklisted : g_app_info_launch_uris

// Blacklisted : g_app_info_remove_supports_type

// Blacklisted : g_app_info_set_as_default_for_extension

// Blacklisted : g_app_info_set_as_default_for_type

// Blacklisted : g_app_info_set_as_last_used_for_type

// Blacklisted : g_app_info_should_show

// Blacklisted : g_app_info_supports_files

// Blacklisted : g_app_info_supports_uris

// AsyncResult is a wrapper around the C record GAsyncResult.
type AsyncResult struct {
	native *C.GAsyncResult
}

func AsyncResultNewFromC(u unsafe.Pointer) *AsyncResult {
	c := (*C.GAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResult{native: c}

	return g
}

func (recv *AsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncResult with another AsyncResult, and returns true if they represent the same GObject.
func (recv *AsyncResult) Equals(other *AsyncResult) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_async_result_get_source_object

// Unsupported : g_async_result_get_user_data : no return generator

// DBusObject is a wrapper around the C record GDBusObject.
type DBusObject struct {
	native *C.GDBusObject
}

func DBusObjectNewFromC(u unsafe.Pointer) *DBusObject {
	c := (*C.GDBusObject)(u)
	if c == nil {
		return nil
	}

	g := &DBusObject{native: c}

	return g
}

func (recv *DBusObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObject with another DBusObject, and returns true if they represent the same GObject.
func (recv *DBusObject) Equals(other *DBusObject) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectManager is a wrapper around the C record GDBusObjectManager.
type DBusObjectManager struct {
	native *C.GDBusObjectManager
}

func DBusObjectManagerNewFromC(u unsafe.Pointer) *DBusObjectManager {
	c := (*C.GDBusObjectManager)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManager{native: c}

	return g
}

func (recv *DBusObjectManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManager with another DBusObjectManager, and returns true if they represent the same GObject.
func (recv *DBusObjectManager) Equals(other *DBusObjectManager) bool {
	return other.ToC() == recv.ToC()
}

// DesktopAppInfoLookup is a wrapper around the C record GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native *C.GDesktopAppInfoLookup
}

func DesktopAppInfoLookupNewFromC(u unsafe.Pointer) *DesktopAppInfoLookup {
	c := (*C.GDesktopAppInfoLookup)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookup{native: c}

	return g
}

func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DesktopAppInfoLookup with another DesktopAppInfoLookup, and returns true if they represent the same GObject.
func (recv *DesktopAppInfoLookup) Equals(other *DesktopAppInfoLookup) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_desktop_app_info_lookup_get_default_for_uri_scheme

// Drive is a wrapper around the C record GDrive.
type Drive struct {
	native *C.GDrive
}

func DriveNewFromC(u unsafe.Pointer) *Drive {
	c := (*C.GDrive)(u)
	if c == nil {
		return nil
	}

	g := &Drive{native: c}

	return g
}

func (recv *Drive) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Drive with another Drive, and returns true if they represent the same GObject.
func (recv *Drive) Equals(other *Drive) bool {
	return other.ToC() == recv.ToC()
}

type signalDriveChangedDetail struct {
	callback  DriveSignalChangedCallback
	handlerID C.gulong
}

var signalDriveChangedId int
var signalDriveChangedMap = make(map[int]signalDriveChangedDetail)
var signalDriveChangedLock sync.RWMutex

// DriveSignalChangedCallback is a callback function for a 'changed' signal emitted from a Drive.
type DriveSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Drive) ConnectChanged(callback DriveSignalChangedCallback) int {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	signalDriveChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_changed(instance, C.gpointer(uintptr(signalDriveChangedId)))

	detail := signalDriveChangedDetail{callback, handlerID}
	signalDriveChangedMap[signalDriveChangedId] = detail

	return signalDriveChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Drive.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Drive) DisconnectChanged(connectionID int) {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	detail, exists := signalDriveChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveChangedMap, connectionID)
}

//export drive_changedHandler
func drive_changedHandler(_ *C.GObject, data C.gpointer) {
	signalDriveChangedLock.RLock()
	defer signalDriveChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveChangedMap[index].callback
	callback()
}

type signalDriveDisconnectedDetail struct {
	callback  DriveSignalDisconnectedCallback
	handlerID C.gulong
}

var signalDriveDisconnectedId int
var signalDriveDisconnectedMap = make(map[int]signalDriveDisconnectedDetail)
var signalDriveDisconnectedLock sync.RWMutex

// DriveSignalDisconnectedCallback is a callback function for a 'disconnected' signal emitted from a Drive.
type DriveSignalDisconnectedCallback func()

/*
ConnectDisconnected connects the callback to the 'disconnected' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectDisconnected to remove it.
*/
func (recv *Drive) ConnectDisconnected(callback DriveSignalDisconnectedCallback) int {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	signalDriveDisconnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_disconnected(instance, C.gpointer(uintptr(signalDriveDisconnectedId)))

	detail := signalDriveDisconnectedDetail{callback, handlerID}
	signalDriveDisconnectedMap[signalDriveDisconnectedId] = detail

	return signalDriveDisconnectedId
}

/*
DisconnectDisconnected disconnects a callback from the 'disconnected' signal for the Drive.

The connectionID should be a value returned from a call to ConnectDisconnected.
*/
func (recv *Drive) DisconnectDisconnected(connectionID int) {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	detail, exists := signalDriveDisconnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveDisconnectedMap, connectionID)
}

//export drive_disconnectedHandler
func drive_disconnectedHandler(_ *C.GObject, data C.gpointer) {
	signalDriveDisconnectedLock.RLock()
	defer signalDriveDisconnectedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveDisconnectedMap[index].callback
	callback()
}

type signalDriveEjectButtonDetail struct {
	callback  DriveSignalEjectButtonCallback
	handlerID C.gulong
}

var signalDriveEjectButtonId int
var signalDriveEjectButtonMap = make(map[int]signalDriveEjectButtonDetail)
var signalDriveEjectButtonLock sync.RWMutex

// DriveSignalEjectButtonCallback is a callback function for a 'eject-button' signal emitted from a Drive.
type DriveSignalEjectButtonCallback func()

/*
ConnectEjectButton connects the callback to the 'eject-button' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectEjectButton to remove it.
*/
func (recv *Drive) ConnectEjectButton(callback DriveSignalEjectButtonCallback) int {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	signalDriveEjectButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_eject_button(instance, C.gpointer(uintptr(signalDriveEjectButtonId)))

	detail := signalDriveEjectButtonDetail{callback, handlerID}
	signalDriveEjectButtonMap[signalDriveEjectButtonId] = detail

	return signalDriveEjectButtonId
}

/*
DisconnectEjectButton disconnects a callback from the 'eject-button' signal for the Drive.

The connectionID should be a value returned from a call to ConnectEjectButton.
*/
func (recv *Drive) DisconnectEjectButton(connectionID int) {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	detail, exists := signalDriveEjectButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveEjectButtonMap, connectionID)
}

//export drive_ejectButtonHandler
func drive_ejectButtonHandler(_ *C.GObject, data C.gpointer) {
	signalDriveEjectButtonLock.RLock()
	defer signalDriveEjectButtonLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveEjectButtonMap[index].callback
	callback()
}

// Blacklisted : g_drive_can_eject

// Blacklisted : g_drive_can_poll_for_media

// Unsupported : g_drive_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_drive_eject_finish

// Blacklisted : g_drive_enumerate_identifiers

// Blacklisted : g_drive_get_icon

// Blacklisted : g_drive_get_identifier

// Blacklisted : g_drive_get_name

// Blacklisted : g_drive_get_volumes

// Blacklisted : g_drive_has_media

// Blacklisted : g_drive_has_volumes

// Blacklisted : g_drive_is_media_check_automatic

// Blacklisted : g_drive_is_media_removable

// Unsupported : g_drive_poll_for_media : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_drive_poll_for_media_finish

// File is a wrapper around the C record GFile.
type File struct {
	native *C.GFile
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this File with another File, and returns true if they represent the same GObject.
func (recv *File) Equals(other *File) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_file_new_for_commandline_arg

// Blacklisted : g_file_new_for_path

// Blacklisted : g_file_new_for_uri

// Blacklisted : g_file_parse_name

// Blacklisted : g_file_append_to

// Unsupported : g_file_append_to_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_append_to_finish

// Unsupported : g_file_copy : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Unsupported : g_file_copy_async : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Blacklisted : g_file_copy_attributes

// Blacklisted : g_file_copy_finish

// Blacklisted : g_file_create

// Unsupported : g_file_create_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_create_finish

// Blacklisted : g_file_delete

// Blacklisted : g_file_dup

// Unsupported : g_file_eject_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_eject_mountable_finish

// Blacklisted : g_file_enumerate_children

// Unsupported : g_file_enumerate_children_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_enumerate_children_finish

// Blacklisted : g_file_equal

// Blacklisted : g_file_find_enclosing_mount

// Unsupported : g_file_find_enclosing_mount_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_find_enclosing_mount_finish

// Blacklisted : g_file_get_basename

// Blacklisted : g_file_get_child

// Blacklisted : g_file_get_child_for_display_name

// Blacklisted : g_file_get_parent

// Blacklisted : g_file_get_parse_name

// Blacklisted : g_file_get_path

// Blacklisted : g_file_get_relative_path

// Blacklisted : g_file_get_uri

// Blacklisted : g_file_get_uri_scheme

// Blacklisted : g_file_has_prefix

// Blacklisted : g_file_has_uri_scheme

// Blacklisted : g_file_hash

// Blacklisted : g_file_is_native

// Unsupported : g_file_load_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_load_contents_finish : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_partial_contents_async : unsupported parameter read_more_callback : no type generator for FileReadMoreCallback (GFileReadMoreCallback) for param read_more_callback

// Unsupported : g_file_load_partial_contents_finish : unsupported parameter contents : output array param contents

// Blacklisted : g_file_make_directory

// Blacklisted : g_file_make_symbolic_link

// Blacklisted : g_file_monitor_directory

// Blacklisted : g_file_monitor_file

// Unsupported : g_file_mount_enclosing_volume : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_mount_enclosing_volume_finish

// Unsupported : g_file_mount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_mount_mountable_finish

// Unsupported : g_file_move : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Blacklisted : g_file_query_default_handler

// Blacklisted : g_file_query_exists

// Blacklisted : g_file_query_filesystem_info

// Unsupported : g_file_query_filesystem_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_query_filesystem_info_finish

// Blacklisted : g_file_query_info

// Unsupported : g_file_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_query_info_finish

// Blacklisted : g_file_query_settable_attributes

// Blacklisted : g_file_query_writable_namespaces

// Blacklisted : g_file_read

// Unsupported : g_file_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_read_finish

// Blacklisted : g_file_replace

// Unsupported : g_file_replace_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_replace_contents

// Unsupported : g_file_replace_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_replace_contents_finish

// Blacklisted : g_file_replace_finish

// Blacklisted : g_file_resolve_relative_path

// Unsupported : g_file_set_attribute : unsupported parameter value_p : no type generator for gpointer (gpointer) for param value_p

// Blacklisted : g_file_set_attribute_byte_string

// Blacklisted : g_file_set_attribute_int32

// Blacklisted : g_file_set_attribute_int64

// Blacklisted : g_file_set_attribute_string

// Blacklisted : g_file_set_attribute_uint32

// Blacklisted : g_file_set_attribute_uint64

// Unsupported : g_file_set_attributes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_set_attributes_finish

// Blacklisted : g_file_set_attributes_from_info

// Blacklisted : g_file_set_display_name

// Unsupported : g_file_set_display_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_set_display_name_finish

// Blacklisted : g_file_trash

// Unsupported : g_file_unmount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_unmount_mountable_finish

// FileDescriptorBased is a wrapper around the C record GFileDescriptorBased.
type FileDescriptorBased struct {
	native *C.GFileDescriptorBased
}

func FileDescriptorBasedNewFromC(u unsafe.Pointer) *FileDescriptorBased {
	c := (*C.GFileDescriptorBased)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBased{native: c}

	return g
}

func (recv *FileDescriptorBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileDescriptorBased with another FileDescriptorBased, and returns true if they represent the same GObject.
func (recv *FileDescriptorBased) Equals(other *FileDescriptorBased) bool {
	return other.ToC() == recv.ToC()
}

// Icon is a wrapper around the C record GIcon.
type Icon struct {
	native *C.GIcon
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	c := (*C.GIcon)(u)
	if c == nil {
		return nil
	}

	g := &Icon{native: c}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Icon with another Icon, and returns true if they represent the same GObject.
func (recv *Icon) Equals(other *Icon) bool {
	return other.ToC() == recv.ToC()
}

// g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon
// Blacklisted : g_icon_equal

// ListModel is a wrapper around the C record GListModel.
type ListModel struct {
	native *C.GListModel
}

func ListModelNewFromC(u unsafe.Pointer) *ListModel {
	c := (*C.GListModel)(u)
	if c == nil {
		return nil
	}

	g := &ListModel{native: c}

	return g
}

func (recv *ListModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ListModel with another ListModel, and returns true if they represent the same GObject.
func (recv *ListModel) Equals(other *ListModel) bool {
	return other.ToC() == recv.ToC()
}

// LoadableIcon is a wrapper around the C record GLoadableIcon.
type LoadableIcon struct {
	native *C.GLoadableIcon
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	c := (*C.GLoadableIcon)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIcon{native: c}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LoadableIcon with another LoadableIcon, and returns true if they represent the same GObject.
func (recv *LoadableIcon) Equals(other *LoadableIcon) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_loadable_icon_load

// Unsupported : g_loadable_icon_load_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_loadable_icon_load_finish

// Mount is a wrapper around the C record GMount.
type Mount struct {
	native *C.GMount
}

func MountNewFromC(u unsafe.Pointer) *Mount {
	c := (*C.GMount)(u)
	if c == nil {
		return nil
	}

	g := &Mount{native: c}

	return g
}

func (recv *Mount) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Mount with another Mount, and returns true if they represent the same GObject.
func (recv *Mount) Equals(other *Mount) bool {
	return other.ToC() == recv.ToC()
}

type signalMountChangedDetail struct {
	callback  MountSignalChangedCallback
	handlerID C.gulong
}

var signalMountChangedId int
var signalMountChangedMap = make(map[int]signalMountChangedDetail)
var signalMountChangedLock sync.RWMutex

// MountSignalChangedCallback is a callback function for a 'changed' signal emitted from a Mount.
type MountSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Mount) ConnectChanged(callback MountSignalChangedCallback) int {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	signalMountChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_changed(instance, C.gpointer(uintptr(signalMountChangedId)))

	detail := signalMountChangedDetail{callback, handlerID}
	signalMountChangedMap[signalMountChangedId] = detail

	return signalMountChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Mount.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Mount) DisconnectChanged(connectionID int) {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	detail, exists := signalMountChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountChangedMap, connectionID)
}

//export mount_changedHandler
func mount_changedHandler(_ *C.GObject, data C.gpointer) {
	signalMountChangedLock.RLock()
	defer signalMountChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountChangedMap[index].callback
	callback()
}

type signalMountUnmountedDetail struct {
	callback  MountSignalUnmountedCallback
	handlerID C.gulong
}

var signalMountUnmountedId int
var signalMountUnmountedMap = make(map[int]signalMountUnmountedDetail)
var signalMountUnmountedLock sync.RWMutex

// MountSignalUnmountedCallback is a callback function for a 'unmounted' signal emitted from a Mount.
type MountSignalUnmountedCallback func()

/*
ConnectUnmounted connects the callback to the 'unmounted' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectUnmounted to remove it.
*/
func (recv *Mount) ConnectUnmounted(callback MountSignalUnmountedCallback) int {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	signalMountUnmountedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_unmounted(instance, C.gpointer(uintptr(signalMountUnmountedId)))

	detail := signalMountUnmountedDetail{callback, handlerID}
	signalMountUnmountedMap[signalMountUnmountedId] = detail

	return signalMountUnmountedId
}

/*
DisconnectUnmounted disconnects a callback from the 'unmounted' signal for the Mount.

The connectionID should be a value returned from a call to ConnectUnmounted.
*/
func (recv *Mount) DisconnectUnmounted(connectionID int) {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	detail, exists := signalMountUnmountedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountUnmountedMap, connectionID)
}

//export mount_unmountedHandler
func mount_unmountedHandler(_ *C.GObject, data C.gpointer) {
	signalMountUnmountedLock.RLock()
	defer signalMountUnmountedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountUnmountedMap[index].callback
	callback()
}

// Blacklisted : g_mount_can_eject

// Blacklisted : g_mount_can_unmount

// Unsupported : g_mount_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_mount_eject_finish

// Blacklisted : g_mount_get_default_location

// Blacklisted : g_mount_get_drive

// Blacklisted : g_mount_get_icon

// Blacklisted : g_mount_get_name

// Blacklisted : g_mount_get_root

// Blacklisted : g_mount_get_uuid

// Blacklisted : g_mount_get_volume

// Unsupported : g_mount_remount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_mount_remount_finish

// Unsupported : g_mount_unmount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_mount_unmount_finish

// RemoteActionGroup is a wrapper around the C record GRemoteActionGroup.
type RemoteActionGroup struct {
	native *C.GRemoteActionGroup
}

func RemoteActionGroupNewFromC(u unsafe.Pointer) *RemoteActionGroup {
	c := (*C.GRemoteActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroup{native: c}

	return g
}

func (recv *RemoteActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RemoteActionGroup with another RemoteActionGroup, and returns true if they represent the same GObject.
func (recv *RemoteActionGroup) Equals(other *RemoteActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Seekable is a wrapper around the C record GSeekable.
type Seekable struct {
	native *C.GSeekable
}

func SeekableNewFromC(u unsafe.Pointer) *Seekable {
	c := (*C.GSeekable)(u)
	if c == nil {
		return nil
	}

	g := &Seekable{native: c}

	return g
}

func (recv *Seekable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Seekable with another Seekable, and returns true if they represent the same GObject.
func (recv *Seekable) Equals(other *Seekable) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_seekable_can_seek

// Blacklisted : g_seekable_can_truncate

// Blacklisted : g_seekable_seek

// Blacklisted : g_seekable_tell

// Blacklisted : g_seekable_truncate

// SocketConnectable is a wrapper around the C record GSocketConnectable.
type SocketConnectable struct {
	native *C.GSocketConnectable
}

func SocketConnectableNewFromC(u unsafe.Pointer) *SocketConnectable {
	c := (*C.GSocketConnectable)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectable{native: c}

	return g
}

func (recv *SocketConnectable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnectable with another SocketConnectable, and returns true if they represent the same GObject.
func (recv *SocketConnectable) Equals(other *SocketConnectable) bool {
	return other.ToC() == recv.ToC()
}

// Volume is a wrapper around the C record GVolume.
type Volume struct {
	native *C.GVolume
}

func VolumeNewFromC(u unsafe.Pointer) *Volume {
	c := (*C.GVolume)(u)
	if c == nil {
		return nil
	}

	g := &Volume{native: c}

	return g
}

func (recv *Volume) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Volume with another Volume, and returns true if they represent the same GObject.
func (recv *Volume) Equals(other *Volume) bool {
	return other.ToC() == recv.ToC()
}

type signalVolumeChangedDetail struct {
	callback  VolumeSignalChangedCallback
	handlerID C.gulong
}

var signalVolumeChangedId int
var signalVolumeChangedMap = make(map[int]signalVolumeChangedDetail)
var signalVolumeChangedLock sync.RWMutex

// VolumeSignalChangedCallback is a callback function for a 'changed' signal emitted from a Volume.
type VolumeSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Volume) ConnectChanged(callback VolumeSignalChangedCallback) int {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	signalVolumeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_changed(instance, C.gpointer(uintptr(signalVolumeChangedId)))

	detail := signalVolumeChangedDetail{callback, handlerID}
	signalVolumeChangedMap[signalVolumeChangedId] = detail

	return signalVolumeChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Volume) DisconnectChanged(connectionID int) {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	detail, exists := signalVolumeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeChangedMap, connectionID)
}

//export volume_changedHandler
func volume_changedHandler(_ *C.GObject, data C.gpointer) {
	signalVolumeChangedLock.RLock()
	defer signalVolumeChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalVolumeChangedMap[index].callback
	callback()
}

type signalVolumeRemovedDetail struct {
	callback  VolumeSignalRemovedCallback
	handlerID C.gulong
}

var signalVolumeRemovedId int
var signalVolumeRemovedMap = make(map[int]signalVolumeRemovedDetail)
var signalVolumeRemovedLock sync.RWMutex

// VolumeSignalRemovedCallback is a callback function for a 'removed' signal emitted from a Volume.
type VolumeSignalRemovedCallback func()

/*
ConnectRemoved connects the callback to the 'removed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectRemoved to remove it.
*/
func (recv *Volume) ConnectRemoved(callback VolumeSignalRemovedCallback) int {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	signalVolumeRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_removed(instance, C.gpointer(uintptr(signalVolumeRemovedId)))

	detail := signalVolumeRemovedDetail{callback, handlerID}
	signalVolumeRemovedMap[signalVolumeRemovedId] = detail

	return signalVolumeRemovedId
}

/*
DisconnectRemoved disconnects a callback from the 'removed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectRemoved.
*/
func (recv *Volume) DisconnectRemoved(connectionID int) {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	detail, exists := signalVolumeRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeRemovedMap, connectionID)
}

//export volume_removedHandler
func volume_removedHandler(_ *C.GObject, data C.gpointer) {
	signalVolumeRemovedLock.RLock()
	defer signalVolumeRemovedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalVolumeRemovedMap[index].callback
	callback()
}

// Blacklisted : g_volume_can_eject

// Blacklisted : g_volume_can_mount

// Unsupported : g_volume_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_volume_eject_finish

// Blacklisted : g_volume_enumerate_identifiers

// Blacklisted : g_volume_get_drive

// Blacklisted : g_volume_get_icon

// Blacklisted : g_volume_get_identifier

// Blacklisted : g_volume_get_mount

// Blacklisted : g_volume_get_name

// Blacklisted : g_volume_get_uuid

// Unsupported : g_volume_mount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_volume_mount_finish

// Blacklisted : g_volume_should_automount
