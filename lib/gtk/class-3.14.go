// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

	void entrycompletion_noMatchesHandler(GObject *, gpointer);

	static gulong EntryCompletion_signal_connect_no_matches(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "no-matches", G_CALLBACK(entrycompletion_noMatchesHandler), data);
	}

*/
/*

	void gesture_beginHandler(GObject *, GdkEventSequence *, gpointer);

	static gulong Gesture_signal_connect_begin(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "begin", G_CALLBACK(gesture_beginHandler), data);
	}

*/
/*

	void gesture_cancelHandler(GObject *, GdkEventSequence *, gpointer);

	static gulong Gesture_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", G_CALLBACK(gesture_cancelHandler), data);
	}

*/
/*

	void gesture_endHandler(GObject *, GdkEventSequence *, gpointer);

	static gulong Gesture_signal_connect_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "end", G_CALLBACK(gesture_endHandler), data);
	}

*/
/*

	void gesture_updateHandler(GObject *, GdkEventSequence *, gpointer);

	static gulong Gesture_signal_connect_update(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update", G_CALLBACK(gesture_updateHandler), data);
	}

*/
/*

	void gesturelongpress_cancelledHandler(GObject *, gpointer);

	static gulong GestureLongPress_signal_connect_cancelled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancelled", G_CALLBACK(gesturelongpress_cancelledHandler), data);
	}

*/
/*

	void gesturemultipress_stoppedHandler(GObject *, gpointer);

	static gulong GestureMultiPress_signal_connect_stopped(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stopped", G_CALLBACK(gesturemultipress_stoppedHandler), data);
	}

*/
/*

	void listbox_selectAllHandler(GObject *, gpointer);

	static gulong ListBox_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", G_CALLBACK(listbox_selectAllHandler), data);
	}

*/
/*

	void listbox_selectedRowsChangedHandler(GObject *, gpointer);

	static gulong ListBox_signal_connect_selected_rows_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selected-rows-changed", G_CALLBACK(listbox_selectedRowsChangedHandler), data);
	}

*/
/*

	void listbox_unselectAllHandler(GObject *, gpointer);

	static gulong ListBox_signal_connect_unselect_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unselect-all", G_CALLBACK(listbox_unselectAllHandler), data);
	}

*/
/*

	void placessidebar_showEnterLocationHandler(GObject *, gpointer);

	static gulong PlacesSidebar_signal_connect_show_enter_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-enter-location", G_CALLBACK(placessidebar_showEnterLocationHandler), data);
	}

*/
/*

	gboolean switch_stateSetHandler(GObject *, gboolean, gpointer);

	static gulong Switch_signal_connect_state_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-set", G_CALLBACK(switch_stateSetHandler), data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File (GFile*) for param file

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

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// AttributeGetColumn is a wrapper around the C function gtk_cell_area_attribute_get_column.
func (recv *CellArea) AttributeGetColumn(renderer *CellRenderer, attribute string) int32 {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.gtk_cell_area_attribute_get_column((*C.GtkCellArea)(recv.native), c_renderer, c_attribute)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

type signalEntryCompletionNoMatchesDetail struct {
	callback  EntryCompletionSignalNoMatchesCallback
	handlerID C.gulong
}

var signalEntryCompletionNoMatchesId int
var signalEntryCompletionNoMatchesMap = make(map[int]signalEntryCompletionNoMatchesDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_no_matches(instance, C.gpointer(uintptr(signalEntryCompletionNoMatchesId)))

	detail := signalEntryCompletionNoMatchesDetail{callback, handlerID}
	signalEntryCompletionNoMatchesMap[signalEntryCompletionNoMatchesId] = detail

	return signalEntryCompletionNoMatchesId
}

/*
DisconnectNoMatches disconnects a callback from the 'no-matches' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectNoMatches.
*/
func (recv *EntryCompletion) DisconnectNoMatches(connectionID int) {
	signalEntryCompletionNoMatchesLock.Lock()
	defer signalEntryCompletionNoMatchesLock.Unlock()

	detail, exists := signalEntryCompletionNoMatchesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionNoMatchesMap, connectionID)
}

