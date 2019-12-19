package generate

func (e *Enumeration) applyAddenda() {
	// internal to glib
	if e.namespace.Name == "GLib" && e.Name == "TestResult" {
		e.blacklist = true
	}

	if e.namespace.Name == "Gtk" && e.Name == "FontChooserLevel" {
		e.Version = "3.24"
	}
}
