package app

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
)

const (
	mentionableUserPrefix  = "user:"
	mentionableGroupPrefix = "group:"
)

type MentionableID string

func MentionableUserID(userID string) MentionableID {
	return MentionableID(fmt.Sprint(mentionableUserPrefix, userID))
}

func MentionableGroupID(groupID string) MentionableID {
	return MentionableID(fmt.Sprint(mentionableGroupPrefix, groupID))
}

func (id MentionableID) AsUserID() (userID string, ok bool) {
	idString := string(id)
	if strings.HasPrefix(idString, mentionableUserPrefix) {
		return idString[len(mentionableUserPrefix):], true
	} else {
		return "", false
	}
}

func (id MentionableID) AsGroupID() (groupID string, ok bool) {
	idString := string(id)
	if strings.HasPrefix(idString, mentionableGroupPrefix) {
		return idString[len(mentionableGroupPrefix):], true
	} else {
		return "", false
	}
}

type MentionKeywords map[string][]MentionableID

func (k MentionKeywords) AddUser(profile *model.User, channelNotifyProps map[string]string, status *model.Status, allowChannelMentions bool) MentionKeywords {
	mentionableID := MentionableUserID(profile.Id)

	userMention := "@" + strings.ToLower(profile.Username)
	k[userMention] = append(k[userMention], mentionableID)

	// Add all the user's mention keys
	for _, mentionKey := range profile.GetMentionKeys() {
		// note that these are made lower case so that we can do a case insensitive check for them
		mentionKey = strings.ToLower(mentionKey)

		if mentionKey != "" {
			k[mentionKey] = append(k[mentionKey], mentionableID)
		}
	}

	// If turned on, add the user's case sensitive first name
	if profile.NotifyProps[model.FirstNameNotifyProp] == "true" && profile.FirstName != "" {
		k[profile.FirstName] = append(k[profile.FirstName], mentionableID)
	}

	// Add @channel and @all to k if user has them turned on and the server allows them
	if allowChannelMentions {
		// Ignore channel mentions if channel is muted and channel mention setting is default
		ignoreChannelMentions := channelNotifyProps[model.IgnoreChannelMentionsNotifyProp] == model.IgnoreChannelMentionsOn || (channelNotifyProps[model.MarkUnreadNotifyProp] == model.UserNotifyMention && channelNotifyProps[model.IgnoreChannelMentionsNotifyProp] == model.IgnoreChannelMentionsDefault)

		if profile.NotifyProps[model.ChannelMentionsNotifyProp] == "true" && !ignoreChannelMentions {
			k["@channel"] = append(k["@channel"], mentionableID)
			k["@all"] = append(k["@all"], mentionableID)

			if status != nil && status.Status == model.StatusOnline {
				k["@here"] = append(k["@here"], mentionableID)
			}
		}
	}

	return k
}

func (k MentionKeywords) AddUserID(userID string, keyword string) MentionKeywords {
	k[keyword] = append(k[keyword], MentionableUserID(userID))

	return k
}

func (k MentionKeywords) AddGroup(group *model.Group) MentionKeywords {
	if group.Name != nil {
		keyword := "@" + *group.Name
		k[keyword] = append(k[keyword], MentionableGroupID(group.Id))
	}

	return k
}

func (k MentionKeywords) AddGroupsMap(groups map[string]*model.Group) MentionKeywords {
	for _, group := range groups {
		k.AddGroup(group)
	}

	return k
}
