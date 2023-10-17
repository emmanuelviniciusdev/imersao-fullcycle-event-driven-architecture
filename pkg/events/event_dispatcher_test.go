package events_test

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type Event struct {
	Name    string
	Payload interface{}
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetDateTime() time.Time {
	return time.Now()
}

func (e *Event) GetPayload() interface{} {
	return e.Payload
}

type EventHandler struct {
	ID int
}

func (eh *EventHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {}

type EventHandlerMock struct {
	mock.Mock
}

func (ehm *EventHandlerMock) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	ehm.Called(event)
	wg.Done()
}

type EventDispatcherTestSuite struct {
	suite.Suite

	eventFoo Event
	eventBar Event

	eventHandler1 EventHandler
	eventHandler2 EventHandler

	eventDispatcher *events.EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventFoo = Event{Name: "foo", Payload: "content... (foo)"}
	suite.eventBar = Event{Name: "bar", Payload: "content... (bar)"}

	suite.eventHandler1 = EventHandler{ID: 1}
	suite.eventHandler2 = EventHandler{ID: 2}

	suite.eventDispatcher = events.NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	eventNameFoo := suite.eventFoo.GetName()
	eventNameBar := suite.eventBar.GetName()

	err := suite.eventDispatcher.Register(eventNameFoo, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameFoo]))

	err = suite.eventDispatcher.Register(eventNameFoo, &suite.eventHandler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.GetHandlers()[eventNameFoo]))

	err = suite.eventDispatcher.Register(eventNameBar, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameBar]))

	suite.eventDispatcher.Remove(eventNameFoo, &suite.eventHandler1)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameFoo]))
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameBar]))

	suite.eventDispatcher.Remove(eventNameFoo, &suite.eventHandler2)
	suite.Equal(0, len(suite.eventDispatcher.GetHandlers()[eventNameFoo]))
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameBar]))

	suite.eventDispatcher.Remove(eventNameFoo, &suite.eventHandler2)
	suite.Equal(0, len(suite.eventDispatcher.GetHandlers()[eventNameFoo]))
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventNameBar]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandlerMock := &EventHandlerMock{}
	eventHandlerMock.On("Handle", &suite.eventFoo).Return()

	eventName := suite.eventFoo.GetName()

	err := suite.eventDispatcher.Register(eventName, eventHandlerMock)
	suite.Nil(err)

	err = suite.eventDispatcher.Dispatch(&suite.eventFoo)
	suite.Nil(err)

	eventHandlerMock.AssertExpectations(suite.T())
	eventHandlerMock.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	eventName := suite.eventFoo.GetName()

	err := suite.eventDispatcher.Register(eventName, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventName]))

	eventHandlerExists := suite.eventDispatcher.Has(eventName, &suite.eventHandler1)
	suite.True(eventHandlerExists)

	eventHandlerExists = suite.eventDispatcher.Has(eventName, &suite.eventHandler2)
	suite.False(eventHandlerExists)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	eventName := suite.eventFoo.GetName()

	err := suite.eventDispatcher.Register(eventName, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.GetHandlers()))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_SameEventHandlerTwice() {
	eventName := suite.eventFoo.GetName()

	err := suite.eventDispatcher.Register(eventName, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventName]))

	err = suite.eventDispatcher.Register(eventName, &suite.eventHandler1)
	suite.NotNil(err)
	suite.Equal(err.Error(), "event handler already registered")
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventName]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	eventName := suite.eventFoo.GetName()

	err := suite.eventDispatcher.Register(eventName, &suite.eventHandler1)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.GetHandlers()[eventName]))

	err = suite.eventDispatcher.Register(eventName, &suite.eventHandler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.GetHandlers()[eventName]))

	assert.Equal(suite.T(), &suite.eventHandler1, suite.eventDispatcher.GetHandlers()[eventName][0])
	assert.Equal(suite.T(), &suite.eventHandler2, suite.eventDispatcher.GetHandlers()[eventName][1])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
