package events

import "time"

// O evento em si, ou seja,
// quem carrega os dados
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() interface{}
}

// É quem executa as operações quando
// o vevento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// É quem gerencia os eventos/operações
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatche(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clean()
}
