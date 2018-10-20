// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void StatusIcon_scrollEventHandler();

	static gulong StatusIcon_signal_connect_scroll_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-event", StatusIcon_scrollEventHandler, data);
	}

*/
/*

	void TextBuffer_pasteDoneHandler();

	static gulong TextBuffer_signal_connect_paste_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paste-done", TextBuffer_pasteDoneHandler, data);
	}

*/
import "C"

// BlockActivate is a wrapper around the C function gtk_action_block_activate.
func (recv *Action) BlockActivate() {
	C.gtk_action_block_activate((*C.GtkAction)(recv.native))

	return
}

// Unsupported : gtk_action_get_gicon : no return generator

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

// Unsupported : gtk_action_set_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_cell_view_get_model : no return generator

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported signal 'icon-press' for Entry : unsupported parameter icon_pos : type EntryIconPosition :

// Unsupported signal 'icon-release' for Entry : unsupported parameter icon_pos : type EntryIconPosition :

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

// Unsupported : gtk_entry_get_icon_gicon : no return generator

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
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

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

	c_target_list := (*C.GtkTargetList)(targetList.ToC())

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_entry_set_icon_drag_source((*C.GtkEntry)(recv.native), c_icon_pos, c_target_list, c_actions)

	return
}

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

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

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

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

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

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
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

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

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

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

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

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

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

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

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

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

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'query-tooltip' for StatusIcon : unsupported parameter x : type gint :

type signalStatusIconScrollEventDetail struct {
	callback  StatusIconSignalScrollEventCallback
	handlerID C.gulong
}

var signalStatusIconScrollEventId int
var signalStatusIconScrollEventMap = make(map[int]signalStatusIconScrollEventDetail)
var signalStatusIconScrollEventLock sync.Mutex

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
	instance := C.gpointer(recv.Object().ToC())
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

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconScrollEventMap, connectionID)
}

//export StatusIcon_scrollEventHandler
func StatusIcon_scrollEventHandler(_ *C.GObject, c_event *C.GdkEventScroll, data C.gpointer) {
	event := gdk.EventScrollNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconScrollEventMap[index].callback
	callback(event)
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

// Unsupported : gtk_style_get : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_style_property : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_valist : unsupported parameter widget_type : no type generator for GType, GType

type signalTextBufferPasteDoneDetail struct {
	callback  TextBufferSignalPasteDoneCallback
	handlerID C.gulong
}

var signalTextBufferPasteDoneId int
var signalTextBufferPasteDoneMap = make(map[int]signalTextBufferPasteDoneDetail)
var signalTextBufferPasteDoneLock sync.Mutex

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
	instance := C.gpointer(recv.Object().ToC())
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

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextBufferPasteDoneMap, connectionID)
}

//export TextBuffer_pasteDoneHandler
func TextBuffer_pasteDoneHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, data C.gpointer) {
	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	index := int(uintptr(data))
	callback := signalTextBufferPasteDoneMap[index].callback
	callback(clipboard)
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType
