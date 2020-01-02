// Code generated - DO NOT EDIT.

package cairo

import "unsafe"

// Status is a representation of the C enumeration cairo_status_t.
type Status int

// Status_success is a representation of the C enumeration member CAIRO_STATUS_SUCCESS.
const Status_success = Status(0)

// Status_no_memory is a representation of the C enumeration member CAIRO_STATUS_NO_MEMORY.
const Status_no_memory = Status(1)

// Status_invalid_restore is a representation of the C enumeration member CAIRO_STATUS_INVALID_RESTORE.
const Status_invalid_restore = Status(2)

// Status_invalid_pop_group is a representation of the C enumeration member CAIRO_STATUS_INVALID_POP_GROUP.
const Status_invalid_pop_group = Status(3)

// Status_no_current_point is a representation of the C enumeration member CAIRO_STATUS_NO_CURRENT_POINT.
const Status_no_current_point = Status(4)

// Status_invalid_matrix is a representation of the C enumeration member CAIRO_STATUS_INVALID_MATRIX.
const Status_invalid_matrix = Status(5)

// Status_invalid_status is a representation of the C enumeration member CAIRO_STATUS_INVALID_STATUS.
const Status_invalid_status = Status(6)

// Status_null_pointer is a representation of the C enumeration member CAIRO_STATUS_NULL_POINTER.
const Status_null_pointer = Status(7)

// Status_invalid_string is a representation of the C enumeration member CAIRO_STATUS_INVALID_STRING.
const Status_invalid_string = Status(8)

// Status_invalid_path_data is a representation of the C enumeration member CAIRO_STATUS_INVALID_PATH_DATA.
const Status_invalid_path_data = Status(9)

// Status_read_error is a representation of the C enumeration member CAIRO_STATUS_READ_ERROR.
const Status_read_error = Status(10)

// Status_write_error is a representation of the C enumeration member CAIRO_STATUS_WRITE_ERROR.
const Status_write_error = Status(11)

// Status_surface_finished is a representation of the C enumeration member CAIRO_STATUS_SURFACE_FINISHED.
const Status_surface_finished = Status(12)

// Status_surface_type_mismatch is a representation of the C enumeration member CAIRO_STATUS_SURFACE_TYPE_MISMATCH.
const Status_surface_type_mismatch = Status(13)

// Status_pattern_type_mismatch is a representation of the C enumeration member CAIRO_STATUS_PATTERN_TYPE_MISMATCH.
const Status_pattern_type_mismatch = Status(14)

// Status_invalid_content is a representation of the C enumeration member CAIRO_STATUS_INVALID_CONTENT.
const Status_invalid_content = Status(15)

// Status_invalid_format is a representation of the C enumeration member CAIRO_STATUS_INVALID_FORMAT.
const Status_invalid_format = Status(16)

// Status_invalid_visual is a representation of the C enumeration member CAIRO_STATUS_INVALID_VISUAL.
const Status_invalid_visual = Status(17)

// Status_file_not_found is a representation of the C enumeration member CAIRO_STATUS_FILE_NOT_FOUND.
const Status_file_not_found = Status(18)

// Status_invalid_dash is a representation of the C enumeration member CAIRO_STATUS_INVALID_DASH.
const Status_invalid_dash = Status(19)

// Status_invalid_dsc_comment is a representation of the C enumeration member CAIRO_STATUS_INVALID_DSC_COMMENT.
const Status_invalid_dsc_comment = Status(20)

// Status_invalid_index is a representation of the C enumeration member CAIRO_STATUS_INVALID_INDEX.
const Status_invalid_index = Status(21)

// Status_clip_not_representable is a representation of the C enumeration member CAIRO_STATUS_CLIP_NOT_REPRESENTABLE.
const Status_clip_not_representable = Status(22)

// Status_temp_file_error is a representation of the C enumeration member CAIRO_STATUS_TEMP_FILE_ERROR.
const Status_temp_file_error = Status(23)

