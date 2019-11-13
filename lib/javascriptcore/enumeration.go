// Code generated - DO NOT EDIT.

package javascriptcore

// Checksyntaxmode is a representation of the C type CheckSyntaxMode.
type Checksyntaxmode int

const (
	// script
	JscCheckSyntaxModeScript Checksyntaxmode = 0
	// module
	JscCheckSyntaxModeModule Checksyntaxmode = 1
)

// Checksyntaxresult is a representation of the C type CheckSyntaxResult.
type Checksyntaxresult int

const (
	// success
	JscCheckSyntaxResultSuccess Checksyntaxresult = 0
	// recoverable_error
	JscCheckSyntaxResultRecoverableError Checksyntaxresult = 1
	// irrecoverable_error
	JscCheckSyntaxResultIrrecoverableError Checksyntaxresult = 2
	// unterminated_literal_error
	JscCheckSyntaxResultUnterminatedLiteralError Checksyntaxresult = 3
	// out_of_memory_error
	JscCheckSyntaxResultOutOfMemoryError Checksyntaxresult = 4
	// stack_overflow_error
	JscCheckSyntaxResultStackOverflowError Checksyntaxresult = 5
)

// Optiontype is a representation of the C type OptionType.
//
// since 2.24
type Optiontype int

const (
	// boolean
	JscOptionBoolean Optiontype = 0
	// int
	JscOptionInt Optiontype = 1
	// uint
	JscOptionUint Optiontype = 2
	// size
	JscOptionSize Optiontype = 3
	// double
	JscOptionDouble Optiontype = 4
	// string
	JscOptionString Optiontype = 5
	// range_string
	JscOptionRangeString Optiontype = 6
)
