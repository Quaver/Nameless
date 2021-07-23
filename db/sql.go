package db

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/config"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type RowScanner interface {
	Scan(dest ...interface{}) error
}

var SQL *sql.DB

// InitializeSQL Initializes the SQL database connection
func InitializeSQL() {
	if SQL != nil {
		return
	}

	connStr := fmt.Sprintf("%v:%v@tcp(%v)/%v", config.Data.SQL.Username, config.Data.SQL.Password,
		config.Data.SQL.Host, config.Data.SQL.Database)

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	SQL = db

	log.Info("Successfully connected to SQL Database!")
}

// CloseSQLConnection Closes the existing SQL connection
func CloseSQLConnection() {
	if SQL == nil {
		return
	}

	err := SQL.Close()

	if err != nil {
		return
	}

	log.Info("SQL Database connection has been closed")
	SQL = nil
}
