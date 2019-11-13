// Code generated - DO NOT EDIT.

package cairo

// Status is a representation of the C type Status.
type Status int

const (
	// success
	CairoStatusSuccess Status = 0
	// no_memory
	CairoStatusNoMemory Status = 1
	// invalid_restore
	CairoStatusInvalidRestore Status = 2
	// invalid_pop_group
	CairoStatusInvalidPopGroup Status = 3
	// no_current_point
	CairoStatusNoCurrentPoint Status = 4
	// invalid_matrix
	CairoStatusInvalidMatrix Status = 5
	// invalid_status
	CairoStatusInvalidStatus Status = 6
	// null_pointer
	CairoStatusNullPointer Status = 7
	// invalid_string
	CairoStatusInvalidString Status = 8
	// invalid_path_data
	CairoStatusInvalidPathData Status = 9
	// read_error
	CairoStatusReadError Status = 10
	// write_error
	CairoStatusWriteError Status = 11
	// surface_finished
	CairoStatusSurfaceFinished Status = 12
	// surface_type_mismatch
	CairoStatusSurfaceTypeMismatch Status = 13
	// pattern_type_mismatch
	CairoStatusPatternTypeMismatch Status = 14
	// invalid_content
	CairoStatusInvalidContent Status = 15
	// invalid_format
	CairoStatusInvalidFormat Status = 16
	// invalid_visual
	CairoStatusInvalidVisual Status = 17
	// file_not_found
	CairoStatusFileNotFound Status = 18
	// invalid_dash
	CairoStatusInvalidDash Status = 19
	// invalid_dsc_comment
	CairoStatusInvalidDscComment Status = 20
	// invalid_index
	CairoStatusInvalidIndex Status = 21
	// clip_not_representable
	CairoStatusClipNotRepresentable Status = 22
	// temp_file_error
	CairoStatusTempFileError Status = 23
	// invalid_stride
	CairoStatusInvalidStride Status = 24
	// font_type_mismatch
	CairoStatusFontTypeMismatch Status = 25
	// user_font_immutable
	CairoStatusUserFontImmutable Status = 26
	// user_font_error
	CairoStatusUserFontError Status = 27
	// negative_count
	CairoStatusNegativeCount Status = 28
	// invalid_clusters
	CairoStatusInvalidClusters Status = 29
	// invalid_slant
	CairoStatusInvalidSlant Status = 30
	// invalid_weight
	CairoStatusInvalidWeight Status = 31
	// invalid_size
	CairoStatusInvalidSize Status = 32
	// user_font_not_implemented
	CairoStatusUserFontNotImplemented Status = 33
	// device_type_mismatch
	CairoStatusDeviceTypeMismatch Status = 34
	// device_error
	CairoStatusDeviceError Status = 35
	// invalid_mesh_construction
	CairoStatusInvalidMeshConstruction Status = 36
	// device_finished
	CairoStatusDeviceFinished Status = 37
	// jbig2_global_missing
	CairoStatusJbig2GlobalMissing Status = 38
)

// Content is a representation of the C type Content.
type Content int

const (
	// color
	CairoContentColor Content = 4096
	// alpha
	CairoContentAlpha Content = 8192
	// color_alpha
	CairoContentColorAlpha Content = 12288
)

// Operator is a representation of the C type Operator.
type Operator int

const (
	// clear
	CairoOperatorClear Operator = 0
	// source
	CairoOperatorSource Operator = 1
	// over
	CairoOperatorOver Operator = 2
	// in
	CairoOperatorIn Operator = 3
	// out
	CairoOperatorOut Operator = 4
	// atop
	CairoOperatorAtop Operator = 5
	// dest
	CairoOperatorDest Operator = 6
	// dest_over
	CairoOperatorDestOver Operator = 7
	// dest_in
	CairoOperatorDestIn Operator = 8
	// dest_out
	CairoOperatorDestOut Operator = 9
	// dest_atop
	CairoOperatorDestAtop Operator = 10
	// xor
	CairoOperatorXor Operator = 11
	// add
	CairoOperatorAdd Operator = 12
	// saturate
	CairoOperatorSaturate Operator = 13
	// multiply
	CairoOperatorMultiply Operator = 14
	// screen
	CairoOperatorScreen Operator = 15
	// overlay
	CairoOperatorOverlay Operator = 16
	// darken
	CairoOperatorDarken Operator = 17
	// lighten
	CairoOperatorLighten Operator = 18
	// color_dodge
	CairoOperatorColorDodge Operator = 19
	// color_burn
	CairoOperatorColorBurn Operator = 20
	// hard_light
	CairoOperatorHardLight Operator = 21
	// soft_light
	CairoOperatorSoftLight Operator = 22
	// difference
	CairoOperatorDifference Operator = 23
	// exclusion
	CairoOperatorExclusion Operator = 24
	// hsl_hue
	CairoOperatorHslHue Operator = 25
	// hsl_saturation
	CairoOperatorHslSaturation Operator = 26
	// hsl_color
	CairoOperatorHslColor Operator = 27
	// hsl_luminosity
	CairoOperatorHslLuminosity Operator = 28
)

