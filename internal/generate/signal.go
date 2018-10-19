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

	record                   *Record
	varNameId                string
	varNameMap               string
	varNameLock              string
	goNameHandler            string
	goNameConnectFunction    string
	goNameDisconnectFunction string
	cNameSignalConnect       string
	callbackTypeName         string
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
	s.goNameDisconnectFunction = fmt.Sprintf("Disconnect%s", signalGoName)

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

func (s *Signal) supported() (bool, string) {
	if supported, reason := s.Parameters.allSupported(); !supported {
		return supported, reason
	}

	if s.ReturnValue != nil && s.ReturnValue.Type.Name != "none" {
		if supported, reason := s.ReturnValue.Type.generator.isSupportedAsReturnCValue(); !supported {
			return supported, fmt.Sprintf("return value %s : %s", s.ReturnValue.Type.Name, reason)
		}
	}

	return true, ""
}

func (s *Signal) generate(g *jen.Group, version *Version, parentVersion string) {
	if supported, reason := s.supported(); !supported {
		g.Commentf("Unsupported signal '%s' for %s : %s", s.Name, s.record.Name, reason)
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
	s.generateDisconnectFunction(g)
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
	g.Commentf("//export %s", s.goNameHandler)

	g.
		Func().
		Id(s.goNameHandler).
		Params(). //ParamFunc(s.Parameters.generateFunctionDeclarationCtypes).
		ParamsFunc(s.ReturnValue.generateFunctionDeclarationCtype).
		BlockFunc(func(g *jen.Group) {
			g.Qual("fmt", "Println").Call(jen.Lit("cb"))
		})

	g.Line()

	//func keyPressEventHandler(target *C.GObject, event *C.GdkEventKey, data uintptr) {
	//	goEvent := gdk.EventKeyNewFromC(unsafe.Pointer(event))
	//
	//	cb := signalKeyPressEventMap[int(data)]
	//	cb(goEvent)
	//}
}

func (s *Signal) generateLockUnlock(g *jen.Group) {
	//	signalKeyPressEventLock.Lock()
	g.Id(s.varNameLock).Op(".").Id("Lock").Call()

	//	defer signalKeyPressEventLock.Unlock()
	g.Defer().Id(s.varNameLock).Op(".").Id("Unlock").Call()
}

func (s *Signal) generateConnectFunction(g *jen.Group) {
	g.Commentf(
		`%s connects the callback to the '%s' signal for the %s.

The returned value represents the connection, and may be passed to %s to remove it.`,
		s.goNameConnectFunction,
		s.Name,
		s.record.GoName,
		s.goNameDisconnectFunction)

	g.
		Func().
		Params(jen.Id("recv").Op("*").Id(s.record.GoName)).
		Id(s.goNameConnectFunction).
		Params(jen.Id("callback").Id(s.callbackTypeName)).
		Params(jen.Int()).
		BlockFunc(func(g *jen.Group) {
			s.generateLockUnlock(g)
			g.Line()

			//	signalKeyPressEventId++
			g.Id(s.varNameId).Op("++")

			//	signalKeyPressEventMap[signalKeyPressEventId] = callback
			g.Id(s.varNameMap).Index(jen.Id(s.varNameId)).Op("=").Id("callback")

			g.Line()

			targetCValue := jen.Id("recv").
				Op(".").Id("Object").Call().
				Op(".").Id("ToC").Call()
			//	instance := C.gpointer(recv.Object().ToC())
			g.
				Id("instance").Op(":=").
				Qual("C", "gpointer").
				Call(targetCValue)

			//	retC := C.signal_connect_key_press_event(instance, C.gpointer(uintptr(signalKeyPressEventId)))
			g.
				Id("retC").
				Op(":=").
				Qual("C", s.cNameSignalConnect).
				CallFunc(func(g *jen.Group) {
					g.Id("instance")

					g.
						Qual("C", "gpointer").
						Call(jen.Id("uintptr").Call(jen.Id(s.varNameId)))
				})

			// return int(retC)
			g.Return(jen.Int().Call(jen.Id("retC")))
		})

	g.Line()
}

func (s *Signal) generateDisconnectFunction(g *jen.Group) {
	g.Commentf(
		`%s disconnects a callback from the '%s' signal for the %s.

The connectionID should be a value returned from a call to %s.`,
		s.goNameDisconnectFunction,
		s.Name,
		s.record.GoName,
		s.goNameConnectFunction)

	g.
		Func().
		Params(jen.Id("recv").Op("*").Id(s.record.GoName)).
		Id(s.goNameDisconnectFunction).
		Params(jen.Id("connectionID").Int()).
		Params().
		BlockFunc(func(g *jen.Group) {
			s.generateLockUnlock(g)
			g.Line()

			// _, exists := signalKeyPressEventMap[connectionID]
			g.
				Id("_").Op(",").Id("exists").
				Op(":=").
				Id(s.varNameMap).Index(jen.Id("connectionID"))

			// if !exists { return }
			g.
				If(jen.Op("!").Id("exists")).
				Block(jen.Return())

			g.Line()

			targetCValue := jen.Id("recv").
				Op(".").Id("Object").Call().
				Op(".").Id("ToC").Call()
			//	instance := C.gpointer(recv.Object().ToC())
			g.
				Id("instance").Op(":=").
				Qual("C", "gpointer").
				Call(targetCValue)

			// C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
			g.
				Qual("C", "g_signal_handler_disconnect").
				Call(
					jen.Id("instance"),
					jen.Qual("C", "gulong").Call(jen.Id("connectionID")),
				)

			//	delete(signalKeyPressEventMap, connectionId)
			g.Delete(jen.Id(s.varNameMap), jen.Id("connectionID"))
		})

	g.Line()
}
