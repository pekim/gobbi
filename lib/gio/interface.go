// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

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

// AddSupportsType is a wrapper around the C function g_app_info_add_supports_type.
func (recv *AppInfo) AddSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_add_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// CanRemoveSupportsType is a wrapper around the C function g_app_info_can_remove_supports_type.
func (recv *AppInfo) CanRemoveSupportsType() bool {
	retC := C.g_app_info_can_remove_supports_type((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Dup is a wrapper around the C function g_app_info_dup.
func (recv *AppInfo) Dup() *AppInfo {
	retC := C.g_app_info_dup((*C.GAppInfo)(recv.native))
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_app_info_equal.
func (recv *AppInfo) Equal(appinfo2 *AppInfo) bool {
	c_appinfo2 := (*C.GAppInfo)(appinfo2.ToC())

	retC := C.g_app_info_equal((*C.GAppInfo)(recv.native), c_appinfo2)
	retGo := retC == C.TRUE

	return retGo
}

// GetDescription is a wrapper around the C function g_app_info_get_description.
func (recv *AppInfo) GetDescription() string {
	retC := C.g_app_info_get_description((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetExecutable is a wrapper around the C function g_app_info_get_executable.
func (recv *AppInfo) GetExecutable() string {
	retC := C.g_app_info_get_executable((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIcon is a wrapper around the C function g_app_info_get_icon.
func (recv *AppInfo) GetIcon() *Icon {
	retC := C.g_app_info_get_icon((*C.GAppInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetId is a wrapper around the C function g_app_info_get_id.
func (recv *AppInfo) GetId() string {
	retC := C.g_app_info_get_id((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_app_info_get_name.
func (recv *AppInfo) GetName() string {
	retC := C.g_app_info_get_name((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Launch is a wrapper around the C function g_app_info_launch.
func (recv *AppInfo) Launch(files *glib.List, context *AppLaunchContext) (bool, error) {
	c_files := (*C.GList)(files.ToC())

	c_context := (*C.GAppLaunchContext)(context.ToC())

	var cThrowableError *C.GError

	retC := C.g_app_info_launch((*C.GAppInfo)(recv.native), c_files, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LaunchUris is a wrapper around the C function g_app_info_launch_uris.
func (recv *AppInfo) LaunchUris(uris *glib.List, context *AppLaunchContext) (bool, error) {
	c_uris := (*C.GList)(uris.ToC())

	c_context := (*C.GAppLaunchContext)(context.ToC())

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_uris((*C.GAppInfo)(recv.native), c_uris, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveSupportsType is a wrapper around the C function g_app_info_remove_supports_type.
func (recv *AppInfo) RemoveSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_remove_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAsDefaultForExtension is a wrapper around the C function g_app_info_set_as_default_for_extension.
func (recv *AppInfo) SetAsDefaultForExtension(extension string) (bool, error) {
	c_extension := C.CString(extension)
	defer C.free(unsafe.Pointer(c_extension))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_extension((*C.GAppInfo)(recv.native), c_extension, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAsDefaultForType is a wrapper around the C function g_app_info_set_as_default_for_type.
func (recv *AppInfo) SetAsDefaultForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAsLastUsedForType is a wrapper around the C function g_app_info_set_as_last_used_for_type.
func (recv *AppInfo) SetAsLastUsedForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_last_used_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ShouldShow is a wrapper around the C function g_app_info_should_show.
func (recv *AppInfo) ShouldShow() bool {
	retC := C.g_app_info_should_show((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsFiles is a wrapper around the C function g_app_info_supports_files.
func (recv *AppInfo) SupportsFiles() bool {
	retC := C.g_app_info_supports_files((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsUris is a wrapper around the C function g_app_info_supports_uris.
func (recv *AppInfo) SupportsUris() bool {
	retC := C.g_app_info_supports_uris((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// GetSourceObject is a wrapper around the C function g_async_result_get_source_object.
func (recv *AsyncResult) GetSourceObject() *gobject.Object {
	retC := C.g_async_result_get_source_object((*C.GAsyncResult)(recv.native))
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetUserData is a wrapper around the C function g_async_result_get_user_data.
func (recv *AsyncResult) GetUserData() uintptr {
	retC := C.g_async_result_get_user_data((*C.GAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// GetDefaultForUriScheme is a wrapper around the C function g_desktop_app_info_lookup_get_default_for_uri_scheme.
func (recv *DesktopAppInfoLookup) GetDefaultForUriScheme(uriScheme string) *AppInfo {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_desktop_app_info_lookup_get_default_for_uri_scheme((*C.GDesktopAppInfoLookup)(recv.native), c_uri_scheme)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

type signalDriveChangedDetail struct {
	callback  DriveSignalChangedCallback
	handlerID C.gulong
}

var signalDriveChangedId int
var signalDriveChangedMap = make(map[int]signalDriveChangedDetail)
var signalDriveChangedLock sync.Mutex

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
var signalDriveDisconnectedLock sync.Mutex

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
var signalDriveEjectButtonLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalDriveEjectButtonMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_drive_can_eject.
func (recv *Drive) CanEject() bool {
	retC := C.g_drive_can_eject((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanPollForMedia is a wrapper around the C function g_drive_can_poll_for_media.
func (recv *Drive) CanPollForMedia() bool {
	retC := C.g_drive_can_poll_for_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_drive_eject_finish.
func (recv *Drive) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_drive_enumerate_identifiers : no return type

// GetIcon is a wrapper around the C function g_drive_get_icon.
func (recv *Drive) GetIcon() *Icon {
	retC := C.g_drive_get_icon((*C.GDrive)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIdentifier is a wrapper around the C function g_drive_get_identifier.
func (recv *Drive) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_drive_get_identifier((*C.GDrive)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_drive_get_name.
func (recv *Drive) GetName() string {
	retC := C.g_drive_get_name((*C.GDrive)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetVolumes is a wrapper around the C function g_drive_get_volumes.
func (recv *Drive) GetVolumes() *glib.List {
	retC := C.g_drive_get_volumes((*C.GDrive)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasMedia is a wrapper around the C function g_drive_has_media.
func (recv *Drive) HasMedia() bool {
	retC := C.g_drive_has_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasVolumes is a wrapper around the C function g_drive_has_volumes.
func (recv *Drive) HasVolumes() bool {
	retC := C.g_drive_has_volumes((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMediaCheckAutomatic is a wrapper around the C function g_drive_is_media_check_automatic.
func (recv *Drive) IsMediaCheckAutomatic() bool {
	retC := C.g_drive_is_media_check_automatic((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMediaRemovable is a wrapper around the C function g_drive_is_media_removable.
func (recv *Drive) IsMediaRemovable() bool {
	retC := C.g_drive_is_media_removable((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_poll_for_media : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// PollForMediaFinish is a wrapper around the C function g_drive_poll_for_media_finish.
func (recv *Drive) PollForMediaFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_poll_for_media_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// AppendTo is a wrapper around the C function g_file_append_to.
func (recv *File) AppendTo(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_append_to((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_append_to_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AppendToFinish is a wrapper around the C function g_file_append_to_finish.
func (recv *File) AppendToFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_append_to_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_copy : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Unsupported : g_file_copy_async : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// CopyAttributes is a wrapper around the C function g_file_copy_attributes.
func (recv *File) CopyAttributes(destination *File, flags FileCopyFlags, cancellable *Cancellable) (bool, error) {
	c_destination := (*C.GFile)(destination.ToC())

	c_flags := (C.GFileCopyFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_copy_attributes((*C.GFile)(recv.native), c_destination, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// CopyFinish is a wrapper around the C function g_file_copy_finish.
func (recv *File) CopyFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_copy_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Create is a wrapper around the C function g_file_create.
func (recv *File) Create(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_create_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CreateFinish is a wrapper around the C function g_file_create_finish.
func (recv *File) CreateFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Delete is a wrapper around the C function g_file_delete.
func (recv *File) Delete(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_delete((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Dup is a wrapper around the C function g_file_dup.
func (recv *File) Dup() *File {
	retC := C.g_file_dup((*C.GFile)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_eject_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectMountableFinish is a wrapper around the C function g_file_eject_mountable_finish.
func (recv *File) EjectMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// EnumerateChildren is a wrapper around the C function g_file_enumerate_children.
func (recv *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileEnumerator, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_enumerate_children_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EnumerateChildrenFinish is a wrapper around the C function g_file_enumerate_children_finish.
func (recv *File) EnumerateChildrenFinish(res *AsyncResult) (*FileEnumerator, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Equal is a wrapper around the C function g_file_equal.
func (recv *File) Equal(file2 *File) bool {
	c_file2 := (*C.GFile)(file2.ToC())

	retC := C.g_file_equal((*C.GFile)(recv.native), c_file2)
	retGo := retC == C.TRUE

	return retGo
}

// FindEnclosingMount is a wrapper around the C function g_file_find_enclosing_mount.
func (recv *File) FindEnclosingMount(cancellable *Cancellable) (*Mount, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_find_enclosing_mount_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// FindEnclosingMountFinish is a wrapper around the C function g_file_find_enclosing_mount_finish.
func (recv *File) FindEnclosingMountFinish(res *AsyncResult) (*Mount, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetBasename is a wrapper around the C function g_file_get_basename.
func (recv *File) GetBasename() string {
	retC := C.g_file_get_basename((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetChild is a wrapper around the C function g_file_get_child.
func (recv *File) GetChild(name string) *File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_file_get_child((*C.GFile)(recv.native), c_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildForDisplayName is a wrapper around the C function g_file_get_child_for_display_name.
func (recv *File) GetChildForDisplayName(displayName string) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	var cThrowableError *C.GError

	retC := C.g_file_get_child_for_display_name((*C.GFile)(recv.native), c_display_name, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetParent is a wrapper around the C function g_file_get_parent.
func (recv *File) GetParent() *File {
	retC := C.g_file_get_parent((*C.GFile)(recv.native))
	var retGo (*File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetParseName is a wrapper around the C function g_file_get_parse_name.
func (recv *File) GetParseName() string {
	retC := C.g_file_get_parse_name((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPath is a wrapper around the C function g_file_get_path.
func (recv *File) GetPath() string {
	retC := C.g_file_get_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetRelativePath is a wrapper around the C function g_file_get_relative_path.
func (recv *File) GetRelativePath(descendant *File) string {
	c_descendant := (*C.GFile)(descendant.ToC())

	retC := C.g_file_get_relative_path((*C.GFile)(recv.native), c_descendant)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function g_file_get_uri.
func (recv *File) GetUri() string {
	retC := C.g_file_get_uri((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUriScheme is a wrapper around the C function g_file_get_uri_scheme.
func (recv *File) GetUriScheme() string {
	retC := C.g_file_get_uri_scheme((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HasPrefix is a wrapper around the C function g_file_has_prefix.
func (recv *File) HasPrefix(prefix *File) bool {
	c_prefix := (*C.GFile)(prefix.ToC())

	retC := C.g_file_has_prefix((*C.GFile)(recv.native), c_prefix)
	retGo := retC == C.TRUE

	return retGo
}

// HasUriScheme is a wrapper around the C function g_file_has_uri_scheme.
func (recv *File) HasUriScheme(uriScheme string) bool {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_file_has_uri_scheme((*C.GFile)(recv.native), c_uri_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// Hash is a wrapper around the C function g_file_hash.
func (recv *File) Hash() uint32 {
	retC := C.g_file_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsNative is a wrapper around the C function g_file_is_native.
func (recv *File) IsNative() bool {
	retC := C.g_file_is_native((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_load_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_load_contents_finish : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_partial_contents_async : unsupported parameter read_more_callback : no type generator for FileReadMoreCallback (GFileReadMoreCallback) for param read_more_callback

// Unsupported : g_file_load_partial_contents_finish : unsupported parameter contents : output array param contents

// MakeDirectory is a wrapper around the C function g_file_make_directory.
func (recv *File) MakeDirectory(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_make_directory((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MakeSymbolicLink is a wrapper around the C function g_file_make_symbolic_link.
func (recv *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable) (bool, error) {
	c_symlink_value := C.CString(symlinkValue)
	defer C.free(unsafe.Pointer(c_symlink_value))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_make_symbolic_link((*C.GFile)(recv.native), c_symlink_value, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MonitorDirectory is a wrapper around the C function g_file_monitor_directory.
func (recv *File) MonitorDirectory(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_monitor_directory((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MonitorFile is a wrapper around the C function g_file_monitor_file.
func (recv *File) MonitorFile(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_monitor_file((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_mount_enclosing_volume : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountEnclosingVolumeFinish is a wrapper around the C function g_file_mount_enclosing_volume_finish.
func (recv *File) MountEnclosingVolumeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_enclosing_volume_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_mount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountMountableFinish is a wrapper around the C function g_file_mount_mountable_finish.
func (recv *File) MountMountableFinish(result *AsyncResult) (*File, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_move : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// QueryDefaultHandler is a wrapper around the C function g_file_query_default_handler.
func (recv *File) QueryDefaultHandler(cancellable *Cancellable) (*AppInfo, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_default_handler((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// QueryExists is a wrapper around the C function g_file_query_exists.
func (recv *File) QueryExists(cancellable *Cancellable) bool {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_file_query_exists((*C.GFile)(recv.native), c_cancellable)
	retGo := retC == C.TRUE

	return retGo
}

// QueryFilesystemInfo is a wrapper around the C function g_file_query_filesystem_info.
func (recv *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info((*C.GFile)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_query_filesystem_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryFilesystemInfoFinish is a wrapper around the C function g_file_query_filesystem_info_finish.
func (recv *File) QueryFilesystemInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// QueryInfo is a wrapper around the C function g_file_query_info.
func (recv *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_info((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryInfoFinish is a wrapper around the C function g_file_query_info_finish.
func (recv *File) QueryInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// QuerySettableAttributes is a wrapper around the C function g_file_query_settable_attributes.
func (recv *File) QuerySettableAttributes(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_settable_attributes((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// QueryWritableNamespaces is a wrapper around the C function g_file_query_writable_namespaces.
func (recv *File) QueryWritableNamespaces(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_writable_namespaces((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Read is a wrapper around the C function g_file_read.
func (recv *File) Read(cancellable *Cancellable) (*FileInputStream, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_read((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadFinish is a wrapper around the C function g_file_read_finish.
func (recv *File) ReadFinish(res *AsyncResult) (*FileInputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_read_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Replace is a wrapper around the C function g_file_replace.
func (recv *File) Replace(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace((*C.GFile)(recv.native), c_etag, c_make_backup, c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_replace_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceContents is a wrapper around the C function g_file_replace_contents.
func (recv *File) ReplaceContents(contents []uint8, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (bool, string, error) {
	c_contents := &contents[0]

	c_length := (C.gsize)(len(contents))

	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	var c_new_etag *C.char

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents((*C.GFile)(recv.native), (*C.char)(unsafe.Pointer(c_contents)), c_length, c_etag, c_make_backup, c_flags, &c_new_etag, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goThrowableError
}

// Unsupported : g_file_replace_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceContentsFinish is a wrapper around the C function g_file_replace_contents_finish.
func (recv *File) ReplaceContentsFinish(res *AsyncResult) (bool, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_new_etag *C.char

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents_finish((*C.GFile)(recv.native), c_res, &c_new_etag, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goThrowableError
}

// ReplaceFinish is a wrapper around the C function g_file_replace_finish.
func (recv *File) ReplaceFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ResolveRelativePath is a wrapper around the C function g_file_resolve_relative_path.
func (recv *File) ResolveRelativePath(relativePath string) *File {
	c_relative_path := C.CString(relativePath)
	defer C.free(unsafe.Pointer(c_relative_path))

	retC := C.g_file_resolve_relative_path((*C.GFile)(recv.native), c_relative_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAttribute is a wrapper around the C function g_file_set_attribute.
func (recv *File) SetAttribute(attribute string, type_ FileAttributeType, valueP uintptr, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_type := (C.GFileAttributeType)(type_)

	c_value_p := (C.gpointer)(valueP)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute((*C.GFile)(recv.native), c_attribute, c_type, c_value_p, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeByteString is a wrapper around the C function g_file_set_attribute_byte_string.
func (recv *File) SetAttributeByteString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_byte_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeInt32 is a wrapper around the C function g_file_set_attribute_int32.
func (recv *File) SetAttributeInt32(attribute string, value int32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeInt64 is a wrapper around the C function g_file_set_attribute_int64.
func (recv *File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeString is a wrapper around the C function g_file_set_attribute_string.
func (recv *File) SetAttributeString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeUint32 is a wrapper around the C function g_file_set_attribute_uint32.
func (recv *File) SetAttributeUint32(attribute string, value uint32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAttributeUint64 is a wrapper around the C function g_file_set_attribute_uint64.
func (recv *File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_set_attributes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_set_attributes_finish : unsupported parameter info : record with indirection level of 2

// SetAttributesFromInfo is a wrapper around the C function g_file_set_attributes_from_info.
func (recv *File) SetAttributesFromInfo(info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_info := (*C.GFileInfo)(info.ToC())

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_attributes_from_info((*C.GFile)(recv.native), c_info, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetDisplayName is a wrapper around the C function g_file_set_display_name.
func (recv *File) SetDisplayName(displayName string, cancellable *Cancellable) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name((*C.GFile)(recv.native), c_display_name, c_cancellable, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_set_display_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SetDisplayNameFinish is a wrapper around the C function g_file_set_display_name_finish.
func (recv *File) SetDisplayNameFinish(res *AsyncResult) (*File, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Trash is a wrapper around the C function g_file_trash.
func (recv *File) Trash(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_trash((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_unmount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountMountableFinish is a wrapper around the C function g_file_unmount_mountable_finish.
func (recv *File) UnmountMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// Equal is a wrapper around the C function g_icon_equal.
func (recv *Icon) Equal(icon2 *Icon) bool {
	c_icon2 := (*C.GIcon)(icon2.ToC())

	retC := C.g_icon_equal((*C.GIcon)(recv.native), c_icon2)
	retGo := retC == C.TRUE

	return retGo
}

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

// Load is a wrapper around the C function g_loadable_icon_load.
func (recv *LoadableIcon) Load(size int32, cancellable *Cancellable) (*InputStream, string, error) {
	c_size := (C.int)(size)

	var c_type *C.char

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load((*C.GLoadableIcon)(recv.native), c_size, &c_type, c_cancellable, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goThrowableError
}

// Unsupported : g_loadable_icon_load_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadFinish is a wrapper around the C function g_loadable_icon_load_finish.
func (recv *LoadableIcon) LoadFinish(res *AsyncResult) (*InputStream, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_type *C.char

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load_finish((*C.GLoadableIcon)(recv.native), c_res, &c_type, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goThrowableError
}

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

type signalMountChangedDetail struct {
	callback  MountSignalChangedCallback
	handlerID C.gulong
}

var signalMountChangedId int
var signalMountChangedMap = make(map[int]signalMountChangedDetail)
var signalMountChangedLock sync.Mutex

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
var signalMountUnmountedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalMountUnmountedMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_mount_can_eject.
func (recv *Mount) CanEject() bool {
	retC := C.g_mount_can_eject((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanUnmount is a wrapper around the C function g_mount_can_unmount.
func (recv *Mount) CanUnmount() bool {
	retC := C.g_mount_can_unmount((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_mount_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_mount_eject_finish.
func (recv *Mount) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetDefaultLocation is a wrapper around the C function g_mount_get_default_location.
func (recv *Mount) GetDefaultLocation() *File {
	retC := C.g_mount_get_default_location((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDrive is a wrapper around the C function g_mount_get_drive.
func (recv *Mount) GetDrive() *Drive {
	retC := C.g_mount_get_drive((*C.GMount)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIcon is a wrapper around the C function g_mount_get_icon.
func (recv *Mount) GetIcon() *Icon {
	retC := C.g_mount_get_icon((*C.GMount)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_mount_get_name.
func (recv *Mount) GetName() string {
	retC := C.g_mount_get_name((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function g_mount_get_root.
func (recv *Mount) GetRoot() *File {
	retC := C.g_mount_get_root((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUuid is a wrapper around the C function g_mount_get_uuid.
func (recv *Mount) GetUuid() string {
	retC := C.g_mount_get_uuid((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetVolume is a wrapper around the C function g_mount_get_volume.
func (recv *Mount) GetVolume() *Volume {
	retC := C.g_mount_get_volume((*C.GMount)(recv.native))
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_mount_remount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// RemountFinish is a wrapper around the C function g_mount_remount_finish.
func (recv *Mount) RemountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_remount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_mount_unmount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountFinish is a wrapper around the C function g_mount_unmount_finish.
func (recv *Mount) UnmountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// CanSeek is a wrapper around the C function g_seekable_can_seek.
func (recv *Seekable) CanSeek() bool {
	retC := C.g_seekable_can_seek((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanTruncate is a wrapper around the C function g_seekable_can_truncate.
func (recv *Seekable) CanTruncate() bool {
	retC := C.g_seekable_can_truncate((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Seek is a wrapper around the C function g_seekable_seek.
func (recv *Seekable) Seek(offset uint64, type_ glib.SeekType, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_type := (C.GSeekType)(type_)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_seekable_seek((*C.GSeekable)(recv.native), c_offset, c_type, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tell is a wrapper around the C function g_seekable_tell.
func (recv *Seekable) Tell() uint64 {
	retC := C.g_seekable_tell((*C.GSeekable)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Truncate is a wrapper around the C function g_seekable_truncate.
func (recv *Seekable) Truncate(offset uint64, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_seekable_truncate((*C.GSeekable)(recv.native), c_offset, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

type signalVolumeChangedDetail struct {
	callback  VolumeSignalChangedCallback
	handlerID C.gulong
}

var signalVolumeChangedId int
var signalVolumeChangedMap = make(map[int]signalVolumeChangedDetail)
var signalVolumeChangedLock sync.Mutex

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
var signalVolumeRemovedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalVolumeRemovedMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_volume_can_eject.
func (recv *Volume) CanEject() bool {
	retC := C.g_volume_can_eject((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanMount is a wrapper around the C function g_volume_can_mount.
func (recv *Volume) CanMount() bool {
	retC := C.g_volume_can_mount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_volume_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_volume_eject_finish.
func (recv *Volume) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_volume_enumerate_identifiers : no return type

// GetDrive is a wrapper around the C function g_volume_get_drive.
func (recv *Volume) GetDrive() *Drive {
	retC := C.g_volume_get_drive((*C.GVolume)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIcon is a wrapper around the C function g_volume_get_icon.
func (recv *Volume) GetIcon() *Icon {
	retC := C.g_volume_get_icon((*C.GVolume)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIdentifier is a wrapper around the C function g_volume_get_identifier.
func (recv *Volume) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_volume_get_identifier((*C.GVolume)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetMount is a wrapper around the C function g_volume_get_mount.
func (recv *Volume) GetMount() *Mount {
	retC := C.g_volume_get_mount((*C.GVolume)(recv.native))
	retGo := MountNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_volume_get_name.
func (recv *Volume) GetName() string {
	retC := C.g_volume_get_name((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUuid is a wrapper around the C function g_volume_get_uuid.
func (recv *Volume) GetUuid() string {
	retC := C.g_volume_get_uuid((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_mount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountFinish is a wrapper around the C function g_volume_mount_finish.
func (recv *Volume) MountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_mount_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ShouldAutomount is a wrapper around the C function g_volume_should_automount.
func (recv *Volume) ShouldAutomount() bool {
	retC := C.g_volume_should_automount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