// Antialias is a representation of the C type Antialias.
type Antialias int

const (
	// default
	CairoAntialiasDefault Antialias = 0
	// none
	CairoAntialiasNone Antialias = 1
	// gray
	CairoAntialiasGray Antialias = 2
	// subpixel
	CairoAntialiasSubpixel Antialias = 3
	// fast
	CairoAntialiasFast Antialias = 4
	// good
	CairoAntialiasGood Antialias = 5
	// best
	CairoAntialiasBest Antialias = 6
)

// Fillrule is a representation of the C type FillRule.
type Fillrule int

const (
	// winding
	CairoFillRuleWinding Fillrule = 0
	// even_odd
	CairoFillRuleEvenOdd Fillrule = 1
)

// Linecap is a representation of the C type LineCap.
type Linecap int

const (
	// butt
	CairoLineCapButt Linecap = 0
	// round
	CairoLineCapRound Linecap = 1
	// square
	CairoLineCapSquare Linecap = 2
)

// Linejoin is a representation of the C type LineJoin.
type Linejoin int

const (
	// miter
	CairoLineJoinMiter Linejoin = 0
	// round
	CairoLineJoinRound Linejoin = 1
	// bevel
	CairoLineJoinBevel Linejoin = 2
)

// Textclusterflags is a representation of the C type TextClusterFlags.
type Textclusterflags int

const (
	// backward
	CairoTextClusterFlagBackward Textclusterflags = 1
)

// Fontslant is a representation of the C type FontSlant.
type Fontslant int

const (
	// normal
	CairoFontSlantNormal Fontslant = 0
	// italic
	CairoFontSlantItalic Fontslant = 1
	// oblique
	CairoFontSlantOblique Fontslant = 2
)

// Fontweight is a representation of the C type FontWeight.
type Fontweight int

const (
	// normal
	CairoFontWeightNormal Fontweight = 0
	// bold
	CairoFontWeightBold Fontweight = 1
)

// Subpixelorder is a representation of the C type SubpixelOrder.
type Subpixelorder int

const (
	// default
	CairoSubpixelOrderDefault Subpixelorder = 0
	// rgb
	CairoSubpixelOrderRgb Subpixelorder = 1
	// bgr
	CairoSubpixelOrderBgr Subpixelorder = 2
	// vrgb
	CairoSubpixelOrderVrgb Subpixelorder = 3
	// vbgr
	CairoSubpixelOrderVbgr Subpixelorder = 4
)

// Hintstyle is a representation of the C type HintStyle.
type Hintstyle int

const (
	// default
	CairoHintStyleDefault Hintstyle = 0
	// none
	CairoHintStyleNone Hintstyle = 1
	// slight
	CairoHintStyleSlight Hintstyle = 2
	// medium
	CairoHintStyleMedium Hintstyle = 3
	// full
	CairoHintStyleFull Hintstyle = 4
)

// Hintmetrics is a representation of the C type HintMetrics.
type Hintmetrics int

const (
	// default
	CairoHintMetricsDefault Hintmetrics = 0
	// off
	CairoHintMetricsOff Hintmetrics = 1
	// on
	CairoHintMetricsOn Hintmetrics = 2
)

// Fonttype is a representation of the C type FontType.
type Fonttype int

const (
	// toy
	CairoFontTypeToy Fonttype = 0
	// ft
	CairoFontTypeFt Fonttype = 1
	// win32
	CairoFontTypeWin32 Fonttype = 2
	// quartz
	CairoFontTypeQuartz Fonttype = 3
	// user
	CairoFontTypeUser Fonttype = 4
)

// Pathdatatype is a representation of the C type PathDataType.
type Pathdatatype int

const (
	// move_to
	CairoPathMoveTo Pathdatatype = 0
	// line_to
	CairoPathLineTo Pathdatatype = 1
	// curve_to
	CairoPathCurveTo Pathdatatype = 2
	// close_path
	CairoPathClosePath Pathdatatype = 3
)

// Devicetype is a representation of the C type DeviceType.
type Devicetype int

