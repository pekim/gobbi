// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
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

	GdkGLContext * glarea_createContextHandler(GObject *, gpointer);

	static gulong GLArea_signal_connect_create_context(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-context", G_CALLBACK(glarea_createContextHandler), data);
	}

*/
/*

	gboolean glarea_renderHandler(GObject *, GdkGLContext *, gpointer);

	static gulong GLArea_signal_connect_render(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "render", G_CALLBACK(glarea_renderHandler), data);
	}

*/
/*

	void glarea_resizeHandler(GObject *, gint, gint, gpointer);

	static gulong GLArea_signal_connect_resize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "resize", G_CALLBACK(glarea_resizeHandler), data);
	}

*/
/*

	void scrolledwindow_edgeOvershotHandler(GObject *, GtkPositionType, gpointer);

	static gulong ScrolledWindow_signal_connect_edge_overshot(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "edge-overshot", G_CALLBACK(scrolledwindow_edgeOvershotHandler), data);
	}

*/
/*

	void scrolledwindow_edgeReachedHandler(GObject *, GtkPositionType, gpointer);

	static gulong ScrolledWindow_signal_connect_edge_reached(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "edge-reached", G_CALLBACK(scrolledwindow_edgeReachedHandler), data);
	}

*/
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
/*

	gboolean textview_extendSelectionHandler(GObject *, GtkTextExtendSelection, GtkTextIter *, GtkTextIter *, GtkTextIter *, gpointer);

	static gulong TextView_signal_connect_extend_selection(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "extend-selection", G_CALLBACK(textview_extendSelectionHandler), data);
	}

*/
import "C"

// ClipboardGetDefault is a wrapper around the C function gtk_clipboard_get_default.
func ClipboardGetDefault(display *gdk.Display) *Clipboard {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gtk_clipboard_get_default(c_display)
	retGo := ClipboardNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadFromResource is a wrapper around the C function gtk_css_provider_load_from_resource.
func (recv *CssProvider) LoadFromResource(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.gtk_css_provider_load_from_resource((*C.GtkCssProvider)(recv.native), c_resource_path)

	return
}

// GrabFocusWithoutSelecting is a wrapper around the C function gtk_entry_grab_focus_without_selecting.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GLArea) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GLArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GLArea with another GLArea, and returns true if they represent the same GObject.
func (recv *GLArea) Equals(other *GLArea) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to GLArea.
// Exercise care, as this is a potentially dangerous function if the Object is not a GLArea.
func CastToGLArea(object *gobject.Object) *GLArea {
	return GLAreaNewFromC(object.ToC())
}

type signalGLAreaCreateContextDetail struct {
	callback  GLAreaSignalCreateContextCallback
	handlerID C.gulong
}

var signalGLAreaCreateContextId int
var signalGLAreaCreateContextMap = make(map[int]signalGLAreaCreateContextDetail)
var signalGLAreaCreateContextLock sync.RWMutex

// GLAreaSignalCreateContextCallback is a callback function for a 'create-context' signal emitted from a GLArea.
type GLAreaSignalCreateContextCallback func() gdk.GLContext

/*
ConnectCreateContext connects the callback to the 'create-context' signal for the GLArea.

The returned value represents the connection, and may be passed to DisconnectCreateContext to remove it.
*/
func (recv *GLArea) ConnectCreateContext(callback GLAreaSignalCreateContextCallback) int {
	signalGLAreaCreateContextLock.Lock()
	defer signalGLAreaCreateContextLock.Unlock()

	signalGLAreaCreateContextId++
	instance := C.gpointer(recv.native)
	handlerID := C.GLArea_signal_connect_create_context(instance, C.gpointer(uintptr(signalGLAreaCreateContextId)))

	detail := signalGLAreaCreateContextDetail{callback, handlerID}
	signalGLAreaCreateContextMap[signalGLAreaCreateContextId] = detail

	return signalGLAreaCreateContextId
}

/*
DisconnectCreateContext disconnects a callback from the 'create-context' signal for the GLArea.

The connectionID should be a value returned from a call to ConnectCreateContext.
*/
func (recv *GLArea) DisconnectCreateContext(connectionID int) {
	signalGLAreaCreateContextLock.Lock()
	defer signalGLAreaCreateContextLock.Unlock()

	detail, exists := signalGLAreaCreateContextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGLAreaCreateContextMap, connectionID)
}

//export glarea_createContextHandler
func glarea_createContextHandler(_ *C.GObject, data C.gpointer) *C.GdkGLContext {
	signalGLAreaCreateContextLock.RLock()
	defer signalGLAreaCreateContextLock.RUnlock()

	index := int(uintptr(data))
	callback := signalGLAreaCreateContextMap[index].callback
	retGo := callback()
	retC :=
		(*C.GdkGLContext)(retGo.ToC())
	return retC
}

