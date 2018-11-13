// This is a generated file - DO NOT EDIT
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
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
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

// Disable activation signals from the action
//
// This is needed when updating the state of your proxy
// #GtkActivatable widget could result in calling gtk_action_activate(),
// this is a convenience function to avoid recursing in those
// cases (updating toggle state for instance).
/*

C function

gtk_action_block_activate
*/
func (recv *Action) BlockActivate() {
	C.gtk_action_block_activate((*C.GtkAction)(recv.native))

	return
}

// Gets the gicon of @action.
/*

C function

gtk_action_get_gicon
*/
func (recv *Action) GetGicon() *gio.Icon {
	retC := C.gtk_action_get_gicon((*C.GtkAction)(recv.native))
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the icon name of @action.
/*

C function

gtk_action_get_icon_name
*/
func (recv *Action) GetIconName() string {
	retC := C.gtk_action_get_icon_name((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks whether @action is important or not
/*

C function

gtk_action_get_is_important
*/
func (recv *Action) GetIsImportant() bool {
	retC := C.gtk_action_get_is_important((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the label text of @action.
/*

C function

gtk_action_get_label
*/
func (recv *Action) GetLabel() string {
	retC := C.gtk_action_get_label((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the short label text of @action.
/*

C function

gtk_action_get_short_label
*/
func (recv *Action) GetShortLabel() string {
	retC := C.gtk_action_get_short_label((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the stock id of @action.
/*

C function

gtk_action_get_stock_id
*/
func (recv *Action) GetStockId() string {
	retC := C.gtk_action_get_stock_id((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the tooltip text of @action.
/*

C function

gtk_action_get_tooltip
*/
func (recv *Action) GetTooltip() string {
	retC := C.gtk_action_get_tooltip((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks whether @action is visible when horizontal
/*

C function

gtk_action_get_visible_horizontal
*/
func (recv *Action) GetVisibleHorizontal() bool {
	retC := C.gtk_action_get_visible_horizontal((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether @action is visible when horizontal
/*

C function

gtk_action_get_visible_vertical
*/
func (recv *Action) GetVisibleVertical() bool {
	retC := C.gtk_action_get_visible_vertical((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the icon of @action.
/*

C function

gtk_action_set_gicon
*/
func (recv *Action) SetGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_action_set_gicon((*C.GtkAction)(recv.native), c_icon)

	return
}

// Sets the icon name on @action
/*

C function

gtk_action_set_icon_name
*/
func (recv *Action) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_action_set_icon_name((*C.GtkAction)(recv.native), c_icon_name)

	return
}

// Sets whether the action is important, this attribute is used
// primarily by toolbar items to decide whether to show a label
// or not.
/*

C function

gtk_action_set_is_important
*/
func (recv *Action) SetIsImportant(isImportant bool) {
	c_is_important :=
		boolToGboolean(isImportant)

	C.gtk_action_set_is_important((*C.GtkAction)(recv.native), c_is_important)

	return
}

// Sets the label of @action.
/*

C function

gtk_action_set_label
*/
func (recv *Action) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_action_set_label((*C.GtkAction)(recv.native), c_label)

	return
}

// Sets a shorter label text on @action.
/*

C function

gtk_action_set_short_label
*/
func (recv *Action) SetShortLabel(shortLabel string) {
	c_short_label := C.CString(shortLabel)
	defer C.free(unsafe.Pointer(c_short_label))

	C.gtk_action_set_short_label((*C.GtkAction)(recv.native), c_short_label)

	return
}

// Sets the stock id on @action
/*

C function

gtk_action_set_stock_id
*/
func (recv *Action) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_action_set_stock_id((*C.GtkAction)(recv.native), c_stock_id)

	return
}

// Sets the tooltip text on @action
/*

C function

gtk_action_set_tooltip
*/
func (recv *Action) SetTooltip(tooltip string) {
	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_action_set_tooltip((*C.GtkAction)(recv.native), c_tooltip)

	return
}

// Sets whether @action is visible when horizontal
/*

C function

gtk_action_set_visible_horizontal
*/
func (recv *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	c_visible_horizontal :=
		boolToGboolean(visibleHorizontal)

	C.gtk_action_set_visible_horizontal((*C.GtkAction)(recv.native), c_visible_horizontal)

	return
}

// Sets whether @action is visible when vertical
/*

C function

gtk_action_set_visible_vertical
*/
func (recv *Action) SetVisibleVertical(visibleVertical bool) {
	c_visible_vertical :=
		boolToGboolean(visibleVertical)

	C.gtk_action_set_visible_vertical((*C.GtkAction)(recv.native), c_visible_vertical)

	return
}

// Reenable activation signals from the action
/*

C function

gtk_action_unblock_activate
*/
func (recv *Action) UnblockActivate() {
	C.gtk_action_unblock_activate((*C.GtkAction)(recv.native))

	return
}

// Returns the model for @cell_view. If no model is used %NULL is
// returned.
/*

C function

gtk_cell_view_get_model
*/
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

// Unsupported signal 'icon-press' for Entry : unsupported parameter icon_pos : type EntryIconPosition :

// Unsupported signal 'icon-release' for Entry : unsupported parameter icon_pos : type EntryIconPosition :

// Returns the index of the icon which is the source of the current
// DND operation, or -1.
//
// This function is meant to be used in a #GtkWidget::drag-data-get
// callback.
/*

C function

gtk_entry_get_current_icon_drag_source
*/
func (recv *Entry) GetCurrentIconDragSource() int32 {
	retC := C.gtk_entry_get_current_icon_drag_source((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the icon is activatable.
/*

C function

gtk_entry_get_icon_activatable
*/
func (recv *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_activatable((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := retC == C.TRUE

	return retGo
}

// Finds the icon at the given position and return its index. The
// position’s coordinates are relative to the @entry’s top left corner.
// If @x, @y doesn’t lie inside an icon, -1 is returned.
// This function is intended for use in a #GtkWidget::query-tooltip
// signal handler.
/*

C function

gtk_entry_get_icon_at_pos
*/
func (recv *Entry) GetIconAtPos(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_entry_get_icon_at_pos((*C.GtkEntry)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the #GIcon used for the icon, or %NULL if there is
// no icon or if the icon was set by some other method (e.g., by
// stock, pixbuf, or icon name).
/*

C function

gtk_entry_get_icon_gicon
*/
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

// Retrieves the icon name used for the icon, or %NULL if there is
// no icon or if the icon was set by some other method (e.g., by
// pixbuf, stock or gicon).
/*

C function

gtk_entry_get_icon_name
*/
func (recv *Entry) GetIconName(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_name((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the image used for the icon.
//
// Unlike the other methods of setting and getting icon data, this
// method will work regardless of whether the icon was set using a
// #GdkPixbuf, a #GIcon, a stock item, or an icon name.
/*

C function

gtk_entry_get_icon_pixbuf
*/
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

// Returns whether the icon appears sensitive or insensitive.
/*

C function

gtk_entry_get_icon_sensitive
*/
func (recv *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_sensitive((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the stock id used for the icon, or %NULL if there is
// no icon or if the icon was set by some other method (e.g., by
// pixbuf, icon name or gicon).
/*

C function

gtk_entry_get_icon_stock
*/
func (recv *Entry) GetIconStock(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_stock((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the type of representation being used by the icon
// to store image data. If the icon has no image data,
// the return value will be %GTK_IMAGE_EMPTY.
/*

C function

gtk_entry_get_icon_storage_type
*/
func (recv *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_storage_type((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := (ImageType)(retC)

	return retGo
}

// Gets the contents of the tooltip on the icon at the specified
// position in @entry.
/*

C function

gtk_entry_get_icon_tooltip_markup
*/
func (recv *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_tooltip_markup((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the contents of the tooltip on the icon at the specified
// position in @entry.
/*

C function

gtk_entry_get_icon_tooltip_text
*/
func (recv *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	retC := C.gtk_entry_get_icon_tooltip_text((*C.GtkEntry)(recv.native), c_icon_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns the current fraction of the task that’s been completed.
// See gtk_entry_set_progress_fraction().
/*

C function

gtk_entry_get_progress_fraction
*/
func (recv *Entry) GetProgressFraction() float64 {
	retC := C.gtk_entry_get_progress_fraction((*C.GtkEntry)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Retrieves the pulse step set with gtk_entry_set_progress_pulse_step().
/*

C function

gtk_entry_get_progress_pulse_step
*/
func (recv *Entry) GetProgressPulseStep() float64 {
	retC := C.gtk_entry_get_progress_pulse_step((*C.GtkEntry)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Indicates that some progress is made, but you don’t know how much.
// Causes the entry’s progress indicator to enter “activity mode,”
// where a block bounces back and forth. Each call to
// gtk_entry_progress_pulse() causes the block to move by a little bit
// (the amount of movement per pulse is determined by
// gtk_entry_set_progress_pulse_step()).
/*

C function

gtk_entry_progress_pulse
*/
func (recv *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse((*C.GtkEntry)(recv.native))

	return
}

// Sets whether the icon is activatable.
/*

C function

gtk_entry_set_icon_activatable
*/
func (recv *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_activatable :=
		boolToGboolean(activatable)

	C.gtk_entry_set_icon_activatable((*C.GtkEntry)(recv.native), c_icon_pos, c_activatable)

	return
}

// Sets up the icon at the given position so that GTK+ will start a drag
// operation when the user clicks and drags the icon.
//
// To handle the drag operation, you need to connect to the usual
// #GtkWidget::drag-data-get (or possibly #GtkWidget::drag-data-delete)
// signal, and use gtk_entry_get_current_icon_drag_source() in
// your signal handler to find out if the drag was started from
// an icon.
//
// By default, GTK+ uses the icon as the drag icon. You can use the
// #GtkWidget::drag-begin signal to set a different icon. Note that you
// have to use g_signal_connect_after() to ensure that your signal handler
// gets executed after the default handler.
/*

C function

gtk_entry_set_icon_drag_source
*/
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

// Sets the icon shown in the entry at the specified position
// from the current icon theme.
// If the icon isn’t known, a “broken image” icon will be displayed
// instead.
//
// If @icon is %NULL, no icon will be shown in the specified position.
/*

C function

gtk_entry_set_icon_from_gicon
*/
func (recv *Entry) SetIconFromGicon(iconPos EntryIconPosition, icon *gio.Icon) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_entry_set_icon_from_gicon((*C.GtkEntry)(recv.native), c_icon_pos, c_icon)

	return
}

// Sets the icon shown in the entry at the specified position
// from the current icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead.
//
// If @icon_name is %NULL, no icon will be shown in the specified position.
/*

C function

gtk_entry_set_icon_from_icon_name
*/
func (recv *Entry) SetIconFromIconName(iconPos EntryIconPosition, iconName string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_entry_set_icon_from_icon_name((*C.GtkEntry)(recv.native), c_icon_pos, c_icon_name)

	return
}

// Sets the icon shown in the specified position using a pixbuf.
//
// If @pixbuf is %NULL, no icon will be shown in the specified position.
/*

C function

gtk_entry_set_icon_from_pixbuf
*/
func (recv *Entry) SetIconFromPixbuf(iconPos EntryIconPosition, pixbuf *gdkpixbuf.Pixbuf) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_entry_set_icon_from_pixbuf((*C.GtkEntry)(recv.native), c_icon_pos, c_pixbuf)

	return
}

// Sets the icon shown in the entry at the specified position from
// a stock image.
//
// If @stock_id is %NULL, no icon will be shown in the specified position.
/*

C function

gtk_entry_set_icon_from_stock
*/
func (recv *Entry) SetIconFromStock(iconPos EntryIconPosition, stockId string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_entry_set_icon_from_stock((*C.GtkEntry)(recv.native), c_icon_pos, c_stock_id)

	return
}

// Sets the sensitivity for the specified icon.
/*

C function

gtk_entry_set_icon_sensitive
*/
func (recv *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_entry_set_icon_sensitive((*C.GtkEntry)(recv.native), c_icon_pos, c_sensitive)

	return
}

// Sets @tooltip as the contents of the tooltip for the icon at
// the specified position. @tooltip is assumed to be marked up with
// the [Pango text markup language][PangoMarkupFormat].
//
// Use %NULL for @tooltip to remove an existing tooltip.
//
// See also gtk_widget_set_tooltip_markup() and
// gtk_entry_set_icon_tooltip_text().
/*

C function

gtk_entry_set_icon_tooltip_markup
*/
func (recv *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_entry_set_icon_tooltip_markup((*C.GtkEntry)(recv.native), c_icon_pos, c_tooltip)

	return
}

// Sets @tooltip as the contents of the tooltip for the icon
// at the specified position.
//
// Use %NULL for @tooltip to remove an existing tooltip.
//
// See also gtk_widget_set_tooltip_text() and
// gtk_entry_set_icon_tooltip_markup().
//
// If you unset the widget tooltip via gtk_widget_set_tooltip_text() or
// gtk_widget_set_tooltip_markup(), this sets GtkWidget:has-tooltip to %FALSE,
// which suppresses icon tooltips too. You can resolve this by then calling
// gtk_widget_set_has_tooltip() to set GtkWidget:has-tooltip back to %TRUE, or
// setting at least one non-empty tooltip on any icon achieves the same result.
/*

C function

gtk_entry_set_icon_tooltip_text
*/
func (recv *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	C.gtk_entry_set_icon_tooltip_text((*C.GtkEntry)(recv.native), c_icon_pos, c_tooltip)

	return
}

// Causes the entry’s progress indicator to “fill in” the given
// fraction of the bar. The fraction should be between 0.0 and 1.0,
// inclusive.
/*

C function

gtk_entry_set_progress_fraction
*/
func (recv *Entry) SetProgressFraction(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_entry_set_progress_fraction((*C.GtkEntry)(recv.native), c_fraction)

	return
}

// Sets the fraction of total entry width to move the progress
// bouncing block for each call to gtk_entry_progress_pulse().
/*

C function

gtk_entry_set_progress_pulse_step
*/
func (recv *Entry) SetProgressPulseStep(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_entry_set_progress_pulse_step((*C.GtkEntry)(recv.native), c_fraction)

	return
}

// Unsets the invisible char previously set with
// gtk_entry_set_invisible_char(). So that the
// default invisible char is used again.
/*

C function

gtk_entry_unset_invisible_char
*/
func (recv *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char((*C.GtkEntry)(recv.native))

	return
}

// Gets the id of the currently active slave of the @context.
/*

C function

gtk_im_multicontext_get_context_id
*/
func (recv *IMMulticontext) GetContextId() string {
	retC := C.gtk_im_multicontext_get_context_id((*C.GtkIMMulticontext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the context id for @context.
//
// This causes the currently active slave of @context to be
// replaced by the slave corresponding to the new context id.
/*

C function

gtk_im_multicontext_set_context_id
*/
func (recv *IMMulticontext) SetContextId(contextId string) {
	c_context_id := C.CString(contextId)
	defer C.free(unsafe.Pointer(c_context_id))

	C.gtk_im_multicontext_set_context_id((*C.GtkIMMulticontext)(recv.native), c_context_id)

	return
}

// Returns whether the menu item will ignore the #GtkSettings:gtk-menu-images
// setting and always show the image, if available.
/*

C function

gtk_image_menu_item_get_always_show_image
*/
func (recv *ImageMenuItem) GetAlwaysShowImage() bool {
	retC := C.gtk_image_menu_item_get_always_show_image((*C.GtkImageMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether the label set in the menuitem is used as a
// stock id to select the stock item for the item.
/*

C function

gtk_image_menu_item_get_use_stock
*/
func (recv *ImageMenuItem) GetUseStock() bool {
	retC := C.gtk_image_menu_item_get_use_stock((*C.GtkImageMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Specifies an @accel_group to add the menu items accelerator to
// (this only applies to stock items so a stock item must already
// be set, make sure to call gtk_image_menu_item_set_use_stock()
// and gtk_menu_item_set_label() with a valid stock item first).
//
// If you want this menu item to have changeable accelerators then
// you shouldnt need this (see gtk_image_menu_item_new_from_stock()).
/*

C function

gtk_image_menu_item_set_accel_group
*/
func (recv *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_image_menu_item_set_accel_group((*C.GtkImageMenuItem)(recv.native), c_accel_group)

	return
}

// If %TRUE, the menu item will ignore the #GtkSettings:gtk-menu-images
// setting and always show the image, if available.
//
// Use this property if the menuitem would be useless or hard to use
// without the image.
/*

C function

gtk_image_menu_item_set_always_show_image
*/
func (recv *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_image_menu_item_set_always_show_image((*C.GtkImageMenuItem)(recv.native), c_always_show)

	return
}

// If %TRUE, the label set in the menuitem is used as a
// stock id to select the stock item for the item.
/*

C function

gtk_image_menu_item_set_use_stock
*/
func (recv *ImageMenuItem) SetUseStock(useStock bool) {
	c_use_stock :=
		boolToGboolean(useStock)

	C.gtk_image_menu_item_set_use_stock((*C.GtkImageMenuItem)(recv.native), c_use_stock)

	return
}

// Sets @text on the @menu_item label
/*

C function

gtk_menu_item_get_label
*/
func (recv *MenuItem) GetLabel() string {
	retC := C.gtk_menu_item_get_label((*C.GtkMenuItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks if an underline in the text indicates the next character
// should be used for the mnemonic accelerator key.
/*

C function

gtk_menu_item_get_use_underline
*/
func (recv *MenuItem) GetUseUnderline() bool {
	retC := C.gtk_menu_item_get_use_underline((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @text on the @menu_item label
/*

C function

gtk_menu_item_set_label
*/
func (recv *MenuItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_menu_item_set_label((*C.GtkMenuItem)(recv.native), c_label)

	return
}

// If true, an underline in the text indicates the next character
// should be used for the mnemonic accelerator key.
/*

C function

gtk_menu_item_set_use_underline
*/
func (recv *MenuItem) SetUseUnderline(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_menu_item_set_use_underline((*C.GtkMenuItem)(recv.native), c_setting)

	return
}

// Signalize that drawing of particular page is complete.
//
// It is called after completion of page drawing (e.g. drawing in another
// thread).
// If gtk_print_operation_set_defer_drawing() was called before, then this function
// has to be called by application. In another case it is called by the library
// itself.
/*

C function

gtk_print_operation_draw_page_finish
*/
func (recv *PrintOperation) DrawPageFinish() {
	C.gtk_print_operation_draw_page_finish((*C.GtkPrintOperation)(recv.native))

	return
}

// Sets up the #GtkPrintOperation to wait for calling of
// gtk_print_operation_draw_page_finish() from application. It can
// be used for drawing page in another thread.
//
// This function must be called in the callback of “draw-page” signal.
/*

C function

gtk_print_operation_set_defer_drawing
*/
func (recv *PrintOperation) SetDeferDrawing() {
	C.gtk_print_operation_set_defer_drawing((*C.GtkPrintOperation)(recv.native))

	return
}

// Gets the value of %GTK_PRINT_SETTINGS_PRINTER_LPI.
/*

C function

gtk_print_settings_get_printer_lpi
*/
func (recv *PrintSettings) GetPrinterLpi() float64 {
	retC := C.gtk_print_settings_get_printer_lpi((*C.GtkPrintSettings)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Gets the value of %GTK_PRINT_SETTINGS_RESOLUTION_X.
/*

C function

gtk_print_settings_get_resolution_x
*/
func (recv *PrintSettings) GetResolutionX() int32 {
	retC := C.gtk_print_settings_get_resolution_x((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of %GTK_PRINT_SETTINGS_RESOLUTION_Y.
/*

C function

gtk_print_settings_get_resolution_y
*/
func (recv *PrintSettings) GetResolutionY() int32 {
	retC := C.gtk_print_settings_get_resolution_y((*C.GtkPrintSettings)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the value of %GTK_PRINT_SETTINGS_PRINTER_LPI.
/*

C function

gtk_print_settings_set_printer_lpi
*/
func (recv *PrintSettings) SetPrinterLpi(lpi float64) {
	c_lpi := (C.gdouble)(lpi)

	C.gtk_print_settings_set_printer_lpi((*C.GtkPrintSettings)(recv.native), c_lpi)

	return
}

// Sets the values of %GTK_PRINT_SETTINGS_RESOLUTION,
// %GTK_PRINT_SETTINGS_RESOLUTION_X and
// %GTK_PRINT_SETTINGS_RESOLUTION_Y.
/*

C function

gtk_print_settings_set_resolution_xy
*/
func (recv *PrintSettings) SetResolutionXy(resolutionX int32, resolutionY int32) {
	c_resolution_x := (C.gint)(resolutionX)

	c_resolution_y := (C.gint)(resolutionY)

	C.gtk_print_settings_set_resolution_xy((*C.GtkPrintSettings)(recv.native), c_resolution_x, c_resolution_y)

	return
}

// Adds a mark at @value.
//
// A mark is indicated visually by drawing a tick mark next to the scale,
// and GTK+ makes it easy for the user to position the scale exactly at the
// marks value.
//
// If @markup is not %NULL, text is shown next to the tick mark.
//
// To remove marks from a scale, use gtk_scale_clear_marks().
/*

C function

gtk_scale_add_mark
*/
func (recv *Scale) AddMark(value float64, position PositionType, markup string) {
	c_value := (C.gdouble)(value)

	c_position := (C.GtkPositionType)(position)

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_scale_add_mark((*C.GtkScale)(recv.native), c_value, c_position, c_markup)

	return
}

// Removes any marks that have been added with gtk_scale_add_mark().
/*

C function

gtk_scale_clear_marks
*/
func (recv *Scale) ClearMarks() {
	C.gtk_scale_clear_marks((*C.GtkScale)(recv.native))

	return
}

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
	event := gdk.EventScrollNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconScrollEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Returns the current value of the has-tooltip property.
// See #GtkStatusIcon:has-tooltip for more information.
/*

C function

gtk_status_icon_get_has_tooltip
*/
func (recv *StatusIcon) GetHasTooltip() bool {
	retC := C.gtk_status_icon_get_has_tooltip((*C.GtkStatusIcon)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the contents of the tooltip for @status_icon.
/*

C function

gtk_status_icon_get_tooltip_markup
*/
func (recv *StatusIcon) GetTooltipMarkup() string {
	retC := C.gtk_status_icon_get_tooltip_markup((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the contents of the tooltip for @status_icon.
/*

C function

gtk_status_icon_get_tooltip_text
*/
func (recv *StatusIcon) GetTooltipText() string {
	retC := C.gtk_status_icon_get_tooltip_text((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Sets the has-tooltip property on @status_icon to @has_tooltip.
// See #GtkStatusIcon:has-tooltip for more information.
/*

C function

gtk_status_icon_set_has_tooltip
*/
func (recv *StatusIcon) SetHasTooltip(hasTooltip bool) {
	c_has_tooltip :=
		boolToGboolean(hasTooltip)

	C.gtk_status_icon_set_has_tooltip((*C.GtkStatusIcon)(recv.native), c_has_tooltip)

	return
}

// Sets @markup as the contents of the tooltip, which is marked up with
// the [Pango text markup language][PangoMarkupFormat].
//
// This function will take care of setting #GtkStatusIcon:has-tooltip to %TRUE
// and of the default handler for the #GtkStatusIcon::query-tooltip signal.
//
// See also the #GtkStatusIcon:tooltip-markup property and
// gtk_tooltip_set_markup().
/*

C function

gtk_status_icon_set_tooltip_markup
*/
func (recv *StatusIcon) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_status_icon_set_tooltip_markup((*C.GtkStatusIcon)(recv.native), c_markup)

	return
}

// Sets @text as the contents of the tooltip.
//
// This function will take care of setting #GtkStatusIcon:has-tooltip to
// %TRUE and of the default handler for the #GtkStatusIcon::query-tooltip
// signal.
//
// See also the #GtkStatusIcon:tooltip-text property and
// gtk_tooltip_set_text().
/*

C function

gtk_status_icon_set_tooltip_text
*/
func (recv *StatusIcon) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_status_icon_set_tooltip_text((*C.GtkStatusIcon)(recv.native), c_text)

	return
}

// Unsupported : gtk_style_get : unsupported parameter ... : varargs

// Queries the value of a style property corresponding to a
// widget class is in the given style.
/*

C function

gtk_style_get_style_property
*/
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
	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	index := int(uintptr(data))
	callback := signalTextBufferPasteDoneMap[index].callback
	callback(clipboard)
}
