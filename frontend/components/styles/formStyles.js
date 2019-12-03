import styled from 'styled-components';

const FormStyles = styled.form`
  --font-f: 'Montserrat', sans-serif;
  /* border: solid 1px red; */
  margin-top: 2em;
  width: 400px;
  label {
    display: block;
    margin-bottom: 1em;
    font-family: var(--font-f);
  }
  /* .form {
    border: solid 1px gray;
  } */

  .form {
    margin-bottom: 1.2em;
  }
  input {
    box-sizing: border-box;
    height: 50px;
    width: 100%;
    padding: 0.5em;
    border: solid 1px #111111;
    font-size: 0.8rem;
  }

  .form .btn {
    background-color: #0074d9;
    cursor: pointer;
    color: #ffffff;
    width: 100%;
    padding: 0.4em;
    border: none;
    font-size: 1.2em;
    font-family: var(--font-f);
  }

  fieldset {
    padding: 1.4em;
  }
`;

export default FormStyles;
