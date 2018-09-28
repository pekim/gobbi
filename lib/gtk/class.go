// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
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

	return g
}

func (recv *AboutDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_about_dialog_add_credit_section : unsupported parameter people : no param type

// Unsupported : gtk_about_dialog_get_artists : no return type

// Unsupported : gtk_about_dialog_get_authors : no return type

// Unsupported : gtk_about_dialog_get_documenters : no return type

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists : no param type

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors : no param type

// Unsupported : gtk_about_dialog_set_comments : no return generator

// Unsupported : gtk_about_dialog_set_copyright : no return generator

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters : no param type

// Unsupported : gtk_about_dialog_set_license : no return generator

// Unsupported : gtk_about_dialog_set_license_type : no return generator

// Unsupported : gtk_about_dialog_set_logo : no return generator

// Unsupported : gtk_about_dialog_set_logo_icon_name : no return generator

// Unsupported : gtk_about_dialog_set_program_name : no return generator

// Unsupported : gtk_about_dialog_set_translator_credits : no return generator

// Unsupported : gtk_about_dialog_set_version : no return generator

// Unsupported : gtk_about_dialog_set_website : no return generator

// Unsupported : gtk_about_dialog_set_website_label : no return generator

// Unsupported : gtk_about_dialog_set_wrap_license : no return generator

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

	return g
}

