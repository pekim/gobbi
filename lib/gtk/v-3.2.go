// Code generated - DO NOT EDIT.
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	"unsafe"
)

// CssSectionType is a representation of the C enumeration GtkCssSectionType.
type CssSectionType int

// CssSectionType_document is a representation of the C enumeration member GTK_CSS_SECTION_DOCUMENT.
const CssSectionType_document = CssSectionType(0)

// CssSectionType_import is a representation of the C enumeration member GTK_CSS_SECTION_IMPORT.
const CssSectionType_import = CssSectionType(1)

// CssSectionType_color_definition is a representation of the C enumeration member GTK_CSS_SECTION_COLOR_DEFINITION.
const CssSectionType_color_definition = CssSectionType(2)

// CssSectionType_binding_set is a representation of the C enumeration member GTK_CSS_SECTION_BINDING_SET.
const CssSectionType_binding_set = CssSectionType(3)

// CssSectionType_ruleset is a representation of the C enumeration member GTK_CSS_SECTION_RULESET.
const CssSectionType_ruleset = CssSectionType(4)

// CssSectionType_selector is a representation of the C enumeration member GTK_CSS_SECTION_SELECTOR.
const CssSectionType_selector = CssSectionType(5)

// CssSectionType_declaration is a representation of the C enumeration member GTK_CSS_SECTION_DECLARATION.
const CssSectionType_declaration = CssSectionType(6)

// CssSectionType_value is a representation of the C enumeration member GTK_CSS_SECTION_VALUE.
const CssSectionType_value = CssSectionType(7)

// CssSectionType_keyframes is a representation of the C enumeration member GTK_CSS_SECTION_KEYFRAMES.
const CssSectionType_keyframes = CssSectionType(8)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// DragSetIconGicon wraps the C function gtk_drag_set_icon_gicon.
//
// since 3.2
func DragSetIconGicon(context *gdk.DragContext, icon *gio.Icon, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_icon := icon.ToC()
	sys_hotX := hotX
	sys_hotY := hotY
	gtk.Fn_gtk_drag_set_icon_gicon(sys_context, sys_icon, sys_hotX, sys_hotY)
}

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// RenderIcon wraps the C function gtk_render_icon.
//
// since 3.2
func RenderIcon(context *StyleContext, cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x float64, y float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_pixbuf := pixbuf.ToC()
	sys_x := x
	sys_y := y
	gtk.Fn_gtk_render_icon(sys_context, sys_cr, sys_pixbuf, sys_x, sys_y)
}

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// CssSection is a representation of the C record GtkCssSection.
//
// since 3.2
type CssSection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkCssSection that represents the CssSection.
func (recv *CssSection) ToC() unsafe.Pointer {
	return recv.native
}

// CssSectionNewFromC creates a new CssSection from a pointer to the C GtkCssSection that represents the CssSection.
func CssSectionNewFromC(native unsafe.Pointer) *CssSection {
	return &CssSection{native: native}
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// FontChooserDialog is a representation of the C record GtkFontChooserDialog.
//
// since 3.2
type FontChooserDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserDialog that represents the FontChooserDialog.
func (recv *FontChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserDialogNewFromC creates a new FontChooserDialog from a pointer to the C GtkFontChooserDialog that represents the FontChooserDialog.
func FontChooserDialogNewFromC(native unsafe.Pointer) *FontChooserDialog {
	return &FontChooserDialog{native: native}
}

// FontChooserWidget is a representation of the C record GtkFontChooserWidget.
//
// since 3.2
type FontChooserWidget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkFontChooserWidget that represents the FontChooserWidget.
func (recv *FontChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// FontChooserWidgetNewFromC creates a new FontChooserWidget from a pointer to the C GtkFontChooserWidget that represents the FontChooserWidget.
func FontChooserWidgetNewFromC(native unsafe.Pointer) *FontChooserWidget {
	return &FontChooserWidget{native: native}
}

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted
