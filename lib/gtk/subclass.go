package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>

gboolean widget_vf_draw(GtkWidget *widget, cairo_t *cr) {
	printf("C draw\n");
	return TRUE;
}

void widget_class_init(GtkWidgetClass g_class, gpointer class_data) {
	g_class.draw = widget_vf_draw;
	printf("class init\n");
}

*/
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

type WidgetVirtualDraw interface {
	Draw(cr *cairo.Context) bool
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
	typeInfo.class_init = C.GClassInitFunc(C.widget_class_init)

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
	f, ok := virtualFunctions.(WidgetVirtualDraw)
	fmt.Println("draw func :", ok)
	f.Draw(nil)

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
