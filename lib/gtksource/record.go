// Code generated - DO NOT EDIT.

package gtksource

import gi "github.com/pekim/gobbi/internal/gi"

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

type BufferPrivate struct {
	native uintptr
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

type CompletionContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'
	// UNSUPPORTED : C value 'cancelled' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
}

type CompletionContextPrivate struct {
	native uintptr
}

type CompletionInfoClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.WindowClass'
	// UNSUPPORTED : C value 'before_show' : missing Type
}

type CompletionInfoPrivate struct {
	native uintptr
}

type CompletionItemClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

type CompletionItemPrivate struct {
	native uintptr
}

type CompletionPrivate struct {
	native uintptr
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

type CompletionWordsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

type CompletionWordsPrivate struct {
	native uintptr
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

type FileClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type FileLoaderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type FileLoaderPrivate struct {
	native uintptr
}

type FilePrivate struct {
	native uintptr
}

type FileSaverClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type FileSaverPrivate struct {
	native uintptr
}

type GutterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

type GutterPrivate struct {
	native uintptr
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

type GutterRendererPixbufClass struct {
	native uintptr
}

type GutterRendererPixbufPrivate struct {
	native uintptr
}

type GutterRendererPrivate struct {
	native uintptr
}

type GutterRendererTextClass struct {
	native uintptr
}

type GutterRendererTextPrivate struct {
	native uintptr
}

type LanguageClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

type LanguageManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved4' : missing Type
}

type LanguageManagerPrivate struct {
	native uintptr
}

type LanguagePrivate struct {
	native uintptr
}

type MapClass struct {
	native      uintptr
	ParentClass *ViewClass
	// UNSUPPORTED : C value 'padding' : missing Type
}

type MarkAttributesClass struct {
	native uintptr
}

type MarkAttributesPrivate struct {
	native uintptr
}

type MarkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextMarkClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

type MarkPrivate struct {
	native uintptr
}

type PrintCompositorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

type PrintCompositorPrivate struct {
	native uintptr
}

type RegionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type RegionIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_source_region_iter_get_subregion' : parameter 'start' of type 'Gtk.TextIter' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_is_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_source_region_iter_next' : return type 'gboolean' not supported

type SearchContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type SearchContextPrivate struct {
	native uintptr
}

type SearchSettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type SearchSettingsPrivate struct {
	native uintptr
}

type SpaceDrawerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

type SpaceDrawerPrivate struct {
	native uintptr
}

type StyleClass struct {
	native uintptr
}

type StyleSchemeChooserButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.ButtonClass'
}

type StyleSchemeChooserInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_interface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_style_scheme' : missing Type
	// UNSUPPORTED : C value 'set_style_scheme' : missing Type
	// UNSUPPORTED : C value 'padding' : missing Type
}

type StyleSchemeChooserWidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'Gtk.BinClass'
}

type StyleSchemeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'base_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
}

type StyleSchemeManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_source_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_source_reserved4' : missing Type
}

type StyleSchemeManagerPrivate struct {
	native uintptr
}

type StyleSchemePrivate struct {
	native uintptr
}

type TagClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gtk.TextTagClass'
	// UNSUPPORTED : C value 'padding' : missing Type
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

type ViewPrivate struct {
	native uintptr
}
