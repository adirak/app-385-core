// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"log"

	"github.com/adirak/app-385-core/dbconn"
)

// InsertWebhookLog to insert data into Webhook log
func InsertWebhookLog(actionType string, jsonStr string) (err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolLogDB()

	if err == nil {

		// Insert webhook log to database
		sqlInsert := "INSERT INTO webhook_log(action_type, log_detail) VALUES($1, $2)"

		// Do execute
		_, err = conn.Exec(sqlInsert, actionType, jsonStr)
	}

	return err
}

// LoadWebhookLog100Rows to query latest of webhook log 100 records
func LoadWebhookLog100Rows() ([]map[string]interface{}, error) {

	list := make([]map[string]interface{}, 0, 5)

	// create connection
	conn, err := dbconn.GetConnectionPoolLogDB()
	if err == nil {

		rows, err := conn.Query(`SELECT * FROM webhook_log ORDER BY webhook_log_id limit 100`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var webhookLogID int
			var actionType string
			var detail string
			var createdDate string
			err := rows.Scan(&webhookLogID, &actionType, &detail, &createdDate)
			if err == nil {
				row := make(map[string]interface{})
				row["webhookLogId"] = webhookLogID
				row["actionType"] = actionType
				row["detail"] = detail
				row["createdDate"] = createdDate
				list = append(list, row)
			} else {
				log.Fatal(err)
			}

		}
	}

	return list, err
}
