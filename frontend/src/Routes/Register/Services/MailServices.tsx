
import { ComponentDispatchStruct} from "./Register";
import { StatusCodes, ERR_MODE_STYLES } from "./utils";

export const MAIL_FIELDNAME = "email"

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
  window.localStorage.setItem(MAIL_FIELDNAME, mail)
  return StatusCodes.Success;
}

export function DoCheckMailValidity(
  mailCompStruct: ComponentDispatchStruct
):boolean {
  
  const {compRef: mailRef, setStyle: setMailStyles, setText: setMailText, originalStyle} = mailCompStruct
  setMailStyles(originalStyle)
  const returnCode = CheckMailValidity(mailRef.current);
  if (returnCode !== StatusCodes.Success) {
    setMailStyles(
      ERR_MODE_STYLES
    );
    setMailText(returnCode);
    return false;
  }

  return true;
}
