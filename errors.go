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

type ErrDialIdNotFound struct {
	err string
}

func NewErrDialIdNotFound(message string) *ErrDialIdNotFound {
	return &ErrDialIdNotFound{err: message}
}

func (e *ErrDialIdNotFound) Error() string {
	return e.err
}

type ErrDialAlreadyExists struct {
	err string
}

func NewErrDialAlreadyExists(message string) *ErrDialAlreadyExists {
	return &ErrDialAlreadyExists{err: message}
}

func (e *ErrDialAlreadyExists) Error() string {
	return e.err
}
