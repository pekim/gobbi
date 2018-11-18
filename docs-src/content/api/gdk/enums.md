+++
title = "enums"
+++
<p class="api-heading">AxisUse</p>
<p class="api-doc">An enumeration describing the way in which a device
axis (valuator) maps onto the predefined valuator
types that GTK+ understands.

Note that the X and Y axes are not really needed; pointer devices
report their location via the x/y members of events regardless. Whether
X and Y are present as axes depends on the GDK backend.</p>
<div class="api-notes">
  <p class="api-ctype">GdkAxisUse</p>
<table>
<tr>
<td class="name">GDK_AXIS_IGNORE</td>
<td class="value">0</td>
<td class="doc">the axis is ignored.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_X</td>
<td class="value">1</td>
<td class="doc">the axis is used as the x axis.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_Y</td>
<td class="value">2</td>
<td class="doc">the axis is used as the y axis.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_PRESSURE</td>
<td class="value">3</td>
<td class="doc">the axis is used for pressure information.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_XTILT</td>
<td class="value">4</td>
<td class="doc">the axis is used for x tilt information.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_YTILT</td>
<td class="value">5</td>
<td class="doc">the axis is used for y tilt information.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_WHEEL</td>
<td class="value">6</td>
<td class="doc">the axis is used for wheel information.</td>
</tr>
<tr>
<td class="name">GDK_AXIS_DISTANCE</td>
<td class="value">7</td>
<td class="doc">the axis is used for pen/tablet distance information. (Since: 3.22)</td>
</tr>
<tr>
<td class="name">GDK_AXIS_ROTATION</td>
<td class="value">8</td>
<td class="doc">the axis is used for pen rotation information. (Since: 3.22)</td>
</tr>
<tr>
<td class="name">GDK_AXIS_SLIDER</td>
<td class="value">9</td>
<td class="doc">the axis is used for pen slider information. (Since: 3.22)</td>
</tr>
<tr>
<td class="name">GDK_AXIS_LAST</td>
<td class="value">10</td>
<td class="doc">a constant equal to the numerically highest axis value.</td>
</tr>
</table>
</div>
<p class="api-heading">ByteOrder</p>
<p class="api-doc">A set of values describing the possible byte-orders
for storing pixel values in memory.</p>
<div class="api-notes">
  <p class="api-ctype">GdkByteOrder</p>
<table>
<tr>
<td class="name">GDK_LSB_FIRST</td>
<td class="value">0</td>
<td class="doc">The values are stored with the least-significant byte
  first. For instance, the 32-bit value 0xffeecc would be stored
  in memory as 0xcc, 0xee, 0xff, 0x00.</td>
</tr>
<tr>
<td class="name">GDK_MSB_FIRST</td>
<td class="value">1</td>
<td class="doc">The values are stored with the most-significant byte
  first. For instance, the 32-bit value 0xffeecc would be stored
  in memory as 0x00, 0xff, 0xee, 0xcc.</td>
</tr>
</table>
</div>
<p class="api-heading">CrossingMode</p>
<p class="api-doc">Specifies the crossing mode for #GdkEventCrossing.</p>
<div class="api-notes">
  <p class="api-ctype">GdkCrossingMode</p>
<table>
<tr>
<td class="name">GDK_CROSSING_NORMAL</td>
<td class="value">0</td>
<td class="doc">crossing because of pointer motion.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_GRAB</td>
<td class="value">1</td>
<td class="doc">crossing because a grab is activated.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_UNGRAB</td>
<td class="value">2</td>
<td class="doc">crossing because a grab is deactivated.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_GTK_GRAB</td>
<td class="value">3</td>
<td class="doc">crossing because a GTK+ grab is activated.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_GTK_UNGRAB</td>
<td class="value">4</td>
<td class="doc">crossing because a GTK+ grab is deactivated.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_STATE_CHANGED</td>
<td class="value">5</td>
<td class="doc">crossing because a GTK+ widget changed
  state (e.g. sensitivity).</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_TOUCH_BEGIN</td>
<td class="value">6</td>
<td class="doc">crossing because a touch sequence has begun,
  this event is synthetic as the pointer might have not left the window.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_TOUCH_END</td>
<td class="value">7</td>
<td class="doc">crossing because a touch sequence has ended,
  this event is synthetic as the pointer might have not left the window.</td>
</tr>
<tr>
<td class="name">GDK_CROSSING_DEVICE_SWITCH</td>
<td class="value">8</td>
<td class="doc">crossing because of a device switch (i.e.
  a mouse taking control of the pointer after a touch device), this event
  is synthetic as the pointer didn’t leave the window.</td>
</tr>
</table>
</div>
<p class="api-heading">CursorType</p>
<p class="api-doc">Predefined cursors.

Note that these IDs are directly taken from the X cursor font, and many
of these cursors are either not useful, or are not available on other platforms.

The recommended way to create cursors is to use gdk_cursor_new_from_name().</p>
<div class="api-notes">
  <p class="api-ctype">GdkCursorType</p>
