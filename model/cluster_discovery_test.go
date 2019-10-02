// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClusterDiscovery(t *testing.T) {
	o := ClusterDiscovery{
		Type:        "test_type",
		ClusterName: "cluster_name",
		Hostname:    "test_hostname",
	}

	json := o.ToJson()
	result1 := ClusterDiscoveryFromJson(strings.NewReader(json))

	require.NotNil(t, result1)
	require.Equal(t, "cluster_name", result1.ClusterName)

	result2 := ClusterDiscoveryFromJson(strings.NewReader(json))
	result3 := ClusterDiscoveryFromJson(strings.NewReader(json))

	o.Id = "0"
	result1.Id = "1"
	result2.Id = "2"
	result3.Id = "3"
	result3.Hostname = "something_diff"

	require.True(t, o.IsEqual(result1))

	list := make([]*ClusterDiscovery, 0)
	list = append(list, &o)
	list = append(list, result1)
	list = append(list, result2)
	list = append(list, result3)

	rlist := FilterClusterDiscovery(list, func(in *ClusterDiscovery) bool {
		return !o.IsEqual(in)
	})

	require.Len(t, rlist, 1)

	o.AutoFillHostname()
	o.Hostname = ""
	o.AutoFillHostname()

	o.AutoFillIpAddress("", "")
	o.Hostname = ""
	o.AutoFillIpAddress("", "")
}
