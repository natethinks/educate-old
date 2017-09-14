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

	// User functions
	router.HandleFunc("/api/user", GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/user", CreateUser).Methods("POST")
	router.HandleFunc("/api/user", DeleteUser).Methods("DELETE")
	// Administration functions
	// is this even necessary? Don't know
	// could be good for promoting people to admin?
	// Subject functions
	router.HandleFunc("/api/subject", GetSubjects).Methods("GET")
	router.HandleFunc("/api/subject/{slug}", GetSubject).Methods("GET")
	router.HandleFunc("/api/subject/{query}", GetSubjectsQuery).Methods("GET")
	router.HandleFunc("/api/subject", CreateSubject).Methods("POST")
	router.HandleFunc("/api/subject", DeleteSubject).Methods("DELETE")
	// Topic functions
	router.HandleFunc("/api/topic", GetTopics).Methods("GET")
	router.HandleFunc("/api/topic/{slug}", GetTopic).Methods("GET")
	router.HandleFunc("/api/topic/{query}", GetTopicsQuery).Methods("GET")
	router.HandleFunc("/api/topic", CreateTopic).Methods("POST")
	router.HandleFunc("/api/topic", DeleteTopic).Methods("POST")
	// Resource functions
	router.HandleFunc("/api/resource", GetResources).Methods("GET")
	router.HandleFunc("/api/resource/{slug}", GetResource).Methods("GET")
	router.HandleFunc("/api/resource/{query}", GetResourceQuery).Methods("GET")
	router.HandleFunc("/api/resource", CreateResource).Methods("POST")
	router.HandleFunc("/api/resource", DeleteResource).Methods("DELETE")

	SetRoute("POST", "/api/secure", Secure)

	return router
}

// SetUnsecuredRoute something
// func SetUnsecuredRoute(method string, path string, fu Handler) {
// 	router.HandleFunc(path, fu).Methods(method)
// }

// func SecureRoute(method string, path string, fu Handler) {
// 	fmt.Println("SecureRoute")
// 	something := r.Header.Get("something")
// 	fmt.Println(something)
// 	router.Handle(path, fu).Methods(method)
// }

// SetRoute something
func SetRoute(method string, path string, fu Handler) {
	fmt.Println("SetRoute")
	router.Handle(path, authedHandler(fu)).Methods(method)
}

func authedHandler(fu Handler) *negroni.Negroni {
	fmt.Println("authedHandler")
	return negroni.New(
		negroni.HandlerFunc(handlerWithNext),
		negroni.Wrap(http.HandlerFunc(fu)),
	)
}

// handlerwithnext is where my control really lies. This is where I check the token

func handlerWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("handlerWithNext")
	err := checkSession(w, r)
	fmt.Println(err)

	// If there was an error, do not call next.
	if err == nil && next != nil {
		fmt.Println("something went wrong")
		next(w, r)
	}
}

func checkSession(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("checkSession")
	return fmt.Errorf("Error extracting token: 1234")
}
