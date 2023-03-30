package main

import (
	"database/sql"
	"log"

	"github.com/aldoclg/project/api"
	db "github.com/aldoclg/project/db/sqlc"
	"github.com/aldoclg/project/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@localhost:5432/root?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	//queries := db.New(conn)

	//os.Exit(m.Run())
}
