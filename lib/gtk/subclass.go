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

type WidgetDrawFunc func(widget *Widget, cr *cairo.Context) bool

type WidgetVirtualFunctions struct {
	draw WidgetDrawFunc
}

type WidgetDerived struct {
	virtualFunctions WidgetVirtualFunctions
}

func WidgetDerive(name string, virtualFunctions WidgetVirtualFunctions) *WidgetDerived {
	derived := &WidgetDerived{
		virtualFunctions: virtualFunctions,
	}

	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkWidgetClass
	typeInfo.instance_size = C.sizeof_GtkWidget

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	//fmt.Println(typeInfo)
	gtype := C.g_type_register_static(C.GTK_TYPE_WIDGET, cTypeName, &typeInfo, 0)
	fmt.Println(gtype)

	return derived
}
