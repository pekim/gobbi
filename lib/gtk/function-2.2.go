// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Obtains the pixel size of a semantic icon size, possibly
// modified by user preferences for a particular
// #GtkSettings. Normally @size would be
// #GTK_ICON_SIZE_MENU, #GTK_ICON_SIZE_BUTTON, etc.  This function
// isn’t normally needed, gtk_widget_render_icon_pixbuf() is the usual
// way to get an icon for rendering, then just look at the size of
// the rendered pixbuf. The rendered pixbuf may not even correspond to
// the width/height returned by gtk_icon_size_lookup(), because themes
// are free to render the pixbuf however they like, including changing
// the usual size.
/*

C function

gtk_icon_size_lookup_for_settings
*/
func IconSizeLookupForSettings(settings *Settings, size IconSize) (bool, int32, int32) {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	var c_width C.gint

	var c_height C.gint

	retC := C.gtk_icon_size_lookup_for_settings(c_settings, c_size, &c_width, &c_height)
	retGo := retC == C.TRUE

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom
