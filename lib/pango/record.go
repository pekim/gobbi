// Code generated - DO NOT EDIT.

package pango

import gi "github.com/pekim/gobbi/internal/gi"

type Analysis struct {
	native uintptr
	// UNSUPPORTED : C value 'shape_engine' : no Go type for 'EngineShape'
	// UNSUPPORTED : C value 'lang_engine' : no Go type for 'EngineLang'
	// UNSUPPORTED : C value 'font' : no Go type for 'Font'
	Level    uint8
	Gravity  uint8
	Flags    uint8
	Script   uint8
	Language *Language
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
	Attr   *Attribute
	Color  *Color
}

type AttrFloat struct {
	native uintptr
	Attr   *Attribute
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'
}

type AttrFontDesc struct {
	native uintptr
	Attr   *Attribute
	Desc   *FontDescription
}

type AttrFontFeatures struct {
	native   uintptr
	Attr     *Attribute
	Features string
}

type AttrInt struct {
	native uintptr
	Attr   *Attribute
	Value  int32
}

type AttrIterator struct {
	native uintptr
}

var copyAttrIteratorInvoker *gi.Function

// Copy is a representation of the C type pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	if copyAttrIteratorInvoker == nil {
		copyAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttrIteratorInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

var destroyAttrIteratorInvoker *gi.Function

// Destroy is a representation of the C type pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	if destroyAttrIteratorInvoker == nil {
		destroyAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyAttrIteratorInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attr_iterator_get' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'desc' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_next' : return type 'gboolean' not supported

var rangeAttrIteratorInvoker *gi.Function

// Range is a representation of the C type pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	if rangeAttrIteratorInvoker == nil {
		rangeAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "range")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	rangeAttrIteratorInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

type AttrLanguage struct {
	native uintptr
	Attr   *Attribute
	Value  *Language
}

type AttrList struct {
	native uintptr
}

var newAttrListInvoker *gi.Function

// AttrListNew is a representation of the C type pango_attr_list_new.
func AttrListNew() *AttrList {
	if newAttrListInvoker == nil {
		newAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "new")
	}

	ret := newAttrListInvoker.Invoke(nil, nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_change' : parameter 'attr' of type 'Attribute' not supported

var copyAttrListInvoker *gi.Function

// Copy is a representation of the C type pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	if copyAttrListInvoker == nil {
		copyAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_filter' : parameter 'func' of type 'AttrFilterFunc' not supported

var getIteratorAttrListInvoker *gi.Function

// GetIterator is a representation of the C type pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	if getIteratorAttrListInvoker == nil {
		getIteratorAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "get_iterator")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIteratorAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_insert' : parameter 'attr' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_list_insert_before' : parameter 'attr' of type 'Attribute' not supported

var refAttrListInvoker *gi.Function

// Ref is a representation of the C type pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	if refAttrListInvoker == nil {
		refAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_splice' : parameter 'other' of type 'AttrList' not supported

var unrefAttrListInvoker *gi.Function

