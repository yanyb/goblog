package messaging

import "github.com/streadway/amqp"

import (
	"github.com/stretchr/testify/mock"
	"context"
)

// MockMessagingClient is an autogenerated mock type for the IMessagingClient type
type MockMessagingClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockMessagingClient) Close() {
	_m.Called()
}

// ConnectToBroker provides a mock function with given fields: connectionString
func (_m *MockMessagingClient) ConnectToBroker(connectionString string) {
	_m.Called(connectionString)
}

// Publish provides a mock function with given fields: msg, exchangeName, exchangeType
func (_m *MockMessagingClient) Publish(msg []byte, exchangeName string, exchangeType string) error {
	ret := _m.Called(msg, exchangeName, exchangeType)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, string, string) error); ok {
		r0 = rf(msg, exchangeName, exchangeType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishOnQueue provides a mock function with given fields: msg, queueName
func (_m *MockMessagingClient) PublishOnQueue(msg []byte, queueName string) error {
	ret := _m.Called(msg, queueName)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, string) error); ok {
		r0 = rf(msg, queueName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishOnQueueWithContext provides a mock function with given fields: ctx, msg, queueName
func (_m *MockMessagingClient) PublishOnQueueWithContext(ctx context.Context, msg []byte, queueName string) error {
	ret := _m.Called(ctx, msg, queueName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) error); ok {
		r0 = rf(ctx, msg, queueName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: exchangeName, exchangeType, consumerName, handlerFunc
func (_m *MockMessagingClient) Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ret := _m.Called(exchangeName, exchangeType, consumerName, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, func(amqp.Delivery)) error); ok {
		r0 = rf(exchangeName, exchangeType, consumerName, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeToQueue provides a mock function with given fields: queueName, consumerName, handlerFunc
func (_m *MockMessagingClient) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ret := _m.Called(queueName, consumerName, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, func(amqp.Delivery)) error); ok {
		r0 = rf(queueName, consumerName, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
