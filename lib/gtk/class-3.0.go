// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
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

	void cellarea_addEditableHandler(GObject *, GtkCellRenderer *, GtkCellEditable *, GdkRectangle *, gchar*, gpointer);

	static gulong CellArea_signal_connect_add_editable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add-editable", G_CALLBACK(cellarea_addEditableHandler), data);
	}

*/
/*

	void cellarea_applyAttributesHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gboolean, gboolean, gpointer);

	static gulong CellArea_signal_connect_apply_attributes(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "apply-attributes", G_CALLBACK(cellarea_applyAttributesHandler), data);
	}

*/
/*

	void cellarea_focusChangedHandler(GObject *, GtkCellRenderer *, gchar*, gpointer);

	static gulong CellArea_signal_connect_focus_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-changed", G_CALLBACK(cellarea_focusChangedHandler), data);
	}

*/
/*

	void cellarea_removeEditableHandler(GObject *, GtkCellRenderer *, GtkCellEditable *, gpointer);

	static gulong CellArea_signal_connect_remove_editable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-editable", G_CALLBACK(cellarea_removeEditableHandler), data);
	}

*/
/*

	void stylecontext_changedHandler(GObject *, gpointer);

	static gulong StyleContext_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(stylecontext_changedHandler), data);
	}

*/
/*

	gboolean widget_drawHandler(GObject *, cairo_t *, gpointer);

	static gulong Widget_signal_connect_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw", G_CALLBACK(widget_drawHandler), data);
	}

*/
/*

	void widget_stateFlagsChangedHandler(GObject *, GtkStateFlags, gpointer);

	static gulong Widget_signal_connect_state_flags_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-flags-changed", G_CALLBACK(widget_stateFlagsChangedHandler), data);
	}

*/
/*

	void widget_styleUpdatedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_style_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-updated", G_CALLBACK(widget_styleUpdatedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_about_dialog_get_license_type

// Blacklisted : gtk_about_dialog_set_license_type

// Blacklisted : gtk_app_chooser_button_new

// Blacklisted : gtk_app_chooser_button_append_custom_item

// Blacklisted : gtk_app_chooser_button_append_separator

// Blacklisted : gtk_app_chooser_button_get_show_dialog_item

// Blacklisted : gtk_app_chooser_button_set_active_custom_item

// Blacklisted : gtk_app_chooser_button_set_show_dialog_item

// Blacklisted : gtk_app_chooser_dialog_new

// Blacklisted : gtk_app_chooser_dialog_new_for_content_type

// Blacklisted : gtk_app_chooser_dialog_get_widget

// Blacklisted : gtk_app_chooser_widget_new

// Blacklisted : gtk_app_chooser_widget_get_default_text

// Blacklisted : gtk_app_chooser_widget_get_show_all

// Blacklisted : gtk_app_chooser_widget_get_show_default

// Blacklisted : gtk_app_chooser_widget_get_show_fallback

// Blacklisted : gtk_app_chooser_widget_get_show_other

// Blacklisted : gtk_app_chooser_widget_get_show_recommended

// Blacklisted : gtk_app_chooser_widget_set_show_all

// Blacklisted : gtk_app_chooser_widget_set_show_default

// Blacklisted : gtk_app_chooser_widget_set_show_fallback

// Blacklisted : gtk_app_chooser_widget_set_show_other

// Blacklisted : gtk_app_chooser_widget_set_show_recommended

// ApplicationNew is a wrapper around the C function gtk_application_new.
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : gtk_application_add_window

// Blacklisted : gtk_application_get_windows

// Blacklisted : gtk_application_remove_window

// Blacklisted : gtk_assistant_next_page

// Blacklisted : gtk_assistant_previous_page

// Blacklisted : gtk_box_new

// Blacklisted : gtk_button_box_new

// Blacklisted : gtk_calendar_get_day_is_marked

type signalCellAreaAddEditableDetail struct {
	callback  CellAreaSignalAddEditableCallback
	handlerID C.gulong
}

var signalCellAreaAddEditableId int
var signalCellAreaAddEditableMap = make(map[int]signalCellAreaAddEditableDetail)
var signalCellAreaAddEditableLock sync.RWMutex

// CellAreaSignalAddEditableCallback is a callback function for a 'add-editable' signal emitted from a CellArea.
type CellAreaSignalAddEditableCallback func(renderer *CellRenderer, editable *CellEditable, cellArea *gdk.Rectangle, path string)

/*
ConnectAddEditable connects the callback to the 'add-editable' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectAddEditable to remove it.
*/
func (recv *CellArea) ConnectAddEditable(callback CellAreaSignalAddEditableCallback) int {
	signalCellAreaAddEditableLock.Lock()
	defer signalCellAreaAddEditableLock.Unlock()

	signalCellAreaAddEditableId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_add_editable(instance, C.gpointer(uintptr(signalCellAreaAddEditableId)))

	detail := signalCellAreaAddEditableDetail{callback, handlerID}
	signalCellAreaAddEditableMap[signalCellAreaAddEditableId] = detail

	return signalCellAreaAddEditableId
}