// Status_invalid_stride is a representation of the C enumeration member CAIRO_STATUS_INVALID_STRIDE.
const Status_invalid_stride = Status(24)

// Status_font_type_mismatch is a representation of the C enumeration member CAIRO_STATUS_FONT_TYPE_MISMATCH.
const Status_font_type_mismatch = Status(25)

// Status_user_font_immutable is a representation of the C enumeration member CAIRO_STATUS_USER_FONT_IMMUTABLE.
const Status_user_font_immutable = Status(26)

// Status_user_font_error is a representation of the C enumeration member CAIRO_STATUS_USER_FONT_ERROR.
const Status_user_font_error = Status(27)

// Status_negative_count is a representation of the C enumeration member CAIRO_STATUS_NEGATIVE_COUNT.
const Status_negative_count = Status(28)

// Status_invalid_clusters is a representation of the C enumeration member CAIRO_STATUS_INVALID_CLUSTERS.
const Status_invalid_clusters = Status(29)

// Status_invalid_slant is a representation of the C enumeration member CAIRO_STATUS_INVALID_SLANT.
const Status_invalid_slant = Status(30)

// Status_invalid_weight is a representation of the C enumeration member CAIRO_STATUS_INVALID_WEIGHT.
const Status_invalid_weight = Status(31)

// Status_invalid_size is a representation of the C enumeration member CAIRO_STATUS_INVALID_SIZE.
const Status_invalid_size = Status(32)

// Status_user_font_not_implemented is a representation of the C enumeration member CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED.
const Status_user_font_not_implemented = Status(33)

// Status_device_type_mismatch is a representation of the C enumeration member CAIRO_STATUS_DEVICE_TYPE_MISMATCH.
const Status_device_type_mismatch = Status(34)

// Status_device_error is a representation of the C enumeration member CAIRO_STATUS_DEVICE_ERROR.
const Status_device_error = Status(35)

// Status_invalid_mesh_construction is a representation of the C enumeration member CAIRO_STATUS_INVALID_MESH_CONSTRUCTION.
const Status_invalid_mesh_construction = Status(36)

// Status_device_finished is a representation of the C enumeration member CAIRO_STATUS_DEVICE_FINISHED.
const Status_device_finished = Status(37)

// Status_jbig2_global_missing is a representation of the C enumeration member CAIRO_STATUS_JBIG2_GLOBAL_MISSING.
const Status_jbig2_global_missing = Status(38)

// Content is a representation of the C enumeration cairo_content_t.
type Content int

// Content_color is a representation of the C enumeration member CAIRO_CONTENT_COLOR.
const Content_color = Content(4096)

// Content_alpha is a representation of the C enumeration member CAIRO_CONTENT_ALPHA.
const Content_alpha = Content(8192)

// Content_color_alpha is a representation of the C enumeration member CAIRO_CONTENT_COLOR_ALPHA.
const Content_color_alpha = Content(12288)

// Operator is a representation of the C enumeration cairo_operator_t.
type Operator int

// Operator_clear is a representation of the C enumeration member CAIRO_OPERATOR_CLEAR.
const Operator_clear = Operator(0)

// Operator_source is a representation of the C enumeration member CAIRO_OPERATOR_SOURCE.
const Operator_source = Operator(1)

// Operator_over is a representation of the C enumeration member CAIRO_OPERATOR_OVER.
const Operator_over = Operator(2)

// Operator_in is a representation of the C enumeration member CAIRO_OPERATOR_IN.
const Operator_in = Operator(3)

// Operator_out is a representation of the C enumeration member CAIRO_OPERATOR_OUT.
const Operator_out = Operator(4)

// Operator_atop is a representation of the C enumeration member CAIRO_OPERATOR_ATOP.
const Operator_atop = Operator(5)

// Operator_dest is a representation of the C enumeration member CAIRO_OPERATOR_DEST.
const Operator_dest = Operator(6)

