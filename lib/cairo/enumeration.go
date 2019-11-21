// Code generated - DO NOT EDIT.

package cairo

// Status is a representation of the C type cairo_status_t.
//
type Status int

const (
	// Status_Success is a representation of the C type CAIRO_STATUS_SUCCESS.
	Status_Success Status = 0
	// Status_NoMemory is a representation of the C type CAIRO_STATUS_NO_MEMORY.
	Status_NoMemory Status = 1
	// Status_InvalidRestore is a representation of the C type CAIRO_STATUS_INVALID_RESTORE.
	Status_InvalidRestore Status = 2
	// Status_InvalidPopGroup is a representation of the C type CAIRO_STATUS_INVALID_POP_GROUP.
	Status_InvalidPopGroup Status = 3
	// Status_NoCurrentPoint is a representation of the C type CAIRO_STATUS_NO_CURRENT_POINT.
	Status_NoCurrentPoint Status = 4
	// Status_InvalidMatrix is a representation of the C type CAIRO_STATUS_INVALID_MATRIX.
	Status_InvalidMatrix Status = 5
	// Status_InvalidStatus is a representation of the C type CAIRO_STATUS_INVALID_STATUS.
	Status_InvalidStatus Status = 6
	// Status_NullPointer is a representation of the C type CAIRO_STATUS_NULL_POINTER.
	Status_NullPointer Status = 7
	// Status_InvalidString is a representation of the C type CAIRO_STATUS_INVALID_STRING.
	Status_InvalidString Status = 8
	// Status_InvalidPathData is a representation of the C type CAIRO_STATUS_INVALID_PATH_DATA.
	Status_InvalidPathData Status = 9
	// Status_ReadError is a representation of the C type CAIRO_STATUS_READ_ERROR.
	Status_ReadError Status = 10
	// Status_WriteError is a representation of the C type CAIRO_STATUS_WRITE_ERROR.
	Status_WriteError Status = 11
	// Status_SurfaceFinished is a representation of the C type CAIRO_STATUS_SURFACE_FINISHED.
	Status_SurfaceFinished Status = 12
	// Status_SurfaceTypeMismatch is a representation of the C type CAIRO_STATUS_SURFACE_TYPE_MISMATCH.
	Status_SurfaceTypeMismatch Status = 13
	// Status_PatternTypeMismatch is a representation of the C type CAIRO_STATUS_PATTERN_TYPE_MISMATCH.
	Status_PatternTypeMismatch Status = 14
	// Status_InvalidContent is a representation of the C type CAIRO_STATUS_INVALID_CONTENT.
	Status_InvalidContent Status = 15
	// Status_InvalidFormat is a representation of the C type CAIRO_STATUS_INVALID_FORMAT.
	Status_InvalidFormat Status = 16
	// Status_InvalidVisual is a representation of the C type CAIRO_STATUS_INVALID_VISUAL.
	Status_InvalidVisual Status = 17
	// Status_FileNotFound is a representation of the C type CAIRO_STATUS_FILE_NOT_FOUND.
	Status_FileNotFound Status = 18
	// Status_InvalidDash is a representation of the C type CAIRO_STATUS_INVALID_DASH.
	Status_InvalidDash Status = 19
	// Status_InvalidDscComment is a representation of the C type CAIRO_STATUS_INVALID_DSC_COMMENT.
	Status_InvalidDscComment Status = 20
	// Status_InvalidIndex is a representation of the C type CAIRO_STATUS_INVALID_INDEX.
	Status_InvalidIndex Status = 21
	// Status_ClipNotRepresentable is a representation of the C type CAIRO_STATUS_CLIP_NOT_REPRESENTABLE.
	Status_ClipNotRepresentable Status = 22
	// Status_TempFileError is a representation of the C type CAIRO_STATUS_TEMP_FILE_ERROR.
	Status_TempFileError Status = 23
	// Status_InvalidStride is a representation of the C type CAIRO_STATUS_INVALID_STRIDE.
	Status_InvalidStride Status = 24
	// Status_FontTypeMismatch is a representation of the C type CAIRO_STATUS_FONT_TYPE_MISMATCH.
	Status_FontTypeMismatch Status = 25
	// Status_UserFontImmutable is a representation of the C type CAIRO_STATUS_USER_FONT_IMMUTABLE.
	Status_UserFontImmutable Status = 26
	// Status_UserFontError is a representation of the C type CAIRO_STATUS_USER_FONT_ERROR.
	Status_UserFontError Status = 27
	// Status_NegativeCount is a representation of the C type CAIRO_STATUS_NEGATIVE_COUNT.
	Status_NegativeCount Status = 28
	// Status_InvalidClusters is a representation of the C type CAIRO_STATUS_INVALID_CLUSTERS.
	Status_InvalidClusters Status = 29
	// Status_InvalidSlant is a representation of the C type CAIRO_STATUS_INVALID_SLANT.
	Status_InvalidSlant Status = 30
	// Status_InvalidWeight is a representation of the C type CAIRO_STATUS_INVALID_WEIGHT.
	Status_InvalidWeight Status = 31
	// Status_InvalidSize is a representation of the C type CAIRO_STATUS_INVALID_SIZE.
	Status_InvalidSize Status = 32
	// Status_UserFontNotImplemented is a representation of the C type CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED.
	Status_UserFontNotImplemented Status = 33
	// Status_DeviceTypeMismatch is a representation of the C type CAIRO_STATUS_DEVICE_TYPE_MISMATCH.
	Status_DeviceTypeMismatch Status = 34
	// Status_DeviceError is a representation of the C type CAIRO_STATUS_DEVICE_ERROR.
	Status_DeviceError Status = 35
	// Status_InvalidMeshConstruction is a representation of the C type CAIRO_STATUS_INVALID_MESH_CONSTRUCTION.
	Status_InvalidMeshConstruction Status = 36
	// Status_DeviceFinished is a representation of the C type CAIRO_STATUS_DEVICE_FINISHED.
	Status_DeviceFinished Status = 37
	// Status_Jbig2GlobalMissing is a representation of the C type CAIRO_STATUS_JBIG2_GLOBAL_MISSING.
	Status_Jbig2GlobalMissing Status = 38
)

