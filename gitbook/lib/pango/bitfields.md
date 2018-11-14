# `pango` bitfields

## `FontMask`

The bits in a #PangoFontMask correspond to fields in a
&num;PangoFontDescription that have been set.

C - `PangoFontMask`

### PANGO_FONT_MASK_FAMILY = 1
the font family is specified.

### PANGO_FONT_MASK_STYLE = 2
the font style is specified.

### PANGO_FONT_MASK_VARIANT = 4
the font variant is specified.

### PANGO_FONT_MASK_WEIGHT = 8
the font weight is specified.

### PANGO_FONT_MASK_STRETCH = 16
the font stretch is specified.

### PANGO_FONT_MASK_SIZE = 32
the font size is specified.

### PANGO_FONT_MASK_GRAVITY = 64
the font gravity is specified (Since: 1.16.)


