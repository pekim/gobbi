// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/gtk"
	"runtime"
	"sync"
)

var bufferStruct *gi.Struct
var bufferStruct_Once sync.Once

func bufferStruct_Set() error {
	var err error
	bufferStruct_Once.Do(func() {
		bufferStruct, err = gi.StructNew("GtkSource", "Buffer")
	})
	return err
}

type Buffer struct {
	gtk.TextBuffer
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Gtk.TextBuffer'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Gtk.TextBuffer'

// FieldPriv returns the C field 'priv'.
func (recv *Buffer) FieldPriv() *BufferPrivate {
	argValue := gi.FieldGet(bufferStruct, recv.Native, "priv")
	value := &BufferPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Buffer) SetFieldPriv(value *BufferPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(bufferStruct, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_buffer_new' : parameter 'table' of type 'Gtk.TextTagTable' not supported

var bufferNewWithLanguageFunction *gi.Function
var bufferNewWithLanguageFunction_Once sync.Once

func bufferNewWithLanguageFunction_Set() error {
	var err error
	bufferNewWithLanguageFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferNewWithLanguageFunction, err = bufferStruct.InvokerNew("new_with_language")
	})
	return err
}

// BufferNewWithLanguage is a representation of the C type gtk_source_buffer_new_with_language.
func BufferNewWithLanguage(language *Language) *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(language.Native)

	var ret gi.Argument

	err := bufferNewWithLanguageFunction_Set()
	if err == nil {
		ret = bufferNewWithLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_backward_iter_to_source_mark' : parameter 'iter' of type 'Gtk.TextIter' not supported

var bufferBeginNotUndoableActionFunction *gi.Function
var bufferBeginNotUndoableActionFunction_Once sync.Once

func bufferBeginNotUndoableActionFunction_Set() error {
	var err error
	bufferBeginNotUndoableActionFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferBeginNotUndoableActionFunction, err = bufferStruct.InvokerNew("begin_not_undoable_action")
	})
	return err
}

// BeginNotUndoableAction is a representation of the C type gtk_source_buffer_begin_not_undoable_action.
func (recv *Buffer) BeginNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := bufferBeginNotUndoableActionFunction_Set()
	if err == nil {
		bufferBeginNotUndoableActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferCanRedoFunction *gi.Function
var bufferCanRedoFunction_Once sync.Once

func bufferCanRedoFunction_Set() error {
	var err error
	bufferCanRedoFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferCanRedoFunction, err = bufferStruct.InvokerNew("can_redo")
	})
	return err
}

// CanRedo is a representation of the C type gtk_source_buffer_can_redo.
func (recv *Buffer) CanRedo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferCanRedoFunction_Set()
	if err == nil {
		ret = bufferCanRedoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferCanUndoFunction *gi.Function
var bufferCanUndoFunction_Once sync.Once

func bufferCanUndoFunction_Set() error {
	var err error
	bufferCanUndoFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferCanUndoFunction, err = bufferStruct.InvokerNew("can_undo")
	})
	return err
}

// CanUndo is a representation of the C type gtk_source_buffer_can_undo.
func (recv *Buffer) CanUndo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferCanUndoFunction_Set()
	if err == nil {
		ret = bufferCanUndoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_change_case' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_create_source_mark' : parameter 'where' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_create_source_tag' : parameter '...' of type 'nil' not supported

var bufferEndNotUndoableActionFunction *gi.Function
var bufferEndNotUndoableActionFunction_Once sync.Once

func bufferEndNotUndoableActionFunction_Set() error {
	var err error
	bufferEndNotUndoableActionFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferEndNotUndoableActionFunction, err = bufferStruct.InvokerNew("end_not_undoable_action")
	})
	return err
}

// EndNotUndoableAction is a representation of the C type gtk_source_buffer_end_not_undoable_action.
func (recv *Buffer) EndNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := bufferEndNotUndoableActionFunction_Set()
	if err == nil {
		bufferEndNotUndoableActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_buffer_ensure_highlight' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_forward_iter_to_source_mark' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_get_context_classes_at_iter' : parameter 'iter' of type 'Gtk.TextIter' not supported

var bufferGetHighlightMatchingBracketsFunction *gi.Function
var bufferGetHighlightMatchingBracketsFunction_Once sync.Once

func bufferGetHighlightMatchingBracketsFunction_Set() error {
	var err error
	bufferGetHighlightMatchingBracketsFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetHighlightMatchingBracketsFunction, err = bufferStruct.InvokerNew("get_highlight_matching_brackets")
	})
	return err
}

// GetHighlightMatchingBrackets is a representation of the C type gtk_source_buffer_get_highlight_matching_brackets.
func (recv *Buffer) GetHighlightMatchingBrackets() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetHighlightMatchingBracketsFunction_Set()
	if err == nil {
		ret = bufferGetHighlightMatchingBracketsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferGetHighlightSyntaxFunction *gi.Function
var bufferGetHighlightSyntaxFunction_Once sync.Once

func bufferGetHighlightSyntaxFunction_Set() error {
	var err error
	bufferGetHighlightSyntaxFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetHighlightSyntaxFunction, err = bufferStruct.InvokerNew("get_highlight_syntax")
	})
	return err
}

// GetHighlightSyntax is a representation of the C type gtk_source_buffer_get_highlight_syntax.
func (recv *Buffer) GetHighlightSyntax() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetHighlightSyntaxFunction_Set()
	if err == nil {
		ret = bufferGetHighlightSyntaxFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferGetImplicitTrailingNewlineFunction *gi.Function
var bufferGetImplicitTrailingNewlineFunction_Once sync.Once

func bufferGetImplicitTrailingNewlineFunction_Set() error {
	var err error
	bufferGetImplicitTrailingNewlineFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetImplicitTrailingNewlineFunction, err = bufferStruct.InvokerNew("get_implicit_trailing_newline")
	})
	return err
}

// GetImplicitTrailingNewline is a representation of the C type gtk_source_buffer_get_implicit_trailing_newline.
func (recv *Buffer) GetImplicitTrailingNewline() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetImplicitTrailingNewlineFunction_Set()
	if err == nil {
		ret = bufferGetImplicitTrailingNewlineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferGetLanguageFunction *gi.Function
var bufferGetLanguageFunction_Once sync.Once

func bufferGetLanguageFunction_Set() error {
	var err error
	bufferGetLanguageFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetLanguageFunction, err = bufferStruct.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type gtk_source_buffer_get_language.
func (recv *Buffer) GetLanguage() *Language {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetLanguageFunction_Set()
	if err == nil {
		ret = bufferGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Language{}
	retGo.Native = ret.Pointer()

	return retGo
}

var bufferGetMaxUndoLevelsFunction *gi.Function
var bufferGetMaxUndoLevelsFunction_Once sync.Once

func bufferGetMaxUndoLevelsFunction_Set() error {
	var err error
	bufferGetMaxUndoLevelsFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetMaxUndoLevelsFunction, err = bufferStruct.InvokerNew("get_max_undo_levels")
	})
	return err
}

// GetMaxUndoLevels is a representation of the C type gtk_source_buffer_get_max_undo_levels.
func (recv *Buffer) GetMaxUndoLevels() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetMaxUndoLevelsFunction_Set()
	if err == nil {
		ret = bufferGetMaxUndoLevelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_get_source_marks_at_iter' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_get_source_marks_at_line' : return type 'GLib.SList' not supported

var bufferGetStyleSchemeFunction *gi.Function
var bufferGetStyleSchemeFunction_Once sync.Once

func bufferGetStyleSchemeFunction_Set() error {
	var err error
	bufferGetStyleSchemeFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferGetStyleSchemeFunction, err = bufferStruct.InvokerNew("get_style_scheme")
	})
	return err
}

// GetStyleScheme is a representation of the C type gtk_source_buffer_get_style_scheme.
func (recv *Buffer) GetStyleScheme() *StyleScheme {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := bufferGetStyleSchemeFunction_Set()
	if err == nil {
		ret = bufferGetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StyleScheme{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_get_undo_manager' : return type 'UndoManager' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_iter_backward_to_context_class_toggle' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_iter_forward_to_context_class_toggle' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_iter_has_context_class' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_join_lines' : parameter 'start' of type 'Gtk.TextIter' not supported

var bufferRedoFunction *gi.Function
var bufferRedoFunction_Once sync.Once

func bufferRedoFunction_Set() error {
	var err error
	bufferRedoFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferRedoFunction, err = bufferStruct.InvokerNew("redo")
	})
	return err
}

// Redo is a representation of the C type gtk_source_buffer_redo.
func (recv *Buffer) Redo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := bufferRedoFunction_Set()
	if err == nil {
		bufferRedoFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_buffer_remove_source_marks' : parameter 'start' of type 'Gtk.TextIter' not supported

var bufferSetHighlightMatchingBracketsFunction *gi.Function
var bufferSetHighlightMatchingBracketsFunction_Once sync.Once

func bufferSetHighlightMatchingBracketsFunction_Set() error {
	var err error
	bufferSetHighlightMatchingBracketsFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetHighlightMatchingBracketsFunction, err = bufferStruct.InvokerNew("set_highlight_matching_brackets")
	})
	return err
}

// SetHighlightMatchingBrackets is a representation of the C type gtk_source_buffer_set_highlight_matching_brackets.
func (recv *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(highlight)

	err := bufferSetHighlightMatchingBracketsFunction_Set()
	if err == nil {
		bufferSetHighlightMatchingBracketsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetHighlightSyntaxFunction *gi.Function
var bufferSetHighlightSyntaxFunction_Once sync.Once

func bufferSetHighlightSyntaxFunction_Set() error {
	var err error
	bufferSetHighlightSyntaxFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetHighlightSyntaxFunction, err = bufferStruct.InvokerNew("set_highlight_syntax")
	})
	return err
}

// SetHighlightSyntax is a representation of the C type gtk_source_buffer_set_highlight_syntax.
func (recv *Buffer) SetHighlightSyntax(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(highlight)

	err := bufferSetHighlightSyntaxFunction_Set()
	if err == nil {
		bufferSetHighlightSyntaxFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetImplicitTrailingNewlineFunction *gi.Function
var bufferSetImplicitTrailingNewlineFunction_Once sync.Once

func bufferSetImplicitTrailingNewlineFunction_Set() error {
	var err error
	bufferSetImplicitTrailingNewlineFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetImplicitTrailingNewlineFunction, err = bufferStruct.InvokerNew("set_implicit_trailing_newline")
	})
	return err
}

// SetImplicitTrailingNewline is a representation of the C type gtk_source_buffer_set_implicit_trailing_newline.
func (recv *Buffer) SetImplicitTrailingNewline(implicitTrailingNewline bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(implicitTrailingNewline)

	err := bufferSetImplicitTrailingNewlineFunction_Set()
	if err == nil {
		bufferSetImplicitTrailingNewlineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetLanguageFunction *gi.Function
var bufferSetLanguageFunction_Once sync.Once

func bufferSetLanguageFunction_Set() error {
	var err error
	bufferSetLanguageFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetLanguageFunction, err = bufferStruct.InvokerNew("set_language")
	})
	return err
}

// SetLanguage is a representation of the C type gtk_source_buffer_set_language.
func (recv *Buffer) SetLanguage(language *Language) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(language.Native)

	err := bufferSetLanguageFunction_Set()
	if err == nil {
		bufferSetLanguageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetMaxUndoLevelsFunction *gi.Function
var bufferSetMaxUndoLevelsFunction_Once sync.Once

func bufferSetMaxUndoLevelsFunction_Set() error {
	var err error
	bufferSetMaxUndoLevelsFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetMaxUndoLevelsFunction, err = bufferStruct.InvokerNew("set_max_undo_levels")
	})
	return err
}

// SetMaxUndoLevels is a representation of the C type gtk_source_buffer_set_max_undo_levels.
func (recv *Buffer) SetMaxUndoLevels(maxUndoLevels int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(maxUndoLevels)

	err := bufferSetMaxUndoLevelsFunction_Set()
	if err == nil {
		bufferSetMaxUndoLevelsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetStyleSchemeFunction *gi.Function
var bufferSetStyleSchemeFunction_Once sync.Once

func bufferSetStyleSchemeFunction_Set() error {
	var err error
	bufferSetStyleSchemeFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferSetStyleSchemeFunction, err = bufferStruct.InvokerNew("set_style_scheme")
	})
	return err
}

