// Code generated - DO NOT EDIT.

package gdkpixbuf

// Colorspace is a representation of the C type Colorspace.
type Colorspace int

const (
	Colorspace_Rgb Colorspace = 0
)

// Interptype is a representation of the C type InterpType.
type Interptype int

const (
	Interptype_Nearest  Interptype = 0
	Interptype_Tiles    Interptype = 1
	Interptype_Bilinear Interptype = 2
	Interptype_Hyper    Interptype = 3
)

// Pixbufalphamode is a representation of the C type PixbufAlphaMode.
type Pixbufalphamode int

const (
	Pixbufalphamode_Bilevel Pixbufalphamode = 0
	Pixbufalphamode_Full    Pixbufalphamode = 1
)

// Pixbuferror is a representation of the C type PixbufError.
type Pixbuferror int

const (
	Pixbuferror_CorruptImage         Pixbuferror = 0
	Pixbuferror_InsufficientMemory   Pixbuferror = 1
	Pixbuferror_BadOption            Pixbuferror = 2
	Pixbuferror_UnknownType          Pixbuferror = 3
	Pixbuferror_UnsupportedOperation Pixbuferror = 4
	Pixbuferror_Failed               Pixbuferror = 5
	Pixbuferror_IncompleteAnimation  Pixbuferror = 6
)

// Pixbufrotation is a representation of the C type PixbufRotation.
type Pixbufrotation int

const (
	Pixbufrotation_None             Pixbufrotation = 0
	Pixbufrotation_Counterclockwise Pixbufrotation = 90
	Pixbufrotation_Upsidedown       Pixbufrotation = 180
	Pixbufrotation_Clockwise        Pixbufrotation = 270
)
