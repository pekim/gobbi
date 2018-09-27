// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

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

// GLAreaNew is a wrapper around the C function gtk_gl_area_new.
func GLAreaNew() *Widget {
	retC := C.gtk_gl_area_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_gl_area_attach_buffers : no return generator

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

// Unsupported : gtk_gl_area_get_required_version : unsupported parameter major : no type generator for gint, gint*

// Unsupported : gtk_gl_area_make_current : no return generator

// Unsupported : gtk_gl_area_queue_render : no return generator

// Unsupported : gtk_gl_area_set_auto_render : no return generator

// Unsupported : gtk_gl_area_set_error : no return generator

// Unsupported : gtk_gl_area_set_has_alpha : no return generator

// Unsupported : gtk_gl_area_set_has_depth_buffer : no return generator

// Unsupported : gtk_gl_area_set_has_stencil_buffer : no return generator

// Unsupported : gtk_gl_area_set_required_version : no return generator

// Unsupported : gtk_gl_area_set_use_es : no return generator
