import { Dispatch, RefObject, SetStateAction } from "react";
import { IDoRegister } from "./utils";
import { useNavigate } from "react-router-dom";

export interface ComponentDispatchStruct {
  setStyle: Dispatch<SetStateAction<string>>;
  setText: Dispatch<SetStateAction<string>>;
  compRef: RefObject<HTMLDivElement | null>;
  originalStyle: string;
}

export async function DoRegister(iDoRegister: IDoRegister) {
  const registerEndpoint = "/api/user/register";
  const navigate = useNavigate()

  const jsonBody = {
    email: iDoRegister.email,
    referralCode: iDoRegister.referralCode,
    password: iDoRegister.password,
    ecdhPublic: iDoRegister.ecdhPublic,
  };

  const req = await fetch(registerEndpoint, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(jsonBody),
  })

  navigate("/register/post?code=" + req.status)


  

  //TODO navigate to page , then, show seed if success otherwise the respective status code
}