/*
DisconnectAddEditable disconnects a callback from the 'add-editable' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectAddEditable.
*/
func (recv *CellArea) DisconnectAddEditable(connectionID int) {
	signalCellAreaAddEditableLock.Lock()
	defer signalCellAreaAddEditableLock.Unlock()

	detail, exists := signalCellAreaAddEditableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaAddEditableMap, connectionID)
}

//export cellarea_addEditableHandler
func cellarea_addEditableHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_editable *C.GtkCellEditable, c_cell_area *C.GdkRectangle, c_path *C.gchar, data C.gpointer) {
	signalCellAreaAddEditableLock.RLock()
	defer signalCellAreaAddEditableLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	cellArea := gdk.RectangleNewFromC(unsafe.Pointer(c_cell_area))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellAreaAddEditableMap[index].callback
	callback(renderer, editable, cellArea, path)
}

type signalCellAreaApplyAttributesDetail struct {
	callback  CellAreaSignalApplyAttributesCallback
	handlerID C.gulong
}

var signalCellAreaApplyAttributesId int
var signalCellAreaApplyAttributesMap = make(map[int]signalCellAreaApplyAttributesDetail)
var signalCellAreaApplyAttributesLock sync.RWMutex

// CellAreaSignalApplyAttributesCallback is a callback function for a 'apply-attributes' signal emitted from a CellArea.
type CellAreaSignalApplyAttributesCallback func(model *TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)

/*
ConnectApplyAttributes connects the callback to the 'apply-attributes' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectApplyAttributes to remove it.
*/
func (recv *CellArea) ConnectApplyAttributes(callback CellAreaSignalApplyAttributesCallback) int {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	signalCellAreaApplyAttributesId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_apply_attributes(instance, C.gpointer(uintptr(signalCellAreaApplyAttributesId)))

	detail := signalCellAreaApplyAttributesDetail{callback, handlerID}
	signalCellAreaApplyAttributesMap[signalCellAreaApplyAttributesId] = detail

	return signalCellAreaApplyAttributesId
}

/*
DisconnectApplyAttributes disconnects a callback from the 'apply-attributes' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectApplyAttributes.
*/
func (recv *CellArea) DisconnectApplyAttributes(connectionID int) {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	detail, exists := signalCellAreaApplyAttributesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaApplyAttributesMap, connectionID)
}

//export cellarea_applyAttributesHandler
func cellarea_applyAttributesHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_is_expander C.gboolean, c_is_expanded C.gboolean, data C.gpointer) {
	signalCellAreaApplyAttributesLock.RLock()
	defer signalCellAreaApplyAttributesLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	isExpander := c_is_expander == C.TRUE

	isExpanded := c_is_expanded == C.TRUE

	index := int(uintptr(data))
	callback := signalCellAreaApplyAttributesMap[index].callback
	callback(model, iter, isExpander, isExpanded)
}

