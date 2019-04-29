// This is a generated file - DO NOT EDIT

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// AboutDialogClass is a wrapper around the C record GtkAboutDialogClass.
type AboutDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate_link
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AboutDialogClassNewFromC(u unsafe.Pointer) *AboutDialogClass {
	if u == nil {
		return nil
	}

	g := &AboutDialogClass{native: u}

	return g
}

func (recv *AboutDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AboutDialogPrivate is a wrapper around the C record GtkAboutDialogPrivate.
type AboutDialogPrivate struct {
	native unsafe.Pointer
}

func AboutDialogPrivateNewFromC(u unsafe.Pointer) *AboutDialogPrivate {
	if u == nil {
		return nil
	}

	g := &AboutDialogPrivate{native: u}

	return g
}

func (recv *AboutDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroupClass is a wrapper around the C record GtkAccelGroupClass.
type AccelGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for accel_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AccelGroupClassNewFromC(u unsafe.Pointer) *AccelGroupClass {
	if u == nil {
		return nil
	}

	g := &AccelGroupClass{native: u}

	return g
}

func (recv *AccelGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroupEntry is a wrapper around the C record GtkAccelGroupEntry.
type AccelGroupEntry struct {
	native unsafe.Pointer
	// key : record
	// closure : record
	AccelPathQuark glib.Quark
}

func AccelGroupEntryNewFromC(u unsafe.Pointer) *AccelGroupEntry {
	if u == nil {
		return nil
	}

	g := &AccelGroupEntry{native: u}

	return g
}

func (recv *AccelGroupEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroupPrivate is a wrapper around the C record GtkAccelGroupPrivate.
type AccelGroupPrivate struct {
	native unsafe.Pointer
}

func AccelGroupPrivateNewFromC(u unsafe.Pointer) *AccelGroupPrivate {
	if u == nil {
		return nil
	}

	g := &AccelGroupPrivate{native: u}

	return g
}

func (recv *AccelGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelKey is a wrapper around the C record GtkAccelKey.
type AccelKey struct {
	native    unsafe.Pointer
	AccelKey  uint32
	AccelMods gdk.ModifierType
	// Bitfield not supported : 16 accel_flags
}

func AccelKeyNewFromC(u unsafe.Pointer) *AccelKey {
	if u == nil {
		return nil
	}

	g := &AccelKey{native: u}

	return g
}

func (recv *AccelKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelLabelClass is a wrapper around the C record GtkAccelLabelClass.
type AccelLabelClass struct {
	native unsafe.Pointer
	// parent_class : record
	SignalQuote1   string
	SignalQuote2   string
	ModNameShift   string
	ModNameControl string
	ModNameAlt     string
	ModSeparator   string
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AccelLabelClassNewFromC(u unsafe.Pointer) *AccelLabelClass {
	if u == nil {
		return nil
	}

	g := &AccelLabelClass{native: u}

	return g
}

func (recv *AccelLabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelLabelPrivate is a wrapper around the C record GtkAccelLabelPrivate.
type AccelLabelPrivate struct {
	native unsafe.Pointer
}

func AccelLabelPrivateNewFromC(u unsafe.Pointer) *AccelLabelPrivate {
	if u == nil {
		return nil
	}

	g := &AccelLabelPrivate{native: u}

	return g
}

func (recv *AccelLabelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelMapClass is a wrapper around the C record GtkAccelMapClass.
type AccelMapClass struct {
	native unsafe.Pointer
}

func AccelMapClassNewFromC(u unsafe.Pointer) *AccelMapClass {
	if u == nil {
		return nil
	}

	g := &AccelMapClass{native: u}

	return g
}

func (recv *AccelMapClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccessibleClass is a wrapper around the C record GtkAccessibleClass.
type AccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for connect_widget_destroyed
	// no type for widget_set
	// no type for widget_unset
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AccessibleClassNewFromC(u unsafe.Pointer) *AccessibleClass {
	if u == nil {
		return nil
	}

	g := &AccessibleClass{native: u}

	return g
}

func (recv *AccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccessiblePrivate is a wrapper around the C record GtkAccessiblePrivate.
type AccessiblePrivate struct {
	native unsafe.Pointer
}

func AccessiblePrivateNewFromC(u unsafe.Pointer) *AccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &AccessiblePrivate{native: u}

	return g
}

func (recv *AccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionBarClass is a wrapper around the C record GtkActionBarClass.
type ActionBarClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ActionBarClassNewFromC(u unsafe.Pointer) *ActionBarClass {
	if u == nil {
		return nil
	}

	g := &ActionBarClass{native: u}

	return g
}

func (recv *ActionBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionBarPrivate is a wrapper around the C record GtkActionBarPrivate.
type ActionBarPrivate struct {
	native unsafe.Pointer
}

func ActionBarPrivateNewFromC(u unsafe.Pointer) *ActionBarPrivate {
	if u == nil {
		return nil
	}

	g := &ActionBarPrivate{native: u}

	return g
}

func (recv *ActionBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionClass is a wrapper around the C record GtkActionClass.
type ActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// Private : menu_item_type
	// Private : toolbar_item_type
	// no type for create_menu_item
	// no type for create_tool_item
	// no type for connect_proxy
	// no type for disconnect_proxy
	// no type for create_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ActionClassNewFromC(u unsafe.Pointer) *ActionClass {
	if u == nil {
		return nil
	}

	g := &ActionClass{native: u}

	return g
}

func (recv *ActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionEntry is a wrapper around the C record GtkActionEntry.
type ActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
}

func ActionEntryNewFromC(u unsafe.Pointer) *ActionEntry {
	if u == nil {
		return nil
	}

	g := &ActionEntry{native: u}

	return g
}

func (recv *ActionEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroupClass is a wrapper around the C record GtkActionGroupClass.
type ActionGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_action
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ActionGroupClassNewFromC(u unsafe.Pointer) *ActionGroupClass {
	if u == nil {
		return nil
	}

	g := &ActionGroupClass{native: u}

	return g
}

func (recv *ActionGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroupPrivate is a wrapper around the C record GtkActionGroupPrivate.
type ActionGroupPrivate struct {
	native unsafe.Pointer
}

func ActionGroupPrivateNewFromC(u unsafe.Pointer) *ActionGroupPrivate {
	if u == nil {
		return nil
	}

	g := &ActionGroupPrivate{native: u}

	return g
}

func (recv *ActionGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionPrivate is a wrapper around the C record GtkActionPrivate.
type ActionPrivate struct {
	native unsafe.Pointer
}

func ActionPrivateNewFromC(u unsafe.Pointer) *ActionPrivate {
	if u == nil {
		return nil
	}

	g := &ActionPrivate{native: u}

	return g
}

func (recv *ActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionableInterface is a wrapper around the C record GtkActionableInterface.
type ActionableInterface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_action_name
	// no type for set_action_name
	// no type for get_action_target_value
	// no type for set_action_target_value
}

func ActionableInterfaceNewFromC(u unsafe.Pointer) *ActionableInterface {
	if u == nil {
		return nil
	}

	g := &ActionableInterface{native: u}

	return g
}

func (recv *ActionableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AdjustmentClass is a wrapper around the C record GtkAdjustmentClass.
type AdjustmentClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AdjustmentClassNewFromC(u unsafe.Pointer) *AdjustmentClass {
	if u == nil {
		return nil
	}

	g := &AdjustmentClass{native: u}

	return g
}

func (recv *AdjustmentClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AdjustmentPrivate is a wrapper around the C record GtkAdjustmentPrivate.
type AdjustmentPrivate struct {
	native unsafe.Pointer
}

func AdjustmentPrivateNewFromC(u unsafe.Pointer) *AdjustmentPrivate {
	if u == nil {
		return nil
	}

	g := &AdjustmentPrivate{native: u}

	return g
}

func (recv *AdjustmentPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AlignmentClass is a wrapper around the C record GtkAlignmentClass.
type AlignmentClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AlignmentClassNewFromC(u unsafe.Pointer) *AlignmentClass {
	if u == nil {
		return nil
	}

	g := &AlignmentClass{native: u}

	return g
}

func (recv *AlignmentClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AlignmentPrivate is a wrapper around the C record GtkAlignmentPrivate.
type AlignmentPrivate struct {
	native unsafe.Pointer
}

func AlignmentPrivateNewFromC(u unsafe.Pointer) *AlignmentPrivate {
	if u == nil {
		return nil
	}

	g := &AlignmentPrivate{native: u}

	return g
}

func (recv *AlignmentPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserButtonClass is a wrapper around the C record GtkAppChooserButtonClass.
type AppChooserButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for custom_item_activated
	// Private : padding
}

func AppChooserButtonClassNewFromC(u unsafe.Pointer) *AppChooserButtonClass {
	if u == nil {
		return nil
	}

	g := &AppChooserButtonClass{native: u}

	return g
}

func (recv *AppChooserButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserButtonPrivate is a wrapper around the C record GtkAppChooserButtonPrivate.
type AppChooserButtonPrivate struct {
	native unsafe.Pointer
}

func AppChooserButtonPrivateNewFromC(u unsafe.Pointer) *AppChooserButtonPrivate {
	if u == nil {
		return nil
	}

	g := &AppChooserButtonPrivate{native: u}

	return g
}

func (recv *AppChooserButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserDialogClass is a wrapper around the C record GtkAppChooserDialogClass.
type AppChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

func AppChooserDialogClassNewFromC(u unsafe.Pointer) *AppChooserDialogClass {
	if u == nil {
		return nil
	}

	g := &AppChooserDialogClass{native: u}

	return g
}

func (recv *AppChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserDialogPrivate is a wrapper around the C record GtkAppChooserDialogPrivate.
type AppChooserDialogPrivate struct {
	native unsafe.Pointer
}

func AppChooserDialogPrivateNewFromC(u unsafe.Pointer) *AppChooserDialogPrivate {
	if u == nil {
		return nil
	}

	g := &AppChooserDialogPrivate{native: u}

	return g
}

func (recv *AppChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserWidgetClass is a wrapper around the C record GtkAppChooserWidgetClass.
type AppChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for application_selected
	// no type for application_activated
	// no type for populate_popup
	// Private : padding
}

func AppChooserWidgetClassNewFromC(u unsafe.Pointer) *AppChooserWidgetClass {
	if u == nil {
		return nil
	}

	g := &AppChooserWidgetClass{native: u}

	return g
}

func (recv *AppChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserWidgetPrivate is a wrapper around the C record GtkAppChooserWidgetPrivate.
type AppChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func AppChooserWidgetPrivateNewFromC(u unsafe.Pointer) *AppChooserWidgetPrivate {
	if u == nil {
		return nil
	}

	g := &AppChooserWidgetPrivate{native: u}

	return g
}

func (recv *AppChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationClass is a wrapper around the C record GtkApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for window_added
	// no type for window_removed
	// Private : padding
}

func ApplicationClassNewFromC(u unsafe.Pointer) *ApplicationClass {
	if u == nil {
		return nil
	}

	g := &ApplicationClass{native: u}

	return g
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationPrivate is a wrapper around the C record GtkApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

func ApplicationPrivateNewFromC(u unsafe.Pointer) *ApplicationPrivate {
	if u == nil {
		return nil
	}

	g := &ApplicationPrivate{native: u}

	return g
}

func (recv *ApplicationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationWindowClass is a wrapper around the C record GtkApplicationWindowClass.
type ApplicationWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

func ApplicationWindowClassNewFromC(u unsafe.Pointer) *ApplicationWindowClass {
	if u == nil {
		return nil
	}

	g := &ApplicationWindowClass{native: u}

	return g
}

func (recv *ApplicationWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationWindowPrivate is a wrapper around the C record GtkApplicationWindowPrivate.
type ApplicationWindowPrivate struct {
	native unsafe.Pointer
}

func ApplicationWindowPrivateNewFromC(u unsafe.Pointer) *ApplicationWindowPrivate {
	if u == nil {
		return nil
	}

	g := &ApplicationWindowPrivate{native: u}

	return g
}

func (recv *ApplicationWindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowAccessibleClass is a wrapper around the C record GtkArrowAccessibleClass.
type ArrowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ArrowAccessibleClassNewFromC(u unsafe.Pointer) *ArrowAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ArrowAccessibleClass{native: u}

	return g
}

func (recv *ArrowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowAccessiblePrivate is a wrapper around the C record GtkArrowAccessiblePrivate.
type ArrowAccessiblePrivate struct {
	native unsafe.Pointer
}

func ArrowAccessiblePrivateNewFromC(u unsafe.Pointer) *ArrowAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ArrowAccessiblePrivate{native: u}

	return g
}

func (recv *ArrowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowClass is a wrapper around the C record GtkArrowClass.
type ArrowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ArrowClassNewFromC(u unsafe.Pointer) *ArrowClass {
	if u == nil {
		return nil
	}

	g := &ArrowClass{native: u}

	return g
}

func (recv *ArrowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowPrivate is a wrapper around the C record GtkArrowPrivate.
type ArrowPrivate struct {
	native unsafe.Pointer
}

func ArrowPrivateNewFromC(u unsafe.Pointer) *ArrowPrivate {
	if u == nil {
		return nil
	}

	g := &ArrowPrivate{native: u}

	return g
}

func (recv *ArrowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AspectFrameClass is a wrapper around the C record GtkAspectFrameClass.
type AspectFrameClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AspectFrameClassNewFromC(u unsafe.Pointer) *AspectFrameClass {
	if u == nil {
		return nil
	}

	g := &AspectFrameClass{native: u}

	return g
}

func (recv *AspectFrameClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AspectFramePrivate is a wrapper around the C record GtkAspectFramePrivate.
type AspectFramePrivate struct {
	native unsafe.Pointer
}

func AspectFramePrivateNewFromC(u unsafe.Pointer) *AspectFramePrivate {
	if u == nil {
		return nil
	}

	g := &AspectFramePrivate{native: u}

	return g
}

func (recv *AspectFramePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AssistantClass is a wrapper around the C record GtkAssistantClass.
type AssistantClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for prepare
	// no type for apply
	// no type for close
	// no type for cancel
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
}

func AssistantClassNewFromC(u unsafe.Pointer) *AssistantClass {
	if u == nil {
		return nil
	}

	g := &AssistantClass{native: u}

	return g
}

func (recv *AssistantClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AssistantPrivate is a wrapper around the C record GtkAssistantPrivate.
type AssistantPrivate struct {
	native unsafe.Pointer
}

func AssistantPrivateNewFromC(u unsafe.Pointer) *AssistantPrivate {
	if u == nil {
		return nil
	}

	g := &AssistantPrivate{native: u}

	return g
}

func (recv *AssistantPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BinClass is a wrapper around the C record GtkBinClass.
type BinClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func BinClassNewFromC(u unsafe.Pointer) *BinClass {
	if u == nil {
		return nil
	}

	g := &BinClass{native: u}

	return g
}

func (recv *BinClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BinPrivate is a wrapper around the C record GtkBinPrivate.
type BinPrivate struct {
	native unsafe.Pointer
}

func BinPrivateNewFromC(u unsafe.Pointer) *BinPrivate {
	if u == nil {
		return nil
	}

	g := &BinPrivate{native: u}

	return g
}

func (recv *BinPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BindingArg is a wrapper around the C record GtkBindingArg.
type BindingArg struct {
	native  unsafe.Pointer
	ArgType gobject.Type
}

func BindingArgNewFromC(u unsafe.Pointer) *BindingArg {
	if u == nil {
		return nil
	}

	g := &BindingArg{native: u}

	return g
}

func (recv *BindingArg) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BindingEntry is a wrapper around the C record GtkBindingEntry.
type BindingEntry struct {
	native    unsafe.Pointer
	Keyval    uint32
	Modifiers gdk.ModifierType
	// binding_set : record
	// Bitfield not supported :  1 destroyed
	// Bitfield not supported :  1 in_emission
	// Bitfield not supported :  1 marks_unbound
	// set_next : record
	// hash_next : record
	// signals : record
}

func BindingEntryNewFromC(u unsafe.Pointer) *BindingEntry {
	if u == nil {
		return nil
	}

	g := &BindingEntry{native: u}

	return g
}

func (recv *BindingEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BindingSet is a wrapper around the C record GtkBindingSet.
type BindingSet struct {
	native   unsafe.Pointer
	SetName  string
	Priority int32
	// widget_path_pspecs : record
	// widget_class_pspecs : record
	// class_branch_pspecs : record
	// entries : record
	// current : record
	// Bitfield not supported :  1 parsed
}

func BindingSetNewFromC(u unsafe.Pointer) *BindingSet {
	if u == nil {
		return nil
	}

	g := &BindingSet{native: u}

	return g
}

func (recv *BindingSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BindingSignal is a wrapper around the C record GtkBindingSignal.
type BindingSignal struct {
	native unsafe.Pointer
	// next : record
	SignalName string
	NArgs      uint32
	// no type for args
}

func BindingSignalNewFromC(u unsafe.Pointer) *BindingSignal {
	if u == nil {
		return nil
	}

	g := &BindingSignal{native: u}

	return g
}

func (recv *BindingSignal) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BooleanCellAccessibleClass is a wrapper around the C record GtkBooleanCellAccessibleClass.
type BooleanCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func BooleanCellAccessibleClassNewFromC(u unsafe.Pointer) *BooleanCellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &BooleanCellAccessibleClass{native: u}

	return g
}

func (recv *BooleanCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BooleanCellAccessiblePrivate is a wrapper around the C record GtkBooleanCellAccessiblePrivate.
type BooleanCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func BooleanCellAccessiblePrivateNewFromC(u unsafe.Pointer) *BooleanCellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &BooleanCellAccessiblePrivate{native: u}

	return g
}

func (recv *BooleanCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Border is a wrapper around the C record GtkBorder.
type Border struct {
	native unsafe.Pointer
	Left   int16
	Right  int16
	Top    int16
	Bottom int16
}

func BorderNewFromC(u unsafe.Pointer) *Border {
	if u == nil {
		return nil
	}

	g := &Border{native: u}

	return g
}

func (recv *Border) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BoxClass is a wrapper around the C record GtkBoxClass.
type BoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func BoxClassNewFromC(u unsafe.Pointer) *BoxClass {
	if u == nil {
		return nil
	}

	g := &BoxClass{native: u}

	return g
}

func (recv *BoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BoxPrivate is a wrapper around the C record GtkBoxPrivate.
type BoxPrivate struct {
	native unsafe.Pointer
}

func BoxPrivateNewFromC(u unsafe.Pointer) *BoxPrivate {
	if u == nil {
		return nil
	}

	g := &BoxPrivate{native: u}

	return g
}

func (recv *BoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BuildableIface is a wrapper around the C record GtkBuildableIface.
type BuildableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for set_name
	// no type for get_name
	// no type for add_child
	// no type for set_buildable_property
	// no type for construct_child
	// no type for custom_tag_start
	// no type for custom_tag_end
	// no type for custom_finished
	// no type for parser_finished
	// no type for get_internal_child
}

func BuildableIfaceNewFromC(u unsafe.Pointer) *BuildableIface {
	if u == nil {
		return nil
	}

	g := &BuildableIface{native: u}

	return g
}

func (recv *BuildableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BuilderClass is a wrapper around the C record GtkBuilderClass.
type BuilderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_type_from_name
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func BuilderClassNewFromC(u unsafe.Pointer) *BuilderClass {
	if u == nil {
		return nil
	}

	g := &BuilderClass{native: u}

	return g
}

func (recv *BuilderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BuilderPrivate is a wrapper around the C record GtkBuilderPrivate.
type BuilderPrivate struct {
	native unsafe.Pointer
}

func BuilderPrivateNewFromC(u unsafe.Pointer) *BuilderPrivate {
	if u == nil {
		return nil
	}

	g := &BuilderPrivate{native: u}

	return g
}

func (recv *BuilderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessibleClass is a wrapper around the C record GtkButtonAccessibleClass.
type ButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ButtonAccessibleClassNewFromC(u unsafe.Pointer) *ButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ButtonAccessibleClass{native: u}

	return g
}

func (recv *ButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessiblePrivate is a wrapper around the C record GtkButtonAccessiblePrivate.
type ButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func ButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ButtonAccessiblePrivate{native: u}

	return g
}

func (recv *ButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonBoxClass is a wrapper around the C record GtkButtonBoxClass.
type ButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ButtonBoxClassNewFromC(u unsafe.Pointer) *ButtonBoxClass {
	if u == nil {
		return nil
	}

	g := &ButtonBoxClass{native: u}

	return g
}

func (recv *ButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonBoxPrivate is a wrapper around the C record GtkButtonBoxPrivate.
type ButtonBoxPrivate struct {
	native unsafe.Pointer
}

func ButtonBoxPrivateNewFromC(u unsafe.Pointer) *ButtonBoxPrivate {
	if u == nil {
		return nil
	}

	g := &ButtonBoxPrivate{native: u}

	return g
}

func (recv *ButtonBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonClass is a wrapper around the C record GtkButtonClass.
type ButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for pressed
	// no type for released
	// no type for clicked
	// no type for enter
	// no type for leave
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ButtonClassNewFromC(u unsafe.Pointer) *ButtonClass {
	if u == nil {
		return nil
	}

	g := &ButtonClass{native: u}

	return g
}

func (recv *ButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonPrivate is a wrapper around the C record GtkButtonPrivate.
type ButtonPrivate struct {
	native unsafe.Pointer
}

func ButtonPrivateNewFromC(u unsafe.Pointer) *ButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ButtonPrivate{native: u}

	return g
}

func (recv *ButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CalendarClass is a wrapper around the C record GtkCalendarClass.
type CalendarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for month_changed
	// no type for day_selected
	// no type for day_selected_double_click
	// no type for prev_month
	// no type for next_month
	// no type for prev_year
	// no type for next_year
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CalendarClassNewFromC(u unsafe.Pointer) *CalendarClass {
	if u == nil {
		return nil
	}

	g := &CalendarClass{native: u}

	return g
}

func (recv *CalendarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CalendarPrivate is a wrapper around the C record GtkCalendarPrivate.
type CalendarPrivate struct {
	native unsafe.Pointer
}

func CalendarPrivateNewFromC(u unsafe.Pointer) *CalendarPrivate {
	if u == nil {
		return nil
	}

	g := &CalendarPrivate{native: u}

	return g
}

func (recv *CalendarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessibleClass is a wrapper around the C record GtkCellAccessibleClass.
type CellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for update_cache
}

func CellAccessibleClassNewFromC(u unsafe.Pointer) *CellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &CellAccessibleClass{native: u}

	return g
}

func (recv *CellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessibleParentIface is a wrapper around the C record GtkCellAccessibleParentIface.
type CellAccessibleParentIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_cell_extents
	// no type for get_cell_area
	// no type for grab_focus
	// no type for get_child_index
	// no type for get_renderer_state
	// no type for expand_collapse
	// no type for activate
	// no type for edit
	// no type for update_relationset
}

func CellAccessibleParentIfaceNewFromC(u unsafe.Pointer) *CellAccessibleParentIface {
	if u == nil {
		return nil
	}

	g := &CellAccessibleParentIface{native: u}

	return g
}

func (recv *CellAccessibleParentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessiblePrivate is a wrapper around the C record GtkCellAccessiblePrivate.
type CellAccessiblePrivate struct {
	native unsafe.Pointer
}

func CellAccessiblePrivateNewFromC(u unsafe.Pointer) *CellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &CellAccessiblePrivate{native: u}

	return g
}

func (recv *CellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaBoxClass is a wrapper around the C record GtkCellAreaBoxClass.
type CellAreaBoxClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellAreaBoxClassNewFromC(u unsafe.Pointer) *CellAreaBoxClass {
	if u == nil {
		return nil
	}

	g := &CellAreaBoxClass{native: u}

	return g
}

func (recv *CellAreaBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaBoxPrivate is a wrapper around the C record GtkCellAreaBoxPrivate.
type CellAreaBoxPrivate struct {
	native unsafe.Pointer
}

func CellAreaBoxPrivateNewFromC(u unsafe.Pointer) *CellAreaBoxPrivate {
	if u == nil {
		return nil
	}

	g := &CellAreaBoxPrivate{native: u}

	return g
}

func (recv *CellAreaBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaClass is a wrapper around the C record GtkCellAreaClass.
type CellAreaClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for add
	// no type for remove
	// no type for foreach
	// no type for foreach_alloc
	// no type for event
	// no type for render
	// no type for apply_attributes
	// no type for create_context
	// no type for copy_context
	// no type for get_request_mode
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for set_cell_property
	// no type for get_cell_property
	// no type for focus
	// no type for is_activatable
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func CellAreaClassNewFromC(u unsafe.Pointer) *CellAreaClass {
	if u == nil {
		return nil
	}

	g := &CellAreaClass{native: u}

	return g
}

func (recv *CellAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaContextClass is a wrapper around the C record GtkCellAreaContextClass.
type CellAreaContextClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for allocate
	// no type for reset
	// no type for get_preferred_height_for_width
	// no type for get_preferred_width_for_height
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

func CellAreaContextClassNewFromC(u unsafe.Pointer) *CellAreaContextClass {
	if u == nil {
		return nil
	}

	g := &CellAreaContextClass{native: u}

	return g
}

func (recv *CellAreaContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaContextPrivate is a wrapper around the C record GtkCellAreaContextPrivate.
type CellAreaContextPrivate struct {
	native unsafe.Pointer
}

func CellAreaContextPrivateNewFromC(u unsafe.Pointer) *CellAreaContextPrivate {
	if u == nil {
		return nil
	}

	g := &CellAreaContextPrivate{native: u}

	return g
}

func (recv *CellAreaContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaPrivate is a wrapper around the C record GtkCellAreaPrivate.
type CellAreaPrivate struct {
	native unsafe.Pointer
}

func CellAreaPrivateNewFromC(u unsafe.Pointer) *CellAreaPrivate {
	if u == nil {
		return nil
	}

	g := &CellAreaPrivate{native: u}

	return g
}

func (recv *CellAreaPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellEditableIface is a wrapper around the C record GtkCellEditableIface.
type CellEditableIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for editing_done
	// no type for remove_widget
	// no type for start_editing
}

func CellEditableIfaceNewFromC(u unsafe.Pointer) *CellEditableIface {
	if u == nil {
		return nil
	}

	g := &CellEditableIface{native: u}

	return g
}

func (recv *CellEditableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellLayoutIface is a wrapper around the C record GtkCellLayoutIface.
type CellLayoutIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for pack_start
	// no type for pack_end
	// no type for clear
	// no type for add_attribute
	// no type for set_cell_data_func
	// no type for clear_attributes
	// no type for reorder
	// no type for get_cells
	// no type for get_area
}

func CellLayoutIfaceNewFromC(u unsafe.Pointer) *CellLayoutIface {
	if u == nil {
		return nil
	}

	g := &CellLayoutIface{native: u}

	return g
}

func (recv *CellLayoutIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererAccelClass is a wrapper around the C record GtkCellRendererAccelClass.
type CellRendererAccelClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for accel_edited
	// no type for accel_cleared
	// no type for _gtk_reserved0
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererAccelClassNewFromC(u unsafe.Pointer) *CellRendererAccelClass {
	if u == nil {
		return nil
	}

	g := &CellRendererAccelClass{native: u}

	return g
}

func (recv *CellRendererAccelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererAccelPrivate is a wrapper around the C record GtkCellRendererAccelPrivate.
type CellRendererAccelPrivate struct {
	native unsafe.Pointer
}

func CellRendererAccelPrivateNewFromC(u unsafe.Pointer) *CellRendererAccelPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererAccelPrivate{native: u}

	return g
}

func (recv *CellRendererAccelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererClass is a wrapper around the C record GtkCellRendererClass.
type CellRendererClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for get_request_mode
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for get_aligned_area
	// no type for get_size
	// no type for render
	// no type for activate
	// no type for start_editing
	// no type for editing_canceled
	// no type for editing_started
	// Private : priv
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererClassNewFromC(u unsafe.Pointer) *CellRendererClass {
	if u == nil {
		return nil
	}

	g := &CellRendererClass{native: u}

	return g
}

func (recv *CellRendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererClassPrivate is a wrapper around the C record GtkCellRendererClassPrivate.
type CellRendererClassPrivate struct {
	native unsafe.Pointer
}

func CellRendererClassPrivateNewFromC(u unsafe.Pointer) *CellRendererClassPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererClassPrivate{native: u}

	return g
}

func (recv *CellRendererClassPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererComboClass is a wrapper around the C record GtkCellRendererComboClass.
type CellRendererComboClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererComboClassNewFromC(u unsafe.Pointer) *CellRendererComboClass {
	if u == nil {
		return nil
	}

	g := &CellRendererComboClass{native: u}

	return g
}

func (recv *CellRendererComboClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererComboPrivate is a wrapper around the C record GtkCellRendererComboPrivate.
type CellRendererComboPrivate struct {
	native unsafe.Pointer
}

func CellRendererComboPrivateNewFromC(u unsafe.Pointer) *CellRendererComboPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererComboPrivate{native: u}

	return g
}

func (recv *CellRendererComboPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbufClass is a wrapper around the C record GtkCellRendererPixbufClass.
type CellRendererPixbufClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererPixbufClassNewFromC(u unsafe.Pointer) *CellRendererPixbufClass {
	if u == nil {
		return nil
	}

	g := &CellRendererPixbufClass{native: u}

	return g
}

func (recv *CellRendererPixbufClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbufPrivate is a wrapper around the C record GtkCellRendererPixbufPrivate.
type CellRendererPixbufPrivate struct {
	native unsafe.Pointer
}

func CellRendererPixbufPrivateNewFromC(u unsafe.Pointer) *CellRendererPixbufPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererPixbufPrivate{native: u}

	return g
}

func (recv *CellRendererPixbufPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPrivate is a wrapper around the C record GtkCellRendererPrivate.
type CellRendererPrivate struct {
	native unsafe.Pointer
}

func CellRendererPrivateNewFromC(u unsafe.Pointer) *CellRendererPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererPrivate{native: u}

	return g
}

func (recv *CellRendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererProgressClass is a wrapper around the C record GtkCellRendererProgressClass.
type CellRendererProgressClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererProgressClassNewFromC(u unsafe.Pointer) *CellRendererProgressClass {
	if u == nil {
		return nil
	}

	g := &CellRendererProgressClass{native: u}

	return g
}

func (recv *CellRendererProgressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererProgressPrivate is a wrapper around the C record GtkCellRendererProgressPrivate.
type CellRendererProgressPrivate struct {
	native unsafe.Pointer
}

func CellRendererProgressPrivateNewFromC(u unsafe.Pointer) *CellRendererProgressPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererProgressPrivate{native: u}

	return g
}

func (recv *CellRendererProgressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinClass is a wrapper around the C record GtkCellRendererSpinClass.
type CellRendererSpinClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererSpinClassNewFromC(u unsafe.Pointer) *CellRendererSpinClass {
	if u == nil {
		return nil
	}

	g := &CellRendererSpinClass{native: u}

	return g
}

func (recv *CellRendererSpinClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinPrivate is a wrapper around the C record GtkCellRendererSpinPrivate.
type CellRendererSpinPrivate struct {
	native unsafe.Pointer
}

func CellRendererSpinPrivateNewFromC(u unsafe.Pointer) *CellRendererSpinPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererSpinPrivate{native: u}

	return g
}

func (recv *CellRendererSpinPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinnerClass is a wrapper around the C record GtkCellRendererSpinnerClass.
type CellRendererSpinnerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererSpinnerClassNewFromC(u unsafe.Pointer) *CellRendererSpinnerClass {
	if u == nil {
		return nil
	}

	g := &CellRendererSpinnerClass{native: u}

	return g
}

func (recv *CellRendererSpinnerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinnerPrivate is a wrapper around the C record GtkCellRendererSpinnerPrivate.
type CellRendererSpinnerPrivate struct {
	native unsafe.Pointer
}

func CellRendererSpinnerPrivateNewFromC(u unsafe.Pointer) *CellRendererSpinnerPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererSpinnerPrivate{native: u}

	return g
}

func (recv *CellRendererSpinnerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererTextClass is a wrapper around the C record GtkCellRendererTextClass.
type CellRendererTextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for edited
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererTextClassNewFromC(u unsafe.Pointer) *CellRendererTextClass {
	if u == nil {
		return nil
	}

	g := &CellRendererTextClass{native: u}

	return g
}

func (recv *CellRendererTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererTextPrivate is a wrapper around the C record GtkCellRendererTextPrivate.
type CellRendererTextPrivate struct {
	native unsafe.Pointer
}

func CellRendererTextPrivateNewFromC(u unsafe.Pointer) *CellRendererTextPrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererTextPrivate{native: u}

	return g
}

func (recv *CellRendererTextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererToggleClass is a wrapper around the C record GtkCellRendererToggleClass.
type CellRendererToggleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererToggleClassNewFromC(u unsafe.Pointer) *CellRendererToggleClass {
	if u == nil {
		return nil
	}

	g := &CellRendererToggleClass{native: u}

	return g
}

func (recv *CellRendererToggleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererTogglePrivate is a wrapper around the C record GtkCellRendererTogglePrivate.
type CellRendererTogglePrivate struct {
	native unsafe.Pointer
}

func CellRendererTogglePrivateNewFromC(u unsafe.Pointer) *CellRendererTogglePrivate {
	if u == nil {
		return nil
	}

	g := &CellRendererTogglePrivate{native: u}

	return g
}

func (recv *CellRendererTogglePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellViewClass is a wrapper around the C record GtkCellViewClass.
type CellViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellViewClassNewFromC(u unsafe.Pointer) *CellViewClass {
	if u == nil {
		return nil
	}

	g := &CellViewClass{native: u}

	return g
}

func (recv *CellViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellViewPrivate is a wrapper around the C record GtkCellViewPrivate.
type CellViewPrivate struct {
	native unsafe.Pointer
}

func CellViewPrivateNewFromC(u unsafe.Pointer) *CellViewPrivate {
	if u == nil {
		return nil
	}

	g := &CellViewPrivate{native: u}

	return g
}

func (recv *CellViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckButtonClass is a wrapper around the C record GtkCheckButtonClass.
type CheckButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CheckButtonClassNewFromC(u unsafe.Pointer) *CheckButtonClass {
	if u == nil {
		return nil
	}

	g := &CheckButtonClass{native: u}

	return g
}

func (recv *CheckButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemAccessibleClass is a wrapper around the C record GtkCheckMenuItemAccessibleClass.
type CheckMenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func CheckMenuItemAccessibleClassNewFromC(u unsafe.Pointer) *CheckMenuItemAccessibleClass {
	if u == nil {
		return nil
	}

	g := &CheckMenuItemAccessibleClass{native: u}

	return g
}

func (recv *CheckMenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemAccessiblePrivate is a wrapper around the C record GtkCheckMenuItemAccessiblePrivate.
type CheckMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func CheckMenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *CheckMenuItemAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &CheckMenuItemAccessiblePrivate{native: u}

	return g
}

func (recv *CheckMenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemClass is a wrapper around the C record GtkCheckMenuItemClass.
type CheckMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CheckMenuItemClassNewFromC(u unsafe.Pointer) *CheckMenuItemClass {
	if u == nil {
		return nil
	}

	g := &CheckMenuItemClass{native: u}

	return g
}

func (recv *CheckMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemPrivate is a wrapper around the C record GtkCheckMenuItemPrivate.
type CheckMenuItemPrivate struct {
	native unsafe.Pointer
}

func CheckMenuItemPrivateNewFromC(u unsafe.Pointer) *CheckMenuItemPrivate {
	if u == nil {
		return nil
	}

	g := &CheckMenuItemPrivate{native: u}

	return g
}

func (recv *CheckMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorButtonClass is a wrapper around the C record GtkColorButtonClass.
type ColorButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for color_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorButtonClassNewFromC(u unsafe.Pointer) *ColorButtonClass {
	if u == nil {
		return nil
	}

	g := &ColorButtonClass{native: u}

	return g
}

func (recv *ColorButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorButtonPrivate is a wrapper around the C record GtkColorButtonPrivate.
type ColorButtonPrivate struct {
	native unsafe.Pointer
}

func ColorButtonPrivateNewFromC(u unsafe.Pointer) *ColorButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ColorButtonPrivate{native: u}

	return g
}

func (recv *ColorButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserDialogClass is a wrapper around the C record GtkColorChooserDialogClass.
type ColorChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorChooserDialogClassNewFromC(u unsafe.Pointer) *ColorChooserDialogClass {
	if u == nil {
		return nil
	}

	g := &ColorChooserDialogClass{native: u}

	return g
}

func (recv *ColorChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserDialogPrivate is a wrapper around the C record GtkColorChooserDialogPrivate.
type ColorChooserDialogPrivate struct {
	native unsafe.Pointer
}

func ColorChooserDialogPrivateNewFromC(u unsafe.Pointer) *ColorChooserDialogPrivate {
	if u == nil {
		return nil
	}

	g := &ColorChooserDialogPrivate{native: u}

	return g
}

func (recv *ColorChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserInterface is a wrapper around the C record GtkColorChooserInterface.
type ColorChooserInterface struct {
	native unsafe.Pointer
	// base_interface : record
	// no type for get_rgba
	// no type for set_rgba
	// no type for add_palette
	// no type for color_activated
	// no type for padding
}

func ColorChooserInterfaceNewFromC(u unsafe.Pointer) *ColorChooserInterface {
	if u == nil {
		return nil
	}

	g := &ColorChooserInterface{native: u}

	return g
}

func (recv *ColorChooserInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserWidgetClass is a wrapper around the C record GtkColorChooserWidgetClass.
type ColorChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func ColorChooserWidgetClassNewFromC(u unsafe.Pointer) *ColorChooserWidgetClass {
	if u == nil {
		return nil
	}

	g := &ColorChooserWidgetClass{native: u}

	return g
}

func (recv *ColorChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserWidgetPrivate is a wrapper around the C record GtkColorChooserWidgetPrivate.
type ColorChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func ColorChooserWidgetPrivateNewFromC(u unsafe.Pointer) *ColorChooserWidgetPrivate {
	if u == nil {
		return nil
	}

	g := &ColorChooserWidgetPrivate{native: u}

	return g
}

func (recv *ColorChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionClass is a wrapper around the C record GtkColorSelectionClass.
type ColorSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for color_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorSelectionClassNewFromC(u unsafe.Pointer) *ColorSelectionClass {
	if u == nil {
		return nil
	}

	g := &ColorSelectionClass{native: u}

	return g
}

func (recv *ColorSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionDialogClass is a wrapper around the C record GtkColorSelectionDialogClass.
type ColorSelectionDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorSelectionDialogClassNewFromC(u unsafe.Pointer) *ColorSelectionDialogClass {
	if u == nil {
		return nil
	}

	g := &ColorSelectionDialogClass{native: u}

	return g
}

func (recv *ColorSelectionDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionDialogPrivate is a wrapper around the C record GtkColorSelectionDialogPrivate.
type ColorSelectionDialogPrivate struct {
	native unsafe.Pointer
}

func ColorSelectionDialogPrivateNewFromC(u unsafe.Pointer) *ColorSelectionDialogPrivate {
	if u == nil {
		return nil
	}

	g := &ColorSelectionDialogPrivate{native: u}

	return g
}

func (recv *ColorSelectionDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionPrivate is a wrapper around the C record GtkColorSelectionPrivate.
type ColorSelectionPrivate struct {
	native unsafe.Pointer
}

func ColorSelectionPrivateNewFromC(u unsafe.Pointer) *ColorSelectionPrivate {
	if u == nil {
		return nil
	}

	g := &ColorSelectionPrivate{native: u}

	return g
}

func (recv *ColorSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxAccessibleClass is a wrapper around the C record GtkComboBoxAccessibleClass.
type ComboBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ComboBoxAccessibleClassNewFromC(u unsafe.Pointer) *ComboBoxAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ComboBoxAccessibleClass{native: u}

	return g
}

func (recv *ComboBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxAccessiblePrivate is a wrapper around the C record GtkComboBoxAccessiblePrivate.
type ComboBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func ComboBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *ComboBoxAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ComboBoxAccessiblePrivate{native: u}

	return g
}

func (recv *ComboBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxClass is a wrapper around the C record GtkComboBoxClass.
type ComboBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for format_entry_text
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

func ComboBoxClassNewFromC(u unsafe.Pointer) *ComboBoxClass {
	if u == nil {
		return nil
	}

	g := &ComboBoxClass{native: u}

	return g
}

func (recv *ComboBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxPrivate is a wrapper around the C record GtkComboBoxPrivate.
type ComboBoxPrivate struct {
	native unsafe.Pointer
}

func ComboBoxPrivateNewFromC(u unsafe.Pointer) *ComboBoxPrivate {
	if u == nil {
		return nil
	}

	g := &ComboBoxPrivate{native: u}

	return g
}

func (recv *ComboBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxTextClass is a wrapper around the C record GtkComboBoxTextClass.
type ComboBoxTextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ComboBoxTextClassNewFromC(u unsafe.Pointer) *ComboBoxTextClass {
	if u == nil {
		return nil
	}

	g := &ComboBoxTextClass{native: u}

	return g
}

func (recv *ComboBoxTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxTextPrivate is a wrapper around the C record GtkComboBoxTextPrivate.
type ComboBoxTextPrivate struct {
	native unsafe.Pointer
}

func ComboBoxTextPrivateNewFromC(u unsafe.Pointer) *ComboBoxTextPrivate {
	if u == nil {
		return nil
	}

	g := &ComboBoxTextPrivate{native: u}

	return g
}

func (recv *ComboBoxTextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessibleClass is a wrapper around the C record GtkContainerAccessibleClass.
type ContainerAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add_gtk
	// no type for remove_gtk
}

func ContainerAccessibleClassNewFromC(u unsafe.Pointer) *ContainerAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ContainerAccessibleClass{native: u}

	return g
}

func (recv *ContainerAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessiblePrivate is a wrapper around the C record GtkContainerAccessiblePrivate.
type ContainerAccessiblePrivate struct {
	native unsafe.Pointer
}

func ContainerAccessiblePrivateNewFromC(u unsafe.Pointer) *ContainerAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ContainerAccessiblePrivate{native: u}

	return g
}

func (recv *ContainerAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerCellAccessibleClass is a wrapper around the C record GtkContainerCellAccessibleClass.
type ContainerCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ContainerCellAccessibleClassNewFromC(u unsafe.Pointer) *ContainerCellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ContainerCellAccessibleClass{native: u}

	return g
}

func (recv *ContainerCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerCellAccessiblePrivate is a wrapper around the C record GtkContainerCellAccessiblePrivate.
type ContainerCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func ContainerCellAccessiblePrivateNewFromC(u unsafe.Pointer) *ContainerCellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ContainerCellAccessiblePrivate{native: u}

	return g
}

func (recv *ContainerCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerClass is a wrapper around the C record GtkContainerClass.
type ContainerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add
	// no type for remove
	// no type for check_resize
	// no type for forall
	// no type for set_focus_child
	// no type for child_type
	// no type for composite_name
	// no type for set_child_property
	// no type for get_child_property
	// no type for get_path_for_child
	// Private : _handle_border_width
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func ContainerClassNewFromC(u unsafe.Pointer) *ContainerClass {
	if u == nil {
		return nil
	}

	g := &ContainerClass{native: u}

	return g
}

func (recv *ContainerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerPrivate is a wrapper around the C record GtkContainerPrivate.
type ContainerPrivate struct {
	native unsafe.Pointer
}

func ContainerPrivateNewFromC(u unsafe.Pointer) *ContainerPrivate {
	if u == nil {
		return nil
	}

	g := &ContainerPrivate{native: u}

	return g
}

func (recv *ContainerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CssProviderClass is a wrapper around the C record GtkCssProviderClass.
type CssProviderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for parsing_error
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CssProviderClassNewFromC(u unsafe.Pointer) *CssProviderClass {
	if u == nil {
		return nil
	}

	g := &CssProviderClass{native: u}

	return g
}

func (recv *CssProviderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CssProviderPrivate is a wrapper around the C record GtkCssProviderPrivate.
type CssProviderPrivate struct {
	native unsafe.Pointer
}

func CssProviderPrivateNewFromC(u unsafe.Pointer) *CssProviderPrivate {
	if u == nil {
		return nil
	}

	g := &CssProviderPrivate{native: u}

	return g
}

func (recv *CssProviderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DialogClass is a wrapper around the C record GtkDialogClass.
type DialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func DialogClassNewFromC(u unsafe.Pointer) *DialogClass {
	if u == nil {
		return nil
	}

	g := &DialogClass{native: u}

	return g
}

func (recv *DialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DialogPrivate is a wrapper around the C record GtkDialogPrivate.
type DialogPrivate struct {
	native unsafe.Pointer
}

func DialogPrivateNewFromC(u unsafe.Pointer) *DialogPrivate {
	if u == nil {
		return nil
	}

	g := &DialogPrivate{native: u}

	return g
}

func (recv *DialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingAreaClass is a wrapper around the C record GtkDrawingAreaClass.
type DrawingAreaClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func DrawingAreaClassNewFromC(u unsafe.Pointer) *DrawingAreaClass {
	if u == nil {
		return nil
	}

	g := &DrawingAreaClass{native: u}

	return g
}

func (recv *DrawingAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EditableInterface is a wrapper around the C record GtkEditableInterface.
type EditableInterface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for insert_text
	// no type for delete_text
	// no type for changed
	// no type for do_insert_text
	// no type for do_delete_text
	// no type for get_chars
	// no type for set_selection_bounds
	// no type for get_selection_bounds
	// no type for set_position
	// no type for get_position
}

func EditableInterfaceNewFromC(u unsafe.Pointer) *EditableInterface {
	if u == nil {
		return nil
	}

	g := &EditableInterface{native: u}

	return g
}

func (recv *EditableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryAccessibleClass is a wrapper around the C record GtkEntryAccessibleClass.
type EntryAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func EntryAccessibleClassNewFromC(u unsafe.Pointer) *EntryAccessibleClass {
	if u == nil {
		return nil
	}

	g := &EntryAccessibleClass{native: u}

	return g
}

func (recv *EntryAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryAccessiblePrivate is a wrapper around the C record GtkEntryAccessiblePrivate.
type EntryAccessiblePrivate struct {
	native unsafe.Pointer
}

func EntryAccessiblePrivateNewFromC(u unsafe.Pointer) *EntryAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &EntryAccessiblePrivate{native: u}

	return g
}

func (recv *EntryAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryBufferClass is a wrapper around the C record GtkEntryBufferClass.
type EntryBufferClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for inserted_text
	// no type for deleted_text
	// no type for get_text
	// no type for get_length
	// no type for insert_text
	// no type for delete_text
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func EntryBufferClassNewFromC(u unsafe.Pointer) *EntryBufferClass {
	if u == nil {
		return nil
	}

	g := &EntryBufferClass{native: u}

	return g
}

func (recv *EntryBufferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryBufferPrivate is a wrapper around the C record GtkEntryBufferPrivate.
type EntryBufferPrivate struct {
	native unsafe.Pointer
}

func EntryBufferPrivateNewFromC(u unsafe.Pointer) *EntryBufferPrivate {
	if u == nil {
		return nil
	}

	g := &EntryBufferPrivate{native: u}

	return g
}

func (recv *EntryBufferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryClass is a wrapper around the C record GtkEntryClass.
type EntryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for populate_popup
	// no type for activate
	// no type for move_cursor
	// no type for insert_at_cursor
	// no type for delete_from_cursor
	// no type for backspace
	// no type for cut_clipboard
	// no type for copy_clipboard
	// no type for paste_clipboard
	// no type for toggle_overwrite
	// no type for get_text_area_size
	// no type for get_frame_size
	// no type for insert_emoji
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

func EntryClassNewFromC(u unsafe.Pointer) *EntryClass {
	if u == nil {
		return nil
	}

	g := &EntryClass{native: u}

	return g
}

func (recv *EntryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryCompletionClass is a wrapper around the C record GtkEntryCompletionClass.
type EntryCompletionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for match_selected
	// no type for action_activated
	// no type for insert_prefix
	// no type for cursor_on_match
	// no type for no_matches
	// no type for _gtk_reserved0
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

func EntryCompletionClassNewFromC(u unsafe.Pointer) *EntryCompletionClass {
	if u == nil {
		return nil
	}

	g := &EntryCompletionClass{native: u}

	return g
}

func (recv *EntryCompletionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryCompletionPrivate is a wrapper around the C record GtkEntryCompletionPrivate.
type EntryCompletionPrivate struct {
	native unsafe.Pointer
}

func EntryCompletionPrivateNewFromC(u unsafe.Pointer) *EntryCompletionPrivate {
	if u == nil {
		return nil
	}

	g := &EntryCompletionPrivate{native: u}

	return g
}

func (recv *EntryCompletionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryPrivate is a wrapper around the C record GtkEntryPrivate.
type EntryPrivate struct {
	native unsafe.Pointer
}

func EntryPrivateNewFromC(u unsafe.Pointer) *EntryPrivate {
	if u == nil {
		return nil
	}

	g := &EntryPrivate{native: u}

	return g
}

func (recv *EntryPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventBoxClass is a wrapper around the C record GtkEventBoxClass.
type EventBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func EventBoxClassNewFromC(u unsafe.Pointer) *EventBoxClass {
	if u == nil {
		return nil
	}

	g := &EventBoxClass{native: u}

	return g
}

func (recv *EventBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventBoxPrivate is a wrapper around the C record GtkEventBoxPrivate.
type EventBoxPrivate struct {
	native unsafe.Pointer
}

func EventBoxPrivateNewFromC(u unsafe.Pointer) *EventBoxPrivate {
	if u == nil {
		return nil
	}

	g := &EventBoxPrivate{native: u}

	return g
}

func (recv *EventBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventControllerClass is a wrapper around the C record GtkEventControllerClass.
type EventControllerClass struct {
	native unsafe.Pointer
}

func EventControllerClassNewFromC(u unsafe.Pointer) *EventControllerClass {
	if u == nil {
		return nil
	}

	g := &EventControllerClass{native: u}

	return g
}

func (recv *EventControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ExpanderAccessibleClass is a wrapper around the C record GtkExpanderAccessibleClass.
type ExpanderAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ExpanderAccessibleClassNewFromC(u unsafe.Pointer) *ExpanderAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ExpanderAccessibleClass{native: u}

	return g
}

func (recv *ExpanderAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ExpanderAccessiblePrivate is a wrapper around the C record GtkExpanderAccessiblePrivate.
type ExpanderAccessiblePrivate struct {
	native unsafe.Pointer
}

func ExpanderAccessiblePrivateNewFromC(u unsafe.Pointer) *ExpanderAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ExpanderAccessiblePrivate{native: u}

	return g
}

func (recv *ExpanderAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ExpanderClass is a wrapper around the C record GtkExpanderClass.
type ExpanderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ExpanderClassNewFromC(u unsafe.Pointer) *ExpanderClass {
	if u == nil {
		return nil
	}

	g := &ExpanderClass{native: u}

	return g
}

func (recv *ExpanderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ExpanderPrivate is a wrapper around the C record GtkExpanderPrivate.
type ExpanderPrivate struct {
	native unsafe.Pointer
}

func ExpanderPrivateNewFromC(u unsafe.Pointer) *ExpanderPrivate {
	if u == nil {
		return nil
	}

	g := &ExpanderPrivate{native: u}

	return g
}

func (recv *ExpanderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserButtonClass is a wrapper around the C record GtkFileChooserButtonClass.
type FileChooserButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for file_set
	// Private : __gtk_reserved1
	// Private : __gtk_reserved2
	// Private : __gtk_reserved3
	// Private : __gtk_reserved4
}

func FileChooserButtonClassNewFromC(u unsafe.Pointer) *FileChooserButtonClass {
	if u == nil {
		return nil
	}

	g := &FileChooserButtonClass{native: u}

	return g
}

func (recv *FileChooserButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserButtonPrivate is a wrapper around the C record GtkFileChooserButtonPrivate.
type FileChooserButtonPrivate struct {
	native unsafe.Pointer
}

func FileChooserButtonPrivateNewFromC(u unsafe.Pointer) *FileChooserButtonPrivate {
	if u == nil {
		return nil
	}

	g := &FileChooserButtonPrivate{native: u}

	return g
}

func (recv *FileChooserButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserDialogClass is a wrapper around the C record GtkFileChooserDialogClass.
type FileChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FileChooserDialogClassNewFromC(u unsafe.Pointer) *FileChooserDialogClass {
	if u == nil {
		return nil
	}

	g := &FileChooserDialogClass{native: u}

	return g
}

func (recv *FileChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserDialogPrivate is a wrapper around the C record GtkFileChooserDialogPrivate.
type FileChooserDialogPrivate struct {
	native unsafe.Pointer
}

func FileChooserDialogPrivateNewFromC(u unsafe.Pointer) *FileChooserDialogPrivate {
	if u == nil {
		return nil
	}

	g := &FileChooserDialogPrivate{native: u}

	return g
}

func (recv *FileChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserWidgetClass is a wrapper around the C record GtkFileChooserWidgetClass.
type FileChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FileChooserWidgetClassNewFromC(u unsafe.Pointer) *FileChooserWidgetClass {
	if u == nil {
		return nil
	}

	g := &FileChooserWidgetClass{native: u}

	return g
}

func (recv *FileChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserWidgetPrivate is a wrapper around the C record GtkFileChooserWidgetPrivate.
type FileChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func FileChooserWidgetPrivateNewFromC(u unsafe.Pointer) *FileChooserWidgetPrivate {
	if u == nil {
		return nil
	}

	g := &FileChooserWidgetPrivate{native: u}

	return g
}

func (recv *FileChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileFilterInfo is a wrapper around the C record GtkFileFilterInfo.
type FileFilterInfo struct {
	native      unsafe.Pointer
	Contains    FileFilterFlags
	Filename    string
	Uri         string
	DisplayName string
	MimeType    string
}

func FileFilterInfoNewFromC(u unsafe.Pointer) *FileFilterInfo {
	if u == nil {
		return nil
	}

	g := &FileFilterInfo{native: u}

	return g
}

func (recv *FileFilterInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FixedChild is a wrapper around the C record GtkFixedChild.
type FixedChild struct {
	native unsafe.Pointer
	// widget : record
	X int32
	Y int32
}

func FixedChildNewFromC(u unsafe.Pointer) *FixedChild {
	if u == nil {
		return nil
	}

	g := &FixedChild{native: u}

	return g
}

func (recv *FixedChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FixedClass is a wrapper around the C record GtkFixedClass.
type FixedClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FixedClassNewFromC(u unsafe.Pointer) *FixedClass {
	if u == nil {
		return nil
	}

	g := &FixedClass{native: u}

	return g
}

func (recv *FixedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FixedPrivate is a wrapper around the C record GtkFixedPrivate.
type FixedPrivate struct {
	native unsafe.Pointer
}

func FixedPrivateNewFromC(u unsafe.Pointer) *FixedPrivate {
	if u == nil {
		return nil
	}

	g := &FixedPrivate{native: u}

	return g
}

func (recv *FixedPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxAccessibleClass is a wrapper around the C record GtkFlowBoxAccessibleClass.
type FlowBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func FlowBoxAccessibleClassNewFromC(u unsafe.Pointer) *FlowBoxAccessibleClass {
	if u == nil {
		return nil
	}

	g := &FlowBoxAccessibleClass{native: u}

	return g
}

func (recv *FlowBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxAccessiblePrivate is a wrapper around the C record GtkFlowBoxAccessiblePrivate.
type FlowBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func FlowBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *FlowBoxAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &FlowBoxAccessiblePrivate{native: u}

	return g
}

func (recv *FlowBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxChildAccessibleClass is a wrapper around the C record GtkFlowBoxChildAccessibleClass.
type FlowBoxChildAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func FlowBoxChildAccessibleClassNewFromC(u unsafe.Pointer) *FlowBoxChildAccessibleClass {
	if u == nil {
		return nil
	}

	g := &FlowBoxChildAccessibleClass{native: u}

	return g
}

func (recv *FlowBoxChildAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxChildClass is a wrapper around the C record GtkFlowBoxChildClass.
type FlowBoxChildClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

func FlowBoxChildClassNewFromC(u unsafe.Pointer) *FlowBoxChildClass {
	if u == nil {
		return nil
	}

	g := &FlowBoxChildClass{native: u}

	return g
}

func (recv *FlowBoxChildClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxClass is a wrapper around the C record GtkFlowBoxClass.
type FlowBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for child_activated
	// no type for selected_children_changed
	// no type for activate_cursor_child
	// no type for toggle_cursor_child
	// no type for move_cursor
	// no type for select_all
	// no type for unselect_all
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

func FlowBoxClassNewFromC(u unsafe.Pointer) *FlowBoxClass {
	if u == nil {
		return nil
	}

	g := &FlowBoxClass{native: u}

	return g
}

func (recv *FlowBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontButtonClass is a wrapper around the C record GtkFontButtonClass.
type FontButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for font_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontButtonClassNewFromC(u unsafe.Pointer) *FontButtonClass {
	if u == nil {
		return nil
	}

	g := &FontButtonClass{native: u}

	return g
}

func (recv *FontButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontButtonPrivate is a wrapper around the C record GtkFontButtonPrivate.
type FontButtonPrivate struct {
	native unsafe.Pointer
}

func FontButtonPrivateNewFromC(u unsafe.Pointer) *FontButtonPrivate {
	if u == nil {
		return nil
	}

	g := &FontButtonPrivate{native: u}

	return g
}

func (recv *FontButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserDialogClass is a wrapper around the C record GtkFontChooserDialogClass.
type FontChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontChooserDialogClassNewFromC(u unsafe.Pointer) *FontChooserDialogClass {
	if u == nil {
		return nil
	}

	g := &FontChooserDialogClass{native: u}

	return g
}

func (recv *FontChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserDialogPrivate is a wrapper around the C record GtkFontChooserDialogPrivate.
type FontChooserDialogPrivate struct {
	native unsafe.Pointer
}

func FontChooserDialogPrivateNewFromC(u unsafe.Pointer) *FontChooserDialogPrivate {
	if u == nil {
		return nil
	}

	g := &FontChooserDialogPrivate{native: u}

	return g
}

func (recv *FontChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserIface is a wrapper around the C record GtkFontChooserIface.
type FontChooserIface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for get_font_family
	// no type for get_font_face
	// no type for get_font_size
	// no type for set_filter_func
	// no type for font_activated
	// no type for set_font_map
	// no type for get_font_map
	// no type for padding
}

func FontChooserIfaceNewFromC(u unsafe.Pointer) *FontChooserIface {
	if u == nil {
		return nil
	}

	g := &FontChooserIface{native: u}

	return g
}

func (recv *FontChooserIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserWidgetClass is a wrapper around the C record GtkFontChooserWidgetClass.
type FontChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func FontChooserWidgetClassNewFromC(u unsafe.Pointer) *FontChooserWidgetClass {
	if u == nil {
		return nil
	}

	g := &FontChooserWidgetClass{native: u}

	return g
}

func (recv *FontChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserWidgetPrivate is a wrapper around the C record GtkFontChooserWidgetPrivate.
type FontChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func FontChooserWidgetPrivateNewFromC(u unsafe.Pointer) *FontChooserWidgetPrivate {
	if u == nil {
		return nil
	}

	g := &FontChooserWidgetPrivate{native: u}

	return g
}

func (recv *FontChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionClass is a wrapper around the C record GtkFontSelectionClass.
type FontSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontSelectionClassNewFromC(u unsafe.Pointer) *FontSelectionClass {
	if u == nil {
		return nil
	}

	g := &FontSelectionClass{native: u}

	return g
}

func (recv *FontSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionDialogClass is a wrapper around the C record GtkFontSelectionDialogClass.
type FontSelectionDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontSelectionDialogClassNewFromC(u unsafe.Pointer) *FontSelectionDialogClass {
	if u == nil {
		return nil
	}

	g := &FontSelectionDialogClass{native: u}

	return g
}

func (recv *FontSelectionDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionDialogPrivate is a wrapper around the C record GtkFontSelectionDialogPrivate.
type FontSelectionDialogPrivate struct {
	native unsafe.Pointer
}

func FontSelectionDialogPrivateNewFromC(u unsafe.Pointer) *FontSelectionDialogPrivate {
	if u == nil {
		return nil
	}

	g := &FontSelectionDialogPrivate{native: u}

	return g
}

func (recv *FontSelectionDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionPrivate is a wrapper around the C record GtkFontSelectionPrivate.
type FontSelectionPrivate struct {
	native unsafe.Pointer
}

func FontSelectionPrivateNewFromC(u unsafe.Pointer) *FontSelectionPrivate {
	if u == nil {
		return nil
	}

	g := &FontSelectionPrivate{native: u}

	return g
}

func (recv *FontSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameAccessibleClass is a wrapper around the C record GtkFrameAccessibleClass.
type FrameAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func FrameAccessibleClassNewFromC(u unsafe.Pointer) *FrameAccessibleClass {
	if u == nil {
		return nil
	}

	g := &FrameAccessibleClass{native: u}

	return g
}

func (recv *FrameAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameAccessiblePrivate is a wrapper around the C record GtkFrameAccessiblePrivate.
type FrameAccessiblePrivate struct {
	native unsafe.Pointer
}

func FrameAccessiblePrivateNewFromC(u unsafe.Pointer) *FrameAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &FrameAccessiblePrivate{native: u}

	return g
}

func (recv *FrameAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameClass is a wrapper around the C record GtkFrameClass.
type FrameClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for compute_child_allocation
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FrameClassNewFromC(u unsafe.Pointer) *FrameClass {
	if u == nil {
		return nil
	}

	g := &FrameClass{native: u}

	return g
}

func (recv *FrameClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FramePrivate is a wrapper around the C record GtkFramePrivate.
type FramePrivate struct {
	native unsafe.Pointer
}

func FramePrivateNewFromC(u unsafe.Pointer) *FramePrivate {
	if u == nil {
		return nil
	}

	g := &FramePrivate{native: u}

	return g
}

func (recv *FramePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureClass is a wrapper around the C record GtkGestureClass.
type GestureClass struct {
	native unsafe.Pointer
}

func GestureClassNewFromC(u unsafe.Pointer) *GestureClass {
	if u == nil {
		return nil
	}

	g := &GestureClass{native: u}

	return g
}

func (recv *GestureClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureDragClass is a wrapper around the C record GtkGestureDragClass.
type GestureDragClass struct {
	native unsafe.Pointer
}

func GestureDragClassNewFromC(u unsafe.Pointer) *GestureDragClass {
	if u == nil {
		return nil
	}

	g := &GestureDragClass{native: u}

	return g
}

func (recv *GestureDragClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureLongPressClass is a wrapper around the C record GtkGestureLongPressClass.
type GestureLongPressClass struct {
	native unsafe.Pointer
}

func GestureLongPressClassNewFromC(u unsafe.Pointer) *GestureLongPressClass {
	if u == nil {
		return nil
	}

	g := &GestureLongPressClass{native: u}

	return g
}

func (recv *GestureLongPressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureMultiPressClass is a wrapper around the C record GtkGestureMultiPressClass.
type GestureMultiPressClass struct {
	native unsafe.Pointer
}

func GestureMultiPressClassNewFromC(u unsafe.Pointer) *GestureMultiPressClass {
	if u == nil {
		return nil
	}

	g := &GestureMultiPressClass{native: u}

	return g
}

func (recv *GestureMultiPressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GesturePanClass is a wrapper around the C record GtkGesturePanClass.
type GesturePanClass struct {
	native unsafe.Pointer
}

func GesturePanClassNewFromC(u unsafe.Pointer) *GesturePanClass {
	if u == nil {
		return nil
	}

	g := &GesturePanClass{native: u}

	return g
}

func (recv *GesturePanClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureRotateClass is a wrapper around the C record GtkGestureRotateClass.
type GestureRotateClass struct {
	native unsafe.Pointer
}

func GestureRotateClassNewFromC(u unsafe.Pointer) *GestureRotateClass {
	if u == nil {
		return nil
	}

	g := &GestureRotateClass{native: u}

	return g
}

func (recv *GestureRotateClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingleClass is a wrapper around the C record GtkGestureSingleClass.
type GestureSingleClass struct {
	native unsafe.Pointer
}

func GestureSingleClassNewFromC(u unsafe.Pointer) *GestureSingleClass {
	if u == nil {
		return nil
	}

	g := &GestureSingleClass{native: u}

	return g
}

func (recv *GestureSingleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSwipeClass is a wrapper around the C record GtkGestureSwipeClass.
type GestureSwipeClass struct {
	native unsafe.Pointer
}

func GestureSwipeClassNewFromC(u unsafe.Pointer) *GestureSwipeClass {
	if u == nil {
		return nil
	}

	g := &GestureSwipeClass{native: u}

	return g
}

func (recv *GestureSwipeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureZoomClass is a wrapper around the C record GtkGestureZoomClass.
type GestureZoomClass struct {
	native unsafe.Pointer
}

func GestureZoomClassNewFromC(u unsafe.Pointer) *GestureZoomClass {
	if u == nil {
		return nil
	}

	g := &GestureZoomClass{native: u}

	return g
}

func (recv *GestureZoomClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gradient is a wrapper around the C record GtkGradient.
type Gradient struct {
	native unsafe.Pointer
}

func GradientNewFromC(u unsafe.Pointer) *Gradient {
	if u == nil {
		return nil
	}

	g := &Gradient{native: u}

	return g
}

func (recv *Gradient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GridClass is a wrapper around the C record GtkGridClass.
type GridClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func GridClassNewFromC(u unsafe.Pointer) *GridClass {
	if u == nil {
		return nil
	}

	g := &GridClass{native: u}

	return g
}

func (recv *GridClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GridPrivate is a wrapper around the C record GtkGridPrivate.
type GridPrivate struct {
	native unsafe.Pointer
}

func GridPrivateNewFromC(u unsafe.Pointer) *GridPrivate {
	if u == nil {
		return nil
	}

	g := &GridPrivate{native: u}

	return g
}

func (recv *GridPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HBoxClass is a wrapper around the C record GtkHBoxClass.
type HBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HBoxClassNewFromC(u unsafe.Pointer) *HBoxClass {
	if u == nil {
		return nil
	}

	g := &HBoxClass{native: u}

	return g
}

func (recv *HBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HButtonBoxClass is a wrapper around the C record GtkHButtonBoxClass.
type HButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HButtonBoxClassNewFromC(u unsafe.Pointer) *HButtonBoxClass {
	if u == nil {
		return nil
	}

	g := &HButtonBoxClass{native: u}

	return g
}

func (recv *HButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HPanedClass is a wrapper around the C record GtkHPanedClass.
type HPanedClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HPanedClassNewFromC(u unsafe.Pointer) *HPanedClass {
	if u == nil {
		return nil
	}

	g := &HPanedClass{native: u}

	return g
}

func (recv *HPanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HSVClass is a wrapper around the C record GtkHSVClass.
type HSVClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for move
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HSVClassNewFromC(u unsafe.Pointer) *HSVClass {
	if u == nil {
		return nil
	}

	g := &HSVClass{native: u}

	return g
}

func (recv *HSVClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HSVPrivate is a wrapper around the C record GtkHSVPrivate.
type HSVPrivate struct {
	native unsafe.Pointer
}

func HSVPrivateNewFromC(u unsafe.Pointer) *HSVPrivate {
	if u == nil {
		return nil
	}

	g := &HSVPrivate{native: u}

	return g
}

func (recv *HSVPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScaleClass is a wrapper around the C record GtkHScaleClass.
type HScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HScaleClassNewFromC(u unsafe.Pointer) *HScaleClass {
	if u == nil {
		return nil
	}

	g := &HScaleClass{native: u}

	return g
}

func (recv *HScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScrollbarClass is a wrapper around the C record GtkHScrollbarClass.
type HScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HScrollbarClassNewFromC(u unsafe.Pointer) *HScrollbarClass {
	if u == nil {
		return nil
	}

	g := &HScrollbarClass{native: u}

	return g
}

func (recv *HScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HSeparatorClass is a wrapper around the C record GtkHSeparatorClass.
type HSeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func HSeparatorClassNewFromC(u unsafe.Pointer) *HSeparatorClass {
	if u == nil {
		return nil
	}

	g := &HSeparatorClass{native: u}

	return g
}

func (recv *HSeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HandleBoxClass is a wrapper around the C record GtkHandleBoxClass.
type HandleBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for child_attached
	// no type for child_detached
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HandleBoxClassNewFromC(u unsafe.Pointer) *HandleBoxClass {
	if u == nil {
		return nil
	}

	g := &HandleBoxClass{native: u}

	return g
}

func (recv *HandleBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HandleBoxPrivate is a wrapper around the C record GtkHandleBoxPrivate.
type HandleBoxPrivate struct {
	native unsafe.Pointer
}

func HandleBoxPrivateNewFromC(u unsafe.Pointer) *HandleBoxPrivate {
	if u == nil {
		return nil
	}

	g := &HandleBoxPrivate{native: u}

	return g
}

func (recv *HandleBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HeaderBarClass is a wrapper around the C record GtkHeaderBarClass.
type HeaderBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HeaderBarClassNewFromC(u unsafe.Pointer) *HeaderBarClass {
	if u == nil {
		return nil
	}

	g := &HeaderBarClass{native: u}

	return g
}

func (recv *HeaderBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HeaderBarPrivate is a wrapper around the C record GtkHeaderBarPrivate.
type HeaderBarPrivate struct {
	native unsafe.Pointer
}

func HeaderBarPrivateNewFromC(u unsafe.Pointer) *HeaderBarPrivate {
	if u == nil {
		return nil
	}

	g := &HeaderBarPrivate{native: u}

	return g
}

func (recv *HeaderBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextClass is a wrapper around the C record GtkIMContextClass.
type IMContextClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for preedit_start
	// no type for preedit_end
	// no type for preedit_changed
	// no type for commit
	// no type for retrieve_surrounding
	// no type for delete_surrounding
	// no type for set_client_window
	// no type for get_preedit_string
	// no type for filter_keypress
	// no type for focus_in
	// no type for focus_out
	// no type for reset
	// no type for set_cursor_location
	// no type for set_use_preedit
	// no type for set_surrounding
	// no type for get_surrounding
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

func IMContextClassNewFromC(u unsafe.Pointer) *IMContextClass {
	if u == nil {
		return nil
	}

	g := &IMContextClass{native: u}

	return g
}

func (recv *IMContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextInfo is a wrapper around the C record GtkIMContextInfo.
type IMContextInfo struct {
	native         unsafe.Pointer
	ContextId      string
	ContextName    string
	Domain         string
	DomainDirname  string
	DefaultLocales string
}

func IMContextInfoNewFromC(u unsafe.Pointer) *IMContextInfo {
	if u == nil {
		return nil
	}

	g := &IMContextInfo{native: u}

	return g
}

func (recv *IMContextInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimpleClass is a wrapper around the C record GtkIMContextSimpleClass.
type IMContextSimpleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func IMContextSimpleClassNewFromC(u unsafe.Pointer) *IMContextSimpleClass {
	if u == nil {
		return nil
	}

	g := &IMContextSimpleClass{native: u}

	return g
}

func (recv *IMContextSimpleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimplePrivate is a wrapper around the C record GtkIMContextSimplePrivate.
type IMContextSimplePrivate struct {
	native unsafe.Pointer
}

func IMContextSimplePrivateNewFromC(u unsafe.Pointer) *IMContextSimplePrivate {
	if u == nil {
		return nil
	}

	g := &IMContextSimplePrivate{native: u}

	return g
}

func (recv *IMContextSimplePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMMulticontextClass is a wrapper around the C record GtkIMMulticontextClass.
type IMMulticontextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IMMulticontextClassNewFromC(u unsafe.Pointer) *IMMulticontextClass {
	if u == nil {
		return nil
	}

	g := &IMMulticontextClass{native: u}

	return g
}

func (recv *IMMulticontextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMMulticontextPrivate is a wrapper around the C record GtkIMMulticontextPrivate.
type IMMulticontextPrivate struct {
	native unsafe.Pointer
}

func IMMulticontextPrivateNewFromC(u unsafe.Pointer) *IMMulticontextPrivate {
	if u == nil {
		return nil
	}

	g := &IMMulticontextPrivate{native: u}

	return g
}

func (recv *IMMulticontextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconFactoryClass is a wrapper around the C record GtkIconFactoryClass.
type IconFactoryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IconFactoryClassNewFromC(u unsafe.Pointer) *IconFactoryClass {
	if u == nil {
		return nil
	}

	g := &IconFactoryClass{native: u}

	return g
}

func (recv *IconFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconFactoryPrivate is a wrapper around the C record GtkIconFactoryPrivate.
type IconFactoryPrivate struct {
	native unsafe.Pointer
}

func IconFactoryPrivateNewFromC(u unsafe.Pointer) *IconFactoryPrivate {
	if u == nil {
		return nil
	}

	g := &IconFactoryPrivate{native: u}

	return g
}

func (recv *IconFactoryPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconInfoClass is a wrapper around the C record GtkIconInfoClass.
type IconInfoClass struct {
	native unsafe.Pointer
}

func IconInfoClassNewFromC(u unsafe.Pointer) *IconInfoClass {
	if u == nil {
		return nil
	}

	g := &IconInfoClass{native: u}

	return g
}

func (recv *IconInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconSet is a wrapper around the C record GtkIconSet.
type IconSet struct {
	native unsafe.Pointer
}

func IconSetNewFromC(u unsafe.Pointer) *IconSet {
	if u == nil {
		return nil
	}

	g := &IconSet{native: u}

	return g
}

func (recv *IconSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconSource is a wrapper around the C record GtkIconSource.
type IconSource struct {
	native unsafe.Pointer
}

func IconSourceNewFromC(u unsafe.Pointer) *IconSource {
	if u == nil {
		return nil
	}

	g := &IconSource{native: u}

	return g
}

func (recv *IconSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconThemeClass is a wrapper around the C record GtkIconThemeClass.
type IconThemeClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IconThemeClassNewFromC(u unsafe.Pointer) *IconThemeClass {
	if u == nil {
		return nil
	}

	g := &IconThemeClass{native: u}

	return g
}

func (recv *IconThemeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconThemePrivate is a wrapper around the C record GtkIconThemePrivate.
type IconThemePrivate struct {
	native unsafe.Pointer
}

func IconThemePrivateNewFromC(u unsafe.Pointer) *IconThemePrivate {
	if u == nil {
		return nil
	}

	g := &IconThemePrivate{native: u}

	return g
}

func (recv *IconThemePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconViewAccessibleClass is a wrapper around the C record GtkIconViewAccessibleClass.
type IconViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func IconViewAccessibleClassNewFromC(u unsafe.Pointer) *IconViewAccessibleClass {
	if u == nil {
		return nil
	}

	g := &IconViewAccessibleClass{native: u}

	return g
}

func (recv *IconViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconViewAccessiblePrivate is a wrapper around the C record GtkIconViewAccessiblePrivate.
type IconViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func IconViewAccessiblePrivateNewFromC(u unsafe.Pointer) *IconViewAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &IconViewAccessiblePrivate{native: u}

	return g
}

func (recv *IconViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconViewClass is a wrapper around the C record GtkIconViewClass.
type IconViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for item_activated
	// no type for selection_changed
	// no type for select_all
	// no type for unselect_all
	// no type for select_cursor_item
	// no type for toggle_cursor_item
	// no type for move_cursor
	// no type for activate_cursor_item
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IconViewClassNewFromC(u unsafe.Pointer) *IconViewClass {
	if u == nil {
		return nil
	}

	g := &IconViewClass{native: u}

	return g
}

func (recv *IconViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconViewPrivate is a wrapper around the C record GtkIconViewPrivate.
type IconViewPrivate struct {
	native unsafe.Pointer
}

func IconViewPrivateNewFromC(u unsafe.Pointer) *IconViewPrivate {
	if u == nil {
		return nil
	}

	g := &IconViewPrivate{native: u}

	return g
}

func (recv *IconViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageAccessibleClass is a wrapper around the C record GtkImageAccessibleClass.
type ImageAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ImageAccessibleClassNewFromC(u unsafe.Pointer) *ImageAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ImageAccessibleClass{native: u}

	return g
}

func (recv *ImageAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageAccessiblePrivate is a wrapper around the C record GtkImageAccessiblePrivate.
type ImageAccessiblePrivate struct {
	native unsafe.Pointer
}

func ImageAccessiblePrivateNewFromC(u unsafe.Pointer) *ImageAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ImageAccessiblePrivate{native: u}

	return g
}

func (recv *ImageAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageCellAccessibleClass is a wrapper around the C record GtkImageCellAccessibleClass.
type ImageCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ImageCellAccessibleClassNewFromC(u unsafe.Pointer) *ImageCellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ImageCellAccessibleClass{native: u}

	return g
}

func (recv *ImageCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageCellAccessiblePrivate is a wrapper around the C record GtkImageCellAccessiblePrivate.
type ImageCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func ImageCellAccessiblePrivateNewFromC(u unsafe.Pointer) *ImageCellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ImageCellAccessiblePrivate{native: u}

	return g
}

func (recv *ImageCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageClass is a wrapper around the C record GtkImageClass.
type ImageClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ImageClassNewFromC(u unsafe.Pointer) *ImageClass {
	if u == nil {
		return nil
	}

	g := &ImageClass{native: u}

	return g
}

func (recv *ImageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageMenuItemClass is a wrapper around the C record GtkImageMenuItemClass.
type ImageMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ImageMenuItemClassNewFromC(u unsafe.Pointer) *ImageMenuItemClass {
	if u == nil {
		return nil
	}

	g := &ImageMenuItemClass{native: u}

	return g
}

func (recv *ImageMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageMenuItemPrivate is a wrapper around the C record GtkImageMenuItemPrivate.
type ImageMenuItemPrivate struct {
	native unsafe.Pointer
}

func ImageMenuItemPrivateNewFromC(u unsafe.Pointer) *ImageMenuItemPrivate {
	if u == nil {
		return nil
	}

	g := &ImageMenuItemPrivate{native: u}

	return g
}

func (recv *ImageMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImagePrivate is a wrapper around the C record GtkImagePrivate.
type ImagePrivate struct {
	native unsafe.Pointer
}

func ImagePrivateNewFromC(u unsafe.Pointer) *ImagePrivate {
	if u == nil {
		return nil
	}

	g := &ImagePrivate{native: u}

	return g
}

func (recv *ImagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InfoBarClass is a wrapper around the C record GtkInfoBarClass.
type InfoBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func InfoBarClassNewFromC(u unsafe.Pointer) *InfoBarClass {
	if u == nil {
		return nil
	}

	g := &InfoBarClass{native: u}

	return g
}

func (recv *InfoBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InfoBarPrivate is a wrapper around the C record GtkInfoBarPrivate.
type InfoBarPrivate struct {
	native unsafe.Pointer
}

func InfoBarPrivateNewFromC(u unsafe.Pointer) *InfoBarPrivate {
	if u == nil {
		return nil
	}

	g := &InfoBarPrivate{native: u}

	return g
}

func (recv *InfoBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InvisibleClass is a wrapper around the C record GtkInvisibleClass.
type InvisibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func InvisibleClassNewFromC(u unsafe.Pointer) *InvisibleClass {
	if u == nil {
		return nil
	}

	g := &InvisibleClass{native: u}

	return g
}

func (recv *InvisibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InvisiblePrivate is a wrapper around the C record GtkInvisiblePrivate.
type InvisiblePrivate struct {
	native unsafe.Pointer
}

func InvisiblePrivateNewFromC(u unsafe.Pointer) *InvisiblePrivate {
	if u == nil {
		return nil
	}

	g := &InvisiblePrivate{native: u}

	return g
}

func (recv *InvisiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelAccessibleClass is a wrapper around the C record GtkLabelAccessibleClass.
type LabelAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func LabelAccessibleClassNewFromC(u unsafe.Pointer) *LabelAccessibleClass {
	if u == nil {
		return nil
	}

	g := &LabelAccessibleClass{native: u}

	return g
}

func (recv *LabelAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelAccessiblePrivate is a wrapper around the C record GtkLabelAccessiblePrivate.
type LabelAccessiblePrivate struct {
	native unsafe.Pointer
}

func LabelAccessiblePrivateNewFromC(u unsafe.Pointer) *LabelAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &LabelAccessiblePrivate{native: u}

	return g
}

func (recv *LabelAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelClass is a wrapper around the C record GtkLabelClass.
type LabelClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for move_cursor
	// no type for copy_clipboard
	// no type for populate_popup
	// no type for activate_link
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func LabelClassNewFromC(u unsafe.Pointer) *LabelClass {
	if u == nil {
		return nil
	}

	g := &LabelClass{native: u}

	return g
}

func (recv *LabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelPrivate is a wrapper around the C record GtkLabelPrivate.
type LabelPrivate struct {
	native unsafe.Pointer
}

func LabelPrivateNewFromC(u unsafe.Pointer) *LabelPrivate {
	if u == nil {
		return nil
	}

	g := &LabelPrivate{native: u}

	return g
}

func (recv *LabelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelSelectionInfo is a wrapper around the C record GtkLabelSelectionInfo.
type LabelSelectionInfo struct {
	native unsafe.Pointer
}

func LabelSelectionInfoNewFromC(u unsafe.Pointer) *LabelSelectionInfo {
	if u == nil {
		return nil
	}

	g := &LabelSelectionInfo{native: u}

	return g
}

func (recv *LabelSelectionInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutClass is a wrapper around the C record GtkLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	if u == nil {
		return nil
	}

	g := &LayoutClass{native: u}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutPrivate is a wrapper around the C record GtkLayoutPrivate.
type LayoutPrivate struct {
	native unsafe.Pointer
}

func LayoutPrivateNewFromC(u unsafe.Pointer) *LayoutPrivate {
	if u == nil {
		return nil
	}

	g := &LayoutPrivate{native: u}

	return g
}

func (recv *LayoutPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LevelBarAccessibleClass is a wrapper around the C record GtkLevelBarAccessibleClass.
type LevelBarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func LevelBarAccessibleClassNewFromC(u unsafe.Pointer) *LevelBarAccessibleClass {
	if u == nil {
		return nil
	}

	g := &LevelBarAccessibleClass{native: u}

	return g
}

func (recv *LevelBarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LevelBarAccessiblePrivate is a wrapper around the C record GtkLevelBarAccessiblePrivate.
type LevelBarAccessiblePrivate struct {
	native unsafe.Pointer
}

func LevelBarAccessiblePrivateNewFromC(u unsafe.Pointer) *LevelBarAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &LevelBarAccessiblePrivate{native: u}

	return g
}

func (recv *LevelBarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LevelBarClass is a wrapper around the C record GtkLevelBarClass.
type LevelBarClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for offset_changed
	// Private : padding
}

func LevelBarClassNewFromC(u unsafe.Pointer) *LevelBarClass {
	if u == nil {
		return nil
	}

	g := &LevelBarClass{native: u}

	return g
}

func (recv *LevelBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LevelBarPrivate is a wrapper around the C record GtkLevelBarPrivate.
type LevelBarPrivate struct {
	native unsafe.Pointer
}

func LevelBarPrivateNewFromC(u unsafe.Pointer) *LevelBarPrivate {
	if u == nil {
		return nil
	}

	g := &LevelBarPrivate{native: u}

	return g
}

func (recv *LevelBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButtonAccessibleClass is a wrapper around the C record GtkLinkButtonAccessibleClass.
type LinkButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func LinkButtonAccessibleClassNewFromC(u unsafe.Pointer) *LinkButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &LinkButtonAccessibleClass{native: u}

	return g
}

func (recv *LinkButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButtonAccessiblePrivate is a wrapper around the C record GtkLinkButtonAccessiblePrivate.
type LinkButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func LinkButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *LinkButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &LinkButtonAccessiblePrivate{native: u}

	return g
}

func (recv *LinkButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButtonClass is a wrapper around the C record GtkLinkButtonClass.
type LinkButtonClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for activate_link
	// no type for _gtk_padding1
	// no type for _gtk_padding2
	// no type for _gtk_padding3
	// no type for _gtk_padding4
}

func LinkButtonClassNewFromC(u unsafe.Pointer) *LinkButtonClass {
	if u == nil {
		return nil
	}

	g := &LinkButtonClass{native: u}

	return g
}

func (recv *LinkButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButtonPrivate is a wrapper around the C record GtkLinkButtonPrivate.
type LinkButtonPrivate struct {
	native unsafe.Pointer
}

func LinkButtonPrivateNewFromC(u unsafe.Pointer) *LinkButtonPrivate {
	if u == nil {
		return nil
	}

	g := &LinkButtonPrivate{native: u}

	return g
}

func (recv *LinkButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxAccessibleClass is a wrapper around the C record GtkListBoxAccessibleClass.
type ListBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ListBoxAccessibleClassNewFromC(u unsafe.Pointer) *ListBoxAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ListBoxAccessibleClass{native: u}

	return g
}

func (recv *ListBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxAccessiblePrivate is a wrapper around the C record GtkListBoxAccessiblePrivate.
type ListBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

func ListBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *ListBoxAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ListBoxAccessiblePrivate{native: u}

	return g
}

func (recv *ListBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxClass is a wrapper around the C record GtkListBoxClass.
type ListBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for row_selected
	// no type for row_activated
	// no type for activate_cursor_row
	// no type for toggle_cursor_row
	// no type for move_cursor
	// no type for selected_rows_changed
	// no type for select_all
	// no type for unselect_all
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

func ListBoxClassNewFromC(u unsafe.Pointer) *ListBoxClass {
	if u == nil {
		return nil
	}

	g := &ListBoxClass{native: u}

	return g
}

func (recv *ListBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxRowAccessibleClass is a wrapper around the C record GtkListBoxRowAccessibleClass.
type ListBoxRowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ListBoxRowAccessibleClassNewFromC(u unsafe.Pointer) *ListBoxRowAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ListBoxRowAccessibleClass{native: u}

	return g
}

func (recv *ListBoxRowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxRowClass is a wrapper around the C record GtkListBoxRowClass.
type ListBoxRowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

func ListBoxRowClassNewFromC(u unsafe.Pointer) *ListBoxRowClass {
	if u == nil {
		return nil
	}

	g := &ListBoxRowClass{native: u}

	return g
}

func (recv *ListBoxRowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStoreClass is a wrapper around the C record GtkListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ListStoreClassNewFromC(u unsafe.Pointer) *ListStoreClass {
	if u == nil {
		return nil
	}

	g := &ListStoreClass{native: u}

	return g
}

func (recv *ListStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStorePrivate is a wrapper around the C record GtkListStorePrivate.
type ListStorePrivate struct {
	native unsafe.Pointer
}

func ListStorePrivateNewFromC(u unsafe.Pointer) *ListStorePrivate {
	if u == nil {
		return nil
	}

	g := &ListStorePrivate{native: u}

	return g
}

func (recv *ListStorePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LockButtonAccessibleClass is a wrapper around the C record GtkLockButtonAccessibleClass.
type LockButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func LockButtonAccessibleClassNewFromC(u unsafe.Pointer) *LockButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &LockButtonAccessibleClass{native: u}

	return g
}

func (recv *LockButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LockButtonAccessiblePrivate is a wrapper around the C record GtkLockButtonAccessiblePrivate.
type LockButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func LockButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *LockButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &LockButtonAccessiblePrivate{native: u}

	return g
}

func (recv *LockButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LockButtonClass is a wrapper around the C record GtkLockButtonClass.
type LockButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for reserved0
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
	// no type for reserved5
	// no type for reserved6
	// no type for reserved7
}

func LockButtonClassNewFromC(u unsafe.Pointer) *LockButtonClass {
	if u == nil {
		return nil
	}

	g := &LockButtonClass{native: u}

	return g
}

func (recv *LockButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LockButtonPrivate is a wrapper around the C record GtkLockButtonPrivate.
type LockButtonPrivate struct {
	native unsafe.Pointer
}

func LockButtonPrivateNewFromC(u unsafe.Pointer) *LockButtonPrivate {
	if u == nil {
		return nil
	}

	g := &LockButtonPrivate{native: u}

	return g
}

func (recv *LockButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuAccessibleClass is a wrapper around the C record GtkMenuAccessibleClass.
type MenuAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func MenuAccessibleClassNewFromC(u unsafe.Pointer) *MenuAccessibleClass {
	if u == nil {
		return nil
	}

	g := &MenuAccessibleClass{native: u}

	return g
}

func (recv *MenuAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuAccessiblePrivate is a wrapper around the C record GtkMenuAccessiblePrivate.
type MenuAccessiblePrivate struct {
	native unsafe.Pointer
}

func MenuAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &MenuAccessiblePrivate{native: u}

	return g
}

func (recv *MenuAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuBarClass is a wrapper around the C record GtkMenuBarClass.
type MenuBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuBarClassNewFromC(u unsafe.Pointer) *MenuBarClass {
	if u == nil {
		return nil
	}

	g := &MenuBarClass{native: u}

	return g
}

func (recv *MenuBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuBarPrivate is a wrapper around the C record GtkMenuBarPrivate.
type MenuBarPrivate struct {
	native unsafe.Pointer
}

func MenuBarPrivateNewFromC(u unsafe.Pointer) *MenuBarPrivate {
	if u == nil {
		return nil
	}

	g := &MenuBarPrivate{native: u}

	return g
}

func (recv *MenuBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuButtonAccessibleClass is a wrapper around the C record GtkMenuButtonAccessibleClass.
type MenuButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func MenuButtonAccessibleClassNewFromC(u unsafe.Pointer) *MenuButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &MenuButtonAccessibleClass{native: u}

	return g
}

func (recv *MenuButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuButtonAccessiblePrivate is a wrapper around the C record GtkMenuButtonAccessiblePrivate.
type MenuButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func MenuButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &MenuButtonAccessiblePrivate{native: u}

	return g
}

func (recv *MenuButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuButtonClass is a wrapper around the C record GtkMenuButtonClass.
type MenuButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuButtonClassNewFromC(u unsafe.Pointer) *MenuButtonClass {
	if u == nil {
		return nil
	}

	g := &MenuButtonClass{native: u}

	return g
}

func (recv *MenuButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuButtonPrivate is a wrapper around the C record GtkMenuButtonPrivate.
type MenuButtonPrivate struct {
	native unsafe.Pointer
}

func MenuButtonPrivateNewFromC(u unsafe.Pointer) *MenuButtonPrivate {
	if u == nil {
		return nil
	}

	g := &MenuButtonPrivate{native: u}

	return g
}

func (recv *MenuButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuClass is a wrapper around the C record GtkMenuClass.
type MenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuClassNewFromC(u unsafe.Pointer) *MenuClass {
	if u == nil {
		return nil
	}

	g := &MenuClass{native: u}

	return g
}

func (recv *MenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemAccessibleClass is a wrapper around the C record GtkMenuItemAccessibleClass.
type MenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func MenuItemAccessibleClassNewFromC(u unsafe.Pointer) *MenuItemAccessibleClass {
	if u == nil {
		return nil
	}

	g := &MenuItemAccessibleClass{native: u}

	return g
}

func (recv *MenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemAccessiblePrivate is a wrapper around the C record GtkMenuItemAccessiblePrivate.
type MenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func MenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuItemAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &MenuItemAccessiblePrivate{native: u}

	return g
}

func (recv *MenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemClass is a wrapper around the C record GtkMenuItemClass.
type MenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Bitfield not supported :  1 hide_on_activate
	// no type for activate
	// no type for activate_item
	// no type for toggle_size_request
	// no type for toggle_size_allocate
	// no type for set_label
	// no type for get_label
	// no type for _select
	// no type for deselect
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuItemClassNewFromC(u unsafe.Pointer) *MenuItemClass {
	if u == nil {
		return nil
	}

	g := &MenuItemClass{native: u}

	return g
}

func (recv *MenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemPrivate is a wrapper around the C record GtkMenuItemPrivate.
type MenuItemPrivate struct {
	native unsafe.Pointer
}

func MenuItemPrivateNewFromC(u unsafe.Pointer) *MenuItemPrivate {
	if u == nil {
		return nil
	}

	g := &MenuItemPrivate{native: u}

	return g
}

func (recv *MenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuPrivate is a wrapper around the C record GtkMenuPrivate.
type MenuPrivate struct {
	native unsafe.Pointer
}

func MenuPrivateNewFromC(u unsafe.Pointer) *MenuPrivate {
	if u == nil {
		return nil
	}

	g := &MenuPrivate{native: u}

	return g
}

func (recv *MenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellAccessibleClass is a wrapper around the C record GtkMenuShellAccessibleClass.
type MenuShellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func MenuShellAccessibleClassNewFromC(u unsafe.Pointer) *MenuShellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &MenuShellAccessibleClass{native: u}

	return g
}

func (recv *MenuShellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellAccessiblePrivate is a wrapper around the C record GtkMenuShellAccessiblePrivate.
type MenuShellAccessiblePrivate struct {
	native unsafe.Pointer
}

func MenuShellAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuShellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &MenuShellAccessiblePrivate{native: u}

	return g
}

func (recv *MenuShellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellClass is a wrapper around the C record GtkMenuShellClass.
type MenuShellClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Bitfield not supported :  1 submenu_placement
	// no type for deactivate
	// no type for selection_done
	// no type for move_current
	// no type for activate_current
	// no type for cancel
	// no type for select_item
	// no type for insert
	// no type for get_popup_delay
	// no type for move_selected
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuShellClassNewFromC(u unsafe.Pointer) *MenuShellClass {
	if u == nil {
		return nil
	}

	g := &MenuShellClass{native: u}

	return g
}

func (recv *MenuShellClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellPrivate is a wrapper around the C record GtkMenuShellPrivate.
type MenuShellPrivate struct {
	native unsafe.Pointer
}

func MenuShellPrivateNewFromC(u unsafe.Pointer) *MenuShellPrivate {
	if u == nil {
		return nil
	}

	g := &MenuShellPrivate{native: u}

	return g
}

func (recv *MenuShellPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuToolButtonClass is a wrapper around the C record GtkMenuToolButtonClass.
type MenuToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for show_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuToolButtonClassNewFromC(u unsafe.Pointer) *MenuToolButtonClass {
	if u == nil {
		return nil
	}

	g := &MenuToolButtonClass{native: u}

	return g
}

func (recv *MenuToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuToolButtonPrivate is a wrapper around the C record GtkMenuToolButtonPrivate.
type MenuToolButtonPrivate struct {
	native unsafe.Pointer
}

func MenuToolButtonPrivateNewFromC(u unsafe.Pointer) *MenuToolButtonPrivate {
	if u == nil {
		return nil
	}

	g := &MenuToolButtonPrivate{native: u}

	return g
}

func (recv *MenuToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MessageDialogClass is a wrapper around the C record GtkMessageDialogClass.
type MessageDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MessageDialogClassNewFromC(u unsafe.Pointer) *MessageDialogClass {
	if u == nil {
		return nil
	}

	g := &MessageDialogClass{native: u}

	return g
}

func (recv *MessageDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MessageDialogPrivate is a wrapper around the C record GtkMessageDialogPrivate.
type MessageDialogPrivate struct {
	native unsafe.Pointer
}

func MessageDialogPrivateNewFromC(u unsafe.Pointer) *MessageDialogPrivate {
	if u == nil {
		return nil
	}

	g := &MessageDialogPrivate{native: u}

	return g
}

func (recv *MessageDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MiscClass is a wrapper around the C record GtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MiscClassNewFromC(u unsafe.Pointer) *MiscClass {
	if u == nil {
		return nil
	}

	g := &MiscClass{native: u}

	return g
}

func (recv *MiscClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MiscPrivate is a wrapper around the C record GtkMiscPrivate.
type MiscPrivate struct {
	native unsafe.Pointer
}

func MiscPrivateNewFromC(u unsafe.Pointer) *MiscPrivate {
	if u == nil {
		return nil
	}

	g := &MiscPrivate{native: u}

	return g
}

func (recv *MiscPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperationClass is a wrapper around the C record GtkMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MountOperationClassNewFromC(u unsafe.Pointer) *MountOperationClass {
	if u == nil {
		return nil
	}

	g := &MountOperationClass{native: u}

	return g
}

func (recv *MountOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperationPrivate is a wrapper around the C record GtkMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

func MountOperationPrivateNewFromC(u unsafe.Pointer) *MountOperationPrivate {
	if u == nil {
		return nil
	}

	g := &MountOperationPrivate{native: u}

	return g
}

func (recv *MountOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookAccessibleClass is a wrapper around the C record GtkNotebookAccessibleClass.
type NotebookAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func NotebookAccessibleClassNewFromC(u unsafe.Pointer) *NotebookAccessibleClass {
	if u == nil {
		return nil
	}

	g := &NotebookAccessibleClass{native: u}

	return g
}

func (recv *NotebookAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookAccessiblePrivate is a wrapper around the C record GtkNotebookAccessiblePrivate.
type NotebookAccessiblePrivate struct {
	native unsafe.Pointer
}

func NotebookAccessiblePrivateNewFromC(u unsafe.Pointer) *NotebookAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &NotebookAccessiblePrivate{native: u}

	return g
}

func (recv *NotebookAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookClass is a wrapper around the C record GtkNotebookClass.
type NotebookClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for switch_page
	// no type for select_page
	// no type for focus_tab
	// no type for change_current_page
	// no type for move_focus_out
	// no type for reorder_tab
	// no type for insert_page
	// no type for create_window
	// no type for page_reordered
	// no type for page_removed
	// no type for page_added
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func NotebookClassNewFromC(u unsafe.Pointer) *NotebookClass {
	if u == nil {
		return nil
	}

	g := &NotebookClass{native: u}

	return g
}

func (recv *NotebookClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPageAccessibleClass is a wrapper around the C record GtkNotebookPageAccessibleClass.
type NotebookPageAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func NotebookPageAccessibleClassNewFromC(u unsafe.Pointer) *NotebookPageAccessibleClass {
	if u == nil {
		return nil
	}

	g := &NotebookPageAccessibleClass{native: u}

	return g
}

func (recv *NotebookPageAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPageAccessiblePrivate is a wrapper around the C record GtkNotebookPageAccessiblePrivate.
type NotebookPageAccessiblePrivate struct {
	native unsafe.Pointer
}

func NotebookPageAccessiblePrivateNewFromC(u unsafe.Pointer) *NotebookPageAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &NotebookPageAccessiblePrivate{native: u}

	return g
}

func (recv *NotebookPageAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPrivate is a wrapper around the C record GtkNotebookPrivate.
type NotebookPrivate struct {
	native unsafe.Pointer
}

func NotebookPrivateNewFromC(u unsafe.Pointer) *NotebookPrivate {
	if u == nil {
		return nil
	}

	g := &NotebookPrivate{native: u}

	return g
}

func (recv *NotebookPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NumerableIconClass is a wrapper around the C record GtkNumerableIconClass.
type NumerableIconClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for padding
}

func NumerableIconClassNewFromC(u unsafe.Pointer) *NumerableIconClass {
	if u == nil {
		return nil
	}

	g := &NumerableIconClass{native: u}

	return g
}

func (recv *NumerableIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NumerableIconPrivate is a wrapper around the C record GtkNumerableIconPrivate.
type NumerableIconPrivate struct {
	native unsafe.Pointer
}

func NumerableIconPrivateNewFromC(u unsafe.Pointer) *NumerableIconPrivate {
	if u == nil {
		return nil
	}

	g := &NumerableIconPrivate{native: u}

	return g
}

func (recv *NumerableIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OffscreenWindowClass is a wrapper around the C record GtkOffscreenWindowClass.
type OffscreenWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func OffscreenWindowClassNewFromC(u unsafe.Pointer) *OffscreenWindowClass {
	if u == nil {
		return nil
	}

	g := &OffscreenWindowClass{native: u}

	return g
}

func (recv *OffscreenWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OrientableIface is a wrapper around the C record GtkOrientableIface.
type OrientableIface struct {
	native unsafe.Pointer
	// base_iface : record
}

func OrientableIfaceNewFromC(u unsafe.Pointer) *OrientableIface {
	if u == nil {
		return nil
	}

	g := &OrientableIface{native: u}

	return g
}

func (recv *OrientableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OverlayClass is a wrapper around the C record GtkOverlayClass.
type OverlayClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_child_position
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func OverlayClassNewFromC(u unsafe.Pointer) *OverlayClass {
	if u == nil {
		return nil
	}

	g := &OverlayClass{native: u}

	return g
}

func (recv *OverlayClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OverlayPrivate is a wrapper around the C record GtkOverlayPrivate.
type OverlayPrivate struct {
	native unsafe.Pointer
}

func OverlayPrivateNewFromC(u unsafe.Pointer) *OverlayPrivate {
	if u == nil {
		return nil
	}

	g := &OverlayPrivate{native: u}

	return g
}

func (recv *OverlayPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PageRange is a wrapper around the C record GtkPageRange.
type PageRange struct {
	native unsafe.Pointer
	Start  int32
	End    int32
}

func PageRangeNewFromC(u unsafe.Pointer) *PageRange {
	if u == nil {
		return nil
	}

	g := &PageRange{native: u}

	return g
}

func (recv *PageRange) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PanedAccessibleClass is a wrapper around the C record GtkPanedAccessibleClass.
type PanedAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func PanedAccessibleClassNewFromC(u unsafe.Pointer) *PanedAccessibleClass {
	if u == nil {
		return nil
	}

	g := &PanedAccessibleClass{native: u}

	return g
}

func (recv *PanedAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PanedAccessiblePrivate is a wrapper around the C record GtkPanedAccessiblePrivate.
type PanedAccessiblePrivate struct {
	native unsafe.Pointer
}

func PanedAccessiblePrivateNewFromC(u unsafe.Pointer) *PanedAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &PanedAccessiblePrivate{native: u}

	return g
}

func (recv *PanedAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PanedClass is a wrapper around the C record GtkPanedClass.
type PanedClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for cycle_child_focus
	// no type for toggle_handle_focus
	// no type for move_handle
	// no type for cycle_handle_focus
	// no type for accept_position
	// no type for cancel_position
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func PanedClassNewFromC(u unsafe.Pointer) *PanedClass {
	if u == nil {
		return nil
	}

	g := &PanedClass{native: u}

	return g
}

func (recv *PanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PanedPrivate is a wrapper around the C record GtkPanedPrivate.
type PanedPrivate struct {
	native unsafe.Pointer
}

func PanedPrivateNewFromC(u unsafe.Pointer) *PanedPrivate {
	if u == nil {
		return nil
	}

	g := &PanedPrivate{native: u}

	return g
}

func (recv *PanedPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PaperSize is a wrapper around the C record GtkPaperSize.
type PaperSize struct {
	native unsafe.Pointer
}

func PaperSizeNewFromC(u unsafe.Pointer) *PaperSize {
	if u == nil {
		return nil
	}

	g := &PaperSize{native: u}

	return g
}

func (recv *PaperSize) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PlacesSidebarClass is a wrapper around the C record GtkPlacesSidebarClass.
type PlacesSidebarClass struct {
	native unsafe.Pointer
}

func PlacesSidebarClassNewFromC(u unsafe.Pointer) *PlacesSidebarClass {
	if u == nil {
		return nil
	}

	g := &PlacesSidebarClass{native: u}

	return g
}

func (recv *PlacesSidebarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PlugClass is a wrapper around the C record GtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for embedded
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func PlugClassNewFromC(u unsafe.Pointer) *PlugClass {
	if u == nil {
		return nil
	}

	g := &PlugClass{native: u}

	return g
}

func (recv *PlugClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PlugPrivate is a wrapper around the C record GtkPlugPrivate.
type PlugPrivate struct {
	native unsafe.Pointer
}

func PlugPrivateNewFromC(u unsafe.Pointer) *PlugPrivate {
	if u == nil {
		return nil
	}

	g := &PlugPrivate{native: u}

	return g
}

func (recv *PlugPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverAccessibleClass is a wrapper around the C record GtkPopoverAccessibleClass.
type PopoverAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func PopoverAccessibleClassNewFromC(u unsafe.Pointer) *PopoverAccessibleClass {
	if u == nil {
		return nil
	}

	g := &PopoverAccessibleClass{native: u}

	return g
}

func (recv *PopoverAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverClass is a wrapper around the C record GtkPopoverClass.
type PopoverClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for closed
	// Private : reserved
}

func PopoverClassNewFromC(u unsafe.Pointer) *PopoverClass {
	if u == nil {
		return nil
	}

	g := &PopoverClass{native: u}

	return g
}

func (recv *PopoverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverMenuClass is a wrapper around the C record GtkPopoverMenuClass.
type PopoverMenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : reserved
}

func PopoverMenuClassNewFromC(u unsafe.Pointer) *PopoverMenuClass {
	if u == nil {
		return nil
	}

	g := &PopoverMenuClass{native: u}

	return g
}

func (recv *PopoverMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverPrivate is a wrapper around the C record GtkPopoverPrivate.
type PopoverPrivate struct {
	native unsafe.Pointer
}

func PopoverPrivateNewFromC(u unsafe.Pointer) *PopoverPrivate {
	if u == nil {
		return nil
	}

	g := &PopoverPrivate{native: u}

	return g
}

func (recv *PopoverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperationClass is a wrapper around the C record GtkPrintOperationClass.
type PrintOperationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for done
	// no type for begin_print
	// no type for paginate
	// no type for request_page_setup
	// no type for draw_page
	// no type for end_print
	// no type for status_changed
	// no type for create_custom_widget
	// no type for custom_widget_apply
	// no type for preview
	// no type for update_custom_widget
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func PrintOperationClassNewFromC(u unsafe.Pointer) *PrintOperationClass {
	if u == nil {
		return nil
	}

	g := &PrintOperationClass{native: u}

	return g
}

func (recv *PrintOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperationPreviewIface is a wrapper around the C record GtkPrintOperationPreviewIface.
type PrintOperationPreviewIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for ready
	// no type for got_page_size
	// no type for render_page
	// no type for is_selected
	// no type for end_preview
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func PrintOperationPreviewIfaceNewFromC(u unsafe.Pointer) *PrintOperationPreviewIface {
	if u == nil {
		return nil
	}

	g := &PrintOperationPreviewIface{native: u}

	return g
}

func (recv *PrintOperationPreviewIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperationPrivate is a wrapper around the C record GtkPrintOperationPrivate.
type PrintOperationPrivate struct {
	native unsafe.Pointer
}

func PrintOperationPrivateNewFromC(u unsafe.Pointer) *PrintOperationPrivate {
	if u == nil {
		return nil
	}

	g := &PrintOperationPrivate{native: u}

	return g
}

func (recv *PrintOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarAccessibleClass is a wrapper around the C record GtkProgressBarAccessibleClass.
type ProgressBarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ProgressBarAccessibleClassNewFromC(u unsafe.Pointer) *ProgressBarAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ProgressBarAccessibleClass{native: u}

	return g
}

func (recv *ProgressBarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarAccessiblePrivate is a wrapper around the C record GtkProgressBarAccessiblePrivate.
type ProgressBarAccessiblePrivate struct {
	native unsafe.Pointer
}

func ProgressBarAccessiblePrivateNewFromC(u unsafe.Pointer) *ProgressBarAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ProgressBarAccessiblePrivate{native: u}

	return g
}

func (recv *ProgressBarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarClass is a wrapper around the C record GtkProgressBarClass.
type ProgressBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ProgressBarClassNewFromC(u unsafe.Pointer) *ProgressBarClass {
	if u == nil {
		return nil
	}

	g := &ProgressBarClass{native: u}

	return g
}

func (recv *ProgressBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarPrivate is a wrapper around the C record GtkProgressBarPrivate.
type ProgressBarPrivate struct {
	native unsafe.Pointer
}

func ProgressBarPrivateNewFromC(u unsafe.Pointer) *ProgressBarPrivate {
	if u == nil {
		return nil
	}

	g := &ProgressBarPrivate{native: u}

	return g
}

func (recv *ProgressBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioActionClass is a wrapper around the C record GtkRadioActionClass.
type RadioActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioActionClassNewFromC(u unsafe.Pointer) *RadioActionClass {
	if u == nil {
		return nil
	}

	g := &RadioActionClass{native: u}

	return g
}

func (recv *RadioActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioActionEntry is a wrapper around the C record GtkRadioActionEntry.
type RadioActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	Value       int32
}

func RadioActionEntryNewFromC(u unsafe.Pointer) *RadioActionEntry {
	if u == nil {
		return nil
	}

	g := &RadioActionEntry{native: u}

	return g
}

func (recv *RadioActionEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioActionPrivate is a wrapper around the C record GtkRadioActionPrivate.
type RadioActionPrivate struct {
	native unsafe.Pointer
}

func RadioActionPrivateNewFromC(u unsafe.Pointer) *RadioActionPrivate {
	if u == nil {
		return nil
	}

	g := &RadioActionPrivate{native: u}

	return g
}

func (recv *RadioActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonAccessibleClass is a wrapper around the C record GtkRadioButtonAccessibleClass.
type RadioButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func RadioButtonAccessibleClassNewFromC(u unsafe.Pointer) *RadioButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &RadioButtonAccessibleClass{native: u}

	return g
}

func (recv *RadioButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonAccessiblePrivate is a wrapper around the C record GtkRadioButtonAccessiblePrivate.
type RadioButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func RadioButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *RadioButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &RadioButtonAccessiblePrivate{native: u}

	return g
}

func (recv *RadioButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonClass is a wrapper around the C record GtkRadioButtonClass.
type RadioButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioButtonClassNewFromC(u unsafe.Pointer) *RadioButtonClass {
	if u == nil {
		return nil
	}

	g := &RadioButtonClass{native: u}

	return g
}

func (recv *RadioButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonPrivate is a wrapper around the C record GtkRadioButtonPrivate.
type RadioButtonPrivate struct {
	native unsafe.Pointer
}

func RadioButtonPrivateNewFromC(u unsafe.Pointer) *RadioButtonPrivate {
	if u == nil {
		return nil
	}

	g := &RadioButtonPrivate{native: u}

	return g
}

func (recv *RadioButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemAccessibleClass is a wrapper around the C record GtkRadioMenuItemAccessibleClass.
type RadioMenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func RadioMenuItemAccessibleClassNewFromC(u unsafe.Pointer) *RadioMenuItemAccessibleClass {
	if u == nil {
		return nil
	}

	g := &RadioMenuItemAccessibleClass{native: u}

	return g
}

func (recv *RadioMenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemAccessiblePrivate is a wrapper around the C record GtkRadioMenuItemAccessiblePrivate.
type RadioMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

func RadioMenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *RadioMenuItemAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &RadioMenuItemAccessiblePrivate{native: u}

	return g
}

func (recv *RadioMenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemClass is a wrapper around the C record GtkRadioMenuItemClass.
type RadioMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioMenuItemClassNewFromC(u unsafe.Pointer) *RadioMenuItemClass {
	if u == nil {
		return nil
	}

	g := &RadioMenuItemClass{native: u}

	return g
}

func (recv *RadioMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemPrivate is a wrapper around the C record GtkRadioMenuItemPrivate.
type RadioMenuItemPrivate struct {
	native unsafe.Pointer
}

func RadioMenuItemPrivateNewFromC(u unsafe.Pointer) *RadioMenuItemPrivate {
	if u == nil {
		return nil
	}

	g := &RadioMenuItemPrivate{native: u}

	return g
}

func (recv *RadioMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioToolButtonClass is a wrapper around the C record GtkRadioToolButtonClass.
type RadioToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioToolButtonClassNewFromC(u unsafe.Pointer) *RadioToolButtonClass {
	if u == nil {
		return nil
	}

	g := &RadioToolButtonClass{native: u}

	return g
}

func (recv *RadioToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangeAccessibleClass is a wrapper around the C record GtkRangeAccessibleClass.
type RangeAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func RangeAccessibleClassNewFromC(u unsafe.Pointer) *RangeAccessibleClass {
	if u == nil {
		return nil
	}

	g := &RangeAccessibleClass{native: u}

	return g
}

func (recv *RangeAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangeAccessiblePrivate is a wrapper around the C record GtkRangeAccessiblePrivate.
type RangeAccessiblePrivate struct {
	native unsafe.Pointer
}

func RangeAccessiblePrivateNewFromC(u unsafe.Pointer) *RangeAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &RangeAccessiblePrivate{native: u}

	return g
}

func (recv *RangeAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangeClass is a wrapper around the C record GtkRangeClass.
type RangeClass struct {
	native unsafe.Pointer
	// parent_class : record
	SliderDetail  string
	StepperDetail string
	// no type for value_changed
	// no type for adjust_bounds
	// no type for move_slider
	// no type for get_range_border
	// no type for change_value
	// no type for get_range_size_request
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

func RangeClassNewFromC(u unsafe.Pointer) *RangeClass {
	if u == nil {
		return nil
	}

	g := &RangeClass{native: u}

	return g
}

func (recv *RangeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangePrivate is a wrapper around the C record GtkRangePrivate.
type RangePrivate struct {
	native unsafe.Pointer
}

func RangePrivateNewFromC(u unsafe.Pointer) *RangePrivate {
	if u == nil {
		return nil
	}

	g := &RangePrivate{native: u}

	return g
}

func (recv *RangePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RcContext is a wrapper around the C record GtkRcContext.
type RcContext struct {
	native unsafe.Pointer
}

func RcContextNewFromC(u unsafe.Pointer) *RcContext {
	if u == nil {
		return nil
	}

	g := &RcContext{native: u}

	return g
}

func (recv *RcContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RcProperty is a wrapper around the C record GtkRcProperty.
type RcProperty struct {
	native       unsafe.Pointer
	TypeName     glib.Quark
	PropertyName glib.Quark
	Origin       string
	// value : record
}

func RcPropertyNewFromC(u unsafe.Pointer) *RcProperty {
	if u == nil {
		return nil
	}

	g := &RcProperty{native: u}

	return g
}

func (recv *RcProperty) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RcStyleClass is a wrapper around the C record GtkRcStyleClass.
type RcStyleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for create_rc_style
	// no type for parse
	// no type for merge
	// no type for create_style
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RcStyleClassNewFromC(u unsafe.Pointer) *RcStyleClass {
	if u == nil {
		return nil
	}

	g := &RcStyleClass{native: u}

	return g
}

func (recv *RcStyleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentActionClass is a wrapper around the C record GtkRecentActionClass.
type RecentActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentActionClassNewFromC(u unsafe.Pointer) *RecentActionClass {
	if u == nil {
		return nil
	}

	g := &RecentActionClass{native: u}

	return g
}

func (recv *RecentActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentActionPrivate is a wrapper around the C record GtkRecentActionPrivate.
type RecentActionPrivate struct {
	native unsafe.Pointer
}

func RecentActionPrivateNewFromC(u unsafe.Pointer) *RecentActionPrivate {
	if u == nil {
		return nil
	}

	g := &RecentActionPrivate{native: u}

	return g
}

func (recv *RecentActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserDialogClass is a wrapper around the C record GtkRecentChooserDialogClass.
type RecentChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentChooserDialogClassNewFromC(u unsafe.Pointer) *RecentChooserDialogClass {
	if u == nil {
		return nil
	}

	g := &RecentChooserDialogClass{native: u}

	return g
}

func (recv *RecentChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserDialogPrivate is a wrapper around the C record GtkRecentChooserDialogPrivate.
type RecentChooserDialogPrivate struct {
	native unsafe.Pointer
}

func RecentChooserDialogPrivateNewFromC(u unsafe.Pointer) *RecentChooserDialogPrivate {
	if u == nil {
		return nil
	}

	g := &RecentChooserDialogPrivate{native: u}

	return g
}

func (recv *RecentChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserIface is a wrapper around the C record GtkRecentChooserIface.
type RecentChooserIface struct {
	native unsafe.Pointer
	// Private : base_iface
	// no type for set_current_uri
	// no type for get_current_uri
	// no type for select_uri
	// no type for unselect_uri
	// no type for select_all
	// no type for unselect_all
	// no type for get_items
	// no type for get_recent_manager
	// no type for add_filter
	// no type for remove_filter
	// no type for list_filters
	// no type for set_sort_func
	// no type for item_activated
	// no type for selection_changed
}

func RecentChooserIfaceNewFromC(u unsafe.Pointer) *RecentChooserIface {
	if u == nil {
		return nil
	}

	g := &RecentChooserIface{native: u}

	return g
}

func (recv *RecentChooserIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserMenuClass is a wrapper around the C record GtkRecentChooserMenuClass.
type RecentChooserMenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for gtk_recent1
	// no type for gtk_recent2
	// no type for gtk_recent3
	// no type for gtk_recent4
}

func RecentChooserMenuClassNewFromC(u unsafe.Pointer) *RecentChooserMenuClass {
	if u == nil {
		return nil
	}

	g := &RecentChooserMenuClass{native: u}

	return g
}

func (recv *RecentChooserMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserMenuPrivate is a wrapper around the C record GtkRecentChooserMenuPrivate.
type RecentChooserMenuPrivate struct {
	native unsafe.Pointer
}

func RecentChooserMenuPrivateNewFromC(u unsafe.Pointer) *RecentChooserMenuPrivate {
	if u == nil {
		return nil
	}

	g := &RecentChooserMenuPrivate{native: u}

	return g
}

func (recv *RecentChooserMenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserWidgetClass is a wrapper around the C record GtkRecentChooserWidgetClass.
type RecentChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentChooserWidgetClassNewFromC(u unsafe.Pointer) *RecentChooserWidgetClass {
	if u == nil {
		return nil
	}

	g := &RecentChooserWidgetClass{native: u}

	return g
}

func (recv *RecentChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserWidgetPrivate is a wrapper around the C record GtkRecentChooserWidgetPrivate.
type RecentChooserWidgetPrivate struct {
	native unsafe.Pointer
}

func RecentChooserWidgetPrivateNewFromC(u unsafe.Pointer) *RecentChooserWidgetPrivate {
	if u == nil {
		return nil
	}

	g := &RecentChooserWidgetPrivate{native: u}

	return g
}

func (recv *RecentChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentData is a wrapper around the C record GtkRecentData.
type RecentData struct {
	native      unsafe.Pointer
	DisplayName string
	Description string
	MimeType    string
	AppName     string
	AppExec     string
	// no type for groups
	IsPrivate bool
}

func RecentDataNewFromC(u unsafe.Pointer) *RecentData {
	if u == nil {
		return nil
	}

	g := &RecentData{native: u}

	return g
}

func (recv *RecentData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentFilterInfo is a wrapper around the C record GtkRecentFilterInfo.
type RecentFilterInfo struct {
	native      unsafe.Pointer
	Contains    RecentFilterFlags
	Uri         string
	DisplayName string
	MimeType    string
	// no type for applications
	// no type for groups
	Age int32
}

func RecentFilterInfoNewFromC(u unsafe.Pointer) *RecentFilterInfo {
	if u == nil {
		return nil
	}

	g := &RecentFilterInfo{native: u}

	return g
}

func (recv *RecentFilterInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentManagerPrivate is a wrapper around the C record GtkRecentManagerPrivate.
type RecentManagerPrivate struct {
	native unsafe.Pointer
}

func RecentManagerPrivateNewFromC(u unsafe.Pointer) *RecentManagerPrivate {
	if u == nil {
		return nil
	}

	g := &RecentManagerPrivate{native: u}

	return g
}

func (recv *RecentManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessibleClass is a wrapper around the C record GtkRendererCellAccessibleClass.
type RendererCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func RendererCellAccessibleClassNewFromC(u unsafe.Pointer) *RendererCellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &RendererCellAccessibleClass{native: u}

	return g
}

func (recv *RendererCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessiblePrivate is a wrapper around the C record GtkRendererCellAccessiblePrivate.
type RendererCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func RendererCellAccessiblePrivateNewFromC(u unsafe.Pointer) *RendererCellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &RendererCellAccessiblePrivate{native: u}

	return g
}

func (recv *RendererCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RequestedSize is a wrapper around the C record GtkRequestedSize.
type RequestedSize struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	MinimumSize int32
	NaturalSize int32
}

func RequestedSizeNewFromC(u unsafe.Pointer) *RequestedSize {
	if u == nil {
		return nil
	}

	g := &RequestedSize{native: u}

	return g
}

func (recv *RequestedSize) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Requisition is a wrapper around the C record GtkRequisition.
type Requisition struct {
	native unsafe.Pointer
	Width  int32
	Height int32
}

func RequisitionNewFromC(u unsafe.Pointer) *Requisition {
	if u == nil {
		return nil
	}

	g := &Requisition{native: u}

	return g
}

func (recv *Requisition) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RevealerClass is a wrapper around the C record GtkRevealerClass.
type RevealerClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func RevealerClassNewFromC(u unsafe.Pointer) *RevealerClass {
	if u == nil {
		return nil
	}

	g := &RevealerClass{native: u}

	return g
}

func (recv *RevealerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleAccessibleClass is a wrapper around the C record GtkScaleAccessibleClass.
type ScaleAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ScaleAccessibleClassNewFromC(u unsafe.Pointer) *ScaleAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ScaleAccessibleClass{native: u}

	return g
}

func (recv *ScaleAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleAccessiblePrivate is a wrapper around the C record GtkScaleAccessiblePrivate.
type ScaleAccessiblePrivate struct {
	native unsafe.Pointer
}

func ScaleAccessiblePrivateNewFromC(u unsafe.Pointer) *ScaleAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ScaleAccessiblePrivate{native: u}

	return g
}

func (recv *ScaleAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButtonAccessibleClass is a wrapper around the C record GtkScaleButtonAccessibleClass.
type ScaleButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ScaleButtonAccessibleClassNewFromC(u unsafe.Pointer) *ScaleButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ScaleButtonAccessibleClass{native: u}

	return g
}

func (recv *ScaleButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButtonAccessiblePrivate is a wrapper around the C record GtkScaleButtonAccessiblePrivate.
type ScaleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func ScaleButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ScaleButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ScaleButtonAccessiblePrivate{native: u}

	return g
}

func (recv *ScaleButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButtonClass is a wrapper around the C record GtkScaleButtonClass.
type ScaleButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScaleButtonClassNewFromC(u unsafe.Pointer) *ScaleButtonClass {
	if u == nil {
		return nil
	}

	g := &ScaleButtonClass{native: u}

	return g
}

func (recv *ScaleButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButtonPrivate is a wrapper around the C record GtkScaleButtonPrivate.
type ScaleButtonPrivate struct {
	native unsafe.Pointer
}

func ScaleButtonPrivateNewFromC(u unsafe.Pointer) *ScaleButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ScaleButtonPrivate{native: u}

	return g
}

func (recv *ScaleButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleClass is a wrapper around the C record GtkScaleClass.
type ScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for format_value
	// no type for draw_value
	// no type for get_layout_offsets
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScaleClassNewFromC(u unsafe.Pointer) *ScaleClass {
	if u == nil {
		return nil
	}

	g := &ScaleClass{native: u}

	return g
}

func (recv *ScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScalePrivate is a wrapper around the C record GtkScalePrivate.
type ScalePrivate struct {
	native unsafe.Pointer
}

func ScalePrivateNewFromC(u unsafe.Pointer) *ScalePrivate {
	if u == nil {
		return nil
	}

	g := &ScalePrivate{native: u}

	return g
}

func (recv *ScalePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrollableInterface is a wrapper around the C record GtkScrollableInterface.
type ScrollableInterface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for get_border
}

func ScrollableInterfaceNewFromC(u unsafe.Pointer) *ScrollableInterface {
	if u == nil {
		return nil
	}

	g := &ScrollableInterface{native: u}

	return g
}

func (recv *ScrollableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrollbarClass is a wrapper around the C record GtkScrollbarClass.
type ScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScrollbarClassNewFromC(u unsafe.Pointer) *ScrollbarClass {
	if u == nil {
		return nil
	}

	g := &ScrollbarClass{native: u}

	return g
}

func (recv *ScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowAccessibleClass is a wrapper around the C record GtkScrolledWindowAccessibleClass.
type ScrolledWindowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ScrolledWindowAccessibleClassNewFromC(u unsafe.Pointer) *ScrolledWindowAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ScrolledWindowAccessibleClass{native: u}

	return g
}

func (recv *ScrolledWindowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowAccessiblePrivate is a wrapper around the C record GtkScrolledWindowAccessiblePrivate.
type ScrolledWindowAccessiblePrivate struct {
	native unsafe.Pointer
}

func ScrolledWindowAccessiblePrivateNewFromC(u unsafe.Pointer) *ScrolledWindowAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ScrolledWindowAccessiblePrivate{native: u}

	return g
}

func (recv *ScrolledWindowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowClass is a wrapper around the C record GtkScrolledWindowClass.
type ScrolledWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	ScrollbarSpacing int32
	// no type for scroll_child
	// no type for move_focus_out
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScrolledWindowClassNewFromC(u unsafe.Pointer) *ScrolledWindowClass {
	if u == nil {
		return nil
	}

	g := &ScrolledWindowClass{native: u}

	return g
}

func (recv *ScrolledWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowPrivate is a wrapper around the C record GtkScrolledWindowPrivate.
type ScrolledWindowPrivate struct {
	native unsafe.Pointer
}

func ScrolledWindowPrivateNewFromC(u unsafe.Pointer) *ScrolledWindowPrivate {
	if u == nil {
		return nil
	}

	g := &ScrolledWindowPrivate{native: u}

	return g
}

func (recv *ScrolledWindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SearchBarClass is a wrapper around the C record GtkSearchBarClass.
type SearchBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SearchBarClassNewFromC(u unsafe.Pointer) *SearchBarClass {
	if u == nil {
		return nil
	}

	g := &SearchBarClass{native: u}

	return g
}

func (recv *SearchBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SearchEntryClass is a wrapper around the C record GtkSearchEntryClass.
type SearchEntryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for search_changed
	// no type for next_match
	// no type for previous_match
	// no type for stop_search
}

func SearchEntryClassNewFromC(u unsafe.Pointer) *SearchEntryClass {
	if u == nil {
		return nil
	}

	g := &SearchEntryClass{native: u}

	return g
}

func (recv *SearchEntryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SelectionData is a wrapper around the C record GtkSelectionData.
type SelectionData struct {
	native unsafe.Pointer
}

func SelectionDataNewFromC(u unsafe.Pointer) *SelectionData {
	if u == nil {
		return nil
	}

	g := &SelectionData{native: u}

	return g
}

func (recv *SelectionData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorClass is a wrapper around the C record GtkSeparatorClass.
type SeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorClassNewFromC(u unsafe.Pointer) *SeparatorClass {
	if u == nil {
		return nil
	}

	g := &SeparatorClass{native: u}

	return g
}

func (recv *SeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorMenuItemClass is a wrapper around the C record GtkSeparatorMenuItemClass.
type SeparatorMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorMenuItemClassNewFromC(u unsafe.Pointer) *SeparatorMenuItemClass {
	if u == nil {
		return nil
	}

	g := &SeparatorMenuItemClass{native: u}

	return g
}

func (recv *SeparatorMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorPrivate is a wrapper around the C record GtkSeparatorPrivate.
type SeparatorPrivate struct {
	native unsafe.Pointer
}

func SeparatorPrivateNewFromC(u unsafe.Pointer) *SeparatorPrivate {
	if u == nil {
		return nil
	}

	g := &SeparatorPrivate{native: u}

	return g
}

func (recv *SeparatorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorToolItemClass is a wrapper around the C record GtkSeparatorToolItemClass.
type SeparatorToolItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorToolItemClassNewFromC(u unsafe.Pointer) *SeparatorToolItemClass {
	if u == nil {
		return nil
	}

	g := &SeparatorToolItemClass{native: u}

	return g
}

func (recv *SeparatorToolItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorToolItemPrivate is a wrapper around the C record GtkSeparatorToolItemPrivate.
type SeparatorToolItemPrivate struct {
	native unsafe.Pointer
}

func SeparatorToolItemPrivateNewFromC(u unsafe.Pointer) *SeparatorToolItemPrivate {
	if u == nil {
		return nil
	}

	g := &SeparatorToolItemPrivate{native: u}

	return g
}

func (recv *SeparatorToolItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsClass is a wrapper around the C record GtkSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	if u == nil {
		return nil
	}

	g := &SettingsClass{native: u}

	return g
}

func (recv *SettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsPrivate is a wrapper around the C record GtkSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	if u == nil {
		return nil
	}

	g := &SettingsPrivate{native: u}

	return g
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsValue is a wrapper around the C record GtkSettingsValue.
type SettingsValue struct {
	native unsafe.Pointer
	Origin string
	// value : record
}

func SettingsValueNewFromC(u unsafe.Pointer) *SettingsValue {
	if u == nil {
		return nil
	}

	g := &SettingsValue{native: u}

	return g
}

func (recv *SettingsValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SizeGroupClass is a wrapper around the C record GtkSizeGroupClass.
type SizeGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SizeGroupClassNewFromC(u unsafe.Pointer) *SizeGroupClass {
	if u == nil {
		return nil
	}

	g := &SizeGroupClass{native: u}

	return g
}

func (recv *SizeGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SizeGroupPrivate is a wrapper around the C record GtkSizeGroupPrivate.
type SizeGroupPrivate struct {
	native unsafe.Pointer
}

func SizeGroupPrivateNewFromC(u unsafe.Pointer) *SizeGroupPrivate {
	if u == nil {
		return nil
	}

	g := &SizeGroupPrivate{native: u}

	return g
}

func (recv *SizeGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClass is a wrapper around the C record GtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for plug_added
	// no type for plug_removed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	if u == nil {
		return nil
	}

	g := &SocketClass{native: u}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketPrivate is a wrapper around the C record GtkSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

func SocketPrivateNewFromC(u unsafe.Pointer) *SocketPrivate {
	if u == nil {
		return nil
	}

	g := &SocketPrivate{native: u}

	return g
}

func (recv *SocketPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonAccessibleClass is a wrapper around the C record GtkSpinButtonAccessibleClass.
type SpinButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func SpinButtonAccessibleClassNewFromC(u unsafe.Pointer) *SpinButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &SpinButtonAccessibleClass{native: u}

	return g
}

func (recv *SpinButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonAccessiblePrivate is a wrapper around the C record GtkSpinButtonAccessiblePrivate.
type SpinButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func SpinButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *SpinButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &SpinButtonAccessiblePrivate{native: u}

	return g
}

func (recv *SpinButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonClass is a wrapper around the C record GtkSpinButtonClass.
type SpinButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for input
	// no type for output
	// no type for value_changed
	// no type for change_value
	// no type for wrapped
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SpinButtonClassNewFromC(u unsafe.Pointer) *SpinButtonClass {
	if u == nil {
		return nil
	}

	g := &SpinButtonClass{native: u}

	return g
}

func (recv *SpinButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonPrivate is a wrapper around the C record GtkSpinButtonPrivate.
type SpinButtonPrivate struct {
	native unsafe.Pointer
}

func SpinButtonPrivateNewFromC(u unsafe.Pointer) *SpinButtonPrivate {
	if u == nil {
		return nil
	}

	g := &SpinButtonPrivate{native: u}

	return g
}

func (recv *SpinButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinnerAccessibleClass is a wrapper around the C record GtkSpinnerAccessibleClass.
type SpinnerAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func SpinnerAccessibleClassNewFromC(u unsafe.Pointer) *SpinnerAccessibleClass {
	if u == nil {
		return nil
	}

	g := &SpinnerAccessibleClass{native: u}

	return g
}

func (recv *SpinnerAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinnerAccessiblePrivate is a wrapper around the C record GtkSpinnerAccessiblePrivate.
type SpinnerAccessiblePrivate struct {
	native unsafe.Pointer
}

func SpinnerAccessiblePrivateNewFromC(u unsafe.Pointer) *SpinnerAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &SpinnerAccessiblePrivate{native: u}

	return g
}

func (recv *SpinnerAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinnerClass is a wrapper around the C record GtkSpinnerClass.
type SpinnerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SpinnerClassNewFromC(u unsafe.Pointer) *SpinnerClass {
	if u == nil {
		return nil
	}

	g := &SpinnerClass{native: u}

	return g
}

func (recv *SpinnerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinnerPrivate is a wrapper around the C record GtkSpinnerPrivate.
type SpinnerPrivate struct {
	native unsafe.Pointer
}

func SpinnerPrivateNewFromC(u unsafe.Pointer) *SpinnerPrivate {
	if u == nil {
		return nil
	}

	g := &SpinnerPrivate{native: u}

	return g
}

func (recv *SpinnerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GtkStackAccessibleClass

// StackClass is a wrapper around the C record GtkStackClass.
type StackClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func StackClassNewFromC(u unsafe.Pointer) *StackClass {
	if u == nil {
		return nil
	}

	g := &StackClass{native: u}

	return g
}

func (recv *StackClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StackSidebarClass is a wrapper around the C record GtkStackSidebarClass.
type StackSidebarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StackSidebarClassNewFromC(u unsafe.Pointer) *StackSidebarClass {
	if u == nil {
		return nil
	}

	g := &StackSidebarClass{native: u}

	return g
}

func (recv *StackSidebarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StackSidebarPrivate is a wrapper around the C record GtkStackSidebarPrivate.
type StackSidebarPrivate struct {
	native unsafe.Pointer
}

func StackSidebarPrivateNewFromC(u unsafe.Pointer) *StackSidebarPrivate {
	if u == nil {
		return nil
	}

	g := &StackSidebarPrivate{native: u}

	return g
}

func (recv *StackSidebarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StackSwitcherClass is a wrapper around the C record GtkStackSwitcherClass.
type StackSwitcherClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StackSwitcherClassNewFromC(u unsafe.Pointer) *StackSwitcherClass {
	if u == nil {
		return nil
	}

	g := &StackSwitcherClass{native: u}

	return g
}

func (recv *StackSwitcherClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusIconClass is a wrapper around the C record GtkStatusIconClass.
type StatusIconClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for popup_menu
	// no type for size_changed
	// no type for button_press_event
	// no type for button_release_event
	// no type for scroll_event
	// no type for query_tooltip
	// __gtk_reserved1 : no type generator for gpointer, void*
	// __gtk_reserved2 : no type generator for gpointer, void*
	// __gtk_reserved3 : no type generator for gpointer, void*
	// __gtk_reserved4 : no type generator for gpointer, void*
}

func StatusIconClassNewFromC(u unsafe.Pointer) *StatusIconClass {
	if u == nil {
		return nil
	}

	g := &StatusIconClass{native: u}

	return g
}

func (recv *StatusIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusIconPrivate is a wrapper around the C record GtkStatusIconPrivate.
type StatusIconPrivate struct {
	native unsafe.Pointer
}

func StatusIconPrivateNewFromC(u unsafe.Pointer) *StatusIconPrivate {
	if u == nil {
		return nil
	}

	g := &StatusIconPrivate{native: u}

	return g
}

func (recv *StatusIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarAccessibleClass is a wrapper around the C record GtkStatusbarAccessibleClass.
type StatusbarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func StatusbarAccessibleClassNewFromC(u unsafe.Pointer) *StatusbarAccessibleClass {
	if u == nil {
		return nil
	}

	g := &StatusbarAccessibleClass{native: u}

	return g
}

func (recv *StatusbarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarAccessiblePrivate is a wrapper around the C record GtkStatusbarAccessiblePrivate.
type StatusbarAccessiblePrivate struct {
	native unsafe.Pointer
}

func StatusbarAccessiblePrivateNewFromC(u unsafe.Pointer) *StatusbarAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &StatusbarAccessiblePrivate{native: u}

	return g
}

func (recv *StatusbarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarClass is a wrapper around the C record GtkStatusbarClass.
type StatusbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// reserved : no type generator for gpointer, gpointer
	// no type for text_pushed
	// no type for text_popped
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StatusbarClassNewFromC(u unsafe.Pointer) *StatusbarClass {
	if u == nil {
		return nil
	}

	g := &StatusbarClass{native: u}

	return g
}

func (recv *StatusbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarPrivate is a wrapper around the C record GtkStatusbarPrivate.
type StatusbarPrivate struct {
	native unsafe.Pointer
}

func StatusbarPrivateNewFromC(u unsafe.Pointer) *StatusbarPrivate {
	if u == nil {
		return nil
	}

	g := &StatusbarPrivate{native: u}

	return g
}

func (recv *StatusbarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StockItem is a wrapper around the C record GtkStockItem.
type StockItem struct {
	native            unsafe.Pointer
	StockId           string
	Label             string
	Modifier          gdk.ModifierType
	Keyval            uint32
	TranslationDomain string
}

func StockItemNewFromC(u unsafe.Pointer) *StockItem {
	if u == nil {
		return nil
	}

	g := &StockItem{native: u}

	return g
}

func (recv *StockItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleClass is a wrapper around the C record GtkStyleClass.
type StyleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for realize
	// no type for unrealize
	// no type for copy
	// no type for clone
	// no type for init_from_rc
	// no type for set_background
	// no type for render_icon
	// no type for draw_hline
	// no type for draw_vline
	// no type for draw_shadow
	// no type for draw_arrow
	// no type for draw_diamond
	// no type for draw_box
	// no type for draw_flat_box
	// no type for draw_check
	// no type for draw_option
	// no type for draw_tab
	// no type for draw_shadow_gap
	// no type for draw_box_gap
	// no type for draw_extension
	// no type for draw_focus
	// no type for draw_slider
	// no type for draw_handle
	// no type for draw_expander
	// no type for draw_layout
	// no type for draw_resize_grip
	// no type for draw_spinner
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
	// no type for _gtk_reserved9
	// no type for _gtk_reserved10
	// no type for _gtk_reserved11
}

func StyleClassNewFromC(u unsafe.Pointer) *StyleClass {
	if u == nil {
		return nil
	}

	g := &StyleClass{native: u}

	return g
}

func (recv *StyleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleContextClass is a wrapper around the C record GtkStyleContextClass.
type StyleContextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StyleContextClassNewFromC(u unsafe.Pointer) *StyleContextClass {
	if u == nil {
		return nil
	}

	g := &StyleContextClass{native: u}

	return g
}

func (recv *StyleContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleContextPrivate is a wrapper around the C record GtkStyleContextPrivate.
type StyleContextPrivate struct {
	native unsafe.Pointer
}

func StyleContextPrivateNewFromC(u unsafe.Pointer) *StyleContextPrivate {
	if u == nil {
		return nil
	}

	g := &StyleContextPrivate{native: u}

	return g
}

func (recv *StyleContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StylePropertiesClass is a wrapper around the C record GtkStylePropertiesClass.
type StylePropertiesClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StylePropertiesClassNewFromC(u unsafe.Pointer) *StylePropertiesClass {
	if u == nil {
		return nil
	}

	g := &StylePropertiesClass{native: u}

	return g
}

func (recv *StylePropertiesClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StylePropertiesPrivate is a wrapper around the C record GtkStylePropertiesPrivate.
type StylePropertiesPrivate struct {
	native unsafe.Pointer
}

func StylePropertiesPrivateNewFromC(u unsafe.Pointer) *StylePropertiesPrivate {
	if u == nil {
		return nil
	}

	g := &StylePropertiesPrivate{native: u}

	return g
}

func (recv *StylePropertiesPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleProviderIface is a wrapper around the C record GtkStyleProviderIface.
type StyleProviderIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_style
	// no type for get_style_property
	// no type for get_icon_factory
}

func StyleProviderIfaceNewFromC(u unsafe.Pointer) *StyleProviderIface {
	if u == nil {
		return nil
	}

	g := &StyleProviderIface{native: u}

	return g
}

func (recv *StyleProviderIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SwitchAccessibleClass is a wrapper around the C record GtkSwitchAccessibleClass.
type SwitchAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func SwitchAccessibleClassNewFromC(u unsafe.Pointer) *SwitchAccessibleClass {
	if u == nil {
		return nil
	}

	g := &SwitchAccessibleClass{native: u}

	return g
}

func (recv *SwitchAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SwitchAccessiblePrivate is a wrapper around the C record GtkSwitchAccessiblePrivate.
type SwitchAccessiblePrivate struct {
	native unsafe.Pointer
}

func SwitchAccessiblePrivateNewFromC(u unsafe.Pointer) *SwitchAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &SwitchAccessiblePrivate{native: u}

	return g
}

func (recv *SwitchAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SwitchClass is a wrapper around the C record GtkSwitchClass.
type SwitchClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for state_set
	// no type for _switch_padding_1
	// no type for _switch_padding_2
	// no type for _switch_padding_3
	// no type for _switch_padding_4
	// no type for _switch_padding_5
}

func SwitchClassNewFromC(u unsafe.Pointer) *SwitchClass {
	if u == nil {
		return nil
	}

	g := &SwitchClass{native: u}

	return g
}

func (recv *SwitchClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SwitchPrivate is a wrapper around the C record GtkSwitchPrivate.
type SwitchPrivate struct {
	native unsafe.Pointer
}

func SwitchPrivateNewFromC(u unsafe.Pointer) *SwitchPrivate {
	if u == nil {
		return nil
	}

	g := &SwitchPrivate{native: u}

	return g
}

func (recv *SwitchPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SymbolicColor is a wrapper around the C record GtkSymbolicColor.
type SymbolicColor struct {
	native unsafe.Pointer
}

func SymbolicColorNewFromC(u unsafe.Pointer) *SymbolicColor {
	if u == nil {
		return nil
	}

	g := &SymbolicColor{native: u}

	return g
}

func (recv *SymbolicColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableChild is a wrapper around the C record GtkTableChild.
type TableChild struct {
	native unsafe.Pointer
	// widget : record
	LeftAttach   uint16
	RightAttach  uint16
	TopAttach    uint16
	BottomAttach uint16
	Xpadding     uint16
	Ypadding     uint16
	// Bitfield not supported :  1 xexpand
	// Bitfield not supported :  1 yexpand
	// Bitfield not supported :  1 xshrink
	// Bitfield not supported :  1 yshrink
	// Bitfield not supported :  1 xfill
	// Bitfield not supported :  1 yfill
}

func TableChildNewFromC(u unsafe.Pointer) *TableChild {
	if u == nil {
		return nil
	}

	g := &TableChild{native: u}

	return g
}

func (recv *TableChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableClass is a wrapper around the C record GtkTableClass.
type TableClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TableClassNewFromC(u unsafe.Pointer) *TableClass {
	if u == nil {
		return nil
	}

	g := &TableClass{native: u}

	return g
}

func (recv *TableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TablePrivate is a wrapper around the C record GtkTablePrivate.
type TablePrivate struct {
	native unsafe.Pointer
}

func TablePrivateNewFromC(u unsafe.Pointer) *TablePrivate {
	if u == nil {
		return nil
	}

	g := &TablePrivate{native: u}

	return g
}

func (recv *TablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableRowCol is a wrapper around the C record GtkTableRowCol.
type TableRowCol struct {
	native      unsafe.Pointer
	Requisition uint16
	Allocation  uint16
	Spacing     uint16
	// Bitfield not supported :  1 need_expand
	// Bitfield not supported :  1 need_shrink
	// Bitfield not supported :  1 expand
	// Bitfield not supported :  1 shrink
	// Bitfield not supported :  1 empty
}

func TableRowColNewFromC(u unsafe.Pointer) *TableRowCol {
	if u == nil {
		return nil
	}

	g := &TableRowCol{native: u}

	return g
}

func (recv *TableRowCol) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TargetEntry is a wrapper around the C record GtkTargetEntry.
type TargetEntry struct {
	native unsafe.Pointer
	Target string
	Flags  uint32
	Info   uint32
}

func TargetEntryNewFromC(u unsafe.Pointer) *TargetEntry {
	if u == nil {
		return nil
	}

	g := &TargetEntry{native: u}

	return g
}

func (recv *TargetEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TargetList is a wrapper around the C record GtkTargetList.
type TargetList struct {
	native unsafe.Pointer
}

func TargetListNewFromC(u unsafe.Pointer) *TargetList {
	if u == nil {
		return nil
	}

	g := &TargetList{native: u}

	return g
}

func (recv *TargetList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TargetPair is a wrapper around the C record GtkTargetPair.
type TargetPair struct {
	native unsafe.Pointer
	// target : record
	Flags uint32
	Info  uint32
}

func TargetPairNewFromC(u unsafe.Pointer) *TargetPair {
	if u == nil {
		return nil
	}

	g := &TargetPair{native: u}

	return g
}

func (recv *TargetPair) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TearoffMenuItemClass is a wrapper around the C record GtkTearoffMenuItemClass.
type TearoffMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TearoffMenuItemClassNewFromC(u unsafe.Pointer) *TearoffMenuItemClass {
	if u == nil {
		return nil
	}

	g := &TearoffMenuItemClass{native: u}

	return g
}

func (recv *TearoffMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TearoffMenuItemPrivate is a wrapper around the C record GtkTearoffMenuItemPrivate.
type TearoffMenuItemPrivate struct {
	native unsafe.Pointer
}

func TearoffMenuItemPrivateNewFromC(u unsafe.Pointer) *TearoffMenuItemPrivate {
	if u == nil {
		return nil
	}

	g := &TearoffMenuItemPrivate{native: u}

	return g
}

func (recv *TearoffMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextAppearance is a wrapper around the C record GtkTextAppearance.
type TextAppearance struct {
	native unsafe.Pointer
	// bg_color : record
	// fg_color : record
	Rise int32
	// Bitfield not supported :  4 underline
	// Bitfield not supported :  1 strikethrough
	// Bitfield not supported :  1 draw_bg
	// Bitfield not supported :  1 inside_selection
	// Bitfield not supported :  1 is_text
}

func TextAppearanceNewFromC(u unsafe.Pointer) *TextAppearance {
	if u == nil {
		return nil
	}

	g := &TextAppearance{native: u}

	return g
}

func (recv *TextAppearance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextAttributes is a wrapper around the C record GtkTextAttributes.
type TextAttributes struct {
	native unsafe.Pointer
	// Private : refcount
	// appearance : record
	Justification Justification
	Direction     TextDirection
	// font : record
	FontScale        float64
	LeftMargin       int32
	RightMargin      int32
	Indent           int32
	PixelsAboveLines int32
	PixelsBelowLines int32
	PixelsInsideWrap int32
	// tabs : record
	WrapMode WrapMode
	// language : record
	// Private : pg_bg_color
	// Bitfield not supported :  1 invisible
	// Bitfield not supported :  1 bg_full_height
	// Bitfield not supported :  1 editable
	// Bitfield not supported :  1 no_fallback
	// Private : pg_bg_rgba
	LetterSpacing int32
}

func TextAttributesNewFromC(u unsafe.Pointer) *TextAttributes {
	if u == nil {
		return nil
	}

	g := &TextAttributes{native: u}

	return g
}

func (recv *TextAttributes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextBTree is a wrapper around the C record GtkTextBTree.
type TextBTree struct {
	native unsafe.Pointer
}

func TextBTreeNewFromC(u unsafe.Pointer) *TextBTree {
	if u == nil {
		return nil
	}

	g := &TextBTree{native: u}

	return g
}

func (recv *TextBTree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextBufferClass is a wrapper around the C record GtkTextBufferClass.
type TextBufferClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for insert_text
	// no type for insert_pixbuf
	// no type for insert_child_anchor
	// no type for delete_range
	// no type for changed
	// no type for modified_changed
	// no type for mark_set
	// no type for mark_deleted
	// no type for apply_tag
	// no type for remove_tag
	// no type for begin_user_action
	// no type for end_user_action
	// no type for paste_done
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextBufferClassNewFromC(u unsafe.Pointer) *TextBufferClass {
	if u == nil {
		return nil
	}

	g := &TextBufferClass{native: u}

	return g
}

func (recv *TextBufferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextBufferPrivate is a wrapper around the C record GtkTextBufferPrivate.
type TextBufferPrivate struct {
	native unsafe.Pointer
}

func TextBufferPrivateNewFromC(u unsafe.Pointer) *TextBufferPrivate {
	if u == nil {
		return nil
	}

	g := &TextBufferPrivate{native: u}

	return g
}

func (recv *TextBufferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextCellAccessibleClass is a wrapper around the C record GtkTextCellAccessibleClass.
type TextCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func TextCellAccessibleClassNewFromC(u unsafe.Pointer) *TextCellAccessibleClass {
	if u == nil {
		return nil
	}

	g := &TextCellAccessibleClass{native: u}

	return g
}

func (recv *TextCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextCellAccessiblePrivate is a wrapper around the C record GtkTextCellAccessiblePrivate.
type TextCellAccessiblePrivate struct {
	native unsafe.Pointer
}

func TextCellAccessiblePrivateNewFromC(u unsafe.Pointer) *TextCellAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &TextCellAccessiblePrivate{native: u}

	return g
}

func (recv *TextCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextChildAnchorClass is a wrapper around the C record GtkTextChildAnchorClass.
type TextChildAnchorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextChildAnchorClassNewFromC(u unsafe.Pointer) *TextChildAnchorClass {
	if u == nil {
		return nil
	}

	g := &TextChildAnchorClass{native: u}

	return g
}

func (recv *TextChildAnchorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextIter is a wrapper around the C record GtkTextIter.
type TextIter struct {
	native unsafe.Pointer
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
	// Private : dummy7
	// Private : dummy8
	// Private : dummy9
	// Private : dummy10
	// Private : dummy11
	// Private : dummy12
	// Private : dummy13
	// Private : dummy14
}

func TextIterNewFromC(u unsafe.Pointer) *TextIter {
	if u == nil {
		return nil
	}

	g := &TextIter{native: u}

	return g
}

func (recv *TextIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextMarkClass is a wrapper around the C record GtkTextMarkClass.
type TextMarkClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextMarkClassNewFromC(u unsafe.Pointer) *TextMarkClass {
	if u == nil {
		return nil
	}

	g := &TextMarkClass{native: u}

	return g
}

func (recv *TextMarkClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagClass is a wrapper around the C record GtkTextTagClass.
type TextTagClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for event
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextTagClassNewFromC(u unsafe.Pointer) *TextTagClass {
	if u == nil {
		return nil
	}

	g := &TextTagClass{native: u}

	return g
}

func (recv *TextTagClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagPrivate is a wrapper around the C record GtkTextTagPrivate.
type TextTagPrivate struct {
	native unsafe.Pointer
}

func TextTagPrivateNewFromC(u unsafe.Pointer) *TextTagPrivate {
	if u == nil {
		return nil
	}

	g := &TextTagPrivate{native: u}

	return g
}

func (recv *TextTagPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagTableClass is a wrapper around the C record GtkTextTagTableClass.
type TextTagTableClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for tag_changed
	// no type for tag_added
	// no type for tag_removed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextTagTableClassNewFromC(u unsafe.Pointer) *TextTagTableClass {
	if u == nil {
		return nil
	}

	g := &TextTagTableClass{native: u}

	return g
}

func (recv *TextTagTableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagTablePrivate is a wrapper around the C record GtkTextTagTablePrivate.
type TextTagTablePrivate struct {
	native unsafe.Pointer
}

func TextTagTablePrivateNewFromC(u unsafe.Pointer) *TextTagTablePrivate {
	if u == nil {
		return nil
	}

	g := &TextTagTablePrivate{native: u}

	return g
}

func (recv *TextTagTablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewAccessibleClass is a wrapper around the C record GtkTextViewAccessibleClass.
type TextViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func TextViewAccessibleClassNewFromC(u unsafe.Pointer) *TextViewAccessibleClass {
	if u == nil {
		return nil
	}

	g := &TextViewAccessibleClass{native: u}

	return g
}

func (recv *TextViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewAccessiblePrivate is a wrapper around the C record GtkTextViewAccessiblePrivate.
type TextViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func TextViewAccessiblePrivateNewFromC(u unsafe.Pointer) *TextViewAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &TextViewAccessiblePrivate{native: u}

	return g
}

func (recv *TextViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewClass is a wrapper around the C record GtkTextViewClass.
type TextViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for populate_popup
	// no type for move_cursor
	// no type for set_anchor
	// no type for insert_at_cursor
	// no type for delete_from_cursor
	// no type for backspace
	// no type for cut_clipboard
	// no type for copy_clipboard
	// no type for paste_clipboard
	// no type for toggle_overwrite
	// no type for create_buffer
	// no type for draw_layer
	// no type for extend_selection
	// no type for insert_emoji
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextViewClassNewFromC(u unsafe.Pointer) *TextViewClass {
	if u == nil {
		return nil
	}

	g := &TextViewClass{native: u}

	return g
}

func (recv *TextViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewPrivate is a wrapper around the C record GtkTextViewPrivate.
type TextViewPrivate struct {
	native unsafe.Pointer
}

func TextViewPrivateNewFromC(u unsafe.Pointer) *TextViewPrivate {
	if u == nil {
		return nil
	}

	g := &TextViewPrivate{native: u}

	return g
}

func (recv *TextViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemeEngine is a wrapper around the C record GtkThemeEngine.
type ThemeEngine struct {
	native unsafe.Pointer
}

func ThemeEngineNewFromC(u unsafe.Pointer) *ThemeEngine {
	if u == nil {
		return nil
	}

	g := &ThemeEngine{native: u}

	return g
}

func (recv *ThemeEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemingEngineClass is a wrapper around the C record GtkThemingEngineClass.
type ThemingEngineClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for render_line
	// no type for render_background
	// no type for render_frame
	// no type for render_frame_gap
	// no type for render_extension
	// no type for render_check
	// no type for render_option
	// no type for render_arrow
	// no type for render_expander
	// no type for render_focus
	// no type for render_layout
	// no type for render_slider
	// no type for render_handle
	// no type for render_activity
	// no type for render_icon_pixbuf
	// no type for render_icon
	// no type for render_icon_surface
	// Private : padding
}

func ThemingEngineClassNewFromC(u unsafe.Pointer) *ThemingEngineClass {
	if u == nil {
		return nil
	}

	g := &ThemingEngineClass{native: u}

	return g
}

func (recv *ThemingEngineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemingEnginePrivate is a wrapper around the C record GtkThemingEnginePrivate.
type ThemingEnginePrivate struct {
	native unsafe.Pointer
}

func ThemingEnginePrivateNewFromC(u unsafe.Pointer) *ThemingEnginePrivate {
	if u == nil {
		return nil
	}

	g := &ThemingEnginePrivate{native: u}

	return g
}

func (recv *ThemingEnginePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleActionClass is a wrapper around the C record GtkToggleActionClass.
type ToggleActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleActionClassNewFromC(u unsafe.Pointer) *ToggleActionClass {
	if u == nil {
		return nil
	}

	g := &ToggleActionClass{native: u}

	return g
}

func (recv *ToggleActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleActionEntry is a wrapper around the C record GtkToggleActionEntry.
type ToggleActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
	IsActive bool
}

func ToggleActionEntryNewFromC(u unsafe.Pointer) *ToggleActionEntry {
	if u == nil {
		return nil
	}

	g := &ToggleActionEntry{native: u}

	return g
}

func (recv *ToggleActionEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleActionPrivate is a wrapper around the C record GtkToggleActionPrivate.
type ToggleActionPrivate struct {
	native unsafe.Pointer
}

func ToggleActionPrivateNewFromC(u unsafe.Pointer) *ToggleActionPrivate {
	if u == nil {
		return nil
	}

	g := &ToggleActionPrivate{native: u}

	return g
}

func (recv *ToggleActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonAccessibleClass is a wrapper around the C record GtkToggleButtonAccessibleClass.
type ToggleButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ToggleButtonAccessibleClassNewFromC(u unsafe.Pointer) *ToggleButtonAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ToggleButtonAccessibleClass{native: u}

	return g
}

func (recv *ToggleButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonAccessiblePrivate is a wrapper around the C record GtkToggleButtonAccessiblePrivate.
type ToggleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

func ToggleButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ToggleButtonAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ToggleButtonAccessiblePrivate{native: u}

	return g
}

func (recv *ToggleButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonClass is a wrapper around the C record GtkToggleButtonClass.
type ToggleButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleButtonClassNewFromC(u unsafe.Pointer) *ToggleButtonClass {
	if u == nil {
		return nil
	}

	g := &ToggleButtonClass{native: u}

	return g
}

func (recv *ToggleButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonPrivate is a wrapper around the C record GtkToggleButtonPrivate.
type ToggleButtonPrivate struct {
	native unsafe.Pointer
}

func ToggleButtonPrivateNewFromC(u unsafe.Pointer) *ToggleButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ToggleButtonPrivate{native: u}

	return g
}

func (recv *ToggleButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleToolButtonClass is a wrapper around the C record GtkToggleToolButtonClass.
type ToggleToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleToolButtonClassNewFromC(u unsafe.Pointer) *ToggleToolButtonClass {
	if u == nil {
		return nil
	}

	g := &ToggleToolButtonClass{native: u}

	return g
}

func (recv *ToggleToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleToolButtonPrivate is a wrapper around the C record GtkToggleToolButtonPrivate.
type ToggleToolButtonPrivate struct {
	native unsafe.Pointer
}

func ToggleToolButtonPrivateNewFromC(u unsafe.Pointer) *ToggleToolButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ToggleToolButtonPrivate{native: u}

	return g
}

func (recv *ToggleToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolButtonClass is a wrapper around the C record GtkToolButtonClass.
type ToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	ButtonType gobject.Type
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolButtonClassNewFromC(u unsafe.Pointer) *ToolButtonClass {
	if u == nil {
		return nil
	}

	g := &ToolButtonClass{native: u}

	return g
}

func (recv *ToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolButtonPrivate is a wrapper around the C record GtkToolButtonPrivate.
type ToolButtonPrivate struct {
	native unsafe.Pointer
}

func ToolButtonPrivateNewFromC(u unsafe.Pointer) *ToolButtonPrivate {
	if u == nil {
		return nil
	}

	g := &ToolButtonPrivate{native: u}

	return g
}

func (recv *ToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItemClass is a wrapper around the C record GtkToolItemClass.
type ToolItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for create_menu_proxy
	// no type for toolbar_reconfigured
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolItemClassNewFromC(u unsafe.Pointer) *ToolItemClass {
	if u == nil {
		return nil
	}

	g := &ToolItemClass{native: u}

	return g
}

func (recv *ToolItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItemGroupClass is a wrapper around the C record GtkToolItemGroupClass.
type ToolItemGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolItemGroupClassNewFromC(u unsafe.Pointer) *ToolItemGroupClass {
	if u == nil {
		return nil
	}

	g := &ToolItemGroupClass{native: u}

	return g
}

func (recv *ToolItemGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItemGroupPrivate is a wrapper around the C record GtkToolItemGroupPrivate.
type ToolItemGroupPrivate struct {
	native unsafe.Pointer
}

func ToolItemGroupPrivateNewFromC(u unsafe.Pointer) *ToolItemGroupPrivate {
	if u == nil {
		return nil
	}

	g := &ToolItemGroupPrivate{native: u}

	return g
}

func (recv *ToolItemGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItemPrivate is a wrapper around the C record GtkToolItemPrivate.
type ToolItemPrivate struct {
	native unsafe.Pointer
}

func ToolItemPrivateNewFromC(u unsafe.Pointer) *ToolItemPrivate {
	if u == nil {
		return nil
	}

	g := &ToolItemPrivate{native: u}

	return g
}

func (recv *ToolItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolPaletteClass is a wrapper around the C record GtkToolPaletteClass.
type ToolPaletteClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolPaletteClassNewFromC(u unsafe.Pointer) *ToolPaletteClass {
	if u == nil {
		return nil
	}

	g := &ToolPaletteClass{native: u}

	return g
}

func (recv *ToolPaletteClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolPalettePrivate is a wrapper around the C record GtkToolPalettePrivate.
type ToolPalettePrivate struct {
	native unsafe.Pointer
}

func ToolPalettePrivateNewFromC(u unsafe.Pointer) *ToolPalettePrivate {
	if u == nil {
		return nil
	}

	g := &ToolPalettePrivate{native: u}

	return g
}

func (recv *ToolPalettePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolShellIface is a wrapper around the C record GtkToolShellIface.
type ToolShellIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_icon_size
	// no type for get_orientation
	// no type for get_style
	// no type for get_relief_style
	// no type for rebuild_menu
	// no type for get_text_orientation
	// no type for get_text_alignment
	// no type for get_ellipsize_mode
	// no type for get_text_size_group
}

func ToolShellIfaceNewFromC(u unsafe.Pointer) *ToolShellIface {
	if u == nil {
		return nil
	}

	g := &ToolShellIface{native: u}

	return g
}

func (recv *ToolShellIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolbarClass is a wrapper around the C record GtkToolbarClass.
type ToolbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for orientation_changed
	// no type for style_changed
	// no type for popup_context_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolbarClassNewFromC(u unsafe.Pointer) *ToolbarClass {
	if u == nil {
		return nil
	}

	g := &ToolbarClass{native: u}

	return g
}

func (recv *ToolbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolbarPrivate is a wrapper around the C record GtkToolbarPrivate.
type ToolbarPrivate struct {
	native unsafe.Pointer
}

func ToolbarPrivateNewFromC(u unsafe.Pointer) *ToolbarPrivate {
	if u == nil {
		return nil
	}

	g := &ToolbarPrivate{native: u}

	return g
}

func (recv *ToolbarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToplevelAccessibleClass is a wrapper around the C record GtkToplevelAccessibleClass.
type ToplevelAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func ToplevelAccessibleClassNewFromC(u unsafe.Pointer) *ToplevelAccessibleClass {
	if u == nil {
		return nil
	}

	g := &ToplevelAccessibleClass{native: u}

	return g
}

func (recv *ToplevelAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToplevelAccessiblePrivate is a wrapper around the C record GtkToplevelAccessiblePrivate.
type ToplevelAccessiblePrivate struct {
	native unsafe.Pointer
}

func ToplevelAccessiblePrivateNewFromC(u unsafe.Pointer) *ToplevelAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &ToplevelAccessiblePrivate{native: u}

	return g
}

func (recv *ToplevelAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeDragDestIface is a wrapper around the C record GtkTreeDragDestIface.
type TreeDragDestIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for drag_data_received
	// no type for row_drop_possible
}

func TreeDragDestIfaceNewFromC(u unsafe.Pointer) *TreeDragDestIface {
	if u == nil {
		return nil
	}

	g := &TreeDragDestIface{native: u}

	return g
}

func (recv *TreeDragDestIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeDragSourceIface is a wrapper around the C record GtkTreeDragSourceIface.
type TreeDragSourceIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for row_draggable
	// no type for drag_data_get
	// no type for drag_data_delete
}

func TreeDragSourceIfaceNewFromC(u unsafe.Pointer) *TreeDragSourceIface {
	if u == nil {
		return nil
	}

	g := &TreeDragSourceIface{native: u}

	return g
}

func (recv *TreeDragSourceIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeIter is a wrapper around the C record GtkTreeIter.
type TreeIter struct {
	native unsafe.Pointer
	Stamp  int32
	// user_data : no type generator for gpointer, gpointer
	// user_data2 : no type generator for gpointer, gpointer
	// user_data3 : no type generator for gpointer, gpointer
}

func TreeIterNewFromC(u unsafe.Pointer) *TreeIter {
	if u == nil {
		return nil
	}

	g := &TreeIter{native: u}

	return g
}

func (recv *TreeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelFilterClass is a wrapper around the C record GtkTreeModelFilterClass.
type TreeModelFilterClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for visible
	// no type for modify
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeModelFilterClassNewFromC(u unsafe.Pointer) *TreeModelFilterClass {
	if u == nil {
		return nil
	}

	g := &TreeModelFilterClass{native: u}

	return g
}

func (recv *TreeModelFilterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelFilterPrivate is a wrapper around the C record GtkTreeModelFilterPrivate.
type TreeModelFilterPrivate struct {
	native unsafe.Pointer
}

func TreeModelFilterPrivateNewFromC(u unsafe.Pointer) *TreeModelFilterPrivate {
	if u == nil {
		return nil
	}

	g := &TreeModelFilterPrivate{native: u}

	return g
}

func (recv *TreeModelFilterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelIface is a wrapper around the C record GtkTreeModelIface.
type TreeModelIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for row_changed
	// no type for row_inserted
	// no type for row_has_child_toggled
	// no type for row_deleted
	// no type for rows_reordered
	// no type for get_flags
	// no type for get_n_columns
	// no type for get_column_type
	// no type for get_iter
	// no type for get_path
	// no type for get_value
	// no type for iter_next
	// no type for iter_previous
	// no type for iter_children
	// no type for iter_has_child
	// no type for iter_n_children
	// no type for iter_nth_child
	// no type for iter_parent
	// no type for ref_node
	// no type for unref_node
}

func TreeModelIfaceNewFromC(u unsafe.Pointer) *TreeModelIface {
	if u == nil {
		return nil
	}

	g := &TreeModelIface{native: u}

	return g
}

func (recv *TreeModelIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelSortClass is a wrapper around the C record GtkTreeModelSortClass.
type TreeModelSortClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeModelSortClassNewFromC(u unsafe.Pointer) *TreeModelSortClass {
	if u == nil {
		return nil
	}

	g := &TreeModelSortClass{native: u}

	return g
}

func (recv *TreeModelSortClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelSortPrivate is a wrapper around the C record GtkTreeModelSortPrivate.
type TreeModelSortPrivate struct {
	native unsafe.Pointer
}

func TreeModelSortPrivateNewFromC(u unsafe.Pointer) *TreeModelSortPrivate {
	if u == nil {
		return nil
	}

	g := &TreeModelSortPrivate{native: u}

	return g
}

func (recv *TreeModelSortPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreePath is a wrapper around the C record GtkTreePath.
type TreePath struct {
	native unsafe.Pointer
}

func TreePathNewFromC(u unsafe.Pointer) *TreePath {
	if u == nil {
		return nil
	}

	g := &TreePath{native: u}

	return g
}

func (recv *TreePath) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeRowReference is a wrapper around the C record GtkTreeRowReference.
type TreeRowReference struct {
	native unsafe.Pointer
}

func TreeRowReferenceNewFromC(u unsafe.Pointer) *TreeRowReference {
	if u == nil {
		return nil
	}

	g := &TreeRowReference{native: u}

	return g
}

func (recv *TreeRowReference) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeSelectionClass is a wrapper around the C record GtkTreeSelectionClass.
type TreeSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeSelectionClassNewFromC(u unsafe.Pointer) *TreeSelectionClass {
	if u == nil {
		return nil
	}

	g := &TreeSelectionClass{native: u}

	return g
}

func (recv *TreeSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeSelectionPrivate is a wrapper around the C record GtkTreeSelectionPrivate.
type TreeSelectionPrivate struct {
	native unsafe.Pointer
}

func TreeSelectionPrivateNewFromC(u unsafe.Pointer) *TreeSelectionPrivate {
	if u == nil {
		return nil
	}

	g := &TreeSelectionPrivate{native: u}

	return g
}

func (recv *TreeSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeSortableIface is a wrapper around the C record GtkTreeSortableIface.
type TreeSortableIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for sort_column_changed
	// no type for get_sort_column_id
	// no type for set_sort_column_id
	// no type for set_sort_func
	// no type for set_default_sort_func
	// no type for has_default_sort_func
}

func TreeSortableIfaceNewFromC(u unsafe.Pointer) *TreeSortableIface {
	if u == nil {
		return nil
	}

	g := &TreeSortableIface{native: u}

	return g
}

func (recv *TreeSortableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeStoreClass is a wrapper around the C record GtkTreeStoreClass.
type TreeStoreClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeStoreClassNewFromC(u unsafe.Pointer) *TreeStoreClass {
	if u == nil {
		return nil
	}

	g := &TreeStoreClass{native: u}

	return g
}

func (recv *TreeStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeStorePrivate is a wrapper around the C record GtkTreeStorePrivate.
type TreeStorePrivate struct {
	native unsafe.Pointer
}

func TreeStorePrivateNewFromC(u unsafe.Pointer) *TreeStorePrivate {
	if u == nil {
		return nil
	}

	g := &TreeStorePrivate{native: u}

	return g
}

func (recv *TreeStorePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewAccessibleClass is a wrapper around the C record GtkTreeViewAccessibleClass.
type TreeViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func TreeViewAccessibleClassNewFromC(u unsafe.Pointer) *TreeViewAccessibleClass {
	if u == nil {
		return nil
	}

	g := &TreeViewAccessibleClass{native: u}

	return g
}

func (recv *TreeViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewAccessiblePrivate is a wrapper around the C record GtkTreeViewAccessiblePrivate.
type TreeViewAccessiblePrivate struct {
	native unsafe.Pointer
}

func TreeViewAccessiblePrivateNewFromC(u unsafe.Pointer) *TreeViewAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &TreeViewAccessiblePrivate{native: u}

	return g
}

func (recv *TreeViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewClass is a wrapper around the C record GtkTreeViewClass.
type TreeViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for row_activated
	// no type for test_expand_row
	// no type for test_collapse_row
	// no type for row_expanded
	// no type for row_collapsed
	// no type for columns_changed
	// no type for cursor_changed
	// no type for move_cursor
	// no type for select_all
	// no type for unselect_all
	// no type for select_cursor_row
	// no type for toggle_cursor_row
	// no type for expand_collapse_cursor_row
	// no type for select_cursor_parent
	// no type for start_interactive_search
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

func TreeViewClassNewFromC(u unsafe.Pointer) *TreeViewClass {
	if u == nil {
		return nil
	}

	g := &TreeViewClass{native: u}

	return g
}

func (recv *TreeViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewColumnClass is a wrapper around the C record GtkTreeViewColumnClass.
type TreeViewColumnClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeViewColumnClassNewFromC(u unsafe.Pointer) *TreeViewColumnClass {
	if u == nil {
		return nil
	}

	g := &TreeViewColumnClass{native: u}

	return g
}

func (recv *TreeViewColumnClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewColumnPrivate is a wrapper around the C record GtkTreeViewColumnPrivate.
type TreeViewColumnPrivate struct {
	native unsafe.Pointer
}

func TreeViewColumnPrivateNewFromC(u unsafe.Pointer) *TreeViewColumnPrivate {
	if u == nil {
		return nil
	}

	g := &TreeViewColumnPrivate{native: u}

	return g
}

func (recv *TreeViewColumnPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewPrivate is a wrapper around the C record GtkTreeViewPrivate.
type TreeViewPrivate struct {
	native unsafe.Pointer
}

func TreeViewPrivateNewFromC(u unsafe.Pointer) *TreeViewPrivate {
	if u == nil {
		return nil
	}

	g := &TreeViewPrivate{native: u}

	return g
}

func (recv *TreeViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UIManagerClass is a wrapper around the C record GtkUIManagerClass.
type UIManagerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add_widget
	// no type for actions_changed
	// no type for connect_proxy
	// no type for disconnect_proxy
	// no type for pre_activate
	// no type for post_activate
	// no type for get_widget
	// no type for get_action
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func UIManagerClassNewFromC(u unsafe.Pointer) *UIManagerClass {
	if u == nil {
		return nil
	}

	g := &UIManagerClass{native: u}

	return g
}

func (recv *UIManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UIManagerPrivate is a wrapper around the C record GtkUIManagerPrivate.
type UIManagerPrivate struct {
	native unsafe.Pointer
}

func UIManagerPrivateNewFromC(u unsafe.Pointer) *UIManagerPrivate {
	if u == nil {
		return nil
	}

	g := &UIManagerPrivate{native: u}

	return g
}

func (recv *UIManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VBoxClass is a wrapper around the C record GtkVBoxClass.
type VBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VBoxClassNewFromC(u unsafe.Pointer) *VBoxClass {
	if u == nil {
		return nil
	}

	g := &VBoxClass{native: u}

	return g
}

func (recv *VBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VButtonBoxClass is a wrapper around the C record GtkVButtonBoxClass.
type VButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VButtonBoxClassNewFromC(u unsafe.Pointer) *VButtonBoxClass {
	if u == nil {
		return nil
	}

	g := &VButtonBoxClass{native: u}

	return g
}

func (recv *VButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VPanedClass is a wrapper around the C record GtkVPanedClass.
type VPanedClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VPanedClassNewFromC(u unsafe.Pointer) *VPanedClass {
	if u == nil {
		return nil
	}

	g := &VPanedClass{native: u}

	return g
}

func (recv *VPanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScaleClass is a wrapper around the C record GtkVScaleClass.
type VScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VScaleClassNewFromC(u unsafe.Pointer) *VScaleClass {
	if u == nil {
		return nil
	}

	g := &VScaleClass{native: u}

	return g
}

func (recv *VScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScrollbarClass is a wrapper around the C record GtkVScrollbarClass.
type VScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VScrollbarClassNewFromC(u unsafe.Pointer) *VScrollbarClass {
	if u == nil {
		return nil
	}

	g := &VScrollbarClass{native: u}

	return g
}

func (recv *VScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VSeparatorClass is a wrapper around the C record GtkVSeparatorClass.
type VSeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func VSeparatorClassNewFromC(u unsafe.Pointer) *VSeparatorClass {
	if u == nil {
		return nil
	}

	g := &VSeparatorClass{native: u}

	return g
}

func (recv *VSeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ViewportClass is a wrapper around the C record GtkViewportClass.
type ViewportClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ViewportClassNewFromC(u unsafe.Pointer) *ViewportClass {
	if u == nil {
		return nil
	}

	g := &ViewportClass{native: u}

	return g
}

func (recv *ViewportClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ViewportPrivate is a wrapper around the C record GtkViewportPrivate.
type ViewportPrivate struct {
	native unsafe.Pointer
}

func ViewportPrivateNewFromC(u unsafe.Pointer) *ViewportPrivate {
	if u == nil {
		return nil
	}

	g := &ViewportPrivate{native: u}

	return g
}

func (recv *ViewportPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeButtonClass is a wrapper around the C record GtkVolumeButtonClass.
type VolumeButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func VolumeButtonClassNewFromC(u unsafe.Pointer) *VolumeButtonClass {
	if u == nil {
		return nil
	}

	g := &VolumeButtonClass{native: u}

	return g
}

func (recv *VolumeButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessibleClass is a wrapper around the C record GtkWidgetAccessibleClass.
type WidgetAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for notify_gtk
}

func WidgetAccessibleClassNewFromC(u unsafe.Pointer) *WidgetAccessibleClass {
	if u == nil {
		return nil
	}

	g := &WidgetAccessibleClass{native: u}

	return g
}

func (recv *WidgetAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessiblePrivate is a wrapper around the C record GtkWidgetAccessiblePrivate.
type WidgetAccessiblePrivate struct {
	native unsafe.Pointer
}

func WidgetAccessiblePrivateNewFromC(u unsafe.Pointer) *WidgetAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &WidgetAccessiblePrivate{native: u}

	return g
}

func (recv *WidgetAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetClass is a wrapper around the C record GtkWidgetClass.
type WidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	ActivateSignal uint32
	// no type for dispatch_child_properties_changed
	// no type for destroy
	// no type for show
	// no type for show_all
	// no type for hide
	// no type for map
	// no type for unmap
	// no type for realize
	// no type for unrealize
	// no type for size_allocate
	// no type for state_changed
	// no type for state_flags_changed
	// no type for parent_set
	// no type for hierarchy_changed
	// no type for style_set
	// no type for direction_changed
	// no type for grab_notify
	// no type for child_notify
	// no type for draw
	// no type for get_request_mode
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for mnemonic_activate
	// no type for grab_focus
	// no type for focus
	// no type for move_focus
	// no type for keynav_failed
	// no type for event
	// no type for button_press_event
	// no type for button_release_event
	// no type for scroll_event
	// no type for motion_notify_event
	// no type for delete_event
	// no type for destroy_event
	// no type for key_press_event
	// no type for key_release_event
	// no type for enter_notify_event
	// no type for leave_notify_event
	// no type for configure_event
	// no type for focus_in_event
	// no type for focus_out_event
	// no type for map_event
	// no type for unmap_event
	// no type for property_notify_event
	// no type for selection_clear_event
	// no type for selection_request_event
	// no type for selection_notify_event
	// no type for proximity_in_event
	// no type for proximity_out_event
	// no type for visibility_notify_event
	// no type for window_state_event
	// no type for damage_event
	// no type for grab_broken_event
	// no type for selection_get
	// no type for selection_received
	// no type for drag_begin
	// no type for drag_end
	// no type for drag_data_get
	// no type for drag_data_delete
	// no type for drag_leave
	// no type for drag_motion
	// no type for drag_drop
	// no type for drag_data_received
	// no type for drag_failed
	// no type for popup_menu
	// no type for show_help
	// no type for get_accessible
	// no type for screen_changed
	// no type for can_activate_accel
	// no type for composited_changed
	// no type for query_tooltip
	// no type for compute_expand
	// no type for adjust_size_request
	// no type for adjust_size_allocation
	// no type for style_updated
	// no type for touch_event
	// no type for get_preferred_height_and_baseline_for_width
	// no type for adjust_baseline_request
	// no type for adjust_baseline_allocation
	// no type for queue_draw_region
	// Private : priv
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
}

func WidgetClassNewFromC(u unsafe.Pointer) *WidgetClass {
	if u == nil {
		return nil
	}

	g := &WidgetClass{native: u}

	return g
}

func (recv *WidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetClassPrivate is a wrapper around the C record GtkWidgetClassPrivate.
type WidgetClassPrivate struct {
	native unsafe.Pointer
}

func WidgetClassPrivateNewFromC(u unsafe.Pointer) *WidgetClassPrivate {
	if u == nil {
		return nil
	}

	g := &WidgetClassPrivate{native: u}

	return g
}

func (recv *WidgetClassPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetPath is a wrapper around the C record GtkWidgetPath.
type WidgetPath struct {
	native unsafe.Pointer
}

func WidgetPathNewFromC(u unsafe.Pointer) *WidgetPath {
	if u == nil {
		return nil
	}

	g := &WidgetPath{native: u}

	return g
}

func (recv *WidgetPath) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetPrivate is a wrapper around the C record GtkWidgetPrivate.
type WidgetPrivate struct {
	native unsafe.Pointer
}

func WidgetPrivateNewFromC(u unsafe.Pointer) *WidgetPrivate {
	if u == nil {
		return nil
	}

	g := &WidgetPrivate{native: u}

	return g
}

func (recv *WidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowAccessibleClass is a wrapper around the C record GtkWindowAccessibleClass.
type WindowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

func WindowAccessibleClassNewFromC(u unsafe.Pointer) *WindowAccessibleClass {
	if u == nil {
		return nil
	}

	g := &WindowAccessibleClass{native: u}

	return g
}

func (recv *WindowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowAccessiblePrivate is a wrapper around the C record GtkWindowAccessiblePrivate.
type WindowAccessiblePrivate struct {
	native unsafe.Pointer
}

func WindowAccessiblePrivateNewFromC(u unsafe.Pointer) *WindowAccessiblePrivate {
	if u == nil {
		return nil
	}

	g := &WindowAccessiblePrivate{native: u}

	return g
}

func (recv *WindowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowClass is a wrapper around the C record GtkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for set_focus
	// no type for activate_focus
	// no type for activate_default
	// no type for keys_changed
	// no type for enable_debugging
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

func WindowClassNewFromC(u unsafe.Pointer) *WindowClass {
	if u == nil {
		return nil
	}

	g := &WindowClass{native: u}

	return g
}

func (recv *WindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGeometryInfo is a wrapper around the C record GtkWindowGeometryInfo.
type WindowGeometryInfo struct {
	native unsafe.Pointer
}

func WindowGeometryInfoNewFromC(u unsafe.Pointer) *WindowGeometryInfo {
	if u == nil {
		return nil
	}

	g := &WindowGeometryInfo{native: u}

	return g
}

func (recv *WindowGeometryInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGroupClass is a wrapper around the C record GtkWindowGroupClass.
type WindowGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func WindowGroupClassNewFromC(u unsafe.Pointer) *WindowGroupClass {
	if u == nil {
		return nil
	}

	g := &WindowGroupClass{native: u}

	return g
}

func (recv *WindowGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGroupPrivate is a wrapper around the C record GtkWindowGroupPrivate.
type WindowGroupPrivate struct {
	native unsafe.Pointer
}

func WindowGroupPrivateNewFromC(u unsafe.Pointer) *WindowGroupPrivate {
	if u == nil {
		return nil
	}

	g := &WindowGroupPrivate{native: u}

	return g
}

func (recv *WindowGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowPrivate is a wrapper around the C record GtkWindowPrivate.
type WindowPrivate struct {
	native unsafe.Pointer
}

func WindowPrivateNewFromC(u unsafe.Pointer) *WindowPrivate {
	if u == nil {
		return nil
	}

	g := &WindowPrivate{native: u}

	return g
}

func (recv *WindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
