// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_format_get_extensions : no return type

// GetLicense is a wrapper around the C function gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() string {
	retC := C.gdk_pixbuf_format_get_license((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_format_get_mime_types : no return type

// IsDisabled is a wrapper around the C function gdk_pixbuf_format_is_disabled.
func (recv *PixbufFormat) IsDisabled() bool {
	retC := C.gdk_pixbuf_format_is_disabled((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsScalable is a wrapper around the C function gdk_pixbuf_format_is_scalable.
func (recv *PixbufFormat) IsScalable() bool {
	retC := C.gdk_pixbuf_format_is_scalable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDisabled is a wrapper around the C function gdk_pixbuf_format_set_disabled.
func (recv *PixbufFormat) SetDisabled(disabled bool) {
	c_disabled :=
		boolToGboolean(disabled)

	C.gdk_pixbuf_format_set_disabled((*C.GdkPixbufFormat)(recv.native), c_disabled)

	return
}
