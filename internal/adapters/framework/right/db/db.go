package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	// blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// Connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB connection failure: %v", err)
	}

	// Test DB connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("DB Ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	if err := da.db.Close(); err != nil {
		log.Fatalf("DB closing connection failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).
		ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
