package main

import (
	"database/sql"
	"fmt"
	"myproject/operations"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the env file: ", err)
	}
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		fmt.Println("Database URL not defined!")
	}
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		fmt.Println("Error connecting the database")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error reaching the database!")
	}
	fmt.Println("Connection Successful!!")
	query := `
		create table if not exists users(
		id serial primary key,
		name text not null
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Error creating the table: ", err)
	}
	fmt.Println("Created Table Successfully!")
	http.HandleFunc("/adduser", func(res http.ResponseWriter, req *http.Request) {
		operations.Add_User(db, res, req)
	})
	http.HandleFunc("/getusers", func(res http.ResponseWriter, req *http.Request) {
		operations.GetAllUsers(db, res, req)
	})
	http.HandleFunc("/getuser", func(res http.ResponseWriter, req *http.Request) {
		operations.GetUser(db, res, req)
	})
	PORT := "8090"
	http.ListenAndServe(":"+PORT, nil)
}
