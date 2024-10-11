package decorators

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

type SlackNotifyDecorator struct {
	NotifyDecorator
	client *slack.Client
}

func NewSlackNotifyDecorator(notifyDecorator Notifier) *SlackNotifyDecorator {
	return &SlackNotifyDecorator{
		client:          slack.New(os.Getenv("SLACK_TOKEN")),
		NotifyDecorator: *NewNotifyDecorator(notifyDecorator),
	}
}

func (s *SlackNotifyDecorator) Notify(message string) {
	_, _, _, err := s.client.SendMessage(os.Getenv("SLACK_CHANNEL"))
	if err != nil {
		log.Println(err.Error())
	}
	s.Notifier.Notify(message)
}