//export entrycompletion_noMatchesHandler
func entrycompletion_noMatchesHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalEntryCompletionNoMatchesMap[index].callback
	callback()
}

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

// Unsupported : gtk_event_controller_handle_event : unsupported parameter event : no type generator for Gdk.Event (const GdkEvent*) for param event

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

type signalGestureBeginDetail struct {
	callback  GestureSignalBeginCallback
	handlerID C.gulong
}

var signalGestureBeginId int
var signalGestureBeginMap = make(map[int]signalGestureBeginDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Gesture_signal_connect_begin(instance, C.gpointer(uintptr(signalGestureBeginId)))

	detail := signalGestureBeginDetail{callback, handlerID}
	signalGestureBeginMap[signalGestureBeginId] = detail

	return signalGestureBeginId
}

/*
DisconnectBegin disconnects a callback from the 'begin' signal for the Gesture.

The connectionID should be a value returned from a call to ConnectBegin.
*/
func (recv *Gesture) DisconnectBegin(connectionID int) {
	signalGestureBeginLock.Lock()
	defer signalGestureBeginLock.Unlock()

	detail, exists := signalGestureBeginMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureBeginMap, connectionID)
}

//export gesture_beginHandler
func gesture_beginHandler(_ *C.GObject, c_sequence *C.GdkEventSequence, data C.gpointer) {
	sequence := gdk.EventSequenceNewFromC(unsafe.Pointer(c_sequence))

	index := int(uintptr(data))
	callback := signalGestureBeginMap[index].callback
	callback(sequence)
}

type signalGestureCancelDetail struct {
	callback  GestureSignalCancelCallback
	handlerID C.gulong
}

var signalGestureCancelId int
var signalGestureCancelMap = make(map[int]signalGestureCancelDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Gesture_signal_connect_cancel(instance, C.gpointer(uintptr(signalGestureCancelId)))

	detail := signalGestureCancelDetail{callback, handlerID}
	signalGestureCancelMap[signalGestureCancelId] = detail

	return signalGestureCancelId
}

/*
DisconnectCancel disconnects a callback from the 'cancel' signal for the Gesture.

The connectionID should be a value returned from a call to ConnectCancel.
*/
func (recv *Gesture) DisconnectCancel(connectionID int) {
	signalGestureCancelLock.Lock()
	defer signalGestureCancelLock.Unlock()

	detail, exists := signalGestureCancelMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureCancelMap, connectionID)
}

//export gesture_cancelHandler
func gesture_cancelHandler(_ *C.GObject, c_sequence *C.GdkEventSequence, data C.gpointer) {
	sequence := gdk.EventSequenceNewFromC(unsafe.Pointer(c_sequence))

	index := int(uintptr(data))
	callback := signalGestureCancelMap[index].callback
	callback(sequence)
}

type signalGestureEndDetail struct {
	callback  GestureSignalEndCallback
	handlerID C.gulong
}

var signalGestureEndId int
var signalGestureEndMap = make(map[int]signalGestureEndDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Gesture_signal_connect_end(instance, C.gpointer(uintptr(signalGestureEndId)))

	detail := signalGestureEndDetail{callback, handlerID}
	signalGestureEndMap[signalGestureEndId] = detail

	return signalGestureEndId
}

/*
DisconnectEnd disconnects a callback from the 'end' signal for the Gesture.

The connectionID should be a value returned from a call to ConnectEnd.
*/
func (recv *Gesture) DisconnectEnd(connectionID int) {
	signalGestureEndLock.Lock()
	defer signalGestureEndLock.Unlock()

	detail, exists := signalGestureEndMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureEndMap, connectionID)
}

//export gesture_endHandler
func gesture_endHandler(_ *C.GObject, c_sequence *C.GdkEventSequence, data C.gpointer) {
	sequence := gdk.EventSequenceNewFromC(unsafe.Pointer(c_sequence))

	index := int(uintptr(data))
	callback := signalGestureEndMap[index].callback
	callback(sequence)
}

