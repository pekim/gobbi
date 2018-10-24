// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_app_chooser_get_app_info : no return generator

// GetContentType is a wrapper around the C function gtk_app_chooser_get_content_type.
func (recv *AppChooser) GetContentType() string {
	retC := C.gtk_app_chooser_get_content_type((*C.GtkAppChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Refresh is a wrapper around the C function gtk_app_chooser_refresh.
func (recv *AppChooser) Refresh() {
	C.gtk_app_chooser_refresh((*C.GtkAppChooser)(recv.native))

	return
}

// GetArea is a wrapper around the C function gtk_cell_layout_get_area.
func (recv *CellLayout) GetArea() *CellArea {
	retC := C.gtk_cell_layout_get_area((*C.GtkCellLayout)(recv.native))
	var retGo (*CellArea)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CellAreaNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_scrollable_get_hadjustment.
func (recv *Scrollable) GetHadjustment() *Adjustment {
	retC := C.gtk_scrollable_get_hadjustment((*C.GtkScrollable)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHscrollPolicy is a wrapper around the C function gtk_scrollable_get_hscroll_policy.
func (recv *Scrollable) GetHscrollPolicy() ScrollablePolicy {
	retC := C.gtk_scrollable_get_hscroll_policy((*C.GtkScrollable)(recv.native))
	retGo := (ScrollablePolicy)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_scrollable_get_vadjustment.
func (recv *Scrollable) GetVadjustment() *Adjustment {
	retC := C.gtk_scrollable_get_vadjustment((*C.GtkScrollable)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVscrollPolicy is a wrapper around the C function gtk_scrollable_get_vscroll_policy.
func (recv *Scrollable) GetVscrollPolicy() ScrollablePolicy {
	retC := C.gtk_scrollable_get_vscroll_policy((*C.GtkScrollable)(recv.native))
	retGo := (ScrollablePolicy)(retC)

	return retGo
}

// SetHadjustment is a wrapper around the C function gtk_scrollable_set_hadjustment.
func (recv *Scrollable) SetHadjustment(hadjustment *Adjustment) {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	C.gtk_scrollable_set_hadjustment((*C.GtkScrollable)(recv.native), c_hadjustment)

	return
}

// SetHscrollPolicy is a wrapper around the C function gtk_scrollable_set_hscroll_policy.
func (recv *Scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	c_policy := (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_hscroll_policy((*C.GtkScrollable)(recv.native), c_policy)

	return
}

// SetVadjustment is a wrapper around the C function gtk_scrollable_set_vadjustment.
func (recv *Scrollable) SetVadjustment(vadjustment *Adjustment) {
	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	C.gtk_scrollable_set_vadjustment((*C.GtkScrollable)(recv.native), c_vadjustment)

	return
}

// SetVscrollPolicy is a wrapper around the C function gtk_scrollable_set_vscroll_policy.
func (recv *Scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	c_policy := (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_vscroll_policy((*C.GtkScrollable)(recv.native), c_policy)

	return
}

// GetIconFactory is a wrapper around the C function gtk_style_provider_get_icon_factory.
func (recv *StyleProvider) GetIconFactory(path *WidgetPath) *IconFactory {
	c_path := (*C.GtkWidgetPath)(path.ToC())

	retC := C.gtk_style_provider_get_icon_factory((*C.GtkStyleProvider)(recv.native), c_path)
	var retGo (*IconFactory)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconFactoryNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStyle is a wrapper around the C function gtk_style_provider_get_style.
func (recv *StyleProvider) GetStyle(path *WidgetPath) *StyleProperties {
	c_path := (*C.GtkWidgetPath)(path.ToC())

	retC := C.gtk_style_provider_get_style((*C.GtkStyleProvider)(recv.native), c_path)
	var retGo (*StyleProperties)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StylePropertiesNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_style_provider_get_style_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// IterPrevious is a wrapper around the C function gtk_tree_model_iter_previous.
func (recv *TreeModel) IterPrevious(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_model_iter_previous((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}
