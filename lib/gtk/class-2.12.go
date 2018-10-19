// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void ComboBox_moveActiveHandler();

	static gulong ComboBox_signal_connect_move_active(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-active", ComboBox_moveActiveHandler, data);
	}

*/
/*

	void ComboBox_popdownHandler();

	static gulong ComboBox_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", ComboBox_popdownHandler, data);
	}

*/
/*

	void ComboBox_popupHandler();

	static gulong ComboBox_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", ComboBox_popupHandler, data);
	}

*/
/*

	void FileChooserButton_fileSetHandler();

	static gulong FileChooserButton_signal_connect_file_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-set", FileChooserButton_fileSetHandler, data);
	}

*/
/*

	void MenuShell_moveSelectedHandler();

	static gulong MenuShell_signal_connect_move_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-selected", MenuShell_moveSelectedHandler, data);
	}

*/
/*

	void ScaleButton_popdownHandler();

	static gulong ScaleButton_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", ScaleButton_popdownHandler, data);
	}

*/
/*

	void ScaleButton_popupHandler();

	static gulong ScaleButton_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", ScaleButton_popupHandler, data);
	}

*/
/*

	void ScaleButton_valueChangedHandler();

	static gulong ScaleButton_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", ScaleButton_valueChangedHandler, data);
	}

*/
/*

	void Widget_dragFailedHandler();

	static gulong Widget_signal_connect_drag_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-failed", Widget_dragFailedHandler, data);
	}

*/
/*

	void Widget_keynavFailedHandler();

	static gulong Widget_signal_connect_keynav_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keynav-failed", Widget_keynavFailedHandler, data);
	}

*/
/*

	void Widget_queryTooltipHandler();

	static gulong Widget_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", Widget_queryTooltipHandler, data);
	}

*/
import "C"

