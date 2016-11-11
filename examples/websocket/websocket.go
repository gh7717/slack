package main

import (
	"fmt"
	"log"
	"os"
        "strings"
	"github.com/gh7717/slack"
        "gopkg.in/mgo.v2"
)

func main() {
	token := os.Getenv("TOKEN")
        bot_id := os.Getenv("BOT_ID")
        channel_id := os.Getenv("CHANNEL_ID")
        mongo := os.Getenv("MONGO")
        session, err := mgo.Dial(mongo)
	if err != nil {
            log.Fatal(err)
        }
        api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)
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
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", channel_id))

			case *slack.MessageEvent:
                                command := strings.Split(ev.Text, " ")
				fmt.Printf("Message: %v\n",  command)
                                if bot_id == command[0][2:len(command[0])-1]{
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