type signalCellAreaFocusChangedDetail struct {
	callback  CellAreaSignalFocusChangedCallback
	handlerID C.gulong
}

var signalCellAreaFocusChangedId int
var signalCellAreaFocusChangedMap = make(map[int]signalCellAreaFocusChangedDetail)
var signalCellAreaFocusChangedLock sync.RWMutex

// CellAreaSignalFocusChangedCallback is a callback function for a 'focus-changed' signal emitted from a CellArea.
type CellAreaSignalFocusChangedCallback func(renderer *CellRenderer, path string)

/*
ConnectFocusChanged connects the callback to the 'focus-changed' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectFocusChanged to remove it.
*/
func (recv *CellArea) ConnectFocusChanged(callback CellAreaSignalFocusChangedCallback) int {
	signalCellAreaFocusChangedLock.Lock()
	defer signalCellAreaFocusChangedLock.Unlock()

	signalCellAreaFocusChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_focus_changed(instance, C.gpointer(uintptr(signalCellAreaFocusChangedId)))

	detail := signalCellAreaFocusChangedDetail{callback, handlerID}
	signalCellAreaFocusChangedMap[signalCellAreaFocusChangedId] = detail

	return signalCellAreaFocusChangedId
}

/*
DisconnectFocusChanged disconnects a callback from the 'focus-changed' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectFocusChanged.
*/
func (recv *CellArea) DisconnectFocusChanged(connectionID int) {
	signalCellAreaFocusChangedLock.Lock()
	defer signalCellAreaFocusChangedLock.Unlock()

	detail, exists := signalCellAreaFocusChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaFocusChangedMap, connectionID)
}

//export cellarea_focusChangedHandler
func cellarea_focusChangedHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_path *C.gchar, data C.gpointer) {
	signalCellAreaFocusChangedLock.RLock()
	defer signalCellAreaFocusChangedLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellAreaFocusChangedMap[index].callback
	callback(renderer, path)
}

type signalCellAreaRemoveEditableDetail struct {
	callback  CellAreaSignalRemoveEditableCallback
	handlerID C.gulong
}

var signalCellAreaRemoveEditableId int
var signalCellAreaRemoveEditableMap = make(map[int]signalCellAreaRemoveEditableDetail)
var signalCellAreaRemoveEditableLock sync.RWMutex

// CellAreaSignalRemoveEditableCallback is a callback function for a 'remove-editable' signal emitted from a CellArea.
type CellAreaSignalRemoveEditableCallback func(renderer *CellRenderer, editable *CellEditable)

/*
ConnectRemoveEditable connects the callback to the 'remove-editable' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectRemoveEditable to remove it.
*/
func (recv *CellArea) ConnectRemoveEditable(callback CellAreaSignalRemoveEditableCallback) int {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	signalCellAreaRemoveEditableId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_remove_editable(instance, C.gpointer(uintptr(signalCellAreaRemoveEditableId)))

	detail := signalCellAreaRemoveEditableDetail{callback, handlerID}
	signalCellAreaRemoveEditableMap[signalCellAreaRemoveEditableId] = detail

	return signalCellAreaRemoveEditableId
}

/*
DisconnectRemoveEditable disconnects a callback from the 'remove-editable' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectRemoveEditable.
*/
func (recv *CellArea) DisconnectRemoveEditable(connectionID int) {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	detail, exists := signalCellAreaRemoveEditableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaRemoveEditableMap, connectionID)
}

//export cellarea_removeEditableHandler
func cellarea_removeEditableHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_editable *C.GtkCellEditable, data C.gpointer) {
	signalCellAreaRemoveEditableLock.RLock()
	defer signalCellAreaRemoveEditableLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	index := int(uintptr(data))
	callback := signalCellAreaRemoveEditableMap[index].callback
	callback(renderer, editable)
}

// Blacklisted : gtk_cell_area_activate

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_cell_area_add