// Unsupported signal 'sequence-state-changed' for Gesture : unsupported parameter state : type EventSequenceState :

type signalGestureUpdateDetail struct {
	callback  GestureSignalUpdateCallback
	handlerID C.gulong
}

var signalGestureUpdateId int
var signalGestureUpdateMap = make(map[int]signalGestureUpdateDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Gesture_signal_connect_update(instance, C.gpointer(uintptr(signalGestureUpdateId)))

	detail := signalGestureUpdateDetail{callback, handlerID}
	signalGestureUpdateMap[signalGestureUpdateId] = detail

	return signalGestureUpdateId
}

/*
DisconnectUpdate disconnects a callback from the 'update' signal for the Gesture.

The connectionID should be a value returned from a call to ConnectUpdate.
*/
func (recv *Gesture) DisconnectUpdate(connectionID int) {
	signalGestureUpdateLock.Lock()
	defer signalGestureUpdateLock.Unlock()

	detail, exists := signalGestureUpdateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureUpdateMap, connectionID)
}

//export gesture_updateHandler
func gesture_updateHandler(_ *C.GObject, c_sequence *C.GdkEventSequence, data C.gpointer) {
	sequence := gdk.EventSequenceNewFromC(unsafe.Pointer(c_sequence))

	index := int(uintptr(data))
	callback := signalGestureUpdateMap[index].callback
	callback(sequence)
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
	var retGo (*gdk.Device)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.DeviceNewFromC(unsafe.Pointer(retC))
	}

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
	var retGo (*gdk.EventSequence)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.EventSequenceNewFromC(unsafe.Pointer(retC))
	}

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
	var retGo (*gdk.Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.WindowNewFromC(unsafe.Pointer(retC))
	}

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

// Unsupported signal 'drag-begin' for GestureDrag : unsupported parameter start_x : type gdouble :

// Unsupported signal 'drag-end' for GestureDrag : unsupported parameter offset_x : type gdouble :

// Unsupported signal 'drag-update' for GestureDrag : unsupported parameter offset_x : type gdouble :

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

type signalGestureLongPressCancelledDetail struct {
	callback  GestureLongPressSignalCancelledCallback
	handlerID C.gulong
}

var signalGestureLongPressCancelledId int
var signalGestureLongPressCancelledMap = make(map[int]signalGestureLongPressCancelledDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.GestureLongPress_signal_connect_cancelled(instance, C.gpointer(uintptr(signalGestureLongPressCancelledId)))

	detail := signalGestureLongPressCancelledDetail{callback, handlerID}
	signalGestureLongPressCancelledMap[signalGestureLongPressCancelledId] = detail

	return signalGestureLongPressCancelledId
}

/*
DisconnectCancelled disconnects a callback from the 'cancelled' signal for the GestureLongPress.

The connectionID should be a value returned from a call to ConnectCancelled.
*/
func (recv *GestureLongPress) DisconnectCancelled(connectionID int) {
	signalGestureLongPressCancelledLock.Lock()
	defer signalGestureLongPressCancelledLock.Unlock()

	detail, exists := signalGestureLongPressCancelledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureLongPressCancelledMap, connectionID)
}

//export gesturelongpress_cancelledHandler
func gesturelongpress_cancelledHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalGestureLongPressCancelledMap[index].callback
	callback()
}

// Unsupported signal 'pressed' for GestureLongPress : unsupported parameter x : type gdouble :

