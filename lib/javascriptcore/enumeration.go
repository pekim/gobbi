// Code generated - DO NOT EDIT.

package javascriptcore

// CheckSyntaxMode is a representation of the C type JSCCheckSyntaxMode.
//
type CheckSyntaxMode int

const (
	// CheckSyntaxMode_Script is a representation of the C type JSC_CHECK_SYNTAX_MODE_SCRIPT.
	CheckSyntaxMode_Script CheckSyntaxMode = 0
	// CheckSyntaxMode_Module is a representation of the C type JSC_CHECK_SYNTAX_MODE_MODULE.
	CheckSyntaxMode_Module CheckSyntaxMode = 1
)

// CheckSyntaxResult is a representation of the C type JSCCheckSyntaxResult.
//
type CheckSyntaxResult int

const (
	// CheckSyntaxResult_Success is a representation of the C type JSC_CHECK_SYNTAX_RESULT_SUCCESS.
	CheckSyntaxResult_Success CheckSyntaxResult = 0
	// CheckSyntaxResult_RecoverableError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_RECOVERABLE_ERROR.
	CheckSyntaxResult_RecoverableError CheckSyntaxResult = 1
	// CheckSyntaxResult_IrrecoverableError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_IRRECOVERABLE_ERROR.
	CheckSyntaxResult_IrrecoverableError CheckSyntaxResult = 2
	// CheckSyntaxResult_UnterminatedLiteralError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_UNTERMINATED_LITERAL_ERROR.
	CheckSyntaxResult_UnterminatedLiteralError CheckSyntaxResult = 3
	// CheckSyntaxResult_OutOfMemoryError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_OUT_OF_MEMORY_ERROR.
	CheckSyntaxResult_OutOfMemoryError CheckSyntaxResult = 4
	// CheckSyntaxResult_StackOverflowError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_STACK_OVERFLOW_ERROR.
	CheckSyntaxResult_StackOverflowError CheckSyntaxResult = 5
)

// OptionType is a representation of the C type JSCOptionType.
//
// since 2.24
type OptionType int

const (
	// OptionType_Boolean is a representation of the C type JSC_OPTION_BOOLEAN.
	OptionType_Boolean OptionType = 0
	// OptionType_Int is a representation of the C type JSC_OPTION_INT.
	OptionType_Int OptionType = 1
	// OptionType_Uint is a representation of the C type JSC_OPTION_UINT.
	OptionType_Uint OptionType = 2
	// OptionType_Size is a representation of the C type JSC_OPTION_SIZE.
	OptionType_Size OptionType = 3
	// OptionType_Double is a representation of the C type JSC_OPTION_DOUBLE.
	OptionType_Double OptionType = 4
	// OptionType_String is a representation of the C type JSC_OPTION_STRING.
	OptionType_String OptionType = 5
	// OptionType_RangeString is a representation of the C type JSC_OPTION_RANGE_STRING.
	OptionType_RangeString OptionType = 6
)
