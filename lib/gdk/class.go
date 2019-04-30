// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// AppLaunchContext is a wrapper around the C record GdkAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
}

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	if u == nil {
		return nil
	}

	g := &AppLaunchContext{native: u}

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Cursor is a wrapper around the C record GdkCursor.
type Cursor struct {
	native unsafe.Pointer
}

func CursorNewFromC(u unsafe.Pointer) *Cursor {
	if u == nil {
		return nil
	}

	g := &Cursor{native: u}

	return g
}

func (recv *Cursor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_cursor_new : return type :

// Device is a wrapper around the C record GdkDevice.
type Device struct {
	native unsafe.Pointer
}

func DeviceNewFromC(u unsafe.Pointer) *Device {
	if u == nil {
		return nil
	}

	g := &Device{native: u}

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DeviceManager is a wrapper around the C record GdkDeviceManager.
type DeviceManager struct {
	native unsafe.Pointer
}

func DeviceManagerNewFromC(u unsafe.Pointer) *DeviceManager {
	if u == nil {
		return nil
	}

	g := &DeviceManager{native: u}

	return g
}

func (recv *DeviceManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Display is a wrapper around the C record GdkDisplay.
type Display struct {
	native unsafe.Pointer
}

func DisplayNewFromC(u unsafe.Pointer) *Display {
	if u == nil {
		return nil
	}

	g := &Display{native: u}

	return g
}

func (recv *Display) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DisplayManager is a wrapper around the C record GdkDisplayManager.
type DisplayManager struct {
	native unsafe.Pointer
}

func DisplayManagerNewFromC(u unsafe.Pointer) *DisplayManager {
	if u == nil {
		return nil
	}

	g := &DisplayManager{native: u}

	return g
}

func (recv *DisplayManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DragContext is a wrapper around the C record GdkDragContext.
type DragContext struct {
	native unsafe.Pointer
}

func DragContextNewFromC(u unsafe.Pointer) *DragContext {
	if u == nil {
		return nil
	}

	g := &DragContext{native: u}

	return g
}

func (recv *DragContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameClock is a wrapper around the C record GdkFrameClock.
type FrameClock struct {
	native unsafe.Pointer
}

func FrameClockNewFromC(u unsafe.Pointer) *FrameClock {
	if u == nil {
		return nil
	}

	g := &FrameClock{native: u}

	return g
}

func (recv *FrameClock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GLContext is a wrapper around the C record GdkGLContext.
type GLContext struct {
	native unsafe.Pointer
}

func GLContextNewFromC(u unsafe.Pointer) *GLContext {
	if u == nil {
		return nil
	}

	g := &GLContext{native: u}

	return g
}

func (recv *GLContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Keymap is a wrapper around the C record GdkKeymap.
type Keymap struct {
	native unsafe.Pointer
}

func KeymapNewFromC(u unsafe.Pointer) *Keymap {
	if u == nil {
		return nil
	}

	g := &Keymap{native: u}

	return g
}

func (recv *Keymap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Screen is a wrapper around the C record GdkScreen.
type Screen struct {
	native unsafe.Pointer
}

func ScreenNewFromC(u unsafe.Pointer) *Screen {
	if u == nil {
		return nil
	}

	g := &Screen{native: u}

	return g
}

func (recv *Screen) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Visual is a wrapper around the C record GdkVisual.
type Visual struct {
	native unsafe.Pointer
}

func VisualNewFromC(u unsafe.Pointer) *Visual {
	if u == nil {
		return nil
	}

	g := &Visual{native: u}

	return g
}

func (recv *Visual) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window is a wrapper around the C record GdkWindow.
type Window struct {
	native unsafe.Pointer
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	if u == nil {
		return nil
	}

	g := &Window{native: u}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_window_new : return type :
