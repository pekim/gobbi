// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	void combobox_moveActiveHandler(GObject *, GtkScrollType, gpointer);

	static gulong ComboBox_signal_connect_move_active(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-active", G_CALLBACK(combobox_moveActiveHandler), data);
	}

*/
/*

	gboolean combobox_popdownHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(combobox_popdownHandler), data);
	}

*/
/*

	void combobox_popupHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(combobox_popupHandler), data);
	}

*/
/*

	gboolean entrycompletion_cursorOnMatchHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gpointer);

	static gulong EntryCompletion_signal_connect_cursor_on_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cursor-on-match", G_CALLBACK(entrycompletion_cursorOnMatchHandler), data);
	}

*/
/*

	void filechooserbutton_fileSetHandler(GObject *, gpointer);

	static gulong FileChooserButton_signal_connect_file_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-set", G_CALLBACK(filechooserbutton_fileSetHandler), data);
	}

*/
/*

	gboolean menushell_moveSelectedHandler(GObject *, gint, gpointer);

	static gulong MenuShell_signal_connect_move_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-selected", G_CALLBACK(menushell_moveSelectedHandler), data);
	}

*/
/*

	GtkNotebook * notebook_createWindowHandler(GObject *, GtkWidget *, gint, gint, gpointer);

	static gulong Notebook_signal_connect_create_window(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-window", G_CALLBACK(notebook_createWindowHandler), data);
	}

*/
/*

	void scalebutton_popdownHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(scalebutton_popdownHandler), data);
	}

*/
/*

	void scalebutton_popupHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(scalebutton_popupHandler), data);
	}

*/
/*

	void scalebutton_valueChangedHandler(GObject *, gdouble, gpointer);

	static gulong ScaleButton_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", G_CALLBACK(scalebutton_valueChangedHandler), data);
	}

*/
/*

	gboolean widget_dragFailedHandler(GObject *, GdkDragContext *, GtkDragResult, gpointer);

	static gulong Widget_signal_connect_drag_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-failed", G_CALLBACK(widget_dragFailedHandler), data);
	}

*/
/*

	gboolean widget_keynavFailedHandler(GObject *, GtkDirectionType, gpointer);

	static gulong Widget_signal_connect_keynav_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keynav-failed", G_CALLBACK(widget_keynavFailedHandler), data);
	}

*/
/*

	gboolean widget_queryTooltipHandler(GObject *, gint, gint, gboolean, GtkTooltip *, gpointer);

	static gulong Widget_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(widget_queryTooltipHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_about_dialog_get_program_name

// Blacklisted : gtk_about_dialog_set_program_name

// Blacklisted : gtk_action_create_menu

// Blacklisted : gtk_builder_new

// Blacklisted : gtk_builder_add_from_file

// Blacklisted : gtk_builder_add_from_string

// Unsupported : gtk_builder_connect_signals : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param func

// Blacklisted : gtk_builder_get_object

// Blacklisted : gtk_builder_get_objects

// Blacklisted : gtk_builder_get_translation_domain

// Blacklisted : gtk_builder_get_type_from_name

// Blacklisted : gtk_builder_set_translation_domain

// Blacklisted : gtk_builder_value_from_string

// Blacklisted : gtk_builder_value_from_string_type

type signalComboBoxMoveActiveDetail struct {
	callback  ComboBoxSignalMoveActiveCallback
	handlerID C.gulong
}

var signalComboBoxMoveActiveId int
var signalComboBoxMoveActiveMap = make(map[int]signalComboBoxMoveActiveDetail)
var signalComboBoxMoveActiveLock sync.RWMutex

// ComboBoxSignalMoveActiveCallback is a callback function for a 'move-active' signal emitted from a ComboBox.
type ComboBoxSignalMoveActiveCallback func(scrollType ScrollType)

/*
ConnectMoveActive connects the callback to the 'move-active' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectMoveActive to remove it.
*/
func (recv *ComboBox) ConnectMoveActive(callback ComboBoxSignalMoveActiveCallback) int {
	signalComboBoxMoveActiveLock.Lock()
	defer signalComboBoxMoveActiveLock.Unlock()

	signalComboBoxMoveActiveId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_move_active(instance, C.gpointer(uintptr(signalComboBoxMoveActiveId)))

	detail := signalComboBoxMoveActiveDetail{callback, handlerID}
	signalComboBoxMoveActiveMap[signalComboBoxMoveActiveId] = detail

	return signalComboBoxMoveActiveId
}

