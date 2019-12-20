package generate

var ClassBlacklist = map[string]bool{
	"GdkPixbuf.PixbufSimpleAnimIter": true,

	"Gtk.EntryIconAccessible":   true,
	"Gtk.EventControllerKey":    true,
	"Gtk.EventControllerMotion": true,
	"Gtk.EventControllerScroll": true,
	"Gtk.GestureStylus":         true,
	"Gtk.HeaderBarAccessible":   true,
	"Gtk.StackAccessible":       true,
}

var classVersionAddenda = map[string]string{}

func (c *Class) applyAddenda() {
	if _, ok := ClassBlacklist[c.namespace.Name+"."+c.Name]; ok {
		c.blacklist = true
	}

	if version, ok := classVersionAddenda[c.namespace.Name+"."+c.Name]; ok {
		c.Version = version
	}
}
