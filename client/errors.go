package client

import (
	"fmt"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("3X-UI API error %d: %s", e.StatusCode, e.Message)
}