/*
DisconnectMoveActive disconnects a callback from the 'move-active' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectMoveActive.
*/
func (recv *ComboBox) DisconnectMoveActive(connectionID int) {
	signalComboBoxMoveActiveLock.Lock()
	defer signalComboBoxMoveActiveLock.Unlock()

	detail, exists := signalComboBoxMoveActiveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxMoveActiveMap, connectionID)
}

//export combobox_moveActiveHandler
func combobox_moveActiveHandler(_ *C.GObject, c_scroll_type C.GtkScrollType, data C.gpointer) {
	signalComboBoxMoveActiveLock.RLock()
	defer signalComboBoxMoveActiveLock.RUnlock()

	scrollType := ScrollType(c_scroll_type)

	index := int(uintptr(data))
	callback := signalComboBoxMoveActiveMap[index].callback
	callback(scrollType)
}

type signalComboBoxPopdownDetail struct {
	callback  ComboBoxSignalPopdownCallback
	handlerID C.gulong
}

var signalComboBoxPopdownId int
var signalComboBoxPopdownMap = make(map[int]signalComboBoxPopdownDetail)
var signalComboBoxPopdownLock sync.RWMutex

// ComboBoxSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ComboBox.
type ComboBoxSignalPopdownCallback func() bool

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ComboBox) ConnectPopdown(callback ComboBoxSignalPopdownCallback) int {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	signalComboBoxPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popdown(instance, C.gpointer(uintptr(signalComboBoxPopdownId)))

	detail := signalComboBoxPopdownDetail{callback, handlerID}
	signalComboBoxPopdownMap[signalComboBoxPopdownId] = detail

	return signalComboBoxPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ComboBox) DisconnectPopdown(connectionID int) {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	detail, exists := signalComboBoxPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopdownMap, connectionID)
}

//export combobox_popdownHandler
func combobox_popdownHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	signalComboBoxPopdownLock.RLock()
	defer signalComboBoxPopdownLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxPopdownMap[index].callback
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalComboBoxPopupDetail struct {
	callback  ComboBoxSignalPopupCallback
	handlerID C.gulong
}

var signalComboBoxPopupId int
var signalComboBoxPopupMap = make(map[int]signalComboBoxPopupDetail)
var signalComboBoxPopupLock sync.RWMutex

// ComboBoxSignalPopupCallback is a callback function for a 'popup' signal emitted from a ComboBox.
type ComboBoxSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ComboBox) ConnectPopup(callback ComboBoxSignalPopupCallback) int {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	signalComboBoxPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popup(instance, C.gpointer(uintptr(signalComboBoxPopupId)))

	detail := signalComboBoxPopupDetail{callback, handlerID}
	signalComboBoxPopupMap[signalComboBoxPopupId] = detail

	return signalComboBoxPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ComboBox) DisconnectPopup(connectionID int) {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	detail, exists := signalComboBoxPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopupMap, connectionID)
}

//export combobox_popupHandler
func combobox_popupHandler(_ *C.GObject, data C.gpointer) {
	signalComboBoxPopupLock.RLock()
	defer signalComboBoxPopupLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxPopupMap[index].callback
	callback()
}

// Blacklisted : gtk_entry_get_cursor_hadjustment

// Blacklisted : gtk_entry_set_cursor_hadjustment

