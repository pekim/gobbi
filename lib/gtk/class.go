// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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

	void container_addHandler(GObject *, GtkWidget *, gpointer);

	static gulong Container_signal_connect_add(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add", G_CALLBACK(container_addHandler), data);
	}

*/
/*

	void container_checkResizeHandler(GObject *, gpointer);

	static gulong Container_signal_connect_check_resize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "check-resize", G_CALLBACK(container_checkResizeHandler), data);
	}

*/
/*

	void container_removeHandler(GObject *, GtkWidget *, gpointer);

	static gulong Container_signal_connect_remove(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove", G_CALLBACK(container_removeHandler), data);
	}

*/
/*

	void container_setFocusChildHandler(GObject *, GtkWidget *, gpointer);

	static gulong Container_signal_connect_set_focus_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "set-focus-child", G_CALLBACK(container_setFocusChildHandler), data);
	}

*/
/*

	void widget_accelClosuresChangedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_accel_closures_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-closures-changed", G_CALLBACK(widget_accelClosuresChangedHandler), data);
	}

*/
/*

	gboolean widget_buttonPressEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong Widget_signal_connect_button_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-press-event", G_CALLBACK(widget_buttonPressEventHandler), data);
	}

*/
/*

	gboolean widget_buttonReleaseEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong Widget_signal_connect_button_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-release-event", G_CALLBACK(widget_buttonReleaseEventHandler), data);
	}

*/
/*

	gboolean widget_canActivateAccelHandler(GObject *, guint, gpointer);

	static gulong Widget_signal_connect_can_activate_accel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "can-activate-accel", G_CALLBACK(widget_canActivateAccelHandler), data);
	}

*/
/*

	void widget_childNotifyHandler(GObject *, GParamSpec *, gpointer);

	static gulong Widget_signal_connect_child_notify(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "child-notify", G_CALLBACK(widget_childNotifyHandler), data);
	}

*/
/*

	void widget_compositedChangedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", G_CALLBACK(widget_compositedChangedHandler), data);
	}

*/
/*

	gboolean widget_configureEventHandler(GObject *, GdkEventConfigure *, gpointer);

	static gulong Widget_signal_connect_configure_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "configure-event", G_CALLBACK(widget_configureEventHandler), data);
	}

*/
/*

	void widget_destroyHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_destroy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "destroy", G_CALLBACK(widget_destroyHandler), data);
	}

*/
/*

	void widget_directionChangedHandler(GObject *, GtkTextDirection, gpointer);

	static gulong Widget_signal_connect_direction_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "direction-changed", G_CALLBACK(widget_directionChangedHandler), data);
	}

*/
/*

	void widget_dragBeginHandler(GObject *, GdkDragContext *, gpointer);

	static gulong Widget_signal_connect_drag_begin(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-begin", G_CALLBACK(widget_dragBeginHandler), data);
	}

*/
/*

	void widget_dragDataDeleteHandler(GObject *, GdkDragContext *, gpointer);

	static gulong Widget_signal_connect_drag_data_delete(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-delete", G_CALLBACK(widget_dragDataDeleteHandler), data);
	}

*/
/*

	void widget_dragDataGetHandler(GObject *, GdkDragContext *, GtkSelectionData *, guint, guint, gpointer);

	static gulong Widget_signal_connect_drag_data_get(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-get", G_CALLBACK(widget_dragDataGetHandler), data);
	}

*/
/*

	void widget_dragDataReceivedHandler(GObject *, GdkDragContext *, gint, gint, GtkSelectionData *, guint, guint, gpointer);

	static gulong Widget_signal_connect_drag_data_received(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-received", G_CALLBACK(widget_dragDataReceivedHandler), data);
	}

*/
/*

	gboolean widget_dragDropHandler(GObject *, GdkDragContext *, gint, gint, guint, gpointer);

	static gulong Widget_signal_connect_drag_drop(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-drop", G_CALLBACK(widget_dragDropHandler), data);
	}

*/
/*

	void widget_dragEndHandler(GObject *, GdkDragContext *, gpointer);

	static gulong Widget_signal_connect_drag_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-end", G_CALLBACK(widget_dragEndHandler), data);
	}

*/
/*

	void widget_dragLeaveHandler(GObject *, GdkDragContext *, guint, gpointer);

	static gulong Widget_signal_connect_drag_leave(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-leave", G_CALLBACK(widget_dragLeaveHandler), data);
	}

*/
/*

	gboolean widget_dragMotionHandler(GObject *, GdkDragContext *, gint, gint, guint, gpointer);

	static gulong Widget_signal_connect_drag_motion(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-motion", G_CALLBACK(widget_dragMotionHandler), data);
	}

*/
/*

	gboolean widget_enterNotifyEventHandler(GObject *, GdkEventCrossing *, gpointer);

	static gulong Widget_signal_connect_enter_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enter-notify-event", G_CALLBACK(widget_enterNotifyEventHandler), data);
	}

*/
/*

	gboolean widget_focusHandler(GObject *, GtkDirectionType, gpointer);

	static gulong Widget_signal_connect_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus", G_CALLBACK(widget_focusHandler), data);
	}

*/
/*

	gboolean widget_focusInEventHandler(GObject *, GdkEventFocus *, gpointer);

	static gulong Widget_signal_connect_focus_in_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-in-event", G_CALLBACK(widget_focusInEventHandler), data);
	}

*/
/*

	gboolean widget_focusOutEventHandler(GObject *, GdkEventFocus *, gpointer);

	static gulong Widget_signal_connect_focus_out_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-out-event", G_CALLBACK(widget_focusOutEventHandler), data);
	}

*/
/*

	void widget_grabFocusHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_grab_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-focus", G_CALLBACK(widget_grabFocusHandler), data);
	}

*/
/*

	void widget_grabNotifyHandler(GObject *, gboolean, gpointer);

	static gulong Widget_signal_connect_grab_notify(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-notify", G_CALLBACK(widget_grabNotifyHandler), data);
	}

*/
/*

	void widget_hideHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_hide(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "hide", G_CALLBACK(widget_hideHandler), data);
	}

*/
/*

	void widget_hierarchyChangedHandler(GObject *, GtkWidget *, gpointer);

	static gulong Widget_signal_connect_hierarchy_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "hierarchy-changed", G_CALLBACK(widget_hierarchyChangedHandler), data);
	}

*/
/*

	gboolean widget_keyPressEventHandler(GObject *, GdkEventKey *, gpointer);

	static gulong Widget_signal_connect_key_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "key-press-event", G_CALLBACK(widget_keyPressEventHandler), data);
	}

*/
/*

	gboolean widget_keyReleaseEventHandler(GObject *, GdkEventKey *, gpointer);

	static gulong Widget_signal_connect_key_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "key-release-event", G_CALLBACK(widget_keyReleaseEventHandler), data);
	}

*/
/*

	gboolean widget_leaveNotifyEventHandler(GObject *, GdkEventCrossing *, gpointer);

	static gulong Widget_signal_connect_leave_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "leave-notify-event", G_CALLBACK(widget_leaveNotifyEventHandler), data);
	}

*/
/*

	void widget_mapHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_map(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "map", G_CALLBACK(widget_mapHandler), data);
	}

*/
/*

	gboolean widget_mapEventHandler(GObject *, GdkEventAny *, gpointer);

	static gulong Widget_signal_connect_map_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "map-event", G_CALLBACK(widget_mapEventHandler), data);
	}

*/
/*

	gboolean widget_mnemonicActivateHandler(GObject *, gboolean, gpointer);

	static gulong Widget_signal_connect_mnemonic_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mnemonic-activate", G_CALLBACK(widget_mnemonicActivateHandler), data);
	}

*/
/*

	gboolean widget_motionNotifyEventHandler(GObject *, GdkEventMotion *, gpointer);

	static gulong Widget_signal_connect_motion_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "motion-notify-event", G_CALLBACK(widget_motionNotifyEventHandler), data);
	}

*/
/*

	void widget_moveFocusHandler(GObject *, GtkDirectionType, gpointer);

	static gulong Widget_signal_connect_move_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-focus", G_CALLBACK(widget_moveFocusHandler), data);
	}

*/
/*

	void widget_parentSetHandler(GObject *, GtkWidget *, gpointer);

	static gulong Widget_signal_connect_parent_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "parent-set", G_CALLBACK(widget_parentSetHandler), data);
	}

*/
/*

	gboolean widget_popupMenuHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_popup_menu(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup-menu", G_CALLBACK(widget_popupMenuHandler), data);
	}

*/
/*

	gboolean widget_propertyNotifyEventHandler(GObject *, GdkEventProperty *, gpointer);

	static gulong Widget_signal_connect_property_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "property-notify-event", G_CALLBACK(widget_propertyNotifyEventHandler), data);
	}

*/
/*

	gboolean widget_proximityInEventHandler(GObject *, GdkEventProximity *, gpointer);

	static gulong Widget_signal_connect_proximity_in_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "proximity-in-event", G_CALLBACK(widget_proximityInEventHandler), data);
	}

*/
/*

	gboolean widget_proximityOutEventHandler(GObject *, GdkEventProximity *, gpointer);

	static gulong Widget_signal_connect_proximity_out_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "proximity-out-event", G_CALLBACK(widget_proximityOutEventHandler), data);
	}

*/
/*

	void widget_realizeHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_realize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "realize", G_CALLBACK(widget_realizeHandler), data);
	}

*/
/*

	void widget_screenChangedHandler(GObject *, GdkScreen *, gpointer);

	static gulong Widget_signal_connect_screen_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "screen-changed", G_CALLBACK(widget_screenChangedHandler), data);
	}

*/
/*

	gboolean widget_scrollEventHandler(GObject *, GdkEventScroll *, gpointer);

	static gulong Widget_signal_connect_scroll_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-event", G_CALLBACK(widget_scrollEventHandler), data);
	}

*/
/*

	gboolean widget_selectionClearEventHandler(GObject *, GdkEventSelection *, gpointer);

	static gulong Widget_signal_connect_selection_clear_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-clear-event", G_CALLBACK(widget_selectionClearEventHandler), data);
	}

*/
/*

	void widget_selectionGetHandler(GObject *, GtkSelectionData *, guint, guint, gpointer);

	static gulong Widget_signal_connect_selection_get(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-get", G_CALLBACK(widget_selectionGetHandler), data);
	}

*/
/*

	gboolean widget_selectionNotifyEventHandler(GObject *, GdkEventSelection *, gpointer);

	static gulong Widget_signal_connect_selection_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-notify-event", G_CALLBACK(widget_selectionNotifyEventHandler), data);
	}

*/
/*

	void widget_selectionReceivedHandler(GObject *, GtkSelectionData *, guint, gpointer);

	static gulong Widget_signal_connect_selection_received(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-received", G_CALLBACK(widget_selectionReceivedHandler), data);
	}

*/
/*

	gboolean widget_selectionRequestEventHandler(GObject *, GdkEventSelection *, gpointer);

	static gulong Widget_signal_connect_selection_request_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-request-event", G_CALLBACK(widget_selectionRequestEventHandler), data);
	}

*/
/*

	void widget_showHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_show(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show", G_CALLBACK(widget_showHandler), data);
	}

*/
/*

	gboolean widget_showHelpHandler(GObject *, GtkWidgetHelpType, gpointer);

	static gulong Widget_signal_connect_show_help(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-help", G_CALLBACK(widget_showHelpHandler), data);
	}

*/
/*

	void widget_sizeAllocateHandler(GObject *, GdkRectangle *, gpointer);

	static gulong Widget_signal_connect_size_allocate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-allocate", G_CALLBACK(widget_sizeAllocateHandler), data);
	}

*/
/*

	void widget_stateChangedHandler(GObject *, GtkStateType, gpointer);

	static gulong Widget_signal_connect_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-changed", G_CALLBACK(widget_stateChangedHandler), data);
	}

*/
/*

	void widget_styleSetHandler(GObject *, GtkStyle *, gpointer);

	static gulong Widget_signal_connect_style_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-set", G_CALLBACK(widget_styleSetHandler), data);
	}

*/
/*

	void widget_unmapHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_unmap(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmap", G_CALLBACK(widget_unmapHandler), data);
	}

*/
/*

	gboolean widget_unmapEventHandler(GObject *, GdkEventAny *, gpointer);

	static gulong Widget_signal_connect_unmap_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmap-event", G_CALLBACK(widget_unmapEventHandler), data);
	}

*/
/*

	void widget_unrealizeHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_unrealize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unrealize", G_CALLBACK(widget_unrealizeHandler), data);
	}

*/
/*

	gboolean widget_visibilityNotifyEventHandler(GObject *, GdkEventVisibility *, gpointer);

	static gulong Widget_signal_connect_visibility_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "visibility-notify-event", G_CALLBACK(widget_visibilityNotifyEventHandler), data);
	}

*/
/*

	gboolean widget_windowStateEventHandler(GObject *, GdkEventWindowState *, gpointer);

	static gulong Widget_signal_connect_window_state_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-state-event", G_CALLBACK(widget_windowStateEventHandler), data);
	}

*/
/*

	void window_activateDefaultHandler(GObject *, gpointer);

	static gulong Window_signal_connect_activate_default(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-default", G_CALLBACK(window_activateDefaultHandler), data);
	}

*/
/*

	void window_activateFocusHandler(GObject *, gpointer);

	static gulong Window_signal_connect_activate_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-focus", G_CALLBACK(window_activateFocusHandler), data);
	}

*/
/*

	gboolean window_enableDebuggingHandler(GObject *, gboolean, gpointer);

	static gulong Window_signal_connect_enable_debugging(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enable-debugging", G_CALLBACK(window_enableDebuggingHandler), data);
	}

*/
/*

	void window_keysChangedHandler(GObject *, gpointer);

	static gulong Window_signal_connect_keys_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keys-changed", G_CALLBACK(window_keysChangedHandler), data);
	}

*/
/*

	void window_setFocusHandler(GObject *, GtkWidget *, gpointer);

	static gulong Window_signal_connect_set_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "set-focus", G_CALLBACK(window_setFocusHandler), data);
	}

*/
import "C"

// Blacklisted : GtkAboutDialog

// Blacklisted : GtkAccelGroup

// Blacklisted : GtkAccelLabel

// Blacklisted : GtkAccelMap

// Blacklisted : GtkAccessible

// Action is a wrapper around the C record GtkAction.
type Action struct {
	native *C.GtkAction
	// object : record
	// Private : private_data
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Action) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Action) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Action.
// Exercise care, as this is a potentially dangerous function if the Object is not a Action.
func CastToAction(object *gobject.Object) *Action {
	return ActionNewFromC(object.ToC())
}

// Buildable returns the Buildable interface implemented by Action
func (recv *Action) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkActionBar

// ActionGroup is a wrapper around the C record GtkActionGroup.
type ActionGroup struct {
	native *C.GtkActionGroup
	// parent : record
	// Private : priv
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GtkActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionGroup with another ActionGroup, and returns true if they represent the same GObject.
func (recv *ActionGroup) Equals(other *ActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ActionGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a ActionGroup.
func CastToActionGroup(object *gobject.Object) *ActionGroup {
	return ActionGroupNewFromC(object.ToC())
}

// Buildable returns the Buildable interface implemented by ActionGroup
func (recv *ActionGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkAdjustment

// Blacklisted : GtkAlignment

// Blacklisted : GtkAppChooserButton

// Blacklisted : GtkAppChooserDialog

// Blacklisted : GtkAppChooserWidget

// Application is a wrapper around the C record GtkApplication.
type Application struct {
	native *C.GtkApplication
	// parent : record
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GtkApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Application) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Application with another Application, and returns true if they represent the same GObject.
func (recv *Application) Equals(other *Application) bool {
	return other.ToC() == recv.ToC()
}

// Application upcasts to *Application
func (recv *Application) Application() *gio.Application {
	return gio.ApplicationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Application) Object() *gobject.Object {
	return recv.Application().Object()
}

// CastToWidget down casts any arbitrary Object to Application.
// Exercise care, as this is a potentially dangerous function if the Object is not a Application.
func CastToApplication(object *gobject.Object) *Application {
	return ApplicationNewFromC(object.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by Application
func (recv *Application) ActionGroup() *gio.ActionGroup {
	return gio.ActionGroupNewFromC(recv.ToC())
}

// ActionMap returns the ActionMap interface implemented by Application
func (recv *Application) ActionMap() *gio.ActionMap {
	return gio.ActionMapNewFromC(recv.ToC())
}

// Blacklisted : GtkApplicationWindow

// Blacklisted : GtkArrow

// Blacklisted : GtkArrowAccessible

// Blacklisted : GtkAspectFrame

// Blacklisted : GtkAssistant

// Bin is a wrapper around the C record GtkBin.
type Bin struct {
	native *C.GtkBin
	// container : record
	// Private : priv
}

func BinNewFromC(u unsafe.Pointer) *Bin {
	c := (*C.GtkBin)(u)
	if c == nil {
		return nil
	}

	g := &Bin{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Bin) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Bin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Bin with another Bin, and returns true if they represent the same GObject.
func (recv *Bin) Equals(other *Bin) bool {
	return other.ToC() == recv.ToC()
}

// Container upcasts to *Container
func (recv *Bin) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Bin) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Bin) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Bin) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitrary Object to Bin.
// Exercise care, as this is a potentially dangerous function if the Object is not a Bin.
func CastToBin(object *gobject.Object) *Bin {
	return BinNewFromC(object.ToC())
}

// Blacklisted : gtk_bin_get_child

// ImplementorIface returns the ImplementorIface interface implemented by Bin
func (recv *Bin) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Bin
func (recv *Bin) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkBooleanCellAccessible

// Blacklisted : GtkBox

// Blacklisted : GtkBuilder

// Blacklisted : GtkButton

// Blacklisted : GtkButtonAccessible

// Blacklisted : GtkButtonBox

// Blacklisted : GtkCalendar

// Blacklisted : GtkCellAccessible

// Blacklisted : GtkCellArea

// Blacklisted : GtkCellAreaBox

// Blacklisted : GtkCellAreaContext

// Blacklisted : GtkCellRenderer

// Blacklisted : GtkCellRendererAccel

// Blacklisted : GtkCellRendererCombo

// Blacklisted : GtkCellRendererPixbuf

// Blacklisted : GtkCellRendererProgress

// Blacklisted : GtkCellRendererSpin

// Blacklisted : GtkCellRendererSpinner

// Blacklisted : GtkCellRendererText

// Blacklisted : GtkCellRendererToggle

// Blacklisted : GtkCellView

// Blacklisted : GtkCheckButton

// Blacklisted : GtkCheckMenuItem

// Blacklisted : GtkCheckMenuItemAccessible

// Blacklisted : GtkClipboard

// Blacklisted : GtkColorButton

// Blacklisted : GtkColorChooserDialog

// Blacklisted : GtkColorChooserWidget

// Blacklisted : GtkColorSelection

// Blacklisted : GtkColorSelectionDialog

// Blacklisted : GtkComboBox

// Blacklisted : GtkComboBoxAccessible

// Blacklisted : GtkComboBoxText

// Container is a wrapper around the C record GtkContainer.
type Container struct {
	native *C.GtkContainer
	// widget : record
	// Private : priv
}

func ContainerNewFromC(u unsafe.Pointer) *Container {
	c := (*C.GtkContainer)(u)
	if c == nil {
		return nil
	}

	g := &Container{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Container) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Container) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Container with another Container, and returns true if they represent the same GObject.
func (recv *Container) Equals(other *Container) bool {
	return other.ToC() == recv.ToC()
}

// Widget upcasts to *Widget
func (recv *Container) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Container) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Container) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitrary Object to Container.
// Exercise care, as this is a potentially dangerous function if the Object is not a Container.
func CastToContainer(object *gobject.Object) *Container {
	return ContainerNewFromC(object.ToC())
}

type signalContainerAddDetail struct {
	callback  ContainerSignalAddCallback
	handlerID C.gulong
}

var signalContainerAddId int
var signalContainerAddMap = make(map[int]signalContainerAddDetail)
var signalContainerAddLock sync.RWMutex

// ContainerSignalAddCallback is a callback function for a 'add' signal emitted from a Container.
type ContainerSignalAddCallback func(object *Widget)

/*
ConnectAdd connects the callback to the 'add' signal for the Container.

The returned value represents the connection, and may be passed to DisconnectAdd to remove it.
*/
func (recv *Container) ConnectAdd(callback ContainerSignalAddCallback) int {
	signalContainerAddLock.Lock()
	defer signalContainerAddLock.Unlock()

	signalContainerAddId++
	instance := C.gpointer(recv.native)
	handlerID := C.Container_signal_connect_add(instance, C.gpointer(uintptr(signalContainerAddId)))

	detail := signalContainerAddDetail{callback, handlerID}
	signalContainerAddMap[signalContainerAddId] = detail

	return signalContainerAddId
}

/*
DisconnectAdd disconnects a callback from the 'add' signal for the Container.

The connectionID should be a value returned from a call to ConnectAdd.
*/
func (recv *Container) DisconnectAdd(connectionID int) {
	signalContainerAddLock.Lock()
	defer signalContainerAddLock.Unlock()

	detail, exists := signalContainerAddMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalContainerAddMap, connectionID)
}

