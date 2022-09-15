// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import "encoding/json"

type MessageExport struct {
	TeamId          *string
	TeamName        *string
	TeamDisplayName *string

	ChannelId          *string
	ChannelName        *string
	ChannelDisplayName *string
	ChannelType        *ChannelType

	UserId    *string
	UserEmail *string
	Username  *string
	IsBot     bool

	PostId         *string
	PostCreateAt   *int64
	PostUpdateAt   *int64
	PostDeleteAt   *int64
	PostMessage    *string
	PostType       *string
	PostRootId     *string
	PostProps      *string
	PostOriginalId *string
	PostFileIds    StringArray
}

type BlockExport struct {
	ID                 string
	ParentID           string
	RootID             string
	ModifiedBy         string
	ModifiedByEmail    string
	ModifiedByUsername string
	Type               string
	Title              string
	Fields             string
	CreateAt           int64
	UpdateAt           int64
	DeleteAt           int64
	WorkspaceID        string
}

type MessageExportCursor struct {
	LastPostUpdateAt int64
	LastPostId       string
}

type BlockExportCursor struct {
	LastBlockUpdateAt int64
	LastBlockId       string
}

// PreviewID returns the value of the post's previewed_post prop, if present, or an empty string.
func (m *MessageExport) PreviewID() string {
	var previewID string
	props := map[string]any{}
	if m.PostProps != nil && json.Unmarshal([]byte(*m.PostProps), &props) == nil {
		if val, ok := props[PostPropsPreviewedPost]; ok {
			previewID = val.(string)
		}
	}
	return previewID
}
