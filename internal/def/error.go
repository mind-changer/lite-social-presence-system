package def

type ClientError struct {
	Code    int
	Message string
}

func (e *ClientError) Error() string {
	return e.Message
}

func (e *ClientError) GetCode() int {
	return e.Code
}

func CreateClientError(code int, msg string) *ClientError {
	return &ClientError{Code: code, Message: msg}
}
