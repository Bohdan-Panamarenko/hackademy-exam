import backgound from "./img/background.svg"
import logo from "./img/Group 3.svg"
import { useState } from "react";
import SignIn from "./SignIn/SignIn";
import SignUp from "./SignUp/SingUp";

export interface SignFormProps {
  changeState: React.Dispatch<React.SetStateAction<boolean>>
}

export default function Sign() {
  const [isSignIn, setIsSignIn] = useState(true);
  
  return (
    <div className="h-screen overflow-hidden relative flex justify-center items-center">
      <img src={backgound} className="absolute insert-0 w-screen z-0" />

      <div className="bg-white sign_form relative z-10">
        <img src={logo} alt="Logo" className="" style={{marginBottom: "5px"}} />
        {isSignIn ? <SignIn changeState={setIsSignIn} /> : <SignUp changeState={setIsSignIn} />}
        
      </div>
    </div>
  )
}