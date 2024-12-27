
package main

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

// DB holds the database connection
var DB *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() {
    var err error
    connStr := "postgres://user:password@localhost:5432/userdb?sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    // Create table if it doesn't exist
    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL
        )
    `)
    if err != nil {
        log.Fatalf("Could not create users table: %v", err)
    }
    log.Println("Database connected and table created (if not exists).")
}

// CreateUserDB inserts a new user into the database
func CreateUserDB(name, email string) (int, error) {
    var id int
    err := DB.QueryRow(
        "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        name, email,
    ).Scan(&id)
    return id, err
}

// GetUserDB retrieves a user by ID from the database
func GetUserDB(id int) (map[string]string, error) {
    var name, email string
    err := DB.QueryRow(
        "SELECT name, email FROM users WHERE id = $1",
        id,
    ).Scan(&name, &email)
    if err != nil {
        return nil, err
    }
    return map[string]string{"name": name, "email": email}, nil
}

// UpdateUserDB updates user details in the database
func UpdateUserDB(id int, name, email string) error {
    _, err := DB.Exec(
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        name, email, id,
    )
    return err
}

// DeleteUserDB deletes a user from the database
func DeleteUserDB(id int) error {
    _, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
    return err
}
