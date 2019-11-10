// Code generated - DO NOT EDIT.

package gdkpixbuf

type Colorspace uint32

const (
	Colorspace_Rgb Colorspace = int64(0)
)

type InterpType uint32

const (
	InterpType_Nearest  InterpType = int64(0)
	InterpType_Tiles    InterpType = int64(1)
	InterpType_Bilinear InterpType = int64(2)
	InterpType_Hyper    InterpType = int64(3)
)

type PixbufAlphaMode uint32

const (
	PixbufAlphaMode_Bilevel PixbufAlphaMode = int64(0)
	PixbufAlphaMode_Full    PixbufAlphaMode = int64(1)
)

type PixbufError uint32

const (
	PixbufError_CorruptImage         PixbufError = int64(0)
	PixbufError_InsufficientMemory   PixbufError = int64(1)
	PixbufError_BadOption            PixbufError = int64(2)
	PixbufError_UnknownType          PixbufError = int64(3)
	PixbufError_UnsupportedOperation PixbufError = int64(4)
	PixbufError_Failed               PixbufError = int64(5)
	PixbufError_IncompleteAnimation  PixbufError = int64(6)
)

type PixbufRotation uint32

const (
	PixbufRotation_None             PixbufRotation = int64(0)
	PixbufRotation_Counterclockwise PixbufRotation = int64(90)
	PixbufRotation_Upsidedown       PixbufRotation = int64(180)
	PixbufRotation_Clockwise        PixbufRotation = int64(270)
)
