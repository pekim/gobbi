// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var bufferClassStruct *gi.Struct
var bufferClassStruct_Once sync.Once

func bufferClassStruct_Set() {
	bufferClassStruct_Once.Do(func() {
		bufferClassStruct = gi.StructNew("GtkSource", "BufferClass")
	})
}

type BufferClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextBufferClass'
	// UNSUPPORTED : C value 'undo' : missing Type
	// UNSUPPORTED : C value 'redo' : missing Type
	// UNSUPPORTED : C value 'bracket_matched' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
}

var bufferPrivateStruct *gi.Struct
var bufferPrivateStruct_Once sync.Once

func bufferPrivateStruct_Set() {
	bufferPrivateStruct_Once.Do(func() {
		bufferPrivateStruct = gi.StructNew("GtkSource", "BufferPrivate")
	})
}

type BufferPrivate struct {
	native uintptr
}

var completionClassStruct *gi.Struct
var completionClassStruct_Once sync.Once

func completionClassStruct_Set() {
	completionClassStruct_Once.Do(func() {
		completionClassStruct = gi.StructNew("GtkSource", "CompletionClass")
	})
}

type CompletionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'proposal_activated' : missing Type
	// UNSUPPORTED : C value 'show' : missing Type
	// UNSUPPORTED : C value 'hide' : missing Type
	// UNSUPPORTED : C value 'populate_context' : missing Type
	// UNSUPPORTED : C value 'move_cursor' : missing Type
	// UNSUPPORTED : C value 'move_page' : missing Type
	// UNSUPPORTED : C value 'activate_proposal' : missing Type
}

var completionContextClassStruct *gi.Struct
var completionContextClassStruct_Once sync.Once

func completionContextClassStruct_Set() {
	completionContextClassStruct_Once.Do(func() {
		completionContextClassStruct = gi.StructNew("GtkSource", "CompletionContextClass")
	})
}

type CompletionContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'
	// UNSUPPORTED : C value 'cancelled' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
}

var completionContextPrivateStruct *gi.Struct
var completionContextPrivateStruct_Once sync.Once

func completionContextPrivateStruct_Set() {
	completionContextPrivateStruct_Once.Do(func() {
		completionContextPrivateStruct = gi.StructNew("GtkSource", "CompletionContextPrivate")
	})
}

type CompletionContextPrivate struct {
	native uintptr
}

var completionInfoClassStruct *gi.Struct
var completionInfoClassStruct_Once sync.Once

func completionInfoClassStruct_Set() {
	completionInfoClassStruct_Once.Do(func() {
		completionInfoClassStruct = gi.StructNew("GtkSource", "CompletionInfoClass")
	})
}

type CompletionInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.WindowClass'
	// UNSUPPORTED : C value 'before_show' : missing Type
}

var completionInfoPrivateStruct *gi.Struct
var completionInfoPrivateStruct_Once sync.Once

func completionInfoPrivateStruct_Set() {
	completionInfoPrivateStruct_Once.Do(func() {
		completionInfoPrivateStruct = gi.StructNew("GtkSource", "CompletionInfoPrivate")
	})
}

type CompletionInfoPrivate struct {
	native uintptr
}

var completionItemClassStruct *gi.Struct
var completionItemClassStruct_Once sync.Once

func completionItemClassStruct_Set() {
	completionItemClassStruct_Once.Do(func() {
		completionItemClassStruct = gi.StructNew("GtkSource", "CompletionItemClass")
	})
}

type CompletionItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var completionItemPrivateStruct *gi.Struct
var completionItemPrivateStruct_Once sync.Once

func completionItemPrivateStruct_Set() {
	completionItemPrivateStruct_Once.Do(func() {
		completionItemPrivateStruct = gi.StructNew("GtkSource", "CompletionItemPrivate")
	})
}

type CompletionItemPrivate struct {
	native uintptr
}

var completionPrivateStruct *gi.Struct
var completionPrivateStruct_Once sync.Once

