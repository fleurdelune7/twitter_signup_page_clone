package drivers

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)


type DB struct {
	SQL * sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10 
const maxIdleDb = 5
const maxDbLifeTime = 5 * time.Minute


func ConnectSQL(dsn string) (*DB, error){
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDb)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil

}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error){
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}