// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"log"

	"github.com/adirak/app-385-core/data"

	"github.com/adirak/app-385-core/dbconn"
)

// InsertRegisterLog to insert data into Register Log
func InsertRegisterLog(userID string, eventID int64) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolLogDB()

	if err == nil {

		// Insert Register Log
		sqlInsert := "INSERT INTO register_log(user_id, event_id) VALUES ($1, $2)"

		// Do execute
		_, err = conn.Exec(sqlInsert, userID, eventID)
	}

	return err
}

// LoadRegisterLog to query Register Log by userId and eventId
func LoadRegisterLog(userID string, eventID int64) (result data.RegisterLog, err error) {

	log.Println("userID=", userID, ", eventId=", eventID)

	// create connection
	conn, err := dbconn.GetConnectionPoolLogDB()
	if err == nil {

		// Queue SQL
		sqlQueue := `SELECT register_log_id, user_id, event_id, user_info, address, item, step, active FROM register_log WHERE user_id=$1 AND event_id=$2 AND active=TRUE`

		// Do execute
		rows, err2 := conn.Query(sqlQueue, userID, eventID)

		if err2 != nil {
			return result, err2
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&result.RegisterLogID, &result.UserID, &result.EventID, &result.UserInfo, &result.Address, &result.Item, &result.Step, &result.Active)
			break
		}
	}

	return result, err
}

// UpdateRegisterLog is function to update register_log table
func UpdateRegisterLog(rLog data.RegisterLog) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolLogDB()

	if err == nil {

		// Update SQL
		sqlUpdate := "UPDATE register_log SET user_info=$1, address=$2, item=$3, step=$4, updated_date=Now() WHERE register_log_id=$5"

		// Do execute
		_, err = conn.Exec(sqlUpdate, rLog.UserInfo, rLog.Address, rLog.Item, rLog.Step, rLog.RegisterLogID)

	}

	return err
}

// SetActiveRegisterLog is function to set active of register_log table
func SetActiveRegisterLog(rLog data.RegisterLog) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolLogDB()

	if err == nil {

		// Update SQL
		sqlUpdate := "UPDATE register_log SET active=$1, updated_date=Now() WHERE register_log_id=$2"

		// Do execute
		_, err = conn.Exec(sqlUpdate, rLog.Active, rLog.RegisterLogID)

	}

	return err
}

// SetStepRegisterLog is function to set step of register_log table
func SetStepRegisterLog(rLog data.RegisterLog) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolLogDB()

	if err == nil {

		// Update SQL
		sqlUpdate := "UPDATE register_log SET step=$1, updated_date=Now() WHERE register_log_id=$2"

		// Do execute
		_, err = conn.Exec(sqlUpdate, rLog.Step, rLog.RegisterLogID)

	}

	return err
}
