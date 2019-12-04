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
    return {
      error,
      data: null
    };
  }

  if (data) {
    return {
      error: null,
      data
    };
  }
}

export default User;
