package gio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, "gio-desktop-app-info-lookup", DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME)
}
