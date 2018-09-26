// This is a generated file - DO NOT EDIT

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

type Status C.cairo_status_t

const (
	CAIRO_STATUS_SUCCESS                   Status = 0
	CAIRO_STATUS_NO_MEMORY                 Status = 1
	CAIRO_STATUS_INVALID_RESTORE           Status = 2
	CAIRO_STATUS_INVALID_POP_GROUP         Status = 3
	CAIRO_STATUS_NO_CURRENT_POINT          Status = 4
	CAIRO_STATUS_INVALID_MATRIX            Status = 5
	CAIRO_STATUS_INVALID_STATUS            Status = 6
	CAIRO_STATUS_NULL_POINTER              Status = 7
	CAIRO_STATUS_INVALID_STRING            Status = 8
	CAIRO_STATUS_INVALID_PATH_DATA         Status = 9
	CAIRO_STATUS_READ_ERROR                Status = 10
	CAIRO_STATUS_WRITE_ERROR               Status = 11
	CAIRO_STATUS_SURFACE_FINISHED          Status = 12
	CAIRO_STATUS_SURFACE_TYPE_MISMATCH     Status = 13
	CAIRO_STATUS_PATTERN_TYPE_MISMATCH     Status = 14
	CAIRO_STATUS_INVALID_CONTENT           Status = 15
	CAIRO_STATUS_INVALID_FORMAT            Status = 16
	CAIRO_STATUS_INVALID_VISUAL            Status = 17
	CAIRO_STATUS_FILE_NOT_FOUND            Status = 18
	CAIRO_STATUS_INVALID_DASH              Status = 19
	CAIRO_STATUS_INVALID_DSC_COMMENT       Status = 20
	CAIRO_STATUS_INVALID_INDEX             Status = 21
	CAIRO_STATUS_CLIP_NOT_REPRESENTABLE    Status = 22
	CAIRO_STATUS_TEMP_FILE_ERROR           Status = 23
	CAIRO_STATUS_INVALID_STRIDE            Status = 24
	CAIRO_STATUS_FONT_TYPE_MISMATCH        Status = 25
	CAIRO_STATUS_USER_FONT_IMMUTABLE       Status = 26
	CAIRO_STATUS_USER_FONT_ERROR           Status = 27
	CAIRO_STATUS_NEGATIVE_COUNT            Status = 28
	CAIRO_STATUS_INVALID_CLUSTERS          Status = 29
	CAIRO_STATUS_INVALID_SLANT             Status = 30
	CAIRO_STATUS_INVALID_WEIGHT            Status = 31
	CAIRO_STATUS_INVALID_SIZE              Status = 32
	CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED Status = 33
	CAIRO_STATUS_DEVICE_TYPE_MISMATCH      Status = 34
	CAIRO_STATUS_DEVICE_ERROR              Status = 35
	CAIRO_STATUS_INVALID_MESH_CONSTRUCTION Status = 36
	CAIRO_STATUS_DEVICE_FINISHED           Status = 37
	CAIRO_STATUS_JBIG2_GLOBAL_MISSING      Status = 38
)

type Content C.cairo_content_t

const (
	CAIRO_CONTENT_COLOR       Content = 4096
	CAIRO_CONTENT_ALPHA       Content = 8192
	CAIRO_CONTENT_COLOR_ALPHA Content = 12288
)

type Operator C.cairo_operator_t

const (
	CAIRO_OPERATOR_CLEAR          Operator = 0
	CAIRO_OPERATOR_SOURCE         Operator = 1
	CAIRO_OPERATOR_OVER           Operator = 2
	CAIRO_OPERATOR_IN             Operator = 3
	CAIRO_OPERATOR_OUT            Operator = 4
	CAIRO_OPERATOR_ATOP           Operator = 5
	CAIRO_OPERATOR_DEST           Operator = 6
	CAIRO_OPERATOR_DEST_OVER      Operator = 7
	CAIRO_OPERATOR_DEST_IN        Operator = 8
	CAIRO_OPERATOR_DEST_OUT       Operator = 9
	CAIRO_OPERATOR_DEST_ATOP      Operator = 10
	CAIRO_OPERATOR_XOR            Operator = 11
	CAIRO_OPERATOR_ADD            Operator = 12
	CAIRO_OPERATOR_SATURATE       Operator = 13
	CAIRO_OPERATOR_MULTIPLY       Operator = 14
	CAIRO_OPERATOR_SCREEN         Operator = 15
	CAIRO_OPERATOR_OVERLAY        Operator = 16
	CAIRO_OPERATOR_DARKEN         Operator = 17
	CAIRO_OPERATOR_LIGHTEN        Operator = 18
	CAIRO_OPERATOR_COLOR_DODGE    Operator = 19
	CAIRO_OPERATOR_COLOR_BURN     Operator = 20
	CAIRO_OPERATOR_HARD_LIGHT     Operator = 21
	CAIRO_OPERATOR_SOFT_LIGHT     Operator = 22
	CAIRO_OPERATOR_DIFFERENCE     Operator = 23
	CAIRO_OPERATOR_EXCLUSION      Operator = 24
	CAIRO_OPERATOR_HSL_HUE        Operator = 25
	CAIRO_OPERATOR_HSL_SATURATION Operator = 26
	CAIRO_OPERATOR_HSL_COLOR      Operator = 27
	CAIRO_OPERATOR_HSL_LUMINOSITY Operator = 28
)

