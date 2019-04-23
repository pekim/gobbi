// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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

	void assistant_applyHandler(GObject *, gpointer);

	static gulong Assistant_signal_connect_apply(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "apply", G_CALLBACK(assistant_applyHandler), data);
	}

*/
/*

	void assistant_cancelHandler(GObject *, gpointer);

	static gulong Assistant_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", G_CALLBACK(assistant_cancelHandler), data);
	}

*/
/*

	void assistant_closeHandler(GObject *, gpointer);

	static gulong Assistant_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", G_CALLBACK(assistant_closeHandler), data);
	}

*/
/*

	void assistant_prepareHandler(GObject *, GtkWidget *, gpointer);

	static gulong Assistant_signal_connect_prepare(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "prepare", G_CALLBACK(assistant_prepareHandler), data);
	}

*/
/*

	void cellrendereraccel_accelClearedHandler(GObject *, gchar*, gpointer);

	static gulong CellRendererAccel_signal_connect_accel_cleared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-cleared", G_CALLBACK(cellrendereraccel_accelClearedHandler), data);
	}

*/
/*

	void cellrendereraccel_accelEditedHandler(GObject *, gchar*, guint, guint, guint, gpointer);

	static gulong CellRendererAccel_signal_connect_accel_edited(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-edited", G_CALLBACK(cellrendereraccel_accelEditedHandler), data);
	}

*/
/*

	void notebook_pageAddedHandler(GObject *, GtkWidget *, guint, gpointer);

	static gulong Notebook_signal_connect_page_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "page-added", G_CALLBACK(notebook_pageAddedHandler), data);
	}

*/
/*

	void notebook_pageRemovedHandler(GObject *, GtkWidget *, guint, gpointer);

	static gulong Notebook_signal_connect_page_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "page-removed", G_CALLBACK(notebook_pageRemovedHandler), data);
	}

*/
/*

	void notebook_pageReorderedHandler(GObject *, GtkWidget *, guint, gpointer);

	static gulong Notebook_signal_connect_page_reordered(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "page-reordered", G_CALLBACK(notebook_pageReorderedHandler), data);
	}

*/
/*

	void printoperation_beginPrintHandler(GObject *, GtkPrintContext *, gpointer);

	static gulong PrintOperation_signal_connect_begin_print(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "begin-print", G_CALLBACK(printoperation_beginPrintHandler), data);
	}

*/
/*

	GObject * printoperation_createCustomWidgetHandler(GObject *, gpointer);

	static gulong PrintOperation_signal_connect_create_custom_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-custom-widget", G_CALLBACK(printoperation_createCustomWidgetHandler), data);
	}

*/
/*

	void printoperation_customWidgetApplyHandler(GObject *, GtkWidget *, gpointer);

	static gulong PrintOperation_signal_connect_custom_widget_apply(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "custom-widget-apply", G_CALLBACK(printoperation_customWidgetApplyHandler), data);
	}

*/
/*

	void printoperation_doneHandler(GObject *, GtkPrintOperationResult, gpointer);

	static gulong PrintOperation_signal_connect_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "done", G_CALLBACK(printoperation_doneHandler), data);
	}

*/
/*

	void printoperation_drawPageHandler(GObject *, GtkPrintContext *, gint, gpointer);

	static gulong PrintOperation_signal_connect_draw_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw-page", G_CALLBACK(printoperation_drawPageHandler), data);
	}

*/
/*

	void printoperation_endPrintHandler(GObject *, GtkPrintContext *, gpointer);

	static gulong PrintOperation_signal_connect_end_print(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "end-print", G_CALLBACK(printoperation_endPrintHandler), data);
	}

*/
/*

	gboolean printoperation_paginateHandler(GObject *, GtkPrintContext *, gpointer);

	static gulong PrintOperation_signal_connect_paginate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paginate", G_CALLBACK(printoperation_paginateHandler), data);
	}

*/
/*

	gboolean printoperation_previewHandler(GObject *, GtkPrintOperationPreview *, GtkPrintContext *, GtkWindow *, gpointer);

	static gulong PrintOperation_signal_connect_preview(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preview", G_CALLBACK(printoperation_previewHandler), data);
	}

*/
/*

	void printoperation_requestPageSetupHandler(GObject *, GtkPrintContext *, gint, GtkPageSetup *, gpointer);

	static gulong PrintOperation_signal_connect_request_page_setup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-page-setup", G_CALLBACK(printoperation_requestPageSetupHandler), data);
	}

*/
/*

	void printoperation_statusChangedHandler(GObject *, gpointer);

	static gulong PrintOperation_signal_connect_status_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "status-changed", G_CALLBACK(printoperation_statusChangedHandler), data);
	}

*/
/*

	void recentmanager_changedHandler(GObject *, gpointer);

	static gulong RecentManager_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(recentmanager_changedHandler), data);
	}

*/
/*

	void spinbutton_wrappedHandler(GObject *, gpointer);

	static gulong SpinButton_signal_connect_wrapped(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrapped", G_CALLBACK(spinbutton_wrappedHandler), data);
	}

*/
/*

	void statusicon_activateHandler(GObject *, gpointer);

	static gulong StatusIcon_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(statusicon_activateHandler), data);
	}

*/
/*

	void statusicon_popupMenuHandler(GObject *, guint, guint, gpointer);

	static gulong StatusIcon_signal_connect_popup_menu(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup-menu", G_CALLBACK(statusicon_popupMenuHandler), data);
	}

*/
/*

	gboolean statusicon_sizeChangedHandler(GObject *, gint, gpointer);

	static gulong StatusIcon_signal_connect_size_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-changed", G_CALLBACK(statusicon_sizeChangedHandler), data);
	}

*/
import "C"

type signalAssistantApplyDetail struct {
	callback  AssistantSignalApplyCallback
	handlerID C.gulong
}

var signalAssistantApplyId int
var signalAssistantApplyMap = make(map[int]signalAssistantApplyDetail)
var signalAssistantApplyLock sync.RWMutex

// AssistantSignalApplyCallback is a callback function for a 'apply' signal emitted from a Assistant.
type AssistantSignalApplyCallback func()

/*
ConnectApply connects the callback to the 'apply' signal for the Assistant.

The returned value represents the connection, and may be passed to DisconnectApply to remove it.
*/
func (recv *Assistant) ConnectApply(callback AssistantSignalApplyCallback) int {
	signalAssistantApplyLock.Lock()
	defer signalAssistantApplyLock.Unlock()

	signalAssistantApplyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Assistant_signal_connect_apply(instance, C.gpointer(uintptr(signalAssistantApplyId)))

	detail := signalAssistantApplyDetail{callback, handlerID}
	signalAssistantApplyMap[signalAssistantApplyId] = detail

	return signalAssistantApplyId
}

/*
DisconnectApply disconnects a callback from the 'apply' signal for the Assistant.

The connectionID should be a value returned from a call to ConnectApply.
*/
func (recv *Assistant) DisconnectApply(connectionID int) {
	signalAssistantApplyLock.Lock()
	defer signalAssistantApplyLock.Unlock()

	detail, exists := signalAssistantApplyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAssistantApplyMap, connectionID)
}

//export assistant_applyHandler
func assistant_applyHandler(_ *C.GObject, data C.gpointer) {
	signalAssistantApplyLock.RLock()
	defer signalAssistantApplyLock.RUnlock()

	index := int(uintptr(data))
	callback := signalAssistantApplyMap[index].callback
	callback()
}

type signalAssistantCancelDetail struct {
	callback  AssistantSignalCancelCallback
	handlerID C.gulong
}

var signalAssistantCancelId int
var signalAssistantCancelMap = make(map[int]signalAssistantCancelDetail)
var signalAssistantCancelLock sync.RWMutex

// AssistantSignalCancelCallback is a callback function for a 'cancel' signal emitted from a Assistant.
type AssistantSignalCancelCallback func()

/*
ConnectCancel connects the callback to the 'cancel' signal for the Assistant.

The returned value represents the connection, and may be passed to DisconnectCancel to remove it.
*/
func (recv *Assistant) ConnectCancel(callback AssistantSignalCancelCallback) int {
	signalAssistantCancelLock.Lock()
	defer signalAssistantCancelLock.Unlock()

	signalAssistantCancelId++
	instance := C.gpointer(recv.native)
	handlerID := C.Assistant_signal_connect_cancel(instance, C.gpointer(uintptr(signalAssistantCancelId)))

	detail := signalAssistantCancelDetail{callback, handlerID}
	signalAssistantCancelMap[signalAssistantCancelId] = detail

	return signalAssistantCancelId
}

/*
DisconnectCancel disconnects a callback from the 'cancel' signal for the Assistant.

The connectionID should be a value returned from a call to ConnectCancel.
*/
func (recv *Assistant) DisconnectCancel(connectionID int) {
	signalAssistantCancelLock.Lock()
	defer signalAssistantCancelLock.Unlock()

	detail, exists := signalAssistantCancelMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAssistantCancelMap, connectionID)
}

