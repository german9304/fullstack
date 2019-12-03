import withApollo from 'next-with-apollo';
import ApolloClient, { InMemoryCache } from 'apollo-boost';

const GRAPHQL_URL = 'http://localhost:8000/';

function client({ headers, initialState }) {
  return new ApolloClient({
    credentials: 'include',
    uri: GRAPHQL_URL,
    cache: new InMemoryCache().restore(initialState || {})
  });
}
export default withApollo(client);
