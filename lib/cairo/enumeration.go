// Code generated - DO NOT EDIT.

package cairo

// Status is a representation of the C type cairo_status_t.
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

// Fillrule is a representation of the C type cairo_fill_rule_t.
type Fillrule int

const (
	// Fillrule_Winding is a representation of the C type CAIRO_FILL_RULE_WINDING.
	Fillrule_Winding Fillrule = 0
	// Fillrule_EvenOdd is a representation of the C type CAIRO_FILL_RULE_EVEN_ODD.
	Fillrule_EvenOdd Fillrule = 1
)

// Linecap is a representation of the C type cairo_line_cap_t.
type Linecap int

const (
	// Linecap_Butt is a representation of the C type CAIRO_LINE_CAP_BUTT.
	Linecap_Butt Linecap = 0
	// Linecap_Round is a representation of the C type CAIRO_LINE_CAP_ROUND.
	Linecap_Round Linecap = 1
	// Linecap_Square is a representation of the C type CAIRO_LINE_CAP_SQUARE.
	Linecap_Square Linecap = 2
)

// Linejoin is a representation of the C type cairo_line_join_t.
type Linejoin int

const (
	// Linejoin_Miter is a representation of the C type CAIRO_LINE_JOIN_MITER.
	Linejoin_Miter Linejoin = 0
	// Linejoin_Round is a representation of the C type CAIRO_LINE_JOIN_ROUND.
	Linejoin_Round Linejoin = 1
	// Linejoin_Bevel is a representation of the C type CAIRO_LINE_JOIN_BEVEL.
	Linejoin_Bevel Linejoin = 2
)

// Textclusterflags is a representation of the C type cairo_text_cluster_flags_t.
type Textclusterflags int

const (
	// Textclusterflags_Backward is a representation of the C type CAIRO_TEXT_CLUSTER_FLAG_BACKWARD.
	Textclusterflags_Backward Textclusterflags = 1
)

// Fontslant is a representation of the C type cairo_font_slant_t.
type Fontslant int

const (
	// Fontslant_Normal is a representation of the C type CAIRO_FONT_SLANT_NORMAL.
	Fontslant_Normal Fontslant = 0
	// Fontslant_Italic is a representation of the C type CAIRO_FONT_SLANT_ITALIC.
	Fontslant_Italic Fontslant = 1
	// Fontslant_Oblique is a representation of the C type CAIRO_FONT_SLANT_OBLIQUE.
	Fontslant_Oblique Fontslant = 2
)

// Fontweight is a representation of the C type cairo_font_weight_t.
type Fontweight int

const (
	// Fontweight_Normal is a representation of the C type CAIRO_FONT_WEIGHT_NORMAL.
	Fontweight_Normal Fontweight = 0
	// Fontweight_Bold is a representation of the C type CAIRO_FONT_WEIGHT_BOLD.
	Fontweight_Bold Fontweight = 1
)

// Subpixelorder is a representation of the C type cairo_subpixel_order_t.
type Subpixelorder int

const (
	// Subpixelorder_Default is a representation of the C type CAIRO_SUBPIXEL_ORDER_DEFAULT.
	Subpixelorder_Default Subpixelorder = 0
	// Subpixelorder_Rgb is a representation of the C type CAIRO_SUBPIXEL_ORDER_RGB.
	Subpixelorder_Rgb Subpixelorder = 1
	// Subpixelorder_Bgr is a representation of the C type CAIRO_SUBPIXEL_ORDER_BGR.
	Subpixelorder_Bgr Subpixelorder = 2
	// Subpixelorder_Vrgb is a representation of the C type CAIRO_SUBPIXEL_ORDER_VRGB.
	Subpixelorder_Vrgb Subpixelorder = 3
	// Subpixelorder_Vbgr is a representation of the C type CAIRO_SUBPIXEL_ORDER_VBGR.
	Subpixelorder_Vbgr Subpixelorder = 4
)

// Hintstyle is a representation of the C type cairo_hint_style_t.
type Hintstyle int

const (
	// Hintstyle_Default is a representation of the C type CAIRO_HINT_STYLE_DEFAULT.
	Hintstyle_Default Hintstyle = 0
	// Hintstyle_None is a representation of the C type CAIRO_HINT_STYLE_NONE.
	Hintstyle_None Hintstyle = 1
	// Hintstyle_Slight is a representation of the C type CAIRO_HINT_STYLE_SLIGHT.
	Hintstyle_Slight Hintstyle = 2
	// Hintstyle_Medium is a representation of the C type CAIRO_HINT_STYLE_MEDIUM.
	Hintstyle_Medium Hintstyle = 3
	// Hintstyle_Full is a representation of the C type CAIRO_HINT_STYLE_FULL.
	Hintstyle_Full Hintstyle = 4
)

