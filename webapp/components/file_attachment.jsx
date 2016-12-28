// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import Constants from 'utils/constants.jsx';
import FileStore from 'stores/file_store.jsx';
import * as Utils from 'utils/utils.jsx';

import {Tooltip, OverlayTrigger} from 'react-bootstrap';

import React from 'react';

import removeCTA from 'images/delete_CTA.svg';
import downloadCTA from 'images/download_CTA.svg';

export default class FileAttachment extends React.Component {
    constructor(props) {
        super(props);

        this.loadFiles = this.loadFiles.bind(this);
        this.onAttachmentClick = this.onAttachmentClick.bind(this);
        this.onRemoveClick = this.onRemoveClick.bind(this);

        this.state = {
            loaded: Utils.getFileType(props.fileInfo.extension) !== 'image'
        };
    }

    componentDidMount() {
        this.loadFiles();
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.fileInfo.id !== this.props.fileInfo.id) {
            this.setState({
                loaded: Utils.getFileType(nextProps.fileInfo.extension) !== 'image'
            });
        }
    }

    componentDidUpdate(prevProps) {
        if (!this.state.loaded && this.props.fileInfo.id !== prevProps.fileInfo.id) {
            this.loadFiles();
        }
    }

    loadFiles() {
        const fileInfo = this.props.fileInfo;
        const fileType = Utils.getFileType(fileInfo.extension);

        if (fileType === 'image') {
            let thumbnailUrl;
            if (this.props.displayType === 'preview') {
                thumbnailUrl = FileStore.getFileUrl(fileInfo.id);
            } else {
                thumbnailUrl = FileStore.getFileThumbnailUrl(fileInfo.id);
            }
            const img = new Image();
            img.onload = () => {
                this.setState({loaded: true});
            };
            img.load(thumbnailUrl);
        }
    }

    onAttachmentClick(e) {
        e.preventDefault();
        if (this.props.displayType !== 'preview') {
            this.props.handleImageClick(this.props.index);
        }
    }

    onRemoveClick(e) {
        e.preventDefault();
        if (this.props.handleRemoveClick) {
            this.props.handleRemoveClick();
        }
    }

    render() {
        const fileInfo = this.props.fileInfo;
        const fileName = fileInfo.name;
        const fileUrl = FileStore.getFileUrl(fileInfo.id);
        const fileType = Utils.getFileType(fileInfo.extension);

        let thumbnail;
        if (this.state.loaded) {
            if (fileType === 'image') {
                let thumbnailUrl;
                if (this.props.displayType === 'preview') {
                    thumbnailUrl = FileStore.getFileUrl(fileInfo.id);
                } else {
                    thumbnailUrl = FileStore.getFileThumbnailUrl(fileInfo.id);
                }

                thumbnail = (
                    <div
                        className='post-image'
                        style={{
                            backgroundImage: `url(${thumbnailUrl})`
                        }}
                    />
                );
            } else {
                thumbnail = <div className={'file-icon ' + Utils.getIconClassName(fileType)}/>;
            }
        } else {
            thumbnail = <div className='post-image__load'/>;
        }

        let trimmedFilename;
        if (fileName.length > 35) {
            trimmedFilename = fileName.substring(0, Math.min(35, fileName.length)) + '...';
        } else {
            trimmedFilename = fileName;
        }

        let filenameOverlay;
        if (this.props.compactDisplay) {
            filenameOverlay = (
                <OverlayTrigger
                    className='hidden-xs'
                    delayShow={1000}
                    placement='top'
                    overlay={(
                        <Tooltip
                            id='file-name__tooltip'
                            className='hidden-xs'
                        >
                            {fileName}
                        </Tooltip>
                    )}
                >
                    <a
                        href='#'
                        onClick={this.onAttachmentClick}
                        className='post-image__name'
                        rel='noopener noreferrer'
                    >
                        <span
                            className='icon'
                            dangerouslySetInnerHTML={{__html: Constants.ATTACHMENT_ICON_SVG}}
                        />
                        {trimmedFilename}
                    </a>
                </OverlayTrigger>
            );
        } else if (this.props.displayType === 'preview') {
            filenameOverlay = (
                <OverlayTrigger
                    className='hidden-xs'
                    delayShow={1000}
                    placement='top'
                    overlay={(
                        <Tooltip
                            id='file-name__tooltip'
                            className='hidden-xs'
                        >
                            {fileName}
                        </Tooltip>
                    )}
                >
                    <a
                        href='#'
                        className='post-image__name'
                    >
                        {trimmedFilename}
                    </a>
                </OverlayTrigger>
            );
        } else {
            filenameOverlay = (
                <OverlayTrigger
                    className='hidden-xs'
                    delayShow={1000}
                    placement='top'
                    overlay={(
                        <Tooltip
                            id='file-name__tooltip'
                            className='hidden-xs'
                        >
                            {Utils.localizeMessage('file_attachment.download', 'Download') + ' "' + fileName + '"'}
                        </Tooltip>
                    )}
                >
                    <a
                        href={fileUrl}
                        download={fileName}
                        className='post-image__name'
                        target='_blank'
                        rel='noopener noreferrer'
                    >
                        {trimmedFilename}
                    </a>
                </OverlayTrigger>
            );
        }

        let fileCTA;
        if (this.props.displayType === 'preview') {
            fileCTA = (
                <a
                    href='#'
                    className='post-image__download'
                    onClick={this.onRemoveClick}
                >
                    <img src={removeCTA}/>
                </a>
            );
        } else {
            fileCTA = (
                <a
                    href={fileUrl}
                    download={fileName}
                    className='post-image__download'
                    target='_blank'
                    rel='noopener noreferrer'
                >
                    <img src={downloadCTA}/>
                </a>
            );
        }

        let fileAttachment;
        if (fileType === 'image' && !this.props.compactDisplay && this.props.displayType !== 'preview') {
            fileAttachment = (
                <div className='post-image__preview'>
                    <a
                        href='#'
                        onClick={this.onAttachmentClick}
                    >
                        <img
                            className='post-image__image'
                            src={FileStore.getFileThumbnailUrl(fileInfo.id)}
                        />
                    </a>
                    <div className='post-image__info'>
                        <a
                            href={fileUrl}
                            download={fileName}
                            className='post-image__download'
                            target='_blank'
                            rel='noopener noreferrer'
                        >
                            <img src={downloadCTA}/>
                            {filenameOverlay}
                            <span className='post-image__size'>{Utils.fileSizeToString(fileInfo.size)}</span>
                        </a>
                    </div>
                </div>
            );
        } else {
            fileAttachment = (
                <div
                    className={this.props.displayType === 'preview' ? 'post-image__column preview' : 'post-image__column'}
                >
                    <a
                        className='post-image__thumbnail'
                        href='#'
                        onClick={this.onAttachmentClick}
                    >
                        {thumbnail}
                    </a>
                    <div className='post-image__details'>
                        {filenameOverlay}
                        <div className='post-image__info'>
                            <span className='post-image__type'>{fileInfo.extension.toUpperCase()}</span>
                            <span className='post-image__size'>{Utils.fileSizeToString(fileInfo.size)}</span>
                        </div>
                        <div>{fileCTA}</div>
                    </div>
                </div>
            );
        }
        return fileAttachment;
    }
}

FileAttachment.propTypes = {
    fileInfo: React.PropTypes.object.isRequired,

    // the index of this attachment preview in the parent FileAttachmentList
    index: React.PropTypes.number,

    // handler for when the thumbnail is clicked passed the index above
    handleImageClick: React.PropTypes.func,

    handleRemoveClick: React.PropTypes.func,

    // available values:
    // (default) inline: the inline display for files uploaded
    // preview: the preview display under the text input field
    displayType: React.PropTypes.string,

    compactDisplay: React.PropTypes.bool
};
