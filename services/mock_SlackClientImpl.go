package services

import (
	"github.com/slack-go/slack"
)

type MockSlackClientImpl struct {
	client *slack.Client
}

func NewMockSlackClientImpl() *MockSlackClientImpl {
	return &MockSlackClientImpl{}
}

func (s *MockSlackClientImpl) SendMessage(channel string) error {
	return nil
}
