import Users from '../components/users';
import Posts from '../components/posts';
import HomeStyle from '../components/styles/homeStyles';
import User from '../components/user';

function Home() {
  return (
    <HomeStyle>
      <Posts />
    </HomeStyle>
  );
}

export default Home;