// GetProgramName is a wrapper around the C function gtk_about_dialog_get_program_name.
func (recv *AboutDialog) GetProgramName() string {
	retC := C.gtk_about_dialog_get_program_name((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetProgramName is a wrapper around the C function gtk_about_dialog_set_program_name.
func (recv *AboutDialog) SetProgramName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_about_dialog_set_program_name((*C.GtkAboutDialog)(recv.native), c_name)

	return
}

// CreateMenu is a wrapper around the C function gtk_action_create_menu.
func (recv *Action) CreateMenu() *Widget {
	retC := C.gtk_action_create_menu((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// BuilderNew is a wrapper around the C function gtk_builder_new.
func BuilderNew() *Builder {
	retC := C.gtk_builder_new()
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFromFile is a wrapper around the C function gtk_builder_add_from_file.
func (recv *Builder) AddFromFile(filename string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_file((*C.GtkBuilder)(recv.native), c_filename, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// AddFromString is a wrapper around the C function gtk_builder_add_from_string.
func (recv *Builder) AddFromString(buffer string, length uint64) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gsize)(length)

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_string((*C.GtkBuilder)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ConnectSignals is a wrapper around the C function gtk_builder_connect_signals.
func (recv *Builder) ConnectSignals(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

	C.gtk_builder_connect_signals((*C.GtkBuilder)(recv.native), c_user_data)

	return
}

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// GetObject is a wrapper around the C function gtk_builder_get_object.
func (recv *Builder) GetObject(name string) *gobject.Object {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_builder_get_object((*C.GtkBuilder)(recv.native), c_name)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObjects is a wrapper around the C function gtk_builder_get_objects.
func (recv *Builder) GetObjects() *glib.SList {
	retC := C.gtk_builder_get_objects((*C.GtkBuilder)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTranslationDomain is a wrapper around the C function gtk_builder_get_translation_domain.
func (recv *Builder) GetTranslationDomain() string {
	retC := C.gtk_builder_get_translation_domain((*C.GtkBuilder)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_builder_get_type_from_name : no return generator

// SetTranslationDomain is a wrapper around the C function gtk_builder_set_translation_domain.
func (recv *Builder) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_builder_set_translation_domain((*C.GtkBuilder)(recv.native), c_domain)

	return
}

// Unsupported : gtk_builder_value_from_string : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_builder_value_from_string_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

var signalComboBoxMoveActiveId int
var signalComboBoxMoveActiveMap = make(map[int]ComboBoxSignalMoveActiveCallback)
var signalComboBoxMoveActiveLock sync.Mutex

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
	signalComboBoxMoveActiveMap[signalComboBoxMoveActiveId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ComboBox_signal_connect_move_active(instance, C.gpointer(uintptr(signalComboBoxMoveActiveId)))
	return int(retC)
}

/*
DisconnectMoveActive disconnects a callback from the 'move-active' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectMoveActive.
*/
func (recv *ComboBox) DisconnectMoveActive(connectionID int) {
	signalComboBoxMoveActiveLock.Lock()
	defer signalComboBoxMoveActiveLock.Unlock()

	_, exists := signalComboBoxMoveActiveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalComboBoxMoveActiveMap, connectionID)
}

//export ComboBox_moveActiveHandler
func ComboBox_moveActiveHandler() {
	fmt.Println("cb")
}

var signalComboBoxPopdownId int
var signalComboBoxPopdownMap = make(map[int]ComboBoxSignalPopdownCallback)
var signalComboBoxPopdownLock sync.Mutex

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
	signalComboBoxPopdownMap[signalComboBoxPopdownId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ComboBox_signal_connect_popdown(instance, C.gpointer(uintptr(signalComboBoxPopdownId)))
	return int(retC)
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ComboBox) DisconnectPopdown(connectionID int) {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	_, exists := signalComboBoxPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalComboBoxPopdownMap, connectionID)
}

//export ComboBox_popdownHandler
func ComboBox_popdownHandler() C.boolean {
	fmt.Println("cb")
}

var signalComboBoxPopupId int
var signalComboBoxPopupMap = make(map[int]ComboBoxSignalPopupCallback)
var signalComboBoxPopupLock sync.Mutex

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
	signalComboBoxPopupMap[signalComboBoxPopupId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ComboBox_signal_connect_popup(instance, C.gpointer(uintptr(signalComboBoxPopupId)))
	return int(retC)
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ComboBox) DisconnectPopup(connectionID int) {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	_, exists := signalComboBoxPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalComboBoxPopupMap, connectionID)
}

//export ComboBox_popupHandler
func ComboBox_popupHandler() {
	fmt.Println("cb")
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetCursorHadjustment is a wrapper around the C function gtk_entry_get_cursor_hadjustment.
func (recv *Entry) GetCursorHadjustment() *Adjustment {
	retC := C.gtk_entry_get_cursor_hadjustment((*C.GtkEntry)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetCursorHadjustment is a wrapper around the C function gtk_entry_set_cursor_hadjustment.
func (recv *Entry) SetCursorHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_entry_set_cursor_hadjustment((*C.GtkEntry)(recv.native), c_adjustment)

	return
}

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// GetCompletionPrefix is a wrapper around the C function gtk_entry_completion_get_completion_prefix.
func (recv *EntryCompletion) GetCompletionPrefix() string {
	retC := C.gtk_entry_completion_get_completion_prefix((*C.GtkEntryCompletion)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInlineSelection is a wrapper around the C function gtk_entry_completion_get_inline_selection.
func (recv *EntryCompletion) GetInlineSelection() bool {
	retC := C.gtk_entry_completion_get_inline_selection((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetInlineSelection is a wrapper around the C function gtk_entry_completion_set_inline_selection.
func (recv *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	c_inline_selection :=
		boolToGboolean(inlineSelection)

	C.gtk_entry_completion_set_inline_selection((*C.GtkEntryCompletion)(recv.native), c_inline_selection)

	return
}

// Unsupported : EntryIconAccessible : no CType

var signalFileChooserButtonFileSetId int
var signalFileChooserButtonFileSetMap = make(map[int]FileChooserButtonSignalFileSetCallback)
var signalFileChooserButtonFileSetLock sync.Mutex

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
	signalFileChooserButtonFileSetMap[signalFileChooserButtonFileSetId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.FileChooserButton_signal_connect_file_set(instance, C.gpointer(uintptr(signalFileChooserButtonFileSetId)))
	return int(retC)
}

/*
DisconnectFileSet disconnects a callback from the 'file-set' signal for the FileChooserButton.

The connectionID should be a value returned from a call to ConnectFileSet.
*/
func (recv *FileChooserButton) DisconnectFileSet(connectionID int) {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	_, exists := signalFileChooserButtonFileSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalFileChooserButtonFileSetMap, connectionID)
}

//export FileChooserButton_fileSetHandler
func FileChooserButton_fileSetHandler() {
	fmt.Println("cb")
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// ListContexts is a wrapper around the C function gtk_icon_theme_list_contexts.
func (recv *IconTheme) ListContexts() *glib.List {
	retC := C.gtk_icon_theme_list_contexts((*C.GtkIconTheme)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// ConvertWidgetToBinWindowCoords is a wrapper around the C function gtk_icon_view_convert_widget_to_bin_window_coords.
func (recv *IconView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_icon_view_convert_widget_to_bin_window_coords((*C.GtkIconView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// GetTooltipColumn is a wrapper around the C function gtk_icon_view_get_tooltip_column.
func (recv *IconView) GetTooltipColumn() int32 {
	retC := C.gtk_icon_view_get_tooltip_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_icon_view_get_tooltip_context : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// SetTooltipCell is a wrapper around the C function gtk_icon_view_set_tooltip_cell.
func (recv *IconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(tooltip.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	C.gtk_icon_view_set_tooltip_cell((*C.GtkIconView)(recv.native), c_tooltip, c_path, c_cell)

	return
}

// SetTooltipColumn is a wrapper around the C function gtk_icon_view_set_tooltip_column.
func (recv *IconView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_tooltip_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// SetTooltipItem is a wrapper around the C function gtk_icon_view_set_tooltip_item.
func (recv *IconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(tooltip.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_icon_view_set_tooltip_item((*C.GtkIconView)(recv.native), c_tooltip, c_path)

	return
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter columns : no param type

var signalMenuShellMoveSelectedId int
var signalMenuShellMoveSelectedMap = make(map[int]MenuShellSignalMoveSelectedCallback)
var signalMenuShellMoveSelectedLock sync.Mutex

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
	signalMenuShellMoveSelectedMap[signalMenuShellMoveSelectedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.MenuShell_signal_connect_move_selected(instance, C.gpointer(uintptr(signalMenuShellMoveSelectedId)))
	return int(retC)
}

/*
DisconnectMoveSelected disconnects a callback from the 'move-selected' signal for the MenuShell.

The connectionID should be a value returned from a call to ConnectMoveSelected.
*/
func (recv *MenuShell) DisconnectMoveSelected(connectionID int) {
	signalMenuShellMoveSelectedLock.Lock()
	defer signalMenuShellMoveSelectedLock.Unlock()

	_, exists := signalMenuShellMoveSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalMenuShellMoveSelectedMap, connectionID)
}

//export MenuShell_moveSelectedHandler
func MenuShell_moveSelectedHandler() C.boolean {
	fmt.Println("cb")
}

// SetArrowTooltipMarkup is a wrapper around the C function gtk_menu_tool_button_set_arrow_tooltip_markup.
func (recv *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup((*C.GtkMenuToolButton)(recv.native), c_markup)

	return
}

// SetArrowTooltipText is a wrapper around the C function gtk_menu_tool_button_set_arrow_tooltip_text.
func (recv *MenuToolButton) SetArrowTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_menu_tool_button_set_arrow_tooltip_text((*C.GtkMenuToolButton)(recv.native), c_text)

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported signal 'create-window' for Notebook : return value Notebook :

// Unsupported signal 'get-child-position' for Overlay : unsupported parameter allocation : Blacklisted record : GdkRectangle

// PageSetupNewFromFile is a wrapper around the C function gtk_page_setup_new_from_file.
func PageSetupNewFromFile(fileName string) (*PageSetup, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_file(c_file_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// PageSetupNewFromKeyFile is a wrapper around the C function gtk_page_setup_new_from_key_file.
func PageSetupNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ToFile is a wrapper around the C function gtk_page_setup_to_file.
func (recv *PageSetup) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_to_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ToKeyFile is a wrapper around the C function gtk_page_setup_to_key_file.
func (recv *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_page_setup_to_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name)

	return
}

// Unsupported signal 'drag-action-ask' for PlacesSidebar : return value gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal 'populate-popup' for PlacesSidebar : unsupported parameter selected_item : no type generator for Gio.File,

// Unsupported signal 'create-custom-widget' for PrintOperation : return value GObject.Object :

// Unsupported signal 'preview' for PrintOperation : unsupported parameter preview : no type generator for PrintOperationPreview,

// PrintSettingsNewFromFile is a wrapper around the C function gtk_print_settings_new_from_file.
func PrintSettingsNewFromFile(fileName string) (*PrintSettings, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_file(c_file_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// PrintSettingsNewFromKeyFile is a wrapper around the C function gtk_print_settings_new_from_key_file.
func PrintSettingsNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PrintSettings, error) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ToFile is a wrapper around the C function gtk_print_settings_to_file.
func (recv *PrintSettings) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_to_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ToKeyFile is a wrapper around the C function gtk_print_settings_to_key_file.
func (recv *PrintSettings) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_print_settings_to_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name)

	return
}

// GetFillLevel is a wrapper around the C function gtk_range_get_fill_level.
func (recv *Range) GetFillLevel() float64 {
	retC := C.gtk_range_get_fill_level((*C.GtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetRestrictToFillLevel is a wrapper around the C function gtk_range_get_restrict_to_fill_level.
func (recv *Range) GetRestrictToFillLevel() bool {
	retC := C.gtk_range_get_restrict_to_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowFillLevel is a wrapper around the C function gtk_range_get_show_fill_level.
func (recv *Range) GetShowFillLevel() bool {
	retC := C.gtk_range_get_show_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFillLevel is a wrapper around the C function gtk_range_set_fill_level.
func (recv *Range) SetFillLevel(fillLevel float64) {
	c_fill_level := (C.gdouble)(fillLevel)

	C.gtk_range_set_fill_level((*C.GtkRange)(recv.native), c_fill_level)

	return
}

// SetRestrictToFillLevel is a wrapper around the C function gtk_range_set_restrict_to_fill_level.
func (recv *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	c_restrict_to_fill_level :=
		boolToGboolean(restrictToFillLevel)

	C.gtk_range_set_restrict_to_fill_level((*C.GtkRange)(recv.native), c_restrict_to_fill_level)

	return
}

// SetShowFillLevel is a wrapper around the C function gtk_range_set_show_fill_level.
func (recv *Range) SetShowFillLevel(showFillLevel bool) {
	c_show_fill_level :=
		boolToGboolean(showFillLevel)

	C.gtk_range_set_show_fill_level((*C.GtkRange)(recv.native), c_show_fill_level)

	return
}

// RecentActionNew is a wrapper around the C function gtk_recent_action_new.
func RecentActionNew(name string, label string, tooltip string, stockId string) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_recent_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentActionNewForManager is a wrapper around the C function gtk_recent_action_new_for_manager.
func RecentActionNewForManager(name string, label string, tooltip string, stockId string, manager *RecentManager) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_manager := (*C.GtkRecentManager)(manager.ToC())

	retC := C.gtk_recent_action_new_for_manager(c_name, c_label, c_tooltip, c_stock_id, c_manager)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShowNumbers is a wrapper around the C function gtk_recent_action_get_show_numbers.
func (recv *RecentAction) GetShowNumbers() bool {
	retC := C.gtk_recent_action_get_show_numbers((*C.GtkRecentAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowNumbers is a wrapper around the C function gtk_recent_action_set_show_numbers.
func (recv *RecentAction) SetShowNumbers(showNumbers bool) {
	c_show_numbers :=
		boolToGboolean(showNumbers)

	C.gtk_recent_action_set_show_numbers((*C.GtkRecentAction)(recv.native), c_show_numbers)

	return
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported signal 'format-value' for Scale : return value utf8 :

var signalScaleButtonPopdownId int
var signalScaleButtonPopdownMap = make(map[int]ScaleButtonSignalPopdownCallback)
var signalScaleButtonPopdownLock sync.Mutex

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
	signalScaleButtonPopdownMap[signalScaleButtonPopdownId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ScaleButton_signal_connect_popdown(instance, C.gpointer(uintptr(signalScaleButtonPopdownId)))
	return int(retC)
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ScaleButton) DisconnectPopdown(connectionID int) {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	_, exists := signalScaleButtonPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalScaleButtonPopdownMap, connectionID)
}

//export ScaleButton_popdownHandler
func ScaleButton_popdownHandler() {
	fmt.Println("cb")
}

var signalScaleButtonPopupId int
var signalScaleButtonPopupMap = make(map[int]ScaleButtonSignalPopupCallback)
var signalScaleButtonPopupLock sync.Mutex

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
	signalScaleButtonPopupMap[signalScaleButtonPopupId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ScaleButton_signal_connect_popup(instance, C.gpointer(uintptr(signalScaleButtonPopupId)))
	return int(retC)
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ScaleButton) DisconnectPopup(connectionID int) {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	_, exists := signalScaleButtonPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalScaleButtonPopupMap, connectionID)
}

//export ScaleButton_popupHandler
func ScaleButton_popupHandler() {
	fmt.Println("cb")
}

var signalScaleButtonValueChangedId int
var signalScaleButtonValueChangedMap = make(map[int]ScaleButtonSignalValueChangedCallback)
var signalScaleButtonValueChangedLock sync.Mutex

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
	signalScaleButtonValueChangedMap[signalScaleButtonValueChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ScaleButton_signal_connect_value_changed(instance, C.gpointer(uintptr(signalScaleButtonValueChangedId)))
	return int(retC)
}

/*
DisconnectValueChanged disconnects a callback from the 'value-changed' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectValueChanged.
*/
func (recv *ScaleButton) DisconnectValueChanged(connectionID int) {
	signalScaleButtonValueChangedLock.Lock()
	defer signalScaleButtonValueChangedLock.Unlock()

	_, exists := signalScaleButtonValueChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalScaleButtonValueChangedMap, connectionID)
}

//export ScaleButton_valueChangedHandler
func ScaleButton_valueChangedHandler() {
	fmt.Println("cb")
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// GetAdjustment is a wrapper around the C function gtk_scale_button_get_adjustment.
func (recv *ScaleButton) GetAdjustment() *Adjustment {
	retC := C.gtk_scale_button_get_adjustment((*C.GtkScaleButton)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValue is a wrapper around the C function gtk_scale_button_get_value.
func (recv *ScaleButton) GetValue() float64 {
	retC := C.gtk_scale_button_get_value((*C.GtkScaleButton)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetAdjustment is a wrapper around the C function gtk_scale_button_set_adjustment.
func (recv *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_scale_button_set_adjustment((*C.GtkScaleButton)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// SetValue is a wrapper around the C function gtk_scale_button_set_value.
func (recv *ScaleButton) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_scale_button_set_value((*C.GtkScaleButton)(recv.native), c_value)

	return
}

// Unsupported signal 'input' for SpinButton : return value gint :

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// GetScreen is a wrapper around the C function gtk_status_icon_get_screen.
func (recv *StatusIcon) GetScreen() *gdk.Screen {
	retC := C.gtk_status_icon_get_screen((*C.GtkStatusIcon)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetScreen is a wrapper around the C function gtk_status_icon_set_screen.
func (recv *StatusIcon) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_status_icon_set_screen((*C.GtkStatusIcon)(recv.native), c_screen)

	return
}

// AddMark is a wrapper around the C function gtk_text_buffer_add_mark.
func (recv *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	c_where := (*C.GtkTextIter)(where.ToC())

	C.gtk_text_buffer_add_mark((*C.GtkTextBuffer)(recv.native), c_mark, c_where)

	return
}

// TextMarkNew is a wrapper around the C function gtk_text_mark_new.
func TextMarkNew(name string, leftGravity bool) *TextMark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_left_gravity :=
		boolToGboolean(leftGravity)

	retC := C.gtk_text_mark_new(c_name, c_left_gravity)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

// SetTooltipMarkup is a wrapper around the C function gtk_tool_item_set_tooltip_markup.
func (recv *ToolItem) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tool_item_set_tooltip_markup((*C.GtkToolItem)(recv.native), c_markup)

	return
}

// SetTooltipText is a wrapper around the C function gtk_tool_item_set_tooltip_text.
func (recv *ToolItem) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tool_item_set_tooltip_text((*C.GtkToolItem)(recv.native), c_text)

	return
}

// SetCustom is a wrapper around the C function gtk_tooltip_set_custom.
func (recv *Tooltip) SetCustom(customWidget *Widget) {
	c_custom_widget := (*C.GtkWidget)(customWidget.ToC())

	C.gtk_tooltip_set_custom((*C.GtkTooltip)(recv.native), c_custom_widget)

	return
}

// SetIcon is a wrapper around the C function gtk_tooltip_set_icon.
func (recv *Tooltip) SetIcon(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_tooltip_set_icon((*C.GtkTooltip)(recv.native), c_pixbuf)

	return
}

// Unsupported : gtk_tooltip_set_icon_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// SetMarkup is a wrapper around the C function gtk_tooltip_set_markup.
func (recv *Tooltip) SetMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tooltip_set_markup((*C.GtkTooltip)(recv.native), c_markup)

	return
}

// SetText is a wrapper around the C function gtk_tooltip_set_text.
func (recv *Tooltip) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tooltip_set_text((*C.GtkTooltip)(recv.native), c_text)

	return
}

// Unsupported : gtk_tooltip_set_tip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// ConvertBinWindowToTreeCoords is a wrapper around the C function gtk_tree_view_convert_bin_window_to_tree_coords.
func (recv *TreeView) ConvertBinWindowToTreeCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_bin_window_to_tree_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// ConvertBinWindowToWidgetCoords is a wrapper around the C function gtk_tree_view_convert_bin_window_to_widget_coords.
func (recv *TreeView) ConvertBinWindowToWidgetCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_bin_window_to_widget_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// ConvertTreeToBinWindowCoords is a wrapper around the C function gtk_tree_view_convert_tree_to_bin_window_coords.
func (recv *TreeView) ConvertTreeToBinWindowCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_tree_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// ConvertTreeToWidgetCoords is a wrapper around the C function gtk_tree_view_convert_tree_to_widget_coords.
func (recv *TreeView) ConvertTreeToWidgetCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_tree_to_widget_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// ConvertWidgetToBinWindowCoords is a wrapper around the C function gtk_tree_view_convert_widget_to_bin_window_coords.
func (recv *TreeView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_widget_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// ConvertWidgetToTreeCoords is a wrapper around the C function gtk_tree_view_convert_widget_to_tree_coords.
func (recv *TreeView) ConvertWidgetToTreeCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_widget_to_tree_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// GetLevelIndentation is a wrapper around the C function gtk_tree_view_get_level_indentation.
func (recv *TreeView) GetLevelIndentation() int32 {
	retC := C.gtk_tree_view_get_level_indentation((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetShowExpanders is a wrapper around the C function gtk_tree_view_get_show_expanders.
func (recv *TreeView) GetShowExpanders() bool {
	retC := C.gtk_tree_view_get_show_expanders((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTooltipColumn is a wrapper around the C function gtk_tree_view_get_tooltip_column.
func (recv *TreeView) GetTooltipColumn() int32 {
	retC := C.gtk_tree_view_get_tooltip_column((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_get_tooltip_context : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// IsRubberBandingActive is a wrapper around the C function gtk_tree_view_is_rubber_banding_active.
func (recv *TreeView) IsRubberBandingActive() bool {
	retC := C.gtk_tree_view_is_rubber_banding_active((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLevelIndentation is a wrapper around the C function gtk_tree_view_set_level_indentation.
func (recv *TreeView) SetLevelIndentation(indentation int32) {
	c_indentation := (C.gint)(indentation)

	C.gtk_tree_view_set_level_indentation((*C.GtkTreeView)(recv.native), c_indentation)

	return
}

// SetShowExpanders is a wrapper around the C function gtk_tree_view_set_show_expanders.
func (recv *TreeView) SetShowExpanders(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_tree_view_set_show_expanders((*C.GtkTreeView)(recv.native), c_enabled)

	return
}

// SetTooltipCell is a wrapper around the C function gtk_tree_view_set_tooltip_cell.
func (recv *TreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(tooltip.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	C.gtk_tree_view_set_tooltip_cell((*C.GtkTreeView)(recv.native), c_tooltip, c_path, c_column, c_cell)

	return
}

// SetTooltipColumn is a wrapper around the C function gtk_tree_view_set_tooltip_column.
func (recv *TreeView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_view_set_tooltip_column((*C.GtkTreeView)(recv.native), c_column)

	return
}

// SetTooltipRow is a wrapper around the C function gtk_tree_view_set_tooltip_row.
func (recv *TreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(tooltip.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_view_set_tooltip_row((*C.GtkTreeView)(recv.native), c_tooltip, c_path)

	return
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetTreeView is a wrapper around the C function gtk_tree_view_column_get_tree_view.
func (recv *TreeViewColumn) GetTreeView() *Widget {
	retC := C.gtk_tree_view_column_get_tree_view((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VolumeButtonNew is a wrapper around the C function gtk_volume_button_new.
func VolumeButtonNew() *VolumeButton {
	retC := C.gtk_volume_button_new()
	retGo := VolumeButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

var signalWidgetDragFailedId int
var signalWidgetDragFailedMap = make(map[int]WidgetSignalDragFailedCallback)
var signalWidgetDragFailedLock sync.Mutex

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
	signalWidgetDragFailedMap[signalWidgetDragFailedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Widget_signal_connect_drag_failed(instance, C.gpointer(uintptr(signalWidgetDragFailedId)))
	return int(retC)
}

/*
DisconnectDragFailed disconnects a callback from the 'drag-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragFailed.
*/
func (recv *Widget) DisconnectDragFailed(connectionID int) {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	_, exists := signalWidgetDragFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWidgetDragFailedMap, connectionID)
}

//export Widget_dragFailedHandler
func Widget_dragFailedHandler() C.boolean {
	fmt.Println("cb")
}

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

var signalWidgetKeynavFailedId int
var signalWidgetKeynavFailedMap = make(map[int]WidgetSignalKeynavFailedCallback)
var signalWidgetKeynavFailedLock sync.Mutex

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
	signalWidgetKeynavFailedMap[signalWidgetKeynavFailedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Widget_signal_connect_keynav_failed(instance, C.gpointer(uintptr(signalWidgetKeynavFailedId)))
	return int(retC)
}

/*
DisconnectKeynavFailed disconnects a callback from the 'keynav-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeynavFailed.
*/
func (recv *Widget) DisconnectKeynavFailed(connectionID int) {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	_, exists := signalWidgetKeynavFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWidgetKeynavFailedMap, connectionID)
}

//export Widget_keynavFailedHandler
func Widget_keynavFailedHandler() C.boolean {
	fmt.Println("cb")
}

var signalWidgetQueryTooltipId int
var signalWidgetQueryTooltipMap = make(map[int]WidgetSignalQueryTooltipCallback)
var signalWidgetQueryTooltipLock sync.Mutex

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
	signalWidgetQueryTooltipMap[signalWidgetQueryTooltipId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Widget_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalWidgetQueryTooltipId)))
	return int(retC)
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the Widget.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *Widget) DisconnectQueryTooltip(connectionID int) {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	_, exists := signalWidgetQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWidgetQueryTooltipMap, connectionID)
}

//export Widget_queryTooltipHandler
func Widget_queryTooltipHandler() C.boolean {
	fmt.Println("cb")
}

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// ErrorBell is a wrapper around the C function gtk_widget_error_bell.
func (recv *Widget) ErrorBell() {
	C.gtk_widget_error_bell((*C.GtkWidget)(recv.native))

	return
}

// GetHasTooltip is a wrapper around the C function gtk_widget_get_has_tooltip.
func (recv *Widget) GetHasTooltip() bool {
	retC := C.gtk_widget_get_has_tooltip((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTooltipMarkup is a wrapper around the C function gtk_widget_get_tooltip_markup.
func (recv *Widget) GetTooltipMarkup() string {
	retC := C.gtk_widget_get_tooltip_markup((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipText is a wrapper around the C function gtk_widget_get_tooltip_text.
func (recv *Widget) GetTooltipText() string {
	retC := C.gtk_widget_get_tooltip_text((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipWindow is a wrapper around the C function gtk_widget_get_tooltip_window.
func (recv *Widget) GetTooltipWindow() *Window {
	retC := C.gtk_widget_get_tooltip_window((*C.GtkWidget)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// KeynavFailed is a wrapper around the C function gtk_widget_keynav_failed.
func (recv *Widget) KeynavFailed(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_widget_keynav_failed((*C.GtkWidget)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// ModifyCursor is a wrapper around the C function gtk_widget_modify_cursor.
func (recv *Widget) ModifyCursor(primary *gdk.Color, secondary *gdk.Color) {
	c_primary := (*C.GdkColor)(primary.ToC())

	c_secondary := (*C.GdkColor)(secondary.ToC())

	C.gtk_widget_modify_cursor((*C.GtkWidget)(recv.native), c_primary, c_secondary)

	return
}

// SetHasTooltip is a wrapper around the C function gtk_widget_set_has_tooltip.
func (recv *Widget) SetHasTooltip(hasTooltip bool) {
	c_has_tooltip :=
		boolToGboolean(hasTooltip)

	C.gtk_widget_set_has_tooltip((*C.GtkWidget)(recv.native), c_has_tooltip)

	return
}

// SetTooltipMarkup is a wrapper around the C function gtk_widget_set_tooltip_markup.
func (recv *Widget) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_widget_set_tooltip_markup((*C.GtkWidget)(recv.native), c_markup)

	return
}

// SetTooltipText is a wrapper around the C function gtk_widget_set_tooltip_text.
func (recv *Widget) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_widget_set_tooltip_text((*C.GtkWidget)(recv.native), c_text)

	return
}

// SetTooltipWindow is a wrapper around the C function gtk_widget_set_tooltip_window.
func (recv *Widget) SetTooltipWindow(customWindow *Window) {
	c_custom_window := (*C.GtkWindow)(customWindow.ToC())

	C.gtk_widget_set_tooltip_window((*C.GtkWidget)(recv.native), c_custom_window)

	return
}

// TriggerTooltipQuery is a wrapper around the C function gtk_widget_trigger_tooltip_query.
func (recv *Widget) TriggerTooltipQuery() {
	C.gtk_widget_trigger_tooltip_query((*C.GtkWidget)(recv.native))

	return
}

// GetOpacity is a wrapper around the C function gtk_window_get_opacity.
func (recv *Window) GetOpacity() float64 {
	retC := C.gtk_window_get_opacity((*C.GtkWindow)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetOpacity is a wrapper around the C function gtk_window_set_opacity.
func (recv *Window) SetOpacity(opacity float64) {
	c_opacity := (C.gdouble)(opacity)

	C.gtk_window_set_opacity((*C.GtkWindow)(recv.native), c_opacity)

	return
}

// SetStartupId is a wrapper around the C function gtk_window_set_startup_id.
func (recv *Window) SetStartupId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gtk_window_set_startup_id((*C.GtkWindow)(recv.native), c_startup_id)

	return
}
