import styled from 'styled-components';
import Link from 'next/link';

const NavStyled = styled.nav`
  border-bottom: solid 1em #111111;
  li[class~='menu'] {
    color: #111111;
    font-weight: bolder;
    font-size: 1.3em;
    font-family: 'Montserrat', sans-serif;
  }

  ul[class~='menu'] {
    display: flex;
    list-style: none;
    justify-content: flex-end;
    margin: 1em;
  }

  ul[class~='menu'] a {
    text-decoration: none;
    cursor: pointer;
    color: #111111;
  }

  .page-title {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .apptitle {
    margin: 0;
    padding: 0;
    font-family: 'Montserrat', sans-serif;
    font-size: 1.5em;
  }
  .inner-container {
    color: #11111;
    display: flex;
    justify-content: center;
    padding: 1.5em;
    width: 200px;
  }
`;

function Nav() {
  return (
    <NavStyled>
      <section className='page-title'>
        <div className='inner-container'>
          <h1 className='apptitle'>POSTS</h1>
        </div>
      </section>
      <ul className='menu'>
        <div className='right'>
          <li className='menu signin'>
            <Link href='/signup'>
              <a>Signup</a>
            </Link>
          </li>
        </div>
      </ul>
    </NavStyled>
  );
}

export default Nav;
