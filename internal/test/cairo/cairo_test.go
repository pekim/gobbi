package cairo

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDashes(t *testing.T) {
	surface := cairo.ImageSurfaceCreate(cairo.CAIRO_FORMAT_ARGB32, 10, 10)
	cr := cairo.Create(surface)

	dashesIn := []float64{1.0, 2.0, 3.0}
	offsetIn := 0.5

	cr.SetDash(dashesIn, offsetIn)
	dashesOut, offsetOut := cr.GetDash()

	assert.Equal(t, dashesIn, dashesOut)
	assert.Equal(t, offsetIn, offsetOut)
}

func TestCopyClipRectangleList(t *testing.T) {
	surface := cairo.ImageSurfaceCreate(cairo.CAIRO_FORMAT_ARGB32, 10, 10)
	cr := cairo.Create(surface)

	rectangles, status := cr.CopyClipRectangleList()

	assert.Equal(t, cairo.CAIRO_STATUS_SUCCESS, status)
	assert.Equal(t, 1, len(rectangles))
	assert.Equal(t, cairo.Rectangle{0, 0, 10, 10}, rectangles[0])
}
