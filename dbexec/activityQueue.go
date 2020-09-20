// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadActivityQueue is function to query from qctivity_queue
func LoadActivityQueue() ([]data.ActivityQueue, error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()
	list := make([]data.ActivityQueue, 0, 5)
	if err == nil {

		rows, err1 := conn.Query(`SELECT activity_queue_id, strava_user_id, active FROM activity_queue WHERE active=TRUE`)
		if err1 != nil {
			return list, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		// Loop to get data
		for rows.Next() {

			actQueue := data.ActivityQueue{}

			err := rows.Scan(&actQueue.ActivityQueueID, &actQueue.StravaUserID, &actQueue.Active)
			if err == nil {
				list = append(list, actQueue)
			} else {
				log.Fatal(err)
			}

		}

	}

	return list, err
}

// UpdateActivityQueue is function to update activity_queue table
func UpdateActivityQueue(actQueue data.ActivityQueue) (err error) {

	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Update active = false in activity_queue
		sqlUpdate := "UPDATE activity_queue SET active=$1, updated_date=Now() WHERE activity_queue_id=$2"

		// Do execute
		_, err = conn.Exec(sqlUpdate, actQueue.Active, actQueue.ActivityQueueID)

	}

	return err
}
