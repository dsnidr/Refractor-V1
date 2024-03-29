/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package refractor

import (
	"errors"
	"github.com/sniddunc/refractor/pkg/config"
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

// Helper equals function for comparing ServiceResponses during testing.
// It compares the following fields: Success, StatusCode and Message.
func (sr *ServiceResponse) Equals(res *ServiceResponse) bool {
	if sr.Success != res.Success {
		return false
	}

	if sr.StatusCode != res.StatusCode {
		return false
	}

	if sr.Message != res.Message {
		return false
	}

	return true
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
