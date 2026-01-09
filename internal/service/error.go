package service

type ConflictError struct {
	message string
}

func (e ConflictError) Error() string {
	return e.message
}

func NewServiceError(message string) ConflictError {
	return ConflictError{message: message}
}

type DataError struct {
	message string
}

func (e DataError) Error() string {
	return e.message
}

func NewDataError(message string) DataError {
	return DataError{message: message}
}
