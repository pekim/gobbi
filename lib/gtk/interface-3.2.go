// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// Gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two
// font descriptions.
/*

C function

gtk_font_chooser_get_font
*/
func (recv *FontChooser) GetFont() string {
	retC := C.gtk_font_chooser_get_font((*C.GtkFontChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two
// font descriptions.
/*

C function

gtk_font_chooser_get_font_desc
*/
func (recv *FontChooser) GetFontDesc() *pango.FontDescription {
	retC := C.gtk_font_chooser_get_font_desc((*C.GtkFontChooser)(recv.native))
	var retGo (*pango.FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the #PangoFontFace representing the selected font group
// details (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns %NULL.
/*

C function

gtk_font_chooser_get_font_face
*/
func (recv *FontChooser) GetFontFace() *pango.FontFace {
	retC := C.gtk_font_chooser_get_font_face((*C.GtkFontChooser)(recv.native))
	var retGo (*pango.FontFace)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontFaceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the #PangoFontFamily representing the selected font family.
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns %NULL.
/*

C function

gtk_font_chooser_get_font_family
*/
func (recv *FontChooser) GetFontFamily() *pango.FontFamily {
	retC := C.gtk_font_chooser_get_font_family((*C.GtkFontChooser)(recv.native))
	var retGo (*pango.FontFamily)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontFamilyNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// The selected font size.
/*

C function

gtk_font_chooser_get_font_size
*/
func (recv *FontChooser) GetFontSize() int32 {
	retC := C.gtk_font_chooser_get_font_size((*C.GtkFontChooser)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the text displayed in the preview area.
/*

C function

gtk_font_chooser_get_preview_text
*/
func (recv *FontChooser) GetPreviewText() string {
	retC := C.gtk_font_chooser_get_preview_text((*C.GtkFontChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the preview entry is shown or not.
/*

C function

gtk_font_chooser_get_show_preview_entry
*/
func (recv *FontChooser) GetShowPreviewEntry() bool {
	retC := C.gtk_font_chooser_get_show_preview_entry((*C.GtkFontChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_font_chooser_set_filter_func : unsupported parameter filter : no type generator for FontFilterFunc (GtkFontFilterFunc) for param filter

// Sets the currently-selected font.
/*

C function

gtk_font_chooser_set_font
*/
func (recv *FontChooser) SetFont(fontname string) {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	C.gtk_font_chooser_set_font((*C.GtkFontChooser)(recv.native), c_fontname)

	return
}

// Sets the currently-selected font from @font_desc.
/*

C function

gtk_font_chooser_set_font_desc
*/
func (recv *FontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	c_font_desc := (*C.PangoFontDescription)(C.NULL)
	if fontDesc != nil {
		c_font_desc = (*C.PangoFontDescription)(fontDesc.ToC())
	}

	C.gtk_font_chooser_set_font_desc((*C.GtkFontChooser)(recv.native), c_font_desc)

	return
}

// Sets the text displayed in the preview area.
// The @text is used to show how the selected font looks.
/*

C function

gtk_font_chooser_set_preview_text
*/
func (recv *FontChooser) SetPreviewText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_font_chooser_set_preview_text((*C.GtkFontChooser)(recv.native), c_text)

	return
}

// Shows or hides the editable preview entry.
/*

C function

gtk_font_chooser_set_show_preview_entry
*/
func (recv *FontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	c_show_preview_entry :=
		boolToGboolean(showPreviewEntry)

	C.gtk_font_chooser_set_show_preview_entry((*C.GtkFontChooser)(recv.native), c_show_preview_entry)

	return
}
