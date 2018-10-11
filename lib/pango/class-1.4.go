// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

func (recv *Context) Object() *gobject.Object {}

func (recv *EngineLang) Engine() *Engine {}

func (recv *EngineShape) Engine() *Engine {}

func (recv *Font) Object() *gobject.Object {}

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

func (recv *FontFace) Object() *gobject.Object {}

// IsMonospace is a wrapper around the C function pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	retC := C.pango_font_family_is_monospace((*C.PangoFontFamily)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *FontFamily) Object() *gobject.Object {}

// GetShapeEngineType is a wrapper around the C function pango_font_map_get_shape_engine_type.
func (recv *FontMap) GetShapeEngineType() string {
	retC := C.pango_font_map_get_shape_engine_type((*C.PangoFontMap)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

func (recv *FontMap) Object() *gobject.Object {}

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

func (recv *Fontset) Object() *gobject.Object {}

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

func (recv *Layout) Object() *gobject.Object {}

func (recv *Renderer) Object() *gobject.Object {}
