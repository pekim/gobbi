// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Returns information about the license of the image loader for the format. The
// returned string should be a shorthand for a wellknown license, e.g. "LGPL",
// "GPL", "QPL", "GPL/QPL", or "other" to indicate some other license.  This
// string should be freed with g_free() when it's no longer needed.
/*

C function : gdk_pixbuf_format_get_license
*/
func (recv *PixbufFormat) GetLicense() string {
	retC := C.gdk_pixbuf_format_get_license((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns whether this image format is disabled. See
// gdk_pixbuf_format_set_disabled().
/*

C function : gdk_pixbuf_format_is_disabled
*/
func (recv *PixbufFormat) IsDisabled() bool {
	retC := C.gdk_pixbuf_format_is_disabled((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether this image format is scalable. If a file is in a
// scalable format, it is preferable to load it at the desired size,
// rather than loading it at the default size and scaling the
// resulting pixbuf to the desired size.
/*

C function : gdk_pixbuf_format_is_scalable
*/
func (recv *PixbufFormat) IsScalable() bool {
	retC := C.gdk_pixbuf_format_is_scalable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Disables or enables an image format. If a format is disabled,
// gdk-pixbuf won't use the image loader for this format to load
// images. Applications can use this to avoid using image loaders
// with an inappropriate license, see gdk_pixbuf_format_get_license().
/*

C function : gdk_pixbuf_format_set_disabled
*/
func (recv *PixbufFormat) SetDisabled(disabled bool) {
	c_disabled :=
		boolToGboolean(disabled)

	C.gdk_pixbuf_format_set_disabled((*C.GdkPixbufFormat)(recv.native), c_disabled)

	return
}
