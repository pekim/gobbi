// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// FileChooserNative is a wrapper around the C record GtkFileChooserNative.
type FileChooserNative struct {
	native unsafe.Pointer
}

func FileChooserNativeNewFromC(u unsafe.Pointer) *FileChooserNative {
	if u == nil {
		return nil
	}

	g := &FileChooserNative{native: u}

	return g
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NativeDialog is a wrapper around the C record GtkNativeDialog.
type NativeDialog struct {
	native unsafe.Pointer
	// parent_instance : record
}

func NativeDialogNewFromC(u unsafe.Pointer) *NativeDialog {
	if u == nil {
		return nil
	}

	g := &NativeDialog{native: u}

	return g
}

func (recv *NativeDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PadController is a wrapper around the C record GtkPadController.
type PadController struct {
	native unsafe.Pointer
}

func PadControllerNewFromC(u unsafe.Pointer) *PadController {
	if u == nil {
		return nil
	}

	g := &PadController{native: u}

	return g
}

func (recv *PadController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutLabel is a wrapper around the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native unsafe.Pointer
}

func ShortcutLabelNewFromC(u unsafe.Pointer) *ShortcutLabel {
	if u == nil {
		return nil
	}

	g := &ShortcutLabel{native: u}

	return g
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsGroup is a wrapper around the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native unsafe.Pointer
}

func ShortcutsGroupNewFromC(u unsafe.Pointer) *ShortcutsGroup {
	if u == nil {
		return nil
	}

	g := &ShortcutsGroup{native: u}

	return g
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsSection is a wrapper around the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native unsafe.Pointer
}

func ShortcutsSectionNewFromC(u unsafe.Pointer) *ShortcutsSection {
	if u == nil {
		return nil
	}

	g := &ShortcutsSection{native: u}

	return g
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsShortcut is a wrapper around the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native unsafe.Pointer
}

func ShortcutsShortcutNewFromC(u unsafe.Pointer) *ShortcutsShortcut {
	if u == nil {
		return nil
	}

	g := &ShortcutsShortcut{native: u}

	return g
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ShortcutsWindow is a wrapper around the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native unsafe.Pointer
	// window : record
}

func ShortcutsWindowNewFromC(u unsafe.Pointer) *ShortcutsWindow {
	if u == nil {
		return nil
	}

	g := &ShortcutsWindow{native: u}

	return g
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
