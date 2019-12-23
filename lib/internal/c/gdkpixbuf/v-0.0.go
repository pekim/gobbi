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

func Fn_gdk_pixbuf_add_alpha(param0 bool, param1 uint8, param2 uint8, param3 uint8) {}

func Fn_gdk_pixbuf_composite(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int) {
}

func Fn_gdk_pixbuf_composite_color(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int, param10 int, param11 int, param12 int, param13 int, param14 uint32, param15 uint32) {
}

func Fn_gdk_pixbuf_composite_color_simple(param0 int, param1 int, param2 int, param3 int, param4 int, param5 uint32, param6 uint32) {
}

func Fn_gdk_pixbuf_copy() {}

func Fn_gdk_pixbuf_copy_area(param0 int, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 int, param6 int) {
}

func Fn_gdk_pixbuf_fill(param0 uint32) {}

func Fn_gdk_pixbuf_get_bits_per_sample() {}

func Fn_gdk_pixbuf_get_colorspace() {}

func Fn_gdk_pixbuf_get_has_alpha() {}

func Fn_gdk_pixbuf_get_height() {}

func Fn_gdk_pixbuf_get_n_channels() {}

func Fn_gdk_pixbuf_get_option(param0 string) {}

func Fn_gdk_pixbuf_get_pixels() {}

func Fn_gdk_pixbuf_get_rowstride() {}

func Fn_gdk_pixbuf_get_width() {}

func Fn_gdk_pixbuf_new_subpixbuf(param0 int, param1 int, param2 int, param3 int) {}

func Fn_gdk_pixbuf_ref() {}

func Fn_gdk_pixbuf_saturate_and_pixelate(param0 unsafe.Pointer, param1 float32, param2 bool) {}

// UNSUPPORTED : save : has varargs

// UNSUPPORTED : save_to_buffer : has varargs

// UNSUPPORTED : save_to_callback : has varargs

// UNSUPPORTED : save_to_callbackv : has callback

// UNSUPPORTED : save_to_stream : has varargs

// UNSUPPORTED : save_to_stream_async : has varargs

// UNSUPPORTED : save_to_streamv_async : has callback

func Fn_gdk_pixbuf_savev(param0 string, param1 string, param2 []string, param3 []string) {}

func Fn_gdk_pixbuf_scale(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int, param5 float64, param6 float64, param7 float64, param8 float64, param9 int) {
}

func Fn_gdk_pixbuf_scale_simple(param0 int, param1 int, param2 int) {}

func Fn_gdk_pixbuf_unref() {}

// UNSUPPORTED : get_file_info_async : has callback

// UNSUPPORTED : new_from_stream_async : has callback

// UNSUPPORTED : new_from_stream_at_scale_async : has callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string) {}

func Fn_gdk_pixbuf_animation_get_height() {}

func Fn_gdk_pixbuf_animation_get_iter(param0 unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_get_static_image() {}

func Fn_gdk_pixbuf_animation_get_width() {}

func Fn_gdk_pixbuf_animation_is_static_image() {}

func Fn_gdk_pixbuf_animation_ref() {}

func Fn_gdk_pixbuf_animation_unref() {}

// UNSUPPORTED : new_from_stream_async : has callback

func Fn_gdk_pixbuf_animation_iter_advance(param0 unsafe.Pointer) {}

func Fn_gdk_pixbuf_animation_iter_get_delay_time() {}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf() {}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame() {}

func Fn_gdk_pixbuf_loader_new() {
	C.gdk_pixbuf_loader_new()
}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string) {}

func Fn_gdk_pixbuf_loader_close() {}

func Fn_gdk_pixbuf_loader_get_animation() {}

func Fn_gdk_pixbuf_loader_get_pixbuf() {}

func Fn_gdk_pixbuf_loader_write(param0 []uint8, param1 uint64) {}
