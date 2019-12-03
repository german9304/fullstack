import styled from 'styled-components';
import Comment from './comment';

const PostStyle = styled.section`
  box-shadow: 0 0 2px 4px rgba(0, 0, 0, 0.2);
  padding: 1em;
  width: 500px;
  margin: 1em;
  .info {
    font-size: 1.4rem;
    margin: 0;
    font-family: 'Montserrat', sans-serif;
  }

  .header {
    font-size: 1.8rem;
    font-weight: bolder;
    margin-bottom: 1.3em;
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
    border: solid 1px #aaaaaa;
    outline: none;
    cursor: pointer;
    font-size: 1em;
  }
`;

function Post({ data }) {
  console.log(data);
  return (
    <PostStyle>
      <div className='post post-info'>
        <p className='info header'> {data.header}</p>
        <p className='info body'>{data.body}</p>
      </div>
      <div className='post posts-article-section'>
        <button className='like'>Like</button>
        <button className='comment'>comment</button>
      </div>
      <div className='posts-article-data'>
        {data.comments.map(comment => {
          return <Comment key={Comment.id} data={comment} />;
        })}
      </div>
    </PostStyle>
  );
}

export default Post;
