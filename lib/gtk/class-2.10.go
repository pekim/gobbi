// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

	void printoperation_statusChangedHandler(GObject *, gpointer);

	static gulong PrintOperation_signal_connect_status_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "status-changed", G_CALLBACK(printoperation_statusChangedHandler), data);
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
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

type signalAssistantApplyDetail struct {
	callback  AssistantSignalApplyCallback
	handlerID C.gulong
}

var signalAssistantApplyId int
var signalAssistantApplyMap = make(map[int]signalAssistantApplyDetail)
var signalAssistantApplyLock sync.Mutex

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
var signalAssistantCancelLock sync.Mutex

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
var signalAssistantCloseLock sync.Mutex

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
var signalAssistantPrepareLock sync.Mutex

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
	page := WidgetNewFromC(unsafe.Pointer(c_page))

	index := int(uintptr(data))
	callback := signalAssistantPrepareMap[index].callback
	callback(page)
}

// AssistantNew is a wrapper around the C function gtk_assistant_new.
func AssistantNew() *Assistant {
	retC := C.gtk_assistant_new()
	retGo := AssistantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddActionWidget is a wrapper around the C function gtk_assistant_add_action_widget.
func (recv *Assistant) AddActionWidget(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_assistant_add_action_widget((*C.GtkAssistant)(recv.native), c_child)

	return
}

// AppendPage is a wrapper around the C function gtk_assistant_append_page.
func (recv *Assistant) AppendPage(page *Widget) int32 {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_append_page((*C.GtkAssistant)(recv.native), c_page)
	retGo := (int32)(retC)

	return retGo
}

// GetCurrentPage is a wrapper around the C function gtk_assistant_get_current_page.
func (recv *Assistant) GetCurrentPage() int32 {
	retC := C.gtk_assistant_get_current_page((*C.GtkAssistant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNPages is a wrapper around the C function gtk_assistant_get_n_pages.
func (recv *Assistant) GetNPages() int32 {
	retC := C.gtk_assistant_get_n_pages((*C.GtkAssistant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNthPage is a wrapper around the C function gtk_assistant_get_nth_page.
func (recv *Assistant) GetNthPage(pageNum int32) *Widget {
	c_page_num := (C.gint)(pageNum)

	retC := C.gtk_assistant_get_nth_page((*C.GtkAssistant)(recv.native), c_page_num)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPageComplete is a wrapper around the C function gtk_assistant_get_page_complete.
func (recv *Assistant) GetPageComplete(page *Widget) bool {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_get_page_complete((*C.GtkAssistant)(recv.native), c_page)
	retGo := retC == C.TRUE

	return retGo
}

// GetPageHeaderImage is a wrapper around the C function gtk_assistant_get_page_header_image.
func (recv *Assistant) GetPageHeaderImage(page *Widget) *gdkpixbuf.Pixbuf {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_get_page_header_image((*C.GtkAssistant)(recv.native), c_page)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPageSideImage is a wrapper around the C function gtk_assistant_get_page_side_image.
func (recv *Assistant) GetPageSideImage(page *Widget) *gdkpixbuf.Pixbuf {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_get_page_side_image((*C.GtkAssistant)(recv.native), c_page)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPageTitle is a wrapper around the C function gtk_assistant_get_page_title.
func (recv *Assistant) GetPageTitle(page *Widget) string {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_get_page_title((*C.GtkAssistant)(recv.native), c_page)
	retGo := C.GoString(retC)

	return retGo
}

// GetPageType is a wrapper around the C function gtk_assistant_get_page_type.
func (recv *Assistant) GetPageType(page *Widget) AssistantPageType {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_get_page_type((*C.GtkAssistant)(recv.native), c_page)
	retGo := (AssistantPageType)(retC)

	return retGo
}

// InsertPage is a wrapper around the C function gtk_assistant_insert_page.
func (recv *Assistant) InsertPage(page *Widget, position int32) int32 {
	c_page := (*C.GtkWidget)(page.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_assistant_insert_page((*C.GtkAssistant)(recv.native), c_page, c_position)
	retGo := (int32)(retC)

	return retGo
}

// PrependPage is a wrapper around the C function gtk_assistant_prepend_page.
func (recv *Assistant) PrependPage(page *Widget) int32 {
	c_page := (*C.GtkWidget)(page.ToC())

	retC := C.gtk_assistant_prepend_page((*C.GtkAssistant)(recv.native), c_page)
	retGo := (int32)(retC)

	return retGo
}

// RemoveActionWidget is a wrapper around the C function gtk_assistant_remove_action_widget.
func (recv *Assistant) RemoveActionWidget(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_assistant_remove_action_widget((*C.GtkAssistant)(recv.native), c_child)

	return
}

// SetCurrentPage is a wrapper around the C function gtk_assistant_set_current_page.
func (recv *Assistant) SetCurrentPage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_assistant_set_current_page((*C.GtkAssistant)(recv.native), c_page_num)

	return
}

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// SetPageComplete is a wrapper around the C function gtk_assistant_set_page_complete.
func (recv *Assistant) SetPageComplete(page *Widget, complete bool) {
	c_page := (*C.GtkWidget)(page.ToC())

	c_complete :=
		boolToGboolean(complete)

	C.gtk_assistant_set_page_complete((*C.GtkAssistant)(recv.native), c_page, c_complete)

	return
}

// SetPageHeaderImage is a wrapper around the C function gtk_assistant_set_page_header_image.
func (recv *Assistant) SetPageHeaderImage(page *Widget, pixbuf *gdkpixbuf.Pixbuf) {
	c_page := (*C.GtkWidget)(page.ToC())

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_assistant_set_page_header_image((*C.GtkAssistant)(recv.native), c_page, c_pixbuf)

	return
}

// SetPageSideImage is a wrapper around the C function gtk_assistant_set_page_side_image.
func (recv *Assistant) SetPageSideImage(page *Widget, pixbuf *gdkpixbuf.Pixbuf) {
	c_page := (*C.GtkWidget)(page.ToC())

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_assistant_set_page_side_image((*C.GtkAssistant)(recv.native), c_page, c_pixbuf)

	return
}

// SetPageTitle is a wrapper around the C function gtk_assistant_set_page_title.
func (recv *Assistant) SetPageTitle(page *Widget, title string) {
	c_page := (*C.GtkWidget)(page.ToC())

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_assistant_set_page_title((*C.GtkAssistant)(recv.native), c_page, c_title)

	return
}

// SetPageType is a wrapper around the C function gtk_assistant_set_page_type.
func (recv *Assistant) SetPageType(page *Widget, type_ AssistantPageType) {
	c_page := (*C.GtkWidget)(page.ToC())

	c_type := (C.GtkAssistantPageType)(type_)

	C.gtk_assistant_set_page_type((*C.GtkAssistant)(recv.native), c_page, c_type)

	return
}

// UpdateButtonsState is a wrapper around the C function gtk_assistant_update_buttons_state.
func (recv *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state((*C.GtkAssistant)(recv.native))

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// GetImagePosition is a wrapper around the C function gtk_button_get_image_position.
func (recv *Button) GetImagePosition() PositionType {
	retC := C.gtk_button_get_image_position((*C.GtkButton)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// SetImagePosition is a wrapper around the C function gtk_button_set_image_position.
func (recv *Button) SetImagePosition(position PositionType) {
	c_position := (C.GtkPositionType)(position)

	C.gtk_button_set_image_position((*C.GtkButton)(recv.native), c_position)

	return
}

// Unsupported signal 'accel-cleared' for CellRendererAccel : unsupported parameter path_string : type utf8 :

// Unsupported signal 'accel-edited' for CellRendererAccel : unsupported parameter path_string : type utf8 :

// CellRendererAccelNew is a wrapper around the C function gtk_cell_renderer_accel_new.
func CellRendererAccelNew() *CellRendererAccel {
	retC := C.gtk_cell_renderer_accel_new()
	retGo := CellRendererAccelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererSpinNew is a wrapper around the C function gtk_cell_renderer_spin_new.
func CellRendererSpinNew() *CellRendererSpin {
	retC := C.gtk_cell_renderer_spin_new()
	retGo := CellRendererSpinNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_clipboard_request_rich_text : unsupported parameter callback : no type generator for ClipboardRichTextReceivedFunc, GtkClipboardRichTextReceivedFunc

// Unsupported : gtk_clipboard_wait_for_rich_text : unsupported parameter format : Blacklisted record : GdkAtom

// WaitIsRichTextAvailable is a wrapper around the C function gtk_clipboard_wait_is_rich_text_available.
func (recv *Clipboard) WaitIsRichTextAvailable(buffer *TextBuffer) bool {
	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	retC := C.gtk_clipboard_wait_is_rich_text_available((*C.GtkClipboard)(recv.native), c_buffer)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

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

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetInnerBorder is a wrapper around the C function gtk_entry_get_inner_border.
func (recv *Entry) GetInnerBorder() *Border {
	retC := C.gtk_entry_get_inner_border((*C.GtkEntry)(recv.native))
	var retGo (*Border)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BorderNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetInnerBorder is a wrapper around the C function gtk_entry_set_inner_border.
func (recv *Entry) SetInnerBorder(border *Border) {
	c_border := (*C.GtkBorder)(border.ToC())

	C.gtk_entry_set_inner_border((*C.GtkEntry)(recv.native), c_border)

	return
}

// Unsupported : EntryIconAccessible : no CType

// GetFocusOnClick is a wrapper around the C function gtk_file_chooser_button_get_focus_on_click.
func (recv *FileChooserButton) GetFocusOnClick() bool {
	retC := C.gtk_file_chooser_button_get_focus_on_click((*C.GtkFileChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFocusOnClick is a wrapper around the C function gtk_file_chooser_button_set_focus_on_click.
func (recv *FileChooserButton) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_file_chooser_button_set_focus_on_click((*C.GtkFileChooserButton)(recv.native), c_focus_on_click)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// GetLineWrapMode is a wrapper around the C function gtk_label_get_line_wrap_mode.
func (recv *Label) GetLineWrapMode() pango.WrapMode {
	retC := C.gtk_label_get_line_wrap_mode((*C.GtkLabel)(recv.native))
	retGo := (pango.WrapMode)(retC)

	return retGo
}

// SetLineWrapMode is a wrapper around the C function gtk_label_set_line_wrap_mode.
func (recv *Label) SetLineWrapMode(wrapMode pango.WrapMode) {
	c_wrap_mode := (C.PangoWrapMode)(wrapMode)

	C.gtk_label_set_line_wrap_mode((*C.GtkLabel)(recv.native), c_wrap_mode)

	return
}

// LinkButtonNew is a wrapper around the C function gtk_link_button_new.
func LinkButtonNew(uri string) *LinkButton {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_link_button_new(c_uri)
	retGo := LinkButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LinkButtonNewWithLabel is a wrapper around the C function gtk_link_button_new_with_label.
func LinkButtonNewWithLabel(uri string, label string) *LinkButton {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_link_button_new_with_label(c_uri, c_label)
	retGo := LinkButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function gtk_link_button_get_uri.
func (recv *LinkButton) GetUri() string {
	retC := C.gtk_link_button_get_uri((*C.GtkLinkButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetUri is a wrapper around the C function gtk_link_button_set_uri.
func (recv *LinkButton) SetUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.gtk_link_button_set_uri((*C.GtkLinkButton)(recv.native), c_uri)

	return
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// SetImage is a wrapper around the C function gtk_message_dialog_set_image.
func (recv *MessageDialog) SetImage(image *Widget) {
	c_image := (*C.GtkWidget)(image.ToC())

	C.gtk_message_dialog_set_image((*C.GtkMessageDialog)(recv.native), c_image)

	return
}

// Unsupported signal 'page-added' for Notebook : unsupported parameter page_num : type guint :

// Unsupported signal 'page-removed' for Notebook : unsupported parameter page_num : type guint :

// Unsupported signal 'page-reordered' for Notebook : unsupported parameter page_num : type guint :

// GetTabDetachable is a wrapper around the C function gtk_notebook_get_tab_detachable.
func (recv *Notebook) GetTabDetachable(child *Widget) bool {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_detachable((*C.GtkNotebook)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// GetTabReorderable is a wrapper around the C function gtk_notebook_get_tab_reorderable.
func (recv *Notebook) GetTabReorderable(child *Widget) bool {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_reorderable((*C.GtkNotebook)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// SetTabDetachable is a wrapper around the C function gtk_notebook_set_tab_detachable.
func (recv *Notebook) SetTabDetachable(child *Widget, detachable bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_detachable :=
		boolToGboolean(detachable)

	C.gtk_notebook_set_tab_detachable((*C.GtkNotebook)(recv.native), c_child, c_detachable)

	return
}

// SetTabReorderable is a wrapper around the C function gtk_notebook_set_tab_reorderable.
func (recv *Notebook) SetTabReorderable(child *Widget, reorderable bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_notebook_set_tab_reorderable((*C.GtkNotebook)(recv.native), c_child, c_reorderable)

	return
}

// PageSetupNew is a wrapper around the C function gtk_page_setup_new.
func PageSetupNew() *PageSetup {
	retC := C.gtk_page_setup_new()
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Copy is a wrapper around the C function gtk_page_setup_copy.
func (recv *PageSetup) Copy() *PageSetup {
	retC := C.gtk_page_setup_copy((*C.GtkPageSetup)(recv.native))
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBottomMargin is a wrapper around the C function gtk_page_setup_get_bottom_margin.
func (recv *PageSetup) GetBottomMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_bottom_margin((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetLeftMargin is a wrapper around the C function gtk_page_setup_get_left_margin.
func (recv *PageSetup) GetLeftMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_left_margin((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_page_setup_get_orientation.
func (recv *PageSetup) GetOrientation() PageOrientation {
	retC := C.gtk_page_setup_get_orientation((*C.GtkPageSetup)(recv.native))
	retGo := (PageOrientation)(retC)

	return retGo
}

// GetPageHeight is a wrapper around the C function gtk_page_setup_get_page_height.
func (recv *PageSetup) GetPageHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_page_height((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetPageWidth is a wrapper around the C function gtk_page_setup_get_page_width.
func (recv *PageSetup) GetPageWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_page_width((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetPaperHeight is a wrapper around the C function gtk_page_setup_get_paper_height.
func (recv *PageSetup) GetPaperHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_paper_height((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetPaperSize is a wrapper around the C function gtk_page_setup_get_paper_size.
func (recv *PageSetup) GetPaperSize() *PaperSize {
	retC := C.gtk_page_setup_get_paper_size((*C.GtkPageSetup)(recv.native))
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPaperWidth is a wrapper around the C function gtk_page_setup_get_paper_width.
func (recv *PageSetup) GetPaperWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_paper_width((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetRightMargin is a wrapper around the C function gtk_page_setup_get_right_margin.
func (recv *PageSetup) GetRightMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_right_margin((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetTopMargin is a wrapper around the C function gtk_page_setup_get_top_margin.
func (recv *PageSetup) GetTopMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_page_setup_get_top_margin((*C.GtkPageSetup)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// SetBottomMargin is a wrapper around the C function gtk_page_setup_set_bottom_margin.
func (recv *PageSetup) SetBottomMargin(margin float64, unit Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_page_setup_set_bottom_margin((*C.GtkPageSetup)(recv.native), c_margin, c_unit)

	return
}

// SetLeftMargin is a wrapper around the C function gtk_page_setup_set_left_margin.
func (recv *PageSetup) SetLeftMargin(margin float64, unit Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_page_setup_set_left_margin((*C.GtkPageSetup)(recv.native), c_margin, c_unit)

	return
}

// SetOrientation is a wrapper around the C function gtk_page_setup_set_orientation.
func (recv *PageSetup) SetOrientation(orientation PageOrientation) {
	c_orientation := (C.GtkPageOrientation)(orientation)

	C.gtk_page_setup_set_orientation((*C.GtkPageSetup)(recv.native), c_orientation)

	return
}

// SetPaperSize is a wrapper around the C function gtk_page_setup_set_paper_size.
func (recv *PageSetup) SetPaperSize(size *PaperSize) {
	c_size := (*C.GtkPaperSize)(size.ToC())

	C.gtk_page_setup_set_paper_size((*C.GtkPageSetup)(recv.native), c_size)

	return
}

// SetPaperSizeAndDefaultMargins is a wrapper around the C function gtk_page_setup_set_paper_size_and_default_margins.
func (recv *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	c_size := (*C.GtkPaperSize)(size.ToC())

	C.gtk_page_setup_set_paper_size_and_default_margins((*C.GtkPageSetup)(recv.native), c_size)

	return
}

// SetRightMargin is a wrapper around the C function gtk_page_setup_set_right_margin.
func (recv *PageSetup) SetRightMargin(margin float64, unit Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_page_setup_set_right_margin((*C.GtkPageSetup)(recv.native), c_margin, c_unit)

	return
}

// SetTopMargin is a wrapper around the C function gtk_page_setup_set_top_margin.
func (recv *PageSetup) SetTopMargin(margin float64, unit Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_page_setup_set_top_margin((*C.GtkPageSetup)(recv.native), c_margin, c_unit)

	return
}

// CreatePangoContext is a wrapper around the C function gtk_print_context_create_pango_context.
func (recv *PrintContext) CreatePangoContext() *pango.Context {
	retC := C.gtk_print_context_create_pango_context((*C.GtkPrintContext)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreatePangoLayout is a wrapper around the C function gtk_print_context_create_pango_layout.
func (recv *PrintContext) CreatePangoLayout() *pango.Layout {
	retC := C.gtk_print_context_create_pango_layout((*C.GtkPrintContext)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCairoContext is a wrapper around the C function gtk_print_context_get_cairo_context.
func (recv *PrintContext) GetCairoContext() *cairo.Context {
	retC := C.gtk_print_context_get_cairo_context((*C.GtkPrintContext)(recv.native))
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDpiX is a wrapper around the C function gtk_print_context_get_dpi_x.
func (recv *PrintContext) GetDpiX() float64 {
	retC := C.gtk_print_context_get_dpi_x((*C.GtkPrintContext)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetDpiY is a wrapper around the C function gtk_print_context_get_dpi_y.
func (recv *PrintContext) GetDpiY() float64 {
	retC := C.gtk_print_context_get_dpi_y((*C.GtkPrintContext)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetHeight is a wrapper around the C function gtk_print_context_get_height.
func (recv *PrintContext) GetHeight() float64 {
	retC := C.gtk_print_context_get_height((*C.GtkPrintContext)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPageSetup is a wrapper around the C function gtk_print_context_get_page_setup.
func (recv *PrintContext) GetPageSetup() *PageSetup {
	retC := C.gtk_print_context_get_page_setup((*C.GtkPrintContext)(recv.native))
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPangoFontmap is a wrapper around the C function gtk_print_context_get_pango_fontmap.
func (recv *PrintContext) GetPangoFontmap() *pango.FontMap {
	retC := C.gtk_print_context_get_pango_fontmap((*C.GtkPrintContext)(recv.native))
	retGo := pango.FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gtk_print_context_get_width.
func (recv *PrintContext) GetWidth() float64 {
	retC := C.gtk_print_context_get_width((*C.GtkPrintContext)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetCairoContext is a wrapper around the C function gtk_print_context_set_cairo_context.
func (recv *PrintContext) SetCairoContext(cr *cairo.Context, dpiX float64, dpiY float64) {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_dpi_x := (C.double)(dpiX)

	c_dpi_y := (C.double)(dpiY)

	C.gtk_print_context_set_cairo_context((*C.GtkPrintContext)(recv.native), c_cr, c_dpi_x, c_dpi_y)

	return
}

type signalPrintOperationBeginPrintDetail struct {
	callback  PrintOperationSignalBeginPrintCallback
	handlerID C.gulong
}

var signalPrintOperationBeginPrintId int
var signalPrintOperationBeginPrintMap = make(map[int]signalPrintOperationBeginPrintDetail)
var signalPrintOperationBeginPrintLock sync.Mutex

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
var signalPrintOperationCreateCustomWidgetLock sync.Mutex

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
var signalPrintOperationCustomWidgetApplyLock sync.Mutex

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
	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	index := int(uintptr(data))
	callback := signalPrintOperationCustomWidgetApplyMap[index].callback
	callback(widget)
}

// Unsupported signal 'done' for PrintOperation : unsupported parameter result : type PrintOperationResult :

// Unsupported signal 'draw-page' for PrintOperation : unsupported parameter page_nr : type gint :

type signalPrintOperationEndPrintDetail struct {
	callback  PrintOperationSignalEndPrintCallback
	handlerID C.gulong
}

var signalPrintOperationEndPrintId int
var signalPrintOperationEndPrintMap = make(map[int]signalPrintOperationEndPrintDetail)
var signalPrintOperationEndPrintLock sync.Mutex

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
var signalPrintOperationPaginateLock sync.Mutex

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
	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationPaginateMap[index].callback
	retGo := callback(context)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported signal 'preview' for PrintOperation : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported signal 'request-page-setup' for PrintOperation : unsupported parameter page_nr : type gint :

type signalPrintOperationStatusChangedDetail struct {
	callback  PrintOperationSignalStatusChangedCallback
	handlerID C.gulong
}

var signalPrintOperationStatusChangedId int
var signalPrintOperationStatusChangedMap = make(map[int]signalPrintOperationStatusChangedDetail)
var signalPrintOperationStatusChangedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalPrintOperationStatusChangedMap[index].callback
	callback()
}

// PrintOperationNew is a wrapper around the C function gtk_print_operation_new.
func PrintOperationNew() *PrintOperation {
	retC := C.gtk_print_operation_new()
	retGo := PrintOperationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Cancel is a wrapper around the C function gtk_print_operation_cancel.
func (recv *PrintOperation) Cancel() {
	C.gtk_print_operation_cancel((*C.GtkPrintOperation)(recv.native))

	return
}

// GetDefaultPageSetup is a wrapper around the C function gtk_print_operation_get_default_page_setup.
func (recv *PrintOperation) GetDefaultPageSetup() *PageSetup {
	retC := C.gtk_print_operation_get_default_page_setup((*C.GtkPrintOperation)(recv.native))
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetError is a wrapper around the C function gtk_print_operation_get_error.
func (recv *PrintOperation) GetError() error {
	var cThrowableError *C.GError

	C.gtk_print_operation_get_error((*C.GtkPrintOperation)(recv.native), &cThrowableError)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return goThrowableError
}

// GetPrintSettings is a wrapper around the C function gtk_print_operation_get_print_settings.
func (recv *PrintOperation) GetPrintSettings() *PrintSettings {
	retC := C.gtk_print_operation_get_print_settings((*C.GtkPrintOperation)(recv.native))
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStatus is a wrapper around the C function gtk_print_operation_get_status.
func (recv *PrintOperation) GetStatus() PrintStatus {
	retC := C.gtk_print_operation_get_status((*C.GtkPrintOperation)(recv.native))
	retGo := (PrintStatus)(retC)

	return retGo
}

// GetStatusString is a wrapper around the C function gtk_print_operation_get_status_string.
func (recv *PrintOperation) GetStatusString() string {
	retC := C.gtk_print_operation_get_status_string((*C.GtkPrintOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsFinished is a wrapper around the C function gtk_print_operation_is_finished.
func (recv *PrintOperation) IsFinished() bool {
	retC := C.gtk_print_operation_is_finished((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Run is a wrapper around the C function gtk_print_operation_run.
func (recv *PrintOperation) Run(action PrintOperationAction, parent *Window) (PrintOperationResult, error) {
	c_action := (C.GtkPrintOperationAction)(action)

	c_parent := (*C.GtkWindow)(parent.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_print_operation_run((*C.GtkPrintOperation)(recv.native), c_action, c_parent, &cThrowableError)
	retGo := (PrintOperationResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetAllowAsync is a wrapper around the C function gtk_print_operation_set_allow_async.
func (recv *PrintOperation) SetAllowAsync(allowAsync bool) {
	c_allow_async :=
		boolToGboolean(allowAsync)

	C.gtk_print_operation_set_allow_async((*C.GtkPrintOperation)(recv.native), c_allow_async)

	return
}

// SetCurrentPage is a wrapper around the C function gtk_print_operation_set_current_page.
func (recv *PrintOperation) SetCurrentPage(currentPage int32) {
	c_current_page := (C.gint)(currentPage)

	C.gtk_print_operation_set_current_page((*C.GtkPrintOperation)(recv.native), c_current_page)

	return
}

// SetCustomTabLabel is a wrapper around the C function gtk_print_operation_set_custom_tab_label.
func (recv *PrintOperation) SetCustomTabLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_print_operation_set_custom_tab_label((*C.GtkPrintOperation)(recv.native), c_label)

	return
}

// SetDefaultPageSetup is a wrapper around the C function gtk_print_operation_set_default_page_setup.
func (recv *PrintOperation) SetDefaultPageSetup(defaultPageSetup *PageSetup) {
	c_default_page_setup := (*C.GtkPageSetup)(defaultPageSetup.ToC())

	C.gtk_print_operation_set_default_page_setup((*C.GtkPrintOperation)(recv.native), c_default_page_setup)

	return
}

// SetExportFilename is a wrapper around the C function gtk_print_operation_set_export_filename.
func (recv *PrintOperation) SetExportFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_print_operation_set_export_filename((*C.GtkPrintOperation)(recv.native), c_filename)

	return
}

// SetJobName is a wrapper around the C function gtk_print_operation_set_job_name.
func (recv *PrintOperation) SetJobName(jobName string) {
	c_job_name := C.CString(jobName)
	defer C.free(unsafe.Pointer(c_job_name))

	C.gtk_print_operation_set_job_name((*C.GtkPrintOperation)(recv.native), c_job_name)

	return
}

// SetNPages is a wrapper around the C function gtk_print_operation_set_n_pages.
func (recv *PrintOperation) SetNPages(nPages int32) {
	c_n_pages := (C.gint)(nPages)

	C.gtk_print_operation_set_n_pages((*C.GtkPrintOperation)(recv.native), c_n_pages)

	return
}

// SetPrintSettings is a wrapper around the C function gtk_print_operation_set_print_settings.
func (recv *PrintOperation) SetPrintSettings(printSettings *PrintSettings) {
	c_print_settings := (*C.GtkPrintSettings)(printSettings.ToC())

	C.gtk_print_operation_set_print_settings((*C.GtkPrintOperation)(recv.native), c_print_settings)

	return
}

// SetShowProgress is a wrapper around the C function gtk_print_operation_set_show_progress.
func (recv *PrintOperation) SetShowProgress(showProgress bool) {
	c_show_progress :=
		boolToGboolean(showProgress)

	C.gtk_print_operation_set_show_progress((*C.GtkPrintOperation)(recv.native), c_show_progress)

	return
}

// SetTrackPrintStatus is a wrapper around the C function gtk_print_operation_set_track_print_status.
func (recv *PrintOperation) SetTrackPrintStatus(trackStatus bool) {
	c_track_status :=
		boolToGboolean(trackStatus)

	C.gtk_print_operation_set_track_print_status((*C.GtkPrintOperation)(recv.native), c_track_status)

	return
}

// SetUnit is a wrapper around the C function gtk_print_operation_set_unit.
func (recv *PrintOperation) SetUnit(unit Unit) {
	c_unit := (C.GtkUnit)(unit)

	C.gtk_print_operation_set_unit((*C.GtkPrintOperation)(recv.native), c_unit)

	return
}

// SetUseFullPage is a wrapper around the C function gtk_print_operation_set_use_full_page.
func (recv *PrintOperation) SetUseFullPage(fullPage bool) {
	c_full_page :=
		boolToGboolean(fullPage)

	C.gtk_print_operation_set_use_full_page((*C.GtkPrintOperation)(recv.native), c_full_page)

	return
}

// PrintSettingsNew is a wrapper around the C function gtk_print_settings_new.
func PrintSettingsNew() *PrintSettings {
	retC := C.gtk_print_settings_new()
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Copy is a wrapper around the C function gtk_print_settings_copy.
func (recv *PrintSettings) Copy() *PrintSettings {
	retC := C.gtk_print_settings_copy((*C.GtkPrintSettings)(recv.native))
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc, GtkPrintSettingsFunc

// Get is a wrapper around the C function gtk_print_settings_get.
func (recv *PrintSettings) Get(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_print_settings_get((*C.GtkPrintSettings)(recv.native), c_key)
	retGo := C.GoString(retC)

	return retGo
}

// GetBool is a wrapper around the C function gtk_print_settings_get_bool.
func (recv *PrintSettings) GetBool(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_print_settings_get_bool((*C.GtkPrintSettings)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetCollate is a wrapper around the C function gtk_print_settings_get_collate.
func (recv *PrintSettings) GetCollate() bool {
	retC := C.gtk_print_settings_get_collate((*C.GtkPrintSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDefaultSource is a wrapper around the C function gtk_print_settings_get_default_source.
func (recv *PrintSettings) GetDefaultSource() string {
	retC := C.gtk_print_settings_get_default_source((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDither is a wrapper around the C function gtk_print_settings_get_dither.
func (recv *PrintSettings) GetDither() string {
	retC := C.gtk_print_settings_get_dither((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDouble is a wrapper around the C function gtk_print_settings_get_double.
func (recv *PrintSettings) GetDouble(key string) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_print_settings_get_double((*C.GtkPrintSettings)(recv.native), c_key)
	retGo := (float64)(retC)

	return retGo
}

// GetDoubleWithDefault is a wrapper around the C function gtk_print_settings_get_double_with_default.
func (recv *PrintSettings) GetDoubleWithDefault(key string, def float64) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_def := (C.gdouble)(def)

	retC := C.gtk_print_settings_get_double_with_default((*C.GtkPrintSettings)(recv.native), c_key, c_def)
	retGo := (float64)(retC)

	return retGo
}

// GetDuplex is a wrapper around the C function gtk_print_settings_get_duplex.
func (recv *PrintSettings) GetDuplex() PrintDuplex {
	retC := C.gtk_print_settings_get_duplex((*C.GtkPrintSettings)(recv.native))
	retGo := (PrintDuplex)(retC)

	return retGo
}

// GetFinishings is a wrapper around the C function gtk_print_settings_get_finishings.
func (recv *PrintSettings) GetFinishings() string {
	retC := C.gtk_print_settings_get_finishings((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInt is a wrapper around the C function gtk_print_settings_get_int.
func (recv *PrintSettings) GetInt(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_print_settings_get_int((*C.GtkPrintSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// GetIntWithDefault is a wrapper around the C function gtk_print_settings_get_int_with_default.
func (recv *PrintSettings) GetIntWithDefault(key string, def int32) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_def := (C.gint)(def)

	retC := C.gtk_print_settings_get_int_with_default((*C.GtkPrintSettings)(recv.native), c_key, c_def)
	retGo := (int32)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_print_settings_get_length.
func (recv *PrintSettings) GetLength(key string, unit Unit) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_print_settings_get_length((*C.GtkPrintSettings)(recv.native), c_key, c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetMediaType is a wrapper around the C function gtk_print_settings_get_media_type.
func (recv *PrintSettings) GetMediaType() string {
	retC := C.gtk_print_settings_get_media_type((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNCopies is a wrapper around the C function gtk_print_settings_get_n_copies.
func (recv *PrintSettings) GetNCopies() int32 {
	retC := C.gtk_print_settings_get_n_copies((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNumberUp is a wrapper around the C function gtk_print_settings_get_number_up.
func (recv *PrintSettings) GetNumberUp() int32 {
	retC := C.gtk_print_settings_get_number_up((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_print_settings_get_orientation.
func (recv *PrintSettings) GetOrientation() PageOrientation {
	retC := C.gtk_print_settings_get_orientation((*C.GtkPrintSettings)(recv.native))
	retGo := (PageOrientation)(retC)

	return retGo
}

// GetOutputBin is a wrapper around the C function gtk_print_settings_get_output_bin.
func (recv *PrintSettings) GetOutputBin() string {
	retC := C.gtk_print_settings_get_output_bin((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_print_settings_get_page_ranges : no return type

// GetPageSet is a wrapper around the C function gtk_print_settings_get_page_set.
func (recv *PrintSettings) GetPageSet() PageSet {
	retC := C.gtk_print_settings_get_page_set((*C.GtkPrintSettings)(recv.native))
	retGo := (PageSet)(retC)

	return retGo
}

// GetPaperHeight is a wrapper around the C function gtk_print_settings_get_paper_height.
func (recv *PrintSettings) GetPaperHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_print_settings_get_paper_height((*C.GtkPrintSettings)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetPaperSize is a wrapper around the C function gtk_print_settings_get_paper_size.
func (recv *PrintSettings) GetPaperSize() *PaperSize {
	retC := C.gtk_print_settings_get_paper_size((*C.GtkPrintSettings)(recv.native))
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPaperWidth is a wrapper around the C function gtk_print_settings_get_paper_width.
func (recv *PrintSettings) GetPaperWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_print_settings_get_paper_width((*C.GtkPrintSettings)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetPrintPages is a wrapper around the C function gtk_print_settings_get_print_pages.
func (recv *PrintSettings) GetPrintPages() PrintPages {
	retC := C.gtk_print_settings_get_print_pages((*C.GtkPrintSettings)(recv.native))
	retGo := (PrintPages)(retC)

	return retGo
}

// GetPrinter is a wrapper around the C function gtk_print_settings_get_printer.
func (recv *PrintSettings) GetPrinter() string {
	retC := C.gtk_print_settings_get_printer((*C.GtkPrintSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetQuality is a wrapper around the C function gtk_print_settings_get_quality.
func (recv *PrintSettings) GetQuality() PrintQuality {
	retC := C.gtk_print_settings_get_quality((*C.GtkPrintSettings)(recv.native))
	retGo := (PrintQuality)(retC)

	return retGo
}

// GetResolution is a wrapper around the C function gtk_print_settings_get_resolution.
func (recv *PrintSettings) GetResolution() int32 {
	retC := C.gtk_print_settings_get_resolution((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReverse is a wrapper around the C function gtk_print_settings_get_reverse.
func (recv *PrintSettings) GetReverse() bool {
	retC := C.gtk_print_settings_get_reverse((*C.GtkPrintSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetScale is a wrapper around the C function gtk_print_settings_get_scale.
func (recv *PrintSettings) GetScale() float64 {
	retC := C.gtk_print_settings_get_scale((*C.GtkPrintSettings)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetUseColor is a wrapper around the C function gtk_print_settings_get_use_color.
func (recv *PrintSettings) GetUseColor() bool {
	retC := C.gtk_print_settings_get_use_color((*C.GtkPrintSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasKey is a wrapper around the C function gtk_print_settings_has_key.
func (recv *PrintSettings) HasKey(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_print_settings_has_key((*C.GtkPrintSettings)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Set is a wrapper around the C function gtk_print_settings_set.
func (recv *PrintSettings) Set(key string, value string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.gtk_print_settings_set((*C.GtkPrintSettings)(recv.native), c_key, c_value)

	return
}

// SetBool is a wrapper around the C function gtk_print_settings_set_bool.
func (recv *PrintSettings) SetBool(key string, value bool) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	C.gtk_print_settings_set_bool((*C.GtkPrintSettings)(recv.native), c_key, c_value)

	return
}

// SetCollate is a wrapper around the C function gtk_print_settings_set_collate.
func (recv *PrintSettings) SetCollate(collate bool) {
	c_collate :=
		boolToGboolean(collate)

	C.gtk_print_settings_set_collate((*C.GtkPrintSettings)(recv.native), c_collate)

	return
}

// SetDefaultSource is a wrapper around the C function gtk_print_settings_set_default_source.
func (recv *PrintSettings) SetDefaultSource(defaultSource string) {
	c_default_source := C.CString(defaultSource)
	defer C.free(unsafe.Pointer(c_default_source))

	C.gtk_print_settings_set_default_source((*C.GtkPrintSettings)(recv.native), c_default_source)

	return
}

// SetDither is a wrapper around the C function gtk_print_settings_set_dither.
func (recv *PrintSettings) SetDither(dither string) {
	c_dither := C.CString(dither)
	defer C.free(unsafe.Pointer(c_dither))

	C.gtk_print_settings_set_dither((*C.GtkPrintSettings)(recv.native), c_dither)

	return
}

// SetDouble is a wrapper around the C function gtk_print_settings_set_double.
func (recv *PrintSettings) SetDouble(key string, value float64) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	C.gtk_print_settings_set_double((*C.GtkPrintSettings)(recv.native), c_key, c_value)

	return
}

// SetDuplex is a wrapper around the C function gtk_print_settings_set_duplex.
func (recv *PrintSettings) SetDuplex(duplex PrintDuplex) {
	c_duplex := (C.GtkPrintDuplex)(duplex)

	C.gtk_print_settings_set_duplex((*C.GtkPrintSettings)(recv.native), c_duplex)

	return
}

// SetFinishings is a wrapper around the C function gtk_print_settings_set_finishings.
func (recv *PrintSettings) SetFinishings(finishings string) {
	c_finishings := C.CString(finishings)
	defer C.free(unsafe.Pointer(c_finishings))

	C.gtk_print_settings_set_finishings((*C.GtkPrintSettings)(recv.native), c_finishings)

	return
}

// SetInt is a wrapper around the C function gtk_print_settings_set_int.
func (recv *PrintSettings) SetInt(key string, value int32) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	C.gtk_print_settings_set_int((*C.GtkPrintSettings)(recv.native), c_key, c_value)

	return
}

// SetLength is a wrapper around the C function gtk_print_settings_set_length.
func (recv *PrintSettings) SetLength(key string, value float64, unit Unit) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_print_settings_set_length((*C.GtkPrintSettings)(recv.native), c_key, c_value, c_unit)

	return
}

// SetMediaType is a wrapper around the C function gtk_print_settings_set_media_type.
func (recv *PrintSettings) SetMediaType(mediaType string) {
	c_media_type := C.CString(mediaType)
	defer C.free(unsafe.Pointer(c_media_type))

	C.gtk_print_settings_set_media_type((*C.GtkPrintSettings)(recv.native), c_media_type)

	return
}

// SetNCopies is a wrapper around the C function gtk_print_settings_set_n_copies.
func (recv *PrintSettings) SetNCopies(numCopies int32) {
	c_num_copies := (C.gint)(numCopies)

	C.gtk_print_settings_set_n_copies((*C.GtkPrintSettings)(recv.native), c_num_copies)

	return
}

// SetNumberUp is a wrapper around the C function gtk_print_settings_set_number_up.
func (recv *PrintSettings) SetNumberUp(numberUp int32) {
	c_number_up := (C.gint)(numberUp)

	C.gtk_print_settings_set_number_up((*C.GtkPrintSettings)(recv.native), c_number_up)

	return
}

// SetOrientation is a wrapper around the C function gtk_print_settings_set_orientation.
func (recv *PrintSettings) SetOrientation(orientation PageOrientation) {
	c_orientation := (C.GtkPageOrientation)(orientation)

	C.gtk_print_settings_set_orientation((*C.GtkPrintSettings)(recv.native), c_orientation)

	return
}

// SetOutputBin is a wrapper around the C function gtk_print_settings_set_output_bin.
func (recv *PrintSettings) SetOutputBin(outputBin string) {
	c_output_bin := C.CString(outputBin)
	defer C.free(unsafe.Pointer(c_output_bin))

	C.gtk_print_settings_set_output_bin((*C.GtkPrintSettings)(recv.native), c_output_bin)

	return
}

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// SetPageSet is a wrapper around the C function gtk_print_settings_set_page_set.
func (recv *PrintSettings) SetPageSet(pageSet PageSet) {
	c_page_set := (C.GtkPageSet)(pageSet)

	C.gtk_print_settings_set_page_set((*C.GtkPrintSettings)(recv.native), c_page_set)

	return
}

// SetPaperHeight is a wrapper around the C function gtk_print_settings_set_paper_height.
func (recv *PrintSettings) SetPaperHeight(height float64, unit Unit) {
	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_print_settings_set_paper_height((*C.GtkPrintSettings)(recv.native), c_height, c_unit)

	return
}

// SetPaperSize is a wrapper around the C function gtk_print_settings_set_paper_size.
func (recv *PrintSettings) SetPaperSize(paperSize *PaperSize) {
	c_paper_size := (*C.GtkPaperSize)(paperSize.ToC())

	C.gtk_print_settings_set_paper_size((*C.GtkPrintSettings)(recv.native), c_paper_size)

	return
}

// SetPaperWidth is a wrapper around the C function gtk_print_settings_set_paper_width.
func (recv *PrintSettings) SetPaperWidth(width float64, unit Unit) {
	c_width := (C.gdouble)(width)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_print_settings_set_paper_width((*C.GtkPrintSettings)(recv.native), c_width, c_unit)

	return
}

// SetPrintPages is a wrapper around the C function gtk_print_settings_set_print_pages.
func (recv *PrintSettings) SetPrintPages(pages PrintPages) {
	c_pages := (C.GtkPrintPages)(pages)

	C.gtk_print_settings_set_print_pages((*C.GtkPrintSettings)(recv.native), c_pages)

	return
}

// SetPrinter is a wrapper around the C function gtk_print_settings_set_printer.
func (recv *PrintSettings) SetPrinter(printer string) {
	c_printer := C.CString(printer)
	defer C.free(unsafe.Pointer(c_printer))

	C.gtk_print_settings_set_printer((*C.GtkPrintSettings)(recv.native), c_printer)

	return
}

// SetQuality is a wrapper around the C function gtk_print_settings_set_quality.
func (recv *PrintSettings) SetQuality(quality PrintQuality) {
	c_quality := (C.GtkPrintQuality)(quality)

	C.gtk_print_settings_set_quality((*C.GtkPrintSettings)(recv.native), c_quality)

	return
}

// SetResolution is a wrapper around the C function gtk_print_settings_set_resolution.
func (recv *PrintSettings) SetResolution(resolution int32) {
	c_resolution := (C.gint)(resolution)

	C.gtk_print_settings_set_resolution((*C.GtkPrintSettings)(recv.native), c_resolution)

	return
}

// SetReverse is a wrapper around the C function gtk_print_settings_set_reverse.
func (recv *PrintSettings) SetReverse(reverse bool) {
	c_reverse :=
		boolToGboolean(reverse)

	C.gtk_print_settings_set_reverse((*C.GtkPrintSettings)(recv.native), c_reverse)

	return
}

// SetScale is a wrapper around the C function gtk_print_settings_set_scale.
func (recv *PrintSettings) SetScale(scale float64) {
	c_scale := (C.gdouble)(scale)

	C.gtk_print_settings_set_scale((*C.GtkPrintSettings)(recv.native), c_scale)

	return
}

// SetUseColor is a wrapper around the C function gtk_print_settings_set_use_color.
func (recv *PrintSettings) SetUseColor(useColor bool) {
	c_use_color :=
		boolToGboolean(useColor)

	C.gtk_print_settings_set_use_color((*C.GtkPrintSettings)(recv.native), c_use_color)

	return
}

// Unset is a wrapper around the C function gtk_print_settings_unset.
func (recv *PrintSettings) Unset(key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	C.gtk_print_settings_unset((*C.GtkPrintSettings)(recv.native), c_key)

	return
}

// SetCurrentValue is a wrapper around the C function gtk_radio_action_set_current_value.
func (recv *RadioAction) SetCurrentValue(currentValue int32) {
	c_current_value := (C.gint)(currentValue)

	C.gtk_radio_action_set_current_value((*C.GtkRadioAction)(recv.native), c_current_value)

	return
}

// GetLowerStepperSensitivity is a wrapper around the C function gtk_range_get_lower_stepper_sensitivity.
func (recv *Range) GetLowerStepperSensitivity() SensitivityType {
	retC := C.gtk_range_get_lower_stepper_sensitivity((*C.GtkRange)(recv.native))
	retGo := (SensitivityType)(retC)

	return retGo
}

// GetUpperStepperSensitivity is a wrapper around the C function gtk_range_get_upper_stepper_sensitivity.
func (recv *Range) GetUpperStepperSensitivity() SensitivityType {
	retC := C.gtk_range_get_upper_stepper_sensitivity((*C.GtkRange)(recv.native))
	retGo := (SensitivityType)(retC)

	return retGo
}

// SetLowerStepperSensitivity is a wrapper around the C function gtk_range_set_lower_stepper_sensitivity.
func (recv *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	c_sensitivity := (C.GtkSensitivityType)(sensitivity)

	C.gtk_range_set_lower_stepper_sensitivity((*C.GtkRange)(recv.native), c_sensitivity)

	return
}

// SetUpperStepperSensitivity is a wrapper around the C function gtk_range_set_upper_stepper_sensitivity.
func (recv *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	c_sensitivity := (C.GtkSensitivityType)(sensitivity)

	C.gtk_range_set_upper_stepper_sensitivity((*C.GtkRange)(recv.native), c_sensitivity)

	return
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// RecentChooserMenuNew is a wrapper around the C function gtk_recent_chooser_menu_new.
func RecentChooserMenuNew() *RecentChooserMenu {
	retC := C.gtk_recent_chooser_menu_new()
	retGo := RecentChooserMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentChooserMenuNewForManager is a wrapper around the C function gtk_recent_chooser_menu_new_for_manager.
func RecentChooserMenuNewForManager(manager *RecentManager) *RecentChooserMenu {
	c_manager := (*C.GtkRecentManager)(manager.ToC())

	retC := C.gtk_recent_chooser_menu_new_for_manager(c_manager)
	retGo := RecentChooserMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShowNumbers is a wrapper around the C function gtk_recent_chooser_menu_get_show_numbers.
func (recv *RecentChooserMenu) GetShowNumbers() bool {
	retC := C.gtk_recent_chooser_menu_get_show_numbers((*C.GtkRecentChooserMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowNumbers is a wrapper around the C function gtk_recent_chooser_menu_set_show_numbers.
func (recv *RecentChooserMenu) SetShowNumbers(showNumbers bool) {
	c_show_numbers :=
		boolToGboolean(showNumbers)

	C.gtk_recent_chooser_menu_set_show_numbers((*C.GtkRecentChooserMenu)(recv.native), c_show_numbers)

	return
}

// RecentChooserWidgetNew is a wrapper around the C function gtk_recent_chooser_widget_new.
func RecentChooserWidgetNew() *RecentChooserWidget {
	retC := C.gtk_recent_chooser_widget_new()
	retGo := RecentChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentChooserWidgetNewForManager is a wrapper around the C function gtk_recent_chooser_widget_new_for_manager.
func RecentChooserWidgetNewForManager(manager *RecentManager) *RecentChooserWidget {
	c_manager := (*C.GtkRecentManager)(manager.ToC())

	retC := C.gtk_recent_chooser_widget_new_for_manager(c_manager)
	retGo := RecentChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentFilterNew is a wrapper around the C function gtk_recent_filter_new.
func RecentFilterNew() *RecentFilter {
	retC := C.gtk_recent_filter_new()
	retGo := RecentFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddAge is a wrapper around the C function gtk_recent_filter_add_age.
func (recv *RecentFilter) AddAge(days int32) {
	c_days := (C.gint)(days)

	C.gtk_recent_filter_add_age((*C.GtkRecentFilter)(recv.native), c_days)

	return
}

// AddApplication is a wrapper around the C function gtk_recent_filter_add_application.
func (recv *RecentFilter) AddApplication(application string) {
	c_application := C.CString(application)
	defer C.free(unsafe.Pointer(c_application))

	C.gtk_recent_filter_add_application((*C.GtkRecentFilter)(recv.native), c_application)

	return
}

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// AddGroup is a wrapper around the C function gtk_recent_filter_add_group.
func (recv *RecentFilter) AddGroup(group string) {
	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	C.gtk_recent_filter_add_group((*C.GtkRecentFilter)(recv.native), c_group)

	return
}

// AddMimeType is a wrapper around the C function gtk_recent_filter_add_mime_type.
func (recv *RecentFilter) AddMimeType(mimeType string) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.gtk_recent_filter_add_mime_type((*C.GtkRecentFilter)(recv.native), c_mime_type)

	return
}

// AddPattern is a wrapper around the C function gtk_recent_filter_add_pattern.
func (recv *RecentFilter) AddPattern(pattern string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_recent_filter_add_pattern((*C.GtkRecentFilter)(recv.native), c_pattern)

	return
}

// AddPixbufFormats is a wrapper around the C function gtk_recent_filter_add_pixbuf_formats.
func (recv *RecentFilter) AddPixbufFormats() {
	C.gtk_recent_filter_add_pixbuf_formats((*C.GtkRecentFilter)(recv.native))

	return
}

// Filter is a wrapper around the C function gtk_recent_filter_filter.
func (recv *RecentFilter) Filter(filterInfo *RecentFilterInfo) bool {
	c_filter_info := (*C.GtkRecentFilterInfo)(filterInfo.ToC())

	retC := C.gtk_recent_filter_filter((*C.GtkRecentFilter)(recv.native), c_filter_info)
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function gtk_recent_filter_get_name.
func (recv *RecentFilter) GetName() string {
	retC := C.gtk_recent_filter_get_name((*C.GtkRecentFilter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNeeded is a wrapper around the C function gtk_recent_filter_get_needed.
func (recv *RecentFilter) GetNeeded() RecentFilterFlags {
	retC := C.gtk_recent_filter_get_needed((*C.GtkRecentFilter)(recv.native))
	retGo := (RecentFilterFlags)(retC)

	return retGo
}

// SetName is a wrapper around the C function gtk_recent_filter_set_name.
func (recv *RecentFilter) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_recent_filter_set_name((*C.GtkRecentFilter)(recv.native), c_name)

	return
}

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

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *RecentManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to RecentManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentManager.
func CastToRecentManager(object *gobject.Object) *RecentManager {
	return RecentManagerNewFromC(object.ToC())
}

// RecentManagerNew is a wrapper around the C function gtk_recent_manager_new.
func RecentManagerNew() *RecentManager {
	retC := C.gtk_recent_manager_new()
	retGo := RecentManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFull is a wrapper around the C function gtk_recent_manager_add_full.
func (recv *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_recent_data := (*C.GtkRecentData)(recentData.ToC())

	retC := C.gtk_recent_manager_add_full((*C.GtkRecentManager)(recv.native), c_uri, c_recent_data)
	retGo := retC == C.TRUE

	return retGo
}

// AddItem is a wrapper around the C function gtk_recent_manager_add_item.
func (recv *RecentManager) AddItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_recent_manager_add_item((*C.GtkRecentManager)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// GetItems is a wrapper around the C function gtk_recent_manager_get_items.
func (recv *RecentManager) GetItems() *glib.List {
	retC := C.gtk_recent_manager_get_items((*C.GtkRecentManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasItem is a wrapper around the C function gtk_recent_manager_has_item.
func (recv *RecentManager) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_recent_manager_has_item((*C.GtkRecentManager)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// LookupItem is a wrapper around the C function gtk_recent_manager_lookup_item.
func (recv *RecentManager) LookupItem(uri string) (*RecentInfo, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_lookup_item((*C.GtkRecentManager)(recv.native), c_uri, &cThrowableError)
	var retGo (*RecentInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = RecentInfoNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MoveItem is a wrapper around the C function gtk_recent_manager_move_item.
func (recv *RecentManager) MoveItem(uri string, newUri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_new_uri := C.CString(newUri)
	defer C.free(unsafe.Pointer(c_new_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_move_item((*C.GtkRecentManager)(recv.native), c_uri, c_new_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PurgeItems is a wrapper around the C function gtk_recent_manager_purge_items.
func (recv *RecentManager) PurgeItems() (int32, error) {
	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_purge_items((*C.GtkRecentManager)(recv.native), &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveItem is a wrapper around the C function gtk_recent_manager_remove_item.
func (recv *RecentManager) RemoveItem(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_recent_manager_remove_item((*C.GtkRecentManager)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// UnsetPlacement is a wrapper around the C function gtk_scrolled_window_unset_placement.
func (recv *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement((*C.GtkScrolledWindow)(recv.native))

	return
}

// GetWidgets is a wrapper around the C function gtk_size_group_get_widgets.
func (recv *SizeGroup) GetWidgets() *glib.SList {
	retC := C.gtk_size_group_get_widgets((*C.GtkSizeGroup)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalSpinButtonWrappedDetail struct {
	callback  SpinButtonSignalWrappedCallback
	handlerID C.gulong
}

var signalSpinButtonWrappedId int
var signalSpinButtonWrappedMap = make(map[int]signalSpinButtonWrappedDetail)
var signalSpinButtonWrappedLock sync.Mutex

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
var signalStatusIconActivateLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalStatusIconActivateMap[index].callback
	callback()
}

// Unsupported signal 'popup-menu' for StatusIcon : unsupported parameter button : type guint :

// Unsupported signal 'size-changed' for StatusIcon : unsupported parameter size : type gint :

// StatusIconNew is a wrapper around the C function gtk_status_icon_new.
func StatusIconNew() *StatusIcon {
	retC := C.gtk_status_icon_new()
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StatusIconNewFromFile is a wrapper around the C function gtk_status_icon_new_from_file.
func StatusIconNewFromFile(filename string) *StatusIcon {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_status_icon_new_from_file(c_filename)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// StatusIconNewFromIconName is a wrapper around the C function gtk_status_icon_new_from_icon_name.
func StatusIconNewFromIconName(iconName string) *StatusIcon {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	retC := C.gtk_status_icon_new_from_icon_name(c_icon_name)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StatusIconNewFromPixbuf is a wrapper around the C function gtk_status_icon_new_from_pixbuf.
func StatusIconNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *StatusIcon {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	retC := C.gtk_status_icon_new_from_pixbuf(c_pixbuf)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StatusIconNewFromStock is a wrapper around the C function gtk_status_icon_new_from_stock.
func StatusIconNewFromStock(stockId string) *StatusIcon {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_status_icon_new_from_stock(c_stock_id)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// GetIconName is a wrapper around the C function gtk_status_icon_get_icon_name.
func (recv *StatusIcon) GetIconName() string {
	retC := C.gtk_status_icon_get_icon_name((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_status_icon_get_pixbuf.
func (recv *StatusIcon) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_status_icon_get_pixbuf((*C.GtkStatusIcon)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSize is a wrapper around the C function gtk_status_icon_get_size.
func (recv *StatusIcon) GetSize() int32 {
	retC := C.gtk_status_icon_get_size((*C.GtkStatusIcon)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStock is a wrapper around the C function gtk_status_icon_get_stock.
func (recv *StatusIcon) GetStock() string {
	retC := C.gtk_status_icon_get_stock((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStorageType is a wrapper around the C function gtk_status_icon_get_storage_type.
func (recv *StatusIcon) GetStorageType() ImageType {
	retC := C.gtk_status_icon_get_storage_type((*C.GtkStatusIcon)(recv.native))
	retGo := (ImageType)(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_status_icon_get_visible.
func (recv *StatusIcon) GetVisible() bool {
	retC := C.gtk_status_icon_get_visible((*C.GtkStatusIcon)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsEmbedded is a wrapper around the C function gtk_status_icon_is_embedded.
func (recv *StatusIcon) IsEmbedded() bool {
	retC := C.gtk_status_icon_is_embedded((*C.GtkStatusIcon)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFromFile is a wrapper around the C function gtk_status_icon_set_from_file.
func (recv *StatusIcon) SetFromFile(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_status_icon_set_from_file((*C.GtkStatusIcon)(recv.native), c_filename)

	return
}

// SetFromIconName is a wrapper around the C function gtk_status_icon_set_from_icon_name.
func (recv *StatusIcon) SetFromIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_status_icon_set_from_icon_name((*C.GtkStatusIcon)(recv.native), c_icon_name)

	return
}

// SetFromPixbuf is a wrapper around the C function gtk_status_icon_set_from_pixbuf.
func (recv *StatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_status_icon_set_from_pixbuf((*C.GtkStatusIcon)(recv.native), c_pixbuf)

	return
}

// SetFromStock is a wrapper around the C function gtk_status_icon_set_from_stock.
func (recv *StatusIcon) SetFromStock(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_status_icon_set_from_stock((*C.GtkStatusIcon)(recv.native), c_stock_id)

	return
}

// SetVisible is a wrapper around the C function gtk_status_icon_set_visible.
func (recv *StatusIcon) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_status_icon_set_visible((*C.GtkStatusIcon)(recv.native), c_visible)

	return
}

// LookupColor is a wrapper around the C function gtk_style_lookup_color.
func (recv *Style) LookupColor(colorName string) (bool, *gdk.Color) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkColor

	retC := C.gtk_style_lookup_color((*C.GtkStyle)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Unsupported : gtk_text_buffer_deserialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_get_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_set_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// GetCopyTargetList is a wrapper around the C function gtk_text_buffer_get_copy_target_list.
func (recv *TextBuffer) GetCopyTargetList() *TargetList {
	retC := C.gtk_text_buffer_get_copy_target_list((*C.GtkTextBuffer)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_get_deserialize_formats : no return type

// GetHasSelection is a wrapper around the C function gtk_text_buffer_get_has_selection.
func (recv *TextBuffer) GetHasSelection() bool {
	retC := C.gtk_text_buffer_get_has_selection((*C.GtkTextBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPasteTargetList is a wrapper around the C function gtk_text_buffer_get_paste_target_list.
func (recv *TextBuffer) GetPasteTargetList() *TargetList {
	retC := C.gtk_text_buffer_get_paste_target_list((*C.GtkTextBuffer)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_get_serialize_formats : no return type

// Unsupported : gtk_text_buffer_register_deserialize_format : unsupported parameter function : no type generator for TextBufferDeserializeFunc, GtkTextBufferDeserializeFunc

// Unsupported : gtk_text_buffer_register_deserialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_register_serialize_format : unsupported parameter function : no type generator for TextBufferSerializeFunc, GtkTextBufferSerializeFunc

// Unsupported : gtk_text_buffer_register_serialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_serialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_deserialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_serialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetEnableTreeLines is a wrapper around the C function gtk_tree_view_get_enable_tree_lines.
func (recv *TreeView) GetEnableTreeLines() bool {
	retC := C.gtk_tree_view_get_enable_tree_lines((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGridLines is a wrapper around the C function gtk_tree_view_get_grid_lines.
func (recv *TreeView) GetGridLines() TreeViewGridLines {
	retC := C.gtk_tree_view_get_grid_lines((*C.GtkTreeView)(recv.native))
	retGo := (TreeViewGridLines)(retC)

	return retGo
}

// GetHeadersClickable is a wrapper around the C function gtk_tree_view_get_headers_clickable.
func (recv *TreeView) GetHeadersClickable() bool {
	retC := C.gtk_tree_view_get_headers_clickable((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRubberBanding is a wrapper around the C function gtk_tree_view_get_rubber_banding.
func (recv *TreeView) GetRubberBanding() bool {
	retC := C.gtk_tree_view_get_rubber_banding((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSearchEntry is a wrapper around the C function gtk_tree_view_get_search_entry.
func (recv *TreeView) GetSearchEntry() *Entry {
	retC := C.gtk_tree_view_get_search_entry((*C.GtkTreeView)(recv.native))
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_search_position_func : no return generator

// SetEnableTreeLines is a wrapper around the C function gtk_tree_view_set_enable_tree_lines.
func (recv *TreeView) SetEnableTreeLines(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_tree_view_set_enable_tree_lines((*C.GtkTreeView)(recv.native), c_enabled)

	return
}

// SetGridLines is a wrapper around the C function gtk_tree_view_set_grid_lines.
func (recv *TreeView) SetGridLines(gridLines TreeViewGridLines) {
	c_grid_lines := (C.GtkTreeViewGridLines)(gridLines)

	C.gtk_tree_view_set_grid_lines((*C.GtkTreeView)(recv.native), c_grid_lines)

	return
}

// SetRubberBanding is a wrapper around the C function gtk_tree_view_set_rubber_banding.
func (recv *TreeView) SetRubberBanding(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_tree_view_set_rubber_banding((*C.GtkTreeView)(recv.native), c_enable)

	return
}

// SetSearchEntry is a wrapper around the C function gtk_tree_view_set_search_entry.
func (recv *TreeView) SetSearchEntry(entry *Entry) {
	c_entry := (*C.GtkEntry)(entry.ToC())

	C.gtk_tree_view_set_search_entry((*C.GtkTreeView)(recv.native), c_entry)

	return
}

// Unsupported : gtk_tree_view_set_search_position_func : unsupported parameter func : no type generator for TreeViewSearchPositionFunc, GtkTreeViewSearchPositionFunc

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// DragDestGetTrackMotion is a wrapper around the C function gtk_drag_dest_get_track_motion.
func (recv *Widget) DragDestGetTrackMotion() bool {
	retC := C.gtk_drag_dest_get_track_motion((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// DragDestSetTrackMotion is a wrapper around the C function gtk_drag_dest_set_track_motion.
func (recv *Widget) DragDestSetTrackMotion(trackMotion bool) {
	c_track_motion :=
		boolToGboolean(trackMotion)

	C.gtk_drag_dest_set_track_motion((*C.GtkWidget)(recv.native), c_track_motion)

	return
}

// IsComposited is a wrapper around the C function gtk_widget_is_composited.
func (recv *Widget) IsComposited() bool {
	retC := C.gtk_widget_is_composited((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDeletable is a wrapper around the C function gtk_window_get_deletable.
func (recv *Window) GetDeletable() bool {
	retC := C.gtk_window_get_deletable((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGroup is a wrapper around the C function gtk_window_get_group.
func (recv *Window) GetGroup() *WindowGroup {
	retC := C.gtk_window_get_group((*C.GtkWindow)(recv.native))
	retGo := WindowGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDeletable is a wrapper around the C function gtk_window_set_deletable.
func (recv *Window) SetDeletable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_deletable((*C.GtkWindow)(recv.native), c_setting)

	return
}
