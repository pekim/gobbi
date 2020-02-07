package ffi

import (
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

func tracing() bool {
	return traceToStdout || traceHandler != nil
}

func trace(message string) {
	if traceToStdout {
		_, err := os.Stdout.WriteString(message)
		if err != nil {
			panic(err)
		}

		return
	}

	if traceHandler == nil {
		return
	}
	traceHandler(message)
}
