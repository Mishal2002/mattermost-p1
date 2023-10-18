// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo} from 'react';

import type {ChannelNotifyProps} from '@mattermost/types/channels';

import {NotificationLevels} from 'utils/constants';

type Actions = {
    updateChannelNotifyProps: (userId: string, channelId: string, props: ChannelNotifyProps) => void;
};

type Props = {
    user: { id: string };
    channel: { id: string };
    actions: Actions;
};

const UnmuteChannelButton = memo(({user, channel, actions}: Props) => {
    const handleClick = () => {
        actions.updateChannelNotifyProps(user.id, channel.id, {mark_unread: NotificationLevels.ALL} as ChannelNotifyProps);
    };

    return (
        <button
            type='button'
            className='navbar-toggle icon icon__mute'
            onClick={handleClick}
        >
            <span className='fa fa-bell-slash-o icon'/>
        </button>
    );
});

export default UnmuteChannelButton;
