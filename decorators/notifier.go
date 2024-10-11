package decorators

type Notifier interface {
	Notify(message string) error
}
