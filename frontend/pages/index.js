import Users from '../components/users';
import Posts from '../components/posts';
import HomeStyle from '../components/styles/homeStyles';
import User from '../components/user';
import CreatePost from '../components/createPost';

function Home() {
  return (
    <HomeStyle>
      <CreatePost />
      <Posts />
    </HomeStyle>
  );
}

export default Home;
