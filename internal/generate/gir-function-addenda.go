package generate

import "strings"

var functionBlacklist = map[string]bool{
	// atk

	// return *glib.SList
	"atk_text_get_default_attributes": true,
	"atk_text_get_run_attributes":     true,

	// gdk
	"gdk_synthesize_window_state": true,
	"gdk_window_destroy_notify":   true,

	// gdkpixbuf
	"gdk_pixbuf_save_to_buffer": true,

	// gio
	"g_io_module_query":                     true,
	"g_settings_schema_source_list_schemas": true,
	"g_input_stream_read":                   true,
	"g_input_stream_read_all":               true,
	"g_file_load_contents":                  true,
	"g_file_load_contents_finish":           true,
	"g_file_load_partial_contents_finish":   true,

	// settings backend implementations
	"g_settings_backend_get_default":           true,
	"g_settings_backend_changed":               true,
	"g_settings_backend_path_changed":          true,
	"g_settings_backend_keys_changed":          true,
	"g_settings_backend_path_writable_changed": true,
	"g_settings_backend_writable_changed":      true,
	"g_settings_backend_changed_tree":          true,
	"g_settings_backend_flatten_tree":          true,
	"g_keyfile_settings_backend_new":           true,
	"g_memory_settings_backend_new":            true,
	"g_null_settings_backend_new":              true,

	// gobject
	"g_signal_set_va_marshaller": true,

	// gtk
	"gtk_cell_accessible_parent_get_cell_position": true,
	"gtk_event_controller_key_forward":             true,
	"gtk_event_controller_key_get_group":           true,
	"gtk_event_controller_key_new":                 true,
	"gtk_event_controller_key_set_im_context":      true,

	// pango
	"pango_module_register":                true,
	"pango_get_lib_subdirectory":           true,
	"pango_get_sysconf_subdirectory":       true,
	"pango_fontset_simple_new":             true,
	"pango_fontset_simple_append":          true,
	"pango_fontset_simple_size":            true,
	"pango_config_key_get":                 true,
	"pango_config_key_get_system":          true,
	"pango_default_break":                  true,
	"pango_find_map":                       true,
	"pango_font_map_get_shape_engine_type": true,
	"pango_lookup_aliases":                 true,
	"pango_font_metrics_new":               true,
	"pango_map_get_engine":                 true,
	"pango_map_get_engines":                true,
	"pango_coverage_to_bytes":              true,

	// xlib
	"XOpenDisplay": true,
}

func (f *Function) applyAddenda() {
	// Parameters with a type of ResponseType do not match the signature of functions
	// of that accept the parameters.
	for _, param := range f.Parameters {
		if param.Type != nil && param.Type.Name == "ResponseType" {
			param.Type.CType = "gint"
		}
	}

	if f.namespace.Name == "cairo" {
		f.blacklist = true
		return
	}

	if strings.HasPrefix(f.CIdentifier, "g_io_module") {
		f.blacklist = true
	}

	// A small number of glib will be manually implemented, but in general
	// glib functions are not needed.
	if f.namespace.Name == "GLib" {
		f.blacklist = true
	}

	if _, ok := functionBlacklist[f.CIdentifier]; ok {
		f.blacklist = true
	}
}
