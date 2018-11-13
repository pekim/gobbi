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

// Returns the product ID of this device, or %NULL if this information couldn't
// be obtained. This ID is retrieved from the device, and is thus constant for
// it. See gdk_device_get_vendor_id() for more information.
/*

C function : gdk_device_get_product_id
*/
func (recv *Device) GetProductId() string {
	retC := C.gdk_device_get_product_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the vendor ID of this device, or %NULL if this information couldn't
// be obtained. This ID is retrieved from the device, and is thus constant for
// it.
//
// This function, together with gdk_device_get_product_id(), can be used to eg.
// compose #GSettings paths to store settings for this device.
//
// |[<!-- language="C" -->
// static GSettings *
// get_device_settings (GdkDevice *device)
// {
// const gchar *vendor, *product;
// GSettings *settings;
// GdkDevice *device;
// gchar *path;
//
// vendor = gdk_device_get_vendor_id (device);
// product = gdk_device_get_product_id (device);
//
// path = g_strdup_printf ("/org/example/app/devices/%s:%s/", vendor, product);
// settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
// g_free (path);
//
// return settings;
// }
// ]|
/*

C function : gdk_device_get_vendor_id
*/
func (recv *Device) GetVendorId() string {
	retC := C.gdk_device_get_vendor_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the value set using gdk_gl_context_set_debug_enabled().
/*

C function : gdk_gl_context_get_debug_enabled
*/
func (recv *GLContext) GetDebugEnabled() bool {
	retC := C.gdk_gl_context_get_debug_enabled((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the #GdkDisplay the @context is created for
/*

C function : gdk_gl_context_get_display
*/
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

// Retrieves the value set using gdk_gl_context_set_forward_compatible().
/*

C function : gdk_gl_context_get_forward_compatible
*/
func (recv *GLContext) GetForwardCompatible() bool {
	retC := C.gdk_gl_context_get_forward_compatible((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the major and minor version requested by calling
// gdk_gl_context_set_required_version().
/*

C function : gdk_gl_context_get_required_version
*/
func (recv *GLContext) GetRequiredVersion() (int32, int32) {
	var c_major C.int

	var c_minor C.int

	C.gdk_gl_context_get_required_version((*C.GdkGLContext)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// Retrieves the #GdkGLContext that this @context share data with.
/*

C function : gdk_gl_context_get_shared_context
*/
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

// Retrieves the OpenGL version of the @context.
//
// The @context must be realized prior to calling this function.
/*

C function : gdk_gl_context_get_version
*/
func (recv *GLContext) GetVersion() (int32, int32) {
	var c_major C.int

	var c_minor C.int

	C.gdk_gl_context_get_version((*C.GdkGLContext)(recv.native), &c_major, &c_minor)

	major := (int32)(c_major)

	minor := (int32)(c_minor)

	return major, minor
}

// Retrieves the #GdkWindow used by the @context.
/*

C function : gdk_gl_context_get_window
*/
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

// Makes the @context the current one.
/*

C function : gdk_gl_context_make_current
*/
func (recv *GLContext) MakeCurrent() {
	C.gdk_gl_context_make_current((*C.GdkGLContext)(recv.native))

	return
}

// Realizes the given #GdkGLContext.
//
// It is safe to call this function on a realized #GdkGLContext.
/*

C function : gdk_gl_context_realize
*/
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

// Sets whether the #GdkGLContext should perform extra validations and
// run time checking. This is useful during development, but has
// additional overhead.
//
// The #GdkGLContext must not be realized or made current prior to
// calling this function.
/*

C function : gdk_gl_context_set_debug_enabled
*/
func (recv *GLContext) SetDebugEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gdk_gl_context_set_debug_enabled((*C.GdkGLContext)(recv.native), c_enabled)

	return
}

// Sets whether the #GdkGLContext should be forward compatible.
//
// Forward compatibile contexts must not support OpenGL functionality that
// has been marked as deprecated in the requested version; non-forward
// compatible contexts, on the other hand, must support both deprecated and
// non deprecated functionality.
//
// The #GdkGLContext must not be realized or made current prior to calling
// this function.
/*

C function : gdk_gl_context_set_forward_compatible
*/
func (recv *GLContext) SetForwardCompatible(compatible bool) {
	c_compatible :=
		boolToGboolean(compatible)

	C.gdk_gl_context_set_forward_compatible((*C.GdkGLContext)(recv.native), c_compatible)

	return
}

// Sets the major and minor version of OpenGL to request.
//
// Setting @major and @minor to zero will use the default values.
//
// The #GdkGLContext must not be realized or made current prior to calling
// this function.
/*

C function : gdk_gl_context_set_required_version
*/
func (recv *GLContext) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.int)(major)

	c_minor := (C.int)(minor)

	C.gdk_gl_context_set_required_version((*C.GdkGLContext)(recv.native), c_major, c_minor)

	return
}

// Creates a new #GdkGLContext matching the
// framebuffer format to the visual of the #GdkWindow. The context
// is disconnected from any particular window or surface.
//
// If the creation of the #GdkGLContext failed, @error will be set.
//
// Before using the returned #GdkGLContext, you will need to
// call gdk_gl_context_make_current() or gdk_gl_context_realize().
/*

C function : gdk_window_create_gl_context
*/
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

// If you call this during a paint (e.g. between gdk_window_begin_paint_region()
// and gdk_window_end_paint() then GDK will mark the current clip region of the
// window as being drawn. This is required when mixing GL rendering via
// gdk_cairo_draw_from_gl() and cairo rendering, as otherwise GDK has no way
// of knowing when something paints over the GL-drawn regions.
//
// This is typically called automatically by GTK+ and you don't need
// to care about this.
/*

C function : gdk_window_mark_paint_from_clip
*/
func (recv *Window) MarkPaintFromClip(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	C.gdk_window_mark_paint_from_clip((*C.GdkWindow)(recv.native), c_cr)

	return
}
