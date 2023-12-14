// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	_ "github.com/denisenkom/go-mssqldb"
// 	"github.com/gorilla/mux" // Import the Gorilla Mux router for routing
// )

// var db *sql.DB

// func main() {
// 	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"
// 	var err error
// 	db, err = sql.Open("sqlserver", connString)
// 	if err != nil {
// 		fmt.Println("Error connecting to the database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// Test the connection
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("Error pinging database:", err.Error())
// 		return
// 	}

// 	fmt.Println("Connected to the database!")

// 	// fs := http.FileServer(http.Dir("../my-vue-app/dist"))
// 	// http.Handle("/", fs)

// 	router := mux.NewRouter() // Create a Gorilla Mux router
// 	router.HandleFunc("/api/login", loginAPI).Methods(http.MethodPost)

// 	// Serve your Vue.js frontend from "../my-vue-app/dist"
// 	fs := http.FileServer(http.Dir("../my-vue-app/dist"))
// 	router.PathPrefix("/").Handler(fs)

// 	// Serve static files (CSS, JavaScript, etc.) from a "static" directory
// 	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// 	http.HandleFunc("/successfulSignup", SignUp)

// 	// Serve static files (CSS, JavaScript, etc.)
// 	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// 	// fmt.Println("Server started on :8080")
// 	// http.ListenAndServe(":8080", nil)

// 	// Start the HTTP server
// 	fmt.Println("Server started on :8080")
// 	http.Handle("/", router)
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux" // Import the Gorilla Mux router for routing
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router := mux.NewRouter() // Create a Gorilla Mux router

	// Apply CORS middleware to your router
	handler := c.Handler(router)

	// Serve your Vue.js frontend from "../my-vue-app/dist"
	fs := http.FileServer(http.Dir("../my-vue-app/dist"))
	router.PathPrefix("/").Handler(fs)

	// Serve static files (CSS, JavaScript, etc.) from a "static" directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/api/login", loginAPI).Methods(http.MethodPost)
	http.HandleFunc("/successfulSignup", SignUp)

	// Start the HTTP server
	fmt.Println("Server started on :8080")
	http.Handle("/", handler) // Use the CORS-enabled handler to handle requests
	http.ListenAndServe(":8080", nil)
}