type signalEntryCompletionCursorOnMatchDetail struct {
	callback  EntryCompletionSignalCursorOnMatchCallback
	handlerID C.gulong
}

var signalEntryCompletionCursorOnMatchId int
var signalEntryCompletionCursorOnMatchMap = make(map[int]signalEntryCompletionCursorOnMatchDetail)
var signalEntryCompletionCursorOnMatchLock sync.RWMutex

// EntryCompletionSignalCursorOnMatchCallback is a callback function for a 'cursor-on-match' signal emitted from a EntryCompletion.
type EntryCompletionSignalCursorOnMatchCallback func(model *TreeModel, iter *TreeIter) bool

/*
ConnectCursorOnMatch connects the callback to the 'cursor-on-match' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectCursorOnMatch to remove it.
*/
func (recv *EntryCompletion) ConnectCursorOnMatch(callback EntryCompletionSignalCursorOnMatchCallback) int {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	signalEntryCompletionCursorOnMatchId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_cursor_on_match(instance, C.gpointer(uintptr(signalEntryCompletionCursorOnMatchId)))

	detail := signalEntryCompletionCursorOnMatchDetail{callback, handlerID}
	signalEntryCompletionCursorOnMatchMap[signalEntryCompletionCursorOnMatchId] = detail

	return signalEntryCompletionCursorOnMatchId
}

/*
DisconnectCursorOnMatch disconnects a callback from the 'cursor-on-match' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectCursorOnMatch.
*/
func (recv *EntryCompletion) DisconnectCursorOnMatch(connectionID int) {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	detail, exists := signalEntryCompletionCursorOnMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionCursorOnMatchMap, connectionID)
}

//export entrycompletion_cursorOnMatchHandler
func entrycompletion_cursorOnMatchHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	signalEntryCompletionCursorOnMatchLock.RLock()
	defer signalEntryCompletionCursorOnMatchLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionCursorOnMatchMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_entry_completion_get_completion_prefix

// Blacklisted : gtk_entry_completion_get_inline_selection

// Blacklisted : gtk_entry_completion_set_inline_selection

type signalFileChooserButtonFileSetDetail struct {
	callback  FileChooserButtonSignalFileSetCallback
	handlerID C.gulong
}

var signalFileChooserButtonFileSetId int
var signalFileChooserButtonFileSetMap = make(map[int]signalFileChooserButtonFileSetDetail)
var signalFileChooserButtonFileSetLock sync.RWMutex

// FileChooserButtonSignalFileSetCallback is a callback function for a 'file-set' signal emitted from a FileChooserButton.
type FileChooserButtonSignalFileSetCallback func()

/*
ConnectFileSet connects the callback to the 'file-set' signal for the FileChooserButton.

The returned value represents the connection, and may be passed to DisconnectFileSet to remove it.
*/
func (recv *FileChooserButton) ConnectFileSet(callback FileChooserButtonSignalFileSetCallback) int {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	signalFileChooserButtonFileSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooserButton_signal_connect_file_set(instance, C.gpointer(uintptr(signalFileChooserButtonFileSetId)))

	detail := signalFileChooserButtonFileSetDetail{callback, handlerID}
	signalFileChooserButtonFileSetMap[signalFileChooserButtonFileSetId] = detail

	return signalFileChooserButtonFileSetId
}

/*
DisconnectFileSet disconnects a callback from the 'file-set' signal for the FileChooserButton.

The connectionID should be a value returned from a call to ConnectFileSet.
*/
func (recv *FileChooserButton) DisconnectFileSet(connectionID int) {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	detail, exists := signalFileChooserButtonFileSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserButtonFileSetMap, connectionID)
}

//export filechooserbutton_fileSetHandler
func filechooserbutton_fileSetHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserButtonFileSetLock.RLock()
	defer signalFileChooserButtonFileSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserButtonFileSetMap[index].callback
	callback()
}

