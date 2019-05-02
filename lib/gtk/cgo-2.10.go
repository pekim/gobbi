// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
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
import "C"

// AssistantNew is a wrapper around the C function gtk_assistant_new.
func AssistantNew() *Assistant {
	retC := C.gtk_assistant_new()
	retGo := AssistantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererAccelNew is a wrapper around the C function gtk_cell_renderer_accel_new.
func CellRendererAccelNew() *CellRendererAccel {
	retC := C.gtk_cell_renderer_accel_new()
	retGo := CellRendererAccelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererSpinNew is a wrapper around the C function gtk_cell_renderer_spin_new.
func CellRendererSpinNew() *CellRendererSpin {
	retC := C.gtk_cell_renderer_spin_new()
	retGo := CellRendererSpinNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LinkButtonNew is a wrapper around the C function gtk_link_button_new.
func LinkButtonNew(uri string) *LinkButton {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_link_button_new(c_uri)
	retGo := LinkButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LinkButtonNewWithLabel is a wrapper around the C function gtk_link_button_new_with_label.
func LinkButtonNewWithLabel(uri string, label string) *LinkButton {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_link_button_new_with_label(c_uri, c_label)
	retGo := LinkButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PageSetupNew is a wrapper around the C function gtk_page_setup_new.
func PageSetupNew() *PageSetup {
	retC := C.gtk_page_setup_new()
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PrintOperationNew is a wrapper around the C function gtk_print_operation_new.
func PrintOperationNew() *PrintOperation {
	retC := C.gtk_print_operation_new()
	retGo := PrintOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PrintSettingsNew is a wrapper around the C function gtk_print_settings_new.
func PrintSettingsNew() *PrintSettings {
	retC := C.gtk_print_settings_new()
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// RecentChooserMenuNew is a wrapper around the C function gtk_recent_chooser_menu_new.
func RecentChooserMenuNew() *RecentChooserMenu {
	retC := C.gtk_recent_chooser_menu_new()
	retGo := RecentChooserMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentChooserMenuNewForManager is a wrapper around the C function gtk_recent_chooser_menu_new_for_manager.
func RecentChooserMenuNewForManager(manager *RecentManager) *RecentChooserMenu {
	c_manager := (*C.GtkRecentManager)(C.NULL)
	if manager != nil {
		c_manager = (*C.GtkRecentManager)(manager.ToC())
	}

	retC := C.gtk_recent_chooser_menu_new_for_manager(c_manager)
	retGo := RecentChooserMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentChooserWidgetNew is a wrapper around the C function gtk_recent_chooser_widget_new.
func RecentChooserWidgetNew() *RecentChooserWidget {
	retC := C.gtk_recent_chooser_widget_new()
	retGo := RecentChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentChooserWidgetNewForManager is a wrapper around the C function gtk_recent_chooser_widget_new_for_manager.
func RecentChooserWidgetNewForManager(manager *RecentManager) *RecentChooserWidget {
	c_manager := (*C.GtkRecentManager)(C.NULL)
	if manager != nil {
		c_manager = (*C.GtkRecentManager)(manager.ToC())
	}

	retC := C.gtk_recent_chooser_widget_new_for_manager(c_manager)
	retGo := RecentChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentFilterNew is a wrapper around the C function gtk_recent_filter_new.
func RecentFilterNew() *RecentFilter {
	retC := C.gtk_recent_filter_new()
	retGo := RecentFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentManager is a wrapper around the C record GtkRecentManager.
type RecentManager struct {
	native *C.GtkRecentManager
	// Private : parent_instance
	// Private : priv
}

func RecentManagerNewFromC(u unsafe.Pointer) *RecentManager {
	c := (*C.GtkRecentManager)(u)
	if c == nil {
		return nil
	}

	g := &RecentManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentManagerNew is a wrapper around the C function gtk_recent_manager_new.
func RecentManagerNew() *RecentManager {
	retC := C.gtk_recent_manager_new()
	retGo := RecentManagerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNew is a wrapper around the C function gtk_status_icon_new.
func StatusIconNew() *StatusIcon {
	retC := C.gtk_status_icon_new()
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNewFromFile is a wrapper around the C function gtk_status_icon_new_from_file.
func StatusIconNewFromFile(filename string) *StatusIcon {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_status_icon_new_from_file(c_filename)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNewFromIconName is a wrapper around the C function gtk_status_icon_new_from_icon_name.
func StatusIconNewFromIconName(iconName string) *StatusIcon {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	retC := C.gtk_status_icon_new_from_icon_name(c_icon_name)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNewFromPixbuf is a wrapper around the C function gtk_status_icon_new_from_pixbuf.
func StatusIconNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *StatusIcon {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_status_icon_new_from_pixbuf(c_pixbuf)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StatusIconNewFromStock is a wrapper around the C function gtk_status_icon_new_from_stock.
func StatusIconNewFromStock(stockId string) *StatusIcon {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_status_icon_new_from_stock(c_stock_id)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : STOCK_ORIENTATION_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_PORTRAIT

// Blacklisted : STOCK_ORIENTATION_REVERSE_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_REVERSE_PORTRAIT

// Blacklisted : STOCK_SELECT_ALL

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_target_table_free : unsupported parameter targets :

// Unsupported : gtk_target_table_new_from_list : array return type :

// Unsupported : gtk_targets_include_image : unsupported parameter targets :

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_uri : unsupported parameter targets :

// PaperSizeNew is a wrapper around the C function gtk_paper_size_new.
func PaperSizeNew(name string) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_paper_size_new(c_name)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PaperSizeNewCustom is a wrapper around the C function gtk_paper_size_new_custom.
func PaperSizeNewCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_new_custom(c_name, c_display_name, c_width, c_height, c_unit)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PaperSizeNewFromPpd is a wrapper around the C function gtk_paper_size_new_from_ppd.
func PaperSizeNewFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	c_ppd_name := C.CString(ppdName)
	defer C.free(unsafe.Pointer(c_ppd_name))

	c_ppd_display_name := C.CString(ppdDisplayName)
	defer C.free(unsafe.Pointer(c_ppd_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	retC := C.gtk_paper_size_new_from_ppd(c_ppd_name, c_ppd_display_name, c_width, c_height)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentInfo is a wrapper around the C record GtkRecentInfo.
type RecentInfo struct {
	native *C.GtkRecentInfo
}

func RecentInfoNewFromC(u unsafe.Pointer) *RecentInfo {
	c := (*C.GtkRecentInfo)(u)
	if c == nil {
		return nil
	}

	g := &RecentInfo{native: c}

	return g
}

func (recv *RecentInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentManagerClass is a wrapper around the C record GtkRecentManagerClass.
type RecentManagerClass struct {
	native *C.GtkRecentManagerClass
	// Private : parent_class
	// no type for changed
	// no type for _gtk_recent1
	// no type for _gtk_recent2
	// no type for _gtk_recent3
	// no type for _gtk_recent4
}

func RecentManagerClassNewFromC(u unsafe.Pointer) *RecentManagerClass {
	c := (*C.GtkRecentManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentManagerClass{native: c}

	return g
}

func (recv *RecentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
