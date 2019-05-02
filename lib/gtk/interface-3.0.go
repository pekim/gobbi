// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// GetAppInfo is a wrapper around the C function gtk_app_chooser_get_app_info.
func (recv *AppChooser) GetAppInfo() *gio.AppInfo {
	retC := C.gtk_app_chooser_get_app_info((*C.GtkAppChooser)(recv.native))
	var retGo (*gio.AppInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.AppInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

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
	c_hadjustment := (*C.GtkAdjustment)(C.NULL)
	if hadjustment != nil {
		c_hadjustment = (*C.GtkAdjustment)(hadjustment.ToC())
	}

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
	c_vadjustment := (*C.GtkAdjustment)(C.NULL)
	if vadjustment != nil {
		c_vadjustment = (*C.GtkAdjustment)(vadjustment.ToC())
	}

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
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

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
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	retC := C.gtk_style_provider_get_style((*C.GtkStyleProvider)(recv.native), c_path)
	var retGo (*StyleProperties)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StylePropertiesNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStyleProperty is a wrapper around the C function gtk_style_provider_get_style_property.
func (recv *StyleProvider) GetStyleProperty(path *WidgetPath, state StateFlags, pspec *gobject.ParamSpec) (bool, *gobject.Value) {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	c_state := (C.GtkStateFlags)(state)

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	var c_value C.GValue

	retC := C.gtk_style_provider_get_style_property((*C.GtkStyleProvider)(recv.native), c_path, c_state, c_pspec, &c_value)
	retGo := retC == C.TRUE

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value
}

// IterPrevious is a wrapper around the C function gtk_tree_model_iter_previous.
func (recv *TreeModel) IterPrevious(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_previous((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}
