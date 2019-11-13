// Code generated - DO NOT EDIT.

package cairo

// Status is a representation of the C type Status.
type Status int

const (
	Status_Success                 Status = 0
	Status_NoMemory                Status = 1
	Status_InvalidRestore          Status = 2
	Status_InvalidPopGroup         Status = 3
	Status_NoCurrentPoint          Status = 4
	Status_InvalidMatrix           Status = 5
	Status_InvalidStatus           Status = 6
	Status_NullPointer             Status = 7
	Status_InvalidString           Status = 8
	Status_InvalidPathData         Status = 9
	Status_ReadError               Status = 10
	Status_WriteError              Status = 11
	Status_SurfaceFinished         Status = 12
	Status_SurfaceTypeMismatch     Status = 13
	Status_PatternTypeMismatch     Status = 14
	Status_InvalidContent          Status = 15
	Status_InvalidFormat           Status = 16
	Status_InvalidVisual           Status = 17
	Status_FileNotFound            Status = 18
	Status_InvalidDash             Status = 19
	Status_InvalidDscComment       Status = 20
	Status_InvalidIndex            Status = 21
	Status_ClipNotRepresentable    Status = 22
	Status_TempFileError           Status = 23
	Status_InvalidStride           Status = 24
	Status_FontTypeMismatch        Status = 25
	Status_UserFontImmutable       Status = 26
	Status_UserFontError           Status = 27
	Status_NegativeCount           Status = 28
	Status_InvalidClusters         Status = 29
	Status_InvalidSlant            Status = 30
	Status_InvalidWeight           Status = 31
	Status_InvalidSize             Status = 32
	Status_UserFontNotImplemented  Status = 33
	Status_DeviceTypeMismatch      Status = 34
	Status_DeviceError             Status = 35
	Status_InvalidMeshConstruction Status = 36
	Status_DeviceFinished          Status = 37
	Status_Jbig2GlobalMissing      Status = 38
)

// Content is a representation of the C type Content.
type Content int

const (
	Content_Color      Content = 4096
	Content_Alpha      Content = 8192
	Content_ColorAlpha Content = 12288
)

// Operator is a representation of the C type Operator.
type Operator int

const (
	Operator_Clear         Operator = 0
	Operator_Source        Operator = 1
	Operator_Over          Operator = 2
	Operator_In            Operator = 3
	Operator_Out           Operator = 4
	Operator_Atop          Operator = 5
	Operator_Dest          Operator = 6
	Operator_DestOver      Operator = 7
	Operator_DestIn        Operator = 8
	Operator_DestOut       Operator = 9
	Operator_DestAtop      Operator = 10
	Operator_Xor           Operator = 11
	Operator_Add           Operator = 12
	Operator_Saturate      Operator = 13
	Operator_Multiply      Operator = 14
	Operator_Screen        Operator = 15
	Operator_Overlay       Operator = 16
	Operator_Darken        Operator = 17
	Operator_Lighten       Operator = 18
	Operator_ColorDodge    Operator = 19
	Operator_ColorBurn     Operator = 20
	Operator_HardLight     Operator = 21
	Operator_SoftLight     Operator = 22
	Operator_Difference    Operator = 23
	Operator_Exclusion     Operator = 24
	Operator_HslHue        Operator = 25
	Operator_HslSaturation Operator = 26
	Operator_HslColor      Operator = 27
	Operator_HslLuminosity Operator = 28
)

// Antialias is a representation of the C type Antialias.
type Antialias int

const (
	Antialias_Default  Antialias = 0
	Antialias_None     Antialias = 1
	Antialias_Gray     Antialias = 2
	Antialias_Subpixel Antialias = 3
	Antialias_Fast     Antialias = 4
	Antialias_Good     Antialias = 5
	Antialias_Best     Antialias = 6
)

// Fillrule is a representation of the C type FillRule.
type Fillrule int

const (
	Fillrule_Winding Fillrule = 0
	Fillrule_EvenOdd Fillrule = 1
)

// Linecap is a representation of the C type LineCap.
type Linecap int

const (
	Linecap_Butt   Linecap = 0
	Linecap_Round  Linecap = 1
	Linecap_Square Linecap = 2
)

// Linejoin is a representation of the C type LineJoin.
type Linejoin int

const (
	Linejoin_Miter Linejoin = 0
	Linejoin_Round Linejoin = 1
	Linejoin_Bevel Linejoin = 2
)

// Textclusterflags is a representation of the C type TextClusterFlags.
type Textclusterflags int

const (
	Textclusterflags_Backward Textclusterflags = 1
)

// Fontslant is a representation of the C type FontSlant.
type Fontslant int

const (
	Fontslant_Normal  Fontslant = 0
	Fontslant_Italic  Fontslant = 1
	Fontslant_Oblique Fontslant = 2
)

// Fontweight is a representation of the C type FontWeight.
type Fontweight int

const (
	Fontweight_Normal Fontweight = 0
	Fontweight_Bold   Fontweight = 1
)

// Subpixelorder is a representation of the C type SubpixelOrder.
type Subpixelorder int

