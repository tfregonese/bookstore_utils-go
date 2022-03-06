package rest_errors

import "net/http"

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Message: message,
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Message: message,
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
