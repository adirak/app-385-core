// Package dbconn contains an database connection
package dbconn

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adirak/app-385-core/data"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Config of this database app
var dbAppConf data.DBConfig
var dbLogConf data.DBConfig

// Initial data for one time
func initAppDB() {

	// Init config data
	dbConf := GetDBConfig()
	dbConf.IsAppDB = true

	dbPoolApp, err := initSocketConnectionPool(dbConf)
	if err != nil {
		log.Println("initial dbConnApp fail : ", err.Error())
	} else {
		log.Println("initial dbConnApp successful")
		dbConf.Conn = dbPoolApp
		dbConf.InitSuccess = true
	}

	dbConf.Err = err

	// Keep instance
	dbAppConf = dbConf
}

// Initial data for one time
func initLogDB() {

	// Init config data
	dbConf := GetDBConfig()
	dbConf.IsAppDB = false

	dbPoolApp, err := initSocketConnectionPool(dbConf)
	if err != nil {
		log.Println("initial dbConnLog fail : ", err.Error())
	} else {
		log.Println("initial dbConnLog successful")
		dbConf.Conn = dbPoolApp
		dbConf.InitSuccess = true
	}

	dbConf.Err = err

	// Keep instance
	dbLogConf = dbConf
}

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool(dbConf data.DBConfig) (*sql.DB, error) {

	var dbURI string

	// App Database
	if dbConf.IsAppDB {
		dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbConf.User, dbConf.Pwd, dbConf.DBNameForApp, dbConf.SocketDir, dbConf.InstanceConnName)
	} else {
		dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbConf.User, dbConf.Pwd, dbConf.DBNameForLog, dbConf.SocketDir, dbConf.InstanceConnName)
	}

	log.Println("***dbURL : ", dbURI)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_postgres_databasesql_create_socket]
}

// ConfigureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_postgres_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(2)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(3)

	// [END cloud_sql_postgres_databasesql_limit]

	// [START cloud_sql_postgres_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_postgres_databasesql_lifetime]
}

// GetConnectionPoolAppDB is function to get connection pool of app database
func GetConnectionPoolAppDB() (*sql.DB, error) {

	if dbAppConf.InitSuccess == false {
		initAppDB()
	}

	return dbAppConf.Conn, dbAppConf.Err
}

// GetConnectionPoolLogDB is function to get connection pool of log database
func GetConnectionPoolLogDB() (*sql.DB, error) {

	if dbLogConf.InitSuccess == false {
		initLogDB()
	}

	return dbLogConf.Conn, dbLogConf.Err
}