//export container_addHandler
func container_addHandler(_ *C.GObject, c_object *C.GtkWidget, data C.gpointer) {
	signalContainerAddLock.RLock()
	defer signalContainerAddLock.RUnlock()

	object := WidgetNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalContainerAddMap[index].callback
	callback(object)
}

type signalContainerCheckResizeDetail struct {
	callback  ContainerSignalCheckResizeCallback
	handlerID C.gulong
}

var signalContainerCheckResizeId int
var signalContainerCheckResizeMap = make(map[int]signalContainerCheckResizeDetail)
var signalContainerCheckResizeLock sync.RWMutex

// ContainerSignalCheckResizeCallback is a callback function for a 'check-resize' signal emitted from a Container.
type ContainerSignalCheckResizeCallback func()

/*
ConnectCheckResize connects the callback to the 'check-resize' signal for the Container.

The returned value represents the connection, and may be passed to DisconnectCheckResize to remove it.
*/
func (recv *Container) ConnectCheckResize(callback ContainerSignalCheckResizeCallback) int {
	signalContainerCheckResizeLock.Lock()
	defer signalContainerCheckResizeLock.Unlock()

	signalContainerCheckResizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Container_signal_connect_check_resize(instance, C.gpointer(uintptr(signalContainerCheckResizeId)))

	detail := signalContainerCheckResizeDetail{callback, handlerID}
	signalContainerCheckResizeMap[signalContainerCheckResizeId] = detail

	return signalContainerCheckResizeId
}

/*
DisconnectCheckResize disconnects a callback from the 'check-resize' signal for the Container.

The connectionID should be a value returned from a call to ConnectCheckResize.
*/
func (recv *Container) DisconnectCheckResize(connectionID int) {
	signalContainerCheckResizeLock.Lock()
	defer signalContainerCheckResizeLock.Unlock()

	detail, exists := signalContainerCheckResizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalContainerCheckResizeMap, connectionID)
}

//export container_checkResizeHandler
func container_checkResizeHandler(_ *C.GObject, data C.gpointer) {
	signalContainerCheckResizeLock.RLock()
	defer signalContainerCheckResizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalContainerCheckResizeMap[index].callback
	callback()
}

type signalContainerRemoveDetail struct {
	callback  ContainerSignalRemoveCallback
	handlerID C.gulong
}

var signalContainerRemoveId int
var signalContainerRemoveMap = make(map[int]signalContainerRemoveDetail)
var signalContainerRemoveLock sync.RWMutex

// ContainerSignalRemoveCallback is a callback function for a 'remove' signal emitted from a Container.
type ContainerSignalRemoveCallback func(object *Widget)

/*
ConnectRemove connects the callback to the 'remove' signal for the Container.

The returned value represents the connection, and may be passed to DisconnectRemove to remove it.
*/
func (recv *Container) ConnectRemove(callback ContainerSignalRemoveCallback) int {
	signalContainerRemoveLock.Lock()
	defer signalContainerRemoveLock.Unlock()

	signalContainerRemoveId++
	instance := C.gpointer(recv.native)
	handlerID := C.Container_signal_connect_remove(instance, C.gpointer(uintptr(signalContainerRemoveId)))

	detail := signalContainerRemoveDetail{callback, handlerID}
	signalContainerRemoveMap[signalContainerRemoveId] = detail

	return signalContainerRemoveId
}

/*
DisconnectRemove disconnects a callback from the 'remove' signal for the Container.

The connectionID should be a value returned from a call to ConnectRemove.
*/
func (recv *Container) DisconnectRemove(connectionID int) {
	signalContainerRemoveLock.Lock()
	defer signalContainerRemoveLock.Unlock()

	detail, exists := signalContainerRemoveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalContainerRemoveMap, connectionID)
}

//export container_removeHandler
func container_removeHandler(_ *C.GObject, c_object *C.GtkWidget, data C.gpointer) {
	signalContainerRemoveLock.RLock()
	defer signalContainerRemoveLock.RUnlock()

	object := WidgetNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalContainerRemoveMap[index].callback
	callback(object)
}

type signalContainerSetFocusChildDetail struct {
	callback  ContainerSignalSetFocusChildCallback
	handlerID C.gulong
}

var signalContainerSetFocusChildId int
var signalContainerSetFocusChildMap = make(map[int]signalContainerSetFocusChildDetail)
var signalContainerSetFocusChildLock sync.RWMutex

// ContainerSignalSetFocusChildCallback is a callback function for a 'set-focus-child' signal emitted from a Container.
type ContainerSignalSetFocusChildCallback func(object *Widget)

/*
ConnectSetFocusChild connects the callback to the 'set-focus-child' signal for the Container.

The returned value represents the connection, and may be passed to DisconnectSetFocusChild to remove it.
*/
func (recv *Container) ConnectSetFocusChild(callback ContainerSignalSetFocusChildCallback) int {
	signalContainerSetFocusChildLock.Lock()
	defer signalContainerSetFocusChildLock.Unlock()

	signalContainerSetFocusChildId++
	instance := C.gpointer(recv.native)
	handlerID := C.Container_signal_connect_set_focus_child(instance, C.gpointer(uintptr(signalContainerSetFocusChildId)))

	detail := signalContainerSetFocusChildDetail{callback, handlerID}
	signalContainerSetFocusChildMap[signalContainerSetFocusChildId] = detail

	return signalContainerSetFocusChildId
}

/*
DisconnectSetFocusChild disconnects a callback from the 'set-focus-child' signal for the Container.

The connectionID should be a value returned from a call to ConnectSetFocusChild.
*/
func (recv *Container) DisconnectSetFocusChild(connectionID int) {
	signalContainerSetFocusChildLock.Lock()
	defer signalContainerSetFocusChildLock.Unlock()

	detail, exists := signalContainerSetFocusChildMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalContainerSetFocusChildMap, connectionID)
}

//export container_setFocusChildHandler
func container_setFocusChildHandler(_ *C.GObject, c_object *C.GtkWidget, data C.gpointer) {
	signalContainerSetFocusChildLock.RLock()
	defer signalContainerSetFocusChildLock.RUnlock()

	object := WidgetNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalContainerSetFocusChildMap[index].callback
	callback(object)
}

// Blacklisted : gtk_container_add

// Unsupported : gtk_container_add_with_properties : unsupported parameter ... : varargs

// Blacklisted : gtk_container_check_resize

// Unsupported : gtk_container_child_get : unsupported parameter ... : varargs

// Blacklisted : gtk_container_child_get_property

// Unsupported : gtk_container_child_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : gtk_container_child_set : unsupported parameter ... : varargs

// Blacklisted : gtk_container_child_set_property

// Unsupported : gtk_container_child_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : gtk_container_child_type

// Unsupported : gtk_container_forall : unsupported parameter callback : no type generator for Callback (GtkCallback) for param callback

// Unsupported : gtk_container_foreach : unsupported parameter callback : no type generator for Callback (GtkCallback) for param callback

// Blacklisted : gtk_container_get_border_width

// Blacklisted : gtk_container_get_children

// Blacklisted : gtk_container_get_focus_chain

// Blacklisted : gtk_container_get_focus_hadjustment

// Blacklisted : gtk_container_get_focus_vadjustment

// Blacklisted : gtk_container_get_path_for_child

// Blacklisted : gtk_container_get_resize_mode

// Blacklisted : gtk_container_propagate_draw

// Blacklisted : gtk_container_remove

// Blacklisted : gtk_container_resize_children

// Blacklisted : gtk_container_set_border_width

// Blacklisted : gtk_container_set_focus_chain

// Blacklisted : gtk_container_set_focus_child

// Blacklisted : gtk_container_set_focus_hadjustment

// Blacklisted : gtk_container_set_focus_vadjustment

// Blacklisted : gtk_container_set_reallocate_redraws

// Blacklisted : gtk_container_set_resize_mode

// Blacklisted : gtk_container_unset_focus_chain

// ImplementorIface returns the ImplementorIface interface implemented by Container
func (recv *Container) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Container
func (recv *Container) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkContainerAccessible

// Blacklisted : GtkContainerCellAccessible

// Blacklisted : GtkCssProvider

// Blacklisted : GtkDialog

// Blacklisted : GtkDrawingArea

// Blacklisted : GtkEntry

// Blacklisted : GtkEntryAccessible

// Blacklisted : GtkEntryBuffer

// Blacklisted : GtkEntryCompletion

// Unsupported : EntryIconAccessible : no CType

// Blacklisted : GtkEventBox

// Blacklisted : GtkEventController

// Blacklisted : GtkExpander

// Blacklisted : GtkExpanderAccessible

// Blacklisted : GtkFileChooserButton

// Blacklisted : GtkFileChooserDialog

// Blacklisted : GtkFileChooserWidget

// Blacklisted : GtkFileFilter

// Blacklisted : GtkFixed

// Blacklisted : GtkFlowBox

// Blacklisted : GtkFlowBoxAccessible

// Blacklisted : GtkFlowBoxChild

// Blacklisted : GtkFlowBoxChildAccessible

// Blacklisted : GtkFontButton

// Blacklisted : GtkFontChooserDialog

// Blacklisted : GtkFontChooserWidget

// Blacklisted : GtkFontSelection

// Blacklisted : GtkFontSelectionDialog

// Blacklisted : GtkFrame

// Blacklisted : GtkFrameAccessible

// Blacklisted : GtkGesture

// Blacklisted : GtkGestureDrag

// Blacklisted : GtkGestureLongPress

// Blacklisted : GtkGestureMultiPress

// Blacklisted : GtkGesturePan

// Blacklisted : GtkGestureRotate

// Blacklisted : GtkGestureSingle

// Blacklisted : GtkGestureSwipe

// Blacklisted : GtkGestureZoom

// Blacklisted : GtkGrid

// Blacklisted : GtkHBox

// Blacklisted : GtkHButtonBox

// Blacklisted : GtkHPaned

// Blacklisted : GtkHSV

// Blacklisted : GtkHScale

// Blacklisted : GtkHScrollbar

// Blacklisted : GtkHSeparator

// Blacklisted : GtkHandleBox

// Blacklisted : GtkHeaderBar

// Blacklisted : GtkIMContext

// Blacklisted : GtkIMContextSimple

// Blacklisted : GtkIMMulticontext

// Blacklisted : GtkIconFactory

// Blacklisted : GtkIconInfo

// Blacklisted : GtkIconTheme

// Blacklisted : GtkIconView

// Blacklisted : GtkIconViewAccessible

// Blacklisted : GtkImage

// Blacklisted : GtkImageAccessible

// Blacklisted : GtkImageCellAccessible

// Blacklisted : GtkImageMenuItem

// Blacklisted : GtkInfoBar

// Blacklisted : GtkInvisible

// Blacklisted : GtkLabel

// Blacklisted : GtkLabelAccessible

// Blacklisted : GtkLayout

// Blacklisted : GtkLevelBar

// Blacklisted : GtkLevelBarAccessible

// Blacklisted : GtkLinkButton

// Blacklisted : GtkLinkButtonAccessible

// Blacklisted : GtkListBox

// Blacklisted : GtkListBoxAccessible

// Blacklisted : GtkListBoxRow

// Blacklisted : GtkListBoxRowAccessible

// Blacklisted : GtkListStore

// Blacklisted : GtkLockButton

// Blacklisted : GtkLockButtonAccessible

// Blacklisted : GtkMenu

// Blacklisted : GtkMenuAccessible

// Blacklisted : GtkMenuBar

// Blacklisted : GtkMenuButton

// Blacklisted : GtkMenuButtonAccessible

// Blacklisted : GtkMenuItem

// Blacklisted : GtkMenuItemAccessible

// Blacklisted : GtkMenuShell

// Blacklisted : GtkMenuShellAccessible

// Blacklisted : GtkMenuToolButton

// Blacklisted : GtkMessageDialog

// Blacklisted : GtkMisc

// Blacklisted : GtkModelButton

// Blacklisted : GtkMountOperation

// Blacklisted : GtkNotebook

// Blacklisted : GtkNotebookAccessible

// Blacklisted : GtkNotebookPageAccessible

// Blacklisted : GtkNumerableIcon

// Blacklisted : GtkOffscreenWindow

// Blacklisted : GtkOverlay

// Blacklisted : GtkPageSetup

// Blacklisted : GtkPaned

// Blacklisted : GtkPanedAccessible

// Blacklisted : GtkPlacesSidebar

// Blacklisted : GtkPlug

// Blacklisted : GtkPopover

// Blacklisted : GtkPopoverAccessible

// Blacklisted : GtkPopoverMenu

// Blacklisted : GtkPrintContext

// Blacklisted : GtkPrintOperation

// Blacklisted : GtkPrintSettings

// Blacklisted : GtkProgressBar

// Blacklisted : GtkProgressBarAccessible

// Blacklisted : GtkRadioAction

// Blacklisted : GtkRadioButton

// Blacklisted : GtkRadioButtonAccessible

// Blacklisted : GtkRadioMenuItem

// Blacklisted : GtkRadioMenuItemAccessible

// Blacklisted : GtkRadioToolButton

// Blacklisted : GtkRange

// Blacklisted : GtkRangeAccessible

// Blacklisted : GtkRcStyle

// Blacklisted : GtkRecentAction

// Blacklisted : GtkRecentChooserDialog

// Blacklisted : GtkRecentChooserMenu

// Blacklisted : GtkRecentChooserWidget

// Blacklisted : GtkRecentFilter

// Blacklisted : GtkRendererCellAccessible

// Blacklisted : GtkRevealer

// Blacklisted : GtkScale

// Blacklisted : GtkScaleAccessible

// Blacklisted : GtkScaleButton

// Blacklisted : GtkScaleButtonAccessible

// Blacklisted : GtkScrollbar

// Blacklisted : GtkScrolledWindow

// Blacklisted : GtkScrolledWindowAccessible

// Blacklisted : GtkSearchBar

// Blacklisted : GtkSearchEntry

// Blacklisted : GtkSeparator

// Blacklisted : GtkSeparatorMenuItem

// Blacklisted : GtkSeparatorToolItem

// Blacklisted : GtkSettings

// Blacklisted : GtkSizeGroup

// Blacklisted : GtkSocket

// Blacklisted : GtkSpinButton

// Blacklisted : GtkSpinButtonAccessible

// Blacklisted : GtkSpinner

// Blacklisted : GtkSpinnerAccessible

// Blacklisted : GtkStack

// Blacklisted : GtkStackAccessible

// Blacklisted : GtkStackSidebar

// Blacklisted : GtkStackSwitcher

// Blacklisted : GtkStatusIcon

// Blacklisted : GtkStatusbar

// Blacklisted : GtkStatusbarAccessible

// Blacklisted : GtkStyle

// Blacklisted : GtkStyleContext

// Blacklisted : GtkStyleProperties

// Blacklisted : GtkSwitch

// Blacklisted : GtkSwitchAccessible

// Blacklisted : GtkTable

// Blacklisted : GtkTearoffMenuItem

// Blacklisted : GtkTextBuffer

// Blacklisted : GtkTextCellAccessible

// Blacklisted : GtkTextChildAnchor

// Blacklisted : GtkTextMark

// Blacklisted : GtkTextTag

// Blacklisted : GtkTextTagTable

// Blacklisted : GtkTextView

// Blacklisted : GtkTextViewAccessible

// Blacklisted : GtkThemingEngine

// Blacklisted : GtkToggleAction

// Blacklisted : GtkToggleButton

// Blacklisted : GtkToggleButtonAccessible

// Blacklisted : GtkToggleToolButton

// Blacklisted : GtkToolButton

// Blacklisted : GtkToolItem

// Blacklisted : GtkToolItemGroup

// Blacklisted : GtkToolPalette

// Blacklisted : GtkToolbar

// Tooltip is a wrapper around the C record GtkTooltip.
type Tooltip struct {
	native *C.GtkTooltip
}

func TooltipNewFromC(u unsafe.Pointer) *Tooltip {
	c := (*C.GtkTooltip)(u)
	if c == nil {
		return nil
	}

	g := &Tooltip{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Tooltip) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Tooltip) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Tooltip with another Tooltip, and returns true if they represent the same GObject.
func (recv *Tooltip) Equals(other *Tooltip) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Tooltip) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Tooltip.
// Exercise care, as this is a potentially dangerous function if the Object is not a Tooltip.
func CastToTooltip(object *gobject.Object) *Tooltip {
	return TooltipNewFromC(object.ToC())
}

// Blacklisted : GtkToplevelAccessible

// Blacklisted : GtkTreeModelFilter

// Blacklisted : GtkTreeModelSort

// Blacklisted : GtkTreeSelection

// Blacklisted : GtkTreeStore

// Blacklisted : GtkTreeView

// Blacklisted : GtkTreeViewAccessible

// Blacklisted : GtkTreeViewColumn

// Blacklisted : GtkUIManager

// Blacklisted : GtkVBox

// Blacklisted : GtkVButtonBox

// Blacklisted : GtkVPaned

// Blacklisted : GtkVScale

// Blacklisted : GtkVScrollbar

// Blacklisted : GtkVSeparator

// Blacklisted : GtkViewport

// Blacklisted : GtkVolumeButton

// Widget is a wrapper around the C record GtkWidget.
type Widget struct {
	native *C.GtkWidget
	// parent_instance : record
	// Private : priv
}

func WidgetNewFromC(u unsafe.Pointer) *Widget {
	c := (*C.GtkWidget)(u)
	if c == nil {
		return nil
	}

	g := &Widget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Widget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Widget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Widget with another Widget, and returns true if they represent the same GObject.
func (recv *Widget) Equals(other *Widget) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Widget) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitrary Object to Widget.
// Exercise care, as this is a potentially dangerous function if the Object is not a Widget.
func CastToWidget(object *gobject.Object) *Widget {
	return WidgetNewFromC(object.ToC())
}

type signalWidgetAccelClosuresChangedDetail struct {
	callback  WidgetSignalAccelClosuresChangedCallback
	handlerID C.gulong
}

var signalWidgetAccelClosuresChangedId int
var signalWidgetAccelClosuresChangedMap = make(map[int]signalWidgetAccelClosuresChangedDetail)
var signalWidgetAccelClosuresChangedLock sync.RWMutex

// WidgetSignalAccelClosuresChangedCallback is a callback function for a 'accel-closures-changed' signal emitted from a Widget.
type WidgetSignalAccelClosuresChangedCallback func()

/*
ConnectAccelClosuresChanged connects the callback to the 'accel-closures-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectAccelClosuresChanged to remove it.
*/
func (recv *Widget) ConnectAccelClosuresChanged(callback WidgetSignalAccelClosuresChangedCallback) int {
	signalWidgetAccelClosuresChangedLock.Lock()
	defer signalWidgetAccelClosuresChangedLock.Unlock()

	signalWidgetAccelClosuresChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_accel_closures_changed(instance, C.gpointer(uintptr(signalWidgetAccelClosuresChangedId)))

	detail := signalWidgetAccelClosuresChangedDetail{callback, handlerID}
	signalWidgetAccelClosuresChangedMap[signalWidgetAccelClosuresChangedId] = detail

	return signalWidgetAccelClosuresChangedId
}

/*
DisconnectAccelClosuresChanged disconnects a callback from the 'accel-closures-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectAccelClosuresChanged.
*/
func (recv *Widget) DisconnectAccelClosuresChanged(connectionID int) {
	signalWidgetAccelClosuresChangedLock.Lock()
	defer signalWidgetAccelClosuresChangedLock.Unlock()

	detail, exists := signalWidgetAccelClosuresChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetAccelClosuresChangedMap, connectionID)
}

//export widget_accelClosuresChangedHandler
func widget_accelClosuresChangedHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetAccelClosuresChangedLock.RLock()
	defer signalWidgetAccelClosuresChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetAccelClosuresChangedMap[index].callback
	callback()
}

type signalWidgetButtonPressEventDetail struct {
	callback  WidgetSignalButtonPressEventCallback
	handlerID C.gulong
}

var signalWidgetButtonPressEventId int
var signalWidgetButtonPressEventMap = make(map[int]signalWidgetButtonPressEventDetail)
var signalWidgetButtonPressEventLock sync.RWMutex

