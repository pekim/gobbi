package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type repository struct {
	// xml mappings
	XMLName   xml.Name  `xml:"repository"`
	Version   string    `xml:"Version,attr"`
	CIncludes CIncludes `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	//Includes  []Include  `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Packages  []Package  `xml:"package"`
	Namespace *Namespace `xml:"namespace"`
}

type Package struct {
	Name string `xml:"name,attr"`
}

func repositoryFromFile(spec repositorySpec) *repository {
	fullname := spec.Name + "-" + spec.Version

	r := &repository{}
	r.loadFromFile(fullname+".gir", true)

	return r
}

func (r *repository) loadFromFile(filename string, required bool) {
	filepath := projectFilepath("gir-files", filename)
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) && !required {
			return
		}

		panic(err)
	}

	err = xml.Unmarshal(source, r)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s : %s", filepath, err))
	}

	r.init()
}

func (r *repository) init() {
	if r.Namespace.Name == "GLib" {
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "glib/gstdio.h"})
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "glib-object.h"})

		r.Packages = append(r.Packages, Package{Name: "gobject-2.0"})
	}

	if r.Namespace.Name == "Gio" {
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "gio/gnetworking.h"})
	}

	if r.Namespace.Name == "Pango" {
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "pango/pango-font.h"})
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "pango/pango-modules.h"})
	}

	if r.Namespace.Name == "xlib" {
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "X11/Xlib.h"})
		r.CIncludes = append(r.CIncludes, &CInclude{Name: "X11/extensions/Xrender.h"})

		r.Packages = append(r.Packages, Package{Name: "x11"})
	}
}
