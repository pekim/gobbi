// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void cellrenderer_editingStartedHandler(GObject *, GtkCellEditable *, gchar*, gpointer);

	static gulong CellRenderer_signal_connect_editing_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-started", G_CALLBACK(cellrenderer_editingStartedHandler), data);
	}

*/
/*

	void clipboard_ownerChangeHandler(GObject *, GdkEventOwnerChange *, gpointer);

	static gulong Clipboard_signal_connect_owner_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "owner-change", G_CALLBACK(clipboard_ownerChangeHandler), data);
	}

*/
/*

	gboolean entrycompletion_insertPrefixHandler(GObject *, gchar*, gpointer);

	static gulong EntryCompletion_signal_connect_insert_prefix(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-prefix", G_CALLBACK(entrycompletion_insertPrefixHandler), data);
	}

*/
/*

	gboolean range_changeValueHandler(GObject *, GtkScrollType, gdouble, gpointer);

	static gulong Range_signal_connect_change_value(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-value", G_CALLBACK(range_changeValueHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_about_dialog_new

// Blacklisted : gtk_about_dialog_get_artists

// Blacklisted : gtk_about_dialog_get_authors

// Blacklisted : gtk_about_dialog_get_comments

// Blacklisted : gtk_about_dialog_get_copyright

// Blacklisted : gtk_about_dialog_get_documenters

// Blacklisted : gtk_about_dialog_get_license

// Blacklisted : gtk_about_dialog_get_logo

// Blacklisted : gtk_about_dialog_get_logo_icon_name

// Blacklisted : gtk_about_dialog_get_translator_credits

// Blacklisted : gtk_about_dialog_get_version

// Blacklisted : gtk_about_dialog_get_website

// Blacklisted : gtk_about_dialog_get_website_label

// Blacklisted : gtk_about_dialog_set_artists

// Blacklisted : gtk_about_dialog_set_authors

// Blacklisted : gtk_about_dialog_set_comments

// Blacklisted : gtk_about_dialog_set_copyright

// Blacklisted : gtk_about_dialog_set_documenters

// Blacklisted : gtk_about_dialog_set_license

// Blacklisted : gtk_about_dialog_set_logo

// Blacklisted : gtk_about_dialog_set_logo_icon_name

// Blacklisted : gtk_about_dialog_set_translator_credits

// Blacklisted : gtk_about_dialog_set_version

// Blacklisted : gtk_about_dialog_set_website

// Blacklisted : gtk_about_dialog_set_website_label

// Blacklisted : gtk_action_get_accel_path

// Blacklisted : gtk_action_set_sensitive

// Blacklisted : gtk_action_set_visible

// Blacklisted : gtk_action_group_translate_string

// Blacklisted : gtk_button_get_image

// Blacklisted : gtk_button_set_image

type signalCellRendererEditingStartedDetail struct {
	callback  CellRendererSignalEditingStartedCallback
	handlerID C.gulong
}

var signalCellRendererEditingStartedId int
var signalCellRendererEditingStartedMap = make(map[int]signalCellRendererEditingStartedDetail)
var signalCellRendererEditingStartedLock sync.RWMutex

// CellRendererSignalEditingStartedCallback is a callback function for a 'editing-started' signal emitted from a CellRenderer.
type CellRendererSignalEditingStartedCallback func(editable *CellEditable, path string)

/*
ConnectEditingStarted connects the callback to the 'editing-started' signal for the CellRenderer.

The returned value represents the connection, and may be passed to DisconnectEditingStarted to remove it.
*/
func (recv *CellRenderer) ConnectEditingStarted(callback CellRendererSignalEditingStartedCallback) int {
	signalCellRendererEditingStartedLock.Lock()
	defer signalCellRendererEditingStartedLock.Unlock()

	signalCellRendererEditingStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRenderer_signal_connect_editing_started(instance, C.gpointer(uintptr(signalCellRendererEditingStartedId)))

	detail := signalCellRendererEditingStartedDetail{callback, handlerID}
	signalCellRendererEditingStartedMap[signalCellRendererEditingStartedId] = detail

	return signalCellRendererEditingStartedId
}

