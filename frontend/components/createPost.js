import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
import FormStyles from './styles/formStyles';
import User from './user';

const CREATE_POST_MUTATION = gql`
  mutation CREATE_POST_MUTATION($pstinpt: PostInput!) {
    createPost(pstinpt: $pstinpt) {
      id
      body
      header
      author {
        id
      }
    }
  }
`;

function CreatePost() {
  const currentUser = User();
  const [createPost, { data, error, loading }] = useMutation(
    CREATE_POST_MUTATION
  );

  if (error) console.error(error);
  if (data) {
    console.log(data);
  }

  function handleSubmit() {
    createPost({
      variables: {
        pstinpt: {
          author: '',
          header: '',
          body: ''
        }
      }
    });
  }

  return (
    <div className='create-post'>
      <FormStyles onSubmit={handleSubmit}>
        <fieldset>
          <div className='form form-group'>
            <label htmlFor='header'>Header</label>
            <input type='text' name='header' />
          </div>
          <div className='form form-group'>
            <label htmlFor='body'>Body</label>
            <input type='text' name='body' />
          </div>
        </fieldset>
      </FormStyles>
    </div>
  );
}

export default CreatePost;
