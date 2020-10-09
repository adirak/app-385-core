// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadRuleFunction is function to query rule_function table
func LoadRuleFunction(ruleID int64) (result data.RuleFunction, err error) {

	// init
	result = data.RuleFunction{}

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()

	if err == nil {

		// Queue SQL
		sqlQueue := `SELECT rule_id, function_name, drescription FROM rule_function WHERE rule_id = $1`

		// Do execute
		rows, err2 := conn.Query(sqlQueue, ruleID)

		if err2 != nil {
			return result, err2
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&result.RuleID, &result.FunctionName, &result.Description)
			break
		}
	}

	return result, err
}