// Blacklisted : gtk_icon_theme_choose_icon

// Blacklisted : gtk_icon_theme_list_contexts

// Blacklisted : gtk_icon_view_convert_widget_to_bin_window_coords

// Blacklisted : gtk_icon_view_get_tooltip_column

// Blacklisted : gtk_icon_view_get_tooltip_context

// Blacklisted : gtk_icon_view_set_tooltip_cell

// Blacklisted : gtk_icon_view_set_tooltip_column

// Blacklisted : gtk_icon_view_set_tooltip_item

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter values :

type signalMenuShellMoveSelectedDetail struct {
	callback  MenuShellSignalMoveSelectedCallback
	handlerID C.gulong
}

var signalMenuShellMoveSelectedId int
var signalMenuShellMoveSelectedMap = make(map[int]signalMenuShellMoveSelectedDetail)
var signalMenuShellMoveSelectedLock sync.RWMutex

// MenuShellSignalMoveSelectedCallback is a callback function for a 'move-selected' signal emitted from a MenuShell.
type MenuShellSignalMoveSelectedCallback func(distance int32) bool

/*
ConnectMoveSelected connects the callback to the 'move-selected' signal for the MenuShell.

The returned value represents the connection, and may be passed to DisconnectMoveSelected to remove it.
*/
func (recv *MenuShell) ConnectMoveSelected(callback MenuShellSignalMoveSelectedCallback) int {
	signalMenuShellMoveSelectedLock.Lock()
	defer signalMenuShellMoveSelectedLock.Unlock()

	signalMenuShellMoveSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.MenuShell_signal_connect_move_selected(instance, C.gpointer(uintptr(signalMenuShellMoveSelectedId)))

	detail := signalMenuShellMoveSelectedDetail{callback, handlerID}
	signalMenuShellMoveSelectedMap[signalMenuShellMoveSelectedId] = detail

	return signalMenuShellMoveSelectedId
}

/*
DisconnectMoveSelected disconnects a callback from the 'move-selected' signal for the MenuShell.

The connectionID should be a value returned from a call to ConnectMoveSelected.
*/
func (recv *MenuShell) DisconnectMoveSelected(connectionID int) {
	signalMenuShellMoveSelectedLock.Lock()
	defer signalMenuShellMoveSelectedLock.Unlock()

	detail, exists := signalMenuShellMoveSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuShellMoveSelectedMap, connectionID)
}

//export menushell_moveSelectedHandler
func menushell_moveSelectedHandler(_ *C.GObject, c_distance C.gint, data C.gpointer) C.gboolean {
	signalMenuShellMoveSelectedLock.RLock()
	defer signalMenuShellMoveSelectedLock.RUnlock()

	distance := int32(c_distance)

	index := int(uintptr(data))
	callback := signalMenuShellMoveSelectedMap[index].callback
	retGo := callback(distance)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_menu_tool_button_set_arrow_tooltip_markup

// Blacklisted : gtk_menu_tool_button_set_arrow_tooltip_text

type signalNotebookCreateWindowDetail struct {
	callback  NotebookSignalCreateWindowCallback
	handlerID C.gulong
}

var signalNotebookCreateWindowId int
var signalNotebookCreateWindowMap = make(map[int]signalNotebookCreateWindowDetail)
var signalNotebookCreateWindowLock sync.RWMutex

// NotebookSignalCreateWindowCallback is a callback function for a 'create-window' signal emitted from a Notebook.
type NotebookSignalCreateWindowCallback func(page *Widget, x int32, y int32) Notebook

/*
ConnectCreateWindow connects the callback to the 'create-window' signal for the Notebook.

The returned value represents the connection, and may be passed to DisconnectCreateWindow to remove it.
*/
func (recv *Notebook) ConnectCreateWindow(callback NotebookSignalCreateWindowCallback) int {
	signalNotebookCreateWindowLock.Lock()
	defer signalNotebookCreateWindowLock.Unlock()

	signalNotebookCreateWindowId++
	instance := C.gpointer(recv.native)
	handlerID := C.Notebook_signal_connect_create_window(instance, C.gpointer(uintptr(signalNotebookCreateWindowId)))

	detail := signalNotebookCreateWindowDetail{callback, handlerID}
	signalNotebookCreateWindowMap[signalNotebookCreateWindowId] = detail

	return signalNotebookCreateWindowId
}

/*
DisconnectCreateWindow disconnects a callback from the 'create-window' signal for the Notebook.

The connectionID should be a value returned from a call to ConnectCreateWindow.
*/
func (recv *Notebook) DisconnectCreateWindow(connectionID int) {
	signalNotebookCreateWindowLock.Lock()
	defer signalNotebookCreateWindowLock.Unlock()

	detail, exists := signalNotebookCreateWindowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNotebookCreateWindowMap, connectionID)
}