// Operator_dest_over is a representation of the C enumeration member CAIRO_OPERATOR_DEST_OVER.
const Operator_dest_over = Operator(7)

// Operator_dest_in is a representation of the C enumeration member CAIRO_OPERATOR_DEST_IN.
const Operator_dest_in = Operator(8)

// Operator_dest_out is a representation of the C enumeration member CAIRO_OPERATOR_DEST_OUT.
const Operator_dest_out = Operator(9)

// Operator_dest_atop is a representation of the C enumeration member CAIRO_OPERATOR_DEST_ATOP.
const Operator_dest_atop = Operator(10)

// Operator_xor is a representation of the C enumeration member CAIRO_OPERATOR_XOR.
const Operator_xor = Operator(11)

// Operator_add is a representation of the C enumeration member CAIRO_OPERATOR_ADD.
const Operator_add = Operator(12)

// Operator_saturate is a representation of the C enumeration member CAIRO_OPERATOR_SATURATE.
const Operator_saturate = Operator(13)

// Operator_multiply is a representation of the C enumeration member CAIRO_OPERATOR_MULTIPLY.
const Operator_multiply = Operator(14)

// Operator_screen is a representation of the C enumeration member CAIRO_OPERATOR_SCREEN.
const Operator_screen = Operator(15)

// Operator_overlay is a representation of the C enumeration member CAIRO_OPERATOR_OVERLAY.
const Operator_overlay = Operator(16)

// Operator_darken is a representation of the C enumeration member CAIRO_OPERATOR_DARKEN.
const Operator_darken = Operator(17)

// Operator_lighten is a representation of the C enumeration member CAIRO_OPERATOR_LIGHTEN.
const Operator_lighten = Operator(18)

// Operator_color_dodge is a representation of the C enumeration member CAIRO_OPERATOR_COLOR_DODGE.
const Operator_color_dodge = Operator(19)

// Operator_color_burn is a representation of the C enumeration member CAIRO_OPERATOR_COLOR_BURN.
const Operator_color_burn = Operator(20)

// Operator_hard_light is a representation of the C enumeration member CAIRO_OPERATOR_HARD_LIGHT.
const Operator_hard_light = Operator(21)

// Operator_soft_light is a representation of the C enumeration member CAIRO_OPERATOR_SOFT_LIGHT.
const Operator_soft_light = Operator(22)

// Operator_difference is a representation of the C enumeration member CAIRO_OPERATOR_DIFFERENCE.
const Operator_difference = Operator(23)

// Operator_exclusion is a representation of the C enumeration member CAIRO_OPERATOR_EXCLUSION.
const Operator_exclusion = Operator(24)

// Operator_hsl_hue is a representation of the C enumeration member CAIRO_OPERATOR_HSL_HUE.
const Operator_hsl_hue = Operator(25)

// Operator_hsl_saturation is a representation of the C enumeration member CAIRO_OPERATOR_HSL_SATURATION.
const Operator_hsl_saturation = Operator(26)

// Operator_hsl_color is a representation of the C enumeration member CAIRO_OPERATOR_HSL_COLOR.
const Operator_hsl_color = Operator(27)

// Operator_hsl_luminosity is a representation of the C enumeration member CAIRO_OPERATOR_HSL_LUMINOSITY.
const Operator_hsl_luminosity = Operator(28)

// Antialias is a representation of the C enumeration cairo_antialias_t.
type Antialias int

// Antialias_default is a representation of the C enumeration member CAIRO_ANTIALIAS_DEFAULT.
const Antialias_default = Antialias(0)

// Antialias_none is a representation of the C enumeration member CAIRO_ANTIALIAS_NONE.
const Antialias_none = Antialias(1)

// Antialias_gray is a representation of the C enumeration member CAIRO_ANTIALIAS_GRAY.
const Antialias_gray = Antialias(2)

// Antialias_subpixel is a representation of the C enumeration member CAIRO_ANTIALIAS_SUBPIXEL.
const Antialias_subpixel = Antialias(3)

