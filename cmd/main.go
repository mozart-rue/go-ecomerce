package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mozart-rue/go-ecomerce/cmd/api"
	"github.com/mozart-rue/go-ecomerce/config"
	"github.com/mozart-rue/go-ecomerce/db"
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
		fmt.Println(" Erro ao connectar ao banco")
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println(" Erro ao fazer PING ao banco")
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
