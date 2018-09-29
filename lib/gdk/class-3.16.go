// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetProductId is a wrapper around the C function gdk_device_get_product_id.
func (recv *Device) GetProductId() string {
	retC := C.gdk_device_get_product_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVendorId is a wrapper around the C function gdk_device_get_vendor_id.
func (recv *Device) GetVendorId() string {
	retC := C.gdk_device_get_vendor_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDebugEnabled is a wrapper around the C function gdk_gl_context_get_debug_enabled.
func (recv *GLContext) GetDebugEnabled() bool {
	retC := C.gdk_gl_context_get_debug_enabled((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_gl_context_get_display.
func (recv *GLContext) GetDisplay() *Display {
	retC := C.gdk_gl_context_get_display((*C.GdkGLContext)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetForwardCompatible is a wrapper around the C function gdk_gl_context_get_forward_compatible.
func (recv *GLContext) GetForwardCompatible() bool {
	retC := C.gdk_gl_context_get_forward_compatible((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// GetSharedContext is a wrapper around the C function gdk_gl_context_get_shared_context.
func (recv *GLContext) GetSharedContext() *GLContext {
	retC := C.gdk_gl_context_get_shared_context((*C.GdkGLContext)(recv.native))
	retGo := GLContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// GetWindow is a wrapper around the C function gdk_gl_context_get_window.
func (recv *GLContext) GetWindow() *Window {
	retC := C.gdk_gl_context_get_window((*C.GdkGLContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MakeCurrent is a wrapper around the C function gdk_gl_context_make_current.
func (recv *GLContext) MakeCurrent() {
	C.gdk_gl_context_make_current((*C.GdkGLContext)(recv.native))

	return
}

// Realize is a wrapper around the C function gdk_gl_context_realize.
func (recv *GLContext) Realize() (bool, error) {
	var cThrowableError *C.GError

	retC := C.gdk_gl_context_realize((*C.GdkGLContext)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetDebugEnabled is a wrapper around the C function gdk_gl_context_set_debug_enabled.
func (recv *GLContext) SetDebugEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gdk_gl_context_set_debug_enabled((*C.GdkGLContext)(recv.native), c_enabled)

	return
}

// SetForwardCompatible is a wrapper around the C function gdk_gl_context_set_forward_compatible.
func (recv *GLContext) SetForwardCompatible(compatible bool) {
	c_compatible :=
		boolToGboolean(compatible)

	C.gdk_gl_context_set_forward_compatible((*C.GdkGLContext)(recv.native), c_compatible)

	return
}

// SetRequiredVersion is a wrapper around the C function gdk_gl_context_set_required_version.
func (recv *GLContext) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.int)(major)

	c_minor := (C.int)(minor)

	C.gdk_gl_context_set_required_version((*C.GdkGLContext)(recv.native), c_major, c_minor)

	return
}

// CreateGlContext is a wrapper around the C function gdk_window_create_gl_context.
func (recv *Window) CreateGlContext() (*GLContext, error) {
	var cThrowableError *C.GError

	retC := C.gdk_window_create_gl_context((*C.GdkWindow)(recv.native), &cThrowableError)
	retGo := GLContextNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MarkPaintFromClip is a wrapper around the C function gdk_window_mark_paint_from_clip.
func (recv *Window) MarkPaintFromClip(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(cr.ToC())

	C.gdk_window_mark_paint_from_clip((*C.GdkWindow)(recv.native), c_cr)

	return
}