//export assistant_cancelHandler
func assistant_cancelHandler(_ *C.GObject, data C.gpointer) {
	signalAssistantCancelLock.RLock()
	defer signalAssistantCancelLock.RUnlock()

	index := int(uintptr(data))
	callback := signalAssistantCancelMap[index].callback
	callback()
}

type signalAssistantCloseDetail struct {
	callback  AssistantSignalCloseCallback
	handlerID C.gulong
}

var signalAssistantCloseId int
var signalAssistantCloseMap = make(map[int]signalAssistantCloseDetail)
var signalAssistantCloseLock sync.RWMutex

// AssistantSignalCloseCallback is a callback function for a 'close' signal emitted from a Assistant.
type AssistantSignalCloseCallback func()

/*
ConnectClose connects the callback to the 'close' signal for the Assistant.

The returned value represents the connection, and may be passed to DisconnectClose to remove it.
*/
func (recv *Assistant) ConnectClose(callback AssistantSignalCloseCallback) int {
	signalAssistantCloseLock.Lock()
	defer signalAssistantCloseLock.Unlock()

	signalAssistantCloseId++
	instance := C.gpointer(recv.native)
	handlerID := C.Assistant_signal_connect_close(instance, C.gpointer(uintptr(signalAssistantCloseId)))

	detail := signalAssistantCloseDetail{callback, handlerID}
	signalAssistantCloseMap[signalAssistantCloseId] = detail

	return signalAssistantCloseId
}

/*
DisconnectClose disconnects a callback from the 'close' signal for the Assistant.

The connectionID should be a value returned from a call to ConnectClose.
*/
func (recv *Assistant) DisconnectClose(connectionID int) {
	signalAssistantCloseLock.Lock()
	defer signalAssistantCloseLock.Unlock()

	detail, exists := signalAssistantCloseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAssistantCloseMap, connectionID)
}

//export assistant_closeHandler
func assistant_closeHandler(_ *C.GObject, data C.gpointer) {
	signalAssistantCloseLock.RLock()
	defer signalAssistantCloseLock.RUnlock()

	index := int(uintptr(data))
	callback := signalAssistantCloseMap[index].callback
	callback()
}

type signalAssistantPrepareDetail struct {
	callback  AssistantSignalPrepareCallback
	handlerID C.gulong
}

var signalAssistantPrepareId int
var signalAssistantPrepareMap = make(map[int]signalAssistantPrepareDetail)
var signalAssistantPrepareLock sync.RWMutex

// AssistantSignalPrepareCallback is a callback function for a 'prepare' signal emitted from a Assistant.
type AssistantSignalPrepareCallback func(page *Widget)

/*
ConnectPrepare connects the callback to the 'prepare' signal for the Assistant.

The returned value represents the connection, and may be passed to DisconnectPrepare to remove it.
*/
func (recv *Assistant) ConnectPrepare(callback AssistantSignalPrepareCallback) int {
	signalAssistantPrepareLock.Lock()
	defer signalAssistantPrepareLock.Unlock()

	signalAssistantPrepareId++
	instance := C.gpointer(recv.native)
	handlerID := C.Assistant_signal_connect_prepare(instance, C.gpointer(uintptr(signalAssistantPrepareId)))

	detail := signalAssistantPrepareDetail{callback, handlerID}
	signalAssistantPrepareMap[signalAssistantPrepareId] = detail

	return signalAssistantPrepareId
}

/*
DisconnectPrepare disconnects a callback from the 'prepare' signal for the Assistant.

The connectionID should be a value returned from a call to ConnectPrepare.
*/
func (recv *Assistant) DisconnectPrepare(connectionID int) {
	signalAssistantPrepareLock.Lock()
	defer signalAssistantPrepareLock.Unlock()

	detail, exists := signalAssistantPrepareMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAssistantPrepareMap, connectionID)
}

//export assistant_prepareHandler
func assistant_prepareHandler(_ *C.GObject, c_page *C.GtkWidget, data C.gpointer) {
	signalAssistantPrepareLock.RLock()
	defer signalAssistantPrepareLock.RUnlock()

	page := WidgetNewFromC(unsafe.Pointer(c_page))

	index := int(uintptr(data))
	callback := signalAssistantPrepareMap[index].callback
	callback(page)
}

// Blacklisted : gtk_assistant_new

// Blacklisted : gtk_assistant_add_action_widget

// Blacklisted : gtk_assistant_append_page

// Blacklisted : gtk_assistant_get_current_page

// Blacklisted : gtk_assistant_get_n_pages

// Blacklisted : gtk_assistant_get_nth_page

// Blacklisted : gtk_assistant_get_page_complete

// Blacklisted : gtk_assistant_get_page_header_image

// Blacklisted : gtk_assistant_get_page_side_image

// Blacklisted : gtk_assistant_get_page_title

// Blacklisted : gtk_assistant_get_page_type

// Blacklisted : gtk_assistant_insert_page

// Blacklisted : gtk_assistant_prepend_page

// Blacklisted : gtk_assistant_remove_action_widget

// Blacklisted : gtk_assistant_set_current_page

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc (GtkAssistantPageFunc) for param page_func

// Blacklisted : gtk_assistant_set_page_complete

// Blacklisted : gtk_assistant_set_page_header_image

// Blacklisted : gtk_assistant_set_page_side_image

// Blacklisted : gtk_assistant_set_page_title

// Blacklisted : gtk_assistant_set_page_type

// Blacklisted : gtk_assistant_update_buttons_state

// Blacklisted : gtk_button_get_image_position

// Blacklisted : gtk_button_set_image_position

type signalCellRendererAccelAccelClearedDetail struct {
	callback  CellRendererAccelSignalAccelClearedCallback
	handlerID C.gulong
}

var signalCellRendererAccelAccelClearedId int
var signalCellRendererAccelAccelClearedMap = make(map[int]signalCellRendererAccelAccelClearedDetail)
var signalCellRendererAccelAccelClearedLock sync.RWMutex

// CellRendererAccelSignalAccelClearedCallback is a callback function for a 'accel-cleared' signal emitted from a CellRendererAccel.
type CellRendererAccelSignalAccelClearedCallback func(pathString string)

/*
ConnectAccelCleared connects the callback to the 'accel-cleared' signal for the CellRendererAccel.

The returned value represents the connection, and may be passed to DisconnectAccelCleared to remove it.
*/
func (recv *CellRendererAccel) ConnectAccelCleared(callback CellRendererAccelSignalAccelClearedCallback) int {
	signalCellRendererAccelAccelClearedLock.Lock()
	defer signalCellRendererAccelAccelClearedLock.Unlock()

	signalCellRendererAccelAccelClearedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRendererAccel_signal_connect_accel_cleared(instance, C.gpointer(uintptr(signalCellRendererAccelAccelClearedId)))

	detail := signalCellRendererAccelAccelClearedDetail{callback, handlerID}
	signalCellRendererAccelAccelClearedMap[signalCellRendererAccelAccelClearedId] = detail

	return signalCellRendererAccelAccelClearedId
}

/*
DisconnectAccelCleared disconnects a callback from the 'accel-cleared' signal for the CellRendererAccel.

The connectionID should be a value returned from a call to ConnectAccelCleared.
*/
func (recv *CellRendererAccel) DisconnectAccelCleared(connectionID int) {
	signalCellRendererAccelAccelClearedLock.Lock()
	defer signalCellRendererAccelAccelClearedLock.Unlock()

	detail, exists := signalCellRendererAccelAccelClearedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererAccelAccelClearedMap, connectionID)
}

//export cellrendereraccel_accelClearedHandler
func cellrendereraccel_accelClearedHandler(_ *C.GObject, c_path_string *C.gchar, data C.gpointer) {
	signalCellRendererAccelAccelClearedLock.RLock()
	defer signalCellRendererAccelAccelClearedLock.RUnlock()

	pathString := C.GoString(c_path_string)

	index := int(uintptr(data))
	callback := signalCellRendererAccelAccelClearedMap[index].callback
	callback(pathString)
}

type signalCellRendererAccelAccelEditedDetail struct {
	callback  CellRendererAccelSignalAccelEditedCallback
	handlerID C.gulong
}

var signalCellRendererAccelAccelEditedId int
var signalCellRendererAccelAccelEditedMap = make(map[int]signalCellRendererAccelAccelEditedDetail)
var signalCellRendererAccelAccelEditedLock sync.RWMutex

// CellRendererAccelSignalAccelEditedCallback is a callback function for a 'accel-edited' signal emitted from a CellRendererAccel.
type CellRendererAccelSignalAccelEditedCallback func(pathString string, accelKey uint32, accelMods gdk.ModifierType, hardwareKeycode uint32)

