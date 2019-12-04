import Comment from '../components/comment';
import { render } from '@testing-library/react';

describe('comments components', () => {
  it('should render with correct props', () => {
    const body = 'comment body';
    const data = {
      body
    };
    const { container } = render(<Comment data={data} />);
    expect(container.textContent).toEqual(body);
  });
});
