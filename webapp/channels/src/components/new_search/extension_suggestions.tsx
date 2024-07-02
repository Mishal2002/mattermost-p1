// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {FormattedMessage} from 'react-intl';
import styled from 'styled-components';

import type {SuggestionProps} from 'components/suggestion/suggestion';

import {getCompassIconClassName} from 'utils/utils';

import type {ExtensionItem} from './extension_suggestions_provider';

const SearchFileExtensionSuggestionContainer = styled.div`
    display: flex;
    align-items: center;
    padding: 8px 2.4rem;
    &.selected, &:hover {
        background: rgba(var(--center-channel-color-rgb), 0.08);
    }

    .file-icon {
        background-size: 16px 20px;
        width: 24px;
        height: 24px;
        font-size: 18px;
        margin-right: 8px;
        display: flex;
        align-items: center;
        &.icon-file-excel-outline {
            color: #339970;
        }
        &.icon-file-powerpoint-outline {
            color: #E07315;
        }
        &.icon-file-pdf-outline {
            color: #C43133;
        }
        &.icon-file-image-outline,&.icon-file-audio-outline, &.icon-file-video-outline, &.icon-file-word-outline {
            color: #5D89EA;
        }
    }
`;

const ExtensionText = styled.span`
    margin-left: 4px;
`;

const SearchFileExtensionSuggestion = React.forwardRef<HTMLDivElement, SuggestionProps<ExtensionItem>>((props, ref) => {
    const {item} = props;

    let labelName: React.ReactNode = item.type;

    switch (item.type) {
    case 'pdf':
        labelName = (
            <FormattedMessage
                id='file_type.pdf'
                defaultMessage='Acrobat'
            />
        );
        break;
    case 'word':
        labelName = (
            <FormattedMessage
                id='file_type.word'
                defaultMessage='Word Document'
            />
        );
        break;
    case 'image':
        labelName = (
            <FormattedMessage
                id='file_type.image'
                defaultMessage='Image'
            />
        );
        break;
    case 'audio':
        labelName = (
            <FormattedMessage
                id='file_type.audio'
                defaultMessage='Audio'
            />
        );
        break;
    case 'video':
        labelName = (
            <FormattedMessage
                id='file_type.video'
                defaultMessage='Video'
            />
        );
        break;
    case 'presentation':
        labelName = (
            <FormattedMessage
                id='file_type.presentation'
                defaultMessage='Powerpoint Presentation'
            />
        );
        break;
    case 'spreadsheet':
        labelName = (
            <FormattedMessage
                id='file_type.spreadsheet'
                defaultMessage='Excel spreadsheet'
            />
        );
        break;
    case 'code':
        labelName = (
            <FormattedMessage
                id='file_type.code'
                defaultMessage='Code file'
            />
        );
        break;
    case 'patch':
        labelName = (
            <FormattedMessage
                id='file_type.patch'
                defaultMessage='Patch file'
            />
        );
        break;
    case 'svg':
        labelName = (
            <FormattedMessage
                id='file_type.svg'
                defaultMessage='Vector graphics'
            />
        );
        break;
    case 'text':
        labelName = (
            <FormattedMessage
                id='file_type.text'
                defaultMessage='Text file'
            />
        );
        break;
    }

    return (
        <SearchFileExtensionSuggestionContainer
            ref={ref}
            className={props.isSelection ? 'selected' : ''}
            onClick={() => props.onClick(item.label, props.matchedPretext)}
        >
            <div className={'file-icon ' + getCompassIconClassName(item.type)}/>
            {labelName}
            <ExtensionText>{`(.${item.value})`}</ExtensionText>
        </SearchFileExtensionSuggestionContainer>
    );
});
SearchFileExtensionSuggestion.displayName = 'SearcFileExtensionSuggestion';

export default SearchFileExtensionSuggestion;