const (
	// drm
	CairoDeviceTypeDrm Devicetype = 0
	// gl
	CairoDeviceTypeGl Devicetype = 1
	// script
	CairoDeviceTypeScript Devicetype = 2
	// xcb
	CairoDeviceTypeXcb Devicetype = 3
	// xlib
	CairoDeviceTypeXlib Devicetype = 4
	// xml
	CairoDeviceTypeXml Devicetype = 5
	// cogl
	CairoDeviceTypeCogl Devicetype = 6
	// win32
	CairoDeviceTypeWin32 Devicetype = 7
	// invalid
	CairoDeviceTypeInvalid Devicetype = -1
)

// Surfacetype is a representation of the C type SurfaceType.
type Surfacetype int

const (
	// image
	CairoSurfaceTypeImage Surfacetype = 0
	// pdf
	CairoSurfaceTypePdf Surfacetype = 1
	// ps
	CairoSurfaceTypePs Surfacetype = 2
	// xlib
	CairoSurfaceTypeXlib Surfacetype = 3
	// xcb
	CairoSurfaceTypeXcb Surfacetype = 4
	// glitz
	CairoSurfaceTypeGlitz Surfacetype = 5
	// quartz
	CairoSurfaceTypeQuartz Surfacetype = 6
	// win32
	CairoSurfaceTypeWin32 Surfacetype = 7
	// beos
	CairoSurfaceTypeBeos Surfacetype = 8
	// directfb
	CairoSurfaceTypeDirectfb Surfacetype = 9
	// svg
	CairoSurfaceTypeSvg Surfacetype = 10
	// os2
	CairoSurfaceTypeOs2 Surfacetype = 11
	// win32_printing
	CairoSurfaceTypeWin32Printing Surfacetype = 12
	// quartz_image
	CairoSurfaceTypeQuartzImage Surfacetype = 13
	// script
	CairoSurfaceTypeScript Surfacetype = 14
	// qt
	CairoSurfaceTypeQt Surfacetype = 15
	// recording
	CairoSurfaceTypeRecording Surfacetype = 16
	// vg
	CairoSurfaceTypeVg Surfacetype = 17
	// gl
	CairoSurfaceTypeGl Surfacetype = 18
	// drm
	CairoSurfaceTypeDrm Surfacetype = 19
	// tee
	CairoSurfaceTypeTee Surfacetype = 20
	// xml
	CairoSurfaceTypeXml Surfacetype = 21
	// skia
	CairoSurfaceTypeSkia Surfacetype = 22
	// subsurface
	CairoSurfaceTypeSubsurface Surfacetype = 23
	// cogl
	CairoSurfaceTypeCogl Surfacetype = 24
)

// Format is a representation of the C type Format.
type Format int

const (
	// invalid
	CairoFormatInvalid Format = -1
	// argb32
	CairoFormatArgb32 Format = 0
	// rgb24
	CairoFormatRgb24 Format = 1
	// a8
	CairoFormatA8 Format = 2
	// a1
	CairoFormatA1 Format = 3
	// rgb16_565
	CairoFormatRgb16565 Format = 4
	// rgb30
	CairoFormatRgb30 Format = 5
)

// Patterntype is a representation of the C type PatternType.
type Patterntype int

const (
	// solid
	CairoPatternTypeSolid Patterntype = 0
	// surface
	CairoPatternTypeSurface Patterntype = 1
	// linear
	CairoPatternTypeLinear Patterntype = 2
	// radial
	CairoPatternTypeRadial Patterntype = 3
	// mesh
	CairoPatternTypeMesh Patterntype = 4
	// raster_source
	CairoPatternTypeRasterSource Patterntype = 5
)

// Extend is a representation of the C type Extend.
type Extend int

const (
	// none
	CairoExtendNone Extend = 0
	// repeat
	CairoExtendRepeat Extend = 1
	// reflect
	CairoExtendReflect Extend = 2
	// pad
	CairoExtendPad Extend = 3
)

// Filter is a representation of the C type Filter.
type Filter int

const (
	// fast
	CairoFilterFast Filter = 0
	// good
	CairoFilterGood Filter = 1
	// best
	CairoFilterBest Filter = 2
	// nearest
	CairoFilterNearest Filter = 3
	// bilinear
	CairoFilterBilinear Filter = 4
	// gaussian
	CairoFilterGaussian Filter = 5
)

// Regionoverlap is a representation of the C type RegionOverlap.
type Regionoverlap int

const (
	// in
	CairoRegionOverlapIn Regionoverlap = 0
	// out
	CairoRegionOverlapOut Regionoverlap = 1
	// part
	CairoRegionOverlapPart Regionoverlap = 2
)
