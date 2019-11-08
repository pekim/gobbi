package webkit2

import (
	"github.com/pekim/gobbi/lib/webkit2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	view := webkit2.WebViewNew()
	assert.NotNil(t, view)
}
