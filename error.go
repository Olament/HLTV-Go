package hltv

import "fmt"

type HTTPError struct {
	Code int
	Description string
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("status code %d: %s", h.Code, h.Description)
}