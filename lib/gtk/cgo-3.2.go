// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// FontChooserDialogNew is a wrapper around the C function gtk_font_chooser_dialog_new.
func FontChooserDialogNew(title string, parent *Window) *FontChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_font_chooser_dialog_new(c_title, c_parent)
	retGo := FontChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontChooserWidgetNew is a wrapper around the C function gtk_font_chooser_widget_new.
func FontChooserWidgetNew() *FontChooserWidget {
	retC := C.gtk_font_chooser_widget_new()
	retGo := FontChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LockButtonNew is a wrapper around the C function gtk_lock_button_new.
func LockButtonNew(permission *gio.Permission) *LockButton {
	c_permission := (*C.GPermission)(C.NULL)
	if permission != nil {
		c_permission = (*C.GPermission)(permission.ToC())
	}

	retC := C.gtk_lock_button_new(c_permission)
	retGo := LockButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OverlayNew is a wrapper around the C function gtk_overlay_new.
func OverlayNew() *Overlay {
	retC := C.gtk_overlay_new()
	retGo := OverlayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CssSection is a wrapper around the C record GtkCssSection.
type CssSection struct {
	native *C.GtkCssSection
}

func CssSectionNewFromC(u unsafe.Pointer) *CssSection {
	c := (*C.GtkCssSection)(u)
	if c == nil {
		return nil
	}

	g := &CssSection{native: c}

	return g
}

func (recv *CssSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
