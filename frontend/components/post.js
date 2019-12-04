import styled from 'styled-components';
import Comment from './comment';

const PostStyle = styled.section`
  box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.08);
  background: rgba(0, 0, 0, 0.02);
  padding: 1em;
  width: 500px;
  margin-bottom: 2em;
  .info {
    font-size: 1rem;
    margin: 0;
    font-family: 'Montserrat', sans-serif;
  }

  .header {
    font-size: 1rem;
    font-weight: bolder;
    margin-bottom: 1em;
  }

  .info.body {
    font-size: 1.2rem;
    font-weight: bolder;
  }
  .posts-article-section {
    display: grid;
    grid-template-columns: repeat(2, 100px);
    padding: 0.3em 0;
    margin-top: 1em;
    grid-gap: 1em;
  }
  .posts-article-section button {
    padding: 0.2em;
    outline: none;
    border: none;
    cursor: pointer;
    font-size: 1em;
  }
  .post .post-meta {
    display: flex;
  }

  .post-meta .icon {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 0.2em;
  }

  .material-icons {
    font-size: 1.2em;
  }
`;

function Post({ data }) {
  return (
    <PostStyle>
      <div className='post post-info'>
        <p className='info header'>{data.header}</p>
        <p className='info body'>{data.body}</p>
      </div>
      <div className='post posts-article-section'>
        <button className='post-meta like'>
          <span className='icon'>
            <i className='material-icons'>thumb_up</i>
          </span>
          <span className='text-icon'>Like</span>
        </button>
        <button className='post-meta comment'>
          <span className='icon'>
            <i className='material-icons'>comment</i>
          </span>
          <span className='text-icon'>comment</span>
        </button>
      </div>
      <div className='posts-article-data'>
        {data.comments.map((comment, i) => {
          return <Comment key={comment.id} data={comment} />;
        })}
      </div>
    </PostStyle>
  );
}

export default Post;
