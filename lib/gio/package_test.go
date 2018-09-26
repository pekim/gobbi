package gio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	app := ApplicationNew("pekim.gobbi.testapp", APPLICATION_FLAGS_NONE)
	assert.NotNil(t, app)
}
