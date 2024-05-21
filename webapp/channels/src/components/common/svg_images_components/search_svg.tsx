// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

export function SearchSVG(props: React.HTMLAttributes<HTMLSpanElement>) {
    const {formatMessage} = useIntl();
    return (
        <span {...props}>
            <svg
                width='16px'
                height='16px'
                viewBox='0 0 16 16'
                role='img'
                aria-label={formatMessage({id: 'generic_icons.search', defaultMessage: 'Search Icon'})}
            >
                <ellipse
                    cx='8.00003'
                    cy='13.9428'
                    rx='5.14286'
                    ry='0.342857'
                    fill='black'
                    fillOpacity='0.06'
                />
                <path
                    opacity='0.4'
                    d='M4.29632 4.28665C4.83967 3.73638 5.52543 3.45295 6.35221 3.43774C7.16378 3.45295 7.84539 3.735 8.39704 4.28665C8.9473 4.8383 9.23073 5.51991 9.24594 6.33148C9.23073 7.15826 8.94868 7.84264 8.39704 8.38738C7.84539 8.93073 7.16378 9.21831 6.35221 9.24596C5.52543 9.21692 4.84106 8.93073 4.29632 8.38738C3.75297 7.84264 3.46678 7.15688 3.43774 6.3301C3.46678 5.51853 3.75297 4.83691 4.29632 4.28665Z'
                    fill='var(--center-channel-bg)'
                />
                <path
                    d='M9.0157 5.81503C8.76979 5.01598 8.30426 4.41787 7.61841 4.0219C6.93256 3.62592 6.18182 3.52182 5.36688 3.70838C4.92123 3.82219 4.53297 4.01894 4.19893 4.2984C4.57246 3.89072 5.04629 3.60843 5.6192 3.45082C6.43414 3.26426 7.18858 3.3705 7.8806 3.77004C8.57262 4.16957 9.04185 4.76982 9.28776 5.56886C9.43894 6.14453 9.43139 6.69603 9.26386 7.22264C9.09757 7.74996 8.79906 8.19292 8.36781 8.54959C8.6916 8.19674 8.91439 7.78236 9.03565 7.30455C9.15568 6.82602 9.14898 6.33055 9.0157 5.81503Z'
                    fill='black'
                    fillOpacity='0.4'
                />
                <path
                    d='M9.91539 6.1061C9.85737 5.24836 9.50928 4.49509 8.87249 3.84628C8.14729 3.15487 7.31021 2.80847 6.36124 2.80847C5.41227 2.80847 4.57657 3.15487 3.85137 3.84628C3.15657 4.56794 2.80847 5.39956 2.80847 6.3439C2.80847 7.28824 3.15657 8.12124 3.85137 8.84152C4.51717 9.47658 5.27828 9.81885 6.13332 9.86834C6.98836 9.9192 7.7771 9.68414 8.50229 9.16592L9.04515 9.70614L9.73996 9.01472L9.1971 8.47451C9.73305 7.75423 9.97202 6.96384 9.91539 6.1061ZM8.41527 8.37691C7.86412 8.91712 7.18313 9.20304 6.37229 9.23053C5.54626 9.20166 4.8625 8.91712 4.31826 8.37691C3.7754 7.8367 3.48946 7.15491 3.46046 6.33291C3.48946 5.52603 3.7754 4.84836 4.31826 4.2999C4.86112 3.75281 5.54626 3.47102 6.37229 3.4559C7.18313 3.47102 7.86412 3.75144 8.41527 4.2999C8.96504 4.84836 9.24821 5.52603 9.2634 6.33291C9.24821 7.15491 8.96642 7.8367 8.41527 8.37691Z'
                    fill='#BABEC9'
                />
                <path
                    d='M12.1374 12.9928C11.9582 13.044 11.8256 13.026 11.7395 12.9388L8.92118 9.84586C8.83503 9.75867 8.79947 9.64658 8.81315 9.50819C8.82682 9.3698 8.91297 9.2148 9.07159 9.04043C9.24389 8.88129 9.39704 8.78995 9.53378 8.76781C9.67053 8.74566 9.78129 8.7858 9.86744 8.8882L12.9442 11.7404C13.0303 11.8276 13.044 11.9577 12.9866 12.132C12.9291 12.3064 12.8211 12.4808 12.6638 12.6552C12.4929 12.8295 12.3165 12.9416 12.1374 12.9928Z'
                    fill='#FFBC1F'
                />
                <path
                    d='M12.3436 11.182L11.1915 12.3437L9.80249 10.819L10.8213 9.77838L12.3436 11.182Z'
                    fill='#7A5600'
                />
            </svg>
        </span>
    );
}
