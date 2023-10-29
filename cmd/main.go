package main

import (
	server "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/bootstrap/server"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/db"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/tweet"
	tweetWorfklow "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/workflow/tweet"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	tweetRepository := tweet.NewRepository(db)
	tweetActivities := tweetWorfklow.NewTweetActivities(tweetRepository)
	workflowInstance := tweetWorfklow.NewTweetWorkflow(tweetActivities)
	w := worker.New(c, "TweetTaskQueue", worker.Options{})
	w.RegisterWorkflow(workflowInstance.ProcessTweet)
	w.RegisterActivity(tweetActivities.SaveTweet)

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			panic(err)
		}
	}()

	config := server.NewConfig("Twitter Clone", "localhost", "8080")

	// Inicia el servidor REST. Puede estar en la goroutine principal
	serverInstance := server.NewServer(config, db)
	err = serverInstance.Run()
	if err != nil {
		panic(err)
	}
}
