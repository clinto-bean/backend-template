package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func NewDB(db, addr, username string) (*DB, error) {
	return &DB{
		Database: db,
		Addr:     addr,
		User:     username,
	}, nil
}

func (db *DB) Start() {
	mux := http.NewServeMux()
	sql.Drivers()
	// CORSHandler := NewCORSHandler(mux)
	srv := &http.Server{
		Addr:    db.Addr,
		Handler: mux,
	}

	log.Fatalf("Application exited: %v", srv.ListenAndServe())
}

func (db *DB) Connect() {

}

func (db *DB) Stop() {

}

func (db *DB) WriteDB(table string, data any) error {
	fmt.Printf("Succesfully added %v to table %v", data, table)
	return nil
}
