// Code generated - DO NOT EDIT.
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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

type FileMeasureFlags C.GFileMeasureFlags

const (
	FILE_MEASURE_NONE             FileMeasureFlags = 0
	FILE_MEASURE_REPORT_ANY_ERROR FileMeasureFlags = 2
	FILE_MEASURE_APPARENT_SIZE    FileMeasureFlags = 4
	FILE_MEASURE_NO_XDEV          FileMeasureFlags = 8
)

// MarkBusy is a wrapper around the C function g_application_mark_busy.
func (recv *Application) MarkBusy() {
	C.g_application_mark_busy((*C.GApplication)(recv.native))

	return
}

// UnmarkBusy is a wrapper around the C function g_application_unmark_busy.
func (recv *Application) UnmarkBusy() {
	C.g_application_unmark_busy((*C.GApplication)(recv.native))

	return
}

// BytesIconNew is a wrapper around the C function g_bytes_icon_new.
func BytesIconNew(bytes *glib.Bytes) *BytesIcon {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	retC := C.g_bytes_icon_new(c_bytes)
	retGo := BytesIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBytes is a wrapper around the C function g_bytes_icon_get_bytes.
func (recv *BytesIcon) GetBytes() *glib.Bytes {
	retC := C.g_bytes_icon_get_bytes((*C.GBytesIcon)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPropertyInfo is a wrapper around the C function g_dbus_method_invocation_get_property_info.
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	retC := C.g_dbus_method_invocation_get_property_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActionName is a wrapper around the C function g_desktop_app_info_get_action_name.
func (recv *DesktopAppInfo) GetActionName(actionName string) string {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_desktop_app_info_get_action_name((*C.GDesktopAppInfo)(recv.native), c_action_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// LaunchAction is a wrapper around the C function g_desktop_app_info_launch_action.
func (recv *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_launch_context := (*C.GAppLaunchContext)(C.NULL)
	if launchContext != nil {
		c_launch_context = (*C.GAppLaunchContext)(launchContext.ToC())
	}

	C.g_desktop_app_info_launch_action((*C.GDesktopAppInfo)(recv.native), c_action_name, c_launch_context)

	return
}

// ListActions is a wrapper around the C function g_desktop_app_info_list_actions.
func (recv *DesktopAppInfo) ListActions() []string {
	retC := C.g_desktop_app_info_list_actions((*C.GDesktopAppInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// RemoveAll is a wrapper around the C function g_menu_remove_all.
func (recv *Menu) RemoveAll() {
	C.g_menu_remove_all((*C.GMenu)(recv.native))

	return
}

// SetIcon is a wrapper around the C function g_menu_item_set_icon.
func (recv *MenuItem) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_menu_item_set_icon((*C.GMenuItem)(recv.native), c_icon)

	return
}

// PropertyAction is a wrapper around the C record GPropertyAction.
type PropertyAction struct {
	native *C.GPropertyAction
}

func PropertyActionNewFromC(u unsafe.Pointer) *PropertyAction {
	c := (*C.GPropertyAction)(u)
	if c == nil {
		return nil
	}

	g := &PropertyAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PropertyAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PropertyAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PropertyAction with another PropertyAction, and returns true if they represent the same GObject.
func (recv *PropertyAction) Equals(other *PropertyAction) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PropertyAction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PropertyAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a PropertyAction.
func CastToPropertyAction(object *gobject.Object) *PropertyAction {
	return PropertyActionNewFromC(object.ToC())
}

// PropertyActionNew is a wrapper around the C function g_property_action_new.
func PropertyActionNew(name string, object *gobject.Object, propertyName string) *PropertyAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_property_action_new(c_name, c_object, c_property_name)
	retGo := PropertyActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

const MENU_ATTRIBUTE_ICON string = C.G_MENU_ATTRIBUTE_ICON

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
