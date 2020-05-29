package jvvt

const (
	VaildationErrorDefective uint32 = 1 << iota		// If token is not in proper format
	ValidationErrorTokenExpired					
)


type ValidationError struct {
	ErrorCode    uint32
	ErrorMessage string
}

func NewValidationError(errorMessage string, errorFlags uint32) *ValidationError {
	return &ValidationError{
		ErrorCode:    errorFlags,
		ErrorMessage: errorMessage,
	}
}
