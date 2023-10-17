package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, eventHandler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, eventHandler EventHandlerInterface)
	Has(eventName string, eventHandler EventHandlerInterface) bool
	Clear()
	GetHandlers() map[string][]EventHandlerInterface
}
