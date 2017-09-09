package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Handler something
type Handler func(w http.ResponseWriter, r *http.Request)

var router *mux.Router

// LoadRoutes something
func LoadRoutes() *mux.Router {
	router = mux.NewRouter()

	router.HandleFunc("/api", Test).Methods("GET")
	router.HandleFunc("/api/login", Login).Methods("POST")
	router.HandleFunc("/api/register", Register).Methods("POST")
	router.HandleFunc("/api/test", Test).Methods("GET")
	router.HandleFunc("/api/test", Something).Methods("GET")

	SetRoute("POST", "/secure", Secure)

	return router
}

// SetUnsecuredRoute something
// func SetUnsecuredRoute(method string, path string, fu Handler) {
// 	router.HandleFunc(path, fu).Methods(method)
// }

// SetRoute something
func SetRoute(method string, path string, fu Handler) {
	router.Handle(path, authedHandler(fu)).Methods(method)
}

func authedHandler(fu Handler) *negroni.Negroni {
	return negroni.New(
		negroni.HandlerFunc(handlerWithNext),
		negroni.Wrap(http.HandlerFunc(fu)),
	)
}

// handlerwithnext is where my control really lies. This is where I check the token

func handlerWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := checkSession(w, r)

	// If there was an error, do not call next.
	if err == nil && next != nil {
		next(w, r)
	}
}

func checkSession(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("Error extracting token: 1234")
}
