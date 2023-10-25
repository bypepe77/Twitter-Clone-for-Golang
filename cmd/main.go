package main

import server "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/bootstrap/server"

func main() {
	config := server.NewConfig("Twitter Clone", "localhost", "8080")
	serverInstance := server.NewServer(config)

	err := serverInstance.Run()
	if err != nil {
		panic(err)
	}
}
