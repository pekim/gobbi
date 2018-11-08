// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_target_list_new : unsupported parameter targets :

// BackwardVisibleCursorPosition is a wrapper around the C function gtk_text_iter_backward_visible_cursor_position.
func (recv *TextIter) BackwardVisibleCursorPosition() bool {
	retC := C.gtk_text_iter_backward_visible_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleCursorPositions is a wrapper around the C function gtk_text_iter_backward_visible_cursor_positions.
func (recv *TextIter) BackwardVisibleCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleWordStart is a wrapper around the C function gtk_text_iter_backward_visible_word_start.
func (recv *TextIter) BackwardVisibleWordStart() bool {
	retC := C.gtk_text_iter_backward_visible_word_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleWordStarts is a wrapper around the C function gtk_text_iter_backward_visible_word_starts.
func (recv *TextIter) BackwardVisibleWordStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_word_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleCursorPosition is a wrapper around the C function gtk_text_iter_forward_visible_cursor_position.
func (recv *TextIter) ForwardVisibleCursorPosition() bool {
	retC := C.gtk_text_iter_forward_visible_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleCursorPositions is a wrapper around the C function gtk_text_iter_forward_visible_cursor_positions.
func (recv *TextIter) ForwardVisibleCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleWordEnd is a wrapper around the C function gtk_text_iter_forward_visible_word_end.
func (recv *TextIter) ForwardVisibleWordEnd() bool {
	retC := C.gtk_text_iter_forward_visible_word_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleWordEnds is a wrapper around the C function gtk_text_iter_forward_visible_word_ends.
func (recv *TextIter) ForwardVisibleWordEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_word_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model
