package cgo

type TraceHandler func(message string)

var traceHandler TraceHandler

func SetTraceHandler(handler TraceHandler) {
	traceHandler = handler
}

func Tracing() bool {
	return traceHandler != nil
}

func Trace(message string) {
	if traceHandler == nil {
		return
	}

	traceHandler(message)
}
