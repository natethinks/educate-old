package main

import (
	"fmt"
	"net/http"
)

// Login checks login credentials, creates session and returns session id to client
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the login page!")
	return
}

// Register accepts new user credentials and creates new user if valid
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the register page!")
	return
}

// Secure is for testing a secure endpoint
func Secure(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the secured page")
	return
}

// JSONarray things
type JSONarray struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Test things
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the test page")
	return
}

func Something(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the something page")
	return
}