<table>
<tr>
<td class="name">GDK_X_CURSOR</td>
<td class="value">0</td>
<td class="doc">![](X_cursor.png)</td>
</tr>
<tr>
<td class="name">GDK_ARROW</td>
<td class="value">2</td>
<td class="doc">![](arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_BASED_ARROW_DOWN</td>
<td class="value">4</td>
<td class="doc">![](based_arrow_down.png)</td>
</tr>
<tr>
<td class="name">GDK_BASED_ARROW_UP</td>
<td class="value">6</td>
<td class="doc">![](based_arrow_up.png)</td>
</tr>
<tr>
<td class="name">GDK_BOAT</td>
<td class="value">8</td>
<td class="doc">![](boat.png)</td>
</tr>
<tr>
<td class="name">GDK_BOGOSITY</td>
<td class="value">10</td>
<td class="doc">![](bogosity.png)</td>
</tr>
<tr>
<td class="name">GDK_BOTTOM_LEFT_CORNER</td>
<td class="value">12</td>
<td class="doc">![](bottom_left_corner.png)</td>
</tr>
<tr>
<td class="name">GDK_BOTTOM_RIGHT_CORNER</td>
<td class="value">14</td>
<td class="doc">![](bottom_right_corner.png)</td>
</tr>
<tr>
<td class="name">GDK_BOTTOM_SIDE</td>
<td class="value">16</td>
<td class="doc">![](bottom_side.png)</td>
</tr>
<tr>
<td class="name">GDK_BOTTOM_TEE</td>
<td class="value">18</td>
<td class="doc">![](bottom_tee.png)</td>
</tr>
<tr>
<td class="name">GDK_BOX_SPIRAL</td>
<td class="value">20</td>
<td class="doc">![](box_spiral.png)</td>
</tr>
<tr>
<td class="name">GDK_CENTER_PTR</td>
<td class="value">22</td>
<td class="doc">![](center_ptr.png)</td>
</tr>
<tr>
<td class="name">GDK_CIRCLE</td>
<td class="value">24</td>
<td class="doc">![](circle.png)</td>
</tr>
<tr>
<td class="name">GDK_CLOCK</td>
<td class="value">26</td>
<td class="doc">![](clock.png)</td>
</tr>
<tr>
<td class="name">GDK_COFFEE_MUG</td>
<td class="value">28</td>
<td class="doc">![](coffee_mug.png)</td>
</tr>
<tr>
<td class="name">GDK_CROSS</td>
<td class="value">30</td>
<td class="doc">![](cross.png)</td>
</tr>
<tr>
<td class="name">GDK_CROSS_REVERSE</td>
<td class="value">32</td>
<td class="doc">![](cross_reverse.png)</td>
</tr>
<tr>
<td class="name">GDK_CROSSHAIR</td>
<td class="value">34</td>
<td class="doc">![](crosshair.png)</td>
</tr>
<tr>
<td class="name">GDK_DIAMOND_CROSS</td>
<td class="value">36</td>
<td class="doc">![](diamond_cross.png)</td>
</tr>
<tr>
<td class="name">GDK_DOT</td>
<td class="value">38</td>
<td class="doc">![](dot.png)</td>
</tr>
<tr>
<td class="name">GDK_DOTBOX</td>
<td class="value">40</td>
<td class="doc">![](dotbox.png)</td>
</tr>
<tr>
<td class="name">GDK_DOUBLE_ARROW</td>
<td class="value">42</td>
<td class="doc">![](double_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_DRAFT_LARGE</td>
<td class="value">44</td>
<td class="doc">![](draft_large.png)</td>
</tr>
<tr>
<td class="name">GDK_DRAFT_SMALL</td>
<td class="value">46</td>
<td class="doc">![](draft_small.png)</td>
</tr>
<tr>
<td class="name">GDK_DRAPED_BOX</td>
<td class="value">48</td>
<td class="doc">![](draped_box.png)</td>
</tr>
<tr>
<td class="name">GDK_EXCHANGE</td>
<td class="value">50</td>
<td class="doc">![](exchange.png)</td>
</tr>
<tr>
<td class="name">GDK_FLEUR</td>
<td class="value">52</td>
<td class="doc">![](fleur.png)</td>
</tr>
<tr>
<td class="name">GDK_GOBBLER</td>
<td class="value">54</td>
<td class="doc">![](gobbler.png)</td>
</tr>
<tr>
<td class="name">GDK_GUMBY</td>
<td class="value">56</td>
<td class="doc">![](gumby.png)</td>
</tr>
<tr>
<td class="name">GDK_HAND1</td>
<td class="value">58</td>
<td class="doc">![](hand1.png)</td>
</tr>
<tr>
<td class="name">GDK_HAND2</td>
<td class="value">60</td>
<td class="doc">![](hand2.png)</td>
</tr>
<tr>
<td class="name">GDK_HEART</td>
<td class="value">62</td>
<td class="doc">![](heart.png)</td>
</tr>
<tr>
<td class="name">GDK_ICON</td>
<td class="value">64</td>
<td class="doc">![](icon.png)</td>
</tr>
<tr>
<td class="name">GDK_IRON_CROSS</td>
<td class="value">66</td>
<td class="doc">![](iron_cross.png)</td>
</tr>
<tr>
<td class="name">GDK_LEFT_PTR</td>
<td class="value">68</td>
<td class="doc">![](left_ptr.png)</td>
</tr>
<tr>
<td class="name">GDK_LEFT_SIDE</td>
<td class="value">70</td>
<td class="doc">![](left_side.png)</td>
</tr>
<tr>
<td class="name">GDK_LEFT_TEE</td>
<td class="value">72</td>
<td class="doc">![](left_tee.png)</td>
</tr>
<tr>
<td class="name">GDK_LEFTBUTTON</td>
<td class="value">74</td>
<td class="doc">![](leftbutton.png)</td>
</tr>
<tr>
<td class="name">GDK_LL_ANGLE</td>
<td class="value">76</td>
<td class="doc">![](ll_angle.png)</td>
</tr>
<tr>
<td class="name">GDK_LR_ANGLE</td>
<td class="value">78</td>
<td class="doc">![](lr_angle.png)</td>
</tr>
<tr>
<td class="name">GDK_MAN</td>
<td class="value">80</td>
<td class="doc">![](man.png)</td>
</tr>
<tr>
<td class="name">GDK_MIDDLEBUTTON</td>
<td class="value">82</td>
<td class="doc">![](middlebutton.png)</td>
</tr>
<tr>
<td class="name">GDK_MOUSE</td>
<td class="value">84</td>
<td class="doc">![](mouse.png)</td>
</tr>
<tr>
<td class="name">GDK_PENCIL</td>
<td class="value">86</td>
<td class="doc">![](pencil.png)</td>
</tr>
<tr>
<td class="name">GDK_PIRATE</td>
<td class="value">88</td>
<td class="doc">![](pirate.png)</td>
</tr>
<tr>
<td class="name">GDK_PLUS</td>
<td class="value">90</td>
<td class="doc">![](plus.png)</td>
</tr>
<tr>
<td class="name">GDK_QUESTION_ARROW</td>
<td class="value">92</td>
<td class="doc">![](question_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_RIGHT_PTR</td>
<td class="value">94</td>
<td class="doc">![](right_ptr.png)</td>
</tr>
<tr>
<td class="name">GDK_RIGHT_SIDE</td>
<td class="value">96</td>
<td class="doc">![](right_side.png)</td>
</tr>
<tr>
<td class="name">GDK_RIGHT_TEE</td>
<td class="value">98</td>
<td class="doc">![](right_tee.png)</td>
</tr>
<tr>
<td class="name">GDK_RIGHTBUTTON</td>
<td class="value">100</td>
<td class="doc">![](rightbutton.png)</td>
</tr>
<tr>
<td class="name">GDK_RTL_LOGO</td>
<td class="value">102</td>
<td class="doc">![](rtl_logo.png)</td>
</tr>
<tr>
<td class="name">GDK_SAILBOAT</td>
<td class="value">104</td>
<td class="doc">![](sailboat.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_DOWN_ARROW</td>
<td class="value">106</td>
<td class="doc">![](sb_down_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_H_DOUBLE_ARROW</td>
<td class="value">108</td>
<td class="doc">![](sb_h_double_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_LEFT_ARROW</td>
<td class="value">110</td>
<td class="doc">![](sb_left_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_RIGHT_ARROW</td>
<td class="value">112</td>
<td class="doc">![](sb_right_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_UP_ARROW</td>
<td class="value">114</td>
<td class="doc">![](sb_up_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SB_V_DOUBLE_ARROW</td>
<td class="value">116</td>
<td class="doc">![](sb_v_double_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_SHUTTLE</td>
<td class="value">118</td>
<td class="doc">![](shuttle.png)</td>
</tr>
<tr>
<td class="name">GDK_SIZING</td>
<td class="value">120</td>
<td class="doc">![](sizing.png)</td>
</tr>
<tr>
<td class="name">GDK_SPIDER</td>
<td class="value">122</td>
<td class="doc">![](spider.png)</td>
</tr>
<tr>
<td class="name">GDK_SPRAYCAN</td>
<td class="value">124</td>
<td class="doc">![](spraycan.png)</td>
</tr>
<tr>
<td class="name">GDK_STAR</td>
<td class="value">126</td>
<td class="doc">![](star.png)</td>
</tr>
<tr>
<td class="name">GDK_TARGET</td>
<td class="value">128</td>
<td class="doc">![](target.png)</td>
</tr>
<tr>
<td class="name">GDK_TCROSS</td>
<td class="value">130</td>
<td class="doc">![](tcross.png)</td>
</tr>
<tr>
<td class="name">GDK_TOP_LEFT_ARROW</td>
<td class="value">132</td>
<td class="doc">![](top_left_arrow.png)</td>
</tr>
<tr>
<td class="name">GDK_TOP_LEFT_CORNER</td>
<td class="value">134</td>
<td class="doc">![](top_left_corner.png)</td>
</tr>
<tr>
<td class="name">GDK_TOP_RIGHT_CORNER</td>
<td class="value">136</td>
<td class="doc">![](top_right_corner.png)</td>
</tr>
<tr>
<td class="name">GDK_TOP_SIDE</td>
<td class="value">138</td>
<td class="doc">![](top_side.png)</td>
</tr>
<tr>
<td class="name">GDK_TOP_TEE</td>
<td class="value">140</td>
<td class="doc">![](top_tee.png)</td>
</tr>
<tr>
<td class="name">GDK_TREK</td>
<td class="value">142</td>
<td class="doc">![](trek.png)</td>
</tr>
<tr>
<td class="name">GDK_UL_ANGLE</td>
<td class="value">144</td>
<td class="doc">![](ul_angle.png)</td>
</tr>
<tr>
<td class="name">GDK_UMBRELLA</td>
<td class="value">146</td>
<td class="doc">![](umbrella.png)</td>
</tr>
<tr>
<td class="name">GDK_UR_ANGLE</td>
<td class="value">148</td>
<td class="doc">![](ur_angle.png)</td>
</tr>
<tr>
<td class="name">GDK_WATCH</td>
<td class="value">150</td>
<td class="doc">![](watch.png)</td>
</tr>
<tr>
<td class="name">GDK_XTERM</td>
<td class="value">152</td>
<td class="doc">![](xterm.png)</td>
</tr>
<tr>
<td class="name">GDK_LAST_CURSOR</td>
<td class="value">153</td>
<td class="doc">last cursor type</td>
</tr>
<tr>
<td class="name">GDK_BLANK_CURSOR</td>
<td class="value">-2</td>
<td class="doc">Blank cursor. Since 2.16</td>
</tr>
<tr>
<td class="name">GDK_CURSOR_IS_PIXMAP</td>
<td class="value">-1</td>
<td class="doc">type of cursors constructed with
  gdk_cursor_new_from_pixbuf()</td>
</tr>
</table>
</div>
<p class="api-heading">DevicePadFeature</p>
<p class="api-doc">A pad feature.</p>
<div class="api-notes">
  <p class="api-ctype">GdkDevicePadFeature</p>
  <p class="api-since">since 3.22</p>
<table>
<tr>
<td class="name">GDK_DEVICE_PAD_FEATURE_BUTTON</td>
<td class="value">0</td>
<td class="doc">a button</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_PAD_FEATURE_RING</td>
<td class="value">1</td>
<td class="doc">a ring-shaped interactive area</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_PAD_FEATURE_STRIP</td>
<td class="value">2</td>
<td class="doc">a straight interactive area</td>
</tr>
</table>
</div>
<p class="api-heading">DeviceToolType</p>
<p class="api-doc">Indicates the specific type of tool being used being a tablet. Such as an
airbrush, pencil, etc.</p>
<div class="api-notes">
  <p class="api-ctype">GdkDeviceToolType</p>
  <p class="api-since">since 3.22</p>
<table>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">Tool is of an unknown type.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_PEN</td>
<td class="value">1</td>
<td class="doc">Tool is a standard tablet stylus.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_ERASER</td>
<td class="value">2</td>
<td class="doc">Tool is standard tablet eraser.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_BRUSH</td>
<td class="value">3</td>
<td class="doc">Tool is a brush stylus.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_PENCIL</td>
<td class="value">4</td>
<td class="doc">Tool is a pencil stylus.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_AIRBRUSH</td>
<td class="value">5</td>
<td class="doc">Tool is an airbrush stylus.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_MOUSE</td>
<td class="value">6</td>
<td class="doc">Tool is a mouse.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TOOL_TYPE_LENS</td>
<td class="value">7</td>
<td class="doc">Tool is a lens cursor.</td>
</tr>
</table>
</div>
<p class="api-heading">DeviceType</p>
<p class="api-doc">Indicates the device type. See [above][GdkDeviceManager.description]
for more information about the meaning of these device types.</p>
<div class="api-notes">
  <p class="api-ctype">GdkDeviceType</p>
<table>
<tr>
<td class="name">GDK_DEVICE_TYPE_MASTER</td>
<td class="value">0</td>
<td class="doc">Device is a master (or virtual) device. There will
                         be an associated focus indicator on the screen.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TYPE_SLAVE</td>
<td class="value">1</td>
<td class="doc">Device is a slave (or physical) device.</td>
</tr>
<tr>
<td class="name">GDK_DEVICE_TYPE_FLOATING</td>
<td class="value">2</td>
<td class="doc">Device is a physical device, currently not attached to
                           any virtual device.</td>
</tr>
</table>
</div>
<p class="api-heading">DragCancelReason</p>
<p class="api-doc">Used in #GdkDragContext to the reason of a cancelled DND operation.</p>
<div class="api-notes">
  <p class="api-ctype">GdkDragCancelReason</p>
  <p class="api-since">since 3.20</p>
<table>
<tr>
<td class="name">GDK_DRAG_CANCEL_NO_TARGET</td>
<td class="value">0</td>
<td class="doc">There is no suitable drop target.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_CANCEL_USER_CANCELLED</td>
<td class="value">1</td>
<td class="doc">Drag cancelled by the user</td>
</tr>
<tr>
<td class="name">GDK_DRAG_CANCEL_ERROR</td>
<td class="value">2</td>
<td class="doc">Unspecified error.</td>
</tr>
</table>
</div>
<p class="api-heading">DragProtocol</p>
<p class="api-doc">Used in #GdkDragContext to indicate the protocol according to
which DND is done.</p>
<div class="api-notes">
  <p class="api-ctype">GdkDragProtocol</p>
<table>
<tr>
<td class="name">GDK_DRAG_PROTO_NONE</td>
<td class="value">0</td>
<td class="doc">no protocol.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_MOTIF</td>
<td class="value">1</td>
<td class="doc">The Motif DND protocol. No longer supported</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_XDND</td>
<td class="value">2</td>
<td class="doc">The Xdnd protocol.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_ROOTWIN</td>
<td class="value">3</td>
<td class="doc">An extension to the Xdnd protocol for
 unclaimed root window drops.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_WIN32_DROPFILES</td>
<td class="value">4</td>
<td class="doc">The simple WM_DROPFILES protocol.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_OLE2</td>
<td class="value">5</td>
<td class="doc">The complex OLE2 DND protocol (not implemented).</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_LOCAL</td>
<td class="value">6</td>
<td class="doc">Intra-application DND.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_PROTO_WAYLAND</td>
<td class="value">7</td>
<td class="doc">Wayland DND protocol.</td>
</tr>
</table>
</div>
<p class="api-heading">EventType</p>
<p class="api-doc">Specifies the type of the event.

Do not confuse these events with the signals that GTK+ widgets emit.
Although many of these events result in corresponding signals being emitted,
the events are often transformed or filtered along the way.

In some language bindings, the values %GDK_2BUTTON_PRESS and
%GDK_3BUTTON_PRESS would translate into something syntactically
invalid (eg `Gdk.EventType.2ButtonPress`, where a
symbol is not allowed to start with a number). In that case, the
aliases %GDK_DOUBLE_BUTTON_PRESS and %GDK_TRIPLE_BUTTON_PRESS can
be used instead.</p>
<div class="api-notes">
  <p class="api-ctype">GdkEventType</p>
<table>
<tr>
<td class="name">GDK_NOTHING</td>
<td class="value">-1</td>
<td class="doc">a special code to indicate a null event.</td>
</tr>
<tr>
<td class="name">GDK_DELETE</td>
<td class="value">0</td>
<td class="doc">the window manager has requested that the toplevel window be
  hidden or destroyed, usually when the user clicks on a special icon in the
  title bar.</td>
</tr>
<tr>
<td class="name">GDK_DESTROY</td>
<td class="value">1</td>
<td class="doc">the window has been destroyed.</td>
</tr>
<tr>
<td class="name">GDK_EXPOSE</td>
<td class="value">2</td>
<td class="doc">all or part of the window has become visible and needs to be
  redrawn.</td>
</tr>
<tr>
<td class="name">GDK_MOTION_NOTIFY</td>
<td class="value">3</td>
<td class="doc">the pointer (usually a mouse) has moved.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON_PRESS</td>
<td class="value">4</td>
<td class="doc">a mouse button has been pressed.</td>
</tr>
<tr>
<td class="name">GDK_2BUTTON_PRESS</td>
<td class="value">5</td>
<td class="doc">a mouse button has been double-clicked (clicked twice
  within a short period of time). Note that each click also generates a
  %GDK_BUTTON_PRESS event.</td>
</tr>
<tr>
<td class="name">GDK_DOUBLE_BUTTON_PRESS</td>
<td class="value">5</td>
<td class="doc">alias for %GDK_2BUTTON_PRESS, added in 3.6.</td>
</tr>
<tr>
<td class="name">GDK_3BUTTON_PRESS</td>
<td class="value">6</td>
<td class="doc">a mouse button has been clicked 3 times in a short period
  of time. Note that each click also generates a %GDK_BUTTON_PRESS event.</td>
</tr>
<tr>
<td class="name">GDK_TRIPLE_BUTTON_PRESS</td>
<td class="value">6</td>
<td class="doc">alias for %GDK_3BUTTON_PRESS, added in 3.6.</td>
</tr>
<tr>
<td class="name">GDK_BUTTON_RELEASE</td>
<td class="value">7</td>
<td class="doc">a mouse button has been released.</td>
</tr>
<tr>
<td class="name">GDK_KEY_PRESS</td>
<td class="value">8</td>
<td class="doc">a key has been pressed.</td>
</tr>
<tr>
<td class="name">GDK_KEY_RELEASE</td>
<td class="value">9</td>
<td class="doc">a key has been released.</td>
</tr>
<tr>
<td class="name">GDK_ENTER_NOTIFY</td>
<td class="value">10</td>
<td class="doc">the pointer has entered the window.</td>
</tr>
<tr>
<td class="name">GDK_LEAVE_NOTIFY</td>
<td class="value">11</td>
<td class="doc">the pointer has left the window.</td>
</tr>
<tr>
<td class="name">GDK_FOCUS_CHANGE</td>
<td class="value">12</td>
<td class="doc">the keyboard focus has entered or left the window.</td>
</tr>
<tr>
<td class="name">GDK_CONFIGURE</td>
<td class="value">13</td>
<td class="doc">the size, position or stacking order of the window has changed.
  Note that GTK+ discards these events for %GDK_WINDOW_CHILD windows.</td>
</tr>
<tr>
<td class="name">GDK_MAP</td>
<td class="value">14</td>
<td class="doc">the window has been mapped.</td>
</tr>
<tr>
<td class="name">GDK_UNMAP</td>
<td class="value">15</td>
<td class="doc">the window has been unmapped.</td>
</tr>
<tr>
<td class="name">GDK_PROPERTY_NOTIFY</td>
<td class="value">16</td>
<td class="doc">a property on the window has been changed or deleted.</td>
</tr>
<tr>
<td class="name">GDK_SELECTION_CLEAR</td>
<td class="value">17</td>
<td class="doc">the application has lost ownership of a selection.</td>
</tr>
<tr>
<td class="name">GDK_SELECTION_REQUEST</td>
<td class="value">18</td>
<td class="doc">another application has requested a selection.</td>
</tr>
<tr>
<td class="name">GDK_SELECTION_NOTIFY</td>
<td class="value">19</td>
<td class="doc">a selection has been received.</td>
</tr>
<tr>
<td class="name">GDK_PROXIMITY_IN</td>
<td class="value">20</td>
<td class="doc">an input device has moved into contact with a sensing
  surface (e.g. a touchscreen or graphics tablet).</td>
</tr>
<tr>
<td class="name">GDK_PROXIMITY_OUT</td>
<td class="value">21</td>
<td class="doc">an input device has moved out of contact with a sensing
  surface.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_ENTER</td>
<td class="value">22</td>
<td class="doc">the mouse has entered the window while a drag is in progress.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_LEAVE</td>
<td class="value">23</td>
<td class="doc">the mouse has left the window while a drag is in progress.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_MOTION</td>
<td class="value">24</td>
<td class="doc">the mouse has moved in the window while a drag is in
  progress.</td>
</tr>
<tr>
<td class="name">GDK_DRAG_STATUS</td>
<td class="value">25</td>
<td class="doc">the status of the drag operation initiated by the window
  has changed.</td>
</tr>
<tr>
<td class="name">GDK_DROP_START</td>
<td class="value">26</td>
<td class="doc">a drop operation onto the window has started.</td>
</tr>
<tr>
<td class="name">GDK_DROP_FINISHED</td>
<td class="value">27</td>
<td class="doc">the drop operation initiated by the window has completed.</td>
</tr>
<tr>
<td class="name">GDK_CLIENT_EVENT</td>
<td class="value">28</td>
<td class="doc">a message has been received from another application.</td>
</tr>
<tr>
<td class="name">GDK_VISIBILITY_NOTIFY</td>
<td class="value">29</td>
<td class="doc">the window visibility status has changed.</td>
</tr>
<tr>
<td class="name">GDK_SCROLL</td>
<td class="value">31</td>
<td class="doc">the scroll wheel was turned</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_STATE</td>
<td class="value">32</td>
<td class="doc">the state of a window has changed. See #GdkWindowState
  for the possible window states</td>
</tr>
<tr>
<td class="name">GDK_SETTING</td>
<td class="value">33</td>
<td class="doc">a setting has been modified.</td>
</tr>
<tr>
<td class="name">GDK_OWNER_CHANGE</td>
<td class="value">34</td>
<td class="doc">the owner of a selection has changed. This event type
  was added in 2.6</td>
</tr>
<tr>
<td class="name">GDK_GRAB_BROKEN</td>
<td class="value">35</td>
<td class="doc">a pointer or keyboard grab was broken. This event type
  was added in 2.8.</td>
</tr>
<tr>
<td class="name">GDK_DAMAGE</td>
<td class="value">36</td>
<td class="doc">the content of the window has been changed. This event type
  was added in 2.14.</td>
</tr>
<tr>
<td class="name">GDK_TOUCH_BEGIN</td>
<td class="value">37</td>
<td class="doc">A new touch event sequence has just started. This event
  type was added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_TOUCH_UPDATE</td>
<td class="value">38</td>
<td class="doc">A touch event sequence has been updated. This event type
  was added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_TOUCH_END</td>
<td class="value">39</td>
<td class="doc">A touch event sequence has finished. This event type
  was added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_TOUCH_CANCEL</td>
<td class="value">40</td>
<td class="doc">A touch event sequence has been canceled. This event type
  was added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_SWIPE</td>
<td class="value">41</td>
<td class="doc">A touchpad swipe gesture event, the current state
  is determined by its phase field. This event type was added in 3.18.</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_PINCH</td>
<td class="value">42</td>
<td class="doc">A touchpad pinch gesture event, the current state
  is determined by its phase field. This event type was added in 3.18.</td>
</tr>
<tr>
<td class="name">GDK_PAD_BUTTON_PRESS</td>
<td class="value">43</td>
<td class="doc">A tablet pad button press event. This event type
  was added in 3.22.</td>
</tr>
<tr>
<td class="name">GDK_PAD_BUTTON_RELEASE</td>
<td class="value">44</td>
<td class="doc">A tablet pad button release event. This event type
  was added in 3.22.</td>
</tr>
<tr>
<td class="name">GDK_PAD_RING</td>
<td class="value">45</td>
<td class="doc">A tablet pad axis event from a "ring". This event type was
  added in 3.22.</td>
</tr>
<tr>
<td class="name">GDK_PAD_STRIP</td>
<td class="value">46</td>
<td class="doc">A tablet pad axis event from a "strip". This event type was
  added in 3.22.</td>
</tr>
<tr>
<td class="name">GDK_PAD_GROUP_MODE</td>
<td class="value">47</td>
<td class="doc">A tablet pad group mode change. This event type was
  added in 3.22.</td>
</tr>
<tr>
<td class="name">GDK_EVENT_LAST</td>
<td class="value">48</td>
<td class="doc">marks the end of the GdkEventType enumeration. Added in 2.18</td>
</tr>
</table>
</div>
<p class="api-heading">FilterReturn</p>
<p class="api-doc">Specifies the result of applying a #GdkFilterFunc to a native event.</p>
<div class="api-notes">
  <p class="api-ctype">GdkFilterReturn</p>
<table>
<tr>
<td class="name">GDK_FILTER_CONTINUE</td>
<td class="value">0</td>
<td class="doc">event not handled, continue processing.</td>
</tr>
<tr>
<td class="name">GDK_FILTER_TRANSLATE</td>
<td class="value">1</td>
<td class="doc">native event translated into a GDK event and stored
 in the `event` structure that was passed in.</td>
</tr>
<tr>
<td class="name">GDK_FILTER_REMOVE</td>
<td class="value">2</td>
<td class="doc">event handled, terminate processing.</td>
</tr>
</table>
</div>
<p class="api-heading">FullscreenMode</p>
<p class="api-doc">Indicates which monitor (in a multi-head setup) a window should span over
when in fullscreen mode.</p>
<div class="api-notes">
  <p class="api-ctype">GdkFullscreenMode</p>
  <p class="api-since">since 3.8</p>
<table>
<tr>
<td class="name">GDK_FULLSCREEN_ON_CURRENT_MONITOR</td>
<td class="value">0</td>
<td class="doc">Fullscreen on current monitor only.</td>
</tr>
<tr>
<td class="name">GDK_FULLSCREEN_ON_ALL_MONITORS</td>
<td class="value">1</td>
<td class="doc">Span across all monitors when fullscreen.</td>
</tr>
</table>
</div>
<p class="api-heading">GLError</p>
<p class="api-doc">Error enumeration for #GdkGLContext.</p>
<div class="api-notes">
  <p class="api-ctype">GdkGLError</p>
  <p class="api-since">since 3.16</p>
<table>
<tr>
<td class="name">GDK_GL_ERROR_NOT_AVAILABLE</td>
<td class="value">0</td>
<td class="doc">OpenGL support is not available</td>
</tr>
<tr>
<td class="name">GDK_GL_ERROR_UNSUPPORTED_FORMAT</td>
<td class="value">1</td>
<td class="doc">The requested visual format is not supported</td>
</tr>
<tr>
<td class="name">GDK_GL_ERROR_UNSUPPORTED_PROFILE</td>
<td class="value">2</td>
<td class="doc">The requested profile is not supported</td>
</tr>
</table>
</div>
<p class="api-heading">GrabOwnership</p>
<p class="api-doc">Defines how device grabs interact with other devices.</p>
<div class="api-notes">
  <p class="api-ctype">GdkGrabOwnership</p>
<table>
<tr>
<td class="name">GDK_OWNERSHIP_NONE</td>
<td class="value">0</td>
<td class="doc">All other devices’ events are allowed.</td>
</tr>
<tr>
<td class="name">GDK_OWNERSHIP_WINDOW</td>
<td class="value">1</td>
<td class="doc">Other devices’ events are blocked for the grab window.</td>
</tr>
<tr>
<td class="name">GDK_OWNERSHIP_APPLICATION</td>
<td class="value">2</td>
<td class="doc">Other devices’ events are blocked for the whole application.</td>
</tr>
</table>
</div>
<p class="api-heading">GrabStatus</p>
<p class="api-doc">Returned by gdk_device_grab(), gdk_pointer_grab() and gdk_keyboard_grab() to
indicate success or the reason for the failure of the grab attempt.</p>
<div class="api-notes">
  <p class="api-ctype">GdkGrabStatus</p>
<table>
<tr>
<td class="name">GDK_GRAB_SUCCESS</td>
<td class="value">0</td>
<td class="doc">the resource was successfully grabbed.</td>
</tr>
<tr>
<td class="name">GDK_GRAB_ALREADY_GRABBED</td>
<td class="value">1</td>
<td class="doc">the resource is actively grabbed by another client.</td>
</tr>
<tr>
<td class="name">GDK_GRAB_INVALID_TIME</td>
<td class="value">2</td>
<td class="doc">the resource was grabbed more recently than the
 specified time.</td>
</tr>
<tr>
<td class="name">GDK_GRAB_NOT_VIEWABLE</td>
<td class="value">3</td>
<td class="doc">the grab window or the @confine_to window are not
 viewable.</td>
</tr>
<tr>
<td class="name">GDK_GRAB_FROZEN</td>
<td class="value">4</td>
<td class="doc">the resource is frozen by an active grab of another client.</td>
</tr>
<tr>
<td class="name">GDK_GRAB_FAILED</td>
<td class="value">5</td>
<td class="doc">the grab failed for some other reason. Since 3.16</td>
</tr>
</table>
</div>
<p class="api-heading">Gravity</p>
<p class="api-doc">Defines the reference point of a window and the meaning of coordinates
passed to gtk_window_move(). See gtk_window_move() and the "implementation
notes" section of the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details.</p>
<div class="api-notes">
  <p class="api-ctype">GdkGravity</p>
<table>
<tr>
<td class="name">GDK_GRAVITY_NORTH_WEST</td>
<td class="value">1</td>
<td class="doc">the reference point is at the top left corner.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_NORTH</td>
<td class="value">2</td>
<td class="doc">the reference point is in the middle of the top edge.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_NORTH_EAST</td>
<td class="value">3</td>
<td class="doc">the reference point is at the top right corner.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_WEST</td>
<td class="value">4</td>
<td class="doc">the reference point is at the middle of the left edge.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_CENTER</td>
<td class="value">5</td>
<td class="doc">the reference point is at the center of the window.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_EAST</td>
<td class="value">6</td>
<td class="doc">the reference point is at the middle of the right edge.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_SOUTH_WEST</td>
<td class="value">7</td>
<td class="doc">the reference point is at the lower left corner.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_SOUTH</td>
<td class="value">8</td>
<td class="doc">the reference point is at the middle of the lower edge.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_SOUTH_EAST</td>
<td class="value">9</td>
<td class="doc">the reference point is at the lower right corner.</td>
</tr>
<tr>
<td class="name">GDK_GRAVITY_STATIC</td>
<td class="value">10</td>
<td class="doc">the reference point is at the top left corner of the
 window itself, ignoring window manager decorations.</td>
</tr>
</table>
</div>
<p class="api-heading">InputMode</p>
<p class="api-doc">An enumeration that describes the mode of an input device.</p>
<div class="api-notes">
  <p class="api-ctype">GdkInputMode</p>
<table>
<tr>
<td class="name">GDK_MODE_DISABLED</td>
<td class="value">0</td>
<td class="doc">the device is disabled and will not report any events.</td>
</tr>
<tr>
<td class="name">GDK_MODE_SCREEN</td>
<td class="value">1</td>
<td class="doc">the device is enabled. The device’s coordinate space
                  maps to the entire screen.</td>
</tr>
<tr>
<td class="name">GDK_MODE_WINDOW</td>
<td class="value">2</td>
<td class="doc">the device is enabled. The device’s coordinate space
                  is mapped to a single window. The manner in which this window
                  is chosen is undefined, but it will typically be the same
                  way in which the focus window for key events is determined.</td>
</tr>
</table>
</div>
<p class="api-heading">InputSource</p>
<p class="api-doc">An enumeration describing the type of an input device in general terms.</p>
<div class="api-notes">
  <p class="api-ctype">GdkInputSource</p>
<table>
<tr>
<td class="name">GDK_SOURCE_MOUSE</td>
<td class="value">0</td>
<td class="doc">the device is a mouse. (This will be reported for the core
                   pointer, even if it is something else, such as a trackball.)</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_PEN</td>
<td class="value">1</td>
<td class="doc">the device is a stylus of a graphics tablet or similar device.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_ERASER</td>
<td class="value">2</td>
<td class="doc">the device is an eraser. Typically, this would be the other end
                    of a stylus on a graphics tablet.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_CURSOR</td>
<td class="value">3</td>
<td class="doc">the device is a graphics tablet “puck” or similar device.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_KEYBOARD</td>
<td class="value">4</td>
<td class="doc">the device is a keyboard.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_TOUCHSCREEN</td>
<td class="value">5</td>
<td class="doc">the device is a direct-input touch device, such
    as a touchscreen or tablet. This device type has been added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_TOUCHPAD</td>
<td class="value">6</td>
<td class="doc">the device is an indirect touch device, such
    as a touchpad. This device type has been added in 3.4.</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_TRACKPOINT</td>
<td class="value">7</td>
<td class="doc">the device is a trackpoint. This device type has been
    added in 3.22</td>
</tr>
<tr>
<td class="name">GDK_SOURCE_TABLET_PAD</td>
<td class="value">8</td>
<td class="doc">the device is a "pad", a collection of buttons,
    rings and strips found in drawing tablets. This device type has been
    added in 3.22.</td>
</tr>
</table>
</div>
<p class="api-heading">ModifierIntent</p>
<p class="api-doc">This enum is used with gdk_keymap_get_modifier_mask()
in order to determine what modifiers the
currently used windowing system backend uses for particular
purposes. For example, on X11/Windows, the Control key is used for
invoking menu shortcuts (accelerators), whereas on Apple computers
it’s the Command key (which correspond to %GDK_CONTROL_MASK and
%GDK_MOD2_MASK, respectively).</p>
<div class="api-notes">
  <p class="api-ctype">GdkModifierIntent</p>
  <p class="api-since">since 3.4</p>
<table>
<tr>
<td class="name">GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR</td>
<td class="value">0</td>
<td class="doc">the primary modifier used to invoke
 menu accelerators.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_CONTEXT_MENU</td>
<td class="value">1</td>
<td class="doc">the modifier used to invoke context menus.
 Note that mouse button 3 always triggers context menus. When this modifier
 is not 0, it additionally triggers context menus when used with mouse button 1.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_EXTEND_SELECTION</td>
<td class="value">2</td>
<td class="doc">the modifier used to extend selections
 using `modifier`-click or `modifier`-cursor-key</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_MODIFY_SELECTION</td>
<td class="value">3</td>
<td class="doc">the modifier used to modify selections,
 which in most cases means toggling the clicked item into or out of the selection.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_NO_TEXT_INPUT</td>
<td class="value">4</td>
<td class="doc">when any of these modifiers is pressed, the
 key event cannot produce a symbol directly. This is meant to be used for
 input methods, and for use cases like typeahead search.</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_SHIFT_GROUP</td>
<td class="value">5</td>
<td class="doc">the modifier that switches between keyboard
 groups (AltGr on X11/Windows and Option/Alt on OS X).</td>
</tr>
<tr>
<td class="name">GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK</td>
<td class="value">6</td>
<td class="doc">The set of modifier masks accepted
as modifiers in accelerators. Needed because Command is mapped to MOD2 on
OSX, which is widely used, but on X11 MOD2 is NumLock and using that for a
mod key is problematic at best.
Ref: https://bugzilla.gnome.org/show_bug.cgi?id=736125.</td>
</tr>
</table>
</div>
<p class="api-heading">NotifyType</p>
<p class="api-doc">Specifies the kind of crossing for #GdkEventCrossing.

See the X11 protocol specification of LeaveNotify for
full details of crossing event generation.</p>
<div class="api-notes">
  <p class="api-ctype">GdkNotifyType</p>
<table>
<tr>
<td class="name">GDK_NOTIFY_ANCESTOR</td>
<td class="value">0</td>
<td class="doc">the window is entered from an ancestor or
  left towards an ancestor.</td>
</tr>
<tr>
<td class="name">GDK_NOTIFY_VIRTUAL</td>
<td class="value">1</td>
<td class="doc">the pointer moves between an ancestor and an
  inferior of the window.</td>
</tr>
<tr>
<td class="name">GDK_NOTIFY_INFERIOR</td>
<td class="value">2</td>
<td class="doc">the window is entered from an inferior or
  left towards an inferior.</td>
</tr>
<tr>
<td class="name">GDK_NOTIFY_NONLINEAR</td>
<td class="value">3</td>
<td class="doc">the window is entered from or left towards
  a window which is neither an ancestor nor an inferior.</td>
</tr>
<tr>
<td class="name">GDK_NOTIFY_NONLINEAR_VIRTUAL</td>
<td class="value">4</td>
<td class="doc">the pointer moves between two windows
  which are not ancestors of each other and the window is part of
  the ancestor chain between one of these windows and their least
  common ancestor.</td>
</tr>
<tr>
<td class="name">GDK_NOTIFY_UNKNOWN</td>
<td class="value">5</td>
<td class="doc">an unknown type of enter/leave event occurred.</td>
</tr>
</table>
</div>
<p class="api-heading">OwnerChange</p>
<p class="api-doc">Specifies why a selection ownership was changed.</p>
<div class="api-notes">
  <p class="api-ctype">GdkOwnerChange</p>
<table>
<tr>
<td class="name">GDK_OWNER_CHANGE_NEW_OWNER</td>
<td class="value">0</td>
<td class="doc">some other app claimed the ownership</td>
</tr>
<tr>
<td class="name">GDK_OWNER_CHANGE_DESTROY</td>
<td class="value">1</td>
<td class="doc">the window was destroyed</td>
</tr>
<tr>
<td class="name">GDK_OWNER_CHANGE_CLOSE</td>
<td class="value">2</td>
<td class="doc">the client was closed</td>
</tr>
</table>
</div>
<p class="api-heading">PropMode</p>
<p class="api-doc">Describes how existing data is combined with new data when
using gdk_property_change().</p>
<div class="api-notes">
  <p class="api-ctype">GdkPropMode</p>
<table>
<tr>
<td class="name">GDK_PROP_MODE_REPLACE</td>
<td class="value">0</td>
<td class="doc">the new data replaces the existing data.</td>
</tr>
<tr>
<td class="name">GDK_PROP_MODE_PREPEND</td>
<td class="value">1</td>
<td class="doc">the new data is prepended to the existing data.</td>
</tr>
<tr>
<td class="name">GDK_PROP_MODE_APPEND</td>
<td class="value">2</td>
<td class="doc">the new data is appended to the existing data.</td>
</tr>
</table>
</div>
<p class="api-heading">PropertyState</p>
<p class="api-doc">Specifies the type of a property change for a #GdkEventProperty.</p>
<div class="api-notes">
  <p class="api-ctype">guint</p>
<table>
<tr>
<td class="name">GDK_PROPERTY_NEW_VALUE</td>
<td class="value">0</td>
<td class="doc">the property value was changed.</td>
</tr>
<tr>
<td class="name">GDK_PROPERTY_DELETE</td>
<td class="value">1</td>
<td class="doc">the property was deleted.</td>
</tr>
</table>
</div>
<p class="api-heading">ScrollDirection</p>
<p class="api-doc">Specifies the direction for #GdkEventScroll.</p>
<div class="api-notes">
  <p class="api-ctype">GdkScrollDirection</p>
<table>
<tr>
<td class="name">GDK_SCROLL_UP</td>
<td class="value">0</td>
<td class="doc">the window is scrolled up.</td>
</tr>
<tr>
<td class="name">GDK_SCROLL_DOWN</td>
<td class="value">1</td>
<td class="doc">the window is scrolled down.</td>
</tr>
<tr>
<td class="name">GDK_SCROLL_LEFT</td>
<td class="value">2</td>
<td class="doc">the window is scrolled to the left.</td>
</tr>
<tr>
<td class="name">GDK_SCROLL_RIGHT</td>
<td class="value">3</td>
<td class="doc">the window is scrolled to the right.</td>
</tr>
<tr>
<td class="name">GDK_SCROLL_SMOOTH</td>
<td class="value">4</td>
<td class="doc">the scrolling is determined by the delta values
  in #GdkEventScroll. See gdk_event_get_scroll_deltas(). Since: 3.4</td>
</tr>
</table>
</div>
<p class="api-heading">SettingAction</p>
<p class="api-doc">Specifies the kind of modification applied to a setting in a
#GdkEventSetting.</p>
<div class="api-notes">
  <p class="api-ctype">GdkSettingAction</p>
<table>
<tr>
<td class="name">GDK_SETTING_ACTION_NEW</td>
<td class="value">0</td>
<td class="doc">a setting was added.</td>
</tr>
<tr>
<td class="name">GDK_SETTING_ACTION_CHANGED</td>
<td class="value">1</td>
<td class="doc">a setting was changed.</td>
</tr>
<tr>
<td class="name">GDK_SETTING_ACTION_DELETED</td>
<td class="value">2</td>
<td class="doc">a setting was deleted.</td>
</tr>
</table>
</div>
<p class="api-heading">Status</p>
<div class="api-notes">
  <p class="api-ctype">GdkStatus</p>
<table>
<tr>
<td class="name">GDK_OK</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">GDK_ERROR</td>
<td class="value">-1</td>
</tr>
<tr>
<td class="name">GDK_ERROR_PARAM</td>
<td class="value">-2</td>
</tr>
<tr>
<td class="name">GDK_ERROR_FILE</td>
<td class="value">-3</td>
</tr>
<tr>
<td class="name">GDK_ERROR_MEM</td>
<td class="value">-4</td>
</tr>
</table>
</div>
<p class="api-heading">SubpixelLayout</p>
<p class="api-doc">This enumeration describes how the red, green and blue components
of physical pixels on an output device are laid out.</p>
<div class="api-notes">
  <p class="api-ctype">GdkSubpixelLayout</p>
  <p class="api-since">since 3.22</p>
<table>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">The layout is not known</td>
</tr>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_NONE</td>
<td class="value">1</td>
<td class="doc">Not organized in this way</td>
</tr>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB</td>
<td class="value">2</td>
<td class="doc">The layout is horizontal, the order is RGB</td>
</tr>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR</td>
<td class="value">3</td>
<td class="doc">The layout is horizontal, the order is BGR</td>
</tr>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB</td>
<td class="value">4</td>
<td class="doc">The layout is vertical, the order is RGB</td>
</tr>
<tr>
<td class="name">GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR</td>
<td class="value">5</td>
<td class="doc">The layout is vertical, the order is BGR</td>
</tr>
</table>
</div>
<p class="api-heading">TouchpadGesturePhase</p>
<p class="api-doc">Specifies the current state of a touchpad gesture. All gestures are
guaranteed to begin with an event with phase %GDK_TOUCHPAD_GESTURE_PHASE_BEGIN,
followed by 0 or several events with phase %GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.

A finished gesture may have 2 possible outcomes, an event with phase
%GDK_TOUCHPAD_GESTURE_PHASE_END will be emitted when the gesture is
considered successful, this should be used as the hint to perform any
permanent changes.

Cancelled gestures may be so for a variety of reasons, due to hardware
or the compositor, or due to the gesture recognition layers hinting the
gesture did not finish resolutely (eg. a 3rd finger being added during
a pinch gesture). In these cases, the last event will report the phase
%GDK_TOUCHPAD_GESTURE_PHASE_CANCEL, this should be used as a hint
to undo any visible/permanent changes that were done throughout the
progress of the gesture.

See also #GdkEventTouchpadSwipe and #GdkEventTouchpadPinch.</p>
<div class="api-notes">
  <p class="api-ctype">GdkTouchpadGesturePhase</p>
<table>
<tr>
<td class="name">GDK_TOUCHPAD_GESTURE_PHASE_BEGIN</td>
<td class="value">0</td>
<td class="doc">The gesture has begun.</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_GESTURE_PHASE_UPDATE</td>
<td class="value">1</td>
<td class="doc">The gesture has been updated.</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_GESTURE_PHASE_END</td>
<td class="value">2</td>
<td class="doc">The gesture was finished, changes
  should be permanently applied.</td>
</tr>
<tr>
<td class="name">GDK_TOUCHPAD_GESTURE_PHASE_CANCEL</td>
<td class="value">3</td>
<td class="doc">The gesture was cancelled, all
  changes should be undone.</td>
</tr>
</table>
</div>
<p class="api-heading">VisibilityState</p>
<p class="api-doc">Specifies the visiblity status of a window for a #GdkEventVisibility.</p>
<div class="api-notes">
  <p class="api-ctype">GdkVisibilityState</p>
<table>
<tr>
<td class="name">GDK_VISIBILITY_UNOBSCURED</td>
<td class="value">0</td>
<td class="doc">the window is completely visible.</td>
</tr>
<tr>
<td class="name">GDK_VISIBILITY_PARTIAL</td>
<td class="value">1</td>
<td class="doc">the window is partially visible.</td>
</tr>
<tr>
<td class="name">GDK_VISIBILITY_FULLY_OBSCURED</td>
<td class="value">2</td>
<td class="doc">the window is not visible at all.</td>
</tr>
</table>
</div>
<p class="api-heading">VisualType</p>
<p class="api-doc">A set of values that describe the manner in which the pixel values
for a visual are converted into RGB values for display.</p>
<div class="api-notes">
  <p class="api-ctype">GdkVisualType</p>
<table>
<tr>
<td class="name">GDK_VISUAL_STATIC_GRAY</td>
<td class="value">0</td>
<td class="doc">Each pixel value indexes a grayscale value
    directly.</td>
</tr>
<tr>
<td class="name">GDK_VISUAL_GRAYSCALE</td>
<td class="value">1</td>
<td class="doc">Each pixel is an index into a color map that
    maps pixel values into grayscale values. The color map can be
    changed by an application.</td>
</tr>
<tr>
<td class="name">GDK_VISUAL_STATIC_COLOR</td>
<td class="value">2</td>
<td class="doc">Each pixel value is an index into a predefined,
    unmodifiable color map that maps pixel values into RGB values.</td>
</tr>
<tr>
<td class="name">GDK_VISUAL_PSEUDO_COLOR</td>
<td class="value">3</td>
<td class="doc">Each pixel is an index into a color map that
    maps pixel values into rgb values. The color map can be changed by
    an application.</td>
</tr>
<tr>
<td class="name">GDK_VISUAL_TRUE_COLOR</td>
<td class="value">4</td>
<td class="doc">Each pixel value directly contains red, green,
    and blue components. Use gdk_visual_get_red_pixel_details(), etc,
    to obtain information about how the components are assembled into
    a pixel value.</td>
</tr>
<tr>
<td class="name">GDK_VISUAL_DIRECT_COLOR</td>
<td class="value">5</td>
<td class="doc">Each pixel value contains red, green, and blue
    components as for %GDK_VISUAL_TRUE_COLOR, but the components are
    mapped via a color table into the final output table instead of
    being converted directly.</td>
</tr>
</table>
</div>
<p class="api-heading">WindowEdge</p>
<p class="api-doc">Determines a window edge or corner.</p>
<div class="api-notes">
  <p class="api-ctype">GdkWindowEdge</p>
<table>
<tr>
<td class="name">GDK_WINDOW_EDGE_NORTH_WEST</td>
<td class="value">0</td>
<td class="doc">the top left corner.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_NORTH</td>
<td class="value">1</td>
<td class="doc">the top edge.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_NORTH_EAST</td>
<td class="value">2</td>
<td class="doc">the top right corner.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_WEST</td>
<td class="value">3</td>
<td class="doc">the left edge.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_EAST</td>
<td class="value">4</td>
<td class="doc">the right edge.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_SOUTH_WEST</td>
<td class="value">5</td>
<td class="doc">the lower left corner.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_SOUTH</td>
<td class="value">6</td>
<td class="doc">the lower edge.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_EDGE_SOUTH_EAST</td>
<td class="value">7</td>
<td class="doc">the lower right corner.</td>
</tr>
</table>
</div>
<p class="api-heading">WindowType</p>
<p class="api-doc">Describes the kind of window.</p>
<div class="api-notes">
  <p class="api-ctype">GdkWindowType</p>
<table>
<tr>
<td class="name">GDK_WINDOW_ROOT</td>
<td class="value">0</td>
<td class="doc">root window; this window has no parent, covers the entire
 screen, and is created by the window system</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TOPLEVEL</td>
<td class="value">1</td>
<td class="doc">toplevel window (used to implement #GtkWindow)</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_CHILD</td>
<td class="value">2</td>
<td class="doc">child window (used to implement e.g. #GtkEntry)</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TEMP</td>
<td class="value">3</td>
<td class="doc">override redirect temporary window (used to implement
 #GtkMenu)</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_FOREIGN</td>
<td class="value">4</td>
<td class="doc">foreign window (see gdk_window_foreign_new())</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_OFFSCREEN</td>
<td class="value">5</td>
<td class="doc">offscreen window (see
 [Offscreen Windows][OFFSCREEN-WINDOWS]). Since 2.18</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_SUBSURFACE</td>
<td class="value">6</td>
<td class="doc">subsurface-based window; This window is visually
 tied to a toplevel, and is moved/stacked with it. Currently this window
 type is only implemented in Wayland. Since 3.14</td>
</tr>
</table>
</div>
<p class="api-heading">WindowTypeHint</p>
<p class="api-doc">These are hints for the window manager that indicate what type of function
the window has. The window manager can use this when determining decoration
and behaviour of the window. The hint must be set before mapping the window.

See the [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
specification for more details about window types.</p>
<div class="api-notes">
  <p class="api-ctype">GdkWindowTypeHint</p>
<table>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_NORMAL</td>
<td class="value">0</td>
<td class="doc">Normal toplevel window.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_DIALOG</td>
<td class="value">1</td>
<td class="doc">Dialog window.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_MENU</td>
<td class="value">2</td>
<td class="doc">Window used to implement a menu; GTK+ uses
 this hint only for torn-off menus, see #GtkTearoffMenuItem.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_TOOLBAR</td>
<td class="value">3</td>
<td class="doc">Window used to implement toolbars.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_SPLASHSCREEN</td>
<td class="value">4</td>
<td class="doc">Window used to display a splash
 screen during application startup.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_UTILITY</td>
<td class="value">5</td>
<td class="doc">Utility windows which are not detached
 toolbars or dialogs.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_DOCK</td>
<td class="value">6</td>
<td class="doc">Used for creating dock or panel windows.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_DESKTOP</td>
<td class="value">7</td>
<td class="doc">Used for creating the desktop background
 window.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU</td>
<td class="value">8</td>
<td class="doc">A menu that belongs to a menubar.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_POPUP_MENU</td>
<td class="value">9</td>
<td class="doc">A menu that does not belong to a menubar,
 e.g. a context menu.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_TOOLTIP</td>
<td class="value">10</td>
<td class="doc">A tooltip.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_NOTIFICATION</td>
<td class="value">11</td>
<td class="doc">A notification - typically a “bubble”
 that belongs to a status icon.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_COMBO</td>
<td class="value">12</td>
<td class="doc">A popup from a combo box.</td>
</tr>
<tr>
<td class="name">GDK_WINDOW_TYPE_HINT_DND</td>
<td class="value">13</td>
<td class="doc">A window that is used to implement a DND cursor.</td>
</tr>
</table>
</div>
<p class="api-heading">WindowWindowClass</p>
<p class="api-doc">@GDK_INPUT_OUTPUT windows are the standard kind of window you might expect.
Such windows receive events and are also displayed on screen.
@GDK_INPUT_ONLY windows are invisible; they are usually placed above other
windows in order to trap or filter the events. You can’t draw on
@GDK_INPUT_ONLY windows.</p>
<div class="api-notes">
  <p class="api-ctype">GdkWindowWindowClass</p>
<table>
<tr>
<td class="name">GDK_INPUT_OUTPUT</td>
<td class="value">0</td>
<td class="doc">window for graphics and events</td>
</tr>
<tr>
<td class="name">GDK_INPUT_ONLY</td>
<td class="value">1</td>
<td class="doc">window for events only</td>
</tr>
</table>
</div>