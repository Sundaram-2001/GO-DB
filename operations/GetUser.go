package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io" // Fix typo here
	"net/http"

	_ "github.com/lib/pq"
)

type UserId struct {
	ID int `json:"id"`
}

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// var db *sql.DB

func GetUser(db *sql.DB, res http.ResponseWriter, req *http.Request) { // Fix signature
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body) // Fix ioutil typo
	if err != nil {
		http.Error(res, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var userId UserId // Declare userId before using it

	err = json.Unmarshal(body, &userId)
	if err != nil {
		http.Error(res, "Error unmarshalling json", http.StatusBadRequest)
		return
	}

	var user User
	row := db.QueryRow("Select id, name from users where id=$1", userId.ID)
	err = row.Scan(&user.ID, &user.Name) // Fix row.Scan usage
	if err == sql.ErrNoRows {
		http.Error(res, "Unable to find the data for given id", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error fetching the data", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Write(jsonResponse)
}
