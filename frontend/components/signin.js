import FormStyles from './styles/formStyles';
import HomeStyles from './styles/homeStyles';
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
import { useState } from 'react';

const SIGNIN_MUTATION = gql`
  mutation SIGNIN_MUTATION($email: String!, $password: String!) {
    signin(email: $email, password: $password) {
      id
      name
      email
    }
  }
`;
function Signin() {
  const [signuser, { data, error }] = useMutation(SIGNIN_MUTATION);
  function useValue(initValue) {
    const [value, setValue] = useState(initValue);
    const handleValue = e => {
      setValue(e.target.value);
    };
    return {
      value,
      handleValue
    };
  }
  if (error) {
    console.error(error);
  }
  const email = useValue('');
  const password = useValue('');

  if (data) {
    console.log(data);
  }

  function handleSubmit(e) {
    e.preventDefault();
    signuser({
      variables: {
        email: email.value,
        password: password.value
      }
    });
  }

  return (
    <HomeStyles>
      <FormStyles onSubmit={handleSubmit}>
        <fieldset>
          <legend>Signin</legend>
          <div className='form form-group-name'>
            <label htmlFor='email'>Email</label>
            <input
              type='text'
              name='email'
              value={email.value}
              onChange={email.handleValue}
            />
          </div>
          <div className='form form-group-password'>
            <label htmlFor='password'> Password</label>
            <input
              type='password'
              name='password'
              value={password.value}
              onChange={password.handleValue}
            />
          </div>
          <div className='form form-group'>
            <button className='btn' type='submit'>
              {' '}
              submit
            </button>
          </div>
        </fieldset>
      </FormStyles>
    </HomeStyles>
  );
}

export default Signin;