// Blacklisted : gtk_cell_area_add_focus_sibling

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Blacklisted : gtk_cell_area_apply_attributes

// Blacklisted : gtk_cell_area_attribute_connect

// Blacklisted : gtk_cell_area_attribute_disconnect

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// Blacklisted : gtk_cell_area_cell_get_property

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// Blacklisted : gtk_cell_area_cell_set_property

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : gtk_cell_area_copy_context

// Blacklisted : gtk_cell_area_create_context

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_cell_area_focus

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback (GtkCellCallback) for param callback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter callback : no type generator for CellAllocCallback (GtkCellAllocCallback) for param callback

// Blacklisted : gtk_cell_area_get_cell_allocation

// Blacklisted : gtk_cell_area_get_cell_at_position

// Blacklisted : gtk_cell_area_get_current_path_string

// Blacklisted : gtk_cell_area_get_edit_widget

// Blacklisted : gtk_cell_area_get_edited_cell

// Blacklisted : gtk_cell_area_get_focus_cell

// Blacklisted : gtk_cell_area_get_focus_from_sibling

// Blacklisted : gtk_cell_area_get_focus_siblings

// Blacklisted : gtk_cell_area_get_preferred_height

// Blacklisted : gtk_cell_area_get_preferred_height_for_width

// Blacklisted : gtk_cell_area_get_preferred_width

// Blacklisted : gtk_cell_area_get_preferred_width_for_height

// Blacklisted : gtk_cell_area_get_request_mode

// Blacklisted : gtk_cell_area_has_renderer

// Blacklisted : gtk_cell_area_inner_cell_area

// Blacklisted : gtk_cell_area_is_activatable

// Blacklisted : gtk_cell_area_is_focus_sibling

// Blacklisted : gtk_cell_area_remove

// Blacklisted : gtk_cell_area_remove_focus_sibling

// Blacklisted : gtk_cell_area_render

// Blacklisted : gtk_cell_area_request_renderer

// Blacklisted : gtk_cell_area_set_focus_cell

// Blacklisted : gtk_cell_area_stop_editing

// Blacklisted : gtk_cell_area_box_new

// Blacklisted : gtk_cell_area_box_get_spacing

// Blacklisted : gtk_cell_area_box_pack_end

// Blacklisted : gtk_cell_area_box_pack_start

// Blacklisted : gtk_cell_area_box_set_spacing

// Blacklisted : gtk_cell_area_context_get_allocation

// Blacklisted : gtk_cell_area_context_get_area

// Blacklisted : gtk_cell_area_context_get_preferred_height

// Blacklisted : gtk_cell_area_context_get_preferred_height_for_width

// Blacklisted : gtk_cell_area_context_get_preferred_width

// Blacklisted : gtk_cell_area_context_get_preferred_width_for_height

// Blacklisted : gtk_cell_area_context_push_preferred_height

// Blacklisted : gtk_cell_area_context_push_preferred_width

// Blacklisted : gtk_cell_renderer_get_aligned_area

// Blacklisted : gtk_cell_renderer_get_preferred_height

// Blacklisted : gtk_cell_renderer_get_preferred_height_for_width

// Blacklisted : gtk_cell_renderer_get_preferred_size

// Blacklisted : gtk_cell_renderer_get_preferred_width

// Blacklisted : gtk_cell_renderer_get_preferred_width_for_height

// Blacklisted : gtk_cell_renderer_get_request_mode

// Blacklisted : gtk_cell_renderer_get_state

// Blacklisted : gtk_cell_renderer_is_activatable

// Blacklisted : gtk_cell_view_get_draw_sensitive

// Blacklisted : gtk_cell_view_get_fit_model

// Blacklisted : gtk_cell_view_set_background_rgba

// Blacklisted : gtk_cell_view_set_draw_sensitive

// Blacklisted : gtk_cell_view_set_fit_model

// Blacklisted : gtk_color_button_new_with_rgba

// Blacklisted : gtk_color_button_get_rgba

// Blacklisted : gtk_color_button_set_rgba

