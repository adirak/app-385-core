// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadRunningEvents is function to query runnint_event
func LoadRunningEvents() ([]data.RunningEvent, error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	// Make result
	events := make([]data.RunningEvent, 0, 5)

	if err == nil {

		sqlQuery := `SELECT event_id, name, data, rule_id FROM running_event WHERE active=TRUE`

		rows, err1 := conn.Query(sqlQuery)
		if err1 != nil {
			return nil, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		for rows.Next() {

			runEvent := data.RunningEvent{}

			err := rows.Scan(&runEvent.EventID, &runEvent.Name, &runEvent.Data, &runEvent.RuleID)
			if err == nil {
				events = append(events, runEvent)
			} else {
				log.Fatal(err)
			}

		}
	}

	return events, err
}

// LoadRunningEvent is function to query running_event table
func LoadRunningEvent(myEvent data.MyEvent) (data.RunningEvent, error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	runEvent := data.RunningEvent{}

	if err == nil {

		sqlQuery := `SELECT event_id, name, data, rule_id FROM running_event WHERE event_id=$1`

		rows, err1 := conn.Query(sqlQuery, myEvent.EventID)
		if err1 != nil {
			return runEvent, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(&runEvent.EventID, &runEvent.Name, &runEvent.Data, &runEvent.RuleID)
			if err == nil {
				break
			} else {
				log.Fatal(err)
			}

		}
	}

	return runEvent, err
}

// LoadRunningEventByID is function to load runnint event by ID
func LoadRunningEventByID(eventID int64) (data.RunningEvent, error) {
	myEvent := data.MyEvent{}
	myEvent.EventID = eventID
	return LoadRunningEvent(myEvent)
}

// CreateRunningEvent is function to create running event
func CreateRunningEvent(name string, data interface{}, active bool, ruleID int64) error {

	// Convert map to json
	json, err1 := json.Marshal(data)

	if err1 == nil {

		// Create connection
		conn, err := dbconn.GetConnectionPoolAppDB()

		if err == nil {

			sqlInsert := "INSERT INTO running_event(name, data, active, rule_id) VALUES ($1, $2, $3, $4)"

			// Do execute
			_, err = conn.Exec(sqlInsert, name, json, active, ruleID)

			err1 = err

		}

	}

	return err1
}
