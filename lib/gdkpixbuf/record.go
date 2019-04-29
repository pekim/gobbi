// This is a generated file - DO NOT EDIT

package gdkpixbuf

import "unsafe"

// PixbufFormat is a wrapper around the C record GdkPixbufFormat.
type PixbufFormat struct {
	native unsafe.Pointer
}

func PixbufFormatNewFromC(u unsafe.Pointer) *PixbufFormat {
	if u == nil {
		return nil
	}

	g := &PixbufFormat{native: u}

	return g
}

func (recv *PixbufFormat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufLoaderClass is a wrapper around the C record GdkPixbufLoaderClass.
type PixbufLoaderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for size_prepared
	// no type for area_prepared
	// no type for area_updated
	// no type for closed
}

func PixbufLoaderClassNewFromC(u unsafe.Pointer) *PixbufLoaderClass {
	if u == nil {
		return nil
	}

	g := &PixbufLoaderClass{native: u}

	return g
}

func (recv *PixbufLoaderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufSimpleAnimClass is a wrapper around the C record GdkPixbufSimpleAnimClass.
type PixbufSimpleAnimClass struct {
	native unsafe.Pointer
}

func PixbufSimpleAnimClassNewFromC(u unsafe.Pointer) *PixbufSimpleAnimClass {
	if u == nil {
		return nil
	}

	g := &PixbufSimpleAnimClass{native: u}

	return g
}

func (recv *PixbufSimpleAnimClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GdkPixdata