// Blacklisted : gtk_color_selection_get_current_rgba

// Blacklisted : gtk_color_selection_get_previous_rgba

// Blacklisted : gtk_color_selection_set_current_rgba

// Blacklisted : gtk_color_selection_set_previous_rgba

// Blacklisted : gtk_combo_box_get_active_id

// Blacklisted : gtk_combo_box_get_id_column

// Blacklisted : gtk_combo_box_get_popup_fixed_width

// Blacklisted : gtk_combo_box_popup_for_device

// Blacklisted : gtk_combo_box_set_active_id

// Blacklisted : gtk_combo_box_set_id_column

// Blacklisted : gtk_combo_box_set_popup_fixed_width

// Blacklisted : gtk_combo_box_text_insert

// Blacklisted : gtk_combo_box_text_remove_all

// Blacklisted : gtk_entry_get_icon_area

// Blacklisted : gtk_entry_get_text_area

// Blacklisted : gtk_entry_completion_new_with_area

// Blacklisted : gtk_icon_info_load_symbolic

// Blacklisted : gtk_icon_info_load_symbolic_for_context

// Blacklisted : gtk_icon_info_load_symbolic_for_style

// Blacklisted : gtk_icon_view_new_with_area

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc (GtkMenuPositionFunc) for param func

// Blacklisted : gtk_menu_item_get_reserve_indicator

// Blacklisted : gtk_menu_item_set_reserve_indicator

// Blacklisted : gtk_menu_shell_get_parent_shell

// Blacklisted : gtk_menu_shell_get_selected_item

// Blacklisted : gtk_numerable_icon_new

// Blacklisted : gtk_numerable_icon_new_with_style_context

// Blacklisted : gtk_numerable_icon_get_background_gicon

// Blacklisted : gtk_numerable_icon_get_background_icon_name

// Blacklisted : gtk_numerable_icon_get_count

// Blacklisted : gtk_numerable_icon_get_label

// Blacklisted : gtk_numerable_icon_get_style_context

// Blacklisted : gtk_numerable_icon_set_background_gicon

// Blacklisted : gtk_numerable_icon_set_background_icon_name

// Blacklisted : gtk_numerable_icon_set_count

// Blacklisted : gtk_numerable_icon_set_label

// Blacklisted : gtk_numerable_icon_set_style_context

// Blacklisted : gtk_paned_new

// Blacklisted : gtk_progress_bar_get_show_text

// Blacklisted : gtk_progress_bar_set_show_text

// Blacklisted : gtk_radio_action_join_group

// Blacklisted : gtk_radio_button_join_group

// Blacklisted : gtk_scale_new

// Blacklisted : gtk_scale_new_with_range

// Blacklisted : gtk_scrollbar_new

// Blacklisted : gtk_scrolled_window_get_min_content_height

// Blacklisted : gtk_scrolled_window_get_min_content_width

// Blacklisted : gtk_scrolled_window_set_min_content_height

// Blacklisted : gtk_scrolled_window_set_min_content_width

// Blacklisted : gtk_separator_new

// Blacklisted : gtk_style_has_context

type signalStyleContextChangedDetail struct {
	callback  StyleContextSignalChangedCallback
	handlerID C.gulong
}

var signalStyleContextChangedId int
var signalStyleContextChangedMap = make(map[int]signalStyleContextChangedDetail)
var signalStyleContextChangedLock sync.RWMutex

// StyleContextSignalChangedCallback is a callback function for a 'changed' signal emitted from a StyleContext.
type StyleContextSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the StyleContext.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *StyleContext) ConnectChanged(callback StyleContextSignalChangedCallback) int {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	signalStyleContextChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.StyleContext_signal_connect_changed(instance, C.gpointer(uintptr(signalStyleContextChangedId)))

	detail := signalStyleContextChangedDetail{callback, handlerID}
	signalStyleContextChangedMap[signalStyleContextChangedId] = detail

	return signalStyleContextChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the StyleContext.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *StyleContext) DisconnectChanged(connectionID int) {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	detail, exists := signalStyleContextChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleContextChangedMap, connectionID)
}

