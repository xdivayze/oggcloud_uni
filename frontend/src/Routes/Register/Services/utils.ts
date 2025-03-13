export const ERR_MODE_STYLES =
  "bg-red-700 hover:text-white hover:bg-indigo-950 text-2xl text-white";

export interface IDoRegister { //TODO change this with a class with methods to create states internally to get rid of redundant code
  email: string;
  password: string; //hash before request
  referralCode: string; //use to update used status of referral code
  ecdhPublic: string;
  secText : string;
}


export const ObeseBarDefaultStyles = "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"

export const REFERRAL_CODE_FIELDNAME = "referralCode"
export const EMAIL_FIELDNAME = "email"
export const PASSWORD_FIELDNAME = "password"
export const ECDH_PUB_FIELDNAME = "ecdhPublic"

export enum StatusCodes {
  Success = "",
  ErrWhiteSpace = "Can't include whitespace",
  ErrPasswordTooLong = "Password exceeds the 9 character password limit",
  ErrDontMatch = "Passwords don't match" ,
  ErrCryptographicFault = "Error occurred while processing your password",
  ErrMailMalformed = "E-Mail address structure is invalid",
  ErrNull = "Field is empty",
  ErrSecurityTextTooShort="Security text must be over 16 characters",

}