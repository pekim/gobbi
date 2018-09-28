// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type StyleContextPrintFlags C.GtkStyleContextPrintFlags

const (
	GTK_STYLE_CONTEXT_PRINT_NONE       StyleContextPrintFlags = 0
	GTK_STYLE_CONTEXT_PRINT_RECURSE    StyleContextPrintFlags = 1
	GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE StyleContextPrintFlags = 2
)