// Content is a representation of the C type cairo_content_t.
//
type Content int

const (
	// Content_Color is a representation of the C type CAIRO_CONTENT_COLOR.
	Content_Color Content = 4096
	// Content_Alpha is a representation of the C type CAIRO_CONTENT_ALPHA.
	Content_Alpha Content = 8192
	// Content_ColorAlpha is a representation of the C type CAIRO_CONTENT_COLOR_ALPHA.
	Content_ColorAlpha Content = 12288
)

// Operator is a representation of the C type cairo_operator_t.
//
type Operator int

const (
	// Operator_Clear is a representation of the C type CAIRO_OPERATOR_CLEAR.
	Operator_Clear Operator = 0
	// Operator_Source is a representation of the C type CAIRO_OPERATOR_SOURCE.
	Operator_Source Operator = 1
	// Operator_Over is a representation of the C type CAIRO_OPERATOR_OVER.
	Operator_Over Operator = 2
	// Operator_In is a representation of the C type CAIRO_OPERATOR_IN.
	Operator_In Operator = 3
	// Operator_Out is a representation of the C type CAIRO_OPERATOR_OUT.
	Operator_Out Operator = 4
	// Operator_Atop is a representation of the C type CAIRO_OPERATOR_ATOP.
	Operator_Atop Operator = 5
	// Operator_Dest is a representation of the C type CAIRO_OPERATOR_DEST.
	Operator_Dest Operator = 6
	// Operator_DestOver is a representation of the C type CAIRO_OPERATOR_DEST_OVER.
	Operator_DestOver Operator = 7
	// Operator_DestIn is a representation of the C type CAIRO_OPERATOR_DEST_IN.
	Operator_DestIn Operator = 8
	// Operator_DestOut is a representation of the C type CAIRO_OPERATOR_DEST_OUT.
	Operator_DestOut Operator = 9
	// Operator_DestAtop is a representation of the C type CAIRO_OPERATOR_DEST_ATOP.
	Operator_DestAtop Operator = 10
	// Operator_Xor is a representation of the C type CAIRO_OPERATOR_XOR.
	Operator_Xor Operator = 11
	// Operator_Add is a representation of the C type CAIRO_OPERATOR_ADD.
	Operator_Add Operator = 12
	// Operator_Saturate is a representation of the C type CAIRO_OPERATOR_SATURATE.
	Operator_Saturate Operator = 13
	// Operator_Multiply is a representation of the C type CAIRO_OPERATOR_MULTIPLY.
	Operator_Multiply Operator = 14
	// Operator_Screen is a representation of the C type CAIRO_OPERATOR_SCREEN.
	Operator_Screen Operator = 15
	// Operator_Overlay is a representation of the C type CAIRO_OPERATOR_OVERLAY.
	Operator_Overlay Operator = 16
	// Operator_Darken is a representation of the C type CAIRO_OPERATOR_DARKEN.
	Operator_Darken Operator = 17
	// Operator_Lighten is a representation of the C type CAIRO_OPERATOR_LIGHTEN.
	Operator_Lighten Operator = 18
	// Operator_ColorDodge is a representation of the C type CAIRO_OPERATOR_COLOR_DODGE.
	Operator_ColorDodge Operator = 19
	// Operator_ColorBurn is a representation of the C type CAIRO_OPERATOR_COLOR_BURN.
	Operator_ColorBurn Operator = 20
	// Operator_HardLight is a representation of the C type CAIRO_OPERATOR_HARD_LIGHT.
	Operator_HardLight Operator = 21
	// Operator_SoftLight is a representation of the C type CAIRO_OPERATOR_SOFT_LIGHT.
	Operator_SoftLight Operator = 22
	// Operator_Difference is a representation of the C type CAIRO_OPERATOR_DIFFERENCE.
	Operator_Difference Operator = 23
	// Operator_Exclusion is a representation of the C type CAIRO_OPERATOR_EXCLUSION.
	Operator_Exclusion Operator = 24
	// Operator_HslHue is a representation of the C type CAIRO_OPERATOR_HSL_HUE.
	Operator_HslHue Operator = 25
	// Operator_HslSaturation is a representation of the C type CAIRO_OPERATOR_HSL_SATURATION.
	Operator_HslSaturation Operator = 26
	// Operator_HslColor is a representation of the C type CAIRO_OPERATOR_HSL_COLOR.
	Operator_HslColor Operator = 27
	// Operator_HslLuminosity is a representation of the C type CAIRO_OPERATOR_HSL_LUMINOSITY.
	Operator_HslLuminosity Operator = 28
)

