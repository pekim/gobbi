// This is a generated file - DO NOT EDIT

package cairo

import "unsafe"

// Context is a wrapper around the C record cairo_t.
type Context struct {
	native unsafe.Pointer
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	if u == nil {
		return nil
	}

	g := &Context{native: u}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Device is a wrapper around the C record cairo_device_t.
type Device struct {
	native unsafe.Pointer
}

func DeviceNewFromC(u unsafe.Pointer) *Device {
	if u == nil {
		return nil
	}

	g := &Device{native: u}

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Surface is a wrapper around the C record cairo_surface_t.
type Surface struct {
	native unsafe.Pointer
}

func SurfaceNewFromC(u unsafe.Pointer) *Surface {
	if u == nil {
		return nil
	}

	g := &Surface{native: u}

	return g
}

func (recv *Surface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : cairo_matrix_t

// Pattern is a wrapper around the C record cairo_pattern_t.
type Pattern struct {
	native unsafe.Pointer
}

func PatternNewFromC(u unsafe.Pointer) *Pattern {
	if u == nil {
		return nil
	}

	g := &Pattern{native: u}

	return g
}

func (recv *Pattern) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Region is a wrapper around the C record cairo_region_t.
type Region struct {
	native unsafe.Pointer
}

func RegionNewFromC(u unsafe.Pointer) *Region {
	if u == nil {
		return nil
	}

	g := &Region{native: u}

	return g
}

func (recv *Region) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontOptions is a wrapper around the C record cairo_font_options_t.
type FontOptions struct {
	native unsafe.Pointer
}

func FontOptionsNewFromC(u unsafe.Pointer) *FontOptions {
	if u == nil {
		return nil
	}

	g := &FontOptions{native: u}

	return g
}

func (recv *FontOptions) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontFace is a wrapper around the C record cairo_font_face_t.
type FontFace struct {
	native unsafe.Pointer
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	if u == nil {
		return nil
	}

	g := &FontFace{native: u}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaledFont is a wrapper around the C record cairo_scaled_font_t.
type ScaledFont struct {
	native unsafe.Pointer
}

func ScaledFontNewFromC(u unsafe.Pointer) *ScaledFont {
	if u == nil {
		return nil
	}

	g := &ScaledFont{native: u}

	return g
}

func (recv *ScaledFont) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Path is a wrapper around the C record cairo_path_t.
type Path struct {
	native unsafe.Pointer
}

func PathNewFromC(u unsafe.Pointer) *Path {
	if u == nil {
		return nil
	}

	g := &Path{native: u}

	return g
}

func (recv *Path) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : cairo_rectangle_int_t
