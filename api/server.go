package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/negroni"
)

func main() {
	// loading routes from routes.go
	myRouter := LoadRoutes()
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(myRouter)

	// set up for docker mysql image
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/mysqldb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", n)
}
