// This is a generated file - DO NOT EDIT

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	static GtkMessageDialog* _gtk_message_dialog_new(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, const gchar* message_format) {
		return gtk_message_dialog_new(parent, flags, type, buttons, message_format);
    }
*/
import "C"

// AboutDialog is a wrapper around the C record GtkAboutDialog.
type AboutDialog struct {
	native *C.GtkAboutDialog
	// parent_instance : record
	// Private : priv
}

func AboutDialogNewFromC(u unsafe.Pointer) *AboutDialog {
	c := (*C.GtkAboutDialog)(u)
	if c == nil {
		return nil
	}

	g := &AboutDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AboutDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AboutDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroup is a wrapper around the C record GtkAccelGroup.
type AccelGroup struct {
	native *C.GtkAccelGroup
	// parent : record
	// priv : record
}

func AccelGroupNewFromC(u unsafe.Pointer) *AccelGroup {
	c := (*C.GtkAccelGroup)(u)
	if c == nil {
		return nil
	}

	g := &AccelGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AccelGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AccelGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroupNew is a wrapper around the C function gtk_accel_group_new.
func AccelGroupNew() *AccelGroup {
	retC := C.gtk_accel_group_new()
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AccelLabel is a wrapper around the C record GtkAccelLabel.
type AccelLabel struct {
	native *C.GtkAccelLabel
	// label : record
	// priv : record
}

func AccelLabelNewFromC(u unsafe.Pointer) *AccelLabel {
	c := (*C.GtkAccelLabel)(u)
	if c == nil {
		return nil
	}

	g := &AccelLabel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AccelLabel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AccelLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelLabelNew is a wrapper around the C function gtk_accel_label_new.
func AccelLabelNew(string_ string) *AccelLabel {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_accel_label_new(c_string)
	retGo := AccelLabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AccelMap is a wrapper around the C record GtkAccelMap.
type AccelMap struct {
	native *C.GtkAccelMap
}

func AccelMapNewFromC(u unsafe.Pointer) *AccelMap {
	c := (*C.GtkAccelMap)(u)
	if c == nil {
		return nil
	}

	g := &AccelMap{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AccelMap) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AccelMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Accessible is a wrapper around the C record GtkAccessible.
type Accessible struct {
	native *C.GtkAccessible
	// parent : record
	// Private : priv
}

func AccessibleNewFromC(u unsafe.Pointer) *Accessible {
	c := (*C.GtkAccessible)(u)
	if c == nil {
		return nil
	}

	g := &Accessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Accessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Accessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action is a wrapper around the C record GtkAction.
type Action struct {
	native *C.GtkAction
	// object : record
	// Private : private_data
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Action) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionBar is a wrapper around the C record GtkActionBar.
type ActionBar struct {
	native *C.GtkActionBar
	// Private : bin
}

func ActionBarNewFromC(u unsafe.Pointer) *ActionBar {
	c := (*C.GtkActionBar)(u)
	if c == nil {
		return nil
	}

	g := &ActionBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ActionBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ActionBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroup is a wrapper around the C record GtkActionGroup.
type ActionGroup struct {
	native *C.GtkActionGroup
	// parent : record
	// Private : priv
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GtkActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Adjustment is a wrapper around the C record GtkAdjustment.
type Adjustment struct {
	native *C.GtkAdjustment
	// parent_instance : record
	// priv : record
}

func AdjustmentNewFromC(u unsafe.Pointer) *Adjustment {
	c := (*C.GtkAdjustment)(u)
	if c == nil {
		return nil
	}

	g := &Adjustment{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Adjustment) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Adjustment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AdjustmentNew is a wrapper around the C function gtk_adjustment_new.
func AdjustmentNew(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) *Adjustment {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_step_increment := (C.gdouble)(stepIncrement)

	c_page_increment := (C.gdouble)(pageIncrement)

	c_page_size := (C.gdouble)(pageSize)

	retC := C.gtk_adjustment_new(c_value, c_lower, c_upper, c_step_increment, c_page_increment, c_page_size)
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Alignment is a wrapper around the C record GtkAlignment.
type Alignment struct {
	native *C.GtkAlignment
	// bin : record
	// Private : priv
}

func AlignmentNewFromC(u unsafe.Pointer) *Alignment {
	c := (*C.GtkAlignment)(u)
	if c == nil {
		return nil
	}

	g := &Alignment{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Alignment) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Alignment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AlignmentNew is a wrapper around the C function gtk_alignment_new.
func AlignmentNew(xalign float32, yalign float32, xscale float32, yscale float32) *Alignment {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	retC := C.gtk_alignment_new(c_xalign, c_yalign, c_xscale, c_yscale)
	retGo := AlignmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserButton is a wrapper around the C record GtkAppChooserButton.
type AppChooserButton struct {
	native *C.GtkAppChooserButton
	// parent : record
	// Private : priv
}

func AppChooserButtonNewFromC(u unsafe.Pointer) *AppChooserButton {
	c := (*C.GtkAppChooserButton)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppChooserButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserDialog is a wrapper around the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native *C.GtkAppChooserDialog
	// parent : record
	// Private : priv
}

func AppChooserDialogNewFromC(u unsafe.Pointer) *AppChooserDialog {
	c := (*C.GtkAppChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppChooserDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppChooserWidget is a wrapper around the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native *C.GtkAppChooserWidget
	// parent : record
	// Private : priv
}

func AppChooserWidgetNewFromC(u unsafe.Pointer) *AppChooserWidget {
	c := (*C.GtkAppChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Application is a wrapper around the C record GtkApplication.
type Application struct {
	native *C.GtkApplication
	// parent : record
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GtkApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Application) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationWindow is a wrapper around the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native *C.GtkApplicationWindow
	// parent_instance : record
	// Private : priv
}

func ApplicationWindowNewFromC(u unsafe.Pointer) *ApplicationWindow {
	c := (*C.GtkApplicationWindow)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationWindow{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ApplicationWindow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ApplicationWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Arrow is a wrapper around the C record GtkArrow.
type Arrow struct {
	native *C.GtkArrow
	// misc : record
	// Private : priv
}

func ArrowNewFromC(u unsafe.Pointer) *Arrow {
	c := (*C.GtkArrow)(u)
	if c == nil {
		return nil
	}

	g := &Arrow{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Arrow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Arrow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowNew is a wrapper around the C function gtk_arrow_new.
func ArrowNew(arrowType ArrowType, shadowType ShadowType) *Arrow {
	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	retC := C.gtk_arrow_new(c_arrow_type, c_shadow_type)
	retGo := ArrowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ArrowAccessible is a wrapper around the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native *C.GtkArrowAccessible
	// parent : record
	// priv : record
}

func ArrowAccessibleNewFromC(u unsafe.Pointer) *ArrowAccessible {
	c := (*C.GtkArrowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ArrowAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ArrowAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ArrowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AspectFrame is a wrapper around the C record GtkAspectFrame.
type AspectFrame struct {
	native *C.GtkAspectFrame
	// frame : record
	// Private : priv
}

func AspectFrameNewFromC(u unsafe.Pointer) *AspectFrame {
	c := (*C.GtkAspectFrame)(u)
	if c == nil {
		return nil
	}

	g := &AspectFrame{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AspectFrame) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AspectFrame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AspectFrameNew is a wrapper around the C function gtk_aspect_frame_new.
func AspectFrameNew(label string, xalign float32, yalign float32, ratio float32, obeyChild bool) *AspectFrame {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obey_child :=
		boolToGboolean(obeyChild)

	retC := C.gtk_aspect_frame_new(c_label, c_xalign, c_yalign, c_ratio, c_obey_child)
	retGo := AspectFrameNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Assistant is a wrapper around the C record GtkAssistant.
type Assistant struct {
	native *C.GtkAssistant
	// parent : record
	// Private : priv
}

func AssistantNewFromC(u unsafe.Pointer) *Assistant {
	c := (*C.GtkAssistant)(u)
	if c == nil {
		return nil
	}

	g := &Assistant{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Assistant) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Assistant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin is a wrapper around the C record GtkBin.
type Bin struct {
	native *C.GtkBin
	// container : record
	// Private : priv
}

func BinNewFromC(u unsafe.Pointer) *Bin {
	c := (*C.GtkBin)(u)
	if c == nil {
		return nil
	}

	g := &Bin{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Bin) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Bin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BooleanCellAccessible is a wrapper around the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native *C.GtkBooleanCellAccessible
	// parent : record
	// priv : record
}

func BooleanCellAccessibleNewFromC(u unsafe.Pointer) *BooleanCellAccessible {
	c := (*C.GtkBooleanCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &BooleanCellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BooleanCellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BooleanCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box is a wrapper around the C record GtkBox.
type Box struct {
	native *C.GtkBox
	// container : record
	// Private : priv
}

func BoxNewFromC(u unsafe.Pointer) *Box {
	c := (*C.GtkBox)(u)
	if c == nil {
		return nil
	}

	g := &Box{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Box) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Box) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Builder is a wrapper around the C record GtkBuilder.
type Builder struct {
	native *C.GtkBuilder
	// parent_instance : record
	// priv : record
}

func BuilderNewFromC(u unsafe.Pointer) *Builder {
	c := (*C.GtkBuilder)(u)
	if c == nil {
		return nil
	}

	g := &Builder{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Builder) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Builder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button is a wrapper around the C record GtkButton.
type Button struct {
	native *C.GtkButton
	// Private : bin
	// Private : priv
}

func ButtonNewFromC(u unsafe.Pointer) *Button {
	c := (*C.GtkButton)(u)
	if c == nil {
		return nil
	}

	g := &Button{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Button) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Button) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonNew is a wrapper around the C function gtk_button_new.
func ButtonNew() *Button {
	retC := C.gtk_button_new()
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewFromStock is a wrapper around the C function gtk_button_new_from_stock.
func ButtonNewFromStock(stockId string) *Button {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_button_new_from_stock(c_stock_id)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithLabel is a wrapper around the C function gtk_button_new_with_label.
func ButtonNewWithLabel(label string) *Button {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_label(c_label)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithMnemonic is a wrapper around the C function gtk_button_new_with_mnemonic.
func ButtonNewWithMnemonic(label string) *Button {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_mnemonic(c_label)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonAccessible is a wrapper around the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native *C.GtkButtonAccessible
	// parent : record
	// priv : record
}

func ButtonAccessibleNewFromC(u unsafe.Pointer) *ButtonAccessible {
	c := (*C.GtkButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonBox is a wrapper around the C record GtkButtonBox.
type ButtonBox struct {
	native *C.GtkButtonBox
	// box : record
	// Private : priv
}

func ButtonBoxNewFromC(u unsafe.Pointer) *ButtonBox {
	c := (*C.GtkButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &ButtonBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ButtonBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Calendar is a wrapper around the C record GtkCalendar.
type Calendar struct {
	native *C.GtkCalendar
	// widget : record
	// priv : record
}

func CalendarNewFromC(u unsafe.Pointer) *Calendar {
	c := (*C.GtkCalendar)(u)
	if c == nil {
		return nil
	}

	g := &Calendar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Calendar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Calendar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CalendarNew is a wrapper around the C function gtk_calendar_new.
func CalendarNew() *Calendar {
	retC := C.gtk_calendar_new()
	retGo := CalendarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellAccessible is a wrapper around the C record GtkCellAccessible.
type CellAccessible struct {
	native *C.GtkCellAccessible
	// parent : record
	// priv : record
}

func CellAccessibleNewFromC(u unsafe.Pointer) *CellAccessible {
	c := (*C.GtkCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellArea is a wrapper around the C record GtkCellArea.
type CellArea struct {
	native *C.GtkCellArea
	// Private : parent_instance
	// Private : priv
}

func CellAreaNewFromC(u unsafe.Pointer) *CellArea {
	c := (*C.GtkCellArea)(u)
	if c == nil {
		return nil
	}

	g := &CellArea{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellArea) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaBox is a wrapper around the C record GtkCellAreaBox.
type CellAreaBox struct {
	native *C.GtkCellAreaBox
	// Private : parent_instance
	// Private : priv
}

func CellAreaBoxNewFromC(u unsafe.Pointer) *CellAreaBox {
	c := (*C.GtkCellAreaBox)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellAreaBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellAreaBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAreaContext is a wrapper around the C record GtkCellAreaContext.
type CellAreaContext struct {
	native *C.GtkCellAreaContext
	// Private : parent_instance
	// Private : priv
}

func CellAreaContextNewFromC(u unsafe.Pointer) *CellAreaContext {
	c := (*C.GtkCellAreaContext)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellAreaContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellAreaContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer is a wrapper around the C record GtkCellRenderer.
type CellRenderer struct {
	native *C.GtkCellRenderer
	// parent_instance : record
	// Private : priv
}

func CellRendererNewFromC(u unsafe.Pointer) *CellRenderer {
	c := (*C.GtkCellRenderer)(u)
	if c == nil {
		return nil
	}

	g := &CellRenderer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRenderer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRenderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererAccel is a wrapper around the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native *C.GtkCellRendererAccel
	// parent : record
	// Private : priv
}

func CellRendererAccelNewFromC(u unsafe.Pointer) *CellRendererAccel {
	c := (*C.GtkCellRendererAccel)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererAccel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererAccel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererAccel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererCombo is a wrapper around the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native *C.GtkCellRendererCombo
	// parent : record
	// Private : priv
}

func CellRendererComboNewFromC(u unsafe.Pointer) *CellRendererCombo {
	c := (*C.GtkCellRendererCombo)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererCombo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererCombo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererCombo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbuf is a wrapper around the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native *C.GtkCellRendererPixbuf
	// parent : record
	// Private : priv
}

func CellRendererPixbufNewFromC(u unsafe.Pointer) *CellRendererPixbuf {
	c := (*C.GtkCellRendererPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererPixbuf{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererPixbuf) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbufNew is a wrapper around the C function gtk_cell_renderer_pixbuf_new.
func CellRendererPixbufNew() *CellRendererPixbuf {
	retC := C.gtk_cell_renderer_pixbuf_new()
	retGo := CellRendererPixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererProgress is a wrapper around the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native *C.GtkCellRendererProgress
	// parent_instance : record
	// Private : priv
}

func CellRendererProgressNewFromC(u unsafe.Pointer) *CellRendererProgress {
	c := (*C.GtkCellRendererProgress)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererProgress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererProgress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererProgress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpin is a wrapper around the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native *C.GtkCellRendererSpin
	// parent : record
	// Private : priv
}

func CellRendererSpinNewFromC(u unsafe.Pointer) *CellRendererSpin {
	c := (*C.GtkCellRendererSpin)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpin{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererSpin) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererSpin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererSpinner is a wrapper around the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native *C.GtkCellRendererSpinner
	// parent : record
	// Private : priv
}

func CellRendererSpinnerNewFromC(u unsafe.Pointer) *CellRendererSpinner {
	c := (*C.GtkCellRendererSpinner)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinner{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererSpinner) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererSpinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererText is a wrapper around the C record GtkCellRendererText.
type CellRendererText struct {
	native *C.GtkCellRendererText
	// parent : record
	// Private : priv
}

func CellRendererTextNewFromC(u unsafe.Pointer) *CellRendererText {
	c := (*C.GtkCellRendererText)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererText{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererText) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererTextNew is a wrapper around the C function gtk_cell_renderer_text_new.
func CellRendererTextNew() *CellRendererText {
	retC := C.gtk_cell_renderer_text_new()
	retGo := CellRendererTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererToggle is a wrapper around the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native *C.GtkCellRendererToggle
	// parent : record
	// Private : priv
}

func CellRendererToggleNewFromC(u unsafe.Pointer) *CellRendererToggle {
	c := (*C.GtkCellRendererToggle)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererToggle{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellRendererToggle) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellRendererToggle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererToggleNew is a wrapper around the C function gtk_cell_renderer_toggle_new.
func CellRendererToggleNew() *CellRendererToggle {
	retC := C.gtk_cell_renderer_toggle_new()
	retGo := CellRendererToggleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellView is a wrapper around the C record GtkCellView.
type CellView struct {
	native *C.GtkCellView
	// parent_instance : record
	// Private : priv
}

func CellViewNewFromC(u unsafe.Pointer) *CellView {
	c := (*C.GtkCellView)(u)
	if c == nil {
		return nil
	}

	g := &CellView{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CellView) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CellView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckButton is a wrapper around the C record GtkCheckButton.
type CheckButton struct {
	native *C.GtkCheckButton
	// toggle_button : record
}

func CheckButtonNewFromC(u unsafe.Pointer) *CheckButton {
	c := (*C.GtkCheckButton)(u)
	if c == nil {
		return nil
	}

	g := &CheckButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CheckButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CheckButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckButtonNew is a wrapper around the C function gtk_check_button_new.
func CheckButtonNew() *CheckButton {
	retC := C.gtk_check_button_new()
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithLabel is a wrapper around the C function gtk_check_button_new_with_label.
func CheckButtonNewWithLabel(label string) *CheckButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_label(c_label)
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithMnemonic is a wrapper around the C function gtk_check_button_new_with_mnemonic.
func CheckButtonNewWithMnemonic(label string) *CheckButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_mnemonic(c_label)
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItem is a wrapper around the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native *C.GtkCheckMenuItem
	// menu_item : record
	// Private : priv
}

func CheckMenuItemNewFromC(u unsafe.Pointer) *CheckMenuItem {
	c := (*C.GtkCheckMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CheckMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CheckMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemNew is a wrapper around the C function gtk_check_menu_item_new.
func CheckMenuItemNew() *CheckMenuItem {
	retC := C.gtk_check_menu_item_new()
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithLabel is a wrapper around the C function gtk_check_menu_item_new_with_label.
func CheckMenuItemNewWithLabel(label string) *CheckMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_label(c_label)
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithMnemonic is a wrapper around the C function gtk_check_menu_item_new_with_mnemonic.
func CheckMenuItemNewWithMnemonic(label string) *CheckMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_mnemonic(c_label)
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemAccessible is a wrapper around the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native *C.GtkCheckMenuItemAccessible
	// parent : record
	// priv : record
}

func CheckMenuItemAccessibleNewFromC(u unsafe.Pointer) *CheckMenuItemAccessible {
	c := (*C.GtkCheckMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CheckMenuItemAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CheckMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Clipboard is a wrapper around the C record GtkClipboard.
type Clipboard struct {
	native *C.GtkClipboard
}

func ClipboardNewFromC(u unsafe.Pointer) *Clipboard {
	c := (*C.GtkClipboard)(u)
	if c == nil {
		return nil
	}

	g := &Clipboard{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Clipboard) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Clipboard) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorButton is a wrapper around the C record GtkColorButton.
type ColorButton struct {
	native *C.GtkColorButton
	// button : record
	// Private : priv
}

func ColorButtonNewFromC(u unsafe.Pointer) *ColorButton {
	c := (*C.GtkColorButton)(u)
	if c == nil {
		return nil
	}

	g := &ColorButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserDialog is a wrapper around the C record GtkColorChooserDialog.
type ColorChooserDialog struct {
	native *C.GtkColorChooserDialog
	// parent_instance : record
	// Private : priv
}

func ColorChooserDialogNewFromC(u unsafe.Pointer) *ColorChooserDialog {
	c := (*C.GtkColorChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorChooserDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorChooserWidget is a wrapper around the C record GtkColorChooserWidget.
type ColorChooserWidget struct {
	native *C.GtkColorChooserWidget
	// parent_instance : record
	// Private : priv
}

func ColorChooserWidgetNewFromC(u unsafe.Pointer) *ColorChooserWidget {
	c := (*C.GtkColorChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelection is a wrapper around the C record GtkColorSelection.
type ColorSelection struct {
	native *C.GtkColorSelection
	// parent_instance : record
	// Private : private_data
}

func ColorSelectionNewFromC(u unsafe.Pointer) *ColorSelection {
	c := (*C.GtkColorSelection)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorSelection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionNew is a wrapper around the C function gtk_color_selection_new.
func ColorSelectionNew() *ColorSelection {
	retC := C.gtk_color_selection_new()
	retGo := ColorSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorSelectionDialog is a wrapper around the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native *C.GtkColorSelectionDialog
	// parent_instance : record
	// Private : priv
}

func ColorSelectionDialogNewFromC(u unsafe.Pointer) *ColorSelectionDialog {
	c := (*C.GtkColorSelectionDialog)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorSelectionDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionDialogNew is a wrapper around the C function gtk_color_selection_dialog_new.
func ColorSelectionDialogNew(title string) *ColorSelectionDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_color_selection_dialog_new(c_title)
	retGo := ColorSelectionDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBox is a wrapper around the C record GtkComboBox.
type ComboBox struct {
	native *C.GtkComboBox
	// parent_instance : record
	// Private : priv
}

func ComboBoxNewFromC(u unsafe.Pointer) *ComboBox {
	c := (*C.GtkComboBox)(u)
	if c == nil {
		return nil
	}

	g := &ComboBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ComboBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ComboBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxNewWithArea is a wrapper around the C function gtk_combo_box_new_with_area.
func ComboBoxNewWithArea(area *CellArea) *ComboBox {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_combo_box_new_with_area(c_area)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithAreaAndEntry is a wrapper around the C function gtk_combo_box_new_with_area_and_entry.
func ComboBoxNewWithAreaAndEntry(area *CellArea) *ComboBox {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_combo_box_new_with_area_and_entry(c_area)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxAccessible is a wrapper around the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native *C.GtkComboBoxAccessible
	// parent : record
	// priv : record
}

func ComboBoxAccessibleNewFromC(u unsafe.Pointer) *ComboBoxAccessible {
	c := (*C.GtkComboBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ComboBoxAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ComboBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxText is a wrapper around the C record GtkComboBoxText.
type ComboBoxText struct {
	native *C.GtkComboBoxText
	// Private : parent_instance
	// Private : priv
}

func ComboBoxTextNewFromC(u unsafe.Pointer) *ComboBoxText {
	c := (*C.GtkComboBoxText)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxText{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ComboBoxText) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ComboBoxText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container is a wrapper around the C record GtkContainer.
type Container struct {
	native *C.GtkContainer
	// widget : record
	// Private : priv
}

func ContainerNewFromC(u unsafe.Pointer) *Container {
	c := (*C.GtkContainer)(u)
	if c == nil {
		return nil
	}

	g := &Container{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Container) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Container) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible is a wrapper around the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native *C.GtkContainerAccessible
	// parent : record
	// priv : record
}

func ContainerAccessibleNewFromC(u unsafe.Pointer) *ContainerAccessible {
	c := (*C.GtkContainerAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ContainerAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContainerAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContainerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerCellAccessible is a wrapper around the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native *C.GtkContainerCellAccessible
	// parent : record
	// priv : record
}

func ContainerCellAccessibleNewFromC(u unsafe.Pointer) *ContainerCellAccessible {
	c := (*C.GtkContainerCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ContainerCellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContainerCellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerCellAccessibleNew is a wrapper around the C function gtk_container_cell_accessible_new.
func ContainerCellAccessibleNew() *ContainerCellAccessible {
	retC := C.gtk_container_cell_accessible_new()
	retGo := ContainerCellAccessibleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CssProvider is a wrapper around the C record GtkCssProvider.
type CssProvider struct {
	native *C.GtkCssProvider
	// parent_instance : record
	// priv : record
}

func CssProviderNewFromC(u unsafe.Pointer) *CssProvider {
	c := (*C.GtkCssProvider)(u)
	if c == nil {
		return nil
	}

	g := &CssProvider{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CssProvider) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CssProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CssProviderNew is a wrapper around the C function gtk_css_provider_new.
func CssProviderNew() *CssProvider {
	retC := C.gtk_css_provider_new()
	retGo := CssProviderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Dialog is a wrapper around the C record GtkDialog.
type Dialog struct {
	native *C.GtkDialog
	// window : record
	// Private : priv
}

func DialogNewFromC(u unsafe.Pointer) *Dialog {
	c := (*C.GtkDialog)(u)
	if c == nil {
		return nil
	}

	g := &Dialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Dialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Dialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DialogNew is a wrapper around the C function gtk_dialog_new.
func DialogNew() *Dialog {
	retC := C.gtk_dialog_new()
	retGo := DialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// DrawingArea is a wrapper around the C record GtkDrawingArea.
type DrawingArea struct {
	native *C.GtkDrawingArea
	// widget : record
	// Private : dummy
}

func DrawingAreaNewFromC(u unsafe.Pointer) *DrawingArea {
	c := (*C.GtkDrawingArea)(u)
	if c == nil {
		return nil
	}

	g := &DrawingArea{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DrawingArea) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DrawingArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingAreaNew is a wrapper around the C function gtk_drawing_area_new.
func DrawingAreaNew() *DrawingArea {
	retC := C.gtk_drawing_area_new()
	retGo := DrawingAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Entry is a wrapper around the C record GtkEntry.
type Entry struct {
	native *C.GtkEntry
	// Private : parent_instance
	// Private : priv
}

func EntryNewFromC(u unsafe.Pointer) *Entry {
	c := (*C.GtkEntry)(u)
	if c == nil {
		return nil
	}

	g := &Entry{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Entry) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Entry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryNew is a wrapper around the C function gtk_entry_new.
func EntryNew() *Entry {
	retC := C.gtk_entry_new()
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EntryAccessible is a wrapper around the C record GtkEntryAccessible.
type EntryAccessible struct {
	native *C.GtkEntryAccessible
	// parent : record
	// priv : record
}

func EntryAccessibleNewFromC(u unsafe.Pointer) *EntryAccessible {
	c := (*C.GtkEntryAccessible)(u)
	if c == nil {
		return nil
	}

	g := &EntryAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EntryAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EntryAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryBuffer is a wrapper around the C record GtkEntryBuffer.
type EntryBuffer struct {
	native *C.GtkEntryBuffer
	// parent_instance : record
	// Private : priv
}

func EntryBufferNewFromC(u unsafe.Pointer) *EntryBuffer {
	c := (*C.GtkEntryBuffer)(u)
	if c == nil {
		return nil
	}

	g := &EntryBuffer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EntryBuffer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EntryBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryCompletion is a wrapper around the C record GtkEntryCompletion.
type EntryCompletion struct {
	native *C.GtkEntryCompletion
	// parent_instance : record
	// Private : priv
}

func EntryCompletionNewFromC(u unsafe.Pointer) *EntryCompletion {
	c := (*C.GtkEntryCompletion)(u)
	if c == nil {
		return nil
	}

	g := &EntryCompletion{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EntryCompletion) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EntryCompletion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : EntryIconAccessible : no CType

// EventBox is a wrapper around the C record GtkEventBox.
type EventBox struct {
	native *C.GtkEventBox
	// bin : record
	// Private : priv
}

func EventBoxNewFromC(u unsafe.Pointer) *EventBox {
	c := (*C.GtkEventBox)(u)
	if c == nil {
		return nil
	}

	g := &EventBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EventBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EventBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventBoxNew is a wrapper around the C function gtk_event_box_new.
func EventBoxNew() *EventBox {
	retC := C.gtk_event_box_new()
	retGo := EventBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EventController is a wrapper around the C record GtkEventController.
type EventController struct {
	native *C.GtkEventController
}

func EventControllerNewFromC(u unsafe.Pointer) *EventController {
	c := (*C.GtkEventController)(u)
	if c == nil {
		return nil
	}

	g := &EventController{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EventController) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EventController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Expander is a wrapper around the C record GtkExpander.
type Expander struct {
	native *C.GtkExpander
	// bin : record
	// priv : record
}

func ExpanderNewFromC(u unsafe.Pointer) *Expander {
	c := (*C.GtkExpander)(u)
	if c == nil {
		return nil
	}

	g := &Expander{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Expander) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Expander) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ExpanderAccessible is a wrapper around the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native *C.GtkExpanderAccessible
	// parent : record
	// priv : record
}

func ExpanderAccessibleNewFromC(u unsafe.Pointer) *ExpanderAccessible {
	c := (*C.GtkExpanderAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ExpanderAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ExpanderAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserButton is a wrapper around the C record GtkFileChooserButton.
type FileChooserButton struct {
	native *C.GtkFileChooserButton
	// parent : record
	// Private : priv
}

func FileChooserButtonNewFromC(u unsafe.Pointer) *FileChooserButton {
	c := (*C.GtkFileChooserButton)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileChooserButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserDialog is a wrapper around the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native *C.GtkFileChooserDialog
	// parent_instance : record
	// priv : record
}

func FileChooserDialogNewFromC(u unsafe.Pointer) *FileChooserDialog {
	c := (*C.GtkFileChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileChooserDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileChooserWidget is a wrapper around the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native *C.GtkFileChooserWidget
	// parent_instance : record
	// priv : record
}

func FileChooserWidgetNewFromC(u unsafe.Pointer) *FileChooserWidget {
	c := (*C.GtkFileChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileFilter is a wrapper around the C record GtkFileFilter.
type FileFilter struct {
	native *C.GtkFileFilter
}

func FileFilterNewFromC(u unsafe.Pointer) *FileFilter {
	c := (*C.GtkFileFilter)(u)
	if c == nil {
		return nil
	}

	g := &FileFilter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileFilter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Fixed is a wrapper around the C record GtkFixed.
type Fixed struct {
	native *C.GtkFixed
	// container : record
	// Private : priv
}

func FixedNewFromC(u unsafe.Pointer) *Fixed {
	c := (*C.GtkFixed)(u)
	if c == nil {
		return nil
	}

	g := &Fixed{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Fixed) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Fixed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FixedNew is a wrapper around the C function gtk_fixed_new.
func FixedNew() *Fixed {
	retC := C.gtk_fixed_new()
	retGo := FixedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FlowBox is a wrapper around the C record GtkFlowBox.
type FlowBox struct {
	native *C.GtkFlowBox
	// container : record
}

func FlowBoxNewFromC(u unsafe.Pointer) *FlowBox {
	c := (*C.GtkFlowBox)(u)
	if c == nil {
		return nil
	}

	g := &FlowBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FlowBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FlowBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxAccessible is a wrapper around the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native *C.GtkFlowBoxAccessible
	// parent : record
	// priv : record
}

func FlowBoxAccessibleNewFromC(u unsafe.Pointer) *FlowBoxAccessible {
	c := (*C.GtkFlowBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FlowBoxAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FlowBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxChild is a wrapper around the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native *C.GtkFlowBoxChild
	// parent_instance : record
}

func FlowBoxChildNewFromC(u unsafe.Pointer) *FlowBoxChild {
	c := (*C.GtkFlowBoxChild)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChild{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FlowBoxChild) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FlowBoxChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlowBoxChildAccessible is a wrapper around the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native *C.GtkFlowBoxChildAccessible
	// parent : record
}

func FlowBoxChildAccessibleNewFromC(u unsafe.Pointer) *FlowBoxChildAccessible {
	c := (*C.GtkFlowBoxChildAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChildAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FlowBoxChildAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FlowBoxChildAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontButton is a wrapper around the C record GtkFontButton.
type FontButton struct {
	native *C.GtkFontButton
	// button : record
	// Private : priv
}

func FontButtonNewFromC(u unsafe.Pointer) *FontButton {
	c := (*C.GtkFontButton)(u)
	if c == nil {
		return nil
	}

	g := &FontButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserDialog is a wrapper around the C record GtkFontChooserDialog.
type FontChooserDialog struct {
	native *C.GtkFontChooserDialog
	// parent_instance : record
	// Private : priv
}

func FontChooserDialogNewFromC(u unsafe.Pointer) *FontChooserDialog {
	c := (*C.GtkFontChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontChooserDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontChooserWidget is a wrapper around the C record GtkFontChooserWidget.
type FontChooserWidget struct {
	native *C.GtkFontChooserWidget
	// parent_instance : record
	// Private : priv
}

func FontChooserWidgetNewFromC(u unsafe.Pointer) *FontChooserWidget {
	c := (*C.GtkFontChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelection is a wrapper around the C record GtkFontSelection.
type FontSelection struct {
	native *C.GtkFontSelection
	// parent_instance : record
	// Private : priv
}

func FontSelectionNewFromC(u unsafe.Pointer) *FontSelection {
	c := (*C.GtkFontSelection)(u)
	if c == nil {
		return nil
	}

	g := &FontSelection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontSelection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionNew is a wrapper around the C function gtk_font_selection_new.
func FontSelectionNew() *FontSelection {
	retC := C.gtk_font_selection_new()
	retGo := FontSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontSelectionDialog is a wrapper around the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native *C.GtkFontSelectionDialog
	// parent_instance : record
	// Private : priv
}

func FontSelectionDialogNewFromC(u unsafe.Pointer) *FontSelectionDialog {
	c := (*C.GtkFontSelectionDialog)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontSelectionDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionDialogNew is a wrapper around the C function gtk_font_selection_dialog_new.
func FontSelectionDialogNew(title string) *FontSelectionDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_font_selection_dialog_new(c_title)
	retGo := FontSelectionDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frame is a wrapper around the C record GtkFrame.
type Frame struct {
	native *C.GtkFrame
	// bin : record
	// Private : priv
}

func FrameNewFromC(u unsafe.Pointer) *Frame {
	c := (*C.GtkFrame)(u)
	if c == nil {
		return nil
	}

	g := &Frame{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Frame) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Frame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameNew is a wrapper around the C function gtk_frame_new.
func FrameNew(label string) *Frame {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_frame_new(c_label)
	retGo := FrameNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FrameAccessible is a wrapper around the C record GtkFrameAccessible.
type FrameAccessible struct {
	native *C.GtkFrameAccessible
	// parent : record
	// priv : record
}

func FrameAccessibleNewFromC(u unsafe.Pointer) *FrameAccessible {
	c := (*C.GtkFrameAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FrameAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FrameAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FrameAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gesture is a wrapper around the C record GtkGesture.
type Gesture struct {
	native *C.GtkGesture
}

func GestureNewFromC(u unsafe.Pointer) *Gesture {
	c := (*C.GtkGesture)(u)
	if c == nil {
		return nil
	}

	g := &Gesture{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Gesture) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Gesture) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureDrag is a wrapper around the C record GtkGestureDrag.
type GestureDrag struct {
	native *C.GtkGestureDrag
}

func GestureDragNewFromC(u unsafe.Pointer) *GestureDrag {
	c := (*C.GtkGestureDrag)(u)
	if c == nil {
		return nil
	}

	g := &GestureDrag{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureDrag) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureDrag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureLongPress is a wrapper around the C record GtkGestureLongPress.
type GestureLongPress struct {
	native *C.GtkGestureLongPress
}

func GestureLongPressNewFromC(u unsafe.Pointer) *GestureLongPress {
	c := (*C.GtkGestureLongPress)(u)
	if c == nil {
		return nil
	}

	g := &GestureLongPress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureLongPress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureLongPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureMultiPress is a wrapper around the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native *C.GtkGestureMultiPress
}

func GestureMultiPressNewFromC(u unsafe.Pointer) *GestureMultiPress {
	c := (*C.GtkGestureMultiPress)(u)
	if c == nil {
		return nil
	}

	g := &GestureMultiPress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureMultiPress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureMultiPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GesturePan is a wrapper around the C record GtkGesturePan.
type GesturePan struct {
	native *C.GtkGesturePan
}

func GesturePanNewFromC(u unsafe.Pointer) *GesturePan {
	c := (*C.GtkGesturePan)(u)
	if c == nil {
		return nil
	}

	g := &GesturePan{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GesturePan) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GesturePan) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureRotate is a wrapper around the C record GtkGestureRotate.
type GestureRotate struct {
	native *C.GtkGestureRotate
}

func GestureRotateNewFromC(u unsafe.Pointer) *GestureRotate {
	c := (*C.GtkGestureRotate)(u)
	if c == nil {
		return nil
	}

	g := &GestureRotate{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureRotate) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureRotate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle is a wrapper around the C record GtkGestureSingle.
type GestureSingle struct {
	native *C.GtkGestureSingle
}

func GestureSingleNewFromC(u unsafe.Pointer) *GestureSingle {
	c := (*C.GtkGestureSingle)(u)
	if c == nil {
		return nil
	}

	g := &GestureSingle{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureSingle) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureSingle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSwipe is a wrapper around the C record GtkGestureSwipe.
type GestureSwipe struct {
	native *C.GtkGestureSwipe
}

func GestureSwipeNewFromC(u unsafe.Pointer) *GestureSwipe {
	c := (*C.GtkGestureSwipe)(u)
	if c == nil {
		return nil
	}

	g := &GestureSwipe{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureSwipe) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureSwipe) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureZoom is a wrapper around the C record GtkGestureZoom.
type GestureZoom struct {
	native *C.GtkGestureZoom
}

func GestureZoomNewFromC(u unsafe.Pointer) *GestureZoom {
	c := (*C.GtkGestureZoom)(u)
	if c == nil {
		return nil
	}

	g := &GestureZoom{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GestureZoom) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GestureZoom) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Grid is a wrapper around the C record GtkGrid.
type Grid struct {
	native *C.GtkGrid
	// Private : container
	// Private : priv
}

func GridNewFromC(u unsafe.Pointer) *Grid {
	c := (*C.GtkGrid)(u)
	if c == nil {
		return nil
	}

	g := &Grid{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Grid) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Grid) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GridNew is a wrapper around the C function gtk_grid_new.
func GridNew() *Grid {
	retC := C.gtk_grid_new()
	retGo := GridNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HBox is a wrapper around the C record GtkHBox.
type HBox struct {
	native *C.GtkHBox
	// box : record
}

func HBoxNewFromC(u unsafe.Pointer) *HBox {
	c := (*C.GtkHBox)(u)
	if c == nil {
		return nil
	}

	g := &HBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HBoxNew is a wrapper around the C function gtk_hbox_new.
func HBoxNew(homogeneous bool, spacing int32) *HBox {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_hbox_new(c_homogeneous, c_spacing)
	retGo := HBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HButtonBox is a wrapper around the C record GtkHButtonBox.
type HButtonBox struct {
	native *C.GtkHButtonBox
	// button_box : record
}

func HButtonBoxNewFromC(u unsafe.Pointer) *HButtonBox {
	c := (*C.GtkHButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &HButtonBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HButtonBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HButtonBoxNew is a wrapper around the C function gtk_hbutton_box_new.
func HButtonBoxNew() *HButtonBox {
	retC := C.gtk_hbutton_box_new()
	retGo := HButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HPaned is a wrapper around the C record GtkHPaned.
type HPaned struct {
	native *C.GtkHPaned
	// paned : record
}

func HPanedNewFromC(u unsafe.Pointer) *HPaned {
	c := (*C.GtkHPaned)(u)
	if c == nil {
		return nil
	}

	g := &HPaned{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HPaned) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HPanedNew is a wrapper around the C function gtk_hpaned_new.
func HPanedNew() *HPaned {
	retC := C.gtk_hpaned_new()
	retGo := HPanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSV is a wrapper around the C record GtkHSV.
type HSV struct {
	native *C.GtkHSV
	// parent_instance : record
	// Private : priv
}

func HSVNewFromC(u unsafe.Pointer) *HSV {
	c := (*C.GtkHSV)(u)
	if c == nil {
		return nil
	}

	g := &HSV{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HSV) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HSV) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScale is a wrapper around the C record GtkHScale.
type HScale struct {
	native *C.GtkHScale
	// scale : record
}

func HScaleNewFromC(u unsafe.Pointer) *HScale {
	c := (*C.GtkHScale)(u)
	if c == nil {
		return nil
	}

	g := &HScale{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HScale) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScaleNew is a wrapper around the C function gtk_hscale_new.
func HScaleNew(adjustment *Adjustment) *HScale {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_hscale_new(c_adjustment)
	retGo := HScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HScaleNewWithRange is a wrapper around the C function gtk_hscale_new_with_range.
func HScaleNewWithRange(min float64, max float64, step float64) *HScale {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_hscale_new_with_range(c_min, c_max, c_step)
	retGo := HScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HScrollbar is a wrapper around the C record GtkHScrollbar.
type HScrollbar struct {
	native *C.GtkHScrollbar
	// scrollbar : record
}

func HScrollbarNewFromC(u unsafe.Pointer) *HScrollbar {
	c := (*C.GtkHScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &HScrollbar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HScrollbar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScrollbarNew is a wrapper around the C function gtk_hscrollbar_new.
func HScrollbarNew(adjustment *Adjustment) *HScrollbar {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_hscrollbar_new(c_adjustment)
	retGo := HScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSeparator is a wrapper around the C record GtkHSeparator.
type HSeparator struct {
	native *C.GtkHSeparator
	// separator : record
}

func HSeparatorNewFromC(u unsafe.Pointer) *HSeparator {
	c := (*C.GtkHSeparator)(u)
	if c == nil {
		return nil
	}

	g := &HSeparator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HSeparator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HSeparatorNew is a wrapper around the C function gtk_hseparator_new.
func HSeparatorNew() *HSeparator {
	retC := C.gtk_hseparator_new()
	retGo := HSeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HandleBox is a wrapper around the C record GtkHandleBox.
type HandleBox struct {
	native *C.GtkHandleBox
	// bin : record
	// Private : priv
}

func HandleBoxNewFromC(u unsafe.Pointer) *HandleBox {
	c := (*C.GtkHandleBox)(u)
	if c == nil {
		return nil
	}

	g := &HandleBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HandleBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HandleBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HandleBoxNew is a wrapper around the C function gtk_handle_box_new.
func HandleBoxNew() *HandleBox {
	retC := C.gtk_handle_box_new()
	retGo := HandleBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HeaderBar is a wrapper around the C record GtkHeaderBar.
type HeaderBar struct {
	native *C.GtkHeaderBar
	// container : record
}

func HeaderBarNewFromC(u unsafe.Pointer) *HeaderBar {
	c := (*C.GtkHeaderBar)(u)
	if c == nil {
		return nil
	}

	g := &HeaderBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HeaderBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HeaderBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContext is a wrapper around the C record GtkIMContext.
type IMContext struct {
	native *C.GtkIMContext
	// parent_instance : record
}

func IMContextNewFromC(u unsafe.Pointer) *IMContext {
	c := (*C.GtkIMContext)(u)
	if c == nil {
		return nil
	}

	g := &IMContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IMContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IMContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimple is a wrapper around the C record GtkIMContextSimple.
type IMContextSimple struct {
	native *C.GtkIMContextSimple
	// object : record
	// Private : priv
}

func IMContextSimpleNewFromC(u unsafe.Pointer) *IMContextSimple {
	c := (*C.GtkIMContextSimple)(u)
	if c == nil {
		return nil
	}

	g := &IMContextSimple{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IMContextSimple) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IMContextSimple) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimpleNew is a wrapper around the C function gtk_im_context_simple_new.
func IMContextSimpleNew() *IMContextSimple {
	retC := C.gtk_im_context_simple_new()
	retGo := IMContextSimpleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IMMulticontext is a wrapper around the C record GtkIMMulticontext.
type IMMulticontext struct {
	native *C.GtkIMMulticontext
	// object : record
	// Private : priv
}

func IMMulticontextNewFromC(u unsafe.Pointer) *IMMulticontext {
	c := (*C.GtkIMMulticontext)(u)
	if c == nil {
		return nil
	}

	g := &IMMulticontext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IMMulticontext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IMMulticontext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMMulticontextNew is a wrapper around the C function gtk_im_multicontext_new.
func IMMulticontextNew() *IMMulticontext {
	retC := C.gtk_im_multicontext_new()
	retGo := IMMulticontextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IconFactory is a wrapper around the C record GtkIconFactory.
type IconFactory struct {
	native *C.GtkIconFactory
	// parent_instance : record
	// Private : priv
}

func IconFactoryNewFromC(u unsafe.Pointer) *IconFactory {
	c := (*C.GtkIconFactory)(u)
	if c == nil {
		return nil
	}

	g := &IconFactory{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IconFactory) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IconFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconFactoryNew is a wrapper around the C function gtk_icon_factory_new.
func IconFactoryNew() *IconFactory {
	retC := C.gtk_icon_factory_new()
	retGo := IconFactoryNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IconInfo is a wrapper around the C record GtkIconInfo.
type IconInfo struct {
	native *C.GtkIconInfo
}

func IconInfoNewFromC(u unsafe.Pointer) *IconInfo {
	c := (*C.GtkIconInfo)(u)
	if c == nil {
		return nil
	}

	g := &IconInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IconInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IconInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconTheme is a wrapper around the C record GtkIconTheme.
type IconTheme struct {
	native *C.GtkIconTheme
	// Private : parent_instance
	// Private : priv
}

func IconThemeNewFromC(u unsafe.Pointer) *IconTheme {
	c := (*C.GtkIconTheme)(u)
	if c == nil {
		return nil
	}

	g := &IconTheme{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IconTheme) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IconTheme) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconView is a wrapper around the C record GtkIconView.
type IconView struct {
	native *C.GtkIconView
	// parent : record
	// Private : priv
}

func IconViewNewFromC(u unsafe.Pointer) *IconView {
	c := (*C.GtkIconView)(u)
	if c == nil {
		return nil
	}

	g := &IconView{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IconView) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IconView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconViewAccessible is a wrapper around the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native *C.GtkIconViewAccessible
	// parent : record
	// priv : record
}

func IconViewAccessibleNewFromC(u unsafe.Pointer) *IconViewAccessible {
	c := (*C.GtkIconViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &IconViewAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IconViewAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IconViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Image is a wrapper around the C record GtkImage.
type Image struct {
	native *C.GtkImage
	// misc : record
	// Private : priv
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	c := (*C.GtkImage)(u)
	if c == nil {
		return nil
	}

	g := &Image{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Image) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageNew is a wrapper around the C function gtk_image_new.
func ImageNew() *Image {
	retC := C.gtk_image_new()
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromAnimation is a wrapper around the C function gtk_image_new_from_animation.
func ImageNewFromAnimation(animation *gdkpixbuf.PixbufAnimation) *Image {
	c_animation := (*C.GdkPixbufAnimation)(C.NULL)
	if animation != nil {
		c_animation = (*C.GdkPixbufAnimation)(animation.ToC())
	}

	retC := C.gtk_image_new_from_animation(c_animation)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromFile is a wrapper around the C function gtk_image_new_from_file.
func ImageNewFromFile(filename string) *Image {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_image_new_from_file(c_filename)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromIconSet is a wrapper around the C function gtk_image_new_from_icon_set.
func ImageNewFromIconSet(iconSet *IconSet, size IconSize) *Image {
	c_icon_set := (*C.GtkIconSet)(C.NULL)
	if iconSet != nil {
		c_icon_set = (*C.GtkIconSet)(iconSet.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_icon_set(c_icon_set, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromPixbuf is a wrapper around the C function gtk_image_new_from_pixbuf.
func ImageNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_image_new_from_pixbuf(c_pixbuf)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromStock is a wrapper around the C function gtk_image_new_from_stock.
func ImageNewFromStock(stockId string, size IconSize) *Image {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_stock(c_stock_id, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageAccessible is a wrapper around the C record GtkImageAccessible.
type ImageAccessible struct {
	native *C.GtkImageAccessible
	// parent : record
	// priv : record
}

func ImageAccessibleNewFromC(u unsafe.Pointer) *ImageAccessible {
	c := (*C.GtkImageAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ImageAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ImageAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ImageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageCellAccessible is a wrapper around the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native *C.GtkImageCellAccessible
	// parent : record
	// priv : record
}

func ImageCellAccessibleNewFromC(u unsafe.Pointer) *ImageCellAccessible {
	c := (*C.GtkImageCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ImageCellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ImageCellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ImageCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageMenuItem is a wrapper around the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native *C.GtkImageMenuItem
	// menu_item : record
	// Private : priv
}

func ImageMenuItemNewFromC(u unsafe.Pointer) *ImageMenuItem {
	c := (*C.GtkImageMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &ImageMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ImageMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ImageMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageMenuItemNew is a wrapper around the C function gtk_image_menu_item_new.
func ImageMenuItemNew() *ImageMenuItem {
	retC := C.gtk_image_menu_item_new()
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewFromStock is a wrapper around the C function gtk_image_menu_item_new_from_stock.
func ImageMenuItemNewFromStock(stockId string, accelGroup *AccelGroup) *ImageMenuItem {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	retC := C.gtk_image_menu_item_new_from_stock(c_stock_id, c_accel_group)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithLabel is a wrapper around the C function gtk_image_menu_item_new_with_label.
func ImageMenuItemNewWithLabel(label string) *ImageMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_label(c_label)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithMnemonic is a wrapper around the C function gtk_image_menu_item_new_with_mnemonic.
func ImageMenuItemNewWithMnemonic(label string) *ImageMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_mnemonic(c_label)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InfoBar is a wrapper around the C record GtkInfoBar.
type InfoBar struct {
	native *C.GtkInfoBar
	// parent : record
	// Private : priv
}

func InfoBarNewFromC(u unsafe.Pointer) *InfoBar {
	c := (*C.GtkInfoBar)(u)
	if c == nil {
		return nil
	}

	g := &InfoBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InfoBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InfoBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Invisible is a wrapper around the C record GtkInvisible.
type Invisible struct {
	native *C.GtkInvisible
	// widget : record
	// Private : priv
}

func InvisibleNewFromC(u unsafe.Pointer) *Invisible {
	c := (*C.GtkInvisible)(u)
	if c == nil {
		return nil
	}

	g := &Invisible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Invisible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Invisible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InvisibleNew is a wrapper around the C function gtk_invisible_new.
func InvisibleNew() *Invisible {
	retC := C.gtk_invisible_new()
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Label is a wrapper around the C record GtkLabel.
type Label struct {
	native *C.GtkLabel
	// misc : record
	// Private : priv
}

func LabelNewFromC(u unsafe.Pointer) *Label {
	c := (*C.GtkLabel)(u)
	if c == nil {
		return nil
	}

	g := &Label{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Label) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Label) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelNew is a wrapper around the C function gtk_label_new.
func LabelNew(str string) *Label {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new(c_str)
	retGo := LabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LabelNewWithMnemonic is a wrapper around the C function gtk_label_new_with_mnemonic.
func LabelNewWithMnemonic(str string) *Label {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new_with_mnemonic(c_str)
	retGo := LabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LabelAccessible is a wrapper around the C record GtkLabelAccessible.
type LabelAccessible struct {
	native *C.GtkLabelAccessible
	// parent : record
	// priv : record
}

func LabelAccessibleNewFromC(u unsafe.Pointer) *LabelAccessible {
	c := (*C.GtkLabelAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LabelAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LabelAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LabelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Layout is a wrapper around the C record GtkLayout.
type Layout struct {
	native *C.GtkLayout
	// container : record
	// Private : priv
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	c := (*C.GtkLayout)(u)
	if c == nil {
		return nil
	}

	g := &Layout{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Layout) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutNew is a wrapper around the C function gtk_layout_new.
func LayoutNew(hadjustment *Adjustment, vadjustment *Adjustment) *Layout {
	c_hadjustment := (*C.GtkAdjustment)(C.NULL)
	if hadjustment != nil {
		c_hadjustment = (*C.GtkAdjustment)(hadjustment.ToC())
	}

	c_vadjustment := (*C.GtkAdjustment)(C.NULL)
	if vadjustment != nil {
		c_vadjustment = (*C.GtkAdjustment)(vadjustment.ToC())
	}

	retC := C.gtk_layout_new(c_hadjustment, c_vadjustment)
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LevelBar is a wrapper around the C record GtkLevelBar.
type LevelBar struct {
	native *C.GtkLevelBar
	// Private : parent
	// Private : priv
}

func LevelBarNewFromC(u unsafe.Pointer) *LevelBar {
	c := (*C.GtkLevelBar)(u)
	if c == nil {
		return nil
	}

	g := &LevelBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LevelBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LevelBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LevelBarAccessible is a wrapper around the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native *C.GtkLevelBarAccessible
	// parent : record
	// priv : record
}

func LevelBarAccessibleNewFromC(u unsafe.Pointer) *LevelBarAccessible {
	c := (*C.GtkLevelBarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LevelBarAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LevelBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButton is a wrapper around the C record GtkLinkButton.
type LinkButton struct {
	native *C.GtkLinkButton
	// Private : parent_instance
	// Private : priv
}

func LinkButtonNewFromC(u unsafe.Pointer) *LinkButton {
	c := (*C.GtkLinkButton)(u)
	if c == nil {
		return nil
	}

	g := &LinkButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LinkButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LinkButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LinkButtonAccessible is a wrapper around the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native *C.GtkLinkButtonAccessible
	// parent : record
	// priv : record
}

func LinkButtonAccessibleNewFromC(u unsafe.Pointer) *LinkButtonAccessible {
	c := (*C.GtkLinkButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LinkButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LinkButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBox is a wrapper around the C record GtkListBox.
type ListBox struct {
	native *C.GtkListBox
	// parent_instance : record
}

func ListBoxNewFromC(u unsafe.Pointer) *ListBox {
	c := (*C.GtkListBox)(u)
	if c == nil {
		return nil
	}

	g := &ListBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxAccessible is a wrapper around the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native *C.GtkListBoxAccessible
	// parent : record
	// priv : record
}

func ListBoxAccessibleNewFromC(u unsafe.Pointer) *ListBoxAccessible {
	c := (*C.GtkListBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListBoxAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxRow is a wrapper around the C record GtkListBoxRow.
type ListBoxRow struct {
	native *C.GtkListBoxRow
	// parent_instance : record
}

func ListBoxRowNewFromC(u unsafe.Pointer) *ListBoxRow {
	c := (*C.GtkListBoxRow)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRow{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListBoxRow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListBoxRow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListBoxRowAccessible is a wrapper around the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native *C.GtkListBoxRowAccessible
	// parent : record
}

func ListBoxRowAccessibleNewFromC(u unsafe.Pointer) *ListBoxRowAccessible {
	c := (*C.GtkListBoxRowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRowAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListBoxRowAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListBoxRowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStore is a wrapper around the C record GtkListStore.
type ListStore struct {
	native *C.GtkListStore
	// parent : record
	// Private : priv
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GtkListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListStore) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// ListStoreNewv is a wrapper around the C function gtk_list_store_newv.
func ListStoreNewv(types []gobject.Type) *ListStore {
	c_n_columns := (C.gint)(len(types))

	c_types_array := make([]C.GType, len(types)+1, len(types)+1)
	for i, item := range types {
		c := (C.GType)(item)
		c_types_array[i] = c
	}
	c_types_array[len(types)] = 0
	c_types_arrayPtr := &c_types_array[0]
	c_types := (*C.GType)(unsafe.Pointer(c_types_arrayPtr))

	retC := C.gtk_list_store_newv(c_n_columns, c_types)
	retGo := ListStoreNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// LockButton is a wrapper around the C record GtkLockButton.
type LockButton struct {
	native *C.GtkLockButton
	// parent : record
	// priv : record
}

func LockButtonNewFromC(u unsafe.Pointer) *LockButton {
	c := (*C.GtkLockButton)(u)
	if c == nil {
		return nil
	}

	g := &LockButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LockButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LockButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LockButtonAccessible is a wrapper around the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native *C.GtkLockButtonAccessible
	// parent : record
	// priv : record
}

func LockButtonAccessibleNewFromC(u unsafe.Pointer) *LockButtonAccessible {
	c := (*C.GtkLockButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LockButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LockButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Menu is a wrapper around the C record GtkMenu.
type Menu struct {
	native *C.GtkMenu
	// menu_shell : record
	// Private : priv
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GtkMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Menu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuNew is a wrapper around the C function gtk_menu_new.
func MenuNew() *Menu {
	retC := C.gtk_menu_new()
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuAccessible is a wrapper around the C record GtkMenuAccessible.
type MenuAccessible struct {
	native *C.GtkMenuAccessible
	// parent : record
	// priv : record
}

func MenuAccessibleNewFromC(u unsafe.Pointer) *MenuAccessible {
	c := (*C.GtkMenuAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuBar is a wrapper around the C record GtkMenuBar.
type MenuBar struct {
	native *C.GtkMenuBar
	// menu_shell : record
	// Private : priv
}

func MenuBarNewFromC(u unsafe.Pointer) *MenuBar {
	c := (*C.GtkMenuBar)(u)
	if c == nil {
		return nil
	}

	g := &MenuBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuBarNew is a wrapper around the C function gtk_menu_bar_new.
func MenuBarNew() *MenuBar {
	retC := C.gtk_menu_bar_new()
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuButton is a wrapper around the C record GtkMenuButton.
type MenuButton struct {
	native *C.GtkMenuButton
	// parent : record
	// Private : priv
}

func MenuButtonNewFromC(u unsafe.Pointer) *MenuButton {
	c := (*C.GtkMenuButton)(u)
	if c == nil {
		return nil
	}

	g := &MenuButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuButtonAccessible is a wrapper around the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native *C.GtkMenuButtonAccessible
	// parent : record
	// priv : record
}

func MenuButtonAccessibleNewFromC(u unsafe.Pointer) *MenuButtonAccessible {
	c := (*C.GtkMenuButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem is a wrapper around the C record GtkMenuItem.
type MenuItem struct {
	native *C.GtkMenuItem
	// bin : record
	// Private : priv
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GtkMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemNew is a wrapper around the C function gtk_menu_item_new.
func MenuItemNew() *MenuItem {
	retC := C.gtk_menu_item_new()
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithLabel is a wrapper around the C function gtk_menu_item_new_with_label.
func MenuItemNewWithLabel(label string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_label(c_label)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithMnemonic is a wrapper around the C function gtk_menu_item_new_with_mnemonic.
func MenuItemNewWithMnemonic(label string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_mnemonic(c_label)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemAccessible is a wrapper around the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native *C.GtkMenuItemAccessible
	// parent : record
	// priv : record
}

func MenuItemAccessibleNewFromC(u unsafe.Pointer) *MenuItemAccessible {
	c := (*C.GtkMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuItemAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShell is a wrapper around the C record GtkMenuShell.
type MenuShell struct {
	native *C.GtkMenuShell
	// container : record
	// Private : priv
}

func MenuShellNewFromC(u unsafe.Pointer) *MenuShell {
	c := (*C.GtkMenuShell)(u)
	if c == nil {
		return nil
	}

	g := &MenuShell{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuShell) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellAccessible is a wrapper around the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native *C.GtkMenuShellAccessible
	// parent : record
	// priv : record
}

func MenuShellAccessibleNewFromC(u unsafe.Pointer) *MenuShellAccessible {
	c := (*C.GtkMenuShellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuShellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuShellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuToolButton is a wrapper around the C record GtkMenuToolButton.
type MenuToolButton struct {
	native *C.GtkMenuToolButton
	// parent : record
	// Private : priv
}

func MenuToolButtonNewFromC(u unsafe.Pointer) *MenuToolButton {
	c := (*C.GtkMenuToolButton)(u)
	if c == nil {
		return nil
	}

	g := &MenuToolButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuToolButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MessageDialog is a wrapper around the C record GtkMessageDialog.
type MessageDialog struct {
	native *C.GtkMessageDialog
	// parent_instance : record
	// Private : priv
}

func MessageDialogNewFromC(u unsafe.Pointer) *MessageDialog {
	c := (*C.GtkMessageDialog)(u)
	if c == nil {
		return nil
	}

	g := &MessageDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MessageDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MessageDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MessageDialogNew is a wrapper around the C function gtk_message_dialog_new.
func MessageDialogNew(parent *Window, flags DialogFlags, type_ MessageType, buttons ButtonsType, messageFormat string, args ...interface{}) *MessageDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_type := (C.GtkMessageType)(type_)

	c_buttons := (C.GtkButtonsType)(buttons)

	goFormattedString := fmt.Sprintf(messageFormat, args...)
	c_message_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_message_format))

	retC := C._gtk_message_dialog_new(c_parent, c_flags, c_type, c_buttons, c_message_format)
	retGo := MessageDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Misc is a wrapper around the C record GtkMisc.
type Misc struct {
	native *C.GtkMisc
	// widget : record
	// Private : priv
}

func MiscNewFromC(u unsafe.Pointer) *Misc {
	c := (*C.GtkMisc)(u)
	if c == nil {
		return nil
	}

	g := &Misc{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Misc) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ModelButton is a wrapper around the C record GtkModelButton.
type ModelButton struct {
	native *C.GtkModelButton
}

func ModelButtonNewFromC(u unsafe.Pointer) *ModelButton {
	c := (*C.GtkModelButton)(u)
	if c == nil {
		return nil
	}

	g := &ModelButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ModelButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ModelButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperation is a wrapper around the C record GtkMountOperation.
type MountOperation struct {
	native *C.GtkMountOperation
	// parent_instance : record
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	c := (*C.GtkMountOperation)(u)
	if c == nil {
		return nil
	}

	g := &MountOperation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MountOperation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Notebook is a wrapper around the C record GtkNotebook.
type Notebook struct {
	native *C.GtkNotebook
	// Private : container
	// Private : priv
}

func NotebookNewFromC(u unsafe.Pointer) *Notebook {
	c := (*C.GtkNotebook)(u)
	if c == nil {
		return nil
	}

	g := &Notebook{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Notebook) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Notebook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookNew is a wrapper around the C function gtk_notebook_new.
func NotebookNew() *Notebook {
	retC := C.gtk_notebook_new()
	retGo := NotebookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NotebookAccessible is a wrapper around the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native *C.GtkNotebookAccessible
	// parent : record
	// priv : record
}

func NotebookAccessibleNewFromC(u unsafe.Pointer) *NotebookAccessible {
	c := (*C.GtkNotebookAccessible)(u)
	if c == nil {
		return nil
	}

	g := &NotebookAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NotebookAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NotebookAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPageAccessible is a wrapper around the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native *C.GtkNotebookPageAccessible
	// parent : record
	// priv : record
}

func NotebookPageAccessibleNewFromC(u unsafe.Pointer) *NotebookPageAccessible {
	c := (*C.GtkNotebookPageAccessible)(u)
	if c == nil {
		return nil
	}

	g := &NotebookPageAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NotebookPageAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPageAccessibleNew is a wrapper around the C function gtk_notebook_page_accessible_new.
func NotebookPageAccessibleNew(notebook *NotebookAccessible, child *Widget) *NotebookPageAccessible {
	c_notebook := (*C.GtkNotebookAccessible)(C.NULL)
	if notebook != nil {
		c_notebook = (*C.GtkNotebookAccessible)(notebook.ToC())
	}

	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	retC := C.gtk_notebook_page_accessible_new(c_notebook, c_child)
	retGo := NotebookPageAccessibleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// NumerableIcon is a wrapper around the C record GtkNumerableIcon.
type NumerableIcon struct {
	native *C.GtkNumerableIcon
	// parent : record
	// Private : priv
}

func NumerableIconNewFromC(u unsafe.Pointer) *NumerableIcon {
	c := (*C.GtkNumerableIcon)(u)
	if c == nil {
		return nil
	}

	g := &NumerableIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NumerableIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NumerableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OffscreenWindow is a wrapper around the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native *C.GtkOffscreenWindow
	// parent_object : record
}

func OffscreenWindowNewFromC(u unsafe.Pointer) *OffscreenWindow {
	c := (*C.GtkOffscreenWindow)(u)
	if c == nil {
		return nil
	}

	g := &OffscreenWindow{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *OffscreenWindow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *OffscreenWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Overlay is a wrapper around the C record GtkOverlay.
type Overlay struct {
	native *C.GtkOverlay
	// parent : record
	// priv : record
}

func OverlayNewFromC(u unsafe.Pointer) *Overlay {
	c := (*C.GtkOverlay)(u)
	if c == nil {
		return nil
	}

	g := &Overlay{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Overlay) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Overlay) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PageSetup is a wrapper around the C record GtkPageSetup.
type PageSetup struct {
	native *C.GtkPageSetup
}

func PageSetupNewFromC(u unsafe.Pointer) *PageSetup {
	c := (*C.GtkPageSetup)(u)
	if c == nil {
		return nil
	}

	g := &PageSetup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PageSetup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PageSetup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Paned is a wrapper around the C record GtkPaned.
type Paned struct {
	native *C.GtkPaned
	// container : record
	// Private : priv
}

func PanedNewFromC(u unsafe.Pointer) *Paned {
	c := (*C.GtkPaned)(u)
	if c == nil {
		return nil
	}

	g := &Paned{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Paned) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Paned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PanedAccessible is a wrapper around the C record GtkPanedAccessible.
type PanedAccessible struct {
	native *C.GtkPanedAccessible
	// parent : record
	// priv : record
}

func PanedAccessibleNewFromC(u unsafe.Pointer) *PanedAccessible {
	c := (*C.GtkPanedAccessible)(u)
	if c == nil {
		return nil
	}

	g := &PanedAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PanedAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PanedAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PlacesSidebar is a wrapper around the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native *C.GtkPlacesSidebar
}

func PlacesSidebarNewFromC(u unsafe.Pointer) *PlacesSidebar {
	c := (*C.GtkPlacesSidebar)(u)
	if c == nil {
		return nil
	}

	g := &PlacesSidebar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PlacesSidebar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PlacesSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GtkPlug

// Popover is a wrapper around the C record GtkPopover.
type Popover struct {
	native *C.GtkPopover
	// parent_instance : record
	// Private : priv
}

func PopoverNewFromC(u unsafe.Pointer) *Popover {
	c := (*C.GtkPopover)(u)
	if c == nil {
		return nil
	}

	g := &Popover{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Popover) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Popover) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverAccessible is a wrapper around the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native *C.GtkPopoverAccessible
	// parent : record
}

func PopoverAccessibleNewFromC(u unsafe.Pointer) *PopoverAccessible {
	c := (*C.GtkPopoverAccessible)(u)
	if c == nil {
		return nil
	}

	g := &PopoverAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PopoverAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PopoverAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PopoverMenu is a wrapper around the C record GtkPopoverMenu.
type PopoverMenu struct {
	native *C.GtkPopoverMenu
}

func PopoverMenuNewFromC(u unsafe.Pointer) *PopoverMenu {
	c := (*C.GtkPopoverMenu)(u)
	if c == nil {
		return nil
	}

	g := &PopoverMenu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PopoverMenu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PopoverMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintContext is a wrapper around the C record GtkPrintContext.
type PrintContext struct {
	native *C.GtkPrintContext
}

func PrintContextNewFromC(u unsafe.Pointer) *PrintContext {
	c := (*C.GtkPrintContext)(u)
	if c == nil {
		return nil
	}

	g := &PrintContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintOperation is a wrapper around the C record GtkPrintOperation.
type PrintOperation struct {
	native *C.GtkPrintOperation
	// parent_instance : record
	// Private : priv
}

func PrintOperationNewFromC(u unsafe.Pointer) *PrintOperation {
	c := (*C.GtkPrintOperation)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintOperation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PrintSettings is a wrapper around the C record GtkPrintSettings.
type PrintSettings struct {
	native *C.GtkPrintSettings
}

func PrintSettingsNewFromC(u unsafe.Pointer) *PrintSettings {
	c := (*C.GtkPrintSettings)(u)
	if c == nil {
		return nil
	}

	g := &PrintSettings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintSettings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintSettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBar is a wrapper around the C record GtkProgressBar.
type ProgressBar struct {
	native *C.GtkProgressBar
	// parent : record
	// Private : priv
}

func ProgressBarNewFromC(u unsafe.Pointer) *ProgressBar {
	c := (*C.GtkProgressBar)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProgressBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProgressBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarNew is a wrapper around the C function gtk_progress_bar_new.
func ProgressBarNew() *ProgressBar {
	retC := C.gtk_progress_bar_new()
	retGo := ProgressBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ProgressBarAccessible is a wrapper around the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native *C.GtkProgressBarAccessible
	// parent : record
	// priv : record
}

func ProgressBarAccessibleNewFromC(u unsafe.Pointer) *ProgressBarAccessible {
	c := (*C.GtkProgressBarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProgressBarAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProgressBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioAction is a wrapper around the C record GtkRadioAction.
type RadioAction struct {
	native *C.GtkRadioAction
	// parent : record
	// Private : private_data
}

func RadioActionNewFromC(u unsafe.Pointer) *RadioAction {
	c := (*C.GtkRadioAction)(u)
	if c == nil {
		return nil
	}

	g := &RadioAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButton is a wrapper around the C record GtkRadioButton.
type RadioButton struct {
	native *C.GtkRadioButton
	// check_button : record
	// Private : priv
}

func RadioButtonNewFromC(u unsafe.Pointer) *RadioButton {
	c := (*C.GtkRadioButton)(u)
	if c == nil {
		return nil
	}

	g := &RadioButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonNew is a wrapper around the C function gtk_radio_button_new.
func RadioButtonNew(group *glib.SList) *RadioButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	retC := C.gtk_radio_button_new(c_group)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewFromWidget is a wrapper around the C function gtk_radio_button_new_from_widget.
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(C.NULL)
	if radioGroupMember != nil {
		c_radio_group_member = (*C.GtkRadioButton)(radioGroupMember.ToC())
	}

	retC := C.gtk_radio_button_new_from_widget(c_radio_group_member)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabel is a wrapper around the C function gtk_radio_button_new_with_label.
func RadioButtonNewWithLabel(group *glib.SList, label string) *RadioButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label(c_group, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabelFromWidget is a wrapper around the C function gtk_radio_button_new_with_label_from_widget.
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(C.NULL)
	if radioGroupMember != nil {
		c_radio_group_member = (*C.GtkRadioButton)(radioGroupMember.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label_from_widget(c_radio_group_member, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonic is a wrapper around the C function gtk_radio_button_new_with_mnemonic.
func RadioButtonNewWithMnemonic(group *glib.SList, label string) *RadioButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic(c_group, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_button_new_with_mnemonic_from_widget.
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(C.NULL)
	if radioGroupMember != nil {
		c_radio_group_member = (*C.GtkRadioButton)(radioGroupMember.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic_from_widget(c_radio_group_member, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonAccessible is a wrapper around the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native *C.GtkRadioButtonAccessible
	// parent : record
	// priv : record
}

func RadioButtonAccessibleNewFromC(u unsafe.Pointer) *RadioButtonAccessible {
	c := (*C.GtkRadioButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItem is a wrapper around the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native *C.GtkRadioMenuItem
	// check_menu_item : record
	// Private : priv
}

func RadioMenuItemNewFromC(u unsafe.Pointer) *RadioMenuItem {
	c := (*C.GtkRadioMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemNew is a wrapper around the C function gtk_radio_menu_item_new.
func RadioMenuItemNew(group *glib.SList) *RadioMenuItem {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	retC := C.gtk_radio_menu_item_new(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabel is a wrapper around the C function gtk_radio_menu_item_new_with_label.
func RadioMenuItemNewWithLabel(group *glib.SList, label string) *RadioMenuItem {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonic is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic.
func RadioMenuItemNewWithMnemonic(group *glib.SList, label string) *RadioMenuItem {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemAccessible is a wrapper around the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native *C.GtkRadioMenuItemAccessible
	// parent : record
	// priv : record
}

func RadioMenuItemAccessibleNewFromC(u unsafe.Pointer) *RadioMenuItemAccessible {
	c := (*C.GtkRadioMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioMenuItemAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioToolButton is a wrapper around the C record GtkRadioToolButton.
type RadioToolButton struct {
	native *C.GtkRadioToolButton
	// parent : record
}

func RadioToolButtonNewFromC(u unsafe.Pointer) *RadioToolButton {
	c := (*C.GtkRadioToolButton)(u)
	if c == nil {
		return nil
	}

	g := &RadioToolButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RadioToolButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RadioToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Range is a wrapper around the C record GtkRange.
type Range struct {
	native *C.GtkRange
	// widget : record
	// priv : record
}

func RangeNewFromC(u unsafe.Pointer) *Range {
	c := (*C.GtkRange)(u)
	if c == nil {
		return nil
	}

	g := &Range{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Range) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangeAccessible is a wrapper around the C record GtkRangeAccessible.
type RangeAccessible struct {
	native *C.GtkRangeAccessible
	// parent : record
	// priv : record
}

func RangeAccessibleNewFromC(u unsafe.Pointer) *RangeAccessible {
	c := (*C.GtkRangeAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RangeAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RangeAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RangeAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RcStyle is a wrapper around the C record GtkRcStyle.
type RcStyle struct {
	native *C.GtkRcStyle
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
	c := (*C.GtkRcStyle)(u)
	if c == nil {
		return nil
	}

	g := &RcStyle{
		Name:       C.GoString(c.name),
		Xthickness: (int32)(c.xthickness),
		Ythickness: (int32)(c.ythickness),
		native:     c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RcStyle) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RcStyle) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.xthickness =
		(C.gint)(recv.Xthickness)
	recv.native.ythickness =
		(C.gint)(recv.Ythickness)

	return (unsafe.Pointer)(recv.native)
}

// RcStyleNew is a wrapper around the C function gtk_rc_style_new.
func RcStyleNew() *RcStyle {
	retC := C.gtk_rc_style_new()
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// RecentAction is a wrapper around the C record GtkRecentAction.
type RecentAction struct {
	native *C.GtkRecentAction
	// parent_instance : record
	// Private : priv
}

func RecentActionNewFromC(u unsafe.Pointer) *RecentAction {
	c := (*C.GtkRecentAction)(u)
	if c == nil {
		return nil
	}

	g := &RecentAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserDialog is a wrapper around the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native *C.GtkRecentChooserDialog
	// parent_instance : record
	// Private : priv
}

func RecentChooserDialogNewFromC(u unsafe.Pointer) *RecentChooserDialog {
	c := (*C.GtkRecentChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserDialog{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentChooserDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserMenu is a wrapper around the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native *C.GtkRecentChooserMenu
	// parent_instance : record
	// Private : priv
}

func RecentChooserMenuNewFromC(u unsafe.Pointer) *RecentChooserMenu {
	c := (*C.GtkRecentChooserMenu)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserMenu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentChooserMenu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentChooserMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentChooserWidget is a wrapper around the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native *C.GtkRecentChooserWidget
	// parent_instance : record
	// Private : priv
}

func RecentChooserWidgetNewFromC(u unsafe.Pointer) *RecentChooserWidget {
	c := (*C.GtkRecentChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentFilter is a wrapper around the C record GtkRecentFilter.
type RecentFilter struct {
	native *C.GtkRecentFilter
}

func RecentFilterNewFromC(u unsafe.Pointer) *RecentFilter {
	c := (*C.GtkRecentFilter)(u)
	if c == nil {
		return nil
	}

	g := &RecentFilter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentFilter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessible is a wrapper around the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native *C.GtkRendererCellAccessible
	// parent : record
	// priv : record
}

func RendererCellAccessibleNewFromC(u unsafe.Pointer) *RendererCellAccessible {
	c := (*C.GtkRendererCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RendererCellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RendererCellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RendererCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessibleNew is a wrapper around the C function gtk_renderer_cell_accessible_new.
func RendererCellAccessibleNew(renderer *CellRenderer) *RendererCellAccessible {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_renderer_cell_accessible_new(c_renderer)
	retGo := RendererCellAccessibleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Revealer is a wrapper around the C record GtkRevealer.
type Revealer struct {
	native *C.GtkRevealer
	// parent_instance : record
}

func RevealerNewFromC(u unsafe.Pointer) *Revealer {
	c := (*C.GtkRevealer)(u)
	if c == nil {
		return nil
	}

	g := &Revealer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Revealer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Revealer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scale is a wrapper around the C record GtkScale.
type Scale struct {
	native *C.GtkScale
	// range : record
	// Private : priv
}

func ScaleNewFromC(u unsafe.Pointer) *Scale {
	c := (*C.GtkScale)(u)
	if c == nil {
		return nil
	}

	g := &Scale{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Scale) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Scale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleAccessible is a wrapper around the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native *C.GtkScaleAccessible
	// parent : record
	// priv : record
}

func ScaleAccessibleNewFromC(u unsafe.Pointer) *ScaleAccessible {
	c := (*C.GtkScaleAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScaleAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ScaleAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ScaleAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButton is a wrapper around the C record GtkScaleButton.
type ScaleButton struct {
	native *C.GtkScaleButton
	// parent : record
	// Private : priv
}

func ScaleButtonNewFromC(u unsafe.Pointer) *ScaleButton {
	c := (*C.GtkScaleButton)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ScaleButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ScaleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButtonAccessible is a wrapper around the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native *C.GtkScaleButtonAccessible
	// parent : record
	// priv : record
}

func ScaleButtonAccessibleNewFromC(u unsafe.Pointer) *ScaleButtonAccessible {
	c := (*C.GtkScaleButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ScaleButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ScaleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scrollbar is a wrapper around the C record GtkScrollbar.
type Scrollbar struct {
	native *C.GtkScrollbar
	// range : record
}

func ScrollbarNewFromC(u unsafe.Pointer) *Scrollbar {
	c := (*C.GtkScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &Scrollbar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Scrollbar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Scrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindow is a wrapper around the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native *C.GtkScrolledWindow
	// container : record
	// priv : record
}

func ScrolledWindowNewFromC(u unsafe.Pointer) *ScrolledWindow {
	c := (*C.GtkScrolledWindow)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindow{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ScrolledWindow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ScrolledWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowNew is a wrapper around the C function gtk_scrolled_window_new.
func ScrolledWindowNew(hadjustment *Adjustment, vadjustment *Adjustment) *ScrolledWindow {
	c_hadjustment := (*C.GtkAdjustment)(C.NULL)
	if hadjustment != nil {
		c_hadjustment = (*C.GtkAdjustment)(hadjustment.ToC())
	}

	c_vadjustment := (*C.GtkAdjustment)(C.NULL)
	if vadjustment != nil {
		c_vadjustment = (*C.GtkAdjustment)(vadjustment.ToC())
	}

	retC := C.gtk_scrolled_window_new(c_hadjustment, c_vadjustment)
	retGo := ScrolledWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScrolledWindowAccessible is a wrapper around the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native *C.GtkScrolledWindowAccessible
	// parent : record
	// priv : record
}

func ScrolledWindowAccessibleNewFromC(u unsafe.Pointer) *ScrolledWindowAccessible {
	c := (*C.GtkScrolledWindowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ScrolledWindowAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ScrolledWindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SearchBar is a wrapper around the C record GtkSearchBar.
type SearchBar struct {
	native *C.GtkSearchBar
	// Private : parent
}

func SearchBarNewFromC(u unsafe.Pointer) *SearchBar {
	c := (*C.GtkSearchBar)(u)
	if c == nil {
		return nil
	}

	g := &SearchBar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SearchBar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SearchBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SearchEntry is a wrapper around the C record GtkSearchEntry.
type SearchEntry struct {
	native *C.GtkSearchEntry
	// parent : record
}

func SearchEntryNewFromC(u unsafe.Pointer) *SearchEntry {
	c := (*C.GtkSearchEntry)(u)
	if c == nil {
		return nil
	}

	g := &SearchEntry{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SearchEntry) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SearchEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Separator is a wrapper around the C record GtkSeparator.
type Separator struct {
	native *C.GtkSeparator
	// widget : record
	// priv : record
}

func SeparatorNewFromC(u unsafe.Pointer) *Separator {
	c := (*C.GtkSeparator)(u)
	if c == nil {
		return nil
	}

	g := &Separator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Separator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Separator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorMenuItem is a wrapper around the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native *C.GtkSeparatorMenuItem
	// menu_item : record
}

func SeparatorMenuItemNewFromC(u unsafe.Pointer) *SeparatorMenuItem {
	c := (*C.GtkSeparatorMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SeparatorMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorMenuItemNew is a wrapper around the C function gtk_separator_menu_item_new.
func SeparatorMenuItemNew() *SeparatorMenuItem {
	retC := C.gtk_separator_menu_item_new()
	retGo := SeparatorMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SeparatorToolItem is a wrapper around the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native *C.GtkSeparatorToolItem
	// parent : record
	// Private : priv
}

func SeparatorToolItemNewFromC(u unsafe.Pointer) *SeparatorToolItem {
	c := (*C.GtkSeparatorToolItem)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorToolItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SeparatorToolItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SeparatorToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Settings is a wrapper around the C record GtkSettings.
type Settings struct {
	native *C.GtkSettings
	// parent_instance : record
	// Private : priv
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.GtkSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Settings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SizeGroup is a wrapper around the C record GtkSizeGroup.
type SizeGroup struct {
	native *C.GtkSizeGroup
	// parent_instance : record
	// Private : priv
}

func SizeGroupNewFromC(u unsafe.Pointer) *SizeGroup {
	c := (*C.GtkSizeGroup)(u)
	if c == nil {
		return nil
	}

	g := &SizeGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SizeGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SizeGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SizeGroupNew is a wrapper around the C function gtk_size_group_new.
func SizeGroupNew(mode SizeGroupMode) *SizeGroup {
	c_mode := (C.GtkSizeGroupMode)(mode)

	retC := C.gtk_size_group_new(c_mode)
	retGo := SizeGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : GtkSocket

// SpinButton is a wrapper around the C record GtkSpinButton.
type SpinButton struct {
	native *C.GtkSpinButton
	// entry : record
	// Private : priv
}

func SpinButtonNewFromC(u unsafe.Pointer) *SpinButton {
	c := (*C.GtkSpinButton)(u)
	if c == nil {
		return nil
	}

	g := &SpinButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SpinButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SpinButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonNew is a wrapper around the C function gtk_spin_button_new.
func SpinButtonNew(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	c_climb_rate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	retC := C.gtk_spin_button_new(c_adjustment, c_climb_rate, c_digits)
	retGo := SpinButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SpinButtonNewWithRange is a wrapper around the C function gtk_spin_button_new_with_range.
func SpinButtonNewWithRange(min float64, max float64, step float64) *SpinButton {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_spin_button_new_with_range(c_min, c_max, c_step)
	retGo := SpinButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SpinButtonAccessible is a wrapper around the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native *C.GtkSpinButtonAccessible
	// parent : record
	// priv : record
}

func SpinButtonAccessibleNewFromC(u unsafe.Pointer) *SpinButtonAccessible {
	c := (*C.GtkSpinButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SpinButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SpinButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Spinner is a wrapper around the C record GtkSpinner.
type Spinner struct {
	native *C.GtkSpinner
	// parent : record
	// Private : priv
}

func SpinnerNewFromC(u unsafe.Pointer) *Spinner {
	c := (*C.GtkSpinner)(u)
	if c == nil {
		return nil
	}

	g := &Spinner{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Spinner) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Spinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinnerAccessible is a wrapper around the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native *C.GtkSpinnerAccessible
	// parent : record
	// priv : record
}

func SpinnerAccessibleNewFromC(u unsafe.Pointer) *SpinnerAccessible {
	c := (*C.GtkSpinnerAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SpinnerAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SpinnerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Stack is a wrapper around the C record GtkStack.
type Stack struct {
	native *C.GtkStack
	// parent_instance : record
}

func StackNewFromC(u unsafe.Pointer) *Stack {
	c := (*C.GtkStack)(u)
	if c == nil {
		return nil
	}

	g := &Stack{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Stack) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Stack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GtkStackAccessible

// StackSidebar is a wrapper around the C record GtkStackSidebar.
type StackSidebar struct {
	native *C.GtkStackSidebar
	// parent : record
}

func StackSidebarNewFromC(u unsafe.Pointer) *StackSidebar {
	c := (*C.GtkStackSidebar)(u)
	if c == nil {
		return nil
	}

	g := &StackSidebar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StackSidebar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StackSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StackSwitcher is a wrapper around the C record GtkStackSwitcher.
type StackSwitcher struct {
	native *C.GtkStackSwitcher
	// widget : record
}

func StackSwitcherNewFromC(u unsafe.Pointer) *StackSwitcher {
	c := (*C.GtkStackSwitcher)(u)
	if c == nil {
		return nil
	}

	g := &StackSwitcher{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StackSwitcher) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StackSwitcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusIcon is a wrapper around the C record GtkStatusIcon.
type StatusIcon struct {
	native *C.GtkStatusIcon
	// parent_instance : record
	// priv : record
}

func StatusIconNewFromC(u unsafe.Pointer) *StatusIcon {
	c := (*C.GtkStatusIcon)(u)
	if c == nil {
		return nil
	}

	g := &StatusIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StatusIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StatusIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Statusbar is a wrapper around the C record GtkStatusbar.
type Statusbar struct {
	native *C.GtkStatusbar
	// parent_widget : record
	// Private : priv
}

func StatusbarNewFromC(u unsafe.Pointer) *Statusbar {
	c := (*C.GtkStatusbar)(u)
	if c == nil {
		return nil
	}

	g := &Statusbar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Statusbar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Statusbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarNew is a wrapper around the C function gtk_statusbar_new.
func StatusbarNew() *Statusbar {
	retC := C.gtk_statusbar_new()
	retGo := StatusbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StatusbarAccessible is a wrapper around the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native *C.GtkStatusbarAccessible
	// parent : record
	// priv : record
}

func StatusbarAccessibleNewFromC(u unsafe.Pointer) *StatusbarAccessible {
	c := (*C.GtkStatusbarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StatusbarAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StatusbarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Style is a wrapper around the C record GtkStyle.
type Style struct {
	native *C.GtkStyle
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
	c := (*C.GtkStyle)(u)
	if c == nil {
		return nil
	}

	g := &Style{
		Xthickness: (int32)(c.xthickness),
		Ythickness: (int32)(c.ythickness),
		native:     c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Style) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Style) ToC() unsafe.Pointer {
	recv.native.xthickness =
		(C.gint)(recv.Xthickness)
	recv.native.ythickness =
		(C.gint)(recv.Ythickness)

	return (unsafe.Pointer)(recv.native)
}

// StyleNew is a wrapper around the C function gtk_style_new.
func StyleNew() *Style {
	retC := C.gtk_style_new()
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StyleContext is a wrapper around the C record GtkStyleContext.
type StyleContext struct {
	native *C.GtkStyleContext
	// parent_object : record
	// priv : record
}

func StyleContextNewFromC(u unsafe.Pointer) *StyleContext {
	c := (*C.GtkStyleContext)(u)
	if c == nil {
		return nil
	}

	g := &StyleContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleContextNew is a wrapper around the C function gtk_style_context_new.
func StyleContextNew() *StyleContext {
	retC := C.gtk_style_context_new()
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StyleProperties is a wrapper around the C record GtkStyleProperties.
type StyleProperties struct {
	native *C.GtkStyleProperties
	// Private : parent_object
	// Private : priv
}

func StylePropertiesNewFromC(u unsafe.Pointer) *StyleProperties {
	c := (*C.GtkStyleProperties)(u)
	if c == nil {
		return nil
	}

	g := &StyleProperties{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleProperties) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleProperties) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StylePropertiesNew is a wrapper around the C function gtk_style_properties_new.
func StylePropertiesNew() *StyleProperties {
	retC := C.gtk_style_properties_new()
	retGo := StylePropertiesNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Switch is a wrapper around the C record GtkSwitch.
type Switch struct {
	native *C.GtkSwitch
	// Private : parent_instance
	// Private : priv
}

func SwitchNewFromC(u unsafe.Pointer) *Switch {
	c := (*C.GtkSwitch)(u)
	if c == nil {
		return nil
	}

	g := &Switch{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Switch) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Switch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SwitchAccessible is a wrapper around the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native *C.GtkSwitchAccessible
	// parent : record
	// priv : record
}

func SwitchAccessibleNewFromC(u unsafe.Pointer) *SwitchAccessible {
	c := (*C.GtkSwitchAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SwitchAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SwitchAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SwitchAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Table is a wrapper around the C record GtkTable.
type Table struct {
	native *C.GtkTable
	// container : record
	// Private : priv
}

func TableNewFromC(u unsafe.Pointer) *Table {
	c := (*C.GtkTable)(u)
	if c == nil {
		return nil
	}

	g := &Table{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Table) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableNew is a wrapper around the C function gtk_table_new.
func TableNew(rows uint32, columns uint32, homogeneous bool) *Table {
	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	c_homogeneous :=
		boolToGboolean(homogeneous)

	retC := C.gtk_table_new(c_rows, c_columns, c_homogeneous)
	retGo := TableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TearoffMenuItem is a wrapper around the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native *C.GtkTearoffMenuItem
	// menu_item : record
	// Private : priv
}

func TearoffMenuItemNewFromC(u unsafe.Pointer) *TearoffMenuItem {
	c := (*C.GtkTearoffMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &TearoffMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TearoffMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TearoffMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TearoffMenuItemNew is a wrapper around the C function gtk_tearoff_menu_item_new.
func TearoffMenuItemNew() *TearoffMenuItem {
	retC := C.gtk_tearoff_menu_item_new()
	retGo := TearoffMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextBuffer is a wrapper around the C record GtkTextBuffer.
type TextBuffer struct {
	native *C.GtkTextBuffer
	// parent_instance : record
	// priv : record
}

func TextBufferNewFromC(u unsafe.Pointer) *TextBuffer {
	c := (*C.GtkTextBuffer)(u)
	if c == nil {
		return nil
	}

	g := &TextBuffer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextBuffer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextBufferNew is a wrapper around the C function gtk_text_buffer_new.
func TextBufferNew(table *TextTagTable) *TextBuffer {
	c_table := (*C.GtkTextTagTable)(C.NULL)
	if table != nil {
		c_table = (*C.GtkTextTagTable)(table.ToC())
	}

	retC := C.gtk_text_buffer_new(c_table)
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TextCellAccessible is a wrapper around the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native *C.GtkTextCellAccessible
	// parent : record
	// priv : record
}

func TextCellAccessibleNewFromC(u unsafe.Pointer) *TextCellAccessible {
	c := (*C.GtkTextCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TextCellAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextCellAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextChildAnchor is a wrapper around the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native *C.GtkTextChildAnchor
	// parent_instance : record
	// Private : segment
}

func TextChildAnchorNewFromC(u unsafe.Pointer) *TextChildAnchor {
	c := (*C.GtkTextChildAnchor)(u)
	if c == nil {
		return nil
	}

	g := &TextChildAnchor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextChildAnchor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextChildAnchor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextChildAnchorNew is a wrapper around the C function gtk_text_child_anchor_new.
func TextChildAnchorNew() *TextChildAnchor {
	retC := C.gtk_text_child_anchor_new()
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TextMark is a wrapper around the C record GtkTextMark.
type TextMark struct {
	native *C.GtkTextMark
	// parent_instance : record
	// Private : segment
}

func TextMarkNewFromC(u unsafe.Pointer) *TextMark {
	c := (*C.GtkTextMark)(u)
	if c == nil {
		return nil
	}

	g := &TextMark{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextMark) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextMark) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTag is a wrapper around the C record GtkTextTag.
type TextTag struct {
	native *C.GtkTextTag
	// parent_instance : record
	// priv : record
}

func TextTagNewFromC(u unsafe.Pointer) *TextTag {
	c := (*C.GtkTextTag)(u)
	if c == nil {
		return nil
	}

	g := &TextTag{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextTag) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextTag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagNew is a wrapper around the C function gtk_text_tag_new.
func TextTagNew(name string) *TextTag {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_tag_new(c_name)
	retGo := TextTagNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TextTagTable is a wrapper around the C record GtkTextTagTable.
type TextTagTable struct {
	native *C.GtkTextTagTable
	// parent_instance : record
	// priv : record
}

func TextTagTableNewFromC(u unsafe.Pointer) *TextTagTable {
	c := (*C.GtkTextTagTable)(u)
	if c == nil {
		return nil
	}

	g := &TextTagTable{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextTagTable) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextTagTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagTableNew is a wrapper around the C function gtk_text_tag_table_new.
func TextTagTableNew() *TextTagTable {
	retC := C.gtk_text_tag_table_new()
	retGo := TextTagTableNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TextView is a wrapper around the C record GtkTextView.
type TextView struct {
	native *C.GtkTextView
	// parent_instance : record
	// Private : priv
}

func TextViewNewFromC(u unsafe.Pointer) *TextView {
	c := (*C.GtkTextView)(u)
	if c == nil {
		return nil
	}

	g := &TextView{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextView) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewNew is a wrapper around the C function gtk_text_view_new.
func TextViewNew() *TextView {
	retC := C.gtk_text_view_new()
	retGo := TextViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextViewNewWithBuffer is a wrapper around the C function gtk_text_view_new_with_buffer.
func TextViewNewWithBuffer(buffer *TextBuffer) *TextView {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	retC := C.gtk_text_view_new_with_buffer(c_buffer)
	retGo := TextViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextViewAccessible is a wrapper around the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native *C.GtkTextViewAccessible
	// parent : record
	// priv : record
}

func TextViewAccessibleNewFromC(u unsafe.Pointer) *TextViewAccessible {
	c := (*C.GtkTextViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TextViewAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TextViewAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TextViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemingEngine is a wrapper around the C record GtkThemingEngine.
type ThemingEngine struct {
	native *C.GtkThemingEngine
	// parent_object : record
	// priv : record
}

func ThemingEngineNewFromC(u unsafe.Pointer) *ThemingEngine {
	c := (*C.GtkThemingEngine)(u)
	if c == nil {
		return nil
	}

	g := &ThemingEngine{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ThemingEngine) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ThemingEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleAction is a wrapper around the C record GtkToggleAction.
type ToggleAction struct {
	native *C.GtkToggleAction
	// parent : record
	// Private : private_data
}

func ToggleActionNewFromC(u unsafe.Pointer) *ToggleAction {
	c := (*C.GtkToggleAction)(u)
	if c == nil {
		return nil
	}

	g := &ToggleAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToggleAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToggleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButton is a wrapper around the C record GtkToggleButton.
type ToggleButton struct {
	native *C.GtkToggleButton
	// Private : button
	// Private : priv
}

func ToggleButtonNewFromC(u unsafe.Pointer) *ToggleButton {
	c := (*C.GtkToggleButton)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToggleButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToggleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonNew is a wrapper around the C function gtk_toggle_button_new.
func ToggleButtonNew() *ToggleButton {
	retC := C.gtk_toggle_button_new()
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithLabel is a wrapper around the C function gtk_toggle_button_new_with_label.
func ToggleButtonNewWithLabel(label string) *ToggleButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_label(c_label)
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithMnemonic is a wrapper around the C function gtk_toggle_button_new_with_mnemonic.
func ToggleButtonNewWithMnemonic(label string) *ToggleButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_mnemonic(c_label)
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonAccessible is a wrapper around the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native *C.GtkToggleButtonAccessible
	// parent : record
	// priv : record
}

func ToggleButtonAccessibleNewFromC(u unsafe.Pointer) *ToggleButtonAccessible {
	c := (*C.GtkToggleButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToggleButtonAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToggleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleToolButton is a wrapper around the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native *C.GtkToggleToolButton
	// parent : record
	// Private : priv
}

func ToggleToolButtonNewFromC(u unsafe.Pointer) *ToggleToolButton {
	c := (*C.GtkToggleToolButton)(u)
	if c == nil {
		return nil
	}

	g := &ToggleToolButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToggleToolButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToggleToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolButton is a wrapper around the C record GtkToolButton.
type ToolButton struct {
	native *C.GtkToolButton
	// parent : record
	// Private : priv
}

func ToolButtonNewFromC(u unsafe.Pointer) *ToolButton {
	c := (*C.GtkToolButton)(u)
	if c == nil {
		return nil
	}

	g := &ToolButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToolButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItem is a wrapper around the C record GtkToolItem.
type ToolItem struct {
	native *C.GtkToolItem
	// parent : record
	// Private : priv
}

func ToolItemNewFromC(u unsafe.Pointer) *ToolItem {
	c := (*C.GtkToolItem)(u)
	if c == nil {
		return nil
	}

	g := &ToolItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToolItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItemGroup is a wrapper around the C record GtkToolItemGroup.
type ToolItemGroup struct {
	native *C.GtkToolItemGroup
	// parent_instance : record
	// priv : record
}

func ToolItemGroupNewFromC(u unsafe.Pointer) *ToolItemGroup {
	c := (*C.GtkToolItemGroup)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToolItemGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToolItemGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolPalette is a wrapper around the C record GtkToolPalette.
type ToolPalette struct {
	native *C.GtkToolPalette
	// parent_instance : record
	// priv : record
}

func ToolPaletteNewFromC(u unsafe.Pointer) *ToolPalette {
	c := (*C.GtkToolPalette)(u)
	if c == nil {
		return nil
	}

	g := &ToolPalette{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToolPalette) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToolPalette) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Toolbar is a wrapper around the C record GtkToolbar.
type Toolbar struct {
	native *C.GtkToolbar
	// container : record
	// priv : record
}

func ToolbarNewFromC(u unsafe.Pointer) *Toolbar {
	c := (*C.GtkToolbar)(u)
	if c == nil {
		return nil
	}

	g := &Toolbar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Toolbar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Toolbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolbarNew is a wrapper around the C function gtk_toolbar_new.
func ToolbarNew() *Toolbar {
	retC := C.gtk_toolbar_new()
	retGo := ToolbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Tooltip is a wrapper around the C record GtkTooltip.
type Tooltip struct {
	native *C.GtkTooltip
}

func TooltipNewFromC(u unsafe.Pointer) *Tooltip {
	c := (*C.GtkTooltip)(u)
	if c == nil {
		return nil
	}

	g := &Tooltip{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Tooltip) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Tooltip) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToplevelAccessible is a wrapper around the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native *C.GtkToplevelAccessible
	// parent : record
	// priv : record
}

func ToplevelAccessibleNewFromC(u unsafe.Pointer) *ToplevelAccessible {
	c := (*C.GtkToplevelAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ToplevelAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ToplevelAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ToplevelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelFilter is a wrapper around the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native *C.GtkTreeModelFilter
	// parent : record
	// Private : priv
}

func TreeModelFilterNewFromC(u unsafe.Pointer) *TreeModelFilter {
	c := (*C.GtkTreeModelFilter)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelFilter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeModelFilter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeModelFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeModelSort is a wrapper around the C record GtkTreeModelSort.
type TreeModelSort struct {
	native *C.GtkTreeModelSort
	// parent : record
	// Private : priv
}

func TreeModelSortNewFromC(u unsafe.Pointer) *TreeModelSort {
	c := (*C.GtkTreeModelSort)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelSort{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeModelSort) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeModelSort) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeSelection is a wrapper around the C record GtkTreeSelection.
type TreeSelection struct {
	native *C.GtkTreeSelection
	// Private : parent
	// Private : priv
}

func TreeSelectionNewFromC(u unsafe.Pointer) *TreeSelection {
	c := (*C.GtkTreeSelection)(u)
	if c == nil {
		return nil
	}

	g := &TreeSelection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeSelection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeStore is a wrapper around the C record GtkTreeStore.
type TreeStore struct {
	native *C.GtkTreeStore
	// parent : record
	// priv : record
}

func TreeStoreNewFromC(u unsafe.Pointer) *TreeStore {
	c := (*C.GtkTreeStore)(u)
	if c == nil {
		return nil
	}

	g := &TreeStore{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeStore) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// TreeStoreNewv is a wrapper around the C function gtk_tree_store_newv.
func TreeStoreNewv(types []gobject.Type) *TreeStore {
	c_n_columns := (C.gint)(len(types))

	c_types_array := make([]C.GType, len(types)+1, len(types)+1)
	for i, item := range types {
		c := (C.GType)(item)
		c_types_array[i] = c
	}
	c_types_array[len(types)] = 0
	c_types_arrayPtr := &c_types_array[0]
	c_types := (*C.GType)(unsafe.Pointer(c_types_arrayPtr))

	retC := C.gtk_tree_store_newv(c_n_columns, c_types)
	retGo := TreeStoreNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TreeView is a wrapper around the C record GtkTreeView.
type TreeView struct {
	native *C.GtkTreeView
	// parent : record
	// Private : priv
}

func TreeViewNewFromC(u unsafe.Pointer) *TreeView {
	c := (*C.GtkTreeView)(u)
	if c == nil {
		return nil
	}

	g := &TreeView{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeView) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewNew is a wrapper around the C function gtk_tree_view_new.
func TreeViewNew() *TreeView {
	retC := C.gtk_tree_view_new()
	retGo := TreeViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreeViewNewWithModel is a wrapper around the C function gtk_tree_view_new_with_model.
func TreeViewNewWithModel(model *TreeModel) *TreeView {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_tree_view_new_with_model(c_model)
	retGo := TreeViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreeViewAccessible is a wrapper around the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native *C.GtkTreeViewAccessible
	// parent : record
	// priv : record
}

func TreeViewAccessibleNewFromC(u unsafe.Pointer) *TreeViewAccessible {
	c := (*C.GtkTreeViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeViewAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewColumn is a wrapper around the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native *C.GtkTreeViewColumn
	// parent_instance : record
	// priv : record
}

func TreeViewColumnNewFromC(u unsafe.Pointer) *TreeViewColumn {
	c := (*C.GtkTreeViewColumn)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewColumn{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TreeViewColumn) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TreeViewColumn) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewColumnNew is a wrapper around the C function gtk_tree_view_column_new.
func TreeViewColumnNew() *TreeViewColumn {
	retC := C.gtk_tree_view_column_new()
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// UIManager is a wrapper around the C record GtkUIManager.
type UIManager struct {
	native *C.GtkUIManager
	// parent : record
	// Private : private_data
}

func UIManagerNewFromC(u unsafe.Pointer) *UIManager {
	c := (*C.GtkUIManager)(u)
	if c == nil {
		return nil
	}

	g := &UIManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UIManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UIManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VBox is a wrapper around the C record GtkVBox.
type VBox struct {
	native *C.GtkVBox
	// box : record
}

func VBoxNewFromC(u unsafe.Pointer) *VBox {
	c := (*C.GtkVBox)(u)
	if c == nil {
		return nil
	}

	g := &VBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VBoxNew is a wrapper around the C function gtk_vbox_new.
func VBoxNew(homogeneous bool, spacing int32) *VBox {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_vbox_new(c_homogeneous, c_spacing)
	retGo := VBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VButtonBox is a wrapper around the C record GtkVButtonBox.
type VButtonBox struct {
	native *C.GtkVButtonBox
	// button_box : record
}

func VButtonBoxNewFromC(u unsafe.Pointer) *VButtonBox {
	c := (*C.GtkVButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &VButtonBox{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VButtonBox) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VButtonBoxNew is a wrapper around the C function gtk_vbutton_box_new.
func VButtonBoxNew() *VButtonBox {
	retC := C.gtk_vbutton_box_new()
	retGo := VButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VPaned is a wrapper around the C record GtkVPaned.
type VPaned struct {
	native *C.GtkVPaned
	// paned : record
}

func VPanedNewFromC(u unsafe.Pointer) *VPaned {
	c := (*C.GtkVPaned)(u)
	if c == nil {
		return nil
	}

	g := &VPaned{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VPaned) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VPanedNew is a wrapper around the C function gtk_vpaned_new.
func VPanedNew() *VPaned {
	retC := C.gtk_vpaned_new()
	retGo := VPanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScale is a wrapper around the C record GtkVScale.
type VScale struct {
	native *C.GtkVScale
	// scale : record
}

func VScaleNewFromC(u unsafe.Pointer) *VScale {
	c := (*C.GtkVScale)(u)
	if c == nil {
		return nil
	}

	g := &VScale{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VScale) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScaleNew is a wrapper around the C function gtk_vscale_new.
func VScaleNew(adjustment *Adjustment) *VScale {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_vscale_new(c_adjustment)
	retGo := VScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScaleNewWithRange is a wrapper around the C function gtk_vscale_new_with_range.
func VScaleNewWithRange(min float64, max float64, step float64) *VScale {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_vscale_new_with_range(c_min, c_max, c_step)
	retGo := VScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScrollbar is a wrapper around the C record GtkVScrollbar.
type VScrollbar struct {
	native *C.GtkVScrollbar
	// scrollbar : record
}

func VScrollbarNewFromC(u unsafe.Pointer) *VScrollbar {
	c := (*C.GtkVScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &VScrollbar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VScrollbar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScrollbarNew is a wrapper around the C function gtk_vscrollbar_new.
func VScrollbarNew(adjustment *Adjustment) *VScrollbar {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_vscrollbar_new(c_adjustment)
	retGo := VScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VSeparator is a wrapper around the C record GtkVSeparator.
type VSeparator struct {
	native *C.GtkVSeparator
	// separator : record
}

func VSeparatorNewFromC(u unsafe.Pointer) *VSeparator {
	c := (*C.GtkVSeparator)(u)
	if c == nil {
		return nil
	}

	g := &VSeparator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VSeparator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VSeparatorNew is a wrapper around the C function gtk_vseparator_new.
func VSeparatorNew() *VSeparator {
	retC := C.gtk_vseparator_new()
	retGo := VSeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Viewport is a wrapper around the C record GtkViewport.
type Viewport struct {
	native *C.GtkViewport
	// bin : record
	// Private : priv
}

func ViewportNewFromC(u unsafe.Pointer) *Viewport {
	c := (*C.GtkViewport)(u)
	if c == nil {
		return nil
	}

	g := &Viewport{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Viewport) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Viewport) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ViewportNew is a wrapper around the C function gtk_viewport_new.
func ViewportNew(hadjustment *Adjustment, vadjustment *Adjustment) *Viewport {
	c_hadjustment := (*C.GtkAdjustment)(C.NULL)
	if hadjustment != nil {
		c_hadjustment = (*C.GtkAdjustment)(hadjustment.ToC())
	}

	c_vadjustment := (*C.GtkAdjustment)(C.NULL)
	if vadjustment != nil {
		c_vadjustment = (*C.GtkAdjustment)(vadjustment.ToC())
	}

	retC := C.gtk_viewport_new(c_hadjustment, c_vadjustment)
	retGo := ViewportNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VolumeButton is a wrapper around the C record GtkVolumeButton.
type VolumeButton struct {
	native *C.GtkVolumeButton
	// parent : record
}

func VolumeButtonNewFromC(u unsafe.Pointer) *VolumeButton {
	c := (*C.GtkVolumeButton)(u)
	if c == nil {
		return nil
	}

	g := &VolumeButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VolumeButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VolumeButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget is a wrapper around the C record GtkWidget.
type Widget struct {
	native *C.GtkWidget
	// parent_instance : record
	// Private : priv
}

func WidgetNewFromC(u unsafe.Pointer) *Widget {
	c := (*C.GtkWidget)(u)
	if c == nil {
		return nil
	}

	g := &Widget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Widget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Widget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// WidgetAccessible is a wrapper around the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native *C.GtkWidgetAccessible
	// parent : record
	// priv : record
}

func WidgetAccessibleNewFromC(u unsafe.Pointer) *WidgetAccessible {
	c := (*C.GtkWidgetAccessible)(u)
	if c == nil {
		return nil
	}

	g := &WidgetAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WidgetAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WidgetAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window is a wrapper around the C record GtkWindow.
type Window struct {
	native *C.GtkWindow
	// bin : record
	// priv : record
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.GtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Window) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowNew is a wrapper around the C function gtk_window_new.
func WindowNew(type_ WindowType) *Window {
	c_type := (C.GtkWindowType)(type_)

	retC := C.gtk_window_new(c_type)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// WindowAccessible is a wrapper around the C record GtkWindowAccessible.
type WindowAccessible struct {
	native *C.GtkWindowAccessible
	// parent : record
	// priv : record
}

func WindowAccessibleNewFromC(u unsafe.Pointer) *WindowAccessible {
	c := (*C.GtkWindowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &WindowAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WindowAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGroup is a wrapper around the C record GtkWindowGroup.
type WindowGroup struct {
	native *C.GtkWindowGroup
	// parent_instance : record
	// priv : record
}

func WindowGroupNewFromC(u unsafe.Pointer) *WindowGroup {
	c := (*C.GtkWindowGroup)(u)
	if c == nil {
		return nil
	}

	g := &WindowGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WindowGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WindowGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGroupNew is a wrapper around the C function gtk_window_group_new.
func WindowGroupNew() *WindowGroup {
	retC := C.gtk_window_group_new()
	retGo := WindowGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : STOCK_ADD

// Blacklisted : STOCK_APPLY

// Blacklisted : STOCK_BOLD

// Blacklisted : STOCK_CANCEL

// Blacklisted : STOCK_CDROM

// Blacklisted : STOCK_CLEAR

// Blacklisted : STOCK_CLOSE

// Blacklisted : STOCK_CONVERT

// Blacklisted : STOCK_COPY

// Blacklisted : STOCK_CUT

// Blacklisted : STOCK_DELETE

// Blacklisted : STOCK_DIALOG_ERROR

// Blacklisted : STOCK_DIALOG_INFO

// Blacklisted : STOCK_DIALOG_QUESTION

// Blacklisted : STOCK_DIALOG_WARNING

// Blacklisted : STOCK_DND

// Blacklisted : STOCK_DND_MULTIPLE

// Blacklisted : STOCK_EXECUTE

// Blacklisted : STOCK_FIND

// Blacklisted : STOCK_FIND_AND_REPLACE

// Blacklisted : STOCK_FLOPPY

// Blacklisted : STOCK_GOTO_BOTTOM

// Blacklisted : STOCK_GOTO_FIRST

// Blacklisted : STOCK_GOTO_LAST

// Blacklisted : STOCK_GOTO_TOP

// Blacklisted : STOCK_GO_BACK

// Blacklisted : STOCK_GO_DOWN

// Blacklisted : STOCK_GO_FORWARD

// Blacklisted : STOCK_GO_UP

// Blacklisted : STOCK_HELP

// Blacklisted : STOCK_HOME

// Blacklisted : STOCK_INDEX

// Blacklisted : STOCK_ITALIC

// Blacklisted : STOCK_JUMP_TO

// Blacklisted : STOCK_JUSTIFY_CENTER

// Blacklisted : STOCK_JUSTIFY_FILL

// Blacklisted : STOCK_JUSTIFY_LEFT

// Blacklisted : STOCK_JUSTIFY_RIGHT

// Blacklisted : STOCK_MISSING_IMAGE

// Blacklisted : STOCK_NEW

// Blacklisted : STOCK_NO

// Blacklisted : STOCK_OK

// Blacklisted : STOCK_OPEN

// Blacklisted : STOCK_PASTE

// Blacklisted : STOCK_PREFERENCES

// Blacklisted : STOCK_PRINT

// Blacklisted : STOCK_PRINT_PREVIEW

// Blacklisted : STOCK_PROPERTIES

// Blacklisted : STOCK_QUIT

// Blacklisted : STOCK_REDO

// Blacklisted : STOCK_REFRESH

// Blacklisted : STOCK_REMOVE

// Blacklisted : STOCK_REVERT_TO_SAVED

// Blacklisted : STOCK_SAVE

// Blacklisted : STOCK_SAVE_AS

// Blacklisted : STOCK_SELECT_COLOR

// Blacklisted : STOCK_SELECT_FONT

// Blacklisted : STOCK_SORT_ASCENDING

// Blacklisted : STOCK_SORT_DESCENDING

// Blacklisted : STOCK_SPELL_CHECK

// Blacklisted : STOCK_STOP

// Blacklisted : STOCK_STRIKETHROUGH

// Blacklisted : STOCK_UNDELETE

// Blacklisted : STOCK_UNDERLINE

// Blacklisted : STOCK_UNDO

// Blacklisted : STOCK_YES

// Blacklisted : STOCK_ZOOM_100

// Blacklisted : STOCK_ZOOM_FIT

// Blacklisted : STOCK_ZOOM_IN

// Blacklisted : STOCK_ZOOM_OUT

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_binding_set_by_class : unsupported parameter object_class : no type generator for gpointer (gpointer) for param object_class

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_selection_add_targets : unsupported parameter targets :

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

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

// IconSourceNew is a wrapper around the C function gtk_icon_source_new.
func IconSourceNew() *IconSource {
	retC := C.gtk_icon_source_new()
	retGo := IconSourceNewFromC(unsafe.Pointer(retC))

	return retGo
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

// RequestedSize is a wrapper around the C record GtkRequestedSize.
type RequestedSize struct {
	native *C.GtkRequestedSize
	// data : no type generator for gpointer, gpointer
	MinimumSize int32
	NaturalSize int32
}

func RequestedSizeNewFromC(u unsafe.Pointer) *RequestedSize {
	c := (*C.GtkRequestedSize)(u)
	if c == nil {
		return nil
	}

	g := &RequestedSize{
		MinimumSize: (int32)(c.minimum_size),
		NaturalSize: (int32)(c.natural_size),
		native:      c,
	}

	return g
}

func (recv *RequestedSize) ToC() unsafe.Pointer {
	recv.native.minimum_size =
		(C.gint)(recv.MinimumSize)
	recv.native.natural_size =
		(C.gint)(recv.NaturalSize)

	return (unsafe.Pointer)(recv.native)
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

// StatusbarClass is a wrapper around the C record GtkStatusbarClass.
type StatusbarClass struct {
	native *C.GtkStatusbarClass
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
	c := (*C.GtkStatusbarClass)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarClass{native: c}

	return g
}

func (recv *StatusbarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// Unsupported : gtk_target_list_new : unsupported parameter targets :

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

// TextAttributesNew is a wrapper around the C function gtk_text_attributes_new.
func TextAttributesNew() *TextAttributes {
	retC := C.gtk_text_attributes_new()
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
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

// TreeIter is a wrapper around the C record GtkTreeIter.
type TreeIter struct {
	native *C.GtkTreeIter
	Stamp  int32
	// user_data : no type generator for gpointer, gpointer
	// user_data2 : no type generator for gpointer, gpointer
	// user_data3 : no type generator for gpointer, gpointer
}

func TreeIterNewFromC(u unsafe.Pointer) *TreeIter {
	c := (*C.GtkTreeIter)(u)
	if c == nil {
		return nil
	}

	g := &TreeIter{
		Stamp:  (int32)(c.stamp),
		native: c,
	}

	return g
}

func (recv *TreeIter) ToC() unsafe.Pointer {
	recv.native.stamp =
		(C.gint)(recv.Stamp)

	return (unsafe.Pointer)(recv.native)
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