/*
ConnectAccelEdited connects the callback to the 'accel-edited' signal for the CellRendererAccel.

The returned value represents the connection, and may be passed to DisconnectAccelEdited to remove it.
*/
func (recv *CellRendererAccel) ConnectAccelEdited(callback CellRendererAccelSignalAccelEditedCallback) int {
	signalCellRendererAccelAccelEditedLock.Lock()
	defer signalCellRendererAccelAccelEditedLock.Unlock()

	signalCellRendererAccelAccelEditedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRendererAccel_signal_connect_accel_edited(instance, C.gpointer(uintptr(signalCellRendererAccelAccelEditedId)))

	detail := signalCellRendererAccelAccelEditedDetail{callback, handlerID}
	signalCellRendererAccelAccelEditedMap[signalCellRendererAccelAccelEditedId] = detail

	return signalCellRendererAccelAccelEditedId
}

/*
DisconnectAccelEdited disconnects a callback from the 'accel-edited' signal for the CellRendererAccel.

The connectionID should be a value returned from a call to ConnectAccelEdited.
*/
func (recv *CellRendererAccel) DisconnectAccelEdited(connectionID int) {
	signalCellRendererAccelAccelEditedLock.Lock()
	defer signalCellRendererAccelAccelEditedLock.Unlock()

	detail, exists := signalCellRendererAccelAccelEditedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererAccelAccelEditedMap, connectionID)
}

//export cellrendereraccel_accelEditedHandler
func cellrendereraccel_accelEditedHandler(_ *C.GObject, c_path_string *C.gchar, c_accel_key C.guint, c_accel_mods C.guint, c_hardware_keycode C.guint, data C.gpointer) {
	signalCellRendererAccelAccelEditedLock.RLock()
	defer signalCellRendererAccelAccelEditedLock.RUnlock()

	pathString := C.GoString(c_path_string)

	accelKey := uint32(c_accel_key)

	accelMods := gdk.ModifierType(c_accel_mods)

	hardwareKeycode := uint32(c_hardware_keycode)

	index := int(uintptr(data))
	callback := signalCellRendererAccelAccelEditedMap[index].callback
	callback(pathString, accelKey, accelMods, hardwareKeycode)
}

// Blacklisted : gtk_cell_renderer_accel_new

// Blacklisted : gtk_cell_renderer_spin_new

// Unsupported : gtk_clipboard_request_rich_text : unsupported parameter callback : no type generator for ClipboardRichTextReceivedFunc (GtkClipboardRichTextReceivedFunc) for param callback

// Unsupported : gtk_clipboard_wait_for_rich_text : array return type :

// Blacklisted : gtk_clipboard_wait_is_rich_text_available

// GetTitle is a wrapper around the C function gtk_combo_box_get_title.
func (recv *ComboBox) GetTitle() string {
	retC := C.gtk_combo_box_get_title((*C.GtkComboBox)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetTitle is a wrapper around the C function gtk_combo_box_set_title.
func (recv *ComboBox) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_combo_box_set_title((*C.GtkComboBox)(recv.native), c_title)

	return
}

// Blacklisted : gtk_entry_get_inner_border

// Blacklisted : gtk_entry_set_inner_border

// Blacklisted : gtk_file_chooser_button_get_focus_on_click

// Blacklisted : gtk_file_chooser_button_set_focus_on_click

// Blacklisted : gtk_label_get_line_wrap_mode

// Blacklisted : gtk_label_set_line_wrap_mode

// Blacklisted : gtk_link_button_new

// Blacklisted : gtk_link_button_new_with_label

// Blacklisted : gtk_link_button_get_uri

// Blacklisted : gtk_link_button_set_uri

// Blacklisted : gtk_message_dialog_set_image

type signalNotebookPageAddedDetail struct {
	callback  NotebookSignalPageAddedCallback
	handlerID C.gulong
}

var signalNotebookPageAddedId int
var signalNotebookPageAddedMap = make(map[int]signalNotebookPageAddedDetail)
var signalNotebookPageAddedLock sync.RWMutex

// NotebookSignalPageAddedCallback is a callback function for a 'page-added' signal emitted from a Notebook.
type NotebookSignalPageAddedCallback func(child *Widget, pageNum uint32)

/*
ConnectPageAdded connects the callback to the 'page-added' signal for the Notebook.

The returned value represents the connection, and may be passed to DisconnectPageAdded to remove it.
*/
func (recv *Notebook) ConnectPageAdded(callback NotebookSignalPageAddedCallback) int {
	signalNotebookPageAddedLock.Lock()
	defer signalNotebookPageAddedLock.Unlock()

	signalNotebookPageAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Notebook_signal_connect_page_added(instance, C.gpointer(uintptr(signalNotebookPageAddedId)))

	detail := signalNotebookPageAddedDetail{callback, handlerID}
	signalNotebookPageAddedMap[signalNotebookPageAddedId] = detail

	return signalNotebookPageAddedId
}

/*
DisconnectPageAdded disconnects a callback from the 'page-added' signal for the Notebook.

The connectionID should be a value returned from a call to ConnectPageAdded.
*/
func (recv *Notebook) DisconnectPageAdded(connectionID int) {
	signalNotebookPageAddedLock.Lock()
	defer signalNotebookPageAddedLock.Unlock()

	detail, exists := signalNotebookPageAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNotebookPageAddedMap, connectionID)
}

//export notebook_pageAddedHandler
func notebook_pageAddedHandler(_ *C.GObject, c_child *C.GtkWidget, c_page_num C.guint, data C.gpointer) {
	signalNotebookPageAddedLock.RLock()
	defer signalNotebookPageAddedLock.RUnlock()

	child := WidgetNewFromC(unsafe.Pointer(c_child))

	pageNum := uint32(c_page_num)

	index := int(uintptr(data))
	callback := signalNotebookPageAddedMap[index].callback
	callback(child, pageNum)
}

type signalNotebookPageRemovedDetail struct {
	callback  NotebookSignalPageRemovedCallback
	handlerID C.gulong
}

var signalNotebookPageRemovedId int
var signalNotebookPageRemovedMap = make(map[int]signalNotebookPageRemovedDetail)
var signalNotebookPageRemovedLock sync.RWMutex

// NotebookSignalPageRemovedCallback is a callback function for a 'page-removed' signal emitted from a Notebook.
type NotebookSignalPageRemovedCallback func(child *Widget, pageNum uint32)

/*
ConnectPageRemoved connects the callback to the 'page-removed' signal for the Notebook.

The returned value represents the connection, and may be passed to DisconnectPageRemoved to remove it.
*/
func (recv *Notebook) ConnectPageRemoved(callback NotebookSignalPageRemovedCallback) int {
	signalNotebookPageRemovedLock.Lock()
	defer signalNotebookPageRemovedLock.Unlock()

	signalNotebookPageRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Notebook_signal_connect_page_removed(instance, C.gpointer(uintptr(signalNotebookPageRemovedId)))

	detail := signalNotebookPageRemovedDetail{callback, handlerID}
	signalNotebookPageRemovedMap[signalNotebookPageRemovedId] = detail

	return signalNotebookPageRemovedId
}

/*
DisconnectPageRemoved disconnects a callback from the 'page-removed' signal for the Notebook.

The connectionID should be a value returned from a call to ConnectPageRemoved.
*/
func (recv *Notebook) DisconnectPageRemoved(connectionID int) {
	signalNotebookPageRemovedLock.Lock()
	defer signalNotebookPageRemovedLock.Unlock()

	detail, exists := signalNotebookPageRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNotebookPageRemovedMap, connectionID)
}

//export notebook_pageRemovedHandler
func notebook_pageRemovedHandler(_ *C.GObject, c_child *C.GtkWidget, c_page_num C.guint, data C.gpointer) {
	signalNotebookPageRemovedLock.RLock()
	defer signalNotebookPageRemovedLock.RUnlock()

	child := WidgetNewFromC(unsafe.Pointer(c_child))

	pageNum := uint32(c_page_num)

	index := int(uintptr(data))
	callback := signalNotebookPageRemovedMap[index].callback
	callback(child, pageNum)
}

type signalNotebookPageReorderedDetail struct {
	callback  NotebookSignalPageReorderedCallback
	handlerID C.gulong
}

var signalNotebookPageReorderedId int
var signalNotebookPageReorderedMap = make(map[int]signalNotebookPageReorderedDetail)
var signalNotebookPageReorderedLock sync.RWMutex

// NotebookSignalPageReorderedCallback is a callback function for a 'page-reordered' signal emitted from a Notebook.
type NotebookSignalPageReorderedCallback func(child *Widget, pageNum uint32)

