package rest_errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}
