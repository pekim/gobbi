// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
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

// Renders an icon using gtk_render_icon_pixbuf() and converts it to a
// cairo surface.
//
// This function never returns %NULL; if the icon can’t be rendered
// (perhaps because an image file fails to load), a default "missing
// image" icon will be returned instead.
/*

C function : gtk_icon_set_render_icon_surface
*/
func (recv *IconSet) RenderIconSurface(context *StyleContext, size IconSize, scale int32, forWindow *gdk.Window) *cairo.Surface {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	c_scale := (C.int)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	retC := C.gtk_icon_set_render_icon_surface((*C.GtkIconSet)(recv.native), c_context, c_size, c_scale, c_for_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// Automatically assign an object declared in the class template XML to be set to a location
// on a freshly built instance’s private data, or alternatively accessible via gtk_widget_get_template_child().
//
// The struct can point either into the public instance, then you should use G_STRUCT_OFFSET(WidgetType, member)
// for @struct_offset,  or in the private struct, then you should use G_PRIVATE_OFFSET(WidgetType, member).
//
// An explicit strong reference will be held automatically for the duration of your
// instance’s life cycle, it will be released automatically when #GObjectClass.dispose() runs
// on your instance and if a @struct_offset that is != 0 is specified, then the automatic location
// in your instance public or private data will be set to %NULL. You can however access an automated child
// pointer the first time your classes #GObjectClass.dispose() runs, or alternatively in
// #GtkWidgetClass.destroy().
//
// If @internal_child is specified, #GtkBuildableIface.get_internal_child() will be automatically
// implemented by the #GtkWidget class so there is no need to implement it manually.
//
// The wrapper macros gtk_widget_class_bind_template_child(), gtk_widget_class_bind_template_child_internal(),
// gtk_widget_class_bind_template_child_private() and gtk_widget_class_bind_template_child_internal_private()
// might be more convenient to use.
//
// Note that this must be called from a composite widget classes class
// initializer after calling gtk_widget_class_set_template().
/*

C function : gtk_widget_class_bind_template_child_full
*/
func (recv *WidgetClass) BindTemplateChildFull(name string, internalChild bool, structOffset int64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_internal_child :=
		boolToGboolean(internalChild)

	c_struct_offset := (C.gssize)(structOffset)

	C.gtk_widget_class_bind_template_child_full((*C.GtkWidgetClass)(recv.native), c_name, c_internal_child, c_struct_offset)

	return
}

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param connect_func

// This should be called at class initialization time to specify
// the GtkBuilder XML to be used to extend a widget.
//
// For convenience, gtk_widget_class_set_template_from_resource() is also provided.
//
// Note that any class that installs templates must call gtk_widget_init_template()
// in the widget’s instance initializer.
/*

C function : gtk_widget_class_set_template
*/
func (recv *WidgetClass) SetTemplate(templateBytes *glib.Bytes) {
	c_template_bytes := (*C.GBytes)(C.NULL)
	if templateBytes != nil {
		c_template_bytes = (*C.GBytes)(templateBytes.ToC())
	}

	C.gtk_widget_class_set_template((*C.GtkWidgetClass)(recv.native), c_template_bytes)

	return
}

// A convenience function to call gtk_widget_class_set_template().
//
// Note that any class that installs templates must call gtk_widget_init_template()
// in the widget’s instance initializer.
/*

C function : gtk_widget_class_set_template_from_resource
*/
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	c_resource_name := C.CString(resourceName)
	defer C.free(unsafe.Pointer(c_resource_name))

	C.gtk_widget_class_set_template_from_resource((*C.GtkWidgetClass)(recv.native), c_resource_name)

	return
}
