package decorators

type NotifyDecorator struct {
	Notifier Notifier
}

func NewNotifyDecorator(notifyDecorator Notifier) *NotifyDecorator {
	return &NotifyDecorator{
		Notifier: notifyDecorator,
	}
}

func (s *NotifyDecorator) Notify(message string) {
	s.Notifier.Notify(message)
}
