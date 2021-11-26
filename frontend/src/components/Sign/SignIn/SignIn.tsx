import { useState, SyntheticEvent } from "react";
import { SignFormProps } from "../Sign"
import axios from "axios";

export default function SignIn(props: SignFormProps) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const updateEmail = (event: SyntheticEvent) => {
    setEmail((event.target as HTMLInputElement).value);
  }

  const updatePassword = (event: SyntheticEvent) => {
    setPassword((event.target as HTMLInputElement).value);
  }

  const handleSubmit = (event: SyntheticEvent) => {
    const article = { email: email, password: password };
    axios.post('http://localhost:8080/signin', article)
        .then(response => localStorage.setItem('jwt', response.data));
  }

  return (
    <div>
      <h1 className="noto font-normal fsz-18 text-left">Sign in</h1>
      <form onSubmit={handleSubmit}>
        <input type="email" name="email" className="sign_input fsz-16 noto"
          placeholder="Email" style={{marginTop:"24px"}} onChange={updateEmail} />

        <input type="password" name="password" className="sign_input fsz-16 noto"
          placeholder="Password" style={{marginTop:"20px"}} onChange={updatePassword} />

        <input type="submit" value="Sign in" className="noto fsz-12 submit" />
      </form>
    </div>
  )
}