func completionPrivateStruct_Set() {
	completionPrivateStruct_Once.Do(func() {
		completionPrivateStruct = gi.StructNew("GtkSource", "CompletionPrivate")
	})
}

type CompletionPrivate struct {
	native uintptr
}

var completionProposalIfaceStruct *gi.Struct
var completionProposalIfaceStruct_Once sync.Once

func completionProposalIfaceStruct_Set() {
	completionProposalIfaceStruct_Once.Do(func() {
		completionProposalIfaceStruct = gi.StructNew("GtkSource", "CompletionProposalIface")
	})
}

type CompletionProposalIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_label' : missing Type
	// UNSUPPORTED : C value 'get_markup' : missing Type
	// UNSUPPORTED : C value 'get_text' : missing Type
	// UNSUPPORTED : C value 'get_icon' : missing Type
	// UNSUPPORTED : C value 'get_icon_name' : missing Type
	// UNSUPPORTED : C value 'get_gicon' : missing Type
	// UNSUPPORTED : C value 'get_info' : missing Type
	// UNSUPPORTED : C value 'hash' : missing Type
	// UNSUPPORTED : C value 'equal' : missing Type
	// UNSUPPORTED : C value 'changed' : missing Type
}

var completionProviderIfaceStruct *gi.Struct
var completionProviderIfaceStruct_Once sync.Once

func completionProviderIfaceStruct_Set() {
	completionProviderIfaceStruct_Once.Do(func() {
		completionProviderIfaceStruct = gi.StructNew("GtkSource", "CompletionProviderIface")
	})
}

type CompletionProviderIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_name' : missing Type
	// UNSUPPORTED : C value 'get_icon' : missing Type
	// UNSUPPORTED : C value 'get_icon_name' : missing Type
	// UNSUPPORTED : C value 'get_gicon' : missing Type
	// UNSUPPORTED : C value 'populate' : missing Type
	// UNSUPPORTED : C value 'match' : missing Type
	// UNSUPPORTED : C value 'get_activation' : missing Type
	// UNSUPPORTED : C value 'get_info_widget' : missing Type
	// UNSUPPORTED : C value 'update_info' : missing Type
	// UNSUPPORTED : C value 'get_start_iter' : missing Type
	// UNSUPPORTED : C value 'activate_proposal' : missing Type
	// UNSUPPORTED : C value 'get_interactive_delay' : missing Type
	// UNSUPPORTED : C value 'get_priority' : missing Type
}

var completionWordsClassStruct *gi.Struct
var completionWordsClassStruct_Once sync.Once

func completionWordsClassStruct_Set() {
	completionWordsClassStruct_Once.Do(func() {
		completionWordsClassStruct = gi.StructNew("GtkSource", "CompletionWordsClass")
	})
}

type CompletionWordsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var completionWordsPrivateStruct *gi.Struct
var completionWordsPrivateStruct_Once sync.Once

func completionWordsPrivateStruct_Set() {
	completionWordsPrivateStruct_Once.Do(func() {
		completionWordsPrivateStruct = gi.StructNew("GtkSource", "CompletionWordsPrivate")
	})
}

type CompletionWordsPrivate struct {
	native uintptr
}

var encodingStruct *gi.Struct
var encodingStruct_Once sync.Once

func encodingStruct_Set() {
	encodingStruct_Once.Do(func() {
		encodingStruct = gi.StructNew("GtkSource", "Encoding")
	})
}

type Encoding struct {
	native uintptr
}

var encodingCopyFunction *gi.Function
var encodingCopyFunction_Once sync.Once

