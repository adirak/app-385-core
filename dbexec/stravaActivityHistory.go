// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadStravaActivityHistory is function to query from strava_activity_history
func LoadStravaActivityHistory(userID string) (history data.StravaActivityHistory, err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		rows, err1 := conn.Query(`SELECT user_id, history FROM strava_activity_history WHERE user_id=$1`, userID)
		if err1 != nil {
			return history, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		// Find result
		for rows.Next() {
			err = rows.Scan(&history.UserID, &history.History)
			if err == nil {
				history.HasHistory = true
			}
		}

		if history.HasHistory == false {
			history.UserID = userID
		}

	}

	return history, err
}

// InsertStravaActivityHistory is function to insert history to strava_activity_history
func InsertStravaActivityHistory(history data.StravaActivityHistory) (err error) {

	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Insert webhook queue to database
		// If strava_user_id is constrain
		sqlInsert := "INSERT INTO strava_activity_history(user_id, history) VALUES ($1, $2)"

		// Do execute
		_, err = conn.Exec(sqlInsert, history.UserID, history.History)

	}

	return err
}

// UpdateStravaActivityHistory is function to update history to strava_activity_history
func UpdateStravaActivityHistory(history data.StravaActivityHistory) (err error) {

	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Update history of strava activity_id
		sqlUpdate := "UPDATE strava_activity_history SET history=$1, updated_date=Now() WHERE user_id=$2"

		// Do execute
		_, err = conn.Exec(sqlUpdate, history.History, history.UserID)

	}

	return err
}
