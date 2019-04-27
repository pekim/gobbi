package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportsSrc(t *testing.T) {
	f := New("", "")
	f.imprt("fmt")
	f.gobbiImprt("glib")
	f.gobbiImprt("glib")

	assert.Equal(t,
		""+
			"import \"fmt\""+"\n"+
			"import \"github.com/pekim/gobbi/glib\""+"\n"+
			"\n",
		f.importsSrc())
}

func TestImportsOrder(t *testing.T) {
	f := New("", "")
	f.gobbiImprt("glib")
	f.imprt("fmt")

	assert.Equal(t,
		""+
			"import \"fmt\""+"\n"+
			"import \"github.com/pekim/gobbi/glib\""+"\n"+
			"\n",
		f.importsSrc())
}

func TestImportsSamePackage(t *testing.T) {
	f := New("gtk", "")
	f.gobbiImprt("glib")
	f.gobbiImprt("gtk")

	assert.Equal(t,
		""+
			"import \"github.com/pekim/gobbi/glib\""+"\n"+
			"\n",
		f.importsSrc())
}

func TestAddConstInt(t *testing.T) {
	f := New("", "")
	f.AddConstInt("name", 42)

	assert.Equal(t, "const name int = 42\n", f.contents)
}

func TestAddConstString(t *testing.T) {
	f := New("", "")
	f.AddConstString("name", "qaz")

	assert.Equal(t, "const name string = \"qaz\"\n", f.contents)
}

func TestAddConsts(t *testing.T) {
	f := New("", "")
	f.AddConsts("int32", func(add func(name string, value int)) {
		add("one", 1)
		add("two", 2)
		add("three", 3)
	})

	assert.Equal(t,
		`const (
	one int32 = 1
	two = 2
	three = 3
)

`, f.contents)
}

func TestSrc(t *testing.T) {
	f := New("pkg", "")
	f.imprt("fmt")

	assert.Equal(t,
		""+
			"package pkg"+"\n"+
			"\n"+
			"import \"fmt\""+"\n"+
			"\n",
		f.src())
}

func TestFormattedSrc(t *testing.T) {
	f := New("pkg", "")
	f.imprt("fmt")

	assert.Equal(t,
		""+
			"package pkg"+"\n"+
			"\n"+
			"import \"fmt\""+"\n",
		string(f.formattedSrc()))
}

//func TestWrite(t *testing.T) {
//	f := New("qaz", "./qqq")
//	f.imprt("fmt")
//	f.gobbiImprt("glib")
//
//	f.Write()
//}
