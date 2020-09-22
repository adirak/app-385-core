// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadMyEvent is function to query event_register table
func LoadMyEvent(userID string) ([]data.MyEvent, error) {

	// Create connection
	conn, err1 := dbconn.GetConnectionPoolAppDB()
	list := make([]data.MyEvent, 0, 5)

	if err1 == nil {

		sqlQuery := `SELECT event_register_id, event_id, user_id FROM event_register WHERE user_id=$1`

		rows, err := conn.Query(sqlQuery, userID)
		if err != nil {
			return nil, fmt.Errorf("DB.Query: %v", err)
		}
		defer rows.Close()

		for rows.Next() {

			myEvent := data.MyEvent{}

			err = rows.Scan(&myEvent.EventRegisterID, &myEvent.EventID, &myEvent.UserID)
			if err == nil {
				list = append(list, myEvent)
			} else {
				log.Fatal(err)
			}

		}

		err1 = err

	}

	return list, err1
}
