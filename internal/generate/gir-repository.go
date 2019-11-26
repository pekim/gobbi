package generate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type repository struct {
	// xml mappings
	XMLName xml.Name `xml:"repository"`
	Version string   `xml:"Version,attr"`
	//Includes  []Include  `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Namespace *Namespace `xml:"namespace"`
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
}

func repositoryFromFile(spec RepositorySpec) *repository {
	fullname := spec.Name + "-" + spec.Version

	r := &repository{}
	r.loadFromFile(fullname+".gir", true)

	return r
}
