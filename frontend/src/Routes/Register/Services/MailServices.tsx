import { Dispatch, SetStateAction } from "react";
import { StatusCodes } from "./Register";

function CheckMailValidity(mailRef: HTMLDivElement | null): StatusCodes {
  if (mailRef === null) {
    return StatusCodes.ErrNull;
  }
  const mail = mailRef.innerHTML;

  if (/\s/.test(mail)) {
    return StatusCodes.ErrWhiteSpace;
  }
  const atIndex = mail.indexOf("@");
  if (atIndex === -1) {
    return StatusCodes.ErrMailMalformed;
  }
  if (mail.lastIndexOf(".") < atIndex) {
    return StatusCodes.ErrMailMalformed;
  }
  return StatusCodes.Success;
}

export function DoCheckMailValidity(
  mailRef: HTMLDivElement | null,
  setMailText: Dispatch<SetStateAction<string>>,
  setMailStyles: Dispatch<SetStateAction<string>>
):boolean {
  const returnCode = CheckMailValidity(mailRef);
  if (returnCode !== StatusCodes.Success) {
    setMailStyles(
      "bg-red-700 hover:text-white hover:bg-indigo-950 text-2xl text-white"
    );
    setMailText(returnCode);
    return false;
  }
  return true;
}