// Unref is a representation of the C type pango_attr_list_unref.
func (recv *AttrList) Unref() {
	if unrefAttrListInvoker == nil {
		unrefAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefAttrListInvoker.Invoke(inArgs[:], nil)

}

type AttrShape struct {
	native      uintptr
	Attr        *Attribute
	InkRect     *Rectangle
	LogicalRect *Rectangle
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'copy_func' : no Go type for 'AttrDataCopyFunc'
	// UNSUPPORTED : C value 'destroy_func' : no Go type for 'GLib.DestroyNotify'
}

type AttrSize struct {
	native   uintptr
	Attr     *Attribute
	Size     int32
	Absolute uint32
}

type AttrString struct {
	native uintptr
	Attr   *Attribute
	Value  string
}

type Attribute struct {
	native     uintptr
	Klass      *AttrClass
	StartIndex uint32
	EndIndex   uint32
}

var copyAttributeInvoker *gi.Function

// Copy is a representation of the C type pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	if copyAttributeInvoker == nil {
		copyAttributeInvoker = gi.StructFunctionInvokerNew("Pango", "Attribute", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttributeInvoker.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var destroyAttributeInvoker *gi.Function

// Destroy is a representation of the C type pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	if destroyAttributeInvoker == nil {
		destroyAttributeInvoker = gi.StructFunctionInvokerNew("Pango", "Attribute", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyAttributeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attribute_equal' : parameter 'attr2' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attribute_init' : parameter 'klass' of type 'AttrClass' not supported

type Color struct {
	native uintptr
	Red    uint16
	Green  uint16
	Blue   uint16
}

var copyColorInvoker *gi.Function

// Copy is a representation of the C type pango_color_copy.
func (recv *Color) Copy() *Color {
	if copyColorInvoker == nil {
		copyColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyColorInvoker.Invoke(inArgs[:], nil)

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

var freeColorInvoker *gi.Function

// Free is a representation of the C type pango_color_free.
func (recv *Color) Free() {
	if freeColorInvoker == nil {
		freeColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeColorInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_color_parse' : return type 'gboolean' not supported

var toStringColorInvoker *gi.Function

// ToString is a representation of the C type pango_color_to_string.
func (recv *Color) ToString() string {
	if toStringColorInvoker == nil {
		toStringColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringColorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

type ContextClass struct {
	native uintptr
}

type Coverage struct {
	native uintptr
}

var copyCoverageInvoker *gi.Function

// Copy is a representation of the C type pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	if copyCoverageInvoker == nil {
		copyCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyCoverageInvoker.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_get' : return type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_max' : parameter 'other' of type 'Coverage' not supported

var refCoverageInvoker *gi.Function

// Ref is a representation of the C type pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	if refCoverageInvoker == nil {
		refCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refCoverageInvoker.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_set' : parameter 'level' of type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_to_bytes' : parameter 'bytes' has no type

var unrefCoverageInvoker *gi.Function

// Unref is a representation of the C type pango_coverage_unref.
func (recv *Coverage) Unref() {
	if unrefCoverageInvoker == nil {
		unrefCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefCoverageInvoker.Invoke(inArgs[:], nil)

}

type EngineClass struct {
	native uintptr
}

type EngineInfo struct {
	native     uintptr
	Id         string
	EngineType string
	RenderType string
	Scripts    *EngineScriptInfo
	NScripts   int32
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

var newFontDescriptionInvoker *gi.Function

// FontDescriptionNew is a representation of the C type pango_font_description_new.
func FontDescriptionNew() *FontDescription {
	if newFontDescriptionInvoker == nil {
		newFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "new")
	}

	ret := newFontDescriptionInvoker.Invoke(nil, nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_better_match' : parameter 'old_match' of type 'FontDescription' not supported

var copyFontDescriptionInvoker *gi.Function

// Copy is a representation of the C type pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	if copyFontDescriptionInvoker == nil {
		copyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var copyStaticFontDescriptionInvoker *gi.Function

// CopyStatic is a representation of the C type pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	if copyStaticFontDescriptionInvoker == nil {
		copyStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "copy_static")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_equal' : parameter 'desc2' of type 'FontDescription' not supported

var freeFontDescriptionInvoker *gi.Function

// Free is a representation of the C type pango_font_description_free.
func (recv *FontDescription) Free() {
	if freeFontDescriptionInvoker == nil {
		freeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var getFamilyFontDescriptionInvoker *gi.Function

// GetFamily is a representation of the C type pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	if getFamilyFontDescriptionInvoker == nil {
		getFamilyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_family")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFamilyFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_gravity' : return type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_font_description_get_set_fields' : return type 'FontMask' not supported

var getSizeFontDescriptionInvoker *gi.Function

// GetSize is a representation of the C type pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	if getSizeFontDescriptionInvoker == nil {
		getSizeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSizeFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_size_is_absolute' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_font_description_get_stretch' : return type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_get_style' : return type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_get_variant' : return type 'Variant' not supported

var getVariationsFontDescriptionInvoker *gi.Function

// GetVariations is a representation of the C type pango_font_description_get_variations.
func (recv *FontDescription) GetVariations() string {
	if getVariationsFontDescriptionInvoker == nil {
		getVariationsFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_variations")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVariationsFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_weight' : return type 'Weight' not supported

var hashFontDescriptionInvoker *gi.Function

// Hash is a representation of the C type pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	if hashFontDescriptionInvoker == nil {
		hashFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_merge' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_merge_static' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_set_absolute_size' : parameter 'size' of type 'gdouble' not supported

var setFamilyFontDescriptionInvoker *gi.Function

// SetFamily is a representation of the C type pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	if setFamilyFontDescriptionInvoker == nil {
		setFamilyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_family")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	setFamilyFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var setFamilyStaticFontDescriptionInvoker *gi.Function

// SetFamilyStatic is a representation of the C type pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	if setFamilyStaticFontDescriptionInvoker == nil {
		setFamilyStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_family_static")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	setFamilyStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_gravity' : parameter 'gravity' of type 'Gravity' not supported

var setSizeFontDescriptionInvoker *gi.Function

// SetSize is a representation of the C type pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	if setSizeFontDescriptionInvoker == nil {
		setSizeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_size")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	setSizeFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_stretch' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_set_style' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_set_variant' : parameter 'variant' of type 'Variant' not supported

var setVariationsFontDescriptionInvoker *gi.Function

// SetVariations is a representation of the C type pango_font_description_set_variations.
func (recv *FontDescription) SetVariations(settings string) {
	if setVariationsFontDescriptionInvoker == nil {
		setVariationsFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_variations")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	setVariationsFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var setVariationsStaticFontDescriptionInvoker *gi.Function

// SetVariationsStatic is a representation of the C type pango_font_description_set_variations_static.
func (recv *FontDescription) SetVariationsStatic(settings string) {
	if setVariationsStaticFontDescriptionInvoker == nil {
		setVariationsStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_variations_static")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	setVariationsStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_weight' : parameter 'weight' of type 'Weight' not supported

var toFilenameFontDescriptionInvoker *gi.Function

// ToFilename is a representation of the C type pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	if toFilenameFontDescriptionInvoker == nil {
		toFilenameFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "to_filename")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toFilenameFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var toStringFontDescriptionInvoker *gi.Function

// ToString is a representation of the C type pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	if toStringFontDescriptionInvoker == nil {
		toStringFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_unset_fields' : parameter 'to_unset' of type 'FontMask' not supported

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

var newFontMetricsInvoker *gi.Function

// FontMetricsNew is a representation of the C type pango_font_metrics_new.
func FontMetricsNew() *FontMetrics {
	if newFontMetricsInvoker == nil {
		newFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "new")
	}

	ret := newFontMetricsInvoker.Invoke(nil, nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var getApproximateCharWidthFontMetricsInvoker *gi.Function

// GetApproximateCharWidth is a representation of the C type pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	if getApproximateCharWidthFontMetricsInvoker == nil {
		getApproximateCharWidthFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_approximate_char_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getApproximateCharWidthFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getApproximateDigitWidthFontMetricsInvoker *gi.Function

// GetApproximateDigitWidth is a representation of the C type pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	if getApproximateDigitWidthFontMetricsInvoker == nil {
		getApproximateDigitWidthFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_approximate_digit_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getApproximateDigitWidthFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getAscentFontMetricsInvoker *gi.Function

// GetAscent is a representation of the C type pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	if getAscentFontMetricsInvoker == nil {
		getAscentFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_ascent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAscentFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getDescentFontMetricsInvoker *gi.Function

// GetDescent is a representation of the C type pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	if getDescentFontMetricsInvoker == nil {
		getDescentFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_descent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescentFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getStrikethroughPositionFontMetricsInvoker *gi.Function

// GetStrikethroughPosition is a representation of the C type pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	if getStrikethroughPositionFontMetricsInvoker == nil {
		getStrikethroughPositionFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_strikethrough_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStrikethroughPositionFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getStrikethroughThicknessFontMetricsInvoker *gi.Function

// GetStrikethroughThickness is a representation of the C type pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	if getStrikethroughThicknessFontMetricsInvoker == nil {
		getStrikethroughThicknessFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_strikethrough_thickness")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStrikethroughThicknessFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getUnderlinePositionFontMetricsInvoker *gi.Function

// GetUnderlinePosition is a representation of the C type pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	if getUnderlinePositionFontMetricsInvoker == nil {
		getUnderlinePositionFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_underline_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUnderlinePositionFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getUnderlineThicknessFontMetricsInvoker *gi.Function

// GetUnderlineThickness is a representation of the C type pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	if getUnderlineThicknessFontMetricsInvoker == nil {
		getUnderlineThicknessFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_underline_thickness")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUnderlineThicknessFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var refFontMetricsInvoker *gi.Function

// Ref is a representation of the C type pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	if refFontMetricsInvoker == nil {
		refFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var unrefFontMetricsInvoker *gi.Function

// Unref is a representation of the C type pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	if unrefFontMetricsInvoker == nil {
		unrefFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefFontMetricsInvoker.Invoke(inArgs[:], nil)

}

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
	native   uintptr
	Glyph    Glyph
	Geometry *GlyphGeometry
	Attr     *GlyphVisAttr
}

type GlyphItem struct {
	native uintptr
	Item   *Item
	Glyphs *GlyphString
}

// UNSUPPORTED : C value 'pango_glyph_item_apply_attrs' : parameter 'list' of type 'AttrList' not supported

var copyGlyphItemInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	if copyGlyphItemInvoker == nil {
		copyGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphItemInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var freeGlyphItemInvoker *gi.Function

// Free is a representation of the C type pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	if freeGlyphItemInvoker == nil {
		freeGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphItemInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_get_logical_widths' : parameter 'logical_widths' has no type

// UNSUPPORTED : C value 'pango_glyph_item_letter_space' : parameter 'log_attrs' has no type

var splitGlyphItemInvoker *gi.Function

// Split is a representation of the C type pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	if splitGlyphItemInvoker == nil {
		splitGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "split")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(splitIndex)

	ret := splitGlyphItemInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

type GlyphItemIter struct {
	native     uintptr
	GlyphItem  *GlyphItem
	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

var copyGlyphItemIterInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	if copyGlyphItemIterInvoker == nil {
		copyGlyphItemIterInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItemIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphItemIterInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItemIter{native: ret.Pointer()}

	return retGo
}

var freeGlyphItemIterInvoker *gi.Function

// Free is a representation of the C type pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	if freeGlyphItemIterInvoker == nil {
		freeGlyphItemIterInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItemIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphItemIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_end' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_start' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_prev_cluster' : return type 'gboolean' not supported

type GlyphString struct {
	native    uintptr
	NumGlyphs int32
	// UNSUPPORTED : C value 'glyphs' : missing Type
	LogClusters int32
}

var newGlyphStringInvoker *gi.Function

// GlyphStringNew is a representation of the C type pango_glyph_string_new.
func GlyphStringNew() *GlyphString {
	if newGlyphStringInvoker == nil {
		newGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "new")
	}

	ret := newGlyphStringInvoker.Invoke(nil, nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

var copyGlyphStringInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	if copyGlyphStringInvoker == nil {
		copyGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphStringInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_extents' : parameter 'font' of type 'Font' not supported

// UNSUPPORTED : C value 'pango_glyph_string_extents_range' : parameter 'font' of type 'Font' not supported

var freeGlyphStringInvoker *gi.Function

// Free is a representation of the C type pango_glyph_string_free.
func (recv *GlyphString) Free() {
	if freeGlyphStringInvoker == nil {
		freeGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphStringInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_get_logical_widths' : parameter 'logical_widths' has no type

var getWidthGlyphStringInvoker *gi.Function

// GetWidth is a representation of the C type pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	if getWidthGlyphStringInvoker == nil {
		getWidthGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "get_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getWidthGlyphStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_index_to_x' : parameter 'analysis' of type 'Analysis' not supported

var setSizeGlyphStringInvoker *gi.Function

// SetSize is a representation of the C type pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	if setSizeGlyphStringInvoker == nil {
		setSizeGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "set_size")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	setSizeGlyphStringInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_x_to_index' : parameter 'analysis' of type 'Analysis' not supported

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
	Analysis *Analysis
}

var newItemInvoker *gi.Function

// ItemNew is a representation of the C type pango_item_new.
func ItemNew() *Item {
	if newItemInvoker == nil {
		newItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "new")
	}

	ret := newItemInvoker.Invoke(nil, nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var copyItemInvoker *gi.Function

// Copy is a representation of the C type pango_item_copy.
func (recv *Item) Copy() *Item {
	if copyItemInvoker == nil {
		copyItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyItemInvoker.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var freeItemInvoker *gi.Function

// Free is a representation of the C type pango_item_free.
func (recv *Item) Free() {
	if freeItemInvoker == nil {
		freeItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeItemInvoker.Invoke(inArgs[:], nil)

}

var splitItemInvoker *gi.Function

// Split is a representation of the C type pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	if splitItemInvoker == nil {
		splitItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "split")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(splitIndex)
	inArgs[2].SetInt32(splitOffset)

	ret := splitItemInvoker.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

type Language struct {
	native uintptr
}

var getSampleStringLanguageInvoker *gi.Function

// GetSampleString is a representation of the C type pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	if getSampleStringLanguageInvoker == nil {
		getSampleStringLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "get_sample_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSampleStringLanguageInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getScriptsLanguageInvoker *gi.Function

// GetScripts is a representation of the C type pango_language_get_scripts.
func (recv *Language) GetScripts() int32 {
	if getScriptsLanguageInvoker == nil {
		getScriptsLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "get_scripts")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	getScriptsLanguageInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'pango_language_includes_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_language_matches' : return type 'gboolean' not supported

var toStringLanguageInvoker *gi.Function

// ToString is a representation of the C type pango_language_to_string.
func (recv *Language) ToString() string {
	if toStringLanguageInvoker == nil {
		toStringLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringLanguageInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

type LayoutClass struct {
	native uintptr
}

type LayoutIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_layout_iter_at_last_line' : return type 'gboolean' not supported

var copyLayoutIterInvoker *gi.Function

// Copy is a representation of the C type pango_layout_iter_copy.
func (recv *LayoutIter) Copy() *LayoutIter {
	if copyLayoutIterInvoker == nil {
		copyLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo
}

var freeLayoutIterInvoker *gi.Function

// Free is a representation of the C type pango_layout_iter_free.
func (recv *LayoutIter) Free() {
	if freeLayoutIterInvoker == nil {
		freeLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeLayoutIterInvoker.Invoke(inArgs[:], nil)

}

var getBaselineLayoutIterInvoker *gi.Function

// GetBaseline is a representation of the C type pango_layout_iter_get_baseline.
func (recv *LayoutIter) GetBaseline() int32 {
	if getBaselineLayoutIterInvoker == nil {
		getBaselineLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_baseline")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getBaselineLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_char_extents' : parameter 'logical_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_cluster_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var getIndexLayoutIterInvoker *gi.Function

// GetIndex is a representation of the C type pango_layout_iter_get_index.
func (recv *LayoutIter) GetIndex() int32 {
	if getIndexLayoutIterInvoker == nil {
		getIndexLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_index")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIndexLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_layout' : return type 'Layout' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_layout_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var getLineLayoutIterInvoker *gi.Function

// GetLine is a representation of the C type pango_layout_iter_get_line.
func (recv *LayoutIter) GetLine() *LayoutLine {
	if getLineLayoutIterInvoker == nil {
		getLineLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_line_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var getLineReadonlyLayoutIterInvoker *gi.Function

// GetLineReadonly is a representation of the C type pango_layout_iter_get_line_readonly.
func (recv *LayoutIter) GetLineReadonly() *LayoutLine {
	if getLineReadonlyLayoutIterInvoker == nil {
		getLineReadonlyLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line_readonly")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineReadonlyLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var getLineYrangeLayoutIterInvoker *gi.Function

// GetLineYrange is a representation of the C type pango_layout_iter_get_line_yrange.
func (recv *LayoutIter) GetLineYrange() (int32, int32) {
	if getLineYrangeLayoutIterInvoker == nil {
		getLineYrangeLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line_yrange")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	getLineYrangeLayoutIterInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run' : return type 'LayoutRun' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_run_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_run_readonly' : return type 'LayoutRun' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_char' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_run' : return type 'gboolean' not supported

type LayoutLine struct {
	native uintptr
	// UNSUPPORTED : C value 'layout' : no Go type for 'Layout'
	StartIndex int32
	Length     int32
	// UNSUPPORTED : C value 'runs' : no Go type for 'GLib.SList'
	IsParagraphStart uint32
	ResolvedDir      uint32
}

// UNSUPPORTED : C value 'pango_layout_line_get_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_line_get_pixel_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_line_get_x_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'pango_layout_line_index_to_x' : parameter 'trailing' of type 'gboolean' not supported

var refLayoutLineInvoker *gi.Function

// Ref is a representation of the C type pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	if refLayoutLineInvoker == nil {
		refLayoutLineInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutLine", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refLayoutLineInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var unrefLayoutLineInvoker *gi.Function

// Unref is a representation of the C type pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	if unrefLayoutLineInvoker == nil {
		unrefLayoutLineInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutLine", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefLayoutLineInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_layout_line_x_to_index' : return type 'gboolean' not supported

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

// UNSUPPORTED : C value 'pango_map_get_engine' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_map_get_engines' : parameter 'script' of type 'Script' not supported

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

// UNSUPPORTED : C value 'pango_matrix_concat' : parameter 'new_matrix' of type 'Matrix' not supported

var copyMatrixInvoker *gi.Function

// Copy is a representation of the C type pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	if copyMatrixInvoker == nil {
		copyMatrixInvoker = gi.StructFunctionInvokerNew("Pango", "Matrix", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyMatrixInvoker.Invoke(inArgs[:], nil)

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

var freeMatrixInvoker *gi.Function

// Free is a representation of the C type pango_matrix_free.
func (recv *Matrix) Free() {
	if freeMatrixInvoker == nil {
		freeMatrixInvoker = gi.StructFunctionInvokerNew("Pango", "Matrix", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMatrixInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_matrix_get_font_scale_factor' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_get_font_scale_factors' : parameter 'xscale' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_rotate' : parameter 'degrees' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_scale' : parameter 'scale_x' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_distance' : parameter 'dx' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_pixel_rectangle' : parameter 'rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_point' : parameter 'x' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_rectangle' : parameter 'rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_matrix_translate' : parameter 'tx' of type 'gdouble' not supported

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

var freeScriptIterInvoker *gi.Function

// Free is a representation of the C type pango_script_iter_free.
func (recv *ScriptIter) Free() {
	if freeScriptIterInvoker == nil {
		freeScriptIterInvoker = gi.StructFunctionInvokerNew("Pango", "ScriptIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeScriptIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_script_iter_get_range' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_script_iter_next' : return type 'gboolean' not supported

type TabArray struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_tab_array_new' : parameter 'positions_in_pixels' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'positions_in_pixels' of type 'gboolean' not supported

var copyTabArrayInvoker *gi.Function

// Copy is a representation of the C type pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	if copyTabArrayInvoker == nil {
		copyTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTabArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &TabArray{native: ret.Pointer()}

	return retGo
}

var freeTabArrayInvoker *gi.Function

// Free is a representation of the C type pango_tab_array_free.
func (recv *TabArray) Free() {
	if freeTabArrayInvoker == nil {
		freeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTabArrayInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_get_positions_in_pixels' : return type 'gboolean' not supported

var getSizeTabArrayInvoker *gi.Function

// GetSize is a representation of the C type pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	if getSizeTabArrayInvoker == nil {
		getSizeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "get_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSizeTabArrayInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_tab_array_get_tab' : parameter 'alignment' of type 'TabAlign' not supported

// UNSUPPORTED : C value 'pango_tab_array_get_tabs' : parameter 'alignments' of type 'TabAlign' not supported

var resizeTabArrayInvoker *gi.Function

// Resize is a representation of the C type pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	if resizeTabArrayInvoker == nil {
		resizeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "resize")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	resizeTabArrayInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_set_tab' : parameter 'alignment' of type 'TabAlign' not supported
