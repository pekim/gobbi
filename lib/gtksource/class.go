// Code generated - DO NOT EDIT.

package gtksource

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/gtk"
	"runtime"
	"sync"
	"unsafe"
)

var bufferObject *gi.Object
var bufferObject_Once sync.Once

func bufferObject_Set() error {
	var err error
	bufferObject_Once.Do(func() {
		bufferObject, err = gi.ObjectNew("GtkSource", "Buffer")
	})
	return err
}

type Buffer struct {
	native unsafe.Pointer
}

func BufferNewFromNative(native unsafe.Pointer) *Buffer {
	instance := &Buffer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TextBuffer upcasts to *TextBuffer
func (recv *Buffer) TextBuffer() *gtk.TextBuffer {
	return gtk.TextBufferNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Buffer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToBuffer down casts any arbitrary Object to Buffer.
Exercise care, as this is a potentially dangerous function
if the Object is not a Buffer.
*/
func CastToBuffer(object *gobject.Object) *Buffer {
	return BufferNewFromNative(object.Native())
}

// Equals compares this Buffer with another Buffer, and returns true if they represent the same Object.
func (recv *Buffer) Equals(other *Buffer) bool {
	return other.Native() == recv.Native()
}

func (recv *Buffer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Buffer) FieldParentInstance() *gtk.TextBuffer {
	var nilValue *gtk.TextBuffer
	err := bufferObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(bufferObject, recv.Native(), "parent_instance")
	value := gtk.TextBufferNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Buffer) SetFieldParentInstance(value *gtk.TextBuffer) {
	err := bufferObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(bufferObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Buffer) FieldPriv() *BufferPrivate {
	var nilValue *BufferPrivate
	err := bufferObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(bufferObject, recv.Native(), "priv")
	value := BufferPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Buffer) SetFieldPriv(value *BufferPrivate) {
	err := bufferObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(bufferObject, recv.Native(), "priv", argValue)
}

var bufferNewFunction *gi.Function
var bufferNewFunction_Once sync.Once

func bufferNewFunction_Set() error {
	var err error
	bufferNewFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferNewFunction, err = bufferObject.InvokerNew("new")
	})
	return err
}

// BufferNew is a representation of the C type gtk_source_buffer_new.
func BufferNew(table *gtk.TextTagTable) *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(table.Native())

	var ret gi.Argument

	err := bufferNewFunction_Set()
	if err == nil {
		ret = bufferNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferNewWithLanguageFunction *gi.Function
var bufferNewWithLanguageFunction_Once sync.Once

func bufferNewWithLanguageFunction_Set() error {
	var err error
	bufferNewWithLanguageFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferNewWithLanguageFunction, err = bufferObject.InvokerNew("new_with_language")
	})
	return err
}

// BufferNewWithLanguage is a representation of the C type gtk_source_buffer_new_with_language.
func BufferNewWithLanguage(language *Language) *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(language.Native())

	var ret gi.Argument

	err := bufferNewWithLanguageFunction_Set()
	if err == nil {
		ret = bufferNewWithLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var bufferBackwardIterToSourceMarkFunction *gi.Function
var bufferBackwardIterToSourceMarkFunction_Once sync.Once

func bufferBackwardIterToSourceMarkFunction_Set() error {
	var err error
	bufferBackwardIterToSourceMarkFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferBackwardIterToSourceMarkFunction, err = bufferObject.InvokerNew("backward_iter_to_source_mark")
	})
	return err
}

// BackwardIterToSourceMark is a representation of the C type gtk_source_buffer_backward_iter_to_source_mark.
func (recv *Buffer) BackwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(category)

	var ret gi.Argument

	err := bufferBackwardIterToSourceMarkFunction_Set()
	if err == nil {
		ret = bufferBackwardIterToSourceMarkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferBeginNotUndoableActionFunction *gi.Function
var bufferBeginNotUndoableActionFunction_Once sync.Once

func bufferBeginNotUndoableActionFunction_Set() error {
	var err error
	bufferBeginNotUndoableActionFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferBeginNotUndoableActionFunction, err = bufferObject.InvokerNew("begin_not_undoable_action")
	})
	return err
}

// BeginNotUndoableAction is a representation of the C type gtk_source_buffer_begin_not_undoable_action.
func (recv *Buffer) BeginNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferCanRedoFunction, err = bufferObject.InvokerNew("can_redo")
	})
	return err
}

// CanRedo is a representation of the C type gtk_source_buffer_can_redo.
func (recv *Buffer) CanRedo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferCanUndoFunction, err = bufferObject.InvokerNew("can_undo")
	})
	return err
}

// CanUndo is a representation of the C type gtk_source_buffer_can_undo.
func (recv *Buffer) CanUndo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferCanUndoFunction_Set()
	if err == nil {
		ret = bufferCanUndoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferChangeCaseFunction *gi.Function
var bufferChangeCaseFunction_Once sync.Once

func bufferChangeCaseFunction_Set() error {
	var err error
	bufferChangeCaseFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferChangeCaseFunction, err = bufferObject.InvokerNew("change_case")
	})
	return err
}

// ChangeCase is a representation of the C type gtk_source_buffer_change_case.
func (recv *Buffer) ChangeCase(caseType ChangeCaseType, start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(caseType))
	inArgs[2].SetPointer(start.Native())
	inArgs[3].SetPointer(end.Native())

	err := bufferChangeCaseFunction_Set()
	if err == nil {
		bufferChangeCaseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferCreateSourceMarkFunction *gi.Function
var bufferCreateSourceMarkFunction_Once sync.Once

func bufferCreateSourceMarkFunction_Set() error {
	var err error
	bufferCreateSourceMarkFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferCreateSourceMarkFunction, err = bufferObject.InvokerNew("create_source_mark")
	})
	return err
}

// CreateSourceMark is a representation of the C type gtk_source_buffer_create_source_mark.
func (recv *Buffer) CreateSourceMark(name string, category string, where *gtk.TextIter) *Mark {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)
	inArgs[2].SetString(category)
	inArgs[3].SetPointer(where.Native())

	var ret gi.Argument

	err := bufferCreateSourceMarkFunction_Set()
	if err == nil {
		ret = bufferCreateSourceMarkFunction.Invoke(inArgs[:], nil)
	}

	retGo := MarkNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_create_source_tag' : parameter '...' of type 'nil' not supported

var bufferEndNotUndoableActionFunction *gi.Function
var bufferEndNotUndoableActionFunction_Once sync.Once

func bufferEndNotUndoableActionFunction_Set() error {
	var err error
	bufferEndNotUndoableActionFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferEndNotUndoableActionFunction, err = bufferObject.InvokerNew("end_not_undoable_action")
	})
	return err
}

// EndNotUndoableAction is a representation of the C type gtk_source_buffer_end_not_undoable_action.
func (recv *Buffer) EndNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := bufferEndNotUndoableActionFunction_Set()
	if err == nil {
		bufferEndNotUndoableActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferEnsureHighlightFunction *gi.Function
var bufferEnsureHighlightFunction_Once sync.Once

func bufferEnsureHighlightFunction_Set() error {
	var err error
	bufferEnsureHighlightFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferEnsureHighlightFunction, err = bufferObject.InvokerNew("ensure_highlight")
	})
	return err
}

// EnsureHighlight is a representation of the C type gtk_source_buffer_ensure_highlight.
func (recv *Buffer) EnsureHighlight(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := bufferEnsureHighlightFunction_Set()
	if err == nil {
		bufferEnsureHighlightFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferForwardIterToSourceMarkFunction *gi.Function
var bufferForwardIterToSourceMarkFunction_Once sync.Once

func bufferForwardIterToSourceMarkFunction_Set() error {
	var err error
	bufferForwardIterToSourceMarkFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferForwardIterToSourceMarkFunction, err = bufferObject.InvokerNew("forward_iter_to_source_mark")
	})
	return err
}

// ForwardIterToSourceMark is a representation of the C type gtk_source_buffer_forward_iter_to_source_mark.
func (recv *Buffer) ForwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(category)

	var ret gi.Argument

	err := bufferForwardIterToSourceMarkFunction_Set()
	if err == nil {
		ret = bufferForwardIterToSourceMarkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferGetContextClassesAtIterFunction *gi.Function
var bufferGetContextClassesAtIterFunction_Once sync.Once

func bufferGetContextClassesAtIterFunction_Set() error {
	var err error
	bufferGetContextClassesAtIterFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetContextClassesAtIterFunction, err = bufferObject.InvokerNew("get_context_classes_at_iter")
	})
	return err
}

// GetContextClassesAtIter is a representation of the C type gtk_source_buffer_get_context_classes_at_iter.
func (recv *Buffer) GetContextClassesAtIter(iter *gtk.TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	err := bufferGetContextClassesAtIterFunction_Set()
	if err == nil {
		bufferGetContextClassesAtIterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferGetHighlightMatchingBracketsFunction *gi.Function
var bufferGetHighlightMatchingBracketsFunction_Once sync.Once

func bufferGetHighlightMatchingBracketsFunction_Set() error {
	var err error
	bufferGetHighlightMatchingBracketsFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetHighlightMatchingBracketsFunction, err = bufferObject.InvokerNew("get_highlight_matching_brackets")
	})
	return err
}

// GetHighlightMatchingBrackets is a representation of the C type gtk_source_buffer_get_highlight_matching_brackets.
func (recv *Buffer) GetHighlightMatchingBrackets() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetHighlightSyntaxFunction, err = bufferObject.InvokerNew("get_highlight_syntax")
	})
	return err
}

// GetHighlightSyntax is a representation of the C type gtk_source_buffer_get_highlight_syntax.
func (recv *Buffer) GetHighlightSyntax() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetImplicitTrailingNewlineFunction, err = bufferObject.InvokerNew("get_implicit_trailing_newline")
	})
	return err
}

// GetImplicitTrailingNewline is a representation of the C type gtk_source_buffer_get_implicit_trailing_newline.
func (recv *Buffer) GetImplicitTrailingNewline() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetLanguageFunction, err = bufferObject.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type gtk_source_buffer_get_language.
func (recv *Buffer) GetLanguage() *Language {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferGetLanguageFunction_Set()
	if err == nil {
		ret = bufferGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := LanguageNewFromNative(ret.Pointer())

	return retGo
}

var bufferGetMaxUndoLevelsFunction *gi.Function
var bufferGetMaxUndoLevelsFunction_Once sync.Once

func bufferGetMaxUndoLevelsFunction_Set() error {
	var err error
	bufferGetMaxUndoLevelsFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetMaxUndoLevelsFunction, err = bufferObject.InvokerNew("get_max_undo_levels")
	})
	return err
}

// GetMaxUndoLevels is a representation of the C type gtk_source_buffer_get_max_undo_levels.
func (recv *Buffer) GetMaxUndoLevels() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferGetMaxUndoLevelsFunction_Set()
	if err == nil {
		ret = bufferGetMaxUndoLevelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var bufferGetSourceMarksAtIterFunction *gi.Function
var bufferGetSourceMarksAtIterFunction_Once sync.Once

func bufferGetSourceMarksAtIterFunction_Set() error {
	var err error
	bufferGetSourceMarksAtIterFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetSourceMarksAtIterFunction, err = bufferObject.InvokerNew("get_source_marks_at_iter")
	})
	return err
}

// GetSourceMarksAtIter is a representation of the C type gtk_source_buffer_get_source_marks_at_iter.
func (recv *Buffer) GetSourceMarksAtIter(iter *gtk.TextIter, category string) *glib.SList {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(category)

	var ret gi.Argument

	err := bufferGetSourceMarksAtIterFunction_Set()
	if err == nil {
		ret = bufferGetSourceMarksAtIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var bufferGetSourceMarksAtLineFunction *gi.Function
var bufferGetSourceMarksAtLineFunction_Once sync.Once

func bufferGetSourceMarksAtLineFunction_Set() error {
	var err error
	bufferGetSourceMarksAtLineFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetSourceMarksAtLineFunction, err = bufferObject.InvokerNew("get_source_marks_at_line")
	})
	return err
}

// GetSourceMarksAtLine is a representation of the C type gtk_source_buffer_get_source_marks_at_line.
func (recv *Buffer) GetSourceMarksAtLine(line int32, category string) *glib.SList {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(line)
	inArgs[2].SetString(category)

	var ret gi.Argument

	err := bufferGetSourceMarksAtLineFunction_Set()
	if err == nil {
		ret = bufferGetSourceMarksAtLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var bufferGetStyleSchemeFunction *gi.Function
var bufferGetStyleSchemeFunction_Once sync.Once

func bufferGetStyleSchemeFunction_Set() error {
	var err error
	bufferGetStyleSchemeFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferGetStyleSchemeFunction, err = bufferObject.InvokerNew("get_style_scheme")
	})
	return err
}

// GetStyleScheme is a representation of the C type gtk_source_buffer_get_style_scheme.
func (recv *Buffer) GetStyleScheme() *StyleScheme {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := bufferGetStyleSchemeFunction_Set()
	if err == nil {
		ret = bufferGetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleSchemeNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_buffer_get_undo_manager' : return type 'UndoManager' not supported

var bufferIterBackwardToContextClassToggleFunction *gi.Function
var bufferIterBackwardToContextClassToggleFunction_Once sync.Once

func bufferIterBackwardToContextClassToggleFunction_Set() error {
	var err error
	bufferIterBackwardToContextClassToggleFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferIterBackwardToContextClassToggleFunction, err = bufferObject.InvokerNew("iter_backward_to_context_class_toggle")
	})
	return err
}

// IterBackwardToContextClassToggle is a representation of the C type gtk_source_buffer_iter_backward_to_context_class_toggle.
func (recv *Buffer) IterBackwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(contextClass)

	var ret gi.Argument

	err := bufferIterBackwardToContextClassToggleFunction_Set()
	if err == nil {
		ret = bufferIterBackwardToContextClassToggleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferIterForwardToContextClassToggleFunction *gi.Function
var bufferIterForwardToContextClassToggleFunction_Once sync.Once

func bufferIterForwardToContextClassToggleFunction_Set() error {
	var err error
	bufferIterForwardToContextClassToggleFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferIterForwardToContextClassToggleFunction, err = bufferObject.InvokerNew("iter_forward_to_context_class_toggle")
	})
	return err
}

// IterForwardToContextClassToggle is a representation of the C type gtk_source_buffer_iter_forward_to_context_class_toggle.
func (recv *Buffer) IterForwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(contextClass)

	var ret gi.Argument

	err := bufferIterForwardToContextClassToggleFunction_Set()
	if err == nil {
		ret = bufferIterForwardToContextClassToggleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferIterHasContextClassFunction *gi.Function
var bufferIterHasContextClassFunction_Once sync.Once

func bufferIterHasContextClassFunction_Set() error {
	var err error
	bufferIterHasContextClassFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferIterHasContextClassFunction, err = bufferObject.InvokerNew("iter_has_context_class")
	})
	return err
}

// IterHasContextClass is a representation of the C type gtk_source_buffer_iter_has_context_class.
func (recv *Buffer) IterHasContextClass(iter *gtk.TextIter, contextClass string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetString(contextClass)

	var ret gi.Argument

	err := bufferIterHasContextClassFunction_Set()
	if err == nil {
		ret = bufferIterHasContextClassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var bufferJoinLinesFunction *gi.Function
var bufferJoinLinesFunction_Once sync.Once

func bufferJoinLinesFunction_Set() error {
	var err error
	bufferJoinLinesFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferJoinLinesFunction, err = bufferObject.InvokerNew("join_lines")
	})
	return err
}

// JoinLines is a representation of the C type gtk_source_buffer_join_lines.
func (recv *Buffer) JoinLines(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := bufferJoinLinesFunction_Set()
	if err == nil {
		bufferJoinLinesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferRedoFunction *gi.Function
var bufferRedoFunction_Once sync.Once

func bufferRedoFunction_Set() error {
	var err error
	bufferRedoFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferRedoFunction, err = bufferObject.InvokerNew("redo")
	})
	return err
}

// Redo is a representation of the C type gtk_source_buffer_redo.
func (recv *Buffer) Redo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := bufferRedoFunction_Set()
	if err == nil {
		bufferRedoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferRemoveSourceMarksFunction *gi.Function
var bufferRemoveSourceMarksFunction_Once sync.Once

func bufferRemoveSourceMarksFunction_Set() error {
	var err error
	bufferRemoveSourceMarksFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferRemoveSourceMarksFunction, err = bufferObject.InvokerNew("remove_source_marks")
	})
	return err
}

// RemoveSourceMarks is a representation of the C type gtk_source_buffer_remove_source_marks.
func (recv *Buffer) RemoveSourceMarks(start *gtk.TextIter, end *gtk.TextIter, category string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())
	inArgs[3].SetString(category)

	err := bufferRemoveSourceMarksFunction_Set()
	if err == nil {
		bufferRemoveSourceMarksFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferSetHighlightMatchingBracketsFunction *gi.Function
var bufferSetHighlightMatchingBracketsFunction_Once sync.Once

func bufferSetHighlightMatchingBracketsFunction_Set() error {
	var err error
	bufferSetHighlightMatchingBracketsFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetHighlightMatchingBracketsFunction, err = bufferObject.InvokerNew("set_highlight_matching_brackets")
	})
	return err
}

// SetHighlightMatchingBrackets is a representation of the C type gtk_source_buffer_set_highlight_matching_brackets.
func (recv *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetHighlightSyntaxFunction, err = bufferObject.InvokerNew("set_highlight_syntax")
	})
	return err
}

// SetHighlightSyntax is a representation of the C type gtk_source_buffer_set_highlight_syntax.
func (recv *Buffer) SetHighlightSyntax(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetImplicitTrailingNewlineFunction, err = bufferObject.InvokerNew("set_implicit_trailing_newline")
	})
	return err
}

// SetImplicitTrailingNewline is a representation of the C type gtk_source_buffer_set_implicit_trailing_newline.
func (recv *Buffer) SetImplicitTrailingNewline(implicitTrailingNewline bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetLanguageFunction, err = bufferObject.InvokerNew("set_language")
	})
	return err
}

// SetLanguage is a representation of the C type gtk_source_buffer_set_language.
func (recv *Buffer) SetLanguage(language *Language) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(language.Native())

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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetMaxUndoLevelsFunction, err = bufferObject.InvokerNew("set_max_undo_levels")
	})
	return err
}

// SetMaxUndoLevels is a representation of the C type gtk_source_buffer_set_max_undo_levels.
func (recv *Buffer) SetMaxUndoLevels(maxUndoLevels int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSetStyleSchemeFunction, err = bufferObject.InvokerNew("set_style_scheme")
	})
	return err
}

// SetStyleScheme is a representation of the C type gtk_source_buffer_set_style_scheme.
func (recv *Buffer) SetStyleScheme(scheme *StyleScheme) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(scheme.Native())

	err := bufferSetStyleSchemeFunction_Set()
	if err == nil {
		bufferSetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_buffer_set_undo_manager' : parameter 'manager' of type 'UndoManager' not supported

var bufferSortLinesFunction *gi.Function
var bufferSortLinesFunction_Once sync.Once

func bufferSortLinesFunction_Set() error {
	var err error
	bufferSortLinesFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferSortLinesFunction, err = bufferObject.InvokerNew("sort_lines")
	})
	return err
}

// SortLines is a representation of the C type gtk_source_buffer_sort_lines.
func (recv *Buffer) SortLines(start *gtk.TextIter, end *gtk.TextIter, flags SortFlags, column int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())
	inArgs[3].SetInt32(int32(flags))
	inArgs[4].SetInt32(column)

	err := bufferSortLinesFunction_Set()
	if err == nil {
		bufferSortLinesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var bufferUndoFunction *gi.Function
var bufferUndoFunction_Once sync.Once

func bufferUndoFunction_Set() error {
	var err error
	bufferUndoFunction_Once.Do(func() {
		err = bufferObject_Set()
		if err != nil {
			return
		}
		bufferUndoFunction, err = bufferObject.InvokerNew("undo")
	})
	return err
}

// Undo is a representation of the C type gtk_source_buffer_undo.
func (recv *Buffer) Undo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := bufferUndoFunction_Set()
	if err == nil {
		bufferUndoFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectBracketMatched connects a callback to the 'bracket-matched' signal of the Buffer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Buffer) ConnectBracketMatched(handler func(instance *Buffer, iter *gtk.TextIter, state BracketMatchType)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := BufferNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextIterNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := (BracketMatchType)(object2.GetInt())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "bracket-matched", marshal)
}

/*
ConnectHighlightUpdated connects a callback to the 'highlight-updated' signal of the Buffer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Buffer) ConnectHighlightUpdated(handler func(instance *Buffer, start *gtk.TextIter, end *gtk.TextIter)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := BufferNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextIterNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := gtk.TextIterNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "highlight-updated", marshal)
}

/*
ConnectRedo connects a callback to the 'redo' signal of the Buffer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Buffer) ConnectRedo(handler func(instance *Buffer)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := BufferNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "redo", marshal)
}

/*
ConnectSourceMarkUpdated connects a callback to the 'source-mark-updated' signal of the Buffer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Buffer) ConnectSourceMarkUpdated(handler func(instance *Buffer, mark *gtk.TextMark)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := BufferNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextMarkNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "source-mark-updated", marshal)
}

/*
ConnectUndo connects a callback to the 'undo' signal of the Buffer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Buffer) ConnectUndo(handler func(instance *Buffer)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := BufferNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "undo", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Buffer) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var completionObject *gi.Object
var completionObject_Once sync.Once

func completionObject_Set() error {
	var err error
	completionObject_Once.Do(func() {
		completionObject, err = gi.ObjectNew("GtkSource", "Completion")
	})
	return err
}

type Completion struct {
	native unsafe.Pointer
}

func CompletionNewFromNative(native unsafe.Pointer) *Completion {
	instance := &Completion{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Completion) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCompletion down casts any arbitrary Object to Completion.
Exercise care, as this is a potentially dangerous function
if the Object is not a Completion.
*/
func CastToCompletion(object *gobject.Object) *Completion {
	return CompletionNewFromNative(object.Native())
}

// Equals compares this Completion with another Completion, and returns true if they represent the same Object.
func (recv *Completion) Equals(other *Completion) bool {
	return other.Native() == recv.Native()
}

func (recv *Completion) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Completion) FieldParentInstance() *gobject.Object {
	var nilValue *gobject.Object
	err := completionObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Completion) SetFieldParentInstance(value *gobject.Object) {
	err := completionObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Completion) FieldPriv() *CompletionPrivate {
	var nilValue *CompletionPrivate
	err := completionObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionObject, recv.Native(), "priv")
	value := CompletionPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Completion) SetFieldPriv(value *CompletionPrivate) {
	err := completionObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionObject, recv.Native(), "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_add_provider' : parameter 'provider' of type 'CompletionProvider' not supported

var completionBlockInteractiveFunction *gi.Function
var completionBlockInteractiveFunction_Once sync.Once

func completionBlockInteractiveFunction_Set() error {
	var err error
	completionBlockInteractiveFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionBlockInteractiveFunction, err = completionObject.InvokerNew("block_interactive")
	})
	return err
}

// BlockInteractive is a representation of the C type gtk_source_completion_block_interactive.
func (recv *Completion) BlockInteractive() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := completionBlockInteractiveFunction_Set()
	if err == nil {
		completionBlockInteractiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionCreateContextFunction *gi.Function
var completionCreateContextFunction_Once sync.Once

func completionCreateContextFunction_Set() error {
	var err error
	completionCreateContextFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionCreateContextFunction, err = completionObject.InvokerNew("create_context")
	})
	return err
}

// CreateContext is a representation of the C type gtk_source_completion_create_context.
func (recv *Completion) CreateContext(position *gtk.TextIter) *CompletionContext {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(position.Native())

	var ret gi.Argument

	err := completionCreateContextFunction_Set()
	if err == nil {
		ret = completionCreateContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionContextNewFromNative(ret.Pointer())

	return retGo
}

var completionGetInfoWindowFunction *gi.Function
var completionGetInfoWindowFunction_Once sync.Once

func completionGetInfoWindowFunction_Set() error {
	var err error
	completionGetInfoWindowFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionGetInfoWindowFunction, err = completionObject.InvokerNew("get_info_window")
	})
	return err
}

