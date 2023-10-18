// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {mount} from 'enzyme';
import React from 'react';

import LoadingWrapper from './loading_wrapper';

describe('components/widgets/loading/LoadingWrapper', () => {
    const testCases = [
        {
            name: 'showing spinner with text',
            loading: true,
            text: 'test',
            children: 'children',
            snapshot: `
<LoadingSpinner
  text="test"
/>
`,
        },
        {
            name: 'showing spinner without text',
            loading: true,
            children: 'text',
            snapshot: `
<LoadingSpinner
  text={null}
/>
`,
        },
        {
            name: 'showing content with children',
            loading: false,
            children: 'text',
            snapshot: '"text"',
        },
        {
            name: 'showing content without children',
            loading: false,
            snapshot: '""',
        },
    ];
    for (const testCase of testCases) {
        test(testCase.name, () => {
            const wrapper = mount(
                <LoadingWrapper
                    loading={testCase.loading}
                    text={testCase.text}
                >
                    {testCase.children}
                </LoadingWrapper>,
            );
            expect(wrapper).toMatchSnapshot(testCase.snapshot);
        });
    }
});
