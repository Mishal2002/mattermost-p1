// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

export function PinSVG(props: React.HTMLAttributes<HTMLSpanElement>) {
    const {formatMessage} = useIntl();
    return (
        <span {...props}>
            <svg
                width='16px'
                height='16px'
                viewBox='0 0 16 16'
                version='1.1'
                role='img'
                aria-label={formatMessage({id: 'generic_icons.pin', defaultMessage: 'Pin Icon'})}
            >
                <path
                    d='M6.38635 10.2446L6.3649 10.266L6.28152 10.3709C6.16941 10.4968 6.05008 10.6162 5.92413 10.7283L5.69302 10.9392L4.51722 11.9685L3.13056 13.0813C3.11601 13.0916 3.09957 13.099 3.08217 13.103C3.06478 13.1069 3.04677 13.1075 3.02918 13.1045C3.01159 13.1015 2.99476 13.0951 2.97965 13.0856C2.96455 13.0761 2.95145 13.0637 2.94113 13.0491C2.91479 13.0237 2.8983 12.9898 2.8946 12.9534C2.89089 12.9169 2.90021 12.8804 2.92089 12.8502L3.34022 12.3248L4.81027 10.5818L5.27249 10.0564C5.41307 9.91657 5.52504 9.80459 5.60843 9.7204L5.75615 9.5941C5.85019 9.52561 5.9653 9.49231 6.08139 9.49999C6.13809 9.50241 6.19362 9.51683 6.24435 9.5423C6.29507 9.56776 6.33981 9.60369 6.37563 9.64772C6.44966 9.72767 6.49202 9.83187 6.49476 9.9408C6.49161 10.051 6.45366 10.1573 6.38635 10.2446V10.2446Z'
                    fill='#A4A9B7'
                />
                <path
                    d='M4.19214 11.6306L5.5788 9.95807L5.82092 10.0901L4.19214 11.6306Z'
                    fill='#E8E9ED'
                />
                <path
                    d='M13.1063 6.64299C13.1074 6.69961 13.0967 6.75583 13.075 6.80812C13.0533 6.86041 13.0209 6.90763 12.9801 6.94679L12.8335 7.09452C12.6153 7.30982 12.3211 7.43052 12.0145 7.43052C11.708 7.43052 11.4137 7.30982 11.1955 7.09452C11.1906 7.08788 11.1843 7.08247 11.1769 7.07875C11.1696 7.07502 11.1615 7.07307 11.1532 7.07307C11.145 7.07307 11.1369 7.07502 11.1295 7.07875C11.1222 7.08247 11.1158 7.08788 11.1109 7.09452L9.36807 8.81609C9.31455 8.8719 9.28291 8.94514 9.27897 9.02237C9.27504 9.09961 9.29905 9.17568 9.34662 9.23665C9.55206 9.53842 9.64523 9.90266 9.60991 10.266C9.57423 10.6302 9.41489 10.9712 9.15841 11.2322L8.48652 11.9042C8.31064 12.079 8.07273 12.1771 7.82475 12.1771C7.57677 12.1771 7.33888 12.079 7.163 11.9042L4.07517 8.83753C3.90036 8.66164 3.80225 8.42371 3.80225 8.17572C3.80225 7.92772 3.90036 7.68979 4.07517 7.5139L4.76851 6.84195C5.02276 6.57429 5.36672 6.40972 5.73464 6.37969C6.09867 6.35206 6.46136 6.4482 6.76392 6.65252C6.82457 6.69406 6.898 6.71268 6.97111 6.70504C7.04423 6.6974 7.11224 6.664 7.163 6.61082L8.90347 4.89164C8.91045 4.87864 8.91411 4.86411 8.91411 4.84935C8.91411 4.83459 8.91045 4.82005 8.90347 4.80705C8.6882 4.58879 8.5675 4.29454 8.5675 3.98796C8.5675 3.68139 8.6882 3.38714 8.90347 3.16888L9.02975 3.02234C9.11874 2.94135 9.23465 2.89634 9.35497 2.89606C9.41174 2.8951 9.4681 2.90581 9.52056 2.92753C9.57302 2.94926 9.62046 2.98153 9.65993 3.02234L12.9777 6.34037C13.0187 6.37928 13.0514 6.42623 13.0735 6.47829C13.0956 6.53035 13.1068 6.58642 13.1063 6.64299Z'
                    fill='#FFBC1F'
                />
                <path
                    d='M7.88983 8.15345C7.16788 7.43161 6.97271 6.92837 6.96538 6.76699C5.99691 6.34884 5.93086 6.5249 6.3931 7.05308C6.86344 7.59052 8.79228 9.05574 7.88983 8.15345Z'
                    fill='#FFD470'
                />
                <path
                    d='M8.12913 7.49364C8.09642 7.52744 8.05714 7.55418 8.01371 7.57221C7.97027 7.59024 7.9236 7.59918 7.87658 7.59848C7.82971 7.5993 7.78319 7.59042 7.73993 7.57237C7.69667 7.55433 7.65761 7.52752 7.62522 7.49364C7.59261 7.46302 7.56652 7.42613 7.54851 7.38519C7.5305 7.34425 7.52093 7.30009 7.52039 7.25536C7.52281 7.1579 7.55999 7.06451 7.62522 6.99206L9.01068 5.62791C9.0589 5.58091 9.11976 5.54896 9.18584 5.53598C9.25191 5.523 9.32032 5.52954 9.38273 5.55482C9.44515 5.58009 9.49885 5.62299 9.53729 5.67829C9.57572 5.73358 9.59722 5.79888 9.59916 5.86619C9.59848 5.9109 9.58887 5.95501 9.57086 5.99594C9.55286 6.03686 9.52683 6.07377 9.49433 6.10447L8.12913 7.49126V7.49364Z'
                    fill='#FFD470'
                />
                <path
                    d='M10.8812 6.86747L10.6679 6.64979C9.4567 5.44732 8.69515 5.26865 8.45715 5.33754L8.93402 4.88055L10.6679 6.64979C10.7375 6.71891 10.8086 6.79142 10.8812 6.86747Z'
                    fill='#F5AB00'
                />
                <ellipse
                    cx='6.74289'
                    cy='13.0857'
                    rx='3.88571'
                    ry='0.285714'
                    fill='black'
                    fillOpacity='0.06'
                />
            </svg>
        </span>
    );
}
