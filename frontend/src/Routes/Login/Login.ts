import { MAIL_FIELDNAME } from "../Register/Services/MailServices";
import { PASSWORD_FIELDNAME } from "../Register/Services/utils";
import { AUTH_CODE_FIELDNAME, EXPIRES_AT_FIELDNAME } from "./constants";

export async function Login(password: string, mail: string) {
  const body = JSON.stringify({
    [PASSWORD_FIELDNAME]: password,
    [MAIL_FIELDNAME]: mail,
  });

  const endpoint = "/api/user/login";

  const resp = await fetch(endpoint, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: body,
  }).catch((e: Error) => {
    throw e;
  });

  if (!resp.ok) {
    throw new Error("request returned non-200 code: " + resp.status);
  }
  const respBody =await resp.json()
  if ( !( AUTH_CODE_FIELDNAME in respBody) ){
    throw new Error(AUTH_CODE_FIELDNAME + " does not exist on response body")

  } 
  if (!(EXPIRES_AT_FIELDNAME in respBody)) {
    throw new Error(EXPIRES_AT_FIELDNAME + " does not exist on response body")
  }
  const authCode = respBody.AUTH_CODE_FIELDNAME
  window.localStorage.setItem(AUTH_CODE_FIELDNAME,JSON.stringify({
    [AUTH_CODE_FIELDNAME] : authCode,
    [EXPIRES_AT_FIELDNAME]: respBody.EXPIRES_AT_FIELDNAME
  }) )
  
  return authCode

  
}
