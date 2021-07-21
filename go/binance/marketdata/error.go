package marketdata

// Error with code and message
type Error struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

// NewError
func NewError() *Error {
	return &Error{
		Code:    "",
		Message: "",
	}
}
