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

// Component is a wrapper around the C record AtkComponent.
type Component struct {
	native *C.AtkComponent
}

func ComponentNewFromC(u unsafe.Pointer) *Component {
	c := (*C.AtkComponent)(u)
	if c == nil {
		return nil
	}

	g := &Component{native: c}

	return g
}

func (recv *Component) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Document is a wrapper around the C record AtkDocument.
type Document struct {
	native *C.AtkDocument
}

func DocumentNewFromC(u unsafe.Pointer) *Document {
	c := (*C.AtkDocument)(u)
	if c == nil {
		return nil
	}

	g := &Document{native: c}

	return g
}

func (recv *Document) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EditableText is a wrapper around the C record AtkEditableText.
type EditableText struct {
	native *C.AtkEditableText
}

func EditableTextNewFromC(u unsafe.Pointer) *EditableText {
	c := (*C.AtkEditableText)(u)
	if c == nil {
		return nil
	}

	g := &EditableText{native: c}

	return g
}

func (recv *EditableText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HyperlinkImpl is a wrapper around the C record AtkHyperlinkImpl.
type HyperlinkImpl struct {
	native *C.AtkHyperlinkImpl
}

func HyperlinkImplNewFromC(u unsafe.Pointer) *HyperlinkImpl {
	c := (*C.AtkHyperlinkImpl)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkImpl{native: c}

	return g
}

func (recv *HyperlinkImpl) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hypertext is a wrapper around the C record AtkHypertext.
type Hypertext struct {
	native *C.AtkHypertext
}

func HypertextNewFromC(u unsafe.Pointer) *Hypertext {
	c := (*C.AtkHypertext)(u)
	if c == nil {
		return nil
	}

	g := &Hypertext{native: c}

	return g
}

func (recv *Hypertext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Image is a wrapper around the C record AtkImage.
type Image struct {
	native *C.AtkImage
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	c := (*C.AtkImage)(u)
	if c == nil {
		return nil
	}

	g := &Image{native: c}

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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

// Selection is a wrapper around the C record AtkSelection.
type Selection struct {
	native *C.AtkSelection
}

func SelectionNewFromC(u unsafe.Pointer) *Selection {
	c := (*C.AtkSelection)(u)
	if c == nil {
		return nil
	}

	g := &Selection{native: c}

	return g
}

func (recv *Selection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StreamableContent is a wrapper around the C record AtkStreamableContent.
type StreamableContent struct {
	native *C.AtkStreamableContent
}

func StreamableContentNewFromC(u unsafe.Pointer) *StreamableContent {
	c := (*C.AtkStreamableContent)(u)
	if c == nil {
		return nil
	}

	g := &StreamableContent{native: c}

	return g
}

func (recv *StreamableContent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Table is a wrapper around the C record AtkTable.
type Table struct {
	native *C.AtkTable
}

func TableNewFromC(u unsafe.Pointer) *Table {
	c := (*C.AtkTable)(u)
	if c == nil {
		return nil
	}

	g := &Table{native: c}

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableCell is a wrapper around the C record AtkTableCell.
type TableCell struct {
	native *C.AtkTableCell
}

func TableCellNewFromC(u unsafe.Pointer) *TableCell {
	c := (*C.AtkTableCell)(u)
	if c == nil {
		return nil
	}

	g := &TableCell{native: c}

	return g
}

func (recv *TableCell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Text is a wrapper around the C record AtkText.
type Text struct {
	native *C.AtkText
}

func TextNewFromC(u unsafe.Pointer) *Text {
	c := (*C.AtkText)(u)
	if c == nil {
		return nil
	}

	g := &Text{native: c}

	return g
}

func (recv *Text) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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