// Antialias is a representation of the C type cairo_antialias_t.
//
type Antialias int

const (
	// Antialias_Default is a representation of the C type CAIRO_ANTIALIAS_DEFAULT.
	Antialias_Default Antialias = 0
	// Antialias_None is a representation of the C type CAIRO_ANTIALIAS_NONE.
	Antialias_None Antialias = 1
	// Antialias_Gray is a representation of the C type CAIRO_ANTIALIAS_GRAY.
	Antialias_Gray Antialias = 2
	// Antialias_Subpixel is a representation of the C type CAIRO_ANTIALIAS_SUBPIXEL.
	Antialias_Subpixel Antialias = 3
	// Antialias_Fast is a representation of the C type CAIRO_ANTIALIAS_FAST.
	Antialias_Fast Antialias = 4
	// Antialias_Good is a representation of the C type CAIRO_ANTIALIAS_GOOD.
	Antialias_Good Antialias = 5
	// Antialias_Best is a representation of the C type CAIRO_ANTIALIAS_BEST.
	Antialias_Best Antialias = 6
)

// FillRule is a representation of the C type cairo_fill_rule_t.
//
type FillRule int

const (
	// FillRule_Winding is a representation of the C type CAIRO_FILL_RULE_WINDING.
	FillRule_Winding FillRule = 0
	// FillRule_EvenOdd is a representation of the C type CAIRO_FILL_RULE_EVEN_ODD.
	FillRule_EvenOdd FillRule = 1
)