// GetInfoWindow is a representation of the C type gtk_source_completion_get_info_window.
func (recv *Completion) GetInfoWindow() *CompletionInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionGetInfoWindowFunction_Set()
	if err == nil {
		ret = completionGetInfoWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionInfoNewFromNative(ret.Pointer())

	return retGo
}

var completionGetProvidersFunction *gi.Function
var completionGetProvidersFunction_Once sync.Once

func completionGetProvidersFunction_Set() error {
	var err error
	completionGetProvidersFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionGetProvidersFunction, err = completionObject.InvokerNew("get_providers")
	})
	return err
}

// GetProviders is a representation of the C type gtk_source_completion_get_providers.
func (recv *Completion) GetProviders() *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionGetProvidersFunction_Set()
	if err == nil {
		ret = completionGetProvidersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var completionGetViewFunction *gi.Function
var completionGetViewFunction_Once sync.Once

func completionGetViewFunction_Set() error {
	var err error
	completionGetViewFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionGetViewFunction, err = completionObject.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_completion_get_view.
func (recv *Completion) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionGetViewFunction_Set()
	if err == nil {
		ret = completionGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ViewNewFromNative(ret.Pointer())

	return retGo
}

var completionHideFunction *gi.Function
var completionHideFunction_Once sync.Once

func completionHideFunction_Set() error {
	var err error
	completionHideFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionHideFunction, err = completionObject.InvokerNew("hide")
	})
	return err
}

// Hide is a representation of the C type gtk_source_completion_hide.
func (recv *Completion) Hide() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := completionHideFunction_Set()
	if err == nil {
		completionHideFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionMoveWindowFunction *gi.Function
var completionMoveWindowFunction_Once sync.Once

func completionMoveWindowFunction_Set() error {
	var err error
	completionMoveWindowFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionMoveWindowFunction, err = completionObject.InvokerNew("move_window")
	})
	return err
}

// MoveWindow is a representation of the C type gtk_source_completion_move_window.
func (recv *Completion) MoveWindow(iter *gtk.TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	err := completionMoveWindowFunction_Set()
	if err == nil {
		completionMoveWindowFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_completion_remove_provider' : parameter 'provider' of type 'CompletionProvider' not supported

var completionShowFunction *gi.Function
var completionShowFunction_Once sync.Once

func completionShowFunction_Set() error {
	var err error
	completionShowFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionShowFunction, err = completionObject.InvokerNew("show")
	})
	return err
}

// Show is a representation of the C type gtk_source_completion_show.
func (recv *Completion) Show(providers *glib.List, context *CompletionContext) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(providers.Native())
	inArgs[2].SetPointer(context.Native())

	var ret gi.Argument

	err := completionShowFunction_Set()
	if err == nil {
		ret = completionShowFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var completionUnblockInteractiveFunction *gi.Function
var completionUnblockInteractiveFunction_Once sync.Once

func completionUnblockInteractiveFunction_Set() error {
	var err error
	completionUnblockInteractiveFunction_Once.Do(func() {
		err = completionObject_Set()
		if err != nil {
			return
		}
		completionUnblockInteractiveFunction, err = completionObject.InvokerNew("unblock_interactive")
	})
	return err
}

// UnblockInteractive is a representation of the C type gtk_source_completion_unblock_interactive.
func (recv *Completion) UnblockInteractive() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := completionUnblockInteractiveFunction_Set()
	if err == nil {
		completionUnblockInteractiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectActivateProposal connects a callback to the 'activate-proposal' signal of the Completion.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Completion) ConnectActivateProposal(handler func(instance *Completion)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "activate-proposal", marshal)
}

/*
ConnectHide connects a callback to the 'hide' signal of the Completion.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Completion) ConnectHide(handler func(instance *Completion)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "hide", marshal)
}

// UNSUPPORTED : C value 'move-cursor' : parameter 'step' of type 'Gtk.ScrollStep' not supported

// UNSUPPORTED : C value 'move-page' : parameter 'step' of type 'Gtk.ScrollStep' not supported

/*
ConnectPopulateContext connects a callback to the 'populate-context' signal of the Completion.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Completion) ConnectPopulateContext(handler func(instance *Completion, context *CompletionContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := CompletionContextNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "populate-context", marshal)
}

/*
ConnectShow connects a callback to the 'show' signal of the Completion.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Completion) ConnectShow(handler func(instance *Completion)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "show", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Completion) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// GtkBuildable returns the Gtk.Buildable interface implemented by Completion
func (recv *Completion) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

var completionContextObject *gi.Object
var completionContextObject_Once sync.Once

func completionContextObject_Set() error {
	var err error
	completionContextObject_Once.Do(func() {
		completionContextObject, err = gi.ObjectNew("GtkSource", "CompletionContext")
	})
	return err
}

type CompletionContext struct {
	native unsafe.Pointer
}

func CompletionContextNewFromNative(native unsafe.Pointer) *CompletionContext {
	instance := &CompletionContext{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CompletionContext) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *CompletionContext) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCompletionContext down casts any arbitrary Object to CompletionContext.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionContext.
*/
func CastToCompletionContext(object *gobject.Object) *CompletionContext {
	return CompletionContextNewFromNative(object.Native())
}

// Equals compares this CompletionContext with another CompletionContext, and returns true if they represent the same Object.
func (recv *CompletionContext) Equals(other *CompletionContext) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionContext) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CompletionContext) FieldParent() *gobject.InitiallyUnowned {
	var nilValue *gobject.InitiallyUnowned
	err := completionContextObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionContextObject, recv.Native(), "parent")
	value := gobject.InitiallyUnownedNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CompletionContext) SetFieldParent(value *gobject.InitiallyUnowned) {
	err := completionContextObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionContextObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *CompletionContext) FieldPriv() *CompletionContextPrivate {
	var nilValue *CompletionContextPrivate
	err := completionContextObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionContextObject, recv.Native(), "priv")
	value := CompletionContextPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionContext) SetFieldPriv(value *CompletionContextPrivate) {
	err := completionContextObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionContextObject, recv.Native(), "priv", argValue)
}

// UNSUPPORTED : C value 'gtk_source_completion_context_add_proposals' : parameter 'provider' of type 'CompletionProvider' not supported

var completionContextGetActivationFunction *gi.Function
var completionContextGetActivationFunction_Once sync.Once

func completionContextGetActivationFunction_Set() error {
	var err error
	completionContextGetActivationFunction_Once.Do(func() {
		err = completionContextObject_Set()
		if err != nil {
			return
		}
		completionContextGetActivationFunction, err = completionContextObject.InvokerNew("get_activation")
	})
	return err
}

// GetActivation is a representation of the C type gtk_source_completion_context_get_activation.
func (recv *CompletionContext) GetActivation() CompletionActivation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionContextGetActivationFunction_Set()
	if err == nil {
		ret = completionContextGetActivationFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionActivation(ret.Int32())

	return retGo
}

var completionContextGetIterFunction *gi.Function
var completionContextGetIterFunction_Once sync.Once

func completionContextGetIterFunction_Set() error {
	var err error
	completionContextGetIterFunction_Once.Do(func() {
		err = completionContextObject_Set()
		if err != nil {
			return
		}
		completionContextGetIterFunction, err = completionContextObject.InvokerNew("get_iter")
	})
	return err
}

// GetIter is a representation of the C type gtk_source_completion_context_get_iter.
func (recv *CompletionContext) GetIter() (bool, *gtk.TextIter) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := completionContextGetIterFunction_Set()
	if err == nil {
		ret = completionContextGetIterFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

/*
ConnectCancelled connects a callback to the 'cancelled' signal of the CompletionContext.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *CompletionContext) ConnectCancelled(handler func(instance *CompletionContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionContextNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "cancelled", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *CompletionContext) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var completionInfoObject *gi.Object
var completionInfoObject_Once sync.Once

func completionInfoObject_Set() error {
	var err error
	completionInfoObject_Once.Do(func() {
		completionInfoObject, err = gi.ObjectNew("GtkSource", "CompletionInfo")
	})
	return err
}

type CompletionInfo struct {
	native unsafe.Pointer
}

func CompletionInfoNewFromNative(native unsafe.Pointer) *CompletionInfo {
	instance := &CompletionInfo{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Window upcasts to *Window
func (recv *CompletionInfo) Window() *gtk.Window {
	return gtk.WindowNewFromNative(recv.Native())
}

// Bin upcasts to *Bin
func (recv *CompletionInfo) Bin() *gtk.Bin {
	return gtk.BinNewFromNative(recv.Native())
}

// Container upcasts to *Container
func (recv *CompletionInfo) Container() *gtk.Container {
	return gtk.ContainerNewFromNative(recv.Native())
}

// Widget upcasts to *Widget
func (recv *CompletionInfo) Widget() *gtk.Widget {
	return gtk.WidgetNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CompletionInfo) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *CompletionInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCompletionInfo down casts any arbitrary Object to CompletionInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionInfo.
*/
func CastToCompletionInfo(object *gobject.Object) *CompletionInfo {
	return CompletionInfoNewFromNative(object.Native())
}

// Equals compares this CompletionInfo with another CompletionInfo, and returns true if they represent the same Object.
func (recv *CompletionInfo) Equals(other *CompletionInfo) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionInfo) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CompletionInfo) FieldParent() *gtk.Window {
	var nilValue *gtk.Window
	err := completionInfoObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionInfoObject, recv.Native(), "parent")
	value := gtk.WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CompletionInfo) SetFieldParent(value *gtk.Window) {
	err := completionInfoObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionInfoObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *CompletionInfo) FieldPriv() *CompletionInfoPrivate {
	var nilValue *CompletionInfoPrivate
	err := completionInfoObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionInfoObject, recv.Native(), "priv")
	value := CompletionInfoPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionInfo) SetFieldPriv(value *CompletionInfoPrivate) {
	err := completionInfoObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionInfoObject, recv.Native(), "priv", argValue)
}

var completionInfoNewFunction *gi.Function
var completionInfoNewFunction_Once sync.Once

func completionInfoNewFunction_Set() error {
	var err error
	completionInfoNewFunction_Once.Do(func() {
		err = completionInfoObject_Set()
		if err != nil {
			return
		}
		completionInfoNewFunction, err = completionInfoObject.InvokerNew("new")
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

	retGo := CompletionInfoNewFromNative(ret.Pointer())

	return retGo
}

var completionInfoGetWidgetFunction *gi.Function
var completionInfoGetWidgetFunction_Once sync.Once

func completionInfoGetWidgetFunction_Set() error {
	var err error
	completionInfoGetWidgetFunction_Once.Do(func() {
		err = completionInfoObject_Set()
		if err != nil {
			return
		}
		completionInfoGetWidgetFunction, err = completionInfoObject.InvokerNew("get_widget")
	})
	return err
}

// GetWidget is a representation of the C type gtk_source_completion_info_get_widget.
func (recv *CompletionInfo) GetWidget() *gtk.Widget {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionInfoGetWidgetFunction_Set()
	if err == nil {
		ret = completionInfoGetWidgetFunction.Invoke(inArgs[:], nil)
	}

	retGo := gtk.WidgetNewFromNative(ret.Pointer())

	return retGo
}

var completionInfoMoveToIterFunction *gi.Function
var completionInfoMoveToIterFunction_Once sync.Once

func completionInfoMoveToIterFunction_Set() error {
	var err error
	completionInfoMoveToIterFunction_Once.Do(func() {
		err = completionInfoObject_Set()
		if err != nil {
			return
		}
		completionInfoMoveToIterFunction, err = completionInfoObject.InvokerNew("move_to_iter")
	})
	return err
}

// MoveToIter is a representation of the C type gtk_source_completion_info_move_to_iter.
func (recv *CompletionInfo) MoveToIter(view *gtk.TextView, iter *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(view.Native())
	inArgs[2].SetPointer(iter.Native())

	err := completionInfoMoveToIterFunction_Set()
	if err == nil {
		completionInfoMoveToIterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionInfoSetWidgetFunction *gi.Function
var completionInfoSetWidgetFunction_Once sync.Once

func completionInfoSetWidgetFunction_Set() error {
	var err error
	completionInfoSetWidgetFunction_Once.Do(func() {
		err = completionInfoObject_Set()
		if err != nil {
			return
		}
		completionInfoSetWidgetFunction, err = completionInfoObject.InvokerNew("set_widget")
	})
	return err
}

// SetWidget is a representation of the C type gtk_source_completion_info_set_widget.
func (recv *CompletionInfo) SetWidget(widget *gtk.Widget) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(widget.Native())

	err := completionInfoSetWidgetFunction_Set()
	if err == nil {
		completionInfoSetWidgetFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectBeforeShow connects a callback to the 'before-show' signal of the CompletionInfo.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *CompletionInfo) ConnectBeforeShow(handler func(instance *CompletionInfo)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CompletionInfoNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "before-show", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *CompletionInfo) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// AtkImplementorIface returns the Atk.ImplementorIface interface implemented by CompletionInfo
func (recv *CompletionInfo) AtkImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromNative(recv.Native())
}

// GtkBuildable returns the Gtk.Buildable interface implemented by CompletionInfo
func (recv *CompletionInfo) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

var completionItemObject *gi.Object
var completionItemObject_Once sync.Once

func completionItemObject_Set() error {
	var err error
	completionItemObject_Once.Do(func() {
		completionItemObject, err = gi.ObjectNew("GtkSource", "CompletionItem")
	})
	return err
}

type CompletionItem struct {
	native unsafe.Pointer
}

func CompletionItemNewFromNative(native unsafe.Pointer) *CompletionItem {
	instance := &CompletionItem{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *CompletionItem) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCompletionItem down casts any arbitrary Object to CompletionItem.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionItem.
*/
func CastToCompletionItem(object *gobject.Object) *CompletionItem {
	return CompletionItemNewFromNative(object.Native())
}

// Equals compares this CompletionItem with another CompletionItem, and returns true if they represent the same Object.
func (recv *CompletionItem) Equals(other *CompletionItem) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionItem) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CompletionItem) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := completionItemObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionItemObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CompletionItem) SetFieldParent(value *gobject.Object) {
	err := completionItemObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionItemObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *CompletionItem) FieldPriv() *CompletionItemPrivate {
	var nilValue *CompletionItemPrivate
	err := completionItemObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionItemObject, recv.Native(), "priv")
	value := CompletionItemPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionItem) SetFieldPriv(value *CompletionItemPrivate) {
	err := completionItemObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionItemObject, recv.Native(), "priv", argValue)
}

var completionItemNewFunction *gi.Function
var completionItemNewFunction_Once sync.Once

func completionItemNewFunction_Set() error {
	var err error
	completionItemNewFunction_Once.Do(func() {
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemNewFunction, err = completionItemObject.InvokerNew("new")
	})
	return err
}

// CompletionItemNew is a representation of the C type gtk_source_completion_item_new.
func CompletionItemNew(label string, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(label)
	inArgs[1].SetString(text)
	inArgs[2].SetPointer(icon.Native())
	inArgs[3].SetString(info)

	var ret gi.Argument

	err := completionItemNewFunction_Set()
	if err == nil {
		ret = completionItemNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var completionItemNewFromStockFunction *gi.Function
var completionItemNewFromStockFunction_Once sync.Once

func completionItemNewFromStockFunction_Set() error {
	var err error
	completionItemNewFromStockFunction_Once.Do(func() {
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemNewFromStockFunction, err = completionItemObject.InvokerNew("new_from_stock")
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

	retGo := CompletionItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var completionItemNewWithMarkupFunction *gi.Function
var completionItemNewWithMarkupFunction_Once sync.Once

func completionItemNewWithMarkupFunction_Set() error {
	var err error
	completionItemNewWithMarkupFunction_Once.Do(func() {
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemNewWithMarkupFunction, err = completionItemObject.InvokerNew("new_with_markup")
	})
	return err
}

// CompletionItemNewWithMarkup is a representation of the C type gtk_source_completion_item_new_with_markup.
func CompletionItemNewWithMarkup(markup string, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(markup)
	inArgs[1].SetString(text)
	inArgs[2].SetPointer(icon.Native())
	inArgs[3].SetString(info)

	var ret gi.Argument

	err := completionItemNewWithMarkupFunction_Set()
	if err == nil {
		ret = completionItemNewWithMarkupFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionItemNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_item_set_gicon' : parameter 'gicon' of type 'Gio.Icon' not supported

var completionItemSetIconFunction *gi.Function
var completionItemSetIconFunction_Once sync.Once

func completionItemSetIconFunction_Set() error {
	var err error
	completionItemSetIconFunction_Once.Do(func() {
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetIconFunction, err = completionItemObject.InvokerNew("set_icon")
	})
	return err
}

// SetIcon is a representation of the C type gtk_source_completion_item_set_icon.
func (recv *CompletionItem) SetIcon(icon *gdkpixbuf.Pixbuf) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(icon.Native())

	err := completionItemSetIconFunction_Set()
	if err == nil {
		completionItemSetIconFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionItemSetIconNameFunction *gi.Function
var completionItemSetIconNameFunction_Once sync.Once

func completionItemSetIconNameFunction_Set() error {
	var err error
	completionItemSetIconNameFunction_Once.Do(func() {
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetIconNameFunction, err = completionItemObject.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_completion_item_set_icon_name.
func (recv *CompletionItem) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetInfoFunction, err = completionItemObject.InvokerNew("set_info")
	})
	return err
}

// SetInfo is a representation of the C type gtk_source_completion_item_set_info.
func (recv *CompletionItem) SetInfo(info string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetLabelFunction, err = completionItemObject.InvokerNew("set_label")
	})
	return err
}

// SetLabel is a representation of the C type gtk_source_completion_item_set_label.
func (recv *CompletionItem) SetLabel(label string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetMarkupFunction, err = completionItemObject.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type gtk_source_completion_item_set_markup.
func (recv *CompletionItem) SetMarkup(markup string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = completionItemObject_Set()
		if err != nil {
			return
		}
		completionItemSetTextFunction, err = completionItemObject.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type gtk_source_completion_item_set_text.
func (recv *CompletionItem) SetText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)

	err := completionItemSetTextFunction_Set()
	if err == nil {
		completionItemSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

// CompletionProposal returns the CompletionProposal interface implemented by CompletionItem
func (recv *CompletionItem) CompletionProposal() *CompletionProposal {
	return CompletionProposalNewFromNative(recv.Native())
}

var completionWordsObject *gi.Object
var completionWordsObject_Once sync.Once

func completionWordsObject_Set() error {
	var err error
	completionWordsObject_Once.Do(func() {
		completionWordsObject, err = gi.ObjectNew("GtkSource", "CompletionWords")
	})
	return err
}

type CompletionWords struct {
	native unsafe.Pointer
}

func CompletionWordsNewFromNative(native unsafe.Pointer) *CompletionWords {
	instance := &CompletionWords{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *CompletionWords) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToCompletionWords down casts any arbitrary Object to CompletionWords.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionWords.
*/
func CastToCompletionWords(object *gobject.Object) *CompletionWords {
	return CompletionWordsNewFromNative(object.Native())
}

// Equals compares this CompletionWords with another CompletionWords, and returns true if they represent the same Object.
func (recv *CompletionWords) Equals(other *CompletionWords) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionWords) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *CompletionWords) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := completionWordsObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionWordsObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CompletionWords) SetFieldParent(value *gobject.Object) {
	err := completionWordsObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionWordsObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *CompletionWords) FieldPriv() *CompletionWordsPrivate {
	var nilValue *CompletionWordsPrivate
	err := completionWordsObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(completionWordsObject, recv.Native(), "priv")
	value := CompletionWordsPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *CompletionWords) SetFieldPriv(value *CompletionWordsPrivate) {
	err := completionWordsObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(completionWordsObject, recv.Native(), "priv", argValue)
}

var completionWordsNewFunction *gi.Function
var completionWordsNewFunction_Once sync.Once

func completionWordsNewFunction_Set() error {
	var err error
	completionWordsNewFunction_Once.Do(func() {
		err = completionWordsObject_Set()
		if err != nil {
			return
		}
		completionWordsNewFunction, err = completionWordsObject.InvokerNew("new")
	})
	return err
}

// CompletionWordsNew is a representation of the C type gtk_source_completion_words_new.
func CompletionWordsNew(name string, icon *gdkpixbuf.Pixbuf) *CompletionWords {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(icon.Native())

	var ret gi.Argument

	err := completionWordsNewFunction_Set()
	if err == nil {
		ret = completionWordsNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionWordsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var completionWordsRegisterFunction *gi.Function
var completionWordsRegisterFunction_Once sync.Once

func completionWordsRegisterFunction_Set() error {
	var err error
	completionWordsRegisterFunction_Once.Do(func() {
		err = completionWordsObject_Set()
		if err != nil {
			return
		}
		completionWordsRegisterFunction, err = completionWordsObject.InvokerNew("register")
	})
	return err
}

// Register is a representation of the C type gtk_source_completion_words_register.
func (recv *CompletionWords) Register(buffer *gtk.TextBuffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(buffer.Native())

	err := completionWordsRegisterFunction_Set()
	if err == nil {
		completionWordsRegisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var completionWordsUnregisterFunction *gi.Function
var completionWordsUnregisterFunction_Once sync.Once

func completionWordsUnregisterFunction_Set() error {
	var err error
	completionWordsUnregisterFunction_Once.Do(func() {
		err = completionWordsObject_Set()
		if err != nil {
			return
		}
		completionWordsUnregisterFunction, err = completionWordsObject.InvokerNew("unregister")
	})
	return err
}

// Unregister is a representation of the C type gtk_source_completion_words_unregister.
func (recv *CompletionWords) Unregister(buffer *gtk.TextBuffer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(buffer.Native())

	err := completionWordsUnregisterFunction_Set()
	if err == nil {
		completionWordsUnregisterFunction.Invoke(inArgs[:], nil)
	}

	return
}

// CompletionProvider returns the CompletionProvider interface implemented by CompletionWords
func (recv *CompletionWords) CompletionProvider() *CompletionProvider {
	return CompletionProviderNewFromNative(recv.Native())
}

var fileObject *gi.Object
var fileObject_Once sync.Once

func fileObject_Set() error {
	var err error
	fileObject_Once.Do(func() {
		fileObject, err = gi.ObjectNew("GtkSource", "File")
	})
	return err
}

type File struct {
	native unsafe.Pointer
}

func FileNewFromNative(native unsafe.Pointer) *File {
	instance := &File{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *File) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFile down casts any arbitrary Object to File.
Exercise care, as this is a potentially dangerous function
if the Object is not a File.
*/
func CastToFile(object *gobject.Object) *File {
	return FileNewFromNative(object.Native())
}

// Equals compares this File with another File, and returns true if they represent the same Object.
func (recv *File) Equals(other *File) bool {
	return other.Native() == recv.Native()
}

func (recv *File) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *File) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := fileObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *File) SetFieldParent(value *gobject.Object) {
	err := fileObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *File) FieldPriv() *FilePrivate {
	var nilValue *FilePrivate
	err := fileObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileObject, recv.Native(), "priv")
	value := FilePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *File) SetFieldPriv(value *FilePrivate) {
	err := fileObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileObject, recv.Native(), "priv", argValue)
}

var fileNewFunction *gi.Function
var fileNewFunction_Once sync.Once

func fileNewFunction_Set() error {
	var err error
	fileNewFunction_Once.Do(func() {
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileNewFunction, err = fileObject.InvokerNew("new")
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

	retGo := FileNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var fileCheckFileOnDiskFunction *gi.Function
var fileCheckFileOnDiskFunction_Once sync.Once

func fileCheckFileOnDiskFunction_Set() error {
	var err error
	fileCheckFileOnDiskFunction_Once.Do(func() {
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileCheckFileOnDiskFunction, err = fileObject.InvokerNew("check_file_on_disk")
	})
	return err
}

// CheckFileOnDisk is a representation of the C type gtk_source_file_check_file_on_disk.
func (recv *File) CheckFileOnDisk() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileGetCompressionTypeFunction, err = fileObject.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_get_compression_type.
func (recv *File) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileGetEncodingFunction, err = fileObject.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_get_encoding.
func (recv *File) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileGetEncodingFunction_Set()
	if err == nil {
		ret = fileGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := EncodingNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_get_location' : return type 'Gio.File' not supported

var fileGetNewlineTypeFunction *gi.Function
var fileGetNewlineTypeFunction_Once sync.Once

func fileGetNewlineTypeFunction_Set() error {
	var err error
	fileGetNewlineTypeFunction_Once.Do(func() {
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileGetNewlineTypeFunction, err = fileObject.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_get_newline_type.
func (recv *File) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileIsDeletedFunction, err = fileObject.InvokerNew("is_deleted")
	})
	return err
}

// IsDeleted is a representation of the C type gtk_source_file_is_deleted.
func (recv *File) IsDeleted() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileIsExternallyModifiedFunction, err = fileObject.InvokerNew("is_externally_modified")
	})
	return err
}

// IsExternallyModified is a representation of the C type gtk_source_file_is_externally_modified.
func (recv *File) IsExternallyModified() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileIsLocalFunction, err = fileObject.InvokerNew("is_local")
	})
	return err
}

// IsLocal is a representation of the C type gtk_source_file_is_local.
func (recv *File) IsLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileObject_Set()
		if err != nil {
			return
		}
		fileIsReadonlyFunction, err = fileObject.InvokerNew("is_readonly")
	})
	return err
}

// IsReadonly is a representation of the C type gtk_source_file_is_readonly.
func (recv *File) IsReadonly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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

var fileLoaderObject *gi.Object
var fileLoaderObject_Once sync.Once

func fileLoaderObject_Set() error {
	var err error
	fileLoaderObject_Once.Do(func() {
		fileLoaderObject, err = gi.ObjectNew("GtkSource", "FileLoader")
	})
	return err
}

type FileLoader struct {
	native unsafe.Pointer
}

func FileLoaderNewFromNative(native unsafe.Pointer) *FileLoader {
	instance := &FileLoader{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileLoader down casts any arbitrary Object to FileLoader.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileLoader.
*/
func CastToFileLoader(object *gobject.Object) *FileLoader {
	return FileLoaderNewFromNative(object.Native())
}

// Equals compares this FileLoader with another FileLoader, and returns true if they represent the same Object.
func (recv *FileLoader) Equals(other *FileLoader) bool {
	return other.Native() == recv.Native()
}

func (recv *FileLoader) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *FileLoader) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := fileLoaderObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileLoaderObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *FileLoader) SetFieldParent(value *gobject.Object) {
	err := fileLoaderObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileLoaderObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *FileLoader) FieldPriv() *FileLoaderPrivate {
	var nilValue *FileLoaderPrivate
	err := fileLoaderObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileLoaderObject, recv.Native(), "priv")
	value := FileLoaderPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *FileLoader) SetFieldPriv(value *FileLoaderPrivate) {
	err := fileLoaderObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileLoaderObject, recv.Native(), "priv", argValue)
}

