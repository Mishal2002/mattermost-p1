// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Integration Action Flow
//
// 1. An integration creates an interactive message button or menu.
// 2. A user clicks on a button or selects an option from the menu.
// 3. The client sends a request to server to complete the post action, calling DoPostAction below.
// 4. DoPostAction will send an HTTP POST request to the integration containing contextual data, including
// an encoded and signed trigger ID. Slash commands also include trigger IDs in their payloads.
// 5. The integration performs any actions it needs to and optionally makes a request back to the MM server
// using the trigger ID to open an interactive dialog.
// 6. If that optional request is made, OpenInteractiveDialog sends a WebSocket event to all connected clients
// for the relevant user, telling them to display the dialog.
// 7. The user fills in the dialog and submits it, where SubmitInteractiveDialog will submit it back to the
// integration for handling.

package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/store"
	"github.com/mattermost/mattermost-server/utils"
)

func (a *App) DoPostAction(postId, actionId, userId, selectedOption string) (string, *model.AppError) {
	return a.DoPostActionWithCookie(postId, actionId, userId, selectedOption, nil)
}

func (a *App) DoPostActionWithCookie(postId, actionId, userId, selectedOption string, cookie *model.PostActionCookie) (string, *model.AppError) {

	// PostAction may result in the original post being updated. For the
	// updated post, we need to unconditionally preserve the original
	// IsPinned and HasReaction attributes, and preserve its entire
	// original Props set unless the plugin returns a replacement value.
	// originalXxx variables are used to preserve these values.
	var originalProps map[string]interface{}
	originalIsPinned := false
	originalHasReactions := false

	// If the updated post does contain a replacement Props set, we still
	// need to preserve some original values, as listed in
	// model.PostActionRetainPropKeys. remove and retain track these.
	remove := []string{}
	retain := map[string]interface{}{}

	datasource := ""
	upstreamURL := ""
	rootPostId := ""
	upstreamRequest := &model.PostActionIntegrationRequest{
		UserId: userId,
		PostId: postId,
	}

	// See if the post exists in the DB, if so ignore the cookie.
	// Start all queries here for parallel execution
	pchan := make(chan store.StoreResult, 1)
	go func() {
		post, err := a.Srv.Store.Post().GetSingle(postId)
		pchan <- store.StoreResult{Data: post, Err: err}
		close(pchan)
	}()

	cchan := make(chan store.StoreResult, 1)
	go func() {
		channel, err := a.Srv.Store.Channel().GetForPost(postId)
		cchan <- store.StoreResult{Data: channel, Err: err}
		close(cchan)
	}()

	result := <-pchan
	if result.Err != nil {
		if cookie == nil {
			return "", result.Err
		}
		if cookie.Integration == nil {
			return "", model.NewAppError("DoPostAction", "api.post.do_action.action_integration.app_error", nil, "no Integration in action cookie", http.StatusBadRequest)
		}

		if postId != cookie.PostId {
			return "", model.NewAppError("DoPostAction", "api.post.do_action.action_integration.app_error", nil, "postId doesn't match", http.StatusBadRequest)
		}

		upstreamRequest.ChannelId = cookie.ChannelId
		upstreamRequest.Type = cookie.Type
		upstreamRequest.Context = cookie.Integration.Context
		datasource = cookie.DataSource

		retain = cookie.RetainProps
		remove = cookie.RemoveProps
		rootPostId = cookie.RootPostId
		upstreamURL = cookie.Integration.URL
	} else {
		post := result.Data.(*model.Post)
		result = <-cchan
		if result.Err != nil {
			return "", result.Err
		}
		channel := result.Data.(*model.Channel)

		action := post.GetAction(actionId)
		if action == nil || action.Integration == nil {
			return "", model.NewAppError("DoPostAction", "api.post.do_action.action_id.app_error", nil, fmt.Sprintf("action=%v", action), http.StatusNotFound)
		}

		upstreamRequest.ChannelId = post.ChannelId
		upstreamRequest.TeamId = channel.TeamId
		upstreamRequest.Type = action.Type
		upstreamRequest.Context = action.Integration.Context
		datasource = action.DataSource

		// Save the original values that may need to be preserved (including selected
		// Props, i.e. override_username, override_icon_url)
		for _, key := range model.PostActionRetainPropKeys {
			value, ok := post.Props[key]
			if ok {
				retain[key] = value
			} else {
				remove = append(remove, key)
			}
		}
		originalProps = post.Props
		originalIsPinned = post.IsPinned
		originalHasReactions = post.HasReactions

		if post.RootId == "" {
			rootPostId = post.Id
		} else {
			rootPostId = post.RootId
		}

		upstreamURL = action.Integration.URL
	}

	if upstreamRequest.Type == model.POST_ACTION_TYPE_SELECT {
		if selectedOption != "" {
			if upstreamRequest.Context == nil {
				upstreamRequest.Context = map[string]interface{}{}
			}
			upstreamRequest.DataSource = datasource
			upstreamRequest.Context["selected_option"] = selectedOption
		}
	}

	clientTriggerId, _, appErr := upstreamRequest.GenerateTriggerId(a.AsymmetricSigningKey())
	if appErr != nil {
		return "", appErr
	}

	resp, appErr := a.DoActionRequest(upstreamURL, upstreamRequest.ToJson())
	if appErr != nil {
		return "", appErr
	}
	defer resp.Body.Close()

	var response model.PostActionIntegrationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", model.NewAppError("DoPostAction", "api.post.do_action.action_integration.app_error", nil, "err="+err.Error(), http.StatusBadRequest)
	}

	if response.Update != nil {
		response.Update.Id = postId

		// Restore the post attributes and Props that need to be preserved
		if response.Update.Props == nil {
			response.Update.Props = originalProps
		} else {
			for key, value := range retain {
				response.Update.AddProp(key, value)
			}
			for _, key := range remove {
				delete(response.Update.Props, key)
			}
		}
		response.Update.IsPinned = originalIsPinned
		response.Update.HasReactions = originalHasReactions

		if _, appErr = a.UpdatePost(response.Update, false); appErr != nil {
			return "", appErr
		}
	}

	if response.EphemeralText != "" {
		ephemeralPost := &model.Post{
			Message:   model.ParseSlackLinksToMarkdown(response.EphemeralText),
			ChannelId: upstreamRequest.ChannelId,
			RootId:    rootPostId,
			UserId:    userId,
		}
		for key, value := range retain {
			ephemeralPost.AddProp(key, value)
		}
		a.SendEphemeralPost(userId, ephemeralPost)
	}

	return clientTriggerId, nil
}

