// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_context_list_families : unsupported parameter families : no param type

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : no param type

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// GetHeight is a wrapper around the C function pango_layout_get_height.
func (recv *Layout) GetHeight() int32 {
	retC := C.pango_layout_get_height((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : no param type

// Unsupported : pango_layout_get_log_attrs_readonly : unsupported parameter n_attrs : no type generator for gint, gint*

// Unsupported : pango_layout_get_pixel_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_get_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_index_to_line_x : unsupported parameter line : no type generator for gint, int*

// Unsupported : pango_layout_move_cursor_visually : unsupported parameter new_index : no type generator for gint, int*

// SetHeight is a wrapper around the C function pango_layout_set_height.
func (recv *Layout) SetHeight(height int32) {
	c_height := (C.int)(height)

	C.pango_layout_set_height((*C.PangoLayout)(recv.native), c_height)

	return
}

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*

// GetLayout is a wrapper around the C function pango_renderer_get_layout.
func (recv *Renderer) GetLayout() *Layout {
	retC := C.pango_renderer_get_layout((*C.PangoRenderer)(recv.native))
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLayoutLine is a wrapper around the C function pango_renderer_get_layout_line.
func (recv *Renderer) GetLayoutLine() *LayoutLine {
	retC := C.pango_renderer_get_layout_line((*C.PangoRenderer)(recv.native))
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}
