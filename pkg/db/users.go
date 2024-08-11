package db

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUserById(id uuid.UUID) (User, error) {
	// query := fmt.Sprintf("SELECT * FROM users WHERE Id = %v", id)
	return User{}, nil
}

func (db *DB) CreateUser(first, last, email string, id uuid.UUID) (User, error) {
	var returnVals = User{
		FirstName: first,
		LastName:  last,
		Email:     email,
		Id:        id,
	}
	fmt.Printf("\n\tDB: User created\n\t\t%v\n\t\t%v\n\t\t%v\n\t\t%v\n\tSuccesfully", first, last, email, id)
	return returnVals, nil
}

func (db *DB) UpdateUser(first, last, email string, id uuid.UUID) (User, error) {
	existing, err := GetUserById(id)
	if err != nil {
		return User{}, fmt.Errorf("user not found with parameter (id = %v)", id)
	}
	existing.FirstName = first
	existing.LastName = last
	existing.Email = email
	err = db.WriteDB("users", existing)
	if err != nil {
		return User{}, fmt.Errorf("could not write to database")
	}
	return existing, err
}