// Perform an HTTP POST request to an integration's action endpoint.
// Caller must consume and close returned http.Response as necessary.
// For internal requests, requests are routed directly to a plugin ServerHTTP hook
func (a *App) DoActionRequest(rawURL string, body []byte) (*http.Response, *model.AppError) {
	inURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	siteURL, _ := url.Parse(*a.Config().ServiceSettings.SiteURL)
	rawURLPath := path.Clean(rawURL)
	if siteURL != nil && (strings.HasPrefix(rawURLPath, "/plugins/") || strings.HasPrefix(rawURLPath, "plugins/")) {
		inURL.Scheme = siteURL.Scheme
		inURL.Host = siteURL.Host
		inURL.Path = path.Join("/", siteURL.Path, rawURLPath)
		rawURL = inURL.String()
	}
	if strings.Contains(rawURL, "plugins/") {
		return a.DoLocalRequest(rawURL, body)
	}

	req, err := http.NewRequest("POST", rawURL, bytes.NewReader(body))
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, err.Error(), http.StatusBadRequest)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Allow access to plugin routes for action buttons
	var httpClient *http.Client
	subpath, _ := utils.GetSubpathFromConfig(a.Config())
	if (inURL.Hostname() == "localhost" || inURL.Hostname() == "127.0.0.1" || inURL.Hostname() == siteURL.Hostname()) && strings.HasPrefix(inURL.Path, path.Join(subpath, "plugins")) {
		req.Header.Set(model.HEADER_AUTH, "Bearer "+a.Session.Token)
		httpClient = a.HTTPService.MakeClient(true)
	} else {
		httpClient = a.HTTPService.MakeClient(false)
	}

	resp, httpErr := httpClient.Do(req)
	if httpErr != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err="+httpErr.Error(), http.StatusBadRequest)
	}

	if resp.StatusCode != http.StatusOK {
		return resp, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, fmt.Sprintf("status=%v", resp.StatusCode), http.StatusBadRequest)
	}

	return resp, nil
}

