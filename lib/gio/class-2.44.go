// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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
// #include <stdlib.h>
import "C"

// BindBusyProperty is a wrapper around the C function g_application_bind_busy_property.
func (recv *Application) BindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_bind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// GetIsBusy is a wrapper around the C function g_application_get_is_busy.
func (recv *Application) GetIsBusy() bool {
	retC := C.g_application_get_is_busy((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnbindBusyProperty is a wrapper around the C function g_application_unbind_busy_property.
func (recv *Application) UnbindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_unbind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// Iterate is a wrapper around the C function g_file_enumerator_iterate.
func (recv *FileEnumerator) Iterate(cancellable *Cancellable) (bool, *FileInfo, *File, error) {
	var c_out_info *C.GFileInfo

	var c_out_child *C.GFile

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_iterate((*C.GFileEnumerator)(recv.native), &c_out_info, &c_out_child, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outInfo := FileInfoNewFromC(unsafe.Pointer(c_out_info))

	outChild := FileNewFromC(unsafe.Pointer(c_out_child))

	return retGo, outInfo, outChild, goError
}

// Unsupported : g_input_stream_read_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadAllFinish is a wrapper around the C function g_input_stream_read_all_finish.
func (recv *InputStream) ReadAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_read C.gsize

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_all_finish((*C.GInputStream)(recv.native), c_result, &c_bytes_read, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	return retGo, bytesRead, goError
}

// ListStoreNew is a wrapper around the C function g_list_store_new.
func ListStoreNew(itemType gobject.Type) *ListStore {
	c_item_type := (C.GType)(itemType)

	retC := C.g_list_store_new(c_item_type)
	retGo := ListStoreNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Append is a wrapper around the C function g_list_store_append.
func (recv *ListStore) Append(item uintptr) {
	c_item := (C.gpointer)(item)

	C.g_list_store_append((*C.GListStore)(recv.native), c_item)

	return
}

// Insert is a wrapper around the C function g_list_store_insert.
func (recv *ListStore) Insert(position uint32, item uintptr) {
	c_position := (C.guint)(position)

	c_item := (C.gpointer)(item)

	C.g_list_store_insert((*C.GListStore)(recv.native), c_position, c_item)

	return
}

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// Remove is a wrapper around the C function g_list_store_remove.
func (recv *ListStore) Remove(position uint32) {
	c_position := (C.guint)(position)

	C.g_list_store_remove((*C.GListStore)(recv.native), c_position)

	return
}

// RemoveAll is a wrapper around the C function g_list_store_remove_all.
func (recv *ListStore) RemoveAll() {
	C.g_list_store_remove_all((*C.GListStore)(recv.native))

	return
}

// Unsupported : g_list_store_splice : unsupported parameter additions :

// NetworkAddressNewLoopback is a wrapper around the C function g_network_address_new_loopback.
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	c_port := (C.guint16)(port)

	retC := C.g_network_address_new_loopback(c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_output_stream_write_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WriteAllFinish is a wrapper around the C function g_output_stream_write_all_finish.
func (recv *OutputStream) WriteAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_all_finish((*C.GOutputStream)(recv.native), c_result, &c_bytes_written, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// SetStateHint is a wrapper around the C function g_simple_action_set_state_hint.
func (recv *SimpleAction) SetStateHint(stateHint *glib.Variant) {
	c_state_hint := (*C.GVariant)(C.NULL)
	if stateHint != nil {
		c_state_hint = (*C.GVariant)(stateHint.ToC())
	}

	C.g_simple_action_set_state_hint((*C.GSimpleAction)(recv.native), c_state_hint)

	return
}

// SimpleIOStream is a wrapper around the C record GSimpleIOStream.
type SimpleIOStream struct {
	native *C.GSimpleIOStream
}

func SimpleIOStreamNewFromC(u unsafe.Pointer) *SimpleIOStream {
	c := (*C.GSimpleIOStream)(u)
	if c == nil {
		return nil
	}

	g := &SimpleIOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleIOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleIOStream with another SimpleIOStream, and returns true if they represent the same GObject.
func (recv *SimpleIOStream) Equals(other *SimpleIOStream) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to SimpleIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleIOStream.
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromC(object.ToC())
}

// SimpleIOStreamNew is a wrapper around the C function g_simple_io_stream_new.
func SimpleIOStreamNew(inputStream *InputStream, outputStream *OutputStream) *SimpleIOStream {
	c_input_stream := (*C.GInputStream)(C.NULL)
	if inputStream != nil {
		c_input_stream = (*C.GInputStream)(inputStream.ToC())
	}

	c_output_stream := (*C.GOutputStream)(C.NULL)
	if outputStream != nil {
		c_output_stream = (*C.GOutputStream)(outputStream.ToC())
	}

	retC := C.g_simple_io_stream_new(c_input_stream, c_output_stream)
	retGo := SimpleIOStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_socket_send_messages : unsupported parameter messages :

// GetCompleted is a wrapper around the C function g_task_get_completed.
func (recv *Task) GetCompleted() bool {
	retC := C.g_task_get_completed((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountMonitorGet is a wrapper around the C function g_unix_mount_monitor_get.
func UnixMountMonitorGet() *UnixMountMonitor {
	retC := C.g_unix_mount_monitor_get()
	retGo := UnixMountMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}
