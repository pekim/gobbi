// Code generated - DO NOT EDIT.

package gdkpixbuf

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

// UNSUPPORTED : get_file_info_async : has callback

// UNSUPPORTED : new_from_stream_async : has callback

// UNSUPPORTED : new_from_stream_at_scale_async : has callback

func Fn_gdk_pixbuf_animation_new_from_file(param0 string) {}

// UNSUPPORTED : new_from_stream_async : has callback

func Fn_gdk_pixbuf_loader_new() {
	C.gdk_pixbuf_loader_new()
}

func Fn_gdk_pixbuf_loader_new_with_type(param0 string) {}
