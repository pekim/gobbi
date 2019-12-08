package gi

type ErrorHandler func(err error)

var errorHandler ErrorHandler

func SetErrorHandler(handler ErrorHandler) {
	errorHandler = handler
}

func handleError(err error) {
	if errorHandler == nil {
		return
	}

	errorHandler(err)
}
