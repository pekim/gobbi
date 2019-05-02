// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GestureDragNew is a wrapper around the C function gtk_gesture_drag_new.
func GestureDragNew(widget *Widget) *GestureDrag {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_drag_new(c_widget)
	retGo := GestureDragNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GestureLongPressNew is a wrapper around the C function gtk_gesture_long_press_new.
func GestureLongPressNew(widget *Widget) *GestureLongPress {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_long_press_new(c_widget)
	retGo := GestureLongPressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GestureMultiPressNew is a wrapper around the C function gtk_gesture_multi_press_new.
func GestureMultiPressNew(widget *Widget) *GestureMultiPress {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_multi_press_new(c_widget)
	retGo := GestureMultiPressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GesturePanNew is a wrapper around the C function gtk_gesture_pan_new.
func GesturePanNew(widget *Widget, orientation Orientation) *GesturePan {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_gesture_pan_new(c_widget, c_orientation)
	retGo := GesturePanNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GestureRotateNew is a wrapper around the C function gtk_gesture_rotate_new.
func GestureRotateNew(widget *Widget) *GestureRotate {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_rotate_new(c_widget)
	retGo := GestureRotateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GestureSwipeNew is a wrapper around the C function gtk_gesture_swipe_new.
func GestureSwipeNew(widget *Widget) *GestureSwipe {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_swipe_new(c_widget)
	retGo := GestureSwipeNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GestureZoomNew is a wrapper around the C function gtk_gesture_zoom_new.
func GestureZoomNew(widget *Widget) *GestureZoom {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_gesture_zoom_new(c_widget)
	retGo := GestureZoomNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}
