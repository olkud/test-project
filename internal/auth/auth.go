package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey Extract API Key from header
// Return error if nothing found
// Example: Authentication: ApiKey {insert API Key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authentication")

	if val == "" {
		return "", errors.New("no auth header found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("mailformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("mailformed first part of the header")
	}

	return vals[1], nil
}
