// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.12

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

static gboolean c_gdk_pixbuf_save(GdkPixbuf* pixbuf, const char* filename, const char* type, GError** error) {
    return gdk_pixbuf_save(pixbuf, filename, type, error, NULL);
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

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

func Fn_gdk_pixbuf_format_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_description(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

func Fn_gdk_pixbuf_format_get_license(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_license(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_format_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_get_name(cValueInstance)

	return C.GoString(ret)
}

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

func Fn_gdk_pixbuf_format_is_writable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_format_is_writable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_set_disabled(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkPixbufFormat)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_pixbuf_format_set_disabled(cValueInstance, cValue0)
}

func Fn_gdk_pixbuf_new(param0 int, param1 bool, param2 int, param3 int, param4 int) unsafe.Pointer {
	cValue0 := (C.GdkColorspace)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	ret := C.gdk_pixbuf_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

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

func Fn_gdk_pixbuf_new_from_file_at_size(param0 string, param1 int, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file_at_size(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_inline(param0 int, param1 []uint8, param2 bool, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := (*C.guint8)(unsafe.Pointer(&param1[0]))

	cValue2 := toCBool(param2)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_inline(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_xpm_data : parameter 'data' is array parameter without length parameter

func Fn_gdk_pixbuf_add_alpha(paramInstance unsafe.Pointer, param0 bool, param1 uint8, param2 uint8, param3 uint8) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := (C.guchar)(param1)

	cValue2 := (C.guchar)(param2)

	cValue3 := (C.guchar)(param3)

	ret := C.gdk_pixbuf_add_alpha(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_apply_embedded_orientation(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_apply_embedded_orientation(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_composite(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.double)(param5)

	cValue6 := (C.double)(param6)

	cValue7 := (C.double)(param7)

	cValue8 := (C.double)(param8)

	cValue9 := (C.GdkInterpType)(param9)

	cValue10 := (C.int)(param10)

	C.gdk_pixbuf_composite(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10)
}

func Fn_gdk_pixbuf_composite_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int, param11 int, param12 int, param13 int, param14 uint32, param15 uint32) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.double)(param5)

	cValue6 := (C.double)(param6)

	cValue7 := (C.double)(param7)

	cValue8 := (C.double)(param8)

	cValue9 := (C.GdkInterpType)(param9)

	cValue10 := (C.int)(param10)

	cValue11 := (C.int)(param11)

	cValue12 := (C.int)(param12)

	cValue13 := (C.int)(param13)

	cValue14 := (C.guint32)(param14)

	cValue15 := (C.guint32)(param15)

	C.gdk_pixbuf_composite_color(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cValue11, cValue12, cValue13, cValue14, cValue15)
}

func Fn_gdk_pixbuf_composite_color_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int, param5 uint32, param6 uint32) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.GdkInterpType)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.guint32)(param5)

	cValue6 := (C.guint32)(param6)

	ret := C.gdk_pixbuf_composite_color_simple(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy_area(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 int, param6 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (*C.GdkPixbuf)(unsafe.Pointer(param4))

	cValue5 := (C.int)(param5)

	cValue6 := (C.int)(param6)

	C.gdk_pixbuf_copy_area(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gdk_pixbuf_fill(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gdk_pixbuf_fill(cValueInstance, cValue0)
}

func Fn_gdk_pixbuf_flip(paramInstance unsafe.Pointer, param0 bool) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gdk_pixbuf_flip(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_get_bits_per_sample(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_bits_per_sample(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_colorspace(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_colorspace(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_has_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_has_alpha(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_get_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_height(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_n_channels(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_n_channels(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_option(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gdk_pixbuf_get_option(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_get_rowstride(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_rowstride(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_get_width(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_new_subpixbuf(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	ret := C.gdk_pixbuf_new_subpixbuf(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_rotate_simple(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkPixbufRotation)(param0)

	ret := C.gdk_pixbuf_rotate_simple(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_saturate_and_pixelate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float32, param2 bool) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	cValue1 := (C.gfloat)(param1)

	cValue2 := toCBool(param2)

	C.gdk_pixbuf_saturate_and_pixelate(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_pixbuf_save(paramInstance unsafe.Pointer, param0 string, param1 string, param2 *unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (**C.GError)(unsafe.Pointer(param2))

	ret := C.c_gdk_pixbuf_save(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_streamv : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'option_keys' is array parameter without length parameter

// UNSUPPORTED : gdk_pixbuf_savev : parameter 'option_keys' is array parameter without length parameter

func Fn_gdk_pixbuf_scale(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.double)(param5)

	cValue6 := (C.double)(param6)

	cValue7 := (C.double)(param7)

	cValue8 := (C.double)(param8)

	cValue9 := (C.GdkInterpType)(param9)

	C.gdk_pixbuf_scale(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gdk_pixbuf_scale_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.GdkInterpType)(param2)

	ret := C.gdk_pixbuf_scale_simple(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_set_option(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gdk_pixbuf_set_option(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_unref(cValueInstance)
}

func Fn_gdk_pixbuf_get_file_info(param0 string, param1 *int, param2 *int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	ret := C.gdk_pixbuf_get_file_info(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_get_formats() unsafe.Pointer {
	ret := C.gdk_pixbuf_get_formats()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_get_height(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_get_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	ret := C.gdk_pixbuf_animation_get_iter(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_static_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_get_static_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_get_width(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_is_static_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_is_static_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_unref(cValueInstance)
}

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_animation_iter_advance(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	ret := C.gdk_pixbuf_animation_iter_advance(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_iter_get_delay_time(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_iter_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_animation_iter_on_currently_loading_frame(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_new() unsafe.Pointer {
	ret := C.gdk_pixbuf_loader_new()

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_new_with_mime_type(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_new_with_mime_type(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_new_with_type(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_close(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_close(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_get_animation(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_loader_get_animation(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_get_format(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_loader_get_format(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	ret := C.gdk_pixbuf_loader_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_set_size(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	C.gdk_pixbuf_loader_set_size(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_pixbuf_loader_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_write(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_simple_anim_new(param0 int, param1 int, param2 float32) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gfloat)(param2)

	ret := C.gdk_pixbuf_simple_anim_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_simple_anim_add_frame(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gdk_pixbuf_simple_anim_add_frame(cValueInstance, cValue0)
}