var fileLoaderNewFunction *gi.Function
var fileLoaderNewFunction_Once sync.Once

func fileLoaderNewFunction_Set() error {
	var err error
	fileLoaderNewFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderNewFunction, err = fileLoaderObject.InvokerNew("new")
	})
	return err
}

// FileLoaderNew is a representation of the C type gtk_source_file_loader_new.
func FileLoaderNew(buffer *Buffer, file *File) *FileLoader {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native())
	inArgs[1].SetPointer(file.Native())

	var ret gi.Argument

	err := fileLoaderNewFunction_Set()
	if err == nil {
		ret = fileLoaderNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileLoaderNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var fileLoaderNewFromStreamFunction *gi.Function
var fileLoaderNewFromStreamFunction_Once sync.Once

func fileLoaderNewFromStreamFunction_Set() error {
	var err error
	fileLoaderNewFromStreamFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderNewFromStreamFunction, err = fileLoaderObject.InvokerNew("new_from_stream")
	})
	return err
}

// FileLoaderNewFromStream is a representation of the C type gtk_source_file_loader_new_from_stream.
func FileLoaderNewFromStream(buffer *Buffer, file *File, stream *gio.InputStream) *FileLoader {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(buffer.Native())
	inArgs[1].SetPointer(file.Native())
	inArgs[2].SetPointer(stream.Native())

	var ret gi.Argument

	err := fileLoaderNewFromStreamFunction_Set()
	if err == nil {
		ret = fileLoaderNewFromStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileLoaderNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var fileLoaderGetBufferFunction *gi.Function
var fileLoaderGetBufferFunction_Once sync.Once

func fileLoaderGetBufferFunction_Set() error {
	var err error
	fileLoaderGetBufferFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetBufferFunction, err = fileLoaderObject.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_file_loader_get_buffer.
func (recv *FileLoader) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileLoaderGetBufferFunction_Set()
	if err == nil {
		ret = fileLoaderGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

	return retGo
}

var fileLoaderGetCompressionTypeFunction *gi.Function
var fileLoaderGetCompressionTypeFunction_Once sync.Once

func fileLoaderGetCompressionTypeFunction_Set() error {
	var err error
	fileLoaderGetCompressionTypeFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetCompressionTypeFunction, err = fileLoaderObject.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_loader_get_compression_type.
func (recv *FileLoader) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetEncodingFunction, err = fileLoaderObject.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_loader_get_encoding.
func (recv *FileLoader) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileLoaderGetEncodingFunction_Set()
	if err == nil {
		ret = fileLoaderGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := EncodingNewFromNative(ret.Pointer())

	return retGo
}

var fileLoaderGetFileFunction *gi.Function
var fileLoaderGetFileFunction_Once sync.Once

func fileLoaderGetFileFunction_Set() error {
	var err error
	fileLoaderGetFileFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetFileFunction, err = fileLoaderObject.InvokerNew("get_file")
	})
	return err
}

// GetFile is a representation of the C type gtk_source_file_loader_get_file.
func (recv *FileLoader) GetFile() *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileLoaderGetFileFunction_Set()
	if err == nil {
		ret = fileLoaderGetFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileLoaderGetInputStreamFunction *gi.Function
var fileLoaderGetInputStreamFunction_Once sync.Once

func fileLoaderGetInputStreamFunction_Set() error {
	var err error
	fileLoaderGetInputStreamFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetInputStreamFunction, err = fileLoaderObject.InvokerNew("get_input_stream")
	})
	return err
}

// GetInputStream is a representation of the C type gtk_source_file_loader_get_input_stream.
func (recv *FileLoader) GetInputStream() *gio.InputStream {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileLoaderGetInputStreamFunction_Set()
	if err == nil {
		ret = fileLoaderGetInputStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := gio.InputStreamNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_get_location' : return type 'Gio.File' not supported

var fileLoaderGetNewlineTypeFunction *gi.Function
var fileLoaderGetNewlineTypeFunction_Once sync.Once

func fileLoaderGetNewlineTypeFunction_Set() error {
	var err error
	fileLoaderGetNewlineTypeFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderGetNewlineTypeFunction, err = fileLoaderObject.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_loader_get_newline_type.
func (recv *FileLoader) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileLoaderGetNewlineTypeFunction_Set()
	if err == nil {
		ret = fileLoaderGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NewlineType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_loader_load_async' : parameter 'progress_callback' of type 'Gio.FileProgressCallback' not supported

// UNSUPPORTED : C value 'gtk_source_file_loader_load_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var fileLoaderSetCandidateEncodingsFunction *gi.Function
var fileLoaderSetCandidateEncodingsFunction_Once sync.Once

func fileLoaderSetCandidateEncodingsFunction_Set() error {
	var err error
	fileLoaderSetCandidateEncodingsFunction_Once.Do(func() {
		err = fileLoaderObject_Set()
		if err != nil {
			return
		}
		fileLoaderSetCandidateEncodingsFunction, err = fileLoaderObject.InvokerNew("set_candidate_encodings")
	})
	return err
}

// SetCandidateEncodings is a representation of the C type gtk_source_file_loader_set_candidate_encodings.
func (recv *FileLoader) SetCandidateEncodings(candidateEncodings *glib.SList) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(candidateEncodings.Native())

	err := fileLoaderSetCandidateEncodingsFunction_Set()
	if err == nil {
		fileLoaderSetCandidateEncodingsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileSaverObject *gi.Object
var fileSaverObject_Once sync.Once

func fileSaverObject_Set() error {
	var err error
	fileSaverObject_Once.Do(func() {
		fileSaverObject, err = gi.ObjectNew("GtkSource", "FileSaver")
	})
	return err
}

type FileSaver struct {
	native unsafe.Pointer
}

func FileSaverNewFromNative(native unsafe.Pointer) *FileSaver {
	instance := &FileSaver{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FileSaver) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFileSaver down casts any arbitrary Object to FileSaver.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileSaver.
*/
func CastToFileSaver(object *gobject.Object) *FileSaver {
	return FileSaverNewFromNative(object.Native())
}

// Equals compares this FileSaver with another FileSaver, and returns true if they represent the same Object.
func (recv *FileSaver) Equals(other *FileSaver) bool {
	return other.Native() == recv.Native()
}

func (recv *FileSaver) Native() unsafe.Pointer {
	return recv.native
}

// FieldObject returns the C field 'object'.
func (recv *FileSaver) FieldObject() *gobject.Object {
	var nilValue *gobject.Object
	err := fileSaverObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileSaverObject, recv.Native(), "object")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldObject sets the value of the C field 'object'.
func (recv *FileSaver) SetFieldObject(value *gobject.Object) {
	err := fileSaverObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileSaverObject, recv.Native(), "object", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *FileSaver) FieldPriv() *FileSaverPrivate {
	var nilValue *FileSaverPrivate
	err := fileSaverObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(fileSaverObject, recv.Native(), "priv")
	value := FileSaverPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *FileSaver) SetFieldPriv(value *FileSaverPrivate) {
	err := fileSaverObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fileSaverObject, recv.Native(), "priv", argValue)
}

var fileSaverNewFunction *gi.Function
var fileSaverNewFunction_Once sync.Once

func fileSaverNewFunction_Set() error {
	var err error
	fileSaverNewFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverNewFunction, err = fileSaverObject.InvokerNew("new")
	})
	return err
}

// FileSaverNew is a representation of the C type gtk_source_file_saver_new.
func FileSaverNew(buffer *Buffer, file *File) *FileSaver {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native())
	inArgs[1].SetPointer(file.Native())

	var ret gi.Argument

	err := fileSaverNewFunction_Set()
	if err == nil {
		ret = fileSaverNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileSaverNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_new_with_target' : parameter 'target_location' of type 'Gio.File' not supported

var fileSaverGetBufferFunction *gi.Function
var fileSaverGetBufferFunction_Once sync.Once

func fileSaverGetBufferFunction_Set() error {
	var err error
	fileSaverGetBufferFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetBufferFunction, err = fileSaverObject.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_file_saver_get_buffer.
func (recv *FileSaver) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSaverGetBufferFunction_Set()
	if err == nil {
		ret = fileSaverGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

	return retGo
}

var fileSaverGetCompressionTypeFunction *gi.Function
var fileSaverGetCompressionTypeFunction_Once sync.Once

func fileSaverGetCompressionTypeFunction_Set() error {
	var err error
	fileSaverGetCompressionTypeFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetCompressionTypeFunction, err = fileSaverObject.InvokerNew("get_compression_type")
	})
	return err
}

// GetCompressionType is a representation of the C type gtk_source_file_saver_get_compression_type.
func (recv *FileSaver) GetCompressionType() CompressionType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetEncodingFunction, err = fileSaverObject.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type gtk_source_file_saver_get_encoding.
func (recv *FileSaver) GetEncoding() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSaverGetEncodingFunction_Set()
	if err == nil {
		ret = fileSaverGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := EncodingNewFromNative(ret.Pointer())

	return retGo
}

var fileSaverGetFileFunction *gi.Function
var fileSaverGetFileFunction_Once sync.Once

func fileSaverGetFileFunction_Set() error {
	var err error
	fileSaverGetFileFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetFileFunction, err = fileSaverObject.InvokerNew("get_file")
	})
	return err
}

// GetFile is a representation of the C type gtk_source_file_saver_get_file.
func (recv *FileSaver) GetFile() *File {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSaverGetFileFunction_Set()
	if err == nil {
		ret = fileSaverGetFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileNewFromNative(ret.Pointer())

	return retGo
}

var fileSaverGetFlagsFunction *gi.Function
var fileSaverGetFlagsFunction_Once sync.Once

func fileSaverGetFlagsFunction_Set() error {
	var err error
	fileSaverGetFlagsFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetFlagsFunction, err = fileSaverObject.InvokerNew("get_flags")
	})
	return err
}

// GetFlags is a representation of the C type gtk_source_file_saver_get_flags.
func (recv *FileSaver) GetFlags() FileSaverFlags {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSaverGetFlagsFunction_Set()
	if err == nil {
		ret = fileSaverGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileSaverFlags(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_get_location' : return type 'Gio.File' not supported

var fileSaverGetNewlineTypeFunction *gi.Function
var fileSaverGetNewlineTypeFunction_Once sync.Once

func fileSaverGetNewlineTypeFunction_Set() error {
	var err error
	fileSaverGetNewlineTypeFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverGetNewlineTypeFunction, err = fileSaverObject.InvokerNew("get_newline_type")
	})
	return err
}

// GetNewlineType is a representation of the C type gtk_source_file_saver_get_newline_type.
func (recv *FileSaver) GetNewlineType() NewlineType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileSaverGetNewlineTypeFunction_Set()
	if err == nil {
		ret = fileSaverGetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NewlineType(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_file_saver_save_async' : parameter 'progress_callback' of type 'Gio.FileProgressCallback' not supported

// UNSUPPORTED : C value 'gtk_source_file_saver_save_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

var fileSaverSetCompressionTypeFunction *gi.Function
var fileSaverSetCompressionTypeFunction_Once sync.Once

func fileSaverSetCompressionTypeFunction_Set() error {
	var err error
	fileSaverSetCompressionTypeFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverSetCompressionTypeFunction, err = fileSaverObject.InvokerNew("set_compression_type")
	})
	return err
}

// SetCompressionType is a representation of the C type gtk_source_file_saver_set_compression_type.
func (recv *FileSaver) SetCompressionType(compressionType CompressionType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverSetEncodingFunction, err = fileSaverObject.InvokerNew("set_encoding")
	})
	return err
}

// SetEncoding is a representation of the C type gtk_source_file_saver_set_encoding.
func (recv *FileSaver) SetEncoding(encoding *Encoding) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(encoding.Native())

	err := fileSaverSetEncodingFunction_Set()
	if err == nil {
		fileSaverSetEncodingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileSaverSetFlagsFunction *gi.Function
var fileSaverSetFlagsFunction_Once sync.Once

func fileSaverSetFlagsFunction_Set() error {
	var err error
	fileSaverSetFlagsFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverSetFlagsFunction, err = fileSaverObject.InvokerNew("set_flags")
	})
	return err
}

// SetFlags is a representation of the C type gtk_source_file_saver_set_flags.
func (recv *FileSaver) SetFlags(flags FileSaverFlags) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(flags))

	err := fileSaverSetFlagsFunction_Set()
	if err == nil {
		fileSaverSetFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileSaverSetNewlineTypeFunction *gi.Function
var fileSaverSetNewlineTypeFunction_Once sync.Once

func fileSaverSetNewlineTypeFunction_Set() error {
	var err error
	fileSaverSetNewlineTypeFunction_Once.Do(func() {
		err = fileSaverObject_Set()
		if err != nil {
			return
		}
		fileSaverSetNewlineTypeFunction, err = fileSaverObject.InvokerNew("set_newline_type")
	})
	return err
}

// SetNewlineType is a representation of the C type gtk_source_file_saver_set_newline_type.
func (recv *FileSaver) SetNewlineType(newlineType NewlineType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(newlineType))

	err := fileSaverSetNewlineTypeFunction_Set()
	if err == nil {
		fileSaverSetNewlineTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterObject *gi.Object
var gutterObject_Once sync.Once

func gutterObject_Set() error {
	var err error
	gutterObject_Once.Do(func() {
		gutterObject, err = gi.ObjectNew("GtkSource", "Gutter")
	})
	return err
}

type Gutter struct {
	native unsafe.Pointer
}

func GutterNewFromNative(native unsafe.Pointer) *Gutter {
	instance := &Gutter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Gutter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToGutter down casts any arbitrary Object to Gutter.
Exercise care, as this is a potentially dangerous function
if the Object is not a Gutter.
*/
func CastToGutter(object *gobject.Object) *Gutter {
	return GutterNewFromNative(object.Native())
}

// Equals compares this Gutter with another Gutter, and returns true if they represent the same Object.
func (recv *Gutter) Equals(other *Gutter) bool {
	return other.Native() == recv.Native()
}

func (recv *Gutter) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Gutter) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := gutterObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(gutterObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Gutter) SetFieldParent(value *gobject.Object) {
	err := gutterObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(gutterObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Gutter) FieldPriv() *GutterPrivate {
	var nilValue *GutterPrivate
	err := gutterObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(gutterObject, recv.Native(), "priv")
	value := GutterPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Gutter) SetFieldPriv(value *GutterPrivate) {
	err := gutterObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(gutterObject, recv.Native(), "priv", argValue)
}

var gutterGetPaddingFunction *gi.Function
var gutterGetPaddingFunction_Once sync.Once

func gutterGetPaddingFunction_Set() error {
	var err error
	gutterGetPaddingFunction_Once.Do(func() {
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterGetPaddingFunction, err = gutterObject.InvokerNew("get_padding")
	})
	return err
}

// GetPadding is a representation of the C type gtk_source_gutter_get_padding.
func (recv *Gutter) GetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterGetRendererAtPosFunction, err = gutterObject.InvokerNew("get_renderer_at_pos")
	})
	return err
}

// GetRendererAtPos is a representation of the C type gtk_source_gutter_get_renderer_at_pos.
func (recv *Gutter) GetRendererAtPos(x int32, y int32) *GutterRenderer {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var ret gi.Argument

	err := gutterGetRendererAtPosFunction_Set()
	if err == nil {
		ret = gutterGetRendererAtPosFunction.Invoke(inArgs[:], nil)
	}

	retGo := GutterRendererNewFromNative(ret.Pointer())

	return retGo
}

var gutterGetViewFunction *gi.Function
var gutterGetViewFunction_Once sync.Once

func gutterGetViewFunction_Set() error {
	var err error
	gutterGetViewFunction_Once.Do(func() {
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterGetViewFunction, err = gutterObject.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_gutter_get_view.
func (recv *Gutter) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterGetViewFunction_Set()
	if err == nil {
		ret = gutterGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ViewNewFromNative(ret.Pointer())

	return retGo
}

var gutterGetWindowFunction *gi.Function
var gutterGetWindowFunction_Once sync.Once

func gutterGetWindowFunction_Set() error {
	var err error
	gutterGetWindowFunction_Once.Do(func() {
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterGetWindowFunction, err = gutterObject.InvokerNew("get_window")
	})
	return err
}

// GetWindow is a representation of the C type gtk_source_gutter_get_window.
func (recv *Gutter) GetWindow() *gdk.Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterGetWindowFunction_Set()
	if err == nil {
		ret = gutterGetWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdk.WindowNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_get_window_type' : return type 'Gtk.TextWindowType' not supported

var gutterInsertFunction *gi.Function
var gutterInsertFunction_Once sync.Once

func gutterInsertFunction_Set() error {
	var err error
	gutterInsertFunction_Once.Do(func() {
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterInsertFunction, err = gutterObject.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type gtk_source_gutter_insert.
func (recv *Gutter) Insert(renderer *GutterRenderer, position int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(renderer.Native())
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
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterQueueDrawFunction, err = gutterObject.InvokerNew("queue_draw")
	})
	return err
}

// QueueDraw is a representation of the C type gtk_source_gutter_queue_draw.
func (recv *Gutter) QueueDraw() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterRemoveFunction, err = gutterObject.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type gtk_source_gutter_remove.
func (recv *Gutter) Remove(renderer *GutterRenderer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(renderer.Native())

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
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterReorderFunction, err = gutterObject.InvokerNew("reorder")
	})
	return err
}

// Reorder is a representation of the C type gtk_source_gutter_reorder.
func (recv *Gutter) Reorder(renderer *GutterRenderer, position int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(renderer.Native())
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
		err = gutterObject_Set()
		if err != nil {
			return
		}
		gutterSetPaddingFunction, err = gutterObject.InvokerNew("set_padding")
	})
	return err
}

// SetPadding is a representation of the C type gtk_source_gutter_set_padding.
func (recv *Gutter) SetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(xpad)
	inArgs[2].SetInt32(ypad)

	err := gutterSetPaddingFunction_Set()
	if err == nil {
		gutterSetPaddingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererObject *gi.Object
var gutterRendererObject_Once sync.Once

func gutterRendererObject_Set() error {
	var err error
	gutterRendererObject_Once.Do(func() {
		gutterRendererObject, err = gi.ObjectNew("GtkSource", "GutterRenderer")
	})
	return err
}

type GutterRenderer struct {
	native unsafe.Pointer
}

func GutterRendererNewFromNative(native unsafe.Pointer) *GutterRenderer {
	instance := &GutterRenderer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRenderer) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *GutterRenderer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToGutterRenderer down casts any arbitrary Object to GutterRenderer.
Exercise care, as this is a potentially dangerous function
if the Object is not a GutterRenderer.
*/
func CastToGutterRenderer(object *gobject.Object) *GutterRenderer {
	return GutterRendererNewFromNative(object.Native())
}

// Equals compares this GutterRenderer with another GutterRenderer, and returns true if they represent the same Object.
func (recv *GutterRenderer) Equals(other *GutterRenderer) bool {
	return other.Native() == recv.Native()
}

func (recv *GutterRenderer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *GutterRenderer) FieldParent() *gobject.InitiallyUnowned {
	var nilValue *gobject.InitiallyUnowned
	err := gutterRendererObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(gutterRendererObject, recv.Native(), "parent")
	value := gobject.InitiallyUnownedNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *GutterRenderer) SetFieldParent(value *gobject.InitiallyUnowned) {
	err := gutterRendererObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(gutterRendererObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_activate' : parameter 'event' of type 'Gdk.Event' not supported

var gutterRendererBeginFunction *gi.Function
var gutterRendererBeginFunction_Once sync.Once

func gutterRendererBeginFunction_Set() error {
	var err error
	gutterRendererBeginFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererBeginFunction, err = gutterRendererObject.InvokerNew("begin")
	})
	return err
}

// Begin is a representation of the C type gtk_source_gutter_renderer_begin.
func (recv *GutterRenderer) Begin(cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cr.Native())
	inArgs[2].SetPointer(backgroundArea.Native())
	inArgs[3].SetPointer(cellArea.Native())
	inArgs[4].SetPointer(start.Native())
	inArgs[5].SetPointer(end.Native())

	err := gutterRendererBeginFunction_Set()
	if err == nil {
		gutterRendererBeginFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererDrawFunction *gi.Function
var gutterRendererDrawFunction_Once sync.Once

func gutterRendererDrawFunction_Set() error {
	var err error
	gutterRendererDrawFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererDrawFunction, err = gutterRendererObject.InvokerNew("draw")
	})
	return err
}

// Draw is a representation of the C type gtk_source_gutter_renderer_draw.
func (recv *GutterRenderer) Draw(cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState) {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cr.Native())
	inArgs[2].SetPointer(backgroundArea.Native())
	inArgs[3].SetPointer(cellArea.Native())
	inArgs[4].SetPointer(start.Native())
	inArgs[5].SetPointer(end.Native())
	inArgs[6].SetInt32(int32(state))

	err := gutterRendererDrawFunction_Set()
	if err == nil {
		gutterRendererDrawFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererEndFunction *gi.Function
var gutterRendererEndFunction_Once sync.Once

func gutterRendererEndFunction_Set() error {
	var err error
	gutterRendererEndFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererEndFunction, err = gutterRendererObject.InvokerNew("end")
	})
	return err
}

