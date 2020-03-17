package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// Database implements `Downstream`
type Database struct {
	address string
	conn    *pgx.Conn
}

// Start initiates the postgres connection
func (db *Database) Start() error {
	conn, err := pgx.Connect(context.Background(), db.address)
	if err != nil {
		log.Fatal("Could not initialize db connection", err)
	}
	db.conn = conn
	return err
}

// Close closes the postgres connection
func (db *Database) Close() {
	if err := db.conn.Close(context.Background()); err != nil {
		log.Fatal("Could not close db connection", err)
	}
	db.conn = nil
}

// New returns a new instance of the database wrapper
func New(address string) *Database {
	return &Database{
		address: address,
		conn:    nil,
	}
}
