// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

func ImportLineFromTeam(team *model.TeamForExport) *LineImportData {
	return &LineImportData{
		Type: "team",
		Team: &TeamImportData{
			Name:            &team.Name,
			DisplayName:     &team.DisplayName,
			Type:            &team.Type,
			Description:     &team.Description,
			AllowOpenInvite: &team.AllowOpenInvite,
			Scheme:          team.SchemeName,
		},
	}
}

func ImportLineFromChannel(channel *model.ChannelForExport) *LineImportData {
	return &LineImportData{
		Type: "channel",
		Channel: &ChannelImportData{
			Team:        &channel.TeamName,
			Name:        &channel.Name,
			DisplayName: &channel.DisplayName,
			Type:        &channel.Type,
			Header:      &channel.Header,
			Purpose:     &channel.Purpose,
			Scheme:      channel.SchemeName,
		},
	}
}

func ImportLineFromUser(user *model.User) *LineImportData {
	return &LineImportData{
		Type: "user",
		User: &UserImportData{
			Username:    &user.Username,
			Email:       &user.Email,
			AuthService: &user.AuthService,
			AuthData:    user.AuthData,
			Nickname:    &user.Nickname,
			FirstName:   &user.FirstName,
			LastName:    &user.LastName,
			Position:    &user.Position,
			Roles:       &user.Roles,
			Locale:      &user.Locale,
		},
	}
}

func ImportUserTeamDataFromTeamMember(member *model.TeamMemberForExport) *UserTeamImportData {
	rolesList := strings.Fields(member.Roles)
	if member.SchemeAdmin {
		rolesList = append(rolesList, model.TEAM_ADMIN_ROLE_ID)
	}
	if member.SchemeUser {
		rolesList = append(rolesList, model.TEAM_USER_ROLE_ID)
	}
	roles := strings.Join(rolesList, " ")
	return &UserTeamImportData{
		Name:  &member.TeamName,
		Roles: &roles,
	}
}

func ImportUserChannelDataFromChannelMember(member *model.ChannelMemberForExport) *UserChannelImportData {
	rolesList := strings.Fields(member.Roles)
	if member.SchemeAdmin {
		rolesList = append(rolesList, model.CHANNEL_ADMIN_ROLE_ID)
	}
	if member.SchemeUser {
		rolesList = append(rolesList, model.CHANNEL_USER_ROLE_ID)
	}
	props := member.NotifyProps
	desktop := props[model.DESKTOP_NOTIFY_PROP]
	mobile := props[model.PUSH_NOTIFY_PROP]
	markUnread := props[model.MARK_UNREAD_NOTIFY_PROP]
	notifyProps := UserChannelNotifyPropsImportData{
		Desktop:    &desktop,
		Mobile:     &mobile,
		MarkUnread: &markUnread,
	}
	roles := strings.Join(rolesList, " ")
	return &UserChannelImportData{
		Name:        &member.ChannelName,
		Roles:       &roles,
		NotifyProps: &notifyProps,
	}
}

func ImportLineForPost(post *model.PostForExport) *LineImportData {
	return &LineImportData{
		Type: "post",
		Post: &PostImportData{
			Team:     &post.TeamName,
			Channel:  &post.ChannelName,
			User:     &post.Username,
			Message:  &post.Message,
			CreateAt: &post.CreateAt,
		},
	}
}

func ImportReplyFromPost(post *model.ReplyForExport) *ReplyImportData {
	return &ReplyImportData{
		User:     &post.Username,
		Message:  &post.Message,
		CreateAt: &post.CreateAt,
	}
}

func ImportReactionFromPost(reaction *model.Reaction) *ReactionImportData {
	return &ReactionImportData{
		User:      &reaction.UserId,
		EmojiName: &reaction.EmojiName,
		CreateAt:  &reaction.CreateAt,
	}
}
