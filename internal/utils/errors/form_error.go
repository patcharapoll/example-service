package errors

type fieldError struct {
	FieldName   string
	Description string
}

// FromError ...
type FromError struct {
	Message string
	Fields  []*fieldError
}

func (e *FromError) Error() string {
	return e.Message
}

// AddErrorField ...
func (e *FromError) AddErrorField(fieldName string, description string) {
	e.Fields = append(e.Fields, &fieldError{
		FieldName:   fieldName,
		Description: description,
	})
}

// NewFormError ...
func NewFormError(message string) *FromError {
	err := &FromError{
		Message: message,
	}
	return err
}

// NewFormErrorWithFields ...
func NewFormErrorWithFields(message string, fields []*fieldError) *FromError {
	err := NewFormError(message)
	err.Fields = fields
	return err
}
