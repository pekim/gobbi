// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.16 pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

type RenderPart C.PangoRenderPart

const (
	PANGO_RENDER_PART_FOREGROUND    RenderPart = 0
	PANGO_RENDER_PART_BACKGROUND    RenderPart = 1
	PANGO_RENDER_PART_UNDERLINE     RenderPart = 2
	PANGO_RENDER_PART_STRIKETHROUGH RenderPart = 3
)
