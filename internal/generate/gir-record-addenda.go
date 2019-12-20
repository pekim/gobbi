package generate

var recordBlacklist = map[string]bool{
	"Gio.NativeSocketAddressClass":   true,
	"Gio.NativeSocketAddressPrivate": true,
	"Gio.SettingsBackendClass":       true,
	"Gio.SettingsBackendPrivate":     true,

	"GLib.TestLogMsg": true,

	"Gtk.EventControllerMotionClass": true,
	"Gtk.EventControllerScrollClass": true,
	"Gtk.GestureStylusClass":         true,
	"Gtk.HeaderBarAccessibleClass":   true,
	"Gtk.HeaderBarAccessiblePrivate": true,
	"Gtk.StackAccessibleClass":       true,
}

var recordVersionAddenda = map[string]string{
	"GLib.StatBuf":                "2.6",
	"Gtk.EventControllerKeyClass": "3.24",
}

func (r *Record) applyAddenda() {
	if _, ok := recordBlacklist[r.namespace.Name+"."+r.Name]; ok {
		r.blacklist = true
	}

	if version, ok := recordVersionAddenda[r.namespace.Name+"."+r.Name]; ok {
		r.Version = version
	}
}
