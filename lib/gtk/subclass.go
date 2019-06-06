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
	"fmt"
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
	new   func() WidgetDerivedInitializer
}

type WidgetDerivedInitializer interface {
	Init()
}

func WidgetDerive(name string, new func() WidgetDerivedInitializer) *WidgetDerivedClass {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkWidgetClass
	typeInfo.instance_size = C.sizeof_GtkWidget

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_WIDGET, cTypeName, &typeInfo, 0)

	class := &WidgetDerivedClass{
		name:  name,
		gtype: gtype,
		new:   new,
	}

	return class
}

func (c *WidgetDerivedClass) New() interface{} {
	instance := c.new()

	object := C.g_object_newv(c.gtype, 0, nil)
	fmt.Println(object)

	instance.Init()

	return instance
}
