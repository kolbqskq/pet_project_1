package errs

import "net/http"

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewWrongMinerClass() error {
	return &HTTPError{Code: http.StatusBadRequest, Message: "wrong miner class request"}
}

func NewNotEnoughBalance() error {
	return &HTTPError{Code: http.StatusConflict, Message: "not enough balance"}
}

func NewErrRequestBody() error {
	return &HTTPError{Code: http.StatusBadRequest, Message: "invalid or missing request body"}
}

func NewCustomErr(code int, message string) error {
	return &HTTPError{Code: code, Message: message}
}

func NewHaveNotMiners() error {
	return &HTTPError{Code: 409, Message: "you have not miners"}
}

func NewNotAllTaskCompleted() error {
	return &HTTPError{Code: 409, Message: "Not all tasks are completed"}
}

func NewInvalidClass() error {
	return &HTTPError{Code: 400, Message: "invalid class"}
}

func NewInvalidEquipment() error {
	return &HTTPError{Code: 400, Message: "invalid equipment"}
}

func NewEquipmentAlreadyOwn() error {
	return &HTTPError{Code: 409, Message: "equipment already own"}
}

func NewYouShouldBuyAllTools() error {
	return &HTTPError{Code: 409, Message: "you should buy all tools"}
}
