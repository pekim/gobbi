// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.32 gdkpixbuf_2.36 gdkpixbuf_2.36.8 gdkpixbuf_2.40

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_new_from_bytes(param0 unsafe.Pointer, param1 int, param2 bool, param3 int, param4 int, param5 int, param6 int) unsafe.Pointer {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cValue1 := (C.GdkColorspace)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.int)(param5)

	cValue6 := (C.int)(param6)

	ret := C.gdk_pixbuf_new_from_bytes(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_get_options(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_options(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_read_pixel_bytes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_read_pixel_bytes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_read_pixels(paramInstance unsafe.Pointer) *uint8 {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_read_pixels(cValueInstance)

	return (*uint8)(ret)
}

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_streamv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_savev : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_get_file_info_finish(param0 unsafe.Pointer, param1 *int, param2 *int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_get_file_info_finish(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback
