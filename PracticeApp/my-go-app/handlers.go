package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	// Import the Gorilla Mux router for routing
)

// func serveLogin(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Serve the login page
// 	http.ServeFile(w, r, "templates/login.html")
// }

func loginAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Printf("Received username: %s, password: %s", username, password)

	// Query the database for the user's credentials
	var storedPassword string
	// err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = ?", username).Scan(&storedPassword)
	err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = @username", sql.Named("username", username)).Scan(&storedPassword)
	// if err != nil {
	// 	log.Printf("Error querying database: %v", err)
	// 	http.Error(w, "User not found", http.StatusUnauthorized)
	// 	return
	// }

	if err == sql.ErrNoRows {
		// User not found in the database
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the stored password
	// if password == storedPassword {
	// 	fmt.Fprintln(w, "Login successful")
	// } else {
	// 	http.Error(w, "Login failed", http.StatusUnauthorized)
	// }

	// Compare the provided password with the stored password
	if password == storedPassword {

		log.Printf("Password = ", password, "StoredPassword = ", storedPassword)

		// Create a JSON response with a redirection URL
		response := map[string]interface{}{
			"success":     true, // Change this to true on successful login
			"redirectURL": "/clusterhome",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Login failed", http.StatusUnauthorized)
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// username := r.FormValue("username")
	// title := r.FormValue("title")
	// firstName := r.FormValue("firstName")
	// middleName := r.FormValue("middleName")
	// lastName := r.FormValue("lastName")
	// password := r.FormValue("password")
	// confirmPassword := r.FormValue("confirmPassword")
	// email := r.FormValue("email")
	// dateOfBirth := r.FormValue("dateOfBirth")
	// nationality := r.FormValue("nationality")
	// cnicNumber := r.FormValue("cnicNumber")
	// passportNumber := r.FormValue("passportNumber")
	// university := r.FormValue("university")
	// thesisTitle := r.FormValue("thesisTitle")

	// // Retrieve checkbox values (services)
	// hadoop := r.FormValue("hadoop")
	// hive := r.FormValue("hive")
	// pig := r.FormValue("Pig")
	// spark := r.FormValue("Spark")
	// sqoop := r.FormValue("Sqoop")
	// mySQL := r.FormValue("MySQL")
	// zookeeper := r.FormValue("Zookeeper")
	// hbase := r.FormValue("Hbase")
	// kafka := r.FormValue("Kafka")
	// flume := r.FormValue("Flume")
	// storm := r.FormValue("Storm")

	// // Retrieve additional fields
	// xyzid := r.FormValue("xyzid")
	// startDate := r.FormValue("startDate")
	// endDate := r.FormValue("endDate")

}