func (recv *AccelGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelGroupNew is a wrapper around the C function gtk_accel_group_new.
func AccelGroupNew() *AccelGroup {
	retC := C.gtk_accel_group_new()
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function gtk_accel_group_activate.
func (recv *AccelGroup) Activate(accelQuark glib.Quark, acceleratable *gobject.Object, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_quark := (C.GQuark)(accelQuark)

	c_acceleratable := (*C.GObject)(acceleratable.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_group_activate((*C.GtkAccelGroup)(recv.native), c_accel_quark, c_acceleratable, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_accel_group_connect : no return generator

// Unsupported : gtk_accel_group_connect_by_path : no return generator

// Disconnect is a wrapper around the C function gtk_accel_group_disconnect.
func (recv *AccelGroup) Disconnect(closure *gobject.Closure) bool {
	c_closure := (*C.GClosure)(closure.ToC())

	retC := C.gtk_accel_group_disconnect((*C.GtkAccelGroup)(recv.native), c_closure)
	retGo := retC == C.TRUE

	return retGo
}

// DisconnectKey is a wrapper around the C function gtk_accel_group_disconnect_key.
func (recv *AccelGroup) DisconnectKey(accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_group_disconnect_key((*C.GtkAccelGroup)(recv.native), c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_accel_group_find : unsupported parameter find_func : no type generator for AccelGroupFindFunc, GtkAccelGroupFindFunc

// Unsupported : gtk_accel_group_lock : no return generator

// Unsupported : gtk_accel_group_query : unsupported parameter n_entries : no type generator for guint, guint*

// Unsupported : gtk_accel_group_unlock : no return generator

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

	return g
}

func (recv *AccelLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AccelLabelNew is a wrapper around the C function gtk_accel_label_new.
func AccelLabelNew(string string) *Widget {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_accel_label_new(c_string)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_key : no type generator for guint, guint*

// GetAccelWidget is a wrapper around the C function gtk_accel_label_get_accel_widget.
func (recv *AccelLabel) GetAccelWidget() *Widget {
	retC := C.gtk_accel_label_get_accel_widget((*C.GtkAccelLabel)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAccelWidth is a wrapper around the C function gtk_accel_label_get_accel_width.
func (recv *AccelLabel) GetAccelWidth() uint32 {
	retC := C.gtk_accel_label_get_accel_width((*C.GtkAccelLabel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Refetch is a wrapper around the C function gtk_accel_label_refetch.
func (recv *AccelLabel) Refetch() bool {
	retC := C.gtk_accel_label_refetch((*C.GtkAccelLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_accel_label_set_accel : no return generator

// Unsupported : gtk_accel_label_set_accel_closure : no return generator

// Unsupported : gtk_accel_label_set_accel_widget : no return generator

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

	return g
}

func (recv *Accessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_accessible_connect_widget_destroyed : no return generator

// Unsupported : gtk_accessible_set_widget : no return generator

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

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_action_activate : no return generator

// Unsupported : gtk_action_block_activate : no return generator

// Unsupported : gtk_action_connect_accelerator : no return generator

// Unsupported : gtk_action_create_icon : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_action_disconnect_accelerator : no return generator

// Unsupported : gtk_action_get_gicon : no return generator

// Unsupported : gtk_action_set_accel_group : no return generator

// Unsupported : gtk_action_set_accel_path : no return generator

// Unsupported : gtk_action_set_always_show_image : no return generator

// Unsupported : gtk_action_set_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_action_set_icon_name : no return generator

// Unsupported : gtk_action_set_is_important : no return generator

// Unsupported : gtk_action_set_label : no return generator

// Unsupported : gtk_action_set_sensitive : no return generator

// Unsupported : gtk_action_set_short_label : no return generator

// Unsupported : gtk_action_set_stock_id : no return generator

// Unsupported : gtk_action_set_tooltip : no return generator

// Unsupported : gtk_action_set_visible : no return generator

// Unsupported : gtk_action_set_visible_horizontal : no return generator

// Unsupported : gtk_action_set_visible_vertical : no return generator

// Unsupported : gtk_action_unblock_activate : no return generator

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

	return g
}

func (recv *ActionBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_action_bar_pack_end : no return generator

// Unsupported : gtk_action_bar_pack_start : no return generator

// Unsupported : gtk_action_bar_set_center_widget : no return generator

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

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_action_group_add_action : no return generator

// Unsupported : gtk_action_group_add_action_with_accel : no return generator

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_remove_action : no return generator

// Unsupported : gtk_action_group_set_accel_group : no return generator

// Unsupported : gtk_action_group_set_sensitive : no return generator

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// Unsupported : gtk_action_group_set_translation_domain : no return generator

// Unsupported : gtk_action_group_set_visible : no return generator

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

// Unsupported : gtk_adjustment_changed : no return generator

// Unsupported : gtk_adjustment_clamp_page : no return generator

// Unsupported : gtk_adjustment_configure : no return generator

// GetValue is a wrapper around the C function gtk_adjustment_get_value.
func (recv *Adjustment) GetValue() float64 {
	retC := C.gtk_adjustment_get_value((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Unsupported : gtk_adjustment_set_lower : no return generator

// Unsupported : gtk_adjustment_set_page_increment : no return generator

// Unsupported : gtk_adjustment_set_page_size : no return generator

// Unsupported : gtk_adjustment_set_step_increment : no return generator

// Unsupported : gtk_adjustment_set_upper : no return generator

// Unsupported : gtk_adjustment_set_value : no return generator

// Unsupported : gtk_adjustment_value_changed : no return generator

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

	return g
}

func (recv *Alignment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AlignmentNew is a wrapper around the C function gtk_alignment_new.
func AlignmentNew(xalign float32, yalign float32, xscale float32, yscale float32) *Widget {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	retC := C.gtk_alignment_new(c_xalign, c_yalign, c_xscale, c_yscale)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_alignment_get_padding : unsupported parameter padding_top : no type generator for guint, guint*

// Unsupported : gtk_alignment_set : no return generator

// Unsupported : gtk_alignment_set_padding : no return generator

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

	return g
}

func (recv *AppChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_app_chooser_button_append_custom_item : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_app_chooser_button_append_separator : no return generator

// GetHeading is a wrapper around the C function gtk_app_chooser_button_get_heading.
func (recv *AppChooserButton) GetHeading() string {
	retC := C.gtk_app_chooser_button_get_heading((*C.GtkAppChooserButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_app_chooser_button_set_active_custom_item : no return generator

// Unsupported : gtk_app_chooser_button_set_heading : no return generator

// Unsupported : gtk_app_chooser_button_set_show_default_item : no return generator

// Unsupported : gtk_app_chooser_button_set_show_dialog_item : no return generator

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

	return g
}

func (recv *AppChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// GetHeading is a wrapper around the C function gtk_app_chooser_dialog_get_heading.
func (recv *AppChooserDialog) GetHeading() string {
	retC := C.gtk_app_chooser_dialog_get_heading((*C.GtkAppChooserDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_app_chooser_dialog_set_heading : no return generator

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

	return g
}

func (recv *AppChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_app_chooser_widget_set_default_text : no return generator

// Unsupported : gtk_app_chooser_widget_set_show_all : no return generator

// Unsupported : gtk_app_chooser_widget_set_show_default : no return generator

// Unsupported : gtk_app_chooser_widget_set_show_fallback : no return generator

// Unsupported : gtk_app_chooser_widget_set_show_other : no return generator

// Unsupported : gtk_app_chooser_widget_set_show_recommended : no return generator

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

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_add_window : no return generator

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_get_actions_for_accel : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_remove_window : no return generator

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// Unsupported : gtk_application_set_app_menu : no return generator

// Unsupported : gtk_application_set_menubar : no return generator

// Unsupported : gtk_application_uninhibit : no return generator

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

	return g
}

func (recv *ApplicationWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_application_window_set_help_overlay : no return generator

// Unsupported : gtk_application_window_set_show_menubar : no return generator

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

	return g
}

func (recv *Arrow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ArrowNew is a wrapper around the C function gtk_arrow_new.
func ArrowNew(arrowType ArrowType, shadowType ShadowType) *Widget {
	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	retC := C.gtk_arrow_new(c_arrow_type, c_shadow_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_arrow_set : no return generator

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

	return g
}

func (recv *AspectFrame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AspectFrameNew is a wrapper around the C function gtk_aspect_frame_new.
func AspectFrameNew(label string, xalign float32, yalign float32, ratio float32, obeyChild bool) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obey_child :=
		boolToGboolean(obeyChild)

	retC := C.gtk_aspect_frame_new(c_label, c_xalign, c_yalign, c_ratio, c_obey_child)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_aspect_frame_set : no return generator

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

	return g
}

func (recv *Assistant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_assistant_add_action_widget : no return generator

// Unsupported : gtk_assistant_commit : no return generator

// Unsupported : gtk_assistant_next_page : no return generator

// Unsupported : gtk_assistant_previous_page : no return generator

// Unsupported : gtk_assistant_remove_action_widget : no return generator

// Unsupported : gtk_assistant_remove_page : no return generator

// Unsupported : gtk_assistant_set_current_page : no return generator

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// Unsupported : gtk_assistant_set_page_complete : no return generator

// Unsupported : gtk_assistant_set_page_has_padding : no return generator

// Unsupported : gtk_assistant_set_page_header_image : no return generator

// Unsupported : gtk_assistant_set_page_side_image : no return generator

// Unsupported : gtk_assistant_set_page_title : no return generator

// Unsupported : gtk_assistant_set_page_type : no return generator

// Unsupported : gtk_assistant_update_buttons_state : no return generator

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

	return g
}

func (recv *Bin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetChild is a wrapper around the C function gtk_bin_get_child.
func (recv *Bin) GetChild() *Widget {
	retC := C.gtk_bin_get_child((*C.GtkBin)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
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

	return g
}

func (recv *Box) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetHomogeneous is a wrapper around the C function gtk_box_get_homogeneous.
func (recv *Box) GetHomogeneous() bool {
	retC := C.gtk_box_get_homogeneous((*C.GtkBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_box_get_spacing.
func (recv *Box) GetSpacing() int32 {
	retC := C.gtk_box_get_spacing((*C.GtkBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_box_pack_end : no return generator

// Unsupported : gtk_box_pack_start : no return generator

// Unsupported : gtk_box_query_child_packing : unsupported parameter padding : no type generator for guint, guint*

// Unsupported : gtk_box_reorder_child : no return generator

// Unsupported : gtk_box_set_baseline_position : no return generator

// Unsupported : gtk_box_set_center_widget : no return generator

// Unsupported : gtk_box_set_child_packing : no return generator

// Unsupported : gtk_box_set_homogeneous : no return generator

// Unsupported : gtk_box_set_spacing : no return generator

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

	return g
}

func (recv *Builder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_objects_from_file : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_string : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_connect_signals : no return generator

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// Unsupported : gtk_builder_expose_object : no return generator

// Unsupported : gtk_builder_extend_with_template : unsupported parameter template_type : no type generator for GType, GType

// Unsupported : gtk_builder_get_type_from_name : no return generator

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// Unsupported : gtk_builder_set_application : no return generator

// Unsupported : gtk_builder_set_translation_domain : no return generator

// Unsupported : gtk_builder_value_from_string_type : unsupported parameter type : no type generator for GType, GType

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

	return g
}

func (recv *Button) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonNew is a wrapper around the C function gtk_button_new.
func ButtonNew() *Widget {
	retC := C.gtk_button_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// ButtonNewFromStock is a wrapper around the C function gtk_button_new_from_stock.
func ButtonNewFromStock(stockId string) *Widget {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_button_new_from_stock(c_stock_id)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithLabel is a wrapper around the C function gtk_button_new_with_label.
func ButtonNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithMnemonic is a wrapper around the C function gtk_button_new_with_mnemonic.
func ButtonNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_button_clicked : no return generator

// Unsupported : gtk_button_enter : no return generator

// Unsupported : gtk_button_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// GetLabel is a wrapper around the C function gtk_button_get_label.
func (recv *Button) GetLabel() string {
	retC := C.gtk_button_get_label((*C.GtkButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetRelief is a wrapper around the C function gtk_button_get_relief.
func (recv *Button) GetRelief() ReliefStyle {
	retC := C.gtk_button_get_relief((*C.GtkButton)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetUseStock is a wrapper around the C function gtk_button_get_use_stock.
func (recv *Button) GetUseStock() bool {
	retC := C.gtk_button_get_use_stock((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_button_get_use_underline.
func (recv *Button) GetUseUnderline() bool {
	retC := C.gtk_button_get_use_underline((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_button_leave : no return generator

// Unsupported : gtk_button_pressed : no return generator

// Unsupported : gtk_button_released : no return generator

// Unsupported : gtk_button_set_alignment : no return generator

// Unsupported : gtk_button_set_always_show_image : no return generator

// Unsupported : gtk_button_set_focus_on_click : no return generator

// Unsupported : gtk_button_set_image : no return generator

// Unsupported : gtk_button_set_image_position : no return generator

// Unsupported : gtk_button_set_label : no return generator

// Unsupported : gtk_button_set_relief : no return generator

// Unsupported : gtk_button_set_use_stock : no return generator

// Unsupported : gtk_button_set_use_underline : no return generator

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

	return g
}

func (recv *ButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetLayout is a wrapper around the C function gtk_button_box_get_layout.
func (recv *ButtonBox) GetLayout() ButtonBoxStyle {
	retC := C.gtk_button_box_get_layout((*C.GtkButtonBox)(recv.native))
	retGo := (ButtonBoxStyle)(retC)

	return retGo
}

// Unsupported : gtk_button_box_set_child_non_homogeneous : no return generator

// Unsupported : gtk_button_box_set_child_secondary : no return generator

// Unsupported : gtk_button_box_set_layout : no return generator

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

	return g
}

func (recv *Calendar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CalendarNew is a wrapper around the C function gtk_calendar_new.
func CalendarNew() *Widget {
	retC := C.gtk_calendar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_calendar_clear_marks : no return generator

// Unsupported : gtk_calendar_get_date : unsupported parameter year : no type generator for guint, guint*

// Unsupported : gtk_calendar_mark_day : no return generator

// Unsupported : gtk_calendar_select_day : no return generator

// Unsupported : gtk_calendar_select_month : no return generator

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc, GtkCalendarDetailFunc

// Unsupported : gtk_calendar_set_detail_height_rows : no return generator

// Unsupported : gtk_calendar_set_detail_width_chars : no return generator

// Unsupported : gtk_calendar_set_display_options : no return generator

// Unsupported : gtk_calendar_unmark_day : no return generator

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

	return g
}

func (recv *CellArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_cell_area_activate : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_add : no return generator

// Unsupported : gtk_cell_area_add_focus_sibling : no return generator

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_apply_attributes : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_cell_area_attribute_connect : no return generator

// Unsupported : gtk_cell_area_attribute_disconnect : no return generator

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_get_property : no return generator

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_set_property : no return generator

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback, GtkCellCallback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_allocation : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_at_position : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_edit_widget : no return generator

// Unsupported : gtk_cell_area_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_inner_cell_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_remove : no return generator

// Unsupported : gtk_cell_area_remove_focus_sibling : no return generator

// Unsupported : gtk_cell_area_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_request_renderer : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_area_set_focus_cell : no return generator

// Unsupported : gtk_cell_area_stop_editing : no return generator

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

	return g
}

func (recv *CellAreaBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_cell_area_box_pack_end : no return generator

// Unsupported : gtk_cell_area_box_pack_start : no return generator

// Unsupported : gtk_cell_area_box_set_spacing : no return generator

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

	return g
}

func (recv *CellAreaContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_cell_area_context_allocate : no return generator

// Unsupported : gtk_cell_area_context_get_allocation : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_push_preferred_height : no return generator

// Unsupported : gtk_cell_area_context_push_preferred_width : no return generator

// Unsupported : gtk_cell_area_context_reset : no return generator

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

	return g
}

func (recv *CellRenderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_cell_renderer_activate : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_cell_renderer_get_fixed_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_size : no return generator

// Unsupported : gtk_cell_renderer_get_preferred_width : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_set_alignment : no return generator

// Unsupported : gtk_cell_renderer_set_fixed_size : no return generator

// Unsupported : gtk_cell_renderer_set_padding : no return generator

// Unsupported : gtk_cell_renderer_set_sensitive : no return generator

// Unsupported : gtk_cell_renderer_set_visible : no return generator

// Unsupported : gtk_cell_renderer_start_editing : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_renderer_stop_editing : no return generator

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

	return g
}

func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererPixbufNew is a wrapper around the C function gtk_cell_renderer_pixbuf_new.
func CellRendererPixbufNew() *CellRenderer {
	retC := C.gtk_cell_renderer_pixbuf_new()
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *CellRendererText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererTextNew is a wrapper around the C function gtk_cell_renderer_text_new.
func CellRendererTextNew() *CellRenderer {
	retC := C.gtk_cell_renderer_text_new()
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_renderer_text_set_fixed_height_from_font : no return generator

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

	return g
}

func (recv *CellRendererToggle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererToggleNew is a wrapper around the C function gtk_cell_renderer_toggle_new.
func CellRendererToggleNew() *CellRenderer {
	retC := C.gtk_cell_renderer_toggle_new()
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_cell_renderer_toggle_get_active.
func (recv *CellRendererToggle) GetActive() bool {
	retC := C.gtk_cell_renderer_toggle_get_active((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRadio is a wrapper around the C function gtk_cell_renderer_toggle_get_radio.
func (recv *CellRendererToggle) GetRadio() bool {
	retC := C.gtk_cell_renderer_toggle_get_radio((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_renderer_toggle_set_activatable : no return generator

// Unsupported : gtk_cell_renderer_toggle_set_active : no return generator

// Unsupported : gtk_cell_renderer_toggle_set_radio : no return generator

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

	return g
}

func (recv *CellView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_cell_view_get_model : no return generator

// Unsupported : gtk_cell_view_set_background_color : no return generator

// Unsupported : gtk_cell_view_set_background_rgba : no return generator

// Unsupported : gtk_cell_view_set_displayed_row : no return generator

// Unsupported : gtk_cell_view_set_draw_sensitive : no return generator

// Unsupported : gtk_cell_view_set_fit_model : no return generator

// Unsupported : gtk_cell_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

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

	return g
}

func (recv *CheckButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckButtonNew is a wrapper around the C function gtk_check_button_new.
func CheckButtonNew() *Widget {
	retC := C.gtk_check_button_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithLabel is a wrapper around the C function gtk_check_button_new_with_label.
func CheckButtonNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithMnemonic is a wrapper around the C function gtk_check_button_new_with_mnemonic.
func CheckButtonNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *CheckMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemNew is a wrapper around the C function gtk_check_menu_item_new.
func CheckMenuItemNew() *Widget {
	retC := C.gtk_check_menu_item_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithLabel is a wrapper around the C function gtk_check_menu_item_new_with_label.
func CheckMenuItemNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithMnemonic is a wrapper around the C function gtk_check_menu_item_new_with_mnemonic.
func CheckMenuItemNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_check_menu_item_get_active.
func (recv *CheckMenuItem) GetActive() bool {
	retC := C.gtk_check_menu_item_get_active((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInconsistent is a wrapper around the C function gtk_check_menu_item_get_inconsistent.
func (recv *CheckMenuItem) GetInconsistent() bool {
	retC := C.gtk_check_menu_item_get_inconsistent((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_check_menu_item_set_active : no return generator

// Unsupported : gtk_check_menu_item_set_draw_as_radio : no return generator

// Unsupported : gtk_check_menu_item_set_inconsistent : no return generator

// Unsupported : gtk_check_menu_item_toggled : no return generator

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

	return g
}

func (recv *Clipboard) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_clipboard_clear : no return generator

// GetOwner is a wrapper around the C function gtk_clipboard_get_owner.
func (recv *Clipboard) GetOwner() *gobject.Object {
	retC := C.gtk_clipboard_get_owner((*C.GtkClipboard)(recv.native))
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_clipboard_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc, GtkClipboardImageReceivedFunc

// Unsupported : gtk_clipboard_request_rich_text : unsupported parameter callback : no type generator for ClipboardRichTextReceivedFunc, GtkClipboardRichTextReceivedFunc

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc, GtkClipboardTargetsReceivedFunc

// Unsupported : gtk_clipboard_request_text : unsupported parameter callback : no type generator for ClipboardTextReceivedFunc, GtkClipboardTextReceivedFunc

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc, GtkClipboardURIReceivedFunc

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_image : no return generator

// Unsupported : gtk_clipboard_set_text : no return generator

// Unsupported : gtk_clipboard_set_with_data : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_with_owner : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_store : no return generator

// Unsupported : gtk_clipboard_wait_for_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_wait_for_rich_text : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : no param type

// WaitForText is a wrapper around the C function gtk_clipboard_wait_for_text.
func (recv *Clipboard) WaitForText() string {
	retC := C.gtk_clipboard_wait_for_text((*C.GtkClipboard)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_clipboard_wait_for_uris : no return type

// Unsupported : gtk_clipboard_wait_is_target_available : unsupported parameter target : Blacklisted record : GdkAtom

// WaitIsTextAvailable is a wrapper around the C function gtk_clipboard_wait_is_text_available.
func (recv *Clipboard) WaitIsTextAvailable() bool {
	retC := C.gtk_clipboard_wait_is_text_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
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

	return g
}

func (recv *ColorButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_color_button_get_color : no return generator

// Unsupported : gtk_color_button_get_rgba : no return generator

// Unsupported : gtk_color_button_set_alpha : no return generator

// Unsupported : gtk_color_button_set_color : no return generator

// Unsupported : gtk_color_button_set_rgba : no return generator

// Unsupported : gtk_color_button_set_title : no return generator

// Unsupported : gtk_color_button_set_use_alpha : no return generator

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

	return g
}

func (recv *ColorSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionNew is a wrapper around the C function gtk_color_selection_new.
func ColorSelectionNew() *Widget {
	retC := C.gtk_color_selection_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentAlpha is a wrapper around the C function gtk_color_selection_get_current_alpha.
func (recv *ColorSelection) GetCurrentAlpha() uint16 {
	retC := C.gtk_color_selection_get_current_alpha((*C.GtkColorSelection)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported : gtk_color_selection_get_current_color : no return generator

// Unsupported : gtk_color_selection_get_current_rgba : no return generator

// GetHasOpacityControl is a wrapper around the C function gtk_color_selection_get_has_opacity_control.
func (recv *ColorSelection) GetHasOpacityControl() bool {
	retC := C.gtk_color_selection_get_has_opacity_control((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasPalette is a wrapper around the C function gtk_color_selection_get_has_palette.
func (recv *ColorSelection) GetHasPalette() bool {
	retC := C.gtk_color_selection_get_has_palette((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPreviousAlpha is a wrapper around the C function gtk_color_selection_get_previous_alpha.
func (recv *ColorSelection) GetPreviousAlpha() uint16 {
	retC := C.gtk_color_selection_get_previous_alpha((*C.GtkColorSelection)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported : gtk_color_selection_get_previous_color : no return generator

// Unsupported : gtk_color_selection_get_previous_rgba : no return generator

// IsAdjusting is a wrapper around the C function gtk_color_selection_is_adjusting.
func (recv *ColorSelection) IsAdjusting() bool {
	retC := C.gtk_color_selection_is_adjusting((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_color_selection_set_current_alpha : no return generator

// Unsupported : gtk_color_selection_set_current_color : no return generator

// Unsupported : gtk_color_selection_set_current_rgba : no return generator

// Unsupported : gtk_color_selection_set_has_opacity_control : no return generator

// Unsupported : gtk_color_selection_set_has_palette : no return generator

// Unsupported : gtk_color_selection_set_previous_alpha : no return generator

// Unsupported : gtk_color_selection_set_previous_color : no return generator

// Unsupported : gtk_color_selection_set_previous_rgba : no return generator

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

	return g
}

func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ColorSelectionDialogNew is a wrapper around the C function gtk_color_selection_dialog_new.
func ColorSelectionDialogNew(title string) *Widget {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_color_selection_dialog_new(c_title)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *ComboBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBoxNewWithArea is a wrapper around the C function gtk_combo_box_new_with_area.
func ComboBoxNewWithArea(area *CellArea) *Widget {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_combo_box_new_with_area(c_area)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithAreaAndEntry is a wrapper around the C function gtk_combo_box_new_with_area_and_entry.
func ComboBoxNewWithAreaAndEntry(area *CellArea) *Widget {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_combo_box_new_with_area_and_entry(c_area)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetAddTearoffs is a wrapper around the C function gtk_combo_box_get_add_tearoffs.
func (recv *ComboBox) GetAddTearoffs() bool {
	retC := C.gtk_combo_box_get_add_tearoffs((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_combo_box_get_model : no return generator

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// Unsupported : gtk_combo_box_popdown : no return generator

// Unsupported : gtk_combo_box_popup : no return generator

// Unsupported : gtk_combo_box_popup_for_device : no return generator

// Unsupported : gtk_combo_box_set_active : no return generator

// Unsupported : gtk_combo_box_set_active_iter : no return generator

// Unsupported : gtk_combo_box_set_add_tearoffs : no return generator

// Unsupported : gtk_combo_box_set_button_sensitivity : no return generator

// Unsupported : gtk_combo_box_set_column_span_column : no return generator

// Unsupported : gtk_combo_box_set_entry_text_column : no return generator

// Unsupported : gtk_combo_box_set_focus_on_click : no return generator

// Unsupported : gtk_combo_box_set_id_column : no return generator

// Unsupported : gtk_combo_box_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_set_popup_fixed_width : no return generator

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_combo_box_set_row_span_column : no return generator

// Unsupported : gtk_combo_box_set_title : no return generator

// Unsupported : gtk_combo_box_set_wrap_width : no return generator

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

	return g
}

func (recv *ComboBoxText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_combo_box_text_append : no return generator

// Unsupported : gtk_combo_box_text_append_text : no return generator

// Unsupported : gtk_combo_box_text_insert : no return generator

// Unsupported : gtk_combo_box_text_insert_text : no return generator

// Unsupported : gtk_combo_box_text_prepend : no return generator

// Unsupported : gtk_combo_box_text_prepend_text : no return generator

// Unsupported : gtk_combo_box_text_remove : no return generator

// Unsupported : gtk_combo_box_text_remove_all : no return generator

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

	return g
}

func (recv *Container) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_container_add : no return generator

// Unsupported : gtk_container_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_container_check_resize : no return generator

// Unsupported : gtk_container_child_get : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_get_property : no return generator

// Unsupported : gtk_container_child_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_notify : no return generator

// Unsupported : gtk_container_child_notify_by_pspec : no return generator

// Unsupported : gtk_container_child_set : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_set_property : no return generator

// Unsupported : gtk_container_child_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_type : no return generator

// Unsupported : gtk_container_forall : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_foreach : unsupported parameter callback : no type generator for Callback, GtkCallback

// GetBorderWidth is a wrapper around the C function gtk_container_get_border_width.
func (recv *Container) GetBorderWidth() uint32 {
	retC := C.gtk_container_get_border_width((*C.GtkContainer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetChildren is a wrapper around the C function gtk_container_get_children.
func (recv *Container) GetChildren() *glib.List {
	retC := C.gtk_container_get_children((*C.GtkContainer)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_container_get_focus_chain : unsupported parameter focusable_widgets : record with indirection level of 2

// GetFocusHadjustment is a wrapper around the C function gtk_container_get_focus_hadjustment.
func (recv *Container) GetFocusHadjustment() *Adjustment {
	retC := C.gtk_container_get_focus_hadjustment((*C.GtkContainer)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusVadjustment is a wrapper around the C function gtk_container_get_focus_vadjustment.
func (recv *Container) GetFocusVadjustment() *Adjustment {
	retC := C.gtk_container_get_focus_vadjustment((*C.GtkContainer)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPathForChild is a wrapper around the C function gtk_container_get_path_for_child.
func (recv *Container) GetPathForChild(child *Widget) *WidgetPath {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_container_get_path_for_child((*C.GtkContainer)(recv.native), c_child)
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResizeMode is a wrapper around the C function gtk_container_get_resize_mode.
func (recv *Container) GetResizeMode() ResizeMode {
	retC := C.gtk_container_get_resize_mode((*C.GtkContainer)(recv.native))
	retGo := (ResizeMode)(retC)

	return retGo
}

// Unsupported : gtk_container_propagate_draw : no return generator

// Unsupported : gtk_container_remove : no return generator

// Unsupported : gtk_container_resize_children : no return generator

// Unsupported : gtk_container_set_border_width : no return generator

// Unsupported : gtk_container_set_focus_chain : no return generator

// Unsupported : gtk_container_set_focus_child : no return generator

// Unsupported : gtk_container_set_focus_hadjustment : no return generator

// Unsupported : gtk_container_set_focus_vadjustment : no return generator

// Unsupported : gtk_container_set_reallocate_redraws : no return generator

// Unsupported : gtk_container_set_resize_mode : no return generator

// Unsupported : gtk_container_unset_focus_chain : no return generator

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

	return g
}

func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerCellAccessibleNew is a wrapper around the C function gtk_container_cell_accessible_new.
func ContainerCellAccessibleNew() *ContainerCellAccessible {
	retC := C.gtk_container_cell_accessible_new()
	retGo := ContainerCellAccessibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_container_cell_accessible_add_child : no return generator

// GetChildren is a wrapper around the C function gtk_container_cell_accessible_get_children.
func (recv *ContainerCellAccessible) GetChildren() *glib.List {
	retC := C.gtk_container_cell_accessible_get_children((*C.GtkContainerCellAccessible)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_container_cell_accessible_remove_child : no return generator

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

	return g
}

func (recv *CssProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CssProviderNew is a wrapper around the C function gtk_css_provider_new.
func CssProviderNew() *CssProvider {
	retC := C.gtk_css_provider_new()
	retGo := CssProviderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_css_provider_load_from_data : unsupported parameter data : no param type

// Unsupported : gtk_css_provider_load_from_file : unsupported parameter file : no type generator for Gio.File, GFile*

// LoadFromPath is a wrapper around the C function gtk_css_provider_load_from_path.
func (recv *CssProvider) LoadFromPath(path string) (bool, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	var cThrowableError *C.GError

	retC := C.gtk_css_provider_load_from_path((*C.GtkCssProvider)(recv.native), c_path, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_css_provider_load_from_resource : no return generator

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

	return g
}

func (recv *Dialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DialogNew is a wrapper around the C function gtk_dialog_new.
func DialogNew() *Widget {
	retC := C.gtk_dialog_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_add_action_widget : no return generator

// AddButton is a wrapper around the C function gtk_dialog_add_button.
func (recv *Dialog) AddButton(buttonText string, responseId int32) *Widget {
	c_button_text := C.CString(buttonText)
	defer C.free(unsafe.Pointer(c_button_text))

	c_response_id := (C.gint)(responseId)

	retC := C.gtk_dialog_add_button((*C.GtkDialog)(recv.native), c_button_text, c_response_id)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_add_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_response : no return generator

// Run is a wrapper around the C function gtk_dialog_run.
func (recv *Dialog) Run() int32 {
	retC := C.gtk_dialog_run((*C.GtkDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order_from_array : unsupported parameter new_order : no param type

// Unsupported : gtk_dialog_set_default_response : no return generator

// Unsupported : gtk_dialog_set_response_sensitive : no return generator

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

	return g
}

func (recv *DrawingArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingAreaNew is a wrapper around the C function gtk_drawing_area_new.
func DrawingAreaNew() *Widget {
	retC := C.gtk_drawing_area_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *Entry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryNew is a wrapper around the C function gtk_entry_new.
func EntryNew() *Widget {
	retC := C.gtk_entry_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActivatesDefault is a wrapper around the C function gtk_entry_get_activates_default.
func (recv *Entry) GetActivatesDefault() bool {
	retC := C.gtk_entry_get_activates_default((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasFrame is a wrapper around the C function gtk_entry_get_has_frame.
func (recv *Entry) GetHasFrame() bool {
	retC := C.gtk_entry_get_has_frame((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_icon_gicon : no return generator

// GetInvisibleChar is a wrapper around the C function gtk_entry_get_invisible_char.
func (recv *Entry) GetInvisibleChar() rune {
	retC := C.gtk_entry_get_invisible_char((*C.GtkEntry)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetLayout is a wrapper around the C function gtk_entry_get_layout.
func (recv *Entry) GetLayout() *pango.Layout {
	retC := C.gtk_entry_get_layout((*C.GtkEntry)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_entry_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetMaxLength is a wrapper around the C function gtk_entry_get_max_length.
func (recv *Entry) GetMaxLength() int32 {
	retC := C.gtk_entry_get_max_length((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_entry_get_text.
func (recv *Entry) GetText() string {
	retC := C.gtk_entry_get_text((*C.GtkEntry)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// GetVisibility is a wrapper around the C function gtk_entry_get_visibility.
func (recv *Entry) GetVisibility() bool {
	retC := C.gtk_entry_get_visibility((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidthChars is a wrapper around the C function gtk_entry_get_width_chars.
func (recv *Entry) GetWidthChars() int32 {
	retC := C.gtk_entry_get_width_chars((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_entry_grab_focus_without_selecting : no return generator

// LayoutIndexToTextIndex is a wrapper around the C function gtk_entry_layout_index_to_text_index.
func (recv *Entry) LayoutIndexToTextIndex(layoutIndex int32) int32 {
	c_layout_index := (C.gint)(layoutIndex)

	retC := C.gtk_entry_layout_index_to_text_index((*C.GtkEntry)(recv.native), c_layout_index)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_entry_progress_pulse : no return generator

// Unsupported : gtk_entry_reset_im_context : no return generator

// Unsupported : gtk_entry_set_activates_default : no return generator

// Unsupported : gtk_entry_set_alignment : no return generator

// Unsupported : gtk_entry_set_attributes : no return generator

// Unsupported : gtk_entry_set_buffer : no return generator

// Unsupported : gtk_entry_set_completion : no return generator

// Unsupported : gtk_entry_set_cursor_hadjustment : no return generator

// Unsupported : gtk_entry_set_has_frame : no return generator

// Unsupported : gtk_entry_set_icon_activatable : no return generator

// Unsupported : gtk_entry_set_icon_drag_source : no return generator

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_entry_set_icon_from_icon_name : no return generator

// Unsupported : gtk_entry_set_icon_from_pixbuf : no return generator

// Unsupported : gtk_entry_set_icon_from_stock : no return generator

// Unsupported : gtk_entry_set_icon_sensitive : no return generator

// Unsupported : gtk_entry_set_icon_tooltip_markup : no return generator

// Unsupported : gtk_entry_set_icon_tooltip_text : no return generator

// Unsupported : gtk_entry_set_inner_border : no return generator

// Unsupported : gtk_entry_set_input_hints : no return generator

// Unsupported : gtk_entry_set_input_purpose : no return generator

// Unsupported : gtk_entry_set_invisible_char : no return generator

// Unsupported : gtk_entry_set_max_length : no return generator

// Unsupported : gtk_entry_set_max_width_chars : no return generator

// Unsupported : gtk_entry_set_overwrite_mode : no return generator

// Unsupported : gtk_entry_set_placeholder_text : no return generator

// Unsupported : gtk_entry_set_progress_fraction : no return generator

// Unsupported : gtk_entry_set_progress_pulse_step : no return generator

// Unsupported : gtk_entry_set_tabs : no return generator

// Unsupported : gtk_entry_set_text : no return generator

// Unsupported : gtk_entry_set_visibility : no return generator

// Unsupported : gtk_entry_set_width_chars : no return generator

// TextIndexToLayoutIndex is a wrapper around the C function gtk_entry_text_index_to_layout_index.
func (recv *Entry) TextIndexToLayoutIndex(textIndex int32) int32 {
	c_text_index := (C.gint)(textIndex)

	retC := C.gtk_entry_text_index_to_layout_index((*C.GtkEntry)(recv.native), c_text_index)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_entry_unset_invisible_char : no return generator

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

	return g
}

func (recv *EntryBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_entry_buffer_emit_deleted_text : no return generator

// Unsupported : gtk_entry_buffer_emit_inserted_text : no return generator

// Unsupported : gtk_entry_buffer_set_max_length : no return generator

// Unsupported : gtk_entry_buffer_set_text : no return generator

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

	return g
}

func (recv *EntryCompletion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_entry_completion_complete : no return generator

// Unsupported : gtk_entry_completion_delete_action : no return generator

// Unsupported : gtk_entry_completion_get_model : no return generator

// Unsupported : gtk_entry_completion_insert_action_markup : no return generator

// Unsupported : gtk_entry_completion_insert_action_text : no return generator

// Unsupported : gtk_entry_completion_insert_prefix : no return generator

// Unsupported : gtk_entry_completion_set_inline_completion : no return generator

// Unsupported : gtk_entry_completion_set_inline_selection : no return generator

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc, GtkEntryCompletionMatchFunc

// Unsupported : gtk_entry_completion_set_minimum_key_length : no return generator

// Unsupported : gtk_entry_completion_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_entry_completion_set_popup_completion : no return generator

// Unsupported : gtk_entry_completion_set_popup_set_width : no return generator

// Unsupported : gtk_entry_completion_set_popup_single_match : no return generator

// Unsupported : gtk_entry_completion_set_text_column : no return generator

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

	return g
}

func (recv *EventBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventBoxNew is a wrapper around the C function gtk_event_box_new.
func EventBoxNew() *Widget {
	retC := C.gtk_event_box_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_event_box_set_above_child : no return generator

// Unsupported : gtk_event_box_set_visible_window : no return generator

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

	return g
}

func (recv *EventController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_event_controller_handle_event : unsupported parameter event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_event_controller_reset : no return generator

// Unsupported : gtk_event_controller_set_propagation_phase : no return generator

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

	return g
}

func (recv *Expander) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_expander_set_expanded : no return generator

// Unsupported : gtk_expander_set_label : no return generator

// Unsupported : gtk_expander_set_label_fill : no return generator

// Unsupported : gtk_expander_set_label_widget : no return generator

// Unsupported : gtk_expander_set_resize_toplevel : no return generator

// Unsupported : gtk_expander_set_spacing : no return generator

// Unsupported : gtk_expander_set_use_markup : no return generator

// Unsupported : gtk_expander_set_use_underline : no return generator

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

	return g
}

func (recv *FileChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_file_chooser_button_set_focus_on_click : no return generator

// Unsupported : gtk_file_chooser_button_set_title : no return generator

// Unsupported : gtk_file_chooser_button_set_width_chars : no return generator

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

	return g
}

func (recv *FileChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

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

	return g
}

func (recv *FileFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc, GtkFileFilterFunc

// Unsupported : gtk_file_filter_add_mime_type : no return generator

// Unsupported : gtk_file_filter_add_pattern : no return generator

// Unsupported : gtk_file_filter_add_pixbuf_formats : no return generator

// Unsupported : gtk_file_filter_set_name : no return generator

// Unsupported : gtk_file_filter_to_gvariant : return type : Blacklisted record : GVariant

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

	return g
}

func (recv *Fixed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FixedNew is a wrapper around the C function gtk_fixed_new.
func FixedNew() *Widget {
	retC := C.gtk_fixed_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_fixed_move : no return generator

// Unsupported : gtk_fixed_put : no return generator

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

	return g
}

func (recv *FlowBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_flow_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_flow_box_insert : no return generator

// Unsupported : gtk_flow_box_invalidate_filter : no return generator

// Unsupported : gtk_flow_box_invalidate_sort : no return generator

// Unsupported : gtk_flow_box_select_all : no return generator

// Unsupported : gtk_flow_box_select_child : no return generator

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc, GtkFlowBoxForeachFunc

// Unsupported : gtk_flow_box_set_activate_on_single_click : no return generator

// Unsupported : gtk_flow_box_set_column_spacing : no return generator

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc, GtkFlowBoxFilterFunc

// Unsupported : gtk_flow_box_set_hadjustment : no return generator

// Unsupported : gtk_flow_box_set_homogeneous : no return generator

// Unsupported : gtk_flow_box_set_max_children_per_line : no return generator

// Unsupported : gtk_flow_box_set_min_children_per_line : no return generator

// Unsupported : gtk_flow_box_set_row_spacing : no return generator

// Unsupported : gtk_flow_box_set_selection_mode : no return generator

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc, GtkFlowBoxSortFunc

// Unsupported : gtk_flow_box_set_vadjustment : no return generator

// Unsupported : gtk_flow_box_unselect_all : no return generator

// Unsupported : gtk_flow_box_unselect_child : no return generator

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

	return g
}

func (recv *FlowBoxChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_flow_box_child_changed : no return generator

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

	return g
}

func (recv *FontButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_font_button_set_show_size : no return generator

// Unsupported : gtk_font_button_set_show_style : no return generator

// Unsupported : gtk_font_button_set_title : no return generator

// Unsupported : gtk_font_button_set_use_font : no return generator

// Unsupported : gtk_font_button_set_use_size : no return generator

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

	return g
}

func (recv *FontSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionNew is a wrapper around the C function gtk_font_selection_new.
func FontSelectionNew() *Widget {
	retC := C.gtk_font_selection_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_selection_get_font_name.
func (recv *FontSelection) GetFontName() string {
	retC := C.gtk_font_selection_get_font_name((*C.GtkFontSelection)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewText is a wrapper around the C function gtk_font_selection_get_preview_text.
func (recv *FontSelection) GetPreviewText() string {
	retC := C.gtk_font_selection_get_preview_text((*C.GtkFontSelection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_selection_set_font_name.
func (recv *FontSelection) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_selection_set_font_name((*C.GtkFontSelection)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_font_selection_set_preview_text : no return generator

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

	return g
}

func (recv *FontSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontSelectionDialogNew is a wrapper around the C function gtk_font_selection_dialog_new.
func FontSelectionDialogNew(title string) *Widget {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_font_selection_dialog_new(c_title)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_selection_dialog_get_font_name.
func (recv *FontSelectionDialog) GetFontName() string {
	retC := C.gtk_font_selection_dialog_get_font_name((*C.GtkFontSelectionDialog)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewText is a wrapper around the C function gtk_font_selection_dialog_get_preview_text.
func (recv *FontSelectionDialog) GetPreviewText() string {
	retC := C.gtk_font_selection_dialog_get_preview_text((*C.GtkFontSelectionDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_selection_dialog_set_font_name.
func (recv *FontSelectionDialog) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_selection_dialog_set_font_name((*C.GtkFontSelectionDialog)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_font_selection_dialog_set_preview_text : no return generator

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

	return g
}

func (recv *Frame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameNew is a wrapper around the C function gtk_frame_new.
func FrameNew(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_frame_new(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLabel is a wrapper around the C function gtk_frame_get_label.
func (recv *Frame) GetLabel() string {
	retC := C.gtk_frame_get_label((*C.GtkFrame)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_frame_get_label_align : unsupported parameter xalign : no type generator for gfloat, gfloat*

// GetLabelWidget is a wrapper around the C function gtk_frame_get_label_widget.
func (recv *Frame) GetLabelWidget() *Widget {
	retC := C.gtk_frame_get_label_widget((*C.GtkFrame)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_frame_get_shadow_type.
func (recv *Frame) GetShadowType() ShadowType {
	retC := C.gtk_frame_get_shadow_type((*C.GtkFrame)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// Unsupported : gtk_frame_set_label : no return generator

// Unsupported : gtk_frame_set_label_align : no return generator

// Unsupported : gtk_frame_set_label_widget : no return generator

// Unsupported : gtk_frame_set_shadow_type : no return generator

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

	return g
}

func (recv *Gesture) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_get_bounding_box : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_get_bounding_box_center : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_get_last_event : no return generator

// Unsupported : gtk_gesture_get_point : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_group : no return generator

// Unsupported : gtk_gesture_set_window : no return generator

// Unsupported : gtk_gesture_ungroup : no return generator

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

	return g
}

func (recv *GestureDrag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_drag_get_offset : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_drag_get_start_point : unsupported parameter x : no type generator for gdouble, gdouble*

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

	return g
}

func (recv *GestureMultiPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_multi_press_get_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_multi_press_set_area : unsupported parameter rect : Blacklisted record : GdkRectangle

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

	return g
}

func (recv *GesturePan) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_pan_set_orientation : no return generator

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

	return g
}

func (recv *GestureSingle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_single_set_button : no return generator

// Unsupported : gtk_gesture_single_set_exclusive : no return generator

// Unsupported : gtk_gesture_single_set_touch_only : no return generator

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

	return g
}

func (recv *GestureSwipe) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gesture_swipe_get_velocity : unsupported parameter velocity_x : no type generator for gdouble, gdouble*

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

	return g
}

func (recv *Grid) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GridNew is a wrapper around the C function gtk_grid_new.
func GridNew() *Widget {
	retC := C.gtk_grid_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_grid_attach : no return generator

// Unsupported : gtk_grid_attach_next_to : no return generator

// GetColumnHomogeneous is a wrapper around the C function gtk_grid_get_column_homogeneous.
func (recv *Grid) GetColumnHomogeneous() bool {
	retC := C.gtk_grid_get_column_homogeneous((*C.GtkGrid)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetColumnSpacing is a wrapper around the C function gtk_grid_get_column_spacing.
func (recv *Grid) GetColumnSpacing() uint32 {
	retC := C.gtk_grid_get_column_spacing((*C.GtkGrid)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetRowHomogeneous is a wrapper around the C function gtk_grid_get_row_homogeneous.
func (recv *Grid) GetRowHomogeneous() bool {
	retC := C.gtk_grid_get_row_homogeneous((*C.GtkGrid)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_grid_get_row_spacing.
func (recv *Grid) GetRowSpacing() uint32 {
	retC := C.gtk_grid_get_row_spacing((*C.GtkGrid)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_grid_insert_column : no return generator

// Unsupported : gtk_grid_insert_next_to : no return generator

// Unsupported : gtk_grid_insert_row : no return generator

// Unsupported : gtk_grid_remove_column : no return generator

// Unsupported : gtk_grid_remove_row : no return generator

// Unsupported : gtk_grid_set_baseline_row : no return generator

// Unsupported : gtk_grid_set_column_homogeneous : no return generator

// Unsupported : gtk_grid_set_column_spacing : no return generator

// Unsupported : gtk_grid_set_row_baseline_position : no return generator

// Unsupported : gtk_grid_set_row_homogeneous : no return generator

// Unsupported : gtk_grid_set_row_spacing : no return generator

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

	return g
}

func (recv *HBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HBoxNew is a wrapper around the C function gtk_hbox_new.
func HBoxNew(homogeneous bool, spacing int32) *Widget {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_hbox_new(c_homogeneous, c_spacing)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HButtonBoxNew is a wrapper around the C function gtk_hbutton_box_new.
func HButtonBoxNew() *Widget {
	retC := C.gtk_hbutton_box_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HPanedNew is a wrapper around the C function gtk_hpaned_new.
func HPanedNew() *Widget {
	retC := C.gtk_hpaned_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HSV) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_hsv_get_color : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_metrics : unsupported parameter size : no type generator for gint, gint*

// Unsupported : gtk_hsv_set_color : no return generator

// Unsupported : gtk_hsv_set_metrics : no return generator

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

	return g
}

func (recv *HScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScaleNew is a wrapper around the C function gtk_hscale_new.
func HScaleNew(adjustment *Adjustment) *Widget {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_hscale_new(c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HScaleNewWithRange is a wrapper around the C function gtk_hscale_new_with_range.
func HScaleNewWithRange(min float64, max float64, step float64) *Widget {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_hscale_new_with_range(c_min, c_max, c_step)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HScrollbarNew is a wrapper around the C function gtk_hscrollbar_new.
func HScrollbarNew(adjustment *Adjustment) *Widget {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_hscrollbar_new(c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HSeparatorNew is a wrapper around the C function gtk_hseparator_new.
func HSeparatorNew() *Widget {
	retC := C.gtk_hseparator_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *HandleBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HandleBoxNew is a wrapper around the C function gtk_handle_box_new.
func HandleBoxNew() *Widget {
	retC := C.gtk_handle_box_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHandlePosition is a wrapper around the C function gtk_handle_box_get_handle_position.
func (recv *HandleBox) GetHandlePosition() PositionType {
	retC := C.gtk_handle_box_get_handle_position((*C.GtkHandleBox)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_handle_box_get_shadow_type.
func (recv *HandleBox) GetShadowType() ShadowType {
	retC := C.gtk_handle_box_get_shadow_type((*C.GtkHandleBox)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetSnapEdge is a wrapper around the C function gtk_handle_box_get_snap_edge.
func (recv *HandleBox) GetSnapEdge() PositionType {
	retC := C.gtk_handle_box_get_snap_edge((*C.GtkHandleBox)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// Unsupported : gtk_handle_box_set_handle_position : no return generator

// Unsupported : gtk_handle_box_set_shadow_type : no return generator

// Unsupported : gtk_handle_box_set_snap_edge : no return generator

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

	return g
}

func (recv *HeaderBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_header_bar_pack_end : no return generator

// Unsupported : gtk_header_bar_pack_start : no return generator

// Unsupported : gtk_header_bar_set_custom_title : no return generator

// Unsupported : gtk_header_bar_set_decoration_layout : no return generator

// Unsupported : gtk_header_bar_set_has_subtitle : no return generator

// Unsupported : gtk_header_bar_set_show_close_button : no return generator

// Unsupported : gtk_header_bar_set_subtitle : no return generator

// Unsupported : gtk_header_bar_set_title : no return generator

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

	return g
}

func (recv *IMContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DeleteSurrounding is a wrapper around the C function gtk_im_context_delete_surrounding.
func (recv *IMContext) DeleteSurrounding(offset int32, nChars int32) bool {
	c_offset := (C.gint)(offset)

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_im_context_delete_surrounding((*C.GtkIMContext)(recv.native), c_offset, c_n_chars)
	retGo := retC == C.TRUE

	return retGo
}

// FilterKeypress is a wrapper around the C function gtk_im_context_filter_keypress.
func (recv *IMContext) FilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_im_context_filter_keypress((*C.GtkIMContext)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_im_context_focus_in : no return generator

// Unsupported : gtk_im_context_focus_out : no return generator

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// Unsupported : gtk_im_context_get_surrounding : unsupported parameter cursor_index : no type generator for gint, gint*

// Unsupported : gtk_im_context_reset : no return generator

// Unsupported : gtk_im_context_set_client_window : no return generator

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_im_context_set_surrounding : no return generator

// Unsupported : gtk_im_context_set_use_preedit : no return generator

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

	return g
}

func (recv *IMContextSimple) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContextSimpleNew is a wrapper around the C function gtk_im_context_simple_new.
func IMContextSimpleNew() *IMContext {
	retC := C.gtk_im_context_simple_new()
	retGo := IMContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_im_context_simple_add_compose_file : no return generator

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

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

	return g
}

func (recv *IMMulticontext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMMulticontextNew is a wrapper around the C function gtk_im_multicontext_new.
func IMMulticontextNew() *IMContext {
	retC := C.gtk_im_multicontext_new()
	retGo := IMContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_im_multicontext_append_menuitems : no return generator

// Unsupported : gtk_im_multicontext_set_context_id : no return generator

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

	return g
}

func (recv *IconFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconFactoryNew is a wrapper around the C function gtk_icon_factory_new.
func IconFactoryNew() *IconFactory {
	retC := C.gtk_icon_factory_new()
	retGo := IconFactoryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_factory_add : no return generator

// Unsupported : gtk_icon_factory_add_default : no return generator

// Lookup is a wrapper around the C function gtk_icon_factory_lookup.
func (recv *IconFactory) Lookup(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_icon_factory_lookup((*C.GtkIconFactory)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_factory_remove_default : no return generator

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

	return g
}

func (recv *IconInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_icon_info_free : no return generator

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : no param type

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_icon_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_for_context_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_set_raw_coordinates : no return generator

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

	return g
}

func (recv *IconTheme) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_icon_theme_add_resource_path : no return generator

// Unsupported : gtk_icon_theme_append_search_path : no return generator

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_theme_lookup_by_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_prepend_search_path : no return generator

// Unsupported : gtk_icon_theme_set_custom_theme : no return generator

// Unsupported : gtk_icon_theme_set_screen : no return generator

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path : no param type

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

	return g
}

func (recv *IconView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_icon_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_model : no return generator

// Unsupported : gtk_icon_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_icon_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_icon_view_item_activated : no return generator

// Unsupported : gtk_icon_view_scroll_to_path : no return generator

// Unsupported : gtk_icon_view_select_all : no return generator

// Unsupported : gtk_icon_view_select_path : no return generator

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc, GtkIconViewForeachFunc

// Unsupported : gtk_icon_view_set_activate_on_single_click : no return generator

// Unsupported : gtk_icon_view_set_column_spacing : no return generator

// Unsupported : gtk_icon_view_set_columns : no return generator

// Unsupported : gtk_icon_view_set_cursor : no return generator

// Unsupported : gtk_icon_view_set_drag_dest_item : no return generator

// Unsupported : gtk_icon_view_set_item_orientation : no return generator

// Unsupported : gtk_icon_view_set_item_padding : no return generator

// Unsupported : gtk_icon_view_set_item_width : no return generator

// Unsupported : gtk_icon_view_set_margin : no return generator

// Unsupported : gtk_icon_view_set_markup_column : no return generator

// Unsupported : gtk_icon_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_icon_view_set_pixbuf_column : no return generator

// Unsupported : gtk_icon_view_set_reorderable : no return generator

// Unsupported : gtk_icon_view_set_row_spacing : no return generator

// Unsupported : gtk_icon_view_set_selection_mode : no return generator

// Unsupported : gtk_icon_view_set_spacing : no return generator

// Unsupported : gtk_icon_view_set_text_column : no return generator

// Unsupported : gtk_icon_view_set_tooltip_cell : no return generator

// Unsupported : gtk_icon_view_set_tooltip_column : no return generator

// Unsupported : gtk_icon_view_set_tooltip_item : no return generator

// Unsupported : gtk_icon_view_unselect_all : no return generator

// Unsupported : gtk_icon_view_unselect_path : no return generator

// Unsupported : gtk_icon_view_unset_model_drag_dest : no return generator

// Unsupported : gtk_icon_view_unset_model_drag_source : no return generator

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

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageNew is a wrapper around the C function gtk_image_new.
func ImageNew() *Widget {
	retC := C.gtk_image_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromAnimation is a wrapper around the C function gtk_image_new_from_animation.
func ImageNewFromAnimation(animation *gdkpixbuf.PixbufAnimation) *Widget {
	c_animation := (*C.GdkPixbufAnimation)(animation.ToC())

	retC := C.gtk_image_new_from_animation(c_animation)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromFile is a wrapper around the C function gtk_image_new_from_file.
func ImageNewFromFile(filename string) *Widget {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_image_new_from_file(c_filename)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// ImageNewFromPixbuf is a wrapper around the C function gtk_image_new_from_pixbuf.
func ImageNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Widget {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	retC := C.gtk_image_new_from_pixbuf(c_pixbuf)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_clear : no return generator

// GetAnimation is a wrapper around the C function gtk_image_get_animation.
func (recv *Image) GetAnimation() *gdkpixbuf.PixbufAnimation {
	retC := C.gtk_image_get_animation((*C.GtkImage)(recv.native))
	retGo := gdkpixbuf.PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon**

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_get_icon_set : unsupported parameter icon_set : record with indirection level of 2

// GetPixbuf is a wrapper around the C function gtk_image_get_pixbuf.
func (recv *Image) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_image_get_pixbuf((*C.GtkImage)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_stock : unsupported parameter size : no type generator for gint, GtkIconSize*

// GetStorageType is a wrapper around the C function gtk_image_get_storage_type.
func (recv *Image) GetStorageType() ImageType {
	retC := C.gtk_image_get_storage_type((*C.GtkImage)(recv.native))
	retGo := (ImageType)(retC)

	return retGo
}

// Unsupported : gtk_image_set_from_animation : no return generator

// Unsupported : gtk_image_set_from_file : no return generator

// Unsupported : gtk_image_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_set_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_pixbuf : no return generator

// Unsupported : gtk_image_set_from_resource : no return generator

// Unsupported : gtk_image_set_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_surface : no return generator

// Unsupported : gtk_image_set_pixel_size : no return generator

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

	return g
}

func (recv *ImageMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageMenuItemNew is a wrapper around the C function gtk_image_menu_item_new.
func ImageMenuItemNew() *Widget {
	retC := C.gtk_image_menu_item_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewFromStock is a wrapper around the C function gtk_image_menu_item_new_from_stock.
func ImageMenuItemNewFromStock(stockId string, accelGroup *AccelGroup) *Widget {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	retC := C.gtk_image_menu_item_new_from_stock(c_stock_id, c_accel_group)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithLabel is a wrapper around the C function gtk_image_menu_item_new_with_label.
func ImageMenuItemNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithMnemonic is a wrapper around the C function gtk_image_menu_item_new_with_mnemonic.
func ImageMenuItemNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetImage is a wrapper around the C function gtk_image_menu_item_get_image.
func (recv *ImageMenuItem) GetImage() *Widget {
	retC := C.gtk_image_menu_item_get_image((*C.GtkImageMenuItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_menu_item_set_accel_group : no return generator

// Unsupported : gtk_image_menu_item_set_always_show_image : no return generator

// Unsupported : gtk_image_menu_item_set_image : no return generator

// Unsupported : gtk_image_menu_item_set_use_stock : no return generator

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

	return g
}

func (recv *InfoBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_info_bar_add_action_widget : no return generator

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_info_bar_response : no return generator

// Unsupported : gtk_info_bar_set_default_response : no return generator

// Unsupported : gtk_info_bar_set_message_type : no return generator

// Unsupported : gtk_info_bar_set_response_sensitive : no return generator

// Unsupported : gtk_info_bar_set_revealed : no return generator

// Unsupported : gtk_info_bar_set_show_close_button : no return generator

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

	return g
}

func (recv *Invisible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InvisibleNew is a wrapper around the C function gtk_invisible_new.
func InvisibleNew() *Widget {
	retC := C.gtk_invisible_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_invisible_set_screen : no return generator

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

	return g
}

func (recv *Label) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LabelNew is a wrapper around the C function gtk_label_new.
func LabelNew(str string) *Widget {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new(c_str)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LabelNewWithMnemonic is a wrapper around the C function gtk_label_new_with_mnemonic.
func LabelNewWithMnemonic(str string) *Widget {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new_with_mnemonic(c_str)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttributes is a wrapper around the C function gtk_label_get_attributes.
func (recv *Label) GetAttributes() *pango.AttrList {
	retC := C.gtk_label_get_attributes((*C.GtkLabel)(recv.native))
	retGo := pango.AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJustify is a wrapper around the C function gtk_label_get_justify.
func (recv *Label) GetJustify() Justification {
	retC := C.gtk_label_get_justify((*C.GtkLabel)(recv.native))
	retGo := (Justification)(retC)

	return retGo
}

// GetLabel is a wrapper around the C function gtk_label_get_label.
func (recv *Label) GetLabel() string {
	retC := C.gtk_label_get_label((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLayout is a wrapper around the C function gtk_label_get_layout.
func (recv *Label) GetLayout() *pango.Layout {
	retC := C.gtk_label_get_layout((*C.GtkLabel)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_label_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetLineWrap is a wrapper around the C function gtk_label_get_line_wrap.
func (recv *Label) GetLineWrap() bool {
	retC := C.gtk_label_get_line_wrap((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMnemonicKeyval is a wrapper around the C function gtk_label_get_mnemonic_keyval.
func (recv *Label) GetMnemonicKeyval() uint32 {
	retC := C.gtk_label_get_mnemonic_keyval((*C.GtkLabel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMnemonicWidget is a wrapper around the C function gtk_label_get_mnemonic_widget.
func (recv *Label) GetMnemonicWidget() *Widget {
	retC := C.gtk_label_get_mnemonic_widget((*C.GtkLabel)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectable is a wrapper around the C function gtk_label_get_selectable.
func (recv *Label) GetSelectable() bool {
	retC := C.gtk_label_get_selectable((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_label_get_selection_bounds : unsupported parameter start : no type generator for gint, gint*

// GetText is a wrapper around the C function gtk_label_get_text.
func (recv *Label) GetText() string {
	retC := C.gtk_label_get_text((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseMarkup is a wrapper around the C function gtk_label_get_use_markup.
func (recv *Label) GetUseMarkup() bool {
	retC := C.gtk_label_get_use_markup((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_label_get_use_underline.
func (recv *Label) GetUseUnderline() bool {
	retC := C.gtk_label_get_use_underline((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_label_select_region : no return generator

// Unsupported : gtk_label_set_angle : no return generator

// Unsupported : gtk_label_set_attributes : no return generator

// Unsupported : gtk_label_set_ellipsize : no return generator

// Unsupported : gtk_label_set_justify : no return generator

// Unsupported : gtk_label_set_label : no return generator

// Unsupported : gtk_label_set_line_wrap : no return generator

// Unsupported : gtk_label_set_line_wrap_mode : no return generator

// Unsupported : gtk_label_set_lines : no return generator

// Unsupported : gtk_label_set_markup : no return generator

// Unsupported : gtk_label_set_markup_with_mnemonic : no return generator

// Unsupported : gtk_label_set_max_width_chars : no return generator

// Unsupported : gtk_label_set_mnemonic_widget : no return generator

// Unsupported : gtk_label_set_pattern : no return generator

// Unsupported : gtk_label_set_selectable : no return generator

// Unsupported : gtk_label_set_single_line_mode : no return generator

// Unsupported : gtk_label_set_text : no return generator

// Unsupported : gtk_label_set_text_with_mnemonic : no return generator

// Unsupported : gtk_label_set_track_visited_links : no return generator

// Unsupported : gtk_label_set_use_markup : no return generator

// Unsupported : gtk_label_set_use_underline : no return generator

// Unsupported : gtk_label_set_width_chars : no return generator

// Unsupported : gtk_label_set_xalign : no return generator

// Unsupported : gtk_label_set_yalign : no return generator

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

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutNew is a wrapper around the C function gtk_layout_new.
func LayoutNew(hadjustment *Adjustment, vadjustment *Adjustment) *Widget {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_layout_new(c_hadjustment, c_vadjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_layout_get_hadjustment.
func (recv *Layout) GetHadjustment() *Adjustment {
	retC := C.gtk_layout_get_hadjustment((*C.GtkLayout)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_layout_get_size : unsupported parameter width : no type generator for guint, guint*

// GetVadjustment is a wrapper around the C function gtk_layout_get_vadjustment.
func (recv *Layout) GetVadjustment() *Adjustment {
	retC := C.gtk_layout_get_vadjustment((*C.GtkLayout)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_layout_move : no return generator

// Unsupported : gtk_layout_put : no return generator

// Unsupported : gtk_layout_set_hadjustment : no return generator

// Unsupported : gtk_layout_set_size : no return generator

// Unsupported : gtk_layout_set_vadjustment : no return generator

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

	return g
}

func (recv *LevelBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_level_bar_add_offset_value : no return generator

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

// Unsupported : gtk_level_bar_remove_offset_value : no return generator

// Unsupported : gtk_level_bar_set_inverted : no return generator

// Unsupported : gtk_level_bar_set_max_value : no return generator

// Unsupported : gtk_level_bar_set_min_value : no return generator

// Unsupported : gtk_level_bar_set_mode : no return generator

// Unsupported : gtk_level_bar_set_value : no return generator

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

	return g
}

func (recv *LinkButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_link_button_set_uri : no return generator

// Unsupported : gtk_link_button_set_visited : no return generator

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

	return g
}

func (recv *ListBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_list_box_drag_highlight_row : no return generator

// Unsupported : gtk_list_box_drag_unhighlight_row : no return generator

// Unsupported : gtk_list_box_insert : no return generator

// Unsupported : gtk_list_box_invalidate_filter : no return generator

// Unsupported : gtk_list_box_invalidate_headers : no return generator

// Unsupported : gtk_list_box_invalidate_sort : no return generator

// Unsupported : gtk_list_box_prepend : no return generator

// Unsupported : gtk_list_box_select_all : no return generator

// Unsupported : gtk_list_box_select_row : no return generator

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc, GtkListBoxForeachFunc

// Unsupported : gtk_list_box_set_activate_on_single_click : no return generator

// Unsupported : gtk_list_box_set_adjustment : no return generator

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc, GtkListBoxFilterFunc

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc, GtkListBoxUpdateHeaderFunc

// Unsupported : gtk_list_box_set_placeholder : no return generator

// Unsupported : gtk_list_box_set_selection_mode : no return generator

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc, GtkListBoxSortFunc

// Unsupported : gtk_list_box_unselect_all : no return generator

// Unsupported : gtk_list_box_unselect_row : no return generator

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

	return g
}

func (recv *ListBoxRow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_list_box_row_changed : no return generator

// Unsupported : gtk_list_box_row_set_activatable : no return generator

// Unsupported : gtk_list_box_row_set_header : no return generator

// Unsupported : gtk_list_box_row_set_selectable : no return generator

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

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_list_store_append : no return generator

// Unsupported : gtk_list_store_clear : no return generator

// Unsupported : gtk_list_store_insert : no return generator

// Unsupported : gtk_list_store_insert_after : no return generator

// Unsupported : gtk_list_store_insert_before : no return generator

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_list_store_move_after : no return generator

// Unsupported : gtk_list_store_move_before : no return generator

// Unsupported : gtk_list_store_prepend : no return generator

// Remove is a wrapper around the C function gtk_list_store_remove.
func (recv *ListStore) Remove(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_list_store_remove((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_list_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_list_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_list_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_list_store_set_value : no return generator

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_list_store_swap : no return generator

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

	return g
}

func (recv *LockButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_lock_button_set_permission : no return generator

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

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuNew is a wrapper around the C function gtk_menu_new.
func MenuNew() *Widget {
	retC := C.gtk_menu_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_attach : no return generator

// Unsupported : gtk_menu_attach_to_widget : unsupported parameter detacher : no type generator for MenuDetachFunc, GtkMenuDetachFunc

// Unsupported : gtk_menu_detach : no return generator

// GetAccelGroup is a wrapper around the C function gtk_menu_get_accel_group.
func (recv *Menu) GetAccelGroup() *AccelGroup {
	retC := C.gtk_menu_get_accel_group((*C.GtkMenu)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_menu_get_active.
func (recv *Menu) GetActive() *Widget {
	retC := C.gtk_menu_get_active((*C.GtkMenu)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttachWidget is a wrapper around the C function gtk_menu_get_attach_widget.
func (recv *Menu) GetAttachWidget() *Widget {
	retC := C.gtk_menu_get_attach_widget((*C.GtkMenu)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTearoffState is a wrapper around the C function gtk_menu_get_tearoff_state.
func (recv *Menu) GetTearoffState() bool {
	retC := C.gtk_menu_get_tearoff_state((*C.GtkMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_menu_get_title.
func (recv *Menu) GetTitle() string {
	retC := C.gtk_menu_get_title((*C.GtkMenu)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_menu_place_on_monitor : no return generator

// Unsupported : gtk_menu_popdown : no return generator

// Unsupported : gtk_menu_popup : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_popup_at_pointer : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_at_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_menu_popup_at_widget : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_reorder_child : no return generator

// Unsupported : gtk_menu_reposition : no return generator

// Unsupported : gtk_menu_set_accel_group : no return generator

// Unsupported : gtk_menu_set_accel_path : no return generator

// Unsupported : gtk_menu_set_active : no return generator

// Unsupported : gtk_menu_set_monitor : no return generator

// Unsupported : gtk_menu_set_reserve_toggle_size : no return generator

// Unsupported : gtk_menu_set_screen : no return generator

// Unsupported : gtk_menu_set_tearoff_state : no return generator

// Unsupported : gtk_menu_set_title : no return generator

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

	return g
}

func (recv *MenuBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuBarNew is a wrapper around the C function gtk_menu_bar_new.
func MenuBarNew() *Widget {
	retC := C.gtk_menu_bar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_bar_set_child_pack_direction : no return generator

// Unsupported : gtk_menu_bar_set_pack_direction : no return generator

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

	return g
}

func (recv *MenuButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_menu_button_set_align_widget : no return generator

// Unsupported : gtk_menu_button_set_direction : no return generator

// Unsupported : gtk_menu_button_set_menu_model : no return generator

// Unsupported : gtk_menu_button_set_popover : no return generator

// Unsupported : gtk_menu_button_set_popup : no return generator

// Unsupported : gtk_menu_button_set_use_popover : no return generator

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

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemNew is a wrapper around the C function gtk_menu_item_new.
func MenuItemNew() *Widget {
	retC := C.gtk_menu_item_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithLabel is a wrapper around the C function gtk_menu_item_new_with_label.
func MenuItemNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithMnemonic is a wrapper around the C function gtk_menu_item_new_with_mnemonic.
func MenuItemNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_item_activate : no return generator

// Unsupported : gtk_menu_item_deselect : no return generator

// GetRightJustified is a wrapper around the C function gtk_menu_item_get_right_justified.
func (recv *MenuItem) GetRightJustified() bool {
	retC := C.gtk_menu_item_get_right_justified((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSubmenu is a wrapper around the C function gtk_menu_item_get_submenu.
func (recv *MenuItem) GetSubmenu() *Widget {
	retC := C.gtk_menu_item_get_submenu((*C.GtkMenuItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_item_select : no return generator

// Unsupported : gtk_menu_item_set_accel_path : no return generator

// Unsupported : gtk_menu_item_set_label : no return generator

// Unsupported : gtk_menu_item_set_reserve_indicator : no return generator

// Unsupported : gtk_menu_item_set_right_justified : no return generator

// Unsupported : gtk_menu_item_set_submenu : no return generator

// Unsupported : gtk_menu_item_set_use_underline : no return generator

// Unsupported : gtk_menu_item_toggle_size_allocate : no return generator

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

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

	return g
}

func (recv *MenuShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_menu_shell_activate_item : no return generator

// Unsupported : gtk_menu_shell_append : no return generator

// Unsupported : gtk_menu_shell_bind_model : no return generator

// Unsupported : gtk_menu_shell_cancel : no return generator

// Unsupported : gtk_menu_shell_deactivate : no return generator

// Unsupported : gtk_menu_shell_deselect : no return generator

// Unsupported : gtk_menu_shell_insert : no return generator

// Unsupported : gtk_menu_shell_prepend : no return generator

// Unsupported : gtk_menu_shell_select_first : no return generator

// Unsupported : gtk_menu_shell_select_item : no return generator

// Unsupported : gtk_menu_shell_set_take_focus : no return generator

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

	return g
}

func (recv *MenuToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_menu_tool_button_set_arrow_tooltip_markup : no return generator

// Unsupported : gtk_menu_tool_button_set_arrow_tooltip_text : no return generator

// Unsupported : gtk_menu_tool_button_set_menu : no return generator

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

	return g
}

func (recv *MessageDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_set_image : no return generator

// Unsupported : gtk_message_dialog_set_markup : no return generator

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

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_misc_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_misc_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_misc_set_alignment : no return generator

// Unsupported : gtk_misc_set_padding : no return generator

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

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_mount_operation_set_parent : no return generator

// Unsupported : gtk_mount_operation_set_screen : no return generator

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

	return g
}

func (recv *Notebook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookNew is a wrapper around the C function gtk_notebook_new.
func NotebookNew() *Widget {
	retC := C.gtk_notebook_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendPage is a wrapper around the C function gtk_notebook_append_page.
func (recv *Notebook) AppendPage(child *Widget, tabLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	retC := C.gtk_notebook_append_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label)
	retGo := (int32)(retC)

	return retGo
}

// AppendPageMenu is a wrapper around the C function gtk_notebook_append_page_menu.
func (recv *Notebook) AppendPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	retC := C.gtk_notebook_append_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_notebook_detach_tab : no return generator

// GetCurrentPage is a wrapper around the C function gtk_notebook_get_current_page.
func (recv *Notebook) GetCurrentPage() int32 {
	retC := C.gtk_notebook_get_current_page((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMenuLabel is a wrapper around the C function gtk_notebook_get_menu_label.
func (recv *Notebook) GetMenuLabel(child *Widget) *Widget {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_menu_label((*C.GtkNotebook)(recv.native), c_child)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMenuLabelText is a wrapper around the C function gtk_notebook_get_menu_label_text.
func (recv *Notebook) GetMenuLabelText(child *Widget) string {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_menu_label_text((*C.GtkNotebook)(recv.native), c_child)
	retGo := C.GoString(retC)

	return retGo
}

// GetNthPage is a wrapper around the C function gtk_notebook_get_nth_page.
func (recv *Notebook) GetNthPage(pageNum int32) *Widget {
	c_page_num := (C.gint)(pageNum)

	retC := C.gtk_notebook_get_nth_page((*C.GtkNotebook)(recv.native), c_page_num)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScrollable is a wrapper around the C function gtk_notebook_get_scrollable.
func (recv *Notebook) GetScrollable() bool {
	retC := C.gtk_notebook_get_scrollable((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowBorder is a wrapper around the C function gtk_notebook_get_show_border.
func (recv *Notebook) GetShowBorder() bool {
	retC := C.gtk_notebook_get_show_border((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowTabs is a wrapper around the C function gtk_notebook_get_show_tabs.
func (recv *Notebook) GetShowTabs() bool {
	retC := C.gtk_notebook_get_show_tabs((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTabLabel is a wrapper around the C function gtk_notebook_get_tab_label.
func (recv *Notebook) GetTabLabel(child *Widget) *Widget {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_label((*C.GtkNotebook)(recv.native), c_child)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTabLabelText is a wrapper around the C function gtk_notebook_get_tab_label_text.
func (recv *Notebook) GetTabLabelText(child *Widget) string {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_label_text((*C.GtkNotebook)(recv.native), c_child)
	retGo := C.GoString(retC)

	return retGo
}

// GetTabPos is a wrapper around the C function gtk_notebook_get_tab_pos.
func (recv *Notebook) GetTabPos() PositionType {
	retC := C.gtk_notebook_get_tab_pos((*C.GtkNotebook)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// InsertPage is a wrapper around the C function gtk_notebook_insert_page.
func (recv *Notebook) InsertPage(child *Widget, tabLabel *Widget, position int32) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_notebook_insert_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_position)
	retGo := (int32)(retC)

	return retGo
}

// InsertPageMenu is a wrapper around the C function gtk_notebook_insert_page_menu.
func (recv *Notebook) InsertPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget, position int32) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_notebook_insert_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label, c_position)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_notebook_next_page : no return generator

// PageNum is a wrapper around the C function gtk_notebook_page_num.
func (recv *Notebook) PageNum(child *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_page_num((*C.GtkNotebook)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_notebook_popup_disable : no return generator

// Unsupported : gtk_notebook_popup_enable : no return generator

// PrependPage is a wrapper around the C function gtk_notebook_prepend_page.
func (recv *Notebook) PrependPage(child *Widget, tabLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	retC := C.gtk_notebook_prepend_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label)
	retGo := (int32)(retC)

	return retGo
}

// PrependPageMenu is a wrapper around the C function gtk_notebook_prepend_page_menu.
func (recv *Notebook) PrependPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	retC := C.gtk_notebook_prepend_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_notebook_prev_page : no return generator

// Unsupported : gtk_notebook_remove_page : no return generator

// Unsupported : gtk_notebook_reorder_child : no return generator

// Unsupported : gtk_notebook_set_action_widget : no return generator

// Unsupported : gtk_notebook_set_current_page : no return generator

// Unsupported : gtk_notebook_set_group_name : no return generator

// Unsupported : gtk_notebook_set_menu_label : no return generator

// Unsupported : gtk_notebook_set_menu_label_text : no return generator

// Unsupported : gtk_notebook_set_scrollable : no return generator

// Unsupported : gtk_notebook_set_show_border : no return generator

// Unsupported : gtk_notebook_set_show_tabs : no return generator

// Unsupported : gtk_notebook_set_tab_detachable : no return generator

// Unsupported : gtk_notebook_set_tab_label : no return generator

// Unsupported : gtk_notebook_set_tab_label_text : no return generator

// Unsupported : gtk_notebook_set_tab_pos : no return generator

// Unsupported : gtk_notebook_set_tab_reorderable : no return generator

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

	return g
}

func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotebookPageAccessibleNew is a wrapper around the C function gtk_notebook_page_accessible_new.
func NotebookPageAccessibleNew(notebook *NotebookAccessible, child *Widget) *atk.Object {
	c_notebook := (*C.GtkNotebookAccessible)(notebook.ToC())

	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_page_accessible_new(c_notebook, c_child)
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_notebook_page_accessible_invalidate : no return generator

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

	return g
}

func (recv *NumerableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_numerable_icon_get_background_gicon : no return generator

// Unsupported : gtk_numerable_icon_set_background_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_numerable_icon_set_background_icon_name : no return generator

// Unsupported : gtk_numerable_icon_set_count : no return generator

// Unsupported : gtk_numerable_icon_set_label : no return generator

// Unsupported : gtk_numerable_icon_set_style_context : no return generator

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

	return g
}

func (recv *Overlay) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_overlay_add_overlay : no return generator

// Unsupported : gtk_overlay_reorder_overlay : no return generator

// Unsupported : gtk_overlay_set_overlay_pass_through : no return generator

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

	return g
}

func (recv *PageSetup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_set_bottom_margin : no return generator

// Unsupported : gtk_page_setup_set_left_margin : no return generator

// Unsupported : gtk_page_setup_set_orientation : no return generator

// Unsupported : gtk_page_setup_set_paper_size : no return generator

// Unsupported : gtk_page_setup_set_paper_size_and_default_margins : no return generator

// Unsupported : gtk_page_setup_set_right_margin : no return generator

// Unsupported : gtk_page_setup_set_top_margin : no return generator

// Unsupported : gtk_page_setup_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_to_key_file : no return generator

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

	return g
}

func (recv *Paned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_paned_add1 : no return generator

// Unsupported : gtk_paned_add2 : no return generator

// GetPosition is a wrapper around the C function gtk_paned_get_position.
func (recv *Paned) GetPosition() int32 {
	retC := C.gtk_paned_get_position((*C.GtkPaned)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_paned_pack1 : no return generator

// Unsupported : gtk_paned_pack2 : no return generator

// Unsupported : gtk_paned_set_position : no return generator

// Unsupported : gtk_paned_set_wide_handle : no return generator

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

	return g
}

func (recv *PlacesSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// GetShowConnectToServer is a wrapper around the C function gtk_places_sidebar_get_show_connect_to_server.
func (recv *PlacesSidebar) GetShowConnectToServer() bool {
	retC := C.gtk_places_sidebar_get_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_drop_targets_visible : no return generator

// Unsupported : gtk_places_sidebar_set_local_only : no return generator

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_open_flags : no return generator

// Unsupported : gtk_places_sidebar_set_show_connect_to_server : no return generator

// Unsupported : gtk_places_sidebar_set_show_desktop : no return generator

// Unsupported : gtk_places_sidebar_set_show_enter_location : no return generator

// Unsupported : gtk_places_sidebar_set_show_other_locations : no return generator

// Unsupported : gtk_places_sidebar_set_show_recent : no return generator

// Unsupported : gtk_places_sidebar_set_show_starred_location : no return generator

// Unsupported : gtk_places_sidebar_set_show_trash : no return generator

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

	return g
}

func (recv *Popover) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_popover_bind_model : no return generator

// Unsupported : gtk_popover_get_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetPosition is a wrapper around the C function gtk_popover_get_position.
func (recv *Popover) GetPosition() PositionType {
	retC := C.gtk_popover_get_position((*C.GtkPopover)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// Unsupported : gtk_popover_popdown : no return generator

// Unsupported : gtk_popover_popup : no return generator

// Unsupported : gtk_popover_set_constrain_to : no return generator

// Unsupported : gtk_popover_set_default_widget : no return generator

// Unsupported : gtk_popover_set_modal : no return generator

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_popover_set_position : no return generator

// Unsupported : gtk_popover_set_relative_to : no return generator

// Unsupported : gtk_popover_set_transitions_enabled : no return generator

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

	return g
}

func (recv *PopoverMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_popover_menu_open_submenu : no return generator

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

	return g
}

func (recv *PrintContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_print_context_get_hard_margins : unsupported parameter top : no type generator for gdouble, gdouble*

// Unsupported : gtk_print_context_set_cairo_context : no return generator

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

	return g
}

func (recv *PrintOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_print_operation_cancel : no return generator

// Unsupported : gtk_print_operation_draw_page_finish : no return generator

// Unsupported : gtk_print_operation_get_error : no return generator

// Unsupported : gtk_print_operation_set_allow_async : no return generator

// Unsupported : gtk_print_operation_set_current_page : no return generator

// Unsupported : gtk_print_operation_set_custom_tab_label : no return generator

// Unsupported : gtk_print_operation_set_default_page_setup : no return generator

// Unsupported : gtk_print_operation_set_defer_drawing : no return generator

// Unsupported : gtk_print_operation_set_embed_page_setup : no return generator

// Unsupported : gtk_print_operation_set_export_filename : no return generator

// Unsupported : gtk_print_operation_set_has_selection : no return generator

// Unsupported : gtk_print_operation_set_job_name : no return generator

// Unsupported : gtk_print_operation_set_n_pages : no return generator

// Unsupported : gtk_print_operation_set_print_settings : no return generator

// Unsupported : gtk_print_operation_set_show_progress : no return generator

// Unsupported : gtk_print_operation_set_support_selection : no return generator

// Unsupported : gtk_print_operation_set_track_print_status : no return generator

// Unsupported : gtk_print_operation_set_unit : no return generator

// Unsupported : gtk_print_operation_set_use_full_page : no return generator

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

	return g
}

func (recv *PrintSettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc, GtkPrintSettingsFunc

// Unsupported : gtk_print_settings_get_page_ranges : unsupported parameter num_ranges : no type generator for gint, gint*

// Unsupported : gtk_print_settings_set : no return generator

// Unsupported : gtk_print_settings_set_bool : no return generator

// Unsupported : gtk_print_settings_set_collate : no return generator

// Unsupported : gtk_print_settings_set_default_source : no return generator

// Unsupported : gtk_print_settings_set_dither : no return generator

// Unsupported : gtk_print_settings_set_double : no return generator

// Unsupported : gtk_print_settings_set_duplex : no return generator

// Unsupported : gtk_print_settings_set_finishings : no return generator

// Unsupported : gtk_print_settings_set_int : no return generator

// Unsupported : gtk_print_settings_set_length : no return generator

// Unsupported : gtk_print_settings_set_media_type : no return generator

// Unsupported : gtk_print_settings_set_n_copies : no return generator

// Unsupported : gtk_print_settings_set_number_up : no return generator

// Unsupported : gtk_print_settings_set_number_up_layout : no return generator

// Unsupported : gtk_print_settings_set_orientation : no return generator

// Unsupported : gtk_print_settings_set_output_bin : no return generator

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// Unsupported : gtk_print_settings_set_page_set : no return generator

// Unsupported : gtk_print_settings_set_paper_height : no return generator

// Unsupported : gtk_print_settings_set_paper_size : no return generator

// Unsupported : gtk_print_settings_set_paper_width : no return generator

// Unsupported : gtk_print_settings_set_print_pages : no return generator

// Unsupported : gtk_print_settings_set_printer : no return generator

// Unsupported : gtk_print_settings_set_printer_lpi : no return generator

// Unsupported : gtk_print_settings_set_quality : no return generator

// Unsupported : gtk_print_settings_set_resolution : no return generator

// Unsupported : gtk_print_settings_set_resolution_xy : no return generator

// Unsupported : gtk_print_settings_set_reverse : no return generator

// Unsupported : gtk_print_settings_set_scale : no return generator

// Unsupported : gtk_print_settings_set_use_color : no return generator

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_to_key_file : no return generator

// Unsupported : gtk_print_settings_unset : no return generator

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

	return g
}

func (recv *ProgressBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProgressBarNew is a wrapper around the C function gtk_progress_bar_new.
func ProgressBarNew() *Widget {
	retC := C.gtk_progress_bar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFraction is a wrapper around the C function gtk_progress_bar_get_fraction.
func (recv *ProgressBar) GetFraction() float64 {
	retC := C.gtk_progress_bar_get_fraction((*C.GtkProgressBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetInverted is a wrapper around the C function gtk_progress_bar_get_inverted.
func (recv *ProgressBar) GetInverted() bool {
	retC := C.gtk_progress_bar_get_inverted((*C.GtkProgressBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPulseStep is a wrapper around the C function gtk_progress_bar_get_pulse_step.
func (recv *ProgressBar) GetPulseStep() float64 {
	retC := C.gtk_progress_bar_get_pulse_step((*C.GtkProgressBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_progress_bar_get_text.
func (recv *ProgressBar) GetText() string {
	retC := C.gtk_progress_bar_get_text((*C.GtkProgressBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_progress_bar_pulse : no return generator

// Unsupported : gtk_progress_bar_set_ellipsize : no return generator

// Unsupported : gtk_progress_bar_set_fraction : no return generator

// Unsupported : gtk_progress_bar_set_inverted : no return generator

// Unsupported : gtk_progress_bar_set_pulse_step : no return generator

// Unsupported : gtk_progress_bar_set_show_text : no return generator

// Unsupported : gtk_progress_bar_set_text : no return generator

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

	return g
}

func (recv *RadioAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_radio_action_join_group : no return generator

// Unsupported : gtk_radio_action_set_current_value : no return generator

// Unsupported : gtk_radio_action_set_group : no return generator

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

	return g
}

func (recv *RadioButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioButtonNew is a wrapper around the C function gtk_radio_button_new.
func RadioButtonNew(group *glib.SList) *Widget {
	c_group := (*C.GSList)(group.ToC())

	retC := C.gtk_radio_button_new(c_group)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewFromWidget is a wrapper around the C function gtk_radio_button_new_from_widget.
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) *Widget {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	retC := C.gtk_radio_button_new_from_widget(c_radio_group_member)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabel is a wrapper around the C function gtk_radio_button_new_with_label.
func RadioButtonNewWithLabel(group *glib.SList, label string) *Widget {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label(c_group, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabelFromWidget is a wrapper around the C function gtk_radio_button_new_with_label_from_widget.
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) *Widget {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label_from_widget(c_radio_group_member, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonic is a wrapper around the C function gtk_radio_button_new_with_mnemonic.
func RadioButtonNewWithMnemonic(group *glib.SList, label string) *Widget {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic(c_group, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_button_new_with_mnemonic_from_widget.
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) *Widget {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic_from_widget(c_radio_group_member, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_button_get_group.
func (recv *RadioButton) GetGroup() *glib.SList {
	retC := C.gtk_radio_button_get_group((*C.GtkRadioButton)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_radio_button_join_group : no return generator

// Unsupported : gtk_radio_button_set_group : no return generator

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

	return g
}

func (recv *RadioMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RadioMenuItemNew is a wrapper around the C function gtk_radio_menu_item_new.
func RadioMenuItemNew(group *glib.SList) *Widget {
	c_group := (*C.GSList)(group.ToC())

	retC := C.gtk_radio_menu_item_new(c_group)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabel is a wrapper around the C function gtk_radio_menu_item_new_with_label.
func RadioMenuItemNewWithLabel(group *glib.SList, label string) *Widget {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label(c_group, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonic is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic.
func RadioMenuItemNewWithMnemonic(group *glib.SList, label string) *Widget {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic(c_group, c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_menu_item_get_group.
func (recv *RadioMenuItem) GetGroup() *glib.SList {
	retC := C.gtk_radio_menu_item_get_group((*C.GtkRadioMenuItem)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_radio_menu_item_join_group : no return generator

// Unsupported : gtk_radio_menu_item_set_group : no return generator

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

	return g
}

func (recv *RadioToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_radio_tool_button_set_group : no return generator

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

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetAdjustment is a wrapper around the C function gtk_range_get_adjustment.
func (recv *Range) GetAdjustment() *Adjustment {
	retC := C.gtk_range_get_adjustment((*C.GtkRange)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInverted is a wrapper around the C function gtk_range_get_inverted.
func (recv *Range) GetInverted() bool {
	retC := C.gtk_range_get_inverted((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_range_get_slider_range : unsupported parameter slider_start : no type generator for gint, gint*

// GetValue is a wrapper around the C function gtk_range_get_value.
func (recv *Range) GetValue() float64 {
	retC := C.gtk_range_get_value((*C.GtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Unsupported : gtk_range_set_adjustment : no return generator

// Unsupported : gtk_range_set_fill_level : no return generator

// Unsupported : gtk_range_set_flippable : no return generator

// Unsupported : gtk_range_set_increments : no return generator

// Unsupported : gtk_range_set_inverted : no return generator

// Unsupported : gtk_range_set_lower_stepper_sensitivity : no return generator

// Unsupported : gtk_range_set_min_slider_size : no return generator

// Unsupported : gtk_range_set_range : no return generator

// Unsupported : gtk_range_set_restrict_to_fill_level : no return generator

// Unsupported : gtk_range_set_round_digits : no return generator

// Unsupported : gtk_range_set_show_fill_level : no return generator

// Unsupported : gtk_range_set_slider_size_fixed : no return generator

// Unsupported : gtk_range_set_upper_stepper_sensitivity : no return generator

// Unsupported : gtk_range_set_value : no return generator

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

	return retGo
}

// Copy is a wrapper around the C function gtk_rc_style_copy.
func (recv *RcStyle) Copy() *RcStyle {
	retC := C.gtk_rc_style_copy((*C.GtkRcStyle)(recv.native))
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *RecentAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_action_set_show_numbers : no return generator

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

	return g
}

func (recv *RecentChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

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

	return g
}

func (recv *RecentChooserMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_chooser_menu_set_show_numbers : no return generator

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

	return g
}

func (recv *RecentFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_filter_add_age : no return generator

// Unsupported : gtk_recent_filter_add_application : no return generator

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// Unsupported : gtk_recent_filter_add_group : no return generator

// Unsupported : gtk_recent_filter_add_mime_type : no return generator

// Unsupported : gtk_recent_filter_add_pattern : no return generator

// Unsupported : gtk_recent_filter_add_pixbuf_formats : no return generator

// Unsupported : gtk_recent_filter_set_name : no return generator

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

	return g
}

func (recv *RendererCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessibleNew is a wrapper around the C function gtk_renderer_cell_accessible_new.
func RendererCellAccessibleNew(renderer *CellRenderer) *atk.Object {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	retC := C.gtk_renderer_cell_accessible_new(c_renderer)
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *Revealer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_revealer_set_reveal_child : no return generator

// Unsupported : gtk_revealer_set_transition_duration : no return generator

// Unsupported : gtk_revealer_set_transition_type : no return generator

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

	return g
}

func (recv *Scale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_scale_add_mark : no return generator

// Unsupported : gtk_scale_clear_marks : no return generator

// GetDigits is a wrapper around the C function gtk_scale_get_digits.
func (recv *Scale) GetDigits() int32 {
	retC := C.gtk_scale_get_digits((*C.GtkScale)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDrawValue is a wrapper around the C function gtk_scale_get_draw_value.
func (recv *Scale) GetDrawValue() bool {
	retC := C.gtk_scale_get_draw_value((*C.GtkScale)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetValuePos is a wrapper around the C function gtk_scale_get_value_pos.
func (recv *Scale) GetValuePos() PositionType {
	retC := C.gtk_scale_get_value_pos((*C.GtkScale)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// Unsupported : gtk_scale_set_digits : no return generator

// Unsupported : gtk_scale_set_draw_value : no return generator

// Unsupported : gtk_scale_set_has_origin : no return generator

// Unsupported : gtk_scale_set_value_pos : no return generator

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

	return g
}

func (recv *ScaleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_scale_button_set_adjustment : no return generator

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// Unsupported : gtk_scale_button_set_value : no return generator

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

	return g
}

func (recv *ScrolledWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindowNew is a wrapper around the C function gtk_scrolled_window_new.
func ScrolledWindowNew(hadjustment *Adjustment, vadjustment *Adjustment) *Widget {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_scrolled_window_new(c_hadjustment, c_vadjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_scrolled_window_add_with_viewport : no return generator

// GetHadjustment is a wrapper around the C function gtk_scrolled_window_get_hadjustment.
func (recv *ScrolledWindow) GetHadjustment() *Adjustment {
	retC := C.gtk_scrolled_window_get_hadjustment((*C.GtkScrolledWindow)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPlacement is a wrapper around the C function gtk_scrolled_window_get_placement.
func (recv *ScrolledWindow) GetPlacement() CornerType {
	retC := C.gtk_scrolled_window_get_placement((*C.GtkScrolledWindow)(recv.native))
	retGo := (CornerType)(retC)

	return retGo
}

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// GetShadowType is a wrapper around the C function gtk_scrolled_window_get_shadow_type.
func (recv *ScrolledWindow) GetShadowType() ShadowType {
	retC := C.gtk_scrolled_window_get_shadow_type((*C.GtkScrolledWindow)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_scrolled_window_get_vadjustment.
func (recv *ScrolledWindow) GetVadjustment() *Adjustment {
	retC := C.gtk_scrolled_window_get_vadjustment((*C.GtkScrolledWindow)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_scrolled_window_set_capture_button_press : no return generator

// Unsupported : gtk_scrolled_window_set_hadjustment : no return generator

// Unsupported : gtk_scrolled_window_set_kinetic_scrolling : no return generator

// Unsupported : gtk_scrolled_window_set_max_content_height : no return generator

// Unsupported : gtk_scrolled_window_set_max_content_width : no return generator

// Unsupported : gtk_scrolled_window_set_min_content_height : no return generator

// Unsupported : gtk_scrolled_window_set_min_content_width : no return generator

// Unsupported : gtk_scrolled_window_set_overlay_scrolling : no return generator

// Unsupported : gtk_scrolled_window_set_placement : no return generator

// Unsupported : gtk_scrolled_window_set_policy : no return generator

// Unsupported : gtk_scrolled_window_set_propagate_natural_height : no return generator

// Unsupported : gtk_scrolled_window_set_propagate_natural_width : no return generator

// Unsupported : gtk_scrolled_window_set_shadow_type : no return generator

// Unsupported : gtk_scrolled_window_set_vadjustment : no return generator

// Unsupported : gtk_scrolled_window_unset_placement : no return generator

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

	return g
}

func (recv *SearchBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_search_bar_connect_entry : no return generator

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_search_bar_set_search_mode : no return generator

// Unsupported : gtk_search_bar_set_show_close_button : no return generator

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

	return g
}

func (recv *SearchEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

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

	return g
}

func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeparatorMenuItemNew is a wrapper around the C function gtk_separator_menu_item_new.
func SeparatorMenuItemNew() *Widget {
	retC := C.gtk_separator_menu_item_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *SeparatorToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_separator_tool_item_set_draw : no return generator

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

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_settings_reset_property : no return generator

// Unsupported : gtk_settings_set_double_property : no return generator

// Unsupported : gtk_settings_set_long_property : no return generator

// Unsupported : gtk_settings_set_property_value : no return generator

// Unsupported : gtk_settings_set_string_property : no return generator

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

	return retGo
}

// Unsupported : gtk_size_group_add_widget : no return generator

// GetMode is a wrapper around the C function gtk_size_group_get_mode.
func (recv *SizeGroup) GetMode() SizeGroupMode {
	retC := C.gtk_size_group_get_mode((*C.GtkSizeGroup)(recv.native))
	retGo := (SizeGroupMode)(retC)

	return retGo
}

// Unsupported : gtk_size_group_remove_widget : no return generator

// Unsupported : gtk_size_group_set_ignore_hidden : no return generator

// Unsupported : gtk_size_group_set_mode : no return generator

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

	return g
}

func (recv *SpinButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SpinButtonNew is a wrapper around the C function gtk_spin_button_new.
func SpinButtonNew(adjustment *Adjustment, climbRate float64, digits uint32) *Widget {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	c_climb_rate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	retC := C.gtk_spin_button_new(c_adjustment, c_climb_rate, c_digits)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SpinButtonNewWithRange is a wrapper around the C function gtk_spin_button_new_with_range.
func SpinButtonNewWithRange(min float64, max float64, step float64) *Widget {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_spin_button_new_with_range(c_min, c_max, c_step)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_spin_button_configure : no return generator

// GetAdjustment is a wrapper around the C function gtk_spin_button_get_adjustment.
func (recv *SpinButton) GetAdjustment() *Adjustment {
	retC := C.gtk_spin_button_get_adjustment((*C.GtkSpinButton)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDigits is a wrapper around the C function gtk_spin_button_get_digits.
func (recv *SpinButton) GetDigits() uint32 {
	retC := C.gtk_spin_button_get_digits((*C.GtkSpinButton)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// GetNumeric is a wrapper around the C function gtk_spin_button_get_numeric.
func (recv *SpinButton) GetNumeric() bool {
	retC := C.gtk_spin_button_get_numeric((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// GetSnapToTicks is a wrapper around the C function gtk_spin_button_get_snap_to_ticks.
func (recv *SpinButton) GetSnapToTicks() bool {
	retC := C.gtk_spin_button_get_snap_to_ticks((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUpdatePolicy is a wrapper around the C function gtk_spin_button_get_update_policy.
func (recv *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	retC := C.gtk_spin_button_get_update_policy((*C.GtkSpinButton)(recv.native))
	retGo := (SpinButtonUpdatePolicy)(retC)

	return retGo
}

// GetValue is a wrapper around the C function gtk_spin_button_get_value.
func (recv *SpinButton) GetValue() float64 {
	retC := C.gtk_spin_button_get_value((*C.GtkSpinButton)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetValueAsInt is a wrapper around the C function gtk_spin_button_get_value_as_int.
func (recv *SpinButton) GetValueAsInt() int32 {
	retC := C.gtk_spin_button_get_value_as_int((*C.GtkSpinButton)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWrap is a wrapper around the C function gtk_spin_button_get_wrap.
func (recv *SpinButton) GetWrap() bool {
	retC := C.gtk_spin_button_get_wrap((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_spin_button_set_adjustment : no return generator

// Unsupported : gtk_spin_button_set_digits : no return generator

// Unsupported : gtk_spin_button_set_increments : no return generator

// Unsupported : gtk_spin_button_set_numeric : no return generator

// Unsupported : gtk_spin_button_set_range : no return generator

// Unsupported : gtk_spin_button_set_snap_to_ticks : no return generator

// Unsupported : gtk_spin_button_set_update_policy : no return generator

// Unsupported : gtk_spin_button_set_value : no return generator

// Unsupported : gtk_spin_button_set_wrap : no return generator

// Unsupported : gtk_spin_button_spin : no return generator

// Unsupported : gtk_spin_button_update : no return generator

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

	return g
}

func (recv *Spinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_spinner_start : no return generator

// Unsupported : gtk_spinner_stop : no return generator

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

	return g
}

func (recv *Stack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_stack_add_named : no return generator

// Unsupported : gtk_stack_add_titled : no return generator

// Unsupported : gtk_stack_set_hhomogeneous : no return generator

// Unsupported : gtk_stack_set_homogeneous : no return generator

// Unsupported : gtk_stack_set_interpolate_size : no return generator

// Unsupported : gtk_stack_set_transition_duration : no return generator

// Unsupported : gtk_stack_set_transition_type : no return generator

// Unsupported : gtk_stack_set_vhomogeneous : no return generator

// Unsupported : gtk_stack_set_visible_child : no return generator

// Unsupported : gtk_stack_set_visible_child_full : no return generator

// Unsupported : gtk_stack_set_visible_child_name : no return generator

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

	return g
}

func (recv *StackSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_stack_sidebar_set_stack : no return generator

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

	return g
}

func (recv *StackSwitcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_stack_switcher_set_stack : no return generator

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

	return g
}

func (recv *StatusIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// Unsupported : gtk_status_icon_get_gicon : no return generator

// Unsupported : gtk_status_icon_set_from_file : no return generator

// Unsupported : gtk_status_icon_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_set_from_icon_name : no return generator

// Unsupported : gtk_status_icon_set_from_pixbuf : no return generator

// Unsupported : gtk_status_icon_set_from_stock : no return generator

// Unsupported : gtk_status_icon_set_has_tooltip : no return generator

// Unsupported : gtk_status_icon_set_name : no return generator

// Unsupported : gtk_status_icon_set_screen : no return generator

// Unsupported : gtk_status_icon_set_title : no return generator

// Unsupported : gtk_status_icon_set_tooltip_markup : no return generator

// Unsupported : gtk_status_icon_set_tooltip_text : no return generator

// Unsupported : gtk_status_icon_set_visible : no return generator

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

	return g
}

func (recv *Statusbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatusbarNew is a wrapper around the C function gtk_statusbar_new.
func StatusbarNew() *Widget {
	retC := C.gtk_statusbar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContextId is a wrapper around the C function gtk_statusbar_get_context_id.
func (recv *Statusbar) GetContextId(contextDescription string) uint32 {
	c_context_description := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(c_context_description))

	retC := C.gtk_statusbar_get_context_id((*C.GtkStatusbar)(recv.native), c_context_description)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_statusbar_pop : no return generator

// Push is a wrapper around the C function gtk_statusbar_push.
func (recv *Statusbar) Push(contextId uint32, text string) uint32 {
	c_context_id := (C.guint)(contextId)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_statusbar_push((*C.GtkStatusbar)(recv.native), c_context_id, c_text)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_statusbar_remove : no return generator

// Unsupported : gtk_statusbar_remove_all : no return generator

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

	return retGo
}

// Unsupported : gtk_style_apply_default_background : no return generator

// Attach is a wrapper around the C function gtk_style_attach.
func (recv *Style) Attach(window *gdk.Window) *Style {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_style_attach((*C.GtkStyle)(recv.native), c_window)
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_style_copy.
func (recv *Style) Copy() *Style {
	retC := C.gtk_style_copy((*C.GtkStyle)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_detach : no return generator

// Unsupported : gtk_style_get : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_style_property : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_valist : unsupported parameter widget_type : no type generator for GType, GType

// LookupIconSet is a wrapper around the C function gtk_style_lookup_icon_set.
func (recv *Style) LookupIconSet(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_style_lookup_icon_set((*C.GtkStyle)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_style_set_background : no return generator

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

	return g
}

func (recv *StyleContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StyleContextNew is a wrapper around the C function gtk_style_context_new.
func StyleContextNew() *StyleContext {
	retC := C.gtk_style_context_new()
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_context_add_class : no return generator

// Unsupported : gtk_style_context_add_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// Unsupported : gtk_style_context_add_region : no return generator

// Unsupported : gtk_style_context_cancel_animations : no return generator

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_background_color : no return generator

// Unsupported : gtk_style_context_get_border : no return generator

// Unsupported : gtk_style_context_get_border_color : no return generator

// Unsupported : gtk_style_context_get_color : no return generator

// Unsupported : gtk_style_context_get_margin : no return generator

// Unsupported : gtk_style_context_get_padding : no return generator

// Unsupported : gtk_style_context_get_property : no return generator

// GetScreen is a wrapper around the C function gtk_style_context_get_screen.
func (recv *StyleContext) GetScreen() *gdk.Screen {
	retC := C.gtk_style_context_get_screen((*C.GtkStyleContext)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSection is a wrapper around the C function gtk_style_context_get_section.
func (recv *StyleContext) GetSection(property string) *CssSection {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	retC := C.gtk_style_context_get_section((*C.GtkStyleContext)(recv.native), c_property)
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_property : no return generator

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_style_context_invalidate : no return generator

// LookupColor is a wrapper around the C function gtk_style_context_lookup_color.
func (recv *StyleContext) LookupColor(colorName string) (bool, *gdk.RGBA) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkRGBA

	retC := C.gtk_style_context_lookup_color((*C.GtkStyleContext)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// LookupIconSet is a wrapper around the C function gtk_style_context_lookup_icon_set.
func (recv *StyleContext) LookupIconSet(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_style_context_lookup_icon_set((*C.GtkStyleContext)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_context_notify_state_change : no return generator

// Unsupported : gtk_style_context_pop_animatable_region : no return generator

// Unsupported : gtk_style_context_push_animatable_region : no return generator

// Unsupported : gtk_style_context_remove_class : no return generator

// Unsupported : gtk_style_context_remove_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// Unsupported : gtk_style_context_remove_region : no return generator

// Unsupported : gtk_style_context_restore : no return generator

// Unsupported : gtk_style_context_save : no return generator

// Unsupported : gtk_style_context_scroll_animations : no return generator

// Unsupported : gtk_style_context_set_background : no return generator

// Unsupported : gtk_style_context_set_direction : no return generator

// Unsupported : gtk_style_context_set_frame_clock : no return generator

// Unsupported : gtk_style_context_set_junction_sides : no return generator

// Unsupported : gtk_style_context_set_parent : no return generator

// Unsupported : gtk_style_context_set_path : no return generator

// Unsupported : gtk_style_context_set_scale : no return generator

// Unsupported : gtk_style_context_set_screen : no return generator

// Unsupported : gtk_style_context_set_state : no return generator

// Unsupported : gtk_style_context_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

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

	return g
}

func (recv *StyleProperties) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StylePropertiesNew is a wrapper around the C function gtk_style_properties_new.
func StylePropertiesNew() *StyleProperties {
	retC := C.gtk_style_properties_new()
	retGo := StylePropertiesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_properties_clear : no return generator

// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_properties_map_color : no return generator

// Unsupported : gtk_style_properties_merge : no return generator

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_set_property : no return generator

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_properties_unset_property : no return generator

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

	return g
}

func (recv *Switch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_switch_set_active : no return generator

// Unsupported : gtk_switch_set_state : no return generator

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

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableNew is a wrapper around the C function gtk_table_new.
func TableNew(rows uint32, columns uint32, homogeneous bool) *Widget {
	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	c_homogeneous :=
		boolToGboolean(homogeneous)

	retC := C.gtk_table_new(c_rows, c_columns, c_homogeneous)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_table_attach : no return generator

// Unsupported : gtk_table_attach_defaults : no return generator

// GetColSpacing is a wrapper around the C function gtk_table_get_col_spacing.
func (recv *Table) GetColSpacing(column uint32) uint32 {
	c_column := (C.guint)(column)

	retC := C.gtk_table_get_col_spacing((*C.GtkTable)(recv.native), c_column)
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultColSpacing is a wrapper around the C function gtk_table_get_default_col_spacing.
func (recv *Table) GetDefaultColSpacing() uint32 {
	retC := C.gtk_table_get_default_col_spacing((*C.GtkTable)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultRowSpacing is a wrapper around the C function gtk_table_get_default_row_spacing.
func (recv *Table) GetDefaultRowSpacing() uint32 {
	retC := C.gtk_table_get_default_row_spacing((*C.GtkTable)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_table_get_homogeneous.
func (recv *Table) GetHomogeneous() bool {
	retC := C.gtk_table_get_homogeneous((*C.GtkTable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_table_get_row_spacing.
func (recv *Table) GetRowSpacing(row uint32) uint32 {
	c_row := (C.guint)(row)

	retC := C.gtk_table_get_row_spacing((*C.GtkTable)(recv.native), c_row)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_table_get_size : unsupported parameter rows : no type generator for guint, guint*

// Unsupported : gtk_table_resize : no return generator

// Unsupported : gtk_table_set_col_spacing : no return generator

// Unsupported : gtk_table_set_col_spacings : no return generator

// Unsupported : gtk_table_set_homogeneous : no return generator

// Unsupported : gtk_table_set_row_spacing : no return generator

// Unsupported : gtk_table_set_row_spacings : no return generator

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

	return g
}

func (recv *TearoffMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TearoffMenuItemNew is a wrapper around the C function gtk_tearoff_menu_item_new.
func TearoffMenuItemNew() *Widget {
	retC := C.gtk_tearoff_menu_item_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *TextBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextBufferNew is a wrapper around the C function gtk_text_buffer_new.
func TextBufferNew(table *TextTagTable) *TextBuffer {
	c_table := (*C.GtkTextTagTable)(table.ToC())

	retC := C.gtk_text_buffer_new(c_table)
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_add_mark : no return generator

// Unsupported : gtk_text_buffer_add_selection_clipboard : no return generator

// Unsupported : gtk_text_buffer_apply_tag : no return generator

// Unsupported : gtk_text_buffer_apply_tag_by_name : no return generator

// Unsupported : gtk_text_buffer_begin_user_action : no return generator

// Unsupported : gtk_text_buffer_copy_clipboard : no return generator

// CreateChildAnchor is a wrapper around the C function gtk_text_buffer_create_child_anchor.
func (recv *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_buffer_create_child_anchor((*C.GtkTextBuffer)(recv.native), c_iter)
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateMark is a wrapper around the C function gtk_text_buffer_create_mark.
func (recv *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity bool) *TextMark {
	c_mark_name := C.CString(markName)
	defer C.free(unsafe.Pointer(c_mark_name))

	c_where := (*C.GtkTextIter)(where.ToC())

	c_left_gravity :=
		boolToGboolean(leftGravity)

	retC := C.gtk_text_buffer_create_mark((*C.GtkTextBuffer)(recv.native), c_mark_name, c_where, c_left_gravity)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_create_tag : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_cut_clipboard : no return generator

// Unsupported : gtk_text_buffer_delete : no return generator

// DeleteInteractive is a wrapper around the C function gtk_text_buffer_delete_interactive.
func (recv *TextBuffer) DeleteInteractive(startIter *TextIter, endIter *TextIter, defaultEditable bool) bool {
	c_start_iter := (*C.GtkTextIter)(startIter.ToC())

	c_end_iter := (*C.GtkTextIter)(endIter.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_delete_interactive((*C.GtkTextBuffer)(recv.native), c_start_iter, c_end_iter, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_delete_mark : no return generator

// Unsupported : gtk_text_buffer_delete_mark_by_name : no return generator

// DeleteSelection is a wrapper around the C function gtk_text_buffer_delete_selection.
func (recv *TextBuffer) DeleteSelection(interactive bool, defaultEditable bool) bool {
	c_interactive :=
		boolToGboolean(interactive)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_delete_selection((*C.GtkTextBuffer)(recv.native), c_interactive, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_deserialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_get_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_set_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_end_user_action : no return generator

// Unsupported : gtk_text_buffer_get_bounds : no return generator

// GetCharCount is a wrapper around the C function gtk_text_buffer_get_char_count.
func (recv *TextBuffer) GetCharCount() int32 {
	retC := C.gtk_text_buffer_get_char_count((*C.GtkTextBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_text_buffer_get_deserialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// Unsupported : gtk_text_buffer_get_end_iter : no return generator

// GetInsert is a wrapper around the C function gtk_text_buffer_get_insert.
func (recv *TextBuffer) GetInsert() *TextMark {
	retC := C.gtk_text_buffer_get_insert((*C.GtkTextBuffer)(recv.native))
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_get_iter_at_child_anchor : no return generator

// Unsupported : gtk_text_buffer_get_iter_at_line : no return generator

// Unsupported : gtk_text_buffer_get_iter_at_line_index : no return generator

// Unsupported : gtk_text_buffer_get_iter_at_line_offset : no return generator

// Unsupported : gtk_text_buffer_get_iter_at_mark : no return generator

// Unsupported : gtk_text_buffer_get_iter_at_offset : no return generator

// GetLineCount is a wrapper around the C function gtk_text_buffer_get_line_count.
func (recv *TextBuffer) GetLineCount() int32 {
	retC := C.gtk_text_buffer_get_line_count((*C.GtkTextBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMark is a wrapper around the C function gtk_text_buffer_get_mark.
func (recv *TextBuffer) GetMark(name string) *TextMark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_buffer_get_mark((*C.GtkTextBuffer)(recv.native), c_name)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetModified is a wrapper around the C function gtk_text_buffer_get_modified.
func (recv *TextBuffer) GetModified() bool {
	retC := C.gtk_text_buffer_get_modified((*C.GtkTextBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectionBound is a wrapper around the C function gtk_text_buffer_get_selection_bound.
func (recv *TextBuffer) GetSelectionBound() *TextMark {
	retC := C.gtk_text_buffer_get_selection_bound((*C.GtkTextBuffer)(recv.native))
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionBounds is a wrapper around the C function gtk_text_buffer_get_selection_bounds.
func (recv *TextBuffer) GetSelectionBounds() (bool, *TextIter, *TextIter) {
	var c_start C.GtkTextIter

	var c_end C.GtkTextIter

	retC := C.gtk_text_buffer_get_selection_bounds((*C.GtkTextBuffer)(recv.native), &c_start, &c_end)
	retGo := retC == C.TRUE

	start := TextIterNewFromC(unsafe.Pointer(&c_start))

	end := TextIterNewFromC(unsafe.Pointer(&c_end))

	return retGo, start, end
}

// Unsupported : gtk_text_buffer_get_serialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// GetSlice is a wrapper around the C function gtk_text_buffer_get_slice.
func (recv *TextBuffer) GetSlice(start *TextIter, end *TextIter, includeHiddenChars bool) string {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_include_hidden_chars :=
		boolToGboolean(includeHiddenChars)

	retC := C.gtk_text_buffer_get_slice((*C.GtkTextBuffer)(recv.native), c_start, c_end, c_include_hidden_chars)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_get_start_iter : no return generator

// GetTagTable is a wrapper around the C function gtk_text_buffer_get_tag_table.
func (recv *TextBuffer) GetTagTable() *TextTagTable {
	retC := C.gtk_text_buffer_get_tag_table((*C.GtkTextBuffer)(recv.native))
	retGo := TextTagTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetText is a wrapper around the C function gtk_text_buffer_get_text.
func (recv *TextBuffer) GetText(start *TextIter, end *TextIter, includeHiddenChars bool) string {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_include_hidden_chars :=
		boolToGboolean(includeHiddenChars)

	retC := C.gtk_text_buffer_get_text((*C.GtkTextBuffer)(recv.native), c_start, c_end, c_include_hidden_chars)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_insert : no return generator

// Unsupported : gtk_text_buffer_insert_at_cursor : no return generator

// Unsupported : gtk_text_buffer_insert_child_anchor : no return generator

// InsertInteractive is a wrapper around the C function gtk_text_buffer_insert_interactive.
func (recv *TextBuffer) InsertInteractive(iter *TextIter, text string, len int32, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_interactive((*C.GtkTextBuffer)(recv.native), c_iter, c_text, c_len, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// InsertInteractiveAtCursor is a wrapper around the C function gtk_text_buffer_insert_interactive_at_cursor.
func (recv *TextBuffer) InsertInteractiveAtCursor(text string, len int32, defaultEditable bool) bool {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_interactive_at_cursor((*C.GtkTextBuffer)(recv.native), c_text, c_len, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_insert_markup : no return generator

// Unsupported : gtk_text_buffer_insert_pixbuf : no return generator

// Unsupported : gtk_text_buffer_insert_range : no return generator

// InsertRangeInteractive is a wrapper around the C function gtk_text_buffer_insert_range_interactive.
func (recv *TextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_range_interactive((*C.GtkTextBuffer)(recv.native), c_iter, c_start, c_end, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_insert_with_tags : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_insert_with_tags_by_name : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_move_mark : no return generator

// Unsupported : gtk_text_buffer_move_mark_by_name : no return generator

// Unsupported : gtk_text_buffer_paste_clipboard : no return generator

// Unsupported : gtk_text_buffer_place_cursor : no return generator

// Unsupported : gtk_text_buffer_register_deserialize_format : unsupported parameter function : no type generator for TextBufferDeserializeFunc, GtkTextBufferDeserializeFunc

// Unsupported : gtk_text_buffer_register_deserialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_register_serialize_format : unsupported parameter function : no type generator for TextBufferSerializeFunc, GtkTextBufferSerializeFunc

// Unsupported : gtk_text_buffer_register_serialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_remove_all_tags : no return generator

// Unsupported : gtk_text_buffer_remove_selection_clipboard : no return generator

// Unsupported : gtk_text_buffer_remove_tag : no return generator

// Unsupported : gtk_text_buffer_remove_tag_by_name : no return generator

// Unsupported : gtk_text_buffer_select_range : no return generator

// Unsupported : gtk_text_buffer_serialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_set_modified : no return generator

// Unsupported : gtk_text_buffer_set_text : no return generator

// Unsupported : gtk_text_buffer_unregister_deserialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_serialize_format : unsupported parameter format : Blacklisted record : GdkAtom

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

	return g
}

func (recv *TextChildAnchor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextChildAnchorNew is a wrapper around the C function gtk_text_child_anchor_new.
func TextChildAnchorNew() *TextChildAnchor {
	retC := C.gtk_text_child_anchor_new()
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeleted is a wrapper around the C function gtk_text_child_anchor_get_deleted.
func (recv *TextChildAnchor) GetDeleted() bool {
	retC := C.gtk_text_child_anchor_get_deleted((*C.GtkTextChildAnchor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidgets is a wrapper around the C function gtk_text_child_anchor_get_widgets.
func (recv *TextChildAnchor) GetWidgets() *glib.List {
	retC := C.gtk_text_child_anchor_get_widgets((*C.GtkTextChildAnchor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *TextMark) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetBuffer is a wrapper around the C function gtk_text_mark_get_buffer.
func (recv *TextMark) GetBuffer() *TextBuffer {
	retC := C.gtk_text_mark_get_buffer((*C.GtkTextMark)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeleted is a wrapper around the C function gtk_text_mark_get_deleted.
func (recv *TextMark) GetDeleted() bool {
	retC := C.gtk_text_mark_get_deleted((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLeftGravity is a wrapper around the C function gtk_text_mark_get_left_gravity.
func (recv *TextMark) GetLeftGravity() bool {
	retC := C.gtk_text_mark_get_left_gravity((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function gtk_text_mark_get_name.
func (recv *TextMark) GetName() string {
	retC := C.gtk_text_mark_get_name((*C.GtkTextMark)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_text_mark_get_visible.
func (recv *TextMark) GetVisible() bool {
	retC := C.gtk_text_mark_get_visible((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_mark_set_visible : no return generator

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

	return retGo
}

// Unsupported : gtk_text_tag_changed : no return generator

// Unsupported : gtk_text_tag_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetPriority is a wrapper around the C function gtk_text_tag_get_priority.
func (recv *TextTag) GetPriority() int32 {
	retC := C.gtk_text_tag_get_priority((*C.GtkTextTag)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_text_tag_set_priority : no return generator

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

	return g
}

func (recv *TextTagTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextTagTableNew is a wrapper around the C function gtk_text_tag_table_new.
func TextTagTableNew() *TextTagTable {
	retC := C.gtk_text_tag_table_new()
	retGo := TextTagTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function gtk_text_tag_table_add.
func (recv *TextTagTable) Add(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	retC := C.gtk_text_tag_table_add((*C.GtkTextTagTable)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_tag_table_foreach : unsupported parameter func : no type generator for TextTagTableForeach, GtkTextTagTableForeach

// GetSize is a wrapper around the C function gtk_text_tag_table_get_size.
func (recv *TextTagTable) GetSize() int32 {
	retC := C.gtk_text_tag_table_get_size((*C.GtkTextTagTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Lookup is a wrapper around the C function gtk_text_tag_table_lookup.
func (recv *TextTagTable) Lookup(name string) *TextTag {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_tag_table_lookup((*C.GtkTextTagTable)(recv.native), c_name)
	retGo := TextTagNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_tag_table_remove : no return generator

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

	return g
}

func (recv *TextView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextViewNew is a wrapper around the C function gtk_text_view_new.
func TextViewNew() *Widget {
	retC := C.gtk_text_view_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextViewNewWithBuffer is a wrapper around the C function gtk_text_view_new_with_buffer.
func TextViewNewWithBuffer(buffer *TextBuffer) *Widget {
	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	retC := C.gtk_text_view_new_with_buffer(c_buffer)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_view_add_child_at_anchor : no return generator

// Unsupported : gtk_text_view_add_child_in_window : no return generator

// BackwardDisplayLine is a wrapper around the C function gtk_text_view_backward_display_line.
func (recv *TextView) BackwardDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_backward_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardDisplayLineStart is a wrapper around the C function gtk_text_view_backward_display_line_start.
func (recv *TextView) BackwardDisplayLineStart(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_backward_display_line_start((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_view_buffer_to_window_coords : unsupported parameter window_x : no type generator for gint, gint*

// ForwardDisplayLine is a wrapper around the C function gtk_text_view_forward_display_line.
func (recv *TextView) ForwardDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_forward_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardDisplayLineEnd is a wrapper around the C function gtk_text_view_forward_display_line_end.
func (recv *TextView) ForwardDisplayLineEnd(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_forward_display_line_end((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// GetBorderWindowSize is a wrapper around the C function gtk_text_view_get_border_window_size.
func (recv *TextView) GetBorderWindowSize(type_ TextWindowType) int32 {
	c_type := (C.GtkTextWindowType)(type_)

	retC := C.gtk_text_view_get_border_window_size((*C.GtkTextView)(recv.native), c_type)
	retGo := (int32)(retC)

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_text_view_get_buffer.
func (recv *TextView) GetBuffer() *TextBuffer {
	retC := C.gtk_text_view_get_buffer((*C.GtkTextView)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_view_get_cursor_locations : unsupported parameter strong : Blacklisted record : GdkRectangle

// GetCursorVisible is a wrapper around the C function gtk_text_view_get_cursor_visible.
func (recv *TextView) GetCursorVisible() bool {
	retC := C.gtk_text_view_get_cursor_visible((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDefaultAttributes is a wrapper around the C function gtk_text_view_get_default_attributes.
func (recv *TextView) GetDefaultAttributes() *TextAttributes {
	retC := C.gtk_text_view_get_default_attributes((*C.GtkTextView)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEditable is a wrapper around the C function gtk_text_view_get_editable.
func (recv *TextView) GetEditable() bool {
	retC := C.gtk_text_view_get_editable((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIndent is a wrapper around the C function gtk_text_view_get_indent.
func (recv *TextView) GetIndent() int32 {
	retC := C.gtk_text_view_get_indent((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIterAtLocation is a wrapper around the C function gtk_text_view_get_iter_at_location.
func (recv *TextView) GetIterAtLocation(x int32, y int32) (bool, *TextIter) {
	var c_iter C.GtkTextIter

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_text_view_get_iter_at_location((*C.GtkTextView)(recv.native), &c_iter, c_x, c_y)
	retGo := retC == C.TRUE

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Unsupported : gtk_text_view_get_iter_at_position : unsupported parameter trailing : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_iter_location : unsupported parameter location : Blacklisted record : GdkRectangle

// GetJustification is a wrapper around the C function gtk_text_view_get_justification.
func (recv *TextView) GetJustification() Justification {
	retC := C.gtk_text_view_get_justification((*C.GtkTextView)(recv.native))
	retGo := (Justification)(retC)

	return retGo
}

// GetLeftMargin is a wrapper around the C function gtk_text_view_get_left_margin.
func (recv *TextView) GetLeftMargin() int32 {
	retC := C.gtk_text_view_get_left_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_text_view_get_line_at_y : unsupported parameter line_top : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_line_yrange : unsupported parameter y : no type generator for gint, gint*

// GetPixelsAboveLines is a wrapper around the C function gtk_text_view_get_pixels_above_lines.
func (recv *TextView) GetPixelsAboveLines() int32 {
	retC := C.gtk_text_view_get_pixels_above_lines((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixelsBelowLines is a wrapper around the C function gtk_text_view_get_pixels_below_lines.
func (recv *TextView) GetPixelsBelowLines() int32 {
	retC := C.gtk_text_view_get_pixels_below_lines((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixelsInsideWrap is a wrapper around the C function gtk_text_view_get_pixels_inside_wrap.
func (recv *TextView) GetPixelsInsideWrap() int32 {
	retC := C.gtk_text_view_get_pixels_inside_wrap((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRightMargin is a wrapper around the C function gtk_text_view_get_right_margin.
func (recv *TextView) GetRightMargin() int32 {
	retC := C.gtk_text_view_get_right_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTabs is a wrapper around the C function gtk_text_view_get_tabs.
func (recv *TextView) GetTabs() *pango.TabArray {
	retC := C.gtk_text_view_get_tabs((*C.GtkTextView)(recv.native))
	retGo := pango.TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// GetWindow is a wrapper around the C function gtk_text_view_get_window.
func (recv *TextView) GetWindow(win TextWindowType) *gdk.Window {
	c_win := (C.GtkTextWindowType)(win)

	retC := C.gtk_text_view_get_window((*C.GtkTextView)(recv.native), c_win)
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowType is a wrapper around the C function gtk_text_view_get_window_type.
func (recv *TextView) GetWindowType(window *gdk.Window) TextWindowType {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_text_view_get_window_type((*C.GtkTextView)(recv.native), c_window)
	retGo := (TextWindowType)(retC)

	return retGo
}

// GetWrapMode is a wrapper around the C function gtk_text_view_get_wrap_mode.
func (recv *TextView) GetWrapMode() WrapMode {
	retC := C.gtk_text_view_get_wrap_mode((*C.GtkTextView)(recv.native))
	retGo := (WrapMode)(retC)

	return retGo
}

// Unsupported : gtk_text_view_move_child : no return generator

// MoveMarkOnscreen is a wrapper around the C function gtk_text_view_move_mark_onscreen.
func (recv *TextView) MoveMarkOnscreen(mark *TextMark) bool {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	retC := C.gtk_text_view_move_mark_onscreen((*C.GtkTextView)(recv.native), c_mark)
	retGo := retC == C.TRUE

	return retGo
}

// MoveVisually is a wrapper around the C function gtk_text_view_move_visually.
func (recv *TextView) MoveVisually(iter *TextIter, count int32) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_count := (C.gint)(count)

	retC := C.gtk_text_view_move_visually((*C.GtkTextView)(recv.native), c_iter, c_count)
	retGo := retC == C.TRUE

	return retGo
}

// PlaceCursorOnscreen is a wrapper around the C function gtk_text_view_place_cursor_onscreen.
func (recv *TextView) PlaceCursorOnscreen() bool {
	retC := C.gtk_text_view_place_cursor_onscreen((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_view_reset_cursor_blink : no return generator

// Unsupported : gtk_text_view_reset_im_context : no return generator

// Unsupported : gtk_text_view_scroll_mark_onscreen : no return generator

// ScrollToIter is a wrapper around the C function gtk_text_view_scroll_to_iter.
func (recv *TextView) ScrollToIter(iter *TextIter, withinMargin float64, useAlign bool, xalign float64, yalign float64) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_within_margin := (C.gdouble)(withinMargin)

	c_use_align :=
		boolToGboolean(useAlign)

	c_xalign := (C.gdouble)(xalign)

	c_yalign := (C.gdouble)(yalign)

	retC := C.gtk_text_view_scroll_to_iter((*C.GtkTextView)(recv.native), c_iter, c_within_margin, c_use_align, c_xalign, c_yalign)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_view_scroll_to_mark : no return generator

// Unsupported : gtk_text_view_set_accepts_tab : no return generator

// Unsupported : gtk_text_view_set_border_window_size : no return generator

// Unsupported : gtk_text_view_set_bottom_margin : no return generator

// Unsupported : gtk_text_view_set_buffer : no return generator

// Unsupported : gtk_text_view_set_cursor_visible : no return generator

// Unsupported : gtk_text_view_set_editable : no return generator

// Unsupported : gtk_text_view_set_indent : no return generator

// Unsupported : gtk_text_view_set_input_hints : no return generator

// Unsupported : gtk_text_view_set_input_purpose : no return generator

// Unsupported : gtk_text_view_set_justification : no return generator

// Unsupported : gtk_text_view_set_left_margin : no return generator

// Unsupported : gtk_text_view_set_monospace : no return generator

// Unsupported : gtk_text_view_set_overwrite : no return generator

// Unsupported : gtk_text_view_set_pixels_above_lines : no return generator

// Unsupported : gtk_text_view_set_pixels_below_lines : no return generator

// Unsupported : gtk_text_view_set_pixels_inside_wrap : no return generator

// Unsupported : gtk_text_view_set_right_margin : no return generator

// Unsupported : gtk_text_view_set_tabs : no return generator

// Unsupported : gtk_text_view_set_top_margin : no return generator

// Unsupported : gtk_text_view_set_wrap_mode : no return generator

// StartsDisplayLine is a wrapper around the C function gtk_text_view_starts_display_line.
func (recv *TextView) StartsDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_starts_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_view_window_to_buffer_coords : unsupported parameter buffer_x : no type generator for gint, gint*

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

	return g
}

func (recv *ThemingEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_background_color : no return generator

// Unsupported : gtk_theming_engine_get_border : no return generator

// Unsupported : gtk_theming_engine_get_border_color : no return generator

// Unsupported : gtk_theming_engine_get_color : no return generator

// Unsupported : gtk_theming_engine_get_margin : no return generator

// Unsupported : gtk_theming_engine_get_padding : no return generator

// Unsupported : gtk_theming_engine_get_property : no return generator

// GetScreen is a wrapper around the C function gtk_theming_engine_get_screen.
func (recv *ThemingEngine) GetScreen() *gdk.Screen {
	retC := C.gtk_theming_engine_get_screen((*C.GtkThemingEngine)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_style_property : no return generator

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_theming_engine_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

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

	return g
}

func (recv *ToggleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_toggle_action_set_active : no return generator

// Unsupported : gtk_toggle_action_set_draw_as_radio : no return generator

// Unsupported : gtk_toggle_action_toggled : no return generator

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

	return g
}

func (recv *ToggleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonNew is a wrapper around the C function gtk_toggle_button_new.
func ToggleButtonNew() *Widget {
	retC := C.gtk_toggle_button_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithLabel is a wrapper around the C function gtk_toggle_button_new_with_label.
func ToggleButtonNewWithLabel(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_label(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithMnemonic is a wrapper around the C function gtk_toggle_button_new_with_mnemonic.
func ToggleButtonNewWithMnemonic(label string) *Widget {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_mnemonic(c_label)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_button_get_active.
func (recv *ToggleButton) GetActive() bool {
	retC := C.gtk_toggle_button_get_active((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInconsistent is a wrapper around the C function gtk_toggle_button_get_inconsistent.
func (recv *ToggleButton) GetInconsistent() bool {
	retC := C.gtk_toggle_button_get_inconsistent((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMode is a wrapper around the C function gtk_toggle_button_get_mode.
func (recv *ToggleButton) GetMode() bool {
	retC := C.gtk_toggle_button_get_mode((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_toggle_button_set_active : no return generator

// Unsupported : gtk_toggle_button_set_inconsistent : no return generator

// Unsupported : gtk_toggle_button_set_mode : no return generator

// Unsupported : gtk_toggle_button_toggled : no return generator

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

	return g
}

func (recv *ToggleToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_toggle_tool_button_set_active : no return generator

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

	return g
}

func (recv *ToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tool_button_set_icon_name : no return generator

// Unsupported : gtk_tool_button_set_icon_widget : no return generator

// Unsupported : gtk_tool_button_set_label : no return generator

// Unsupported : gtk_tool_button_set_label_widget : no return generator

// Unsupported : gtk_tool_button_set_stock_id : no return generator

// Unsupported : gtk_tool_button_set_use_underline : no return generator

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

	return g
}

func (recv *ToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tool_item_get_icon_size : no return generator

// Unsupported : gtk_tool_item_rebuild_menu : no return generator

// Unsupported : gtk_tool_item_set_expand : no return generator

// Unsupported : gtk_tool_item_set_homogeneous : no return generator

// Unsupported : gtk_tool_item_set_is_important : no return generator

// Unsupported : gtk_tool_item_set_proxy_menu_item : no return generator

// Unsupported : gtk_tool_item_set_tooltip_markup : no return generator

// Unsupported : gtk_tool_item_set_tooltip_text : no return generator

// Unsupported : gtk_tool_item_set_use_drag_window : no return generator

// Unsupported : gtk_tool_item_set_visible_horizontal : no return generator

// Unsupported : gtk_tool_item_set_visible_vertical : no return generator

// Unsupported : gtk_tool_item_toolbar_reconfigured : no return generator

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

	return g
}

func (recv *ToolItemGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tool_item_group_insert : no return generator

// Unsupported : gtk_tool_item_group_set_collapsed : no return generator

// Unsupported : gtk_tool_item_group_set_ellipsize : no return generator

// Unsupported : gtk_tool_item_group_set_header_relief : no return generator

// Unsupported : gtk_tool_item_group_set_item_position : no return generator

// Unsupported : gtk_tool_item_group_set_label : no return generator

// Unsupported : gtk_tool_item_group_set_label_widget : no return generator

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

	return g
}

func (recv *ToolPalette) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tool_palette_add_drag_dest : no return generator

// Unsupported : gtk_tool_palette_get_icon_size : no return generator

// Unsupported : gtk_tool_palette_set_drag_source : no return generator

// Unsupported : gtk_tool_palette_set_exclusive : no return generator

// Unsupported : gtk_tool_palette_set_expand : no return generator

// Unsupported : gtk_tool_palette_set_group_position : no return generator

// Unsupported : gtk_tool_palette_set_icon_size : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tool_palette_set_style : no return generator

// Unsupported : gtk_tool_palette_unset_icon_size : no return generator

// Unsupported : gtk_tool_palette_unset_style : no return generator

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

	return g
}

func (recv *Toolbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolbarNew is a wrapper around the C function gtk_toolbar_new.
func ToolbarNew() *Widget {
	retC := C.gtk_toolbar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconSize is a wrapper around the C function gtk_toolbar_get_icon_size.
func (recv *Toolbar) GetIconSize() IconSize {
	retC := C.gtk_toolbar_get_icon_size((*C.GtkToolbar)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function gtk_toolbar_get_style.
func (recv *Toolbar) GetStyle() ToolbarStyle {
	retC := C.gtk_toolbar_get_style((*C.GtkToolbar)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// Unsupported : gtk_toolbar_insert : no return generator

// Unsupported : gtk_toolbar_set_drop_highlight_item : no return generator

// Unsupported : gtk_toolbar_set_icon_size : no return generator

// Unsupported : gtk_toolbar_set_show_arrow : no return generator

// Unsupported : gtk_toolbar_set_style : no return generator

// Unsupported : gtk_toolbar_unset_icon_size : no return generator

// Unsupported : gtk_toolbar_unset_style : no return generator

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

	return g
}

func (recv *Tooltip) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tooltip_set_custom : no return generator

// Unsupported : gtk_tooltip_set_icon : no return generator

// Unsupported : gtk_tooltip_set_icon_from_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_tooltip_set_icon_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_icon_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_markup : no return generator

// Unsupported : gtk_tooltip_set_text : no return generator

// Unsupported : gtk_tooltip_set_tip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

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

	return g
}

func (recv *ToplevelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetChildren is a wrapper around the C function gtk_toplevel_accessible_get_children.
func (recv *ToplevelAccessible) GetChildren() *glib.List {
	retC := C.gtk_toplevel_accessible_get_children((*C.GtkToplevelAccessible)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
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

	return g
}

func (recv *TreeModelFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tree_model_filter_clear_cache : no return generator

// Unsupported : gtk_tree_model_filter_convert_iter_to_child_iter : no return generator

// Unsupported : gtk_tree_model_filter_get_model : no return generator

// Unsupported : gtk_tree_model_filter_refilter : no return generator

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter types : no param type

// Unsupported : gtk_tree_model_filter_set_visible_column : no return generator

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc, GtkTreeModelFilterVisibleFunc

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

	return g
}

func (recv *TreeModelSort) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tree_model_sort_clear_cache : no return generator

// ConvertChildIterToIter is a wrapper around the C function gtk_tree_model_sort_convert_child_iter_to_iter.
func (recv *TreeModelSort) ConvertChildIterToIter(childIter *TreeIter) (bool, *TreeIter) {
	var c_sort_iter C.GtkTreeIter

	c_child_iter := (*C.GtkTreeIter)(childIter.ToC())

	retC := C.gtk_tree_model_sort_convert_child_iter_to_iter((*C.GtkTreeModelSort)(recv.native), &c_sort_iter, c_child_iter)
	retGo := retC == C.TRUE

	sortIter := TreeIterNewFromC(unsafe.Pointer(&c_sort_iter))

	return retGo, sortIter
}

// ConvertChildPathToPath is a wrapper around the C function gtk_tree_model_sort_convert_child_path_to_path.
func (recv *TreeModelSort) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	c_child_path := (*C.GtkTreePath)(childPath.ToC())

	retC := C.gtk_tree_model_sort_convert_child_path_to_path((*C.GtkTreeModelSort)(recv.native), c_child_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_sort_convert_iter_to_child_iter : no return generator

// ConvertPathToChildPath is a wrapper around the C function gtk_tree_model_sort_convert_path_to_child_path.
func (recv *TreeModelSort) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	c_sorted_path := (*C.GtkTreePath)(sortedPath.ToC())

	retC := C.gtk_tree_model_sort_convert_path_to_child_path((*C.GtkTreeModelSort)(recv.native), c_sorted_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_sort_get_model : no return generator

// Unsupported : gtk_tree_model_sort_reset_default_sort_func : no return generator

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

	return g
}

func (recv *TreeSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetMode is a wrapper around the C function gtk_tree_selection_get_mode.
func (recv *TreeSelection) GetMode() SelectionMode {
	retC := C.gtk_tree_selection_get_mode((*C.GtkTreeSelection)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Unsupported : gtk_tree_selection_get_select_function : no return generator

// Unsupported : gtk_tree_selection_get_selected : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_selection_get_selected_rows : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// GetTreeView is a wrapper around the C function gtk_tree_selection_get_tree_view.
func (recv *TreeSelection) GetTreeView() *TreeView {
	retC := C.gtk_tree_selection_get_tree_view((*C.GtkTreeSelection)(recv.native))
	retGo := TreeViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUserData is a wrapper around the C function gtk_tree_selection_get_user_data.
func (recv *TreeSelection) GetUserData() uintptr {
	retC := C.gtk_tree_selection_get_user_data((*C.GtkTreeSelection)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// IterIsSelected is a wrapper around the C function gtk_tree_selection_iter_is_selected.
func (recv *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_selection_iter_is_selected((*C.GtkTreeSelection)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// PathIsSelected is a wrapper around the C function gtk_tree_selection_path_is_selected.
func (recv *TreeSelection) PathIsSelected(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_selection_path_is_selected((*C.GtkTreeSelection)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_selection_select_all : no return generator

// Unsupported : gtk_tree_selection_select_iter : no return generator

// Unsupported : gtk_tree_selection_select_path : no return generator

// Unsupported : gtk_tree_selection_select_range : no return generator

// Unsupported : gtk_tree_selection_selected_foreach : unsupported parameter func : no type generator for TreeSelectionForeachFunc, GtkTreeSelectionForeachFunc

// Unsupported : gtk_tree_selection_set_mode : no return generator

// Unsupported : gtk_tree_selection_set_select_function : unsupported parameter func : no type generator for TreeSelectionFunc, GtkTreeSelectionFunc

// Unsupported : gtk_tree_selection_unselect_all : no return generator

// Unsupported : gtk_tree_selection_unselect_iter : no return generator

// Unsupported : gtk_tree_selection_unselect_path : no return generator

// Unsupported : gtk_tree_selection_unselect_range : no return generator

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

	return g
}

func (recv *TreeStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_append : no return generator

// Unsupported : gtk_tree_store_clear : no return generator

// Unsupported : gtk_tree_store_insert : no return generator

// Unsupported : gtk_tree_store_insert_after : no return generator

// Unsupported : gtk_tree_store_insert_before : no return generator

// Unsupported : gtk_tree_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_insert_with_valuesv : unsupported parameter columns : no param type

// IsAncestor is a wrapper around the C function gtk_tree_store_is_ancestor.
func (recv *TreeStore) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_descendant := (*C.GtkTreeIter)(descendant.ToC())

	retC := C.gtk_tree_store_is_ancestor((*C.GtkTreeStore)(recv.native), c_iter, c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// IterDepth is a wrapper around the C function gtk_tree_store_iter_depth.
func (recv *TreeStore) IterDepth(iter *TreeIter) int32 {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_iter_depth((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_store_move_after : no return generator

// Unsupported : gtk_tree_store_move_before : no return generator

// Unsupported : gtk_tree_store_prepend : no return generator

// Remove is a wrapper around the C function gtk_tree_store_remove.
func (recv *TreeStore) Remove(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_remove((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_tree_store_set_value : no return generator

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_store_swap : no return generator

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

	return g
}

func (recv *TreeView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TreeViewNew is a wrapper around the C function gtk_tree_view_new.
func TreeViewNew() *Widget {
	retC := C.gtk_tree_view_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// AppendColumn is a wrapper around the C function gtk_tree_view_append_column.
func (recv *TreeView) AppendColumn(column *TreeViewColumn) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	retC := C.gtk_tree_view_append_column((*C.GtkTreeView)(recv.native), c_column)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_collapse_all : no return generator

// CollapseRow is a wrapper around the C function gtk_tree_view_collapse_row.
func (recv *TreeView) CollapseRow(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_collapse_row((*C.GtkTreeView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_columns_autosize : no return generator

// Unsupported : gtk_tree_view_convert_bin_window_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_bin_window_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// CreateRowDragIcon is a wrapper around the C function gtk_tree_view_create_row_drag_icon.
func (recv *TreeView) CreateRowDragIcon(path *TreePath) *cairo.Surface {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_create_row_drag_icon((*C.GtkTreeView)(recv.native), c_path)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_expand_all : no return generator

// ExpandRow is a wrapper around the C function gtk_tree_view_expand_row.
func (recv *TreeView) ExpandRow(path *TreePath, openAll bool) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_open_all :=
		boolToGboolean(openAll)

	retC := C.gtk_tree_view_expand_row((*C.GtkTreeView)(recv.native), c_path, c_open_all)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_expand_to_path : no return generator

// Unsupported : gtk_tree_view_get_background_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetBinWindow is a wrapper around the C function gtk_tree_view_get_bin_window.
func (recv *TreeView) GetBinWindow() *gdk.Window {
	retC := C.gtk_tree_view_get_bin_window((*C.GtkTreeView)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_cell_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetColumn is a wrapper around the C function gtk_tree_view_get_column.
func (recv *TreeView) GetColumn(n int32) *TreeViewColumn {
	c_n := (C.gint)(n)

	retC := C.gtk_tree_view_get_column((*C.GtkTreeView)(recv.native), c_n)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetColumns is a wrapper around the C function gtk_tree_view_get_columns.
func (recv *TreeView) GetColumns() *glib.List {
	retC := C.gtk_tree_view_get_columns((*C.GtkTreeView)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_dest_row_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_drag_dest_row : unsupported parameter path : record with indirection level of 2

// GetEnableSearch is a wrapper around the C function gtk_tree_view_get_enable_search.
func (recv *TreeView) GetEnableSearch() bool {
	retC := C.gtk_tree_view_get_enable_search((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetExpanderColumn is a wrapper around the C function gtk_tree_view_get_expander_column.
func (recv *TreeView) GetExpanderColumn() *TreeViewColumn {
	retC := C.gtk_tree_view_get_expander_column((*C.GtkTreeView)(recv.native))
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_tree_view_get_hadjustment.
func (recv *TreeView) GetHadjustment() *Adjustment {
	retC := C.gtk_tree_view_get_hadjustment((*C.GtkTreeView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHeadersVisible is a wrapper around the C function gtk_tree_view_get_headers_visible.
func (recv *TreeView) GetHeadersVisible() bool {
	retC := C.gtk_tree_view_get_headers_visible((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_model : no return generator

// Unsupported : gtk_tree_view_get_path_at_pos : unsupported parameter path : record with indirection level of 2

// GetReorderable is a wrapper around the C function gtk_tree_view_get_reorderable.
func (recv *TreeView) GetReorderable() bool {
	retC := C.gtk_tree_view_get_reorderable((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

// GetRulesHint is a wrapper around the C function gtk_tree_view_get_rules_hint.
func (recv *TreeView) GetRulesHint() bool {
	retC := C.gtk_tree_view_get_rules_hint((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSearchColumn is a wrapper around the C function gtk_tree_view_get_search_column.
func (recv *TreeView) GetSearchColumn() int32 {
	retC := C.gtk_tree_view_get_search_column((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_get_search_equal_func : no return generator

// Unsupported : gtk_tree_view_get_search_position_func : no return generator

// GetSelection is a wrapper around the C function gtk_tree_view_get_selection.
func (recv *TreeView) GetSelection() *TreeSelection {
	retC := C.gtk_tree_view_get_selection((*C.GtkTreeView)(recv.native))
	retGo := TreeSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// GetVadjustment is a wrapper around the C function gtk_tree_view_get_vadjustment.
func (recv *TreeView) GetVadjustment() *Adjustment {
	retC := C.gtk_tree_view_get_vadjustment((*C.GtkTreeView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// InsertColumn is a wrapper around the C function gtk_tree_view_insert_column.
func (recv *TreeView) InsertColumn(column *TreeViewColumn, position int32) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_tree_view_insert_column((*C.GtkTreeView)(recv.native), c_column, c_position)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_insert_column_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_insert_column_with_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_tree_view_is_blank_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_map_expanded_rows : unsupported parameter func : no type generator for TreeViewMappingFunc, GtkTreeViewMappingFunc

// Unsupported : gtk_tree_view_move_column_after : no return generator

// RemoveColumn is a wrapper around the C function gtk_tree_view_remove_column.
func (recv *TreeView) RemoveColumn(column *TreeViewColumn) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	retC := C.gtk_tree_view_remove_column((*C.GtkTreeView)(recv.native), c_column)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_row_activated : no return generator

// RowExpanded is a wrapper around the C function gtk_tree_view_row_expanded.
func (recv *TreeView) RowExpanded(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_row_expanded((*C.GtkTreeView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_scroll_to_cell : no return generator

// Unsupported : gtk_tree_view_scroll_to_point : no return generator

// Unsupported : gtk_tree_view_set_activate_on_single_click : no return generator

// Unsupported : gtk_tree_view_set_column_drag_function : unsupported parameter func : no type generator for TreeViewColumnDropFunc, GtkTreeViewColumnDropFunc

// Unsupported : gtk_tree_view_set_cursor : no return generator

// Unsupported : gtk_tree_view_set_cursor_on_cell : no return generator

// Unsupported : gtk_tree_view_set_destroy_count_func : unsupported parameter func : no type generator for TreeDestroyCountFunc, GtkTreeDestroyCountFunc

// Unsupported : gtk_tree_view_set_drag_dest_row : no return generator

// Unsupported : gtk_tree_view_set_enable_search : no return generator

// Unsupported : gtk_tree_view_set_enable_tree_lines : no return generator

// Unsupported : gtk_tree_view_set_expander_column : no return generator

// Unsupported : gtk_tree_view_set_fixed_height_mode : no return generator

// Unsupported : gtk_tree_view_set_grid_lines : no return generator

// Unsupported : gtk_tree_view_set_hadjustment : no return generator

// Unsupported : gtk_tree_view_set_headers_clickable : no return generator

// Unsupported : gtk_tree_view_set_headers_visible : no return generator

// Unsupported : gtk_tree_view_set_hover_expand : no return generator

// Unsupported : gtk_tree_view_set_hover_selection : no return generator

// Unsupported : gtk_tree_view_set_level_indentation : no return generator

// Unsupported : gtk_tree_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_set_reorderable : no return generator

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_tree_view_set_rubber_banding : no return generator

// Unsupported : gtk_tree_view_set_rules_hint : no return generator

// Unsupported : gtk_tree_view_set_search_column : no return generator

// Unsupported : gtk_tree_view_set_search_entry : no return generator

// Unsupported : gtk_tree_view_set_search_equal_func : unsupported parameter search_equal_func : no type generator for TreeViewSearchEqualFunc, GtkTreeViewSearchEqualFunc

// Unsupported : gtk_tree_view_set_search_position_func : unsupported parameter func : no type generator for TreeViewSearchPositionFunc, GtkTreeViewSearchPositionFunc

// Unsupported : gtk_tree_view_set_show_expanders : no return generator

// Unsupported : gtk_tree_view_set_tooltip_cell : no return generator

// Unsupported : gtk_tree_view_set_tooltip_column : no return generator

// Unsupported : gtk_tree_view_set_tooltip_row : no return generator

// Unsupported : gtk_tree_view_set_vadjustment : no return generator

// Unsupported : gtk_tree_view_unset_rows_drag_dest : no return generator

// Unsupported : gtk_tree_view_unset_rows_drag_source : no return generator

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

// Unsupported : gtk_tree_view_column_add_attribute : no return generator

// Unsupported : gtk_tree_view_column_cell_get_position : unsupported parameter x_offset : no type generator for gint, gint*

// Unsupported : gtk_tree_view_column_cell_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// CellIsVisible is a wrapper around the C function gtk_tree_view_column_cell_is_visible.
func (recv *TreeViewColumn) CellIsVisible() bool {
	retC := C.gtk_tree_view_column_cell_is_visible((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_column_cell_set_cell_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_clear : no return generator

// Unsupported : gtk_tree_view_column_clear_attributes : no return generator

// Unsupported : gtk_tree_view_column_clicked : no return generator

// Unsupported : gtk_tree_view_column_focus_cell : no return generator

// GetAlignment is a wrapper around the C function gtk_tree_view_column_get_alignment.
func (recv *TreeViewColumn) GetAlignment() float32 {
	retC := C.gtk_tree_view_column_get_alignment((*C.GtkTreeViewColumn)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetClickable is a wrapper around the C function gtk_tree_view_column_get_clickable.
func (recv *TreeViewColumn) GetClickable() bool {
	retC := C.gtk_tree_view_column_get_clickable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFixedWidth is a wrapper around the C function gtk_tree_view_column_get_fixed_width.
func (recv *TreeViewColumn) GetFixedWidth() int32 {
	retC := C.gtk_tree_view_column_get_fixed_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMaxWidth is a wrapper around the C function gtk_tree_view_column_get_max_width.
func (recv *TreeViewColumn) GetMaxWidth() int32 {
	retC := C.gtk_tree_view_column_get_max_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinWidth is a wrapper around the C function gtk_tree_view_column_get_min_width.
func (recv *TreeViewColumn) GetMinWidth() int32 {
	retC := C.gtk_tree_view_column_get_min_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReorderable is a wrapper around the C function gtk_tree_view_column_get_reorderable.
func (recv *TreeViewColumn) GetReorderable() bool {
	retC := C.gtk_tree_view_column_get_reorderable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetResizable is a wrapper around the C function gtk_tree_view_column_get_resizable.
func (recv *TreeViewColumn) GetResizable() bool {
	retC := C.gtk_tree_view_column_get_resizable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSizing is a wrapper around the C function gtk_tree_view_column_get_sizing.
func (recv *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	retC := C.gtk_tree_view_column_get_sizing((*C.GtkTreeViewColumn)(recv.native))
	retGo := (TreeViewColumnSizing)(retC)

	return retGo
}

// GetSortColumnId is a wrapper around the C function gtk_tree_view_column_get_sort_column_id.
func (recv *TreeViewColumn) GetSortColumnId() int32 {
	retC := C.gtk_tree_view_column_get_sort_column_id((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSortIndicator is a wrapper around the C function gtk_tree_view_column_get_sort_indicator.
func (recv *TreeViewColumn) GetSortIndicator() bool {
	retC := C.gtk_tree_view_column_get_sort_indicator((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSortOrder is a wrapper around the C function gtk_tree_view_column_get_sort_order.
func (recv *TreeViewColumn) GetSortOrder() SortType {
	retC := C.gtk_tree_view_column_get_sort_order((*C.GtkTreeViewColumn)(recv.native))
	retGo := (SortType)(retC)

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_tree_view_column_get_spacing.
func (recv *TreeViewColumn) GetSpacing() int32 {
	retC := C.gtk_tree_view_column_get_spacing((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTitle is a wrapper around the C function gtk_tree_view_column_get_title.
func (recv *TreeViewColumn) GetTitle() string {
	retC := C.gtk_tree_view_column_get_title((*C.GtkTreeViewColumn)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_tree_view_column_get_visible.
func (recv *TreeViewColumn) GetVisible() bool {
	retC := C.gtk_tree_view_column_get_visible((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidget is a wrapper around the C function gtk_tree_view_column_get_widget.
func (recv *TreeViewColumn) GetWidget() *Widget {
	retC := C.gtk_tree_view_column_get_widget((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gtk_tree_view_column_get_width.
func (recv *TreeViewColumn) GetWidth() int32 {
	retC := C.gtk_tree_view_column_get_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_column_pack_end : no return generator

// Unsupported : gtk_tree_view_column_pack_start : no return generator

// Unsupported : gtk_tree_view_column_queue_resize : no return generator

// Unsupported : gtk_tree_view_column_set_alignment : no return generator

// Unsupported : gtk_tree_view_column_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_set_cell_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_tree_view_column_set_clickable : no return generator

// Unsupported : gtk_tree_view_column_set_expand : no return generator

// Unsupported : gtk_tree_view_column_set_fixed_width : no return generator

// Unsupported : gtk_tree_view_column_set_max_width : no return generator

// Unsupported : gtk_tree_view_column_set_min_width : no return generator

// Unsupported : gtk_tree_view_column_set_reorderable : no return generator

// Unsupported : gtk_tree_view_column_set_resizable : no return generator

// Unsupported : gtk_tree_view_column_set_sizing : no return generator

// Unsupported : gtk_tree_view_column_set_sort_column_id : no return generator

// Unsupported : gtk_tree_view_column_set_sort_indicator : no return generator

// Unsupported : gtk_tree_view_column_set_sort_order : no return generator

// Unsupported : gtk_tree_view_column_set_spacing : no return generator

// Unsupported : gtk_tree_view_column_set_title : no return generator

// Unsupported : gtk_tree_view_column_set_visible : no return generator

// Unsupported : gtk_tree_view_column_set_widget : no return generator

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

	return g
}

func (recv *UIManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_ui_manager_add_ui : no return generator

// Unsupported : gtk_ui_manager_ensure_update : no return generator

// Unsupported : gtk_ui_manager_insert_action_group : no return generator

// Unsupported : gtk_ui_manager_remove_action_group : no return generator

// Unsupported : gtk_ui_manager_remove_ui : no return generator

// Unsupported : gtk_ui_manager_set_add_tearoffs : no return generator

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

	return g
}

func (recv *VBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VBoxNew is a wrapper around the C function gtk_vbox_new.
func VBoxNew(homogeneous bool, spacing int32) *Widget {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_vbox_new(c_homogeneous, c_spacing)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *VButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VButtonBoxNew is a wrapper around the C function gtk_vbutton_box_new.
func VButtonBoxNew() *Widget {
	retC := C.gtk_vbutton_box_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *VPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VPanedNew is a wrapper around the C function gtk_vpaned_new.
func VPanedNew() *Widget {
	retC := C.gtk_vpaned_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *VScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScaleNew is a wrapper around the C function gtk_vscale_new.
func VScaleNew(adjustment *Adjustment) *Widget {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_vscale_new(c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScaleNewWithRange is a wrapper around the C function gtk_vscale_new_with_range.
func VScaleNewWithRange(min float64, max float64, step float64) *Widget {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_vscale_new_with_range(c_min, c_max, c_step)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *VScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VScrollbarNew is a wrapper around the C function gtk_vscrollbar_new.
func VScrollbarNew(adjustment *Adjustment) *Widget {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_vscrollbar_new(c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *VSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VSeparatorNew is a wrapper around the C function gtk_vseparator_new.
func VSeparatorNew() *Widget {
	retC := C.gtk_vseparator_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *Viewport) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ViewportNew is a wrapper around the C function gtk_viewport_new.
func ViewportNew(hadjustment *Adjustment, vadjustment *Adjustment) *Widget {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_viewport_new(c_hadjustment, c_vadjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_viewport_get_hadjustment.
func (recv *Viewport) GetHadjustment() *Adjustment {
	retC := C.gtk_viewport_get_hadjustment((*C.GtkViewport)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_viewport_get_shadow_type.
func (recv *Viewport) GetShadowType() ShadowType {
	retC := C.gtk_viewport_get_shadow_type((*C.GtkViewport)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_viewport_get_vadjustment.
func (recv *Viewport) GetVadjustment() *Adjustment {
	retC := C.gtk_viewport_get_vadjustment((*C.GtkViewport)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_viewport_set_hadjustment : no return generator

// Unsupported : gtk_viewport_set_shadow_type : no return generator

// Unsupported : gtk_viewport_set_vadjustment : no return generator

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

	return g
}

func (recv *Widget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Activate is a wrapper around the C function gtk_widget_activate.
func (recv *Widget) Activate() bool {
	retC := C.gtk_widget_activate((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_add_accelerator : no return generator

// Unsupported : gtk_widget_add_device_events : no return generator

// Unsupported : gtk_widget_add_events : no return generator

// Unsupported : gtk_widget_add_mnemonic_label : no return generator

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback, GtkTickCallback

// ChildFocus is a wrapper around the C function gtk_widget_child_focus.
func (recv *Widget) ChildFocus(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_widget_child_focus((*C.GtkWidget)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_child_notify : no return generator

// Unsupported : gtk_widget_class_path : unsupported parameter path_length : no type generator for guint, guint*

// ComputeExpand is a wrapper around the C function gtk_widget_compute_expand.
func (recv *Widget) ComputeExpand(orientation Orientation) bool {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_widget_compute_expand((*C.GtkWidget)(recv.native), c_orientation)
	retGo := retC == C.TRUE

	return retGo
}

// CreatePangoContext is a wrapper around the C function gtk_widget_create_pango_context.
func (recv *Widget) CreatePangoContext() *pango.Context {
	retC := C.gtk_widget_create_pango_context((*C.GtkWidget)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreatePangoLayout is a wrapper around the C function gtk_widget_create_pango_layout.
func (recv *Widget) CreatePangoLayout(text string) *pango.Layout {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_widget_create_pango_layout((*C.GtkWidget)(recv.native), c_text)
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_destroy : no return generator

// Unsupported : gtk_widget_destroyed : unsupported parameter widget_pointer : record with indirection level of 2

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// DragCheckThreshold is a wrapper around the C function gtk_drag_check_threshold.
func (recv *Widget) DragCheckThreshold(startX int32, startY int32, currentX int32, currentY int32) bool {
	c_start_x := (C.gint)(startX)

	c_start_y := (C.gint)(startY)

	c_current_x := (C.gint)(currentX)

	c_current_y := (C.gint)(currentY)

	retC := C.gtk_drag_check_threshold((*C.GtkWidget)(recv.native), c_start_x, c_start_y, c_current_x, c_current_y)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_drag_dest_add_image_targets : no return generator

// Unsupported : gtk_drag_dest_add_text_targets : no return generator

// Unsupported : gtk_drag_dest_add_uri_targets : no return generator

// Unsupported : gtk_drag_dest_find_target : return type : Blacklisted record : GdkAtom

// DragDestGetTargetList is a wrapper around the C function gtk_drag_dest_get_target_list.
func (recv *Widget) DragDestGetTargetList() *TargetList {
	retC := C.gtk_drag_dest_get_target_list((*C.GtkWidget)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_drag_dest_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_dest_set_proxy : no return generator

// Unsupported : gtk_drag_dest_set_target_list : no return generator

// Unsupported : gtk_drag_dest_set_track_motion : no return generator

// Unsupported : gtk_drag_dest_unset : no return generator

// Unsupported : gtk_drag_get_data : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_highlight : no return generator

// Unsupported : gtk_drag_source_add_image_targets : no return generator

// Unsupported : gtk_drag_source_add_text_targets : no return generator

// Unsupported : gtk_drag_source_add_uri_targets : no return generator

// Unsupported : gtk_drag_source_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_drag_source_set_icon_name : no return generator

// Unsupported : gtk_drag_source_set_icon_pixbuf : no return generator

// Unsupported : gtk_drag_source_set_icon_stock : no return generator

// Unsupported : gtk_drag_source_set_target_list : no return generator

// Unsupported : gtk_drag_source_unset : no return generator

// Unsupported : gtk_drag_unhighlight : no return generator

// Unsupported : gtk_widget_draw : no return generator

// Unsupported : gtk_widget_ensure_style : no return generator

// Unsupported : gtk_widget_error_bell : no return generator

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_freeze_child_notify : no return generator

// GetAccessible is a wrapper around the C function gtk_widget_get_accessible.
func (recv *Widget) GetAccessible() *atk.Object {
	retC := C.gtk_widget_get_accessible((*C.GtkWidget)(recv.native))
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_get_action_group : no return generator

// GetAllocatedHeight is a wrapper around the C function gtk_widget_get_allocated_height.
func (recv *Widget) GetAllocatedHeight() int32 {
	retC := C.gtk_widget_get_allocated_height((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetAllocatedWidth is a wrapper around the C function gtk_widget_get_allocated_width.
func (recv *Widget) GetAllocatedWidth() int32 {
	retC := C.gtk_widget_get_allocated_width((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_ancestor : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_widget_get_child_requisition : no return generator

// GetChildVisible is a wrapper around the C function gtk_widget_get_child_visible.
func (recv *Widget) GetChildVisible() bool {
	retC := C.gtk_widget_get_child_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// GetCompositeName is a wrapper around the C function gtk_widget_get_composite_name.
func (recv *Widget) GetCompositeName() string {
	retC := C.gtk_widget_get_composite_name((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDirection is a wrapper around the C function gtk_widget_get_direction.
func (recv *Widget) GetDirection() TextDirection {
	retC := C.gtk_widget_get_direction((*C.GtkWidget)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetEvents is a wrapper around the C function gtk_widget_get_events.
func (recv *Widget) GetEvents() int32 {
	retC := C.gtk_widget_get_events((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHalign is a wrapper around the C function gtk_widget_get_halign.
func (recv *Widget) GetHalign() Align {
	retC := C.gtk_widget_get_halign((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// GetHexpand is a wrapper around the C function gtk_widget_get_hexpand.
func (recv *Widget) GetHexpand() bool {
	retC := C.gtk_widget_get_hexpand((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHexpandSet is a wrapper around the C function gtk_widget_get_hexpand_set.
func (recv *Widget) GetHexpandSet() bool {
	retC := C.gtk_widget_get_hexpand_set((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModifierStyle is a wrapper around the C function gtk_widget_get_modifier_style.
func (recv *Widget) GetModifierStyle() *RcStyle {
	retC := C.gtk_widget_get_modifier_style((*C.GtkWidget)(recv.native))
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_widget_get_name.
func (recv *Widget) GetName() string {
	retC := C.gtk_widget_get_name((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPangoContext is a wrapper around the C function gtk_widget_get_pango_context.
func (recv *Widget) GetPangoContext() *pango.Context {
	retC := C.gtk_widget_get_pango_context((*C.GtkWidget)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParent is a wrapper around the C function gtk_widget_get_parent.
func (recv *Widget) GetParent() *Widget {
	retC := C.gtk_widget_get_parent((*C.GtkWidget)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParentWindow is a wrapper around the C function gtk_widget_get_parent_window.
func (recv *Widget) GetParentWindow() *gdk.Window {
	retC := C.gtk_widget_get_parent_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPath is a wrapper around the C function gtk_widget_get_path.
func (recv *Widget) GetPath() *WidgetPath {
	retC := C.gtk_widget_get_path((*C.GtkWidget)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_and_baseline_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_size : no return generator

// Unsupported : gtk_widget_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_requisition : no return generator

// GetSettings is a wrapper around the C function gtk_widget_get_settings.
func (recv *Widget) GetSettings() *Settings {
	retC := C.gtk_widget_get_settings((*C.GtkWidget)(recv.native))
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_get_size_request : unsupported parameter width : no type generator for gint, gint*

// GetStyle is a wrapper around the C function gtk_widget_get_style.
func (recv *Widget) GetStyle() *Style {
	retC := C.gtk_widget_get_style((*C.GtkWidget)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStyleContext is a wrapper around the C function gtk_widget_get_style_context.
func (recv *Widget) GetStyleContext() *StyleContext {
	retC := C.gtk_widget_get_style_context((*C.GtkWidget)(recv.native))
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSupportMultidevice is a wrapper around the C function gtk_widget_get_support_multidevice.
func (recv *Widget) GetSupportMultidevice() bool {
	retC := C.gtk_widget_get_support_multidevice((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

// GetToplevel is a wrapper around the C function gtk_widget_get_toplevel.
func (recv *Widget) GetToplevel() *Widget {
	retC := C.gtk_widget_get_toplevel((*C.GtkWidget)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValign is a wrapper around the C function gtk_widget_get_valign.
func (recv *Widget) GetValign() Align {
	retC := C.gtk_widget_get_valign((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// GetVexpand is a wrapper around the C function gtk_widget_get_vexpand.
func (recv *Widget) GetVexpand() bool {
	retC := C.gtk_widget_get_vexpand((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVexpandSet is a wrapper around the C function gtk_widget_get_vexpand_set.
func (recv *Widget) GetVexpandSet() bool {
	retC := C.gtk_widget_get_vexpand_set((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisual is a wrapper around the C function gtk_widget_get_visual.
func (recv *Widget) GetVisual() *gdk.Visual {
	retC := C.gtk_widget_get_visual((*C.GtkWidget)(recv.native))
	retGo := gdk.VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_grab_add : no return generator

// Unsupported : gtk_widget_grab_default : no return generator

// Unsupported : gtk_widget_grab_focus : no return generator

// Unsupported : gtk_grab_remove : no return generator

// Unsupported : gtk_widget_hide : no return generator

// HideOnDelete is a wrapper around the C function gtk_widget_hide_on_delete.
func (recv *Widget) HideOnDelete() bool {
	retC := C.gtk_widget_hide_on_delete((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// InDestruction is a wrapper around the C function gtk_widget_in_destruction.
func (recv *Widget) InDestruction() bool {
	retC := C.gtk_widget_in_destruction((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_init_template : no return generator

// Unsupported : gtk_widget_input_shape_combine_region : no return generator

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_widget_intersect : unsupported parameter area : Blacklisted record : GdkRectangle

// IsAncestor is a wrapper around the C function gtk_widget_is_ancestor.
func (recv *Widget) IsAncestor(ancestor *Widget) bool {
	c_ancestor := (*C.GtkWidget)(ancestor.ToC())

	retC := C.gtk_widget_is_ancestor((*C.GtkWidget)(recv.native), c_ancestor)
	retGo := retC == C.TRUE

	return retGo
}

// IsFocus is a wrapper around the C function gtk_widget_is_focus.
func (recv *Widget) IsFocus() bool {
	retC := C.gtk_widget_is_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListAccelClosures is a wrapper around the C function gtk_widget_list_accel_closures.
func (recv *Widget) ListAccelClosures() *glib.List {
	retC := C.gtk_widget_list_accel_closures((*C.GtkWidget)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_list_action_prefixes : no return type

// Unsupported : gtk_widget_map : no return generator

// MnemonicActivate is a wrapper around the C function gtk_widget_mnemonic_activate.
func (recv *Widget) MnemonicActivate(groupCycling bool) bool {
	c_group_cycling :=
		boolToGboolean(groupCycling)

	retC := C.gtk_widget_mnemonic_activate((*C.GtkWidget)(recv.native), c_group_cycling)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_modify_base : no return generator

// Unsupported : gtk_widget_modify_bg : no return generator

// Unsupported : gtk_widget_modify_cursor : no return generator

// Unsupported : gtk_widget_modify_fg : no return generator

// Unsupported : gtk_widget_modify_font : no return generator

// Unsupported : gtk_widget_modify_style : no return generator

// Unsupported : gtk_widget_modify_text : no return generator

// Unsupported : gtk_widget_override_background_color : no return generator

// Unsupported : gtk_widget_override_color : no return generator

// Unsupported : gtk_widget_override_cursor : no return generator

// Unsupported : gtk_widget_override_font : no return generator

// Unsupported : gtk_widget_override_symbolic_color : no return generator

// Unsupported : gtk_widget_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_queue_allocate : no return generator

// Unsupported : gtk_widget_queue_compute_expand : no return generator

// Unsupported : gtk_widget_queue_draw : no return generator

// Unsupported : gtk_widget_queue_draw_area : no return generator

// Unsupported : gtk_widget_queue_draw_region : no return generator

// Unsupported : gtk_widget_queue_resize : no return generator

// Unsupported : gtk_widget_queue_resize_no_redraw : no return generator

// Unsupported : gtk_widget_realize : no return generator

// RegionIntersect is a wrapper around the C function gtk_widget_region_intersect.
func (recv *Widget) RegionIntersect(region *cairo.Region) *cairo.Region {
	c_region := (*C.cairo_region_t)(region.ToC())

	retC := C.gtk_widget_region_intersect((*C.GtkWidget)(recv.native), c_region)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_register_window : no return generator

// RemoveAccelerator is a wrapper around the C function gtk_widget_remove_accelerator.
func (recv *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_widget_remove_accelerator((*C.GtkWidget)(recv.native), c_accel_group, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_remove_mnemonic_label : no return generator

// Unsupported : gtk_widget_remove_tick_callback : no return generator

// Unsupported : gtk_widget_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_reparent : no return generator

// Unsupported : gtk_widget_reset_rc_styles : no return generator

// Unsupported : gtk_widget_reset_style : no return generator

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_set_accel_path : no return generator

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_app_paintable : no return generator

// Unsupported : gtk_widget_set_can_default : no return generator

// Unsupported : gtk_widget_set_can_focus : no return generator

// Unsupported : gtk_widget_set_child_visible : no return generator

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_composite_name : no return generator

// Unsupported : gtk_widget_set_device_enabled : no return generator

// Unsupported : gtk_widget_set_device_events : no return generator

// Unsupported : gtk_widget_set_direction : no return generator

// Unsupported : gtk_widget_set_double_buffered : no return generator

// Unsupported : gtk_widget_set_events : no return generator

// Unsupported : gtk_widget_set_focus_on_click : no return generator

// Unsupported : gtk_widget_set_font_map : no return generator

// Unsupported : gtk_widget_set_font_options : no return generator

// Unsupported : gtk_widget_set_halign : no return generator

// Unsupported : gtk_widget_set_has_tooltip : no return generator

// Unsupported : gtk_widget_set_has_window : no return generator

// Unsupported : gtk_widget_set_hexpand : no return generator

// Unsupported : gtk_widget_set_hexpand_set : no return generator

// Unsupported : gtk_widget_set_mapped : no return generator

// Unsupported : gtk_widget_set_margin_bottom : no return generator

// Unsupported : gtk_widget_set_margin_end : no return generator

// Unsupported : gtk_widget_set_margin_left : no return generator

// Unsupported : gtk_widget_set_margin_right : no return generator

// Unsupported : gtk_widget_set_margin_start : no return generator

// Unsupported : gtk_widget_set_margin_top : no return generator

// Unsupported : gtk_widget_set_name : no return generator

// Unsupported : gtk_widget_set_no_show_all : no return generator

// Unsupported : gtk_widget_set_opacity : no return generator

// Unsupported : gtk_widget_set_parent : no return generator

// Unsupported : gtk_widget_set_parent_window : no return generator

// Unsupported : gtk_widget_set_realized : no return generator

// Unsupported : gtk_widget_set_receives_default : no return generator

// Unsupported : gtk_widget_set_redraw_on_allocate : no return generator

// Unsupported : gtk_widget_set_sensitive : no return generator

// Unsupported : gtk_widget_set_size_request : no return generator

// Unsupported : gtk_widget_set_state : no return generator

// Unsupported : gtk_widget_set_state_flags : no return generator

// Unsupported : gtk_widget_set_style : no return generator

// Unsupported : gtk_widget_set_support_multidevice : no return generator

// Unsupported : gtk_widget_set_tooltip_markup : no return generator

// Unsupported : gtk_widget_set_tooltip_text : no return generator

// Unsupported : gtk_widget_set_tooltip_window : no return generator

// Unsupported : gtk_widget_set_valign : no return generator

// Unsupported : gtk_widget_set_vexpand : no return generator

// Unsupported : gtk_widget_set_vexpand_set : no return generator

// Unsupported : gtk_widget_set_visible : no return generator

// Unsupported : gtk_widget_set_visual : no return generator

// Unsupported : gtk_widget_set_window : no return generator

// Unsupported : gtk_widget_shape_combine_region : no return generator

// Unsupported : gtk_widget_show : no return generator

// Unsupported : gtk_widget_show_all : no return generator

// Unsupported : gtk_widget_show_now : no return generator

// Unsupported : gtk_widget_size_allocate : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_request : no return generator

// Unsupported : gtk_widget_style_attach : no return generator

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// Unsupported : gtk_widget_style_get_property : no return generator

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_widget_thaw_child_notify : no return generator

// Unsupported : gtk_widget_translate_coordinates : unsupported parameter dest_x : no type generator for gint, gint*

// Unsupported : gtk_widget_trigger_tooltip_query : no return generator

// Unsupported : gtk_widget_unmap : no return generator

// Unsupported : gtk_widget_unparent : no return generator

// Unsupported : gtk_widget_unrealize : no return generator

// Unsupported : gtk_widget_unregister_window : no return generator

// Unsupported : gtk_widget_unset_state_flags : no return generator

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

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowNew is a wrapper around the C function gtk_window_new.
func WindowNew(type_ WindowType) *Widget {
	c_type := (C.GtkWindowType)(type_)

	retC := C.gtk_window_new(c_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ActivateDefault is a wrapper around the C function gtk_window_activate_default.
func (recv *Window) ActivateDefault() bool {
	retC := C.gtk_window_activate_default((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ActivateFocus is a wrapper around the C function gtk_window_activate_focus.
func (recv *Window) ActivateFocus() bool {
	retC := C.gtk_window_activate_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_add_accel_group : no return generator

// Unsupported : gtk_window_add_mnemonic : no return generator

// Unsupported : gtk_window_begin_move_drag : no return generator

// Unsupported : gtk_window_begin_resize_drag : no return generator

// Unsupported : gtk_window_close : no return generator

// Unsupported : gtk_window_deiconify : no return generator

// Unsupported : gtk_window_fullscreen : no return generator

// Unsupported : gtk_window_fullscreen_on_monitor : no return generator

// GetDecorated is a wrapper around the C function gtk_window_get_decorated.
func (recv *Window) GetDecorated() bool {
	retC := C.gtk_window_get_decorated((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// GetDestroyWithParent is a wrapper around the C function gtk_window_get_destroy_with_parent.
func (recv *Window) GetDestroyWithParent() bool {
	retC := C.gtk_window_get_destroy_with_parent((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFocus is a wrapper around the C function gtk_window_get_focus.
func (recv *Window) GetFocus() *Widget {
	retC := C.gtk_window_get_focus((*C.GtkWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGravity is a wrapper around the C function gtk_window_get_gravity.
func (recv *Window) GetGravity() gdk.Gravity {
	retC := C.gtk_window_get_gravity((*C.GtkWindow)(recv.native))
	retGo := (gdk.Gravity)(retC)

	return retGo
}

// GetIcon is a wrapper around the C function gtk_window_get_icon.
func (recv *Window) GetIcon() *gdkpixbuf.Pixbuf {
	retC := C.gtk_window_get_icon((*C.GtkWindow)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconList is a wrapper around the C function gtk_window_get_icon_list.
func (recv *Window) GetIconList() *glib.List {
	retC := C.gtk_window_get_icon_list((*C.GtkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMnemonicModifier is a wrapper around the C function gtk_window_get_mnemonic_modifier.
func (recv *Window) GetMnemonicModifier() gdk.ModifierType {
	retC := C.gtk_window_get_mnemonic_modifier((*C.GtkWindow)(recv.native))
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// GetModal is a wrapper around the C function gtk_window_get_modal.
func (recv *Window) GetModal() bool {
	retC := C.gtk_window_get_modal((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// GetResizable is a wrapper around the C function gtk_window_get_resizable.
func (recv *Window) GetResizable() bool {
	retC := C.gtk_window_get_resizable((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetRole is a wrapper around the C function gtk_window_get_role.
func (recv *Window) GetRole() string {
	retC := C.gtk_window_get_role((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// GetTitle is a wrapper around the C function gtk_window_get_title.
func (recv *Window) GetTitle() string {
	retC := C.gtk_window_get_title((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTransientFor is a wrapper around the C function gtk_window_get_transient_for.
func (recv *Window) GetTransientFor() *Window {
	retC := C.gtk_window_get_transient_for((*C.GtkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTypeHint is a wrapper around the C function gtk_window_get_type_hint.
func (recv *Window) GetTypeHint() gdk.WindowTypeHint {
	retC := C.gtk_window_get_type_hint((*C.GtkWindow)(recv.native))
	retGo := (gdk.WindowTypeHint)(retC)

	return retGo
}

// HasGroup is a wrapper around the C function gtk_window_has_group.
func (recv *Window) HasGroup() bool {
	retC := C.gtk_window_has_group((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_iconify : no return generator

// Unsupported : gtk_window_maximize : no return generator

// MnemonicActivate is a wrapper around the C function gtk_window_mnemonic_activate.
func (recv *Window) MnemonicActivate(keyval uint32, modifier gdk.ModifierType) bool {
	c_keyval := (C.guint)(keyval)

	c_modifier := (C.GdkModifierType)(modifier)

	retC := C.gtk_window_mnemonic_activate((*C.GtkWindow)(recv.native), c_keyval, c_modifier)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_move : no return generator

// ParseGeometry is a wrapper around the C function gtk_window_parse_geometry.
func (recv *Window) ParseGeometry(geometry string) bool {
	c_geometry := C.CString(geometry)
	defer C.free(unsafe.Pointer(c_geometry))

	retC := C.gtk_window_parse_geometry((*C.GtkWindow)(recv.native), c_geometry)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_present : no return generator

// Unsupported : gtk_window_present_with_time : no return generator

// Unsupported : gtk_window_remove_accel_group : no return generator

// Unsupported : gtk_window_remove_mnemonic : no return generator

// Unsupported : gtk_window_reshow_with_initial_size : no return generator

// Unsupported : gtk_window_resize : no return generator

// Unsupported : gtk_window_resize_to_geometry : no return generator

// Unsupported : gtk_window_set_accept_focus : no return generator

// Unsupported : gtk_window_set_application : no return generator

// Unsupported : gtk_window_set_attached_to : no return generator

// Unsupported : gtk_window_set_decorated : no return generator

// Unsupported : gtk_window_set_default : no return generator

// Unsupported : gtk_window_set_default_geometry : no return generator

// Unsupported : gtk_window_set_default_size : no return generator

// Unsupported : gtk_window_set_deletable : no return generator

// Unsupported : gtk_window_set_destroy_with_parent : no return generator

// Unsupported : gtk_window_set_focus : no return generator

// Unsupported : gtk_window_set_focus_on_map : no return generator

// Unsupported : gtk_window_set_focus_visible : no return generator

// Unsupported : gtk_window_set_geometry_hints : no return generator

// Unsupported : gtk_window_set_gravity : no return generator

// Unsupported : gtk_window_set_has_resize_grip : no return generator

// Unsupported : gtk_window_set_has_user_ref_count : no return generator

// Unsupported : gtk_window_set_hide_titlebar_when_maximized : no return generator

// Unsupported : gtk_window_set_icon : no return generator

// Unsupported : gtk_window_set_icon_list : no return generator

// Unsupported : gtk_window_set_icon_name : no return generator

// Unsupported : gtk_window_set_keep_above : no return generator

// Unsupported : gtk_window_set_keep_below : no return generator

// Unsupported : gtk_window_set_mnemonic_modifier : no return generator

// Unsupported : gtk_window_set_mnemonics_visible : no return generator

// Unsupported : gtk_window_set_modal : no return generator

// Unsupported : gtk_window_set_opacity : no return generator

// Unsupported : gtk_window_set_position : no return generator

// Unsupported : gtk_window_set_resizable : no return generator

// Unsupported : gtk_window_set_role : no return generator

// Unsupported : gtk_window_set_screen : no return generator

// Unsupported : gtk_window_set_skip_pager_hint : no return generator

// Unsupported : gtk_window_set_skip_taskbar_hint : no return generator

// Unsupported : gtk_window_set_startup_id : no return generator

// Unsupported : gtk_window_set_title : no return generator

// Unsupported : gtk_window_set_titlebar : no return generator

// Unsupported : gtk_window_set_transient_for : no return generator

// Unsupported : gtk_window_set_type_hint : no return generator

// Unsupported : gtk_window_set_urgency_hint : no return generator

// Unsupported : gtk_window_set_wmclass : no return generator

// Unsupported : gtk_window_stick : no return generator

// Unsupported : gtk_window_unfullscreen : no return generator

// Unsupported : gtk_window_unmaximize : no return generator

// Unsupported : gtk_window_unstick : no return generator

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

	return g
}

func (recv *WindowGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowGroupNew is a wrapper around the C function gtk_window_group_new.
func WindowGroupNew() *WindowGroup {
	retC := C.gtk_window_group_new()
	retGo := WindowGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_window_group_add_window : no return generator

// Unsupported : gtk_window_group_remove_window : no return generator
