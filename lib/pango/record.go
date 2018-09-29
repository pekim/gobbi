// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Analysis is a wrapper around the C record PangoAnalysis.
type Analysis struct {
	native *C.PangoAnalysis
	// shape_engine : record
	// lang_engine : record
	// font : record
	Level   uint8
	Gravity uint8
	Flags   uint8
	Script  uint8
	// language : record
	// extra_attrs : record
}

func AnalysisNewFromC(u unsafe.Pointer) *Analysis {
	c := (*C.PangoAnalysis)(u)
	if c == nil {
		return nil
	}

	g := &Analysis{
		Flags:   (uint8)(c.flags),
		Gravity: (uint8)(c.gravity),
		Level:   (uint8)(c.level),
		Script:  (uint8)(c.script),
		native:  c,
	}

	return g
}

func (recv *Analysis) ToC() unsafe.Pointer {
	recv.native.level =
		(C.guint8)(recv.Level)
	recv.native.gravity =
		(C.guint8)(recv.Gravity)
	recv.native.flags =
		(C.guint8)(recv.Flags)
	recv.native.script =
		(C.guint8)(recv.Script)

	return (unsafe.Pointer)(recv.native)
}

// AttrClass is a wrapper around the C record PangoAttrClass.
type AttrClass struct {
	native *C.PangoAttrClass
	Type   AttrType
	// no type for copy
	// no type for destroy
	// no type for equal
}

func AttrClassNewFromC(u unsafe.Pointer) *AttrClass {
	c := (*C.PangoAttrClass)(u)
	if c == nil {
		return nil
	}

	g := &AttrClass{
		Type:   (AttrType)(c._type),
		native: c,
	}

	return g
}

func (recv *AttrClass) ToC() unsafe.Pointer {
	recv.native._type =
		(C.PangoAttrType)(recv.Type)

	return (unsafe.Pointer)(recv.native)
}

// AttrColor is a wrapper around the C record PangoAttrColor.
type AttrColor struct {
	native *C.PangoAttrColor
	// attr : record
	// color : record
}

func AttrColorNewFromC(u unsafe.Pointer) *AttrColor {
	c := (*C.PangoAttrColor)(u)
	if c == nil {
		return nil
	}

	g := &AttrColor{native: c}

	return g
}