const (
	Subpixelorder_Default Subpixelorder = 0
	Subpixelorder_Rgb     Subpixelorder = 1
	Subpixelorder_Bgr     Subpixelorder = 2
	Subpixelorder_Vrgb    Subpixelorder = 3
	Subpixelorder_Vbgr    Subpixelorder = 4
)

// Hintstyle is a representation of the C type HintStyle.
type Hintstyle int

const (
	Hintstyle_Default Hintstyle = 0
	Hintstyle_None    Hintstyle = 1
	Hintstyle_Slight  Hintstyle = 2
	Hintstyle_Medium  Hintstyle = 3
	Hintstyle_Full    Hintstyle = 4
)

// Hintmetrics is a representation of the C type HintMetrics.
type Hintmetrics int

const (
	Hintmetrics_Default Hintmetrics = 0
	Hintmetrics_Off     Hintmetrics = 1
	Hintmetrics_On      Hintmetrics = 2
)

// Fonttype is a representation of the C type FontType.
type Fonttype int

const (
	Fonttype_Toy    Fonttype = 0
	Fonttype_Ft     Fonttype = 1
	Fonttype_Win32  Fonttype = 2
	Fonttype_Quartz Fonttype = 3
	Fonttype_User   Fonttype = 4
)

// Pathdatatype is a representation of the C type PathDataType.
type Pathdatatype int

const (
	Pathdatatype_MoveTo    Pathdatatype = 0
	Pathdatatype_LineTo    Pathdatatype = 1
	Pathdatatype_CurveTo   Pathdatatype = 2
	Pathdatatype_ClosePath Pathdatatype = 3
)

// Devicetype is a representation of the C type DeviceType.
type Devicetype int

const (
	Devicetype_Drm     Devicetype = 0
	Devicetype_Gl      Devicetype = 1
	Devicetype_Script  Devicetype = 2
	Devicetype_Xcb     Devicetype = 3
	Devicetype_Xlib    Devicetype = 4
	Devicetype_Xml     Devicetype = 5
	Devicetype_Cogl    Devicetype = 6
	Devicetype_Win32   Devicetype = 7
	Devicetype_Invalid Devicetype = -1
)

// Surfacetype is a representation of the C type SurfaceType.
type Surfacetype int

const (
	Surfacetype_Image         Surfacetype = 0
	Surfacetype_Pdf           Surfacetype = 1
	Surfacetype_Ps            Surfacetype = 2
	Surfacetype_Xlib          Surfacetype = 3
	Surfacetype_Xcb           Surfacetype = 4
	Surfacetype_Glitz         Surfacetype = 5
	Surfacetype_Quartz        Surfacetype = 6
	Surfacetype_Win32         Surfacetype = 7
	Surfacetype_Beos          Surfacetype = 8
	Surfacetype_Directfb      Surfacetype = 9
	Surfacetype_Svg           Surfacetype = 10
	Surfacetype_Os2           Surfacetype = 11
	Surfacetype_Win32Printing Surfacetype = 12
	Surfacetype_QuartzImage   Surfacetype = 13
	Surfacetype_Script        Surfacetype = 14
	Surfacetype_Qt            Surfacetype = 15
	Surfacetype_Recording     Surfacetype = 16
	Surfacetype_Vg            Surfacetype = 17
	Surfacetype_Gl            Surfacetype = 18
	Surfacetype_Drm           Surfacetype = 19
	Surfacetype_Tee           Surfacetype = 20
	Surfacetype_Xml           Surfacetype = 21
	Surfacetype_Skia          Surfacetype = 22
	Surfacetype_Subsurface    Surfacetype = 23
	Surfacetype_Cogl          Surfacetype = 24
)

// Format is a representation of the C type Format.
type Format int

const (
	Format_Invalid  Format = -1
	Format_Argb32   Format = 0
	Format_Rgb24    Format = 1
	Format_A8       Format = 2
	Format_A1       Format = 3
	Format_Rgb16565 Format = 4
	Format_Rgb30    Format = 5
)

// Patterntype is a representation of the C type PatternType.
type Patterntype int

const (
	Patterntype_Solid        Patterntype = 0
	Patterntype_Surface      Patterntype = 1
	Patterntype_Linear       Patterntype = 2
	Patterntype_Radial       Patterntype = 3
	Patterntype_Mesh         Patterntype = 4
	Patterntype_RasterSource Patterntype = 5
)

// Extend is a representation of the C type Extend.
type Extend int

const (
	Extend_None    Extend = 0
	Extend_Repeat  Extend = 1
	Extend_Reflect Extend = 2
	Extend_Pad     Extend = 3
)

// Filter is a representation of the C type Filter.
type Filter int

const (
	Filter_Fast     Filter = 0
	Filter_Good     Filter = 1
	Filter_Best     Filter = 2
	Filter_Nearest  Filter = 3
	Filter_Bilinear Filter = 4
	Filter_Gaussian Filter = 5
)

// Regionoverlap is a representation of the C type RegionOverlap.
type Regionoverlap int

const (
	Regionoverlap_In   Regionoverlap = 0
	Regionoverlap_Out  Regionoverlap = 1
	Regionoverlap_Part Regionoverlap = 2
)
