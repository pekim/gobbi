package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>

gboolean widget_vf_draw(GtkWidget *widget, cairo_t *cr) {
	return TRUE;
}

*/
import "C"

import (
	"github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

type WidgetVirtualDraw interface {
	Draw(widget *Widget, cr *cairo.Context) bool
}

//type WidgetVirtualFunctions struct {
//	draw WidgetDrawFunc
//}

type WidgetDerivedClass struct {
	name  string
	gtype C.GType
}

type WidgetDerived struct {
	class  *WidgetDerivedClass
	native *C.GtkWidget
}

func WidgetDerive(name string) *WidgetDerivedClass {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkWidgetClass
	typeInfo.instance_size = C.sizeof_GtkWidget

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_WIDGET, cTypeName, &typeInfo, 0)

	class := &WidgetDerivedClass{
		name:  name,
		gtype: gtype,
	}

	return class
}

func (c *WidgetDerivedClass) New(virtualFunctions interface{}) *WidgetDerived {
	native := (*C.GtkWidget)(C.g_object_newv(c.gtype, 0, nil))

	instance := &WidgetDerived{
		class:  c,
		native: native,
	}

	return instance
}

// Widget upcasts to *Widget
func (recv *WidgetDerived) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))

}
