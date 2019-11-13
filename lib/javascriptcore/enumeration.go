// Code generated - DO NOT EDIT.

package javascriptcore

// Checksyntaxmode is a representation of the C type CheckSyntaxMode.
type Checksyntaxmode int

const (
	Checksyntaxmode_Script Checksyntaxmode = 0
	Checksyntaxmode_Module Checksyntaxmode = 1
)

// Checksyntaxresult is a representation of the C type CheckSyntaxResult.
type Checksyntaxresult int

const (
	Checksyntaxresult_Success                  Checksyntaxresult = 0
	Checksyntaxresult_RecoverableError         Checksyntaxresult = 1
	Checksyntaxresult_IrrecoverableError       Checksyntaxresult = 2
	Checksyntaxresult_UnterminatedLiteralError Checksyntaxresult = 3
	Checksyntaxresult_OutOfMemoryError         Checksyntaxresult = 4
	Checksyntaxresult_StackOverflowError       Checksyntaxresult = 5
)

// Optiontype is a representation of the C type OptionType.
//
// since 2.24
type Optiontype int

const (
	Optiontype_Boolean     Optiontype = 0
	Optiontype_Int         Optiontype = 1
	Optiontype_Uint        Optiontype = 2
	Optiontype_Size        Optiontype = 3
	Optiontype_Double      Optiontype = 4
	Optiontype_String      Optiontype = 5
	Optiontype_RangeString Optiontype = 6
)
