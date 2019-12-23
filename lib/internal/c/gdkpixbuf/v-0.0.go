// Code generated - DO NOT EDIT.

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

func Fn_gdk_pixbuf_error_quark() {
	C.gdk_pixbuf_error_quark()
}

func Fn_gdk_pixbuf_new(param0 int, param1 bool, param2 int, param3 int, param4 int) {}

// UNSUPPORTED : new_from_data : has callback

func Fn_gdk_pixbuf_new_from_file(param0 string) {}

func Fn_gdk_pixbuf_new_from_inline(param0 int, param1 []uint8, param2 bool) {}

func Fn_gdk_pixbuf_new_from_xpm_data(param0 []string) {}

func Fn_gdk_pixbuf_add_alpha(paramInstance unsafe.Pointer, param0 bool, param1 uint8, param2 uint8, param3 uint8) {
}

func Fn_gdk_pixbuf_composite(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int) {
}

func Fn_gdk_pixbuf_composite_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int, param11 int, param12 int, param13 int, param14 uint32, param15 uint32) {
}

func Fn_gdk_pixbuf_composite_color_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int, param5 uint32, param6 uint32) {
}

func Fn_gdk_pixbuf_copy(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_copy_area(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 int, param6 int) {
}

func Fn_gdk_pixbuf_fill(paramInstance unsafe.Pointer, param0 uint32) {}

func Fn_gdk_pixbuf_get_bits_per_sample(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_colorspace(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_has_alpha(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_height(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_n_channels(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_option(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gdk_pixbuf_get_pixels(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_rowstride(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_get_width(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_new_subpixbuf(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
}

func Fn_gdk_pixbuf_ref(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_saturate_and_pixelate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float32, param2 bool) {
}

// UNSUPPORTED : save : has varargs

// UNSUPPORTED : save_to_buffer : has varargs

// UNSUPPORTED : save_to_callback : has varargs

// UNSUPPORTED : save_to_callbackv : has callback

// UNSUPPORTED : save_to_stream : has varargs

// UNSUPPORTED : save_to_stream_async : has varargs

// UNSUPPORTED : save_to_streamv_async : has callback

func Fn_gdk_pixbuf_savev(paramInstance unsafe.Pointer, param0 string, param1 string, param2 []string, param3 []string) {
}

func Fn_gdk_pixbuf_scale(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int) {
}

func Fn_gdk_pixbuf_scale_simple(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {}

func Fn_gdk_pixbuf_unref(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : get_file_info_async : has callback

// UNSUPPORTED : new_from_stream_async : has callback

// UNSUPPORTED : new_from_stream_at_scale_async : has callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string) {}

func Fn_gdk_pixbuf_animation_get_height(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_get_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_get_static_image(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_get_width(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_is_static_image(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_ref(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_unref(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : new_from_stream_async : has callback

func Fn_gdk_pixbuf_animation_iter_advance(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_loader_new() {
	C.gdk_pixbuf_loader_new()
}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string) {}

func Fn_gdk_pixbuf_loader_close(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_loader_get_animation(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_loader_get_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gdk_pixbuf_loader_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64) {}
