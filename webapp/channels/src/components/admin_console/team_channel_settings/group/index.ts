// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {connect, ConnectedProps} from 'react-redux';
import {t} from 'utils/i18n';
import {Group} from '@mattermost/types/groups';
import List from './group_list';
import {GlobalState} from 'types/store';

type OwnProps = {
    groups: Group[];
    totalGroups: number;
    isModeSync: boolean;
    onGroupRemoved: (gid: string) => void;
    setNewGroupRole: (gid: string) => void;
}

function mapStateToProps(state: GlobalState, ownProps: OwnProps) {
    return {
        data: ownProps.groups,
        total: ownProps.totalGroups,
        emptyListTextId: ownProps.isModeSync ? t('admin.team_channel_settings.group_list.no-synced-groups') : t('admin.team_channel_settings.group_list.no-groups'),
        emptyListTextDefaultMessage: ownProps.isModeSync ? 'At least one group must be specified' : 'No groups specified yet',
        removeGroup: ownProps.onGroupRemoved,
        setNewGroupRole: ownProps.setNewGroupRole,
    };
}

const mapDispatchToProps = {
    getData: () => Promise.resolve(),
};

const connector = connect(mapStateToProps, mapDispatchToProps);

export type PropsFromRedux = ConnectedProps<typeof connector>;

export default connector(List);
