// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void volumemonitor_driveEjectButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-eject-button", G_CALLBACK(volumemonitor_driveEjectButtonHandler), data);
	}

*/
import "C"

// Blacklisted : g_desktop_app_info_new_from_keyfile

// Blacklisted : g_emblem_new

// Blacklisted : g_emblem_new_with_origin

// Blacklisted : g_emblem_get_icon

// Blacklisted : g_emblem_get_origin

// Blacklisted : g_emblemed_icon_new

// Blacklisted : g_emblemed_icon_add_emblem

// Blacklisted : g_emblemed_icon_get_emblems

// Blacklisted : g_emblemed_icon_get_icon

// Blacklisted : g_file_enumerator_get_container

// Blacklisted : g_memory_output_stream_get_data_size

// Blacklisted : g_themed_icon_prepend_name

// Blacklisted : g_unix_mount_monitor_set_rate_limit

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