// SetStyleScheme is a representation of the C type gtk_source_buffer_set_style_scheme.
func (recv *Buffer) SetStyleScheme(scheme *StyleScheme) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(scheme.Native)

	err := bufferSetStyleSchemeFunction_Set()
	if err == nil {
		bufferSetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_buffer_set_undo_manager' : parameter 'manager' of type 'UndoManager' not supported

// UNSUPPORTED : C value 'gtk_source_buffer_sort_lines' : parameter 'start' of type 'Gtk.TextIter' not supported

var bufferUndoFunction *gi.Function
var bufferUndoFunction_Once sync.Once

func bufferUndoFunction_Set() error {
	var err error
	bufferUndoFunction_Once.Do(func() {
		err = bufferStruct_Set()
		if err != nil {
			return
		}
		bufferUndoFunction, err = bufferStruct.InvokerNew("undo")
	})
	return err
}

// Undo is a representation of the C type gtk_source_buffer_undo.
func (recv *Buffer) Undo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := bufferUndoFunction_Set()
	if err == nil {
		bufferUndoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionStruct *gi.Struct
var completionStruct_Once sync.Once

func completionStruct_Set() error {
	var err error
	completionStruct_Once.Do(func() {
		completionStruct, err = gi.StructNew("GtkSource", "Completion")
	})
	return err
}

type Completion struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Completion) FieldPriv() *CompletionPrivate {
	argValue := gi.FieldGet(completionStruct, recv.Native, "priv")
	value := &CompletionPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Completion) SetFieldPriv(value *CompletionPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(completionStruct, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_add_provider' : parameter 'provider' of type 'CompletionProvider' not supported

var completionBlockInteractiveFunction *gi.Function
var completionBlockInteractiveFunction_Once sync.Once

func completionBlockInteractiveFunction_Set() error {
	var err error
	completionBlockInteractiveFunction_Once.Do(func() {
		err = completionStruct_Set()
		if err != nil {
			return
		}
		completionBlockInteractiveFunction, err = completionStruct.InvokerNew("block_interactive")
	})
	return err
}

// BlockInteractive is a representation of the C type gtk_source_completion_block_interactive.
func (recv *Completion) BlockInteractive() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := completionBlockInteractiveFunction_Set()
	if err == nil {
		completionBlockInteractiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_completion_create_context' : parameter 'position' of type 'Gtk.TextIter' not supported

var completionGetInfoWindowFunction *gi.Function
var completionGetInfoWindowFunction_Once sync.Once

func completionGetInfoWindowFunction_Set() error {
	var err error
	completionGetInfoWindowFunction_Once.Do(func() {
		err = completionStruct_Set()
		if err != nil {
			return
		}
		completionGetInfoWindowFunction, err = completionStruct.InvokerNew("get_info_window")
	})
	return err
}

// GetInfoWindow is a representation of the C type gtk_source_completion_get_info_window.
func (recv *Completion) GetInfoWindow() *CompletionInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := completionGetInfoWindowFunction_Set()
	if err == nil {
		ret = completionGetInfoWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &CompletionInfo{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_get_providers' : return type 'GLib.List' not supported

var completionGetViewFunction *gi.Function
var completionGetViewFunction_Once sync.Once

func completionGetViewFunction_Set() error {
	var err error
	completionGetViewFunction_Once.Do(func() {
		err = completionStruct_Set()
		if err != nil {
			return
		}
		completionGetViewFunction, err = completionStruct.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_completion_get_view.
func (recv *Completion) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := completionGetViewFunction_Set()
	if err == nil {
		ret = completionGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &View{}
	retGo.Native = ret.Pointer()

	return retGo
}

var completionHideFunction *gi.Function
var completionHideFunction_Once sync.Once

func completionHideFunction_Set() error {
	var err error
	completionHideFunction_Once.Do(func() {
		err = completionStruct_Set()
		if err != nil {
			return
		}
		completionHideFunction, err = completionStruct.InvokerNew("hide")
	})
	return err
}

// Hide is a representation of the C type gtk_source_completion_hide.
func (recv *Completion) Hide() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := completionHideFunction_Set()
	if err == nil {
		completionHideFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_completion_move_window' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_completion_remove_provider' : parameter 'provider' of type 'CompletionProvider' not supported

// UNSUPPORTED : C value 'gtk_source_completion_show' : parameter 'providers' of type 'GLib.List' not supported

var completionUnblockInteractiveFunction *gi.Function
var completionUnblockInteractiveFunction_Once sync.Once

func completionUnblockInteractiveFunction_Set() error {
	var err error
	completionUnblockInteractiveFunction_Once.Do(func() {
		err = completionStruct_Set()
		if err != nil {
			return
		}
		completionUnblockInteractiveFunction, err = completionStruct.InvokerNew("unblock_interactive")
	})
	return err
}

// UnblockInteractive is a representation of the C type gtk_source_completion_unblock_interactive.
func (recv *Completion) UnblockInteractive() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := completionUnblockInteractiveFunction_Set()
	if err == nil {
		completionUnblockInteractiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

// CompletionStruct creates an uninitialised Completion.
func CompletionStruct() *Completion {
	err := completionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Completion{}
	structGo.Native = completionStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCompletion)
	return structGo
}
func finalizeCompletion(obj *Completion) {
	completionStruct.Free(obj.Native)
}

var completionContextStruct *gi.Struct
var completionContextStruct_Once sync.Once

func completionContextStruct_Set() error {
	var err error
	completionContextStruct_Once.Do(func() {
		completionContextStruct, err = gi.StructNew("GtkSource", "CompletionContext")
	})
	return err
}

type CompletionContext struct {
	gobject.InitiallyUnowned
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.InitiallyUnowned'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.InitiallyUnowned'

// FieldPriv returns the C field 'priv'.
func (recv *CompletionContext) FieldPriv() *CompletionContextPrivate {
	argValue := gi.FieldGet(completionContextStruct, recv.Native, "priv")
	value := &CompletionContextPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionContext) SetFieldPriv(value *CompletionContextPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(completionContextStruct, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_context_add_proposals' : parameter 'provider' of type 'CompletionProvider' not supported

// UNSUPPORTED : C value 'gtk_source_completion_context_get_activation' : return type 'CompletionActivation' not supported

// UNSUPPORTED : C value 'gtk_source_completion_context_get_iter' : parameter 'iter' of type 'Gtk.TextIter' not supported

// CompletionContextStruct creates an uninitialised CompletionContext.
func CompletionContextStruct() *CompletionContext {
	err := completionContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionContext{}
	structGo.Native = completionContextStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeCompletionContext)
	return structGo
}
func finalizeCompletionContext(obj *CompletionContext) {
	completionContextStruct.Free(obj.Native)
}

var completionInfoStruct *gi.Struct
var completionInfoStruct_Once sync.Once

func completionInfoStruct_Set() error {
	var err error
	completionInfoStruct_Once.Do(func() {
		completionInfoStruct, err = gi.StructNew("GtkSource", "CompletionInfo")
	})
	return err
}

type CompletionInfo struct {
	gtk.Window
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.Window'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.Window'

// FieldPriv returns the C field 'priv'.
func (recv *CompletionInfo) FieldPriv() *CompletionInfoPrivate {
	argValue := gi.FieldGet(completionInfoStruct, recv.Native, "priv")
	value := &CompletionInfoPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionInfo) SetFieldPriv(value *CompletionInfoPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(completionInfoStruct, recv.Native, "priv", argValue)
}

var completionInfoNewFunction *gi.Function
var completionInfoNewFunction_Once sync.Once

func completionInfoNewFunction_Set() error {
	var err error
	completionInfoNewFunction_Once.Do(func() {
		err = completionInfoStruct_Set()
		if err != nil {
			return
		}
		completionInfoNewFunction, err = completionInfoStruct.InvokerNew("new")
	})
	return err
}

// CompletionInfoNew is a representation of the C type gtk_source_completion_info_new.
func CompletionInfoNew() *CompletionInfo {

	var ret gi.Argument

	err := completionInfoNewFunction_Set()
	if err == nil {
		ret = completionInfoNewFunction.Invoke(nil, nil)
	}

	retGo := &CompletionInfo{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_info_get_widget' : return type 'Gtk.Widget' not supported

// UNSUPPORTED : C value 'gtk_source_completion_info_move_to_iter' : parameter 'view' of type 'Gtk.TextView' not supported

// UNSUPPORTED : C value 'gtk_source_completion_info_set_widget' : parameter 'widget' of type 'Gtk.Widget' not supported

var completionItemStruct *gi.Struct
var completionItemStruct_Once sync.Once

func completionItemStruct_Set() error {
	var err error
	completionItemStruct_Once.Do(func() {
		completionItemStruct, err = gi.StructNew("GtkSource", "CompletionItem")
	})
	return err
}

type CompletionItem struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *CompletionItem) FieldPriv() *CompletionItemPrivate {
	argValue := gi.FieldGet(completionItemStruct, recv.Native, "priv")
	value := &CompletionItemPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionItem) SetFieldPriv(value *CompletionItemPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(completionItemStruct, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_item_new' : parameter 'icon' of type 'GdkPixbuf.Pixbuf' not supported

var completionItemNewFromStockFunction *gi.Function
var completionItemNewFromStockFunction_Once sync.Once

func completionItemNewFromStockFunction_Set() error {
	var err error
	completionItemNewFromStockFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemNewFromStockFunction, err = completionItemStruct.InvokerNew("new_from_stock")
	})
	return err
}

// CompletionItemNewFromStock is a representation of the C type gtk_source_completion_item_new_from_stock.
func CompletionItemNewFromStock(label string, text string, stock string, info string) *CompletionItem {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(label)
	inArgs[1].SetString(text)
	inArgs[2].SetString(stock)
	inArgs[3].SetString(info)

	var ret gi.Argument

	err := completionItemNewFromStockFunction_Set()
	if err == nil {
		ret = completionItemNewFromStockFunction.Invoke(inArgs[:], nil)
	}

	retGo := &CompletionItem{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_item_new_with_markup' : parameter 'icon' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_source_completion_item_set_gicon' : parameter 'gicon' of type 'Gio.Icon' not supported

// UNSUPPORTED : C value 'gtk_source_completion_item_set_icon' : parameter 'icon' of type 'GdkPixbuf.Pixbuf' not supported

var completionItemSetIconNameFunction *gi.Function
var completionItemSetIconNameFunction_Once sync.Once

func completionItemSetIconNameFunction_Set() error {
	var err error
	completionItemSetIconNameFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemSetIconNameFunction, err = completionItemStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_completion_item_set_icon_name.
func (recv *CompletionItem) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(iconName)

	err := completionItemSetIconNameFunction_Set()
	if err == nil {
		completionItemSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionItemSetInfoFunction *gi.Function
var completionItemSetInfoFunction_Once sync.Once

func completionItemSetInfoFunction_Set() error {
	var err error
	completionItemSetInfoFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemSetInfoFunction, err = completionItemStruct.InvokerNew("set_info")
	})
	return err
}

// SetInfo is a representation of the C type gtk_source_completion_item_set_info.
func (recv *CompletionItem) SetInfo(info string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(info)

	err := completionItemSetInfoFunction_Set()
	if err == nil {
		completionItemSetInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionItemSetLabelFunction *gi.Function
var completionItemSetLabelFunction_Once sync.Once

func completionItemSetLabelFunction_Set() error {
	var err error
	completionItemSetLabelFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemSetLabelFunction, err = completionItemStruct.InvokerNew("set_label")
	})
	return err
}

// SetLabel is a representation of the C type gtk_source_completion_item_set_label.
func (recv *CompletionItem) SetLabel(label string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(label)

	err := completionItemSetLabelFunction_Set()
	if err == nil {
		completionItemSetLabelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionItemSetMarkupFunction *gi.Function
var completionItemSetMarkupFunction_Once sync.Once

func completionItemSetMarkupFunction_Set() error {
	var err error
	completionItemSetMarkupFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemSetMarkupFunction, err = completionItemStruct.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type gtk_source_completion_item_set_markup.
func (recv *CompletionItem) SetMarkup(markup string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(markup)

	err := completionItemSetMarkupFunction_Set()
	if err == nil {
		completionItemSetMarkupFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionItemSetTextFunction *gi.Function
var completionItemSetTextFunction_Once sync.Once

func completionItemSetTextFunction_Set() error {
	var err error
	completionItemSetTextFunction_Once.Do(func() {
		err = completionItemStruct_Set()
		if err != nil {
			return
		}
		completionItemSetTextFunction, err = completionItemStruct.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type gtk_source_completion_item_set_text.
func (recv *CompletionItem) SetText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(text)

	err := completionItemSetTextFunction_Set()
	if err == nil {
		completionItemSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionWordsStruct *gi.Struct
var completionWordsStruct_Once sync.Once

func completionWordsStruct_Set() error {
	var err error
	completionWordsStruct_Once.Do(func() {
		completionWordsStruct, err = gi.StructNew("GtkSource", "CompletionWords")
	})
	return err
}

type CompletionWords struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *CompletionWords) FieldPriv() *CompletionWordsPrivate {
	argValue := gi.FieldGet(completionWordsStruct, recv.Native, "priv")
	value := &CompletionWordsPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionWords) SetFieldPriv(value *CompletionWordsPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(completionWordsStruct, recv.Native, "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_words_new' : parameter 'icon' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_source_completion_words_register' : parameter 'buffer' of type 'Gtk.TextBuffer' not supported

// UNSUPPORTED : C value 'gtk_source_completion_words_unregister' : parameter 'buffer' of type 'Gtk.TextBuffer' not supported

var fileStruct *gi.Struct
var fileStruct_Once sync.Once

func fileStruct_Set() error {
	var err error
	fileStruct_Once.Do(func() {
		fileStruct, err = gi.StructNew("GtkSource", "File")
	})
	return err
}

type File struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *File) FieldPriv() *FilePrivate {
	argValue := gi.FieldGet(fileStruct, recv.Native, "priv")
	value := &FilePrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *File) SetFieldPriv(value *FilePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(fileStruct, recv.Native, "priv", argValue)
}

var fileNewFunction *gi.Function
var fileNewFunction_Once sync.Once

func fileNewFunction_Set() error {
	var err error
	fileNewFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileNewFunction, err = fileStruct.InvokerNew("new")
	})
	return err
}

// FileNew is a representation of the C type gtk_source_file_new.
func FileNew() *File {

	var ret gi.Argument

	err := fileNewFunction_Set()
	if err == nil {
		ret = fileNewFunction.Invoke(nil, nil)
	}

	retGo := &File{}
	retGo.Native = ret.Pointer()

	return retGo
}

var fileCheckFileOnDiskFunction *gi.Function
var fileCheckFileOnDiskFunction_Once sync.Once

func fileCheckFileOnDiskFunction_Set() error {
	var err error
	fileCheckFileOnDiskFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileCheckFileOnDiskFunction, err = fileStruct.InvokerNew("check_file_on_disk")
	})
	return err
}

// CheckFileOnDisk is a representation of the C type gtk_source_file_check_file_on_disk.
func (recv *File) CheckFileOnDisk() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := fileCheckFileOnDiskFunction_Set()
	if err == nil {
		fileCheckFileOnDiskFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileGetCompressionTypeFunction *gi.Function
var fileGetCompressionTypeFunction_Once sync.Once

func fileGetCompressionTypeFunction_Set() error {
	var err error
	fileGetCompressionTypeFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileGetCompressionTypeFunction, err = fileStruct.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_get_compression_type.
func (recv *File) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileGetCompressionTypeFunction_Set()
	if err == nil {
		ret = fileGetCompressionTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompressionType(ret.Int32())

	return retGo
}

var fileGetEncodingFunction *gi.Function
var fileGetEncodingFunction_Once sync.Once

func fileGetEncodingFunction_Set() error {
	var err error
	fileGetEncodingFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileGetEncodingFunction, err = fileStruct.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_get_encoding.
func (recv *File) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileGetEncodingFunction_Set()
	if err == nil {
		ret = fileGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Encoding{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_get_location' : return type 'Gio.File' not supported

var fileGetNewlineTypeFunction *gi.Function
var fileGetNewlineTypeFunction_Once sync.Once

func fileGetNewlineTypeFunction_Set() error {
	var err error
	fileGetNewlineTypeFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileGetNewlineTypeFunction, err = fileStruct.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_get_newline_type.
func (recv *File) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileGetNewlineTypeFunction_Set()
	if err == nil {
		ret = fileGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NewlineType(ret.Int32())

	return retGo
}

var fileIsDeletedFunction *gi.Function
var fileIsDeletedFunction_Once sync.Once

func fileIsDeletedFunction_Set() error {
	var err error
	fileIsDeletedFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileIsDeletedFunction, err = fileStruct.InvokerNew("is_deleted")
	})
	return err
}

// IsDeleted is a representation of the C type gtk_source_file_is_deleted.
func (recv *File) IsDeleted() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileIsDeletedFunction_Set()
	if err == nil {
		ret = fileIsDeletedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileIsExternallyModifiedFunction *gi.Function
var fileIsExternallyModifiedFunction_Once sync.Once

func fileIsExternallyModifiedFunction_Set() error {
	var err error
	fileIsExternallyModifiedFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileIsExternallyModifiedFunction, err = fileStruct.InvokerNew("is_externally_modified")
	})
	return err
}

// IsExternallyModified is a representation of the C type gtk_source_file_is_externally_modified.
func (recv *File) IsExternallyModified() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileIsExternallyModifiedFunction_Set()
	if err == nil {
		ret = fileIsExternallyModifiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileIsLocalFunction *gi.Function
var fileIsLocalFunction_Once sync.Once

func fileIsLocalFunction_Set() error {
	var err error
	fileIsLocalFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileIsLocalFunction, err = fileStruct.InvokerNew("is_local")
	})
	return err
}

// IsLocal is a representation of the C type gtk_source_file_is_local.
func (recv *File) IsLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileIsLocalFunction_Set()
	if err == nil {
		ret = fileIsLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileIsReadonlyFunction *gi.Function
var fileIsReadonlyFunction_Once sync.Once

func fileIsReadonlyFunction_Set() error {
	var err error
	fileIsReadonlyFunction_Once.Do(func() {
		err = fileStruct_Set()
		if err != nil {
			return
		}
		fileIsReadonlyFunction, err = fileStruct.InvokerNew("is_readonly")
	})
	return err
}

// IsReadonly is a representation of the C type gtk_source_file_is_readonly.
func (recv *File) IsReadonly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileIsReadonlyFunction_Set()
	if err == nil {
		ret = fileIsReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_set_location' : parameter 'location' of type 'Gio.File' not supported

// UNSUPPORTED : C value 'gtk_source_file_set_mount_operation_factory' : parameter 'callback' of type 'MountOperationFactory' not supported

var fileLoaderStruct *gi.Struct
var fileLoaderStruct_Once sync.Once

func fileLoaderStruct_Set() error {
	var err error
	fileLoaderStruct_Once.Do(func() {
		fileLoaderStruct, err = gi.StructNew("GtkSource", "FileLoader")
	})
	return err
}

type FileLoader struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *FileLoader) FieldPriv() *FileLoaderPrivate {
	argValue := gi.FieldGet(fileLoaderStruct, recv.Native, "priv")
	value := &FileLoaderPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *FileLoader) SetFieldPriv(value *FileLoaderPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(fileLoaderStruct, recv.Native, "priv", argValue)
}

var fileLoaderNewFunction *gi.Function
var fileLoaderNewFunction_Once sync.Once

func fileLoaderNewFunction_Set() error {
	var err error
	fileLoaderNewFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderNewFunction, err = fileLoaderStruct.InvokerNew("new")
	})
	return err
}

// FileLoaderNew is a representation of the C type gtk_source_file_loader_new.
func FileLoaderNew(buffer *Buffer, file *File) *FileLoader {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native)
	inArgs[1].SetPointer(file.Native)

	var ret gi.Argument

	err := fileLoaderNewFunction_Set()
	if err == nil {
		ret = fileLoaderNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileLoader{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_new_from_stream' : parameter 'stream' of type 'Gio.InputStream' not supported

var fileLoaderGetBufferFunction *gi.Function
var fileLoaderGetBufferFunction_Once sync.Once

func fileLoaderGetBufferFunction_Set() error {
	var err error
	fileLoaderGetBufferFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderGetBufferFunction, err = fileLoaderStruct.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_file_loader_get_buffer.
func (recv *FileLoader) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileLoaderGetBufferFunction_Set()
	if err == nil {
		ret = fileLoaderGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var fileLoaderGetCompressionTypeFunction *gi.Function
var fileLoaderGetCompressionTypeFunction_Once sync.Once

func fileLoaderGetCompressionTypeFunction_Set() error {
	var err error
	fileLoaderGetCompressionTypeFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderGetCompressionTypeFunction, err = fileLoaderStruct.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_loader_get_compression_type.
func (recv *FileLoader) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileLoaderGetCompressionTypeFunction_Set()
	if err == nil {
		ret = fileLoaderGetCompressionTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompressionType(ret.Int32())

	return retGo
}

var fileLoaderGetEncodingFunction *gi.Function
var fileLoaderGetEncodingFunction_Once sync.Once

func fileLoaderGetEncodingFunction_Set() error {
	var err error
	fileLoaderGetEncodingFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderGetEncodingFunction, err = fileLoaderStruct.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_loader_get_encoding.
func (recv *FileLoader) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileLoaderGetEncodingFunction_Set()
	if err == nil {
		ret = fileLoaderGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Encoding{}
	retGo.Native = ret.Pointer()

	return retGo
}

var fileLoaderGetFileFunction *gi.Function
var fileLoaderGetFileFunction_Once sync.Once

func fileLoaderGetFileFunction_Set() error {
	var err error
	fileLoaderGetFileFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderGetFileFunction, err = fileLoaderStruct.InvokerNew("get_file")
	})
	return err
}

// GetFile is a representation of the C type gtk_source_file_loader_get_file.
func (recv *FileLoader) GetFile() *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileLoaderGetFileFunction_Set()
	if err == nil {
		ret = fileLoaderGetFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := &File{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_get_input_stream' : return type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gtk_source_file_loader_get_location' : return type 'Gio.File' not supported

var fileLoaderGetNewlineTypeFunction *gi.Function
var fileLoaderGetNewlineTypeFunction_Once sync.Once

func fileLoaderGetNewlineTypeFunction_Set() error {
	var err error
	fileLoaderGetNewlineTypeFunction_Once.Do(func() {
		err = fileLoaderStruct_Set()
		if err != nil {
			return
		}
		fileLoaderGetNewlineTypeFunction, err = fileLoaderStruct.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_loader_get_newline_type.
func (recv *FileLoader) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileLoaderGetNewlineTypeFunction_Set()
	if err == nil {
		ret = fileLoaderGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NewlineType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_load_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'gtk_source_file_loader_load_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_file_loader_set_candidate_encodings' : parameter 'candidate_encodings' of type 'GLib.SList' not supported

var fileSaverStruct *gi.Struct
var fileSaverStruct_Once sync.Once

func fileSaverStruct_Set() error {
	var err error
	fileSaverStruct_Once.Do(func() {
		fileSaverStruct, err = gi.StructNew("GtkSource", "FileSaver")
	})
	return err
}

type FileSaver struct {
	gobject.Object
}

// UNSUPPORTED : C value 'object' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'object' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *FileSaver) FieldPriv() *FileSaverPrivate {
	argValue := gi.FieldGet(fileSaverStruct, recv.Native, "priv")
	value := &FileSaverPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *FileSaver) SetFieldPriv(value *FileSaverPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(fileSaverStruct, recv.Native, "priv", argValue)
}

var fileSaverNewFunction *gi.Function
var fileSaverNewFunction_Once sync.Once

func fileSaverNewFunction_Set() error {
	var err error
	fileSaverNewFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverNewFunction, err = fileSaverStruct.InvokerNew("new")
	})
	return err
}

// FileSaverNew is a representation of the C type gtk_source_file_saver_new.
func FileSaverNew(buffer *Buffer, file *File) *FileSaver {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native)
	inArgs[1].SetPointer(file.Native)

	var ret gi.Argument

	err := fileSaverNewFunction_Set()
	if err == nil {
		ret = fileSaverNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FileSaver{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_new_with_target' : parameter 'target_location' of type 'Gio.File' not supported

var fileSaverGetBufferFunction *gi.Function
var fileSaverGetBufferFunction_Once sync.Once

func fileSaverGetBufferFunction_Set() error {
	var err error
	fileSaverGetBufferFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverGetBufferFunction, err = fileSaverStruct.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_file_saver_get_buffer.
func (recv *FileSaver) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileSaverGetBufferFunction_Set()
	if err == nil {
		ret = fileSaverGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var fileSaverGetCompressionTypeFunction *gi.Function
var fileSaverGetCompressionTypeFunction_Once sync.Once

func fileSaverGetCompressionTypeFunction_Set() error {
	var err error
	fileSaverGetCompressionTypeFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverGetCompressionTypeFunction, err = fileSaverStruct.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_saver_get_compression_type.
func (recv *FileSaver) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileSaverGetCompressionTypeFunction_Set()
	if err == nil {
		ret = fileSaverGetCompressionTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompressionType(ret.Int32())

	return retGo
}

var fileSaverGetEncodingFunction *gi.Function
var fileSaverGetEncodingFunction_Once sync.Once

func fileSaverGetEncodingFunction_Set() error {
	var err error
	fileSaverGetEncodingFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverGetEncodingFunction, err = fileSaverStruct.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_saver_get_encoding.
func (recv *FileSaver) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileSaverGetEncodingFunction_Set()
	if err == nil {
		ret = fileSaverGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Encoding{}
	retGo.Native = ret.Pointer()

	return retGo
}

var fileSaverGetFileFunction *gi.Function
var fileSaverGetFileFunction_Once sync.Once

func fileSaverGetFileFunction_Set() error {
	var err error
	fileSaverGetFileFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverGetFileFunction, err = fileSaverStruct.InvokerNew("get_file")
	})
	return err
}

// GetFile is a representation of the C type gtk_source_file_saver_get_file.
func (recv *FileSaver) GetFile() *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileSaverGetFileFunction_Set()
	if err == nil {
		ret = fileSaverGetFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := &File{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_get_flags' : return type 'FileSaverFlags' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_get_location' : return type 'Gio.File' not supported

var fileSaverGetNewlineTypeFunction *gi.Function
var fileSaverGetNewlineTypeFunction_Once sync.Once

func fileSaverGetNewlineTypeFunction_Set() error {
	var err error
	fileSaverGetNewlineTypeFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverGetNewlineTypeFunction, err = fileSaverStruct.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_saver_get_newline_type.
func (recv *FileSaver) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := fileSaverGetNewlineTypeFunction_Set()
	if err == nil {
		ret = fileSaverGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NewlineType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_save_async' : parameter 'cancellable' of type 'Gio.Cancellable' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_save_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var fileSaverSetCompressionTypeFunction *gi.Function
var fileSaverSetCompressionTypeFunction_Once sync.Once

func fileSaverSetCompressionTypeFunction_Set() error {
	var err error
	fileSaverSetCompressionTypeFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverSetCompressionTypeFunction, err = fileSaverStruct.InvokerNew("set_compression_type")
	})
	return err
}

// SetCompressionType is a representation of the C type gtk_source_file_saver_set_compression_type.
func (recv *FileSaver) SetCompressionType(compressionType CompressionType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(int32(compressionType))

	err := fileSaverSetCompressionTypeFunction_Set()
	if err == nil {
		fileSaverSetCompressionTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileSaverSetEncodingFunction *gi.Function
var fileSaverSetEncodingFunction_Once sync.Once

func fileSaverSetEncodingFunction_Set() error {
	var err error
	fileSaverSetEncodingFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverSetEncodingFunction, err = fileSaverStruct.InvokerNew("set_encoding")
	})
	return err
}

// SetEncoding is a representation of the C type gtk_source_file_saver_set_encoding.
func (recv *FileSaver) SetEncoding(encoding *Encoding) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(encoding.Native)

	err := fileSaverSetEncodingFunction_Set()
	if err == nil {
		fileSaverSetEncodingFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_file_saver_set_flags' : parameter 'flags' of type 'FileSaverFlags' not supported

var fileSaverSetNewlineTypeFunction *gi.Function
var fileSaverSetNewlineTypeFunction_Once sync.Once

func fileSaverSetNewlineTypeFunction_Set() error {
	var err error
	fileSaverSetNewlineTypeFunction_Once.Do(func() {
		err = fileSaverStruct_Set()
		if err != nil {
			return
		}
		fileSaverSetNewlineTypeFunction, err = fileSaverStruct.InvokerNew("set_newline_type")
	})
	return err
}

// SetNewlineType is a representation of the C type gtk_source_file_saver_set_newline_type.
func (recv *FileSaver) SetNewlineType(newlineType NewlineType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(int32(newlineType))

	err := fileSaverSetNewlineTypeFunction_Set()
	if err == nil {
		fileSaverSetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterStruct *gi.Struct
var gutterStruct_Once sync.Once

func gutterStruct_Set() error {
	var err error
	gutterStruct_Once.Do(func() {
		gutterStruct, err = gi.StructNew("GtkSource", "Gutter")
	})
	return err
}

type Gutter struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Gutter) FieldPriv() *GutterPrivate {
	argValue := gi.FieldGet(gutterStruct, recv.Native, "priv")
	value := &GutterPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Gutter) SetFieldPriv(value *GutterPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(gutterStruct, recv.Native, "priv", argValue)
}

var gutterGetPaddingFunction *gi.Function
var gutterGetPaddingFunction_Once sync.Once

func gutterGetPaddingFunction_Set() error {
	var err error
	gutterGetPaddingFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterGetPaddingFunction, err = gutterStruct.InvokerNew("get_padding")
	})
	return err
}

// GetPadding is a representation of the C type gtk_source_gutter_get_padding.
func (recv *Gutter) GetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(xpad)
	inArgs[2].SetInt32(ypad)

	err := gutterGetPaddingFunction_Set()
	if err == nil {
		gutterGetPaddingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterGetRendererAtPosFunction *gi.Function
var gutterGetRendererAtPosFunction_Once sync.Once

func gutterGetRendererAtPosFunction_Set() error {
	var err error
	gutterGetRendererAtPosFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterGetRendererAtPosFunction, err = gutterStruct.InvokerNew("get_renderer_at_pos")
	})
	return err
}

// GetRendererAtPos is a representation of the C type gtk_source_gutter_get_renderer_at_pos.
func (recv *Gutter) GetRendererAtPos(x int32, y int32) *GutterRenderer {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var ret gi.Argument

	err := gutterGetRendererAtPosFunction_Set()
	if err == nil {
		ret = gutterGetRendererAtPosFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GutterRenderer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var gutterGetViewFunction *gi.Function
var gutterGetViewFunction_Once sync.Once

func gutterGetViewFunction_Set() error {
	var err error
	gutterGetViewFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterGetViewFunction, err = gutterStruct.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_gutter_get_view.
func (recv *Gutter) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterGetViewFunction_Set()
	if err == nil {
		ret = gutterGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &View{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_get_window' : return type 'Gdk.Window' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_get_window_type' : return type 'Gtk.TextWindowType' not supported

var gutterInsertFunction *gi.Function
var gutterInsertFunction_Once sync.Once

func gutterInsertFunction_Set() error {
	var err error
	gutterInsertFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterInsertFunction, err = gutterStruct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type gtk_source_gutter_insert.
func (recv *Gutter) Insert(renderer *GutterRenderer, position int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(renderer.Native)
	inArgs[2].SetInt32(position)

	var ret gi.Argument

	err := gutterInsertFunction_Set()
	if err == nil {
		ret = gutterInsertFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gutterQueueDrawFunction *gi.Function
var gutterQueueDrawFunction_Once sync.Once

func gutterQueueDrawFunction_Set() error {
	var err error
	gutterQueueDrawFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterQueueDrawFunction, err = gutterStruct.InvokerNew("queue_draw")
	})
	return err
}

// QueueDraw is a representation of the C type gtk_source_gutter_queue_draw.
func (recv *Gutter) QueueDraw() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := gutterQueueDrawFunction_Set()
	if err == nil {
		gutterQueueDrawFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRemoveFunction *gi.Function
var gutterRemoveFunction_Once sync.Once

func gutterRemoveFunction_Set() error {
	var err error
	gutterRemoveFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterRemoveFunction, err = gutterStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type gtk_source_gutter_remove.
func (recv *Gutter) Remove(renderer *GutterRenderer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(renderer.Native)

	err := gutterRemoveFunction_Set()
	if err == nil {
		gutterRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterReorderFunction *gi.Function
var gutterReorderFunction_Once sync.Once

func gutterReorderFunction_Set() error {
	var err error
	gutterReorderFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterReorderFunction, err = gutterStruct.InvokerNew("reorder")
	})
	return err
}

// Reorder is a representation of the C type gtk_source_gutter_reorder.
func (recv *Gutter) Reorder(renderer *GutterRenderer, position int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(renderer.Native)
	inArgs[2].SetInt32(position)

	err := gutterReorderFunction_Set()
	if err == nil {
		gutterReorderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterSetPaddingFunction *gi.Function
var gutterSetPaddingFunction_Once sync.Once

func gutterSetPaddingFunction_Set() error {
	var err error
	gutterSetPaddingFunction_Once.Do(func() {
		err = gutterStruct_Set()
		if err != nil {
			return
		}
		gutterSetPaddingFunction, err = gutterStruct.InvokerNew("set_padding")
	})
	return err
}

// SetPadding is a representation of the C type gtk_source_gutter_set_padding.
func (recv *Gutter) SetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(xpad)
	inArgs[2].SetInt32(ypad)

	err := gutterSetPaddingFunction_Set()
	if err == nil {
		gutterSetPaddingFunction.Invoke(inArgs[:], nil)
	}

	return
}

// GutterStruct creates an uninitialised Gutter.
func GutterStruct() *Gutter {
	err := gutterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Gutter{}
	structGo.Native = gutterStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeGutter)
	return structGo
}
func finalizeGutter(obj *Gutter) {
	gutterStruct.Free(obj.Native)
}

var gutterRendererStruct *gi.Struct
var gutterRendererStruct_Once sync.Once

func gutterRendererStruct_Set() error {
	var err error
	gutterRendererStruct_Once.Do(func() {
		gutterRendererStruct, err = gi.StructNew("GtkSource", "GutterRenderer")
	})
	return err
}

type GutterRenderer struct {
	gobject.InitiallyUnowned
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.InitiallyUnowned'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.InitiallyUnowned'

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_activate' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_begin' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_draw' : parameter 'cr' of type 'cairo.Context' not supported

var gutterRendererEndFunction *gi.Function
var gutterRendererEndFunction_Once sync.Once

func gutterRendererEndFunction_Set() error {
	var err error
	gutterRendererEndFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererEndFunction, err = gutterRendererStruct.InvokerNew("end")
	})
	return err
}

// End is a representation of the C type gtk_source_gutter_renderer_end.
func (recv *GutterRenderer) End() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := gutterRendererEndFunction_Set()
	if err == nil {
		gutterRendererEndFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererGetAlignmentFunction *gi.Function
var gutterRendererGetAlignmentFunction_Once sync.Once

func gutterRendererGetAlignmentFunction_Set() error {
	var err error
	gutterRendererGetAlignmentFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererGetAlignmentFunction, err = gutterRendererStruct.InvokerNew("get_alignment")
	})
	return err
}

// GetAlignment is a representation of the C type gtk_source_gutter_renderer_get_alignment.
func (recv *GutterRenderer) GetAlignment() (float32, float32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := gutterRendererGetAlignmentFunction_Set()
	if err == nil {
		gutterRendererGetAlignmentFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float32()
	out1 := outArgs[1].Float32()

	return out0, out1
}

var gutterRendererGetAlignmentModeFunction *gi.Function
var gutterRendererGetAlignmentModeFunction_Once sync.Once

func gutterRendererGetAlignmentModeFunction_Set() error {
	var err error
	gutterRendererGetAlignmentModeFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererGetAlignmentModeFunction, err = gutterRendererStruct.InvokerNew("get_alignment_mode")
	})
	return err
}

// GetAlignmentMode is a representation of the C type gtk_source_gutter_renderer_get_alignment_mode.
func (recv *GutterRenderer) GetAlignmentMode() GutterRendererAlignmentMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterRendererGetAlignmentModeFunction_Set()
	if err == nil {
		ret = gutterRendererGetAlignmentModeFunction.Invoke(inArgs[:], nil)
	}

	retGo := GutterRendererAlignmentMode(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_get_background' : parameter 'color' of type 'Gdk.RGBA' not supported

var gutterRendererGetPaddingFunction *gi.Function
var gutterRendererGetPaddingFunction_Once sync.Once

func gutterRendererGetPaddingFunction_Set() error {
	var err error
	gutterRendererGetPaddingFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererGetPaddingFunction, err = gutterRendererStruct.InvokerNew("get_padding")
	})
	return err
}

// GetPadding is a representation of the C type gtk_source_gutter_renderer_get_padding.
func (recv *GutterRenderer) GetPadding() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := gutterRendererGetPaddingFunction_Set()
	if err == nil {
		gutterRendererGetPaddingFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var gutterRendererGetSizeFunction *gi.Function
var gutterRendererGetSizeFunction_Once sync.Once

func gutterRendererGetSizeFunction_Set() error {
	var err error
	gutterRendererGetSizeFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererGetSizeFunction, err = gutterRendererStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type gtk_source_gutter_renderer_get_size.
func (recv *GutterRenderer) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterRendererGetSizeFunction_Set()
	if err == nil {
		ret = gutterRendererGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_get_view' : return type 'Gtk.TextView' not supported

var gutterRendererGetVisibleFunction *gi.Function
var gutterRendererGetVisibleFunction_Once sync.Once

func gutterRendererGetVisibleFunction_Set() error {
	var err error
	gutterRendererGetVisibleFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererGetVisibleFunction, err = gutterRendererStruct.InvokerNew("get_visible")
	})
	return err
}

// GetVisible is a representation of the C type gtk_source_gutter_renderer_get_visible.
func (recv *GutterRenderer) GetVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterRendererGetVisibleFunction_Set()
	if err == nil {
		ret = gutterRendererGetVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_get_window_type' : return type 'Gtk.TextWindowType' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_query_activatable' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_query_data' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_query_tooltip' : parameter 'iter' of type 'Gtk.TextIter' not supported

var gutterRendererQueueDrawFunction *gi.Function
var gutterRendererQueueDrawFunction_Once sync.Once

func gutterRendererQueueDrawFunction_Set() error {
	var err error
	gutterRendererQueueDrawFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererQueueDrawFunction, err = gutterRendererStruct.InvokerNew("queue_draw")
	})
	return err
}

// QueueDraw is a representation of the C type gtk_source_gutter_renderer_queue_draw.
func (recv *GutterRenderer) QueueDraw() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := gutterRendererQueueDrawFunction_Set()
	if err == nil {
		gutterRendererQueueDrawFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetAlignmentFunction *gi.Function
var gutterRendererSetAlignmentFunction_Once sync.Once

func gutterRendererSetAlignmentFunction_Set() error {
	var err error
	gutterRendererSetAlignmentFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererSetAlignmentFunction, err = gutterRendererStruct.InvokerNew("set_alignment")
	})
	return err
}

// SetAlignment is a representation of the C type gtk_source_gutter_renderer_set_alignment.
func (recv *GutterRenderer) SetAlignment(xalign float32, yalign float32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetFloat32(xalign)
	inArgs[2].SetFloat32(yalign)

	err := gutterRendererSetAlignmentFunction_Set()
	if err == nil {
		gutterRendererSetAlignmentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetAlignmentModeFunction *gi.Function
var gutterRendererSetAlignmentModeFunction_Once sync.Once

func gutterRendererSetAlignmentModeFunction_Set() error {
	var err error
	gutterRendererSetAlignmentModeFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererSetAlignmentModeFunction, err = gutterRendererStruct.InvokerNew("set_alignment_mode")
	})
	return err
}

// SetAlignmentMode is a representation of the C type gtk_source_gutter_renderer_set_alignment_mode.
func (recv *GutterRenderer) SetAlignmentMode(mode GutterRendererAlignmentMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(int32(mode))

	err := gutterRendererSetAlignmentModeFunction_Set()
	if err == nil {
		gutterRendererSetAlignmentModeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_set_background' : parameter 'color' of type 'Gdk.RGBA' not supported

var gutterRendererSetPaddingFunction *gi.Function
var gutterRendererSetPaddingFunction_Once sync.Once

func gutterRendererSetPaddingFunction_Set() error {
	var err error
	gutterRendererSetPaddingFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererSetPaddingFunction, err = gutterRendererStruct.InvokerNew("set_padding")
	})
	return err
}

// SetPadding is a representation of the C type gtk_source_gutter_renderer_set_padding.
func (recv *GutterRenderer) SetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(xpad)
	inArgs[2].SetInt32(ypad)

	err := gutterRendererSetPaddingFunction_Set()
	if err == nil {
		gutterRendererSetPaddingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetSizeFunction *gi.Function
var gutterRendererSetSizeFunction_Once sync.Once

func gutterRendererSetSizeFunction_Set() error {
	var err error
	gutterRendererSetSizeFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererSetSizeFunction, err = gutterRendererStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type gtk_source_gutter_renderer_set_size.
func (recv *GutterRenderer) SetSize(size int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(size)

	err := gutterRendererSetSizeFunction_Set()
	if err == nil {
		gutterRendererSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetVisibleFunction *gi.Function
var gutterRendererSetVisibleFunction_Once sync.Once

func gutterRendererSetVisibleFunction_Set() error {
	var err error
	gutterRendererSetVisibleFunction_Once.Do(func() {
		err = gutterRendererStruct_Set()
		if err != nil {
			return
		}
		gutterRendererSetVisibleFunction, err = gutterRendererStruct.InvokerNew("set_visible")
	})
	return err
}

// SetVisible is a representation of the C type gtk_source_gutter_renderer_set_visible.
func (recv *GutterRenderer) SetVisible(visible bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(visible)

	err := gutterRendererSetVisibleFunction_Set()
	if err == nil {
		gutterRendererSetVisibleFunction.Invoke(inArgs[:], nil)
	}

	return
}

// GutterRendererStruct creates an uninitialised GutterRenderer.
func GutterRendererStruct() *GutterRenderer {
	err := gutterRendererStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRenderer{}
	structGo.Native = gutterRendererStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeGutterRenderer)
	return structGo
}
func finalizeGutterRenderer(obj *GutterRenderer) {
	gutterRendererStruct.Free(obj.Native)
}

var gutterRendererPixbufStruct *gi.Struct
var gutterRendererPixbufStruct_Once sync.Once

func gutterRendererPixbufStruct_Set() error {
	var err error
	gutterRendererPixbufStruct_Once.Do(func() {
		gutterRendererPixbufStruct, err = gi.StructNew("GtkSource", "GutterRendererPixbuf")
	})
	return err
}

type GutterRendererPixbuf struct {
	GutterRenderer
}

var gutterRendererPixbufNewFunction *gi.Function
var gutterRendererPixbufNewFunction_Once sync.Once

func gutterRendererPixbufNewFunction_Set() error {
	var err error
	gutterRendererPixbufNewFunction_Once.Do(func() {
		err = gutterRendererPixbufStruct_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufNewFunction, err = gutterRendererPixbufStruct.InvokerNew("new")
	})
	return err
}

// GutterRendererPixbufNew is a representation of the C type gtk_source_gutter_renderer_pixbuf_new.
func GutterRendererPixbufNew() *GutterRendererPixbuf {

	var ret gi.Argument

	err := gutterRendererPixbufNewFunction_Set()
	if err == nil {
		ret = gutterRendererPixbufNewFunction.Invoke(nil, nil)
	}

	retGo := &GutterRendererPixbuf{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_pixbuf_get_gicon' : return type 'Gio.Icon' not supported

var gutterRendererPixbufGetIconNameFunction *gi.Function
var gutterRendererPixbufGetIconNameFunction_Once sync.Once

func gutterRendererPixbufGetIconNameFunction_Set() error {
	var err error
	gutterRendererPixbufGetIconNameFunction_Once.Do(func() {
		err = gutterRendererPixbufStruct_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufGetIconNameFunction, err = gutterRendererPixbufStruct.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_gutter_renderer_pixbuf_get_icon_name.
func (recv *GutterRendererPixbuf) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterRendererPixbufGetIconNameFunction_Set()
	if err == nil {
		ret = gutterRendererPixbufGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_pixbuf_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

var gutterRendererPixbufGetStockIdFunction *gi.Function
var gutterRendererPixbufGetStockIdFunction_Once sync.Once

func gutterRendererPixbufGetStockIdFunction_Set() error {
	var err error
	gutterRendererPixbufGetStockIdFunction_Once.Do(func() {
		err = gutterRendererPixbufStruct_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufGetStockIdFunction, err = gutterRendererPixbufStruct.InvokerNew("get_stock_id")
	})
	return err
}

// GetStockId is a representation of the C type gtk_source_gutter_renderer_pixbuf_get_stock_id.
func (recv *GutterRendererPixbuf) GetStockId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gutterRendererPixbufGetStockIdFunction_Set()
	if err == nil {
		ret = gutterRendererPixbufGetStockIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_pixbuf_set_gicon' : parameter 'icon' of type 'Gio.Icon' not supported

var gutterRendererPixbufSetIconNameFunction *gi.Function
var gutterRendererPixbufSetIconNameFunction_Once sync.Once

func gutterRendererPixbufSetIconNameFunction_Set() error {
	var err error
	gutterRendererPixbufSetIconNameFunction_Once.Do(func() {
		err = gutterRendererPixbufStruct_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufSetIconNameFunction, err = gutterRendererPixbufStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_gutter_renderer_pixbuf_set_icon_name.
func (recv *GutterRendererPixbuf) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(iconName)

	err := gutterRendererPixbufSetIconNameFunction_Set()
	if err == nil {
		gutterRendererPixbufSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_pixbuf_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

var gutterRendererPixbufSetStockIdFunction *gi.Function
var gutterRendererPixbufSetStockIdFunction_Once sync.Once

func gutterRendererPixbufSetStockIdFunction_Set() error {
	var err error
	gutterRendererPixbufSetStockIdFunction_Once.Do(func() {
		err = gutterRendererPixbufStruct_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufSetStockIdFunction, err = gutterRendererPixbufStruct.InvokerNew("set_stock_id")
	})
	return err
}

// SetStockId is a representation of the C type gtk_source_gutter_renderer_pixbuf_set_stock_id.
func (recv *GutterRendererPixbuf) SetStockId(stockId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(stockId)

	err := gutterRendererPixbufSetStockIdFunction_Set()
	if err == nil {
		gutterRendererPixbufSetStockIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererTextStruct *gi.Struct
var gutterRendererTextStruct_Once sync.Once

func gutterRendererTextStruct_Set() error {
	var err error
	gutterRendererTextStruct_Once.Do(func() {
		gutterRendererTextStruct, err = gi.StructNew("GtkSource", "GutterRendererText")
	})
	return err
}

type GutterRendererText struct {
	GutterRenderer
}

var gutterRendererTextNewFunction *gi.Function
var gutterRendererTextNewFunction_Once sync.Once

func gutterRendererTextNewFunction_Set() error {
	var err error
	gutterRendererTextNewFunction_Once.Do(func() {
		err = gutterRendererTextStruct_Set()
		if err != nil {
			return
		}
		gutterRendererTextNewFunction, err = gutterRendererTextStruct.InvokerNew("new")
	})
	return err
}

// GutterRendererTextNew is a representation of the C type gtk_source_gutter_renderer_text_new.
func GutterRendererTextNew() *GutterRendererText {

	var ret gi.Argument

	err := gutterRendererTextNewFunction_Set()
	if err == nil {
		ret = gutterRendererTextNewFunction.Invoke(nil, nil)
	}

	retGo := &GutterRendererText{}
	retGo.Native = ret.Pointer()

	return retGo
}

var gutterRendererTextMeasureFunction *gi.Function
var gutterRendererTextMeasureFunction_Once sync.Once

func gutterRendererTextMeasureFunction_Set() error {
	var err error
	gutterRendererTextMeasureFunction_Once.Do(func() {
		err = gutterRendererTextStruct_Set()
		if err != nil {
			return
		}
		gutterRendererTextMeasureFunction, err = gutterRendererTextStruct.InvokerNew("measure")
	})
	return err
}

// Measure is a representation of the C type gtk_source_gutter_renderer_text_measure.
func (recv *GutterRendererText) Measure(text string) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(text)

	var outArgs [2]gi.Argument

	err := gutterRendererTextMeasureFunction_Set()
	if err == nil {
		gutterRendererTextMeasureFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var gutterRendererTextMeasureMarkupFunction *gi.Function
var gutterRendererTextMeasureMarkupFunction_Once sync.Once

func gutterRendererTextMeasureMarkupFunction_Set() error {
	var err error
	gutterRendererTextMeasureMarkupFunction_Once.Do(func() {
		err = gutterRendererTextStruct_Set()
		if err != nil {
			return
		}
		gutterRendererTextMeasureMarkupFunction, err = gutterRendererTextStruct.InvokerNew("measure_markup")
	})
	return err
}

// MeasureMarkup is a representation of the C type gtk_source_gutter_renderer_text_measure_markup.
func (recv *GutterRendererText) MeasureMarkup(markup string) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(markup)

	var outArgs [2]gi.Argument

	err := gutterRendererTextMeasureMarkupFunction_Set()
	if err == nil {
		gutterRendererTextMeasureMarkupFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var gutterRendererTextSetMarkupFunction *gi.Function
var gutterRendererTextSetMarkupFunction_Once sync.Once

func gutterRendererTextSetMarkupFunction_Set() error {
	var err error
	gutterRendererTextSetMarkupFunction_Once.Do(func() {
		err = gutterRendererTextStruct_Set()
		if err != nil {
			return
		}
		gutterRendererTextSetMarkupFunction, err = gutterRendererTextStruct.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type gtk_source_gutter_renderer_text_set_markup.
func (recv *GutterRendererText) SetMarkup(markup string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(markup)
	inArgs[2].SetInt32(length)

	err := gutterRendererTextSetMarkupFunction_Set()
	if err == nil {
		gutterRendererTextSetMarkupFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererTextSetTextFunction *gi.Function
var gutterRendererTextSetTextFunction_Once sync.Once

func gutterRendererTextSetTextFunction_Set() error {
	var err error
	gutterRendererTextSetTextFunction_Once.Do(func() {
		err = gutterRendererTextStruct_Set()
		if err != nil {
			return
		}
		gutterRendererTextSetTextFunction, err = gutterRendererTextStruct.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type gtk_source_gutter_renderer_text_set_text.
func (recv *GutterRendererText) SetText(text string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)

	err := gutterRendererTextSetTextFunction_Set()
	if err == nil {
		gutterRendererTextSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageStruct *gi.Struct
var languageStruct_Once sync.Once

func languageStruct_Set() error {
	var err error
	languageStruct_Once.Do(func() {
		languageStruct, err = gi.StructNew("GtkSource", "Language")
	})
	return err
}

type Language struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *Language) FieldPriv() *LanguagePrivate {
	argValue := gi.FieldGet(languageStruct, recv.Native, "priv")
	value := &LanguagePrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Language) SetFieldPriv(value *LanguagePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(languageStruct, recv.Native, "priv", argValue)
}

var languageGetGlobsFunction *gi.Function
var languageGetGlobsFunction_Once sync.Once

func languageGetGlobsFunction_Set() error {
	var err error
	languageGetGlobsFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetGlobsFunction, err = languageStruct.InvokerNew("get_globs")
	})
	return err
}

// GetGlobs is a representation of the C type gtk_source_language_get_globs.
func (recv *Language) GetGlobs() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := languageGetGlobsFunction_Set()
	if err == nil {
		languageGetGlobsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageGetHiddenFunction *gi.Function
var languageGetHiddenFunction_Once sync.Once

func languageGetHiddenFunction_Set() error {
	var err error
	languageGetHiddenFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetHiddenFunction, err = languageStruct.InvokerNew("get_hidden")
	})
	return err
}

// GetHidden is a representation of the C type gtk_source_language_get_hidden.
func (recv *Language) GetHidden() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := languageGetHiddenFunction_Set()
	if err == nil {
		ret = languageGetHiddenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var languageGetIdFunction *gi.Function
var languageGetIdFunction_Once sync.Once

func languageGetIdFunction_Set() error {
	var err error
	languageGetIdFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetIdFunction, err = languageStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type gtk_source_language_get_id.
func (recv *Language) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := languageGetIdFunction_Set()
	if err == nil {
		ret = languageGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetMetadataFunction *gi.Function
var languageGetMetadataFunction_Once sync.Once

func languageGetMetadataFunction_Set() error {
	var err error
	languageGetMetadataFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetMetadataFunction, err = languageStruct.InvokerNew("get_metadata")
	})
	return err
}

// GetMetadata is a representation of the C type gtk_source_language_get_metadata.
func (recv *Language) GetMetadata(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := languageGetMetadataFunction_Set()
	if err == nil {
		ret = languageGetMetadataFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetMimeTypesFunction *gi.Function
var languageGetMimeTypesFunction_Once sync.Once

func languageGetMimeTypesFunction_Set() error {
	var err error
	languageGetMimeTypesFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetMimeTypesFunction, err = languageStruct.InvokerNew("get_mime_types")
	})
	return err
}

// GetMimeTypes is a representation of the C type gtk_source_language_get_mime_types.
func (recv *Language) GetMimeTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := languageGetMimeTypesFunction_Set()
	if err == nil {
		languageGetMimeTypesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageGetNameFunction *gi.Function
var languageGetNameFunction_Once sync.Once

func languageGetNameFunction_Set() error {
	var err error
	languageGetNameFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetNameFunction, err = languageStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_language_get_name.
func (recv *Language) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := languageGetNameFunction_Set()
	if err == nil {
		ret = languageGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetSectionFunction *gi.Function
var languageGetSectionFunction_Once sync.Once

func languageGetSectionFunction_Set() error {
	var err error
	languageGetSectionFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetSectionFunction, err = languageStruct.InvokerNew("get_section")
	})
	return err
}

// GetSection is a representation of the C type gtk_source_language_get_section.
func (recv *Language) GetSection() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := languageGetSectionFunction_Set()
	if err == nil {
		ret = languageGetSectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetStyleFallbackFunction *gi.Function
var languageGetStyleFallbackFunction_Once sync.Once

func languageGetStyleFallbackFunction_Set() error {
	var err error
	languageGetStyleFallbackFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetStyleFallbackFunction, err = languageStruct.InvokerNew("get_style_fallback")
	})
	return err
}

// GetStyleFallback is a representation of the C type gtk_source_language_get_style_fallback.
func (recv *Language) GetStyleFallback(styleId string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(styleId)

	var ret gi.Argument

	err := languageGetStyleFallbackFunction_Set()
	if err == nil {
		ret = languageGetStyleFallbackFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetStyleIdsFunction *gi.Function
var languageGetStyleIdsFunction_Once sync.Once

func languageGetStyleIdsFunction_Set() error {
	var err error
	languageGetStyleIdsFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetStyleIdsFunction, err = languageStruct.InvokerNew("get_style_ids")
	})
	return err
}

// GetStyleIds is a representation of the C type gtk_source_language_get_style_ids.
func (recv *Language) GetStyleIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := languageGetStyleIdsFunction_Set()
	if err == nil {
		languageGetStyleIdsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageGetStyleNameFunction *gi.Function
var languageGetStyleNameFunction_Once sync.Once

func languageGetStyleNameFunction_Set() error {
	var err error
	languageGetStyleNameFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetStyleNameFunction, err = languageStruct.InvokerNew("get_style_name")
	})
	return err
}

// GetStyleName is a representation of the C type gtk_source_language_get_style_name.
func (recv *Language) GetStyleName(styleId string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(styleId)

	var ret gi.Argument

	err := languageGetStyleNameFunction_Set()
	if err == nil {
		ret = languageGetStyleNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// LanguageStruct creates an uninitialised Language.
func LanguageStruct() *Language {
	err := languageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Language{}
	structGo.Native = languageStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeLanguage)
	return structGo
}
func finalizeLanguage(obj *Language) {
	languageStruct.Free(obj.Native)
}

var languageManagerStruct *gi.Struct
var languageManagerStruct_Once sync.Once

func languageManagerStruct_Set() error {
	var err error
	languageManagerStruct_Once.Do(func() {
		languageManagerStruct, err = gi.StructNew("GtkSource", "LanguageManager")
	})
	return err
}

type LanguageManager struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *LanguageManager) FieldPriv() *LanguageManagerPrivate {
	argValue := gi.FieldGet(languageManagerStruct, recv.Native, "priv")
	value := &LanguageManagerPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *LanguageManager) SetFieldPriv(value *LanguageManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(languageManagerStruct, recv.Native, "priv", argValue)
}

var languageManagerNewFunction *gi.Function
var languageManagerNewFunction_Once sync.Once

func languageManagerNewFunction_Set() error {
	var err error
	languageManagerNewFunction_Once.Do(func() {
		err = languageManagerStruct_Set()
		if err != nil {
			return
		}
		languageManagerNewFunction, err = languageManagerStruct.InvokerNew("new")
	})
	return err
}

// LanguageManagerNew is a representation of the C type gtk_source_language_manager_new.
func LanguageManagerNew() *LanguageManager {

	var ret gi.Argument

	err := languageManagerNewFunction_Set()
	if err == nil {
		ret = languageManagerNewFunction.Invoke(nil, nil)
	}

	retGo := &LanguageManager{}
	retGo.Native = ret.Pointer()

	return retGo
}

var languageManagerGetLanguageFunction *gi.Function
var languageManagerGetLanguageFunction_Once sync.Once

func languageManagerGetLanguageFunction_Set() error {
	var err error
	languageManagerGetLanguageFunction_Once.Do(func() {
		err = languageManagerStruct_Set()
		if err != nil {
			return
		}
		languageManagerGetLanguageFunction, err = languageManagerStruct.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type gtk_source_language_manager_get_language.
func (recv *LanguageManager) GetLanguage(id string) *Language {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(id)

	var ret gi.Argument

	err := languageManagerGetLanguageFunction_Set()
	if err == nil {
		ret = languageManagerGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Language{}
	retGo.Native = ret.Pointer()

	return retGo
}

var languageManagerGetLanguageIdsFunction *gi.Function
var languageManagerGetLanguageIdsFunction_Once sync.Once

func languageManagerGetLanguageIdsFunction_Set() error {
	var err error
	languageManagerGetLanguageIdsFunction_Once.Do(func() {
		err = languageManagerStruct_Set()
		if err != nil {
			return
		}
		languageManagerGetLanguageIdsFunction, err = languageManagerStruct.InvokerNew("get_language_ids")
	})
	return err
}

// GetLanguageIds is a representation of the C type gtk_source_language_manager_get_language_ids.
func (recv *LanguageManager) GetLanguageIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := languageManagerGetLanguageIdsFunction_Set()
	if err == nil {
		languageManagerGetLanguageIdsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageManagerGetSearchPathFunction *gi.Function
var languageManagerGetSearchPathFunction_Once sync.Once

func languageManagerGetSearchPathFunction_Set() error {
	var err error
	languageManagerGetSearchPathFunction_Once.Do(func() {
		err = languageManagerStruct_Set()
		if err != nil {
			return
		}
		languageManagerGetSearchPathFunction, err = languageManagerStruct.InvokerNew("get_search_path")
	})
	return err
}

// GetSearchPath is a representation of the C type gtk_source_language_manager_get_search_path.
func (recv *LanguageManager) GetSearchPath() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := languageManagerGetSearchPathFunction_Set()
	if err == nil {
		languageManagerGetSearchPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageManagerGuessLanguageFunction *gi.Function
var languageManagerGuessLanguageFunction_Once sync.Once

func languageManagerGuessLanguageFunction_Set() error {
	var err error
	languageManagerGuessLanguageFunction_Once.Do(func() {
		err = languageManagerStruct_Set()
		if err != nil {
			return
		}
		languageManagerGuessLanguageFunction, err = languageManagerStruct.InvokerNew("guess_language")
	})
	return err
}

// GuessLanguage is a representation of the C type gtk_source_language_manager_guess_language.
func (recv *LanguageManager) GuessLanguage(filename string, contentType string) *Language {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(filename)
	inArgs[2].SetString(contentType)

	var ret gi.Argument

	err := languageManagerGuessLanguageFunction_Set()
	if err == nil {
		ret = languageManagerGuessLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Language{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_language_manager_set_search_path' : parameter 'dirs' of type 'nil' not supported

var mapStruct *gi.Struct
var mapStruct_Once sync.Once

func mapStruct_Set() error {
	var err error
	mapStruct_Once.Do(func() {
		mapStruct, err = gi.StructNew("GtkSource", "Map")
	})
	return err
}

type Map struct {
	View
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Map) FieldParentInstance() *View {
	argValue := gi.FieldGet(mapStruct, recv.Native, "parent_instance")
	value := &View{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Map) SetFieldParentInstance(value *View) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(mapStruct, recv.Native, "parent_instance", argValue)
}

var mapNewFunction *gi.Function
var mapNewFunction_Once sync.Once

func mapNewFunction_Set() error {
	var err error
	mapNewFunction_Once.Do(func() {
		err = mapStruct_Set()
		if err != nil {
			return
		}
		mapNewFunction, err = mapStruct.InvokerNew("new")
	})
	return err
}

// MapNew is a representation of the C type gtk_source_map_new.
func MapNew() *Map {

	var ret gi.Argument

	err := mapNewFunction_Set()
	if err == nil {
		ret = mapNewFunction.Invoke(nil, nil)
	}

	retGo := &Map{}
	retGo.Native = ret.Pointer()

	return retGo
}

var mapGetViewFunction *gi.Function
var mapGetViewFunction_Once sync.Once

func mapGetViewFunction_Set() error {
	var err error
	mapGetViewFunction_Once.Do(func() {
		err = mapStruct_Set()
		if err != nil {
			return
		}
		mapGetViewFunction, err = mapStruct.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_map_get_view.
func (recv *Map) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := mapGetViewFunction_Set()
	if err == nil {
		ret = mapGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &View{}
	retGo.Native = ret.Pointer()

	return retGo
}

var mapSetViewFunction *gi.Function
var mapSetViewFunction_Once sync.Once

func mapSetViewFunction_Set() error {
	var err error
	mapSetViewFunction_Once.Do(func() {
		err = mapStruct_Set()
		if err != nil {
			return
		}
		mapSetViewFunction, err = mapStruct.InvokerNew("set_view")
	})
	return err
}

// SetView is a representation of the C type gtk_source_map_set_view.
func (recv *Map) SetView(view *View) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(view.Native)

	err := mapSetViewFunction_Set()
	if err == nil {
		mapSetViewFunction.Invoke(inArgs[:], nil)
	}

	return
}

var markStruct *gi.Struct
var markStruct_Once sync.Once

func markStruct_Set() error {
	var err error
	markStruct_Once.Do(func() {
		markStruct, err = gi.StructNew("GtkSource", "Mark")
	})
	return err
}

type Mark struct {
	gtk.TextMark
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Gtk.TextMark'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Gtk.TextMark'

// FieldPriv returns the C field 'priv'.
func (recv *Mark) FieldPriv() *MarkPrivate {
	argValue := gi.FieldGet(markStruct, recv.Native, "priv")
	value := &MarkPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Mark) SetFieldPriv(value *MarkPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(markStruct, recv.Native, "priv", argValue)
}

var markNewFunction *gi.Function
var markNewFunction_Once sync.Once

func markNewFunction_Set() error {
	var err error
	markNewFunction_Once.Do(func() {
		err = markStruct_Set()
		if err != nil {
			return
		}
		markNewFunction, err = markStruct.InvokerNew("new")
	})
	return err
}

// MarkNew is a representation of the C type gtk_source_mark_new.
func MarkNew(name string, category string) *Mark {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(category)

	var ret gi.Argument

	err := markNewFunction_Set()
	if err == nil {
		ret = markNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Mark{}
	retGo.Native = ret.Pointer()

	return retGo
}

var markGetCategoryFunction *gi.Function
var markGetCategoryFunction_Once sync.Once

func markGetCategoryFunction_Set() error {
	var err error
	markGetCategoryFunction_Once.Do(func() {
		err = markStruct_Set()
		if err != nil {
			return
		}
		markGetCategoryFunction, err = markStruct.InvokerNew("get_category")
	})
	return err
}

// GetCategory is a representation of the C type gtk_source_mark_get_category.
func (recv *Mark) GetCategory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := markGetCategoryFunction_Set()
	if err == nil {
		ret = markGetCategoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var markNextFunction *gi.Function
var markNextFunction_Once sync.Once

func markNextFunction_Set() error {
	var err error
	markNextFunction_Once.Do(func() {
		err = markStruct_Set()
		if err != nil {
			return
		}
		markNextFunction, err = markStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type gtk_source_mark_next.
func (recv *Mark) Next(category string) *Mark {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(category)

	var ret gi.Argument

	err := markNextFunction_Set()
	if err == nil {
		ret = markNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Mark{}
	retGo.Native = ret.Pointer()

	return retGo
}

var markPrevFunction *gi.Function
var markPrevFunction_Once sync.Once

func markPrevFunction_Set() error {
	var err error
	markPrevFunction_Once.Do(func() {
		err = markStruct_Set()
		if err != nil {
			return
		}
		markPrevFunction, err = markStruct.InvokerNew("prev")
	})
	return err
}

// Prev is a representation of the C type gtk_source_mark_prev.
func (recv *Mark) Prev(category string) *Mark {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(category)

	var ret gi.Argument

	err := markPrevFunction_Set()
	if err == nil {
		ret = markPrevFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Mark{}
	retGo.Native = ret.Pointer()

	return retGo
}

var markAttributesStruct *gi.Struct
var markAttributesStruct_Once sync.Once

func markAttributesStruct_Set() error {
	var err error
	markAttributesStruct_Once.Do(func() {
		markAttributesStruct, err = gi.StructNew("GtkSource", "MarkAttributes")
	})
	return err
}

type MarkAttributes struct {
	gobject.Object
}

var markAttributesNewFunction *gi.Function
var markAttributesNewFunction_Once sync.Once

func markAttributesNewFunction_Set() error {
	var err error
	markAttributesNewFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesNewFunction, err = markAttributesStruct.InvokerNew("new")
	})
	return err
}

// MarkAttributesNew is a representation of the C type gtk_source_mark_attributes_new.
func MarkAttributesNew() *MarkAttributes {

	var ret gi.Argument

	err := markAttributesNewFunction_Set()
	if err == nil {
		ret = markAttributesNewFunction.Invoke(nil, nil)
	}

	retGo := &MarkAttributes{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_get_background' : parameter 'background' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'gtk_source_mark_attributes_get_gicon' : return type 'Gio.Icon' not supported

var markAttributesGetIconNameFunction *gi.Function
var markAttributesGetIconNameFunction_Once sync.Once

func markAttributesGetIconNameFunction_Set() error {
	var err error
	markAttributesGetIconNameFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesGetIconNameFunction, err = markAttributesStruct.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_mark_attributes_get_icon_name.
func (recv *MarkAttributes) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := markAttributesGetIconNameFunction_Set()
	if err == nil {
		ret = markAttributesGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

var markAttributesGetStockIdFunction *gi.Function
var markAttributesGetStockIdFunction_Once sync.Once

func markAttributesGetStockIdFunction_Set() error {
	var err error
	markAttributesGetStockIdFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesGetStockIdFunction, err = markAttributesStruct.InvokerNew("get_stock_id")
	})
	return err
}

// GetStockId is a representation of the C type gtk_source_mark_attributes_get_stock_id.
func (recv *MarkAttributes) GetStockId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := markAttributesGetStockIdFunction_Set()
	if err == nil {
		ret = markAttributesGetStockIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var markAttributesGetTooltipMarkupFunction *gi.Function
var markAttributesGetTooltipMarkupFunction_Once sync.Once

func markAttributesGetTooltipMarkupFunction_Set() error {
	var err error
	markAttributesGetTooltipMarkupFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesGetTooltipMarkupFunction, err = markAttributesStruct.InvokerNew("get_tooltip_markup")
	})
	return err
}

// GetTooltipMarkup is a representation of the C type gtk_source_mark_attributes_get_tooltip_markup.
func (recv *MarkAttributes) GetTooltipMarkup(mark *Mark) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(mark.Native)

	var ret gi.Argument

	err := markAttributesGetTooltipMarkupFunction_Set()
	if err == nil {
		ret = markAttributesGetTooltipMarkupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var markAttributesGetTooltipTextFunction *gi.Function
var markAttributesGetTooltipTextFunction_Once sync.Once

func markAttributesGetTooltipTextFunction_Set() error {
	var err error
	markAttributesGetTooltipTextFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesGetTooltipTextFunction, err = markAttributesStruct.InvokerNew("get_tooltip_text")
	})
	return err
}

// GetTooltipText is a representation of the C type gtk_source_mark_attributes_get_tooltip_text.
func (recv *MarkAttributes) GetTooltipText(mark *Mark) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(mark.Native)

	var ret gi.Argument

	err := markAttributesGetTooltipTextFunction_Set()
	if err == nil {
		ret = markAttributesGetTooltipTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_render_icon' : parameter 'widget' of type 'Gtk.Widget' not supported

// UNSUPPORTED : C value 'gtk_source_mark_attributes_set_background' : parameter 'background' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'gtk_source_mark_attributes_set_gicon' : parameter 'gicon' of type 'Gio.Icon' not supported

var markAttributesSetIconNameFunction *gi.Function
var markAttributesSetIconNameFunction_Once sync.Once

func markAttributesSetIconNameFunction_Set() error {
	var err error
	markAttributesSetIconNameFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesSetIconNameFunction, err = markAttributesStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_mark_attributes_set_icon_name.
func (recv *MarkAttributes) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(iconName)

	err := markAttributesSetIconNameFunction_Set()
	if err == nil {
		markAttributesSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

var markAttributesSetStockIdFunction *gi.Function
var markAttributesSetStockIdFunction_Once sync.Once

func markAttributesSetStockIdFunction_Set() error {
	var err error
	markAttributesSetStockIdFunction_Once.Do(func() {
		err = markAttributesStruct_Set()
		if err != nil {
			return
		}
		markAttributesSetStockIdFunction, err = markAttributesStruct.InvokerNew("set_stock_id")
	})
	return err
}

// SetStockId is a representation of the C type gtk_source_mark_attributes_set_stock_id.
func (recv *MarkAttributes) SetStockId(stockId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(stockId)

	err := markAttributesSetStockIdFunction_Set()
	if err == nil {
		markAttributesSetStockIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorStruct *gi.Struct
var printCompositorStruct_Once sync.Once

func printCompositorStruct_Set() error {
	var err error
	printCompositorStruct_Once.Do(func() {
		printCompositorStruct, err = gi.StructNew("GtkSource", "PrintCompositor")
	})
	return err
}

type PrintCompositor struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *PrintCompositor) FieldPriv() *PrintCompositorPrivate {
	argValue := gi.FieldGet(printCompositorStruct, recv.Native, "priv")
	value := &PrintCompositorPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *PrintCompositor) SetFieldPriv(value *PrintCompositorPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(printCompositorStruct, recv.Native, "priv", argValue)
}

var printCompositorNewFunction *gi.Function
var printCompositorNewFunction_Once sync.Once

func printCompositorNewFunction_Set() error {
	var err error
	printCompositorNewFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorNewFunction, err = printCompositorStruct.InvokerNew("new")
	})
	return err
}

// PrintCompositorNew is a representation of the C type gtk_source_print_compositor_new.
func PrintCompositorNew(buffer *Buffer) *PrintCompositor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(buffer.Native)

	var ret gi.Argument

	err := printCompositorNewFunction_Set()
	if err == nil {
		ret = printCompositorNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PrintCompositor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var printCompositorNewFromViewFunction *gi.Function
var printCompositorNewFromViewFunction_Once sync.Once

func printCompositorNewFromViewFunction_Set() error {
	var err error
	printCompositorNewFromViewFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorNewFromViewFunction, err = printCompositorStruct.InvokerNew("new_from_view")
	})
	return err
}

// PrintCompositorNewFromView is a representation of the C type gtk_source_print_compositor_new_from_view.
func PrintCompositorNewFromView(view *View) *PrintCompositor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(view.Native)

	var ret gi.Argument

	err := printCompositorNewFromViewFunction_Set()
	if err == nil {
		ret = printCompositorNewFromViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PrintCompositor{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_draw_page' : parameter 'context' of type 'Gtk.PrintContext' not supported

var printCompositorGetBodyFontNameFunction *gi.Function
var printCompositorGetBodyFontNameFunction_Once sync.Once

func printCompositorGetBodyFontNameFunction_Set() error {
	var err error
	printCompositorGetBodyFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetBodyFontNameFunction, err = printCompositorStruct.InvokerNew("get_body_font_name")
	})
	return err
}

// GetBodyFontName is a representation of the C type gtk_source_print_compositor_get_body_font_name.
func (recv *PrintCompositor) GetBodyFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetBodyFontNameFunction_Set()
	if err == nil {
		ret = printCompositorGetBodyFontNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_get_bottom_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorGetBufferFunction *gi.Function
var printCompositorGetBufferFunction_Once sync.Once

func printCompositorGetBufferFunction_Set() error {
	var err error
	printCompositorGetBufferFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetBufferFunction, err = printCompositorStruct.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_print_compositor_get_buffer.
func (recv *PrintCompositor) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetBufferFunction_Set()
	if err == nil {
		ret = printCompositorGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var printCompositorGetFooterFontNameFunction *gi.Function
var printCompositorGetFooterFontNameFunction_Once sync.Once

func printCompositorGetFooterFontNameFunction_Set() error {
	var err error
	printCompositorGetFooterFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetFooterFontNameFunction, err = printCompositorStruct.InvokerNew("get_footer_font_name")
	})
	return err
}

// GetFooterFontName is a representation of the C type gtk_source_print_compositor_get_footer_font_name.
func (recv *PrintCompositor) GetFooterFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetFooterFontNameFunction_Set()
	if err == nil {
		ret = printCompositorGetFooterFontNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var printCompositorGetHeaderFontNameFunction *gi.Function
var printCompositorGetHeaderFontNameFunction_Once sync.Once

func printCompositorGetHeaderFontNameFunction_Set() error {
	var err error
	printCompositorGetHeaderFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetHeaderFontNameFunction, err = printCompositorStruct.InvokerNew("get_header_font_name")
	})
	return err
}

// GetHeaderFontName is a representation of the C type gtk_source_print_compositor_get_header_font_name.
func (recv *PrintCompositor) GetHeaderFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetHeaderFontNameFunction_Set()
	if err == nil {
		ret = printCompositorGetHeaderFontNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var printCompositorGetHighlightSyntaxFunction *gi.Function
var printCompositorGetHighlightSyntaxFunction_Once sync.Once

func printCompositorGetHighlightSyntaxFunction_Set() error {
	var err error
	printCompositorGetHighlightSyntaxFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetHighlightSyntaxFunction, err = printCompositorStruct.InvokerNew("get_highlight_syntax")
	})
	return err
}

// GetHighlightSyntax is a representation of the C type gtk_source_print_compositor_get_highlight_syntax.
func (recv *PrintCompositor) GetHighlightSyntax() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetHighlightSyntaxFunction_Set()
	if err == nil {
		ret = printCompositorGetHighlightSyntaxFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_get_left_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorGetLineNumbersFontNameFunction *gi.Function
var printCompositorGetLineNumbersFontNameFunction_Once sync.Once

func printCompositorGetLineNumbersFontNameFunction_Set() error {
	var err error
	printCompositorGetLineNumbersFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetLineNumbersFontNameFunction, err = printCompositorStruct.InvokerNew("get_line_numbers_font_name")
	})
	return err
}

// GetLineNumbersFontName is a representation of the C type gtk_source_print_compositor_get_line_numbers_font_name.
func (recv *PrintCompositor) GetLineNumbersFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetLineNumbersFontNameFunction_Set()
	if err == nil {
		ret = printCompositorGetLineNumbersFontNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var printCompositorGetNPagesFunction *gi.Function
var printCompositorGetNPagesFunction_Once sync.Once

func printCompositorGetNPagesFunction_Set() error {
	var err error
	printCompositorGetNPagesFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetNPagesFunction, err = printCompositorStruct.InvokerNew("get_n_pages")
	})
	return err
}

// GetNPages is a representation of the C type gtk_source_print_compositor_get_n_pages.
func (recv *PrintCompositor) GetNPages() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetNPagesFunction_Set()
	if err == nil {
		ret = printCompositorGetNPagesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var printCompositorGetPaginationProgressFunction *gi.Function
var printCompositorGetPaginationProgressFunction_Once sync.Once

func printCompositorGetPaginationProgressFunction_Set() error {
	var err error
	printCompositorGetPaginationProgressFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetPaginationProgressFunction, err = printCompositorStruct.InvokerNew("get_pagination_progress")
	})
	return err
}

// GetPaginationProgress is a representation of the C type gtk_source_print_compositor_get_pagination_progress.
func (recv *PrintCompositor) GetPaginationProgress() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetPaginationProgressFunction_Set()
	if err == nil {
		ret = printCompositorGetPaginationProgressFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var printCompositorGetPrintFooterFunction *gi.Function
var printCompositorGetPrintFooterFunction_Once sync.Once

func printCompositorGetPrintFooterFunction_Set() error {
	var err error
	printCompositorGetPrintFooterFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintFooterFunction, err = printCompositorStruct.InvokerNew("get_print_footer")
	})
	return err
}

// GetPrintFooter is a representation of the C type gtk_source_print_compositor_get_print_footer.
func (recv *PrintCompositor) GetPrintFooter() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetPrintFooterFunction_Set()
	if err == nil {
		ret = printCompositorGetPrintFooterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var printCompositorGetPrintHeaderFunction *gi.Function
var printCompositorGetPrintHeaderFunction_Once sync.Once

func printCompositorGetPrintHeaderFunction_Set() error {
	var err error
	printCompositorGetPrintHeaderFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintHeaderFunction, err = printCompositorStruct.InvokerNew("get_print_header")
	})
	return err
}

// GetPrintHeader is a representation of the C type gtk_source_print_compositor_get_print_header.
func (recv *PrintCompositor) GetPrintHeader() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetPrintHeaderFunction_Set()
	if err == nil {
		ret = printCompositorGetPrintHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var printCompositorGetPrintLineNumbersFunction *gi.Function
var printCompositorGetPrintLineNumbersFunction_Once sync.Once

func printCompositorGetPrintLineNumbersFunction_Set() error {
	var err error
	printCompositorGetPrintLineNumbersFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintLineNumbersFunction, err = printCompositorStruct.InvokerNew("get_print_line_numbers")
	})
	return err
}

// GetPrintLineNumbers is a representation of the C type gtk_source_print_compositor_get_print_line_numbers.
func (recv *PrintCompositor) GetPrintLineNumbers() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetPrintLineNumbersFunction_Set()
	if err == nil {
		ret = printCompositorGetPrintLineNumbersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_get_right_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorGetTabWidthFunction *gi.Function
var printCompositorGetTabWidthFunction_Once sync.Once

func printCompositorGetTabWidthFunction_Set() error {
	var err error
	printCompositorGetTabWidthFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorGetTabWidthFunction, err = printCompositorStruct.InvokerNew("get_tab_width")
	})
	return err
}

// GetTabWidth is a representation of the C type gtk_source_print_compositor_get_tab_width.
func (recv *PrintCompositor) GetTabWidth() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := printCompositorGetTabWidthFunction_Set()
	if err == nil {
		ret = printCompositorGetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_get_top_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

// UNSUPPORTED : C value 'gtk_source_print_compositor_get_wrap_mode' : return type 'Gtk.WrapMode' not supported

// UNSUPPORTED : C value 'gtk_source_print_compositor_paginate' : parameter 'context' of type 'Gtk.PrintContext' not supported

var printCompositorSetBodyFontNameFunction *gi.Function
var printCompositorSetBodyFontNameFunction_Once sync.Once

func printCompositorSetBodyFontNameFunction_Set() error {
	var err error
	printCompositorSetBodyFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetBodyFontNameFunction, err = printCompositorStruct.InvokerNew("set_body_font_name")
	})
	return err
}

// SetBodyFontName is a representation of the C type gtk_source_print_compositor_set_body_font_name.
func (recv *PrintCompositor) SetBodyFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(fontName)

	err := printCompositorSetBodyFontNameFunction_Set()
	if err == nil {
		printCompositorSetBodyFontNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_bottom_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorSetFooterFontNameFunction *gi.Function
var printCompositorSetFooterFontNameFunction_Once sync.Once

func printCompositorSetFooterFontNameFunction_Set() error {
	var err error
	printCompositorSetFooterFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetFooterFontNameFunction, err = printCompositorStruct.InvokerNew("set_footer_font_name")
	})
	return err
}

// SetFooterFontName is a representation of the C type gtk_source_print_compositor_set_footer_font_name.
func (recv *PrintCompositor) SetFooterFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(fontName)

	err := printCompositorSetFooterFontNameFunction_Set()
	if err == nil {
		printCompositorSetFooterFontNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetFooterFormatFunction *gi.Function
var printCompositorSetFooterFormatFunction_Once sync.Once

func printCompositorSetFooterFormatFunction_Set() error {
	var err error
	printCompositorSetFooterFormatFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetFooterFormatFunction, err = printCompositorStruct.InvokerNew("set_footer_format")
	})
	return err
}

// SetFooterFormat is a representation of the C type gtk_source_print_compositor_set_footer_format.
func (recv *PrintCompositor) SetFooterFormat(separator bool, left string, center string, right string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(separator)
	inArgs[2].SetString(left)
	inArgs[3].SetString(center)
	inArgs[4].SetString(right)

	err := printCompositorSetFooterFormatFunction_Set()
	if err == nil {
		printCompositorSetFooterFormatFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetHeaderFontNameFunction *gi.Function
var printCompositorSetHeaderFontNameFunction_Once sync.Once

func printCompositorSetHeaderFontNameFunction_Set() error {
	var err error
	printCompositorSetHeaderFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetHeaderFontNameFunction, err = printCompositorStruct.InvokerNew("set_header_font_name")
	})
	return err
}

// SetHeaderFontName is a representation of the C type gtk_source_print_compositor_set_header_font_name.
func (recv *PrintCompositor) SetHeaderFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(fontName)

	err := printCompositorSetHeaderFontNameFunction_Set()
	if err == nil {
		printCompositorSetHeaderFontNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetHeaderFormatFunction *gi.Function
var printCompositorSetHeaderFormatFunction_Once sync.Once

func printCompositorSetHeaderFormatFunction_Set() error {
	var err error
	printCompositorSetHeaderFormatFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetHeaderFormatFunction, err = printCompositorStruct.InvokerNew("set_header_format")
	})
	return err
}

// SetHeaderFormat is a representation of the C type gtk_source_print_compositor_set_header_format.
func (recv *PrintCompositor) SetHeaderFormat(separator bool, left string, center string, right string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(separator)
	inArgs[2].SetString(left)
	inArgs[3].SetString(center)
	inArgs[4].SetString(right)

	err := printCompositorSetHeaderFormatFunction_Set()
	if err == nil {
		printCompositorSetHeaderFormatFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetHighlightSyntaxFunction *gi.Function
var printCompositorSetHighlightSyntaxFunction_Once sync.Once

func printCompositorSetHighlightSyntaxFunction_Set() error {
	var err error
	printCompositorSetHighlightSyntaxFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetHighlightSyntaxFunction, err = printCompositorStruct.InvokerNew("set_highlight_syntax")
	})
	return err
}

// SetHighlightSyntax is a representation of the C type gtk_source_print_compositor_set_highlight_syntax.
func (recv *PrintCompositor) SetHighlightSyntax(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(highlight)

	err := printCompositorSetHighlightSyntaxFunction_Set()
	if err == nil {
		printCompositorSetHighlightSyntaxFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_left_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorSetLineNumbersFontNameFunction *gi.Function
var printCompositorSetLineNumbersFontNameFunction_Once sync.Once

func printCompositorSetLineNumbersFontNameFunction_Set() error {
	var err error
	printCompositorSetLineNumbersFontNameFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetLineNumbersFontNameFunction, err = printCompositorStruct.InvokerNew("set_line_numbers_font_name")
	})
	return err
}

// SetLineNumbersFontName is a representation of the C type gtk_source_print_compositor_set_line_numbers_font_name.
func (recv *PrintCompositor) SetLineNumbersFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(fontName)

	err := printCompositorSetLineNumbersFontNameFunction_Set()
	if err == nil {
		printCompositorSetLineNumbersFontNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetPrintFooterFunction *gi.Function
var printCompositorSetPrintFooterFunction_Once sync.Once

func printCompositorSetPrintFooterFunction_Set() error {
	var err error
	printCompositorSetPrintFooterFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintFooterFunction, err = printCompositorStruct.InvokerNew("set_print_footer")
	})
	return err
}

// SetPrintFooter is a representation of the C type gtk_source_print_compositor_set_print_footer.
func (recv *PrintCompositor) SetPrintFooter(print bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(print)

	err := printCompositorSetPrintFooterFunction_Set()
	if err == nil {
		printCompositorSetPrintFooterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetPrintHeaderFunction *gi.Function
var printCompositorSetPrintHeaderFunction_Once sync.Once

func printCompositorSetPrintHeaderFunction_Set() error {
	var err error
	printCompositorSetPrintHeaderFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintHeaderFunction, err = printCompositorStruct.InvokerNew("set_print_header")
	})
	return err
}

// SetPrintHeader is a representation of the C type gtk_source_print_compositor_set_print_header.
func (recv *PrintCompositor) SetPrintHeader(print bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(print)

	err := printCompositorSetPrintHeaderFunction_Set()
	if err == nil {
		printCompositorSetPrintHeaderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorSetPrintLineNumbersFunction *gi.Function
var printCompositorSetPrintLineNumbersFunction_Once sync.Once

func printCompositorSetPrintLineNumbersFunction_Set() error {
	var err error
	printCompositorSetPrintLineNumbersFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintLineNumbersFunction, err = printCompositorStruct.InvokerNew("set_print_line_numbers")
	})
	return err
}

// SetPrintLineNumbers is a representation of the C type gtk_source_print_compositor_set_print_line_numbers.
func (recv *PrintCompositor) SetPrintLineNumbers(interval uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(interval)

	err := printCompositorSetPrintLineNumbersFunction_Set()
	if err == nil {
		printCompositorSetPrintLineNumbersFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_right_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

var printCompositorSetTabWidthFunction *gi.Function
var printCompositorSetTabWidthFunction_Once sync.Once

func printCompositorSetTabWidthFunction_Set() error {
	var err error
	printCompositorSetTabWidthFunction_Once.Do(func() {
		err = printCompositorStruct_Set()
		if err != nil {
			return
		}
		printCompositorSetTabWidthFunction, err = printCompositorStruct.InvokerNew("set_tab_width")
	})
	return err
}

// SetTabWidth is a representation of the C type gtk_source_print_compositor_set_tab_width.
func (recv *PrintCompositor) SetTabWidth(width uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(width)

	err := printCompositorSetTabWidthFunction_Set()
	if err == nil {
		printCompositorSetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_top_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_wrap_mode' : parameter 'wrap_mode' of type 'Gtk.WrapMode' not supported

var regionStruct *gi.Struct
var regionStruct_Once sync.Once

func regionStruct_Set() error {
	var err error
	regionStruct_Once.Do(func() {
		regionStruct, err = gi.StructNew("GtkSource", "Region")
	})
	return err
}

type Region struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'gtk_source_region_new' : parameter 'buffer' of type 'Gtk.TextBuffer' not supported

var regionAddRegionFunction *gi.Function
var regionAddRegionFunction_Once sync.Once

func regionAddRegionFunction_Set() error {
	var err error
	regionAddRegionFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionAddRegionFunction, err = regionStruct.InvokerNew("add_region")
	})
	return err
}

// AddRegion is a representation of the C type gtk_source_region_add_region.
func (recv *Region) AddRegion(regionToAdd *Region) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(regionToAdd.Native)

	err := regionAddRegionFunction_Set()
	if err == nil {
		regionAddRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_region_add_subregion' : parameter '_start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_region_get_bounds' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_region_get_buffer' : return type 'Gtk.TextBuffer' not supported

var regionGetStartRegionIterFunction *gi.Function
var regionGetStartRegionIterFunction_Once sync.Once

func regionGetStartRegionIterFunction_Set() error {
	var err error
	regionGetStartRegionIterFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionGetStartRegionIterFunction, err = regionStruct.InvokerNew("get_start_region_iter")
	})
	return err
}

// GetStartRegionIter is a representation of the C type gtk_source_region_get_start_region_iter.
func (recv *Region) GetStartRegionIter() *RegionIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [1]gi.Argument

	err := regionGetStartRegionIterFunction_Set()
	if err == nil {
		regionGetStartRegionIterFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &RegionIter{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var regionIntersectRegionFunction *gi.Function
var regionIntersectRegionFunction_Once sync.Once

func regionIntersectRegionFunction_Set() error {
	var err error
	regionIntersectRegionFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionIntersectRegionFunction, err = regionStruct.InvokerNew("intersect_region")
	})
	return err
}

// IntersectRegion is a representation of the C type gtk_source_region_intersect_region.
func (recv *Region) IntersectRegion(region2 *Region) *Region {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(region2.Native)

	var ret gi.Argument

	err := regionIntersectRegionFunction_Set()
	if err == nil {
		ret = regionIntersectRegionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Region{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_region_intersect_subregion' : parameter '_start' of type 'Gtk.TextIter' not supported

var regionIsEmptyFunction *gi.Function
var regionIsEmptyFunction_Once sync.Once

func regionIsEmptyFunction_Set() error {
	var err error
	regionIsEmptyFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionIsEmptyFunction, err = regionStruct.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type gtk_source_region_is_empty.
func (recv *Region) IsEmpty() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := regionIsEmptyFunction_Set()
	if err == nil {
		ret = regionIsEmptyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var regionSubtractRegionFunction *gi.Function
var regionSubtractRegionFunction_Once sync.Once

func regionSubtractRegionFunction_Set() error {
	var err error
	regionSubtractRegionFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionSubtractRegionFunction, err = regionStruct.InvokerNew("subtract_region")
	})
	return err
}

// SubtractRegion is a representation of the C type gtk_source_region_subtract_region.
func (recv *Region) SubtractRegion(regionToSubtract *Region) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(regionToSubtract.Native)

	err := regionSubtractRegionFunction_Set()
	if err == nil {
		regionSubtractRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_region_subtract_subregion' : parameter '_start' of type 'Gtk.TextIter' not supported

var regionToStringFunction *gi.Function
var regionToStringFunction_Once sync.Once

func regionToStringFunction_Set() error {
	var err error
	regionToStringFunction_Once.Do(func() {
		err = regionStruct_Set()
		if err != nil {
			return
		}
		regionToStringFunction, err = regionStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_source_region_to_string.
func (recv *Region) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := regionToStringFunction_Set()
	if err == nil {
		ret = regionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var searchContextStruct *gi.Struct
var searchContextStruct_Once sync.Once

func searchContextStruct_Set() error {
	var err error
	searchContextStruct_Once.Do(func() {
		searchContextStruct, err = gi.StructNew("GtkSource", "SearchContext")
	})
	return err
}

type SearchContext struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *SearchContext) FieldPriv() *SearchContextPrivate {
	argValue := gi.FieldGet(searchContextStruct, recv.Native, "priv")
	value := &SearchContextPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SearchContext) SetFieldPriv(value *SearchContextPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(searchContextStruct, recv.Native, "priv", argValue)
}

var searchContextNewFunction *gi.Function
var searchContextNewFunction_Once sync.Once

func searchContextNewFunction_Set() error {
	var err error
	searchContextNewFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextNewFunction, err = searchContextStruct.InvokerNew("new")
	})
	return err
}

// SearchContextNew is a representation of the C type gtk_source_search_context_new.
func SearchContextNew(buffer *Buffer, settings *SearchSettings) *SearchContext {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native)
	inArgs[1].SetPointer(settings.Native)

	var ret gi.Argument

	err := searchContextNewFunction_Set()
	if err == nil {
		ret = searchContextNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SearchContext{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_search_context_backward' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward2' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward_async' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward_finish2' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward2' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward_async' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward_finish2' : parameter 'result' of type 'Gio.AsyncResult' not supported

var searchContextGetBufferFunction *gi.Function
var searchContextGetBufferFunction_Once sync.Once

func searchContextGetBufferFunction_Set() error {
	var err error
	searchContextGetBufferFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextGetBufferFunction, err = searchContextStruct.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_search_context_get_buffer.
func (recv *SearchContext) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchContextGetBufferFunction_Set()
	if err == nil {
		ret = searchContextGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Buffer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var searchContextGetHighlightFunction *gi.Function
var searchContextGetHighlightFunction_Once sync.Once

func searchContextGetHighlightFunction_Set() error {
	var err error
	searchContextGetHighlightFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextGetHighlightFunction, err = searchContextStruct.InvokerNew("get_highlight")
	})
	return err
}

// GetHighlight is a representation of the C type gtk_source_search_context_get_highlight.
func (recv *SearchContext) GetHighlight() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchContextGetHighlightFunction_Set()
	if err == nil {
		ret = searchContextGetHighlightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchContextGetMatchStyleFunction *gi.Function
var searchContextGetMatchStyleFunction_Once sync.Once

func searchContextGetMatchStyleFunction_Set() error {
	var err error
	searchContextGetMatchStyleFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextGetMatchStyleFunction, err = searchContextStruct.InvokerNew("get_match_style")
	})
	return err
}

// GetMatchStyle is a representation of the C type gtk_source_search_context_get_match_style.
func (recv *SearchContext) GetMatchStyle() *Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchContextGetMatchStyleFunction_Set()
	if err == nil {
		ret = searchContextGetMatchStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Style{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_search_context_get_occurrence_position' : parameter 'match_start' of type 'Gtk.TextIter' not supported

var searchContextGetOccurrencesCountFunction *gi.Function
var searchContextGetOccurrencesCountFunction_Once sync.Once

func searchContextGetOccurrencesCountFunction_Set() error {
	var err error
	searchContextGetOccurrencesCountFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextGetOccurrencesCountFunction, err = searchContextStruct.InvokerNew("get_occurrences_count")
	})
	return err
}

// GetOccurrencesCount is a representation of the C type gtk_source_search_context_get_occurrences_count.
func (recv *SearchContext) GetOccurrencesCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchContextGetOccurrencesCountFunction_Set()
	if err == nil {
		ret = searchContextGetOccurrencesCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_search_context_get_regex_error' : return type 'GLib.Error' not supported

var searchContextGetSettingsFunction *gi.Function
var searchContextGetSettingsFunction_Once sync.Once

func searchContextGetSettingsFunction_Set() error {
	var err error
	searchContextGetSettingsFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextGetSettingsFunction, err = searchContextStruct.InvokerNew("get_settings")
	})
	return err
}

// GetSettings is a representation of the C type gtk_source_search_context_get_settings.
func (recv *SearchContext) GetSettings() *SearchSettings {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchContextGetSettingsFunction_Set()
	if err == nil {
		ret = searchContextGetSettingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SearchSettings{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_search_context_replace' : parameter 'match_start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_replace2' : parameter 'match_start' of type 'Gtk.TextIter' not supported

var searchContextReplaceAllFunction *gi.Function
var searchContextReplaceAllFunction_Once sync.Once

func searchContextReplaceAllFunction_Set() error {
	var err error
	searchContextReplaceAllFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextReplaceAllFunction, err = searchContextStruct.InvokerNew("replace_all")
	})
	return err
}

// ReplaceAll is a representation of the C type gtk_source_search_context_replace_all.
func (recv *SearchContext) ReplaceAll(replace string, replaceLength int32) uint32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(replace)
	inArgs[2].SetInt32(replaceLength)

	var ret gi.Argument

	err := searchContextReplaceAllFunction_Set()
	if err == nil {
		ret = searchContextReplaceAllFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var searchContextSetHighlightFunction *gi.Function
var searchContextSetHighlightFunction_Once sync.Once

func searchContextSetHighlightFunction_Set() error {
	var err error
	searchContextSetHighlightFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextSetHighlightFunction, err = searchContextStruct.InvokerNew("set_highlight")
	})
	return err
}

// SetHighlight is a representation of the C type gtk_source_search_context_set_highlight.
func (recv *SearchContext) SetHighlight(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(highlight)

	err := searchContextSetHighlightFunction_Set()
	if err == nil {
		searchContextSetHighlightFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchContextSetMatchStyleFunction *gi.Function
var searchContextSetMatchStyleFunction_Once sync.Once

func searchContextSetMatchStyleFunction_Set() error {
	var err error
	searchContextSetMatchStyleFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextSetMatchStyleFunction, err = searchContextStruct.InvokerNew("set_match_style")
	})
	return err
}

// SetMatchStyle is a representation of the C type gtk_source_search_context_set_match_style.
func (recv *SearchContext) SetMatchStyle(matchStyle *Style) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(matchStyle.Native)

	err := searchContextSetMatchStyleFunction_Set()
	if err == nil {
		searchContextSetMatchStyleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchContextSetSettingsFunction *gi.Function
var searchContextSetSettingsFunction_Once sync.Once

func searchContextSetSettingsFunction_Set() error {
	var err error
	searchContextSetSettingsFunction_Once.Do(func() {
		err = searchContextStruct_Set()
		if err != nil {
			return
		}
		searchContextSetSettingsFunction, err = searchContextStruct.InvokerNew("set_settings")
	})
	return err
}

// SetSettings is a representation of the C type gtk_source_search_context_set_settings.
func (recv *SearchContext) SetSettings(settings *SearchSettings) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(settings.Native)

	err := searchContextSetSettingsFunction_Set()
	if err == nil {
		searchContextSetSettingsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsStruct *gi.Struct
var searchSettingsStruct_Once sync.Once

func searchSettingsStruct_Set() error {
	var err error
	searchSettingsStruct_Once.Do(func() {
		searchSettingsStruct, err = gi.StructNew("GtkSource", "SearchSettings")
	})
	return err
}

type SearchSettings struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *SearchSettings) FieldPriv() *SearchSettingsPrivate {
	argValue := gi.FieldGet(searchSettingsStruct, recv.Native, "priv")
	value := &SearchSettingsPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SearchSettings) SetFieldPriv(value *SearchSettingsPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(searchSettingsStruct, recv.Native, "priv", argValue)
}

var searchSettingsNewFunction *gi.Function
var searchSettingsNewFunction_Once sync.Once

func searchSettingsNewFunction_Set() error {
	var err error
	searchSettingsNewFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsNewFunction, err = searchSettingsStruct.InvokerNew("new")
	})
	return err
}

// SearchSettingsNew is a representation of the C type gtk_source_search_settings_new.
func SearchSettingsNew() *SearchSettings {

	var ret gi.Argument

	err := searchSettingsNewFunction_Set()
	if err == nil {
		ret = searchSettingsNewFunction.Invoke(nil, nil)
	}

	retGo := &SearchSettings{}
	retGo.Native = ret.Pointer()

	return retGo
}

var searchSettingsGetAtWordBoundariesFunction *gi.Function
var searchSettingsGetAtWordBoundariesFunction_Once sync.Once

func searchSettingsGetAtWordBoundariesFunction_Set() error {
	var err error
	searchSettingsGetAtWordBoundariesFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsGetAtWordBoundariesFunction, err = searchSettingsStruct.InvokerNew("get_at_word_boundaries")
	})
	return err
}

// GetAtWordBoundaries is a representation of the C type gtk_source_search_settings_get_at_word_boundaries.
func (recv *SearchSettings) GetAtWordBoundaries() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchSettingsGetAtWordBoundariesFunction_Set()
	if err == nil {
		ret = searchSettingsGetAtWordBoundariesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchSettingsGetCaseSensitiveFunction *gi.Function
var searchSettingsGetCaseSensitiveFunction_Once sync.Once

func searchSettingsGetCaseSensitiveFunction_Set() error {
	var err error
	searchSettingsGetCaseSensitiveFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsGetCaseSensitiveFunction, err = searchSettingsStruct.InvokerNew("get_case_sensitive")
	})
	return err
}

// GetCaseSensitive is a representation of the C type gtk_source_search_settings_get_case_sensitive.
func (recv *SearchSettings) GetCaseSensitive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchSettingsGetCaseSensitiveFunction_Set()
	if err == nil {
		ret = searchSettingsGetCaseSensitiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchSettingsGetRegexEnabledFunction *gi.Function
var searchSettingsGetRegexEnabledFunction_Once sync.Once

func searchSettingsGetRegexEnabledFunction_Set() error {
	var err error
	searchSettingsGetRegexEnabledFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsGetRegexEnabledFunction, err = searchSettingsStruct.InvokerNew("get_regex_enabled")
	})
	return err
}

// GetRegexEnabled is a representation of the C type gtk_source_search_settings_get_regex_enabled.
func (recv *SearchSettings) GetRegexEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchSettingsGetRegexEnabledFunction_Set()
	if err == nil {
		ret = searchSettingsGetRegexEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchSettingsGetSearchTextFunction *gi.Function
var searchSettingsGetSearchTextFunction_Once sync.Once

func searchSettingsGetSearchTextFunction_Set() error {
	var err error
	searchSettingsGetSearchTextFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsGetSearchTextFunction, err = searchSettingsStruct.InvokerNew("get_search_text")
	})
	return err
}

// GetSearchText is a representation of the C type gtk_source_search_settings_get_search_text.
func (recv *SearchSettings) GetSearchText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchSettingsGetSearchTextFunction_Set()
	if err == nil {
		ret = searchSettingsGetSearchTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var searchSettingsGetWrapAroundFunction *gi.Function
var searchSettingsGetWrapAroundFunction_Once sync.Once

func searchSettingsGetWrapAroundFunction_Set() error {
	var err error
	searchSettingsGetWrapAroundFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsGetWrapAroundFunction, err = searchSettingsStruct.InvokerNew("get_wrap_around")
	})
	return err
}

// GetWrapAround is a representation of the C type gtk_source_search_settings_get_wrap_around.
func (recv *SearchSettings) GetWrapAround() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := searchSettingsGetWrapAroundFunction_Set()
	if err == nil {
		ret = searchSettingsGetWrapAroundFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchSettingsSetAtWordBoundariesFunction *gi.Function
var searchSettingsSetAtWordBoundariesFunction_Once sync.Once

func searchSettingsSetAtWordBoundariesFunction_Set() error {
	var err error
	searchSettingsSetAtWordBoundariesFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsSetAtWordBoundariesFunction, err = searchSettingsStruct.InvokerNew("set_at_word_boundaries")
	})
	return err
}

// SetAtWordBoundaries is a representation of the C type gtk_source_search_settings_set_at_word_boundaries.
func (recv *SearchSettings) SetAtWordBoundaries(atWordBoundaries bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(atWordBoundaries)

	err := searchSettingsSetAtWordBoundariesFunction_Set()
	if err == nil {
		searchSettingsSetAtWordBoundariesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsSetCaseSensitiveFunction *gi.Function
var searchSettingsSetCaseSensitiveFunction_Once sync.Once

func searchSettingsSetCaseSensitiveFunction_Set() error {
	var err error
	searchSettingsSetCaseSensitiveFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsSetCaseSensitiveFunction, err = searchSettingsStruct.InvokerNew("set_case_sensitive")
	})
	return err
}

// SetCaseSensitive is a representation of the C type gtk_source_search_settings_set_case_sensitive.
func (recv *SearchSettings) SetCaseSensitive(caseSensitive bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(caseSensitive)

	err := searchSettingsSetCaseSensitiveFunction_Set()
	if err == nil {
		searchSettingsSetCaseSensitiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsSetRegexEnabledFunction *gi.Function
var searchSettingsSetRegexEnabledFunction_Once sync.Once

func searchSettingsSetRegexEnabledFunction_Set() error {
	var err error
	searchSettingsSetRegexEnabledFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsSetRegexEnabledFunction, err = searchSettingsStruct.InvokerNew("set_regex_enabled")
	})
	return err
}

// SetRegexEnabled is a representation of the C type gtk_source_search_settings_set_regex_enabled.
func (recv *SearchSettings) SetRegexEnabled(regexEnabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(regexEnabled)

	err := searchSettingsSetRegexEnabledFunction_Set()
	if err == nil {
		searchSettingsSetRegexEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsSetSearchTextFunction *gi.Function
var searchSettingsSetSearchTextFunction_Once sync.Once

func searchSettingsSetSearchTextFunction_Set() error {
	var err error
	searchSettingsSetSearchTextFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsSetSearchTextFunction, err = searchSettingsStruct.InvokerNew("set_search_text")
	})
	return err
}

// SetSearchText is a representation of the C type gtk_source_search_settings_set_search_text.
func (recv *SearchSettings) SetSearchText(searchText string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(searchText)

	err := searchSettingsSetSearchTextFunction_Set()
	if err == nil {
		searchSettingsSetSearchTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsSetWrapAroundFunction *gi.Function
var searchSettingsSetWrapAroundFunction_Once sync.Once

func searchSettingsSetWrapAroundFunction_Set() error {
	var err error
	searchSettingsSetWrapAroundFunction_Once.Do(func() {
		err = searchSettingsStruct_Set()
		if err != nil {
			return
		}
		searchSettingsSetWrapAroundFunction, err = searchSettingsStruct.InvokerNew("set_wrap_around")
	})
	return err
}

// SetWrapAround is a representation of the C type gtk_source_search_settings_set_wrap_around.
func (recv *SearchSettings) SetWrapAround(wrapAround bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(wrapAround)

	err := searchSettingsSetWrapAroundFunction_Set()
	if err == nil {
		searchSettingsSetWrapAroundFunction.Invoke(inArgs[:], nil)
	}

	return
}

var spaceDrawerStruct *gi.Struct
var spaceDrawerStruct_Once sync.Once

func spaceDrawerStruct_Set() error {
	var err error
	spaceDrawerStruct_Once.Do(func() {
		spaceDrawerStruct, err = gi.StructNew("GtkSource", "SpaceDrawer")
	})
	return err
}

type SpaceDrawer struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *SpaceDrawer) FieldPriv() *SpaceDrawerPrivate {
	argValue := gi.FieldGet(spaceDrawerStruct, recv.Native, "priv")
	value := &SpaceDrawerPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SpaceDrawer) SetFieldPriv(value *SpaceDrawerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(spaceDrawerStruct, recv.Native, "priv", argValue)
}

var spaceDrawerNewFunction *gi.Function
var spaceDrawerNewFunction_Once sync.Once

func spaceDrawerNewFunction_Set() error {
	var err error
	spaceDrawerNewFunction_Once.Do(func() {
		err = spaceDrawerStruct_Set()
		if err != nil {
			return
		}
		spaceDrawerNewFunction, err = spaceDrawerStruct.InvokerNew("new")
	})
	return err
}

// SpaceDrawerNew is a representation of the C type gtk_source_space_drawer_new.
func SpaceDrawerNew() *SpaceDrawer {

	var ret gi.Argument

	err := spaceDrawerNewFunction_Set()
	if err == nil {
		ret = spaceDrawerNewFunction.Invoke(nil, nil)
	}

	retGo := &SpaceDrawer{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_space_drawer_bind_matrix_setting' : parameter 'settings' of type 'Gio.Settings' not supported

var spaceDrawerGetEnableMatrixFunction *gi.Function
var spaceDrawerGetEnableMatrixFunction_Once sync.Once

func spaceDrawerGetEnableMatrixFunction_Set() error {
	var err error
	spaceDrawerGetEnableMatrixFunction_Once.Do(func() {
		err = spaceDrawerStruct_Set()
		if err != nil {
			return
		}
		spaceDrawerGetEnableMatrixFunction, err = spaceDrawerStruct.InvokerNew("get_enable_matrix")
	})
	return err
}

// GetEnableMatrix is a representation of the C type gtk_source_space_drawer_get_enable_matrix.
func (recv *SpaceDrawer) GetEnableMatrix() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := spaceDrawerGetEnableMatrixFunction_Set()
	if err == nil {
		ret = spaceDrawerGetEnableMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_space_drawer_get_matrix' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_source_space_drawer_get_types_for_locations' : parameter 'locations' of type 'SpaceLocationFlags' not supported

var spaceDrawerSetEnableMatrixFunction *gi.Function
var spaceDrawerSetEnableMatrixFunction_Once sync.Once

func spaceDrawerSetEnableMatrixFunction_Set() error {
	var err error
	spaceDrawerSetEnableMatrixFunction_Once.Do(func() {
		err = spaceDrawerStruct_Set()
		if err != nil {
			return
		}
		spaceDrawerSetEnableMatrixFunction, err = spaceDrawerStruct.InvokerNew("set_enable_matrix")
	})
	return err
}

// SetEnableMatrix is a representation of the C type gtk_source_space_drawer_set_enable_matrix.
func (recv *SpaceDrawer) SetEnableMatrix(enableMatrix bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(enableMatrix)

	err := spaceDrawerSetEnableMatrixFunction_Set()
	if err == nil {
		spaceDrawerSetEnableMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_space_drawer_set_matrix' : parameter 'matrix' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_source_space_drawer_set_types_for_locations' : parameter 'locations' of type 'SpaceLocationFlags' not supported

var styleStruct *gi.Struct
var styleStruct_Once sync.Once

func styleStruct_Set() error {
	var err error
	styleStruct_Once.Do(func() {
		styleStruct, err = gi.StructNew("GtkSource", "Style")
	})
	return err
}

type Style struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gtk_source_style_apply' : parameter 'tag' of type 'Gtk.TextTag' not supported

var styleCopyFunction *gi.Function
var styleCopyFunction_Once sync.Once

func styleCopyFunction_Set() error {
	var err error
	styleCopyFunction_Once.Do(func() {
		err = styleStruct_Set()
		if err != nil {
			return
		}
		styleCopyFunction, err = styleStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_source_style_copy.
func (recv *Style) Copy() *Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := styleCopyFunction_Set()
	if err == nil {
		ret = styleCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Style{}
	retGo.Native = ret.Pointer()

	return retGo
}

// StyleStruct creates an uninitialised Style.
func StyleStruct() *Style {
	err := styleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Style{}
	structGo.Native = styleStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeStyle)
	return structGo
}
func finalizeStyle(obj *Style) {
	styleStruct.Free(obj.Native)
}

var styleSchemeStruct *gi.Struct
var styleSchemeStruct_Once sync.Once

func styleSchemeStruct_Set() error {
	var err error
	styleSchemeStruct_Once.Do(func() {
		styleSchemeStruct, err = gi.StructNew("GtkSource", "StyleScheme")
	})
	return err
}

type StyleScheme struct {
	gobject.Object
}

// UNSUPPORTED : C value 'base' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'base' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *StyleScheme) FieldPriv() *StyleSchemePrivate {
	argValue := gi.FieldGet(styleSchemeStruct, recv.Native, "priv")
	value := &StyleSchemePrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *StyleScheme) SetFieldPriv(value *StyleSchemePrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(styleSchemeStruct, recv.Native, "priv", argValue)
}

var styleSchemeGetAuthorsFunction *gi.Function
var styleSchemeGetAuthorsFunction_Once sync.Once

func styleSchemeGetAuthorsFunction_Set() error {
	var err error
	styleSchemeGetAuthorsFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetAuthorsFunction, err = styleSchemeStruct.InvokerNew("get_authors")
	})
	return err
}

// GetAuthors is a representation of the C type gtk_source_style_scheme_get_authors.
func (recv *StyleScheme) GetAuthors() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := styleSchemeGetAuthorsFunction_Set()
	if err == nil {
		styleSchemeGetAuthorsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleSchemeGetDescriptionFunction *gi.Function
var styleSchemeGetDescriptionFunction_Once sync.Once

func styleSchemeGetDescriptionFunction_Set() error {
	var err error
	styleSchemeGetDescriptionFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetDescriptionFunction, err = styleSchemeStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type gtk_source_style_scheme_get_description.
func (recv *StyleScheme) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := styleSchemeGetDescriptionFunction_Set()
	if err == nil {
		ret = styleSchemeGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var styleSchemeGetFilenameFunction *gi.Function
var styleSchemeGetFilenameFunction_Once sync.Once

func styleSchemeGetFilenameFunction_Set() error {
	var err error
	styleSchemeGetFilenameFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetFilenameFunction, err = styleSchemeStruct.InvokerNew("get_filename")
	})
	return err
}

// GetFilename is a representation of the C type gtk_source_style_scheme_get_filename.
func (recv *StyleScheme) GetFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := styleSchemeGetFilenameFunction_Set()
	if err == nil {
		ret = styleSchemeGetFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var styleSchemeGetIdFunction *gi.Function
var styleSchemeGetIdFunction_Once sync.Once

func styleSchemeGetIdFunction_Set() error {
	var err error
	styleSchemeGetIdFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetIdFunction, err = styleSchemeStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type gtk_source_style_scheme_get_id.
func (recv *StyleScheme) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := styleSchemeGetIdFunction_Set()
	if err == nil {
		ret = styleSchemeGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var styleSchemeGetNameFunction *gi.Function
var styleSchemeGetNameFunction_Once sync.Once

func styleSchemeGetNameFunction_Set() error {
	var err error
	styleSchemeGetNameFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetNameFunction, err = styleSchemeStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_style_scheme_get_name.
func (recv *StyleScheme) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := styleSchemeGetNameFunction_Set()
	if err == nil {
		ret = styleSchemeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var styleSchemeGetStyleFunction *gi.Function
var styleSchemeGetStyleFunction_Once sync.Once

func styleSchemeGetStyleFunction_Set() error {
	var err error
	styleSchemeGetStyleFunction_Once.Do(func() {
		err = styleSchemeStruct_Set()
		if err != nil {
			return
		}
		styleSchemeGetStyleFunction, err = styleSchemeStruct.InvokerNew("get_style")
	})
	return err
}

// GetStyle is a representation of the C type gtk_source_style_scheme_get_style.
func (recv *StyleScheme) GetStyle(styleId string) *Style {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(styleId)

	var ret gi.Argument

	err := styleSchemeGetStyleFunction_Set()
	if err == nil {
		ret = styleSchemeGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Style{}
	retGo.Native = ret.Pointer()

	return retGo
}

// StyleSchemeStruct creates an uninitialised StyleScheme.
func StyleSchemeStruct() *StyleScheme {
	err := styleSchemeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleScheme{}
	structGo.Native = styleSchemeStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeStyleScheme)
	return structGo
}
func finalizeStyleScheme(obj *StyleScheme) {
	styleSchemeStruct.Free(obj.Native)
}

var styleSchemeChooserButtonStruct *gi.Struct
var styleSchemeChooserButtonStruct_Once sync.Once

func styleSchemeChooserButtonStruct_Set() error {
	var err error
	styleSchemeChooserButtonStruct_Once.Do(func() {
		styleSchemeChooserButtonStruct, err = gi.StructNew("GtkSource", "StyleSchemeChooserButton")
	})
	return err
}

type StyleSchemeChooserButton struct {
	gtk.Button
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.Button'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.Button'

var styleSchemeChooserButtonNewFunction *gi.Function
var styleSchemeChooserButtonNewFunction_Once sync.Once

func styleSchemeChooserButtonNewFunction_Set() error {
	var err error
	styleSchemeChooserButtonNewFunction_Once.Do(func() {
		err = styleSchemeChooserButtonStruct_Set()
		if err != nil {
			return
		}
		styleSchemeChooserButtonNewFunction, err = styleSchemeChooserButtonStruct.InvokerNew("new")
	})
	return err
}

// StyleSchemeChooserButtonNew is a representation of the C type gtk_source_style_scheme_chooser_button_new.
func StyleSchemeChooserButtonNew() *StyleSchemeChooserButton {

	var ret gi.Argument

	err := styleSchemeChooserButtonNewFunction_Set()
	if err == nil {
		ret = styleSchemeChooserButtonNewFunction.Invoke(nil, nil)
	}

	retGo := &StyleSchemeChooserButton{}
	retGo.Native = ret.Pointer()

	return retGo
}

var styleSchemeChooserWidgetStruct *gi.Struct
var styleSchemeChooserWidgetStruct_Once sync.Once

func styleSchemeChooserWidgetStruct_Set() error {
	var err error
	styleSchemeChooserWidgetStruct_Once.Do(func() {
		styleSchemeChooserWidgetStruct, err = gi.StructNew("GtkSource", "StyleSchemeChooserWidget")
	})
	return err
}

type StyleSchemeChooserWidget struct {
	gtk.Bin
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.Bin'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.Bin'

var styleSchemeChooserWidgetNewFunction *gi.Function
var styleSchemeChooserWidgetNewFunction_Once sync.Once

func styleSchemeChooserWidgetNewFunction_Set() error {
	var err error
	styleSchemeChooserWidgetNewFunction_Once.Do(func() {
		err = styleSchemeChooserWidgetStruct_Set()
		if err != nil {
			return
		}
		styleSchemeChooserWidgetNewFunction, err = styleSchemeChooserWidgetStruct.InvokerNew("new")
	})
	return err
}

// StyleSchemeChooserWidgetNew is a representation of the C type gtk_source_style_scheme_chooser_widget_new.
func StyleSchemeChooserWidgetNew() *StyleSchemeChooserWidget {

	var ret gi.Argument

	err := styleSchemeChooserWidgetNewFunction_Set()
	if err == nil {
		ret = styleSchemeChooserWidgetNewFunction.Invoke(nil, nil)
	}

	retGo := &StyleSchemeChooserWidget{}
	retGo.Native = ret.Pointer()

	return retGo
}

var styleSchemeManagerStruct *gi.Struct
var styleSchemeManagerStruct_Once sync.Once

func styleSchemeManagerStruct_Set() error {
	var err error
	styleSchemeManagerStruct_Once.Do(func() {
		styleSchemeManagerStruct, err = gi.StructNew("GtkSource", "StyleSchemeManager")
	})
	return err
}

type StyleSchemeManager struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldPriv returns the C field 'priv'.
func (recv *StyleSchemeManager) FieldPriv() *StyleSchemeManagerPrivate {
	argValue := gi.FieldGet(styleSchemeManagerStruct, recv.Native, "priv")
	value := &StyleSchemeManagerPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *StyleSchemeManager) SetFieldPriv(value *StyleSchemeManagerPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(styleSchemeManagerStruct, recv.Native, "priv", argValue)
}

var styleSchemeManagerNewFunction *gi.Function
var styleSchemeManagerNewFunction_Once sync.Once

func styleSchemeManagerNewFunction_Set() error {
	var err error
	styleSchemeManagerNewFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerNewFunction, err = styleSchemeManagerStruct.InvokerNew("new")
	})
	return err
}

// StyleSchemeManagerNew is a representation of the C type gtk_source_style_scheme_manager_new.
func StyleSchemeManagerNew() *StyleSchemeManager {

	var ret gi.Argument

	err := styleSchemeManagerNewFunction_Set()
	if err == nil {
		ret = styleSchemeManagerNewFunction.Invoke(nil, nil)
	}

	retGo := &StyleSchemeManager{}
	retGo.Native = ret.Pointer()

	return retGo
}

var styleSchemeManagerAppendSearchPathFunction *gi.Function
var styleSchemeManagerAppendSearchPathFunction_Once sync.Once

func styleSchemeManagerAppendSearchPathFunction_Set() error {
	var err error
	styleSchemeManagerAppendSearchPathFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerAppendSearchPathFunction, err = styleSchemeManagerStruct.InvokerNew("append_search_path")
	})
	return err
}

// AppendSearchPath is a representation of the C type gtk_source_style_scheme_manager_append_search_path.
func (recv *StyleSchemeManager) AppendSearchPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(path)

	err := styleSchemeManagerAppendSearchPathFunction_Set()
	if err == nil {
		styleSchemeManagerAppendSearchPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleSchemeManagerForceRescanFunction *gi.Function
var styleSchemeManagerForceRescanFunction_Once sync.Once

func styleSchemeManagerForceRescanFunction_Set() error {
	var err error
	styleSchemeManagerForceRescanFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerForceRescanFunction, err = styleSchemeManagerStruct.InvokerNew("force_rescan")
	})
	return err
}

// ForceRescan is a representation of the C type gtk_source_style_scheme_manager_force_rescan.
func (recv *StyleSchemeManager) ForceRescan() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := styleSchemeManagerForceRescanFunction_Set()
	if err == nil {
		styleSchemeManagerForceRescanFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleSchemeManagerGetSchemeFunction *gi.Function
var styleSchemeManagerGetSchemeFunction_Once sync.Once

func styleSchemeManagerGetSchemeFunction_Set() error {
	var err error
	styleSchemeManagerGetSchemeFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSchemeFunction, err = styleSchemeManagerStruct.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type gtk_source_style_scheme_manager_get_scheme.
func (recv *StyleSchemeManager) GetScheme(schemeId string) *StyleScheme {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(schemeId)

	var ret gi.Argument

	err := styleSchemeManagerGetSchemeFunction_Set()
	if err == nil {
		ret = styleSchemeManagerGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StyleScheme{}
	retGo.Native = ret.Pointer()

	return retGo
}

var styleSchemeManagerGetSchemeIdsFunction *gi.Function
var styleSchemeManagerGetSchemeIdsFunction_Once sync.Once

func styleSchemeManagerGetSchemeIdsFunction_Set() error {
	var err error
	styleSchemeManagerGetSchemeIdsFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSchemeIdsFunction, err = styleSchemeManagerStruct.InvokerNew("get_scheme_ids")
	})
	return err
}

// GetSchemeIds is a representation of the C type gtk_source_style_scheme_manager_get_scheme_ids.
func (recv *StyleSchemeManager) GetSchemeIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := styleSchemeManagerGetSchemeIdsFunction_Set()
	if err == nil {
		styleSchemeManagerGetSchemeIdsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleSchemeManagerGetSearchPathFunction *gi.Function
var styleSchemeManagerGetSearchPathFunction_Once sync.Once

func styleSchemeManagerGetSearchPathFunction_Set() error {
	var err error
	styleSchemeManagerGetSearchPathFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSearchPathFunction, err = styleSchemeManagerStruct.InvokerNew("get_search_path")
	})
	return err
}

// GetSearchPath is a representation of the C type gtk_source_style_scheme_manager_get_search_path.
func (recv *StyleSchemeManager) GetSearchPath() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := styleSchemeManagerGetSearchPathFunction_Set()
	if err == nil {
		styleSchemeManagerGetSearchPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleSchemeManagerPrependSearchPathFunction *gi.Function
var styleSchemeManagerPrependSearchPathFunction_Once sync.Once

func styleSchemeManagerPrependSearchPathFunction_Set() error {
	var err error
	styleSchemeManagerPrependSearchPathFunction_Once.Do(func() {
		err = styleSchemeManagerStruct_Set()
		if err != nil {
			return
		}
		styleSchemeManagerPrependSearchPathFunction, err = styleSchemeManagerStruct.InvokerNew("prepend_search_path")
	})
	return err
}

// PrependSearchPath is a representation of the C type gtk_source_style_scheme_manager_prepend_search_path.
func (recv *StyleSchemeManager) PrependSearchPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(path)

	err := styleSchemeManagerPrependSearchPathFunction_Set()
	if err == nil {
		styleSchemeManagerPrependSearchPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_style_scheme_manager_set_search_path' : parameter 'path' of type 'nil' not supported

var tagStruct *gi.Struct
var tagStruct_Once sync.Once

func tagStruct_Set() error {
	var err error
	tagStruct_Once.Do(func() {
		tagStruct, err = gi.StructNew("GtkSource", "Tag")
	})
	return err
}

type Tag struct {
	gtk.TextTag
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Gtk.TextTag'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Gtk.TextTag'

var tagNewFunction *gi.Function
var tagNewFunction_Once sync.Once

func tagNewFunction_Set() error {
	var err error
	tagNewFunction_Once.Do(func() {
		err = tagStruct_Set()
		if err != nil {
			return
		}
		tagNewFunction, err = tagStruct.InvokerNew("new")
	})
	return err
}

// TagNew is a representation of the C type gtk_source_tag_new.
func TagNew(name string) *Tag {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := tagNewFunction_Set()
	if err == nil {
		ret = tagNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Tag{}
	retGo.Native = ret.Pointer()

	return retGo
}

var viewStruct *gi.Struct
var viewStruct_Once sync.Once

func viewStruct_Set() error {
	var err error
	viewStruct_Once.Do(func() {
		viewStruct, err = gi.StructNew("GtkSource", "View")
	})
	return err
}

type View struct {
	gtk.TextView
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.TextView'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.TextView'

// FieldPriv returns the C field 'priv'.
func (recv *View) FieldPriv() *ViewPrivate {
	argValue := gi.FieldGet(viewStruct, recv.Native, "priv")
	value := &ViewPrivate{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *View) SetFieldPriv(value *ViewPrivate) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(viewStruct, recv.Native, "priv", argValue)
}

var viewNewFunction *gi.Function
var viewNewFunction_Once sync.Once

func viewNewFunction_Set() error {
	var err error
	viewNewFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewNewFunction, err = viewStruct.InvokerNew("new")
	})
	return err
}

// ViewNew is a representation of the C type gtk_source_view_new.
func ViewNew() *View {

	var ret gi.Argument

	err := viewNewFunction_Set()
	if err == nil {
		ret = viewNewFunction.Invoke(nil, nil)
	}

	retGo := &View{}
	retGo.Native = ret.Pointer()

	return retGo
}

var viewNewWithBufferFunction *gi.Function
var viewNewWithBufferFunction_Once sync.Once

func viewNewWithBufferFunction_Set() error {
	var err error
	viewNewWithBufferFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewNewWithBufferFunction, err = viewStruct.InvokerNew("new_with_buffer")
	})
	return err
}

// ViewNewWithBuffer is a representation of the C type gtk_source_view_new_with_buffer.
func ViewNewWithBuffer(buffer *Buffer) *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(buffer.Native)

	var ret gi.Argument

	err := viewNewWithBufferFunction_Set()
	if err == nil {
		ret = viewNewWithBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &View{}
	retGo.Native = ret.Pointer()

	return retGo
}

var viewGetAutoIndentFunction *gi.Function
var viewGetAutoIndentFunction_Once sync.Once

func viewGetAutoIndentFunction_Set() error {
	var err error
	viewGetAutoIndentFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetAutoIndentFunction, err = viewStruct.InvokerNew("get_auto_indent")
	})
	return err
}

// GetAutoIndent is a representation of the C type gtk_source_view_get_auto_indent.
func (recv *View) GetAutoIndent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetAutoIndentFunction_Set()
	if err == nil {
		ret = viewGetAutoIndentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetBackgroundPatternFunction *gi.Function
var viewGetBackgroundPatternFunction_Once sync.Once

func viewGetBackgroundPatternFunction_Set() error {
	var err error
	viewGetBackgroundPatternFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetBackgroundPatternFunction, err = viewStruct.InvokerNew("get_background_pattern")
	})
	return err
}

// GetBackgroundPattern is a representation of the C type gtk_source_view_get_background_pattern.
func (recv *View) GetBackgroundPattern() BackgroundPatternType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetBackgroundPatternFunction_Set()
	if err == nil {
		ret = viewGetBackgroundPatternFunction.Invoke(inArgs[:], nil)
	}

	retGo := BackgroundPatternType(ret.Int32())

	return retGo
}

var viewGetCompletionFunction *gi.Function
var viewGetCompletionFunction_Once sync.Once

func viewGetCompletionFunction_Set() error {
	var err error
	viewGetCompletionFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetCompletionFunction, err = viewStruct.InvokerNew("get_completion")
	})
	return err
}

// GetCompletion is a representation of the C type gtk_source_view_get_completion.
func (recv *View) GetCompletion() *Completion {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetCompletionFunction_Set()
	if err == nil {
		ret = viewGetCompletionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Completion{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_view_get_draw_spaces' : return type 'DrawSpacesFlags' not supported

// UNSUPPORTED : C value 'gtk_source_view_get_gutter' : parameter 'window_type' of type 'Gtk.TextWindowType' not supported

var viewGetHighlightCurrentLineFunction *gi.Function
var viewGetHighlightCurrentLineFunction_Once sync.Once

func viewGetHighlightCurrentLineFunction_Set() error {
	var err error
	viewGetHighlightCurrentLineFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetHighlightCurrentLineFunction, err = viewStruct.InvokerNew("get_highlight_current_line")
	})
	return err
}

// GetHighlightCurrentLine is a representation of the C type gtk_source_view_get_highlight_current_line.
func (recv *View) GetHighlightCurrentLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetHighlightCurrentLineFunction_Set()
	if err == nil {
		ret = viewGetHighlightCurrentLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetIndentOnTabFunction *gi.Function
var viewGetIndentOnTabFunction_Once sync.Once

func viewGetIndentOnTabFunction_Set() error {
	var err error
	viewGetIndentOnTabFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetIndentOnTabFunction, err = viewStruct.InvokerNew("get_indent_on_tab")
	})
	return err
}

// GetIndentOnTab is a representation of the C type gtk_source_view_get_indent_on_tab.
func (recv *View) GetIndentOnTab() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetIndentOnTabFunction_Set()
	if err == nil {
		ret = viewGetIndentOnTabFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetIndentWidthFunction *gi.Function
var viewGetIndentWidthFunction_Once sync.Once

func viewGetIndentWidthFunction_Set() error {
	var err error
	viewGetIndentWidthFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetIndentWidthFunction, err = viewStruct.InvokerNew("get_indent_width")
	})
	return err
}

// GetIndentWidth is a representation of the C type gtk_source_view_get_indent_width.
func (recv *View) GetIndentWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetIndentWidthFunction_Set()
	if err == nil {
		ret = viewGetIndentWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var viewGetInsertSpacesInsteadOfTabsFunction *gi.Function
var viewGetInsertSpacesInsteadOfTabsFunction_Once sync.Once

func viewGetInsertSpacesInsteadOfTabsFunction_Set() error {
	var err error
	viewGetInsertSpacesInsteadOfTabsFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetInsertSpacesInsteadOfTabsFunction, err = viewStruct.InvokerNew("get_insert_spaces_instead_of_tabs")
	})
	return err
}

// GetInsertSpacesInsteadOfTabs is a representation of the C type gtk_source_view_get_insert_spaces_instead_of_tabs.
func (recv *View) GetInsertSpacesInsteadOfTabs() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetInsertSpacesInsteadOfTabsFunction_Set()
	if err == nil {
		ret = viewGetInsertSpacesInsteadOfTabsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetMarkAttributesFunction *gi.Function
var viewGetMarkAttributesFunction_Once sync.Once

func viewGetMarkAttributesFunction_Set() error {
	var err error
	viewGetMarkAttributesFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetMarkAttributesFunction, err = viewStruct.InvokerNew("get_mark_attributes")
	})
	return err
}

// GetMarkAttributes is a representation of the C type gtk_source_view_get_mark_attributes.
func (recv *View) GetMarkAttributes(category string, priority int32) *MarkAttributes {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(category)
	inArgs[2].SetInt32(priority)

	var ret gi.Argument

	err := viewGetMarkAttributesFunction_Set()
	if err == nil {
		ret = viewGetMarkAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MarkAttributes{}
	retGo.Native = ret.Pointer()

	return retGo
}

var viewGetRightMarginPositionFunction *gi.Function
var viewGetRightMarginPositionFunction_Once sync.Once

func viewGetRightMarginPositionFunction_Set() error {
	var err error
	viewGetRightMarginPositionFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetRightMarginPositionFunction, err = viewStruct.InvokerNew("get_right_margin_position")
	})
	return err
}

// GetRightMarginPosition is a representation of the C type gtk_source_view_get_right_margin_position.
func (recv *View) GetRightMarginPosition() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetRightMarginPositionFunction_Set()
	if err == nil {
		ret = viewGetRightMarginPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var viewGetShowLineMarksFunction *gi.Function
var viewGetShowLineMarksFunction_Once sync.Once

func viewGetShowLineMarksFunction_Set() error {
	var err error
	viewGetShowLineMarksFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetShowLineMarksFunction, err = viewStruct.InvokerNew("get_show_line_marks")
	})
	return err
}

// GetShowLineMarks is a representation of the C type gtk_source_view_get_show_line_marks.
func (recv *View) GetShowLineMarks() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetShowLineMarksFunction_Set()
	if err == nil {
		ret = viewGetShowLineMarksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetShowLineNumbersFunction *gi.Function
var viewGetShowLineNumbersFunction_Once sync.Once

func viewGetShowLineNumbersFunction_Set() error {
	var err error
	viewGetShowLineNumbersFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetShowLineNumbersFunction, err = viewStruct.InvokerNew("get_show_line_numbers")
	})
	return err
}

// GetShowLineNumbers is a representation of the C type gtk_source_view_get_show_line_numbers.
func (recv *View) GetShowLineNumbers() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetShowLineNumbersFunction_Set()
	if err == nil {
		ret = viewGetShowLineNumbersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetShowRightMarginFunction *gi.Function
var viewGetShowRightMarginFunction_Once sync.Once

func viewGetShowRightMarginFunction_Set() error {
	var err error
	viewGetShowRightMarginFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetShowRightMarginFunction, err = viewStruct.InvokerNew("get_show_right_margin")
	})
	return err
}

// GetShowRightMargin is a representation of the C type gtk_source_view_get_show_right_margin.
func (recv *View) GetShowRightMargin() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetShowRightMarginFunction_Set()
	if err == nil {
		ret = viewGetShowRightMarginFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetSmartBackspaceFunction *gi.Function
var viewGetSmartBackspaceFunction_Once sync.Once

func viewGetSmartBackspaceFunction_Set() error {
	var err error
	viewGetSmartBackspaceFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetSmartBackspaceFunction, err = viewStruct.InvokerNew("get_smart_backspace")
	})
	return err
}

// GetSmartBackspace is a representation of the C type gtk_source_view_get_smart_backspace.
func (recv *View) GetSmartBackspace() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetSmartBackspaceFunction_Set()
	if err == nil {
		ret = viewGetSmartBackspaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var viewGetSmartHomeEndFunction *gi.Function
var viewGetSmartHomeEndFunction_Once sync.Once

func viewGetSmartHomeEndFunction_Set() error {
	var err error
	viewGetSmartHomeEndFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetSmartHomeEndFunction, err = viewStruct.InvokerNew("get_smart_home_end")
	})
	return err
}

// GetSmartHomeEnd is a representation of the C type gtk_source_view_get_smart_home_end.
func (recv *View) GetSmartHomeEnd() SmartHomeEndType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetSmartHomeEndFunction_Set()
	if err == nil {
		ret = viewGetSmartHomeEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := SmartHomeEndType(ret.Int32())

	return retGo
}

var viewGetSpaceDrawerFunction *gi.Function
var viewGetSpaceDrawerFunction_Once sync.Once

func viewGetSpaceDrawerFunction_Set() error {
	var err error
	viewGetSpaceDrawerFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetSpaceDrawerFunction, err = viewStruct.InvokerNew("get_space_drawer")
	})
	return err
}

// GetSpaceDrawer is a representation of the C type gtk_source_view_get_space_drawer.
func (recv *View) GetSpaceDrawer() *SpaceDrawer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetSpaceDrawerFunction_Set()
	if err == nil {
		ret = viewGetSpaceDrawerFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SpaceDrawer{}
	retGo.Native = ret.Pointer()

	return retGo
}

var viewGetTabWidthFunction *gi.Function
var viewGetTabWidthFunction_Once sync.Once

func viewGetTabWidthFunction_Set() error {
	var err error
	viewGetTabWidthFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewGetTabWidthFunction, err = viewStruct.InvokerNew("get_tab_width")
	})
	return err
}

// GetTabWidth is a representation of the C type gtk_source_view_get_tab_width.
func (recv *View) GetTabWidth() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := viewGetTabWidthFunction_Set()
	if err == nil {
		ret = viewGetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_view_get_visual_column' : parameter 'iter' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_view_indent_lines' : parameter 'start' of type 'Gtk.TextIter' not supported

var viewSetAutoIndentFunction *gi.Function
var viewSetAutoIndentFunction_Once sync.Once

func viewSetAutoIndentFunction_Set() error {
	var err error
	viewSetAutoIndentFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetAutoIndentFunction, err = viewStruct.InvokerNew("set_auto_indent")
	})
	return err
}

// SetAutoIndent is a representation of the C type gtk_source_view_set_auto_indent.
func (recv *View) SetAutoIndent(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(enable)

	err := viewSetAutoIndentFunction_Set()
	if err == nil {
		viewSetAutoIndentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetBackgroundPatternFunction *gi.Function
var viewSetBackgroundPatternFunction_Once sync.Once

func viewSetBackgroundPatternFunction_Set() error {
	var err error
	viewSetBackgroundPatternFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetBackgroundPatternFunction, err = viewStruct.InvokerNew("set_background_pattern")
	})
	return err
}

// SetBackgroundPattern is a representation of the C type gtk_source_view_set_background_pattern.
func (recv *View) SetBackgroundPattern(backgroundPattern BackgroundPatternType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(int32(backgroundPattern))

	err := viewSetBackgroundPatternFunction_Set()
	if err == nil {
		viewSetBackgroundPatternFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_view_set_draw_spaces' : parameter 'flags' of type 'DrawSpacesFlags' not supported

var viewSetHighlightCurrentLineFunction *gi.Function
var viewSetHighlightCurrentLineFunction_Once sync.Once

func viewSetHighlightCurrentLineFunction_Set() error {
	var err error
	viewSetHighlightCurrentLineFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetHighlightCurrentLineFunction, err = viewStruct.InvokerNew("set_highlight_current_line")
	})
	return err
}

// SetHighlightCurrentLine is a representation of the C type gtk_source_view_set_highlight_current_line.
func (recv *View) SetHighlightCurrentLine(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(highlight)

	err := viewSetHighlightCurrentLineFunction_Set()
	if err == nil {
		viewSetHighlightCurrentLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetIndentOnTabFunction *gi.Function
var viewSetIndentOnTabFunction_Once sync.Once

func viewSetIndentOnTabFunction_Set() error {
	var err error
	viewSetIndentOnTabFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetIndentOnTabFunction, err = viewStruct.InvokerNew("set_indent_on_tab")
	})
	return err
}

// SetIndentOnTab is a representation of the C type gtk_source_view_set_indent_on_tab.
func (recv *View) SetIndentOnTab(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(enable)

	err := viewSetIndentOnTabFunction_Set()
	if err == nil {
		viewSetIndentOnTabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetIndentWidthFunction *gi.Function
var viewSetIndentWidthFunction_Once sync.Once

func viewSetIndentWidthFunction_Set() error {
	var err error
	viewSetIndentWidthFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetIndentWidthFunction, err = viewStruct.InvokerNew("set_indent_width")
	})
	return err
}

// SetIndentWidth is a representation of the C type gtk_source_view_set_indent_width.
func (recv *View) SetIndentWidth(width int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(width)

	err := viewSetIndentWidthFunction_Set()
	if err == nil {
		viewSetIndentWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetInsertSpacesInsteadOfTabsFunction *gi.Function
var viewSetInsertSpacesInsteadOfTabsFunction_Once sync.Once

func viewSetInsertSpacesInsteadOfTabsFunction_Set() error {
	var err error
	viewSetInsertSpacesInsteadOfTabsFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetInsertSpacesInsteadOfTabsFunction, err = viewStruct.InvokerNew("set_insert_spaces_instead_of_tabs")
	})
	return err
}

// SetInsertSpacesInsteadOfTabs is a representation of the C type gtk_source_view_set_insert_spaces_instead_of_tabs.
func (recv *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(enable)

	err := viewSetInsertSpacesInsteadOfTabsFunction_Set()
	if err == nil {
		viewSetInsertSpacesInsteadOfTabsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetMarkAttributesFunction *gi.Function
var viewSetMarkAttributesFunction_Once sync.Once

func viewSetMarkAttributesFunction_Set() error {
	var err error
	viewSetMarkAttributesFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetMarkAttributesFunction, err = viewStruct.InvokerNew("set_mark_attributes")
	})
	return err
}

// SetMarkAttributes is a representation of the C type gtk_source_view_set_mark_attributes.
func (recv *View) SetMarkAttributes(category string, attributes *MarkAttributes, priority int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(category)
	inArgs[2].SetPointer(attributes.Native)
	inArgs[3].SetInt32(priority)

	err := viewSetMarkAttributesFunction_Set()
	if err == nil {
		viewSetMarkAttributesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetRightMarginPositionFunction *gi.Function
var viewSetRightMarginPositionFunction_Once sync.Once

func viewSetRightMarginPositionFunction_Set() error {
	var err error
	viewSetRightMarginPositionFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetRightMarginPositionFunction, err = viewStruct.InvokerNew("set_right_margin_position")
	})
	return err
}

// SetRightMarginPosition is a representation of the C type gtk_source_view_set_right_margin_position.
func (recv *View) SetRightMarginPosition(pos uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(pos)

	err := viewSetRightMarginPositionFunction_Set()
	if err == nil {
		viewSetRightMarginPositionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetShowLineMarksFunction *gi.Function
var viewSetShowLineMarksFunction_Once sync.Once

func viewSetShowLineMarksFunction_Set() error {
	var err error
	viewSetShowLineMarksFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetShowLineMarksFunction, err = viewStruct.InvokerNew("set_show_line_marks")
	})
	return err
}

// SetShowLineMarks is a representation of the C type gtk_source_view_set_show_line_marks.
func (recv *View) SetShowLineMarks(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(show)

	err := viewSetShowLineMarksFunction_Set()
	if err == nil {
		viewSetShowLineMarksFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetShowLineNumbersFunction *gi.Function
var viewSetShowLineNumbersFunction_Once sync.Once

func viewSetShowLineNumbersFunction_Set() error {
	var err error
	viewSetShowLineNumbersFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetShowLineNumbersFunction, err = viewStruct.InvokerNew("set_show_line_numbers")
	})
	return err
}

// SetShowLineNumbers is a representation of the C type gtk_source_view_set_show_line_numbers.
func (recv *View) SetShowLineNumbers(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(show)

	err := viewSetShowLineNumbersFunction_Set()
	if err == nil {
		viewSetShowLineNumbersFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetShowRightMarginFunction *gi.Function
var viewSetShowRightMarginFunction_Once sync.Once

func viewSetShowRightMarginFunction_Set() error {
	var err error
	viewSetShowRightMarginFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetShowRightMarginFunction, err = viewStruct.InvokerNew("set_show_right_margin")
	})
	return err
}

// SetShowRightMargin is a representation of the C type gtk_source_view_set_show_right_margin.
func (recv *View) SetShowRightMargin(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(show)

	err := viewSetShowRightMarginFunction_Set()
	if err == nil {
		viewSetShowRightMarginFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetSmartBackspaceFunction *gi.Function
var viewSetSmartBackspaceFunction_Once sync.Once

func viewSetSmartBackspaceFunction_Set() error {
	var err error
	viewSetSmartBackspaceFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetSmartBackspaceFunction, err = viewStruct.InvokerNew("set_smart_backspace")
	})
	return err
}

// SetSmartBackspace is a representation of the C type gtk_source_view_set_smart_backspace.
func (recv *View) SetSmartBackspace(smartBackspace bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(smartBackspace)

	err := viewSetSmartBackspaceFunction_Set()
	if err == nil {
		viewSetSmartBackspaceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetSmartHomeEndFunction *gi.Function
var viewSetSmartHomeEndFunction_Once sync.Once

func viewSetSmartHomeEndFunction_Set() error {
	var err error
	viewSetSmartHomeEndFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetSmartHomeEndFunction, err = viewStruct.InvokerNew("set_smart_home_end")
	})
	return err
}

// SetSmartHomeEnd is a representation of the C type gtk_source_view_set_smart_home_end.
func (recv *View) SetSmartHomeEnd(smartHomeEnd SmartHomeEndType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(int32(smartHomeEnd))

	err := viewSetSmartHomeEndFunction_Set()
	if err == nil {
		viewSetSmartHomeEndFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetTabWidthFunction *gi.Function
var viewSetTabWidthFunction_Once sync.Once

func viewSetTabWidthFunction_Set() error {
	var err error
	viewSetTabWidthFunction_Once.Do(func() {
		err = viewStruct_Set()
		if err != nil {
			return
		}
		viewSetTabWidthFunction, err = viewStruct.InvokerNew("set_tab_width")
	})
	return err
}

// SetTabWidth is a representation of the C type gtk_source_view_set_tab_width.
func (recv *View) SetTabWidth(width uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(width)

	err := viewSetTabWidthFunction_Set()
	if err == nil {
		viewSetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_view_unindent_lines' : parameter 'start' of type 'Gtk.TextIter' not supported
