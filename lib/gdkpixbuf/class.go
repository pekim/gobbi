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

func PixbufNewFromC(u unsafe.Pointer) *Pixbuf {
	if u == nil {
		return nil
	}

	g := &Pixbuf{native: u}

	return g
}

func (recv *Pixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_pixbuf_new : return type :

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// Unsupported : gdk_pixbuf_new_from_file : return type :

// Unsupported : gdk_pixbuf_new_from_inline : return type :

// Unsupported : gdk_pixbuf_new_from_xpm_data : return type :

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

func PixbufAnimationNewFromC(u unsafe.Pointer) *PixbufAnimation {
	if u == nil {
		return nil
	}

	g := &PixbufAnimation{native: u}

	return g
}

func (recv *PixbufAnimation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_pixbuf_animation_new_from_file : return type :

// PixbufAnimationIter is a wrapper around the C record GdkPixbufAnimationIter.
type PixbufAnimationIter struct {
	native unsafe.Pointer
}

func PixbufAnimationIterNewFromC(u unsafe.Pointer) *PixbufAnimationIter {
	if u == nil {
		return nil
	}

	g := &PixbufAnimationIter{native: u}

	return g
}

func (recv *PixbufAnimationIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
type PixbufLoader struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func PixbufLoaderNewFromC(u unsafe.Pointer) *PixbufLoader {
	if u == nil {
		return nil
	}

	g := &PixbufLoader{native: u}

	return g
}

func (recv *PixbufLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_pixbuf_loader_new : return type :

// Unsupported : gdk_pixbuf_loader_new_with_type : return type :

// PixbufSimpleAnim is a wrapper around the C record GdkPixbufSimpleAnim.
type PixbufSimpleAnim struct {
	native unsafe.Pointer
}

func PixbufSimpleAnimNewFromC(u unsafe.Pointer) *PixbufSimpleAnim {
	if u == nil {
		return nil
	}

	g := &PixbufSimpleAnim{native: u}

	return g
}

func (recv *PixbufSimpleAnim) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : PixbufSimpleAnimIter : no CType