// Antialias_fast is a representation of the C enumeration member CAIRO_ANTIALIAS_FAST.
const Antialias_fast = Antialias(4)

// Antialias_good is a representation of the C enumeration member CAIRO_ANTIALIAS_GOOD.
const Antialias_good = Antialias(5)

// Antialias_best is a representation of the C enumeration member CAIRO_ANTIALIAS_BEST.
const Antialias_best = Antialias(6)

// FillRule is a representation of the C enumeration cairo_fill_rule_t.
type FillRule int

// FillRule_winding is a representation of the C enumeration member CAIRO_FILL_RULE_WINDING.
const FillRule_winding = FillRule(0)

// FillRule_even_odd is a representation of the C enumeration member CAIRO_FILL_RULE_EVEN_ODD.
const FillRule_even_odd = FillRule(1)

// LineCap is a representation of the C enumeration cairo_line_cap_t.
type LineCap int

// LineCap_butt is a representation of the C enumeration member CAIRO_LINE_CAP_BUTT.
const LineCap_butt = LineCap(0)

// LineCap_round is a representation of the C enumeration member CAIRO_LINE_CAP_ROUND.
const LineCap_round = LineCap(1)

// LineCap_square is a representation of the C enumeration member CAIRO_LINE_CAP_SQUARE.
const LineCap_square = LineCap(2)

// LineJoin is a representation of the C enumeration cairo_line_join_t.
type LineJoin int

// LineJoin_miter is a representation of the C enumeration member CAIRO_LINE_JOIN_MITER.
const LineJoin_miter = LineJoin(0)

// LineJoin_round is a representation of the C enumeration member CAIRO_LINE_JOIN_ROUND.
const LineJoin_round = LineJoin(1)

// LineJoin_bevel is a representation of the C enumeration member CAIRO_LINE_JOIN_BEVEL.
const LineJoin_bevel = LineJoin(2)

// TextClusterFlags is a representation of the C enumeration cairo_text_cluster_flags_t.
type TextClusterFlags int

// TextClusterFlags_backward is a representation of the C enumeration member CAIRO_TEXT_CLUSTER_FLAG_BACKWARD.
const TextClusterFlags_backward = TextClusterFlags(1)

// FontSlant is a representation of the C enumeration cairo_font_slant_t.
type FontSlant int

// FontSlant_normal is a representation of the C enumeration member CAIRO_FONT_SLANT_NORMAL.
const FontSlant_normal = FontSlant(0)

// FontSlant_italic is a representation of the C enumeration member CAIRO_FONT_SLANT_ITALIC.
const FontSlant_italic = FontSlant(1)

// FontSlant_oblique is a representation of the C enumeration member CAIRO_FONT_SLANT_OBLIQUE.
const FontSlant_oblique = FontSlant(2)

// FontWeight is a representation of the C enumeration cairo_font_weight_t.
type FontWeight int

// FontWeight_normal is a representation of the C enumeration member CAIRO_FONT_WEIGHT_NORMAL.
const FontWeight_normal = FontWeight(0)

// FontWeight_bold is a representation of the C enumeration member CAIRO_FONT_WEIGHT_BOLD.
const FontWeight_bold = FontWeight(1)

// SubpixelOrder is a representation of the C enumeration cairo_subpixel_order_t.
type SubpixelOrder int

// SubpixelOrder_default is a representation of the C enumeration member CAIRO_SUBPIXEL_ORDER_DEFAULT.
const SubpixelOrder_default = SubpixelOrder(0)

// SubpixelOrder_rgb is a representation of the C enumeration member CAIRO_SUBPIXEL_ORDER_RGB.
const SubpixelOrder_rgb = SubpixelOrder(1)

// SubpixelOrder_bgr is a representation of the C enumeration member CAIRO_SUBPIXEL_ORDER_BGR.
const SubpixelOrder_bgr = SubpixelOrder(2)

