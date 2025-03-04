import { createHash } from "crypto";
import {
  ComponentDispatchStruct,
  ERR_MODE_STYLES,
  StatusCodes,
} from "./Register";

const PASSWORD_FIELDNAME = "passwordHash";

export function DoPasswordOperations(
  passwordCompStruct: ComponentDispatchStruct,
  passwordRepeatCompStruct: ComponentDispatchStruct
): string {
  const {
    compRef: passwd,
    setStyle: setPasswordStyles,
    setText: setPasswordText,
    originalStyle: ogPasswdStyle,
  } = passwordCompStruct;
  const {
    compRef: passwdRepeat,
    setStyle: setPasswordRepeatStyles,
    setText: setPasswordRepeatText,
    originalStyle: originalPasswordRepeatStyle,
  } = passwordRepeatCompStruct;

  setPasswordRepeatStyles(originalPasswordRepeatStyle);
  setPasswordStyles(ogPasswdStyle);

  if (passwd.current === null) {
    setPasswordText(StatusCodes.ErrNull);
    setPasswordStyles(ERR_MODE_STYLES);
    return "";
  }
  if (passwdRepeat.current === null) {
    setPasswordRepeatStyles(ERR_MODE_STYLES);
    setPasswordRepeatText(StatusCodes.ErrNull);
    return "";
  }
  const passwordContent = passwd.current.innerText;
  const passwordRepeatContent = passwdRepeat.current.innerText;

  const { code, data } = CheckPasswordValidity(
    passwordContent,
    passwordRepeatContent
  );
  if (code !== StatusCodes.Success) {
    setPasswordText(code);
    setPasswordStyles(ERR_MODE_STYLES);
    if (code === StatusCodes.ErrDontMatch) {
      setPasswordRepeatStyles(ERR_MODE_STYLES);

      setPasswordRepeatText(code);
    }
    return "";
  }
  window.localStorage.setItem(PASSWORD_FIELDNAME, data);
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
    return { code: StatusCodes.ErrDontMatch, data: "" };
  }
  return { code: StatusCodes.Success, data: CreatePasswordHash(p1) };
}

export function CreatePasswordHash(s1: string): string {
  const hashedPassword = createHash("sha256").update(s1).digest("hex");
  return hashedPassword;
}
