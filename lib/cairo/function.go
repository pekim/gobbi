// Code generated - DO NOT EDIT.

package cairo

import gi "github.com/pekim/gobbi/internal/gi"

var imageSurfaceCreateInvoker *gi.FunctionInvoker

// ImageSurfaceCreate is a representation of the C type cairo_image_surface_create.
func ImageSurfaceCreate() {
	if imageSurfaceCreateInvoker == nil {
		imageSurfaceCreateInvoker = gi.FunctionInvokerNew("cairo", "image_surface_create")
	}

	imageSurfaceCreateInvoker.Call()
}