// SubpixelOrder_vrgb is a representation of the C enumeration member CAIRO_SUBPIXEL_ORDER_VRGB.
const SubpixelOrder_vrgb = SubpixelOrder(3)

// SubpixelOrder_vbgr is a representation of the C enumeration member CAIRO_SUBPIXEL_ORDER_VBGR.
const SubpixelOrder_vbgr = SubpixelOrder(4)

// HintStyle is a representation of the C enumeration cairo_hint_style_t.
type HintStyle int

// HintStyle_default is a representation of the C enumeration member CAIRO_HINT_STYLE_DEFAULT.
const HintStyle_default = HintStyle(0)

// HintStyle_none is a representation of the C enumeration member CAIRO_HINT_STYLE_NONE.
const HintStyle_none = HintStyle(1)

// HintStyle_slight is a representation of the C enumeration member CAIRO_HINT_STYLE_SLIGHT.
const HintStyle_slight = HintStyle(2)

// HintStyle_medium is a representation of the C enumeration member CAIRO_HINT_STYLE_MEDIUM.
const HintStyle_medium = HintStyle(3)

// HintStyle_full is a representation of the C enumeration member CAIRO_HINT_STYLE_FULL.
const HintStyle_full = HintStyle(4)

// HintMetrics is a representation of the C enumeration cairo_hint_metrics_t.
type HintMetrics int

// HintMetrics_default is a representation of the C enumeration member CAIRO_HINT_METRICS_DEFAULT.
const HintMetrics_default = HintMetrics(0)

// HintMetrics_off is a representation of the C enumeration member CAIRO_HINT_METRICS_OFF.
const HintMetrics_off = HintMetrics(1)

// HintMetrics_on is a representation of the C enumeration member CAIRO_HINT_METRICS_ON.
const HintMetrics_on = HintMetrics(2)

// FontType is a representation of the C enumeration cairo_font_type_t.
type FontType int

// FontType_toy is a representation of the C enumeration member CAIRO_FONT_TYPE_TOY.
const FontType_toy = FontType(0)

// FontType_ft is a representation of the C enumeration member CAIRO_FONT_TYPE_FT.
const FontType_ft = FontType(1)

// FontType_win32 is a representation of the C enumeration member CAIRO_FONT_TYPE_WIN32.
const FontType_win32 = FontType(2)

// FontType_quartz is a representation of the C enumeration member CAIRO_FONT_TYPE_QUARTZ.
const FontType_quartz = FontType(3)

// FontType_user is a representation of the C enumeration member CAIRO_FONT_TYPE_USER.
const FontType_user = FontType(4)

// PathDataType is a representation of the C enumeration cairo_path_data_type_t.
type PathDataType int

// PathDataType_move_to is a representation of the C enumeration member CAIRO_PATH_MOVE_TO.
const PathDataType_move_to = PathDataType(0)

// PathDataType_line_to is a representation of the C enumeration member CAIRO_PATH_LINE_TO.
const PathDataType_line_to = PathDataType(1)

// PathDataType_curve_to is a representation of the C enumeration member CAIRO_PATH_CURVE_TO.
const PathDataType_curve_to = PathDataType(2)

// PathDataType_close_path is a representation of the C enumeration member CAIRO_PATH_CLOSE_PATH.
const PathDataType_close_path = PathDataType(3)

// DeviceType is a representation of the C enumeration cairo_device_type_t.
type DeviceType int

// DeviceType_drm is a representation of the C enumeration member CAIRO_DEVICE_TYPE_DRM.
const DeviceType_drm = DeviceType(0)

// DeviceType_gl is a representation of the C enumeration member CAIRO_DEVICE_TYPE_GL.
const DeviceType_gl = DeviceType(1)

// DeviceType_script is a representation of the C enumeration member CAIRO_DEVICE_TYPE_SCRIPT.
const DeviceType_script = DeviceType(2)

// DeviceType_xcb is a representation of the C enumeration member CAIRO_DEVICE_TYPE_XCB.
const DeviceType_xcb = DeviceType(3)