// End is a representation of the C type gtk_source_gutter_renderer_end.
func (recv *GutterRenderer) End() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetAlignmentFunction, err = gutterRendererObject.InvokerNew("get_alignment")
	})
	return err
}

// GetAlignment is a representation of the C type gtk_source_gutter_renderer_get_alignment.
func (recv *GutterRenderer) GetAlignment() (float32, float32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetAlignmentModeFunction, err = gutterRendererObject.InvokerNew("get_alignment_mode")
	})
	return err
}

// GetAlignmentMode is a representation of the C type gtk_source_gutter_renderer_get_alignment_mode.
func (recv *GutterRenderer) GetAlignmentMode() GutterRendererAlignmentMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererGetAlignmentModeFunction_Set()
	if err == nil {
		ret = gutterRendererGetAlignmentModeFunction.Invoke(inArgs[:], nil)
	}

	retGo := GutterRendererAlignmentMode(ret.Int32())

	return retGo
}

var gutterRendererGetBackgroundFunction *gi.Function
var gutterRendererGetBackgroundFunction_Once sync.Once

func gutterRendererGetBackgroundFunction_Set() error {
	var err error
	gutterRendererGetBackgroundFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetBackgroundFunction, err = gutterRendererObject.InvokerNew("get_background")
	})
	return err
}

// GetBackground is a representation of the C type gtk_source_gutter_renderer_get_background.
func (recv *GutterRenderer) GetBackground() (bool, *gdk.RGBA) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := gutterRendererGetBackgroundFunction_Set()
	if err == nil {
		ret = gutterRendererGetBackgroundFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gdk.RGBANewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var gutterRendererGetPaddingFunction *gi.Function
var gutterRendererGetPaddingFunction_Once sync.Once

func gutterRendererGetPaddingFunction_Set() error {
	var err error
	gutterRendererGetPaddingFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetPaddingFunction, err = gutterRendererObject.InvokerNew("get_padding")
	})
	return err
}

// GetPadding is a representation of the C type gtk_source_gutter_renderer_get_padding.
func (recv *GutterRenderer) GetPadding() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetSizeFunction, err = gutterRendererObject.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type gtk_source_gutter_renderer_get_size.
func (recv *GutterRenderer) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererGetSizeFunction_Set()
	if err == nil {
		ret = gutterRendererGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var gutterRendererGetViewFunction *gi.Function
var gutterRendererGetViewFunction_Once sync.Once

func gutterRendererGetViewFunction_Set() error {
	var err error
	gutterRendererGetViewFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetViewFunction, err = gutterRendererObject.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_gutter_renderer_get_view.
func (recv *GutterRenderer) GetView() *gtk.TextView {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererGetViewFunction_Set()
	if err == nil {
		ret = gutterRendererGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := gtk.TextViewNewFromNative(ret.Pointer())

	return retGo
}

var gutterRendererGetVisibleFunction *gi.Function
var gutterRendererGetVisibleFunction_Once sync.Once

func gutterRendererGetVisibleFunction_Set() error {
	var err error
	gutterRendererGetVisibleFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererGetVisibleFunction, err = gutterRendererObject.InvokerNew("get_visible")
	})
	return err
}

// GetVisible is a representation of the C type gtk_source_gutter_renderer_get_visible.
func (recv *GutterRenderer) GetVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererGetVisibleFunction_Set()
	if err == nil {
		ret = gutterRendererGetVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_get_window_type' : return type 'Gtk.TextWindowType' not supported

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_query_activatable' : parameter 'event' of type 'Gdk.Event' not supported

var gutterRendererQueryDataFunction *gi.Function
var gutterRendererQueryDataFunction_Once sync.Once

func gutterRendererQueryDataFunction_Set() error {
	var err error
	gutterRendererQueryDataFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererQueryDataFunction, err = gutterRendererObject.InvokerNew("query_data")
	})
	return err
}

// QueryData is a representation of the C type gtk_source_gutter_renderer_query_data.
func (recv *GutterRenderer) QueryData(start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())
	inArgs[3].SetInt32(int32(state))

	err := gutterRendererQueryDataFunction_Set()
	if err == nil {
		gutterRendererQueryDataFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererQueryTooltipFunction *gi.Function
var gutterRendererQueryTooltipFunction_Once sync.Once

func gutterRendererQueryTooltipFunction_Set() error {
	var err error
	gutterRendererQueryTooltipFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererQueryTooltipFunction, err = gutterRendererObject.InvokerNew("query_tooltip")
	})
	return err
}

// QueryTooltip is a representation of the C type gtk_source_gutter_renderer_query_tooltip.
func (recv *GutterRenderer) QueryTooltip(iter *gtk.TextIter, area *gdk.Rectangle, x int32, y int32, tooltip *gtk.Tooltip) bool {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetPointer(area.Native())
	inArgs[3].SetInt32(x)
	inArgs[4].SetInt32(y)
	inArgs[5].SetPointer(tooltip.Native())

	var ret gi.Argument

	err := gutterRendererQueryTooltipFunction_Set()
	if err == nil {
		ret = gutterRendererQueryTooltipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gutterRendererQueueDrawFunction *gi.Function
var gutterRendererQueueDrawFunction_Once sync.Once

func gutterRendererQueueDrawFunction_Set() error {
	var err error
	gutterRendererQueueDrawFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererQueueDrawFunction, err = gutterRendererObject.InvokerNew("queue_draw")
	})
	return err
}

// QueueDraw is a representation of the C type gtk_source_gutter_renderer_queue_draw.
func (recv *GutterRenderer) QueueDraw() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetAlignmentFunction, err = gutterRendererObject.InvokerNew("set_alignment")
	})
	return err
}

// SetAlignment is a representation of the C type gtk_source_gutter_renderer_set_alignment.
func (recv *GutterRenderer) SetAlignment(xalign float32, yalign float32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetAlignmentModeFunction, err = gutterRendererObject.InvokerNew("set_alignment_mode")
	})
	return err
}

// SetAlignmentMode is a representation of the C type gtk_source_gutter_renderer_set_alignment_mode.
func (recv *GutterRenderer) SetAlignmentMode(mode GutterRendererAlignmentMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(mode))

	err := gutterRendererSetAlignmentModeFunction_Set()
	if err == nil {
		gutterRendererSetAlignmentModeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetBackgroundFunction *gi.Function
var gutterRendererSetBackgroundFunction_Once sync.Once

func gutterRendererSetBackgroundFunction_Set() error {
	var err error
	gutterRendererSetBackgroundFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetBackgroundFunction, err = gutterRendererObject.InvokerNew("set_background")
	})
	return err
}

// SetBackground is a representation of the C type gtk_source_gutter_renderer_set_background.
func (recv *GutterRenderer) SetBackground(color *gdk.RGBA) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(color.Native())

	err := gutterRendererSetBackgroundFunction_Set()
	if err == nil {
		gutterRendererSetBackgroundFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererSetPaddingFunction *gi.Function
var gutterRendererSetPaddingFunction_Once sync.Once

func gutterRendererSetPaddingFunction_Set() error {
	var err error
	gutterRendererSetPaddingFunction_Once.Do(func() {
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetPaddingFunction, err = gutterRendererObject.InvokerNew("set_padding")
	})
	return err
}

// SetPadding is a representation of the C type gtk_source_gutter_renderer_set_padding.
func (recv *GutterRenderer) SetPadding(xpad int32, ypad int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetSizeFunction, err = gutterRendererObject.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type gtk_source_gutter_renderer_set_size.
func (recv *GutterRenderer) SetSize(size int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererObject_Set()
		if err != nil {
			return
		}
		gutterRendererSetVisibleFunction, err = gutterRendererObject.InvokerNew("set_visible")
	})
	return err
}

// SetVisible is a representation of the C type gtk_source_gutter_renderer_set_visible.
func (recv *GutterRenderer) SetVisible(visible bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(visible)

	err := gutterRendererSetVisibleFunction_Set()
	if err == nil {
		gutterRendererSetVisibleFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'activate' : parameter 'event' of type 'Gdk.Event' not supported

// UNSUPPORTED : C value 'query-activatable' : parameter 'event' of type 'Gdk.Event' not supported

/*
ConnectQueryData connects a callback to the 'query-data' signal of the GutterRenderer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *GutterRenderer) ConnectQueryData(handler func(instance *GutterRenderer, start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := GutterRendererNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextIterNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := gtk.TextIterNewFromNative(object2.GetObject().Native())

		object3 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[3]))
		arg3 := (GutterRendererState)(object3.GetInt())

		handler(argInstance, arg1, arg2, arg3)
	}

	return callback.ConnectSignal(recv.Native(), "query-data", marshal)
}

/*
ConnectQueryTooltip connects a callback to the 'query-tooltip' signal of the GutterRenderer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *GutterRenderer) ConnectQueryTooltip(handler func(instance *GutterRenderer, iter *gtk.TextIter, area *gdk.Rectangle, x int32, y int32, tooltip *gtk.Tooltip) bool) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := GutterRendererNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextIterNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := gdk.RectangleNewFromNative(object2.GetObject().Native())

		object3 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[3]))
		arg3 := object3.GetInt()

		object4 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[4]))
		arg4 := object4.GetInt()

		object5 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[5]))
		arg5 := gtk.TooltipNewFromNative(object5.GetObject().Native())

		handler(argInstance, arg1, arg2, arg3, arg4, arg5)
	}

	return callback.ConnectSignal(recv.Native(), "query-tooltip", marshal)
}

/*
ConnectQueueDraw connects a callback to the 'queue-draw' signal of the GutterRenderer.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *GutterRenderer) ConnectQueueDraw(handler func(instance *GutterRenderer)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := GutterRendererNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "queue-draw", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *GutterRenderer) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var gutterRendererPixbufObject *gi.Object
var gutterRendererPixbufObject_Once sync.Once

func gutterRendererPixbufObject_Set() error {
	var err error
	gutterRendererPixbufObject_Once.Do(func() {
		gutterRendererPixbufObject, err = gi.ObjectNew("GtkSource", "GutterRendererPixbuf")
	})
	return err
}

type GutterRendererPixbuf struct {
	native unsafe.Pointer
}

func GutterRendererPixbufNewFromNative(native unsafe.Pointer) *GutterRendererPixbuf {
	instance := &GutterRendererPixbuf{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// GutterRenderer upcasts to *GutterRenderer
func (recv *GutterRendererPixbuf) GutterRenderer() *GutterRenderer {
	return GutterRendererNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRendererPixbuf) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *GutterRendererPixbuf) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToGutterRendererPixbuf down casts any arbitrary Object to GutterRendererPixbuf.
Exercise care, as this is a potentially dangerous function
if the Object is not a GutterRendererPixbuf.
*/
func CastToGutterRendererPixbuf(object *gobject.Object) *GutterRendererPixbuf {
	return GutterRendererPixbufNewFromNative(object.Native())
}

// Equals compares this GutterRendererPixbuf with another GutterRendererPixbuf, and returns true if they represent the same Object.
func (recv *GutterRendererPixbuf) Equals(other *GutterRendererPixbuf) bool {
	return other.Native() == recv.Native()
}

func (recv *GutterRendererPixbuf) Native() unsafe.Pointer {
	return recv.native
}

var gutterRendererPixbufNewFunction *gi.Function
var gutterRendererPixbufNewFunction_Once sync.Once

func gutterRendererPixbufNewFunction_Set() error {
	var err error
	gutterRendererPixbufNewFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufNewFunction, err = gutterRendererPixbufObject.InvokerNew("new")
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

	retGo := GutterRendererPixbufNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_gutter_renderer_pixbuf_get_gicon' : return type 'Gio.Icon' not supported

var gutterRendererPixbufGetIconNameFunction *gi.Function
var gutterRendererPixbufGetIconNameFunction_Once sync.Once

func gutterRendererPixbufGetIconNameFunction_Set() error {
	var err error
	gutterRendererPixbufGetIconNameFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufGetIconNameFunction, err = gutterRendererPixbufObject.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_gutter_renderer_pixbuf_get_icon_name.
func (recv *GutterRendererPixbuf) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererPixbufGetIconNameFunction_Set()
	if err == nil {
		ret = gutterRendererPixbufGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var gutterRendererPixbufGetPixbufFunction *gi.Function
var gutterRendererPixbufGetPixbufFunction_Once sync.Once

func gutterRendererPixbufGetPixbufFunction_Set() error {
	var err error
	gutterRendererPixbufGetPixbufFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufGetPixbufFunction, err = gutterRendererPixbufObject.InvokerNew("get_pixbuf")
	})
	return err
}

// GetPixbuf is a representation of the C type gtk_source_gutter_renderer_pixbuf_get_pixbuf.
func (recv *GutterRendererPixbuf) GetPixbuf() *gdkpixbuf.Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gutterRendererPixbufGetPixbufFunction_Set()
	if err == nil {
		ret = gutterRendererPixbufGetPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var gutterRendererPixbufGetStockIdFunction *gi.Function
var gutterRendererPixbufGetStockIdFunction_Once sync.Once

func gutterRendererPixbufGetStockIdFunction_Set() error {
	var err error
	gutterRendererPixbufGetStockIdFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufGetStockIdFunction, err = gutterRendererPixbufObject.InvokerNew("get_stock_id")
	})
	return err
}

// GetStockId is a representation of the C type gtk_source_gutter_renderer_pixbuf_get_stock_id.
func (recv *GutterRendererPixbuf) GetStockId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufSetIconNameFunction, err = gutterRendererPixbufObject.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_gutter_renderer_pixbuf_set_icon_name.
func (recv *GutterRendererPixbuf) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(iconName)

	err := gutterRendererPixbufSetIconNameFunction_Set()
	if err == nil {
		gutterRendererPixbufSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererPixbufSetPixbufFunction *gi.Function
var gutterRendererPixbufSetPixbufFunction_Once sync.Once

func gutterRendererPixbufSetPixbufFunction_Set() error {
	var err error
	gutterRendererPixbufSetPixbufFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufSetPixbufFunction, err = gutterRendererPixbufObject.InvokerNew("set_pixbuf")
	})
	return err
}

// SetPixbuf is a representation of the C type gtk_source_gutter_renderer_pixbuf_set_pixbuf.
func (recv *GutterRendererPixbuf) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(pixbuf.Native())

	err := gutterRendererPixbufSetPixbufFunction_Set()
	if err == nil {
		gutterRendererPixbufSetPixbufFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererPixbufSetStockIdFunction *gi.Function
var gutterRendererPixbufSetStockIdFunction_Once sync.Once

func gutterRendererPixbufSetStockIdFunction_Set() error {
	var err error
	gutterRendererPixbufSetStockIdFunction_Once.Do(func() {
		err = gutterRendererPixbufObject_Set()
		if err != nil {
			return
		}
		gutterRendererPixbufSetStockIdFunction, err = gutterRendererPixbufObject.InvokerNew("set_stock_id")
	})
	return err
}

// SetStockId is a representation of the C type gtk_source_gutter_renderer_pixbuf_set_stock_id.
func (recv *GutterRendererPixbuf) SetStockId(stockId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(stockId)

	err := gutterRendererPixbufSetStockIdFunction_Set()
	if err == nil {
		gutterRendererPixbufSetStockIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gutterRendererTextObject *gi.Object
var gutterRendererTextObject_Once sync.Once

func gutterRendererTextObject_Set() error {
	var err error
	gutterRendererTextObject_Once.Do(func() {
		gutterRendererTextObject, err = gi.ObjectNew("GtkSource", "GutterRendererText")
	})
	return err
}

type GutterRendererText struct {
	native unsafe.Pointer
}

func GutterRendererTextNewFromNative(native unsafe.Pointer) *GutterRendererText {
	instance := &GutterRendererText{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// GutterRenderer upcasts to *GutterRenderer
func (recv *GutterRendererText) GutterRenderer() *GutterRenderer {
	return GutterRendererNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRendererText) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *GutterRendererText) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToGutterRendererText down casts any arbitrary Object to GutterRendererText.
Exercise care, as this is a potentially dangerous function
if the Object is not a GutterRendererText.
*/
func CastToGutterRendererText(object *gobject.Object) *GutterRendererText {
	return GutterRendererTextNewFromNative(object.Native())
}

// Equals compares this GutterRendererText with another GutterRendererText, and returns true if they represent the same Object.
func (recv *GutterRendererText) Equals(other *GutterRendererText) bool {
	return other.Native() == recv.Native()
}

func (recv *GutterRendererText) Native() unsafe.Pointer {
	return recv.native
}

var gutterRendererTextNewFunction *gi.Function
var gutterRendererTextNewFunction_Once sync.Once

func gutterRendererTextNewFunction_Set() error {
	var err error
	gutterRendererTextNewFunction_Once.Do(func() {
		err = gutterRendererTextObject_Set()
		if err != nil {
			return
		}
		gutterRendererTextNewFunction, err = gutterRendererTextObject.InvokerNew("new")
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

	retGo := GutterRendererTextNewFromNative(ret.Pointer())

	return retGo
}

var gutterRendererTextMeasureFunction *gi.Function
var gutterRendererTextMeasureFunction_Once sync.Once

func gutterRendererTextMeasureFunction_Set() error {
	var err error
	gutterRendererTextMeasureFunction_Once.Do(func() {
		err = gutterRendererTextObject_Set()
		if err != nil {
			return
		}
		gutterRendererTextMeasureFunction, err = gutterRendererTextObject.InvokerNew("measure")
	})
	return err
}

// Measure is a representation of the C type gtk_source_gutter_renderer_text_measure.
func (recv *GutterRendererText) Measure(text string) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererTextObject_Set()
		if err != nil {
			return
		}
		gutterRendererTextMeasureMarkupFunction, err = gutterRendererTextObject.InvokerNew("measure_markup")
	})
	return err
}

// MeasureMarkup is a representation of the C type gtk_source_gutter_renderer_text_measure_markup.
func (recv *GutterRendererText) MeasureMarkup(markup string) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererTextObject_Set()
		if err != nil {
			return
		}
		gutterRendererTextSetMarkupFunction, err = gutterRendererTextObject.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type gtk_source_gutter_renderer_text_set_markup.
func (recv *GutterRendererText) SetMarkup(markup string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = gutterRendererTextObject_Set()
		if err != nil {
			return
		}
		gutterRendererTextSetTextFunction, err = gutterRendererTextObject.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type gtk_source_gutter_renderer_text_set_text.
func (recv *GutterRendererText) SetText(text string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)

	err := gutterRendererTextSetTextFunction_Set()
	if err == nil {
		gutterRendererTextSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var languageObject *gi.Object
var languageObject_Once sync.Once

func languageObject_Set() error {
	var err error
	languageObject_Once.Do(func() {
		languageObject, err = gi.ObjectNew("GtkSource", "Language")
	})
	return err
}

type Language struct {
	native unsafe.Pointer
}

func LanguageNewFromNative(native unsafe.Pointer) *Language {
	instance := &Language{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Language) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToLanguage down casts any arbitrary Object to Language.
Exercise care, as this is a potentially dangerous function
if the Object is not a Language.
*/
func CastToLanguage(object *gobject.Object) *Language {
	return LanguageNewFromNative(object.Native())
}

// Equals compares this Language with another Language, and returns true if they represent the same Object.
func (recv *Language) Equals(other *Language) bool {
	return other.Native() == recv.Native()
}

func (recv *Language) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Language) FieldParentInstance() *gobject.Object {
	var nilValue *gobject.Object
	err := languageObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(languageObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Language) SetFieldParentInstance(value *gobject.Object) {
	err := languageObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(languageObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Language) FieldPriv() *LanguagePrivate {
	var nilValue *LanguagePrivate
	err := languageObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(languageObject, recv.Native(), "priv")
	value := LanguagePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Language) SetFieldPriv(value *LanguagePrivate) {
	err := languageObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(languageObject, recv.Native(), "priv", argValue)
}

var languageGetGlobsFunction *gi.Function
var languageGetGlobsFunction_Once sync.Once

func languageGetGlobsFunction_Set() error {
	var err error
	languageGetGlobsFunction_Once.Do(func() {
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetGlobsFunction, err = languageObject.InvokerNew("get_globs")
	})
	return err
}

// GetGlobs is a representation of the C type gtk_source_language_get_globs.
func (recv *Language) GetGlobs() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetHiddenFunction, err = languageObject.InvokerNew("get_hidden")
	})
	return err
}

// GetHidden is a representation of the C type gtk_source_language_get_hidden.
func (recv *Language) GetHidden() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetIdFunction, err = languageObject.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type gtk_source_language_get_id.
func (recv *Language) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetMetadataFunction, err = languageObject.InvokerNew("get_metadata")
	})
	return err
}

// GetMetadata is a representation of the C type gtk_source_language_get_metadata.
func (recv *Language) GetMetadata(name string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetMimeTypesFunction, err = languageObject.InvokerNew("get_mime_types")
	})
	return err
}

// GetMimeTypes is a representation of the C type gtk_source_language_get_mime_types.
func (recv *Language) GetMimeTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetNameFunction, err = languageObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_language_get_name.
func (recv *Language) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetSectionFunction, err = languageObject.InvokerNew("get_section")
	})
	return err
}

// GetSection is a representation of the C type gtk_source_language_get_section.
func (recv *Language) GetSection() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetStyleFallbackFunction, err = languageObject.InvokerNew("get_style_fallback")
	})
	return err
}

// GetStyleFallback is a representation of the C type gtk_source_language_get_style_fallback.
func (recv *Language) GetStyleFallback(styleId string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetStyleIdsFunction, err = languageObject.InvokerNew("get_style_ids")
	})
	return err
}

// GetStyleIds is a representation of the C type gtk_source_language_get_style_ids.
func (recv *Language) GetStyleIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageObject_Set()
		if err != nil {
			return
		}
		languageGetStyleNameFunction, err = languageObject.InvokerNew("get_style_name")
	})
	return err
}

