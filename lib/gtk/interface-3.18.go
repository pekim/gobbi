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

// GetFontMap is a wrapper around the C function gtk_font_chooser_get_font_map.
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

// SetFontMap is a wrapper around the C function gtk_font_chooser_set_font_map.
func (recv *FontChooser) SetFontMap(fontmap *pango.FontMap) {
	c_fontmap := (*C.PangoFontMap)(C.NULL)
	if fontmap != nil {
		c_fontmap = (*C.PangoFontMap)(fontmap.ToC())
	}

	C.gtk_font_chooser_set_font_map((*C.GtkFontChooser)(recv.native), c_fontmap)

	return
}