// WidgetSignalButtonPressEventCallback is a callback function for a 'button-press-event' signal emitted from a Widget.
type WidgetSignalButtonPressEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonPressEvent connects the callback to the 'button-press-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectButtonPressEvent to remove it.
*/
func (recv *Widget) ConnectButtonPressEvent(callback WidgetSignalButtonPressEventCallback) int {
	signalWidgetButtonPressEventLock.Lock()
	defer signalWidgetButtonPressEventLock.Unlock()

	signalWidgetButtonPressEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_button_press_event(instance, C.gpointer(uintptr(signalWidgetButtonPressEventId)))

	detail := signalWidgetButtonPressEventDetail{callback, handlerID}
	signalWidgetButtonPressEventMap[signalWidgetButtonPressEventId] = detail

	return signalWidgetButtonPressEventId
}

/*
DisconnectButtonPressEvent disconnects a callback from the 'button-press-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectButtonPressEvent.
*/
func (recv *Widget) DisconnectButtonPressEvent(connectionID int) {
	signalWidgetButtonPressEventLock.Lock()
	defer signalWidgetButtonPressEventLock.Unlock()

	detail, exists := signalWidgetButtonPressEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetButtonPressEventMap, connectionID)
}

//export widget_buttonPressEventHandler
func widget_buttonPressEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalWidgetButtonPressEventLock.RLock()
	defer signalWidgetButtonPressEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetButtonPressEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetButtonReleaseEventDetail struct {
	callback  WidgetSignalButtonReleaseEventCallback
	handlerID C.gulong
}

var signalWidgetButtonReleaseEventId int
var signalWidgetButtonReleaseEventMap = make(map[int]signalWidgetButtonReleaseEventDetail)
var signalWidgetButtonReleaseEventLock sync.RWMutex

// WidgetSignalButtonReleaseEventCallback is a callback function for a 'button-release-event' signal emitted from a Widget.
type WidgetSignalButtonReleaseEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonReleaseEvent connects the callback to the 'button-release-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectButtonReleaseEvent to remove it.
*/
func (recv *Widget) ConnectButtonReleaseEvent(callback WidgetSignalButtonReleaseEventCallback) int {
	signalWidgetButtonReleaseEventLock.Lock()
	defer signalWidgetButtonReleaseEventLock.Unlock()

	signalWidgetButtonReleaseEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_button_release_event(instance, C.gpointer(uintptr(signalWidgetButtonReleaseEventId)))

	detail := signalWidgetButtonReleaseEventDetail{callback, handlerID}
	signalWidgetButtonReleaseEventMap[signalWidgetButtonReleaseEventId] = detail

	return signalWidgetButtonReleaseEventId
}

/*
DisconnectButtonReleaseEvent disconnects a callback from the 'button-release-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectButtonReleaseEvent.
*/
func (recv *Widget) DisconnectButtonReleaseEvent(connectionID int) {
	signalWidgetButtonReleaseEventLock.Lock()
	defer signalWidgetButtonReleaseEventLock.Unlock()

	detail, exists := signalWidgetButtonReleaseEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetButtonReleaseEventMap, connectionID)
}

//export widget_buttonReleaseEventHandler
func widget_buttonReleaseEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalWidgetButtonReleaseEventLock.RLock()
	defer signalWidgetButtonReleaseEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetButtonReleaseEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetCanActivateAccelDetail struct {
	callback  WidgetSignalCanActivateAccelCallback
	handlerID C.gulong
}

var signalWidgetCanActivateAccelId int
var signalWidgetCanActivateAccelMap = make(map[int]signalWidgetCanActivateAccelDetail)
var signalWidgetCanActivateAccelLock sync.RWMutex

// WidgetSignalCanActivateAccelCallback is a callback function for a 'can-activate-accel' signal emitted from a Widget.
type WidgetSignalCanActivateAccelCallback func(signalId uint32) bool

/*
ConnectCanActivateAccel connects the callback to the 'can-activate-accel' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectCanActivateAccel to remove it.
*/
func (recv *Widget) ConnectCanActivateAccel(callback WidgetSignalCanActivateAccelCallback) int {
	signalWidgetCanActivateAccelLock.Lock()
	defer signalWidgetCanActivateAccelLock.Unlock()

	signalWidgetCanActivateAccelId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_can_activate_accel(instance, C.gpointer(uintptr(signalWidgetCanActivateAccelId)))

	detail := signalWidgetCanActivateAccelDetail{callback, handlerID}
	signalWidgetCanActivateAccelMap[signalWidgetCanActivateAccelId] = detail

	return signalWidgetCanActivateAccelId
}

/*
DisconnectCanActivateAccel disconnects a callback from the 'can-activate-accel' signal for the Widget.

The connectionID should be a value returned from a call to ConnectCanActivateAccel.
*/
func (recv *Widget) DisconnectCanActivateAccel(connectionID int) {
	signalWidgetCanActivateAccelLock.Lock()
	defer signalWidgetCanActivateAccelLock.Unlock()

	detail, exists := signalWidgetCanActivateAccelMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetCanActivateAccelMap, connectionID)
}

//export widget_canActivateAccelHandler
func widget_canActivateAccelHandler(_ *C.GObject, c_signal_id C.guint, data C.gpointer) C.gboolean {
	signalWidgetCanActivateAccelLock.RLock()
	defer signalWidgetCanActivateAccelLock.RUnlock()

	signalId := uint32(c_signal_id)

	index := int(uintptr(data))
	callback := signalWidgetCanActivateAccelMap[index].callback
	retGo := callback(signalId)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetChildNotifyDetail struct {
	callback  WidgetSignalChildNotifyCallback
	handlerID C.gulong
}

var signalWidgetChildNotifyId int
var signalWidgetChildNotifyMap = make(map[int]signalWidgetChildNotifyDetail)
var signalWidgetChildNotifyLock sync.RWMutex

// WidgetSignalChildNotifyCallback is a callback function for a 'child-notify' signal emitted from a Widget.
type WidgetSignalChildNotifyCallback func(childProperty *gobject.ParamSpec)

/*
ConnectChildNotify connects the callback to the 'child-notify' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectChildNotify to remove it.
*/
func (recv *Widget) ConnectChildNotify(callback WidgetSignalChildNotifyCallback) int {
	signalWidgetChildNotifyLock.Lock()
	defer signalWidgetChildNotifyLock.Unlock()

	signalWidgetChildNotifyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_child_notify(instance, C.gpointer(uintptr(signalWidgetChildNotifyId)))

	detail := signalWidgetChildNotifyDetail{callback, handlerID}
	signalWidgetChildNotifyMap[signalWidgetChildNotifyId] = detail

	return signalWidgetChildNotifyId
}

/*
DisconnectChildNotify disconnects a callback from the 'child-notify' signal for the Widget.

The connectionID should be a value returned from a call to ConnectChildNotify.
*/
func (recv *Widget) DisconnectChildNotify(connectionID int) {
	signalWidgetChildNotifyLock.Lock()
	defer signalWidgetChildNotifyLock.Unlock()

	detail, exists := signalWidgetChildNotifyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetChildNotifyMap, connectionID)
}

//export widget_childNotifyHandler
func widget_childNotifyHandler(_ *C.GObject, c_child_property *C.GParamSpec, data C.gpointer) {
	signalWidgetChildNotifyLock.RLock()
	defer signalWidgetChildNotifyLock.RUnlock()

	childProperty := gobject.ParamSpecNewFromC(unsafe.Pointer(c_child_property))

	index := int(uintptr(data))
	callback := signalWidgetChildNotifyMap[index].callback
	callback(childProperty)
}

type signalWidgetCompositedChangedDetail struct {
	callback  WidgetSignalCompositedChangedCallback
	handlerID C.gulong
}

var signalWidgetCompositedChangedId int
var signalWidgetCompositedChangedMap = make(map[int]signalWidgetCompositedChangedDetail)
var signalWidgetCompositedChangedLock sync.RWMutex

// WidgetSignalCompositedChangedCallback is a callback function for a 'composited-changed' signal emitted from a Widget.
type WidgetSignalCompositedChangedCallback func()

/*
ConnectCompositedChanged connects the callback to the 'composited-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectCompositedChanged to remove it.
*/
func (recv *Widget) ConnectCompositedChanged(callback WidgetSignalCompositedChangedCallback) int {
	signalWidgetCompositedChangedLock.Lock()
	defer signalWidgetCompositedChangedLock.Unlock()

	signalWidgetCompositedChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_composited_changed(instance, C.gpointer(uintptr(signalWidgetCompositedChangedId)))

	detail := signalWidgetCompositedChangedDetail{callback, handlerID}
	signalWidgetCompositedChangedMap[signalWidgetCompositedChangedId] = detail

	return signalWidgetCompositedChangedId
}

/*
DisconnectCompositedChanged disconnects a callback from the 'composited-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectCompositedChanged.
*/
func (recv *Widget) DisconnectCompositedChanged(connectionID int) {
	signalWidgetCompositedChangedLock.Lock()
	defer signalWidgetCompositedChangedLock.Unlock()

	detail, exists := signalWidgetCompositedChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetCompositedChangedMap, connectionID)
}

//export widget_compositedChangedHandler
func widget_compositedChangedHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetCompositedChangedLock.RLock()
	defer signalWidgetCompositedChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetCompositedChangedMap[index].callback
	callback()
}

type signalWidgetConfigureEventDetail struct {
	callback  WidgetSignalConfigureEventCallback
	handlerID C.gulong
}

var signalWidgetConfigureEventId int
var signalWidgetConfigureEventMap = make(map[int]signalWidgetConfigureEventDetail)
var signalWidgetConfigureEventLock sync.RWMutex

// WidgetSignalConfigureEventCallback is a callback function for a 'configure-event' signal emitted from a Widget.
type WidgetSignalConfigureEventCallback func(event *gdk.EventConfigure) bool

/*
ConnectConfigureEvent connects the callback to the 'configure-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectConfigureEvent to remove it.
*/
func (recv *Widget) ConnectConfigureEvent(callback WidgetSignalConfigureEventCallback) int {
	signalWidgetConfigureEventLock.Lock()
	defer signalWidgetConfigureEventLock.Unlock()

	signalWidgetConfigureEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_configure_event(instance, C.gpointer(uintptr(signalWidgetConfigureEventId)))

	detail := signalWidgetConfigureEventDetail{callback, handlerID}
	signalWidgetConfigureEventMap[signalWidgetConfigureEventId] = detail

	return signalWidgetConfigureEventId
}

/*
DisconnectConfigureEvent disconnects a callback from the 'configure-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectConfigureEvent.
*/
func (recv *Widget) DisconnectConfigureEvent(connectionID int) {
	signalWidgetConfigureEventLock.Lock()
	defer signalWidgetConfigureEventLock.Unlock()

	detail, exists := signalWidgetConfigureEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetConfigureEventMap, connectionID)
}

//export widget_configureEventHandler
func widget_configureEventHandler(_ *C.GObject, c_event *C.GdkEventConfigure, data C.gpointer) C.gboolean {
	signalWidgetConfigureEventLock.RLock()
	defer signalWidgetConfigureEventLock.RUnlock()

	event := gdk.EventConfigureNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetConfigureEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

type signalWidgetDestroyDetail struct {
	callback  WidgetSignalDestroyCallback
	handlerID C.gulong
}

var signalWidgetDestroyId int
var signalWidgetDestroyMap = make(map[int]signalWidgetDestroyDetail)
var signalWidgetDestroyLock sync.RWMutex

// WidgetSignalDestroyCallback is a callback function for a 'destroy' signal emitted from a Widget.
type WidgetSignalDestroyCallback func()

/*
ConnectDestroy connects the callback to the 'destroy' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDestroy to remove it.
*/
func (recv *Widget) ConnectDestroy(callback WidgetSignalDestroyCallback) int {
	signalWidgetDestroyLock.Lock()
	defer signalWidgetDestroyLock.Unlock()

	signalWidgetDestroyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_destroy(instance, C.gpointer(uintptr(signalWidgetDestroyId)))

	detail := signalWidgetDestroyDetail{callback, handlerID}
	signalWidgetDestroyMap[signalWidgetDestroyId] = detail

	return signalWidgetDestroyId
}

/*
DisconnectDestroy disconnects a callback from the 'destroy' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDestroy.
*/
func (recv *Widget) DisconnectDestroy(connectionID int) {
	signalWidgetDestroyLock.Lock()
	defer signalWidgetDestroyLock.Unlock()

	detail, exists := signalWidgetDestroyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDestroyMap, connectionID)
}

//export widget_destroyHandler
func widget_destroyHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetDestroyLock.RLock()
	defer signalWidgetDestroyLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetDestroyMap[index].callback
	callback()
}

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

type signalWidgetDirectionChangedDetail struct {
	callback  WidgetSignalDirectionChangedCallback
	handlerID C.gulong
}

var signalWidgetDirectionChangedId int
var signalWidgetDirectionChangedMap = make(map[int]signalWidgetDirectionChangedDetail)
var signalWidgetDirectionChangedLock sync.RWMutex

// WidgetSignalDirectionChangedCallback is a callback function for a 'direction-changed' signal emitted from a Widget.
type WidgetSignalDirectionChangedCallback func(previousDirection TextDirection)

/*
ConnectDirectionChanged connects the callback to the 'direction-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDirectionChanged to remove it.
*/
func (recv *Widget) ConnectDirectionChanged(callback WidgetSignalDirectionChangedCallback) int {
	signalWidgetDirectionChangedLock.Lock()
	defer signalWidgetDirectionChangedLock.Unlock()

	signalWidgetDirectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_direction_changed(instance, C.gpointer(uintptr(signalWidgetDirectionChangedId)))

	detail := signalWidgetDirectionChangedDetail{callback, handlerID}
	signalWidgetDirectionChangedMap[signalWidgetDirectionChangedId] = detail

	return signalWidgetDirectionChangedId
}

/*
DisconnectDirectionChanged disconnects a callback from the 'direction-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDirectionChanged.
*/
func (recv *Widget) DisconnectDirectionChanged(connectionID int) {
	signalWidgetDirectionChangedLock.Lock()
	defer signalWidgetDirectionChangedLock.Unlock()

	detail, exists := signalWidgetDirectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDirectionChangedMap, connectionID)
}

//export widget_directionChangedHandler
func widget_directionChangedHandler(_ *C.GObject, c_previous_direction C.GtkTextDirection, data C.gpointer) {
	signalWidgetDirectionChangedLock.RLock()
	defer signalWidgetDirectionChangedLock.RUnlock()

	previousDirection := TextDirection(c_previous_direction)

	index := int(uintptr(data))
	callback := signalWidgetDirectionChangedMap[index].callback
	callback(previousDirection)
}

type signalWidgetDragBeginDetail struct {
	callback  WidgetSignalDragBeginCallback
	handlerID C.gulong
}

var signalWidgetDragBeginId int
var signalWidgetDragBeginMap = make(map[int]signalWidgetDragBeginDetail)
var signalWidgetDragBeginLock sync.RWMutex

// WidgetSignalDragBeginCallback is a callback function for a 'drag-begin' signal emitted from a Widget.
type WidgetSignalDragBeginCallback func(context *gdk.DragContext)

/*
ConnectDragBegin connects the callback to the 'drag-begin' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragBegin to remove it.
*/
func (recv *Widget) ConnectDragBegin(callback WidgetSignalDragBeginCallback) int {
	signalWidgetDragBeginLock.Lock()
	defer signalWidgetDragBeginLock.Unlock()

	signalWidgetDragBeginId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_begin(instance, C.gpointer(uintptr(signalWidgetDragBeginId)))

	detail := signalWidgetDragBeginDetail{callback, handlerID}
	signalWidgetDragBeginMap[signalWidgetDragBeginId] = detail

	return signalWidgetDragBeginId
}

/*
DisconnectDragBegin disconnects a callback from the 'drag-begin' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragBegin.
*/
func (recv *Widget) DisconnectDragBegin(connectionID int) {
	signalWidgetDragBeginLock.Lock()
	defer signalWidgetDragBeginLock.Unlock()

	detail, exists := signalWidgetDragBeginMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragBeginMap, connectionID)
}

//export widget_dragBeginHandler
func widget_dragBeginHandler(_ *C.GObject, c_context *C.GdkDragContext, data C.gpointer) {
	signalWidgetDragBeginLock.RLock()
	defer signalWidgetDragBeginLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalWidgetDragBeginMap[index].callback
	callback(context)
}

type signalWidgetDragDataDeleteDetail struct {
	callback  WidgetSignalDragDataDeleteCallback
	handlerID C.gulong
}

var signalWidgetDragDataDeleteId int
var signalWidgetDragDataDeleteMap = make(map[int]signalWidgetDragDataDeleteDetail)
var signalWidgetDragDataDeleteLock sync.RWMutex

// WidgetSignalDragDataDeleteCallback is a callback function for a 'drag-data-delete' signal emitted from a Widget.
type WidgetSignalDragDataDeleteCallback func(context *gdk.DragContext)

/*
ConnectDragDataDelete connects the callback to the 'drag-data-delete' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragDataDelete to remove it.
*/
func (recv *Widget) ConnectDragDataDelete(callback WidgetSignalDragDataDeleteCallback) int {
	signalWidgetDragDataDeleteLock.Lock()
	defer signalWidgetDragDataDeleteLock.Unlock()

	signalWidgetDragDataDeleteId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_data_delete(instance, C.gpointer(uintptr(signalWidgetDragDataDeleteId)))

	detail := signalWidgetDragDataDeleteDetail{callback, handlerID}
	signalWidgetDragDataDeleteMap[signalWidgetDragDataDeleteId] = detail

	return signalWidgetDragDataDeleteId
}

/*
DisconnectDragDataDelete disconnects a callback from the 'drag-data-delete' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragDataDelete.
*/
func (recv *Widget) DisconnectDragDataDelete(connectionID int) {
	signalWidgetDragDataDeleteLock.Lock()
	defer signalWidgetDragDataDeleteLock.Unlock()

	detail, exists := signalWidgetDragDataDeleteMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragDataDeleteMap, connectionID)
}

//export widget_dragDataDeleteHandler
func widget_dragDataDeleteHandler(_ *C.GObject, c_context *C.GdkDragContext, data C.gpointer) {
	signalWidgetDragDataDeleteLock.RLock()
	defer signalWidgetDragDataDeleteLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalWidgetDragDataDeleteMap[index].callback
	callback(context)
}

type signalWidgetDragDataGetDetail struct {
	callback  WidgetSignalDragDataGetCallback
	handlerID C.gulong
}

var signalWidgetDragDataGetId int
var signalWidgetDragDataGetMap = make(map[int]signalWidgetDragDataGetDetail)
var signalWidgetDragDataGetLock sync.RWMutex

// WidgetSignalDragDataGetCallback is a callback function for a 'drag-data-get' signal emitted from a Widget.
type WidgetSignalDragDataGetCallback func(context *gdk.DragContext, Data *SelectionData, info uint32, time uint32)

/*
ConnectDragDataGet connects the callback to the 'drag-data-get' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragDataGet to remove it.
*/
func (recv *Widget) ConnectDragDataGet(callback WidgetSignalDragDataGetCallback) int {
	signalWidgetDragDataGetLock.Lock()
	defer signalWidgetDragDataGetLock.Unlock()

	signalWidgetDragDataGetId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_data_get(instance, C.gpointer(uintptr(signalWidgetDragDataGetId)))

	detail := signalWidgetDragDataGetDetail{callback, handlerID}
	signalWidgetDragDataGetMap[signalWidgetDragDataGetId] = detail

	return signalWidgetDragDataGetId
}

/*
DisconnectDragDataGet disconnects a callback from the 'drag-data-get' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragDataGet.
*/
func (recv *Widget) DisconnectDragDataGet(connectionID int) {
	signalWidgetDragDataGetLock.Lock()
	defer signalWidgetDragDataGetLock.Unlock()

	detail, exists := signalWidgetDragDataGetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragDataGetMap, connectionID)
}

//export widget_dragDataGetHandler
func widget_dragDataGetHandler(_ *C.GObject, c_context *C.GdkDragContext, c__data *C.GtkSelectionData, c_info C.guint, c_time C.guint, data C.gpointer) {
	signalWidgetDragDataGetLock.RLock()
	defer signalWidgetDragDataGetLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	Data := SelectionDataNewFromC(unsafe.Pointer(c__data))

	info := uint32(c_info)

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetDragDataGetMap[index].callback
	callback(context, Data, info, time)
}

type signalWidgetDragDataReceivedDetail struct {
	callback  WidgetSignalDragDataReceivedCallback
	handlerID C.gulong
}

var signalWidgetDragDataReceivedId int
var signalWidgetDragDataReceivedMap = make(map[int]signalWidgetDragDataReceivedDetail)
var signalWidgetDragDataReceivedLock sync.RWMutex

