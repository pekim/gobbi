// This is a generated file - DO NOT EDIT

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Actionable is a wrapper around the C record GtkActionable.
type Actionable struct {
	native *C.GtkActionable
}

func ActionableNewFromC(u unsafe.Pointer) *Actionable {
	c := (*C.GtkActionable)(u)
	if c == nil {
		return nil
	}

	g := &Actionable{native: c}

	return g
}

func (recv *Actionable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Activatable is a wrapper around the C record GtkActivatable.
type Activatable struct {
	native *C.GtkActivatable
}

func ActivatableNewFromC(u unsafe.Pointer) *Activatable {
	c := (*C.GtkActivatable)(u)
	if c == nil {
		return nil
	}

	g := &Activatable{native: c}

	return g
}

func (recv *Activatable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooser is a wrapper around the C record GtkAppChooser.
type AppChooser struct {
	native *C.GtkAppChooser
}

func AppChooserNewFromC(u unsafe.Pointer) *AppChooser {
	c := (*C.GtkAppChooser)(u)
	if c == nil {
		return nil
	}

	g := &AppChooser{native: c}

	return g
}

func (recv *AppChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Buildable is a wrapper around the C record GtkBuildable.
type Buildable struct {
	native *C.GtkBuildable
}

func BuildableNewFromC(u unsafe.Pointer) *Buildable {
	c := (*C.GtkBuildable)(u)
	if c == nil {
		return nil
	}

	g := &Buildable{native: c}

	return g
}

func (recv *Buildable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessibleParent is a wrapper around the C record GtkCellAccessibleParent.
type CellAccessibleParent struct {
	native *C.GtkCellAccessibleParent
}

func CellAccessibleParentNewFromC(u unsafe.Pointer) *CellAccessibleParent {
	c := (*C.GtkCellAccessibleParent)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessibleParent{native: c}

	return g
}

func (recv *CellAccessibleParent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellEditable is a wrapper around the C record GtkCellEditable.
type CellEditable struct {
	native *C.GtkCellEditable
}

func CellEditableNewFromC(u unsafe.Pointer) *CellEditable {
	c := (*C.GtkCellEditable)(u)
	if c == nil {
		return nil
	}

	g := &CellEditable{native: c}

	return g
}

func (recv *CellEditable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellLayout is a wrapper around the C record GtkCellLayout.
type CellLayout struct {
	native *C.GtkCellLayout
}

func CellLayoutNewFromC(u unsafe.Pointer) *CellLayout {
	c := (*C.GtkCellLayout)(u)
	if c == nil {
		return nil
	}

	g := &CellLayout{native: c}

	return g
}

func (recv *CellLayout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooser is a wrapper around the C record GtkColorChooser.
type ColorChooser struct {
	native *C.GtkColorChooser
}

func ColorChooserNewFromC(u unsafe.Pointer) *ColorChooser {
	c := (*C.GtkColorChooser)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooser{native: c}

	return g
}

func (recv *ColorChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Editable is a wrapper around the C record GtkEditable.
type Editable struct {
	native *C.GtkEditable
}

func EditableNewFromC(u unsafe.Pointer) *Editable {
	c := (*C.GtkEditable)(u)
	if c == nil {
		return nil
	}

	g := &Editable{native: c}

	return g
}

func (recv *Editable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooser is a wrapper around the C record GtkFileChooser.
type FileChooser struct {
	native *C.GtkFileChooser
}

func FileChooserNewFromC(u unsafe.Pointer) *FileChooser {
	c := (*C.GtkFileChooser)(u)
	if c == nil {
		return nil
	}

	g := &FileChooser{native: c}

	return g
}

func (recv *FileChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooser is a wrapper around the C record GtkFontChooser.
type FontChooser struct {
	native *C.GtkFontChooser
}

func FontChooserNewFromC(u unsafe.Pointer) *FontChooser {
	c := (*C.GtkFontChooser)(u)
	if c == nil {
		return nil
	}

	g := &FontChooser{native: c}

	return g
}

func (recv *FontChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Orientable is a wrapper around the C record GtkOrientable.
type Orientable struct {
	native *C.GtkOrientable
}

func OrientableNewFromC(u unsafe.Pointer) *Orientable {
	c := (*C.GtkOrientable)(u)
	if c == nil {
		return nil
	}

	g := &Orientable{native: c}

	return g
}

func (recv *Orientable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperationPreview is a wrapper around the C record GtkPrintOperationPreview.
type PrintOperationPreview struct {
	native *C.GtkPrintOperationPreview
}

func PrintOperationPreviewNewFromC(u unsafe.Pointer) *PrintOperationPreview {
	c := (*C.GtkPrintOperationPreview)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationPreview{native: c}

	return g
}

func (recv *PrintOperationPreview) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooser is a wrapper around the C record GtkRecentChooser.
type RecentChooser struct {
	native *C.GtkRecentChooser
}

func RecentChooserNewFromC(u unsafe.Pointer) *RecentChooser {
	c := (*C.GtkRecentChooser)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooser{native: c}

	return g
}

func (recv *RecentChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scrollable is a wrapper around the C record GtkScrollable.
type Scrollable struct {
	native *C.GtkScrollable
}

func ScrollableNewFromC(u unsafe.Pointer) *Scrollable {
	c := (*C.GtkScrollable)(u)
	if c == nil {
		return nil
	}

	g := &Scrollable{native: c}

	return g
}

func (recv *Scrollable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleProvider is a wrapper around the C record GtkStyleProvider.
type StyleProvider struct {
	native *C.GtkStyleProvider
}

func StyleProviderNewFromC(u unsafe.Pointer) *StyleProvider {
	c := (*C.GtkStyleProvider)(u)
	if c == nil {
		return nil
	}

	g := &StyleProvider{native: c}

	return g
}

func (recv *StyleProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolShell is a wrapper around the C record GtkToolShell.
type ToolShell struct {
	native *C.GtkToolShell
}

func ToolShellNewFromC(u unsafe.Pointer) *ToolShell {
	c := (*C.GtkToolShell)(u)
	if c == nil {
		return nil
	}

	g := &ToolShell{native: c}

	return g
}

func (recv *ToolShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeDragDest is a wrapper around the C record GtkTreeDragDest.
type TreeDragDest struct {
	native *C.GtkTreeDragDest
}

func TreeDragDestNewFromC(u unsafe.Pointer) *TreeDragDest {
	c := (*C.GtkTreeDragDest)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragDest{native: c}

	return g
}

func (recv *TreeDragDest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeDragSource is a wrapper around the C record GtkTreeDragSource.
type TreeDragSource struct {
	native *C.GtkTreeDragSource
}

func TreeDragSourceNewFromC(u unsafe.Pointer) *TreeDragSource {
	c := (*C.GtkTreeDragSource)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragSource{native: c}

	return g
}

func (recv *TreeDragSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModel is a wrapper around the C record GtkTreeModel.
type TreeModel struct {
	native *C.GtkTreeModel
}

func TreeModelNewFromC(u unsafe.Pointer) *TreeModel {
	c := (*C.GtkTreeModel)(u)
	if c == nil {
		return nil
	}

	g := &TreeModel{native: c}

	return g
}

func (recv *TreeModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeSortable is a wrapper around the C record GtkTreeSortable.
type TreeSortable struct {
	native *C.GtkTreeSortable
}

func TreeSortableNewFromC(u unsafe.Pointer) *TreeSortable {
	c := (*C.GtkTreeSortable)(u)
	if c == nil {
		return nil
	}

	g := &TreeSortable{native: c}

	return g
}

func (recv *TreeSortable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
