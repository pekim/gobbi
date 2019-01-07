// +build cairo_1.4 cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"

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