// GestureLongPressNew is a wrapper around the C function gtk_gesture_long_press_new.
func GestureLongPressNew(widget *Widget) *GestureLongPress {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_gesture_long_press_new(c_widget)
	retGo := GestureLongPressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'pressed' for GestureMultiPress : unsupported parameter n_press : type gint :

// Unsupported signal 'released' for GestureMultiPress : unsupported parameter n_press : type gint :

type signalGestureMultiPressStoppedDetail struct {
	callback  GestureMultiPressSignalStoppedCallback
	handlerID C.gulong
}

var signalGestureMultiPressStoppedId int
var signalGestureMultiPressStoppedMap = make(map[int]signalGestureMultiPressStoppedDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.GestureMultiPress_signal_connect_stopped(instance, C.gpointer(uintptr(signalGestureMultiPressStoppedId)))

	detail := signalGestureMultiPressStoppedDetail{callback, handlerID}
	signalGestureMultiPressStoppedMap[signalGestureMultiPressStoppedId] = detail

	return signalGestureMultiPressStoppedId
}

/*
DisconnectStopped disconnects a callback from the 'stopped' signal for the GestureMultiPress.

The connectionID should be a value returned from a call to ConnectStopped.
*/
func (recv *GestureMultiPress) DisconnectStopped(connectionID int) {
	signalGestureMultiPressStoppedLock.Lock()
	defer signalGestureMultiPressStoppedLock.Unlock()

	detail, exists := signalGestureMultiPressStoppedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGestureMultiPressStoppedMap, connectionID)
}

//export gesturemultipress_stoppedHandler
func gesturemultipress_stoppedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalGestureMultiPressStoppedMap[index].callback
	callback()
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

// Unsupported signal 'pan' for GesturePan : unsupported parameter direction : type PanDirection :

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

// Unsupported signal 'angle-changed' for GestureRotate : unsupported parameter angle : type gdouble :

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
	var retGo (*gdk.EventSequence)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.EventSequenceNewFromC(unsafe.Pointer(retC))
	}

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

// Unsupported signal 'swipe' for GestureSwipe : unsupported parameter velocity_x : type gdouble :

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

// Unsupported signal 'scale-changed' for GestureZoom : unsupported parameter scale : type gdouble :

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

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

type signalListBoxSelectAllDetail struct {
	callback  ListBoxSignalSelectAllCallback
	handlerID C.gulong
}

var signalListBoxSelectAllId int
var signalListBoxSelectAllMap = make(map[int]signalListBoxSelectAllDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_select_all(instance, C.gpointer(uintptr(signalListBoxSelectAllId)))

	detail := signalListBoxSelectAllDetail{callback, handlerID}
	signalListBoxSelectAllMap[signalListBoxSelectAllId] = detail

	return signalListBoxSelectAllId
}

/*
DisconnectSelectAll disconnects a callback from the 'select-all' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectSelectAll.
*/
func (recv *ListBox) DisconnectSelectAll(connectionID int) {
	signalListBoxSelectAllLock.Lock()
	defer signalListBoxSelectAllLock.Unlock()

	detail, exists := signalListBoxSelectAllMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxSelectAllMap, connectionID)
}

//export listbox_selectAllHandler
func listbox_selectAllHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalListBoxSelectAllMap[index].callback
	callback()
}

type signalListBoxSelectedRowsChangedDetail struct {
	callback  ListBoxSignalSelectedRowsChangedCallback
	handlerID C.gulong
}

var signalListBoxSelectedRowsChangedId int
var signalListBoxSelectedRowsChangedMap = make(map[int]signalListBoxSelectedRowsChangedDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_selected_rows_changed(instance, C.gpointer(uintptr(signalListBoxSelectedRowsChangedId)))

	detail := signalListBoxSelectedRowsChangedDetail{callback, handlerID}
	signalListBoxSelectedRowsChangedMap[signalListBoxSelectedRowsChangedId] = detail

	return signalListBoxSelectedRowsChangedId
}

/*
DisconnectSelectedRowsChanged disconnects a callback from the 'selected-rows-changed' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectSelectedRowsChanged.
*/
func (recv *ListBox) DisconnectSelectedRowsChanged(connectionID int) {
	signalListBoxSelectedRowsChangedLock.Lock()
	defer signalListBoxSelectedRowsChangedLock.Unlock()

	detail, exists := signalListBoxSelectedRowsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxSelectedRowsChangedMap, connectionID)
}