// LineCap is a representation of the C type cairo_line_cap_t.
//
type LineCap int

const (
	// LineCap_Butt is a representation of the C type CAIRO_LINE_CAP_BUTT.
	LineCap_Butt LineCap = 0
	// LineCap_Round is a representation of the C type CAIRO_LINE_CAP_ROUND.
	LineCap_Round LineCap = 1
	// LineCap_Square is a representation of the C type CAIRO_LINE_CAP_SQUARE.
	LineCap_Square LineCap = 2
)

// LineJoin is a representation of the C type cairo_line_join_t.
//
type LineJoin int

const (
	// LineJoin_Miter is a representation of the C type CAIRO_LINE_JOIN_MITER.
	LineJoin_Miter LineJoin = 0
	// LineJoin_Round is a representation of the C type CAIRO_LINE_JOIN_ROUND.
	LineJoin_Round LineJoin = 1
	// LineJoin_Bevel is a representation of the C type CAIRO_LINE_JOIN_BEVEL.
	LineJoin_Bevel LineJoin = 2
)

// TextClusterFlags is a representation of the C type cairo_text_cluster_flags_t.
//
type TextClusterFlags int

const (
	// TextClusterFlags_Backward is a representation of the C type CAIRO_TEXT_CLUSTER_FLAG_BACKWARD.
	TextClusterFlags_Backward TextClusterFlags = 1
)

// FontSlant is a representation of the C type cairo_font_slant_t.
//
type FontSlant int

const (
	// FontSlant_Normal is a representation of the C type CAIRO_FONT_SLANT_NORMAL.
	FontSlant_Normal FontSlant = 0
	// FontSlant_Italic is a representation of the C type CAIRO_FONT_SLANT_ITALIC.
	FontSlant_Italic FontSlant = 1
	// FontSlant_Oblique is a representation of the C type CAIRO_FONT_SLANT_OBLIQUE.
	FontSlant_Oblique FontSlant = 2
)

// FontWeight is a representation of the C type cairo_font_weight_t.
//
type FontWeight int

const (
	// FontWeight_Normal is a representation of the C type CAIRO_FONT_WEIGHT_NORMAL.
	FontWeight_Normal FontWeight = 0
	// FontWeight_Bold is a representation of the C type CAIRO_FONT_WEIGHT_BOLD.
	FontWeight_Bold FontWeight = 1
)

// SubpixelOrder is a representation of the C type cairo_subpixel_order_t.
//
type SubpixelOrder int

const (
	// SubpixelOrder_Default is a representation of the C type CAIRO_SUBPIXEL_ORDER_DEFAULT.
	SubpixelOrder_Default SubpixelOrder = 0
	// SubpixelOrder_Rgb is a representation of the C type CAIRO_SUBPIXEL_ORDER_RGB.
	SubpixelOrder_Rgb SubpixelOrder = 1
	// SubpixelOrder_Bgr is a representation of the C type CAIRO_SUBPIXEL_ORDER_BGR.
	SubpixelOrder_Bgr SubpixelOrder = 2
	// SubpixelOrder_Vrgb is a representation of the C type CAIRO_SUBPIXEL_ORDER_VRGB.
	SubpixelOrder_Vrgb SubpixelOrder = 3
	// SubpixelOrder_Vbgr is a representation of the C type CAIRO_SUBPIXEL_ORDER_VBGR.
	SubpixelOrder_Vbgr SubpixelOrder = 4
)

// HintStyle is a representation of the C type cairo_hint_style_t.
//
type HintStyle int

const (
	// HintStyle_Default is a representation of the C type CAIRO_HINT_STYLE_DEFAULT.
	HintStyle_Default HintStyle = 0
	// HintStyle_None is a representation of the C type CAIRO_HINT_STYLE_NONE.
	HintStyle_None HintStyle = 1
	// HintStyle_Slight is a representation of the C type CAIRO_HINT_STYLE_SLIGHT.
	HintStyle_Slight HintStyle = 2
	// HintStyle_Medium is a representation of the C type CAIRO_HINT_STYLE_MEDIUM.
	HintStyle_Medium HintStyle = 3
	// HintStyle_Full is a representation of the C type CAIRO_HINT_STYLE_FULL.
	HintStyle_Full HintStyle = 4
)

