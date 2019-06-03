package da

/*
#cgo pkg-config: gtk+-3.0

#include <gtk/gtk.h>
#include <stdlib.h>
#include "drawingarea-derived.h"
*/
import "C"

import (
	"github.com/pekim/gobbi/lib/cairo"
	"sync"
	"unsafe"
)

type DrawingAreaDerivedNative *C.GtkDrawingArea

/*
	A map of ids to Go objects representing the derived widget instance.
	The ids are allocated by incrementing an int for each new instance.

	An id is stored in a C class instance struct when the instance is created.
	It is later retrieved in the Go virtual function Draw, to lookup the
	correspond Go object in the map.
*/
var (
	daIntancesLock sync.Mutex
	daInstanceId   = 0
	daInstances    = make(map[int]*DrawingAreaDerived)
)

type DrawingAreaDerivedClass struct {
	gtype C.GType
}

/*
	Register a new class derived from GtkDrawingArea.
*/
func DrawingAreaDerive() *DrawingAreaDerivedClass {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkDrawingAreaClass
	typeInfo.instance_size = C.sizeof_da_d_instance
	typeInfo.class_init = C.GClassInitFunc(C.drawing_area_class_init)

	cTypeName := C.CString("drawing_area_derived")
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_DRAWING_AREA, cTypeName, &typeInfo, 0)

	class := &DrawingAreaDerivedClass{
		gtype: gtype,
	}

	return class
}

/*
	Create a new instance of the derived class, and use a
	Go object (DrawingAreadDerived) to represent it.
*/
func (c *DrawingAreaDerivedClass) New() *DrawingAreaDerived {
	native := (DrawingAreaDerivedNative)(C.g_object_newv(c.gtype, 0, nil))

	instance := &DrawingAreaDerived{native: native}
	instance.init()

	daIntancesLock.Lock()
	defer daIntancesLock.Unlock()
	daInstanceId++
	// map the id to the Go object
	daInstances[daInstanceId] = instance

	// note the id in the class's instance data struct
	daC := (*C.da_d_instance)(unsafe.Pointer(native))
	daC.instanceId = C.int(daInstanceId)

	return instance
}

/*
	drawingAreaDerivedFromCWidget uses the stored id to lookup a
	DrawingAreaDerived.
*/
func drawingAreaDerivedFromCWidget(widgetC unsafe.Pointer) *DrawingAreaDerived {
	daC := (*C.da_d_instance)(widgetC)
	instanceId := int(daC.instanceId)
	return daInstances[instanceId]
}

// DrawingAreaDraw is called from the C virtual function.
//
//export DrawingAreaDraw
func DrawingAreaDraw(widgetC *C.GtkWidget, contextC *C.cairo_t) C.gboolean {
	cr := cairo.ContextNewFromC(unsafe.Pointer(contextC))

	da := drawingAreaDerivedFromCWidget(unsafe.Pointer(widgetC))
	da.Draw(cr)

	return C.FALSE
}
