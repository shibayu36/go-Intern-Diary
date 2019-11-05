import React from 'react';
import wait from 'waait';
import { mount } from 'enzyme';
import { getMyDiariesQuery, Diaries } from '../diaries';
import { MockedProvider } from '@apollo/react-testing';
import { act } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';

describe('<Diaries />', () => {
  it('正常にレンダリングできる', async () => {
    const mocks = [
      {
        request: {
          query: getMyDiariesQuery
        },
        result: {
          data: {
            visitor: {
              id: '1',
              name: 'shibayu36',
              diaries: [
                {
                  id: '2',
                  name: 'shibayu36の日記'
                },
                {
                  id: '3',
                  name: 'shibayu36のブログ'
                }
              ]
            }
          }
        }
      }
    ];

    const wrapper = mount(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Router>
          <Diaries />
        </Router>
      </MockedProvider>
    );

    await act(async () => {
      await wait(0);
    });
    wrapper.update();

    expect(wrapper.find('h1').text()).toBe('shibayu36のダイアリー一覧');

    const diaries = wrapper.find('Link');
    expect(diaries).toHaveLength(3);
    expect(diaries.at(0).prop('to')).toBe('/diaries/2');
    expect(diaries.at(0).find('p').text()).toBe('shibayu36の日記');
    expect(diaries.at(1).prop('to')).toBe('/diaries/3');
    expect(diaries.at(1).find('p').text()).toBe("shibayu36のブログ");
  });
});
