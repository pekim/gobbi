package generate

import (
	"fmt"
	"os"
	"path"
)

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}

func (cc Constants) versionList() Versions {
	var versions Versions

	for _, c := range cc {
		if c.Version != "" {
			versions = append(versions, VersionNew(c.Version))
		}
	}

	return versions
}

func (cc Constants) entities() []Generatable {
	var generatables []Generatable

	for _, constant := range cc {
		generatables = append(generatables, constant)
	}

	return generatables
}

func (cc Constants) forName(name string) *Constant {
	for _, constant := range cc {
		if constant.Name == name {
			return constant
		}
	}

	return nil
}

func (cc Constants) mergeAddenda(addenda Constants) {
	for _, addendaContant := range addenda {
		if constant := cc.forName(addendaContant.Name); constant != nil {
			constant.mergeAddenda(addendaContant)
		}
	}
}

func (cc Constants) generateDocs(ns *Namespace) {
	filename := path.Join(ns.docsDir, "constant.md")
	//fmt.Println(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(fmt.Sprintf(
		"# `%s` Constants\n\n",
		ns.goPackageName,
	))
	if err != nil {
		panic(err)
	}

	for _, constant := range cc {
		blacklisted, _ := constant.blacklisted()
		supported, _ := constant.supported()
		if blacklisted || !supported {
			continue
		}

		constant.generateDocs(file)
	}

	file.Close()
}
