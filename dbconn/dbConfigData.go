// Package dbconn contains an database connection
package dbconn

import (
	"github.com/adirak/app-385-core/data"
)

// GetDBConfig is function to get the database config
func GetDBConfig() data.DBConfig {

	// Database Config
	dbConfig := data.DBConfig{}
	dbConfig.User = "postgres"
	dbConfig.Pwd = "TunTongTee@385app"
	dbConfig.TCPHost = "34.101.126.60"
	dbConfig.InstanceConnName = "app-385:asia-southeast2:app-385-db"
	dbConfig.SocketDir = "/cloudsql"
	dbConfig.Port = "5432"
	dbConfig.DBNameForApp = "app_385_activity"
	dbConfig.DBNameForLog = "app_log_db"

	dbConfig.InitSuccess = false

	return dbConfig
}
