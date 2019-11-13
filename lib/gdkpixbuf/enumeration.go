// Code generated - DO NOT EDIT.

package gdkpixbuf

// Colorspace is a representation of the C type Colorspace.
type Colorspace int

const (
	// rgb
	GdkColorspaceRgb Colorspace = 0
)

// Interptype is a representation of the C type InterpType.
type Interptype int

const (
	// nearest
	GdkInterpNearest Interptype = 0
	// tiles
	GdkInterpTiles Interptype = 1
	// bilinear
	GdkInterpBilinear Interptype = 2
	// hyper
	GdkInterpHyper Interptype = 3
)

// Pixbufalphamode is a representation of the C type PixbufAlphaMode.
type Pixbufalphamode int

const (
	// bilevel
	GdkPixbufAlphaBilevel Pixbufalphamode = 0
	// full
	GdkPixbufAlphaFull Pixbufalphamode = 1
)

// Pixbuferror is a representation of the C type PixbufError.
type Pixbuferror int

const (
	// corrupt_image
	GdkPixbufErrorCorruptImage Pixbuferror = 0
	// insufficient_memory
	GdkPixbufErrorInsufficientMemory Pixbuferror = 1
	// bad_option
	GdkPixbufErrorBadOption Pixbuferror = 2
	// unknown_type
	GdkPixbufErrorUnknownType Pixbuferror = 3
	// unsupported_operation
	GdkPixbufErrorUnsupportedOperation Pixbuferror = 4
	// failed
	GdkPixbufErrorFailed Pixbuferror = 5
	// incomplete_animation
	GdkPixbufErrorIncompleteAnimation Pixbuferror = 6
)

// Pixbufrotation is a representation of the C type PixbufRotation.
type Pixbufrotation int

const (
	// none
	GdkPixbufRotateNone Pixbufrotation = 0
	// counterclockwise
	GdkPixbufRotateCounterclockwise Pixbufrotation = 90
	// upsidedown
	GdkPixbufRotateUpsidedown Pixbufrotation = 180
	// clockwise
	GdkPixbufRotateClockwise Pixbufrotation = 270
)
