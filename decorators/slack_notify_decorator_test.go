package decorators

import (
	"testing"

	"golang-design-patterns/services"
)

func TestSlackNotifyDecoratorNotify(t *testing.T) {
	notifier := NewImplNotifier()
	slackClient := services.NewMockSlackClientImpl()
	slackNotifyDecorator := NewSlackNotifyDecorator(notifier, slackClient)

	err := slackNotifyDecorator.Notify("test")
	if err != nil {
		t.Fatalf("Failed to test Notify, err=%v\n", err)
	}
}
