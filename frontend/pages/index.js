import Users from '../components/users';
import Posts from '../components/posts';
import styled from 'styled-components';

const HomeStyle = styled.div`
  display: flex;
  justify-content: center;
`;
function Home() {
  return (
    <HomeStyle>
      <Posts />
    </HomeStyle>
  );
}

export default Home;
