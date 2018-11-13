// This is a generated file - DO NOT EDIT

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// The #PangoAnalysis structure stores information about
// the properties of a segment of text.
/*

C type

PangoAnalysis
*/
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

// The #PangoAttrClass structure stores the type and operations for
// a particular type of attribute. The functions in this structure should
// not be called directly. Instead, one should use the wrapper functions
// provided for #PangoAttribute.
/*

C type

PangoAttrClass
*/
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

// The #PangoAttrColor structure is used to represent attributes that
// are colors.
/*

C type

PangoAttrColor
*/
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

// The #PangoAttrFloat structure is used to represent attributes with
// a float or double value.
/*

C type

PangoAttrFloat
*/
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

// The #PangoAttrFontDesc structure is used to store an attribute that
// sets all aspects of the font description at once.
/*

C type

PangoAttrFontDesc
*/
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

// The #PangoAttrInt structure is used to represent attributes with
// an integer or enumeration value.
/*

C type

PangoAttrInt
*/
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

// The #PangoAttrIterator structure is used to represent an
// iterator through a #PangoAttrList. A new iterator is created
// with pango_attr_list_get_iterator(). Once the iterator
// is created, it can be advanced through the style changes
// in the text using pango_attr_iterator_next(). At each
// style change, the range of the current style segment and the
// attributes currently in effect can be queried.
/*

C type

PangoAttrIterator
*/
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