// DeviceType_xlib is a representation of the C enumeration member CAIRO_DEVICE_TYPE_XLIB.
const DeviceType_xlib = DeviceType(4)

// DeviceType_xml is a representation of the C enumeration member CAIRO_DEVICE_TYPE_XML.
const DeviceType_xml = DeviceType(5)

// DeviceType_cogl is a representation of the C enumeration member CAIRO_DEVICE_TYPE_COGL.
const DeviceType_cogl = DeviceType(6)

// DeviceType_win32 is a representation of the C enumeration member CAIRO_DEVICE_TYPE_WIN32.
const DeviceType_win32 = DeviceType(7)

// DeviceType_invalid is a representation of the C enumeration member CAIRO_DEVICE_TYPE_INVALID.
const DeviceType_invalid = DeviceType(-1)

// SurfaceType is a representation of the C enumeration cairo_surface_type_t.
type SurfaceType int

// SurfaceType_image is a representation of the C enumeration member CAIRO_SURFACE_TYPE_IMAGE.
const SurfaceType_image = SurfaceType(0)

// SurfaceType_pdf is a representation of the C enumeration member CAIRO_SURFACE_TYPE_PDF.
const SurfaceType_pdf = SurfaceType(1)

// SurfaceType_ps is a representation of the C enumeration member CAIRO_SURFACE_TYPE_PS.
const SurfaceType_ps = SurfaceType(2)

// SurfaceType_xlib is a representation of the C enumeration member CAIRO_SURFACE_TYPE_XLIB.
const SurfaceType_xlib = SurfaceType(3)

// SurfaceType_xcb is a representation of the C enumeration member CAIRO_SURFACE_TYPE_XCB.
const SurfaceType_xcb = SurfaceType(4)

// SurfaceType_glitz is a representation of the C enumeration member CAIRO_SURFACE_TYPE_GLITZ.
const SurfaceType_glitz = SurfaceType(5)

// SurfaceType_quartz is a representation of the C enumeration member CAIRO_SURFACE_TYPE_QUARTZ.
const SurfaceType_quartz = SurfaceType(6)

// SurfaceType_win32 is a representation of the C enumeration member CAIRO_SURFACE_TYPE_WIN32.
const SurfaceType_win32 = SurfaceType(7)

// SurfaceType_beos is a representation of the C enumeration member CAIRO_SURFACE_TYPE_BEOS.
const SurfaceType_beos = SurfaceType(8)

// SurfaceType_directfb is a representation of the C enumeration member CAIRO_SURFACE_TYPE_DIRECTFB.
const SurfaceType_directfb = SurfaceType(9)

// SurfaceType_svg is a representation of the C enumeration member CAIRO_SURFACE_TYPE_SVG.
const SurfaceType_svg = SurfaceType(10)

// SurfaceType_os2 is a representation of the C enumeration member CAIRO_SURFACE_TYPE_OS2.
const SurfaceType_os2 = SurfaceType(11)

// SurfaceType_win32_printing is a representation of the C enumeration member CAIRO_SURFACE_TYPE_WIN32_PRINTING.
const SurfaceType_win32_printing = SurfaceType(12)

// SurfaceType_quartz_image is a representation of the C enumeration member CAIRO_SURFACE_TYPE_QUARTZ_IMAGE.
const SurfaceType_quartz_image = SurfaceType(13)

// SurfaceType_script is a representation of the C enumeration member CAIRO_SURFACE_TYPE_SCRIPT.
const SurfaceType_script = SurfaceType(14)

// SurfaceType_qt is a representation of the C enumeration member CAIRO_SURFACE_TYPE_QT.
const SurfaceType_qt = SurfaceType(15)

// SurfaceType_recording is a representation of the C enumeration member CAIRO_SURFACE_TYPE_RECORDING.
const SurfaceType_recording = SurfaceType(16)

