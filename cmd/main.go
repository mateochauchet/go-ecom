package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mateochauchet/go-ecom/cmd/api"
	"github.com/mateochauchet/go-ecom/config"
	"github.com/mateochauchet/go-ecom/db"
)

var PORT = ":8080"

func main() {
	db, dbErr := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAdress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	initSorage(db)

	server := api.NewAPIServer(PORT, db)

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func initSorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected !")
}
