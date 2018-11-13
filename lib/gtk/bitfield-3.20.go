// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

/*
Flags that modify the behavior of gtk_style_context_to_string().
New values may be added to this enumeration.
*/
type StyleContextPrintFlags C.GtkStyleContextPrintFlags

const (
	GTK_STYLE_CONTEXT_PRINT_NONE StyleContextPrintFlags = 0
	/*
	   Print the entire tree of
	       CSS nodes starting at the style context's node
	*/
	GTK_STYLE_CONTEXT_PRINT_RECURSE StyleContextPrintFlags = 1
	/*
	   Show the values of the
	       CSS properties for each node
	*/
	GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE StyleContextPrintFlags = 2
)
