// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// GetPixbuf is a wrapper around the C function gtk_selection_data_get_pixbuf.
func (recv *SelectionData) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_selection_data_get_pixbuf((*C.GtkSelectionData)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetUris is a wrapper around the C function gtk_selection_data_get_uris.
func (recv *SelectionData) GetUris() []string {
	retC := C.gtk_selection_data_get_uris((*C.GtkSelectionData)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// SetPixbuf is a wrapper around the C function gtk_selection_data_set_pixbuf.
func (recv *SelectionData) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) bool {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_selection_data_set_pixbuf((*C.GtkSelectionData)(recv.native), c_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// SetUris is a wrapper around the C function gtk_selection_data_set_uris.
func (recv *SelectionData) SetUris(uris []string) bool {
	c_uris_array := make([]*C.gchar, len(uris)+1, len(uris)+1)
	for i, item := range uris {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_uris_array[i] = c
	}
	c_uris_array[len(uris)] = nil
	c_uris_arrayPtr := &c_uris_array[0]
	c_uris := (**C.gchar)(unsafe.Pointer(c_uris_arrayPtr))

	retC := C.gtk_selection_data_set_uris((*C.GtkSelectionData)(recv.native), c_uris)
	retGo := retC == C.TRUE

	return retGo
}

// TargetsIncludeImage is a wrapper around the C function gtk_selection_data_targets_include_image.
func (recv *SelectionData) TargetsIncludeImage(writable bool) bool {
	c_writable :=
		boolToGboolean(writable)

	retC := C.gtk_selection_data_targets_include_image((*C.GtkSelectionData)(recv.native), c_writable)
	retGo := retC == C.TRUE

	return retGo
}

// AddImageTargets is a wrapper around the C function gtk_target_list_add_image_targets.
func (recv *TargetList) AddImageTargets(info uint32, writable bool) {
	c_info := (C.guint)(info)

	c_writable :=
		boolToGboolean(writable)

	C.gtk_target_list_add_image_targets((*C.GtkTargetList)(recv.native), c_info, c_writable)

	return
}

// AddTextTargets is a wrapper around the C function gtk_target_list_add_text_targets.
func (recv *TargetList) AddTextTargets(info uint32) {
	c_info := (C.guint)(info)

	C.gtk_target_list_add_text_targets((*C.GtkTargetList)(recv.native), c_info)

	return
}

// AddUriTargets is a wrapper around the C function gtk_target_list_add_uri_targets.
func (recv *TargetList) AddUriTargets(info uint32) {
	c_info := (C.guint)(info)

	C.gtk_target_list_add_uri_targets((*C.GtkTargetList)(recv.native), c_info)

	return
}
