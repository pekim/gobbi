// Code generated - DO NOT EDIT.

package atk

// Coordtype is a representation of the C type CoordType.
type Coordtype int

const (
	Coordtype_Screen Coordtype = 0
	Coordtype_Window Coordtype = 1
)

// Keyeventtype is a representation of the C type KeyEventType.
type Keyeventtype int

const (
	Keyeventtype_Press       Keyeventtype = 0
	Keyeventtype_Release     Keyeventtype = 1
	Keyeventtype_LastDefined Keyeventtype = 2
)

// Layer is a representation of the C type Layer.
type Layer int

const (
	Layer_Invalid    Layer = 0
	Layer_Background Layer = 1
	Layer_Canvas     Layer = 2
	Layer_Widget     Layer = 3
	Layer_Mdi        Layer = 4
	Layer_Popup      Layer = 5
	Layer_Overlay    Layer = 6
	Layer_Window     Layer = 7
)

// Relationtype is a representation of the C type RelationType.
type Relationtype int

const (
	Relationtype_Null           Relationtype = 0
	Relationtype_ControlledBy   Relationtype = 1
	Relationtype_ControllerFor  Relationtype = 2
	Relationtype_LabelFor       Relationtype = 3
	Relationtype_LabelledBy     Relationtype = 4
	Relationtype_MemberOf       Relationtype = 5
	Relationtype_NodeChildOf    Relationtype = 6
	Relationtype_FlowsTo        Relationtype = 7
	Relationtype_FlowsFrom      Relationtype = 8
	Relationtype_SubwindowOf    Relationtype = 9
	Relationtype_Embeds         Relationtype = 10
	Relationtype_EmbeddedBy     Relationtype = 11
	Relationtype_PopupFor       Relationtype = 12
	Relationtype_ParentWindowOf Relationtype = 13
	Relationtype_DescribedBy    Relationtype = 14
	Relationtype_DescriptionFor Relationtype = 15
	Relationtype_NodeParentOf   Relationtype = 16
	Relationtype_Details        Relationtype = 17
	Relationtype_DetailsFor     Relationtype = 18
	Relationtype_ErrorMessage   Relationtype = 19
	Relationtype_ErrorFor       Relationtype = 20
	Relationtype_LastDefined    Relationtype = 21
)

// Role is a representation of the C type Role.
type Role int

