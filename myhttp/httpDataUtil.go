// Package myhttp contains an http connection
package myhttp

import (
	"strings"
)

// IsJSONMap is funtion to check json data whether it is map
func IsJSONMap(body []byte) bool {
	json := string(body)
	json = strings.TrimSpace(json)
	r := strings.HasPrefix(json, "{")
	return r
}

// IsJSONArray is funtion to check json data whether it is array
func IsJSONArray(body []byte) bool {
	json := string(body)
	json = strings.TrimSpace(json)
	r := strings.HasPrefix(json, "[")
	return r
}
