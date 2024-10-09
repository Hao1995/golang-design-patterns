package decorators

import (
	"log"

	"github.com/slack-go/slack"
)

type SlackNotifyDecorator struct {
	NotifyDecorator
	client *slack.Client
}

func NewSlackNotifyDecorator(notifyDecorator Notifier) *SlackNotifyDecorator {
	return &SlackNotifyDecorator{
		client:          slack.New("YOUR_TOKEN_HERE"),
		NotifyDecorator: *NewNotifyDecorator(notifyDecorator),
	}
}

func (s *SlackNotifyDecorator) Notify(message string) {
	_, _, _, err := s.client.SendMessage("CHANNEL")
	if err != nil {
		log.Println(err.Error())
	}
	s.Notifier.Notify(message)
}
