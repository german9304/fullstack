import Nav from './nav';
import Meta from './meta';

function Page({ children }) {
  return (
    <div>
      <Meta />
      <Nav />
      <main className='container'>{children}</main>
    </div>
  );
}

export default Page;
