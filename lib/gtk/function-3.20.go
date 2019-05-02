// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// RenderBackgroundGetClip is a wrapper around the C function gtk_render_background_get_clip.
func RenderBackgroundGetClip(context *StyleContext, x float64, y float64, width float64, height float64) *gdk.Rectangle {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	var c_out_clip C.GdkRectangle

	C.gtk_render_background_get_clip(c_context, c_x, c_y, c_width, c_height, &c_out_clip)

	outClip := gdk.RectangleNewFromC(unsafe.Pointer(&c_out_clip))

	return outClip
}
