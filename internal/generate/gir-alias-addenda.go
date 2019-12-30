package generate

var aliasBlacklist = map[string]bool{
	"GLib.MutexLocker":        true,
	"GLib.RecMutexLocker":     true,
	"GLib.RWLockReaderLocker": true,
	"GLib.RWLockWriterLocker": true,

	"Gdk.XEvent": true,

	"GObject.SignalCMarshaller":   true,
	"GObject.SignalCVaMarshaller": true,
}

var aliasVersionAddenda = map[string]string{
	"GLib.RecMutexLocker":     "2.60",
	"GLib.RefString":          "2.58",
	"GLib.RWLockReaderLocker": "2.62",
	"GLib.RWLockWriterLocker": "2.62",
	//"GLib.Type":               "",
}

func (a *Alias) applyAddenda() {
	if _, ok := aliasBlacklist[a.namespace.Name+"."+a.Name]; ok {
		a.blacklist = true
	}

	if version, ok := aliasVersionAddenda[a.namespace.Name+"."+a.Name]; ok {
		a.Version = version
	}
}
