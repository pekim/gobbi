// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
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

	void SearchEntry_nextMatchHandler();

	static gulong SearchEntry_signal_connect_next_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "next-match", SearchEntry_nextMatchHandler, data);
	}

*/
/*

	void SearchEntry_previousMatchHandler();

	static gulong SearchEntry_signal_connect_previous_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "previous-match", SearchEntry_previousMatchHandler, data);
	}

*/
/*

	void SearchEntry_stopSearchHandler();

	static gulong SearchEntry_signal_connect_stop_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stop-search", SearchEntry_stopSearchHandler, data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// LoadFromResource is a wrapper around the C function gtk_css_provider_load_from_resource.
func (recv *CssProvider) LoadFromResource(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.gtk_css_provider_load_from_resource((*C.GtkCssProvider)(recv.native), c_resource_path)

	return
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GrabFocusWithoutSelecting is a wrapper around the C function gtk_entry_grab_focus_without_selecting.
func (recv *Entry) GrabFocusWithoutSelecting() {
	C.gtk_entry_grab_focus_without_selecting((*C.GtkEntry)(recv.native))

	return
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GLArea is a wrapper around the C record GtkGLArea.
type GLArea struct {
	native *C.GtkGLArea
	// Private : parent_instance
}

func GLAreaNewFromC(u unsafe.Pointer) *GLArea {
	c := (*C.GtkGLArea)(u)
	if c == nil {
		return nil
	}

	g := &GLArea{native: c}

	return g
}

func (recv *GLArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *GLArea) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GLArea) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *GLArea) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to GLArea.
// Exercise care, as this is a potentially dangerous function if the Object is not a GLArea.
func CastToGLArea(object *gobject.Object) *GLArea {
	return GLAreaNewFromC(object.ToC())
}

// GLAreaNew is a wrapper around the C function gtk_gl_area_new.
func GLAreaNew() *GLArea {
	retC := C.gtk_gl_area_new()
	retGo := GLAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttachBuffers is a wrapper around the C function gtk_gl_area_attach_buffers.
func (recv *GLArea) AttachBuffers() {
	C.gtk_gl_area_attach_buffers((*C.GtkGLArea)(recv.native))

	return
}

// GetAutoRender is a wrapper around the C function gtk_gl_area_get_auto_render.
func (recv *GLArea) GetAutoRender() bool {
	retC := C.gtk_gl_area_get_auto_render((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function gtk_gl_area_get_context.
func (recv *GLArea) GetContext() *gdk.GLContext {
	retC := C.gtk_gl_area_get_context((*C.GtkGLArea)(recv.native))
	retGo := gdk.GLContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetError is a wrapper around the C function gtk_gl_area_get_error.
func (recv *GLArea) GetError() *glib.Error {
	retC := C.gtk_gl_area_get_error((*C.GtkGLArea)(recv.native))
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHasAlpha is a wrapper around the C function gtk_gl_area_get_has_alpha.
func (recv *GLArea) GetHasAlpha() bool {
	retC := C.gtk_gl_area_get_has_alpha((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasDepthBuffer is a wrapper around the C function gtk_gl_area_get_has_depth_buffer.
func (recv *GLArea) GetHasDepthBuffer() bool {
	retC := C.gtk_gl_area_get_has_depth_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasStencilBuffer is a wrapper around the C function gtk_gl_area_get_has_stencil_buffer.
func (recv *GLArea) GetHasStencilBuffer() bool {
	retC := C.gtk_gl_area_get_has_stencil_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRequiredVersion is a wrapper around the C function gtk_gl_area_get_required_version.
func (recv *GLArea) GetRequiredVersion() (int32, int32) {
	var c_major C.gint

	var c_minor C.gint

	C.gtk_gl_area_get_required_version((*C.GtkGLArea)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// MakeCurrent is a wrapper around the C function gtk_gl_area_make_current.
func (recv *GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current((*C.GtkGLArea)(recv.native))

	return
}

// QueueRender is a wrapper around the C function gtk_gl_area_queue_render.
func (recv *GLArea) QueueRender() {
	C.gtk_gl_area_queue_render((*C.GtkGLArea)(recv.native))

	return
}

// SetAutoRender is a wrapper around the C function gtk_gl_area_set_auto_render.
func (recv *GLArea) SetAutoRender(autoRender bool) {
	c_auto_render :=
		boolToGboolean(autoRender)

	C.gtk_gl_area_set_auto_render((*C.GtkGLArea)(recv.native), c_auto_render)

	return
}

// SetError is a wrapper around the C function gtk_gl_area_set_error.
func (recv *GLArea) SetError(error *glib.Error) {
	c_error := (*C.GError)(error.ToC())

	C.gtk_gl_area_set_error((*C.GtkGLArea)(recv.native), c_error)

	return
}

// SetHasAlpha is a wrapper around the C function gtk_gl_area_set_has_alpha.
func (recv *GLArea) SetHasAlpha(hasAlpha bool) {
	c_has_alpha :=
		boolToGboolean(hasAlpha)

	C.gtk_gl_area_set_has_alpha((*C.GtkGLArea)(recv.native), c_has_alpha)

	return
}

// SetHasDepthBuffer is a wrapper around the C function gtk_gl_area_set_has_depth_buffer.
func (recv *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	c_has_depth_buffer :=
		boolToGboolean(hasDepthBuffer)

	C.gtk_gl_area_set_has_depth_buffer((*C.GtkGLArea)(recv.native), c_has_depth_buffer)

	return
}

// SetHasStencilBuffer is a wrapper around the C function gtk_gl_area_set_has_stencil_buffer.
func (recv *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	c_has_stencil_buffer :=
		boolToGboolean(hasStencilBuffer)

	C.gtk_gl_area_set_has_stencil_buffer((*C.GtkGLArea)(recv.native), c_has_stencil_buffer)

	return
}

// SetRequiredVersion is a wrapper around the C function gtk_gl_area_set_required_version.
func (recv *GLArea) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.gint)(major)

	c_minor := (C.gint)(minor)

	C.gtk_gl_area_set_required_version((*C.GtkGLArea)(recv.native), c_major, c_minor)

	return
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// GetXalign is a wrapper around the C function gtk_label_get_xalign.
func (recv *Label) GetXalign() float32 {
	retC := C.gtk_label_get_xalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetYalign is a wrapper around the C function gtk_label_get_yalign.
func (recv *Label) GetYalign() float32 {
	retC := C.gtk_label_get_yalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// SetXalign is a wrapper around the C function gtk_label_set_xalign.
func (recv *Label) SetXalign(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_label_set_xalign((*C.GtkLabel)(recv.native), c_xalign)

	return
}

// SetYalign is a wrapper around the C function gtk_label_set_yalign.
func (recv *Label) SetYalign(yalign float32) {
	c_yalign := (C.gfloat)(yalign)

	C.gtk_label_set_yalign((*C.GtkLabel)(recv.native), c_yalign)

	return
}

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// ModelButtonNew is a wrapper around the C function gtk_model_button_new.
func ModelButtonNew() *ModelButton {
	retC := C.gtk_model_button_new()
	retGo := ModelButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DetachTab is a wrapper around the C function gtk_notebook_detach_tab.
func (recv *Notebook) DetachTab(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_notebook_detach_tab((*C.GtkNotebook)(recv.native), c_child)

	return
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GetWideHandle is a wrapper around the C function gtk_paned_get_wide_handle.
func (recv *Paned) GetWideHandle() bool {
	retC := C.gtk_paned_get_wide_handle((*C.GtkPaned)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetWideHandle is a wrapper around the C function gtk_paned_set_wide_handle.
func (recv *Paned) SetWideHandle(wide bool) {
	c_wide :=
		boolToGboolean(wide)

	C.gtk_paned_set_wide_handle((*C.GtkPaned)(recv.native), c_wide)

	return
}

// GetTransitionsEnabled is a wrapper around the C function gtk_popover_get_transitions_enabled.
func (recv *Popover) GetTransitionsEnabled() bool {
	retC := C.gtk_popover_get_transitions_enabled((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTransitionsEnabled is a wrapper around the C function gtk_popover_set_transitions_enabled.
func (recv *Popover) SetTransitionsEnabled(transitionsEnabled bool) {
	c_transitions_enabled :=
		boolToGboolean(transitionsEnabled)

	C.gtk_popover_set_transitions_enabled((*C.GtkPopover)(recv.native), c_transitions_enabled)

	return
}

// PopoverMenuNew is a wrapper around the C function gtk_popover_menu_new.
func PopoverMenuNew() *PopoverMenu {
	retC := C.gtk_popover_menu_new()
	retGo := PopoverMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OpenSubmenu is a wrapper around the C function gtk_popover_menu_open_submenu.
func (recv *PopoverMenu) OpenSubmenu(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_popover_menu_open_submenu((*C.GtkPopoverMenu)(recv.native), c_name)

	return
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'edge-overshot' for ScrolledWindow : unsupported parameter pos : type PositionType :

// Unsupported signal 'edge-reached' for ScrolledWindow : unsupported parameter pos : type PositionType :

// GetOverlayScrolling is a wrapper around the C function gtk_scrolled_window_get_overlay_scrolling.
func (recv *ScrolledWindow) GetOverlayScrolling() bool {
	retC := C.gtk_scrolled_window_get_overlay_scrolling((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetOverlayScrolling is a wrapper around the C function gtk_scrolled_window_set_overlay_scrolling.
func (recv *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	c_overlay_scrolling :=
		boolToGboolean(overlayScrolling)

	C.gtk_scrolled_window_set_overlay_scrolling((*C.GtkScrolledWindow)(recv.native), c_overlay_scrolling)

	return
}

var signalSearchEntryNextMatchId int
var signalSearchEntryNextMatchMap = make(map[int]SearchEntrySignalNextMatchCallback)
var signalSearchEntryNextMatchLock sync.Mutex

// SearchEntrySignalNextMatchCallback is a callback function for a 'next-match' signal emitted from a SearchEntry.
type SearchEntrySignalNextMatchCallback func()

/*
ConnectNextMatch connects the callback to the 'next-match' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectNextMatch to remove it.
*/
func (recv *SearchEntry) ConnectNextMatch(callback SearchEntrySignalNextMatchCallback) int {
	signalSearchEntryNextMatchLock.Lock()
	defer signalSearchEntryNextMatchLock.Unlock()

	signalSearchEntryNextMatchId++
	signalSearchEntryNextMatchMap[signalSearchEntryNextMatchId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.SearchEntry_signal_connect_next_match(instance, C.gpointer(uintptr(signalSearchEntryNextMatchId)))
	return int(retC)
}

/*
DisconnectNextMatch disconnects a callback from the 'next-match' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectNextMatch.
*/
func (recv *SearchEntry) DisconnectNextMatch(connectionID int) {
	signalSearchEntryNextMatchLock.Lock()
	defer signalSearchEntryNextMatchLock.Unlock()

	_, exists := signalSearchEntryNextMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalSearchEntryNextMatchMap, connectionID)
}

