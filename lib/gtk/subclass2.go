package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr);
void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data);
*/
import "C"

import (
	"github.com/pekim/gobbi/lib/cairo"
	"sync"
	"unsafe"
)

var (
	virtualFunctionsInterfacesLock sync.Mutex
	virtualFunctionsInterfacesId   = 0
	virtualFunctionsInterfaces     = make(map[int]interface{})
)

type DrawingAreaVirtualDraw interface {
	Draw(cr *cairo.Context) bool
}

type DrawingAreaDerivedClass struct {
	name  string
	gtype C.GType
}

type DrawingAreaDerived struct {
	class  *DrawingAreaDerivedClass
	native *C.GtkDrawingArea
}

func DrawingAreaDerive(name string, virtualFunctions interface{}) *DrawingAreaDerivedClass {
	virtualFunctionsInterfacesLock.Lock()
	defer virtualFunctionsInterfacesLock.Unlock()
	virtualFunctionsInterfacesId++
	virtualFunctionsInterfaces[virtualFunctionsInterfacesId] = virtualFunctions

	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkDrawingAreaClass
	typeInfo.class_data = C.gconstpointer(uintptr(virtualFunctionsInterfacesId))
	typeInfo.instance_size = C.sizeof_GtkDrawingArea + C.sizeof_gpointer
	typeInfo.class_init = C.GClassInitFunc(C.drawing_area_class_init)

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_DRAWING_AREA, cTypeName, &typeInfo, 0)

	class := &DrawingAreaDerivedClass{
		name:  name,
		gtype: gtype,
	}

	return class
}

func (c *DrawingAreaDerivedClass) New(virtualFunctions interface{}) *DrawingAreaDerived {
	virtualFunctionsInterfacesLock.Lock()
	defer virtualFunctionsInterfacesLock.Unlock()
	virtualFunctionsInterfacesId++
	virtualFunctionsInterfaces[virtualFunctionsInterfacesId] = virtualFunctions

	object := C.g_object_newv(c.gtype, 0, nil)
	idPointer := C.gpointer(uintptr(object) + uintptr(C.sizeof_GtkDrawingArea))
	*((*int)(idPointer)) = virtualFunctionsInterfacesId

	instance := &DrawingAreaDerived{
		class:  c,
		native: (*C.GtkDrawingArea)(object),
	}

	return instance
}

// DrawingArea upcasts to *DrawingArea
func (recv *DrawingAreaDerived) DrawingArea() *DrawingArea {
	return DrawingAreaNewFromC(unsafe.Pointer(recv.native))

}

//export DrawingAreaClassInit
func DrawingAreaClassInit(class *C.GtkDrawingAreaClass, data unsafe.Pointer) {
	widgetClass := (*C.GtkWidgetClass)(unsafe.Pointer(class))

	virtualFunctions := virtualFunctionsInterfaces[int(uintptr(data))]

	if _, ok := virtualFunctions.(DrawingAreaVirtualDraw); ok {
		widgetClass.draw = (*[0]byte)(C.drawing_area_vf_draw)
	}
}

//export DrawingAreaDraw
func DrawingAreaDraw(daC *C.GtkDrawingArea, contextC *C.cairo_t) C.gboolean {
	idPointer := C.gpointer(uintptr(unsafe.Pointer(daC)) + uintptr(C.sizeof_GtkDrawingArea))
	virtualFunctionsId := *((*int)(idPointer))
	virtualFunctions := virtualFunctionsInterfaces[virtualFunctionsId]

	cr := cairo.ContextNewFromC(unsafe.Pointer(contextC))

	virtualFunctions.(DrawingAreaVirtualDraw).Draw(cr)

	return C.FALSE
}
