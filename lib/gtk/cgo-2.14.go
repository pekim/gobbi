// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
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

// HSVNew is a wrapper around the C function gtk_hsv_new.
func HSVNew() *HSV {
	retC := C.gtk_hsv_new()
	retGo := HSVNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconInfoNewForPixbuf is a wrapper around the C function gtk_icon_info_new_for_pixbuf.
func IconInfoNewForPixbuf(iconTheme *IconTheme, pixbuf *gdkpixbuf.Pixbuf) *IconInfo {
	c_icon_theme := (*C.GtkIconTheme)(C.NULL)
	if iconTheme != nil {
		c_icon_theme = (*C.GtkIconTheme)(iconTheme.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_icon_info_new_for_pixbuf(c_icon_theme, c_pixbuf)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ImageNewFromGicon is a wrapper around the C function gtk_image_new_from_gicon.
func ImageNewFromGicon(icon *gio.Icon, size IconSize) *Image {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_gicon(c_icon, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MountOperationNew is a wrapper around the C function gtk_mount_operation_new.
func MountOperationNew(parent *Window) *MountOperation {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_mount_operation_new(c_parent)
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNewFromGicon is a wrapper around the C function gtk_status_icon_new_from_gicon.
func StatusIconNewFromGicon(icon *gio.Icon) *StatusIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.gtk_status_icon_new_from_gicon(c_icon)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : STOCK_PAGE_SETUP

// Blacklisted : STOCK_PRINT_ERROR

// Blacklisted : STOCK_PRINT_PAUSED

// Blacklisted : STOCK_PRINT_REPORT

// Blacklisted : STOCK_PRINT_WARNING

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_init : unsupported parameter argcp : array length param argcp is pointer (int*)

// Unsupported : gtk_test_list_all_types : array return type :

// BorderNew is a wrapper around the C function gtk_border_new.
func BorderNew() *Border {
	retC := C.gtk_border_new()
	retGo := BorderNewFromC(unsafe.Pointer(retC))

	return retGo
}