type Antialias C.cairo_antialias_t

const (
	CAIRO_ANTIALIAS_DEFAULT  Antialias = 0
	CAIRO_ANTIALIAS_NONE     Antialias = 1
	CAIRO_ANTIALIAS_GRAY     Antialias = 2
	CAIRO_ANTIALIAS_SUBPIXEL Antialias = 3
	CAIRO_ANTIALIAS_FAST     Antialias = 4
	CAIRO_ANTIALIAS_GOOD     Antialias = 5
	CAIRO_ANTIALIAS_BEST     Antialias = 6
)

type FillRule C.cairo_fill_rule_t

const (
	CAIRO_FILL_RULE_WINDING  FillRule = 0
	CAIRO_FILL_RULE_EVEN_ODD FillRule = 1
)

type LineCap C.cairo_line_cap_t

const (
	CAIRO_LINE_CAP_BUTT   LineCap = 0
	CAIRO_LINE_CAP_ROUND  LineCap = 1
	CAIRO_LINE_CAP_SQUARE LineCap = 2
)

type LineJoin C.cairo_line_join_t

const (
	CAIRO_LINE_JOIN_MITER LineJoin = 0
	CAIRO_LINE_JOIN_ROUND LineJoin = 1
	CAIRO_LINE_JOIN_BEVEL LineJoin = 2
)

type TextClusterFlags C.cairo_text_cluster_flags_t

const (
	CAIRO_TEXT_CLUSTER_FLAG_BACKWARD TextClusterFlags = 1
)

type FontSlant C.cairo_font_slant_t

const (
	CAIRO_FONT_SLANT_NORMAL  FontSlant = 0
	CAIRO_FONT_SLANT_ITALIC  FontSlant = 1
	CAIRO_FONT_SLANT_OBLIQUE FontSlant = 2
)

type FontWeight C.cairo_font_weight_t

const (
	CAIRO_FONT_WEIGHT_NORMAL FontWeight = 0
	CAIRO_FONT_WEIGHT_BOLD   FontWeight = 1
)

type SubpixelOrder C.cairo_subpixel_order_t

const (
	CAIRO_SUBPIXEL_ORDER_DEFAULT SubpixelOrder = 0
	CAIRO_SUBPIXEL_ORDER_RGB     SubpixelOrder = 1
	CAIRO_SUBPIXEL_ORDER_BGR     SubpixelOrder = 2
	CAIRO_SUBPIXEL_ORDER_VRGB    SubpixelOrder = 3
	CAIRO_SUBPIXEL_ORDER_VBGR    SubpixelOrder = 4
)

type HintStyle C.cairo_hint_style_t

const (
	CAIRO_HINT_STYLE_DEFAULT HintStyle = 0
	CAIRO_HINT_STYLE_NONE    HintStyle = 1
	CAIRO_HINT_STYLE_SLIGHT  HintStyle = 2
	CAIRO_HINT_STYLE_MEDIUM  HintStyle = 3
	CAIRO_HINT_STYLE_FULL    HintStyle = 4
)

type HintMetrics C.cairo_hint_metrics_t

const (
	CAIRO_HINT_METRICS_DEFAULT HintMetrics = 0
	CAIRO_HINT_METRICS_OFF     HintMetrics = 1
	CAIRO_HINT_METRICS_ON      HintMetrics = 2
)

type FontType C.cairo_font_type_t

const (
	CAIRO_FONT_TYPE_TOY    FontType = 0
	CAIRO_FONT_TYPE_FT     FontType = 1
	CAIRO_FONT_TYPE_WIN32  FontType = 2
	CAIRO_FONT_TYPE_QUARTZ FontType = 3
	CAIRO_FONT_TYPE_USER   FontType = 4
)

type PathDataType C.cairo_path_data_type_t

const (
	CAIRO_PATH_MOVE_TO    PathDataType = 0
	CAIRO_PATH_LINE_TO    PathDataType = 1
	CAIRO_PATH_CURVE_TO   PathDataType = 2
	CAIRO_PATH_CLOSE_PATH PathDataType = 3
)

type DeviceType C.cairo_device_type_t

