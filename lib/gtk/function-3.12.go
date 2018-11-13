// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Get the direction of the current locale. This is the expected
// reading direction for text and UI.
//
// This function depends on the current locale being set with
// setlocale() and will default to setting the %GTK_TEXT_DIR_LTR
// direction otherwise. %GTK_TEXT_DIR_NONE will never be returned.
//
// GTK+ sets the default text direction according to the locale
// during gtk_init(), and you should normally use
// gtk_widget_get_direction() or gtk_widget_get_default_direction()
// to obtain the current direcion.
//
// This function is only needed rare cases when the locale is
// changed after GTK+ has already been initialized. In this case,
// you can use it to update the default text direction as follows:
//
// |[<!-- language="C" -->
// setlocale (LC_ALL, new_locale);
// direction = gtk_get_locale_direction ();
// gtk_widget_set_default_direction (direction);
// ]|
/*

C function

gtk_get_locale_direction
*/
func GetLocaleDirection() TextDirection {
	retC := C.gtk_get_locale_direction()
	retGo := (TextDirection)(retC)

	return retGo
}
