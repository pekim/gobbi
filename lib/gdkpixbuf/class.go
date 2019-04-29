// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	"unsafe"
)

// Pixbuf is a wrapper around the C record GdkPixbuf.
type Pixbuf struct {
	native unsafe.Pointer
}

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// PixbufAnimation is a wrapper around the C record GdkPixbufAnimation.
type PixbufAnimation struct {
	native unsafe.Pointer
}

// PixbufAnimationIter is a wrapper around the C record GdkPixbufAnimationIter.
type PixbufAnimationIter struct {
	native unsafe.Pointer
}

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
type PixbufLoader struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PixbufSimpleAnim is a wrapper around the C record GdkPixbufSimpleAnim.
type PixbufSimpleAnim struct {
	native unsafe.Pointer
}

// Unsupported : PixbufSimpleAnimIter : no CType
