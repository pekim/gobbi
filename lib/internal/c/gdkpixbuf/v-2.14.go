// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.14

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

static gboolean c_gdk_pixbuf_save_to_stream(GdkPixbuf* pixbuf, GOutputStream* stream, const char* type, GCancellable* cancellable, GError** error) {
    return gdk_pixbuf_save_to_stream(pixbuf, stream, type, cancellable, error, NULL);
}
*/
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

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_new_from_stream(param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_stream(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_stream_at_scale(param0 unsafe.Pointer, param1 int, param2 int, param3 bool, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := toCBool(param3)

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_stream_at_scale(cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

func Fn_gdk_pixbuf_save_to_stream(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 *unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cValue3 := (**C.GError)(unsafe.Pointer(param3))

	ret := C.c_gdk_pixbuf_save_to_stream(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_streamv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_savev : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback
