import gql from 'graphql-tag';
import styled from 'styled-components';
import { useQuery } from '@apollo/react-hooks';
import Post from './post';

const PostStyle = styled.section``;

export const POSTS_QUERY = gql`
  query POSTS_QUERY {
    posts {
      id
      header
      body
      comments {
        id
        body
      }
    }
  }
`;
function Posts() {
  const { data, loading, error } = useQuery(POSTS_QUERY);
  if (loading) return <p>loading...</p>;
  if (error) console.error(error);
  return (
    <PostStyle>
      {data.posts.map(post => (
        <Post key={post.id} data={post} />
      ))}
    </PostStyle>
  );
}

export default Posts;