const (
	Role_Invalid              Role = 0
	Role_AcceleratorLabel     Role = 1
	Role_Alert                Role = 2
	Role_Animation            Role = 3
	Role_Arrow                Role = 4
	Role_Calendar             Role = 5
	Role_Canvas               Role = 6
	Role_CheckBox             Role = 7
	Role_CheckMenuItem        Role = 8
	Role_ColorChooser         Role = 9
	Role_ColumnHeader         Role = 10
	Role_ComboBox             Role = 11
	Role_DateEditor           Role = 12
	Role_DesktopIcon          Role = 13
	Role_DesktopFrame         Role = 14
	Role_Dial                 Role = 15
	Role_Dialog               Role = 16
	Role_DirectoryPane        Role = 17
	Role_DrawingArea          Role = 18
	Role_FileChooser          Role = 19
	Role_Filler               Role = 20
	Role_FontChooser          Role = 21
	Role_Frame                Role = 22
	Role_GlassPane            Role = 23
	Role_HtmlContainer        Role = 24
	Role_Icon                 Role = 25
	Role_Image                Role = 26
	Role_InternalFrame        Role = 27
	Role_Label                Role = 28
	Role_LayeredPane          Role = 29
	Role_List                 Role = 30
	Role_ListItem             Role = 31
	Role_Menu                 Role = 32
	Role_MenuBar              Role = 33
	Role_MenuItem             Role = 34
	Role_OptionPane           Role = 35
	Role_PageTab              Role = 36
	Role_PageTabList          Role = 37
	Role_Panel                Role = 38
	Role_PasswordText         Role = 39
	Role_PopupMenu            Role = 40
	Role_ProgressBar          Role = 41
	Role_PushButton           Role = 42
	Role_RadioButton          Role = 43
	Role_RadioMenuItem        Role = 44
	Role_RootPane             Role = 45
	Role_RowHeader            Role = 46
	Role_ScrollBar            Role = 47
	Role_ScrollPane           Role = 48
	Role_Separator            Role = 49
	Role_Slider               Role = 50
	Role_SplitPane            Role = 51
	Role_SpinButton           Role = 52
	Role_Statusbar            Role = 53
	Role_Table                Role = 54
	Role_TableCell            Role = 55
	Role_TableColumnHeader    Role = 56
	Role_TableRowHeader       Role = 57
	Role_TearOffMenuItem      Role = 58
	Role_Terminal             Role = 59
	Role_Text                 Role = 60
	Role_ToggleButton         Role = 61
	Role_ToolBar              Role = 62
	Role_ToolTip              Role = 63
	Role_Tree                 Role = 64
	Role_TreeTable            Role = 65
	Role_Unknown              Role = 66
	Role_Viewport             Role = 67
	Role_Window               Role = 68
	Role_Header               Role = 69
	Role_Footer               Role = 70
	Role_Paragraph            Role = 71
	Role_Ruler                Role = 72
	Role_Application          Role = 73
	Role_Autocomplete         Role = 74
	Role_EditBar              Role = 75
	Role_Embedded             Role = 76
	Role_Entry                Role = 77
	Role_Chart                Role = 78
	Role_Caption              Role = 79
	Role_DocumentFrame        Role = 80
	Role_Heading              Role = 81
	Role_Page                 Role = 82
	Role_Section              Role = 83
	Role_RedundantObject      Role = 84
	Role_Form                 Role = 85
	Role_Link                 Role = 86
	Role_InputMethodWindow    Role = 87
	Role_TableRow             Role = 88
	Role_TreeItem             Role = 89
	Role_DocumentSpreadsheet  Role = 90
	Role_DocumentPresentation Role = 91
	Role_DocumentText         Role = 92
	Role_DocumentWeb          Role = 93
	Role_DocumentEmail        Role = 94
	Role_Comment              Role = 95
	Role_ListBox              Role = 96
	Role_Grouping             Role = 97
	Role_ImageMap             Role = 98
	Role_Notification         Role = 99
	Role_InfoBar              Role = 100
	Role_LevelBar             Role = 101
	Role_TitleBar             Role = 102
	Role_BlockQuote           Role = 103
	Role_Audio                Role = 104
	Role_Video                Role = 105
	Role_Definition           Role = 106
	Role_Article              Role = 107
	Role_Landmark             Role = 108
	Role_Log                  Role = 109
	Role_Marquee              Role = 110
	Role_Math                 Role = 111
	Role_Rating               Role = 112
	Role_Timer                Role = 113
	Role_DescriptionList      Role = 114
	Role_DescriptionTerm      Role = 115
	Role_DescriptionValue     Role = 116
	Role_Static               Role = 117
	Role_MathFraction         Role = 118
	Role_MathRoot             Role = 119
	Role_Subscript            Role = 120
	Role_Superscript          Role = 121
	Role_Footnote             Role = 122
	Role_LastDefined          Role = 123
)

// Statetype is a representation of the C type StateType.
type Statetype int

const (
	Statetype_Invalid                Statetype = 0
	Statetype_Active                 Statetype = 1
	Statetype_Armed                  Statetype = 2
	Statetype_Busy                   Statetype = 3
	Statetype_Checked                Statetype = 4
	Statetype_Defunct                Statetype = 5
	Statetype_Editable               Statetype = 6
	Statetype_Enabled                Statetype = 7
	Statetype_Expandable             Statetype = 8
	Statetype_Expanded               Statetype = 9
	Statetype_Focusable              Statetype = 10
	Statetype_Focused                Statetype = 11
	Statetype_Horizontal             Statetype = 12
	Statetype_Iconified              Statetype = 13
	Statetype_Modal                  Statetype = 14
	Statetype_MultiLine              Statetype = 15
	Statetype_Multiselectable        Statetype = 16
	Statetype_Opaque                 Statetype = 17
	Statetype_Pressed                Statetype = 18
	Statetype_Resizable              Statetype = 19
	Statetype_Selectable             Statetype = 20
	Statetype_Selected               Statetype = 21
	Statetype_Sensitive              Statetype = 22
	Statetype_Showing                Statetype = 23
	Statetype_SingleLine             Statetype = 24
	Statetype_Stale                  Statetype = 25
	Statetype_Transient              Statetype = 26
	Statetype_Vertical               Statetype = 27
	Statetype_Visible                Statetype = 28
	Statetype_ManagesDescendants     Statetype = 29
	Statetype_Indeterminate          Statetype = 30
	Statetype_Truncated              Statetype = 31
	Statetype_Required               Statetype = 32
	Statetype_InvalidEntry           Statetype = 33
	Statetype_SupportsAutocompletion Statetype = 34
	Statetype_SelectableText         Statetype = 35
	Statetype_Default                Statetype = 36
	Statetype_Animated               Statetype = 37
	Statetype_Visited                Statetype = 38
	Statetype_Checkable              Statetype = 39
	Statetype_HasPopup               Statetype = 40
	Statetype_HasTooltip             Statetype = 41
	Statetype_ReadOnly               Statetype = 42
	Statetype_LastDefined            Statetype = 43
)

