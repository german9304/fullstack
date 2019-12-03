import App from 'next/app';
import { ApolloProvider } from '@apollo/react-hooks';
import withApollo from '../lib/withApollo';
import Page from '../components/Page';

function MyApp({ Component, pageProps, apollo }) {
  return (
    <Page>
      <ApolloProvider client={apollo}>
        <Component {...pageProps} />
      </ApolloProvider>
    </Page>
  );
}

export default withApollo(MyApp);
