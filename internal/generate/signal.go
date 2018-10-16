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
}

func (s *Signal) init(ns *Namespace, record *Record) {
	s.Namespace = ns
	s.record = record
}

func (s *Signal) generate(g *jen.Group, version *Version) {
	s.goNameHandler = fmt.Sprintf("%s_%s_%sHandler",
		s.Namespace.goPackageName,
		s.record.Name,
		makeGoNameInternal(s.Name, false))

	s.cNameSignalConnect = fmt.Sprintf("%s_signal_connect_%s",
		s.record.Name,
		strings.Replace(s.Name, "-", "_", -1))

	s.generateCgoPreamble()
}

func (s *Signal) generateCgoPreamble() {
	s.Namespace.jenFile.CgoPreamble(
		fmt.Sprintf(`
	extern void %s();

	static void %s(gpointer instance, gpointer data) {
		g_signal_connect(instance, "%s", %s, data);
	}

`,
			s.goNameHandler,
			s.cNameSignalConnect,
			s.Name,
			s.goNameHandler))
}