/*
DisconnectEditingStarted disconnects a callback from the 'editing-started' signal for the CellRenderer.

The connectionID should be a value returned from a call to ConnectEditingStarted.
*/
func (recv *CellRenderer) DisconnectEditingStarted(connectionID int) {
	signalCellRendererEditingStartedLock.Lock()
	defer signalCellRendererEditingStartedLock.Unlock()

	detail, exists := signalCellRendererEditingStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererEditingStartedMap, connectionID)
}

//export cellrenderer_editingStartedHandler
func cellrenderer_editingStartedHandler(_ *C.GObject, c_editable *C.GtkCellEditable, c_path *C.gchar, data C.gpointer) {
	signalCellRendererEditingStartedLock.RLock()
	defer signalCellRendererEditingStartedLock.RUnlock()

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellRendererEditingStartedMap[index].callback
	callback(editable, path)
}

// Blacklisted : gtk_cell_renderer_stop_editing

// Blacklisted : gtk_cell_renderer_combo_new

// Blacklisted : gtk_cell_renderer_progress_new

// Blacklisted : gtk_cell_view_new

// Blacklisted : gtk_cell_view_new_with_context

// Blacklisted : gtk_cell_view_new_with_markup

// Blacklisted : gtk_cell_view_new_with_pixbuf

// Blacklisted : gtk_cell_view_new_with_text

// Blacklisted : gtk_cell_view_get_displayed_row

// Blacklisted : gtk_cell_view_get_size_of_row

// Blacklisted : gtk_cell_view_set_background_color

// Blacklisted : gtk_cell_view_set_displayed_row

// Blacklisted : gtk_cell_view_set_model

type signalClipboardOwnerChangeDetail struct {
	callback  ClipboardSignalOwnerChangeCallback
	handlerID C.gulong
}

var signalClipboardOwnerChangeId int
var signalClipboardOwnerChangeMap = make(map[int]signalClipboardOwnerChangeDetail)
var signalClipboardOwnerChangeLock sync.RWMutex

// ClipboardSignalOwnerChangeCallback is a callback function for a 'owner-change' signal emitted from a Clipboard.
type ClipboardSignalOwnerChangeCallback func(event *gdk.EventOwnerChange)

/*
ConnectOwnerChange connects the callback to the 'owner-change' signal for the Clipboard.

The returned value represents the connection, and may be passed to DisconnectOwnerChange to remove it.
*/
func (recv *Clipboard) ConnectOwnerChange(callback ClipboardSignalOwnerChangeCallback) int {
	signalClipboardOwnerChangeLock.Lock()
	defer signalClipboardOwnerChangeLock.Unlock()

	signalClipboardOwnerChangeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Clipboard_signal_connect_owner_change(instance, C.gpointer(uintptr(signalClipboardOwnerChangeId)))

	detail := signalClipboardOwnerChangeDetail{callback, handlerID}
	signalClipboardOwnerChangeMap[signalClipboardOwnerChangeId] = detail

	return signalClipboardOwnerChangeId
}

/*
DisconnectOwnerChange disconnects a callback from the 'owner-change' signal for the Clipboard.

The connectionID should be a value returned from a call to ConnectOwnerChange.
*/
func (recv *Clipboard) DisconnectOwnerChange(connectionID int) {
	signalClipboardOwnerChangeLock.Lock()
	defer signalClipboardOwnerChangeLock.Unlock()

	detail, exists := signalClipboardOwnerChangeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalClipboardOwnerChangeMap, connectionID)
}

