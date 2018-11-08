// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : output array param sizes

// IsMonospace is a wrapper around the C function pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	retC := C.pango_font_family_is_monospace((*C.PangoFontFamily)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShapeEngineType is a wrapper around the C function pango_font_map_get_shape_engine_type.
func (recv *FontMap) GetShapeEngineType() string {
	retC := C.pango_font_map_get_shape_engine_type((*C.PangoFontMap)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc (PangoFontsetForeachFunc) for param func

// GetAutoDir is a wrapper around the C function pango_layout_get_auto_dir.
func (recv *Layout) GetAutoDir() bool {
	retC := C.pango_layout_get_auto_dir((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAutoDir is a wrapper around the C function pango_layout_set_auto_dir.
func (recv *Layout) SetAutoDir(autoDir bool) {
	c_auto_dir :=
		boolToGboolean(autoDir)

	C.pango_layout_set_auto_dir((*C.PangoLayout)(recv.native), c_auto_dir)

	return
}
