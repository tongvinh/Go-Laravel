package data

import (
	"database/sql"
	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	// any models inserted here (and in the New Function)
	// are easily accessible throughout the entire application
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if os.Getenv("DATABASE_TYPE") == "mysql" || os.Getenv("DATABASE_TYPE") == "marialdb" {
		upper, _ = mysql.New(databasePool)
	} else {
		upper, _ = postgresql.New(databasePool)
	}
	return Models{}
}