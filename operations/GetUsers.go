package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			fmt.Println("Error scanning row:", err)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}