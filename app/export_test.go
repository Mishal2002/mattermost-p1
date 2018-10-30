package app

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReactionsOfPost(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	post := th.BasicPost
	post.HasReactions = true

	reactionObject := model.Reaction{
		UserId:    model.NewId(),
		PostId:    post.Id,
		EmojiName: "emoji",
		CreateAt:  model.GetMillis(),
	}

	th.App.SaveReactionForPost(&reactionObject)
	reactionsOfPost, err := th.App.BuildPostReactions(post.Id)

	if err != nil {
		t.Fatal("should have reactions")
	}

	assert.Equal(t, reactionObject.EmojiName, *(*reactionsOfPost)[0].EmojiName)
}

func TestExportUserNotifyProps(t *testing.T) {

	th := Setup().InitBasic()
	defer th.TearDown()

	userNotifyProps := model.StringMap{
		model.DESKTOP_NOTIFY_PROP:            model.USER_NOTIFY_ALL,
		model.DESKTOP_SOUND_NOTIFY_PROP:      "true",
		model.EMAIL_NOTIFY_PROP:              "true",
		model.MOBILE_NOTIFY_PROP:             model.USER_NOTIFY_ALL,
		model.MOBILE_PUSH_STATUS_NOTIFY_PROP: model.STATUS_ONLINE,
		model.CHANNEL_MENTIONS_NOTIFY_PROP:   "true",
		model.COMMENTS_NOTIFY_PROP:           model.COMMENTS_NOTIFY_ROOT,
		model.MENTION_KEYS_NOTIFY_PROP:       "valid,misc",
	}

	exportNotifyProps := th.App.buildUserNotifyProps(userNotifyProps)

	require.Equal(t, userNotifyProps[model.DESKTOP_NOTIFY_PROP], *exportNotifyProps.Desktop)
	require.Equal(t, userNotifyProps[model.DESKTOP_SOUND_NOTIFY_PROP], *exportNotifyProps.DesktopSound)
	require.Equal(t, userNotifyProps[model.EMAIL_NOTIFY_PROP], *exportNotifyProps.Email)
	require.Equal(t, userNotifyProps[model.MOBILE_NOTIFY_PROP], *exportNotifyProps.Mobile)
	require.Equal(t, userNotifyProps[model.MOBILE_PUSH_STATUS_NOTIFY_PROP], *exportNotifyProps.MobilePushStatus)
	require.Equal(t, userNotifyProps[model.CHANNEL_MENTIONS_NOTIFY_PROP], *exportNotifyProps.ChannelTrigger)
	require.Equal(t, userNotifyProps[model.COMMENTS_NOTIFY_PROP], *exportNotifyProps.CommentsTrigger)
	require.Equal(t, userNotifyProps[model.MENTION_KEYS_NOTIFY_PROP], *exportNotifyProps.MentionKeys)
}

func TestExportUserChannelsNotifyProps(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()
	channel := th.BasicChannel
	user := th.BasicUser
	team := th.BasicTeam
	channelName := channel.Name
	notifyProps := model.StringMap{
		model.DESKTOP_NOTIFY_PROP: model.USER_NOTIFY_ALL,
		model.PUSH_NOTIFY_PROP:    model.USER_NOTIFY_NONE,
	}
	channelMember := model.ChannelMember{
		ChannelId: channel.Id,
		UserId:    user.Id,
	}
	th.App.Srv.Store.Channel().SaveMember(&channelMember)
	th.App.UpdateChannelMemberNotifyProps(notifyProps, channel.Id, user.Id)
	exportData, _ := th.App.buildUserChannelMemberships(user.Id, team.Id)
	for _, data := range *exportData {
		if *data.Name == channelName {
			assert.Equal(t, *data.NotifyProps.Desktop, "all")
			assert.Equal(t, *data.NotifyProps.Mobile, "none")
			assert.Equal(t, *data.NotifyProps.MarkUnread, "all")
		} else { // default values
			assert.Equal(t, *data.NotifyProps.Desktop, "default")
			assert.Equal(t, *data.NotifyProps.Mobile, "default")
			assert.Equal(t, *data.NotifyProps.MarkUnread, "all")
		}
	}
}