//export notebook_createWindowHandler
func notebook_createWindowHandler(_ *C.GObject, c_page *C.GtkWidget, c_x C.gint, c_y C.gint, data C.gpointer) *C.GtkNotebook {
	signalNotebookCreateWindowLock.RLock()
	defer signalNotebookCreateWindowLock.RUnlock()

	page := WidgetNewFromC(unsafe.Pointer(c_page))

	x := int32(c_x)

	y := int32(c_y)

	index := int(uintptr(data))
	callback := signalNotebookCreateWindowMap[index].callback
	retGo := callback(page, x, y)
	retC :=
		(*C.GtkNotebook)(retGo.ToC())
	return retC
}

// Blacklisted : gtk_page_setup_new_from_file

// Blacklisted : gtk_page_setup_new_from_key_file

// Blacklisted : gtk_page_setup_to_file

// Blacklisted : gtk_page_setup_to_key_file

// Blacklisted : gtk_print_settings_new_from_file

// Blacklisted : gtk_print_settings_new_from_key_file

// Blacklisted : gtk_print_settings_to_file

// Blacklisted : gtk_print_settings_to_key_file

// Blacklisted : gtk_range_get_fill_level

// Blacklisted : gtk_range_get_restrict_to_fill_level

// Blacklisted : gtk_range_get_show_fill_level

// Blacklisted : gtk_range_set_fill_level

// Blacklisted : gtk_range_set_restrict_to_fill_level

// Blacklisted : gtk_range_set_show_fill_level

// Blacklisted : gtk_recent_action_new

// Blacklisted : gtk_recent_action_new_for_manager

// Blacklisted : gtk_recent_action_get_show_numbers

// Blacklisted : gtk_recent_action_set_show_numbers

type signalScaleButtonPopdownDetail struct {
	callback  ScaleButtonSignalPopdownCallback
	handlerID C.gulong
}

var signalScaleButtonPopdownId int
var signalScaleButtonPopdownMap = make(map[int]signalScaleButtonPopdownDetail)
var signalScaleButtonPopdownLock sync.RWMutex

// ScaleButtonSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ScaleButton.
type ScaleButtonSignalPopdownCallback func()

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ScaleButton) ConnectPopdown(callback ScaleButtonSignalPopdownCallback) int {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	signalScaleButtonPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popdown(instance, C.gpointer(uintptr(signalScaleButtonPopdownId)))

	detail := signalScaleButtonPopdownDetail{callback, handlerID}
	signalScaleButtonPopdownMap[signalScaleButtonPopdownId] = detail

	return signalScaleButtonPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ScaleButton) DisconnectPopdown(connectionID int) {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	detail, exists := signalScaleButtonPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopdownMap, connectionID)
}

//export scalebutton_popdownHandler
func scalebutton_popdownHandler(_ *C.GObject, data C.gpointer) {
	signalScaleButtonPopdownLock.RLock()
	defer signalScaleButtonPopdownLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScaleButtonPopdownMap[index].callback
	callback()
}

