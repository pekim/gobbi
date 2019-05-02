// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AboutDialogNew is a wrapper around the C function gtk_about_dialog_new.
func AboutDialogNew() *AboutDialog {
	retC := C.gtk_about_dialog_new()
	retGo := AboutDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererComboNew is a wrapper around the C function gtk_cell_renderer_combo_new.
func CellRendererComboNew() *CellRendererCombo {
	retC := C.gtk_cell_renderer_combo_new()
	retGo := CellRendererComboNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererProgressNew is a wrapper around the C function gtk_cell_renderer_progress_new.
func CellRendererProgressNew() *CellRendererProgress {
	retC := C.gtk_cell_renderer_progress_new()
	retGo := CellRendererProgressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNew is a wrapper around the C function gtk_cell_view_new.
func CellViewNew() *CellView {
	retC := C.gtk_cell_view_new()
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithContext is a wrapper around the C function gtk_cell_view_new_with_context.
func CellViewNewWithContext(area *CellArea, context *CellAreaContext) *CellView {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	retC := C.gtk_cell_view_new_with_context(c_area, c_context)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithMarkup is a wrapper around the C function gtk_cell_view_new_with_markup.
func CellViewNewWithMarkup(markup string) *CellView {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	retC := C.gtk_cell_view_new_with_markup(c_markup)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithPixbuf is a wrapper around the C function gtk_cell_view_new_with_pixbuf.
func CellViewNewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *CellView {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_cell_view_new_with_pixbuf(c_pixbuf)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellViewNewWithText is a wrapper around the C function gtk_cell_view_new_with_text.
func CellViewNewWithText(text string) *CellView {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_cell_view_new_with_text(c_text)
	retGo := CellViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileChooserButtonNew is a wrapper around the C function gtk_file_chooser_button_new.
func FileChooserButtonNew(title string, action FileChooserAction) *FileChooserButton {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_button_new(c_title, c_action)
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileChooserButtonNewWithDialog is a wrapper around the C function gtk_file_chooser_button_new_with_dialog.
func FileChooserButtonNewWithDialog(dialog *Dialog) *FileChooserButton {
	c_dialog := (*C.GtkWidget)(C.NULL)
	if dialog != nil {
		c_dialog = (*C.GtkWidget)(dialog.ToC())
	}

	retC := C.gtk_file_chooser_button_new_with_dialog(c_dialog)
	retGo := FileChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconViewNew is a wrapper around the C function gtk_icon_view_new.
func IconViewNew() *IconView {
	retC := C.gtk_icon_view_new()
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconViewNewWithModel is a wrapper around the C function gtk_icon_view_new_with_model.
func IconViewNewWithModel(model *TreeModel) *IconView {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_icon_view_new_with_model(c_model)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromIconName is a wrapper around the C function gtk_image_new_from_icon_name.
func ImageNewFromIconName(iconName string, size IconSize) *Image {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_icon_name(c_icon_name, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuToolButtonNew is a wrapper around the C function gtk_menu_tool_button_new.
func MenuToolButtonNew(iconWidget *Widget, label string) *MenuToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_tool_button_new(c_icon_widget, c_label)
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuToolButtonNewFromStock is a wrapper around the C function gtk_menu_tool_button_new_from_stock.
func MenuToolButtonNewFromStock(stockId string) *MenuToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_menu_tool_button_new_from_stock(c_stock_id)
	retGo := MenuToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : STOCK_ABOUT

// Blacklisted : STOCK_CONNECT

// Blacklisted : STOCK_DIRECTORY

// Blacklisted : STOCK_DISCONNECT

// Blacklisted : STOCK_EDIT

// Blacklisted : STOCK_FILE

// Blacklisted : STOCK_MEDIA_FORWARD

// Blacklisted : STOCK_MEDIA_NEXT

// Blacklisted : STOCK_MEDIA_PAUSE

// Blacklisted : STOCK_MEDIA_PLAY

// Blacklisted : STOCK_MEDIA_PREVIOUS

// Blacklisted : STOCK_MEDIA_RECORD

// Blacklisted : STOCK_MEDIA_REWIND

// Blacklisted : STOCK_MEDIA_STOP

// Unsupported : gtk_init_with_args : unsupported parameter entries :

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs
