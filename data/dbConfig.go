// Package data contains the data model of this app
package data

import "database/sql"

// DBConfig is struct for config
type DBConfig struct {

	// Data to connect to database
	User             string
	Pwd              string
	TCPHost          string
	InstanceConnName string
	SocketDir        string
	Port             string
	DBNameForLog     string
	DBNameForApp     string

	// Database connection pool
	Conn        *sql.DB
	Err         error
	IsAppDB     bool
	InitSuccess bool
}
