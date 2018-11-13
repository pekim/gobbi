// This is a generated file - DO NOT EDIT

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// The #PangoContext structure stores global information
// used to control the itemization process.
/*

C record/class : PangoContext
*/
type Context struct {
	native *C.PangoContext
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.PangoContext)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Context) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Context.
// Exercise care, as this is a potentially dangerous function if the Object is not a Context.
func CastToContext(object *gobject.Object) *Context {
	return ContextNewFromC(object.ToC())
}

// Creates a new #PangoContext initialized to default values.
//
// This function is not particularly useful as it should always
// be followed by a pango_context_set_font_map() call, and the
// function pango_font_map_create_context() does these two steps
// together and hence users are recommended to use that.
//
// If you are using Pango as part of a higher-level system,
// that system may have it's own way of create a #PangoContext.
// For instance, the GTK+ toolkit has, among others,
// gdk_pango_context_get_for_screen(), and
// gtk_widget_get_pango_context().  Use those instead.
/*

C function : pango_context_new
*/
func ContextNew() *Context {
	retC := C.pango_context_new()
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the base direction for the context. See
// pango_context_set_base_dir().
/*

C function : pango_context_get_base_dir
*/
func (recv *Context) GetBaseDir() Direction {
	retC := C.pango_context_get_base_dir((*C.PangoContext)(recv.native))
	retGo := (Direction)(retC)

	return retGo
}

// Retrieve the default font description for the context.
/*

C function : pango_context_get_font_description
*/
func (recv *Context) GetFontDescription() *FontDescription {
	retC := C.pango_context_get_font_description((*C.PangoContext)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the global language tag for the context.
/*

C function : pango_context_get_language
*/
func (recv *Context) GetLanguage() *Language {
	retC := C.pango_context_get_language((*C.PangoContext)(recv.native))
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_context_list_families : unsupported parameter families : output array param families

// Loads the font in one of the fontmaps in the context
// that is the closest match for @desc.
/*

C function : pango_context_load_font
*/
func (recv *Context) LoadFont(desc *FontDescription) *Font {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	retC := C.pango_context_load_font((*C.PangoContext)(recv.native), c_desc)
	var retGo (*Font)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Load a set of fonts in the context that can be used to render
// a font matching @desc.
/*

C function : pango_context_load_fontset
*/
func (recv *Context) LoadFontset(desc *FontDescription, language *Language) *Fontset {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_context_load_fontset((*C.PangoContext)(recv.native), c_desc, c_language)
	var retGo (*Fontset)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontsetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the base direction for the context.
//
// The base direction is used in applying the Unicode bidirectional
// algorithm; if the @direction is %PANGO_DIRECTION_LTR or
// %PANGO_DIRECTION_RTL, then the value will be used as the paragraph
// direction in the Unicode bidirectional algorithm.  A value of
// %PANGO_DIRECTION_WEAK_LTR or %PANGO_DIRECTION_WEAK_RTL is used only
// for paragraphs that do not contain any strong characters themselves.
/*

C function : pango_context_set_base_dir
*/
func (recv *Context) SetBaseDir(direction Direction) {
	c_direction := (C.PangoDirection)(direction)

	C.pango_context_set_base_dir((*C.PangoContext)(recv.native), c_direction)

	return
}

// Set the default font description for the context
/*

C function : pango_context_set_font_description
*/
func (recv *Context) SetFontDescription(desc *FontDescription) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	C.pango_context_set_font_description((*C.PangoContext)(recv.native), c_desc)

	return
}

// Sets the font map to be searched when fonts are looked-up in this context.
// This is only for internal use by Pango backends, a #PangoContext obtained
// via one of the recommended methods should already have a suitable font map.
/*

C function : pango_context_set_font_map
*/
func (recv *Context) SetFontMap(fontMap *FontMap) {
	c_font_map := (*C.PangoFontMap)(C.NULL)
	if fontMap != nil {
		c_font_map = (*C.PangoFontMap)(fontMap.ToC())
	}

	C.pango_context_set_font_map((*C.PangoContext)(recv.native), c_font_map)

	return
}

// Sets the global language tag for the context.  The default language
// for the locale of the running process can be found using
// pango_language_get_default().
/*

C function : pango_context_set_language
*/
func (recv *Context) SetLanguage(language *Language) {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	C.pango_context_set_language((*C.PangoContext)(recv.native), c_language)

	return
}

// Blacklisted : PangoEngine

// The #PangoEngineLang class is implemented by engines that
// customize the rendering-system independent part of the
// Pango pipeline for a particular script or language. For
// instance, a custom #PangoEngineLang could be provided for
// Thai to implement the dictionary-based word boundary
// lookups needed for that language.
/*

C record/class : PangoEngineLang
*/
type EngineLang struct {
	native *C.PangoEngineLang
	// Private : parent_instance
}

func EngineLangNewFromC(u unsafe.Pointer) *EngineLang {
	c := (*C.PangoEngineLang)(u)
	if c == nil {
		return nil
	}

	g := &EngineLang{native: c}

	return g
}

func (recv *EngineLang) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CastToWidget down casts any arbitary Object to EngineLang.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineLang.
func CastToEngineLang(object *gobject.Object) *EngineLang {
	return EngineLangNewFromC(object.ToC())
}

// The #PangoEngineShape class is implemented by engines that
// customize the rendering-system dependent part of the
// Pango pipeline for a particular script or language.
// A #PangoEngineShape implementation is then specific to both
// a particular rendering system or group of rendering systems
// and to a particular script. For instance, there is one
// #PangoEngineShape implementation to handle shaping Arabic
// for Fontconfig-based backends.
/*

C record/class : PangoEngineShape
*/
type EngineShape struct {
	native *C.PangoEngineShape
	// parent_instance : record
}

func EngineShapeNewFromC(u unsafe.Pointer) *EngineShape {
	c := (*C.PangoEngineShape)(u)
	if c == nil {
		return nil
	}

	g := &EngineShape{native: c}

	return g
}

func (recv *EngineShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CastToWidget down casts any arbitary Object to EngineShape.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineShape.
func CastToEngineShape(object *gobject.Object) *EngineShape {
	return EngineShapeNewFromC(object.ToC())
}

// The #PangoFont structure is used to represent
// a font in a rendering-system-independent matter.
// To create an implementation of a #PangoFont,
// the rendering-system specific code should allocate
// a larger structure that contains a nested
// #PangoFont, fill in the <structfield>klass</structfield> member of
// the nested #PangoFont with a pointer to
// a appropriate #PangoFontClass, then call
// pango_font_init() on the structure.
//
// The #PangoFont structure contains one member
// which the implementation fills in.
/*

C record/class : PangoFont
*/
type Font struct {
	native *C.PangoFont
	// parent_instance : record
}

func FontNewFromC(u unsafe.Pointer) *Font {
	c := (*C.PangoFont)(u)
	if c == nil {
		return nil
	}

	g := &Font{native: c}

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Font) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Font.
// Exercise care, as this is a potentially dangerous function if the Object is not a Font.
func CastToFont(object *gobject.Object) *Font {
	return FontNewFromC(object.ToC())
}

// Returns a description of the font, with font size set in points.
// Use pango_font_describe_with_absolute_size() if you want the font
// size in device units.
/*

C function : pango_font_describe
*/
func (recv *Font) Describe() *FontDescription {
	retC := C.pango_font_describe((*C.PangoFont)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds the best matching shaper for a font for a particular
// language tag and character point.
/*

C function : pango_font_find_shaper
*/
func (recv *Font) FindShaper(language *Language, ch uint32) *EngineShape {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	c_ch := (C.guint32)(ch)

	retC := C.pango_font_find_shaper((*C.PangoFont)(recv.native), c_language, c_ch)
	retGo := EngineShapeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Computes the coverage map for a given font and language tag.
/*

C function : pango_font_get_coverage
*/
func (recv *Font) GetCoverage(language *Language) *Coverage {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_font_get_coverage((*C.PangoFont)(recv.native), c_language)
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the logical and ink extents of a glyph within a font. The
// coordinate system for each rectangle has its origin at the
// base line and horizontal origin of the character with increasing
// coordinates extending to the right and down. The macros PANGO_ASCENT(),
// PANGO_DESCENT(), PANGO_LBEARING(), and PANGO_RBEARING() can be used to convert
// from the extents rectangle to more traditional font metrics. The units
// of the rectangles are in 1/PANGO_SCALE of a device unit.
//
// If @font is %NULL, this function gracefully sets some sane values in the
// output variables and returns.
/*

C function : pango_font_get_glyph_extents
*/
func (recv *Font) GetGlyphExtents(glyph Glyph) (*Rectangle, *Rectangle) {
	c_glyph := (C.PangoGlyph)(glyph)

	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_font_get_glyph_extents((*C.PangoFont)(recv.native), c_glyph, &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// The #PangoFontFace structure is used to represent a group of fonts with
// the same family, slant, weight, width, but varying sizes.
/*

C record/class : PangoFontFace
*/
type FontFace struct {
	native *C.PangoFontFace
	// parent_instance : record
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	c := (*C.PangoFontFace)(u)
	if c == nil {
		return nil
	}

	g := &FontFace{native: c}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FontFace) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontFace.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontFace.
func CastToFontFace(object *gobject.Object) *FontFace {
	return FontFaceNewFromC(object.ToC())
}

// Returns the family, style, variant, weight and stretch of
// a #PangoFontFace. The size field of the resulting font description
// will be unset.
/*

C function : pango_font_face_describe
*/
func (recv *FontFace) Describe() *FontDescription {
	retC := C.pango_font_face_describe((*C.PangoFontFace)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a name representing the style of this face among the
// different faces in the #PangoFontFamily for the face. This
// name is unique among all faces in the family and is suitable
// for displaying to users.
/*

C function : pango_font_face_get_face_name
*/
func (recv *FontFace) GetFaceName() string {
	retC := C.pango_font_face_get_face_name((*C.PangoFontFace)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// The #PangoFontFamily structure is used to represent a family of related
// font faces. The faces in a family share a common design, but differ in
// slant, weight, width and other aspects.
/*

C record/class : PangoFontFamily
*/
type FontFamily struct {
	native *C.PangoFontFamily
	// parent_instance : record
}

func FontFamilyNewFromC(u unsafe.Pointer) *FontFamily {
	c := (*C.PangoFontFamily)(u)
	if c == nil {
		return nil
	}

	g := &FontFamily{native: c}

	return g
}

func (recv *FontFamily) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FontFamily) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontFamily.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontFamily.
func CastToFontFamily(object *gobject.Object) *FontFamily {
	return FontFamilyNewFromC(object.ToC())
}

// Gets the name of the family. The name is unique among all
// fonts for the font backend and can be used in a #PangoFontDescription
// to specify that a face from this family is desired.
/*

C function : pango_font_family_get_name
*/
func (recv *FontFamily) GetName() string {
	retC := C.pango_font_family_get_name((*C.PangoFontFamily)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : output array param faces

// The #PangoFontMap represents the set of fonts available for a
// particular rendering system. This is a virtual object with
// implementations being specific to particular rendering systems.  To
// create an implementation of a #PangoFontMap, the rendering-system
// specific code should allocate a larger structure that contains a nested
// #PangoFontMap, fill in the <structfield>klass</structfield> member of the nested #PangoFontMap with a
// pointer to a appropriate #PangoFontMapClass, then call
// pango_font_map_init() on the structure.
//
// The #PangoFontMap structure contains one member which the implementation
// fills in.
/*

C record/class : PangoFontMap
*/
type FontMap struct {
	native *C.PangoFontMap
	// parent_instance : record
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoFontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FontMap) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontMap.
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromC(object.ToC())
}

// Unsupported : pango_font_map_list_families : unsupported parameter families : output array param families

// Load the font in the fontmap that is the closest match for @desc.
/*

C function : pango_font_map_load_font
*/
func (recv *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	retC := C.pango_font_map_load_font((*C.PangoFontMap)(recv.native), c_context, c_desc)
	var retGo (*Font)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Load a set of fonts in the fontmap that can be used to render
// a font matching @desc.
/*

C function : pango_font_map_load_fontset
*/
func (recv *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_font_map_load_fontset((*C.PangoFontMap)(recv.native), c_context, c_desc, c_language)
	var retGo (*Fontset)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontsetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// A #PangoFontset represents a set of #PangoFont to use
// when rendering text. It is the result of resolving a
// #PangoFontDescription against a particular #PangoContext.
// It has operations for finding the component font for
// a particular Unicode character, and for finding a composite
// set of metrics for the entire fontset.
/*

C record/class : PangoFontset
*/
type Fontset struct {
	native *C.PangoFontset
	// parent_instance : record
}

func FontsetNewFromC(u unsafe.Pointer) *Fontset {
	c := (*C.PangoFontset)(u)
	if c == nil {
		return nil
	}

	g := &Fontset{native: c}

	return g
}

func (recv *Fontset) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Fontset) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Fontset.
// Exercise care, as this is a potentially dangerous function if the Object is not a Fontset.
func CastToFontset(object *gobject.Object) *Fontset {
	return FontsetNewFromC(object.ToC())
}

// Returns the font in the fontset that contains the best glyph for the
// Unicode character @wc.
/*

C function : pango_fontset_get_font
*/
func (recv *Fontset) GetFont(wc uint32) *Font {
	c_wc := (C.guint)(wc)

	retC := C.pango_fontset_get_font((*C.PangoFontset)(recv.native), c_wc)
	retGo := FontNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Blacklisted : PangoFontsetSimple

// The #PangoLayout structure represents an entire paragraph
// of text. It is initialized with a #PangoContext, UTF-8 string
// and set of attributes for that string. Once that is done, the
// set of formatted lines can be extracted from the object,
// the layout can be rendered, and conversion between logical
// character positions within the layout's text, and the physical
// position of the resulting glyphs can be made.
//
// There are also a number of parameters to adjust the formatting
// of a #PangoLayout, which are illustrated in <xref linkend="parameters"/>.
// It is possible, as well, to ignore the 2-D setup, and simply
// treat the results of a #PangoLayout as a list of lines.
//
// <figure id="parameters">
// <title>Adjustable parameters for a PangoLayout</title>
// <graphic fileref="layout.gif" format="GIF"></graphic>
// </figure>
//
// The #PangoLayout structure is opaque, and has no user-visible
// fields.
/*

C record/class : PangoLayout
*/
type Layout struct {
	native *C.PangoLayout
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	c := (*C.PangoLayout)(u)
	if c == nil {
		return nil
	}

	g := &Layout{native: c}

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Layout) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Layout.
// Exercise care, as this is a potentially dangerous function if the Object is not a Layout.
func CastToLayout(object *gobject.Object) *Layout {
	return LayoutNewFromC(object.ToC())
}

// Create a new #PangoLayout object with attributes initialized to
// default values for a particular #PangoContext.
/*

C function : pango_layout_new
*/
func LayoutNew(context *Context) *Layout {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	retC := C.pango_layout_new(c_context)
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Forces recomputation of any state in the #PangoLayout that
// might depend on the layout's context. This function should
// be called if you make changes to the context subsequent
// to creating the layout.
/*

C function : pango_layout_context_changed
*/
func (recv *Layout) ContextChanged() {
	C.pango_layout_context_changed((*C.PangoLayout)(recv.native))

	return
}

// Does a deep copy-by-value of the @src layout. The attribute list,
// tab array, and text from the original layout are all copied by
// value.
/*

C function : pango_layout_copy
*/
func (recv *Layout) Copy() *Layout {
	retC := C.pango_layout_copy((*C.PangoLayout)(recv.native))
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the alignment for the layout: how partial lines are
// positioned within the horizontal space available.
/*

C function : pango_layout_get_alignment
*/
func (recv *Layout) GetAlignment() Alignment {
	retC := C.pango_layout_get_alignment((*C.PangoLayout)(recv.native))
	retGo := (Alignment)(retC)

	return retGo
}

// Gets the attribute list for the layout, if any.
/*

C function : pango_layout_get_attributes
*/
func (recv *Layout) GetAttributes() *AttrList {
	retC := C.pango_layout_get_attributes((*C.PangoLayout)(recv.native))
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the #PangoContext used for this layout.
/*

C function : pango_layout_get_context
*/
func (recv *Layout) GetContext() *Context {
	retC := C.pango_layout_get_context((*C.PangoLayout)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Given an index within a layout, determines the positions that of the
// strong and weak cursors if the insertion point is at that
// index. The position of each cursor is stored as a zero-width
// rectangle. The strong cursor location is the location where
// characters of the directionality equal to the base direction of the
// layout are inserted.  The weak cursor location is the location
// where characters of the directionality opposite to the base
// direction of the layout are inserted.
/*

C function : pango_layout_get_cursor_pos
*/
func (recv *Layout) GetCursorPos(index int32) (*Rectangle, *Rectangle) {
	c_index_ := (C.int)(index)

	var c_strong_pos C.PangoRectangle

	var c_weak_pos C.PangoRectangle

	C.pango_layout_get_cursor_pos((*C.PangoLayout)(recv.native), c_index_, &c_strong_pos, &c_weak_pos)

	strongPos := RectangleNewFromC(unsafe.Pointer(&c_strong_pos))

	weakPos := RectangleNewFromC(unsafe.Pointer(&c_weak_pos))

	return strongPos, weakPos
}

// Computes the logical and ink extents of @layout. Logical extents
// are usually what you want for positioning things.  Note that both extents
// may have non-zero x and y.  You may want to use those to offset where you
// render the layout.  Not doing that is a very typical bug that shows up as
// right-to-left layouts not being correctly positioned in a layout with
// a set width.
//
// The extents are given in layout coordinates and in Pango units; layout
// coordinates begin at the top left corner of the layout.
/*

C function : pango_layout_get_extents
*/
func (recv *Layout) GetExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_get_extents((*C.PangoLayout)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Gets the paragraph indent width in Pango units. A negative value
// indicates a hanging indentation.
/*

C function : pango_layout_get_indent
*/
func (recv *Layout) GetIndent() int32 {
	retC := C.pango_layout_get_indent((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Gets whether each complete line should be stretched to fill the entire
// width of the layout.
/*

C function : pango_layout_get_justify
*/
func (recv *Layout) GetJustify() bool {
	retC := C.pango_layout_get_justify((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves a particular line from a #PangoLayout.
//
// Use the faster pango_layout_get_line_readonly() if you do not plan
// to modify the contents of the line (glyphs, glyph widths, etc.).
/*

C function : pango_layout_get_line
*/
func (recv *Layout) GetLine(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line((*C.PangoLayout)(recv.native), c_line)
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Retrieves the count of lines for the @layout.
/*

C function : pango_layout_get_line_count
*/
func (recv *Layout) GetLineCount() int32 {
	retC := C.pango_layout_get_line_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the lines of the @layout as a list.
//
// Use the faster pango_layout_get_lines_readonly() if you do not plan
// to modify the contents of the lines (glyphs, glyph widths, etc.).
/*

C function : pango_layout_get_lines
*/
func (recv *Layout) GetLines() *glib.SList {
	retC := C.pango_layout_get_lines((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : output array param attrs

// Computes the logical and ink extents of @layout in device units.
// This function just calls pango_layout_get_extents() followed by
// two pango_extents_to_pixels() calls, rounding @ink_rect and @logical_rect
// such that the rounded rectangles fully contain the unrounded one (that is,
// passes them as first argument to pango_extents_to_pixels()).
/*

C function : pango_layout_get_pixel_extents
*/
func (recv *Layout) GetPixelExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_get_pixel_extents((*C.PangoLayout)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Determines the logical width and height of a #PangoLayout
// in device units. (pango_layout_get_size() returns the width
// and height scaled by %PANGO_SCALE.) This
// is simply a convenience function around
// pango_layout_get_pixel_extents().
/*

C function : pango_layout_get_pixel_size
*/
func (recv *Layout) GetPixelSize() (int32, int32) {
	var c_width C.int

	var c_height C.int

	C.pango_layout_get_pixel_size((*C.PangoLayout)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Obtains the value set by pango_layout_set_single_paragraph_mode().
/*

C function : pango_layout_get_single_paragraph_mode
*/
func (recv *Layout) GetSingleParagraphMode() bool {
	retC := C.pango_layout_get_single_paragraph_mode((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines the logical width and height of a #PangoLayout
// in Pango units (device units scaled by %PANGO_SCALE). This
// is simply a convenience function around pango_layout_get_extents().
/*

C function : pango_layout_get_size
*/
func (recv *Layout) GetSize() (int32, int32) {
	var c_width C.int

	var c_height C.int

	C.pango_layout_get_size((*C.PangoLayout)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Gets the amount of spacing between the lines of the layout.
/*

C function : pango_layout_get_spacing
*/
func (recv *Layout) GetSpacing() int32 {
	retC := C.pango_layout_get_spacing((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the current #PangoTabArray used by this layout. If no
// #PangoTabArray has been set, then the default tabs are in use
// and %NULL is returned. Default tabs are every 8 spaces.
// The return value should be freed with pango_tab_array_free().
/*

C function : pango_layout_get_tabs
*/
func (recv *Layout) GetTabs() *TabArray {
	retC := C.pango_layout_get_tabs((*C.PangoLayout)(recv.native))
	var retGo (*TabArray)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TabArrayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the text in the layout. The returned text should not
// be freed or modified.
/*

C function : pango_layout_get_text
*/
func (recv *Layout) GetText() string {
	retC := C.pango_layout_get_text((*C.PangoLayout)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the width to which the lines of the #PangoLayout should wrap.
/*

C function : pango_layout_get_width
*/
func (recv *Layout) GetWidth() int32 {
	retC := C.pango_layout_get_width((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the wrap mode for the layout.
//
// Use pango_layout_is_wrapped() to query whether any paragraphs
// were actually wrapped.
/*

C function : pango_layout_get_wrap
*/
func (recv *Layout) GetWrap() WrapMode {
	retC := C.pango_layout_get_wrap((*C.PangoLayout)(recv.native))
	retGo := (WrapMode)(retC)

	return retGo
}

// Converts from byte @index_ within the @layout to line and X position.
// (X position is measured from the left edge of the line)
/*

C function : pango_layout_index_to_line_x
*/
func (recv *Layout) IndexToLineX(index int32, trailing bool) (int32, int32) {
	c_index_ := (C.int)(index)

	c_trailing :=
		boolToGboolean(trailing)

	var c_line C.int

	var c_x_pos C.int

	C.pango_layout_index_to_line_x((*C.PangoLayout)(recv.native), c_index_, c_trailing, &c_line, &c_x_pos)

	line := (int32)(c_line)

	xPos := (int32)(c_x_pos)

	return line, xPos
}

// Converts from an index within a #PangoLayout to the onscreen position
// corresponding to the grapheme at that index, which is represented
// as rectangle.  Note that <literal>pos->x</literal> is always the leading
// edge of the grapheme and <literal>pos->x + pos->width</literal> the trailing
// edge of the grapheme. If the directionality of the grapheme is right-to-left,
// then <literal>pos->width</literal> will be negative.
/*

C function : pango_layout_index_to_pos
*/
func (recv *Layout) IndexToPos(index int32) *Rectangle {
	c_index_ := (C.int)(index)

	var c_pos C.PangoRectangle

	C.pango_layout_index_to_pos((*C.PangoLayout)(recv.native), c_index_, &c_pos)

	pos := RectangleNewFromC(unsafe.Pointer(&c_pos))

	return pos
}

// Computes a new cursor position from an old position and
// a count of positions to move visually. If @direction is positive,
// then the new strong cursor position will be one position
// to the right of the old cursor position. If @direction is negative,
// then the new strong cursor position will be one position
// to the left of the old cursor position.
//
// In the presence of bidirectional text, the correspondence
// between logical and visual order will depend on the direction
// of the current run, and there may be jumps when the cursor
// is moved off of the end of a run.
//
// Motion here is in cursor positions, not in characters, so a
// single call to pango_layout_move_cursor_visually() may move the
// cursor over multiple characters when multiple characters combine
// to form a single grapheme.
/*

C function : pango_layout_move_cursor_visually
*/
func (recv *Layout) MoveCursorVisually(strong bool, oldIndex int32, oldTrailing int32, direction int32) (int32, int32) {
	c_strong :=
		boolToGboolean(strong)

	c_old_index := (C.int)(oldIndex)

	c_old_trailing := (C.int)(oldTrailing)

	c_direction := (C.int)(direction)

	var c_new_index C.int

	var c_new_trailing C.int

	C.pango_layout_move_cursor_visually((*C.PangoLayout)(recv.native), c_strong, c_old_index, c_old_trailing, c_direction, &c_new_index, &c_new_trailing)

	newIndex := (int32)(c_new_index)

	newTrailing := (int32)(c_new_trailing)

	return newIndex, newTrailing
}

// Sets the alignment for the layout: how partial lines are
// positioned within the horizontal space available.
/*

C function : pango_layout_set_alignment
*/
func (recv *Layout) SetAlignment(alignment Alignment) {
	c_alignment := (C.PangoAlignment)(alignment)

	C.pango_layout_set_alignment((*C.PangoLayout)(recv.native), c_alignment)

	return
}

// Sets the text attributes for a layout object.
// References @attrs, so the caller can unref its reference.
/*

C function : pango_layout_set_attributes
*/
func (recv *Layout) SetAttributes(attrs *AttrList) {
	c_attrs := (*C.PangoAttrList)(C.NULL)
	if attrs != nil {
		c_attrs = (*C.PangoAttrList)(attrs.ToC())
	}

	C.pango_layout_set_attributes((*C.PangoLayout)(recv.native), c_attrs)

	return
}

// Sets the default font description for the layout. If no font
// description is set on the layout, the font description from
// the layout's context is used.
/*

C function : pango_layout_set_font_description
*/
func (recv *Layout) SetFontDescription(desc *FontDescription) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	C.pango_layout_set_font_description((*C.PangoLayout)(recv.native), c_desc)

	return
}

// Sets the width in Pango units to indent each paragraph. A negative value
// of @indent will produce a hanging indentation. That is, the first line will
// have the full width, and subsequent lines will be indented by the
// absolute value of @indent.
//
// The indent setting is ignored if layout alignment is set to
// %PANGO_ALIGN_CENTER.
/*

C function : pango_layout_set_indent
*/
func (recv *Layout) SetIndent(indent int32) {
	c_indent := (C.int)(indent)

	C.pango_layout_set_indent((*C.PangoLayout)(recv.native), c_indent)

	return
}

// Sets whether each complete line should be stretched to
// fill the entire width of the layout. This stretching is typically
// done by adding whitespace, but for some scripts (such as Arabic),
// the justification may be done in more complex ways, like extending
// the characters.
//
// Note that this setting is not implemented and so is ignored in Pango
// older than 1.18.
/*

C function : pango_layout_set_justify
*/
func (recv *Layout) SetJustify(justify bool) {
	c_justify :=
		boolToGboolean(justify)

	C.pango_layout_set_justify((*C.PangoLayout)(recv.native), c_justify)

	return
}

// Same as pango_layout_set_markup_with_accel(), but
// the markup text isn't scanned for accelerators.
/*

C function : pango_layout_set_markup
*/
func (recv *Layout) SetMarkup(markup string, length int32) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_length := (C.int)(length)

	C.pango_layout_set_markup((*C.PangoLayout)(recv.native), c_markup, c_length)

	return
}

// Blacklisted : pango_layout_set_markup_with_accel

// If @setting is %TRUE, do not treat newlines and similar characters
// as paragraph separators; instead, keep all text in a single paragraph,
// and display a glyph for paragraph separator characters. Used when
// you want to allow editing of newlines on a single text line.
/*

C function : pango_layout_set_single_paragraph_mode
*/
func (recv *Layout) SetSingleParagraphMode(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.pango_layout_set_single_paragraph_mode((*C.PangoLayout)(recv.native), c_setting)

	return
}

// Sets the amount of spacing in Pango unit between the lines of the
// layout.
/*

C function : pango_layout_set_spacing
*/
func (recv *Layout) SetSpacing(spacing int32) {
	c_spacing := (C.int)(spacing)

	C.pango_layout_set_spacing((*C.PangoLayout)(recv.native), c_spacing)

	return
}

// Sets the tabs to use for @layout, overriding the default tabs
// (by default, tabs are every 8 spaces). If @tabs is %NULL, the default
// tabs are reinstated. @tabs is copied into the layout; you must
// free your copy of @tabs yourself.
/*

C function : pango_layout_set_tabs
*/
func (recv *Layout) SetTabs(tabs *TabArray) {
	c_tabs := (*C.PangoTabArray)(C.NULL)
	if tabs != nil {
		c_tabs = (*C.PangoTabArray)(tabs.ToC())
	}

	C.pango_layout_set_tabs((*C.PangoLayout)(recv.native), c_tabs)

	return
}

// Sets the text of the layout.
//
// Note that if you have used
// pango_layout_set_markup() or pango_layout_set_markup_with_accel() on
// @layout before, you may want to call pango_layout_set_attributes() to clear
// the attributes set on the layout from the markup as this function does not
// clear attributes.
/*

C function : pango_layout_set_text
*/
func (recv *Layout) SetText(text string, length int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	C.pango_layout_set_text((*C.PangoLayout)(recv.native), c_text, c_length)

	return
}

// Sets the width to which the lines of the #PangoLayout should wrap or
// ellipsized.  The default value is -1: no width set.
/*

C function : pango_layout_set_width
*/
func (recv *Layout) SetWidth(width int32) {
	c_width := (C.int)(width)

	C.pango_layout_set_width((*C.PangoLayout)(recv.native), c_width)

	return
}

// Sets the wrap mode; the wrap mode only has effect if a width
// is set on the layout with pango_layout_set_width().
// To turn off wrapping, set the width to -1.
/*

C function : pango_layout_set_wrap
*/
func (recv *Layout) SetWrap(wrap WrapMode) {
	c_wrap := (C.PangoWrapMode)(wrap)

	C.pango_layout_set_wrap((*C.PangoLayout)(recv.native), c_wrap)

	return
}

// Converts from X and Y position within a layout to the byte
// index to the character at that logical position. If the
// Y position is not inside the layout, the closest position is chosen
// (the position will be clamped inside the layout). If the
// X position is not within the layout, then the start or the
// end of the line is chosen as described for pango_layout_line_x_to_index().
// If either the X or Y positions were not inside the layout, then the
// function returns %FALSE; on an exact hit, it returns %TRUE.
/*

C function : pango_layout_xy_to_index
*/
func (recv *Layout) XyToIndex(x int32, y int32) (bool, int32, int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	var c_index_ C.int

	var c_trailing C.int

	retC := C.pango_layout_xy_to_index((*C.PangoLayout)(recv.native), c_x, c_y, &c_index_, &c_trailing)
	retGo := retC == C.TRUE

	index := (int32)(c_index_)

	trailing := (int32)(c_trailing)

	return retGo, index, trailing
}
