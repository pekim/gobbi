// This is a generated file - DO NOT EDIT
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Gets the custom font map of this font chooser widget,
// or %NULL if it does not have one.
/*

C function : gtk_font_chooser_get_font_map
*/
func (recv *FontChooser) GetFontMap() *pango.FontMap {
	retC := C.gtk_font_chooser_get_font_map((*C.GtkFontChooser)(recv.native))
	var retGo (*pango.FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets a custom font map to use for this font chooser widget.
// A custom font map can be used to present application-specific
// fonts instead of or in addition to the normal system fonts.
//
// |[<!-- language="C" -->
// FcConfig *config;
// PangoFontMap *fontmap;
//
// config = FcInitLoadConfigAndFonts ();
// FcConfigAppFontAddFile (config, my_app_font_file);
//
// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
// gtk_font_chooser_set_font_map (font_chooser, fontmap);
// ]|
//
// Note that other GTK+ widgets will only be able to use the application-specific
// font if it is present in the font map they use:
//
// |[
// context = gtk_widget_get_pango_context (label);
// pango_context_set_font_map (context, fontmap);
// ]|
/*

C function : gtk_font_chooser_set_font_map
*/
func (recv *FontChooser) SetFontMap(fontmap *pango.FontMap) {
	c_fontmap := (*C.PangoFontMap)(C.NULL)
	if fontmap != nil {
		c_fontmap = (*C.PangoFontMap)(fontmap.ToC())
	}

	C.gtk_font_chooser_set_font_map((*C.GtkFontChooser)(recv.native), c_fontmap)

	return
}