/*
ConnectPageReordered connects the callback to the 'page-reordered' signal for the Notebook.

The returned value represents the connection, and may be passed to DisconnectPageReordered to remove it.
*/
func (recv *Notebook) ConnectPageReordered(callback NotebookSignalPageReorderedCallback) int {
	signalNotebookPageReorderedLock.Lock()
	defer signalNotebookPageReorderedLock.Unlock()

	signalNotebookPageReorderedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Notebook_signal_connect_page_reordered(instance, C.gpointer(uintptr(signalNotebookPageReorderedId)))

	detail := signalNotebookPageReorderedDetail{callback, handlerID}
	signalNotebookPageReorderedMap[signalNotebookPageReorderedId] = detail

	return signalNotebookPageReorderedId
}

/*
DisconnectPageReordered disconnects a callback from the 'page-reordered' signal for the Notebook.

The connectionID should be a value returned from a call to ConnectPageReordered.
*/
func (recv *Notebook) DisconnectPageReordered(connectionID int) {
	signalNotebookPageReorderedLock.Lock()
	defer signalNotebookPageReorderedLock.Unlock()

	detail, exists := signalNotebookPageReorderedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNotebookPageReorderedMap, connectionID)
}

//export notebook_pageReorderedHandler
func notebook_pageReorderedHandler(_ *C.GObject, c_child *C.GtkWidget, c_page_num C.guint, data C.gpointer) {
	signalNotebookPageReorderedLock.RLock()
	defer signalNotebookPageReorderedLock.RUnlock()

	child := WidgetNewFromC(unsafe.Pointer(c_child))

	pageNum := uint32(c_page_num)

	index := int(uintptr(data))
	callback := signalNotebookPageReorderedMap[index].callback
	callback(child, pageNum)
}

// Blacklisted : gtk_notebook_get_tab_detachable

// Blacklisted : gtk_notebook_get_tab_reorderable

// Blacklisted : gtk_notebook_set_tab_detachable

// Blacklisted : gtk_notebook_set_tab_reorderable

// Blacklisted : gtk_page_setup_new

// Blacklisted : gtk_page_setup_copy

// Blacklisted : gtk_page_setup_get_bottom_margin

// Blacklisted : gtk_page_setup_get_left_margin

// Blacklisted : gtk_page_setup_get_orientation

// Blacklisted : gtk_page_setup_get_page_height

// Blacklisted : gtk_page_setup_get_page_width

// Blacklisted : gtk_page_setup_get_paper_height

// Blacklisted : gtk_page_setup_get_paper_size

// Blacklisted : gtk_page_setup_get_paper_width

// Blacklisted : gtk_page_setup_get_right_margin

// Blacklisted : gtk_page_setup_get_top_margin

// Blacklisted : gtk_page_setup_set_bottom_margin

// Blacklisted : gtk_page_setup_set_left_margin

// Blacklisted : gtk_page_setup_set_orientation

// Blacklisted : gtk_page_setup_set_paper_size

// Blacklisted : gtk_page_setup_set_paper_size_and_default_margins

// Blacklisted : gtk_page_setup_set_right_margin

// Blacklisted : gtk_page_setup_set_top_margin

// Blacklisted : gtk_print_context_create_pango_context

// Blacklisted : gtk_print_context_create_pango_layout

// Blacklisted : gtk_print_context_get_cairo_context

// Blacklisted : gtk_print_context_get_dpi_x

// Blacklisted : gtk_print_context_get_dpi_y

// Blacklisted : gtk_print_context_get_height

// Blacklisted : gtk_print_context_get_page_setup

// Blacklisted : gtk_print_context_get_pango_fontmap

// Blacklisted : gtk_print_context_get_width

// Blacklisted : gtk_print_context_set_cairo_context

type signalPrintOperationBeginPrintDetail struct {
	callback  PrintOperationSignalBeginPrintCallback
	handlerID C.gulong
}

var signalPrintOperationBeginPrintId int
var signalPrintOperationBeginPrintMap = make(map[int]signalPrintOperationBeginPrintDetail)
var signalPrintOperationBeginPrintLock sync.RWMutex

// PrintOperationSignalBeginPrintCallback is a callback function for a 'begin-print' signal emitted from a PrintOperation.
type PrintOperationSignalBeginPrintCallback func(context *PrintContext)

/*
ConnectBeginPrint connects the callback to the 'begin-print' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectBeginPrint to remove it.
*/
func (recv *PrintOperation) ConnectBeginPrint(callback PrintOperationSignalBeginPrintCallback) int {
	signalPrintOperationBeginPrintLock.Lock()
	defer signalPrintOperationBeginPrintLock.Unlock()

	signalPrintOperationBeginPrintId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_begin_print(instance, C.gpointer(uintptr(signalPrintOperationBeginPrintId)))

	detail := signalPrintOperationBeginPrintDetail{callback, handlerID}
	signalPrintOperationBeginPrintMap[signalPrintOperationBeginPrintId] = detail

	return signalPrintOperationBeginPrintId
}

/*
DisconnectBeginPrint disconnects a callback from the 'begin-print' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectBeginPrint.
*/
func (recv *PrintOperation) DisconnectBeginPrint(connectionID int) {
	signalPrintOperationBeginPrintLock.Lock()
	defer signalPrintOperationBeginPrintLock.Unlock()

	detail, exists := signalPrintOperationBeginPrintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationBeginPrintMap, connectionID)
}

//export printoperation_beginPrintHandler
func printoperation_beginPrintHandler(_ *C.GObject, c_context *C.GtkPrintContext, data C.gpointer) {
	signalPrintOperationBeginPrintLock.RLock()
	defer signalPrintOperationBeginPrintLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationBeginPrintMap[index].callback
	callback(context)
}

type signalPrintOperationCreateCustomWidgetDetail struct {
	callback  PrintOperationSignalCreateCustomWidgetCallback
	handlerID C.gulong
}

var signalPrintOperationCreateCustomWidgetId int
var signalPrintOperationCreateCustomWidgetMap = make(map[int]signalPrintOperationCreateCustomWidgetDetail)
var signalPrintOperationCreateCustomWidgetLock sync.RWMutex

// PrintOperationSignalCreateCustomWidgetCallback is a callback function for a 'create-custom-widget' signal emitted from a PrintOperation.
type PrintOperationSignalCreateCustomWidgetCallback func() gobject.Object

/*
ConnectCreateCustomWidget connects the callback to the 'create-custom-widget' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectCreateCustomWidget to remove it.
*/
func (recv *PrintOperation) ConnectCreateCustomWidget(callback PrintOperationSignalCreateCustomWidgetCallback) int {
	signalPrintOperationCreateCustomWidgetLock.Lock()
	defer signalPrintOperationCreateCustomWidgetLock.Unlock()

	signalPrintOperationCreateCustomWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_create_custom_widget(instance, C.gpointer(uintptr(signalPrintOperationCreateCustomWidgetId)))

	detail := signalPrintOperationCreateCustomWidgetDetail{callback, handlerID}
	signalPrintOperationCreateCustomWidgetMap[signalPrintOperationCreateCustomWidgetId] = detail

	return signalPrintOperationCreateCustomWidgetId
}

/*
DisconnectCreateCustomWidget disconnects a callback from the 'create-custom-widget' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectCreateCustomWidget.
*/
func (recv *PrintOperation) DisconnectCreateCustomWidget(connectionID int) {
	signalPrintOperationCreateCustomWidgetLock.Lock()
	defer signalPrintOperationCreateCustomWidgetLock.Unlock()

	detail, exists := signalPrintOperationCreateCustomWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationCreateCustomWidgetMap, connectionID)
}

//export printoperation_createCustomWidgetHandler
func printoperation_createCustomWidgetHandler(_ *C.GObject, data C.gpointer) *C.GObject {
	signalPrintOperationCreateCustomWidgetLock.RLock()
	defer signalPrintOperationCreateCustomWidgetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPrintOperationCreateCustomWidgetMap[index].callback
	retGo := callback()
	retC :=
		(*C.GObject)(retGo.ToC())
	return retC
}

type signalPrintOperationCustomWidgetApplyDetail struct {
	callback  PrintOperationSignalCustomWidgetApplyCallback
	handlerID C.gulong
}

var signalPrintOperationCustomWidgetApplyId int
var signalPrintOperationCustomWidgetApplyMap = make(map[int]signalPrintOperationCustomWidgetApplyDetail)
var signalPrintOperationCustomWidgetApplyLock sync.RWMutex

// PrintOperationSignalCustomWidgetApplyCallback is a callback function for a 'custom-widget-apply' signal emitted from a PrintOperation.
type PrintOperationSignalCustomWidgetApplyCallback func(widget *Widget)

/*
ConnectCustomWidgetApply connects the callback to the 'custom-widget-apply' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectCustomWidgetApply to remove it.
*/
func (recv *PrintOperation) ConnectCustomWidgetApply(callback PrintOperationSignalCustomWidgetApplyCallback) int {
	signalPrintOperationCustomWidgetApplyLock.Lock()
	defer signalPrintOperationCustomWidgetApplyLock.Unlock()

	signalPrintOperationCustomWidgetApplyId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_custom_widget_apply(instance, C.gpointer(uintptr(signalPrintOperationCustomWidgetApplyId)))

	detail := signalPrintOperationCustomWidgetApplyDetail{callback, handlerID}
	signalPrintOperationCustomWidgetApplyMap[signalPrintOperationCustomWidgetApplyId] = detail

	return signalPrintOperationCustomWidgetApplyId
}