// SurfaceType_vg is a representation of the C enumeration member CAIRO_SURFACE_TYPE_VG.
const SurfaceType_vg = SurfaceType(17)

// SurfaceType_gl is a representation of the C enumeration member CAIRO_SURFACE_TYPE_GL.
const SurfaceType_gl = SurfaceType(18)

// SurfaceType_drm is a representation of the C enumeration member CAIRO_SURFACE_TYPE_DRM.
const SurfaceType_drm = SurfaceType(19)

// SurfaceType_tee is a representation of the C enumeration member CAIRO_SURFACE_TYPE_TEE.
const SurfaceType_tee = SurfaceType(20)

// SurfaceType_xml is a representation of the C enumeration member CAIRO_SURFACE_TYPE_XML.
const SurfaceType_xml = SurfaceType(21)

// SurfaceType_skia is a representation of the C enumeration member CAIRO_SURFACE_TYPE_SKIA.
const SurfaceType_skia = SurfaceType(22)

// SurfaceType_subsurface is a representation of the C enumeration member CAIRO_SURFACE_TYPE_SUBSURFACE.
const SurfaceType_subsurface = SurfaceType(23)

// SurfaceType_cogl is a representation of the C enumeration member CAIRO_SURFACE_TYPE_COGL.
const SurfaceType_cogl = SurfaceType(24)

// Format is a representation of the C enumeration cairo_format_t.
type Format int

// Format_invalid is a representation of the C enumeration member CAIRO_FORMAT_INVALID.
const Format_invalid = Format(-1)

// Format_argb32 is a representation of the C enumeration member CAIRO_FORMAT_ARGB32.
const Format_argb32 = Format(0)

// Format_rgb24 is a representation of the C enumeration member CAIRO_FORMAT_RGB24.
const Format_rgb24 = Format(1)

// Format_a8 is a representation of the C enumeration member CAIRO_FORMAT_A8.
const Format_a8 = Format(2)

// Format_a1 is a representation of the C enumeration member CAIRO_FORMAT_A1.
const Format_a1 = Format(3)

// Format_rgb16_565 is a representation of the C enumeration member CAIRO_FORMAT_RGB16_565.
const Format_rgb16_565 = Format(4)

// Format_rgb30 is a representation of the C enumeration member CAIRO_FORMAT_RGB30.
const Format_rgb30 = Format(5)

// PatternType is a representation of the C enumeration cairo_pattern_type_t.
type PatternType int

// PatternType_solid is a representation of the C enumeration member CAIRO_PATTERN_TYPE_SOLID.
const PatternType_solid = PatternType(0)

// PatternType_surface is a representation of the C enumeration member CAIRO_PATTERN_TYPE_SURFACE.
const PatternType_surface = PatternType(1)

// PatternType_linear is a representation of the C enumeration member CAIRO_PATTERN_TYPE_LINEAR.
const PatternType_linear = PatternType(2)

// PatternType_radial is a representation of the C enumeration member CAIRO_PATTERN_TYPE_RADIAL.
const PatternType_radial = PatternType(3)

// PatternType_mesh is a representation of the C enumeration member CAIRO_PATTERN_TYPE_MESH.
const PatternType_mesh = PatternType(4)

// PatternType_raster_source is a representation of the C enumeration member CAIRO_PATTERN_TYPE_RASTER_SOURCE.
const PatternType_raster_source = PatternType(5)

// Extend is a representation of the C enumeration cairo_extend_t.
type Extend int

// Extend_none is a representation of the C enumeration member CAIRO_EXTEND_NONE.
const Extend_none = Extend(0)

// Extend_repeat is a representation of the C enumeration member CAIRO_EXTEND_REPEAT.
const Extend_repeat = Extend(1)

// Extend_reflect is a representation of the C enumeration member CAIRO_EXTEND_REFLECT.
const Extend_reflect = Extend(2)

// Extend_pad is a representation of the C enumeration member CAIRO_EXTEND_PAD.
const Extend_pad = Extend(3)

