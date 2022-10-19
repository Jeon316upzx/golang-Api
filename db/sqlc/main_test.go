package bankdb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries

const (
	dbDriver = "mysql"
	dbSource = "root:bankapi18@tcp(localhost:5432)/bankdb"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())

}
