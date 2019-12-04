import { render } from '@testing-library/react';
import Post from '../components/post';
import Posts, { POSTS_QUERY } from '../components/posts';
import { MockedProvider } from '@apollo/react-testing';
import wait from 'waait';

describe('POSTS and POST component', () => {
  const consoleError = console.error;
  beforeAll(() => {
    jest.spyOn(console, 'error').mockImplementation((...args) => {
      if (
        !args[0].includes(
          'Warning: An update to %s inside a test was not wrapped in act'
        )
      ) {
        consoleError(...args);
      }
    });
  });
  it('posts component should render without failing', async () => {
    const postInfo = ['postheader', 'postbody'];
    const mocks = [
      {
        request: {
          query: POSTS_QUERY
        },
        result: {
          data: {
            posts: [
              {
                id: '1',
                header: postInfo[0],
                body: postInfo[1],
                comments: [
                  {
                    id: '2',
                    body: 'comment body'
                  }
                ]
              }
            ]
          }
        }
      }
    ];

    let { container } = render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Posts />
      </MockedProvider>
    );
    await wait(0);
    const info = container.querySelectorAll('.info');
    info.forEach((el, i) => {
      expect(el.textContent).toEqual(postInfo[i]);
    });
  });

  it('post component should render correct props', () => {
    const dataHeader = 'header';
    const dataBody = 'body';
    const data = {
      header: dataHeader,
      body: dataBody,
      comments: [
        {
          id: '123',
          body: 'body comment'
        }
      ]
    };
    const { container, queryByText } = render(<Post data={data} />);
    const info = container.querySelector('.post-info');
    const header = info.querySelector('.header');
    const body = info.querySelector('.body');
    expect(header.textContent).toEqual(dataHeader);
    expect(body.textContent).toEqual(dataBody);
  });
});