// GetStyleName is a representation of the C type gtk_source_language_get_style_name.
func (recv *Language) GetStyleName(styleId string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(styleId)

	var ret gi.Argument

	err := languageGetStyleNameFunction_Set()
	if err == nil {
		ret = languageGetStyleNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageManagerObject *gi.Object
var languageManagerObject_Once sync.Once

func languageManagerObject_Set() error {
	var err error
	languageManagerObject_Once.Do(func() {
		languageManagerObject, err = gi.ObjectNew("GtkSource", "LanguageManager")
	})
	return err
}

type LanguageManager struct {
	native unsafe.Pointer
}

func LanguageManagerNewFromNative(native unsafe.Pointer) *LanguageManager {
	instance := &LanguageManager{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *LanguageManager) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToLanguageManager down casts any arbitrary Object to LanguageManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a LanguageManager.
*/
func CastToLanguageManager(object *gobject.Object) *LanguageManager {
	return LanguageManagerNewFromNative(object.Native())
}

// Equals compares this LanguageManager with another LanguageManager, and returns true if they represent the same Object.
func (recv *LanguageManager) Equals(other *LanguageManager) bool {
	return other.Native() == recv.Native()
}

func (recv *LanguageManager) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *LanguageManager) FieldParentInstance() *gobject.Object {
	var nilValue *gobject.Object
	err := languageManagerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(languageManagerObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *LanguageManager) SetFieldParentInstance(value *gobject.Object) {
	err := languageManagerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(languageManagerObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *LanguageManager) FieldPriv() *LanguageManagerPrivate {
	var nilValue *LanguageManagerPrivate
	err := languageManagerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(languageManagerObject, recv.Native(), "priv")
	value := LanguageManagerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *LanguageManager) SetFieldPriv(value *LanguageManagerPrivate) {
	err := languageManagerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(languageManagerObject, recv.Native(), "priv", argValue)
}

var languageManagerNewFunction *gi.Function
var languageManagerNewFunction_Once sync.Once

func languageManagerNewFunction_Set() error {
	var err error
	languageManagerNewFunction_Once.Do(func() {
		err = languageManagerObject_Set()
		if err != nil {
			return
		}
		languageManagerNewFunction, err = languageManagerObject.InvokerNew("new")
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

	retGo := LanguageManagerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var languageManagerGetLanguageFunction *gi.Function
var languageManagerGetLanguageFunction_Once sync.Once

func languageManagerGetLanguageFunction_Set() error {
	var err error
	languageManagerGetLanguageFunction_Once.Do(func() {
		err = languageManagerObject_Set()
		if err != nil {
			return
		}
		languageManagerGetLanguageFunction, err = languageManagerObject.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type gtk_source_language_manager_get_language.
func (recv *LanguageManager) GetLanguage(id string) *Language {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)

	var ret gi.Argument

	err := languageManagerGetLanguageFunction_Set()
	if err == nil {
		ret = languageManagerGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := LanguageNewFromNative(ret.Pointer())

	return retGo
}

var languageManagerGetLanguageIdsFunction *gi.Function
var languageManagerGetLanguageIdsFunction_Once sync.Once

func languageManagerGetLanguageIdsFunction_Set() error {
	var err error
	languageManagerGetLanguageIdsFunction_Once.Do(func() {
		err = languageManagerObject_Set()
		if err != nil {
			return
		}
		languageManagerGetLanguageIdsFunction, err = languageManagerObject.InvokerNew("get_language_ids")
	})
	return err
}

// GetLanguageIds is a representation of the C type gtk_source_language_manager_get_language_ids.
func (recv *LanguageManager) GetLanguageIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageManagerObject_Set()
		if err != nil {
			return
		}
		languageManagerGetSearchPathFunction, err = languageManagerObject.InvokerNew("get_search_path")
	})
	return err
}

// GetSearchPath is a representation of the C type gtk_source_language_manager_get_search_path.
func (recv *LanguageManager) GetSearchPath() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = languageManagerObject_Set()
		if err != nil {
			return
		}
		languageManagerGuessLanguageFunction, err = languageManagerObject.InvokerNew("guess_language")
	})
	return err
}

// GuessLanguage is a representation of the C type gtk_source_language_manager_guess_language.
func (recv *LanguageManager) GuessLanguage(filename string, contentType string) *Language {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(filename)
	inArgs[2].SetString(contentType)

	var ret gi.Argument

	err := languageManagerGuessLanguageFunction_Set()
	if err == nil {
		ret = languageManagerGuessLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := LanguageNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_language_manager_set_search_path' : parameter 'dirs' of type 'nil' not supported

var mapObject *gi.Object
var mapObject_Once sync.Once

func mapObject_Set() error {
	var err error
	mapObject_Once.Do(func() {
		mapObject, err = gi.ObjectNew("GtkSource", "Map")
	})
	return err
}

type Map struct {
	native unsafe.Pointer
}

func MapNewFromNative(native unsafe.Pointer) *Map {
	instance := &Map{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// View upcasts to *View
func (recv *Map) View() *View {
	return ViewNewFromNative(recv.Native())
}

// TextView upcasts to *TextView
func (recv *Map) TextView() *gtk.TextView {
	return gtk.TextViewNewFromNative(recv.Native())
}

// Container upcasts to *Container
func (recv *Map) Container() *gtk.Container {
	return gtk.ContainerNewFromNative(recv.Native())
}

// Widget upcasts to *Widget
func (recv *Map) Widget() *gtk.Widget {
	return gtk.WidgetNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Map) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Map) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMap down casts any arbitrary Object to Map.
Exercise care, as this is a potentially dangerous function
if the Object is not a Map.
*/
func CastToMap(object *gobject.Object) *Map {
	return MapNewFromNative(object.Native())
}

// Equals compares this Map with another Map, and returns true if they represent the same Object.
func (recv *Map) Equals(other *Map) bool {
	return other.Native() == recv.Native()
}

func (recv *Map) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Map) FieldParentInstance() *View {
	var nilValue *View
	err := mapObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(mapObject, recv.Native(), "parent_instance")
	value := ViewNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Map) SetFieldParentInstance(value *View) {
	err := mapObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(mapObject, recv.Native(), "parent_instance", argValue)
}

var mapNewFunction *gi.Function
var mapNewFunction_Once sync.Once

func mapNewFunction_Set() error {
	var err error
	mapNewFunction_Once.Do(func() {
		err = mapObject_Set()
		if err != nil {
			return
		}
		mapNewFunction, err = mapObject.InvokerNew("new")
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

	retGo := MapNewFromNative(ret.Pointer())

	return retGo
}

var mapGetViewFunction *gi.Function
var mapGetViewFunction_Once sync.Once

func mapGetViewFunction_Set() error {
	var err error
	mapGetViewFunction_Once.Do(func() {
		err = mapObject_Set()
		if err != nil {
			return
		}
		mapGetViewFunction, err = mapObject.InvokerNew("get_view")
	})
	return err
}

// GetView is a representation of the C type gtk_source_map_get_view.
func (recv *Map) GetView() *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mapGetViewFunction_Set()
	if err == nil {
		ret = mapGetViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ViewNewFromNative(ret.Pointer())

	return retGo
}

var mapSetViewFunction *gi.Function
var mapSetViewFunction_Once sync.Once

func mapSetViewFunction_Set() error {
	var err error
	mapSetViewFunction_Once.Do(func() {
		err = mapObject_Set()
		if err != nil {
			return
		}
		mapSetViewFunction, err = mapObject.InvokerNew("set_view")
	})
	return err
}

// SetView is a representation of the C type gtk_source_map_set_view.
func (recv *Map) SetView(view *View) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(view.Native())

	err := mapSetViewFunction_Set()
	if err == nil {
		mapSetViewFunction.Invoke(inArgs[:], nil)
	}

	return
}

// AtkImplementorIface returns the Atk.ImplementorIface interface implemented by Map
func (recv *Map) AtkImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromNative(recv.Native())
}

// GtkBuildable returns the Gtk.Buildable interface implemented by Map
func (recv *Map) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

// GtkScrollable returns the Gtk.Scrollable interface implemented by Map
func (recv *Map) GtkScrollable() *gtk.Scrollable {
	return gtk.ScrollableNewFromNative(recv.Native())
}

var markObject *gi.Object
var markObject_Once sync.Once

func markObject_Set() error {
	var err error
	markObject_Once.Do(func() {
		markObject, err = gi.ObjectNew("GtkSource", "Mark")
	})
	return err
}

type Mark struct {
	native unsafe.Pointer
}

func MarkNewFromNative(native unsafe.Pointer) *Mark {
	instance := &Mark{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TextMark upcasts to *TextMark
func (recv *Mark) TextMark() *gtk.TextMark {
	return gtk.TextMarkNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Mark) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMark down casts any arbitrary Object to Mark.
Exercise care, as this is a potentially dangerous function
if the Object is not a Mark.
*/
func CastToMark(object *gobject.Object) *Mark {
	return MarkNewFromNative(object.Native())
}

// Equals compares this Mark with another Mark, and returns true if they represent the same Object.
func (recv *Mark) Equals(other *Mark) bool {
	return other.Native() == recv.Native()
}

func (recv *Mark) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Mark) FieldParentInstance() *gtk.TextMark {
	var nilValue *gtk.TextMark
	err := markObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(markObject, recv.Native(), "parent_instance")
	value := gtk.TextMarkNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Mark) SetFieldParentInstance(value *gtk.TextMark) {
	err := markObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(markObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *Mark) FieldPriv() *MarkPrivate {
	var nilValue *MarkPrivate
	err := markObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(markObject, recv.Native(), "priv")
	value := MarkPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *Mark) SetFieldPriv(value *MarkPrivate) {
	err := markObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(markObject, recv.Native(), "priv", argValue)
}

var markNewFunction *gi.Function
var markNewFunction_Once sync.Once

func markNewFunction_Set() error {
	var err error
	markNewFunction_Once.Do(func() {
		err = markObject_Set()
		if err != nil {
			return
		}
		markNewFunction, err = markObject.InvokerNew("new")
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

	retGo := MarkNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var markGetCategoryFunction *gi.Function
var markGetCategoryFunction_Once sync.Once

func markGetCategoryFunction_Set() error {
	var err error
	markGetCategoryFunction_Once.Do(func() {
		err = markObject_Set()
		if err != nil {
			return
		}
		markGetCategoryFunction, err = markObject.InvokerNew("get_category")
	})
	return err
}

// GetCategory is a representation of the C type gtk_source_mark_get_category.
func (recv *Mark) GetCategory() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = markObject_Set()
		if err != nil {
			return
		}
		markNextFunction, err = markObject.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type gtk_source_mark_next.
func (recv *Mark) Next(category string) *Mark {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(category)

	var ret gi.Argument

	err := markNextFunction_Set()
	if err == nil {
		ret = markNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := MarkNewFromNative(ret.Pointer())

	return retGo
}

var markPrevFunction *gi.Function
var markPrevFunction_Once sync.Once

func markPrevFunction_Set() error {
	var err error
	markPrevFunction_Once.Do(func() {
		err = markObject_Set()
		if err != nil {
			return
		}
		markPrevFunction, err = markObject.InvokerNew("prev")
	})
	return err
}

// Prev is a representation of the C type gtk_source_mark_prev.
func (recv *Mark) Prev(category string) *Mark {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(category)

	var ret gi.Argument

	err := markPrevFunction_Set()
	if err == nil {
		ret = markPrevFunction.Invoke(inArgs[:], nil)
	}

	retGo := MarkNewFromNative(ret.Pointer())

	return retGo
}

var markAttributesObject *gi.Object
var markAttributesObject_Once sync.Once

func markAttributesObject_Set() error {
	var err error
	markAttributesObject_Once.Do(func() {
		markAttributesObject, err = gi.ObjectNew("GtkSource", "MarkAttributes")
	})
	return err
}

type MarkAttributes struct {
	native unsafe.Pointer
}

func MarkAttributesNewFromNative(native unsafe.Pointer) *MarkAttributes {
	instance := &MarkAttributes{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *MarkAttributes) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMarkAttributes down casts any arbitrary Object to MarkAttributes.
Exercise care, as this is a potentially dangerous function
if the Object is not a MarkAttributes.
*/
func CastToMarkAttributes(object *gobject.Object) *MarkAttributes {
	return MarkAttributesNewFromNative(object.Native())
}

// Equals compares this MarkAttributes with another MarkAttributes, and returns true if they represent the same Object.
func (recv *MarkAttributes) Equals(other *MarkAttributes) bool {
	return other.Native() == recv.Native()
}

func (recv *MarkAttributes) Native() unsafe.Pointer {
	return recv.native
}

var markAttributesNewFunction *gi.Function
var markAttributesNewFunction_Once sync.Once

func markAttributesNewFunction_Set() error {
	var err error
	markAttributesNewFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesNewFunction, err = markAttributesObject.InvokerNew("new")
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

	retGo := MarkAttributesNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var markAttributesGetBackgroundFunction *gi.Function
var markAttributesGetBackgroundFunction_Once sync.Once

func markAttributesGetBackgroundFunction_Set() error {
	var err error
	markAttributesGetBackgroundFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetBackgroundFunction, err = markAttributesObject.InvokerNew("get_background")
	})
	return err
}

// GetBackground is a representation of the C type gtk_source_mark_attributes_get_background.
func (recv *MarkAttributes) GetBackground() (bool, *gdk.RGBA) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := markAttributesGetBackgroundFunction_Set()
	if err == nil {
		ret = markAttributesGetBackgroundFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gdk.RGBANewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_get_gicon' : return type 'Gio.Icon' not supported

var markAttributesGetIconNameFunction *gi.Function
var markAttributesGetIconNameFunction_Once sync.Once

func markAttributesGetIconNameFunction_Set() error {
	var err error
	markAttributesGetIconNameFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetIconNameFunction, err = markAttributesObject.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_mark_attributes_get_icon_name.
func (recv *MarkAttributes) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := markAttributesGetIconNameFunction_Set()
	if err == nil {
		ret = markAttributesGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var markAttributesGetPixbufFunction *gi.Function
var markAttributesGetPixbufFunction_Once sync.Once

func markAttributesGetPixbufFunction_Set() error {
	var err error
	markAttributesGetPixbufFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetPixbufFunction, err = markAttributesObject.InvokerNew("get_pixbuf")
	})
	return err
}

// GetPixbuf is a representation of the C type gtk_source_mark_attributes_get_pixbuf.
func (recv *MarkAttributes) GetPixbuf() *gdkpixbuf.Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := markAttributesGetPixbufFunction_Set()
	if err == nil {
		ret = markAttributesGetPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var markAttributesGetStockIdFunction *gi.Function
var markAttributesGetStockIdFunction_Once sync.Once

func markAttributesGetStockIdFunction_Set() error {
	var err error
	markAttributesGetStockIdFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetStockIdFunction, err = markAttributesObject.InvokerNew("get_stock_id")
	})
	return err
}

// GetStockId is a representation of the C type gtk_source_mark_attributes_get_stock_id.
func (recv *MarkAttributes) GetStockId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetTooltipMarkupFunction, err = markAttributesObject.InvokerNew("get_tooltip_markup")
	})
	return err
}

// GetTooltipMarkup is a representation of the C type gtk_source_mark_attributes_get_tooltip_markup.
func (recv *MarkAttributes) GetTooltipMarkup(mark *Mark) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(mark.Native())

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
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesGetTooltipTextFunction, err = markAttributesObject.InvokerNew("get_tooltip_text")
	})
	return err
}

// GetTooltipText is a representation of the C type gtk_source_mark_attributes_get_tooltip_text.
func (recv *MarkAttributes) GetTooltipText(mark *Mark) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(mark.Native())

	var ret gi.Argument

	err := markAttributesGetTooltipTextFunction_Set()
	if err == nil {
		ret = markAttributesGetTooltipTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var markAttributesRenderIconFunction *gi.Function
var markAttributesRenderIconFunction_Once sync.Once

func markAttributesRenderIconFunction_Set() error {
	var err error
	markAttributesRenderIconFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesRenderIconFunction, err = markAttributesObject.InvokerNew("render_icon")
	})
	return err
}

// RenderIcon is a representation of the C type gtk_source_mark_attributes_render_icon.
func (recv *MarkAttributes) RenderIcon(widget *gtk.Widget, size int32) *gdkpixbuf.Pixbuf {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(widget.Native())
	inArgs[2].SetInt32(size)

	var ret gi.Argument

	err := markAttributesRenderIconFunction_Set()
	if err == nil {
		ret = markAttributesRenderIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var markAttributesSetBackgroundFunction *gi.Function
var markAttributesSetBackgroundFunction_Once sync.Once

func markAttributesSetBackgroundFunction_Set() error {
	var err error
	markAttributesSetBackgroundFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesSetBackgroundFunction, err = markAttributesObject.InvokerNew("set_background")
	})
	return err
}

// SetBackground is a representation of the C type gtk_source_mark_attributes_set_background.
func (recv *MarkAttributes) SetBackground(background *gdk.RGBA) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(background.Native())

	err := markAttributesSetBackgroundFunction_Set()
	if err == nil {
		markAttributesSetBackgroundFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_mark_attributes_set_gicon' : parameter 'gicon' of type 'Gio.Icon' not supported

var markAttributesSetIconNameFunction *gi.Function
var markAttributesSetIconNameFunction_Once sync.Once

func markAttributesSetIconNameFunction_Set() error {
	var err error
	markAttributesSetIconNameFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesSetIconNameFunction, err = markAttributesObject.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_source_mark_attributes_set_icon_name.
func (recv *MarkAttributes) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(iconName)

	err := markAttributesSetIconNameFunction_Set()
	if err == nil {
		markAttributesSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var markAttributesSetPixbufFunction *gi.Function
var markAttributesSetPixbufFunction_Once sync.Once

func markAttributesSetPixbufFunction_Set() error {
	var err error
	markAttributesSetPixbufFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesSetPixbufFunction, err = markAttributesObject.InvokerNew("set_pixbuf")
	})
	return err
}

// SetPixbuf is a representation of the C type gtk_source_mark_attributes_set_pixbuf.
func (recv *MarkAttributes) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(pixbuf.Native())

	err := markAttributesSetPixbufFunction_Set()
	if err == nil {
		markAttributesSetPixbufFunction.Invoke(inArgs[:], nil)
	}

	return
}

var markAttributesSetStockIdFunction *gi.Function
var markAttributesSetStockIdFunction_Once sync.Once

func markAttributesSetStockIdFunction_Set() error {
	var err error
	markAttributesSetStockIdFunction_Once.Do(func() {
		err = markAttributesObject_Set()
		if err != nil {
			return
		}
		markAttributesSetStockIdFunction, err = markAttributesObject.InvokerNew("set_stock_id")
	})
	return err
}

// SetStockId is a representation of the C type gtk_source_mark_attributes_set_stock_id.
func (recv *MarkAttributes) SetStockId(stockId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(stockId)

	err := markAttributesSetStockIdFunction_Set()
	if err == nil {
		markAttributesSetStockIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectQueryTooltipMarkup connects a callback to the 'query-tooltip-markup' signal of the MarkAttributes.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *MarkAttributes) ConnectQueryTooltipMarkup(handler func(instance *MarkAttributes, mark *Mark) string) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MarkAttributesNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MarkNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "query-tooltip-markup", marshal)
}

/*
ConnectQueryTooltipText connects a callback to the 'query-tooltip-text' signal of the MarkAttributes.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *MarkAttributes) ConnectQueryTooltipText(handler func(instance *MarkAttributes, mark *Mark) string) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := MarkAttributesNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := MarkNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "query-tooltip-text", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *MarkAttributes) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var printCompositorObject *gi.Object
var printCompositorObject_Once sync.Once

func printCompositorObject_Set() error {
	var err error
	printCompositorObject_Once.Do(func() {
		printCompositorObject, err = gi.ObjectNew("GtkSource", "PrintCompositor")
	})
	return err
}

type PrintCompositor struct {
	native unsafe.Pointer
}

func PrintCompositorNewFromNative(native unsafe.Pointer) *PrintCompositor {
	instance := &PrintCompositor{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *PrintCompositor) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToPrintCompositor down casts any arbitrary Object to PrintCompositor.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintCompositor.
*/
func CastToPrintCompositor(object *gobject.Object) *PrintCompositor {
	return PrintCompositorNewFromNative(object.Native())
}

// Equals compares this PrintCompositor with another PrintCompositor, and returns true if they represent the same Object.
func (recv *PrintCompositor) Equals(other *PrintCompositor) bool {
	return other.Native() == recv.Native()
}

func (recv *PrintCompositor) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *PrintCompositor) FieldParentInstance() *gobject.Object {
	var nilValue *gobject.Object
	err := printCompositorObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(printCompositorObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *PrintCompositor) SetFieldParentInstance(value *gobject.Object) {
	err := printCompositorObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(printCompositorObject, recv.Native(), "parent_instance", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *PrintCompositor) FieldPriv() *PrintCompositorPrivate {
	var nilValue *PrintCompositorPrivate
	err := printCompositorObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(printCompositorObject, recv.Native(), "priv")
	value := PrintCompositorPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *PrintCompositor) SetFieldPriv(value *PrintCompositorPrivate) {
	err := printCompositorObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(printCompositorObject, recv.Native(), "priv", argValue)
}

var printCompositorNewFunction *gi.Function
var printCompositorNewFunction_Once sync.Once

func printCompositorNewFunction_Set() error {
	var err error
	printCompositorNewFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorNewFunction, err = printCompositorObject.InvokerNew("new")
	})
	return err
}

// PrintCompositorNew is a representation of the C type gtk_source_print_compositor_new.
func PrintCompositorNew(buffer *Buffer) *PrintCompositor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(buffer.Native())

	var ret gi.Argument

	err := printCompositorNewFunction_Set()
	if err == nil {
		ret = printCompositorNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := PrintCompositorNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var printCompositorNewFromViewFunction *gi.Function
var printCompositorNewFromViewFunction_Once sync.Once

func printCompositorNewFromViewFunction_Set() error {
	var err error
	printCompositorNewFromViewFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorNewFromViewFunction, err = printCompositorObject.InvokerNew("new_from_view")
	})
	return err
}

// PrintCompositorNewFromView is a representation of the C type gtk_source_print_compositor_new_from_view.
func PrintCompositorNewFromView(view *View) *PrintCompositor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(view.Native())

	var ret gi.Argument

	err := printCompositorNewFromViewFunction_Set()
	if err == nil {
		ret = printCompositorNewFromViewFunction.Invoke(inArgs[:], nil)
	}

	retGo := PrintCompositorNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var printCompositorDrawPageFunction *gi.Function
var printCompositorDrawPageFunction_Once sync.Once

func printCompositorDrawPageFunction_Set() error {
	var err error
	printCompositorDrawPageFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorDrawPageFunction, err = printCompositorObject.InvokerNew("draw_page")
	})
	return err
}

// DrawPage is a representation of the C type gtk_source_print_compositor_draw_page.
func (recv *PrintCompositor) DrawPage(context *gtk.PrintContext, pageNr int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())
	inArgs[2].SetInt32(pageNr)

	err := printCompositorDrawPageFunction_Set()
	if err == nil {
		printCompositorDrawPageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printCompositorGetBodyFontNameFunction *gi.Function
var printCompositorGetBodyFontNameFunction_Once sync.Once

func printCompositorGetBodyFontNameFunction_Set() error {
	var err error
	printCompositorGetBodyFontNameFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetBodyFontNameFunction, err = printCompositorObject.InvokerNew("get_body_font_name")
	})
	return err
}

// GetBodyFontName is a representation of the C type gtk_source_print_compositor_get_body_font_name.
func (recv *PrintCompositor) GetBodyFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetBufferFunction, err = printCompositorObject.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_print_compositor_get_buffer.
func (recv *PrintCompositor) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := printCompositorGetBufferFunction_Set()
	if err == nil {
		ret = printCompositorGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

	return retGo
}

var printCompositorGetFooterFontNameFunction *gi.Function
var printCompositorGetFooterFontNameFunction_Once sync.Once

func printCompositorGetFooterFontNameFunction_Set() error {
	var err error
	printCompositorGetFooterFontNameFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetFooterFontNameFunction, err = printCompositorObject.InvokerNew("get_footer_font_name")
	})
	return err
}

// GetFooterFontName is a representation of the C type gtk_source_print_compositor_get_footer_font_name.
func (recv *PrintCompositor) GetFooterFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetHeaderFontNameFunction, err = printCompositorObject.InvokerNew("get_header_font_name")
	})
	return err
}

// GetHeaderFontName is a representation of the C type gtk_source_print_compositor_get_header_font_name.
func (recv *PrintCompositor) GetHeaderFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetHighlightSyntaxFunction, err = printCompositorObject.InvokerNew("get_highlight_syntax")
	})
	return err
}

// GetHighlightSyntax is a representation of the C type gtk_source_print_compositor_get_highlight_syntax.
func (recv *PrintCompositor) GetHighlightSyntax() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetLineNumbersFontNameFunction, err = printCompositorObject.InvokerNew("get_line_numbers_font_name")
	})
	return err
}

// GetLineNumbersFontName is a representation of the C type gtk_source_print_compositor_get_line_numbers_font_name.
func (recv *PrintCompositor) GetLineNumbersFontName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetNPagesFunction, err = printCompositorObject.InvokerNew("get_n_pages")
	})
	return err
}

// GetNPages is a representation of the C type gtk_source_print_compositor_get_n_pages.
func (recv *PrintCompositor) GetNPages() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetPaginationProgressFunction, err = printCompositorObject.InvokerNew("get_pagination_progress")
	})
	return err
}

// GetPaginationProgress is a representation of the C type gtk_source_print_compositor_get_pagination_progress.
func (recv *PrintCompositor) GetPaginationProgress() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintFooterFunction, err = printCompositorObject.InvokerNew("get_print_footer")
	})
	return err
}

// GetPrintFooter is a representation of the C type gtk_source_print_compositor_get_print_footer.
func (recv *PrintCompositor) GetPrintFooter() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintHeaderFunction, err = printCompositorObject.InvokerNew("get_print_header")
	})
	return err
}

// GetPrintHeader is a representation of the C type gtk_source_print_compositor_get_print_header.
func (recv *PrintCompositor) GetPrintHeader() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetPrintLineNumbersFunction, err = printCompositorObject.InvokerNew("get_print_line_numbers")
	})
	return err
}

// GetPrintLineNumbers is a representation of the C type gtk_source_print_compositor_get_print_line_numbers.
func (recv *PrintCompositor) GetPrintLineNumbers() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorGetTabWidthFunction, err = printCompositorObject.InvokerNew("get_tab_width")
	})
	return err
}

