package decorators

import (
	"os"

	"golang-design-patterns/services"
)

type SlackNotifyDecorator struct {
	NotifyDecorator
	client services.SlackClient
}

func NewSlackNotifyDecorator(notifyDecorator Notifier, client services.SlackClient) *SlackNotifyDecorator {
	return &SlackNotifyDecorator{
		client:          client,
		NotifyDecorator: *NewNotifyDecorator(notifyDecorator),
	}
}

func (s *SlackNotifyDecorator) Notify(message string) error {
	if err := s.client.SendMessage(os.Getenv("SLACK_CHANNEL")); err != nil {
		return err
	}
	return s.Notifier.Notify(message)
}
