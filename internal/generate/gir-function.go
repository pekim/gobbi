package generate

import "github.com/dave/jennifer/jen"

type Function struct {
	Name              string `xml:"name,attr"`
	Blacklist         bool   `xml:"blacklist,attr"`
	Version           string `xml:"version,attr"`
	MovedTo           string `xml:"moved-to,attr"`
	CIdentifier       string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	//Doc               *Doc         `xml:"doc"`
	//InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters     Parameters   `xml:"parameters>parameter"`
	ReturnValue    *ReturnValue `xml:"return-value"`
	Throws         int          `xml:"throws,attr"`
	Introspectable string       `xml:"introspectable,attr"`

	namespace *Namespace
	goName    string

	//receiver   *Record
	//ctorRecord *Record
	//
	//throwableErrorType      *Type
	//throwableErrorCVarName  string
	//throwableErrorGoVarName string
}

func (f *Function) init(ns *Namespace /*receiver *Record, namePrefix string*/) {
	f.namespace = ns
	f.goName = makeExportedGoName(f.Name)

	//f.receiver = receiver
	//if f.GoName == "" {
	//	f.GoName = MakeExportedGoName(f.Name)
	//}
	//f.GoName = namePrefix + f.GoName
	//f.Parameters.init(ns)
	//if f.InstanceParameter != nil {
	//	f.InstanceParameter.init(ns)
	//}
	//f.initThrowableError()
	//
	//if f.ReturnValue != nil {
	//	f.ReturnValue.init(ns)
	//}
}

func (f *Function) generate(fi *file) {
	if len(f.Parameters) > 0 || f.ReturnValue != nil {
		fi.unsupported(f.CIdentifier, "non trivial function")
	}

	fi.docForC(f.goName, f.CIdentifier)

	fi.
		Func().
		Id(f.goName).
		Add(f.generateParams()).
		Add(f.generateBody())
}

func (f *Function) generateParams() *jen.Statement {
	return jen.
		ParamsFunc(func(jg *jen.Group) {
			//g := &group{jg}
		})
}

func (f *Function) generateBody() *jen.Statement {
	return jen.
		BlockFunc(func(jg *jen.Group) {
			g := &group{jg}

			g.
				Qual("fmt", "Println").
				Call(jen.
					Qual("github.com/pekim/gobbi/internal/gi", "Qaz"))
		})
}
