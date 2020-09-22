// Package data contains the data model of this app
package data

// ReqtData is data from http request
type ReqtData struct {
	InData map[string]interface{}
}

// RespData is data will response to frontend
type RespData struct {
	Code              int
	Msg               string
	Err               error
	Success           bool
	OutData           map[string]interface{}
	ForceResponseData interface{}
}
