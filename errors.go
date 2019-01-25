package wtf

type AuthenticationError struct {
	err string
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{err: message}
}

func (e *AuthenticationError) Error() string {
	return e.err
}

type DialNotFoundError struct {
	err string
}

func NewDialNotFoundError(message string) *DialNotFoundError {
	return &DialNotFoundError{err: message}
}

func (e *DialNotFoundError) Error() string {
	return e.err
}
