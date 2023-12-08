package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err.Error())
		return
	}

	fmt.Println("Connected to the database!")

	http.HandleFunc("/", serveLogin)
	http.HandleFunc("/api/login", loginAPI)

	// Serve static files (CSS, JavaScript, etc.)
	http.Handle("/", http.FileServer(http.Dir("frontend-dist")))

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
