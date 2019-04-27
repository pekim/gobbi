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
		f.formattedSrc())
}