// WidgetSignalDragDataReceivedCallback is a callback function for a 'drag-data-received' signal emitted from a Widget.
type WidgetSignalDragDataReceivedCallback func(context *gdk.DragContext, x int32, y int32, Data *SelectionData, info uint32, time uint32)

/*
ConnectDragDataReceived connects the callback to the 'drag-data-received' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragDataReceived to remove it.
*/
func (recv *Widget) ConnectDragDataReceived(callback WidgetSignalDragDataReceivedCallback) int {
	signalWidgetDragDataReceivedLock.Lock()
	defer signalWidgetDragDataReceivedLock.Unlock()

	signalWidgetDragDataReceivedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_data_received(instance, C.gpointer(uintptr(signalWidgetDragDataReceivedId)))

	detail := signalWidgetDragDataReceivedDetail{callback, handlerID}
	signalWidgetDragDataReceivedMap[signalWidgetDragDataReceivedId] = detail

	return signalWidgetDragDataReceivedId
}

/*
DisconnectDragDataReceived disconnects a callback from the 'drag-data-received' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragDataReceived.
*/
func (recv *Widget) DisconnectDragDataReceived(connectionID int) {
	signalWidgetDragDataReceivedLock.Lock()
	defer signalWidgetDragDataReceivedLock.Unlock()

	detail, exists := signalWidgetDragDataReceivedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragDataReceivedMap, connectionID)
}

//export widget_dragDataReceivedHandler
func widget_dragDataReceivedHandler(_ *C.GObject, c_context *C.GdkDragContext, c_x C.gint, c_y C.gint, c__data *C.GtkSelectionData, c_info C.guint, c_time C.guint, data C.gpointer) {
	signalWidgetDragDataReceivedLock.RLock()
	defer signalWidgetDragDataReceivedLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	x := int32(c_x)

	y := int32(c_y)

	Data := SelectionDataNewFromC(unsafe.Pointer(c__data))

	info := uint32(c_info)

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetDragDataReceivedMap[index].callback
	callback(context, x, y, Data, info, time)
}

type signalWidgetDragDropDetail struct {
	callback  WidgetSignalDragDropCallback
	handlerID C.gulong
}

var signalWidgetDragDropId int
var signalWidgetDragDropMap = make(map[int]signalWidgetDragDropDetail)
var signalWidgetDragDropLock sync.RWMutex

// WidgetSignalDragDropCallback is a callback function for a 'drag-drop' signal emitted from a Widget.
type WidgetSignalDragDropCallback func(context *gdk.DragContext, x int32, y int32, time uint32) bool

/*
ConnectDragDrop connects the callback to the 'drag-drop' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragDrop to remove it.
*/
func (recv *Widget) ConnectDragDrop(callback WidgetSignalDragDropCallback) int {
	signalWidgetDragDropLock.Lock()
	defer signalWidgetDragDropLock.Unlock()

	signalWidgetDragDropId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_drop(instance, C.gpointer(uintptr(signalWidgetDragDropId)))

	detail := signalWidgetDragDropDetail{callback, handlerID}
	signalWidgetDragDropMap[signalWidgetDragDropId] = detail

	return signalWidgetDragDropId
}

/*
DisconnectDragDrop disconnects a callback from the 'drag-drop' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragDrop.
*/
func (recv *Widget) DisconnectDragDrop(connectionID int) {
	signalWidgetDragDropLock.Lock()
	defer signalWidgetDragDropLock.Unlock()

	detail, exists := signalWidgetDragDropMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragDropMap, connectionID)
}

//export widget_dragDropHandler
func widget_dragDropHandler(_ *C.GObject, c_context *C.GdkDragContext, c_x C.gint, c_y C.gint, c_time C.guint, data C.gpointer) C.gboolean {
	signalWidgetDragDropLock.RLock()
	defer signalWidgetDragDropLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	x := int32(c_x)

	y := int32(c_y)

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetDragDropMap[index].callback
	retGo := callback(context, x, y, time)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetDragEndDetail struct {
	callback  WidgetSignalDragEndCallback
	handlerID C.gulong
}

var signalWidgetDragEndId int
var signalWidgetDragEndMap = make(map[int]signalWidgetDragEndDetail)
var signalWidgetDragEndLock sync.RWMutex

// WidgetSignalDragEndCallback is a callback function for a 'drag-end' signal emitted from a Widget.
type WidgetSignalDragEndCallback func(context *gdk.DragContext)

/*
ConnectDragEnd connects the callback to the 'drag-end' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragEnd to remove it.
*/
func (recv *Widget) ConnectDragEnd(callback WidgetSignalDragEndCallback) int {
	signalWidgetDragEndLock.Lock()
	defer signalWidgetDragEndLock.Unlock()

	signalWidgetDragEndId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_end(instance, C.gpointer(uintptr(signalWidgetDragEndId)))

	detail := signalWidgetDragEndDetail{callback, handlerID}
	signalWidgetDragEndMap[signalWidgetDragEndId] = detail

	return signalWidgetDragEndId
}

/*
DisconnectDragEnd disconnects a callback from the 'drag-end' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragEnd.
*/
func (recv *Widget) DisconnectDragEnd(connectionID int) {
	signalWidgetDragEndLock.Lock()
	defer signalWidgetDragEndLock.Unlock()

	detail, exists := signalWidgetDragEndMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragEndMap, connectionID)
}

//export widget_dragEndHandler
func widget_dragEndHandler(_ *C.GObject, c_context *C.GdkDragContext, data C.gpointer) {
	signalWidgetDragEndLock.RLock()
	defer signalWidgetDragEndLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalWidgetDragEndMap[index].callback
	callback(context)
}

type signalWidgetDragLeaveDetail struct {
	callback  WidgetSignalDragLeaveCallback
	handlerID C.gulong
}

var signalWidgetDragLeaveId int
var signalWidgetDragLeaveMap = make(map[int]signalWidgetDragLeaveDetail)
var signalWidgetDragLeaveLock sync.RWMutex

// WidgetSignalDragLeaveCallback is a callback function for a 'drag-leave' signal emitted from a Widget.
type WidgetSignalDragLeaveCallback func(context *gdk.DragContext, time uint32)

/*
ConnectDragLeave connects the callback to the 'drag-leave' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragLeave to remove it.
*/
func (recv *Widget) ConnectDragLeave(callback WidgetSignalDragLeaveCallback) int {
	signalWidgetDragLeaveLock.Lock()
	defer signalWidgetDragLeaveLock.Unlock()

	signalWidgetDragLeaveId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_leave(instance, C.gpointer(uintptr(signalWidgetDragLeaveId)))

	detail := signalWidgetDragLeaveDetail{callback, handlerID}
	signalWidgetDragLeaveMap[signalWidgetDragLeaveId] = detail

	return signalWidgetDragLeaveId
}

/*
DisconnectDragLeave disconnects a callback from the 'drag-leave' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragLeave.
*/
func (recv *Widget) DisconnectDragLeave(connectionID int) {
	signalWidgetDragLeaveLock.Lock()
	defer signalWidgetDragLeaveLock.Unlock()

	detail, exists := signalWidgetDragLeaveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragLeaveMap, connectionID)
}

//export widget_dragLeaveHandler
func widget_dragLeaveHandler(_ *C.GObject, c_context *C.GdkDragContext, c_time C.guint, data C.gpointer) {
	signalWidgetDragLeaveLock.RLock()
	defer signalWidgetDragLeaveLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetDragLeaveMap[index].callback
	callback(context, time)
}

type signalWidgetDragMotionDetail struct {
	callback  WidgetSignalDragMotionCallback
	handlerID C.gulong
}

var signalWidgetDragMotionId int
var signalWidgetDragMotionMap = make(map[int]signalWidgetDragMotionDetail)
var signalWidgetDragMotionLock sync.RWMutex

// WidgetSignalDragMotionCallback is a callback function for a 'drag-motion' signal emitted from a Widget.
type WidgetSignalDragMotionCallback func(context *gdk.DragContext, x int32, y int32, time uint32) bool

/*
ConnectDragMotion connects the callback to the 'drag-motion' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragMotion to remove it.
*/
func (recv *Widget) ConnectDragMotion(callback WidgetSignalDragMotionCallback) int {
	signalWidgetDragMotionLock.Lock()
	defer signalWidgetDragMotionLock.Unlock()

	signalWidgetDragMotionId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_motion(instance, C.gpointer(uintptr(signalWidgetDragMotionId)))

	detail := signalWidgetDragMotionDetail{callback, handlerID}
	signalWidgetDragMotionMap[signalWidgetDragMotionId] = detail

	return signalWidgetDragMotionId
}

/*
DisconnectDragMotion disconnects a callback from the 'drag-motion' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragMotion.
*/
func (recv *Widget) DisconnectDragMotion(connectionID int) {
	signalWidgetDragMotionLock.Lock()
	defer signalWidgetDragMotionLock.Unlock()

	detail, exists := signalWidgetDragMotionMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragMotionMap, connectionID)
}

//export widget_dragMotionHandler
func widget_dragMotionHandler(_ *C.GObject, c_context *C.GdkDragContext, c_x C.gint, c_y C.gint, c_time C.guint, data C.gpointer) C.gboolean {
	signalWidgetDragMotionLock.RLock()
	defer signalWidgetDragMotionLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	x := int32(c_x)

	y := int32(c_y)

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetDragMotionMap[index].callback
	retGo := callback(context, x, y, time)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetEnterNotifyEventDetail struct {
	callback  WidgetSignalEnterNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetEnterNotifyEventId int
var signalWidgetEnterNotifyEventMap = make(map[int]signalWidgetEnterNotifyEventDetail)
var signalWidgetEnterNotifyEventLock sync.RWMutex

// WidgetSignalEnterNotifyEventCallback is a callback function for a 'enter-notify-event' signal emitted from a Widget.
type WidgetSignalEnterNotifyEventCallback func(event *gdk.EventCrossing) bool

/*
ConnectEnterNotifyEvent connects the callback to the 'enter-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectEnterNotifyEvent to remove it.
*/
func (recv *Widget) ConnectEnterNotifyEvent(callback WidgetSignalEnterNotifyEventCallback) int {
	signalWidgetEnterNotifyEventLock.Lock()
	defer signalWidgetEnterNotifyEventLock.Unlock()

	signalWidgetEnterNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_enter_notify_event(instance, C.gpointer(uintptr(signalWidgetEnterNotifyEventId)))

	detail := signalWidgetEnterNotifyEventDetail{callback, handlerID}
	signalWidgetEnterNotifyEventMap[signalWidgetEnterNotifyEventId] = detail

	return signalWidgetEnterNotifyEventId
}

/*
DisconnectEnterNotifyEvent disconnects a callback from the 'enter-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectEnterNotifyEvent.
*/
func (recv *Widget) DisconnectEnterNotifyEvent(connectionID int) {
	signalWidgetEnterNotifyEventLock.Lock()
	defer signalWidgetEnterNotifyEventLock.Unlock()

	detail, exists := signalWidgetEnterNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetEnterNotifyEventMap, connectionID)
}

//export widget_enterNotifyEventHandler
func widget_enterNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventCrossing, data C.gpointer) C.gboolean {
	signalWidgetEnterNotifyEventLock.RLock()
	defer signalWidgetEnterNotifyEventLock.RUnlock()

	event := gdk.EventCrossingNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetEnterNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

type signalWidgetFocusDetail struct {
	callback  WidgetSignalFocusCallback
	handlerID C.gulong
}

var signalWidgetFocusId int
var signalWidgetFocusMap = make(map[int]signalWidgetFocusDetail)
var signalWidgetFocusLock sync.RWMutex

// WidgetSignalFocusCallback is a callback function for a 'focus' signal emitted from a Widget.
type WidgetSignalFocusCallback func(direction DirectionType) bool

/*
ConnectFocus connects the callback to the 'focus' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectFocus to remove it.
*/
func (recv *Widget) ConnectFocus(callback WidgetSignalFocusCallback) int {
	signalWidgetFocusLock.Lock()
	defer signalWidgetFocusLock.Unlock()

	signalWidgetFocusId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_focus(instance, C.gpointer(uintptr(signalWidgetFocusId)))

	detail := signalWidgetFocusDetail{callback, handlerID}
	signalWidgetFocusMap[signalWidgetFocusId] = detail

	return signalWidgetFocusId
}

/*
DisconnectFocus disconnects a callback from the 'focus' signal for the Widget.

The connectionID should be a value returned from a call to ConnectFocus.
*/
func (recv *Widget) DisconnectFocus(connectionID int) {
	signalWidgetFocusLock.Lock()
	defer signalWidgetFocusLock.Unlock()

	detail, exists := signalWidgetFocusMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetFocusMap, connectionID)
}

//export widget_focusHandler
func widget_focusHandler(_ *C.GObject, c_direction C.GtkDirectionType, data C.gpointer) C.gboolean {
	signalWidgetFocusLock.RLock()
	defer signalWidgetFocusLock.RUnlock()

	direction := DirectionType(c_direction)

	index := int(uintptr(data))
	callback := signalWidgetFocusMap[index].callback
	retGo := callback(direction)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetFocusInEventDetail struct {
	callback  WidgetSignalFocusInEventCallback
	handlerID C.gulong
}

var signalWidgetFocusInEventId int
var signalWidgetFocusInEventMap = make(map[int]signalWidgetFocusInEventDetail)
var signalWidgetFocusInEventLock sync.RWMutex

// WidgetSignalFocusInEventCallback is a callback function for a 'focus-in-event' signal emitted from a Widget.
type WidgetSignalFocusInEventCallback func(event *gdk.EventFocus) bool

/*
ConnectFocusInEvent connects the callback to the 'focus-in-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectFocusInEvent to remove it.
*/
func (recv *Widget) ConnectFocusInEvent(callback WidgetSignalFocusInEventCallback) int {
	signalWidgetFocusInEventLock.Lock()
	defer signalWidgetFocusInEventLock.Unlock()

	signalWidgetFocusInEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_focus_in_event(instance, C.gpointer(uintptr(signalWidgetFocusInEventId)))

	detail := signalWidgetFocusInEventDetail{callback, handlerID}
	signalWidgetFocusInEventMap[signalWidgetFocusInEventId] = detail

	return signalWidgetFocusInEventId
}

/*
DisconnectFocusInEvent disconnects a callback from the 'focus-in-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectFocusInEvent.
*/
func (recv *Widget) DisconnectFocusInEvent(connectionID int) {
	signalWidgetFocusInEventLock.Lock()
	defer signalWidgetFocusInEventLock.Unlock()

	detail, exists := signalWidgetFocusInEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetFocusInEventMap, connectionID)
}

//export widget_focusInEventHandler
func widget_focusInEventHandler(_ *C.GObject, c_event *C.GdkEventFocus, data C.gpointer) C.gboolean {
	signalWidgetFocusInEventLock.RLock()
	defer signalWidgetFocusInEventLock.RUnlock()

	event := gdk.EventFocusNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetFocusInEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetFocusOutEventDetail struct {
	callback  WidgetSignalFocusOutEventCallback
	handlerID C.gulong
}

var signalWidgetFocusOutEventId int
var signalWidgetFocusOutEventMap = make(map[int]signalWidgetFocusOutEventDetail)
var signalWidgetFocusOutEventLock sync.RWMutex

// WidgetSignalFocusOutEventCallback is a callback function for a 'focus-out-event' signal emitted from a Widget.
type WidgetSignalFocusOutEventCallback func(event *gdk.EventFocus) bool

/*
ConnectFocusOutEvent connects the callback to the 'focus-out-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectFocusOutEvent to remove it.
*/
func (recv *Widget) ConnectFocusOutEvent(callback WidgetSignalFocusOutEventCallback) int {
	signalWidgetFocusOutEventLock.Lock()
	defer signalWidgetFocusOutEventLock.Unlock()

	signalWidgetFocusOutEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_focus_out_event(instance, C.gpointer(uintptr(signalWidgetFocusOutEventId)))

	detail := signalWidgetFocusOutEventDetail{callback, handlerID}
	signalWidgetFocusOutEventMap[signalWidgetFocusOutEventId] = detail

	return signalWidgetFocusOutEventId
}

/*
DisconnectFocusOutEvent disconnects a callback from the 'focus-out-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectFocusOutEvent.
*/
func (recv *Widget) DisconnectFocusOutEvent(connectionID int) {
	signalWidgetFocusOutEventLock.Lock()
	defer signalWidgetFocusOutEventLock.Unlock()

	detail, exists := signalWidgetFocusOutEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetFocusOutEventMap, connectionID)
}

//export widget_focusOutEventHandler
func widget_focusOutEventHandler(_ *C.GObject, c_event *C.GdkEventFocus, data C.gpointer) C.gboolean {
	signalWidgetFocusOutEventLock.RLock()
	defer signalWidgetFocusOutEventLock.RUnlock()

	event := gdk.EventFocusNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetFocusOutEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetGrabFocusDetail struct {
	callback  WidgetSignalGrabFocusCallback
	handlerID C.gulong
}

var signalWidgetGrabFocusId int
var signalWidgetGrabFocusMap = make(map[int]signalWidgetGrabFocusDetail)
var signalWidgetGrabFocusLock sync.RWMutex

// WidgetSignalGrabFocusCallback is a callback function for a 'grab-focus' signal emitted from a Widget.
type WidgetSignalGrabFocusCallback func()

/*
ConnectGrabFocus connects the callback to the 'grab-focus' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectGrabFocus to remove it.
*/
func (recv *Widget) ConnectGrabFocus(callback WidgetSignalGrabFocusCallback) int {
	signalWidgetGrabFocusLock.Lock()
	defer signalWidgetGrabFocusLock.Unlock()

	signalWidgetGrabFocusId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_grab_focus(instance, C.gpointer(uintptr(signalWidgetGrabFocusId)))

	detail := signalWidgetGrabFocusDetail{callback, handlerID}
	signalWidgetGrabFocusMap[signalWidgetGrabFocusId] = detail

	return signalWidgetGrabFocusId
}

/*
DisconnectGrabFocus disconnects a callback from the 'grab-focus' signal for the Widget.

The connectionID should be a value returned from a call to ConnectGrabFocus.
*/
func (recv *Widget) DisconnectGrabFocus(connectionID int) {
	signalWidgetGrabFocusLock.Lock()
	defer signalWidgetGrabFocusLock.Unlock()

	detail, exists := signalWidgetGrabFocusMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetGrabFocusMap, connectionID)
}

//export widget_grabFocusHandler
func widget_grabFocusHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetGrabFocusLock.RLock()
	defer signalWidgetGrabFocusLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetGrabFocusMap[index].callback
	callback()
}

type signalWidgetGrabNotifyDetail struct {
	callback  WidgetSignalGrabNotifyCallback
	handlerID C.gulong
}

var signalWidgetGrabNotifyId int
var signalWidgetGrabNotifyMap = make(map[int]signalWidgetGrabNotifyDetail)
var signalWidgetGrabNotifyLock sync.RWMutex

// WidgetSignalGrabNotifyCallback is a callback function for a 'grab-notify' signal emitted from a Widget.
type WidgetSignalGrabNotifyCallback func(wasGrabbed bool)

/*
ConnectGrabNotify connects the callback to the 'grab-notify' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectGrabNotify to remove it.
*/
func (recv *Widget) ConnectGrabNotify(callback WidgetSignalGrabNotifyCallback) int {
	signalWidgetGrabNotifyLock.Lock()
	defer signalWidgetGrabNotifyLock.Unlock()

	signalWidgetGrabNotifyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_grab_notify(instance, C.gpointer(uintptr(signalWidgetGrabNotifyId)))

	detail := signalWidgetGrabNotifyDetail{callback, handlerID}
	signalWidgetGrabNotifyMap[signalWidgetGrabNotifyId] = detail

	return signalWidgetGrabNotifyId
}

/*
DisconnectGrabNotify disconnects a callback from the 'grab-notify' signal for the Widget.

The connectionID should be a value returned from a call to ConnectGrabNotify.
*/
func (recv *Widget) DisconnectGrabNotify(connectionID int) {
	signalWidgetGrabNotifyLock.Lock()
	defer signalWidgetGrabNotifyLock.Unlock()

	detail, exists := signalWidgetGrabNotifyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetGrabNotifyMap, connectionID)
}

//export widget_grabNotifyHandler
func widget_grabNotifyHandler(_ *C.GObject, c_was_grabbed C.gboolean, data C.gpointer) {
	signalWidgetGrabNotifyLock.RLock()
	defer signalWidgetGrabNotifyLock.RUnlock()

	wasGrabbed := c_was_grabbed == C.TRUE

	index := int(uintptr(data))
	callback := signalWidgetGrabNotifyMap[index].callback
	callback(wasGrabbed)
}