/*
DisconnectCustomWidgetApply disconnects a callback from the 'custom-widget-apply' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectCustomWidgetApply.
*/
func (recv *PrintOperation) DisconnectCustomWidgetApply(connectionID int) {
	signalPrintOperationCustomWidgetApplyLock.Lock()
	defer signalPrintOperationCustomWidgetApplyLock.Unlock()

	detail, exists := signalPrintOperationCustomWidgetApplyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationCustomWidgetApplyMap, connectionID)
}

//export printoperation_customWidgetApplyHandler
func printoperation_customWidgetApplyHandler(_ *C.GObject, c_widget *C.GtkWidget, data C.gpointer) {
	signalPrintOperationCustomWidgetApplyLock.RLock()
	defer signalPrintOperationCustomWidgetApplyLock.RUnlock()

	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	index := int(uintptr(data))
	callback := signalPrintOperationCustomWidgetApplyMap[index].callback
	callback(widget)
}

type signalPrintOperationDoneDetail struct {
	callback  PrintOperationSignalDoneCallback
	handlerID C.gulong
}

var signalPrintOperationDoneId int
var signalPrintOperationDoneMap = make(map[int]signalPrintOperationDoneDetail)
var signalPrintOperationDoneLock sync.RWMutex

// PrintOperationSignalDoneCallback is a callback function for a 'done' signal emitted from a PrintOperation.
type PrintOperationSignalDoneCallback func(result PrintOperationResult)

/*
ConnectDone connects the callback to the 'done' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectDone to remove it.
*/
func (recv *PrintOperation) ConnectDone(callback PrintOperationSignalDoneCallback) int {
	signalPrintOperationDoneLock.Lock()
	defer signalPrintOperationDoneLock.Unlock()

	signalPrintOperationDoneId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_done(instance, C.gpointer(uintptr(signalPrintOperationDoneId)))

	detail := signalPrintOperationDoneDetail{callback, handlerID}
	signalPrintOperationDoneMap[signalPrintOperationDoneId] = detail

	return signalPrintOperationDoneId
}

/*
DisconnectDone disconnects a callback from the 'done' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectDone.
*/
func (recv *PrintOperation) DisconnectDone(connectionID int) {
	signalPrintOperationDoneLock.Lock()
	defer signalPrintOperationDoneLock.Unlock()

	detail, exists := signalPrintOperationDoneMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationDoneMap, connectionID)
}

//export printoperation_doneHandler
func printoperation_doneHandler(_ *C.GObject, c_result C.GtkPrintOperationResult, data C.gpointer) {
	signalPrintOperationDoneLock.RLock()
	defer signalPrintOperationDoneLock.RUnlock()

	result := PrintOperationResult(c_result)

	index := int(uintptr(data))
	callback := signalPrintOperationDoneMap[index].callback
	callback(result)
}

type signalPrintOperationDrawPageDetail struct {
	callback  PrintOperationSignalDrawPageCallback
	handlerID C.gulong
}

var signalPrintOperationDrawPageId int
var signalPrintOperationDrawPageMap = make(map[int]signalPrintOperationDrawPageDetail)
var signalPrintOperationDrawPageLock sync.RWMutex

// PrintOperationSignalDrawPageCallback is a callback function for a 'draw-page' signal emitted from a PrintOperation.
type PrintOperationSignalDrawPageCallback func(context *PrintContext, pageNr int32)

/*
ConnectDrawPage connects the callback to the 'draw-page' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectDrawPage to remove it.
*/
func (recv *PrintOperation) ConnectDrawPage(callback PrintOperationSignalDrawPageCallback) int {
	signalPrintOperationDrawPageLock.Lock()
	defer signalPrintOperationDrawPageLock.Unlock()

	signalPrintOperationDrawPageId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_draw_page(instance, C.gpointer(uintptr(signalPrintOperationDrawPageId)))

	detail := signalPrintOperationDrawPageDetail{callback, handlerID}
	signalPrintOperationDrawPageMap[signalPrintOperationDrawPageId] = detail

	return signalPrintOperationDrawPageId
}

/*
DisconnectDrawPage disconnects a callback from the 'draw-page' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectDrawPage.
*/
func (recv *PrintOperation) DisconnectDrawPage(connectionID int) {
	signalPrintOperationDrawPageLock.Lock()
	defer signalPrintOperationDrawPageLock.Unlock()

	detail, exists := signalPrintOperationDrawPageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationDrawPageMap, connectionID)
}

//export printoperation_drawPageHandler
func printoperation_drawPageHandler(_ *C.GObject, c_context *C.GtkPrintContext, c_page_nr C.gint, data C.gpointer) {
	signalPrintOperationDrawPageLock.RLock()
	defer signalPrintOperationDrawPageLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	pageNr := int32(c_page_nr)

	index := int(uintptr(data))
	callback := signalPrintOperationDrawPageMap[index].callback
	callback(context, pageNr)
}

type signalPrintOperationEndPrintDetail struct {
	callback  PrintOperationSignalEndPrintCallback
	handlerID C.gulong
}

var signalPrintOperationEndPrintId int
var signalPrintOperationEndPrintMap = make(map[int]signalPrintOperationEndPrintDetail)
var signalPrintOperationEndPrintLock sync.RWMutex

// PrintOperationSignalEndPrintCallback is a callback function for a 'end-print' signal emitted from a PrintOperation.
type PrintOperationSignalEndPrintCallback func(context *PrintContext)

/*
ConnectEndPrint connects the callback to the 'end-print' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectEndPrint to remove it.
*/
func (recv *PrintOperation) ConnectEndPrint(callback PrintOperationSignalEndPrintCallback) int {
	signalPrintOperationEndPrintLock.Lock()
	defer signalPrintOperationEndPrintLock.Unlock()

	signalPrintOperationEndPrintId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_end_print(instance, C.gpointer(uintptr(signalPrintOperationEndPrintId)))

	detail := signalPrintOperationEndPrintDetail{callback, handlerID}
	signalPrintOperationEndPrintMap[signalPrintOperationEndPrintId] = detail

	return signalPrintOperationEndPrintId
}

/*
DisconnectEndPrint disconnects a callback from the 'end-print' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectEndPrint.
*/
func (recv *PrintOperation) DisconnectEndPrint(connectionID int) {
	signalPrintOperationEndPrintLock.Lock()
	defer signalPrintOperationEndPrintLock.Unlock()

	detail, exists := signalPrintOperationEndPrintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationEndPrintMap, connectionID)
}

//export printoperation_endPrintHandler
func printoperation_endPrintHandler(_ *C.GObject, c_context *C.GtkPrintContext, data C.gpointer) {
	signalPrintOperationEndPrintLock.RLock()
	defer signalPrintOperationEndPrintLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationEndPrintMap[index].callback
	callback(context)
}

type signalPrintOperationPaginateDetail struct {
	callback  PrintOperationSignalPaginateCallback
	handlerID C.gulong
}

var signalPrintOperationPaginateId int
var signalPrintOperationPaginateMap = make(map[int]signalPrintOperationPaginateDetail)
var signalPrintOperationPaginateLock sync.RWMutex

// PrintOperationSignalPaginateCallback is a callback function for a 'paginate' signal emitted from a PrintOperation.
type PrintOperationSignalPaginateCallback func(context *PrintContext) bool

/*
ConnectPaginate connects the callback to the 'paginate' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectPaginate to remove it.
*/
func (recv *PrintOperation) ConnectPaginate(callback PrintOperationSignalPaginateCallback) int {
	signalPrintOperationPaginateLock.Lock()
	defer signalPrintOperationPaginateLock.Unlock()

	signalPrintOperationPaginateId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_paginate(instance, C.gpointer(uintptr(signalPrintOperationPaginateId)))

	detail := signalPrintOperationPaginateDetail{callback, handlerID}
	signalPrintOperationPaginateMap[signalPrintOperationPaginateId] = detail

	return signalPrintOperationPaginateId
}

/*
DisconnectPaginate disconnects a callback from the 'paginate' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectPaginate.
*/
func (recv *PrintOperation) DisconnectPaginate(connectionID int) {
	signalPrintOperationPaginateLock.Lock()
	defer signalPrintOperationPaginateLock.Unlock()

	detail, exists := signalPrintOperationPaginateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPaginateMap, connectionID)
}