// GetTabWidth is a representation of the C type gtk_source_print_compositor_get_tab_width.
func (recv *PrintCompositor) GetTabWidth() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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

var printCompositorPaginateFunction *gi.Function
var printCompositorPaginateFunction_Once sync.Once

func printCompositorPaginateFunction_Set() error {
	var err error
	printCompositorPaginateFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorPaginateFunction, err = printCompositorObject.InvokerNew("paginate")
	})
	return err
}

// Paginate is a representation of the C type gtk_source_print_compositor_paginate.
func (recv *PrintCompositor) Paginate(context *gtk.PrintContext) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())

	var ret gi.Argument

	err := printCompositorPaginateFunction_Set()
	if err == nil {
		ret = printCompositorPaginateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var printCompositorSetBodyFontNameFunction *gi.Function
var printCompositorSetBodyFontNameFunction_Once sync.Once

func printCompositorSetBodyFontNameFunction_Set() error {
	var err error
	printCompositorSetBodyFontNameFunction_Once.Do(func() {
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetBodyFontNameFunction, err = printCompositorObject.InvokerNew("set_body_font_name")
	})
	return err
}

// SetBodyFontName is a representation of the C type gtk_source_print_compositor_set_body_font_name.
func (recv *PrintCompositor) SetBodyFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetFooterFontNameFunction, err = printCompositorObject.InvokerNew("set_footer_font_name")
	})
	return err
}

// SetFooterFontName is a representation of the C type gtk_source_print_compositor_set_footer_font_name.
func (recv *PrintCompositor) SetFooterFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetFooterFormatFunction, err = printCompositorObject.InvokerNew("set_footer_format")
	})
	return err
}

// SetFooterFormat is a representation of the C type gtk_source_print_compositor_set_footer_format.
func (recv *PrintCompositor) SetFooterFormat(separator bool, left string, center string, right string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetHeaderFontNameFunction, err = printCompositorObject.InvokerNew("set_header_font_name")
	})
	return err
}

// SetHeaderFontName is a representation of the C type gtk_source_print_compositor_set_header_font_name.
func (recv *PrintCompositor) SetHeaderFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetHeaderFormatFunction, err = printCompositorObject.InvokerNew("set_header_format")
	})
	return err
}

// SetHeaderFormat is a representation of the C type gtk_source_print_compositor_set_header_format.
func (recv *PrintCompositor) SetHeaderFormat(separator bool, left string, center string, right string) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetHighlightSyntaxFunction, err = printCompositorObject.InvokerNew("set_highlight_syntax")
	})
	return err
}

// SetHighlightSyntax is a representation of the C type gtk_source_print_compositor_set_highlight_syntax.
func (recv *PrintCompositor) SetHighlightSyntax(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetLineNumbersFontNameFunction, err = printCompositorObject.InvokerNew("set_line_numbers_font_name")
	})
	return err
}

// SetLineNumbersFontName is a representation of the C type gtk_source_print_compositor_set_line_numbers_font_name.
func (recv *PrintCompositor) SetLineNumbersFontName(fontName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintFooterFunction, err = printCompositorObject.InvokerNew("set_print_footer")
	})
	return err
}

// SetPrintFooter is a representation of the C type gtk_source_print_compositor_set_print_footer.
func (recv *PrintCompositor) SetPrintFooter(print bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintHeaderFunction, err = printCompositorObject.InvokerNew("set_print_header")
	})
	return err
}

// SetPrintHeader is a representation of the C type gtk_source_print_compositor_set_print_header.
func (recv *PrintCompositor) SetPrintHeader(print bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetPrintLineNumbersFunction, err = printCompositorObject.InvokerNew("set_print_line_numbers")
	})
	return err
}

// SetPrintLineNumbers is a representation of the C type gtk_source_print_compositor_set_print_line_numbers.
func (recv *PrintCompositor) SetPrintLineNumbers(interval uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = printCompositorObject_Set()
		if err != nil {
			return
		}
		printCompositorSetTabWidthFunction, err = printCompositorObject.InvokerNew("set_tab_width")
	})
	return err
}

// SetTabWidth is a representation of the C type gtk_source_print_compositor_set_tab_width.
func (recv *PrintCompositor) SetTabWidth(width uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(width)

	err := printCompositorSetTabWidthFunction_Set()
	if err == nil {
		printCompositorSetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_top_margin' : parameter 'unit' of type 'Gtk.Unit' not supported

// UNSUPPORTED : C value 'gtk_source_print_compositor_set_wrap_mode' : parameter 'wrap_mode' of type 'Gtk.WrapMode' not supported

var regionObject *gi.Object
var regionObject_Once sync.Once

func regionObject_Set() error {
	var err error
	regionObject_Once.Do(func() {
		regionObject, err = gi.ObjectNew("GtkSource", "Region")
	})
	return err
}

type Region struct {
	native unsafe.Pointer
}

func RegionNewFromNative(native unsafe.Pointer) *Region {
	instance := &Region{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Region) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRegion down casts any arbitrary Object to Region.
Exercise care, as this is a potentially dangerous function
if the Object is not a Region.
*/
func CastToRegion(object *gobject.Object) *Region {
	return RegionNewFromNative(object.Native())
}

// Equals compares this Region with another Region, and returns true if they represent the same Object.
func (recv *Region) Equals(other *Region) bool {
	return other.Native() == recv.Native()
}

func (recv *Region) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Region) FieldParentInstance() *gobject.Object {
	var nilValue *gobject.Object
	err := regionObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(regionObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Region) SetFieldParentInstance(value *gobject.Object) {
	err := regionObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(regionObject, recv.Native(), "parent_instance", argValue)
}

var regionNewFunction *gi.Function
var regionNewFunction_Once sync.Once

func regionNewFunction_Set() error {
	var err error
	regionNewFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionNewFunction, err = regionObject.InvokerNew("new")
	})
	return err
}

// RegionNew is a representation of the C type gtk_source_region_new.
func RegionNew(buffer *gtk.TextBuffer) *Region {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(buffer.Native())

	var ret gi.Argument

	err := regionNewFunction_Set()
	if err == nil {
		ret = regionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := RegionNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var regionAddRegionFunction *gi.Function
var regionAddRegionFunction_Once sync.Once

func regionAddRegionFunction_Set() error {
	var err error
	regionAddRegionFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionAddRegionFunction, err = regionObject.InvokerNew("add_region")
	})
	return err
}

// AddRegion is a representation of the C type gtk_source_region_add_region.
func (recv *Region) AddRegion(regionToAdd *Region) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(regionToAdd.Native())

	err := regionAddRegionFunction_Set()
	if err == nil {
		regionAddRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var regionAddSubregionFunction *gi.Function
var regionAddSubregionFunction_Once sync.Once

func regionAddSubregionFunction_Set() error {
	var err error
	regionAddSubregionFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionAddSubregionFunction, err = regionObject.InvokerNew("add_subregion")
	})
	return err
}

// AddSubregion is a representation of the C type gtk_source_region_add_subregion.
func (recv *Region) AddSubregion(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := regionAddSubregionFunction_Set()
	if err == nil {
		regionAddSubregionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var regionGetBoundsFunction *gi.Function
var regionGetBoundsFunction_Once sync.Once

func regionGetBoundsFunction_Set() error {
	var err error
	regionGetBoundsFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionGetBoundsFunction, err = regionObject.InvokerNew("get_bounds")
	})
	return err
}

// GetBounds is a representation of the C type gtk_source_region_get_bounds.
func (recv *Region) GetBounds() (bool, *gtk.TextIter, *gtk.TextIter) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := regionGetBoundsFunction_Set()
	if err == nil {
		ret = regionGetBoundsFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())
	out1 := gtk.TextIterNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var regionGetBufferFunction *gi.Function
var regionGetBufferFunction_Once sync.Once

func regionGetBufferFunction_Set() error {
	var err error
	regionGetBufferFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionGetBufferFunction, err = regionObject.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_region_get_buffer.
func (recv *Region) GetBuffer() *gtk.TextBuffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := regionGetBufferFunction_Set()
	if err == nil {
		ret = regionGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := gtk.TextBufferNewFromNative(ret.Pointer())

	return retGo
}

var regionGetStartRegionIterFunction *gi.Function
var regionGetStartRegionIterFunction_Once sync.Once

func regionGetStartRegionIterFunction_Set() error {
	var err error
	regionGetStartRegionIterFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionGetStartRegionIterFunction, err = regionObject.InvokerNew("get_start_region_iter")
	})
	return err
}

// GetStartRegionIter is a representation of the C type gtk_source_region_get_start_region_iter.
func (recv *Region) GetStartRegionIter() *RegionIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := regionGetStartRegionIterFunction_Set()
	if err == nil {
		regionGetStartRegionIterFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RegionIterNewFromNative(outArgs[0].Pointer())

	return out0
}

var regionIntersectRegionFunction *gi.Function
var regionIntersectRegionFunction_Once sync.Once

func regionIntersectRegionFunction_Set() error {
	var err error
	regionIntersectRegionFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionIntersectRegionFunction, err = regionObject.InvokerNew("intersect_region")
	})
	return err
}

// IntersectRegion is a representation of the C type gtk_source_region_intersect_region.
func (recv *Region) IntersectRegion(region2 *Region) *Region {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(region2.Native())

	var ret gi.Argument

	err := regionIntersectRegionFunction_Set()
	if err == nil {
		ret = regionIntersectRegionFunction.Invoke(inArgs[:], nil)
	}

	retGo := RegionNewFromNative(ret.Pointer())

	return retGo
}

var regionIntersectSubregionFunction *gi.Function
var regionIntersectSubregionFunction_Once sync.Once

func regionIntersectSubregionFunction_Set() error {
	var err error
	regionIntersectSubregionFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionIntersectSubregionFunction, err = regionObject.InvokerNew("intersect_subregion")
	})
	return err
}

// IntersectSubregion is a representation of the C type gtk_source_region_intersect_subregion.
func (recv *Region) IntersectSubregion(start *gtk.TextIter, end *gtk.TextIter) *Region {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	var ret gi.Argument

	err := regionIntersectSubregionFunction_Set()
	if err == nil {
		ret = regionIntersectSubregionFunction.Invoke(inArgs[:], nil)
	}

	retGo := RegionNewFromNative(ret.Pointer())

	return retGo
}

var regionIsEmptyFunction *gi.Function
var regionIsEmptyFunction_Once sync.Once

func regionIsEmptyFunction_Set() error {
	var err error
	regionIsEmptyFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionIsEmptyFunction, err = regionObject.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type gtk_source_region_is_empty.
func (recv *Region) IsEmpty() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionSubtractRegionFunction, err = regionObject.InvokerNew("subtract_region")
	})
	return err
}

// SubtractRegion is a representation of the C type gtk_source_region_subtract_region.
func (recv *Region) SubtractRegion(regionToSubtract *Region) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(regionToSubtract.Native())

	err := regionSubtractRegionFunction_Set()
	if err == nil {
		regionSubtractRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var regionSubtractSubregionFunction *gi.Function
var regionSubtractSubregionFunction_Once sync.Once

func regionSubtractSubregionFunction_Set() error {
	var err error
	regionSubtractSubregionFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionSubtractSubregionFunction, err = regionObject.InvokerNew("subtract_subregion")
	})
	return err
}

// SubtractSubregion is a representation of the C type gtk_source_region_subtract_subregion.
func (recv *Region) SubtractSubregion(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := regionSubtractSubregionFunction_Set()
	if err == nil {
		regionSubtractSubregionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var regionToStringFunction *gi.Function
var regionToStringFunction_Once sync.Once

func regionToStringFunction_Set() error {
	var err error
	regionToStringFunction_Once.Do(func() {
		err = regionObject_Set()
		if err != nil {
			return
		}
		regionToStringFunction, err = regionObject.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_source_region_to_string.
func (recv *Region) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := regionToStringFunction_Set()
	if err == nil {
		ret = regionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var searchContextObject *gi.Object
var searchContextObject_Once sync.Once

func searchContextObject_Set() error {
	var err error
	searchContextObject_Once.Do(func() {
		searchContextObject, err = gi.ObjectNew("GtkSource", "SearchContext")
	})
	return err
}

type SearchContext struct {
	native unsafe.Pointer
}

func SearchContextNewFromNative(native unsafe.Pointer) *SearchContext {
	instance := &SearchContext{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SearchContext) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSearchContext down casts any arbitrary Object to SearchContext.
Exercise care, as this is a potentially dangerous function
if the Object is not a SearchContext.
*/
func CastToSearchContext(object *gobject.Object) *SearchContext {
	return SearchContextNewFromNative(object.Native())
}

// Equals compares this SearchContext with another SearchContext, and returns true if they represent the same Object.
func (recv *SearchContext) Equals(other *SearchContext) bool {
	return other.Native() == recv.Native()
}

func (recv *SearchContext) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SearchContext) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := searchContextObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(searchContextObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SearchContext) SetFieldParent(value *gobject.Object) {
	err := searchContextObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(searchContextObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SearchContext) FieldPriv() *SearchContextPrivate {
	var nilValue *SearchContextPrivate
	err := searchContextObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(searchContextObject, recv.Native(), "priv")
	value := SearchContextPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SearchContext) SetFieldPriv(value *SearchContextPrivate) {
	err := searchContextObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(searchContextObject, recv.Native(), "priv", argValue)
}

var searchContextNewFunction *gi.Function
var searchContextNewFunction_Once sync.Once

func searchContextNewFunction_Set() error {
	var err error
	searchContextNewFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextNewFunction, err = searchContextObject.InvokerNew("new")
	})
	return err
}

// SearchContextNew is a representation of the C type gtk_source_search_context_new.
func SearchContextNew(buffer *Buffer, settings *SearchSettings) *SearchContext {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(buffer.Native())
	inArgs[1].SetPointer(settings.Native())

	var ret gi.Argument

	err := searchContextNewFunction_Set()
	if err == nil {
		ret = searchContextNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SearchContextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var searchContextBackwardFunction *gi.Function
var searchContextBackwardFunction_Once sync.Once

func searchContextBackwardFunction_Set() error {
	var err error
	searchContextBackwardFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextBackwardFunction, err = searchContextObject.InvokerNew("backward")
	})
	return err
}

// Backward is a representation of the C type gtk_source_search_context_backward.
func (recv *SearchContext) Backward(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := searchContextBackwardFunction_Set()
	if err == nil {
		ret = searchContextBackwardFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())
	out1 := gtk.TextIterNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var searchContextBackward2Function *gi.Function
var searchContextBackward2Function_Once sync.Once

func searchContextBackward2Function_Set() error {
	var err error
	searchContextBackward2Function_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextBackward2Function, err = searchContextObject.InvokerNew("backward2")
	})
	return err
}

// Backward2 is a representation of the C type gtk_source_search_context_backward2.
func (recv *SearchContext) Backward2(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter, bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := searchContextBackward2Function_Set()
	if err == nil {
		ret = searchContextBackward2Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())
	out1 := gtk.TextIterNewFromNative(outArgs[1].Pointer())
	out2 := outArgs[2].Boolean()

	return retGo, out0, out1, out2
}

// UNSUPPORTED : C value 'gtk_source_search_context_backward_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_backward_finish2' : parameter 'result' of type 'Gio.AsyncResult' not supported

var searchContextForwardFunction *gi.Function
var searchContextForwardFunction_Once sync.Once

func searchContextForwardFunction_Set() error {
	var err error
	searchContextForwardFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextForwardFunction, err = searchContextObject.InvokerNew("forward")
	})
	return err
}

// Forward is a representation of the C type gtk_source_search_context_forward.
func (recv *SearchContext) Forward(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := searchContextForwardFunction_Set()
	if err == nil {
		ret = searchContextForwardFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())
	out1 := gtk.TextIterNewFromNative(outArgs[1].Pointer())

	return retGo, out0, out1
}

var searchContextForward2Function *gi.Function
var searchContextForward2Function_Once sync.Once

func searchContextForward2Function_Set() error {
	var err error
	searchContextForward2Function_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextForward2Function, err = searchContextObject.InvokerNew("forward2")
	})
	return err
}

// Forward2 is a representation of the C type gtk_source_search_context_forward2.
func (recv *SearchContext) Forward2(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter, bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := searchContextForward2Function_Set()
	if err == nil {
		ret = searchContextForward2Function.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gtk.TextIterNewFromNative(outArgs[0].Pointer())
	out1 := gtk.TextIterNewFromNative(outArgs[1].Pointer())
	out2 := outArgs[2].Boolean()

	return retGo, out0, out1, out2
}

// UNSUPPORTED : C value 'gtk_source_search_context_forward_async' : parameter 'callback' of type 'Gio.AsyncReadyCallback' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward_finish' : parameter 'result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gtk_source_search_context_forward_finish2' : parameter 'result' of type 'Gio.AsyncResult' not supported

var searchContextGetBufferFunction *gi.Function
var searchContextGetBufferFunction_Once sync.Once

func searchContextGetBufferFunction_Set() error {
	var err error
	searchContextGetBufferFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetBufferFunction, err = searchContextObject.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_source_search_context_get_buffer.
func (recv *SearchContext) GetBuffer() *Buffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := searchContextGetBufferFunction_Set()
	if err == nil {
		ret = searchContextGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := BufferNewFromNative(ret.Pointer())

	return retGo
}

var searchContextGetHighlightFunction *gi.Function
var searchContextGetHighlightFunction_Once sync.Once

func searchContextGetHighlightFunction_Set() error {
	var err error
	searchContextGetHighlightFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetHighlightFunction, err = searchContextObject.InvokerNew("get_highlight")
	})
	return err
}

// GetHighlight is a representation of the C type gtk_source_search_context_get_highlight.
func (recv *SearchContext) GetHighlight() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetMatchStyleFunction, err = searchContextObject.InvokerNew("get_match_style")
	})
	return err
}

// GetMatchStyle is a representation of the C type gtk_source_search_context_get_match_style.
func (recv *SearchContext) GetMatchStyle() *Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := searchContextGetMatchStyleFunction_Set()
	if err == nil {
		ret = searchContextGetMatchStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleNewFromNative(ret.Pointer())

	return retGo
}

var searchContextGetOccurrencePositionFunction *gi.Function
var searchContextGetOccurrencePositionFunction_Once sync.Once

func searchContextGetOccurrencePositionFunction_Set() error {
	var err error
	searchContextGetOccurrencePositionFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetOccurrencePositionFunction, err = searchContextObject.InvokerNew("get_occurrence_position")
	})
	return err
}

// GetOccurrencePosition is a representation of the C type gtk_source_search_context_get_occurrence_position.
func (recv *SearchContext) GetOccurrencePosition(matchStart *gtk.TextIter, matchEnd *gtk.TextIter) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matchStart.Native())
	inArgs[2].SetPointer(matchEnd.Native())

	var ret gi.Argument

	err := searchContextGetOccurrencePositionFunction_Set()
	if err == nil {
		ret = searchContextGetOccurrencePositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var searchContextGetOccurrencesCountFunction *gi.Function
var searchContextGetOccurrencesCountFunction_Once sync.Once

func searchContextGetOccurrencesCountFunction_Set() error {
	var err error
	searchContextGetOccurrencesCountFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetOccurrencesCountFunction, err = searchContextObject.InvokerNew("get_occurrences_count")
	})
	return err
}

// GetOccurrencesCount is a representation of the C type gtk_source_search_context_get_occurrences_count.
func (recv *SearchContext) GetOccurrencesCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := searchContextGetOccurrencesCountFunction_Set()
	if err == nil {
		ret = searchContextGetOccurrencesCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var searchContextGetRegexErrorFunction *gi.Function
var searchContextGetRegexErrorFunction_Once sync.Once

func searchContextGetRegexErrorFunction_Set() error {
	var err error
	searchContextGetRegexErrorFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetRegexErrorFunction, err = searchContextObject.InvokerNew("get_regex_error")
	})
	return err
}

// GetRegexError is a representation of the C type gtk_source_search_context_get_regex_error.
func (recv *SearchContext) GetRegexError() *glib.Error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := searchContextGetRegexErrorFunction_Set()
	if err == nil {
		ret = searchContextGetRegexErrorFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ErrorNewFromNative(ret.Pointer())

	return retGo
}

var searchContextGetSettingsFunction *gi.Function
var searchContextGetSettingsFunction_Once sync.Once

func searchContextGetSettingsFunction_Set() error {
	var err error
	searchContextGetSettingsFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextGetSettingsFunction, err = searchContextObject.InvokerNew("get_settings")
	})
	return err
}

// GetSettings is a representation of the C type gtk_source_search_context_get_settings.
func (recv *SearchContext) GetSettings() *SearchSettings {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := searchContextGetSettingsFunction_Set()
	if err == nil {
		ret = searchContextGetSettingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := SearchSettingsNewFromNative(ret.Pointer())

	return retGo
}

var searchContextReplaceFunction *gi.Function
var searchContextReplaceFunction_Once sync.Once

func searchContextReplaceFunction_Set() error {
	var err error
	searchContextReplaceFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextReplaceFunction, err = searchContextObject.InvokerNew("replace")
	})
	return err
}

// Replace is a representation of the C type gtk_source_search_context_replace.
func (recv *SearchContext) Replace(matchStart *gtk.TextIter, matchEnd *gtk.TextIter, replace string, replaceLength int32) bool {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matchStart.Native())
	inArgs[2].SetPointer(matchEnd.Native())
	inArgs[3].SetString(replace)
	inArgs[4].SetInt32(replaceLength)

	var ret gi.Argument

	err := searchContextReplaceFunction_Set()
	if err == nil {
		ret = searchContextReplaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchContextReplace2Function *gi.Function
var searchContextReplace2Function_Once sync.Once

func searchContextReplace2Function_Set() error {
	var err error
	searchContextReplace2Function_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextReplace2Function, err = searchContextObject.InvokerNew("replace2")
	})
	return err
}

// Replace2 is a representation of the C type gtk_source_search_context_replace2.
func (recv *SearchContext) Replace2(matchStart *gtk.TextIter, matchEnd *gtk.TextIter, replace string, replaceLength int32) bool {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matchStart.Native())
	inArgs[2].SetPointer(matchEnd.Native())
	inArgs[3].SetString(replace)
	inArgs[4].SetInt32(replaceLength)

	var ret gi.Argument

	err := searchContextReplace2Function_Set()
	if err == nil {
		ret = searchContextReplace2Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var searchContextReplaceAllFunction *gi.Function
var searchContextReplaceAllFunction_Once sync.Once

func searchContextReplaceAllFunction_Set() error {
	var err error
	searchContextReplaceAllFunction_Once.Do(func() {
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextReplaceAllFunction, err = searchContextObject.InvokerNew("replace_all")
	})
	return err
}

// ReplaceAll is a representation of the C type gtk_source_search_context_replace_all.
func (recv *SearchContext) ReplaceAll(replace string, replaceLength int32) uint32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextSetHighlightFunction, err = searchContextObject.InvokerNew("set_highlight")
	})
	return err
}

// SetHighlight is a representation of the C type gtk_source_search_context_set_highlight.
func (recv *SearchContext) SetHighlight(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextSetMatchStyleFunction, err = searchContextObject.InvokerNew("set_match_style")
	})
	return err
}

// SetMatchStyle is a representation of the C type gtk_source_search_context_set_match_style.
func (recv *SearchContext) SetMatchStyle(matchStyle *Style) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matchStyle.Native())

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
		err = searchContextObject_Set()
		if err != nil {
			return
		}
		searchContextSetSettingsFunction, err = searchContextObject.InvokerNew("set_settings")
	})
	return err
}

