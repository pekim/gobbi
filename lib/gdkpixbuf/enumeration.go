// Code generated - DO NOT EDIT.

package gdkpixbuf

// Colorspace is a representation of the C type GdkColorspace.
type Colorspace int

const (
	// Colorspace_Rgb is a representation of the C type GDK_COLORSPACE_RGB.
	Colorspace_Rgb Colorspace = 0
)

// Interptype is a representation of the C type GdkInterpType.
type Interptype int

const (
	// Interptype_Nearest is a representation of the C type GDK_INTERP_NEAREST.
	Interptype_Nearest Interptype = 0
	// Interptype_Tiles is a representation of the C type GDK_INTERP_TILES.
	Interptype_Tiles Interptype = 1
	// Interptype_Bilinear is a representation of the C type GDK_INTERP_BILINEAR.
	Interptype_Bilinear Interptype = 2
	// Interptype_Hyper is a representation of the C type GDK_INTERP_HYPER.
	Interptype_Hyper Interptype = 3
)

// Pixbufalphamode is a representation of the C type GdkPixbufAlphaMode.
type Pixbufalphamode int

const (
	// Pixbufalphamode_Bilevel is a representation of the C type GDK_PIXBUF_ALPHA_BILEVEL.
	Pixbufalphamode_Bilevel Pixbufalphamode = 0
	// Pixbufalphamode_Full is a representation of the C type GDK_PIXBUF_ALPHA_FULL.
	Pixbufalphamode_Full Pixbufalphamode = 1
)

// Pixbuferror is a representation of the C type GdkPixbufError.
type Pixbuferror int

const (
	// Pixbuferror_CorruptImage is a representation of the C type GDK_PIXBUF_ERROR_CORRUPT_IMAGE.
	Pixbuferror_CorruptImage Pixbuferror = 0
	// Pixbuferror_InsufficientMemory is a representation of the C type GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY.
	Pixbuferror_InsufficientMemory Pixbuferror = 1
	// Pixbuferror_BadOption is a representation of the C type GDK_PIXBUF_ERROR_BAD_OPTION.
	Pixbuferror_BadOption Pixbuferror = 2
	// Pixbuferror_UnknownType is a representation of the C type GDK_PIXBUF_ERROR_UNKNOWN_TYPE.
	Pixbuferror_UnknownType Pixbuferror = 3
	// Pixbuferror_UnsupportedOperation is a representation of the C type GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION.
	Pixbuferror_UnsupportedOperation Pixbuferror = 4
	// Pixbuferror_Failed is a representation of the C type GDK_PIXBUF_ERROR_FAILED.
	Pixbuferror_Failed Pixbuferror = 5
	// Pixbuferror_IncompleteAnimation is a representation of the C type GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION.
	Pixbuferror_IncompleteAnimation Pixbuferror = 6
)

// Pixbufrotation is a representation of the C type GdkPixbufRotation.
type Pixbufrotation int

const (
	// Pixbufrotation_None is a representation of the C type GDK_PIXBUF_ROTATE_NONE.
	Pixbufrotation_None Pixbufrotation = 0
	// Pixbufrotation_Counterclockwise is a representation of the C type GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE.
	Pixbufrotation_Counterclockwise Pixbufrotation = 90
	// Pixbufrotation_Upsidedown is a representation of the C type GDK_PIXBUF_ROTATE_UPSIDEDOWN.
	Pixbufrotation_Upsidedown Pixbufrotation = 180
	// Pixbufrotation_Clockwise is a representation of the C type GDK_PIXBUF_ROTATE_CLOCKWISE.
	Pixbufrotation_Clockwise Pixbufrotation = 270
)
