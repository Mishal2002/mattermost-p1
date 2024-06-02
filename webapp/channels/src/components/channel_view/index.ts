// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {connect} from 'react-redux';
import type {ConnectedProps} from 'react-redux';
import {withRouter} from 'react-router-dom';

import {getCurrentChannel, getDirectTeammate} from 'mattermost-redux/selectors/entities/channels';
import {getConfig, getLicense} from 'mattermost-redux/selectors/entities/general';
import {getTheme} from 'mattermost-redux/selectors/entities/preferences';
import {getCurrentRelativeTeamUrl} from 'mattermost-redux/selectors/entities/teams';
import {isFirstAdmin} from 'mattermost-redux/selectors/entities/users';

import {goToLastViewedChannel} from 'actions/views/channel';
import {getChannelTabPluginComponent} from 'selectors/plugins';

import type {GlobalState} from 'types/store';

import ChannelView from './channel_view';

function isDeactivatedChannel(state: GlobalState, channelId: string) {
    const teammate = getDirectTeammate(state, channelId);

    return Boolean(teammate && teammate.delete_at);
}

function mapStateToProps(state: GlobalState) {
    const channel = getCurrentChannel(state);

    const config = getConfig(state);

    const viewArchivedChannels = config.ExperimentalViewArchivedChannels === 'true';
    const enableOnboardingFlow = config.EnableOnboardingFlow === 'true';
    const enableWebSocketEventScope = config.FeatureFlagWebSocketEventScope === 'true';

    return {
        channelId: channel ? channel.id : '',
        deactivatedChannel: channel ? isDeactivatedChannel(state, channel.id) : false,
        tabContent: getChannelTabPluginComponent(state, state.views.channel.channelTab),
        enableOnboardingFlow,
        channelIsArchived: channel ? channel.delete_at !== 0 : false,
        viewArchivedChannels,
        isCloud: getLicense(state).Cloud === 'true',
        teamUrl: getCurrentRelativeTeamUrl(state),
        isFirstAdmin: isFirstAdmin(state),
        enableWebSocketEventScope,
        channelContentPlugin: channel ? state.plugins.channelContent[channel.id] : null,
        theme: getTheme(state),
    };
}

const mapDispatchToProps = ({
    goToLastViewedChannel,
});

const connector = connect(mapStateToProps, mapDispatchToProps);

export type PropsFromRedux = ConnectedProps<typeof connector>;

export default withRouter(connector(ChannelView));
