package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var WITPGH *sql.database

func getConnection() string {
	return os.Geetenv("DATABASE_URL")
}