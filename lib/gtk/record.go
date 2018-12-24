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
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AboutDialogClass is a wrapper around the C record GtkAboutDialogClass.
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

// Equals compares this AboutDialogClass with another AboutDialogClass, and returns true if they represent the same GObject.
func (recv *AboutDialogClass) Equals(other *AboutDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// AboutDialogPrivate is a wrapper around the C record GtkAboutDialogPrivate.
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

// Equals compares this AboutDialogPrivate with another AboutDialogPrivate, and returns true if they represent the same GObject.
func (recv *AboutDialogPrivate) Equals(other *AboutDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AccelGroupClass is a wrapper around the C record GtkAccelGroupClass.
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

// Equals compares this AccelGroupClass with another AccelGroupClass, and returns true if they represent the same GObject.
func (recv *AccelGroupClass) Equals(other *AccelGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// AccelGroupEntry is a wrapper around the C record GtkAccelGroupEntry.
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

// Equals compares this AccelGroupEntry with another AccelGroupEntry, and returns true if they represent the same GObject.
func (recv *AccelGroupEntry) Equals(other *AccelGroupEntry) bool {
	return other.ToC() == recv.ToC()
}

// AccelGroupPrivate is a wrapper around the C record GtkAccelGroupPrivate.
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

// Equals compares this AccelGroupPrivate with another AccelGroupPrivate, and returns true if they represent the same GObject.
func (recv *AccelGroupPrivate) Equals(other *AccelGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AccelKey is a wrapper around the C record GtkAccelKey.
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

// Equals compares this AccelKey with another AccelKey, and returns true if they represent the same GObject.
func (recv *AccelKey) Equals(other *AccelKey) bool {
	return other.ToC() == recv.ToC()
}

// AccelLabelClass is a wrapper around the C record GtkAccelLabelClass.
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

// Equals compares this AccelLabelClass with another AccelLabelClass, and returns true if they represent the same GObject.
func (recv *AccelLabelClass) Equals(other *AccelLabelClass) bool {
	return other.ToC() == recv.ToC()
}

// AccelLabelPrivate is a wrapper around the C record GtkAccelLabelPrivate.
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

// Equals compares this AccelLabelPrivate with another AccelLabelPrivate, and returns true if they represent the same GObject.
func (recv *AccelLabelPrivate) Equals(other *AccelLabelPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AccelMapClass is a wrapper around the C record GtkAccelMapClass.
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

// Equals compares this AccelMapClass with another AccelMapClass, and returns true if they represent the same GObject.
func (recv *AccelMapClass) Equals(other *AccelMapClass) bool {
	return other.ToC() == recv.ToC()
}

// AccessibleClass is a wrapper around the C record GtkAccessibleClass.
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

// Equals compares this AccessibleClass with another AccessibleClass, and returns true if they represent the same GObject.
func (recv *AccessibleClass) Equals(other *AccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// AccessiblePrivate is a wrapper around the C record GtkAccessiblePrivate.
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

// Equals compares this AccessiblePrivate with another AccessiblePrivate, and returns true if they represent the same GObject.
func (recv *AccessiblePrivate) Equals(other *AccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ActionBarClass is a wrapper around the C record GtkActionBarClass.
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

// Equals compares this ActionBarClass with another ActionBarClass, and returns true if they represent the same GObject.
func (recv *ActionBarClass) Equals(other *ActionBarClass) bool {
	return other.ToC() == recv.ToC()
}

// ActionBarPrivate is a wrapper around the C record GtkActionBarPrivate.
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

// Equals compares this ActionBarPrivate with another ActionBarPrivate, and returns true if they represent the same GObject.
func (recv *ActionBarPrivate) Equals(other *ActionBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ActionClass is a wrapper around the C record GtkActionClass.
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

// Equals compares this ActionClass with another ActionClass, and returns true if they represent the same GObject.
func (recv *ActionClass) Equals(other *ActionClass) bool {
	return other.ToC() == recv.ToC()
}

// ActionEntry is a wrapper around the C record GtkActionEntry.
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

// Equals compares this ActionEntry with another ActionEntry, and returns true if they represent the same GObject.
func (recv *ActionEntry) Equals(other *ActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// ActionGroupClass is a wrapper around the C record GtkActionGroupClass.
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

// Equals compares this ActionGroupClass with another ActionGroupClass, and returns true if they represent the same GObject.
func (recv *ActionGroupClass) Equals(other *ActionGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// ActionGroupPrivate is a wrapper around the C record GtkActionGroupPrivate.
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

// Equals compares this ActionGroupPrivate with another ActionGroupPrivate, and returns true if they represent the same GObject.
func (recv *ActionGroupPrivate) Equals(other *ActionGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ActionPrivate is a wrapper around the C record GtkActionPrivate.
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

// Equals compares this ActionPrivate with another ActionPrivate, and returns true if they represent the same GObject.
func (recv *ActionPrivate) Equals(other *ActionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ActionableInterface is a wrapper around the C record GtkActionableInterface.
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

// Equals compares this ActionableInterface with another ActionableInterface, and returns true if they represent the same GObject.
func (recv *ActionableInterface) Equals(other *ActionableInterface) bool {
	return other.ToC() == recv.ToC()
}

// AdjustmentClass is a wrapper around the C record GtkAdjustmentClass.
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

// Equals compares this AdjustmentClass with another AdjustmentClass, and returns true if they represent the same GObject.
func (recv *AdjustmentClass) Equals(other *AdjustmentClass) bool {
	return other.ToC() == recv.ToC()
}

// AdjustmentPrivate is a wrapper around the C record GtkAdjustmentPrivate.
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

// Equals compares this AdjustmentPrivate with another AdjustmentPrivate, and returns true if they represent the same GObject.
func (recv *AdjustmentPrivate) Equals(other *AdjustmentPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AlignmentClass is a wrapper around the C record GtkAlignmentClass.
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

// Equals compares this AlignmentClass with another AlignmentClass, and returns true if they represent the same GObject.
func (recv *AlignmentClass) Equals(other *AlignmentClass) bool {
	return other.ToC() == recv.ToC()
}

// AlignmentPrivate is a wrapper around the C record GtkAlignmentPrivate.
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

// Equals compares this AlignmentPrivate with another AlignmentPrivate, and returns true if they represent the same GObject.
func (recv *AlignmentPrivate) Equals(other *AlignmentPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserButtonClass is a wrapper around the C record GtkAppChooserButtonClass.
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

// Equals compares this AppChooserButtonClass with another AppChooserButtonClass, and returns true if they represent the same GObject.
func (recv *AppChooserButtonClass) Equals(other *AppChooserButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserButtonPrivate is a wrapper around the C record GtkAppChooserButtonPrivate.
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

// Equals compares this AppChooserButtonPrivate with another AppChooserButtonPrivate, and returns true if they represent the same GObject.
func (recv *AppChooserButtonPrivate) Equals(other *AppChooserButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserDialogClass is a wrapper around the C record GtkAppChooserDialogClass.
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

// Equals compares this AppChooserDialogClass with another AppChooserDialogClass, and returns true if they represent the same GObject.
func (recv *AppChooserDialogClass) Equals(other *AppChooserDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserDialogPrivate is a wrapper around the C record GtkAppChooserDialogPrivate.
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

// Equals compares this AppChooserDialogPrivate with another AppChooserDialogPrivate, and returns true if they represent the same GObject.
func (recv *AppChooserDialogPrivate) Equals(other *AppChooserDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserWidgetClass is a wrapper around the C record GtkAppChooserWidgetClass.
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

// Equals compares this AppChooserWidgetClass with another AppChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *AppChooserWidgetClass) Equals(other *AppChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// AppChooserWidgetPrivate is a wrapper around the C record GtkAppChooserWidgetPrivate.
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

// Equals compares this AppChooserWidgetPrivate with another AppChooserWidgetPrivate, and returns true if they represent the same GObject.
func (recv *AppChooserWidgetPrivate) Equals(other *AppChooserWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationClass is a wrapper around the C record GtkApplicationClass.
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

// Equals compares this ApplicationClass with another ApplicationClass, and returns true if they represent the same GObject.
func (recv *ApplicationClass) Equals(other *ApplicationClass) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationPrivate is a wrapper around the C record GtkApplicationPrivate.
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

// Equals compares this ApplicationPrivate with another ApplicationPrivate, and returns true if they represent the same GObject.
func (recv *ApplicationPrivate) Equals(other *ApplicationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationWindowClass is a wrapper around the C record GtkApplicationWindowClass.
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

// Equals compares this ApplicationWindowClass with another ApplicationWindowClass, and returns true if they represent the same GObject.
func (recv *ApplicationWindowClass) Equals(other *ApplicationWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationWindowPrivate is a wrapper around the C record GtkApplicationWindowPrivate.
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

// Equals compares this ApplicationWindowPrivate with another ApplicationWindowPrivate, and returns true if they represent the same GObject.
func (recv *ApplicationWindowPrivate) Equals(other *ApplicationWindowPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ArrowAccessibleClass is a wrapper around the C record GtkArrowAccessibleClass.
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

// Equals compares this ArrowAccessibleClass with another ArrowAccessibleClass, and returns true if they represent the same GObject.
func (recv *ArrowAccessibleClass) Equals(other *ArrowAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ArrowAccessiblePrivate is a wrapper around the C record GtkArrowAccessiblePrivate.
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

// Equals compares this ArrowAccessiblePrivate with another ArrowAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ArrowAccessiblePrivate) Equals(other *ArrowAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ArrowClass is a wrapper around the C record GtkArrowClass.
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

// Equals compares this ArrowClass with another ArrowClass, and returns true if they represent the same GObject.
func (recv *ArrowClass) Equals(other *ArrowClass) bool {
	return other.ToC() == recv.ToC()
}

// ArrowPrivate is a wrapper around the C record GtkArrowPrivate.
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

// Equals compares this ArrowPrivate with another ArrowPrivate, and returns true if they represent the same GObject.
func (recv *ArrowPrivate) Equals(other *ArrowPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AspectFrameClass is a wrapper around the C record GtkAspectFrameClass.
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

// Equals compares this AspectFrameClass with another AspectFrameClass, and returns true if they represent the same GObject.
func (recv *AspectFrameClass) Equals(other *AspectFrameClass) bool {
	return other.ToC() == recv.ToC()
}

// AspectFramePrivate is a wrapper around the C record GtkAspectFramePrivate.
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

// Equals compares this AspectFramePrivate with another AspectFramePrivate, and returns true if they represent the same GObject.
func (recv *AspectFramePrivate) Equals(other *AspectFramePrivate) bool {
	return other.ToC() == recv.ToC()
}

// AssistantClass is a wrapper around the C record GtkAssistantClass.
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

// Equals compares this AssistantClass with another AssistantClass, and returns true if they represent the same GObject.
func (recv *AssistantClass) Equals(other *AssistantClass) bool {
	return other.ToC() == recv.ToC()
}

// AssistantPrivate is a wrapper around the C record GtkAssistantPrivate.
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

// Equals compares this AssistantPrivate with another AssistantPrivate, and returns true if they represent the same GObject.
func (recv *AssistantPrivate) Equals(other *AssistantPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BinClass is a wrapper around the C record GtkBinClass.
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

// Equals compares this BinClass with another BinClass, and returns true if they represent the same GObject.
func (recv *BinClass) Equals(other *BinClass) bool {
	return other.ToC() == recv.ToC()
}

// BinPrivate is a wrapper around the C record GtkBinPrivate.
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

// Equals compares this BinPrivate with another BinPrivate, and returns true if they represent the same GObject.
func (recv *BinPrivate) Equals(other *BinPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BindingArg is a wrapper around the C record GtkBindingArg.
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

// Equals compares this BindingArg with another BindingArg, and returns true if they represent the same GObject.
func (recv *BindingArg) Equals(other *BindingArg) bool {
	return other.ToC() == recv.ToC()
}

// BindingEntry is a wrapper around the C record GtkBindingEntry.
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

// Equals compares this BindingEntry with another BindingEntry, and returns true if they represent the same GObject.
func (recv *BindingEntry) Equals(other *BindingEntry) bool {
	return other.ToC() == recv.ToC()
}

// gtk_binding_entry_add_signal : unsupported parameter ... : varargs
// BindingEntryAddSignall is a wrapper around the C function gtk_binding_entry_add_signall.
func BindingEntryAddSignall(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType, signalName string, bindingArgs *glib.SList) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_signal_name := C.CString(signalName)
	defer C.free(unsafe.Pointer(c_signal_name))

	c_binding_args := (*C.GSList)(C.NULL)
	if bindingArgs != nil {
		c_binding_args = (*C.GSList)(bindingArgs.ToC())
	}

	C.gtk_binding_entry_add_signall(c_binding_set, c_keyval, c_modifiers, c_signal_name, c_binding_args)

	return
}

// BindingEntryRemove is a wrapper around the C function gtk_binding_entry_remove.
func BindingEntryRemove(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_remove(c_binding_set, c_keyval, c_modifiers)

	return
}

// BindingSet is a wrapper around the C record GtkBindingSet.
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

// Equals compares this BindingSet with another BindingSet, and returns true if they represent the same GObject.
func (recv *BindingSet) Equals(other *BindingSet) bool {
	return other.ToC() == recv.ToC()
}

// BindingSetByClass is a wrapper around the C function gtk_binding_set_by_class.
func BindingSetByClass(objectClass uintptr) *BindingSet {
	c_object_class := (C.gpointer)(objectClass)

	retC := C.gtk_binding_set_by_class(c_object_class)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindingSetFind is a wrapper around the C function gtk_binding_set_find.
func BindingSetFind(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_find(c_set_name)
	var retGo (*BindingSet)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BindingSetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// BindingSetNew is a wrapper around the C function gtk_binding_set_new.
func BindingSetNew(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_new(c_set_name)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function gtk_binding_set_activate.
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

// AddPath is a wrapper around the C function gtk_binding_set_add_path.
func (recv *BindingSet) AddPath(pathType PathType, pathPattern string, priority PathPriorityType) {
	c_path_type := (C.GtkPathType)(pathType)

	c_path_pattern := C.CString(pathPattern)
	defer C.free(unsafe.Pointer(c_path_pattern))

	c_priority := (C.GtkPathPriorityType)(priority)

	C.gtk_binding_set_add_path((*C.GtkBindingSet)(recv.native), c_path_type, c_path_pattern, c_priority)

	return
}

// BindingSignal is a wrapper around the C record GtkBindingSignal.
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

// Equals compares this BindingSignal with another BindingSignal, and returns true if they represent the same GObject.
func (recv *BindingSignal) Equals(other *BindingSignal) bool {
	return other.ToC() == recv.ToC()
}

// BooleanCellAccessibleClass is a wrapper around the C record GtkBooleanCellAccessibleClass.
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

// Equals compares this BooleanCellAccessibleClass with another BooleanCellAccessibleClass, and returns true if they represent the same GObject.
func (recv *BooleanCellAccessibleClass) Equals(other *BooleanCellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// BooleanCellAccessiblePrivate is a wrapper around the C record GtkBooleanCellAccessiblePrivate.
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

// Equals compares this BooleanCellAccessiblePrivate with another BooleanCellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *BooleanCellAccessiblePrivate) Equals(other *BooleanCellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// Border is a wrapper around the C record GtkBorder.
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

// Equals compares this Border with another Border, and returns true if they represent the same GObject.
func (recv *Border) Equals(other *Border) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gtk_border_copy.
func (recv *Border) Copy() *Border {
	retC := C.gtk_border_copy((*C.GtkBorder)(recv.native))
	retGo := BorderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_border_free.
func (recv *Border) Free() {
	C.gtk_border_free((*C.GtkBorder)(recv.native))

	return
}

// BoxClass is a wrapper around the C record GtkBoxClass.
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

// Equals compares this BoxClass with another BoxClass, and returns true if they represent the same GObject.
func (recv *BoxClass) Equals(other *BoxClass) bool {
	return other.ToC() == recv.ToC()
}

// BoxPrivate is a wrapper around the C record GtkBoxPrivate.
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

// Equals compares this BoxPrivate with another BoxPrivate, and returns true if they represent the same GObject.
func (recv *BoxPrivate) Equals(other *BoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BuildableIface is a wrapper around the C record GtkBuildableIface.
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

// Equals compares this BuildableIface with another BuildableIface, and returns true if they represent the same GObject.
func (recv *BuildableIface) Equals(other *BuildableIface) bool {
	return other.ToC() == recv.ToC()
}

// BuilderClass is a wrapper around the C record GtkBuilderClass.
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

// Equals compares this BuilderClass with another BuilderClass, and returns true if they represent the same GObject.
func (recv *BuilderClass) Equals(other *BuilderClass) bool {
	return other.ToC() == recv.ToC()
}

// BuilderPrivate is a wrapper around the C record GtkBuilderPrivate.
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

// Equals compares this BuilderPrivate with another BuilderPrivate, and returns true if they represent the same GObject.
func (recv *BuilderPrivate) Equals(other *BuilderPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ButtonAccessibleClass is a wrapper around the C record GtkButtonAccessibleClass.
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

// Equals compares this ButtonAccessibleClass with another ButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *ButtonAccessibleClass) Equals(other *ButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ButtonAccessiblePrivate is a wrapper around the C record GtkButtonAccessiblePrivate.
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

// Equals compares this ButtonAccessiblePrivate with another ButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ButtonAccessiblePrivate) Equals(other *ButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ButtonBoxClass is a wrapper around the C record GtkButtonBoxClass.
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

// Equals compares this ButtonBoxClass with another ButtonBoxClass, and returns true if they represent the same GObject.
func (recv *ButtonBoxClass) Equals(other *ButtonBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// ButtonBoxPrivate is a wrapper around the C record GtkButtonBoxPrivate.
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

// Equals compares this ButtonBoxPrivate with another ButtonBoxPrivate, and returns true if they represent the same GObject.
func (recv *ButtonBoxPrivate) Equals(other *ButtonBoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ButtonClass is a wrapper around the C record GtkButtonClass.
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

// Equals compares this ButtonClass with another ButtonClass, and returns true if they represent the same GObject.
func (recv *ButtonClass) Equals(other *ButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ButtonPrivate is a wrapper around the C record GtkButtonPrivate.
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

// Equals compares this ButtonPrivate with another ButtonPrivate, and returns true if they represent the same GObject.
func (recv *ButtonPrivate) Equals(other *ButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CalendarClass is a wrapper around the C record GtkCalendarClass.
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

// Equals compares this CalendarClass with another CalendarClass, and returns true if they represent the same GObject.
func (recv *CalendarClass) Equals(other *CalendarClass) bool {
	return other.ToC() == recv.ToC()
}

// CalendarPrivate is a wrapper around the C record GtkCalendarPrivate.
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

// Equals compares this CalendarPrivate with another CalendarPrivate, and returns true if they represent the same GObject.
func (recv *CalendarPrivate) Equals(other *CalendarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellAccessibleClass is a wrapper around the C record GtkCellAccessibleClass.
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

// Equals compares this CellAccessibleClass with another CellAccessibleClass, and returns true if they represent the same GObject.
func (recv *CellAccessibleClass) Equals(other *CellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// CellAccessibleParentIface is a wrapper around the C record GtkCellAccessibleParentIface.
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

// Equals compares this CellAccessibleParentIface with another CellAccessibleParentIface, and returns true if they represent the same GObject.
func (recv *CellAccessibleParentIface) Equals(other *CellAccessibleParentIface) bool {
	return other.ToC() == recv.ToC()
}

// CellAccessiblePrivate is a wrapper around the C record GtkCellAccessiblePrivate.
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

// Equals compares this CellAccessiblePrivate with another CellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *CellAccessiblePrivate) Equals(other *CellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaBoxClass is a wrapper around the C record GtkCellAreaBoxClass.
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

// Equals compares this CellAreaBoxClass with another CellAreaBoxClass, and returns true if they represent the same GObject.
func (recv *CellAreaBoxClass) Equals(other *CellAreaBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaBoxPrivate is a wrapper around the C record GtkCellAreaBoxPrivate.
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

// Equals compares this CellAreaBoxPrivate with another CellAreaBoxPrivate, and returns true if they represent the same GObject.
func (recv *CellAreaBoxPrivate) Equals(other *CellAreaBoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaClass is a wrapper around the C record GtkCellAreaClass.
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

// Equals compares this CellAreaClass with another CellAreaClass, and returns true if they represent the same GObject.
func (recv *CellAreaClass) Equals(other *CellAreaClass) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaContextClass is a wrapper around the C record GtkCellAreaContextClass.
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

// Equals compares this CellAreaContextClass with another CellAreaContextClass, and returns true if they represent the same GObject.
func (recv *CellAreaContextClass) Equals(other *CellAreaContextClass) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaContextPrivate is a wrapper around the C record GtkCellAreaContextPrivate.
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

// Equals compares this CellAreaContextPrivate with another CellAreaContextPrivate, and returns true if they represent the same GObject.
func (recv *CellAreaContextPrivate) Equals(other *CellAreaContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellAreaPrivate is a wrapper around the C record GtkCellAreaPrivate.
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

// Equals compares this CellAreaPrivate with another CellAreaPrivate, and returns true if they represent the same GObject.
func (recv *CellAreaPrivate) Equals(other *CellAreaPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellEditableIface is a wrapper around the C record GtkCellEditableIface.
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

// Equals compares this CellEditableIface with another CellEditableIface, and returns true if they represent the same GObject.
func (recv *CellEditableIface) Equals(other *CellEditableIface) bool {
	return other.ToC() == recv.ToC()
}

// CellLayoutIface is a wrapper around the C record GtkCellLayoutIface.
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

// Equals compares this CellLayoutIface with another CellLayoutIface, and returns true if they represent the same GObject.
func (recv *CellLayoutIface) Equals(other *CellLayoutIface) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererAccelClass is a wrapper around the C record GtkCellRendererAccelClass.
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

// Equals compares this CellRendererAccelClass with another CellRendererAccelClass, and returns true if they represent the same GObject.
func (recv *CellRendererAccelClass) Equals(other *CellRendererAccelClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererAccelPrivate is a wrapper around the C record GtkCellRendererAccelPrivate.
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

// Equals compares this CellRendererAccelPrivate with another CellRendererAccelPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererAccelPrivate) Equals(other *CellRendererAccelPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererClass is a wrapper around the C record GtkCellRendererClass.
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

// Equals compares this CellRendererClass with another CellRendererClass, and returns true if they represent the same GObject.
func (recv *CellRendererClass) Equals(other *CellRendererClass) bool {
	return other.ToC() == recv.ToC()
}

// SetAccessibleType is a wrapper around the C function gtk_cell_renderer_class_set_accessible_type.
func (recv *CellRendererClass) SetAccessibleType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_cell_renderer_class_set_accessible_type((*C.GtkCellRendererClass)(recv.native), c_type)

	return
}

// CellRendererClassPrivate is a wrapper around the C record GtkCellRendererClassPrivate.
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

// Equals compares this CellRendererClassPrivate with another CellRendererClassPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererClassPrivate) Equals(other *CellRendererClassPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererComboClass is a wrapper around the C record GtkCellRendererComboClass.
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

// Equals compares this CellRendererComboClass with another CellRendererComboClass, and returns true if they represent the same GObject.
func (recv *CellRendererComboClass) Equals(other *CellRendererComboClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererComboPrivate is a wrapper around the C record GtkCellRendererComboPrivate.
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

// Equals compares this CellRendererComboPrivate with another CellRendererComboPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererComboPrivate) Equals(other *CellRendererComboPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererPixbufClass is a wrapper around the C record GtkCellRendererPixbufClass.
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

// Equals compares this CellRendererPixbufClass with another CellRendererPixbufClass, and returns true if they represent the same GObject.
func (recv *CellRendererPixbufClass) Equals(other *CellRendererPixbufClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererPixbufPrivate is a wrapper around the C record GtkCellRendererPixbufPrivate.
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

// Equals compares this CellRendererPixbufPrivate with another CellRendererPixbufPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererPixbufPrivate) Equals(other *CellRendererPixbufPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererPrivate is a wrapper around the C record GtkCellRendererPrivate.
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

// Equals compares this CellRendererPrivate with another CellRendererPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererPrivate) Equals(other *CellRendererPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererProgressClass is a wrapper around the C record GtkCellRendererProgressClass.
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

// Equals compares this CellRendererProgressClass with another CellRendererProgressClass, and returns true if they represent the same GObject.
func (recv *CellRendererProgressClass) Equals(other *CellRendererProgressClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererProgressPrivate is a wrapper around the C record GtkCellRendererProgressPrivate.
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

// Equals compares this CellRendererProgressPrivate with another CellRendererProgressPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererProgressPrivate) Equals(other *CellRendererProgressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererSpinClass is a wrapper around the C record GtkCellRendererSpinClass.
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

// Equals compares this CellRendererSpinClass with another CellRendererSpinClass, and returns true if they represent the same GObject.
func (recv *CellRendererSpinClass) Equals(other *CellRendererSpinClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererSpinPrivate is a wrapper around the C record GtkCellRendererSpinPrivate.
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

// Equals compares this CellRendererSpinPrivate with another CellRendererSpinPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererSpinPrivate) Equals(other *CellRendererSpinPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererSpinnerClass is a wrapper around the C record GtkCellRendererSpinnerClass.
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

// Equals compares this CellRendererSpinnerClass with another CellRendererSpinnerClass, and returns true if they represent the same GObject.
func (recv *CellRendererSpinnerClass) Equals(other *CellRendererSpinnerClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererSpinnerPrivate is a wrapper around the C record GtkCellRendererSpinnerPrivate.
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

// Equals compares this CellRendererSpinnerPrivate with another CellRendererSpinnerPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererSpinnerPrivate) Equals(other *CellRendererSpinnerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererTextClass is a wrapper around the C record GtkCellRendererTextClass.
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

// Equals compares this CellRendererTextClass with another CellRendererTextClass, and returns true if they represent the same GObject.
func (recv *CellRendererTextClass) Equals(other *CellRendererTextClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererTextPrivate is a wrapper around the C record GtkCellRendererTextPrivate.
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

// Equals compares this CellRendererTextPrivate with another CellRendererTextPrivate, and returns true if they represent the same GObject.
func (recv *CellRendererTextPrivate) Equals(other *CellRendererTextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererToggleClass is a wrapper around the C record GtkCellRendererToggleClass.
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

// Equals compares this CellRendererToggleClass with another CellRendererToggleClass, and returns true if they represent the same GObject.
func (recv *CellRendererToggleClass) Equals(other *CellRendererToggleClass) bool {
	return other.ToC() == recv.ToC()
}

// CellRendererTogglePrivate is a wrapper around the C record GtkCellRendererTogglePrivate.
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

// Equals compares this CellRendererTogglePrivate with another CellRendererTogglePrivate, and returns true if they represent the same GObject.
func (recv *CellRendererTogglePrivate) Equals(other *CellRendererTogglePrivate) bool {
	return other.ToC() == recv.ToC()
}

// CellViewClass is a wrapper around the C record GtkCellViewClass.
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

// Equals compares this CellViewClass with another CellViewClass, and returns true if they represent the same GObject.
func (recv *CellViewClass) Equals(other *CellViewClass) bool {
	return other.ToC() == recv.ToC()
}

// CellViewPrivate is a wrapper around the C record GtkCellViewPrivate.
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

// Equals compares this CellViewPrivate with another CellViewPrivate, and returns true if they represent the same GObject.
func (recv *CellViewPrivate) Equals(other *CellViewPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CheckButtonClass is a wrapper around the C record GtkCheckButtonClass.
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

// Equals compares this CheckButtonClass with another CheckButtonClass, and returns true if they represent the same GObject.
func (recv *CheckButtonClass) Equals(other *CheckButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// CheckMenuItemAccessibleClass is a wrapper around the C record GtkCheckMenuItemAccessibleClass.
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

// Equals compares this CheckMenuItemAccessibleClass with another CheckMenuItemAccessibleClass, and returns true if they represent the same GObject.
func (recv *CheckMenuItemAccessibleClass) Equals(other *CheckMenuItemAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// CheckMenuItemAccessiblePrivate is a wrapper around the C record GtkCheckMenuItemAccessiblePrivate.
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

// Equals compares this CheckMenuItemAccessiblePrivate with another CheckMenuItemAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *CheckMenuItemAccessiblePrivate) Equals(other *CheckMenuItemAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// CheckMenuItemClass is a wrapper around the C record GtkCheckMenuItemClass.
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

// Equals compares this CheckMenuItemClass with another CheckMenuItemClass, and returns true if they represent the same GObject.
func (recv *CheckMenuItemClass) Equals(other *CheckMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// CheckMenuItemPrivate is a wrapper around the C record GtkCheckMenuItemPrivate.
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

// Equals compares this CheckMenuItemPrivate with another CheckMenuItemPrivate, and returns true if they represent the same GObject.
func (recv *CheckMenuItemPrivate) Equals(other *CheckMenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorButtonClass is a wrapper around the C record GtkColorButtonClass.
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

// Equals compares this ColorButtonClass with another ColorButtonClass, and returns true if they represent the same GObject.
func (recv *ColorButtonClass) Equals(other *ColorButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorButtonPrivate is a wrapper around the C record GtkColorButtonPrivate.
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

// Equals compares this ColorButtonPrivate with another ColorButtonPrivate, and returns true if they represent the same GObject.
func (recv *ColorButtonPrivate) Equals(other *ColorButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserDialogClass is a wrapper around the C record GtkColorChooserDialogClass.
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

// Equals compares this ColorChooserDialogClass with another ColorChooserDialogClass, and returns true if they represent the same GObject.
func (recv *ColorChooserDialogClass) Equals(other *ColorChooserDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserDialogPrivate is a wrapper around the C record GtkColorChooserDialogPrivate.
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

// Equals compares this ColorChooserDialogPrivate with another ColorChooserDialogPrivate, and returns true if they represent the same GObject.
func (recv *ColorChooserDialogPrivate) Equals(other *ColorChooserDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserInterface is a wrapper around the C record GtkColorChooserInterface.
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

// Equals compares this ColorChooserInterface with another ColorChooserInterface, and returns true if they represent the same GObject.
func (recv *ColorChooserInterface) Equals(other *ColorChooserInterface) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserWidgetClass is a wrapper around the C record GtkColorChooserWidgetClass.
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

// Equals compares this ColorChooserWidgetClass with another ColorChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *ColorChooserWidgetClass) Equals(other *ColorChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserWidgetPrivate is a wrapper around the C record GtkColorChooserWidgetPrivate.
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

// Equals compares this ColorChooserWidgetPrivate with another ColorChooserWidgetPrivate, and returns true if they represent the same GObject.
func (recv *ColorChooserWidgetPrivate) Equals(other *ColorChooserWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorSelectionClass is a wrapper around the C record GtkColorSelectionClass.
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

// Equals compares this ColorSelectionClass with another ColorSelectionClass, and returns true if they represent the same GObject.
func (recv *ColorSelectionClass) Equals(other *ColorSelectionClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorSelectionDialogClass is a wrapper around the C record GtkColorSelectionDialogClass.
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

// Equals compares this ColorSelectionDialogClass with another ColorSelectionDialogClass, and returns true if they represent the same GObject.
func (recv *ColorSelectionDialogClass) Equals(other *ColorSelectionDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorSelectionDialogPrivate is a wrapper around the C record GtkColorSelectionDialogPrivate.
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

// Equals compares this ColorSelectionDialogPrivate with another ColorSelectionDialogPrivate, and returns true if they represent the same GObject.
func (recv *ColorSelectionDialogPrivate) Equals(other *ColorSelectionDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorSelectionPrivate is a wrapper around the C record GtkColorSelectionPrivate.
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

// Equals compares this ColorSelectionPrivate with another ColorSelectionPrivate, and returns true if they represent the same GObject.
func (recv *ColorSelectionPrivate) Equals(other *ColorSelectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxAccessibleClass is a wrapper around the C record GtkComboBoxAccessibleClass.
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

// Equals compares this ComboBoxAccessibleClass with another ComboBoxAccessibleClass, and returns true if they represent the same GObject.
func (recv *ComboBoxAccessibleClass) Equals(other *ComboBoxAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxAccessiblePrivate is a wrapper around the C record GtkComboBoxAccessiblePrivate.
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

// Equals compares this ComboBoxAccessiblePrivate with another ComboBoxAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ComboBoxAccessiblePrivate) Equals(other *ComboBoxAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxClass is a wrapper around the C record GtkComboBoxClass.
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

// Equals compares this ComboBoxClass with another ComboBoxClass, and returns true if they represent the same GObject.
func (recv *ComboBoxClass) Equals(other *ComboBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxPrivate is a wrapper around the C record GtkComboBoxPrivate.
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

// Equals compares this ComboBoxPrivate with another ComboBoxPrivate, and returns true if they represent the same GObject.
func (recv *ComboBoxPrivate) Equals(other *ComboBoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxTextClass is a wrapper around the C record GtkComboBoxTextClass.
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

// Equals compares this ComboBoxTextClass with another ComboBoxTextClass, and returns true if they represent the same GObject.
func (recv *ComboBoxTextClass) Equals(other *ComboBoxTextClass) bool {
	return other.ToC() == recv.ToC()
}

// ComboBoxTextPrivate is a wrapper around the C record GtkComboBoxTextPrivate.
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

// Equals compares this ComboBoxTextPrivate with another ComboBoxTextPrivate, and returns true if they represent the same GObject.
func (recv *ComboBoxTextPrivate) Equals(other *ComboBoxTextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContainerAccessibleClass is a wrapper around the C record GtkContainerAccessibleClass.
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

// Equals compares this ContainerAccessibleClass with another ContainerAccessibleClass, and returns true if they represent the same GObject.
func (recv *ContainerAccessibleClass) Equals(other *ContainerAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ContainerAccessiblePrivate is a wrapper around the C record GtkContainerAccessiblePrivate.
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

// Equals compares this ContainerAccessiblePrivate with another ContainerAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ContainerAccessiblePrivate) Equals(other *ContainerAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContainerCellAccessibleClass is a wrapper around the C record GtkContainerCellAccessibleClass.
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

// Equals compares this ContainerCellAccessibleClass with another ContainerCellAccessibleClass, and returns true if they represent the same GObject.
func (recv *ContainerCellAccessibleClass) Equals(other *ContainerCellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ContainerCellAccessiblePrivate is a wrapper around the C record GtkContainerCellAccessiblePrivate.
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

// Equals compares this ContainerCellAccessiblePrivate with another ContainerCellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ContainerCellAccessiblePrivate) Equals(other *ContainerCellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContainerClass is a wrapper around the C record GtkContainerClass.
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

// Equals compares this ContainerClass with another ContainerClass, and returns true if they represent the same GObject.
func (recv *ContainerClass) Equals(other *ContainerClass) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : gtk_container_class_find_child_property : return type : Blacklisted record : GParamSpec

// HandleBorderWidth is a wrapper around the C function gtk_container_class_handle_border_width.
func (recv *ContainerClass) HandleBorderWidth() {
	C.gtk_container_class_handle_border_width((*C.GtkContainerClass)(recv.native))

	return
}

// Unsupported : gtk_container_class_install_child_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_container_class_list_child_properties : no return type

// ContainerPrivate is a wrapper around the C record GtkContainerPrivate.
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

// Equals compares this ContainerPrivate with another ContainerPrivate, and returns true if they represent the same GObject.
func (recv *ContainerPrivate) Equals(other *ContainerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CssProviderClass is a wrapper around the C record GtkCssProviderClass.
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

// Equals compares this CssProviderClass with another CssProviderClass, and returns true if they represent the same GObject.
func (recv *CssProviderClass) Equals(other *CssProviderClass) bool {
	return other.ToC() == recv.ToC()
}

// CssProviderPrivate is a wrapper around the C record GtkCssProviderPrivate.
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

// Equals compares this CssProviderPrivate with another CssProviderPrivate, and returns true if they represent the same GObject.
func (recv *CssProviderPrivate) Equals(other *CssProviderPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DialogClass is a wrapper around the C record GtkDialogClass.
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

// Equals compares this DialogClass with another DialogClass, and returns true if they represent the same GObject.
func (recv *DialogClass) Equals(other *DialogClass) bool {
	return other.ToC() == recv.ToC()
}

// DialogPrivate is a wrapper around the C record GtkDialogPrivate.
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

// Equals compares this DialogPrivate with another DialogPrivate, and returns true if they represent the same GObject.
func (recv *DialogPrivate) Equals(other *DialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DrawingAreaClass is a wrapper around the C record GtkDrawingAreaClass.
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

// Equals compares this DrawingAreaClass with another DrawingAreaClass, and returns true if they represent the same GObject.
func (recv *DrawingAreaClass) Equals(other *DrawingAreaClass) bool {
	return other.ToC() == recv.ToC()
}

// EditableInterface is a wrapper around the C record GtkEditableInterface.
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

// Equals compares this EditableInterface with another EditableInterface, and returns true if they represent the same GObject.
func (recv *EditableInterface) Equals(other *EditableInterface) bool {
	return other.ToC() == recv.ToC()
}

// EntryAccessibleClass is a wrapper around the C record GtkEntryAccessibleClass.
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

// Equals compares this EntryAccessibleClass with another EntryAccessibleClass, and returns true if they represent the same GObject.
func (recv *EntryAccessibleClass) Equals(other *EntryAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// EntryAccessiblePrivate is a wrapper around the C record GtkEntryAccessiblePrivate.
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

// Equals compares this EntryAccessiblePrivate with another EntryAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *EntryAccessiblePrivate) Equals(other *EntryAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// EntryBufferClass is a wrapper around the C record GtkEntryBufferClass.
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

// Equals compares this EntryBufferClass with another EntryBufferClass, and returns true if they represent the same GObject.
func (recv *EntryBufferClass) Equals(other *EntryBufferClass) bool {
	return other.ToC() == recv.ToC()
}

// EntryBufferPrivate is a wrapper around the C record GtkEntryBufferPrivate.
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

// Equals compares this EntryBufferPrivate with another EntryBufferPrivate, and returns true if they represent the same GObject.
func (recv *EntryBufferPrivate) Equals(other *EntryBufferPrivate) bool {
	return other.ToC() == recv.ToC()
}

// EntryClass is a wrapper around the C record GtkEntryClass.
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

// Equals compares this EntryClass with another EntryClass, and returns true if they represent the same GObject.
func (recv *EntryClass) Equals(other *EntryClass) bool {
	return other.ToC() == recv.ToC()
}

// EntryCompletionClass is a wrapper around the C record GtkEntryCompletionClass.
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

// Equals compares this EntryCompletionClass with another EntryCompletionClass, and returns true if they represent the same GObject.
func (recv *EntryCompletionClass) Equals(other *EntryCompletionClass) bool {
	return other.ToC() == recv.ToC()
}

// EntryCompletionPrivate is a wrapper around the C record GtkEntryCompletionPrivate.
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

// Equals compares this EntryCompletionPrivate with another EntryCompletionPrivate, and returns true if they represent the same GObject.
func (recv *EntryCompletionPrivate) Equals(other *EntryCompletionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// EntryPrivate is a wrapper around the C record GtkEntryPrivate.
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

// Equals compares this EntryPrivate with another EntryPrivate, and returns true if they represent the same GObject.
func (recv *EntryPrivate) Equals(other *EntryPrivate) bool {
	return other.ToC() == recv.ToC()
}

// EventBoxClass is a wrapper around the C record GtkEventBoxClass.
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

// Equals compares this EventBoxClass with another EventBoxClass, and returns true if they represent the same GObject.
func (recv *EventBoxClass) Equals(other *EventBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// EventBoxPrivate is a wrapper around the C record GtkEventBoxPrivate.
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

// Equals compares this EventBoxPrivate with another EventBoxPrivate, and returns true if they represent the same GObject.
func (recv *EventBoxPrivate) Equals(other *EventBoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// EventControllerClass is a wrapper around the C record GtkEventControllerClass.
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

// Equals compares this EventControllerClass with another EventControllerClass, and returns true if they represent the same GObject.
func (recv *EventControllerClass) Equals(other *EventControllerClass) bool {
	return other.ToC() == recv.ToC()
}

// ExpanderAccessibleClass is a wrapper around the C record GtkExpanderAccessibleClass.
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

// Equals compares this ExpanderAccessibleClass with another ExpanderAccessibleClass, and returns true if they represent the same GObject.
func (recv *ExpanderAccessibleClass) Equals(other *ExpanderAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ExpanderAccessiblePrivate is a wrapper around the C record GtkExpanderAccessiblePrivate.
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

// Equals compares this ExpanderAccessiblePrivate with another ExpanderAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ExpanderAccessiblePrivate) Equals(other *ExpanderAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ExpanderClass is a wrapper around the C record GtkExpanderClass.
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

// Equals compares this ExpanderClass with another ExpanderClass, and returns true if they represent the same GObject.
func (recv *ExpanderClass) Equals(other *ExpanderClass) bool {
	return other.ToC() == recv.ToC()
}

// ExpanderPrivate is a wrapper around the C record GtkExpanderPrivate.
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

// Equals compares this ExpanderPrivate with another ExpanderPrivate, and returns true if they represent the same GObject.
func (recv *ExpanderPrivate) Equals(other *ExpanderPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserButtonClass is a wrapper around the C record GtkFileChooserButtonClass.
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

// Equals compares this FileChooserButtonClass with another FileChooserButtonClass, and returns true if they represent the same GObject.
func (recv *FileChooserButtonClass) Equals(other *FileChooserButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserButtonPrivate is a wrapper around the C record GtkFileChooserButtonPrivate.
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

// Equals compares this FileChooserButtonPrivate with another FileChooserButtonPrivate, and returns true if they represent the same GObject.
func (recv *FileChooserButtonPrivate) Equals(other *FileChooserButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserDialogClass is a wrapper around the C record GtkFileChooserDialogClass.
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

// Equals compares this FileChooserDialogClass with another FileChooserDialogClass, and returns true if they represent the same GObject.
func (recv *FileChooserDialogClass) Equals(other *FileChooserDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserDialogPrivate is a wrapper around the C record GtkFileChooserDialogPrivate.
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

// Equals compares this FileChooserDialogPrivate with another FileChooserDialogPrivate, and returns true if they represent the same GObject.
func (recv *FileChooserDialogPrivate) Equals(other *FileChooserDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserWidgetClass is a wrapper around the C record GtkFileChooserWidgetClass.
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

// Equals compares this FileChooserWidgetClass with another FileChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *FileChooserWidgetClass) Equals(other *FileChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserWidgetPrivate is a wrapper around the C record GtkFileChooserWidgetPrivate.
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

// Equals compares this FileChooserWidgetPrivate with another FileChooserWidgetPrivate, and returns true if they represent the same GObject.
func (recv *FileChooserWidgetPrivate) Equals(other *FileChooserWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileFilterInfo is a wrapper around the C record GtkFileFilterInfo.
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

// Equals compares this FileFilterInfo with another FileFilterInfo, and returns true if they represent the same GObject.
func (recv *FileFilterInfo) Equals(other *FileFilterInfo) bool {
	return other.ToC() == recv.ToC()
}

// FixedChild is a wrapper around the C record GtkFixedChild.
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

// Equals compares this FixedChild with another FixedChild, and returns true if they represent the same GObject.
func (recv *FixedChild) Equals(other *FixedChild) bool {
	return other.ToC() == recv.ToC()
}

// FixedClass is a wrapper around the C record GtkFixedClass.
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

// Equals compares this FixedClass with another FixedClass, and returns true if they represent the same GObject.
func (recv *FixedClass) Equals(other *FixedClass) bool {
	return other.ToC() == recv.ToC()
}

// FixedPrivate is a wrapper around the C record GtkFixedPrivate.
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

// Equals compares this FixedPrivate with another FixedPrivate, and returns true if they represent the same GObject.
func (recv *FixedPrivate) Equals(other *FixedPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FlowBoxAccessibleClass is a wrapper around the C record GtkFlowBoxAccessibleClass.
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

// Equals compares this FlowBoxAccessibleClass with another FlowBoxAccessibleClass, and returns true if they represent the same GObject.
func (recv *FlowBoxAccessibleClass) Equals(other *FlowBoxAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// FlowBoxAccessiblePrivate is a wrapper around the C record GtkFlowBoxAccessiblePrivate.
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

// Equals compares this FlowBoxAccessiblePrivate with another FlowBoxAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *FlowBoxAccessiblePrivate) Equals(other *FlowBoxAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// FlowBoxChildAccessibleClass is a wrapper around the C record GtkFlowBoxChildAccessibleClass.
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

// Equals compares this FlowBoxChildAccessibleClass with another FlowBoxChildAccessibleClass, and returns true if they represent the same GObject.
func (recv *FlowBoxChildAccessibleClass) Equals(other *FlowBoxChildAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// FlowBoxChildClass is a wrapper around the C record GtkFlowBoxChildClass.
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

// Equals compares this FlowBoxChildClass with another FlowBoxChildClass, and returns true if they represent the same GObject.
func (recv *FlowBoxChildClass) Equals(other *FlowBoxChildClass) bool {
	return other.ToC() == recv.ToC()
}

// FlowBoxClass is a wrapper around the C record GtkFlowBoxClass.
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

// Equals compares this FlowBoxClass with another FlowBoxClass, and returns true if they represent the same GObject.
func (recv *FlowBoxClass) Equals(other *FlowBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// FontButtonClass is a wrapper around the C record GtkFontButtonClass.
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

// Equals compares this FontButtonClass with another FontButtonClass, and returns true if they represent the same GObject.
func (recv *FontButtonClass) Equals(other *FontButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// FontButtonPrivate is a wrapper around the C record GtkFontButtonPrivate.
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

// Equals compares this FontButtonPrivate with another FontButtonPrivate, and returns true if they represent the same GObject.
func (recv *FontButtonPrivate) Equals(other *FontButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FontChooserDialogClass is a wrapper around the C record GtkFontChooserDialogClass.
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

// Equals compares this FontChooserDialogClass with another FontChooserDialogClass, and returns true if they represent the same GObject.
func (recv *FontChooserDialogClass) Equals(other *FontChooserDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// FontChooserDialogPrivate is a wrapper around the C record GtkFontChooserDialogPrivate.
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

// Equals compares this FontChooserDialogPrivate with another FontChooserDialogPrivate, and returns true if they represent the same GObject.
func (recv *FontChooserDialogPrivate) Equals(other *FontChooserDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FontChooserIface is a wrapper around the C record GtkFontChooserIface.
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

// Equals compares this FontChooserIface with another FontChooserIface, and returns true if they represent the same GObject.
func (recv *FontChooserIface) Equals(other *FontChooserIface) bool {
	return other.ToC() == recv.ToC()
}

// FontChooserWidgetClass is a wrapper around the C record GtkFontChooserWidgetClass.
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

// Equals compares this FontChooserWidgetClass with another FontChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *FontChooserWidgetClass) Equals(other *FontChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// FontChooserWidgetPrivate is a wrapper around the C record GtkFontChooserWidgetPrivate.
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

// Equals compares this FontChooserWidgetPrivate with another FontChooserWidgetPrivate, and returns true if they represent the same GObject.
func (recv *FontChooserWidgetPrivate) Equals(other *FontChooserWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FontSelectionClass is a wrapper around the C record GtkFontSelectionClass.
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

// Equals compares this FontSelectionClass with another FontSelectionClass, and returns true if they represent the same GObject.
func (recv *FontSelectionClass) Equals(other *FontSelectionClass) bool {
	return other.ToC() == recv.ToC()
}

// FontSelectionDialogClass is a wrapper around the C record GtkFontSelectionDialogClass.
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

// Equals compares this FontSelectionDialogClass with another FontSelectionDialogClass, and returns true if they represent the same GObject.
func (recv *FontSelectionDialogClass) Equals(other *FontSelectionDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// FontSelectionDialogPrivate is a wrapper around the C record GtkFontSelectionDialogPrivate.
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

// Equals compares this FontSelectionDialogPrivate with another FontSelectionDialogPrivate, and returns true if they represent the same GObject.
func (recv *FontSelectionDialogPrivate) Equals(other *FontSelectionDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FontSelectionPrivate is a wrapper around the C record GtkFontSelectionPrivate.
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

// Equals compares this FontSelectionPrivate with another FontSelectionPrivate, and returns true if they represent the same GObject.
func (recv *FontSelectionPrivate) Equals(other *FontSelectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FrameAccessibleClass is a wrapper around the C record GtkFrameAccessibleClass.
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

// Equals compares this FrameAccessibleClass with another FrameAccessibleClass, and returns true if they represent the same GObject.
func (recv *FrameAccessibleClass) Equals(other *FrameAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// FrameAccessiblePrivate is a wrapper around the C record GtkFrameAccessiblePrivate.
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

// Equals compares this FrameAccessiblePrivate with another FrameAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *FrameAccessiblePrivate) Equals(other *FrameAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// FrameClass is a wrapper around the C record GtkFrameClass.
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

// Equals compares this FrameClass with another FrameClass, and returns true if they represent the same GObject.
func (recv *FrameClass) Equals(other *FrameClass) bool {
	return other.ToC() == recv.ToC()
}

// FramePrivate is a wrapper around the C record GtkFramePrivate.
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

// Equals compares this FramePrivate with another FramePrivate, and returns true if they represent the same GObject.
func (recv *FramePrivate) Equals(other *FramePrivate) bool {
	return other.ToC() == recv.ToC()
}

// GestureClass is a wrapper around the C record GtkGestureClass.
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

// Equals compares this GestureClass with another GestureClass, and returns true if they represent the same GObject.
func (recv *GestureClass) Equals(other *GestureClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureDragClass is a wrapper around the C record GtkGestureDragClass.
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

// Equals compares this GestureDragClass with another GestureDragClass, and returns true if they represent the same GObject.
func (recv *GestureDragClass) Equals(other *GestureDragClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureLongPressClass is a wrapper around the C record GtkGestureLongPressClass.
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

// Equals compares this GestureLongPressClass with another GestureLongPressClass, and returns true if they represent the same GObject.
func (recv *GestureLongPressClass) Equals(other *GestureLongPressClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureMultiPressClass is a wrapper around the C record GtkGestureMultiPressClass.
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

// Equals compares this GestureMultiPressClass with another GestureMultiPressClass, and returns true if they represent the same GObject.
func (recv *GestureMultiPressClass) Equals(other *GestureMultiPressClass) bool {
	return other.ToC() == recv.ToC()
}

// GesturePanClass is a wrapper around the C record GtkGesturePanClass.
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

// Equals compares this GesturePanClass with another GesturePanClass, and returns true if they represent the same GObject.
func (recv *GesturePanClass) Equals(other *GesturePanClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureRotateClass is a wrapper around the C record GtkGestureRotateClass.
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

// Equals compares this GestureRotateClass with another GestureRotateClass, and returns true if they represent the same GObject.
func (recv *GestureRotateClass) Equals(other *GestureRotateClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureSingleClass is a wrapper around the C record GtkGestureSingleClass.
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

// Equals compares this GestureSingleClass with another GestureSingleClass, and returns true if they represent the same GObject.
func (recv *GestureSingleClass) Equals(other *GestureSingleClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureSwipeClass is a wrapper around the C record GtkGestureSwipeClass.
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

// Equals compares this GestureSwipeClass with another GestureSwipeClass, and returns true if they represent the same GObject.
func (recv *GestureSwipeClass) Equals(other *GestureSwipeClass) bool {
	return other.ToC() == recv.ToC()
}

// GestureZoomClass is a wrapper around the C record GtkGestureZoomClass.
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

// Equals compares this GestureZoomClass with another GestureZoomClass, and returns true if they represent the same GObject.
func (recv *GestureZoomClass) Equals(other *GestureZoomClass) bool {
	return other.ToC() == recv.ToC()
}

// Gradient is a wrapper around the C record GtkGradient.
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

// Equals compares this Gradient with another Gradient, and returns true if they represent the same GObject.
func (recv *Gradient) Equals(other *Gradient) bool {
	return other.ToC() == recv.ToC()
}

// ResolveForContext is a wrapper around the C function gtk_gradient_resolve_for_context.
func (recv *Gradient) ResolveForContext(context *StyleContext) *cairo.Pattern {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	retC := C.gtk_gradient_resolve_for_context((*C.GtkGradient)(recv.native), c_context)
	retGo := cairo.PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function gtk_gradient_to_string.
func (recv *Gradient) ToString() string {
	retC := C.gtk_gradient_to_string((*C.GtkGradient)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GridClass is a wrapper around the C record GtkGridClass.
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

// Equals compares this GridClass with another GridClass, and returns true if they represent the same GObject.
func (recv *GridClass) Equals(other *GridClass) bool {
	return other.ToC() == recv.ToC()
}

// GridPrivate is a wrapper around the C record GtkGridPrivate.
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

// Equals compares this GridPrivate with another GridPrivate, and returns true if they represent the same GObject.
func (recv *GridPrivate) Equals(other *GridPrivate) bool {
	return other.ToC() == recv.ToC()
}

// HBoxClass is a wrapper around the C record GtkHBoxClass.
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

// Equals compares this HBoxClass with another HBoxClass, and returns true if they represent the same GObject.
func (recv *HBoxClass) Equals(other *HBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// HButtonBoxClass is a wrapper around the C record GtkHButtonBoxClass.
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

// Equals compares this HButtonBoxClass with another HButtonBoxClass, and returns true if they represent the same GObject.
func (recv *HButtonBoxClass) Equals(other *HButtonBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// HPanedClass is a wrapper around the C record GtkHPanedClass.
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

// Equals compares this HPanedClass with another HPanedClass, and returns true if they represent the same GObject.
func (recv *HPanedClass) Equals(other *HPanedClass) bool {
	return other.ToC() == recv.ToC()
}

// HSVClass is a wrapper around the C record GtkHSVClass.
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

// Equals compares this HSVClass with another HSVClass, and returns true if they represent the same GObject.
func (recv *HSVClass) Equals(other *HSVClass) bool {
	return other.ToC() == recv.ToC()
}

// HSVPrivate is a wrapper around the C record GtkHSVPrivate.
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

// Equals compares this HSVPrivate with another HSVPrivate, and returns true if they represent the same GObject.
func (recv *HSVPrivate) Equals(other *HSVPrivate) bool {
	return other.ToC() == recv.ToC()
}

// HScaleClass is a wrapper around the C record GtkHScaleClass.
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

// Equals compares this HScaleClass with another HScaleClass, and returns true if they represent the same GObject.
func (recv *HScaleClass) Equals(other *HScaleClass) bool {
	return other.ToC() == recv.ToC()
}

// HScrollbarClass is a wrapper around the C record GtkHScrollbarClass.
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

// Equals compares this HScrollbarClass with another HScrollbarClass, and returns true if they represent the same GObject.
func (recv *HScrollbarClass) Equals(other *HScrollbarClass) bool {
	return other.ToC() == recv.ToC()
}

// HSeparatorClass is a wrapper around the C record GtkHSeparatorClass.
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

// Equals compares this HSeparatorClass with another HSeparatorClass, and returns true if they represent the same GObject.
func (recv *HSeparatorClass) Equals(other *HSeparatorClass) bool {
	return other.ToC() == recv.ToC()
}

// HandleBoxClass is a wrapper around the C record GtkHandleBoxClass.
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

// Equals compares this HandleBoxClass with another HandleBoxClass, and returns true if they represent the same GObject.
func (recv *HandleBoxClass) Equals(other *HandleBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// HandleBoxPrivate is a wrapper around the C record GtkHandleBoxPrivate.
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

// Equals compares this HandleBoxPrivate with another HandleBoxPrivate, and returns true if they represent the same GObject.
func (recv *HandleBoxPrivate) Equals(other *HandleBoxPrivate) bool {
	return other.ToC() == recv.ToC()
}

// HeaderBarClass is a wrapper around the C record GtkHeaderBarClass.
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

// Equals compares this HeaderBarClass with another HeaderBarClass, and returns true if they represent the same GObject.
func (recv *HeaderBarClass) Equals(other *HeaderBarClass) bool {
	return other.ToC() == recv.ToC()
}

// HeaderBarPrivate is a wrapper around the C record GtkHeaderBarPrivate.
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

// Equals compares this HeaderBarPrivate with another HeaderBarPrivate, and returns true if they represent the same GObject.
func (recv *HeaderBarPrivate) Equals(other *HeaderBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// IMContextClass is a wrapper around the C record GtkIMContextClass.
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

// Equals compares this IMContextClass with another IMContextClass, and returns true if they represent the same GObject.
func (recv *IMContextClass) Equals(other *IMContextClass) bool {
	return other.ToC() == recv.ToC()
}

// IMContextInfo is a wrapper around the C record GtkIMContextInfo.
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

// Equals compares this IMContextInfo with another IMContextInfo, and returns true if they represent the same GObject.
func (recv *IMContextInfo) Equals(other *IMContextInfo) bool {
	return other.ToC() == recv.ToC()
}

// IMContextSimpleClass is a wrapper around the C record GtkIMContextSimpleClass.
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

// Equals compares this IMContextSimpleClass with another IMContextSimpleClass, and returns true if they represent the same GObject.
func (recv *IMContextSimpleClass) Equals(other *IMContextSimpleClass) bool {
	return other.ToC() == recv.ToC()
}

// IMContextSimplePrivate is a wrapper around the C record GtkIMContextSimplePrivate.
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

// Equals compares this IMContextSimplePrivate with another IMContextSimplePrivate, and returns true if they represent the same GObject.
func (recv *IMContextSimplePrivate) Equals(other *IMContextSimplePrivate) bool {
	return other.ToC() == recv.ToC()
}

// IMMulticontextClass is a wrapper around the C record GtkIMMulticontextClass.
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

// Equals compares this IMMulticontextClass with another IMMulticontextClass, and returns true if they represent the same GObject.
func (recv *IMMulticontextClass) Equals(other *IMMulticontextClass) bool {
	return other.ToC() == recv.ToC()
}

// IMMulticontextPrivate is a wrapper around the C record GtkIMMulticontextPrivate.
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

// Equals compares this IMMulticontextPrivate with another IMMulticontextPrivate, and returns true if they represent the same GObject.
func (recv *IMMulticontextPrivate) Equals(other *IMMulticontextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// IconFactoryClass is a wrapper around the C record GtkIconFactoryClass.
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

// Equals compares this IconFactoryClass with another IconFactoryClass, and returns true if they represent the same GObject.
func (recv *IconFactoryClass) Equals(other *IconFactoryClass) bool {
	return other.ToC() == recv.ToC()
}

// IconFactoryPrivate is a wrapper around the C record GtkIconFactoryPrivate.
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

// Equals compares this IconFactoryPrivate with another IconFactoryPrivate, and returns true if they represent the same GObject.
func (recv *IconFactoryPrivate) Equals(other *IconFactoryPrivate) bool {
	return other.ToC() == recv.ToC()
}

// IconInfoClass is a wrapper around the C record GtkIconInfoClass.
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

// Equals compares this IconInfoClass with another IconInfoClass, and returns true if they represent the same GObject.
func (recv *IconInfoClass) Equals(other *IconInfoClass) bool {
	return other.ToC() == recv.ToC()
}

// IconSet is a wrapper around the C record GtkIconSet.
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

// Equals compares this IconSet with another IconSet, and returns true if they represent the same GObject.
func (recv *IconSet) Equals(other *IconSet) bool {
	return other.ToC() == recv.ToC()
}

// IconSetNew is a wrapper around the C function gtk_icon_set_new.
func IconSetNew() *IconSet {
	retC := C.gtk_icon_set_new()
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconSetNewFromPixbuf is a wrapper around the C function gtk_icon_set_new_from_pixbuf.
func IconSetNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *IconSet {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_icon_set_new_from_pixbuf(c_pixbuf)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddSource is a wrapper around the C function gtk_icon_set_add_source.
func (recv *IconSet) AddSource(source *IconSource) {
	c_source := (*C.GtkIconSource)(C.NULL)
	if source != nil {
		c_source = (*C.GtkIconSource)(source.ToC())
	}

	C.gtk_icon_set_add_source((*C.GtkIconSet)(recv.native), c_source)

	return
}

// Copy is a wrapper around the C function gtk_icon_set_copy.
func (recv *IconSet) Copy() *IconSet {
	retC := C.gtk_icon_set_copy((*C.GtkIconSet)(recv.native))
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : output array param sizes

// Ref is a wrapper around the C function gtk_icon_set_ref.
func (recv *IconSet) Ref() *IconSet {
	retC := C.gtk_icon_set_ref((*C.GtkIconSet)(recv.native))
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RenderIcon is a wrapper around the C function gtk_icon_set_render_icon.
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

// Unref is a wrapper around the C function gtk_icon_set_unref.
func (recv *IconSet) Unref() {
	C.gtk_icon_set_unref((*C.GtkIconSet)(recv.native))

	return
}

// IconSource is a wrapper around the C record GtkIconSource.
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

// Equals compares this IconSource with another IconSource, and returns true if they represent the same GObject.
func (recv *IconSource) Equals(other *IconSource) bool {
	return other.ToC() == recv.ToC()
}

// IconSourceNew is a wrapper around the C function gtk_icon_source_new.
func IconSourceNew() *IconSource {
	retC := C.gtk_icon_source_new()
	retGo := IconSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_icon_source_copy.
func (recv *IconSource) Copy() *IconSource {
	retC := C.gtk_icon_source_copy((*C.GtkIconSource)(recv.native))
	retGo := IconSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_icon_source_free.
func (recv *IconSource) Free() {
	C.gtk_icon_source_free((*C.GtkIconSource)(recv.native))

	return
}

// GetDirection is a wrapper around the C function gtk_icon_source_get_direction.
func (recv *IconSource) GetDirection() TextDirection {
	retC := C.gtk_icon_source_get_direction((*C.GtkIconSource)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetDirectionWildcarded is a wrapper around the C function gtk_icon_source_get_direction_wildcarded.
func (recv *IconSource) GetDirectionWildcarded() bool {
	retC := C.gtk_icon_source_get_direction_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFilename is a wrapper around the C function gtk_icon_source_get_filename.
func (recv *IconSource) GetFilename() string {
	retC := C.gtk_icon_source_get_filename((*C.GtkIconSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIconName is a wrapper around the C function gtk_icon_source_get_icon_name.
func (recv *IconSource) GetIconName() string {
	retC := C.gtk_icon_source_get_icon_name((*C.GtkIconSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_icon_source_get_pixbuf.
func (recv *IconSource) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_icon_source_get_pixbuf((*C.GtkIconSource)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSize is a wrapper around the C function gtk_icon_source_get_size.
func (recv *IconSource) GetSize() IconSize {
	retC := C.gtk_icon_source_get_size((*C.GtkIconSource)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// GetSizeWildcarded is a wrapper around the C function gtk_icon_source_get_size_wildcarded.
func (recv *IconSource) GetSizeWildcarded() bool {
	retC := C.gtk_icon_source_get_size_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetState is a wrapper around the C function gtk_icon_source_get_state.
func (recv *IconSource) GetState() StateType {
	retC := C.gtk_icon_source_get_state((*C.GtkIconSource)(recv.native))
	retGo := (StateType)(retC)

	return retGo
}

// GetStateWildcarded is a wrapper around the C function gtk_icon_source_get_state_wildcarded.
func (recv *IconSource) GetStateWildcarded() bool {
	retC := C.gtk_icon_source_get_state_wildcarded((*C.GtkIconSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDirection is a wrapper around the C function gtk_icon_source_set_direction.
func (recv *IconSource) SetDirection(direction TextDirection) {
	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_icon_source_set_direction((*C.GtkIconSource)(recv.native), c_direction)

	return
}

// SetDirectionWildcarded is a wrapper around the C function gtk_icon_source_set_direction_wildcarded.
func (recv *IconSource) SetDirectionWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_direction_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

// SetFilename is a wrapper around the C function gtk_icon_source_set_filename.
func (recv *IconSource) SetFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_icon_source_set_filename((*C.GtkIconSource)(recv.native), c_filename)

	return
}

// SetIconName is a wrapper around the C function gtk_icon_source_set_icon_name.
func (recv *IconSource) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_icon_source_set_icon_name((*C.GtkIconSource)(recv.native), c_icon_name)

	return
}

// SetPixbuf is a wrapper around the C function gtk_icon_source_set_pixbuf.
func (recv *IconSource) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_icon_source_set_pixbuf((*C.GtkIconSource)(recv.native), c_pixbuf)

	return
}

// SetSize is a wrapper around the C function gtk_icon_source_set_size.
func (recv *IconSource) SetSize(size IconSize) {
	c_size := (C.GtkIconSize)(size)

	C.gtk_icon_source_set_size((*C.GtkIconSource)(recv.native), c_size)

	return
}

// SetSizeWildcarded is a wrapper around the C function gtk_icon_source_set_size_wildcarded.
func (recv *IconSource) SetSizeWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_size_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

// SetState is a wrapper around the C function gtk_icon_source_set_state.
func (recv *IconSource) SetState(state StateType) {
	c_state := (C.GtkStateType)(state)

	C.gtk_icon_source_set_state((*C.GtkIconSource)(recv.native), c_state)

	return
}

// SetStateWildcarded is a wrapper around the C function gtk_icon_source_set_state_wildcarded.
func (recv *IconSource) SetStateWildcarded(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_icon_source_set_state_wildcarded((*C.GtkIconSource)(recv.native), c_setting)

	return
}

// IconThemeClass is a wrapper around the C record GtkIconThemeClass.
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

// Equals compares this IconThemeClass with another IconThemeClass, and returns true if they represent the same GObject.
func (recv *IconThemeClass) Equals(other *IconThemeClass) bool {
	return other.ToC() == recv.ToC()
}

// IconThemePrivate is a wrapper around the C record GtkIconThemePrivate.
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

// Equals compares this IconThemePrivate with another IconThemePrivate, and returns true if they represent the same GObject.
func (recv *IconThemePrivate) Equals(other *IconThemePrivate) bool {
	return other.ToC() == recv.ToC()
}

// IconViewAccessibleClass is a wrapper around the C record GtkIconViewAccessibleClass.
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

// Equals compares this IconViewAccessibleClass with another IconViewAccessibleClass, and returns true if they represent the same GObject.
func (recv *IconViewAccessibleClass) Equals(other *IconViewAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// IconViewAccessiblePrivate is a wrapper around the C record GtkIconViewAccessiblePrivate.
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

// Equals compares this IconViewAccessiblePrivate with another IconViewAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *IconViewAccessiblePrivate) Equals(other *IconViewAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// IconViewClass is a wrapper around the C record GtkIconViewClass.
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

// Equals compares this IconViewClass with another IconViewClass, and returns true if they represent the same GObject.
func (recv *IconViewClass) Equals(other *IconViewClass) bool {
	return other.ToC() == recv.ToC()
}

// IconViewPrivate is a wrapper around the C record GtkIconViewPrivate.
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

// Equals compares this IconViewPrivate with another IconViewPrivate, and returns true if they represent the same GObject.
func (recv *IconViewPrivate) Equals(other *IconViewPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ImageAccessibleClass is a wrapper around the C record GtkImageAccessibleClass.
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

// Equals compares this ImageAccessibleClass with another ImageAccessibleClass, and returns true if they represent the same GObject.
func (recv *ImageAccessibleClass) Equals(other *ImageAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ImageAccessiblePrivate is a wrapper around the C record GtkImageAccessiblePrivate.
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

// Equals compares this ImageAccessiblePrivate with another ImageAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ImageAccessiblePrivate) Equals(other *ImageAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ImageCellAccessibleClass is a wrapper around the C record GtkImageCellAccessibleClass.
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

// Equals compares this ImageCellAccessibleClass with another ImageCellAccessibleClass, and returns true if they represent the same GObject.
func (recv *ImageCellAccessibleClass) Equals(other *ImageCellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ImageCellAccessiblePrivate is a wrapper around the C record GtkImageCellAccessiblePrivate.
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

// Equals compares this ImageCellAccessiblePrivate with another ImageCellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ImageCellAccessiblePrivate) Equals(other *ImageCellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ImageClass is a wrapper around the C record GtkImageClass.
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

// Equals compares this ImageClass with another ImageClass, and returns true if they represent the same GObject.
func (recv *ImageClass) Equals(other *ImageClass) bool {
	return other.ToC() == recv.ToC()
}

// ImageMenuItemClass is a wrapper around the C record GtkImageMenuItemClass.
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

// Equals compares this ImageMenuItemClass with another ImageMenuItemClass, and returns true if they represent the same GObject.
func (recv *ImageMenuItemClass) Equals(other *ImageMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// ImageMenuItemPrivate is a wrapper around the C record GtkImageMenuItemPrivate.
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

// Equals compares this ImageMenuItemPrivate with another ImageMenuItemPrivate, and returns true if they represent the same GObject.
func (recv *ImageMenuItemPrivate) Equals(other *ImageMenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ImagePrivate is a wrapper around the C record GtkImagePrivate.
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

// Equals compares this ImagePrivate with another ImagePrivate, and returns true if they represent the same GObject.
func (recv *ImagePrivate) Equals(other *ImagePrivate) bool {
	return other.ToC() == recv.ToC()
}

// InfoBarClass is a wrapper around the C record GtkInfoBarClass.
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

// Equals compares this InfoBarClass with another InfoBarClass, and returns true if they represent the same GObject.
func (recv *InfoBarClass) Equals(other *InfoBarClass) bool {
	return other.ToC() == recv.ToC()
}

// InfoBarPrivate is a wrapper around the C record GtkInfoBarPrivate.
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

// Equals compares this InfoBarPrivate with another InfoBarPrivate, and returns true if they represent the same GObject.
func (recv *InfoBarPrivate) Equals(other *InfoBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// InvisibleClass is a wrapper around the C record GtkInvisibleClass.
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

// Equals compares this InvisibleClass with another InvisibleClass, and returns true if they represent the same GObject.
func (recv *InvisibleClass) Equals(other *InvisibleClass) bool {
	return other.ToC() == recv.ToC()
}

// InvisiblePrivate is a wrapper around the C record GtkInvisiblePrivate.
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

// Equals compares this InvisiblePrivate with another InvisiblePrivate, and returns true if they represent the same GObject.
func (recv *InvisiblePrivate) Equals(other *InvisiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LabelAccessibleClass is a wrapper around the C record GtkLabelAccessibleClass.
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

// Equals compares this LabelAccessibleClass with another LabelAccessibleClass, and returns true if they represent the same GObject.
func (recv *LabelAccessibleClass) Equals(other *LabelAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// LabelAccessiblePrivate is a wrapper around the C record GtkLabelAccessiblePrivate.
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

// Equals compares this LabelAccessiblePrivate with another LabelAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *LabelAccessiblePrivate) Equals(other *LabelAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LabelClass is a wrapper around the C record GtkLabelClass.
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

// Equals compares this LabelClass with another LabelClass, and returns true if they represent the same GObject.
func (recv *LabelClass) Equals(other *LabelClass) bool {
	return other.ToC() == recv.ToC()
}

// LabelPrivate is a wrapper around the C record GtkLabelPrivate.
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

// Equals compares this LabelPrivate with another LabelPrivate, and returns true if they represent the same GObject.
func (recv *LabelPrivate) Equals(other *LabelPrivate) bool {
	return other.ToC() == recv.ToC()
}

// LabelSelectionInfo is a wrapper around the C record GtkLabelSelectionInfo.
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

// Equals compares this LabelSelectionInfo with another LabelSelectionInfo, and returns true if they represent the same GObject.
func (recv *LabelSelectionInfo) Equals(other *LabelSelectionInfo) bool {
	return other.ToC() == recv.ToC()
}

// LayoutClass is a wrapper around the C record GtkLayoutClass.
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

// Equals compares this LayoutClass with another LayoutClass, and returns true if they represent the same GObject.
func (recv *LayoutClass) Equals(other *LayoutClass) bool {
	return other.ToC() == recv.ToC()
}

// LayoutPrivate is a wrapper around the C record GtkLayoutPrivate.
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

// Equals compares this LayoutPrivate with another LayoutPrivate, and returns true if they represent the same GObject.
func (recv *LayoutPrivate) Equals(other *LayoutPrivate) bool {
	return other.ToC() == recv.ToC()
}

// LevelBarAccessibleClass is a wrapper around the C record GtkLevelBarAccessibleClass.
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

// Equals compares this LevelBarAccessibleClass with another LevelBarAccessibleClass, and returns true if they represent the same GObject.
func (recv *LevelBarAccessibleClass) Equals(other *LevelBarAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// LevelBarAccessiblePrivate is a wrapper around the C record GtkLevelBarAccessiblePrivate.
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

// Equals compares this LevelBarAccessiblePrivate with another LevelBarAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *LevelBarAccessiblePrivate) Equals(other *LevelBarAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LevelBarClass is a wrapper around the C record GtkLevelBarClass.
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

// Equals compares this LevelBarClass with another LevelBarClass, and returns true if they represent the same GObject.
func (recv *LevelBarClass) Equals(other *LevelBarClass) bool {
	return other.ToC() == recv.ToC()
}

// LevelBarPrivate is a wrapper around the C record GtkLevelBarPrivate.
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

// Equals compares this LevelBarPrivate with another LevelBarPrivate, and returns true if they represent the same GObject.
func (recv *LevelBarPrivate) Equals(other *LevelBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// LinkButtonAccessibleClass is a wrapper around the C record GtkLinkButtonAccessibleClass.
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

// Equals compares this LinkButtonAccessibleClass with another LinkButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *LinkButtonAccessibleClass) Equals(other *LinkButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// LinkButtonAccessiblePrivate is a wrapper around the C record GtkLinkButtonAccessiblePrivate.
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

// Equals compares this LinkButtonAccessiblePrivate with another LinkButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *LinkButtonAccessiblePrivate) Equals(other *LinkButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LinkButtonClass is a wrapper around the C record GtkLinkButtonClass.
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

// Equals compares this LinkButtonClass with another LinkButtonClass, and returns true if they represent the same GObject.
func (recv *LinkButtonClass) Equals(other *LinkButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// LinkButtonPrivate is a wrapper around the C record GtkLinkButtonPrivate.
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

// Equals compares this LinkButtonPrivate with another LinkButtonPrivate, and returns true if they represent the same GObject.
func (recv *LinkButtonPrivate) Equals(other *LinkButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ListBoxAccessibleClass is a wrapper around the C record GtkListBoxAccessibleClass.
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

// Equals compares this ListBoxAccessibleClass with another ListBoxAccessibleClass, and returns true if they represent the same GObject.
func (recv *ListBoxAccessibleClass) Equals(other *ListBoxAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ListBoxAccessiblePrivate is a wrapper around the C record GtkListBoxAccessiblePrivate.
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

// Equals compares this ListBoxAccessiblePrivate with another ListBoxAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ListBoxAccessiblePrivate) Equals(other *ListBoxAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ListBoxClass is a wrapper around the C record GtkListBoxClass.
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

// Equals compares this ListBoxClass with another ListBoxClass, and returns true if they represent the same GObject.
func (recv *ListBoxClass) Equals(other *ListBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// ListBoxRowAccessibleClass is a wrapper around the C record GtkListBoxRowAccessibleClass.
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

// Equals compares this ListBoxRowAccessibleClass with another ListBoxRowAccessibleClass, and returns true if they represent the same GObject.
func (recv *ListBoxRowAccessibleClass) Equals(other *ListBoxRowAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ListBoxRowClass is a wrapper around the C record GtkListBoxRowClass.
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

// Equals compares this ListBoxRowClass with another ListBoxRowClass, and returns true if they represent the same GObject.
func (recv *ListBoxRowClass) Equals(other *ListBoxRowClass) bool {
	return other.ToC() == recv.ToC()
}

// ListStoreClass is a wrapper around the C record GtkListStoreClass.
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

// Equals compares this ListStoreClass with another ListStoreClass, and returns true if they represent the same GObject.
func (recv *ListStoreClass) Equals(other *ListStoreClass) bool {
	return other.ToC() == recv.ToC()
}

// ListStorePrivate is a wrapper around the C record GtkListStorePrivate.
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

// Equals compares this ListStorePrivate with another ListStorePrivate, and returns true if they represent the same GObject.
func (recv *ListStorePrivate) Equals(other *ListStorePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LockButtonAccessibleClass is a wrapper around the C record GtkLockButtonAccessibleClass.
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

// Equals compares this LockButtonAccessibleClass with another LockButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *LockButtonAccessibleClass) Equals(other *LockButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// LockButtonAccessiblePrivate is a wrapper around the C record GtkLockButtonAccessiblePrivate.
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

// Equals compares this LockButtonAccessiblePrivate with another LockButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *LockButtonAccessiblePrivate) Equals(other *LockButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// LockButtonClass is a wrapper around the C record GtkLockButtonClass.
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

// Equals compares this LockButtonClass with another LockButtonClass, and returns true if they represent the same GObject.
func (recv *LockButtonClass) Equals(other *LockButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// LockButtonPrivate is a wrapper around the C record GtkLockButtonPrivate.
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

// Equals compares this LockButtonPrivate with another LockButtonPrivate, and returns true if they represent the same GObject.
func (recv *LockButtonPrivate) Equals(other *LockButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuAccessibleClass is a wrapper around the C record GtkMenuAccessibleClass.
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

// Equals compares this MenuAccessibleClass with another MenuAccessibleClass, and returns true if they represent the same GObject.
func (recv *MenuAccessibleClass) Equals(other *MenuAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuAccessiblePrivate is a wrapper around the C record GtkMenuAccessiblePrivate.
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

// Equals compares this MenuAccessiblePrivate with another MenuAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *MenuAccessiblePrivate) Equals(other *MenuAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuBarClass is a wrapper around the C record GtkMenuBarClass.
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

// Equals compares this MenuBarClass with another MenuBarClass, and returns true if they represent the same GObject.
func (recv *MenuBarClass) Equals(other *MenuBarClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuBarPrivate is a wrapper around the C record GtkMenuBarPrivate.
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

// Equals compares this MenuBarPrivate with another MenuBarPrivate, and returns true if they represent the same GObject.
func (recv *MenuBarPrivate) Equals(other *MenuBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuButtonAccessibleClass is a wrapper around the C record GtkMenuButtonAccessibleClass.
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

// Equals compares this MenuButtonAccessibleClass with another MenuButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *MenuButtonAccessibleClass) Equals(other *MenuButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuButtonAccessiblePrivate is a wrapper around the C record GtkMenuButtonAccessiblePrivate.
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

// Equals compares this MenuButtonAccessiblePrivate with another MenuButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *MenuButtonAccessiblePrivate) Equals(other *MenuButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuButtonClass is a wrapper around the C record GtkMenuButtonClass.
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

// Equals compares this MenuButtonClass with another MenuButtonClass, and returns true if they represent the same GObject.
func (recv *MenuButtonClass) Equals(other *MenuButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuButtonPrivate is a wrapper around the C record GtkMenuButtonPrivate.
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

// Equals compares this MenuButtonPrivate with another MenuButtonPrivate, and returns true if they represent the same GObject.
func (recv *MenuButtonPrivate) Equals(other *MenuButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuClass is a wrapper around the C record GtkMenuClass.
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

// Equals compares this MenuClass with another MenuClass, and returns true if they represent the same GObject.
func (recv *MenuClass) Equals(other *MenuClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuItemAccessibleClass is a wrapper around the C record GtkMenuItemAccessibleClass.
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

// Equals compares this MenuItemAccessibleClass with another MenuItemAccessibleClass, and returns true if they represent the same GObject.
func (recv *MenuItemAccessibleClass) Equals(other *MenuItemAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuItemAccessiblePrivate is a wrapper around the C record GtkMenuItemAccessiblePrivate.
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

// Equals compares this MenuItemAccessiblePrivate with another MenuItemAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *MenuItemAccessiblePrivate) Equals(other *MenuItemAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuItemClass is a wrapper around the C record GtkMenuItemClass.
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

// Equals compares this MenuItemClass with another MenuItemClass, and returns true if they represent the same GObject.
func (recv *MenuItemClass) Equals(other *MenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuItemPrivate is a wrapper around the C record GtkMenuItemPrivate.
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

// Equals compares this MenuItemPrivate with another MenuItemPrivate, and returns true if they represent the same GObject.
func (recv *MenuItemPrivate) Equals(other *MenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuPrivate is a wrapper around the C record GtkMenuPrivate.
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

// Equals compares this MenuPrivate with another MenuPrivate, and returns true if they represent the same GObject.
func (recv *MenuPrivate) Equals(other *MenuPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuShellAccessibleClass is a wrapper around the C record GtkMenuShellAccessibleClass.
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

// Equals compares this MenuShellAccessibleClass with another MenuShellAccessibleClass, and returns true if they represent the same GObject.
func (recv *MenuShellAccessibleClass) Equals(other *MenuShellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuShellAccessiblePrivate is a wrapper around the C record GtkMenuShellAccessiblePrivate.
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

// Equals compares this MenuShellAccessiblePrivate with another MenuShellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *MenuShellAccessiblePrivate) Equals(other *MenuShellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuShellClass is a wrapper around the C record GtkMenuShellClass.
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

// Equals compares this MenuShellClass with another MenuShellClass, and returns true if they represent the same GObject.
func (recv *MenuShellClass) Equals(other *MenuShellClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuShellPrivate is a wrapper around the C record GtkMenuShellPrivate.
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

// Equals compares this MenuShellPrivate with another MenuShellPrivate, and returns true if they represent the same GObject.
func (recv *MenuShellPrivate) Equals(other *MenuShellPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuToolButtonClass is a wrapper around the C record GtkMenuToolButtonClass.
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

// Equals compares this MenuToolButtonClass with another MenuToolButtonClass, and returns true if they represent the same GObject.
func (recv *MenuToolButtonClass) Equals(other *MenuToolButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuToolButtonPrivate is a wrapper around the C record GtkMenuToolButtonPrivate.
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

// Equals compares this MenuToolButtonPrivate with another MenuToolButtonPrivate, and returns true if they represent the same GObject.
func (recv *MenuToolButtonPrivate) Equals(other *MenuToolButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MessageDialogClass is a wrapper around the C record GtkMessageDialogClass.
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

// Equals compares this MessageDialogClass with another MessageDialogClass, and returns true if they represent the same GObject.
func (recv *MessageDialogClass) Equals(other *MessageDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// MessageDialogPrivate is a wrapper around the C record GtkMessageDialogPrivate.
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

// Equals compares this MessageDialogPrivate with another MessageDialogPrivate, and returns true if they represent the same GObject.
func (recv *MessageDialogPrivate) Equals(other *MessageDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MiscClass is a wrapper around the C record GtkMiscClass.
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

// Equals compares this MiscClass with another MiscClass, and returns true if they represent the same GObject.
func (recv *MiscClass) Equals(other *MiscClass) bool {
	return other.ToC() == recv.ToC()
}

// MiscPrivate is a wrapper around the C record GtkMiscPrivate.
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

// Equals compares this MiscPrivate with another MiscPrivate, and returns true if they represent the same GObject.
func (recv *MiscPrivate) Equals(other *MiscPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MountOperationClass is a wrapper around the C record GtkMountOperationClass.
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

// Equals compares this MountOperationClass with another MountOperationClass, and returns true if they represent the same GObject.
func (recv *MountOperationClass) Equals(other *MountOperationClass) bool {
	return other.ToC() == recv.ToC()
}

// MountOperationPrivate is a wrapper around the C record GtkMountOperationPrivate.
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

// Equals compares this MountOperationPrivate with another MountOperationPrivate, and returns true if they represent the same GObject.
func (recv *MountOperationPrivate) Equals(other *MountOperationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NotebookAccessibleClass is a wrapper around the C record GtkNotebookAccessibleClass.
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

// Equals compares this NotebookAccessibleClass with another NotebookAccessibleClass, and returns true if they represent the same GObject.
func (recv *NotebookAccessibleClass) Equals(other *NotebookAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// NotebookAccessiblePrivate is a wrapper around the C record GtkNotebookAccessiblePrivate.
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

// Equals compares this NotebookAccessiblePrivate with another NotebookAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *NotebookAccessiblePrivate) Equals(other *NotebookAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// NotebookClass is a wrapper around the C record GtkNotebookClass.
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

// Equals compares this NotebookClass with another NotebookClass, and returns true if they represent the same GObject.
func (recv *NotebookClass) Equals(other *NotebookClass) bool {
	return other.ToC() == recv.ToC()
}

// NotebookPageAccessibleClass is a wrapper around the C record GtkNotebookPageAccessibleClass.
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

// Equals compares this NotebookPageAccessibleClass with another NotebookPageAccessibleClass, and returns true if they represent the same GObject.
func (recv *NotebookPageAccessibleClass) Equals(other *NotebookPageAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// NotebookPageAccessiblePrivate is a wrapper around the C record GtkNotebookPageAccessiblePrivate.
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

// Equals compares this NotebookPageAccessiblePrivate with another NotebookPageAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *NotebookPageAccessiblePrivate) Equals(other *NotebookPageAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// NotebookPrivate is a wrapper around the C record GtkNotebookPrivate.
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

// Equals compares this NotebookPrivate with another NotebookPrivate, and returns true if they represent the same GObject.
func (recv *NotebookPrivate) Equals(other *NotebookPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NumerableIconClass is a wrapper around the C record GtkNumerableIconClass.
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

// Equals compares this NumerableIconClass with another NumerableIconClass, and returns true if they represent the same GObject.
func (recv *NumerableIconClass) Equals(other *NumerableIconClass) bool {
	return other.ToC() == recv.ToC()
}

// NumerableIconPrivate is a wrapper around the C record GtkNumerableIconPrivate.
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

// Equals compares this NumerableIconPrivate with another NumerableIconPrivate, and returns true if they represent the same GObject.
func (recv *NumerableIconPrivate) Equals(other *NumerableIconPrivate) bool {
	return other.ToC() == recv.ToC()
}

// OffscreenWindowClass is a wrapper around the C record GtkOffscreenWindowClass.
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

// Equals compares this OffscreenWindowClass with another OffscreenWindowClass, and returns true if they represent the same GObject.
func (recv *OffscreenWindowClass) Equals(other *OffscreenWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// OrientableIface is a wrapper around the C record GtkOrientableIface.
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

// Equals compares this OrientableIface with another OrientableIface, and returns true if they represent the same GObject.
func (recv *OrientableIface) Equals(other *OrientableIface) bool {
	return other.ToC() == recv.ToC()
}

// OverlayClass is a wrapper around the C record GtkOverlayClass.
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

// Equals compares this OverlayClass with another OverlayClass, and returns true if they represent the same GObject.
func (recv *OverlayClass) Equals(other *OverlayClass) bool {
	return other.ToC() == recv.ToC()
}

// OverlayPrivate is a wrapper around the C record GtkOverlayPrivate.
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

// Equals compares this OverlayPrivate with another OverlayPrivate, and returns true if they represent the same GObject.
func (recv *OverlayPrivate) Equals(other *OverlayPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PageRange is a wrapper around the C record GtkPageRange.
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

// Equals compares this PageRange with another PageRange, and returns true if they represent the same GObject.
func (recv *PageRange) Equals(other *PageRange) bool {
	return other.ToC() == recv.ToC()
}

// PanedAccessibleClass is a wrapper around the C record GtkPanedAccessibleClass.
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

// Equals compares this PanedAccessibleClass with another PanedAccessibleClass, and returns true if they represent the same GObject.
func (recv *PanedAccessibleClass) Equals(other *PanedAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// PanedAccessiblePrivate is a wrapper around the C record GtkPanedAccessiblePrivate.
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

// Equals compares this PanedAccessiblePrivate with another PanedAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *PanedAccessiblePrivate) Equals(other *PanedAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// PanedClass is a wrapper around the C record GtkPanedClass.
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

// Equals compares this PanedClass with another PanedClass, and returns true if they represent the same GObject.
func (recv *PanedClass) Equals(other *PanedClass) bool {
	return other.ToC() == recv.ToC()
}

// PanedPrivate is a wrapper around the C record GtkPanedPrivate.
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

// Equals compares this PanedPrivate with another PanedPrivate, and returns true if they represent the same GObject.
func (recv *PanedPrivate) Equals(other *PanedPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PaperSize is a wrapper around the C record GtkPaperSize.
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

// Equals compares this PaperSize with another PaperSize, and returns true if they represent the same GObject.
func (recv *PaperSize) Equals(other *PaperSize) bool {
	return other.ToC() == recv.ToC()
}

// IsCustom is a wrapper around the C function gtk_paper_size_is_custom.
func (recv *PaperSize) IsCustom() bool {
	retC := C.gtk_paper_size_is_custom((*C.GtkPaperSize)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsIpp is a wrapper around the C function gtk_paper_size_is_ipp.
func (recv *PaperSize) IsIpp() bool {
	retC := C.gtk_paper_size_is_ipp((*C.GtkPaperSize)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PlacesSidebarClass is a wrapper around the C record GtkPlacesSidebarClass.
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

// Equals compares this PlacesSidebarClass with another PlacesSidebarClass, and returns true if they represent the same GObject.
func (recv *PlacesSidebarClass) Equals(other *PlacesSidebarClass) bool {
	return other.ToC() == recv.ToC()
}

// PlugClass is a wrapper around the C record GtkPlugClass.
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

// Equals compares this PlugClass with another PlugClass, and returns true if they represent the same GObject.
func (recv *PlugClass) Equals(other *PlugClass) bool {
	return other.ToC() == recv.ToC()
}

// PlugPrivate is a wrapper around the C record GtkPlugPrivate.
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

// Equals compares this PlugPrivate with another PlugPrivate, and returns true if they represent the same GObject.
func (recv *PlugPrivate) Equals(other *PlugPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PopoverAccessibleClass is a wrapper around the C record GtkPopoverAccessibleClass.
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

// Equals compares this PopoverAccessibleClass with another PopoverAccessibleClass, and returns true if they represent the same GObject.
func (recv *PopoverAccessibleClass) Equals(other *PopoverAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// PopoverClass is a wrapper around the C record GtkPopoverClass.
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

// Equals compares this PopoverClass with another PopoverClass, and returns true if they represent the same GObject.
func (recv *PopoverClass) Equals(other *PopoverClass) bool {
	return other.ToC() == recv.ToC()
}

// PopoverMenuClass is a wrapper around the C record GtkPopoverMenuClass.
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

// Equals compares this PopoverMenuClass with another PopoverMenuClass, and returns true if they represent the same GObject.
func (recv *PopoverMenuClass) Equals(other *PopoverMenuClass) bool {
	return other.ToC() == recv.ToC()
}

// PopoverPrivate is a wrapper around the C record GtkPopoverPrivate.
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

// Equals compares this PopoverPrivate with another PopoverPrivate, and returns true if they represent the same GObject.
func (recv *PopoverPrivate) Equals(other *PopoverPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PrintOperationClass is a wrapper around the C record GtkPrintOperationClass.
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

// Equals compares this PrintOperationClass with another PrintOperationClass, and returns true if they represent the same GObject.
func (recv *PrintOperationClass) Equals(other *PrintOperationClass) bool {
	return other.ToC() == recv.ToC()
}

// PrintOperationPreviewIface is a wrapper around the C record GtkPrintOperationPreviewIface.
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

// Equals compares this PrintOperationPreviewIface with another PrintOperationPreviewIface, and returns true if they represent the same GObject.
func (recv *PrintOperationPreviewIface) Equals(other *PrintOperationPreviewIface) bool {
	return other.ToC() == recv.ToC()
}

// PrintOperationPrivate is a wrapper around the C record GtkPrintOperationPrivate.
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

// Equals compares this PrintOperationPrivate with another PrintOperationPrivate, and returns true if they represent the same GObject.
func (recv *PrintOperationPrivate) Equals(other *PrintOperationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ProgressBarAccessibleClass is a wrapper around the C record GtkProgressBarAccessibleClass.
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

// Equals compares this ProgressBarAccessibleClass with another ProgressBarAccessibleClass, and returns true if they represent the same GObject.
func (recv *ProgressBarAccessibleClass) Equals(other *ProgressBarAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ProgressBarAccessiblePrivate is a wrapper around the C record GtkProgressBarAccessiblePrivate.
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

// Equals compares this ProgressBarAccessiblePrivate with another ProgressBarAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ProgressBarAccessiblePrivate) Equals(other *ProgressBarAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ProgressBarClass is a wrapper around the C record GtkProgressBarClass.
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

// Equals compares this ProgressBarClass with another ProgressBarClass, and returns true if they represent the same GObject.
func (recv *ProgressBarClass) Equals(other *ProgressBarClass) bool {
	return other.ToC() == recv.ToC()
}

// ProgressBarPrivate is a wrapper around the C record GtkProgressBarPrivate.
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

// Equals compares this ProgressBarPrivate with another ProgressBarPrivate, and returns true if they represent the same GObject.
func (recv *ProgressBarPrivate) Equals(other *ProgressBarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioActionClass is a wrapper around the C record GtkRadioActionClass.
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

// Equals compares this RadioActionClass with another RadioActionClass, and returns true if they represent the same GObject.
func (recv *RadioActionClass) Equals(other *RadioActionClass) bool {
	return other.ToC() == recv.ToC()
}

// RadioActionEntry is a wrapper around the C record GtkRadioActionEntry.
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

// Equals compares this RadioActionEntry with another RadioActionEntry, and returns true if they represent the same GObject.
func (recv *RadioActionEntry) Equals(other *RadioActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// RadioActionPrivate is a wrapper around the C record GtkRadioActionPrivate.
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

// Equals compares this RadioActionPrivate with another RadioActionPrivate, and returns true if they represent the same GObject.
func (recv *RadioActionPrivate) Equals(other *RadioActionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioButtonAccessibleClass is a wrapper around the C record GtkRadioButtonAccessibleClass.
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

// Equals compares this RadioButtonAccessibleClass with another RadioButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *RadioButtonAccessibleClass) Equals(other *RadioButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// RadioButtonAccessiblePrivate is a wrapper around the C record GtkRadioButtonAccessiblePrivate.
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

// Equals compares this RadioButtonAccessiblePrivate with another RadioButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *RadioButtonAccessiblePrivate) Equals(other *RadioButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioButtonClass is a wrapper around the C record GtkRadioButtonClass.
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

// Equals compares this RadioButtonClass with another RadioButtonClass, and returns true if they represent the same GObject.
func (recv *RadioButtonClass) Equals(other *RadioButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// RadioButtonPrivate is a wrapper around the C record GtkRadioButtonPrivate.
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

// Equals compares this RadioButtonPrivate with another RadioButtonPrivate, and returns true if they represent the same GObject.
func (recv *RadioButtonPrivate) Equals(other *RadioButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioMenuItemAccessibleClass is a wrapper around the C record GtkRadioMenuItemAccessibleClass.
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

// Equals compares this RadioMenuItemAccessibleClass with another RadioMenuItemAccessibleClass, and returns true if they represent the same GObject.
func (recv *RadioMenuItemAccessibleClass) Equals(other *RadioMenuItemAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// RadioMenuItemAccessiblePrivate is a wrapper around the C record GtkRadioMenuItemAccessiblePrivate.
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

// Equals compares this RadioMenuItemAccessiblePrivate with another RadioMenuItemAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *RadioMenuItemAccessiblePrivate) Equals(other *RadioMenuItemAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioMenuItemClass is a wrapper around the C record GtkRadioMenuItemClass.
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

// Equals compares this RadioMenuItemClass with another RadioMenuItemClass, and returns true if they represent the same GObject.
func (recv *RadioMenuItemClass) Equals(other *RadioMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// RadioMenuItemPrivate is a wrapper around the C record GtkRadioMenuItemPrivate.
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

// Equals compares this RadioMenuItemPrivate with another RadioMenuItemPrivate, and returns true if they represent the same GObject.
func (recv *RadioMenuItemPrivate) Equals(other *RadioMenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RadioToolButtonClass is a wrapper around the C record GtkRadioToolButtonClass.
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

// Equals compares this RadioToolButtonClass with another RadioToolButtonClass, and returns true if they represent the same GObject.
func (recv *RadioToolButtonClass) Equals(other *RadioToolButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// RangeAccessibleClass is a wrapper around the C record GtkRangeAccessibleClass.
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

// Equals compares this RangeAccessibleClass with another RangeAccessibleClass, and returns true if they represent the same GObject.
func (recv *RangeAccessibleClass) Equals(other *RangeAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// RangeAccessiblePrivate is a wrapper around the C record GtkRangeAccessiblePrivate.
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

// Equals compares this RangeAccessiblePrivate with another RangeAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *RangeAccessiblePrivate) Equals(other *RangeAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RangeClass is a wrapper around the C record GtkRangeClass.
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

// Equals compares this RangeClass with another RangeClass, and returns true if they represent the same GObject.
func (recv *RangeClass) Equals(other *RangeClass) bool {
	return other.ToC() == recv.ToC()
}

// RangePrivate is a wrapper around the C record GtkRangePrivate.
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

// Equals compares this RangePrivate with another RangePrivate, and returns true if they represent the same GObject.
func (recv *RangePrivate) Equals(other *RangePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RcContext is a wrapper around the C record GtkRcContext.
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

// Equals compares this RcContext with another RcContext, and returns true if they represent the same GObject.
func (recv *RcContext) Equals(other *RcContext) bool {
	return other.ToC() == recv.ToC()
}

// RcProperty is a wrapper around the C record GtkRcProperty.
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

// Equals compares this RcProperty with another RcProperty, and returns true if they represent the same GObject.
func (recv *RcProperty) Equals(other *RcProperty) bool {
	return other.ToC() == recv.ToC()
}

// gtk_rc_property_parse_border : unsupported parameter pspec : Blacklisted record : GParamSpec
// gtk_rc_property_parse_color : unsupported parameter pspec : Blacklisted record : GParamSpec
// gtk_rc_property_parse_enum : unsupported parameter pspec : Blacklisted record : GParamSpec
// gtk_rc_property_parse_flags : unsupported parameter pspec : Blacklisted record : GParamSpec
// gtk_rc_property_parse_requisition : unsupported parameter pspec : Blacklisted record : GParamSpec
// RcStyleClass is a wrapper around the C record GtkRcStyleClass.
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

// Equals compares this RcStyleClass with another RcStyleClass, and returns true if they represent the same GObject.
func (recv *RcStyleClass) Equals(other *RcStyleClass) bool {
	return other.ToC() == recv.ToC()
}

// RecentActionClass is a wrapper around the C record GtkRecentActionClass.
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

// Equals compares this RecentActionClass with another RecentActionClass, and returns true if they represent the same GObject.
func (recv *RecentActionClass) Equals(other *RecentActionClass) bool {
	return other.ToC() == recv.ToC()
}

// RecentActionPrivate is a wrapper around the C record GtkRecentActionPrivate.
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

// Equals compares this RecentActionPrivate with another RecentActionPrivate, and returns true if they represent the same GObject.
func (recv *RecentActionPrivate) Equals(other *RecentActionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserDialogClass is a wrapper around the C record GtkRecentChooserDialogClass.
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

// Equals compares this RecentChooserDialogClass with another RecentChooserDialogClass, and returns true if they represent the same GObject.
func (recv *RecentChooserDialogClass) Equals(other *RecentChooserDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserDialogPrivate is a wrapper around the C record GtkRecentChooserDialogPrivate.
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

// Equals compares this RecentChooserDialogPrivate with another RecentChooserDialogPrivate, and returns true if they represent the same GObject.
func (recv *RecentChooserDialogPrivate) Equals(other *RecentChooserDialogPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserIface is a wrapper around the C record GtkRecentChooserIface.
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

// Equals compares this RecentChooserIface with another RecentChooserIface, and returns true if they represent the same GObject.
func (recv *RecentChooserIface) Equals(other *RecentChooserIface) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserMenuClass is a wrapper around the C record GtkRecentChooserMenuClass.
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

// Equals compares this RecentChooserMenuClass with another RecentChooserMenuClass, and returns true if they represent the same GObject.
func (recv *RecentChooserMenuClass) Equals(other *RecentChooserMenuClass) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserMenuPrivate is a wrapper around the C record GtkRecentChooserMenuPrivate.
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

// Equals compares this RecentChooserMenuPrivate with another RecentChooserMenuPrivate, and returns true if they represent the same GObject.
func (recv *RecentChooserMenuPrivate) Equals(other *RecentChooserMenuPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserWidgetClass is a wrapper around the C record GtkRecentChooserWidgetClass.
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

// Equals compares this RecentChooserWidgetClass with another RecentChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *RecentChooserWidgetClass) Equals(other *RecentChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// RecentChooserWidgetPrivate is a wrapper around the C record GtkRecentChooserWidgetPrivate.
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

// Equals compares this RecentChooserWidgetPrivate with another RecentChooserWidgetPrivate, and returns true if they represent the same GObject.
func (recv *RecentChooserWidgetPrivate) Equals(other *RecentChooserWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RecentData is a wrapper around the C record GtkRecentData.
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

// Equals compares this RecentData with another RecentData, and returns true if they represent the same GObject.
func (recv *RecentData) Equals(other *RecentData) bool {
	return other.ToC() == recv.ToC()
}

// RecentFilterInfo is a wrapper around the C record GtkRecentFilterInfo.
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

// Equals compares this RecentFilterInfo with another RecentFilterInfo, and returns true if they represent the same GObject.
func (recv *RecentFilterInfo) Equals(other *RecentFilterInfo) bool {
	return other.ToC() == recv.ToC()
}

// RecentManagerPrivate is a wrapper around the C record GtkRecentManagerPrivate.
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

// Equals compares this RecentManagerPrivate with another RecentManagerPrivate, and returns true if they represent the same GObject.
func (recv *RecentManagerPrivate) Equals(other *RecentManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RendererCellAccessibleClass is a wrapper around the C record GtkRendererCellAccessibleClass.
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

// Equals compares this RendererCellAccessibleClass with another RendererCellAccessibleClass, and returns true if they represent the same GObject.
func (recv *RendererCellAccessibleClass) Equals(other *RendererCellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// RendererCellAccessiblePrivate is a wrapper around the C record GtkRendererCellAccessiblePrivate.
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

// Equals compares this RendererCellAccessiblePrivate with another RendererCellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *RendererCellAccessiblePrivate) Equals(other *RendererCellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RequestedSize is a wrapper around the C record GtkRequestedSize.
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

// Equals compares this RequestedSize with another RequestedSize, and returns true if they represent the same GObject.
func (recv *RequestedSize) Equals(other *RequestedSize) bool {
	return other.ToC() == recv.ToC()
}

// Requisition is a wrapper around the C record GtkRequisition.
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

// Equals compares this Requisition with another Requisition, and returns true if they represent the same GObject.
func (recv *Requisition) Equals(other *Requisition) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gtk_requisition_copy.
func (recv *Requisition) Copy() *Requisition {
	retC := C.gtk_requisition_copy((*C.GtkRequisition)(recv.native))
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_requisition_free.
func (recv *Requisition) Free() {
	C.gtk_requisition_free((*C.GtkRequisition)(recv.native))

	return
}

// RevealerClass is a wrapper around the C record GtkRevealerClass.
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

// Equals compares this RevealerClass with another RevealerClass, and returns true if they represent the same GObject.
func (recv *RevealerClass) Equals(other *RevealerClass) bool {
	return other.ToC() == recv.ToC()
}

// ScaleAccessibleClass is a wrapper around the C record GtkScaleAccessibleClass.
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

// Equals compares this ScaleAccessibleClass with another ScaleAccessibleClass, and returns true if they represent the same GObject.
func (recv *ScaleAccessibleClass) Equals(other *ScaleAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ScaleAccessiblePrivate is a wrapper around the C record GtkScaleAccessiblePrivate.
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

// Equals compares this ScaleAccessiblePrivate with another ScaleAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ScaleAccessiblePrivate) Equals(other *ScaleAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScaleButtonAccessibleClass is a wrapper around the C record GtkScaleButtonAccessibleClass.
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

// Equals compares this ScaleButtonAccessibleClass with another ScaleButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *ScaleButtonAccessibleClass) Equals(other *ScaleButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ScaleButtonAccessiblePrivate is a wrapper around the C record GtkScaleButtonAccessiblePrivate.
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

// Equals compares this ScaleButtonAccessiblePrivate with another ScaleButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ScaleButtonAccessiblePrivate) Equals(other *ScaleButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScaleButtonClass is a wrapper around the C record GtkScaleButtonClass.
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

// Equals compares this ScaleButtonClass with another ScaleButtonClass, and returns true if they represent the same GObject.
func (recv *ScaleButtonClass) Equals(other *ScaleButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ScaleButtonPrivate is a wrapper around the C record GtkScaleButtonPrivate.
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

// Equals compares this ScaleButtonPrivate with another ScaleButtonPrivate, and returns true if they represent the same GObject.
func (recv *ScaleButtonPrivate) Equals(other *ScaleButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScaleClass is a wrapper around the C record GtkScaleClass.
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

// Equals compares this ScaleClass with another ScaleClass, and returns true if they represent the same GObject.
func (recv *ScaleClass) Equals(other *ScaleClass) bool {
	return other.ToC() == recv.ToC()
}

// ScalePrivate is a wrapper around the C record GtkScalePrivate.
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

// Equals compares this ScalePrivate with another ScalePrivate, and returns true if they represent the same GObject.
func (recv *ScalePrivate) Equals(other *ScalePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScrollableInterface is a wrapper around the C record GtkScrollableInterface.
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

// Equals compares this ScrollableInterface with another ScrollableInterface, and returns true if they represent the same GObject.
func (recv *ScrollableInterface) Equals(other *ScrollableInterface) bool {
	return other.ToC() == recv.ToC()
}

// ScrollbarClass is a wrapper around the C record GtkScrollbarClass.
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

// Equals compares this ScrollbarClass with another ScrollbarClass, and returns true if they represent the same GObject.
func (recv *ScrollbarClass) Equals(other *ScrollbarClass) bool {
	return other.ToC() == recv.ToC()
}

// ScrolledWindowAccessibleClass is a wrapper around the C record GtkScrolledWindowAccessibleClass.
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

// Equals compares this ScrolledWindowAccessibleClass with another ScrolledWindowAccessibleClass, and returns true if they represent the same GObject.
func (recv *ScrolledWindowAccessibleClass) Equals(other *ScrolledWindowAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ScrolledWindowAccessiblePrivate is a wrapper around the C record GtkScrolledWindowAccessiblePrivate.
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

// Equals compares this ScrolledWindowAccessiblePrivate with another ScrolledWindowAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ScrolledWindowAccessiblePrivate) Equals(other *ScrolledWindowAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScrolledWindowClass is a wrapper around the C record GtkScrolledWindowClass.
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

// Equals compares this ScrolledWindowClass with another ScrolledWindowClass, and returns true if they represent the same GObject.
func (recv *ScrolledWindowClass) Equals(other *ScrolledWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// ScrolledWindowPrivate is a wrapper around the C record GtkScrolledWindowPrivate.
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

// Equals compares this ScrolledWindowPrivate with another ScrolledWindowPrivate, and returns true if they represent the same GObject.
func (recv *ScrolledWindowPrivate) Equals(other *ScrolledWindowPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SearchBarClass is a wrapper around the C record GtkSearchBarClass.
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

// Equals compares this SearchBarClass with another SearchBarClass, and returns true if they represent the same GObject.
func (recv *SearchBarClass) Equals(other *SearchBarClass) bool {
	return other.ToC() == recv.ToC()
}

// SearchEntryClass is a wrapper around the C record GtkSearchEntryClass.
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

// Equals compares this SearchEntryClass with another SearchEntryClass, and returns true if they represent the same GObject.
func (recv *SearchEntryClass) Equals(other *SearchEntryClass) bool {
	return other.ToC() == recv.ToC()
}

// SelectionData is a wrapper around the C record GtkSelectionData.
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

// Equals compares this SelectionData with another SelectionData, and returns true if they represent the same GObject.
func (recv *SelectionData) Equals(other *SelectionData) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gtk_selection_data_copy.
func (recv *SelectionData) Copy() *SelectionData {
	retC := C.gtk_selection_data_copy((*C.GtkSelectionData)(recv.native))
	retGo := SelectionDataNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_selection_data_free.
func (recv *SelectionData) Free() {
	C.gtk_selection_data_free((*C.GtkSelectionData)(recv.native))

	return
}

// Unsupported : gtk_selection_data_get_targets : unsupported parameter targets : output array param targets

// Blacklisted : gtk_selection_data_get_text

// Unsupported : gtk_selection_data_set : unsupported parameter type : Blacklisted record : GdkAtom

// SetText is a wrapper around the C function gtk_selection_data_set_text.
func (recv *SelectionData) SetText(str string, len int32) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gint)(len)

	retC := C.gtk_selection_data_set_text((*C.GtkSelectionData)(recv.native), c_str, c_len)
	retGo := retC == C.TRUE

	return retGo
}

// TargetsIncludeText is a wrapper around the C function gtk_selection_data_targets_include_text.
func (recv *SelectionData) TargetsIncludeText() bool {
	retC := C.gtk_selection_data_targets_include_text((*C.GtkSelectionData)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SeparatorClass is a wrapper around the C record GtkSeparatorClass.
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

// Equals compares this SeparatorClass with another SeparatorClass, and returns true if they represent the same GObject.
func (recv *SeparatorClass) Equals(other *SeparatorClass) bool {
	return other.ToC() == recv.ToC()
}

// SeparatorMenuItemClass is a wrapper around the C record GtkSeparatorMenuItemClass.
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

// Equals compares this SeparatorMenuItemClass with another SeparatorMenuItemClass, and returns true if they represent the same GObject.
func (recv *SeparatorMenuItemClass) Equals(other *SeparatorMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// SeparatorPrivate is a wrapper around the C record GtkSeparatorPrivate.
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

// Equals compares this SeparatorPrivate with another SeparatorPrivate, and returns true if they represent the same GObject.
func (recv *SeparatorPrivate) Equals(other *SeparatorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SeparatorToolItemClass is a wrapper around the C record GtkSeparatorToolItemClass.
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

// Equals compares this SeparatorToolItemClass with another SeparatorToolItemClass, and returns true if they represent the same GObject.
func (recv *SeparatorToolItemClass) Equals(other *SeparatorToolItemClass) bool {
	return other.ToC() == recv.ToC()
}

// SeparatorToolItemPrivate is a wrapper around the C record GtkSeparatorToolItemPrivate.
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

// Equals compares this SeparatorToolItemPrivate with another SeparatorToolItemPrivate, and returns true if they represent the same GObject.
func (recv *SeparatorToolItemPrivate) Equals(other *SeparatorToolItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SettingsClass is a wrapper around the C record GtkSettingsClass.
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

// Equals compares this SettingsClass with another SettingsClass, and returns true if they represent the same GObject.
func (recv *SettingsClass) Equals(other *SettingsClass) bool {
	return other.ToC() == recv.ToC()
}

// SettingsPrivate is a wrapper around the C record GtkSettingsPrivate.
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

// Equals compares this SettingsPrivate with another SettingsPrivate, and returns true if they represent the same GObject.
func (recv *SettingsPrivate) Equals(other *SettingsPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SettingsValue is a wrapper around the C record GtkSettingsValue.
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

// Equals compares this SettingsValue with another SettingsValue, and returns true if they represent the same GObject.
func (recv *SettingsValue) Equals(other *SettingsValue) bool {
	return other.ToC() == recv.ToC()
}

// SizeGroupClass is a wrapper around the C record GtkSizeGroupClass.
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

// Equals compares this SizeGroupClass with another SizeGroupClass, and returns true if they represent the same GObject.
func (recv *SizeGroupClass) Equals(other *SizeGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// SizeGroupPrivate is a wrapper around the C record GtkSizeGroupPrivate.
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

// Equals compares this SizeGroupPrivate with another SizeGroupPrivate, and returns true if they represent the same GObject.
func (recv *SizeGroupPrivate) Equals(other *SizeGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketClass is a wrapper around the C record GtkSocketClass.
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

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same GObject.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketPrivate is a wrapper around the C record GtkSocketPrivate.
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

// Equals compares this SocketPrivate with another SocketPrivate, and returns true if they represent the same GObject.
func (recv *SocketPrivate) Equals(other *SocketPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SpinButtonAccessibleClass is a wrapper around the C record GtkSpinButtonAccessibleClass.
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

// Equals compares this SpinButtonAccessibleClass with another SpinButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *SpinButtonAccessibleClass) Equals(other *SpinButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// SpinButtonAccessiblePrivate is a wrapper around the C record GtkSpinButtonAccessiblePrivate.
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

// Equals compares this SpinButtonAccessiblePrivate with another SpinButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *SpinButtonAccessiblePrivate) Equals(other *SpinButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// SpinButtonClass is a wrapper around the C record GtkSpinButtonClass.
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

// Equals compares this SpinButtonClass with another SpinButtonClass, and returns true if they represent the same GObject.
func (recv *SpinButtonClass) Equals(other *SpinButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// SpinButtonPrivate is a wrapper around the C record GtkSpinButtonPrivate.
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

// Equals compares this SpinButtonPrivate with another SpinButtonPrivate, and returns true if they represent the same GObject.
func (recv *SpinButtonPrivate) Equals(other *SpinButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SpinnerAccessibleClass is a wrapper around the C record GtkSpinnerAccessibleClass.
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

// Equals compares this SpinnerAccessibleClass with another SpinnerAccessibleClass, and returns true if they represent the same GObject.
func (recv *SpinnerAccessibleClass) Equals(other *SpinnerAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// SpinnerAccessiblePrivate is a wrapper around the C record GtkSpinnerAccessiblePrivate.
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

// Equals compares this SpinnerAccessiblePrivate with another SpinnerAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *SpinnerAccessiblePrivate) Equals(other *SpinnerAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// SpinnerClass is a wrapper around the C record GtkSpinnerClass.
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

// Equals compares this SpinnerClass with another SpinnerClass, and returns true if they represent the same GObject.
func (recv *SpinnerClass) Equals(other *SpinnerClass) bool {
	return other.ToC() == recv.ToC()
}

// SpinnerPrivate is a wrapper around the C record GtkSpinnerPrivate.
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

// Equals compares this SpinnerPrivate with another SpinnerPrivate, and returns true if they represent the same GObject.
func (recv *SpinnerPrivate) Equals(other *SpinnerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GtkStackAccessibleClass

// StackClass is a wrapper around the C record GtkStackClass.
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

// Equals compares this StackClass with another StackClass, and returns true if they represent the same GObject.
func (recv *StackClass) Equals(other *StackClass) bool {
	return other.ToC() == recv.ToC()
}

// StackSidebarClass is a wrapper around the C record GtkStackSidebarClass.
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

// Equals compares this StackSidebarClass with another StackSidebarClass, and returns true if they represent the same GObject.
func (recv *StackSidebarClass) Equals(other *StackSidebarClass) bool {
	return other.ToC() == recv.ToC()
}

// StackSidebarPrivate is a wrapper around the C record GtkStackSidebarPrivate.
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

// Equals compares this StackSidebarPrivate with another StackSidebarPrivate, and returns true if they represent the same GObject.
func (recv *StackSidebarPrivate) Equals(other *StackSidebarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StackSwitcherClass is a wrapper around the C record GtkStackSwitcherClass.
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

// Equals compares this StackSwitcherClass with another StackSwitcherClass, and returns true if they represent the same GObject.
func (recv *StackSwitcherClass) Equals(other *StackSwitcherClass) bool {
	return other.ToC() == recv.ToC()
}

// StatusIconClass is a wrapper around the C record GtkStatusIconClass.
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

// Equals compares this StatusIconClass with another StatusIconClass, and returns true if they represent the same GObject.
func (recv *StatusIconClass) Equals(other *StatusIconClass) bool {
	return other.ToC() == recv.ToC()
}

// StatusIconPrivate is a wrapper around the C record GtkStatusIconPrivate.
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

// Equals compares this StatusIconPrivate with another StatusIconPrivate, and returns true if they represent the same GObject.
func (recv *StatusIconPrivate) Equals(other *StatusIconPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StatusbarAccessibleClass is a wrapper around the C record GtkStatusbarAccessibleClass.
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

// Equals compares this StatusbarAccessibleClass with another StatusbarAccessibleClass, and returns true if they represent the same GObject.
func (recv *StatusbarAccessibleClass) Equals(other *StatusbarAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// StatusbarAccessiblePrivate is a wrapper around the C record GtkStatusbarAccessiblePrivate.
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

// Equals compares this StatusbarAccessiblePrivate with another StatusbarAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *StatusbarAccessiblePrivate) Equals(other *StatusbarAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// StatusbarClass is a wrapper around the C record GtkStatusbarClass.
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

// Equals compares this StatusbarClass with another StatusbarClass, and returns true if they represent the same GObject.
func (recv *StatusbarClass) Equals(other *StatusbarClass) bool {
	return other.ToC() == recv.ToC()
}

// StatusbarPrivate is a wrapper around the C record GtkStatusbarPrivate.
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

// Equals compares this StatusbarPrivate with another StatusbarPrivate, and returns true if they represent the same GObject.
func (recv *StatusbarPrivate) Equals(other *StatusbarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StockItem is a wrapper around the C record GtkStockItem.
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

// Equals compares this StockItem with another StockItem, and returns true if they represent the same GObject.
func (recv *StockItem) Equals(other *StockItem) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gtk_stock_item_copy.
func (recv *StockItem) Copy() *StockItem {
	retC := C.gtk_stock_item_copy((*C.GtkStockItem)(recv.native))
	retGo := StockItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_stock_item_free.
func (recv *StockItem) Free() {
	C.gtk_stock_item_free((*C.GtkStockItem)(recv.native))

	return
}

// StyleClass is a wrapper around the C record GtkStyleClass.
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

// Equals compares this StyleClass with another StyleClass, and returns true if they represent the same GObject.
func (recv *StyleClass) Equals(other *StyleClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleContextClass is a wrapper around the C record GtkStyleContextClass.
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

// Equals compares this StyleContextClass with another StyleContextClass, and returns true if they represent the same GObject.
func (recv *StyleContextClass) Equals(other *StyleContextClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleContextPrivate is a wrapper around the C record GtkStyleContextPrivate.
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

// Equals compares this StyleContextPrivate with another StyleContextPrivate, and returns true if they represent the same GObject.
func (recv *StyleContextPrivate) Equals(other *StyleContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StylePropertiesClass is a wrapper around the C record GtkStylePropertiesClass.
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

// Equals compares this StylePropertiesClass with another StylePropertiesClass, and returns true if they represent the same GObject.
func (recv *StylePropertiesClass) Equals(other *StylePropertiesClass) bool {
	return other.ToC() == recv.ToC()
}

// StylePropertiesPrivate is a wrapper around the C record GtkStylePropertiesPrivate.
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

// Equals compares this StylePropertiesPrivate with another StylePropertiesPrivate, and returns true if they represent the same GObject.
func (recv *StylePropertiesPrivate) Equals(other *StylePropertiesPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StyleProviderIface is a wrapper around the C record GtkStyleProviderIface.
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

// Equals compares this StyleProviderIface with another StyleProviderIface, and returns true if they represent the same GObject.
func (recv *StyleProviderIface) Equals(other *StyleProviderIface) bool {
	return other.ToC() == recv.ToC()
}

// SwitchAccessibleClass is a wrapper around the C record GtkSwitchAccessibleClass.
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

// Equals compares this SwitchAccessibleClass with another SwitchAccessibleClass, and returns true if they represent the same GObject.
func (recv *SwitchAccessibleClass) Equals(other *SwitchAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// SwitchAccessiblePrivate is a wrapper around the C record GtkSwitchAccessiblePrivate.
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

// Equals compares this SwitchAccessiblePrivate with another SwitchAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *SwitchAccessiblePrivate) Equals(other *SwitchAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// SwitchClass is a wrapper around the C record GtkSwitchClass.
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

// Equals compares this SwitchClass with another SwitchClass, and returns true if they represent the same GObject.
func (recv *SwitchClass) Equals(other *SwitchClass) bool {
	return other.ToC() == recv.ToC()
}

// SwitchPrivate is a wrapper around the C record GtkSwitchPrivate.
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

// Equals compares this SwitchPrivate with another SwitchPrivate, and returns true if they represent the same GObject.
func (recv *SwitchPrivate) Equals(other *SwitchPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SymbolicColor is a wrapper around the C record GtkSymbolicColor.
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

// Equals compares this SymbolicColor with another SymbolicColor, and returns true if they represent the same GObject.
func (recv *SymbolicColor) Equals(other *SymbolicColor) bool {
	return other.ToC() == recv.ToC()
}

// ToString is a wrapper around the C function gtk_symbolic_color_to_string.
func (recv *SymbolicColor) ToString() string {
	retC := C.gtk_symbolic_color_to_string((*C.GtkSymbolicColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TableChild is a wrapper around the C record GtkTableChild.
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

// Equals compares this TableChild with another TableChild, and returns true if they represent the same GObject.
func (recv *TableChild) Equals(other *TableChild) bool {
	return other.ToC() == recv.ToC()
}

// TableClass is a wrapper around the C record GtkTableClass.
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

// Equals compares this TableClass with another TableClass, and returns true if they represent the same GObject.
func (recv *TableClass) Equals(other *TableClass) bool {
	return other.ToC() == recv.ToC()
}

// TablePrivate is a wrapper around the C record GtkTablePrivate.
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

// Equals compares this TablePrivate with another TablePrivate, and returns true if they represent the same GObject.
func (recv *TablePrivate) Equals(other *TablePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TableRowCol is a wrapper around the C record GtkTableRowCol.
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

// Equals compares this TableRowCol with another TableRowCol, and returns true if they represent the same GObject.
func (recv *TableRowCol) Equals(other *TableRowCol) bool {
	return other.ToC() == recv.ToC()
}

// TargetEntry is a wrapper around the C record GtkTargetEntry.
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

// Equals compares this TargetEntry with another TargetEntry, and returns true if they represent the same GObject.
func (recv *TargetEntry) Equals(other *TargetEntry) bool {
	return other.ToC() == recv.ToC()
}

// TargetEntryNew is a wrapper around the C function gtk_target_entry_new.
func TargetEntryNew(target string, flags uint32, info uint32) *TargetEntry {
	c_target := C.CString(target)
	defer C.free(unsafe.Pointer(c_target))

	c_flags := (C.guint)(flags)

	c_info := (C.guint)(info)

	retC := C.gtk_target_entry_new(c_target, c_flags, c_info)
	retGo := TargetEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_target_entry_copy.
func (recv *TargetEntry) Copy() *TargetEntry {
	retC := C.gtk_target_entry_copy((*C.GtkTargetEntry)(recv.native))
	retGo := TargetEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_target_entry_free.
func (recv *TargetEntry) Free() {
	C.gtk_target_entry_free((*C.GtkTargetEntry)(recv.native))

	return
}

// TargetList is a wrapper around the C record GtkTargetList.
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

// Equals compares this TargetList with another TargetList, and returns true if they represent the same GObject.
func (recv *TargetList) Equals(other *TargetList) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : gtk_target_list_new : unsupported parameter targets :

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_add_table : unsupported parameter targets :

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Ref is a wrapper around the C function gtk_target_list_ref.
func (recv *TargetList) Ref() *TargetList {
	retC := C.gtk_target_list_ref((*C.GtkTargetList)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

// Unref is a wrapper around the C function gtk_target_list_unref.
func (recv *TargetList) Unref() {
	C.gtk_target_list_unref((*C.GtkTargetList)(recv.native))

	return
}

// TargetPair is a wrapper around the C record GtkTargetPair.
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

// Equals compares this TargetPair with another TargetPair, and returns true if they represent the same GObject.
func (recv *TargetPair) Equals(other *TargetPair) bool {
	return other.ToC() == recv.ToC()
}

// TearoffMenuItemClass is a wrapper around the C record GtkTearoffMenuItemClass.
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

// Equals compares this TearoffMenuItemClass with another TearoffMenuItemClass, and returns true if they represent the same GObject.
func (recv *TearoffMenuItemClass) Equals(other *TearoffMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// TearoffMenuItemPrivate is a wrapper around the C record GtkTearoffMenuItemPrivate.
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

// Equals compares this TearoffMenuItemPrivate with another TearoffMenuItemPrivate, and returns true if they represent the same GObject.
func (recv *TearoffMenuItemPrivate) Equals(other *TearoffMenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextAppearance is a wrapper around the C record GtkTextAppearance.
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

// Equals compares this TextAppearance with another TextAppearance, and returns true if they represent the same GObject.
func (recv *TextAppearance) Equals(other *TextAppearance) bool {
	return other.ToC() == recv.ToC()
}

// TextAttributes is a wrapper around the C record GtkTextAttributes.
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

// Equals compares this TextAttributes with another TextAttributes, and returns true if they represent the same GObject.
func (recv *TextAttributes) Equals(other *TextAttributes) bool {
	return other.ToC() == recv.ToC()
}

// TextAttributesNew is a wrapper around the C function gtk_text_attributes_new.
func TextAttributesNew() *TextAttributes {
	retC := C.gtk_text_attributes_new()
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_text_attributes_copy.
func (recv *TextAttributes) Copy() *TextAttributes {
	retC := C.gtk_text_attributes_copy((*C.GtkTextAttributes)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CopyValues is a wrapper around the C function gtk_text_attributes_copy_values.
func (recv *TextAttributes) CopyValues(dest *TextAttributes) {
	c_dest := (*C.GtkTextAttributes)(C.NULL)
	if dest != nil {
		c_dest = (*C.GtkTextAttributes)(dest.ToC())
	}

	C.gtk_text_attributes_copy_values((*C.GtkTextAttributes)(recv.native), c_dest)

	return
}

// Ref is a wrapper around the C function gtk_text_attributes_ref.
func (recv *TextAttributes) Ref() *TextAttributes {
	retC := C.gtk_text_attributes_ref((*C.GtkTextAttributes)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_text_attributes_unref.
func (recv *TextAttributes) Unref() {
	C.gtk_text_attributes_unref((*C.GtkTextAttributes)(recv.native))

	return
}

// TextBTree is a wrapper around the C record GtkTextBTree.
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

// Equals compares this TextBTree with another TextBTree, and returns true if they represent the same GObject.
func (recv *TextBTree) Equals(other *TextBTree) bool {
	return other.ToC() == recv.ToC()
}

// TextBufferClass is a wrapper around the C record GtkTextBufferClass.
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

// Equals compares this TextBufferClass with another TextBufferClass, and returns true if they represent the same GObject.
func (recv *TextBufferClass) Equals(other *TextBufferClass) bool {
	return other.ToC() == recv.ToC()
}

// TextBufferPrivate is a wrapper around the C record GtkTextBufferPrivate.
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

// Equals compares this TextBufferPrivate with another TextBufferPrivate, and returns true if they represent the same GObject.
func (recv *TextBufferPrivate) Equals(other *TextBufferPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextCellAccessibleClass is a wrapper around the C record GtkTextCellAccessibleClass.
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

// Equals compares this TextCellAccessibleClass with another TextCellAccessibleClass, and returns true if they represent the same GObject.
func (recv *TextCellAccessibleClass) Equals(other *TextCellAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// TextCellAccessiblePrivate is a wrapper around the C record GtkTextCellAccessiblePrivate.
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

// Equals compares this TextCellAccessiblePrivate with another TextCellAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *TextCellAccessiblePrivate) Equals(other *TextCellAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextChildAnchorClass is a wrapper around the C record GtkTextChildAnchorClass.
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

// Equals compares this TextChildAnchorClass with another TextChildAnchorClass, and returns true if they represent the same GObject.
func (recv *TextChildAnchorClass) Equals(other *TextChildAnchorClass) bool {
	return other.ToC() == recv.ToC()
}

// TextIter is a wrapper around the C record GtkTextIter.
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

// Equals compares this TextIter with another TextIter, and returns true if they represent the same GObject.
func (recv *TextIter) Equals(other *TextIter) bool {
	return other.ToC() == recv.ToC()
}

// BackwardChar is a wrapper around the C function gtk_text_iter_backward_char.
func (recv *TextIter) BackwardChar() bool {
	retC := C.gtk_text_iter_backward_char((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardChars is a wrapper around the C function gtk_text_iter_backward_chars.
func (recv *TextIter) BackwardChars(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_chars((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardCursorPosition is a wrapper around the C function gtk_text_iter_backward_cursor_position.
func (recv *TextIter) BackwardCursorPosition() bool {
	retC := C.gtk_text_iter_backward_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardCursorPositions is a wrapper around the C function gtk_text_iter_backward_cursor_positions.
func (recv *TextIter) BackwardCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_iter_backward_find_char : unsupported parameter pred : no type generator for TextCharPredicate (GtkTextCharPredicate) for param pred

// BackwardLine is a wrapper around the C function gtk_text_iter_backward_line.
func (recv *TextIter) BackwardLine() bool {
	retC := C.gtk_text_iter_backward_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardLines is a wrapper around the C function gtk_text_iter_backward_lines.
func (recv *TextIter) BackwardLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardSearch is a wrapper around the C function gtk_text_iter_backward_search.
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

// BackwardSentenceStart is a wrapper around the C function gtk_text_iter_backward_sentence_start.
func (recv *TextIter) BackwardSentenceStart() bool {
	retC := C.gtk_text_iter_backward_sentence_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardSentenceStarts is a wrapper around the C function gtk_text_iter_backward_sentence_starts.
func (recv *TextIter) BackwardSentenceStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_sentence_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardToTagToggle is a wrapper around the C function gtk_text_iter_backward_to_tag_toggle.
func (recv *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_backward_to_tag_toggle((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardWordStart is a wrapper around the C function gtk_text_iter_backward_word_start.
func (recv *TextIter) BackwardWordStart() bool {
	retC := C.gtk_text_iter_backward_word_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardWordStarts is a wrapper around the C function gtk_text_iter_backward_word_starts.
func (recv *TextIter) BackwardWordStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_word_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BeginsTag is a wrapper around the C function gtk_text_iter_begins_tag.
func (recv *TextIter) BeginsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_begins_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// CanInsert is a wrapper around the C function gtk_text_iter_can_insert.
func (recv *TextIter) CanInsert(defaultEditability bool) bool {
	c_default_editability :=
		boolToGboolean(defaultEditability)

	retC := C.gtk_text_iter_can_insert((*C.GtkTextIter)(recv.native), c_default_editability)
	retGo := retC == C.TRUE

	return retGo
}

// Compare is a wrapper around the C function gtk_text_iter_compare.
func (recv *TextIter) Compare(rhs *TextIter) int32 {
	c_rhs := (*C.GtkTextIter)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GtkTextIter)(rhs.ToC())
	}

	retC := C.gtk_text_iter_compare((*C.GtkTextIter)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// Copy is a wrapper around the C function gtk_text_iter_copy.
func (recv *TextIter) Copy() *TextIter {
	retC := C.gtk_text_iter_copy((*C.GtkTextIter)(recv.native))
	retGo := TextIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Editable is a wrapper around the C function gtk_text_iter_editable.
func (recv *TextIter) Editable(defaultSetting bool) bool {
	c_default_setting :=
		boolToGboolean(defaultSetting)

	retC := C.gtk_text_iter_editable((*C.GtkTextIter)(recv.native), c_default_setting)
	retGo := retC == C.TRUE

	return retGo
}

// EndsLine is a wrapper around the C function gtk_text_iter_ends_line.
func (recv *TextIter) EndsLine() bool {
	retC := C.gtk_text_iter_ends_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// EndsSentence is a wrapper around the C function gtk_text_iter_ends_sentence.
func (recv *TextIter) EndsSentence() bool {
	retC := C.gtk_text_iter_ends_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// EndsTag is a wrapper around the C function gtk_text_iter_ends_tag.
func (recv *TextIter) EndsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_ends_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// EndsWord is a wrapper around the C function gtk_text_iter_ends_word.
func (recv *TextIter) EndsWord() bool {
	retC := C.gtk_text_iter_ends_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Equal is a wrapper around the C function gtk_text_iter_equal.
func (recv *TextIter) Equal(rhs *TextIter) bool {
	c_rhs := (*C.GtkTextIter)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GtkTextIter)(rhs.ToC())
	}

	retC := C.gtk_text_iter_equal((*C.GtkTextIter)(recv.native), c_rhs)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardChar is a wrapper around the C function gtk_text_iter_forward_char.
func (recv *TextIter) ForwardChar() bool {
	retC := C.gtk_text_iter_forward_char((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardChars is a wrapper around the C function gtk_text_iter_forward_chars.
func (recv *TextIter) ForwardChars(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_chars((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardCursorPosition is a wrapper around the C function gtk_text_iter_forward_cursor_position.
func (recv *TextIter) ForwardCursorPosition() bool {
	retC := C.gtk_text_iter_forward_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardCursorPositions is a wrapper around the C function gtk_text_iter_forward_cursor_positions.
func (recv *TextIter) ForwardCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_iter_forward_find_char : unsupported parameter pred : no type generator for TextCharPredicate (GtkTextCharPredicate) for param pred

// ForwardLine is a wrapper around the C function gtk_text_iter_forward_line.
func (recv *TextIter) ForwardLine() bool {
	retC := C.gtk_text_iter_forward_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardLines is a wrapper around the C function gtk_text_iter_forward_lines.
func (recv *TextIter) ForwardLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardSearch is a wrapper around the C function gtk_text_iter_forward_search.
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

// ForwardSentenceEnd is a wrapper around the C function gtk_text_iter_forward_sentence_end.
func (recv *TextIter) ForwardSentenceEnd() bool {
	retC := C.gtk_text_iter_forward_sentence_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardSentenceEnds is a wrapper around the C function gtk_text_iter_forward_sentence_ends.
func (recv *TextIter) ForwardSentenceEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_sentence_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardToEnd is a wrapper around the C function gtk_text_iter_forward_to_end.
func (recv *TextIter) ForwardToEnd() {
	C.gtk_text_iter_forward_to_end((*C.GtkTextIter)(recv.native))

	return
}

// ForwardToLineEnd is a wrapper around the C function gtk_text_iter_forward_to_line_end.
func (recv *TextIter) ForwardToLineEnd() bool {
	retC := C.gtk_text_iter_forward_to_line_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardToTagToggle is a wrapper around the C function gtk_text_iter_forward_to_tag_toggle.
func (recv *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_forward_to_tag_toggle((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardWordEnd is a wrapper around the C function gtk_text_iter_forward_word_end.
func (recv *TextIter) ForwardWordEnd() bool {
	retC := C.gtk_text_iter_forward_word_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardWordEnds is a wrapper around the C function gtk_text_iter_forward_word_ends.
func (recv *TextIter) ForwardWordEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_word_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gtk_text_iter_free.
func (recv *TextIter) Free() {
	C.gtk_text_iter_free((*C.GtkTextIter)(recv.native))

	return
}

// GetAttributes is a wrapper around the C function gtk_text_iter_get_attributes.
func (recv *TextIter) GetAttributes() (bool, *TextAttributes) {
	var c_values C.GtkTextAttributes

	retC := C.gtk_text_iter_get_attributes((*C.GtkTextIter)(recv.native), &c_values)
	retGo := retC == C.TRUE

	values := TextAttributesNewFromC(unsafe.Pointer(&c_values))

	return retGo, values
}

// GetBuffer is a wrapper around the C function gtk_text_iter_get_buffer.
func (recv *TextIter) GetBuffer() *TextBuffer {
	retC := C.gtk_text_iter_get_buffer((*C.GtkTextIter)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBytesInLine is a wrapper around the C function gtk_text_iter_get_bytes_in_line.
func (recv *TextIter) GetBytesInLine() int32 {
	retC := C.gtk_text_iter_get_bytes_in_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetChar is a wrapper around the C function gtk_text_iter_get_char.
func (recv *TextIter) GetChar() rune {
	retC := C.gtk_text_iter_get_char((*C.GtkTextIter)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetCharsInLine is a wrapper around the C function gtk_text_iter_get_chars_in_line.
func (recv *TextIter) GetCharsInLine() int32 {
	retC := C.gtk_text_iter_get_chars_in_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetChildAnchor is a wrapper around the C function gtk_text_iter_get_child_anchor.
func (recv *TextIter) GetChildAnchor() *TextChildAnchor {
	retC := C.gtk_text_iter_get_child_anchor((*C.GtkTextIter)(recv.native))
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLanguage is a wrapper around the C function gtk_text_iter_get_language.
func (recv *TextIter) GetLanguage() *pango.Language {
	retC := C.gtk_text_iter_get_language((*C.GtkTextIter)(recv.native))
	retGo := pango.LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLine is a wrapper around the C function gtk_text_iter_get_line.
func (recv *TextIter) GetLine() int32 {
	retC := C.gtk_text_iter_get_line((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLineIndex is a wrapper around the C function gtk_text_iter_get_line_index.
func (recv *TextIter) GetLineIndex() int32 {
	retC := C.gtk_text_iter_get_line_index((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLineOffset is a wrapper around the C function gtk_text_iter_get_line_offset.
func (recv *TextIter) GetLineOffset() int32 {
	retC := C.gtk_text_iter_get_line_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarks is a wrapper around the C function gtk_text_iter_get_marks.
func (recv *TextIter) GetMarks() *glib.SList {
	retC := C.gtk_text_iter_get_marks((*C.GtkTextIter)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOffset is a wrapper around the C function gtk_text_iter_get_offset.
func (recv *TextIter) GetOffset() int32 {
	retC := C.gtk_text_iter_get_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_text_iter_get_pixbuf.
func (recv *TextIter) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_text_iter_get_pixbuf((*C.GtkTextIter)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSlice is a wrapper around the C function gtk_text_iter_get_slice.
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

// GetTags is a wrapper around the C function gtk_text_iter_get_tags.
func (recv *TextIter) GetTags() *glib.SList {
	retC := C.gtk_text_iter_get_tags((*C.GtkTextIter)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetText is a wrapper around the C function gtk_text_iter_get_text.
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

// GetToggledTags is a wrapper around the C function gtk_text_iter_get_toggled_tags.
func (recv *TextIter) GetToggledTags(toggledOn bool) *glib.SList {
	c_toggled_on :=
		boolToGboolean(toggledOn)

	retC := C.gtk_text_iter_get_toggled_tags((*C.GtkTextIter)(recv.native), c_toggled_on)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisibleLineIndex is a wrapper around the C function gtk_text_iter_get_visible_line_index.
func (recv *TextIter) GetVisibleLineIndex() int32 {
	retC := C.gtk_text_iter_get_visible_line_index((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetVisibleLineOffset is a wrapper around the C function gtk_text_iter_get_visible_line_offset.
func (recv *TextIter) GetVisibleLineOffset() int32 {
	retC := C.gtk_text_iter_get_visible_line_offset((*C.GtkTextIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetVisibleSlice is a wrapper around the C function gtk_text_iter_get_visible_slice.
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

// GetVisibleText is a wrapper around the C function gtk_text_iter_get_visible_text.
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

// HasTag is a wrapper around the C function gtk_text_iter_has_tag.
func (recv *TextIter) HasTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_has_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// InRange is a wrapper around the C function gtk_text_iter_in_range.
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

// InsideSentence is a wrapper around the C function gtk_text_iter_inside_sentence.
func (recv *TextIter) InsideSentence() bool {
	retC := C.gtk_text_iter_inside_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// InsideWord is a wrapper around the C function gtk_text_iter_inside_word.
func (recv *TextIter) InsideWord() bool {
	retC := C.gtk_text_iter_inside_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsCursorPosition is a wrapper around the C function gtk_text_iter_is_cursor_position.
func (recv *TextIter) IsCursorPosition() bool {
	retC := C.gtk_text_iter_is_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsEnd is a wrapper around the C function gtk_text_iter_is_end.
func (recv *TextIter) IsEnd() bool {
	retC := C.gtk_text_iter_is_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsStart is a wrapper around the C function gtk_text_iter_is_start.
func (recv *TextIter) IsStart() bool {
	retC := C.gtk_text_iter_is_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Order is a wrapper around the C function gtk_text_iter_order.
func (recv *TextIter) Order(second *TextIter) {
	c_second := (*C.GtkTextIter)(C.NULL)
	if second != nil {
		c_second = (*C.GtkTextIter)(second.ToC())
	}

	C.gtk_text_iter_order((*C.GtkTextIter)(recv.native), c_second)

	return
}

// SetLine is a wrapper around the C function gtk_text_iter_set_line.
func (recv *TextIter) SetLine(lineNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	C.gtk_text_iter_set_line((*C.GtkTextIter)(recv.native), c_line_number)

	return
}

// SetLineIndex is a wrapper around the C function gtk_text_iter_set_line_index.
func (recv *TextIter) SetLineIndex(byteOnLine int32) {
	c_byte_on_line := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_line_index((*C.GtkTextIter)(recv.native), c_byte_on_line)

	return
}

// SetLineOffset is a wrapper around the C function gtk_text_iter_set_line_offset.
func (recv *TextIter) SetLineOffset(charOnLine int32) {
	c_char_on_line := (C.gint)(charOnLine)

	C.gtk_text_iter_set_line_offset((*C.GtkTextIter)(recv.native), c_char_on_line)

	return
}

// SetOffset is a wrapper around the C function gtk_text_iter_set_offset.
func (recv *TextIter) SetOffset(charOffset int32) {
	c_char_offset := (C.gint)(charOffset)

	C.gtk_text_iter_set_offset((*C.GtkTextIter)(recv.native), c_char_offset)

	return
}

// SetVisibleLineIndex is a wrapper around the C function gtk_text_iter_set_visible_line_index.
func (recv *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	c_byte_on_line := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_visible_line_index((*C.GtkTextIter)(recv.native), c_byte_on_line)

	return
}

// SetVisibleLineOffset is a wrapper around the C function gtk_text_iter_set_visible_line_offset.
func (recv *TextIter) SetVisibleLineOffset(charOnLine int32) {
	c_char_on_line := (C.gint)(charOnLine)

	C.gtk_text_iter_set_visible_line_offset((*C.GtkTextIter)(recv.native), c_char_on_line)

	return
}

// StartsLine is a wrapper around the C function gtk_text_iter_starts_line.
func (recv *TextIter) StartsLine() bool {
	retC := C.gtk_text_iter_starts_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// StartsSentence is a wrapper around the C function gtk_text_iter_starts_sentence.
func (recv *TextIter) StartsSentence() bool {
	retC := C.gtk_text_iter_starts_sentence((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// StartsWord is a wrapper around the C function gtk_text_iter_starts_word.
func (recv *TextIter) StartsWord() bool {
	retC := C.gtk_text_iter_starts_word((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TogglesTag is a wrapper around the C function gtk_text_iter_toggles_tag.
func (recv *TextIter) TogglesTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_toggles_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// TextMarkClass is a wrapper around the C record GtkTextMarkClass.
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

// Equals compares this TextMarkClass with another TextMarkClass, and returns true if they represent the same GObject.
func (recv *TextMarkClass) Equals(other *TextMarkClass) bool {
	return other.ToC() == recv.ToC()
}

// TextTagClass is a wrapper around the C record GtkTextTagClass.
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

// Equals compares this TextTagClass with another TextTagClass, and returns true if they represent the same GObject.
func (recv *TextTagClass) Equals(other *TextTagClass) bool {
	return other.ToC() == recv.ToC()
}

// TextTagPrivate is a wrapper around the C record GtkTextTagPrivate.
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

// Equals compares this TextTagPrivate with another TextTagPrivate, and returns true if they represent the same GObject.
func (recv *TextTagPrivate) Equals(other *TextTagPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextTagTableClass is a wrapper around the C record GtkTextTagTableClass.
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

// Equals compares this TextTagTableClass with another TextTagTableClass, and returns true if they represent the same GObject.
func (recv *TextTagTableClass) Equals(other *TextTagTableClass) bool {
	return other.ToC() == recv.ToC()
}

// TextTagTablePrivate is a wrapper around the C record GtkTextTagTablePrivate.
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

// Equals compares this TextTagTablePrivate with another TextTagTablePrivate, and returns true if they represent the same GObject.
func (recv *TextTagTablePrivate) Equals(other *TextTagTablePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextViewAccessibleClass is a wrapper around the C record GtkTextViewAccessibleClass.
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

// Equals compares this TextViewAccessibleClass with another TextViewAccessibleClass, and returns true if they represent the same GObject.
func (recv *TextViewAccessibleClass) Equals(other *TextViewAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// TextViewAccessiblePrivate is a wrapper around the C record GtkTextViewAccessiblePrivate.
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

// Equals compares this TextViewAccessiblePrivate with another TextViewAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *TextViewAccessiblePrivate) Equals(other *TextViewAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TextViewClass is a wrapper around the C record GtkTextViewClass.
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

// Equals compares this TextViewClass with another TextViewClass, and returns true if they represent the same GObject.
func (recv *TextViewClass) Equals(other *TextViewClass) bool {
	return other.ToC() == recv.ToC()
}

// TextViewPrivate is a wrapper around the C record GtkTextViewPrivate.
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

// Equals compares this TextViewPrivate with another TextViewPrivate, and returns true if they represent the same GObject.
func (recv *TextViewPrivate) Equals(other *TextViewPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ThemeEngine is a wrapper around the C record GtkThemeEngine.
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

// Equals compares this ThemeEngine with another ThemeEngine, and returns true if they represent the same GObject.
func (recv *ThemeEngine) Equals(other *ThemeEngine) bool {
	return other.ToC() == recv.ToC()
}

// ThemingEngineClass is a wrapper around the C record GtkThemingEngineClass.
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

// Equals compares this ThemingEngineClass with another ThemingEngineClass, and returns true if they represent the same GObject.
func (recv *ThemingEngineClass) Equals(other *ThemingEngineClass) bool {
	return other.ToC() == recv.ToC()
}

// ThemingEnginePrivate is a wrapper around the C record GtkThemingEnginePrivate.
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

// Equals compares this ThemingEnginePrivate with another ThemingEnginePrivate, and returns true if they represent the same GObject.
func (recv *ThemingEnginePrivate) Equals(other *ThemingEnginePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToggleActionClass is a wrapper around the C record GtkToggleActionClass.
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

// Equals compares this ToggleActionClass with another ToggleActionClass, and returns true if they represent the same GObject.
func (recv *ToggleActionClass) Equals(other *ToggleActionClass) bool {
	return other.ToC() == recv.ToC()
}

// ToggleActionEntry is a wrapper around the C record GtkToggleActionEntry.
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

// Equals compares this ToggleActionEntry with another ToggleActionEntry, and returns true if they represent the same GObject.
func (recv *ToggleActionEntry) Equals(other *ToggleActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// ToggleActionPrivate is a wrapper around the C record GtkToggleActionPrivate.
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

// Equals compares this ToggleActionPrivate with another ToggleActionPrivate, and returns true if they represent the same GObject.
func (recv *ToggleActionPrivate) Equals(other *ToggleActionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToggleButtonAccessibleClass is a wrapper around the C record GtkToggleButtonAccessibleClass.
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

// Equals compares this ToggleButtonAccessibleClass with another ToggleButtonAccessibleClass, and returns true if they represent the same GObject.
func (recv *ToggleButtonAccessibleClass) Equals(other *ToggleButtonAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ToggleButtonAccessiblePrivate is a wrapper around the C record GtkToggleButtonAccessiblePrivate.
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

// Equals compares this ToggleButtonAccessiblePrivate with another ToggleButtonAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ToggleButtonAccessiblePrivate) Equals(other *ToggleButtonAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToggleButtonClass is a wrapper around the C record GtkToggleButtonClass.
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

// Equals compares this ToggleButtonClass with another ToggleButtonClass, and returns true if they represent the same GObject.
func (recv *ToggleButtonClass) Equals(other *ToggleButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ToggleButtonPrivate is a wrapper around the C record GtkToggleButtonPrivate.
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

// Equals compares this ToggleButtonPrivate with another ToggleButtonPrivate, and returns true if they represent the same GObject.
func (recv *ToggleButtonPrivate) Equals(other *ToggleButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToggleToolButtonClass is a wrapper around the C record GtkToggleToolButtonClass.
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

// Equals compares this ToggleToolButtonClass with another ToggleToolButtonClass, and returns true if they represent the same GObject.
func (recv *ToggleToolButtonClass) Equals(other *ToggleToolButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ToggleToolButtonPrivate is a wrapper around the C record GtkToggleToolButtonPrivate.
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

// Equals compares this ToggleToolButtonPrivate with another ToggleToolButtonPrivate, and returns true if they represent the same GObject.
func (recv *ToggleToolButtonPrivate) Equals(other *ToggleToolButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToolButtonClass is a wrapper around the C record GtkToolButtonClass.
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

// Equals compares this ToolButtonClass with another ToolButtonClass, and returns true if they represent the same GObject.
func (recv *ToolButtonClass) Equals(other *ToolButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// ToolButtonPrivate is a wrapper around the C record GtkToolButtonPrivate.
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

// Equals compares this ToolButtonPrivate with another ToolButtonPrivate, and returns true if they represent the same GObject.
func (recv *ToolButtonPrivate) Equals(other *ToolButtonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToolItemClass is a wrapper around the C record GtkToolItemClass.
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

// Equals compares this ToolItemClass with another ToolItemClass, and returns true if they represent the same GObject.
func (recv *ToolItemClass) Equals(other *ToolItemClass) bool {
	return other.ToC() == recv.ToC()
}

// ToolItemGroupClass is a wrapper around the C record GtkToolItemGroupClass.
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

// Equals compares this ToolItemGroupClass with another ToolItemGroupClass, and returns true if they represent the same GObject.
func (recv *ToolItemGroupClass) Equals(other *ToolItemGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// ToolItemGroupPrivate is a wrapper around the C record GtkToolItemGroupPrivate.
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

// Equals compares this ToolItemGroupPrivate with another ToolItemGroupPrivate, and returns true if they represent the same GObject.
func (recv *ToolItemGroupPrivate) Equals(other *ToolItemGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToolItemPrivate is a wrapper around the C record GtkToolItemPrivate.
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

// Equals compares this ToolItemPrivate with another ToolItemPrivate, and returns true if they represent the same GObject.
func (recv *ToolItemPrivate) Equals(other *ToolItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToolPaletteClass is a wrapper around the C record GtkToolPaletteClass.
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

// Equals compares this ToolPaletteClass with another ToolPaletteClass, and returns true if they represent the same GObject.
func (recv *ToolPaletteClass) Equals(other *ToolPaletteClass) bool {
	return other.ToC() == recv.ToC()
}

// ToolPalettePrivate is a wrapper around the C record GtkToolPalettePrivate.
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

// Equals compares this ToolPalettePrivate with another ToolPalettePrivate, and returns true if they represent the same GObject.
func (recv *ToolPalettePrivate) Equals(other *ToolPalettePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToolShellIface is a wrapper around the C record GtkToolShellIface.
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

// Equals compares this ToolShellIface with another ToolShellIface, and returns true if they represent the same GObject.
func (recv *ToolShellIface) Equals(other *ToolShellIface) bool {
	return other.ToC() == recv.ToC()
}

// ToolbarClass is a wrapper around the C record GtkToolbarClass.
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

// Equals compares this ToolbarClass with another ToolbarClass, and returns true if they represent the same GObject.
func (recv *ToolbarClass) Equals(other *ToolbarClass) bool {
	return other.ToC() == recv.ToC()
}

// ToolbarPrivate is a wrapper around the C record GtkToolbarPrivate.
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

// Equals compares this ToolbarPrivate with another ToolbarPrivate, and returns true if they represent the same GObject.
func (recv *ToolbarPrivate) Equals(other *ToolbarPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ToplevelAccessibleClass is a wrapper around the C record GtkToplevelAccessibleClass.
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

// Equals compares this ToplevelAccessibleClass with another ToplevelAccessibleClass, and returns true if they represent the same GObject.
func (recv *ToplevelAccessibleClass) Equals(other *ToplevelAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// ToplevelAccessiblePrivate is a wrapper around the C record GtkToplevelAccessiblePrivate.
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

// Equals compares this ToplevelAccessiblePrivate with another ToplevelAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *ToplevelAccessiblePrivate) Equals(other *ToplevelAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeDragDestIface is a wrapper around the C record GtkTreeDragDestIface.
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

// Equals compares this TreeDragDestIface with another TreeDragDestIface, and returns true if they represent the same GObject.
func (recv *TreeDragDestIface) Equals(other *TreeDragDestIface) bool {
	return other.ToC() == recv.ToC()
}

// TreeDragSourceIface is a wrapper around the C record GtkTreeDragSourceIface.
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

// Equals compares this TreeDragSourceIface with another TreeDragSourceIface, and returns true if they represent the same GObject.
func (recv *TreeDragSourceIface) Equals(other *TreeDragSourceIface) bool {
	return other.ToC() == recv.ToC()
}

// TreeIter is a wrapper around the C record GtkTreeIter.
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

// Equals compares this TreeIter with another TreeIter, and returns true if they represent the same GObject.
func (recv *TreeIter) Equals(other *TreeIter) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gtk_tree_iter_copy.
func (recv *TreeIter) Copy() *TreeIter {
	retC := C.gtk_tree_iter_copy((*C.GtkTreeIter)(recv.native))
	retGo := TreeIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_tree_iter_free.
func (recv *TreeIter) Free() {
	C.gtk_tree_iter_free((*C.GtkTreeIter)(recv.native))

	return
}

// TreeModelFilterClass is a wrapper around the C record GtkTreeModelFilterClass.
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

// Equals compares this TreeModelFilterClass with another TreeModelFilterClass, and returns true if they represent the same GObject.
func (recv *TreeModelFilterClass) Equals(other *TreeModelFilterClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeModelFilterPrivate is a wrapper around the C record GtkTreeModelFilterPrivate.
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

// Equals compares this TreeModelFilterPrivate with another TreeModelFilterPrivate, and returns true if they represent the same GObject.
func (recv *TreeModelFilterPrivate) Equals(other *TreeModelFilterPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeModelIface is a wrapper around the C record GtkTreeModelIface.
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

// Equals compares this TreeModelIface with another TreeModelIface, and returns true if they represent the same GObject.
func (recv *TreeModelIface) Equals(other *TreeModelIface) bool {
	return other.ToC() == recv.ToC()
}

// TreeModelSortClass is a wrapper around the C record GtkTreeModelSortClass.
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

// Equals compares this TreeModelSortClass with another TreeModelSortClass, and returns true if they represent the same GObject.
func (recv *TreeModelSortClass) Equals(other *TreeModelSortClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeModelSortPrivate is a wrapper around the C record GtkTreeModelSortPrivate.
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

// Equals compares this TreeModelSortPrivate with another TreeModelSortPrivate, and returns true if they represent the same GObject.
func (recv *TreeModelSortPrivate) Equals(other *TreeModelSortPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreePath is a wrapper around the C record GtkTreePath.
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

// Equals compares this TreePath with another TreePath, and returns true if they represent the same GObject.
func (recv *TreePath) Equals(other *TreePath) bool {
	return other.ToC() == recv.ToC()
}

// TreePathNew is a wrapper around the C function gtk_tree_path_new.
func TreePathNew() *TreePath {
	retC := C.gtk_tree_path_new()
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreePathNewFirst is a wrapper around the C function gtk_tree_path_new_first.
func TreePathNewFirst() *TreePath {
	retC := C.gtk_tree_path_new_first()
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreePathNewFromString is a wrapper around the C function gtk_tree_path_new_from_string.
func TreePathNewFromString(path string) *TreePath {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_tree_path_new_from_string(c_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendIndex is a wrapper around the C function gtk_tree_path_append_index.
func (recv *TreePath) AppendIndex(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_tree_path_append_index((*C.GtkTreePath)(recv.native), c_index_)

	return
}

// Compare is a wrapper around the C function gtk_tree_path_compare.
func (recv *TreePath) Compare(b *TreePath) int32 {
	c_b := (*C.GtkTreePath)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreePath)(b.ToC())
	}

	retC := C.gtk_tree_path_compare((*C.GtkTreePath)(recv.native), c_b)
	retGo := (int32)(retC)

	return retGo
}

// Copy is a wrapper around the C function gtk_tree_path_copy.
func (recv *TreePath) Copy() *TreePath {
	retC := C.gtk_tree_path_copy((*C.GtkTreePath)(recv.native))
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Down is a wrapper around the C function gtk_tree_path_down.
func (recv *TreePath) Down() {
	C.gtk_tree_path_down((*C.GtkTreePath)(recv.native))

	return
}

// Free is a wrapper around the C function gtk_tree_path_free.
func (recv *TreePath) Free() {
	C.gtk_tree_path_free((*C.GtkTreePath)(recv.native))

	return
}

// GetDepth is a wrapper around the C function gtk_tree_path_get_depth.
func (recv *TreePath) GetDepth() int32 {
	retC := C.gtk_tree_path_get_depth((*C.GtkTreePath)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : gtk_tree_path_get_indices

// IsAncestor is a wrapper around the C function gtk_tree_path_is_ancestor.
func (recv *TreePath) IsAncestor(descendant *TreePath) bool {
	c_descendant := (*C.GtkTreePath)(C.NULL)
	if descendant != nil {
		c_descendant = (*C.GtkTreePath)(descendant.ToC())
	}

	retC := C.gtk_tree_path_is_ancestor((*C.GtkTreePath)(recv.native), c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// IsDescendant is a wrapper around the C function gtk_tree_path_is_descendant.
func (recv *TreePath) IsDescendant(ancestor *TreePath) bool {
	c_ancestor := (*C.GtkTreePath)(C.NULL)
	if ancestor != nil {
		c_ancestor = (*C.GtkTreePath)(ancestor.ToC())
	}

	retC := C.gtk_tree_path_is_descendant((*C.GtkTreePath)(recv.native), c_ancestor)
	retGo := retC == C.TRUE

	return retGo
}

// Next is a wrapper around the C function gtk_tree_path_next.
func (recv *TreePath) Next() {
	C.gtk_tree_path_next((*C.GtkTreePath)(recv.native))

	return
}

// PrependIndex is a wrapper around the C function gtk_tree_path_prepend_index.
func (recv *TreePath) PrependIndex(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_tree_path_prepend_index((*C.GtkTreePath)(recv.native), c_index_)

	return
}

// Prev is a wrapper around the C function gtk_tree_path_prev.
func (recv *TreePath) Prev() bool {
	retC := C.gtk_tree_path_prev((*C.GtkTreePath)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function gtk_tree_path_to_string.
func (recv *TreePath) ToString() string {
	retC := C.gtk_tree_path_to_string((*C.GtkTreePath)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Up is a wrapper around the C function gtk_tree_path_up.
func (recv *TreePath) Up() bool {
	retC := C.gtk_tree_path_up((*C.GtkTreePath)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TreeRowReference is a wrapper around the C record GtkTreeRowReference.
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

// Equals compares this TreeRowReference with another TreeRowReference, and returns true if they represent the same GObject.
func (recv *TreeRowReference) Equals(other *TreeRowReference) bool {
	return other.ToC() == recv.ToC()
}

// TreeRowReferenceNew is a wrapper around the C function gtk_tree_row_reference_new.
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

// TreeRowReferenceNewProxy is a wrapper around the C function gtk_tree_row_reference_new_proxy.
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

// TreeRowReferenceDeleted is a wrapper around the C function gtk_tree_row_reference_deleted.
func TreeRowReferenceDeleted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_row_reference_deleted(c_proxy, c_path)

	return
}

// TreeRowReferenceInserted is a wrapper around the C function gtk_tree_row_reference_inserted.
func TreeRowReferenceInserted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_row_reference_inserted(c_proxy, c_path)

	return
}

// TreeRowReferenceReordered is a wrapper around the C function gtk_tree_row_reference_reordered.
func TreeRowReferenceReordered(proxy *gobject.Object, path *TreePath, iter *TreeIter, newOrder []int32) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order := &newOrder[0]

	C.gtk_tree_row_reference_reordered(c_proxy, c_path, c_iter, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Free is a wrapper around the C function gtk_tree_row_reference_free.
func (recv *TreeRowReference) Free() {
	C.gtk_tree_row_reference_free((*C.GtkTreeRowReference)(recv.native))

	return
}

// GetPath is a wrapper around the C function gtk_tree_row_reference_get_path.
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

// Valid is a wrapper around the C function gtk_tree_row_reference_valid.
func (recv *TreeRowReference) Valid() bool {
	retC := C.gtk_tree_row_reference_valid((*C.GtkTreeRowReference)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TreeSelectionClass is a wrapper around the C record GtkTreeSelectionClass.
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

// Equals compares this TreeSelectionClass with another TreeSelectionClass, and returns true if they represent the same GObject.
func (recv *TreeSelectionClass) Equals(other *TreeSelectionClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeSelectionPrivate is a wrapper around the C record GtkTreeSelectionPrivate.
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

// Equals compares this TreeSelectionPrivate with another TreeSelectionPrivate, and returns true if they represent the same GObject.
func (recv *TreeSelectionPrivate) Equals(other *TreeSelectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeSortableIface is a wrapper around the C record GtkTreeSortableIface.
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

// Equals compares this TreeSortableIface with another TreeSortableIface, and returns true if they represent the same GObject.
func (recv *TreeSortableIface) Equals(other *TreeSortableIface) bool {
	return other.ToC() == recv.ToC()
}

// TreeStoreClass is a wrapper around the C record GtkTreeStoreClass.
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

// Equals compares this TreeStoreClass with another TreeStoreClass, and returns true if they represent the same GObject.
func (recv *TreeStoreClass) Equals(other *TreeStoreClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeStorePrivate is a wrapper around the C record GtkTreeStorePrivate.
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

// Equals compares this TreeStorePrivate with another TreeStorePrivate, and returns true if they represent the same GObject.
func (recv *TreeStorePrivate) Equals(other *TreeStorePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewAccessibleClass is a wrapper around the C record GtkTreeViewAccessibleClass.
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

// Equals compares this TreeViewAccessibleClass with another TreeViewAccessibleClass, and returns true if they represent the same GObject.
func (recv *TreeViewAccessibleClass) Equals(other *TreeViewAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewAccessiblePrivate is a wrapper around the C record GtkTreeViewAccessiblePrivate.
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

// Equals compares this TreeViewAccessiblePrivate with another TreeViewAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *TreeViewAccessiblePrivate) Equals(other *TreeViewAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewClass is a wrapper around the C record GtkTreeViewClass.
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

// Equals compares this TreeViewClass with another TreeViewClass, and returns true if they represent the same GObject.
func (recv *TreeViewClass) Equals(other *TreeViewClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewColumnClass is a wrapper around the C record GtkTreeViewColumnClass.
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

// Equals compares this TreeViewColumnClass with another TreeViewColumnClass, and returns true if they represent the same GObject.
func (recv *TreeViewColumnClass) Equals(other *TreeViewColumnClass) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewColumnPrivate is a wrapper around the C record GtkTreeViewColumnPrivate.
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

// Equals compares this TreeViewColumnPrivate with another TreeViewColumnPrivate, and returns true if they represent the same GObject.
func (recv *TreeViewColumnPrivate) Equals(other *TreeViewColumnPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TreeViewPrivate is a wrapper around the C record GtkTreeViewPrivate.
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

// Equals compares this TreeViewPrivate with another TreeViewPrivate, and returns true if they represent the same GObject.
func (recv *TreeViewPrivate) Equals(other *TreeViewPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UIManagerClass is a wrapper around the C record GtkUIManagerClass.
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

// Equals compares this UIManagerClass with another UIManagerClass, and returns true if they represent the same GObject.
func (recv *UIManagerClass) Equals(other *UIManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// UIManagerPrivate is a wrapper around the C record GtkUIManagerPrivate.
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

// Equals compares this UIManagerPrivate with another UIManagerPrivate, and returns true if they represent the same GObject.
func (recv *UIManagerPrivate) Equals(other *UIManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// VBoxClass is a wrapper around the C record GtkVBoxClass.
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

// Equals compares this VBoxClass with another VBoxClass, and returns true if they represent the same GObject.
func (recv *VBoxClass) Equals(other *VBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// VButtonBoxClass is a wrapper around the C record GtkVButtonBoxClass.
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

// Equals compares this VButtonBoxClass with another VButtonBoxClass, and returns true if they represent the same GObject.
func (recv *VButtonBoxClass) Equals(other *VButtonBoxClass) bool {
	return other.ToC() == recv.ToC()
}

// VPanedClass is a wrapper around the C record GtkVPanedClass.
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

// Equals compares this VPanedClass with another VPanedClass, and returns true if they represent the same GObject.
func (recv *VPanedClass) Equals(other *VPanedClass) bool {
	return other.ToC() == recv.ToC()
}

// VScaleClass is a wrapper around the C record GtkVScaleClass.
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

// Equals compares this VScaleClass with another VScaleClass, and returns true if they represent the same GObject.
func (recv *VScaleClass) Equals(other *VScaleClass) bool {
	return other.ToC() == recv.ToC()
}

// VScrollbarClass is a wrapper around the C record GtkVScrollbarClass.
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

// Equals compares this VScrollbarClass with another VScrollbarClass, and returns true if they represent the same GObject.
func (recv *VScrollbarClass) Equals(other *VScrollbarClass) bool {
	return other.ToC() == recv.ToC()
}

// VSeparatorClass is a wrapper around the C record GtkVSeparatorClass.
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

// Equals compares this VSeparatorClass with another VSeparatorClass, and returns true if they represent the same GObject.
func (recv *VSeparatorClass) Equals(other *VSeparatorClass) bool {
	return other.ToC() == recv.ToC()
}

// ViewportClass is a wrapper around the C record GtkViewportClass.
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

// Equals compares this ViewportClass with another ViewportClass, and returns true if they represent the same GObject.
func (recv *ViewportClass) Equals(other *ViewportClass) bool {
	return other.ToC() == recv.ToC()
}

// ViewportPrivate is a wrapper around the C record GtkViewportPrivate.
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

// Equals compares this ViewportPrivate with another ViewportPrivate, and returns true if they represent the same GObject.
func (recv *ViewportPrivate) Equals(other *ViewportPrivate) bool {
	return other.ToC() == recv.ToC()
}

// VolumeButtonClass is a wrapper around the C record GtkVolumeButtonClass.
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

// Equals compares this VolumeButtonClass with another VolumeButtonClass, and returns true if they represent the same GObject.
func (recv *VolumeButtonClass) Equals(other *VolumeButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// WidgetAccessibleClass is a wrapper around the C record GtkWidgetAccessibleClass.
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

// Equals compares this WidgetAccessibleClass with another WidgetAccessibleClass, and returns true if they represent the same GObject.
func (recv *WidgetAccessibleClass) Equals(other *WidgetAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// WidgetAccessiblePrivate is a wrapper around the C record GtkWidgetAccessiblePrivate.
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

// Equals compares this WidgetAccessiblePrivate with another WidgetAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *WidgetAccessiblePrivate) Equals(other *WidgetAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// WidgetClass is a wrapper around the C record GtkWidgetClass.
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

// Equals compares this WidgetClass with another WidgetClass, and returns true if they represent the same GObject.
func (recv *WidgetClass) Equals(other *WidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : gtk_widget_class_install_style_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_widget_class_install_style_property_parser : unsupported parameter pspec : Blacklisted record : GParamSpec

// WidgetClassPrivate is a wrapper around the C record GtkWidgetClassPrivate.
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

// Equals compares this WidgetClassPrivate with another WidgetClassPrivate, and returns true if they represent the same GObject.
func (recv *WidgetClassPrivate) Equals(other *WidgetClassPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WidgetPath is a wrapper around the C record GtkWidgetPath.
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

// Equals compares this WidgetPath with another WidgetPath, and returns true if they represent the same GObject.
func (recv *WidgetPath) Equals(other *WidgetPath) bool {
	return other.ToC() == recv.ToC()
}

// IterGetName is a wrapper around the C function gtk_widget_path_iter_get_name.
func (recv *WidgetPath) IterGetName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// IterGetSiblingIndex is a wrapper around the C function gtk_widget_path_iter_get_sibling_index.
func (recv *WidgetPath) IterGetSiblingIndex(pos int32) uint32 {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_sibling_index((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (uint32)(retC)

	return retGo
}

// IterGetSiblings is a wrapper around the C function gtk_widget_path_iter_get_siblings.
func (recv *WidgetPath) IterGetSiblings(pos int32) *WidgetPath {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_siblings((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// WidgetPrivate is a wrapper around the C record GtkWidgetPrivate.
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

// Equals compares this WidgetPrivate with another WidgetPrivate, and returns true if they represent the same GObject.
func (recv *WidgetPrivate) Equals(other *WidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WindowAccessibleClass is a wrapper around the C record GtkWindowAccessibleClass.
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

// Equals compares this WindowAccessibleClass with another WindowAccessibleClass, and returns true if they represent the same GObject.
func (recv *WindowAccessibleClass) Equals(other *WindowAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// WindowAccessiblePrivate is a wrapper around the C record GtkWindowAccessiblePrivate.
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

// Equals compares this WindowAccessiblePrivate with another WindowAccessiblePrivate, and returns true if they represent the same GObject.
func (recv *WindowAccessiblePrivate) Equals(other *WindowAccessiblePrivate) bool {
	return other.ToC() == recv.ToC()
}

// WindowClass is a wrapper around the C record GtkWindowClass.
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

// Equals compares this WindowClass with another WindowClass, and returns true if they represent the same GObject.
func (recv *WindowClass) Equals(other *WindowClass) bool {
	return other.ToC() == recv.ToC()
}

// WindowGeometryInfo is a wrapper around the C record GtkWindowGeometryInfo.
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

// Equals compares this WindowGeometryInfo with another WindowGeometryInfo, and returns true if they represent the same GObject.
func (recv *WindowGeometryInfo) Equals(other *WindowGeometryInfo) bool {
	return other.ToC() == recv.ToC()
}

// WindowGroupClass is a wrapper around the C record GtkWindowGroupClass.
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

// Equals compares this WindowGroupClass with another WindowGroupClass, and returns true if they represent the same GObject.
func (recv *WindowGroupClass) Equals(other *WindowGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// WindowGroupPrivate is a wrapper around the C record GtkWindowGroupPrivate.
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

// Equals compares this WindowGroupPrivate with another WindowGroupPrivate, and returns true if they represent the same GObject.
func (recv *WindowGroupPrivate) Equals(other *WindowGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WindowPrivate is a wrapper around the C record GtkWindowPrivate.
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

// Equals compares this WindowPrivate with another WindowPrivate, and returns true if they represent the same GObject.
func (recv *WindowPrivate) Equals(other *WindowPrivate) bool {
	return other.ToC() == recv.ToC()
}
