// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {FormattedMessage} from 'react-intl';

type Props = {
    customDescriptionText?: string;
    siteName: string;
};
const SiteNameAndDescription = ({customDescriptionText, siteName = 'Mattermost'}: Props) => {
    let description = null;
    if (customDescriptionText) {
        description = customDescriptionText;
    } else {
        description = (
            <FormattedMessage
                id='web.root.signup_info'
                defaultMessage='All team communication in one place, searchable and accessible anywhere'
            />
        );
    }

    return (
        <React.Fragment>
            <h1 id='site_name'>{siteName}</h1>
            <h3
                id='site_description'
                className='color--light'
            >
                {description}
            </h3>
        </React.Fragment>
    );
};

export default React.memo(SiteNameAndDescription);