// SetSettings is a representation of the C type gtk_source_search_context_set_settings.
func (recv *SearchContext) SetSettings(settings *SearchSettings) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(settings.Native())

	err := searchContextSetSettingsFunction_Set()
	if err == nil {
		searchContextSetSettingsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var searchSettingsObject *gi.Object
var searchSettingsObject_Once sync.Once

func searchSettingsObject_Set() error {
	var err error
	searchSettingsObject_Once.Do(func() {
		searchSettingsObject, err = gi.ObjectNew("GtkSource", "SearchSettings")
	})
	return err
}

type SearchSettings struct {
	native unsafe.Pointer
}

func SearchSettingsNewFromNative(native unsafe.Pointer) *SearchSettings {
	instance := &SearchSettings{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SearchSettings) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSearchSettings down casts any arbitrary Object to SearchSettings.
Exercise care, as this is a potentially dangerous function
if the Object is not a SearchSettings.
*/
func CastToSearchSettings(object *gobject.Object) *SearchSettings {
	return SearchSettingsNewFromNative(object.Native())
}

// Equals compares this SearchSettings with another SearchSettings, and returns true if they represent the same Object.
func (recv *SearchSettings) Equals(other *SearchSettings) bool {
	return other.Native() == recv.Native()
}

func (recv *SearchSettings) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SearchSettings) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := searchSettingsObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(searchSettingsObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SearchSettings) SetFieldParent(value *gobject.Object) {
	err := searchSettingsObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(searchSettingsObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SearchSettings) FieldPriv() *SearchSettingsPrivate {
	var nilValue *SearchSettingsPrivate
	err := searchSettingsObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(searchSettingsObject, recv.Native(), "priv")
	value := SearchSettingsPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SearchSettings) SetFieldPriv(value *SearchSettingsPrivate) {
	err := searchSettingsObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(searchSettingsObject, recv.Native(), "priv", argValue)
}

var searchSettingsNewFunction *gi.Function
var searchSettingsNewFunction_Once sync.Once

func searchSettingsNewFunction_Set() error {
	var err error
	searchSettingsNewFunction_Once.Do(func() {
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsNewFunction, err = searchSettingsObject.InvokerNew("new")
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

	retGo := SearchSettingsNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var searchSettingsGetAtWordBoundariesFunction *gi.Function
var searchSettingsGetAtWordBoundariesFunction_Once sync.Once

func searchSettingsGetAtWordBoundariesFunction_Set() error {
	var err error
	searchSettingsGetAtWordBoundariesFunction_Once.Do(func() {
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsGetAtWordBoundariesFunction, err = searchSettingsObject.InvokerNew("get_at_word_boundaries")
	})
	return err
}

// GetAtWordBoundaries is a representation of the C type gtk_source_search_settings_get_at_word_boundaries.
func (recv *SearchSettings) GetAtWordBoundaries() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsGetCaseSensitiveFunction, err = searchSettingsObject.InvokerNew("get_case_sensitive")
	})
	return err
}

// GetCaseSensitive is a representation of the C type gtk_source_search_settings_get_case_sensitive.
func (recv *SearchSettings) GetCaseSensitive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsGetRegexEnabledFunction, err = searchSettingsObject.InvokerNew("get_regex_enabled")
	})
	return err
}

// GetRegexEnabled is a representation of the C type gtk_source_search_settings_get_regex_enabled.
func (recv *SearchSettings) GetRegexEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsGetSearchTextFunction, err = searchSettingsObject.InvokerNew("get_search_text")
	})
	return err
}

// GetSearchText is a representation of the C type gtk_source_search_settings_get_search_text.
func (recv *SearchSettings) GetSearchText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsGetWrapAroundFunction, err = searchSettingsObject.InvokerNew("get_wrap_around")
	})
	return err
}

// GetWrapAround is a representation of the C type gtk_source_search_settings_get_wrap_around.
func (recv *SearchSettings) GetWrapAround() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsSetAtWordBoundariesFunction, err = searchSettingsObject.InvokerNew("set_at_word_boundaries")
	})
	return err
}

// SetAtWordBoundaries is a representation of the C type gtk_source_search_settings_set_at_word_boundaries.
func (recv *SearchSettings) SetAtWordBoundaries(atWordBoundaries bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsSetCaseSensitiveFunction, err = searchSettingsObject.InvokerNew("set_case_sensitive")
	})
	return err
}

// SetCaseSensitive is a representation of the C type gtk_source_search_settings_set_case_sensitive.
func (recv *SearchSettings) SetCaseSensitive(caseSensitive bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsSetRegexEnabledFunction, err = searchSettingsObject.InvokerNew("set_regex_enabled")
	})
	return err
}

// SetRegexEnabled is a representation of the C type gtk_source_search_settings_set_regex_enabled.
func (recv *SearchSettings) SetRegexEnabled(regexEnabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsSetSearchTextFunction, err = searchSettingsObject.InvokerNew("set_search_text")
	})
	return err
}

// SetSearchText is a representation of the C type gtk_source_search_settings_set_search_text.
func (recv *SearchSettings) SetSearchText(searchText string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = searchSettingsObject_Set()
		if err != nil {
			return
		}
		searchSettingsSetWrapAroundFunction, err = searchSettingsObject.InvokerNew("set_wrap_around")
	})
	return err
}

// SetWrapAround is a representation of the C type gtk_source_search_settings_set_wrap_around.
func (recv *SearchSettings) SetWrapAround(wrapAround bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(wrapAround)

	err := searchSettingsSetWrapAroundFunction_Set()
	if err == nil {
		searchSettingsSetWrapAroundFunction.Invoke(inArgs[:], nil)
	}

	return
}

var spaceDrawerObject *gi.Object
var spaceDrawerObject_Once sync.Once

func spaceDrawerObject_Set() error {
	var err error
	spaceDrawerObject_Once.Do(func() {
		spaceDrawerObject, err = gi.ObjectNew("GtkSource", "SpaceDrawer")
	})
	return err
}

type SpaceDrawer struct {
	native unsafe.Pointer
}

func SpaceDrawerNewFromNative(native unsafe.Pointer) *SpaceDrawer {
	instance := &SpaceDrawer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *SpaceDrawer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSpaceDrawer down casts any arbitrary Object to SpaceDrawer.
Exercise care, as this is a potentially dangerous function
if the Object is not a SpaceDrawer.
*/
func CastToSpaceDrawer(object *gobject.Object) *SpaceDrawer {
	return SpaceDrawerNewFromNative(object.Native())
}

// Equals compares this SpaceDrawer with another SpaceDrawer, and returns true if they represent the same Object.
func (recv *SpaceDrawer) Equals(other *SpaceDrawer) bool {
	return other.Native() == recv.Native()
}

func (recv *SpaceDrawer) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SpaceDrawer) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := spaceDrawerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(spaceDrawerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SpaceDrawer) SetFieldParent(value *gobject.Object) {
	err := spaceDrawerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(spaceDrawerObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *SpaceDrawer) FieldPriv() *SpaceDrawerPrivate {
	var nilValue *SpaceDrawerPrivate
	err := spaceDrawerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(spaceDrawerObject, recv.Native(), "priv")
	value := SpaceDrawerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *SpaceDrawer) SetFieldPriv(value *SpaceDrawerPrivate) {
	err := spaceDrawerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(spaceDrawerObject, recv.Native(), "priv", argValue)
}

var spaceDrawerNewFunction *gi.Function
var spaceDrawerNewFunction_Once sync.Once

func spaceDrawerNewFunction_Set() error {
	var err error
	spaceDrawerNewFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerNewFunction, err = spaceDrawerObject.InvokerNew("new")
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

	retGo := SpaceDrawerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_space_drawer_bind_matrix_setting' : parameter 'flags' of type 'Gio.SettingsBindFlags' not supported

var spaceDrawerGetEnableMatrixFunction *gi.Function
var spaceDrawerGetEnableMatrixFunction_Once sync.Once

func spaceDrawerGetEnableMatrixFunction_Set() error {
	var err error
	spaceDrawerGetEnableMatrixFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerGetEnableMatrixFunction, err = spaceDrawerObject.InvokerNew("get_enable_matrix")
	})
	return err
}

// GetEnableMatrix is a representation of the C type gtk_source_space_drawer_get_enable_matrix.
func (recv *SpaceDrawer) GetEnableMatrix() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := spaceDrawerGetEnableMatrixFunction_Set()
	if err == nil {
		ret = spaceDrawerGetEnableMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var spaceDrawerGetMatrixFunction *gi.Function
var spaceDrawerGetMatrixFunction_Once sync.Once

func spaceDrawerGetMatrixFunction_Set() error {
	var err error
	spaceDrawerGetMatrixFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerGetMatrixFunction, err = spaceDrawerObject.InvokerNew("get_matrix")
	})
	return err
}

// GetMatrix is a representation of the C type gtk_source_space_drawer_get_matrix.
func (recv *SpaceDrawer) GetMatrix() *glib.Variant {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := spaceDrawerGetMatrixFunction_Set()
	if err == nil {
		ret = spaceDrawerGetMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}

var spaceDrawerGetTypesForLocationsFunction *gi.Function
var spaceDrawerGetTypesForLocationsFunction_Once sync.Once

func spaceDrawerGetTypesForLocationsFunction_Set() error {
	var err error
	spaceDrawerGetTypesForLocationsFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerGetTypesForLocationsFunction, err = spaceDrawerObject.InvokerNew("get_types_for_locations")
	})
	return err
}

// GetTypesForLocations is a representation of the C type gtk_source_space_drawer_get_types_for_locations.
func (recv *SpaceDrawer) GetTypesForLocations(locations SpaceLocationFlags) SpaceTypeFlags {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(locations))

	var ret gi.Argument

	err := spaceDrawerGetTypesForLocationsFunction_Set()
	if err == nil {
		ret = spaceDrawerGetTypesForLocationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := SpaceTypeFlags(ret.Int32())

	return retGo
}

var spaceDrawerSetEnableMatrixFunction *gi.Function
var spaceDrawerSetEnableMatrixFunction_Once sync.Once

func spaceDrawerSetEnableMatrixFunction_Set() error {
	var err error
	spaceDrawerSetEnableMatrixFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerSetEnableMatrixFunction, err = spaceDrawerObject.InvokerNew("set_enable_matrix")
	})
	return err
}

// SetEnableMatrix is a representation of the C type gtk_source_space_drawer_set_enable_matrix.
func (recv *SpaceDrawer) SetEnableMatrix(enableMatrix bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(enableMatrix)

	err := spaceDrawerSetEnableMatrixFunction_Set()
	if err == nil {
		spaceDrawerSetEnableMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

var spaceDrawerSetMatrixFunction *gi.Function
var spaceDrawerSetMatrixFunction_Once sync.Once

func spaceDrawerSetMatrixFunction_Set() error {
	var err error
	spaceDrawerSetMatrixFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerSetMatrixFunction, err = spaceDrawerObject.InvokerNew("set_matrix")
	})
	return err
}

// SetMatrix is a representation of the C type gtk_source_space_drawer_set_matrix.
func (recv *SpaceDrawer) SetMatrix(matrix *glib.Variant) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matrix.Native())

	err := spaceDrawerSetMatrixFunction_Set()
	if err == nil {
		spaceDrawerSetMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

var spaceDrawerSetTypesForLocationsFunction *gi.Function
var spaceDrawerSetTypesForLocationsFunction_Once sync.Once

func spaceDrawerSetTypesForLocationsFunction_Set() error {
	var err error
	spaceDrawerSetTypesForLocationsFunction_Once.Do(func() {
		err = spaceDrawerObject_Set()
		if err != nil {
			return
		}
		spaceDrawerSetTypesForLocationsFunction, err = spaceDrawerObject.InvokerNew("set_types_for_locations")
	})
	return err
}

// SetTypesForLocations is a representation of the C type gtk_source_space_drawer_set_types_for_locations.
func (recv *SpaceDrawer) SetTypesForLocations(locations SpaceLocationFlags, types SpaceTypeFlags) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(locations))
	inArgs[2].SetInt32(int32(types))

	err := spaceDrawerSetTypesForLocationsFunction_Set()
	if err == nil {
		spaceDrawerSetTypesForLocationsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleObject *gi.Object
var styleObject_Once sync.Once

func styleObject_Set() error {
	var err error
	styleObject_Once.Do(func() {
		styleObject, err = gi.ObjectNew("GtkSource", "Style")
	})
	return err
}

type Style struct {
	native unsafe.Pointer
}

func StyleNewFromNative(native unsafe.Pointer) *Style {
	instance := &Style{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Style) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStyle down casts any arbitrary Object to Style.
Exercise care, as this is a potentially dangerous function
if the Object is not a Style.
*/
func CastToStyle(object *gobject.Object) *Style {
	return StyleNewFromNative(object.Native())
}

// Equals compares this Style with another Style, and returns true if they represent the same Object.
func (recv *Style) Equals(other *Style) bool {
	return other.Native() == recv.Native()
}

func (recv *Style) Native() unsafe.Pointer {
	return recv.native
}

var styleApplyFunction *gi.Function
var styleApplyFunction_Once sync.Once

func styleApplyFunction_Set() error {
	var err error
	styleApplyFunction_Once.Do(func() {
		err = styleObject_Set()
		if err != nil {
			return
		}
		styleApplyFunction, err = styleObject.InvokerNew("apply")
	})
	return err
}

// Apply is a representation of the C type gtk_source_style_apply.
func (recv *Style) Apply(tag *gtk.TextTag) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(tag.Native())

	err := styleApplyFunction_Set()
	if err == nil {
		styleApplyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleCopyFunction *gi.Function
var styleCopyFunction_Once sync.Once

func styleCopyFunction_Set() error {
	var err error
	styleCopyFunction_Once.Do(func() {
		err = styleObject_Set()
		if err != nil {
			return
		}
		styleCopyFunction, err = styleObject.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_source_style_copy.
func (recv *Style) Copy() *Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := styleCopyFunction_Set()
	if err == nil {
		ret = styleCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleNewFromNative(ret.Pointer())

	return retGo
}

var styleSchemeObject *gi.Object
var styleSchemeObject_Once sync.Once

func styleSchemeObject_Set() error {
	var err error
	styleSchemeObject_Once.Do(func() {
		styleSchemeObject, err = gi.ObjectNew("GtkSource", "StyleScheme")
	})
	return err
}

type StyleScheme struct {
	native unsafe.Pointer
}

func StyleSchemeNewFromNative(native unsafe.Pointer) *StyleScheme {
	instance := &StyleScheme{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *StyleScheme) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStyleScheme down casts any arbitrary Object to StyleScheme.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleScheme.
*/
func CastToStyleScheme(object *gobject.Object) *StyleScheme {
	return StyleSchemeNewFromNative(object.Native())
}

// Equals compares this StyleScheme with another StyleScheme, and returns true if they represent the same Object.
func (recv *StyleScheme) Equals(other *StyleScheme) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleScheme) Native() unsafe.Pointer {
	return recv.native
}

// FieldBase returns the C field 'base'.
func (recv *StyleScheme) FieldBase() *gobject.Object {
	var nilValue *gobject.Object
	err := styleSchemeObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeObject, recv.Native(), "base")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBase sets the value of the C field 'base'.
func (recv *StyleScheme) SetFieldBase(value *gobject.Object) {
	err := styleSchemeObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeObject, recv.Native(), "base", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *StyleScheme) FieldPriv() *StyleSchemePrivate {
	var nilValue *StyleSchemePrivate
	err := styleSchemeObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeObject, recv.Native(), "priv")
	value := StyleSchemePrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *StyleScheme) SetFieldPriv(value *StyleSchemePrivate) {
	err := styleSchemeObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeObject, recv.Native(), "priv", argValue)
}

var styleSchemeGetAuthorsFunction *gi.Function
var styleSchemeGetAuthorsFunction_Once sync.Once

func styleSchemeGetAuthorsFunction_Set() error {
	var err error
	styleSchemeGetAuthorsFunction_Once.Do(func() {
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetAuthorsFunction, err = styleSchemeObject.InvokerNew("get_authors")
	})
	return err
}

// GetAuthors is a representation of the C type gtk_source_style_scheme_get_authors.
func (recv *StyleScheme) GetAuthors() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetDescriptionFunction, err = styleSchemeObject.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type gtk_source_style_scheme_get_description.
func (recv *StyleScheme) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetFilenameFunction, err = styleSchemeObject.InvokerNew("get_filename")
	})
	return err
}

// GetFilename is a representation of the C type gtk_source_style_scheme_get_filename.
func (recv *StyleScheme) GetFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetIdFunction, err = styleSchemeObject.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type gtk_source_style_scheme_get_id.
func (recv *StyleScheme) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetNameFunction, err = styleSchemeObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_style_scheme_get_name.
func (recv *StyleScheme) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeObject_Set()
		if err != nil {
			return
		}
		styleSchemeGetStyleFunction, err = styleSchemeObject.InvokerNew("get_style")
	})
	return err
}

// GetStyle is a representation of the C type gtk_source_style_scheme_get_style.
func (recv *StyleScheme) GetStyle(styleId string) *Style {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(styleId)

	var ret gi.Argument

	err := styleSchemeGetStyleFunction_Set()
	if err == nil {
		ret = styleSchemeGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleNewFromNative(ret.Pointer())

	return retGo
}

var styleSchemeChooserButtonObject *gi.Object
var styleSchemeChooserButtonObject_Once sync.Once

func styleSchemeChooserButtonObject_Set() error {
	var err error
	styleSchemeChooserButtonObject_Once.Do(func() {
		styleSchemeChooserButtonObject, err = gi.ObjectNew("GtkSource", "StyleSchemeChooserButton")
	})
	return err
}

type StyleSchemeChooserButton struct {
	native unsafe.Pointer
}

func StyleSchemeChooserButtonNewFromNative(native unsafe.Pointer) *StyleSchemeChooserButton {
	instance := &StyleSchemeChooserButton{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Button upcasts to *Button
func (recv *StyleSchemeChooserButton) Button() *gtk.Button {
	return gtk.ButtonNewFromNative(recv.Native())
}

// Bin upcasts to *Bin
func (recv *StyleSchemeChooserButton) Bin() *gtk.Bin {
	return gtk.BinNewFromNative(recv.Native())
}

// Container upcasts to *Container
func (recv *StyleSchemeChooserButton) Container() *gtk.Container {
	return gtk.ContainerNewFromNative(recv.Native())
}

// Widget upcasts to *Widget
func (recv *StyleSchemeChooserButton) Widget() *gtk.Widget {
	return gtk.WidgetNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StyleSchemeChooserButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *StyleSchemeChooserButton) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStyleSchemeChooserButton down casts any arbitrary Object to StyleSchemeChooserButton.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleSchemeChooserButton.
*/
func CastToStyleSchemeChooserButton(object *gobject.Object) *StyleSchemeChooserButton {
	return StyleSchemeChooserButtonNewFromNative(object.Native())
}

// Equals compares this StyleSchemeChooserButton with another StyleSchemeChooserButton, and returns true if they represent the same Object.
func (recv *StyleSchemeChooserButton) Equals(other *StyleSchemeChooserButton) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleSchemeChooserButton) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StyleSchemeChooserButton) FieldParent() *gtk.Button {
	var nilValue *gtk.Button
	err := styleSchemeChooserButtonObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeChooserButtonObject, recv.Native(), "parent")
	value := gtk.ButtonNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StyleSchemeChooserButton) SetFieldParent(value *gtk.Button) {
	err := styleSchemeChooserButtonObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeChooserButtonObject, recv.Native(), "parent", argValue)
}

var styleSchemeChooserButtonNewFunction *gi.Function
var styleSchemeChooserButtonNewFunction_Once sync.Once

func styleSchemeChooserButtonNewFunction_Set() error {
	var err error
	styleSchemeChooserButtonNewFunction_Once.Do(func() {
		err = styleSchemeChooserButtonObject_Set()
		if err != nil {
			return
		}
		styleSchemeChooserButtonNewFunction, err = styleSchemeChooserButtonObject.InvokerNew("new")
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

	retGo := StyleSchemeChooserButtonNewFromNative(ret.Pointer())

	return retGo
}

// AtkImplementorIface returns the Atk.ImplementorIface interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) AtkImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromNative(recv.Native())
}

// GtkActionable returns the Gtk.Actionable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) GtkActionable() *gtk.Actionable {
	return gtk.ActionableNewFromNative(recv.Native())
}

// GtkActivatable returns the Gtk.Activatable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) GtkActivatable() *gtk.Activatable {
	return gtk.ActivatableNewFromNative(recv.Native())
}

// GtkBuildable returns the Gtk.Buildable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

// StyleSchemeChooser returns the StyleSchemeChooser interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) StyleSchemeChooser() *StyleSchemeChooser {
	return StyleSchemeChooserNewFromNative(recv.Native())
}

var styleSchemeChooserWidgetObject *gi.Object
var styleSchemeChooserWidgetObject_Once sync.Once

func styleSchemeChooserWidgetObject_Set() error {
	var err error
	styleSchemeChooserWidgetObject_Once.Do(func() {
		styleSchemeChooserWidgetObject, err = gi.ObjectNew("GtkSource", "StyleSchemeChooserWidget")
	})
	return err
}

type StyleSchemeChooserWidget struct {
	native unsafe.Pointer
}

func StyleSchemeChooserWidgetNewFromNative(native unsafe.Pointer) *StyleSchemeChooserWidget {
	instance := &StyleSchemeChooserWidget{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Bin upcasts to *Bin
func (recv *StyleSchemeChooserWidget) Bin() *gtk.Bin {
	return gtk.BinNewFromNative(recv.Native())
}

// Container upcasts to *Container
func (recv *StyleSchemeChooserWidget) Container() *gtk.Container {
	return gtk.ContainerNewFromNative(recv.Native())
}

// Widget upcasts to *Widget
func (recv *StyleSchemeChooserWidget) Widget() *gtk.Widget {
	return gtk.WidgetNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StyleSchemeChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *StyleSchemeChooserWidget) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStyleSchemeChooserWidget down casts any arbitrary Object to StyleSchemeChooserWidget.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleSchemeChooserWidget.
*/
func CastToStyleSchemeChooserWidget(object *gobject.Object) *StyleSchemeChooserWidget {
	return StyleSchemeChooserWidgetNewFromNative(object.Native())
}

// Equals compares this StyleSchemeChooserWidget with another StyleSchemeChooserWidget, and returns true if they represent the same Object.
func (recv *StyleSchemeChooserWidget) Equals(other *StyleSchemeChooserWidget) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleSchemeChooserWidget) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StyleSchemeChooserWidget) FieldParent() *gtk.Bin {
	var nilValue *gtk.Bin
	err := styleSchemeChooserWidgetObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeChooserWidgetObject, recv.Native(), "parent")
	value := gtk.BinNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StyleSchemeChooserWidget) SetFieldParent(value *gtk.Bin) {
	err := styleSchemeChooserWidgetObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeChooserWidgetObject, recv.Native(), "parent", argValue)
}

var styleSchemeChooserWidgetNewFunction *gi.Function
var styleSchemeChooserWidgetNewFunction_Once sync.Once

func styleSchemeChooserWidgetNewFunction_Set() error {
	var err error
	styleSchemeChooserWidgetNewFunction_Once.Do(func() {
		err = styleSchemeChooserWidgetObject_Set()
		if err != nil {
			return
		}
		styleSchemeChooserWidgetNewFunction, err = styleSchemeChooserWidgetObject.InvokerNew("new")
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

	retGo := StyleSchemeChooserWidgetNewFromNative(ret.Pointer())

	return retGo
}

// AtkImplementorIface returns the Atk.ImplementorIface interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) AtkImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromNative(recv.Native())
}

// GtkBuildable returns the Gtk.Buildable interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

// StyleSchemeChooser returns the StyleSchemeChooser interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) StyleSchemeChooser() *StyleSchemeChooser {
	return StyleSchemeChooserNewFromNative(recv.Native())
}

var styleSchemeManagerObject *gi.Object
var styleSchemeManagerObject_Once sync.Once

func styleSchemeManagerObject_Set() error {
	var err error
	styleSchemeManagerObject_Once.Do(func() {
		styleSchemeManagerObject, err = gi.ObjectNew("GtkSource", "StyleSchemeManager")
	})
	return err
}

type StyleSchemeManager struct {
	native unsafe.Pointer
}

