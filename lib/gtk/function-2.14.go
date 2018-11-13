// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range;
// output values will be in the same range.
/*

C function : gtk_rgb_to_hsv
*/
func RgbToHsv(r float64, g float64, b float64) (float64, float64, float64) {
	c_r := (C.gdouble)(r)

	c_g := (C.gdouble)(g)

	c_b := (C.gdouble)(b)

	var c_h C.gdouble

	var c_s C.gdouble

	var c_v C.gdouble

	C.gtk_rgb_to_hsv(c_r, c_g, c_b, &c_h, &c_s, &c_v)

	h := (float64)(c_h)

	s := (float64)(c_s)

	v := (float64)(c_v)

	return h, s, v
}

// A convenience function for launching the default application
// to show the uri. Like gtk_show_uri_on_window(), but takes a screen
// as transient parent instead of a window.
//
// Note that this function is deprecated as it does not pass the necessary
// information for helpers to parent their dialog properly, when run from
// sandboxed applications for example.
/*

C function : gtk_show_uri
*/
func ShowUri(screen *gdk.Screen, uri string, timestamp uint32) (bool, error) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

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

// Create a simple window with window title @window_title and
// text contents @dialog_text.
// The window will quit any running gtk_main()-loop when destroyed, and it
// will automatically be destroyed upon test function teardown.
/*

C function : gtk_test_create_simple_window
*/
func TestCreateSimpleWindow(windowTitle string, dialogText string) *Widget {
	c_window_title := C.CString(windowTitle)
	defer C.free(unsafe.Pointer(c_window_title))

	c_dialog_text := C.CString(dialogText)
	defer C.free(unsafe.Pointer(c_dialog_text))

	retC := C.gtk_test_create_simple_window(c_window_title, c_dialog_text)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// This function will search @widget and all its descendants for a GtkLabel
// widget with a text string matching @label_pattern.
// The @label_pattern may contain asterisks “*” and question marks “?” as
// placeholders, g_pattern_match() is used for the matching.
// Note that locales other than "C“ tend to alter (translate” label strings,
// so this function is genrally only useful in test programs with
// predetermined locales, see gtk_test_init() for more details.
/*

C function : gtk_test_find_label
*/
func TestFindLabel(widget *Widget, labelPattern string) *Widget {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_label_pattern := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(c_label_pattern))

	retC := C.gtk_test_find_label(c_widget, c_label_pattern)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This function will search siblings of @base_widget and siblings of its
// ancestors for all widgets matching @widget_type.
// Of the matching widgets, the one that is geometrically closest to
// @base_widget will be returned.
// The general purpose of this function is to find the most likely “action”
// widget, relative to another labeling widget. Such as finding a
// button or text entry widget, given its corresponding label widget.
/*

C function : gtk_test_find_sibling
*/
func TestFindSibling(baseWidget *Widget, widgetType gobject.Type) *Widget {
	c_base_widget := (*C.GtkWidget)(C.NULL)
	if baseWidget != nil {
		c_base_widget = (*C.GtkWidget)(baseWidget.ToC())
	}

	c_widget_type := (C.GType)(widgetType)

	retC := C.gtk_test_find_sibling(c_base_widget, c_widget_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This function will search the descendants of @widget for a widget
// of type @widget_type that has a label matching @label_pattern next
// to it. This is most useful for automated GUI testing, e.g. to find
// the “OK” button in a dialog and synthesize clicks on it.
// However see gtk_test_find_label(), gtk_test_find_sibling() and
// gtk_test_widget_click() for possible caveats involving the search of
// such widgets and synthesizing widget events.
/*

C function : gtk_test_find_widget
*/
func TestFindWidget(widget *Widget, labelPattern string, widgetType gobject.Type) *Widget {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_label_pattern := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(c_label_pattern))

	c_widget_type := (C.GType)(widgetType)

	retC := C.gtk_test_find_widget(c_widget, c_label_pattern, c_widget_type)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_test_init : unsupported parameter argcp : array length param argcp is pointer (int*)

// Unsupported : gtk_test_list_all_types : no return type

// Force registration of all core Gtk+ and Gdk object types.
// This allowes to refer to any of those object types via
// g_type_from_name() after calling this function.
/*

C function : gtk_test_register_all_types
*/
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()

	return
}

// Retrive the literal adjustment value for GtkRange based
// widgets and spin buttons. Note that the value returned by
// this function is anything between the lower and upper bounds
// of the adjustment belonging to @widget, and is not a percentage
// as passed in to gtk_test_slider_set_perc().
/*

C function : gtk_test_slider_get_value
*/
func TestSliderGetValue(widget *Widget) float64 {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_test_slider_get_value(c_widget)
	retGo := (float64)(retC)

	return retGo
}

// This function will adjust the slider position of all GtkRange
// based widgets, such as scrollbars or scales, it’ll also adjust
// spin buttons. The adjustment value of these widgets is set to
// a value between the lower and upper limits, according to the
// @percentage argument.
/*

C function : gtk_test_slider_set_perc
*/
func TestSliderSetPerc(widget *Widget, percentage float64) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_percentage := (C.double)(percentage)

	C.gtk_test_slider_set_perc(c_widget, c_percentage)

	return
}

// This function will generate a @button click in the upwards or downwards
// spin button arrow areas, usually leading to an increase or decrease of
// spin button’s value.
/*

C function : gtk_test_spin_button_click
*/
func TestSpinButtonClick(spinner *SpinButton, button uint32, upwards bool) bool {
	c_spinner := (*C.GtkSpinButton)(C.NULL)
	if spinner != nil {
		c_spinner = (*C.GtkSpinButton)(spinner.ToC())
	}

	c_button := (C.guint)(button)

	c_upwards :=
		boolToGboolean(upwards)

	retC := C.gtk_test_spin_button_click(c_spinner, c_button, c_upwards)
	retGo := retC == C.TRUE

	return retGo
}

// Retrive the text string of @widget if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
/*

C function : gtk_test_text_get
*/
func TestTextGet(widget *Widget) string {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_test_text_get(c_widget)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Set the text string of @widget to @string if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
/*

C function : gtk_test_text_set
*/
func TestTextSet(widget *Widget, string string) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.gtk_test_text_set(c_widget, c_string)

	return
}

// This function will generate a @button click (button press and button
// release event) in the middle of the first GdkWindow found that belongs
// to @widget.
// For windowless widgets like #GtkButton (which returns %FALSE from
// gtk_widget_get_has_window()), this will often be an
// input-only event window. For other widgets, this is usually widget->window.
// Certain caveats should be considered when using this function, in
// particular because the mouse pointer is warped to the button click
// location, see gdk_test_simulate_button() for details.
/*

C function : gtk_test_widget_click
*/
func TestWidgetClick(widget *Widget, button uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_click(c_widget, c_button, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// This function will generate keyboard press and release events in
// the middle of the first GdkWindow found that belongs to @widget.
// For windowless widgets like #GtkButton (which returns %FALSE from
// gtk_widget_get_has_window()), this will often be an
// input-only event window. For other widgets, this is usually widget->window.
// Certain caveats should be considered when using this function, in
// particular because the mouse pointer is warped to the key press
// location, see gdk_test_simulate_key() for details.
/*

C function : gtk_test_widget_send_key
*/
func TestWidgetSendKey(widget *Widget, keyval uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_send_key(c_widget, c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}
