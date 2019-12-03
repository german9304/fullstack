import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

const ME_QUERY = gql`
  query Me {
    me {
      id
      email
    }
  }
`;

function User() {
  const { data, loading, error } = useQuery(ME_QUERY);

  if (error) {
    console.error(error);
  }

  if (data) {
    console.log('current user ', data);
  }

  return (
    <div className='user'>
      <h1>user</h1>
    </div>
  );
}

export default User;
