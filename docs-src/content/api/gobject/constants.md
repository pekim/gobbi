+++
title = "constants"
+++
<p class="api-heading">PARAM_MASK</p>
<p class="api-doc">Mask containing the bits of #GParamSpec.flags which are reserved for GLib.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_PARAM_MASK</p>
</div>
<p class="api-heading">PARAM_STATIC_STRINGS</p>
<p class="api-doc">#GParamFlags value alias for %G_PARAM_STATIC_NAME | %G_PARAM_STATIC_NICK | %G_PARAM_STATIC_BLURB.

Since 2.13.0</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_PARAM_STATIC_STRINGS</p>
</div>
<p class="api-heading">PARAM_USER_SHIFT</p>
<p class="api-doc">Minimum shift count to be used for user defined flags, to be stored in
#GParamSpec.flags. The maximum allowed is 10.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_PARAM_USER_SHIFT</p>
</div>
<p class="api-heading">SIGNAL_FLAGS_MASK</p>
<p class="api-doc">A mask for all #GSignalFlags bits.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_SIGNAL_FLAGS_MASK</p>
</div>
<p class="api-heading">SIGNAL_MATCH_MASK</p>
<p class="api-doc">A mask for all #GSignalMatchType bits.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_SIGNAL_MATCH_MASK</p>
</div>
<p class="api-heading">TYPE_FLAG_RESERVED_ID_BIT</p>
<p class="api-doc">A bit in the type number that's supposed to be left untouched.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_FLAG_RESERVED_ID_BIT</p>
</div>
<p class="api-heading">TYPE_FUNDAMENTAL_MAX</p>
<p class="api-doc">An integer constant that represents the number of identifiers reserved
for types that are assigned at compile-time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_FUNDAMENTAL_MAX</p>
</div>
<p class="api-heading">TYPE_FUNDAMENTAL_SHIFT</p>
<p class="api-doc">Shift value used in converting numbers to type IDs.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_FUNDAMENTAL_SHIFT</p>
</div>
<p class="api-heading">TYPE_RESERVED_BSE_FIRST</p>
<p class="api-doc">First fundamental type number to create a new fundamental type id with
G_TYPE_MAKE_FUNDAMENTAL() reserved for BSE.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_RESERVED_BSE_FIRST</p>
</div>
<p class="api-heading">TYPE_RESERVED_BSE_LAST</p>
<p class="api-doc">Last fundamental type number reserved for BSE.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_RESERVED_BSE_LAST</p>
</div>
<p class="api-heading">TYPE_RESERVED_GLIB_FIRST</p>
<p class="api-doc">First fundamental type number to create a new fundamental type id with
G_TYPE_MAKE_FUNDAMENTAL() reserved for GLib.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_RESERVED_GLIB_FIRST</p>
</div>
<p class="api-heading">TYPE_RESERVED_GLIB_LAST</p>
<p class="api-doc">Last fundamental type number reserved for GLib.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_RESERVED_GLIB_LAST</p>
</div>
<p class="api-heading">TYPE_RESERVED_USER_FIRST</p>
<p class="api-doc">First available fundamental type number to create new fundamental
type id with G_TYPE_MAKE_FUNDAMENTAL().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TYPE_RESERVED_USER_FIRST</p>
</div>
<p class="api-heading">VALUE_NOCOPY_CONTENTS</p>
<p class="api-doc">If passed to G_VALUE_COLLECT(), allocated data won't be copied
but used verbatim. This does not affect ref-counted types like
objects.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VALUE_NOCOPY_CONTENTS</p>
</div>
