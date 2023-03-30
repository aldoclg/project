package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/aldoclg/project/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config file")
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) //go get github.com/lib/pq
	// go get github.com/stretchr/testify
}
