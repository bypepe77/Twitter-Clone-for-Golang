package main

import (
	server "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/bootstrap/server"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/db"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/tweet"
	tweetWorfklow "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/workflow/tweetworfklow"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
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
	w := worker.New(c, "TweetTaskQueue", worker.Options{})
	w.RegisterWorkflowWithOptions(tweetWorfklow.ProcessTweet, workflow.RegisterOptions{Name: "ProcessTweetWorkflow"})
	w.RegisterActivity(tweetActivities.SaveTweet)

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			panic(err)
		}
	}()

	config := server.NewConfig("Twitter Clone", "localhost", "8080")

	serverInstance := server.NewServer(config, db, c)
	err = serverInstance.Run()
	if err != nil {
		panic(err)
	}
}
