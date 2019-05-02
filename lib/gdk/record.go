// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Equals compares this Atom with another Atom, and returns true if they represent the same GObject.
func (recv *Atom) Equals(other *Atom) bool {
	return other.ToC() == recv.ToC()
}

// AtomIntern is a wrapper around the C function gdk_atom_intern.
func AtomIntern(atomName string, onlyIfExists bool) Atom {
	c_atom_name := C.CString(atomName)
	defer C.free(unsafe.Pointer(c_atom_name))

	c_only_if_exists :=
		boolToGboolean(onlyIfExists)

	retC := C.gdk_atom_intern(c_atom_name, c_only_if_exists)
	retGo := *AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Name is a wrapper around the C function gdk_atom_name.
func (recv *Atom) Name() string {
	retC := C.gdk_atom_name((C.GdkAtom)(*recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Equals compares this Color with another Color, and returns true if they represent the same GObject.
func (recv *Color) Equals(other *Color) bool {
	return other.ToC() == recv.ToC()
}

// ColorParse is a wrapper around the C function gdk_color_parse.
func ColorParse(spec string) (bool, *Color) {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	var c_color C.GdkColor

	retC := C.gdk_color_parse(c_spec, &c_color)
	retGo := retC == C.TRUE

	color := ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Copy is a wrapper around the C function gdk_color_copy.
func (recv *Color) Copy() *Color {
	retC := C.gdk_color_copy((*C.GdkColor)(recv.native))
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function gdk_color_equal.
func (recv *Color) Equal(colorb *Color) bool {
	c_colorb := (*C.GdkColor)(C.NULL)
	if colorb != nil {
		c_colorb = (*C.GdkColor)(colorb.ToC())
	}

	retC := C.gdk_color_equal((*C.GdkColor)(recv.native), c_colorb)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gdk_color_free.
func (recv *Color) Free() {
	C.gdk_color_free((*C.GdkColor)(recv.native))

	return
}

// Hash is a wrapper around the C function gdk_color_hash.
func (recv *Color) Hash() uint32 {
	retC := C.gdk_color_hash((*C.GdkColor)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Equals compares this EventAny with another EventAny, and returns true if they represent the same GObject.
func (recv *EventAny) Equals(other *EventAny) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventButton with another EventButton, and returns true if they represent the same GObject.
func (recv *EventButton) Equals(other *EventButton) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventConfigure with another EventConfigure, and returns true if they represent the same GObject.
func (recv *EventConfigure) Equals(other *EventConfigure) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventCrossing with another EventCrossing, and returns true if they represent the same GObject.
func (recv *EventCrossing) Equals(other *EventCrossing) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventDND with another EventDND, and returns true if they represent the same GObject.
func (recv *EventDND) Equals(other *EventDND) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventExpose with another EventExpose, and returns true if they represent the same GObject.
func (recv *EventExpose) Equals(other *EventExpose) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventFocus with another EventFocus, and returns true if they represent the same GObject.
func (recv *EventFocus) Equals(other *EventFocus) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventKey with another EventKey, and returns true if they represent the same GObject.
func (recv *EventKey) Equals(other *EventKey) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventMotion with another EventMotion, and returns true if they represent the same GObject.
func (recv *EventMotion) Equals(other *EventMotion) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventProperty with another EventProperty, and returns true if they represent the same GObject.
func (recv *EventProperty) Equals(other *EventProperty) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventProximity with another EventProximity, and returns true if they represent the same GObject.
func (recv *EventProximity) Equals(other *EventProximity) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventScroll with another EventScroll, and returns true if they represent the same GObject.
func (recv *EventScroll) Equals(other *EventScroll) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventSelection with another EventSelection, and returns true if they represent the same GObject.
func (recv *EventSelection) Equals(other *EventSelection) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventSequence with another EventSequence, and returns true if they represent the same GObject.
func (recv *EventSequence) Equals(other *EventSequence) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventSetting with another EventSetting, and returns true if they represent the same GObject.
func (recv *EventSetting) Equals(other *EventSetting) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventTouch with another EventTouch, and returns true if they represent the same GObject.
func (recv *EventTouch) Equals(other *EventTouch) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventTouchpadPinch with another EventTouchpadPinch, and returns true if they represent the same GObject.
func (recv *EventTouchpadPinch) Equals(other *EventTouchpadPinch) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventTouchpadSwipe with another EventTouchpadSwipe, and returns true if they represent the same GObject.
func (recv *EventTouchpadSwipe) Equals(other *EventTouchpadSwipe) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventVisibility with another EventVisibility, and returns true if they represent the same GObject.
func (recv *EventVisibility) Equals(other *EventVisibility) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventWindowState with another EventWindowState, and returns true if they represent the same GObject.
func (recv *EventWindowState) Equals(other *EventWindowState) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this FrameClockClass with another FrameClockClass, and returns true if they represent the same GObject.
func (recv *FrameClockClass) Equals(other *FrameClockClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this FrameClockPrivate with another FrameClockPrivate, and returns true if they represent the same GObject.
func (recv *FrameClockPrivate) Equals(other *FrameClockPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this FrameTimings with another FrameTimings, and returns true if they represent the same GObject.
func (recv *FrameTimings) Equals(other *FrameTimings) bool {
	return other.ToC() == recv.ToC()
}

// GetFrameTime is a wrapper around the C function gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() int64 {
	retC := C.gdk_frame_timings_get_frame_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Equals compares this Geometry with another Geometry, and returns true if they represent the same GObject.
func (recv *Geometry) Equals(other *Geometry) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this KeymapKey with another KeymapKey, and returns true if they represent the same GObject.
func (recv *KeymapKey) Equals(other *KeymapKey) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Point with another Point, and returns true if they represent the same GObject.
func (recv *Point) Equals(other *Point) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this RGBA with another RGBA, and returns true if they represent the same GObject.
func (recv *RGBA) Equals(other *RGBA) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// Intersect is a wrapper around the C function gdk_rectangle_intersect.
func (recv *Rectangle) Intersect(src2 *Rectangle) (bool, *Rectangle) {
	c_src2 := (*C.GdkRectangle)(C.NULL)
	if src2 != nil {
		c_src2 = (*C.GdkRectangle)(src2.ToC())
	}

	var c_dest C.GdkRectangle

	retC := C.gdk_rectangle_intersect((*C.GdkRectangle)(recv.native), c_src2, &c_dest)
	retGo := retC == C.TRUE

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return retGo, dest
}

// Union is a wrapper around the C function gdk_rectangle_union.
func (recv *Rectangle) Union(src2 *Rectangle) *Rectangle {
	c_src2 := (*C.GdkRectangle)(C.NULL)
	if src2 != nil {
		c_src2 = (*C.GdkRectangle)(src2.ToC())
	}

	var c_dest C.GdkRectangle

	C.gdk_rectangle_union((*C.GdkRectangle)(recv.native), c_src2, &c_dest)

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return dest
}

// Equals compares this TimeCoord with another TimeCoord, and returns true if they represent the same GObject.
func (recv *TimeCoord) Equals(other *TimeCoord) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this WindowAttr with another WindowAttr, and returns true if they represent the same GObject.
func (recv *WindowAttr) Equals(other *WindowAttr) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this WindowClass with another WindowClass, and returns true if they represent the same GObject.
func (recv *WindowClass) Equals(other *WindowClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this WindowRedirect with another WindowRedirect, and returns true if they represent the same GObject.
func (recv *WindowRedirect) Equals(other *WindowRedirect) bool {
	return other.ToC() == recv.ToC()
}