type LocalResponseWriter struct {
	data   []byte
	status int
}

func (w *LocalResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (w *LocalResponseWriter) Write(bytes []byte) (int, error) {
	w.data = make([]byte, len(bytes))
	copy(w.data, bytes)
	return len(w.data), nil
}

func (w *LocalResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (a *App) DoLocalRequest(rawURL string, body []byte) (*http.Response, *model.AppError) {
	inURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err="+err.Error(), http.StatusBadRequest)
	}
	pluginsEnvironment := a.GetPluginsEnvironment()

	regx := regexp.MustCompile(`plugins\/\s*([^\/]*)`)
	matches := regx.FindStringSubmatch(inURL.Path)
	if len(matches) != 2 {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err=Unable to find pluginId", http.StatusBadRequest)
	}
	pluginId := matches[1]

	regx = regexp.MustCompile(`(` + pluginId + `)(\/.*)`)
	matches = regx.FindStringSubmatch(inURL.Path)
	if len(matches) != 3 {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err=Unable to find path", http.StatusBadRequest)
	}
	path := matches[2]

	hooks, err := pluginsEnvironment.HooksForPlugin(pluginId)
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err="+err.Error(), http.StatusBadRequest)
	}
	w := LocalResponseWriter{}
	r, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err="+err.Error(), http.StatusBadRequest)
	}
	r.Header.Set("Mattermost-User-ID", a.Session.UserId)
	r.Header.Set(model.HEADER_AUTH, "Bearer "+a.Session.Token)
	c := &plugin.Context{
		SessionId:      model.NewId(),
		RequestId:      model.NewId(),
		IpAddress:      "localhost",
		AcceptLanguage: "*",
		UserAgent:      "",
	}

	hooks.ServeHTTP(c, &w, r)

	resp := &http.Response{
		Status:           "",
		StatusCode:       w.status,
		Proto:            "HTTP/1.1",
		ProtoMajor:       1,
		ProtoMinor:       1,
		Header:           make(http.Header),
		Body:             ioutil.NopCloser(bytes.NewReader(w.data)),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            true,
		Uncompressed:     true,
		Trailer:          make(http.Header),
		Request:          nil,
		TLS:              nil,
	}

	return resp, nil
}

func (a *App) OpenInteractiveDialog(request model.OpenDialogRequest) *model.AppError {
	clientTriggerId, userId, err := request.DecodeAndVerifyTriggerId(a.AsymmetricSigningKey())
	if err != nil {
		return err
	}

	request.TriggerId = clientTriggerId

	jsonRequest, _ := json.Marshal(request)

	message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_OPEN_DIALOG, "", "", userId, nil)
	message.Add("dialog", string(jsonRequest))
	a.Publish(message)

	return nil
}

func (a *App) SubmitInteractiveDialog(request model.SubmitDialogRequest) (*model.SubmitDialogResponse, *model.AppError) {
	url := request.URL
	request.URL = ""
	request.Type = "dialog_submission"

	b, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return nil, model.NewAppError("SubmitInteractiveDialog", "app.submit_interactive_dialog.json_error", nil, jsonErr.Error(), http.StatusBadRequest)
	}

	resp, err := a.DoActionRequest(url, b)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response model.SubmitDialogResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		// Don't fail, an empty response is acceptable
		return &response, nil
	}

	return &response, nil
}
