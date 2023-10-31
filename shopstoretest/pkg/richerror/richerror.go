package richerror

type Kind int

const (
	KindInvalid = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation    string
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func New(err error, operation, message string, kind Kind) RichError {
	newRichError := RichError{
		operation:    operation,
		wrappedError: err,
		message:      message,
		kind:         kind,
	}

	return newRichError
}
