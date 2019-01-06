// +build cairo_1.4 cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

func rectangleNew(native *C.cairo_rectangle_t) Rectangle {
	return Rectangle{
		X:      float64(native.x),
		Y:      float64(native.y),
		Width:  float64(native.width),
		Height: float64(native.height),
	}
}

type RectangleList struct {
	native     *C.cairo_rectangle_list_t
	Status     Status
	Rectangles []Rectangle
}

func rectangleListNew(native *C.cairo_rectangle_list_t) *RectangleList {
	numRectangles := int(native.num_rectangles)
	rectangles := make([]Rectangle, numRectangles, numRectangles)

	rectNative := (*[1 << 30]C.cairo_rectangle_t)(unsafe.Pointer(native.rectangles))
	for i := 0; i < numRectangles; i++ {
		rectangles[i] = rectangleNew(&rectNative[i])
	}

	list := &RectangleList{
		native:     native,
		Status:     Status(native.status),
		Rectangles: rectangles,
	}

	runtime.SetFinalizer(list, func(o *RectangleList) {
		C.cairo_rectangle_list_destroy(o.native)
	})

	return list
}
