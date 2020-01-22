// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.2

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}
func Fn_gdk_pixbuf_format_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_description(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_format_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_pixbuf_format_is_writable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_is_writable(cValueInstance)

	return toGoBool(ret)
}

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

func Fn_gdk_pixbuf_set_option(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gdk_pixbuf_set_option(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_get_formats() unsafe.Pointer {
	ret := C.gdk_pixbuf_get_formats()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_loader_get_format(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_loader_get_format(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_set_size(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	C.gdk_pixbuf_loader_set_size(cValueInstance, cValue0, cValue1)
}
