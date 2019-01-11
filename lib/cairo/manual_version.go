// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

func Version() int {
	return int(C.cairo_version())
}

func VersionString() string {
	return C.GoString(C.cairo_version_string())
}
