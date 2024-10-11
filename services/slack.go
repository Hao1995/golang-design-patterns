package services

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackClient interface {
	SendMessage(channel string) error
}

type SlackClientImpl struct {
	client *slack.Client
}

func NewSlackClientImpl() *SlackClientImpl {
	return &SlackClientImpl{
		client: slack.New(os.Getenv("SLACK_TOKEN")),
	}
}

func (s *SlackClientImpl) SendMessage(channel string) error {
	_, _, _, err := s.client.SendMessage(channel)
	if err != nil {
		return err
	}
	return nil
}
