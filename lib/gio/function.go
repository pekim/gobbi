// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
import "C"

// ContentTypeCanBeExecutable is a wrapper around the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_can_be_executable(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeEquals is a wrapper around the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	c_type1 := C.CString(type1)
	defer C.free(unsafe.Pointer(c_type1))

	c_type2 := C.CString(type2)
	defer C.free(unsafe.Pointer(c_type2))

	retC := C.g_content_type_equals(c_type1, c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeGetDescription is a wrapper around the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_description(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGetIcon is a wrapper around the C function g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) *Icon {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_icon(c_type)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGetMimeType is a wrapper around the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_mime_type(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGuess is a wrapper around the C function g_content_type_guess.
func ContentTypeGuess(filename string, data []uint8) (string, bool) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_data_array := make([]C.guchar, len(data), len(data))
	c_data := &c_data_array[0]

	c_data_size := (C.gsize)(len(data))

	var c_result_uncertain C.gboolean

	retC := C.g_content_type_guess(c_filename, c_data, c_data_size, &c_result_uncertain)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	resultUncertain := c_result_uncertain == C.TRUE

	return retGo, resultUncertain
}

// ContentTypeIsA is a wrapper around the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_supertype := C.CString(supertype)
	defer C.free(unsafe.Pointer(c_supertype))

	retC := C.g_content_type_is_a(c_type, c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeIsUnknown is a wrapper around the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_is_unknown(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypesGetRegistered is a wrapper around the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retC := C.g_content_types_get_registered()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon

// IoErrorFromErrno is a wrapper around the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	c_err_no := (C.gint)(errNo)

	retC := C.g_io_error_from_errno(c_err_no)
	retGo := (IOErrorEnum)(retC)

	return retGo
}

// IoErrorQuark is a wrapper around the C function g_io_error_quark.
func IoErrorQuark() glib.Quark {
	retC := C.g_io_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// IoModulesLoadAllInDirectory is a wrapper around the C function g_io_modules_load_all_in_directory.
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	retC := C.g_io_modules_load_all_in_directory(c_dirname)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoSchedulerCancelAllJobs is a wrapper around the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()

	return
}

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Blacklisted : g_keyfile_settings_backend_new

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnixIsMountPathSystemInternal is a wrapper around the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	retC := C.g_unix_is_mount_path_system_internal(c_mount_path)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountAt is a wrapper around the C function g_unix_mount_at.
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_at(c_mount_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountCompare is a wrapper around the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	c_mount1 := (*C.GUnixMountEntry)(C.NULL)
	if mount1 != nil {
		c_mount1 = (*C.GUnixMountEntry)(mount1.ToC())
	}

	c_mount2 := (*C.GUnixMountEntry)(C.NULL)
	if mount2 != nil {
		c_mount2 = (*C.GUnixMountEntry)(mount2.ToC())
	}

	retC := C.g_unix_mount_compare(c_mount1, c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// UnixMountFree is a wrapper around the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	C.g_unix_mount_free(c_mount_entry)

	return
}

// UnixMountGetDevicePath is a wrapper around the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_device_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetFsType is a wrapper around the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_fs_type(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetMountPath is a wrapper around the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_mount_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGuessCanEject is a wrapper around the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_can_eject(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountGuessIcon is a wrapper around the C function g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_icon(c_mount_entry)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnixMountGuessName is a wrapper around the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_name(c_mount_entry)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnixMountGuessShouldDisplay is a wrapper around the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_should_display(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountIsReadonly is a wrapper around the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_readonly(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountIsSystemInternal is a wrapper around the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_system_internal(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsChangedSince is a wrapper around the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mount_points_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsGet is a wrapper around the C function g_unix_mount_points_get.
func UnixMountPointsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mount_points_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountsChangedSince is a wrapper around the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mounts_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountsGet is a wrapper around the C function g_unix_mounts_get.
func UnixMountsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mounts_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}
