# `pango` Constants

## `ANALYSIS_FLAG_CENTERED_BASELINE`

Whether the segment should be shifted to center around the baseline.
Used in vertical writing directions mostly.

C - `PANGO_ANALYSIS_FLAG_CENTERED_BASELINE`

## `ANALYSIS_FLAG_IS_ELLIPSIS`

This flag is used to mark runs that hold ellipsized text,
in an ellipsized layout.

C - `PANGO_ANALYSIS_FLAG_IS_ELLIPSIS`

## `ATTR_INDEX_FROM_TEXT_BEGINNING`

This value can be used to set the start_index member of a #PangoAttribute
such that the attribute covers from the beginning of the text.

C - `PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING`

## `SCALE`

The %PANGO_SCALE macro represents the scale between dimensions used
for Pango distances and device units. (The definition of device
units is dependent on the output device; it will typically be pixels
for a screen, and points for a printer.) %PANGO_SCALE is currently
1024, but this may be changed in the future.

When setting font sizes, device units are always considered to be
points (as in "12 point font"), rather than pixels.

C - `PANGO_SCALE`

