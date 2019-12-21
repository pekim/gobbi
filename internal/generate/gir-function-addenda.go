package generate

var functionBlacklist = map[string]bool{
	"Pango.module_register": true,
}

func (f *Function) applyAddenda() {
	if _, ok := functionBlacklist[f.namespace.Name+"."+f.Name]; ok {
		f.blacklist = true
	}
}