//export stylecontext_changedHandler
func stylecontext_changedHandler(_ *C.GObject, data C.gpointer) {
	signalStyleContextChangedLock.RLock()
	defer signalStyleContextChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleContextChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_style_context_add_provider_for_screen

// Blacklisted : gtk_style_context_remove_provider_for_screen

// Blacklisted : gtk_style_context_reset_widgets

// Blacklisted : gtk_style_context_add_class

// Blacklisted : gtk_style_context_add_provider

// Blacklisted : gtk_style_context_add_region

// Unsupported : gtk_style_context_cancel_animations : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// Blacklisted : gtk_style_context_get_background_color

// Blacklisted : gtk_style_context_get_border

// Blacklisted : gtk_style_context_get_border_color

// Blacklisted : gtk_style_context_get_color

// Blacklisted : gtk_style_context_get_direction

// Blacklisted : gtk_style_context_get_font

// Blacklisted : gtk_style_context_get_junction_sides

// Blacklisted : gtk_style_context_get_margin

// Blacklisted : gtk_style_context_get_padding

// Blacklisted : gtk_style_context_get_path

// Blacklisted : gtk_style_context_get_property

// Blacklisted : gtk_style_context_get_state

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : gtk_style_context_has_class

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Blacklisted : gtk_style_context_invalidate

// Blacklisted : gtk_style_context_list_classes

// Blacklisted : gtk_style_context_list_regions

// Unsupported : gtk_style_context_notify_state_change : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// Blacklisted : gtk_style_context_pop_animatable_region

// Unsupported : gtk_style_context_push_animatable_region : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// Blacklisted : gtk_style_context_remove_class

// Blacklisted : gtk_style_context_remove_provider

// Blacklisted : gtk_style_context_remove_region

// Blacklisted : gtk_style_context_restore

// Blacklisted : gtk_style_context_save

// Blacklisted : gtk_style_context_scroll_animations

// Blacklisted : gtk_style_context_set_background

// Blacklisted : gtk_style_context_set_direction

// Blacklisted : gtk_style_context_set_junction_sides

// Blacklisted : gtk_style_context_set_path

// Blacklisted : gtk_style_context_set_screen

// Blacklisted : gtk_style_context_set_state

// Blacklisted : gtk_style_context_state_is_running

// gtk_style_properties_lookup_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser*) for param parse_func
// gtk_style_properties_register_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser) for param parse_func
// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// Blacklisted : gtk_style_properties_get_property

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : gtk_style_properties_lookup_color

// Blacklisted : gtk_style_properties_map_color

// Blacklisted : gtk_style_properties_merge

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// Blacklisted : gtk_style_properties_set_property

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : gtk_style_properties_unset_property

// Blacklisted : gtk_switch_new

// Blacklisted : gtk_switch_get_active

// Blacklisted : gtk_switch_set_active

// Blacklisted : gtk_text_view_get_cursor_locations

// gtk_theming_engine_register_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser) for param parse_func
// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// Blacklisted : gtk_theming_engine_get_background_color

// Blacklisted : gtk_theming_engine_get_border

// Blacklisted : gtk_theming_engine_get_border_color

// Blacklisted : gtk_theming_engine_get_color

// Blacklisted : gtk_theming_engine_get_direction

// Blacklisted : gtk_theming_engine_get_font

// Blacklisted : gtk_theming_engine_get_junction_sides

// Blacklisted : gtk_theming_engine_get_margin

// Blacklisted : gtk_theming_engine_get_padding

// Blacklisted : gtk_theming_engine_get_path

// Blacklisted : gtk_theming_engine_get_property

// Blacklisted : gtk_theming_engine_get_state

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// Blacklisted : gtk_theming_engine_get_style_property

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : gtk_theming_engine_has_class

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Blacklisted : gtk_theming_engine_lookup_color

// Blacklisted : gtk_theming_engine_state_is_running

// Blacklisted : gtk_tree_view_is_blank_at_pos

// Blacklisted : gtk_tree_view_column_new_with_area

// Blacklisted : gtk_tree_view_column_get_button

type signalWidgetDrawDetail struct {
	callback  WidgetSignalDrawCallback
	handlerID C.gulong
}

var signalWidgetDrawId int
var signalWidgetDrawMap = make(map[int]signalWidgetDrawDetail)
var signalWidgetDrawLock sync.RWMutex

// WidgetSignalDrawCallback is a callback function for a 'draw' signal emitted from a Widget.
type WidgetSignalDrawCallback func(cr *cairo.Context) bool

/*
ConnectDraw connects the callback to the 'draw' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDraw to remove it.
*/
func (recv *Widget) ConnectDraw(callback WidgetSignalDrawCallback) int {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	signalWidgetDrawId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_draw(instance, C.gpointer(uintptr(signalWidgetDrawId)))

	detail := signalWidgetDrawDetail{callback, handlerID}
	signalWidgetDrawMap[signalWidgetDrawId] = detail

	return signalWidgetDrawId
}

/*
DisconnectDraw disconnects a callback from the 'draw' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDraw.
*/
func (recv *Widget) DisconnectDraw(connectionID int) {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	detail, exists := signalWidgetDrawMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDrawMap, connectionID)
}

