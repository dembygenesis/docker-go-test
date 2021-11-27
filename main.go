package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Load .env
	// Open up our database connection.
	// db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't load ENV")
	}

	envDBPort := os.Getenv("DB_PORT")
	envDBHost := os.Getenv("DB_HOST")
	envDBUser := os.Getenv("DB_USERNAME")
	envDBPass := os.Getenv("DB_PASSWORD")
	envDBSchema := os.Getenv("DB_SCHEMA")

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		envDBUser,
		envDBPass,
		envDBHost,
		envDBPort,
		envDBSchema,
	))

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
