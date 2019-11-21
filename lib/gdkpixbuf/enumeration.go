// Code generated - DO NOT EDIT.

package gdkpixbuf

// Colorspace is a representation of the C type GdkColorspace.
//
type Colorspace int

const (
	// Colorspace_Rgb is a representation of the C type GDK_COLORSPACE_RGB.
	Colorspace_Rgb Colorspace = 0
)

// InterpType is a representation of the C type GdkInterpType.
//
type InterpType int

const (
	// InterpType_Nearest is a representation of the C type GDK_INTERP_NEAREST.
	InterpType_Nearest InterpType = 0
	// InterpType_Tiles is a representation of the C type GDK_INTERP_TILES.
	InterpType_Tiles InterpType = 1
	// InterpType_Bilinear is a representation of the C type GDK_INTERP_BILINEAR.
	InterpType_Bilinear InterpType = 2
	// InterpType_Hyper is a representation of the C type GDK_INTERP_HYPER.
	InterpType_Hyper InterpType = 3
)

// PixbufAlphaMode is a representation of the C type GdkPixbufAlphaMode.
//
type PixbufAlphaMode int

const (
	// PixbufAlphaMode_Bilevel is a representation of the C type GDK_PIXBUF_ALPHA_BILEVEL.
	PixbufAlphaMode_Bilevel PixbufAlphaMode = 0
	// PixbufAlphaMode_Full is a representation of the C type GDK_PIXBUF_ALPHA_FULL.
	PixbufAlphaMode_Full PixbufAlphaMode = 1
)

// PixbufError is a representation of the C type GdkPixbufError.
//
type PixbufError int

const (
	// PixbufError_CorruptImage is a representation of the C type GDK_PIXBUF_ERROR_CORRUPT_IMAGE.
	PixbufError_CorruptImage PixbufError = 0
	// PixbufError_InsufficientMemory is a representation of the C type GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY.
	PixbufError_InsufficientMemory PixbufError = 1
	// PixbufError_BadOption is a representation of the C type GDK_PIXBUF_ERROR_BAD_OPTION.
	PixbufError_BadOption PixbufError = 2
	// PixbufError_UnknownType is a representation of the C type GDK_PIXBUF_ERROR_UNKNOWN_TYPE.
	PixbufError_UnknownType PixbufError = 3
	// PixbufError_UnsupportedOperation is a representation of the C type GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION.
	PixbufError_UnsupportedOperation PixbufError = 4
	// PixbufError_Failed is a representation of the C type GDK_PIXBUF_ERROR_FAILED.
	PixbufError_Failed PixbufError = 5
	// PixbufError_IncompleteAnimation is a representation of the C type GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION.
	PixbufError_IncompleteAnimation PixbufError = 6
)

// PixbufRotation is a representation of the C type GdkPixbufRotation.
//
type PixbufRotation int

const (
	// PixbufRotation_None is a representation of the C type GDK_PIXBUF_ROTATE_NONE.
	PixbufRotation_None PixbufRotation = 0
	// PixbufRotation_Counterclockwise is a representation of the C type GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE.
	PixbufRotation_Counterclockwise PixbufRotation = 90
	// PixbufRotation_Upsidedown is a representation of the C type GDK_PIXBUF_ROTATE_UPSIDEDOWN.
	PixbufRotation_Upsidedown PixbufRotation = 180
	// PixbufRotation_Clockwise is a representation of the C type GDK_PIXBUF_ROTATE_CLOCKWISE.
	PixbufRotation_Clockwise PixbufRotation = 270
)
