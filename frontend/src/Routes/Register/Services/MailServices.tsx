import { ComponentDispatchStructType } from "../Components/ComponentDispatchStruct";
import { StatusCodes, ERR_MODE_STYLES } from "./utils";

export const MAIL_FIELDNAME = "email";

function CheckMailValidity(mailRef: HTMLDivElement | null): StatusCodes {
  if (mailRef === null) {
    return StatusCodes.ErrNull;
  }
  const mail = mailRef.innerText;

  if (mail.length < 5) {
    return StatusCodes.ErrMailMalformed;
  }

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
  mailCompStruct: ComponentDispatchStructType
): boolean {
  const mailRef = mailCompStruct.getRef();

  mailCompStruct.setStyles(mailCompStruct.originalStyles);
  const returnCode = CheckMailValidity(mailRef.current);
  if (returnCode !== StatusCodes.Success) {
    mailCompStruct.setStyles(ERR_MODE_STYLES);
    mailCompStruct.setText(returnCode);
    return false;
  }

  return true;
}
