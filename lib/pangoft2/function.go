// Code generated - DO NOT EDIT.

package pangoft2

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'pango_ft2_font_get_coverage' : parameter 'font' of type 'Pango.Font' not supported

// UNSUPPORTED : C value 'pango_ft2_font_get_face' : parameter 'font' of type 'Pango.Font' not supported

// UNSUPPORTED : C value 'pango_ft2_font_get_kerning' : parameter 'font' of type 'Pango.Font' not supported

// UNSUPPORTED : C value 'pango_ft2_get_context' : parameter 'dpi_x' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_ft2_get_unknown_glyph' : parameter 'font' of type 'Pango.Font' not supported

// UNSUPPORTED : C value 'pango_ft2_render' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_line' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_line_subpixel' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_subpixel' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_transformed' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

var shutdownDisplayInvoker *gi.Function

// ShutdownDisplay is a representation of the C type pango_ft2_shutdown_display.
func ShutdownDisplay() {
	if shutdownDisplayInvoker == nil {
		shutdownDisplayInvoker = gi.FunctionInvokerNew("PangoFT2", "shutdown_display")
	}

	shutdownDisplayInvoker.Invoke(nil, nil)

}