type signalScaleButtonPopupDetail struct {
	callback  ScaleButtonSignalPopupCallback
	handlerID C.gulong
}

var signalScaleButtonPopupId int
var signalScaleButtonPopupMap = make(map[int]signalScaleButtonPopupDetail)
var signalScaleButtonPopupLock sync.RWMutex

// ScaleButtonSignalPopupCallback is a callback function for a 'popup' signal emitted from a ScaleButton.
type ScaleButtonSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ScaleButton) ConnectPopup(callback ScaleButtonSignalPopupCallback) int {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	signalScaleButtonPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popup(instance, C.gpointer(uintptr(signalScaleButtonPopupId)))

	detail := signalScaleButtonPopupDetail{callback, handlerID}
	signalScaleButtonPopupMap[signalScaleButtonPopupId] = detail

	return signalScaleButtonPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ScaleButton) DisconnectPopup(connectionID int) {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	detail, exists := signalScaleButtonPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopupMap, connectionID)
}

//export scalebutton_popupHandler
func scalebutton_popupHandler(_ *C.GObject, data C.gpointer) {
	signalScaleButtonPopupLock.RLock()
	defer signalScaleButtonPopupLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScaleButtonPopupMap[index].callback
	callback()
}

type signalScaleButtonValueChangedDetail struct {
	callback  ScaleButtonSignalValueChangedCallback
	handlerID C.gulong
}

var signalScaleButtonValueChangedId int
var signalScaleButtonValueChangedMap = make(map[int]signalScaleButtonValueChangedDetail)
var signalScaleButtonValueChangedLock sync.RWMutex

// ScaleButtonSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a ScaleButton.
type ScaleButtonSignalValueChangedCallback func(value float64)

/*
ConnectValueChanged connects the callback to the 'value-changed' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectValueChanged to remove it.
*/
func (recv *ScaleButton) ConnectValueChanged(callback ScaleButtonSignalValueChangedCallback) int {
	signalScaleButtonValueChangedLock.Lock()
	defer signalScaleButtonValueChangedLock.Unlock()

	signalScaleButtonValueChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_value_changed(instance, C.gpointer(uintptr(signalScaleButtonValueChangedId)))

	detail := signalScaleButtonValueChangedDetail{callback, handlerID}
	signalScaleButtonValueChangedMap[signalScaleButtonValueChangedId] = detail

	return signalScaleButtonValueChangedId
}

/*
DisconnectValueChanged disconnects a callback from the 'value-changed' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectValueChanged.
*/
func (recv *ScaleButton) DisconnectValueChanged(connectionID int) {
	signalScaleButtonValueChangedLock.Lock()
	defer signalScaleButtonValueChangedLock.Unlock()

	detail, exists := signalScaleButtonValueChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonValueChangedMap, connectionID)
}

//export scalebutton_valueChangedHandler
func scalebutton_valueChangedHandler(_ *C.GObject, c_value C.gdouble, data C.gpointer) {
	signalScaleButtonValueChangedLock.RLock()
	defer signalScaleButtonValueChangedLock.RUnlock()

	value := float64(c_value)

	index := int(uintptr(data))
	callback := signalScaleButtonValueChangedMap[index].callback
	callback(value)
}

// Blacklisted : gtk_scale_button_new

// Blacklisted : gtk_scale_button_get_adjustment

// Blacklisted : gtk_scale_button_get_value

// Blacklisted : gtk_scale_button_set_adjustment

// Blacklisted : gtk_scale_button_set_icons

// Blacklisted : gtk_scale_button_set_value

// Blacklisted : gtk_status_icon_get_screen

// Blacklisted : gtk_status_icon_set_screen

// Blacklisted : gtk_text_buffer_add_mark

// Blacklisted : gtk_text_mark_new

// Blacklisted : gtk_tool_item_set_tooltip_markup

