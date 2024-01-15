package main

import (
	"log"
	"time"

	"github.com/s-yakubovskiy/inst2tg/pkg/inst2tg"
	"github.com/zelenin/go-tdlib/client"
)

func main() {
	// client authorizer
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	const (
		apiId   = 00000       // get from https://my.telegram.org/apps
		apiHash = "your-hash" // get from https://my.telegram.org/apps
	)

	tgClient := inst2tg.NewTGClient(apiId, apiHash, 2)
	err := tgClient.Initialize()
	if err != nil {
		log.Fatalf("Error initializing Telegram client: %s", err)
	}
	user, err := tgClient.GetMe()
	if err != nil {
		log.Fatalf("Error: Failed to get user")
	}

	// tgClient.SendStory(user.Id, inst2tg.SendStoryRequest{Photo: true, Path: "/tmp/l2", Local: true, ActivePeriod: 43200})
	tgClient.SendStory(user.Id, inst2tg.SendStoryRequest{Video: true, Path: "path-to-file(remote or local)", Local: false, ActivePeriod: 43200})
	time.Sleep(5 * time.Second)
	tgClient.Client.Close()

	// time for flush
	time.Sleep(5 * time.Second)
}
