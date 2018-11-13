// This is a generated file - DO NOT EDIT

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

/*

C type

GtkAboutDialogClass
*/
type AboutDialogClass struct {
	native *C.GtkAboutDialogClass
	// parent_class : record
	// no type for activate_link
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AboutDialogClassNewFromC(u unsafe.Pointer) *AboutDialogClass {
	c := (*C.GtkAboutDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &AboutDialogClass{native: c}

	return g
}

func (recv *AboutDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAboutDialogPrivate
*/
type AboutDialogPrivate struct {
	native *C.GtkAboutDialogPrivate
}

func AboutDialogPrivateNewFromC(u unsafe.Pointer) *AboutDialogPrivate {
	c := (*C.GtkAboutDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AboutDialogPrivate{native: c}

	return g
}

func (recv *AboutDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelGroupClass
*/
type AccelGroupClass struct {
	native *C.GtkAccelGroupClass
	// parent_class : record
	// no type for accel_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AccelGroupClassNewFromC(u unsafe.Pointer) *AccelGroupClass {
	c := (*C.GtkAccelGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &AccelGroupClass{native: c}

	return g
}

func (recv *AccelGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelGroupEntry
*/
type AccelGroupEntry struct {
	native *C.GtkAccelGroupEntry
	// key : record
	// closure : record
	AccelPathQuark glib.Quark
}

func AccelGroupEntryNewFromC(u unsafe.Pointer) *AccelGroupEntry {
	c := (*C.GtkAccelGroupEntry)(u)
	if c == nil {
		return nil
	}

	g := &AccelGroupEntry{
		AccelPathQuark: (glib.Quark)(c.accel_path_quark),
		native:         c,
	}

	return g
}

func (recv *AccelGroupEntry) ToC() unsafe.Pointer {
	recv.native.accel_path_quark =
		(C.GQuark)(recv.AccelPathQuark)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelGroupPrivate
*/
type AccelGroupPrivate struct {
	native *C.GtkAccelGroupPrivate
}

func AccelGroupPrivateNewFromC(u unsafe.Pointer) *AccelGroupPrivate {
	c := (*C.GtkAccelGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AccelGroupPrivate{native: c}

	return g
}

func (recv *AccelGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelKey
*/
type AccelKey struct {
	native    *C.GtkAccelKey
	AccelKey  uint32
	AccelMods gdk.ModifierType
	// Bitfield not supported : 16 accel_flags
}

func AccelKeyNewFromC(u unsafe.Pointer) *AccelKey {
	c := (*C.GtkAccelKey)(u)
	if c == nil {
		return nil
	}

	g := &AccelKey{
		AccelKey:  (uint32)(c.accel_key),
		AccelMods: (gdk.ModifierType)(c.accel_mods),
		native:    c,
	}

	return g
}

func (recv *AccelKey) ToC() unsafe.Pointer {
	recv.native.accel_key =
		(C.guint)(recv.AccelKey)
	recv.native.accel_mods =
		(C.GdkModifierType)(recv.AccelMods)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelLabelClass
*/
type AccelLabelClass struct {
	native *C.GtkAccelLabelClass
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
	c := (*C.GtkAccelLabelClass)(u)
	if c == nil {
		return nil
	}

	g := &AccelLabelClass{
		ModNameAlt:     C.GoString(c.mod_name_alt),
		ModNameControl: C.GoString(c.mod_name_control),
		ModNameShift:   C.GoString(c.mod_name_shift),
		ModSeparator:   C.GoString(c.mod_separator),
		SignalQuote1:   C.GoString(c.signal_quote1),
		SignalQuote2:   C.GoString(c.signal_quote2),
		native:         c,
	}

	return g
}

func (recv *AccelLabelClass) ToC() unsafe.Pointer {
	recv.native.signal_quote1 =
		C.CString(recv.SignalQuote1)
	recv.native.signal_quote2 =
		C.CString(recv.SignalQuote2)
	recv.native.mod_name_shift =
		C.CString(recv.ModNameShift)
	recv.native.mod_name_control =
		C.CString(recv.ModNameControl)
	recv.native.mod_name_alt =
		C.CString(recv.ModNameAlt)
	recv.native.mod_separator =
		C.CString(recv.ModSeparator)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelLabelPrivate
*/
type AccelLabelPrivate struct {
	native *C.GtkAccelLabelPrivate
}

func AccelLabelPrivateNewFromC(u unsafe.Pointer) *AccelLabelPrivate {
	c := (*C.GtkAccelLabelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AccelLabelPrivate{native: c}

	return g
}

func (recv *AccelLabelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccelMapClass
*/
type AccelMapClass struct {
	native *C.GtkAccelMapClass
}

func AccelMapClassNewFromC(u unsafe.Pointer) *AccelMapClass {
	c := (*C.GtkAccelMapClass)(u)
	if c == nil {
		return nil
	}

	g := &AccelMapClass{native: c}

	return g
}

func (recv *AccelMapClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccessibleClass
*/
type AccessibleClass struct {
	native *C.GtkAccessibleClass
	// parent_class : record
	// no type for connect_widget_destroyed
	// no type for widget_set
	// no type for widget_unset
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AccessibleClassNewFromC(u unsafe.Pointer) *AccessibleClass {
	c := (*C.GtkAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &AccessibleClass{native: c}

	return g
}

func (recv *AccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAccessiblePrivate
*/
type AccessiblePrivate struct {
	native *C.GtkAccessiblePrivate
}

func AccessiblePrivateNewFromC(u unsafe.Pointer) *AccessiblePrivate {
	c := (*C.GtkAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &AccessiblePrivate{native: c}

	return g
}

func (recv *AccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionBarClass
*/
type ActionBarClass struct {
	native *C.GtkActionBarClass
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ActionBarClassNewFromC(u unsafe.Pointer) *ActionBarClass {
	c := (*C.GtkActionBarClass)(u)
	if c == nil {
		return nil
	}

	g := &ActionBarClass{native: c}

	return g
}

func (recv *ActionBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionBarPrivate
*/
type ActionBarPrivate struct {
	native *C.GtkActionBarPrivate
}

func ActionBarPrivateNewFromC(u unsafe.Pointer) *ActionBarPrivate {
	c := (*C.GtkActionBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ActionBarPrivate{native: c}

	return g
}

func (recv *ActionBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionClass
*/
type ActionClass struct {
	native *C.GtkActionClass
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
	c := (*C.GtkActionClass)(u)
	if c == nil {
		return nil
	}

	g := &ActionClass{native: c}

	return g
}

func (recv *ActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkActionEntry structs are used with gtk_action_group_add_actions() to
// construct actions.
/*

C type

GtkActionEntry
*/
type ActionEntry struct {
	native      *C.GtkActionEntry
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
}

func ActionEntryNewFromC(u unsafe.Pointer) *ActionEntry {
	c := (*C.GtkActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &ActionEntry{
		Accelerator: C.GoString(c.accelerator),
		Label:       C.GoString(c.label),
		Name:        C.GoString(c.name),
		StockId:     C.GoString(c.stock_id),
		Tooltip:     C.GoString(c.tooltip),
		native:      c,
	}

	return g
}

func (recv *ActionEntry) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.stock_id =
		C.CString(recv.StockId)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.accelerator =
		C.CString(recv.Accelerator)
	recv.native.tooltip =
		C.CString(recv.Tooltip)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionGroupClass
*/
type ActionGroupClass struct {
	native *C.GtkActionGroupClass
	// parent_class : record
	// no type for get_action
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ActionGroupClassNewFromC(u unsafe.Pointer) *ActionGroupClass {
	c := (*C.GtkActionGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroupClass{native: c}

	return g
}

func (recv *ActionGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionGroupPrivate
*/
type ActionGroupPrivate struct {
	native *C.GtkActionGroupPrivate
}

func ActionGroupPrivateNewFromC(u unsafe.Pointer) *ActionGroupPrivate {
	c := (*C.GtkActionGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroupPrivate{native: c}

	return g
}

func (recv *ActionGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkActionPrivate
*/
type ActionPrivate struct {
	native *C.GtkActionPrivate
}

func ActionPrivateNewFromC(u unsafe.Pointer) *ActionPrivate {
	c := (*C.GtkActionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ActionPrivate{native: c}

	return g
}

func (recv *ActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The interface vtable for #GtkActionable.
/*

C type

GtkActionableInterface
*/
type ActionableInterface struct {
	native *C.GtkActionableInterface
	// Private : g_iface
	// no type for get_action_name
	// no type for set_action_name
	// no type for get_action_target_value
	// no type for set_action_target_value
}

func ActionableInterfaceNewFromC(u unsafe.Pointer) *ActionableInterface {
	c := (*C.GtkActionableInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionableInterface{native: c}

	return g
}

func (recv *ActionableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAdjustmentClass
*/
type AdjustmentClass struct {
	native *C.GtkAdjustmentClass
	// parent_class : record
	// no type for changed
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AdjustmentClassNewFromC(u unsafe.Pointer) *AdjustmentClass {
	c := (*C.GtkAdjustmentClass)(u)
	if c == nil {
		return nil
	}

	g := &AdjustmentClass{native: c}

	return g
}

func (recv *AdjustmentClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAdjustmentPrivate
*/
type AdjustmentPrivate struct {
	native *C.GtkAdjustmentPrivate
}

func AdjustmentPrivateNewFromC(u unsafe.Pointer) *AdjustmentPrivate {
	c := (*C.GtkAdjustmentPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AdjustmentPrivate{native: c}

	return g
}

func (recv *AdjustmentPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAlignmentClass
*/
type AlignmentClass struct {
	native *C.GtkAlignmentClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AlignmentClassNewFromC(u unsafe.Pointer) *AlignmentClass {
	c := (*C.GtkAlignmentClass)(u)
	if c == nil {
		return nil
	}

	g := &AlignmentClass{native: c}

	return g
}

func (recv *AlignmentClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAlignmentPrivate
*/
type AlignmentPrivate struct {
	native *C.GtkAlignmentPrivate
}

func AlignmentPrivateNewFromC(u unsafe.Pointer) *AlignmentPrivate {
	c := (*C.GtkAlignmentPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AlignmentPrivate{native: c}

	return g
}

func (recv *AlignmentPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserButtonClass
*/
type AppChooserButtonClass struct {
	native *C.GtkAppChooserButtonClass
	// parent_class : record
	// no type for custom_item_activated
	// Private : padding
}

func AppChooserButtonClassNewFromC(u unsafe.Pointer) *AppChooserButtonClass {
	c := (*C.GtkAppChooserButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserButtonClass{native: c}

	return g
}

func (recv *AppChooserButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserButtonPrivate
*/
type AppChooserButtonPrivate struct {
	native *C.GtkAppChooserButtonPrivate
}

func AppChooserButtonPrivateNewFromC(u unsafe.Pointer) *AppChooserButtonPrivate {
	c := (*C.GtkAppChooserButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserButtonPrivate{native: c}

	return g
}

func (recv *AppChooserButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserDialogClass
*/
type AppChooserDialogClass struct {
	native *C.GtkAppChooserDialogClass
	// parent_class : record
	// Private : padding
}

func AppChooserDialogClassNewFromC(u unsafe.Pointer) *AppChooserDialogClass {
	c := (*C.GtkAppChooserDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserDialogClass{native: c}

	return g
}

func (recv *AppChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserDialogPrivate
*/
type AppChooserDialogPrivate struct {
	native *C.GtkAppChooserDialogPrivate
}

func AppChooserDialogPrivateNewFromC(u unsafe.Pointer) *AppChooserDialogPrivate {
	c := (*C.GtkAppChooserDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserDialogPrivate{native: c}

	return g
}

func (recv *AppChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserWidgetClass
*/
type AppChooserWidgetClass struct {
	native *C.GtkAppChooserWidgetClass
	// parent_class : record
	// no type for application_selected
	// no type for application_activated
	// no type for populate_popup
	// Private : padding
}

func AppChooserWidgetClassNewFromC(u unsafe.Pointer) *AppChooserWidgetClass {
	c := (*C.GtkAppChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserWidgetClass{native: c}

	return g
}

func (recv *AppChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAppChooserWidgetPrivate
*/
type AppChooserWidgetPrivate struct {
	native *C.GtkAppChooserWidgetPrivate
}

func AppChooserWidgetPrivateNewFromC(u unsafe.Pointer) *AppChooserWidgetPrivate {
	c := (*C.GtkAppChooserWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserWidgetPrivate{native: c}

	return g
}

func (recv *AppChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkApplicationClass
*/
type ApplicationClass struct {
	native *C.GtkApplicationClass
	// parent_class : record
	// no type for window_added
	// no type for window_removed
	// Private : padding
}

func ApplicationClassNewFromC(u unsafe.Pointer) *ApplicationClass {
	c := (*C.GtkApplicationClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationClass{native: c}

	return g
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkApplicationPrivate
*/
type ApplicationPrivate struct {
	native *C.GtkApplicationPrivate
}

func ApplicationPrivateNewFromC(u unsafe.Pointer) *ApplicationPrivate {
	c := (*C.GtkApplicationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationPrivate{native: c}

	return g
}

func (recv *ApplicationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkApplicationWindowClass
*/
type ApplicationWindowClass struct {
	native *C.GtkApplicationWindowClass
	// parent_class : record
	// Private : padding
}

func ApplicationWindowClassNewFromC(u unsafe.Pointer) *ApplicationWindowClass {
	c := (*C.GtkApplicationWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationWindowClass{native: c}

	return g
}

func (recv *ApplicationWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkApplicationWindowPrivate
*/
type ApplicationWindowPrivate struct {
	native *C.GtkApplicationWindowPrivate
}

func ApplicationWindowPrivateNewFromC(u unsafe.Pointer) *ApplicationWindowPrivate {
	c := (*C.GtkApplicationWindowPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationWindowPrivate{native: c}

	return g
}

func (recv *ApplicationWindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkArrowAccessibleClass
*/
type ArrowAccessibleClass struct {
	native *C.GtkArrowAccessibleClass
	// parent_class : record
}

func ArrowAccessibleClassNewFromC(u unsafe.Pointer) *ArrowAccessibleClass {
	c := (*C.GtkArrowAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ArrowAccessibleClass{native: c}

	return g
}

func (recv *ArrowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkArrowAccessiblePrivate
*/
type ArrowAccessiblePrivate struct {
	native *C.GtkArrowAccessiblePrivate
}

func ArrowAccessiblePrivateNewFromC(u unsafe.Pointer) *ArrowAccessiblePrivate {
	c := (*C.GtkArrowAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ArrowAccessiblePrivate{native: c}

	return g
}

func (recv *ArrowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkArrowClass
*/
type ArrowClass struct {
	native *C.GtkArrowClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ArrowClassNewFromC(u unsafe.Pointer) *ArrowClass {
	c := (*C.GtkArrowClass)(u)
	if c == nil {
		return nil
	}

	g := &ArrowClass{native: c}

	return g
}

func (recv *ArrowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkArrowPrivate
*/
type ArrowPrivate struct {
	native *C.GtkArrowPrivate
}

func ArrowPrivateNewFromC(u unsafe.Pointer) *ArrowPrivate {
	c := (*C.GtkArrowPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ArrowPrivate{native: c}

	return g
}

func (recv *ArrowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAspectFrameClass
*/
type AspectFrameClass struct {
	native *C.GtkAspectFrameClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func AspectFrameClassNewFromC(u unsafe.Pointer) *AspectFrameClass {
	c := (*C.GtkAspectFrameClass)(u)
	if c == nil {
		return nil
	}

	g := &AspectFrameClass{native: c}

	return g
}

func (recv *AspectFrameClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAspectFramePrivate
*/
type AspectFramePrivate struct {
	native *C.GtkAspectFramePrivate
}

func AspectFramePrivateNewFromC(u unsafe.Pointer) *AspectFramePrivate {
	c := (*C.GtkAspectFramePrivate)(u)
	if c == nil {
		return nil
	}

	g := &AspectFramePrivate{native: c}

	return g
}

func (recv *AspectFramePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAssistantClass
*/
type AssistantClass struct {
	native *C.GtkAssistantClass
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
	c := (*C.GtkAssistantClass)(u)
	if c == nil {
		return nil
	}

	g := &AssistantClass{native: c}

	return g
}

func (recv *AssistantClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkAssistantPrivate
*/
type AssistantPrivate struct {
	native *C.GtkAssistantPrivate
}

func AssistantPrivateNewFromC(u unsafe.Pointer) *AssistantPrivate {
	c := (*C.GtkAssistantPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AssistantPrivate{native: c}

	return g
}

func (recv *AssistantPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBinClass
*/
type BinClass struct {
	native *C.GtkBinClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func BinClassNewFromC(u unsafe.Pointer) *BinClass {
	c := (*C.GtkBinClass)(u)
	if c == nil {
		return nil
	}

	g := &BinClass{native: c}

	return g
}

func (recv *BinClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBinPrivate
*/
type BinPrivate struct {
	native *C.GtkBinPrivate
}

func BinPrivateNewFromC(u unsafe.Pointer) *BinPrivate {
	c := (*C.GtkBinPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BinPrivate{native: c}

	return g
}

func (recv *BinPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A #GtkBindingArg holds the data associated with
// an argument for a key binding signal emission as
// stored in #GtkBindingSignal.
/*

C type

GtkBindingArg
*/
type BindingArg struct {
	native  *C.GtkBindingArg
	ArgType gobject.Type
}

func BindingArgNewFromC(u unsafe.Pointer) *BindingArg {
	c := (*C.GtkBindingArg)(u)
	if c == nil {
		return nil
	}

	g := &BindingArg{
		ArgType: (gobject.Type)(c.arg_type),
		native:  c,
	}

	return g
}

func (recv *BindingArg) ToC() unsafe.Pointer {
	recv.native.arg_type =
		(C.GType)(recv.ArgType)

	return (unsafe.Pointer)(recv.native)
}

// Each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
/*

C type

GtkBindingEntry
*/
type BindingEntry struct {
	native    *C.GtkBindingEntry
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
	c := (*C.GtkBindingEntry)(u)
	if c == nil {
		return nil
	}

	g := &BindingEntry{
		Keyval:    (uint32)(c.keyval),
		Modifiers: (gdk.ModifierType)(c.modifiers),
		native:    c,
	}

	return g
}

func (recv *BindingEntry) ToC() unsafe.Pointer {
	recv.native.keyval =
		(C.guint)(recv.Keyval)
	recv.native.modifiers =
		(C.GdkModifierType)(recv.Modifiers)

	return (unsafe.Pointer)(recv.native)
}

// A binding set maintains a list of activatable key bindings.
// A single binding set can match multiple types of widgets.
// Similar to style contexts, can be matched by any information contained
// in a widgets #GtkWidgetPath. When a binding within a set is matched upon
// activation, an action signal is emitted on the target widget to carry out
// the actual activation.
/*

C type

GtkBindingSet
*/
type BindingSet struct {
	native   *C.GtkBindingSet
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
	c := (*C.GtkBindingSet)(u)
	if c == nil {
		return nil
	}

	g := &BindingSet{
		Priority: (int32)(c.priority),
		SetName:  C.GoString(c.set_name),
		native:   c,
	}

	return g
}

func (recv *BindingSet) ToC() unsafe.Pointer {
	recv.native.set_name =
		C.CString(recv.SetName)
	recv.native.priority =
		(C.gint)(recv.Priority)

	return (unsafe.Pointer)(recv.native)
}

// Find a key binding matching @keyval and @modifiers within
// @binding_set and activate the binding on @object.
/*

C function

gtk_binding_set_activate
*/
func (recv *BindingSet) Activate(keyval uint32, modifiers gdk.ModifierType, object *gobject.Object) bool {
	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	retC := C.gtk_binding_set_activate((*C.GtkBindingSet)(recv.native), c_keyval, c_modifiers, c_object)
	retGo := retC == C.TRUE

	return retGo
}

// This function was used internally by the GtkRC parsing mechanism
// to assign match patterns to #GtkBindingSet structures.
//
// In GTK+ 3, these match patterns are unused.
/*

C function

gtk_binding_set_add_path
*/
func (recv *BindingSet) AddPath(pathType PathType, pathPattern string, priority PathPriorityType) {
	c_path_type := (C.GtkPathType)(pathType)

	c_path_pattern := C.CString(pathPattern)
	defer C.free(unsafe.Pointer(c_path_pattern))

	c_priority := (C.GtkPathPriorityType)(priority)

	C.gtk_binding_set_add_path((*C.GtkBindingSet)(recv.native), c_path_type, c_path_pattern, c_priority)

	return
}

// A GtkBindingSignal stores the necessary information to
// activate a widget in response to a key press via a signal
// emission.
/*

C type

GtkBindingSignal
*/
type BindingSignal struct {
	native *C.GtkBindingSignal
	// next : record
	SignalName string
	NArgs      uint32
	// no type for args
}

func BindingSignalNewFromC(u unsafe.Pointer) *BindingSignal {
	c := (*C.GtkBindingSignal)(u)
	if c == nil {
		return nil
	}

	g := &BindingSignal{
		NArgs:      (uint32)(c.n_args),
		SignalName: C.GoString(c.signal_name),
		native:     c,
	}

	return g
}

func (recv *BindingSignal) ToC() unsafe.Pointer {
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.n_args =
		(C.guint)(recv.NArgs)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBooleanCellAccessibleClass
*/
type BooleanCellAccessibleClass struct {
	native *C.GtkBooleanCellAccessibleClass
	// parent_class : record
}

func BooleanCellAccessibleClassNewFromC(u unsafe.Pointer) *BooleanCellAccessibleClass {
	c := (*C.GtkBooleanCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &BooleanCellAccessibleClass{native: c}

	return g
}

func (recv *BooleanCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBooleanCellAccessiblePrivate
*/
type BooleanCellAccessiblePrivate struct {
	native *C.GtkBooleanCellAccessiblePrivate
}

func BooleanCellAccessiblePrivateNewFromC(u unsafe.Pointer) *BooleanCellAccessiblePrivate {
	c := (*C.GtkBooleanCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &BooleanCellAccessiblePrivate{native: c}

	return g
}

func (recv *BooleanCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A struct that specifies a border around a rectangular area
// that can be of different width on each side.
/*

C type

GtkBorder
*/
type Border struct {
	native *C.GtkBorder
	Left   int16
	Right  int16
	Top    int16
	Bottom int16
}

func BorderNewFromC(u unsafe.Pointer) *Border {
	c := (*C.GtkBorder)(u)
	if c == nil {
		return nil
	}

	g := &Border{
		Bottom: (int16)(c.bottom),
		Left:   (int16)(c.left),
		Right:  (int16)(c.right),
		Top:    (int16)(c.top),
		native: c,
	}

	return g
}

func (recv *Border) ToC() unsafe.Pointer {
	recv.native.left =
		(C.gint16)(recv.Left)
	recv.native.right =
		(C.gint16)(recv.Right)
	recv.native.top =
		(C.gint16)(recv.Top)
	recv.native.bottom =
		(C.gint16)(recv.Bottom)

	return (unsafe.Pointer)(recv.native)
}

// Copies a #GtkBorder-struct.
/*

C function

gtk_border_copy
*/
func (recv *Border) Copy() *Border {
	retC := C.gtk_border_copy((*C.GtkBorder)(recv.native))
	retGo := BorderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a #GtkBorder-struct.
/*

C function

gtk_border_free
*/
func (recv *Border) Free() {
	C.gtk_border_free((*C.GtkBorder)(recv.native))

	return
}

/*

C type

GtkBoxClass
*/
type BoxClass struct {
	native *C.GtkBoxClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func BoxClassNewFromC(u unsafe.Pointer) *BoxClass {
	c := (*C.GtkBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &BoxClass{native: c}

	return g
}

func (recv *BoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBoxPrivate
*/
type BoxPrivate struct {
	native *C.GtkBoxPrivate
}

func BoxPrivateNewFromC(u unsafe.Pointer) *BoxPrivate {
	c := (*C.GtkBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BoxPrivate{native: c}

	return g
}

func (recv *BoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GtkBuildableIface interface contains method that are
// necessary to allow #GtkBuilder to construct an object from
// a #GtkBuilder UI definition.
/*

C type

GtkBuildableIface
*/
type BuildableIface struct {
	native *C.GtkBuildableIface
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
	c := (*C.GtkBuildableIface)(u)
	if c == nil {
		return nil
	}

	g := &BuildableIface{native: c}

	return g
}

func (recv *BuildableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBuilderClass
*/
type BuilderClass struct {
	native *C.GtkBuilderClass
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
	c := (*C.GtkBuilderClass)(u)
	if c == nil {
		return nil
	}

	g := &BuilderClass{native: c}

	return g
}

func (recv *BuilderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkBuilderPrivate
*/
type BuilderPrivate struct {
	native *C.GtkBuilderPrivate
}

func BuilderPrivateNewFromC(u unsafe.Pointer) *BuilderPrivate {
	c := (*C.GtkBuilderPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BuilderPrivate{native: c}

	return g
}

func (recv *BuilderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonAccessibleClass
*/
type ButtonAccessibleClass struct {
	native *C.GtkButtonAccessibleClass
	// parent_class : record
}

func ButtonAccessibleClassNewFromC(u unsafe.Pointer) *ButtonAccessibleClass {
	c := (*C.GtkButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ButtonAccessibleClass{native: c}

	return g
}

func (recv *ButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonAccessiblePrivate
*/
type ButtonAccessiblePrivate struct {
	native *C.GtkButtonAccessiblePrivate
}

func ButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ButtonAccessiblePrivate {
	c := (*C.GtkButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ButtonAccessiblePrivate{native: c}

	return g
}

func (recv *ButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonBoxClass
*/
type ButtonBoxClass struct {
	native *C.GtkButtonBoxClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ButtonBoxClassNewFromC(u unsafe.Pointer) *ButtonBoxClass {
	c := (*C.GtkButtonBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &ButtonBoxClass{native: c}

	return g
}

func (recv *ButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonBoxPrivate
*/
type ButtonBoxPrivate struct {
	native *C.GtkButtonBoxPrivate
}

func ButtonBoxPrivateNewFromC(u unsafe.Pointer) *ButtonBoxPrivate {
	c := (*C.GtkButtonBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ButtonBoxPrivate{native: c}

	return g
}

func (recv *ButtonBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonClass
*/
type ButtonClass struct {
	native *C.GtkButtonClass
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
	c := (*C.GtkButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ButtonClass{native: c}

	return g
}

func (recv *ButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkButtonPrivate
*/
type ButtonPrivate struct {
	native *C.GtkButtonPrivate
}

func ButtonPrivateNewFromC(u unsafe.Pointer) *ButtonPrivate {
	c := (*C.GtkButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ButtonPrivate{native: c}

	return g
}

func (recv *ButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCalendarClass
*/
type CalendarClass struct {
	native *C.GtkCalendarClass
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
	c := (*C.GtkCalendarClass)(u)
	if c == nil {
		return nil
	}

	g := &CalendarClass{native: c}

	return g
}

func (recv *CalendarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCalendarPrivate
*/
type CalendarPrivate struct {
	native *C.GtkCalendarPrivate
}

func CalendarPrivateNewFromC(u unsafe.Pointer) *CalendarPrivate {
	c := (*C.GtkCalendarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CalendarPrivate{native: c}

	return g
}

func (recv *CalendarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAccessibleClass
*/
type CellAccessibleClass struct {
	native *C.GtkCellAccessibleClass
	// parent_class : record
	// no type for update_cache
}

func CellAccessibleClassNewFromC(u unsafe.Pointer) *CellAccessibleClass {
	c := (*C.GtkCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessibleClass{native: c}

	return g
}

func (recv *CellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAccessibleParentIface
*/
type CellAccessibleParentIface struct {
	native *C.GtkCellAccessibleParentIface
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
	c := (*C.GtkCellAccessibleParentIface)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessibleParentIface{native: c}

	return g
}

func (recv *CellAccessibleParentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAccessiblePrivate
*/
type CellAccessiblePrivate struct {
	native *C.GtkCellAccessiblePrivate
}

func CellAccessiblePrivateNewFromC(u unsafe.Pointer) *CellAccessiblePrivate {
	c := (*C.GtkCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessiblePrivate{native: c}

	return g
}

func (recv *CellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaBoxClass
*/
type CellAreaBoxClass struct {
	native *C.GtkCellAreaBoxClass
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellAreaBoxClassNewFromC(u unsafe.Pointer) *CellAreaBoxClass {
	c := (*C.GtkCellAreaBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaBoxClass{native: c}

	return g
}

func (recv *CellAreaBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaBoxPrivate
*/
type CellAreaBoxPrivate struct {
	native *C.GtkCellAreaBoxPrivate
}

func CellAreaBoxPrivateNewFromC(u unsafe.Pointer) *CellAreaBoxPrivate {
	c := (*C.GtkCellAreaBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaBoxPrivate{native: c}

	return g
}

func (recv *CellAreaBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaClass
*/
type CellAreaClass struct {
	native *C.GtkCellAreaClass
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
	c := (*C.GtkCellAreaClass)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaClass{native: c}

	return g
}

func (recv *CellAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaContextClass
*/
type CellAreaContextClass struct {
	native *C.GtkCellAreaContextClass
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
	c := (*C.GtkCellAreaContextClass)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaContextClass{native: c}

	return g
}

func (recv *CellAreaContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaContextPrivate
*/
type CellAreaContextPrivate struct {
	native *C.GtkCellAreaContextPrivate
}

func CellAreaContextPrivateNewFromC(u unsafe.Pointer) *CellAreaContextPrivate {
	c := (*C.GtkCellAreaContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaContextPrivate{native: c}

	return g
}

func (recv *CellAreaContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellAreaPrivate
*/
type CellAreaPrivate struct {
	native *C.GtkCellAreaPrivate
}

func CellAreaPrivateNewFromC(u unsafe.Pointer) *CellAreaPrivate {
	c := (*C.GtkCellAreaPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaPrivate{native: c}

	return g
}

func (recv *CellAreaPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellEditableIface
*/
type CellEditableIface struct {
	native *C.GtkCellEditableIface
	// Private : g_iface
	// no type for editing_done
	// no type for remove_widget
	// no type for start_editing
}

func CellEditableIfaceNewFromC(u unsafe.Pointer) *CellEditableIface {
	c := (*C.GtkCellEditableIface)(u)
	if c == nil {
		return nil
	}

	g := &CellEditableIface{native: c}

	return g
}

func (recv *CellEditableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellLayoutIface
*/
type CellLayoutIface struct {
	native *C.GtkCellLayoutIface
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
	c := (*C.GtkCellLayoutIface)(u)
	if c == nil {
		return nil
	}

	g := &CellLayoutIface{native: c}

	return g
}

func (recv *CellLayoutIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererAccelClass
*/
type CellRendererAccelClass struct {
	native *C.GtkCellRendererAccelClass
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
	c := (*C.GtkCellRendererAccelClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererAccelClass{native: c}

	return g
}

func (recv *CellRendererAccelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererAccelPrivate
*/
type CellRendererAccelPrivate struct {
	native *C.GtkCellRendererAccelPrivate
}

func CellRendererAccelPrivateNewFromC(u unsafe.Pointer) *CellRendererAccelPrivate {
	c := (*C.GtkCellRendererAccelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererAccelPrivate{native: c}

	return g
}

func (recv *CellRendererAccelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererClass
*/
type CellRendererClass struct {
	native *C.GtkCellRendererClass
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
	c := (*C.GtkCellRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererClass{native: c}

	return g
}

func (recv *CellRendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Sets the type to be used for creating accessibles for cells rendered by
// cell renderers of @renderer_class. Note that multiple accessibles will
// be created.
//
// This function should only be called from class init functions of cell
// renderers.
/*

C function

gtk_cell_renderer_class_set_accessible_type
*/
func (recv *CellRendererClass) SetAccessibleType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_cell_renderer_class_set_accessible_type((*C.GtkCellRendererClass)(recv.native), c_type)

	return
}

/*

C type

GtkCellRendererClassPrivate
*/
type CellRendererClassPrivate struct {
	native *C.GtkCellRendererClassPrivate
}

func CellRendererClassPrivateNewFromC(u unsafe.Pointer) *CellRendererClassPrivate {
	c := (*C.GtkCellRendererClassPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererClassPrivate{native: c}

	return g
}

func (recv *CellRendererClassPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererComboClass
*/
type CellRendererComboClass struct {
	native *C.GtkCellRendererComboClass
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererComboClassNewFromC(u unsafe.Pointer) *CellRendererComboClass {
	c := (*C.GtkCellRendererComboClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererComboClass{native: c}

	return g
}

func (recv *CellRendererComboClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererComboPrivate
*/
type CellRendererComboPrivate struct {
	native *C.GtkCellRendererComboPrivate
}

func CellRendererComboPrivateNewFromC(u unsafe.Pointer) *CellRendererComboPrivate {
	c := (*C.GtkCellRendererComboPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererComboPrivate{native: c}

	return g
}

func (recv *CellRendererComboPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererPixbufClass
*/
type CellRendererPixbufClass struct {
	native *C.GtkCellRendererPixbufClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererPixbufClassNewFromC(u unsafe.Pointer) *CellRendererPixbufClass {
	c := (*C.GtkCellRendererPixbufClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererPixbufClass{native: c}

	return g
}

func (recv *CellRendererPixbufClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererPixbufPrivate
*/
type CellRendererPixbufPrivate struct {
	native *C.GtkCellRendererPixbufPrivate
}

func CellRendererPixbufPrivateNewFromC(u unsafe.Pointer) *CellRendererPixbufPrivate {
	c := (*C.GtkCellRendererPixbufPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererPixbufPrivate{native: c}

	return g
}

func (recv *CellRendererPixbufPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererPrivate
*/
type CellRendererPrivate struct {
	native *C.GtkCellRendererPrivate
}

func CellRendererPrivateNewFromC(u unsafe.Pointer) *CellRendererPrivate {
	c := (*C.GtkCellRendererPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererPrivate{native: c}

	return g
}

func (recv *CellRendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererProgressClass
*/
type CellRendererProgressClass struct {
	native *C.GtkCellRendererProgressClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererProgressClassNewFromC(u unsafe.Pointer) *CellRendererProgressClass {
	c := (*C.GtkCellRendererProgressClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererProgressClass{native: c}

	return g
}

func (recv *CellRendererProgressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererProgressPrivate
*/
type CellRendererProgressPrivate struct {
	native *C.GtkCellRendererProgressPrivate
}

func CellRendererProgressPrivateNewFromC(u unsafe.Pointer) *CellRendererProgressPrivate {
	c := (*C.GtkCellRendererProgressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererProgressPrivate{native: c}

	return g
}

func (recv *CellRendererProgressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererSpinClass
*/
type CellRendererSpinClass struct {
	native *C.GtkCellRendererSpinClass
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererSpinClassNewFromC(u unsafe.Pointer) *CellRendererSpinClass {
	c := (*C.GtkCellRendererSpinClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinClass{native: c}

	return g
}

func (recv *CellRendererSpinClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererSpinPrivate
*/
type CellRendererSpinPrivate struct {
	native *C.GtkCellRendererSpinPrivate
}

func CellRendererSpinPrivateNewFromC(u unsafe.Pointer) *CellRendererSpinPrivate {
	c := (*C.GtkCellRendererSpinPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinPrivate{native: c}

	return g
}

func (recv *CellRendererSpinPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererSpinnerClass
*/
type CellRendererSpinnerClass struct {
	native *C.GtkCellRendererSpinnerClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererSpinnerClassNewFromC(u unsafe.Pointer) *CellRendererSpinnerClass {
	c := (*C.GtkCellRendererSpinnerClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinnerClass{native: c}

	return g
}

func (recv *CellRendererSpinnerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererSpinnerPrivate
*/
type CellRendererSpinnerPrivate struct {
	native *C.GtkCellRendererSpinnerPrivate
}

func CellRendererSpinnerPrivateNewFromC(u unsafe.Pointer) *CellRendererSpinnerPrivate {
	c := (*C.GtkCellRendererSpinnerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinnerPrivate{native: c}

	return g
}

func (recv *CellRendererSpinnerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererTextClass
*/
type CellRendererTextClass struct {
	native *C.GtkCellRendererTextClass
	// parent_class : record
	// no type for edited
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererTextClassNewFromC(u unsafe.Pointer) *CellRendererTextClass {
	c := (*C.GtkCellRendererTextClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererTextClass{native: c}

	return g
}

func (recv *CellRendererTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererTextPrivate
*/
type CellRendererTextPrivate struct {
	native *C.GtkCellRendererTextPrivate
}

func CellRendererTextPrivateNewFromC(u unsafe.Pointer) *CellRendererTextPrivate {
	c := (*C.GtkCellRendererTextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererTextPrivate{native: c}

	return g
}

func (recv *CellRendererTextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererToggleClass
*/
type CellRendererToggleClass struct {
	native *C.GtkCellRendererToggleClass
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellRendererToggleClassNewFromC(u unsafe.Pointer) *CellRendererToggleClass {
	c := (*C.GtkCellRendererToggleClass)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererToggleClass{native: c}

	return g
}

func (recv *CellRendererToggleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellRendererTogglePrivate
*/
type CellRendererTogglePrivate struct {
	native *C.GtkCellRendererTogglePrivate
}

func CellRendererTogglePrivateNewFromC(u unsafe.Pointer) *CellRendererTogglePrivate {
	c := (*C.GtkCellRendererTogglePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererTogglePrivate{native: c}

	return g
}

func (recv *CellRendererTogglePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellViewClass
*/
type CellViewClass struct {
	native *C.GtkCellViewClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CellViewClassNewFromC(u unsafe.Pointer) *CellViewClass {
	c := (*C.GtkCellViewClass)(u)
	if c == nil {
		return nil
	}

	g := &CellViewClass{native: c}

	return g
}

func (recv *CellViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCellViewPrivate
*/
type CellViewPrivate struct {
	native *C.GtkCellViewPrivate
}

func CellViewPrivateNewFromC(u unsafe.Pointer) *CellViewPrivate {
	c := (*C.GtkCellViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CellViewPrivate{native: c}

	return g
}

func (recv *CellViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCheckButtonClass
*/
type CheckButtonClass struct {
	native *C.GtkCheckButtonClass
	// parent_class : record
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CheckButtonClassNewFromC(u unsafe.Pointer) *CheckButtonClass {
	c := (*C.GtkCheckButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &CheckButtonClass{native: c}

	return g
}

func (recv *CheckButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCheckMenuItemAccessibleClass
*/
type CheckMenuItemAccessibleClass struct {
	native *C.GtkCheckMenuItemAccessibleClass
	// parent_class : record
}

func CheckMenuItemAccessibleClassNewFromC(u unsafe.Pointer) *CheckMenuItemAccessibleClass {
	c := (*C.GtkCheckMenuItemAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemAccessibleClass{native: c}

	return g
}

func (recv *CheckMenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCheckMenuItemAccessiblePrivate
*/
type CheckMenuItemAccessiblePrivate struct {
	native *C.GtkCheckMenuItemAccessiblePrivate
}

func CheckMenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *CheckMenuItemAccessiblePrivate {
	c := (*C.GtkCheckMenuItemAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemAccessiblePrivate{native: c}

	return g
}

func (recv *CheckMenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCheckMenuItemClass
*/
type CheckMenuItemClass struct {
	native *C.GtkCheckMenuItemClass
	// parent_class : record
	// no type for toggled
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CheckMenuItemClassNewFromC(u unsafe.Pointer) *CheckMenuItemClass {
	c := (*C.GtkCheckMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemClass{native: c}

	return g
}

func (recv *CheckMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCheckMenuItemPrivate
*/
type CheckMenuItemPrivate struct {
	native *C.GtkCheckMenuItemPrivate
}

func CheckMenuItemPrivateNewFromC(u unsafe.Pointer) *CheckMenuItemPrivate {
	c := (*C.GtkCheckMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemPrivate{native: c}

	return g
}

func (recv *CheckMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorButtonClass
*/
type ColorButtonClass struct {
	native *C.GtkColorButtonClass
	// parent_class : record
	// no type for color_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorButtonClassNewFromC(u unsafe.Pointer) *ColorButtonClass {
	c := (*C.GtkColorButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorButtonClass{native: c}

	return g
}

func (recv *ColorButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorButtonPrivate
*/
type ColorButtonPrivate struct {
	native *C.GtkColorButtonPrivate
}

func ColorButtonPrivateNewFromC(u unsafe.Pointer) *ColorButtonPrivate {
	c := (*C.GtkColorButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorButtonPrivate{native: c}

	return g
}

func (recv *ColorButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorChooserDialogClass
*/
type ColorChooserDialogClass struct {
	native *C.GtkColorChooserDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorChooserDialogClassNewFromC(u unsafe.Pointer) *ColorChooserDialogClass {
	c := (*C.GtkColorChooserDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserDialogClass{native: c}

	return g
}

func (recv *ColorChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorChooserDialogPrivate
*/
type ColorChooserDialogPrivate struct {
	native *C.GtkColorChooserDialogPrivate
}

func ColorChooserDialogPrivateNewFromC(u unsafe.Pointer) *ColorChooserDialogPrivate {
	c := (*C.GtkColorChooserDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserDialogPrivate{native: c}

	return g
}

func (recv *ColorChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorChooserInterface
*/
type ColorChooserInterface struct {
	native *C.GtkColorChooserInterface
	// base_interface : record
	// no type for get_rgba
	// no type for set_rgba
	// no type for add_palette
	// no type for color_activated
	// no type for padding
}

func ColorChooserInterfaceNewFromC(u unsafe.Pointer) *ColorChooserInterface {
	c := (*C.GtkColorChooserInterface)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserInterface{native: c}

	return g
}

func (recv *ColorChooserInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorChooserWidgetClass
*/
type ColorChooserWidgetClass struct {
	native *C.GtkColorChooserWidgetClass
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
	c := (*C.GtkColorChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserWidgetClass{native: c}

	return g
}

func (recv *ColorChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorChooserWidgetPrivate
*/
type ColorChooserWidgetPrivate struct {
	native *C.GtkColorChooserWidgetPrivate
}

func ColorChooserWidgetPrivateNewFromC(u unsafe.Pointer) *ColorChooserWidgetPrivate {
	c := (*C.GtkColorChooserWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserWidgetPrivate{native: c}

	return g
}

func (recv *ColorChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorSelectionClass
*/
type ColorSelectionClass struct {
	native *C.GtkColorSelectionClass
	// parent_class : record
	// no type for color_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorSelectionClassNewFromC(u unsafe.Pointer) *ColorSelectionClass {
	c := (*C.GtkColorSelectionClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionClass{native: c}

	return g
}

func (recv *ColorSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorSelectionDialogClass
*/
type ColorSelectionDialogClass struct {
	native *C.GtkColorSelectionDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ColorSelectionDialogClassNewFromC(u unsafe.Pointer) *ColorSelectionDialogClass {
	c := (*C.GtkColorSelectionDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionDialogClass{native: c}

	return g
}

func (recv *ColorSelectionDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorSelectionDialogPrivate
*/
type ColorSelectionDialogPrivate struct {
	native *C.GtkColorSelectionDialogPrivate
}

func ColorSelectionDialogPrivateNewFromC(u unsafe.Pointer) *ColorSelectionDialogPrivate {
	c := (*C.GtkColorSelectionDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionDialogPrivate{native: c}

	return g
}

func (recv *ColorSelectionDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkColorSelectionPrivate
*/
type ColorSelectionPrivate struct {
	native *C.GtkColorSelectionPrivate
}

func ColorSelectionPrivateNewFromC(u unsafe.Pointer) *ColorSelectionPrivate {
	c := (*C.GtkColorSelectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionPrivate{native: c}

	return g
}

func (recv *ColorSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxAccessibleClass
*/
type ComboBoxAccessibleClass struct {
	native *C.GtkComboBoxAccessibleClass
	// parent_class : record
}

func ComboBoxAccessibleClassNewFromC(u unsafe.Pointer) *ComboBoxAccessibleClass {
	c := (*C.GtkComboBoxAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxAccessibleClass{native: c}

	return g
}

func (recv *ComboBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxAccessiblePrivate
*/
type ComboBoxAccessiblePrivate struct {
	native *C.GtkComboBoxAccessiblePrivate
}

func ComboBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *ComboBoxAccessiblePrivate {
	c := (*C.GtkComboBoxAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxAccessiblePrivate{native: c}

	return g
}

func (recv *ComboBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxClass
*/
type ComboBoxClass struct {
	native *C.GtkComboBoxClass
	// parent_class : record
	// no type for changed
	// no type for format_entry_text
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

func ComboBoxClassNewFromC(u unsafe.Pointer) *ComboBoxClass {
	c := (*C.GtkComboBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxClass{native: c}

	return g
}

func (recv *ComboBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxPrivate
*/
type ComboBoxPrivate struct {
	native *C.GtkComboBoxPrivate
}

func ComboBoxPrivateNewFromC(u unsafe.Pointer) *ComboBoxPrivate {
	c := (*C.GtkComboBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxPrivate{native: c}

	return g
}

func (recv *ComboBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxTextClass
*/
type ComboBoxTextClass struct {
	native *C.GtkComboBoxTextClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ComboBoxTextClassNewFromC(u unsafe.Pointer) *ComboBoxTextClass {
	c := (*C.GtkComboBoxTextClass)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxTextClass{native: c}

	return g
}

func (recv *ComboBoxTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkComboBoxTextPrivate
*/
type ComboBoxTextPrivate struct {
	native *C.GtkComboBoxTextPrivate
}

func ComboBoxTextPrivateNewFromC(u unsafe.Pointer) *ComboBoxTextPrivate {
	c := (*C.GtkComboBoxTextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxTextPrivate{native: c}

	return g
}

func (recv *ComboBoxTextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkContainerAccessibleClass
*/
type ContainerAccessibleClass struct {
	native *C.GtkContainerAccessibleClass
	// parent_class : record
	// no type for add_gtk
	// no type for remove_gtk
}

func ContainerAccessibleClassNewFromC(u unsafe.Pointer) *ContainerAccessibleClass {
	c := (*C.GtkContainerAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ContainerAccessibleClass{native: c}

	return g
}

func (recv *ContainerAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkContainerAccessiblePrivate
*/
type ContainerAccessiblePrivate struct {
	native *C.GtkContainerAccessiblePrivate
}

func ContainerAccessiblePrivateNewFromC(u unsafe.Pointer) *ContainerAccessiblePrivate {
	c := (*C.GtkContainerAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContainerAccessiblePrivate{native: c}

	return g
}

func (recv *ContainerAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkContainerCellAccessibleClass
*/
type ContainerCellAccessibleClass struct {
	native *C.GtkContainerCellAccessibleClass
	// parent_class : record
}

func ContainerCellAccessibleClassNewFromC(u unsafe.Pointer) *ContainerCellAccessibleClass {
	c := (*C.GtkContainerCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ContainerCellAccessibleClass{native: c}

	return g
}

func (recv *ContainerCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkContainerCellAccessiblePrivate
*/
type ContainerCellAccessiblePrivate struct {
	native *C.GtkContainerCellAccessiblePrivate
}

func ContainerCellAccessiblePrivateNewFromC(u unsafe.Pointer) *ContainerCellAccessiblePrivate {
	c := (*C.GtkContainerCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContainerCellAccessiblePrivate{native: c}

	return g
}

func (recv *ContainerCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Base class for containers.
/*

C type

GtkContainerClass
*/
type ContainerClass struct {
	native *C.GtkContainerClass
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
	c := (*C.GtkContainerClass)(u)
	if c == nil {
		return nil
	}

	g := &ContainerClass{native: c}

	return g
}

func (recv *ContainerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_container_class_find_child_property : return type : Blacklisted record : GParamSpec

// Modifies a subclass of #GtkContainerClass to automatically add and
// remove the border-width setting on GtkContainer.  This allows the
// subclass to ignore the border width in its size request and
// allocate methods. The intent is for a subclass to invoke this
// in its class_init function.
//
// gtk_container_class_handle_border_width() is necessary because it
// would break API too badly to make this behavior the default. So
// subclasses must “opt in” to the parent class handling border_width
// for them.
/*

C function

gtk_container_class_handle_border_width
*/
func (recv *ContainerClass) HandleBorderWidth() {
	C.gtk_container_class_handle_border_width((*C.GtkContainerClass)(recv.native))

	return
}

// Unsupported : gtk_container_class_install_child_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_container_class_list_child_properties : no return type

/*

C type

GtkContainerPrivate
*/
type ContainerPrivate struct {
	native *C.GtkContainerPrivate
}

func ContainerPrivateNewFromC(u unsafe.Pointer) *ContainerPrivate {
	c := (*C.GtkContainerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContainerPrivate{native: c}

	return g
}

func (recv *ContainerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCssProviderClass
*/
type CssProviderClass struct {
	native *C.GtkCssProviderClass
	// parent_class : record
	// no type for parsing_error
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func CssProviderClassNewFromC(u unsafe.Pointer) *CssProviderClass {
	c := (*C.GtkCssProviderClass)(u)
	if c == nil {
		return nil
	}

	g := &CssProviderClass{native: c}

	return g
}

func (recv *CssProviderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkCssProviderPrivate
*/
type CssProviderPrivate struct {
	native *C.GtkCssProviderPrivate
}

func CssProviderPrivateNewFromC(u unsafe.Pointer) *CssProviderPrivate {
	c := (*C.GtkCssProviderPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CssProviderPrivate{native: c}

	return g
}

func (recv *CssProviderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkDialogClass
*/
type DialogClass struct {
	native *C.GtkDialogClass
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func DialogClassNewFromC(u unsafe.Pointer) *DialogClass {
	c := (*C.GtkDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &DialogClass{native: c}

	return g
}

func (recv *DialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkDialogPrivate
*/
type DialogPrivate struct {
	native *C.GtkDialogPrivate
}

func DialogPrivateNewFromC(u unsafe.Pointer) *DialogPrivate {
	c := (*C.GtkDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DialogPrivate{native: c}

	return g
}

func (recv *DialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkDrawingAreaClass
*/
type DrawingAreaClass struct {
	native *C.GtkDrawingAreaClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func DrawingAreaClassNewFromC(u unsafe.Pointer) *DrawingAreaClass {
	c := (*C.GtkDrawingAreaClass)(u)
	if c == nil {
		return nil
	}

	g := &DrawingAreaClass{native: c}

	return g
}

func (recv *DrawingAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEditableInterface
*/
type EditableInterface struct {
	native *C.GtkEditableInterface
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
	c := (*C.GtkEditableInterface)(u)
	if c == nil {
		return nil
	}

	g := &EditableInterface{native: c}

	return g
}

func (recv *EditableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryAccessibleClass
*/
type EntryAccessibleClass struct {
	native *C.GtkEntryAccessibleClass
	// parent_class : record
}

func EntryAccessibleClassNewFromC(u unsafe.Pointer) *EntryAccessibleClass {
	c := (*C.GtkEntryAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &EntryAccessibleClass{native: c}

	return g
}

func (recv *EntryAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryAccessiblePrivate
*/
type EntryAccessiblePrivate struct {
	native *C.GtkEntryAccessiblePrivate
}

func EntryAccessiblePrivateNewFromC(u unsafe.Pointer) *EntryAccessiblePrivate {
	c := (*C.GtkEntryAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &EntryAccessiblePrivate{native: c}

	return g
}

func (recv *EntryAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryBufferClass
*/
type EntryBufferClass struct {
	native *C.GtkEntryBufferClass
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
	c := (*C.GtkEntryBufferClass)(u)
	if c == nil {
		return nil
	}

	g := &EntryBufferClass{native: c}

	return g
}

func (recv *EntryBufferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryBufferPrivate
*/
type EntryBufferPrivate struct {
	native *C.GtkEntryBufferPrivate
}

func EntryBufferPrivateNewFromC(u unsafe.Pointer) *EntryBufferPrivate {
	c := (*C.GtkEntryBufferPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EntryBufferPrivate{native: c}

	return g
}

func (recv *EntryBufferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Class structure for #GtkEntry. All virtual functions have a default
// implementation. Derived classes may set the virtual function pointers for the
// signal handlers to %NULL, but must keep @get_text_area_size and
// @get_frame_size non-%NULL; either use the default implementation, or provide
// a custom one.
/*

C type

GtkEntryClass
*/
type EntryClass struct {
	native *C.GtkEntryClass
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
	c := (*C.GtkEntryClass)(u)
	if c == nil {
		return nil
	}

	g := &EntryClass{native: c}

	return g
}

func (recv *EntryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryCompletionClass
*/
type EntryCompletionClass struct {
	native *C.GtkEntryCompletionClass
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
	c := (*C.GtkEntryCompletionClass)(u)
	if c == nil {
		return nil
	}

	g := &EntryCompletionClass{native: c}

	return g
}

func (recv *EntryCompletionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryCompletionPrivate
*/
type EntryCompletionPrivate struct {
	native *C.GtkEntryCompletionPrivate
}

func EntryCompletionPrivateNewFromC(u unsafe.Pointer) *EntryCompletionPrivate {
	c := (*C.GtkEntryCompletionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EntryCompletionPrivate{native: c}

	return g
}

func (recv *EntryCompletionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEntryPrivate
*/
type EntryPrivate struct {
	native *C.GtkEntryPrivate
}

func EntryPrivateNewFromC(u unsafe.Pointer) *EntryPrivate {
	c := (*C.GtkEntryPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EntryPrivate{native: c}

	return g
}

func (recv *EntryPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEventBoxClass
*/
type EventBoxClass struct {
	native *C.GtkEventBoxClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func EventBoxClassNewFromC(u unsafe.Pointer) *EventBoxClass {
	c := (*C.GtkEventBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &EventBoxClass{native: c}

	return g
}

func (recv *EventBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEventBoxPrivate
*/
type EventBoxPrivate struct {
	native *C.GtkEventBoxPrivate
}

func EventBoxPrivateNewFromC(u unsafe.Pointer) *EventBoxPrivate {
	c := (*C.GtkEventBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EventBoxPrivate{native: c}

	return g
}

func (recv *EventBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkEventControllerClass
*/
type EventControllerClass struct {
	native *C.GtkEventControllerClass
}

func EventControllerClassNewFromC(u unsafe.Pointer) *EventControllerClass {
	c := (*C.GtkEventControllerClass)(u)
	if c == nil {
		return nil
	}

	g := &EventControllerClass{native: c}

	return g
}

func (recv *EventControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkExpanderAccessibleClass
*/
type ExpanderAccessibleClass struct {
	native *C.GtkExpanderAccessibleClass
	// parent_class : record
}

func ExpanderAccessibleClassNewFromC(u unsafe.Pointer) *ExpanderAccessibleClass {
	c := (*C.GtkExpanderAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderAccessibleClass{native: c}

	return g
}

func (recv *ExpanderAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkExpanderAccessiblePrivate
*/
type ExpanderAccessiblePrivate struct {
	native *C.GtkExpanderAccessiblePrivate
}

func ExpanderAccessiblePrivateNewFromC(u unsafe.Pointer) *ExpanderAccessiblePrivate {
	c := (*C.GtkExpanderAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderAccessiblePrivate{native: c}

	return g
}

func (recv *ExpanderAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkExpanderClass
*/
type ExpanderClass struct {
	native *C.GtkExpanderClass
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ExpanderClassNewFromC(u unsafe.Pointer) *ExpanderClass {
	c := (*C.GtkExpanderClass)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderClass{native: c}

	return g
}

func (recv *ExpanderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkExpanderPrivate
*/
type ExpanderPrivate struct {
	native *C.GtkExpanderPrivate
}

func ExpanderPrivateNewFromC(u unsafe.Pointer) *ExpanderPrivate {
	c := (*C.GtkExpanderPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderPrivate{native: c}

	return g
}

func (recv *ExpanderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserButtonClass
*/
type FileChooserButtonClass struct {
	native *C.GtkFileChooserButtonClass
	// parent_class : record
	// no type for file_set
	// Private : __gtk_reserved1
	// Private : __gtk_reserved2
	// Private : __gtk_reserved3
	// Private : __gtk_reserved4
}

func FileChooserButtonClassNewFromC(u unsafe.Pointer) *FileChooserButtonClass {
	c := (*C.GtkFileChooserButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserButtonClass{native: c}

	return g
}

func (recv *FileChooserButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserButtonPrivate
*/
type FileChooserButtonPrivate struct {
	native *C.GtkFileChooserButtonPrivate
}

func FileChooserButtonPrivateNewFromC(u unsafe.Pointer) *FileChooserButtonPrivate {
	c := (*C.GtkFileChooserButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserButtonPrivate{native: c}

	return g
}

func (recv *FileChooserButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserDialogClass
*/
type FileChooserDialogClass struct {
	native *C.GtkFileChooserDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FileChooserDialogClassNewFromC(u unsafe.Pointer) *FileChooserDialogClass {
	c := (*C.GtkFileChooserDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserDialogClass{native: c}

	return g
}

func (recv *FileChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserDialogPrivate
*/
type FileChooserDialogPrivate struct {
	native *C.GtkFileChooserDialogPrivate
}

func FileChooserDialogPrivateNewFromC(u unsafe.Pointer) *FileChooserDialogPrivate {
	c := (*C.GtkFileChooserDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserDialogPrivate{native: c}

	return g
}

func (recv *FileChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserWidgetClass
*/
type FileChooserWidgetClass struct {
	native *C.GtkFileChooserWidgetClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FileChooserWidgetClassNewFromC(u unsafe.Pointer) *FileChooserWidgetClass {
	c := (*C.GtkFileChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserWidgetClass{native: c}

	return g
}

func (recv *FileChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFileChooserWidgetPrivate
*/
type FileChooserWidgetPrivate struct {
	native *C.GtkFileChooserWidgetPrivate
}

func FileChooserWidgetPrivateNewFromC(u unsafe.Pointer) *FileChooserWidgetPrivate {
	c := (*C.GtkFileChooserWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserWidgetPrivate{native: c}

	return g
}

func (recv *FileChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A #GtkFileFilterInfo-struct is used to pass information about the
// tested file to gtk_file_filter_filter().
/*

C type

GtkFileFilterInfo
*/
type FileFilterInfo struct {
	native      *C.GtkFileFilterInfo
	Contains    FileFilterFlags
	Filename    string
	Uri         string
	DisplayName string
	MimeType    string
}

func FileFilterInfoNewFromC(u unsafe.Pointer) *FileFilterInfo {
	c := (*C.GtkFileFilterInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileFilterInfo{
		Contains:    (FileFilterFlags)(c.contains),
		DisplayName: C.GoString(c.display_name),
		Filename:    C.GoString(c.filename),
		MimeType:    C.GoString(c.mime_type),
		Uri:         C.GoString(c.uri),
		native:      c,
	}

	return g
}

func (recv *FileFilterInfo) ToC() unsafe.Pointer {
	recv.native.contains =
		(C.GtkFileFilterFlags)(recv.Contains)
	recv.native.filename =
		C.CString(recv.Filename)
	recv.native.uri =
		C.CString(recv.Uri)
	recv.native.display_name =
		C.CString(recv.DisplayName)
	recv.native.mime_type =
		C.CString(recv.MimeType)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFixedChild
*/
type FixedChild struct {
	native *C.GtkFixedChild
	// widget : record
	X int32
	Y int32
}

func FixedChildNewFromC(u unsafe.Pointer) *FixedChild {
	c := (*C.GtkFixedChild)(u)
	if c == nil {
		return nil
	}

	g := &FixedChild{
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *FixedChild) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFixedClass
*/
type FixedClass struct {
	native *C.GtkFixedClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FixedClassNewFromC(u unsafe.Pointer) *FixedClass {
	c := (*C.GtkFixedClass)(u)
	if c == nil {
		return nil
	}

	g := &FixedClass{native: c}

	return g
}

func (recv *FixedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFixedPrivate
*/
type FixedPrivate struct {
	native *C.GtkFixedPrivate
}

func FixedPrivateNewFromC(u unsafe.Pointer) *FixedPrivate {
	c := (*C.GtkFixedPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FixedPrivate{native: c}

	return g
}

func (recv *FixedPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFlowBoxAccessibleClass
*/
type FlowBoxAccessibleClass struct {
	native *C.GtkFlowBoxAccessibleClass
	// parent_class : record
}

func FlowBoxAccessibleClassNewFromC(u unsafe.Pointer) *FlowBoxAccessibleClass {
	c := (*C.GtkFlowBoxAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxAccessibleClass{native: c}

	return g
}

func (recv *FlowBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFlowBoxAccessiblePrivate
*/
type FlowBoxAccessiblePrivate struct {
	native *C.GtkFlowBoxAccessiblePrivate
}

func FlowBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *FlowBoxAccessiblePrivate {
	c := (*C.GtkFlowBoxAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxAccessiblePrivate{native: c}

	return g
}

func (recv *FlowBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFlowBoxChildAccessibleClass
*/
type FlowBoxChildAccessibleClass struct {
	native *C.GtkFlowBoxChildAccessibleClass
	// parent_class : record
}

func FlowBoxChildAccessibleClassNewFromC(u unsafe.Pointer) *FlowBoxChildAccessibleClass {
	c := (*C.GtkFlowBoxChildAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChildAccessibleClass{native: c}

	return g
}

func (recv *FlowBoxChildAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFlowBoxChildClass
*/
type FlowBoxChildClass struct {
	native *C.GtkFlowBoxChildClass
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

func FlowBoxChildClassNewFromC(u unsafe.Pointer) *FlowBoxChildClass {
	c := (*C.GtkFlowBoxChildClass)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChildClass{native: c}

	return g
}

func (recv *FlowBoxChildClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFlowBoxClass
*/
type FlowBoxClass struct {
	native *C.GtkFlowBoxClass
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
	c := (*C.GtkFlowBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxClass{native: c}

	return g
}

func (recv *FlowBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontButtonClass
*/
type FontButtonClass struct {
	native *C.GtkFontButtonClass
	// parent_class : record
	// no type for font_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontButtonClassNewFromC(u unsafe.Pointer) *FontButtonClass {
	c := (*C.GtkFontButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &FontButtonClass{native: c}

	return g
}

func (recv *FontButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontButtonPrivate
*/
type FontButtonPrivate struct {
	native *C.GtkFontButtonPrivate
}

func FontButtonPrivateNewFromC(u unsafe.Pointer) *FontButtonPrivate {
	c := (*C.GtkFontButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FontButtonPrivate{native: c}

	return g
}

func (recv *FontButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontChooserDialogClass
*/
type FontChooserDialogClass struct {
	native *C.GtkFontChooserDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontChooserDialogClassNewFromC(u unsafe.Pointer) *FontChooserDialogClass {
	c := (*C.GtkFontChooserDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserDialogClass{native: c}

	return g
}

func (recv *FontChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontChooserDialogPrivate
*/
type FontChooserDialogPrivate struct {
	native *C.GtkFontChooserDialogPrivate
}

func FontChooserDialogPrivateNewFromC(u unsafe.Pointer) *FontChooserDialogPrivate {
	c := (*C.GtkFontChooserDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserDialogPrivate{native: c}

	return g
}

func (recv *FontChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontChooserIface
*/
type FontChooserIface struct {
	native *C.GtkFontChooserIface
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
	c := (*C.GtkFontChooserIface)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserIface{native: c}

	return g
}

func (recv *FontChooserIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontChooserWidgetClass
*/
type FontChooserWidgetClass struct {
	native *C.GtkFontChooserWidgetClass
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
	c := (*C.GtkFontChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserWidgetClass{native: c}

	return g
}

func (recv *FontChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontChooserWidgetPrivate
*/
type FontChooserWidgetPrivate struct {
	native *C.GtkFontChooserWidgetPrivate
}

func FontChooserWidgetPrivateNewFromC(u unsafe.Pointer) *FontChooserWidgetPrivate {
	c := (*C.GtkFontChooserWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserWidgetPrivate{native: c}

	return g
}

func (recv *FontChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontSelectionClass
*/
type FontSelectionClass struct {
	native *C.GtkFontSelectionClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontSelectionClassNewFromC(u unsafe.Pointer) *FontSelectionClass {
	c := (*C.GtkFontSelectionClass)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionClass{native: c}

	return g
}

func (recv *FontSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontSelectionDialogClass
*/
type FontSelectionDialogClass struct {
	native *C.GtkFontSelectionDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FontSelectionDialogClassNewFromC(u unsafe.Pointer) *FontSelectionDialogClass {
	c := (*C.GtkFontSelectionDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionDialogClass{native: c}

	return g
}

func (recv *FontSelectionDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontSelectionDialogPrivate
*/
type FontSelectionDialogPrivate struct {
	native *C.GtkFontSelectionDialogPrivate
}

func FontSelectionDialogPrivateNewFromC(u unsafe.Pointer) *FontSelectionDialogPrivate {
	c := (*C.GtkFontSelectionDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionDialogPrivate{native: c}

	return g
}

func (recv *FontSelectionDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFontSelectionPrivate
*/
type FontSelectionPrivate struct {
	native *C.GtkFontSelectionPrivate
}

func FontSelectionPrivateNewFromC(u unsafe.Pointer) *FontSelectionPrivate {
	c := (*C.GtkFontSelectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionPrivate{native: c}

	return g
}

func (recv *FontSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFrameAccessibleClass
*/
type FrameAccessibleClass struct {
	native *C.GtkFrameAccessibleClass
	// parent_class : record
}

func FrameAccessibleClassNewFromC(u unsafe.Pointer) *FrameAccessibleClass {
	c := (*C.GtkFrameAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &FrameAccessibleClass{native: c}

	return g
}

func (recv *FrameAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFrameAccessiblePrivate
*/
type FrameAccessiblePrivate struct {
	native *C.GtkFrameAccessiblePrivate
}

func FrameAccessiblePrivateNewFromC(u unsafe.Pointer) *FrameAccessiblePrivate {
	c := (*C.GtkFrameAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &FrameAccessiblePrivate{native: c}

	return g
}

func (recv *FrameAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFrameClass
*/
type FrameClass struct {
	native *C.GtkFrameClass
	// parent_class : record
	// no type for compute_child_allocation
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func FrameClassNewFromC(u unsafe.Pointer) *FrameClass {
	c := (*C.GtkFrameClass)(u)
	if c == nil {
		return nil
	}

	g := &FrameClass{native: c}

	return g
}

func (recv *FrameClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkFramePrivate
*/
type FramePrivate struct {
	native *C.GtkFramePrivate
}

func FramePrivateNewFromC(u unsafe.Pointer) *FramePrivate {
	c := (*C.GtkFramePrivate)(u)
	if c == nil {
		return nil
	}

	g := &FramePrivate{native: c}

	return g
}

func (recv *FramePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureClass
*/
type GestureClass struct {
	native *C.GtkGestureClass
}

func GestureClassNewFromC(u unsafe.Pointer) *GestureClass {
	c := (*C.GtkGestureClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureClass{native: c}

	return g
}

func (recv *GestureClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureDragClass
*/
type GestureDragClass struct {
	native *C.GtkGestureDragClass
}

func GestureDragClassNewFromC(u unsafe.Pointer) *GestureDragClass {
	c := (*C.GtkGestureDragClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureDragClass{native: c}

	return g
}

func (recv *GestureDragClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureLongPressClass
*/
type GestureLongPressClass struct {
	native *C.GtkGestureLongPressClass
}

func GestureLongPressClassNewFromC(u unsafe.Pointer) *GestureLongPressClass {
	c := (*C.GtkGestureLongPressClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureLongPressClass{native: c}

	return g
}

func (recv *GestureLongPressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureMultiPressClass
*/
type GestureMultiPressClass struct {
	native *C.GtkGestureMultiPressClass
}

func GestureMultiPressClassNewFromC(u unsafe.Pointer) *GestureMultiPressClass {
	c := (*C.GtkGestureMultiPressClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureMultiPressClass{native: c}

	return g
}

func (recv *GestureMultiPressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGesturePanClass
*/
type GesturePanClass struct {
	native *C.GtkGesturePanClass
}

func GesturePanClassNewFromC(u unsafe.Pointer) *GesturePanClass {
	c := (*C.GtkGesturePanClass)(u)
	if c == nil {
		return nil
	}

	g := &GesturePanClass{native: c}

	return g
}

func (recv *GesturePanClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureRotateClass
*/
type GestureRotateClass struct {
	native *C.GtkGestureRotateClass
}

func GestureRotateClassNewFromC(u unsafe.Pointer) *GestureRotateClass {
	c := (*C.GtkGestureRotateClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureRotateClass{native: c}

	return g
}

func (recv *GestureRotateClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureSingleClass
*/
type GestureSingleClass struct {
	native *C.GtkGestureSingleClass
}

func GestureSingleClassNewFromC(u unsafe.Pointer) *GestureSingleClass {
	c := (*C.GtkGestureSingleClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureSingleClass{native: c}

	return g
}

func (recv *GestureSingleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureSwipeClass
*/
type GestureSwipeClass struct {
	native *C.GtkGestureSwipeClass
}

func GestureSwipeClassNewFromC(u unsafe.Pointer) *GestureSwipeClass {
	c := (*C.GtkGestureSwipeClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureSwipeClass{native: c}

	return g
}

func (recv *GestureSwipeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGestureZoomClass
*/
type GestureZoomClass struct {
	native *C.GtkGestureZoomClass
}

func GestureZoomClassNewFromC(u unsafe.Pointer) *GestureZoomClass {
	c := (*C.GtkGestureZoomClass)(u)
	if c == nil {
		return nil
	}

	g := &GestureZoomClass{native: c}

	return g
}

func (recv *GestureZoomClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkGradient is a boxed type that represents a gradient.
// It is the result of parsing a
// [gradient expression][gtkcssprovider-gradients].
// To obtain the gradient represented by a GtkGradient, it has to
// be resolved with gtk_gradient_resolve(), which replaces all
// symbolic color references by the colors they refer to (in a given
// context) and constructs a #cairo_pattern_t value.
//
// It is not normally necessary to deal directly with #GtkGradients,
// since they are mostly used behind the scenes by #GtkStyleContext and
// #GtkCssProvider.
//
// #GtkGradient is deprecated. It was used internally by GTK’s CSS engine
// to represent gradients. As its handling is not conforming to modern
// web standards, it is not used anymore. If you want to use gradients in
// your own code, please use Cairo directly.
/*

C type

GtkGradient
*/
type Gradient struct {
	native *C.GtkGradient
}

func GradientNewFromC(u unsafe.Pointer) *Gradient {
	c := (*C.GtkGradient)(u)
	if c == nil {
		return nil
	}

	g := &Gradient{native: c}

	return g
}

func (recv *Gradient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C function

gtk_gradient_resolve_for_context
*/
func (recv *Gradient) ResolveForContext(context *StyleContext) *cairo.Pattern {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	retC := C.gtk_gradient_resolve_for_context((*C.GtkGradient)(recv.native), c_context)
	retGo := cairo.PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a string representation for @gradient that is suitable
// for using in GTK CSS files.
/*

C function

gtk_gradient_to_string
*/
func (recv *Gradient) ToString() string {
	retC := C.gtk_gradient_to_string((*C.GtkGradient)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

/*

C type

GtkGridClass
*/
type GridClass struct {
	native *C.GtkGridClass
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
	c := (*C.GtkGridClass)(u)
	if c == nil {
		return nil
	}

	g := &GridClass{native: c}

	return g
}

func (recv *GridClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkGridPrivate
*/
type GridPrivate struct {
	native *C.GtkGridPrivate
}

func GridPrivateNewFromC(u unsafe.Pointer) *GridPrivate {
	c := (*C.GtkGridPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GridPrivate{native: c}

	return g
}

func (recv *GridPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHBoxClass
*/
type HBoxClass struct {
	native *C.GtkHBoxClass
	// parent_class : record
}

func HBoxClassNewFromC(u unsafe.Pointer) *HBoxClass {
	c := (*C.GtkHBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &HBoxClass{native: c}

	return g
}

func (recv *HBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHButtonBoxClass
*/
type HButtonBoxClass struct {
	native *C.GtkHButtonBoxClass
	// parent_class : record
}

func HButtonBoxClassNewFromC(u unsafe.Pointer) *HButtonBoxClass {
	c := (*C.GtkHButtonBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &HButtonBoxClass{native: c}

	return g
}

func (recv *HButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHPanedClass
*/
type HPanedClass struct {
	native *C.GtkHPanedClass
	// parent_class : record
}

func HPanedClassNewFromC(u unsafe.Pointer) *HPanedClass {
	c := (*C.GtkHPanedClass)(u)
	if c == nil {
		return nil
	}

	g := &HPanedClass{native: c}

	return g
}

func (recv *HPanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHSVClass
*/
type HSVClass struct {
	native *C.GtkHSVClass
	// parent_class : record
	// no type for changed
	// no type for move
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HSVClassNewFromC(u unsafe.Pointer) *HSVClass {
	c := (*C.GtkHSVClass)(u)
	if c == nil {
		return nil
	}

	g := &HSVClass{native: c}

	return g
}

func (recv *HSVClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHSVPrivate
*/
type HSVPrivate struct {
	native *C.GtkHSVPrivate
}

func HSVPrivateNewFromC(u unsafe.Pointer) *HSVPrivate {
	c := (*C.GtkHSVPrivate)(u)
	if c == nil {
		return nil
	}

	g := &HSVPrivate{native: c}

	return g
}

func (recv *HSVPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHScaleClass
*/
type HScaleClass struct {
	native *C.GtkHScaleClass
	// parent_class : record
}

func HScaleClassNewFromC(u unsafe.Pointer) *HScaleClass {
	c := (*C.GtkHScaleClass)(u)
	if c == nil {
		return nil
	}

	g := &HScaleClass{native: c}

	return g
}

func (recv *HScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHScrollbarClass
*/
type HScrollbarClass struct {
	native *C.GtkHScrollbarClass
	// parent_class : record
}

func HScrollbarClassNewFromC(u unsafe.Pointer) *HScrollbarClass {
	c := (*C.GtkHScrollbarClass)(u)
	if c == nil {
		return nil
	}

	g := &HScrollbarClass{native: c}

	return g
}

func (recv *HScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHSeparatorClass
*/
type HSeparatorClass struct {
	native *C.GtkHSeparatorClass
	// parent_class : record
}

func HSeparatorClassNewFromC(u unsafe.Pointer) *HSeparatorClass {
	c := (*C.GtkHSeparatorClass)(u)
	if c == nil {
		return nil
	}

	g := &HSeparatorClass{native: c}

	return g
}

func (recv *HSeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHandleBoxClass
*/
type HandleBoxClass struct {
	native *C.GtkHandleBoxClass
	// parent_class : record
	// no type for child_attached
	// no type for child_detached
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HandleBoxClassNewFromC(u unsafe.Pointer) *HandleBoxClass {
	c := (*C.GtkHandleBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &HandleBoxClass{native: c}

	return g
}

func (recv *HandleBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHandleBoxPrivate
*/
type HandleBoxPrivate struct {
	native *C.GtkHandleBoxPrivate
}

func HandleBoxPrivateNewFromC(u unsafe.Pointer) *HandleBoxPrivate {
	c := (*C.GtkHandleBoxPrivate)(u)
	if c == nil {
		return nil
	}

	g := &HandleBoxPrivate{native: c}

	return g
}

func (recv *HandleBoxPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHeaderBarClass
*/
type HeaderBarClass struct {
	native *C.GtkHeaderBarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func HeaderBarClassNewFromC(u unsafe.Pointer) *HeaderBarClass {
	c := (*C.GtkHeaderBarClass)(u)
	if c == nil {
		return nil
	}

	g := &HeaderBarClass{native: c}

	return g
}

func (recv *HeaderBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkHeaderBarPrivate
*/
type HeaderBarPrivate struct {
	native *C.GtkHeaderBarPrivate
}

func HeaderBarPrivateNewFromC(u unsafe.Pointer) *HeaderBarPrivate {
	c := (*C.GtkHeaderBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &HeaderBarPrivate{native: c}

	return g
}

func (recv *HeaderBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIMContextClass
*/
type IMContextClass struct {
	native *C.GtkIMContextClass
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
	c := (*C.GtkIMContextClass)(u)
	if c == nil {
		return nil
	}

	g := &IMContextClass{native: c}

	return g
}

func (recv *IMContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bookkeeping information about a loadable input method.
/*

C type

GtkIMContextInfo
*/
type IMContextInfo struct {
	native         *C.GtkIMContextInfo
	ContextId      string
	ContextName    string
	Domain         string
	DomainDirname  string
	DefaultLocales string
}

func IMContextInfoNewFromC(u unsafe.Pointer) *IMContextInfo {
	c := (*C.GtkIMContextInfo)(u)
	if c == nil {
		return nil
	}

	g := &IMContextInfo{
		ContextId:      C.GoString(c.context_id),
		ContextName:    C.GoString(c.context_name),
		DefaultLocales: C.GoString(c.default_locales),
		Domain:         C.GoString(c.domain),
		DomainDirname:  C.GoString(c.domain_dirname),
		native:         c,
	}

	return g
}

func (recv *IMContextInfo) ToC() unsafe.Pointer {
	recv.native.context_id =
		C.CString(recv.ContextId)
	recv.native.context_name =
		C.CString(recv.ContextName)
	recv.native.domain =
		C.CString(recv.Domain)
	recv.native.domain_dirname =
		C.CString(recv.DomainDirname)
	recv.native.default_locales =
		C.CString(recv.DefaultLocales)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIMContextSimpleClass
*/
type IMContextSimpleClass struct {
	native *C.GtkIMContextSimpleClass
	// parent_class : record
}

func IMContextSimpleClassNewFromC(u unsafe.Pointer) *IMContextSimpleClass {
	c := (*C.GtkIMContextSimpleClass)(u)
	if c == nil {
		return nil
	}

	g := &IMContextSimpleClass{native: c}

	return g
}

func (recv *IMContextSimpleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIMContextSimplePrivate
*/
type IMContextSimplePrivate struct {
	native *C.GtkIMContextSimplePrivate
}

func IMContextSimplePrivateNewFromC(u unsafe.Pointer) *IMContextSimplePrivate {
	c := (*C.GtkIMContextSimplePrivate)(u)
	if c == nil {
		return nil
	}

	g := &IMContextSimplePrivate{native: c}

	return g
}

func (recv *IMContextSimplePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIMMulticontextClass
*/
type IMMulticontextClass struct {
	native *C.GtkIMMulticontextClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IMMulticontextClassNewFromC(u unsafe.Pointer) *IMMulticontextClass {
	c := (*C.GtkIMMulticontextClass)(u)
	if c == nil {
		return nil
	}

	g := &IMMulticontextClass{native: c}

	return g
}

func (recv *IMMulticontextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIMMulticontextPrivate
*/
type IMMulticontextPrivate struct {
	native *C.GtkIMMulticontextPrivate
}

func IMMulticontextPrivateNewFromC(u unsafe.Pointer) *IMMulticontextPrivate {
	c := (*C.GtkIMMulticontextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IMMulticontextPrivate{native: c}

	return g
}

func (recv *IMMulticontextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconFactoryClass
*/
type IconFactoryClass struct {
	native *C.GtkIconFactoryClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IconFactoryClassNewFromC(u unsafe.Pointer) *IconFactoryClass {
	c := (*C.GtkIconFactoryClass)(u)
	if c == nil {
		return nil
	}

	g := &IconFactoryClass{native: c}

	return g
}

func (recv *IconFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconFactoryPrivate
*/
type IconFactoryPrivate struct {
	native *C.GtkIconFactoryPrivate
}

func IconFactoryPrivateNewFromC(u unsafe.Pointer) *IconFactoryPrivate {
	c := (*C.GtkIconFactoryPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IconFactoryPrivate{native: c}

	return g
}

func (recv *IconFactoryPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconInfoClass
*/
type IconInfoClass struct {
	native *C.GtkIconInfoClass
}

func IconInfoClassNewFromC(u unsafe.Pointer) *IconInfoClass {
	c := (*C.GtkIconInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &IconInfoClass{native: c}

	return g
}

func (recv *IconInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconSet
*/
type IconSet struct {
	native *C.GtkIconSet
}

func IconSetNewFromC(u unsafe.Pointer) *IconSet {
	c := (*C.GtkIconSet)(u)
	if c == nil {
		return nil
	}

	g := &IconSet{native: c}

	return g
}

func (recv *IconSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GtkIconSet. A #GtkIconSet represents a single icon
// in various sizes and widget states. It can provide a #GdkPixbuf
// for a given size and state on request, and automatically caches
// some of the rendered #GdkPixbuf objects.
//
// Normally you would use gtk_widget_render_icon_pixbuf() instead of
// using #GtkIconSet directly. The one case where you’d use
// #GtkIconSet is to create application-specific icon sets to place in
// a #GtkIconFactory.
/*

C function

gtk_icon_set_new
*/
func IconSetNew() *IconSet {
	retC := C.gtk_icon_set_new()
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkIconSet with @pixbuf as the default/fallback
// source image. If you don’t add any additional #GtkIconSource to the
// icon set, all variants of the icon will be created from @pixbuf,
// using scaling, pixelation, etc. as required to adjust the icon size
// or make the icon look insensitive/prelighted.
/*

C function

gtk_icon_set_new_from_pixbuf
*/
func IconSetNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *IconSet {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_icon_set_new_from_pixbuf(c_pixbuf)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Icon sets have a list of #GtkIconSource, which they use as base
// icons for rendering icons in different states and sizes. Icons are
// scaled, made to look insensitive, etc. in
// gtk_icon_set_render_icon(), but #GtkIconSet needs base images to
// work with. The base images and when to use them are described by
// a #GtkIconSource.
//
// This function copies @source, so you can reuse the same source immediately
// without affecting the icon set.
//
// An example of when you’d use this function: a web browser’s "Back
// to Previous Page" icon might point in a different direction in
// Hebrew and in English; it might look different when insensitive;
// and it might change size depending on toolbar mode (small/large
// icons). So a single icon set would contain all those variants of
// the icon, and you might add a separate source for each one.
//
// You should nearly always add a “default” icon source with all
// fields wildcarded, which will be used as a fallback if no more
// specific source matches. #GtkIconSet always prefers more specific
// icon sources to more generic icon sources. The order in which you
// add the sources to the icon set does not matter.
//
// gtk_icon_set_new_from_pixbuf() creates a new icon set with a
// default icon source based on the given pixbuf.
/*

C function

gtk_icon_set_add_source
*/
func (recv *IconSet) AddSource(source *IconSource) {
	c_source := (*C.GtkIconSource)(C.NULL)
	if source != nil {
		c_source = (*C.GtkIconSource)(source.ToC())
	}

	C.gtk_icon_set_add_source((*C.GtkIconSet)(recv.native), c_source)

	return
}

// Copies @icon_set by value.
/*

C function

gtk_icon_set_copy
*/
func (recv *IconSet) Copy() *IconSet {
	retC := C.gtk_icon_set_copy((*C.GtkIconSet)(recv.native))
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : output array param sizes

// Increments the reference count on @icon_set.
/*

C function

gtk_icon_set_ref
*/
func (recv *IconSet) Ref() *IconSet {
	retC := C.gtk_icon_set_ref((*C.GtkIconSet)(recv.native))
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Renders an icon using gtk_style_render_icon(). In most cases,
// gtk_widget_render_icon() is better, since it automatically provides
// most of the arguments from the current widget settings.  This
// function never returns %NULL; if the icon can’t be rendered
// (perhaps because an image file fails to load), a default "missing
// image" icon will be returned instead.
/*

C function

gtk_icon_set_render_icon
*/
func (recv *IconSet) RenderIcon(style *Style, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *gdkpixbuf.Pixbuf {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_direction := (C.GtkTextDirection)(direction)

	c_state := (C.GtkStateType)(state)

	c_size := (C.GtkIconSize)(size)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	retC := C.gtk_icon_set_render_icon((*C.GtkIconSet)(recv.native), c_style, c_direction, c_state, c_size, c_widget, c_detail)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count on @icon_set, and frees memory
// if the reference count reaches 0.
/*

C function

gtk_icon_set_unref
*/
func (recv *IconSet) Unref() {
	C.gtk_icon_set_unref((*C.GtkIconSet)(recv.native))

	return
}

/*

C type

GtkIconSource
*/
type IconSource struct {
	native *C.GtkIconSource
}

func IconSourceNewFromC(u unsafe.Pointer) *IconSource {
	c := (*C.GtkIconSource)(u)
	if c == nil {
		return nil
	}

	g := &IconSource{native: c}

	return g
}

func (recv *IconSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GtkIconSource. A #GtkIconSource contains a #GdkPixbuf (or
// image filename) that serves as the base image for one or more of the
// icons in a #GtkIconSet, along with a specification for which icons in the
// icon set will be based on that pixbuf or image file. An icon set contains
// a set of icons that represent “the same” logical concept in different states,
// different global text directions, and different sizes.
//
// So for example a web browser’s “Back to Previous Page” icon might
// point in a different direction in Hebrew and in English; it might
// look different when insensitive; and it might change size depending
// on toolbar mode (small/large icons). So a single icon set would
// contain all those variants of the icon. #GtkIconSet contains a list
// of #GtkIconSource from which it can derive specific icon variants in
// the set.
//
// In the simplest case, #GtkIconSet contains one source pixbuf from
// which it derives all variants. The convenience function
// gtk_icon_set_new_from_pixbuf() handles this case; if you only have
// one source pixbuf, just use that function.
//
// If you want to use a different base pixbuf for different icon
// variants, you create multiple icon sources, mark which variants
// they’ll be used to create, and add them to the icon set with
// gtk_icon_set_add_source().
//
// By default, the icon source has all parameters wildcarded. That is,
// the icon source will be used as the base icon for any desired text
// direction, widget state, or icon size.
/*

C function

gtk_icon_source_new
*/
func IconSourceNew() *IconSource {
	retC := C.gtk_icon_source_new()
	retGo := IconSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @source; mostly useful for language bindings.
/*

C function

gtk_icon_source_copy
*/
func (recv *IconSource) Copy() *IconSource {
	retC := C.gtk_icon_source_copy((*C.GtkIconSource)(recv.native))
	retGo := IconSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a dynamically-allocated icon source, along with its
// filename, size, and pixbuf fields if those are not %NULL.
/*

C function

gtk_icon_source_free
*/
func (recv *IconSource) Free() {
	C.gtk_icon_source_free((*C.GtkIconSource)(recv.native))

	return
}

// Obtains the text direction this icon source applies to. The return
// value is only useful/meaningful if the text direction is not
// wildcarded.
/*

C function

gtk_icon_source_get_direction
*/
func (recv *IconSource) GetDirection() TextDirection {
	retC := C.gtk_icon_source_get_direction((*C.GtkIconSource)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// Gets the value set by gtk_icon_source_set_direction_wildcarded().
/*

C function

gtk_icon_source_get_direction_wildcarded
*/
func (recv *IconSource) GetDirectionWildcarded() bool {
	retC := C.gtk_icon_source_get_direction_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the source filename, or %NULL if none is set. The
// filename is not a copy, and should not be modified or expected to
// persist beyond the lifetime of the icon source.
/*

C function

gtk_icon_source_get_filename
*/
func (recv *IconSource) GetFilename() string {
	retC := C.gtk_icon_source_get_filename((*C.GtkIconSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the source icon name, or %NULL if none is set. The
// icon_name is not a copy, and should not be modified or expected to
// persist beyond the lifetime of the icon source.
/*

C function

gtk_icon_source_get_icon_name
*/
func (recv *IconSource) GetIconName() string {
	retC := C.gtk_icon_source_get_icon_name((*C.GtkIconSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the source pixbuf, or %NULL if none is set.
// In addition, if a filename source is in use, this
// function in some cases will return the pixbuf from
// loaded from the filename. This is, for example, true
// for the GtkIconSource passed to the #GtkStyle render_icon()
// virtual function. The reference count on the pixbuf is
// not incremented.
/*

C function

gtk_icon_source_get_pixbuf
*/
func (recv *IconSource) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_icon_source_get_pixbuf((*C.GtkIconSource)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains the icon size this source applies to. The return value
// is only useful/meaningful if the icon size is not wildcarded.
/*

C function

gtk_icon_source_get_size
*/
func (recv *IconSource) GetSize() IconSize {
	retC := C.gtk_icon_source_get_size((*C.GtkIconSource)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// Gets the value set by gtk_icon_source_set_size_wildcarded().
/*

C function

gtk_icon_source_get_size_wildcarded
*/
func (recv *IconSource) GetSizeWildcarded() bool {
	retC := C.gtk_icon_source_get_size_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Obtains the widget state this icon source applies to. The return
// value is only useful/meaningful if the widget state is not
// wildcarded.
/*

C function

gtk_icon_source_get_state
*/
func (recv *IconSource) GetState() StateType {
	retC := C.gtk_icon_source_get_state((*C.GtkIconSource)(recv.native))
	retGo := (StateType)(retC)

	return retGo
}

// Gets the value set by gtk_icon_source_set_state_wildcarded().
/*

C function

gtk_icon_source_get_state_wildcarded
*/
func (recv *IconSource) GetStateWildcarded() bool {
	retC := C.gtk_icon_source_get_state_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the text direction this icon source is intended to be used
// with.
//
// Setting the text direction on an icon source makes no difference
// if the text direction is wildcarded. Therefore, you should usually
// call gtk_icon_source_set_direction_wildcarded() to un-wildcard it
// in addition to calling this function.
/*

C function

gtk_icon_source_set_direction
*/
func (recv *IconSource) SetDirection(direction TextDirection) {
	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_icon_source_set_direction((*C.GtkIconSource)(recv.native), c_direction)

	return
}

// If the text direction is wildcarded, this source can be used
// as the base image for an icon in any #GtkTextDirection.
// If the text direction is not wildcarded, then the
// text direction the icon source applies to should be set
// with gtk_icon_source_set_direction(), and the icon source
// will only be used with that text direction.
//
// #GtkIconSet prefers non-wildcarded sources (exact matches) over
// wildcarded sources, and will use an exact match when possible.
/*

C function

gtk_icon_source_set_direction_wildcarded
*/
func (recv *IconSource) SetDirectionWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_direction_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

// Sets the name of an image file to use as a base image when creating
// icon variants for #GtkIconSet. The filename must be absolute.
/*

C function

gtk_icon_source_set_filename
*/
func (recv *IconSource) SetFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_icon_source_set_filename((*C.GtkIconSource)(recv.native), c_filename)

	return
}

// Sets the name of an icon to look up in the current icon theme
// to use as a base image when creating icon variants for #GtkIconSet.
/*

C function

gtk_icon_source_set_icon_name
*/
func (recv *IconSource) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_icon_source_set_icon_name((*C.GtkIconSource)(recv.native), c_icon_name)

	return
}

// Sets a pixbuf to use as a base image when creating icon variants
// for #GtkIconSet.
/*

C function

gtk_icon_source_set_pixbuf
*/
func (recv *IconSource) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_icon_source_set_pixbuf((*C.GtkIconSource)(recv.native), c_pixbuf)

	return
}

// Sets the icon size this icon source is intended to be used
// with.
//
// Setting the icon size on an icon source makes no difference
// if the size is wildcarded. Therefore, you should usually
// call gtk_icon_source_set_size_wildcarded() to un-wildcard it
// in addition to calling this function.
/*

C function

gtk_icon_source_set_size
*/
func (recv *IconSource) SetSize(size IconSize) {
	c_size := (C.GtkIconSize)(size)

	C.gtk_icon_source_set_size((*C.GtkIconSource)(recv.native), c_size)

	return
}

// If the icon size is wildcarded, this source can be used as the base
// image for an icon of any size.  If the size is not wildcarded, then
// the size the source applies to should be set with
// gtk_icon_source_set_size() and the icon source will only be used
// with that specific size.
//
// #GtkIconSet prefers non-wildcarded sources (exact matches) over
// wildcarded sources, and will use an exact match when possible.
//
// #GtkIconSet will normally scale wildcarded source images to produce
// an appropriate icon at a given size, but will not change the size
// of source images that match exactly.
/*

C function

gtk_icon_source_set_size_wildcarded
*/
func (recv *IconSource) SetSizeWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_size_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

// Sets the widget state this icon source is intended to be used
// with.
//
// Setting the widget state on an icon source makes no difference
// if the state is wildcarded. Therefore, you should usually
// call gtk_icon_source_set_state_wildcarded() to un-wildcard it
// in addition to calling this function.
/*

C function

gtk_icon_source_set_state
*/
func (recv *IconSource) SetState(state StateType) {
	c_state := (C.GtkStateType)(state)

	C.gtk_icon_source_set_state((*C.GtkIconSource)(recv.native), c_state)

	return
}

// If the widget state is wildcarded, this source can be used as the
// base image for an icon in any #GtkStateType.  If the widget state
// is not wildcarded, then the state the source applies to should be
// set with gtk_icon_source_set_state() and the icon source will
// only be used with that specific state.
//
// #GtkIconSet prefers non-wildcarded sources (exact matches) over
// wildcarded sources, and will use an exact match when possible.
//
// #GtkIconSet will normally transform wildcarded source images to
// produce an appropriate icon for a given state, for example
// lightening an image on prelight, but will not modify source images
// that match exactly.
/*

C function

gtk_icon_source_set_state_wildcarded
*/
func (recv *IconSource) SetStateWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_state_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

/*

C type

GtkIconThemeClass
*/
type IconThemeClass struct {
	native *C.GtkIconThemeClass
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func IconThemeClassNewFromC(u unsafe.Pointer) *IconThemeClass {
	c := (*C.GtkIconThemeClass)(u)
	if c == nil {
		return nil
	}

	g := &IconThemeClass{native: c}

	return g
}

func (recv *IconThemeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconThemePrivate
*/
type IconThemePrivate struct {
	native *C.GtkIconThemePrivate
}

func IconThemePrivateNewFromC(u unsafe.Pointer) *IconThemePrivate {
	c := (*C.GtkIconThemePrivate)(u)
	if c == nil {
		return nil
	}

	g := &IconThemePrivate{native: c}

	return g
}

func (recv *IconThemePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconViewAccessibleClass
*/
type IconViewAccessibleClass struct {
	native *C.GtkIconViewAccessibleClass
	// parent_class : record
}

func IconViewAccessibleClassNewFromC(u unsafe.Pointer) *IconViewAccessibleClass {
	c := (*C.GtkIconViewAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &IconViewAccessibleClass{native: c}

	return g
}

func (recv *IconViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconViewAccessiblePrivate
*/
type IconViewAccessiblePrivate struct {
	native *C.GtkIconViewAccessiblePrivate
}

func IconViewAccessiblePrivateNewFromC(u unsafe.Pointer) *IconViewAccessiblePrivate {
	c := (*C.GtkIconViewAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &IconViewAccessiblePrivate{native: c}

	return g
}

func (recv *IconViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconViewClass
*/
type IconViewClass struct {
	native *C.GtkIconViewClass
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
	c := (*C.GtkIconViewClass)(u)
	if c == nil {
		return nil
	}

	g := &IconViewClass{native: c}

	return g
}

func (recv *IconViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkIconViewPrivate
*/
type IconViewPrivate struct {
	native *C.GtkIconViewPrivate
}

func IconViewPrivateNewFromC(u unsafe.Pointer) *IconViewPrivate {
	c := (*C.GtkIconViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IconViewPrivate{native: c}

	return g
}

func (recv *IconViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageAccessibleClass
*/
type ImageAccessibleClass struct {
	native *C.GtkImageAccessibleClass
	// parent_class : record
}

func ImageAccessibleClassNewFromC(u unsafe.Pointer) *ImageAccessibleClass {
	c := (*C.GtkImageAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ImageAccessibleClass{native: c}

	return g
}

func (recv *ImageAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageAccessiblePrivate
*/
type ImageAccessiblePrivate struct {
	native *C.GtkImageAccessiblePrivate
}

func ImageAccessiblePrivateNewFromC(u unsafe.Pointer) *ImageAccessiblePrivate {
	c := (*C.GtkImageAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ImageAccessiblePrivate{native: c}

	return g
}

func (recv *ImageAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageCellAccessibleClass
*/
type ImageCellAccessibleClass struct {
	native *C.GtkImageCellAccessibleClass
	// parent_class : record
}

func ImageCellAccessibleClassNewFromC(u unsafe.Pointer) *ImageCellAccessibleClass {
	c := (*C.GtkImageCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ImageCellAccessibleClass{native: c}

	return g
}

func (recv *ImageCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageCellAccessiblePrivate
*/
type ImageCellAccessiblePrivate struct {
	native *C.GtkImageCellAccessiblePrivate
}

func ImageCellAccessiblePrivateNewFromC(u unsafe.Pointer) *ImageCellAccessiblePrivate {
	c := (*C.GtkImageCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ImageCellAccessiblePrivate{native: c}

	return g
}

func (recv *ImageCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageClass
*/
type ImageClass struct {
	native *C.GtkImageClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ImageClassNewFromC(u unsafe.Pointer) *ImageClass {
	c := (*C.GtkImageClass)(u)
	if c == nil {
		return nil
	}

	g := &ImageClass{native: c}

	return g
}

func (recv *ImageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageMenuItemClass
*/
type ImageMenuItemClass struct {
	native *C.GtkImageMenuItemClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ImageMenuItemClassNewFromC(u unsafe.Pointer) *ImageMenuItemClass {
	c := (*C.GtkImageMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &ImageMenuItemClass{native: c}

	return g
}

func (recv *ImageMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImageMenuItemPrivate
*/
type ImageMenuItemPrivate struct {
	native *C.GtkImageMenuItemPrivate
}

func ImageMenuItemPrivateNewFromC(u unsafe.Pointer) *ImageMenuItemPrivate {
	c := (*C.GtkImageMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ImageMenuItemPrivate{native: c}

	return g
}

func (recv *ImageMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkImagePrivate
*/
type ImagePrivate struct {
	native *C.GtkImagePrivate
}

func ImagePrivateNewFromC(u unsafe.Pointer) *ImagePrivate {
	c := (*C.GtkImagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ImagePrivate{native: c}

	return g
}

func (recv *ImagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkInfoBarClass
*/
type InfoBarClass struct {
	native *C.GtkInfoBarClass
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func InfoBarClassNewFromC(u unsafe.Pointer) *InfoBarClass {
	c := (*C.GtkInfoBarClass)(u)
	if c == nil {
		return nil
	}

	g := &InfoBarClass{native: c}

	return g
}

func (recv *InfoBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkInfoBarPrivate
*/
type InfoBarPrivate struct {
	native *C.GtkInfoBarPrivate
}

func InfoBarPrivateNewFromC(u unsafe.Pointer) *InfoBarPrivate {
	c := (*C.GtkInfoBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InfoBarPrivate{native: c}

	return g
}

func (recv *InfoBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkInvisibleClass
*/
type InvisibleClass struct {
	native *C.GtkInvisibleClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func InvisibleClassNewFromC(u unsafe.Pointer) *InvisibleClass {
	c := (*C.GtkInvisibleClass)(u)
	if c == nil {
		return nil
	}

	g := &InvisibleClass{native: c}

	return g
}

func (recv *InvisibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkInvisiblePrivate
*/
type InvisiblePrivate struct {
	native *C.GtkInvisiblePrivate
}

func InvisiblePrivateNewFromC(u unsafe.Pointer) *InvisiblePrivate {
	c := (*C.GtkInvisiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &InvisiblePrivate{native: c}

	return g
}

func (recv *InvisiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLabelAccessibleClass
*/
type LabelAccessibleClass struct {
	native *C.GtkLabelAccessibleClass
	// parent_class : record
}

func LabelAccessibleClassNewFromC(u unsafe.Pointer) *LabelAccessibleClass {
	c := (*C.GtkLabelAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &LabelAccessibleClass{native: c}

	return g
}

func (recv *LabelAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLabelAccessiblePrivate
*/
type LabelAccessiblePrivate struct {
	native *C.GtkLabelAccessiblePrivate
}

func LabelAccessiblePrivateNewFromC(u unsafe.Pointer) *LabelAccessiblePrivate {
	c := (*C.GtkLabelAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &LabelAccessiblePrivate{native: c}

	return g
}

func (recv *LabelAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLabelClass
*/
type LabelClass struct {
	native *C.GtkLabelClass
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
	c := (*C.GtkLabelClass)(u)
	if c == nil {
		return nil
	}

	g := &LabelClass{native: c}

	return g
}

func (recv *LabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLabelPrivate
*/
type LabelPrivate struct {
	native *C.GtkLabelPrivate
}

func LabelPrivateNewFromC(u unsafe.Pointer) *LabelPrivate {
	c := (*C.GtkLabelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LabelPrivate{native: c}

	return g
}

func (recv *LabelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLabelSelectionInfo
*/
type LabelSelectionInfo struct {
	native *C.GtkLabelSelectionInfo
}

func LabelSelectionInfoNewFromC(u unsafe.Pointer) *LabelSelectionInfo {
	c := (*C.GtkLabelSelectionInfo)(u)
	if c == nil {
		return nil
	}

	g := &LabelSelectionInfo{native: c}

	return g
}

func (recv *LabelSelectionInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLayoutClass
*/
type LayoutClass struct {
	native *C.GtkLayoutClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	c := (*C.GtkLayoutClass)(u)
	if c == nil {
		return nil
	}

	g := &LayoutClass{native: c}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLayoutPrivate
*/
type LayoutPrivate struct {
	native *C.GtkLayoutPrivate
}

func LayoutPrivateNewFromC(u unsafe.Pointer) *LayoutPrivate {
	c := (*C.GtkLayoutPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LayoutPrivate{native: c}

	return g
}

func (recv *LayoutPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLevelBarAccessibleClass
*/
type LevelBarAccessibleClass struct {
	native *C.GtkLevelBarAccessibleClass
	// parent_class : record
}

func LevelBarAccessibleClassNewFromC(u unsafe.Pointer) *LevelBarAccessibleClass {
	c := (*C.GtkLevelBarAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarAccessibleClass{native: c}

	return g
}

func (recv *LevelBarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLevelBarAccessiblePrivate
*/
type LevelBarAccessiblePrivate struct {
	native *C.GtkLevelBarAccessiblePrivate
}

func LevelBarAccessiblePrivateNewFromC(u unsafe.Pointer) *LevelBarAccessiblePrivate {
	c := (*C.GtkLevelBarAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarAccessiblePrivate{native: c}

	return g
}

func (recv *LevelBarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLevelBarClass
*/
type LevelBarClass struct {
	native *C.GtkLevelBarClass
	// Private : parent_class
	// no type for offset_changed
	// Private : padding
}

func LevelBarClassNewFromC(u unsafe.Pointer) *LevelBarClass {
	c := (*C.GtkLevelBarClass)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarClass{native: c}

	return g
}

func (recv *LevelBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLevelBarPrivate
*/
type LevelBarPrivate struct {
	native *C.GtkLevelBarPrivate
}

func LevelBarPrivateNewFromC(u unsafe.Pointer) *LevelBarPrivate {
	c := (*C.GtkLevelBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarPrivate{native: c}

	return g
}

func (recv *LevelBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLinkButtonAccessibleClass
*/
type LinkButtonAccessibleClass struct {
	native *C.GtkLinkButtonAccessibleClass
	// parent_class : record
}

func LinkButtonAccessibleClassNewFromC(u unsafe.Pointer) *LinkButtonAccessibleClass {
	c := (*C.GtkLinkButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonAccessibleClass{native: c}

	return g
}

func (recv *LinkButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLinkButtonAccessiblePrivate
*/
type LinkButtonAccessiblePrivate struct {
	native *C.GtkLinkButtonAccessiblePrivate
}

func LinkButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *LinkButtonAccessiblePrivate {
	c := (*C.GtkLinkButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonAccessiblePrivate{native: c}

	return g
}

func (recv *LinkButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GtkLinkButtonClass contains only
// private data.
/*

C type

GtkLinkButtonClass
*/
type LinkButtonClass struct {
	native *C.GtkLinkButtonClass
	// Private : parent_class
	// no type for activate_link
	// no type for _gtk_padding1
	// no type for _gtk_padding2
	// no type for _gtk_padding3
	// no type for _gtk_padding4
}

func LinkButtonClassNewFromC(u unsafe.Pointer) *LinkButtonClass {
	c := (*C.GtkLinkButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonClass{native: c}

	return g
}

func (recv *LinkButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLinkButtonPrivate
*/
type LinkButtonPrivate struct {
	native *C.GtkLinkButtonPrivate
}

func LinkButtonPrivateNewFromC(u unsafe.Pointer) *LinkButtonPrivate {
	c := (*C.GtkLinkButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonPrivate{native: c}

	return g
}

func (recv *LinkButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListBoxAccessibleClass
*/
type ListBoxAccessibleClass struct {
	native *C.GtkListBoxAccessibleClass
	// parent_class : record
}

func ListBoxAccessibleClassNewFromC(u unsafe.Pointer) *ListBoxAccessibleClass {
	c := (*C.GtkListBoxAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxAccessibleClass{native: c}

	return g
}

func (recv *ListBoxAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListBoxAccessiblePrivate
*/
type ListBoxAccessiblePrivate struct {
	native *C.GtkListBoxAccessiblePrivate
}

func ListBoxAccessiblePrivateNewFromC(u unsafe.Pointer) *ListBoxAccessiblePrivate {
	c := (*C.GtkListBoxAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxAccessiblePrivate{native: c}

	return g
}

func (recv *ListBoxAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListBoxClass
*/
type ListBoxClass struct {
	native *C.GtkListBoxClass
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
	c := (*C.GtkListBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxClass{native: c}

	return g
}

func (recv *ListBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListBoxRowAccessibleClass
*/
type ListBoxRowAccessibleClass struct {
	native *C.GtkListBoxRowAccessibleClass
	// parent_class : record
}

func ListBoxRowAccessibleClassNewFromC(u unsafe.Pointer) *ListBoxRowAccessibleClass {
	c := (*C.GtkListBoxRowAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRowAccessibleClass{native: c}

	return g
}

func (recv *ListBoxRowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListBoxRowClass
*/
type ListBoxRowClass struct {
	native *C.GtkListBoxRowClass
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

func ListBoxRowClassNewFromC(u unsafe.Pointer) *ListBoxRowClass {
	c := (*C.GtkListBoxRowClass)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRowClass{native: c}

	return g
}

func (recv *ListBoxRowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListStoreClass
*/
type ListStoreClass struct {
	native *C.GtkListStoreClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ListStoreClassNewFromC(u unsafe.Pointer) *ListStoreClass {
	c := (*C.GtkListStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &ListStoreClass{native: c}

	return g
}

func (recv *ListStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkListStorePrivate
*/
type ListStorePrivate struct {
	native *C.GtkListStorePrivate
}

func ListStorePrivateNewFromC(u unsafe.Pointer) *ListStorePrivate {
	c := (*C.GtkListStorePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ListStorePrivate{native: c}

	return g
}

func (recv *ListStorePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLockButtonAccessibleClass
*/
type LockButtonAccessibleClass struct {
	native *C.GtkLockButtonAccessibleClass
	// parent_class : record
}

func LockButtonAccessibleClassNewFromC(u unsafe.Pointer) *LockButtonAccessibleClass {
	c := (*C.GtkLockButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonAccessibleClass{native: c}

	return g
}

func (recv *LockButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLockButtonAccessiblePrivate
*/
type LockButtonAccessiblePrivate struct {
	native *C.GtkLockButtonAccessiblePrivate
}

func LockButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *LockButtonAccessiblePrivate {
	c := (*C.GtkLockButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonAccessiblePrivate{native: c}

	return g
}

func (recv *LockButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLockButtonClass
*/
type LockButtonClass struct {
	native *C.GtkLockButtonClass
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
	c := (*C.GtkLockButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonClass{native: c}

	return g
}

func (recv *LockButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkLockButtonPrivate
*/
type LockButtonPrivate struct {
	native *C.GtkLockButtonPrivate
}

func LockButtonPrivateNewFromC(u unsafe.Pointer) *LockButtonPrivate {
	c := (*C.GtkLockButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonPrivate{native: c}

	return g
}

func (recv *LockButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuAccessibleClass
*/
type MenuAccessibleClass struct {
	native *C.GtkMenuAccessibleClass
	// parent_class : record
}

func MenuAccessibleClassNewFromC(u unsafe.Pointer) *MenuAccessibleClass {
	c := (*C.GtkMenuAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuAccessibleClass{native: c}

	return g
}

func (recv *MenuAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuAccessiblePrivate
*/
type MenuAccessiblePrivate struct {
	native *C.GtkMenuAccessiblePrivate
}

func MenuAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuAccessiblePrivate {
	c := (*C.GtkMenuAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuAccessiblePrivate{native: c}

	return g
}

func (recv *MenuAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuBarClass
*/
type MenuBarClass struct {
	native *C.GtkMenuBarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuBarClassNewFromC(u unsafe.Pointer) *MenuBarClass {
	c := (*C.GtkMenuBarClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuBarClass{native: c}

	return g
}

func (recv *MenuBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuBarPrivate
*/
type MenuBarPrivate struct {
	native *C.GtkMenuBarPrivate
}

func MenuBarPrivateNewFromC(u unsafe.Pointer) *MenuBarPrivate {
	c := (*C.GtkMenuBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuBarPrivate{native: c}

	return g
}

func (recv *MenuBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuButtonAccessibleClass
*/
type MenuButtonAccessibleClass struct {
	native *C.GtkMenuButtonAccessibleClass
	// parent_class : record
}

func MenuButtonAccessibleClassNewFromC(u unsafe.Pointer) *MenuButtonAccessibleClass {
	c := (*C.GtkMenuButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonAccessibleClass{native: c}

	return g
}

func (recv *MenuButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuButtonAccessiblePrivate
*/
type MenuButtonAccessiblePrivate struct {
	native *C.GtkMenuButtonAccessiblePrivate
}

func MenuButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuButtonAccessiblePrivate {
	c := (*C.GtkMenuButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonAccessiblePrivate{native: c}

	return g
}

func (recv *MenuButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuButtonClass
*/
type MenuButtonClass struct {
	native *C.GtkMenuButtonClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuButtonClassNewFromC(u unsafe.Pointer) *MenuButtonClass {
	c := (*C.GtkMenuButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonClass{native: c}

	return g
}

func (recv *MenuButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuButtonPrivate
*/
type MenuButtonPrivate struct {
	native *C.GtkMenuButtonPrivate
}

func MenuButtonPrivateNewFromC(u unsafe.Pointer) *MenuButtonPrivate {
	c := (*C.GtkMenuButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonPrivate{native: c}

	return g
}

func (recv *MenuButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuClass
*/
type MenuClass struct {
	native *C.GtkMenuClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuClassNewFromC(u unsafe.Pointer) *MenuClass {
	c := (*C.GtkMenuClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuClass{native: c}

	return g
}

func (recv *MenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuItemAccessibleClass
*/
type MenuItemAccessibleClass struct {
	native *C.GtkMenuItemAccessibleClass
	// parent_class : record
}

func MenuItemAccessibleClassNewFromC(u unsafe.Pointer) *MenuItemAccessibleClass {
	c := (*C.GtkMenuItemAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemAccessibleClass{native: c}

	return g
}

func (recv *MenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuItemAccessiblePrivate
*/
type MenuItemAccessiblePrivate struct {
	native *C.GtkMenuItemAccessiblePrivate
}

func MenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuItemAccessiblePrivate {
	c := (*C.GtkMenuItemAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemAccessiblePrivate{native: c}

	return g
}

func (recv *MenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuItemClass
*/
type MenuItemClass struct {
	native *C.GtkMenuItemClass
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
	c := (*C.GtkMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemClass{native: c}

	return g
}

func (recv *MenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuItemPrivate
*/
type MenuItemPrivate struct {
	native *C.GtkMenuItemPrivate
}

func MenuItemPrivateNewFromC(u unsafe.Pointer) *MenuItemPrivate {
	c := (*C.GtkMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemPrivate{native: c}

	return g
}

func (recv *MenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuPrivate
*/
type MenuPrivate struct {
	native *C.GtkMenuPrivate
}

func MenuPrivateNewFromC(u unsafe.Pointer) *MenuPrivate {
	c := (*C.GtkMenuPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuPrivate{native: c}

	return g
}

func (recv *MenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuShellAccessibleClass
*/
type MenuShellAccessibleClass struct {
	native *C.GtkMenuShellAccessibleClass
	// parent_class : record
}

func MenuShellAccessibleClassNewFromC(u unsafe.Pointer) *MenuShellAccessibleClass {
	c := (*C.GtkMenuShellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellAccessibleClass{native: c}

	return g
}

func (recv *MenuShellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuShellAccessiblePrivate
*/
type MenuShellAccessiblePrivate struct {
	native *C.GtkMenuShellAccessiblePrivate
}

func MenuShellAccessiblePrivateNewFromC(u unsafe.Pointer) *MenuShellAccessiblePrivate {
	c := (*C.GtkMenuShellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellAccessiblePrivate{native: c}

	return g
}

func (recv *MenuShellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuShellClass
*/
type MenuShellClass struct {
	native *C.GtkMenuShellClass
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
	c := (*C.GtkMenuShellClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellClass{native: c}

	return g
}

func (recv *MenuShellClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuShellPrivate
*/
type MenuShellPrivate struct {
	native *C.GtkMenuShellPrivate
}

func MenuShellPrivateNewFromC(u unsafe.Pointer) *MenuShellPrivate {
	c := (*C.GtkMenuShellPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellPrivate{native: c}

	return g
}

func (recv *MenuShellPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuToolButtonClass
*/
type MenuToolButtonClass struct {
	native *C.GtkMenuToolButtonClass
	// parent_class : record
	// no type for show_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MenuToolButtonClassNewFromC(u unsafe.Pointer) *MenuToolButtonClass {
	c := (*C.GtkMenuToolButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuToolButtonClass{native: c}

	return g
}

func (recv *MenuToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMenuToolButtonPrivate
*/
type MenuToolButtonPrivate struct {
	native *C.GtkMenuToolButtonPrivate
}

func MenuToolButtonPrivateNewFromC(u unsafe.Pointer) *MenuToolButtonPrivate {
	c := (*C.GtkMenuToolButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuToolButtonPrivate{native: c}

	return g
}

func (recv *MenuToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMessageDialogClass
*/
type MessageDialogClass struct {
	native *C.GtkMessageDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MessageDialogClassNewFromC(u unsafe.Pointer) *MessageDialogClass {
	c := (*C.GtkMessageDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &MessageDialogClass{native: c}

	return g
}

func (recv *MessageDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMessageDialogPrivate
*/
type MessageDialogPrivate struct {
	native *C.GtkMessageDialogPrivate
}

func MessageDialogPrivateNewFromC(u unsafe.Pointer) *MessageDialogPrivate {
	c := (*C.GtkMessageDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MessageDialogPrivate{native: c}

	return g
}

func (recv *MessageDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMiscClass
*/
type MiscClass struct {
	native *C.GtkMiscClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MiscClassNewFromC(u unsafe.Pointer) *MiscClass {
	c := (*C.GtkMiscClass)(u)
	if c == nil {
		return nil
	}

	g := &MiscClass{native: c}

	return g
}

func (recv *MiscClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMiscPrivate
*/
type MiscPrivate struct {
	native *C.GtkMiscPrivate
}

func MiscPrivateNewFromC(u unsafe.Pointer) *MiscPrivate {
	c := (*C.GtkMiscPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MiscPrivate{native: c}

	return g
}

func (recv *MiscPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMountOperationClass
*/
type MountOperationClass struct {
	native *C.GtkMountOperationClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func MountOperationClassNewFromC(u unsafe.Pointer) *MountOperationClass {
	c := (*C.GtkMountOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationClass{native: c}

	return g
}

func (recv *MountOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkMountOperationPrivate
*/
type MountOperationPrivate struct {
	native *C.GtkMountOperationPrivate
}

func MountOperationPrivateNewFromC(u unsafe.Pointer) *MountOperationPrivate {
	c := (*C.GtkMountOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationPrivate{native: c}

	return g
}

func (recv *MountOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookAccessibleClass
*/
type NotebookAccessibleClass struct {
	native *C.GtkNotebookAccessibleClass
	// parent_class : record
}

func NotebookAccessibleClassNewFromC(u unsafe.Pointer) *NotebookAccessibleClass {
	c := (*C.GtkNotebookAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &NotebookAccessibleClass{native: c}

	return g
}

func (recv *NotebookAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookAccessiblePrivate
*/
type NotebookAccessiblePrivate struct {
	native *C.GtkNotebookAccessiblePrivate
}

func NotebookAccessiblePrivateNewFromC(u unsafe.Pointer) *NotebookAccessiblePrivate {
	c := (*C.GtkNotebookAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &NotebookAccessiblePrivate{native: c}

	return g
}

func (recv *NotebookAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookClass
*/
type NotebookClass struct {
	native *C.GtkNotebookClass
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
	c := (*C.GtkNotebookClass)(u)
	if c == nil {
		return nil
	}

	g := &NotebookClass{native: c}

	return g
}

func (recv *NotebookClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookPageAccessibleClass
*/
type NotebookPageAccessibleClass struct {
	native *C.GtkNotebookPageAccessibleClass
	// parent_class : record
}

func NotebookPageAccessibleClassNewFromC(u unsafe.Pointer) *NotebookPageAccessibleClass {
	c := (*C.GtkNotebookPageAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &NotebookPageAccessibleClass{native: c}

	return g
}

func (recv *NotebookPageAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookPageAccessiblePrivate
*/
type NotebookPageAccessiblePrivate struct {
	native *C.GtkNotebookPageAccessiblePrivate
}

func NotebookPageAccessiblePrivateNewFromC(u unsafe.Pointer) *NotebookPageAccessiblePrivate {
	c := (*C.GtkNotebookPageAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &NotebookPageAccessiblePrivate{native: c}

	return g
}

func (recv *NotebookPageAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNotebookPrivate
*/
type NotebookPrivate struct {
	native *C.GtkNotebookPrivate
}

func NotebookPrivateNewFromC(u unsafe.Pointer) *NotebookPrivate {
	c := (*C.GtkNotebookPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NotebookPrivate{native: c}

	return g
}

func (recv *NotebookPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNumerableIconClass
*/
type NumerableIconClass struct {
	native *C.GtkNumerableIconClass
	// parent_class : record
	// no type for padding
}

func NumerableIconClassNewFromC(u unsafe.Pointer) *NumerableIconClass {
	c := (*C.GtkNumerableIconClass)(u)
	if c == nil {
		return nil
	}

	g := &NumerableIconClass{native: c}

	return g
}

func (recv *NumerableIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkNumerableIconPrivate
*/
type NumerableIconPrivate struct {
	native *C.GtkNumerableIconPrivate
}

func NumerableIconPrivateNewFromC(u unsafe.Pointer) *NumerableIconPrivate {
	c := (*C.GtkNumerableIconPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NumerableIconPrivate{native: c}

	return g
}

func (recv *NumerableIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkOffscreenWindowClass
*/
type OffscreenWindowClass struct {
	native *C.GtkOffscreenWindowClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func OffscreenWindowClassNewFromC(u unsafe.Pointer) *OffscreenWindowClass {
	c := (*C.GtkOffscreenWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &OffscreenWindowClass{native: c}

	return g
}

func (recv *OffscreenWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkOrientableIface
*/
type OrientableIface struct {
	native *C.GtkOrientableIface
	// base_iface : record
}

func OrientableIfaceNewFromC(u unsafe.Pointer) *OrientableIface {
	c := (*C.GtkOrientableIface)(u)
	if c == nil {
		return nil
	}

	g := &OrientableIface{native: c}

	return g
}

func (recv *OrientableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkOverlayClass
*/
type OverlayClass struct {
	native *C.GtkOverlayClass
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
	c := (*C.GtkOverlayClass)(u)
	if c == nil {
		return nil
	}

	g := &OverlayClass{native: c}

	return g
}

func (recv *OverlayClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkOverlayPrivate
*/
type OverlayPrivate struct {
	native *C.GtkOverlayPrivate
}

func OverlayPrivateNewFromC(u unsafe.Pointer) *OverlayPrivate {
	c := (*C.GtkOverlayPrivate)(u)
	if c == nil {
		return nil
	}

	g := &OverlayPrivate{native: c}

	return g
}

func (recv *OverlayPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// See also gtk_print_settings_set_page_ranges().
/*

C type

GtkPageRange
*/
type PageRange struct {
	native *C.GtkPageRange
	Start  int32
	End    int32
}

func PageRangeNewFromC(u unsafe.Pointer) *PageRange {
	c := (*C.GtkPageRange)(u)
	if c == nil {
		return nil
	}

	g := &PageRange{
		End:    (int32)(c.end),
		Start:  (int32)(c.start),
		native: c,
	}

	return g
}

func (recv *PageRange) ToC() unsafe.Pointer {
	recv.native.start =
		(C.gint)(recv.Start)
	recv.native.end =
		(C.gint)(recv.End)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPanedAccessibleClass
*/
type PanedAccessibleClass struct {
	native *C.GtkPanedAccessibleClass
	// parent_class : record
}

func PanedAccessibleClassNewFromC(u unsafe.Pointer) *PanedAccessibleClass {
	c := (*C.GtkPanedAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &PanedAccessibleClass{native: c}

	return g
}

func (recv *PanedAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPanedAccessiblePrivate
*/
type PanedAccessiblePrivate struct {
	native *C.GtkPanedAccessiblePrivate
}

func PanedAccessiblePrivateNewFromC(u unsafe.Pointer) *PanedAccessiblePrivate {
	c := (*C.GtkPanedAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &PanedAccessiblePrivate{native: c}

	return g
}

func (recv *PanedAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPanedClass
*/
type PanedClass struct {
	native *C.GtkPanedClass
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
	c := (*C.GtkPanedClass)(u)
	if c == nil {
		return nil
	}

	g := &PanedClass{native: c}

	return g
}

func (recv *PanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPanedPrivate
*/
type PanedPrivate struct {
	native *C.GtkPanedPrivate
}

func PanedPrivateNewFromC(u unsafe.Pointer) *PanedPrivate {
	c := (*C.GtkPanedPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PanedPrivate{native: c}

	return g
}

func (recv *PanedPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkPaperSize handles paper sizes. It uses the standard called
// [PWG 5101.1-2002 PWG: Standard for Media Standardized Names](http://www.pwg.org/standards.html)
// to name the paper sizes (and to get the data for the page sizes).
// In addition to standard paper sizes, GtkPaperSize allows to
// construct custom paper sizes with arbitrary dimensions.
//
// The #GtkPaperSize object stores not only the dimensions (width
// and height) of a paper size and its name, it also provides
// default [print margins][print-margins].
//
// Printing support has been added in GTK+ 2.10.
/*

C type

GtkPaperSize
*/
type PaperSize struct {
	native *C.GtkPaperSize
}

func PaperSizeNewFromC(u unsafe.Pointer) *PaperSize {
	c := (*C.GtkPaperSize)(u)
	if c == nil {
		return nil
	}

	g := &PaperSize{native: c}

	return g
}

func (recv *PaperSize) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns %TRUE if @size is not a standard paper size.
/*

C function

gtk_paper_size_is_custom
*/
func (recv *PaperSize) IsCustom() bool {
	retC := C.gtk_paper_size_is_custom((*C.GtkPaperSize)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @size is an IPP standard paper size.
/*

C function

gtk_paper_size_is_ipp
*/
func (recv *PaperSize) IsIpp() bool {
	retC := C.gtk_paper_size_is_ipp((*C.GtkPaperSize)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

/*

C type

GtkPlacesSidebarClass
*/
type PlacesSidebarClass struct {
	native *C.GtkPlacesSidebarClass
}

func PlacesSidebarClassNewFromC(u unsafe.Pointer) *PlacesSidebarClass {
	c := (*C.GtkPlacesSidebarClass)(u)
	if c == nil {
		return nil
	}

	g := &PlacesSidebarClass{native: c}

	return g
}

func (recv *PlacesSidebarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPlugClass
*/
type PlugClass struct {
	native *C.GtkPlugClass
	// parent_class : record
	// no type for embedded
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func PlugClassNewFromC(u unsafe.Pointer) *PlugClass {
	c := (*C.GtkPlugClass)(u)
	if c == nil {
		return nil
	}

	g := &PlugClass{native: c}

	return g
}

func (recv *PlugClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPlugPrivate
*/
type PlugPrivate struct {
	native *C.GtkPlugPrivate
}

func PlugPrivateNewFromC(u unsafe.Pointer) *PlugPrivate {
	c := (*C.GtkPlugPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PlugPrivate{native: c}

	return g
}

func (recv *PlugPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPopoverAccessibleClass
*/
type PopoverAccessibleClass struct {
	native *C.GtkPopoverAccessibleClass
	// parent_class : record
}

func PopoverAccessibleClassNewFromC(u unsafe.Pointer) *PopoverAccessibleClass {
	c := (*C.GtkPopoverAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &PopoverAccessibleClass{native: c}

	return g
}

func (recv *PopoverAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPopoverClass
*/
type PopoverClass struct {
	native *C.GtkPopoverClass
	// parent_class : record
	// no type for closed
	// Private : reserved
}

func PopoverClassNewFromC(u unsafe.Pointer) *PopoverClass {
	c := (*C.GtkPopoverClass)(u)
	if c == nil {
		return nil
	}

	g := &PopoverClass{native: c}

	return g
}

func (recv *PopoverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPopoverMenuClass
*/
type PopoverMenuClass struct {
	native *C.GtkPopoverMenuClass
	// parent_class : record
	// Private : reserved
}

func PopoverMenuClassNewFromC(u unsafe.Pointer) *PopoverMenuClass {
	c := (*C.GtkPopoverMenuClass)(u)
	if c == nil {
		return nil
	}

	g := &PopoverMenuClass{native: c}

	return g
}

func (recv *PopoverMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPopoverPrivate
*/
type PopoverPrivate struct {
	native *C.GtkPopoverPrivate
}

func PopoverPrivateNewFromC(u unsafe.Pointer) *PopoverPrivate {
	c := (*C.GtkPopoverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PopoverPrivate{native: c}

	return g
}

func (recv *PopoverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPrintOperationClass
*/
type PrintOperationClass struct {
	native *C.GtkPrintOperationClass
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
	c := (*C.GtkPrintOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationClass{native: c}

	return g
}

func (recv *PrintOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPrintOperationPreviewIface
*/
type PrintOperationPreviewIface struct {
	native *C.GtkPrintOperationPreviewIface
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
	c := (*C.GtkPrintOperationPreviewIface)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationPreviewIface{native: c}

	return g
}

func (recv *PrintOperationPreviewIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkPrintOperationPrivate
*/
type PrintOperationPrivate struct {
	native *C.GtkPrintOperationPrivate
}

func PrintOperationPrivateNewFromC(u unsafe.Pointer) *PrintOperationPrivate {
	c := (*C.GtkPrintOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationPrivate{native: c}

	return g
}

func (recv *PrintOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkProgressBarAccessibleClass
*/
type ProgressBarAccessibleClass struct {
	native *C.GtkProgressBarAccessibleClass
	// parent_class : record
}

func ProgressBarAccessibleClassNewFromC(u unsafe.Pointer) *ProgressBarAccessibleClass {
	c := (*C.GtkProgressBarAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarAccessibleClass{native: c}

	return g
}

func (recv *ProgressBarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkProgressBarAccessiblePrivate
*/
type ProgressBarAccessiblePrivate struct {
	native *C.GtkProgressBarAccessiblePrivate
}

func ProgressBarAccessiblePrivateNewFromC(u unsafe.Pointer) *ProgressBarAccessiblePrivate {
	c := (*C.GtkProgressBarAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarAccessiblePrivate{native: c}

	return g
}

func (recv *ProgressBarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkProgressBarClass
*/
type ProgressBarClass struct {
	native *C.GtkProgressBarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ProgressBarClassNewFromC(u unsafe.Pointer) *ProgressBarClass {
	c := (*C.GtkProgressBarClass)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarClass{native: c}

	return g
}

func (recv *ProgressBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkProgressBarPrivate
*/
type ProgressBarPrivate struct {
	native *C.GtkProgressBarPrivate
}

func ProgressBarPrivateNewFromC(u unsafe.Pointer) *ProgressBarPrivate {
	c := (*C.GtkProgressBarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarPrivate{native: c}

	return g
}

func (recv *ProgressBarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioActionClass
*/
type RadioActionClass struct {
	native *C.GtkRadioActionClass
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioActionClassNewFromC(u unsafe.Pointer) *RadioActionClass {
	c := (*C.GtkRadioActionClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioActionClass{native: c}

	return g
}

func (recv *RadioActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkRadioActionEntry structs are used with
// gtk_action_group_add_radio_actions() to construct groups of radio actions.
/*

C type

GtkRadioActionEntry
*/
type RadioActionEntry struct {
	native      *C.GtkRadioActionEntry
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	Value       int32
}

func RadioActionEntryNewFromC(u unsafe.Pointer) *RadioActionEntry {
	c := (*C.GtkRadioActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &RadioActionEntry{
		Accelerator: C.GoString(c.accelerator),
		Label:       C.GoString(c.label),
		Name:        C.GoString(c.name),
		StockId:     C.GoString(c.stock_id),
		Tooltip:     C.GoString(c.tooltip),
		Value:       (int32)(c.value),
		native:      c,
	}

	return g
}

func (recv *RadioActionEntry) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.stock_id =
		C.CString(recv.StockId)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.accelerator =
		C.CString(recv.Accelerator)
	recv.native.tooltip =
		C.CString(recv.Tooltip)
	recv.native.value =
		(C.gint)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioActionPrivate
*/
type RadioActionPrivate struct {
	native *C.GtkRadioActionPrivate
}

func RadioActionPrivateNewFromC(u unsafe.Pointer) *RadioActionPrivate {
	c := (*C.GtkRadioActionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RadioActionPrivate{native: c}

	return g
}

func (recv *RadioActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioButtonAccessibleClass
*/
type RadioButtonAccessibleClass struct {
	native *C.GtkRadioButtonAccessibleClass
	// parent_class : record
}

func RadioButtonAccessibleClassNewFromC(u unsafe.Pointer) *RadioButtonAccessibleClass {
	c := (*C.GtkRadioButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonAccessibleClass{native: c}

	return g
}

func (recv *RadioButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioButtonAccessiblePrivate
*/
type RadioButtonAccessiblePrivate struct {
	native *C.GtkRadioButtonAccessiblePrivate
}

func RadioButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *RadioButtonAccessiblePrivate {
	c := (*C.GtkRadioButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonAccessiblePrivate{native: c}

	return g
}

func (recv *RadioButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioButtonClass
*/
type RadioButtonClass struct {
	native *C.GtkRadioButtonClass
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioButtonClassNewFromC(u unsafe.Pointer) *RadioButtonClass {
	c := (*C.GtkRadioButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonClass{native: c}

	return g
}

func (recv *RadioButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioButtonPrivate
*/
type RadioButtonPrivate struct {
	native *C.GtkRadioButtonPrivate
}

func RadioButtonPrivateNewFromC(u unsafe.Pointer) *RadioButtonPrivate {
	c := (*C.GtkRadioButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonPrivate{native: c}

	return g
}

func (recv *RadioButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioMenuItemAccessibleClass
*/
type RadioMenuItemAccessibleClass struct {
	native *C.GtkRadioMenuItemAccessibleClass
	// parent_class : record
}

func RadioMenuItemAccessibleClassNewFromC(u unsafe.Pointer) *RadioMenuItemAccessibleClass {
	c := (*C.GtkRadioMenuItemAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemAccessibleClass{native: c}

	return g
}

func (recv *RadioMenuItemAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioMenuItemAccessiblePrivate
*/
type RadioMenuItemAccessiblePrivate struct {
	native *C.GtkRadioMenuItemAccessiblePrivate
}

func RadioMenuItemAccessiblePrivateNewFromC(u unsafe.Pointer) *RadioMenuItemAccessiblePrivate {
	c := (*C.GtkRadioMenuItemAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemAccessiblePrivate{native: c}

	return g
}

func (recv *RadioMenuItemAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioMenuItemClass
*/
type RadioMenuItemClass struct {
	native *C.GtkRadioMenuItemClass
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioMenuItemClassNewFromC(u unsafe.Pointer) *RadioMenuItemClass {
	c := (*C.GtkRadioMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemClass{native: c}

	return g
}

func (recv *RadioMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioMenuItemPrivate
*/
type RadioMenuItemPrivate struct {
	native *C.GtkRadioMenuItemPrivate
}

func RadioMenuItemPrivateNewFromC(u unsafe.Pointer) *RadioMenuItemPrivate {
	c := (*C.GtkRadioMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemPrivate{native: c}

	return g
}

func (recv *RadioMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRadioToolButtonClass
*/
type RadioToolButtonClass struct {
	native *C.GtkRadioToolButtonClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RadioToolButtonClassNewFromC(u unsafe.Pointer) *RadioToolButtonClass {
	c := (*C.GtkRadioToolButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &RadioToolButtonClass{native: c}

	return g
}

func (recv *RadioToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRangeAccessibleClass
*/
type RangeAccessibleClass struct {
	native *C.GtkRangeAccessibleClass
	// parent_class : record
}

func RangeAccessibleClassNewFromC(u unsafe.Pointer) *RangeAccessibleClass {
	c := (*C.GtkRangeAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &RangeAccessibleClass{native: c}

	return g
}

func (recv *RangeAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRangeAccessiblePrivate
*/
type RangeAccessiblePrivate struct {
	native *C.GtkRangeAccessiblePrivate
}

func RangeAccessiblePrivateNewFromC(u unsafe.Pointer) *RangeAccessiblePrivate {
	c := (*C.GtkRangeAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RangeAccessiblePrivate{native: c}

	return g
}

func (recv *RangeAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRangeClass
*/
type RangeClass struct {
	native *C.GtkRangeClass
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
	c := (*C.GtkRangeClass)(u)
	if c == nil {
		return nil
	}

	g := &RangeClass{
		SliderDetail:  C.GoString(c.slider_detail),
		StepperDetail: C.GoString(c.stepper_detail),
		native:        c,
	}

	return g
}

func (recv *RangeClass) ToC() unsafe.Pointer {
	recv.native.slider_detail =
		C.CString(recv.SliderDetail)
	recv.native.stepper_detail =
		C.CString(recv.StepperDetail)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRangePrivate
*/
type RangePrivate struct {
	native *C.GtkRangePrivate
}

func RangePrivateNewFromC(u unsafe.Pointer) *RangePrivate {
	c := (*C.GtkRangePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RangePrivate{native: c}

	return g
}

func (recv *RangePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRcContext
*/
type RcContext struct {
	native *C.GtkRcContext
}

func RcContextNewFromC(u unsafe.Pointer) *RcContext {
	c := (*C.GtkRcContext)(u)
	if c == nil {
		return nil
	}

	g := &RcContext{native: c}

	return g
}

func (recv *RcContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Deprecated
/*

C type

GtkRcProperty
*/
type RcProperty struct {
	native       *C.GtkRcProperty
	TypeName     glib.Quark
	PropertyName glib.Quark
	Origin       string
	// value : record
}

func RcPropertyNewFromC(u unsafe.Pointer) *RcProperty {
	c := (*C.GtkRcProperty)(u)
	if c == nil {
		return nil
	}

	g := &RcProperty{
		Origin:       C.GoString(c.origin),
		PropertyName: (glib.Quark)(c.property_name),
		TypeName:     (glib.Quark)(c.type_name),
		native:       c,
	}

	return g
}

func (recv *RcProperty) ToC() unsafe.Pointer {
	recv.native.type_name =
		(C.GQuark)(recv.TypeName)
	recv.native.property_name =
		(C.GQuark)(recv.PropertyName)
	recv.native.origin =
		C.CString(recv.Origin)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRcStyleClass
*/
type RcStyleClass struct {
	native *C.GtkRcStyleClass
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
	c := (*C.GtkRcStyleClass)(u)
	if c == nil {
		return nil
	}

	g := &RcStyleClass{native: c}

	return g
}

func (recv *RcStyleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentActionClass
*/
type RecentActionClass struct {
	native *C.GtkRecentActionClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentActionClassNewFromC(u unsafe.Pointer) *RecentActionClass {
	c := (*C.GtkRecentActionClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentActionClass{native: c}

	return g
}

func (recv *RecentActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentActionPrivate
*/
type RecentActionPrivate struct {
	native *C.GtkRecentActionPrivate
}

func RecentActionPrivateNewFromC(u unsafe.Pointer) *RecentActionPrivate {
	c := (*C.GtkRecentActionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RecentActionPrivate{native: c}

	return g
}

func (recv *RecentActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserDialogClass
*/
type RecentChooserDialogClass struct {
	native *C.GtkRecentChooserDialogClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentChooserDialogClassNewFromC(u unsafe.Pointer) *RecentChooserDialogClass {
	c := (*C.GtkRecentChooserDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserDialogClass{native: c}

	return g
}

func (recv *RecentChooserDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserDialogPrivate
*/
type RecentChooserDialogPrivate struct {
	native *C.GtkRecentChooserDialogPrivate
}

func RecentChooserDialogPrivateNewFromC(u unsafe.Pointer) *RecentChooserDialogPrivate {
	c := (*C.GtkRecentChooserDialogPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserDialogPrivate{native: c}

	return g
}

func (recv *RecentChooserDialogPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserIface
*/
type RecentChooserIface struct {
	native *C.GtkRecentChooserIface
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
	c := (*C.GtkRecentChooserIface)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserIface{native: c}

	return g
}

func (recv *RecentChooserIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserMenuClass
*/
type RecentChooserMenuClass struct {
	native *C.GtkRecentChooserMenuClass
	// parent_class : record
	// no type for gtk_recent1
	// no type for gtk_recent2
	// no type for gtk_recent3
	// no type for gtk_recent4
}

func RecentChooserMenuClassNewFromC(u unsafe.Pointer) *RecentChooserMenuClass {
	c := (*C.GtkRecentChooserMenuClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserMenuClass{native: c}

	return g
}

func (recv *RecentChooserMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserMenuPrivate
*/
type RecentChooserMenuPrivate struct {
	native *C.GtkRecentChooserMenuPrivate
}

func RecentChooserMenuPrivateNewFromC(u unsafe.Pointer) *RecentChooserMenuPrivate {
	c := (*C.GtkRecentChooserMenuPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserMenuPrivate{native: c}

	return g
}

func (recv *RecentChooserMenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserWidgetClass
*/
type RecentChooserWidgetClass struct {
	native *C.GtkRecentChooserWidgetClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func RecentChooserWidgetClassNewFromC(u unsafe.Pointer) *RecentChooserWidgetClass {
	c := (*C.GtkRecentChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserWidgetClass{native: c}

	return g
}

func (recv *RecentChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentChooserWidgetPrivate
*/
type RecentChooserWidgetPrivate struct {
	native *C.GtkRecentChooserWidgetPrivate
}

func RecentChooserWidgetPrivateNewFromC(u unsafe.Pointer) *RecentChooserWidgetPrivate {
	c := (*C.GtkRecentChooserWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserWidgetPrivate{native: c}

	return g
}

func (recv *RecentChooserWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
/*

C type

GtkRecentData
*/
type RecentData struct {
	native      *C.GtkRecentData
	DisplayName string
	Description string
	MimeType    string
	AppName     string
	AppExec     string
	// no type for groups
	IsPrivate bool
}

func RecentDataNewFromC(u unsafe.Pointer) *RecentData {
	c := (*C.GtkRecentData)(u)
	if c == nil {
		return nil
	}

	g := &RecentData{
		AppExec:     C.GoString(c.app_exec),
		AppName:     C.GoString(c.app_name),
		Description: C.GoString(c.description),
		DisplayName: C.GoString(c.display_name),
		IsPrivate:   c.is_private == C.TRUE,
		MimeType:    C.GoString(c.mime_type),
		native:      c,
	}

	return g
}

func (recv *RecentData) ToC() unsafe.Pointer {
	recv.native.display_name =
		C.CString(recv.DisplayName)
	recv.native.description =
		C.CString(recv.Description)
	recv.native.mime_type =
		C.CString(recv.MimeType)
	recv.native.app_name =
		C.CString(recv.AppName)
	recv.native.app_exec =
		C.CString(recv.AppExec)
	recv.native.is_private =
		boolToGboolean(recv.IsPrivate)

	return (unsafe.Pointer)(recv.native)
}

// A GtkRecentFilterInfo struct is used
// to pass information about the tested file to gtk_recent_filter_filter().
/*

C type

GtkRecentFilterInfo
*/
type RecentFilterInfo struct {
	native      *C.GtkRecentFilterInfo
	Contains    RecentFilterFlags
	Uri         string
	DisplayName string
	MimeType    string
	// no type for applications
	// no type for groups
	Age int32
}

func RecentFilterInfoNewFromC(u unsafe.Pointer) *RecentFilterInfo {
	c := (*C.GtkRecentFilterInfo)(u)
	if c == nil {
		return nil
	}

	g := &RecentFilterInfo{
		Age:         (int32)(c.age),
		Contains:    (RecentFilterFlags)(c.contains),
		DisplayName: C.GoString(c.display_name),
		MimeType:    C.GoString(c.mime_type),
		Uri:         C.GoString(c.uri),
		native:      c,
	}

	return g
}

func (recv *RecentFilterInfo) ToC() unsafe.Pointer {
	recv.native.contains =
		(C.GtkRecentFilterFlags)(recv.Contains)
	recv.native.uri =
		C.CString(recv.Uri)
	recv.native.display_name =
		C.CString(recv.DisplayName)
	recv.native.mime_type =
		C.CString(recv.MimeType)
	recv.native.age =
		(C.gint)(recv.Age)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRecentManagerPrivate
*/
type RecentManagerPrivate struct {
	native *C.GtkRecentManagerPrivate
}

func RecentManagerPrivateNewFromC(u unsafe.Pointer) *RecentManagerPrivate {
	c := (*C.GtkRecentManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RecentManagerPrivate{native: c}

	return g
}

func (recv *RecentManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRendererCellAccessibleClass
*/
type RendererCellAccessibleClass struct {
	native *C.GtkRendererCellAccessibleClass
	// parent_class : record
}

func RendererCellAccessibleClassNewFromC(u unsafe.Pointer) *RendererCellAccessibleClass {
	c := (*C.GtkRendererCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &RendererCellAccessibleClass{native: c}

	return g
}

func (recv *RendererCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkRendererCellAccessiblePrivate
*/
type RendererCellAccessiblePrivate struct {
	native *C.GtkRendererCellAccessiblePrivate
}

func RendererCellAccessiblePrivateNewFromC(u unsafe.Pointer) *RendererCellAccessiblePrivate {
	c := (*C.GtkRendererCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RendererCellAccessiblePrivate{native: c}

	return g
}

func (recv *RendererCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Represents a request of a screen object in a given orientation. These
// are primarily used in container implementations when allocating a natural
// size for children calling. See gtk_distribute_natural_allocation().
/*

C type

GtkRequestedSize
*/
type RequestedSize struct {
	native      *C.GtkRequestedSize
	Data        uintptr
	MinimumSize int32
	NaturalSize int32
}

func RequestedSizeNewFromC(u unsafe.Pointer) *RequestedSize {
	c := (*C.GtkRequestedSize)(u)
	if c == nil {
		return nil
	}

	g := &RequestedSize{
		Data:        (uintptr)(c.data),
		MinimumSize: (int32)(c.minimum_size),
		NaturalSize: (int32)(c.natural_size),
		native:      c,
	}

	return g
}

func (recv *RequestedSize) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)
	recv.native.minimum_size =
		(C.gint)(recv.MinimumSize)
	recv.native.natural_size =
		(C.gint)(recv.NaturalSize)

	return (unsafe.Pointer)(recv.native)
}

// A #GtkRequisition-struct represents the desired size of a widget. See
// [GtkWidget’s geometry management section][geometry-management] for
// more information.
/*

C type

GtkRequisition
*/
type Requisition struct {
	native *C.GtkRequisition
	Width  int32
	Height int32
}

func RequisitionNewFromC(u unsafe.Pointer) *Requisition {
	c := (*C.GtkRequisition)(u)
	if c == nil {
		return nil
	}

	g := &Requisition{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		native: c,
	}

	return g
}

func (recv *Requisition) ToC() unsafe.Pointer {
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Copies a #GtkRequisition.
/*

C function

gtk_requisition_copy
*/
func (recv *Requisition) Copy() *Requisition {
	retC := C.gtk_requisition_copy((*C.GtkRequisition)(recv.native))
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a #GtkRequisition.
/*

C function

gtk_requisition_free
*/
func (recv *Requisition) Free() {
	C.gtk_requisition_free((*C.GtkRequisition)(recv.native))

	return
}

/*

C type

GtkRevealerClass
*/
type RevealerClass struct {
	native *C.GtkRevealerClass
	// parent_class : record
}

func RevealerClassNewFromC(u unsafe.Pointer) *RevealerClass {
	c := (*C.GtkRevealerClass)(u)
	if c == nil {
		return nil
	}

	g := &RevealerClass{native: c}

	return g
}

func (recv *RevealerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleAccessibleClass
*/
type ScaleAccessibleClass struct {
	native *C.GtkScaleAccessibleClass
	// parent_class : record
}

func ScaleAccessibleClassNewFromC(u unsafe.Pointer) *ScaleAccessibleClass {
	c := (*C.GtkScaleAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ScaleAccessibleClass{native: c}

	return g
}

func (recv *ScaleAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleAccessiblePrivate
*/
type ScaleAccessiblePrivate struct {
	native *C.GtkScaleAccessiblePrivate
}

func ScaleAccessiblePrivateNewFromC(u unsafe.Pointer) *ScaleAccessiblePrivate {
	c := (*C.GtkScaleAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScaleAccessiblePrivate{native: c}

	return g
}

func (recv *ScaleAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleButtonAccessibleClass
*/
type ScaleButtonAccessibleClass struct {
	native *C.GtkScaleButtonAccessibleClass
	// parent_class : record
}

func ScaleButtonAccessibleClassNewFromC(u unsafe.Pointer) *ScaleButtonAccessibleClass {
	c := (*C.GtkScaleButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonAccessibleClass{native: c}

	return g
}

func (recv *ScaleButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleButtonAccessiblePrivate
*/
type ScaleButtonAccessiblePrivate struct {
	native *C.GtkScaleButtonAccessiblePrivate
}

func ScaleButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ScaleButtonAccessiblePrivate {
	c := (*C.GtkScaleButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonAccessiblePrivate{native: c}

	return g
}

func (recv *ScaleButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleButtonClass
*/
type ScaleButtonClass struct {
	native *C.GtkScaleButtonClass
	// parent_class : record
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScaleButtonClassNewFromC(u unsafe.Pointer) *ScaleButtonClass {
	c := (*C.GtkScaleButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonClass{native: c}

	return g
}

func (recv *ScaleButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleButtonPrivate
*/
type ScaleButtonPrivate struct {
	native *C.GtkScaleButtonPrivate
}

func ScaleButtonPrivateNewFromC(u unsafe.Pointer) *ScaleButtonPrivate {
	c := (*C.GtkScaleButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonPrivate{native: c}

	return g
}

func (recv *ScaleButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScaleClass
*/
type ScaleClass struct {
	native *C.GtkScaleClass
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
	c := (*C.GtkScaleClass)(u)
	if c == nil {
		return nil
	}

	g := &ScaleClass{native: c}

	return g
}

func (recv *ScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScalePrivate
*/
type ScalePrivate struct {
	native *C.GtkScalePrivate
}

func ScalePrivateNewFromC(u unsafe.Pointer) *ScalePrivate {
	c := (*C.GtkScalePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScalePrivate{native: c}

	return g
}

func (recv *ScalePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrollableInterface
*/
type ScrollableInterface struct {
	native *C.GtkScrollableInterface
	// base_iface : record
	// no type for get_border
}

func ScrollableInterfaceNewFromC(u unsafe.Pointer) *ScrollableInterface {
	c := (*C.GtkScrollableInterface)(u)
	if c == nil {
		return nil
	}

	g := &ScrollableInterface{native: c}

	return g
}

func (recv *ScrollableInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrollbarClass
*/
type ScrollbarClass struct {
	native *C.GtkScrollbarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ScrollbarClassNewFromC(u unsafe.Pointer) *ScrollbarClass {
	c := (*C.GtkScrollbarClass)(u)
	if c == nil {
		return nil
	}

	g := &ScrollbarClass{native: c}

	return g
}

func (recv *ScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrolledWindowAccessibleClass
*/
type ScrolledWindowAccessibleClass struct {
	native *C.GtkScrolledWindowAccessibleClass
	// parent_class : record
}

func ScrolledWindowAccessibleClassNewFromC(u unsafe.Pointer) *ScrolledWindowAccessibleClass {
	c := (*C.GtkScrolledWindowAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowAccessibleClass{native: c}

	return g
}

func (recv *ScrolledWindowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrolledWindowAccessiblePrivate
*/
type ScrolledWindowAccessiblePrivate struct {
	native *C.GtkScrolledWindowAccessiblePrivate
}

func ScrolledWindowAccessiblePrivateNewFromC(u unsafe.Pointer) *ScrolledWindowAccessiblePrivate {
	c := (*C.GtkScrolledWindowAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowAccessiblePrivate{native: c}

	return g
}

func (recv *ScrolledWindowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrolledWindowClass
*/
type ScrolledWindowClass struct {
	native *C.GtkScrolledWindowClass
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
	c := (*C.GtkScrolledWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowClass{
		ScrollbarSpacing: (int32)(c.scrollbar_spacing),
		native:           c,
	}

	return g
}

func (recv *ScrolledWindowClass) ToC() unsafe.Pointer {
	recv.native.scrollbar_spacing =
		(C.gint)(recv.ScrollbarSpacing)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkScrolledWindowPrivate
*/
type ScrolledWindowPrivate struct {
	native *C.GtkScrolledWindowPrivate
}

func ScrolledWindowPrivateNewFromC(u unsafe.Pointer) *ScrolledWindowPrivate {
	c := (*C.GtkScrolledWindowPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowPrivate{native: c}

	return g
}

func (recv *ScrolledWindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSearchBarClass
*/
type SearchBarClass struct {
	native *C.GtkSearchBarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SearchBarClassNewFromC(u unsafe.Pointer) *SearchBarClass {
	c := (*C.GtkSearchBarClass)(u)
	if c == nil {
		return nil
	}

	g := &SearchBarClass{native: c}

	return g
}

func (recv *SearchBarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSearchEntryClass
*/
type SearchEntryClass struct {
	native *C.GtkSearchEntryClass
	// parent_class : record
	// no type for search_changed
	// no type for next_match
	// no type for previous_match
	// no type for stop_search
}

func SearchEntryClassNewFromC(u unsafe.Pointer) *SearchEntryClass {
	c := (*C.GtkSearchEntryClass)(u)
	if c == nil {
		return nil
	}

	g := &SearchEntryClass{native: c}

	return g
}

func (recv *SearchEntryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSelectionData
*/
type SelectionData struct {
	native *C.GtkSelectionData
}

func SelectionDataNewFromC(u unsafe.Pointer) *SelectionData {
	c := (*C.GtkSelectionData)(u)
	if c == nil {
		return nil
	}

	g := &SelectionData{native: c}

	return g
}

func (recv *SelectionData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Makes a copy of a #GtkSelectionData-struct and its data.
/*

C function

gtk_selection_data_copy
*/
func (recv *SelectionData) Copy() *SelectionData {
	retC := C.gtk_selection_data_copy((*C.GtkSelectionData)(recv.native))
	retGo := SelectionDataNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a #GtkSelectionData-struct returned from
// gtk_selection_data_copy().
/*

C function

gtk_selection_data_free
*/
func (recv *SelectionData) Free() {
	C.gtk_selection_data_free((*C.GtkSelectionData)(recv.native))

	return
}

// Unsupported : gtk_selection_data_get_targets : unsupported parameter targets : output array param targets

// Blacklisted : gtk_selection_data_get_text

// Unsupported : gtk_selection_data_set : unsupported parameter type : Blacklisted record : GdkAtom

// Sets the contents of the selection from a UTF-8 encoded string.
// The string is converted to the form determined by
// @selection_data->target.
/*

C function

gtk_selection_data_set_text
*/
func (recv *SelectionData) SetText(str string, len int32) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gint)(len)

	retC := C.gtk_selection_data_set_text((*C.GtkSelectionData)(recv.native), c_str, c_len)
	retGo := retC == C.TRUE

	return retGo
}

// Given a #GtkSelectionData object holding a list of targets,
// determines if any of the targets in @targets can be used to
// provide text.
/*

C function

gtk_selection_data_targets_include_text
*/
func (recv *SelectionData) TargetsIncludeText() bool {
	retC := C.gtk_selection_data_targets_include_text((*C.GtkSelectionData)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

/*

C type

GtkSeparatorClass
*/
type SeparatorClass struct {
	native *C.GtkSeparatorClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorClassNewFromC(u unsafe.Pointer) *SeparatorClass {
	c := (*C.GtkSeparatorClass)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorClass{native: c}

	return g
}

func (recv *SeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSeparatorMenuItemClass
*/
type SeparatorMenuItemClass struct {
	native *C.GtkSeparatorMenuItemClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorMenuItemClassNewFromC(u unsafe.Pointer) *SeparatorMenuItemClass {
	c := (*C.GtkSeparatorMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorMenuItemClass{native: c}

	return g
}

func (recv *SeparatorMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSeparatorPrivate
*/
type SeparatorPrivate struct {
	native *C.GtkSeparatorPrivate
}

func SeparatorPrivateNewFromC(u unsafe.Pointer) *SeparatorPrivate {
	c := (*C.GtkSeparatorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorPrivate{native: c}

	return g
}

func (recv *SeparatorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSeparatorToolItemClass
*/
type SeparatorToolItemClass struct {
	native *C.GtkSeparatorToolItemClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SeparatorToolItemClassNewFromC(u unsafe.Pointer) *SeparatorToolItemClass {
	c := (*C.GtkSeparatorToolItemClass)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorToolItemClass{native: c}

	return g
}

func (recv *SeparatorToolItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSeparatorToolItemPrivate
*/
type SeparatorToolItemPrivate struct {
	native *C.GtkSeparatorToolItemPrivate
}

func SeparatorToolItemPrivateNewFromC(u unsafe.Pointer) *SeparatorToolItemPrivate {
	c := (*C.GtkSeparatorToolItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorToolItemPrivate{native: c}

	return g
}

func (recv *SeparatorToolItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSettingsClass
*/
type SettingsClass struct {
	native *C.GtkSettingsClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	c := (*C.GtkSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SettingsClass{native: c}

	return g
}

func (recv *SettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSettingsPrivate
*/
type SettingsPrivate struct {
	native *C.GtkSettingsPrivate
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	c := (*C.GtkSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SettingsPrivate{native: c}

	return g
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSettingsValue
*/
type SettingsValue struct {
	native *C.GtkSettingsValue
	Origin string
	// value : record
}

func SettingsValueNewFromC(u unsafe.Pointer) *SettingsValue {
	c := (*C.GtkSettingsValue)(u)
	if c == nil {
		return nil
	}

	g := &SettingsValue{
		Origin: C.GoString(c.origin),
		native: c,
	}

	return g
}

func (recv *SettingsValue) ToC() unsafe.Pointer {
	recv.native.origin =
		C.CString(recv.Origin)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSizeGroupClass
*/
type SizeGroupClass struct {
	native *C.GtkSizeGroupClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SizeGroupClassNewFromC(u unsafe.Pointer) *SizeGroupClass {
	c := (*C.GtkSizeGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &SizeGroupClass{native: c}

	return g
}

func (recv *SizeGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSizeGroupPrivate
*/
type SizeGroupPrivate struct {
	native *C.GtkSizeGroupPrivate
}

func SizeGroupPrivateNewFromC(u unsafe.Pointer) *SizeGroupPrivate {
	c := (*C.GtkSizeGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SizeGroupPrivate{native: c}

	return g
}

func (recv *SizeGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSocketClass
*/
type SocketClass struct {
	native *C.GtkSocketClass
	// parent_class : record
	// no type for plug_added
	// no type for plug_removed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.GtkSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSocketPrivate
*/
type SocketPrivate struct {
	native *C.GtkSocketPrivate
}

func SocketPrivateNewFromC(u unsafe.Pointer) *SocketPrivate {
	c := (*C.GtkSocketPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketPrivate{native: c}

	return g
}

func (recv *SocketPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinButtonAccessibleClass
*/
type SpinButtonAccessibleClass struct {
	native *C.GtkSpinButtonAccessibleClass
	// parent_class : record
}

func SpinButtonAccessibleClassNewFromC(u unsafe.Pointer) *SpinButtonAccessibleClass {
	c := (*C.GtkSpinButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonAccessibleClass{native: c}

	return g
}

func (recv *SpinButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinButtonAccessiblePrivate
*/
type SpinButtonAccessiblePrivate struct {
	native *C.GtkSpinButtonAccessiblePrivate
}

func SpinButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *SpinButtonAccessiblePrivate {
	c := (*C.GtkSpinButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonAccessiblePrivate{native: c}

	return g
}

func (recv *SpinButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinButtonClass
*/
type SpinButtonClass struct {
	native *C.GtkSpinButtonClass
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
	c := (*C.GtkSpinButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonClass{native: c}

	return g
}

func (recv *SpinButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinButtonPrivate
*/
type SpinButtonPrivate struct {
	native *C.GtkSpinButtonPrivate
}

func SpinButtonPrivateNewFromC(u unsafe.Pointer) *SpinButtonPrivate {
	c := (*C.GtkSpinButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonPrivate{native: c}

	return g
}

func (recv *SpinButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinnerAccessibleClass
*/
type SpinnerAccessibleClass struct {
	native *C.GtkSpinnerAccessibleClass
	// parent_class : record
}

func SpinnerAccessibleClassNewFromC(u unsafe.Pointer) *SpinnerAccessibleClass {
	c := (*C.GtkSpinnerAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerAccessibleClass{native: c}

	return g
}

func (recv *SpinnerAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinnerAccessiblePrivate
*/
type SpinnerAccessiblePrivate struct {
	native *C.GtkSpinnerAccessiblePrivate
}

func SpinnerAccessiblePrivateNewFromC(u unsafe.Pointer) *SpinnerAccessiblePrivate {
	c := (*C.GtkSpinnerAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerAccessiblePrivate{native: c}

	return g
}

func (recv *SpinnerAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinnerClass
*/
type SpinnerClass struct {
	native *C.GtkSpinnerClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func SpinnerClassNewFromC(u unsafe.Pointer) *SpinnerClass {
	c := (*C.GtkSpinnerClass)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerClass{native: c}

	return g
}

func (recv *SpinnerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSpinnerPrivate
*/
type SpinnerPrivate struct {
	native *C.GtkSpinnerPrivate
}

func SpinnerPrivateNewFromC(u unsafe.Pointer) *SpinnerPrivate {
	c := (*C.GtkSpinnerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerPrivate{native: c}

	return g
}

func (recv *SpinnerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GtkStackAccessibleClass

/*

C type

GtkStackClass
*/
type StackClass struct {
	native *C.GtkStackClass
	// parent_class : record
}

func StackClassNewFromC(u unsafe.Pointer) *StackClass {
	c := (*C.GtkStackClass)(u)
	if c == nil {
		return nil
	}

	g := &StackClass{native: c}

	return g
}

func (recv *StackClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStackSidebarClass
*/
type StackSidebarClass struct {
	native *C.GtkStackSidebarClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StackSidebarClassNewFromC(u unsafe.Pointer) *StackSidebarClass {
	c := (*C.GtkStackSidebarClass)(u)
	if c == nil {
		return nil
	}

	g := &StackSidebarClass{native: c}

	return g
}

func (recv *StackSidebarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStackSidebarPrivate
*/
type StackSidebarPrivate struct {
	native *C.GtkStackSidebarPrivate
}

func StackSidebarPrivateNewFromC(u unsafe.Pointer) *StackSidebarPrivate {
	c := (*C.GtkStackSidebarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StackSidebarPrivate{native: c}

	return g
}

func (recv *StackSidebarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStackSwitcherClass
*/
type StackSwitcherClass struct {
	native *C.GtkStackSwitcherClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StackSwitcherClassNewFromC(u unsafe.Pointer) *StackSwitcherClass {
	c := (*C.GtkStackSwitcherClass)(u)
	if c == nil {
		return nil
	}

	g := &StackSwitcherClass{native: c}

	return g
}

func (recv *StackSwitcherClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusIconClass
*/
type StatusIconClass struct {
	native *C.GtkStatusIconClass
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
	c := (*C.GtkStatusIconClass)(u)
	if c == nil {
		return nil
	}

	g := &StatusIconClass{native: c}

	return g
}

func (recv *StatusIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusIconPrivate
*/
type StatusIconPrivate struct {
	native *C.GtkStatusIconPrivate
}

func StatusIconPrivateNewFromC(u unsafe.Pointer) *StatusIconPrivate {
	c := (*C.GtkStatusIconPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StatusIconPrivate{native: c}

	return g
}

func (recv *StatusIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusbarAccessibleClass
*/
type StatusbarAccessibleClass struct {
	native *C.GtkStatusbarAccessibleClass
	// parent_class : record
}

func StatusbarAccessibleClassNewFromC(u unsafe.Pointer) *StatusbarAccessibleClass {
	c := (*C.GtkStatusbarAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarAccessibleClass{native: c}

	return g
}

func (recv *StatusbarAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusbarAccessiblePrivate
*/
type StatusbarAccessiblePrivate struct {
	native *C.GtkStatusbarAccessiblePrivate
}

func StatusbarAccessiblePrivateNewFromC(u unsafe.Pointer) *StatusbarAccessiblePrivate {
	c := (*C.GtkStatusbarAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarAccessiblePrivate{native: c}

	return g
}

func (recv *StatusbarAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusbarClass
*/
type StatusbarClass struct {
	native *C.GtkStatusbarClass
	// parent_class : record
	Reserved uintptr
	// no type for text_pushed
	// no type for text_popped
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StatusbarClassNewFromC(u unsafe.Pointer) *StatusbarClass {
	c := (*C.GtkStatusbarClass)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarClass{
		Reserved: (uintptr)(c.reserved),
		native:   c,
	}

	return g
}

func (recv *StatusbarClass) ToC() unsafe.Pointer {
	recv.native.reserved =
		(C.gpointer)(recv.Reserved)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStatusbarPrivate
*/
type StatusbarPrivate struct {
	native *C.GtkStatusbarPrivate
}

func StatusbarPrivateNewFromC(u unsafe.Pointer) *StatusbarPrivate {
	c := (*C.GtkStatusbarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarPrivate{native: c}

	return g
}

func (recv *StatusbarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStockItem
*/
type StockItem struct {
	native            *C.GtkStockItem
	StockId           string
	Label             string
	Modifier          gdk.ModifierType
	Keyval            uint32
	TranslationDomain string
}

func StockItemNewFromC(u unsafe.Pointer) *StockItem {
	c := (*C.GtkStockItem)(u)
	if c == nil {
		return nil
	}

	g := &StockItem{
		Keyval:            (uint32)(c.keyval),
		Label:             C.GoString(c.label),
		Modifier:          (gdk.ModifierType)(c.modifier),
		StockId:           C.GoString(c.stock_id),
		TranslationDomain: C.GoString(c.translation_domain),
		native:            c,
	}

	return g
}

func (recv *StockItem) ToC() unsafe.Pointer {
	recv.native.stock_id =
		C.CString(recv.StockId)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.modifier =
		(C.GdkModifierType)(recv.Modifier)
	recv.native.keyval =
		(C.guint)(recv.Keyval)
	recv.native.translation_domain =
		C.CString(recv.TranslationDomain)

	return (unsafe.Pointer)(recv.native)
}

// Copies a stock item, mostly useful for language bindings and not in applications.
/*

C function

gtk_stock_item_copy
*/
func (recv *StockItem) Copy() *StockItem {
	retC := C.gtk_stock_item_copy((*C.GtkStockItem)(recv.native))
	retGo := StockItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a stock item allocated on the heap, such as one returned by
// gtk_stock_item_copy(). Also frees the fields inside the stock item,
// if they are not %NULL.
/*

C function

gtk_stock_item_free
*/
func (recv *StockItem) Free() {
	C.gtk_stock_item_free((*C.GtkStockItem)(recv.native))

	return
}

/*

C type

GtkStyleClass
*/
type StyleClass struct {
	native *C.GtkStyleClass
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
	c := (*C.GtkStyleClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleClass{native: c}

	return g
}

func (recv *StyleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStyleContextClass
*/
type StyleContextClass struct {
	native *C.GtkStyleContextClass
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StyleContextClassNewFromC(u unsafe.Pointer) *StyleContextClass {
	c := (*C.GtkStyleContextClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleContextClass{native: c}

	return g
}

func (recv *StyleContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStyleContextPrivate
*/
type StyleContextPrivate struct {
	native *C.GtkStyleContextPrivate
}

func StyleContextPrivateNewFromC(u unsafe.Pointer) *StyleContextPrivate {
	c := (*C.GtkStyleContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StyleContextPrivate{native: c}

	return g
}

func (recv *StyleContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStylePropertiesClass
*/
type StylePropertiesClass struct {
	native *C.GtkStylePropertiesClass
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func StylePropertiesClassNewFromC(u unsafe.Pointer) *StylePropertiesClass {
	c := (*C.GtkStylePropertiesClass)(u)
	if c == nil {
		return nil
	}

	g := &StylePropertiesClass{native: c}

	return g
}

func (recv *StylePropertiesClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStylePropertiesPrivate
*/
type StylePropertiesPrivate struct {
	native *C.GtkStylePropertiesPrivate
}

func StylePropertiesPrivateNewFromC(u unsafe.Pointer) *StylePropertiesPrivate {
	c := (*C.GtkStylePropertiesPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StylePropertiesPrivate{native: c}

	return g
}

func (recv *StylePropertiesPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkStyleProviderIface
*/
type StyleProviderIface struct {
	native *C.GtkStyleProviderIface
	// Private : g_iface
	// no type for get_style
	// no type for get_style_property
	// no type for get_icon_factory
}

func StyleProviderIfaceNewFromC(u unsafe.Pointer) *StyleProviderIface {
	c := (*C.GtkStyleProviderIface)(u)
	if c == nil {
		return nil
	}

	g := &StyleProviderIface{native: c}

	return g
}

func (recv *StyleProviderIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSwitchAccessibleClass
*/
type SwitchAccessibleClass struct {
	native *C.GtkSwitchAccessibleClass
	// parent_class : record
}

func SwitchAccessibleClassNewFromC(u unsafe.Pointer) *SwitchAccessibleClass {
	c := (*C.GtkSwitchAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &SwitchAccessibleClass{native: c}

	return g
}

func (recv *SwitchAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSwitchAccessiblePrivate
*/
type SwitchAccessiblePrivate struct {
	native *C.GtkSwitchAccessiblePrivate
}

func SwitchAccessiblePrivateNewFromC(u unsafe.Pointer) *SwitchAccessiblePrivate {
	c := (*C.GtkSwitchAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SwitchAccessiblePrivate{native: c}

	return g
}

func (recv *SwitchAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSwitchClass
*/
type SwitchClass struct {
	native *C.GtkSwitchClass
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
	c := (*C.GtkSwitchClass)(u)
	if c == nil {
		return nil
	}

	g := &SwitchClass{native: c}

	return g
}

func (recv *SwitchClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkSwitchPrivate
*/
type SwitchPrivate struct {
	native *C.GtkSwitchPrivate
}

func SwitchPrivateNewFromC(u unsafe.Pointer) *SwitchPrivate {
	c := (*C.GtkSwitchPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SwitchPrivate{native: c}

	return g
}

func (recv *SwitchPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkSymbolicColor is a boxed type that represents a symbolic color.
// It is the result of parsing a
// [color expression][gtkcssprovider-symbolic-colors].
// To obtain the color represented by a GtkSymbolicColor, it has to
// be resolved with gtk_symbolic_color_resolve(), which replaces all
// symbolic color references by the colors they refer to (in a given
// context) and evaluates mix, shade and other expressions, resulting
// in a #GdkRGBA value.
//
// It is not normally necessary to deal directly with #GtkSymbolicColors,
// since they are mostly used behind the scenes by #GtkStyleContext and
// #GtkCssProvider.
//
// #GtkSymbolicColor is deprecated. Symbolic colors are considered an
// implementation detail of GTK+.
/*

C type

GtkSymbolicColor
*/
type SymbolicColor struct {
	native *C.GtkSymbolicColor
}

func SymbolicColorNewFromC(u unsafe.Pointer) *SymbolicColor {
	c := (*C.GtkSymbolicColor)(u)
	if c == nil {
		return nil
	}

	g := &SymbolicColor{native: c}

	return g
}

func (recv *SymbolicColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Converts the given @color to a string representation. This is useful
// both for debugging and for serialization of strings. The format of
// the string may change between different versions of GTK, but it is
// guaranteed that the GTK css parser is able to read the string and
// create the same symbolic color from it.
/*

C function

gtk_symbolic_color_to_string
*/
func (recv *SymbolicColor) ToString() string {
	retC := C.gtk_symbolic_color_to_string((*C.GtkSymbolicColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

/*

C type

GtkTableChild
*/
type TableChild struct {
	native *C.GtkTableChild
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
	c := (*C.GtkTableChild)(u)
	if c == nil {
		return nil
	}

	g := &TableChild{
		BottomAttach: (uint16)(c.bottom_attach),
		LeftAttach:   (uint16)(c.left_attach),
		RightAttach:  (uint16)(c.right_attach),
		TopAttach:    (uint16)(c.top_attach),
		Xpadding:     (uint16)(c.xpadding),
		Ypadding:     (uint16)(c.ypadding),
		native:       c,
	}

	return g
}

func (recv *TableChild) ToC() unsafe.Pointer {
	recv.native.left_attach =
		(C.guint16)(recv.LeftAttach)
	recv.native.right_attach =
		(C.guint16)(recv.RightAttach)
	recv.native.top_attach =
		(C.guint16)(recv.TopAttach)
	recv.native.bottom_attach =
		(C.guint16)(recv.BottomAttach)
	recv.native.xpadding =
		(C.guint16)(recv.Xpadding)
	recv.native.ypadding =
		(C.guint16)(recv.Ypadding)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTableClass
*/
type TableClass struct {
	native *C.GtkTableClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TableClassNewFromC(u unsafe.Pointer) *TableClass {
	c := (*C.GtkTableClass)(u)
	if c == nil {
		return nil
	}

	g := &TableClass{native: c}

	return g
}

func (recv *TableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTablePrivate
*/
type TablePrivate struct {
	native *C.GtkTablePrivate
}

func TablePrivateNewFromC(u unsafe.Pointer) *TablePrivate {
	c := (*C.GtkTablePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TablePrivate{native: c}

	return g
}

func (recv *TablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTableRowCol
*/
type TableRowCol struct {
	native      *C.GtkTableRowCol
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
	c := (*C.GtkTableRowCol)(u)
	if c == nil {
		return nil
	}

	g := &TableRowCol{
		Allocation:  (uint16)(c.allocation),
		Requisition: (uint16)(c.requisition),
		Spacing:     (uint16)(c.spacing),
		native:      c,
	}

	return g
}

func (recv *TableRowCol) ToC() unsafe.Pointer {
	recv.native.requisition =
		(C.guint16)(recv.Requisition)
	recv.native.allocation =
		(C.guint16)(recv.Allocation)
	recv.native.spacing =
		(C.guint16)(recv.Spacing)

	return (unsafe.Pointer)(recv.native)
}

// A #GtkTargetEntry represents a single type of
// data than can be supplied for by a widget for a selection
// or for supplied or received during drag-and-drop.
/*

C type

GtkTargetEntry
*/
type TargetEntry struct {
	native *C.GtkTargetEntry
	Target string
	Flags  uint32
	Info   uint32
}

func TargetEntryNewFromC(u unsafe.Pointer) *TargetEntry {
	c := (*C.GtkTargetEntry)(u)
	if c == nil {
		return nil
	}

	g := &TargetEntry{
		Flags:  (uint32)(c.flags),
		Info:   (uint32)(c.info),
		Target: C.GoString(c.target),
		native: c,
	}

	return g
}

func (recv *TargetEntry) ToC() unsafe.Pointer {
	recv.native.target =
		C.CString(recv.Target)
	recv.native.flags =
		(C.guint)(recv.Flags)
	recv.native.info =
		(C.guint)(recv.Info)

	return (unsafe.Pointer)(recv.native)
}

// Makes a new #GtkTargetEntry.
/*

C function

gtk_target_entry_new
*/
func TargetEntryNew(target string, flags uint32, info uint32) *TargetEntry {
	c_target := C.CString(target)
	defer C.free(unsafe.Pointer(c_target))

	c_flags := (C.guint)(flags)

	c_info := (C.guint)(info)

	retC := C.gtk_target_entry_new(c_target, c_flags, c_info)
	retGo := TargetEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Makes a copy of a #GtkTargetEntry and its data.
/*

C function

gtk_target_entry_copy
*/
func (recv *TargetEntry) Copy() *TargetEntry {
	retC := C.gtk_target_entry_copy((*C.GtkTargetEntry)(recv.native))
	retGo := TargetEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a #GtkTargetEntry returned from
// gtk_target_entry_new() or gtk_target_entry_copy().
/*

C function

gtk_target_entry_free
*/
func (recv *TargetEntry) Free() {
	C.gtk_target_entry_free((*C.GtkTargetEntry)(recv.native))

	return
}

// A #GtkTargetList-struct is a reference counted list
// of #GtkTargetPair and should be treated as
// opaque.
/*

C type

GtkTargetList
*/
type TargetList struct {
	native *C.GtkTargetList
}

func TargetListNewFromC(u unsafe.Pointer) *TargetList {
	c := (*C.GtkTargetList)(u)
	if c == nil {
		return nil
	}

	g := &TargetList{native: c}

	return g
}

func (recv *TargetList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_target_list_new : unsupported parameter targets :

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_add_table : unsupported parameter targets :

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Increases the reference count of a #GtkTargetList by one.
/*

C function

gtk_target_list_ref
*/
func (recv *TargetList) Ref() *TargetList {
	retC := C.gtk_target_list_ref((*C.GtkTargetList)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

// Decreases the reference count of a #GtkTargetList by one.
// If the resulting reference count is zero, frees the list.
/*

C function

gtk_target_list_unref
*/
func (recv *TargetList) Unref() {
	C.gtk_target_list_unref((*C.GtkTargetList)(recv.native))

	return
}

// A #GtkTargetPair is used to represent the same
// information as a table of #GtkTargetEntry, but in
// an efficient form.
/*

C type

GtkTargetPair
*/
type TargetPair struct {
	native *C.GtkTargetPair
	// target : record
	Flags uint32
	Info  uint32
}

func TargetPairNewFromC(u unsafe.Pointer) *TargetPair {
	c := (*C.GtkTargetPair)(u)
	if c == nil {
		return nil
	}

	g := &TargetPair{
		Flags:  (uint32)(c.flags),
		Info:   (uint32)(c.info),
		native: c,
	}

	return g
}

func (recv *TargetPair) ToC() unsafe.Pointer {
	recv.native.flags =
		(C.guint)(recv.Flags)
	recv.native.info =
		(C.guint)(recv.Info)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTearoffMenuItemClass
*/
type TearoffMenuItemClass struct {
	native *C.GtkTearoffMenuItemClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TearoffMenuItemClassNewFromC(u unsafe.Pointer) *TearoffMenuItemClass {
	c := (*C.GtkTearoffMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &TearoffMenuItemClass{native: c}

	return g
}

func (recv *TearoffMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTearoffMenuItemPrivate
*/
type TearoffMenuItemPrivate struct {
	native *C.GtkTearoffMenuItemPrivate
}

func TearoffMenuItemPrivateNewFromC(u unsafe.Pointer) *TearoffMenuItemPrivate {
	c := (*C.GtkTearoffMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TearoffMenuItemPrivate{native: c}

	return g
}

func (recv *TearoffMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextAppearance
*/
type TextAppearance struct {
	native *C.GtkTextAppearance
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
	c := (*C.GtkTextAppearance)(u)
	if c == nil {
		return nil
	}

	g := &TextAppearance{
		Rise:   (int32)(c.rise),
		native: c,
	}

	return g
}

func (recv *TextAppearance) ToC() unsafe.Pointer {
	recv.native.rise =
		(C.gint)(recv.Rise)

	return (unsafe.Pointer)(recv.native)
}

// Using #GtkTextAttributes directly should rarely be necessary.
// It’s primarily useful with gtk_text_iter_get_attributes().
// As with most GTK+ structs, the fields in this struct should only
// be read, never modified directly.
/*

C type

GtkTextAttributes
*/
type TextAttributes struct {
	native *C.GtkTextAttributes
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
	c := (*C.GtkTextAttributes)(u)
	if c == nil {
		return nil
	}

	g := &TextAttributes{
		Direction:        (TextDirection)(c.direction),
		FontScale:        (float64)(c.font_scale),
		Indent:           (int32)(c.indent),
		Justification:    (Justification)(c.justification),
		LeftMargin:       (int32)(c.left_margin),
		LetterSpacing:    (int32)(c.letter_spacing),
		PixelsAboveLines: (int32)(c.pixels_above_lines),
		PixelsBelowLines: (int32)(c.pixels_below_lines),
		PixelsInsideWrap: (int32)(c.pixels_inside_wrap),
		RightMargin:      (int32)(c.right_margin),
		WrapMode:         (WrapMode)(c.wrap_mode),
		native:           c,
	}

	return g
}

func (recv *TextAttributes) ToC() unsafe.Pointer {
	recv.native.justification =
		(C.GtkJustification)(recv.Justification)
	recv.native.direction =
		(C.GtkTextDirection)(recv.Direction)
	recv.native.font_scale =
		(C.gdouble)(recv.FontScale)
	recv.native.left_margin =
		(C.gint)(recv.LeftMargin)
	recv.native.right_margin =
		(C.gint)(recv.RightMargin)
	recv.native.indent =
		(C.gint)(recv.Indent)
	recv.native.pixels_above_lines =
		(C.gint)(recv.PixelsAboveLines)
	recv.native.pixels_below_lines =
		(C.gint)(recv.PixelsBelowLines)
	recv.native.pixels_inside_wrap =
		(C.gint)(recv.PixelsInsideWrap)
	recv.native.wrap_mode =
		(C.GtkWrapMode)(recv.WrapMode)
	recv.native.letter_spacing =
		(C.gint)(recv.LetterSpacing)

	return (unsafe.Pointer)(recv.native)
}

// Creates a #GtkTextAttributes, which describes
// a set of properties on some text.
/*

C function

gtk_text_attributes_new
*/
func TextAttributesNew() *TextAttributes {
	retC := C.gtk_text_attributes_new()
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies @src and returns a new #GtkTextAttributes.
/*

C function

gtk_text_attributes_copy
*/
func (recv *TextAttributes) Copy() *TextAttributes {
	retC := C.gtk_text_attributes_copy((*C.GtkTextAttributes)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies the values from @src to @dest so that @dest has
// the same values as @src. Frees existing values in @dest.
/*

C function

gtk_text_attributes_copy_values
*/
func (recv *TextAttributes) CopyValues(dest *TextAttributes) {
	c_dest := (*C.GtkTextAttributes)(C.NULL)
	if dest != nil {
		c_dest = (*C.GtkTextAttributes)(dest.ToC())
	}

	C.gtk_text_attributes_copy_values((*C.GtkTextAttributes)(recv.native), c_dest)

	return
}

// Increments the reference count on @values.
/*

C function

gtk_text_attributes_ref
*/
func (recv *TextAttributes) Ref() *TextAttributes {
	retC := C.gtk_text_attributes_ref((*C.GtkTextAttributes)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count on @values, freeing the structure
// if the reference count reaches 0.
/*

C function

gtk_text_attributes_unref
*/
func (recv *TextAttributes) Unref() {
	C.gtk_text_attributes_unref((*C.GtkTextAttributes)(recv.native))

	return
}

/*

C type

GtkTextBTree
*/
type TextBTree struct {
	native *C.GtkTextBTree
}

func TextBTreeNewFromC(u unsafe.Pointer) *TextBTree {
	c := (*C.GtkTextBTree)(u)
	if c == nil {
		return nil
	}

	g := &TextBTree{native: c}

	return g
}

func (recv *TextBTree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextBufferClass
*/
type TextBufferClass struct {
	native *C.GtkTextBufferClass
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
	c := (*C.GtkTextBufferClass)(u)
	if c == nil {
		return nil
	}

	g := &TextBufferClass{native: c}

	return g
}

func (recv *TextBufferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextBufferPrivate
*/
type TextBufferPrivate struct {
	native *C.GtkTextBufferPrivate
}

func TextBufferPrivateNewFromC(u unsafe.Pointer) *TextBufferPrivate {
	c := (*C.GtkTextBufferPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextBufferPrivate{native: c}

	return g
}

func (recv *TextBufferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextCellAccessibleClass
*/
type TextCellAccessibleClass struct {
	native *C.GtkTextCellAccessibleClass
	// parent_class : record
}

func TextCellAccessibleClassNewFromC(u unsafe.Pointer) *TextCellAccessibleClass {
	c := (*C.GtkTextCellAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &TextCellAccessibleClass{native: c}

	return g
}

func (recv *TextCellAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextCellAccessiblePrivate
*/
type TextCellAccessiblePrivate struct {
	native *C.GtkTextCellAccessiblePrivate
}

func TextCellAccessiblePrivateNewFromC(u unsafe.Pointer) *TextCellAccessiblePrivate {
	c := (*C.GtkTextCellAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextCellAccessiblePrivate{native: c}

	return g
}

func (recv *TextCellAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextChildAnchorClass
*/
type TextChildAnchorClass struct {
	native *C.GtkTextChildAnchorClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextChildAnchorClassNewFromC(u unsafe.Pointer) *TextChildAnchorClass {
	c := (*C.GtkTextChildAnchorClass)(u)
	if c == nil {
		return nil
	}

	g := &TextChildAnchorClass{native: c}

	return g
}

func (recv *TextChildAnchorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// You may wish to begin by reading the
// [text widget conceptual overview][TextWidget]
// which gives an overview of all the objects and data
// types related to the text widget and how they work together.
/*

C type

GtkTextIter
*/
type TextIter struct {
	native *C.GtkTextIter
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
	c := (*C.GtkTextIter)(u)
	if c == nil {
		return nil
	}

	g := &TextIter{native: c}

	return g
}

func (recv *TextIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Moves backward by one character offset. Returns %TRUE if movement
// was possible; if @iter was the first in the buffer (character
// offset 0), gtk_text_iter_backward_char() returns %FALSE for convenience when
// writing loops.
/*

C function

gtk_text_iter_backward_char
*/
func (recv *TextIter) BackwardChar() bool {
	retC := C.gtk_text_iter_backward_char((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count characters backward, if possible (if @count would move
// past the start or end of the buffer, moves to the start or end of
// the buffer).  The return value indicates whether the iterator moved
// onto a dereferenceable position; if the iterator didn’t move, or
// moved onto the end iterator, then %FALSE is returned. If @count is 0,
// the function does nothing and returns %FALSE.
/*

C function

gtk_text_iter_backward_chars
*/
func (recv *TextIter) BackwardChars(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_chars((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Like gtk_text_iter_forward_cursor_position(), but moves backward.
/*

C function

gtk_text_iter_backward_cursor_position
*/
func (recv *TextIter) BackwardCursorPosition() bool {
	retC := C.gtk_text_iter_backward_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves up to @count cursor positions. See
// gtk_text_iter_forward_cursor_position() for details.
/*

C function

gtk_text_iter_backward_cursor_positions
*/
func (recv *TextIter) BackwardCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_iter_backward_find_char : unsupported parameter pred : no type generator for TextCharPredicate (GtkTextCharPredicate) for param pred

// Moves @iter to the start of the previous line. Returns %TRUE if
// @iter could be moved; i.e. if @iter was at character offset 0, this
// function returns %FALSE. Therefore if @iter was already on line 0,
// but not at the start of the line, @iter is snapped to the start of
// the line and the function returns %TRUE. (Note that this implies that
// in a loop calling this function, the line number may not change on
// every iteration, if your first iteration is on line 0.)
/*

C function

gtk_text_iter_backward_line
*/
func (recv *TextIter) BackwardLine() bool {
	retC := C.gtk_text_iter_backward_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count lines backward, if possible (if @count would move
// past the start or end of the buffer, moves to the start or end of
// the buffer).  The return value indicates whether the iterator moved
// onto a dereferenceable position; if the iterator didn’t move, or
// moved onto the end iterator, then %FALSE is returned. If @count is 0,
// the function does nothing and returns %FALSE. If @count is negative,
// moves forward by 0 - @count lines.
/*

C function

gtk_text_iter_backward_lines
*/
func (recv *TextIter) BackwardLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Same as gtk_text_iter_forward_search(), but moves backward.
//
// @match_end will never be set to a #GtkTextIter located after @iter, even if
// there is a possible @match_start before or at @iter.
/*

C function

gtk_text_iter_backward_search
*/
func (recv *TextIter) BackwardSearch(str string, flags TextSearchFlags, limit *TextIter) (bool, *TextIter, *TextIter) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_flags := (C.GtkTextSearchFlags)(flags)

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	c_limit := (*C.GtkTextIter)(C.NULL)
	if limit != nil {
		c_limit = (*C.GtkTextIter)(limit.ToC())
	}

	retC := C.gtk_text_iter_backward_search((*C.GtkTextIter)(recv.native), c_str, c_flags, &c_match_start, &c_match_end, c_limit)
	retGo := retC == C.TRUE

	matchStart := TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd
}

// Moves backward to the previous sentence start; if @iter is already at
// the start of a sentence, moves backward to the next one.  Sentence
// boundaries are determined by Pango and should be correct for nearly
// any language (if not, the correct fix would be to the Pango text
// boundary algorithms).
/*

C function

gtk_text_iter_backward_sentence_start
*/
func (recv *TextIter) BackwardSentenceStart() bool {
	retC := C.gtk_text_iter_backward_sentence_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Calls gtk_text_iter_backward_sentence_start() up to @count times,
// or until it returns %FALSE. If @count is negative, moves forward
// instead of backward.
/*

C function

gtk_text_iter_backward_sentence_starts
*/
func (recv *TextIter) BackwardSentenceStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_sentence_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Moves backward to the next toggle (on or off) of the
// #GtkTextTag @tag, or to the next toggle of any tag if
// @tag is %NULL. If no matching tag toggles are found,
// returns %FALSE, otherwise %TRUE. Does not return toggles
// located at @iter, only toggles before @iter. Sets @iter
// to the location of the toggle, or the start of the buffer
// if no toggle is found.
/*

C function

gtk_text_iter_backward_to_tag_toggle
*/
func (recv *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_backward_to_tag_toggle((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Moves backward to the previous word start. (If @iter is currently on a
// word start, moves backward to the next one after that.) Word breaks
// are determined by Pango and should be correct for nearly any
// language (if not, the correct fix would be to the Pango word break
// algorithms).
/*

C function

gtk_text_iter_backward_word_start
*/
func (recv *TextIter) BackwardWordStart() bool {
	retC := C.gtk_text_iter_backward_word_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Calls gtk_text_iter_backward_word_start() up to @count times.
/*

C function

gtk_text_iter_backward_word_starts
*/
func (recv *TextIter) BackwardWordStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_word_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @tag is toggled on at exactly this point. If @tag
// is %NULL, returns %TRUE if any tag is toggled on at this point.
//
// Note that if gtk_text_iter_begins_tag() returns %TRUE, it means that @iter is
// at the beginning of the tagged range, and that the
// character at @iter is inside the tagged range. In other
// words, unlike gtk_text_iter_ends_tag(), if gtk_text_iter_begins_tag() returns
// %TRUE, gtk_text_iter_has_tag() will also return %TRUE for the same
// parameters.
/*

C function

gtk_text_iter_begins_tag
*/
func (recv *TextIter) BeginsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_begins_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Considering the default editability of the buffer, and tags that
// affect editability, determines whether text inserted at @iter would
// be editable. If text inserted at @iter would be editable then the
// user should be allowed to insert text at @iter.
// gtk_text_buffer_insert_interactive() uses this function to decide
// whether insertions are allowed at a given position.
/*

C function

gtk_text_iter_can_insert
*/
func (recv *TextIter) CanInsert(defaultEditability bool) bool {
	c_default_editability :=
		boolToGboolean(defaultEditability)

	retC := C.gtk_text_iter_can_insert((*C.GtkTextIter)(recv.native), c_default_editability)
	retGo := retC == C.TRUE

	return retGo
}

// A qsort()-style function that returns negative if @lhs is less than
// @rhs, positive if @lhs is greater than @rhs, and 0 if they’re equal.
// Ordering is in character offset order, i.e. the first character in the buffer
// is less than the second character in the buffer.
/*

C function

gtk_text_iter_compare
*/
func (recv *TextIter) Compare(rhs *TextIter) int32 {
	c_rhs := (*C.GtkTextIter)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GtkTextIter)(rhs.ToC())
	}

	retC := C.gtk_text_iter_compare((*C.GtkTextIter)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// Creates a dynamically-allocated copy of an iterator. This function
// is not useful in applications, because iterators can be copied with a
// simple assignment (`GtkTextIter i = j;`). The
// function is used by language bindings.
/*

C function

gtk_text_iter_copy
*/
func (recv *TextIter) Copy() *TextIter {
	retC := C.gtk_text_iter_copy((*C.GtkTextIter)(recv.native))
	retGo := TextIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the character at @iter is within an editable region
// of text.  Non-editable text is “locked” and can’t be changed by the
// user via #GtkTextView. This function is simply a convenience
// wrapper around gtk_text_iter_get_attributes(). If no tags applied
// to this text affect editability, @default_setting will be returned.
//
// You don’t want to use this function to decide whether text can be
// inserted at @iter, because for insertion you don’t want to know
// whether the char at @iter is inside an editable range, you want to
// know whether a new character inserted at @iter would be inside an
// editable range. Use gtk_text_iter_can_insert() to handle this
// case.
/*

C function

gtk_text_iter_editable
*/
func (recv *TextIter) Editable(defaultSetting bool) bool {
	c_default_setting :=
		boolToGboolean(defaultSetting)

	retC := C.gtk_text_iter_editable((*C.GtkTextIter)(recv.native), c_default_setting)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @iter points to the start of the paragraph
// delimiter characters for a line (delimiters will be either a
// newline, a carriage return, a carriage return followed by a
// newline, or a Unicode paragraph separator character). Note that an
// iterator pointing to the \n of a \r\n pair will not be counted as
// the end of a line, the line ends before the \r. The end iterator is
// considered to be at the end of a line, even though there are no
// paragraph delimiter chars there.
/*

C function

gtk_text_iter_ends_line
*/
func (recv *TextIter) EndsLine() bool {
	retC := C.gtk_text_iter_ends_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @iter ends a sentence.  Sentence boundaries are
// determined by Pango and should be correct for nearly any language
// (if not, the correct fix would be to the Pango text boundary
// algorithms).
/*

C function

gtk_text_iter_ends_sentence
*/
func (recv *TextIter) EndsSentence() bool {
	retC := C.gtk_text_iter_ends_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @tag is toggled off at exactly this point. If @tag
// is %NULL, returns %TRUE if any tag is toggled off at this point.
//
// Note that if gtk_text_iter_ends_tag() returns %TRUE, it means that @iter is
// at the end of the tagged range, but that the character
// at @iter is outside the tagged range. In other words,
// unlike gtk_text_iter_starts_tag(), if gtk_text_iter_ends_tag() returns %TRUE,
// gtk_text_iter_has_tag() will return %FALSE for the same parameters.
/*

C function

gtk_text_iter_ends_tag
*/
func (recv *TextIter) EndsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_ends_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @iter ends a natural-language word.  Word breaks
// are determined by Pango and should be correct for nearly any
// language (if not, the correct fix would be to the Pango word break
// algorithms).
/*

C function

gtk_text_iter_ends_word
*/
func (recv *TextIter) EndsWord() bool {
	retC := C.gtk_text_iter_ends_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether two iterators are equal, using the fastest possible
// mechanism. This function is very fast; you can expect it to perform
// better than e.g. getting the character offset for each iterator and
// comparing the offsets yourself. Also, it’s a bit faster than
// gtk_text_iter_compare().
/*

C function

gtk_text_iter_equal
*/
func (recv *TextIter) Equal(rhs *TextIter) bool {
	c_rhs := (*C.GtkTextIter)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GtkTextIter)(rhs.ToC())
	}

	retC := C.gtk_text_iter_equal((*C.GtkTextIter)(recv.native), c_rhs)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter forward by one character offset. Note that images
// embedded in the buffer occupy 1 character slot, so
// gtk_text_iter_forward_char() may actually move onto an image instead
// of a character, if you have images in your buffer.  If @iter is the
// end iterator or one character before it, @iter will now point at
// the end iterator, and gtk_text_iter_forward_char() returns %FALSE for
// convenience when writing loops.
/*

C function

gtk_text_iter_forward_char
*/
func (recv *TextIter) ForwardChar() bool {
	retC := C.gtk_text_iter_forward_char((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count characters if possible (if @count would move past the
// start or end of the buffer, moves to the start or end of the
// buffer). The return value indicates whether the new position of
// @iter is different from its original position, and dereferenceable
// (the last iterator in the buffer is not dereferenceable). If @count
// is 0, the function does nothing and returns %FALSE.
/*

C function

gtk_text_iter_forward_chars
*/
func (recv *TextIter) ForwardChars(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_chars((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter forward by a single cursor position. Cursor positions
// are (unsurprisingly) positions where the cursor can appear. Perhaps
// surprisingly, there may not be a cursor position between all
// characters. The most common example for European languages would be
// a carriage return/newline sequence. For some Unicode characters,
// the equivalent of say the letter “a” with an accent mark will be
// represented as two characters, first the letter then a "combining
// mark" that causes the accent to be rendered; so the cursor can’t go
// between those two characters. See also the #PangoLogAttr-struct and
// pango_break() function.
/*

C function

gtk_text_iter_forward_cursor_position
*/
func (recv *TextIter) ForwardCursorPosition() bool {
	retC := C.gtk_text_iter_forward_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves up to @count cursor positions. See
// gtk_text_iter_forward_cursor_position() for details.
/*

C function

gtk_text_iter_forward_cursor_positions
*/
func (recv *TextIter) ForwardCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_iter_forward_find_char : unsupported parameter pred : no type generator for TextCharPredicate (GtkTextCharPredicate) for param pred

// Moves @iter to the start of the next line. If the iter is already on the
// last line of the buffer, moves the iter to the end of the current line.
// If after the operation, the iter is at the end of the buffer and not
// dereferencable, returns %FALSE. Otherwise, returns %TRUE.
/*

C function

gtk_text_iter_forward_line
*/
func (recv *TextIter) ForwardLine() bool {
	retC := C.gtk_text_iter_forward_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count lines forward, if possible (if @count would move
// past the start or end of the buffer, moves to the start or end of
// the buffer).  The return value indicates whether the iterator moved
// onto a dereferenceable position; if the iterator didn’t move, or
// moved onto the end iterator, then %FALSE is returned. If @count is 0,
// the function does nothing and returns %FALSE. If @count is negative,
// moves backward by 0 - @count lines.
/*

C function

gtk_text_iter_forward_lines
*/
func (recv *TextIter) ForwardLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Searches forward for @str. Any match is returned by setting
// @match_start to the first character of the match and @match_end to the
// first character after the match. The search will not continue past
// @limit. Note that a search is a linear or O(n) operation, so you
// may wish to use @limit to avoid locking up your UI on large
// buffers.
//
// @match_start will never be set to a #GtkTextIter located before @iter, even if
// there is a possible @match_end after or at @iter.
/*

C function

gtk_text_iter_forward_search
*/
func (recv *TextIter) ForwardSearch(str string, flags TextSearchFlags, limit *TextIter) (bool, *TextIter, *TextIter) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_flags := (C.GtkTextSearchFlags)(flags)

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	c_limit := (*C.GtkTextIter)(C.NULL)
	if limit != nil {
		c_limit = (*C.GtkTextIter)(limit.ToC())
	}

	retC := C.gtk_text_iter_forward_search((*C.GtkTextIter)(recv.native), c_str, c_flags, &c_match_start, &c_match_end, c_limit)
	retGo := retC == C.TRUE

	matchStart := TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd
}

// Moves forward to the next sentence end. (If @iter is at the end of
// a sentence, moves to the next end of sentence.)  Sentence
// boundaries are determined by Pango and should be correct for nearly
// any language (if not, the correct fix would be to the Pango text
// boundary algorithms).
/*

C function

gtk_text_iter_forward_sentence_end
*/
func (recv *TextIter) ForwardSentenceEnd() bool {
	retC := C.gtk_text_iter_forward_sentence_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Calls gtk_text_iter_forward_sentence_end() @count times (or until
// gtk_text_iter_forward_sentence_end() returns %FALSE). If @count is
// negative, moves backward instead of forward.
/*

C function

gtk_text_iter_forward_sentence_ends
*/
func (recv *TextIter) ForwardSentenceEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_sentence_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter forward to the “end iterator,” which points one past the last
// valid character in the buffer. gtk_text_iter_get_char() called on the
// end iterator returns 0, which is convenient for writing loops.
/*

C function

gtk_text_iter_forward_to_end
*/
func (recv *TextIter) ForwardToEnd() {
	C.gtk_text_iter_forward_to_end((*C.GtkTextIter)(recv.native))

	return
}

// Moves the iterator to point to the paragraph delimiter characters,
// which will be either a newline, a carriage return, a carriage
// return/newline in sequence, or the Unicode paragraph separator
// character. If the iterator is already at the paragraph delimiter
// characters, moves to the paragraph delimiter characters for the
// next line. If @iter is on the last line in the buffer, which does
// not end in paragraph delimiters, moves to the end iterator (end of
// the last line), and returns %FALSE.
/*

C function

gtk_text_iter_forward_to_line_end
*/
func (recv *TextIter) ForwardToLineEnd() bool {
	retC := C.gtk_text_iter_forward_to_line_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves forward to the next toggle (on or off) of the
// #GtkTextTag @tag, or to the next toggle of any tag if
// @tag is %NULL. If no matching tag toggles are found,
// returns %FALSE, otherwise %TRUE. Does not return toggles
// located at @iter, only toggles after @iter. Sets @iter to
// the location of the toggle, or to the end of the buffer
// if no toggle is found.
/*

C function

gtk_text_iter_forward_to_tag_toggle
*/
func (recv *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_forward_to_tag_toggle((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Moves forward to the next word end. (If @iter is currently on a
// word end, moves forward to the next one after that.) Word breaks
// are determined by Pango and should be correct for nearly any
// language (if not, the correct fix would be to the Pango word break
// algorithms).
/*

C function

gtk_text_iter_forward_word_end
*/
func (recv *TextIter) ForwardWordEnd() bool {
	retC := C.gtk_text_iter_forward_word_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Calls gtk_text_iter_forward_word_end() up to @count times.
/*

C function

gtk_text_iter_forward_word_ends
*/
func (recv *TextIter) ForwardWordEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_word_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Free an iterator allocated on the heap. This function
// is intended for use in language bindings, and is not
// especially useful for applications, because iterators can
// simply be allocated on the stack.
/*

C function

gtk_text_iter_free
*/
func (recv *TextIter) Free() {
	C.gtk_text_iter_free((*C.GtkTextIter)(recv.native))

	return
}

// Computes the effect of any tags applied to this spot in the
// text. The @values parameter should be initialized to the default
// settings you wish to use if no tags are in effect. You’d typically
// obtain the defaults from gtk_text_view_get_default_attributes().
//
// gtk_text_iter_get_attributes() will modify @values, applying the
// effects of any tags present at @iter. If any tags affected @values,
// the function returns %TRUE.
/*

C function

gtk_text_iter_get_attributes
*/
func (recv *TextIter) GetAttributes() (bool, *TextAttributes) {
	var c_values C.GtkTextAttributes

	retC := C.gtk_text_iter_get_attributes((*C.GtkTextIter)(recv.native), &c_values)
	retGo := retC == C.TRUE

	values := TextAttributesNewFromC(unsafe.Pointer(&c_values))

	return retGo, values
}

// Returns the #GtkTextBuffer this iterator is associated with.
/*

C function

gtk_text_iter_get_buffer
*/
func (recv *TextIter) GetBuffer() *TextBuffer {
	retC := C.gtk_text_iter_get_buffer((*C.GtkTextIter)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the number of bytes in the line containing @iter,
// including the paragraph delimiters.
/*

C function

gtk_text_iter_get_bytes_in_line
*/
func (recv *TextIter) GetBytesInLine() int32 {
	retC := C.gtk_text_iter_get_bytes_in_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// The Unicode character at this iterator is returned.  (Equivalent to
// operator* on a C++ iterator.)  If the element at this iterator is a
// non-character element, such as an image embedded in the buffer, the
// Unicode “unknown” character 0xFFFC is returned. If invoked on
// the end iterator, zero is returned; zero is not a valid Unicode character.
// So you can write a loop which ends when gtk_text_iter_get_char()
// returns 0.
/*

C function

gtk_text_iter_get_char
*/
func (recv *TextIter) GetChar() rune {
	retC := C.gtk_text_iter_get_char((*C.GtkTextIter)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// Returns the number of characters in the line containing @iter,
// including the paragraph delimiters.
/*

C function

gtk_text_iter_get_chars_in_line
*/
func (recv *TextIter) GetCharsInLine() int32 {
	retC := C.gtk_text_iter_get_chars_in_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// If the location at @iter contains a child anchor, the
// anchor is returned (with no new reference count added). Otherwise,
// %NULL is returned.
/*

C function

gtk_text_iter_get_child_anchor
*/
func (recv *TextIter) GetChildAnchor() *TextChildAnchor {
	retC := C.gtk_text_iter_get_child_anchor((*C.GtkTextIter)(recv.native))
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// A convenience wrapper around gtk_text_iter_get_attributes(),
// which returns the language in effect at @iter. If no tags affecting
// language apply to @iter, the return value is identical to that of
// gtk_get_default_language().
/*

C function

gtk_text_iter_get_language
*/
func (recv *TextIter) GetLanguage() *pango.Language {
	retC := C.gtk_text_iter_get_language((*C.GtkTextIter)(recv.native))
	retGo := pango.LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the line number containing the iterator. Lines in
// a #GtkTextBuffer are numbered beginning with 0 for the first
// line in the buffer.
/*

C function

gtk_text_iter_get_line
*/
func (recv *TextIter) GetLine() int32 {
	retC := C.gtk_text_iter_get_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the byte index of the iterator, counting
// from the start of a newline-terminated line.
// Remember that #GtkTextBuffer encodes text in
// UTF-8, and that characters can require a variable
// number of bytes to represent.
/*

C function

gtk_text_iter_get_line_index
*/
func (recv *TextIter) GetLineIndex() int32 {
	retC := C.gtk_text_iter_get_line_index((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the character offset of the iterator,
// counting from the start of a newline-terminated line.
// The first character on the line has offset 0.
/*

C function

gtk_text_iter_get_line_offset
*/
func (recv *TextIter) GetLineOffset() int32 {
	retC := C.gtk_text_iter_get_line_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns a list of all #GtkTextMark at this location. Because marks
// are not iterable (they don’t take up any "space" in the buffer,
// they are just marks in between iterable locations), multiple marks
// can exist in the same place. The returned list is not in any
// meaningful order.
/*

C function

gtk_text_iter_get_marks
*/
func (recv *TextIter) GetMarks() *glib.SList {
	retC := C.gtk_text_iter_get_marks((*C.GtkTextIter)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the character offset of an iterator.
// Each character in a #GtkTextBuffer has an offset,
// starting with 0 for the first character in the buffer.
// Use gtk_text_buffer_get_iter_at_offset() to convert an
// offset back into an iterator.
/*

C function

gtk_text_iter_get_offset
*/
func (recv *TextIter) GetOffset() int32 {
	retC := C.gtk_text_iter_get_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// If the element at @iter is a pixbuf, the pixbuf is returned
// (with no new reference count added). Otherwise,
// %NULL is returned.
/*

C function

gtk_text_iter_get_pixbuf
*/
func (recv *TextIter) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_text_iter_get_pixbuf((*C.GtkTextIter)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the text in the given range. A “slice” is an array of
// characters encoded in UTF-8 format, including the Unicode “unknown”
// character 0xFFFC for iterable non-character elements in the buffer,
// such as images.  Because images are encoded in the slice, byte and
// character offsets in the returned array will correspond to byte
// offsets in the text buffer. Note that 0xFFFC can occur in normal
// text as well, so it is not a reliable indicator that a pixbuf or
// widget is in the buffer.
/*

C function

gtk_text_iter_get_slice
*/
func (recv *TextIter) GetSlice(end *TextIter) string {
	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	retC := C.gtk_text_iter_get_slice((*C.GtkTextIter)(recv.native), c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns a list of tags that apply to @iter, in ascending order of
// priority (highest-priority tags are last). The #GtkTextTag in the
// list don’t have a reference added, but you have to free the list
// itself.
/*

C function

gtk_text_iter_get_tags
*/
func (recv *TextIter) GetTags() *glib.SList {
	retC := C.gtk_text_iter_get_tags((*C.GtkTextIter)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns text in the given range.  If the range
// contains non-text elements such as images, the character and byte
// offsets in the returned string will not correspond to character and
// byte offsets in the buffer. If you want offsets to correspond, see
// gtk_text_iter_get_slice().
/*

C function

gtk_text_iter_get_text
*/
func (recv *TextIter) GetText(end *TextIter) string {
	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	retC := C.gtk_text_iter_get_text((*C.GtkTextIter)(recv.native), c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns a list of #GtkTextTag that are toggled on or off at this
// point.  (If @toggled_on is %TRUE, the list contains tags that are
// toggled on.) If a tag is toggled on at @iter, then some non-empty
// range of characters following @iter has that tag applied to it.  If
// a tag is toggled off, then some non-empty range following @iter
// does not have the tag applied to it.
/*

C function

gtk_text_iter_get_toggled_tags
*/
func (recv *TextIter) GetToggledTags(toggledOn bool) *glib.SList {
	c_toggled_on :=
		boolToGboolean(toggledOn)

	retC := C.gtk_text_iter_get_toggled_tags((*C.GtkTextIter)(recv.native), c_toggled_on)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the number of bytes from the start of the
// line to the given @iter, not counting bytes that
// are invisible due to tags with the “invisible” flag
// toggled on.
/*

C function

gtk_text_iter_get_visible_line_index
*/
func (recv *TextIter) GetVisibleLineIndex() int32 {
	retC := C.gtk_text_iter_get_visible_line_index((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the offset in characters from the start of the
// line to the given @iter, not counting characters that
// are invisible due to tags with the “invisible” flag
// toggled on.
/*

C function

gtk_text_iter_get_visible_line_offset
*/
func (recv *TextIter) GetVisibleLineOffset() int32 {
	retC := C.gtk_text_iter_get_visible_line_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Like gtk_text_iter_get_slice(), but invisible text is not included.
// Invisible text is usually invisible because a #GtkTextTag with the
// “invisible” attribute turned on has been applied to it.
/*

C function

gtk_text_iter_get_visible_slice
*/
func (recv *TextIter) GetVisibleSlice(end *TextIter) string {
	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	retC := C.gtk_text_iter_get_visible_slice((*C.GtkTextIter)(recv.native), c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Like gtk_text_iter_get_text(), but invisible text is not included.
// Invisible text is usually invisible because a #GtkTextTag with the
// “invisible” attribute turned on has been applied to it.
/*

C function

gtk_text_iter_get_visible_text
*/
func (recv *TextIter) GetVisibleText(end *TextIter) string {
	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	retC := C.gtk_text_iter_get_visible_text((*C.GtkTextIter)(recv.native), c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns %TRUE if @iter points to a character that is part of a range tagged
// with @tag. See also gtk_text_iter_starts_tag() and gtk_text_iter_ends_tag().
/*

C function

gtk_text_iter_has_tag
*/
func (recv *TextIter) HasTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_has_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether @iter falls in the range [@start, @end).
// @start and @end must be in ascending order.
/*

C function

gtk_text_iter_in_range
*/
func (recv *TextIter) InRange(start *TextIter, end *TextIter) bool {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	retC := C.gtk_text_iter_in_range((*C.GtkTextIter)(recv.native), c_start, c_end)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @iter is inside a sentence (as opposed to in
// between two sentences, e.g. after a period and before the first
// letter of the next sentence).  Sentence boundaries are determined
// by Pango and should be correct for nearly any language (if not, the
// correct fix would be to the Pango text boundary algorithms).
/*

C function

gtk_text_iter_inside_sentence
*/
func (recv *TextIter) InsideSentence() bool {
	retC := C.gtk_text_iter_inside_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether the character pointed by @iter is part of a
// natural-language word (as opposed to say inside some whitespace).  Word
// breaks are determined by Pango and should be correct for nearly any language
// (if not, the correct fix would be to the Pango word break algorithms).
//
// Note that if gtk_text_iter_starts_word() returns %TRUE, then this function
// returns %TRUE too, since @iter points to the first character of the word.
/*

C function

gtk_text_iter_inside_word
*/
func (recv *TextIter) InsideWord() bool {
	retC := C.gtk_text_iter_inside_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// See gtk_text_iter_forward_cursor_position() or #PangoLogAttr or
// pango_break() for details on what a cursor position is.
/*

C function

gtk_text_iter_is_cursor_position
*/
func (recv *TextIter) IsCursorPosition() bool {
	retC := C.gtk_text_iter_is_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @iter is the end iterator, i.e. one past the last
// dereferenceable iterator in the buffer. gtk_text_iter_is_end() is
// the most efficient way to check whether an iterator is the end
// iterator.
/*

C function

gtk_text_iter_is_end
*/
func (recv *TextIter) IsEnd() bool {
	retC := C.gtk_text_iter_is_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @iter is the first iterator in the buffer, that is
// if @iter has a character offset of 0.
/*

C function

gtk_text_iter_is_start
*/
func (recv *TextIter) IsStart() bool {
	retC := C.gtk_text_iter_is_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Swaps the value of @first and @second if @second comes before
// @first in the buffer. That is, ensures that @first and @second are
// in sequence. Most text buffer functions that take a range call this
// automatically on your behalf, so there’s no real reason to call it yourself
// in those cases. There are some exceptions, such as gtk_text_iter_in_range(),
// that expect a pre-sorted range.
/*

C function

gtk_text_iter_order
*/
func (recv *TextIter) Order(second *TextIter) {
	c_second := (*C.GtkTextIter)(C.NULL)
	if second != nil {
		c_second = (*C.GtkTextIter)(second.ToC())
	}

	C.gtk_text_iter_order((*C.GtkTextIter)(recv.native), c_second)

	return
}

// Moves iterator @iter to the start of the line @line_number.  If
// @line_number is negative or larger than the number of lines in the
// buffer, moves @iter to the start of the last line in the buffer.
/*

C function

gtk_text_iter_set_line
*/
func (recv *TextIter) SetLine(lineNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	C.gtk_text_iter_set_line((*C.GtkTextIter)(recv.native), c_line_number)

	return
}

// Same as gtk_text_iter_set_line_offset(), but works with a
// byte index. The given byte index must be at
// the start of a character, it can’t be in the middle of a UTF-8
// encoded character.
/*

C function

gtk_text_iter_set_line_index
*/
func (recv *TextIter) SetLineIndex(byteOnLine int32) {
	c_byte_on_line := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_line_index((*C.GtkTextIter)(recv.native), c_byte_on_line)

	return
}

// Moves @iter within a line, to a new character
// (not byte) offset. The given character offset must be less than or
// equal to the number of characters in the line; if equal, @iter
// moves to the start of the next line. See
// gtk_text_iter_set_line_index() if you have a byte index rather than
// a character offset.
/*

C function

gtk_text_iter_set_line_offset
*/
func (recv *TextIter) SetLineOffset(charOnLine int32) {
	c_char_on_line := (C.gint)(charOnLine)

	C.gtk_text_iter_set_line_offset((*C.GtkTextIter)(recv.native), c_char_on_line)

	return
}

// Sets @iter to point to @char_offset. @char_offset counts from the start
// of the entire text buffer, starting with 0.
/*

C function

gtk_text_iter_set_offset
*/
func (recv *TextIter) SetOffset(charOffset int32) {
	c_char_offset := (C.gint)(charOffset)

	C.gtk_text_iter_set_offset((*C.GtkTextIter)(recv.native), c_char_offset)

	return
}

// Like gtk_text_iter_set_line_index(), but the index is in visible
// bytes, i.e. text with a tag making it invisible is not counted
// in the index.
/*

C function

gtk_text_iter_set_visible_line_index
*/
func (recv *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	c_byte_on_line := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_visible_line_index((*C.GtkTextIter)(recv.native), c_byte_on_line)

	return
}

// Like gtk_text_iter_set_line_offset(), but the offset is in visible
// characters, i.e. text with a tag making it invisible is not
// counted in the offset.
/*

C function

gtk_text_iter_set_visible_line_offset
*/
func (recv *TextIter) SetVisibleLineOffset(charOnLine int32) {
	c_char_on_line := (C.gint)(charOnLine)

	C.gtk_text_iter_set_visible_line_offset((*C.GtkTextIter)(recv.native), c_char_on_line)

	return
}

// Returns %TRUE if @iter begins a paragraph,
// i.e. if gtk_text_iter_get_line_offset() would return 0.
// However this function is potentially more efficient than
// gtk_text_iter_get_line_offset() because it doesn’t have to compute
// the offset, it just has to see whether it’s 0.
/*

C function

gtk_text_iter_starts_line
*/
func (recv *TextIter) StartsLine() bool {
	retC := C.gtk_text_iter_starts_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @iter begins a sentence.  Sentence boundaries are
// determined by Pango and should be correct for nearly any language
// (if not, the correct fix would be to the Pango text boundary
// algorithms).
/*

C function

gtk_text_iter_starts_sentence
*/
func (recv *TextIter) StartsSentence() bool {
	retC := C.gtk_text_iter_starts_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @iter begins a natural-language word.  Word
// breaks are determined by Pango and should be correct for nearly any
// language (if not, the correct fix would be to the Pango word break
// algorithms).
/*

C function

gtk_text_iter_starts_word
*/
func (recv *TextIter) StartsWord() bool {
	retC := C.gtk_text_iter_starts_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This is equivalent to (gtk_text_iter_starts_tag() ||
// gtk_text_iter_ends_tag()), i.e. it tells you whether a range with
// @tag applied to it begins or ends at @iter.
/*

C function

gtk_text_iter_toggles_tag
*/
func (recv *TextIter) TogglesTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_toggles_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

/*

C type

GtkTextMarkClass
*/
type TextMarkClass struct {
	native *C.GtkTextMarkClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextMarkClassNewFromC(u unsafe.Pointer) *TextMarkClass {
	c := (*C.GtkTextMarkClass)(u)
	if c == nil {
		return nil
	}

	g := &TextMarkClass{native: c}

	return g
}

func (recv *TextMarkClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextTagClass
*/
type TextTagClass struct {
	native *C.GtkTextTagClass
	// parent_class : record
	// no type for event
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TextTagClassNewFromC(u unsafe.Pointer) *TextTagClass {
	c := (*C.GtkTextTagClass)(u)
	if c == nil {
		return nil
	}

	g := &TextTagClass{native: c}

	return g
}

func (recv *TextTagClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextTagPrivate
*/
type TextTagPrivate struct {
	native *C.GtkTextTagPrivate
}

func TextTagPrivateNewFromC(u unsafe.Pointer) *TextTagPrivate {
	c := (*C.GtkTextTagPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextTagPrivate{native: c}

	return g
}

func (recv *TextTagPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextTagTableClass
*/
type TextTagTableClass struct {
	native *C.GtkTextTagTableClass
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
	c := (*C.GtkTextTagTableClass)(u)
	if c == nil {
		return nil
	}

	g := &TextTagTableClass{native: c}

	return g
}

func (recv *TextTagTableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextTagTablePrivate
*/
type TextTagTablePrivate struct {
	native *C.GtkTextTagTablePrivate
}

func TextTagTablePrivateNewFromC(u unsafe.Pointer) *TextTagTablePrivate {
	c := (*C.GtkTextTagTablePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextTagTablePrivate{native: c}

	return g
}

func (recv *TextTagTablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextViewAccessibleClass
*/
type TextViewAccessibleClass struct {
	native *C.GtkTextViewAccessibleClass
	// parent_class : record
}

func TextViewAccessibleClassNewFromC(u unsafe.Pointer) *TextViewAccessibleClass {
	c := (*C.GtkTextViewAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &TextViewAccessibleClass{native: c}

	return g
}

func (recv *TextViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextViewAccessiblePrivate
*/
type TextViewAccessiblePrivate struct {
	native *C.GtkTextViewAccessiblePrivate
}

func TextViewAccessiblePrivateNewFromC(u unsafe.Pointer) *TextViewAccessiblePrivate {
	c := (*C.GtkTextViewAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextViewAccessiblePrivate{native: c}

	return g
}

func (recv *TextViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextViewClass
*/
type TextViewClass struct {
	native *C.GtkTextViewClass
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
	c := (*C.GtkTextViewClass)(u)
	if c == nil {
		return nil
	}

	g := &TextViewClass{native: c}

	return g
}

func (recv *TextViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTextViewPrivate
*/
type TextViewPrivate struct {
	native *C.GtkTextViewPrivate
}

func TextViewPrivateNewFromC(u unsafe.Pointer) *TextViewPrivate {
	c := (*C.GtkTextViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TextViewPrivate{native: c}

	return g
}

func (recv *TextViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkThemeEngine
*/
type ThemeEngine struct {
	native *C.GtkThemeEngine
}

func ThemeEngineNewFromC(u unsafe.Pointer) *ThemeEngine {
	c := (*C.GtkThemeEngine)(u)
	if c == nil {
		return nil
	}

	g := &ThemeEngine{native: c}

	return g
}

func (recv *ThemeEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Base class for theming engines.
/*

C type

GtkThemingEngineClass
*/
type ThemingEngineClass struct {
	native *C.GtkThemingEngineClass
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
	c := (*C.GtkThemingEngineClass)(u)
	if c == nil {
		return nil
	}

	g := &ThemingEngineClass{native: c}

	return g
}

func (recv *ThemingEngineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkThemingEnginePrivate
*/
type ThemingEnginePrivate struct {
	native *C.GtkThemingEnginePrivate
}

func ThemingEnginePrivateNewFromC(u unsafe.Pointer) *ThemingEnginePrivate {
	c := (*C.GtkThemingEnginePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ThemingEnginePrivate{native: c}

	return g
}

func (recv *ThemingEnginePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleActionClass
*/
type ToggleActionClass struct {
	native *C.GtkToggleActionClass
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleActionClassNewFromC(u unsafe.Pointer) *ToggleActionClass {
	c := (*C.GtkToggleActionClass)(u)
	if c == nil {
		return nil
	}

	g := &ToggleActionClass{native: c}

	return g
}

func (recv *ToggleActionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkToggleActionEntry structs are used with
// gtk_action_group_add_toggle_actions() to construct toggle actions.
/*

C type

GtkToggleActionEntry
*/
type ToggleActionEntry struct {
	native      *C.GtkToggleActionEntry
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
	IsActive bool
}

func ToggleActionEntryNewFromC(u unsafe.Pointer) *ToggleActionEntry {
	c := (*C.GtkToggleActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &ToggleActionEntry{
		Accelerator: C.GoString(c.accelerator),
		IsActive:    c.is_active == C.TRUE,
		Label:       C.GoString(c.label),
		Name:        C.GoString(c.name),
		StockId:     C.GoString(c.stock_id),
		Tooltip:     C.GoString(c.tooltip),
		native:      c,
	}

	return g
}

func (recv *ToggleActionEntry) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.stock_id =
		C.CString(recv.StockId)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.accelerator =
		C.CString(recv.Accelerator)
	recv.native.tooltip =
		C.CString(recv.Tooltip)
	recv.native.is_active =
		boolToGboolean(recv.IsActive)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleActionPrivate
*/
type ToggleActionPrivate struct {
	native *C.GtkToggleActionPrivate
}

func ToggleActionPrivateNewFromC(u unsafe.Pointer) *ToggleActionPrivate {
	c := (*C.GtkToggleActionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToggleActionPrivate{native: c}

	return g
}

func (recv *ToggleActionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleButtonAccessibleClass
*/
type ToggleButtonAccessibleClass struct {
	native *C.GtkToggleButtonAccessibleClass
	// parent_class : record
}

func ToggleButtonAccessibleClassNewFromC(u unsafe.Pointer) *ToggleButtonAccessibleClass {
	c := (*C.GtkToggleButtonAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonAccessibleClass{native: c}

	return g
}

func (recv *ToggleButtonAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleButtonAccessiblePrivate
*/
type ToggleButtonAccessiblePrivate struct {
	native *C.GtkToggleButtonAccessiblePrivate
}

func ToggleButtonAccessiblePrivateNewFromC(u unsafe.Pointer) *ToggleButtonAccessiblePrivate {
	c := (*C.GtkToggleButtonAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonAccessiblePrivate{native: c}

	return g
}

func (recv *ToggleButtonAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleButtonClass
*/
type ToggleButtonClass struct {
	native *C.GtkToggleButtonClass
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleButtonClassNewFromC(u unsafe.Pointer) *ToggleButtonClass {
	c := (*C.GtkToggleButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonClass{native: c}

	return g
}

func (recv *ToggleButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleButtonPrivate
*/
type ToggleButtonPrivate struct {
	native *C.GtkToggleButtonPrivate
}

func ToggleButtonPrivateNewFromC(u unsafe.Pointer) *ToggleButtonPrivate {
	c := (*C.GtkToggleButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonPrivate{native: c}

	return g
}

func (recv *ToggleButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleToolButtonClass
*/
type ToggleToolButtonClass struct {
	native *C.GtkToggleToolButtonClass
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToggleToolButtonClassNewFromC(u unsafe.Pointer) *ToggleToolButtonClass {
	c := (*C.GtkToggleToolButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ToggleToolButtonClass{native: c}

	return g
}

func (recv *ToggleToolButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToggleToolButtonPrivate
*/
type ToggleToolButtonPrivate struct {
	native *C.GtkToggleToolButtonPrivate
}

func ToggleToolButtonPrivateNewFromC(u unsafe.Pointer) *ToggleToolButtonPrivate {
	c := (*C.GtkToggleToolButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToggleToolButtonPrivate{native: c}

	return g
}

func (recv *ToggleToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolButtonClass
*/
type ToolButtonClass struct {
	native *C.GtkToolButtonClass
	// parent_class : record
	ButtonType gobject.Type
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolButtonClassNewFromC(u unsafe.Pointer) *ToolButtonClass {
	c := (*C.GtkToolButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &ToolButtonClass{
		ButtonType: (gobject.Type)(c.button_type),
		native:     c,
	}

	return g
}

func (recv *ToolButtonClass) ToC() unsafe.Pointer {
	recv.native.button_type =
		(C.GType)(recv.ButtonType)

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolButtonPrivate
*/
type ToolButtonPrivate struct {
	native *C.GtkToolButtonPrivate
}

func ToolButtonPrivateNewFromC(u unsafe.Pointer) *ToolButtonPrivate {
	c := (*C.GtkToolButtonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToolButtonPrivate{native: c}

	return g
}

func (recv *ToolButtonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolItemClass
*/
type ToolItemClass struct {
	native *C.GtkToolItemClass
	// parent_class : record
	// no type for create_menu_proxy
	// no type for toolbar_reconfigured
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolItemClassNewFromC(u unsafe.Pointer) *ToolItemClass {
	c := (*C.GtkToolItemClass)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemClass{native: c}

	return g
}

func (recv *ToolItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolItemGroupClass
*/
type ToolItemGroupClass struct {
	native *C.GtkToolItemGroupClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolItemGroupClassNewFromC(u unsafe.Pointer) *ToolItemGroupClass {
	c := (*C.GtkToolItemGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemGroupClass{native: c}

	return g
}

func (recv *ToolItemGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolItemGroupPrivate
*/
type ToolItemGroupPrivate struct {
	native *C.GtkToolItemGroupPrivate
}

func ToolItemGroupPrivateNewFromC(u unsafe.Pointer) *ToolItemGroupPrivate {
	c := (*C.GtkToolItemGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemGroupPrivate{native: c}

	return g
}

func (recv *ToolItemGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolItemPrivate
*/
type ToolItemPrivate struct {
	native *C.GtkToolItemPrivate
}

func ToolItemPrivateNewFromC(u unsafe.Pointer) *ToolItemPrivate {
	c := (*C.GtkToolItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemPrivate{native: c}

	return g
}

func (recv *ToolItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolPaletteClass
*/
type ToolPaletteClass struct {
	native *C.GtkToolPaletteClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ToolPaletteClassNewFromC(u unsafe.Pointer) *ToolPaletteClass {
	c := (*C.GtkToolPaletteClass)(u)
	if c == nil {
		return nil
	}

	g := &ToolPaletteClass{native: c}

	return g
}

func (recv *ToolPaletteClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolPalettePrivate
*/
type ToolPalettePrivate struct {
	native *C.GtkToolPalettePrivate
}

func ToolPalettePrivateNewFromC(u unsafe.Pointer) *ToolPalettePrivate {
	c := (*C.GtkToolPalettePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToolPalettePrivate{native: c}

	return g
}

func (recv *ToolPalettePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Virtual function table for the #GtkToolShell interface.
/*

C type

GtkToolShellIface
*/
type ToolShellIface struct {
	native *C.GtkToolShellIface
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
	c := (*C.GtkToolShellIface)(u)
	if c == nil {
		return nil
	}

	g := &ToolShellIface{native: c}

	return g
}

func (recv *ToolShellIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolbarClass
*/
type ToolbarClass struct {
	native *C.GtkToolbarClass
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
	c := (*C.GtkToolbarClass)(u)
	if c == nil {
		return nil
	}

	g := &ToolbarClass{native: c}

	return g
}

func (recv *ToolbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToolbarPrivate
*/
type ToolbarPrivate struct {
	native *C.GtkToolbarPrivate
}

func ToolbarPrivateNewFromC(u unsafe.Pointer) *ToolbarPrivate {
	c := (*C.GtkToolbarPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToolbarPrivate{native: c}

	return g
}

func (recv *ToolbarPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToplevelAccessibleClass
*/
type ToplevelAccessibleClass struct {
	native *C.GtkToplevelAccessibleClass
	// parent_class : record
}

func ToplevelAccessibleClassNewFromC(u unsafe.Pointer) *ToplevelAccessibleClass {
	c := (*C.GtkToplevelAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &ToplevelAccessibleClass{native: c}

	return g
}

func (recv *ToplevelAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkToplevelAccessiblePrivate
*/
type ToplevelAccessiblePrivate struct {
	native *C.GtkToplevelAccessiblePrivate
}

func ToplevelAccessiblePrivateNewFromC(u unsafe.Pointer) *ToplevelAccessiblePrivate {
	c := (*C.GtkToplevelAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ToplevelAccessiblePrivate{native: c}

	return g
}

func (recv *ToplevelAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeDragDestIface
*/
type TreeDragDestIface struct {
	native *C.GtkTreeDragDestIface
	// Private : g_iface
	// no type for drag_data_received
	// no type for row_drop_possible
}

func TreeDragDestIfaceNewFromC(u unsafe.Pointer) *TreeDragDestIface {
	c := (*C.GtkTreeDragDestIface)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragDestIface{native: c}

	return g
}

func (recv *TreeDragDestIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeDragSourceIface
*/
type TreeDragSourceIface struct {
	native *C.GtkTreeDragSourceIface
	// Private : g_iface
	// no type for row_draggable
	// no type for drag_data_get
	// no type for drag_data_delete
}

func TreeDragSourceIfaceNewFromC(u unsafe.Pointer) *TreeDragSourceIface {
	c := (*C.GtkTreeDragSourceIface)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragSourceIface{native: c}

	return g
}

func (recv *TreeDragSourceIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GtkTreeIter is the primary structure
// for accessing a #GtkTreeModel. Models are expected to put a unique
// integer in the @stamp member, and put
// model-specific data in the three @user_data
// members.
/*

C type

GtkTreeIter
*/
type TreeIter struct {
	native    *C.GtkTreeIter
	Stamp     int32
	UserData  uintptr
	UserData2 uintptr
	UserData3 uintptr
}

func TreeIterNewFromC(u unsafe.Pointer) *TreeIter {
	c := (*C.GtkTreeIter)(u)
	if c == nil {
		return nil
	}

	g := &TreeIter{
		Stamp:     (int32)(c.stamp),
		UserData:  (uintptr)(c.user_data),
		UserData2: (uintptr)(c.user_data2),
		UserData3: (uintptr)(c.user_data3),
		native:    c,
	}

	return g
}

func (recv *TreeIter) ToC() unsafe.Pointer {
	recv.native.stamp =
		(C.gint)(recv.Stamp)
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
	recv.native.user_data2 =
		(C.gpointer)(recv.UserData2)
	recv.native.user_data3 =
		(C.gpointer)(recv.UserData3)

	return (unsafe.Pointer)(recv.native)
}

// Creates a dynamically allocated tree iterator as a copy of @iter.
//
// This function is not intended for use in applications,
// because you can just copy the structs by value
// (`GtkTreeIter new_iter = iter;`).
// You must free this iter with gtk_tree_iter_free().
/*

C function

gtk_tree_iter_copy
*/
func (recv *TreeIter) Copy() *TreeIter {
	retC := C.gtk_tree_iter_copy((*C.GtkTreeIter)(recv.native))
	retGo := TreeIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees an iterator that has been allocated by gtk_tree_iter_copy().
//
// This function is mainly used for language bindings.
/*

C function

gtk_tree_iter_free
*/
func (recv *TreeIter) Free() {
	C.gtk_tree_iter_free((*C.GtkTreeIter)(recv.native))

	return
}

/*

C type

GtkTreeModelFilterClass
*/
type TreeModelFilterClass struct {
	native *C.GtkTreeModelFilterClass
	// parent_class : record
	// no type for visible
	// no type for modify
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeModelFilterClassNewFromC(u unsafe.Pointer) *TreeModelFilterClass {
	c := (*C.GtkTreeModelFilterClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelFilterClass{native: c}

	return g
}

func (recv *TreeModelFilterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeModelFilterPrivate
*/
type TreeModelFilterPrivate struct {
	native *C.GtkTreeModelFilterPrivate
}

func TreeModelFilterPrivateNewFromC(u unsafe.Pointer) *TreeModelFilterPrivate {
	c := (*C.GtkTreeModelFilterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelFilterPrivate{native: c}

	return g
}

func (recv *TreeModelFilterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeModelIface
*/
type TreeModelIface struct {
	native *C.GtkTreeModelIface
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
	c := (*C.GtkTreeModelIface)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelIface{native: c}

	return g
}

func (recv *TreeModelIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeModelSortClass
*/
type TreeModelSortClass struct {
	native *C.GtkTreeModelSortClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeModelSortClassNewFromC(u unsafe.Pointer) *TreeModelSortClass {
	c := (*C.GtkTreeModelSortClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelSortClass{native: c}

	return g
}

func (recv *TreeModelSortClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeModelSortPrivate
*/
type TreeModelSortPrivate struct {
	native *C.GtkTreeModelSortPrivate
}

func TreeModelSortPrivateNewFromC(u unsafe.Pointer) *TreeModelSortPrivate {
	c := (*C.GtkTreeModelSortPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelSortPrivate{native: c}

	return g
}

func (recv *TreeModelSortPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreePath
*/
type TreePath struct {
	native *C.GtkTreePath
}

func TreePathNewFromC(u unsafe.Pointer) *TreePath {
	c := (*C.GtkTreePath)(u)
	if c == nil {
		return nil
	}

	g := &TreePath{native: c}

	return g
}

func (recv *TreePath) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GtkTreePath-struct.
// This refers to a row.
/*

C function

gtk_tree_path_new
*/
func TreePathNew() *TreePath {
	retC := C.gtk_tree_path_new()
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkTreePath-struct.
//
// The string representation of this path is “0”.
/*

C function

gtk_tree_path_new_first
*/
func TreePathNewFirst() *TreePath {
	retC := C.gtk_tree_path_new_first()
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkTreePath-struct initialized to @path.
//
// @path is expected to be a colon separated list of numbers.
// For example, the string “10:4:0” would create a path of depth
// 3 pointing to the 11th child of the root node, the 5th
// child of that 11th child, and the 1st child of that 5th child.
// If an invalid path string is passed in, %NULL is returned.
/*

C function

gtk_tree_path_new_from_string
*/
func TreePathNewFromString(path string) *TreePath {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_tree_path_new_from_string(c_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends a new index to a path.
//
// As a result, the depth of the path is increased.
/*

C function

gtk_tree_path_append_index
*/
func (recv *TreePath) AppendIndex(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_tree_path_append_index((*C.GtkTreePath)(recv.native), c_index_)

	return
}

// Compares two paths.
//
// If @a appears before @b in a tree, then -1 is returned.
// If @b appears before @a, then 1 is returned.
// If the two nodes are equal, then 0 is returned.
/*

C function

gtk_tree_path_compare
*/
func (recv *TreePath) Compare(b *TreePath) int32 {
	c_b := (*C.GtkTreePath)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreePath)(b.ToC())
	}

	retC := C.gtk_tree_path_compare((*C.GtkTreePath)(recv.native), c_b)
	retGo := (int32)(retC)

	return retGo
}

// Creates a new #GtkTreePath-struct as a copy of @path.
/*

C function

gtk_tree_path_copy
*/
func (recv *TreePath) Copy() *TreePath {
	retC := C.gtk_tree_path_copy((*C.GtkTreePath)(recv.native))
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Moves @path to point to the first child of the current path.
/*

C function

gtk_tree_path_down
*/
func (recv *TreePath) Down() {
	C.gtk_tree_path_down((*C.GtkTreePath)(recv.native))

	return
}

// Frees @path. If @path is %NULL, it simply returns.
/*

C function

gtk_tree_path_free
*/
func (recv *TreePath) Free() {
	C.gtk_tree_path_free((*C.GtkTreePath)(recv.native))

	return
}

// Returns the current depth of @path.
/*

C function

gtk_tree_path_get_depth
*/
func (recv *TreePath) GetDepth() int32 {
	retC := C.gtk_tree_path_get_depth((*C.GtkTreePath)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : gtk_tree_path_get_indices

// Returns %TRUE if @descendant is a descendant of @path.
/*

C function

gtk_tree_path_is_ancestor
*/
func (recv *TreePath) IsAncestor(descendant *TreePath) bool {
	c_descendant := (*C.GtkTreePath)(C.NULL)
	if descendant != nil {
		c_descendant = (*C.GtkTreePath)(descendant.ToC())
	}

	retC := C.gtk_tree_path_is_ancestor((*C.GtkTreePath)(recv.native), c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if @path is a descendant of @ancestor.
/*

C function

gtk_tree_path_is_descendant
*/
func (recv *TreePath) IsDescendant(ancestor *TreePath) bool {
	c_ancestor := (*C.GtkTreePath)(C.NULL)
	if ancestor != nil {
		c_ancestor = (*C.GtkTreePath)(ancestor.ToC())
	}

	retC := C.gtk_tree_path_is_descendant((*C.GtkTreePath)(recv.native), c_ancestor)
	retGo := retC == C.TRUE

	return retGo
}

// Moves the @path to point to the next node at the current depth.
/*

C function

gtk_tree_path_next
*/
func (recv *TreePath) Next() {
	C.gtk_tree_path_next((*C.GtkTreePath)(recv.native))

	return
}

// Prepends a new index to a path.
//
// As a result, the depth of the path is increased.
/*

C function

gtk_tree_path_prepend_index
*/
func (recv *TreePath) PrependIndex(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_tree_path_prepend_index((*C.GtkTreePath)(recv.native), c_index_)

	return
}

// Moves the @path to point to the previous node at the
// current depth, if it exists.
/*

C function

gtk_tree_path_prev
*/
func (recv *TreePath) Prev() bool {
	retC := C.gtk_tree_path_prev((*C.GtkTreePath)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Generates a string representation of the path.
//
// This string is a “:” separated list of numbers.
// For example, “4:10:0:3” would be an acceptable
// return value for this string.
/*

C function

gtk_tree_path_to_string
*/
func (recv *TreePath) ToString() string {
	retC := C.gtk_tree_path_to_string((*C.GtkTreePath)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Moves the @path to point to its parent node, if it has a parent.
/*

C function

gtk_tree_path_up
*/
func (recv *TreePath) Up() bool {
	retC := C.gtk_tree_path_up((*C.GtkTreePath)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// A GtkTreeRowReference tracks model changes so that it always refers to the
// same row (a #GtkTreePath refers to a position, not a fixed row). Create a
// new GtkTreeRowReference with gtk_tree_row_reference_new().
/*

C type

GtkTreeRowReference
*/
type TreeRowReference struct {
	native *C.GtkTreeRowReference
}

func TreeRowReferenceNewFromC(u unsafe.Pointer) *TreeRowReference {
	c := (*C.GtkTreeRowReference)(u)
	if c == nil {
		return nil
	}

	g := &TreeRowReference{native: c}

	return g
}

func (recv *TreeRowReference) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a row reference based on @path.
//
// This reference will keep pointing to the node pointed to
// by @path, so long as it exists. Any changes that occur on @model are
// propagated, and the path is updated appropriately. If
// @path isn’t a valid path in @model, then %NULL is returned.
/*

C function

gtk_tree_row_reference_new
*/
func TreeRowReferenceNew(model *TreeModel, path *TreePath) *TreeRowReference {
	c_model := (*C.GtkTreeModel)(model.ToC())

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_row_reference_new(c_model, c_path)
	retGo := TreeRowReferenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// You do not need to use this function.
//
// Creates a row reference based on @path.
//
// This reference will keep pointing to the node pointed to
// by @path, so long as it exists. If @path isn’t a valid
// path in @model, then %NULL is returned. However, unlike
// references created with gtk_tree_row_reference_new(), it
// does not listen to the model for changes. The creator of
// the row reference must do this explicitly using
// gtk_tree_row_reference_inserted(), gtk_tree_row_reference_deleted(),
// gtk_tree_row_reference_reordered().
//
// These functions must be called exactly once per proxy when the
// corresponding signal on the model is emitted. This single call
// updates all row references for that proxy. Since built-in GTK+
// objects like #GtkTreeView already use this mechanism internally,
// using them as the proxy object will produce unpredictable results.
// Further more, passing the same object as @model and @proxy
// doesn’t work for reasons of internal implementation.
//
// This type of row reference is primarily meant by structures that
// need to carefully monitor exactly when a row reference updates
// itself, and is not generally needed by most applications.
/*

C function

gtk_tree_row_reference_new_proxy
*/
func TreeRowReferenceNewProxy(proxy *gobject.Object, model *TreeModel, path *TreePath) *TreeRowReference {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_model := (*C.GtkTreeModel)(model.ToC())

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_row_reference_new_proxy(c_proxy, c_model, c_path)
	retGo := TreeRowReferenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free’s @reference. @reference may be %NULL
/*

C function

gtk_tree_row_reference_free
*/
func (recv *TreeRowReference) Free() {
	C.gtk_tree_row_reference_free((*C.GtkTreeRowReference)(recv.native))

	return
}

// Returns a path that the row reference currently points to,
// or %NULL if the path pointed to is no longer valid.
/*

C function

gtk_tree_row_reference_get_path
*/
func (recv *TreeRowReference) GetPath() *TreePath {
	retC := C.gtk_tree_row_reference_get_path((*C.GtkTreeRowReference)(recv.native))
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns %TRUE if the @reference is non-%NULL and refers to
// a current valid path.
/*

C function

gtk_tree_row_reference_valid
*/
func (recv *TreeRowReference) Valid() bool {
	retC := C.gtk_tree_row_reference_valid((*C.GtkTreeRowReference)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

/*

C type

GtkTreeSelectionClass
*/
type TreeSelectionClass struct {
	native *C.GtkTreeSelectionClass
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeSelectionClassNewFromC(u unsafe.Pointer) *TreeSelectionClass {
	c := (*C.GtkTreeSelectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeSelectionClass{native: c}

	return g
}

func (recv *TreeSelectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeSelectionPrivate
*/
type TreeSelectionPrivate struct {
	native *C.GtkTreeSelectionPrivate
}

func TreeSelectionPrivateNewFromC(u unsafe.Pointer) *TreeSelectionPrivate {
	c := (*C.GtkTreeSelectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeSelectionPrivate{native: c}

	return g
}

func (recv *TreeSelectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeSortableIface
*/
type TreeSortableIface struct {
	native *C.GtkTreeSortableIface
	// Private : g_iface
	// no type for sort_column_changed
	// no type for get_sort_column_id
	// no type for set_sort_column_id
	// no type for set_sort_func
	// no type for set_default_sort_func
	// no type for has_default_sort_func
}

func TreeSortableIfaceNewFromC(u unsafe.Pointer) *TreeSortableIface {
	c := (*C.GtkTreeSortableIface)(u)
	if c == nil {
		return nil
	}

	g := &TreeSortableIface{native: c}

	return g
}

func (recv *TreeSortableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeStoreClass
*/
type TreeStoreClass struct {
	native *C.GtkTreeStoreClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeStoreClassNewFromC(u unsafe.Pointer) *TreeStoreClass {
	c := (*C.GtkTreeStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeStoreClass{native: c}

	return g
}

func (recv *TreeStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeStorePrivate
*/
type TreeStorePrivate struct {
	native *C.GtkTreeStorePrivate
}

func TreeStorePrivateNewFromC(u unsafe.Pointer) *TreeStorePrivate {
	c := (*C.GtkTreeStorePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeStorePrivate{native: c}

	return g
}

func (recv *TreeStorePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewAccessibleClass
*/
type TreeViewAccessibleClass struct {
	native *C.GtkTreeViewAccessibleClass
	// parent_class : record
}

func TreeViewAccessibleClassNewFromC(u unsafe.Pointer) *TreeViewAccessibleClass {
	c := (*C.GtkTreeViewAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewAccessibleClass{native: c}

	return g
}

func (recv *TreeViewAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewAccessiblePrivate
*/
type TreeViewAccessiblePrivate struct {
	native *C.GtkTreeViewAccessiblePrivate
}

func TreeViewAccessiblePrivateNewFromC(u unsafe.Pointer) *TreeViewAccessiblePrivate {
	c := (*C.GtkTreeViewAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewAccessiblePrivate{native: c}

	return g
}

func (recv *TreeViewAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewClass
*/
type TreeViewClass struct {
	native *C.GtkTreeViewClass
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
	c := (*C.GtkTreeViewClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewClass{native: c}

	return g
}

func (recv *TreeViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewColumnClass
*/
type TreeViewColumnClass struct {
	native *C.GtkTreeViewColumnClass
	// parent_class : record
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func TreeViewColumnClassNewFromC(u unsafe.Pointer) *TreeViewColumnClass {
	c := (*C.GtkTreeViewColumnClass)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewColumnClass{native: c}

	return g
}

func (recv *TreeViewColumnClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewColumnPrivate
*/
type TreeViewColumnPrivate struct {
	native *C.GtkTreeViewColumnPrivate
}

func TreeViewColumnPrivateNewFromC(u unsafe.Pointer) *TreeViewColumnPrivate {
	c := (*C.GtkTreeViewColumnPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewColumnPrivate{native: c}

	return g
}

func (recv *TreeViewColumnPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkTreeViewPrivate
*/
type TreeViewPrivate struct {
	native *C.GtkTreeViewPrivate
}

func TreeViewPrivateNewFromC(u unsafe.Pointer) *TreeViewPrivate {
	c := (*C.GtkTreeViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewPrivate{native: c}

	return g
}

func (recv *TreeViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkUIManagerClass
*/
type UIManagerClass struct {
	native *C.GtkUIManagerClass
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
	c := (*C.GtkUIManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &UIManagerClass{native: c}

	return g
}

func (recv *UIManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkUIManagerPrivate
*/
type UIManagerPrivate struct {
	native *C.GtkUIManagerPrivate
}

func UIManagerPrivateNewFromC(u unsafe.Pointer) *UIManagerPrivate {
	c := (*C.GtkUIManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UIManagerPrivate{native: c}

	return g
}

func (recv *UIManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVBoxClass
*/
type VBoxClass struct {
	native *C.GtkVBoxClass
	// parent_class : record
}

func VBoxClassNewFromC(u unsafe.Pointer) *VBoxClass {
	c := (*C.GtkVBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &VBoxClass{native: c}

	return g
}

func (recv *VBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVButtonBoxClass
*/
type VButtonBoxClass struct {
	native *C.GtkVButtonBoxClass
	// parent_class : record
}

func VButtonBoxClassNewFromC(u unsafe.Pointer) *VButtonBoxClass {
	c := (*C.GtkVButtonBoxClass)(u)
	if c == nil {
		return nil
	}

	g := &VButtonBoxClass{native: c}

	return g
}

func (recv *VButtonBoxClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVPanedClass
*/
type VPanedClass struct {
	native *C.GtkVPanedClass
	// parent_class : record
}

func VPanedClassNewFromC(u unsafe.Pointer) *VPanedClass {
	c := (*C.GtkVPanedClass)(u)
	if c == nil {
		return nil
	}

	g := &VPanedClass{native: c}

	return g
}

func (recv *VPanedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVScaleClass
*/
type VScaleClass struct {
	native *C.GtkVScaleClass
	// parent_class : record
}

func VScaleClassNewFromC(u unsafe.Pointer) *VScaleClass {
	c := (*C.GtkVScaleClass)(u)
	if c == nil {
		return nil
	}

	g := &VScaleClass{native: c}

	return g
}

func (recv *VScaleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVScrollbarClass
*/
type VScrollbarClass struct {
	native *C.GtkVScrollbarClass
	// parent_class : record
}

func VScrollbarClassNewFromC(u unsafe.Pointer) *VScrollbarClass {
	c := (*C.GtkVScrollbarClass)(u)
	if c == nil {
		return nil
	}

	g := &VScrollbarClass{native: c}

	return g
}

func (recv *VScrollbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVSeparatorClass
*/
type VSeparatorClass struct {
	native *C.GtkVSeparatorClass
	// parent_class : record
}

func VSeparatorClassNewFromC(u unsafe.Pointer) *VSeparatorClass {
	c := (*C.GtkVSeparatorClass)(u)
	if c == nil {
		return nil
	}

	g := &VSeparatorClass{native: c}

	return g
}

func (recv *VSeparatorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkViewportClass
*/
type ViewportClass struct {
	native *C.GtkViewportClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func ViewportClassNewFromC(u unsafe.Pointer) *ViewportClass {
	c := (*C.GtkViewportClass)(u)
	if c == nil {
		return nil
	}

	g := &ViewportClass{native: c}

	return g
}

func (recv *ViewportClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkViewportPrivate
*/
type ViewportPrivate struct {
	native *C.GtkViewportPrivate
}

func ViewportPrivateNewFromC(u unsafe.Pointer) *ViewportPrivate {
	c := (*C.GtkViewportPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ViewportPrivate{native: c}

	return g
}

func (recv *ViewportPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkVolumeButtonClass
*/
type VolumeButtonClass struct {
	native *C.GtkVolumeButtonClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func VolumeButtonClassNewFromC(u unsafe.Pointer) *VolumeButtonClass {
	c := (*C.GtkVolumeButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &VolumeButtonClass{native: c}

	return g
}

func (recv *VolumeButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWidgetAccessibleClass
*/
type WidgetAccessibleClass struct {
	native *C.GtkWidgetAccessibleClass
	// parent_class : record
	// no type for notify_gtk
}

func WidgetAccessibleClassNewFromC(u unsafe.Pointer) *WidgetAccessibleClass {
	c := (*C.GtkWidgetAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &WidgetAccessibleClass{native: c}

	return g
}

func (recv *WidgetAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWidgetAccessiblePrivate
*/
type WidgetAccessiblePrivate struct {
	native *C.GtkWidgetAccessiblePrivate
}

func WidgetAccessiblePrivateNewFromC(u unsafe.Pointer) *WidgetAccessiblePrivate {
	c := (*C.GtkWidgetAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &WidgetAccessiblePrivate{native: c}

	return g
}

func (recv *WidgetAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWidgetClass
*/
type WidgetClass struct {
	native *C.GtkWidgetClass
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
	c := (*C.GtkWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &WidgetClass{
		ActivateSignal: (uint32)(c.activate_signal),
		native:         c,
	}

	return g
}

func (recv *WidgetClass) ToC() unsafe.Pointer {
	recv.native.activate_signal =
		(C.guint)(recv.ActivateSignal)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_widget_class_install_style_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_widget_class_install_style_property_parser : unsupported parameter pspec : Blacklisted record : GParamSpec

/*

C type

GtkWidgetClassPrivate
*/
type WidgetClassPrivate struct {
	native *C.GtkWidgetClassPrivate
}

func WidgetClassPrivateNewFromC(u unsafe.Pointer) *WidgetClassPrivate {
	c := (*C.GtkWidgetClassPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WidgetClassPrivate{native: c}

	return g
}

func (recv *WidgetClassPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkWidgetPath is a boxed type that represents a widget hierarchy from
// the topmost widget, typically a toplevel, to any child. This widget
// path abstraction is used in #GtkStyleContext on behalf of the real
// widget in order to query style information.
//
// If you are using GTK+ widgets, you probably will not need to use
// this API directly, as there is gtk_widget_get_path(), and the style
// context returned by gtk_widget_get_style_context() will be automatically
// updated on widget hierarchy changes.
//
// The widget path generation is generally simple:
//
// Defining a button within a window
//
// |[<!-- language="C" -->
// {
// GtkWidgetPath *path;
//
// path = gtk_widget_path_new ();
// gtk_widget_path_append_type (path, GTK_TYPE_WINDOW);
// gtk_widget_path_append_type (path, GTK_TYPE_BUTTON);
// }
// ]|
//
// Although more complex information, such as widget names, or
// different classes (property that may be used by other widget
// types) and intermediate regions may be included:
//
// Defining the first tab widget in a notebook
//
// |[<!-- language="C" -->
// {
// GtkWidgetPath *path;
// guint pos;
//
// path = gtk_widget_path_new ();
//
// pos = gtk_widget_path_append_type (path, GTK_TYPE_NOTEBOOK);
// gtk_widget_path_iter_add_region (path, pos, "tab", GTK_REGION_EVEN | GTK_REGION_FIRST);
//
// pos = gtk_widget_path_append_type (path, GTK_TYPE_LABEL);
// gtk_widget_path_iter_set_name (path, pos, "first tab label");
// }
// ]|
//
// All this information will be used to match the style information
// that applies to the described widget.
/*

C type

GtkWidgetPath
*/
type WidgetPath struct {
	native *C.GtkWidgetPath
}

func WidgetPathNewFromC(u unsafe.Pointer) *WidgetPath {
	c := (*C.GtkWidgetPath)(u)
	if c == nil {
		return nil
	}

	g := &WidgetPath{native: c}

	return g
}

func (recv *WidgetPath) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the name corresponding to the widget found at
// the position @pos in the widget hierarchy defined by
// @path
/*

C function

gtk_widget_path_iter_get_name
*/
func (recv *WidgetPath) IterGetName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// Returns the index into the list of siblings for the element at @pos as
// returned by gtk_widget_path_iter_get_siblings(). If that function would
// return %NULL because the element at @pos has no siblings, this function
// will return 0.
/*

C function

gtk_widget_path_iter_get_sibling_index
*/
func (recv *WidgetPath) IterGetSiblingIndex(pos int32) uint32 {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_sibling_index((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (uint32)(retC)

	return retGo
}

// Returns the list of siblings for the element at @pos. If the element
// was not added with siblings, %NULL is returned.
/*

C function

gtk_widget_path_iter_get_siblings
*/
func (recv *WidgetPath) IterGetSiblings(pos int32) *WidgetPath {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_siblings((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

/*

C type

GtkWidgetPrivate
*/
type WidgetPrivate struct {
	native *C.GtkWidgetPrivate
}

func WidgetPrivateNewFromC(u unsafe.Pointer) *WidgetPrivate {
	c := (*C.GtkWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WidgetPrivate{native: c}

	return g
}

func (recv *WidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowAccessibleClass
*/
type WindowAccessibleClass struct {
	native *C.GtkWindowAccessibleClass
	// parent_class : record
}

func WindowAccessibleClassNewFromC(u unsafe.Pointer) *WindowAccessibleClass {
	c := (*C.GtkWindowAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &WindowAccessibleClass{native: c}

	return g
}

func (recv *WindowAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowAccessiblePrivate
*/
type WindowAccessiblePrivate struct {
	native *C.GtkWindowAccessiblePrivate
}

func WindowAccessiblePrivateNewFromC(u unsafe.Pointer) *WindowAccessiblePrivate {
	c := (*C.GtkWindowAccessiblePrivate)(u)
	if c == nil {
		return nil
	}

	g := &WindowAccessiblePrivate{native: c}

	return g
}

func (recv *WindowAccessiblePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowClass
*/
type WindowClass struct {
	native *C.GtkWindowClass
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
	c := (*C.GtkWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &WindowClass{native: c}

	return g
}

func (recv *WindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowGeometryInfo
*/
type WindowGeometryInfo struct {
	native *C.GtkWindowGeometryInfo
}

func WindowGeometryInfoNewFromC(u unsafe.Pointer) *WindowGeometryInfo {
	c := (*C.GtkWindowGeometryInfo)(u)
	if c == nil {
		return nil
	}

	g := &WindowGeometryInfo{native: c}

	return g
}

func (recv *WindowGeometryInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowGroupClass
*/
type WindowGroupClass struct {
	native *C.GtkWindowGroupClass
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func WindowGroupClassNewFromC(u unsafe.Pointer) *WindowGroupClass {
	c := (*C.GtkWindowGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &WindowGroupClass{native: c}

	return g
}

func (recv *WindowGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowGroupPrivate
*/
type WindowGroupPrivate struct {
	native *C.GtkWindowGroupPrivate
}

func WindowGroupPrivateNewFromC(u unsafe.Pointer) *WindowGroupPrivate {
	c := (*C.GtkWindowGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WindowGroupPrivate{native: c}

	return g
}

func (recv *WindowGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GtkWindowPrivate
*/
type WindowPrivate struct {
	native *C.GtkWindowPrivate
}

func WindowPrivateNewFromC(u unsafe.Pointer) *WindowPrivate {
	c := (*C.GtkWindowPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WindowPrivate{native: c}

	return g
}

func (recv *WindowPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
