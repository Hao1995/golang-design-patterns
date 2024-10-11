package main

import (
	"golang-design-patterns/decorators"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {

	notifier := decorators.NewImplNotifier()
	slackNotifyDecorator := decorators.NewSlackNotifyDecorator(notifier)

	slackNotifyDecorator.Notify("Notification")
}
