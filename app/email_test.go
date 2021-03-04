// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/services/mailservice"
)

func TestCondenseSiteURL(t *testing.T) {
	require.Equal(t, "", condenseSiteURL(""))
	require.Equal(t, "mattermost.com", condenseSiteURL("mattermost.com"))
	require.Equal(t, "mattermost.com", condenseSiteURL("mattermost.com/"))
	require.Equal(t, "chat.mattermost.com", condenseSiteURL("chat.mattermost.com"))
	require.Equal(t, "chat.mattermost.com", condenseSiteURL("chat.mattermost.com/"))
	require.Equal(t, "mattermost.com/subpath", condenseSiteURL("mattermost.com/subpath"))
	require.Equal(t, "mattermost.com/subpath", condenseSiteURL("mattermost.com/subpath/"))
	require.Equal(t, "chat.mattermost.com/subpath", condenseSiteURL("chat.mattermost.com/subpath"))
	require.Equal(t, "chat.mattermost.com/subpath", condenseSiteURL("chat.mattermost.com/subpath/"))

	require.Equal(t, "mattermost.com:8080", condenseSiteURL("http://mattermost.com:8080"))
	require.Equal(t, "mattermost.com:8080", condenseSiteURL("http://mattermost.com:8080/"))
	require.Equal(t, "chat.mattermost.com:8080", condenseSiteURL("http://chat.mattermost.com:8080"))
	require.Equal(t, "chat.mattermost.com:8080", condenseSiteURL("http://chat.mattermost.com:8080/"))
	require.Equal(t, "mattermost.com:8080/subpath", condenseSiteURL("http://mattermost.com:8080/subpath"))
	require.Equal(t, "mattermost.com:8080/subpath", condenseSiteURL("http://mattermost.com:8080/subpath/"))
	require.Equal(t, "chat.mattermost.com:8080/subpath", condenseSiteURL("http://chat.mattermost.com:8080/subpath"))
	require.Equal(t, "chat.mattermost.com:8080/subpath", condenseSiteURL("http://chat.mattermost.com:8080/subpath/"))
}

func TestSendInviteEmailRateLimits(t *testing.T) {
	th := Setup(t).InitBasic()
	defer th.TearDown()

	th.BasicTeam.AllowedDomains = "common.com"
	_, err := th.App.UpdateTeam(th.BasicTeam)
	require.Nilf(t, err, "%v, Should update the team", err)

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.EnableEmailInvitations = true
	})

	emailList := make([]string, 22)
	for i := 0; i < 22; i++ {
		emailList[i] = "test-" + strconv.Itoa(i) + "@common.com"
	}
	err = th.App.InviteNewUsersToTeam(emailList, th.BasicTeam.Id, th.BasicUser.Id)
	require.NotNil(t, err)
	assert.Equal(t, "app.email.rate_limit_exceeded.app_error", err.Id)
	assert.Equal(t, http.StatusRequestEntityTooLarge, err.StatusCode)

	_, err = th.App.InviteNewUsersToTeamGracefully(emailList, th.BasicTeam.Id, th.BasicUser.Id)
	require.NotNil(t, err)
	assert.Equal(t, "app.email.rate_limit_exceeded.app_error", err.Id)
	assert.Equal(t, http.StatusRequestEntityTooLarge, err.StatusCode)
}

func TestSendAdminUpgradeRequestEmail(t *testing.T) {
	th := Setup(t).InitBasic()
	defer th.TearDown()

	th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

	mockSubscription := &model.Subscription{
		ID:         "MySubscriptionID",
		CustomerID: "MyCustomer",
		ProductID:  "SomeProductId",
		AddOns:     []string{},
		StartAt:    1000000000,
		EndAt:      2000000000,
		CreateAt:   1000000000,
		Seats:      100,
		DNS:        "some.dns.server",
		IsPaidTier: "false",
	}

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ExperimentalSettings.CloudUserLimit = 10
	})

	err := th.App.SendAdminUpgradeRequestEmail(th.BasicUser.Username, mockSubscription, model.InviteLimitation)
	require.Nil(t, err)

	// other attempts by the same user or other users to send emails are blocked by rate limiter
	err = th.App.SendAdminUpgradeRequestEmail(th.BasicUser.Username, mockSubscription, model.InviteLimitation)
	require.NotNil(t, err)
	assert.Equal(t, err.Id, "app.email.rate_limit_exceeded.app_error")

	err = th.App.SendAdminUpgradeRequestEmail(th.BasicUser2.Username, mockSubscription, model.InviteLimitation)
	require.NotNil(t, err)
	assert.Equal(t, err.Id, "app.email.rate_limit_exceeded.app_error")
}

