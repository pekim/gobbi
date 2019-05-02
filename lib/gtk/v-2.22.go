// This is a generated file - DO NOT EDIT
// +build gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// GetWidget is a wrapper around the C function gtk_accessible_get_widget.
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

// SetWidget is a wrapper around the C function gtk_accessible_set_widget.
func (recv *Accessible) SetWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_accessible_set_widget((*C.GtkAccessible)(recv.native), c_widget)

	return
}

// Commit is a wrapper around the C function gtk_assistant_commit.
func (recv *Assistant) Commit() {
	C.gtk_assistant_commit((*C.GtkAssistant)(recv.native))

	return
}

// GetEventWindow is a wrapper around the C function gtk_button_get_event_window.
func (recv *Button) GetEventWindow() *gdk.Window {
	retC := C.gtk_button_get_event_window((*C.GtkButton)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImContextFilterKeypress is a wrapper around the C function gtk_entry_im_context_filter_keypress.
func (recv *Entry) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_entry_im_context_filter_keypress((*C.GtkEntry)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// ResetImContext is a wrapper around the C function gtk_entry_reset_im_context.
func (recv *Entry) ResetImContext() {
	C.gtk_entry_reset_im_context((*C.GtkEntry)(recv.native))

	return
}

// GetLabelFill is a wrapper around the C function gtk_expander_get_label_fill.
func (recv *Expander) GetLabelFill() bool {
	retC := C.gtk_expander_get_label_fill((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLabelFill is a wrapper around the C function gtk_expander_set_label_fill.
func (recv *Expander) SetLabelFill(labelFill bool) {
	c_label_fill :=
		boolToGboolean(labelFill)

	C.gtk_expander_set_label_fill((*C.GtkExpander)(recv.native), c_label_fill)

	return
}

// GetFontSelection is a wrapper around the C function gtk_font_selection_dialog_get_font_selection.
func (recv *FontSelectionDialog) GetFontSelection() *Widget {
	retC := C.gtk_font_selection_dialog_get_font_selection((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetItemColumn is a wrapper around the C function gtk_icon_view_get_item_column.
func (recv *IconView) GetItemColumn(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_get_item_column((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// GetItemRow is a wrapper around the C function gtk_icon_view_get_item_row.
func (recv *IconView) GetItemRow(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_get_item_row((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// GetMessageArea is a wrapper around the C function gtk_message_dialog_get_message_area.
func (recv *MessageDialog) GetMessageArea() *Widget {
	retC := C.gtk_message_dialog_get_message_area((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTabHborder is a wrapper around the C function gtk_notebook_get_tab_hborder.
func (recv *Notebook) GetTabHborder() uint16 {
	retC := C.gtk_notebook_get_tab_hborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetTabVborder is a wrapper around the C function gtk_notebook_get_tab_vborder.
func (recv *Notebook) GetTabVborder() uint16 {
	retC := C.gtk_notebook_get_tab_vborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// RemoveAll is a wrapper around the C function gtk_statusbar_remove_all.
func (recv *Statusbar) RemoveAll(contextId uint32) {
	c_context_id := (C.guint)(contextId)

	C.gtk_statusbar_remove_all((*C.GtkStatusbar)(recv.native), c_context_id)

	return
}

// GetSize is a wrapper around the C function gtk_table_get_size.
func (recv *Table) GetSize() (uint32, uint32) {
	var c_rows C.guint

	var c_columns C.guint

	C.gtk_table_get_size((*C.GtkTable)(recv.native), &c_rows, &c_columns)

	rows := (uint32)(c_rows)

	columns := (uint32)(c_columns)

	return rows, columns
}

// GetHadjustment is a wrapper around the C function gtk_text_view_get_hadjustment.
func (recv *TextView) GetHadjustment() *Adjustment {
	retC := C.gtk_text_view_get_hadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_text_view_get_vadjustment.
func (recv *TextView) GetVadjustment() *Adjustment {
	retC := C.gtk_text_view_get_vadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImContextFilterKeypress is a wrapper around the C function gtk_text_view_im_context_filter_keypress.
func (recv *TextView) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_text_view_im_context_filter_keypress((*C.GtkTextView)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// ResetImContext is a wrapper around the C function gtk_text_view_reset_im_context.
func (recv *TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context((*C.GtkTextView)(recv.native))

	return
}

// GetViewWindow is a wrapper around the C function gtk_viewport_get_view_window.
func (recv *Viewport) GetViewWindow() *gdk.Window {
	retC := C.gtk_viewport_get_view_window((*C.GtkViewport)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentGrab is a wrapper around the C function gtk_window_group_get_current_grab.
func (recv *WindowGroup) GetCurrentGrab() *Widget {
	retC := C.gtk_window_group_get_current_grab((*C.GtkWindowGroup)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGicon is a wrapper around the C function gtk_recent_info_get_gicon.
func (recv *RecentInfo) GetGicon() *gio.Icon {
	retC := C.gtk_recent_info_get_gicon((*C.GtkRecentInfo)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
