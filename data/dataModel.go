// Package data contains the data model of this app
package data

import "time"

// RefreshTokenResult is result data after call strava refresh token
type RefreshTokenResult struct {
	UserToken    StravaUserToken
	AccessToken  string
	RefreshToken string
	TokenExpired bool
	ExpiredAt    time.Time
	ExpiredIn    int64
	Success      bool
}
