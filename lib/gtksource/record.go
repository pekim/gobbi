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