// HintMetrics is a representation of the C type cairo_hint_metrics_t.
//
type HintMetrics int

const (
	// HintMetrics_Default is a representation of the C type CAIRO_HINT_METRICS_DEFAULT.
	HintMetrics_Default HintMetrics = 0
	// HintMetrics_Off is a representation of the C type CAIRO_HINT_METRICS_OFF.
	HintMetrics_Off HintMetrics = 1
	// HintMetrics_On is a representation of the C type CAIRO_HINT_METRICS_ON.
	HintMetrics_On HintMetrics = 2
)

// FontType is a representation of the C type cairo_font_type_t.
//
type FontType int

const (
	// FontType_Toy is a representation of the C type CAIRO_FONT_TYPE_TOY.
	FontType_Toy FontType = 0
	// FontType_Ft is a representation of the C type CAIRO_FONT_TYPE_FT.
	FontType_Ft FontType = 1
	// FontType_Win32 is a representation of the C type CAIRO_FONT_TYPE_WIN32.
	FontType_Win32 FontType = 2
	// FontType_Quartz is a representation of the C type CAIRO_FONT_TYPE_QUARTZ.
	FontType_Quartz FontType = 3
	// FontType_User is a representation of the C type CAIRO_FONT_TYPE_USER.
	FontType_User FontType = 4
)

// PathDataType is a representation of the C type cairo_path_data_type_t.
//
type PathDataType int

const (
	// PathDataType_MoveTo is a representation of the C type CAIRO_PATH_MOVE_TO.
	PathDataType_MoveTo PathDataType = 0
	// PathDataType_LineTo is a representation of the C type CAIRO_PATH_LINE_TO.
	PathDataType_LineTo PathDataType = 1
	// PathDataType_CurveTo is a representation of the C type CAIRO_PATH_CURVE_TO.
	PathDataType_CurveTo PathDataType = 2
	// PathDataType_ClosePath is a representation of the C type CAIRO_PATH_CLOSE_PATH.
	PathDataType_ClosePath PathDataType = 3
)

// DeviceType is a representation of the C type cairo_device_type_t.
//
type DeviceType int

const (
	// DeviceType_Drm is a representation of the C type CAIRO_DEVICE_TYPE_DRM.
	DeviceType_Drm DeviceType = 0
	// DeviceType_Gl is a representation of the C type CAIRO_DEVICE_TYPE_GL.
	DeviceType_Gl DeviceType = 1
	// DeviceType_Script is a representation of the C type CAIRO_DEVICE_TYPE_SCRIPT.
	DeviceType_Script DeviceType = 2
	// DeviceType_Xcb is a representation of the C type CAIRO_DEVICE_TYPE_XCB.
	DeviceType_Xcb DeviceType = 3
	// DeviceType_Xlib is a representation of the C type CAIRO_DEVICE_TYPE_XLIB.
	DeviceType_Xlib DeviceType = 4
	// DeviceType_Xml is a representation of the C type CAIRO_DEVICE_TYPE_XML.
	DeviceType_Xml DeviceType = 5
	// DeviceType_Cogl is a representation of the C type CAIRO_DEVICE_TYPE_COGL.
	DeviceType_Cogl DeviceType = 6
	// DeviceType_Win32 is a representation of the C type CAIRO_DEVICE_TYPE_WIN32.
	DeviceType_Win32 DeviceType = 7
	// DeviceType_Invalid is a representation of the C type CAIRO_DEVICE_TYPE_INVALID.
	DeviceType_Invalid DeviceType = -1
)

// SurfaceType is a representation of the C type cairo_surface_type_t.
//
type SurfaceType int

