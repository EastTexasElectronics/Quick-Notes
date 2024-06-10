package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "os"
)

var DB *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS notes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

func InitDB() {
    var err error
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/notes_app"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Cannot ping database:", err)
    }

    _, err = DB.Exec(schema)
    if err != nil {
        log.Fatal("Cannot create tables:", err)
    }

    log.Println("Database initialized and tables created")
}