//export SearchEntry_nextMatchHandler
func SearchEntry_nextMatchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryNextMatchMap[index]
	callback()
}

var signalSearchEntryPreviousMatchId int
var signalSearchEntryPreviousMatchMap = make(map[int]SearchEntrySignalPreviousMatchCallback)
var signalSearchEntryPreviousMatchLock sync.Mutex

// SearchEntrySignalPreviousMatchCallback is a callback function for a 'previous-match' signal emitted from a SearchEntry.
type SearchEntrySignalPreviousMatchCallback func()

/*
ConnectPreviousMatch connects the callback to the 'previous-match' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectPreviousMatch to remove it.
*/
func (recv *SearchEntry) ConnectPreviousMatch(callback SearchEntrySignalPreviousMatchCallback) int {
	signalSearchEntryPreviousMatchLock.Lock()
	defer signalSearchEntryPreviousMatchLock.Unlock()

	signalSearchEntryPreviousMatchId++
	signalSearchEntryPreviousMatchMap[signalSearchEntryPreviousMatchId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.SearchEntry_signal_connect_previous_match(instance, C.gpointer(uintptr(signalSearchEntryPreviousMatchId)))
	return int(retC)
}

/*
DisconnectPreviousMatch disconnects a callback from the 'previous-match' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectPreviousMatch.
*/
func (recv *SearchEntry) DisconnectPreviousMatch(connectionID int) {
	signalSearchEntryPreviousMatchLock.Lock()
	defer signalSearchEntryPreviousMatchLock.Unlock()

	_, exists := signalSearchEntryPreviousMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalSearchEntryPreviousMatchMap, connectionID)
}

