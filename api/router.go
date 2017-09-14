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

// User functions
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUsers")
	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUser")
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUser")
	return
}

// Subject functions

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetSubjects")
	return
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func GetSubjectsQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

// Topic functions

func GetTopics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func GetTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func GetTopicsQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

//ResourceFunctions

func GetResources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func GetResourceQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	return
}
