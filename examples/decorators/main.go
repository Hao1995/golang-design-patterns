package main

import (
	"golang-design-patterns/decorators"
)

func main() {
	notifier := decorators.NewImplNotifier()
	slackNotifyDecorator := decorators.NewSlackNotifyDecorator(notifier)

	slackNotifyDecorator.Notify("Notification test")
}
