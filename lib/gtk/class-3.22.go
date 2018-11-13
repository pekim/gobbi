// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_clipboard_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_to_gvariant : return type : Blacklisted record : GVariant

// Retrieves the value set by gtk_gl_area_set_use_es().
/*

C function

gtk_gl_area_get_use_es
*/
func (recv *GLArea) GetUseEs() bool {
	retC := C.gtk_gl_area_get_use_es((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the @area should create an OpenGL or an OpenGL ES context.
//
// You should check the capabilities of the #GdkGLContext before drawing
// with either API.
/*

C function

gtk_gl_area_set_use_es
*/
func (recv *GLArea) SetUseEs(useEs bool) {
	c_use_es :=
		boolToGboolean(useEs)

	C.gtk_gl_area_set_use_es((*C.GtkGLArea)(recv.native), c_use_es)

	return
}

// Unsupported signal 'popped-up' for Menu : unsupported parameter flipped_rect : type gpointer :

// Places @menu on the given monitor.
/*

C function

gtk_menu_place_on_monitor
*/
func (recv *Menu) PlaceOnMonitor(monitor *gdk.Monitor) {
	c_monitor := (*C.GdkMonitor)(C.NULL)
	if monitor != nil {
		c_monitor = (*C.GdkMonitor)(monitor.ToC())
	}

	C.gtk_menu_place_on_monitor((*C.GtkMenu)(recv.native), c_monitor)

	return
}

// Unsupported : gtk_menu_popup_at_pointer : unsupported parameter trigger_event : no type generator for Gdk.Event (const GdkEvent*) for param trigger_event

// Unsupported : gtk_menu_popup_at_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_menu_popup_at_widget : unsupported parameter trigger_event : no type generator for Gdk.Event (const GdkEvent*) for param trigger_event

// Creates a new #GtkPadController that will associate events from @pad to
// actions. A %NULL pad may be provided so the controller manages all pad devices
// generically, it is discouraged to mix #GtkPadController objects with %NULL
// and non-%NULL @pad argument on the same @window, as execution order is not
// guaranteed.
//
// The #GtkPadController is created with no mapped actions. In order to map pad
// events to actions, use gtk_pad_controller_set_action_entries() or
// gtk_pad_controller_set_action().
/*

C function

gtk_pad_controller_new
*/
func PadControllerNew(window *Window, group *gio.ActionGroup, pad *gdk.Device) *PadController {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	c_group := (*C.GActionGroup)(group.ToC())

	c_pad := (*C.GdkDevice)(C.NULL)
	if pad != nil {
		c_pad = (*C.GdkDevice)(pad.ToC())
	}

	retC := C.gtk_pad_controller_new(c_window, c_group, c_pad)
	retGo := PadControllerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds an individual action to @controller. This action will only be activated
// if the given button/ring/strip number in @index is interacted while
// the current mode is @mode. -1 may be used for simple cases, so the action
// is triggered on all modes.
//
// The given @label should be considered user-visible, so internationalization
// rules apply. Some windowing systems may be able to use those for user
// feedback.
/*

C function

gtk_pad_controller_set_action
*/
func (recv *PadController) SetAction(type_ PadActionType, index int32, mode int32, label string, actionName string) {
	c_type := (C.GtkPadActionType)(type_)

	c_index := (C.gint)(index)

	c_mode := (C.gint)(mode)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.gtk_pad_controller_set_action((*C.GtkPadController)(recv.native), c_type, c_index, c_mode, c_label, c_action_name)

	return
}

// Unsupported : gtk_pad_controller_set_action_entries : unsupported parameter entries :

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_to_gvariant : return type : Blacklisted record : GVariant

// Pops @popover down.This is different than a gtk_widget_hide() call
// in that it shows the popover with a transition. If you want to hide
// the popover without a transition, use gtk_widget_hide().
/*

C function

gtk_popover_popdown
*/
func (recv *Popover) Popdown() {
	C.gtk_popover_popdown((*C.GtkPopover)(recv.native))

	return
}

// Pops @popover up. This is different than a gtk_widget_show() call
// in that it shows the popover with a transition. If you want to show
// the popover without a transition, use gtk_widget_show().
/*

C function

gtk_popover_popup
*/
func (recv *Popover) Popup() {
	C.gtk_popover_popup((*C.GtkPopover)(recv.native))

	return
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// Returns the maximum content height set.
/*

C function

gtk_scrolled_window_get_max_content_height
*/
func (recv *ScrolledWindow) GetMaxContentHeight() int32 {
	retC := C.gtk_scrolled_window_get_max_content_height((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the maximum content width set.
/*

C function

gtk_scrolled_window_get_max_content_width
*/
func (recv *ScrolledWindow) GetMaxContentWidth() int32 {
	retC := C.gtk_scrolled_window_get_max_content_width((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Reports whether the natural height of the child will be calculated and propagated
// through the scrolled windows requested natural height.
/*

C function

gtk_scrolled_window_get_propagate_natural_height
*/
func (recv *ScrolledWindow) GetPropagateNaturalHeight() bool {
	retC := C.gtk_scrolled_window_get_propagate_natural_height((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Reports whether the natural width of the child will be calculated and propagated
// through the scrolled windows requested natural width.
/*

C function

gtk_scrolled_window_get_propagate_natural_width
*/
func (recv *ScrolledWindow) GetPropagateNaturalWidth() bool {
	retC := C.gtk_scrolled_window_get_propagate_natural_width((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the maximum height that @scrolled_window should keep visible. The
// @scrolled_window will grow up to this height before it starts scrolling
// the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than #GtkScrolledWindow:min-content-height.
/*

C function

gtk_scrolled_window_set_max_content_height
*/
func (recv *ScrolledWindow) SetMaxContentHeight(height int32) {
	c_height := (C.gint)(height)

	C.gtk_scrolled_window_set_max_content_height((*C.GtkScrolledWindow)(recv.native), c_height)

	return
}

// Sets the maximum width that @scrolled_window should keep visible. The
// @scrolled_window will grow up to this width before it starts scrolling
// the content.
//
// It is a programming error to set the maximum content width to a value
// smaller than #GtkScrolledWindow:min-content-width.
/*

C function

gtk_scrolled_window_set_max_content_width
*/
func (recv *ScrolledWindow) SetMaxContentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_scrolled_window_set_max_content_width((*C.GtkScrolledWindow)(recv.native), c_width)

	return
}

// Sets whether the natural height of the child should be calculated and propagated
// through the scrolled windows requested natural height.
/*

C function

gtk_scrolled_window_set_propagate_natural_height
*/
func (recv *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	c_propagate :=
		boolToGboolean(propagate)

	C.gtk_scrolled_window_set_propagate_natural_height((*C.GtkScrolledWindow)(recv.native), c_propagate)

	return
}

// Sets whether the natural width of the child should be calculated and propagated
// through the scrolled windows requested natural width.
/*

C function

gtk_scrolled_window_set_propagate_natural_width
*/
func (recv *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	c_propagate :=
		boolToGboolean(propagate)

	C.gtk_scrolled_window_set_propagate_natural_width((*C.GtkScrolledWindow)(recv.native), c_propagate)

	return
}

// Creates a new #GtkShortcutLabel with @accelerator set.
/*

C function

gtk_shortcut_label_new
*/
func ShortcutLabelNew(accelerator string) *ShortcutLabel {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	retC := C.gtk_shortcut_label_new(c_accelerator)
	retGo := ShortcutLabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the current accelerator of @self.
/*

C function

gtk_shortcut_label_get_accelerator
*/
func (recv *ShortcutLabel) GetAccelerator() string {
	retC := C.gtk_shortcut_label_get_accelerator((*C.GtkShortcutLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the text that is displayed when no accelerator is set.
/*

C function

gtk_shortcut_label_get_disabled_text
*/
func (recv *ShortcutLabel) GetDisabledText() string {
	retC := C.gtk_shortcut_label_get_disabled_text((*C.GtkShortcutLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the accelerator to be displayed by @self.
/*

C function

gtk_shortcut_label_set_accelerator
*/
func (recv *ShortcutLabel) SetAccelerator(accelerator string) {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	C.gtk_shortcut_label_set_accelerator((*C.GtkShortcutLabel)(recv.native), c_accelerator)

	return
}

// Sets the text to be displayed by @self when no accelerator is set.
/*

C function

gtk_shortcut_label_set_disabled_text
*/
func (recv *ShortcutLabel) SetDisabledText(disabledText string) {
	c_disabled_text := C.CString(disabledText)
	defer C.free(unsafe.Pointer(c_disabled_text))

	C.gtk_shortcut_label_set_disabled_text((*C.GtkShortcutLabel)(recv.native), c_disabled_text)

	return
}
