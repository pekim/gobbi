// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// Action is a wrapper around the C record AtkAction.
type Action struct {
	native unsafe.Pointer
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	if u == nil {
		return nil
	}

	g := &Action{native: u}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component is a wrapper around the C record AtkComponent.
type Component struct {
	native unsafe.Pointer
}

func ComponentNewFromC(u unsafe.Pointer) *Component {
	if u == nil {
		return nil
	}

	g := &Component{native: u}

	return g
}

func (recv *Component) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Document is a wrapper around the C record AtkDocument.
type Document struct {
	native unsafe.Pointer
}

func DocumentNewFromC(u unsafe.Pointer) *Document {
	if u == nil {
		return nil
	}

	g := &Document{native: u}

	return g
}

func (recv *Document) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EditableText is a wrapper around the C record AtkEditableText.
type EditableText struct {
	native unsafe.Pointer
}

func EditableTextNewFromC(u unsafe.Pointer) *EditableText {
	if u == nil {
		return nil
	}

	g := &EditableText{native: u}

	return g
}

func (recv *EditableText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HyperlinkImpl is a wrapper around the C record AtkHyperlinkImpl.
type HyperlinkImpl struct {
	native unsafe.Pointer
}

func HyperlinkImplNewFromC(u unsafe.Pointer) *HyperlinkImpl {
	if u == nil {
		return nil
	}

	g := &HyperlinkImpl{native: u}

	return g
}

func (recv *HyperlinkImpl) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hypertext is a wrapper around the C record AtkHypertext.
type Hypertext struct {
	native unsafe.Pointer
}

func HypertextNewFromC(u unsafe.Pointer) *Hypertext {
	if u == nil {
		return nil
	}

	g := &Hypertext{native: u}

	return g
}

func (recv *Hypertext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Image is a wrapper around the C record AtkImage.
type Image struct {
	native unsafe.Pointer
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	if u == nil {
		return nil
	}

	g := &Image{native: u}

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImplementorIface is a wrapper around the C record AtkImplementorIface.
type ImplementorIface struct {
	native unsafe.Pointer
}

func ImplementorIfaceNewFromC(u unsafe.Pointer) *ImplementorIface {
	if u == nil {
		return nil
	}

	g := &ImplementorIface{native: u}

	return g
}

func (recv *ImplementorIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Selection is a wrapper around the C record AtkSelection.
type Selection struct {
	native unsafe.Pointer
}

func SelectionNewFromC(u unsafe.Pointer) *Selection {
	if u == nil {
		return nil
	}

	g := &Selection{native: u}

	return g
}

func (recv *Selection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StreamableContent is a wrapper around the C record AtkStreamableContent.
type StreamableContent struct {
	native unsafe.Pointer
}

func StreamableContentNewFromC(u unsafe.Pointer) *StreamableContent {
	if u == nil {
		return nil
	}

	g := &StreamableContent{native: u}

	return g
}

func (recv *StreamableContent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Table is a wrapper around the C record AtkTable.
type Table struct {
	native unsafe.Pointer
}

func TableNewFromC(u unsafe.Pointer) *Table {
	if u == nil {
		return nil
	}

	g := &Table{native: u}

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableCell is a wrapper around the C record AtkTableCell.
type TableCell struct {
	native unsafe.Pointer
}

func TableCellNewFromC(u unsafe.Pointer) *TableCell {
	if u == nil {
		return nil
	}

	g := &TableCell{native: u}

	return g
}

func (recv *TableCell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Text is a wrapper around the C record AtkText.
type Text struct {
	native unsafe.Pointer
}

func TextNewFromC(u unsafe.Pointer) *Text {
	if u == nil {
		return nil
	}

	g := &Text{native: u}

	return g
}

func (recv *Text) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Value is a wrapper around the C record AtkValue.
type Value struct {
	native unsafe.Pointer
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	if u == nil {
		return nil
	}

	g := &Value{native: u}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window is a wrapper around the C record AtkWindow.
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
