import { createHash } from "crypto";
import {

} from "./Register";
import { ERR_MODE_STYLES, StatusCodes } from "./utils";
import  { ComponentDispatchStructType } from "../Components/ComponentDispatchStruct";

export function DoPasswordOperations(
  passwordCompStruct: ComponentDispatchStructType,
  passwordRepeatCompStruct: ComponentDispatchStructType
): string {

  const passwd = passwordCompStruct.getRef()
  const passwdRepeat = passwordRepeatCompStruct.getRef()

  passwordRepeatCompStruct.setStyles(passwordRepeatCompStruct.originalStyles);
  passwordCompStruct.setStyles(passwordCompStruct.originalStyles);

  if (passwd.current === null || passwordCompStruct.getRefContent().innerText === "" ) {
    passwordRepeatCompStruct.setText(StatusCodes.ErrNull);
    passwordCompStruct.setStyles(ERR_MODE_STYLES);
    return "";
  }
  if (passwdRepeat.current === null || passwordCompStruct.getRefContent().innerText === "") {
    passwordRepeatCompStruct.setStyles(ERR_MODE_STYLES);
    passwordRepeatCompStruct.setText(StatusCodes.ErrNull);
    return "";
  }
  const passwordContent = passwd.current.innerText;
  const passwordRepeatContent = passwdRepeat.current.innerText;

  const { code, data } = CheckPasswordValidity(
    passwordContent,
    passwordRepeatContent
  );
  if (code !== StatusCodes.Success) {
    passwordCompStruct.setText(code);
    passwordCompStruct.setStyles(ERR_MODE_STYLES);
    if (code === StatusCodes.ErrDontMatch) {
      console.error("err")
      passwordRepeatCompStruct.setStyles(ERR_MODE_STYLES);

      passwordRepeatCompStruct.setText(code);
    }
    return "";
  }
  return data;
}

export function CheckPasswordValidity(
  p1: string,
  p2: string
): {
  code: StatusCodes;
  data: string; //returns password hash
} {
  if (p1.length > 9) {
    return { code: StatusCodes.ErrPasswordTooLong, data: "" };
  }
  if (/\s/.test(p1)) {
    return { code: StatusCodes.ErrWhiteSpace, data: "" };
  }
  if (p1 !== p2) {
    console.error("no match")
    return { code: StatusCodes.ErrDontMatch, data: "" };
  }
  return { code: StatusCodes.Success, data: CreatePasswordHash(p1) };
}

export function CreatePasswordHash(s1: string): string {
  const hashedPassword = createHash("sha256").update(s1).digest("hex");
  return hashedPassword;
}