type signalWidgetHideDetail struct {
	callback  WidgetSignalHideCallback
	handlerID C.gulong
}

var signalWidgetHideId int
var signalWidgetHideMap = make(map[int]signalWidgetHideDetail)
var signalWidgetHideLock sync.RWMutex

// WidgetSignalHideCallback is a callback function for a 'hide' signal emitted from a Widget.
type WidgetSignalHideCallback func()

/*
ConnectHide connects the callback to the 'hide' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectHide to remove it.
*/
func (recv *Widget) ConnectHide(callback WidgetSignalHideCallback) int {
	signalWidgetHideLock.Lock()
	defer signalWidgetHideLock.Unlock()

	signalWidgetHideId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_hide(instance, C.gpointer(uintptr(signalWidgetHideId)))

	detail := signalWidgetHideDetail{callback, handlerID}
	signalWidgetHideMap[signalWidgetHideId] = detail

	return signalWidgetHideId
}

/*
DisconnectHide disconnects a callback from the 'hide' signal for the Widget.

The connectionID should be a value returned from a call to ConnectHide.
*/
func (recv *Widget) DisconnectHide(connectionID int) {
	signalWidgetHideLock.Lock()
	defer signalWidgetHideLock.Unlock()

	detail, exists := signalWidgetHideMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetHideMap, connectionID)
}

//export widget_hideHandler
func widget_hideHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetHideLock.RLock()
	defer signalWidgetHideLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetHideMap[index].callback
	callback()
}

type signalWidgetHierarchyChangedDetail struct {
	callback  WidgetSignalHierarchyChangedCallback
	handlerID C.gulong
}

var signalWidgetHierarchyChangedId int
var signalWidgetHierarchyChangedMap = make(map[int]signalWidgetHierarchyChangedDetail)
var signalWidgetHierarchyChangedLock sync.RWMutex

// WidgetSignalHierarchyChangedCallback is a callback function for a 'hierarchy-changed' signal emitted from a Widget.
type WidgetSignalHierarchyChangedCallback func(previousToplevel *Widget)

/*
ConnectHierarchyChanged connects the callback to the 'hierarchy-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectHierarchyChanged to remove it.
*/
func (recv *Widget) ConnectHierarchyChanged(callback WidgetSignalHierarchyChangedCallback) int {
	signalWidgetHierarchyChangedLock.Lock()
	defer signalWidgetHierarchyChangedLock.Unlock()

	signalWidgetHierarchyChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_hierarchy_changed(instance, C.gpointer(uintptr(signalWidgetHierarchyChangedId)))

	detail := signalWidgetHierarchyChangedDetail{callback, handlerID}
	signalWidgetHierarchyChangedMap[signalWidgetHierarchyChangedId] = detail

	return signalWidgetHierarchyChangedId
}

/*
DisconnectHierarchyChanged disconnects a callback from the 'hierarchy-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectHierarchyChanged.
*/
func (recv *Widget) DisconnectHierarchyChanged(connectionID int) {
	signalWidgetHierarchyChangedLock.Lock()
	defer signalWidgetHierarchyChangedLock.Unlock()

	detail, exists := signalWidgetHierarchyChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetHierarchyChangedMap, connectionID)
}

//export widget_hierarchyChangedHandler
func widget_hierarchyChangedHandler(_ *C.GObject, c_previous_toplevel *C.GtkWidget, data C.gpointer) {
	signalWidgetHierarchyChangedLock.RLock()
	defer signalWidgetHierarchyChangedLock.RUnlock()

	previousToplevel := WidgetNewFromC(unsafe.Pointer(c_previous_toplevel))

	index := int(uintptr(data))
	callback := signalWidgetHierarchyChangedMap[index].callback
	callback(previousToplevel)
}

type signalWidgetKeyPressEventDetail struct {
	callback  WidgetSignalKeyPressEventCallback
	handlerID C.gulong
}

var signalWidgetKeyPressEventId int
var signalWidgetKeyPressEventMap = make(map[int]signalWidgetKeyPressEventDetail)
var signalWidgetKeyPressEventLock sync.RWMutex

// WidgetSignalKeyPressEventCallback is a callback function for a 'key-press-event' signal emitted from a Widget.
type WidgetSignalKeyPressEventCallback func(event *gdk.EventKey) bool

/*
ConnectKeyPressEvent connects the callback to the 'key-press-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectKeyPressEvent to remove it.
*/
func (recv *Widget) ConnectKeyPressEvent(callback WidgetSignalKeyPressEventCallback) int {
	signalWidgetKeyPressEventLock.Lock()
	defer signalWidgetKeyPressEventLock.Unlock()

	signalWidgetKeyPressEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_key_press_event(instance, C.gpointer(uintptr(signalWidgetKeyPressEventId)))

	detail := signalWidgetKeyPressEventDetail{callback, handlerID}
	signalWidgetKeyPressEventMap[signalWidgetKeyPressEventId] = detail

	return signalWidgetKeyPressEventId
}

/*
DisconnectKeyPressEvent disconnects a callback from the 'key-press-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeyPressEvent.
*/
func (recv *Widget) DisconnectKeyPressEvent(connectionID int) {
	signalWidgetKeyPressEventLock.Lock()
	defer signalWidgetKeyPressEventLock.Unlock()

	detail, exists := signalWidgetKeyPressEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetKeyPressEventMap, connectionID)
}

//export widget_keyPressEventHandler
func widget_keyPressEventHandler(_ *C.GObject, c_event *C.GdkEventKey, data C.gpointer) C.gboolean {
	signalWidgetKeyPressEventLock.RLock()
	defer signalWidgetKeyPressEventLock.RUnlock()

	event := gdk.EventKeyNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetKeyPressEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetKeyReleaseEventDetail struct {
	callback  WidgetSignalKeyReleaseEventCallback
	handlerID C.gulong
}

var signalWidgetKeyReleaseEventId int
var signalWidgetKeyReleaseEventMap = make(map[int]signalWidgetKeyReleaseEventDetail)
var signalWidgetKeyReleaseEventLock sync.RWMutex

// WidgetSignalKeyReleaseEventCallback is a callback function for a 'key-release-event' signal emitted from a Widget.
type WidgetSignalKeyReleaseEventCallback func(event *gdk.EventKey) bool

/*
ConnectKeyReleaseEvent connects the callback to the 'key-release-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectKeyReleaseEvent to remove it.
*/
func (recv *Widget) ConnectKeyReleaseEvent(callback WidgetSignalKeyReleaseEventCallback) int {
	signalWidgetKeyReleaseEventLock.Lock()
	defer signalWidgetKeyReleaseEventLock.Unlock()

	signalWidgetKeyReleaseEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_key_release_event(instance, C.gpointer(uintptr(signalWidgetKeyReleaseEventId)))

	detail := signalWidgetKeyReleaseEventDetail{callback, handlerID}
	signalWidgetKeyReleaseEventMap[signalWidgetKeyReleaseEventId] = detail

	return signalWidgetKeyReleaseEventId
}

/*
DisconnectKeyReleaseEvent disconnects a callback from the 'key-release-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeyReleaseEvent.
*/
func (recv *Widget) DisconnectKeyReleaseEvent(connectionID int) {
	signalWidgetKeyReleaseEventLock.Lock()
	defer signalWidgetKeyReleaseEventLock.Unlock()

	detail, exists := signalWidgetKeyReleaseEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetKeyReleaseEventMap, connectionID)
}

//export widget_keyReleaseEventHandler
func widget_keyReleaseEventHandler(_ *C.GObject, c_event *C.GdkEventKey, data C.gpointer) C.gboolean {
	signalWidgetKeyReleaseEventLock.RLock()
	defer signalWidgetKeyReleaseEventLock.RUnlock()

	event := gdk.EventKeyNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetKeyReleaseEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetLeaveNotifyEventDetail struct {
	callback  WidgetSignalLeaveNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetLeaveNotifyEventId int
var signalWidgetLeaveNotifyEventMap = make(map[int]signalWidgetLeaveNotifyEventDetail)
var signalWidgetLeaveNotifyEventLock sync.RWMutex

// WidgetSignalLeaveNotifyEventCallback is a callback function for a 'leave-notify-event' signal emitted from a Widget.
type WidgetSignalLeaveNotifyEventCallback func(event *gdk.EventCrossing) bool

/*
ConnectLeaveNotifyEvent connects the callback to the 'leave-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectLeaveNotifyEvent to remove it.
*/
func (recv *Widget) ConnectLeaveNotifyEvent(callback WidgetSignalLeaveNotifyEventCallback) int {
	signalWidgetLeaveNotifyEventLock.Lock()
	defer signalWidgetLeaveNotifyEventLock.Unlock()

	signalWidgetLeaveNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_leave_notify_event(instance, C.gpointer(uintptr(signalWidgetLeaveNotifyEventId)))

	detail := signalWidgetLeaveNotifyEventDetail{callback, handlerID}
	signalWidgetLeaveNotifyEventMap[signalWidgetLeaveNotifyEventId] = detail

	return signalWidgetLeaveNotifyEventId
}

/*
DisconnectLeaveNotifyEvent disconnects a callback from the 'leave-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectLeaveNotifyEvent.
*/
func (recv *Widget) DisconnectLeaveNotifyEvent(connectionID int) {
	signalWidgetLeaveNotifyEventLock.Lock()
	defer signalWidgetLeaveNotifyEventLock.Unlock()

	detail, exists := signalWidgetLeaveNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetLeaveNotifyEventMap, connectionID)
}

//export widget_leaveNotifyEventHandler
func widget_leaveNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventCrossing, data C.gpointer) C.gboolean {
	signalWidgetLeaveNotifyEventLock.RLock()
	defer signalWidgetLeaveNotifyEventLock.RUnlock()

	event := gdk.EventCrossingNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetLeaveNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetMapDetail struct {
	callback  WidgetSignalMapCallback
	handlerID C.gulong
}

var signalWidgetMapId int
var signalWidgetMapMap = make(map[int]signalWidgetMapDetail)
var signalWidgetMapLock sync.RWMutex

// WidgetSignalMapCallback is a callback function for a 'map' signal emitted from a Widget.
type WidgetSignalMapCallback func()

/*
ConnectMap connects the callback to the 'map' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectMap to remove it.
*/
func (recv *Widget) ConnectMap(callback WidgetSignalMapCallback) int {
	signalWidgetMapLock.Lock()
	defer signalWidgetMapLock.Unlock()

	signalWidgetMapId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_map(instance, C.gpointer(uintptr(signalWidgetMapId)))

	detail := signalWidgetMapDetail{callback, handlerID}
	signalWidgetMapMap[signalWidgetMapId] = detail

	return signalWidgetMapId
}

/*
DisconnectMap disconnects a callback from the 'map' signal for the Widget.

The connectionID should be a value returned from a call to ConnectMap.
*/
func (recv *Widget) DisconnectMap(connectionID int) {
	signalWidgetMapLock.Lock()
	defer signalWidgetMapLock.Unlock()

	detail, exists := signalWidgetMapMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetMapMap, connectionID)
}

//export widget_mapHandler
func widget_mapHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetMapLock.RLock()
	defer signalWidgetMapLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetMapMap[index].callback
	callback()
}

type signalWidgetMapEventDetail struct {
	callback  WidgetSignalMapEventCallback
	handlerID C.gulong
}

var signalWidgetMapEventId int
var signalWidgetMapEventMap = make(map[int]signalWidgetMapEventDetail)
var signalWidgetMapEventLock sync.RWMutex

// WidgetSignalMapEventCallback is a callback function for a 'map-event' signal emitted from a Widget.
type WidgetSignalMapEventCallback func(event *gdk.EventAny) bool

/*
ConnectMapEvent connects the callback to the 'map-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectMapEvent to remove it.
*/
func (recv *Widget) ConnectMapEvent(callback WidgetSignalMapEventCallback) int {
	signalWidgetMapEventLock.Lock()
	defer signalWidgetMapEventLock.Unlock()

	signalWidgetMapEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_map_event(instance, C.gpointer(uintptr(signalWidgetMapEventId)))

	detail := signalWidgetMapEventDetail{callback, handlerID}
	signalWidgetMapEventMap[signalWidgetMapEventId] = detail

	return signalWidgetMapEventId
}

/*
DisconnectMapEvent disconnects a callback from the 'map-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectMapEvent.
*/
func (recv *Widget) DisconnectMapEvent(connectionID int) {
	signalWidgetMapEventLock.Lock()
	defer signalWidgetMapEventLock.Unlock()

	detail, exists := signalWidgetMapEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetMapEventMap, connectionID)
}

//export widget_mapEventHandler
func widget_mapEventHandler(_ *C.GObject, c_event *C.GdkEventAny, data C.gpointer) C.gboolean {
	signalWidgetMapEventLock.RLock()
	defer signalWidgetMapEventLock.RUnlock()

	event := gdk.EventAnyNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetMapEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetMnemonicActivateDetail struct {
	callback  WidgetSignalMnemonicActivateCallback
	handlerID C.gulong
}

var signalWidgetMnemonicActivateId int
var signalWidgetMnemonicActivateMap = make(map[int]signalWidgetMnemonicActivateDetail)
var signalWidgetMnemonicActivateLock sync.RWMutex

// WidgetSignalMnemonicActivateCallback is a callback function for a 'mnemonic-activate' signal emitted from a Widget.
type WidgetSignalMnemonicActivateCallback func(groupCycling bool) bool

/*
ConnectMnemonicActivate connects the callback to the 'mnemonic-activate' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectMnemonicActivate to remove it.
*/
func (recv *Widget) ConnectMnemonicActivate(callback WidgetSignalMnemonicActivateCallback) int {
	signalWidgetMnemonicActivateLock.Lock()
	defer signalWidgetMnemonicActivateLock.Unlock()

	signalWidgetMnemonicActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_mnemonic_activate(instance, C.gpointer(uintptr(signalWidgetMnemonicActivateId)))

	detail := signalWidgetMnemonicActivateDetail{callback, handlerID}
	signalWidgetMnemonicActivateMap[signalWidgetMnemonicActivateId] = detail

	return signalWidgetMnemonicActivateId
}

/*
DisconnectMnemonicActivate disconnects a callback from the 'mnemonic-activate' signal for the Widget.

The connectionID should be a value returned from a call to ConnectMnemonicActivate.
*/
func (recv *Widget) DisconnectMnemonicActivate(connectionID int) {
	signalWidgetMnemonicActivateLock.Lock()
	defer signalWidgetMnemonicActivateLock.Unlock()

	detail, exists := signalWidgetMnemonicActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetMnemonicActivateMap, connectionID)
}

//export widget_mnemonicActivateHandler
func widget_mnemonicActivateHandler(_ *C.GObject, c_group_cycling C.gboolean, data C.gpointer) C.gboolean {
	signalWidgetMnemonicActivateLock.RLock()
	defer signalWidgetMnemonicActivateLock.RUnlock()

	groupCycling := c_group_cycling == C.TRUE

	index := int(uintptr(data))
	callback := signalWidgetMnemonicActivateMap[index].callback
	retGo := callback(groupCycling)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetMotionNotifyEventDetail struct {
	callback  WidgetSignalMotionNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetMotionNotifyEventId int
var signalWidgetMotionNotifyEventMap = make(map[int]signalWidgetMotionNotifyEventDetail)
var signalWidgetMotionNotifyEventLock sync.RWMutex

// WidgetSignalMotionNotifyEventCallback is a callback function for a 'motion-notify-event' signal emitted from a Widget.
type WidgetSignalMotionNotifyEventCallback func(event *gdk.EventMotion) bool

/*
ConnectMotionNotifyEvent connects the callback to the 'motion-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectMotionNotifyEvent to remove it.
*/
func (recv *Widget) ConnectMotionNotifyEvent(callback WidgetSignalMotionNotifyEventCallback) int {
	signalWidgetMotionNotifyEventLock.Lock()
	defer signalWidgetMotionNotifyEventLock.Unlock()

	signalWidgetMotionNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_motion_notify_event(instance, C.gpointer(uintptr(signalWidgetMotionNotifyEventId)))

	detail := signalWidgetMotionNotifyEventDetail{callback, handlerID}
	signalWidgetMotionNotifyEventMap[signalWidgetMotionNotifyEventId] = detail

	return signalWidgetMotionNotifyEventId
}

/*
DisconnectMotionNotifyEvent disconnects a callback from the 'motion-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectMotionNotifyEvent.
*/
func (recv *Widget) DisconnectMotionNotifyEvent(connectionID int) {
	signalWidgetMotionNotifyEventLock.Lock()
	defer signalWidgetMotionNotifyEventLock.Unlock()

	detail, exists := signalWidgetMotionNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetMotionNotifyEventMap, connectionID)
}

//export widget_motionNotifyEventHandler
func widget_motionNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventMotion, data C.gpointer) C.gboolean {
	signalWidgetMotionNotifyEventLock.RLock()
	defer signalWidgetMotionNotifyEventLock.RUnlock()

	event := gdk.EventMotionNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetMotionNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetMoveFocusDetail struct {
	callback  WidgetSignalMoveFocusCallback
	handlerID C.gulong
}

var signalWidgetMoveFocusId int
var signalWidgetMoveFocusMap = make(map[int]signalWidgetMoveFocusDetail)
var signalWidgetMoveFocusLock sync.RWMutex

// WidgetSignalMoveFocusCallback is a callback function for a 'move-focus' signal emitted from a Widget.
type WidgetSignalMoveFocusCallback func(direction DirectionType)

/*
ConnectMoveFocus connects the callback to the 'move-focus' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectMoveFocus to remove it.
*/
func (recv *Widget) ConnectMoveFocus(callback WidgetSignalMoveFocusCallback) int {
	signalWidgetMoveFocusLock.Lock()
	defer signalWidgetMoveFocusLock.Unlock()

	signalWidgetMoveFocusId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_move_focus(instance, C.gpointer(uintptr(signalWidgetMoveFocusId)))

	detail := signalWidgetMoveFocusDetail{callback, handlerID}
	signalWidgetMoveFocusMap[signalWidgetMoveFocusId] = detail

	return signalWidgetMoveFocusId
}

/*
DisconnectMoveFocus disconnects a callback from the 'move-focus' signal for the Widget.

The connectionID should be a value returned from a call to ConnectMoveFocus.
*/
func (recv *Widget) DisconnectMoveFocus(connectionID int) {
	signalWidgetMoveFocusLock.Lock()
	defer signalWidgetMoveFocusLock.Unlock()

	detail, exists := signalWidgetMoveFocusMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetMoveFocusMap, connectionID)
}

//export widget_moveFocusHandler
func widget_moveFocusHandler(_ *C.GObject, c_direction C.GtkDirectionType, data C.gpointer) {
	signalWidgetMoveFocusLock.RLock()
	defer signalWidgetMoveFocusLock.RUnlock()

	direction := DirectionType(c_direction)

	index := int(uintptr(data))
	callback := signalWidgetMoveFocusMap[index].callback
	callback(direction)
}

type signalWidgetParentSetDetail struct {
	callback  WidgetSignalParentSetCallback
	handlerID C.gulong
}

var signalWidgetParentSetId int
var signalWidgetParentSetMap = make(map[int]signalWidgetParentSetDetail)
var signalWidgetParentSetLock sync.RWMutex

// WidgetSignalParentSetCallback is a callback function for a 'parent-set' signal emitted from a Widget.
type WidgetSignalParentSetCallback func(oldParent *Widget)

/*
ConnectParentSet connects the callback to the 'parent-set' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectParentSet to remove it.
*/
func (recv *Widget) ConnectParentSet(callback WidgetSignalParentSetCallback) int {
	signalWidgetParentSetLock.Lock()
	defer signalWidgetParentSetLock.Unlock()

	signalWidgetParentSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_parent_set(instance, C.gpointer(uintptr(signalWidgetParentSetId)))

	detail := signalWidgetParentSetDetail{callback, handlerID}
	signalWidgetParentSetMap[signalWidgetParentSetId] = detail

	return signalWidgetParentSetId
}

/*
DisconnectParentSet disconnects a callback from the 'parent-set' signal for the Widget.

The connectionID should be a value returned from a call to ConnectParentSet.
*/
func (recv *Widget) DisconnectParentSet(connectionID int) {
	signalWidgetParentSetLock.Lock()
	defer signalWidgetParentSetLock.Unlock()

	detail, exists := signalWidgetParentSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetParentSetMap, connectionID)
}

