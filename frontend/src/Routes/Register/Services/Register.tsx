import { Dispatch, RefObject, SetStateAction } from "react";
import { IDoRegister } from "./utils";

export interface ComponentDispatchStruct {
  setStyle: Dispatch<SetStateAction<string>>;
  setText: Dispatch<SetStateAction<string>>;
  compRef: RefObject<HTMLDivElement | null>;
  originalStyle: string;
}

export function DoRegister(iDoRegister: IDoRegister) {
  const jsonBody = {
    email: iDoRegister.email,
    referralCode: iDoRegister.referralCode,
    password: iDoRegister.password,
    ecdhPublic: iDoRegister.ecdhPublic

  }; 
  
  //TODO implement register requests
}