//export clipboard_ownerChangeHandler
func clipboard_ownerChangeHandler(_ *C.GObject, c_event *C.GdkEventOwnerChange, data C.gpointer) {
	signalClipboardOwnerChangeLock.RLock()
	defer signalClipboardOwnerChangeLock.RUnlock()

	event := gdk.EventOwnerChangeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalClipboardOwnerChangeMap[index].callback
	callback(event)
}

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc (GtkClipboardImageReceivedFunc) for param callback

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets :

// Blacklisted : gtk_clipboard_set_image

// Blacklisted : gtk_clipboard_store

// Blacklisted : gtk_clipboard_wait_for_image

// Blacklisted : gtk_clipboard_wait_is_image_available

// Blacklisted : gtk_clipboard_wait_is_target_available

// Blacklisted : gtk_combo_box_get_column_span_column

// Blacklisted : gtk_combo_box_get_focus_on_click

// Blacklisted : gtk_combo_box_get_popup_accessible

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// Blacklisted : gtk_combo_box_get_row_span_column

// Blacklisted : gtk_combo_box_get_wrap_width

// Blacklisted : gtk_combo_box_set_add_tearoffs

// Blacklisted : gtk_combo_box_set_focus_on_click

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Blacklisted : gtk_dialog_set_alternative_button_order_from_array

type signalEntryCompletionInsertPrefixDetail struct {
	callback  EntryCompletionSignalInsertPrefixCallback
	handlerID C.gulong
}

var signalEntryCompletionInsertPrefixId int
var signalEntryCompletionInsertPrefixMap = make(map[int]signalEntryCompletionInsertPrefixDetail)
var signalEntryCompletionInsertPrefixLock sync.RWMutex

// EntryCompletionSignalInsertPrefixCallback is a callback function for a 'insert-prefix' signal emitted from a EntryCompletion.
type EntryCompletionSignalInsertPrefixCallback func(prefix string) bool

/*
ConnectInsertPrefix connects the callback to the 'insert-prefix' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectInsertPrefix to remove it.
*/
func (recv *EntryCompletion) ConnectInsertPrefix(callback EntryCompletionSignalInsertPrefixCallback) int {
	signalEntryCompletionInsertPrefixLock.Lock()
	defer signalEntryCompletionInsertPrefixLock.Unlock()

	signalEntryCompletionInsertPrefixId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_insert_prefix(instance, C.gpointer(uintptr(signalEntryCompletionInsertPrefixId)))

	detail := signalEntryCompletionInsertPrefixDetail{callback, handlerID}
	signalEntryCompletionInsertPrefixMap[signalEntryCompletionInsertPrefixId] = detail

	return signalEntryCompletionInsertPrefixId
}

/*
DisconnectInsertPrefix disconnects a callback from the 'insert-prefix' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectInsertPrefix.
*/
func (recv *EntryCompletion) DisconnectInsertPrefix(connectionID int) {
	signalEntryCompletionInsertPrefixLock.Lock()
	defer signalEntryCompletionInsertPrefixLock.Unlock()

	detail, exists := signalEntryCompletionInsertPrefixMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionInsertPrefixMap, connectionID)
}

