# `pango` enums

## `Alignment`

A #PangoAlignment describes how to align the lines of a #PangoLayout within the
available space. If the #PangoLayout is set to justify
using pango_layout_set_justify(), this only has effect for partial lines.

C - `PangoAlignment`

## `AttrType`

The #PangoAttrType
distinguishes between different types of attributes. Along with the
predefined values, it is possible to allocate additional values
for custom attributes using pango_attr_type_register(). The predefined
values are given below. The type of structure used to store the
attribute is listed in parentheses after the description.

C - `PangoAttrType`

## `BidiType`

The #PangoBidiType type represents the bidirectional character
type of a Unicode character as specified by the
<ulink url="http://www.unicode.org/reports/tr9/">Unicode bidirectional algorithm</ulink>.

C - `PangoBidiType`

## `CoverageLevel`

Used to indicate how well a font can represent a particular Unicode
character point for a particular script.

C - `PangoCoverageLevel`

## `Direction`

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

