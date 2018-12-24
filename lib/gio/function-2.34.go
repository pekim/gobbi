// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// ContentTypeGetGenericIconName is a wrapper around the C function g_content_type_get_generic_icon_name.
func ContentTypeGetGenericIconName(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_generic_icon_name(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGetSymbolicIcon is a wrapper around the C function g_content_type_get_symbolic_icon.
func ContentTypeGetSymbolicIcon(type_ string) *Icon {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_symbolic_icon(c_type)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PollableSourceNewFull is a wrapper around the C function g_pollable_source_new_full.
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

// PollableStreamRead is a wrapper around the C function g_pollable_stream_read.
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PollableStreamWrite is a wrapper around the C function g_pollable_stream_write.
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PollableStreamWriteAll is a wrapper around the C function g_pollable_stream_write_all.
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// UnixMountGuessSymbolicIcon is a wrapper around the C function g_unix_mount_guess_symbolic_icon.
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) *Icon {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_symbolic_icon(c_mount_entry)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
