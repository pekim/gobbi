# `atk` Enums

## `CoordType`

Specifies how xy coordinates are to be interpreted. Used by functions such
as atk_component_get_position() and atk_text_get_character_extents()

C - `AtkCoordType`

## `KeyEventType`

Specifies the type of a keyboard evemt.

C - `AtkKeyEventType`

## `Layer`

Describes the layer of a component

These enumerated "layer values" are used when determining which UI
rendering layer a component is drawn into, which can help in making
determinations of when components occlude one another.

C - `AtkLayer`

## `RelationType`

Describes the type of the relation

C - `AtkRelationType`

## `Role`

Describes the role of an object

These are the built-in enumerated roles that UI components can have in
ATK.  Other roles may be added at runtime, so an AtkRole >=
ATK_ROLE_LAST_DEFINED is not necessarily an error.

C - `AtkRole`

## `StateType`

The possible types of states of an object

C - `AtkStateType`

## `TextAttribute`

Describes the text attributes supported

C - `AtkTextAttribute`

## `TextBoundary`

Text boundary types used for specifying boundaries for regions of text.
This enumeration is deprecated since 2.9.4 and should not be used. Use
AtkTextGranularity with #atk_text_get_string_at_offset instead.

C - `AtkTextBoundary`

## `TextClipType`

Describes the type of clipping required.

C - `AtkTextClipType`

## `TextGranularity`

Text granularity types used for specifying the granularity of the region of
text we are interested in.

C - `AtkTextGranularity`

## `ValueType`

Default types for a given value. Those are defined in order to
easily get localized strings to describe a given value or a given
subrange, using atk_value_type_get_localized_name().

C - `AtkValueType`

