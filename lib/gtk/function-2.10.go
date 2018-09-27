// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// PaperSizeGetDefault is a wrapper around the C function gtk_paper_size_get_default.
func PaperSizeGetDefault() string {
	retC := C.gtk_paper_size_get_default()
	retGo := C.GoString(retC)

	return retGo
}

// PrintErrorQuark is a wrapper around the C function gtk_print_error_quark.
func PrintErrorQuark() glib.Quark {
	retC := C.gtk_print_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// PrintRunPageSetupDialog is a wrapper around the C function gtk_print_run_page_setup_dialog.
func PrintRunPageSetupDialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) *PageSetup {
	c_parent := (*C.GtkWindow)(parent.ToC())

	c_page_setup := (*C.GtkPageSetup)(pageSetup.ToC())

	c_settings := (*C.GtkPrintSettings)(settings.ToC())

	retC := C.gtk_print_run_page_setup_dialog(c_parent, c_page_setup, c_settings)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc, GtkPageSetupDoneFunc

// Unsupported : gtk_target_table_free : unsupported parameter targets : no param type

// Unsupported : gtk_target_table_new_from_list : unsupported parameter n_targets : no type generator for gint, gint*

// Unsupported : gtk_targets_include_image : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_uri : unsupported parameter targets : no param type
