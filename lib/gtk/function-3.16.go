// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import gdk "github.com/pekim/gobbi/lib/gdk"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Cancels an ongoing drag operation on the source side.
//
// If you want to be able to cancel a drag operation in this way,
// you need to keep a pointer to the drag context, either from an
// explicit call to gtk_drag_begin_with_coordinates(), or by
// connecting to #GtkWidget::drag-begin.
//
// If @context does not refer to an ongoing drag operation, this
// function does nothing.
//
// If a drag is cancelled in this way, the @result argument of
// #GtkWidget::drag-failed is set to @GTK_DRAG_RESULT_ERROR.
/*

C function : gtk_drag_cancel
*/
func DragCancel(context *gdk.DragContext) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	C.gtk_drag_cancel(c_context)

	return
}