//export listbox_selectedRowsChangedHandler
func listbox_selectedRowsChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalListBoxSelectedRowsChangedMap[index].callback
	callback()
}

type signalListBoxUnselectAllDetail struct {
	callback  ListBoxSignalUnselectAllCallback
	handlerID C.gulong
}

var signalListBoxUnselectAllId int
var signalListBoxUnselectAllMap = make(map[int]signalListBoxUnselectAllDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_unselect_all(instance, C.gpointer(uintptr(signalListBoxUnselectAllId)))

	detail := signalListBoxUnselectAllDetail{callback, handlerID}
	signalListBoxUnselectAllMap[signalListBoxUnselectAllId] = detail

	return signalListBoxUnselectAllId
}

/*
DisconnectUnselectAll disconnects a callback from the 'unselect-all' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectUnselectAll.
*/
func (recv *ListBox) DisconnectUnselectAll(connectionID int) {
	signalListBoxUnselectAllLock.Lock()
	defer signalListBoxUnselectAllLock.Unlock()

	detail, exists := signalListBoxUnselectAllMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxUnselectAllMap, connectionID)
}

//export listbox_unselectAllHandler
func listbox_unselectAllHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalListBoxUnselectAllMap[index].callback
	callback()
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

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc (GtkListBoxForeachFunc) for param func

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

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

type signalPlacesSidebarShowEnterLocationDetail struct {
	callback  PlacesSidebarSignalShowEnterLocationCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowEnterLocationId int
var signalPlacesSidebarShowEnterLocationMap = make(map[int]signalPlacesSidebarShowEnterLocationDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_enter_location(instance, C.gpointer(uintptr(signalPlacesSidebarShowEnterLocationId)))

	detail := signalPlacesSidebarShowEnterLocationDetail{callback, handlerID}
	signalPlacesSidebarShowEnterLocationMap[signalPlacesSidebarShowEnterLocationId] = detail

	return signalPlacesSidebarShowEnterLocationId
}

/*
DisconnectShowEnterLocation disconnects a callback from the 'show-enter-location' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowEnterLocation.
*/
func (recv *PlacesSidebar) DisconnectShowEnterLocation(connectionID int) {
	signalPlacesSidebarShowEnterLocationLock.Lock()
	defer signalPlacesSidebarShowEnterLocationLock.Unlock()

	detail, exists := signalPlacesSidebarShowEnterLocationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowEnterLocationMap, connectionID)
}

//export placessidebar_showEnterLocationHandler
func placessidebar_showEnterLocationHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalPlacesSidebarShowEnterLocationMap[index].callback
	callback()
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

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

type signalSwitchStateSetDetail struct {
	callback  SwitchSignalStateSetCallback
	handlerID C.gulong
}

var signalSwitchStateSetId int
var signalSwitchStateSetMap = make(map[int]signalSwitchStateSetDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Switch_signal_connect_state_set(instance, C.gpointer(uintptr(signalSwitchStateSetId)))

	detail := signalSwitchStateSetDetail{callback, handlerID}
	signalSwitchStateSetMap[signalSwitchStateSetId] = detail

	return signalSwitchStateSetId
}

/*
DisconnectStateSet disconnects a callback from the 'state-set' signal for the Switch.

The connectionID should be a value returned from a call to ConnectStateSet.
*/
func (recv *Switch) DisconnectStateSet(connectionID int) {
	signalSwitchStateSetLock.Lock()
	defer signalSwitchStateSetLock.Unlock()

	detail, exists := signalSwitchStateSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSwitchStateSetMap, connectionID)
}

//export switch_stateSetHandler
func switch_stateSetHandler(_ *C.GObject, c_state C.gboolean, data C.gpointer) C.gboolean {
	state := c_state == C.TRUE

	index := int(uintptr(data))
	callback := signalSwitchStateSetMap[index].callback
	retGo := callback(state)
	retC :=
		boolToGboolean(retGo)
	return retC
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

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle
