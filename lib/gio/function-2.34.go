// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Gets the generic icon name for a content type.
//
// See the
// [shared-mime-info](http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on the generic icon name.
/*

C function : g_content_type_get_generic_icon_name
*/
func ContentTypeGetGenericIconName(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_generic_icon_name(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the symbolic icon for a content type.
/*

C function : g_content_type_get_symbolic_icon
*/
func ContentTypeGetSymbolicIcon(type_ string) *Icon {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_symbolic_icon(c_type)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Utility method for #GPollableInputStream and #GPollableOutputStream
// implementations. Creates a new #GSource, as with
// g_pollable_source_new(), but also attaching @child_source (with a
// dummy callback), and @cancellable, if they are non-%NULL.
/*

C function : g_pollable_source_new_full
*/
func PollableSourceNewFull(pollableStream uintptr, childSource *glib.Source, cancellable *Cancellable) *glib.Source {
	c_pollable_stream := (C.gpointer)(pollableStream)

	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_source_new_full(c_pollable_stream, c_child_source, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Tries to read from @stream, as with g_input_stream_read() (if
// @blocking is %TRUE) or g_pollable_input_stream_read_nonblocking()
// (if @blocking is %FALSE). This can be used to more easily share
// code between blocking and non-blocking implementations of a method.
//
// If @blocking is %FALSE, then @stream must be a
// #GPollableInputStream for which g_pollable_input_stream_can_poll()
// returns %TRUE, or else the behavior is undefined. If @blocking is
// %TRUE, then @stream does not need to be a #GPollableInputStream.
/*

C function : g_pollable_stream_read
*/
func PollableStreamRead(stream *InputStream, buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_stream_read(c_stream, (unsafe.Pointer(c_buffer)), c_count, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to write to @stream, as with g_output_stream_write() (if
// @blocking is %TRUE) or g_pollable_output_stream_write_nonblocking()
// (if @blocking is %FALSE). This can be used to more easily share
// code between blocking and non-blocking implementations of a method.
//
// If @blocking is %FALSE, then @stream must be a
// #GPollableOutputStream for which
// g_pollable_output_stream_can_poll() returns %TRUE or else the
// behavior is undefined. If @blocking is %TRUE, then @stream does not
// need to be a #GPollableOutputStream.
/*

C function : g_pollable_stream_write
*/
func PollableStreamWrite(stream *OutputStream, buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_stream := (*C.GOutputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GOutputStream)(stream.ToC())
	}

	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_stream_write(c_stream, (unsafe.Pointer(c_buffer)), c_count, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to write @count bytes to @stream, as with
// g_output_stream_write_all(), but using g_pollable_stream_write()
// rather than g_output_stream_write().
//
// On a successful write of @count bytes, %TRUE is returned, and
// @bytes_written is set to @count.
//
// If there is an error during the operation (including
// %G_IO_ERROR_WOULD_BLOCK in the non-blocking case), %FALSE is
// returned and @error is set to indicate the error status,
// @bytes_written is updated to contain the number of bytes written
// into the stream before the error occurred.
//
// As with g_pollable_stream_write(), if @blocking is %FALSE, then
// @stream must be a #GPollableOutputStream for which
// g_pollable_output_stream_can_poll() returns %TRUE or else the
// behavior is undefined. If @blocking is %TRUE, then @stream does not
// need to be a #GPollableOutputStream.
/*

C function : g_pollable_stream_write_all
*/
func PollableStreamWriteAll(stream *OutputStream, buffer []uint8, blocking bool, cancellable *Cancellable) (bool, uint64, error) {
	c_stream := (*C.GOutputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GOutputStream)(stream.ToC())
	}

	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	var c_bytes_written C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_stream_write_all(c_stream, (unsafe.Pointer(c_buffer)), c_count, c_blocking, &c_bytes_written, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goThrowableError
}

// Guesses the symbolic icon of a Unix mount.
/*

C function : g_unix_mount_guess_symbolic_icon
*/
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) *Icon {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_symbolic_icon(c_mount_entry)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
