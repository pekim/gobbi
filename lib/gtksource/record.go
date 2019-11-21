// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var bufferClassStruct *gi.Struct
var bufferClassStructOnce sync.Once

func bufferClassStructSet() {
	bufferClassStructOnce.Do(func() {
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
var bufferPrivateStructOnce sync.Once

func bufferPrivateStructSet() {
	bufferPrivateStructOnce.Do(func() {
		bufferPrivateStruct = gi.StructNew("GtkSource", "BufferPrivate")
	})
}

type BufferPrivate struct {
	native uintptr
}

var completionClassStruct *gi.Struct
var completionClassStructOnce sync.Once

func completionClassStructSet() {
	completionClassStructOnce.Do(func() {
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
var completionContextClassStructOnce sync.Once

func completionContextClassStructSet() {
	completionContextClassStructOnce.Do(func() {
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
var completionContextPrivateStructOnce sync.Once

func completionContextPrivateStructSet() {
	completionContextPrivateStructOnce.Do(func() {
		completionContextPrivateStruct = gi.StructNew("GtkSource", "CompletionContextPrivate")
	})
}

type CompletionContextPrivate struct {
	native uintptr
}

var completionInfoClassStruct *gi.Struct
var completionInfoClassStructOnce sync.Once

func completionInfoClassStructSet() {
	completionInfoClassStructOnce.Do(func() {
		completionInfoClassStruct = gi.StructNew("GtkSource", "CompletionInfoClass")
	})
}

type CompletionInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.WindowClass'
	// UNSUPPORTED : C value 'before_show' : missing Type
}

var completionInfoPrivateStruct *gi.Struct
var completionInfoPrivateStructOnce sync.Once

func completionInfoPrivateStructSet() {
	completionInfoPrivateStructOnce.Do(func() {
		completionInfoPrivateStruct = gi.StructNew("GtkSource", "CompletionInfoPrivate")
	})
}

type CompletionInfoPrivate struct {
	native uintptr
}

var completionItemClassStruct *gi.Struct
var completionItemClassStructOnce sync.Once

func completionItemClassStructSet() {
	completionItemClassStructOnce.Do(func() {
		completionItemClassStruct = gi.StructNew("GtkSource", "CompletionItemClass")
	})
}

type CompletionItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var completionItemPrivateStruct *gi.Struct
var completionItemPrivateStructOnce sync.Once

func completionItemPrivateStructSet() {
	completionItemPrivateStructOnce.Do(func() {
		completionItemPrivateStruct = gi.StructNew("GtkSource", "CompletionItemPrivate")
	})
}

type CompletionItemPrivate struct {
	native uintptr
}

var completionPrivateStruct *gi.Struct
var completionPrivateStructOnce sync.Once

func completionPrivateStructSet() {
	completionPrivateStructOnce.Do(func() {
		completionPrivateStruct = gi.StructNew("GtkSource", "CompletionPrivate")
	})
}

type CompletionPrivate struct {
	native uintptr
}

var completionProposalIfaceStruct *gi.Struct
var completionProposalIfaceStructOnce sync.Once

func completionProposalIfaceStructSet() {
	completionProposalIfaceStructOnce.Do(func() {
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
var completionProviderIfaceStructOnce sync.Once

func completionProviderIfaceStructSet() {
	completionProviderIfaceStructOnce.Do(func() {
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
var completionWordsClassStructOnce sync.Once

func completionWordsClassStructSet() {
	completionWordsClassStructOnce.Do(func() {
		completionWordsClassStruct = gi.StructNew("GtkSource", "CompletionWordsClass")
	})
}

type CompletionWordsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var completionWordsPrivateStruct *gi.Struct
var completionWordsPrivateStructOnce sync.Once

func completionWordsPrivateStructSet() {
	completionWordsPrivateStructOnce.Do(func() {
		completionWordsPrivateStruct = gi.StructNew("GtkSource", "CompletionWordsPrivate")
	})
}

type CompletionWordsPrivate struct {
	native uintptr
}

var encodingStruct *gi.Struct
var encodingStructOnce sync.Once

func encodingStructSet() {
	encodingStructOnce.Do(func() {
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
		encodingCopyFunction = gi.FunctionInvokerNew("GtkSource", "copy")
	})
}

var copyEncodingInvoker *gi.Function

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
		encodingFreeFunction = gi.FunctionInvokerNew("GtkSource", "free")
	})
}

var freeEncodingInvoker *gi.Function

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
		encodingGetCharsetFunction = gi.FunctionInvokerNew("GtkSource", "get_charset")
	})
}

var getCharsetEncodingInvoker *gi.Function

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
		encodingGetNameFunction = gi.FunctionInvokerNew("GtkSource", "get_name")
	})
}

var getNameEncodingInvoker *gi.Function

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
		encodingToStringFunction = gi.FunctionInvokerNew("GtkSource", "to_string")
	})
}

var toStringEncodingInvoker *gi.Function

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
var fileClassStructOnce sync.Once

func fileClassStructSet() {
	fileClassStructOnce.Do(func() {
		fileClassStruct = gi.StructNew("GtkSource", "FileClass")
	})
}

type FileClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileLoaderClassStruct *gi.Struct
var fileLoaderClassStructOnce sync.Once

func fileLoaderClassStructSet() {
	fileLoaderClassStructOnce.Do(func() {
		fileLoaderClassStruct = gi.StructNew("GtkSource", "FileLoaderClass")
	})
}

type FileLoaderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileLoaderPrivateStruct *gi.Struct
var fileLoaderPrivateStructOnce sync.Once

func fileLoaderPrivateStructSet() {
	fileLoaderPrivateStructOnce.Do(func() {
		fileLoaderPrivateStruct = gi.StructNew("GtkSource", "FileLoaderPrivate")
	})
}

type FileLoaderPrivate struct {
	native uintptr
}

var filePrivateStruct *gi.Struct
var filePrivateStructOnce sync.Once

func filePrivateStructSet() {
	filePrivateStructOnce.Do(func() {
		filePrivateStruct = gi.StructNew("GtkSource", "FilePrivate")
	})
}

type FilePrivate struct {
	native uintptr
}

var fileSaverClassStruct *gi.Struct
var fileSaverClassStructOnce sync.Once

func fileSaverClassStructSet() {
	fileSaverClassStructOnce.Do(func() {
		fileSaverClassStruct = gi.StructNew("GtkSource", "FileSaverClass")
	})
}

type FileSaverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var fileSaverPrivateStruct *gi.Struct
var fileSaverPrivateStructOnce sync.Once

func fileSaverPrivateStructSet() {
	fileSaverPrivateStructOnce.Do(func() {
		fileSaverPrivateStruct = gi.StructNew("GtkSource", "FileSaverPrivate")
	})
}

type FileSaverPrivate struct {
	native uintptr
}

var gutterClassStruct *gi.Struct
var gutterClassStructOnce sync.Once

func gutterClassStructSet() {
	gutterClassStructOnce.Do(func() {
		gutterClassStruct = gi.StructNew("GtkSource", "GutterClass")
	})
}

type GutterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var gutterPrivateStruct *gi.Struct
var gutterPrivateStructOnce sync.Once

func gutterPrivateStructSet() {
	gutterPrivateStructOnce.Do(func() {
		gutterPrivateStruct = gi.StructNew("GtkSource", "GutterPrivate")
	})
}

type GutterPrivate struct {
	native uintptr
}

var gutterRendererClassStruct *gi.Struct
var gutterRendererClassStructOnce sync.Once

func gutterRendererClassStructSet() {
	gutterRendererClassStructOnce.Do(func() {
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
var gutterRendererPixbufClassStructOnce sync.Once

func gutterRendererPixbufClassStructSet() {
	gutterRendererPixbufClassStructOnce.Do(func() {
		gutterRendererPixbufClassStruct = gi.StructNew("GtkSource", "GutterRendererPixbufClass")
	})
}

type GutterRendererPixbufClass struct {
	native uintptr
}

var gutterRendererPixbufPrivateStruct *gi.Struct
var gutterRendererPixbufPrivateStructOnce sync.Once

func gutterRendererPixbufPrivateStructSet() {
	gutterRendererPixbufPrivateStructOnce.Do(func() {
		gutterRendererPixbufPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPixbufPrivate")
	})
}

type GutterRendererPixbufPrivate struct {
	native uintptr
}

var gutterRendererPrivateStruct *gi.Struct
var gutterRendererPrivateStructOnce sync.Once

func gutterRendererPrivateStructSet() {
	gutterRendererPrivateStructOnce.Do(func() {
		gutterRendererPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPrivate")
	})
}

type GutterRendererPrivate struct {
	native uintptr
}

var gutterRendererTextClassStruct *gi.Struct
var gutterRendererTextClassStructOnce sync.Once

func gutterRendererTextClassStructSet() {
	gutterRendererTextClassStructOnce.Do(func() {
		gutterRendererTextClassStruct = gi.StructNew("GtkSource", "GutterRendererTextClass")
	})
}

type GutterRendererTextClass struct {
	native uintptr
}

var gutterRendererTextPrivateStruct *gi.Struct
var gutterRendererTextPrivateStructOnce sync.Once

func gutterRendererTextPrivateStructSet() {
	gutterRendererTextPrivateStructOnce.Do(func() {
		gutterRendererTextPrivateStruct = gi.StructNew("GtkSource", "GutterRendererTextPrivate")
	})
}

type GutterRendererTextPrivate struct {
	native uintptr
}

var languageClassStruct *gi.Struct
var languageClassStructOnce sync.Once

func languageClassStructSet() {
	languageClassStructOnce.Do(func() {
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
var languageManagerClassStructOnce sync.Once

func languageManagerClassStructSet() {
	languageManagerClassStructOnce.Do(func() {
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
var languageManagerPrivateStructOnce sync.Once

func languageManagerPrivateStructSet() {
	languageManagerPrivateStructOnce.Do(func() {
		languageManagerPrivateStruct = gi.StructNew("GtkSource", "LanguageManagerPrivate")
	})
}

type LanguageManagerPrivate struct {
	native uintptr
}

var languagePrivateStruct *gi.Struct
var languagePrivateStructOnce sync.Once

func languagePrivateStructSet() {
	languagePrivateStructOnce.Do(func() {
		languagePrivateStruct = gi.StructNew("GtkSource", "LanguagePrivate")
	})
}

type LanguagePrivate struct {
	native uintptr
}

var mapClassStruct *gi.Struct
var mapClassStructOnce sync.Once

func mapClassStructSet() {
	mapClassStructOnce.Do(func() {
		mapClassStruct = gi.StructNew("GtkSource", "MapClass")
	})
}

type MapClass struct {
	native      uintptr
	ParentClass *ViewClass
	// UNSUPPORTED : C value 'padding' : missing Type
}

var markAttributesClassStruct *gi.Struct
var markAttributesClassStructOnce sync.Once

func markAttributesClassStructSet() {
	markAttributesClassStructOnce.Do(func() {
		markAttributesClassStruct = gi.StructNew("GtkSource", "MarkAttributesClass")
	})
}

type MarkAttributesClass struct {
	native uintptr
}

var markAttributesPrivateStruct *gi.Struct
var markAttributesPrivateStructOnce sync.Once

func markAttributesPrivateStructSet() {
	markAttributesPrivateStructOnce.Do(func() {
		markAttributesPrivateStruct = gi.StructNew("GtkSource", "MarkAttributesPrivate")
	})
}

type MarkAttributesPrivate struct {
	native uintptr
}

var markClassStruct *gi.Struct
var markClassStructOnce sync.Once

func markClassStructSet() {
	markClassStructOnce.Do(func() {
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
var markPrivateStructOnce sync.Once

func markPrivateStructSet() {
	markPrivateStructOnce.Do(func() {
		markPrivateStruct = gi.StructNew("GtkSource", "MarkPrivate")
	})
}

type MarkPrivate struct {
	native uintptr
}

var printCompositorClassStruct *gi.Struct
var printCompositorClassStructOnce sync.Once

func printCompositorClassStructSet() {
	printCompositorClassStructOnce.Do(func() {
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
var printCompositorPrivateStructOnce sync.Once

func printCompositorPrivateStructSet() {
	printCompositorPrivateStructOnce.Do(func() {
		printCompositorPrivateStruct = gi.StructNew("GtkSource", "PrintCompositorPrivate")
	})
}

type PrintCompositorPrivate struct {
	native uintptr
}

var regionClassStruct *gi.Struct
var regionClassStructOnce sync.Once

func regionClassStructSet() {
	regionClassStructOnce.Do(func() {
		regionClassStruct = gi.StructNew("GtkSource", "RegionClass")
	})
}

type RegionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var regionIterStruct *gi.Struct
var regionIterStructOnce sync.Once

func regionIterStructSet() {
	regionIterStructOnce.Do(func() {
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
var searchContextClassStructOnce sync.Once

func searchContextClassStructSet() {
	searchContextClassStructOnce.Do(func() {
		searchContextClassStruct = gi.StructNew("GtkSource", "SearchContextClass")
	})
}

type SearchContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var searchContextPrivateStruct *gi.Struct
var searchContextPrivateStructOnce sync.Once

func searchContextPrivateStructSet() {
	searchContextPrivateStructOnce.Do(func() {
		searchContextPrivateStruct = gi.StructNew("GtkSource", "SearchContextPrivate")
	})
}

type SearchContextPrivate struct {
	native uintptr
}

var searchSettingsClassStruct *gi.Struct
var searchSettingsClassStructOnce sync.Once

func searchSettingsClassStructSet() {
	searchSettingsClassStructOnce.Do(func() {
		searchSettingsClassStruct = gi.StructNew("GtkSource", "SearchSettingsClass")
	})
}

type SearchSettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var searchSettingsPrivateStruct *gi.Struct
var searchSettingsPrivateStructOnce sync.Once

func searchSettingsPrivateStructSet() {
	searchSettingsPrivateStructOnce.Do(func() {
		searchSettingsPrivateStruct = gi.StructNew("GtkSource", "SearchSettingsPrivate")
	})
}

type SearchSettingsPrivate struct {
	native uintptr
}

var spaceDrawerClassStruct *gi.Struct
var spaceDrawerClassStructOnce sync.Once

func spaceDrawerClassStructSet() {
	spaceDrawerClassStructOnce.Do(func() {
		spaceDrawerClassStruct = gi.StructNew("GtkSource", "SpaceDrawerClass")
	})
}

type SpaceDrawerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var spaceDrawerPrivateStruct *gi.Struct
var spaceDrawerPrivateStructOnce sync.Once

func spaceDrawerPrivateStructSet() {
	spaceDrawerPrivateStructOnce.Do(func() {
		spaceDrawerPrivateStruct = gi.StructNew("GtkSource", "SpaceDrawerPrivate")
	})
}

type SpaceDrawerPrivate struct {
	native uintptr
}

var styleClassStruct *gi.Struct
var styleClassStructOnce sync.Once

func styleClassStructSet() {
	styleClassStructOnce.Do(func() {
		styleClassStruct = gi.StructNew("GtkSource", "StyleClass")
	})
}

type StyleClass struct {
	native uintptr
}

var styleSchemeChooserButtonClassStruct *gi.Struct
var styleSchemeChooserButtonClassStructOnce sync.Once

func styleSchemeChooserButtonClassStructSet() {
	styleSchemeChooserButtonClassStructOnce.Do(func() {
		styleSchemeChooserButtonClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserButtonClass")
	})
}

type StyleSchemeChooserButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.ButtonClass'
}

var styleSchemeChooserInterfaceStruct *gi.Struct
var styleSchemeChooserInterfaceStructOnce sync.Once

func styleSchemeChooserInterfaceStructSet() {
	styleSchemeChooserInterfaceStructOnce.Do(func() {
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
var styleSchemeChooserWidgetClassStructOnce sync.Once

func styleSchemeChooserWidgetClassStructSet() {
	styleSchemeChooserWidgetClassStructOnce.Do(func() {
		styleSchemeChooserWidgetClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserWidgetClass")
	})
}

type StyleSchemeChooserWidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.BinClass'
}

var styleSchemeClassStruct *gi.Struct
var styleSchemeClassStructOnce sync.Once

func styleSchemeClassStructSet() {
	styleSchemeClassStructOnce.Do(func() {
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
var styleSchemeManagerClassStructOnce sync.Once

func styleSchemeManagerClassStructSet() {
	styleSchemeManagerClassStructOnce.Do(func() {
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
var styleSchemeManagerPrivateStructOnce sync.Once

func styleSchemeManagerPrivateStructSet() {
	styleSchemeManagerPrivateStructOnce.Do(func() {
		styleSchemeManagerPrivateStruct = gi.StructNew("GtkSource", "StyleSchemeManagerPrivate")
	})
}

type StyleSchemeManagerPrivate struct {
	native uintptr
}

var styleSchemePrivateStruct *gi.Struct
var styleSchemePrivateStructOnce sync.Once

func styleSchemePrivateStructSet() {
	styleSchemePrivateStructOnce.Do(func() {
		styleSchemePrivateStruct = gi.StructNew("GtkSource", "StyleSchemePrivate")
	})
}

type StyleSchemePrivate struct {
	native uintptr
}

var tagClassStruct *gi.Struct
var tagClassStructOnce sync.Once

func tagClassStructSet() {
	tagClassStructOnce.Do(func() {
		tagClassStruct = gi.StructNew("GtkSource", "TagClass")
	})
}

type TagClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextTagClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var undoManagerIfaceStruct *gi.Struct
var undoManagerIfaceStructOnce sync.Once

func undoManagerIfaceStructSet() {
	undoManagerIfaceStructOnce.Do(func() {
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
var viewClassStructOnce sync.Once

func viewClassStructSet() {
	viewClassStructOnce.Do(func() {
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
var viewPrivateStructOnce sync.Once

func viewPrivateStructSet() {
	viewPrivateStructOnce.Do(func() {
		viewPrivateStruct = gi.StructNew("GtkSource", "ViewPrivate")
	})
}

type ViewPrivate struct {
	native uintptr
}
