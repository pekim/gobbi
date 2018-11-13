// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
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

	void searchentry_nextMatchHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_next_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "next-match", G_CALLBACK(searchentry_nextMatchHandler), data);
	}

*/
/*

	void searchentry_previousMatchHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_previous_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "previous-match", G_CALLBACK(searchentry_previousMatchHandler), data);
	}

*/
/*

	void searchentry_stopSearchHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_stop_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stop-search", G_CALLBACK(searchentry_stopSearchHandler), data);
	}

*/
import "C"

// Loads the data contained in the resource at @resource_path into
// the #GtkCssProvider, clearing any previously loaded information.
//
// To track errors while loading CSS, connect to the
// #GtkCssProvider::parsing-error signal.
/*

C function : gtk_css_provider_load_from_resource
*/
func (recv *CssProvider) LoadFromResource(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.gtk_css_provider_load_from_resource((*C.GtkCssProvider)(recv.native), c_resource_path)

	return
}

// Causes @entry to have keyboard focus.
//
// It behaves like gtk_widget_grab_focus(),
// except that it doesn't select the contents of the entry.
// You only want to call this on some special entries
// which the user usually doesn't want to replace all text in,
// such as search-as-you-type entries.
/*

C function : gtk_entry_grab_focus_without_selecting
*/
func (recv *Entry) GrabFocusWithoutSelecting() {
	C.gtk_entry_grab_focus_without_selecting((*C.GtkEntry)(recv.native))

	return
}

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

