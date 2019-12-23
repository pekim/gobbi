// Code generated - DO NOT EDIT.

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

func Fn_gdk_pixbuf_error_quark() {

}

func Fn_gdk_pixbuf_new(param0 int, param1 bool, param2 int, param3 int, param4 int) {
	cValue0 := (C.GdkColorspace)(param0)
	cValue1 := (C.gboolean)(param1)
	cValue2 := (C.int)(param2)
	cValue3 := (C.int)(param3)
	cValue4 := (C.int)(param4)

}

// UNSUPPORTED : new_from_data : has callback

func Fn_gdk_pixbuf_new_from_file(param0 string) {
	cValue0 := 42

}

func Fn_gdk_pixbuf_new_from_inline(param0 int, param1 []uint8, param2 bool) {
	// has array param
}

func Fn_gdk_pixbuf_new_from_xpm_data(param0 []string) {
	// has array param
}

func Fn_gdk_pixbuf_add_alpha(paramInstance unsafe.Pointer, param0 bool, param1 uint8, param2 uint8, param3 uint8) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gboolean)(param0)
	cValue1 := (C.guchar)(param1)
	cValue2 := (C.guchar)(param2)
	cValue3 := (C.guchar)(param3)

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

}

func Fn_gdk_pixbuf_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

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

}

func Fn_gdk_pixbuf_fill(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

}

func Fn_gdk_pixbuf_get_bits_per_sample(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_colorspace(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_has_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_n_channels(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_option(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := 42

}

func Fn_gdk_pixbuf_get_pixels(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_rowstride(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_new_subpixbuf(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.int)(param2)
	cValue3 := (C.int)(param3)

}

func Fn_gdk_pixbuf_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_saturate_and_pixelate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float32, param2 bool) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))
	cValue1 := (C.gfloat)(param1)
	cValue2 := (C.gboolean)(param2)

}

// UNSUPPORTED : save : has varargs

// UNSUPPORTED : save_to_buffer : has varargs

// UNSUPPORTED : save_to_callback : has varargs

// UNSUPPORTED : save_to_callbackv : has callback

// UNSUPPORTED : save_to_stream : has varargs

// UNSUPPORTED : save_to_stream_async : has varargs

// UNSUPPORTED : save_to_streamv_async : has callback

func Fn_gdk_pixbuf_savev(paramInstance unsafe.Pointer, param0 string, param1 string, param2 []string, param3 []string) {
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

}

func Fn_gdk_pixbuf_scale_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.GdkInterpType)(param2)

}

func Fn_gdk_pixbuf_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbuf)(unsafe.Pointer(paramInstance))

}

// UNSUPPORTED : get_file_info_async : has callback

// UNSUPPORTED : new_from_stream_async : has callback

// UNSUPPORTED : new_from_stream_at_scale_async : has callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string) {
	cValue0 := 42

}

func Fn_gdk_pixbuf_animation_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_get_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

}

func Fn_gdk_pixbuf_animation_get_static_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_is_static_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_ref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimation)(unsafe.Pointer(paramInstance))

}

// UNSUPPORTED : new_from_stream_async : has callback

func Fn_gdk_pixbuf_animation_iter_advance(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_loader_new() {

}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string) {
	cValue0 := 42

}

func Fn_gdk_pixbuf_loader_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_loader_get_animation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_loader_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkPixbufLoader)(unsafe.Pointer(paramInstance))

}

func Fn_gdk_pixbuf_loader_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64) {
	// has array param
}
