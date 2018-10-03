// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
import "C"

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// DesktopAppInfoNewFromKeyfile is a wrapper around the C function g_desktop_app_info_new_from_keyfile.
func DesktopAppInfoNewFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	retC := C.g_desktop_app_info_new_from_keyfile(c_key_file)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_get_icon : no return generator

// GetOrigin is a wrapper around the C function g_emblem_get_origin.
func (recv *Emblem) GetOrigin() EmblemOrigin {
	retC := C.g_emblem_get_origin((*C.GEmblem)(recv.native))
	retGo := (EmblemOrigin)(retC)

	return retGo
}

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// AddEmblem is a wrapper around the C function g_emblemed_icon_add_emblem.
func (recv *EmblemedIcon) AddEmblem(emblem *Emblem) {
	c_emblem := (*C.GEmblem)(emblem.ToC())

	C.g_emblemed_icon_add_emblem((*C.GEmblemedIcon)(recv.native), c_emblem)

	return
}

// GetEmblems is a wrapper around the C function g_emblemed_icon_get_emblems.
func (recv *EmblemedIcon) GetEmblems() *glib.List {
	retC := C.g_emblemed_icon_get_emblems((*C.GEmblemedIcon)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_emblemed_icon_get_icon : no return generator

// Unsupported : g_file_enumerator_get_container : no return generator

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// GetDataSize is a wrapper around the C function g_memory_output_stream_get_data_size.
func (recv *MemoryOutputStream) GetDataSize() uint64 {
	retC := C.g_memory_output_stream_get_data_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// PrependName is a wrapper around the C function g_themed_icon_prepend_name.
func (recv *ThemedIcon) PrependName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_prepend_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// SetRateLimit is a wrapper around the C function g_unix_mount_monitor_set_rate_limit.
func (recv *UnixMountMonitor) SetRateLimit(limitMsec int32) {
	c_limit_msec := (C.int)(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit((*C.GUnixMountMonitor)(recv.native), c_limit_msec)

	return
}

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type