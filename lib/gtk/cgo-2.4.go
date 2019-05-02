// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
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

	static GtkMessageDialog* _gtk_message_dialog_new_with_markup(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, const gchar* message_format) {
		return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, message_format);
    }
*/
import "C"

// ActionNew is a wrapper around the C function gtk_action_new.
func ActionNew(name string, label string, tooltip string, stockId string) *Action {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ActionGroupNew is a wrapper around the C function gtk_action_group_new.
func ActionGroupNew(name string) *ActionGroup {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_action_group_new(c_name)
	retGo := ActionGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ColorButtonNew is a wrapper around the C function gtk_color_button_new.
func ColorButtonNew() *ColorButton {
	retC := C.gtk_color_button_new()
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorButtonNewWithColor is a wrapper around the C function gtk_color_button_new_with_color.
func ColorButtonNewWithColor(color *gdk.Color) *ColorButton {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	retC := C.gtk_color_button_new_with_color(c_color)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNew is a wrapper around the C function gtk_combo_box_new.
func ComboBoxNew() *ComboBox {
	retC := C.gtk_combo_box_new()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithModel is a wrapper around the C function gtk_combo_box_new_with_model.
func ComboBoxNewWithModel(model *TreeModel) *ComboBox {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_combo_box_new_with_model(c_model)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EntryCompletionNew is a wrapper around the C function gtk_entry_completion_new.
func EntryCompletionNew() *EntryCompletion {
	retC := C.gtk_entry_completion_new()
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ExpanderNew is a wrapper around the C function gtk_expander_new.
func ExpanderNew(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExpanderNewWithMnemonic is a wrapper around the C function gtk_expander_new_with_mnemonic.
func ExpanderNewWithMnemonic(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new_with_mnemonic(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// FileChooserWidgetNew is a wrapper around the C function gtk_file_chooser_widget_new.
func FileChooserWidgetNew(action FileChooserAction) *FileChooserWidget {
	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_widget_new(c_action)
	retGo := FileChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileFilterNew is a wrapper around the C function gtk_file_filter_new.
func FileFilterNew() *FileFilter {
	retC := C.gtk_file_filter_new()
	retGo := FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontButtonNew is a wrapper around the C function gtk_font_button_new.
func FontButtonNew() *FontButton {
	retC := C.gtk_font_button_new()
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontButtonNewWithFont is a wrapper around the C function gtk_font_button_new_with_font.
func FontButtonNewWithFont(fontname string) *FontButton {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_new_with_font(c_fontname)
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconThemeNew is a wrapper around the C function gtk_icon_theme_new.
func IconThemeNew() *IconTheme {
	retC := C.gtk_icon_theme_new()
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// MessageDialogNewWithMarkup is a wrapper around the C function gtk_message_dialog_new_with_markup.
func MessageDialogNewWithMarkup(parent *Window, flags DialogFlags, type_ MessageType, buttons ButtonsType, messageFormat string, args ...interface{}) *MessageDialog {
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

	retC := C._gtk_message_dialog_new_with_markup(c_parent, c_flags, c_type, c_buttons, c_message_format)
	retGo := MessageDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioActionNew is a wrapper around the C function gtk_radio_action_new.
func RadioActionNew(name string, label string, tooltip string, stockId string, value int32) *RadioAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_value := (C.gint)(value)

	retC := C.gtk_radio_action_new(c_name, c_label, c_tooltip, c_stock_id, c_value)
	retGo := RadioActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// RadioMenuItemNewFromWidget is a wrapper around the C function gtk_radio_menu_item_new_from_widget.
func RadioMenuItemNewFromWidget(group *RadioMenuItem) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	retC := C.gtk_radio_menu_item_new_from_widget(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_label_from_widget.
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic_from_widget.
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNew is a wrapper around the C function gtk_radio_tool_button_new.
func RadioToolButtonNew(group *glib.SList) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromStock is a wrapper around the C function gtk_radio_tool_button_new_from_stock.
func RadioToolButtonNewFromStock(group *glib.SList, stockId string) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_from_stock(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromWidget is a wrapper around the C function gtk_radio_tool_button_new_from_widget.
func RadioToolButtonNewFromWidget(group *RadioToolButton) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new_from_widget(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewWithStockFromWidget is a wrapper around the C function gtk_radio_tool_button_new_with_stock_from_widget.
func RadioToolButtonNewWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_with_stock_from_widget(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SeparatorToolItemNew is a wrapper around the C function gtk_separator_tool_item_new.
func SeparatorToolItemNew() *SeparatorToolItem {
	retC := C.gtk_separator_tool_item_new()
	retGo := SeparatorToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleActionNew is a wrapper around the C function gtk_toggle_action_new.
func ToggleActionNew(name string, label string, tooltip string, stockId string) *ToggleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ToggleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ToggleToolButtonNew is a wrapper around the C function gtk_toggle_tool_button_new.
func ToggleToolButtonNew() *ToggleToolButton {
	retC := C.gtk_toggle_tool_button_new()
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleToolButtonNewFromStock is a wrapper around the C function gtk_toggle_tool_button_new_from_stock.
func ToggleToolButtonNewFromStock(stockId string) *ToggleToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_tool_button_new_from_stock(c_stock_id)
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolButtonNew is a wrapper around the C function gtk_tool_button_new.
func ToolButtonNew(iconWidget *Widget, label string) *ToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_button_new(c_icon_widget, c_label)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolButtonNewFromStock is a wrapper around the C function gtk_tool_button_new_from_stock.
func ToolButtonNewFromStock(stockId string) *ToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_tool_button_new_from_stock(c_stock_id)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolItemNew is a wrapper around the C function gtk_tool_item_new.
func ToolItemNew() *ToolItem {
	retC := C.gtk_tool_item_new()
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UIManagerNew is a wrapper around the C function gtk_ui_manager_new.
func UIManagerNew() *UIManager {
	retC := C.gtk_ui_manager_new()
	retGo := UIManagerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : STOCK_DIALOG_AUTHENTICATION

// Blacklisted : STOCK_HARDDISK

// Blacklisted : STOCK_INDENT

// Blacklisted : STOCK_NETWORK

// Blacklisted : STOCK_UNINDENT
