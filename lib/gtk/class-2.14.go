// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
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

	gboolean statusicon_buttonPressEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-press-event", G_CALLBACK(statusicon_buttonPressEventHandler), data);
	}

*/
/*

	gboolean statusicon_buttonReleaseEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-release-event", G_CALLBACK(statusicon_buttonReleaseEventHandler), data);
	}

*/
/*

	gboolean widget_damageEventHandler(GObject *, GdkEventExpose *, gpointer);

	static gulong Widget_signal_connect_damage_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "damage-event", G_CALLBACK(widget_damageEventHandler), data);
	}

*/
import "C"

// Locks are added and removed using gtk_accel_group_lock() and
// gtk_accel_group_unlock().
/*

C function

gtk_accel_group_get_is_locked
*/
func (recv *AccelGroup) GetIsLocked() bool {
	retC := C.gtk_accel_group_get_is_locked((*C.GtkAccelGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets a #GdkModifierType representing the mask for this
// @accel_group. For example, #GDK_CONTROL_MASK, #GDK_SHIFT_MASK, etc.
/*

C function

gtk_accel_group_get_modifier_mask
*/
func (recv *AccelGroup) GetModifierMask() gdk.ModifierType {
	retC := C.gtk_accel_group_get_modifier_mask((*C.GtkAccelGroup)(recv.native))
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// Sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the
// #GtkAdjustment::changed signal. See gtk_adjustment_set_lower()
// for an alternative way of compressing multiple emissions of
// #GtkAdjustment::changed into one.
/*

C function

gtk_adjustment_configure
*/
func (recv *Adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_step_increment := (C.gdouble)(stepIncrement)

	c_page_increment := (C.gdouble)(pageIncrement)

	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_configure((*C.GtkAdjustment)(recv.native), c_value, c_lower, c_upper, c_step_increment, c_page_increment, c_page_size)

	return
}

// Retrieves the minimum value of the adjustment.
/*

C function

gtk_adjustment_get_lower
*/
func (recv *Adjustment) GetLower() float64 {
	retC := C.gtk_adjustment_get_lower((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Retrieves the page increment of the adjustment.
/*

C function

gtk_adjustment_get_page_increment
*/
func (recv *Adjustment) GetPageIncrement() float64 {
	retC := C.gtk_adjustment_get_page_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Retrieves the page size of the adjustment.
/*

C function

gtk_adjustment_get_page_size
*/
func (recv *Adjustment) GetPageSize() float64 {
	retC := C.gtk_adjustment_get_page_size((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Retrieves the step increment of the adjustment.
/*

C function

gtk_adjustment_get_step_increment
*/
func (recv *Adjustment) GetStepIncrement() float64 {
	retC := C.gtk_adjustment_get_step_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Retrieves the maximum value of the adjustment.
/*

C function

gtk_adjustment_get_upper
*/
func (recv *Adjustment) GetUpper() float64 {
	retC := C.gtk_adjustment_get_upper((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual
// setters, multiple #GtkAdjustment::changed signals will be emitted.
// However, since the emission of the #GtkAdjustment::changed signal
// is tied to the emission of the #GObject::notify signals of the changed
// properties, it’s possible to compress the #GtkAdjustment::changed
// signals into one by calling g_object_freeze_notify() and
// g_object_thaw_notify() around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties
// to change, or using gtk_adjustment_configure() has the same effect
// of compressing #GtkAdjustment::changed emissions.
/*

C function

gtk_adjustment_set_lower
*/
func (recv *Adjustment) SetLower(lower float64) {
	c_lower := (C.gdouble)(lower)

	C.gtk_adjustment_set_lower((*C.GtkAdjustment)(recv.native), c_lower)

	return
}

// Sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple
// emissions of the #GtkAdjustment::changed signal when setting
// multiple adjustment properties.
/*

C function

gtk_adjustment_set_page_increment
*/
func (recv *Adjustment) SetPageIncrement(pageIncrement float64) {
	c_page_increment := (C.gdouble)(pageIncrement)

	C.gtk_adjustment_set_page_increment((*C.GtkAdjustment)(recv.native), c_page_increment)

	return
}

// Sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple
// emissions of the GtkAdjustment::changed signal when setting
// multiple adjustment properties.
/*

C function

gtk_adjustment_set_page_size
*/
func (recv *Adjustment) SetPageSize(pageSize float64) {
	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_set_page_size((*C.GtkAdjustment)(recv.native), c_page_size)

	return
}

// Sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple
// emissions of the #GtkAdjustment::changed signal when setting
// multiple adjustment properties.
/*

C function

gtk_adjustment_set_step_increment
*/
func (recv *Adjustment) SetStepIncrement(stepIncrement float64) {
	c_step_increment := (C.gdouble)(stepIncrement)

	C.gtk_adjustment_set_step_increment((*C.GtkAdjustment)(recv.native), c_step_increment)

	return
}

// Sets the maximum value of the adjustment.
//
// Note that values will be restricted by `upper - page-size`
// if the page-size property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple
// emissions of the #GtkAdjustment::changed signal when setting
// multiple adjustment properties.
/*

C function

gtk_adjustment_set_upper
*/
func (recv *Adjustment) SetUpper(upper float64) {
	c_upper := (C.gdouble)(upper)

	C.gtk_adjustment_set_upper((*C.GtkAdjustment)(recv.native), c_upper)

	return
}

// Unsupported : gtk_builder_add_objects_from_file : unsupported parameter object_ids :

// Unsupported : gtk_builder_add_objects_from_string : unsupported parameter object_ids :

// Queries the height of detail cells, in rows.
// See #GtkCalendar:detail-width-chars.
/*

C function

gtk_calendar_get_detail_height_rows
*/
func (recv *Calendar) GetDetailHeightRows() int32 {
	retC := C.gtk_calendar_get_detail_height_rows((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries the width of detail cells, in characters.
// See #GtkCalendar:detail-width-chars.
/*

C function

gtk_calendar_get_detail_width_chars
*/
func (recv *Calendar) GetDetailWidthChars() int32 {
	retC := C.gtk_calendar_get_detail_width_chars((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc (GtkCalendarDetailFunc) for param func

// Updates the height of detail cells.
// See #GtkCalendar:detail-height-rows.
/*

C function

gtk_calendar_set_detail_height_rows
*/
func (recv *Calendar) SetDetailHeightRows(rows int32) {
	c_rows := (C.gint)(rows)

	C.gtk_calendar_set_detail_height_rows((*C.GtkCalendar)(recv.native), c_rows)

	return
}

// Updates the width of detail cells.
// See #GtkCalendar:detail-width-chars.
/*

C function

gtk_calendar_set_detail_width_chars
*/
func (recv *Calendar) SetDetailWidthChars(chars int32) {
	c_chars := (C.gint)(chars)

	C.gtk_calendar_set_detail_width_chars((*C.GtkCalendar)(recv.native), c_chars)

	return
}

// Unsupported signal 'changed' for CellRendererCombo : unsupported parameter path_string : type utf8 :

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc (GtkClipboardURIReceivedFunc) for param callback

// Unsupported : gtk_clipboard_wait_for_uris : no return type

// Test to see if there is a list of URIs available to be pasted
// This is done by requesting the TARGETS atom and checking
// if it contains the URI targets. This function
// waits for the data to be received using the main loop, so events,
// timeouts, etc, may be dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_uris() since it doesn’t need to retrieve
// the actual URI data.
/*

C function

gtk_clipboard_wait_is_uris_available
*/
func (recv *Clipboard) WaitIsUrisAvailable() bool {
	retC := C.gtk_clipboard_wait_is_uris_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the #GtkColorSelection widget embedded in the dialog.
/*

C function

gtk_color_selection_dialog_get_color_selection
*/
func (recv *ColorSelectionDialog) GetColorSelection() *Widget {
	retC := C.gtk_color_selection_dialog_get_color_selection((*C.GtkColorSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
/*

C function

gtk_combo_box_get_button_sensitivity
*/
func (recv *ComboBox) GetButtonSensitivity() SensitivityType {
	retC := C.gtk_combo_box_get_button_sensitivity((*C.GtkComboBox)(recv.native))
	retGo := (SensitivityType)(retC)

	return retGo
}

// Sets whether the dropdown button of the combo box should be
// always sensitive (%GTK_SENSITIVITY_ON), never sensitive (%GTK_SENSITIVITY_OFF)
// or only if there is at least one item to display (%GTK_SENSITIVITY_AUTO).
/*

C function

gtk_combo_box_set_button_sensitivity
*/
func (recv *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	c_sensitivity := (C.GtkSensitivityType)(sensitivity)

	C.gtk_combo_box_set_button_sensitivity((*C.GtkComboBox)(recv.native), c_sensitivity)

	return
}

// Returns the current focus child widget inside @container. This is not the
// currently focused widget. That can be obtained by calling
// gtk_window_get_focus().
/*

C function

gtk_container_get_focus_child
*/
func (recv *Container) GetFocusChild() *Widget {
	retC := C.gtk_container_get_focus_child((*C.GtkContainer)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the action area of @dialog.
/*

C function

gtk_dialog_get_action_area
*/
func (recv *Dialog) GetActionArea() *Widget {
	retC := C.gtk_dialog_get_action_area((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the content area of @dialog.
/*

C function

gtk_dialog_get_content_area
*/
func (recv *Dialog) GetContentArea() *Box {
	retC := C.gtk_dialog_get_content_area((*C.GtkDialog)(recv.native))
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value set by gtk_entry_set_overwrite_mode().
/*

C function

gtk_entry_get_overwrite_mode
*/
func (recv *Entry) GetOverwriteMode() bool {
	retC := C.gtk_entry_get_overwrite_mode((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the current length of the text in
// @entry.
//
// This is equivalent to getting @entry's #GtkEntryBuffer and
// calling gtk_entry_buffer_get_length() on it.
/*

C function

gtk_entry_get_text_length
*/
func (recv *Entry) GetTextLength() uint16 {
	retC := C.gtk_entry_get_text_length((*C.GtkEntry)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Sets whether the text is overwritten when typing in the #GtkEntry.
/*

C function

gtk_entry_set_overwrite_mode
*/
func (recv *Entry) SetOverwriteMode(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_entry_set_overwrite_mode((*C.GtkEntry)(recv.native), c_overwrite)

	return
}

// Gets the #PangoFontFace representing the selected font group
// details (i.e. family, slant, weight, width, etc).
/*

C function

gtk_font_selection_get_face
*/
func (recv *FontSelection) GetFace() *pango.FontFace {
	retC := C.gtk_font_selection_get_face((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This returns the #GtkTreeView which lists all styles available for
// the selected font. For example, “Regular”, “Bold”, etc.
/*

C function

gtk_font_selection_get_face_list
*/
func (recv *FontSelection) GetFaceList() *Widget {
	retC := C.gtk_font_selection_get_face_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #PangoFontFamily representing the selected font family.
/*

C function

gtk_font_selection_get_family
*/
func (recv *FontSelection) GetFamily() *pango.FontFamily {
	retC := C.gtk_font_selection_get_family((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFamilyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This returns the #GtkTreeView that lists font families, for
// example, “Sans”, “Serif”, etc.
/*

C function

gtk_font_selection_get_family_list
*/
func (recv *FontSelection) GetFamilyList() *Widget {
	retC := C.gtk_font_selection_get_family_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This returns the #GtkEntry used to display the font as a preview.
/*

C function

gtk_font_selection_get_preview_entry
*/
func (recv *FontSelection) GetPreviewEntry() *Widget {
	retC := C.gtk_font_selection_get_preview_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// The selected font size.
/*

C function

gtk_font_selection_get_size
*/
func (recv *FontSelection) GetSize() int32 {
	retC := C.gtk_font_selection_get_size((*C.GtkFontSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// This returns the #GtkEntry used to allow the user to edit the font
// number manually instead of selecting it from the list of font sizes.
/*

C function

gtk_font_selection_get_size_entry
*/
func (recv *FontSelection) GetSizeEntry() *Widget {
	retC := C.gtk_font_selection_get_size_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This returns the #GtkTreeView used to list font sizes.
/*

C function

gtk_font_selection_get_size_list
*/
func (recv *FontSelection) GetSizeList() *Widget {
	retC := C.gtk_font_selection_get_size_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the “Cancel” button.
/*

C function

gtk_font_selection_dialog_get_cancel_button
*/
func (recv *FontSelectionDialog) GetCancelButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_cancel_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the “OK” button.
/*

C function

gtk_font_selection_dialog_get_ok_button
*/
func (recv *FontSelectionDialog) GetOkButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_ok_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new HSV color selector.
/*

C function

gtk_hsv_new
*/
func HSVNew() *HSV {
	retC := C.gtk_hsv_new()
	retGo := HSVNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries the current color in an HSV color selector.
// Returned values will be in the [0.0, 1.0] range.
/*

C function

gtk_hsv_get_color
*/
func (recv *HSV) GetColor() (float64, float64, float64) {
	var c_h C.gdouble

	var c_s C.gdouble

	var c_v C.gdouble

	C.gtk_hsv_get_color((*C.GtkHSV)(recv.native), &c_h, &c_s, &c_v)

	h := (float64)(c_h)

	s := (float64)(c_s)

	v := (float64)(c_v)

	return h, s, v
}

// Queries the size and ring width of an HSV color selector.
/*

C function

gtk_hsv_get_metrics
*/
func (recv *HSV) GetMetrics() (int32, int32) {
	var c_size C.gint

	var c_ring_width C.gint

	C.gtk_hsv_get_metrics((*C.GtkHSV)(recv.native), &c_size, &c_ring_width)

	size := (int32)(c_size)

	ringWidth := (int32)(c_ring_width)

	return size, ringWidth
}

// An HSV color selector can be said to be adjusting if multiple rapid
// changes are being made to its value, for example, when the user is
// adjusting the value with the mouse. This function queries whether
// the HSV color selector is being adjusted or not.
/*

C function

gtk_hsv_is_adjusting
*/
func (recv *HSV) IsAdjusting() bool {
	retC := C.gtk_hsv_is_adjusting((*C.GtkHSV)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the current color in an HSV color selector.
// Color component values must be in the [0.0, 1.0] range.
/*

C function

gtk_hsv_set_color
*/
func (recv *HSV) SetColor(h float64, s float64, v float64) {
	c_h := (C.double)(h)

	c_s := (C.double)(s)

	c_v := (C.double)(v)

	C.gtk_hsv_set_color((*C.GtkHSV)(recv.native), c_h, c_s, c_v)

	return
}

// Sets the size and ring width of an HSV color selector.
/*

C function

gtk_hsv_set_metrics
*/
func (recv *HSV) SetMetrics(size int32, ringWidth int32) {
	c_size := (C.gint)(size)

	c_ring_width := (C.gint)(ringWidth)

	C.gtk_hsv_set_metrics((*C.GtkHSV)(recv.native), c_size, c_ring_width)

	return
}

// Whether the handlebox’s child is currently detached.
/*

C function

gtk_handle_box_get_child_detached
*/
func (recv *HandleBox) GetChildDetached() bool {
	retC := C.gtk_handle_box_get_child_detached((*C.GtkHandleBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Creates a #GtkIconInfo for a #GdkPixbuf.
/*

C function

gtk_icon_info_new_for_pixbuf
*/
func IconInfoNewForPixbuf(iconTheme *IconTheme, pixbuf *gdkpixbuf.Pixbuf) *IconInfo {
	c_icon_theme := (*C.GtkIconTheme)(C.NULL)
	if iconTheme != nil {
		c_icon_theme = (*C.GtkIconTheme)(iconTheme.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_icon_info_new_for_pixbuf(c_icon_theme, c_pixbuf)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up an icon and returns a #GtkIconInfo containing information
// such as the filename of the icon. The icon can then be rendered
// into a pixbuf using gtk_icon_info_load_icon().
//
// When rendering on displays with high pixel densities you should not
// use a @size multiplied by the scaling factor returned by functions
// like gdk_window_get_scale_factor(). Instead, you should use
// gtk_icon_theme_lookup_by_gicon_for_scale(), as the assets loaded
// for a given scaling factor may be different.
/*

C function

gtk_icon_theme_lookup_by_gicon
*/
func (recv *IconTheme) LookupByGicon(icon *gio.Icon, size int32, flags IconLookupFlags) *IconInfo {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_by_gicon((*C.GtkIconTheme)(recv.native), c_icon, c_size, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Creates a #GtkImage displaying an icon from the current icon theme.
// If the icon name isn’t known, a “broken image” icon will be
// displayed instead.  If the current icon theme is changed, the icon
// will be updated appropriately.
/*

C function

gtk_image_new_from_gicon
*/
func ImageNewFromGicon(icon *gio.Icon, size IconSize) *Image {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_gicon(c_icon, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_gicon : unsupported parameter size : no type generator for gint (GtkIconSize*) for param size

// See gtk_image_new_from_gicon() for details.
/*

C function

gtk_image_set_from_gicon
*/
func (recv *Image) SetFromGicon(icon *gio.Icon, size IconSize) {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_gicon((*C.GtkImage)(recv.native), c_icon, c_size)

	return
}

// Retrieve the bin window of the layout used for drawing operations.
/*

C function

gtk_layout_get_bin_window
*/
func (recv *Layout) GetBinWindow() *gdk.Window {
	retC := C.gtk_layout_get_bin_window((*C.GtkLayout)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the “visited” state of the URI where the #GtkLinkButton
// points. The button becomes visited when it is clicked. If the URI
// is changed on the button, the “visited” state is unset again.
//
// The state may also be changed using gtk_link_button_set_visited().
/*

C function

gtk_link_button_get_visited
*/
func (recv *LinkButton) GetVisited() bool {
	retC := C.gtk_link_button_get_visited((*C.GtkLinkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the “visited” state of the URI where the #GtkLinkButton
// points.  See gtk_link_button_get_visited() for more details.
/*

C function

gtk_link_button_set_visited
*/
func (recv *LinkButton) SetVisited(visited bool) {
	c_visited :=
		boolToGboolean(visited)

	C.gtk_link_button_set_visited((*C.GtkLinkButton)(recv.native), c_visited)

	return
}

// Retrieves the accelerator path set on the menu.
/*

C function

gtk_menu_get_accel_path
*/
func (recv *Menu) GetAccelPath() string {
	retC := C.gtk_menu_get_accel_path((*C.GtkMenu)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the number of the monitor on which to show the menu.
/*

C function

gtk_menu_get_monitor
*/
func (recv *Menu) GetMonitor() int32 {
	retC := C.gtk_menu_get_monitor((*C.GtkMenu)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieve the accelerator path that was previously set on @menu_item.
//
// See gtk_menu_item_set_accel_path() for details.
/*

C function

gtk_menu_item_get_accel_path
*/
func (recv *MenuItem) GetAccelPath() string {
	retC := C.gtk_menu_item_get_accel_path((*C.GtkMenuItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the dialog’s image.
/*

C function

gtk_message_dialog_get_image
*/
func (recv *MessageDialog) GetImage() *Widget {
	retC := C.gtk_message_dialog_get_image((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkMountOperation
/*

C function

gtk_mount_operation_new
*/
func MountOperationNew(parent *Window) *MountOperation {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_mount_operation_new(c_parent)
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the transient parent used by the #GtkMountOperation
/*

C function

gtk_mount_operation_get_parent
*/
func (recv *MountOperation) GetParent() *Window {
	retC := C.gtk_mount_operation_get_parent((*C.GtkMountOperation)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the screen on which windows of the #GtkMountOperation
// will be shown.
/*

C function

gtk_mount_operation_get_screen
*/
func (recv *MountOperation) GetScreen() *gdk.Screen {
	retC := C.gtk_mount_operation_get_screen((*C.GtkMountOperation)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the #GtkMountOperation is currently displaying
// a window.
/*

C function

gtk_mount_operation_is_showing
*/
func (recv *MountOperation) IsShowing() bool {
	retC := C.gtk_mount_operation_is_showing((*C.GtkMountOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the transient parent for windows shown by the
// #GtkMountOperation.
/*

C function

gtk_mount_operation_set_parent
*/
func (recv *MountOperation) SetParent(parent *Window) {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	C.gtk_mount_operation_set_parent((*C.GtkMountOperation)(recv.native), c_parent)

	return
}

// Sets the screen to show windows of the #GtkMountOperation on.
/*

C function

gtk_mount_operation_set_screen
*/
func (recv *MountOperation) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_mount_operation_set_screen((*C.GtkMountOperation)(recv.native), c_screen)

	return
}

// Reads the page setup from the file @file_name.
// See gtk_page_setup_to_file().
/*

C function

gtk_page_setup_load_file
*/
func (recv *PageSetup) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads the page setup from the group @group_name in the key file
// @key_file.
/*

C function

gtk_page_setup_load_key_file
*/
func (recv *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the value of %GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
/*

C function

gtk_print_settings_get_number_up_layout
*/
func (recv *PrintSettings) GetNumberUpLayout() NumberUpLayout {
	retC := C.gtk_print_settings_get_number_up_layout((*C.GtkPrintSettings)(recv.native))
	retGo := (NumberUpLayout)(retC)

	return retGo
}

// Reads the print settings from @file_name. If the file could not be loaded
// then error is set to either a #GFileError or #GKeyFileError.
// See gtk_print_settings_to_file().
/*

C function

gtk_print_settings_load_file
*/
func (recv *PrintSettings) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads the print settings from the group @group_name in @key_file. If the
// file could not be loaded then error is set to either a #GFileError or
// #GKeyFileError.
/*

C function

gtk_print_settings_load_key_file
*/
func (recv *PrintSettings) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the value of %GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
/*

C function

gtk_print_settings_set_number_up_layout
*/
func (recv *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	c_number_up_layout := (C.GtkNumberUpLayout)(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout((*C.GtkPrintSettings)(recv.native), c_number_up_layout)

	return
}

// Retrieves the minus button of the #GtkScaleButton.
/*

C function

gtk_scale_button_get_minus_button
*/
func (recv *ScaleButton) GetMinusButton() *Button {
	retC := C.gtk_scale_button_get_minus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the plus button of the #GtkScaleButton.
/*

C function

gtk_scale_button_get_plus_button
*/
func (recv *ScaleButton) GetPlusButton() *Button {
	retC := C.gtk_scale_button_get_plus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the popup of the #GtkScaleButton.
/*

C function

gtk_scale_button_get_popup
*/
func (recv *ScaleButton) GetPopup() *Widget {
	retC := C.gtk_scale_button_get_popup((*C.GtkScaleButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalStatusIconButtonPressEventDetail struct {
	callback  StatusIconSignalButtonPressEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonPressEventId int
var signalStatusIconButtonPressEventMap = make(map[int]signalStatusIconButtonPressEventDetail)
var signalStatusIconButtonPressEventLock sync.Mutex

// StatusIconSignalButtonPressEventCallback is a callback function for a 'button-press-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonPressEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonPressEvent connects the callback to the 'button-press-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonPressEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonPressEvent(callback StatusIconSignalButtonPressEventCallback) int {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	signalStatusIconButtonPressEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_press_event(instance, C.gpointer(uintptr(signalStatusIconButtonPressEventId)))

	detail := signalStatusIconButtonPressEventDetail{callback, handlerID}
	signalStatusIconButtonPressEventMap[signalStatusIconButtonPressEventId] = detail

	return signalStatusIconButtonPressEventId
}

/*
DisconnectButtonPressEvent disconnects a callback from the 'button-press-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonPressEvent.
*/
func (recv *StatusIcon) DisconnectButtonPressEvent(connectionID int) {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	detail, exists := signalStatusIconButtonPressEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonPressEventMap, connectionID)
}

//export statusicon_buttonPressEventHandler
func statusicon_buttonPressEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonPressEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalStatusIconButtonReleaseEventDetail struct {
	callback  StatusIconSignalButtonReleaseEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonReleaseEventId int
var signalStatusIconButtonReleaseEventMap = make(map[int]signalStatusIconButtonReleaseEventDetail)
var signalStatusIconButtonReleaseEventLock sync.Mutex

// StatusIconSignalButtonReleaseEventCallback is a callback function for a 'button-release-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonReleaseEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonReleaseEvent connects the callback to the 'button-release-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonReleaseEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonReleaseEvent(callback StatusIconSignalButtonReleaseEventCallback) int {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	signalStatusIconButtonReleaseEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_release_event(instance, C.gpointer(uintptr(signalStatusIconButtonReleaseEventId)))

	detail := signalStatusIconButtonReleaseEventDetail{callback, handlerID}
	signalStatusIconButtonReleaseEventMap[signalStatusIconButtonReleaseEventId] = detail

	return signalStatusIconButtonReleaseEventId
}

/*
DisconnectButtonReleaseEvent disconnects a callback from the 'button-release-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonReleaseEvent.
*/
func (recv *StatusIcon) DisconnectButtonReleaseEvent(connectionID int) {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	detail, exists := signalStatusIconButtonReleaseEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonReleaseEventMap, connectionID)
}

//export statusicon_buttonReleaseEventHandler
func statusicon_buttonReleaseEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonReleaseEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Creates a status icon displaying a #GIcon. If the icon is a
// themed icon, it will be updated when the theme changes.
/*

C function

gtk_status_icon_new_from_gicon
*/
func StatusIconNewFromGicon(icon *gio.Icon) *StatusIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.gtk_status_icon_new_from_gicon(c_icon)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the #GIcon being displayed by the #GtkStatusIcon.
// The storage type of the status icon must be %GTK_IMAGE_EMPTY or
// %GTK_IMAGE_GICON (see gtk_status_icon_get_storage_type()).
// The caller of this function does not own a reference to the
// returned #GIcon.
//
// If this function fails, @icon is left unchanged;
/*

C function

gtk_status_icon_get_gicon
*/
func (recv *StatusIcon) GetGicon() *gio.Icon {
	retC := C.gtk_status_icon_get_gicon((*C.GtkStatusIcon)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// This function is only useful on the X11/freedesktop.org platform.
//
// It returns a window ID for the widget in the underlying
// status icon implementation.  This is useful for the Galago
// notification service, which can send a window ID in the protocol
// in order for the server to position notification windows
// pointing to a status icon reliably.
//
// This function is not intended for other use cases which are
// more likely to be met by one of the non-X11 specific methods, such
// as gtk_status_icon_position_menu().
/*

C function

gtk_status_icon_get_x11_window_id
*/
func (recv *StatusIcon) GetX11WindowId() uint32 {
	retC := C.gtk_status_icon_get_x11_window_id((*C.GtkStatusIcon)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Makes @status_icon display the #GIcon.
// See gtk_status_icon_new_from_gicon() for details.
/*

C function

gtk_status_icon_set_from_gicon
*/
func (recv *StatusIcon) SetFromGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_status_icon_set_from_gicon((*C.GtkStatusIcon)(recv.native), c_icon)

	return
}

// Emits the signal #GtkToolItem::toolbar_reconfigured on @tool_item.
// #GtkToolbar and other #GtkToolShell implementations use this function
// to notify children, when some aspect of their configuration changes.
/*

C function

gtk_tool_item_toolbar_reconfigured
*/
func (recv *ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured((*C.GtkToolItem)(recv.native))

	return
}

// Sets the icon of the tooltip (which is in front of the text) to be
// the icon indicated by @icon_name with the size indicated
// by @size.  If @icon_name is %NULL, the image will be hidden.
/*

C function

gtk_tooltip_set_icon_from_icon_name
*/
func (recv *Tooltip) SetIconFromIconName(iconName string, size IconSize) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	C.gtk_tooltip_set_icon_from_icon_name((*C.GtkTooltip)(recv.native), c_icon_name, c_size)

	return
}

// Unsupported : gtk_tree_selection_get_select_function : no return generator

type signalWidgetDamageEventDetail struct {
	callback  WidgetSignalDamageEventCallback
	handlerID C.gulong
}

var signalWidgetDamageEventId int
var signalWidgetDamageEventMap = make(map[int]signalWidgetDamageEventDetail)
var signalWidgetDamageEventLock sync.Mutex

// WidgetSignalDamageEventCallback is a callback function for a 'damage-event' signal emitted from a Widget.
type WidgetSignalDamageEventCallback func(event *gdk.EventExpose) bool

/*
ConnectDamageEvent connects the callback to the 'damage-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDamageEvent to remove it.
*/
func (recv *Widget) ConnectDamageEvent(callback WidgetSignalDamageEventCallback) int {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	signalWidgetDamageEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_damage_event(instance, C.gpointer(uintptr(signalWidgetDamageEventId)))

	detail := signalWidgetDamageEventDetail{callback, handlerID}
	signalWidgetDamageEventMap[signalWidgetDamageEventId] = detail

	return signalWidgetDamageEventId
}

/*
DisconnectDamageEvent disconnects a callback from the 'damage-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDamageEvent.
*/
func (recv *Widget) DisconnectDamageEvent(connectionID int) {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	detail, exists := signalWidgetDamageEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDamageEventMap, connectionID)
}

//export widget_damageEventHandler
func widget_damageEventHandler(_ *C.GObject, c_event *C.GdkEventExpose, data C.gpointer) C.gboolean {
	event := gdk.EventExposeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetDamageEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Returns the widget’s window if it is realized, %NULL otherwise
/*

C function

gtk_widget_get_window
*/
func (recv *Widget) GetWindow() *gdk.Window {
	retC := C.gtk_widget_get_window((*C.GtkWidget)(recv.native))
	var retGo (*gdk.Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the default widget for @window. See
// gtk_window_set_default() for more details.
/*

C function

gtk_window_get_default_widget
*/
func (recv *Window) GetDefaultWidget() *Widget {
	retC := C.gtk_window_get_default_widget((*C.GtkWindow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns a list of the #GtkWindows that belong to @window_group.
/*

C function

gtk_window_group_list_windows
*/
func (recv *WindowGroup) ListWindows() *glib.List {
	retC := C.gtk_window_group_list_windows((*C.GtkWindowGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}