//export widget_parentSetHandler
func widget_parentSetHandler(_ *C.GObject, c_old_parent *C.GtkWidget, data C.gpointer) {
	signalWidgetParentSetLock.RLock()
	defer signalWidgetParentSetLock.RUnlock()

	oldParent := WidgetNewFromC(unsafe.Pointer(c_old_parent))

	index := int(uintptr(data))
	callback := signalWidgetParentSetMap[index].callback
	callback(oldParent)
}

type signalWidgetPopupMenuDetail struct {
	callback  WidgetSignalPopupMenuCallback
	handlerID C.gulong
}

var signalWidgetPopupMenuId int
var signalWidgetPopupMenuMap = make(map[int]signalWidgetPopupMenuDetail)
var signalWidgetPopupMenuLock sync.RWMutex

// WidgetSignalPopupMenuCallback is a callback function for a 'popup-menu' signal emitted from a Widget.
type WidgetSignalPopupMenuCallback func() bool

/*
ConnectPopupMenu connects the callback to the 'popup-menu' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectPopupMenu to remove it.
*/
func (recv *Widget) ConnectPopupMenu(callback WidgetSignalPopupMenuCallback) int {
	signalWidgetPopupMenuLock.Lock()
	defer signalWidgetPopupMenuLock.Unlock()

	signalWidgetPopupMenuId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_popup_menu(instance, C.gpointer(uintptr(signalWidgetPopupMenuId)))

	detail := signalWidgetPopupMenuDetail{callback, handlerID}
	signalWidgetPopupMenuMap[signalWidgetPopupMenuId] = detail

	return signalWidgetPopupMenuId
}

/*
DisconnectPopupMenu disconnects a callback from the 'popup-menu' signal for the Widget.

The connectionID should be a value returned from a call to ConnectPopupMenu.
*/
func (recv *Widget) DisconnectPopupMenu(connectionID int) {
	signalWidgetPopupMenuLock.Lock()
	defer signalWidgetPopupMenuLock.Unlock()

	detail, exists := signalWidgetPopupMenuMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetPopupMenuMap, connectionID)
}

//export widget_popupMenuHandler
func widget_popupMenuHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	signalWidgetPopupMenuLock.RLock()
	defer signalWidgetPopupMenuLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetPopupMenuMap[index].callback
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetPropertyNotifyEventDetail struct {
	callback  WidgetSignalPropertyNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetPropertyNotifyEventId int
var signalWidgetPropertyNotifyEventMap = make(map[int]signalWidgetPropertyNotifyEventDetail)
var signalWidgetPropertyNotifyEventLock sync.RWMutex

// WidgetSignalPropertyNotifyEventCallback is a callback function for a 'property-notify-event' signal emitted from a Widget.
type WidgetSignalPropertyNotifyEventCallback func(event *gdk.EventProperty) bool

/*
ConnectPropertyNotifyEvent connects the callback to the 'property-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectPropertyNotifyEvent to remove it.
*/
func (recv *Widget) ConnectPropertyNotifyEvent(callback WidgetSignalPropertyNotifyEventCallback) int {
	signalWidgetPropertyNotifyEventLock.Lock()
	defer signalWidgetPropertyNotifyEventLock.Unlock()

	signalWidgetPropertyNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_property_notify_event(instance, C.gpointer(uintptr(signalWidgetPropertyNotifyEventId)))

	detail := signalWidgetPropertyNotifyEventDetail{callback, handlerID}
	signalWidgetPropertyNotifyEventMap[signalWidgetPropertyNotifyEventId] = detail

	return signalWidgetPropertyNotifyEventId
}

/*
DisconnectPropertyNotifyEvent disconnects a callback from the 'property-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectPropertyNotifyEvent.
*/
func (recv *Widget) DisconnectPropertyNotifyEvent(connectionID int) {
	signalWidgetPropertyNotifyEventLock.Lock()
	defer signalWidgetPropertyNotifyEventLock.Unlock()

	detail, exists := signalWidgetPropertyNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetPropertyNotifyEventMap, connectionID)
}

//export widget_propertyNotifyEventHandler
func widget_propertyNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventProperty, data C.gpointer) C.gboolean {
	signalWidgetPropertyNotifyEventLock.RLock()
	defer signalWidgetPropertyNotifyEventLock.RUnlock()

	event := gdk.EventPropertyNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetPropertyNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetProximityInEventDetail struct {
	callback  WidgetSignalProximityInEventCallback
	handlerID C.gulong
}

var signalWidgetProximityInEventId int
var signalWidgetProximityInEventMap = make(map[int]signalWidgetProximityInEventDetail)
var signalWidgetProximityInEventLock sync.RWMutex

// WidgetSignalProximityInEventCallback is a callback function for a 'proximity-in-event' signal emitted from a Widget.
type WidgetSignalProximityInEventCallback func(event *gdk.EventProximity) bool

/*
ConnectProximityInEvent connects the callback to the 'proximity-in-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectProximityInEvent to remove it.
*/
func (recv *Widget) ConnectProximityInEvent(callback WidgetSignalProximityInEventCallback) int {
	signalWidgetProximityInEventLock.Lock()
	defer signalWidgetProximityInEventLock.Unlock()

	signalWidgetProximityInEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_proximity_in_event(instance, C.gpointer(uintptr(signalWidgetProximityInEventId)))

	detail := signalWidgetProximityInEventDetail{callback, handlerID}
	signalWidgetProximityInEventMap[signalWidgetProximityInEventId] = detail

	return signalWidgetProximityInEventId
}

/*
DisconnectProximityInEvent disconnects a callback from the 'proximity-in-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectProximityInEvent.
*/
func (recv *Widget) DisconnectProximityInEvent(connectionID int) {
	signalWidgetProximityInEventLock.Lock()
	defer signalWidgetProximityInEventLock.Unlock()

	detail, exists := signalWidgetProximityInEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetProximityInEventMap, connectionID)
}

//export widget_proximityInEventHandler
func widget_proximityInEventHandler(_ *C.GObject, c_event *C.GdkEventProximity, data C.gpointer) C.gboolean {
	signalWidgetProximityInEventLock.RLock()
	defer signalWidgetProximityInEventLock.RUnlock()

	event := gdk.EventProximityNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetProximityInEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetProximityOutEventDetail struct {
	callback  WidgetSignalProximityOutEventCallback
	handlerID C.gulong
}

var signalWidgetProximityOutEventId int
var signalWidgetProximityOutEventMap = make(map[int]signalWidgetProximityOutEventDetail)
var signalWidgetProximityOutEventLock sync.RWMutex

// WidgetSignalProximityOutEventCallback is a callback function for a 'proximity-out-event' signal emitted from a Widget.
type WidgetSignalProximityOutEventCallback func(event *gdk.EventProximity) bool

/*
ConnectProximityOutEvent connects the callback to the 'proximity-out-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectProximityOutEvent to remove it.
*/
func (recv *Widget) ConnectProximityOutEvent(callback WidgetSignalProximityOutEventCallback) int {
	signalWidgetProximityOutEventLock.Lock()
	defer signalWidgetProximityOutEventLock.Unlock()

	signalWidgetProximityOutEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_proximity_out_event(instance, C.gpointer(uintptr(signalWidgetProximityOutEventId)))

	detail := signalWidgetProximityOutEventDetail{callback, handlerID}
	signalWidgetProximityOutEventMap[signalWidgetProximityOutEventId] = detail

	return signalWidgetProximityOutEventId
}

/*
DisconnectProximityOutEvent disconnects a callback from the 'proximity-out-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectProximityOutEvent.
*/
func (recv *Widget) DisconnectProximityOutEvent(connectionID int) {
	signalWidgetProximityOutEventLock.Lock()
	defer signalWidgetProximityOutEventLock.Unlock()

	detail, exists := signalWidgetProximityOutEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetProximityOutEventMap, connectionID)
}

//export widget_proximityOutEventHandler
func widget_proximityOutEventHandler(_ *C.GObject, c_event *C.GdkEventProximity, data C.gpointer) C.gboolean {
	signalWidgetProximityOutEventLock.RLock()
	defer signalWidgetProximityOutEventLock.RUnlock()

	event := gdk.EventProximityNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetProximityOutEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetRealizeDetail struct {
	callback  WidgetSignalRealizeCallback
	handlerID C.gulong
}

var signalWidgetRealizeId int
var signalWidgetRealizeMap = make(map[int]signalWidgetRealizeDetail)
var signalWidgetRealizeLock sync.RWMutex

// WidgetSignalRealizeCallback is a callback function for a 'realize' signal emitted from a Widget.
type WidgetSignalRealizeCallback func()

/*
ConnectRealize connects the callback to the 'realize' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectRealize to remove it.
*/
func (recv *Widget) ConnectRealize(callback WidgetSignalRealizeCallback) int {
	signalWidgetRealizeLock.Lock()
	defer signalWidgetRealizeLock.Unlock()

	signalWidgetRealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_realize(instance, C.gpointer(uintptr(signalWidgetRealizeId)))

	detail := signalWidgetRealizeDetail{callback, handlerID}
	signalWidgetRealizeMap[signalWidgetRealizeId] = detail

	return signalWidgetRealizeId
}

/*
DisconnectRealize disconnects a callback from the 'realize' signal for the Widget.

The connectionID should be a value returned from a call to ConnectRealize.
*/
func (recv *Widget) DisconnectRealize(connectionID int) {
	signalWidgetRealizeLock.Lock()
	defer signalWidgetRealizeLock.Unlock()

	detail, exists := signalWidgetRealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetRealizeMap, connectionID)
}

//export widget_realizeHandler
func widget_realizeHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetRealizeLock.RLock()
	defer signalWidgetRealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetRealizeMap[index].callback
	callback()
}

type signalWidgetScreenChangedDetail struct {
	callback  WidgetSignalScreenChangedCallback
	handlerID C.gulong
}

var signalWidgetScreenChangedId int
var signalWidgetScreenChangedMap = make(map[int]signalWidgetScreenChangedDetail)
var signalWidgetScreenChangedLock sync.RWMutex

// WidgetSignalScreenChangedCallback is a callback function for a 'screen-changed' signal emitted from a Widget.
type WidgetSignalScreenChangedCallback func(previousScreen *gdk.Screen)

/*
ConnectScreenChanged connects the callback to the 'screen-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectScreenChanged to remove it.
*/
func (recv *Widget) ConnectScreenChanged(callback WidgetSignalScreenChangedCallback) int {
	signalWidgetScreenChangedLock.Lock()
	defer signalWidgetScreenChangedLock.Unlock()

	signalWidgetScreenChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_screen_changed(instance, C.gpointer(uintptr(signalWidgetScreenChangedId)))

	detail := signalWidgetScreenChangedDetail{callback, handlerID}
	signalWidgetScreenChangedMap[signalWidgetScreenChangedId] = detail

	return signalWidgetScreenChangedId
}

/*
DisconnectScreenChanged disconnects a callback from the 'screen-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectScreenChanged.
*/
func (recv *Widget) DisconnectScreenChanged(connectionID int) {
	signalWidgetScreenChangedLock.Lock()
	defer signalWidgetScreenChangedLock.Unlock()

	detail, exists := signalWidgetScreenChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetScreenChangedMap, connectionID)
}

//export widget_screenChangedHandler
func widget_screenChangedHandler(_ *C.GObject, c_previous_screen *C.GdkScreen, data C.gpointer) {
	signalWidgetScreenChangedLock.RLock()
	defer signalWidgetScreenChangedLock.RUnlock()

	previousScreen := gdk.ScreenNewFromC(unsafe.Pointer(c_previous_screen))

	index := int(uintptr(data))
	callback := signalWidgetScreenChangedMap[index].callback
	callback(previousScreen)
}

type signalWidgetScrollEventDetail struct {
	callback  WidgetSignalScrollEventCallback
	handlerID C.gulong
}

var signalWidgetScrollEventId int
var signalWidgetScrollEventMap = make(map[int]signalWidgetScrollEventDetail)
var signalWidgetScrollEventLock sync.RWMutex

// WidgetSignalScrollEventCallback is a callback function for a 'scroll-event' signal emitted from a Widget.
type WidgetSignalScrollEventCallback func(event *gdk.EventScroll) bool

/*
ConnectScrollEvent connects the callback to the 'scroll-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectScrollEvent to remove it.
*/
func (recv *Widget) ConnectScrollEvent(callback WidgetSignalScrollEventCallback) int {
	signalWidgetScrollEventLock.Lock()
	defer signalWidgetScrollEventLock.Unlock()

	signalWidgetScrollEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_scroll_event(instance, C.gpointer(uintptr(signalWidgetScrollEventId)))

	detail := signalWidgetScrollEventDetail{callback, handlerID}
	signalWidgetScrollEventMap[signalWidgetScrollEventId] = detail

	return signalWidgetScrollEventId
}

/*
DisconnectScrollEvent disconnects a callback from the 'scroll-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectScrollEvent.
*/
func (recv *Widget) DisconnectScrollEvent(connectionID int) {
	signalWidgetScrollEventLock.Lock()
	defer signalWidgetScrollEventLock.Unlock()

	detail, exists := signalWidgetScrollEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetScrollEventMap, connectionID)
}

//export widget_scrollEventHandler
func widget_scrollEventHandler(_ *C.GObject, c_event *C.GdkEventScroll, data C.gpointer) C.gboolean {
	signalWidgetScrollEventLock.RLock()
	defer signalWidgetScrollEventLock.RUnlock()

	event := gdk.EventScrollNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetScrollEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetSelectionClearEventDetail struct {
	callback  WidgetSignalSelectionClearEventCallback
	handlerID C.gulong
}

var signalWidgetSelectionClearEventId int
var signalWidgetSelectionClearEventMap = make(map[int]signalWidgetSelectionClearEventDetail)
var signalWidgetSelectionClearEventLock sync.RWMutex

// WidgetSignalSelectionClearEventCallback is a callback function for a 'selection-clear-event' signal emitted from a Widget.
type WidgetSignalSelectionClearEventCallback func(event *gdk.EventSelection) bool

/*
ConnectSelectionClearEvent connects the callback to the 'selection-clear-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSelectionClearEvent to remove it.
*/
func (recv *Widget) ConnectSelectionClearEvent(callback WidgetSignalSelectionClearEventCallback) int {
	signalWidgetSelectionClearEventLock.Lock()
	defer signalWidgetSelectionClearEventLock.Unlock()

	signalWidgetSelectionClearEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_selection_clear_event(instance, C.gpointer(uintptr(signalWidgetSelectionClearEventId)))

	detail := signalWidgetSelectionClearEventDetail{callback, handlerID}
	signalWidgetSelectionClearEventMap[signalWidgetSelectionClearEventId] = detail

	return signalWidgetSelectionClearEventId
}

/*
DisconnectSelectionClearEvent disconnects a callback from the 'selection-clear-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSelectionClearEvent.
*/
func (recv *Widget) DisconnectSelectionClearEvent(connectionID int) {
	signalWidgetSelectionClearEventLock.Lock()
	defer signalWidgetSelectionClearEventLock.Unlock()

	detail, exists := signalWidgetSelectionClearEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSelectionClearEventMap, connectionID)
}

//export widget_selectionClearEventHandler
func widget_selectionClearEventHandler(_ *C.GObject, c_event *C.GdkEventSelection, data C.gpointer) C.gboolean {
	signalWidgetSelectionClearEventLock.RLock()
	defer signalWidgetSelectionClearEventLock.RUnlock()

	event := gdk.EventSelectionNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetSelectionClearEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetSelectionGetDetail struct {
	callback  WidgetSignalSelectionGetCallback
	handlerID C.gulong
}

var signalWidgetSelectionGetId int
var signalWidgetSelectionGetMap = make(map[int]signalWidgetSelectionGetDetail)
var signalWidgetSelectionGetLock sync.RWMutex

// WidgetSignalSelectionGetCallback is a callback function for a 'selection-get' signal emitted from a Widget.
type WidgetSignalSelectionGetCallback func(Data *SelectionData, info uint32, time uint32)

/*
ConnectSelectionGet connects the callback to the 'selection-get' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSelectionGet to remove it.
*/
func (recv *Widget) ConnectSelectionGet(callback WidgetSignalSelectionGetCallback) int {
	signalWidgetSelectionGetLock.Lock()
	defer signalWidgetSelectionGetLock.Unlock()

	signalWidgetSelectionGetId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_selection_get(instance, C.gpointer(uintptr(signalWidgetSelectionGetId)))

	detail := signalWidgetSelectionGetDetail{callback, handlerID}
	signalWidgetSelectionGetMap[signalWidgetSelectionGetId] = detail

	return signalWidgetSelectionGetId
}

/*
DisconnectSelectionGet disconnects a callback from the 'selection-get' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSelectionGet.
*/
func (recv *Widget) DisconnectSelectionGet(connectionID int) {
	signalWidgetSelectionGetLock.Lock()
	defer signalWidgetSelectionGetLock.Unlock()

	detail, exists := signalWidgetSelectionGetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSelectionGetMap, connectionID)
}

//export widget_selectionGetHandler
func widget_selectionGetHandler(_ *C.GObject, c__data *C.GtkSelectionData, c_info C.guint, c_time C.guint, data C.gpointer) {
	signalWidgetSelectionGetLock.RLock()
	defer signalWidgetSelectionGetLock.RUnlock()

	Data := SelectionDataNewFromC(unsafe.Pointer(c__data))

	info := uint32(c_info)

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetSelectionGetMap[index].callback
	callback(Data, info, time)
}

type signalWidgetSelectionNotifyEventDetail struct {
	callback  WidgetSignalSelectionNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetSelectionNotifyEventId int
var signalWidgetSelectionNotifyEventMap = make(map[int]signalWidgetSelectionNotifyEventDetail)
var signalWidgetSelectionNotifyEventLock sync.RWMutex

// WidgetSignalSelectionNotifyEventCallback is a callback function for a 'selection-notify-event' signal emitted from a Widget.
type WidgetSignalSelectionNotifyEventCallback func(event *gdk.EventSelection) bool

/*
ConnectSelectionNotifyEvent connects the callback to the 'selection-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSelectionNotifyEvent to remove it.
*/
func (recv *Widget) ConnectSelectionNotifyEvent(callback WidgetSignalSelectionNotifyEventCallback) int {
	signalWidgetSelectionNotifyEventLock.Lock()
	defer signalWidgetSelectionNotifyEventLock.Unlock()

	signalWidgetSelectionNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_selection_notify_event(instance, C.gpointer(uintptr(signalWidgetSelectionNotifyEventId)))

	detail := signalWidgetSelectionNotifyEventDetail{callback, handlerID}
	signalWidgetSelectionNotifyEventMap[signalWidgetSelectionNotifyEventId] = detail

	return signalWidgetSelectionNotifyEventId
}

/*
DisconnectSelectionNotifyEvent disconnects a callback from the 'selection-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSelectionNotifyEvent.
*/
func (recv *Widget) DisconnectSelectionNotifyEvent(connectionID int) {
	signalWidgetSelectionNotifyEventLock.Lock()
	defer signalWidgetSelectionNotifyEventLock.Unlock()

	detail, exists := signalWidgetSelectionNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSelectionNotifyEventMap, connectionID)
}

//export widget_selectionNotifyEventHandler
func widget_selectionNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventSelection, data C.gpointer) C.gboolean {
	signalWidgetSelectionNotifyEventLock.RLock()
	defer signalWidgetSelectionNotifyEventLock.RUnlock()

	event := gdk.EventSelectionNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetSelectionNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetSelectionReceivedDetail struct {
	callback  WidgetSignalSelectionReceivedCallback
	handlerID C.gulong
}

var signalWidgetSelectionReceivedId int
var signalWidgetSelectionReceivedMap = make(map[int]signalWidgetSelectionReceivedDetail)
var signalWidgetSelectionReceivedLock sync.RWMutex

// WidgetSignalSelectionReceivedCallback is a callback function for a 'selection-received' signal emitted from a Widget.
type WidgetSignalSelectionReceivedCallback func(Data *SelectionData, time uint32)

/*
ConnectSelectionReceived connects the callback to the 'selection-received' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSelectionReceived to remove it.
*/
func (recv *Widget) ConnectSelectionReceived(callback WidgetSignalSelectionReceivedCallback) int {
	signalWidgetSelectionReceivedLock.Lock()
	defer signalWidgetSelectionReceivedLock.Unlock()

	signalWidgetSelectionReceivedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_selection_received(instance, C.gpointer(uintptr(signalWidgetSelectionReceivedId)))

	detail := signalWidgetSelectionReceivedDetail{callback, handlerID}
	signalWidgetSelectionReceivedMap[signalWidgetSelectionReceivedId] = detail

	return signalWidgetSelectionReceivedId
}

/*
DisconnectSelectionReceived disconnects a callback from the 'selection-received' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSelectionReceived.
*/
func (recv *Widget) DisconnectSelectionReceived(connectionID int) {
	signalWidgetSelectionReceivedLock.Lock()
	defer signalWidgetSelectionReceivedLock.Unlock()

	detail, exists := signalWidgetSelectionReceivedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSelectionReceivedMap, connectionID)
}

