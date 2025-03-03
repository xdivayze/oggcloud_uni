export enum StatusCodes {
  Success = "",
  ErrWhiteSpace = "Password includes whitespace",
  ErrPasswordTooLong = "Password exceeds the 9 character password limit",
  ErrDontMatch = "Passwords don't match" ,
  ErrCryptographicFault = "Error occurred while processing your password",
  ErrMailMalformed = "E-Mail address structure is invalid",
  ErrNull = "Field is empty",
  ErrSecurityTextTooShort="Security text must be over 16 characters"
}

export interface IDoRegister {
  email: string;
  password: string; //hash before request
  securityText: string; //process into ecdhPrivate
  referralCode: string; //use to update used status of referral code
  ecdhPrivate: string;
}

export default function DoRegister(iDoRegister: IDoRegister) {
  //call after checking password repeat
  const jsonBody = {
    email: iDoRegister.email,
    referralCode: iDoRegister.referralCode,
  };
}

