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
			"import fmt"+"\n"+
			"import github.com/pekim/gobbi/glib"+"\n"+
			"\n",
		f.importsSrc())
}
