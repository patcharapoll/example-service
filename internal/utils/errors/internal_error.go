package errors

import "fmt"

// InternalErrorCodeInterface ...
type InternalErrorCodeInterface interface {
	String() string
}

// InternalError ...
type InternalError struct {
	Code     InternalErrorCodeInterface
	Message  string
	Metadata map[string]string
}

func (e *InternalError) Error() string {
	return fmt.Sprintln("%s: %s", e.Code.String(), e.Message)
}

// AddMetaData ...
func (e *InternalError) AddMetaData(key string, value string) {
	e.Metadata[key] = value
}

// ErrorWithMetadata ...
func (e *InternalError) ErrorWithMetadata() string {
	errMsg := e.Error()

	for k, v := range e.Metadata {
		errMsg += fmt.Sprintf("  %s=%s", k, v)
	}

	return errMsg
}

// NewInternalError ...
func NewInternalError(code InternalErrorCodeInterface, message string) *InternalError {
	err := &InternalError{
		Code:     code,
		Message:  message,
		Metadata: map[string]string{},
	}
	return err
}

// NewInternalErrorWithMetadata ...
func NewInternalErrorWithMetadata(code InternalErrorCodeInterface, message string, meta map[string]string) *InternalError {
	err := &InternalError{
		Code:     code,
		Message:  message,
		Metadata: meta,
	}
	return err
}
