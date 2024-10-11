package decorators

type NotifyDecorator struct {
	Notifier Notifier
}

func NewNotifyDecorator(notifier Notifier) *NotifyDecorator {
	return &NotifyDecorator{
		Notifier: notifier,
	}
}

func (s *NotifyDecorator) Notify(message string) error {
	return s.Notifier.Notify(message)
}