type signalGLAreaRenderDetail struct {
	callback  GLAreaSignalRenderCallback
	handlerID C.gulong
}

var signalGLAreaRenderId int
var signalGLAreaRenderMap = make(map[int]signalGLAreaRenderDetail)
var signalGLAreaRenderLock sync.RWMutex

// GLAreaSignalRenderCallback is a callback function for a 'render' signal emitted from a GLArea.
type GLAreaSignalRenderCallback func(context *gdk.GLContext) bool

/*
ConnectRender connects the callback to the 'render' signal for the GLArea.

The returned value represents the connection, and may be passed to DisconnectRender to remove it.
*/
func (recv *GLArea) ConnectRender(callback GLAreaSignalRenderCallback) int {
	signalGLAreaRenderLock.Lock()
	defer signalGLAreaRenderLock.Unlock()

	signalGLAreaRenderId++
	instance := C.gpointer(recv.native)
	handlerID := C.GLArea_signal_connect_render(instance, C.gpointer(uintptr(signalGLAreaRenderId)))

	detail := signalGLAreaRenderDetail{callback, handlerID}
	signalGLAreaRenderMap[signalGLAreaRenderId] = detail

	return signalGLAreaRenderId
}

/*
DisconnectRender disconnects a callback from the 'render' signal for the GLArea.

The connectionID should be a value returned from a call to ConnectRender.
*/
func (recv *GLArea) DisconnectRender(connectionID int) {
	signalGLAreaRenderLock.Lock()
	defer signalGLAreaRenderLock.Unlock()

	detail, exists := signalGLAreaRenderMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGLAreaRenderMap, connectionID)
}

//export glarea_renderHandler
func glarea_renderHandler(_ *C.GObject, c_context *C.GdkGLContext, data C.gpointer) C.gboolean {
	signalGLAreaRenderLock.RLock()
	defer signalGLAreaRenderLock.RUnlock()

	context := gdk.GLContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalGLAreaRenderMap[index].callback
	retGo := callback(context)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalGLAreaResizeDetail struct {
	callback  GLAreaSignalResizeCallback
	handlerID C.gulong
}

var signalGLAreaResizeId int
var signalGLAreaResizeMap = make(map[int]signalGLAreaResizeDetail)
var signalGLAreaResizeLock sync.RWMutex

// GLAreaSignalResizeCallback is a callback function for a 'resize' signal emitted from a GLArea.
type GLAreaSignalResizeCallback func(width int32, height int32)

/*
ConnectResize connects the callback to the 'resize' signal for the GLArea.

The returned value represents the connection, and may be passed to DisconnectResize to remove it.
*/
func (recv *GLArea) ConnectResize(callback GLAreaSignalResizeCallback) int {
	signalGLAreaResizeLock.Lock()
	defer signalGLAreaResizeLock.Unlock()

	signalGLAreaResizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.GLArea_signal_connect_resize(instance, C.gpointer(uintptr(signalGLAreaResizeId)))

	detail := signalGLAreaResizeDetail{callback, handlerID}
	signalGLAreaResizeMap[signalGLAreaResizeId] = detail

	return signalGLAreaResizeId
}

/*
DisconnectResize disconnects a callback from the 'resize' signal for the GLArea.

The connectionID should be a value returned from a call to ConnectResize.
*/
func (recv *GLArea) DisconnectResize(connectionID int) {
	signalGLAreaResizeLock.Lock()
	defer signalGLAreaResizeLock.Unlock()

	detail, exists := signalGLAreaResizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGLAreaResizeMap, connectionID)
}

//export glarea_resizeHandler
func glarea_resizeHandler(_ *C.GObject, c_width C.gint, c_height C.gint, data C.gpointer) {
	signalGLAreaResizeLock.RLock()
	defer signalGLAreaResizeLock.RUnlock()

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalGLAreaResizeMap[index].callback
	callback(width, height)
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
	var retGo (*glib.Error)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ErrorNewFromC(unsafe.Pointer(retC))
	}

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
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

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

// Unsupported : gtk_list_box_bind_model : unsupported parameter create_widget_func : no type generator for ListBoxCreateWidgetFunc (GtkListBoxCreateWidgetFunc) for param create_widget_func

