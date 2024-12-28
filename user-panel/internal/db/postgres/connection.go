package postgres

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(host, port, user, password, dbname string) error {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        return err
    }
    return DB.Ping()
}
