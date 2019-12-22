package generate

var functionBlacklist = map[string]bool{
	// settings backend implementations
	"Gio.settings_backend_get_default":           true,
	"Gio.settings_backend_changed":               true,
	"Gio.settings_backend_path_changed":          true,
	"Gio.settings_backend_keys_changed":          true,
	"Gio.settings_backend_path_writable_changed": true,
	"Gio.settings_backend_writable_changed":      true,
	"Gio.settings_backend_changed_tree":          true,
	"Gio.settings_backend_flatten_tree":          true,
	"Gio.keyfile_settings_backend_new":           true,
	"Gio.memory_settings_backend_new":            true,
	"Gio.null_settings_backend_new":              true,

	"GLib.unix_error_quark": true,

	"Pango.module_register":          true,
	"Pango.get_lib_subdirectory":     true,
	"Pango.get_sysconf_subdirectory": true,
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

	if _, ok := functionBlacklist[f.namespace.Name+"."+f.Name]; ok {
		f.blacklist = true
	}
}
