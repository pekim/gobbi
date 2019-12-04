package gi

type TraceHandler func(message string)

var traceHandler TraceHandler

func SetTraceHandler(handler TraceHandler) {
	traceHandler = handler
}

func tracing() bool {
	return traceHandler != nil
}

func trace(message string) {
	if traceHandler == nil {
		return
	}

	traceHandler(message)
}
