package events

import (
	"errors"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already resgistered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clean() {
	ed.handlers = make(map[string][]EventHandlerInterface)

}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatche(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, h := range handlers {
			h.Handle(event)
		}
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				// na lógica dentro do append temos:
				// ed.handlers[eventName][:i] => o próprio handler a ser removido
				// ed.handlers[eventName][i+1:] => com excessão do próprio acrescenta todo o resto
				// ... => indica que dé uma adição de slices

			}
		}
	}
	return nil
}
