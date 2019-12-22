// Code generated - DO NOT EDIT.

package cairo

// #include <cairo-gobject.h>
import "C"

// aliases

// bitfields

// enumerations
type Status C.cairo_status_t
type Content C.cairo_content_t
type Operator C.cairo_operator_t
type Antialias C.cairo_antialias_t
type FillRule C.cairo_fill_rule_t
type LineCap C.cairo_line_cap_t
type LineJoin C.cairo_line_join_t
type TextClusterFlags C.cairo_text_cluster_flags_t
type FontSlant C.cairo_font_slant_t
type FontWeight C.cairo_font_weight_t
type SubpixelOrder C.cairo_subpixel_order_t
type HintStyle C.cairo_hint_style_t
type HintMetrics C.cairo_hint_metrics_t
type FontType C.cairo_font_type_t
type PathDataType C.cairo_path_data_type_t
type DeviceType C.cairo_device_type_t
type SurfaceType C.cairo_surface_type_t
type Format C.cairo_format_t
type PatternType C.cairo_pattern_type_t
type Extend C.cairo_extend_t
type Filter C.cairo_filter_t
type RegionOverlap C.cairo_region_overlap_t

// unions

// records
type Context C.cairo_t
type Device C.cairo_device_t
type Surface C.cairo_surface_t
type Matrix C.cairo_matrix_t
type Pattern C.cairo_pattern_t
type Region C.cairo_region_t
type FontOptions C.cairo_font_options_t
type FontFace C.cairo_font_face_t
type ScaledFont C.cairo_scaled_font_t
type Path C.cairo_path_t
type Rectangle C.cairo_rectangle_t
type RectangleInt C.cairo_rectangle_int_t

// classes

// interfaces

func Fn_image_surface_create() {}
