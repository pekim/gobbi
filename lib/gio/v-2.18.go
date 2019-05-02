// Code generated - DO NOT EDIT.
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

	void volumemonitor_driveEjectButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-eject-button", G_CALLBACK(volumemonitor_driveEjectButtonHandler), data);
	}

*/
import "C"

// DesktopAppInfoNewFromKeyfile is a wrapper around the C function g_desktop_app_info_new_from_keyfile.
func DesktopAppInfoNewFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	retC := C.g_desktop_app_info_new_from_keyfile(c_key_file)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EmblemNew is a wrapper around the C function g_emblem_new.
func EmblemNew(icon *Icon) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.g_emblem_new(c_icon)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EmblemNewWithOrigin is a wrapper around the C function g_emblem_new_with_origin.
func EmblemNewWithOrigin(icon *Icon, origin EmblemOrigin) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	c_origin := (C.GEmblemOrigin)(origin)

	retC := C.g_emblem_new_with_origin(c_icon, c_origin)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetIcon is a wrapper around the C function g_emblem_get_icon.
func (recv *Emblem) GetIcon() *Icon {
	retC := C.g_emblem_get_icon((*C.GEmblem)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOrigin is a wrapper around the C function g_emblem_get_origin.
func (recv *Emblem) GetOrigin() EmblemOrigin {
	retC := C.g_emblem_get_origin((*C.GEmblem)(recv.native))
	retGo := (EmblemOrigin)(retC)

	return retGo
}

// EmblemedIconNew is a wrapper around the C function g_emblemed_icon_new.
func EmblemedIconNew(icon *Icon, emblem *Emblem) *EmblemedIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	c_emblem := (*C.GEmblem)(C.NULL)
	if emblem != nil {
		c_emblem = (*C.GEmblem)(emblem.ToC())
	}

	retC := C.g_emblemed_icon_new(c_icon, c_emblem)
	retGo := EmblemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddEmblem is a wrapper around the C function g_emblemed_icon_add_emblem.
func (recv *EmblemedIcon) AddEmblem(emblem *Emblem) {
	c_emblem := (*C.GEmblem)(C.NULL)
	if emblem != nil {
		c_emblem = (*C.GEmblem)(emblem.ToC())
	}

	C.g_emblemed_icon_add_emblem((*C.GEmblemedIcon)(recv.native), c_emblem)

	return
}

// GetEmblems is a wrapper around the C function g_emblemed_icon_get_emblems.
func (recv *EmblemedIcon) GetEmblems() *glib.List {
	retC := C.g_emblemed_icon_get_emblems((*C.GEmblemedIcon)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIcon is a wrapper around the C function g_emblemed_icon_get_icon.
func (recv *EmblemedIcon) GetIcon() *Icon {
	retC := C.g_emblemed_icon_get_icon((*C.GEmblemedIcon)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContainer is a wrapper around the C function g_file_enumerator_get_container.
func (recv *FileEnumerator) GetContainer() *File {
	retC := C.g_file_enumerator_get_container((*C.GFileEnumerator)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDataSize is a wrapper around the C function g_memory_output_stream_get_data_size.
func (recv *MemoryOutputStream) GetDataSize() uint64 {
	retC := C.g_memory_output_stream_get_data_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// PrependName is a wrapper around the C function g_themed_icon_prepend_name.
func (recv *ThemedIcon) PrependName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_prepend_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
}

// SetRateLimit is a wrapper around the C function g_unix_mount_monitor_set_rate_limit.
func (recv *UnixMountMonitor) SetRateLimit(limitMsec int32) {
	c_limit_msec := (C.int)(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit((*C.GUnixMountMonitor)(recv.native), c_limit_msec)

	return
}

type signalVolumeMonitorDriveEjectButtonDetail struct {
	callback  VolumeMonitorSignalDriveEjectButtonCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveEjectButtonId int
var signalVolumeMonitorDriveEjectButtonMap = make(map[int]signalVolumeMonitorDriveEjectButtonDetail)
var signalVolumeMonitorDriveEjectButtonLock sync.RWMutex

// VolumeMonitorSignalDriveEjectButtonCallback is a callback function for a 'drive-eject-button' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveEjectButtonCallback func(drive *Drive)

/*
ConnectDriveEjectButton connects the callback to the 'drive-eject-button' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveEjectButton to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveEjectButton(callback VolumeMonitorSignalDriveEjectButtonCallback) int {
	signalVolumeMonitorDriveEjectButtonLock.Lock()
	defer signalVolumeMonitorDriveEjectButtonLock.Unlock()

	signalVolumeMonitorDriveEjectButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_eject_button(instance, C.gpointer(uintptr(signalVolumeMonitorDriveEjectButtonId)))

	detail := signalVolumeMonitorDriveEjectButtonDetail{callback, handlerID}
	signalVolumeMonitorDriveEjectButtonMap[signalVolumeMonitorDriveEjectButtonId] = detail

	return signalVolumeMonitorDriveEjectButtonId
}

/*
DisconnectDriveEjectButton disconnects a callback from the 'drive-eject-button' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveEjectButton.
*/
func (recv *VolumeMonitor) DisconnectDriveEjectButton(connectionID int) {
	signalVolumeMonitorDriveEjectButtonLock.Lock()
	defer signalVolumeMonitorDriveEjectButtonLock.Unlock()

	detail, exists := signalVolumeMonitorDriveEjectButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveEjectButtonMap, connectionID)
}

//export volumemonitor_driveEjectButtonHandler
func volumemonitor_driveEjectButtonHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveEjectButtonLock.RLock()
	defer signalVolumeMonitorDriveEjectButtonLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveEjectButtonMap[index].callback
	callback(drive)
}

type EmblemOrigin C.GEmblemOrigin

const (
	EMBLEM_ORIGIN_UNKNOWN      EmblemOrigin = 0
	EMBLEM_ORIGIN_DEVICE       EmblemOrigin = 1
	EMBLEM_ORIGIN_LIVEMETADATA EmblemOrigin = 2
	EMBLEM_ORIGIN_TAG          EmblemOrigin = 3
)

// ContentTypeFromMimeType is a wrapper around the C function g_content_type_from_mime_type.
func ContentTypeFromMimeType(mimeType string) string {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.g_content_type_from_mime_type(c_mime_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGuessForTree is a wrapper around the C function g_content_type_guess_for_tree.
func ContentTypeGuessForTree(root *File) []string {
	c_root := (*C.GFile)(root.ToC())

	retC := C.g_content_type_guess_for_tree(c_root)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// MakeDirectoryWithParents is a wrapper around the C function g_file_make_directory_with_parents.
func (recv *File) MakeDirectoryWithParents(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_directory_with_parents((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Monitor is a wrapper around the C function g_file_monitor.
func (recv *File) Monitor(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QueryFileType is a wrapper around the C function g_file_query_file_type.
func (recv *File) QueryFileType(flags FileQueryInfoFlags, cancellable *Cancellable) FileType {
	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_file_query_file_type((*C.GFile)(recv.native), c_flags, c_cancellable)
	retGo := (FileType)(retC)

	return retGo
}

// Unsupported : g_mount_guess_content_type : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GuessContentTypeFinish is a wrapper around the C function g_mount_guess_content_type_finish.
func (recv *Mount) GuessContentTypeFinish(result *AsyncResult) ([]string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_guess_content_type_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GuessContentTypeSync is a wrapper around the C function g_mount_guess_content_type_sync.
func (recv *Mount) GuessContentTypeSync(forceRescan bool, cancellable *Cancellable) ([]string, error) {
	c_force_rescan :=
		boolToGboolean(forceRescan)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_mount_guess_content_type_sync((*C.GMount)(recv.native), c_force_rescan, c_cancellable, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetActivationRoot is a wrapper around the C function g_volume_get_activation_root.
func (recv *Volume) GetActivationRoot() *File {
	retC := C.g_volume_get_activation_root((*C.GVolume)(recv.native))
	var retGo (*File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
