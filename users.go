package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type UserParams struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
}

func (a *API) handleUsers(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		err := a.handleGetUser(w, r)
		if err != nil {
			msg := fmt.Sprint("GET /v1/users returned an error: ", err)
			return errors.New(msg)
		}
	case "POST":
		err := a.handleCreateUser(w, r)
		if err != nil {
			msg := fmt.Sprint("POST /v1/users returned an error: ", err)
			return errors.New(msg)
		}
	case "PUT":
		err := a.handleModifyUser(w, r)
		if err != nil {
			msg := fmt.Sprint("PUT /v1/users returned an error: ", err)
			return errors.New(msg)
		}
	case "DELETE":
		err := a.handleDeleteUser(w, r)
		if err != nil {
			msg := fmt.Sprint("DELETE /v1/users returned an error: ", err)
			return errors.New(msg)
		}
	}
	return nil
}

func (a *API) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	err := decode(r.Body)
	if err != nil {
		return err
	}
	// dbUser, err := db.GetUser(user.Id)
	RespondWithJSON(w, http.StatusOK, r.Body)
	return nil
}

// func (a *API) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

func (a *API) handleCreateUser(w http.ResponseWriter, r *http.Request) error {

	decoder := json.NewDecoder(r.Body)
	var parameters UserParams
	err := decoder.Decode(&parameters)

	if err != nil {
		RespondWithJSON(w, http.StatusBadRequest, err)
		return err
	}

	userid, err := uuid.NewUUID()
	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, err)
	}

	_, err = a.DB.CreateUser(parameters.FirstName, parameters.LastName, parameters.Email, userid)
	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, err)
	}
	RespondWithJSON(w, http.StatusCreated, fmt.Sprintf("Successfully created user: %v %v (%v)", parameters.FirstName, parameters.LastName, parameters.Email))
	return nil

}

func (a *API) handleModifyUser(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(w, r)

	// user, err := parse(r.Body)

	// if err != nil {
	// 	RespondWithJSON(w, http.StatusBadRequest, err)
	// 	return err
	// }

	// if user.id != dbUser.id {
	// 	RespondWithJSON(w, http.StatusUnauthorized, err)
	// 	return err
	// }

	// err = db.UpdateUser(dbUser, user)

	// if err != nil {
	// 	RespondWithJSON(w, http.StatusInternalServerError, err)
	// 	return err
	// }

	return nil

}

func (a *API) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(w, r)
	return nil

}

func (a *API) handleTest(w http.ResponseWriter, r *http.Request) error {

	RespondWithJSON(w, http.StatusOK, "test succeeded")
	return nil
}