// Blacklisted : gtk_tool_item_set_tooltip_text

// Blacklisted : gtk_tooltip_trigger_tooltip_query

// Blacklisted : gtk_tooltip_set_custom

// Blacklisted : gtk_tooltip_set_icon

// Blacklisted : gtk_tooltip_set_icon_from_stock

// Blacklisted : gtk_tooltip_set_markup

// SetText is a wrapper around the C function gtk_tooltip_set_text.
func (recv *Tooltip) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tooltip_set_text((*C.GtkTooltip)(recv.native), c_text)

	return
}

// Blacklisted : gtk_tooltip_set_tip_area

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter values :

// Blacklisted : gtk_tree_view_convert_bin_window_to_tree_coords

// Blacklisted : gtk_tree_view_convert_bin_window_to_widget_coords

// Blacklisted : gtk_tree_view_convert_tree_to_bin_window_coords

// Blacklisted : gtk_tree_view_convert_tree_to_widget_coords

// Blacklisted : gtk_tree_view_convert_widget_to_bin_window_coords

// Blacklisted : gtk_tree_view_convert_widget_to_tree_coords

// Blacklisted : gtk_tree_view_get_level_indentation

// Blacklisted : gtk_tree_view_get_show_expanders

// Blacklisted : gtk_tree_view_get_tooltip_column

// Blacklisted : gtk_tree_view_get_tooltip_context

// Blacklisted : gtk_tree_view_is_rubber_banding_active

// Blacklisted : gtk_tree_view_set_level_indentation

// Blacklisted : gtk_tree_view_set_show_expanders

// Blacklisted : gtk_tree_view_set_tooltip_cell

// Blacklisted : gtk_tree_view_set_tooltip_column

// Blacklisted : gtk_tree_view_set_tooltip_row

// Blacklisted : gtk_tree_view_column_get_tree_view

// Blacklisted : gtk_volume_button_new

type signalWidgetDragFailedDetail struct {
	callback  WidgetSignalDragFailedCallback
	handlerID C.gulong
}

var signalWidgetDragFailedId int
var signalWidgetDragFailedMap = make(map[int]signalWidgetDragFailedDetail)
var signalWidgetDragFailedLock sync.RWMutex

// WidgetSignalDragFailedCallback is a callback function for a 'drag-failed' signal emitted from a Widget.
type WidgetSignalDragFailedCallback func(context *gdk.DragContext, result DragResult) bool

/*
ConnectDragFailed connects the callback to the 'drag-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragFailed to remove it.
*/
func (recv *Widget) ConnectDragFailed(callback WidgetSignalDragFailedCallback) int {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	signalWidgetDragFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_failed(instance, C.gpointer(uintptr(signalWidgetDragFailedId)))

	detail := signalWidgetDragFailedDetail{callback, handlerID}
	signalWidgetDragFailedMap[signalWidgetDragFailedId] = detail

	return signalWidgetDragFailedId
}

/*
DisconnectDragFailed disconnects a callback from the 'drag-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragFailed.
*/
func (recv *Widget) DisconnectDragFailed(connectionID int) {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	detail, exists := signalWidgetDragFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragFailedMap, connectionID)
}

