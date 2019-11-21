// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var imageSurfaceCreateFunction *gi.Function
var imageSurfaceCreateFunction_Once sync.Once

func imageSurfaceCreateFunction_Set() {
	imageSurfaceCreateFunction_Once.Do(func() {
		imageSurfaceCreateFunction = gi.FunctionInvokerNew("cairo", "image_surface_create")
	})
}

// ImageSurfaceCreate is a representation of the C type cairo_image_surface_create.
func ImageSurfaceCreate() {
	imageSurfaceCreateFunction_Set()

	imageSurfaceCreateFunction.Invoke(nil, nil)

}