//export printoperation_paginateHandler
func printoperation_paginateHandler(_ *C.GObject, c_context *C.GtkPrintContext, data C.gpointer) C.gboolean {
	signalPrintOperationPaginateLock.RLock()
	defer signalPrintOperationPaginateLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationPaginateMap[index].callback
	retGo := callback(context)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalPrintOperationPreviewDetail struct {
	callback  PrintOperationSignalPreviewCallback
	handlerID C.gulong
}

var signalPrintOperationPreviewId int
var signalPrintOperationPreviewMap = make(map[int]signalPrintOperationPreviewDetail)
var signalPrintOperationPreviewLock sync.RWMutex

// PrintOperationSignalPreviewCallback is a callback function for a 'preview' signal emitted from a PrintOperation.
type PrintOperationSignalPreviewCallback func(preview *PrintOperationPreview, context *PrintContext, parent *Window) bool

/*
ConnectPreview connects the callback to the 'preview' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectPreview to remove it.
*/
func (recv *PrintOperation) ConnectPreview(callback PrintOperationSignalPreviewCallback) int {
	signalPrintOperationPreviewLock.Lock()
	defer signalPrintOperationPreviewLock.Unlock()

	signalPrintOperationPreviewId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_preview(instance, C.gpointer(uintptr(signalPrintOperationPreviewId)))

	detail := signalPrintOperationPreviewDetail{callback, handlerID}
	signalPrintOperationPreviewMap[signalPrintOperationPreviewId] = detail

	return signalPrintOperationPreviewId
}

/*
DisconnectPreview disconnects a callback from the 'preview' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectPreview.
*/
func (recv *PrintOperation) DisconnectPreview(connectionID int) {
	signalPrintOperationPreviewLock.Lock()
	defer signalPrintOperationPreviewLock.Unlock()

	detail, exists := signalPrintOperationPreviewMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPreviewMap, connectionID)
}

//export printoperation_previewHandler
func printoperation_previewHandler(_ *C.GObject, c_preview *C.GtkPrintOperationPreview, c_context *C.GtkPrintContext, c_parent *C.GtkWindow, data C.gpointer) C.gboolean {
	signalPrintOperationPreviewLock.RLock()
	defer signalPrintOperationPreviewLock.RUnlock()

	preview := PrintOperationPreviewNewFromC(unsafe.Pointer(c_preview))

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	parent := WindowNewFromC(unsafe.Pointer(c_parent))

	index := int(uintptr(data))
	callback := signalPrintOperationPreviewMap[index].callback
	retGo := callback(preview, context, parent)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalPrintOperationRequestPageSetupDetail struct {
	callback  PrintOperationSignalRequestPageSetupCallback
	handlerID C.gulong
}

var signalPrintOperationRequestPageSetupId int
var signalPrintOperationRequestPageSetupMap = make(map[int]signalPrintOperationRequestPageSetupDetail)
var signalPrintOperationRequestPageSetupLock sync.RWMutex

// PrintOperationSignalRequestPageSetupCallback is a callback function for a 'request-page-setup' signal emitted from a PrintOperation.
type PrintOperationSignalRequestPageSetupCallback func(context *PrintContext, pageNr int32, setup *PageSetup)

/*
ConnectRequestPageSetup connects the callback to the 'request-page-setup' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectRequestPageSetup to remove it.
*/
func (recv *PrintOperation) ConnectRequestPageSetup(callback PrintOperationSignalRequestPageSetupCallback) int {
	signalPrintOperationRequestPageSetupLock.Lock()
	defer signalPrintOperationRequestPageSetupLock.Unlock()

	signalPrintOperationRequestPageSetupId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_request_page_setup(instance, C.gpointer(uintptr(signalPrintOperationRequestPageSetupId)))

	detail := signalPrintOperationRequestPageSetupDetail{callback, handlerID}
	signalPrintOperationRequestPageSetupMap[signalPrintOperationRequestPageSetupId] = detail

	return signalPrintOperationRequestPageSetupId
}

/*
DisconnectRequestPageSetup disconnects a callback from the 'request-page-setup' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectRequestPageSetup.
*/
func (recv *PrintOperation) DisconnectRequestPageSetup(connectionID int) {
	signalPrintOperationRequestPageSetupLock.Lock()
	defer signalPrintOperationRequestPageSetupLock.Unlock()

	detail, exists := signalPrintOperationRequestPageSetupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationRequestPageSetupMap, connectionID)
}

//export printoperation_requestPageSetupHandler
func printoperation_requestPageSetupHandler(_ *C.GObject, c_context *C.GtkPrintContext, c_page_nr C.gint, c_setup *C.GtkPageSetup, data C.gpointer) {
	signalPrintOperationRequestPageSetupLock.RLock()
	defer signalPrintOperationRequestPageSetupLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	pageNr := int32(c_page_nr)

	setup := PageSetupNewFromC(unsafe.Pointer(c_setup))

	index := int(uintptr(data))
	callback := signalPrintOperationRequestPageSetupMap[index].callback
	callback(context, pageNr, setup)
}

type signalPrintOperationStatusChangedDetail struct {
	callback  PrintOperationSignalStatusChangedCallback
	handlerID C.gulong
}

var signalPrintOperationStatusChangedId int
var signalPrintOperationStatusChangedMap = make(map[int]signalPrintOperationStatusChangedDetail)
var signalPrintOperationStatusChangedLock sync.RWMutex

// PrintOperationSignalStatusChangedCallback is a callback function for a 'status-changed' signal emitted from a PrintOperation.
type PrintOperationSignalStatusChangedCallback func()

/*
ConnectStatusChanged connects the callback to the 'status-changed' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectStatusChanged to remove it.
*/
func (recv *PrintOperation) ConnectStatusChanged(callback PrintOperationSignalStatusChangedCallback) int {
	signalPrintOperationStatusChangedLock.Lock()
	defer signalPrintOperationStatusChangedLock.Unlock()

	signalPrintOperationStatusChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_status_changed(instance, C.gpointer(uintptr(signalPrintOperationStatusChangedId)))

	detail := signalPrintOperationStatusChangedDetail{callback, handlerID}
	signalPrintOperationStatusChangedMap[signalPrintOperationStatusChangedId] = detail

	return signalPrintOperationStatusChangedId
}

/*
DisconnectStatusChanged disconnects a callback from the 'status-changed' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectStatusChanged.
*/
func (recv *PrintOperation) DisconnectStatusChanged(connectionID int) {
	signalPrintOperationStatusChangedLock.Lock()
	defer signalPrintOperationStatusChangedLock.Unlock()

	detail, exists := signalPrintOperationStatusChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationStatusChangedMap, connectionID)
}

