import React from "react";
import { shallow } from 'enzyme';
import { Index } from "../index";

describe('<Index />', () => {
  it('renders correctly', () => {
    const wrapper = shallow(<Index />);
    expect(wrapper.find('[className="Index"]').text()).toBe('トップページ');
  });
});
