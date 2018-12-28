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
	typeStructName           string
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

	s.fixParameterNames()
	s.Parameters.init(ns)

	if s.ReturnValue != nil {
		s.ReturnValue.init(ns)
	}

	s.initNames()
}

// fixParameterNames adjusts some parameter names to avoid a conflict
// generated vars.
func (s *Signal) fixParameterNames() {
	for _, param := range s.Parameters {
		if param.Name == "index" || param.Name == "data" {
			param.Name = "_" + param.Name
		}
	}
}

func (s *Signal) initNames() {
	signalGoName := makeExportedGoName(s.Name)

	s.typeStructName = fmt.Sprintf("signal%s%sDetail", s.record.Name, signalGoName)
	s.varNameId = fmt.Sprintf("signal%s%sId", s.record.Name, signalGoName)
	s.varNameMap = fmt.Sprintf("signal%s%sMap", s.record.Name, signalGoName)
	s.varNameLock = fmt.Sprintf("signal%s%sLock", s.record.Name, signalGoName)

	s.goNameHandler = fmt.Sprintf("%s_%sHandler",
		makeGoNameInternal(s.record.Name, false),
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
	if supported, reason := s.Parameters.allSupportedC(); !supported {
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
	if !((parentVersion == "" && s.Version == version.value) ||
		(parentVersion != "" && (s.Version == "" && parentVersion == version.value)) ||
		(s.Version == version.value)) {
		return
	}

	if supported, reason := s.supported(); !supported {
		g.Commentf("Unsupported signal '%s' for %s : %s", s.Name, s.record.Name, reason)
		g.Line()
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
	handlerReturn := s.cTypeDeclaration(s.ReturnValue.Type)

	params := []string{"GObject *"}
	for _, param := range s.Parameters {
		if param.Array != nil {
			params = append(params, param.Array.CType)
		} else {
			params = append(params, s.cTypeDeclaration(param.Type))
		}
	}
	params = append(params, "gpointer")
	handlerParams := strings.Join(params, ", ")

	s.Namespace.jenFile.CgoPreamble(
		fmt.Sprintf(`
	%s %s(%s);

	static gulong %s(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "%s", G_CALLBACK(%s), data);
	}

`,
			handlerReturn,
			s.goNameHandler,
			handlerParams,
			s.cNameSignalConnect,
			s.Name,
			s.goNameHandler))
}

func (s *Signal) cTypeDeclaration(typ *Type) string {
	cDeclaration := typ.CType

	if cDeclaration == "" {
		qname := QNameNew(s.Namespace, typ.Name)
		record, found := qname.namespace.recordOrClassRecordForName(qname.name)
		if !found {
			iface, found := qname.namespace.interfaceForName(qname.name)
			if !found {
				panic(fmt.Sprintf("Not found class, record, or interface %s, for signal %s, for class %s",
					qname.name, s.Name, s.record.Name))
			}
			cDeclaration = iface.CType + " *"
		} else {
			cDeclaration = record.CType + " *"
		}
	}

	return cDeclaration
}

func (s *Signal) generateVariables(g *jen.Group) {
	g.Type().Id(s.typeStructName).StructFunc(func(g *jen.Group) {
		g.Id("callback").Id(s.callbackTypeName)
		g.Id("handlerID").Qual("C", "gulong")
	})

	g.Var().Id(s.varNameId).Int()
	g.Var().Id(s.varNameMap).Op("=").Make(jen.Map(jen.Int()).Id(s.typeStructName))
	g.Var().Id(s.varNameLock).Qual("sync", "RWMutex")

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
		ParamsFunc(s.generateHandlerParameters).
		ParamsFunc(s.ReturnValue.generateFunctionDeclarationCtype).
		BlockFunc(func(g *jen.Group) {
			s.generateLockUnlock(g, true)
			g.Line()

			s.Parameters.generateGoVars(g)
			s.generateHandlerCall(g)
			s.generateHandlerReturn(g)
		})

	g.Line()
}

func (s *Signal) generateHandlerParameters(g *jen.Group) {
	g.Id("_").Op("*").Qual("C", "GObject")
	s.Parameters.generateFunctionDeclarationCtypes(g)
	g.Id("data").Qual("C", "gpointer")
}

func (s *Signal) generateHandlerCall(g *jen.Group) {
	// index := int(c_index)
	g.Id("index").Op(":=").Int().Call(jen.Uintptr().Call(jen.Id("data")))

	//	callback := signalKeyPressEventMap[signalKeyPressEventId].callback
	g.Id("callback").Op(":=").Id(s.varNameMap).Index(jen.Id("index")).Dot("callback")

	if s.ReturnValue.Type.Name != "none" {
		// retGo := callback(...)
		g.Id("retGo").Op(":=").Id("callback").CallFunc(s.generateHandleCallParams)
	} else {
		// callback(...)
		g.Id("callback").CallFunc(s.generateHandleCallParams)
	}
}

func (s *Signal) generateHandleCallParams(g *jen.Group) {
	for _, p := range s.Parameters {
		if p.arrayLengthFor != nil {
			continue
		}

		g.Id(p.goVarName)
	}
}

func (s *Signal) generateHandlerReturn(g *jen.Group) {
	if s.ReturnValue.Type.Name == "none" {
		return
	}

	g.Id("retC").Op(":=")
	s.ReturnValue.Type.generator.generateGoToC(g, jen.Id("retGo"))

	g.Return().Id("retC")

}

func (s *Signal) generateLockUnlock(g *jen.Group, readonly bool) {
	lockFunction := "Lock"
	unlockFunction := "Unlock"
	if readonly {
		lockFunction = "RLock"
		unlockFunction = "RUnlock"
	}

	//	signalKeyPressEventLock.Lock()
	g.Id(s.varNameLock).Op(".").Id(lockFunction).Call()

	//	defer signalKeyPressEventLock.Unlock()
	g.Defer().Id(s.varNameLock).Op(".").Id(unlockFunction).Call()
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
			s.generateLockUnlock(g, false)
			g.Line()

			//	signalKeyPressEventId++
			g.Id(s.varNameId).Op("++")

			// instance := C.gpointer(recv.native)
			g.
				Id("instance").Op(":=").
				Qual("C", "gpointer").Call(jen.Id("recv").Dot("native"))

			//	retC := C.signal_connect_key_press_event(instance, C.gpointer(uintptr(signalKeyPressEventId)))
			g.
				Id("handlerID").
				Op(":=").
				Qual("C", s.cNameSignalConnect).
				CallFunc(func(g *jen.Group) {
					g.Id("instance")

					g.
						Qual("C", "gpointer").
						Call(jen.Id("uintptr").Call(jen.Id(s.varNameId)))
				})
			g.Line()

			// detail := signalWidgetKeyPressEventDetail{callback, handlerId}
			g.Id("detail").Op(":=").Id(s.typeStructName).Values(
				jen.Id("callback"),
				jen.Id("handlerID"))
			//	signalKeyPressEventMap[signalKeyPressEventId] = callback
			g.Id(s.varNameMap).Index(jen.Id(s.varNameId)).Op("=").Id("detail")
			g.Line()

			// return signalKeyPressEventId
			g.Return(jen.Id(s.varNameId))
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
			s.generateLockUnlock(g, false)
			g.Line()

			// detail, exists := signalKeyPressEventMap[connectionID]
			g.
				Id("detail").Op(",").Id("exists").
				Op(":=").
				Id(s.varNameMap).Index(jen.Id("connectionID"))

			// if !exists { return }
			g.
				If(jen.Op("!").Id("exists")).
				Block(jen.Return())

			g.Line()

			// instance := C.gpointer(recv.native)
			g.
				Id("instance").Op(":=").
				Qual("C", "gpointer").Call(jen.Id("recv").Dot("native"))

			// C.g_signal_handler_disconnect(instance, detail.handlerID)
			g.
				Qual("C", "g_signal_handler_disconnect").
				Call(
					jen.Id("instance"),
					jen.Id("detail").Dot("handlerID"),
				)

			//	delete(signalKeyPressEventMap, connectionId)
			g.Delete(jen.Id(s.varNameMap), jen.Id("connectionID"))
		})

	g.Line()
}
