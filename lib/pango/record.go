// Code generated - DO NOT EDIT.

package pango

type Analysis struct {
	native uintptr
	// UNSUPPORTED : C value 'shape_engine' : no Go type for 'EngineShape'

	// UNSUPPORTED : C value 'lang_engine' : no Go type for 'EngineLang'

	// UNSUPPORTED : C value 'font' : no Go type for 'Font'

	Level   uint8
	Gravity uint8
	Flags   uint8
	Script  uint8
	// UNSUPPORTED : C value 'language' : no Go type for 'Language'

	// UNSUPPORTED : C value 'extra_attrs' : no Go type for 'GLib.SList'

}

type AttrClass struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'AttrType'

	// UNSUPPORTED : C value 'copy' : missing Type

	// UNSUPPORTED : C value 'destroy' : missing Type

	// UNSUPPORTED : C value 'equal' : missing Type

}

type AttrColor struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	// UNSUPPORTED : C value 'color' : no Go type for 'Color'

}

type AttrFloat struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'

}

type AttrFontDesc struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	// UNSUPPORTED : C value 'desc' : no Go type for 'FontDescription'

}

type AttrFontFeatures struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	Features string
}

type AttrInt struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	Value int32
}

type AttrIterator struct {
	native uintptr
}

type AttrLanguage struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	// UNSUPPORTED : C value 'value' : no Go type for 'Language'

}

type AttrList struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_attr_list_new' : return type 'AttrList' not supported

type AttrShape struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	// UNSUPPORTED : C value 'ink_rect' : no Go type for 'Rectangle'

	// UNSUPPORTED : C value 'logical_rect' : no Go type for 'Rectangle'

	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'copy_func' : no Go type for 'AttrDataCopyFunc'

	// UNSUPPORTED : C value 'destroy_func' : no Go type for 'GLib.DestroyNotify'

}

type AttrSize struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	Size     int32
	Absolute uint32
}

type AttrString struct {
	native uintptr
	// UNSUPPORTED : C value 'attr' : no Go type for 'Attribute'

	Value string
}

type Attribute struct {
	native uintptr
	// UNSUPPORTED : C value 'klass' : no Go type for 'AttrClass'

	StartIndex uint32
	EndIndex   uint32
}

type Color struct {
	native uintptr
	Red    uint16
	Green  uint16
	Blue   uint16
}

type ContextClass struct {
	native uintptr
}

type Coverage struct {
	native uintptr
}

type EngineClass struct {
	native uintptr
}

type EngineInfo struct {
	native     uintptr
	Id         string
	EngineType string
	RenderType string
	// UNSUPPORTED : C value 'scripts' : no Go type for 'EngineScriptInfo'

	NScripts int32
}

type EngineLangClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_break' : missing Type

}

type EngineScriptInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'script' : no Go type for 'Script'

	Langs string
}

type EngineShapeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_shape' : missing Type

	// UNSUPPORTED : C value 'covers' : missing Type

}

type FontClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'describe' : missing Type

	// UNSUPPORTED : C value 'get_coverage' : missing Type

	// UNSUPPORTED : C value 'find_shaper' : missing Type

	// UNSUPPORTED : C value 'get_glyph_extents' : missing Type

	// UNSUPPORTED : C value 'get_metrics' : missing Type

	// UNSUPPORTED : C value 'get_font_map' : missing Type

	// UNSUPPORTED : C value 'describe_absolute' : missing Type

	// UNSUPPORTED : C value '_pango_reserved1' : missing Type

	// UNSUPPORTED : C value '_pango_reserved2' : missing Type

}

type FontDescription struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_font_description_new' : return type 'FontDescription' not supported

type FontFaceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_face_name' : missing Type

	// UNSUPPORTED : C value 'describe' : missing Type

	// UNSUPPORTED : C value 'list_sizes' : missing Type

	// UNSUPPORTED : C value 'is_synthesized' : missing Type

	// UNSUPPORTED : C value '_pango_reserved3' : missing Type

	// UNSUPPORTED : C value '_pango_reserved4' : missing Type

}

type FontFamilyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'list_faces' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'is_monospace' : missing Type

	// UNSUPPORTED : C value '_pango_reserved2' : missing Type

	// UNSUPPORTED : C value '_pango_reserved3' : missing Type

	// UNSUPPORTED : C value '_pango_reserved4' : missing Type

}

type FontMapClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'load_font' : missing Type

	// UNSUPPORTED : C value 'list_families' : missing Type

	// UNSUPPORTED : C value 'load_fontset' : missing Type

	ShapeEngineType string
	// UNSUPPORTED : C value 'get_serial' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_pango_reserved1' : missing Type

	// UNSUPPORTED : C value '_pango_reserved2' : missing Type

}

type FontMetrics struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_font_metrics_new' : return type 'FontMetrics' not supported

type FontsetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_font' : missing Type

	// UNSUPPORTED : C value 'get_metrics' : missing Type

	// UNSUPPORTED : C value 'get_language' : missing Type

	// UNSUPPORTED : C value 'foreach' : missing Type

	// UNSUPPORTED : C value '_pango_reserved1' : missing Type

	// UNSUPPORTED : C value '_pango_reserved2' : missing Type

	// UNSUPPORTED : C value '_pango_reserved3' : missing Type

	// UNSUPPORTED : C value '_pango_reserved4' : missing Type

}

type FontsetSimpleClass struct {
	native uintptr
}

type GlyphGeometry struct {
	native  uintptr
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

type GlyphInfo struct {
	native uintptr
	Glyph  Glyph
	// UNSUPPORTED : C value 'geometry' : no Go type for 'GlyphGeometry'

	// UNSUPPORTED : C value 'attr' : no Go type for 'GlyphVisAttr'

}

type GlyphItem struct {
	native uintptr
	// UNSUPPORTED : C value 'item' : no Go type for 'Item'

	// UNSUPPORTED : C value 'glyphs' : no Go type for 'GlyphString'

}

type GlyphItemIter struct {
	native uintptr
	// UNSUPPORTED : C value 'glyph_item' : no Go type for 'GlyphItem'

	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

type GlyphString struct {
	native    uintptr
	NumGlyphs int32
	// UNSUPPORTED : C value 'glyphs' : missing Type

	LogClusters int32
}

// UNSUPPORTED : C value 'pango_glyph_string_new' : return type 'GlyphString' not supported

type GlyphVisAttr struct {
	native         uintptr
	IsClusterStart uint32
}

type IncludedModule struct {
	native uintptr
	// UNSUPPORTED : C value 'list' : missing Type

	// UNSUPPORTED : C value 'init' : missing Type

	// UNSUPPORTED : C value 'exit' : missing Type

	// UNSUPPORTED : C value 'create' : missing Type

}

type Item struct {
	native   uintptr
	Offset   int32
	Length   int32
	NumChars int32
	// UNSUPPORTED : C value 'analysis' : no Go type for 'Analysis'

}

// UNSUPPORTED : C value 'pango_item_new' : return type 'Item' not supported

type Language struct {
	native uintptr
}

type LayoutClass struct {
	native uintptr
}

type LayoutIter struct {
	native uintptr
}

type LayoutLine struct {
	native uintptr
	// UNSUPPORTED : C value 'layout' : no Go type for 'Layout'

	StartIndex int32
	Length     int32
	// UNSUPPORTED : C value 'runs' : no Go type for 'GLib.SList'

	IsParagraphStart uint32
	ResolvedDir      uint32
}

type LogAttr struct {
	native                    uintptr
	IsLineBreak               uint32
	IsMandatoryBreak          uint32
	IsCharBreak               uint32
	IsWhite                   uint32
	IsCursorPosition          uint32
	IsWordStart               uint32
	IsWordEnd                 uint32
	IsSentenceBoundary        uint32
	IsSentenceStart           uint32
	IsSentenceEnd             uint32
	BackspaceDeletesCharacter uint32
	IsExpandableSpace         uint32
	IsWordBoundary            uint32
}

type Map struct {
	native uintptr
}

type MapEntry struct {
	native uintptr
}

type Matrix struct {
	native uintptr
	// UNSUPPORTED : C value 'xx' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'xy' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'yx' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'yy' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'x0' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y0' : no Go type for 'gdouble'

}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type RendererClass struct {
	native uintptr
	// UNSUPPORTED : C value 'draw_glyphs' : missing Type

	// UNSUPPORTED : C value 'draw_rectangle' : missing Type

	// UNSUPPORTED : C value 'draw_error_underline' : missing Type

	// UNSUPPORTED : C value 'draw_shape' : missing Type

	// UNSUPPORTED : C value 'draw_trapezoid' : missing Type

	// UNSUPPORTED : C value 'draw_glyph' : missing Type

	// UNSUPPORTED : C value 'part_changed' : missing Type

	// UNSUPPORTED : C value 'begin' : missing Type

	// UNSUPPORTED : C value 'end' : missing Type

	// UNSUPPORTED : C value 'prepare_run' : missing Type

	// UNSUPPORTED : C value 'draw_glyph_item' : missing Type

	// UNSUPPORTED : C value '_pango_reserved2' : missing Type

	// UNSUPPORTED : C value '_pango_reserved3' : missing Type

	// UNSUPPORTED : C value '_pango_reserved4' : missing Type

}

type RendererPrivate struct {
	native uintptr
}

type ScriptIter struct {
	native uintptr
}

type TabArray struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_tab_array_new' : parameter 'positions_in_pixels' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'positions_in_pixels' of type 'gboolean' not supported
