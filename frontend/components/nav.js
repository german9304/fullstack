import styled from 'styled-components';

const NavStyled = styled.nav`
  background-color: #0074d9;
  li[class~='menu'] {
    color: #ffff;
    font-family: 'Montserrat', sans-serif;
  }

  ul[class~='menu'] {
    padding: 1.5em;
    display: flex;
    list-style: none;
    justify-content: flex-end;
  }

  .right {
    display: grid;
    grid-gap: 1em;
    grid-template-columns: repeat(2, 1fr);
  }
`;

function Nav() {
  return (
    <NavStyled>
      <ul className='menu'>
        <div className='right'>
          <li className='menu signin'>Signin</li>
          <li className='menu signout'>Signout</li>
        </div>
      </ul>
    </NavStyled>
  );
}

export default Nav;
