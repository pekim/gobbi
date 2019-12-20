package generate

var enumerationBlacklist = map[string]bool{
	// internal to glib
	"GLib.TestResult": true,
}

var enumerationVersionAddenda = map[string]string{
	"Gtk.FontChooserLevel": "3.24",
}

func (e *Enumeration) applyAddenda() {
	if _, ok := enumerationBlacklist[e.namespace.Name+"."+e.Name]; ok {
		e.blacklist = true
	}

	if version, ok := enumerationVersionAddenda[e.namespace.Name+"."+e.Name]; ok {
		e.Version = version
	}
}
