package generate

var functionBlacklist = map[string]bool{
	"GLib.unix_error_quark": true,

	"Pango.module_register": true,
}

func (f *Function) applyAddenda() {
	if f.namespace.Name == "cairo" {
		f.blacklist = true
		return
	}

	if _, ok := functionBlacklist[f.namespace.Name+"."+f.Name]; ok {
		f.blacklist = true
	}
}
