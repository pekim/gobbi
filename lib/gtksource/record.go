// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
	"unsafe"
)

var bufferClassStruct *gi.Struct
var bufferClassStruct_Once sync.Once

func bufferClassStruct_Set() error {
	var err error
	bufferClassStruct_Once.Do(func() {
		bufferClassStruct, err = gi.StructNew("GtkSource", "BufferClass")
	})
	return err
}

type BufferClass struct {
	native unsafe.Pointer
}

func BufferClassNewFromNative(native unsafe.Pointer) *BufferClass {
	return &BufferClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextBufferClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gtk.TextBufferClass'

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'undo' : for field setter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field setter : missing Type

// UNSUPPORTED : C value 'bracket_matched' : for field getter : missing Type

// UNSUPPORTED : C value 'bracket_matched' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field setter : missing Type

// BufferClassStruct creates an uninitialised BufferClass.
func BufferClassStruct() *BufferClass {
	err := bufferClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferClassNewFromNative(bufferClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferClass)
	return structGo
}
func finalizeBufferClass(obj *BufferClass) {
	bufferClassStruct.Free(obj.native)
}

var bufferPrivateStruct *gi.Struct
var bufferPrivateStruct_Once sync.Once

func bufferPrivateStruct_Set() error {
	var err error
	bufferPrivateStruct_Once.Do(func() {
		bufferPrivateStruct, err = gi.StructNew("GtkSource", "BufferPrivate")
	})
	return err
}

type BufferPrivate struct {
	native unsafe.Pointer
}

func BufferPrivateNewFromNative(native unsafe.Pointer) *BufferPrivate {
	return &BufferPrivate{native: native}
}

// BufferPrivateStruct creates an uninitialised BufferPrivate.
func BufferPrivateStruct() *BufferPrivate {
	err := bufferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BufferPrivateNewFromNative(bufferPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBufferPrivate)
	return structGo
}
func finalizeBufferPrivate(obj *BufferPrivate) {
	bufferPrivateStruct.Free(obj.native)
}

var completionClassStruct *gi.Struct
var completionClassStruct_Once sync.Once

func completionClassStruct_Set() error {
	var err error
	completionClassStruct_Once.Do(func() {
		completionClassStruct, err = gi.StructNew("GtkSource", "CompletionClass")
	})
	return err
}

type CompletionClass struct {
	native unsafe.Pointer
}

func CompletionClassNewFromNative(native unsafe.Pointer) *CompletionClass {
	return &CompletionClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'proposal_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'proposal_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'show' : for field getter : missing Type

// UNSUPPORTED : C value 'show' : for field setter : missing Type

// UNSUPPORTED : C value 'hide' : for field getter : missing Type

// UNSUPPORTED : C value 'hide' : for field setter : missing Type

// UNSUPPORTED : C value 'populate_context' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_context' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'move_page' : for field getter : missing Type

// UNSUPPORTED : C value 'move_page' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field setter : missing Type

// CompletionClassStruct creates an uninitialised CompletionClass.
func CompletionClassStruct() *CompletionClass {
	err := completionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionClassNewFromNative(completionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionClass)
	return structGo
}
func finalizeCompletionClass(obj *CompletionClass) {
	completionClassStruct.Free(obj.native)
}

var completionContextClassStruct *gi.Struct
var completionContextClassStruct_Once sync.Once

func completionContextClassStruct_Set() error {
	var err error
	completionContextClassStruct_Once.Do(func() {
		completionContextClassStruct, err = gi.StructNew("GtkSource", "CompletionContextClass")
	})
	return err
}

type CompletionContextClass struct {
	native unsafe.Pointer
}

func CompletionContextClassNewFromNative(native unsafe.Pointer) *CompletionContextClass {
	return &CompletionContextClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'cancelled' : for field getter : missing Type

// UNSUPPORTED : C value 'cancelled' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field setter : missing Type

// CompletionContextClassStruct creates an uninitialised CompletionContextClass.
func CompletionContextClassStruct() *CompletionContextClass {
	err := completionContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionContextClassNewFromNative(completionContextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionContextClass)
	return structGo
}
func finalizeCompletionContextClass(obj *CompletionContextClass) {
	completionContextClassStruct.Free(obj.native)
}

var completionContextPrivateStruct *gi.Struct
var completionContextPrivateStruct_Once sync.Once

func completionContextPrivateStruct_Set() error {
	var err error
	completionContextPrivateStruct_Once.Do(func() {
		completionContextPrivateStruct, err = gi.StructNew("GtkSource", "CompletionContextPrivate")
	})
	return err
}

type CompletionContextPrivate struct {
	native unsafe.Pointer
}

func CompletionContextPrivateNewFromNative(native unsafe.Pointer) *CompletionContextPrivate {
	return &CompletionContextPrivate{native: native}
}

// CompletionContextPrivateStruct creates an uninitialised CompletionContextPrivate.
func CompletionContextPrivateStruct() *CompletionContextPrivate {
	err := completionContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionContextPrivateNewFromNative(completionContextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionContextPrivate)
	return structGo
}
func finalizeCompletionContextPrivate(obj *CompletionContextPrivate) {
	completionContextPrivateStruct.Free(obj.native)
}

var completionInfoClassStruct *gi.Struct
var completionInfoClassStruct_Once sync.Once

func completionInfoClassStruct_Set() error {
	var err error
	completionInfoClassStruct_Once.Do(func() {
		completionInfoClassStruct, err = gi.StructNew("GtkSource", "CompletionInfoClass")
	})
	return err
}

type CompletionInfoClass struct {
	native unsafe.Pointer
}

func CompletionInfoClassNewFromNative(native unsafe.Pointer) *CompletionInfoClass {
	return &CompletionInfoClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.WindowClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gtk.WindowClass'

// UNSUPPORTED : C value 'before_show' : for field getter : missing Type

// UNSUPPORTED : C value 'before_show' : for field setter : missing Type

// CompletionInfoClassStruct creates an uninitialised CompletionInfoClass.
func CompletionInfoClassStruct() *CompletionInfoClass {
	err := completionInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionInfoClassNewFromNative(completionInfoClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionInfoClass)
	return structGo
}
func finalizeCompletionInfoClass(obj *CompletionInfoClass) {
	completionInfoClassStruct.Free(obj.native)
}

var completionInfoPrivateStruct *gi.Struct
var completionInfoPrivateStruct_Once sync.Once

func completionInfoPrivateStruct_Set() error {
	var err error
	completionInfoPrivateStruct_Once.Do(func() {
		completionInfoPrivateStruct, err = gi.StructNew("GtkSource", "CompletionInfoPrivate")
	})
	return err
}

type CompletionInfoPrivate struct {
	native unsafe.Pointer
}

func CompletionInfoPrivateNewFromNative(native unsafe.Pointer) *CompletionInfoPrivate {
	return &CompletionInfoPrivate{native: native}
}

// CompletionInfoPrivateStruct creates an uninitialised CompletionInfoPrivate.
func CompletionInfoPrivateStruct() *CompletionInfoPrivate {
	err := completionInfoPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionInfoPrivateNewFromNative(completionInfoPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionInfoPrivate)
	return structGo
}
func finalizeCompletionInfoPrivate(obj *CompletionInfoPrivate) {
	completionInfoPrivateStruct.Free(obj.native)
}

var completionItemClassStruct *gi.Struct
var completionItemClassStruct_Once sync.Once

func completionItemClassStruct_Set() error {
	var err error
	completionItemClassStruct_Once.Do(func() {
		completionItemClassStruct, err = gi.StructNew("GtkSource", "CompletionItemClass")
	})
	return err
}

type CompletionItemClass struct {
	native unsafe.Pointer
}

func CompletionItemClassNewFromNative(native unsafe.Pointer) *CompletionItemClass {
	return &CompletionItemClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// CompletionItemClassStruct creates an uninitialised CompletionItemClass.
func CompletionItemClassStruct() *CompletionItemClass {
	err := completionItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionItemClassNewFromNative(completionItemClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionItemClass)
	return structGo
}
func finalizeCompletionItemClass(obj *CompletionItemClass) {
	completionItemClassStruct.Free(obj.native)
}

var completionItemPrivateStruct *gi.Struct
var completionItemPrivateStruct_Once sync.Once

func completionItemPrivateStruct_Set() error {
	var err error
	completionItemPrivateStruct_Once.Do(func() {
		completionItemPrivateStruct, err = gi.StructNew("GtkSource", "CompletionItemPrivate")
	})
	return err
}

type CompletionItemPrivate struct {
	native unsafe.Pointer
}

func CompletionItemPrivateNewFromNative(native unsafe.Pointer) *CompletionItemPrivate {
	return &CompletionItemPrivate{native: native}
}

// CompletionItemPrivateStruct creates an uninitialised CompletionItemPrivate.
func CompletionItemPrivateStruct() *CompletionItemPrivate {
	err := completionItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionItemPrivateNewFromNative(completionItemPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionItemPrivate)
	return structGo
}
func finalizeCompletionItemPrivate(obj *CompletionItemPrivate) {
	completionItemPrivateStruct.Free(obj.native)
}

var completionPrivateStruct *gi.Struct
var completionPrivateStruct_Once sync.Once

func completionPrivateStruct_Set() error {
	var err error
	completionPrivateStruct_Once.Do(func() {
		completionPrivateStruct, err = gi.StructNew("GtkSource", "CompletionPrivate")
	})
	return err
}

type CompletionPrivate struct {
	native unsafe.Pointer
}

func CompletionPrivateNewFromNative(native unsafe.Pointer) *CompletionPrivate {
	return &CompletionPrivate{native: native}
}

// CompletionPrivateStruct creates an uninitialised CompletionPrivate.
func CompletionPrivateStruct() *CompletionPrivate {
	err := completionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionPrivateNewFromNative(completionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionPrivate)
	return structGo
}
func finalizeCompletionPrivate(obj *CompletionPrivate) {
	completionPrivateStruct.Free(obj.native)
}

var completionProposalIfaceStruct *gi.Struct
var completionProposalIfaceStruct_Once sync.Once

func completionProposalIfaceStruct_Set() error {
	var err error
	completionProposalIfaceStruct_Once.Do(func() {
		completionProposalIfaceStruct, err = gi.StructNew("GtkSource", "CompletionProposalIface")
	})
	return err
}

type CompletionProposalIface struct {
	native unsafe.Pointer
}

func CompletionProposalIfaceNewFromNative(native unsafe.Pointer) *CompletionProposalIface {
	return &CompletionProposalIface{native: native}
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_label' : for field getter : missing Type

// UNSUPPORTED : C value 'get_label' : for field setter : missing Type

// UNSUPPORTED : C value 'get_markup' : for field getter : missing Type

// UNSUPPORTED : C value 'get_markup' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field setter : missing Type

// UNSUPPORTED : C value 'get_info' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info' : for field setter : missing Type

// UNSUPPORTED : C value 'hash' : for field getter : missing Type

// UNSUPPORTED : C value 'hash' : for field setter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// CompletionProposalIfaceStruct creates an uninitialised CompletionProposalIface.
func CompletionProposalIfaceStruct() *CompletionProposalIface {
	err := completionProposalIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionProposalIfaceNewFromNative(completionProposalIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionProposalIface)
	return structGo
}
func finalizeCompletionProposalIface(obj *CompletionProposalIface) {
	completionProposalIfaceStruct.Free(obj.native)
}

var completionProviderIfaceStruct *gi.Struct
var completionProviderIfaceStruct_Once sync.Once

func completionProviderIfaceStruct_Set() error {
	var err error
	completionProviderIfaceStruct_Once.Do(func() {
		completionProviderIfaceStruct, err = gi.StructNew("GtkSource", "CompletionProviderIface")
	})
	return err
}

type CompletionProviderIface struct {
	native unsafe.Pointer
}

func CompletionProviderIfaceNewFromNative(native unsafe.Pointer) *CompletionProviderIface {
	return &CompletionProviderIface{native: native}
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field setter : missing Type

// UNSUPPORTED : C value 'populate' : for field getter : missing Type

// UNSUPPORTED : C value 'populate' : for field setter : missing Type

// UNSUPPORTED : C value 'match' : for field getter : missing Type

// UNSUPPORTED : C value 'match' : for field setter : missing Type

// UNSUPPORTED : C value 'get_activation' : for field getter : missing Type

// UNSUPPORTED : C value 'get_activation' : for field setter : missing Type

// UNSUPPORTED : C value 'get_info_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info_widget' : for field setter : missing Type

// UNSUPPORTED : C value 'update_info' : for field getter : missing Type

// UNSUPPORTED : C value 'update_info' : for field setter : missing Type

// UNSUPPORTED : C value 'get_start_iter' : for field getter : missing Type

// UNSUPPORTED : C value 'get_start_iter' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field setter : missing Type

// UNSUPPORTED : C value 'get_interactive_delay' : for field getter : missing Type

// UNSUPPORTED : C value 'get_interactive_delay' : for field setter : missing Type

// UNSUPPORTED : C value 'get_priority' : for field getter : missing Type

// UNSUPPORTED : C value 'get_priority' : for field setter : missing Type

// CompletionProviderIfaceStruct creates an uninitialised CompletionProviderIface.
func CompletionProviderIfaceStruct() *CompletionProviderIface {
	err := completionProviderIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionProviderIfaceNewFromNative(completionProviderIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionProviderIface)
	return structGo
}
func finalizeCompletionProviderIface(obj *CompletionProviderIface) {
	completionProviderIfaceStruct.Free(obj.native)
}

var completionWordsClassStruct *gi.Struct
var completionWordsClassStruct_Once sync.Once

func completionWordsClassStruct_Set() error {
	var err error
	completionWordsClassStruct_Once.Do(func() {
		completionWordsClassStruct, err = gi.StructNew("GtkSource", "CompletionWordsClass")
	})
	return err
}

type CompletionWordsClass struct {
	native unsafe.Pointer
}

func CompletionWordsClassNewFromNative(native unsafe.Pointer) *CompletionWordsClass {
	return &CompletionWordsClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// CompletionWordsClassStruct creates an uninitialised CompletionWordsClass.
func CompletionWordsClassStruct() *CompletionWordsClass {
	err := completionWordsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionWordsClassNewFromNative(completionWordsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionWordsClass)
	return structGo
}
func finalizeCompletionWordsClass(obj *CompletionWordsClass) {
	completionWordsClassStruct.Free(obj.native)
}

var completionWordsPrivateStruct *gi.Struct
var completionWordsPrivateStruct_Once sync.Once

func completionWordsPrivateStruct_Set() error {
	var err error
	completionWordsPrivateStruct_Once.Do(func() {
		completionWordsPrivateStruct, err = gi.StructNew("GtkSource", "CompletionWordsPrivate")
	})
	return err
}

type CompletionWordsPrivate struct {
	native unsafe.Pointer
}

func CompletionWordsPrivateNewFromNative(native unsafe.Pointer) *CompletionWordsPrivate {
	return &CompletionWordsPrivate{native: native}
}

// CompletionWordsPrivateStruct creates an uninitialised CompletionWordsPrivate.
func CompletionWordsPrivateStruct() *CompletionWordsPrivate {
	err := completionWordsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CompletionWordsPrivateNewFromNative(completionWordsPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCompletionWordsPrivate)
	return structGo
}
func finalizeCompletionWordsPrivate(obj *CompletionWordsPrivate) {
	completionWordsPrivateStruct.Free(obj.native)
}

var encodingStruct *gi.Struct
var encodingStruct_Once sync.Once

func encodingStruct_Set() error {
	var err error
	encodingStruct_Once.Do(func() {
		encodingStruct, err = gi.StructNew("GtkSource", "Encoding")
	})
	return err
}

type Encoding struct {
	native unsafe.Pointer
}

func EncodingNewFromNative(native unsafe.Pointer) *Encoding {
	return &Encoding{native: native}
}

var encodingCopyFunction *gi.Function
var encodingCopyFunction_Once sync.Once

func encodingCopyFunction_Set() error {
	var err error
	encodingCopyFunction_Once.Do(func() {
		err = encodingStruct_Set()
		if err != nil {
			return
		}
		encodingCopyFunction, err = encodingStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_source_encoding_copy.
func (recv *Encoding) Copy() *Encoding {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := encodingCopyFunction_Set()
	if err == nil {
		ret = encodingCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := EncodingNewFromNative(ret.Pointer())

	return retGo
}

var encodingFreeFunction *gi.Function
var encodingFreeFunction_Once sync.Once

func encodingFreeFunction_Set() error {
	var err error
	encodingFreeFunction_Once.Do(func() {
		err = encodingStruct_Set()
		if err != nil {
			return
		}
		encodingFreeFunction, err = encodingStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_source_encoding_free.
func (recv *Encoding) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := encodingFreeFunction_Set()
	if err == nil {
		encodingFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var encodingGetCharsetFunction *gi.Function
var encodingGetCharsetFunction_Once sync.Once

func encodingGetCharsetFunction_Set() error {
	var err error
	encodingGetCharsetFunction_Once.Do(func() {
		err = encodingStruct_Set()
		if err != nil {
			return
		}
		encodingGetCharsetFunction, err = encodingStruct.InvokerNew("get_charset")
	})
	return err
}

// GetCharset is a representation of the C type gtk_source_encoding_get_charset.
func (recv *Encoding) GetCharset() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := encodingGetCharsetFunction_Set()
	if err == nil {
		ret = encodingGetCharsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var encodingGetNameFunction *gi.Function
var encodingGetNameFunction_Once sync.Once

func encodingGetNameFunction_Set() error {
	var err error
	encodingGetNameFunction_Once.Do(func() {
		err = encodingStruct_Set()
		if err != nil {
			return
		}
		encodingGetNameFunction, err = encodingStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_encoding_get_name.
func (recv *Encoding) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := encodingGetNameFunction_Set()
	if err == nil {
		ret = encodingGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var encodingToStringFunction *gi.Function
var encodingToStringFunction_Once sync.Once

func encodingToStringFunction_Set() error {
	var err error
	encodingToStringFunction_Once.Do(func() {
		err = encodingStruct_Set()
		if err != nil {
			return
		}
		encodingToStringFunction, err = encodingStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_source_encoding_to_string.
func (recv *Encoding) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := encodingToStringFunction_Set()
	if err == nil {
		ret = encodingToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// EncodingStruct creates an uninitialised Encoding.
func EncodingStruct() *Encoding {
	err := encodingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EncodingNewFromNative(encodingStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEncoding)
	return structGo
}
func finalizeEncoding(obj *Encoding) {
	encodingStruct.Free(obj.native)
}

var fileClassStruct *gi.Struct
var fileClassStruct_Once sync.Once

func fileClassStruct_Set() error {
	var err error
	fileClassStruct_Once.Do(func() {
		fileClassStruct, err = gi.StructNew("GtkSource", "FileClass")
	})
	return err
}

type FileClass struct {
	native unsafe.Pointer
}

func FileClassNewFromNative(native unsafe.Pointer) *FileClass {
	return &FileClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// FileClassStruct creates an uninitialised FileClass.
func FileClassStruct() *FileClass {
	err := fileClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileClassNewFromNative(fileClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileClass)
	return structGo
}
func finalizeFileClass(obj *FileClass) {
	fileClassStruct.Free(obj.native)
}

var fileLoaderClassStruct *gi.Struct
var fileLoaderClassStruct_Once sync.Once

func fileLoaderClassStruct_Set() error {
	var err error
	fileLoaderClassStruct_Once.Do(func() {
		fileLoaderClassStruct, err = gi.StructNew("GtkSource", "FileLoaderClass")
	})
	return err
}

type FileLoaderClass struct {
	native unsafe.Pointer
}

func FileLoaderClassNewFromNative(native unsafe.Pointer) *FileLoaderClass {
	return &FileLoaderClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// FileLoaderClassStruct creates an uninitialised FileLoaderClass.
func FileLoaderClassStruct() *FileLoaderClass {
	err := fileLoaderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileLoaderClassNewFromNative(fileLoaderClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileLoaderClass)
	return structGo
}
func finalizeFileLoaderClass(obj *FileLoaderClass) {
	fileLoaderClassStruct.Free(obj.native)
}

var fileLoaderPrivateStruct *gi.Struct
var fileLoaderPrivateStruct_Once sync.Once

func fileLoaderPrivateStruct_Set() error {
	var err error
	fileLoaderPrivateStruct_Once.Do(func() {
		fileLoaderPrivateStruct, err = gi.StructNew("GtkSource", "FileLoaderPrivate")
	})
	return err
}

type FileLoaderPrivate struct {
	native unsafe.Pointer
}

func FileLoaderPrivateNewFromNative(native unsafe.Pointer) *FileLoaderPrivate {
	return &FileLoaderPrivate{native: native}
}

// FileLoaderPrivateStruct creates an uninitialised FileLoaderPrivate.
func FileLoaderPrivateStruct() *FileLoaderPrivate {
	err := fileLoaderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileLoaderPrivateNewFromNative(fileLoaderPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileLoaderPrivate)
	return structGo
}
func finalizeFileLoaderPrivate(obj *FileLoaderPrivate) {
	fileLoaderPrivateStruct.Free(obj.native)
}

var filePrivateStruct *gi.Struct
var filePrivateStruct_Once sync.Once

func filePrivateStruct_Set() error {
	var err error
	filePrivateStruct_Once.Do(func() {
		filePrivateStruct, err = gi.StructNew("GtkSource", "FilePrivate")
	})
	return err
}

type FilePrivate struct {
	native unsafe.Pointer
}

func FilePrivateNewFromNative(native unsafe.Pointer) *FilePrivate {
	return &FilePrivate{native: native}
}

// FilePrivateStruct creates an uninitialised FilePrivate.
func FilePrivateStruct() *FilePrivate {
	err := filePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FilePrivateNewFromNative(filePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFilePrivate)
	return structGo
}
func finalizeFilePrivate(obj *FilePrivate) {
	filePrivateStruct.Free(obj.native)
}

var fileSaverClassStruct *gi.Struct
var fileSaverClassStruct_Once sync.Once

func fileSaverClassStruct_Set() error {
	var err error
	fileSaverClassStruct_Once.Do(func() {
		fileSaverClassStruct, err = gi.StructNew("GtkSource", "FileSaverClass")
	})
	return err
}

type FileSaverClass struct {
	native unsafe.Pointer
}

func FileSaverClassNewFromNative(native unsafe.Pointer) *FileSaverClass {
	return &FileSaverClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// FileSaverClassStruct creates an uninitialised FileSaverClass.
func FileSaverClassStruct() *FileSaverClass {
	err := fileSaverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileSaverClassNewFromNative(fileSaverClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileSaverClass)
	return structGo
}
func finalizeFileSaverClass(obj *FileSaverClass) {
	fileSaverClassStruct.Free(obj.native)
}

var fileSaverPrivateStruct *gi.Struct
var fileSaverPrivateStruct_Once sync.Once

func fileSaverPrivateStruct_Set() error {
	var err error
	fileSaverPrivateStruct_Once.Do(func() {
		fileSaverPrivateStruct, err = gi.StructNew("GtkSource", "FileSaverPrivate")
	})
	return err
}

type FileSaverPrivate struct {
	native unsafe.Pointer
}

func FileSaverPrivateNewFromNative(native unsafe.Pointer) *FileSaverPrivate {
	return &FileSaverPrivate{native: native}
}

// FileSaverPrivateStruct creates an uninitialised FileSaverPrivate.
func FileSaverPrivateStruct() *FileSaverPrivate {
	err := fileSaverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileSaverPrivateNewFromNative(fileSaverPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileSaverPrivate)
	return structGo
}
func finalizeFileSaverPrivate(obj *FileSaverPrivate) {
	fileSaverPrivateStruct.Free(obj.native)
}

var gutterClassStruct *gi.Struct
var gutterClassStruct_Once sync.Once

func gutterClassStruct_Set() error {
	var err error
	gutterClassStruct_Once.Do(func() {
		gutterClassStruct, err = gi.StructNew("GtkSource", "GutterClass")
	})
	return err
}

type GutterClass struct {
	native unsafe.Pointer
}

func GutterClassNewFromNative(native unsafe.Pointer) *GutterClass {
	return &GutterClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// GutterClassStruct creates an uninitialised GutterClass.
func GutterClassStruct() *GutterClass {
	err := gutterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterClassNewFromNative(gutterClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterClass)
	return structGo
}
func finalizeGutterClass(obj *GutterClass) {
	gutterClassStruct.Free(obj.native)
}

var gutterPrivateStruct *gi.Struct
var gutterPrivateStruct_Once sync.Once

func gutterPrivateStruct_Set() error {
	var err error
	gutterPrivateStruct_Once.Do(func() {
		gutterPrivateStruct, err = gi.StructNew("GtkSource", "GutterPrivate")
	})
	return err
}

type GutterPrivate struct {
	native unsafe.Pointer
}

func GutterPrivateNewFromNative(native unsafe.Pointer) *GutterPrivate {
	return &GutterPrivate{native: native}
}

// GutterPrivateStruct creates an uninitialised GutterPrivate.
func GutterPrivateStruct() *GutterPrivate {
	err := gutterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterPrivateNewFromNative(gutterPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterPrivate)
	return structGo
}
func finalizeGutterPrivate(obj *GutterPrivate) {
	gutterPrivateStruct.Free(obj.native)
}

var gutterRendererClassStruct *gi.Struct
var gutterRendererClassStruct_Once sync.Once

func gutterRendererClassStruct_Set() error {
	var err error
	gutterRendererClassStruct_Once.Do(func() {
		gutterRendererClassStruct, err = gi.StructNew("GtkSource", "GutterRendererClass")
	})
	return err
}

type GutterRendererClass struct {
	native unsafe.Pointer
}

func GutterRendererClassNewFromNative(native unsafe.Pointer) *GutterRendererClass {
	return &GutterRendererClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'begin' : for field getter : missing Type

// UNSUPPORTED : C value 'begin' : for field setter : missing Type

// UNSUPPORTED : C value 'draw' : for field getter : missing Type

// UNSUPPORTED : C value 'draw' : for field setter : missing Type

// UNSUPPORTED : C value 'end' : for field getter : missing Type

// UNSUPPORTED : C value 'end' : for field setter : missing Type

// UNSUPPORTED : C value 'change_view' : for field getter : missing Type

// UNSUPPORTED : C value 'change_view' : for field setter : missing Type

// UNSUPPORTED : C value 'change_buffer' : for field getter : missing Type

// UNSUPPORTED : C value 'change_buffer' : for field setter : missing Type

// UNSUPPORTED : C value 'query_activatable' : for field getter : missing Type

// UNSUPPORTED : C value 'query_activatable' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'queue_draw' : for field getter : missing Type

// UNSUPPORTED : C value 'queue_draw' : for field setter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field getter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field setter : missing Type

// UNSUPPORTED : C value 'query_data' : for field getter : missing Type

// UNSUPPORTED : C value 'query_data' : for field setter : missing Type

// GutterRendererClassStruct creates an uninitialised GutterRendererClass.
func GutterRendererClassStruct() *GutterRendererClass {
	err := gutterRendererClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererClassNewFromNative(gutterRendererClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererClass)
	return structGo
}
func finalizeGutterRendererClass(obj *GutterRendererClass) {
	gutterRendererClassStruct.Free(obj.native)
}

var gutterRendererPixbufClassStruct *gi.Struct
var gutterRendererPixbufClassStruct_Once sync.Once

func gutterRendererPixbufClassStruct_Set() error {
	var err error
	gutterRendererPixbufClassStruct_Once.Do(func() {
		gutterRendererPixbufClassStruct, err = gi.StructNew("GtkSource", "GutterRendererPixbufClass")
	})
	return err
}

type GutterRendererPixbufClass struct {
	native unsafe.Pointer
}

func GutterRendererPixbufClassNewFromNative(native unsafe.Pointer) *GutterRendererPixbufClass {
	return &GutterRendererPixbufClass{native: native}
}

// GutterRendererPixbufClassStruct creates an uninitialised GutterRendererPixbufClass.
func GutterRendererPixbufClassStruct() *GutterRendererPixbufClass {
	err := gutterRendererPixbufClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererPixbufClassNewFromNative(gutterRendererPixbufClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererPixbufClass)
	return structGo
}
func finalizeGutterRendererPixbufClass(obj *GutterRendererPixbufClass) {
	gutterRendererPixbufClassStruct.Free(obj.native)
}

var gutterRendererPixbufPrivateStruct *gi.Struct
var gutterRendererPixbufPrivateStruct_Once sync.Once

func gutterRendererPixbufPrivateStruct_Set() error {
	var err error
	gutterRendererPixbufPrivateStruct_Once.Do(func() {
		gutterRendererPixbufPrivateStruct, err = gi.StructNew("GtkSource", "GutterRendererPixbufPrivate")
	})
	return err
}

type GutterRendererPixbufPrivate struct {
	native unsafe.Pointer
}

func GutterRendererPixbufPrivateNewFromNative(native unsafe.Pointer) *GutterRendererPixbufPrivate {
	return &GutterRendererPixbufPrivate{native: native}
}

// GutterRendererPixbufPrivateStruct creates an uninitialised GutterRendererPixbufPrivate.
func GutterRendererPixbufPrivateStruct() *GutterRendererPixbufPrivate {
	err := gutterRendererPixbufPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererPixbufPrivateNewFromNative(gutterRendererPixbufPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererPixbufPrivate)
	return structGo
}
func finalizeGutterRendererPixbufPrivate(obj *GutterRendererPixbufPrivate) {
	gutterRendererPixbufPrivateStruct.Free(obj.native)
}

var gutterRendererPrivateStruct *gi.Struct
var gutterRendererPrivateStruct_Once sync.Once

func gutterRendererPrivateStruct_Set() error {
	var err error
	gutterRendererPrivateStruct_Once.Do(func() {
		gutterRendererPrivateStruct, err = gi.StructNew("GtkSource", "GutterRendererPrivate")
	})
	return err
}

type GutterRendererPrivate struct {
	native unsafe.Pointer
}

func GutterRendererPrivateNewFromNative(native unsafe.Pointer) *GutterRendererPrivate {
	return &GutterRendererPrivate{native: native}
}

// GutterRendererPrivateStruct creates an uninitialised GutterRendererPrivate.
func GutterRendererPrivateStruct() *GutterRendererPrivate {
	err := gutterRendererPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererPrivateNewFromNative(gutterRendererPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererPrivate)
	return structGo
}
func finalizeGutterRendererPrivate(obj *GutterRendererPrivate) {
	gutterRendererPrivateStruct.Free(obj.native)
}

var gutterRendererTextClassStruct *gi.Struct
var gutterRendererTextClassStruct_Once sync.Once

func gutterRendererTextClassStruct_Set() error {
	var err error
	gutterRendererTextClassStruct_Once.Do(func() {
		gutterRendererTextClassStruct, err = gi.StructNew("GtkSource", "GutterRendererTextClass")
	})
	return err
}

type GutterRendererTextClass struct {
	native unsafe.Pointer
}

func GutterRendererTextClassNewFromNative(native unsafe.Pointer) *GutterRendererTextClass {
	return &GutterRendererTextClass{native: native}
}

// GutterRendererTextClassStruct creates an uninitialised GutterRendererTextClass.
func GutterRendererTextClassStruct() *GutterRendererTextClass {
	err := gutterRendererTextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererTextClassNewFromNative(gutterRendererTextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererTextClass)
	return structGo
}
func finalizeGutterRendererTextClass(obj *GutterRendererTextClass) {
	gutterRendererTextClassStruct.Free(obj.native)
}

var gutterRendererTextPrivateStruct *gi.Struct
var gutterRendererTextPrivateStruct_Once sync.Once

func gutterRendererTextPrivateStruct_Set() error {
	var err error
	gutterRendererTextPrivateStruct_Once.Do(func() {
		gutterRendererTextPrivateStruct, err = gi.StructNew("GtkSource", "GutterRendererTextPrivate")
	})
	return err
}

type GutterRendererTextPrivate struct {
	native unsafe.Pointer
}

func GutterRendererTextPrivateNewFromNative(native unsafe.Pointer) *GutterRendererTextPrivate {
	return &GutterRendererTextPrivate{native: native}
}

// GutterRendererTextPrivateStruct creates an uninitialised GutterRendererTextPrivate.
func GutterRendererTextPrivateStruct() *GutterRendererTextPrivate {
	err := gutterRendererTextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GutterRendererTextPrivateNewFromNative(gutterRendererTextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGutterRendererTextPrivate)
	return structGo
}
func finalizeGutterRendererTextPrivate(obj *GutterRendererTextPrivate) {
	gutterRendererTextPrivateStruct.Free(obj.native)
}

var languageClassStruct *gi.Struct
var languageClassStruct_Once sync.Once

func languageClassStruct_Set() error {
	var err error
	languageClassStruct_Once.Do(func() {
		languageClassStruct, err = gi.StructNew("GtkSource", "LanguageClass")
	})
	return err
}

type LanguageClass struct {
	native unsafe.Pointer
}

func LanguageClassNewFromNative(native unsafe.Pointer) *LanguageClass {
	return &LanguageClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// LanguageClassStruct creates an uninitialised LanguageClass.
func LanguageClassStruct() *LanguageClass {
	err := languageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LanguageClassNewFromNative(languageClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLanguageClass)
	return structGo
}
func finalizeLanguageClass(obj *LanguageClass) {
	languageClassStruct.Free(obj.native)
}

var languageManagerClassStruct *gi.Struct
var languageManagerClassStruct_Once sync.Once

func languageManagerClassStruct_Set() error {
	var err error
	languageManagerClassStruct_Once.Do(func() {
		languageManagerClassStruct, err = gi.StructNew("GtkSource", "LanguageManagerClass")
	})
	return err
}

type LanguageManagerClass struct {
	native unsafe.Pointer
}

func LanguageManagerClassNewFromNative(native unsafe.Pointer) *LanguageManagerClass {
	return &LanguageManagerClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field setter : missing Type

// LanguageManagerClassStruct creates an uninitialised LanguageManagerClass.
func LanguageManagerClassStruct() *LanguageManagerClass {
	err := languageManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LanguageManagerClassNewFromNative(languageManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLanguageManagerClass)
	return structGo
}
func finalizeLanguageManagerClass(obj *LanguageManagerClass) {
	languageManagerClassStruct.Free(obj.native)
}

var languageManagerPrivateStruct *gi.Struct
var languageManagerPrivateStruct_Once sync.Once

func languageManagerPrivateStruct_Set() error {
	var err error
	languageManagerPrivateStruct_Once.Do(func() {
		languageManagerPrivateStruct, err = gi.StructNew("GtkSource", "LanguageManagerPrivate")
	})
	return err
}

type LanguageManagerPrivate struct {
	native unsafe.Pointer
}

func LanguageManagerPrivateNewFromNative(native unsafe.Pointer) *LanguageManagerPrivate {
	return &LanguageManagerPrivate{native: native}
}

// LanguageManagerPrivateStruct creates an uninitialised LanguageManagerPrivate.
func LanguageManagerPrivateStruct() *LanguageManagerPrivate {
	err := languageManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LanguageManagerPrivateNewFromNative(languageManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLanguageManagerPrivate)
	return structGo
}
func finalizeLanguageManagerPrivate(obj *LanguageManagerPrivate) {
	languageManagerPrivateStruct.Free(obj.native)
}

var languagePrivateStruct *gi.Struct
var languagePrivateStruct_Once sync.Once

func languagePrivateStruct_Set() error {
	var err error
	languagePrivateStruct_Once.Do(func() {
		languagePrivateStruct, err = gi.StructNew("GtkSource", "LanguagePrivate")
	})
	return err
}

type LanguagePrivate struct {
	native unsafe.Pointer
}

func LanguagePrivateNewFromNative(native unsafe.Pointer) *LanguagePrivate {
	return &LanguagePrivate{native: native}
}

// LanguagePrivateStruct creates an uninitialised LanguagePrivate.
func LanguagePrivateStruct() *LanguagePrivate {
	err := languagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LanguagePrivateNewFromNative(languagePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLanguagePrivate)
	return structGo
}
func finalizeLanguagePrivate(obj *LanguagePrivate) {
	languagePrivateStruct.Free(obj.native)
}

var mapClassStruct *gi.Struct
var mapClassStruct_Once sync.Once

func mapClassStruct_Set() error {
	var err error
	mapClassStruct_Once.Do(func() {
		mapClassStruct, err = gi.StructNew("GtkSource", "MapClass")
	})
	return err
}

type MapClass struct {
	native unsafe.Pointer
}

func MapClassNewFromNative(native unsafe.Pointer) *MapClass {
	return &MapClass{native: native}
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *MapClass) FieldParentClass() *ViewClass {
	argValue := gi.StructFieldGet(mapClassStruct, recv.native, "parent_class")
	value := ViewClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MapClass) SetFieldParentClass(value *ViewClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(mapClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// MapClassStruct creates an uninitialised MapClass.
func MapClassStruct() *MapClass {
	err := mapClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MapClassNewFromNative(mapClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMapClass)
	return structGo
}
func finalizeMapClass(obj *MapClass) {
	mapClassStruct.Free(obj.native)
}

var markAttributesClassStruct *gi.Struct
var markAttributesClassStruct_Once sync.Once

func markAttributesClassStruct_Set() error {
	var err error
	markAttributesClassStruct_Once.Do(func() {
		markAttributesClassStruct, err = gi.StructNew("GtkSource", "MarkAttributesClass")
	})
	return err
}

type MarkAttributesClass struct {
	native unsafe.Pointer
}

func MarkAttributesClassNewFromNative(native unsafe.Pointer) *MarkAttributesClass {
	return &MarkAttributesClass{native: native}
}

// MarkAttributesClassStruct creates an uninitialised MarkAttributesClass.
func MarkAttributesClassStruct() *MarkAttributesClass {
	err := markAttributesClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MarkAttributesClassNewFromNative(markAttributesClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMarkAttributesClass)
	return structGo
}
func finalizeMarkAttributesClass(obj *MarkAttributesClass) {
	markAttributesClassStruct.Free(obj.native)
}

var markAttributesPrivateStruct *gi.Struct
var markAttributesPrivateStruct_Once sync.Once

func markAttributesPrivateStruct_Set() error {
	var err error
	markAttributesPrivateStruct_Once.Do(func() {
		markAttributesPrivateStruct, err = gi.StructNew("GtkSource", "MarkAttributesPrivate")
	})
	return err
}

type MarkAttributesPrivate struct {
	native unsafe.Pointer
}

func MarkAttributesPrivateNewFromNative(native unsafe.Pointer) *MarkAttributesPrivate {
	return &MarkAttributesPrivate{native: native}
}

// MarkAttributesPrivateStruct creates an uninitialised MarkAttributesPrivate.
func MarkAttributesPrivateStruct() *MarkAttributesPrivate {
	err := markAttributesPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MarkAttributesPrivateNewFromNative(markAttributesPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMarkAttributesPrivate)
	return structGo
}
func finalizeMarkAttributesPrivate(obj *MarkAttributesPrivate) {
	markAttributesPrivateStruct.Free(obj.native)
}

var markClassStruct *gi.Struct
var markClassStruct_Once sync.Once

func markClassStruct_Set() error {
	var err error
	markClassStruct_Once.Do(func() {
		markClassStruct, err = gi.StructNew("GtkSource", "MarkClass")
	})
	return err
}

type MarkClass struct {
	native unsafe.Pointer
}

func MarkClassNewFromNative(native unsafe.Pointer) *MarkClass {
	return &MarkClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextMarkClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gtk.TextMarkClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// MarkClassStruct creates an uninitialised MarkClass.
func MarkClassStruct() *MarkClass {
	err := markClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MarkClassNewFromNative(markClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMarkClass)
	return structGo
}
func finalizeMarkClass(obj *MarkClass) {
	markClassStruct.Free(obj.native)
}

var markPrivateStruct *gi.Struct
var markPrivateStruct_Once sync.Once

func markPrivateStruct_Set() error {
	var err error
	markPrivateStruct_Once.Do(func() {
		markPrivateStruct, err = gi.StructNew("GtkSource", "MarkPrivate")
	})
	return err
}

type MarkPrivate struct {
	native unsafe.Pointer
}

func MarkPrivateNewFromNative(native unsafe.Pointer) *MarkPrivate {
	return &MarkPrivate{native: native}
}

// MarkPrivateStruct creates an uninitialised MarkPrivate.
func MarkPrivateStruct() *MarkPrivate {
	err := markPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MarkPrivateNewFromNative(markPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMarkPrivate)
	return structGo
}
func finalizeMarkPrivate(obj *MarkPrivate) {
	markPrivateStruct.Free(obj.native)
}

var printCompositorClassStruct *gi.Struct
var printCompositorClassStruct_Once sync.Once

func printCompositorClassStruct_Set() error {
	var err error
	printCompositorClassStruct_Once.Do(func() {
		printCompositorClassStruct, err = gi.StructNew("GtkSource", "PrintCompositorClass")
	})
	return err
}

type PrintCompositorClass struct {
	native unsafe.Pointer
}

func PrintCompositorClassNewFromNative(native unsafe.Pointer) *PrintCompositorClass {
	return &PrintCompositorClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// PrintCompositorClassStruct creates an uninitialised PrintCompositorClass.
func PrintCompositorClassStruct() *PrintCompositorClass {
	err := printCompositorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintCompositorClassNewFromNative(printCompositorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintCompositorClass)
	return structGo
}
func finalizePrintCompositorClass(obj *PrintCompositorClass) {
	printCompositorClassStruct.Free(obj.native)
}

var printCompositorPrivateStruct *gi.Struct
var printCompositorPrivateStruct_Once sync.Once

func printCompositorPrivateStruct_Set() error {
	var err error
	printCompositorPrivateStruct_Once.Do(func() {
		printCompositorPrivateStruct, err = gi.StructNew("GtkSource", "PrintCompositorPrivate")
	})
	return err
}

type PrintCompositorPrivate struct {
	native unsafe.Pointer
}

func PrintCompositorPrivateNewFromNative(native unsafe.Pointer) *PrintCompositorPrivate {
	return &PrintCompositorPrivate{native: native}
}

// PrintCompositorPrivateStruct creates an uninitialised PrintCompositorPrivate.
func PrintCompositorPrivateStruct() *PrintCompositorPrivate {
	err := printCompositorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintCompositorPrivateNewFromNative(printCompositorPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintCompositorPrivate)
	return structGo
}
func finalizePrintCompositorPrivate(obj *PrintCompositorPrivate) {
	printCompositorPrivateStruct.Free(obj.native)
}

var regionClassStruct *gi.Struct
var regionClassStruct_Once sync.Once

func regionClassStruct_Set() error {
	var err error
	regionClassStruct_Once.Do(func() {
		regionClassStruct, err = gi.StructNew("GtkSource", "RegionClass")
	})
	return err
}

type RegionClass struct {
	native unsafe.Pointer
}

func RegionClassNewFromNative(native unsafe.Pointer) *RegionClass {
	return &RegionClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// RegionClassStruct creates an uninitialised RegionClass.
func RegionClassStruct() *RegionClass {
	err := regionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RegionClassNewFromNative(regionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRegionClass)
	return structGo
}
func finalizeRegionClass(obj *RegionClass) {
	regionClassStruct.Free(obj.native)
}

var regionIterStruct *gi.Struct
var regionIterStruct_Once sync.Once

func regionIterStruct_Set() error {
	var err error
	regionIterStruct_Once.Do(func() {
		regionIterStruct, err = gi.StructNew("GtkSource", "RegionIter")
	})
	return err
}

type RegionIter struct {
	native unsafe.Pointer
}

func RegionIterNewFromNative(native unsafe.Pointer) *RegionIter {
	return &RegionIter{native: native}
}

// UNSUPPORTED : C value 'gtk_source_region_iter_get_subregion' : parameter 'start' of type 'Gtk.TextIter' not supported

var regionIterIsEndFunction *gi.Function
var regionIterIsEndFunction_Once sync.Once

func regionIterIsEndFunction_Set() error {
	var err error
	regionIterIsEndFunction_Once.Do(func() {
		err = regionIterStruct_Set()
		if err != nil {
			return
		}
		regionIterIsEndFunction, err = regionIterStruct.InvokerNew("is_end")
	})
	return err
}

// IsEnd is a representation of the C type gtk_source_region_iter_is_end.
func (recv *RegionIter) IsEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regionIterIsEndFunction_Set()
	if err == nil {
		ret = regionIterIsEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var regionIterNextFunction *gi.Function
var regionIterNextFunction_Once sync.Once

func regionIterNextFunction_Set() error {
	var err error
	regionIterNextFunction_Once.Do(func() {
		err = regionIterStruct_Set()
		if err != nil {
			return
		}
		regionIterNextFunction, err = regionIterStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type gtk_source_region_iter_next.
func (recv *RegionIter) Next() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regionIterNextFunction_Set()
	if err == nil {
		ret = regionIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// RegionIterStruct creates an uninitialised RegionIter.
func RegionIterStruct() *RegionIter {
	err := regionIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RegionIterNewFromNative(regionIterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRegionIter)
	return structGo
}
func finalizeRegionIter(obj *RegionIter) {
	regionIterStruct.Free(obj.native)
}

var searchContextClassStruct *gi.Struct
var searchContextClassStruct_Once sync.Once

func searchContextClassStruct_Set() error {
	var err error
	searchContextClassStruct_Once.Do(func() {
		searchContextClassStruct, err = gi.StructNew("GtkSource", "SearchContextClass")
	})
	return err
}

type SearchContextClass struct {
	native unsafe.Pointer
}

func SearchContextClassNewFromNative(native unsafe.Pointer) *SearchContextClass {
	return &SearchContextClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// SearchContextClassStruct creates an uninitialised SearchContextClass.
func SearchContextClassStruct() *SearchContextClass {
	err := searchContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SearchContextClassNewFromNative(searchContextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSearchContextClass)
	return structGo
}
func finalizeSearchContextClass(obj *SearchContextClass) {
	searchContextClassStruct.Free(obj.native)
}

var searchContextPrivateStruct *gi.Struct
var searchContextPrivateStruct_Once sync.Once

func searchContextPrivateStruct_Set() error {
	var err error
	searchContextPrivateStruct_Once.Do(func() {
		searchContextPrivateStruct, err = gi.StructNew("GtkSource", "SearchContextPrivate")
	})
	return err
}

type SearchContextPrivate struct {
	native unsafe.Pointer
}

func SearchContextPrivateNewFromNative(native unsafe.Pointer) *SearchContextPrivate {
	return &SearchContextPrivate{native: native}
}

// SearchContextPrivateStruct creates an uninitialised SearchContextPrivate.
func SearchContextPrivateStruct() *SearchContextPrivate {
	err := searchContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SearchContextPrivateNewFromNative(searchContextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSearchContextPrivate)
	return structGo
}
func finalizeSearchContextPrivate(obj *SearchContextPrivate) {
	searchContextPrivateStruct.Free(obj.native)
}

var searchSettingsClassStruct *gi.Struct
var searchSettingsClassStruct_Once sync.Once

func searchSettingsClassStruct_Set() error {
	var err error
	searchSettingsClassStruct_Once.Do(func() {
		searchSettingsClassStruct, err = gi.StructNew("GtkSource", "SearchSettingsClass")
	})
	return err
}

type SearchSettingsClass struct {
	native unsafe.Pointer
}

func SearchSettingsClassNewFromNative(native unsafe.Pointer) *SearchSettingsClass {
	return &SearchSettingsClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// SearchSettingsClassStruct creates an uninitialised SearchSettingsClass.
func SearchSettingsClassStruct() *SearchSettingsClass {
	err := searchSettingsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SearchSettingsClassNewFromNative(searchSettingsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSearchSettingsClass)
	return structGo
}
func finalizeSearchSettingsClass(obj *SearchSettingsClass) {
	searchSettingsClassStruct.Free(obj.native)
}

var searchSettingsPrivateStruct *gi.Struct
var searchSettingsPrivateStruct_Once sync.Once

func searchSettingsPrivateStruct_Set() error {
	var err error
	searchSettingsPrivateStruct_Once.Do(func() {
		searchSettingsPrivateStruct, err = gi.StructNew("GtkSource", "SearchSettingsPrivate")
	})
	return err
}

type SearchSettingsPrivate struct {
	native unsafe.Pointer
}

func SearchSettingsPrivateNewFromNative(native unsafe.Pointer) *SearchSettingsPrivate {
	return &SearchSettingsPrivate{native: native}
}

// SearchSettingsPrivateStruct creates an uninitialised SearchSettingsPrivate.
func SearchSettingsPrivateStruct() *SearchSettingsPrivate {
	err := searchSettingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SearchSettingsPrivateNewFromNative(searchSettingsPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSearchSettingsPrivate)
	return structGo
}
func finalizeSearchSettingsPrivate(obj *SearchSettingsPrivate) {
	searchSettingsPrivateStruct.Free(obj.native)
}

var spaceDrawerClassStruct *gi.Struct
var spaceDrawerClassStruct_Once sync.Once

func spaceDrawerClassStruct_Set() error {
	var err error
	spaceDrawerClassStruct_Once.Do(func() {
		spaceDrawerClassStruct, err = gi.StructNew("GtkSource", "SpaceDrawerClass")
	})
	return err
}

type SpaceDrawerClass struct {
	native unsafe.Pointer
}

func SpaceDrawerClassNewFromNative(native unsafe.Pointer) *SpaceDrawerClass {
	return &SpaceDrawerClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// SpaceDrawerClassStruct creates an uninitialised SpaceDrawerClass.
func SpaceDrawerClassStruct() *SpaceDrawerClass {
	err := spaceDrawerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SpaceDrawerClassNewFromNative(spaceDrawerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSpaceDrawerClass)
	return structGo
}
func finalizeSpaceDrawerClass(obj *SpaceDrawerClass) {
	spaceDrawerClassStruct.Free(obj.native)
}

var spaceDrawerPrivateStruct *gi.Struct
var spaceDrawerPrivateStruct_Once sync.Once

func spaceDrawerPrivateStruct_Set() error {
	var err error
	spaceDrawerPrivateStruct_Once.Do(func() {
		spaceDrawerPrivateStruct, err = gi.StructNew("GtkSource", "SpaceDrawerPrivate")
	})
	return err
}

type SpaceDrawerPrivate struct {
	native unsafe.Pointer
}

func SpaceDrawerPrivateNewFromNative(native unsafe.Pointer) *SpaceDrawerPrivate {
	return &SpaceDrawerPrivate{native: native}
}

// SpaceDrawerPrivateStruct creates an uninitialised SpaceDrawerPrivate.
func SpaceDrawerPrivateStruct() *SpaceDrawerPrivate {
	err := spaceDrawerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SpaceDrawerPrivateNewFromNative(spaceDrawerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSpaceDrawerPrivate)
	return structGo
}
func finalizeSpaceDrawerPrivate(obj *SpaceDrawerPrivate) {
	spaceDrawerPrivateStruct.Free(obj.native)
}

var styleClassStruct *gi.Struct
var styleClassStruct_Once sync.Once

func styleClassStruct_Set() error {
	var err error
	styleClassStruct_Once.Do(func() {
		styleClassStruct, err = gi.StructNew("GtkSource", "StyleClass")
	})
	return err
}

type StyleClass struct {
	native unsafe.Pointer
}

func StyleClassNewFromNative(native unsafe.Pointer) *StyleClass {
	return &StyleClass{native: native}
}

// StyleClassStruct creates an uninitialised StyleClass.
func StyleClassStruct() *StyleClass {
	err := styleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleClassNewFromNative(styleClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleClass)
	return structGo
}
func finalizeStyleClass(obj *StyleClass) {
	styleClassStruct.Free(obj.native)
}

var styleSchemeChooserButtonClassStruct *gi.Struct
var styleSchemeChooserButtonClassStruct_Once sync.Once

func styleSchemeChooserButtonClassStruct_Set() error {
	var err error
	styleSchemeChooserButtonClassStruct_Once.Do(func() {
		styleSchemeChooserButtonClassStruct, err = gi.StructNew("GtkSource", "StyleSchemeChooserButtonClass")
	})
	return err
}

type StyleSchemeChooserButtonClass struct {
	native unsafe.Pointer
}

func StyleSchemeChooserButtonClassNewFromNative(native unsafe.Pointer) *StyleSchemeChooserButtonClass {
	return &StyleSchemeChooserButtonClass{native: native}
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.ButtonClass'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.ButtonClass'

// StyleSchemeChooserButtonClassStruct creates an uninitialised StyleSchemeChooserButtonClass.
func StyleSchemeChooserButtonClassStruct() *StyleSchemeChooserButtonClass {
	err := styleSchemeChooserButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeChooserButtonClassNewFromNative(styleSchemeChooserButtonClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeChooserButtonClass)
	return structGo
}
func finalizeStyleSchemeChooserButtonClass(obj *StyleSchemeChooserButtonClass) {
	styleSchemeChooserButtonClassStruct.Free(obj.native)
}

var styleSchemeChooserInterfaceStruct *gi.Struct
var styleSchemeChooserInterfaceStruct_Once sync.Once

func styleSchemeChooserInterfaceStruct_Set() error {
	var err error
	styleSchemeChooserInterfaceStruct_Once.Do(func() {
		styleSchemeChooserInterfaceStruct, err = gi.StructNew("GtkSource", "StyleSchemeChooserInterface")
	})
	return err
}

type StyleSchemeChooserInterface struct {
	native unsafe.Pointer
}

func StyleSchemeChooserInterfaceNewFromNative(native unsafe.Pointer) *StyleSchemeChooserInterface {
	return &StyleSchemeChooserInterface{native: native}
}

// UNSUPPORTED : C value 'base_interface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_interface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_style_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'get_style_scheme' : for field setter : missing Type

// UNSUPPORTED : C value 'set_style_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'set_style_scheme' : for field setter : missing Type

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// StyleSchemeChooserInterfaceStruct creates an uninitialised StyleSchemeChooserInterface.
func StyleSchemeChooserInterfaceStruct() *StyleSchemeChooserInterface {
	err := styleSchemeChooserInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeChooserInterfaceNewFromNative(styleSchemeChooserInterfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeChooserInterface)
	return structGo
}
func finalizeStyleSchemeChooserInterface(obj *StyleSchemeChooserInterface) {
	styleSchemeChooserInterfaceStruct.Free(obj.native)
}

var styleSchemeChooserWidgetClassStruct *gi.Struct
var styleSchemeChooserWidgetClassStruct_Once sync.Once

func styleSchemeChooserWidgetClassStruct_Set() error {
	var err error
	styleSchemeChooserWidgetClassStruct_Once.Do(func() {
		styleSchemeChooserWidgetClassStruct, err = gi.StructNew("GtkSource", "StyleSchemeChooserWidgetClass")
	})
	return err
}

type StyleSchemeChooserWidgetClass struct {
	native unsafe.Pointer
}

func StyleSchemeChooserWidgetClassNewFromNative(native unsafe.Pointer) *StyleSchemeChooserWidgetClass {
	return &StyleSchemeChooserWidgetClass{native: native}
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.BinClass'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'Gtk.BinClass'

// StyleSchemeChooserWidgetClassStruct creates an uninitialised StyleSchemeChooserWidgetClass.
func StyleSchemeChooserWidgetClassStruct() *StyleSchemeChooserWidgetClass {
	err := styleSchemeChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeChooserWidgetClassNewFromNative(styleSchemeChooserWidgetClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeChooserWidgetClass)
	return structGo
}
func finalizeStyleSchemeChooserWidgetClass(obj *StyleSchemeChooserWidgetClass) {
	styleSchemeChooserWidgetClassStruct.Free(obj.native)
}

var styleSchemeClassStruct *gi.Struct
var styleSchemeClassStruct_Once sync.Once

func styleSchemeClassStruct_Set() error {
	var err error
	styleSchemeClassStruct_Once.Do(func() {
		styleSchemeClassStruct, err = gi.StructNew("GtkSource", "StyleSchemeClass")
	})
	return err
}

type StyleSchemeClass struct {
	native unsafe.Pointer
}

func StyleSchemeClassNewFromNative(native unsafe.Pointer) *StyleSchemeClass {
	return &StyleSchemeClass{native: native}
}

// UNSUPPORTED : C value 'base_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'base_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// StyleSchemeClassStruct creates an uninitialised StyleSchemeClass.
func StyleSchemeClassStruct() *StyleSchemeClass {
	err := styleSchemeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeClassNewFromNative(styleSchemeClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeClass)
	return structGo
}
func finalizeStyleSchemeClass(obj *StyleSchemeClass) {
	styleSchemeClassStruct.Free(obj.native)
}

var styleSchemeManagerClassStruct *gi.Struct
var styleSchemeManagerClassStruct_Once sync.Once

func styleSchemeManagerClassStruct_Set() error {
	var err error
	styleSchemeManagerClassStruct_Once.Do(func() {
		styleSchemeManagerClassStruct, err = gi.StructNew("GtkSource", "StyleSchemeManagerClass")
	})
	return err
}

type StyleSchemeManagerClass struct {
	native unsafe.Pointer
}

func StyleSchemeManagerClassNewFromNative(native unsafe.Pointer) *StyleSchemeManagerClass {
	return &StyleSchemeManagerClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field setter : missing Type

// StyleSchemeManagerClassStruct creates an uninitialised StyleSchemeManagerClass.
func StyleSchemeManagerClassStruct() *StyleSchemeManagerClass {
	err := styleSchemeManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeManagerClassNewFromNative(styleSchemeManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeManagerClass)
	return structGo
}
func finalizeStyleSchemeManagerClass(obj *StyleSchemeManagerClass) {
	styleSchemeManagerClassStruct.Free(obj.native)
}

var styleSchemeManagerPrivateStruct *gi.Struct
var styleSchemeManagerPrivateStruct_Once sync.Once

func styleSchemeManagerPrivateStruct_Set() error {
	var err error
	styleSchemeManagerPrivateStruct_Once.Do(func() {
		styleSchemeManagerPrivateStruct, err = gi.StructNew("GtkSource", "StyleSchemeManagerPrivate")
	})
	return err
}

type StyleSchemeManagerPrivate struct {
	native unsafe.Pointer
}

func StyleSchemeManagerPrivateNewFromNative(native unsafe.Pointer) *StyleSchemeManagerPrivate {
	return &StyleSchemeManagerPrivate{native: native}
}

// StyleSchemeManagerPrivateStruct creates an uninitialised StyleSchemeManagerPrivate.
func StyleSchemeManagerPrivateStruct() *StyleSchemeManagerPrivate {
	err := styleSchemeManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemeManagerPrivateNewFromNative(styleSchemeManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemeManagerPrivate)
	return structGo
}
func finalizeStyleSchemeManagerPrivate(obj *StyleSchemeManagerPrivate) {
	styleSchemeManagerPrivateStruct.Free(obj.native)
}

var styleSchemePrivateStruct *gi.Struct
var styleSchemePrivateStruct_Once sync.Once

func styleSchemePrivateStruct_Set() error {
	var err error
	styleSchemePrivateStruct_Once.Do(func() {
		styleSchemePrivateStruct, err = gi.StructNew("GtkSource", "StyleSchemePrivate")
	})
	return err
}

type StyleSchemePrivate struct {
	native unsafe.Pointer
}

func StyleSchemePrivateNewFromNative(native unsafe.Pointer) *StyleSchemePrivate {
	return &StyleSchemePrivate{native: native}
}

// StyleSchemePrivateStruct creates an uninitialised StyleSchemePrivate.
func StyleSchemePrivateStruct() *StyleSchemePrivate {
	err := styleSchemePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StyleSchemePrivateNewFromNative(styleSchemePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStyleSchemePrivate)
	return structGo
}
func finalizeStyleSchemePrivate(obj *StyleSchemePrivate) {
	styleSchemePrivateStruct.Free(obj.native)
}

var tagClassStruct *gi.Struct
var tagClassStruct_Once sync.Once

func tagClassStruct_Set() error {
	var err error
	tagClassStruct_Once.Do(func() {
		tagClassStruct, err = gi.StructNew("GtkSource", "TagClass")
	})
	return err
}

type TagClass struct {
	native unsafe.Pointer
}

func TagClassNewFromNative(native unsafe.Pointer) *TagClass {
	return &TagClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextTagClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gtk.TextTagClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// TagClassStruct creates an uninitialised TagClass.
func TagClassStruct() *TagClass {
	err := tagClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TagClassNewFromNative(tagClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTagClass)
	return structGo
}
func finalizeTagClass(obj *TagClass) {
	tagClassStruct.Free(obj.native)
}

var undoManagerIfaceStruct *gi.Struct
var undoManagerIfaceStruct_Once sync.Once

func undoManagerIfaceStruct_Set() error {
	var err error
	undoManagerIfaceStruct_Once.Do(func() {
		undoManagerIfaceStruct, err = gi.StructNew("GtkSource", "UndoManagerIface")
	})
	return err
}

type UndoManagerIface struct {
	native unsafe.Pointer
}

func UndoManagerIfaceNewFromNative(native unsafe.Pointer) *UndoManagerIface {
	return &UndoManagerIface{native: native}
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'can_undo' : for field getter : missing Type

// UNSUPPORTED : C value 'can_undo' : for field setter : missing Type

// UNSUPPORTED : C value 'can_redo' : for field getter : missing Type

// UNSUPPORTED : C value 'can_redo' : for field setter : missing Type

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'undo' : for field setter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field setter : missing Type

// UNSUPPORTED : C value 'begin_not_undoable_action' : for field getter : missing Type

// UNSUPPORTED : C value 'begin_not_undoable_action' : for field setter : missing Type

// UNSUPPORTED : C value 'end_not_undoable_action' : for field getter : missing Type

// UNSUPPORTED : C value 'end_not_undoable_action' : for field setter : missing Type

// UNSUPPORTED : C value 'can_undo_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'can_undo_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'can_redo_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'can_redo_changed' : for field setter : missing Type

// UndoManagerIfaceStruct creates an uninitialised UndoManagerIface.
func UndoManagerIfaceStruct() *UndoManagerIface {
	err := undoManagerIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UndoManagerIfaceNewFromNative(undoManagerIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUndoManagerIface)
	return structGo
}
func finalizeUndoManagerIface(obj *UndoManagerIface) {
	undoManagerIfaceStruct.Free(obj.native)
}

var viewClassStruct *gi.Struct
var viewClassStruct_Once sync.Once

func viewClassStruct_Set() error {
	var err error
	viewClassStruct_Once.Do(func() {
		viewClassStruct, err = gi.StructNew("GtkSource", "ViewClass")
	})
	return err
}

type ViewClass struct {
	native unsafe.Pointer
}

func ViewClassNewFromNative(native unsafe.Pointer) *ViewClass {
	return &ViewClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextViewClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gtk.TextViewClass'

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'undo' : for field setter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field setter : missing Type

// UNSUPPORTED : C value 'line_mark_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'line_mark_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'show_completion' : for field getter : missing Type

// UNSUPPORTED : C value 'show_completion' : for field setter : missing Type

// UNSUPPORTED : C value 'move_lines' : for field getter : missing Type

// UNSUPPORTED : C value 'move_lines' : for field setter : missing Type

// UNSUPPORTED : C value 'move_words' : for field getter : missing Type

// UNSUPPORTED : C value 'move_words' : for field setter : missing Type

// ViewClassStruct creates an uninitialised ViewClass.
func ViewClassStruct() *ViewClass {
	err := viewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ViewClassNewFromNative(viewClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeViewClass)
	return structGo
}
func finalizeViewClass(obj *ViewClass) {
	viewClassStruct.Free(obj.native)
}

var viewPrivateStruct *gi.Struct
var viewPrivateStruct_Once sync.Once

func viewPrivateStruct_Set() error {
	var err error
	viewPrivateStruct_Once.Do(func() {
		viewPrivateStruct, err = gi.StructNew("GtkSource", "ViewPrivate")
	})
	return err
}

type ViewPrivate struct {
	native unsafe.Pointer
}

func ViewPrivateNewFromNative(native unsafe.Pointer) *ViewPrivate {
	return &ViewPrivate{native: native}
}

// ViewPrivateStruct creates an uninitialised ViewPrivate.
func ViewPrivateStruct() *ViewPrivate {
	err := viewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ViewPrivateNewFromNative(viewPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeViewPrivate)
	return structGo
}
func finalizeViewPrivate(obj *ViewPrivate) {
	viewPrivateStruct.Free(obj.native)
}
