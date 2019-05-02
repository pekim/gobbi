// Code generated - DO NOT EDIT.
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

	void entry_iconPressHandler(GObject *, GtkEntryIconPosition, GdkEventButton *, gpointer);

	static gulong Entry_signal_connect_icon_press(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "icon-press", G_CALLBACK(entry_iconPressHandler), data);
	}

*/
/*

	void entry_iconReleaseHandler(GObject *, GtkEntryIconPosition, GdkEventButton *, gpointer);

	static gulong Entry_signal_connect_icon_release(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "icon-release", G_CALLBACK(entry_iconReleaseHandler), data);
	}

*/
/*

	gboolean statusicon_queryTooltipHandler(GObject *, gint, gint, gboolean, GtkTooltip *, gpointer);

	static gulong StatusIcon_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(statusicon_queryTooltipHandler), data);
	}

*/
/*

	gboolean statusicon_scrollEventHandler(GObject *, GdkEventScroll *, gpointer);

	static gulong StatusIcon_signal_connect_scroll_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-event", G_CALLBACK(statusicon_scrollEventHandler), data);
	}

*/
/*

	void textbuffer_pasteDoneHandler(GObject *, GtkClipboard *, gpointer);

	static gulong TextBuffer_signal_connect_paste_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paste-done", G_CALLBACK(textbuffer_pasteDoneHandler), data);
	}

*/
import "C"

// BlockActivate is a wrapper around the C function gtk_action_block_activate.
func (recv *Action) BlockActivate() {
	C.gtk_action_block_activate((*C.GtkAction)(recv.native))

	return
}

