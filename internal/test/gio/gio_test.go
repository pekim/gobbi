package gio

import (
	"github.com/pekim/gobbi/lib/gio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	app := gio.ApplicationNew("pekim.gobbi.testapp",
		gio.APPLICATION_FLAGS_NONE)
	assert.NotNil(t, app)
}

// g_networking_init is defined in a header file that is not included
// in the gio gir repo file.
func TestAddendaCInclude(t *testing.T) {
	gio.NetworkingInit()
}