const (
	// SurfaceType_Image is a representation of the C type CAIRO_SURFACE_TYPE_IMAGE.
	SurfaceType_Image SurfaceType = 0
	// SurfaceType_Pdf is a representation of the C type CAIRO_SURFACE_TYPE_PDF.
	SurfaceType_Pdf SurfaceType = 1
	// SurfaceType_Ps is a representation of the C type CAIRO_SURFACE_TYPE_PS.
	SurfaceType_Ps SurfaceType = 2
	// SurfaceType_Xlib is a representation of the C type CAIRO_SURFACE_TYPE_XLIB.
	SurfaceType_Xlib SurfaceType = 3
	// SurfaceType_Xcb is a representation of the C type CAIRO_SURFACE_TYPE_XCB.
	SurfaceType_Xcb SurfaceType = 4
	// SurfaceType_Glitz is a representation of the C type CAIRO_SURFACE_TYPE_GLITZ.
	SurfaceType_Glitz SurfaceType = 5
	// SurfaceType_Quartz is a representation of the C type CAIRO_SURFACE_TYPE_QUARTZ.
	SurfaceType_Quartz SurfaceType = 6
	// SurfaceType_Win32 is a representation of the C type CAIRO_SURFACE_TYPE_WIN32.
	SurfaceType_Win32 SurfaceType = 7
	// SurfaceType_Beos is a representation of the C type CAIRO_SURFACE_TYPE_BEOS.
	SurfaceType_Beos SurfaceType = 8
	// SurfaceType_Directfb is a representation of the C type CAIRO_SURFACE_TYPE_DIRECTFB.
	SurfaceType_Directfb SurfaceType = 9
	// SurfaceType_Svg is a representation of the C type CAIRO_SURFACE_TYPE_SVG.
	SurfaceType_Svg SurfaceType = 10
	// SurfaceType_Os2 is a representation of the C type CAIRO_SURFACE_TYPE_OS2.
	SurfaceType_Os2 SurfaceType = 11
	// SurfaceType_Win32Printing is a representation of the C type CAIRO_SURFACE_TYPE_WIN32_PRINTING.
	SurfaceType_Win32Printing SurfaceType = 12
	// SurfaceType_QuartzImage is a representation of the C type CAIRO_SURFACE_TYPE_QUARTZ_IMAGE.
	SurfaceType_QuartzImage SurfaceType = 13
	// SurfaceType_Script is a representation of the C type CAIRO_SURFACE_TYPE_SCRIPT.
	SurfaceType_Script SurfaceType = 14
	// SurfaceType_Qt is a representation of the C type CAIRO_SURFACE_TYPE_QT.
	SurfaceType_Qt SurfaceType = 15
	// SurfaceType_Recording is a representation of the C type CAIRO_SURFACE_TYPE_RECORDING.
	SurfaceType_Recording SurfaceType = 16
	// SurfaceType_Vg is a representation of the C type CAIRO_SURFACE_TYPE_VG.
	SurfaceType_Vg SurfaceType = 17
	// SurfaceType_Gl is a representation of the C type CAIRO_SURFACE_TYPE_GL.
	SurfaceType_Gl SurfaceType = 18
	// SurfaceType_Drm is a representation of the C type CAIRO_SURFACE_TYPE_DRM.
	SurfaceType_Drm SurfaceType = 19
	// SurfaceType_Tee is a representation of the C type CAIRO_SURFACE_TYPE_TEE.
	SurfaceType_Tee SurfaceType = 20
	// SurfaceType_Xml is a representation of the C type CAIRO_SURFACE_TYPE_XML.
	SurfaceType_Xml SurfaceType = 21
	// SurfaceType_Skia is a representation of the C type CAIRO_SURFACE_TYPE_SKIA.
	SurfaceType_Skia SurfaceType = 22
	// SurfaceType_Subsurface is a representation of the C type CAIRO_SURFACE_TYPE_SUBSURFACE.
	SurfaceType_Subsurface SurfaceType = 23
	// SurfaceType_Cogl is a representation of the C type CAIRO_SURFACE_TYPE_COGL.
	SurfaceType_Cogl SurfaceType = 24
)

