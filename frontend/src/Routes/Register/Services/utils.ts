export const ERR_MODE_STYLES =
  "bg-red-700 hover:text-white hover:bg-indigo-950 text-2xl text-white";

export interface IDoRegister {
  email: string;
  password: string; //hash before request
  referralCode: string; //use to update used status of referral code
  ecdhPublic: string;
  secText : string;
}


export enum StatusCodes {
  Success = "",
  ErrWhiteSpace = "Password includes whitespace",
  ErrPasswordTooLong = "Password exceeds the 9 character password limit",
  ErrDontMatch = "Passwords don't match" ,
  ErrCryptographicFault = "Error occurred while processing your password",
  ErrMailMalformed = "E-Mail address structure is invalid",
  ErrNull = "Field is empty",
  ErrSecurityTextTooShort="Security text must be over 16 characters",

}