// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36 gdkpixbuf_2.36.8 gdkpixbuf_2.40

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_new_from_resource(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_resource(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_resource_at_scale(param0 string, param1 int, param2 int, param3 bool, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := toCBool(param3)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_resource_at_scale(cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_get_byte_length(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_byte_length(cValueInstance)

	return (uint64)(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_get_pixels_with_length(paramInstance unsafe.Pointer, param0 *uint) []uint8 {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	ret := C.gdk_pixbuf_get_pixels_with_length(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
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
