package generate

type Functions []*Function

func (ff Functions) init(ns *Namespace /*, namePrefix string*/) {
	for _, function := range ff {
		function.init(ns /*nil*/, "")
	}
}

func (ff Functions) generate(f *file) {
	for _, fn := range ff {
		fn.generate(f)
	}
}