//export printoperation_statusChangedHandler
func printoperation_statusChangedHandler(_ *C.GObject, data C.gpointer) {
	signalPrintOperationStatusChangedLock.RLock()
	defer signalPrintOperationStatusChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPrintOperationStatusChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_print_operation_new

// Blacklisted : gtk_print_operation_cancel

// Blacklisted : gtk_print_operation_get_default_page_setup

// Blacklisted : gtk_print_operation_get_error

// Blacklisted : gtk_print_operation_get_print_settings

// Blacklisted : gtk_print_operation_get_status

// Blacklisted : gtk_print_operation_get_status_string

// Blacklisted : gtk_print_operation_is_finished

// Blacklisted : gtk_print_operation_run

// Blacklisted : gtk_print_operation_set_allow_async

// Blacklisted : gtk_print_operation_set_current_page

// Blacklisted : gtk_print_operation_set_custom_tab_label

// Blacklisted : gtk_print_operation_set_default_page_setup

// Blacklisted : gtk_print_operation_set_export_filename

// Blacklisted : gtk_print_operation_set_job_name

// Blacklisted : gtk_print_operation_set_n_pages

// Blacklisted : gtk_print_operation_set_print_settings

// Blacklisted : gtk_print_operation_set_show_progress

// Blacklisted : gtk_print_operation_set_track_print_status

// Blacklisted : gtk_print_operation_set_unit

// Blacklisted : gtk_print_operation_set_use_full_page

// Blacklisted : gtk_print_settings_new

// Blacklisted : gtk_print_settings_copy

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc (GtkPrintSettingsFunc) for param func

// Blacklisted : gtk_print_settings_get

// Blacklisted : gtk_print_settings_get_bool

// Blacklisted : gtk_print_settings_get_collate

// Blacklisted : gtk_print_settings_get_default_source

// Blacklisted : gtk_print_settings_get_dither

// Blacklisted : gtk_print_settings_get_double

// Blacklisted : gtk_print_settings_get_double_with_default

// Blacklisted : gtk_print_settings_get_duplex

// Blacklisted : gtk_print_settings_get_finishings

// Blacklisted : gtk_print_settings_get_int

// Blacklisted : gtk_print_settings_get_int_with_default

// Blacklisted : gtk_print_settings_get_length

// Blacklisted : gtk_print_settings_get_media_type

// Blacklisted : gtk_print_settings_get_n_copies

// Blacklisted : gtk_print_settings_get_number_up

// Blacklisted : gtk_print_settings_get_orientation

// Blacklisted : gtk_print_settings_get_output_bin

// Unsupported : gtk_print_settings_get_page_ranges : array return type :

// Blacklisted : gtk_print_settings_get_page_set

// Blacklisted : gtk_print_settings_get_paper_height

// Blacklisted : gtk_print_settings_get_paper_size

// Blacklisted : gtk_print_settings_get_paper_width

// Blacklisted : gtk_print_settings_get_print_pages

// Blacklisted : gtk_print_settings_get_printer

// Blacklisted : gtk_print_settings_get_quality

// Blacklisted : gtk_print_settings_get_resolution

// Blacklisted : gtk_print_settings_get_reverse

// Blacklisted : gtk_print_settings_get_scale

// Blacklisted : gtk_print_settings_get_use_color

// Blacklisted : gtk_print_settings_has_key

// Blacklisted : gtk_print_settings_set

// Blacklisted : gtk_print_settings_set_bool

// Blacklisted : gtk_print_settings_set_collate

// Blacklisted : gtk_print_settings_set_default_source

// Blacklisted : gtk_print_settings_set_dither

// Blacklisted : gtk_print_settings_set_double

// Blacklisted : gtk_print_settings_set_duplex

// Blacklisted : gtk_print_settings_set_finishings

// Blacklisted : gtk_print_settings_set_int

// Blacklisted : gtk_print_settings_set_length

// Blacklisted : gtk_print_settings_set_media_type

// Blacklisted : gtk_print_settings_set_n_copies

// Blacklisted : gtk_print_settings_set_number_up

// Blacklisted : gtk_print_settings_set_orientation

// Blacklisted : gtk_print_settings_set_output_bin

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges :

// Blacklisted : gtk_print_settings_set_page_set

// Blacklisted : gtk_print_settings_set_paper_height

// Blacklisted : gtk_print_settings_set_paper_size

// Blacklisted : gtk_print_settings_set_paper_width

// Blacklisted : gtk_print_settings_set_print_pages

// Blacklisted : gtk_print_settings_set_printer

// Blacklisted : gtk_print_settings_set_quality

// Blacklisted : gtk_print_settings_set_resolution

// Blacklisted : gtk_print_settings_set_reverse

// Blacklisted : gtk_print_settings_set_scale

// Blacklisted : gtk_print_settings_set_use_color

// Blacklisted : gtk_print_settings_unset

// Blacklisted : gtk_radio_action_set_current_value

// Blacklisted : gtk_range_get_lower_stepper_sensitivity

// Blacklisted : gtk_range_get_upper_stepper_sensitivity

// Blacklisted : gtk_range_set_lower_stepper_sensitivity

// Blacklisted : gtk_range_set_upper_stepper_sensitivity

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Blacklisted : gtk_recent_chooser_menu_new

// Blacklisted : gtk_recent_chooser_menu_new_for_manager

// Blacklisted : gtk_recent_chooser_menu_get_show_numbers

// Blacklisted : gtk_recent_chooser_menu_set_show_numbers

// Blacklisted : gtk_recent_chooser_widget_new

// Blacklisted : gtk_recent_chooser_widget_new_for_manager

// Blacklisted : gtk_recent_filter_new

// Blacklisted : gtk_recent_filter_add_age

// Blacklisted : gtk_recent_filter_add_application

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc (GtkRecentFilterFunc) for param func

// Blacklisted : gtk_recent_filter_add_group

// Blacklisted : gtk_recent_filter_add_mime_type

// Blacklisted : gtk_recent_filter_add_pattern

// Blacklisted : gtk_recent_filter_add_pixbuf_formats

// Blacklisted : gtk_recent_filter_filter

// Blacklisted : gtk_recent_filter_get_name

// Blacklisted : gtk_recent_filter_get_needed

// Blacklisted : gtk_recent_filter_set_name

// RecentManager is a wrapper around the C record GtkRecentManager.
type RecentManager struct {
	native *C.GtkRecentManager
	// Private : parent_instance
	// Private : priv
}

func RecentManagerNewFromC(u unsafe.Pointer) *RecentManager {
	c := (*C.GtkRecentManager)(u)
	if c == nil {
		return nil
	}

	g := &RecentManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RecentManager with another RecentManager, and returns true if they represent the same GObject.
func (recv *RecentManager) Equals(other *RecentManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *RecentManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to RecentManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentManager.
func CastToRecentManager(object *gobject.Object) *RecentManager {
	return RecentManagerNewFromC(object.ToC())
}

type signalRecentManagerChangedDetail struct {
	callback  RecentManagerSignalChangedCallback
	handlerID C.gulong
}

var signalRecentManagerChangedId int
var signalRecentManagerChangedMap = make(map[int]signalRecentManagerChangedDetail)
var signalRecentManagerChangedLock sync.RWMutex

// RecentManagerSignalChangedCallback is a callback function for a 'changed' signal emitted from a RecentManager.
type RecentManagerSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the RecentManager.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *RecentManager) ConnectChanged(callback RecentManagerSignalChangedCallback) int {
	signalRecentManagerChangedLock.Lock()
	defer signalRecentManagerChangedLock.Unlock()

	signalRecentManagerChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RecentManager_signal_connect_changed(instance, C.gpointer(uintptr(signalRecentManagerChangedId)))

	detail := signalRecentManagerChangedDetail{callback, handlerID}
	signalRecentManagerChangedMap[signalRecentManagerChangedId] = detail

	return signalRecentManagerChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the RecentManager.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *RecentManager) DisconnectChanged(connectionID int) {
	signalRecentManagerChangedLock.Lock()
	defer signalRecentManagerChangedLock.Unlock()

	detail, exists := signalRecentManagerChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRecentManagerChangedMap, connectionID)
}

//export recentmanager_changedHandler
func recentmanager_changedHandler(_ *C.GObject, data C.gpointer) {
	signalRecentManagerChangedLock.RLock()
	defer signalRecentManagerChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalRecentManagerChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_recent_manager_new

// Blacklisted : gtk_recent_manager_get_default

// Blacklisted : gtk_recent_manager_add_full

// Blacklisted : gtk_recent_manager_add_item

// Blacklisted : gtk_recent_manager_get_items

// Blacklisted : gtk_recent_manager_has_item

// Blacklisted : gtk_recent_manager_lookup_item

// Blacklisted : gtk_recent_manager_move_item

// Blacklisted : gtk_recent_manager_purge_items

// Blacklisted : gtk_recent_manager_remove_item

// Blacklisted : gtk_scrolled_window_unset_placement

// Blacklisted : gtk_size_group_get_widgets

type signalSpinButtonWrappedDetail struct {
	callback  SpinButtonSignalWrappedCallback
	handlerID C.gulong
}

var signalSpinButtonWrappedId int
var signalSpinButtonWrappedMap = make(map[int]signalSpinButtonWrappedDetail)
var signalSpinButtonWrappedLock sync.RWMutex

// SpinButtonSignalWrappedCallback is a callback function for a 'wrapped' signal emitted from a SpinButton.
type SpinButtonSignalWrappedCallback func()

/*
ConnectWrapped connects the callback to the 'wrapped' signal for the SpinButton.

The returned value represents the connection, and may be passed to DisconnectWrapped to remove it.
*/
func (recv *SpinButton) ConnectWrapped(callback SpinButtonSignalWrappedCallback) int {
	signalSpinButtonWrappedLock.Lock()
	defer signalSpinButtonWrappedLock.Unlock()

	signalSpinButtonWrappedId++
	instance := C.gpointer(recv.native)
	handlerID := C.SpinButton_signal_connect_wrapped(instance, C.gpointer(uintptr(signalSpinButtonWrappedId)))

	detail := signalSpinButtonWrappedDetail{callback, handlerID}
	signalSpinButtonWrappedMap[signalSpinButtonWrappedId] = detail

	return signalSpinButtonWrappedId
}

/*
DisconnectWrapped disconnects a callback from the 'wrapped' signal for the SpinButton.

The connectionID should be a value returned from a call to ConnectWrapped.
*/
func (recv *SpinButton) DisconnectWrapped(connectionID int) {
	signalSpinButtonWrappedLock.Lock()
	defer signalSpinButtonWrappedLock.Unlock()

	detail, exists := signalSpinButtonWrappedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSpinButtonWrappedMap, connectionID)
}

//export spinbutton_wrappedHandler
func spinbutton_wrappedHandler(_ *C.GObject, data C.gpointer) {
	signalSpinButtonWrappedLock.RLock()
	defer signalSpinButtonWrappedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalSpinButtonWrappedMap[index].callback
	callback()
}

type signalStatusIconActivateDetail struct {
	callback  StatusIconSignalActivateCallback
	handlerID C.gulong
}

var signalStatusIconActivateId int
var signalStatusIconActivateMap = make(map[int]signalStatusIconActivateDetail)
var signalStatusIconActivateLock sync.RWMutex

// StatusIconSignalActivateCallback is a callback function for a 'activate' signal emitted from a StatusIcon.
type StatusIconSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *StatusIcon) ConnectActivate(callback StatusIconSignalActivateCallback) int {
	signalStatusIconActivateLock.Lock()
	defer signalStatusIconActivateLock.Unlock()

	signalStatusIconActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_activate(instance, C.gpointer(uintptr(signalStatusIconActivateId)))

	detail := signalStatusIconActivateDetail{callback, handlerID}
	signalStatusIconActivateMap[signalStatusIconActivateId] = detail

	return signalStatusIconActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *StatusIcon) DisconnectActivate(connectionID int) {
	signalStatusIconActivateLock.Lock()
	defer signalStatusIconActivateLock.Unlock()

	detail, exists := signalStatusIconActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconActivateMap, connectionID)
}

