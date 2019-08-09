package repository

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
	Close()
}

const sqlVersion = "sqlite3"

func ConnectSqlDB(path string) (*SqlDB, error) {
	var sqlDB SqlDB
	var err error
	sqlDB.database, err = sql.Open(sqlVersion, path)
	return &sqlDB, err
}

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

func (sqlDB *SqlDB) Close() {
	defer sqlDB.database.Close()
}