const (
	CAIRO_DEVICE_TYPE_DRM     DeviceType = 0
	CAIRO_DEVICE_TYPE_GL      DeviceType = 1
	CAIRO_DEVICE_TYPE_SCRIPT  DeviceType = 2
	CAIRO_DEVICE_TYPE_XCB     DeviceType = 3
	CAIRO_DEVICE_TYPE_XLIB    DeviceType = 4
	CAIRO_DEVICE_TYPE_XML     DeviceType = 5
	CAIRO_DEVICE_TYPE_COGL    DeviceType = 6
	CAIRO_DEVICE_TYPE_WIN32   DeviceType = 7
	CAIRO_DEVICE_TYPE_INVALID DeviceType = -1
)

type SurfaceType C.cairo_surface_type_t

const (
	CAIRO_SURFACE_TYPE_IMAGE          SurfaceType = 0
	CAIRO_SURFACE_TYPE_PDF            SurfaceType = 1
	CAIRO_SURFACE_TYPE_PS             SurfaceType = 2
	CAIRO_SURFACE_TYPE_XLIB           SurfaceType = 3
	CAIRO_SURFACE_TYPE_XCB            SurfaceType = 4
	CAIRO_SURFACE_TYPE_GLITZ          SurfaceType = 5
	CAIRO_SURFACE_TYPE_QUARTZ         SurfaceType = 6
	CAIRO_SURFACE_TYPE_WIN32          SurfaceType = 7
	CAIRO_SURFACE_TYPE_BEOS           SurfaceType = 8
	CAIRO_SURFACE_TYPE_DIRECTFB       SurfaceType = 9
	CAIRO_SURFACE_TYPE_SVG            SurfaceType = 10
	CAIRO_SURFACE_TYPE_OS2            SurfaceType = 11
	CAIRO_SURFACE_TYPE_WIN32_PRINTING SurfaceType = 12
	CAIRO_SURFACE_TYPE_QUARTZ_IMAGE   SurfaceType = 13
	CAIRO_SURFACE_TYPE_SCRIPT         SurfaceType = 14
	CAIRO_SURFACE_TYPE_QT             SurfaceType = 15
	CAIRO_SURFACE_TYPE_RECORDING      SurfaceType = 16
	CAIRO_SURFACE_TYPE_VG             SurfaceType = 17
	CAIRO_SURFACE_TYPE_GL             SurfaceType = 18
	CAIRO_SURFACE_TYPE_DRM            SurfaceType = 19
	CAIRO_SURFACE_TYPE_TEE            SurfaceType = 20
	CAIRO_SURFACE_TYPE_XML            SurfaceType = 21
	CAIRO_SURFACE_TYPE_SKIA           SurfaceType = 22
	CAIRO_SURFACE_TYPE_SUBSURFACE     SurfaceType = 23
	CAIRO_SURFACE_TYPE_COGL           SurfaceType = 24
)

type Format C.cairo_format_t

const (
	CAIRO_FORMAT_INVALID   Format = -1
	CAIRO_FORMAT_ARGB32    Format = 0
	CAIRO_FORMAT_RGB24     Format = 1
	CAIRO_FORMAT_A8        Format = 2
	CAIRO_FORMAT_A1        Format = 3
	CAIRO_FORMAT_RGB16_565 Format = 4
	CAIRO_FORMAT_RGB30     Format = 5
)

type PatternType C.cairo_pattern_type_t

const (
	CAIRO_PATTERN_TYPE_SOLID         PatternType = 0
	CAIRO_PATTERN_TYPE_SURFACE       PatternType = 1
	CAIRO_PATTERN_TYPE_LINEAR        PatternType = 2
	CAIRO_PATTERN_TYPE_RADIAL        PatternType = 3
	CAIRO_PATTERN_TYPE_MESH          PatternType = 4
	CAIRO_PATTERN_TYPE_RASTER_SOURCE PatternType = 5
)

type Extend C.cairo_extend_t

const (
	CAIRO_EXTEND_NONE    Extend = 0
	CAIRO_EXTEND_REPEAT  Extend = 1
	CAIRO_EXTEND_REFLECT Extend = 2
	CAIRO_EXTEND_PAD     Extend = 3
)

type Filter C.cairo_filter_t

const (
	CAIRO_FILTER_FAST     Filter = 0
	CAIRO_FILTER_GOOD     Filter = 1
	CAIRO_FILTER_BEST     Filter = 2
	CAIRO_FILTER_NEAREST  Filter = 3
	CAIRO_FILTER_BILINEAR Filter = 4
	CAIRO_FILTER_GAUSSIAN Filter = 5
)

type RegionOverlap C.cairo_region_overlap_t

const (
	CAIRO_REGION_OVERLAP_IN   RegionOverlap = 0
	CAIRO_REGION_OVERLAP_OUT  RegionOverlap = 1
	CAIRO_REGION_OVERLAP_PART RegionOverlap = 2
)
