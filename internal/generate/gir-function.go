package generate

import "C"
import (
	"github.com/blang/semver"
)

type Function struct {
	context           *context
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	MovedTo           string `xml:"moved-to,attr"`
	ShadowedBy        string `xml:"shadowed-by,attr"`
	CIdentifier       string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	//Doc               *Doc   `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        Parameters   `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            bool         `xml:"throws,attr"`
	Introspectable    string       `xml:"introspectable,attr"`

	namespace *Namespace
	blacklist bool
	version   semver.Version
	goName    string
	sysName   string
}

func (f *Function) init(ns *Namespace, record *Record, receiver bool) {
	f.context = newContext(ns.context, "Function", f.Name)
	f.namespace = ns
	f.applyAddenda()
	f.version = versionNew(f.Version)
	f.namespace.versions.add(f.version)

	f.goName = makeExportedGoName(f.Name)
	if _, found := ns.Classes.byName(f.goName); found {
		f.goName += "_"
	}
	if _, found := ns.Records.byName(f.goName); found {
		f.goName += "_"
	}

	f.sysName = "Fn_" + f.CIdentifier
	if f.InstanceParameter != nil {
		f.InstanceParameter.init(ns, f.context)
	}
	f.Parameters.init(ns, f.context)
	f.ReturnValue.init(ns, f.Parameters)
}

func (f *Function) isSupported() (bool, string) {
	if f.blacklist {
		return false, "blacklisted"
	}

	if supported, reason := f.Parameters.allSupported(); !supported {
		return false, reason
	}

	if supported, reason := f.ReturnValue.isSupported(); !supported {
		return false, reason
	}

	return true, ""
}
