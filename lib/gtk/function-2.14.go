// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_rgb_to_hsv : unsupported parameter h : no type generator for gdouble, gdouble*

// ShowUri is a wrapper around the C function gtk_show_uri.
func ShowUri(screen *gdk.Screen, uri string, timestamp uint32) (bool, error) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_timestamp := (C.guint32)(timestamp)

	var cThrowableError *C.GError

	retC := C.gtk_show_uri(c_screen, c_uri, c_timestamp, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// TestCreateSimpleWindow is a wrapper around the C function gtk_test_create_simple_window.
func TestCreateSimpleWindow(windowTitle string, dialogText string) *Widget {
	c_window_title := C.CString(windowTitle)
	defer C.free(unsafe.Pointer(c_window_title))

	c_dialog_text := C.CString(dialogText)
	defer C.free(unsafe.Pointer(c_dialog_text))

	retC := C.gtk_test_create_simple_window(c_window_title, c_dialog_text)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_test_create_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// TestFindLabel is a wrapper around the C function gtk_test_find_label.
func TestFindLabel(widget *Widget, labelPattern string) *Widget {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_label_pattern := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(c_label_pattern))

	retC := C.gtk_test_find_label(c_widget, c_label_pattern)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_test_find_sibling : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_find_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_init : unsupported parameter argcp : no type generator for gint, int*

// Unsupported : gtk_test_list_all_types : unsupported parameter n_types : no type generator for guint, guint*

// TestRegisterAllTypes is a wrapper around the C function gtk_test_register_all_types.
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()

	return
}

// TestSliderGetValue is a wrapper around the C function gtk_test_slider_get_value.
func TestSliderGetValue(widget *Widget) float64 {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_test_slider_get_value(c_widget)
	retGo := (float64)(retC)

	return retGo
}

// TestSliderSetPerc is a wrapper around the C function gtk_test_slider_set_perc.
func TestSliderSetPerc(widget *Widget, percentage float64) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_percentage := (C.double)(percentage)

	C.gtk_test_slider_set_perc(c_widget, c_percentage)

	return
}

// TestSpinButtonClick is a wrapper around the C function gtk_test_spin_button_click.
func TestSpinButtonClick(spinner *SpinButton, button uint32, upwards bool) bool {
	c_spinner := (*C.GtkSpinButton)(spinner.ToC())

	c_button := (C.guint)(button)

	c_upwards :=
		boolToGboolean(upwards)

	retC := C.gtk_test_spin_button_click(c_spinner, c_button, c_upwards)
	retGo := retC == C.TRUE

	return retGo
}

// TestTextGet is a wrapper around the C function gtk_test_text_get.
func TestTextGet(widget *Widget) string {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_test_text_get(c_widget)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TestTextSet is a wrapper around the C function gtk_test_text_set.
func TestTextSet(widget *Widget, string string) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.gtk_test_text_set(c_widget, c_string)

	return
}

// TestWidgetClick is a wrapper around the C function gtk_test_widget_click.
func TestWidgetClick(widget *Widget, button uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_click(c_widget, c_button, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// TestWidgetSendKey is a wrapper around the C function gtk_test_widget_send_key.
func TestWidgetSendKey(widget *Widget, keyval uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_send_key(c_widget, c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}