//export SearchEntry_previousMatchHandler
func SearchEntry_previousMatchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryPreviousMatchMap[index]
	callback()
}

var signalSearchEntryStopSearchId int
var signalSearchEntryStopSearchMap = make(map[int]SearchEntrySignalStopSearchCallback)
var signalSearchEntryStopSearchLock sync.Mutex

// SearchEntrySignalStopSearchCallback is a callback function for a 'stop-search' signal emitted from a SearchEntry.
type SearchEntrySignalStopSearchCallback func()

/*
ConnectStopSearch connects the callback to the 'stop-search' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectStopSearch to remove it.
*/
func (recv *SearchEntry) ConnectStopSearch(callback SearchEntrySignalStopSearchCallback) int {
	signalSearchEntryStopSearchLock.Lock()
	defer signalSearchEntryStopSearchLock.Unlock()

	signalSearchEntryStopSearchId++
	signalSearchEntryStopSearchMap[signalSearchEntryStopSearchId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.SearchEntry_signal_connect_stop_search(instance, C.gpointer(uintptr(signalSearchEntryStopSearchId)))
	return int(retC)
}

/*
DisconnectStopSearch disconnects a callback from the 'stop-search' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectStopSearch.
*/
func (recv *SearchEntry) DisconnectStopSearch(connectionID int) {
	signalSearchEntryStopSearchLock.Lock()
	defer signalSearchEntryStopSearchLock.Unlock()

	_, exists := signalSearchEntryStopSearchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalSearchEntryStopSearchMap, connectionID)
}

