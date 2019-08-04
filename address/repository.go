package address

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqlDB struct {
	database *sql.DB
}

type QueryInterface interface {
	Query(queryStr string, args ...interface{}) (*sql.Rows, error)
	QueryRow(queryStr string, args ...interface{}) *sql.Row
}

const sqlVersion = "sqlite3"

func ConnectSqlDB(path string) (SqlDB, error) {
	var sqlDB SqlDB
	var err error
	sqlDB.database, err = sql.Open(sqlVersion, path)
	return sqlDB, err
}

/*
func Init() error {
	var err error
	database, err = sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	return database.Ping()
}
*/

func (sqlDB *SqlDB) Query(queryStr string, args ...interface{}) (*sql.Rows, error) {
	if len(args) == 0 {
		return sqlDB.database.Query(queryStr)
	} else {
		return sqlDB.database.Query(queryStr, args...)
	}
}

func (sqlDB *SqlDB) QueryRow(queryStr string, args ...interface{}) *sql.Row {
	if len(args) == 0 {
		return sqlDB.database.QueryRow(queryStr)
	} else {
		return sqlDB.database.QueryRow(queryStr, args...)
	}
}
