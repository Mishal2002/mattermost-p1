// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// EmojiStore is an autogenerated mock type for the EmojiStore type
type EmojiStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: emoji, timestamp
func (_m *EmojiStore) Delete(emoji *model.Emoji, timestamp int64) error {
	ret := _m.Called(emoji, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Emoji, int64) error); ok {
		r0 = rf(emoji, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: c, id, allowFromCache
func (_m *EmojiStore) Get(c request.CTX, id string, allowFromCache bool) (*model.Emoji, error) {
	ret := _m.Called(c, id, allowFromCache)

	var r0 *model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string, bool) (*model.Emoji, error)); ok {
		return rf(c, id, allowFromCache)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, bool) *model.Emoji); ok {
		r0 = rf(c, id, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, bool) error); ok {
		r1 = rf(c, id, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: c, name, allowFromCache
func (_m *EmojiStore) GetByName(c request.CTX, name string, allowFromCache bool) (*model.Emoji, error) {
	ret := _m.Called(c, name, allowFromCache)

	var r0 *model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string, bool) (*model.Emoji, error)); ok {
		return rf(c, name, allowFromCache)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, bool) *model.Emoji); ok {
		r0 = rf(c, name, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, bool) error); ok {
		r1 = rf(c, name, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: offset, limit, sort
func (_m *EmojiStore) GetList(offset int, limit int, sort string) ([]*model.Emoji, error) {
	ret := _m.Called(offset, limit, sort)

	var r0 []*model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*model.Emoji, error)); ok {
		return rf(offset, limit, sort)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.Emoji); ok {
		r0 = rf(offset, limit, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(offset, limit, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMultipleByName provides a mock function with given fields: c, names
func (_m *EmojiStore) GetMultipleByName(c request.CTX, names []string) ([]*model.Emoji, error) {
	ret := _m.Called(c, names)

	var r0 []*model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, []string) ([]*model.Emoji, error)); ok {
		return rf(c, names)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, []string) []*model.Emoji); ok {
		r0 = rf(c, names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, []string) error); ok {
		r1 = rf(c, names)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: emoji
func (_m *EmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {
	ret := _m.Called(emoji)

	var r0 *model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Emoji) (*model.Emoji, error)); ok {
		return rf(emoji)
	}
	if rf, ok := ret.Get(0).(func(*model.Emoji) *model.Emoji); ok {
		r0 = rf(emoji)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Emoji) error); ok {
		r1 = rf(emoji)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: name, prefixOnly, limit
func (_m *EmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {
	ret := _m.Called(name, prefixOnly, limit)

	var r0 []*model.Emoji
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool, int) ([]*model.Emoji, error)); ok {
		return rf(name, prefixOnly, limit)
	}
	if rf, ok := ret.Get(0).(func(string, bool, int) []*model.Emoji); ok {
		r0 = rf(name, prefixOnly, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool, int) error); ok {
		r1 = rf(name, prefixOnly, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEmojiStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmojiStore creates a new instance of EmojiStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmojiStore(t mockConstructorTestingTNewEmojiStore) *EmojiStore {
	mock := &EmojiStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