//export widget_dragFailedHandler
func widget_dragFailedHandler(_ *C.GObject, c_context *C.GdkDragContext, c_result C.GtkDragResult, data C.gpointer) C.gboolean {
	signalWidgetDragFailedLock.RLock()
	defer signalWidgetDragFailedLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	result := DragResult(c_result)

	index := int(uintptr(data))
	callback := signalWidgetDragFailedMap[index].callback
	retGo := callback(context, result)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetKeynavFailedDetail struct {
	callback  WidgetSignalKeynavFailedCallback
	handlerID C.gulong
}

var signalWidgetKeynavFailedId int
var signalWidgetKeynavFailedMap = make(map[int]signalWidgetKeynavFailedDetail)
var signalWidgetKeynavFailedLock sync.RWMutex

// WidgetSignalKeynavFailedCallback is a callback function for a 'keynav-failed' signal emitted from a Widget.
type WidgetSignalKeynavFailedCallback func(direction DirectionType) bool

/*
ConnectKeynavFailed connects the callback to the 'keynav-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectKeynavFailed to remove it.
*/
func (recv *Widget) ConnectKeynavFailed(callback WidgetSignalKeynavFailedCallback) int {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	signalWidgetKeynavFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_keynav_failed(instance, C.gpointer(uintptr(signalWidgetKeynavFailedId)))

	detail := signalWidgetKeynavFailedDetail{callback, handlerID}
	signalWidgetKeynavFailedMap[signalWidgetKeynavFailedId] = detail

	return signalWidgetKeynavFailedId
}

/*
DisconnectKeynavFailed disconnects a callback from the 'keynav-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeynavFailed.
*/
func (recv *Widget) DisconnectKeynavFailed(connectionID int) {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	detail, exists := signalWidgetKeynavFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetKeynavFailedMap, connectionID)
}

//export widget_keynavFailedHandler
func widget_keynavFailedHandler(_ *C.GObject, c_direction C.GtkDirectionType, data C.gpointer) C.gboolean {
	signalWidgetKeynavFailedLock.RLock()
	defer signalWidgetKeynavFailedLock.RUnlock()

	direction := DirectionType(c_direction)

	index := int(uintptr(data))
	callback := signalWidgetKeynavFailedMap[index].callback
	retGo := callback(direction)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetQueryTooltipDetail struct {
	callback  WidgetSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalWidgetQueryTooltipId int
var signalWidgetQueryTooltipMap = make(map[int]signalWidgetQueryTooltipDetail)
var signalWidgetQueryTooltipLock sync.RWMutex

// WidgetSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a Widget.
type WidgetSignalQueryTooltipCallback func(x int32, y int32, keyboardMode bool, tooltip *Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *Widget) ConnectQueryTooltip(callback WidgetSignalQueryTooltipCallback) int {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	signalWidgetQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalWidgetQueryTooltipId)))

	detail := signalWidgetQueryTooltipDetail{callback, handlerID}
	signalWidgetQueryTooltipMap[signalWidgetQueryTooltipId] = detail

	return signalWidgetQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the Widget.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *Widget) DisconnectQueryTooltip(connectionID int) {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	detail, exists := signalWidgetQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetQueryTooltipMap, connectionID)
}

//export widget_queryTooltipHandler
func widget_queryTooltipHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_keyboard_mode C.gboolean, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalWidgetQueryTooltipLock.RLock()
	defer signalWidgetQueryTooltipLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	keyboardMode := c_keyboard_mode == C.TRUE

	tooltip := TooltipNewFromC(unsafe.Pointer(c_tooltip))

	index := int(uintptr(data))
	callback := signalWidgetQueryTooltipMap[index].callback
	retGo := callback(x, y, keyboardMode, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_widget_error_bell

// Blacklisted : gtk_widget_get_has_tooltip

// Blacklisted : gtk_widget_get_tooltip_markup

// Blacklisted : gtk_widget_get_tooltip_text

// Blacklisted : gtk_widget_get_tooltip_window

// Blacklisted : gtk_widget_keynav_failed

// Blacklisted : gtk_widget_modify_cursor

// Blacklisted : gtk_widget_set_has_tooltip

// Blacklisted : gtk_widget_set_tooltip_markup

// Blacklisted : gtk_widget_set_tooltip_text

// Blacklisted : gtk_widget_set_tooltip_window

// Blacklisted : gtk_widget_trigger_tooltip_query

// Blacklisted : gtk_window_get_opacity

// Blacklisted : gtk_window_set_opacity

// Blacklisted : gtk_window_set_startup_id
