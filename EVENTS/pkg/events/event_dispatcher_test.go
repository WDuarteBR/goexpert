package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

type TestEventHandler struct{}

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
	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.handler3 = TestEventHandler{}
	suite.event = TestEvent{Name: "test", PayLoad: "teste"}
	suite.event2 = TestEvent{Name: "test2", PayLoad: "teste2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
