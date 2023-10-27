package main

import (
	server "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/bootstrap/server"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/db"
	"github.com/joho/godotenv"
)

func main() {
	config := server.NewConfig("Twitter Clone", "localhost", "8080")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}

	serverInstance := server.NewServer(config, db)
	err = serverInstance.Run()
	if err != nil {
		panic(err)
	}
}
