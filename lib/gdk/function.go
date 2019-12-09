// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

var addOptionEntriesLibgtkOnlyFunction *gi.Function
var addOptionEntriesLibgtkOnlyFunction_Once sync.Once

func addOptionEntriesLibgtkOnlyFunction_Set() error {
	var err error
	addOptionEntriesLibgtkOnlyFunction_Once.Do(func() {
		addOptionEntriesLibgtkOnlyFunction, err = gi.FunctionInvokerNew("Gdk", "add_option_entries_libgtk_only")
	})
	return err
}

// AddOptionEntriesLibgtkOnly is a representation of the C type gdk_add_option_entries_libgtk_only.
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(group.Native())

	err := addOptionEntriesLibgtkOnlyFunction_Set()
	if err == nil {
		addOptionEntriesLibgtkOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var atomInternFunction *gi.Function
var atomInternFunction_Once sync.Once

func atomInternFunction_Set() error {
	var err error
	atomInternFunction_Once.Do(func() {
		atomInternFunction, err = gi.FunctionInvokerNew("Gdk", "atom_intern")
	})
	return err
}

// AtomIntern is a representation of the C type gdk_atom_intern.
func AtomIntern(atomName string, onlyIfExists bool) *Atom {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(atomName)
	inArgs[1].SetBoolean(onlyIfExists)

	var ret gi.Argument

	err := atomInternFunction_Set()
	if err == nil {
		ret = atomInternFunction.Invoke(inArgs[:], nil)
	}

	retGo := AtomNewFromNative(ret.Pointer())

	return retGo
}

var atomInternStaticStringFunction *gi.Function
var atomInternStaticStringFunction_Once sync.Once

func atomInternStaticStringFunction_Set() error {
	var err error
	atomInternStaticStringFunction_Once.Do(func() {
		atomInternStaticStringFunction, err = gi.FunctionInvokerNew("Gdk", "atom_intern_static_string")
	})
	return err
}

// AtomInternStaticString is a representation of the C type gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) *Atom {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(atomName)

	var ret gi.Argument

	err := atomInternStaticStringFunction_Set()
	if err == nil {
		ret = atomInternStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := AtomNewFromNative(ret.Pointer())

	return retGo
}

var beepFunction *gi.Function
var beepFunction_Once sync.Once

func beepFunction_Set() error {
	var err error
	beepFunction_Once.Do(func() {
		beepFunction, err = gi.FunctionInvokerNew("Gdk", "beep")
	})
	return err
}

// Beep is a representation of the C type gdk_beep.
func Beep() {

	err := beepFunction_Set()
	if err == nil {
		beepFunction.Invoke(nil, nil)
	}

	return
}

var cairoCreateFunction *gi.Function
var cairoCreateFunction_Once sync.Once

func cairoCreateFunction_Set() error {
	var err error
	cairoCreateFunction_Once.Do(func() {
		cairoCreateFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_create")
	})
	return err
}

// CairoCreate is a representation of the C type gdk_cairo_create.
func CairoCreate(window *Window) *cairo.Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native())

	var ret gi.Argument

	err := cairoCreateFunction_Set()
	if err == nil {
		ret = cairoCreateFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.ContextNewFromNative(ret.Pointer())

	return retGo
}

var cairoDrawFromGlFunction *gi.Function
var cairoDrawFromGlFunction_Once sync.Once

func cairoDrawFromGlFunction_Set() error {
	var err error
	cairoDrawFromGlFunction_Once.Do(func() {
		cairoDrawFromGlFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_draw_from_gl")
	})
	return err
}

