// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {expect} from '@playwright/test';

import {test} from '@e2e-support/test_fixture';
import {getRandomId} from '@e2e-support/util';

const keywords = [`AB${getRandomId()}`, `CD${getRandomId()}`, `EF${getRandomId()}`, `Highlight me ${getRandomId()}`];

const highlightWithoutNotificationClass = 'non-notification-highlight';

test('MM-T5465-1 Should add the keyword when enter, comma or tab is pressed on the textbox', async ({pw, pages}) => {
    const {user} = await pw.initSetup();

    // # Log in as a user in new browser context
    const {page} = await pw.testBrowser.login(user);

    // # Visit default channel page
    const channelPage = new pages.ChannelsPage(page);
    await channelPage.goto();
    await channelPage.toBeVisible();

    await channelPage.centerView.postCreate.postMessage('Hello World');

    // # Open settings modal
    await channelPage.globalHeader.openSettings();
    await channelPage.settingsModal.toBeVisible();

    // # Open notifications tab
    await channelPage.settingsModal.openNotificationsTab();

    // # Open keywords that get highlighted section
    await channelPage.settingsModal.notificationsSettings.expandSection('keysWithHighlight');

    const keywordsInput = await channelPage.settingsModal.notificationsSettings.getKeywordsInput();

    // # Enter keyword 1
    await keywordsInput.type(keywords[0]);

    // # Press Comma on the textbox
    await keywordsInput.press(',');

    // # Enter keyword 2
    await keywordsInput.type(keywords[1]);

    // # Press Tab on the textbox
    await keywordsInput.press('Tab');

    // # Enter keyword 3
    await keywordsInput.type(keywords[2]);

    // # Press Enter on the textbox
    await keywordsInput.press('Enter');

    // * Verify that the keywords have been added to the collapsed description
    await expect(channelPage.settingsModal.notificationsSettings.container.getByText(keywords[0])).toBeVisible();
    await expect(channelPage.settingsModal.notificationsSettings.container.getByText(keywords[1])).toBeVisible();
    await expect(channelPage.settingsModal.notificationsSettings.container.getByText(keywords[2])).toBeVisible();
});

test('MM-T5465-2 Should highlight the keywords when a message is sent with the keyword in center', async ({pw, pages}) => {
    const {user} = await pw.initSetup();

    // # Log in as a user in new browser context
    const {page} = await pw.testBrowser.login(user);

    // # Visit default channel page
    const channelPage = new pages.ChannelsPage(page);
    await channelPage.goto();
    await channelPage.toBeVisible();

    // # Open settings modal
    await channelPage.globalHeader.openSettings();
    await channelPage.settingsModal.toBeVisible();

    // # Open notifications tab
    await channelPage.settingsModal.openNotificationsTab();

    // # Open keywords that get highlighted section
    await channelPage.settingsModal.notificationsSettings.expandSection('keysWithHighlight');

    // # Enter the keyword
    const keywordsInput = await channelPage.settingsModal.notificationsSettings.getKeywordsInput();
    await keywordsInput.type(keywords[3]);
    await keywordsInput.press('Tab');

    // # Save the keyword
    await channelPage.settingsModal.notificationsSettings.save();

    // # Close the settings modal
    await channelPage.settingsModal.closeModal();

    // # Post a message without the keyword
    const messageWithoutKeyword = 'This message does not contain the keyword';
    await channelPage.centerView.postCreate.postMessage(messageWithoutKeyword);
    const lastPostWithoutHighlight = await channelPage.centerView.getLastPost();

    // * Verify that the keywords are not highlighted
    await expect(lastPostWithoutHighlight.container.getByText(messageWithoutKeyword)).toBeVisible();
    await expect(lastPostWithoutHighlight.container.getByText(messageWithoutKeyword)).not.toHaveClass(
        highlightWithoutNotificationClass
    );

    // # Post a message with the keyword
    const messageWithKeyword = `This message contains the keyword ${keywords[3]}`;
    await channelPage.centerView.postCreate.postMessage(messageWithKeyword);
    const lastPostWithHighlight = await channelPage.centerView.getLastPost();

    // * Verify that the keywords are highlighted
    await expect(lastPostWithHighlight.container.getByText(messageWithKeyword)).toBeVisible();
    await expect(lastPostWithHighlight.container.getByText(keywords[3])).toHaveClass(highlightWithoutNotificationClass);
});

test('MM-T5465-3 Should highlight the keywords when a message is sent with the keyword in rhs', async ({pw, pages}) => {
    const {user} = await pw.initSetup();

    // # Log in as a user in new browser context
    const {page} = await pw.testBrowser.login(user);

    // # Visit default channel page
    const channelPage = new pages.ChannelsPage(page);
    await channelPage.goto();
    await channelPage.toBeVisible();

    // # Open settings modal
    await channelPage.globalHeader.openSettings();
    await channelPage.settingsModal.toBeVisible();

    // # Open notifications tab
    await channelPage.settingsModal.openNotificationsTab();

    // # Open keywords that get highlighted section
    await channelPage.settingsModal.notificationsSettings.expandSection('keysWithHighlight');

    // # Enter the keyword
    const keywordsInput = await channelPage.settingsModal.notificationsSettings.getKeywordsInput();
    await keywordsInput.type(keywords[3]);
    await keywordsInput.press('Tab');

    // # Save the keyword
    await channelPage.settingsModal.notificationsSettings.save();

    // # Close the settings modal
    await channelPage.settingsModal.closeModal();

    // # Post a message without the keyword
    const messageWithoutKeyword = 'This message does not contain the keyword';
    await channelPage.centerView.postCreate.postMessage(messageWithoutKeyword);
    const lastPostWithoutHighlight = await channelPage.centerView.getLastPost();

    // # Open the message in the RHS
    await lastPostWithoutHighlight.hover();
    await lastPostWithoutHighlight.postMenu.toBeVisible();
    await lastPostWithoutHighlight.postMenu.reply();
    await channelPage.sidebarRight.toBeVisible();

    // # Post a message with the keyword in the RHS
    const messageWithKeyword = `This message contains the keyword ${keywords[3]}`;
    await channelPage.sidebarRight.postCreate.postMessage(messageWithKeyword)

    // * Verify that the keywords are highlighted
    const lastPostWithHighlightInRHS = await channelPage.sidebarRight.getLastPost();
    await expect(lastPostWithHighlightInRHS.container.getByText(messageWithKeyword)).toBeVisible();
    await expect(lastPostWithHighlightInRHS.container.getByText(keywords[3])).toHaveClass(highlightWithoutNotificationClass);
});
