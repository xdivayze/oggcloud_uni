import { Dispatch, RefObject, SetStateAction } from "react";
import { IDoRegister } from "./utils";

export interface ComponentDispatchStruct {
  setStyle: Dispatch<SetStateAction<string>>;
  setText: Dispatch<SetStateAction<string>>;
  compRef: RefObject<HTMLDivElement | null>;
  originalStyle: string;
}

export async function DoRegister(iDoRegister: IDoRegister) {
  const registerEndpoint = "/api/user/register";

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
  }).catch((e: Error) => {
    throw e;
  });

  //TODO navigate to page , then, show seed if success otherwise the respective status code
}