func (recv *AttrColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrFloat is a wrapper around the C record PangoAttrFloat.
type AttrFloat struct {
	native *C.PangoAttrFloat
	// attr : record
	Value float64
}

func AttrFloatNewFromC(u unsafe.Pointer) *AttrFloat {
	c := (*C.PangoAttrFloat)(u)
	if c == nil {
		return nil
	}

	g := &AttrFloat{
		Value:  (float64)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrFloat) ToC() unsafe.Pointer {
	recv.native.value =
		(C.double)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// AttrFontDesc is a wrapper around the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native *C.PangoAttrFontDesc
	// attr : record
	// desc : record
}

func AttrFontDescNewFromC(u unsafe.Pointer) *AttrFontDesc {
	c := (*C.PangoAttrFontDesc)(u)
	if c == nil {
		return nil
	}

	g := &AttrFontDesc{native: c}

	return g
}

func (recv *AttrFontDesc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrInt is a wrapper around the C record PangoAttrInt.
type AttrInt struct {
	native *C.PangoAttrInt
	// attr : record
	Value int32
}

func AttrIntNewFromC(u unsafe.Pointer) *AttrInt {
	c := (*C.PangoAttrInt)(u)
	if c == nil {
		return nil
	}

	g := &AttrInt{
		Value:  (int32)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrInt) ToC() unsafe.Pointer {
	recv.native.value =
		(C.int)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// AttrIterator is a wrapper around the C record PangoAttrIterator.
type AttrIterator struct {
	native *C.PangoAttrIterator
}

func AttrIteratorNewFromC(u unsafe.Pointer) *AttrIterator {
	c := (*C.PangoAttrIterator)(u)
	if c == nil {
		return nil
	}

	g := &AttrIterator{native: c}

	return g
}

func (recv *AttrIterator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Copy is a wrapper around the C function pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	retC := C.pango_attr_iterator_copy((*C.PangoAttrIterator)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	C.pango_attr_iterator_destroy((*C.PangoAttrIterator)(recv.native))

	return
}

// Get is a wrapper around the C function pango_attr_iterator_get.
func (recv *AttrIterator) Get(type_ AttrType) *Attribute {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_iterator_get((*C.PangoAttrIterator)(recv.native), c_type)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_attr_iterator_get_font : unsupported parameter language : record with indirection level of 2

// Next is a wrapper around the C function pango_attr_iterator_next.
func (recv *AttrIterator) Next() bool {
	retC := C.pango_attr_iterator_next((*C.PangoAttrIterator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_attr_iterator_range : unsupported parameter start : no type generator for gint, gint*

// AttrLanguage is a wrapper around the C record PangoAttrLanguage.
type AttrLanguage struct {
	native *C.PangoAttrLanguage
	// attr : record
	// value : record
}

func AttrLanguageNewFromC(u unsafe.Pointer) *AttrLanguage {
	c := (*C.PangoAttrLanguage)(u)
	if c == nil {
		return nil
	}

	g := &AttrLanguage{native: c}

	return g
}

func (recv *AttrLanguage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrList is a wrapper around the C record PangoAttrList.
type AttrList struct {
	native *C.PangoAttrList
}

func AttrListNewFromC(u unsafe.Pointer) *AttrList {
	c := (*C.PangoAttrList)(u)
	if c == nil {
		return nil
	}

	g := &AttrList{native: c}

	return g
}

func (recv *AttrList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrListNew is a wrapper around the C function pango_attr_list_new.
func AttrListNew() *AttrList {
	retC := C.pango_attr_list_new()
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Change is a wrapper around the C function pango_attr_list_change.
func (recv *AttrList) Change(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(attr.ToC())

	C.pango_attr_list_change((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Copy is a wrapper around the C function pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	retC := C.pango_attr_list_copy((*C.PangoAttrList)(recv.native))
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc, PangoAttrFilterFunc

// GetIterator is a wrapper around the C function pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	retC := C.pango_attr_list_get_iterator((*C.PangoAttrList)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function pango_attr_list_insert.
func (recv *AttrList) Insert(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(attr.ToC())

	C.pango_attr_list_insert((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// InsertBefore is a wrapper around the C function pango_attr_list_insert_before.
func (recv *AttrList) InsertBefore(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(attr.ToC())

	C.pango_attr_list_insert_before((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Splice is a wrapper around the C function pango_attr_list_splice.
func (recv *AttrList) Splice(other *AttrList, pos int32, len int32) {
	c_other := (*C.PangoAttrList)(other.ToC())

	c_pos := (C.gint)(pos)

	c_len := (C.gint)(len)

	C.pango_attr_list_splice((*C.PangoAttrList)(recv.native), c_other, c_pos, c_len)

	return
}

// Unref is a wrapper around the C function pango_attr_list_unref.
func (recv *AttrList) Unref() {
	C.pango_attr_list_unref((*C.PangoAttrList)(recv.native))

	return
}

// AttrShape is a wrapper around the C record PangoAttrShape.
type AttrShape struct {
	native *C.PangoAttrShape
	// attr : record
	// ink_rect : record
	// logical_rect : record
	Data uintptr
	// copy_func : no type generator for AttrDataCopyFunc, PangoAttrDataCopyFunc
	// destroy_func : no type generator for GLib.DestroyNotify, GDestroyNotify
}

func AttrShapeNewFromC(u unsafe.Pointer) *AttrShape {
	c := (*C.PangoAttrShape)(u)
	if c == nil {
		return nil
	}

	g := &AttrShape{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *AttrShape) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// AttrSize is a wrapper around the C record PangoAttrSize.
type AttrSize struct {
	native *C.PangoAttrSize
	// attr : record
	Size int32
	// Bitfield not supported :  1 absolute
}

func AttrSizeNewFromC(u unsafe.Pointer) *AttrSize {
	c := (*C.PangoAttrSize)(u)
	if c == nil {
		return nil
	}

	g := &AttrSize{
		Size:   (int32)(c.size),
		native: c,
	}

	return g
}

func (recv *AttrSize) ToC() unsafe.Pointer {
	recv.native.size =
		(C.int)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// AttrString is a wrapper around the C record PangoAttrString.
type AttrString struct {
	native *C.PangoAttrString
	// attr : record
	Value string
}

func AttrStringNewFromC(u unsafe.Pointer) *AttrString {
	c := (*C.PangoAttrString)(u)
	if c == nil {
		return nil
	}

	g := &AttrString{
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *AttrString) ToC() unsafe.Pointer {
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Attribute is a wrapper around the C record PangoAttribute.
type Attribute struct {
	native *C.PangoAttribute
	// klass : record
	StartIndex uint32
	EndIndex   uint32
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	c := (*C.PangoAttribute)(u)
	if c == nil {
		return nil
	}

	g := &Attribute{
		EndIndex:   (uint32)(c.end_index),
		StartIndex: (uint32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.guint)(recv.StartIndex)
	recv.native.end_index =
		(C.guint)(recv.EndIndex)

	return (unsafe.Pointer)(recv.native)
}

// Copy is a wrapper around the C function pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	retC := C.pango_attribute_copy((*C.PangoAttribute)(recv.native))
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	C.pango_attribute_destroy((*C.PangoAttribute)(recv.native))

	return
}

// Equal is a wrapper around the C function pango_attribute_equal.
func (recv *Attribute) Equal(attr2 *Attribute) bool {
	c_attr2 := (*C.PangoAttribute)(attr2.ToC())

	retC := C.pango_attribute_equal((*C.PangoAttribute)(recv.native), c_attr2)
	retGo := retC == C.TRUE

	return retGo
}

// Color is a wrapper around the C record PangoColor.
type Color struct {
	native *C.PangoColor
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	c := (*C.PangoColor)(u)
	if c == nil {
		return nil
	}

	g := &Color{
		Blue:   (uint16)(c.blue),
		Green:  (uint16)(c.green),
		Red:    (uint16)(c.red),
		native: c,
	}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {
	recv.native.red =
		(C.guint16)(recv.Red)
	recv.native.green =
		(C.guint16)(recv.Green)
	recv.native.blue =
		(C.guint16)(recv.Blue)

	return (unsafe.Pointer)(recv.native)
}

// Copy is a wrapper around the C function pango_color_copy.
func (recv *Color) Copy() *Color {
	retC := C.pango_color_copy((*C.PangoColor)(recv.native))
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_color_free.
func (recv *Color) Free() {
	C.pango_color_free((*C.PangoColor)(recv.native))

	return
}

// Parse is a wrapper around the C function pango_color_parse.
func (recv *Color) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.pango_color_parse((*C.PangoColor)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

// ContextClass is a wrapper around the C record PangoContextClass.
type ContextClass struct {
	native *C.PangoContextClass
}

func ContextClassNewFromC(u unsafe.Pointer) *ContextClass {
	c := (*C.PangoContextClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextClass{native: c}

	return g
}

func (recv *ContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Coverage is a wrapper around the C record PangoCoverage.
type Coverage struct {
	native *C.PangoCoverage
}

func CoverageNewFromC(u unsafe.Pointer) *Coverage {
	c := (*C.PangoCoverage)(u)
	if c == nil {
		return nil
	}

	g := &Coverage{native: c}

	return g
}

func (recv *Coverage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Copy is a wrapper around the C function pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	retC := C.pango_coverage_copy((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get is a wrapper around the C function pango_coverage_get.
func (recv *Coverage) Get(index int32) CoverageLevel {
	c_index_ := (C.int)(index)

	retC := C.pango_coverage_get((*C.PangoCoverage)(recv.native), c_index_)
	retGo := (CoverageLevel)(retC)

	return retGo
}

// Max is a wrapper around the C function pango_coverage_max.
func (recv *Coverage) Max(other *Coverage) {
	c_other := (*C.PangoCoverage)(other.ToC())

	C.pango_coverage_max((*C.PangoCoverage)(recv.native), c_other)

	return
}

// Ref is a wrapper around the C function pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	retC := C.pango_coverage_ref((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function pango_coverage_set.
func (recv *Coverage) Set(index int32, level CoverageLevel) {
	c_index_ := (C.int)(index)

	c_level := (C.PangoCoverageLevel)(level)

	C.pango_coverage_set((*C.PangoCoverage)(recv.native), c_index_, c_level)

	return
}

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : no param type

// Unref is a wrapper around the C function pango_coverage_unref.
func (recv *Coverage) Unref() {
	C.pango_coverage_unref((*C.PangoCoverage)(recv.native))

	return
}

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// FontDescription is a wrapper around the C record PangoFontDescription.
type FontDescription struct {
	native *C.PangoFontDescription
}

func FontDescriptionNewFromC(u unsafe.Pointer) *FontDescription {
	c := (*C.PangoFontDescription)(u)
	if c == nil {
		return nil
	}

	g := &FontDescription{native: c}

	return g
}

func (recv *FontDescription) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontDescriptionNew is a wrapper around the C function pango_font_description_new.
func FontDescriptionNew() *FontDescription {
	retC := C.pango_font_description_new()
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BetterMatch is a wrapper around the C function pango_font_description_better_match.
func (recv *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	c_old_match := (*C.PangoFontDescription)(oldMatch.ToC())

	c_new_match := (*C.PangoFontDescription)(newMatch.ToC())

	retC := C.pango_font_description_better_match((*C.PangoFontDescription)(recv.native), c_old_match, c_new_match)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	retC := C.pango_font_description_copy((*C.PangoFontDescription)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CopyStatic is a wrapper around the C function pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	retC := C.pango_font_description_copy_static((*C.PangoFontDescription)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function pango_font_description_equal.
func (recv *FontDescription) Equal(desc2 *FontDescription) bool {
	c_desc2 := (*C.PangoFontDescription)(desc2.ToC())

	retC := C.pango_font_description_equal((*C.PangoFontDescription)(recv.native), c_desc2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function pango_font_description_free.
func (recv *FontDescription) Free() {
	C.pango_font_description_free((*C.PangoFontDescription)(recv.native))

	return
}

// GetFamily is a wrapper around the C function pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	retC := C.pango_font_description_get_family((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSetFields is a wrapper around the C function pango_font_description_get_set_fields.
func (recv *FontDescription) GetSetFields() FontMask {
	retC := C.pango_font_description_get_set_fields((*C.PangoFontDescription)(recv.native))
	retGo := (FontMask)(retC)

	return retGo
}

// GetSize is a wrapper around the C function pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	retC := C.pango_font_description_get_size((*C.PangoFontDescription)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStretch is a wrapper around the C function pango_font_description_get_stretch.
func (recv *FontDescription) GetStretch() Stretch {
	retC := C.pango_font_description_get_stretch((*C.PangoFontDescription)(recv.native))
	retGo := (Stretch)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function pango_font_description_get_style.
func (recv *FontDescription) GetStyle() Style {
	retC := C.pango_font_description_get_style((*C.PangoFontDescription)(recv.native))
	retGo := (Style)(retC)

	return retGo
}

// GetVariant is a wrapper around the C function pango_font_description_get_variant.
func (recv *FontDescription) GetVariant() Variant {
	retC := C.pango_font_description_get_variant((*C.PangoFontDescription)(recv.native))
	retGo := (Variant)(retC)

	return retGo
}

// GetWeight is a wrapper around the C function pango_font_description_get_weight.
func (recv *FontDescription) GetWeight() Weight {
	retC := C.pango_font_description_get_weight((*C.PangoFontDescription)(recv.native))
	retGo := (Weight)(retC)

	return retGo
}

// Hash is a wrapper around the C function pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	retC := C.pango_font_description_hash((*C.PangoFontDescription)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Merge is a wrapper around the C function pango_font_description_merge.
func (recv *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(descToMerge.ToC())

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// MergeStatic is a wrapper around the C function pango_font_description_merge_static.
func (recv *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(descToMerge.ToC())

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge_static((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// SetFamily is a wrapper around the C function pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// SetFamilyStatic is a wrapper around the C function pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family_static((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// SetSize is a wrapper around the C function pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	c_size := (C.gint)(size)

	C.pango_font_description_set_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// SetStretch is a wrapper around the C function pango_font_description_set_stretch.
func (recv *FontDescription) SetStretch(stretch Stretch) {
	c_stretch := (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch((*C.PangoFontDescription)(recv.native), c_stretch)

	return
}

// SetStyle is a wrapper around the C function pango_font_description_set_style.
func (recv *FontDescription) SetStyle(style Style) {
	c_style := (C.PangoStyle)(style)

	C.pango_font_description_set_style((*C.PangoFontDescription)(recv.native), c_style)

	return
}

// SetVariant is a wrapper around the C function pango_font_description_set_variant.
func (recv *FontDescription) SetVariant(variant Variant) {
	c_variant := (C.PangoVariant)(variant)

	C.pango_font_description_set_variant((*C.PangoFontDescription)(recv.native), c_variant)

	return
}

// SetWeight is a wrapper around the C function pango_font_description_set_weight.
func (recv *FontDescription) SetWeight(weight Weight) {
	c_weight := (C.PangoWeight)(weight)

	C.pango_font_description_set_weight((*C.PangoFontDescription)(recv.native), c_weight)

	return
}

// ToFilename is a wrapper around the C function pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	retC := C.pango_font_description_to_filename((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	retC := C.pango_font_description_to_string((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnsetFields is a wrapper around the C function pango_font_description_unset_fields.
func (recv *FontDescription) UnsetFields(toUnset FontMask) {
	c_to_unset := (C.PangoFontMask)(toUnset)

	C.pango_font_description_unset_fields((*C.PangoFontDescription)(recv.native), c_to_unset)

	return
}

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// Blacklisted : PangoFontMetrics

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// GlyphGeometry is a wrapper around the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native  *C.PangoGlyphGeometry
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

func GlyphGeometryNewFromC(u unsafe.Pointer) *GlyphGeometry {
	c := (*C.PangoGlyphGeometry)(u)
	if c == nil {
		return nil
	}

	g := &GlyphGeometry{
		Width:   (GlyphUnit)(c.width),
		XOffset: (GlyphUnit)(c.x_offset),
		YOffset: (GlyphUnit)(c.y_offset),
		native:  c,
	}

	return g
}

func (recv *GlyphGeometry) ToC() unsafe.Pointer {
	recv.native.width =
		(C.PangoGlyphUnit)(recv.Width)
	recv.native.x_offset =
		(C.PangoGlyphUnit)(recv.XOffset)
	recv.native.y_offset =
		(C.PangoGlyphUnit)(recv.YOffset)

	return (unsafe.Pointer)(recv.native)
}

// GlyphInfo is a wrapper around the C record PangoGlyphInfo.
type GlyphInfo struct {
	native *C.PangoGlyphInfo
	Glyph  Glyph
	// geometry : record
	// attr : record
}

func GlyphInfoNewFromC(u unsafe.Pointer) *GlyphInfo {
	c := (*C.PangoGlyphInfo)(u)
	if c == nil {
		return nil
	}

	g := &GlyphInfo{
		Glyph:  (Glyph)(c.glyph),
		native: c,
	}

	return g
}

func (recv *GlyphInfo) ToC() unsafe.Pointer {
	recv.native.glyph =
		(C.PangoGlyph)(recv.Glyph)

	return (unsafe.Pointer)(recv.native)
}

// GlyphItem is a wrapper around the C record PangoGlyphItem.
type GlyphItem struct {
	native *C.PangoGlyphItem
	// item : record
	// glyphs : record
}

func GlyphItemNewFromC(u unsafe.Pointer) *GlyphItem {
	c := (*C.PangoGlyphItem)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItem{native: c}

	return g
}

func (recv *GlyphItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_glyph_item_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs : no param type

// GlyphString is a wrapper around the C record PangoGlyphString.
type GlyphString struct {
	native    *C.PangoGlyphString
	NumGlyphs int32
	// no type for glyphs
	// log_clusters : no type generator for gint, gint*
	// Private : space
}

func GlyphStringNewFromC(u unsafe.Pointer) *GlyphString {
	c := (*C.PangoGlyphString)(u)
	if c == nil {
		return nil
	}

	g := &GlyphString{
		NumGlyphs: (int32)(c.num_glyphs),
		native:    c,
	}

	return g
}

func (recv *GlyphString) ToC() unsafe.Pointer {
	recv.native.num_glyphs =
		(C.gint)(recv.NumGlyphs)

	return (unsafe.Pointer)(recv.native)
}

// GlyphStringNew is a wrapper around the C function pango_glyph_string_new.
func GlyphStringNew() *GlyphString {
	retC := C.pango_glyph_string_new()
	retGo := GlyphStringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	retC := C.pango_glyph_string_copy((*C.PangoGlyphString)(recv.native))
	retGo := GlyphStringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Extents is a wrapper around the C function pango_glyph_string_extents.
func (recv *GlyphString) Extents(font *Font) (*Rectangle, *Rectangle) {
	c_font := (*C.PangoFont)(font.ToC())

	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_glyph_string_extents((*C.PangoGlyphString)(recv.native), c_font, &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// ExtentsRange is a wrapper around the C function pango_glyph_string_extents_range.
func (recv *GlyphString) ExtentsRange(start int32, end int32, font *Font) (*Rectangle, *Rectangle) {
	c_start := (C.int)(start)

	c_end := (C.int)(end)

	c_font := (*C.PangoFont)(font.ToC())

	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_glyph_string_extents_range((*C.PangoGlyphString)(recv.native), c_start, c_end, c_font, &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Free is a wrapper around the C function pango_glyph_string_free.
func (recv *GlyphString) Free() {
	C.pango_glyph_string_free((*C.PangoGlyphString)(recv.native))

	return
}

// Unsupported : pango_glyph_string_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_string_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// SetSize is a wrapper around the C function pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	c_new_len := (C.gint)(newLen)

	C.pango_glyph_string_set_size((*C.PangoGlyphString)(recv.native), c_new_len)

	return
}

// Unsupported : pango_glyph_string_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// GlyphVisAttr is a wrapper around the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native *C.PangoGlyphVisAttr
	// Bitfield not supported :  1 is_cluster_start
}

func GlyphVisAttrNewFromC(u unsafe.Pointer) *GlyphVisAttr {
	c := (*C.PangoGlyphVisAttr)(u)
	if c == nil {
		return nil
	}

	g := &GlyphVisAttr{native: c}

	return g
}

func (recv *GlyphVisAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoIncludedModule

// Item is a wrapper around the C record PangoItem.
type Item struct {
	native   *C.PangoItem
	Offset   int32
	Length   int32
	NumChars int32
	// analysis : record
}

func ItemNewFromC(u unsafe.Pointer) *Item {
	c := (*C.PangoItem)(u)
	if c == nil {
		return nil
	}

	g := &Item{
		Length:   (int32)(c.length),
		NumChars: (int32)(c.num_chars),
		Offset:   (int32)(c.offset),
		native:   c,
	}

	return g
}

func (recv *Item) ToC() unsafe.Pointer {
	recv.native.offset =
		(C.gint)(recv.Offset)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.num_chars =
		(C.gint)(recv.NumChars)

	return (unsafe.Pointer)(recv.native)
}

// ItemNew is a wrapper around the C function pango_item_new.
func ItemNew() *Item {
	retC := C.pango_item_new()
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function pango_item_copy.
func (recv *Item) Copy() *Item {
	retC := C.pango_item_copy((*C.PangoItem)(recv.native))
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_item_free.
func (recv *Item) Free() {
	C.pango_item_free((*C.PangoItem)(recv.native))

	return
}

// Split is a wrapper around the C function pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	c_split_index := (C.int)(splitIndex)

	c_split_offset := (C.int)(splitOffset)

	retC := C.pango_item_split((*C.PangoItem)(recv.native), c_split_index, c_split_offset)
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Language is a wrapper around the C record PangoLanguage.
type Language struct {
	native *C.PangoLanguage
}

func LanguageNewFromC(u unsafe.Pointer) *Language {
	c := (*C.PangoLanguage)(u)
	if c == nil {
		return nil
	}

	g := &Language{native: c}

	return g
}

func (recv *Language) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetSampleString is a wrapper around the C function pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	retC := C.pango_language_get_sample_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_language_get_scripts : unsupported parameter num_scripts : no type generator for gint, int*

// Matches is a wrapper around the C function pango_language_matches.
func (recv *Language) Matches(rangeList string) bool {
	c_range_list := C.CString(rangeList)
	defer C.free(unsafe.Pointer(c_range_list))

	retC := C.pango_language_matches((*C.PangoLanguage)(recv.native), c_range_list)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function pango_language_to_string.
func (recv *Language) ToString() string {
	retC := C.pango_language_to_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LayoutClass is a wrapper around the C record PangoLayoutClass.
type LayoutClass struct {
	native *C.PangoLayoutClass
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	c := (*C.PangoLayoutClass)(u)
	if c == nil {
		return nil
	}

	g := &LayoutClass{native: c}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoLayoutIter

// LayoutLine is a wrapper around the C record PangoLayoutLine.
type LayoutLine struct {
	native *C.PangoLayoutLine
	// layout : record
	StartIndex int32
	Length     int32
	// runs : record
	// Bitfield not supported :  1 is_paragraph_start
	// Bitfield not supported :  3 resolved_dir
}

func LayoutLineNewFromC(u unsafe.Pointer) *LayoutLine {
	c := (*C.PangoLayoutLine)(u)
	if c == nil {
		return nil
	}

	g := &LayoutLine{
		Length:     (int32)(c.length),
		StartIndex: (int32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *LayoutLine) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.gint)(recv.StartIndex)
	recv.native.length =
		(C.gint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// GetExtents is a wrapper around the C function pango_layout_line_get_extents.
func (recv *LayoutLine) GetExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// GetPixelExtents is a wrapper around the C function pango_layout_line_get_pixel_extents.
func (recv *LayoutLine) GetPixelExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_pixel_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : no param type

// Unsupported : pango_layout_line_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unref is a wrapper around the C function pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	C.pango_layout_line_unref((*C.PangoLayoutLine)(recv.native))

	return
}

// Unsupported : pango_layout_line_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// LogAttr is a wrapper around the C record PangoLogAttr.
type LogAttr struct {
	native *C.PangoLogAttr
	// Bitfield not supported :  1 is_line_break
	// Bitfield not supported :  1 is_mandatory_break
	// Bitfield not supported :  1 is_char_break
	// Bitfield not supported :  1 is_white
	// Bitfield not supported :  1 is_cursor_position
	// Bitfield not supported :  1 is_word_start
	// Bitfield not supported :  1 is_word_end
	// Bitfield not supported :  1 is_sentence_boundary
	// Bitfield not supported :  1 is_sentence_start
	// Bitfield not supported :  1 is_sentence_end
	// Bitfield not supported :  1 backspace_deletes_character
	// Bitfield not supported :  1 is_expandable_space
	// Bitfield not supported :  1 is_word_boundary
}

func LogAttrNewFromC(u unsafe.Pointer) *LogAttr {
	c := (*C.PangoLogAttr)(u)
	if c == nil {
		return nil
	}

	g := &LogAttr{native: c}

	return g
}

func (recv *LogAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Rectangle is a wrapper around the C record PangoRectangle.
type Rectangle struct {
	native *C.PangoRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.PangoRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.int)(recv.X)
	recv.native.y =
		(C.int)(recv.Y)
	recv.native.width =
		(C.int)(recv.Width)
	recv.native.height =
		(C.int)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// RendererPrivate is a wrapper around the C record PangoRendererPrivate.
type RendererPrivate struct {
	native *C.PangoRendererPrivate
}

func RendererPrivateNewFromC(u unsafe.Pointer) *RendererPrivate {
	c := (*C.PangoRendererPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RendererPrivate{native: c}

	return g
}

func (recv *RendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoScriptForLang

// ScriptIter is a wrapper around the C record PangoScriptIter.
type ScriptIter struct {
	native *C.PangoScriptIter
}

func ScriptIterNewFromC(u unsafe.Pointer) *ScriptIter {
	c := (*C.PangoScriptIter)(u)
	if c == nil {
		return nil
	}

	g := &ScriptIter{native: c}

	return g
}

func (recv *ScriptIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// TabArray is a wrapper around the C record PangoTabArray.
type TabArray struct {
	native *C.PangoTabArray
}

func TabArrayNewFromC(u unsafe.Pointer) *TabArray {
	c := (*C.PangoTabArray)(u)
	if c == nil {
		return nil
	}

	g := &TabArray{native: c}

	return g
}

func (recv *TabArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TabArrayNew is a wrapper around the C function pango_tab_array_new.
func TabArrayNew(initialSize int32, positionsInPixels bool) *TabArray {
	c_initial_size := (C.gint)(initialSize)

	c_positions_in_pixels :=
		boolToGboolean(positionsInPixels)

	retC := C.pango_tab_array_new(c_initial_size, c_positions_in_pixels)
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Copy is a wrapper around the C function pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	retC := C.pango_tab_array_copy((*C.PangoTabArray)(recv.native))
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_tab_array_free.
func (recv *TabArray) Free() {
	C.pango_tab_array_free((*C.PangoTabArray)(recv.native))

	return
}

// GetPositionsInPixels is a wrapper around the C function pango_tab_array_get_positions_in_pixels.
func (recv *TabArray) GetPositionsInPixels() bool {
	retC := C.pango_tab_array_get_positions_in_pixels((*C.PangoTabArray)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSize is a wrapper around the C function pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	retC := C.pango_tab_array_get_size((*C.PangoTabArray)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2

// Resize is a wrapper around the C function pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	c_new_size := (C.gint)(newSize)

	C.pango_tab_array_resize((*C.PangoTabArray)(recv.native), c_new_size)

	return
}

// SetTab is a wrapper around the C function pango_tab_array_set_tab.
func (recv *TabArray) SetTab(tabIndex int32, alignment TabAlign, location int32) {
	c_tab_index := (C.gint)(tabIndex)

	c_alignment := (C.PangoTabAlign)(alignment)

	c_location := (C.gint)(location)

	C.pango_tab_array_set_tab((*C.PangoTabArray)(recv.native), c_tab_index, c_alignment, c_location)

	return
}
