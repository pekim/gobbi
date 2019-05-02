// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
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

// ApplicationWindowNew is a wrapper around the C function gtk_application_window_new.
func ApplicationWindowNew(application *Application) *ApplicationWindow {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	retC := C.gtk_application_window_new(c_application)
	retGo := ApplicationWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorChooserDialogNew is a wrapper around the C function gtk_color_chooser_dialog_new.
func ColorChooserDialogNew(title string, parent *Window) *ColorChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_color_chooser_dialog_new(c_title, c_parent)
	retGo := ColorChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorChooserWidgetNew is a wrapper around the C function gtk_color_chooser_widget_new.
func ColorChooserWidgetNew() *ColorChooserWidget {
	retC := C.gtk_color_chooser_widget_new()
	retGo := ColorChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromResource is a wrapper around the C function gtk_image_new_from_resource.
func ImageNewFromResource(resourcePath string) *Image {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_image_new_from_resource(c_resource_path)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuNewFromModel is a wrapper around the C function gtk_menu_new_from_model.
func MenuNewFromModel(model *gio.MenuModel) *Menu {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_new_from_model(c_model)
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuBarNewFromModel is a wrapper around the C function gtk_menu_bar_new_from_model.
func MenuBarNewFromModel(model *gio.MenuModel) *MenuBar {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_bar_new_from_model(c_model)
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : output array param accelerator_codes

// SymbolicColorNewWin32 is a wrapper around the C function gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	c_theme_class := C.CString(themeClass)
	defer C.free(unsafe.Pointer(c_theme_class))

	c_id := (C.gint)(id)

	retC := C.gtk_symbolic_color_new_win32(c_theme_class, c_id)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}
