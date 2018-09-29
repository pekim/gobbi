// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

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

// CreateContext is a wrapper around the C function pango_font_map_create_context.
func (recv *FontMap) CreateContext() *Context {
	retC := C.pango_font_map_create_context((*C.PangoFontMap)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// GetBaseline is a wrapper around the C function pango_layout_get_baseline.
func (recv *Layout) GetBaseline() int32 {
	retC := C.pango_layout_get_baseline((*C.PangoLayout)(recv.native))
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

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*

// DrawGlyphItem is a wrapper around the C function pango_renderer_draw_glyph_item.
func (recv *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int32, y int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_glyph_item := (*C.PangoGlyphItem)(glyphItem.ToC())

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_glyph_item((*C.PangoRenderer)(recv.native), c_text, c_glyph_item, c_x, c_y)

	return
}
