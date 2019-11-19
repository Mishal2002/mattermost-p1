// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommandWebhookPreSave(t *testing.T) {
	h := CommandWebhook{}
	h.PreSave()

	require.Len(t, h.Id, 26, "Id should be generated")
	require.NotEqual(t, 0, h.CreateAt, "CreateAt should be set")
}

func TestCommandWebhookIsValid(t *testing.T) {
	h := CommandWebhook{}
	h.Id = NewId()
	h.CreateAt = GetMillis()
	h.CommandId = NewId()
	h.UserId = NewId()
	h.ChannelId = NewId()

	for _, test := range []struct {
		Transform     func()
		ExpectedError string
	}{
		{func() {}, ""},
		{func() { h.Id = "asd" }, "model.command_hook.id.app_error"},
		{func() { h.CreateAt = 0 }, "model.command_hook.create_at.app_error"},
		{func() { h.CommandId = "asd" }, "model.command_hook.command_id.app_error"},
		{func() { h.UserId = "asd" }, "model.command_hook.user_id.app_error"},
		{func() { h.ChannelId = "asd" }, "model.command_hook.channel_id.app_error"},
		{func() { h.RootId = "asd" }, "model.command_hook.root_id.app_error"},
		{func() { h.RootId = NewId() }, ""},
		{func() { h.ParentId = "asd" }, "model.command_hook.parent_id.app_error"},
		{func() { h.ParentId = NewId() }, ""},
	} {
		tmp := h
		test.Transform()
		err := h.IsValid()

		if test.ExpectedError == "" {
			assert.Error(t, err, "hook should be valid")
		} else {
			require.NotNil(t, err)
			assert.Equal(t, test.ExpectedError, err.Id, "expected "+test.ExpectedError+" error")
		}

		h = tmp
	}
}
