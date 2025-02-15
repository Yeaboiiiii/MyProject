package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.sql") // Change the DB type accordingly
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	createTables()
}

func createTables() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE, 
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Fatalf("could not create user table: %v", err) // Print the actual error
	}

	query := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userID INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	)`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err) // Print the actual error
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER , 
	user_id INTEGER ,
	FOREIGN KEY(event_id) REFERENCES events(id)
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		log.Fatalf("could not create registration table: %v", err) // Print the actual error
	}
}
