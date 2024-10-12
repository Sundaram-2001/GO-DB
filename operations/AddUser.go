package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// var users []User

func Add_User(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusBadRequest)
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(res, "Error in request body", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(query, user.Name).Scan(&user.ID)
	if err != nil {
		fmt.Println("Error inserting data", err)
		http.Error(res, "Error Inserting Data", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"Message": "Data Inserted Successfully",
		"Name":    user.Name,
		"ID":      user.ID,
	}
	res.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "INternal Server Error", http.StatusInternalServerError)
		return
	}
	res.Write(jsonResponse)
}