//export widget_selectionReceivedHandler
func widget_selectionReceivedHandler(_ *C.GObject, c__data *C.GtkSelectionData, c_time C.guint, data C.gpointer) {
	signalWidgetSelectionReceivedLock.RLock()
	defer signalWidgetSelectionReceivedLock.RUnlock()

	Data := SelectionDataNewFromC(unsafe.Pointer(c__data))

	time := uint32(c_time)

	index := int(uintptr(data))
	callback := signalWidgetSelectionReceivedMap[index].callback
	callback(Data, time)
}

type signalWidgetSelectionRequestEventDetail struct {
	callback  WidgetSignalSelectionRequestEventCallback
	handlerID C.gulong
}

var signalWidgetSelectionRequestEventId int
var signalWidgetSelectionRequestEventMap = make(map[int]signalWidgetSelectionRequestEventDetail)
var signalWidgetSelectionRequestEventLock sync.RWMutex

// WidgetSignalSelectionRequestEventCallback is a callback function for a 'selection-request-event' signal emitted from a Widget.
type WidgetSignalSelectionRequestEventCallback func(event *gdk.EventSelection) bool

/*
ConnectSelectionRequestEvent connects the callback to the 'selection-request-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSelectionRequestEvent to remove it.
*/
func (recv *Widget) ConnectSelectionRequestEvent(callback WidgetSignalSelectionRequestEventCallback) int {
	signalWidgetSelectionRequestEventLock.Lock()
	defer signalWidgetSelectionRequestEventLock.Unlock()

	signalWidgetSelectionRequestEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_selection_request_event(instance, C.gpointer(uintptr(signalWidgetSelectionRequestEventId)))

	detail := signalWidgetSelectionRequestEventDetail{callback, handlerID}
	signalWidgetSelectionRequestEventMap[signalWidgetSelectionRequestEventId] = detail

	return signalWidgetSelectionRequestEventId
}

/*
DisconnectSelectionRequestEvent disconnects a callback from the 'selection-request-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSelectionRequestEvent.
*/
func (recv *Widget) DisconnectSelectionRequestEvent(connectionID int) {
	signalWidgetSelectionRequestEventLock.Lock()
	defer signalWidgetSelectionRequestEventLock.Unlock()

	detail, exists := signalWidgetSelectionRequestEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSelectionRequestEventMap, connectionID)
}

//export widget_selectionRequestEventHandler
func widget_selectionRequestEventHandler(_ *C.GObject, c_event *C.GdkEventSelection, data C.gpointer) C.gboolean {
	signalWidgetSelectionRequestEventLock.RLock()
	defer signalWidgetSelectionRequestEventLock.RUnlock()

	event := gdk.EventSelectionNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetSelectionRequestEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetShowDetail struct {
	callback  WidgetSignalShowCallback
	handlerID C.gulong
}

var signalWidgetShowId int
var signalWidgetShowMap = make(map[int]signalWidgetShowDetail)
var signalWidgetShowLock sync.RWMutex

// WidgetSignalShowCallback is a callback function for a 'show' signal emitted from a Widget.
type WidgetSignalShowCallback func()

/*
ConnectShow connects the callback to the 'show' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectShow to remove it.
*/
func (recv *Widget) ConnectShow(callback WidgetSignalShowCallback) int {
	signalWidgetShowLock.Lock()
	defer signalWidgetShowLock.Unlock()

	signalWidgetShowId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_show(instance, C.gpointer(uintptr(signalWidgetShowId)))

	detail := signalWidgetShowDetail{callback, handlerID}
	signalWidgetShowMap[signalWidgetShowId] = detail

	return signalWidgetShowId
}

/*
DisconnectShow disconnects a callback from the 'show' signal for the Widget.

The connectionID should be a value returned from a call to ConnectShow.
*/
func (recv *Widget) DisconnectShow(connectionID int) {
	signalWidgetShowLock.Lock()
	defer signalWidgetShowLock.Unlock()

	detail, exists := signalWidgetShowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetShowMap, connectionID)
}

//export widget_showHandler
func widget_showHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetShowLock.RLock()
	defer signalWidgetShowLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetShowMap[index].callback
	callback()
}

type signalWidgetShowHelpDetail struct {
	callback  WidgetSignalShowHelpCallback
	handlerID C.gulong
}

var signalWidgetShowHelpId int
var signalWidgetShowHelpMap = make(map[int]signalWidgetShowHelpDetail)
var signalWidgetShowHelpLock sync.RWMutex

// WidgetSignalShowHelpCallback is a callback function for a 'show-help' signal emitted from a Widget.
type WidgetSignalShowHelpCallback func(helpType WidgetHelpType) bool

/*
ConnectShowHelp connects the callback to the 'show-help' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectShowHelp to remove it.
*/
func (recv *Widget) ConnectShowHelp(callback WidgetSignalShowHelpCallback) int {
	signalWidgetShowHelpLock.Lock()
	defer signalWidgetShowHelpLock.Unlock()

	signalWidgetShowHelpId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_show_help(instance, C.gpointer(uintptr(signalWidgetShowHelpId)))

	detail := signalWidgetShowHelpDetail{callback, handlerID}
	signalWidgetShowHelpMap[signalWidgetShowHelpId] = detail

	return signalWidgetShowHelpId
}

/*
DisconnectShowHelp disconnects a callback from the 'show-help' signal for the Widget.

The connectionID should be a value returned from a call to ConnectShowHelp.
*/
func (recv *Widget) DisconnectShowHelp(connectionID int) {
	signalWidgetShowHelpLock.Lock()
	defer signalWidgetShowHelpLock.Unlock()

	detail, exists := signalWidgetShowHelpMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetShowHelpMap, connectionID)
}

//export widget_showHelpHandler
func widget_showHelpHandler(_ *C.GObject, c_help_type C.GtkWidgetHelpType, data C.gpointer) C.gboolean {
	signalWidgetShowHelpLock.RLock()
	defer signalWidgetShowHelpLock.RUnlock()

	helpType := WidgetHelpType(c_help_type)

	index := int(uintptr(data))
	callback := signalWidgetShowHelpMap[index].callback
	retGo := callback(helpType)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetSizeAllocateDetail struct {
	callback  WidgetSignalSizeAllocateCallback
	handlerID C.gulong
}

var signalWidgetSizeAllocateId int
var signalWidgetSizeAllocateMap = make(map[int]signalWidgetSizeAllocateDetail)
var signalWidgetSizeAllocateLock sync.RWMutex

// WidgetSignalSizeAllocateCallback is a callback function for a 'size-allocate' signal emitted from a Widget.
type WidgetSignalSizeAllocateCallback func(allocation *gdk.Rectangle)

/*
ConnectSizeAllocate connects the callback to the 'size-allocate' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectSizeAllocate to remove it.
*/
func (recv *Widget) ConnectSizeAllocate(callback WidgetSignalSizeAllocateCallback) int {
	signalWidgetSizeAllocateLock.Lock()
	defer signalWidgetSizeAllocateLock.Unlock()

	signalWidgetSizeAllocateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_size_allocate(instance, C.gpointer(uintptr(signalWidgetSizeAllocateId)))

	detail := signalWidgetSizeAllocateDetail{callback, handlerID}
	signalWidgetSizeAllocateMap[signalWidgetSizeAllocateId] = detail

	return signalWidgetSizeAllocateId
}

/*
DisconnectSizeAllocate disconnects a callback from the 'size-allocate' signal for the Widget.

The connectionID should be a value returned from a call to ConnectSizeAllocate.
*/
func (recv *Widget) DisconnectSizeAllocate(connectionID int) {
	signalWidgetSizeAllocateLock.Lock()
	defer signalWidgetSizeAllocateLock.Unlock()

	detail, exists := signalWidgetSizeAllocateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetSizeAllocateMap, connectionID)
}

//export widget_sizeAllocateHandler
func widget_sizeAllocateHandler(_ *C.GObject, c_allocation *C.GdkRectangle, data C.gpointer) {
	signalWidgetSizeAllocateLock.RLock()
	defer signalWidgetSizeAllocateLock.RUnlock()

	allocation := gdk.RectangleNewFromC(unsafe.Pointer(c_allocation))

	index := int(uintptr(data))
	callback := signalWidgetSizeAllocateMap[index].callback
	callback(allocation)
}

type signalWidgetStateChangedDetail struct {
	callback  WidgetSignalStateChangedCallback
	handlerID C.gulong
}

var signalWidgetStateChangedId int
var signalWidgetStateChangedMap = make(map[int]signalWidgetStateChangedDetail)
var signalWidgetStateChangedLock sync.RWMutex

// WidgetSignalStateChangedCallback is a callback function for a 'state-changed' signal emitted from a Widget.
type WidgetSignalStateChangedCallback func(state StateType)

/*
ConnectStateChanged connects the callback to the 'state-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStateChanged to remove it.
*/
func (recv *Widget) ConnectStateChanged(callback WidgetSignalStateChangedCallback) int {
	signalWidgetStateChangedLock.Lock()
	defer signalWidgetStateChangedLock.Unlock()

	signalWidgetStateChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_state_changed(instance, C.gpointer(uintptr(signalWidgetStateChangedId)))

	detail := signalWidgetStateChangedDetail{callback, handlerID}
	signalWidgetStateChangedMap[signalWidgetStateChangedId] = detail

	return signalWidgetStateChangedId
}

/*
DisconnectStateChanged disconnects a callback from the 'state-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStateChanged.
*/
func (recv *Widget) DisconnectStateChanged(connectionID int) {
	signalWidgetStateChangedLock.Lock()
	defer signalWidgetStateChangedLock.Unlock()

	detail, exists := signalWidgetStateChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStateChangedMap, connectionID)
}

//export widget_stateChangedHandler
func widget_stateChangedHandler(_ *C.GObject, c_state C.GtkStateType, data C.gpointer) {
	signalWidgetStateChangedLock.RLock()
	defer signalWidgetStateChangedLock.RUnlock()

	state := StateType(c_state)

	index := int(uintptr(data))
	callback := signalWidgetStateChangedMap[index].callback
	callback(state)
}

type signalWidgetStyleSetDetail struct {
	callback  WidgetSignalStyleSetCallback
	handlerID C.gulong
}

var signalWidgetStyleSetId int
var signalWidgetStyleSetMap = make(map[int]signalWidgetStyleSetDetail)
var signalWidgetStyleSetLock sync.RWMutex

// WidgetSignalStyleSetCallback is a callback function for a 'style-set' signal emitted from a Widget.
type WidgetSignalStyleSetCallback func(previousStyle *Style)

/*
ConnectStyleSet connects the callback to the 'style-set' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStyleSet to remove it.
*/
func (recv *Widget) ConnectStyleSet(callback WidgetSignalStyleSetCallback) int {
	signalWidgetStyleSetLock.Lock()
	defer signalWidgetStyleSetLock.Unlock()

	signalWidgetStyleSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_style_set(instance, C.gpointer(uintptr(signalWidgetStyleSetId)))

	detail := signalWidgetStyleSetDetail{callback, handlerID}
	signalWidgetStyleSetMap[signalWidgetStyleSetId] = detail

	return signalWidgetStyleSetId
}

/*
DisconnectStyleSet disconnects a callback from the 'style-set' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStyleSet.
*/
func (recv *Widget) DisconnectStyleSet(connectionID int) {
	signalWidgetStyleSetLock.Lock()
	defer signalWidgetStyleSetLock.Unlock()

	detail, exists := signalWidgetStyleSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStyleSetMap, connectionID)
}

//export widget_styleSetHandler
func widget_styleSetHandler(_ *C.GObject, c_previous_style *C.GtkStyle, data C.gpointer) {
	signalWidgetStyleSetLock.RLock()
	defer signalWidgetStyleSetLock.RUnlock()

	previousStyle := StyleNewFromC(unsafe.Pointer(c_previous_style))

	index := int(uintptr(data))
	callback := signalWidgetStyleSetMap[index].callback
	callback(previousStyle)
}

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

type signalWidgetUnmapDetail struct {
	callback  WidgetSignalUnmapCallback
	handlerID C.gulong
}

var signalWidgetUnmapId int
var signalWidgetUnmapMap = make(map[int]signalWidgetUnmapDetail)
var signalWidgetUnmapLock sync.RWMutex

// WidgetSignalUnmapCallback is a callback function for a 'unmap' signal emitted from a Widget.
type WidgetSignalUnmapCallback func()

/*
ConnectUnmap connects the callback to the 'unmap' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectUnmap to remove it.
*/
func (recv *Widget) ConnectUnmap(callback WidgetSignalUnmapCallback) int {
	signalWidgetUnmapLock.Lock()
	defer signalWidgetUnmapLock.Unlock()

	signalWidgetUnmapId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_unmap(instance, C.gpointer(uintptr(signalWidgetUnmapId)))

	detail := signalWidgetUnmapDetail{callback, handlerID}
	signalWidgetUnmapMap[signalWidgetUnmapId] = detail

	return signalWidgetUnmapId
}

/*
DisconnectUnmap disconnects a callback from the 'unmap' signal for the Widget.

The connectionID should be a value returned from a call to ConnectUnmap.
*/
func (recv *Widget) DisconnectUnmap(connectionID int) {
	signalWidgetUnmapLock.Lock()
	defer signalWidgetUnmapLock.Unlock()

	detail, exists := signalWidgetUnmapMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetUnmapMap, connectionID)
}

//export widget_unmapHandler
func widget_unmapHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetUnmapLock.RLock()
	defer signalWidgetUnmapLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetUnmapMap[index].callback
	callback()
}

type signalWidgetUnmapEventDetail struct {
	callback  WidgetSignalUnmapEventCallback
	handlerID C.gulong
}

var signalWidgetUnmapEventId int
var signalWidgetUnmapEventMap = make(map[int]signalWidgetUnmapEventDetail)
var signalWidgetUnmapEventLock sync.RWMutex

// WidgetSignalUnmapEventCallback is a callback function for a 'unmap-event' signal emitted from a Widget.
type WidgetSignalUnmapEventCallback func(event *gdk.EventAny) bool

/*
ConnectUnmapEvent connects the callback to the 'unmap-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectUnmapEvent to remove it.
*/
func (recv *Widget) ConnectUnmapEvent(callback WidgetSignalUnmapEventCallback) int {
	signalWidgetUnmapEventLock.Lock()
	defer signalWidgetUnmapEventLock.Unlock()

	signalWidgetUnmapEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_unmap_event(instance, C.gpointer(uintptr(signalWidgetUnmapEventId)))

	detail := signalWidgetUnmapEventDetail{callback, handlerID}
	signalWidgetUnmapEventMap[signalWidgetUnmapEventId] = detail

	return signalWidgetUnmapEventId
}

/*
DisconnectUnmapEvent disconnects a callback from the 'unmap-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectUnmapEvent.
*/
func (recv *Widget) DisconnectUnmapEvent(connectionID int) {
	signalWidgetUnmapEventLock.Lock()
	defer signalWidgetUnmapEventLock.Unlock()

	detail, exists := signalWidgetUnmapEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetUnmapEventMap, connectionID)
}

//export widget_unmapEventHandler
func widget_unmapEventHandler(_ *C.GObject, c_event *C.GdkEventAny, data C.gpointer) C.gboolean {
	signalWidgetUnmapEventLock.RLock()
	defer signalWidgetUnmapEventLock.RUnlock()

	event := gdk.EventAnyNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetUnmapEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetUnrealizeDetail struct {
	callback  WidgetSignalUnrealizeCallback
	handlerID C.gulong
}

var signalWidgetUnrealizeId int
var signalWidgetUnrealizeMap = make(map[int]signalWidgetUnrealizeDetail)
var signalWidgetUnrealizeLock sync.RWMutex

// WidgetSignalUnrealizeCallback is a callback function for a 'unrealize' signal emitted from a Widget.
type WidgetSignalUnrealizeCallback func()

/*
ConnectUnrealize connects the callback to the 'unrealize' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectUnrealize to remove it.
*/
func (recv *Widget) ConnectUnrealize(callback WidgetSignalUnrealizeCallback) int {
	signalWidgetUnrealizeLock.Lock()
	defer signalWidgetUnrealizeLock.Unlock()

	signalWidgetUnrealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_unrealize(instance, C.gpointer(uintptr(signalWidgetUnrealizeId)))

	detail := signalWidgetUnrealizeDetail{callback, handlerID}
	signalWidgetUnrealizeMap[signalWidgetUnrealizeId] = detail

	return signalWidgetUnrealizeId
}

/*
DisconnectUnrealize disconnects a callback from the 'unrealize' signal for the Widget.

The connectionID should be a value returned from a call to ConnectUnrealize.
*/
func (recv *Widget) DisconnectUnrealize(connectionID int) {
	signalWidgetUnrealizeLock.Lock()
	defer signalWidgetUnrealizeLock.Unlock()

	detail, exists := signalWidgetUnrealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetUnrealizeMap, connectionID)
}

//export widget_unrealizeHandler
func widget_unrealizeHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetUnrealizeLock.RLock()
	defer signalWidgetUnrealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetUnrealizeMap[index].callback
	callback()
}

type signalWidgetVisibilityNotifyEventDetail struct {
	callback  WidgetSignalVisibilityNotifyEventCallback
	handlerID C.gulong
}

var signalWidgetVisibilityNotifyEventId int
var signalWidgetVisibilityNotifyEventMap = make(map[int]signalWidgetVisibilityNotifyEventDetail)
var signalWidgetVisibilityNotifyEventLock sync.RWMutex

// WidgetSignalVisibilityNotifyEventCallback is a callback function for a 'visibility-notify-event' signal emitted from a Widget.
type WidgetSignalVisibilityNotifyEventCallback func(event *gdk.EventVisibility) bool

/*
ConnectVisibilityNotifyEvent connects the callback to the 'visibility-notify-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectVisibilityNotifyEvent to remove it.
*/
func (recv *Widget) ConnectVisibilityNotifyEvent(callback WidgetSignalVisibilityNotifyEventCallback) int {
	signalWidgetVisibilityNotifyEventLock.Lock()
	defer signalWidgetVisibilityNotifyEventLock.Unlock()

	signalWidgetVisibilityNotifyEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_visibility_notify_event(instance, C.gpointer(uintptr(signalWidgetVisibilityNotifyEventId)))

	detail := signalWidgetVisibilityNotifyEventDetail{callback, handlerID}
	signalWidgetVisibilityNotifyEventMap[signalWidgetVisibilityNotifyEventId] = detail

	return signalWidgetVisibilityNotifyEventId
}

/*
DisconnectVisibilityNotifyEvent disconnects a callback from the 'visibility-notify-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectVisibilityNotifyEvent.
*/
func (recv *Widget) DisconnectVisibilityNotifyEvent(connectionID int) {
	signalWidgetVisibilityNotifyEventLock.Lock()
	defer signalWidgetVisibilityNotifyEventLock.Unlock()

	detail, exists := signalWidgetVisibilityNotifyEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetVisibilityNotifyEventMap, connectionID)
}

//export widget_visibilityNotifyEventHandler
func widget_visibilityNotifyEventHandler(_ *C.GObject, c_event *C.GdkEventVisibility, data C.gpointer) C.gboolean {
	signalWidgetVisibilityNotifyEventLock.RLock()
	defer signalWidgetVisibilityNotifyEventLock.RUnlock()

	event := gdk.EventVisibilityNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetVisibilityNotifyEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetWindowStateEventDetail struct {
	callback  WidgetSignalWindowStateEventCallback
	handlerID C.gulong
}

var signalWidgetWindowStateEventId int
var signalWidgetWindowStateEventMap = make(map[int]signalWidgetWindowStateEventDetail)
var signalWidgetWindowStateEventLock sync.RWMutex

// WidgetSignalWindowStateEventCallback is a callback function for a 'window-state-event' signal emitted from a Widget.
type WidgetSignalWindowStateEventCallback func(event *gdk.EventWindowState) bool

/*
ConnectWindowStateEvent connects the callback to the 'window-state-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectWindowStateEvent to remove it.
*/
func (recv *Widget) ConnectWindowStateEvent(callback WidgetSignalWindowStateEventCallback) int {
	signalWidgetWindowStateEventLock.Lock()
	defer signalWidgetWindowStateEventLock.Unlock()

	signalWidgetWindowStateEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_window_state_event(instance, C.gpointer(uintptr(signalWidgetWindowStateEventId)))

	detail := signalWidgetWindowStateEventDetail{callback, handlerID}
	signalWidgetWindowStateEventMap[signalWidgetWindowStateEventId] = detail

	return signalWidgetWindowStateEventId
}

/*
DisconnectWindowStateEvent disconnects a callback from the 'window-state-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectWindowStateEvent.
*/
func (recv *Widget) DisconnectWindowStateEvent(connectionID int) {
	signalWidgetWindowStateEventLock.Lock()
	defer signalWidgetWindowStateEventLock.Unlock()

	detail, exists := signalWidgetWindowStateEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetWindowStateEventMap, connectionID)
}

