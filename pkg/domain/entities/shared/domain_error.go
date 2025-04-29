package shared

type DomainError struct {
	Field       string
	Cause       string
	Message     string
	Description string
	Code        int
}

func NewDomainError(field, cause, message, description string, code int) *DomainError {
	return &DomainError{
		Field:       field,
		Cause:       cause,
		Message:     message,
		Description: description,
		Code:        code,
	}
}

func (e *DomainError) Error() string {
	return e.Message
}
