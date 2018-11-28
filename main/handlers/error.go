package handlers

type Error struct {
	Message string `json:"error"`
}

func NewError(message string) *Error {
	return &Error{Message:message}
}

func (e *Error) Error() string {
	return e.Message
}