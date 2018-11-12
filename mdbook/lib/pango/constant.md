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

n`

The #PangoDirection type represents a direction in the
Unicode bidirectional algorithm; not every value in this
enumeration makes sense for every usage of #PangoDirection;
for example, the return value of pango_unichar_direction()
and pango_find_base_dir() cannot be %PANGO_DIRECTION_WEAK_LTR
or %PANGO_DIRECTION_WEAK_RTL, since every character is either
neutral or has a strong direction; on the other hand
%PANGO_DIRECTION_NEUTRAL doesn't make sense to pass
to pango_itemize_with_base_dir().

The %PANGO_DIRECTION_TTB_LTR, %PANGO_DIRECTION_TTB_RTL
values come from an earlier interpretation of this
enumeration as the writing direction of a block of
text and are no longer used; See #PangoGravity for how
vertical text is handled in Pango.

C - `PangoDirection`

## `EllipsizeMode`

The #PangoEllipsizeMode type describes what sort of (if any)
ellipsization should be applied to a line of text. In
the ellipsization process characters are removed from the
text in order to make it fit to a given width and replaced
with an ellipsis.

C - `PangoEllipsizeMode`

## `Gravity`

The #PangoGravity type represents the orientation of glyphs in a segment
of text.  This is useful when rendering vertical text layouts.  In
those situations, the layout is rotated using a non-identity PangoMatrix,
and then glyph orientation is controlled using #PangoGravity.
Not every value in this enumeration makes sense for every usage of
&num;PangoGravity; for example, %PANGO_GRAVITY_AUTO only can be passed to
pango_context_set_base_gravity() and can only be returned by
pango_context_get_base_gravity().

See also: #PangoGravityHint

C - `PangoGravity`

## `GravityHint`

The #PangoGravityHint defines how horizontal scripts should behave in a
vertical context.  That is, English excerpt in a vertical paragraph for
example.

See #PangoGravity.

C - `PangoGravityHint`

## `RenderPart`

#PangoRenderPart defines different items to render for such
purposes as setting colors.

C - `PangoRenderPart`

## `Script`

The #PangoScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard.
Note that new types may be added in the future. Applications should be ready
to handle unknown values.  This enumeration is interchangeable with
&num;GUnicodeScript.  See <ulink
url="http://www.unicode.org/reports/tr24/">Unicode Standard Annex
&num;24: Script names</ulink>.

C - `PangoScript`

## `Stretch`

An enumeration specifying the width of the font relative to other designs
within a family.

C - `PangoStretch`

## `Style`

An enumeration specifying the various slant styles possible for a font.

C - `PangoStyle`

## `TabAlign`

A #PangoTabAlign specifies where a tab stop appears relative to the text.

C - `PangoTabAlign`

## `Underline`

The #PangoUnderline enumeration is used to specify
whether text should be underlined, and if so, the type
of underlining.

C - `PangoUnderline`

## `Variant`

An enumeration specifying capitalization variant of the font.

C - `PangoVariant`

## `Weight`

An enumeration specifying the weight (boldness) of a font. This is a numerical
value ranging from 100 to 1000, but there are some predefined values:

C - `PangoWeight`

## `WrapMode`

A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.

C - `PangoWrapMode`