//export statusicon_activateHandler
func statusicon_activateHandler(_ *C.GObject, data C.gpointer) {
	signalStatusIconActivateLock.RLock()
	defer signalStatusIconActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStatusIconActivateMap[index].callback
	callback()
}

type signalStatusIconPopupMenuDetail struct {
	callback  StatusIconSignalPopupMenuCallback
	handlerID C.gulong
}

var signalStatusIconPopupMenuId int
var signalStatusIconPopupMenuMap = make(map[int]signalStatusIconPopupMenuDetail)
var signalStatusIconPopupMenuLock sync.RWMutex

// StatusIconSignalPopupMenuCallback is a callback function for a 'popup-menu' signal emitted from a StatusIcon.
type StatusIconSignalPopupMenuCallback func(button uint32, activateTime uint32)

/*
ConnectPopupMenu connects the callback to the 'popup-menu' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectPopupMenu to remove it.
*/
func (recv *StatusIcon) ConnectPopupMenu(callback StatusIconSignalPopupMenuCallback) int {
	signalStatusIconPopupMenuLock.Lock()
	defer signalStatusIconPopupMenuLock.Unlock()

	signalStatusIconPopupMenuId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_popup_menu(instance, C.gpointer(uintptr(signalStatusIconPopupMenuId)))

	detail := signalStatusIconPopupMenuDetail{callback, handlerID}
	signalStatusIconPopupMenuMap[signalStatusIconPopupMenuId] = detail

	return signalStatusIconPopupMenuId
}

/*
DisconnectPopupMenu disconnects a callback from the 'popup-menu' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectPopupMenu.
*/
func (recv *StatusIcon) DisconnectPopupMenu(connectionID int) {
	signalStatusIconPopupMenuLock.Lock()
	defer signalStatusIconPopupMenuLock.Unlock()

	detail, exists := signalStatusIconPopupMenuMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconPopupMenuMap, connectionID)
}

//export statusicon_popupMenuHandler
func statusicon_popupMenuHandler(_ *C.GObject, c_button C.guint, c_activate_time C.guint, data C.gpointer) {
	signalStatusIconPopupMenuLock.RLock()
	defer signalStatusIconPopupMenuLock.RUnlock()

	button := uint32(c_button)

	activateTime := uint32(c_activate_time)

	index := int(uintptr(data))
	callback := signalStatusIconPopupMenuMap[index].callback
	callback(button, activateTime)
}

type signalStatusIconSizeChangedDetail struct {
	callback  StatusIconSignalSizeChangedCallback
	handlerID C.gulong
}

var signalStatusIconSizeChangedId int
var signalStatusIconSizeChangedMap = make(map[int]signalStatusIconSizeChangedDetail)
var signalStatusIconSizeChangedLock sync.RWMutex

// StatusIconSignalSizeChangedCallback is a callback function for a 'size-changed' signal emitted from a StatusIcon.
type StatusIconSignalSizeChangedCallback func(size int32) bool

/*
ConnectSizeChanged connects the callback to the 'size-changed' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectSizeChanged to remove it.
*/
func (recv *StatusIcon) ConnectSizeChanged(callback StatusIconSignalSizeChangedCallback) int {
	signalStatusIconSizeChangedLock.Lock()
	defer signalStatusIconSizeChangedLock.Unlock()

	signalStatusIconSizeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_size_changed(instance, C.gpointer(uintptr(signalStatusIconSizeChangedId)))

	detail := signalStatusIconSizeChangedDetail{callback, handlerID}
	signalStatusIconSizeChangedMap[signalStatusIconSizeChangedId] = detail

	return signalStatusIconSizeChangedId
}

/*
DisconnectSizeChanged disconnects a callback from the 'size-changed' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectSizeChanged.
*/
func (recv *StatusIcon) DisconnectSizeChanged(connectionID int) {
	signalStatusIconSizeChangedLock.Lock()
	defer signalStatusIconSizeChangedLock.Unlock()

	detail, exists := signalStatusIconSizeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconSizeChangedMap, connectionID)
}

//export statusicon_sizeChangedHandler
func statusicon_sizeChangedHandler(_ *C.GObject, c_size C.gint, data C.gpointer) C.gboolean {
	signalStatusIconSizeChangedLock.RLock()
	defer signalStatusIconSizeChangedLock.RUnlock()

	size := int32(c_size)

	index := int(uintptr(data))
	callback := signalStatusIconSizeChangedMap[index].callback
	retGo := callback(size)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_status_icon_new

// Blacklisted : gtk_status_icon_new_from_file

// Blacklisted : gtk_status_icon_new_from_icon_name

// Blacklisted : gtk_status_icon_new_from_pixbuf

// Blacklisted : gtk_status_icon_new_from_stock

// Blacklisted : gtk_status_icon_position_menu

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter orientation : GtkOrientation* with indirection level of 1

// Blacklisted : gtk_status_icon_get_icon_name

// Blacklisted : gtk_status_icon_get_pixbuf

// Blacklisted : gtk_status_icon_get_size

// Blacklisted : gtk_status_icon_get_stock

// Blacklisted : gtk_status_icon_get_storage_type

// Blacklisted : gtk_status_icon_get_visible

// Blacklisted : gtk_status_icon_is_embedded

// Blacklisted : gtk_status_icon_set_from_file

// Blacklisted : gtk_status_icon_set_from_icon_name

// Blacklisted : gtk_status_icon_set_from_pixbuf

// Blacklisted : gtk_status_icon_set_from_stock

// Blacklisted : gtk_status_icon_set_visible

// Blacklisted : gtk_style_lookup_color

// Blacklisted : gtk_text_buffer_deserialize

// Blacklisted : gtk_text_buffer_deserialize_get_can_create_tags

// Blacklisted : gtk_text_buffer_deserialize_set_can_create_tags

// Blacklisted : gtk_text_buffer_get_copy_target_list

// Unsupported : gtk_text_buffer_get_deserialize_formats : array return type :

// Blacklisted : gtk_text_buffer_get_has_selection

// Blacklisted : gtk_text_buffer_get_paste_target_list

// Unsupported : gtk_text_buffer_get_serialize_formats : array return type :

// Unsupported : gtk_text_buffer_register_deserialize_format : unsupported parameter function : no type generator for TextBufferDeserializeFunc (GtkTextBufferDeserializeFunc) for param function

// Blacklisted : gtk_text_buffer_register_deserialize_tagset

// Unsupported : gtk_text_buffer_register_serialize_format : unsupported parameter function : no type generator for TextBufferSerializeFunc (GtkTextBufferSerializeFunc) for param function

// Blacklisted : gtk_text_buffer_register_serialize_tagset

// Unsupported : gtk_text_buffer_serialize : array return type :

// Blacklisted : gtk_text_buffer_unregister_deserialize_format

// Blacklisted : gtk_text_buffer_unregister_serialize_format

// Unsupported : gtk_tree_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_insert_with_valuesv : unsupported parameter values :

// Blacklisted : gtk_tree_view_get_enable_tree_lines

// Blacklisted : gtk_tree_view_get_grid_lines

// Blacklisted : gtk_tree_view_get_headers_clickable

// Blacklisted : gtk_tree_view_get_rubber_banding

// Blacklisted : gtk_tree_view_get_search_entry

// Unsupported : gtk_tree_view_get_search_position_func : no return generator

// Blacklisted : gtk_tree_view_set_enable_tree_lines

// Blacklisted : gtk_tree_view_set_grid_lines

// Blacklisted : gtk_tree_view_set_rubber_banding

// Blacklisted : gtk_tree_view_set_search_entry

// Unsupported : gtk_tree_view_set_search_position_func : unsupported parameter func : no type generator for TreeViewSearchPositionFunc (GtkTreeViewSearchPositionFunc) for param func

// Blacklisted : gtk_drag_dest_get_track_motion

// Blacklisted : gtk_drag_dest_set_track_motion

// Blacklisted : gtk_widget_is_composited

// Blacklisted : gtk_window_get_deletable

// Blacklisted : gtk_window_get_group

// Blacklisted : gtk_window_set_deletable
