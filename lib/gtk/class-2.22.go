// This is a generated file - DO NOT EDIT
// +build gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Gets the #GtkWidget corresponding to the #GtkAccessible.
// The returned widget does not have a reference added, so
// you do not need to unref it.
/*

C function : gtk_accessible_get_widget
*/
func (recv *Accessible) GetWidget() *Widget {
	retC := C.gtk_accessible_get_widget((*C.GtkAccessible)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the #GtkWidget corresponding to the #GtkAccessible.
//
// @accessible will not hold a reference to @widget.
// It is the caller’s responsibility to ensure that when @widget
// is destroyed, the widget is unset by calling this function
// again with @widget set to %NULL.
/*

C function : gtk_accessible_set_widget
*/
func (recv *Accessible) SetWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_accessible_set_widget((*C.GtkAccessible)(recv.native), c_widget)

	return
}

// Erases the visited page history so the back button is not
// shown on the current page, and removes the cancel button
// from subsequent pages.
//
// Use this when the information provided up to the current
// page is hereafter deemed permanent and cannot be modified
// or undone. For example, showing a progress page to track
// a long-running, unreversible operation after the user has
// clicked apply on a confirmation page.
/*

C function : gtk_assistant_commit
*/
func (recv *Assistant) Commit() {
	C.gtk_assistant_commit((*C.GtkAssistant)(recv.native))

	return
}

// Returns the button’s event window if it is realized, %NULL otherwise.
// This function should be rarely needed.
/*

C function : gtk_button_get_event_window
*/
func (recv *Button) GetEventWindow() *gdk.Window {
	retC := C.gtk_button_get_event_window((*C.GtkButton)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Allow the #GtkEntry input method to internally handle key press
// and release events. If this function returns %TRUE, then no further
// processing should be done for this key event. See
// gtk_im_context_filter_keypress().
//
// Note that you are expected to call this function from your handler
// when overriding key event handling. This is needed in the case when
// you need to insert your own key handling between the input method
// and the default key event handling of the #GtkEntry.
// See gtk_text_view_reset_im_context() for an example of use.
/*

C function : gtk_entry_im_context_filter_keypress
*/
func (recv *Entry) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_entry_im_context_filter_keypress((*C.GtkEntry)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Reset the input method context of the entry if needed.
//
// This can be necessary in the case where modifying the buffer
// would confuse on-going input method behavior.
/*

C function : gtk_entry_reset_im_context
*/
func (recv *Entry) ResetImContext() {
	C.gtk_entry_reset_im_context((*C.GtkEntry)(recv.native))

	return
}

// Returns whether the label widget will fill all available
// horizontal space allocated to @expander.
/*

C function : gtk_expander_get_label_fill
*/
func (recv *Expander) GetLabelFill() bool {
	retC := C.gtk_expander_get_label_fill((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the label widget should fill all available
// horizontal space allocated to @expander.
/*

C function : gtk_expander_set_label_fill
*/
func (recv *Expander) SetLabelFill(labelFill bool) {
	c_label_fill :=
		boolToGboolean(labelFill)

	C.gtk_expander_set_label_fill((*C.GtkExpander)(recv.native), c_label_fill)

	return
}

// Retrieves the #GtkFontSelection widget embedded in the dialog.
/*

C function : gtk_font_selection_dialog_get_font_selection
*/
func (recv *FontSelectionDialog) GetFontSelection() *Widget {
	retC := C.gtk_font_selection_dialog_get_font_selection((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the column in which the item @path is currently
// displayed. Column numbers start at 0.
/*

C function : gtk_icon_view_get_item_column
*/
func (recv *IconView) GetItemColumn(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_get_item_column((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// Gets the row in which the item @path is currently
// displayed. Row numbers start at 0.
/*

C function : gtk_icon_view_get_item_row
*/
func (recv *IconView) GetItemRow(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_get_item_row((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// Returns the message area of the dialog. This is the box where the
// dialog’s primary and secondary labels are packed. You can add your
// own extra content to that box and it will appear below those labels.
// See gtk_dialog_get_content_area() for the corresponding
// function in the parent #GtkDialog.
/*

C function : gtk_message_dialog_get_message_area
*/
func (recv *MessageDialog) GetMessageArea() *Widget {
	retC := C.gtk_message_dialog_get_message_area((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the horizontal width of a tab border.
/*

C function : gtk_notebook_get_tab_hborder
*/
func (recv *Notebook) GetTabHborder() uint16 {
	retC := C.gtk_notebook_get_tab_hborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Returns the vertical width of a tab border.
/*

C function : gtk_notebook_get_tab_vborder
*/
func (recv *Notebook) GetTabVborder() uint16 {
	retC := C.gtk_notebook_get_tab_vborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Forces the removal of all messages from a statusbar's
// stack with the exact @context_id.
/*

C function : gtk_statusbar_remove_all
*/
func (recv *Statusbar) RemoveAll(contextId uint32) {
	c_context_id := (C.guint)(contextId)

	C.gtk_statusbar_remove_all((*C.GtkStatusbar)(recv.native), c_context_id)

	return
}

// Gets the number of rows and columns in the table.
/*

C function : gtk_table_get_size
*/
func (recv *Table) GetSize() (uint32, uint32) {
	var c_rows C.guint

	var c_columns C.guint

	C.gtk_table_get_size((*C.GtkTable)(recv.native), &c_rows, &c_columns)

	rows := (uint32)(c_rows)

	columns := (uint32)(c_columns)

	return rows, columns
}

// Gets the horizontal-scrolling #GtkAdjustment.
/*

C function : gtk_text_view_get_hadjustment
*/
func (recv *TextView) GetHadjustment() *Adjustment {
	retC := C.gtk_text_view_get_hadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the vertical-scrolling #GtkAdjustment.
/*

C function : gtk_text_view_get_vadjustment
*/
func (recv *TextView) GetVadjustment() *Adjustment {
	retC := C.gtk_text_view_get_vadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Allow the #GtkTextView input method to internally handle key press
// and release events. If this function returns %TRUE, then no further
// processing should be done for this key event. See
// gtk_im_context_filter_keypress().
//
// Note that you are expected to call this function from your handler
// when overriding key event handling. This is needed in the case when
// you need to insert your own key handling between the input method
// and the default key event handling of the #GtkTextView.
//
// |[<!-- language="C" -->
// static gboolean
// gtk_foo_bar_key_press_event (GtkWidget   *widget,
// GdkEventKey *event)
// {
// guint keyval;
//
// gdk_event_get_keyval ((GdkEvent*)event, &keyval);
//
// if (keyval == GDK_KEY_Return || keyval == GDK_KEY_KP_Enter)
// {
// if (gtk_text_view_im_context_filter_keypress (GTK_TEXT_VIEW (widget), event))
// return TRUE;
// }
//
// Do some stuff
//
// return GTK_WIDGET_CLASS (gtk_foo_bar_parent_class)->key_press_event (widget, event);
// }
// ]|
/*

C function : gtk_text_view_im_context_filter_keypress
*/
func (recv *TextView) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_text_view_im_context_filter_keypress((*C.GtkTextView)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Reset the input method context of the text view if needed.
//
// This can be necessary in the case where modifying the buffer
// would confuse on-going input method behavior.
/*

C function : gtk_text_view_reset_im_context
*/
func (recv *TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context((*C.GtkTextView)(recv.native))

	return
}

// Gets the view window of the #GtkViewport.
/*

C function : gtk_viewport_get_view_window
*/
func (recv *Viewport) GetViewWindow() *gdk.Window {
	retC := C.gtk_viewport_get_view_window((*C.GtkViewport)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the current grab widget of the given group,
// see gtk_grab_add().
/*

C function : gtk_window_group_get_current_grab
*/
func (recv *WindowGroup) GetCurrentGrab() *Widget {
	retC := C.gtk_window_group_get_current_grab((*C.GtkWindowGroup)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}