// Filter is a representation of the C enumeration cairo_filter_t.
type Filter int

// Filter_fast is a representation of the C enumeration member CAIRO_FILTER_FAST.
const Filter_fast = Filter(0)

// Filter_good is a representation of the C enumeration member CAIRO_FILTER_GOOD.
const Filter_good = Filter(1)

// Filter_best is a representation of the C enumeration member CAIRO_FILTER_BEST.
const Filter_best = Filter(2)

// Filter_nearest is a representation of the C enumeration member CAIRO_FILTER_NEAREST.
const Filter_nearest = Filter(3)

// Filter_bilinear is a representation of the C enumeration member CAIRO_FILTER_BILINEAR.
const Filter_bilinear = Filter(4)

// Filter_gaussian is a representation of the C enumeration member CAIRO_FILTER_GAUSSIAN.
const Filter_gaussian = Filter(5)

// RegionOverlap is a representation of the C enumeration cairo_region_overlap_t.
type RegionOverlap int

// RegionOverlap_in is a representation of the C enumeration member CAIRO_REGION_OVERLAP_IN.
const RegionOverlap_in = RegionOverlap(0)

// RegionOverlap_out is a representation of the C enumeration member CAIRO_REGION_OVERLAP_OUT.
const RegionOverlap_out = RegionOverlap(1)

// RegionOverlap_part is a representation of the C enumeration member CAIRO_REGION_OVERLAP_PART.
const RegionOverlap_part = RegionOverlap(2)

// UNSUPPORTED : cairo_image_surface_create : blacklisted

// Context is a representation of the C record cairo_t.
type Context struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_t that represents the Context.
func (recv *Context) ToC() unsafe.Pointer {
	return recv.native
}

// Device is a representation of the C record cairo_device_t.
type Device struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_device_t that represents the Device.
func (recv *Device) ToC() unsafe.Pointer {
	return recv.native
}

// Surface is a representation of the C record cairo_surface_t.
type Surface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_surface_t that represents the Surface.
func (recv *Surface) ToC() unsafe.Pointer {
	return recv.native
}

// Matrix is a representation of the C record cairo_matrix_t.
type Matrix struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_matrix_t that represents the Matrix.
func (recv *Matrix) ToC() unsafe.Pointer {
	return recv.native
}

// Pattern is a representation of the C record cairo_pattern_t.
type Pattern struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_pattern_t that represents the Pattern.
func (recv *Pattern) ToC() unsafe.Pointer {
	return recv.native
}

// Region is a representation of the C record cairo_region_t.
type Region struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_region_t that represents the Region.
func (recv *Region) ToC() unsafe.Pointer {
	return recv.native
}

// FontOptions is a representation of the C record cairo_font_options_t.
type FontOptions struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_font_options_t that represents the FontOptions.
func (recv *FontOptions) ToC() unsafe.Pointer {
	return recv.native
}

// FontFace is a representation of the C record cairo_font_face_t.
type FontFace struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_font_face_t that represents the FontFace.
func (recv *FontFace) ToC() unsafe.Pointer {
	return recv.native
}

// ScaledFont is a representation of the C record cairo_scaled_font_t.
type ScaledFont struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_scaled_font_t that represents the ScaledFont.
func (recv *ScaledFont) ToC() unsafe.Pointer {
	return recv.native
}

// Path is a representation of the C record cairo_path_t.
type Path struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_path_t that represents the Path.
func (recv *Path) ToC() unsafe.Pointer {
	return recv.native
}

// Rectangle is a representation of the C record cairo_rectangle_t.
type Rectangle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_rectangle_t that represents the Rectangle.
func (recv *Rectangle) ToC() unsafe.Pointer {
	return recv.native
}

// RectangleInt is a representation of the C record cairo_rectangle_int_t.
type RectangleInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C cairo_rectangle_int_t that represents the RectangleInt.
func (recv *RectangleInt) ToC() unsafe.Pointer {
	return recv.native
}