//export widget_drawHandler
func widget_drawHandler(_ *C.GObject, c_cr *C.cairo_t, data C.gpointer) C.gboolean {
	signalWidgetDrawLock.RLock()
	defer signalWidgetDrawLock.RUnlock()

	cr := cairo.ContextNewFromC(unsafe.Pointer(c_cr))

	index := int(uintptr(data))
	callback := signalWidgetDrawMap[index].callback
	retGo := callback(cr)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetStateFlagsChangedDetail struct {
	callback  WidgetSignalStateFlagsChangedCallback
	handlerID C.gulong
}

var signalWidgetStateFlagsChangedId int
var signalWidgetStateFlagsChangedMap = make(map[int]signalWidgetStateFlagsChangedDetail)
var signalWidgetStateFlagsChangedLock sync.RWMutex

// WidgetSignalStateFlagsChangedCallback is a callback function for a 'state-flags-changed' signal emitted from a Widget.
type WidgetSignalStateFlagsChangedCallback func(flags StateFlags)

/*
ConnectStateFlagsChanged connects the callback to the 'state-flags-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStateFlagsChanged to remove it.
*/
func (recv *Widget) ConnectStateFlagsChanged(callback WidgetSignalStateFlagsChangedCallback) int {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	signalWidgetStateFlagsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_state_flags_changed(instance, C.gpointer(uintptr(signalWidgetStateFlagsChangedId)))

	detail := signalWidgetStateFlagsChangedDetail{callback, handlerID}
	signalWidgetStateFlagsChangedMap[signalWidgetStateFlagsChangedId] = detail

	return signalWidgetStateFlagsChangedId
}

/*
DisconnectStateFlagsChanged disconnects a callback from the 'state-flags-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStateFlagsChanged.
*/
func (recv *Widget) DisconnectStateFlagsChanged(connectionID int) {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	detail, exists := signalWidgetStateFlagsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStateFlagsChangedMap, connectionID)
}

//export widget_stateFlagsChangedHandler
func widget_stateFlagsChangedHandler(_ *C.GObject, c_flags C.GtkStateFlags, data C.gpointer) {
	signalWidgetStateFlagsChangedLock.RLock()
	defer signalWidgetStateFlagsChangedLock.RUnlock()

	flags := StateFlags(c_flags)

	index := int(uintptr(data))
	callback := signalWidgetStateFlagsChangedMap[index].callback
	callback(flags)
}

type signalWidgetStyleUpdatedDetail struct {
	callback  WidgetSignalStyleUpdatedCallback
	handlerID C.gulong
}

