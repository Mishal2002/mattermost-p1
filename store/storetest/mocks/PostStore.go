// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// PostStore is an autogenerated mock type for the PostStore type
type PostStore struct {
	mock.Mock
}

// AnalyticsPostCount provides a mock function with given fields: teamID, mustHaveFile, mustHaveHashtag
func (_m *PostStore) AnalyticsPostCount(teamID string, mustHaveFile bool, mustHaveHashtag bool) (int64, error) {
	ret := _m.Called(teamID, mustHaveFile, mustHaveHashtag)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, bool, bool) int64); ok {
		r0 = rf(teamID, mustHaveFile, mustHaveHashtag)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, bool) error); ok {
		r1 = rf(teamID, mustHaveFile, mustHaveHashtag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsPostCountsByDay provides a mock function with given fields: options
func (_m *PostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, error) {
	ret := _m.Called(options)

	var r0 model.AnalyticsRows
	if rf, ok := ret.Get(0).(func(*model.AnalyticsPostCountsOptions) model.AnalyticsRows); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.AnalyticsRows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.AnalyticsPostCountsOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsUserCountsWithPostsByDay provides a mock function with given fields: teamID
func (_m *PostStore) AnalyticsUserCountsWithPostsByDay(teamID string) (model.AnalyticsRows, error) {
	ret := _m.Called(teamID)

	var r0 model.AnalyticsRows
	if rf, ok := ret.Get(0).(func(string) model.AnalyticsRows); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.AnalyticsRows)
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

// ClearCaches provides a mock function with given fields:
func (_m *PostStore) ClearCaches() {
	_m.Called()
}

// Delete provides a mock function with given fields: postID, time, deleteByID
func (_m *PostStore) Delete(postID string, time int64, deleteByID string) error {
	ret := _m.Called(postID, time, deleteByID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, string) error); ok {
		r0 = rf(postID, time, deleteByID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id, skipFetchThreads, collapsedThreads, collapsedThreadsExtended, userID
func (_m *PostStore) Get(ctx context.Context, id string, skipFetchThreads bool, collapsedThreads bool, collapsedThreadsExtended bool, userID string) (*model.PostList, error) {
	ret := _m.Called(ctx, id, skipFetchThreads, collapsedThreads, collapsedThreadsExtended, userID)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, bool, bool, string) *model.PostList); ok {
		r0 = rf(ctx, id, skipFetchThreads, collapsedThreads, collapsedThreadsExtended, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, bool, bool, string) error); ok {
		r1 = rf(ctx, id, skipFetchThreads, collapsedThreads, collapsedThreadsExtended, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectPostParentsForExportAfter provides a mock function with given fields: limit, afterID
func (_m *PostStore) GetDirectPostParentsForExportAfter(limit int, afterID string) ([]*model.DirectPostForExport, error) {
	ret := _m.Called(limit, afterID)

	var r0 []*model.DirectPostForExport
	if rf, ok := ret.Get(0).(func(int, string) []*model.DirectPostForExport); ok {
		r0 = rf(limit, afterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DirectPostForExport)
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

// GetEtag provides a mock function with given fields: channelID, allowFromCache, collapsedThreads
func (_m *PostStore) GetEtag(channelID string, allowFromCache bool, collapsedThreads bool) string {
	ret := _m.Called(channelID, allowFromCache, collapsedThreads)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, bool, bool) string); ok {
		r0 = rf(channelID, allowFromCache, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFlaggedPosts provides a mock function with given fields: userID, offset, limit
func (_m *PostStore) GetFlaggedPosts(userID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, offset, limit)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(string, int, int) *model.PostList); ok {
		r0 = rf(userID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(userID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlaggedPostsForChannel provides a mock function with given fields: userID, channelID, offset, limit
func (_m *PostStore) GetFlaggedPostsForChannel(userID string, channelID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, channelID, offset, limit)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(string, string, int, int) *model.PostList); ok {
		r0 = rf(userID, channelID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(userID, channelID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlaggedPostsForTeam provides a mock function with given fields: userID, teamID, offset, limit
func (_m *PostStore) GetFlaggedPostsForTeam(userID string, teamID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, teamID, offset, limit)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(string, string, int, int) *model.PostList); ok {
		r0 = rf(userID, teamID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(userID, teamID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxPostSize provides a mock function with given fields:
func (_m *PostStore) GetMaxPostSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetOldest provides a mock function with given fields:
func (_m *PostStore) GetOldest() (*model.Post, error) {
	ret := _m.Called()

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func() *model.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
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

// GetOldestEntityCreationTime provides a mock function with given fields:
func (_m *PostStore) GetOldestEntityCreationTime() (int64, error) {
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

// GetParentsForExportAfter provides a mock function with given fields: limit, afterID
func (_m *PostStore) GetParentsForExportAfter(limit int, afterID string) ([]*model.PostForExport, error) {
	ret := _m.Called(limit, afterID)

	var r0 []*model.PostForExport
	if rf, ok := ret.Get(0).(func(int, string) []*model.PostForExport); ok {
		r0 = rf(limit, afterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostForExport)
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

// GetPostAfterTime provides a mock function with given fields: channelID, time, collapsedThreads
func (_m *PostStore) GetPostAfterTime(channelID string, time int64, collapsedThreads bool) (*model.Post, error) {
	ret := _m.Called(channelID, time, collapsedThreads)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(string, int64, bool) *model.Post); ok {
		r0 = rf(channelID, time, collapsedThreads)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, time, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostIdAfterTime provides a mock function with given fields: channelID, time, collapsedThreads
func (_m *PostStore) GetPostIdAfterTime(channelID string, time int64, collapsedThreads bool) (string, error) {
	ret := _m.Called(channelID, time, collapsedThreads)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int64, bool) string); ok {
		r0 = rf(channelID, time, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, time, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostIdBeforeTime provides a mock function with given fields: channelID, time, collapsedThreads
func (_m *PostStore) GetPostIdBeforeTime(channelID string, time int64, collapsedThreads bool) (string, error) {
	ret := _m.Called(channelID, time, collapsedThreads)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int64, bool) string); ok {
		r0 = rf(channelID, time, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, time, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with given fields: options, allowFromCache
func (_m *PostStore) GetPosts(options model.GetPostsOptions, allowFromCache bool) (*model.PostList, error) {
	ret := _m.Called(options, allowFromCache)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, bool) *model.PostList); ok {
		r0 = rf(options, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPostsOptions, bool) error); ok {
		r1 = rf(options, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsAfter provides a mock function with given fields: options
func (_m *PostStore) GetPostsAfter(options model.GetPostsOptions) (*model.PostList, error) {
	ret := _m.Called(options)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions) *model.PostList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPostsOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsBatchForIndexing provides a mock function with given fields: startTime, endTime, limit
func (_m *PostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.PostForIndexing, error) {
	ret := _m.Called(startTime, endTime, limit)

	var r0 []*model.PostForIndexing
	if rf, ok := ret.Get(0).(func(int64, int64, int) []*model.PostForIndexing); ok {
		r0 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostForIndexing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(startTime, endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsBefore provides a mock function with given fields: options
func (_m *PostStore) GetPostsBefore(options model.GetPostsOptions) (*model.PostList, error) {
	ret := _m.Called(options)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions) *model.PostList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPostsOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByIds provides a mock function with given fields: postIds
func (_m *PostStore) GetPostsByIds(postIds []string) ([]*model.Post, error) {
	ret := _m.Called(postIds)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func([]string) []*model.Post); ok {
		r0 = rf(postIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(postIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsCreatedAt provides a mock function with given fields: channelID, time
func (_m *PostStore) GetPostsCreatedAt(channelID string, time int64) ([]*model.Post, error) {
	ret := _m.Called(channelID, time)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func(string, int64) []*model.Post); ok {
		r0 = rf(channelID, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(channelID, time)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsSince provides a mock function with given fields: options, allowFromCache
func (_m *PostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool) (*model.PostList, error) {
	ret := _m.Called(options, allowFromCache)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, bool) *model.PostList); ok {
		r0 = rf(options, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPostsSinceOptions, bool) error); ok {
		r1 = rf(options, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsSinceForSync provides a mock function with given fields: options, cursor
func (_m *PostStore) GetPostsSinceForSync(options model.GetPostsSinceForSyncOptions, cursor model.GetPostsSinceForSyncCursor) ([]*model.Post, model.GetPostsSinceForSyncCursor, error) {
	ret := _m.Called(options, cursor)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor) []*model.Post); ok {
		r0 = rf(options, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 model.GetPostsSinceForSyncCursor
	if rf, ok := ret.Get(1).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor) model.GetPostsSinceForSyncCursor); ok {
		r1 = rf(options, cursor)
	} else {
		r1 = ret.Get(1).(model.GetPostsSinceForSyncCursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor) error); ok {
		r2 = rf(options, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRepliesForExport provides a mock function with given fields: parentID
func (_m *PostStore) GetRepliesForExport(parentID string) ([]*model.ReplyForExport, error) {
	ret := _m.Called(parentID)

	var r0 []*model.ReplyForExport
	if rf, ok := ret.Get(0).(func(string) []*model.ReplyForExport); ok {
		r0 = rf(parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ReplyForExport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(parentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingle provides a mock function with given fields: id, inclDeleted
func (_m *PostStore) GetSingle(id string, inclDeleted bool) (*model.Post, error) {
	ret := _m.Called(id, inclDeleted)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(string, bool) *model.Post); ok {
		r0 = rf(id, inclDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(id, inclDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAutoResponsePostByUserSince provides a mock function with given fields: options, userId
func (_m *PostStore) HasAutoResponsePostByUserSince(options model.GetPostsSinceOptions, userId string) (bool, error) {
	ret := _m.Called(options, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, string) bool); ok {
		r0 = rf(options, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPostsSinceOptions, string) error); ok {
		r1 = rf(options, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvalidateLastPostTimeCache provides a mock function with given fields: channelID
func (_m *PostStore) InvalidateLastPostTimeCache(channelID string) {
	_m.Called(channelID)
}

// Overwrite provides a mock function with given fields: post
func (_m *PostStore) Overwrite(post *model.Post) (*model.Post, error) {
	ret := _m.Called(post)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(*model.Post) *model.Post); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OverwriteMultiple provides a mock function with given fields: posts
func (_m *PostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	ret := _m.Called(posts)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func([]*model.Post) []*model.Post); ok {
		r0 = rf(posts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func([]*model.Post) int); ok {
		r1 = rf(posts)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]*model.Post) error); ok {
		r2 = rf(posts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *PostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByChannel provides a mock function with given fields: channelID
func (_m *PostStore) PermanentDeleteByChannel(channelID string) error {
	ret := _m.Called(channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userID
func (_m *PostStore) PermanentDeleteByUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: post
func (_m *PostStore) Save(post *model.Post) (*model.Post, error) {
	ret := _m.Called(post)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(*model.Post) *model.Post); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMultiple provides a mock function with given fields: posts
func (_m *PostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	ret := _m.Called(posts)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func([]*model.Post) []*model.Post); ok {
		r0 = rf(posts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func([]*model.Post) int); ok {
		r1 = rf(posts)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]*model.Post) error); ok {
		r2 = rf(posts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Search provides a mock function with given fields: teamID, userID, params
func (_m *PostStore) Search(teamID string, userID string, params *model.SearchParams) (*model.PostList, error) {
	ret := _m.Called(teamID, userID, params)

	var r0 *model.PostList
	if rf, ok := ret.Get(0).(func(string, string, *model.SearchParams) *model.PostList); ok {
		r0 = rf(teamID, userID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.SearchParams) error); ok {
		r1 = rf(teamID, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPostsInTeamForUser provides a mock function with given fields: paramsList, userID, teamID, page, perPage
func (_m *PostStore) SearchPostsInTeamForUser(paramsList []*model.SearchParams, userID string, teamID string, page int, perPage int) (*model.PostSearchResults, error) {
	ret := _m.Called(paramsList, userID, teamID, page, perPage)

	var r0 *model.PostSearchResults
	if rf, ok := ret.Get(0).(func([]*model.SearchParams, string, string, int, int) *model.PostSearchResults); ok {
		r0 = rf(paramsList, userID, teamID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostSearchResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*model.SearchParams, string, string, int, int) error); ok {
		r1 = rf(paramsList, userID, teamID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newPost, oldPost
func (_m *PostStore) Update(newPost *model.Post, oldPost *model.Post) (*model.Post, error) {
	ret := _m.Called(newPost, oldPost)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(*model.Post, *model.Post) *model.Post); ok {
		r0 = rf(newPost, oldPost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Post, *model.Post) error); ok {
		r1 = rf(newPost, oldPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
