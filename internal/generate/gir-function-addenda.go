package generate

var functionBlacklist = map[string]bool{
	// gio
	"g_io_module_query": true,

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

	// glib
	"g_unix_error_quark": true,

	// pango
	"pango_module_register":          true,
	"pango_get_lib_subdirectory":     true,
	"pango_get_sysconf_subdirectory": true,
	"pango_fontset_simple_new":       true,
	"pango_fontset_simple_append":    true,
	"pango_fontset_simple_size":      true,

	// xlib
	"XOpenDisplay": true,
}

func (f *Function) applyAddenda() {
	if f.namespace.Name == "cairo" {
		f.blacklist = true
		return
	}

	if f.namespace.Name == "GLib" && f.Name == "clear_error" {
		f.Parameters = append(f.Parameters, &Parameter{
			Name: "err",
			Argument: Argument{
				Type: &Type{
					Name:  "Error",
					CType: "GError**",
				},
			},
		})
		return
	}

	if _, ok := functionBlacklist[f.CIdentifier]; ok {
		f.blacklist = true
	}
}