// ModelButtonNew is a wrapper around the C function gtk_model_button_new.
func ModelButtonNew() *ModelButton {
	retC := C.gtk_model_button_new()
	retGo := ModelButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DetachTab is a wrapper around the C function gtk_notebook_detach_tab.
func (recv *Notebook) DetachTab(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_notebook_detach_tab((*C.GtkNotebook)(recv.native), c_child)

	return
}

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

type signalScrolledWindowEdgeOvershotDetail struct {
	callback  ScrolledWindowSignalEdgeOvershotCallback
	handlerID C.gulong
}

var signalScrolledWindowEdgeOvershotId int
var signalScrolledWindowEdgeOvershotMap = make(map[int]signalScrolledWindowEdgeOvershotDetail)
var signalScrolledWindowEdgeOvershotLock sync.RWMutex

// ScrolledWindowSignalEdgeOvershotCallback is a callback function for a 'edge-overshot' signal emitted from a ScrolledWindow.
type ScrolledWindowSignalEdgeOvershotCallback func(pos PositionType)

/*
ConnectEdgeOvershot connects the callback to the 'edge-overshot' signal for the ScrolledWindow.

The returned value represents the connection, and may be passed to DisconnectEdgeOvershot to remove it.
*/
func (recv *ScrolledWindow) ConnectEdgeOvershot(callback ScrolledWindowSignalEdgeOvershotCallback) int {
	signalScrolledWindowEdgeOvershotLock.Lock()
	defer signalScrolledWindowEdgeOvershotLock.Unlock()

	signalScrolledWindowEdgeOvershotId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScrolledWindow_signal_connect_edge_overshot(instance, C.gpointer(uintptr(signalScrolledWindowEdgeOvershotId)))

	detail := signalScrolledWindowEdgeOvershotDetail{callback, handlerID}
	signalScrolledWindowEdgeOvershotMap[signalScrolledWindowEdgeOvershotId] = detail

	return signalScrolledWindowEdgeOvershotId
}

/*
DisconnectEdgeOvershot disconnects a callback from the 'edge-overshot' signal for the ScrolledWindow.

The connectionID should be a value returned from a call to ConnectEdgeOvershot.
*/
func (recv *ScrolledWindow) DisconnectEdgeOvershot(connectionID int) {
	signalScrolledWindowEdgeOvershotLock.Lock()
	defer signalScrolledWindowEdgeOvershotLock.Unlock()

	detail, exists := signalScrolledWindowEdgeOvershotMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScrolledWindowEdgeOvershotMap, connectionID)
}

//export scrolledwindow_edgeOvershotHandler
func scrolledwindow_edgeOvershotHandler(_ *C.GObject, c_pos C.GtkPositionType, data C.gpointer) {
	signalScrolledWindowEdgeOvershotLock.RLock()
	defer signalScrolledWindowEdgeOvershotLock.RUnlock()

	pos := PositionType(c_pos)

	index := int(uintptr(data))
	callback := signalScrolledWindowEdgeOvershotMap[index].callback
	callback(pos)
}

type signalScrolledWindowEdgeReachedDetail struct {
	callback  ScrolledWindowSignalEdgeReachedCallback
	handlerID C.gulong
}

var signalScrolledWindowEdgeReachedId int
var signalScrolledWindowEdgeReachedMap = make(map[int]signalScrolledWindowEdgeReachedDetail)
var signalScrolledWindowEdgeReachedLock sync.RWMutex

// ScrolledWindowSignalEdgeReachedCallback is a callback function for a 'edge-reached' signal emitted from a ScrolledWindow.
type ScrolledWindowSignalEdgeReachedCallback func(pos PositionType)

/*
ConnectEdgeReached connects the callback to the 'edge-reached' signal for the ScrolledWindow.

The returned value represents the connection, and may be passed to DisconnectEdgeReached to remove it.
*/
func (recv *ScrolledWindow) ConnectEdgeReached(callback ScrolledWindowSignalEdgeReachedCallback) int {
	signalScrolledWindowEdgeReachedLock.Lock()
	defer signalScrolledWindowEdgeReachedLock.Unlock()

	signalScrolledWindowEdgeReachedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScrolledWindow_signal_connect_edge_reached(instance, C.gpointer(uintptr(signalScrolledWindowEdgeReachedId)))

	detail := signalScrolledWindowEdgeReachedDetail{callback, handlerID}
	signalScrolledWindowEdgeReachedMap[signalScrolledWindowEdgeReachedId] = detail

	return signalScrolledWindowEdgeReachedId
}

/*
DisconnectEdgeReached disconnects a callback from the 'edge-reached' signal for the ScrolledWindow.

The connectionID should be a value returned from a call to ConnectEdgeReached.
*/
func (recv *ScrolledWindow) DisconnectEdgeReached(connectionID int) {
	signalScrolledWindowEdgeReachedLock.Lock()
	defer signalScrolledWindowEdgeReachedLock.Unlock()

	detail, exists := signalScrolledWindowEdgeReachedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScrolledWindowEdgeReachedMap, connectionID)
}

//export scrolledwindow_edgeReachedHandler
func scrolledwindow_edgeReachedHandler(_ *C.GObject, c_pos C.GtkPositionType, data C.gpointer) {
	signalScrolledWindowEdgeReachedLock.RLock()
	defer signalScrolledWindowEdgeReachedLock.RUnlock()

	pos := PositionType(c_pos)

	index := int(uintptr(data))
	callback := signalScrolledWindowEdgeReachedMap[index].callback
	callback(pos)
}

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

