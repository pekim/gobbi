package generate

var unionBlacklistBlacklist = map[string]bool{
	"GObject._Value__data__union": true,
}

func (u *Union) applyAddenda() {
	if _, ok := unionBlacklistBlacklist[u.namespace.Name+"."+u.Name]; ok {
		u.blacklist = true
	}
}
