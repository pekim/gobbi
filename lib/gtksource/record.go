// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextBufferClass'

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'bracket_matched' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// BufferClassStruct creates an uninitialised BufferClass.
func BufferClassStruct() *BufferClass {
	err := bufferClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BufferClass{native: bufferClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// BufferPrivateStruct creates an uninitialised BufferPrivate.
func BufferPrivateStruct() *BufferPrivate {
	err := bufferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BufferPrivate{native: bufferPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'proposal_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'show' : for field getter : missing Type

// UNSUPPORTED : C value 'hide' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_context' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_page' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field getter : missing Type

// CompletionClassStruct creates an uninitialised CompletionClass.
func CompletionClassStruct() *CompletionClass {
	err := completionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionClass{native: completionClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'cancelled' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// CompletionContextClassStruct creates an uninitialised CompletionContextClass.
func CompletionContextClassStruct() *CompletionContextClass {
	err := completionContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionContextClass{native: completionContextClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// CompletionContextPrivateStruct creates an uninitialised CompletionContextPrivate.
func CompletionContextPrivateStruct() *CompletionContextPrivate {
	err := completionContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionContextPrivate{native: completionContextPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.WindowClass'

// UNSUPPORTED : C value 'before_show' : for field getter : missing Type

// CompletionInfoClassStruct creates an uninitialised CompletionInfoClass.
func CompletionInfoClassStruct() *CompletionInfoClass {
	err := completionInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionInfoClass{native: completionInfoClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// CompletionInfoPrivateStruct creates an uninitialised CompletionInfoPrivate.
func CompletionInfoPrivateStruct() *CompletionInfoPrivate {
	err := completionInfoPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionInfoPrivate{native: completionInfoPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// CompletionItemClassStruct creates an uninitialised CompletionItemClass.
func CompletionItemClassStruct() *CompletionItemClass {
	err := completionItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionItemClass{native: completionItemClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// CompletionItemPrivateStruct creates an uninitialised CompletionItemPrivate.
func CompletionItemPrivateStruct() *CompletionItemPrivate {
	err := completionItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionItemPrivate{native: completionItemPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// CompletionPrivateStruct creates an uninitialised CompletionPrivate.
func CompletionPrivateStruct() *CompletionPrivate {
	err := completionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionPrivate{native: completionPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_label' : for field getter : missing Type

// UNSUPPORTED : C value 'get_markup' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info' : for field getter : missing Type

// UNSUPPORTED : C value 'hash' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// CompletionProposalIfaceStruct creates an uninitialised CompletionProposalIface.
func CompletionProposalIfaceStruct() *CompletionProposalIface {
	err := completionProposalIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionProposalIface{native: completionProposalIfaceStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_gicon' : for field getter : missing Type

// UNSUPPORTED : C value 'populate' : for field getter : missing Type

// UNSUPPORTED : C value 'match' : for field getter : missing Type

// UNSUPPORTED : C value 'get_activation' : for field getter : missing Type

// UNSUPPORTED : C value 'get_info_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'update_info' : for field getter : missing Type

// UNSUPPORTED : C value 'get_start_iter' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_proposal' : for field getter : missing Type

// UNSUPPORTED : C value 'get_interactive_delay' : for field getter : missing Type

// UNSUPPORTED : C value 'get_priority' : for field getter : missing Type

// CompletionProviderIfaceStruct creates an uninitialised CompletionProviderIface.
func CompletionProviderIfaceStruct() *CompletionProviderIface {
	err := completionProviderIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionProviderIface{native: completionProviderIfaceStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// CompletionWordsClassStruct creates an uninitialised CompletionWordsClass.
func CompletionWordsClassStruct() *CompletionWordsClass {
	err := completionWordsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionWordsClass{native: completionWordsClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// CompletionWordsPrivateStruct creates an uninitialised CompletionWordsPrivate.
func CompletionWordsPrivateStruct() *CompletionWordsPrivate {
	err := completionWordsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CompletionWordsPrivate{native: completionWordsPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// EncodingStruct creates an uninitialised Encoding.
func EncodingStruct() *Encoding {
	err := encodingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Encoding{native: encodingStruct.Alloc()}
	return structGo
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

	retGo := &Encoding{native: ret.Pointer()}

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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// FileClassStruct creates an uninitialised FileClass.
func FileClassStruct() *FileClass {
	err := fileClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileClass{native: fileClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// FileLoaderClassStruct creates an uninitialised FileLoaderClass.
func FileLoaderClassStruct() *FileLoaderClass {
	err := fileLoaderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileLoaderClass{native: fileLoaderClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// FileLoaderPrivateStruct creates an uninitialised FileLoaderPrivate.
func FileLoaderPrivateStruct() *FileLoaderPrivate {
	err := fileLoaderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileLoaderPrivate{native: fileLoaderPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// FilePrivateStruct creates an uninitialised FilePrivate.
func FilePrivateStruct() *FilePrivate {
	err := filePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FilePrivate{native: filePrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// FileSaverClassStruct creates an uninitialised FileSaverClass.
func FileSaverClassStruct() *FileSaverClass {
	err := fileSaverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileSaverClass{native: fileSaverClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// FileSaverPrivateStruct creates an uninitialised FileSaverPrivate.
func FileSaverPrivateStruct() *FileSaverPrivate {
	err := fileSaverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileSaverPrivate{native: fileSaverPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// GutterClassStruct creates an uninitialised GutterClass.
func GutterClassStruct() *GutterClass {
	err := gutterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterClass{native: gutterClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterPrivateStruct creates an uninitialised GutterPrivate.
func GutterPrivateStruct() *GutterPrivate {
	err := gutterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterPrivate{native: gutterPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'begin' : for field getter : missing Type

// UNSUPPORTED : C value 'draw' : for field getter : missing Type

// UNSUPPORTED : C value 'end' : for field getter : missing Type

// UNSUPPORTED : C value 'change_view' : for field getter : missing Type

// UNSUPPORTED : C value 'change_buffer' : for field getter : missing Type

// UNSUPPORTED : C value 'query_activatable' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'queue_draw' : for field getter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field getter : missing Type

// UNSUPPORTED : C value 'query_data' : for field getter : missing Type

// GutterRendererClassStruct creates an uninitialised GutterRendererClass.
func GutterRendererClassStruct() *GutterRendererClass {
	err := gutterRendererClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererClass{native: gutterRendererClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterRendererPixbufClassStruct creates an uninitialised GutterRendererPixbufClass.
func GutterRendererPixbufClassStruct() *GutterRendererPixbufClass {
	err := gutterRendererPixbufClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererPixbufClass{native: gutterRendererPixbufClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterRendererPixbufPrivateStruct creates an uninitialised GutterRendererPixbufPrivate.
func GutterRendererPixbufPrivateStruct() *GutterRendererPixbufPrivate {
	err := gutterRendererPixbufPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererPixbufPrivate{native: gutterRendererPixbufPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterRendererPrivateStruct creates an uninitialised GutterRendererPrivate.
func GutterRendererPrivateStruct() *GutterRendererPrivate {
	err := gutterRendererPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererPrivate{native: gutterRendererPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterRendererTextClassStruct creates an uninitialised GutterRendererTextClass.
func GutterRendererTextClassStruct() *GutterRendererTextClass {
	err := gutterRendererTextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererTextClass{native: gutterRendererTextClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// GutterRendererTextPrivateStruct creates an uninitialised GutterRendererTextPrivate.
func GutterRendererTextPrivateStruct() *GutterRendererTextPrivate {
	err := gutterRendererTextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GutterRendererTextPrivate{native: gutterRendererTextPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// LanguageClassStruct creates an uninitialised LanguageClass.
func LanguageClassStruct() *LanguageClass {
	err := languageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LanguageClass{native: languageClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field getter : missing Type

// LanguageManagerClassStruct creates an uninitialised LanguageManagerClass.
func LanguageManagerClassStruct() *LanguageManagerClass {
	err := languageManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LanguageManagerClass{native: languageManagerClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// LanguageManagerPrivateStruct creates an uninitialised LanguageManagerPrivate.
func LanguageManagerPrivateStruct() *LanguageManagerPrivate {
	err := languageManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LanguageManagerPrivate{native: languageManagerPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// LanguagePrivateStruct creates an uninitialised LanguagePrivate.
func LanguagePrivateStruct() *LanguagePrivate {
	err := languagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LanguagePrivate{native: languagePrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *MapClass) ParentClass() *ViewClass {
	argValue := gi.FieldGet(mapClassStruct, recv.native, "parent_class")
	value := &ViewClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// MapClassStruct creates an uninitialised MapClass.
func MapClassStruct() *MapClass {
	err := mapClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MapClass{native: mapClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// MarkAttributesClassStruct creates an uninitialised MarkAttributesClass.
func MarkAttributesClassStruct() *MarkAttributesClass {
	err := markAttributesClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MarkAttributesClass{native: markAttributesClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// MarkAttributesPrivateStruct creates an uninitialised MarkAttributesPrivate.
func MarkAttributesPrivateStruct() *MarkAttributesPrivate {
	err := markAttributesPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MarkAttributesPrivate{native: markAttributesPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextMarkClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// MarkClassStruct creates an uninitialised MarkClass.
func MarkClassStruct() *MarkClass {
	err := markClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MarkClass{native: markClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// MarkPrivateStruct creates an uninitialised MarkPrivate.
func MarkPrivateStruct() *MarkPrivate {
	err := markPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MarkPrivate{native: markPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// PrintCompositorClassStruct creates an uninitialised PrintCompositorClass.
func PrintCompositorClassStruct() *PrintCompositorClass {
	err := printCompositorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PrintCompositorClass{native: printCompositorClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// PrintCompositorPrivateStruct creates an uninitialised PrintCompositorPrivate.
func PrintCompositorPrivateStruct() *PrintCompositorPrivate {
	err := printCompositorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PrintCompositorPrivate{native: printCompositorPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// RegionClassStruct creates an uninitialised RegionClass.
func RegionClassStruct() *RegionClass {
	err := regionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RegionClass{native: regionClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// RegionIterStruct creates an uninitialised RegionIter.
func RegionIterStruct() *RegionIter {
	err := regionIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RegionIter{native: regionIterStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// SearchContextClassStruct creates an uninitialised SearchContextClass.
func SearchContextClassStruct() *SearchContextClass {
	err := searchContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchContextClass{native: searchContextClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// SearchContextPrivateStruct creates an uninitialised SearchContextPrivate.
func SearchContextPrivateStruct() *SearchContextPrivate {
	err := searchContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchContextPrivate{native: searchContextPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// SearchSettingsClassStruct creates an uninitialised SearchSettingsClass.
func SearchSettingsClassStruct() *SearchSettingsClass {
	err := searchSettingsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchSettingsClass{native: searchSettingsClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// SearchSettingsPrivateStruct creates an uninitialised SearchSettingsPrivate.
func SearchSettingsPrivateStruct() *SearchSettingsPrivate {
	err := searchSettingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchSettingsPrivate{native: searchSettingsPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// SpaceDrawerClassStruct creates an uninitialised SpaceDrawerClass.
func SpaceDrawerClassStruct() *SpaceDrawerClass {
	err := spaceDrawerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpaceDrawerClass{native: spaceDrawerClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// SpaceDrawerPrivateStruct creates an uninitialised SpaceDrawerPrivate.
func SpaceDrawerPrivateStruct() *SpaceDrawerPrivate {
	err := spaceDrawerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpaceDrawerPrivate{native: spaceDrawerPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// StyleClassStruct creates an uninitialised StyleClass.
func StyleClassStruct() *StyleClass {
	err := styleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleClass{native: styleClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.ButtonClass'

// StyleSchemeChooserButtonClassStruct creates an uninitialised StyleSchemeChooserButtonClass.
func StyleSchemeChooserButtonClassStruct() *StyleSchemeChooserButtonClass {
	err := styleSchemeChooserButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeChooserButtonClass{native: styleSchemeChooserButtonClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'base_interface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_style_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'set_style_scheme' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// StyleSchemeChooserInterfaceStruct creates an uninitialised StyleSchemeChooserInterface.
func StyleSchemeChooserInterfaceStruct() *StyleSchemeChooserInterface {
	err := styleSchemeChooserInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeChooserInterface{native: styleSchemeChooserInterfaceStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'Gtk.BinClass'

// StyleSchemeChooserWidgetClassStruct creates an uninitialised StyleSchemeChooserWidgetClass.
func StyleSchemeChooserWidgetClassStruct() *StyleSchemeChooserWidgetClass {
	err := styleSchemeChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeChooserWidgetClass{native: styleSchemeChooserWidgetClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'base_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// StyleSchemeClassStruct creates an uninitialised StyleSchemeClass.
func StyleSchemeClassStruct() *StyleSchemeClass {
	err := styleSchemeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeClass{native: styleSchemeClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_source_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_source_reserved4' : for field getter : missing Type

// StyleSchemeManagerClassStruct creates an uninitialised StyleSchemeManagerClass.
func StyleSchemeManagerClassStruct() *StyleSchemeManagerClass {
	err := styleSchemeManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeManagerClass{native: styleSchemeManagerClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// StyleSchemeManagerPrivateStruct creates an uninitialised StyleSchemeManagerPrivate.
func StyleSchemeManagerPrivateStruct() *StyleSchemeManagerPrivate {
	err := styleSchemeManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemeManagerPrivate{native: styleSchemeManagerPrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// StyleSchemePrivateStruct creates an uninitialised StyleSchemePrivate.
func StyleSchemePrivateStruct() *StyleSchemePrivate {
	err := styleSchemePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleSchemePrivate{native: styleSchemePrivateStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextTagClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// TagClassStruct creates an uninitialised TagClass.
func TagClassStruct() *TagClass {
	err := tagClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TagClass{native: tagClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'can_undo' : for field getter : missing Type

// UNSUPPORTED : C value 'can_redo' : for field getter : missing Type

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'begin_not_undoable_action' : for field getter : missing Type

// UNSUPPORTED : C value 'end_not_undoable_action' : for field getter : missing Type

// UNSUPPORTED : C value 'can_undo_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'can_redo_changed' : for field getter : missing Type

// UndoManagerIfaceStruct creates an uninitialised UndoManagerIface.
func UndoManagerIfaceStruct() *UndoManagerIface {
	err := undoManagerIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UndoManagerIface{native: undoManagerIfaceStruct.Alloc()}
	return structGo
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gtk.TextViewClass'

// UNSUPPORTED : C value 'undo' : for field getter : missing Type

// UNSUPPORTED : C value 'redo' : for field getter : missing Type

// UNSUPPORTED : C value 'line_mark_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'show_completion' : for field getter : missing Type

// UNSUPPORTED : C value 'move_lines' : for field getter : missing Type

// UNSUPPORTED : C value 'move_words' : for field getter : missing Type

// ViewClassStruct creates an uninitialised ViewClass.
func ViewClassStruct() *ViewClass {
	err := viewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ViewClass{native: viewClassStruct.Alloc()}
	return structGo
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
	native uintptr
}

// ViewPrivateStruct creates an uninitialised ViewPrivate.
func ViewPrivateStruct() *ViewPrivate {
	err := viewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ViewPrivate{native: viewPrivateStruct.Alloc()}
	return structGo
}