// Hintmetrics is a representation of the C type cairo_hint_metrics_t.
type Hintmetrics int

const (
	// Hintmetrics_Default is a representation of the C type CAIRO_HINT_METRICS_DEFAULT.
	Hintmetrics_Default Hintmetrics = 0
	// Hintmetrics_Off is a representation of the C type CAIRO_HINT_METRICS_OFF.
	Hintmetrics_Off Hintmetrics = 1
	// Hintmetrics_On is a representation of the C type CAIRO_HINT_METRICS_ON.
	Hintmetrics_On Hintmetrics = 2
)

// Fonttype is a representation of the C type cairo_font_type_t.
type Fonttype int

const (
	// Fonttype_Toy is a representation of the C type CAIRO_FONT_TYPE_TOY.
	Fonttype_Toy Fonttype = 0
	// Fonttype_Ft is a representation of the C type CAIRO_FONT_TYPE_FT.
	Fonttype_Ft Fonttype = 1
	// Fonttype_Win32 is a representation of the C type CAIRO_FONT_TYPE_WIN32.
	Fonttype_Win32 Fonttype = 2
	// Fonttype_Quartz is a representation of the C type CAIRO_FONT_TYPE_QUARTZ.
	Fonttype_Quartz Fonttype = 3
	// Fonttype_User is a representation of the C type CAIRO_FONT_TYPE_USER.
	Fonttype_User Fonttype = 4
)

// Pathdatatype is a representation of the C type cairo_path_data_type_t.
type Pathdatatype int

const (
	// Pathdatatype_MoveTo is a representation of the C type CAIRO_PATH_MOVE_TO.
	Pathdatatype_MoveTo Pathdatatype = 0
	// Pathdatatype_LineTo is a representation of the C type CAIRO_PATH_LINE_TO.
	Pathdatatype_LineTo Pathdatatype = 1
	// Pathdatatype_CurveTo is a representation of the C type CAIRO_PATH_CURVE_TO.
	Pathdatatype_CurveTo Pathdatatype = 2
	// Pathdatatype_ClosePath is a representation of the C type CAIRO_PATH_CLOSE_PATH.
	Pathdatatype_ClosePath Pathdatatype = 3
)

// Devicetype is a representation of the C type cairo_device_type_t.
type Devicetype int

const (
	// Devicetype_Drm is a representation of the C type CAIRO_DEVICE_TYPE_DRM.
	Devicetype_Drm Devicetype = 0
	// Devicetype_Gl is a representation of the C type CAIRO_DEVICE_TYPE_GL.
	Devicetype_Gl Devicetype = 1
	// Devicetype_Script is a representation of the C type CAIRO_DEVICE_TYPE_SCRIPT.
	Devicetype_Script Devicetype = 2
	// Devicetype_Xcb is a representation of the C type CAIRO_DEVICE_TYPE_XCB.
	Devicetype_Xcb Devicetype = 3
	// Devicetype_Xlib is a representation of the C type CAIRO_DEVICE_TYPE_XLIB.
	Devicetype_Xlib Devicetype = 4
	// Devicetype_Xml is a representation of the C type CAIRO_DEVICE_TYPE_XML.
	Devicetype_Xml Devicetype = 5
	// Devicetype_Cogl is a representation of the C type CAIRO_DEVICE_TYPE_COGL.
	Devicetype_Cogl Devicetype = 6
	// Devicetype_Win32 is a representation of the C type CAIRO_DEVICE_TYPE_WIN32.
	Devicetype_Win32 Devicetype = 7
	// Devicetype_Invalid is a representation of the C type CAIRO_DEVICE_TYPE_INVALID.
	Devicetype_Invalid Devicetype = -1
)

// Surfacetype is a representation of the C type cairo_surface_type_t.
type Surfacetype int

