// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// CopyOptions is a wrapper around the C function gdk_pixbuf_copy_options.
func (recv *Pixbuf) CopyOptions(destPixbuf *Pixbuf) bool {
	c_dest_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if destPixbuf != nil {
		c_dest_pixbuf = (*C.GdkPixbuf)(destPixbuf.ToC())
	}

	retC := C.gdk_pixbuf_copy_options((*C.GdkPixbuf)(recv.native), c_dest_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveOption is a wrapper around the C function gdk_pixbuf_remove_option.
func (recv *Pixbuf) RemoveOption(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_remove_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// SaveToStreamv is a wrapper around the C function gdk_pixbuf_save_to_streamv.
func (recv *Pixbuf) SaveToStreamv(stream *gio.OutputStream, type_ string, optionKeys []string, optionValues []string, cancellable *gio.Cancellable) (bool, error) {
	c_stream := (*C.GOutputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GOutputStream)(stream.ToC())
	}

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_option_keys_array := make([]*C.char, len(optionKeys), len(optionKeys))
	for i, item := range optionKeys {
		g := optionKeys[i]
		c := C.CString(g)
		c_option_keys_array[i] = c
	}
	c_option_keys_arrayPtr := &c_option_keys_array[0]
	c_option_keys := (**C.char)(unsafe.Pointer(c_option_keys_arrayPtr))

	c_option_values_array := make([]*C.char, len(optionValues), len(optionValues))
	for i, item := range optionValues {
		g := optionValues[i]
		c := C.CString(g)
		c_option_values_array[i] = c
	}
	c_option_values_arrayPtr := &c_option_values_array[0]
	c_option_values := (**C.char)(unsafe.Pointer(c_option_values_arrayPtr))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_save_to_streamv((*C.GdkPixbuf)(recv.native), c_stream, c_type, c_option_keys, c_option_values, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
