import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

const USERS_QUERY = gql`
  query USERS_QUERY {
    users {
      id
      name
      email
    }
  }
`;

function Users() {
  const { data, loading, error } = useQuery(USERS_QUERY);
  if (loading) return <p>loading...</p>;
  if (error) console.error(error);

  return (
    <div className='users'>
      <ol>
        {data.users.map(({ id, email, name }) => {
          return (
            <li key={id}>
              <div className='card'>
                <p className='email'>{email}</p>
                <p className='name'>{name}</p>
              </div>
            </li>
          );
        })}
      </ol>
      <style jsx>{`
        .card {
          border: solid 1px red;
        }
      `}</style>
    </div>
  );
}

export default Users;
