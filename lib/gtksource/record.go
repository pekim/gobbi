// Code generated - DO NOT EDIT.

package gtksource

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var BufferClassStruct *gi.Struct
var BufferClassStructOnce sync.Once

func BufferClassStructSet() {
	BufferClassStructOnce.Do(func() {
		BufferClassStruct = gi.StructNew("GtkSource", "BufferClass")
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

var BufferPrivateStruct *gi.Struct
var BufferPrivateStructOnce sync.Once

func BufferPrivateStructSet() {
	BufferPrivateStructOnce.Do(func() {
		BufferPrivateStruct = gi.StructNew("GtkSource", "BufferPrivate")
	})
}

type BufferPrivate struct {
	native uintptr
}

var CompletionClassStruct *gi.Struct
var CompletionClassStructOnce sync.Once

func CompletionClassStructSet() {
	CompletionClassStructOnce.Do(func() {
		CompletionClassStruct = gi.StructNew("GtkSource", "CompletionClass")
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

var CompletionContextClassStruct *gi.Struct
var CompletionContextClassStructOnce sync.Once

func CompletionContextClassStructSet() {
	CompletionContextClassStructOnce.Do(func() {
		CompletionContextClassStruct = gi.StructNew("GtkSource", "CompletionContextClass")
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

var CompletionContextPrivateStruct *gi.Struct
var CompletionContextPrivateStructOnce sync.Once

func CompletionContextPrivateStructSet() {
	CompletionContextPrivateStructOnce.Do(func() {
		CompletionContextPrivateStruct = gi.StructNew("GtkSource", "CompletionContextPrivate")
	})
}

type CompletionContextPrivate struct {
	native uintptr
}

var CompletionInfoClassStruct *gi.Struct
var CompletionInfoClassStructOnce sync.Once

func CompletionInfoClassStructSet() {
	CompletionInfoClassStructOnce.Do(func() {
		CompletionInfoClassStruct = gi.StructNew("GtkSource", "CompletionInfoClass")
	})
}

type CompletionInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.WindowClass'
	// UNSUPPORTED : C value 'before_show' : missing Type
}

var CompletionInfoPrivateStruct *gi.Struct
var CompletionInfoPrivateStructOnce sync.Once

func CompletionInfoPrivateStructSet() {
	CompletionInfoPrivateStructOnce.Do(func() {
		CompletionInfoPrivateStruct = gi.StructNew("GtkSource", "CompletionInfoPrivate")
	})
}

type CompletionInfoPrivate struct {
	native uintptr
}

var CompletionItemClassStruct *gi.Struct
var CompletionItemClassStructOnce sync.Once

func CompletionItemClassStructSet() {
	CompletionItemClassStructOnce.Do(func() {
		CompletionItemClassStruct = gi.StructNew("GtkSource", "CompletionItemClass")
	})
}

type CompletionItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var CompletionItemPrivateStruct *gi.Struct
var CompletionItemPrivateStructOnce sync.Once

func CompletionItemPrivateStructSet() {
	CompletionItemPrivateStructOnce.Do(func() {
		CompletionItemPrivateStruct = gi.StructNew("GtkSource", "CompletionItemPrivate")
	})
}

type CompletionItemPrivate struct {
	native uintptr
}

var CompletionPrivateStruct *gi.Struct
var CompletionPrivateStructOnce sync.Once

func CompletionPrivateStructSet() {
	CompletionPrivateStructOnce.Do(func() {
		CompletionPrivateStruct = gi.StructNew("GtkSource", "CompletionPrivate")
	})
}

type CompletionPrivate struct {
	native uintptr
}

var CompletionProposalIfaceStruct *gi.Struct
var CompletionProposalIfaceStructOnce sync.Once

func CompletionProposalIfaceStructSet() {
	CompletionProposalIfaceStructOnce.Do(func() {
		CompletionProposalIfaceStruct = gi.StructNew("GtkSource", "CompletionProposalIface")
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

var CompletionProviderIfaceStruct *gi.Struct
var CompletionProviderIfaceStructOnce sync.Once

func CompletionProviderIfaceStructSet() {
	CompletionProviderIfaceStructOnce.Do(func() {
		CompletionProviderIfaceStruct = gi.StructNew("GtkSource", "CompletionProviderIface")
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

var CompletionWordsClassStruct *gi.Struct
var CompletionWordsClassStructOnce sync.Once

func CompletionWordsClassStructSet() {
	CompletionWordsClassStructOnce.Do(func() {
		CompletionWordsClassStruct = gi.StructNew("GtkSource", "CompletionWordsClass")
	})
}

type CompletionWordsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var CompletionWordsPrivateStruct *gi.Struct
var CompletionWordsPrivateStructOnce sync.Once

func CompletionWordsPrivateStructSet() {
	CompletionWordsPrivateStructOnce.Do(func() {
		CompletionWordsPrivateStruct = gi.StructNew("GtkSource", "CompletionWordsPrivate")
	})
}

type CompletionWordsPrivate struct {
	native uintptr
}

var EncodingStruct *gi.Struct
var EncodingStructOnce sync.Once

func EncodingStructSet() {
	EncodingStructOnce.Do(func() {
		EncodingStruct = gi.StructNew("GtkSource", "Encoding")
	})
}

type Encoding struct {
	native uintptr
}

var copyEncodingInvoker *gi.Function

// Copy is a representation of the C type gtk_source_encoding_copy.
func (recv *Encoding) Copy() *Encoding {
	if copyEncodingInvoker == nil {
		copyEncodingInvoker = gi.StructFunctionInvokerNew("GtkSource", "Encoding", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyEncodingInvoker.Invoke(inArgs[:], nil)

	retGo := &Encoding{native: ret.Pointer()}

	return retGo
}

var freeEncodingInvoker *gi.Function

// Free is a representation of the C type gtk_source_encoding_free.
func (recv *Encoding) Free() {
	if freeEncodingInvoker == nil {
		freeEncodingInvoker = gi.StructFunctionInvokerNew("GtkSource", "Encoding", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeEncodingInvoker.Invoke(inArgs[:], nil)

}

var getCharsetEncodingInvoker *gi.Function

// GetCharset is a representation of the C type gtk_source_encoding_get_charset.
func (recv *Encoding) GetCharset() string {
	if getCharsetEncodingInvoker == nil {
		getCharsetEncodingInvoker = gi.StructFunctionInvokerNew("GtkSource", "Encoding", "get_charset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getCharsetEncodingInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getNameEncodingInvoker *gi.Function

// GetName is a representation of the C type gtk_source_encoding_get_name.
func (recv *Encoding) GetName() string {
	if getNameEncodingInvoker == nil {
		getNameEncodingInvoker = gi.StructFunctionInvokerNew("GtkSource", "Encoding", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameEncodingInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var toStringEncodingInvoker *gi.Function

// ToString is a representation of the C type gtk_source_encoding_to_string.
func (recv *Encoding) ToString() string {
	if toStringEncodingInvoker == nil {
		toStringEncodingInvoker = gi.StructFunctionInvokerNew("GtkSource", "Encoding", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringEncodingInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var FileClassStruct *gi.Struct
var FileClassStructOnce sync.Once

func FileClassStructSet() {
	FileClassStructOnce.Do(func() {
		FileClassStruct = gi.StructNew("GtkSource", "FileClass")
	})
}

type FileClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var FileLoaderClassStruct *gi.Struct
var FileLoaderClassStructOnce sync.Once

func FileLoaderClassStructSet() {
	FileLoaderClassStructOnce.Do(func() {
		FileLoaderClassStruct = gi.StructNew("GtkSource", "FileLoaderClass")
	})
}

type FileLoaderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var FileLoaderPrivateStruct *gi.Struct
var FileLoaderPrivateStructOnce sync.Once

func FileLoaderPrivateStructSet() {
	FileLoaderPrivateStructOnce.Do(func() {
		FileLoaderPrivateStruct = gi.StructNew("GtkSource", "FileLoaderPrivate")
	})
}

type FileLoaderPrivate struct {
	native uintptr
}

var FilePrivateStruct *gi.Struct
var FilePrivateStructOnce sync.Once

func FilePrivateStructSet() {
	FilePrivateStructOnce.Do(func() {
		FilePrivateStruct = gi.StructNew("GtkSource", "FilePrivate")
	})
}

type FilePrivate struct {
	native uintptr
}

var FileSaverClassStruct *gi.Struct
var FileSaverClassStructOnce sync.Once

func FileSaverClassStructSet() {
	FileSaverClassStructOnce.Do(func() {
		FileSaverClassStruct = gi.StructNew("GtkSource", "FileSaverClass")
	})
}

type FileSaverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var FileSaverPrivateStruct *gi.Struct
var FileSaverPrivateStructOnce sync.Once

func FileSaverPrivateStructSet() {
	FileSaverPrivateStructOnce.Do(func() {
		FileSaverPrivateStruct = gi.StructNew("GtkSource", "FileSaverPrivate")
	})
}

type FileSaverPrivate struct {
	native uintptr
}

var GutterClassStruct *gi.Struct
var GutterClassStructOnce sync.Once

func GutterClassStructSet() {
	GutterClassStructOnce.Do(func() {
		GutterClassStruct = gi.StructNew("GtkSource", "GutterClass")
	})
}

type GutterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var GutterPrivateStruct *gi.Struct
var GutterPrivateStructOnce sync.Once

func GutterPrivateStructSet() {
	GutterPrivateStructOnce.Do(func() {
		GutterPrivateStruct = gi.StructNew("GtkSource", "GutterPrivate")
	})
}

type GutterPrivate struct {
	native uintptr
}

var GutterRendererClassStruct *gi.Struct
var GutterRendererClassStructOnce sync.Once

func GutterRendererClassStructSet() {
	GutterRendererClassStructOnce.Do(func() {
		GutterRendererClassStruct = gi.StructNew("GtkSource", "GutterRendererClass")
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

var GutterRendererPixbufClassStruct *gi.Struct
var GutterRendererPixbufClassStructOnce sync.Once

func GutterRendererPixbufClassStructSet() {
	GutterRendererPixbufClassStructOnce.Do(func() {
		GutterRendererPixbufClassStruct = gi.StructNew("GtkSource", "GutterRendererPixbufClass")
	})
}

type GutterRendererPixbufClass struct {
	native uintptr
}

var GutterRendererPixbufPrivateStruct *gi.Struct
var GutterRendererPixbufPrivateStructOnce sync.Once

func GutterRendererPixbufPrivateStructSet() {
	GutterRendererPixbufPrivateStructOnce.Do(func() {
		GutterRendererPixbufPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPixbufPrivate")
	})
}

type GutterRendererPixbufPrivate struct {
	native uintptr
}

var GutterRendererPrivateStruct *gi.Struct
var GutterRendererPrivateStructOnce sync.Once

func GutterRendererPrivateStructSet() {
	GutterRendererPrivateStructOnce.Do(func() {
		GutterRendererPrivateStruct = gi.StructNew("GtkSource", "GutterRendererPrivate")
	})
}

type GutterRendererPrivate struct {
	native uintptr
}

var GutterRendererTextClassStruct *gi.Struct
var GutterRendererTextClassStructOnce sync.Once

func GutterRendererTextClassStructSet() {
	GutterRendererTextClassStructOnce.Do(func() {
		GutterRendererTextClassStruct = gi.StructNew("GtkSource", "GutterRendererTextClass")
	})
}

type GutterRendererTextClass struct {
	native uintptr
}

var GutterRendererTextPrivateStruct *gi.Struct
var GutterRendererTextPrivateStructOnce sync.Once

func GutterRendererTextPrivateStructSet() {
	GutterRendererTextPrivateStructOnce.Do(func() {
		GutterRendererTextPrivateStruct = gi.StructNew("GtkSource", "GutterRendererTextPrivate")
	})
}

type GutterRendererTextPrivate struct {
	native uintptr
}

var LanguageClassStruct *gi.Struct
var LanguageClassStructOnce sync.Once

func LanguageClassStructSet() {
	LanguageClassStructOnce.Do(func() {
		LanguageClassStruct = gi.StructNew("GtkSource", "LanguageClass")
	})
}

type LanguageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var LanguageManagerClassStruct *gi.Struct
var LanguageManagerClassStructOnce sync.Once

func LanguageManagerClassStructSet() {
	LanguageManagerClassStructOnce.Do(func() {
		LanguageManagerClassStruct = gi.StructNew("GtkSource", "LanguageManagerClass")
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

var LanguageManagerPrivateStruct *gi.Struct
var LanguageManagerPrivateStructOnce sync.Once

func LanguageManagerPrivateStructSet() {
	LanguageManagerPrivateStructOnce.Do(func() {
		LanguageManagerPrivateStruct = gi.StructNew("GtkSource", "LanguageManagerPrivate")
	})
}

type LanguageManagerPrivate struct {
	native uintptr
}

var LanguagePrivateStruct *gi.Struct
var LanguagePrivateStructOnce sync.Once

func LanguagePrivateStructSet() {
	LanguagePrivateStructOnce.Do(func() {
		LanguagePrivateStruct = gi.StructNew("GtkSource", "LanguagePrivate")
	})
}

type LanguagePrivate struct {
	native uintptr
}

var MapClassStruct *gi.Struct
var MapClassStructOnce sync.Once

func MapClassStructSet() {
	MapClassStructOnce.Do(func() {
		MapClassStruct = gi.StructNew("GtkSource", "MapClass")
	})
}

type MapClass struct {
	native      uintptr
	ParentClass *ViewClass
	// UNSUPPORTED : C value 'padding' : missing Type
}

var MarkAttributesClassStruct *gi.Struct
var MarkAttributesClassStructOnce sync.Once

func MarkAttributesClassStructSet() {
	MarkAttributesClassStructOnce.Do(func() {
		MarkAttributesClassStruct = gi.StructNew("GtkSource", "MarkAttributesClass")
	})
}

type MarkAttributesClass struct {
	native uintptr
}

var MarkAttributesPrivateStruct *gi.Struct
var MarkAttributesPrivateStructOnce sync.Once

func MarkAttributesPrivateStructSet() {
	MarkAttributesPrivateStructOnce.Do(func() {
		MarkAttributesPrivateStruct = gi.StructNew("GtkSource", "MarkAttributesPrivate")
	})
}

type MarkAttributesPrivate struct {
	native uintptr
}

var MarkClassStruct *gi.Struct
var MarkClassStructOnce sync.Once

func MarkClassStructSet() {
	MarkClassStructOnce.Do(func() {
		MarkClassStruct = gi.StructNew("GtkSource", "MarkClass")
	})
}

type MarkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextMarkClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var MarkPrivateStruct *gi.Struct
var MarkPrivateStructOnce sync.Once

func MarkPrivateStructSet() {
	MarkPrivateStructOnce.Do(func() {
		MarkPrivateStruct = gi.StructNew("GtkSource", "MarkPrivate")
	})
}

type MarkPrivate struct {
	native uintptr
}

var PrintCompositorClassStruct *gi.Struct
var PrintCompositorClassStructOnce sync.Once

func PrintCompositorClassStructSet() {
	PrintCompositorClassStructOnce.Do(func() {
		PrintCompositorClassStruct = gi.StructNew("GtkSource", "PrintCompositorClass")
	})
}

type PrintCompositorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var PrintCompositorPrivateStruct *gi.Struct
var PrintCompositorPrivateStructOnce sync.Once

func PrintCompositorPrivateStructSet() {
	PrintCompositorPrivateStructOnce.Do(func() {
		PrintCompositorPrivateStruct = gi.StructNew("GtkSource", "PrintCompositorPrivate")
	})
}

type PrintCompositorPrivate struct {
	native uintptr
}

var RegionClassStruct *gi.Struct
var RegionClassStructOnce sync.Once

func RegionClassStructSet() {
	RegionClassStructOnce.Do(func() {
		RegionClassStruct = gi.StructNew("GtkSource", "RegionClass")
	})
}

type RegionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var RegionIterStruct *gi.Struct
var RegionIterStructOnce sync.Once

func RegionIterStructSet() {
	RegionIterStructOnce.Do(func() {
		RegionIterStruct = gi.StructNew("GtkSource", "RegionIter")
	})
}

type RegionIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_source_region_iter_get_subregion' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_is_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_next' : return type 'gboolean' not supported

var SearchContextClassStruct *gi.Struct
var SearchContextClassStructOnce sync.Once

func SearchContextClassStructSet() {
	SearchContextClassStructOnce.Do(func() {
		SearchContextClassStruct = gi.StructNew("GtkSource", "SearchContextClass")
	})
}

type SearchContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var SearchContextPrivateStruct *gi.Struct
var SearchContextPrivateStructOnce sync.Once

func SearchContextPrivateStructSet() {
	SearchContextPrivateStructOnce.Do(func() {
		SearchContextPrivateStruct = gi.StructNew("GtkSource", "SearchContextPrivate")
	})
}

type SearchContextPrivate struct {
	native uintptr
}

var SearchSettingsClassStruct *gi.Struct
var SearchSettingsClassStructOnce sync.Once

func SearchSettingsClassStructSet() {
	SearchSettingsClassStructOnce.Do(func() {
		SearchSettingsClassStruct = gi.StructNew("GtkSource", "SearchSettingsClass")
	})
}

type SearchSettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var SearchSettingsPrivateStruct *gi.Struct
var SearchSettingsPrivateStructOnce sync.Once

func SearchSettingsPrivateStructSet() {
	SearchSettingsPrivateStructOnce.Do(func() {
		SearchSettingsPrivateStruct = gi.StructNew("GtkSource", "SearchSettingsPrivate")
	})
}

type SearchSettingsPrivate struct {
	native uintptr
}

var SpaceDrawerClassStruct *gi.Struct
var SpaceDrawerClassStructOnce sync.Once

func SpaceDrawerClassStructSet() {
	SpaceDrawerClassStructOnce.Do(func() {
		SpaceDrawerClassStruct = gi.StructNew("GtkSource", "SpaceDrawerClass")
	})
}

type SpaceDrawerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var SpaceDrawerPrivateStruct *gi.Struct
var SpaceDrawerPrivateStructOnce sync.Once

func SpaceDrawerPrivateStructSet() {
	SpaceDrawerPrivateStructOnce.Do(func() {
		SpaceDrawerPrivateStruct = gi.StructNew("GtkSource", "SpaceDrawerPrivate")
	})
}

type SpaceDrawerPrivate struct {
	native uintptr
}

var StyleClassStruct *gi.Struct
var StyleClassStructOnce sync.Once

func StyleClassStructSet() {
	StyleClassStructOnce.Do(func() {
		StyleClassStruct = gi.StructNew("GtkSource", "StyleClass")
	})
}

type StyleClass struct {
	native uintptr
}

var StyleSchemeChooserButtonClassStruct *gi.Struct
var StyleSchemeChooserButtonClassStructOnce sync.Once

func StyleSchemeChooserButtonClassStructSet() {
	StyleSchemeChooserButtonClassStructOnce.Do(func() {
		StyleSchemeChooserButtonClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserButtonClass")
	})
}

type StyleSchemeChooserButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.ButtonClass'
}

var StyleSchemeChooserInterfaceStruct *gi.Struct
var StyleSchemeChooserInterfaceStructOnce sync.Once

func StyleSchemeChooserInterfaceStructSet() {
	StyleSchemeChooserInterfaceStructOnce.Do(func() {
		StyleSchemeChooserInterfaceStruct = gi.StructNew("GtkSource", "StyleSchemeChooserInterface")
	})
}

type StyleSchemeChooserInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_style_scheme' : missing Type
	// UNSUPPORTED : C value 'set_style_scheme' : missing Type
	// UNSUPPORTED : C value 'padding' : missing Type
}

var StyleSchemeChooserWidgetClassStruct *gi.Struct
var StyleSchemeChooserWidgetClassStructOnce sync.Once

func StyleSchemeChooserWidgetClassStructSet() {
	StyleSchemeChooserWidgetClassStructOnce.Do(func() {
		StyleSchemeChooserWidgetClassStruct = gi.StructNew("GtkSource", "StyleSchemeChooserWidgetClass")
	})
}

type StyleSchemeChooserWidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.BinClass'
}

var StyleSchemeClassStruct *gi.Struct
var StyleSchemeClassStructOnce sync.Once

func StyleSchemeClassStructSet() {
	StyleSchemeClassStructOnce.Do(func() {
		StyleSchemeClassStruct = gi.StructNew("GtkSource", "StyleSchemeClass")
	})
}

type StyleSchemeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'base_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

var StyleSchemeManagerClassStruct *gi.Struct
var StyleSchemeManagerClassStructOnce sync.Once

func StyleSchemeManagerClassStructSet() {
	StyleSchemeManagerClassStructOnce.Do(func() {
		StyleSchemeManagerClassStruct = gi.StructNew("GtkSource", "StyleSchemeManagerClass")
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

var StyleSchemeManagerPrivateStruct *gi.Struct
var StyleSchemeManagerPrivateStructOnce sync.Once

func StyleSchemeManagerPrivateStructSet() {
	StyleSchemeManagerPrivateStructOnce.Do(func() {
		StyleSchemeManagerPrivateStruct = gi.StructNew("GtkSource", "StyleSchemeManagerPrivate")
	})
}

type StyleSchemeManagerPrivate struct {
	native uintptr
}

var StyleSchemePrivateStruct *gi.Struct
var StyleSchemePrivateStructOnce sync.Once

func StyleSchemePrivateStructSet() {
	StyleSchemePrivateStructOnce.Do(func() {
		StyleSchemePrivateStruct = gi.StructNew("GtkSource", "StyleSchemePrivate")
	})
}

type StyleSchemePrivate struct {
	native uintptr
}

var TagClassStruct *gi.Struct
var TagClassStructOnce sync.Once

func TagClassStructSet() {
	TagClassStructOnce.Do(func() {
		TagClassStruct = gi.StructNew("GtkSource", "TagClass")
	})
}

type TagClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextTagClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var UndoManagerIfaceStruct *gi.Struct
var UndoManagerIfaceStructOnce sync.Once

func UndoManagerIfaceStructSet() {
	UndoManagerIfaceStructOnce.Do(func() {
		UndoManagerIfaceStruct = gi.StructNew("GtkSource", "UndoManagerIface")
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

var ViewClassStruct *gi.Struct
var ViewClassStructOnce sync.Once

func ViewClassStructSet() {
	ViewClassStructOnce.Do(func() {
		ViewClassStruct = gi.StructNew("GtkSource", "ViewClass")
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

var ViewPrivateStruct *gi.Struct
var ViewPrivateStructOnce sync.Once

func ViewPrivateStructSet() {
	ViewPrivateStructOnce.Do(func() {
		ViewPrivateStruct = gi.StructNew("GtkSource", "ViewPrivate")
	})
}

type ViewPrivate struct {
	native uintptr
}
