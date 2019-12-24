// Code generated - DO NOT EDIT.

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

func Fn_gdk_pixbuf_error_quark() {

	C.gdk_pixbuf_error_quark()
}

func Fn_gdk_pixbuf_new(param0 int, param1 bool, param2 int, param3 int, param4 int) {
	cValue0 := (C.GdkColorspace)(param0)
	cValue1 := toCBool(param1)
	cValue2 := (C.int)(param2)
	cValue3 := (C.int)(param3)
	cValue4 := (C.int)(param4)

	C.gdk_pixbuf_new(cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : new_from_data : has callback

func Fn_gdk_pixbuf_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.gdk_pixbuf_new_from_file(cValue0, cError)
}

func Fn_gdk_pixbuf_new_from_inline(param0 int, param1 []uint8, param2 bool, error unsafe.Pointer) {
	// has array param
}

func Fn_gdk_pixbuf_new_from_xpm_data(param0 []string) {
	// has array param
}

func Fn_gdk_pixbuf_add_alpha(paramInstance unsafe.Pointer, param0 bool, param1 uint8, param2 uint8, param3 uint8) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)
	cValue1 := (C.guchar)(param1)
	cValue2 := (C.guchar)(param2)
	cValue3 := (C.guchar)(param3)

	C.gdk_pixbuf_add_alpha(cValueInstance, cValue0, cValue1, cValue2, cValue3)
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

func Fn_gdk_pixbuf_composite_color_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int, param5 uint32, param6 uint32) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.GdkInterpType)(param2)
	cValue3 := (C.int)(param3)
	cValue4 := (C.int)(param4)
	cValue5 := (C.guint32)(param5)
	cValue6 := (C.guint32)(param6)

	C.gdk_pixbuf_composite_color_simple(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gdk_pixbuf_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_copy(cValueInstance)
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

func Fn_gdk_pixbuf_get_bits_per_sample(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_bits_per_sample(cValueInstance)
}

func Fn_gdk_pixbuf_get_colorspace(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_colorspace(cValueInstance)
}

func Fn_gdk_pixbuf_get_has_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_has_alpha(cValueInstance)
}

func Fn_gdk_pixbuf_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_height(cValueInstance)
}

func Fn_gdk_pixbuf_get_n_channels(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_n_channels(cValueInstance)
}

func Fn_gdk_pixbuf_get_option(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_pixbuf_get_option(cValueInstance, cValue0)
}

func Fn_gdk_pixbuf_get_pixels(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_pixels(cValueInstance)
}

func Fn_gdk_pixbuf_get_rowstride(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_rowstride(cValueInstance)
}

func Fn_gdk_pixbuf_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_get_width(cValueInstance)
}

func Fn_gdk_pixbuf_new_subpixbuf(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.int)(param2)
	cValue3 := (C.int)(param3)

	C.gdk_pixbuf_new_subpixbuf(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_pixbuf_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_ref(cValueInstance)
}

func Fn_gdk_pixbuf_saturate_and_pixelate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float32, param2 bool) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))
	cValue1 := (C.gfloat)(param1)
	cValue2 := toCBool(param2)

	C.gdk_pixbuf_saturate_and_pixelate(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : save : has varargs

// UNSUPPORTED : save_to_buffer : has varargs

// UNSUPPORTED : save_to_callback : has varargs

// UNSUPPORTED : save_to_callbackv : has callback

// UNSUPPORTED : save_to_stream : has varargs

// UNSUPPORTED : save_to_stream_async : has varargs

// UNSUPPORTED : save_to_streamv_async : has callback

func Fn_gdk_pixbuf_savev(paramInstance unsafe.Pointer, param0 string, param1 string, param2 []string, param3 []string, error unsafe.Pointer) {
	// has array param
}

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

func Fn_gdk_pixbuf_scale_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.GdkInterpType)(param2)

	C.gdk_pixbuf_scale_simple(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_pixbuf_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_unref(cValueInstance)
}

// UNSUPPORTED : get_file_info_async : has callback

// UNSUPPORTED : new_from_stream_async : has callback

// UNSUPPORTED : new_from_stream_at_scale_async : has callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.gdk_pixbuf_animation_new_from_file(cValue0, cError)
}

func Fn_gdk_pixbuf_animation_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_get_height(cValueInstance)
}

func Fn_gdk_pixbuf_animation_get_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	C.gdk_pixbuf_animation_get_iter(cValueInstance, cValue0)
}

func Fn_gdk_pixbuf_animation_get_static_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_get_static_image(cValueInstance)
}

func Fn_gdk_pixbuf_animation_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_get_width(cValueInstance)
}

func Fn_gdk_pixbuf_animation_is_static_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_is_static_image(cValueInstance)
}

func Fn_gdk_pixbuf_animation_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_ref(cValueInstance)
}

func Fn_gdk_pixbuf_animation_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_unref(cValueInstance)
}

// UNSUPPORTED : new_from_stream_async : has callback

func Fn_gdk_pixbuf_animation_iter_advance(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	C.gdk_pixbuf_animation_iter_advance(cValueInstance, cValue0)
}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_iter_get_delay_time(cValueInstance)
}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_iter_get_pixbuf(cValueInstance)
}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_animation_iter_on_currently_loading_frame(cValueInstance)
}

func Fn_gdk_pixbuf_loader_new() {

	C.gdk_pixbuf_loader_new()
}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.gdk_pixbuf_loader_new_with_type(cValue0, cError)
}

func Fn_gdk_pixbuf_loader_close(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.gdk_pixbuf_loader_close(cValueInstance, cError)
}

func Fn_gdk_pixbuf_loader_get_animation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_loader_get_animation(cValueInstance)
}

func Fn_gdk_pixbuf_loader_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

	C.gdk_pixbuf_loader_get_pixbuf(cValueInstance)
}

func Fn_gdk_pixbuf_loader_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, error unsafe.Pointer) {
	// has array param
}