func StyleSchemeManagerNewFromNative(native unsafe.Pointer) *StyleSchemeManager {
	instance := &StyleSchemeManager{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *StyleSchemeManager) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStyleSchemeManager down casts any arbitrary Object to StyleSchemeManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleSchemeManager.
*/
func CastToStyleSchemeManager(object *gobject.Object) *StyleSchemeManager {
	return StyleSchemeManagerNewFromNative(object.Native())
}

// Equals compares this StyleSchemeManager with another StyleSchemeManager, and returns true if they represent the same Object.
func (recv *StyleSchemeManager) Equals(other *StyleSchemeManager) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleSchemeManager) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StyleSchemeManager) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := styleSchemeManagerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeManagerObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StyleSchemeManager) SetFieldParent(value *gobject.Object) {
	err := styleSchemeManagerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeManagerObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *StyleSchemeManager) FieldPriv() *StyleSchemeManagerPrivate {
	var nilValue *StyleSchemeManagerPrivate
	err := styleSchemeManagerObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(styleSchemeManagerObject, recv.Native(), "priv")
	value := StyleSchemeManagerPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *StyleSchemeManager) SetFieldPriv(value *StyleSchemeManagerPrivate) {
	err := styleSchemeManagerObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(styleSchemeManagerObject, recv.Native(), "priv", argValue)
}

var styleSchemeManagerNewFunction *gi.Function
var styleSchemeManagerNewFunction_Once sync.Once

func styleSchemeManagerNewFunction_Set() error {
	var err error
	styleSchemeManagerNewFunction_Once.Do(func() {
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerNewFunction, err = styleSchemeManagerObject.InvokerNew("new")
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

	retGo := StyleSchemeManagerNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var styleSchemeManagerAppendSearchPathFunction *gi.Function
var styleSchemeManagerAppendSearchPathFunction_Once sync.Once

func styleSchemeManagerAppendSearchPathFunction_Set() error {
	var err error
	styleSchemeManagerAppendSearchPathFunction_Once.Do(func() {
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerAppendSearchPathFunction, err = styleSchemeManagerObject.InvokerNew("append_search_path")
	})
	return err
}

// AppendSearchPath is a representation of the C type gtk_source_style_scheme_manager_append_search_path.
func (recv *StyleSchemeManager) AppendSearchPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerForceRescanFunction, err = styleSchemeManagerObject.InvokerNew("force_rescan")
	})
	return err
}

// ForceRescan is a representation of the C type gtk_source_style_scheme_manager_force_rescan.
func (recv *StyleSchemeManager) ForceRescan() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSchemeFunction, err = styleSchemeManagerObject.InvokerNew("get_scheme")
	})
	return err
}

// GetScheme is a representation of the C type gtk_source_style_scheme_manager_get_scheme.
func (recv *StyleSchemeManager) GetScheme(schemeId string) *StyleScheme {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(schemeId)

	var ret gi.Argument

	err := styleSchemeManagerGetSchemeFunction_Set()
	if err == nil {
		ret = styleSchemeManagerGetSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleSchemeNewFromNative(ret.Pointer())

	return retGo
}

var styleSchemeManagerGetSchemeIdsFunction *gi.Function
var styleSchemeManagerGetSchemeIdsFunction_Once sync.Once

func styleSchemeManagerGetSchemeIdsFunction_Set() error {
	var err error
	styleSchemeManagerGetSchemeIdsFunction_Once.Do(func() {
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSchemeIdsFunction, err = styleSchemeManagerObject.InvokerNew("get_scheme_ids")
	})
	return err
}

// GetSchemeIds is a representation of the C type gtk_source_style_scheme_manager_get_scheme_ids.
func (recv *StyleSchemeManager) GetSchemeIds() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerGetSearchPathFunction, err = styleSchemeManagerObject.InvokerNew("get_search_path")
	})
	return err
}

// GetSearchPath is a representation of the C type gtk_source_style_scheme_manager_get_search_path.
func (recv *StyleSchemeManager) GetSearchPath() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = styleSchemeManagerObject_Set()
		if err != nil {
			return
		}
		styleSchemeManagerPrependSearchPathFunction, err = styleSchemeManagerObject.InvokerNew("prepend_search_path")
	})
	return err
}

// PrependSearchPath is a representation of the C type gtk_source_style_scheme_manager_prepend_search_path.
func (recv *StyleSchemeManager) PrependSearchPath(path string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(path)

	err := styleSchemeManagerPrependSearchPathFunction_Set()
	if err == nil {
		styleSchemeManagerPrependSearchPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_style_scheme_manager_set_search_path' : parameter 'path' of type 'nil' not supported

var tagObject *gi.Object
var tagObject_Once sync.Once

func tagObject_Set() error {
	var err error
	tagObject_Once.Do(func() {
		tagObject, err = gi.ObjectNew("GtkSource", "Tag")
	})
	return err
}

type Tag struct {
	native unsafe.Pointer
}

func TagNewFromNative(native unsafe.Pointer) *Tag {
	instance := &Tag{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TextTag upcasts to *TextTag
func (recv *Tag) TextTag() *gtk.TextTag {
	return gtk.TextTagNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Tag) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToTag down casts any arbitrary Object to Tag.
Exercise care, as this is a potentially dangerous function
if the Object is not a Tag.
*/
func CastToTag(object *gobject.Object) *Tag {
	return TagNewFromNative(object.Native())
}

// Equals compares this Tag with another Tag, and returns true if they represent the same Object.
func (recv *Tag) Equals(other *Tag) bool {
	return other.Native() == recv.Native()
}

func (recv *Tag) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Tag) FieldParentInstance() *gtk.TextTag {
	var nilValue *gtk.TextTag
	err := tagObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(tagObject, recv.Native(), "parent_instance")
	value := gtk.TextTagNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Tag) SetFieldParentInstance(value *gtk.TextTag) {
	err := tagObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(tagObject, recv.Native(), "parent_instance", argValue)
}

var tagNewFunction *gi.Function
var tagNewFunction_Once sync.Once

func tagNewFunction_Set() error {
	var err error
	tagNewFunction_Once.Do(func() {
		err = tagObject_Set()
		if err != nil {
			return
		}
		tagNewFunction, err = tagObject.InvokerNew("new")
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

	retGo := TagNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var viewObject *gi.Object
var viewObject_Once sync.Once

func viewObject_Set() error {
	var err error
	viewObject_Once.Do(func() {
		viewObject, err = gi.ObjectNew("GtkSource", "View")
	})
	return err
}

type View struct {
	native unsafe.Pointer
}

func ViewNewFromNative(native unsafe.Pointer) *View {
	instance := &View{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// TextView upcasts to *TextView
func (recv *View) TextView() *gtk.TextView {
	return gtk.TextViewNewFromNative(recv.Native())
}

// Container upcasts to *Container
func (recv *View) Container() *gtk.Container {
	return gtk.ContainerNewFromNative(recv.Native())
}

// Widget upcasts to *Widget
func (recv *View) Widget() *gtk.Widget {
	return gtk.WidgetNewFromNative(recv.Native())
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *View) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *View) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToView down casts any arbitrary Object to View.
Exercise care, as this is a potentially dangerous function
if the Object is not a View.
*/
func CastToView(object *gobject.Object) *View {
	return ViewNewFromNative(object.Native())
}

// Equals compares this View with another View, and returns true if they represent the same Object.
func (recv *View) Equals(other *View) bool {
	return other.Native() == recv.Native()
}

func (recv *View) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *View) FieldParent() *gtk.TextView {
	var nilValue *gtk.TextView
	err := viewObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(viewObject, recv.Native(), "parent")
	value := gtk.TextViewNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *View) SetFieldParent(value *gtk.TextView) {
	err := viewObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(viewObject, recv.Native(), "parent", argValue)
}

// FieldPriv returns the C field 'priv'.
func (recv *View) FieldPriv() *ViewPrivate {
	var nilValue *ViewPrivate
	err := viewObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(viewObject, recv.Native(), "priv")
	value := ViewPrivateNewFromNative(argValue.Pointer())
	return value
}

// SetFieldPriv sets the value of the C field 'priv'.
func (recv *View) SetFieldPriv(value *ViewPrivate) {
	err := viewObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(viewObject, recv.Native(), "priv", argValue)
}

var viewNewFunction *gi.Function
var viewNewFunction_Once sync.Once

func viewNewFunction_Set() error {
	var err error
	viewNewFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewNewFunction, err = viewObject.InvokerNew("new")
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

	retGo := ViewNewFromNative(ret.Pointer())

	return retGo
}

var viewNewWithBufferFunction *gi.Function
var viewNewWithBufferFunction_Once sync.Once

func viewNewWithBufferFunction_Set() error {
	var err error
	viewNewWithBufferFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewNewWithBufferFunction, err = viewObject.InvokerNew("new_with_buffer")
	})
	return err
}

// ViewNewWithBuffer is a representation of the C type gtk_source_view_new_with_buffer.
func ViewNewWithBuffer(buffer *Buffer) *View {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(buffer.Native())

	var ret gi.Argument

	err := viewNewWithBufferFunction_Set()
	if err == nil {
		ret = viewNewWithBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := ViewNewFromNative(ret.Pointer())

	return retGo
}

var viewGetAutoIndentFunction *gi.Function
var viewGetAutoIndentFunction_Once sync.Once

func viewGetAutoIndentFunction_Set() error {
	var err error
	viewGetAutoIndentFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetAutoIndentFunction, err = viewObject.InvokerNew("get_auto_indent")
	})
	return err
}

// GetAutoIndent is a representation of the C type gtk_source_view_get_auto_indent.
func (recv *View) GetAutoIndent() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetBackgroundPatternFunction, err = viewObject.InvokerNew("get_background_pattern")
	})
	return err
}

// GetBackgroundPattern is a representation of the C type gtk_source_view_get_background_pattern.
func (recv *View) GetBackgroundPattern() BackgroundPatternType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetCompletionFunction, err = viewObject.InvokerNew("get_completion")
	})
	return err
}

// GetCompletion is a representation of the C type gtk_source_view_get_completion.
func (recv *View) GetCompletion() *Completion {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := viewGetCompletionFunction_Set()
	if err == nil {
		ret = viewGetCompletionFunction.Invoke(inArgs[:], nil)
	}

	retGo := CompletionNewFromNative(ret.Pointer())

	return retGo
}

var viewGetDrawSpacesFunction *gi.Function
var viewGetDrawSpacesFunction_Once sync.Once

func viewGetDrawSpacesFunction_Set() error {
	var err error
	viewGetDrawSpacesFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetDrawSpacesFunction, err = viewObject.InvokerNew("get_draw_spaces")
	})
	return err
}

// GetDrawSpaces is a representation of the C type gtk_source_view_get_draw_spaces.
func (recv *View) GetDrawSpaces() DrawSpacesFlags {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := viewGetDrawSpacesFunction_Set()
	if err == nil {
		ret = viewGetDrawSpacesFunction.Invoke(inArgs[:], nil)
	}

	retGo := DrawSpacesFlags(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_view_get_gutter' : parameter 'window_type' of type 'Gtk.TextWindowType' not supported

var viewGetHighlightCurrentLineFunction *gi.Function
var viewGetHighlightCurrentLineFunction_Once sync.Once

func viewGetHighlightCurrentLineFunction_Set() error {
	var err error
	viewGetHighlightCurrentLineFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetHighlightCurrentLineFunction, err = viewObject.InvokerNew("get_highlight_current_line")
	})
	return err
}

// GetHighlightCurrentLine is a representation of the C type gtk_source_view_get_highlight_current_line.
func (recv *View) GetHighlightCurrentLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetIndentOnTabFunction, err = viewObject.InvokerNew("get_indent_on_tab")
	})
	return err
}

// GetIndentOnTab is a representation of the C type gtk_source_view_get_indent_on_tab.
func (recv *View) GetIndentOnTab() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetIndentWidthFunction, err = viewObject.InvokerNew("get_indent_width")
	})
	return err
}

// GetIndentWidth is a representation of the C type gtk_source_view_get_indent_width.
func (recv *View) GetIndentWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetInsertSpacesInsteadOfTabsFunction, err = viewObject.InvokerNew("get_insert_spaces_instead_of_tabs")
	})
	return err
}

// GetInsertSpacesInsteadOfTabs is a representation of the C type gtk_source_view_get_insert_spaces_instead_of_tabs.
func (recv *View) GetInsertSpacesInsteadOfTabs() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetMarkAttributesFunction, err = viewObject.InvokerNew("get_mark_attributes")
	})
	return err
}

// GetMarkAttributes is a representation of the C type gtk_source_view_get_mark_attributes.
func (recv *View) GetMarkAttributes(category string, priority int32) *MarkAttributes {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(category)
	inArgs[2].SetInt32(priority)

	var ret gi.Argument

	err := viewGetMarkAttributesFunction_Set()
	if err == nil {
		ret = viewGetMarkAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := MarkAttributesNewFromNative(ret.Pointer())

	return retGo
}

var viewGetRightMarginPositionFunction *gi.Function
var viewGetRightMarginPositionFunction_Once sync.Once

func viewGetRightMarginPositionFunction_Set() error {
	var err error
	viewGetRightMarginPositionFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetRightMarginPositionFunction, err = viewObject.InvokerNew("get_right_margin_position")
	})
	return err
}

// GetRightMarginPosition is a representation of the C type gtk_source_view_get_right_margin_position.
func (recv *View) GetRightMarginPosition() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetShowLineMarksFunction, err = viewObject.InvokerNew("get_show_line_marks")
	})
	return err
}

// GetShowLineMarks is a representation of the C type gtk_source_view_get_show_line_marks.
func (recv *View) GetShowLineMarks() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetShowLineNumbersFunction, err = viewObject.InvokerNew("get_show_line_numbers")
	})
	return err
}

// GetShowLineNumbers is a representation of the C type gtk_source_view_get_show_line_numbers.
func (recv *View) GetShowLineNumbers() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetShowRightMarginFunction, err = viewObject.InvokerNew("get_show_right_margin")
	})
	return err
}

// GetShowRightMargin is a representation of the C type gtk_source_view_get_show_right_margin.
func (recv *View) GetShowRightMargin() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetSmartBackspaceFunction, err = viewObject.InvokerNew("get_smart_backspace")
	})
	return err
}

// GetSmartBackspace is a representation of the C type gtk_source_view_get_smart_backspace.
func (recv *View) GetSmartBackspace() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetSmartHomeEndFunction, err = viewObject.InvokerNew("get_smart_home_end")
	})
	return err
}

// GetSmartHomeEnd is a representation of the C type gtk_source_view_get_smart_home_end.
func (recv *View) GetSmartHomeEnd() SmartHomeEndType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetSpaceDrawerFunction, err = viewObject.InvokerNew("get_space_drawer")
	})
	return err
}

// GetSpaceDrawer is a representation of the C type gtk_source_view_get_space_drawer.
func (recv *View) GetSpaceDrawer() *SpaceDrawer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := viewGetSpaceDrawerFunction_Set()
	if err == nil {
		ret = viewGetSpaceDrawerFunction.Invoke(inArgs[:], nil)
	}

	retGo := SpaceDrawerNewFromNative(ret.Pointer())

	return retGo
}

var viewGetTabWidthFunction *gi.Function
var viewGetTabWidthFunction_Once sync.Once

func viewGetTabWidthFunction_Set() error {
	var err error
	viewGetTabWidthFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetTabWidthFunction, err = viewObject.InvokerNew("get_tab_width")
	})
	return err
}

// GetTabWidth is a representation of the C type gtk_source_view_get_tab_width.
func (recv *View) GetTabWidth() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := viewGetTabWidthFunction_Set()
	if err == nil {
		ret = viewGetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var viewGetVisualColumnFunction *gi.Function
var viewGetVisualColumnFunction_Once sync.Once

func viewGetVisualColumnFunction_Set() error {
	var err error
	viewGetVisualColumnFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewGetVisualColumnFunction, err = viewObject.InvokerNew("get_visual_column")
	})
	return err
}

// GetVisualColumn is a representation of the C type gtk_source_view_get_visual_column.
func (recv *View) GetVisualColumn(iter *gtk.TextIter) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := viewGetVisualColumnFunction_Set()
	if err == nil {
		ret = viewGetVisualColumnFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var viewIndentLinesFunction *gi.Function
var viewIndentLinesFunction_Once sync.Once

func viewIndentLinesFunction_Set() error {
	var err error
	viewIndentLinesFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewIndentLinesFunction, err = viewObject.InvokerNew("indent_lines")
	})
	return err
}

// IndentLines is a representation of the C type gtk_source_view_indent_lines.
func (recv *View) IndentLines(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := viewIndentLinesFunction_Set()
	if err == nil {
		viewIndentLinesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetAutoIndentFunction *gi.Function
var viewSetAutoIndentFunction_Once sync.Once

func viewSetAutoIndentFunction_Set() error {
	var err error
	viewSetAutoIndentFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetAutoIndentFunction, err = viewObject.InvokerNew("set_auto_indent")
	})
	return err
}

// SetAutoIndent is a representation of the C type gtk_source_view_set_auto_indent.
func (recv *View) SetAutoIndent(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetBackgroundPatternFunction, err = viewObject.InvokerNew("set_background_pattern")
	})
	return err
}

// SetBackgroundPattern is a representation of the C type gtk_source_view_set_background_pattern.
func (recv *View) SetBackgroundPattern(backgroundPattern BackgroundPatternType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(backgroundPattern))

	err := viewSetBackgroundPatternFunction_Set()
	if err == nil {
		viewSetBackgroundPatternFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetDrawSpacesFunction *gi.Function
var viewSetDrawSpacesFunction_Once sync.Once

func viewSetDrawSpacesFunction_Set() error {
	var err error
	viewSetDrawSpacesFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetDrawSpacesFunction, err = viewObject.InvokerNew("set_draw_spaces")
	})
	return err
}

// SetDrawSpaces is a representation of the C type gtk_source_view_set_draw_spaces.
func (recv *View) SetDrawSpaces(flags DrawSpacesFlags) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(flags))

	err := viewSetDrawSpacesFunction_Set()
	if err == nil {
		viewSetDrawSpacesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewSetHighlightCurrentLineFunction *gi.Function
var viewSetHighlightCurrentLineFunction_Once sync.Once

func viewSetHighlightCurrentLineFunction_Set() error {
	var err error
	viewSetHighlightCurrentLineFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetHighlightCurrentLineFunction, err = viewObject.InvokerNew("set_highlight_current_line")
	})
	return err
}

// SetHighlightCurrentLine is a representation of the C type gtk_source_view_set_highlight_current_line.
func (recv *View) SetHighlightCurrentLine(highlight bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetIndentOnTabFunction, err = viewObject.InvokerNew("set_indent_on_tab")
	})
	return err
}

// SetIndentOnTab is a representation of the C type gtk_source_view_set_indent_on_tab.
func (recv *View) SetIndentOnTab(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetIndentWidthFunction, err = viewObject.InvokerNew("set_indent_width")
	})
	return err
}

// SetIndentWidth is a representation of the C type gtk_source_view_set_indent_width.
func (recv *View) SetIndentWidth(width int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetInsertSpacesInsteadOfTabsFunction, err = viewObject.InvokerNew("set_insert_spaces_instead_of_tabs")
	})
	return err
}

// SetInsertSpacesInsteadOfTabs is a representation of the C type gtk_source_view_set_insert_spaces_instead_of_tabs.
func (recv *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetMarkAttributesFunction, err = viewObject.InvokerNew("set_mark_attributes")
	})
	return err
}

// SetMarkAttributes is a representation of the C type gtk_source_view_set_mark_attributes.
func (recv *View) SetMarkAttributes(category string, attributes *MarkAttributes, priority int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(category)
	inArgs[2].SetPointer(attributes.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetRightMarginPositionFunction, err = viewObject.InvokerNew("set_right_margin_position")
	})
	return err
}

// SetRightMarginPosition is a representation of the C type gtk_source_view_set_right_margin_position.
func (recv *View) SetRightMarginPosition(pos uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetShowLineMarksFunction, err = viewObject.InvokerNew("set_show_line_marks")
	})
	return err
}

// SetShowLineMarks is a representation of the C type gtk_source_view_set_show_line_marks.
func (recv *View) SetShowLineMarks(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetShowLineNumbersFunction, err = viewObject.InvokerNew("set_show_line_numbers")
	})
	return err
}

// SetShowLineNumbers is a representation of the C type gtk_source_view_set_show_line_numbers.
func (recv *View) SetShowLineNumbers(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetShowRightMarginFunction, err = viewObject.InvokerNew("set_show_right_margin")
	})
	return err
}

// SetShowRightMargin is a representation of the C type gtk_source_view_set_show_right_margin.
func (recv *View) SetShowRightMargin(show bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetSmartBackspaceFunction, err = viewObject.InvokerNew("set_smart_backspace")
	})
	return err
}

// SetSmartBackspace is a representation of the C type gtk_source_view_set_smart_backspace.
func (recv *View) SetSmartBackspace(smartBackspace bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetSmartHomeEndFunction, err = viewObject.InvokerNew("set_smart_home_end")
	})
	return err
}

// SetSmartHomeEnd is a representation of the C type gtk_source_view_set_smart_home_end.
func (recv *View) SetSmartHomeEnd(smartHomeEnd SmartHomeEndType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
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
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewSetTabWidthFunction, err = viewObject.InvokerNew("set_tab_width")
	})
	return err
}

// SetTabWidth is a representation of the C type gtk_source_view_set_tab_width.
func (recv *View) SetTabWidth(width uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(width)

	err := viewSetTabWidthFunction_Set()
	if err == nil {
		viewSetTabWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

var viewUnindentLinesFunction *gi.Function
var viewUnindentLinesFunction_Once sync.Once

func viewUnindentLinesFunction_Set() error {
	var err error
	viewUnindentLinesFunction_Once.Do(func() {
		err = viewObject_Set()
		if err != nil {
			return
		}
		viewUnindentLinesFunction, err = viewObject.InvokerNew("unindent_lines")
	})
	return err
}

// UnindentLines is a representation of the C type gtk_source_view_unindent_lines.
func (recv *View) UnindentLines(start *gtk.TextIter, end *gtk.TextIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(start.Native())
	inArgs[2].SetPointer(end.Native())

	err := viewUnindentLinesFunction_Set()
	if err == nil {
		viewUnindentLinesFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectChangeCase connects a callback to the 'change-case' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectChangeCase(handler func(instance *View, caseType ChangeCaseType)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := (ChangeCaseType)(object1.GetInt())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "change-case", marshal)
}

/*
ConnectChangeNumber connects a callback to the 'change-number' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectChangeNumber(handler func(instance *View, count int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetInt()

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "change-number", marshal)
}

/*
ConnectJoinLines connects a callback to the 'join-lines' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectJoinLines(handler func(instance *View)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "join-lines", marshal)
}

// UNSUPPORTED : C value 'line-mark-activated' : parameter 'event' of type 'Gdk.Event' not supported

/*
ConnectMoveLines connects a callback to the 'move-lines' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectMoveLines(handler func(instance *View, copy bool, count int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetBoolean()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := object2.GetInt()

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "move-lines", marshal)
}

/*
ConnectMoveToMatchingBracket connects a callback to the 'move-to-matching-bracket' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectMoveToMatchingBracket(handler func(instance *View, extendSelection bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetBoolean()

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "move-to-matching-bracket", marshal)
}

/*
ConnectMoveWords connects a callback to the 'move-words' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectMoveWords(handler func(instance *View, count int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetInt()

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "move-words", marshal)
}

/*
ConnectRedo connects a callback to the 'redo' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectRedo(handler func(instance *View)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "redo", marshal)
}

/*
ConnectShowCompletion connects a callback to the 'show-completion' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectShowCompletion(handler func(instance *View)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "show-completion", marshal)
}

/*
ConnectSmartHomeEnd connects a callback to the 'smart-home-end' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectSmartHomeEnd(handler func(instance *View, iter *gtk.TextIter, count int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gtk.TextIterNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := object2.GetInt()

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "smart-home-end", marshal)
}

/*
ConnectUndo connects a callback to the 'undo' signal of the View.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *View) ConnectUndo(handler func(instance *View)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ViewNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "undo", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *View) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// AtkImplementorIface returns the Atk.ImplementorIface interface implemented by View
func (recv *View) AtkImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromNative(recv.Native())
}

// GtkBuildable returns the Gtk.Buildable interface implemented by View
func (recv *View) GtkBuildable() *gtk.Buildable {
	return gtk.BuildableNewFromNative(recv.Native())
}

// GtkScrollable returns the Gtk.Scrollable interface implemented by View
func (recv *View) GtkScrollable() *gtk.Scrollable {
	return gtk.ScrollableNewFromNative(recv.Native())
}
