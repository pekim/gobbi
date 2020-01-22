// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.6

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

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

func Fn_gdk_pixbuf_format_get_license(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_license(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_format_is_disabled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_is_disabled(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_is_scalable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_is_scalable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_set_disabled(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_pixbuf_format_set_disabled(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_new_from_file_at_scale(param0 string, param1 int, param2 int, param3 bool, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := toCBool(param3)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file_at_scale(cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_flip(paramInstance unsafe.Pointer, param0 bool) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gdk_pixbuf_flip(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_rotate_simple(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkPixbufRotation)(param0)

	ret := C.gdk_pixbuf_rotate_simple(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback
