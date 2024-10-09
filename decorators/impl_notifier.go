package decorators

import (
	"log"
)

type ImplNotifier struct {
}

func NewImplNotifier() *ImplNotifier {
	return &ImplNotifier{}
}

func (s *ImplNotifier) Notify(message string) {
	log.Println(message)
}
