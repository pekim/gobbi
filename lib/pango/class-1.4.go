// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : output array param sizes

// A monospace font is a font designed for text display where the the
// characters form a regular grid. For Western languages this would
// mean that the advance width of all characters are the same, but
// this categorization also includes Asian fonts which include
// double-width characters: characters that occupy two grid cells.
// g_unichar_iswide() returns a result that indicates whether a
// character is typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// pango_font_metrics_get_approximate_digit_width(), since the results
// of pango_font_metrics_get_approximate_char_width() may be affected
// by double-width characters.
/*

C function

pango_font_family_is_monospace
*/
func (recv *FontFamily) IsMonospace() bool {
	retC := C.pango_font_family_is_monospace((*C.PangoFontFamily)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the render ID for shape engines for this fontmap.
// See the <structfield>render_type</structfield> field of
// #PangoEngineInfo.
/*

C function

pango_font_map_get_shape_engine_type
*/
func (recv *FontMap) GetShapeEngineType() string {
	retC := C.pango_font_map_get_shape_engine_type((*C.PangoFontMap)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc (PangoFontsetForeachFunc) for param func

// Gets whether to calculate the bidirectional base direction
// for the layout according to the contents of the layout.
// See pango_layout_set_auto_dir().
/*

C function

pango_layout_get_auto_dir
*/
func (recv *Layout) GetAutoDir() bool {
	retC := C.pango_layout_get_auto_dir((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether to calculate the bidirectional base direction
// for the layout according to the contents of the layout;
// when this flag is on (the default), then paragraphs in
// @layout that begin with strong right-to-left characters
// (Arabic and Hebrew principally), will have right-to-left
// layout, paragraphs with letters from other scripts will
// have left-to-right layout. Paragraphs with only neutral
// characters get their direction from the surrounding paragraphs.
//
// When %FALSE, the choice between left-to-right and
// right-to-left layout is done according to the base direction
// of the layout's #PangoContext. (See pango_context_set_base_dir()).
//
// When the auto-computed direction of a paragraph differs from the
// base direction of the context, the interpretation of
// %PANGO_ALIGN_LEFT and %PANGO_ALIGN_RIGHT are swapped.
/*

C function

pango_layout_set_auto_dir
*/
func (recv *Layout) SetAutoDir(autoDir bool) {
	c_auto_dir :=
		boolToGboolean(autoDir)

	C.pango_layout_set_auto_dir((*C.PangoLayout)(recv.native), c_auto_dir)

	return
}
