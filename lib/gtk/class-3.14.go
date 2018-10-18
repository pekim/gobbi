// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
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

var signalEntryCompletionNoMatchesId int
var signalEntryCompletionNoMatchesMap = make(map[int]EntryCompletionSignalNoMatchesCallback)
var signalEntryCompletionNoMatchesLock sync.Mutex

// EntryCompletionSignalNoMatchesCallback is a callback function for a 'no-matches' signal emitted from a EntryCompletion.
type EntryCompletionSignalNoMatchesCallback func()

/*
ConnectNoMatches connects the callback to the 'no-matches' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectNoMatches to remove it.
*/
func (recv *EntryCompletion) ConnectNoMatches(callback EntryCompletionSignalNoMatchesCallback) int {
	signalEntryCompletionNoMatchesLock.Lock()
	defer signalEntryCompletionNoMatchesLock.Unlock()

	signalEntryCompletionNoMatchesId++
	signalEntryCompletionNoMatchesMap[signalEntryCompletionNoMatchesId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.EntryCompletion_signal_connect_no_matches(instance, C.gpointer(uintptr(signalEntryCompletionNoMatchesId)))
	return int(retC)
}

//export EntryCompletion_noMatchesHandler
func EntryCompletion_noMatchesHandler() {
	fmt.Println("cb")
}

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

var signalGestureBeginId int
var signalGestureBeginMap = make(map[int]GestureSignalBeginCallback)
var signalGestureBeginLock sync.Mutex

// GestureSignalBeginCallback is a callback function for a 'begin' signal emitted from a Gesture.
type GestureSignalBeginCallback func(sequence *gdk.EventSequence)

/*
ConnectBegin connects the callback to the 'begin' signal for the Gesture.

The returned value represents the connection, and may be passed to DisconnectBegin to remove it.
*/
func (recv *Gesture) ConnectBegin(callback GestureSignalBeginCallback) int {
	signalGestureBeginLock.Lock()
	defer signalGestureBeginLock.Unlock()

	signalGestureBeginId++
	signalGestureBeginMap[signalGestureBeginId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Gesture_signal_connect_begin(instance, C.gpointer(uintptr(signalGestureBeginId)))
	return int(retC)
}

//export Gesture_beginHandler
func Gesture_beginHandler() {
	fmt.Println("cb")
}

var signalGestureCancelId int
var signalGestureCancelMap = make(map[int]GestureSignalCancelCallback)
var signalGestureCancelLock sync.Mutex

// GestureSignalCancelCallback is a callback function for a 'cancel' signal emitted from a Gesture.
type GestureSignalCancelCallback func(sequence *gdk.EventSequence)

/*
ConnectCancel connects the callback to the 'cancel' signal for the Gesture.

The returned value represents the connection, and may be passed to DisconnectCancel to remove it.
*/
func (recv *Gesture) ConnectCancel(callback GestureSignalCancelCallback) int {
	signalGestureCancelLock.Lock()
	defer signalGestureCancelLock.Unlock()

	signalGestureCancelId++
	signalGestureCancelMap[signalGestureCancelId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Gesture_signal_connect_cancel(instance, C.gpointer(uintptr(signalGestureCancelId)))
	return int(retC)
}

//export Gesture_cancelHandler
func Gesture_cancelHandler() {
	fmt.Println("cb")
}

var signalGestureEndId int
var signalGestureEndMap = make(map[int]GestureSignalEndCallback)
var signalGestureEndLock sync.Mutex

// GestureSignalEndCallback is a callback function for a 'end' signal emitted from a Gesture.
type GestureSignalEndCallback func(sequence *gdk.EventSequence)

/*
ConnectEnd connects the callback to the 'end' signal for the Gesture.

The returned value represents the connection, and may be passed to DisconnectEnd to remove it.
*/
func (recv *Gesture) ConnectEnd(callback GestureSignalEndCallback) int {
	signalGestureEndLock.Lock()
	defer signalGestureEndLock.Unlock()

	signalGestureEndId++
	signalGestureEndMap[signalGestureEndId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Gesture_signal_connect_end(instance, C.gpointer(uintptr(signalGestureEndId)))
	return int(retC)
}

//export Gesture_endHandler
func Gesture_endHandler() {
	fmt.Println("cb")
}

var signalGestureSequenceStateChangedId int
var signalGestureSequenceStateChangedMap = make(map[int]GestureSignalSequenceStateChangedCallback)
var signalGestureSequenceStateChangedLock sync.Mutex

// GestureSignalSequenceStateChangedCallback is a callback function for a 'sequence-state-changed' signal emitted from a Gesture.
type GestureSignalSequenceStateChangedCallback func(sequence *gdk.EventSequence, state EventSequenceState)

/*
ConnectSequenceStateChanged connects the callback to the 'sequence-state-changed' signal for the Gesture.

The returned value represents the connection, and may be passed to DisconnectSequenceStateChanged to remove it.
*/
func (recv *Gesture) ConnectSequenceStateChanged(callback GestureSignalSequenceStateChangedCallback) int {
	signalGestureSequenceStateChangedLock.Lock()
	defer signalGestureSequenceStateChangedLock.Unlock()

	signalGestureSequenceStateChangedId++
	signalGestureSequenceStateChangedMap[signalGestureSequenceStateChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Gesture_signal_connect_sequence_state_changed(instance, C.gpointer(uintptr(signalGestureSequenceStateChangedId)))
	return int(retC)
}

//export Gesture_sequenceStateChangedHandler
func Gesture_sequenceStateChangedHandler() {
	fmt.Println("cb")
}

var signalGestureUpdateId int
var signalGestureUpdateMap = make(map[int]GestureSignalUpdateCallback)
var signalGestureUpdateLock sync.Mutex

// GestureSignalUpdateCallback is a callback function for a 'update' signal emitted from a Gesture.
type GestureSignalUpdateCallback func(sequence *gdk.EventSequence)

/*
ConnectUpdate connects the callback to the 'update' signal for the Gesture.

The returned value represents the connection, and may be passed to DisconnectUpdate to remove it.
*/
func (recv *Gesture) ConnectUpdate(callback GestureSignalUpdateCallback) int {
	signalGestureUpdateLock.Lock()
	defer signalGestureUpdateLock.Unlock()

	signalGestureUpdateId++
	signalGestureUpdateMap[signalGestureUpdateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Gesture_signal_connect_update(instance, C.gpointer(uintptr(signalGestureUpdateId)))
	return int(retC)
}

//export Gesture_updateHandler
func Gesture_updateHandler() {
	fmt.Println("cb")
}

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

var signalGestureDragDragBeginId int
var signalGestureDragDragBeginMap = make(map[int]GestureDragSignalDragBeginCallback)
var signalGestureDragDragBeginLock sync.Mutex

// GestureDragSignalDragBeginCallback is a callback function for a 'drag-begin' signal emitted from a GestureDrag.
type GestureDragSignalDragBeginCallback func(startX float64, startY float64)

/*
ConnectDragBegin connects the callback to the 'drag-begin' signal for the GestureDrag.

The returned value represents the connection, and may be passed to DisconnectDragBegin to remove it.
*/
func (recv *GestureDrag) ConnectDragBegin(callback GestureDragSignalDragBeginCallback) int {
	signalGestureDragDragBeginLock.Lock()
	defer signalGestureDragDragBeginLock.Unlock()

	signalGestureDragDragBeginId++
	signalGestureDragDragBeginMap[signalGestureDragDragBeginId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureDrag_signal_connect_drag_begin(instance, C.gpointer(uintptr(signalGestureDragDragBeginId)))
	return int(retC)
}

//export GestureDrag_dragBeginHandler
func GestureDrag_dragBeginHandler() {
	fmt.Println("cb")
}

var signalGestureDragDragEndId int
var signalGestureDragDragEndMap = make(map[int]GestureDragSignalDragEndCallback)
var signalGestureDragDragEndLock sync.Mutex

// GestureDragSignalDragEndCallback is a callback function for a 'drag-end' signal emitted from a GestureDrag.
type GestureDragSignalDragEndCallback func(offsetX float64, offsetY float64)

/*
ConnectDragEnd connects the callback to the 'drag-end' signal for the GestureDrag.

The returned value represents the connection, and may be passed to DisconnectDragEnd to remove it.
*/
func (recv *GestureDrag) ConnectDragEnd(callback GestureDragSignalDragEndCallback) int {
	signalGestureDragDragEndLock.Lock()
	defer signalGestureDragDragEndLock.Unlock()

	signalGestureDragDragEndId++
	signalGestureDragDragEndMap[signalGestureDragDragEndId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureDrag_signal_connect_drag_end(instance, C.gpointer(uintptr(signalGestureDragDragEndId)))
	return int(retC)
}

//export GestureDrag_dragEndHandler
func GestureDrag_dragEndHandler() {
	fmt.Println("cb")
}

var signalGestureDragDragUpdateId int
var signalGestureDragDragUpdateMap = make(map[int]GestureDragSignalDragUpdateCallback)
var signalGestureDragDragUpdateLock sync.Mutex

// GestureDragSignalDragUpdateCallback is a callback function for a 'drag-update' signal emitted from a GestureDrag.
type GestureDragSignalDragUpdateCallback func(offsetX float64, offsetY float64)

/*
ConnectDragUpdate connects the callback to the 'drag-update' signal for the GestureDrag.

The returned value represents the connection, and may be passed to DisconnectDragUpdate to remove it.
*/
func (recv *GestureDrag) ConnectDragUpdate(callback GestureDragSignalDragUpdateCallback) int {
	signalGestureDragDragUpdateLock.Lock()
	defer signalGestureDragDragUpdateLock.Unlock()

	signalGestureDragDragUpdateId++
	signalGestureDragDragUpdateMap[signalGestureDragDragUpdateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureDrag_signal_connect_drag_update(instance, C.gpointer(uintptr(signalGestureDragDragUpdateId)))
	return int(retC)
}

//export GestureDrag_dragUpdateHandler
func GestureDrag_dragUpdateHandler() {
	fmt.Println("cb")
}

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

var signalGestureLongPressCancelledId int
var signalGestureLongPressCancelledMap = make(map[int]GestureLongPressSignalCancelledCallback)
var signalGestureLongPressCancelledLock sync.Mutex

// GestureLongPressSignalCancelledCallback is a callback function for a 'cancelled' signal emitted from a GestureLongPress.
type GestureLongPressSignalCancelledCallback func()

/*
ConnectCancelled connects the callback to the 'cancelled' signal for the GestureLongPress.

The returned value represents the connection, and may be passed to DisconnectCancelled to remove it.
*/
func (recv *GestureLongPress) ConnectCancelled(callback GestureLongPressSignalCancelledCallback) int {
	signalGestureLongPressCancelledLock.Lock()
	defer signalGestureLongPressCancelledLock.Unlock()

	signalGestureLongPressCancelledId++
	signalGestureLongPressCancelledMap[signalGestureLongPressCancelledId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureLongPress_signal_connect_cancelled(instance, C.gpointer(uintptr(signalGestureLongPressCancelledId)))
	return int(retC)
}

//export GestureLongPress_cancelledHandler
func GestureLongPress_cancelledHandler() {
	fmt.Println("cb")
}

var signalGestureLongPressPressedId int
var signalGestureLongPressPressedMap = make(map[int]GestureLongPressSignalPressedCallback)
var signalGestureLongPressPressedLock sync.Mutex

// GestureLongPressSignalPressedCallback is a callback function for a 'pressed' signal emitted from a GestureLongPress.
type GestureLongPressSignalPressedCallback func(x float64, y float64)

/*
ConnectPressed connects the callback to the 'pressed' signal for the GestureLongPress.

The returned value represents the connection, and may be passed to DisconnectPressed to remove it.
*/
func (recv *GestureLongPress) ConnectPressed(callback GestureLongPressSignalPressedCallback) int {
	signalGestureLongPressPressedLock.Lock()
	defer signalGestureLongPressPressedLock.Unlock()

	signalGestureLongPressPressedId++
	signalGestureLongPressPressedMap[signalGestureLongPressPressedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureLongPress_signal_connect_pressed(instance, C.gpointer(uintptr(signalGestureLongPressPressedId)))
	return int(retC)
}

//export GestureLongPress_pressedHandler
func GestureLongPress_pressedHandler() {
	fmt.Println("cb")
}

// GestureLongPressNew is a wrapper around the C function gtk_gesture_long_press_new.
func GestureLongPressNew(widget *Widget) *GestureLongPress {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_long_press_new(c_widget)
	retGo := GestureLongPressNewFromC(unsafe.Pointer(retC))

	return retGo
}

var signalGestureMultiPressPressedId int
var signalGestureMultiPressPressedMap = make(map[int]GestureMultiPressSignalPressedCallback)
var signalGestureMultiPressPressedLock sync.Mutex

// GestureMultiPressSignalPressedCallback is a callback function for a 'pressed' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalPressedCallback func(nPress int32, x float64, y float64)

/*
ConnectPressed connects the callback to the 'pressed' signal for the GestureMultiPress.

The returned value represents the connection, and may be passed to DisconnectPressed to remove it.
*/
func (recv *GestureMultiPress) ConnectPressed(callback GestureMultiPressSignalPressedCallback) int {
	signalGestureMultiPressPressedLock.Lock()
	defer signalGestureMultiPressPressedLock.Unlock()

	signalGestureMultiPressPressedId++
	signalGestureMultiPressPressedMap[signalGestureMultiPressPressedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureMultiPress_signal_connect_pressed(instance, C.gpointer(uintptr(signalGestureMultiPressPressedId)))
	return int(retC)
}

//export GestureMultiPress_pressedHandler
func GestureMultiPress_pressedHandler() {
	fmt.Println("cb")
}

var signalGestureMultiPressReleasedId int
var signalGestureMultiPressReleasedMap = make(map[int]GestureMultiPressSignalReleasedCallback)
var signalGestureMultiPressReleasedLock sync.Mutex

// GestureMultiPressSignalReleasedCallback is a callback function for a 'released' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalReleasedCallback func(nPress int32, x float64, y float64)

/*
ConnectReleased connects the callback to the 'released' signal for the GestureMultiPress.

The returned value represents the connection, and may be passed to DisconnectReleased to remove it.
*/
func (recv *GestureMultiPress) ConnectReleased(callback GestureMultiPressSignalReleasedCallback) int {
	signalGestureMultiPressReleasedLock.Lock()
	defer signalGestureMultiPressReleasedLock.Unlock()

	signalGestureMultiPressReleasedId++
	signalGestureMultiPressReleasedMap[signalGestureMultiPressReleasedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureMultiPress_signal_connect_released(instance, C.gpointer(uintptr(signalGestureMultiPressReleasedId)))
	return int(retC)
}

//export GestureMultiPress_releasedHandler
func GestureMultiPress_releasedHandler() {
	fmt.Println("cb")
}

var signalGestureMultiPressStoppedId int
var signalGestureMultiPressStoppedMap = make(map[int]GestureMultiPressSignalStoppedCallback)
var signalGestureMultiPressStoppedLock sync.Mutex

// GestureMultiPressSignalStoppedCallback is a callback function for a 'stopped' signal emitted from a GestureMultiPress.
type GestureMultiPressSignalStoppedCallback func()

/*
ConnectStopped connects the callback to the 'stopped' signal for the GestureMultiPress.

The returned value represents the connection, and may be passed to DisconnectStopped to remove it.
*/
func (recv *GestureMultiPress) ConnectStopped(callback GestureMultiPressSignalStoppedCallback) int {
	signalGestureMultiPressStoppedLock.Lock()
	defer signalGestureMultiPressStoppedLock.Unlock()

	signalGestureMultiPressStoppedId++
	signalGestureMultiPressStoppedMap[signalGestureMultiPressStoppedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureMultiPress_signal_connect_stopped(instance, C.gpointer(uintptr(signalGestureMultiPressStoppedId)))
	return int(retC)
}

//export GestureMultiPress_stoppedHandler
func GestureMultiPress_stoppedHandler() {
	fmt.Println("cb")
}

// GestureMultiPressNew is a wrapper around the C function gtk_gesture_multi_press_new.
func GestureMultiPressNew(widget *Widget) *GestureMultiPress {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_multi_press_new(c_widget)
	retGo := GestureMultiPressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_gesture_multi_press_get_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_multi_press_set_area : unsupported parameter rect : Blacklisted record : GdkRectangle

var signalGesturePanPanId int
var signalGesturePanPanMap = make(map[int]GesturePanSignalPanCallback)
var signalGesturePanPanLock sync.Mutex

// GesturePanSignalPanCallback is a callback function for a 'pan' signal emitted from a GesturePan.
type GesturePanSignalPanCallback func(direction PanDirection, offset float64)

/*
ConnectPan connects the callback to the 'pan' signal for the GesturePan.

The returned value represents the connection, and may be passed to DisconnectPan to remove it.
*/
func (recv *GesturePan) ConnectPan(callback GesturePanSignalPanCallback) int {
	signalGesturePanPanLock.Lock()
	defer signalGesturePanPanLock.Unlock()

	signalGesturePanPanId++
	signalGesturePanPanMap[signalGesturePanPanId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GesturePan_signal_connect_pan(instance, C.gpointer(uintptr(signalGesturePanPanId)))
	return int(retC)
}

//export GesturePan_panHandler
func GesturePan_panHandler() {
	fmt.Println("cb")
}

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

var signalGestureRotateAngleChangedId int
var signalGestureRotateAngleChangedMap = make(map[int]GestureRotateSignalAngleChangedCallback)
var signalGestureRotateAngleChangedLock sync.Mutex

// GestureRotateSignalAngleChangedCallback is a callback function for a 'angle-changed' signal emitted from a GestureRotate.
type GestureRotateSignalAngleChangedCallback func(angle float64, angleDelta float64)

/*
ConnectAngleChanged connects the callback to the 'angle-changed' signal for the GestureRotate.

The returned value represents the connection, and may be passed to DisconnectAngleChanged to remove it.
*/
func (recv *GestureRotate) ConnectAngleChanged(callback GestureRotateSignalAngleChangedCallback) int {
	signalGestureRotateAngleChangedLock.Lock()
	defer signalGestureRotateAngleChangedLock.Unlock()

	signalGestureRotateAngleChangedId++
	signalGestureRotateAngleChangedMap[signalGestureRotateAngleChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureRotate_signal_connect_angle_changed(instance, C.gpointer(uintptr(signalGestureRotateAngleChangedId)))
	return int(retC)
}

//export GestureRotate_angleChangedHandler
func GestureRotate_angleChangedHandler() {
	fmt.Println("cb")
}

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

var signalGestureSwipeSwipeId int
var signalGestureSwipeSwipeMap = make(map[int]GestureSwipeSignalSwipeCallback)
var signalGestureSwipeSwipeLock sync.Mutex

// GestureSwipeSignalSwipeCallback is a callback function for a 'swipe' signal emitted from a GestureSwipe.
type GestureSwipeSignalSwipeCallback func(velocityX float64, velocityY float64)

/*
ConnectSwipe connects the callback to the 'swipe' signal for the GestureSwipe.

The returned value represents the connection, and may be passed to DisconnectSwipe to remove it.
*/
func (recv *GestureSwipe) ConnectSwipe(callback GestureSwipeSignalSwipeCallback) int {
	signalGestureSwipeSwipeLock.Lock()
	defer signalGestureSwipeSwipeLock.Unlock()

	signalGestureSwipeSwipeId++
	signalGestureSwipeSwipeMap[signalGestureSwipeSwipeId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureSwipe_signal_connect_swipe(instance, C.gpointer(uintptr(signalGestureSwipeSwipeId)))
	return int(retC)
}

//export GestureSwipe_swipeHandler
func GestureSwipe_swipeHandler() {
	fmt.Println("cb")
}

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

var signalGestureZoomScaleChangedId int
var signalGestureZoomScaleChangedMap = make(map[int]GestureZoomSignalScaleChangedCallback)
var signalGestureZoomScaleChangedLock sync.Mutex

// GestureZoomSignalScaleChangedCallback is a callback function for a 'scale-changed' signal emitted from a GestureZoom.
type GestureZoomSignalScaleChangedCallback func(scale float64)

/*
ConnectScaleChanged connects the callback to the 'scale-changed' signal for the GestureZoom.

The returned value represents the connection, and may be passed to DisconnectScaleChanged to remove it.
*/
func (recv *GestureZoom) ConnectScaleChanged(callback GestureZoomSignalScaleChangedCallback) int {
	signalGestureZoomScaleChangedLock.Lock()
	defer signalGestureZoomScaleChangedLock.Unlock()

	signalGestureZoomScaleChangedId++
	signalGestureZoomScaleChangedMap[signalGestureZoomScaleChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.GestureZoom_signal_connect_scale_changed(instance, C.gpointer(uintptr(signalGestureZoomScaleChangedId)))
	return int(retC)
}

//export GestureZoom_scaleChangedHandler
func GestureZoom_scaleChangedHandler() {
	fmt.Println("cb")
}

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

var signalListBoxSelectAllId int
var signalListBoxSelectAllMap = make(map[int]ListBoxSignalSelectAllCallback)
var signalListBoxSelectAllLock sync.Mutex

// ListBoxSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a ListBox.
type ListBoxSignalSelectAllCallback func()

/*
ConnectSelectAll connects the callback to the 'select-all' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectSelectAll to remove it.
*/
func (recv *ListBox) ConnectSelectAll(callback ListBoxSignalSelectAllCallback) int {
	signalListBoxSelectAllLock.Lock()
	defer signalListBoxSelectAllLock.Unlock()

	signalListBoxSelectAllId++
	signalListBoxSelectAllMap[signalListBoxSelectAllId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ListBox_signal_connect_select_all(instance, C.gpointer(uintptr(signalListBoxSelectAllId)))
	return int(retC)
}

//export ListBox_selectAllHandler
func ListBox_selectAllHandler() {
	fmt.Println("cb")
}

var signalListBoxSelectedRowsChangedId int
var signalListBoxSelectedRowsChangedMap = make(map[int]ListBoxSignalSelectedRowsChangedCallback)
var signalListBoxSelectedRowsChangedLock sync.Mutex

// ListBoxSignalSelectedRowsChangedCallback is a callback function for a 'selected-rows-changed' signal emitted from a ListBox.
type ListBoxSignalSelectedRowsChangedCallback func()

/*
ConnectSelectedRowsChanged connects the callback to the 'selected-rows-changed' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectSelectedRowsChanged to remove it.
*/
func (recv *ListBox) ConnectSelectedRowsChanged(callback ListBoxSignalSelectedRowsChangedCallback) int {
	signalListBoxSelectedRowsChangedLock.Lock()
	defer signalListBoxSelectedRowsChangedLock.Unlock()

	signalListBoxSelectedRowsChangedId++
	signalListBoxSelectedRowsChangedMap[signalListBoxSelectedRowsChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ListBox_signal_connect_selected_rows_changed(instance, C.gpointer(uintptr(signalListBoxSelectedRowsChangedId)))
	return int(retC)
}

//export ListBox_selectedRowsChangedHandler
func ListBox_selectedRowsChangedHandler() {
	fmt.Println("cb")
}

var signalListBoxUnselectAllId int
var signalListBoxUnselectAllMap = make(map[int]ListBoxSignalUnselectAllCallback)
var signalListBoxUnselectAllLock sync.Mutex

// ListBoxSignalUnselectAllCallback is a callback function for a 'unselect-all' signal emitted from a ListBox.
type ListBoxSignalUnselectAllCallback func()

/*
ConnectUnselectAll connects the callback to the 'unselect-all' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectUnselectAll to remove it.
*/
func (recv *ListBox) ConnectUnselectAll(callback ListBoxSignalUnselectAllCallback) int {
	signalListBoxUnselectAllLock.Lock()
	defer signalListBoxUnselectAllLock.Unlock()

	signalListBoxUnselectAllId++
	signalListBoxUnselectAllMap[signalListBoxUnselectAllId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ListBox_signal_connect_unselect_all(instance, C.gpointer(uintptr(signalListBoxUnselectAllId)))
	return int(retC)
}

//export ListBox_unselectAllHandler
func ListBox_unselectAllHandler() {
	fmt.Println("cb")
}

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

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal : unsupported parameter selected_item : no type generator for Gio.File,

var signalPlacesSidebarShowEnterLocationId int
var signalPlacesSidebarShowEnterLocationMap = make(map[int]PlacesSidebarSignalShowEnterLocationCallback)
var signalPlacesSidebarShowEnterLocationLock sync.Mutex

// PlacesSidebarSignalShowEnterLocationCallback is a callback function for a 'show-enter-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowEnterLocationCallback func()

/*
ConnectShowEnterLocation connects the callback to the 'show-enter-location' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowEnterLocation to remove it.
*/
func (recv *PlacesSidebar) ConnectShowEnterLocation(callback PlacesSidebarSignalShowEnterLocationCallback) int {
	signalPlacesSidebarShowEnterLocationLock.Lock()
	defer signalPlacesSidebarShowEnterLocationLock.Unlock()

	signalPlacesSidebarShowEnterLocationId++
	signalPlacesSidebarShowEnterLocationMap[signalPlacesSidebarShowEnterLocationId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.PlacesSidebar_signal_connect_show_enter_location(instance, C.gpointer(uintptr(signalPlacesSidebarShowEnterLocationId)))
	return int(retC)
}

//export PlacesSidebar_showEnterLocationHandler
func PlacesSidebar_showEnterLocationHandler() {
	fmt.Println("cb")
}

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

var signalSwitchStateSetId int
var signalSwitchStateSetMap = make(map[int]SwitchSignalStateSetCallback)
var signalSwitchStateSetLock sync.Mutex

// SwitchSignalStateSetCallback is a callback function for a 'state-set' signal emitted from a Switch.
type SwitchSignalStateSetCallback func(state bool) bool

/*
ConnectStateSet connects the callback to the 'state-set' signal for the Switch.

The returned value represents the connection, and may be passed to DisconnectStateSet to remove it.
*/
func (recv *Switch) ConnectStateSet(callback SwitchSignalStateSetCallback) int {
	signalSwitchStateSetLock.Lock()
	defer signalSwitchStateSetLock.Unlock()

	signalSwitchStateSetId++
	signalSwitchStateSetMap[signalSwitchStateSetId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Switch_signal_connect_state_set(instance, C.gpointer(uintptr(signalSwitchStateSetId)))
	return int(retC)
}

//export Switch_stateSetHandler
func Switch_stateSetHandler() {
	fmt.Println("cb")
}

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
