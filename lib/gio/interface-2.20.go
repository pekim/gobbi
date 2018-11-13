// This is a generated file - DO NOT EDIT
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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
import "C"

// Obtains the information whether the #GAppInfo can be deleted.
// See g_app_info_delete().
/*

C function : g_app_info_can_delete
*/
func (recv *AppInfo) CanDelete() bool {
	retC := C.g_app_info_can_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tries to delete a #GAppInfo.
//
// On some platforms, there may be a difference between user-defined
// #GAppInfos which can be deleted, and system-wide ones which cannot.
// See g_app_info_can_delete().
/*

C function : g_app_info_delete
*/
func (recv *AppInfo) Delete() bool {
	retC := C.g_app_info_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the commandline with which the application will be
// started.
/*

C function : g_app_info_get_commandline
*/
func (recv *AppInfo) GetCommandline() string {
	retC := C.g_app_info_get_commandline((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Generates a textual representation of @icon that can be used for
// serialization such as when passing @icon to a different process or
// saving it to persistent storage. Use g_icon_new_for_string() to
// get @icon back from the returned string.
//
// The encoding of the returned string is proprietary to #GIcon except
// in the following two cases
//
// - If @icon is a #GFileIcon, the returned string is a native path
// (such as `/path/to/my icon.png`) without escaping
// if the #GFile for @icon is a native file.  If the file is not
// native, the returned string is the result of g_file_get_uri()
// (such as `sftp://path/to/my%20icon.png`).
//
// - If @icon is a #GThemedIcon with exactly one name, the encoding is
// simply the name (such as `network-server`).
/*

C function : g_icon_to_string
*/
func (recv *Icon) ToString() string {
	retC := C.g_icon_to_string((*C.GIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Determines if @mount is shadowed. Applications or libraries should
// avoid displaying @mount in the user interface if it is shadowed.
//
// A mount is said to be shadowed if there exists one or more user
// visible objects (currently #GMount objects) with a root that is
// inside the root of @mount.
//
// One application of shadow mounts is when exposing a single file
// system that is used to address several logical volumes. In this
// situation, a #GVolumeMonitor implementation would create two
// #GVolume objects (for example, one for the camera functionality of
// the device and one for a SD card reader on the device) with
// activation URIs `gphoto2://[usb:001,002]/store1/`
// and `gphoto2://[usb:001,002]/store2/`. When the
// underlying mount (with root
// `gphoto2://[usb:001,002]/`) is mounted, said
// #GVolumeMonitor implementation would create two #GMount objects
// (each with their root matching the corresponding volume activation
// root) that would shadow the original mount.
//
// The proxy monitor in GVfs 2.26 and later, automatically creates and
// manage shadow mounts (and shadows the underlying mount) if the
// activation root on a #GVolume is set.
/*

C function : g_mount_is_shadowed
*/
func (recv *Mount) IsShadowed() bool {
	retC := C.g_mount_is_shadowed((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Increments the shadow count on @mount. Usually used by
// #GVolumeMonitor implementations when creating a shadow mount for
// @mount, see g_mount_is_shadowed() for more information. The caller
// will need to emit the #GMount::changed signal on @mount manually.
/*

C function : g_mount_shadow
*/
func (recv *Mount) Shadow() {
	C.g_mount_shadow((*C.GMount)(recv.native))

	return
}

// Decrements the shadow count on @mount. Usually used by
// #GVolumeMonitor implementations when destroying a shadow mount for
// @mount, see g_mount_is_shadowed() for more information. The caller
// will need to emit the #GMount::changed signal on @mount manually.
/*

C function : g_mount_unshadow
*/
func (recv *Mount) Unshadow() {
	C.g_mount_unshadow((*C.GMount)(recv.native))

	return
}