//export widget_windowStateEventHandler
func widget_windowStateEventHandler(_ *C.GObject, c_event *C.GdkEventWindowState, data C.gpointer) C.gboolean {
	signalWidgetWindowStateEventLock.RLock()
	defer signalWidgetWindowStateEventLock.RUnlock()

	event := gdk.EventWindowStateNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetWindowStateEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// Blacklisted : gtk_widget_get_default_direction

// Blacklisted : gtk_widget_get_default_style

// Blacklisted : gtk_widget_pop_composite_child

// Blacklisted : gtk_widget_push_composite_child

// Blacklisted : gtk_widget_set_default_direction

// Blacklisted : gtk_widget_activate

// Blacklisted : gtk_widget_add_accelerator

// Blacklisted : gtk_widget_add_events

// Blacklisted : gtk_widget_child_focus

// Blacklisted : gtk_widget_child_notify

// Blacklisted : gtk_widget_class_path

// Blacklisted : gtk_widget_compute_expand

// Blacklisted : gtk_widget_create_pango_context

// Blacklisted : gtk_widget_create_pango_layout

// Destroy is a wrapper around the C function gtk_widget_destroy.
func (recv *Widget) Destroy() {
	C.gtk_widget_destroy((*C.GtkWidget)(recv.native))

	return
}

// Blacklisted : gtk_widget_destroyed

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_drag_check_threshold

// Blacklisted : gtk_drag_dest_find_target

// Blacklisted : gtk_drag_dest_get_target_list

// Unsupported : gtk_drag_dest_set : unsupported parameter targets :

// Blacklisted : gtk_drag_dest_set_proxy

// Blacklisted : gtk_drag_dest_set_target_list

// Blacklisted : gtk_drag_dest_unset

// Blacklisted : gtk_drag_get_data

// Blacklisted : gtk_drag_highlight

// Unsupported : gtk_drag_source_set : unsupported parameter targets :

// Blacklisted : gtk_drag_source_set_icon_pixbuf

// Blacklisted : gtk_drag_source_set_icon_stock

// Blacklisted : gtk_drag_source_unset

// Blacklisted : gtk_drag_unhighlight

// Blacklisted : gtk_widget_ensure_style

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_widget_freeze_child_notify

// Blacklisted : gtk_widget_get_accessible

// Blacklisted : gtk_widget_get_allocated_height

// Blacklisted : gtk_widget_get_allocated_width

// Blacklisted : gtk_widget_get_ancestor

// Blacklisted : gtk_widget_get_child_requisition

// Blacklisted : gtk_widget_get_child_visible

// Blacklisted : gtk_widget_get_composite_name

// Blacklisted : gtk_widget_get_direction

// Blacklisted : gtk_widget_get_events

// Blacklisted : gtk_widget_get_halign

// Blacklisted : gtk_widget_get_hexpand

// Blacklisted : gtk_widget_get_hexpand_set

// Blacklisted : gtk_widget_get_modifier_style

// Blacklisted : gtk_widget_get_name

// Blacklisted : gtk_widget_get_pango_context

// Blacklisted : gtk_widget_get_parent

// Blacklisted : gtk_widget_get_parent_window

// Blacklisted : gtk_widget_get_path

// Blacklisted : gtk_widget_get_pointer

// Blacklisted : gtk_widget_get_settings

// Blacklisted : gtk_widget_get_size_request

// Blacklisted : gtk_widget_get_style

// Blacklisted : gtk_widget_get_style_context

// Blacklisted : gtk_widget_get_support_multidevice

// Blacklisted : gtk_widget_get_template_child

// Blacklisted : gtk_widget_get_toplevel

// Blacklisted : gtk_widget_get_valign

// Blacklisted : gtk_widget_get_vexpand

// Blacklisted : gtk_widget_get_vexpand_set

// Blacklisted : gtk_widget_get_visual

// Blacklisted : gtk_grab_add

// Blacklisted : gtk_widget_grab_default

// Blacklisted : gtk_widget_grab_focus

// Blacklisted : gtk_grab_remove

// Blacklisted : gtk_widget_hide

// Blacklisted : gtk_widget_hide_on_delete

// Blacklisted : gtk_widget_in_destruction

// Blacklisted : gtk_widget_intersect

// Blacklisted : gtk_widget_is_ancestor

// Blacklisted : gtk_widget_is_focus

// Blacklisted : gtk_widget_list_accel_closures

// Blacklisted : gtk_widget_map

// Blacklisted : gtk_widget_mnemonic_activate

// Blacklisted : gtk_widget_modify_base

// Blacklisted : gtk_widget_modify_bg

// Blacklisted : gtk_widget_modify_fg

// Blacklisted : gtk_widget_modify_font

// Blacklisted : gtk_widget_modify_style

// Blacklisted : gtk_widget_modify_text

// Blacklisted : gtk_widget_path

// Blacklisted : gtk_widget_queue_compute_expand

// Blacklisted : gtk_widget_queue_draw

// Blacklisted : gtk_widget_queue_draw_area

// Blacklisted : gtk_widget_queue_resize

// Blacklisted : gtk_widget_realize

// Blacklisted : gtk_widget_region_intersect

// Blacklisted : gtk_widget_remove_accelerator

// Blacklisted : gtk_widget_render_icon

// Blacklisted : gtk_widget_reparent

// Blacklisted : gtk_widget_reset_rc_styles

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_widget_set_accel_path

// Blacklisted : gtk_widget_set_app_paintable

// Blacklisted : gtk_widget_set_child_visible

// Blacklisted : gtk_widget_set_composite_name

// Blacklisted : gtk_widget_set_direction

// Blacklisted : gtk_widget_set_double_buffered

// Blacklisted : gtk_widget_set_events

// Blacklisted : gtk_widget_set_halign

// Blacklisted : gtk_widget_set_hexpand

// Blacklisted : gtk_widget_set_hexpand_set

// Blacklisted : gtk_widget_set_name

// Blacklisted : gtk_widget_set_parent

// Blacklisted : gtk_widget_set_parent_window

// Blacklisted : gtk_widget_set_redraw_on_allocate

// Blacklisted : gtk_widget_set_sensitive

// Blacklisted : gtk_widget_set_size_request

// Blacklisted : gtk_widget_set_state

// Blacklisted : gtk_widget_set_style

// Blacklisted : gtk_widget_set_valign

// Blacklisted : gtk_widget_set_vexpand

// Blacklisted : gtk_widget_set_vexpand_set

// Blacklisted : gtk_widget_set_visual

// Blacklisted : gtk_widget_show

// ShowAll is a wrapper around the C function gtk_widget_show_all.
func (recv *Widget) ShowAll() {
	C.gtk_widget_show_all((*C.GtkWidget)(recv.native))

	return
}

// Blacklisted : gtk_widget_show_now

// Blacklisted : gtk_widget_size_allocate

// Blacklisted : gtk_widget_size_request

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// Blacklisted : gtk_widget_style_get_property

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : gtk_widget_thaw_child_notify

// Blacklisted : gtk_widget_translate_coordinates

// Blacklisted : gtk_widget_unmap

// Blacklisted : gtk_widget_unparent

// Blacklisted : gtk_widget_unrealize

// ImplementorIface returns the ImplementorIface interface implemented by Widget
func (recv *Widget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Widget
func (recv *Widget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkWidgetAccessible

// Window is a wrapper around the C record GtkWindow.
type Window struct {
	native *C.GtkWindow
	// bin : record
	// priv : record
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.GtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Window) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}

// Bin upcasts to *Bin
func (recv *Window) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Window) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Window) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Window) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Window) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitrary Object to Window.
// Exercise care, as this is a potentially dangerous function if the Object is not a Window.
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromC(object.ToC())
}

type signalWindowActivateDefaultDetail struct {
	callback  WindowSignalActivateDefaultCallback
	handlerID C.gulong
}

var signalWindowActivateDefaultId int
var signalWindowActivateDefaultMap = make(map[int]signalWindowActivateDefaultDetail)
var signalWindowActivateDefaultLock sync.RWMutex

// WindowSignalActivateDefaultCallback is a callback function for a 'activate-default' signal emitted from a Window.
type WindowSignalActivateDefaultCallback func()

/*
ConnectActivateDefault connects the callback to the 'activate-default' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectActivateDefault to remove it.
*/
func (recv *Window) ConnectActivateDefault(callback WindowSignalActivateDefaultCallback) int {
	signalWindowActivateDefaultLock.Lock()
	defer signalWindowActivateDefaultLock.Unlock()

	signalWindowActivateDefaultId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_activate_default(instance, C.gpointer(uintptr(signalWindowActivateDefaultId)))

	detail := signalWindowActivateDefaultDetail{callback, handlerID}
	signalWindowActivateDefaultMap[signalWindowActivateDefaultId] = detail

	return signalWindowActivateDefaultId
}

/*
DisconnectActivateDefault disconnects a callback from the 'activate-default' signal for the Window.

The connectionID should be a value returned from a call to ConnectActivateDefault.
*/
func (recv *Window) DisconnectActivateDefault(connectionID int) {
	signalWindowActivateDefaultLock.Lock()
	defer signalWindowActivateDefaultLock.Unlock()

	detail, exists := signalWindowActivateDefaultMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowActivateDefaultMap, connectionID)
}

//export window_activateDefaultHandler
func window_activateDefaultHandler(_ *C.GObject, data C.gpointer) {
	signalWindowActivateDefaultLock.RLock()
	defer signalWindowActivateDefaultLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWindowActivateDefaultMap[index].callback
	callback()
}

type signalWindowActivateFocusDetail struct {
	callback  WindowSignalActivateFocusCallback
	handlerID C.gulong
}

var signalWindowActivateFocusId int
var signalWindowActivateFocusMap = make(map[int]signalWindowActivateFocusDetail)
var signalWindowActivateFocusLock sync.RWMutex

// WindowSignalActivateFocusCallback is a callback function for a 'activate-focus' signal emitted from a Window.
type WindowSignalActivateFocusCallback func()

/*
ConnectActivateFocus connects the callback to the 'activate-focus' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectActivateFocus to remove it.
*/
func (recv *Window) ConnectActivateFocus(callback WindowSignalActivateFocusCallback) int {
	signalWindowActivateFocusLock.Lock()
	defer signalWindowActivateFocusLock.Unlock()

	signalWindowActivateFocusId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_activate_focus(instance, C.gpointer(uintptr(signalWindowActivateFocusId)))

	detail := signalWindowActivateFocusDetail{callback, handlerID}
	signalWindowActivateFocusMap[signalWindowActivateFocusId] = detail

	return signalWindowActivateFocusId
}

/*
DisconnectActivateFocus disconnects a callback from the 'activate-focus' signal for the Window.

The connectionID should be a value returned from a call to ConnectActivateFocus.
*/
func (recv *Window) DisconnectActivateFocus(connectionID int) {
	signalWindowActivateFocusLock.Lock()
	defer signalWindowActivateFocusLock.Unlock()

	detail, exists := signalWindowActivateFocusMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowActivateFocusMap, connectionID)
}

//export window_activateFocusHandler
func window_activateFocusHandler(_ *C.GObject, data C.gpointer) {
	signalWindowActivateFocusLock.RLock()
	defer signalWindowActivateFocusLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWindowActivateFocusMap[index].callback
	callback()
}

type signalWindowEnableDebuggingDetail struct {
	callback  WindowSignalEnableDebuggingCallback
	handlerID C.gulong
}

var signalWindowEnableDebuggingId int
var signalWindowEnableDebuggingMap = make(map[int]signalWindowEnableDebuggingDetail)
var signalWindowEnableDebuggingLock sync.RWMutex

// WindowSignalEnableDebuggingCallback is a callback function for a 'enable-debugging' signal emitted from a Window.
type WindowSignalEnableDebuggingCallback func(toggle bool) bool

/*
ConnectEnableDebugging connects the callback to the 'enable-debugging' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectEnableDebugging to remove it.
*/
func (recv *Window) ConnectEnableDebugging(callback WindowSignalEnableDebuggingCallback) int {
	signalWindowEnableDebuggingLock.Lock()
	defer signalWindowEnableDebuggingLock.Unlock()

	signalWindowEnableDebuggingId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_enable_debugging(instance, C.gpointer(uintptr(signalWindowEnableDebuggingId)))

	detail := signalWindowEnableDebuggingDetail{callback, handlerID}
	signalWindowEnableDebuggingMap[signalWindowEnableDebuggingId] = detail

	return signalWindowEnableDebuggingId
}

/*
DisconnectEnableDebugging disconnects a callback from the 'enable-debugging' signal for the Window.

The connectionID should be a value returned from a call to ConnectEnableDebugging.
*/
func (recv *Window) DisconnectEnableDebugging(connectionID int) {
	signalWindowEnableDebuggingLock.Lock()
	defer signalWindowEnableDebuggingLock.Unlock()

	detail, exists := signalWindowEnableDebuggingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowEnableDebuggingMap, connectionID)
}

//export window_enableDebuggingHandler
func window_enableDebuggingHandler(_ *C.GObject, c_toggle C.gboolean, data C.gpointer) C.gboolean {
	signalWindowEnableDebuggingLock.RLock()
	defer signalWindowEnableDebuggingLock.RUnlock()

	toggle := c_toggle == C.TRUE

	index := int(uintptr(data))
	callback := signalWindowEnableDebuggingMap[index].callback
	retGo := callback(toggle)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWindowKeysChangedDetail struct {
	callback  WindowSignalKeysChangedCallback
	handlerID C.gulong
}

var signalWindowKeysChangedId int
var signalWindowKeysChangedMap = make(map[int]signalWindowKeysChangedDetail)
var signalWindowKeysChangedLock sync.RWMutex

// WindowSignalKeysChangedCallback is a callback function for a 'keys-changed' signal emitted from a Window.
type WindowSignalKeysChangedCallback func()

/*
ConnectKeysChanged connects the callback to the 'keys-changed' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectKeysChanged to remove it.
*/
func (recv *Window) ConnectKeysChanged(callback WindowSignalKeysChangedCallback) int {
	signalWindowKeysChangedLock.Lock()
	defer signalWindowKeysChangedLock.Unlock()

	signalWindowKeysChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_keys_changed(instance, C.gpointer(uintptr(signalWindowKeysChangedId)))

	detail := signalWindowKeysChangedDetail{callback, handlerID}
	signalWindowKeysChangedMap[signalWindowKeysChangedId] = detail

	return signalWindowKeysChangedId
}

/*
DisconnectKeysChanged disconnects a callback from the 'keys-changed' signal for the Window.

The connectionID should be a value returned from a call to ConnectKeysChanged.
*/
func (recv *Window) DisconnectKeysChanged(connectionID int) {
	signalWindowKeysChangedLock.Lock()
	defer signalWindowKeysChangedLock.Unlock()

	detail, exists := signalWindowKeysChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowKeysChangedMap, connectionID)
}

//export window_keysChangedHandler
func window_keysChangedHandler(_ *C.GObject, data C.gpointer) {
	signalWindowKeysChangedLock.RLock()
	defer signalWindowKeysChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWindowKeysChangedMap[index].callback
	callback()
}

type signalWindowSetFocusDetail struct {
	callback  WindowSignalSetFocusCallback
	handlerID C.gulong
}

var signalWindowSetFocusId int
var signalWindowSetFocusMap = make(map[int]signalWindowSetFocusDetail)
var signalWindowSetFocusLock sync.RWMutex

// WindowSignalSetFocusCallback is a callback function for a 'set-focus' signal emitted from a Window.
type WindowSignalSetFocusCallback func(object *Widget)

/*
ConnectSetFocus connects the callback to the 'set-focus' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectSetFocus to remove it.
*/
func (recv *Window) ConnectSetFocus(callback WindowSignalSetFocusCallback) int {
	signalWindowSetFocusLock.Lock()
	defer signalWindowSetFocusLock.Unlock()

	signalWindowSetFocusId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_set_focus(instance, C.gpointer(uintptr(signalWindowSetFocusId)))

	detail := signalWindowSetFocusDetail{callback, handlerID}
	signalWindowSetFocusMap[signalWindowSetFocusId] = detail

	return signalWindowSetFocusId
}

/*
DisconnectSetFocus disconnects a callback from the 'set-focus' signal for the Window.

The connectionID should be a value returned from a call to ConnectSetFocus.
*/
func (recv *Window) DisconnectSetFocus(connectionID int) {
	signalWindowSetFocusLock.Lock()
	defer signalWindowSetFocusLock.Unlock()

	detail, exists := signalWindowSetFocusMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowSetFocusMap, connectionID)
}

//export window_setFocusHandler
func window_setFocusHandler(_ *C.GObject, c_object *C.GtkWidget, data C.gpointer) {
	signalWindowSetFocusLock.RLock()
	defer signalWindowSetFocusLock.RUnlock()

	object := WidgetNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalWindowSetFocusMap[index].callback
	callback(object)
}

// WindowNew is a wrapper around the C function gtk_window_new.
func WindowNew(type_ WindowType) *Window {
	c_type := (C.GtkWindowType)(type_)

	retC := C.gtk_window_new(c_type)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gtk_window_get_default_icon_list

// Blacklisted : gtk_window_list_toplevels

// Blacklisted : gtk_window_set_default_icon_list

// Blacklisted : gtk_window_activate_default

// Blacklisted : gtk_window_activate_focus

// Blacklisted : gtk_window_add_accel_group

// Blacklisted : gtk_window_add_mnemonic

// Blacklisted : gtk_window_begin_move_drag

// Blacklisted : gtk_window_begin_resize_drag

// Blacklisted : gtk_window_deiconify

// Blacklisted : gtk_window_get_decorated

// Blacklisted : gtk_window_get_default_size

// Blacklisted : gtk_window_get_destroy_with_parent

// Blacklisted : gtk_window_get_focus

// Blacklisted : gtk_window_get_gravity

// Blacklisted : gtk_window_get_icon

// Blacklisted : gtk_window_get_icon_list

// Blacklisted : gtk_window_get_mnemonic_modifier

// Blacklisted : gtk_window_get_modal

// Blacklisted : gtk_window_get_position

// Blacklisted : gtk_window_get_resizable

// Blacklisted : gtk_window_get_role

// Blacklisted : gtk_window_get_size

// GetTitle is a wrapper around the C function gtk_window_get_title.
func (recv *Window) GetTitle() string {
	retC := C.gtk_window_get_title((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_window_get_transient_for

// Blacklisted : gtk_window_get_type_hint

// Blacklisted : gtk_window_has_group

// Blacklisted : gtk_window_iconify

// Blacklisted : gtk_window_maximize

// Blacklisted : gtk_window_mnemonic_activate

// Blacklisted : gtk_window_move

// Blacklisted : gtk_window_parse_geometry

// Blacklisted : gtk_window_present

// Blacklisted : gtk_window_remove_accel_group

// Blacklisted : gtk_window_remove_mnemonic

// Blacklisted : gtk_window_reshow_with_initial_size

// Blacklisted : gtk_window_resize

// Blacklisted : gtk_window_set_decorated

// Blacklisted : gtk_window_set_default

// Blacklisted : gtk_window_set_default_size

// Blacklisted : gtk_window_set_destroy_with_parent

// Blacklisted : gtk_window_set_focus

// Blacklisted : gtk_window_set_geometry_hints

// Blacklisted : gtk_window_set_gravity

// Blacklisted : gtk_window_set_icon

// Blacklisted : gtk_window_set_icon_list

// Blacklisted : gtk_window_set_mnemonic_modifier

// Blacklisted : gtk_window_set_modal

// Blacklisted : gtk_window_set_position

// Blacklisted : gtk_window_set_resizable

// Blacklisted : gtk_window_set_role

// SetTitle is a wrapper around the C function gtk_window_set_title.
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_window_set_title((*C.GtkWindow)(recv.native), c_title)

	return
}

// Blacklisted : gtk_window_set_transient_for

// Blacklisted : gtk_window_set_type_hint

// Blacklisted : gtk_window_set_wmclass

// Blacklisted : gtk_window_stick

// Blacklisted : gtk_window_unmaximize

// Blacklisted : gtk_window_unstick

// ImplementorIface returns the ImplementorIface interface implemented by Window
func (recv *Window) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Window
func (recv *Window) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkWindowAccessible

// Blacklisted : GtkWindowGroup
