package config

type DBConnError struct {
	Message string
	Code    int
}

func (e *DBConnError) Error() string {
	return e.Message
}

type UserCreateError struct {
	Message string
	Code    int
}

func (e *UserCreateError) Error() string {
	return e.Message
}

type RequestValidationError struct {
	Message string
	Code    int
}

func (e *RequestValidationError) Error() string {
	return e.Message
}
