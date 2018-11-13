// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Marks @application as busy (see g_application_mark_busy()) while
// @property on @object is %TRUE.
//
// The binding holds a reference to @application while it is active, but
// not to @object. Instead, the binding is destroyed when @object is
// finalized.
/*

C function : g_application_bind_busy_property
*/
func (recv *Application) BindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_bind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// Gets the application's current busy state, as set through
// g_application_mark_busy() or g_application_bind_busy_property().
/*

C function : g_application_get_is_busy
*/
func (recv *Application) GetIsBusy() bool {
	retC := C.g_application_get_is_busy((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Destroys a binding between @property and the busy state of
// @application that was previously created with
// g_application_bind_busy_property().
/*

C function : g_application_unbind_busy_property
*/
func (recv *Application) UnbindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_unbind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// This is a version of g_file_enumerator_next_file() that's easier to
// use correctly from C programs.  With g_file_enumerator_next_file(),
// the gboolean return value signifies "end of iteration or error", which
// requires allocation of a temporary #GError.
//
// In contrast, with this function, a %FALSE return from
// g_file_enumerator_iterate() *always* means
// "error".  End of iteration is signaled by @out_info or @out_child being %NULL.
//
// Another crucial difference is that the references for @out_info and
// @out_child are owned by @direnum (they are cached as hidden
// properties).  You must not unref them in your own code.  This makes
// memory management significantly easier for C code in combination
// with loops.
//
// Finally, this function optionally allows retrieving a #GFile as
// well.
//
// You must specify at least one of @out_info or @out_child.
//
// The code pattern for correctly using g_file_enumerator_iterate() from C
// is:
//
// |[
// direnum = g_file_enumerate_children (file, ...);
// while (TRUE)
// {
// GFileInfo *info;
// if (!g_file_enumerator_iterate (direnum, &info, NULL, cancellable, error))
// goto out;
// if (!info)
// break;
// ... do stuff with "info"; do not unref it! ...
// }
//
// out:
// g_object_unref (direnum); // Note: frees the last @info
// ]|
/*

C function : g_file_enumerator_iterate
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outInfo := FileInfoNewFromC(unsafe.Pointer(c_out_info))

	outChild := FileNewFromC(unsafe.Pointer(c_out_child))

	return retGo, outInfo, outChild, goThrowableError
}

// Unsupported : g_input_stream_read_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous stream read operation started with
// g_input_stream_read_all_async().
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_read will be set to the number of bytes that were successfully
// read before the error was encountered.  This functionality is only
// available from C.  If you need it from another language then you must
// write your own loop around g_input_stream_read_async().
/*

C function : g_input_stream_read_all_finish
*/
func (recv *InputStream) ReadAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_read C.gsize

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_all_finish((*C.GInputStream)(recv.native), c_result, &c_bytes_read, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	return retGo, bytesRead, goThrowableError
}

// Creates a new #GListStore with items of type @item_type. @item_type
// must be a subclass of #GObject.
/*

C function : g_list_store_new
*/
func ListStoreNew(itemType gobject.Type) *ListStore {
	c_item_type := (C.GType)(itemType)

	retC := C.g_list_store_new(c_item_type)
	retGo := ListStoreNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends @item to @store. @item must be of type #GListStore:item-type.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to append multiple items at the same time
// efficiently.
/*

C function : g_list_store_append
*/
func (recv *ListStore) Append(item uintptr) {
	c_item := (C.gpointer)(item)

	C.g_list_store_append((*C.GListStore)(recv.native), c_item)

	return
}

// Inserts @item into @store at @position. @item must be of type
// #GListStore:item-type or derived from it. @position must be smaller
// than the length of the list, or equal to it to append.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to insert multiple items at the same time
// efficiently.
/*

C function : g_list_store_insert
*/
func (recv *ListStore) Insert(position uint32, item uintptr) {
	c_position := (C.guint)(position)

	c_item := (C.gpointer)(item)

	C.g_list_store_insert((*C.GListStore)(recv.native), c_position, c_item)

	return
}

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// Removes the item from @store that is at @position. @position must be
// smaller than the current length of the list.
//
// Use g_list_store_splice() to remove multiple items at the same time
// efficiently.
/*

C function : g_list_store_remove
*/
func (recv *ListStore) Remove(position uint32) {
	c_position := (C.guint)(position)

	C.g_list_store_remove((*C.GListStore)(recv.native), c_position)

	return
}

// Removes all items from @store.
/*

C function : g_list_store_remove_all
*/
func (recv *ListStore) RemoveAll() {
	C.g_list_store_remove_all((*C.GListStore)(recv.native))

	return
}

// Unsupported : g_list_store_splice : unsupported parameter additions :

// Creates a new #GSocketConnectable for connecting to the local host
// over a loopback connection to the given @port. This is intended for
// use in connecting to local services which may be running on IPv4 or
// IPv6.
//
// The connectable will return IPv4 and IPv6 loopback addresses,
// regardless of how the host resolves `localhost`. By contrast,
// g_network_address_new() will often only return an IPv4 address when
// resolving `localhost`, and an IPv6 address for `localhost6`.
//
// g_network_address_get_hostname() will always return `localhost` for
// #GNetworkAddresses created with this constructor.
/*

C function : g_network_address_new_loopback
*/
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	c_port := (C.guint16)(port)

	retC := C.g_network_address_new_loopback(c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_output_stream_write_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous stream write operation started with
// g_output_stream_write_all_async().
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C.  If you need it from another
// language then you must write your own loop around
// g_output_stream_write_async().
/*

C function : g_output_stream_write_all_finish
*/
func (recv *OutputStream) WriteAllFinish(result *AsyncResult) (bool, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_all_finish((*C.GOutputStream)(recv.native), c_result, &c_bytes_written, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goThrowableError
}

// Unsupported : g_simple_action_set_state_hint : unsupported parameter state_hint : Blacklisted record : GVariant

// GSimpleIOStream creates a #GIOStream from an arbitrary #GInputStream and
// #GOutputStream. This allows any pair of input and output streams to be used
// with #GIOStream methods.
//
// This is useful when you obtained a #GInputStream and a #GOutputStream
// by other means, for instance creating them with platform specific methods as
// g_unix_input_stream_new() or g_win32_input_stream_new(), and you want
// to take advantage of the methods provided by #GIOStream.
/*

C record/class : GSimpleIOStream
*/
type SimpleIOStream struct {
	native *C.GSimpleIOStream
}

func SimpleIOStreamNewFromC(u unsafe.Pointer) *SimpleIOStream {
	c := (*C.GSimpleIOStream)(u)
	if c == nil {
		return nil
	}

	g := &SimpleIOStream{native: c}

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitary Object to SimpleIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleIOStream.
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromC(object.ToC())
}

// Creates a new #GSimpleIOStream wrapping @input_stream and @output_stream.
// See also #GIOStream.
/*

C function : g_simple_io_stream_new
*/
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

	return retGo
}

// Unsupported : g_socket_send_messages : unsupported parameter messages :

// Gets the value of #GTask:completed. This changes from %FALSE to %TRUE after
// the task’s callback is invoked, and will return %FALSE if called from inside
// the callback.
/*

C function : g_task_get_completed
*/
func (recv *Task) GetCompleted() bool {
	retC := C.g_task_get_completed((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
