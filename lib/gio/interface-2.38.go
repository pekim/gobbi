// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// ActionNameIsValid is a wrapper around the C function g_action_name_is_valid.
func ActionNameIsValid(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_name_is_valid(c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// ActionParseDetailedName is a wrapper around the C function g_action_parse_detailed_name.
func ActionParseDetailedName(detailedName string) (bool, string, *glib.Variant, error) {
	c_detailed_name := C.CString(detailedName)
	defer C.free(unsafe.Pointer(c_detailed_name))

	var c_action_name *C.gchar

	var c_target_value *C.GVariant

	var cThrowableError *C.GError

	retC := C.g_action_parse_detailed_name(c_detailed_name, &c_action_name, &c_target_value, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	actionName := C.GoString(c_action_name)
	defer C.free(unsafe.Pointer(c_action_name))

	targetValue := glib.VariantNewFromC(unsafe.Pointer(c_target_value))

	return retGo, actionName, targetValue, goError
}

// ActionPrintDetailedName is a wrapper around the C function g_action_print_detailed_name.
func ActionPrintDetailedName(actionName string, targetValue *glib.Variant) string {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_target_value := (*C.GVariant)(C.NULL)
	if targetValue != nil {
		c_target_value = (*C.GVariant)(targetValue.ToC())
	}

	retC := C.g_action_print_detailed_name(c_action_name, c_target_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_make_directory_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MakeDirectoryFinish is a wrapper around the C function g_file_make_directory_finish.
func (recv *File) MakeDirectoryFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_make_directory_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_measure_disk_usage : unsupported parameter progress_callback : no type generator for FileMeasureProgressCallback (GFileMeasureProgressCallback) for param progress_callback

// Unsupported : g_file_measure_disk_usage_async : unsupported parameter progress_callback : no type generator for FileMeasureProgressCallback (GFileMeasureProgressCallback) for param progress_callback

// MeasureDiskUsageFinish is a wrapper around the C function g_file_measure_disk_usage_finish.
func (recv *File) MeasureDiskUsageFinish(result *AsyncResult) (bool, uint64, uint64, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_disk_usage C.guint64

	var c_num_dirs C.guint64

	var c_num_files C.guint64

	var cThrowableError *C.GError

	retC := C.g_file_measure_disk_usage_finish((*C.GFile)(recv.native), c_result, &c_disk_usage, &c_num_dirs, &c_num_files, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	diskUsage := (uint64)(c_disk_usage)

	numDirs := (uint64)(c_num_dirs)

	numFiles := (uint64)(c_num_files)

	return retGo, diskUsage, numDirs, numFiles, goError
}

// Unsupported : g_file_trash_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// TrashFinish is a wrapper around the C function g_file_trash_finish.
func (recv *File) TrashFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_trash_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IconDeserialize is a wrapper around the C function g_icon_deserialize.
func IconDeserialize(value *glib.Variant) *Icon {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_icon_deserialize(c_value)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Serialize is a wrapper around the C function g_icon_serialize.
func (recv *Icon) Serialize() *glib.Variant {
	retC := C.g_icon_serialize((*C.GIcon)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}
