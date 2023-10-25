package main

import server "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/bootstrap/server"

func main() {
	config := server.NewConfig("Twitter Clone", "localhost", "8080")
	apiServer := server.NewServer(config)

	err := apiServer.Run()
	if err != nil {
		panic(err)
	}
}
