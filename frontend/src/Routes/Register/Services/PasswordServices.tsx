import { createHash } from "crypto";
import { StatusCodes } from "./Register";


export function DoPasswordOperations(
  passwd: HTMLDivElement | null,
  passwdRepeat: HTMLDivElement | null,
  setPasswordStyles: React.Dispatch<React.SetStateAction<string>>,
  setPasswordRepeatStyles: React.Dispatch<React.SetStateAction<string>>,
  setPasswordText: React.Dispatch<React.SetStateAction<string>>,
  setPasswordRepeatText: React.Dispatch<React.SetStateAction<string>>
): string {
  if (passwd === null) {
    return "";
  }
  if (passwdRepeat === null) {
    return "";
  }
  const passwordContent = passwd.innerText;
  const passwordRepeatContent = passwdRepeat.innerText;

  const { code, data } = CheckPasswordValidity(
    passwordContent,
    passwordRepeatContent
  );
  if (code !== StatusCodes.Success) {
    setPasswordText(code);
    setPasswordStyles(
      "bg-red-700 hover:text-white hover:bg-indigo-950 text-2xl text-white"
    );
    if (code === StatusCodes.ErrDontMatch) {
      setPasswordRepeatStyles(
        "bg-red-700 hover:text-white hover:bg-indigo-950 text-2xl text-white"
      );

      setPasswordRepeatText(code);
    }
    return "";
  }
  return data;
}


export function CheckPasswordValidity(p1: string, p2: string ): {
  code: StatusCodes;
  data: string;
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

