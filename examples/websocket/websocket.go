package main

import (
	"fmt"
	"log"
	"os"
        "strings"
	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("TOKEN")
        bot_id := os.Getenv("BOT_ID")
	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				// Replace #general with your Channel ID
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "G2NJ8BD5E"))

			case *slack.MessageEvent:
                                command := strings.Split(ev.Text, " ")
				fmt.Printf("Message: %v\n",  command)
                                if bot_id == command[0]{
                                    fmt.Printf("This is command for bot")
                                }
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:

				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