//export SearchEntry_stopSearchHandler
func SearchEntry_stopSearchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryStopSearchMap[index]
	callback()
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetHhomogeneous is a wrapper around the C function gtk_stack_get_hhomogeneous.
func (recv *Stack) GetHhomogeneous() bool {
	retC := C.gtk_stack_get_hhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVhomogeneous is a wrapper around the C function gtk_stack_get_vhomogeneous.
func (recv *Stack) GetVhomogeneous() bool {
	retC := C.gtk_stack_get_vhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetHhomogeneous is a wrapper around the C function gtk_stack_set_hhomogeneous.
func (recv *Stack) SetHhomogeneous(hhomogeneous bool) {
	c_hhomogeneous :=
		boolToGboolean(hhomogeneous)

	C.gtk_stack_set_hhomogeneous((*C.GtkStack)(recv.native), c_hhomogeneous)

	return
}

// SetVhomogeneous is a wrapper around the C function gtk_stack_set_vhomogeneous.
func (recv *Stack) SetVhomogeneous(vhomogeneous bool) {
	c_vhomogeneous :=
		boolToGboolean(vhomogeneous)

	C.gtk_stack_set_vhomogeneous((*C.GtkStack)(recv.native), c_vhomogeneous)

	return
}

// StackSidebarNew is a wrapper around the C function gtk_stack_sidebar_new.
func StackSidebarNew() *StackSidebar {
	retC := C.gtk_stack_sidebar_new()
	retGo := StackSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_sidebar_get_stack.
func (recv *StackSidebar) GetStack() *Stack {
	retC := C.gtk_stack_sidebar_get_stack((*C.GtkStackSidebar)(recv.native))
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_sidebar_set_stack.
func (recv *StackSidebar) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(stack.ToC())

	C.gtk_stack_sidebar_set_stack((*C.GtkStackSidebar)(recv.native), c_stack)

	return
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// InsertMarkup is a wrapper around the C function gtk_text_buffer_insert_markup.
func (recv *TextBuffer) InsertMarkup(iter *TextIter, markup string, len int32) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_insert_markup((*C.GtkTextBuffer)(recv.native), c_iter, c_markup, c_len)

	return
}

// Unsupported signal 'extend-selection' for TextView : unsupported parameter granularity : type TextExtendSelection :

// GetMonospace is a wrapper around the C function gtk_text_view_get_monospace.
func (recv *TextView) GetMonospace() bool {
	retC := C.gtk_text_view_get_monospace((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetMonospace is a wrapper around the C function gtk_text_view_set_monospace.
func (recv *TextView) SetMonospace(monospace bool) {
	c_monospace :=
		boolToGboolean(monospace)

	C.gtk_text_view_set_monospace((*C.GtkTextView)(recv.native), c_monospace)

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_get_action_group : no return generator

// Unsupported : gtk_widget_list_action_prefixes : no return type

// GetTitlebar is a wrapper around the C function gtk_window_get_titlebar.
func (recv *Window) GetTitlebar() *Widget {
	retC := C.gtk_window_get_titlebar((*C.GtkWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}
