import { Dispatch, RefObject, SetStateAction } from "react";
import { ECDH_PUB_FIELDNAME, EMAIL_FIELDNAME, IDoRegister, PASSWORD_FIELDNAME, REFERRAL_CODE_FIELDNAME } from "./utils";
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
    [EMAIL_FIELDNAME]: iDoRegister.email,
    [REFERRAL_CODE_FIELDNAME]: iDoRegister.referralCode,
    [PASSWORD_FIELDNAME]: iDoRegister.password,
    [ECDH_PUB_FIELDNAME]: Buffer.from(iDoRegister.ecdhPublic).toString("hex"),
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
