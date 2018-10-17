// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void EntryCompletion_noMatchesHandler();

	static gulong EntryCompletion_signal_connect_no_matches(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "no-matches", EntryCompletion_noMatchesHandler, data);
	}

*/
/*

	void Gesture_beginHandler();

	static gulong Gesture_signal_connect_begin(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "begin", Gesture_beginHandler, data);
	}

*/
/*

	void Gesture_cancelHandler();

	static gulong Gesture_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", Gesture_cancelHandler, data);
	}

*/
/*

	void Gesture_endHandler();

	static gulong Gesture_signal_connect_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "end", Gesture_endHandler, data);
	}

*/
/*

	void Gesture_sequenceStateChangedHandler();

	static gulong Gesture_signal_connect_sequence_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "sequence-state-changed", Gesture_sequenceStateChangedHandler, data);
	}

*/
/*

	void Gesture_updateHandler();

	static gulong Gesture_signal_connect_update(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update", Gesture_updateHandler, data);
	}

*/
/*

	void GestureDrag_dragBeginHandler();

	static gulong GestureDrag_signal_connect_drag_begin(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-begin", GestureDrag_dragBeginHandler, data);
	}

*/
/*

	void GestureDrag_dragEndHandler();

	static gulong GestureDrag_signal_connect_drag_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-end", GestureDrag_dragEndHandler, data);
	}

*/
/*

	void GestureDrag_dragUpdateHandler();

	static gulong GestureDrag_signal_connect_drag_update(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-update", GestureDrag_dragUpdateHandler, data);
	}

*/
/*

	void GestureLongPress_cancelledHandler();

	static gulong GestureLongPress_signal_connect_cancelled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancelled", GestureLongPress_cancelledHandler, data);
	}

*/
/*

	void GestureLongPress_pressedHandler();

	static gulong GestureLongPress_signal_connect_pressed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pressed", GestureLongPress_pressedHandler, data);
	}

*/
/*

	void GestureMultiPress_pressedHandler();

	static gulong GestureMultiPress_signal_connect_pressed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pressed", GestureMultiPress_pressedHandler, data);
	}

*/
/*

	void GestureMultiPress_releasedHandler();

	static gulong GestureMultiPress_signal_connect_released(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "released", GestureMultiPress_releasedHandler, data);
	}

*/
/*

	void GestureMultiPress_stoppedHandler();

	static gulong GestureMultiPress_signal_connect_stopped(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stopped", GestureMultiPress_stoppedHandler, data);
	}

*/
/*

	void GesturePan_panHandler();

	static gulong GesturePan_signal_connect_pan(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pan", GesturePan_panHandler, data);
	}

*/
/*

	void GestureRotate_angleChangedHandler();

	static gulong GestureRotate_signal_connect_angle_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "angle-changed", GestureRotate_angleChangedHandler, data);
	}

*/
/*

	void GestureSwipe_swipeHandler();

	static gulong GestureSwipe_signal_connect_swipe(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "swipe", GestureSwipe_swipeHandler, data);
	}

*/
/*

	void GestureZoom_scaleChangedHandler();

	static gulong GestureZoom_signal_connect_scale_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scale-changed", GestureZoom_scaleChangedHandler, data);
	}

*/
/*

	void ListBox_selectAllHandler();

	static gulong ListBox_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", ListBox_selectAllHandler, data);
	}

*/
/*

	void ListBox_selectedRowsChangedHandler();

	static gulong ListBox_signal_connect_selected_rows_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selected-rows-changed", ListBox_selectedRowsChangedHandler, data);
	}

*/
/*

	void ListBox_unselectAllHandler();

	static gulong ListBox_signal_connect_unselect_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unselect-all", ListBox_unselectAllHandler, data);
	}

*/
/*

	void PlacesSidebar_showEnterLocationHandler();

	static gulong PlacesSidebar_signal_connect_show_enter_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-enter-location", PlacesSidebar_showEnterLocationHandler, data);
	}

*/
/*

	void Switch_stateSetHandler();

	static gulong Switch_signal_connect_state_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-set", Switch_stateSetHandler, data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported : gtk_application_get_actions_for_accel : no return type

// GetMenuById is a wrapper around the C function gtk_application_get_menu_by_id.
func (recv *Application) GetMenuById(id string) *gio.Menu {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	retC := C.gtk_application_get_menu_by_id((*C.GtkApplication)(recv.native), c_id)
	retGo := gio.MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrefersAppMenu is a wrapper around the C function gtk_application_prefers_app_menu.
func (recv *Application) PrefersAppMenu() bool {
	retC := C.gtk_application_prefers_app_menu((*C.GtkApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// AttributeGetColumn is a wrapper around the C function gtk_cell_area_attribute_get_column.
func (recv *CellArea) AttributeGetColumn(renderer *CellRenderer, attribute string) int32 {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.gtk_cell_area_attribute_get_column((*C.GtkCellArea)(recv.native), c_renderer, c_attribute)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// EntryCompletionSignalNoMatchesCallback is a callback function for a 'no-matches' signal emitted from a EntryCompletion.
type EntryCompletionSignalNoMatchesCallback func()

// Unsupported : EntryIconAccessible : no CType

// GetPropagationPhase is a wrapper around the C function gtk_event_controller_get_propagation_phase.
func (recv *EventController) GetPropagationPhase() PropagationPhase {
	retC := C.gtk_event_controller_get_propagation_phase((*C.GtkEventController)(recv.native))
	retGo := (PropagationPhase)(retC)

	return retGo
}

// GetWidget is a wrapper around the C function gtk_event_controller_get_widget.
func (recv *EventController) GetWidget() *Widget {
	retC := C.gtk_event_controller_get_widget((*C.GtkEventController)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_event_controller_handle_event : unsupported parameter event : no type generator for Gdk.Event, const GdkEvent*

// Reset is a wrapper around the C function gtk_event_controller_reset.
func (recv *EventController) Reset() {
	C.gtk_event_controller_reset((*C.GtkEventController)(recv.native))

	return
}

// SetPropagationPhase is a wrapper around the C function gtk_event_controller_set_propagation_phase.
func (recv *EventController) SetPropagationPhase(phase PropagationPhase) {
	c_phase := (C.GtkPropagationPhase)(phase)

	C.gtk_event_controller_set_propagation_phase((*C.GtkEventController)(recv.native), c_phase)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GestureSignalBeginCallback is a callback function for a 'begin' signal emitted from a Gesture.
type GestureSignalBeginCallback func(sequence *gdk.EventSequence)

// GestureSignalCancelCallback is a callback function for a 'cancel' signal emitted from a Gesture.
type GestureSignalCancelCallback func(sequence *gdk.EventSequence)

// GestureSignalEndCallback is a callback function for a 'end' signal emitted from a Gesture.
type GestureSignalEndCallback func(sequence *gdk.EventSequence)

// GestureSignalSequenceStateChangedCallback is a callback function for a 'sequence-state-changed' signal emitted from a Gesture.
type GestureSignalSequenceStateChangedCallback func(sequence *gdk.EventSequence, state EventSequenceState)

// GestureSignalUpdateCallback is a callback function for a 'update' signal emitted from a Gesture.
type GestureSignalUpdateCallback func(sequence *gdk.EventSequence)

// Unsupported : gtk_gesture_get_bounding_box : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetBoundingBoxCenter is a wrapper around the C function gtk_gesture_get_bounding_box_center.
func (recv *Gesture) GetBoundingBoxCenter() (bool, float64, float64) {
	var c_x C.gdouble

	var c_y C.gdouble

	retC := C.gtk_gesture_get_bounding_box_center((*C.GtkGesture)(recv.native), &c_x, &c_y)
	retGo := retC == C.TRUE

	x := (float64)(c_x)

	y := (float64)(c_y)

	return retGo, x, y
}

// GetDevice is a wrapper around the C function gtk_gesture_get_device.
func (recv *Gesture) GetDevice() *gdk.Device {
	retC := C.gtk_gesture_get_device((*C.GtkGesture)(recv.native))
	retGo := gdk.DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_gesture_get_group.
func (recv *Gesture) GetGroup() *glib.List {
	retC := C.gtk_gesture_get_group((*C.GtkGesture)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLastUpdatedSequence is a wrapper around the C function gtk_gesture_get_last_updated_sequence.
func (recv *Gesture) GetLastUpdatedSequence() *gdk.EventSequence {
	retC := C.gtk_gesture_get_last_updated_sequence((*C.GtkGesture)(recv.native))
	retGo := gdk.EventSequenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPoint is a wrapper around the C function gtk_gesture_get_point.
func (recv *Gesture) GetPoint(sequence *gdk.EventSequence) (bool, float64, float64) {
	c_sequence := (*C.GdkEventSequence)(sequence.ToC())

	var c_x C.gdouble

	var c_y C.gdouble

	retC := C.gtk_gesture_get_point((*C.GtkGesture)(recv.native), c_sequence, &c_x, &c_y)
	retGo := retC == C.TRUE

	x := (float64)(c_x)

	y := (float64)(c_y)

	return retGo, x, y
}

// GetSequenceState is a wrapper around the C function gtk_gesture_get_sequence_state.
func (recv *Gesture) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	c_sequence := (*C.GdkEventSequence)(sequence.ToC())

	retC := C.gtk_gesture_get_sequence_state((*C.GtkGesture)(recv.native), c_sequence)
	retGo := (EventSequenceState)(retC)

	return retGo
}

// GetSequences is a wrapper around the C function gtk_gesture_get_sequences.
func (recv *Gesture) GetSequences() *glib.List {
	retC := C.gtk_gesture_get_sequences((*C.GtkGesture)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindow is a wrapper around the C function gtk_gesture_get_window.
func (recv *Gesture) GetWindow() *gdk.Window {
	retC := C.gtk_gesture_get_window((*C.GtkGesture)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Group is a wrapper around the C function gtk_gesture_group.
func (recv *Gesture) Group(gesture *Gesture) {
	c_gesture := (*C.GtkGesture)(gesture.ToC())

	C.gtk_gesture_group((*C.GtkGesture)(recv.native), c_gesture)

	return
}

// HandlesSequence is a wrapper around the C function gtk_gesture_handles_sequence.
func (recv *Gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	c_sequence := (*C.GdkEventSequence)(sequence.ToC())

	retC := C.gtk_gesture_handles_sequence((*C.GtkGesture)(recv.native), c_sequence)
	retGo := retC == C.TRUE

	return retGo
}

// IsActive is a wrapper around the C function gtk_gesture_is_active.
func (recv *Gesture) IsActive() bool {
	retC := C.gtk_gesture_is_active((*C.GtkGesture)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsGroupedWith is a wrapper around the C function gtk_gesture_is_grouped_with.
func (recv *Gesture) IsGroupedWith(other *Gesture) bool {
	c_other := (*C.GtkGesture)(other.ToC())

	retC := C.gtk_gesture_is_grouped_with((*C.GtkGesture)(recv.native), c_other)
	retGo := retC == C.TRUE

	return retGo
}

// IsRecognized is a wrapper around the C function gtk_gesture_is_recognized.
func (recv *Gesture) IsRecognized() bool {
	retC := C.gtk_gesture_is_recognized((*C.GtkGesture)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetSequenceState is a wrapper around the C function gtk_gesture_set_sequence_state.
func (recv *Gesture) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	c_sequence := (*C.GdkEventSequence)(sequence.ToC())

	c_state := (C.GtkEventSequenceState)(state)

	retC := C.gtk_gesture_set_sequence_state((*C.GtkGesture)(recv.native), c_sequence, c_state)
	retGo := retC == C.TRUE

	return retGo
}

// SetState is a wrapper around the C function gtk_gesture_set_state.
func (recv *Gesture) SetState(state EventSequenceState) bool {
	c_state := (C.GtkEventSequenceState)(state)

	retC := C.gtk_gesture_set_state((*C.GtkGesture)(recv.native), c_state)
	retGo := retC == C.TRUE

	return retGo
}

// SetWindow is a wrapper around the C function gtk_gesture_set_window.
func (recv *Gesture) SetWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gtk_gesture_set_window((*C.GtkGesture)(recv.native), c_window)

	return
}

// Ungroup is a wrapper around the C function gtk_gesture_ungroup.
func (recv *Gesture) Ungroup() {
	C.gtk_gesture_ungroup((*C.GtkGesture)(recv.native))

	return
}

// GestureDragSignalDragBeginCallback is a callback function for a 'drag-begin' signal emitted from a GestureDrag.
type GestureDragSignalDragBeginCallback func(startX float64, startY float64)

// GestureDragSignalDragEndCallback is a callback function for a 'drag-end' signal emitted from a GestureDrag.
type GestureDragSignalDragEndCallback func(offsetX float64, offsetY float64)

// GestureDragSignalDragUpdateCallback is a callback function for a 'drag-update' signal emitted from a GestureDrag.
type GestureDragSignalDragUpdateCallback func(offsetX float64, offsetY float64)

// GestureDragNew is a wrapper around the C function gtk_gesture_drag_new.
func GestureDragNew(widget *Widget) *GestureDrag {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_drag_new(c_widget)
	retGo := GestureDragNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOffset is a wrapper around the C function gtk_gesture_drag_get_offset.
func (recv *GestureDrag) GetOffset() (bool, float64, float64) {
	var c_x C.gdouble

	var c_y C.gdouble

	retC := C.gtk_gesture_drag_get_offset((*C.GtkGestureDrag)(recv.native), &c_x, &c_y)
	retGo := retC == C.TRUE

	x := (float64)(c_x)

	y := (float64)(c_y)

	return retGo, x, y
}

// GetStartPoint is a wrapper around the C function gtk_gesture_drag_get_start_point.
func (recv *GestureDrag) GetStartPoint() (bool, float64, float64) {
	var c_x C.gdouble

	var c_y C.gdouble

	retC := C.gtk_gesture_drag_get_start_point((*C.GtkGestureDrag)(recv.native), &c_x, &c_y)
	retGo := retC == C.TRUE

	x := (float64)(c_x)

	y := (float64)(c_y)

	return retGo, x, y
}

// GestureLongPressSignalCancelledCallback is a callback function for a 'cancelled' signal emitted from a GestureLongPress.
type GestureLongPressSignalCancelledCallback func()

// GestureLongPressSignalPressedCallback is a callback function for a 'pressed' signal emitted from a GestureLongPress.
type GestureLongPressSignalPressedCallback func(x float64, y float64)

// GestureLongPressNew is a wrapper around the C function gtk_gesture_long_press_new.
func GestureLongPressNew(widget *Widget) *GestureLongPress {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_long_press_new(c_widget)
	retGo := GestureLongPressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GestureMultiPressSignalPressedCallback is a callback function for a 'pressed' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalPressedCallback func(nPress int32, x float64, y float64)

// GestureMultiPressSignalReleasedCallback is a callback function for a 'released' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalReleasedCallback func(nPress int32, x float64, y float64)

// GestureMultiPressSignalStoppedCallback is a callback function for a 'stopped' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalStoppedCallback func()

// GestureMultiPressNew is a wrapper around the C function gtk_gesture_multi_press_new.
func GestureMultiPressNew(widget *Widget) *GestureMultiPress {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_multi_press_new(c_widget)
	retGo := GestureMultiPressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_gesture_multi_press_get_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_multi_press_set_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GesturePanSignalPanCallback is a callback function for a 'pan' signal emitted from a GesturePan.
type GesturePanSignalPanCallback func(direction PanDirection, offset float64)

// GesturePanNew is a wrapper around the C function gtk_gesture_pan_new.
func GesturePanNew(widget *Widget, orientation Orientation) *GesturePan {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_gesture_pan_new(c_widget, c_orientation)
	retGo := GesturePanNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_gesture_pan_get_orientation.
func (recv *GesturePan) GetOrientation() Orientation {
	retC := C.gtk_gesture_pan_get_orientation((*C.GtkGesturePan)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// SetOrientation is a wrapper around the C function gtk_gesture_pan_set_orientation.
func (recv *GesturePan) SetOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_gesture_pan_set_orientation((*C.GtkGesturePan)(recv.native), c_orientation)

	return
}

// GestureRotateSignalAngleChangedCallback is a callback function for a 'angle-changed' signal emitted from a GestureRotate.
type GestureRotateSignalAngleChangedCallback func(angle float64, angleDelta float64)

// GestureRotateNew is a wrapper around the C function gtk_gesture_rotate_new.
func GestureRotateNew(widget *Widget) *GestureRotate {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_rotate_new(c_widget)
	retGo := GestureRotateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAngleDelta is a wrapper around the C function gtk_gesture_rotate_get_angle_delta.
func (recv *GestureRotate) GetAngleDelta() float64 {
	retC := C.gtk_gesture_rotate_get_angle_delta((*C.GtkGestureRotate)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetButton is a wrapper around the C function gtk_gesture_single_get_button.
func (recv *GestureSingle) GetButton() uint32 {
	retC := C.gtk_gesture_single_get_button((*C.GtkGestureSingle)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetCurrentButton is a wrapper around the C function gtk_gesture_single_get_current_button.
func (recv *GestureSingle) GetCurrentButton() uint32 {
	retC := C.gtk_gesture_single_get_current_button((*C.GtkGestureSingle)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetCurrentSequence is a wrapper around the C function gtk_gesture_single_get_current_sequence.
func (recv *GestureSingle) GetCurrentSequence() *gdk.EventSequence {
	retC := C.gtk_gesture_single_get_current_sequence((*C.GtkGestureSingle)(recv.native))
	retGo := gdk.EventSequenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExclusive is a wrapper around the C function gtk_gesture_single_get_exclusive.
func (recv *GestureSingle) GetExclusive() bool {
	retC := C.gtk_gesture_single_get_exclusive((*C.GtkGestureSingle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTouchOnly is a wrapper around the C function gtk_gesture_single_get_touch_only.
func (recv *GestureSingle) GetTouchOnly() bool {
	retC := C.gtk_gesture_single_get_touch_only((*C.GtkGestureSingle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetButton is a wrapper around the C function gtk_gesture_single_set_button.
func (recv *GestureSingle) SetButton(button uint32) {
	c_button := (C.guint)(button)

	C.gtk_gesture_single_set_button((*C.GtkGestureSingle)(recv.native), c_button)

	return
}

// SetExclusive is a wrapper around the C function gtk_gesture_single_set_exclusive.
func (recv *GestureSingle) SetExclusive(exclusive bool) {
	c_exclusive :=
		boolToGboolean(exclusive)

	C.gtk_gesture_single_set_exclusive((*C.GtkGestureSingle)(recv.native), c_exclusive)

	return
}

// SetTouchOnly is a wrapper around the C function gtk_gesture_single_set_touch_only.
func (recv *GestureSingle) SetTouchOnly(touchOnly bool) {
	c_touch_only :=
		boolToGboolean(touchOnly)

	C.gtk_gesture_single_set_touch_only((*C.GtkGestureSingle)(recv.native), c_touch_only)

	return
}

// GestureSwipeSignalSwipeCallback is a callback function for a 'swipe' signal emitted from a GestureSwipe.
type GestureSwipeSignalSwipeCallback func(velocityX float64, velocityY float64)

// GestureSwipeNew is a wrapper around the C function gtk_gesture_swipe_new.
func GestureSwipeNew(widget *Widget) *GestureSwipe {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_swipe_new(c_widget)
	retGo := GestureSwipeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVelocity is a wrapper around the C function gtk_gesture_swipe_get_velocity.
func (recv *GestureSwipe) GetVelocity() (bool, float64, float64) {
	var c_velocity_x C.gdouble

	var c_velocity_y C.gdouble

	retC := C.gtk_gesture_swipe_get_velocity((*C.GtkGestureSwipe)(recv.native), &c_velocity_x, &c_velocity_y)
	retGo := retC == C.TRUE

	velocityX := (float64)(c_velocity_x)

	velocityY := (float64)(c_velocity_y)

	return retGo, velocityX, velocityY
}

// GestureZoomSignalScaleChangedCallback is a callback function for a 'scale-changed' signal emitted from a GestureZoom.
type GestureZoomSignalScaleChangedCallback func(scale float64)

// GestureZoomNew is a wrapper around the C function gtk_gesture_zoom_new.
func GestureZoomNew(widget *Widget) *GestureZoom {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_zoom_new(c_widget)
	retGo := GestureZoomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScaleDelta is a wrapper around the C function gtk_gesture_zoom_get_scale_delta.
func (recv *GestureZoom) GetScaleDelta() float64 {
	retC := C.gtk_gesture_zoom_get_scale_delta((*C.GtkGestureZoom)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// AddResourcePath is a wrapper around the C function gtk_icon_theme_add_resource_path.
func (recv *IconTheme) AddResourcePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_add_resource_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// ListBoxSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a ListBox.
type ListBoxSignalSelectAllCallback func()

// ListBoxSignalSelectedRowsChangedCallback is a callback function for a 'selected-rows-changed' signal emitted from a ListBox.
type ListBoxSignalSelectedRowsChangedCallback func()

// ListBoxSignalUnselectAllCallback is a callback function for a 'unselect-all' signal emitted from a ListBox.
type ListBoxSignalUnselectAllCallback func()

// GetSelectedRows is a wrapper around the C function gtk_list_box_get_selected_rows.
func (recv *ListBox) GetSelectedRows() *glib.List {
	retC := C.gtk_list_box_get_selected_rows((*C.GtkListBox)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SelectAll is a wrapper around the C function gtk_list_box_select_all.
func (recv *ListBox) SelectAll() {
	C.gtk_list_box_select_all((*C.GtkListBox)(recv.native))

	return
}

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc, GtkListBoxForeachFunc

// UnselectAll is a wrapper around the C function gtk_list_box_unselect_all.
func (recv *ListBox) UnselectAll() {
	C.gtk_list_box_unselect_all((*C.GtkListBox)(recv.native))

	return
}

// UnselectRow is a wrapper around the C function gtk_list_box_unselect_row.
func (recv *ListBox) UnselectRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(row.ToC())

	C.gtk_list_box_unselect_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// GetActivatable is a wrapper around the C function gtk_list_box_row_get_activatable.
func (recv *ListBoxRow) GetActivatable() bool {
	retC := C.gtk_list_box_row_get_activatable((*C.GtkListBoxRow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectable is a wrapper around the C function gtk_list_box_row_get_selectable.
func (recv *ListBoxRow) GetSelectable() bool {
	retC := C.gtk_list_box_row_get_selectable((*C.GtkListBoxRow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSelected is a wrapper around the C function gtk_list_box_row_is_selected.
func (recv *ListBoxRow) IsSelected() bool {
	retC := C.gtk_list_box_row_is_selected((*C.GtkListBoxRow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActivatable is a wrapper around the C function gtk_list_box_row_set_activatable.
func (recv *ListBoxRow) SetActivatable(activatable bool) {
	c_activatable :=
		boolToGboolean(activatable)

	C.gtk_list_box_row_set_activatable((*C.GtkListBoxRow)(recv.native), c_activatable)

	return
}

// SetSelectable is a wrapper around the C function gtk_list_box_row_set_selectable.
func (recv *ListBoxRow) SetSelectable(selectable bool) {
	c_selectable :=
		boolToGboolean(selectable)

	C.gtk_list_box_row_set_selectable((*C.GtkListBoxRow)(recv.native), c_selectable)

	return
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal : unsupported parameter selected_item : no type generator for Gio.File,

// PlacesSidebarSignalShowEnterLocationCallback is a callback function for a 'show-enter-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowEnterLocationCallback func()

// GetShowEnterLocation is a wrapper around the C function gtk_places_sidebar_get_show_enter_location.
func (recv *PlacesSidebar) GetShowEnterLocation() bool {
	retC := C.gtk_places_sidebar_get_show_enter_location((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowEnterLocation is a wrapper around the C function gtk_places_sidebar_set_show_enter_location.
func (recv *PlacesSidebar) SetShowEnterLocation(showEnterLocation bool) {
	c_show_enter_location :=
		boolToGboolean(showEnterLocation)

	C.gtk_places_sidebar_set_show_enter_location((*C.GtkPlacesSidebar)(recv.native), c_show_enter_location)

	return
}

// Unsupported signal : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SwitchSignalStateSetCallback is a callback function for a 'state-set' signal emitted from a Switch.
type SwitchSignalStateSetCallback func(state bool) bool

// GetState is a wrapper around the C function gtk_switch_get_state.
func (recv *Switch) GetState() bool {
	retC := C.gtk_switch_get_state((*C.GtkSwitch)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetState is a wrapper around the C function gtk_switch_set_state.
func (recv *Switch) SetState(state bool) {
	c_state :=
		boolToGboolean(state)

	C.gtk_switch_set_state((*C.GtkSwitch)(recv.native), c_state)

	return
}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported signal : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle
