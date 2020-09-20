// Package data contains the data model of this app
package data

import "time"

// StravaUserToken is data from strava_user_token table
type StravaUserToken struct {
	UserID         string    `json:"userId"`
	StravaUserID   int64     `json:"stravaUserId"`
	AppID          int64     `json:"appId"`
	Token          string    `json:"token"`
	RefreshToken   string    `json:"refreshToken"`
	Active         bool      `json:"active"`
	CreatedDate    time.Time `json:"createdDate"`
	UpdatedDate    time.Time `json:"updatedDate"`
	TokenExpiredAt time.Time `json:"tokenExpiredAt"`
	TokenExpiredIn int64     `json:"tokenExpiredIn"`
}

// ActivityQueue is data from activity_queue table
type ActivityQueue struct {
	ActivityQueueID int64     `json:"activityQueueId"`
	StravaUserID    int64     `json:"stravaUserId"`
	Active          bool      `json:"active"`
	CreatedDate     time.Time `json:"createdDate"`
	UpdatedDate     time.Time `json:"updatedDate"`
}

// EventProgress is data from event_progress table
type EventProgress struct {
	EventProgressID  int64     `json:"eventProgressId"`
	EventID          int64     `json:"eventId"`
	UserID           string    `json:"userId"`
	Active           bool      `json:"active"`
	Distance         float64   `json:"distance"`
	ElapsedTime      float64   `json:"elapsedTime"`
	ValidDistance    float64   `json:"validDistance"`
	ValidElapsedTime float64   `json:"validElapsedTime"`
	CreatedDate      time.Time `json:"createdDate"`
	UpdatedDate      time.Time `json:"updatedDate"`
}

// StravaActivityHistory is data from strava_activity_history table
type StravaActivityHistory struct {
	UserID      string    `json:"userId"`
	History     string    `json:"history"`
	HasHistory  bool      `json:"hasHistory"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// StravaApp is data from strava_app table
type StravaApp struct {
	AppID           int64     `json:"appId"`
	ClientID        int64     `json:"clientId"`
	ClientSecret    string    `json:"clientSecret"`
	RefreshURL      string    `json:"refreshUrl"`
	AppName         string    `json:"appName"`
	Remark          string    `json:"remark"`
	LoadActivityURL string    `json:"loadActivityUrl"`
	CreatedDate     time.Time `json:"createdDate"`
	UpdatedDate     time.Time `json:"updatedDate"`
}

// MyEvent is data from event_register table
type MyEvent struct {
	EventRegisterID int64  `json:"eventRegisterId"`
	EventID         int64  `json:"eventId"`
	UserID          string `json:"userId"`
}

// RunningEvent is data from running_event table
type RunningEvent struct {
	EventID int64  `json:"eventId"`
	Name    string `json:"name"`
	Data    string `json:"data"`
}