type signalSearchEntryNextMatchDetail struct {
	callback  SearchEntrySignalNextMatchCallback
	handlerID C.gulong
}

var signalSearchEntryNextMatchId int
var signalSearchEntryNextMatchMap = make(map[int]signalSearchEntryNextMatchDetail)
var signalSearchEntryNextMatchLock sync.RWMutex

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
	signalSearchEntryNextMatchLock.RLock()
	defer signalSearchEntryNextMatchLock.RUnlock()

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
var signalSearchEntryPreviousMatchLock sync.RWMutex

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
	signalSearchEntryPreviousMatchLock.RLock()
	defer signalSearchEntryPreviousMatchLock.RUnlock()

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
var signalSearchEntryStopSearchLock sync.RWMutex

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
	signalSearchEntryStopSearchLock.RLock()
	defer signalSearchEntryStopSearchLock.RUnlock()

	index := int(uintptr(data))
	callback := signalSearchEntryStopSearchMap[index].callback
	callback()
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

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
	var retGo (*Stack)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StackNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_sidebar_set_stack.
func (recv *StackSidebar) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(C.NULL)
	if stack != nil {
		c_stack = (*C.GtkStack)(stack.ToC())
	}

	C.gtk_stack_sidebar_set_stack((*C.GtkStackSidebar)(recv.native), c_stack)

	return
}

// InsertMarkup is a wrapper around the C function gtk_text_buffer_insert_markup.
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

type signalTextViewExtendSelectionDetail struct {
	callback  TextViewSignalExtendSelectionCallback
	handlerID C.gulong
}

var signalTextViewExtendSelectionId int
var signalTextViewExtendSelectionMap = make(map[int]signalTextViewExtendSelectionDetail)
var signalTextViewExtendSelectionLock sync.RWMutex

// TextViewSignalExtendSelectionCallback is a callback function for a 'extend-selection' signal emitted from a TextView.
type TextViewSignalExtendSelectionCallback func(granularity TextExtendSelection, location *TextIter, start *TextIter, end *TextIter) bool

/*
ConnectExtendSelection connects the callback to the 'extend-selection' signal for the TextView.

The returned value represents the connection, and may be passed to DisconnectExtendSelection to remove it.
*/
func (recv *TextView) ConnectExtendSelection(callback TextViewSignalExtendSelectionCallback) int {
	signalTextViewExtendSelectionLock.Lock()
	defer signalTextViewExtendSelectionLock.Unlock()

	signalTextViewExtendSelectionId++
	instance := C.gpointer(recv.native)
	handlerID := C.TextView_signal_connect_extend_selection(instance, C.gpointer(uintptr(signalTextViewExtendSelectionId)))

	detail := signalTextViewExtendSelectionDetail{callback, handlerID}
	signalTextViewExtendSelectionMap[signalTextViewExtendSelectionId] = detail

	return signalTextViewExtendSelectionId
}

/*
DisconnectExtendSelection disconnects a callback from the 'extend-selection' signal for the TextView.

The connectionID should be a value returned from a call to ConnectExtendSelection.
*/
func (recv *TextView) DisconnectExtendSelection(connectionID int) {
	signalTextViewExtendSelectionLock.Lock()
	defer signalTextViewExtendSelectionLock.Unlock()

	detail, exists := signalTextViewExtendSelectionMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextViewExtendSelectionMap, connectionID)
}

//export textview_extendSelectionHandler
func textview_extendSelectionHandler(_ *C.GObject, c_granularity C.GtkTextExtendSelection, c_location *C.GtkTextIter, c_start *C.GtkTextIter, c_end *C.GtkTextIter, data C.gpointer) C.gboolean {
	signalTextViewExtendSelectionLock.RLock()
	defer signalTextViewExtendSelectionLock.RUnlock()

	granularity := TextExtendSelection(c_granularity)

	location := TextIterNewFromC(unsafe.Pointer(c_location))

	start := TextIterNewFromC(unsafe.Pointer(c_start))

	end := TextIterNewFromC(unsafe.Pointer(c_end))

	index := int(uintptr(data))
	callback := signalTextViewExtendSelectionMap[index].callback
	retGo := callback(granularity, location, start, end)
	retC :=
		boolToGboolean(retGo)
	return retC
}

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

// GetActionGroup is a wrapper around the C function gtk_widget_get_action_group.
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

// ListActionPrefixes is a wrapper around the C function gtk_widget_list_action_prefixes.
func (recv *Widget) ListActionPrefixes() []string {
	retC := C.gtk_widget_list_action_prefixes((*C.GtkWidget)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetTitlebar is a wrapper around the C function gtk_window_get_titlebar.
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