//export entrycompletion_insertPrefixHandler
func entrycompletion_insertPrefixHandler(_ *C.GObject, c_prefix *C.gchar, data C.gpointer) C.gboolean {
	signalEntryCompletionInsertPrefixLock.RLock()
	defer signalEntryCompletionInsertPrefixLock.RUnlock()

	prefix := C.GoString(c_prefix)

	index := int(uintptr(data))
	callback := signalEntryCompletionInsertPrefixMap[index].callback
	retGo := callback(prefix)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_entry_completion_get_inline_completion

// Blacklisted : gtk_entry_completion_get_popup_completion

// Blacklisted : gtk_entry_completion_get_text_column

// Blacklisted : gtk_entry_completion_insert_prefix

// Blacklisted : gtk_entry_completion_set_inline_completion

// Blacklisted : gtk_entry_completion_set_popup_completion

// Blacklisted : gtk_file_chooser_button_new

// Blacklisted : gtk_file_chooser_button_new_with_dialog

// GetTitle is a wrapper around the C function gtk_file_chooser_button_get_title.
func (recv *FileChooserButton) GetTitle() string {
	retC := C.gtk_file_chooser_button_get_title((*C.GtkFileChooserButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_file_chooser_button_get_width_chars

// SetTitle is a wrapper around the C function gtk_file_chooser_button_set_title.
func (recv *FileChooserButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_file_chooser_button_set_title((*C.GtkFileChooserButton)(recv.native), c_title)

	return
}

// Blacklisted : gtk_file_chooser_button_set_width_chars

// Blacklisted : gtk_file_filter_add_pixbuf_formats

// Unsupported : gtk_icon_theme_get_icon_sizes : array return type :

// Blacklisted : gtk_icon_view_new

// Blacklisted : gtk_icon_view_new_with_model

// Blacklisted : gtk_icon_view_get_column_spacing

// Blacklisted : gtk_icon_view_get_columns

// Blacklisted : gtk_icon_view_get_item_orientation

// Blacklisted : gtk_icon_view_get_item_width

// Blacklisted : gtk_icon_view_get_margin

// Blacklisted : gtk_icon_view_get_markup_column

// Blacklisted : gtk_icon_view_get_model

// Blacklisted : gtk_icon_view_get_path_at_pos

// Blacklisted : gtk_icon_view_get_pixbuf_column

// Blacklisted : gtk_icon_view_get_row_spacing

// Blacklisted : gtk_icon_view_get_selected_items

// Blacklisted : gtk_icon_view_get_selection_mode

// Blacklisted : gtk_icon_view_get_spacing

// Blacklisted : gtk_icon_view_get_text_column

// Blacklisted : gtk_icon_view_item_activated

// Blacklisted : gtk_icon_view_path_is_selected

// Blacklisted : gtk_icon_view_select_all

// Blacklisted : gtk_icon_view_select_path

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc (GtkIconViewForeachFunc) for param func

// Blacklisted : gtk_icon_view_set_column_spacing

// Blacklisted : gtk_icon_view_set_columns

// Blacklisted : gtk_icon_view_set_item_orientation

// Blacklisted : gtk_icon_view_set_item_width

// Blacklisted : gtk_icon_view_set_margin

// Blacklisted : gtk_icon_view_set_markup_column

// Blacklisted : gtk_icon_view_set_model

// Blacklisted : gtk_icon_view_set_pixbuf_column

// Blacklisted : gtk_icon_view_set_row_spacing

// Blacklisted : gtk_icon_view_set_selection_mode

// Blacklisted : gtk_icon_view_set_spacing

// Blacklisted : gtk_icon_view_set_text_column

// Blacklisted : gtk_icon_view_unselect_all

// Blacklisted : gtk_icon_view_unselect_path

// Blacklisted : gtk_image_new_from_icon_name

// Blacklisted : gtk_image_get_icon_name

// Blacklisted : gtk_image_get_pixel_size

// Blacklisted : gtk_image_set_from_icon_name

// Blacklisted : gtk_image_set_pixel_size

// Blacklisted : gtk_label_get_angle

// Blacklisted : gtk_label_get_ellipsize

// Blacklisted : gtk_label_get_max_width_chars

// Blacklisted : gtk_label_get_single_line_mode

// Blacklisted : gtk_label_get_width_chars

// Blacklisted : gtk_label_set_angle

// Blacklisted : gtk_label_set_ellipsize

// Blacklisted : gtk_label_set_max_width_chars

// Blacklisted : gtk_label_set_single_line_mode

// Blacklisted : gtk_label_set_width_chars

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter values :

// Blacklisted : gtk_menu_get_for_attach_widget

// Blacklisted : gtk_menu_tool_button_new

// Blacklisted : gtk_menu_tool_button_new_from_stock

// Blacklisted : gtk_menu_tool_button_get_menu

// Blacklisted : gtk_menu_tool_button_set_menu

// Blacklisted : gtk_message_dialog_format_secondary_markup

// Blacklisted : gtk_message_dialog_format_secondary_text

// Blacklisted : gtk_progress_bar_get_ellipsize

// Blacklisted : gtk_progress_bar_set_ellipsize

type signalRangeChangeValueDetail struct {
	callback  RangeSignalChangeValueCallback
	handlerID C.gulong
}

var signalRangeChangeValueId int
var signalRangeChangeValueMap = make(map[int]signalRangeChangeValueDetail)
var signalRangeChangeValueLock sync.RWMutex

// RangeSignalChangeValueCallback is a callback function for a 'change-value' signal emitted from a Range.
type RangeSignalChangeValueCallback func(scroll ScrollType, value float64) bool

/*
ConnectChangeValue connects the callback to the 'change-value' signal for the Range.

The returned value represents the connection, and may be passed to DisconnectChangeValue to remove it.
*/
func (recv *Range) ConnectChangeValue(callback RangeSignalChangeValueCallback) int {
	signalRangeChangeValueLock.Lock()
	defer signalRangeChangeValueLock.Unlock()

	signalRangeChangeValueId++
	instance := C.gpointer(recv.native)
	handlerID := C.Range_signal_connect_change_value(instance, C.gpointer(uintptr(signalRangeChangeValueId)))

	detail := signalRangeChangeValueDetail{callback, handlerID}
	signalRangeChangeValueMap[signalRangeChangeValueId] = detail

	return signalRangeChangeValueId
}

/*
DisconnectChangeValue disconnects a callback from the 'change-value' signal for the Range.

The connectionID should be a value returned from a call to ConnectChangeValue.
*/
func (recv *Range) DisconnectChangeValue(connectionID int) {
	signalRangeChangeValueLock.Lock()
	defer signalRangeChangeValueLock.Unlock()

	detail, exists := signalRangeChangeValueMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRangeChangeValueMap, connectionID)
}

//export range_changeValueHandler
func range_changeValueHandler(_ *C.GObject, c_scroll C.GtkScrollType, c_value C.gdouble, data C.gpointer) C.gboolean {
	signalRangeChangeValueLock.RLock()
	defer signalRangeChangeValueLock.RUnlock()

	scroll := ScrollType(c_scroll)

	value := float64(c_value)

	index := int(uintptr(data))
	callback := signalRangeChangeValueMap[index].callback
	retGo := callback(scroll, value)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_text_buffer_backspace

// Blacklisted : gtk_text_view_get_iter_at_position

// Blacklisted : gtk_tool_item_rebuild_menu

// Blacklisted : gtk_tree_view_get_fixed_height_mode

// Blacklisted : gtk_tree_view_get_hover_expand

// Blacklisted : gtk_tree_view_get_hover_selection

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

// Blacklisted : gtk_tree_view_set_fixed_height_mode

// Blacklisted : gtk_tree_view_set_hover_expand

// Blacklisted : gtk_tree_view_set_hover_selection

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc (GtkTreeViewRowSeparatorFunc) for param func

// Blacklisted : gtk_drag_dest_add_image_targets

// Blacklisted : gtk_drag_dest_add_text_targets

// Blacklisted : gtk_drag_dest_add_uri_targets

// Blacklisted : gtk_drag_source_add_image_targets

// Blacklisted : gtk_drag_source_add_text_targets

// Blacklisted : gtk_drag_source_add_uri_targets

// Blacklisted : gtk_window_set_default_icon_name

// Blacklisted : gtk_window_get_focus_on_map

// Blacklisted : gtk_window_get_icon_name

// Blacklisted : gtk_window_set_focus_on_map

// Blacklisted : gtk_window_set_icon_name
