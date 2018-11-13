// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

	void volumemonitor_driveEjectButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-eject-button", G_CALLBACK(volumemonitor_driveEjectButtonHandler), data);
	}

*/
import "C"

// Creates a new #GDesktopAppInfo.
/*

C function : g_desktop_app_info_new_from_keyfile
*/
func DesktopAppInfoNewFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	retC := C.g_desktop_app_info_new_from_keyfile(c_key_file)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new emblem for @icon.
/*

C function : g_emblem_new
*/
func EmblemNew(icon *Icon) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.g_emblem_new(c_icon)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new emblem for @icon.
/*

C function : g_emblem_new_with_origin
*/
func EmblemNewWithOrigin(icon *Icon, origin EmblemOrigin) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	c_origin := (C.GEmblemOrigin)(origin)

	retC := C.g_emblem_new_with_origin(c_icon, c_origin)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gives back the icon from @emblem.
/*

C function : g_emblem_get_icon
*/
func (recv *Emblem) GetIcon() *Icon {
	retC := C.g_emblem_get_icon((*C.GEmblem)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the origin of the emblem.
/*

C function : g_emblem_get_origin
*/
func (recv *Emblem) GetOrigin() EmblemOrigin {
	retC := C.g_emblem_get_origin((*C.GEmblem)(recv.native))
	retGo := (EmblemOrigin)(retC)

	return retGo
}

// Creates a new emblemed icon for @icon with the emblem @emblem.
/*

C function : g_emblemed_icon_new
*/
func EmblemedIconNew(icon *Icon, emblem *Emblem) *EmblemedIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	c_emblem := (*C.GEmblem)(C.NULL)
	if emblem != nil {
		c_emblem = (*C.GEmblem)(emblem.ToC())
	}

	retC := C.g_emblemed_icon_new(c_icon, c_emblem)
	retGo := EmblemedIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds @emblem to the #GList of #GEmblems.
/*

C function : g_emblemed_icon_add_emblem
*/
func (recv *EmblemedIcon) AddEmblem(emblem *Emblem) {
	c_emblem := (*C.GEmblem)(C.NULL)
	if emblem != nil {
		c_emblem = (*C.GEmblem)(emblem.ToC())
	}

	C.g_emblemed_icon_add_emblem((*C.GEmblemedIcon)(recv.native), c_emblem)

	return
}

// Gets the list of emblems for the @icon.
/*

C function : g_emblemed_icon_get_emblems
*/
func (recv *EmblemedIcon) GetEmblems() *glib.List {
	retC := C.g_emblemed_icon_get_emblems((*C.GEmblemedIcon)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the main icon for @emblemed.
/*

C function : g_emblemed_icon_get_icon
*/
func (recv *EmblemedIcon) GetIcon() *Icon {
	retC := C.g_emblemed_icon_get_icon((*C.GEmblemedIcon)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the #GFile container which is being enumerated.
/*

C function : g_file_enumerator_get_container
*/
func (recv *FileEnumerator) GetContainer() *File {
	retC := C.g_file_enumerator_get_container((*C.GFileEnumerator)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the number of bytes from the start up to including the last
// byte written in the stream that has not been truncated away.
/*

C function : g_memory_output_stream_get_data_size
*/
func (recv *MemoryOutputStream) GetDataSize() uint64 {
	retC := C.g_memory_output_stream_get_data_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Prepend a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls
// to g_icon_hash().
/*

C function : g_themed_icon_prepend_name
*/
func (recv *ThemedIcon) PrependName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_prepend_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
}

// This function does nothing.
//
// Before 2.44, this was a partially-effective way of controlling the
// rate at which events would be reported under some uncommon
// circumstances.  Since @mount_monitor is a singleton, it also meant
// that calling this function would have side effects for other users of
// the monitor.
/*

C function : g_unix_mount_monitor_set_rate_limit
*/
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
var signalVolumeMonitorDriveEjectButtonLock sync.Mutex

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
	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveEjectButtonMap[index].callback
	callback(drive)
}
