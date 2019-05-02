// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// gdk_pixbuf_get_file_info_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
// PixbufGetFileInfoFinish is a wrapper around the C function gdk_pixbuf_get_file_info_finish.
func PixbufGetFileInfoFinish(asyncResult *gio.AsyncResult) (*PixbufFormat, int32, int32, error) {
	c_async_result := (*C.GAsyncResult)(asyncResult.ToC())

	var c_width C.gint

	var c_height C.gint

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_get_file_info_finish(c_async_result, &c_width, &c_height, &cThrowableError)
	retGo := PixbufFormatNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height, goError
}

// GetOptions is a wrapper around the C function gdk_pixbuf_get_options.
func (recv *Pixbuf) GetOptions() *glib.HashTable {
	retC := C.gdk_pixbuf_get_options((*C.GdkPixbuf)(recv.native))
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReadPixelBytes is a wrapper around the C function gdk_pixbuf_read_pixel_bytes.
func (recv *Pixbuf) ReadPixelBytes() *glib.Bytes {
	retC := C.gdk_pixbuf_read_pixel_bytes((*C.GdkPixbuf)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_pixbuf_read_pixels
