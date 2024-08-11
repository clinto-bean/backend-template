package db

import "github.com/google/uuid"

type DB struct {
	Database string
	Addr     string
	User     string
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Id        uuid.UUID
}
