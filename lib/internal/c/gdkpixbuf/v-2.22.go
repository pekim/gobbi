// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36 gdkpixbuf_2.36.8 gdkpixbuf_2.40

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func Fn_gdk_pixbuf_format_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_format_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_format_free(cValueInstance)
}

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_streamv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_savev : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback
