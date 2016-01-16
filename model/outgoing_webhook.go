// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"encoding/json"
	"fmt"
	"io"
)

type OutgoingWebhook struct {
	Id           string      `json:"id"`
	Token        string      `json:"token"`
	CreateAt     int64       `json:"create_at"`
	UpdateAt     int64       `json:"update_at"`
	DeleteAt     int64       `json:"delete_at"`
	CreatorId    string      `json:"creator_id"`
	ChannelId    string      `json:"channel_id"`
	TeamId       string      `json:"team_id"`
	TriggerWords StringArray `json:"trigger_words"`
	CallbackURLs StringArray `json:"callback_urls"`
}

func (o *OutgoingWebhook) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func OutgoingWebhookFromJson(data io.Reader) *OutgoingWebhook {
	decoder := json.NewDecoder(data)
	var o OutgoingWebhook
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func OutgoingWebhookListToJson(l []*OutgoingWebhook) string {
	b, err := json.Marshal(l)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func OutgoingWebhookListFromJson(data io.Reader) []*OutgoingWebhook {
	decoder := json.NewDecoder(data)
	var o []*OutgoingWebhook
	err := decoder.Decode(&o)
	if err == nil {
		return o
	} else {
		return nil
	}
}

func (o *OutgoingWebhook) IsValid(T goi18n.TranslateFunc) *AppError {

	if len(o.Id) != 26 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid Id"), "")
	}

	if len(o.Token) != 26 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid token"), "")
	}

	if o.CreateAt == 0 {
		return NewAppError("OutgoingWebhook.IsValid", T("Create at must be a valid time"), "id="+o.Id)
	}

	if o.UpdateAt == 0 {
		return NewAppError("OutgoingWebhook.IsValid", T("Update at must be a valid time"), "id="+o.Id)
	}

	if len(o.CreatorId) != 26 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid user id"), "")
	}

	if len(o.ChannelId) != 0 && len(o.ChannelId) != 26 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid channel id"), "")
	}

	if len(o.TeamId) != 26 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid team id"), "")
	}

	if len(fmt.Sprintf("%s", o.TriggerWords)) > 1024 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid trigger words"), "")
	}

	if len(o.CallbackURLs) == 0 || len(fmt.Sprintf("%s", o.CallbackURLs)) > 1024 {
		return NewAppError("OutgoingWebhook.IsValid", T("Invalid callback urls"), "")
	}

	for _, callback := range o.CallbackURLs {
		if !IsValidHttpUrl(callback) {
			return NewAppError("OutgoingWebhook.IsValid", T("Invalid callback URLs. Each must be a valid URL and start with http:// or https://"), "")
		}
	}

	return nil
}

func (o *OutgoingWebhook) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	if o.Token == "" {
		o.Token = NewId()
	}

	o.CreateAt = GetMillis()
	o.UpdateAt = o.CreateAt
}

func (o *OutgoingWebhook) PreUpdate() {
	o.UpdateAt = GetMillis()
}

func (o *OutgoingWebhook) HasTriggerWord(word string) bool {
	if len(o.TriggerWords) == 0 || len(word) == 0 {
		return false
	}

	for _, trigger := range o.TriggerWords {
		if trigger == word {
			return true
		}
	}

	return false
}
