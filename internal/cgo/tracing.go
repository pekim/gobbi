package cgo

import (
	"fmt"
	"os"
	"strconv"
)

type TraceHandler func(message string)

var traceHandler TraceHandler

var traceToStdout bool

func init() {
	traceToStdout, _ = strconv.ParseBool(os.Getenv("GOBBI_TRACE"))
}

func SetTraceHandler(handler TraceHandler) {
	traceHandler = handler
}

func Tracing() bool {
	return traceToStdout || traceHandler != nil
}

func Trace(message string) {
	fmt.Println(traceToStdout, message)
	if traceToStdout {
		os.Stdout.WriteString(message)
		return
	}

	if traceHandler == nil {
		return
	}
	traceHandler(message)
}
