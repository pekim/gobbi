// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var imageSurfaceCreateFunction *gi.Function
var imageSurfaceCreateFunction_Once sync.Once

func imageSurfaceCreateFunction_Set() error {
	var err error
	imageSurfaceCreateFunction_Once.Do(func() {
		imageSurfaceCreateFunction, err = gi.FunctionInvokerNew("cairo", "image_surface_create")
	})
	return err
}

// ImageSurfaceCreate is a representation of the C type cairo_image_surface_create.
func ImageSurfaceCreate() error {

	err := imageSurfaceCreateFunction_Set()
	if err == nil {
		imageSurfaceCreateFunction.Invoke(nil, nil)
	}

	return err
}
