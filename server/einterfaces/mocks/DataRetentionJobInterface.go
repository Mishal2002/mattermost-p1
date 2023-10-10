// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	jobs "github.com/mattermost/mattermost/server/v8/einterfaces/jobs"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"
)

// DataRetentionJobInterface is an autogenerated mock type for the DataRetentionJobInterface type
type DataRetentionJobInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with given fields:
func (_m *DataRetentionJobInterface) MakeScheduler() jobs.Scheduler {
	ret := _m.Called()

	var r0 jobs.Scheduler
	if rf, ok := ret.Get(0).(func() jobs.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jobs.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with given fields:
func (_m *DataRetentionJobInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}

type mockConstructorTestingTNewDataRetentionJobInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataRetentionJobInterface creates a new instance of DataRetentionJobInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataRetentionJobInterface(t mockConstructorTestingTNewDataRetentionJobInterface) *DataRetentionJobInterface {
	mock := &DataRetentionJobInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
