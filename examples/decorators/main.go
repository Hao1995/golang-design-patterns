package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"golang-design-patterns/decorators"
	"golang-design-patterns/services"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {

	notifier := decorators.NewImplNotifier()

	slackClient := services.NewSlackClientImpl()
	slackNotifyDecorator := decorators.NewSlackNotifyDecorator(notifier, slackClient)

	slackNotifyDecorator.Notify("Notification")
}
