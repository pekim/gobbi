// This is a generated file - DO NOT EDIT

package cairo

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

// Context is a wrapper around the C record cairo_t.
type Context struct {
	native *C.cairo_t
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.cairo_t)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Context with another Context, and returns true if they represent the same GObject.
func (recv *Context) Equals(other *Context) bool {
	return other.ToC() == recv.ToC()
}

// Device is a wrapper around the C record cairo_device_t.
type Device struct {
	native *C.cairo_device_t
}

func DeviceNewFromC(u unsafe.Pointer) *Device {
	c := (*C.cairo_device_t)(u)
	if c == nil {
		return nil
	}

	g := &Device{native: c}

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Device with another Device, and returns true if they represent the same GObject.
func (recv *Device) Equals(other *Device) bool {
	return other.ToC() == recv.ToC()
}

// Surface is a wrapper around the C record cairo_surface_t.
type Surface struct {
	native *C.cairo_surface_t
}

func SurfaceNewFromC(u unsafe.Pointer) *Surface {
	c := (*C.cairo_surface_t)(u)
	if c == nil {
		return nil
	}

	g := &Surface{native: c}

	return g
}

func (recv *Surface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Surface with another Surface, and returns true if they represent the same GObject.
func (recv *Surface) Equals(other *Surface) bool {
	return other.ToC() == recv.ToC()
}

// Matrix is a wrapper around the C record cairo_matrix_t.
type Matrix struct {
	native *C.cairo_matrix_t
}

func MatrixNewFromC(u unsafe.Pointer) *Matrix {
	c := (*C.cairo_matrix_t)(u)
	if c == nil {
		return nil
	}

	g := &Matrix{native: c}

	return g
}

func (recv *Matrix) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Matrix with another Matrix, and returns true if they represent the same GObject.
func (recv *Matrix) Equals(other *Matrix) bool {
	return other.ToC() == recv.ToC()
}

// Pattern is a wrapper around the C record cairo_pattern_t.
type Pattern struct {
	native *C.cairo_pattern_t
}

func PatternNewFromC(u unsafe.Pointer) *Pattern {
	c := (*C.cairo_pattern_t)(u)
	if c == nil {
		return nil
	}

	g := &Pattern{native: c}

	return g
}

func (recv *Pattern) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Pattern with another Pattern, and returns true if they represent the same GObject.
func (recv *Pattern) Equals(other *Pattern) bool {
	return other.ToC() == recv.ToC()
}

// Region is a wrapper around the C record cairo_region_t.
type Region struct {
	native *C.cairo_region_t
}

func RegionNewFromC(u unsafe.Pointer) *Region {
	c := (*C.cairo_region_t)(u)
	if c == nil {
		return nil
	}

	g := &Region{native: c}

	return g
}

func (recv *Region) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Region with another Region, and returns true if they represent the same GObject.
func (recv *Region) Equals(other *Region) bool {
	return other.ToC() == recv.ToC()
}

// FontOptions is a wrapper around the C record cairo_font_options_t.
type FontOptions struct {
	native *C.cairo_font_options_t
}

func FontOptionsNewFromC(u unsafe.Pointer) *FontOptions {
	c := (*C.cairo_font_options_t)(u)
	if c == nil {
		return nil
	}

	g := &FontOptions{native: c}

	return g
}

func (recv *FontOptions) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontOptions with another FontOptions, and returns true if they represent the same GObject.
func (recv *FontOptions) Equals(other *FontOptions) bool {
	return other.ToC() == recv.ToC()
}

// FontFace is a wrapper around the C record cairo_font_face_t.
type FontFace struct {
	native *C.cairo_font_face_t
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	c := (*C.cairo_font_face_t)(u)
	if c == nil {
		return nil
	}

	g := &FontFace{native: c}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontFace with another FontFace, and returns true if they represent the same GObject.
func (recv *FontFace) Equals(other *FontFace) bool {
	return other.ToC() == recv.ToC()
}

// ScaledFont is a wrapper around the C record cairo_scaled_font_t.
type ScaledFont struct {
	native *C.cairo_scaled_font_t
}

func ScaledFontNewFromC(u unsafe.Pointer) *ScaledFont {
	c := (*C.cairo_scaled_font_t)(u)
	if c == nil {
		return nil
	}

	g := &ScaledFont{native: c}

	return g
}

func (recv *ScaledFont) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ScaledFont with another ScaledFont, and returns true if they represent the same GObject.
func (recv *ScaledFont) Equals(other *ScaledFont) bool {
	return other.ToC() == recv.ToC()
}

// Path is a wrapper around the C record cairo_path_t.
type Path struct {
	native *C.cairo_path_t
}

func PathNewFromC(u unsafe.Pointer) *Path {
	c := (*C.cairo_path_t)(u)
	if c == nil {
		return nil
	}

	g := &Path{native: c}

	return g
}

func (recv *Path) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Path with another Path, and returns true if they represent the same GObject.
func (recv *Path) Equals(other *Path) bool {
	return other.ToC() == recv.ToC()
}

// RectangleInt is a wrapper around the C record cairo_rectangle_int_t.
type RectangleInt struct {
	native *C.cairo_rectangle_int_t
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleIntNewFromC(u unsafe.Pointer) *RectangleInt {
	c := (*C.cairo_rectangle_int_t)(u)
	if c == nil {
		return nil
	}

	g := &RectangleInt{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *RectangleInt) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RectangleInt with another RectangleInt, and returns true if they represent the same GObject.
func (recv *RectangleInt) Equals(other *RectangleInt) bool {
	return other.ToC() == recv.ToC()
}