func encodingCopyFunction_Set() {
	encodingCopyFunction_Once.Do(func() {
		encodingStruct_Set()
		encodingCopyFunction = encodingStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type gtk_source_encoding_copy.
func (recv *Encoding) Copy() *Encoding {
	encodingCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := encodingCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

var encodingFreeFunction *gi.Function
var encodingFreeFunction_Once sync.Once

func encodingFreeFunction_Set() {
	encodingFreeFunction_Once.Do(func() {
		encodingStruct_Set()
		encodingFreeFunction = encodingStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_source_encoding_free.
func (recv *Encoding) Free() {
	encodingFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	encodingFreeFunction.Invoke(inArgs[:], nil)

}

var encodingGetCharsetFunction *gi.Function
var encodingGetCharsetFunction_Once sync.Once

func encodingGetCharsetFunction_Set() {
	encodingGetCharsetFunction_Once.Do(func() {
		encodingStruct_Set()
		encodingGetCharsetFunction = encodingStruct.InvokerNew("get_charset")
	})
}

// GetCharset is a representation of the C type gtk_source_encoding_get_charset.
func (recv *Encoding) GetCharset() string {
	encodingGetCharsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := encodingGetCharsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var encodingGetNameFunction *gi.Function
var encodingGetNameFunction_Once sync.Once

func encodingGetNameFunction_Set() {
	encodingGetNameFunction_Once.Do(func() {
		encodingStruct_Set()
		encodingGetNameFunction = encodingStruct.InvokerNew("get_name")
	})
}

// GetName is a representation of the C type gtk_source_encoding_get_name.
func (recv *Encoding) GetName() string {
	encodingGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := encodingGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var encodingToStringFunction *gi.Function
var encodingToStringFunction_Once sync.Once

func encodingToStringFunction_Set() {
	encodingToStringFunction_Once.Do(func() {
		encodingStruct_Set()
		encodingToStringFunction = encodingStruct.InvokerNew("to_string")
	})
}

// ToString is a representation of the C type gtk_source_encoding_to_string.
func (recv *Encoding) ToString() string {
	encodingToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := encodingToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fileClassStruct *gi.Struct
var fileClassStruct_Once sync.Once

func fileClassStruct_Set() {
	fileClassStruct_Once.Do(func() {
		fileClassStruct = gi.StructNew("GtkSource", "FileClass")
	})
}

type FileClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileLoaderClassStruct *gi.Struct
var fileLoaderClassStruct_Once sync.Once

func fileLoaderClassStruct_Set() {
	fileLoaderClassStruct_Once.Do(func() {
		fileLoaderClassStruct = gi.StructNew("GtkSource", "FileLoaderClass")
	})
}

type FileLoaderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileLoaderPrivateStruct *gi.Struct
var fileLoaderPrivateStruct_Once sync.Once

func fileLoaderPrivateStruct_Set() {
	fileLoaderPrivateStruct_Once.Do(func() {
		fileLoaderPrivateStruct = gi.StructNew("GtkSource", "FileLoaderPrivate")
	})
}

type FileLoaderPrivate struct {
	native uintptr
}

var filePrivateStruct *gi.Struct
var filePrivateStruct_Once sync.Once

func filePrivateStruct_Set() {
	filePrivateStruct_Once.Do(func() {
		filePrivateStruct = gi.StructNew("GtkSource", "FilePrivate")
	})
}

type FilePrivate struct {
	native uintptr
}

var fileSaverClassStruct *gi.Struct
var fileSaverClassStruct_Once sync.Once

func fileSaverClassStruct_Set() {
	fileSaverClassStruct_Once.Do(func() {
		fileSaverClassStruct = gi.StructNew("GtkSource", "FileSaverClass")
	})
}

type FileSaverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileSaverPrivateStruct *gi.Struct
var fileSaverPrivateStruct_Once sync.Once

func fileSaverPrivateStruct_Set() {
	fileSaverPrivateStruct_Once.Do(func() {
		fileSaverPrivateStruct = gi.StructNew("GtkSource", "FileSaverPrivate")
	})
}

type FileSaverPrivate struct {
	native uintptr
}

var gutterClassStruct *gi.Struct
var gutterClassStruct_Once sync.Once

func gutterClassStruct_Set() {
	gutterClassStruct_Once.Do(func() {
		gutterClassStruct = gi.StructNew("GtkSource", "GutterClass")
	})
}

type GutterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var gutterPrivateStruct *gi.Struct
var gutterPrivateStruct_Once sync.Once

func gutterPrivateStruct_Set() {
	gutterPrivateStruct_Once.Do(func() {
		gutterPrivateStruct = gi.StructNew("GtkSource", "GutterPrivate")
	})
}

type GutterPrivate struct {
	native uintptr
}

var gutterRendererClassStruct *gi.Struct
var gutterRendererClassStruct_Once sync.Once

func gutterRendererClassStruct_Set() {
	gutterRendererClassStruct_Once.Do(func() {
		gutterRendererClassStruct = gi.StructNew("GtkSource", "GutterRendererClass")
	})
}

type GutterRendererClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'
	// UNSUPPORTED : C value 'begin' : missing Type
	// UNSUPPORTED : C value 'draw' : missing Type
	// UNSUPPORTED : C value 'end' : missing Type
	// UNSUPPORTED : C value 'change_view' : missing Type
	// UNSUPPORTED : C value 'change_buffer' : missing Type
	// UNSUPPORTED : C value 'query_activatable' : missing Type
	// UNSUPPORTED : C value 'activate' : missing Type
	// UNSUPPORTED : C value 'queue_draw' : missing Type
	// UNSUPPORTED : C value 'query_tooltip' : missing Type
	// UNSUPPORTED : C value 'query_data' : missing Type
}

var gutterRendererPixbufClassStruct *gi.Struct
var gutterRendererPixbufClassStruct_Once sync.Once

func gutterRendererPixbufClassStruct_Set() {
	gutterRendererPixbufClassStruct_Once.Do(func() {
		gutterRendererPixbufClassStruct = gi.StructNew("GtkSource", "GutterRendererPixbufClass")
	})
}

type GutterRendererPixbufClass struct {
	native uintptr
}

var gutterRendererPixbufPrivateStruct *gi.Struct
var gutterRendererPixbufPrivateStruct_Once sync.Once

func gutterRendererPixbufPrivateStruct_Set() {
	gutterRendererPixbufPrivateStruct_Once.Do(func() {
		gutterRendererPixbufPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPixbufPrivate")
	})
}

type GutterRendererPixbufPrivate struct {
	native uintptr
}

var gutterRendererPrivateStruct *gi.Struct
var gutterRendererPrivateStruct_Once sync.Once

func gutterRendererPrivateStruct_Set() {
	gutterRendererPrivateStruct_Once.Do(func() {
		gutterRendererPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPrivate")
	})
}

type GutterRendererPrivate struct {
	native uintptr
}

var gutterRendererTextClassStruct *gi.Struct
var gutterRendererTextClassStruct_Once sync.Once

func gutterRendererTextClassStruct_Set() {
	gutterRendererTextClassStruct_Once.Do(func() {
		gutterRendererTextClassStruct = gi.StructNew("GtkSource", "GutterRendererTextClass")
	})
}

type GutterRendererTextClass struct {
	native uintptr
}

var gutterRendererTextPrivateStruct *gi.Struct
var gutterRendererTextPrivateStruct_Once sync.Once

func gutterRendererTextPrivateStruct_Set() {
	gutterRendererTextPrivateStruct_Once.Do(func() {
		gutterRendererTextPrivateStruct = gi.StructNew("GtkSource", "GutterRendererTextPrivate")
	})
}

type GutterRendererTextPrivate struct {
	native uintptr
}

var languageClassStruct *gi.Struct
var languageClassStruct_Once sync.Once

func languageClassStruct_Set() {
	languageClassStruct_Once.Do(func() {
		languageClassStruct = gi.StructNew("GtkSource", "LanguageClass")
	})
}

type LanguageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var languageManagerClassStruct *gi.Struct
var languageManagerClassStruct_Once sync.Once

func languageManagerClassStruct_Set() {
	languageManagerClassStruct_Once.Do(func() {
		languageManagerClassStruct = gi.StructNew("GtkSource", "LanguageManagerClass")
	})
}

type LanguageManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved4' : missing Type
}

var languageManagerPrivateStruct *gi.Struct
var languageManagerPrivateStruct_Once sync.Once

func languageManagerPrivateStruct_Set() {
	languageManagerPrivateStruct_Once.Do(func() {
		languageManagerPrivateStruct = gi.StructNew("GtkSource", "LanguageManagerPrivate")
	})
}

type LanguageManagerPrivate struct {
	native uintptr
}

var languagePrivateStruct *gi.Struct
var languagePrivateStruct_Once sync.Once

func languagePrivateStruct_Set() {
	languagePrivateStruct_Once.Do(func() {
		languagePrivateStruct = gi.StructNew("GtkSource", "LanguagePrivate")
	})
}

type LanguagePrivate struct {
	native uintptr
}

var mapClassStruct *gi.Struct
var mapClassStruct_Once sync.Once

func mapClassStruct_Set() {
	mapClassStruct_Once.Do(func() {
		mapClassStruct = gi.StructNew("GtkSource", "MapClass")
	})
}

type MapClass struct {
	native      uintptr
	ParentClass *ViewClass
	// UNSUPPORTED : C value 'padding' : missing Type
}

var markAttributesClassStruct *gi.Struct
var markAttributesClassStruct_Once sync.Once

func markAttributesClassStruct_Set() {
	markAttributesClassStruct_Once.Do(func() {
		markAttributesClassStruct = gi.StructNew("GtkSource", "MarkAttributesClass")
	})
}

type MarkAttributesClass struct {
	native uintptr
}

var markAttributesPrivateStruct *gi.Struct
var markAttributesPrivateStruct_Once sync.Once

func markAttributesPrivateStruct_Set() {
	markAttributesPrivateStruct_Once.Do(func() {
		markAttributesPrivateStruct = gi.StructNew("GtkSource", "MarkAttributesPrivate")
	})
}

type MarkAttributesPrivate struct {
	native uintptr
}

var markClassStruct *gi.Struct
var markClassStruct_Once sync.Once

func markClassStruct_Set() {
	markClassStruct_Once.Do(func() {
		markClassStruct = gi.StructNew("GtkSource", "MarkClass")
	})
}

type MarkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextMarkClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var markPrivateStruct *gi.Struct
var markPrivateStruct_Once sync.Once

func markPrivateStruct_Set() {
	markPrivateStruct_Once.Do(func() {
		markPrivateStruct = gi.StructNew("GtkSource", "MarkPrivate")
	})
}

type MarkPrivate struct {
	native uintptr
}

var printCompositorClassStruct *gi.Struct
var printCompositorClassStruct_Once sync.Once

func printCompositorClassStruct_Set() {
	printCompositorClassStruct_Once.Do(func() {
		printCompositorClassStruct = gi.StructNew("GtkSource", "PrintCompositorClass")
	})
}

type PrintCompositorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var printCompositorPrivateStruct *gi.Struct
var printCompositorPrivateStruct_Once sync.Once

func printCompositorPrivateStruct_Set() {
	printCompositorPrivateStruct_Once.Do(func() {
		printCompositorPrivateStruct = gi.StructNew("GtkSource", "PrintCompositorPrivate")
	})
}

type PrintCompositorPrivate struct {
	native uintptr
}

var regionClassStruct *gi.Struct
var regionClassStruct_Once sync.Once

func regionClassStruct_Set() {
	regionClassStruct_Once.Do(func() {
		regionClassStruct = gi.StructNew("GtkSource", "RegionClass")
	})
}

type RegionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var regionIterStruct *gi.Struct
var regionIterStruct_Once sync.Once

func regionIterStruct_Set() {
	regionIterStruct_Once.Do(func() {
		regionIterStruct = gi.StructNew("GtkSource", "RegionIter")
	})
}

type RegionIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_source_region_iter_get_subregion' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_is_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_next' : return type 'gboolean' not supported

var searchContextClassStruct *gi.Struct
var searchContextClassStruct_Once sync.Once

func searchContextClassStruct_Set() {
	searchContextClassStruct_Once.Do(func() {
		searchContextClassStruct = gi.StructNew("GtkSource", "SearchContextClass")
	})
}

type SearchContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var searchContextPrivateStruct *gi.Struct
var searchContextPrivateStruct_Once sync.Once

func searchContextPrivateStruct_Set() {
	searchContextPrivateStruct_Once.Do(func() {
		searchContextPrivateStruct = gi.StructNew("GtkSource", "SearchContextPrivate")
	})
}

type SearchContextPrivate struct {
	native uintptr
}

var searchSettingsClassStruct *gi.Struct
var searchSettingsClassStruct_Once sync.Once

func searchSettingsClassStruct_Set() {
	searchSettingsClassStruct_Once.Do(func() {
		searchSettingsClassStruct = gi.StructNew("GtkSource", "SearchSettingsClass")
	})
}

type SearchSettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var searchSettingsPrivateStruct *gi.Struct
var searchSettingsPrivateStruct_Once sync.Once

func searchSettingsPrivateStruct_Set() {
	searchSettingsPrivateStruct_Once.Do(func() {
		searchSettingsPrivateStruct = gi.StructNew("GtkSource", "SearchSettingsPrivate")
	})
}

type SearchSettingsPrivate struct {
	native uintptr
}

var spaceDrawerClassStruct *gi.Struct
var spaceDrawerClassStruct_Once sync.Once

func spaceDrawerClassStruct_Set() {
	spaceDrawerClassStruct_Once.Do(func() {
		spaceDrawerClassStruct = gi.StructNew("GtkSource", "SpaceDrawerClass")
	})
}

type SpaceDrawerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var spaceDrawerPrivateStruct *gi.Struct
var spaceDrawerPrivateStruct_Once sync.Once

func spaceDrawerPrivateStruct_Set() {
	spaceDrawerPrivateStruct_Once.Do(func() {
		spaceDrawerPrivateStruct = gi.StructNew("GtkSource", "SpaceDrawerPrivate")
	})
}

type SpaceDrawerPrivate struct {
	native uintptr
}

var styleClassStruct *gi.Struct
var styleClassStruct_Once sync.Once

func styleClassStruct_Set() {
	styleClassStruct_Once.Do(func() {
		styleClassStruct = gi.StructNew("GtkSource", "StyleClass")
	})
}

type StyleClass struct {
	native uintptr
}

var styleSchemeChooserButtonClassStruct *gi.Struct
var styleSchemeChooserButtonClassStruct_Once sync.Once

func styleSchemeChooserButtonClassStruct_Set() {
	styleSchemeChooserButtonClassStruct_Once.Do(func() {
		styleSchemeChooserButtonClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserButtonClass")
	})
}

type StyleSchemeChooserButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.ButtonClass'
}

var styleSchemeChooserInterfaceStruct *gi.Struct
var styleSchemeChooserInterfaceStruct_Once sync.Once

func styleSchemeChooserInterfaceStruct_Set() {
	styleSchemeChooserInterfaceStruct_Once.Do(func() {
		styleSchemeChooserInterfaceStruct = gi.StructNew("GtkSource", "StyleSchemeChooserInterface")
	})
}

type StyleSchemeChooserInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_style_scheme' : missing Type
	// UNSUPPORTED : C value 'set_style_scheme' : missing Type
	// UNSUPPORTED : C value 'padding' : missing Type
}

var styleSchemeChooserWidgetClassStruct *gi.Struct
var styleSchemeChooserWidgetClassStruct_Once sync.Once

func styleSchemeChooserWidgetClassStruct_Set() {
	styleSchemeChooserWidgetClassStruct_Once.Do(func() {
		styleSchemeChooserWidgetClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserWidgetClass")
	})
}

type StyleSchemeChooserWidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.BinClass'
}

var styleSchemeClassStruct *gi.Struct
var styleSchemeClassStruct_Once sync.Once

func styleSchemeClassStruct_Set() {
	styleSchemeClassStruct_Once.Do(func() {
		styleSchemeClassStruct = gi.StructNew("GtkSource", "StyleSchemeClass")
	})
}

type StyleSchemeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'base_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var styleSchemeManagerClassStruct *gi.Struct
var styleSchemeManagerClassStruct_Once sync.Once

func styleSchemeManagerClassStruct_Set() {
	styleSchemeManagerClassStruct_Once.Do(func() {
		styleSchemeManagerClassStruct = gi.StructNew("GtkSource", "StyleSchemeManagerClass")
	})
}

type StyleSchemeManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved4' : missing Type
}

var styleSchemeManagerPrivateStruct *gi.Struct
var styleSchemeManagerPrivateStruct_Once sync.Once

func styleSchemeManagerPrivateStruct_Set() {
	styleSchemeManagerPrivateStruct_Once.Do(func() {
		styleSchemeManagerPrivateStruct = gi.StructNew("GtkSource", "StyleSchemeManagerPrivate")
	})
}

type StyleSchemeManagerPrivate struct {
	native uintptr
}

var styleSchemePrivateStruct *gi.Struct
var styleSchemePrivateStruct_Once sync.Once

func styleSchemePrivateStruct_Set() {
	styleSchemePrivateStruct_Once.Do(func() {
		styleSchemePrivateStruct = gi.StructNew("GtkSource", "StyleSchemePrivate")
	})
}

type StyleSchemePrivate struct {
	native uintptr
}

var tagClassStruct *gi.Struct
var tagClassStruct_Once sync.Once

func tagClassStruct_Set() {
	tagClassStruct_Once.Do(func() {
		tagClassStruct = gi.StructNew("GtkSource", "TagClass")
	})
}

type TagClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextTagClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var undoManagerIfaceStruct *gi.Struct
var undoManagerIfaceStruct_Once sync.Once

func undoManagerIfaceStruct_Set() {
	undoManagerIfaceStruct_Once.Do(func() {
		undoManagerIfaceStruct = gi.StructNew("GtkSource", "UndoManagerIface")
	})
}

type UndoManagerIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'can_undo' : missing Type
	// UNSUPPORTED : C value 'can_redo' : missing Type
	// UNSUPPORTED : C value 'undo' : missing Type
	// UNSUPPORTED : C value 'redo' : missing Type
	// UNSUPPORTED : C value 'begin_not_undoable_action' : missing Type
	// UNSUPPORTED : C value 'end_not_undoable_action' : missing Type
	// UNSUPPORTED : C value 'can_undo_changed' : missing Type
	// UNSUPPORTED : C value 'can_redo_changed' : missing Type
}

var viewClassStruct *gi.Struct
var viewClassStruct_Once sync.Once

func viewClassStruct_Set() {
	viewClassStruct_Once.Do(func() {
		viewClassStruct = gi.StructNew("GtkSource", "ViewClass")
	})
}

type ViewClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextViewClass'
	// UNSUPPORTED : C value 'undo' : missing Type
	// UNSUPPORTED : C value 'redo' : missing Type
	// UNSUPPORTED : C value 'line_mark_activated' : missing Type
	// UNSUPPORTED : C value 'show_completion' : missing Type
	// UNSUPPORTED : C value 'move_lines' : missing Type
	// UNSUPPORTED : C value 'move_words' : missing Type
}

var viewPrivateStruct *gi.Struct
var viewPrivateStruct_Once sync.Once

func viewPrivateStruct_Set() {
	viewPrivateStruct_Once.Do(func() {
		viewPrivateStruct = gi.StructNew("GtkSource", "ViewPrivate")
	})
}

type ViewPrivate struct {
	native uintptr
}
