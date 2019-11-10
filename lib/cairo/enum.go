// Code generated - DO NOT EDIT.

package cairo

type Status uint32

const (
	Status_Success                 Status = int64(0)
	Status_NoMemory                Status = int64(1)
	Status_InvalidRestore          Status = int64(2)
	Status_InvalidPopGroup         Status = int64(3)
	Status_NoCurrentPoint          Status = int64(4)
	Status_InvalidMatrix           Status = int64(5)
	Status_InvalidStatus           Status = int64(6)
	Status_NullPointer             Status = int64(7)
	Status_InvalidString           Status = int64(8)
	Status_InvalidPathData         Status = int64(9)
	Status_ReadError               Status = int64(10)
	Status_WriteError              Status = int64(11)
	Status_SurfaceFinished         Status = int64(12)
	Status_SurfaceTypeMismatch     Status = int64(13)
	Status_PatternTypeMismatch     Status = int64(14)
	Status_InvalidContent          Status = int64(15)
	Status_InvalidFormat           Status = int64(16)
	Status_InvalidVisual           Status = int64(17)
	Status_FileNotFound            Status = int64(18)
	Status_InvalidDash             Status = int64(19)
	Status_InvalidDscComment       Status = int64(20)
	Status_InvalidIndex            Status = int64(21)
	Status_ClipNotRepresentable    Status = int64(22)
	Status_TempFileError           Status = int64(23)
	Status_InvalidStride           Status = int64(24)
	Status_FontTypeMismatch        Status = int64(25)
	Status_UserFontImmutable       Status = int64(26)
	Status_UserFontError           Status = int64(27)
	Status_NegativeCount           Status = int64(28)
	Status_InvalidClusters         Status = int64(29)
	Status_InvalidSlant            Status = int64(30)
	Status_InvalidWeight           Status = int64(31)
	Status_InvalidSize             Status = int64(32)
	Status_UserFontNotImplemented  Status = int64(33)
	Status_DeviceTypeMismatch      Status = int64(34)
	Status_DeviceError             Status = int64(35)
	Status_InvalidMeshConstruction Status = int64(36)
	Status_DeviceFinished          Status = int64(37)
	Status_Jbig2GlobalMissing      Status = int64(38)
)

type Content uint32

const (
	Content_Color      Content = int64(4096)
	Content_Alpha      Content = int64(8192)
	Content_ColorAlpha Content = int64(12288)
)

type Operator uint32

const (
	Operator_Clear         Operator = int64(0)
	Operator_Source        Operator = int64(1)
	Operator_Over          Operator = int64(2)
	Operator_In            Operator = int64(3)
	Operator_Out           Operator = int64(4)
	Operator_Atop          Operator = int64(5)
	Operator_Dest          Operator = int64(6)
	Operator_DestOver      Operator = int64(7)
	Operator_DestIn        Operator = int64(8)
	Operator_DestOut       Operator = int64(9)
	Operator_DestAtop      Operator = int64(10)
	Operator_Xor           Operator = int64(11)
	Operator_Add           Operator = int64(12)
	Operator_Saturate      Operator = int64(13)
	Operator_Multiply      Operator = int64(14)
	Operator_Screen        Operator = int64(15)
	Operator_Overlay       Operator = int64(16)
	Operator_Darken        Operator = int64(17)
	Operator_Lighten       Operator = int64(18)
	Operator_ColorDodge    Operator = int64(19)
	Operator_ColorBurn     Operator = int64(20)
	Operator_HardLight     Operator = int64(21)
	Operator_SoftLight     Operator = int64(22)
	Operator_Difference    Operator = int64(23)
	Operator_Exclusion     Operator = int64(24)
	Operator_HslHue        Operator = int64(25)
	Operator_HslSaturation Operator = int64(26)
	Operator_HslColor      Operator = int64(27)
	Operator_HslLuminosity Operator = int64(28)
)

type Antialias uint32

const (
	Antialias_Default  Antialias = int64(0)
	Antialias_None     Antialias = int64(1)
	Antialias_Gray     Antialias = int64(2)
	Antialias_Subpixel Antialias = int64(3)
	Antialias_Fast     Antialias = int64(4)
	Antialias_Good     Antialias = int64(5)
	Antialias_Best     Antialias = int64(6)
)

type FillRule uint32

const (
	FillRule_Winding FillRule = int64(0)
	FillRule_EvenOdd FillRule = int64(1)
)

type LineCap uint32

const (
	LineCap_Butt   LineCap = int64(0)
	LineCap_Round  LineCap = int64(1)
	LineCap_Square LineCap = int64(2)
)

type LineJoin uint32

const (
	LineJoin_Miter LineJoin = int64(0)
	LineJoin_Round LineJoin = int64(1)
	LineJoin_Bevel LineJoin = int64(2)
)

type TextClusterFlags uint32

const (
	TextClusterFlags_Backward TextClusterFlags = int64(1)
)

type FontSlant uint32

const (
	FontSlant_Normal  FontSlant = int64(0)
	FontSlant_Italic  FontSlant = int64(1)
	FontSlant_Oblique FontSlant = int64(2)
)

type FontWeight uint32

const (
	FontWeight_Normal FontWeight = int64(0)
	FontWeight_Bold   FontWeight = int64(1)
)

