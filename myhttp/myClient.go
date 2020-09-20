// Package myhttp contains an http connection
package myhttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GetMyClient is function to get custom client
func GetMyClient() *http.Client {

	// Transport to skip verify ssl
	trnsp := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create client
	client := &http.Client{Transport: trnsp, Timeout: time.Second * 30}

	return client
}

// NewPostRequest is function to create Post request
func NewPostRequest(url string, formData url.Values) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(formData.Encode()))

	return req, err
}

// NewGetRequest is function to create Get request
func NewGetRequest(url string) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	return req, err
}

// SetTokenHeader is function to set Authorization token
func SetTokenHeader(req *http.Request, accessToken string) {
	req.Header.Set("Authorization", "Bearer "+accessToken)
}

// SetAcceptJSONHeader is function to set accept JSON header
func SetAcceptJSONHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
}