const (
	// Surfacetype_Image is a representation of the C type CAIRO_SURFACE_TYPE_IMAGE.
	Surfacetype_Image Surfacetype = 0
	// Surfacetype_Pdf is a representation of the C type CAIRO_SURFACE_TYPE_PDF.
	Surfacetype_Pdf Surfacetype = 1
	// Surfacetype_Ps is a representation of the C type CAIRO_SURFACE_TYPE_PS.
	Surfacetype_Ps Surfacetype = 2
	// Surfacetype_Xlib is a representation of the C type CAIRO_SURFACE_TYPE_XLIB.
	Surfacetype_Xlib Surfacetype = 3
	// Surfacetype_Xcb is a representation of the C type CAIRO_SURFACE_TYPE_XCB.
	Surfacetype_Xcb Surfacetype = 4
	// Surfacetype_Glitz is a representation of the C type CAIRO_SURFACE_TYPE_GLITZ.
	Surfacetype_Glitz Surfacetype = 5
	// Surfacetype_Quartz is a representation of the C type CAIRO_SURFACE_TYPE_QUARTZ.
	Surfacetype_Quartz Surfacetype = 6
	// Surfacetype_Win32 is a representation of the C type CAIRO_SURFACE_TYPE_WIN32.
	Surfacetype_Win32 Surfacetype = 7
	// Surfacetype_Beos is a representation of the C type CAIRO_SURFACE_TYPE_BEOS.
	Surfacetype_Beos Surfacetype = 8
	// Surfacetype_Directfb is a representation of the C type CAIRO_SURFACE_TYPE_DIRECTFB.
	Surfacetype_Directfb Surfacetype = 9
	// Surfacetype_Svg is a representation of the C type CAIRO_SURFACE_TYPE_SVG.
	Surfacetype_Svg Surfacetype = 10
	// Surfacetype_Os2 is a representation of the C type CAIRO_SURFACE_TYPE_OS2.
	Surfacetype_Os2 Surfacetype = 11
	// Surfacetype_Win32Printing is a representation of the C type CAIRO_SURFACE_TYPE_WIN32_PRINTING.
	Surfacetype_Win32Printing Surfacetype = 12
	// Surfacetype_QuartzImage is a representation of the C type CAIRO_SURFACE_TYPE_QUARTZ_IMAGE.
	Surfacetype_QuartzImage Surfacetype = 13
	// Surfacetype_Script is a representation of the C type CAIRO_SURFACE_TYPE_SCRIPT.
	Surfacetype_Script Surfacetype = 14
	// Surfacetype_Qt is a representation of the C type CAIRO_SURFACE_TYPE_QT.
	Surfacetype_Qt Surfacetype = 15
	// Surfacetype_Recording is a representation of the C type CAIRO_SURFACE_TYPE_RECORDING.
	Surfacetype_Recording Surfacetype = 16
	// Surfacetype_Vg is a representation of the C type CAIRO_SURFACE_TYPE_VG.
	Surfacetype_Vg Surfacetype = 17
	// Surfacetype_Gl is a representation of the C type CAIRO_SURFACE_TYPE_GL.
	Surfacetype_Gl Surfacetype = 18
	// Surfacetype_Drm is a representation of the C type CAIRO_SURFACE_TYPE_DRM.
	Surfacetype_Drm Surfacetype = 19
	// Surfacetype_Tee is a representation of the C type CAIRO_SURFACE_TYPE_TEE.
	Surfacetype_Tee Surfacetype = 20
	// Surfacetype_Xml is a representation of the C type CAIRO_SURFACE_TYPE_XML.
	Surfacetype_Xml Surfacetype = 21
	// Surfacetype_Skia is a representation of the C type CAIRO_SURFACE_TYPE_SKIA.
	Surfacetype_Skia Surfacetype = 22
	// Surfacetype_Subsurface is a representation of the C type CAIRO_SURFACE_TYPE_SUBSURFACE.
	Surfacetype_Subsurface Surfacetype = 23
	// Surfacetype_Cogl is a representation of the C type CAIRO_SURFACE_TYPE_COGL.
	Surfacetype_Cogl Surfacetype = 24
)

// Format is a representation of the C type cairo_format_t.
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

// Patterntype is a representation of the C type cairo_pattern_type_t.
type Patterntype int

const (
	// Patterntype_Solid is a representation of the C type CAIRO_PATTERN_TYPE_SOLID.
	Patterntype_Solid Patterntype = 0
	// Patterntype_Surface is a representation of the C type CAIRO_PATTERN_TYPE_SURFACE.
	Patterntype_Surface Patterntype = 1
	// Patterntype_Linear is a representation of the C type CAIRO_PATTERN_TYPE_LINEAR.
	Patterntype_Linear Patterntype = 2
	// Patterntype_Radial is a representation of the C type CAIRO_PATTERN_TYPE_RADIAL.
	Patterntype_Radial Patterntype = 3
	// Patterntype_Mesh is a representation of the C type CAIRO_PATTERN_TYPE_MESH.
	Patterntype_Mesh Patterntype = 4
	// Patterntype_RasterSource is a representation of the C type CAIRO_PATTERN_TYPE_RASTER_SOURCE.
	Patterntype_RasterSource Patterntype = 5
)

// Extend is a representation of the C type cairo_extend_t.
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

// Regionoverlap is a representation of the C type cairo_region_overlap_t.
type Regionoverlap int

const (
	// Regionoverlap_In is a representation of the C type CAIRO_REGION_OVERLAP_IN.
	Regionoverlap_In Regionoverlap = 0
	// Regionoverlap_Out is a representation of the C type CAIRO_REGION_OVERLAP_OUT.
	Regionoverlap_Out Regionoverlap = 1
	// Regionoverlap_Part is a representation of the C type CAIRO_REGION_OVERLAP_PART.
	Regionoverlap_Part Regionoverlap = 2
)
