// This is a generated file - DO NOT EDIT
// +build pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Changed is a wrapper around the C function pango_context_changed.
func (recv *Context) Changed() {
	C.pango_context_changed((*C.PangoContext)(recv.native))

	return
}

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// GetSerial is a wrapper around the C function pango_context_get_serial.
func (recv *Context) GetSerial() uint32 {
	retC := C.pango_context_get_serial((*C.PangoContext)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : pango_context_list_families : unsupported parameter families : no param type

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : no param type

// GetSerial is a wrapper around the C function pango_font_map_get_serial.
func (recv *FontMap) GetSerial() uint32 {
	retC := C.pango_font_map_get_serial((*C.PangoFontMap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : no param type

// Unsupported : pango_layout_get_log_attrs_readonly : unsupported parameter n_attrs : no type generator for gint, gint*

// Unsupported : pango_layout_get_pixel_size : unsupported parameter width : no type generator for gint, int*

// GetSerial is a wrapper around the C function pango_layout_get_serial.
func (recv *Layout) GetSerial() uint32 {
	retC := C.pango_layout_get_serial((*C.PangoLayout)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_index_to_line_x : unsupported parameter line : no type generator for gint, int*

// Unsupported : pango_layout_move_cursor_visually : unsupported parameter new_index : no type generator for gint, int*

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*
