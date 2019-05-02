// This is a generated file - DO NOT EDIT

package gdkpixbuf

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Equals compares this PixbufFormat with another PixbufFormat, and returns true if they represent the same GObject.
func (recv *PixbufFormat) Equals(other *PixbufFormat) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PixbufLoaderClass with another PixbufLoaderClass, and returns true if they represent the same GObject.
func (recv *PixbufLoaderClass) Equals(other *PixbufLoaderClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PixbufSimpleAnimClass with another PixbufSimpleAnimClass, and returns true if they represent the same GObject.
func (recv *PixbufSimpleAnimClass) Equals(other *PixbufSimpleAnimClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GdkPixdata
