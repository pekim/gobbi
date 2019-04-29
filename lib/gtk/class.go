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

func AboutDialogNewFromC(u unsafe.Pointer) *AboutDialog {
	if u == nil {
		return nil
	}

	g := &AboutDialog{native: u}

	return g
}

func (recv *AboutDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AccelGroupNewFromC(u unsafe.Pointer) *AccelGroup {
	if u == nil {
		return nil
	}

	g := &AccelGroup{native: u}

	return g
}

func (recv *AccelGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelLabel is a wrapper around the C record GtkAccelLabel.
type AccelLabel struct {
	native unsafe.Pointer
	// label : record
	// priv : record
}

func AccelLabelNewFromC(u unsafe.Pointer) *AccelLabel {
	if u == nil {
		return nil
	}

	g := &AccelLabel{native: u}

	return g
}

func (recv *AccelLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AccelMapNewFromC(u unsafe.Pointer) *AccelMap {
	if u == nil {
		return nil
	}

	g := &AccelMap{native: u}

	return g
}

func (recv *AccelMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Accessible is a wrapper around the C record GtkAccessible.
type Accessible struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func AccessibleNewFromC(u unsafe.Pointer) *Accessible {
	if u == nil {
		return nil
	}

	g := &Accessible{native: u}

	return g
}

func (recv *Accessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action is a wrapper around the C record GtkAction.
type Action struct {
	native unsafe.Pointer
	// object : record
	// Private : private_data
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

// Buildable returns the Buildable interface implemented by Action
func (recv *Action) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ActionBar is a wrapper around the C record GtkActionBar.
type ActionBar struct {
	native unsafe.Pointer
	// Private : bin
}

func ActionBarNewFromC(u unsafe.Pointer) *ActionBar {
	if u == nil {
		return nil
	}

	g := &ActionBar{native: u}

	return g
}

func (recv *ActionBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	if u == nil {
		return nil
	}

	g := &ActionGroup{native: u}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AdjustmentNewFromC(u unsafe.Pointer) *Adjustment {
	if u == nil {
		return nil
	}

	g := &Adjustment{native: u}

	return g
}

func (recv *Adjustment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Alignment is a wrapper around the C record GtkAlignment.
type Alignment struct {
	native unsafe.Pointer
	// bin : record
	// Private : priv
}

func AlignmentNewFromC(u unsafe.Pointer) *Alignment {
	if u == nil {
		return nil
	}

	g := &Alignment{native: u}

	return g
}

func (recv *Alignment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AppChooserButtonNewFromC(u unsafe.Pointer) *AppChooserButton {
	if u == nil {
		return nil
	}

	g := &AppChooserButton{native: u}

	return g
}

func (recv *AppChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AppChooserDialogNewFromC(u unsafe.Pointer) *AppChooserDialog {
	if u == nil {
		return nil
	}

	g := &AppChooserDialog{native: u}

	return g
}

func (recv *AppChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AppChooserWidgetNewFromC(u unsafe.Pointer) *AppChooserWidget {
	if u == nil {
		return nil
	}

	g := &AppChooserWidget{native: u}

	return g
}

func (recv *AppChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	if u == nil {
		return nil
	}

	g := &Application{native: u}

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ApplicationWindowNewFromC(u unsafe.Pointer) *ApplicationWindow {
	if u == nil {
		return nil
	}

	g := &ApplicationWindow{native: u}

	return g
}

func (recv *ApplicationWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ArrowNewFromC(u unsafe.Pointer) *Arrow {
	if u == nil {
		return nil
	}

	g := &Arrow{native: u}

	return g
}

func (recv *Arrow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ArrowAccessibleNewFromC(u unsafe.Pointer) *ArrowAccessible {
	if u == nil {
		return nil
	}

	g := &ArrowAccessible{native: u}

	return g
}

func (recv *ArrowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AspectFrameNewFromC(u unsafe.Pointer) *AspectFrame {
	if u == nil {
		return nil
	}

	g := &AspectFrame{native: u}

	return g
}

func (recv *AspectFrame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func AssistantNewFromC(u unsafe.Pointer) *Assistant {
	if u == nil {
		return nil
	}

	g := &Assistant{native: u}

	return g
}

func (recv *Assistant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func BinNewFromC(u unsafe.Pointer) *Bin {
	if u == nil {
		return nil
	}

	g := &Bin{native: u}

	return g
}

func (recv *Bin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func BooleanCellAccessibleNewFromC(u unsafe.Pointer) *BooleanCellAccessible {
	if u == nil {
		return nil
	}

	g := &BooleanCellAccessible{native: u}

	return g
}

func (recv *BooleanCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func BoxNewFromC(u unsafe.Pointer) *Box {
	if u == nil {
		return nil
	}

	g := &Box{native: u}

	return g
}

func (recv *Box) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func BuilderNewFromC(u unsafe.Pointer) *Builder {
	if u == nil {
		return nil
	}

	g := &Builder{native: u}

	return g
}

func (recv *Builder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button is a wrapper around the C record GtkButton.
type Button struct {
	native unsafe.Pointer
	// Private : bin
	// Private : priv
}

func ButtonNewFromC(u unsafe.Pointer) *Button {
	if u == nil {
		return nil
	}

	g := &Button{native: u}

	return g
}

func (recv *Button) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ButtonAccessibleNewFromC(u unsafe.Pointer) *ButtonAccessible {
	if u == nil {
		return nil
	}

	g := &ButtonAccessible{native: u}

	return g
}

func (recv *ButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ButtonBoxNewFromC(u unsafe.Pointer) *ButtonBox {
	if u == nil {
		return nil
	}

	g := &ButtonBox{native: u}

	return g
}

func (recv *ButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CalendarNewFromC(u unsafe.Pointer) *Calendar {
	if u == nil {
		return nil
	}

	g := &Calendar{native: u}

	return g
}

func (recv *Calendar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CellAccessibleNewFromC(u unsafe.Pointer) *CellAccessible {
	if u == nil {
		return nil
	}

	g := &CellAccessible{native: u}

	return g
}

func (recv *CellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CellAreaNewFromC(u unsafe.Pointer) *CellArea {
	if u == nil {
		return nil
	}

	g := &CellArea{native: u}

	return g
}

func (recv *CellArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CellAreaBoxNewFromC(u unsafe.Pointer) *CellAreaBox {
	if u == nil {
		return nil
	}

	g := &CellAreaBox{native: u}

	return g
}

func (recv *CellAreaBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CellAreaContextNewFromC(u unsafe.Pointer) *CellAreaContext {
	if u == nil {
		return nil
	}

	g := &CellAreaContext{native: u}

	return g
}

func (recv *CellAreaContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer is a wrapper around the C record GtkCellRenderer.
type CellRenderer struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func CellRendererNewFromC(u unsafe.Pointer) *CellRenderer {
	if u == nil {
		return nil
	}

	g := &CellRenderer{native: u}

	return g
}

func (recv *CellRenderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererAccel is a wrapper around the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererAccelNewFromC(u unsafe.Pointer) *CellRendererAccel {
	if u == nil {
		return nil
	}

	g := &CellRendererAccel{native: u}

	return g
}

func (recv *CellRendererAccel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererCombo is a wrapper around the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererComboNewFromC(u unsafe.Pointer) *CellRendererCombo {
	if u == nil {
		return nil
	}

	g := &CellRendererCombo{native: u}

	return g
}

func (recv *CellRendererCombo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbuf is a wrapper around the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererPixbufNewFromC(u unsafe.Pointer) *CellRendererPixbuf {
	if u == nil {
		return nil
	}

	g := &CellRendererPixbuf{native: u}

	return g
}

func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererProgress is a wrapper around the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func CellRendererProgressNewFromC(u unsafe.Pointer) *CellRendererProgress {
	if u == nil {
		return nil
	}

	g := &CellRendererProgress{native: u}

	return g
}

func (recv *CellRendererProgress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CellRendererSpinNewFromC(u unsafe.Pointer) *CellRendererSpin {
	if u == nil {
		return nil
	}

	g := &CellRendererSpin{native: u}

	return g
}

func (recv *CellRendererSpin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinner is a wrapper around the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererSpinnerNewFromC(u unsafe.Pointer) *CellRendererSpinner {
	if u == nil {
		return nil
	}

	g := &CellRendererSpinner{native: u}

	return g
}

func (recv *CellRendererSpinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererText is a wrapper around the C record GtkCellRendererText.
type CellRendererText struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererTextNewFromC(u unsafe.Pointer) *CellRendererText {
	if u == nil {
		return nil
	}

	g := &CellRendererText{native: u}

	return g
}

func (recv *CellRendererText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererToggle is a wrapper around the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func CellRendererToggleNewFromC(u unsafe.Pointer) *CellRendererToggle {
	if u == nil {
		return nil
	}

	g := &CellRendererToggle{native: u}

	return g
}

func (recv *CellRendererToggle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellView is a wrapper around the C record GtkCellView.
type CellView struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func CellViewNewFromC(u unsafe.Pointer) *CellView {
	if u == nil {
		return nil
	}

	g := &CellView{native: u}

	return g
}

func (recv *CellView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CheckButtonNewFromC(u unsafe.Pointer) *CheckButton {
	if u == nil {
		return nil
	}

	g := &CheckButton{native: u}

	return g
}

func (recv *CheckButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CheckMenuItemNewFromC(u unsafe.Pointer) *CheckMenuItem {
	if u == nil {
		return nil
	}

	g := &CheckMenuItem{native: u}

	return g
}

func (recv *CheckMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CheckMenuItemAccessibleNewFromC(u unsafe.Pointer) *CheckMenuItemAccessible {
	if u == nil {
		return nil
	}

	g := &CheckMenuItemAccessible{native: u}

	return g
}

func (recv *CheckMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ClipboardNewFromC(u unsafe.Pointer) *Clipboard {
	if u == nil {
		return nil
	}

	g := &Clipboard{native: u}

	return g
}

func (recv *Clipboard) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorButton is a wrapper around the C record GtkColorButton.
type ColorButton struct {
	native unsafe.Pointer
	// button : record
	// Private : priv
}

func ColorButtonNewFromC(u unsafe.Pointer) *ColorButton {
	if u == nil {
		return nil
	}

	g := &ColorButton{native: u}

	return g
}

func (recv *ColorButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ColorChooserDialogNewFromC(u unsafe.Pointer) *ColorChooserDialog {
	if u == nil {
		return nil
	}

	g := &ColorChooserDialog{native: u}

	return g
}

func (recv *ColorChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ColorChooserWidgetNewFromC(u unsafe.Pointer) *ColorChooserWidget {
	if u == nil {
		return nil
	}

	g := &ColorChooserWidget{native: u}

	return g
}

func (recv *ColorChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ColorSelectionNewFromC(u unsafe.Pointer) *ColorSelection {
	if u == nil {
		return nil
	}

	g := &ColorSelection{native: u}

	return g
}

func (recv *ColorSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ColorSelectionDialogNewFromC(u unsafe.Pointer) *ColorSelectionDialog {
	if u == nil {
		return nil
	}

	g := &ColorSelectionDialog{native: u}

	return g
}

func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ComboBoxNewFromC(u unsafe.Pointer) *ComboBox {
	if u == nil {
		return nil
	}

	g := &ComboBox{native: u}

	return g
}

func (recv *ComboBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ComboBoxAccessibleNewFromC(u unsafe.Pointer) *ComboBoxAccessible {
	if u == nil {
		return nil
	}

	g := &ComboBoxAccessible{native: u}

	return g
}

func (recv *ComboBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ComboBoxTextNewFromC(u unsafe.Pointer) *ComboBoxText {
	if u == nil {
		return nil
	}

	g := &ComboBoxText{native: u}

	return g
}

func (recv *ComboBoxText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ContainerNewFromC(u unsafe.Pointer) *Container {
	if u == nil {
		return nil
	}

	g := &Container{native: u}

	return g
}

func (recv *Container) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ContainerAccessibleNewFromC(u unsafe.Pointer) *ContainerAccessible {
	if u == nil {
		return nil
	}

	g := &ContainerAccessible{native: u}

	return g
}

func (recv *ContainerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ContainerCellAccessibleNewFromC(u unsafe.Pointer) *ContainerCellAccessible {
	if u == nil {
		return nil
	}

	g := &ContainerCellAccessible{native: u}

	return g
}

func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CssProviderNewFromC(u unsafe.Pointer) *CssProvider {
	if u == nil {
		return nil
	}

	g := &CssProvider{native: u}

	return g
}

func (recv *CssProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func DialogNewFromC(u unsafe.Pointer) *Dialog {
	if u == nil {
		return nil
	}

	g := &Dialog{native: u}

	return g
}

func (recv *Dialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func DrawingAreaNewFromC(u unsafe.Pointer) *DrawingArea {
	if u == nil {
		return nil
	}

	g := &DrawingArea{native: u}

	return g
}

func (recv *DrawingArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EntryNewFromC(u unsafe.Pointer) *Entry {
	if u == nil {
		return nil
	}

	g := &Entry{native: u}

	return g
}

func (recv *Entry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EntryAccessibleNewFromC(u unsafe.Pointer) *EntryAccessible {
	if u == nil {
		return nil
	}

	g := &EntryAccessible{native: u}

	return g
}

func (recv *EntryAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EntryBufferNewFromC(u unsafe.Pointer) *EntryBuffer {
	if u == nil {
		return nil
	}

	g := &EntryBuffer{native: u}

	return g
}

func (recv *EntryBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryCompletion is a wrapper around the C record GtkEntryCompletion.
type EntryCompletion struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func EntryCompletionNewFromC(u unsafe.Pointer) *EntryCompletion {
	if u == nil {
		return nil
	}

	g := &EntryCompletion{native: u}

	return g
}

func (recv *EntryCompletion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EventBoxNewFromC(u unsafe.Pointer) *EventBox {
	if u == nil {
		return nil
	}

	g := &EventBox{native: u}

	return g
}

func (recv *EventBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EventControllerNewFromC(u unsafe.Pointer) *EventController {
	if u == nil {
		return nil
	}

	g := &EventController{native: u}

	return g
}

func (recv *EventController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Expander is a wrapper around the C record GtkExpander.
type Expander struct {
	native unsafe.Pointer
	// bin : record
	// priv : record
}

func ExpanderNewFromC(u unsafe.Pointer) *Expander {
	if u == nil {
		return nil
	}

	g := &Expander{native: u}

	return g
}

func (recv *Expander) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ExpanderAccessibleNewFromC(u unsafe.Pointer) *ExpanderAccessible {
	if u == nil {
		return nil
	}

	g := &ExpanderAccessible{native: u}

	return g
}

func (recv *ExpanderAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileChooserButtonNewFromC(u unsafe.Pointer) *FileChooserButton {
	if u == nil {
		return nil
	}

	g := &FileChooserButton{native: u}

	return g
}

func (recv *FileChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileChooserDialogNewFromC(u unsafe.Pointer) *FileChooserDialog {
	if u == nil {
		return nil
	}

	g := &FileChooserDialog{native: u}

	return g
}

func (recv *FileChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileChooserWidgetNewFromC(u unsafe.Pointer) *FileChooserWidget {
	if u == nil {
		return nil
	}

	g := &FileChooserWidget{native: u}

	return g
}

func (recv *FileChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileFilterNewFromC(u unsafe.Pointer) *FileFilter {
	if u == nil {
		return nil
	}

	g := &FileFilter{native: u}

	return g
}

func (recv *FileFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FixedNewFromC(u unsafe.Pointer) *Fixed {
	if u == nil {
		return nil
	}

	g := &Fixed{native: u}

	return g
}

func (recv *Fixed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FlowBoxNewFromC(u unsafe.Pointer) *FlowBox {
	if u == nil {
		return nil
	}

	g := &FlowBox{native: u}

	return g
}

func (recv *FlowBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FlowBoxAccessibleNewFromC(u unsafe.Pointer) *FlowBoxAccessible {
	if u == nil {
		return nil
	}

	g := &FlowBoxAccessible{native: u}

	return g
}

func (recv *FlowBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FlowBoxChildNewFromC(u unsafe.Pointer) *FlowBoxChild {
	if u == nil {
		return nil
	}

	g := &FlowBoxChild{native: u}

	return g
}

func (recv *FlowBoxChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FlowBoxChildAccessibleNewFromC(u unsafe.Pointer) *FlowBoxChildAccessible {
	if u == nil {
		return nil
	}

	g := &FlowBoxChildAccessible{native: u}

	return g
}

func (recv *FlowBoxChildAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontButtonNewFromC(u unsafe.Pointer) *FontButton {
	if u == nil {
		return nil
	}

	g := &FontButton{native: u}

	return g
}

func (recv *FontButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontChooserDialogNewFromC(u unsafe.Pointer) *FontChooserDialog {
	if u == nil {
		return nil
	}

	g := &FontChooserDialog{native: u}

	return g
}

func (recv *FontChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontChooserWidgetNewFromC(u unsafe.Pointer) *FontChooserWidget {
	if u == nil {
		return nil
	}

	g := &FontChooserWidget{native: u}

	return g
}

func (recv *FontChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontSelectionNewFromC(u unsafe.Pointer) *FontSelection {
	if u == nil {
		return nil
	}

	g := &FontSelection{native: u}

	return g
}

func (recv *FontSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontSelectionDialogNewFromC(u unsafe.Pointer) *FontSelectionDialog {
	if u == nil {
		return nil
	}

	g := &FontSelectionDialog{native: u}

	return g
}

func (recv *FontSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FrameNewFromC(u unsafe.Pointer) *Frame {
	if u == nil {
		return nil
	}

	g := &Frame{native: u}

	return g
}

func (recv *Frame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FrameAccessibleNewFromC(u unsafe.Pointer) *FrameAccessible {
	if u == nil {
		return nil
	}

	g := &FrameAccessible{native: u}

	return g
}

func (recv *FrameAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component returns the Component interface implemented by FrameAccessible
func (recv *FrameAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Gesture is a wrapper around the C record GtkGesture.
type Gesture struct {
	native unsafe.Pointer
}

func GestureNewFromC(u unsafe.Pointer) *Gesture {
	if u == nil {
		return nil
	}

	g := &Gesture{native: u}

	return g
}

func (recv *Gesture) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureDrag is a wrapper around the C record GtkGestureDrag.
type GestureDrag struct {
	native unsafe.Pointer
}

func GestureDragNewFromC(u unsafe.Pointer) *GestureDrag {
	if u == nil {
		return nil
	}

	g := &GestureDrag{native: u}

	return g
}

func (recv *GestureDrag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureLongPress is a wrapper around the C record GtkGestureLongPress.
type GestureLongPress struct {
	native unsafe.Pointer
}

func GestureLongPressNewFromC(u unsafe.Pointer) *GestureLongPress {
	if u == nil {
		return nil
	}

	g := &GestureLongPress{native: u}

	return g
}

func (recv *GestureLongPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureMultiPress is a wrapper around the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native unsafe.Pointer
}

func GestureMultiPressNewFromC(u unsafe.Pointer) *GestureMultiPress {
	if u == nil {
		return nil
	}

	g := &GestureMultiPress{native: u}

	return g
}

func (recv *GestureMultiPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GesturePan is a wrapper around the C record GtkGesturePan.
type GesturePan struct {
	native unsafe.Pointer
}

func GesturePanNewFromC(u unsafe.Pointer) *GesturePan {
	if u == nil {
		return nil
	}

	g := &GesturePan{native: u}

	return g
}

func (recv *GesturePan) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureRotate is a wrapper around the C record GtkGestureRotate.
type GestureRotate struct {
	native unsafe.Pointer
}

func GestureRotateNewFromC(u unsafe.Pointer) *GestureRotate {
	if u == nil {
		return nil
	}

	g := &GestureRotate{native: u}

	return g
}

func (recv *GestureRotate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle is a wrapper around the C record GtkGestureSingle.
type GestureSingle struct {
	native unsafe.Pointer
}

func GestureSingleNewFromC(u unsafe.Pointer) *GestureSingle {
	if u == nil {
		return nil
	}

	g := &GestureSingle{native: u}

	return g
}

func (recv *GestureSingle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSwipe is a wrapper around the C record GtkGestureSwipe.
type GestureSwipe struct {
	native unsafe.Pointer
}

func GestureSwipeNewFromC(u unsafe.Pointer) *GestureSwipe {
	if u == nil {
		return nil
	}

	g := &GestureSwipe{native: u}

	return g
}

func (recv *GestureSwipe) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureZoom is a wrapper around the C record GtkGestureZoom.
type GestureZoom struct {
	native unsafe.Pointer
}

func GestureZoomNewFromC(u unsafe.Pointer) *GestureZoom {
	if u == nil {
		return nil
	}

	g := &GestureZoom{native: u}

	return g
}

func (recv *GestureZoom) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Grid is a wrapper around the C record GtkGrid.
type Grid struct {
	native unsafe.Pointer
	// Private : container
	// Private : priv
}

func GridNewFromC(u unsafe.Pointer) *Grid {
	if u == nil {
		return nil
	}

	g := &Grid{native: u}

	return g
}

func (recv *Grid) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HBoxNewFromC(u unsafe.Pointer) *HBox {
	if u == nil {
		return nil
	}

	g := &HBox{native: u}

	return g
}

func (recv *HBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HButtonBoxNewFromC(u unsafe.Pointer) *HButtonBox {
	if u == nil {
		return nil
	}

	g := &HButtonBox{native: u}

	return g
}

func (recv *HButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HPanedNewFromC(u unsafe.Pointer) *HPaned {
	if u == nil {
		return nil
	}

	g := &HPaned{native: u}

	return g
}

func (recv *HPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HSVNewFromC(u unsafe.Pointer) *HSV {
	if u == nil {
		return nil
	}

	g := &HSV{native: u}

	return g
}

func (recv *HSV) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HScaleNewFromC(u unsafe.Pointer) *HScale {
	if u == nil {
		return nil
	}

	g := &HScale{native: u}

	return g
}

func (recv *HScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HScrollbarNewFromC(u unsafe.Pointer) *HScrollbar {
	if u == nil {
		return nil
	}

	g := &HScrollbar{native: u}

	return g
}

func (recv *HScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HSeparatorNewFromC(u unsafe.Pointer) *HSeparator {
	if u == nil {
		return nil
	}

	g := &HSeparator{native: u}

	return g
}

func (recv *HSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HandleBoxNewFromC(u unsafe.Pointer) *HandleBox {
	if u == nil {
		return nil
	}

	g := &HandleBox{native: u}

	return g
}

func (recv *HandleBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HeaderBarNewFromC(u unsafe.Pointer) *HeaderBar {
	if u == nil {
		return nil
	}

	g := &HeaderBar{native: u}

	return g
}

func (recv *HeaderBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func IMContextNewFromC(u unsafe.Pointer) *IMContext {
	if u == nil {
		return nil
	}

	g := &IMContext{native: u}

	return g
}

func (recv *IMContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimple is a wrapper around the C record GtkIMContextSimple.
type IMContextSimple struct {
	native unsafe.Pointer
	// object : record
	// Private : priv
}

func IMContextSimpleNewFromC(u unsafe.Pointer) *IMContextSimple {
	if u == nil {
		return nil
	}

	g := &IMContextSimple{native: u}

	return g
}

func (recv *IMContextSimple) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMMulticontext is a wrapper around the C record GtkIMMulticontext.
type IMMulticontext struct {
	native unsafe.Pointer
	// object : record
	// Private : priv
}

func IMMulticontextNewFromC(u unsafe.Pointer) *IMMulticontext {
	if u == nil {
		return nil
	}

	g := &IMMulticontext{native: u}

	return g
}

func (recv *IMMulticontext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconFactory is a wrapper around the C record GtkIconFactory.
type IconFactory struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func IconFactoryNewFromC(u unsafe.Pointer) *IconFactory {
	if u == nil {
		return nil
	}

	g := &IconFactory{native: u}

	return g
}

func (recv *IconFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Buildable returns the Buildable interface implemented by IconFactory
func (recv *IconFactory) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// IconInfo is a wrapper around the C record GtkIconInfo.
type IconInfo struct {
	native unsafe.Pointer
}

func IconInfoNewFromC(u unsafe.Pointer) *IconInfo {
	if u == nil {
		return nil
	}

	g := &IconInfo{native: u}

	return g
}

func (recv *IconInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconTheme is a wrapper around the C record GtkIconTheme.
type IconTheme struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func IconThemeNewFromC(u unsafe.Pointer) *IconTheme {
	if u == nil {
		return nil
	}

	g := &IconTheme{native: u}

	return g
}

func (recv *IconTheme) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconView is a wrapper around the C record GtkIconView.
type IconView struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func IconViewNewFromC(u unsafe.Pointer) *IconView {
	if u == nil {
		return nil
	}

	g := &IconView{native: u}

	return g
}

func (recv *IconView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func IconViewAccessibleNewFromC(u unsafe.Pointer) *IconViewAccessible {
	if u == nil {
		return nil
	}

	g := &IconViewAccessible{native: u}

	return g
}

func (recv *IconViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ImageAccessibleNewFromC(u unsafe.Pointer) *ImageAccessible {
	if u == nil {
		return nil
	}

	g := &ImageAccessible{native: u}

	return g
}

func (recv *ImageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ImageCellAccessibleNewFromC(u unsafe.Pointer) *ImageCellAccessible {
	if u == nil {
		return nil
	}

	g := &ImageCellAccessible{native: u}

	return g
}

func (recv *ImageCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ImageMenuItemNewFromC(u unsafe.Pointer) *ImageMenuItem {
	if u == nil {
		return nil
	}

	g := &ImageMenuItem{native: u}

	return g
}

func (recv *ImageMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func InfoBarNewFromC(u unsafe.Pointer) *InfoBar {
	if u == nil {
		return nil
	}

	g := &InfoBar{native: u}

	return g
}

func (recv *InfoBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func InvisibleNewFromC(u unsafe.Pointer) *Invisible {
	if u == nil {
		return nil
	}

	g := &Invisible{native: u}

	return g
}

func (recv *Invisible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LabelNewFromC(u unsafe.Pointer) *Label {
	if u == nil {
		return nil
	}

	g := &Label{native: u}

	return g
}

func (recv *Label) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LabelAccessibleNewFromC(u unsafe.Pointer) *LabelAccessible {
	if u == nil {
		return nil
	}

	g := &LabelAccessible{native: u}

	return g
}

func (recv *LabelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	if u == nil {
		return nil
	}

	g := &Layout{native: u}

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LevelBarNewFromC(u unsafe.Pointer) *LevelBar {
	if u == nil {
		return nil
	}

	g := &LevelBar{native: u}

	return g
}

func (recv *LevelBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LevelBarAccessibleNewFromC(u unsafe.Pointer) *LevelBarAccessible {
	if u == nil {
		return nil
	}

	g := &LevelBarAccessible{native: u}

	return g
}

func (recv *LevelBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LinkButtonNewFromC(u unsafe.Pointer) *LinkButton {
	if u == nil {
		return nil
	}

	g := &LinkButton{native: u}

	return g
}

func (recv *LinkButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LinkButtonAccessibleNewFromC(u unsafe.Pointer) *LinkButtonAccessible {
	if u == nil {
		return nil
	}

	g := &LinkButtonAccessible{native: u}

	return g
}

func (recv *LinkButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ListBoxNewFromC(u unsafe.Pointer) *ListBox {
	if u == nil {
		return nil
	}

	g := &ListBox{native: u}

	return g
}

func (recv *ListBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ListBoxAccessibleNewFromC(u unsafe.Pointer) *ListBoxAccessible {
	if u == nil {
		return nil
	}

	g := &ListBoxAccessible{native: u}

	return g
}

func (recv *ListBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ListBoxRowNewFromC(u unsafe.Pointer) *ListBoxRow {
	if u == nil {
		return nil
	}

	g := &ListBoxRow{native: u}

	return g
}

func (recv *ListBoxRow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ListBoxRowAccessibleNewFromC(u unsafe.Pointer) *ListBoxRowAccessible {
	if u == nil {
		return nil
	}

	g := &ListBoxRowAccessible{native: u}

	return g
}

func (recv *ListBoxRowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	if u == nil {
		return nil
	}

	g := &ListStore{native: u}

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LockButtonNewFromC(u unsafe.Pointer) *LockButton {
	if u == nil {
		return nil
	}

	g := &LockButton{native: u}

	return g
}

func (recv *LockButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LockButtonAccessibleNewFromC(u unsafe.Pointer) *LockButtonAccessible {
	if u == nil {
		return nil
	}

	g := &LockButtonAccessible{native: u}

	return g
}

func (recv *LockButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuNewFromC(u unsafe.Pointer) *Menu {
	if u == nil {
		return nil
	}

	g := &Menu{native: u}

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuAccessibleNewFromC(u unsafe.Pointer) *MenuAccessible {
	if u == nil {
		return nil
	}

	g := &MenuAccessible{native: u}

	return g
}

func (recv *MenuAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuBarNewFromC(u unsafe.Pointer) *MenuBar {
	if u == nil {
		return nil
	}

	g := &MenuBar{native: u}

	return g
}

func (recv *MenuBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuButtonNewFromC(u unsafe.Pointer) *MenuButton {
	if u == nil {
		return nil
	}

	g := &MenuButton{native: u}

	return g
}

func (recv *MenuButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuButtonAccessibleNewFromC(u unsafe.Pointer) *MenuButtonAccessible {
	if u == nil {
		return nil
	}

	g := &MenuButtonAccessible{native: u}

	return g
}

func (recv *MenuButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	if u == nil {
		return nil
	}

	g := &MenuItem{native: u}

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuItemAccessibleNewFromC(u unsafe.Pointer) *MenuItemAccessible {
	if u == nil {
		return nil
	}

	g := &MenuItemAccessible{native: u}

	return g
}

func (recv *MenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuShellNewFromC(u unsafe.Pointer) *MenuShell {
	if u == nil {
		return nil
	}

	g := &MenuShell{native: u}

	return g
}

func (recv *MenuShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuShellAccessibleNewFromC(u unsafe.Pointer) *MenuShellAccessible {
	if u == nil {
		return nil
	}

	g := &MenuShellAccessible{native: u}

	return g
}

func (recv *MenuShellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MenuToolButtonNewFromC(u unsafe.Pointer) *MenuToolButton {
	if u == nil {
		return nil
	}

	g := &MenuToolButton{native: u}

	return g
}

func (recv *MenuToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MessageDialogNewFromC(u unsafe.Pointer) *MessageDialog {
	if u == nil {
		return nil
	}

	g := &MessageDialog{native: u}

	return g
}

func (recv *MessageDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MiscNewFromC(u unsafe.Pointer) *Misc {
	if u == nil {
		return nil
	}

	g := &Misc{native: u}

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ModelButtonNewFromC(u unsafe.Pointer) *ModelButton {
	if u == nil {
		return nil
	}

	g := &ModelButton{native: u}

	return g
}

func (recv *ModelButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	if u == nil {
		return nil
	}

	g := &MountOperation{native: u}

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Notebook is a wrapper around the C record GtkNotebook.
type Notebook struct {
	native unsafe.Pointer
	// Private : container
	// Private : priv
}

func NotebookNewFromC(u unsafe.Pointer) *Notebook {
	if u == nil {
		return nil
	}

	g := &Notebook{native: u}

	return g
}

func (recv *Notebook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func NotebookAccessibleNewFromC(u unsafe.Pointer) *NotebookAccessible {
	if u == nil {
		return nil
	}

	g := &NotebookAccessible{native: u}

	return g
}

func (recv *NotebookAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func NotebookPageAccessibleNewFromC(u unsafe.Pointer) *NotebookPageAccessible {
	if u == nil {
		return nil
	}

	g := &NotebookPageAccessible{native: u}

	return g
}

func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func NumerableIconNewFromC(u unsafe.Pointer) *NumerableIcon {
	if u == nil {
		return nil
	}

	g := &NumerableIcon{native: u}

	return g
}

func (recv *NumerableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func OffscreenWindowNewFromC(u unsafe.Pointer) *OffscreenWindow {
	if u == nil {
		return nil
	}

	g := &OffscreenWindow{native: u}

	return g
}

func (recv *OffscreenWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func OverlayNewFromC(u unsafe.Pointer) *Overlay {
	if u == nil {
		return nil
	}

	g := &Overlay{native: u}

	return g
}

func (recv *Overlay) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PageSetupNewFromC(u unsafe.Pointer) *PageSetup {
	if u == nil {
		return nil
	}

	g := &PageSetup{native: u}

	return g
}

func (recv *PageSetup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Paned is a wrapper around the C record GtkPaned.
type Paned struct {
	native unsafe.Pointer
	// container : record
	// Private : priv
}

func PanedNewFromC(u unsafe.Pointer) *Paned {
	if u == nil {
		return nil
	}

	g := &Paned{native: u}

	return g
}

func (recv *Paned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PanedAccessibleNewFromC(u unsafe.Pointer) *PanedAccessible {
	if u == nil {
		return nil
	}

	g := &PanedAccessible{native: u}

	return g
}

func (recv *PanedAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PlacesSidebarNewFromC(u unsafe.Pointer) *PlacesSidebar {
	if u == nil {
		return nil
	}

	g := &PlacesSidebar{native: u}

	return g
}

func (recv *PlacesSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PopoverNewFromC(u unsafe.Pointer) *Popover {
	if u == nil {
		return nil
	}

	g := &Popover{native: u}

	return g
}

func (recv *Popover) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PopoverAccessibleNewFromC(u unsafe.Pointer) *PopoverAccessible {
	if u == nil {
		return nil
	}

	g := &PopoverAccessible{native: u}

	return g
}

func (recv *PopoverAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component returns the Component interface implemented by PopoverAccessible
func (recv *PopoverAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// PopoverMenu is a wrapper around the C record GtkPopoverMenu.
type PopoverMenu struct {
	native unsafe.Pointer
}

func PopoverMenuNewFromC(u unsafe.Pointer) *PopoverMenu {
	if u == nil {
		return nil
	}

	g := &PopoverMenu{native: u}

	return g
}

func (recv *PopoverMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func PrintContextNewFromC(u unsafe.Pointer) *PrintContext {
	if u == nil {
		return nil
	}

	g := &PrintContext{native: u}

	return g
}

func (recv *PrintContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperation is a wrapper around the C record GtkPrintOperation.
type PrintOperation struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func PrintOperationNewFromC(u unsafe.Pointer) *PrintOperation {
	if u == nil {
		return nil
	}

	g := &PrintOperation{native: u}

	return g
}

func (recv *PrintOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperationPreview returns the PrintOperationPreview interface implemented by PrintOperation
func (recv *PrintOperation) PrintOperationPreview() *PrintOperationPreview {
	return PrintOperationPreviewNewFromC(recv.ToC())
}

// PrintSettings is a wrapper around the C record GtkPrintSettings.
type PrintSettings struct {
	native unsafe.Pointer
}

func PrintSettingsNewFromC(u unsafe.Pointer) *PrintSettings {
	if u == nil {
		return nil
	}

	g := &PrintSettings{native: u}

	return g
}

func (recv *PrintSettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBar is a wrapper around the C record GtkProgressBar.
type ProgressBar struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func ProgressBarNewFromC(u unsafe.Pointer) *ProgressBar {
	if u == nil {
		return nil
	}

	g := &ProgressBar{native: u}

	return g
}

func (recv *ProgressBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ProgressBarAccessibleNewFromC(u unsafe.Pointer) *ProgressBarAccessible {
	if u == nil {
		return nil
	}

	g := &ProgressBarAccessible{native: u}

	return g
}

func (recv *ProgressBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioActionNewFromC(u unsafe.Pointer) *RadioAction {
	if u == nil {
		return nil
	}

	g := &RadioAction{native: u}

	return g
}

func (recv *RadioAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioButtonNewFromC(u unsafe.Pointer) *RadioButton {
	if u == nil {
		return nil
	}

	g := &RadioButton{native: u}

	return g
}

func (recv *RadioButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioButtonAccessibleNewFromC(u unsafe.Pointer) *RadioButtonAccessible {
	if u == nil {
		return nil
	}

	g := &RadioButtonAccessible{native: u}

	return g
}

func (recv *RadioButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioMenuItemNewFromC(u unsafe.Pointer) *RadioMenuItem {
	if u == nil {
		return nil
	}

	g := &RadioMenuItem{native: u}

	return g
}

func (recv *RadioMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioMenuItemAccessibleNewFromC(u unsafe.Pointer) *RadioMenuItemAccessible {
	if u == nil {
		return nil
	}

	g := &RadioMenuItemAccessible{native: u}

	return g
}

func (recv *RadioMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RadioToolButtonNewFromC(u unsafe.Pointer) *RadioToolButton {
	if u == nil {
		return nil
	}

	g := &RadioToolButton{native: u}

	return g
}

func (recv *RadioToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RangeNewFromC(u unsafe.Pointer) *Range {
	if u == nil {
		return nil
	}

	g := &Range{native: u}

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RangeAccessibleNewFromC(u unsafe.Pointer) *RangeAccessible {
	if u == nil {
		return nil
	}

	g := &RangeAccessible{native: u}

	return g
}

func (recv *RangeAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RcStyleNewFromC(u unsafe.Pointer) *RcStyle {
	if u == nil {
		return nil
	}

	g := &RcStyle{native: u}

	return g
}

func (recv *RcStyle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentAction is a wrapper around the C record GtkRecentAction.
type RecentAction struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func RecentActionNewFromC(u unsafe.Pointer) *RecentAction {
	if u == nil {
		return nil
	}

	g := &RecentAction{native: u}

	return g
}

func (recv *RecentAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RecentChooserDialogNewFromC(u unsafe.Pointer) *RecentChooserDialog {
	if u == nil {
		return nil
	}

	g := &RecentChooserDialog{native: u}

	return g
}

func (recv *RecentChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RecentChooserMenuNewFromC(u unsafe.Pointer) *RecentChooserMenu {
	if u == nil {
		return nil
	}

	g := &RecentChooserMenu{native: u}

	return g
}

func (recv *RecentChooserMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RecentChooserWidgetNewFromC(u unsafe.Pointer) *RecentChooserWidget {
	if u == nil {
		return nil
	}

	g := &RecentChooserWidget{native: u}

	return g
}

func (recv *RecentChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RecentFilterNewFromC(u unsafe.Pointer) *RecentFilter {
	if u == nil {
		return nil
	}

	g := &RecentFilter{native: u}

	return g
}

func (recv *RecentFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RendererCellAccessibleNewFromC(u unsafe.Pointer) *RendererCellAccessible {
	if u == nil {
		return nil
	}

	g := &RendererCellAccessible{native: u}

	return g
}

func (recv *RendererCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RevealerNewFromC(u unsafe.Pointer) *Revealer {
	if u == nil {
		return nil
	}

	g := &Revealer{native: u}

	return g
}

func (recv *Revealer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScaleNewFromC(u unsafe.Pointer) *Scale {
	if u == nil {
		return nil
	}

	g := &Scale{native: u}

	return g
}

func (recv *Scale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScaleAccessibleNewFromC(u unsafe.Pointer) *ScaleAccessible {
	if u == nil {
		return nil
	}

	g := &ScaleAccessible{native: u}

	return g
}

func (recv *ScaleAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScaleButtonNewFromC(u unsafe.Pointer) *ScaleButton {
	if u == nil {
		return nil
	}

	g := &ScaleButton{native: u}

	return g
}

func (recv *ScaleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScaleButtonAccessibleNewFromC(u unsafe.Pointer) *ScaleButtonAccessible {
	if u == nil {
		return nil
	}

	g := &ScaleButtonAccessible{native: u}

	return g
}

func (recv *ScaleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScrollbarNewFromC(u unsafe.Pointer) *Scrollbar {
	if u == nil {
		return nil
	}

	g := &Scrollbar{native: u}

	return g
}

func (recv *Scrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScrolledWindowNewFromC(u unsafe.Pointer) *ScrolledWindow {
	if u == nil {
		return nil
	}

	g := &ScrolledWindow{native: u}

	return g
}

func (recv *ScrolledWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScrolledWindowAccessibleNewFromC(u unsafe.Pointer) *ScrolledWindowAccessible {
	if u == nil {
		return nil
	}

	g := &ScrolledWindowAccessible{native: u}

	return g
}

func (recv *ScrolledWindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SearchBarNewFromC(u unsafe.Pointer) *SearchBar {
	if u == nil {
		return nil
	}

	g := &SearchBar{native: u}

	return g
}

func (recv *SearchBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SearchEntryNewFromC(u unsafe.Pointer) *SearchEntry {
	if u == nil {
		return nil
	}

	g := &SearchEntry{native: u}

	return g
}

func (recv *SearchEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SeparatorNewFromC(u unsafe.Pointer) *Separator {
	if u == nil {
		return nil
	}

	g := &Separator{native: u}

	return g
}

func (recv *Separator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SeparatorMenuItemNewFromC(u unsafe.Pointer) *SeparatorMenuItem {
	if u == nil {
		return nil
	}

	g := &SeparatorMenuItem{native: u}

	return g
}

func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SeparatorToolItemNewFromC(u unsafe.Pointer) *SeparatorToolItem {
	if u == nil {
		return nil
	}

	g := &SeparatorToolItem{native: u}

	return g
}

func (recv *SeparatorToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	if u == nil {
		return nil
	}

	g := &Settings{native: u}

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SizeGroupNewFromC(u unsafe.Pointer) *SizeGroup {
	if u == nil {
		return nil
	}

	g := &SizeGroup{native: u}

	return g
}

func (recv *SizeGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SpinButtonNewFromC(u unsafe.Pointer) *SpinButton {
	if u == nil {
		return nil
	}

	g := &SpinButton{native: u}

	return g
}

func (recv *SpinButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SpinButtonAccessibleNewFromC(u unsafe.Pointer) *SpinButtonAccessible {
	if u == nil {
		return nil
	}

	g := &SpinButtonAccessible{native: u}

	return g
}

func (recv *SpinButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SpinnerNewFromC(u unsafe.Pointer) *Spinner {
	if u == nil {
		return nil
	}

	g := &Spinner{native: u}

	return g
}

func (recv *Spinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SpinnerAccessibleNewFromC(u unsafe.Pointer) *SpinnerAccessible {
	if u == nil {
		return nil
	}

	g := &SpinnerAccessible{native: u}

	return g
}

func (recv *SpinnerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StackNewFromC(u unsafe.Pointer) *Stack {
	if u == nil {
		return nil
	}

	g := &Stack{native: u}

	return g
}

func (recv *Stack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StackSidebarNewFromC(u unsafe.Pointer) *StackSidebar {
	if u == nil {
		return nil
	}

	g := &StackSidebar{native: u}

	return g
}

func (recv *StackSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StackSwitcherNewFromC(u unsafe.Pointer) *StackSwitcher {
	if u == nil {
		return nil
	}

	g := &StackSwitcher{native: u}

	return g
}

func (recv *StackSwitcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StatusIconNewFromC(u unsafe.Pointer) *StatusIcon {
	if u == nil {
		return nil
	}

	g := &StatusIcon{native: u}

	return g
}

func (recv *StatusIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Statusbar is a wrapper around the C record GtkStatusbar.
type Statusbar struct {
	native unsafe.Pointer
	// parent_widget : record
	// Private : priv
}

func StatusbarNewFromC(u unsafe.Pointer) *Statusbar {
	if u == nil {
		return nil
	}

	g := &Statusbar{native: u}

	return g
}

func (recv *Statusbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StatusbarAccessibleNewFromC(u unsafe.Pointer) *StatusbarAccessible {
	if u == nil {
		return nil
	}

	g := &StatusbarAccessible{native: u}

	return g
}

func (recv *StatusbarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func StyleNewFromC(u unsafe.Pointer) *Style {
	if u == nil {
		return nil
	}

	g := &Style{native: u}

	return g
}

func (recv *Style) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleContext is a wrapper around the C record GtkStyleContext.
type StyleContext struct {
	native unsafe.Pointer
	// parent_object : record
	// priv : record
}

func StyleContextNewFromC(u unsafe.Pointer) *StyleContext {
	if u == nil {
		return nil
	}

	g := &StyleContext{native: u}

	return g
}

func (recv *StyleContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleProperties is a wrapper around the C record GtkStyleProperties.
type StyleProperties struct {
	native unsafe.Pointer
	// Private : parent_object
	// Private : priv
}

func StylePropertiesNewFromC(u unsafe.Pointer) *StyleProperties {
	if u == nil {
		return nil
	}

	g := &StyleProperties{native: u}

	return g
}

func (recv *StyleProperties) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SwitchNewFromC(u unsafe.Pointer) *Switch {
	if u == nil {
		return nil
	}

	g := &Switch{native: u}

	return g
}

func (recv *Switch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SwitchAccessibleNewFromC(u unsafe.Pointer) *SwitchAccessible {
	if u == nil {
		return nil
	}

	g := &SwitchAccessible{native: u}

	return g
}

func (recv *SwitchAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TearoffMenuItemNewFromC(u unsafe.Pointer) *TearoffMenuItem {
	if u == nil {
		return nil
	}

	g := &TearoffMenuItem{native: u}

	return g
}

func (recv *TearoffMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TextBufferNewFromC(u unsafe.Pointer) *TextBuffer {
	if u == nil {
		return nil
	}

	g := &TextBuffer{native: u}

	return g
}

func (recv *TextBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextCellAccessible is a wrapper around the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

func TextCellAccessibleNewFromC(u unsafe.Pointer) *TextCellAccessible {
	if u == nil {
		return nil
	}

	g := &TextCellAccessible{native: u}

	return g
}

func (recv *TextCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TextChildAnchorNewFromC(u unsafe.Pointer) *TextChildAnchor {
	if u == nil {
		return nil
	}

	g := &TextChildAnchor{native: u}

	return g
}

func (recv *TextChildAnchor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextMark is a wrapper around the C record GtkTextMark.
type TextMark struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : segment
}

func TextMarkNewFromC(u unsafe.Pointer) *TextMark {
	if u == nil {
		return nil
	}

	g := &TextMark{native: u}

	return g
}

func (recv *TextMark) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTag is a wrapper around the C record GtkTextTag.
type TextTag struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TextTagNewFromC(u unsafe.Pointer) *TextTag {
	if u == nil {
		return nil
	}

	g := &TextTag{native: u}

	return g
}

func (recv *TextTag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagTable is a wrapper around the C record GtkTextTagTable.
type TextTagTable struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TextTagTableNewFromC(u unsafe.Pointer) *TextTagTable {
	if u == nil {
		return nil
	}

	g := &TextTagTable{native: u}

	return g
}

func (recv *TextTagTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TextViewNewFromC(u unsafe.Pointer) *TextView {
	if u == nil {
		return nil
	}

	g := &TextView{native: u}

	return g
}

func (recv *TextView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TextViewAccessibleNewFromC(u unsafe.Pointer) *TextViewAccessible {
	if u == nil {
		return nil
	}

	g := &TextViewAccessible{native: u}

	return g
}

func (recv *TextViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ThemingEngineNewFromC(u unsafe.Pointer) *ThemingEngine {
	if u == nil {
		return nil
	}

	g := &ThemingEngine{native: u}

	return g
}

func (recv *ThemingEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleAction is a wrapper around the C record GtkToggleAction.
type ToggleAction struct {
	native unsafe.Pointer
	// parent : record
	// Private : private_data
}

func ToggleActionNewFromC(u unsafe.Pointer) *ToggleAction {
	if u == nil {
		return nil
	}

	g := &ToggleAction{native: u}

	return g
}

func (recv *ToggleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToggleButtonNewFromC(u unsafe.Pointer) *ToggleButton {
	if u == nil {
		return nil
	}

	g := &ToggleButton{native: u}

	return g
}

func (recv *ToggleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToggleButtonAccessibleNewFromC(u unsafe.Pointer) *ToggleButtonAccessible {
	if u == nil {
		return nil
	}

	g := &ToggleButtonAccessible{native: u}

	return g
}

func (recv *ToggleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToggleToolButtonNewFromC(u unsafe.Pointer) *ToggleToolButton {
	if u == nil {
		return nil
	}

	g := &ToggleToolButton{native: u}

	return g
}

func (recv *ToggleToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToolButtonNewFromC(u unsafe.Pointer) *ToolButton {
	if u == nil {
		return nil
	}

	g := &ToolButton{native: u}

	return g
}

func (recv *ToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToolItemNewFromC(u unsafe.Pointer) *ToolItem {
	if u == nil {
		return nil
	}

	g := &ToolItem{native: u}

	return g
}

func (recv *ToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToolItemGroupNewFromC(u unsafe.Pointer) *ToolItemGroup {
	if u == nil {
		return nil
	}

	g := &ToolItemGroup{native: u}

	return g
}

func (recv *ToolItemGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToolPaletteNewFromC(u unsafe.Pointer) *ToolPalette {
	if u == nil {
		return nil
	}

	g := &ToolPalette{native: u}

	return g
}

func (recv *ToolPalette) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ToolbarNewFromC(u unsafe.Pointer) *Toolbar {
	if u == nil {
		return nil
	}

	g := &Toolbar{native: u}

	return g
}

func (recv *Toolbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TooltipNewFromC(u unsafe.Pointer) *Tooltip {
	if u == nil {
		return nil
	}

	g := &Tooltip{native: u}

	return g
}

func (recv *Tooltip) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToplevelAccessible is a wrapper around the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

func ToplevelAccessibleNewFromC(u unsafe.Pointer) *ToplevelAccessible {
	if u == nil {
		return nil
	}

	g := &ToplevelAccessible{native: u}

	return g
}

func (recv *ToplevelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelFilter is a wrapper around the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native unsafe.Pointer
	// parent : record
	// Private : priv
}

func TreeModelFilterNewFromC(u unsafe.Pointer) *TreeModelFilter {
	if u == nil {
		return nil
	}

	g := &TreeModelFilter{native: u}

	return g
}

func (recv *TreeModelFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TreeModelSortNewFromC(u unsafe.Pointer) *TreeModelSort {
	if u == nil {
		return nil
	}

	g := &TreeModelSort{native: u}

	return g
}

func (recv *TreeModelSort) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TreeSelectionNewFromC(u unsafe.Pointer) *TreeSelection {
	if u == nil {
		return nil
	}

	g := &TreeSelection{native: u}

	return g
}

func (recv *TreeSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeStore is a wrapper around the C record GtkTreeStore.
type TreeStore struct {
	native unsafe.Pointer
	// parent : record
	// priv : record
}

func TreeStoreNewFromC(u unsafe.Pointer) *TreeStore {
	if u == nil {
		return nil
	}

	g := &TreeStore{native: u}

	return g
}

func (recv *TreeStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TreeViewNewFromC(u unsafe.Pointer) *TreeView {
	if u == nil {
		return nil
	}

	g := &TreeView{native: u}

	return g
}

func (recv *TreeView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TreeViewAccessibleNewFromC(u unsafe.Pointer) *TreeViewAccessible {
	if u == nil {
		return nil
	}

	g := &TreeViewAccessible{native: u}

	return g
}

func (recv *TreeViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TreeViewColumnNewFromC(u unsafe.Pointer) *TreeViewColumn {
	if u == nil {
		return nil
	}

	g := &TreeViewColumn{native: u}

	return g
}

func (recv *TreeViewColumn) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func UIManagerNewFromC(u unsafe.Pointer) *UIManager {
	if u == nil {
		return nil
	}

	g := &UIManager{native: u}

	return g
}

func (recv *UIManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VBoxNewFromC(u unsafe.Pointer) *VBox {
	if u == nil {
		return nil
	}

	g := &VBox{native: u}

	return g
}

func (recv *VBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VButtonBoxNewFromC(u unsafe.Pointer) *VButtonBox {
	if u == nil {
		return nil
	}

	g := &VButtonBox{native: u}

	return g
}

func (recv *VButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VPanedNewFromC(u unsafe.Pointer) *VPaned {
	if u == nil {
		return nil
	}

	g := &VPaned{native: u}

	return g
}

func (recv *VPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VScaleNewFromC(u unsafe.Pointer) *VScale {
	if u == nil {
		return nil
	}

	g := &VScale{native: u}

	return g
}

func (recv *VScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VScrollbarNewFromC(u unsafe.Pointer) *VScrollbar {
	if u == nil {
		return nil
	}

	g := &VScrollbar{native: u}

	return g
}

func (recv *VScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VSeparatorNewFromC(u unsafe.Pointer) *VSeparator {
	if u == nil {
		return nil
	}

	g := &VSeparator{native: u}

	return g
}

func (recv *VSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ViewportNewFromC(u unsafe.Pointer) *Viewport {
	if u == nil {
		return nil
	}

	g := &Viewport{native: u}

	return g
}

func (recv *Viewport) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func VolumeButtonNewFromC(u unsafe.Pointer) *VolumeButton {
	if u == nil {
		return nil
	}

	g := &VolumeButton{native: u}

	return g
}

func (recv *VolumeButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func WidgetNewFromC(u unsafe.Pointer) *Widget {
	if u == nil {
		return nil
	}

	g := &Widget{native: u}

	return g
}

func (recv *Widget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func WidgetAccessibleNewFromC(u unsafe.Pointer) *WidgetAccessible {
	if u == nil {
		return nil
	}

	g := &WidgetAccessible{native: u}

	return g
}

func (recv *WidgetAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func WindowAccessibleNewFromC(u unsafe.Pointer) *WindowAccessible {
	if u == nil {
		return nil
	}

	g := &WindowAccessible{native: u}

	return g
}

func (recv *WindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func WindowGroupNewFromC(u unsafe.Pointer) *WindowGroup {
	if u == nil {
		return nil
	}

	g := &WindowGroup{native: u}

	return g
}

func (recv *WindowGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
