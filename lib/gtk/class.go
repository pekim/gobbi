// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gio "github.com/pekim/gobbi/lib/gio"
	"unsafe"
)

// AboutDialog is a wrapper around the C record GtkAboutDialog.
type AboutDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by AboutDialog
func (recv *AboutDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AboutDialog
func (recv *AboutDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// AccelGroup is a wrapper around the C record GtkAccelGroup.
type AccelGroup struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// AccelLabel is a wrapper around the C record GtkAccelLabel.
type AccelLabel struct {
	native unsafe.Pointer
	// label : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by AccelLabel
func (recv *AccelLabel) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AccelLabel
func (recv *AccelLabel) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// AccelMap is a wrapper around the C record GtkAccelMap.
type AccelMap struct {
	native unsafe.Pointer
}

// Accessible is a wrapper around the C record GtkAccessible.
type Accessible struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// Action is a wrapper around the C record GtkAction.
type Action struct {
	native unsafe.Pointer
	// object : record
	// Private : private_data
}

// Buildable returns the Buildable interface implemented by Action
func (recv *Action) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ActionBar is a wrapper around the C record GtkActionBar.
type ActionBar struct {
	native unsafe.Pointer
	// Private : bin
}

// ImplementorIface returns the ImplementorIface interface implemented by ActionBar
func (recv *ActionBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ActionBar
func (recv *ActionBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ActionGroup is a wrapper around the C record GtkActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by ActionGroup
func (recv *ActionGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Adjustment is a wrapper around the C record GtkAdjustment.
type Adjustment struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Alignment is a wrapper around the C record GtkAlignment.
type Alignment struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Alignment
func (recv *Alignment) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Alignment
func (recv *Alignment) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// AppChooserButton is a wrapper around the C record GtkAppChooserButton.
type AppChooserButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserButton
func (recv *AppChooserButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserButton
func (recv *AppChooserButton) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserButton
func (recv *AppChooserButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by AppChooserButton
func (recv *AppChooserButton) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by AppChooserButton
func (recv *AppChooserButton) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// AppChooserDialog is a wrapper around the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserDialog
func (recv *AppChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserDialog
func (recv *AppChooserDialog) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserDialog
func (recv *AppChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// AppChooserWidget is a wrapper around the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserWidget
func (recv *AppChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserWidget
func (recv *AppChooserWidget) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserWidget
func (recv *AppChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by AppChooserWidget
func (recv *AppChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Application is a wrapper around the C record GtkApplication.
type Application struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ActionGroup returns the ActionGroup interface implemented by Application
func (recv *Application) ActionGroup() *gio.ActionGroup {
	return gio.ActionGroupNewFromC(recv.ToC())
}

// ActionMap returns the ActionMap interface implemented by Application
func (recv *Application) ActionMap() *gio.ActionMap {
	return gio.ActionMapNewFromC(recv.ToC())
}

// ApplicationWindow is a wrapper around the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ActionGroup() *gio.ActionGroup {
	return gio.ActionGroupNewFromC(recv.ToC())
}

// ActionMap returns the ActionMap interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ActionMap() *gio.ActionMap {
	return gio.ActionMapNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ApplicationWindow
func (recv *ApplicationWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Arrow is a wrapper around the C record GtkArrow.
type Arrow struct {
	native unsafe.Pointer
	// misc : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Arrow
func (recv *Arrow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Arrow
func (recv *Arrow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ArrowAccessible is a wrapper around the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ArrowAccessible
func (recv *ArrowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ArrowAccessible
func (recv *ArrowAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// AspectFrame is a wrapper around the C record GtkAspectFrame.
type AspectFrame struct {
	native unsafe.Pointer
	// frame : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by AspectFrame
func (recv *AspectFrame) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AspectFrame
func (recv *AspectFrame) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Assistant is a wrapper around the C record GtkAssistant.
type Assistant struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Assistant
func (recv *Assistant) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Assistant
func (recv *Assistant) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Bin is a wrapper around the C record GtkBin.
type Bin struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Bin
func (recv *Bin) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Bin
func (recv *Bin) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// BooleanCellAccessible is a wrapper around the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by BooleanCellAccessible
func (recv *BooleanCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by BooleanCellAccessible
func (recv *BooleanCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Box is a wrapper around the C record GtkBox.
type Box struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Box
func (recv *Box) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Box
func (recv *Box) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Box
func (recv *Box) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Builder is a wrapper around the C record GtkBuilder.
type Builder struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Button is a wrapper around the C record GtkButton.
type Button struct {
	native unsafe.Pointer
	// Private : bin
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Button
func (recv *Button) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by Button
func (recv *Button) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by Button
func (recv *Button) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Button
func (recv *Button) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ButtonAccessible is a wrapper around the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ButtonBox is a wrapper around the C record GtkButtonBox.
type ButtonBox struct {
	native unsafe.Pointer
	// box : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ButtonBox
func (recv *ButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ButtonBox
func (recv *ButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ButtonBox
func (recv *ButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Calendar is a wrapper around the C record GtkCalendar.
type Calendar struct {
	native unsafe.Pointer
	// widget : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Calendar
func (recv *Calendar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Calendar
func (recv *Calendar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellAccessible is a wrapper around the C record GtkCellAccessible.
type CellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by CellAccessible
func (recv *CellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by CellAccessible
func (recv *CellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// CellArea is a wrapper around the C record GtkCellArea.
type CellArea struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// Buildable returns the Buildable interface implemented by CellArea
func (recv *CellArea) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellArea
func (recv *CellArea) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// CellAreaBox is a wrapper around the C record GtkCellAreaBox.
type CellAreaBox struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// Buildable returns the Buildable interface implemented by CellAreaBox
func (recv *CellAreaBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellAreaBox
func (recv *CellAreaBox) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by CellAreaBox
func (recv *CellAreaBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// CellAreaContext is a wrapper around the C record GtkCellAreaContext.
type CellAreaContext struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// CellRenderer is a wrapper around the C record GtkCellRenderer.
type CellRenderer struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// CellRendererAccel is a wrapper around the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererCombo is a wrapper around the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererPixbuf is a wrapper around the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererProgress is a wrapper around the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Orientable returns the Orientable interface implemented by CellRendererProgress
func (recv *CellRendererProgress) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// CellRendererSpin is a wrapper around the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererSpinner is a wrapper around the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererText is a wrapper around the C record GtkCellRendererText.
type CellRendererText struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellRendererToggle is a wrapper around the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// CellView is a wrapper around the C record GtkCellView.
type CellView struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by CellView
func (recv *CellView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CellView
func (recv *CellView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellView
func (recv *CellView) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by CellView
func (recv *CellView) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// CheckButton is a wrapper around the C record GtkCheckButton.
type CheckButton struct {
	native unsafe.Pointer
	// toggle_button : record
}

// ImplementorIface returns the ImplementorIface interface implemented by CheckButton
func (recv *CheckButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by CheckButton
func (recv *CheckButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by CheckButton
func (recv *CheckButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CheckButton
func (recv *CheckButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CheckMenuItem is a wrapper around the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native unsafe.Pointer
	// menu_item : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by CheckMenuItem
func (recv *CheckMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CheckMenuItemAccessible is a wrapper around the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Clipboard is a wrapper around the C record GtkClipboard.
type Clipboard struct {
	native unsafe.Pointer
}

// ColorButton is a wrapper around the C record GtkColorButton.
type ColorButton struct {
	native unsafe.Pointer
	// button : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorButton
func (recv *ColorButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ColorButton
func (recv *ColorButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ColorButton
func (recv *ColorButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorButton
func (recv *ColorButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorButton
func (recv *ColorButton) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// ColorChooserDialog is a wrapper around the C record GtkColorChooserDialog.
type ColorChooserDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// ColorChooserWidget is a wrapper around the C record GtkColorChooserWidget.
type ColorChooserWidget struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ColorSelection is a wrapper around the C record GtkColorSelection.
type ColorSelection struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : private_data
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorSelection
func (recv *ColorSelection) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorSelection
func (recv *ColorSelection) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ColorSelection
func (recv *ColorSelection) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ColorSelectionDialog is a wrapper around the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorSelectionDialog
func (recv *ColorSelectionDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorSelectionDialog
func (recv *ColorSelectionDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ComboBox is a wrapper around the C record GtkComboBox.
type ComboBox struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ComboBox
func (recv *ComboBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ComboBox
func (recv *ComboBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by ComboBox
func (recv *ComboBox) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by ComboBox
func (recv *ComboBox) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// ComboBoxAccessible is a wrapper around the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ComboBoxText is a wrapper around the C record GtkComboBoxText.
type ComboBoxText struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ComboBoxText
func (recv *ComboBoxText) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ComboBoxText
func (recv *ComboBoxText) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by ComboBoxText
func (recv *ComboBoxText) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by ComboBoxText
func (recv *ComboBoxText) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Container is a wrapper around the C record GtkContainer.
type Container struct {
	native unsafe.Pointer
	// widget : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Container
func (recv *Container) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Container
func (recv *Container) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ContainerAccessible is a wrapper around the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ContainerAccessible
func (recv *ContainerAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ContainerCellAccessible is a wrapper around the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ContainerCellAccessible
func (recv *ContainerCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ContainerCellAccessible
func (recv *ContainerCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// CssProvider is a wrapper around the C record GtkCssProvider.
type CssProvider struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// StyleProvider returns the StyleProvider interface implemented by CssProvider
func (recv *CssProvider) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// Dialog is a wrapper around the C record GtkDialog.
type Dialog struct {
	native unsafe.Pointer
	// window : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Dialog
func (recv *Dialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Dialog
func (recv *Dialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// DrawingArea is a wrapper around the C record GtkDrawingArea.
type DrawingArea struct {
	native unsafe.Pointer
	// widget : record
	// Private : dummy
}

// ImplementorIface returns the ImplementorIface interface implemented by DrawingArea
func (recv *DrawingArea) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by DrawingArea
func (recv *DrawingArea) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Entry is a wrapper around the C record GtkEntry.
type Entry struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Entry
func (recv *Entry) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Entry
func (recv *Entry) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by Entry
func (recv *Entry) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by Entry
func (recv *Entry) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// EntryAccessible is a wrapper around the C record GtkEntryAccessible.
type EntryAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by EntryAccessible
func (recv *EntryAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by EntryAccessible
func (recv *EntryAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by EntryAccessible
func (recv *EntryAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by EntryAccessible
func (recv *EntryAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// EntryBuffer is a wrapper around the C record GtkEntryBuffer.
type EntryBuffer struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// EntryCompletion is a wrapper around the C record GtkEntryCompletion.
type EntryCompletion struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by EntryCompletion
func (recv *EntryCompletion) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by EntryCompletion
func (recv *EntryCompletion) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Unsupported : EntryIconAccessible : no CType

// EventBox is a wrapper around the C record GtkEventBox.
type EventBox struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by EventBox
func (recv *EventBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by EventBox
func (recv *EventBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// EventController is a wrapper around the C record GtkEventController.
type EventController struct {
	native unsafe.Pointer
}

// Expander is a wrapper around the C record GtkExpander.
type Expander struct {
	native unsafe.Pointer
	// bin : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Expander
func (recv *Expander) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Expander
func (recv *Expander) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ExpanderAccessible is a wrapper around the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ExpanderAccessible
func (recv *ExpanderAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ExpanderAccessible
func (recv *ExpanderAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// FileChooserButton is a wrapper around the C record GtkFileChooserButton.
type FileChooserButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserButton
func (recv *FileChooserButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserButton
func (recv *FileChooserButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserButton
func (recv *FileChooserButton) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FileChooserButton
func (recv *FileChooserButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// FileChooserDialog is a wrapper around the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserDialog
func (recv *FileChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserDialog
func (recv *FileChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserDialog
func (recv *FileChooserDialog) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// FileChooserWidget is a wrapper around the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserWidget
func (recv *FileChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserWidget
func (recv *FileChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserWidget
func (recv *FileChooserWidget) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FileChooserWidget
func (recv *FileChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// FileFilter is a wrapper around the C record GtkFileFilter.
type FileFilter struct {
	native unsafe.Pointer
}

// Buildable returns the Buildable interface implemented by FileFilter
func (recv *FileFilter) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Fixed is a wrapper around the C record GtkFixed.
type Fixed struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Fixed
func (recv *Fixed) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Fixed
func (recv *Fixed) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FlowBox is a wrapper around the C record GtkFlowBox.
type FlowBox struct {
	native unsafe.Pointer
	// container : record
}

// ImplementorIface returns the ImplementorIface interface implemented by FlowBox
func (recv *FlowBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FlowBox
func (recv *FlowBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FlowBox
func (recv *FlowBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// FlowBoxAccessible is a wrapper around the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by FlowBoxAccessible
func (recv *FlowBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by FlowBoxAccessible
func (recv *FlowBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// FlowBoxChild is a wrapper around the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ImplementorIface returns the ImplementorIface interface implemented by FlowBoxChild
func (recv *FlowBoxChild) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FlowBoxChild
func (recv *FlowBoxChild) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FlowBoxChildAccessible is a wrapper around the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native unsafe.Pointer
	// parent : record
}

// Component returns the Component interface implemented by FlowBoxChildAccessible
func (recv *FlowBoxChildAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// FontButton is a wrapper around the C record GtkFontButton.
type FontButton struct {
	native unsafe.Pointer
	// button : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FontButton
func (recv *FontButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by FontButton
func (recv *FontButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by FontButton
func (recv *FontButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontButton
func (recv *FontButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontButton
func (recv *FontButton) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// FontChooserDialog is a wrapper around the C record GtkFontChooserDialog.
type FontChooserDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FontChooserDialog
func (recv *FontChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontChooserDialog
func (recv *FontChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontChooserDialog
func (recv *FontChooserDialog) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// FontChooserWidget is a wrapper around the C record GtkFontChooserWidget.
type FontChooserWidget struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FontChooserWidget
func (recv *FontChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontChooserWidget
func (recv *FontChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontChooserWidget
func (recv *FontChooserWidget) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FontChooserWidget
func (recv *FontChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// FontSelection is a wrapper around the C record GtkFontSelection.
type FontSelection struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FontSelection
func (recv *FontSelection) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontSelection
func (recv *FontSelection) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FontSelection
func (recv *FontSelection) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// FontSelectionDialog is a wrapper around the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by FontSelectionDialog
func (recv *FontSelectionDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontSelectionDialog
func (recv *FontSelectionDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Frame is a wrapper around the C record GtkFrame.
type Frame struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Frame
func (recv *Frame) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Frame
func (recv *Frame) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FrameAccessible is a wrapper around the C record GtkFrameAccessible.
type FrameAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by FrameAccessible
func (recv *FrameAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Gesture is a wrapper around the C record GtkGesture.
type Gesture struct {
	native unsafe.Pointer
}

// GestureDrag is a wrapper around the C record GtkGestureDrag.
type GestureDrag struct {
	native unsafe.Pointer
}

// GestureLongPress is a wrapper around the C record GtkGestureLongPress.
type GestureLongPress struct {
	native unsafe.Pointer
}

// GestureMultiPress is a wrapper around the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native unsafe.Pointer
}

// GesturePan is a wrapper around the C record GtkGesturePan.
type GesturePan struct {
	native unsafe.Pointer
}

// GestureRotate is a wrapper around the C record GtkGestureRotate.
type GestureRotate struct {
	native unsafe.Pointer
}

// GestureSingle is a wrapper around the C record GtkGestureSingle.
type GestureSingle struct {
	native unsafe.Pointer
}

// GestureSwipe is a wrapper around the C record GtkGestureSwipe.
type GestureSwipe struct {
	native unsafe.Pointer
}

// GestureZoom is a wrapper around the C record GtkGestureZoom.
type GestureZoom struct {
	native unsafe.Pointer
}

// Grid is a wrapper around the C record GtkGrid.
type Grid struct {
	native unsafe.Pointer
	// Private : container
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Grid
func (recv *Grid) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Grid
func (recv *Grid) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Grid
func (recv *Grid) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HBox is a wrapper around the C record GtkHBox.
type HBox struct {
	native unsafe.Pointer
	// box : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HBox
func (recv *HBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HBox
func (recv *HBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HBox
func (recv *HBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HButtonBox is a wrapper around the C record GtkHButtonBox.
type HButtonBox struct {
	native unsafe.Pointer
	// button_box : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HButtonBox
func (recv *HButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HButtonBox
func (recv *HButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HButtonBox
func (recv *HButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HPaned is a wrapper around the C record GtkHPaned.
type HPaned struct {
	native unsafe.Pointer
	// paned : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HPaned
func (recv *HPaned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HPaned
func (recv *HPaned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HPaned
func (recv *HPaned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HSV is a wrapper around the C record GtkHSV.
type HSV struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by HSV
func (recv *HSV) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HSV
func (recv *HSV) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// HScale is a wrapper around the C record GtkHScale.
type HScale struct {
	native unsafe.Pointer
	// scale : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HScale
func (recv *HScale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HScale
func (recv *HScale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HScale
func (recv *HScale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HScrollbar is a wrapper around the C record GtkHScrollbar.
type HScrollbar struct {
	native unsafe.Pointer
	// scrollbar : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HScrollbar
func (recv *HScrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HScrollbar
func (recv *HScrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HScrollbar
func (recv *HScrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HSeparator is a wrapper around the C record GtkHSeparator.
type HSeparator struct {
	native unsafe.Pointer
	// separator : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HSeparator
func (recv *HSeparator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HSeparator
func (recv *HSeparator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HSeparator
func (recv *HSeparator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// HandleBox is a wrapper around the C record GtkHandleBox.
type HandleBox struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by HandleBox
func (recv *HandleBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HandleBox
func (recv *HandleBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// HeaderBar is a wrapper around the C record GtkHeaderBar.
type HeaderBar struct {
	native unsafe.Pointer
	// container : record
}

// ImplementorIface returns the ImplementorIface interface implemented by HeaderBar
func (recv *HeaderBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HeaderBar
func (recv *HeaderBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// IMContext is a wrapper around the C record GtkIMContext.
type IMContext struct {
	native unsafe.Pointer
	// parent_instance : record
}

// IMContextSimple is a wrapper around the C record GtkIMContextSimple.
type IMContextSimple struct {
	native unsafe.Pointer
	// object : record
	// Private : priv
}

// IMMulticontext is a wrapper around the C record GtkIMMulticontext.
type IMMulticontext struct {
	native unsafe.Pointer
	// object : record
	// Private : priv
}

// IconFactory is a wrapper around the C record GtkIconFactory.
type IconFactory struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by IconFactory
func (recv *IconFactory) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// IconInfo is a wrapper around the C record GtkIconInfo.
type IconInfo struct {
	native unsafe.Pointer
}

// IconTheme is a wrapper around the C record GtkIconTheme.
type IconTheme struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// IconView is a wrapper around the C record GtkIconView.
type IconView struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by IconView
func (recv *IconView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by IconView
func (recv *IconView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by IconView
func (recv *IconView) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by IconView
func (recv *IconView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// IconViewAccessible is a wrapper around the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by IconViewAccessible
func (recv *IconViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by IconViewAccessible
func (recv *IconViewAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Image is a wrapper around the C record GtkImage.
type Image struct {
	native unsafe.Pointer
	// misc : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Image
func (recv *Image) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Image
func (recv *Image) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImageAccessible is a wrapper around the C record GtkImageAccessible.
type ImageAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ImageAccessible
func (recv *ImageAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ImageAccessible
func (recv *ImageAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImageCellAccessible is a wrapper around the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImageMenuItem is a wrapper around the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native unsafe.Pointer
	// menu_item : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ImageMenuItem
func (recv *ImageMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// InfoBar is a wrapper around the C record GtkInfoBar.
type InfoBar struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by InfoBar
func (recv *InfoBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by InfoBar
func (recv *InfoBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by InfoBar
func (recv *InfoBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Invisible is a wrapper around the C record GtkInvisible.
type Invisible struct {
	native unsafe.Pointer
	// widget : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Invisible
func (recv *Invisible) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Invisible
func (recv *Invisible) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Label is a wrapper around the C record GtkLabel.
type Label struct {
	native unsafe.Pointer
	// misc : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Label
func (recv *Label) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Label
func (recv *Label) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// LabelAccessible is a wrapper around the C record GtkLabelAccessible.
type LabelAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by LabelAccessible
func (recv *LabelAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Hypertext returns the Hypertext interface implemented by LabelAccessible
func (recv *LabelAccessible) Hypertext() *atk.Hypertext {
	return atk.HypertextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by LabelAccessible
func (recv *LabelAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Layout is a wrapper around the C record GtkLayout.
type Layout struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Layout
func (recv *Layout) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Layout
func (recv *Layout) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by Layout
func (recv *Layout) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// LevelBar is a wrapper around the C record GtkLevelBar.
type LevelBar struct {
	native unsafe.Pointer
	// Private : parent
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by LevelBar
func (recv *LevelBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LevelBar
func (recv *LevelBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by LevelBar
func (recv *LevelBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// LevelBarAccessible is a wrapper around the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by LevelBarAccessible
func (recv *LevelBarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by LevelBarAccessible
func (recv *LevelBarAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// LinkButton is a wrapper around the C record GtkLinkButton.
type LinkButton struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by LinkButton
func (recv *LinkButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by LinkButton
func (recv *LinkButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by LinkButton
func (recv *LinkButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LinkButton
func (recv *LinkButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// LinkButtonAccessible is a wrapper around the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// HyperlinkImpl returns the HyperlinkImpl interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) HyperlinkImpl() *atk.HyperlinkImpl {
	return atk.HyperlinkImplNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ListBox is a wrapper around the C record GtkListBox.
type ListBox struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ImplementorIface returns the ImplementorIface interface implemented by ListBox
func (recv *ListBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ListBox
func (recv *ListBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ListBoxAccessible is a wrapper around the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ListBoxAccessible
func (recv *ListBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by ListBoxAccessible
func (recv *ListBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ListBoxRow is a wrapper around the C record GtkListBoxRow.
type ListBoxRow struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ImplementorIface returns the ImplementorIface interface implemented by ListBoxRow
func (recv *ListBoxRow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ListBoxRow
func (recv *ListBoxRow) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ListBoxRow
func (recv *ListBoxRow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ListBoxRowAccessible is a wrapper around the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native unsafe.Pointer
	// parent : record
}

// Component returns the Component interface implemented by ListBoxRowAccessible
func (recv *ListBoxRowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ListStore is a wrapper around the C record GtkListStore.
type ListStore struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by ListStore
func (recv *ListStore) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TreeDragDest returns the TreeDragDest interface implemented by ListStore
func (recv *ListStore) TreeDragDest() *TreeDragDest {
	return TreeDragDestNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by ListStore
func (recv *ListStore) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by ListStore
func (recv *ListStore) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by ListStore
func (recv *ListStore) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// LockButton is a wrapper around the C record GtkLockButton.
type LockButton struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by LockButton
func (recv *LockButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by LockButton
func (recv *LockButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by LockButton
func (recv *LockButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LockButton
func (recv *LockButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// LockButtonAccessible is a wrapper around the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Menu is a wrapper around the C record GtkMenu.
type Menu struct {
	native unsafe.Pointer
	// menu_shell : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Menu
func (recv *Menu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Menu
func (recv *Menu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MenuAccessible is a wrapper around the C record GtkMenuAccessible.
type MenuAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by MenuAccessible
func (recv *MenuAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuAccessible
func (recv *MenuAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// MenuBar is a wrapper around the C record GtkMenuBar.
type MenuBar struct {
	native unsafe.Pointer
	// menu_shell : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuBar
func (recv *MenuBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuBar
func (recv *MenuBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MenuButton is a wrapper around the C record GtkMenuButton.
type MenuButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuButton
func (recv *MenuButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuButton
func (recv *MenuButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuButton
func (recv *MenuButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuButton
func (recv *MenuButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MenuButtonAccessible is a wrapper around the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// MenuItem is a wrapper around the C record GtkMenuItem.
type MenuItem struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuItem
func (recv *MenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuItem
func (recv *MenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuItem
func (recv *MenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuItem
func (recv *MenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MenuItemAccessible is a wrapper around the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// MenuShell is a wrapper around the C record GtkMenuShell.
type MenuShell struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuShell
func (recv *MenuShell) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuShell
func (recv *MenuShell) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MenuShellAccessible is a wrapper around the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by MenuShellAccessible
func (recv *MenuShellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuShellAccessible
func (recv *MenuShellAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// MenuToolButton is a wrapper around the C record GtkMenuToolButton.
type MenuToolButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuToolButton
func (recv *MenuToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuToolButton
func (recv *MenuToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuToolButton
func (recv *MenuToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuToolButton
func (recv *MenuToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MessageDialog is a wrapper around the C record GtkMessageDialog.
type MessageDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by MessageDialog
func (recv *MessageDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MessageDialog
func (recv *MessageDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Misc is a wrapper around the C record GtkMisc.
type Misc struct {
	native unsafe.Pointer
	// widget : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Misc
func (recv *Misc) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Misc
func (recv *Misc) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ModelButton is a wrapper around the C record GtkModelButton.
type ModelButton struct {
	native unsafe.Pointer
}

// ImplementorIface returns the ImplementorIface interface implemented by ModelButton
func (recv *ModelButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ModelButton
func (recv *ModelButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ModelButton
func (recv *ModelButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ModelButton
func (recv *ModelButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// MountOperation is a wrapper around the C record GtkMountOperation.
type MountOperation struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Notebook is a wrapper around the C record GtkNotebook.
type Notebook struct {
	native unsafe.Pointer
	// Private : container
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Notebook
func (recv *Notebook) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Notebook
func (recv *Notebook) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// NotebookAccessible is a wrapper around the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by NotebookAccessible
func (recv *NotebookAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by NotebookAccessible
func (recv *NotebookAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// NotebookPageAccessible is a wrapper around the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by NotebookPageAccessible
func (recv *NotebookPageAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// NumerableIcon is a wrapper around the C record GtkNumerableIcon.
type NumerableIcon struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// Icon returns the Icon interface implemented by NumerableIcon
func (recv *NumerableIcon) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// OffscreenWindow is a wrapper around the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native unsafe.Pointer
	// parent_object : record
}

// ImplementorIface returns the ImplementorIface interface implemented by OffscreenWindow
func (recv *OffscreenWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by OffscreenWindow
func (recv *OffscreenWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Overlay is a wrapper around the C record GtkOverlay.
type Overlay struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Overlay
func (recv *Overlay) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Overlay
func (recv *Overlay) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// PageSetup is a wrapper around the C record GtkPageSetup.
type PageSetup struct {
	native unsafe.Pointer
}

// Paned is a wrapper around the C record GtkPaned.
type Paned struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Paned
func (recv *Paned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Paned
func (recv *Paned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Paned
func (recv *Paned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// PanedAccessible is a wrapper around the C record GtkPanedAccessible.
type PanedAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by PanedAccessible
func (recv *PanedAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by PanedAccessible
func (recv *PanedAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// PlacesSidebar is a wrapper around the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native unsafe.Pointer
}

// ImplementorIface returns the ImplementorIface interface implemented by PlacesSidebar
func (recv *PlacesSidebar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by PlacesSidebar
func (recv *PlacesSidebar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkPlug

// Popover is a wrapper around the C record GtkPopover.
type Popover struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Popover
func (recv *Popover) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Popover
func (recv *Popover) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// PopoverAccessible is a wrapper around the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native unsafe.Pointer
	// parent : record
}

// Component returns the Component interface implemented by PopoverAccessible
func (recv *PopoverAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// PopoverMenu is a wrapper around the C record GtkPopoverMenu.
type PopoverMenu struct {
	native unsafe.Pointer
}

// ImplementorIface returns the ImplementorIface interface implemented by PopoverMenu
func (recv *PopoverMenu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by PopoverMenu
func (recv *PopoverMenu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// PrintContext is a wrapper around the C record GtkPrintContext.
type PrintContext struct {
	native unsafe.Pointer
}

// PrintOperation is a wrapper around the C record GtkPrintOperation.
type PrintOperation struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PrintOperationPreview returns the PrintOperationPreview interface implemented by PrintOperation
func (recv *PrintOperation) PrintOperationPreview() *PrintOperationPreview {
	return PrintOperationPreviewNewFromC(recv.ToC())
}

// PrintSettings is a wrapper around the C record GtkPrintSettings.
type PrintSettings struct {
	native unsafe.Pointer
}

// ProgressBar is a wrapper around the C record GtkProgressBar.
type ProgressBar struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ProgressBar
func (recv *ProgressBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ProgressBar
func (recv *ProgressBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ProgressBar
func (recv *ProgressBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ProgressBarAccessible is a wrapper around the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ProgressBarAccessible
func (recv *ProgressBarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ProgressBarAccessible
func (recv *ProgressBarAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// RadioAction is a wrapper around the C record GtkRadioAction.
type RadioAction struct {
	native unsafe.Pointer
	// parent : record
	// Private : private_data
}

// Buildable returns the Buildable interface implemented by RadioAction
func (recv *RadioAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RadioButton is a wrapper around the C record GtkRadioButton.
type RadioButton struct {
	native unsafe.Pointer
	// check_button : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by RadioButton
func (recv *RadioButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioButton
func (recv *RadioButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioButton
func (recv *RadioButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioButton
func (recv *RadioButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RadioButtonAccessible is a wrapper around the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// RadioMenuItem is a wrapper around the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native unsafe.Pointer
	// check_menu_item : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by RadioMenuItem
func (recv *RadioMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RadioMenuItemAccessible is a wrapper around the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// RadioToolButton is a wrapper around the C record GtkRadioToolButton.
type RadioToolButton struct {
	native unsafe.Pointer
	// parent : record
}

// ImplementorIface returns the ImplementorIface interface implemented by RadioToolButton
func (recv *RadioToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioToolButton
func (recv *RadioToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioToolButton
func (recv *RadioToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioToolButton
func (recv *RadioToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Range is a wrapper around the C record GtkRange.
type Range struct {
	native unsafe.Pointer
	// widget : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Range
func (recv *Range) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Range
func (recv *Range) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Range
func (recv *Range) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// RangeAccessible is a wrapper around the C record GtkRangeAccessible.
type RangeAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by RangeAccessible
func (recv *RangeAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by RangeAccessible
func (recv *RangeAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// RcStyle is a wrapper around the C record GtkRcStyle.
type RcStyle struct {
	native unsafe.Pointer
	// parent_instance : record
	Name string
	// no type for bg_pixmap_name
	// font_desc : record
	// no type for color_flags
	// no type for fg
	// no type for bg
	// no type for text
	// no type for base
	Xthickness int32
	Ythickness int32
	// Private : rc_properties
	// Private : rc_style_lists
	// Private : icon_factories
	// Private : engine_specified
}

// RecentAction is a wrapper around the C record GtkRecentAction.
type RecentAction struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by RecentAction
func (recv *RecentAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentAction
func (recv *RecentAction) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// RecentChooserDialog is a wrapper around the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// RecentChooserMenu is a wrapper around the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// RecentChooserWidget is a wrapper around the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// RecentFilter is a wrapper around the C record GtkRecentFilter.
type RecentFilter struct {
	native unsafe.Pointer
}

// Buildable returns the Buildable interface implemented by RecentFilter
func (recv *RecentFilter) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RendererCellAccessible is a wrapper around the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by RendererCellAccessible
func (recv *RendererCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RendererCellAccessible
func (recv *RendererCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Revealer is a wrapper around the C record GtkRevealer.
type Revealer struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Revealer
func (recv *Revealer) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Revealer
func (recv *Revealer) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scale is a wrapper around the C record GtkScale.
type Scale struct {
	native unsafe.Pointer
	// range : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Scale
func (recv *Scale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Scale
func (recv *Scale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Scale
func (recv *Scale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ScaleAccessible is a wrapper around the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ScaleAccessible
func (recv *ScaleAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ScaleAccessible
func (recv *ScaleAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ScaleButton is a wrapper around the C record GtkScaleButton.
type ScaleButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ScaleButton
func (recv *ScaleButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ScaleButton
func (recv *ScaleButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ScaleButton
func (recv *ScaleButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ScaleButton
func (recv *ScaleButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ScaleButton
func (recv *ScaleButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ScaleButtonAccessible is a wrapper around the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// Scrollbar is a wrapper around the C record GtkScrollbar.
type Scrollbar struct {
	native unsafe.Pointer
	// range : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Scrollbar
func (recv *Scrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Scrollbar
func (recv *Scrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Scrollbar
func (recv *Scrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ScrolledWindow is a wrapper around the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native unsafe.Pointer
	// container : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by ScrolledWindow
func (recv *ScrolledWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ScrolledWindow
func (recv *ScrolledWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ScrolledWindowAccessible is a wrapper around the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by ScrolledWindowAccessible
func (recv *ScrolledWindowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// SearchBar is a wrapper around the C record GtkSearchBar.
type SearchBar struct {
	native unsafe.Pointer
	// Private : parent
}

// ImplementorIface returns the ImplementorIface interface implemented by SearchBar
func (recv *SearchBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SearchBar
func (recv *SearchBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// SearchEntry is a wrapper around the C record GtkSearchEntry.
type SearchEntry struct {
	native unsafe.Pointer
	// parent : record
}

// ImplementorIface returns the ImplementorIface interface implemented by SearchEntry
func (recv *SearchEntry) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SearchEntry
func (recv *SearchEntry) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by SearchEntry
func (recv *SearchEntry) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by SearchEntry
func (recv *SearchEntry) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// Separator is a wrapper around the C record GtkSeparator.
type Separator struct {
	native unsafe.Pointer
	// widget : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Separator
func (recv *Separator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Separator
func (recv *Separator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Separator
func (recv *Separator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// SeparatorMenuItem is a wrapper around the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native unsafe.Pointer
	// menu_item : record
}

// ImplementorIface returns the ImplementorIface interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// SeparatorToolItem is a wrapper around the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Settings is a wrapper around the C record GtkSettings.
type Settings struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// StyleProvider returns the StyleProvider interface implemented by Settings
func (recv *Settings) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// SizeGroup is a wrapper around the C record GtkSizeGroup.
type SizeGroup struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Buildable returns the Buildable interface implemented by SizeGroup
func (recv *SizeGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkSocket

// SpinButton is a wrapper around the C record GtkSpinButton.
type SpinButton struct {
	native unsafe.Pointer
	// entry : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by SpinButton
func (recv *SpinButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SpinButton
func (recv *SpinButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by SpinButton
func (recv *SpinButton) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by SpinButton
func (recv *SpinButton) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by SpinButton
func (recv *SpinButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// SpinButtonAccessible is a wrapper around the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// Spinner is a wrapper around the C record GtkSpinner.
type Spinner struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Spinner
func (recv *Spinner) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Spinner
func (recv *Spinner) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// SpinnerAccessible is a wrapper around the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by SpinnerAccessible
func (recv *SpinnerAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by SpinnerAccessible
func (recv *SpinnerAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Stack is a wrapper around the C record GtkStack.
type Stack struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Stack
func (recv *Stack) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Stack
func (recv *Stack) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkStackAccessible

// StackSidebar is a wrapper around the C record GtkStackSidebar.
type StackSidebar struct {
	native unsafe.Pointer
	// parent : record
}

// ImplementorIface returns the ImplementorIface interface implemented by StackSidebar
func (recv *StackSidebar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StackSidebar
func (recv *StackSidebar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// StackSwitcher is a wrapper around the C record GtkStackSwitcher.
type StackSwitcher struct {
	native unsafe.Pointer
	// widget : record
}

// ImplementorIface returns the ImplementorIface interface implemented by StackSwitcher
func (recv *StackSwitcher) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StackSwitcher
func (recv *StackSwitcher) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by StackSwitcher
func (recv *StackSwitcher) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// StatusIcon is a wrapper around the C record GtkStatusIcon.
type StatusIcon struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Statusbar is a wrapper around the C record GtkStatusbar.
type Statusbar struct {
	native unsafe.Pointer
	// parent_widget : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Statusbar
func (recv *Statusbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Statusbar
func (recv *Statusbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Statusbar
func (recv *Statusbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// StatusbarAccessible is a wrapper around the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by StatusbarAccessible
func (recv *StatusbarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Style is a wrapper around the C record GtkStyle.
type Style struct {
	native unsafe.Pointer
	// Private : parent_instance
	// no type for fg
	// no type for bg
	// no type for light
	// no type for dark
	// no type for mid
	// no type for text
	// no type for base
	// no type for text_aa
	// black : record
	// white : record
	// font_desc : record
	Xthickness int32
	Ythickness int32
	// no type for background
	// Private : attach_count
	// Private : visual
	// Private : private_font_desc
	// Private : rc_style
	// Private : styles
	// Private : property_cache
	// Private : icon_factories
}

// StyleContext is a wrapper around the C record GtkStyleContext.
type StyleContext struct {
	native unsafe.Pointer
	// parent_object : record
	// priv : record
}

// StyleProperties is a wrapper around the C record GtkStyleProperties.
type StyleProperties struct {
	native unsafe.Pointer
	// Private : parent_object
	// Private : priv
}

// StyleProvider returns the StyleProvider interface implemented by StyleProperties
func (recv *StyleProperties) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// Switch is a wrapper around the C record GtkSwitch.
type Switch struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Switch
func (recv *Switch) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by Switch
func (recv *Switch) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by Switch
func (recv *Switch) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Switch
func (recv *Switch) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// SwitchAccessible is a wrapper around the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by SwitchAccessible
func (recv *SwitchAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by SwitchAccessible
func (recv *SwitchAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Table is a wrapper around the C record GtkTable.
type Table struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Table
func (recv *Table) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Table
func (recv *Table) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TearoffMenuItem is a wrapper around the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native unsafe.Pointer
	// menu_item : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TextBuffer is a wrapper around the C record GtkTextBuffer.
type TextBuffer struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// TextCellAccessible is a wrapper around the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// TextChildAnchor is a wrapper around the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : segment
}

// TextMark is a wrapper around the C record GtkTextMark.
type TextMark struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : segment
}

// TextTag is a wrapper around the C record GtkTextTag.
type TextTag struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// TextTagTable is a wrapper around the C record GtkTextTagTable.
type TextTagTable struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Buildable returns the Buildable interface implemented by TextTagTable
func (recv *TextTagTable) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TextView is a wrapper around the C record GtkTextView.
type TextView struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by TextView
func (recv *TextView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TextView
func (recv *TextView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by TextView
func (recv *TextView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// TextViewAccessible is a wrapper around the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by TextViewAccessible
func (recv *TextViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by TextViewAccessible
func (recv *TextViewAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// StreamableContent returns the StreamableContent interface implemented by TextViewAccessible
func (recv *TextViewAccessible) StreamableContent() *atk.StreamableContent {
	return atk.StreamableContentNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by TextViewAccessible
func (recv *TextViewAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// ThemingEngine is a wrapper around the C record GtkThemingEngine.
type ThemingEngine struct {
	native unsafe.Pointer
	// parent_object : record
	// priv : record
}

// ToggleAction is a wrapper around the C record GtkToggleAction.
type ToggleAction struct {
	native unsafe.Pointer
	// parent : record
	// Private : private_data
}

// Buildable returns the Buildable interface implemented by ToggleAction
func (recv *ToggleAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToggleButton is a wrapper around the C record GtkToggleButton.
type ToggleButton struct {
	native unsafe.Pointer
	// Private : button
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ToggleButton
func (recv *ToggleButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToggleButton
func (recv *ToggleButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToggleButton
func (recv *ToggleButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToggleButton
func (recv *ToggleButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToggleButtonAccessible is a wrapper around the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Action returns the Action interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ToggleToolButton is a wrapper around the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ToggleToolButton
func (recv *ToggleToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToolButton is a wrapper around the C record GtkToolButton.
type ToolButton struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolButton
func (recv *ToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToolButton
func (recv *ToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToolButton
func (recv *ToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolButton
func (recv *ToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToolItem is a wrapper around the C record GtkToolItem.
type ToolItem struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolItem
func (recv *ToolItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToolItem
func (recv *ToolItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolItem
func (recv *ToolItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToolItemGroup is a wrapper around the C record GtkToolItemGroup.
type ToolItemGroup struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolItemGroup
func (recv *ToolItemGroup) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolItemGroup
func (recv *ToolItemGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToolShell returns the ToolShell interface implemented by ToolItemGroup
func (recv *ToolItemGroup) ToolShell() *ToolShell {
	return ToolShellNewFromC(recv.ToC())
}

// ToolPalette is a wrapper around the C record GtkToolPalette.
type ToolPalette struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolPalette
func (recv *ToolPalette) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolPalette
func (recv *ToolPalette) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ToolPalette
func (recv *ToolPalette) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by ToolPalette
func (recv *ToolPalette) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// Toolbar is a wrapper around the C record GtkToolbar.
type Toolbar struct {
	native unsafe.Pointer
	// container : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Toolbar
func (recv *Toolbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Toolbar
func (recv *Toolbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Toolbar
func (recv *Toolbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ToolShell returns the ToolShell interface implemented by Toolbar
func (recv *Toolbar) ToolShell() *ToolShell {
	return ToolShellNewFromC(recv.ToC())
}

// Tooltip is a wrapper around the C record GtkTooltip.
type Tooltip struct {
	native unsafe.Pointer
}

// ToplevelAccessible is a wrapper around the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// TreeModelFilter is a wrapper around the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeModelFilter
func (recv *TreeModelFilter) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeModelFilter
func (recv *TreeModelFilter) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeModelSort is a wrapper around the C record GtkTreeModelSort.
type TreeModelSort struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// TreeSelection is a wrapper around the C record GtkTreeSelection.
type TreeSelection struct {
	native unsafe.Pointer
	// Private : parent
	// Private : priv
}

// TreeStore is a wrapper around the C record GtkTreeStore.
type TreeStore struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Buildable returns the Buildable interface implemented by TreeStore
func (recv *TreeStore) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TreeDragDest returns the TreeDragDest interface implemented by TreeStore
func (recv *TreeStore) TreeDragDest() *TreeDragDest {
	return TreeDragDestNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeStore
func (recv *TreeStore) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeStore
func (recv *TreeStore) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by TreeStore
func (recv *TreeStore) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// TreeView is a wrapper around the C record GtkTreeView.
type TreeView struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by TreeView
func (recv *TreeView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TreeView
func (recv *TreeView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by TreeView
func (recv *TreeView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// TreeViewAccessible is a wrapper around the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Table returns the Table interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Table() *atk.Table {
	return atk.TableNewFromC(recv.ToC())
}

// CellAccessibleParent returns the CellAccessibleParent interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) CellAccessibleParent() *CellAccessibleParent {
	return CellAccessibleParentNewFromC(recv.ToC())
}

// TreeViewColumn is a wrapper around the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Buildable returns the Buildable interface implemented by TreeViewColumn
func (recv *TreeViewColumn) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by TreeViewColumn
func (recv *TreeViewColumn) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// UIManager is a wrapper around the C record GtkUIManager.
type UIManager struct {
	native unsafe.Pointer
	// parent : record
	// Private : private_data
}

// Buildable returns the Buildable interface implemented by UIManager
func (recv *UIManager) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// VBox is a wrapper around the C record GtkVBox.
type VBox struct {
	native unsafe.Pointer
	// box : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VBox
func (recv *VBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VBox
func (recv *VBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VBox
func (recv *VBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// VButtonBox is a wrapper around the C record GtkVButtonBox.
type VButtonBox struct {
	native unsafe.Pointer
	// button_box : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VButtonBox
func (recv *VButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VButtonBox
func (recv *VButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VButtonBox
func (recv *VButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// VPaned is a wrapper around the C record GtkVPaned.
type VPaned struct {
	native unsafe.Pointer
	// paned : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VPaned
func (recv *VPaned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VPaned
func (recv *VPaned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VPaned
func (recv *VPaned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// VScale is a wrapper around the C record GtkVScale.
type VScale struct {
	native unsafe.Pointer
	// scale : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VScale
func (recv *VScale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VScale
func (recv *VScale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VScale
func (recv *VScale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// VScrollbar is a wrapper around the C record GtkVScrollbar.
type VScrollbar struct {
	native unsafe.Pointer
	// scrollbar : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VScrollbar
func (recv *VScrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VScrollbar
func (recv *VScrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VScrollbar
func (recv *VScrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// VSeparator is a wrapper around the C record GtkVSeparator.
type VSeparator struct {
	native unsafe.Pointer
	// separator : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VSeparator
func (recv *VSeparator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VSeparator
func (recv *VSeparator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VSeparator
func (recv *VSeparator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Viewport is a wrapper around the C record GtkViewport.
type Viewport struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Viewport
func (recv *Viewport) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Viewport
func (recv *Viewport) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by Viewport
func (recv *Viewport) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// VolumeButton is a wrapper around the C record GtkVolumeButton.
type VolumeButton struct {
	native unsafe.Pointer
	// parent : record
}

// ImplementorIface returns the ImplementorIface interface implemented by VolumeButton
func (recv *VolumeButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by VolumeButton
func (recv *VolumeButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by VolumeButton
func (recv *VolumeButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VolumeButton
func (recv *VolumeButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VolumeButton
func (recv *VolumeButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Widget is a wrapper around the C record GtkWidget.
type Widget struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ImplementorIface returns the ImplementorIface interface implemented by Widget
func (recv *Widget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Widget
func (recv *Widget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// WidgetAccessible is a wrapper around the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by WidgetAccessible
func (recv *WidgetAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Window is a wrapper around the C record GtkWindow.
type Window struct {
	native unsafe.Pointer
	// bin : record
	// priv : record
}

// ImplementorIface returns the ImplementorIface interface implemented by Window
func (recv *Window) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Window
func (recv *Window) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// WindowAccessible is a wrapper around the C record GtkWindowAccessible.
type WindowAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

// Component returns the Component interface implemented by WindowAccessible
func (recv *WindowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Window returns the Window interface implemented by WindowAccessible
func (recv *WindowAccessible) Window() *atk.Window {
	return atk.WindowNewFromC(recv.ToC())
}

// WindowGroup is a wrapper around the C record GtkWindowGroup.
type WindowGroup struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}
