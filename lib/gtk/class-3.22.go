// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void menu_poppedUpHandler(GObject *, gpointer, gpointer, gboolean, gboolean, gpointer);

	static gulong Menu_signal_connect_popped_up(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popped-up", G_CALLBACK(menu_poppedUpHandler), data);
	}

*/
import "C"

// Unsupported : gtk_clipboard_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_to_gvariant : return type : Blacklisted record : GVariant

// GetUseEs is a wrapper around the C function gtk_gl_area_get_use_es.
func (recv *GLArea) GetUseEs() bool {
	retC := C.gtk_gl_area_get_use_es((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetUseEs is a wrapper around the C function gtk_gl_area_set_use_es.
func (recv *GLArea) SetUseEs(useEs bool) {
	c_use_es :=
		boolToGboolean(useEs)

	C.gtk_gl_area_set_use_es((*C.GtkGLArea)(recv.native), c_use_es)

	return
}

type signalMenuPoppedUpDetail struct {
	callback  MenuSignalPoppedUpCallback
	handlerID C.gulong
}

var signalMenuPoppedUpId int
var signalMenuPoppedUpMap = make(map[int]signalMenuPoppedUpDetail)
var signalMenuPoppedUpLock sync.RWMutex

// MenuSignalPoppedUpCallback is a callback function for a 'popped-up' signal emitted from a Menu.
type MenuSignalPoppedUpCallback func(flippedRect uintptr, finalRect uintptr, flippedX bool, flippedY bool)

/*
ConnectPoppedUp connects the callback to the 'popped-up' signal for the Menu.

The returned value represents the connection, and may be passed to DisconnectPoppedUp to remove it.
*/
func (recv *Menu) ConnectPoppedUp(callback MenuSignalPoppedUpCallback) int {
	signalMenuPoppedUpLock.Lock()
	defer signalMenuPoppedUpLock.Unlock()

	signalMenuPoppedUpId++
	instance := C.gpointer(recv.native)
	handlerID := C.Menu_signal_connect_popped_up(instance, C.gpointer(uintptr(signalMenuPoppedUpId)))

	detail := signalMenuPoppedUpDetail{callback, handlerID}
	signalMenuPoppedUpMap[signalMenuPoppedUpId] = detail

	return signalMenuPoppedUpId
}

/*
DisconnectPoppedUp disconnects a callback from the 'popped-up' signal for the Menu.

The connectionID should be a value returned from a call to ConnectPoppedUp.
*/
func (recv *Menu) DisconnectPoppedUp(connectionID int) {
	signalMenuPoppedUpLock.Lock()
	defer signalMenuPoppedUpLock.Unlock()

	detail, exists := signalMenuPoppedUpMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuPoppedUpMap, connectionID)
}

//export menu_poppedUpHandler
func menu_poppedUpHandler(_ *C.GObject, c_flipped_rect C.gpointer, c_final_rect C.gpointer, c_flipped_x C.gboolean, c_flipped_y C.gboolean, data C.gpointer) {
	signalMenuPoppedUpLock.RLock()
	defer signalMenuPoppedUpLock.RUnlock()

	flippedRect := uintptr(c_flipped_rect)

	finalRect := uintptr(c_final_rect)

	flippedX := c_flipped_x == C.TRUE

	flippedY := c_flipped_y == C.TRUE

	index := int(uintptr(data))
	callback := signalMenuPoppedUpMap[index].callback
	callback(flippedRect, finalRect, flippedX, flippedY)
}

// PlaceOnMonitor is a wrapper around the C function gtk_menu_place_on_monitor.
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

// PadControllerNew is a wrapper around the C function gtk_pad_controller_new.
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

	C.g_object_unref((C.gpointer)(retC))

	return retGo
}

// SetAction is a wrapper around the C function gtk_pad_controller_set_action.
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

// Popdown is a wrapper around the C function gtk_popover_popdown.
func (recv *Popover) Popdown() {
	C.gtk_popover_popdown((*C.GtkPopover)(recv.native))

	return
}

// Popup is a wrapper around the C function gtk_popover_popup.
func (recv *Popover) Popup() {
	C.gtk_popover_popup((*C.GtkPopover)(recv.native))

	return
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// GetMaxContentHeight is a wrapper around the C function gtk_scrolled_window_get_max_content_height.
func (recv *ScrolledWindow) GetMaxContentHeight() int32 {
	retC := C.gtk_scrolled_window_get_max_content_height((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMaxContentWidth is a wrapper around the C function gtk_scrolled_window_get_max_content_width.
func (recv *ScrolledWindow) GetMaxContentWidth() int32 {
	retC := C.gtk_scrolled_window_get_max_content_width((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPropagateNaturalHeight is a wrapper around the C function gtk_scrolled_window_get_propagate_natural_height.
func (recv *ScrolledWindow) GetPropagateNaturalHeight() bool {
	retC := C.gtk_scrolled_window_get_propagate_natural_height((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPropagateNaturalWidth is a wrapper around the C function gtk_scrolled_window_get_propagate_natural_width.
func (recv *ScrolledWindow) GetPropagateNaturalWidth() bool {
	retC := C.gtk_scrolled_window_get_propagate_natural_width((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetMaxContentHeight is a wrapper around the C function gtk_scrolled_window_set_max_content_height.
func (recv *ScrolledWindow) SetMaxContentHeight(height int32) {
	c_height := (C.gint)(height)

	C.gtk_scrolled_window_set_max_content_height((*C.GtkScrolledWindow)(recv.native), c_height)

	return
}

// SetMaxContentWidth is a wrapper around the C function gtk_scrolled_window_set_max_content_width.
func (recv *ScrolledWindow) SetMaxContentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_scrolled_window_set_max_content_width((*C.GtkScrolledWindow)(recv.native), c_width)

	return
}

// SetPropagateNaturalHeight is a wrapper around the C function gtk_scrolled_window_set_propagate_natural_height.
func (recv *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	c_propagate :=
		boolToGboolean(propagate)

	C.gtk_scrolled_window_set_propagate_natural_height((*C.GtkScrolledWindow)(recv.native), c_propagate)

	return
}

// SetPropagateNaturalWidth is a wrapper around the C function gtk_scrolled_window_set_propagate_natural_width.
func (recv *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	c_propagate :=
		boolToGboolean(propagate)

	C.gtk_scrolled_window_set_propagate_natural_width((*C.GtkScrolledWindow)(recv.native), c_propagate)

	return
}

// ShortcutLabelNew is a wrapper around the C function gtk_shortcut_label_new.
func ShortcutLabelNew(accelerator string) *ShortcutLabel {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	retC := C.gtk_shortcut_label_new(c_accelerator)
	retGo := ShortcutLabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAccelerator is a wrapper around the C function gtk_shortcut_label_get_accelerator.
func (recv *ShortcutLabel) GetAccelerator() string {
	retC := C.gtk_shortcut_label_get_accelerator((*C.GtkShortcutLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisabledText is a wrapper around the C function gtk_shortcut_label_get_disabled_text.
func (recv *ShortcutLabel) GetDisabledText() string {
	retC := C.gtk_shortcut_label_get_disabled_text((*C.GtkShortcutLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetAccelerator is a wrapper around the C function gtk_shortcut_label_set_accelerator.
func (recv *ShortcutLabel) SetAccelerator(accelerator string) {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	C.gtk_shortcut_label_set_accelerator((*C.GtkShortcutLabel)(recv.native), c_accelerator)

	return
}

// SetDisabledText is a wrapper around the C function gtk_shortcut_label_set_disabled_text.
func (recv *ShortcutLabel) SetDisabledText(disabledText string) {
	c_disabled_text := C.CString(disabledText)
	defer C.free(unsafe.Pointer(c_disabled_text))

	C.gtk_shortcut_label_set_disabled_text((*C.GtkShortcutLabel)(recv.native), c_disabled_text)

	return
}