var signalWidgetStyleUpdatedId int
var signalWidgetStyleUpdatedMap = make(map[int]signalWidgetStyleUpdatedDetail)
var signalWidgetStyleUpdatedLock sync.RWMutex

// WidgetSignalStyleUpdatedCallback is a callback function for a 'style-updated' signal emitted from a Widget.
type WidgetSignalStyleUpdatedCallback func()

/*
ConnectStyleUpdated connects the callback to the 'style-updated' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStyleUpdated to remove it.
*/
func (recv *Widget) ConnectStyleUpdated(callback WidgetSignalStyleUpdatedCallback) int {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	signalWidgetStyleUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_style_updated(instance, C.gpointer(uintptr(signalWidgetStyleUpdatedId)))

	detail := signalWidgetStyleUpdatedDetail{callback, handlerID}
	signalWidgetStyleUpdatedMap[signalWidgetStyleUpdatedId] = detail

	return signalWidgetStyleUpdatedId
}

/*
DisconnectStyleUpdated disconnects a callback from the 'style-updated' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStyleUpdated.
*/
func (recv *Widget) DisconnectStyleUpdated(connectionID int) {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	detail, exists := signalWidgetStyleUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStyleUpdatedMap, connectionID)
}

//export widget_styleUpdatedHandler
func widget_styleUpdatedHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetStyleUpdatedLock.RLock()
	defer signalWidgetStyleUpdatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetStyleUpdatedMap[index].callback
	callback()
}

// Blacklisted : gtk_widget_add_device_events

// Blacklisted : gtk_widget_device_is_shadowed

// Blacklisted : gtk_widget_draw

// Blacklisted : gtk_widget_get_device_enabled

// Blacklisted : gtk_widget_get_device_events

// Blacklisted : gtk_widget_get_margin_bottom

// Blacklisted : gtk_widget_get_margin_left

// Blacklisted : gtk_widget_get_margin_right

// Blacklisted : gtk_widget_get_margin_top

// Blacklisted : gtk_widget_get_preferred_height

// Blacklisted : gtk_widget_get_preferred_height_for_width

// Blacklisted : gtk_widget_get_preferred_size

// Blacklisted : gtk_widget_get_preferred_width

// Blacklisted : gtk_widget_get_preferred_width_for_height

// Blacklisted : gtk_widget_get_request_mode

// Blacklisted : gtk_widget_get_state_flags

// Blacklisted : gtk_widget_input_shape_combine_region

// Blacklisted : gtk_widget_override_background_color

// Blacklisted : gtk_widget_override_color

// Blacklisted : gtk_widget_override_cursor

// Blacklisted : gtk_widget_override_font

// Blacklisted : gtk_widget_override_symbolic_color

// Blacklisted : gtk_widget_queue_draw_region

// Blacklisted : gtk_widget_render_icon_pixbuf

// Blacklisted : gtk_widget_reset_style

// Blacklisted : gtk_widget_set_device_enabled

// Blacklisted : gtk_widget_set_device_events

// Blacklisted : gtk_widget_set_margin_bottom

// Blacklisted : gtk_widget_set_margin_left

// Blacklisted : gtk_widget_set_margin_right

// Blacklisted : gtk_widget_set_margin_top

// Blacklisted : gtk_widget_set_state_flags

// Blacklisted : gtk_widget_set_support_multidevice

// Blacklisted : gtk_widget_shape_combine_region

// Blacklisted : gtk_widget_unset_state_flags

// Blacklisted : gtk_window_get_application

// Blacklisted : gtk_window_get_has_resize_grip

// Blacklisted : gtk_window_get_resize_grip_area

// Blacklisted : gtk_window_resize_grip_is_visible

// Blacklisted : gtk_window_resize_to_geometry

// Blacklisted : gtk_window_set_application

// Blacklisted : gtk_window_set_default_geometry

// Blacklisted : gtk_window_set_has_resize_grip

// Blacklisted : gtk_window_set_has_user_ref_count

// Blacklisted : gtk_window_group_get_current_device_grab
