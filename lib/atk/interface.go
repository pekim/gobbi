// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Action is a wrapper around the C record AtkAction.
type Action struct {
	native *C.AtkAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.AtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_action_do_action

// Blacklisted : atk_action_get_description

// Blacklisted : atk_action_get_keybinding

// Blacklisted : atk_action_get_localized_name

// Blacklisted : atk_action_get_n_actions

// Blacklisted : atk_action_get_name

// Blacklisted : atk_action_set_description

// Blacklisted : AtkComponent

// Blacklisted : AtkDocument

// Blacklisted : AtkEditableText

// Blacklisted : AtkHyperlinkImpl

// Blacklisted : AtkHypertext

// Blacklisted : AtkImage

// ImplementorIface is a wrapper around the C record AtkImplementorIface.
type ImplementorIface struct {
	native *C.AtkImplementorIface
}

func ImplementorIfaceNewFromC(u unsafe.Pointer) *ImplementorIface {
	c := (*C.AtkImplementorIface)(u)
	if c == nil {
		return nil
	}

	g := &ImplementorIface{native: c}

	return g
}

func (recv *ImplementorIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ImplementorIface with another ImplementorIface, and returns true if they represent the same GObject.
func (recv *ImplementorIface) Equals(other *ImplementorIface) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : AtkSelection

// Blacklisted : AtkStreamableContent

// Blacklisted : AtkTable

// Blacklisted : AtkTableCell

// Blacklisted : AtkText

// Value is a wrapper around the C record AtkValue.
type Value struct {
	native *C.AtkValue
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.AtkValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_value_get_current_value

// Blacklisted : atk_value_get_maximum_value

// Blacklisted : atk_value_get_minimum_value

// Blacklisted : atk_value_set_current_value

// Window is a wrapper around the C record AtkWindow.
type Window struct {
	native *C.AtkWindow
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.AtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}
