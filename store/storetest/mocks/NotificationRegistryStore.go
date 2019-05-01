// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"

// NotificationRegistryStore is an autogenerated mock type for the NotificationRegistryStore type
type NotificationRegistryStore struct {
	mock.Mock
}

// MarkAsReceived provides a mock function with given fields: ackId, time
func (_m *NotificationRegistryStore) MarkAsReceived(ackId string, time int64) *model.AppError {
	ret := _m.Called(ackId, time)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int64) *model.AppError); ok {
		r0 = rf(ackId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Save provides a mock function with given fields: notification
func (_m *NotificationRegistryStore) Save(notification *model.NotificationRegistry) (*model.NotificationRegistry, *model.AppError) {
	ret := _m.Called(notification)

	var r0 *model.NotificationRegistry
	if rf, ok := ret.Get(0).(func(*model.NotificationRegistry) *model.NotificationRegistry); ok {
		r0 = rf(notification)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NotificationRegistry)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.NotificationRegistry) *model.AppError); ok {
		r1 = rf(notification)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateSendStatus provides a mock function with given fields: ackId, status
func (_m *NotificationRegistryStore) UpdateSendStatus(ackId string, status string) *model.AppError {
	ret := _m.Called(ackId, status)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(ackId, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}
