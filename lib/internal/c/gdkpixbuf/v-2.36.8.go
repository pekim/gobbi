// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.36.8 gdkpixbuf_2.40

package gdkpixbuf

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

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

func Fn_gdk_pixbuf_calculate_rowstride(param0 int, param1 bool, param2 int, param3 int, param4 int) int {
	cValue0 := (C.GdkColorspace)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	ret := C.gdk_pixbuf_calculate_rowstride(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (int)(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback
