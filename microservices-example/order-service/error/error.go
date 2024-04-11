package serviceerror

type invalidRequestError struct{
	m string
}

func (e invalidRequestError) Error() string {
	return e.m
}

func InvalidRequestError(message string) error {
	return invalidRequestError{m: message}
}