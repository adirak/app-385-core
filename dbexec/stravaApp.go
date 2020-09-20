// Package dbexec contains an database execution such as query insert or update table
package dbexec

import (
	"fmt"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbconn"
)

// LoadStravaApp is function to query from strava_app table
func LoadStravaApp(appID int64) (stravaApp data.StravaApp, err error) {

	// Create connection
	conn, err := dbconn.GetConnectionPoolAppDB()
	if err == nil {

		rows, err1 := conn.Query(`SELECT app_id, client_id, client_secret, refresh_url, load_activity_url, app_name, remark FROM strava_app WHERE app_id=$1`, appID)
		if err1 != nil {
			return stravaApp, fmt.Errorf("DB.Query: %v", err1)
		}
		defer rows.Close()

		// Find result
		for rows.Next() {
			err = rows.Scan(&stravaApp.AppID, &stravaApp.ClientID, &stravaApp.ClientSecret, &stravaApp.RefreshURL, &stravaApp.LoadActivityURL, &stravaApp.AppName, &stravaApp.Remark)
		}

	}

	return stravaApp, err
}