func TestSendAdminUpgradeRequestEmailOnJoin(t *testing.T) {
	th := Setup(t).InitBasic()
	defer th.TearDown()

	th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

	mockSubscription := &model.Subscription{
		ID:         "MySubscriptionID",
		CustomerID: "MyCustomer",
		ProductID:  "SomeProductId",
		AddOns:     []string{},
		StartAt:    1000000000,
		EndAt:      2000000000,
		CreateAt:   1000000000,
		Seats:      100,
		DNS:        "some.dns.server",
		IsPaidTier: "false",
	}

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ExperimentalSettings.CloudUserLimit = 10
	})

	err := th.App.SendAdminUpgradeRequestEmail(th.BasicUser.Username, mockSubscription, model.JoinLimitation)
	require.Nil(t, err)

	// other attempts by the same user or other users to send emails are blocked by rate limiter
	err = th.App.SendAdminUpgradeRequestEmail(th.BasicUser.Username, mockSubscription, model.JoinLimitation)
	require.NotNil(t, err)
	assert.Equal(t, err.Id, "app.email.rate_limit_exceeded.app_error")

	err = th.App.SendAdminUpgradeRequestEmail(th.BasicUser2.Username, mockSubscription, model.JoinLimitation)
	require.NotNil(t, err)
	assert.Equal(t, err.Id, "app.email.rate_limit_exceeded.app_error")
}

func TestSendInviteEmails(t *testing.T) {
	th := Setup(t).InitBasic()
	defer th.TearDown()

	inbucket_host := os.Getenv("CI_INBUCKET_HOST")
	if inbucket_host == "" {
		inbucket_host = "localhost"
	}
	inbucket_port := os.Getenv("CI_INBUCKET_SMTP_PORT")
	if inbucket_port == "" {
		inbucket_port = "10025"
	}
	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.EnableEmailInvitations = true
		*cfg.EmailSettings.SMTPServer = inbucket_host
		*cfg.EmailSettings.SMTPPort = inbucket_port
	})

	emailTo := "test@example.com"
	mailservice.DeleteMailBox(emailTo)

	err := th.App.Srv().EmailService.SendInviteEmails(th.BasicTeam, "test-user", th.BasicUser.Id, []string{emailTo}, "http://testserver")
	require.Nil(t, err)

	var resultsMailbox mailservice.JSONMessageHeaderInbucket
	err3 := mailservice.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = mailservice.GetMailBox(emailTo)
		return err
	})
	if err3 != nil {
		t.Log(err3)
		t.Log("No email was received, maybe due load on the server. Skipping this verification")
	} else {
		if len(resultsMailbox) > 0 {
			require.Contains(t, resultsMailbox[0].To[0], emailTo, "Wrong To: recipient")
			resultsEmail, err := mailservice.GetMessageFromMailbox(emailTo, resultsMailbox[0].ID)
			require.NoError(t, err, "Could not get message from mailbox")
			require.Contains(t, resultsEmail.Body.HTML, "http://testserver", "Wrong received message %s", resultsEmail.Body.Text)
			require.Contains(t, resultsEmail.Body.HTML, "test-user", "Wrong received message %s", resultsEmail.Body.Text)
			require.Contains(t, resultsEmail.Body.Text, "http://testserver", "Wrong received message %s", resultsEmail.Body.Text)
			require.Contains(t, resultsEmail.Body.Text, "test-user", "Wrong received message %s", resultsEmail.Body.Text)
		}
	}
}
