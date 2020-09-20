// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadEventProgress is function to query from event_progres table
func LoadEventProgress(userID string) (listEventProg []data.EventProgress, err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		rows, err1 := conn.Query(`SELECT event_progress_id, event_id, user_id, active, distance, elapsed_time, valid_distance, valid_elapsed_time FROM event_progress WHERE user_id=$1`, userID)
		if err1 != nil {
			return listEventProg, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		listEventProg = make([]data.EventProgress, 0, 5)

		// Find result
		for rows.Next() {
			eventProg := data.EventProgress{}
			err = rows.Scan(&eventProg.EventProgressID, &eventProg.EventID, &eventProg.UserID, &eventProg.Active, &eventProg.Distance, &eventProg.ElapsedTime, &eventProg.ValidDistance, &eventProg.ValidElapsedTime)
			if err == nil {
				listEventProg = append(listEventProg, eventProg)
			} else {
				break
			}
		}
	}

	return listEventProg, err
}

// UpdateEventProgress is function to update event_progres table
func UpdateEventProgress(evtProg data.EventProgress) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Update active = false in strava_user_token
		sqlUpdate := "UPDATE event_progress SET distance=$1, elapsed_time=$2, valid_distance=$3, valid_elapsed_time=$4, updated_date=Now() WHERE event_progress_id=$5"

		// Do execute
		_, err = conn.Exec(sqlUpdate, evtProg.Distance, evtProg.ElapsedTime, evtProg.ValidDistance, evtProg.ValidElapsedTime, evtProg.EventProgressID)

	}

	return
}
