// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	io "io"

	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"

	request "github.com/mattermost/mattermost/server/public/shared/request"
)

// OAuthProvider is an autogenerated mock type for the OAuthProvider type
type OAuthProvider struct {
	mock.Mock
}

// GetSSOSettings provides a mock function with given fields: c, config, service
func (_m *OAuthProvider) GetSSOSettings(c request.CTX, config *model.Config, service string) (*model.SSOSettings, error) {
	ret := _m.Called(c, config, service)

	var r0 *model.SSOSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Config, string) (*model.SSOSettings, error)); ok {
		return rf(c, config, service)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Config, string) *model.SSOSettings); ok {
		r0 = rf(c, config, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SSOSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Config, string) error); ok {
		r1 = rf(c, config, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFromIdToken provides a mock function with given fields: c, idToken
func (_m *OAuthProvider) GetUserFromIdToken(c request.CTX, idToken string) (*model.User, error) {
	ret := _m.Called(c, idToken)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) (*model.User, error)); ok {
		return rf(c, idToken)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.User); ok {
		r0 = rf(c, idToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) error); ok {
		r1 = rf(c, idToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFromJSON provides a mock function with given fields: c, data, tokenUser
func (_m *OAuthProvider) GetUserFromJSON(c request.CTX, data io.Reader, tokenUser *model.User) (*model.User, error) {
	ret := _m.Called(c, data, tokenUser)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, io.Reader, *model.User) (*model.User, error)); ok {
		return rf(c, data, tokenUser)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, io.Reader, *model.User) *model.User); ok {
		r0 = rf(c, data, tokenUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, io.Reader, *model.User) error); ok {
		r1 = rf(c, data, tokenUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsSameUser provides a mock function with given fields: c, dbUser, oAuthUser
func (_m *OAuthProvider) IsSameUser(c request.CTX, dbUser *model.User, oAuthUser *model.User) bool {
	ret := _m.Called(c, dbUser, oAuthUser)

	var r0 bool
	if rf, ok := ret.Get(0).(func(request.CTX, *model.User, *model.User) bool); ok {
		r0 = rf(c, dbUser, oAuthUser)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewOAuthProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewOAuthProvider creates a new instance of OAuthProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOAuthProvider(t mockConstructorTestingTNewOAuthProvider) *OAuthProvider {
	mock := &OAuthProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
