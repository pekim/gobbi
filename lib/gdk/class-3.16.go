// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"C"
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// GLContextClearCurrent is a wrapper around the C function gdk_gl_context_clear_current.
func GLContextClearCurrent() {
	C.gdk_gl_context_clear_current()

	return
}

// GLContextGetCurrent is a wrapper around the C function gdk_gl_context_get_current.
func GLContextGetCurrent() *GLContext {
	retC := C.gdk_gl_context_get_current()
	var retGo (*GLContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GLContextNewFromC(unsafe.Pointer(retC))
	}

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
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetForwardCompatible is a wrapper around the C function gdk_gl_context_get_forward_compatible.
func (recv *GLContext) GetForwardCompatible() bool {
	retC := C.gdk_gl_context_get_forward_compatible((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRequiredVersion is a wrapper around the C function gdk_gl_context_get_required_version.
func (recv *GLContext) GetRequiredVersion() (int32, int32) {
	var c_major C.int

	var c_minor C.int

	C.gdk_gl_context_get_required_version((*C.GdkGLContext)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// GetSharedContext is a wrapper around the C function gdk_gl_context_get_shared_context.
func (recv *GLContext) GetSharedContext() *GLContext {
	retC := C.gdk_gl_context_get_shared_context((*C.GdkGLContext)(recv.native))
	var retGo (*GLContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GLContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetVersion is a wrapper around the C function gdk_gl_context_get_version.
func (recv *GLContext) GetVersion() (int32, int32) {
	var c_major C.int

	var c_minor C.int

	C.gdk_gl_context_get_version((*C.GdkGLContext)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// GetWindow is a wrapper around the C function gdk_gl_context_get_window.
func (recv *GLContext) GetWindow() *Window {
	retC := C.gdk_gl_context_get_window((*C.GdkGLContext)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MarkPaintFromClip is a wrapper around the C function gdk_window_mark_paint_from_clip.
func (recv *Window) MarkPaintFromClip(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	C.gdk_window_mark_paint_from_clip((*C.GdkWindow)(recv.native), c_cr)

	return
}
