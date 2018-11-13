// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// A #PangoGlyphItemIter is an iterator over the clusters in a
// #PangoGlyphItem.  The <firstterm>forward direction</firstterm> of the
// iterator is the logical direction of text.  That is, with increasing
// @start_index and @start_char values.  If @glyph_item is right-to-left
// (that is, if <literal>@glyph_item->item->analysis.level</literal> is odd),
// then @start_glyph decreases as the iterator moves forward.  Moreover,
// in right-to-left cases, @start_glyph is greater than @end_glyph.
//
// An iterator should be initialized using either of
// pango_glyph_item_iter_init_start() and
// pango_glyph_item_iter_init_end(), for forward and backward iteration
// respectively, and walked over using any desired mixture of
// pango_glyph_item_iter_next_cluster() and
// pango_glyph_item_iter_prev_cluster().  A common idiom for doing a
// forward iteration over the clusters is:
// <programlisting>
// PangoGlyphItemIter cluster_iter;
// gboolean have_cluster;
//
// for (have_cluster = pango_glyph_item_iter_init_start (&amp;cluster_iter,
// glyph_item, text);
// have_cluster;
// have_cluster = pango_glyph_item_iter_next_cluster (&amp;cluster_iter))
// {
// ...
// }
// </programlisting>
//
// Note that @text is the start of the text for layout, which is then
// indexed by <literal>@glyph_item->item->offset</literal> to get to the
// text of @glyph_item.  The @start_index and @end_index values can directly
// index into @text.  The @start_glyph, @end_glyph, @start_char, and @end_char
// values however are zero-based for the @glyph_item.  For each cluster, the
// item pointed at by the start variables is included in the cluster while
// the one pointed at by end variables is not.
//
// None of the members of a #PangoGlyphItemIter should be modified manually.
/*

C record/class : PangoGlyphItemIter
*/
type GlyphItemIter struct {
	native *C.PangoGlyphItemIter
	// glyph_item : record
	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

func GlyphItemIterNewFromC(u unsafe.Pointer) *GlyphItemIter {
	c := (*C.PangoGlyphItemIter)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItemIter{
		EndChar:    (int32)(c.end_char),
		EndGlyph:   (int32)(c.end_glyph),
		EndIndex:   (int32)(c.end_index),
		StartChar:  (int32)(c.start_char),
		StartGlyph: (int32)(c.start_glyph),
		StartIndex: (int32)(c.start_index),
		Text:       C.GoString(c.text),
		native:     c,
	}

	return g
}

func (recv *GlyphItemIter) ToC() unsafe.Pointer {
	recv.native.text =
		C.CString(recv.Text)
	recv.native.start_glyph =
		(C.int)(recv.StartGlyph)
	recv.native.start_index =
		(C.int)(recv.StartIndex)
	recv.native.start_char =
		(C.int)(recv.StartChar)
	recv.native.end_glyph =
		(C.int)(recv.EndGlyph)
	recv.native.end_index =
		(C.int)(recv.EndIndex)
	recv.native.end_char =
		(C.int)(recv.EndChar)

	return (unsafe.Pointer)(recv.native)
}

// Make a shallow copy of an existing #PangoGlyphItemIter structure.
/*

C function : pango_glyph_item_iter_copy
*/
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	retC := C.pango_glyph_item_iter_copy((*C.PangoGlyphItemIter)(recv.native))
	var retGo (*GlyphItemIter)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GlyphItemIterNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Frees a #PangoGlyphItemIter created by pango_glyph_item_iter_copy().
/*

C function : pango_glyph_item_iter_free
*/
func (recv *GlyphItemIter) Free() {
	C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(recv.native))

	return
}

// Initializes a #PangoGlyphItemIter structure to point to the
// last cluster in a glyph item.
// See #PangoGlyphItemIter for details of cluster orders.
/*

C function : pango_glyph_item_iter_init_end
*/
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_end((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// Initializes a #PangoGlyphItemIter structure to point to the
// first cluster in a glyph item.
// See #PangoGlyphItemIter for details of cluster orders.
/*

C function : pango_glyph_item_iter_init_start
*/
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_start((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// Advances the iterator to the next cluster in the glyph item.
// See #PangoGlyphItemIter for details of cluster orders.
/*

C function : pango_glyph_item_iter_next_cluster
*/
func (recv *GlyphItemIter) NextCluster() bool {
	retC := C.pango_glyph_item_iter_next_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves the iterator to the preceding cluster in the glyph item.
// See #PangoGlyphItemIter for details of cluster orders.
/*

C function : pango_glyph_item_iter_prev_cluster
*/
func (recv *GlyphItemIter) PrevCluster() bool {
	retC := C.pango_glyph_item_iter_prev_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_language_get_scripts : no return type
