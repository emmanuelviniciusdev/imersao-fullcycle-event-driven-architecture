package events

import (
	"errors"
	"sync"
)

var ErrEventHandlerAlreadyRegistered = errors.New("event handler already registered")
var ErrEventHandlerNotFoundByEventName = errors.New("no event handlers found by the specified event name")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, eventHandler EventHandlerInterface) error {
	if _, eventNameExists := ed.handlers[eventName]; eventNameExists {
		for _, mappedEventHandler := range ed.handlers[eventName] {
			if mappedEventHandler == eventHandler {
				return ErrEventHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], eventHandler)

	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	handlersByEventName, handlersByEventNameExists := ed.handlers[event.GetName()]

	if !handlersByEventNameExists {
		return ErrEventHandlerNotFoundByEventName
	}

	wg := &sync.WaitGroup{}

	for _, eventHandler := range handlersByEventName {
		wg.Add(1)
		go eventHandler.Handle(event, wg)
	}

	wg.Wait()

	return nil
}

func (ed *EventDispatcher) Remove(eventName string, eventHandler EventHandlerInterface) {
	if _, eventNameExists := ed.handlers[eventName]; eventNameExists {
		for index, mappedEventHandler := range ed.handlers[eventName] {
			if mappedEventHandler == eventHandler {
				ed.handlers[eventName] =
					append(
						ed.handlers[eventName][:index],
						ed.handlers[eventName][index+1:]...)
			}
		}
	}
}

func (ed *EventDispatcher) Has(eventName string, eventHandler EventHandlerInterface) bool {
	if _, eventNameExists := ed.handlers[eventName]; eventNameExists {
		for _, mappedEventHandler := range ed.handlers[eventName] {
			if mappedEventHandler == eventHandler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) GetHandlers() map[string][]EventHandlerInterface {
	return ed.handlers
}
