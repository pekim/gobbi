// Code generated - DO NOT EDIT.

package pangoft2

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'pango_ft2_font_get_coverage' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_font_get_face' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_font_get_kerning' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_get_context' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_get_unknown_glyph' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render_layout' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render_layout_line' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render_layout_line_subpixel' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render_layout_subpixel' : non trivial function

// UNSUPPORTED : C value 'pango_ft2_render_transformed' : non trivial function

var shutdownDisplayInvoker *gi.Function

// ShutdownDisplay is a representation of the C type pango_ft2_shutdown_display.
func ShutdownDisplay() {
	if shutdownDisplayInvoker == nil {
		shutdownDisplayInvoker = gi.FunctionInvokerNew("PangoFT2", "shutdown_display")
	}

	shutdownDisplayInvoker.Invoke()
}
