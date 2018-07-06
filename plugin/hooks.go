// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package plugin

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

// These assignments are part of the wire protocol. You can add more, but should not change existing
// assignments. Follow the naming convention of <HookName>Id as the autogenerated glue code depends on that.
const (
	OnActivateId            = 0
	OnDeactivateId          = 1
	ServeHTTPId             = 2
	OnConfigurationChangeId = 3
	ExecuteCommandId        = 4
	MessageWillBePostedId   = 5
	MessageWillBeUpdatedId  = 6
	MessageHasBeenPostedId  = 7
	MessageHasBeenUpdatedId = 8
	TotalHooksId            = iota
)

// Methods from the Hooks interface can be used by a plugin to respond to events. Methods are likely
// to be added over time, and plugins are not expected to implement all of them. Instead, plugins
// are expected to implement a subset of them and pass an instance to plugin/rpcplugin.Main, which
// will take over execution of the process and add default behaviors for missing hooks.
type Hooks interface {
	// OnActivate is invoked when the plugin is activated.
	OnActivate() error

	// Implemented returns a list of hooks that are implmented by the plugin.
	// Plugins do not need to provide an implementation. Any given will be ignored.
	Implemented() ([]string, error)

	// OnDeactivate is invoked when the plugin is deactivated. This is the plugin's last chance to
	// use the API, and the plugin will be terminated shortly after this invocation.
	OnDeactivate() error

	// OnConfigurationChange is invoked when configuration changes may have been made.
	OnConfigurationChange() error

	// ServeHTTP allows the plugin to implement the http.Handler interface. Requests destined for
	// the /plugins/{id} path will be routed to the plugin.
	//
	// The Mattermost-User-Id header will be present if (and only if) the request is by an
	// authenticated user.
	ServeHTTP(c *Context, w http.ResponseWriter, r *http.Request)

	// ExecuteCommand executes a command that has been previously registered via the RegisterCommand
	// API.
	ExecuteCommand(c *Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError)

	// MessageWillBePosted is invoked when a message is posted by a user before it is commited
	// to the database. If you also want to act on edited posts, see MessageWillBeUpdated.
	// Return values should be the modified post or nil if rejected and an explanation for the user.
	//
	// If you don't need to modify or reject posts, use MessageHasBeenPosted instead.
	//
	// Note that this method will be called for posts created by plugins, including the plugin that
	// created the post.
	MessageWillBePosted(c *Context, post *model.Post) (*model.Post, string)

	// MessageWillBeUpdated is invoked when a message is updated by a user before it is commited
	// to the database. If you also want to act on new posts, see MessageWillBePosted.
	// Return values should be the modified post or nil if rejected and an explanation for the user.
	// On rejection, the post will be kept in its previous state.
	//
	// If you don't need to modify or rejected updated posts, use MessageHasBeenUpdated instead.
	//
	// Note that this method will be called for posts updated by plugins, including the plugin that
	// updated the post.
	MessageWillBeUpdated(c *Context, newPost, oldPost *model.Post) (*model.Post, string)

	// MessageHasBeenPosted is invoked after the message has been commited to the databse.
	// If you need to modify or reject the post, see MessageWillBePosted
	// Note that this method will be called for posts created by plugins, including the plugin that
	// created the post.
	MessageHasBeenPosted(c *Context, post *model.Post)

	// MessageHasBeenUpdated is invoked after a message is updated and has been updated in the databse.
	// If you need to modify or reject the post, see MessageWillBeUpdated
	// Note that this method will be called for posts created by plugins, including the plugin that
	// created the post.
	MessageHasBeenUpdated(c *Context, newPost, oldPost *model.Post)
}
