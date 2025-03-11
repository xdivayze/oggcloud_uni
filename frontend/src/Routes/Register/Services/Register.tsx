import { Dispatch, JSX, RefObject, SetStateAction } from "react";
import { IDoRegister } from "./utils";
import { Buffer } from "buffer/";
export interface ComponentDispatchStruct {
  setStyle: Dispatch<SetStateAction<string>>;
  setText: Dispatch<SetStateAction<string>>;
  compRef: RefObject<HTMLDivElement | null>;
  originalStyle: string;
}

export async function DoRegister(
  iDoRegister: IDoRegister,
) {
  const registerEndpoint = "/api/user/register";

  const jsonBody = {
    email: iDoRegister.email,
    referralCode: iDoRegister.referralCode,
    password: iDoRegister.password,
    ecdhPublic: Buffer.from(iDoRegister.ecdhPublic).toString("hex"),
  };

  const req = await fetch(registerEndpoint, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(jsonBody),
  });
  return req.status
 
    
}
