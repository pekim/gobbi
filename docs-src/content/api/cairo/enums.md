+++
title = "enums"
+++
<p class="api-heading">Status</p>
<table>
<tr>
<td class="name">CAIRO_STATUS_SUCCESS</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_NO_MEMORY</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_RESTORE</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_POP_GROUP</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_NO_CURRENT_POINT</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_MATRIX</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_STATUS</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_NULL_POINTER</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_STRING</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_PATH_DATA</td>
<td class="value">9</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_READ_ERROR</td>
<td class="value">10</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_WRITE_ERROR</td>
<td class="value">11</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_SURFACE_FINISHED</td>
<td class="value">12</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_SURFACE_TYPE_MISMATCH</td>
<td class="value">13</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_PATTERN_TYPE_MISMATCH</td>
<td class="value">14</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_CONTENT</td>
<td class="value">15</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_FORMAT</td>
<td class="value">16</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_VISUAL</td>
<td class="value">17</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_FILE_NOT_FOUND</td>
<td class="value">18</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_DASH</td>
<td class="value">19</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_DSC_COMMENT</td>
<td class="value">20</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_INDEX</td>
<td class="value">21</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_CLIP_NOT_REPRESENTABLE</td>
<td class="value">22</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_TEMP_FILE_ERROR</td>
<td class="value">23</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_STRIDE</td>
<td class="value">24</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_FONT_TYPE_MISMATCH</td>
<td class="value">25</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_USER_FONT_IMMUTABLE</td>
<td class="value">26</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_USER_FONT_ERROR</td>
<td class="value">27</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_NEGATIVE_COUNT</td>
<td class="value">28</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_CLUSTERS</td>
<td class="value">29</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_SLANT</td>
<td class="value">30</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_WEIGHT</td>
<td class="value">31</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_SIZE</td>
<td class="value">32</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED</td>
<td class="value">33</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_DEVICE_TYPE_MISMATCH</td>
<td class="value">34</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_DEVICE_ERROR</td>
<td class="value">35</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_INVALID_MESH_CONSTRUCTION</td>
<td class="value">36</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_DEVICE_FINISHED</td>
<td class="value">37</td>
</tr>
<tr>
<td class="name">CAIRO_STATUS_JBIG2_GLOBAL_MISSING</td>
<td class="value">38</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_status_t</p>
</div>
<p class="api-heading">Content</p>
<table>
<tr>
<td class="name">CAIRO_CONTENT_COLOR</td>
<td class="value">4096</td>
</tr>
<tr>
<td class="name">CAIRO_CONTENT_ALPHA</td>
<td class="value">8192</td>
</tr>
<tr>
<td class="name">CAIRO_CONTENT_COLOR_ALPHA</td>
<td class="value">12288</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_content_t</p>
</div>
<p class="api-heading">Operator</p>
<table>
<tr>
<td class="name">CAIRO_OPERATOR_CLEAR</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_SOURCE</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_OVER</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_IN</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_OUT</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_ATOP</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DEST</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DEST_OVER</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DEST_IN</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DEST_OUT</td>
<td class="value">9</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DEST_ATOP</td>
<td class="value">10</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_XOR</td>
<td class="value">11</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_ADD</td>
<td class="value">12</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_SATURATE</td>
<td class="value">13</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_MULTIPLY</td>
<td class="value">14</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_SCREEN</td>
<td class="value">15</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_OVERLAY</td>
<td class="value">16</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DARKEN</td>
<td class="value">17</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_LIGHTEN</td>
<td class="value">18</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_COLOR_DODGE</td>
<td class="value">19</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_COLOR_BURN</td>
<td class="value">20</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_HARD_LIGHT</td>
<td class="value">21</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_SOFT_LIGHT</td>
<td class="value">22</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_DIFFERENCE</td>
<td class="value">23</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_EXCLUSION</td>
<td class="value">24</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_HSL_HUE</td>
<td class="value">25</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_HSL_SATURATION</td>
<td class="value">26</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_HSL_COLOR</td>
<td class="value">27</td>
</tr>
<tr>
<td class="name">CAIRO_OPERATOR_HSL_LUMINOSITY</td>
<td class="value">28</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_operator_t</p>
</div>
<p class="api-heading">Antialias</p>
<table>
<tr>
<td class="name">CAIRO_ANTIALIAS_DEFAULT</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_NONE</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_GRAY</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_SUBPIXEL</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_FAST</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_GOOD</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">CAIRO_ANTIALIAS_BEST</td>
<td class="value">6</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_antialias_t</p>
</div>
<p class="api-heading">FillRule</p>
<table>
<tr>
<td class="name">CAIRO_FILL_RULE_WINDING</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FILL_RULE_EVEN_ODD</td>
<td class="value">1</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_fill_rule_t</p>
</div>
<p class="api-heading">LineCap</p>
<table>
<tr>
<td class="name">CAIRO_LINE_CAP_BUTT</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_LINE_CAP_ROUND</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_LINE_CAP_SQUARE</td>
<td class="value">2</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_line_cap_t</p>
</div>
<p class="api-heading">LineJoin</p>
<table>
<tr>
<td class="name">CAIRO_LINE_JOIN_MITER</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_LINE_JOIN_ROUND</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_LINE_JOIN_BEVEL</td>
<td class="value">2</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_line_join_t</p>
</div>
<p class="api-heading">TextClusterFlags</p>
<table>
<tr>
<td class="name">CAIRO_TEXT_CLUSTER_FLAG_BACKWARD</td>
<td class="value">1</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_text_cluster_flags_t</p>
</div>
<p class="api-heading">FontSlant</p>
<table>
<tr>
<td class="name">CAIRO_FONT_SLANT_NORMAL</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_SLANT_ITALIC</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_SLANT_OBLIQUE</td>
<td class="value">2</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_font_slant_t</p>
</div>
<p class="api-heading">FontWeight</p>
<table>
<tr>
<td class="name">CAIRO_FONT_WEIGHT_NORMAL</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_WEIGHT_BOLD</td>
<td class="value">1</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_font_weight_t</p>
</div>
<p class="api-heading">SubpixelOrder</p>
<table>
<tr>
<td class="name">CAIRO_SUBPIXEL_ORDER_DEFAULT</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_SUBPIXEL_ORDER_RGB</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_SUBPIXEL_ORDER_BGR</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_SUBPIXEL_ORDER_VRGB</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_SUBPIXEL_ORDER_VBGR</td>
<td class="value">4</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_subpixel_order_t</p>
</div>
<p class="api-heading">HintStyle</p>
<table>
<tr>
<td class="name">CAIRO_HINT_STYLE_DEFAULT</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_STYLE_NONE</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_STYLE_SLIGHT</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_STYLE_MEDIUM</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_STYLE_FULL</td>
<td class="value">4</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_hint_style_t</p>
</div>
<p class="api-heading">HintMetrics</p>
<table>
<tr>
<td class="name">CAIRO_HINT_METRICS_DEFAULT</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_METRICS_OFF</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_HINT_METRICS_ON</td>
<td class="value">2</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_hint_metrics_t</p>
</div>
<p class="api-heading">FontType</p>
<table>
<tr>
<td class="name">CAIRO_FONT_TYPE_TOY</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_TYPE_FT</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_TYPE_WIN32</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_TYPE_QUARTZ</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_FONT_TYPE_USER</td>
<td class="value">4</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_font_type_t</p>
</div>
<p class="api-heading">PathDataType</p>
<table>
<tr>
<td class="name">CAIRO_PATH_MOVE_TO</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_PATH_LINE_TO</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_PATH_CURVE_TO</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_PATH_CLOSE_PATH</td>
<td class="value">3</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_path_data_type_t</p>
</div>
<p class="api-heading">DeviceType</p>
<table>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_DRM</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_GL</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_SCRIPT</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_XCB</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_XLIB</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_XML</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_COGL</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_WIN32</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">CAIRO_DEVICE_TYPE_INVALID</td>
<td class="value">-1</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_device_type_t</p>
</div>
<p class="api-heading">SurfaceType</p>
<table>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_IMAGE</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_PDF</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_PS</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_XLIB</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_XCB</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_GLITZ</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_QUARTZ</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_WIN32</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_BEOS</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_DIRECTFB</td>
<td class="value">9</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_SVG</td>
<td class="value">10</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_OS2</td>
<td class="value">11</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_WIN32_PRINTING</td>
<td class="value">12</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_QUARTZ_IMAGE</td>
<td class="value">13</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_SCRIPT</td>
<td class="value">14</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_QT</td>
<td class="value">15</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_RECORDING</td>
<td class="value">16</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_VG</td>
<td class="value">17</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_GL</td>
<td class="value">18</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_DRM</td>
<td class="value">19</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_TEE</td>
<td class="value">20</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_XML</td>
<td class="value">21</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_SKIA</td>
<td class="value">22</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_SUBSURFACE</td>
<td class="value">23</td>
</tr>
<tr>
<td class="name">CAIRO_SURFACE_TYPE_COGL</td>
<td class="value">24</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_surface_type_t</p>
</div>
<p class="api-heading">Format</p>
<table>
<tr>
<td class="name">CAIRO_FORMAT_INVALID</td>
<td class="value">-1</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_ARGB32</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_RGB24</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_A8</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_A1</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_RGB16_565</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_FORMAT_RGB30</td>
<td class="value">5</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_format_t</p>
</div>
<p class="api-heading">PatternType</p>
<table>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_SOLID</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_SURFACE</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_LINEAR</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_RADIAL</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_MESH</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_PATTERN_TYPE_RASTER_SOURCE</td>
<td class="value">5</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_pattern_type_t</p>
</div>
<p class="api-heading">Extend</p>
<table>
<tr>
<td class="name">CAIRO_EXTEND_NONE</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_EXTEND_REPEAT</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_EXTEND_REFLECT</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_EXTEND_PAD</td>
<td class="value">3</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_extend_t</p>
</div>
<p class="api-heading">Filter</p>
<table>
<tr>
<td class="name">CAIRO_FILTER_FAST</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_FILTER_GOOD</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_FILTER_BEST</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">CAIRO_FILTER_NEAREST</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">CAIRO_FILTER_BILINEAR</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">CAIRO_FILTER_GAUSSIAN</td>
<td class="value">5</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_filter_t</p>
</div>
<p class="api-heading">RegionOverlap</p>
<table>
<tr>
<td class="name">CAIRO_REGION_OVERLAP_IN</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">CAIRO_REGION_OVERLAP_OUT</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">CAIRO_REGION_OVERLAP_PART</td>
<td class="value">2</td>
</tr>
</table>
<div class="api-notes">
  <p class="api-ctype">C type: cairo_region_overlap_t</p>
</div>