// CairoDrawFromGl is a representation of the C type gdk_cairo_draw_from_gl.
func CairoDrawFromGl(cr *cairo.Context, window *Window, source int32, sourceType int32, bufferScale int32, x int32, y int32, width int32, height int32) {
	var inArgs [9]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(window.Native())
	inArgs[2].SetInt32(source)
	inArgs[3].SetInt32(sourceType)
	inArgs[4].SetInt32(bufferScale)
	inArgs[5].SetInt32(x)
	inArgs[6].SetInt32(y)
	inArgs[7].SetInt32(width)
	inArgs[8].SetInt32(height)

	err := cairoDrawFromGlFunction_Set()
	if err == nil {
		cairoDrawFromGlFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoGetClipRectangleFunction *gi.Function
var cairoGetClipRectangleFunction_Once sync.Once

func cairoGetClipRectangleFunction_Set() error {
	var err error
	cairoGetClipRectangleFunction_Once.Do(func() {
		cairoGetClipRectangleFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_get_clip_rectangle")
	})
	return err
}

// CairoGetClipRectangle is a representation of the C type gdk_cairo_get_clip_rectangle.
func CairoGetClipRectangle(cr *cairo.Context) (bool, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cr.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := cairoGetClipRectangleFunction_Set()
	if err == nil {
		ret = cairoGetClipRectangleFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var cairoGetDrawingContextFunction *gi.Function
var cairoGetDrawingContextFunction_Once sync.Once

func cairoGetDrawingContextFunction_Set() error {
	var err error
	cairoGetDrawingContextFunction_Once.Do(func() {
		cairoGetDrawingContextFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_get_drawing_context")
	})
	return err
}

// CairoGetDrawingContext is a representation of the C type gdk_cairo_get_drawing_context.
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cr.Native())

	var ret gi.Argument

	err := cairoGetDrawingContextFunction_Set()
	if err == nil {
		ret = cairoGetDrawingContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := DrawingContextNewFromNative(ret.Pointer())

	return retGo
}

var cairoRectangleFunction *gi.Function
var cairoRectangleFunction_Once sync.Once

func cairoRectangleFunction_Set() error {
	var err error
	cairoRectangleFunction_Once.Do(func() {
		cairoRectangleFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_rectangle")
	})
	return err
}

// CairoRectangle is a representation of the C type gdk_cairo_rectangle.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(rectangle.Native())

	err := cairoRectangleFunction_Set()
	if err == nil {
		cairoRectangleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoRegionFunction *gi.Function
var cairoRegionFunction_Once sync.Once

func cairoRegionFunction_Set() error {
	var err error
	cairoRegionFunction_Once.Do(func() {
		cairoRegionFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_region")
	})
	return err
}

// CairoRegion is a representation of the C type gdk_cairo_region.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(region.Native())

	err := cairoRegionFunction_Set()
	if err == nil {
		cairoRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoRegionCreateFromSurfaceFunction *gi.Function
var cairoRegionCreateFromSurfaceFunction_Once sync.Once

func cairoRegionCreateFromSurfaceFunction_Set() error {
	var err error
	cairoRegionCreateFromSurfaceFunction_Once.Do(func() {
		cairoRegionCreateFromSurfaceFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_region_create_from_surface")
	})
	return err
}

// CairoRegionCreateFromSurface is a representation of the C type gdk_cairo_region_create_from_surface.
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(surface.Native())

	var ret gi.Argument

	err := cairoRegionCreateFromSurfaceFunction_Set()
	if err == nil {
		ret = cairoRegionCreateFromSurfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.RegionNewFromNative(ret.Pointer())

	return retGo
}

var cairoSetSourceColorFunction *gi.Function
var cairoSetSourceColorFunction_Once sync.Once

func cairoSetSourceColorFunction_Set() error {
	var err error
	cairoSetSourceColorFunction_Once.Do(func() {
		cairoSetSourceColorFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_set_source_color")
	})
	return err
}

// CairoSetSourceColor is a representation of the C type gdk_cairo_set_source_color.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(color.Native())

	err := cairoSetSourceColorFunction_Set()
	if err == nil {
		cairoSetSourceColorFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoSetSourcePixbufFunction *gi.Function
var cairoSetSourcePixbufFunction_Once sync.Once

func cairoSetSourcePixbufFunction_Set() error {
	var err error
	cairoSetSourcePixbufFunction_Once.Do(func() {
		cairoSetSourcePixbufFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_set_source_pixbuf")
	})
	return err
}

// CairoSetSourcePixbuf is a representation of the C type gdk_cairo_set_source_pixbuf.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(pixbuf.Native())
	inArgs[2].SetFloat64(pixbufX)
	inArgs[3].SetFloat64(pixbufY)

	err := cairoSetSourcePixbufFunction_Set()
	if err == nil {
		cairoSetSourcePixbufFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoSetSourceRgbaFunction *gi.Function
var cairoSetSourceRgbaFunction_Once sync.Once

func cairoSetSourceRgbaFunction_Set() error {
	var err error
	cairoSetSourceRgbaFunction_Once.Do(func() {
		cairoSetSourceRgbaFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_set_source_rgba")
	})
	return err
}

// CairoSetSourceRgba is a representation of the C type gdk_cairo_set_source_rgba.
func CairoSetSourceRgba(cr *cairo.Context, rgba *RGBA) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(rgba.Native())

	err := cairoSetSourceRgbaFunction_Set()
	if err == nil {
		cairoSetSourceRgbaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoSetSourceWindowFunction *gi.Function
var cairoSetSourceWindowFunction_Once sync.Once

func cairoSetSourceWindowFunction_Set() error {
	var err error
	cairoSetSourceWindowFunction_Once.Do(func() {
		cairoSetSourceWindowFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_set_source_window")
	})
	return err
}

// CairoSetSourceWindow is a representation of the C type gdk_cairo_set_source_window.
func CairoSetSourceWindow(cr *cairo.Context, window *Window, x float64, y float64) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(window.Native())
	inArgs[2].SetFloat64(x)
	inArgs[3].SetFloat64(y)

	err := cairoSetSourceWindowFunction_Set()
	if err == nil {
		cairoSetSourceWindowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cairoSurfaceCreateFromPixbufFunction *gi.Function
var cairoSurfaceCreateFromPixbufFunction_Once sync.Once

func cairoSurfaceCreateFromPixbufFunction_Set() error {
	var err error
	cairoSurfaceCreateFromPixbufFunction_Once.Do(func() {
		cairoSurfaceCreateFromPixbufFunction, err = gi.FunctionInvokerNew("Gdk", "cairo_surface_create_from_pixbuf")
	})
	return err
}

// CairoSurfaceCreateFromPixbuf is a representation of the C type gdk_cairo_surface_create_from_pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int32, forWindow *Window) *cairo.Surface {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(pixbuf.Native())
	inArgs[1].SetInt32(scale)
	inArgs[2].SetPointer(forWindow.Native())

	var ret gi.Argument

	err := cairoSurfaceCreateFromPixbufFunction_Set()
	if err == nil {
		ret = cairoSurfaceCreateFromPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.SurfaceNewFromNative(ret.Pointer())

	return retGo
}

var colorParseFunction *gi.Function
var colorParseFunction_Once sync.Once

func colorParseFunction_Set() error {
	var err error
	colorParseFunction_Once.Do(func() {
		colorParseFunction, err = gi.FunctionInvokerNew("Gdk", "color_parse")
	})
	return err
}

// ColorParse is a representation of the C type gdk_color_parse.
func ColorParse(spec string) (bool, *Color) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(spec)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := colorParseFunction_Set()
	if err == nil {
		ret = colorParseFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := ColorNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var disableMultideviceFunction *gi.Function
var disableMultideviceFunction_Once sync.Once

func disableMultideviceFunction_Set() error {
	var err error
	disableMultideviceFunction_Once.Do(func() {
		disableMultideviceFunction, err = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	})
	return err
}

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {

	err := disableMultideviceFunction_Set()
	if err == nil {
		disableMultideviceFunction.Invoke(nil, nil)
	}

	return
}

var dragAbortFunction *gi.Function
var dragAbortFunction_Once sync.Once

func dragAbortFunction_Set() error {
	var err error
	dragAbortFunction_Once.Do(func() {
		dragAbortFunction, err = gi.FunctionInvokerNew("Gdk", "drag_abort")
	})
	return err
}

// DragAbort is a representation of the C type gdk_drag_abort.
func DragAbort(context *DragContext, time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetUint32(time)

	err := dragAbortFunction_Set()
	if err == nil {
		dragAbortFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragBeginFunction *gi.Function
var dragBeginFunction_Once sync.Once

func dragBeginFunction_Set() error {
	var err error
	dragBeginFunction_Once.Do(func() {
		dragBeginFunction, err = gi.FunctionInvokerNew("Gdk", "drag_begin")
	})
	return err
}

// DragBegin is a representation of the C type gdk_drag_begin.
func DragBegin(window *Window, targets *glib.List) *DragContext {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(targets.Native())

	var ret gi.Argument

	err := dragBeginFunction_Set()
	if err == nil {
		ret = dragBeginFunction.Invoke(inArgs[:], nil)
	}

	retGo := DragContextNewFromNative(ret.Pointer())

	return retGo
}

var dragBeginForDeviceFunction *gi.Function
var dragBeginForDeviceFunction_Once sync.Once

func dragBeginForDeviceFunction_Set() error {
	var err error
	dragBeginForDeviceFunction_Once.Do(func() {
		dragBeginForDeviceFunction, err = gi.FunctionInvokerNew("Gdk", "drag_begin_for_device")
	})
	return err
}

// DragBeginForDevice is a representation of the C type gdk_drag_begin_for_device.
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) *DragContext {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(device.Native())
	inArgs[2].SetPointer(targets.Native())

	var ret gi.Argument

	err := dragBeginForDeviceFunction_Set()
	if err == nil {
		ret = dragBeginForDeviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := DragContextNewFromNative(ret.Pointer())

	return retGo
}

var dragBeginFromPointFunction *gi.Function
var dragBeginFromPointFunction_Once sync.Once

func dragBeginFromPointFunction_Set() error {
	var err error
	dragBeginFromPointFunction_Once.Do(func() {
		dragBeginFromPointFunction, err = gi.FunctionInvokerNew("Gdk", "drag_begin_from_point")
	})
	return err
}

// DragBeginFromPoint is a representation of the C type gdk_drag_begin_from_point.
func DragBeginFromPoint(window *Window, device *Device, targets *glib.List, xRoot int32, yRoot int32) *DragContext {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(device.Native())
	inArgs[2].SetPointer(targets.Native())
	inArgs[3].SetInt32(xRoot)
	inArgs[4].SetInt32(yRoot)

	var ret gi.Argument

	err := dragBeginFromPointFunction_Set()
	if err == nil {
		ret = dragBeginFromPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := DragContextNewFromNative(ret.Pointer())

	return retGo
}

var dragDropFunction *gi.Function
var dragDropFunction_Once sync.Once

func dragDropFunction_Set() error {
	var err error
	dragDropFunction_Once.Do(func() {
		dragDropFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop")
	})
	return err
}

// DragDrop is a representation of the C type gdk_drag_drop.
func DragDrop(context *DragContext, time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetUint32(time)

	err := dragDropFunction_Set()
	if err == nil {
		dragDropFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragDropDoneFunction *gi.Function
var dragDropDoneFunction_Once sync.Once

func dragDropDoneFunction_Set() error {
	var err error
	dragDropDoneFunction_Once.Do(func() {
		dragDropDoneFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop_done")
	})
	return err
}

// DragDropDone is a representation of the C type gdk_drag_drop_done.
func DragDropDone(context *DragContext, success bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetBoolean(success)

	err := dragDropDoneFunction_Set()
	if err == nil {
		dragDropDoneFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragDropSucceededFunction *gi.Function
var dragDropSucceededFunction_Once sync.Once

func dragDropSucceededFunction_Set() error {
	var err error
	dragDropSucceededFunction_Once.Do(func() {
		dragDropSucceededFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop_succeeded")
	})
	return err
}

// DragDropSucceeded is a representation of the C type gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := dragDropSucceededFunction_Set()
	if err == nil {
		ret = dragDropSucceededFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var dragFindWindowForScreenFunction *gi.Function
var dragFindWindowForScreenFunction_Once sync.Once

func dragFindWindowForScreenFunction_Set() error {
	var err error
	dragFindWindowForScreenFunction_Once.Do(func() {
		dragFindWindowForScreenFunction, err = gi.FunctionInvokerNew("Gdk", "drag_find_window_for_screen")
	})
	return err
}

// DragFindWindowForScreen is a representation of the C type gdk_drag_find_window_for_screen.
func DragFindWindowForScreen(context *DragContext, dragWindow *Window, screen *Screen, xRoot int32, yRoot int32) (*Window, DragProtocol) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(dragWindow.Native())
	inArgs[2].SetPointer(screen.Native())
	inArgs[3].SetInt32(xRoot)
	inArgs[4].SetInt32(yRoot)

	var outArgs [2]gi.Argument

	err := dragFindWindowForScreenFunction_Set()
	if err == nil {
		dragFindWindowForScreenFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := WindowNewFromNative(outArgs[0].Pointer())
	out1 := DragProtocol(outArgs[1].Int32())

	return out0, out1
}

var dragGetSelectionFunction *gi.Function
var dragGetSelectionFunction_Once sync.Once

func dragGetSelectionFunction_Set() error {
	var err error
	dragGetSelectionFunction_Once.Do(func() {
		dragGetSelectionFunction, err = gi.FunctionInvokerNew("Gdk", "drag_get_selection")
	})
	return err
}

// DragGetSelection is a representation of the C type gdk_drag_get_selection.
func DragGetSelection(context *DragContext) *Atom {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := dragGetSelectionFunction_Set()
	if err == nil {
		ret = dragGetSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := AtomNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gdk_drag_motion' : parameter 'suggested_action' of type 'DragAction' not supported

// UNSUPPORTED : C value 'gdk_drag_status' : parameter 'action' of type 'DragAction' not supported

var dropFinishFunction *gi.Function
var dropFinishFunction_Once sync.Once

func dropFinishFunction_Set() error {
	var err error
	dropFinishFunction_Once.Do(func() {
		dropFinishFunction, err = gi.FunctionInvokerNew("Gdk", "drop_finish")
	})
	return err
}

// DropFinish is a representation of the C type gdk_drop_finish.
func DropFinish(context *DragContext, success bool, time uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetBoolean(success)
	inArgs[2].SetUint32(time)

	err := dropFinishFunction_Set()
	if err == nil {
		dropFinishFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dropReplyFunction *gi.Function
var dropReplyFunction_Once sync.Once

func dropReplyFunction_Set() error {
	var err error
	dropReplyFunction_Once.Do(func() {
		dropReplyFunction, err = gi.FunctionInvokerNew("Gdk", "drop_reply")
	})
	return err
}

// DropReply is a representation of the C type gdk_drop_reply.
func DropReply(context *DragContext, accepted bool, time uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetBoolean(accepted)
	inArgs[2].SetUint32(time)

	err := dropReplyFunction_Set()
	if err == nil {
		dropReplyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var errorTrapPopFunction *gi.Function
var errorTrapPopFunction_Once sync.Once

func errorTrapPopFunction_Set() error {
	var err error
	errorTrapPopFunction_Once.Do(func() {
		errorTrapPopFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	})
	return err
}

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {

	var ret gi.Argument

	err := errorTrapPopFunction_Set()
	if err == nil {
		ret = errorTrapPopFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
}

var errorTrapPopIgnoredFunction *gi.Function
var errorTrapPopIgnoredFunction_Once sync.Once

func errorTrapPopIgnoredFunction_Set() error {
	var err error
	errorTrapPopIgnoredFunction_Once.Do(func() {
		errorTrapPopIgnoredFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	})
	return err
}

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {

	err := errorTrapPopIgnoredFunction_Set()
	if err == nil {
		errorTrapPopIgnoredFunction.Invoke(nil, nil)
	}

	return
}

var errorTrapPushFunction *gi.Function
var errorTrapPushFunction_Once sync.Once

func errorTrapPushFunction_Set() error {
	var err error
	errorTrapPushFunction_Once.Do(func() {
		errorTrapPushFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	})
	return err
}

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {

	err := errorTrapPushFunction_Set()
	if err == nil {
		errorTrapPushFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : parameter 'func' of type 'EventFunc' not supported

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

var eventRequestMotionsFunction *gi.Function
var eventRequestMotionsFunction_Once sync.Once

func eventRequestMotionsFunction_Set() error {
	var err error
	eventRequestMotionsFunction_Once.Do(func() {
		eventRequestMotionsFunction, err = gi.FunctionInvokerNew("Gdk", "event_request_motions")
	})
	return err
}

// EventRequestMotions is a representation of the C type gdk_event_request_motions.
func EventRequestMotions(event *EventMotion) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(event.Native())

	err := eventRequestMotionsFunction_Set()
	if err == nil {
		eventRequestMotionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_events_get_angle' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_center' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_distance' : parameter 'event1' of type 'Event' not supported

var eventsPendingFunction *gi.Function
var eventsPendingFunction_Once sync.Once

func eventsPendingFunction_Set() error {
	var err error
	eventsPendingFunction_Once.Do(func() {
		eventsPendingFunction, err = gi.FunctionInvokerNew("Gdk", "events_pending")
	})
	return err
}

// EventsPending is a representation of the C type gdk_events_pending.
func EventsPending() bool {

	var ret gi.Argument

	err := eventsPendingFunction_Set()
	if err == nil {
		ret = eventsPendingFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var flushFunction *gi.Function
var flushFunction_Once sync.Once

func flushFunction_Set() error {
	var err error
	flushFunction_Once.Do(func() {
		flushFunction, err = gi.FunctionInvokerNew("Gdk", "flush")
	})
	return err
}

// Flush is a representation of the C type gdk_flush.
func Flush() {

	err := flushFunction_Set()
	if err == nil {
		flushFunction.Invoke(nil, nil)
	}

	return
}

var getDefaultRootWindowFunction *gi.Function
var getDefaultRootWindowFunction_Once sync.Once

func getDefaultRootWindowFunction_Set() error {
	var err error
	getDefaultRootWindowFunction_Once.Do(func() {
		getDefaultRootWindowFunction, err = gi.FunctionInvokerNew("Gdk", "get_default_root_window")
	})
	return err
}

// GetDefaultRootWindow is a representation of the C type gdk_get_default_root_window.
func GetDefaultRootWindow() *Window {

	var ret gi.Argument

	err := getDefaultRootWindowFunction_Set()
	if err == nil {
		ret = getDefaultRootWindowFunction.Invoke(nil, nil)
	}

	retGo := WindowNewFromNative(ret.Pointer())

	return retGo
}

var getDisplayFunction *gi.Function
var getDisplayFunction_Once sync.Once

func getDisplayFunction_Set() error {
	var err error
	getDisplayFunction_Once.Do(func() {
		getDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {

	var ret gi.Argument

	err := getDisplayFunction_Set()
	if err == nil {
		ret = getDisplayFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var getDisplayArgNameFunction *gi.Function
var getDisplayArgNameFunction_Once sync.Once

func getDisplayArgNameFunction_Set() error {
	var err error
	getDisplayArgNameFunction_Once.Do(func() {
		getDisplayArgNameFunction, err = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	})
	return err
}

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {

	var ret gi.Argument

	err := getDisplayArgNameFunction_Set()
	if err == nil {
		ret = getDisplayArgNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getProgramClassFunction *gi.Function
var getProgramClassFunction_Once sync.Once

func getProgramClassFunction_Set() error {
	var err error
	getProgramClassFunction_Once.Do(func() {
		getProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "get_program_class")
	})
	return err
}

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {

	var ret gi.Argument

	err := getProgramClassFunction_Set()
	if err == nil {
		ret = getProgramClassFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getShowEventsFunction *gi.Function
var getShowEventsFunction_Once sync.Once

func getShowEventsFunction_Set() error {
	var err error
	getShowEventsFunction_Once.Do(func() {
		getShowEventsFunction, err = gi.FunctionInvokerNew("Gdk", "get_show_events")
	})
	return err
}

// GetShowEvents is a representation of the C type gdk_get_show_events.
func GetShowEvents() bool {

	var ret gi.Argument

	err := getShowEventsFunction_Set()
	if err == nil {
		ret = getShowEventsFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var glErrorQuarkFunction *gi.Function
var glErrorQuarkFunction_Once sync.Once

func glErrorQuarkFunction_Set() error {
	var err error
	glErrorQuarkFunction_Once.Do(func() {
		glErrorQuarkFunction, err = gi.FunctionInvokerNew("Gdk", "gl_error_quark")
	})
	return err
}

// GlErrorQuark is a representation of the C type gdk_gl_error_quark.
func GlErrorQuark() glib.Quark {

	var ret gi.Argument

	err := glErrorQuarkFunction_Set()
	if err == nil {
		ret = glErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'gdk_init' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_init_check' : parameter 'argv' of type 'nil' not supported

var keyboardGrabFunction *gi.Function
var keyboardGrabFunction_Once sync.Once

func keyboardGrabFunction_Set() error {
	var err error
	keyboardGrabFunction_Once.Do(func() {
		keyboardGrabFunction, err = gi.FunctionInvokerNew("Gdk", "keyboard_grab")
	})
	return err
}

// KeyboardGrab is a representation of the C type gdk_keyboard_grab.
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) GrabStatus {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetBoolean(ownerEvents)
	inArgs[2].SetUint32(time)

	var ret gi.Argument

	err := keyboardGrabFunction_Set()
	if err == nil {
		ret = keyboardGrabFunction.Invoke(inArgs[:], nil)
	}

	retGo := GrabStatus(ret.Int32())

	return retGo
}

var keyboardUngrabFunction *gi.Function
var keyboardUngrabFunction_Once sync.Once

func keyboardUngrabFunction_Set() error {
	var err error
	keyboardUngrabFunction_Once.Do(func() {
		keyboardUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "keyboard_ungrab")
	})
	return err
}

// KeyboardUngrab is a representation of the C type gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := keyboardUngrabFunction_Set()
	if err == nil {
		keyboardUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var keyvalConvertCaseFunction *gi.Function
var keyvalConvertCaseFunction_Once sync.Once

func keyvalConvertCaseFunction_Set() error {
	var err error
	keyvalConvertCaseFunction_Once.Do(func() {
		keyvalConvertCaseFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_convert_case")
	})
	return err
}

// KeyvalConvertCase is a representation of the C type gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(symbol)

	var outArgs [2]gi.Argument

	err := keyvalConvertCaseFunction_Set()
	if err == nil {
		keyvalConvertCaseFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Uint32()

	return out0, out1
}

var keyvalFromNameFunction *gi.Function
var keyvalFromNameFunction_Once sync.Once

func keyvalFromNameFunction_Set() error {
	var err error
	keyvalFromNameFunction_Once.Do(func() {
		keyvalFromNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_from_name")
	})
	return err
}

// KeyvalFromName is a representation of the C type gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(keyvalName)

	var ret gi.Argument

	err := keyvalFromNameFunction_Set()
	if err == nil {
		ret = keyvalFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalIsLowerFunction *gi.Function
var keyvalIsLowerFunction_Once sync.Once

func keyvalIsLowerFunction_Set() error {
	var err error
	keyvalIsLowerFunction_Once.Do(func() {
		keyvalIsLowerFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_is_lower")
	})
	return err
}

// KeyvalIsLower is a representation of the C type gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalIsLowerFunction_Set()
	if err == nil {
		ret = keyvalIsLowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keyvalIsUpperFunction *gi.Function
var keyvalIsUpperFunction_Once sync.Once

func keyvalIsUpperFunction_Set() error {
	var err error
	keyvalIsUpperFunction_Once.Do(func() {
		keyvalIsUpperFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_is_upper")
	})
	return err
}

// KeyvalIsUpper is a representation of the C type gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalIsUpperFunction_Set()
	if err == nil {
		ret = keyvalIsUpperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keyvalNameFunction *gi.Function
var keyvalNameFunction_Once sync.Once

func keyvalNameFunction_Set() error {
	var err error
	keyvalNameFunction_Once.Do(func() {
		keyvalNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_name")
	})
	return err
}

// KeyvalName is a representation of the C type gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalNameFunction_Set()
	if err == nil {
		ret = keyvalNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var keyvalToLowerFunction *gi.Function
var keyvalToLowerFunction_Once sync.Once

func keyvalToLowerFunction_Set() error {
	var err error
	keyvalToLowerFunction_Once.Do(func() {
		keyvalToLowerFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_lower")
	})
	return err
}

// KeyvalToLower is a representation of the C type gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToLowerFunction_Set()
	if err == nil {
		ret = keyvalToLowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUnicodeFunction *gi.Function
var keyvalToUnicodeFunction_Once sync.Once

func keyvalToUnicodeFunction_Set() error {
	var err error
	keyvalToUnicodeFunction_Once.Do(func() {
		keyvalToUnicodeFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_unicode")
	})
	return err
}

// KeyvalToUnicode is a representation of the C type gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUnicodeFunction_Set()
	if err == nil {
		ret = keyvalToUnicodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUpperFunction *gi.Function
var keyvalToUpperFunction_Once sync.Once

func keyvalToUpperFunction_Set() error {
	var err error
	keyvalToUpperFunction_Once.Do(func() {
		keyvalToUpperFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_upper")
	})
	return err
}

// KeyvalToUpper is a representation of the C type gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUpperFunction_Set()
	if err == nil {
		ret = keyvalToUpperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var listVisualsFunction *gi.Function
var listVisualsFunction_Once sync.Once

func listVisualsFunction_Set() error {
	var err error
	listVisualsFunction_Once.Do(func() {
		listVisualsFunction, err = gi.FunctionInvokerNew("Gdk", "list_visuals")
	})
	return err
}

// ListVisuals is a representation of the C type gdk_list_visuals.
func ListVisuals() *glib.List {

	var ret gi.Argument

	err := listVisualsFunction_Set()
	if err == nil {
		ret = listVisualsFunction.Invoke(nil, nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var notifyStartupCompleteFunction *gi.Function
var notifyStartupCompleteFunction_Once sync.Once

func notifyStartupCompleteFunction_Set() error {
	var err error
	notifyStartupCompleteFunction_Once.Do(func() {
		notifyStartupCompleteFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	})
	return err
}

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {

	err := notifyStartupCompleteFunction_Set()
	if err == nil {
		notifyStartupCompleteFunction.Invoke(nil, nil)
	}

	return
}

var notifyStartupCompleteWithIdFunction *gi.Function
var notifyStartupCompleteWithIdFunction_Once sync.Once

func notifyStartupCompleteWithIdFunction_Set() error {
	var err error
	notifyStartupCompleteWithIdFunction_Once.Do(func() {
		notifyStartupCompleteWithIdFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete_with_id")
	})
	return err
}

// NotifyStartupCompleteWithId is a representation of the C type gdk_notify_startup_complete_with_id.
func NotifyStartupCompleteWithId(startupId string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(startupId)

	err := notifyStartupCompleteWithIdFunction_Set()
	if err == nil {
		notifyStartupCompleteWithIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var offscreenWindowGetEmbedderFunction *gi.Function
var offscreenWindowGetEmbedderFunction_Once sync.Once

func offscreenWindowGetEmbedderFunction_Set() error {
	var err error
	offscreenWindowGetEmbedderFunction_Once.Do(func() {
		offscreenWindowGetEmbedderFunction, err = gi.FunctionInvokerNew("Gdk", "offscreen_window_get_embedder")
	})
	return err
}

// OffscreenWindowGetEmbedder is a representation of the C type gdk_offscreen_window_get_embedder.
func OffscreenWindowGetEmbedder(window *Window) *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native())

	var ret gi.Argument

	err := offscreenWindowGetEmbedderFunction_Set()
	if err == nil {
		ret = offscreenWindowGetEmbedderFunction.Invoke(inArgs[:], nil)
	}

	retGo := WindowNewFromNative(ret.Pointer())

	return retGo
}

var offscreenWindowGetSurfaceFunction *gi.Function
var offscreenWindowGetSurfaceFunction_Once sync.Once

func offscreenWindowGetSurfaceFunction_Set() error {
	var err error
	offscreenWindowGetSurfaceFunction_Once.Do(func() {
		offscreenWindowGetSurfaceFunction, err = gi.FunctionInvokerNew("Gdk", "offscreen_window_get_surface")
	})
	return err
}

// OffscreenWindowGetSurface is a representation of the C type gdk_offscreen_window_get_surface.
func OffscreenWindowGetSurface(window *Window) *cairo.Surface {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native())

	var ret gi.Argument

	err := offscreenWindowGetSurfaceFunction_Set()
	if err == nil {
		ret = offscreenWindowGetSurfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.SurfaceNewFromNative(ret.Pointer())

	return retGo
}

var offscreenWindowSetEmbedderFunction *gi.Function
var offscreenWindowSetEmbedderFunction_Once sync.Once

func offscreenWindowSetEmbedderFunction_Set() error {
	var err error
	offscreenWindowSetEmbedderFunction_Once.Do(func() {
		offscreenWindowSetEmbedderFunction, err = gi.FunctionInvokerNew("Gdk", "offscreen_window_set_embedder")
	})
	return err
}

// OffscreenWindowSetEmbedder is a representation of the C type gdk_offscreen_window_set_embedder.
func OffscreenWindowSetEmbedder(window *Window, embedder *Window) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(embedder.Native())

	err := offscreenWindowSetEmbedderFunction_Set()
	if err == nil {
		offscreenWindowSetEmbedderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pangoContextGetFunction *gi.Function
var pangoContextGetFunction_Once sync.Once

func pangoContextGetFunction_Set() error {
	var err error
	pangoContextGetFunction_Once.Do(func() {
		pangoContextGetFunction, err = gi.FunctionInvokerNew("Gdk", "pango_context_get")
	})
	return err
}

// PangoContextGet is a representation of the C type gdk_pango_context_get.
func PangoContextGet() *pango.Context {

	var ret gi.Argument

	err := pangoContextGetFunction_Set()
	if err == nil {
		ret = pangoContextGetFunction.Invoke(nil, nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

var pangoContextGetForDisplayFunction *gi.Function
var pangoContextGetForDisplayFunction_Once sync.Once

func pangoContextGetForDisplayFunction_Set() error {
	var err error
	pangoContextGetForDisplayFunction_Once.Do(func() {
		pangoContextGetForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "pango_context_get_for_display")
	})
	return err
}

// PangoContextGetForDisplay is a representation of the C type gdk_pango_context_get_for_display.
func PangoContextGetForDisplay(display *Display) *pango.Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(display.Native())

	var ret gi.Argument

	err := pangoContextGetForDisplayFunction_Set()
	if err == nil {
		ret = pangoContextGetForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

var pangoContextGetForScreenFunction *gi.Function
var pangoContextGetForScreenFunction_Once sync.Once

func pangoContextGetForScreenFunction_Set() error {
	var err error
	pangoContextGetForScreenFunction_Once.Do(func() {
		pangoContextGetForScreenFunction, err = gi.FunctionInvokerNew("Gdk", "pango_context_get_for_screen")
	})
	return err
}

// PangoContextGetForScreen is a representation of the C type gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(screen.Native())

	var ret gi.Argument

	err := pangoContextGetForScreenFunction_Set()
	if err == nil {
		ret = pangoContextGetForScreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

var pangoLayoutGetClipRegionFunction *gi.Function
var pangoLayoutGetClipRegionFunction_Once sync.Once

func pangoLayoutGetClipRegionFunction_Set() error {
	var err error
	pangoLayoutGetClipRegionFunction_Once.Do(func() {
		pangoLayoutGetClipRegionFunction, err = gi.FunctionInvokerNew("Gdk", "pango_layout_get_clip_region")
	})
	return err
}

// PangoLayoutGetClipRegion is a representation of the C type gdk_pango_layout_get_clip_region.
func PangoLayoutGetClipRegion(layout *pango.Layout, xOrigin int32, yOrigin int32, indexRanges int32, nRanges int32) *cairo.Region {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(layout.Native())
	inArgs[1].SetInt32(xOrigin)
	inArgs[2].SetInt32(yOrigin)
	inArgs[3].SetInt32(indexRanges)
	inArgs[4].SetInt32(nRanges)

	var ret gi.Argument

	err := pangoLayoutGetClipRegionFunction_Set()
	if err == nil {
		ret = pangoLayoutGetClipRegionFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.RegionNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gdk_pango_layout_line_get_clip_region' : parameter 'index_ranges' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_parse_args' : parameter 'argv' of type 'nil' not supported

var pixbufGetFromSurfaceFunction *gi.Function
var pixbufGetFromSurfaceFunction_Once sync.Once

func pixbufGetFromSurfaceFunction_Set() error {
	var err error
	pixbufGetFromSurfaceFunction_Once.Do(func() {
		pixbufGetFromSurfaceFunction, err = gi.FunctionInvokerNew("Gdk", "pixbuf_get_from_surface")
	})
	return err
}

// PixbufGetFromSurface is a representation of the C type gdk_pixbuf_get_from_surface.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(surface.Native())
	inArgs[1].SetInt32(srcX)
	inArgs[2].SetInt32(srcY)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	var ret gi.Argument

	err := pixbufGetFromSurfaceFunction_Set()
	if err == nil {
		ret = pixbufGetFromSurfaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufGetFromWindowFunction *gi.Function
var pixbufGetFromWindowFunction_Once sync.Once

func pixbufGetFromWindowFunction_Set() error {
	var err error
	pixbufGetFromWindowFunction_Once.Do(func() {
		pixbufGetFromWindowFunction, err = gi.FunctionInvokerNew("Gdk", "pixbuf_get_from_window")
	})
	return err
}

// PixbufGetFromWindow is a representation of the C type gdk_pixbuf_get_from_window.
func PixbufGetFromWindow(window *Window, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetInt32(srcX)
	inArgs[2].SetInt32(srcY)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	var ret gi.Argument

	err := pixbufGetFromWindowFunction_Set()
	if err == nil {
		ret = pixbufGetFromWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gdk_pointer_grab' : parameter 'event_mask' of type 'EventMask' not supported

var pointerIsGrabbedFunction *gi.Function
var pointerIsGrabbedFunction_Once sync.Once

func pointerIsGrabbedFunction_Set() error {
	var err error
	pointerIsGrabbedFunction_Once.Do(func() {
		pointerIsGrabbedFunction, err = gi.FunctionInvokerNew("Gdk", "pointer_is_grabbed")
	})
	return err
}

// PointerIsGrabbed is a representation of the C type gdk_pointer_is_grabbed.
func PointerIsGrabbed() bool {

	var ret gi.Argument

	err := pointerIsGrabbedFunction_Set()
	if err == nil {
		ret = pointerIsGrabbedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pointerUngrabFunction *gi.Function
var pointerUngrabFunction_Once sync.Once

func pointerUngrabFunction_Set() error {
	var err error
	pointerUngrabFunction_Once.Do(func() {
		pointerUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "pointer_ungrab")
	})
	return err
}

// PointerUngrab is a representation of the C type gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := pointerUngrabFunction_Set()
	if err == nil {
		pointerUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var preParseLibgtkOnlyFunction *gi.Function
var preParseLibgtkOnlyFunction_Once sync.Once

func preParseLibgtkOnlyFunction_Set() error {
	var err error
	preParseLibgtkOnlyFunction_Once.Do(func() {
		preParseLibgtkOnlyFunction, err = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	})
	return err
}

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {

	err := preParseLibgtkOnlyFunction_Set()
	if err == nil {
		preParseLibgtkOnlyFunction.Invoke(nil, nil)
	}

	return
}

var propertyChangeFunction *gi.Function
var propertyChangeFunction_Once sync.Once

func propertyChangeFunction_Set() error {
	var err error
	propertyChangeFunction_Once.Do(func() {
		propertyChangeFunction, err = gi.FunctionInvokerNew("Gdk", "property_change")
	})
	return err
}

// PropertyChange is a representation of the C type gdk_property_change.
func PropertyChange(window *Window, property *Atom, type_ *Atom, format int32, mode PropMode, data uint8, nelements int32) {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(property.Native())
	inArgs[2].SetPointer(type_.Native())
	inArgs[3].SetInt32(format)
	inArgs[4].SetInt32(int32(mode))
	inArgs[5].SetUint8(data)
	inArgs[6].SetInt32(nelements)

	err := propertyChangeFunction_Set()
	if err == nil {
		propertyChangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var propertyDeleteFunction *gi.Function
var propertyDeleteFunction_Once sync.Once

func propertyDeleteFunction_Set() error {
	var err error
	propertyDeleteFunction_Once.Do(func() {
		propertyDeleteFunction, err = gi.FunctionInvokerNew("Gdk", "property_delete")
	})
	return err
}

// PropertyDelete is a representation of the C type gdk_property_delete.
func PropertyDelete(window *Window, property *Atom) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(window.Native())
	inArgs[1].SetPointer(property.Native())

	err := propertyDeleteFunction_Set()
	if err == nil {
		propertyDeleteFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_property_get' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_query_depths' : parameter 'depths' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_query_visual_types' : parameter 'visual_types' of type 'nil' not supported

var selectionConvertFunction *gi.Function
var selectionConvertFunction_Once sync.Once

func selectionConvertFunction_Set() error {
	var err error
	selectionConvertFunction_Once.Do(func() {
		selectionConvertFunction, err = gi.FunctionInvokerNew("Gdk", "selection_convert")
	})
	return err
}

// SelectionConvert is a representation of the C type gdk_selection_convert.
func SelectionConvert(requestor *Window, selection *Atom, target *Atom, time uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(requestor.Native())
	inArgs[1].SetPointer(selection.Native())
	inArgs[2].SetPointer(target.Native())
	inArgs[3].SetUint32(time)

	err := selectionConvertFunction_Set()
	if err == nil {
		selectionConvertFunction.Invoke(inArgs[:], nil)
	}

	return
}

var selectionOwnerGetFunction *gi.Function
var selectionOwnerGetFunction_Once sync.Once

func selectionOwnerGetFunction_Set() error {
	var err error
	selectionOwnerGetFunction_Once.Do(func() {
		selectionOwnerGetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_get")
	})
	return err
}

// SelectionOwnerGet is a representation of the C type gdk_selection_owner_get.
func SelectionOwnerGet(selection *Atom) *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(selection.Native())

	var ret gi.Argument

	err := selectionOwnerGetFunction_Set()
	if err == nil {
		ret = selectionOwnerGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := WindowNewFromNative(ret.Pointer())

	return retGo
}

var selectionOwnerGetForDisplayFunction *gi.Function
var selectionOwnerGetForDisplayFunction_Once sync.Once

func selectionOwnerGetForDisplayFunction_Set() error {
	var err error
	selectionOwnerGetForDisplayFunction_Once.Do(func() {
		selectionOwnerGetForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_get_for_display")
	})
	return err
}

// SelectionOwnerGetForDisplay is a representation of the C type gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection *Atom) *Window {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(display.Native())
	inArgs[1].SetPointer(selection.Native())

	var ret gi.Argument

	err := selectionOwnerGetForDisplayFunction_Set()
	if err == nil {
		ret = selectionOwnerGetForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := WindowNewFromNative(ret.Pointer())

	return retGo
}

var selectionOwnerSetFunction *gi.Function
var selectionOwnerSetFunction_Once sync.Once

func selectionOwnerSetFunction_Set() error {
	var err error
	selectionOwnerSetFunction_Once.Do(func() {
		selectionOwnerSetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_set")
	})
	return err
}

// SelectionOwnerSet is a representation of the C type gdk_selection_owner_set.
func SelectionOwnerSet(owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(owner.Native())
	inArgs[1].SetPointer(selection.Native())
	inArgs[2].SetUint32(time)
	inArgs[3].SetBoolean(sendEvent)

	var ret gi.Argument

	err := selectionOwnerSetFunction_Set()
	if err == nil {
		ret = selectionOwnerSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionOwnerSetForDisplayFunction *gi.Function
var selectionOwnerSetForDisplayFunction_Once sync.Once

func selectionOwnerSetForDisplayFunction_Set() error {
	var err error
	selectionOwnerSetForDisplayFunction_Once.Do(func() {
		selectionOwnerSetForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_set_for_display")
	})
	return err
}

// SelectionOwnerSetForDisplay is a representation of the C type gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(display.Native())
	inArgs[1].SetPointer(owner.Native())
	inArgs[2].SetPointer(selection.Native())
	inArgs[3].SetUint32(time)
	inArgs[4].SetBoolean(sendEvent)

	var ret gi.Argument

	err := selectionOwnerSetForDisplayFunction_Set()
	if err == nil {
		ret = selectionOwnerSetForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionPropertyGetFunction *gi.Function
var selectionPropertyGetFunction_Once sync.Once

func selectionPropertyGetFunction_Set() error {
	var err error
	selectionPropertyGetFunction_Once.Do(func() {
		selectionPropertyGetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_property_get")
	})
	return err
}

// SelectionPropertyGet is a representation of the C type gdk_selection_property_get.
func SelectionPropertyGet(requestor *Window, data uint8, propType *Atom, propFormat int32) int32 {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(requestor.Native())
	inArgs[1].SetUint8(data)
	inArgs[2].SetPointer(propType.Native())
	inArgs[3].SetInt32(propFormat)

	var ret gi.Argument

	err := selectionPropertyGetFunction_Set()
	if err == nil {
		ret = selectionPropertyGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var selectionSendNotifyFunction *gi.Function
var selectionSendNotifyFunction_Once sync.Once

func selectionSendNotifyFunction_Set() error {
	var err error
	selectionSendNotifyFunction_Once.Do(func() {
		selectionSendNotifyFunction, err = gi.FunctionInvokerNew("Gdk", "selection_send_notify")
	})
	return err
}

// SelectionSendNotify is a representation of the C type gdk_selection_send_notify.
func SelectionSendNotify(requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(requestor.Native())
	inArgs[1].SetPointer(selection.Native())
	inArgs[2].SetPointer(target.Native())
	inArgs[3].SetPointer(property.Native())
	inArgs[4].SetUint32(time)

	err := selectionSendNotifyFunction_Set()
	if err == nil {
		selectionSendNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var selectionSendNotifyForDisplayFunction *gi.Function
var selectionSendNotifyForDisplayFunction_Once sync.Once

func selectionSendNotifyForDisplayFunction_Set() error {
	var err error
	selectionSendNotifyForDisplayFunction_Once.Do(func() {
		selectionSendNotifyForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_send_notify_for_display")
	})
	return err
}

// SelectionSendNotifyForDisplay is a representation of the C type gdk_selection_send_notify_for_display.
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(display.Native())
	inArgs[1].SetPointer(requestor.Native())
	inArgs[2].SetPointer(selection.Native())
	inArgs[3].SetPointer(target.Native())
	inArgs[4].SetPointer(property.Native())
	inArgs[5].SetUint32(time)

	err := selectionSendNotifyForDisplayFunction_Set()
	if err == nil {
		selectionSendNotifyForDisplayFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setAllowedBackendsFunction *gi.Function
var setAllowedBackendsFunction_Once sync.Once

func setAllowedBackendsFunction_Set() error {
	var err error
	setAllowedBackendsFunction_Once.Do(func() {
		setAllowedBackendsFunction, err = gi.FunctionInvokerNew("Gdk", "set_allowed_backends")
	})
	return err
}

// SetAllowedBackends is a representation of the C type gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(backends)

	err := setAllowedBackendsFunction_Set()
	if err == nil {
		setAllowedBackendsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setDoubleClickTimeFunction *gi.Function
var setDoubleClickTimeFunction_Once sync.Once

func setDoubleClickTimeFunction_Set() error {
	var err error
	setDoubleClickTimeFunction_Once.Do(func() {
		setDoubleClickTimeFunction, err = gi.FunctionInvokerNew("Gdk", "set_double_click_time")
	})
	return err
}

// SetDoubleClickTime is a representation of the C type gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(msec)

	err := setDoubleClickTimeFunction_Set()
	if err == nil {
		setDoubleClickTimeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setProgramClassFunction *gi.Function
var setProgramClassFunction_Once sync.Once

func setProgramClassFunction_Set() error {
	var err error
	setProgramClassFunction_Once.Do(func() {
		setProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "set_program_class")
	})
	return err
}

// SetProgramClass is a representation of the C type gdk_set_program_class.
func SetProgramClass(programClass string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(programClass)

	err := setProgramClassFunction_Set()
	if err == nil {
		setProgramClassFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setShowEventsFunction *gi.Function
var setShowEventsFunction_Once sync.Once

func setShowEventsFunction_Set() error {
	var err error
	setShowEventsFunction_Once.Do(func() {
		setShowEventsFunction, err = gi.FunctionInvokerNew("Gdk", "set_show_events")
	})
	return err
}

// SetShowEvents is a representation of the C type gdk_set_show_events.
func SetShowEvents(showEvents bool) {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(showEvents)

	err := setShowEventsFunction_Set()
	if err == nil {
		setShowEventsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingGetFunction *gi.Function
var settingGetFunction_Once sync.Once

func settingGetFunction_Set() error {
	var err error
	settingGetFunction_Once.Do(func() {
		settingGetFunction, err = gi.FunctionInvokerNew("Gdk", "setting_get")
	})
	return err
}

// SettingGet is a representation of the C type gdk_setting_get.
func SettingGet(name string, value *gobject.Value) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(value.Native())

	var ret gi.Argument

	err := settingGetFunction_Set()
	if err == nil {
		ret = settingGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_synthesize_window_state' : parameter 'unset_flags' of type 'WindowState' not supported

var testRenderSyncFunction *gi.Function
var testRenderSyncFunction_Once sync.Once

func testRenderSyncFunction_Set() error {
	var err error
	testRenderSyncFunction_Once.Do(func() {
		testRenderSyncFunction, err = gi.FunctionInvokerNew("Gdk", "test_render_sync")
	})
	return err
}

// TestRenderSync is a representation of the C type gdk_test_render_sync.
func TestRenderSync(window *Window) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native())

	err := testRenderSyncFunction_Set()
	if err == nil {
		testRenderSyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_test_simulate_button' : parameter 'modifiers' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_key' : parameter 'modifiers' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_text_property_to_utf8_list_for_display' : parameter 'text' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

var threadsEnterFunction *gi.Function
var threadsEnterFunction_Once sync.Once

func threadsEnterFunction_Set() error {
	var err error
	threadsEnterFunction_Once.Do(func() {
		threadsEnterFunction, err = gi.FunctionInvokerNew("Gdk", "threads_enter")
	})
	return err
}

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {

	err := threadsEnterFunction_Set()
	if err == nil {
		threadsEnterFunction.Invoke(nil, nil)
	}

	return
}

var threadsInitFunction *gi.Function
var threadsInitFunction_Once sync.Once

func threadsInitFunction_Set() error {
	var err error
	threadsInitFunction_Once.Do(func() {
		threadsInitFunction, err = gi.FunctionInvokerNew("Gdk", "threads_init")
	})
	return err
}

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {

	err := threadsInitFunction_Set()
	if err == nil {
		threadsInitFunction.Invoke(nil, nil)
	}

	return
}

var threadsLeaveFunction *gi.Function
var threadsLeaveFunction_Once sync.Once

func threadsLeaveFunction_Set() error {
	var err error
	threadsLeaveFunction_Once.Do(func() {
		threadsLeaveFunction, err = gi.FunctionInvokerNew("Gdk", "threads_leave")
	})
	return err
}

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {

	err := threadsLeaveFunction_Set()
	if err == nil {
		threadsLeaveFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : parameter 'enter_fn' of type 'GObject.Callback' not supported

var unicodeToKeyvalFunction *gi.Function
var unicodeToKeyvalFunction_Once sync.Once

func unicodeToKeyvalFunction_Set() error {
	var err error
	unicodeToKeyvalFunction_Once.Do(func() {
		unicodeToKeyvalFunction, err = gi.FunctionInvokerNew("Gdk", "unicode_to_keyval")
	})
	return err
}

// UnicodeToKeyval is a representation of the C type gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(wc)

	var ret gi.Argument

	err := unicodeToKeyvalFunction_Set()
	if err == nil {
		ret = unicodeToKeyvalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var utf8ToStringTargetFunction *gi.Function
var utf8ToStringTargetFunction_Once sync.Once

func utf8ToStringTargetFunction_Set() error {
	var err error
	utf8ToStringTargetFunction_Once.Do(func() {
		utf8ToStringTargetFunction, err = gi.FunctionInvokerNew("Gdk", "utf8_to_string_target")
	})
	return err
}

// Utf8ToStringTarget is a representation of the C type gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := utf8ToStringTargetFunction_Set()
	if err == nil {
		ret = utf8ToStringTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}
