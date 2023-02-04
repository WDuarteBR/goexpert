package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	PayLoad interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayLoad() interface{} {
	return e.PayLoad
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event2 = TestEvent{Name: "test2", PayLoad: "teste2"}
	suite.event = TestEvent{Name: "test", PayLoad: "teste"}
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clean() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	suite.dispatcher.Clean()
	suite.Equal(0, len(suite.dispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.dispatcher.Has(suite.event.GetName(), &suite.handler3))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.dispatcher.Register(suite.event.GetName(), eh)

	suite.dispatcher.Dispatche(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	suite.dispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
	suite.dispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
