package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/clinto-bean/backend-template/pkg/db"
)

type API struct {
	Addr string
	DB   *db.DB
}

type APIHandler func(http.ResponseWriter, *http.Request) error

func NewAPI(addr string, db *db.DB) *API {
	log.Printf("Starting API at %v\n", addr)
	return &API{
		Addr: addr,
		DB:   db,
	}
}

func NewAPIError(msg string, err error) string {
	return fmt.Sprintf("API Error: \n\t%v%v\n,", msg, err.Error())
}

func (a *API) Start() {

	mux := http.NewServeMux()
	CORSHandler := NewCORSHandler(mux)
	srv := &http.Server{
		Addr:    a.Addr,
		Handler: CORSHandler,
	}

	mux.HandleFunc("/v1/test", handleAll(a.handleTest))
	mux.HandleFunc("/v1/users", handleAll(a.handleUsers))

	log.Fatalf("Application exited: %v", srv.ListenAndServe())
}

func decode(r io.Reader) error {
	decoder := json.NewDecoder(r)
	data := UserParams{}
	err := decoder.Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

func (a *API) Shutdown() {

}

func RespondWithJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func handleAll(f APIHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			RespondWithJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}

func NewCORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html;encoding=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
