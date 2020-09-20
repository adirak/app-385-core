// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadStravaUserToken is function to query from strava_user_token
func LoadStravaUserToken(stravaUserID int64) (userToken data.StravaUserToken, err error) {

	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		rows, err1 := conn.Query(`SELECT user_id, strava_user_id, app_id, token, refresh_tokken, active, token_expired_at, token_expired_in FROM strava_user_token WHERE strava_user_id=$1`, stravaUserID)
		if err1 != nil {
			return userToken, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		// Find result
		for rows.Next() {
			err = rows.Scan(&userToken.UserID, &userToken.StravaUserID, &userToken.AppID, &userToken.Token, &userToken.RefreshToken, &userToken.Active, &userToken.TokenExpiredAt, &userToken.TokenExpiredIn)
		}

	}

	return userToken, err
}

// UpdateStravaUserToken is function to update refresh token expired for app, and push notification to user
func UpdateStravaUserToken(token data.StravaUserToken) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Update active = false in strava_user_token
		sqlUpdate := "UPDATE strava_user_token SET token=$1, refresh_tokken=$2, active=$3, token_expired_at=$4, token_expired_in=$5, updated_date=Now() WHERE user_id=$6"

		// Do execute
		_, err = conn.Exec(sqlUpdate, token.Token, token.RefreshToken, token.Active, token.TokenExpiredAt, token.TokenExpiredIn, token.UserID)

		if err != nil {
			err = fmt.Errorf("DB.Exec: %v", err)
		}
	}

	return
}

// UpdateTokenExpired is function to set active flag=false from strava_user_token
func UpdateTokenExpired(token data.StravaUserToken) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Update active = false in strava_user_token
		sqlUpdate := "UPDATE strava_user_token SET active=FALSE, updated_date=Now() WHERE user_id=$1"

		// Do execute
		_, err = conn.Exec(sqlUpdate, token.UserID)

	}

	return err
}