// Creates a new #GtkGLArea widget.
/*

C function : gtk_gl_area_new
*/
func GLAreaNew() *GLArea {
	retC := C.gtk_gl_area_new()
	retGo := GLAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ensures that the @area framebuffer object is made the current draw
// and read target, and that all the required buffers for the @area
// are created and bound to the frambuffer.
//
// This function is automatically called before emitting the
// #GtkGLArea::render signal, and doesn't normally need to be called
// by application code.
/*

C function : gtk_gl_area_attach_buffers
*/
func (recv *GLArea) AttachBuffers() {
	C.gtk_gl_area_attach_buffers((*C.GtkGLArea)(recv.native))

	return
}

// Returns whether the area is in auto render mode or not.
/*

C function : gtk_gl_area_get_auto_render
*/
func (recv *GLArea) GetAutoRender() bool {
	retC := C.gtk_gl_area_get_auto_render((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the #GdkGLContext used by @area.
/*

C function : gtk_gl_area_get_context
*/
func (recv *GLArea) GetContext() *gdk.GLContext {
	retC := C.gtk_gl_area_get_context((*C.GtkGLArea)(recv.native))
	retGo := gdk.GLContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the current error set on the @area.
/*

C function : gtk_gl_area_get_error
*/
func (recv *GLArea) GetError() *glib.Error {
	retC := C.gtk_gl_area_get_error((*C.GtkGLArea)(recv.native))
	var retGo (*glib.Error)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ErrorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether the area has an alpha component.
/*

C function : gtk_gl_area_get_has_alpha
*/
func (recv *GLArea) GetHasAlpha() bool {
	retC := C.gtk_gl_area_get_has_alpha((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the area has a depth buffer.
/*

C function : gtk_gl_area_get_has_depth_buffer
*/
func (recv *GLArea) GetHasDepthBuffer() bool {
	retC := C.gtk_gl_area_get_has_depth_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the area has a stencil buffer.
/*

C function : gtk_gl_area_get_has_stencil_buffer
*/
func (recv *GLArea) GetHasStencilBuffer() bool {
	retC := C.gtk_gl_area_get_has_stencil_buffer((*C.GtkGLArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the required version of OpenGL set
// using gtk_gl_area_set_required_version().
/*

C function : gtk_gl_area_get_required_version
*/
func (recv *GLArea) GetRequiredVersion() (int32, int32) {
	var c_major C.gint

	var c_minor C.gint

	C.gtk_gl_area_get_required_version((*C.GtkGLArea)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// Ensures that the #GdkGLContext used by @area is associated with
// the #GtkGLArea.
//
// This function is automatically called before emitting the
// #GtkGLArea::render signal, and doesn't normally need to be called
// by application code.
/*

C function : gtk_gl_area_make_current
*/
func (recv *GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current((*C.GtkGLArea)(recv.native))

	return
}

// Marks the currently rendered data (if any) as invalid, and queues
// a redraw of the widget, ensuring that the #GtkGLArea::render signal
// is emitted during the draw.
//
// This is only needed when the gtk_gl_area_set_auto_render() has
// been called with a %FALSE value. The default behaviour is to
// emit #GtkGLArea::render on each draw.
/*

C function : gtk_gl_area_queue_render
*/
func (recv *GLArea) QueueRender() {
	C.gtk_gl_area_queue_render((*C.GtkGLArea)(recv.native))

	return
}

// If @auto_render is %TRUE the #GtkGLArea::render signal will be
// emitted every time the widget draws. This is the default and is
// useful if drawing the widget is faster.
//
// If @auto_render is %FALSE the data from previous rendering is kept
// around and will be used for drawing the widget the next time,
// unless the window is resized. In order to force a rendering
// gtk_gl_area_queue_render() must be called. This mode is useful when
// the scene changes seldomly, but takes a long time to redraw.
/*

C function : gtk_gl_area_set_auto_render
*/
func (recv *GLArea) SetAutoRender(autoRender bool) {
	c_auto_render :=
		boolToGboolean(autoRender)

	C.gtk_gl_area_set_auto_render((*C.GtkGLArea)(recv.native), c_auto_render)

	return
}

// Sets an error on the area which will be shown instead of the
// GL rendering. This is useful in the #GtkGLArea::create-context
// signal if GL context creation fails.
/*

C function : gtk_gl_area_set_error
*/
func (recv *GLArea) SetError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.gtk_gl_area_set_error((*C.GtkGLArea)(recv.native), c_error)

	return
}

// If @has_alpha is %TRUE the buffer allocated by the widget will have
// an alpha channel component, and when rendering to the window the
// result will be composited over whatever is below the widget.
//
// If @has_alpha is %FALSE there will be no alpha channel, and the
// buffer will fully replace anything below the widget.
/*

C function : gtk_gl_area_set_has_alpha
*/
func (recv *GLArea) SetHasAlpha(hasAlpha bool) {
	c_has_alpha :=
		boolToGboolean(hasAlpha)

	C.gtk_gl_area_set_has_alpha((*C.GtkGLArea)(recv.native), c_has_alpha)

	return
}

// If @has_depth_buffer is %TRUE the widget will allocate and
// enable a depth buffer for the target framebuffer. Otherwise
// there will be none.
/*

C function : gtk_gl_area_set_has_depth_buffer
*/
func (recv *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	c_has_depth_buffer :=
		boolToGboolean(hasDepthBuffer)

	C.gtk_gl_area_set_has_depth_buffer((*C.GtkGLArea)(recv.native), c_has_depth_buffer)

	return
}

// If @has_stencil_buffer is %TRUE the widget will allocate and
// enable a stencil buffer for the target framebuffer. Otherwise
// there will be none.
/*

C function : gtk_gl_area_set_has_stencil_buffer
*/
func (recv *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	c_has_stencil_buffer :=
		boolToGboolean(hasStencilBuffer)

	C.gtk_gl_area_set_has_stencil_buffer((*C.GtkGLArea)(recv.native), c_has_stencil_buffer)

	return
}

// Sets the required version of OpenGL to be used when creating the context
// for the widget.
//
// This function must be called before the area has been realized.
/*

C function : gtk_gl_area_set_required_version
*/
func (recv *GLArea) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.gint)(major)

	c_minor := (C.gint)(minor)

	C.gtk_gl_area_set_required_version((*C.GtkGLArea)(recv.native), c_major, c_minor)

	return
}

// Gets the #GtkLabel:xalign property for @label.
/*

C function : gtk_label_get_xalign
*/
func (recv *Label) GetXalign() float32 {
	retC := C.gtk_label_get_xalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Gets the #GtkLabel:yalign property for @label.
/*

C function : gtk_label_get_yalign
*/
func (recv *Label) GetYalign() float32 {
	retC := C.gtk_label_get_yalign((*C.GtkLabel)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Sets the #GtkLabel:xalign property for @label.
/*

C function : gtk_label_set_xalign
*/
func (recv *Label) SetXalign(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_label_set_xalign((*C.GtkLabel)(recv.native), c_xalign)

	return
}

// Sets the #GtkLabel:yalign property for @label.
/*

C function : gtk_label_set_yalign
*/
func (recv *Label) SetYalign(yalign float32) {
	c_yalign := (C.gfloat)(yalign)

	C.gtk_label_set_yalign((*C.GtkLabel)(recv.native), c_yalign)

	return
}

// Unsupported : gtk_list_box_bind_model : unsupported parameter create_widget_func : no type generator for ListBoxCreateWidgetFunc (GtkListBoxCreateWidgetFunc) for param create_widget_func

// Creates a new GtkModelButton.
/*

C function : gtk_model_button_new
*/
func ModelButtonNew() *ModelButton {
	retC := C.gtk_model_button_new()
	retGo := ModelButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes the child from the notebook.
//
// This function is very similar to gtk_container_remove(),
// but additionally informs the notebook that the removal
// is happening as part of a tab DND operation, which should
// not be cancelled.
/*

C function : gtk_notebook_detach_tab
*/
func (recv *Notebook) DetachTab(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_notebook_detach_tab((*C.GtkNotebook)(recv.native), c_child)

	return
}

// Gets the #GtkPaned:wide-handle property.
/*

C function : gtk_paned_get_wide_handle
*/
func (recv *Paned) GetWideHandle() bool {
	retC := C.gtk_paned_get_wide_handle((*C.GtkPaned)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkPaned:wide-handle property.
/*

C function : gtk_paned_set_wide_handle
*/
func (recv *Paned) SetWideHandle(wide bool) {
	c_wide :=
		boolToGboolean(wide)

	C.gtk_paned_set_wide_handle((*C.GtkPaned)(recv.native), c_wide)

	return
}

// Returns whether show/hide transitions are enabled on this popover.
/*

C function : gtk_popover_get_transitions_enabled
*/
func (recv *Popover) GetTransitionsEnabled() bool {
	retC := C.gtk_popover_get_transitions_enabled((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether show/hide transitions are enabled on this popover
/*

C function : gtk_popover_set_transitions_enabled
*/
func (recv *Popover) SetTransitionsEnabled(transitionsEnabled bool) {
	c_transitions_enabled :=
		boolToGboolean(transitionsEnabled)

	C.gtk_popover_set_transitions_enabled((*C.GtkPopover)(recv.native), c_transitions_enabled)

	return
}

// Creates a new popover menu.
/*

C function : gtk_popover_menu_new
*/
func PopoverMenuNew() *PopoverMenu {
	retC := C.gtk_popover_menu_new()
	retGo := PopoverMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Opens a submenu of the @popover. The @name
// must be one of the names given to the submenus
// of @popover with #GtkPopoverMenu:submenu, or
// "main" to switch back to the main menu.
//
// #GtkModelButton will open submenus automatically
// when the #GtkModelButton:menu-name property is set,
// so this function is only needed when you are using
// other kinds of widgets to initiate menu changes.
/*

C function : gtk_popover_menu_open_submenu
*/
func (recv *PopoverMenu) OpenSubmenu(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_popover_menu_open_submenu((*C.GtkPopoverMenu)(recv.native), c_name)

	return
}

// Unsupported signal 'edge-overshot' for ScrolledWindow : unsupported parameter pos : type PositionType :

// Unsupported signal 'edge-reached' for ScrolledWindow : unsupported parameter pos : type PositionType :

// Returns whether overlay scrolling is enabled for this scrolled window.
/*

C function : gtk_scrolled_window_get_overlay_scrolling
*/
func (recv *ScrolledWindow) GetOverlayScrolling() bool {
	retC := C.gtk_scrolled_window_get_overlay_scrolling((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Enables or disables overlay scrolling for this scrolled window.
/*

C function : gtk_scrolled_window_set_overlay_scrolling
*/
func (recv *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	c_overlay_scrolling :=
		boolToGboolean(overlayScrolling)

	C.gtk_scrolled_window_set_overlay_scrolling((*C.GtkScrolledWindow)(recv.native), c_overlay_scrolling)

	return
}

type signalSearchEntryNextMatchDetail struct {
	callback  SearchEntrySignalNextMatchCallback
	handlerID C.gulong
}

var signalSearchEntryNextMatchId int
var signalSearchEntryNextMatchMap = make(map[int]signalSearchEntryNextMatchDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_next_match(instance, C.gpointer(uintptr(signalSearchEntryNextMatchId)))

	detail := signalSearchEntryNextMatchDetail{callback, handlerID}
	signalSearchEntryNextMatchMap[signalSearchEntryNextMatchId] = detail

	return signalSearchEntryNextMatchId
}

/*
DisconnectNextMatch disconnects a callback from the 'next-match' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectNextMatch.
*/
func (recv *SearchEntry) DisconnectNextMatch(connectionID int) {
	signalSearchEntryNextMatchLock.Lock()
	defer signalSearchEntryNextMatchLock.Unlock()

	detail, exists := signalSearchEntryNextMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntryNextMatchMap, connectionID)
}

//export searchentry_nextMatchHandler
func searchentry_nextMatchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryNextMatchMap[index].callback
	callback()
}

type signalSearchEntryPreviousMatchDetail struct {
	callback  SearchEntrySignalPreviousMatchCallback
	handlerID C.gulong
}

var signalSearchEntryPreviousMatchId int
var signalSearchEntryPreviousMatchMap = make(map[int]signalSearchEntryPreviousMatchDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_previous_match(instance, C.gpointer(uintptr(signalSearchEntryPreviousMatchId)))

	detail := signalSearchEntryPreviousMatchDetail{callback, handlerID}
	signalSearchEntryPreviousMatchMap[signalSearchEntryPreviousMatchId] = detail

	return signalSearchEntryPreviousMatchId
}

/*
DisconnectPreviousMatch disconnects a callback from the 'previous-match' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectPreviousMatch.
*/
func (recv *SearchEntry) DisconnectPreviousMatch(connectionID int) {
	signalSearchEntryPreviousMatchLock.Lock()
	defer signalSearchEntryPreviousMatchLock.Unlock()

	detail, exists := signalSearchEntryPreviousMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntryPreviousMatchMap, connectionID)
}

//export searchentry_previousMatchHandler
func searchentry_previousMatchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryPreviousMatchMap[index].callback
	callback()
}

type signalSearchEntryStopSearchDetail struct {
	callback  SearchEntrySignalStopSearchCallback
	handlerID C.gulong
}

var signalSearchEntryStopSearchId int
var signalSearchEntryStopSearchMap = make(map[int]signalSearchEntryStopSearchDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_stop_search(instance, C.gpointer(uintptr(signalSearchEntryStopSearchId)))

	detail := signalSearchEntryStopSearchDetail{callback, handlerID}
	signalSearchEntryStopSearchMap[signalSearchEntryStopSearchId] = detail

	return signalSearchEntryStopSearchId
}

/*
DisconnectStopSearch disconnects a callback from the 'stop-search' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectStopSearch.
*/
func (recv *SearchEntry) DisconnectStopSearch(connectionID int) {
	signalSearchEntryStopSearchLock.Lock()
	defer signalSearchEntryStopSearchLock.Unlock()

	detail, exists := signalSearchEntryStopSearchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntryStopSearchMap, connectionID)
}

//export searchentry_stopSearchHandler
func searchentry_stopSearchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSearchEntryStopSearchMap[index].callback
	callback()
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Gets whether @stack is horizontally homogeneous.
// See gtk_stack_set_hhomogeneous().
/*

C function : gtk_stack_get_hhomogeneous
*/
func (recv *Stack) GetHhomogeneous() bool {
	retC := C.gtk_stack_get_hhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @stack is vertically homogeneous.
// See gtk_stack_set_vhomogeneous().
/*

C function : gtk_stack_get_vhomogeneous
*/
func (recv *Stack) GetVhomogeneous() bool {
	retC := C.gtk_stack_get_vhomogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkStack to be horizontally homogeneous or not.
// If it is homogeneous, the #GtkStack will request the same
// width for all its children. If it isn't, the stack
// may change width when a different child becomes visible.
/*

C function : gtk_stack_set_hhomogeneous
*/
func (recv *Stack) SetHhomogeneous(hhomogeneous bool) {
	c_hhomogeneous :=
		boolToGboolean(hhomogeneous)

	C.gtk_stack_set_hhomogeneous((*C.GtkStack)(recv.native), c_hhomogeneous)

	return
}

// Sets the #GtkStack to be vertically homogeneous or not.
// If it is homogeneous, the #GtkStack will request the same
// height for all its children. If it isn't, the stack
// may change height when a different child becomes visible.
/*

C function : gtk_stack_set_vhomogeneous
*/
func (recv *Stack) SetVhomogeneous(vhomogeneous bool) {
	c_vhomogeneous :=
		boolToGboolean(vhomogeneous)

	C.gtk_stack_set_vhomogeneous((*C.GtkStack)(recv.native), c_vhomogeneous)

	return
}

// Creates a new sidebar.
/*

C function : gtk_stack_sidebar_new
*/
func StackSidebarNew() *StackSidebar {
	retC := C.gtk_stack_sidebar_new()
	retGo := StackSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the stack.
// See gtk_stack_sidebar_set_stack().
/*

C function : gtk_stack_sidebar_get_stack
*/
func (recv *StackSidebar) GetStack() *Stack {
	retC := C.gtk_stack_sidebar_get_stack((*C.GtkStackSidebar)(recv.native))
	var retGo (*Stack)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StackNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Set the #GtkStack associated with this #GtkStackSidebar.
//
// The sidebar widget will automatically update according to the order
// (packing) and items within the given #GtkStack.
/*

C function : gtk_stack_sidebar_set_stack
*/
func (recv *StackSidebar) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(C.NULL)
	if stack != nil {
		c_stack = (*C.GtkStack)(stack.ToC())
	}

	C.gtk_stack_sidebar_set_stack((*C.GtkStackSidebar)(recv.native), c_stack)

	return
}

// Inserts the text in @markup at position @iter. @markup will be inserted
// in its entirety and must be nul-terminated and valid UTF-8. Emits the
// #GtkTextBuffer::insert-text signal, possibly multiple times; insertion
// actually occurs in the default handler for the signal. @iter will point
// to the end of the inserted text on return.
/*

C function : gtk_text_buffer_insert_markup
*/
func (recv *TextBuffer) InsertMarkup(iter *TextIter, markup string, len int32) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_insert_markup((*C.GtkTextBuffer)(recv.native), c_iter, c_markup, c_len)

	return
}

// Unsupported signal 'extend-selection' for TextView : unsupported parameter granularity : type TextExtendSelection :

// Gets the value of the #GtkTextView:monospace property.
/*

C function : gtk_text_view_get_monospace
*/
func (recv *TextView) GetMonospace() bool {
	retC := C.gtk_text_view_get_monospace((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkTextView:monospace property, which
// indicates that the text view should use monospace
// fonts.
/*

C function : gtk_text_view_set_monospace
*/
func (recv *TextView) SetMonospace(monospace bool) {
	c_monospace :=
		boolToGboolean(monospace)

	C.gtk_text_view_set_monospace((*C.GtkTextView)(recv.native), c_monospace)

	return
}

// Retrieves the #GActionGroup that was registered using @prefix. The resulting
// #GActionGroup may have been registered to @widget or any #GtkWidget in its
// ancestry.
//
// If no action group was found matching @prefix, then %NULL is returned.
/*

C function : gtk_widget_get_action_group
*/
func (recv *Widget) GetActionGroup(prefix string) *gio.ActionGroup {
	c_prefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(c_prefix))

	retC := C.gtk_widget_get_action_group((*C.GtkWidget)(recv.native), c_prefix)
	var retGo (*gio.ActionGroup)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.ActionGroupNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_widget_list_action_prefixes : no return type

// Returns the custom titlebar that has been set with
// gtk_window_set_titlebar().
/*

C function : gtk_window_get_titlebar
*/
func (recv *Window) GetTitlebar() *Widget {
	retC := C.gtk_window_get_titlebar((*C.GtkWindow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
