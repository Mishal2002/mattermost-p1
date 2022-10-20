// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/utils"
)

func (a *App) registerCollectionAndTopic(pluginID, collectionType, topicType string) error {
	// we have a race condition due to multiple plugins calling this method
	a.ch.collectionAndTopicTypesMut.Lock()
	defer a.ch.collectionAndTopicTypesMut.Unlock()

	// check if collectionType was already registered by other plugin
	for plugID, colTypes := range a.ch.collectionTypes {
		if utils.StringInSlice(collectionType, colTypes) && plugID != pluginID {
			return model.NewAppError("registerCollectionAndTopic", "app.collection.add_collection.exists.app_error", nil, "", http.StatusBadRequest)
		}
	}

	// check if topicType was already registered to other collection
	for collectionTypeForTopic, topicTypes := range a.ch.topicTypes {
		if utils.StringInSlice(topicType, topicTypes) && collectionTypeForTopic != collectionType {
			return model.NewAppError("registerCollectionAndTopic", "app.collection.add_topic.exists.app_error", nil, "", http.StatusBadRequest)
		}
	}

	a.ch.collectionTypes[pluginID] = appendIfUnique(a.ch.collectionTypes[pluginID], collectionType)
	a.ch.topicTypes[collectionType] = appendIfUnique(a.ch.topicTypes[collectionType], topicType)

	return nil
}

func appendIfUnique(slice []string, a string) []string {
	if utils.StringInSlice(a, slice) {
		return slice
	}
	return append(slice, a)
}
