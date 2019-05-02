// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "C"

// AddFrame is a wrapper around the C function gdk_pixbuf_simple_anim_add_frame.
func (recv *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gdk_pixbuf_simple_anim_add_frame((*C.GdkPixbufSimpleAnim)(recv.native), c_pixbuf)

	return
}
