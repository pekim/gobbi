package generate

func (e *Enumeration) applyAddenda() {
	if e.namespace.Name == "Gtk" && e.Name == "FontChooserLevel" {
		e.Version = "3.24"
	}
}
