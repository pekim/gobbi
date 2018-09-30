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