type SubpixelOrder uint32

const (
	SubpixelOrder_Default SubpixelOrder = int64(0)
	SubpixelOrder_Rgb     SubpixelOrder = int64(1)
	SubpixelOrder_Bgr     SubpixelOrder = int64(2)
	SubpixelOrder_Vrgb    SubpixelOrder = int64(3)
	SubpixelOrder_Vbgr    SubpixelOrder = int64(4)
)

type HintStyle uint32

const (
	HintStyle_Default HintStyle = int64(0)
	HintStyle_None    HintStyle = int64(1)
	HintStyle_Slight  HintStyle = int64(2)
	HintStyle_Medium  HintStyle = int64(3)
	HintStyle_Full    HintStyle = int64(4)
)

type HintMetrics uint32

const (
	HintMetrics_Default HintMetrics = int64(0)
	HintMetrics_Off     HintMetrics = int64(1)
	HintMetrics_On      HintMetrics = int64(2)
)

type FontType uint32

const (
	FontType_Toy    FontType = int64(0)
	FontType_Ft     FontType = int64(1)
	FontType_Win32  FontType = int64(2)
	FontType_Quartz FontType = int64(3)
	FontType_User   FontType = int64(4)
)

type PathDataType uint32

const (
	PathDataType_MoveTo    PathDataType = int64(0)
	PathDataType_LineTo    PathDataType = int64(1)
	PathDataType_CurveTo   PathDataType = int64(2)
	PathDataType_ClosePath PathDataType = int64(3)
)

type DeviceType int32

const (
	DeviceType_Drm     DeviceType = int64(0)
	DeviceType_Gl      DeviceType = int64(1)
	DeviceType_Script  DeviceType = int64(2)
	DeviceType_Xcb     DeviceType = int64(3)
	DeviceType_Xlib    DeviceType = int64(4)
	DeviceType_Xml     DeviceType = int64(5)
	DeviceType_Cogl    DeviceType = int64(6)
	DeviceType_Win32   DeviceType = int64(7)
	DeviceType_Invalid DeviceType = int64(-1)
)

type SurfaceType uint32

const (
	SurfaceType_Image         SurfaceType = int64(0)
	SurfaceType_Pdf           SurfaceType = int64(1)
	SurfaceType_Ps            SurfaceType = int64(2)
	SurfaceType_Xlib          SurfaceType = int64(3)
	SurfaceType_Xcb           SurfaceType = int64(4)
	SurfaceType_Glitz         SurfaceType = int64(5)
	SurfaceType_Quartz        SurfaceType = int64(6)
	SurfaceType_Win32         SurfaceType = int64(7)
	SurfaceType_Beos          SurfaceType = int64(8)
	SurfaceType_Directfb      SurfaceType = int64(9)
	SurfaceType_Svg           SurfaceType = int64(10)
	SurfaceType_Os2           SurfaceType = int64(11)
	SurfaceType_Win32Printing SurfaceType = int64(12)
	SurfaceType_QuartzImage   SurfaceType = int64(13)
	SurfaceType_Script        SurfaceType = int64(14)
	SurfaceType_Qt            SurfaceType = int64(15)
	SurfaceType_Recording     SurfaceType = int64(16)
	SurfaceType_Vg            SurfaceType = int64(17)
	SurfaceType_Gl            SurfaceType = int64(18)
	SurfaceType_Drm           SurfaceType = int64(19)
	SurfaceType_Tee           SurfaceType = int64(20)
	SurfaceType_Xml           SurfaceType = int64(21)
	SurfaceType_Skia          SurfaceType = int64(22)
	SurfaceType_Subsurface    SurfaceType = int64(23)
	SurfaceType_Cogl          SurfaceType = int64(24)
)

type Format int32

const (
	Format_Invalid  Format = int64(-1)
	Format_Argb32   Format = int64(0)
	Format_Rgb24    Format = int64(1)
	Format_A8       Format = int64(2)
	Format_A1       Format = int64(3)
	Format_Rgb16565 Format = int64(4)
	Format_Rgb30    Format = int64(5)
)

type PatternType uint32

const (
	PatternType_Solid        PatternType = int64(0)
	PatternType_Surface      PatternType = int64(1)
	PatternType_Linear       PatternType = int64(2)
	PatternType_Radial       PatternType = int64(3)
	PatternType_Mesh         PatternType = int64(4)
	PatternType_RasterSource PatternType = int64(5)
)

type Extend uint32

const (
	Extend_None    Extend = int64(0)
	Extend_Repeat  Extend = int64(1)
	Extend_Reflect Extend = int64(2)
	Extend_Pad     Extend = int64(3)
)

type Filter uint32

const (
	Filter_Fast     Filter = int64(0)
	Filter_Good     Filter = int64(1)
	Filter_Best     Filter = int64(2)
	Filter_Nearest  Filter = int64(3)
	Filter_Bilinear Filter = int64(4)
	Filter_Gaussian Filter = int64(5)
)

type RegionOverlap uint32

const (
	RegionOverlap_In   RegionOverlap = int64(0)
	RegionOverlap_Out  RegionOverlap = int64(1)
	RegionOverlap_Part RegionOverlap = int64(2)
)
