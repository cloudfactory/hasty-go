package hasty

import (
	"fmt"
	"strings"
)

var _ error = Error{}

// Error is used to represent API call error with details
type Error struct {
	Status    int    `json:"-"`
	RequestID string `json:"-"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// Error returns string implementation of error
func (e Error) Error() string {
	message := []string{
		fmt.Sprintf("API request %s failed with status %d", e.RequestID, e.Status),
	}
	if e.Code != "" {
		message = append(message, fmt.Sprintf("code %s", e.Code))
	}
	if e.Message != "" {
		message = append(message, fmt.Sprintf("message %s", e.Message))
	}
	return strings.Join(message, ", ")
}
