// Code generated - DO NOT EDIT.

package pangoft2

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'pango_ft2_font_get_coverage' : has parameters

// UNSUPPORTED : C value 'pango_ft2_font_get_face' : has parameters

// UNSUPPORTED : C value 'pango_ft2_font_get_kerning' : has parameters

// UNSUPPORTED : C value 'pango_ft2_get_context' : has parameters

// UNSUPPORTED : C value 'pango_ft2_get_unknown_glyph' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render_layout' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render_layout_line' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render_layout_line_subpixel' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render_layout_subpixel' : has parameters

// UNSUPPORTED : C value 'pango_ft2_render_transformed' : has parameters

var shutdownDisplayInvoker *gi.Function

// ShutdownDisplay is a representation of the C type pango_ft2_shutdown_display.
func ShutdownDisplay() {
	if shutdownDisplayInvoker == nil {
		shutdownDisplayInvoker = gi.FunctionInvokerNew("PangoFT2", "shutdown_display")
	}

	shutdownDisplayInvoker.Invoke()
}