// Copy a #PangoAttrIterator
/*

C function

pango_attr_iterator_copy
*/
func (recv *AttrIterator) Copy() *AttrIterator {
	retC := C.pango_attr_iterator_copy((*C.PangoAttrIterator)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy a #PangoAttrIterator and free all associated memory.
/*

C function

pango_attr_iterator_destroy
*/
func (recv *AttrIterator) Destroy() {
	C.pango_attr_iterator_destroy((*C.PangoAttrIterator)(recv.native))

	return
}

// Find the current attribute of a particular type at the iterator
// location. When multiple attributes of the same type overlap,
// the attribute whose range starts closest to the current location
// is used.
/*

C function

pango_attr_iterator_get
*/
func (recv *AttrIterator) Get(type_ AttrType) *Attribute {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_iterator_get((*C.PangoAttrIterator)(recv.native), c_type)
	var retGo (*Attribute)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AttributeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Get the font and other attributes at the current iterator position.
/*

C function

pango_attr_iterator_get_font
*/
func (recv *AttrIterator) GetFont(desc *FontDescription, language *Language, extraAttrs *glib.SList) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (**C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (**C.PangoLanguage)(language.ToC())
	}

	c_extra_attrs := (**C.GSList)(C.NULL)
	if extraAttrs != nil {
		c_extra_attrs = (**C.GSList)(extraAttrs.ToC())
	}

	C.pango_attr_iterator_get_font((*C.PangoAttrIterator)(recv.native), c_desc, c_language, c_extra_attrs)

	return
}

// Advance the iterator until the next change of style.
/*

C function

pango_attr_iterator_next
*/
func (recv *AttrIterator) Next() bool {
	retC := C.pango_attr_iterator_next((*C.PangoAttrIterator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Get the range of the current segment. Note that the
// stored return values are signed, not unsigned like
// the values in #PangoAttribute. To deal with this API
// oversight, stored return values that wouldn't fit into
// a signed integer are clamped to %G_MAXINT.
/*

C function

pango_attr_iterator_range
*/
func (recv *AttrIterator) Range() (int32, int32) {
	var c_start C.gint

	var c_end C.gint

	C.pango_attr_iterator_range((*C.PangoAttrIterator)(recv.native), &c_start, &c_end)

	start := (int32)(c_start)

	end := (int32)(c_end)

	return start, end
}

// The #PangoAttrLanguage structure is used to represent attributes that
// are languages.
/*

C type

PangoAttrLanguage
*/
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

// The #PangoAttrList structure represents a list of attributes
// that apply to a section of text. The attributes are, in general,
// allowed to overlap in an arbitrary fashion, however, if the
// attributes are manipulated only through pango_attr_list_change(),
// the overlap between properties will meet stricter criteria.
//
// Since the #PangoAttrList structure is stored as a linear list,
// it is not suitable for storing attributes for large amounts
// of text. In general, you should not use a single #PangoAttrList
// for more than one paragraph of text.
/*

C type

PangoAttrList
*/
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

// Create a new empty attribute list with a reference count of one.
/*

C function

pango_attr_list_new
*/
func AttrListNew() *AttrList {
	retC := C.pango_attr_list_new()
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert the given attribute into the #PangoAttrList. It will
// replace any attributes of the same type on that segment
// and be merged with any adjoining attributes that are identical.
//
// This function is slower than pango_attr_list_insert() for
// creating a attribute list in order (potentially much slower
// for large lists). However, pango_attr_list_insert() is not
// suitable for continually changing a set of attributes
// since it never removes or combines existing attributes.
/*

C function

pango_attr_list_change
*/
func (recv *AttrList) Change(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_change((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Copy @list and return an identical new list.
/*

C function

pango_attr_list_copy
*/
func (recv *AttrList) Copy() *AttrList {
	retC := C.pango_attr_list_copy((*C.PangoAttrList)(recv.native))
	var retGo (*AttrList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AttrListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Create a iterator initialized to the beginning of the list.
// @list must not be modified until this iterator is freed.
/*

C function

pango_attr_list_get_iterator
*/
func (recv *AttrList) GetIterator() *AttrIterator {
	retC := C.pango_attr_list_get_iterator((*C.PangoAttrList)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert the given attribute into the #PangoAttrList. It will
// be inserted after all other attributes with a matching
// @start_index.
/*

C function

pango_attr_list_insert
*/
func (recv *AttrList) Insert(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_insert((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Insert the given attribute into the #PangoAttrList. It will
// be inserted before all other attributes with a matching
// @start_index.
/*

C function

pango_attr_list_insert_before
*/
func (recv *AttrList) InsertBefore(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_insert_before((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// This function opens up a hole in @list, fills it in with attributes from
// the left, and then merges @other on top of the hole.
//
// This operation is equivalent to stretching every attribute
// that applies at position @pos in @list by an amount @len,
// and then calling pango_attr_list_change() with a copy
// of each attribute in @other in sequence (offset in position by @pos).
//
// This operation proves useful for, for instance, inserting
// a pre-edit string in the middle of an edit buffer.
/*

C function

pango_attr_list_splice
*/
func (recv *AttrList) Splice(other *AttrList, pos int32, len int32) {
	c_other := (*C.PangoAttrList)(C.NULL)
	if other != nil {
		c_other = (*C.PangoAttrList)(other.ToC())
	}

	c_pos := (C.gint)(pos)

	c_len := (C.gint)(len)

	C.pango_attr_list_splice((*C.PangoAttrList)(recv.native), c_other, c_pos, c_len)

	return
}

// Decrease the reference count of the given attribute list by one.
// If the result is zero, free the attribute list and the attributes
// it contains.
/*

C function

pango_attr_list_unref
*/
func (recv *AttrList) Unref() {
	C.pango_attr_list_unref((*C.PangoAttrList)(recv.native))

	return
}

// The #PangoAttrShape structure is used to represent attributes which
// impose shape restrictions.
/*

C type

PangoAttrShape
*/
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

// The #PangoAttrSize structure is used to represent attributes which
// set font size.
/*

C type

PangoAttrSize
*/
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

// The #PangoAttrString structure is used to represent attributes with
// a string value.
/*

C type

PangoAttrString
*/
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

// The #PangoAttribute structure represents the common portions of all
// attributes. Particular types of attributes include this structure
// as their initial portion. The common portion of the attribute holds
// the range to which the value in the type-specific part of the attribute
// applies and should be initialized using pango_attribute_init().
// By default an attribute will have an all-inclusive range of [0,%G_MAXUINT].
/*

C type

PangoAttribute
*/
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

// Make a copy of an attribute.
/*

C function

pango_attribute_copy
*/
func (recv *Attribute) Copy() *Attribute {
	retC := C.pango_attribute_copy((*C.PangoAttribute)(recv.native))
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy a #PangoAttribute and free all associated memory.
/*

C function

pango_attribute_destroy
*/
func (recv *Attribute) Destroy() {
	C.pango_attribute_destroy((*C.PangoAttribute)(recv.native))

	return
}

// Compare two attributes for equality. This compares only the
// actual value of the two attributes and not the ranges that the
// attributes apply to.
/*

C function

pango_attribute_equal
*/
func (recv *Attribute) Equal(attr2 *Attribute) bool {
	c_attr2 := (*C.PangoAttribute)(C.NULL)
	if attr2 != nil {
		c_attr2 = (*C.PangoAttribute)(attr2.ToC())
	}

	retC := C.pango_attribute_equal((*C.PangoAttribute)(recv.native), c_attr2)
	retGo := retC == C.TRUE

	return retGo
}

// The #PangoColor structure is used to
// represent a color in an uncalibrated RGB color-space.
/*

C type

PangoColor
*/
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

// Creates a copy of @src, which should be freed with
// pango_color_free(). Primarily used by language bindings,
// not that useful otherwise (since colors can just be copied
// by assignment in C).
/*

C function

pango_color_copy
*/
func (recv *Color) Copy() *Color {
	retC := C.pango_color_copy((*C.PangoColor)(recv.native))
	var retGo (*Color)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ColorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Frees a color allocated by pango_color_copy().
/*

C function

pango_color_free
*/
func (recv *Color) Free() {
	C.pango_color_free((*C.PangoColor)(recv.native))

	return
}

// Fill in the fields of a color from a string specification. The
// string can either one of a large set of standard names. (Taken
// from the CSS <ulink url="http://dev.w3.org/csswg/css-color/#named-colors">specification</ulink>), or it can be a hexadecimal
// value in the
// form '&num;rgb' '&num;rrggbb' '&num;rrrgggbbb' or '&num;rrrrggggbbbb' where
// 'r', 'g' and 'b' are hex digits of the red, green, and blue
// components of the color, respectively. (White in the four
// forms is '&num;fff' '&num;ffffff' '&num;fffffffff' and '&num;ffffffffffff')
/*

C function

pango_color_parse
*/
func (recv *Color) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.pango_color_parse((*C.PangoColor)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

/*

C type

PangoContextClass
*/
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

// The #PangoCoverage structure represents a map from Unicode characters
// to #PangoCoverageLevel. It is an opaque structure with no public fields.
/*

C type

PangoCoverage
*/
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

// Copy an existing #PangoCoverage. (This function may now be unnecessary
// since we refcount the structure. File a bug if you use it.)
/*

C function

pango_coverage_copy
*/
func (recv *Coverage) Copy() *Coverage {
	retC := C.pango_coverage_copy((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determine whether a particular index is covered by @coverage
/*

C function

pango_coverage_get
*/
func (recv *Coverage) Get(index int32) CoverageLevel {
	c_index_ := (C.int)(index)

	retC := C.pango_coverage_get((*C.PangoCoverage)(recv.native), c_index_)
	retGo := (CoverageLevel)(retC)

	return retGo
}

// Set the coverage for each index in @coverage to be the max (better)
// value of the current coverage for the index and the coverage for
// the corresponding index in @other.
/*

C function

pango_coverage_max
*/
func (recv *Coverage) Max(other *Coverage) {
	c_other := (*C.PangoCoverage)(C.NULL)
	if other != nil {
		c_other = (*C.PangoCoverage)(other.ToC())
	}

	C.pango_coverage_max((*C.PangoCoverage)(recv.native), c_other)

	return
}

// Increase the reference count on the #PangoCoverage by one
/*

C function

pango_coverage_ref
*/
func (recv *Coverage) Ref() *Coverage {
	retC := C.pango_coverage_ref((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Modify a particular index within @coverage
/*

C function

pango_coverage_set
*/
func (recv *Coverage) Set(index int32, level CoverageLevel) {
	c_index_ := (C.int)(index)

	c_level := (C.PangoCoverageLevel)(level)

	C.pango_coverage_set((*C.PangoCoverage)(recv.native), c_index_, c_level)

	return
}

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : output array param bytes

// Decrease the reference count on the #PangoCoverage by one.
// If the result is zero, free the coverage and all associated memory.
/*

C function

pango_coverage_unref
*/
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

// The #PangoFontDescription structure represents the description
// of an ideal font. These structures are used both to list
// what fonts are available on the system and also for specifying
// the characteristics of a font to load.
/*

C type

PangoFontDescription
*/
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

// Creates a new font description structure with all fields unset.
/*

C function

pango_font_description_new
*/
func FontDescriptionNew() *FontDescription {
	retC := C.pango_font_description_new()
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines if the style attributes of @new_match are a closer match
// for @desc than those of @old_match are, or if @old_match is %NULL,
// determines if @new_match is a match at all.
// Approximate matching is done for
// weight and style; other style attributes must match exactly.
// Style attributes are all attributes other than family and size-related
// attributes.  Approximate matching for style considers PANGO_STYLE_OBLIQUE
// and PANGO_STYLE_ITALIC as matches, but not as good a match as when the
// styles are equal.
//
// Note that @old_match must match @desc.
/*

C function

pango_font_description_better_match
*/
func (recv *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	c_old_match := (*C.PangoFontDescription)(C.NULL)
	if oldMatch != nil {
		c_old_match = (*C.PangoFontDescription)(oldMatch.ToC())
	}

	c_new_match := (*C.PangoFontDescription)(C.NULL)
	if newMatch != nil {
		c_new_match = (*C.PangoFontDescription)(newMatch.ToC())
	}

	retC := C.pango_font_description_better_match((*C.PangoFontDescription)(recv.native), c_old_match, c_new_match)
	retGo := retC == C.TRUE

	return retGo
}

// Make a copy of a #PangoFontDescription.
/*

C function

pango_font_description_copy
*/
func (recv *FontDescription) Copy() *FontDescription {
	retC := C.pango_font_description_copy((*C.PangoFontDescription)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Like pango_font_description_copy(), but only a shallow copy is made
// of the family name and other allocated fields. The result can only
// be used until @desc is modified or freed. This is meant to be used
// when the copy is only needed temporarily.
/*

C function

pango_font_description_copy_static
*/
func (recv *FontDescription) CopyStatic() *FontDescription {
	retC := C.pango_font_description_copy_static((*C.PangoFontDescription)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Compares two font descriptions for equality. Two font descriptions
// are considered equal if the fonts they describe are provably identical.
// This means that their masks do not have to match, as long as other fields
// are all the same. (Two font descriptions may result in identical fonts
// being loaded, but still compare %FALSE.)
/*

C function

pango_font_description_equal
*/
func (recv *FontDescription) Equal(desc2 *FontDescription) bool {
	c_desc2 := (*C.PangoFontDescription)(C.NULL)
	if desc2 != nil {
		c_desc2 = (*C.PangoFontDescription)(desc2.ToC())
	}

	retC := C.pango_font_description_equal((*C.PangoFontDescription)(recv.native), c_desc2)
	retGo := retC == C.TRUE

	return retGo
}

// Frees a font description.
/*

C function

pango_font_description_free
*/
func (recv *FontDescription) Free() {
	C.pango_font_description_free((*C.PangoFontDescription)(recv.native))

	return
}

// Gets the family name field of a font description. See
// pango_font_description_set_family().
/*

C function

pango_font_description_get_family
*/
func (recv *FontDescription) GetFamily() string {
	retC := C.pango_font_description_get_family((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Determines which fields in a font description have been set.
/*

C function

pango_font_description_get_set_fields
*/
func (recv *FontDescription) GetSetFields() FontMask {
	retC := C.pango_font_description_get_set_fields((*C.PangoFontDescription)(recv.native))
	retGo := (FontMask)(retC)

	return retGo
}

// Gets the size field of a font description.
// See pango_font_description_set_size().
/*

C function

pango_font_description_get_size
*/
func (recv *FontDescription) GetSize() int32 {
	retC := C.pango_font_description_get_size((*C.PangoFontDescription)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the stretch field of a font description.
// See pango_font_description_set_stretch().
/*

C function

pango_font_description_get_stretch
*/
func (recv *FontDescription) GetStretch() Stretch {
	retC := C.pango_font_description_get_stretch((*C.PangoFontDescription)(recv.native))
	retGo := (Stretch)(retC)

	return retGo
}

// Gets the style field of a #PangoFontDescription. See
// pango_font_description_set_style().
/*

C function

pango_font_description_get_style
*/
func (recv *FontDescription) GetStyle() Style {
	retC := C.pango_font_description_get_style((*C.PangoFontDescription)(recv.native))
	retGo := (Style)(retC)

	return retGo
}

// Gets the variant field of a #PangoFontDescription. See
// pango_font_description_set_variant().
/*

C function

pango_font_description_get_variant
*/
func (recv *FontDescription) GetVariant() Variant {
	retC := C.pango_font_description_get_variant((*C.PangoFontDescription)(recv.native))
	retGo := (Variant)(retC)

	return retGo
}

// Gets the weight field of a font description. See
// pango_font_description_set_weight().
/*

C function

pango_font_description_get_weight
*/
func (recv *FontDescription) GetWeight() Weight {
	retC := C.pango_font_description_get_weight((*C.PangoFontDescription)(recv.native))
	retGo := (Weight)(retC)

	return retGo
}

// Computes a hash of a #PangoFontDescription structure suitable
// to be used, for example, as an argument to g_hash_table_new().
// The hash value is independent of @desc->mask.
/*

C function

pango_font_description_hash
*/
func (recv *FontDescription) Hash() uint32 {
	retC := C.pango_font_description_hash((*C.PangoFontDescription)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Merges the fields that are set in @desc_to_merge into the fields in
// @desc.  If @replace_existing is %FALSE, only fields in @desc that
// are not already set are affected. If %TRUE, then fields that are
// already set will be replaced as well.
//
// If @desc_to_merge is %NULL, this function performs nothing.
/*

C function

pango_font_description_merge
*/
func (recv *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(C.NULL)
	if descToMerge != nil {
		c_desc_to_merge = (*C.PangoFontDescription)(descToMerge.ToC())
	}

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// Like pango_font_description_merge(), but only a shallow copy is made
// of the family name and other allocated fields. @desc can only be
// used until @desc_to_merge is modified or freed. This is meant
// to be used when the merged font description is only needed temporarily.
/*

C function

pango_font_description_merge_static
*/
func (recv *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(C.NULL)
	if descToMerge != nil {
		c_desc_to_merge = (*C.PangoFontDescription)(descToMerge.ToC())
	}

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge_static((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// Sets the family name field of a font description. The family
// name represents a family of related font styles, and will
// resolve to a particular #PangoFontFamily. In some uses of
// #PangoFontDescription, it is also possible to use a comma
// separated list of family names for this field.
/*

C function

pango_font_description_set_family
*/
func (recv *FontDescription) SetFamily(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// Like pango_font_description_set_family(), except that no
// copy of @family is made. The caller must make sure that the
// string passed in stays around until @desc has been freed
// or the name is set again. This function can be used if
// @family is a static string such as a C string literal, or
// if @desc is only needed temporarily.
/*

C function

pango_font_description_set_family_static
*/
func (recv *FontDescription) SetFamilyStatic(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family_static((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// Sets the size field of a font description in fractional points. This is mutually
// exclusive with pango_font_description_set_absolute_size().
/*

C function

pango_font_description_set_size
*/
func (recv *FontDescription) SetSize(size int32) {
	c_size := (C.gint)(size)

	C.pango_font_description_set_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// Sets the stretch field of a font description. The stretch field
// specifies how narrow or wide the font should be.
/*

C function

pango_font_description_set_stretch
*/
func (recv *FontDescription) SetStretch(stretch Stretch) {
	c_stretch := (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch((*C.PangoFontDescription)(recv.native), c_stretch)

	return
}

// Sets the style field of a #PangoFontDescription. The
// #PangoStyle enumeration describes whether the font is slanted and
// the manner in which it is slanted; it can be either
// #PANGO_STYLE_NORMAL, #PANGO_STYLE_ITALIC, or #PANGO_STYLE_OBLIQUE.
// Most fonts will either have a italic style or an oblique
// style, but not both, and font matching in Pango will
// match italic specifications with oblique fonts and vice-versa
// if an exact match is not found.
/*

C function

pango_font_description_set_style
*/
func (recv *FontDescription) SetStyle(style Style) {
	c_style := (C.PangoStyle)(style)

	C.pango_font_description_set_style((*C.PangoFontDescription)(recv.native), c_style)

	return
}

// Sets the variant field of a font description. The #PangoVariant
// can either be %PANGO_VARIANT_NORMAL or %PANGO_VARIANT_SMALL_CAPS.
/*

C function

pango_font_description_set_variant
*/
func (recv *FontDescription) SetVariant(variant Variant) {
	c_variant := (C.PangoVariant)(variant)

	C.pango_font_description_set_variant((*C.PangoFontDescription)(recv.native), c_variant)

	return
}

// Sets the weight field of a font description. The weight field
// specifies how bold or light the font should be. In addition
// to the values of the #PangoWeight enumeration, other intermediate
// numeric values are possible.
/*

C function

pango_font_description_set_weight
*/
func (recv *FontDescription) SetWeight(weight Weight) {
	c_weight := (C.PangoWeight)(weight)

	C.pango_font_description_set_weight((*C.PangoFontDescription)(recv.native), c_weight)

	return
}

// Creates a filename representation of a font description. The
// filename is identical to the result from calling
// pango_font_description_to_string(), but with underscores instead of
// characters that are untypical in filenames, and in lower case only.
/*

C function

pango_font_description_to_filename
*/
func (recv *FontDescription) ToFilename() string {
	retC := C.pango_font_description_to_filename((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a string representation of a font description. See
// pango_font_description_from_string() for a description of the
// format of the string representation. The family list in the
// string description will only have a terminating comma if the
// last word of the list is a valid style option.
/*

C function

pango_font_description_to_string
*/
func (recv *FontDescription) ToString() string {
	retC := C.pango_font_description_to_string((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsets some of the fields in a #PangoFontDescription.  The unset
// fields will get back to their default values.
/*

C function

pango_font_description_unset_fields
*/
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

// The #PangoGlyphGeometry structure contains width and positioning
// information for a single glyph.
/*

C type

PangoGlyphGeometry
*/
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

// The #PangoGlyphInfo structure represents a single glyph together with
// positioning information and visual attributes.
// It contains the following fields.
/*

C type

PangoGlyphInfo
*/
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

// A #PangoGlyphItem is a pair of a #PangoItem and the glyphs
// resulting from shaping the text corresponding to an item.
// As an example of the usage of #PangoGlyphItem, the results
// of shaping text with #PangoLayout is a list of #PangoLayoutLine,
// each of which contains a list of #PangoGlyphItem.
/*

C type

PangoGlyphItem
*/
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

// Blacklisted : PangoGlyphString

// The PangoGlyphVisAttr is used to communicate information between
// the shaping phase and the rendering phase.  More attributes may be
// added in the future.
/*

C type

PangoGlyphVisAttr
*/
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

// The #PangoItem structure stores information about a segment of text.
/*

C type

PangoItem
*/
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

// Creates a new #PangoItem structure initialized to default values.
/*

C function

pango_item_new
*/
func ItemNew() *Item {
	retC := C.pango_item_new()
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy an existing #PangoItem structure.
/*

C function

pango_item_copy
*/
func (recv *Item) Copy() *Item {
	retC := C.pango_item_copy((*C.PangoItem)(recv.native))
	var retGo (*Item)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free a #PangoItem and all associated memory.
/*

C function

pango_item_free
*/
func (recv *Item) Free() {
	C.pango_item_free((*C.PangoItem)(recv.native))

	return
}

// Modifies @orig to cover only the text after @split_index, and
// returns a new item that covers the text before @split_index that
// used to be in @orig. You can think of @split_index as the length of
// the returned item. @split_index may not be 0, and it may not be
// greater than or equal to the length of @orig (that is, there must
// be at least one byte assigned to each item, you can't create a
// zero-length item). @split_offset is the length of the first item in
// chars, and must be provided because the text used to generate the
// item isn't available, so pango_item_split() can't count the char
// length of the split items itself.
/*

C function

pango_item_split
*/
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	c_split_index := (C.int)(splitIndex)

	c_split_offset := (C.int)(splitOffset)

	retC := C.pango_item_split((*C.PangoItem)(recv.native), c_split_index, c_split_offset)
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// The #PangoLanguage structure is used to
// represent a language.
//
// #PangoLanguage pointers can be efficiently
// copied and compared with each other.
/*

C type

PangoLanguage
*/
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

// Get a string that is representative of the characters needed to
// render a particular language.
//
// The sample text may be a pangram, but is not necessarily.  It is chosen to
// be demonstrative of normal text in the language, as well as exposing font
// feature requirements unique to the language.  It is suitable for use
// as sample text in a font selection dialog.
//
// If @language is %NULL, the default language as found by
// pango_language_get_default() is used.
//
// If Pango does not have a sample string for @language, the classic
// "The quick brown fox..." is returned.  This can be detected by
// comparing the returned pointer value to that returned for (non-existent)
// language code "xx".  That is, compare to:
// <informalexample><programlisting>
// pango_language_get_sample_string (pango_language_from_string ("xx"))
// </programlisting></informalexample>
/*

C function

pango_language_get_sample_string
*/
func (recv *Language) GetSampleString() string {
	retC := C.pango_language_get_sample_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks if a language tag matches one of the elements in a list of
// language ranges. A language tag is considered to match a range
// in the list if the range is '*', the range is exactly the tag,
// or the range is a prefix of the tag, and the character after it
// in the tag is '-'.
/*

C function

pango_language_matches
*/
func (recv *Language) Matches(rangeList string) bool {
	c_range_list := C.CString(rangeList)
	defer C.free(unsafe.Pointer(c_range_list))

	retC := C.pango_language_matches((*C.PangoLanguage)(recv.native), c_range_list)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the RFC-3066 format string representing the given language tag.
/*

C function

pango_language_to_string
*/
func (recv *Language) ToString() string {
	retC := C.pango_language_to_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

/*

C type

PangoLayoutClass
*/
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

// The #PangoLayoutLine structure represents one of the lines resulting
// from laying out a paragraph via #PangoLayout. #PangoLayoutLine
// structures are obtained by calling pango_layout_get_line() and
// are only valid until the text, attributes, or settings of the
// parent #PangoLayout are modified.
//
// Routines for rendering PangoLayout objects are provided in
// code specific to each rendering system.
/*

C type

PangoLayoutLine
*/
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

// Computes the logical and ink extents of a layout line. See
// pango_font_get_glyph_extents() for details about the interpretation
// of the rectangles.
/*

C function

pango_layout_line_get_extents
*/
func (recv *LayoutLine) GetExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Computes the logical and ink extents of @layout_line in device units.
// This function just calls pango_layout_line_get_extents() followed by
// two pango_extents_to_pixels() calls, rounding @ink_rect and @logical_rect
// such that the rounded rectangles fully contain the unrounded one (that is,
// passes them as first argument to pango_extents_to_pixels()).
/*

C function

pango_layout_line_get_pixel_extents
*/
func (recv *LayoutLine) GetPixelExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_pixel_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : output array param ranges

// Converts an index within a line to a X position.
/*

C function

pango_layout_line_index_to_x
*/
func (recv *LayoutLine) IndexToX(index int32, trailing bool) int32 {
	c_index_ := (C.int)(index)

	c_trailing :=
		boolToGboolean(trailing)

	var c_x_pos C.int

	C.pango_layout_line_index_to_x((*C.PangoLayoutLine)(recv.native), c_index_, c_trailing, &c_x_pos)

	xPos := (int32)(c_x_pos)

	return xPos
}

// Decrease the reference count of a #PangoLayoutLine by one.
// If the result is zero, the line and all associated memory
// will be freed.
/*

C function

pango_layout_line_unref
*/
func (recv *LayoutLine) Unref() {
	C.pango_layout_line_unref((*C.PangoLayoutLine)(recv.native))

	return
}

// Converts from x offset to the byte index of the corresponding
// character within the text of the layout. If @x_pos is outside the line,
// @index_ and @trailing will point to the very first or very last position
// in the line. This determination is based on the resolved direction
// of the paragraph; for example, if the resolved direction is
// right-to-left, then an X position to the right of the line (after it)
// results in 0 being stored in @index_ and @trailing. An X position to the
// left of the line results in @index_ pointing to the (logical) last
// grapheme in the line and @trailing being set to the number of characters
// in that grapheme. The reverse is true for a left-to-right line.
/*

C function

pango_layout_line_x_to_index
*/
func (recv *LayoutLine) XToIndex(xPos int32) (bool, int32, int32) {
	c_x_pos := (C.int)(xPos)

	var c_index_ C.int

	var c_trailing C.int

	retC := C.pango_layout_line_x_to_index((*C.PangoLayoutLine)(recv.native), c_x_pos, &c_index_, &c_trailing)
	retGo := retC == C.TRUE

	index := (int32)(c_index_)

	trailing := (int32)(c_trailing)

	return retGo, index, trailing
}

// The #PangoLogAttr structure stores information
// about the attributes of a single character.
/*

C type

PangoLogAttr
*/
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

// The #PangoRectangle structure represents a rectangle. It is frequently
// used to represent the logical or ink extents of a single glyph or section
// of text. (See, for instance, pango_font_get_glyph_extents())
/*

C type

PangoRectangle
*/
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

/*

C type

PangoRendererPrivate
*/
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

// A #PangoScriptIter is used to iterate through a string
// and identify ranges in different scripts.
/*

C type

PangoScriptIter
*/
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

// A #PangoTabArray struct contains an array
// of tab stops. Each tab stop has an alignment and a position.
/*

C type

PangoTabArray
*/
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

// Creates an array of @initial_size tab stops. Tab stops are specified in
// pixel units if @positions_in_pixels is %TRUE, otherwise in Pango
// units. All stops are initially at position 0.
/*

C function

pango_tab_array_new
*/
func TabArrayNew(initialSize int32, positionsInPixels bool) *TabArray {
	c_initial_size := (C.gint)(initialSize)

	c_positions_in_pixels :=
		boolToGboolean(positionsInPixels)

	retC := C.pango_tab_array_new(c_initial_size, c_positions_in_pixels)
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Copies a #PangoTabArray
/*

C function

pango_tab_array_copy
*/
func (recv *TabArray) Copy() *TabArray {
	retC := C.pango_tab_array_copy((*C.PangoTabArray)(recv.native))
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a tab array and associated resources.
/*

C function

pango_tab_array_free
*/
func (recv *TabArray) Free() {
	C.pango_tab_array_free((*C.PangoTabArray)(recv.native))

	return
}

// Returns %TRUE if the tab positions are in pixels, %FALSE if they are
// in Pango units.
/*

C function

pango_tab_array_get_positions_in_pixels
*/
func (recv *TabArray) GetPositionsInPixels() bool {
	retC := C.pango_tab_array_get_positions_in_pixels((*C.PangoTabArray)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the number of tab stops in @tab_array.
/*

C function

pango_tab_array_get_size
*/
func (recv *TabArray) GetSize() int32 {
	retC := C.pango_tab_array_get_size((*C.PangoTabArray)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2

// Resizes a tab array. You must subsequently initialize any tabs that
// were added as a result of growing the array.
/*

C function

pango_tab_array_resize
*/
func (recv *TabArray) Resize(newSize int32) {
	c_new_size := (C.gint)(newSize)

	C.pango_tab_array_resize((*C.PangoTabArray)(recv.native), c_new_size)

	return
}

// Sets the alignment and location of a tab stop.
// @alignment must always be #PANGO_TAB_LEFT in the current
// implementation.
/*

C function

pango_tab_array_set_tab
*/
func (recv *TabArray) SetTab(tabIndex int32, alignment TabAlign, location int32) {
	c_tab_index := (C.gint)(tabIndex)

	c_alignment := (C.PangoTabAlign)(alignment)

	c_location := (C.gint)(location)

	C.pango_tab_array_set_tab((*C.PangoTabArray)(recv.native), c_tab_index, c_alignment, c_location)

	return
}
