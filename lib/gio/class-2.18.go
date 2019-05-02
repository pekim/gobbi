// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

/*

	void volumemonitor_driveEjectButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-eject-button", G_CALLBACK(volumemonitor_driveEjectButtonHandler), data);
	}

*/
import "C"

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
