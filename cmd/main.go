package main

import (
	"database/sql"
	"log"

	"github.com/Chandra5468/golangproject3/cmd/api"
	"github.com/Chandra5468/golangproject3/config"
	"github.com/Chandra5468/golangproject3/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Unable to establish connection with database.", err)
	}

	initStorage(db)
	server := api.NewApiServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal("Error running server", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal("Unable to ping database.", err)
	}

	log.Println("Database is alive and successfully connected")
}
