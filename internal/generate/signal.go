package generate

import "C"
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

	record                *Record
	varNameId             string
	varNameMap            string
	varNameLock           string
	goNameHandler         string
	goNameConnectFunction string
	cNameSignalConnect    string
	callbackTypeName      string
}

func (s *Signal) init(ns *Namespace, record *Record) {
	s.Namespace = ns
	s.record = record

	s.Parameters.init(ns)

	if s.ReturnValue != nil {
		s.ReturnValue.init(ns)
	}

	s.initNames()
}

func (s *Signal) initNames() {
	signalGoName := makeExportedGoName(s.Name)

	s.varNameId = fmt.Sprintf("signal%s%sId", s.record.Name, signalGoName)
	s.varNameMap = fmt.Sprintf("signal%s%sMap", s.record.Name, signalGoName)
	s.varNameLock = fmt.Sprintf("signal%s%sLock", s.record.Name, signalGoName)

	s.goNameHandler = fmt.Sprintf("%s_%sHandler",
		s.record.Name,
		makeGoNameInternal(s.Name, false))

	s.goNameConnectFunction = fmt.Sprintf("Connect%s", signalGoName)

	s.cNameSignalConnect = fmt.Sprintf("%s_signal_connect_%s",
		s.record.Name,
		strings.Replace(s.Name, "-", "_", -1))

	s.callbackTypeName = fmt.Sprintf("%sSignal%sCallback",
		s.record.Name,
		signalGoName)
}

func (s *Signal) version() string {
	return s.Version
}

func (s *Signal) generate(g *jen.Group, version *Version, parentVersion string) {
	supported, reason := s.Parameters.allSupported()
	if !supported {
		g.Commentf("Unsupported signal : %s", reason)
		g.Line()
		return
	}

	if !((parentVersion == "" && s.Version == version.value) ||
		(parentVersion != "" && (s.Version == "" && parentVersion == version.value))) {
		return
	}

	s.generateCgoPreamble()
	s.generateVariables(g)
	s.generateCallbackType(g)
	s.generateConnectFunction(g)
	s.generateHandlerFunction(g)
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

func (s *Signal) generateVariables(g *jen.Group) {
	g.Var().Id(s.varNameId).Int()
	g.Var().Id(s.varNameMap).Op("=").Make(jen.Map(jen.Int()).Id(s.callbackTypeName))
	g.Var().Id(s.varNameLock).Qual("sync", "Mutex")
	g.Line()
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
		ParamsFunc(s.Parameters.generateFunctionDeclaration).
		ParamsFunc(s.generateCallbackReturnDeclaration)

	g.Line()
}

func (s *Signal) generateCallbackReturnDeclaration(g *jen.Group) {
	s.ReturnValue.generateFunctionDeclaration(g)
}

func (s *Signal) generateHandlerFunction(g *jen.Group) {
	g.
		Func().
		Id(s.goNameHandler).
		Params().
		Params().
		BlockFunc(func(g *jen.Group) {

		})

	g.Line()
}

func (s *Signal) generateConnectFunction(g *jen.Group) {
	g.
		Func().
		Params(jen.Id("recv").Op("*").Id(s.record.GoName)).
		Id(s.goNameConnectFunction).
		Params().
		BlockFunc(func(g *jen.Group) {

		})

	g.Line()
}

//func connectKeyPressEvent(target *gtk.Widget, callback KeyPressEventCallback) {
//	signalKeyPressEventLock.Lock()
//	defer signalKeyPressEventLock.Unlock()
//
//	signalKeyPressEventId++
//	signalKeyPressEventMap[signalKeyPressEventId] = callback
//
//	instance := C.gpointer(target.Object().ToC())
//
//	C.signal_connect_key_press_event(instance, C.gpointer(uintptr(signalKeyPressEventId)))
//}
