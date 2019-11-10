// Code generated - DO NOT EDIT.

package javascriptcore

type CheckSyntaxMode uint32

const (
	CheckSyntaxMode_Script CheckSyntaxMode = int64(0)
	CheckSyntaxMode_Module CheckSyntaxMode = int64(1)
)

type CheckSyntaxResult uint32

const (
	CheckSyntaxResult_Success                  CheckSyntaxResult = int64(0)
	CheckSyntaxResult_RecoverableError         CheckSyntaxResult = int64(1)
	CheckSyntaxResult_IrrecoverableError       CheckSyntaxResult = int64(2)
	CheckSyntaxResult_UnterminatedLiteralError CheckSyntaxResult = int64(3)
	CheckSyntaxResult_OutOfMemoryError         CheckSyntaxResult = int64(4)
	CheckSyntaxResult_StackOverflowError       CheckSyntaxResult = int64(5)
)

type OptionType uint32

const (
	OptionType_Boolean     OptionType = int64(0)
	OptionType_Int         OptionType = int64(1)
	OptionType_Uint        OptionType = int64(2)
	OptionType_Size        OptionType = int64(3)
	OptionType_Double      OptionType = int64(4)
	OptionType_String      OptionType = int64(5)
	OptionType_RangeString OptionType = int64(6)
)
