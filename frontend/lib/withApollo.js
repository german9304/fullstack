import withApollo from 'next-with-apollo';
import ApolloClient, { InMemoryCache } from 'apollo-boost';

const GRAPHQL_URL = 'http://backend:8000/';

function client({ headers, initialState }) {
  return new ApolloClient({
    uri: HERO_URL,
    cache: new InMemoryCache().restore(initialState || {})
  });
}
export default withApollo(client);
