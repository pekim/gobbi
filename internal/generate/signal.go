package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Signal struct {
	Namespace *Namespace

	//Blacklist     bool `xml:"blacklist,attr"`

	Name        string       `xml:"name,attr"`
	When        string       `xml:"when,attr"`
	Version     string       `xml:"version,attr"`
	Doc         *Doc         `xml:"doc"`
	Parameters  Parameters   `xml:"parameters>parameter"`
	ReturnValue *ReturnValue `xml:"return-value"`

	record             *Record
	goNameHandler      string
	cNameSignalConnect string
	callbackTypeName   string
}

func (s *Signal) init(ns *Namespace, record *Record) {
	s.Namespace = ns
	s.record = record

	s.goNameHandler = fmt.Sprintf("%s_%sHandler",
		s.record.Name,
		makeGoNameInternal(s.Name, false))

	s.cNameSignalConnect = fmt.Sprintf("%s_signal_connect_%s",
		s.record.Name,
		strings.Replace(s.Name, "-", "_", -1))

	s.callbackTypeName = fmt.Sprintf("%sSignal%sCallback",
		s.record.Name,
		makeExportedGoName(s.Name))
}

func (s *Signal) generate(g *jen.Group, version *Version) {
	s.generateCgoPreamble()
	s.generateCallbackType(g)
}

func (s *Signal) generateCgoPreamble() {
	s.Namespace.jenFile.CgoPreamble(
		fmt.Sprintf(`
	void %s();

	static gulong %s(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "%s", %s, data);
	}

`,
			s.goNameHandler,
			s.cNameSignalConnect,
			s.Name,
			s.goNameHandler))
}

func (s *Signal) generateCallbackType(g *jen.Group) {
	g.Commentf("%s is a callback function for a '%s' signal emitted from a %s.",
		s.callbackTypeName,
		s.Name,
		s.record.Name)

	g.
		Type().
		Id(s.callbackTypeName).
		Func().
		Params()

	g.Line()
}