// Textattribute is a representation of the C type TextAttribute.
type Textattribute int

const (
	Textattribute_Invalid          Textattribute = 0
	Textattribute_LeftMargin       Textattribute = 1
	Textattribute_RightMargin      Textattribute = 2
	Textattribute_Indent           Textattribute = 3
	Textattribute_Invisible        Textattribute = 4
	Textattribute_Editable         Textattribute = 5
	Textattribute_PixelsAboveLines Textattribute = 6
	Textattribute_PixelsBelowLines Textattribute = 7
	Textattribute_PixelsInsideWrap Textattribute = 8
	Textattribute_BgFullHeight     Textattribute = 9
	Textattribute_Rise             Textattribute = 10
	Textattribute_Underline        Textattribute = 11
	Textattribute_Strikethrough    Textattribute = 12
	Textattribute_Size             Textattribute = 13
	Textattribute_Scale            Textattribute = 14
	Textattribute_Weight           Textattribute = 15
	Textattribute_Language         Textattribute = 16
	Textattribute_FamilyName       Textattribute = 17
	Textattribute_BgColor          Textattribute = 18
	Textattribute_FgColor          Textattribute = 19
	Textattribute_BgStipple        Textattribute = 20
	Textattribute_FgStipple        Textattribute = 21
	Textattribute_WrapMode         Textattribute = 22
	Textattribute_Direction        Textattribute = 23
	Textattribute_Justification    Textattribute = 24
	Textattribute_Stretch          Textattribute = 25
	Textattribute_Variant          Textattribute = 26
	Textattribute_Style            Textattribute = 27
	Textattribute_LastDefined      Textattribute = 28
)

// Textboundary is a representation of the C type TextBoundary.
type Textboundary int

const (
	Textboundary_Char          Textboundary = 0
	Textboundary_WordStart     Textboundary = 1
	Textboundary_WordEnd       Textboundary = 2
	Textboundary_SentenceStart Textboundary = 3
	Textboundary_SentenceEnd   Textboundary = 4
	Textboundary_LineStart     Textboundary = 5
	Textboundary_LineEnd       Textboundary = 6
)

// Textcliptype is a representation of the C type TextClipType.
type Textcliptype int

const (
	Textcliptype_None Textcliptype = 0
	Textcliptype_Min  Textcliptype = 1
	Textcliptype_Max  Textcliptype = 2
	Textcliptype_Both Textcliptype = 3
)

// Textgranularity is a representation of the C type TextGranularity.
type Textgranularity int

const (
	Textgranularity_Char      Textgranularity = 0
	Textgranularity_Word      Textgranularity = 1
	Textgranularity_Sentence  Textgranularity = 2
	Textgranularity_Line      Textgranularity = 3
	Textgranularity_Paragraph Textgranularity = 4
)

// Valuetype is a representation of the C type ValueType.
type Valuetype int

const (
	Valuetype_VeryWeak    Valuetype = 0
	Valuetype_Weak        Valuetype = 1
	Valuetype_Acceptable  Valuetype = 2
	Valuetype_Strong      Valuetype = 3
	Valuetype_VeryStrong  Valuetype = 4
	Valuetype_VeryLow     Valuetype = 5
	Valuetype_Low         Valuetype = 6
	Valuetype_Medium      Valuetype = 7
	Valuetype_High        Valuetype = 8
	Valuetype_VeryHigh    Valuetype = 9
	Valuetype_VeryBad     Valuetype = 10
	Valuetype_Bad         Valuetype = 11
	Valuetype_Good        Valuetype = 12
	Valuetype_VeryGood    Valuetype = 13
	Valuetype_Best        Valuetype = 14
	Valuetype_LastDefined Valuetype = 15
)