// GetGicon is a wrapper around the C function gtk_action_get_gicon.
func (recv *Action) GetGicon() *gio.Icon {
	retC := C.gtk_action_get_gicon((*C.GtkAction)(recv.native))
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconName is a wrapper around the C function gtk_action_get_icon_name.
func (recv *Action) GetIconName() string {
	retC := C.gtk_action_get_icon_name((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIsImportant is a wrapper around the C function gtk_action_get_is_important.
func (recv *Action) GetIsImportant() bool {
	retC := C.gtk_action_get_is_important((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLabel is a wrapper around the C function gtk_action_get_label.
func (recv *Action) GetLabel() string {
	retC := C.gtk_action_get_label((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetShortLabel is a wrapper around the C function gtk_action_get_short_label.
func (recv *Action) GetShortLabel() string {
	retC := C.gtk_action_get_short_label((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStockId is a wrapper around the C function gtk_action_get_stock_id.
func (recv *Action) GetStockId() string {
	retC := C.gtk_action_get_stock_id((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTooltip is a wrapper around the C function gtk_action_get_tooltip.
func (recv *Action) GetTooltip() string {
	retC := C.gtk_action_get_tooltip((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisibleHorizontal is a wrapper around the C function gtk_action_get_visible_horizontal.
func (recv *Action) GetVisibleHorizontal() bool {
	retC := C.gtk_action_get_visible_horizontal((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleVertical is a wrapper around the C function gtk_action_get_visible_vertical.
func (recv *Action) GetVisibleVertical() bool {
	retC := C.gtk_action_get_visible_vertical((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetGicon is a wrapper around the C function gtk_action_set_gicon.
func (recv *Action) SetGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_action_set_gicon((*C.GtkAction)(recv.native), c_icon)

	return
}

// SetIconName is a wrapper around the C function gtk_action_set_icon_name.
func (recv *Action) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_action_set_icon_name((*C.GtkAction)(recv.native), c_icon_name)

	return
}

// SetIsImportant is a wrapper around the C function gtk_action_set_is_important.
func (recv *Action) SetIsImportant(isImportant bool) {
	c_is_important :=
		boolToGboolean(isImportant)

	C.gtk_action_set_is_important((*C.GtkAction)(recv.native), c_is_important)

	return
}

// SetLabel is a wrapper around the C function gtk_action_set_label.
func (recv *Action) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_action_set_label((*C.GtkAction)(recv.native), c_label)

	return
}

// SetShortLabel is a wrapper around the C function gtk_action_set_short_label.
func (recv *Action) SetShortLabel(shortLabel string) {
	c_short_label := C.CString(shortLabel)
	defer C.free(unsafe.Pointer(c_short_label))

	C.gtk_action_set_short_label((*C.GtkAction)(recv.native), c_short_label)

	return
}

// SetStockId is a wrapper around the C function gtk_action_set_stock_id.
func (recv *Action) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_action_set_stock_id((*C.GtkAction)(recv.native), c_stock_id)

	return
}

// SetTooltip is a wrapper around the C function gtk_action_set_tooltip.
func (recv *Action) SetTooltip(tooltip string) {
	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_action_set_tooltip((*C.GtkAction)(recv.native), c_tooltip)

	return
}

// SetVisibleHorizontal is a wrapper around the C function gtk_action_set_visible_horizontal.
func (recv *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	c_visible_horizontal :=
		boolToGboolean(visibleHorizontal)

	C.gtk_action_set_visible_horizontal((*C.GtkAction)(recv.native), c_visible_horizontal)

	return
}

// SetVisibleVertical is a wrapper around the C function gtk_action_set_visible_vertical.
func (recv *Action) SetVisibleVertical(visibleVertical bool) {
	c_visible_vertical :=
		boolToGboolean(visibleVertical)

	C.gtk_action_set_visible_vertical((*C.GtkAction)(recv.native), c_visible_vertical)

	return
}

// UnblockActivate is a wrapper around the C function gtk_action_unblock_activate.
func (recv *Action) UnblockActivate() {
	C.gtk_action_unblock_activate((*C.GtkAction)(recv.native))

	return
}

// GetModel is a wrapper around the C function gtk_cell_view_get_model.
func (recv *CellView) GetModel() *TreeModel {
	retC := C.gtk_cell_view_get_model((*C.GtkCellView)(recv.native))
	var retGo (*TreeModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreeModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type signalEntryIconPressDetail struct {
	callback  EntrySignalIconPressCallback
	handlerID C.gulong
}

var signalEntryIconPressId int
var signalEntryIconPressMap = make(map[int]signalEntryIconPressDetail)
var signalEntryIconPressLock sync.RWMutex

// EntrySignalIconPressCallback is a callback function for a 'icon-press' signal emitted from a Entry.
type EntrySignalIconPressCallback func(iconPos EntryIconPosition, event *gdk.EventButton)

/*
ConnectIconPress connects the callback to the 'icon-press' signal for the Entry.

The returned value represents the connection, and may be passed to DisconnectIconPress to remove it.
*/
func (recv *Entry) ConnectIconPress(callback EntrySignalIconPressCallback) int {
	signalEntryIconPressLock.Lock()
	defer signalEntryIconPressLock.Unlock()

	signalEntryIconPressId++
	instance := C.gpointer(recv.native)
	handlerID := C.Entry_signal_connect_icon_press(instance, C.gpointer(uintptr(signalEntryIconPressId)))

	detail := signalEntryIconPressDetail{callback, handlerID}
	signalEntryIconPressMap[signalEntryIconPressId] = detail

	return signalEntryIconPressId
}

/*
DisconnectIconPress disconnects a callback from the 'icon-press' signal for the Entry.

The connectionID should be a value returned from a call to ConnectIconPress.
*/
func (recv *Entry) DisconnectIconPress(connectionID int) {
	signalEntryIconPressLock.Lock()
	defer signalEntryIconPressLock.Unlock()

	detail, exists := signalEntryIconPressMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryIconPressMap, connectionID)
}

//export entry_iconPressHandler
func entry_iconPressHandler(_ *C.GObject, c_icon_pos C.GtkEntryIconPosition, c_event *C.GdkEventButton, data C.gpointer) {
	signalEntryIconPressLock.RLock()
	defer signalEntryIconPressLock.RUnlock()

	iconPos := EntryIconPosition(c_icon_pos)

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalEntryIconPressMap[index].callback
	callback(iconPos, event)
}

type signalEntryIconReleaseDetail struct {
	callback  EntrySignalIconReleaseCallback
	handlerID C.gulong
}

var signalEntryIconReleaseId int
var signalEntryIconReleaseMap = make(map[int]signalEntryIconReleaseDetail)
var signalEntryIconReleaseLock sync.RWMutex

// EntrySignalIconReleaseCallback is a callback function for a 'icon-release' signal emitted from a Entry.
type EntrySignalIconReleaseCallback func(iconPos EntryIconPosition, event *gdk.EventButton)

/*
ConnectIconRelease connects the callback to the 'icon-release' signal for the Entry.

The returned value represents the connection, and may be passed to DisconnectIconRelease to remove it.
*/
func (recv *Entry) ConnectIconRelease(callback EntrySignalIconReleaseCallback) int {
	signalEntryIconReleaseLock.Lock()
	defer signalEntryIconReleaseLock.Unlock()

	signalEntryIconReleaseId++
	instance := C.gpointer(recv.native)
	handlerID := C.Entry_signal_connect_icon_release(instance, C.gpointer(uintptr(signalEntryIconReleaseId)))

	detail := signalEntryIconReleaseDetail{callback, handlerID}
	signalEntryIconReleaseMap[signalEntryIconReleaseId] = detail

	return signalEntryIconReleaseId
}

/*
DisconnectIconRelease disconnects a callback from the 'icon-release' signal for the Entry.

The connectionID should be a value returned from a call to ConnectIconRelease.
*/
func (recv *Entry) DisconnectIconRelease(connectionID int) {
	signalEntryIconReleaseLock.Lock()
	defer signalEntryIconReleaseLock.Unlock()

	detail, exists := signalEntryIconReleaseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryIconReleaseMap, connectionID)
}

//export entry_iconReleaseHandler
func entry_iconReleaseHandler(_ *C.GObject, c_icon_pos C.GtkEntryIconPosition, c_event *C.GdkEventButton, data C.gpointer) {
	signalEntryIconReleaseLock.RLock()
	defer signalEntryIconReleaseLock.RUnlock()

	iconPos := EntryIconPosition(c_icon_pos)

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalEntryIconReleaseMap[index].callback
	callback(iconPos, event)
}

// GetCurrentIconDragSource is a wrapper around the C function gtk_entry_get_current_icon_drag_source.
func (recv *Entry) GetCurrentIconDragSource() int32 {
	retC := C.gtk_entry_get_current_icon_drag_source((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIconActivatable is a wrapper around the C function gtk_entry_get_icon_activatable.
func (recv *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_activatable((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := retC == C.TRUE

	return retGo
}

// GetIconAtPos is a wrapper around the C function gtk_entry_get_icon_at_pos.
func (recv *Entry) GetIconAtPos(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_entry_get_icon_at_pos((*C.GtkEntry)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// GetIconGicon is a wrapper around the C function gtk_entry_get_icon_gicon.
func (recv *Entry) GetIconGicon(iconPos EntryIconPosition) *gio.Icon {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_gicon((*C.GtkEntry)(recv.native), c_icon_pos)
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIconName is a wrapper around the C function gtk_entry_get_icon_name.
func (recv *Entry) GetIconName(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_name((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)

	return retGo
}

// GetIconPixbuf is a wrapper around the C function gtk_entry_get_icon_pixbuf.
func (recv *Entry) GetIconPixbuf(iconPos EntryIconPosition) *gdkpixbuf.Pixbuf {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_pixbuf((*C.GtkEntry)(recv.native), c_icon_pos)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIconSensitive is a wrapper around the C function gtk_entry_get_icon_sensitive.
func (recv *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_sensitive((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := retC == C.TRUE

	return retGo
}

// GetIconStock is a wrapper around the C function gtk_entry_get_icon_stock.
func (recv *Entry) GetIconStock(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_stock((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)

	return retGo
}

// GetIconStorageType is a wrapper around the C function gtk_entry_get_icon_storage_type.
func (recv *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_storage_type((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := (ImageType)(retC)

	return retGo
}

// GetIconTooltipMarkup is a wrapper around the C function gtk_entry_get_icon_tooltip_markup.
func (recv *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_tooltip_markup((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetIconTooltipText is a wrapper around the C function gtk_entry_get_icon_tooltip_text.
func (recv *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_tooltip_text((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetProgressFraction is a wrapper around the C function gtk_entry_get_progress_fraction.
func (recv *Entry) GetProgressFraction() float64 {
	retC := C.gtk_entry_get_progress_fraction((*C.GtkEntry)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetProgressPulseStep is a wrapper around the C function gtk_entry_get_progress_pulse_step.
func (recv *Entry) GetProgressPulseStep() float64 {
	retC := C.gtk_entry_get_progress_pulse_step((*C.GtkEntry)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// ProgressPulse is a wrapper around the C function gtk_entry_progress_pulse.
func (recv *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse((*C.GtkEntry)(recv.native))

	return
}

// SetIconActivatable is a wrapper around the C function gtk_entry_set_icon_activatable.
func (recv *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_activatable :=
		boolToGboolean(activatable)

	C.gtk_entry_set_icon_activatable((*C.GtkEntry)(recv.native), c_icon_pos, c_activatable)

	return
}

// SetIconDragSource is a wrapper around the C function gtk_entry_set_icon_drag_source.
func (recv *Entry) SetIconDragSource(iconPos EntryIconPosition, targetList *TargetList, actions gdk.DragAction) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_target_list := (*C.GtkTargetList)(C.NULL)
	if targetList != nil {
		c_target_list = (*C.GtkTargetList)(targetList.ToC())
	}

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_entry_set_icon_drag_source((*C.GtkEntry)(recv.native), c_icon_pos, c_target_list, c_actions)

	return
}

// SetIconFromGicon is a wrapper around the C function gtk_entry_set_icon_from_gicon.
func (recv *Entry) SetIconFromGicon(iconPos EntryIconPosition, icon *gio.Icon) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_entry_set_icon_from_gicon((*C.GtkEntry)(recv.native), c_icon_pos, c_icon)

	return
}

// SetIconFromIconName is a wrapper around the C function gtk_entry_set_icon_from_icon_name.
func (recv *Entry) SetIconFromIconName(iconPos EntryIconPosition, iconName string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_entry_set_icon_from_icon_name((*C.GtkEntry)(recv.native), c_icon_pos, c_icon_name)

	return
}

// SetIconFromPixbuf is a wrapper around the C function gtk_entry_set_icon_from_pixbuf.
func (recv *Entry) SetIconFromPixbuf(iconPos EntryIconPosition, pixbuf *gdkpixbuf.Pixbuf) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_entry_set_icon_from_pixbuf((*C.GtkEntry)(recv.native), c_icon_pos, c_pixbuf)

	return
}

// SetIconFromStock is a wrapper around the C function gtk_entry_set_icon_from_stock.
func (recv *Entry) SetIconFromStock(iconPos EntryIconPosition, stockId string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_entry_set_icon_from_stock((*C.GtkEntry)(recv.native), c_icon_pos, c_stock_id)

	return
}

// SetIconSensitive is a wrapper around the C function gtk_entry_set_icon_sensitive.
func (recv *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_entry_set_icon_sensitive((*C.GtkEntry)(recv.native), c_icon_pos, c_sensitive)

	return
}

// SetIconTooltipMarkup is a wrapper around the C function gtk_entry_set_icon_tooltip_markup.
func (recv *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_entry_set_icon_tooltip_markup((*C.GtkEntry)(recv.native), c_icon_pos, c_tooltip)

	return
}

// SetIconTooltipText is a wrapper around the C function gtk_entry_set_icon_tooltip_text.
func (recv *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_entry_set_icon_tooltip_text((*C.GtkEntry)(recv.native), c_icon_pos, c_tooltip)

	return
}

// SetProgressFraction is a wrapper around the C function gtk_entry_set_progress_fraction.
func (recv *Entry) SetProgressFraction(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_entry_set_progress_fraction((*C.GtkEntry)(recv.native), c_fraction)

	return
}

// SetProgressPulseStep is a wrapper around the C function gtk_entry_set_progress_pulse_step.
func (recv *Entry) SetProgressPulseStep(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_entry_set_progress_pulse_step((*C.GtkEntry)(recv.native), c_fraction)

	return
}

// UnsetInvisibleChar is a wrapper around the C function gtk_entry_unset_invisible_char.
func (recv *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char((*C.GtkEntry)(recv.native))

	return
}

// GetContextId is a wrapper around the C function gtk_im_multicontext_get_context_id.
func (recv *IMMulticontext) GetContextId() string {
	retC := C.gtk_im_multicontext_get_context_id((*C.GtkIMMulticontext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetContextId is a wrapper around the C function gtk_im_multicontext_set_context_id.
func (recv *IMMulticontext) SetContextId(contextId string) {
	c_context_id := C.CString(contextId)
	defer C.free(unsafe.Pointer(c_context_id))

	C.gtk_im_multicontext_set_context_id((*C.GtkIMMulticontext)(recv.native), c_context_id)

	return
}

// GetAlwaysShowImage is a wrapper around the C function gtk_image_menu_item_get_always_show_image.
func (recv *ImageMenuItem) GetAlwaysShowImage() bool {
	retC := C.gtk_image_menu_item_get_always_show_image((*C.GtkImageMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseStock is a wrapper around the C function gtk_image_menu_item_get_use_stock.
func (recv *ImageMenuItem) GetUseStock() bool {
	retC := C.gtk_image_menu_item_get_use_stock((*C.GtkImageMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAccelGroup is a wrapper around the C function gtk_image_menu_item_set_accel_group.
func (recv *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_image_menu_item_set_accel_group((*C.GtkImageMenuItem)(recv.native), c_accel_group)

	return
}

// SetAlwaysShowImage is a wrapper around the C function gtk_image_menu_item_set_always_show_image.
func (recv *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_image_menu_item_set_always_show_image((*C.GtkImageMenuItem)(recv.native), c_always_show)

	return
}

// SetUseStock is a wrapper around the C function gtk_image_menu_item_set_use_stock.
func (recv *ImageMenuItem) SetUseStock(useStock bool) {
	c_use_stock :=
		boolToGboolean(useStock)

	C.gtk_image_menu_item_set_use_stock((*C.GtkImageMenuItem)(recv.native), c_use_stock)

	return
}

// GetLabel is a wrapper around the C function gtk_menu_item_get_label.
func (recv *MenuItem) GetLabel() string {
	retC := C.gtk_menu_item_get_label((*C.GtkMenuItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_menu_item_get_use_underline.
func (recv *MenuItem) GetUseUnderline() bool {
	retC := C.gtk_menu_item_get_use_underline((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLabel is a wrapper around the C function gtk_menu_item_set_label.
func (recv *MenuItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_menu_item_set_label((*C.GtkMenuItem)(recv.native), c_label)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_menu_item_set_use_underline.
func (recv *MenuItem) SetUseUnderline(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_menu_item_set_use_underline((*C.GtkMenuItem)(recv.native), c_setting)

	return
}

// DrawPageFinish is a wrapper around the C function gtk_print_operation_draw_page_finish.
func (recv *PrintOperation) DrawPageFinish() {
	C.gtk_print_operation_draw_page_finish((*C.GtkPrintOperation)(recv.native))

	return
}

// SetDeferDrawing is a wrapper around the C function gtk_print_operation_set_defer_drawing.
func (recv *PrintOperation) SetDeferDrawing() {
	C.gtk_print_operation_set_defer_drawing((*C.GtkPrintOperation)(recv.native))

	return
}

// GetPrinterLpi is a wrapper around the C function gtk_print_settings_get_printer_lpi.
func (recv *PrintSettings) GetPrinterLpi() float64 {
	retC := C.gtk_print_settings_get_printer_lpi((*C.GtkPrintSettings)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetResolutionX is a wrapper around the C function gtk_print_settings_get_resolution_x.
func (recv *PrintSettings) GetResolutionX() int32 {
	retC := C.gtk_print_settings_get_resolution_x((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetResolutionY is a wrapper around the C function gtk_print_settings_get_resolution_y.
func (recv *PrintSettings) GetResolutionY() int32 {
	retC := C.gtk_print_settings_get_resolution_y((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetPrinterLpi is a wrapper around the C function gtk_print_settings_set_printer_lpi.
func (recv *PrintSettings) SetPrinterLpi(lpi float64) {
	c_lpi := (C.gdouble)(lpi)

	C.gtk_print_settings_set_printer_lpi((*C.GtkPrintSettings)(recv.native), c_lpi)

	return
}

// SetResolutionXy is a wrapper around the C function gtk_print_settings_set_resolution_xy.
func (recv *PrintSettings) SetResolutionXy(resolutionX int32, resolutionY int32) {
	c_resolution_x := (C.gint)(resolutionX)

	c_resolution_y := (C.gint)(resolutionY)

	C.gtk_print_settings_set_resolution_xy((*C.GtkPrintSettings)(recv.native), c_resolution_x, c_resolution_y)

	return
}

// AddMark is a wrapper around the C function gtk_scale_add_mark.
func (recv *Scale) AddMark(value float64, position PositionType, markup string) {
	c_value := (C.gdouble)(value)

	c_position := (C.GtkPositionType)(position)

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_scale_add_mark((*C.GtkScale)(recv.native), c_value, c_position, c_markup)

	return
}

// ClearMarks is a wrapper around the C function gtk_scale_clear_marks.
func (recv *Scale) ClearMarks() {
	C.gtk_scale_clear_marks((*C.GtkScale)(recv.native))

	return
}

type signalStatusIconQueryTooltipDetail struct {
	callback  StatusIconSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalStatusIconQueryTooltipId int
var signalStatusIconQueryTooltipMap = make(map[int]signalStatusIconQueryTooltipDetail)
var signalStatusIconQueryTooltipLock sync.RWMutex

// StatusIconSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a StatusIcon.
type StatusIconSignalQueryTooltipCallback func(x int32, y int32, keyboardMode bool, tooltip *Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *StatusIcon) ConnectQueryTooltip(callback StatusIconSignalQueryTooltipCallback) int {
	signalStatusIconQueryTooltipLock.Lock()
	defer signalStatusIconQueryTooltipLock.Unlock()

	signalStatusIconQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalStatusIconQueryTooltipId)))

	detail := signalStatusIconQueryTooltipDetail{callback, handlerID}
	signalStatusIconQueryTooltipMap[signalStatusIconQueryTooltipId] = detail

	return signalStatusIconQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *StatusIcon) DisconnectQueryTooltip(connectionID int) {
	signalStatusIconQueryTooltipLock.Lock()
	defer signalStatusIconQueryTooltipLock.Unlock()

	detail, exists := signalStatusIconQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconQueryTooltipMap, connectionID)
}

//export statusicon_queryTooltipHandler
func statusicon_queryTooltipHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_keyboard_mode C.gboolean, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalStatusIconQueryTooltipLock.RLock()
	defer signalStatusIconQueryTooltipLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	keyboardMode := c_keyboard_mode == C.TRUE

	tooltip := TooltipNewFromC(unsafe.Pointer(c_tooltip))

	index := int(uintptr(data))
	callback := signalStatusIconQueryTooltipMap[index].callback
	retGo := callback(x, y, keyboardMode, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalStatusIconScrollEventDetail struct {
	callback  StatusIconSignalScrollEventCallback
	handlerID C.gulong
}

var signalStatusIconScrollEventId int
var signalStatusIconScrollEventMap = make(map[int]signalStatusIconScrollEventDetail)
var signalStatusIconScrollEventLock sync.RWMutex

// StatusIconSignalScrollEventCallback is a callback function for a 'scroll-event' signal emitted from a StatusIcon.
type StatusIconSignalScrollEventCallback func(event *gdk.EventScroll) bool

/*
ConnectScrollEvent connects the callback to the 'scroll-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectScrollEvent to remove it.
*/
func (recv *StatusIcon) ConnectScrollEvent(callback StatusIconSignalScrollEventCallback) int {
	signalStatusIconScrollEventLock.Lock()
	defer signalStatusIconScrollEventLock.Unlock()

	signalStatusIconScrollEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_scroll_event(instance, C.gpointer(uintptr(signalStatusIconScrollEventId)))

	detail := signalStatusIconScrollEventDetail{callback, handlerID}
	signalStatusIconScrollEventMap[signalStatusIconScrollEventId] = detail

	return signalStatusIconScrollEventId
}

/*
DisconnectScrollEvent disconnects a callback from the 'scroll-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectScrollEvent.
*/
func (recv *StatusIcon) DisconnectScrollEvent(connectionID int) {
	signalStatusIconScrollEventLock.Lock()
	defer signalStatusIconScrollEventLock.Unlock()

	detail, exists := signalStatusIconScrollEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconScrollEventMap, connectionID)
}

//export statusicon_scrollEventHandler
func statusicon_scrollEventHandler(_ *C.GObject, c_event *C.GdkEventScroll, data C.gpointer) C.gboolean {
	signalStatusIconScrollEventLock.RLock()
	defer signalStatusIconScrollEventLock.RUnlock()

	event := gdk.EventScrollNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconScrollEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// GetHasTooltip is a wrapper around the C function gtk_status_icon_get_has_tooltip.
func (recv *StatusIcon) GetHasTooltip() bool {
	retC := C.gtk_status_icon_get_has_tooltip((*C.GtkStatusIcon)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTooltipMarkup is a wrapper around the C function gtk_status_icon_get_tooltip_markup.
func (recv *StatusIcon) GetTooltipMarkup() string {
	retC := C.gtk_status_icon_get_tooltip_markup((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipText is a wrapper around the C function gtk_status_icon_get_tooltip_text.
func (recv *StatusIcon) GetTooltipText() string {
	retC := C.gtk_status_icon_get_tooltip_text((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// SetHasTooltip is a wrapper around the C function gtk_status_icon_set_has_tooltip.
func (recv *StatusIcon) SetHasTooltip(hasTooltip bool) {
	c_has_tooltip :=
		boolToGboolean(hasTooltip)

	C.gtk_status_icon_set_has_tooltip((*C.GtkStatusIcon)(recv.native), c_has_tooltip)

	return
}

// SetTooltipMarkup is a wrapper around the C function gtk_status_icon_set_tooltip_markup.
func (recv *StatusIcon) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_status_icon_set_tooltip_markup((*C.GtkStatusIcon)(recv.native), c_markup)

	return
}

// SetTooltipText is a wrapper around the C function gtk_status_icon_set_tooltip_text.
func (recv *StatusIcon) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_status_icon_set_tooltip_text((*C.GtkStatusIcon)(recv.native), c_text)

	return
}

// Unsupported : gtk_style_get : unsupported parameter ... : varargs

// GetStyleProperty is a wrapper around the C function gtk_style_get_style_property.
func (recv *Style) GetStyleProperty(widgetType gobject.Type, propertyName string) *gobject.Value {
	c_widget_type := (C.GType)(widgetType)

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	var c_value C.GValue

	C.gtk_style_get_style_property((*C.GtkStyle)(recv.native), c_widget_type, c_property_name, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Unsupported : gtk_style_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

type signalTextBufferPasteDoneDetail struct {
	callback  TextBufferSignalPasteDoneCallback
	handlerID C.gulong
}

var signalTextBufferPasteDoneId int
var signalTextBufferPasteDoneMap = make(map[int]signalTextBufferPasteDoneDetail)
var signalTextBufferPasteDoneLock sync.RWMutex

// TextBufferSignalPasteDoneCallback is a callback function for a 'paste-done' signal emitted from a TextBuffer.
type TextBufferSignalPasteDoneCallback func(clipboard *Clipboard)

/*
ConnectPasteDone connects the callback to the 'paste-done' signal for the TextBuffer.

The returned value represents the connection, and may be passed to DisconnectPasteDone to remove it.
*/
func (recv *TextBuffer) ConnectPasteDone(callback TextBufferSignalPasteDoneCallback) int {
	signalTextBufferPasteDoneLock.Lock()
	defer signalTextBufferPasteDoneLock.Unlock()

	signalTextBufferPasteDoneId++
	instance := C.gpointer(recv.native)
	handlerID := C.TextBuffer_signal_connect_paste_done(instance, C.gpointer(uintptr(signalTextBufferPasteDoneId)))

	detail := signalTextBufferPasteDoneDetail{callback, handlerID}
	signalTextBufferPasteDoneMap[signalTextBufferPasteDoneId] = detail

	return signalTextBufferPasteDoneId
}

/*
DisconnectPasteDone disconnects a callback from the 'paste-done' signal for the TextBuffer.

The connectionID should be a value returned from a call to ConnectPasteDone.
*/
func (recv *TextBuffer) DisconnectPasteDone(connectionID int) {
	signalTextBufferPasteDoneLock.Lock()
	defer signalTextBufferPasteDoneLock.Unlock()

	detail, exists := signalTextBufferPasteDoneMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextBufferPasteDoneMap, connectionID)
}

//export textbuffer_pasteDoneHandler
func textbuffer_pasteDoneHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, data C.gpointer) {
	signalTextBufferPasteDoneLock.RLock()
	defer signalTextBufferPasteDoneLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	index := int(uintptr(data))
	callback := signalTextBufferPasteDoneMap[index].callback
	callback(clipboard)
}

// WindowGetDefaultIconName is a wrapper around the C function gtk_window_get_default_icon_name.
func WindowGetDefaultIconName() string {
	retC := C.gtk_window_get_default_icon_name()
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : STOCK_CAPS_LOCK_WARNING

type EntryIconPosition C.GtkEntryIconPosition

const (
	GTK_ENTRY_ICON_PRIMARY   EntryIconPosition = 0
	GTK_ENTRY_ICON_SECONDARY EntryIconPosition = 1
)

// DoSetRelatedAction is a wrapper around the C function gtk_activatable_do_set_related_action.
func (recv *Activatable) DoSetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_do_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// GetRelatedAction is a wrapper around the C function gtk_activatable_get_related_action.
func (recv *Activatable) GetRelatedAction() *Action {
	retC := C.gtk_activatable_get_related_action((*C.GtkActivatable)(recv.native))
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUseActionAppearance is a wrapper around the C function gtk_activatable_get_use_action_appearance.
func (recv *Activatable) GetUseActionAppearance() bool {
	retC := C.gtk_activatable_get_use_action_appearance((*C.GtkActivatable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetRelatedAction is a wrapper around the C function gtk_activatable_set_related_action.
func (recv *Activatable) SetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// SetUseActionAppearance is a wrapper around the C function gtk_activatable_set_use_action_appearance.
func (recv *Activatable) SetUseActionAppearance(useAppearance bool) {
	c_use_appearance :=
		boolToGboolean(useAppearance)

	C.gtk_activatable_set_use_action_appearance((*C.GtkActivatable)(recv.native), c_use_appearance)

	return
}

// SyncActionProperties is a wrapper around the C function gtk_activatable_sync_action_properties.
func (recv *Activatable) SyncActionProperties(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_sync_action_properties((*C.GtkActivatable)(recv.native), c_action)

	return
}

// GetOrientation is a wrapper around the C function gtk_orientable_get_orientation.
func (recv *Orientable) GetOrientation() Orientation {
	retC := C.gtk_orientable_get_orientation((*C.GtkOrientable)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// SetOrientation is a wrapper around the C function gtk_orientable_set_orientation.
func (recv *Orientable) SetOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_orientable_set_orientation((*C.GtkOrientable)(recv.native), c_orientation)

	return
}

// ActivatableIface is a wrapper around the C record GtkActivatableIface.
type ActivatableIface struct {
	native *C.GtkActivatableIface
	// Private : g_iface
	// no type for update
	// no type for sync_action_properties
}

func ActivatableIfaceNewFromC(u unsafe.Pointer) *ActivatableIface {
	c := (*C.GtkActivatableIface)(u)
	if c == nil {
		return nil
	}

	g := &ActivatableIface{native: c}

	return g
}

func (recv *ActivatableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActivatableIface with another ActivatableIface, and returns true if they represent the same GObject.
func (recv *ActivatableIface) Equals(other *ActivatableIface) bool {
	return other.ToC() == recv.ToC()
}

// GetSelection is a wrapper around the C function gtk_selection_data_get_selection.
func (recv *SelectionData) GetSelection() gdk.Atom {
	retC := C.gtk_selection_data_get_selection((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}