// Format is a representation of the C type cairo_format_t.
//
type Format int

const (
	// Format_Invalid is a representation of the C type CAIRO_FORMAT_INVALID.
	Format_Invalid Format = -1
	// Format_Argb32 is a representation of the C type CAIRO_FORMAT_ARGB32.
	Format_Argb32 Format = 0
	// Format_Rgb24 is a representation of the C type CAIRO_FORMAT_RGB24.
	Format_Rgb24 Format = 1
	// Format_A8 is a representation of the C type CAIRO_FORMAT_A8.
	Format_A8 Format = 2
	// Format_A1 is a representation of the C type CAIRO_FORMAT_A1.
	Format_A1 Format = 3
	// Format_Rgb16565 is a representation of the C type CAIRO_FORMAT_RGB16_565.
	Format_Rgb16565 Format = 4
	// Format_Rgb30 is a representation of the C type CAIRO_FORMAT_RGB30.
	Format_Rgb30 Format = 5
)

// PatternType is a representation of the C type cairo_pattern_type_t.
//
type PatternType int

const (
	// PatternType_Solid is a representation of the C type CAIRO_PATTERN_TYPE_SOLID.
	PatternType_Solid PatternType = 0
	// PatternType_Surface is a representation of the C type CAIRO_PATTERN_TYPE_SURFACE.
	PatternType_Surface PatternType = 1
	// PatternType_Linear is a representation of the C type CAIRO_PATTERN_TYPE_LINEAR.
	PatternType_Linear PatternType = 2
	// PatternType_Radial is a representation of the C type CAIRO_PATTERN_TYPE_RADIAL.
	PatternType_Radial PatternType = 3
	// PatternType_Mesh is a representation of the C type CAIRO_PATTERN_TYPE_MESH.
	PatternType_Mesh PatternType = 4
	// PatternType_RasterSource is a representation of the C type CAIRO_PATTERN_TYPE_RASTER_SOURCE.
	PatternType_RasterSource PatternType = 5
)

// Extend is a representation of the C type cairo_extend_t.
//
type Extend int

const (
	// Extend_None is a representation of the C type CAIRO_EXTEND_NONE.
	Extend_None Extend = 0
	// Extend_Repeat is a representation of the C type CAIRO_EXTEND_REPEAT.
	Extend_Repeat Extend = 1
	// Extend_Reflect is a representation of the C type CAIRO_EXTEND_REFLECT.
	Extend_Reflect Extend = 2
	// Extend_Pad is a representation of the C type CAIRO_EXTEND_PAD.
	Extend_Pad Extend = 3
)

// Filter is a representation of the C type cairo_filter_t.
//
type Filter int

const (
	// Filter_Fast is a representation of the C type CAIRO_FILTER_FAST.
	Filter_Fast Filter = 0
	// Filter_Good is a representation of the C type CAIRO_FILTER_GOOD.
	Filter_Good Filter = 1
	// Filter_Best is a representation of the C type CAIRO_FILTER_BEST.
	Filter_Best Filter = 2
	// Filter_Nearest is a representation of the C type CAIRO_FILTER_NEAREST.
	Filter_Nearest Filter = 3
	// Filter_Bilinear is a representation of the C type CAIRO_FILTER_BILINEAR.
	Filter_Bilinear Filter = 4
	// Filter_Gaussian is a representation of the C type CAIRO_FILTER_GAUSSIAN.
	Filter_Gaussian Filter = 5
)

// RegionOverlap is a representation of the C type cairo_region_overlap_t.
//
type RegionOverlap int

const (
	// RegionOverlap_In is a representation of the C type CAIRO_REGION_OVERLAP_IN.
	RegionOverlap_In RegionOverlap = 0
	// RegionOverlap_Out is a representation of the C type CAIRO_REGION_OVERLAP_OUT.
	RegionOverlap_Out RegionOverlap = 1
	// RegionOverlap_Part is a representation of the C type CAIRO_REGION_OVERLAP_PART.
	RegionOverlap_Part RegionOverlap = 2
)
