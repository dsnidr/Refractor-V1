package refractor

import (
	"backend/pkg/config"
	"errors"
	"net/http"
	"net/url"
)

// FindArgs is used to specify search criteria for querying a datastore
type FindArgs map[string]interface{}

// UpdateArgs is used to specify field updates for an object in a database
type UpdateArgs map[string]interface{}

// This struct is what is returned by services to communicate with their handlers
type ServiceResponse struct {
	Success          bool
	StatusCode       int
	Message          string
	Error            error
	ValidationErrors url.Values
}

var (
	// ErrNotFound is used when a record could not be found in storage
	ErrNotFound = errors.New("record not found")

	// ErrInternalError is used when something goes wrong on our end
	ErrInternalError = errors.New("something went wrong. Please try again later")

	// InternalErrorResponse stores a pointer to a ServiceResponse struct containing the response
	// fields used during an internal error service response. This exists because it is used by
	// many service implementations across Refractor, and rewriting it every time shouldn't be necessary.
	InternalErrorResponse = &ServiceResponse{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Message:    config.MessageInternalError,
	}
)
