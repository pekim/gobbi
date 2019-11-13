// Code generated - DO NOT EDIT.

package javascriptcore

// Checksyntaxmode is a representation of the C type JSCCheckSyntaxMode.
type Checksyntaxmode int

const (
	// Checksyntaxmode_Script is a representation of the C type JSC_CHECK_SYNTAX_MODE_SCRIPT.
	Checksyntaxmode_Script Checksyntaxmode = 0
	// Checksyntaxmode_Module is a representation of the C type JSC_CHECK_SYNTAX_MODE_MODULE.
	Checksyntaxmode_Module Checksyntaxmode = 1
)

// Checksyntaxresult is a representation of the C type JSCCheckSyntaxResult.
type Checksyntaxresult int

const (
	// Checksyntaxresult_Success is a representation of the C type JSC_CHECK_SYNTAX_RESULT_SUCCESS.
	Checksyntaxresult_Success Checksyntaxresult = 0
	// Checksyntaxresult_RecoverableError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_RECOVERABLE_ERROR.
	Checksyntaxresult_RecoverableError Checksyntaxresult = 1
	// Checksyntaxresult_IrrecoverableError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_IRRECOVERABLE_ERROR.
	Checksyntaxresult_IrrecoverableError Checksyntaxresult = 2
	// Checksyntaxresult_UnterminatedLiteralError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_UNTERMINATED_LITERAL_ERROR.
	Checksyntaxresult_UnterminatedLiteralError Checksyntaxresult = 3
	// Checksyntaxresult_OutOfMemoryError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_OUT_OF_MEMORY_ERROR.
	Checksyntaxresult_OutOfMemoryError Checksyntaxresult = 4
	// Checksyntaxresult_StackOverflowError is a representation of the C type JSC_CHECK_SYNTAX_RESULT_STACK_OVERFLOW_ERROR.
	Checksyntaxresult_StackOverflowError Checksyntaxresult = 5
)

// Optiontype is a representation of the C type JSCOptionType.
//
// since 2.24
type Optiontype int

const (
	// Optiontype_Boolean is a representation of the C type JSC_OPTION_BOOLEAN.
	Optiontype_Boolean Optiontype = 0
	// Optiontype_Int is a representation of the C type JSC_OPTION_INT.
	Optiontype_Int Optiontype = 1
	// Optiontype_Uint is a representation of the C type JSC_OPTION_UINT.
	Optiontype_Uint Optiontype = 2
	// Optiontype_Size is a representation of the C type JSC_OPTION_SIZE.
	Optiontype_Size Optiontype = 3
	// Optiontype_Double is a representation of the C type JSC_OPTION_DOUBLE.
	Optiontype_Double Optiontype = 4
	// Optiontype_String is a representation of the C type JSC_OPTION_STRING.
	Optiontype_String Optiontype = 5
	// Optiontype_RangeString is a representation of the C type JSC_OPTION_RANGE_STRING.
	Optiontype_RangeString Optiontype = 6
)
