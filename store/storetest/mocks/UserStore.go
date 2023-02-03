// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"

	store "github.com/mattermost/mattermost-server/v6/store"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// AnalyticsActiveCount provides a mock function with given fields: timestamp, options
func (_m *UserStore) AnalyticsActiveCount(timestamp int64, options model.UserCountOptions) (int64, error) {
	ret := _m.Called(timestamp, options)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, model.UserCountOptions) int64); ok {
		r0 = rf(timestamp, options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, model.UserCountOptions) error); ok {
		r1 = rf(timestamp, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsActiveCountForPeriod provides a mock function with given fields: startTime, endTime, options
func (_m *UserStore) AnalyticsActiveCountForPeriod(startTime int64, endTime int64, options model.UserCountOptions) (int64, error) {
	ret := _m.Called(startTime, endTime, options)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64, model.UserCountOptions) int64); ok {
		r0 = rf(startTime, endTime, options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, model.UserCountOptions) error); ok {
		r1 = rf(startTime, endTime, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsGetExternalUsers provides a mock function with given fields: hostDomain
func (_m *UserStore) AnalyticsGetExternalUsers(hostDomain string) (bool, error) {
	ret := _m.Called(hostDomain)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(hostDomain)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsGetGuestCount provides a mock function with given fields:
func (_m *UserStore) AnalyticsGetGuestCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsGetInactiveUsersCount provides a mock function with given fields:
func (_m *UserStore) AnalyticsGetInactiveUsersCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsGetSystemAdminCount provides a mock function with given fields:
func (_m *UserStore) AnalyticsGetSystemAdminCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AutocompleteUsersInChannel provides a mock function with given fields: teamID, channelID, term, options
func (_m *UserStore) AutocompleteUsersInChannel(teamID string, channelID string, term string, options *model.UserSearchOptions) (*model.UserAutocompleteInChannel, error) {
	ret := _m.Called(teamID, channelID, term, options)

	var r0 *model.UserAutocompleteInChannel
	if rf, ok := ret.Get(0).(func(string, string, string, *model.UserSearchOptions) *model.UserAutocompleteInChannel); ok {
		r0 = rf(teamID, channelID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserAutocompleteInChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(teamID, channelID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearAllCustomRoleAssignments provides a mock function with given fields:
func (_m *UserStore) ClearAllCustomRoleAssignments() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *UserStore) ClearCaches() {
	_m.Called()
}

// Count provides a mock function with given fields: options
func (_m *UserStore) Count(options model.UserCountOptions) (int64, error) {
	ret := _m.Called(options)

	var r0 int64
	if rf, ok := ret.Get(0).(func(model.UserCountOptions) int64); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.UserCountOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeactivateGuests provides a mock function with given fields:
func (_m *UserStore) DeactivateGuests() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DemoteUserToGuest provides a mock function with given fields: userID
func (_m *UserStore) DemoteUserToGuest(userID string) (*model.User, error) {
	ret := _m.Called(userID)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *UserStore) Get(ctx context.Context, id string) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *UserStore) GetAll() ([]*model.User, error) {
	ret := _m.Called()

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func() []*model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAfter provides a mock function with given fields: limit, afterID
func (_m *UserStore) GetAllAfter(limit int, afterID string) ([]*model.User, error) {
	ret := _m.Called(limit, afterID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(int, string) []*model.User); ok {
		r0 = rf(limit, afterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(limit, afterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllNotInAuthService provides a mock function with given fields: authServices
func (_m *UserStore) GetAllNotInAuthService(authServices []string) ([]*model.User, error) {
	ret := _m.Called(authServices)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func([]string) []*model.User); ok {
		r0 = rf(authServices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(authServices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProfiles provides a mock function with given fields: options
func (_m *UserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, int64, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) int64); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*model.UserGetOptions) error); ok {
		r2 = rf(options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllProfilesInChannel provides a mock function with given fields: ctx, channelID, allowFromCache
func (_m *UserStore) GetAllProfilesInChannel(ctx context.Context, channelID string, allowFromCache bool) (map[string]*model.User, error) {
	ret := _m.Called(ctx, channelID, allowFromCache)

	var r0 map[string]*model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) map[string]*model.User); ok {
		r0 = rf(ctx, channelID, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, channelID, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsingAuthService provides a mock function with given fields: authService
func (_m *UserStore) GetAllUsingAuthService(authService string) ([]*model.User, error) {
	ret := _m.Called(authService)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(authService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(authService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnyUnreadPostCountForChannel provides a mock function with given fields: userID, channelID
func (_m *UserStore) GetAnyUnreadPostCountForChannel(userID string, channelID string) (int64, error) {
	ret := _m.Called(userID, channelID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(userID, channelID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAuth provides a mock function with given fields: authData, authService
func (_m *UserStore) GetByAuth(authData *string, authService string) (*model.User, error) {
	ret := _m.Called(authData, authService)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*string, string) *model.User); ok {
		r0 = rf(authData, authService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, string) error); ok {
		r1 = rf(authData, authService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserStore) GetByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *UserStore) GetByUsername(username string) (*model.User, error) {
	ret := _m.Called(username)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelGroupUsers provides a mock function with given fields: channelID
func (_m *UserStore) GetChannelGroupUsers(channelID string) ([]*model.User, error) {
	ret := _m.Called(channelID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtagForAllProfiles provides a mock function with given fields:
func (_m *UserStore) GetEtagForAllProfiles() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEtagForProfiles provides a mock function with given fields: teamID
func (_m *UserStore) GetEtagForProfiles(teamID string) string {
	ret := _m.Called(teamID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(teamID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEtagForProfilesNotInTeam provides a mock function with given fields: teamID
func (_m *UserStore) GetEtagForProfilesNotInTeam(teamID string) string {
	ret := _m.Called(teamID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(teamID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFirstSystemAdminID provides a mock function with given fields:
func (_m *UserStore) GetFirstSystemAdminID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForLogin provides a mock function with given fields: loginID, allowSignInWithUsername, allowSignInWithEmail
func (_m *UserStore) GetForLogin(loginID string, allowSignInWithUsername bool, allowSignInWithEmail bool) (*model.User, error) {
	ret := _m.Called(loginID, allowSignInWithUsername, allowSignInWithEmail)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string, bool, bool) *model.User); ok {
		r0 = rf(loginID, allowSignInWithUsername, allowSignInWithEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, bool) error); ok {
		r1 = rf(loginID, allowSignInWithUsername, allowSignInWithEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKnownUsers provides a mock function with given fields: userID
func (_m *UserStore) GetKnownUsers(userID string) ([]string, error) {
	ret := _m.Called(userID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: ctx, ids
func (_m *UserStore) GetMany(ctx context.Context, ids []string) ([]*model.User, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*model.User); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewUsersForTeam provides a mock function with given fields: teamID, offset, limit, viewRestrictions
func (_m *UserStore) GetNewUsersForTeam(teamID string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	ret := _m.Called(teamID, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamID, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, offset, limit, viewRestrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileByGroupChannelIdsForUser provides a mock function with given fields: userID, channelIds
func (_m *UserStore) GetProfileByGroupChannelIdsForUser(userID string, channelIds []string) (map[string][]*model.User, error) {
	ret := _m.Called(userID, channelIds)

	var r0 map[string][]*model.User
	if rf, ok := ret.Get(0).(func(string, []string) map[string][]*model.User); ok {
		r0 = rf(userID, channelIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(userID, channelIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileByIds provides a mock function with given fields: ctx, userIds, options, allowFromCache
func (_m *UserStore) GetProfileByIds(ctx context.Context, userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, error) {
	ret := _m.Called(ctx, userIds, options, allowFromCache)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(context.Context, []string, *store.UserGetByIdsOpts, bool) []*model.User); ok {
		r0 = rf(ctx, userIds, options, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, *store.UserGetByIdsOpts, bool) error); ok {
		r1 = rf(ctx, userIds, options, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfiles provides a mock function with given fields: options
func (_m *UserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, int64, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) int64); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*model.UserGetOptions) error); ok {
		r2 = rf(options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProfilesByUsernames provides a mock function with given fields: usernames, viewRestrictions
func (_m *UserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	ret := _m.Called(usernames, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func([]string, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(usernames, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(usernames, viewRestrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesInChannel provides a mock function with given fields: options
func (_m *UserStore) GetProfilesInChannel(options *model.UserGetOptions) ([]*model.User, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesInChannelByAdmin provides a mock function with given fields: options
func (_m *UserStore) GetProfilesInChannelByAdmin(options *model.UserGetOptions) ([]*model.User, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesInChannelByStatus provides a mock function with given fields: options
func (_m *UserStore) GetProfilesInChannelByStatus(options *model.UserGetOptions) ([]*model.User, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesNotInChannel provides a mock function with given fields: teamID, channelId, groupConstrained, offset, limit, viewRestrictions
func (_m *UserStore) GetProfilesNotInChannel(teamID string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	ret := _m.Called(teamID, channelId, groupConstrained, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, bool, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamID, channelId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool, int, int, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, channelId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesNotInTeam provides a mock function with given fields: teamID, groupConstrained, offset, limit, viewRestrictions
func (_m *UserStore) GetProfilesNotInTeam(teamID string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	ret := _m.Called(teamID, groupConstrained, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, bool, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamID, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, int, int, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, groupConstrained, offset, limit, viewRestrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilesWithoutTeam provides a mock function with given fields: options
func (_m *UserStore) GetProfilesWithoutTeam(options *model.UserGetOptions) ([]*model.User, error) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecentlyActiveUsersForTeam provides a mock function with given fields: teamID, offset, limit, viewRestrictions
func (_m *UserStore) GetRecentlyActiveUsersForTeam(teamID string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	ret := _m.Called(teamID, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamID, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, offset, limit, viewRestrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSystemAdminProfiles provides a mock function with given fields:
func (_m *UserStore) GetSystemAdminProfiles() (map[string]*model.User, error) {
	ret := _m.Called()

	var r0 map[string]*model.User
	if rf, ok := ret.Get(0).(func() map[string]*model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamGroupUsers provides a mock function with given fields: teamID
func (_m *UserStore) GetTeamGroupUsers(teamID string) ([]*model.User, error) {
	ret := _m.Called(teamID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnreadCount provides a mock function with given fields: userID, isCRTEnabled
func (_m *UserStore) GetUnreadCount(userID string, isCRTEnabled bool) (int64, error) {
	ret := _m.Called(userID, isCRTEnabled)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, bool) int64); ok {
		r0 = rf(userID, isCRTEnabled)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(userID, isCRTEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnreadCountForChannel provides a mock function with given fields: userID, channelID
func (_m *UserStore) GetUnreadCountForChannel(userID string, channelID string) (int64, error) {
	ret := _m.Called(userID, channelID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(userID, channelID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersBatchForIndexing provides a mock function with given fields: startTime, startFileID, limit
func (_m *UserStore) GetUsersBatchForIndexing(startTime int64, startFileID string, limit int) ([]*model.UserForIndexing, error) {
	ret := _m.Called(startTime, startFileID, limit)

	var r0 []*model.UserForIndexing
	if rf, ok := ret.Get(0).(func(int64, string, int) []*model.UserForIndexing); ok {
		r0 = rf(startTime, startFileID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserForIndexing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, int) error); ok {
		r1 = rf(startTime, startFileID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersWithInvalidEmails provides a mock function with given fields: page, perPage, restrictedDomains
func (_m *UserStore) GetUsersWithInvalidEmails(page int, perPage int, restrictedDomains string) ([]*model.User, error) {
	ret := _m.Called(page, perPage, restrictedDomains)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.User); ok {
		r0 = rf(page, perPage, restrictedDomains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, perPage, restrictedDomains)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InferSystemInstallDate provides a mock function with given fields:
func (_m *UserStore) InferSystemInstallDate() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUsers provides a mock function with given fields: users
func (_m *UserStore) InsertUsers(users []*model.User) error {
	ret := _m.Called(users)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*model.User) error); ok {
		r0 = rf(users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InvalidateProfileCacheForUser provides a mock function with given fields: userID
func (_m *UserStore) InvalidateProfileCacheForUser(userID string) {
	_m.Called(userID)
}

// InvalidateProfilesInChannelCache provides a mock function with given fields: channelID
func (_m *UserStore) InvalidateProfilesInChannelCache(channelID string) {
	_m.Called(channelID)
}

// InvalidateProfilesInChannelCacheByUser provides a mock function with given fields: userID
func (_m *UserStore) InvalidateProfilesInChannelCacheByUser(userID string) {
	_m.Called(userID)
}

// IsEmpty provides a mock function with given fields: excludeBots
func (_m *UserStore) IsEmpty(excludeBots bool) (bool, error) {
	ret := _m.Called(excludeBots)

	var r0 bool
	if rf, ok := ret.Get(0).(func(bool) bool); ok {
		r0 = rf(excludeBots)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(excludeBots)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDelete provides a mock function with given fields: userID
func (_m *UserStore) PermanentDelete(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PromoteGuestToUser provides a mock function with given fields: userID
func (_m *UserStore) PromoteGuestToUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetAuthDataToEmailForUsers provides a mock function with given fields: service, userIDs, includeDeleted, dryRun
func (_m *UserStore) ResetAuthDataToEmailForUsers(service string, userIDs []string, includeDeleted bool, dryRun bool) (int, error) {
	ret := _m.Called(service, userIDs, includeDeleted, dryRun)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, []string, bool, bool) int); ok {
		r0 = rf(service, userIDs, includeDeleted, dryRun)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, bool, bool) error); ok {
		r1 = rf(service, userIDs, includeDeleted, dryRun)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetLastPictureUpdate provides a mock function with given fields: userID
func (_m *UserStore) ResetLastPictureUpdate(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: user
func (_m *UserStore) Save(user *model.User) (*model.User, error) {
	ret := _m.Called(user)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: teamID, term, options
func (_m *UserStore) Search(teamID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(teamID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(teamID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(teamID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchInChannel provides a mock function with given fields: channelID, term, options
func (_m *UserStore) SearchInChannel(channelID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(channelID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(channelID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(channelID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchInGroup provides a mock function with given fields: groupID, term, options
func (_m *UserStore) SearchInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(groupID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(groupID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(groupID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchNotInChannel provides a mock function with given fields: teamID, channelID, term, options
func (_m *UserStore) SearchNotInChannel(teamID string, channelID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(teamID, channelID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(teamID, channelID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(teamID, channelID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchNotInGroup provides a mock function with given fields: groupID, term, options
func (_m *UserStore) SearchNotInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(groupID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(groupID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(groupID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchNotInTeam provides a mock function with given fields: notInTeamID, term, options
func (_m *UserStore) SearchNotInTeam(notInTeamID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(notInTeamID, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(notInTeamID, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) error); ok {
		r1 = rf(notInTeamID, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchWithoutTeam provides a mock function with given fields: term, options
func (_m *UserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, error) {
	ret := _m.Called(term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.UserSearchOptions) error); ok {
		r1 = rf(term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user, allowRoleUpdate
func (_m *UserStore) Update(user *model.User, allowRoleUpdate bool) (*model.UserUpdate, error) {
	ret := _m.Called(user, allowRoleUpdate)

	var r0 *model.UserUpdate
	if rf, ok := ret.Get(0).(func(*model.User, bool) *model.UserUpdate); ok {
		r0 = rf(user, allowRoleUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserUpdate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, bool) error); ok {
		r1 = rf(user, allowRoleUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAuthData provides a mock function with given fields: userID, service, authData, email, resetMfa
func (_m *UserStore) UpdateAuthData(userID string, service string, authData *string, email string, resetMfa bool) (string, error) {
	ret := _m.Called(userID, service, authData, email, resetMfa)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, *string, string, bool) string); ok {
		r0 = rf(userID, service, authData, email, resetMfa)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *string, string, bool) error); ok {
		r1 = rf(userID, service, authData, email, resetMfa)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFailedPasswordAttempts provides a mock function with given fields: userID, attempts
func (_m *UserStore) UpdateFailedPasswordAttempts(userID string, attempts int) error {
	ret := _m.Called(userID, attempts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(userID, attempts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastPictureUpdate provides a mock function with given fields: userID
func (_m *UserStore) UpdateLastPictureUpdate(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMfaActive provides a mock function with given fields: userID, active
func (_m *UserStore) UpdateMfaActive(userID string, active bool) error {
	ret := _m.Called(userID, active)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(userID, active)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMfaSecret provides a mock function with given fields: userID, secret
func (_m *UserStore) UpdateMfaSecret(userID string, secret string) error {
	ret := _m.Called(userID, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotifyProps provides a mock function with given fields: userID, props
func (_m *UserStore) UpdateNotifyProps(userID string, props map[string]string) error {
	ret := _m.Called(userID, props)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(userID, props)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: userID, newPassword
func (_m *UserStore) UpdatePassword(userID string, newPassword string) error {
	ret := _m.Called(userID, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUpdateAt provides a mock function with given fields: userID
func (_m *UserStore) UpdateUpdateAt(userID string) (int64, error) {
	ret := _m.Called(userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyEmail provides a mock function with given fields: userID, email
func (_m *UserStore) VerifyEmail(userID string, email string) (string, error) {
	ret := _m.Called(userID, email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(userID, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
