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

var (
	daIntancesLock sync.Mutex
	daInstanceId   = 0
	daInstances    = make(map[int]*DrawingAreaDerived)
)

type DrawingAreaDerivedClass struct {
	gtype C.GType
}

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

func (c *DrawingAreaDerivedClass) New() *DrawingAreaDerived {
	native := (DrawingAreaDerivedNative)(C.g_object_newv(c.gtype, 0, nil))

	instance := &DrawingAreaDerived{
		class:  c,
		native: native,
	}

	daIntancesLock.Lock()
	defer daIntancesLock.Unlock()
	daInstanceId++
	daInstances[daInstanceId] = instance

	daC := (*C.da_d_instance)(unsafe.Pointer(native))
	daC.instanceId = C.int(daInstanceId)

	return instance
}

//export DrawingAreaDraw
func DrawingAreaDraw(widgetC *C.GtkWidget, contextC *C.cairo_t) C.gboolean {
	daC := (*C.da_d_instance)(unsafe.Pointer(widgetC))
	instanceId := int(daC.instanceId)
	instance := daInstances[instanceId]

	cr := cairo.ContextNewFromC(unsafe.Pointer(contextC))

	instance.Draw(cr)

	return C.FALSE
}
