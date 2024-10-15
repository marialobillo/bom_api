package repository

type RepositoryError struct {
	Message string
	Err     error
}

func (e *RepositoryError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func NewRepositoryError(message string, err error) *RepositoryError {
	return &RepositoryError{
		Message: message,
		Err:     err,
	}
}
